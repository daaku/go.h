package h

type Body struct {
	ID    string
	Class string
	Style string
	Inner HTML
}

func (b *Body) HTML() (HTML, error) {
	attrs := Attributes{}
	if b.ID != "" {
		attrs["id"] = b.ID
	}
	if b.Class != "" {
		attrs["class"] = b.Class
	}
	if b.Style != "" {
		attrs["style"] = b.Style
	}
	return &Node{
		Tag:        "body",
		Attributes: attrs,
		Inner:      b.Inner,
	}, nil
}
