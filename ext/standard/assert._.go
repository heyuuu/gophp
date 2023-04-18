package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var AssertGlobals ZendAssertGlobals
var AssertionErrorCe *types.ClassEntry

const (
	ASSERT_ACTIVE = 1
	ASSERT_CALLBACK
	ASSERT_BAIL
	ASSERT_WARNING
	ASSERT_QUIET_EVAL
	ASSERT_EXCEPTION
)

var IniEntries []zend.ZendIniEntryDef = []zend.ZendIniEntryDef{
	*zend.NewZendIniEntryDef("assert.active", core.PHP_INI_ALL).Value("1").
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetActive()))-(*byte)(nil))), any(&AssertGlobals), nil,
		),
	*zend.NewZendIniEntryDef("assert.bail", core.PHP_INI_ALL).Value("0").
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetBail()))-(*byte)(nil))), any(&AssertGlobals), nil,
		),
	*zend.NewZendIniEntryDef("assert.warning", core.PHP_INI_ALL).Value("1").
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetWarning()))-(*byte)(nil))), any(&AssertGlobals), nil,
		),
	*zend.NewZendIniEntryDef("assert.callback", core.PHP_INI_ALL).
		OnModifyArgs(
			OnChangeCallback, nil, nil, nil,
		),
	*zend.NewZendIniEntryDef("assert.quiet_eval", core.PHP_INI_ALL).Value("0").
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetQuietEval()))-(*byte)(nil))), any(&AssertGlobals), nil,
		),
	*zend.NewZendIniEntryDef("assert.exception", core.PHP_INI_ALL).Value("0").
		OnModifyArgs(
			zend.OnUpdateBool, any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetException()))-(*byte)(nil))), any(&AssertGlobals), nil,
		),
}
