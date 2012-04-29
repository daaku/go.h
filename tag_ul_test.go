package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyUl(t *testing.T) {
	t.Parallel()
	ul := &h.Ul{}
	assertRender(t, ul, `<ul></ul>`)
}
