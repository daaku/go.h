package h

type Script struct {
	Inner HTML
}

func (s *Script) HTML() (HTML, error) {
	return &Node{
		Tag:   "script",
		Inner: s.Inner,
	}, nil
}
