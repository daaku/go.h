package h

type Input struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Name string `h:"attr"`
	Style string `h:"attr"`
	Type  string `h:"attr"`
	Value string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (i *Input) HTML() (HTML, error) {
	return &ReflectNode{
		Tag:         "input",
		Node:        i,
		SelfClosing: true,
	}, nil
}
