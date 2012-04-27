package h

type Body struct {
	ID    string
	Class string
	Style string
	Title string
	Inner HTML
}

func (b *Body) HTML() (HTML, error) {
	return &Node{
		Tag: "body",
	}, nil
}
