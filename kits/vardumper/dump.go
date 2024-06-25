package vardumper

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

func Dump(v any) string {
	return DumpEx(v, Config{})
}

func DumpEx(v any, c Config) string {
	// default ident
	if c.Ident == "" {
		c.Ident = "    "
	}

	// setup printer
	p := printer{
		c:         c,
		ptrmap:    make(map[any]int),
		isNewLine: true,
	}

	// print x
	p.printValue(reflect.ValueOf(v))
	p.printf("\n")

	return p.buf.String()
}

type Config struct {
	ShowAnonymousField bool
	ShowLineNum        bool
	Ident              string
}

type printer struct {
	c         Config
	buf       strings.Builder
	ptrmap    map[any]int // *T -> line number
	indent    int         // current indentation level
	line      int         // current line number
	isNewLine bool
}

func (p *printer) printIdent() {
	if p.c.ShowLineNum {
		p.buf.WriteString(fmt.Sprintf("%4d: ", p.line))
	}
	p.buf.WriteString(strings.Repeat(p.c.Ident, p.indent))
}

func (p *printer) printf(s string, args ...any) {
	if len(args) > 0 {
		s = fmt.Sprintf(s, args...)
	}

	for s != "" {
		if p.isNewLine {
			p.printIdent()
			p.isNewLine = false
		}
		if idx := strings.IndexByte(s, '\n'); idx >= 0 {
			p.buf.WriteString(s[:idx+1])
			s = s[idx+1:]
			p.isNewLine = true
			p.line++
		} else {
			p.buf.WriteString(s)
			s = ""
		}
	}
}

func (p *printer) printValue(x reflect.Value) {
	if isNilValue(x) {
		p.printf("nil")
		return
	}

	switch x.Kind() {
	case reflect.Interface:
		p.printValue(x.Elem())

	case reflect.Map:
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for _, key := range x.MapKeys() {
				p.printValue(key)
				p.printf(": ")
				p.printValue(x.MapIndex(key))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Pointer:
		p.printf("*")
		// type-checked ASTs may contain cycles - use ptrmap
		// to keep track of objects that have been printed
		// already and print the respective line number instead
		ptr := x.Interface()
		if line, exists := p.ptrmap[ptr]; exists {
			p.printf("(obj @ %d)", line)
		} else {
			p.ptrmap[ptr] = p.line
			p.printValue(x.Elem())
		}

	case reflect.Array:
		p.printf("%s {", x.Type())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.printValue(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Slice:
		if s, ok := x.Interface().([]byte); ok {
			p.printf("%#q", s)
			return
		}
		p.printf("%s (len = %d) {", x.Type(), x.Len())
		if x.Len() > 0 {
			p.indent++
			p.printf("\n")
			for i, n := 0, x.Len(); i < n; i++ {
				p.printf("%d: ", i)
				p.printValue(x.Index(i))
				p.printf("\n")
			}
			p.indent--
		}
		p.printf("}")

	case reflect.Struct:
		t := x.Type()
		p.printf("%s {", t)
		p.indent++

		first := true
		p.eachFields("", x, func(name string, value reflect.Value) {
			if first {
				p.printf("\n")
				first = false
			}
			p.printf("%s: ", name)
			p.printValue(value)
			p.printf("\n")
		})

		p.indent--
		p.printf("}")

	default:
		v := x.Interface()
		switch v := v.(type) {
		case string:
			// print strings in quotes
			p.printf("%q", v)
			return
		}
		// default
		p.printf("%s(%v)", x.Type(), v)
	}
}

func (p *printer) eachFields(prefix string, v reflect.Value, f func(name string, v reflect.Value)) {
	if v.Kind() != reflect.Struct {
		return
	}

	t := v.Type()
	for i, n := 0, t.NumField(); i < n; i++ {
		fieldTyp := t.Field(i)
		fieldVal := v.Field(i)

		if isExported(fieldTyp.Name) {
			f(prefix+fieldTyp.Name, fieldVal)
		} else if fieldTyp.Anonymous && p.c.ShowAnonymousField {
			p.eachFields(prefix+fieldTyp.Name+".", fieldVal, f)
		}
	}
}

func isExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}

func isNilValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Invalid: // v 未初始化时的类型，此 v 值可来源于 reflect.ValueOf(nil)
		return true
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	}
	return false
}
