package replace

import (
	"fmt"
	"sik/script/util"
	"sik/script/util/finder"
	"testing"
)

func TestReplaceMakeArgInfo(t *testing.T) {
	f := finder.DefaultProjectFinder()
	f.Walk(func(file finder.FileInfo) {
		//fmt.Println(">>>> FILE: " + file.RelativePath)
		code := util.MustReadFileString(file.Path)
		result := ReplaceMakeArgInfo(code)
		if result != code {
			util.MustWriteFileString(file.Path, result)
		}
	})
}

func pSet(name string, record map[string]int) {
	if len(record) == 0 {
		return
	}

	fmt.Println("Name: " + name)
	for v, _ := range record {
		fmt.Println("    " + v)
	}
}

func TestReplaceMakeArgInfoDev(t *testing.T) {
	file := "/Users/heyu/Code/sik/sik-go-gen-2/zend/zend_builtin_functions._.go"
	code := util.MustReadFileString(file)

	ReplaceMakeArgInfo(code)
}
