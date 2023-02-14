package filesplit

import (
	"fmt"
	"sik/script/util"
	"testing"
)

func TestFuncFileSplit(t *testing.T) {
	var fileMap = map[string]string{
		//"/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_compile.func.go": "/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_compile_%d.go",
		//"/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_API.func.go":     "/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_API_%d.go",
		//"/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_execute.func.go": "/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_execute_%d.go",
	}
	for file, newFileTpl := range fileMap {
		code := util.MustReadFileString(file)
		splitFiles := FuncFileSplit(code, 10)
		for i, splitFile := range splitFiles {
			util.MustWriteFileString(fmt.Sprintf(newFileTpl, i), splitFile)
		}
	}
}
