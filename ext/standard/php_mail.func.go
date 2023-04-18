package standard

import (
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
)

func PHP_MAIL_BUILD_HEADER_CHECK(target string, s zend.SmartStr, key *types2.String, val *types2.Zval) {
	for {
		if val.IsType(types2.IS_STRING) {
			PhpMailBuildHeadersElem(&s, key, val)
		} else if val.IsType(types2.IS_ARRAY) {
			if !(strncasecmp(target, key.GetVal(), key.GetLen())) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "'%s' header must be at most one header. Array is passed for '%s'", target, target)
				continue
			}
			PhpMailBuildHeadersElems(&s, key, val)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header element '%s' cannot be other than string or array.", key.GetVal())
		}
		break
	}
}
func PHP_MAIL_BUILD_HEADER_DEFAULT(s zend.SmartStr, key *types2.String, val *types2.Zval) {
	if val.IsType(types2.IS_STRING) {
		PhpMailBuildHeadersElem(&s, key, val)
	} else if val.IsType(types2.IS_ARRAY) {
		PhpMailBuildHeadersElems(&s, key, val)
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Extra header element '%s' cannot be other than string or array.", key.GetVal())
	}
}
