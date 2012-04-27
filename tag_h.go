package h

type H1 struct {
	ID    string
	Class string
	Style string
	Title string
	Inner HTML
}

func (h *H1) HTML() (HTML, error) {
	return &Node{
		Tag:   "h1",
		Inner: h.Inner,
	}, nil
}
