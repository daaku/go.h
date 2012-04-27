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
	return &Node{
		Tag:   "a",
		Inner: a.Inner,
	}, nil
}
