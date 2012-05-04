package h

type Strong struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Strong) HTML() (HTML, error) {
	return &ReflectNode{Tag: "strong", Node: s}, nil
}
