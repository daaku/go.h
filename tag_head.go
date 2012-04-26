package h

type Head struct {
	Inner HTML
}

func (h *Head) HTML() HTML {
	return &Node{
		Tag:   "head",
		Inner: h.Inner,
	}
}
