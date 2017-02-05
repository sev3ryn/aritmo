package parse

import (
	"strconv"

	"github.com/sev3ryn/aritmo/scan"
)

type Result struct {
	val float64
}

type Parser struct {
	tokens []scan.Item
}

func (p *Parser) peek() scan.Item {
	return p.tokens[0]
}

func (p *Parser) next() scan.Item {
	if len(p.tokens) > 1 {
		p.tokens = p.tokens[1:]
	}
	return p.tokens[0]
}

func (p *Parser) exec() float64 {

	var f func(float64, float64) float64
	var itmOp scan.Item
	itmLeft := p.peek()
	valLeft, _ := strconv.ParseFloat(itmLeft.Val, 64)
	for {
		itmOp = p.next()
		switch itmOp.Typ {
		case scan.ItemAdd:
			f = func(a, b float64) float64 { return a + b }
		case scan.ItemSub:
			f = func(a, b float64) float64 { return a - b }
		case scan.ItemMul:
			f = func(a, b float64) float64 { return a * b }
		case scan.ItemDiv:
			f = func(a, b float64) float64 { return a / b }
		case scan.ItemEOF:
			return valLeft
		default:
			panic("yo")
		}

		itmLeft = p.next()
		tmpVal, _ := strconv.ParseFloat(itmLeft.Val, 64)

		valLeft = f(valLeft, tmpVal)

	}
	return valLeft

}

func New(s *scan.Scanner) *Parser {
	//p.tokens = p.tokens[:0]
	p := &Parser{tokens: []scan.Item{}}
	for {
		tok := s.NextItem()
		switch tok.Typ {
		case scan.ItemError, scan.ItemEOF:
			p.tokens = append(p.tokens, tok)
			return p
		default:
			p.tokens = append(p.tokens, tok)
		}
	}
}

// func (p *parser) peek()  {

// 	return
// }

// func (p *parser) run() {
// 	if p.token.Typ == scan.ItemVariable {
// 		val, ok := variables[p.token.val]
// 		if !ok {
// 			return "expect = "
// 		}
// 		return val
// 	}

// 	if p.token.Typ == scan.ItemOperator

// }
