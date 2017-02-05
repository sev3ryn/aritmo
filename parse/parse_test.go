package parse

import (
	"testing"

	"github.com/sev3ryn/aritmo/scan"
)

func TestParser(t *testing.T) {

	s := scan.New("1+2+3+4")
	p := New(s)
	if res := p.exec(); res != 10 {
		t.Errorf("1+2+3+4=%d instead of 10", res)
	}

	s = scan.New("1+2*3+4")
	p = New(s)
	if res := p.exec(); res != 11 {
		t.Errorf("1+2*3+4=%d instead of 11", res)
	}

}
