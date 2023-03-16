// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZEND_CONSTANT_FLAGS(c *ZendConstant) int {
	return c.GetValue().GetConstantFlags() & 0xff
}
func ZEND_CONSTANT_MODULE_NUMBER(c *ZendConstant) int {
	return c.GetValue().GetConstantFlags() >> 8
}
func ZEND_CONSTANT_SET_FLAGS(c *ZendConstant, _flags int, _module_number int) {
	c.GetValue().GetConstantFlags() = _flags&0xff | _module_number<<8
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
func REGISTER_MAIN_STRINGL_CONSTANT(name string, str string, flags int) {
	ZendRegisterStringlConstant(name, str, flags, 0)
}
func IS_CONSTANT_VISITED(zv *Zval) int {
	return zv.GetAccessFlags() & IS_CONSTANT_VISITED_MARK
}
func MARK_CONSTANT_VISITED(zv *Zval) uint32 {
	zv.AddAccessFlags(IS_CONSTANT_VISITED_MARK)
	return zv.GetAccessFlags()
}
func RESET_CONSTANT_VISITED(zv *Zval) uint32 {
	zv.SubAccessFlags(IS_CONSTANT_VISITED_MARK)
	return zv.GetAccessFlags()
}
func FreeZendConstant(zv *Zval) {
	var c *ZendConstant = zv.GetPtr()
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
		ZvalPtrDtorNogc(c.GetValue())
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 0)
		}
		Efree(c)
	} else {
		ZvalInternalPtrDtor(c.GetValue())
		if c.GetName() != nil {
			ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}
func CleanModuleConstant(el *Zval, arg any) int {
	var c *ZendConstant = (*ZendConstant)(el.GetPtr())
	var module_number int = *((*int)(arg))
	if ZEND_CONSTANT_MODULE_NUMBER(c) == module_number {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}
func CleanModuleConstants(module_number int) {
	ZendHashApplyWithArgument(EG__().GetZendConstants(), CleanModuleConstant, any(&module_number))
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
	REGISTER_MAIN_BOOL_CONSTANT("ZEND_DEBUG_BUILD", 0, CONST_PERSISTENT|CONST_CS)
	REGISTER_MAIN_NULL_CONSTANT("NULL", CONST_PERSISTENT|CONST_CT_SUBST)
}
func ZendRegisterNullConstant(name *byte, name_len int, flags int, module_number int) {
	var c ZendConstant
	c.GetValue().SetNull()
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterBoolConstant(name *byte, name_len int, bval ZendBool, flags int, module_number int) {
	var c ZendConstant
	ZVAL_BOOL(c.GetValue(), bval != 0)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterLongConstant(name *byte, name_len int, lval ZendLong, flags int, module_number int) {
	var c ZendConstant
	c.GetValue().SetLong(lval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterDoubleConstant(name *byte, name_len int, dval float64, flags int, module_number int) {
	var c ZendConstant
	c.GetValue().SetDouble(dval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringlConstant(name string, str string, flags int, module_number int) {
	c := NewZendConstant(name)
	c.Value().SetRawString(str)

	ZEND_CONSTANT_SET_FLAGS(c, flags, module_number)
	ZendRegisterConstant(c)
}
func ZendRegisterStringConstant(name *byte, name_len int, strval *byte, flags int, module_number int) {
	ZendRegisterStringlConstant(name, strval, flags, module_number)
}
func ZendGetSpecialConstant(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	var haltoff = "__COMPILER_HALT_OFFSET__"
	if CurrEX() == nil {
		return nil
	} else if name_len == b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(name, "__COMPILER_HALT_OFFSET__", b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) {
		cfilename := ZendGetExecutedFilename()

		/* check for __COMPILER_HALT_OFFSET__ */

		haltname := ZendManglePropertyName_Ex(haltoff, cfilename)
		c = ZendHashFindPtr(EG__().GetZendConstants(), haltname)
		return c
	} else {
		return nil
	}
}
func ZendVerifyConstAccess(c *ZendClassConstant, scope *ZendClassEntry) int {
	if (c.GetValue().GetAccessFlags() & ZEND_ACC_PUBLIC) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & ZEND_ACC_PRIVATE) != 0 {
		return c.GetCe() == scope
	} else {
		ZEND_ASSERT((c.GetValue().GetAccessFlags() & ZEND_ACC_PROTECTED) != 0)
		return ZendCheckProtected(c.GetCe(), scope)
	}
}
func ZendGetConstantStrImpl(name *byte, name_len int) *ZendConstant {
	var c *ZendConstant
	if b.Assign(&c, ZendHashStrFindPtr(EG__().GetZendConstants(), name, name_len)) == nil {
		var lcname *byte = DoAlloca(name_len+1, use_heap)
		ZendStrTolowerCopy(lcname, name, name_len)
		if b.Assign(&c, ZendHashStrFindPtr(EG__().GetZendConstants(), lcname, name_len)) != nil {
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
func ZendGetConstantImpl(name *ZendString) *ZendConstant {
	var zv *Zval
	var c *ZendConstant
	zv = EG__().GetZendConstants().KeyFind(name.GetStr())
	if zv == nil {
		var lcname *byte = DoAlloca(name.GetLen()+1, use_heap)
		ZendStrTolowerCopy(lcname, name.GetVal(), name.GetLen())
		zv = EG__().GetZendConstants().KeyFind(b.CastStr(lcname, name.GetLen()))
		if zv != nil {
			c = zv.GetPtr()
			if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(name.GetVal(), name.GetLen())
		}
		FreeAlloca(lcname, use_heap)
		return c
	} else {
		return (*ZendConstant)(zv.GetPtr())
	}
}
func ZendGetConstant(name *ZendString) *Zval {
	var c *ZendConstant = ZendGetConstantImpl(name)
	if c != nil {
		return c.GetValue()
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
	if b.Assign(&colon, ZendMemrchr(name, ':', name_len)) && colon > name && (*(colon - 1)) == ':' {
		var class_name_len int = colon - name - 1
		var const_name_len int = name_len - class_name_len - 2
		var constant_name *ZendString = ZendStringInit(colon+1, const_name_len, 0)
		var class_name *ZendString = ZendStringInit(name, class_name_len, 0)
		var c *ZendClassConstant = nil
		var ret_constant *Zval = nil
		if ZendStringEqualsLiteralCi(class_name, "self") {
			if scope == nil {
				ZendThrowError(nil, "Cannot access self:: when no class scope is active")
				goto failure
			}
			ce = scope
		} else if ZendStringEqualsLiteralCi(class_name, "parent") {
			if scope == nil {
				ZendThrowError(nil, "Cannot access parent:: when no class scope is active")
				goto failure
			} else if !(scope.GetParent()) {
				ZendThrowError(nil, "Cannot access parent:: when current class scope has no parent")
				goto failure
			} else {
				ce = scope.GetParent()
			}
		} else if ZendStringEqualsLiteralCi(class_name, "static") {
			ce = ZendGetCalledScope(CurrEX())
			if ce == nil {
				ZendThrowError(nil, "Cannot access static:: when no class scope is active")
				goto failure
			}
		} else {
			ce = ZendFetchClass(class_name, flags)
		}
		if ce != nil {
			c = ZendHashFindPtr(ce.GetConstantsTable(), constant_name)
			if c == nil {
				if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
					ZendThrowError(nil, "Undefined class constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
					goto failure
				}
				ret_constant = nil
			} else {
				if ZendVerifyConstAccess(c, scope) == 0 {
					if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
						ZendThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), class_name.GetVal(), constant_name.GetVal())
					}
					goto failure
				}
				ret_constant = c.GetValue()
			}
		}
		if ret_constant != nil && ret_constant.IsConstant() {
			var ret int
			if IS_CONSTANT_VISITED(ret_constant) != 0 {
				ZendThrowError(nil, "Cannot declare self-referencing constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
				ret_constant = nil
				goto failure
			}
			MARK_CONSTANT_VISITED(ret_constant)
			ret = ZvalUpdateConstantEx(ret_constant, c.GetCe())
			RESET_CONSTANT_VISITED(ret_constant)
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
		if b.Assign(&c, ZendHashStrFindPtr(EG__().GetZendConstants(), lcname, lcname_len)) == nil {

			/* try lowercase */

			ZendStrTolower(lcname+prefix_len+1, const_name_len)
			if b.Assign(&c, ZendHashStrFindPtr(EG__().GetZendConstants(), lcname, lcname_len)) != nil {
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
			ZendError(E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
		}
	}
	return c.GetValue()
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
		var slash *byte = strrchr(c.GetName().GetVal(), '\\')
		if slash != nil {
			lowercase_name = ZendStringInit(c.GetName().GetVal(), c.GetName().GetLen(), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
			ZendStrTolower(lowercase_name.GetVal(), slash-c.GetName().GetVal())
			lowercase_name = ZendNewInternedString(lowercase_name)
			name = lowercase_name
		} else {
			name = c.GetName()
		}
	}

	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */

	if ZendStringEqualsLiteral(name, "__COMPILER_HALT_OFFSET__") || ZendHashAddConstant(EG__().GetZendConstants(), name, c) == nil {

		/* The internal __COMPILER_HALT_OFFSET__ is prefixed by NULL byte */

		if c.GetName().GetVal()[0] == '0' && c.GetName().GetLen() > b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")-1 && memcmp(name.GetVal(), "0__COMPILER_HALT_OFFSET__", b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")) == 0 {

		}
		ZendError(E_NOTICE, "Constant %s already defined", name.GetVal())
		ZendStringRelease(c.GetName())
		if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
			ZvalPtrDtorNogc(c.GetValue())
		}
		ret = FAILURE
	}
	if lowercase_name != nil {
		ZendStringRelease(lowercase_name)
	}
	return ret
}
