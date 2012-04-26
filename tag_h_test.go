package h

import (
	"testing"
)

func TestEmptyH1(t *testing.T) {
	t.Parallel()
	h1 := &H1{}
	assertRender(t, h1, `<h1></h1>`)
}
