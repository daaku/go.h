package h

type Span struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Span) HTML() (HTML, error) {
	return &ReflectNode{Tag: "span", Node: s}, nil
}
