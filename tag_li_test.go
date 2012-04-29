package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyLi(t *testing.T) {
	t.Parallel()
	li := &h.Li{}
	assertRender(t, li, `<li></li>`)
}
