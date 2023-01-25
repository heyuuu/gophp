// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_ini.h>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// #define ZEND_INI_H

const ZEND_INI_USER = 1 << 0
const ZEND_INI_PERDIR ZendLong = 1 << 1
const ZEND_INI_SYSTEM uint8 = 1 << 2
const ZEND_INI_ALL ZendLong = ZEND_INI_USER | ZEND_INI_PERDIR | ZEND_INI_SYSTEM

// #define ZEND_INI_MH(name) int name ( zend_ini_entry * entry , zend_string * new_value , void * mh_arg1 , void * mh_arg2 , void * mh_arg3 , int stage )

// #define ZEND_INI_DISP(name) ZEND_COLD void name ( zend_ini_entry * ini_entry , int type )

var ZendCopyIniDirectives func() int
var ZendIniRefreshCaches func(stage int)
var DisplayIniEntries func(module *ZendModuleEntry)

// #define ZEND_INI_BEGIN() static const zend_ini_entry_def ini_entries [ ] = {

// #define ZEND_INI_END() { NULL , NULL , NULL , NULL , NULL , NULL , NULL , 0 , 0 , 0 } } ;

// #define ZEND_INI_ENTRY3_EX(name,default_value,modifiable,on_modify,arg1,arg2,arg3,displayer) { name , on_modify , arg1 , arg2 , arg3 , default_value , displayer , sizeof ( default_value ) - 1 , sizeof ( name ) - 1 , modifiable } ,

// #define ZEND_INI_ENTRY3(name,default_value,modifiable,on_modify,arg1,arg2,arg3) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , arg3 , NULL )

// #define ZEND_INI_ENTRY2_EX(name,default_value,modifiable,on_modify,arg1,arg2,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , NULL , displayer )

// #define ZEND_INI_ENTRY2(name,default_value,modifiable,on_modify,arg1,arg2) ZEND_INI_ENTRY2_EX ( name , default_value , modifiable , on_modify , arg1 , arg2 , NULL )

// #define ZEND_INI_ENTRY1_EX(name,default_value,modifiable,on_modify,arg1,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , arg1 , NULL , NULL , displayer )

// #define ZEND_INI_ENTRY1(name,default_value,modifiable,on_modify,arg1) ZEND_INI_ENTRY1_EX ( name , default_value , modifiable , on_modify , arg1 , NULL )

// #define ZEND_INI_ENTRY_EX(name,default_value,modifiable,on_modify,displayer) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , NULL , NULL , NULL , displayer )

// #define ZEND_INI_ENTRY(name,default_value,modifiable,on_modify) ZEND_INI_ENTRY_EX ( name , default_value , modifiable , on_modify , NULL )

// #define STD_ZEND_INI_ENTRY(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr) ZEND_INI_ENTRY2 ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr )

// #define STD_ZEND_INI_ENTRY_EX(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr,displayer) ZEND_INI_ENTRY2_EX ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr , displayer )

// #define STD_ZEND_INI_BOOLEAN(name,default_value,modifiable,on_modify,property_name,struct_type,struct_ptr) ZEND_INI_ENTRY3_EX ( name , default_value , modifiable , on_modify , ( void * ) XtOffsetOf ( struct_type , property_name ) , ( void * ) & struct_ptr , NULL , zend_ini_boolean_displayer_cb )

func INI_INT(name string) ZendLong {
	return ZendIniLong(name, b.SizeOf("name")-1, 0)
}
func INI_FLT(name *byte) float64 {
	return ZendIniDouble(name, b.SizeOf("name")-1, 0)
}
func INI_STR(name string) *byte {
	return ZendIniStringEx(name, b.SizeOf("name")-1, 0, nil)
}
func INI_BOOL(name *byte) ZendBool { return ZendBool(INI_INT(name)) }
func INI_ORIG_INT(name *byte) ZendLong {
	return ZendIniLong(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_FLT(name *byte) float64 {
	return ZendIniDouble(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_STR(name *byte) *byte {
	return ZendIniString(name, b.SizeOf("name")-1, 1)
}
func INI_ORIG_BOOL(name *byte) ZendBool { return ZendBool(INI_ORIG_INT(name)) }
func REGISTER_INI_ENTRIES() int {
	return ZendRegisterIniEntries(IniEntries, module_number)
}
func UNREGISTER_INI_ENTRIES() { ZendUnregisterIniEntries(module_number) }
func DISPLAY_INI_ENTRIES()    { DisplayIniEntries(zend_module) }
func REGISTER_INI_DISPLAYER(name *byte, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	return ZendIniRegisterDisplayer(name, b.SizeOf("name")-1, displayer)
}
func REGISTER_INI_BOOLEAN(name *byte) int {
	return REGISTER_INI_DISPLAYER(name, ZendIniBooleanDisplayerCb)
}

/* Standard message handlers */

const ZEND_INI_DISPLAY_ORIG = 1
const ZEND_INI_DISPLAY_ACTIVE = 2
const ZEND_INI_STAGE_STARTUP = 1 << 0
const ZEND_INI_STAGE_SHUTDOWN = 1 << 1
const ZEND_INI_STAGE_ACTIVATE = 1 << 2
const ZEND_INI_STAGE_DEACTIVATE = 1 << 3
const ZEND_INI_STAGE_RUNTIME = 1 << 4
const ZEND_INI_STAGE_HTACCESS = 1 << 5
const ZEND_INI_STAGE_IN_REQUEST = ZEND_INI_STAGE_ACTIVATE | ZEND_INI_STAGE_DEACTIVATE | ZEND_INI_STAGE_RUNTIME | ZEND_INI_STAGE_HTACCESS

/* INI parsing engine */

type ZendIniParserCbT func(arg1 *Zval, arg2 *Zval, arg3 *Zval, callback_type int, arg any)

/* INI entries */

const ZEND_INI_PARSER_ENTRY = 1
const ZEND_INI_PARSER_SECTION = 2
const ZEND_INI_PARSER_POP_ENTRY = 3

// Source: <Zend/zend_ini.c>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_sort.h"

// # include "zend_API.h"

// # include "zend_ini.h"

// # include "zend_alloc.h"

// # include "zend_operators.h"

// # include "zend_strtod.h"

var RegisteredZendIniDirectives *HashTable

const NO_VALUE_PLAINTEXT = "no value"
const NO_VALUE_HTML = "<i>no value</i>"

/*
 * hash_apply functions
 */

func ZendRemoveIniEntries(el *Zval, arg any) int {
	var ini_entry *ZendIniEntry = (*ZendIniEntry)(Z_PTR_P(el))
	var module_number int = *((*int)(arg))
	return ini_entry.GetModuleNumber() == module_number
}

/* }}} */

func ZendRestoreIniEntryCb(ini_entry *ZendIniEntry, stage int) int {
	var result int = FAILURE
	if ini_entry.GetModified() != 0 {
		if ini_entry.GetOnModify() != nil {
			var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
			var __bailout JMP_BUF
			ExecutorGlobals.SetBailout(&__bailout)
			if SETJMP(__bailout) == 0 {

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

				result = ini_entry.GetOnModify()(ini_entry, ini_entry.GetOrigValue(), ini_entry.GetMhArg1(), ini_entry.GetMhArg2(), ini_entry.GetMhArg3(), stage)

				/* even if on_modify bails out, we have to continue on with restoring,
				   since there can be allocated variables that would be freed on MM shutdown
				   and would lead to memory corruption later ini entry is modified again */

			}
			ExecutorGlobals.SetBailout(__orig_bailout)
		}
		if stage == ZEND_INI_STAGE_RUNTIME && result == FAILURE {

			/* runtime failure is OK */

			return 1

			/* runtime failure is OK */

		}
		if ini_entry.GetValue() != ini_entry.GetOrigValue() {
			ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(ini_entry.GetOrigValue())
		ini_entry.SetModifiable(ini_entry.GetOrigModifiable())
		ini_entry.SetModified(0)
		ini_entry.SetOrigValue(nil)
		ini_entry.SetOrigModifiable(0)
	}
	return 0
}

/* }}} */

func FreeIniEntry(zv *Zval) {
	var entry *ZendIniEntry = (*ZendIniEntry)(Z_PTR_P(zv))
	ZendStringReleaseEx(entry.GetName(), 1)
	if entry.GetValue() != nil {
		ZendStringRelease(entry.GetValue())
	}
	if entry.GetOrigValue() != nil {
		ZendStringReleaseEx(entry.GetOrigValue(), 1)
	}
	Free(entry)
}

/* }}} */

func ZendIniStartup() int {
	RegisteredZendIniDirectives = (*HashTable)(Malloc(b.SizeOf("HashTable")))
	ExecutorGlobals.SetIniDirectives(RegisteredZendIniDirectives)
	ExecutorGlobals.SetModifiedIniDirectives(nil)
	ExecutorGlobals.SetErrorReportingIniEntry(nil)
	ZendHashInitEx(RegisteredZendIniDirectives, 128, nil, FreeIniEntry, 1, 0)
	return SUCCESS
}

/* }}} */

func ZendIniShutdown() int {
	ZendIniDtor(ExecutorGlobals.GetIniDirectives())
	return SUCCESS
}

/* }}} */

func ZendIniDtor(ini_directives *HashTable) {
	ZendHashDestroy(ini_directives)
	Free(ini_directives)
}

/* }}} */

func ZendIniGlobalShutdown() int {
	ZendHashDestroy(RegisteredZendIniDirectives)
	Free(RegisteredZendIniDirectives)
	return SUCCESS
}

/* }}} */

func ZendIniDeactivate() int {
	if ExecutorGlobals.GetModifiedIniDirectives() != nil {
		var ini_entry *ZendIniEntry
		for {
			var __ht *HashTable = ExecutorGlobals.GetModifiedIniDirectives()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				ini_entry = Z_PTR_P(_z)
				ZendRestoreIniEntryCb(ini_entry, ZEND_INI_STAGE_DEACTIVATE)
			}
			break
		}
		ZendHashDestroy(ExecutorGlobals.GetModifiedIniDirectives())
		FREE_HASHTABLE(ExecutorGlobals.GetModifiedIniDirectives())
		ExecutorGlobals.SetModifiedIniDirectives(nil)
	}
	return SUCCESS
}

/* }}} */

func IniKeyCompare(a any, b any) int {
	var f *Bucket
	var s *Bucket
	f = (*Bucket)(a)
	s = (*Bucket)(b)
	if f.GetKey() == nil && s.GetKey() == nil {
		if f.GetH() > s.GetH() {
			return -1
		} else if f.GetH() < s.GetH() {
			return 1
		}
		return 0
	} else if f.GetKey() == nil {
		return -1
	} else if s.GetKey() == nil {
		return 1
	} else {
		return ZendBinaryStrcasecmp(ZSTR_VAL(f.GetKey()), ZSTR_LEN(f.GetKey()), ZSTR_VAL(s.GetKey()), ZSTR_LEN(s.GetKey()))
	}
}

/* }}} */

func ZendIniSortEntries() {
	ZendHashSort(ExecutorGlobals.GetIniDirectives(), IniKeyCompare, 0)
}

/* }}} */

func ZendRegisterIniEntries(ini_entry *ZendIniEntryDef, module_number int) int {
	var p *ZendIniEntry
	var default_value *Zval
	var directives *HashTable = RegisteredZendIniDirectives
	for ini_entry.GetName() != nil {
		p = Pemalloc(b.SizeOf("zend_ini_entry"), 1)
		p.SetName(ZendStringInitInterned(ini_entry.GetName(), ini_entry.GetNameLength(), 1))
		p.SetOnModify(ini_entry.GetOnModify())
		p.SetMhArg1(ini_entry.GetMhArg1())
		p.SetMhArg2(ini_entry.GetMhArg2())
		p.SetMhArg3(ini_entry.GetMhArg3())
		p.SetValue(nil)
		p.SetOrigValue(nil)
		p.SetDisplayer(ini_entry.GetDisplayer())
		p.SetModifiable(ini_entry.GetModifiable())
		p.SetOrigModifiable(0)
		p.SetModified(0)
		p.SetModuleNumber(module_number)
		if ZendHashAddPtr(directives, p.GetName(), any(p)) == nil {
			if p.GetName() != nil {
				ZendStringReleaseEx(p.GetName(), 1)
			}
			ZendUnregisterIniEntries(module_number)
			return FAILURE
		}
		if b.Assign(&default_value, ZendGetConfigurationDirective(p.GetName())) != nil && (p.GetOnModify() == nil || p.GetOnModify()(p, Z_STR_P(default_value), p.GetMhArg1(), p.GetMhArg2(), p.GetMhArg3(), ZEND_INI_STAGE_STARTUP) == SUCCESS) {
			p.SetValue(ZendNewInternedString(ZendStringCopy(Z_STR_P(default_value))))
		} else {
			if ini_entry.GetValue() != nil {
				p.SetValue(ZendStringInitInterned(ini_entry.GetValue(), ini_entry.GetValueLength(), 1))
			} else {
				p.SetValue(nil)
			}
			if p.GetOnModify() != nil {
				p.GetOnModify()(p, p.GetValue(), p.GetMhArg1(), p.GetMhArg2(), p.GetMhArg3(), ZEND_INI_STAGE_STARTUP)
			}
		}
		ini_entry++
	}
	return SUCCESS
}

/* }}} */

func ZendUnregisterIniEntries(module_number int) {
	ZendHashApplyWithArgument(RegisteredZendIniDirectives, ZendRemoveIniEntries, any(&module_number))
}

/* }}} */

func ZendAlterIniEntry(name *ZendString, new_value *ZendString, modify_type int, stage int) int {
	return ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
}

/* }}} */

func ZendAlterIniEntryChars(name *ZendString, value string, value_length int, modify_type int, stage int) int {
	var ret int
	var new_value *ZendString
	new_value = ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, 0)
	ZendStringRelease(new_value)
	return ret
}

/* }}} */

func ZendAlterIniEntryCharsEx(name *ZendString, value *byte, value_length int, modify_type int, stage int, force_change int) int {
	var ret int
	var new_value *ZendString
	new_value = ZendStringInit(value, value_length, !(stage & ZEND_INI_STAGE_IN_REQUEST))
	ret = ZendAlterIniEntryEx(name, new_value, modify_type, stage, force_change)
	ZendStringRelease(new_value)
	return ret
}

/* }}} */

func ZendAlterIniEntryEx(name *ZendString, new_value *ZendString, modify_type int, stage int, force_change int) int {
	var ini_entry *ZendIniEntry
	var duplicate *ZendString
	var modifiable uint8
	var modified ZendBool
	if b.Assign(&ini_entry, ZendHashFindPtr(ExecutorGlobals.GetIniDirectives(), name)) == nil {
		return FAILURE
	}
	modifiable = ini_entry.GetModifiable()
	modified = ini_entry.GetModified()
	if stage == ZEND_INI_STAGE_ACTIVATE && modify_type == ZEND_INI_SYSTEM {
		ini_entry.SetModifiable(ZEND_INI_SYSTEM)
	}
	if force_change == 0 {
		if (ini_entry.GetModifiable() & modify_type) == 0 {
			return FAILURE
		}
	}
	if ExecutorGlobals.GetModifiedIniDirectives() == nil {
		ALLOC_HASHTABLE(ExecutorGlobals.GetModifiedIniDirectives())
		ZendHashInit(ExecutorGlobals.GetModifiedIniDirectives(), 8, nil, nil, 0)
	}
	if modified == 0 {
		ini_entry.SetOrigValue(ini_entry.GetValue())
		ini_entry.SetOrigModifiable(modifiable)
		ini_entry.SetModified(1)
		ZendHashAddPtr(ExecutorGlobals.GetModifiedIniDirectives(), ini_entry.GetName(), ini_entry)
	}
	duplicate = ZendStringCopy(new_value)
	if ini_entry.GetOnModify() == nil || ini_entry.GetOnModify()(ini_entry, duplicate, ini_entry.GetMhArg1(), ini_entry.GetMhArg2(), ini_entry.GetMhArg3(), stage) == SUCCESS {
		if modified != 0 && ini_entry.GetOrigValue() != ini_entry.GetValue() {
			ZendStringRelease(ini_entry.GetValue())
		}
		ini_entry.SetValue(duplicate)
	} else {
		ZendStringRelease(duplicate)
		return FAILURE
	}
	return SUCCESS
}

/* }}} */

func ZendRestoreIniEntry(name *ZendString, stage int) int {
	var ini_entry *ZendIniEntry
	if b.Assign(&ini_entry, ZendHashFindPtr(ExecutorGlobals.GetIniDirectives(), name)) == nil || stage == ZEND_INI_STAGE_RUNTIME && (ini_entry.GetModifiable()&ZEND_INI_USER) == 0 {
		return FAILURE
	}
	if ExecutorGlobals.GetModifiedIniDirectives() != nil {
		if ZendRestoreIniEntryCb(ini_entry, stage) == 0 {
			ZendHashDel(ExecutorGlobals.GetModifiedIniDirectives(), name)
		} else {
			return FAILURE
		}
	}
	return SUCCESS
}

/* }}} */

func ZendIniRegisterDisplayer(name *byte, name_length uint32, displayer func(ini_entry *ZendIniEntry, type_ int)) int {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(RegisteredZendIniDirectives, name, name_length)
	if ini_entry == nil {
		return FAILURE
	}
	ini_entry.SetDisplayer(displayer)
	return SUCCESS
}

/* }}} */

func ZendIniLong(name *byte, name_length int, orig int) ZendLong {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(ExecutorGlobals.GetIniDirectives(), name, name_length)
	if ini_entry != nil {
		if orig != 0 && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil {
				return ZEND_STRTOL(ZSTR_VAL(ini_entry.GetOrigValue()), nil, 0)
			} else {
				return 0
			}
		} else {
			if ini_entry.GetValue() != nil {
				return ZEND_STRTOL(ZSTR_VAL(ini_entry.GetValue()), nil, 0)
			} else {
				return 0
			}
		}
	}
	return 0
}

/* }}} */

func ZendIniDouble(name *byte, name_length int, orig int) float64 {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(ExecutorGlobals.GetIniDirectives(), name, name_length)
	if ini_entry != nil {
		if orig != 0 && ini_entry.GetModified() != 0 {
			return float64(b.CondF1(ini_entry.GetOrigValue() != nil, func() float64 { return ZendStrtod(ZSTR_VAL(ini_entry.GetOrigValue()), nil) }, 0.0))
		} else {
			return float64(b.CondF1(ini_entry.GetValue() != nil, func() float64 { return ZendStrtod(ZSTR_VAL(ini_entry.GetValue()), nil) }, 0.0))
		}
	}
	return 0.0
}

/* }}} */

func ZendIniStringEx(name *byte, name_length int, orig int, exists *ZendBool) *byte {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashStrFindPtr(ExecutorGlobals.GetIniDirectives(), name, name_length)
	if ini_entry != nil {
		if exists != nil {
			*exists = 1
		}
		if orig != 0 && ini_entry.GetModified() != 0 {
			if ini_entry.GetOrigValue() != nil {
				return ZSTR_VAL(ini_entry.GetOrigValue())
			} else {
				return nil
			}
		} else {
			if ini_entry.GetValue() != nil {
				return ZSTR_VAL(ini_entry.GetValue())
			} else {
				return nil
			}
		}
	} else {
		if exists != nil {
			*exists = 0
		}
		return nil
	}
}

/* }}} */

func ZendIniString(name string, name_length int, orig int) *byte {
	var exists ZendBool = 1
	var return_value *byte
	return_value = ZendIniStringEx(name, name_length, orig, &exists)
	if exists == 0 {
		return nil
	} else if return_value == nil {
		return_value = ""
	}
	return return_value
}

/* }}} */

func ZendIniGetValue(name *ZendString) *ZendString {
	var ini_entry *ZendIniEntry
	ini_entry = ZendHashFindPtr(ExecutorGlobals.GetIniDirectives(), name)
	if ini_entry != nil {
		if ini_entry.GetValue() != nil {
			return ini_entry.GetValue()
		} else {
			return ZSTR_EMPTY_ALLOC()
		}
	} else {
		return nil
	}
}

/* }}} */

func ZendIniParseBool(str *ZendString) ZendBool {
	if ZSTR_LEN(str) == 4 && strcasecmp(ZSTR_VAL(str), "true") == 0 || ZSTR_LEN(str) == 3 && strcasecmp(ZSTR_VAL(str), "yes") == 0 || ZSTR_LEN(str) == 2 && strcasecmp(ZSTR_VAL(str), "on") == 0 {
		return 1
	} else {
		return atoi(ZSTR_VAL(str)) != 0
	}
}
func ZendIniBooleanDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value int
	var tmp_value *ZendString
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		if ini_entry.GetOrigValue() != nil {
			tmp_value = ini_entry.GetOrigValue()
		} else {
			tmp_value = nil
		}
	} else if ini_entry.GetValue() != nil {
		tmp_value = ini_entry.GetValue()
	} else {
		tmp_value = nil
	}
	if tmp_value != nil {
		value = ZendIniParseBool(tmp_value)
	} else {
		value = 0
	}
	if value != 0 {
		ZEND_PUTS("On")
	} else {
		ZEND_PUTS("Off")
	}
}

/* }}} */

func ZendIniColorDisplayerCb(ini_entry *ZendIniEntry, type_ int) {
	var value *byte
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		value = ZSTR_VAL(ini_entry.GetOrigValue())
	} else if ini_entry.GetValue() != nil {
		value = ZSTR_VAL(ini_entry.GetValue())
	} else {
		value = nil
	}
	if value != nil {
		if ZendUv.GetHtmlErrors() != 0 {
			ZendPrintf("<font style=\"color: %s\">%s</font>", value, value)
		} else {
			ZEND_PUTS(value)
		}
	} else {
		if ZendUv.GetHtmlErrors() != 0 {
			ZEND_PUTS(NO_VALUE_HTML)
		} else {
			ZEND_PUTS(NO_VALUE_PLAINTEXT)
		}
	}
}

/* }}} */

func DisplayLinkNumbers(ini_entry *ZendIniEntry, type_ int) {
	var value *byte
	if type_ == ZEND_INI_DISPLAY_ORIG && ini_entry.GetModified() != 0 {
		value = ZSTR_VAL(ini_entry.GetOrigValue())
	} else if ini_entry.GetValue() != nil {
		value = ZSTR_VAL(ini_entry.GetValue())
	} else {
		value = nil
	}
	if value != nil {
		if atoi(value) == -1 {
			ZEND_PUTS("Unlimited")
		} else {
			ZendPrintf("%s", value)
		}
	}
}

/* }}} */

func OnUpdateBool(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendBool
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendBool)(base + int(mh_arg1))
	*p = ZendIniParseBool(new_value)
	return SUCCESS
}

/* }}} */

func OnUpdateLong(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var base *byte = (*byte)(mh_arg2)
	p = (*ZendLong)(base + int(mh_arg1))
	*p = ZendAtol(ZSTR_VAL(new_value), ZSTR_LEN(new_value))
	return SUCCESS
}

/* }}} */

func OnUpdateLongGEZero(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *ZendLong
	var tmp ZendLong
	var base *byte = (*byte)(mh_arg2)
	tmp = ZendAtol(ZSTR_VAL(new_value), ZSTR_LEN(new_value))
	if tmp < 0 {
		return FAILURE
	}
	p = (*ZendLong)(base + int(mh_arg1))
	*p = tmp
	return SUCCESS
}

/* }}} */

func OnUpdateReal(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p *float64
	var base *byte = (*byte)(mh_arg2)
	p = (*float64)(base + int(mh_arg1))
	*p = ZendStrtod(ZSTR_VAL(new_value), nil)
	return SUCCESS
}

/* }}} */

func OnUpdateString(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = ZSTR_VAL(new_value)
	} else {
		*p = nil
	}
	return SUCCESS
}

/* }}} */

func OnUpdateStringUnempty(entry *ZendIniEntry, new_value *ZendString, mh_arg1 any, mh_arg2 any, mh_arg3 any, stage int) int {
	var p **byte
	var base *byte = (*byte)(mh_arg2)
	if new_value != nil && !(ZSTR_VAL(new_value)[0]) {
		return FAILURE
	}
	p = (**byte)(base + int(mh_arg1))
	if new_value != nil {
		*p = ZSTR_VAL(new_value)
	} else {
		*p = nil
	}
	return SUCCESS
}

/* }}} */
