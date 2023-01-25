// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

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
	{
		"assert.active",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetActive())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"assert.active\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.bail",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetBail())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.bail\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.warning",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetWarning())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"1",
		nil,
		b.SizeOf("\"1\"") - 1,
		b.SizeOf("\"assert.warning\"") - 1,
		core.PHP_INI_ALL,
	},
	{"assert.callback", OnChangeCallback, nil, nil, nil, nil, nil, b.SizeOf("NULL") - 1, b.SizeOf("\"assert.callback\"") - 1, core.PHP_INI_ALL},
	{
		"assert.quiet_eval",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetQuietEval())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.quiet_eval\"") - 1,
		core.PHP_INI_ALL,
	},
	{
		"assert.exception",
		zend.OnUpdateBool,
		any(zend_long((*byte)(&((*ZendAssertGlobals)(nil).GetException())) - (*byte)(nil))),
		any(&AssertGlobals),
		nil,
		"0",
		nil,
		b.SizeOf("\"0\"") - 1,
		b.SizeOf("\"assert.exception\"") - 1,
		core.PHP_INI_ALL,
	},
	{nil, nil, nil, nil, nil, nil, nil, 0, 0, 0},
}
