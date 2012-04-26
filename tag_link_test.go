package h

import (
	"testing"
)

func TestEmptyLink(t *testing.T) {
	t.Parallel()
	link := &Link{}
	assertRender(t, link, `<link>`)
}
