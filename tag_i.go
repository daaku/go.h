package h

type I struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Inner HTML                   `h:"inner"`
	Data  map[string]interface{} `h:"dict"`
}

func (i *I) HTML() (HTML, error) {
	return &ReflectNode{Tag: "i", Node: i}, nil
}
