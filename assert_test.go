package h

import (
	"testing"
)

func assertRender(t *testing.T, h HTML, expected string) {
	actual, err := Render(h)
	if err != nil {
		t.Fatalf("Failed to render document %v with error %s", h, err)
	}
	if actual != expected {
		t.Fatalf("Did not find expected:\n%s\ninstead found:\n%s", expected, actual)
	}
}
