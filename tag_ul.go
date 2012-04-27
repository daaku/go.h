package h

type Ul struct {
	ID    string
	Class string
	Style string
	Title string
	Inner HTML
}

func (ul *Ul) HTML() (HTML, error) {
	return &Node{
		Tag:   "ul",
		Inner: ul.Inner,
	}, nil
}