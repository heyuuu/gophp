// <<generate>>

package standard

import "sik/zend"

// Source: <ext/standard/php_var.h>

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
   | Author: Jani Lehtimäki <jkl@njet.net>                                |
   +----------------------------------------------------------------------+
*/

// #define PHP_VAR_H

// # include "ext/standard/basic_functions.h"

// # include "zend_smart_str_public.h"

type PhpSerializeDataT *PhpSerializeData
type PhpUnserializeDataT *PhpUnserializeData

var PhpVarUnserializeRef func(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int
var PhpVarUnserializeIntern func(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int

// #define PHP_VAR_SERIALIZE_INIT(d) ( d ) = php_var_serialize_init ( )

// #define PHP_VAR_SERIALIZE_DESTROY(d) php_var_serialize_destroy ( d )

// #define PHP_VAR_UNSERIALIZE_INIT(d) ( d ) = php_var_unserialize_init ( )

// #define PHP_VAR_UNSERIALIZE_DESTROY(d) php_var_unserialize_destroy ( d )
