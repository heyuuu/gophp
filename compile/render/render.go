package render

import (
	"github.com/heyuuu/gophp/compile/ir"
)

func Render(proj *ir.Project) (map[string]string, error) {
	p := defaultPrinter()

	var result = make(map[string]string)
	for _, ns := range proj.Namespaces() {
		p.reset()
		p.pNamespace(ns)
		content, err := p.result()
		if err != nil {
			return nil, err
		}

		name := ns.Name
		result[name] = content
	}
	return result, nil
}

// NameResolver
type NameResolver interface {
	Namespace(string) string
	Class(string) string
}

type defaultNameResolver struct{}

func newDefaultNameResolver() *defaultNameResolver {
	return &defaultNameResolver{}
}

func (d defaultNameResolver) Namespace(name string) string {
	return name
}

func (d defaultNameResolver) Class(name string) string {
	return name
}
