package h

type Body struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (b *Body) HTML() (HTML, error) {
	return &ReflectNode{Tag: "body", Node: b}, nil
}
