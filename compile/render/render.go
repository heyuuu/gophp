package render

import (
	"github.com/heyuuu/gophp/compile/ir"
)

// functions
func Render(proj *ir.Project) (map[string]string, error) {
	return defaultConfig.Render(proj)
}

// Config
var defaultConfig = &Config{}

type Config struct{}

func (c *Config) Render(proj *ir.Project) (map[string]string, error) {
	r := newRender(c)

	var result = make(map[string]string)
	for _, ns := range proj.Namespaces() {
		name, content, err := r.renderNamespace(ns)
		if err != nil {
			return nil, err
		}
		result[name] = content
	}
	return result, nil
}

// NameResolver
type NameResolver func(ir.Name) string

func defaultNameResolver(name ir.Name) string {
	return name.ToCodeString()
}

// render
type render struct {
	nameResolver NameResolver
}

func newRender(c *Config) *render {
	return &render{
		nameResolver: defaultNameResolver,
	}
}

func (r *render) renderNamespace(ns *ir.Namespace) (name string, content string, err error) {
	// build ast
	// todo

	// render
	// todo

	return name, content, nil
}
