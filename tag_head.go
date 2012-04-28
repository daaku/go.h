package h

type Head struct {
	Inner HTML `h:"inner"`
}

func (h *Head) HTML() (HTML, error) {
	return &ReflectNode{Tag: "head", Node: h}, nil
}
