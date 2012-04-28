package h

type Ul struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (ul *Ul) HTML() (HTML, error) {
	return &ReflectNode{Tag: "ul", Node: ul}, nil
}
