package main

import (
	"math"
	"strconv"

	"fmt"

	"github.com/sev3ryn/aritmo/frontend/core/datatype"
	"github.com/sev3ryn/aritmo/frontend/core/parse"
	"github.com/sev3ryn/aritmo/frontend/core/scan"
	"github.com/sev3ryn/aritmo/frontend/core/storage"
)

const precision = 2

func roundFloat(x float64, prec int) float64 {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return x
	}

	sign := 1.0
	if x < 0 {
		sign = -1
		x *= -1
	}

	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)

	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow * sign
}

var store = storage.RAMStore
var currUpdCh = make(chan []byte)
var typeMap = datatype.Init(currUpdCh)

func calculate(line int, input string) string {
	store.CurrLine = line
	p := parse.New(scan.New(input), store, typeMap)
	output, err := p.ExecStatement()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	output.Val = roundFloat(output.Val, precision)

	return fmt.Sprintf("%s %s", strconv.FormatFloat(output.Val, 'f', -1, 64), output.Typ.GetBase().DisplayName)

}

func main() {

}
