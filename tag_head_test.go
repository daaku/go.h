package h

import (
	"testing"
)

func TestEmptyHead(t *testing.T) {
	t.Parallel()
	head := &Head{}
	assertRender(t, head, `<head></head>`)
}
