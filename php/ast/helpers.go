package ast

import (
	"unicode"
	"unicode/utf8"
)

func isExported(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(ch)
}
