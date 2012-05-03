package h

type Script struct {
	Src   string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Script) HTML() (HTML, error) {
	return &ReflectNode{Tag: "script", Node: s}, nil
}
