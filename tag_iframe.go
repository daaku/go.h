package h

type Iframe struct {
	ID                string `h:"attr"`
	Class             string `h:"attr"`
	Style             string `h:"attr"`
	Src               string `h:"attr"`
	Scrolling         bool   `h:"attr"`
	FrameBorder       int    `h:"attr"`
	AllowTransparency bool   `h:"attr"`
}

func (t *Iframe) HTML() (HTML, error) {
	return &ReflectNode{Tag: "iframe", Node: t}, nil
}
