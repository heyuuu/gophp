// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

func PHP_MAIL_BUILD_HEADER_CHECK(target string, s zend.SmartStr, key *zend.ZendString, val *zend.Zval) {
	for {
		if zend.Z_TYPE_P(val) == zend.IS_STRING {
			PhpMailBuildHeadersElem(&s, key, val)
		} else if zend.Z_TYPE_P(val) == zend.IS_ARRAY {
			if !(strncasecmp(target, zend.ZSTR_VAL(key), zend.ZSTR_LEN(key))) {
				core.PhpErrorDocref(nil, zend.E_WARNING, "'%s' header must be at most one header. Array is passed for '%s'", target, target)
				continue
			}
			PhpMailBuildHeadersElems(&s, key, val)
		} else {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header element '%s' cannot be other than string or array.", zend.ZSTR_VAL(key))
		}
		break
	}
}
func PHP_MAIL_BUILD_HEADER_DEFAULT(s zend.SmartStr, key *zend.ZendString, val *zend.Zval) {
	if zend.Z_TYPE_P(val) == zend.IS_STRING {
		PhpMailBuildHeadersElem(&s, key, val)
	} else if zend.Z_TYPE_P(val) == zend.IS_ARRAY {
		PhpMailBuildHeadersElems(&s, key, val)
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Extra header element '%s' cannot be other than string or array.", zend.ZSTR_VAL(key))
	}
}
