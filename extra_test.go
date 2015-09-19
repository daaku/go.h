package h

import (
	"net/url"
	"testing"

	"golang.org/x/net/context"

	"github.com/facebookgo/ensure"
)

func TestHiddenInputs(t *testing.T) {
	const (
		k1  = "k1"
		v1a = "v1a"
		v1b = "v1b"
		k2  = "k2"
		v2  = "v2"
	)
	v := url.Values{
		k1: []string{v1a, v1b},
		k2: []string{v2},
	}
	d := HiddenInputs(v).(*Div)
	ensure.DeepEqual(t, d.Style, "display:none")
	f := d.Inner.(Frag)
	ensure.SameElements(t, f, Frag{
		&Input{Name: k1, Value: v1a},
		&Input{Name: k1, Value: v1b},
		&Input{Name: k2, Value: v2},
	})
}

func TestLinkStyleMissingHREF(t *testing.T) {
	h, err := (&LinkStyle{}).HTML(context.Background())
	ensure.True(t, err == errLinkStyleMissingHREF)
	ensure.Nil(t, h)
}

func TestLinkStyleMissing(t *testing.T) {
	const href = "42"
	h, err := (&LinkStyle{HREF: href}).HTML(context.Background())
	ensure.Nil(t, err)
	ensure.DeepEqual(t, h, &Link{
		Type: "text/css",
		Rel:  "stylesheet",
		HREF: href,
	})
}
