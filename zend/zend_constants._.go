// <<generate>>

package zend

import "sik/zend/types"

// Source: <Zend/zend_constants.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

const CONST_CS = 1 << 0
const CONST_PERSISTENT = 1 << 1
const CONST_CT_SUBST = 1 << 2
const CONST_NO_FILE_CACHE = 1 << 3
const PHP_USER_CONSTANT = 0x7fffff

/* Flag for zend_get_constant_ex(). Must not class with ZEND_FETCH_CLASS_* flags. */

const ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK = 0x1000

const ZEND_CONSTANT_DTOR types.DtorFuncT = FreeZendConstant

// Source: <Zend/zend_constants.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/* Protection from recursive self-referencing class constants */

const IS_CONSTANT_VISITED_MARK = 0x80
