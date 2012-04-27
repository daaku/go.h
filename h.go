package h

import (
	"bytes"
	"fmt"
	"io"
)

type HTML interface {
	HTML() (HTML, error)
}

type Primitive interface {
	Write(io.Writer) (int, error)
}

// Write HTML into a writer.
func Write(w io.Writer, h HTML) (int, error) {
	var err error
	for {
		switch t := h.(type) {
		case nil:
			return 0, nil
		case Primitive:
			return t.Write(w)
		case HTML:
			h, err = h.HTML()
			if err != nil {
				return 0, err
			}
		default:
			return 0, fmt.Errorf("Value %+v of unknown type %T", h, h)
		}
	}
	panic("not reached")
}

// Render HTML as a string.
func Render(h HTML) (string, error) {
	buffer := bytes.NewBufferString("")
	_, err := Write(buffer, h)
	return buffer.String(), err
}
