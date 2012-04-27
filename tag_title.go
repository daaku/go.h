package h

type Title string

func (t Title) HTML() (HTML, error) {
	return &Node{
		Tag:   "title",
		Inner: String(t),
	}, nil
}
