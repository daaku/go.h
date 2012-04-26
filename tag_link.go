package h

type Link struct {
	Type string
	Rel  string
	HREF string
}

func (l *Link) HTML() HTML {
	return &Node{
		Tag:         "link",
		SelfClosing: true,
	}
}
