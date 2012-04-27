package h

type Li struct {
	ID    string
	Class string
	Style string
	Title string
	Inner HTML
}

func (l *Li) HTML() (HTML, error) {
	return &Node{
		Tag:   "li",
		Inner: l.Inner,
	}, nil
}
