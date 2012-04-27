package h

type A struct {
	ID    string
	Class string
	Style string
	Title string
	HREF  string
	Inner HTML
}

func (a *A) HTML() (HTML, error) {
	attrs := Attributes{}
	if a.ID != "" {
		attrs["id"] = a.ID
	}
	if a.Class != "" {
		attrs["class"] = a.Class
	}
	if a.Style != "" {
		attrs["style"] = a.Style
	}
	return &Node{
		Tag:   "a",
		Attributes: attrs,
		Inner: a.Inner,
	}, nil
}
