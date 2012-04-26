package h

type LinkedStyle struct {
	HREF string
}

func (l *LinkedStyle) HTML() HTML {
	return &Link{
		Type: "text/css",
		Rel:  "stylesheet",
		HREF: l.HREF,
	}
}
