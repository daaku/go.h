package h

import "golang.org/x/net/context"

type XMLNS map[string]string

func (ns XMLNS) Attributes() Attributes {
	attrs := Attributes{}
	for key, val := range ns {
		attrs["xmlns:"+key] = val
	}
	return attrs
}

var _ HTML = (*Document)(nil)

type Document struct {
	XMLNS XMLNS
	Inner HTML
	ID    string
	Lang  string
}

func (d *Document) HTML(ctx context.Context) (HTML, error) {
	attrs := d.XMLNS.Attributes()
	if d.ID != "" {
		attrs["id"] = d.ID
	}
	if d.Lang != "" {
		attrs["lang"] = "en"
	}

	return &Frag{
		Unsafe("<!doctype html>"),
		&Node{
			Tag:        "html",
			Attributes: attrs,
			Inner:      d.Inner,
		},
	}, nil
}

var _ HTML = (*A)(nil)

type A struct {
	ID     string                 `h:"attr"`
	Class  string                 `h:"attr"`
	Style  string                 `h:"attr"`
	Title  string                 `h:"attr"`
	HREF   string                 `h:"attr"`
	Target string                 `h:"attr"`
	Rel    string                 `h:"attr"`
	Inner  HTML                   `h:"inner"`
	Data   map[string]interface{} `h:"dict"`
}

func (a *A) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "a", Node: a}, nil
}

var _ HTML = (*Body)(nil)

type Body struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (b *Body) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "body", Node: b}, nil
}

var _ HTML = (*Button)(nil)

type Button struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Type  string                 `h:"attr"`
	Style string                 `h:"attr"`
	Title string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (b *Button) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "button", Node: b}, nil
}

var _ HTML = (*Div)(nil)

type Div struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Div) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "div", Node: d}, nil
}

var _ HTML = (*P)(nil)

type P struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (p *P) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "p", Node: p}, nil
}

const (
	Post = "post"
	Get  = "get"
)

var _ HTML = (*Form)(nil)

type Form struct {
	ID      string `h:"attr"`
	Class   string `h:"attr"`
	Style   string `h:"attr"`
	Action  string `h:"attr"`
	Method  string `h:"attr"`
	EncType string `h:"attr"`
	Target  string `h:"attr"`
	Title   string `h:"attr"`
	Inner   HTML   `h:"inner"`
}

func (f *Form) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "form", Node: f}, nil
}

var _ HTML = (*FieldSet)(nil)

type FieldSet struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (f *FieldSet) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "fieldset", Node: f}, nil
}

var _ HTML = (*Legend)(nil)

type Legend struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (l *Legend) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "legend", Node: l}, nil
}

var _ HTML = (*Head)(nil)

type Head struct {
	Inner HTML `h:"inner"`
}

func (h *Head) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "head", Node: h}, nil
}

var _ HTML = (*H1)(nil)

type H1 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H1) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h1", Node: h}, nil
}

var _ HTML = (*H2)(nil)

type H2 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H2) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h2", Node: h}, nil
}

var _ HTML = (*H3)(nil)

type H3 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H3) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h3", Node: h}, nil
}

var _ HTML = (*H4)(nil)

type H4 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H4) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h4", Node: h}, nil
}

var _ HTML = (*H5)(nil)

type H5 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H5) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h5", Node: h}, nil
}

var _ HTML = (*H6)(nil)

type H6 struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (h *H6) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "h6", Node: h}, nil
}

var _ HTML = (*Iframe)(nil)

type Iframe struct {
	ID              string `h:"attr"`
	Class           string `h:"attr"`
	Style           string `h:"attr"`
	Src             string `h:"attr"`
	Width           int    `h:"attr"`
	Height          int    `h:"attr"`
	FrameBorder     int    `h:"attr"`
	AllowFullScreen bool   `h:"attr"`
}

func (t *Iframe) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "iframe", Node: t}, nil
}

var _ HTML = (*I)(nil)

type I struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Inner HTML                   `h:"inner"`
	Data  map[string]interface{} `h:"dict"`
}

func (i *I) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "i", Node: i}, nil
}

var _ HTML = (*Img)(nil)

type Img struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Src   string                 `h:"attr"`
	Alt   string                 `h:"attr"`
	Inner HTML                   `h:"inner"`
	Data  map[string]interface{} `h:"dict"`
}

func (t *Img) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{
		Tag:         "img",
		Node:        t,
		SelfClosing: true,
	}, nil
}

var _ HTML = (*Input)(nil)

type Input struct {
	ID          string                 `h:"attr"`
	Class       string                 `h:"attr"`
	Name        string                 `h:"attr"`
	Style       string                 `h:"attr"`
	Type        string                 `h:"attr"`
	Value       string                 `h:"attr"`
	Placeholder string                 `h:"attr"`
	Checked     bool                   `h:"attr"`
	Multiple    bool                   `h:"attr"`
	Data        map[string]interface{} `h:"dict"`
	Inner       HTML                   `h:"inner"`
}

func (i *Input) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{
		Tag:         "input",
		Node:        i,
		SelfClosing: true,
	}, nil
}

var _ HTML = (*Label)(nil)

type Label struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	For   string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (l *Label) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "label", Node: l}, nil
}

var _ HTML = (*Li)(nil)

type Li struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (l *Li) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "li", Node: l}, nil
}

var _ HTML = (*Link)(nil)

type Link struct {
	HREF string `h:"attr"`
	Type string `h:"attr"`
	Rel  string `h:"attr"`
}

func (l *Link) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{"link", l, true}, nil
}

var _ HTML = (*Meta)(nil)

type Meta struct {
	Charset  string `h:"attr"`
	Name     string `h:"attr"`
	Property string `h:"attr"`
	Content  string `h:"attr"`
}

func (m *Meta) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{"meta", m, true}, nil
}

var _ HTML = (*Option)(nil)

type Option struct {
	ID       string                 `h:"attr"`
	Class    string                 `h:"attr"`
	Style    string                 `h:"attr"`
	Title    string                 `h:"attr"`
	Value    string                 `h:"attr"`
	Selected bool                   `h:"attr"`
	Data     map[string]interface{} `h:"dict"`
	Inner    HTML                   `h:"inner"`
}

func (o *Option) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "option", Node: o}, nil
}

var _ HTML = (*Pre)(nil)

type Pre struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (d *Pre) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "pre", Node: d}, nil
}

var _ HTML = (*Script)(nil)

type Script struct {
	Src   string `h:"attr"`
	Async bool   `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Script) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "script", Node: s}, nil
}

var _ HTML = (*Select)(nil)

type Select struct {
	ID       string `h:"attr"`
	Class    string `h:"attr"`
	Name     string `h:"attr"`
	Style    string `h:"attr"`
	Title    string `h:"attr"`
	Multiple bool   `h:"attr"`
	Size     int    `h:"size"`
	Inner    HTML   `h:"inner"`
}

func (s *Select) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "select", Node: s}, nil
}

var _ HTML = (*Span)(nil)

type Span struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Span) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "span", Node: s}, nil
}

var _ HTML = (*Strong)(nil)

type Strong struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Strong) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "strong", Node: s}, nil
}

var _ HTML = (*Table)(nil)

type Table struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Table) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "table", Node: t}, nil
}

var _ HTML = (*Tr)(nil)

type Tr struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (t *Tr) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "tr", Node: t}, nil
}

var _ HTML = (*Th)(nil)

type Th struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Th) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "th", Node: t}, nil
}

var _ HTML = (*Td)(nil)

type Td struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Td) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "td", Node: t}, nil
}

var _ HTML = (*Thead)(nil)

type Thead struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Thead) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "thead", Node: t}, nil
}

var _ HTML = (*Tbody)(nil)

type Tbody struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Tbody) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "tbody", Node: t}, nil
}

var _ HTML = (*Textarea)(nil)

type Textarea struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Name  string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (t *Textarea) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "textarea", Node: t}, nil
}

type Title Frag

func (t *Title) HTML(ctx context.Context) (HTML, error) {
	f := Frag(*t)
	return &Node{
		Tag:   "title",
		Inner: &f,
	}, nil
}

var _ HTML = (*Ul)(nil)

type Ul struct {
	ID    string `h:"attr"`
	Class string `h:"attr"`
	Style string `h:"attr"`
	Title string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (ul *Ul) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "ul", Node: ul}, nil
}

var _ HTML = (*Style)(nil)

type Style struct {
	ID    string `h:"attr"`
	Inner HTML   `h:"inner"`
}

func (s *Style) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "style", Node: s}, nil
}

var _ HTML = (*Header)(nil)

type Header struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Header) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "header", Node: d}, nil
}

var _ HTML = (*Footer)(nil)

type Footer struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Footer) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "footer", Node: d}, nil
}

var _ HTML = (*Main)(nil)

type Main struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Main) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "main", Node: d}, nil
}

var _ HTML = (*Nav)(nil)

type Nav struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Nav) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "nav", Node: d}, nil
}

var _ HTML = (*Aside)(nil)

type Aside struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Aside) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "aside", Node: d}, nil
}

var _ HTML = (*Small)(nil)

type Small struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (d *Small) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "small", Node: d}, nil
}

var _ HTML = (*Figure)(nil)

type Figure struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (p *Figure) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "figure", Node: p}, nil
}

var _ HTML = (*FigCaption)(nil)

type FigCaption struct {
	ID    string                 `h:"attr"`
	Class string                 `h:"attr"`
	Style string                 `h:"attr"`
	Data  map[string]interface{} `h:"dict"`
	Inner HTML                   `h:"inner"`
}

func (p *FigCaption) HTML(ctx context.Context) (HTML, error) {
	return &ReflectNode{Tag: "figcaption", Node: p}, nil
}
