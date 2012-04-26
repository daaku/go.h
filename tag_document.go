package h

type XMLNS map[string]string

type Document struct {
	XMLNS XMLNS
	Head  HTML
	Body  HTML
}

func (ns XMLNS) Attributes() Attributes {
	attrs := Attributes{}
	for key, val := range ns {
		attrs["xmlns:"+key] = val
	}
	return attrs
}

func (d *Document) HTML() HTML {
	return &Node{
		Tag:        "html",
		Attributes: d.XMLNS.Attributes(),
		Inner:      &Frag{d.Head, d.Body},
	}
}
