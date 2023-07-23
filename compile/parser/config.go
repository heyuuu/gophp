package parser

type Config interface {
	EachSource(func(relativePath string, content string) bool)
	LoadClass(class string) []string
}
