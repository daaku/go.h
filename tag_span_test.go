package h_test

import (
	"github.com/daaku/go.h"
	"testing"
)

func TestEmptySpan(t *testing.T) {
	t.Parallel()
	span := &h.Span{}
	assertRender(t, span, `<span></span>`)
}
