package render

import (
	"strings"
)

type centerRender struct{}

var CenterRender centerRender

func (r centerRender) Render(w, h int, c ...string) (string, bool) {
	cheight, cwidth := 0, 0
	blocks := make([][]string, 0)
	bwidth := make([]int, 0)
	for _, block := range c {
		lines := strings.Split(block, "\n")
		cheight += len(lines)
		bw := 0
		for _, line := range lines {
			if len(line) > bw {
				bw = len(line)
			}
		}
		blocks = append(blocks, lines)
		bwidth = append(bwidth, bw)
		if bw > cwidth {
			cwidth = bw
		}
	}
	if cheight > h || cwidth > w {
		return strings.Join(c, "\n"), false
	}
	topSpace := (h - cheight) / 2
	bottomSpace := h - cheight - topSpace

	rendered := []string{}
	for i := 0; i < topSpace; i++ {
		rendered = append(rendered, strings.Repeat(" ", w))
	}
	for i, block := range blocks {
		leftSpace := (w - bwidth[i]) / 2
		rightSpace := w - bwidth[i] - leftSpace
		for _, line := range block {
			rendered = append(rendered, strings.Repeat(" ", leftSpace)+line+strings.Repeat(" ", rightSpace))
		}
	}
	for i := 0; i < bottomSpace; i++ {
		rendered = append(rendered, strings.Repeat(" ", w))
	}
	return strings.Join(rendered, "\n"), true
}
