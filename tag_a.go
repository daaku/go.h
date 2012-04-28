package h

type A struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	HREF  string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (a *A) HTML() (HTML, error) {
	return &ReflectNode{Tag: "a", Node: a}, nil
}
