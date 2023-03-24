package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifLevenshtein
var DefZifLevenshtein = def.DefFunc("levenshtein", 2, 5, []def.ArgInfo{{name: "str1"}, {name: "str2"}, {name: "cost_ins"}, {name: "cost_rep"}, {name: "cost_del"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
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
