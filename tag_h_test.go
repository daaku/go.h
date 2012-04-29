package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyH1(t *testing.T) {
	t.Parallel()
	h1 := &h.H1{}
	assertRender(t, h1, `<h1></h1>`)
}
