package h_test

import (
	"github.com/nshah/go.h"
	"testing"
)

func TestEmptyBody(t *testing.T) {
	t.Parallel()
	body := &h.Body{}
	assertRender(t, body, `<body></body>`)
}
