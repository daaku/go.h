package h

type Table struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Table) HTML() (HTML, error) {
	return &ReflectNode{Tag: "table", Node: t}, nil
}

type Tr struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Tr) HTML() (HTML, error) {
	return &ReflectNode{Tag: "tr", Node: t}, nil
}

type Th struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Th) HTML() (HTML, error) {
	return &ReflectNode{Tag: "th", Node: t}, nil
}

type Td struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Td) HTML() (HTML, error) {
	return &ReflectNode{Tag: "td", Node: t}, nil
}

type Thead struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Thead) HTML() (HTML, error) {
	return &ReflectNode{Tag: "thead", Node: t}, nil
}

type Tbody struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Tbody) HTML() (HTML, error) {
	return &ReflectNode{Tag: "tbody", Node: t}, nil
}
