package compile

import (
	"github.com/heyuuu/gophp/compile/ir"
	"github.com/heyuuu/gophp/compile/parser"
)

// Compiler
type Compiler struct{}

func (c *Compiler) Compile(sources Sources) (*ir.Project, error) {
	proj := ir.NewProject()
	_, err := sources.EachSource(func(source Source) bool {
		err := c.compileSource(proj, source)
		return err == nil
	})
	if err != nil {
		return nil, err
	}

	return proj, nil
}

func (c *Compiler) compileSource(proj *ir.Project, source Source) error {
	astFile, err := parser.ParseCode(source.Content)
	if err != nil {
		return err
	}

	irFile, err := ir.ParseAstFile(astFile)
	if err != nil {
		return err
	}

	return proj.AddFile(source.RelativePath, irFile)
}
