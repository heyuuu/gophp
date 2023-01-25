// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/php_mail.h>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// #define PHP_MAIL_H

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
