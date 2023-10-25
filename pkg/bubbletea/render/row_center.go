package render

import "strings"

type rowCenterRender struct{}

var RowCenterRender rowCenterRender

func (r rowCenterRender) Render(w, h int, c ...string) (string, bool) {
	n := len(c)
	perw := w / n
	texts := make([][]string, 0, n)
	for _, t := range c {
		t, ok := CenterRender.Render(perw, h, t)
		if !ok {
			return "", ok
		}
		texts = append(texts, strings.Split(t, "\n"))
	}
	rendered := []string{}
	for i := 0; i < h; i++ {
		line := ""
		for _, t := range texts {
			line += t[i]
		}
		rendered = append(rendered, line)
	}
	return CenterRender.Render(w, h, rendered...)
}
