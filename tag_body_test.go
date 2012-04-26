package h

import (
	"testing"
)

func TestEmptyBody(t *testing.T) {
	t.Parallel()
	body := &Body{}
	assertRender(t, body, `<body></body>`)
}
