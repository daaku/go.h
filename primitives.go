package h

import (
	"fmt"
	"html"
	"io"
	"reflect"
	"strings"

	"golang.org/x/net/context"
)

var _ HTML = (*Frag)(nil)
var _ Primitive = (*Frag)(nil)

type Frag []HTML

func (f *Frag) HTML(ctx context.Context) (HTML, error) {
	return f, fmt.Errorf("Frag.HTML called for %s", f)
}

func (f *Frag) Append(h HTML) {
	*f = append(*f, h)
}

func (f *Frag) Write(ctx context.Context, w io.Writer) (int, error) {
	written := 0
	for _, e := range *f {
		i, err := Write(ctx, w, e)
		written += i
		if err != nil {
			return written, err
		}
	}
	return written, nil
}

var _ HTML = String("")

type String string

func (s String) HTML(ctx context.Context) (HTML, error) {
	return Unsafe(html.EscapeString(string(s))), nil
}

var _ HTML = Unsafe("")
var _ Primitive = Unsafe("")

type Unsafe string

func (u Unsafe) HTML(ctx context.Context) (HTML, error) {
	return u, fmt.Errorf("Unsafe.HTML called for %s", u)
}

func (u Unsafe) Write(ctx context.Context, w io.Writer) (int, error) {
	return fmt.Fprint(w, u)
}

var _ HTML = UnsafeBytes(nil)
var _ Primitive = UnsafeBytes(nil)

type UnsafeBytes []byte

func (u UnsafeBytes) HTML(ctx context.Context) (HTML, error) {
	return u, fmt.Errorf("UnsafeBytes.HTML called for %s", u)
}

func (u UnsafeBytes) Write(ctx context.Context, w io.Writer) (int, error) {
	return fmt.Fprintf(w, "%s", u)
}

var _ HTML = (*ReflectNode)(nil)
var _ Primitive = (*ReflectNode)(nil)

type ReflectNode struct {
	Tag         string
	Node        interface{}
	SelfClosing bool
}

func (n *ReflectNode) HTML(ctx context.Context) (HTML, error) {
	return n, fmt.Errorf("Called HTML for ReflectNode: %+v", n)
}

// Write the generated markup for a ReflectNode.
func (n *ReflectNode) Write(ctx context.Context, w io.Writer) (int, error) {
	written := 0
	i := 0
	var err error

	i, err = fmt.Fprint(w, "<", n.Tag)
	written += i
	if err != nil {
		return written, err
	}

	i, err = n.writeAttributes(w)
	written += i
	if err != nil {
		return written, err
	}

	i, err = fmt.Fprint(w, ">")
	written += i
	if err != nil {
		return written, err
	}

	i, err = n.writeInner(ctx, w)
	written += i
	if err != nil {
		return written, err
	}

	if !n.SelfClosing {
		i, err = fmt.Fprint(w, "</", n.Tag, ">")
		written += i
		if err != nil {
			return written, err
		}
	}

	return written, nil
}

// Use reflection to write attributes.
func (n *ReflectNode) writeAttributes(w io.Writer) (int, error) {
	value := reflect.ValueOf(n.Node).Elem()
	typeOf := value.Type()
	written := 0
	tmp := 0
	var err error
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		switch field.Tag.Get("h") {
		case "attr":
			tmp, err = writeKeyValue(
				w, strings.ToLower(field.Name), value.Field(i).Interface())
			written += tmp
			if err != nil {
				return written, err
			}
		case "dict":
			val := value.Field(i).Interface()
			rawAttrs, ok := val.(map[string]interface{})
			if !ok {
				return written, fmt.Errorf(
					"Invalid dict2 value: %+v of type %T", val, val)
			}
			attrs := Attributes(rawAttrs)
			tmp, err = attrs.Write(w, strings.ToLower(field.Name)+"-")
			written += tmp
			if err != nil {
				return written, err
			}
		}
	}
	return written, nil
}

// Use reflection to write inner HTML.
func (n *ReflectNode) writeInner(ctx context.Context, w io.Writer) (int, error) {
	value := reflect.ValueOf(n.Node).Elem()
	typeOf := value.Type()
	written := 0
	tmp := 0
	var err error
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		if field.Tag.Get("h") != "inner" {
			continue
		}

		fieldValue := value.Field(i).Interface()
		if fieldValue == nil {
			continue
		}
		html, ok := fieldValue.(HTML)
		if !ok {
			return written, fmt.Errorf(
				"Field %s was marked as inner but does not satisfy the HTML interface",
				field.Name)
		}
		tmp, err = Write(ctx, w, html)
		written += tmp
		if err != nil {
			return written, err
		}
	}
	return written, nil
}