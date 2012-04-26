package h

import (
	"io"
	"log"
)

type Frag []HTML

func (f *Frag) HTML() HTML {
	log.Fatalf("Frag.HTML called for %s", f)
	return f
}

func (f *Frag) Append(h HTML) {
	*f = append(*f, h)
}

func (f *Frag) Write(w io.Writer) (int, error) {
	written := 0
	for _, e := range *f {
		i, err := Write(w, e)
		written += i
		if err != nil {
			return written, err
		}
	}
	return written, nil
}
