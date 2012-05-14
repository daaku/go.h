package h

type Option struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Value string `h:"attr"`
	Selected bool `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (o *Option) HTML() (HTML, error) {
	return &ReflectNode{Tag: "option", Node: o}, nil
}
