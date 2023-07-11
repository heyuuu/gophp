package vardumper

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Print(v any) error {
	fmt.Println()
	return Fprint(os.Stdout, v)
}

func Sprint(v any) (string, error) {
	var buf strings.Builder
	err := Fprint(&buf, v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Fprint(w io.Writer, v any) error {
	return fprint(w, v)
}
