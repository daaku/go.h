package h

import (
	"fmt"
	"html"
	"io"
	"reflect"
	"strconv"
)

type Attributes map[string]interface{}

// Render an attribute value.
func writeValue(w io.Writer, i interface{}) (int, error) {
	var res string
	value := reflect.ValueOf(i)
	switch value.Kind() {
	case reflect.Bool:
		res = strconv.FormatBool(value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = strconv.FormatInt(value.Int(), 10)
	case reflect.Float32, reflect.Float64:
		res = strconv.FormatFloat(value.Float(), 'E', 3, 64)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = strconv.FormatUint(value.Uint(), 10)
	case reflect.String:
		res = value.String()
	default:
		return 0, fmt.Errorf(
			`Could not write attribute value "%v" with kind %s`, i, value.Kind())
	}
	return fmt.Fprint(w, html.EscapeString(res))
}

// Render attributes with using the optional key prefix.
func (attrs Attributes) Write(w io.Writer, prefix string) (int, error) {
	written := 0
	i := 0
	var err error
	for key, val := range attrs {
		i, err = fmt.Fprintf(w, ` %s="`, prefix+key)
		written += i
		if err != nil {
			return written, err
		}
		i, err = writeValue(w, val)
		written += i
		if err != nil {
			return written, err
		}
		i, err = fmt.Fprint(w, `"`)
		written += i
		if err != nil {
			return written, err
		}
	}
	return written, nil
}