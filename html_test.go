package h_test

import (
	"testing"

	"github.com/daaku/go.h"
)

func TestEmptyDocument(t *testing.T) {
	t.Parallel()
	doc := &h.Document{}
	assertRender(t, doc, `<!doctype html><html></html>`)
}

func TestFacebookXMLNS(t *testing.T) {
	t.Parallel()
	doc := &h.Document{
		XMLNS: h.XMLNS{
			"fb": "http://ogp.me/ns/fb#",
		},
	}
	assertRender(t, doc,
		`<!doctype html><html xmlns:fb="http://ogp.me/ns/fb#"></html>`)
}

func TestDocumentLangAttr(t *testing.T) {
	t.Parallel()
	doc := &h.Document{
		Lang: "en",
	}
	assertRender(t, doc, `<!doctype html><html lang="en"></html>`)
}

func TestEmptyBody(t *testing.T) {
	t.Parallel()
	body := &h.Body{}
	assertRender(t, body, `<body></body>`)
}

func TestEmptyHead(t *testing.T) {
	t.Parallel()
	head := &h.Head{}
	assertRender(t, head, `<head></head>`)
}

func TestEmptyH1(t *testing.T) {
	t.Parallel()
	h1 := &h.H1{}
	assertRender(t, h1, `<h1></h1>`)
}

func TestEmptyLink(t *testing.T) {
	t.Parallel()
	link := &h.Link{}
	assertRender(t, link, `<link>`)
}

func TestEmptyLi(t *testing.T) {
	t.Parallel()
	li := &h.Li{}
	assertRender(t, li, `<li></li>`)
}

func TestEmptyMeta(t *testing.T) {
	t.Parallel()
	meta := &h.Meta{}
	assertRender(t, meta, `<meta>`)
}

func TestEmptySpan(t *testing.T) {
	t.Parallel()
	span := &h.Span{}
	assertRender(t, span, `<span></span>`)
}

func TestEmptyTitle(t *testing.T) {
	t.Parallel()
	title := &h.Title{h.String("")}
	assertRender(t, title, `<title></title>`)
}
