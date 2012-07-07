package h_test

import (
	"github.com/daaku/go.h"
	"testing"
)

func TestEmptyTitle(t *testing.T) {
	t.Parallel()
	title := &h.Title{h.String("")}
	assertRender(t, title, `<title></title>`)
}
