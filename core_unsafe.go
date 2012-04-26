package h

import (
	"fmt"
	"io"
	"log"
)

type Unsafe string

func (u Unsafe) HTML() HTML {
	log.Fatalf("Unsafe.HTML called for %s", u)
	return u
}

func (u Unsafe) Write(w io.Writer) (int, error) {
	return fmt.Fprint(w, u)
}
