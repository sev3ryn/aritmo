package parse

import (
	"testing"

	"github.com/sev3ryn/aritmo/scan"
)

type parseTest struct {
	input  string
	result float64
}

var parseTests = []parseTest{
	{"1+2+3+4", 10},
	{"1+2*3+4", 11},
	{"1+2*(3+4)", 15},
	{"1+2*(3+4)-6", 9},
	{"1+2*(3+4)/2-1", 7},
	{"(1+2)*3", 9},
	{"1+2*(3+4)-5*6", -15},
	{"1+2*(3+4)-5*(6+7)", -50},
	{"1+2*((3+4))", 15},
	{"1+2*(3+4)-(5)", 10},
	{"1+2*((3+4)*5+6)", 83},
	{"1+(1+(3+4)*5+6)*2", 85},
}

func TestParser(t *testing.T) {

	for _, tst := range parseTests {
		t.Run(tst.input, func(t *testing.T) {
			s := scan.New(tst.input)
			p := New(s)
			if res := p.exec(); res != tst.result {
				t.Errorf("%s=%f instead of %f", tst.input, res, tst.result)
			}
		})
	}
}
