package main

import (
	"math"
	"strconv"

	"fmt"

	"github.com/sev3ryn/aritmo/datatype"
	"github.com/sev3ryn/aritmo/parse"
	"github.com/sev3ryn/aritmo/scan"
	"github.com/sev3ryn/aritmo/storage"
)

const precision = 2

func RoundFloat(x float64, prec int) float64 {
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
var typeMap = datatype.Init()

func calculate(line int, input string) string {
	store.CurrLine = line
	p := parse.New(scan.New(input), store, typeMap)
	output, err := p.ExecStatement()
	if err != nil {
		return ""
	}

	output.Val = RoundFloat(output.Val, precision)

	return fmt.Sprintf("%s %s", strconv.FormatFloat(output.Val, 'f', -1, 64), output.Typ.GetBase().DisplayName)

}

func main() {

}
