package vardumper

import (
	"io"
	"os"
	"strings"
)

func Print(v any) error {
	return Fprint(os.Stdout, v)
}

func Sprint(v any) string {
	var buf strings.Builder
	_ = Fprint(&buf, v) // 使用 strings.Builder 做 io.Writer 时，预期不会产生 error
	return buf.String()
}

func Fprint(w io.Writer, v any) error {
	return fprint(w, v)
}
