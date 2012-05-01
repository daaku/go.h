package h

type Img struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Src string `h:"attr"`
	Inner HTML `h:"inner"`
}

func (t *Img) HTML() (HTML, error) {
	return &ReflectNode{Tag: "img", Node: t}, nil
}
