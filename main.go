// +build js

package main

import (
	"maunium.net/go/gopher-ace"
	"github.com/gopherjs/gopherjs/js"
	"github.com/sev3ryn/aritmo/scan"
	"github.com/sev3ryn/aritmo/parse"
	"fmt"
)

func getLine(line int) ace.Range {

	return ace.NewRange(line, 0, line, int(^uint(0) >> 1))
}

func getLineRange(line int) ace.Range {
	o := js.Global.Get("Object").New()
	o.Set("start", map[string]interface{}{
		"row": line,
		"column": 0,
	})
	o.Set("end",map[string]interface{}{
		"row": line,
		"column": int(^uint(0) >> 1),
	})
	return ace.Range{Object:o}
}

func expandResultPane(r ace.Editor, rses ace.EditSession, numLines int){
	for i:=0;i<numLines;i++{
		rses.Insert(0,0,"\n")
	}
	r.GotoLine(0,0,false)
}

func syncScroll(s1,s2 ace.EditSession){
	s1.On("changeScrollTop", func() {
		s2.SetScrollTop(s1.GetScrollTop())
	});
}

func main(){
	e := ace.Edit("editor")
	e.SetFontSize(22)
	e.SetTheme("ace/theme/iplastic")
	sess := e.GetSession()
	sess.SetMode("ace/mode/golang")

	r := ace.Edit("result")
	r.SetFontSize(22)
	r.SetOption("showGutter", false)
	r.SetOption("highlightActiveLine", false)
	r.SetOption("cursorStyle", "none")
	r.SetReadOnly(true)
	rses := r.GetSession()

	// hack while no enter handler found
	expandResultPane(r, rses, 100)

	syncScroll(sess, rses)
	syncScroll(rses,sess)
	




	//fmt.Printf("%#v",e.GetKeyBinding())

	e.OnChange(func (j *js.Object){
		line := e.GetSelectionRange().StartRow()
		fmt.Println(line)
		inp := sess.GetLine(line)
		s := scan.New(inp)
		p := parse.New(s)
		if res, err := p.ExecStatement(); err == nil {
			rses.Replace(getLineRange(line), fmt.Sprintf("%f",res))
		} else {
			rses.Replace(getLineRange(line), "")
		}
	})

}


