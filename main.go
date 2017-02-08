// +build js

package main

import (
	"maunium.net/go/gopher-ace"
	"github.com/gopherjs/gopherjs/js"
	"github.com/sev3ryn/aritmo/scan"
	"github.com/sev3ryn/aritmo/parse"
	"fmt"
)

func main(){
	e := ace.Edit("editor")
	e.SetFontSize(22)
	
	e.SetTheme("ace/theme/iplastic")
	sess := e.GetSession()
	sess.SetMode("ace/mode/golang")
	
	e.OnChange(func (j *js.Object){
		inp := sess.GetLine(e.GetSelectionRange().StartRow())
		s := scan.New(inp)
		p := parse.New(s)
		if res, err := p.ExecStatement(); err == nil {
			fmt.Println(res)
		}
		
	})

}


