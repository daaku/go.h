package h

type Link struct {
	Type string `h:"attr"`
	Rel  string `h:"attr"`
	HREF string `h:"attr"`
}

func (l *Link) HTML() (HTML, error) {
	return &ReflectNode{"link", l, true}, nil
}
