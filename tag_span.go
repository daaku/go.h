package h

type Span struct {
	ID    string
	Class string
	Style string
	Inner HTML
}

func (s *Span) HTML() (HTML, error) {
	attrs := Attributes{}
	if s.ID != "" {
		attrs["id"] = s.ID
	}
	if s.Class != "" {
		attrs["class"] = s.Class
	}
	if s.Style != "" {
		attrs["style"] = s.Style
	}
	return &Node{
		Tag:   "span",
		Attributes: attrs,
		Inner: s.Inner,
	}, nil
}
