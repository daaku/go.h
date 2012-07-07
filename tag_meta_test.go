package h_test

import (
	"github.com/daaku/go.h"
	"testing"
)

func TestEmptyMeta(t *testing.T) {
	t.Parallel()
	meta := &h.Meta{}
	assertRender(t, meta, `<meta>`)
}
