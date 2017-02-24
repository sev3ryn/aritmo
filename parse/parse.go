package parse

import (
	"fmt"
	"strconv"

	"github.com/sev3ryn/aritmo/datatype"
	"github.com/sev3ryn/aritmo/scan"
	"github.com/sev3ryn/aritmo/storage"
)

type Parser struct {
	tokens  []scan.Item
	store   storage.Store
	typeMap datatype.TypeMap
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

func (p *Parser) getOperand(tok scan.Item) (r storage.Result, err error) {
	switch tok.Typ {
	case scan.ItemError:
		return storage.Result{}, fmt.Errorf(tok.Val)
	case scan.ItemLParen:
		// calculate statement in parenteses as new statement
		p.next()
		return p.execStatement()
	case scan.ItemVariable:
		return p.store.Get(tok.Val)
	default:
		//r = storage.Result{}
		r.Val, err = getVal(tok)
		if err != nil {
			return
		}
		r.Typ, err = p.getTyp(p.next())

		return r, err
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

func (p *Parser) execOperation(valLeft storage.Result, op operation) (storage.Result, error) {

	valRight, err := p.getOperand(p.next())
	if err != nil {
		return storage.Result{}, err
	}

	nextOp, err := p.getOperationFn(p.next())
	if err != nil {
		return storage.Result{}, err
	}

	if nextOp == nil {
		return op.Exec([]storage.Result{valLeft, valRight})
	}

	if nextOp.GetPrec() > op.GetPrec() {
		valRight, err = p.execOperation(valRight, nextOp)
		if err != nil {
			return storage.Result{}, err
		}
		return op.Exec([]storage.Result{valLeft, valRight})
	}

	valLeft, err = op.Exec([]storage.Result{valLeft, valRight})
	if err != nil {
		return storage.Result{}, err
	}
	return p.execOperation(valLeft, nextOp)

}

func getVal(i scan.Item) (float64, error) {
	return strconv.ParseFloat(i.Val, 64)
}

func (p *Parser) getTyp(i scan.Item) (datatype.DataType, error) {
	if i.Typ == scan.ItemBareDataType {
		return datatype.BareDataType, nil
	} else if i.Typ == scan.ItemDataType {
		return p.typeMap.GetType(i.Val)
	}

	panic("unexpected behaviour of scan - expected ItemDataType not found")

}

func (p *Parser) execStatement() (storage.Result, error) {

	v, err := p.getOperand(p.peek())

	if err != nil {
		return storage.Result{}, err
	}

	f, err := p.getOperationFn(p.next())
	if err != nil {
		return storage.Result{}, err
	} else if f == nil {
		return v, nil
	}
	return p.execOperation(v, f)
}

func (p *Parser) ExecStatement() (storage.Result, error) {

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
func New(s *scan.Scanner, st storage.Store, tm datatype.TypeMap) *Parser {
	//p.tokens = p.tokens[:0]
	p := &Parser{tokens: []scan.Item{}, store: st, typeMap: tm}
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
