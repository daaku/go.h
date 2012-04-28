package h

type H1 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H1) HTML() (HTML, error) {
	return &ReflectNode{Tag: "h1", Node: h}, nil
}
