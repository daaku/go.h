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

type H2 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H2) HTML() (HTML, error) {
	return &ReflectNode{Tag: "h2", Node: h}, nil
}
