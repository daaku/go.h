package h_test

import (
	"testing"

	"github.com/daaku/go.h"
	"golang.org/x/net/context"
)

func assertRender(t *testing.T, html h.HTML, expected string) {
	actual, err := h.Render(context.Background(), html)
	if err != nil {
		t.Fatalf("Failed to render document %v with error %s", html, err)
	}
	if actual != expected {
		t.Fatalf("Did not find expected:\n%s\ninstead found:\n%s", expected, actual)
	}
}
