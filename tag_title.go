package h

type Title string

func (t Title) HTML() HTML {
	return &Node{
		Tag:   "title",
		Inner: String(t),
	}
}
