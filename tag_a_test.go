package h

import (
	"testing"
)

func TestEmptyA(t *testing.T) {
	t.Parallel()
	a := &A{}
	assertRender(t, a, `<a></a>`)
}
