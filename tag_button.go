package h

type Button struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (b *Button) HTML() (HTML, error) {
	return &ReflectNode{Tag: "button", Node: b}, nil
}
