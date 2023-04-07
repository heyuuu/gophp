package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifVersionCompare
var DefZifVersionCompare = def.DefFunc("version_compare", 2, 3, []def.ArgInfo{{Name: "ver1"}, {Name: "ver2"}, {Name: "oper"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	ver1 := fp.ParseZval()
	ver2 := fp.ParseZval()
	fp.StartOptional()
	oper := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifVersionCompare(ver1, ver2, nil, oper)
	returnValue.SetBy(ret)
})
