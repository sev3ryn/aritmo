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
	input []rune    // input
	start int       // item start
	col   int       // current end
	items chan item // output channel
	state stateFn
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

func (l *lexer) emitEOF() {
	l.items <- item{itemEOF, "", l.start}
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
		return lexAny
	} else {
		// catch identifier declaration. First symbol is letter - else alphanumeric
		l.next()
		l.iterate(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' })
		l.emit(itemIdentifier)

		l.skipSpaces()

		if l.peek() == '=' {
			l.next()
			l.emit(itemEqual)
		}
	}
	return lexAny

}

func lexAny(l *lexer) stateFn {
	// LOOP:
	for {
		r := l.peek()
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
		case r == '(':
			l.next()
			l.emit(itemLParen)
		case r == ')':
			l.next()
			l.emit(itemRParen)
		case isSpace(r):
			l.skipSpaces()
		case r == '\n' || r == eof:
			l.next()
			l.emitEOF()
			return nil
		case unicode.IsDigit(r):
			return lexNumber
		default:
			l.errorf("Unexpected char - %q at col %d", r, l.start)
		}
	}

}

func lexNumber(l *lexer) stateFn {

	l.iterate(unicode.IsDigit)

	if l.peek() == '.' {
		l.next()

		if r := l.peek(); !unicode.IsDigit(r) {
			l.errorf("Unexpected char after dot %q at col %d", r, l.col)
		}
		l.iterate(unicode.IsDigit)
	}

	l.emit(itemNumber)
	return lexAny
}
