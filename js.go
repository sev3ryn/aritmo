// +build js

package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/sev3ryn/aritmo/storage"

	"maunium.net/go/gopher-ace"
)

// for incremental refres - not used for now as has bugs
func getLine(line int) ace.Range {

	return ace.NewRange(line, 0, line, int(^uint(0)>>1))
}

func getLineRange(start, end int) ace.Range {
	o := js.Global.Get("Object").New()
	o.Set("start", map[string]interface{}{
		"row":    start,
		"column": 0,
	})
	o.Set("end", map[string]interface{}{
		"row":    end,
		"column": int(^uint(0) >> 1),
	})
	return ace.Range{Object: o}
}

// for incremental refres - not used for now as has bugs
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
	return &e, &s
}

// for incremental refres - not used for now as has bugs
func refresh(
	eSession, rSession *ace.EditSession,
	selection *ace.Range,
) {

	startLine, endLine := selection.StartRow(), selection.EndRow()

	if startLine > endLine {
		return
	}

	editLength := eSession.GetLength()
	resutLength := rSession.GetLength()

	var input string
	if editLength == resutLength {
		for line := startLine; line <= endLine; line++ {
			input = eSession.GetLine(line)
			rSession.Replace(getLineRange(line, line), calculate(line, input))
		}
	} else if editLength > resutLength {

		for i := 0; i <= editLength-resutLength; i++ {
			if i != editLength-resutLength {
				rSession.Insert(startLine+i, 20, "\n")
			}

			input = eSession.GetLine(startLine + i)
			rSession.Replace(getLineRange(startLine+i, startLine+i), calculate(startLine+i, input))
		}

	} else if editLength < resutLength {
		input = eSession.GetLine(startLine)
		rSession.Remove(getLineRange(startLine, endLine))
		rSession.Replace(getLineRange(startLine, startLine), calculate(startLine, input))
	}

}

func recompute(eSession, rSession *ace.EditSession) {
	rSession.Remove(getLineRange(0, rSession.GetLength()-1))

	store = storage.RAMStore
	for i := 0; i < eSession.GetLength(); i++ {
		input := eSession.GetLine(i)
		rSession.Insert(i, 20, calculate(i, input)+"\n")
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

	//e.Get("commands").Get("byName").Get("enter").Set("exec", func(){fmt.Println("enter")})

	//var selection ace.Range
	e.OnChange(func(j *js.Object) {

		// temporary - recalculate all window. TODO - do diff updates as below
		recompute(eSession, rSession)

		//selection = e.GetSelectionRange()
		//refresh(eSession, rSession, &selection)
	})

}
