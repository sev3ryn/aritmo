package parse

import "fmt"

type OperationFn func(float64, float64) (float64, error)

type binaryOp struct {
	f           OperationFn
	precendance int
}

func (op *binaryOp) Exec(m []float64) (float64, error) {
	return op.f(m[0], m[1])
}

func (op *binaryOp) GetPrec() int {
	return op.precendance
}

type operation interface {
	Exec([]float64) (float64, error)
	GetPrec() int
}

var operationMap = map[string]operation{
	"+": &binaryOp{f: Add, precendance: 1},
	"-": &binaryOp{f: Sub, precendance: 1},
	"*": &binaryOp{f: Mul, precendance: 10},
	"/": &binaryOp{f: Div, precendance: 10},
}

func Add(a, b float64) (float64, error) { return a + b, nil }
func Sub(a, b float64) (float64, error) { return a - b, nil }
func Mul(a, b float64) (float64, error) { return a * b, nil }
func Div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Can't divide by zero")
	}
	return a / b, nil
}
