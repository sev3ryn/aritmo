package parse

import (
	"testing"

	"github.com/sev3ryn/aritmo/frontend/core/datatype"
	"github.com/sev3ryn/aritmo/frontend/core/scan"
	"github.com/sev3ryn/aritmo/frontend/core/storage"
)

type parseTest struct {
	input  string
	result storage.Result
}

var parseTests = []parseTest{
	{"1+2+3+4", storage.Result{Val: 10, Typ: datatype.BareDataType}},
	{"1+2*3+4", storage.Result{Val: 11, Typ: datatype.BareDataType}},
	{"1+2*(3+4)", storage.Result{Val: 15, Typ: datatype.BareDataType}},
	{"1+2*(3+4)-6", storage.Result{Val: 9, Typ: datatype.BareDataType}},
	{"1+2*(3+4)/2-1", storage.Result{Val: 7, Typ: datatype.BareDataType}},
	{"(1+2)*3", storage.Result{Val: 9, Typ: datatype.BareDataType}},
	{"1+2*(3+4)-5*6", storage.Result{Val: -15, Typ: datatype.BareDataType}},
	{"1+2*(3+4)-5*(6+7)", storage.Result{Val: -50, Typ: datatype.BareDataType}},
	{"1+2*((3+4))", storage.Result{Val: 15, Typ: datatype.BareDataType}},
	{"1+2*(3+4)-(5)", storage.Result{Val: 10, Typ: datatype.BareDataType}},
	{"1+2*((3+4)*5+6)", storage.Result{Val: 83, Typ: datatype.BareDataType}},
	{"1+(1+(3+4)*5+6)*2", storage.Result{Val: 85, Typ: datatype.BareDataType}},
	{"1+*2", storage.Result{Val: 0}},
	{"1/0", storage.Result{Val: 0}},
}

func TestParser(t *testing.T) {
	var currUpdCh = make(chan []byte)
	dt := datatype.Init(currUpdCh)
	for _, tst := range parseTests {
		t.Run(tst.input, func(t *testing.T) {
			s := scan.New(tst.input)
			p := New(s, storage.RAMStore, dt)
			if res, _ := p.execStatement(); res.Typ != tst.result.Typ || res.Val != tst.result.Val {
				t.Errorf("%s=%#v instead of %#v", tst.input, res, tst.result)
			}
		})
	}
}
