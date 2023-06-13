package zend

import "github.com/heyuuu/gophp/php/types"

var currCompiler = &Compiler{}

func CurrCompiler() *Compiler {
	return currCompiler
}

/**
 * Compiler
 */
type Compiler struct {
	lineno uint32 // 从 CG__.zend_lineno 中剥离，专用于 OpCode 生成过程
}

// helpers
func (compiler *Compiler) initOp(op *types.ZendOp) *types.ZendOp {
	if op == nil {
		op = types.NewOp(ZEND_NOP, compiler.lineno)
	} else {
		op.SetNop()
		op.SetExtendedValue(0)
		op.SetLineno(compiler.lineno)
	}
	return op
}

func (compiler *Compiler) setLineno(lineno uint32) {
	compiler.lineno = lineno
}

func (compiler *Compiler) setLinenoByAst(ast *ZendAst) {
	compiler.setLineno(ast.Lineno())
}

func (compiler *Compiler) setLinenoByOpline(opline *types.ZendOp) {
	compiler.setLineno(opline.GetLineno())
}

func (compiler *Compiler) setLinenoByDeclEnd(decl *ZendAstDecl) {
	compiler.setLineno(decl.GetEndLineno())
}

// compile
