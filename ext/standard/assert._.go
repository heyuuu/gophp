// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/assert.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

var AssertGlobals ZendAssertGlobals
var AssertionErrorCe *zend.ZendClassEntry

const (
	ASSERT_ACTIVE = 1
	ASSERT_CALLBACK
	ASSERT_BAIL
	ASSERT_WARNING
	ASSERT_QUIET_EVAL
	ASSERT_EXCEPTION
)

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	zend.MakeZendIniEntryDef("assert.active", zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetActive()))-(*byte)(nil))), any(&AssertGlobals), nil, "1", nil, core.PHP_INI_ALL),
	zend.MakeZendIniEntryDef("assert.bail", zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetBail()))-(*byte)(nil))), any(&AssertGlobals), nil, "0", nil, core.PHP_INI_ALL),
	zend.MakeZendIniEntryDef("assert.warning", zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetWarning()))-(*byte)(nil))), any(&AssertGlobals), nil, "1", nil, core.PHP_INI_ALL),
	zend.MakeZendIniEntryDef("assert.callback", OnChangeCallback, nil, nil, nil, nil, nil, core.PHP_INI_ALL),
	zend.MakeZendIniEntryDef("assert.quiet_eval", zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetQuietEval()))-(*byte)(nil))), any(&AssertGlobals), nil, "0", nil, core.PHP_INI_ALL),
	zend.MakeZendIniEntryDef("assert.exception", zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetException()))-(*byte)(nil))), any(&AssertGlobals), nil, "0", nil, core.PHP_INI_ALL),
}
