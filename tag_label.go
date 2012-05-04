package h

type Label struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	For   string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (l *Label) HTML() (HTML, error) {
	return &ReflectNode{Tag: "label", Node: l}, nil
}
