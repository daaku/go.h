package h_test

import (
	"github.com/daaku/go.h"
	"testing"
)

func TestEmptyUl(t *testing.T) {
	t.Parallel()
	ul := &h.Ul{}
	assertRender(t, ul, `<ul></ul>`)
}
