package h

import (
	"testing"
)

func TestEmptyUl(t *testing.T) {
	t.Parallel()
	ul := &Ul{}
	assertRender(t, ul, `<ul></ul>`)
}
