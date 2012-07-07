package h_test

import (
	"github.com/daaku/go.h"
	"testing"
)

func TestEmptyA(t *testing.T) {
	t.Parallel()
	a := &h.A{}
	assertRender(t, a, `<a></a>`)
}

func TestWithHREF(t *testing.T) {
	t.Parallel()
	a := &h.A{HREF: "/"}
	assertRender(t, a, `<a href="/"></a>`)
}
