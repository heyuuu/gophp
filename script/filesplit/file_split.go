package filesplit

import "strings"

func FuncFileSplit(code string, n int) []string {
	lines := strings.Split(code, "\n")
	chunkSize := len(lines)/n + 1

	headerBuilder := strings.Builder{}
	fileBuilders := make([]strings.Builder, n)

	builder := &headerBuilder
	for i, line := range lines {
		if strings.HasPrefix(line, "func ") {
			builder = &fileBuilders[i/chunkSize]
		}

		builder.WriteString(line)
		builder.WriteByte('\n')
	}

	header := headerBuilder.String()
	files := make([]string, n)
	for i := range files {
		files[i] = header + "\n\n" + fileBuilders[i].String()
	}
	return files
}
