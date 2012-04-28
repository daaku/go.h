package h

type Div struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (d *Div) HTML() (HTML, error) {
	return &ReflectNode{Tag: "div", Node: d}, nil
}
