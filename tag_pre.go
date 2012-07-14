package h

type Pre struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (d *Pre) HTML() (HTML, error) {
	return &ReflectNode{Tag: "pre", Node: d}, nil
}
