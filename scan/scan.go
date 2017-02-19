package scan

import (
	"fmt"
	"unicode"
)

type itemType int

//go:generate stringer -type=itemType
const (
	ItemEOF itemType = iota
	ItemNumber
	ItemVariable
	ItemEqual
	ItemOperation
	ItemLParen
	ItemRParen
	ItemDataType
	ItemBareDataType
	ItemError
	ItemTypeConv
)

// Item atomic piece of scanner output. May be number, operator or variable
type Item struct {
	Typ itemType
	Val string
	col int
}

func (i Item) String() string {
	return fmt.Sprintf("%s:%q", i.Typ, i.Val)
}

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

type stateFn func(*Scanner) stateFn

// Scanner - holds current state of scan
type Scanner struct {
	input        []rune    // input
	start        int       // item start
	col          int       // current end
	items        chan Item // output channel
	state        stateFn   // function for current execution
	openParenCnt int       // number of opened parenteses
}

// NextItem - returns next emited item by scanner
func (s *Scanner) NextItem() Item {
	item := <-s.items
	return item
}

// New - costructor for Scanner
func New(input string) *Scanner {
	s := &Scanner{
		input: []rune(input),
		start: 0,
		col:   0,
		items: make(chan Item),
	}
	go s.run()
	return s
}

func (s *Scanner) run() {
	for s.state = lexStart; s.state != nil; {
		s.state = s.state(s)
	}
}

// peek - return current rune of input
func (s *Scanner) peek() rune {

	if len(s.input) < s.col+1 {
		return eof
	}

	return s.input[s.col]
}

// next - return next rune of input
func (s *Scanner) next() rune {

	s.col++
	return s.peek()
}

func (s *Scanner) currentStr() string {
	return string(s.input[s.start:s.col])
}

// emit - puts item to output channel and shifts start for next iteration
func (s *Scanner) emit(t itemType) {
	s.items <- Item{t, s.currentStr(), s.start}
	s.start = s.col
}

// errorf - emits itemError
func (s *Scanner) errorf(format string, args ...interface{}) stateFn {
	s.items <- Item{ItemError, fmt.Sprintf(format, args...), s.start}
	return nil
}

// emitEOF - emits EOF if it is expected and all parenteses are closed otherwise - error
func (s *Scanner) emitEOF(unexpected bool) stateFn {
	if s.openParenCnt != 0 || unexpected {
		return s.unexpectedErr("EOF")
	}

	s.items <- Item{ItemEOF, "", s.col}
	return nil
}

func (s *Scanner) unexpectedErr(typ string) stateFn {
	return s.errorf("Unexpected %s - at col %d", typ, s.col)
}

// iterate - move through runes until whileIsTrue returns true for the specific rune
func (s *Scanner) iterate(whileIsTrue func(rune) bool) {
	for r := s.peek(); whileIsTrue(r); r = s.next() {
	}
}

// skipSpaces - ignore spaces
func (s *Scanner) skipSpaces() {
	s.iterate(func(r rune) bool { return r == ' ' || r == '\t' })
	s.start = s.col
}

func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

func isEOF(r rune) bool {
	return r == '\n' || r == eof
}

func lexStart(s *Scanner) stateFn {
	if r := s.peek(); isSpace(r) {
		s.skipSpaces()
		return s.state
	} else if !unicode.IsLetter(r) {
		return lexOperand
	} else {
		// catch identifier declaration. First symbol is letter - else alphanumeric
		lexVariable(s)

		s.skipSpaces()

		if s.peek() == '=' {
			s.next()
			s.emit(ItemEqual)
			return lexOperand
		}
		return lexOperator
	}
}

func lexOperand(s *Scanner) stateFn {

	if isSpace(s.peek()) {
		s.skipSpaces()
	}

	r := s.peek()
	if isEOF(r) {
		return s.emitEOF(true)
	}

	switch {
	case unicode.IsDigit(r):
		lexNumber(s)
		return lexDataType
	// handling negative numbers
	case r == '-':

		if r = s.next(); isEOF(r) {
			return s.emitEOF(true)
		} else if unicode.IsDigit(r) {
			lexNumber(s)
			return lexDataType
		}

		return s.unexpectedErr("char")

	case unicode.IsLetter(r):
		return lexVariable
	case r == '(':
		s.openParenCnt++
		s.next()
		s.emit(ItemLParen)
		return lexOperand
	default:
		return s.unexpectedErr("char")
	}

}

func lexDataType(s *Scanner) stateFn {
	if isSpace(s.peek()) {
		s.skipSpaces()
	}

	r := s.peek()
	if isEOF(r) {
		s.emit(ItemBareDataType)
		return s.emitEOF(false)
	}

	if !unicode.IsLetter(r) {
		s.emit(ItemBareDataType)
		return lexOperator
	}

	s.iterate(unicode.IsLetter)

	s.emit(ItemDataType)

	return lexOperator
}

func lexOperator(s *Scanner) stateFn {

	if isSpace(s.peek()) {
		s.skipSpaces()
	}

	r := s.peek()
	if isEOF(r) {
		return s.emitEOF(false)
	}

	switch r {
	case '+', '-', '*', '/':
		s.next()
		s.emit(ItemOperation)
	case 't':
		if s.next() != 'o' {
			s.unexpectedErr("char")
		}
		s.next()
		s.emit(ItemTypeConv)
	case ')':
		if s.openParenCnt == 0 {
			return s.errorf("No matching opening parenteses for closing one at col %d", s.col)
		}
		s.openParenCnt--
		s.next()
		s.emit(ItemRParen)
		return lexOperator
	default:
		return s.unexpectedErr("char")
	}

	return lexOperand
}

func lexNumber(s *Scanner) stateFn {
	s.next()
	s.iterate(unicode.IsDigit)

	if s.peek() == '.' {
		s.next()

		if r := s.peek(); !unicode.IsDigit(r) {
			return s.unexpectedErr("char")
		}
		s.iterate(unicode.IsDigit)
	}

	s.emit(ItemNumber)
	return lexOperator
}

func lexVariable(s *Scanner) stateFn {
	// catch identifier declaration. First symbol is letter - else alphanumeric
	s.next()
	s.iterate(func(r rune) bool { return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' })
	s.emit(ItemVariable)
	return lexOperator
}
