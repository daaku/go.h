package h

type Link struct {
	HREF string `h:"attr"`
	Type string `h:"attr"`
	Rel  string `h:"attr"`
}

func (l *Link) HTML() (HTML, error) {
	return &ReflectNode{"link", l, true}, nil
}
