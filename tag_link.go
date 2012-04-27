package h

type Link struct {
	Type string
	Rel  string
	HREF string
}

func (l *Link) HTML() (HTML, error) {
	attrs := Attributes{}
	if l.Type != "" {
		attrs["type"] = l.Type
	}
	if l.Rel != "" {
		attrs["rel"] = l.Rel
	}
	if l.HREF != "" {
		attrs["href"] = l.HREF
	}
	return &Node{
		Tag:         "link",
		Attributes:  attrs,
		SelfClosing: true,
	}, nil
}
