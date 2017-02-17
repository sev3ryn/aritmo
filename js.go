// +build js

package main

import (
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/sev3ryn/aritmo/parse"
	"github.com/sev3ryn/aritmo/scan"
	"maunium.net/go/gopher-ace"
)

func getLine(line int) ace.Range {

	return ace.NewRange(line, 0, line, int(^uint(0)>>1))
}

func getLineRange(line int) ace.Range {
	o := js.Global.Get("Object").New()
	o.Set("start", map[string]interface{}{
		"row":    line,
		"column": 0,
	})
	o.Set("end", map[string]interface{}{
		"row":    line,
		"column": int(^uint(0) >> 1),
	})
	return ace.Range{Object: o}
}

func expandResultPane(r *ace.Editor, rses *ace.EditSession, numLines int) {
	for i := 0; i < numLines; i++ {
		rses.Insert(0, 0, "\n")
	}
	r.GotoLine(0, 0, false)
}

func syncScroll(s1, s2 *ace.EditSession) {
	s1.On("changeScrollTop", func() {
		s2.SetScrollTop(s1.GetScrollTop())
	})
}

func setupEditor(divId string) (*ace.Editor, *ace.EditSession) {
	e := ace.Edit(divId)
	e.SetFontSize(22)
	e.SetTheme("ace/theme/iplastic")
	s := e.GetSession()
	s.SetMode("ace/mode/golang")
	return &e, &s
}

func setupResultPane(divId string) (*ace.Editor, *ace.EditSession) {
	e := ace.Edit("result")
	e.SetFontSize(22)
	e.SetOptions(map[string]interface{}{
		"showGutter":          false,
		"highlightActiveLine": false,
		"readOnly":            true,
	})
	s := e.GetSession()
	e.Set("$blockScrolling", 1)
	// hack while keyboard handler not found yet
	expandResultPane(&e, &s, 99)
	return &e, &s
}

func refresh(
	eSession, rSession *ace.EditSession,
	startLine, endLine int,
) {
	if startLine > endLine {
		return
	}

	var input string
	var output float64
	var err error
	var p *parse.Parser
	for line := startLine; line < endLine+1; line++ {
		input = eSession.GetLine(line)
		p = parse.New(scan.New(input))
		if output, err = p.ExecStatement(); err == nil {
			rSession.Replace(getLineRange(line), fmt.Sprintf("%f", output))
		} else {
			rSession.Replace(getLineRange(line), "")
		}
	}
}

func init() {

	e, eSession := setupEditor("editor")
	_, rSession := setupResultPane("result")
	//set start focus
	e.Focus()
	// sync scrolling in both windows
	syncScroll(eSession, rSession)
	syncScroll(rSession, eSession)

	e.Get("commands").Get("byName").Get("backspace").Set("exec", func() { fmt.Println("backspace") })
	//e.Get("commands").Get("byName").Get("enter").Set("exec", func(){fmt.Println("enter")})

	var line int
	e.OnChange(func(j *js.Object) {
		line = e.GetSelectionRange().StartRow()
		refresh(eSession, rSession, line, line)
	})

}
