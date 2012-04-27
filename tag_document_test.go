package h

import (
	"testing"
)

func TestEmptyDocument(t *testing.T) {
	t.Parallel()
	doc := &Document{}
	assertRender(t, doc, `<!doctype html><html></html>`)
}

func TestFacebookXMLNS(t *testing.T) {
	t.Parallel()
	doc := &Document{
		XMLNS: XMLNS{
			"fb": "http://ogp.me/ns/fb#",
		},
	}
	assertRender(t, doc,
		`<!doctype html><html xmlns:fb="http://ogp.me/ns/fb#"></html>`)
}
