package standard

import (
	"sik/zend/def"
	"sik/zend/zpp"
)

// generate by ZifLevenshtein
var DefZifLevenshtein = def.DefFunc("levenshtein", 2, 5, []def.ArgInfo{{Name: "str1"}, {Name: "str2"}, {Name: "cost_ins"}, {Name: "cost_rep"}, {Name: "cost_del"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 5, 0)
	str1 := fp.ParseZval()
	str2 := fp.ParseZval()
	fp.StartOptional()
	cost_ins := fp.ParseZval()
	cost_rep := fp.ParseZval()
	cost_del := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLevenshtein(executeData, returnValue, str1, str2, nil, cost_ins, cost_rep, cost_del)
})
