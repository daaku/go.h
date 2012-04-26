package h

import (
	"testing"
)

func TestEmptyTitle(t *testing.T) {
	t.Parallel()
	title := Title("")
	assertRender(t, title, `<title></title>`)
}
