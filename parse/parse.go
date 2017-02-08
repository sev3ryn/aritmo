package parse

import (
	"fmt"
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

func (p *Parser) getOperand(tok scan.Item) (float64, error) {
	switch tok.Typ {
	case scan.ItemError:
		return 0, fmt.Errorf(tok.Val)
	case scan.ItemLParen:
		// calculate statement in parenteses as new statement
		p.next()
		return p.ExecStatement()
	default:
		return getVal(tok)
	}

}

type operation struct {
	f           func(float64, float64) (float64, error)
	precendance int
}

func (p *Parser) getOperationFn(i scan.Item) (*operation, error) {
	switch i.Typ {
	case scan.ItemAdd:
		return &operation{
			f:           func(a, b float64) (float64, error) { return a + b, nil },
			precendance: 1,
		}, nil
	case scan.ItemSub:
		return &operation{
			f:           func(a, b float64) (float64, error) { return a - b, nil },
			precendance: 1,
		}, nil
	case scan.ItemMul:
		return &operation{
			f:           func(a, b float64) (float64, error) { return a * b, nil },
			precendance: 10,
		}, nil

	case scan.ItemDiv:
		return &operation{
			f: func(a, b float64) (float64, error) {
				if b == 0 {
					return 0, fmt.Errorf("Can't divide by zero")
				}
				return a / b, nil
			},
			precendance: 10,
		}, nil
	case scan.ItemError:
		return nil, fmt.Errorf(i.Val)
	case scan.ItemEOF:
		return nil, nil
	case scan.ItemRParen: // end of substatement
		return nil, nil
	default: //should never happen
		panic("Unsupported operation")
	}
}

func (p *Parser) execOperation(valLeft float64, op *operation) (float64, error) {

	valRight, err := p.getOperand(p.next())
	if err != nil {
		return 0, err
	}

	nextOp, err := p.getOperationFn(p.next())
	if err != nil {
		return 0, err
	}

	if nextOp == nil {
		return op.f(valLeft, valRight)
	}

	if nextOp.precendance > op.precendance {
		valRight, err = p.execOperation(valRight, nextOp)
		if err != nil {
			return 0, err
		}
		return op.f(valLeft, valRight)
	}

	valLeft, err = op.f(valLeft, valRight)
	if err != nil {
		return 0, err
	}
	return p.execOperation(valLeft, nextOp)

}

func getVal(i scan.Item) (float64, error) {
	return strconv.ParseFloat(i.Val, 64)
}

func (p *Parser) exe() (float64, error) {
	v, err := p.getOperand(p.peek())

	if err != nil {
		return 0, err
	}

	f, err := p.getOperationFn(p.next())
	if f == nil {
		return v, nil
	} else if err != nil {
		return 0, err
	}
	return p.execOperation(v, f)

}

func (p *Parser) ExecStatement() (float64, error) {
	//i++
	//z := i
	v, err := p.exe()
	//fmt.Printf("%d: %f\n", z, v)
	return v, err
}

// New - costructor for Parser
func New(s *scan.Scanner) *Parser {
	//p.tokens = p.tokens[:0]
	p := &Parser{tokens: []scan.Item{}}
	for {
		tok := s.NextItem()
		switch tok.Typ {
		case scan.ItemError:
			p.tokens = []scan.Item{tok}
			return p
		case scan.ItemEOF:
			p.tokens = append(p.tokens, tok)
			return p
		default:
			p.tokens = append(p.tokens, tok)
		}
	}
}
