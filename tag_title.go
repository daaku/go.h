package h

type Title Frag

func (t *Title) HTML() (HTML, error) {
	f := Frag(*t)
	return &Node{
		Tag:   "title",
		Inner: &f,
	}, nil
}
