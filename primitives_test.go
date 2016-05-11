package h

import (
	"context"
	"testing"

	"github.com/facebookgo/ensure"
)

func TestPrimitivesHTML(t *testing.T) {
	cases := []struct {
		Name string
		HTML HTML
	}{
		{Name: "Frag", HTML: &Frag{}},
		{Name: "Unsafe", HTML: Unsafe("")},
		{Name: "UnsafeBytes", HTML: UnsafeBytes(nil)},
		{Name: "Node", HTML: &Node{}},
		{Name: "ReflectNode", HTML: &ReflectNode{}},
	}
	for _, c := range cases {
		v, err := c.HTML.HTML(context.Background())
		ensure.DeepEqual(t, err, errHTMLOnPrimitive(c.Name))
		ensure.Nil(t, v)
	}
}

func TestUnsafeBytes(t *testing.T) {
	raw := []byte("<foo></bar>")
	s, err := Render(context.Background(), UnsafeBytes(raw))
	ensure.Nil(t, err)
	ensure.DeepEqual(t, s, string(raw))
}
