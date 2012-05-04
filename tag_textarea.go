package h

type Textarea struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Name  string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Textarea) HTML() (HTML, error) {
	return &ReflectNode{Tag: "textarea", Node: t}, nil
}
