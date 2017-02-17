package parse

import (
	"fmt"
	"strconv"

	"github.com/sev3ryn/aritmo/scan"
	"github.com/sev3ryn/aritmo/storage"
)

type Result struct {
	val float64
}

type Parser struct {
	tokens []scan.Item
	store  storage.Store
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
		return p.execStatement()
	case scan.ItemVariable:
		return storage.RAMStore.Get(tok.Val)
	default:
		return getVal(tok)
	}

}

func (p *Parser) getOperationFn(i scan.Item) (operation, error) {
	if i.Typ == scan.ItemError {
		return nil, fmt.Errorf(i.Val)
	} else if i.Typ == scan.ItemEOF {
		return nil, nil
	} else if i.Typ == scan.ItemRParen { // end of substatement
		return nil, nil
		// } else if i.Typ == scan.ItemTypeConv {
		// 	return &operation{
		// 		f:           func(a, _ float64) (float64, error) { return 0, nil },
		// 		precendance: 5,
		// 	}, nil
	} else if i.Typ == scan.ItemOperation {

		op, ok := operationMap[i.Val]
		if !ok {
			panic("Unsupported operation. Something went wrong")
		}
		return op, nil
	}
	panic("Unsupported operation")
}

func (p *Parser) execOperation(valLeft float64, op operation) (float64, error) {

	valRight, err := p.getOperand(p.next())
	if err != nil {
		return 0, err
	}

	nextOp, err := p.getOperationFn(p.next())
	if err != nil {
		return 0, err
	}

	if nextOp == nil {
		return op.Exec([]float64{valLeft, valRight})
	}

	if nextOp.GetPrec() > op.GetPrec() {
		valRight, err = p.execOperation(valRight, nextOp)
		if err != nil {
			return 0, err
		}
		return op.Exec([]float64{valLeft, valRight})
	}

	valLeft, err = op.Exec([]float64{valLeft, valRight})
	if err != nil {
		return 0, err
	}
	return p.execOperation(valLeft, nextOp)

}

func getVal(i scan.Item) (float64, error) {
	return strconv.ParseFloat(i.Val, 64)
}

func (p *Parser) execStatement() (float64, error) {

	v, err := p.getOperand(p.peek())

	if err != nil {
		return 0, err
	}

	f, err := p.getOperationFn(p.next())
	if err != nil {
		return 0, err
	} else if f == nil {
		return v, nil
	}
	return p.execOperation(v, f)
}

func (p *Parser) ExecStatement() (float64, error) {

	// case of assignment
	if len(p.tokens) > 2 {

		firstTok := p.tokens[0]
		secondTok := p.tokens[1]
		if firstTok.Typ == scan.ItemVariable && secondTok.Typ == scan.ItemEqual {
			p.next()
			p.next()

			v, err := p.execStatement()
			// IO errors on storage should crash application
			ioErr := p.store.Save(firstTok.Val, v, err)
			if ioErr != nil {
				panic(ioErr)
			}

			return v, err
		}
	}

	return p.execStatement()
}

// New - costructor for Parser
func New(s *scan.Scanner, st storage.Store) *Parser {
	//p.tokens = p.tokens[:0]
	p := &Parser{tokens: []scan.Item{}, store: st}
	for {
		tok := s.NextItem()
		switch tok.Typ {
		case scan.ItemError:
			p.tokens = append(p.tokens, tok)
			return p
		case scan.ItemEOF:
			p.tokens = append(p.tokens, tok)
			return p
		default:
			p.tokens = append(p.tokens, tok)
		}
	}
}
