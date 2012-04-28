package h

type Meta struct {
	Charset  string `h:"attr"`
	Name     string `h:"attr"`
	Property string `h:"attr"`
	Content  string `h:"attr"`
}

func (m *Meta) HTML() (HTML, error) {
	return &ReflectNode{"meta", m, true}, nil
}
