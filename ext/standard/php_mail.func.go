// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

func PHP_MAIL_BUILD_HEADER_CHECK(target string, s zend.SmartStr, key *types.ZendString, val *types.Zval) {
	for {
		if val.IsType(types.IS_STRING) {
			PhpMailBuildHeadersElem(&s, key, val)
		} else if val.IsType(types.IS_ARRAY) {
			if !(strncasecmp(target, key.GetVal(), key.GetLen())) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "'%s' header must be at most one header. Array is passed for '%s'", target, target)
				continue
			}
			PhpMailBuildHeadersElems(&s, key, val)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header element '%s' cannot be other than string or array.", key.GetVal())
		}
		break
	}
}
func PHP_MAIL_BUILD_HEADER_DEFAULT(s zend.SmartStr, key *types.ZendString, val *types.Zval) {
	if val.IsType(types.IS_STRING) {
		PhpMailBuildHeadersElem(&s, key, val)
	} else if val.IsType(types.IS_ARRAY) {
		PhpMailBuildHeadersElems(&s, key, val)
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header element '%s' cannot be other than string or array.", key.GetVal())
	}
}
