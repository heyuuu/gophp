package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifOpenlog
var DefZifOpenlog = def.DefFunc("openlog", 3, 3, []def.ArgInfo{{Name: "ident"}, {Name: "option"}, {Name: "facility"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
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
var DefZifCloselog = def.DefFunc("closelog", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifCloselog(executeData, returnValue)
})

// generate by ZifSyslog
var DefZifSyslog = def.DefFunc("syslog", 2, 2, []def.ArgInfo{{Name: "priority"}, {Name: "message"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	priority := fp.ParseZval()
	message := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSyslog(executeData, returnValue, priority, message)
})
