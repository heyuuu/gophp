package astutil

import (
	"go/ast"
	"go/token"
	"sort"
)

type FileBuilder struct {
	name    string
	imports map[string]bool
	decls   []ast.Decl
}

func NewFileBuilder(name string) *FileBuilder {
	return &FileBuilder{name: name, imports: make(map[string]bool)}
}
func (f *FileBuilder) AddImport(name string) { f.imports[name] = true }
func (f *FileBuilder) AddDecl(decl ast.Decl) { f.decls = append(f.decls, decl) }

func (f *FileBuilder) Build() *ast.File {
	var decls []ast.Decl
	if importDecl := f.buildImportDecl(); importDecl != nil {
		decls = append(decls, importDecl)
	}
	decls = append(decls, f.decls...)

	return &ast.File{
		Name:  Ident(f.name),
		Decls: decls,
	}
}

func (f *FileBuilder) buildImportDecl() *ast.GenDecl {
	if len(f.imports) == 0 {
		return nil
	}

	var names []string
	for name, _ := range f.imports {
		names = append(names, name)
	}
	sort.Strings(names)

	importSpecs := make([]ast.Spec, 0, len(f.imports))
	for _, name := range names {
		importSpecs = append(importSpecs, &ast.ImportSpec{
			Path: StrLit(name),
		})
	}

	return &ast.GenDecl{Tok: token.IMPORT, Specs: importSpecs}
}
