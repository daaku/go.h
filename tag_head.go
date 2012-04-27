package h

type Head struct {
	Inner HTML
}

func (h *Head) HTML() (HTML, error) {
	return &Node{
		Tag:   "head",
		Inner: h.Inner,
	}, nil
}
