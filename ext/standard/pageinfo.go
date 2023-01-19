// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/pageinfo.h>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// #define PAGEINFO_H

// Source: <ext/standard/pageinfo.c>

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
   | Author: Jim Winstead <jimw@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "pageinfo.h"

// # include "SAPI.h"

// # include < stdio . h >

// # include < stdlib . h >

// # include < pwd . h >

// # include < grp . h >

// # include < unistd . h >

// # include < sys / stat . h >

// # include < sys / types . h >

// # include "ext/standard/basic_functions.h"

/* {{{ php_statpage
 */

func PhpStatpage() {
	var pstat *zend.ZendStatT
	pstat = core.SapiGetStat()
	if BasicGlobals.GetPageUid() == -1 || BasicGlobals.GetPageGid() == -1 {
		if pstat != nil {
			BasicGlobals.SetPageUid(pstat.st_uid)
			BasicGlobals.SetPageGid(pstat.st_gid)
			BasicGlobals.SetPageInode(pstat.st_ino)
			BasicGlobals.SetPageMtime(pstat.st_mtime)
		} else {
			BasicGlobals.SetPageUid(getuid())
			BasicGlobals.SetPageGid(getgid())
		}
	}
}

/* }}} */

func PhpGetuid() zend.ZendLong {
	PhpStatpage()
	return BasicGlobals.GetPageUid()
}

/* }}} */

func PhpGetgid() zend.ZendLong {
	PhpStatpage()
	return BasicGlobals.GetPageGid()
}

/* {{{ proto int getmyuid(void)
   Get PHP script owner's UID */

func ZifGetmyuid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var uid zend.ZendLong
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	uid = PhpGetuid()
	if uid < 0 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = uid
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func ZifGetmygid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var gid zend.ZendLong
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	gid = PhpGetgid()
	if gid < 0 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = gid
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func ZifGetmypid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pid zend.ZendLong
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	pid = getpid()
	if pid < 0 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = pid
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func ZifGetmyinode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	PhpStatpage()
	if BasicGlobals.GetPageInode() < 0 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = BasicGlobals.GetPageInode()
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func PhpGetlastmod() int64 {
	PhpStatpage()
	return BasicGlobals.GetPageMtime()
}

/* {{{ proto int getlastmod(void)
   Get time of last page modification */

func ZifGetlastmod(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var lm zend.ZendLong
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	lm = PhpGetlastmod()
	if lm < 0 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = lm
		__z.u1.type_info = 4
		return
	}
}

/* }}} */
