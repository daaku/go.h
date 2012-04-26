package h

import (
	"html"
)

type String string

func (s String) HTML() HTML {
	return Unsafe(html.EscapeString(string(s)))
}
