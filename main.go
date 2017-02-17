package main

import (
	"math"
	"strconv"

	"github.com/sev3ryn/aritmo/parse"
	"github.com/sev3ryn/aritmo/scan"
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

func calculate(input string) string {
	p := parse.New(scan.New(input))
	output, err := p.ExecStatement()
	if err != nil {
		return ""
	}

	output = RoundFloat(output, precision)

	return strconv.FormatFloat(output, 'f', -1, 64)

}

func main() {

}
