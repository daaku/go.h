package h

type Div struct {
	ID    string
	Class string
	Style string
	Inner HTML
}

func (d *Div) HTML() (HTML, error) {
	attrs := Attributes{}
	if d.ID != "" {
		attrs["id"] = d.ID
	}
	if d.Class != "" {
		attrs["class"] = d.Class
	}
	if d.Style != "" {
		attrs["style"] = d.Style
	}
	return &Node{
		Tag:   "div",
		Attributes: attrs,
		Inner: d.Inner,
	}, nil
}
