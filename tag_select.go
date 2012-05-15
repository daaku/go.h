package h

type Select struct {
	ID       string `h:"attr"`
	Class    string `h:"attr"`
	Name     string `h:"attr"`
	Style    string `h:"attr"`
	Title    string `h:"attr"`
	Multiple bool   `h:"attr"`
	Size     int    `h:"size"`
	Inner    HTML   `h:"inner"`
}

func (s *Select) HTML() (HTML, error) {
	return &ReflectNode{Tag: "select", Node: s}, nil
}
