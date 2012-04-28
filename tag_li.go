package h

type Li struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (l *Li) HTML() (HTML, error) {
	return &ReflectNode{Tag: "li", Node: l}, nil
}
