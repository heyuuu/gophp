package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifOpenlog
var DefZifOpenlog = def.DefFunc("openlog", 3, 3, []def.ArgInfo{{name: "ident"}, {name: "option"}, {name: "facility"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	ident := fp.ParseZval()
	option := fp.ParseZval()
	facility := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOpenlog(executeData, returnValue, ident, option, facility)
})

// generate by ZifCloselog
var DefZifCloselog = def.DefFunc("closelog", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifCloselog(executeData, returnValue)
})

// generate by ZifSyslog
var DefZifSyslog = def.DefFunc("syslog", 2, 2, []def.ArgInfo{{name: "priority"}, {name: "message"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	priority := fp.ParseZval()
	message := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSyslog(executeData, returnValue, priority, message)
})
