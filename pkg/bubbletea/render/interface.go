package render

type Render interface {
	Render(w, h int, c ...string) (string, bool)
}
