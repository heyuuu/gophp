// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
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

// #define CONST_CS       ( 1 << 0 )

// #define CONST_PERSISTENT       ( 1 << 1 )

// #define CONST_CT_SUBST       ( 1 << 2 )

// #define CONST_NO_FILE_CACHE       ( 1 << 3 )

// #define PHP_USER_CONSTANT       0x7fffff

/* Flag for zend_get_constant_ex(). Must not class with ZEND_FETCH_CLASS_* flags. */

// #define ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK       0x1000

// @type ZendConstant struct

// #define ZEND_CONSTANT_FLAGS(c) ( Z_CONSTANT_FLAGS ( ( c ) -> value ) & 0xff )

// #define ZEND_CONSTANT_MODULE_NUMBER(c) ( Z_CONSTANT_FLAGS ( ( c ) -> value ) >> 8 )

// #define ZEND_CONSTANT_SET_FLAGS(c,_flags,_module_number) do { Z_CONSTANT_FLAGS ( ( c ) -> value ) = ( ( _flags ) & 0xff ) | ( ( _module_number ) << 8 ) ; } while ( 0 )

// #define REGISTER_NULL_CONSTANT(name,flags) zend_register_null_constant ( ( name ) , sizeof ( name ) - 1 , ( flags ) , module_number )

// #define REGISTER_BOOL_CONSTANT(name,bval,flags) zend_register_bool_constant ( ( name ) , sizeof ( name ) - 1 , ( bval ) , ( flags ) , module_number )

// #define REGISTER_LONG_CONSTANT(name,lval,flags) zend_register_long_constant ( ( name ) , sizeof ( name ) - 1 , ( lval ) , ( flags ) , module_number )

// #define REGISTER_DOUBLE_CONSTANT(name,dval,flags) zend_register_double_constant ( ( name ) , sizeof ( name ) - 1 , ( dval ) , ( flags ) , module_number )

// #define REGISTER_STRING_CONSTANT(name,str,flags) zend_register_string_constant ( ( name ) , sizeof ( name ) - 1 , ( str ) , ( flags ) , module_number )

// #define REGISTER_STRINGL_CONSTANT(name,str,len,flags) zend_register_stringl_constant ( ( name ) , sizeof ( name ) - 1 , ( str ) , ( len ) , ( flags ) , module_number )

// #define REGISTER_NS_NULL_CONSTANT(ns,name,flags) zend_register_null_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( flags ) , module_number )

// #define REGISTER_NS_BOOL_CONSTANT(ns,name,bval,flags) zend_register_bool_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( bval ) , ( flags ) , module_number )

// #define REGISTER_NS_LONG_CONSTANT(ns,name,lval,flags) zend_register_long_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( lval ) , ( flags ) , module_number )

// #define REGISTER_NS_DOUBLE_CONSTANT(ns,name,dval,flags) zend_register_double_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( dval ) , ( flags ) , module_number )

// #define REGISTER_NS_STRING_CONSTANT(ns,name,str,flags) zend_register_string_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( str ) , ( flags ) , module_number )

// #define REGISTER_NS_STRINGL_CONSTANT(ns,name,str,len,flags) zend_register_stringl_constant ( ZEND_NS_NAME ( ns , name ) , sizeof ( ZEND_NS_NAME ( ns , name ) ) - 1 , ( str ) , ( len ) , ( flags ) , module_number )

// #define REGISTER_MAIN_NULL_CONSTANT(name,flags) zend_register_null_constant ( ( name ) , sizeof ( name ) - 1 , ( flags ) , 0 )

// #define REGISTER_MAIN_BOOL_CONSTANT(name,bval,flags) zend_register_bool_constant ( ( name ) , sizeof ( name ) - 1 , ( bval ) , ( flags ) , 0 )

// #define REGISTER_MAIN_LONG_CONSTANT(name,lval,flags) zend_register_long_constant ( ( name ) , sizeof ( name ) - 1 , ( lval ) , ( flags ) , 0 )

// #define REGISTER_MAIN_DOUBLE_CONSTANT(name,dval,flags) zend_register_double_constant ( ( name ) , sizeof ( name ) - 1 , ( dval ) , ( flags ) , 0 )

// #define REGISTER_MAIN_STRING_CONSTANT(name,str,flags) zend_register_string_constant ( ( name ) , sizeof ( name ) - 1 , ( str ) , ( flags ) , 0 )

// #define REGISTER_MAIN_STRINGL_CONSTANT(name,str,len,flags) zend_register_stringl_constant ( ( name ) , sizeof ( name ) - 1 , ( str ) , ( len ) , ( flags ) , 0 )

// #define ZEND_CONSTANT_DTOR       free_zend_constant

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

// #define IS_CONSTANT_VISITED_MARK       0x80

// #define IS_CONSTANT_VISITED(zv) ( Z_ACCESS_FLAGS_P ( zv ) & IS_CONSTANT_VISITED_MARK )

// #define MARK_CONSTANT_VISITED(zv) Z_ACCESS_FLAGS_P ( zv ) |= IS_CONSTANT_VISITED_MARK

// #define RESET_CONSTANT_VISITED(zv) Z_ACCESS_FLAGS_P ( zv ) &= ~ IS_CONSTANT_VISITED_MARK

func FreeZendConstant(zv *Zval) {
	var c *ZendConstant = zv.GetValue().GetPtr()
	if (c.GetValue().GetConstantFlags() & 0xff & 1 << 1) == 0 {
		ZvalPtrDtorNogc(&c.value)
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 0)
		}
		_efree(c)
	} else {
		ZvalInternalPtrDtor(&c.value)
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}
func CleanModuleConstant(el *Zval, arg any) int {
	var c *ZendConstant = (*ZendConstant)(el.GetValue().GetPtr())
	var module_number int = *((*int)(arg))
	if c.GetValue().GetConstantFlags()>>8 == module_number {
		return 1 << 0
	} else {
		return 0
	}
}
func CleanModuleConstants(module_number int) {
	ZendHashApplyWithArgument(EG.GetZendConstants(), CleanModuleConstant, any(&module_number))
}
func ZendStartupConstants() int {
	EG.SetZendConstants((*HashTable)(Malloc(g.SizeOf("HashTable"))))
	_zendHashInit(EG.GetZendConstants(), 128, FreeZendConstant, 1)
	return SUCCESS
}
func ZendRegisterStandardConstants() {
	ZendRegisterLongConstant("E_ERROR", g.SizeOf("\"E_ERROR\"")-1, 1<<0, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_RECOVERABLE_ERROR", g.SizeOf("\"E_RECOVERABLE_ERROR\"")-1, 1<<12, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_WARNING", g.SizeOf("\"E_WARNING\"")-1, 1<<1, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_PARSE", g.SizeOf("\"E_PARSE\"")-1, 1<<2, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_NOTICE", g.SizeOf("\"E_NOTICE\"")-1, 1<<3, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_STRICT", g.SizeOf("\"E_STRICT\"")-1, 1<<11, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_DEPRECATED", g.SizeOf("\"E_DEPRECATED\"")-1, 1<<13, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_CORE_ERROR", g.SizeOf("\"E_CORE_ERROR\"")-1, 1<<4, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_CORE_WARNING", g.SizeOf("\"E_CORE_WARNING\"")-1, 1<<5, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_COMPILE_ERROR", g.SizeOf("\"E_COMPILE_ERROR\"")-1, 1<<6, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_COMPILE_WARNING", g.SizeOf("\"E_COMPILE_WARNING\"")-1, 1<<7, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_USER_ERROR", g.SizeOf("\"E_USER_ERROR\"")-1, 1<<8, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_USER_WARNING", g.SizeOf("\"E_USER_WARNING\"")-1, 1<<9, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_USER_NOTICE", g.SizeOf("\"E_USER_NOTICE\"")-1, 1<<10, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_USER_DEPRECATED", g.SizeOf("\"E_USER_DEPRECATED\"")-1, 1<<14, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("E_ALL", g.SizeOf("\"E_ALL\"")-1, 1<<0|1<<1|1<<2|1<<3|1<<4|1<<5|1<<6|1<<7|1<<8|1<<9|1<<10|1<<12|1<<13|1<<14|1<<11, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("DEBUG_BACKTRACE_PROVIDE_OBJECT", g.SizeOf("\"DEBUG_BACKTRACE_PROVIDE_OBJECT\"")-1, 1<<0, 1<<1|1<<0, 0)
	ZendRegisterLongConstant("DEBUG_BACKTRACE_IGNORE_ARGS", g.SizeOf("\"DEBUG_BACKTRACE_IGNORE_ARGS\"")-1, 1<<1, 1<<1|1<<0, 0)

	/* true/false constants */

	ZendRegisterBoolConstant("TRUE", g.SizeOf("\"TRUE\"")-1, 1, 1<<1|1<<2, 0)
	ZendRegisterBoolConstant("FALSE", g.SizeOf("\"FALSE\"")-1, 0, 1<<1|1<<2, 0)
	ZendRegisterBoolConstant("ZEND_THREAD_SAFE", g.SizeOf("\"ZEND_THREAD_SAFE\"")-1, 0, 1<<1|1<<0, 0)
	ZendRegisterBoolConstant("ZEND_DEBUG_BUILD", g.SizeOf("\"ZEND_DEBUG_BUILD\"")-1, 0, 1<<1|1<<0, 0)
	ZendRegisterNullConstant("NULL", g.SizeOf("\"NULL\"")-1, 1<<1|1<<2, 0)
}
func ZendShutdownConstants() int {
	ZendHashDestroy(EG.GetZendConstants())
	Free(EG.GetZendConstants())
	return SUCCESS
}
func ZendRegisterNullConstant(name string, name_len int, flags int, module_number int) {
	var c ZendConstant
	&c.value.u1.type_info = 1
	&c.GetValue().SetConstantFlags(flags&0xff | module_number<<8)
	c.SetName(ZendStringInitInterned(name, name_len, flags&1<<1))
	ZendRegisterConstant(&c)
}
func ZendRegisterBoolConstant(name string, name_len int, bval ZendBool, flags int, module_number int) {
	var c ZendConstant
	if bval != 0 {
		&c.value.u1.type_info = 3
	} else {
		&c.value.u1.type_info = 2
	}
	&c.GetValue().SetConstantFlags(flags&0xff | module_number<<8)
	c.SetName(ZendStringInitInterned(name, name_len, flags&1<<1))
	ZendRegisterConstant(&c)
}
func ZendRegisterLongConstant(name string, name_len int, lval ZendLong, flags int, module_number int) {
	var c ZendConstant
	var __z *Zval = &c.value
	__z.GetValue().SetLval(lval)
	__z.SetTypeInfo(4)
	&c.GetValue().SetConstantFlags(flags&0xff | module_number<<8)
	c.SetName(ZendStringInitInterned(name, name_len, flags&1<<1))
	ZendRegisterConstant(&c)
}
func ZendRegisterDoubleConstant(name string, name_len int, dval float64, flags int, module_number int) {
	var c ZendConstant
	var __z *Zval = &c.value
	__z.GetValue().SetDval(dval)
	__z.SetTypeInfo(5)
	&c.GetValue().SetConstantFlags(flags&0xff | module_number<<8)
	c.SetName(ZendStringInitInterned(name, name_len, flags&1<<1))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringlConstant(name string, name_len int, strval *byte, strlen int, flags int, module_number int) {
	var c ZendConstant
	var __z *Zval = &c.value
	var __s *ZendString = ZendStringInitInterned(strval, strlen, flags&1<<1)
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	&c.GetValue().SetConstantFlags(flags&0xff | module_number<<8)
	c.SetName(ZendStringInitInterned(name, name_len, flags&1<<1))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringConstant(name string, name_len int, strval *byte, flags int, module_number int) {
	ZendRegisterStringlConstant(name, name_len, strval, strlen(strval), flags, module_number)
}
func ZendGetSpecialConstant(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	var haltoff []byte = "__COMPILER_HALT_OFFSET__"
	if EG.GetCurrentExecuteData() == nil {
		return nil
	} else if name_len == g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(name, "__COMPILER_HALT_OFFSET__", g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) {
		var cfilename *byte
		var haltname *ZendString
		var clen int
		cfilename = ZendGetExecutedFilename()
		clen = strlen(cfilename)

		/* check for __COMPILER_HALT_OFFSET__ */

		haltname = ZendManglePropertyName(haltoff, g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1, cfilename, clen, 0)
		c = ZendHashFindPtr(EG.GetZendConstants(), haltname)
		ZendStringEfree(haltname)
		return c
	} else {
		return nil
	}
}
func ZendVerifyConstAccess(c *ZendClassConstant, scope *ZendClassEntry) int {
	if (c.GetValue().GetAccessFlags() & 1 << 0) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & 1 << 2) != 0 {
		return c.GetCe() == scope
	} else {
		assert((c.GetValue().GetAccessFlags() & 1 << 1) != 0)
		return ZendCheckProtected(c.GetCe(), scope)
	}
}

/* }}} */

func ZendGetConstantStrImpl(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	if g.Assign(&c, ZendHashStrFindPtr(EG.GetZendConstants(), name, name_len)) == nil {
		var lcname *byte = _emalloc(name_len + 1)
		ZendStrTolowerCopy(lcname, name, name_len)
		if g.Assign(&c, ZendHashStrFindPtr(EG.GetZendConstants(), lcname, name_len)) != nil {
			if (c.GetValue().GetConstantFlags() & 0xff & 1 << 0) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(name, name_len)
		}
		_efree(lcname)
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
	zv = ZendHashFind(EG.GetZendConstants(), name)
	if zv == nil {
		var lcname *byte = _emalloc(name.GetLen() + 1)
		ZendStrTolowerCopy(lcname, name.GetVal(), name.GetLen())
		zv = ZendHashStrFind(EG.GetZendConstants(), lcname, name.GetLen())
		if zv != nil {
			c = zv.GetValue().GetPtr()
			if (c.GetValue().GetConstantFlags() & 0xff & 1 << 0) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(name.GetVal(), name.GetLen())
		}
		_efree(lcname)
		return c
	} else {
		return (*ZendConstant)(zv.GetValue().GetPtr())
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
	var ns_sep *byte = ZendMemrchr(c.GetName().GetVal(), '\\', c.GetName().GetLen())
	if ns_sep != nil {

		/* Namespaces are always case-insensitive. Only compare shortname. */

		var shortname_offset int = ns_sep - c.GetName().GetVal() + 1
		var shortname_len int = c.GetName().GetLen() - shortname_offset
		return memcmp(access_name+shortname_offset, c.GetName().GetVal()+shortname_offset, shortname_len) != 0
	} else {

		/* No namespace, compare whole name */

		return memcmp(access_name, c.GetName().GetVal(), c.GetName().GetLen()) != 0

		/* No namespace, compare whole name */

	}
}
func ZendGetConstantEx(cname *ZendString, scope *ZendClassEntry, flags uint32) *Zval {
	var c *ZendConstant
	var colon *byte
	var ce *ZendClassEntry = nil
	var name *byte = cname.GetVal()
	var name_len int = cname.GetLen()

	/* Skip leading \\ */

	if name[0] == '\\' {
		name += 1
		name_len -= 1
		cname = nil
	}
	if g.Assign(&colon, ZendMemrchr(name, ':', name_len)) && colon > name && (*(colon - 1)) == ':' {
		var class_name_len int = colon - name - 1
		var const_name_len int = name_len - class_name_len - 2
		var constant_name *ZendString = ZendStringInit(colon+1, const_name_len, 0)
		var class_name *ZendString = ZendStringInit(name, class_name_len, 0)
		var c *ZendClassConstant = nil
		var ret_constant *Zval = nil
		if class_name.GetLen() == g.SizeOf("\"self\"")-1 && ZendBinaryStrcasecmp(class_name.GetVal(), class_name.GetLen(), "self", g.SizeOf("\"self\"")-1) == 0 {
			if scope == nil {
				ZendThrowError(nil, "Cannot access self:: when no class scope is active")
				goto failure
			}
			ce = scope
		} else if class_name.GetLen() == g.SizeOf("\"parent\"")-1 && ZendBinaryStrcasecmp(class_name.GetVal(), class_name.GetLen(), "parent", g.SizeOf("\"parent\"")-1) == 0 {
			if scope == nil {
				ZendThrowError(nil, "Cannot access parent:: when no class scope is active")
				goto failure
			} else if !(scope.parent) {
				ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
				goto failure
			} else {
				ce = scope.parent
			}
		} else if class_name.GetLen() == g.SizeOf("\"static\"")-1 && ZendBinaryStrcasecmp(class_name.GetVal(), class_name.GetLen(), "static", g.SizeOf("\"static\"")-1) == 0 {
			ce = ZendGetCalledScope(EG.GetCurrentExecuteData())
			if ce == nil {
				ZendThrowError(nil, "Cannot access static:: when no class scope is active")
				goto failure
			}
		} else {
			ce = ZendFetchClass(class_name, flags)
		}
		if ce != nil {
			c = ZendHashFindPtr(&ce.constants_table, constant_name)
			if c == nil {
				if (flags & 0x100) == 0 {
					ZendThrowError(nil, "Undefined class constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
					goto failure
				}
				ret_constant = nil
			} else {
				if ZendVerifyConstAccess(c, scope) == 0 {
					if (flags & 0x100) == 0 {
						ZendThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), class_name.GetVal(), constant_name.GetVal())
					}
					goto failure
				}
				ret_constant = &c.value
			}
		}
		if ret_constant != nil && ret_constant.GetType() == 11 {
			var ret int
			if (ret_constant.GetAccessFlags() & 0x80) != 0 {
				ZendThrowError(nil, "Cannot declare self-referencing constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
				ret_constant = nil
				goto failure
			}
			ret_constant.SetAccessFlags(ret_constant.GetAccessFlags() | 0x80)
			ret = ZvalUpdateConstantEx(ret_constant, c.GetCe())
			ret_constant.SetAccessFlags(ret_constant.GetAccessFlags() &^ 0x80)
			if ret != SUCCESS {
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

	if g.Assign(&colon, ZendMemrchr(name, '\\', name_len)) != nil {

		/* compound constant name */

		var prefix_len int = colon - name
		var const_name_len int = name_len - prefix_len - 1
		var constant_name *byte = colon + 1
		var lcname *byte
		var lcname_len int
		lcname_len = prefix_len + 1 + const_name_len
		lcname = _emalloc(lcname_len + 1)
		ZendStrTolowerCopy(lcname, name, prefix_len)

		/* Check for namespace constant */

		lcname[prefix_len] = '\\'
		memcpy(lcname+prefix_len+1, constant_name, const_name_len+1)
		if g.Assign(&c, ZendHashStrFindPtr(EG.GetZendConstants(), lcname, lcname_len)) == nil {

			/* try lowercase */

			ZendStrTolower(lcname+prefix_len+1, const_name_len)
			if g.Assign(&c, ZendHashStrFindPtr(EG.GetZendConstants(), lcname, lcname_len)) != nil {
				if (c.GetValue().GetConstantFlags() & 0xff & 1 << 0) != 0 {
					c = nil
				}
			}
		}
		_efree(lcname)
		if c == nil {
			if (flags & 0x10) == 0 {
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
	if (flags & 0x1000) == 0 {
		if (c.GetValue().GetConstantFlags()&0xff&(1<<0|1<<2)) == 0 && IsAccessDeprecated(c, name) != 0 {
			ZendError(1<<13, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
		}
	}
	return &c.value
}
func ZendHashAddConstant(ht *HashTable, key *ZendString, c *ZendConstant) any {
	var ret any
	var copy *ZendConstant = g.CondF((c.GetValue().GetConstantFlags()&0xff&1<<1) != 0, func() any { return __zendMalloc(g.SizeOf("zend_constant")) }, func() any { return _emalloc(g.SizeOf("zend_constant")) })
	memcpy(copy, c, g.SizeOf("zend_constant"))
	ret = ZendHashAddPtr(ht, key, copy)
	if !ret {
		g.CondF((c.GetValue().GetConstantFlags()&0xff&1<<1) != 0, func() { return Free(copy) }, func() { return _efree(copy) })
	}
	return ret
}
func ZendRegisterConstant(c *ZendConstant) int {
	var lowercase_name *ZendString = nil
	var name *ZendString
	var ret int = SUCCESS
	if (c.GetValue().GetConstantFlags() & 0xff & 1 << 0) == 0 {
		lowercase_name = ZendStringTolowerEx(c.GetName(), c.GetValue().GetConstantFlags()&0xff&1<<1)
		lowercase_name = ZendNewInternedString(lowercase_name)
		name = lowercase_name
	} else {
		var slash *byte = strrchr(c.GetName().GetVal(), '\\')
		if slash != nil {
			lowercase_name = ZendStringInit(c.GetName().GetVal(), c.GetName().GetLen(), c.GetValue().GetConstantFlags()&0xff&1<<1)
			ZendStrTolower(lowercase_name.GetVal(), slash-c.GetName().GetVal())
			lowercase_name = ZendNewInternedString(lowercase_name)
			name = lowercase_name
		} else {
			name = c.GetName()
		}
	}

	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */

	if name.GetLen() == g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(name.GetVal(), "__COMPILER_HALT_OFFSET__", g.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) || ZendHashAddConstant(EG.GetZendConstants(), name, c) == nil {

		/* The internal __COMPILER_HALT_OFFSET__ is prefixed by NULL byte */

		if c.GetName().GetVal()[0] == '0' && c.GetName().GetLen() > g.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")-1 && memcmp(name.GetVal(), "0__COMPILER_HALT_OFFSET__", g.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")) == 0 {

		}
		ZendError(1<<3, "Constant %s already defined", name.GetVal())
		ZendStringRelease(c.GetName())
		if (c.GetValue().GetConstantFlags() & 0xff & 1 << 1) == 0 {
			ZvalPtrDtorNogc(&c.value)
		}
		ret = FAILURE
	}
	if lowercase_name != nil {
		ZendStringRelease(lowercase_name)
	}
	return ret
}
