package render

import "strings"

type frameCenterRender struct{}

var FrameCenterRender frameCenterRender

func (r frameCenterRender) Render(w, h int, c ...string) (string, bool) {
	s, ok := CenterRender.Render(w-2, h-2, c...)
	if !ok {
		return s, ok
	}
	rendered := []string{}
	rendered = append(rendered, "+"+strings.Repeat("-", w-2)+"+")
	for _, s := range strings.Split(s, "\n") {
		rendered = append(rendered, "|"+s+"|")
	}
	rendered = append(rendered, "+"+strings.Repeat("-", w-2)+"+")
	return strings.Join(rendered, "\n"), true
}
