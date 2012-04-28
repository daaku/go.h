package h

import (
	"testing"
)

func TestEmptySpan(t *testing.T) {
	t.Parallel()
	span := &Span{}
	assertRender(t, span, `<span></span>`)
}
