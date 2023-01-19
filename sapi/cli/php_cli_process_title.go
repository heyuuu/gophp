// <<generate>>

package cli

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <sapi/cli/php_cli_process_title.h>

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
  | Author: Keyur Govande (kgovande@gmail.com)                           |
  +----------------------------------------------------------------------+
*/

// #define PHP_PS_TITLE_HEADER

var ArginfoCliSetProcessTitle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"title", 0, 0, 0}}
var ArginfoCliGetProcessTitle []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

// Source: <sapi/cli/php_cli_process_title.c>

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
  | Author: Keyur Govande (kgovande@gmail.com)                           |
  +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_cli_process_title.h"

// # include "ps_title.h"

/* {{{ proto bool cli_set_process_title(string arg)
   Return a boolean to confirm if the process title was successfully changed or not */

func ZifCliSetProcessTitle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var title *byte = nil
	var title_len int
	var rc int
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s", &title, &title_len) == zend.FAILURE {
		return
	}
	rc = SetPsTitle(title)
	if rc == 0 {
		return_value.u1.type_info = 3
		return
	}
	core.PhpErrorDocref(nil, 1<<1, "cli_set_process_title had an error: %s", PsTitleErrno(rc))
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifCliGetProcessTitle(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var length int = 0
	var title *byte = nil
	var rc int
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	rc = GetPsTitle(&length, &title)
	if rc != 0 {
		core.PhpErrorDocref(nil, 1<<1, "cli_get_process_title had an error: %s", PsTitleErrno(rc))
		return_value.u1.type_info = 1
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(title, length, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */
