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
	{"1+2*(3+4)-5", 10},
	{"1+2*(3+4)-5*6", -15},
	{"1+2*(3+4)-5*(6+7)", -50},
	//{"1+2*((3+4))", 15},
	//{"1+2*((3+4)*5+6)", 83},
}

func TestParser(t *testing.T) {

	for _, tst := range parseTests {
		t.Run(tst.input, func(t *testing.T) {
			s := scan.New(tst.input)
			p := New(s)
			if res := p.exec(); res != tst.result {
				t.Errorf("1+2+3+4=%f instead of %f", res, tst.result)
			}
		})
	}
}
