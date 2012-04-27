package h

type LinkStyle struct {
	HREF string
}

func (l *LinkStyle) HTML() (HTML, error) {
	return &Link{
		Type: "text/css",
		Rel:  "stylesheet",
		HREF: l.HREF,
	}, nil
}
