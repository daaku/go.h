package h

import (
	"testing"
)

func TestEmptyMeta(t *testing.T) {
	t.Parallel()
	meta := &Meta{}
	assertRender(t, meta, `<meta>`)
}
