// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

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

// #define ZEND_CONSTANTS_H

// # include "zend_globals.h"

const CONST_CS = 1 << 0
const CONST_PERSISTENT = 1 << 1
const CONST_CT_SUBST = 1 << 2
const CONST_NO_FILE_CACHE = 1 << 3
const PHP_USER_CONSTANT = 0x7fffff

/* Flag for zend_get_constant_ex(). Must not class with ZEND_FETCH_CLASS_* flags. */

const ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK = 0x1000

func ZEND_CONSTANT_FLAGS(c *ZendConstant) int {
	return Z_CONSTANT_FLAGS(c.GetValue()) & 0xff
}
func ZEND_CONSTANT_MODULE_NUMBER(c *ZendConstant) int {
	return Z_CONSTANT_FLAGS(c.GetValue()) >> 8
}
func ZEND_CONSTANT_SET_FLAGS(c *ZendConstant, _flags int, _module_number int) {
	Z_CONSTANT_FLAGS(c.GetValue()) = _flags&0xff | _module_number<<8
}
func REGISTER_NULL_CONSTANT(name *byte, flags int) {
	ZendRegisterNullConstant(name, b.SizeOf("name")-1, flags, module_number)
}
func REGISTER_BOOL_CONSTANT(name *byte, bval ZendBool, flags int) {
	ZendRegisterBoolConstant(name, b.SizeOf("name")-1, bval, flags, module_number)
}
func REGISTER_LONG_CONSTANT(name string, lval ZendLong, flags int) {
	ZendRegisterLongConstant(name, b.SizeOf("name")-1, lval, flags, module_number)
}
func REGISTER_DOUBLE_CONSTANT(name string, dval float64, flags int) {
	ZendRegisterDoubleConstant(name, b.SizeOf("name")-1, dval, flags, module_number)
}
func REGISTER_STRING_CONSTANT(name string, str *byte, flags int) {
	ZendRegisterStringConstant(name, b.SizeOf("name")-1, str, flags, module_number)
}
func REGISTER_STRINGL_CONSTANT(name *byte, str *byte, len_ int, flags int) {
	ZendRegisterStringlConstant(name, b.SizeOf("name")-1, str, len_, flags, module_number)
}
func REGISTER_NS_NULL_CONSTANT(ns string, name string, flags int) {
	ZendRegisterNullConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, flags, module_number)
}
func REGISTER_NS_BOOL_CONSTANT(ns string, name string, bval ZendBool, flags int) {
	ZendRegisterBoolConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, bval, flags, module_number)
}
func REGISTER_NS_LONG_CONSTANT(ns string, name string, lval ZendLong, flags int) {
	ZendRegisterLongConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, lval, flags, module_number)
}
func REGISTER_NS_DOUBLE_CONSTANT(ns string, name string, dval float64, flags int) {
	ZendRegisterDoubleConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, dval, flags, module_number)
}
func REGISTER_NS_STRING_CONSTANT(ns string, name string, str *byte, flags int) {
	ZendRegisterStringConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, str, flags, module_number)
}
func REGISTER_NS_STRINGL_CONSTANT(ns string, name string, str *byte, len_ int, flags int) {
	ZendRegisterStringlConstant(ZEND_NS_NAME(ns, name), b.SizeOf("ZEND_NS_NAME ( ns , name )")-1, str, len_, flags, module_number)
}
func REGISTER_MAIN_NULL_CONSTANT(name string, flags int) {
	ZendRegisterNullConstant(name, b.SizeOf("name")-1, flags, 0)
}
func REGISTER_MAIN_BOOL_CONSTANT(name string, bval ZendBool, flags int) {
	ZendRegisterBoolConstant(name, b.SizeOf("name")-1, bval, flags, 0)
}
func REGISTER_MAIN_LONG_CONSTANT(name string, lval ZendLong, flags int) {
	ZendRegisterLongConstant(name, b.SizeOf("name")-1, lval, flags, 0)
}
func REGISTER_MAIN_DOUBLE_CONSTANT(name string, dval float64, flags int) {
	ZendRegisterDoubleConstant(name, b.SizeOf("name")-1, dval, flags, 0)
}
func REGISTER_MAIN_STRING_CONSTANT(name *byte, str *byte, flags int) {
	ZendRegisterStringConstant(name, b.SizeOf("name")-1, str, flags, 0)
}
func REGISTER_MAIN_STRINGL_CONSTANT(name string, str *byte, len_ int, flags int) {
	ZendRegisterStringlConstant(name, b.SizeOf("name")-1, str, len_, flags, 0)
}

const ZEND_CONSTANT_DTOR DtorFuncT = FreeZendConstant

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

// # include "zend.h"

// # include "zend_constants.h"

// # include "zend_exceptions.h"

// # include "zend_execute.h"

// # include "zend_variables.h"

// # include "zend_operators.h"

// # include "zend_globals.h"

// # include "zend_API.h"

/* Protection from recursive self-referencing class constants */

const IS_CONSTANT_VISITED_MARK = 0x80

func IS_CONSTANT_VISITED(zv *Zval) int {
	return Z_ACCESS_FLAGS_P(zv) & IS_CONSTANT_VISITED_MARK
}
func MARK_CONSTANT_VISITED(zv *Zval) uint32 {
	Z_ACCESS_FLAGS_P(zv) |= IS_CONSTANT_VISITED_MARK
	return Z_ACCESS_FLAGS_P(zv)
}
func RESET_CONSTANT_VISITED(zv *Zval) uint32 {
	Z_ACCESS_FLAGS_P(zv) &= ^IS_CONSTANT_VISITED_MARK
	return Z_ACCESS_FLAGS_P(zv)
}
func FreeZendConstant(zv *Zval) {
	var c *ZendConstant = Z_PTR_P(zv)
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
		ZvalPtrDtorNogc(&c.value)
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 0)
		}
		Efree(c)
	} else {
		ZvalInternalPtrDtor(&c.value)
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}
func CleanModuleConstant(el *Zval, arg any) int {
	var c *ZendConstant = (*ZendConstant)(Z_PTR_P(el))
	var module_number int = *((*int)(arg))
	if ZEND_CONSTANT_MODULE_NUMBER(c) == module_number {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}
func CleanModuleConstants(module_number int) {
	ZendHashApplyWithArgument(ExecutorGlobals.GetZendConstants(), CleanModuleConstant, any(&module_number))
}
func ZendStartupConstants() int {
	ExecutorGlobals.SetZendConstants((*HashTable)(Malloc(b.SizeOf("HashTable"))))
	ZendHashInit(ExecutorGlobals.GetZendConstants(), 128, nil, ZEND_CONSTANT_DTOR, 1)
	return SUCCESS
}
func ZendRegisterStandardConstants() {
	REGISTER_MAIN_LONG_CONSTANT("E_ERROR", E_ERROR, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_RECOVERABLE_ERROR", E_RECOVERABLE_ERROR, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_WARNING", E_WARNING, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_PARSE", E_PARSE, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_NOTICE", E_NOTICE, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_STRICT", E_STRICT, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_DEPRECATED", E_DEPRECATED, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_CORE_ERROR", E_CORE_ERROR, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_CORE_WARNING", E_CORE_WARNING, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_COMPILE_ERROR", E_COMPILE_ERROR, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_COMPILE_WARNING", E_COMPILE_WARNING, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_USER_ERROR", E_USER_ERROR, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_USER_WARNING", E_USER_WARNING, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_USER_NOTICE", E_USER_NOTICE, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_USER_DEPRECATED", E_USER_DEPRECATED, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("E_ALL", E_ALL, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("DEBUG_BACKTRACE_PROVIDE_OBJECT", DEBUG_BACKTRACE_PROVIDE_OBJECT, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_LONG_CONSTANT("DEBUG_BACKTRACE_IGNORE_ARGS", DEBUG_BACKTRACE_IGNORE_ARGS, CONST_PERSISTENT|CONST_CS)

	/* true/false constants */

	REGISTER_MAIN_BOOL_CONSTANT("TRUE", 1, CONST_PERSISTENT|CONST_CT_SUBST)
	REGISTER_MAIN_BOOL_CONSTANT("FALSE", 0, CONST_PERSISTENT|CONST_CT_SUBST)
	REGISTER_MAIN_BOOL_CONSTANT("ZEND_THREAD_SAFE", ZTS_V, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_BOOL_CONSTANT("ZEND_DEBUG_BUILD", core.ZEND_DEBUG, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_NULL_CONSTANT("NULL", CONST_PERSISTENT|CONST_CT_SUBST)
}
func ZendShutdownConstants() int {
	ZendHashDestroy(ExecutorGlobals.GetZendConstants())
	Free(ExecutorGlobals.GetZendConstants())
	return SUCCESS
}
func ZendRegisterNullConstant(name *byte, name_len int, flags int, module_number int) {
	var c ZendConstant
	ZVAL_NULL(&c.value)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterBoolConstant(name *byte, name_len int, bval ZendBool, flags int, module_number int) {
	var c ZendConstant
	ZVAL_BOOL(&c.value, bval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterLongConstant(name *byte, name_len int, lval ZendLong, flags int, module_number int) {
	var c ZendConstant
	ZVAL_LONG(&c.value, lval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterDoubleConstant(name *byte, name_len int, dval float64, flags int, module_number int) {
	var c ZendConstant
	ZVAL_DOUBLE(&c.value, dval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringlConstant(name *byte, name_len int, strval *byte, strlen int, flags int, module_number int) {
	var c ZendConstant
	ZVAL_STR(&c.value, ZendStringInitInterned(strval, strlen, flags&CONST_PERSISTENT))
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringConstant(name *byte, name_len int, strval *byte, flags int, module_number int) {
	ZendRegisterStringlConstant(name, name_len, strval, strlen(strval), flags, module_number)
}
func ZendGetSpecialConstant(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	var haltoff []byte = "__COMPILER_HALT_OFFSET__"
	if ExecutorGlobals.GetCurrentExecuteData() == nil {
		return nil
	} else if name_len == b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(name, "__COMPILER_HALT_OFFSET__", b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) {
		var cfilename *byte
		var haltname *ZendString
		var clen int
		cfilename = ZendGetExecutedFilename()
		clen = strlen(cfilename)

		/* check for __COMPILER_HALT_OFFSET__ */

		haltname = ZendManglePropertyName(haltoff, b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1, cfilename, clen, 0)
		c = ZendHashFindPtr(ExecutorGlobals.GetZendConstants(), haltname)
		ZendStringEfree(haltname)
		return c
	} else {
		return nil
	}
}
func ZendVerifyConstAccess(c *ZendClassConstant, scope *ZendClassEntry) int {
	if (Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PUBLIC) != 0 {
		return 1
	} else if (Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PRIVATE) != 0 {
		return c.GetCe() == scope
	} else {
		ZEND_ASSERT((Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PROTECTED) != 0)
		return ZendCheckProtected(c.GetCe(), scope)
	}
}

/* }}} */

func ZendGetConstantStrImpl(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	if b.Assign(&c, ZendHashStrFindPtr(ExecutorGlobals.GetZendConstants(), name, name_len)) == nil {
		var lcname *byte = DoAlloca(name_len+1, use_heap)
		ZendStrTolowerCopy(lcname, name, name_len)
		if b.Assign(&c, ZendHashStrFindPtr(ExecutorGlobals.GetZendConstants(), lcname, name_len)) != nil {
			if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(name, name_len)
		}
		FreeAlloca(lcname, use_heap)
	}
	return c
}
func ZendGetConstantStr(name *byte, name_len int) *Zval {
	var c *ZendConstant = ZendGetConstantStrImpl(name, name_len)
	if c != nil {
		return &c.value
	} else {
		return nil
	}
}
func ZendGetConstantImpl(name *ZendString) *ZendConstant {
	var zv *Zval
	var c *ZendConstant
	zv = ZendHashFind(ExecutorGlobals.GetZendConstants(), name)
	if zv == nil {
		var lcname *byte = DoAlloca(ZSTR_LEN(name)+1, use_heap)
		ZendStrTolowerCopy(lcname, ZSTR_VAL(name), ZSTR_LEN(name))
		zv = ZendHashStrFind(ExecutorGlobals.GetZendConstants(), lcname, ZSTR_LEN(name))
		if zv != nil {
			c = Z_PTR_P(zv)
			if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(ZSTR_VAL(name), ZSTR_LEN(name))
		}
		FreeAlloca(lcname, use_heap)
		return c
	} else {
		return (*ZendConstant)(Z_PTR_P(zv))
	}
}
func ZendGetConstant(name *ZendString) *Zval {
	var c *ZendConstant = ZendGetConstantImpl(name)
	if c != nil {
		return &c.value
	} else {
		return nil
	}
}
func IsAccessDeprecated(c *ZendConstant, access_name *byte) ZendBool {
	var ns_sep *byte = ZendMemrchr(ZSTR_VAL(c.GetName()), '\\', ZSTR_LEN(c.GetName()))
	if ns_sep != nil {

		/* Namespaces are always case-insensitive. Only compare shortname. */

		var shortname_offset int = ns_sep - ZSTR_VAL(c.GetName()) + 1
		var shortname_len int = ZSTR_LEN(c.GetName()) - shortname_offset
		return memcmp(access_name+shortname_offset, ZSTR_VAL(c.GetName())+shortname_offset, shortname_len) != 0
	} else {

		/* No namespace, compare whole name */

		return memcmp(access_name, ZSTR_VAL(c.GetName()), ZSTR_LEN(c.GetName())) != 0

		/* No namespace, compare whole name */

	}
}
func ZendGetConstantEx(cname *ZendString, scope *ZendClassEntry, flags uint32) *Zval {
	var c *ZendConstant
	var colon *byte
	var ce *ZendClassEntry = nil
	var name *byte = ZSTR_VAL(cname)
	var name_len int = ZSTR_LEN(cname)

	/* Skip leading \\ */

	if name[0] == '\\' {
		name += 1
		name_len -= 1
		cname = nil
	}
	if b.Assign(&colon, ZendMemrchr(name, ':', name_len)) && colon > name && (*(colon - 1)) == ':' {
		var class_name_len int = colon - name - 1
		var const_name_len int = name_len - class_name_len - 2
		var constant_name *ZendString = ZendStringInit(colon+1, const_name_len, 0)
		var class_name *ZendString = ZendStringInit(name, class_name_len, 0)
		var c *ZendClassConstant = nil
		var ret_constant *Zval = nil
		if ZendStringEqualsLiteralCi(class_name, "self") {
			if UNEXPECTED(scope == nil) {
				ZendThrowError(nil, "Cannot access self:: when no class scope is active")
				goto failure
			}
			ce = scope
		} else if ZendStringEqualsLiteralCi(class_name, "parent") {
			if UNEXPECTED(scope == nil) {
				ZendThrowError(nil, "Cannot access parent:: when no class scope is active")
				goto failure
			} else if UNEXPECTED(!(scope.parent)) {
				ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
				goto failure
			} else {
				ce = scope.parent
			}
		} else if ZendStringEqualsLiteralCi(class_name, "static") {
			ce = ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData())
			if UNEXPECTED(ce == nil) {
				ZendThrowError(nil, "Cannot access static:: when no class scope is active")
				goto failure
			}
		} else {
			ce = ZendFetchClass(class_name, flags)
		}
		if ce != nil {
			c = ZendHashFindPtr(&ce.constants_table, constant_name)
			if c == nil {
				if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
					ZendThrowError(nil, "Undefined class constant '%s::%s'", ZSTR_VAL(class_name), ZSTR_VAL(constant_name))
					goto failure
				}
				ret_constant = nil
			} else {
				if ZendVerifyConstAccess(c, scope) == 0 {
					if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
						ZendThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(Z_ACCESS_FLAGS(c.GetValue())), ZSTR_VAL(class_name), ZSTR_VAL(constant_name))
					}
					goto failure
				}
				ret_constant = &c.value
			}
		}
		if ret_constant != nil && Z_TYPE_P(ret_constant) == IS_CONSTANT_AST {
			var ret int
			if IS_CONSTANT_VISITED(ret_constant) != 0 {
				ZendThrowError(nil, "Cannot declare self-referencing constant '%s::%s'", ZSTR_VAL(class_name), ZSTR_VAL(constant_name))
				ret_constant = nil
				goto failure
			}
			MARK_CONSTANT_VISITED(ret_constant)
			ret = ZvalUpdateConstantEx(ret_constant, c.GetCe())
			RESET_CONSTANT_VISITED(ret_constant)
			if UNEXPECTED(ret != SUCCESS) {
				ret_constant = nil
				goto failure
			}
		}
	failure:
		ZendStringReleaseEx(class_name, 0)
		ZendStringEfree(constant_name)
		return ret_constant
	}

	/* non-class constant */

	if b.Assign(&colon, ZendMemrchr(name, '\\', name_len)) != nil {

		/* compound constant name */

		var prefix_len int = colon - name
		var const_name_len int = name_len - prefix_len - 1
		var constant_name *byte = colon + 1
		var lcname *byte
		var lcname_len int
		lcname_len = prefix_len + 1 + const_name_len
		lcname = DoAlloca(lcname_len+1, use_heap)
		ZendStrTolowerCopy(lcname, name, prefix_len)

		/* Check for namespace constant */

		lcname[prefix_len] = '\\'
		memcpy(lcname+prefix_len+1, constant_name, const_name_len+1)
		if b.Assign(&c, ZendHashStrFindPtr(ExecutorGlobals.GetZendConstants(), lcname, lcname_len)) == nil {

			/* try lowercase */

			ZendStrTolower(lcname+prefix_len+1, const_name_len)
			if b.Assign(&c, ZendHashStrFindPtr(ExecutorGlobals.GetZendConstants(), lcname, lcname_len)) != nil {
				if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
					c = nil
				}
			}
		}
		FreeAlloca(lcname, use_heap)
		if c == nil {
			if (flags & IS_CONSTANT_UNQUALIFIED) == 0 {
				return nil
			}

			/* name requires runtime resolution, need to check non-namespaced name */

			c = ZendGetConstantStrImpl(constant_name, const_name_len)
			name = constant_name
		}
	} else {
		if cname != nil {
			c = ZendGetConstantImpl(cname)
		} else {
			c = ZendGetConstantStrImpl(name, name_len)
		}
	}
	if c == nil {
		return nil
	}
	if (flags & ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) == 0 {
		if (ZEND_CONSTANT_FLAGS(c)&(CONST_CS|CONST_CT_SUBST)) == 0 && IsAccessDeprecated(c, name) != 0 {
			ZendError(E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", ZSTR_VAL(c.GetName()))
		}
	}
	return &c.value
}
func ZendHashAddConstant(ht *HashTable, key *ZendString, c *ZendConstant) any {
	var ret any
	var copy *ZendConstant = Pemalloc(b.SizeOf("zend_constant"), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
	memcpy(copy, c, b.SizeOf("zend_constant"))
	ret = ZendHashAddPtr(ht, key, copy)
	if !ret {
		Pefree(copy, ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
	}
	return ret
}
func ZendRegisterConstant(c *ZendConstant) int {
	var lowercase_name *ZendString = nil
	var name *ZendString
	var ret int = SUCCESS
	if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) == 0 {
		lowercase_name = ZendStringTolowerEx(c.GetName(), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
		lowercase_name = ZendNewInternedString(lowercase_name)
		name = lowercase_name
	} else {
		var slash *byte = strrchr(ZSTR_VAL(c.GetName()), '\\')
		if slash != nil {
			lowercase_name = ZendStringInit(ZSTR_VAL(c.GetName()), ZSTR_LEN(c.GetName()), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
			ZendStrTolower(ZSTR_VAL(lowercase_name), slash-ZSTR_VAL(c.GetName()))
			lowercase_name = ZendNewInternedString(lowercase_name)
			name = lowercase_name
		} else {
			name = c.GetName()
		}
	}

	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */

	if ZendStringEqualsLiteral(name, "__COMPILER_HALT_OFFSET__") || ZendHashAddConstant(ExecutorGlobals.GetZendConstants(), name, c) == nil {

		/* The internal __COMPILER_HALT_OFFSET__ is prefixed by NULL byte */

		if ZSTR_VAL(c.GetName())[0] == '0' && ZSTR_LEN(c.GetName()) > b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")-1 && memcmp(ZSTR_VAL(name), "0__COMPILER_HALT_OFFSET__", b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")) == 0 {

		}
		ZendError(E_NOTICE, "Constant %s already defined", ZSTR_VAL(name))
		ZendStringRelease(c.GetName())
		if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
			ZvalPtrDtorNogc(&c.value)
		}
		ret = FAILURE
	}
	if lowercase_name != nil {
		ZendStringRelease(lowercase_name)
	}
	return ret
}
