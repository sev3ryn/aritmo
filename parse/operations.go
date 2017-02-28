package parse

import (
	"fmt"

	"math"

	"github.com/sev3ryn/aritmo/datatype"
	"github.com/sev3ryn/aritmo/storage"
)

type binaryFn func(storage.Result, storage.Result) (storage.Result, error)

type binaryOp struct {
	f           binaryFn
	precendance int
}

func (op *binaryOp) Exec(m []storage.Result) (storage.Result, error) {
	return op.f(m[0], m[1])
}

func (op *binaryOp) GetPrec() int {
	return op.precendance
}

type unaryFn func(storage.Result) (storage.Result, error)

type unaryOp struct {
	f           binaryFn
	precendance int
}

func (op *unaryOp) Exec(m []storage.Result) (storage.Result, error) {
	return op.f(m[0], storage.Result{})
}

func (op *unaryOp) GetPrec() int {
	return op.precendance
}

type operation interface {
	Exec([]storage.Result) (storage.Result, error)
	GetPrec() int
}

var operationMap = map[string]operation{
	"+":   &binaryOp{f: Add, precendance: 1},
	"-":   &binaryOp{f: Sub, precendance: 1},
	"*":   &binaryOp{f: Mul, precendance: 10},
	"/":   &binaryOp{f: Div, precendance: 10},
	"to":  &binaryOp{f: Conv, precendance: 0},
	"sin": &unaryOp{f: Conv, precendance: 100},
}

func Add(a, b storage.Result) (r storage.Result, err error) {
	return calcResult(a, b, func(a, b float64) float64 { return a + b })
}

func Sub(a, b storage.Result) (storage.Result, error) {
	return calcResult(a, b, func(a, b float64) float64 { return a - b })
}

func Mul(a, b storage.Result) (storage.Result, error) {
	return calcResult(a, b, func(a, b float64) float64 { return a * b })
}

func Div(a, b storage.Result) (storage.Result, error) {
	if b.Val == 0 {
		return storage.Result{}, fmt.Errorf("Can't divide by zero")
	}
	return calcResult(a, b, func(a, b float64) float64 { return a / b })
}

func Conv(a, b storage.Result) (storage.Result, error) {
	f, err := a.Typ.GetConvFunc(b.Typ)
	if err != nil {
		return storage.Result{}, err
	}
	return storage.Result{Val: f(a.Val), Typ: b.Typ}, err
}

func Sin(a storage.Result) (r storage.Result, err error) {
	return storage.Result{Val: math.Sin(a.Val), Typ: datatype.BareDataType}, nil
}

func calcResult(a, b storage.Result, opFunc func(float64, float64) float64) (storage.Result, error) {
	if a.Typ.GetBase().Group == datatype.GroupBare {
		return storage.Result{Val: opFunc(a.Val, b.Val), Typ: b.Typ}, nil
	}
	f, err := b.Typ.GetConvFunc(a.Typ)
	if err != nil {
		return storage.Result{}, err
	}
	return storage.Result{Val: opFunc(a.Val, f(b.Val)), Typ: a.Typ}, err
}
