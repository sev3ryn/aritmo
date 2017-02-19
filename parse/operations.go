package parse

import (
	"fmt"

	"github.com/sev3ryn/aritmo/datatype"
	"github.com/sev3ryn/aritmo/storage"
)

type OperationFn func(storage.Result, storage.Result) (storage.Result, error)

type binaryOp struct {
	f           OperationFn
	precendance int
}

func (op *binaryOp) Exec(m []storage.Result) (storage.Result, error) {
	return op.f(m[0], m[1])
}

func (op *binaryOp) GetPrec() int {
	return op.precendance
}

type operation interface {
	Exec([]storage.Result) (storage.Result, error)
	GetPrec() int
}

var operationMap = map[string]operation{
	"+": &binaryOp{f: Add, precendance: 1},
	"-": &binaryOp{f: Sub, precendance: 1},
	"*": &binaryOp{f: Mul, precendance: 10},
	"/": &binaryOp{f: Div, precendance: 10},
}

func Add(a, b storage.Result) (r storage.Result, err error) {
	if a.Typ.Group == datatype.GroupBare {
		return storage.Result{a.Val + b.Val, b.Typ}, nil
	}
	m, err := b.Typ.GetConversionMultipl(a.Typ)
	return storage.Result{Val: a.Val + b.Val*m, Typ: a.Typ}, err
}

func Sub(a, b storage.Result) (storage.Result, error) {
	if a.Typ.Group == datatype.GroupBare {
		return storage.Result{a.Val - b.Val, b.Typ}, nil
	}
	m, err := b.Typ.GetConversionMultipl(a.Typ)
	return storage.Result{Val: a.Val - b.Val*m, Typ: a.Typ}, err
}

func Mul(a, b storage.Result) (storage.Result, error) {
	if a.Typ.Group == datatype.GroupBare {
		return storage.Result{a.Val * b.Val, b.Typ}, nil
	}
	m, err := b.Typ.GetConversionMultipl(a.Typ)
	return storage.Result{Val: a.Val * b.Val * m, Typ: a.Typ}, err
}

func Div(a, b storage.Result) (storage.Result, error) {
	if b.Val == 0 {
		return storage.Result{}, fmt.Errorf("Can't divide by zero")
	}
	if a.Typ.Group == datatype.GroupBare {
		return storage.Result{a.Val / b.Val, b.Typ}, nil
	}
	m, err := b.Typ.GetConversionMultipl(a.Typ)
	return storage.Result{Val: a.Val / (b.Val * m), Typ: a.Typ}, err
}
