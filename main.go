package main

import (
	"strconv"

	"github.com/sev3ryn/aritmo/parse"
	"github.com/sev3ryn/aritmo/scan"
)

const precision = 2

func RoundFloat(x float64, prec int) float64 {
	f, _ := strconv.ParseFloat(strconv.FormatFloat(x, 'g', prec, 64), 64)
	return f
}

func calculate(input string) string {
	p := parse.New(scan.New(input))
	output, err := p.ExecStatement()
	if err != nil {
		return ""
	}

	output = RoundFloat(output, precision+1)

	return strconv.FormatFloat(output, 'f', -1, 64)

}

func main() {

}
