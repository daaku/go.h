package h

import (
	"errors"
	"net/url"

	"golang.org/x/net/context"
)

func HiddenInputs(values url.Values) HTML {
	frag := &Frag{}
	for key, list := range values {
		for _, each := range list {
			frag.Append(&Input{Name: key, Value: each})
		}
	}
	return &Div{Style: "display:none", Inner: frag}
}

var _ HTML = (*LinkStyle)(nil)

type LinkStyle struct {
	HREF string
}

func (l *LinkStyle) HTML(ctx context.Context) (HTML, error) {
	if l.HREF == "" {
		return nil, errors.New("Missing HREF in LinkStyle.")
	}
	return &Link{
		Type: "text/css",
		Rel:  "stylesheet",
		HREF: l.HREF,
	}, nil
}
