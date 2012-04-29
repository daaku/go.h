package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyHead(t *testing.T) {
	t.Parallel()
	head := &h.Head{}
	assertRender(t, head, `<head></head>`)
}
