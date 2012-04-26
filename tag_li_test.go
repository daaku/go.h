package h

import (
	"testing"
)

func TestEmptyLi(t *testing.T) {
	t.Parallel()
	li := &Li{}
	assertRender(t, li, `<li></li>`)
}
