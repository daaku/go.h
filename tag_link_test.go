package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyLink(t *testing.T) {
	t.Parallel()
	link := &h.Link{}
	assertRender(t, link, `<link>`)
}
