package h

type Meta struct {
	Charset  string
	Name     string
	Property string
	Content  string
}

func (m *Meta) HTML() HTML {
	attrs := Attributes{}
	if m.Charset != "" {
		attrs["charset"] = m.Charset
	}
	if m.Name != "" {
		attrs["name"] = m.Name
	}
	if m.Property != "" {
		attrs["property"] = m.Property
	}
	if m.Content != "" {
		attrs["content"] = m.Content
	}
	return &Node{
		Tag:         "meta",
		Attributes:  attrs,
		SelfClosing: true,
	}
}
