package parser

import (
	"github.com/heyuuu/gophp/compile/ast"
	"github.com/heyuuu/gophp/compile/parser/internal/phpparse"
	"os"
	"path/filepath"
)

// 设置项目目录，用于在 test 等场景下，临时解决程序运行目录不为项目根目录导致找不到脚本的问题
func SetProjRoot(projRoot string) {
	phpparse.SetScriptPath(filepath.Join(projRoot, "tools/parser/parser.php"))
}

func ParseCode(code string) (*ast.File, error) {
	return phpparse.ParseCode(code)
}

func ParseFile(file string) (*ast.File, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return phpparse.ParseCode(string(bytes))
}
