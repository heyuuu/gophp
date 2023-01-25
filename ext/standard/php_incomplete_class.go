// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/standard/php_incomplete_class.h>

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
   | Author:  Sascha Schumann <sascha@schumann.cx>                        |
   +----------------------------------------------------------------------+
*/

// #define PHP_INCOMPLETE_CLASS_H

// # include "ext/standard/basic_functions.h"

const PHP_IC_ENTRY *zend.ZendClassEntry = BG(incomplete_class)

func PHP_SET_CLASS_ATTRIBUTES(struc *zend.Zval) {
	if zend.Z_OBJCE_P(struc) == BG(incomplete_class) {
		class_name = PhpLookupClassName(struc)
		if !class_name {
			class_name = zend.ZendStringInit(INCOMPLETE_CLASS, b.SizeOf("INCOMPLETE_CLASS")-1, 0)
		}
		incomplete_class = 1
	} else {
		class_name = zend.ZendStringCopy(zend.Z_OBJCE_P(struc).name)
	}
}
func PHP_CLEANUP_CLASS_ATTRIBUTES() { zend.ZendStringReleaseEx(class_name, 0) }

// #define PHP_CLASS_ATTRIBUTES       zend_string * class_name ; zend_bool incomplete_class ZEND_ATTRIBUTE_UNUSED = 0

const INCOMPLETE_CLASS = "__PHP_Incomplete_Class"
const MAGIC_MEMBER = "__PHP_Incomplete_Class_Name"
