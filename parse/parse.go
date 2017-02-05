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

func (p *Parser) exec2(valLeft float64, op operation, pr int) float64 {

	tmpLeft := getVal(p.next())
	f, nextPr := getOpFunc(p.next())

	if f == nil {
		return op(valLeft, tmpLeft)
	}

	if nextPr > pr {
		return op(valLeft, p.exec2(tmpLeft, f, nextPr))
	}

	return p.exec2(op(valLeft, tmpLeft), f, nextPr)

}

func getVal(i scan.Item) float64 {
	v, _ := strconv.ParseFloat(i.Val, 64)
	return v
}

type operation func(float64, float64) float64

func getOpFunc(i scan.Item) (operation, int) {
	switch i.Typ {
	case scan.ItemAdd:
		return func(a, b float64) float64 { return a + b }, 1
	case scan.ItemSub:
		return func(a, b float64) float64 { return a - b }, 1
	case scan.ItemMul:
		return func(a, b float64) float64 { return a * b }, 10
	case scan.ItemDiv:
		return func(a, b float64) float64 { return a / b }, 10
	case scan.ItemEOF:
		return nil, 0
	default:
		panic("yo")
	}
}

func (p *Parser) exec() float64 {
	v := getVal(p.peek())
	f, pr := getOpFunc(p.next())
	return p.exec2(v, f, pr)

}

// func (p *Parser) exec() float64 {

// 	var f operation
// 	valLeft := getVal(p.peek())

// 	for {

// 		f, _ = getOpFunc(p.next())

// 		if f == nil {
// 			return valLeft
// 		}

// 		tmpVal := getVal(p.next())

// 		valLeft = f(valLeft, tmpVal)

// 	}
// 	// unreachable
// 	return 0
// }

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
