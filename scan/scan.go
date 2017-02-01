package scan

import (
	"fmt"
	"unicode"
)

type item struct {
	typ itemType
	val string
	col int
}

func (i item) String() string {
	return fmt.Sprintf("{%s:%q}", i.typ, i.val)
}

type itemType int

//go:generate stringer -type=itemType
const (
	itemEOF itemType = iota
	itemNumber
	itemIdentifier
	itemEqual
	itemAdd
	itemSub
	itemMul
	itemDiv
	itemLParen
	itemRParen
	itemError
)

type currencyType int

//go:generate stringer -type=currencyType
const (
	nilType currencyType = iota
	usd
	eur
	pln
	gbp
	chf
)

const eof = -1

type stateFn func(*lexer) stateFn

type lexer struct {
	input        []rune    // input
	start        int       // item start
	col          int       // current end
	items        chan item // output channel
	state        stateFn
	openParenCnt int // number of opened parenteses
}

func (l *lexer) nextItem() item {
	item := <-l.items
	return item
}

func lex(input string) *lexer {
	l := &lexer{
		input: []rune(input),
		start: 0,
		col:   0,
		items: make(chan item),
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) next() rune {

	l.col++
	return l.peek()
}

func (l *lexer) peek() rune {

	if len(l.input) < l.col+1 {
		return eof
	}

	return l.input[l.col]
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, string(l.input[l.start:l.col]), l.start}
	l.start = l.col
}

func (l *lexer) emitEOF(unexpected bool) stateFn {
	if l.openParenCnt != 0 || unexpected {
		return l.errorf("Unexpected EOF - at col %d", l.col)
	}

	l.items <- item{itemEOF, "", l.col}
	return nil
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, fmt.Sprintf(format, args...), l.start}
	return nil
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func (l *lexer) iterate(whileIsTrue func(rune) bool) {
	for r := l.peek(); whileIsTrue(r); r = l.next() {
	}
}

func (l *lexer) skipSpaces() {
	l.iterate(func(r rune) bool { return r == ' ' || r == '\t' })
	l.start = l.col
}

func lexStart(l *lexer) stateFn {
	if r := l.peek(); isSpace(r) {
		l.skipSpaces()
		return l.state
	} else if !unicode.IsLetter(r) {
		return lexOperand
	} else {
		// catch identifier declaration. First symbol is letter - else alphanumeric
		lexVariable(l)

		l.skipSpaces()

		if l.peek() == '=' {
			l.next()
			l.emit(itemEqual)
			return lexOperand
		}
		return lexOperator
	}

}

func isEOF(r rune) bool {
	return r == '\n' || r == eof
}

func lexOperand(l *lexer) stateFn {

	if isSpace(l.peek()) {
		l.skipSpaces()
	}

	r := l.peek()
	if isEOF(r) {
		return l.emitEOF(true)
	}

	switch {
	case unicode.IsDigit(r):
		return lexNumber
	case unicode.IsLetter(r):
		return lexVariable
	case r == '(':
		l.openParenCnt++
		l.next()
		l.emit(itemLParen)
		return lexOperand
	default:
		return l.errorf("Unexpected char - at col %d", l.col)
	}

}

func lexOperator(l *lexer) stateFn {

	if isSpace(l.peek()) {
		l.skipSpaces()
	}

	r := l.peek()
	if isEOF(r) {
		return l.emitEOF(false)
	}

	switch {
	case r == '+':
		l.next()
		l.emit(itemAdd)
	case r == '-':
		l.next()
		l.emit(itemSub)
	case r == '*':
		l.next()
		l.emit(itemMul)
	case r == '/':
		l.next()
		l.emit(itemDiv)
	// case r == '(':
	// 	l.openParenCnt++
	// 	l.next()
	// 	l.emit(itemLParen)
	case r == ')':
		if l.openParenCnt == 0 {
			return l.errorf("No matching opening parenteses for closing one at col %d", l.col)
		}
		l.openParenCnt--
		l.next()
		l.emit(itemRParen)
		return lexOperator
	default:
		return l.errorf("Unexpected char - at col %d", l.col)
	}

	return lexOperand
}

func lexNumber(l *lexer) stateFn {
	l.next()
	l.iterate(unicode.IsDigit)

	if l.peek() == '.' {
		l.next()

		if r := l.peek(); !unicode.IsDigit(r) {
			return l.errorf("Unexpected char - at col %d", l.col)
		}
		l.iterate(unicode.IsDigit)
	}

	l.emit(itemNumber)
	return lexOperator
}

func lexVariable(l *lexer) stateFn {
	// catch identifier declaration. First symbol is letter - else alphanumeric
	l.next()
	l.iterate(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' })
	l.emit(itemIdentifier)
	return lexOperator
}
