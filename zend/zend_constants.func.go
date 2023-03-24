package zend

import (
	b "sik/builtin"
	"sik/zend/faults"
	"sik/zend/types"
)

func ZEND_CONSTANT_FLAGS(c *ZendConstant) uint8       { return c.Flags() }
func ZEND_CONSTANT_MODULE_NUMBER(c *ZendConstant) int { return c.ModuleNumber() }
func ZEND_CONSTANT_SET_FLAGS(c *ZendConstant, _flags int, _module_number int) {
	c.Value().GetConstantFlags() = _flags&0xff | _module_number<<8
}
func IS_CONSTANT_VISITED(zv *types.Zval) bool {
	return zv.GetAccessFlags()&IS_CONSTANT_VISITED_MARK != 0
}
func MARK_CONSTANT_VISITED(zv *types.Zval) uint32 {
	zv.AddAccessFlags(IS_CONSTANT_VISITED_MARK)
	return zv.GetAccessFlags()
}
func RESET_CONSTANT_VISITED(zv *types.Zval) uint32 {
	zv.SubAccessFlags(IS_CONSTANT_VISITED_MARK)
	return zv.GetAccessFlags()
}

func RegisterLongConstant(name string, lval ZendLong, flags int, module_number int) {
	ZendRegisterLongConstant(name, lval, flags, module_number)
}
func RegisterDoubleConstant(name string, dval float64, flags int, module_number int) {
	ZendRegisterDoubleConstant(name, dval, flags, module_number)
}
func RegisterStringConstant(name string, str *byte, flags int, module_number int) {
	ZendRegisterStringConstant(name, str, flags, module_number)
}
func RegisterMainNullConstant(name string, flags int) {
	ZendRegisterNullConstant(name, flags, 0)
}
func RegisterMainBoolConstant(name string, bval types.ZendBool, flags int) {
	ZendRegisterBoolConstant(name, bval, flags, 0)
}
func RegisterMainLongConstant(name string, lval ZendLong, flags int) {
	ZendRegisterLongConstant(name, lval, flags, 0)
}
func RegisterMainDoubleConstant(name string, dval float64, flags int) {
	ZendRegisterDoubleConstant(name, dval, flags, 0)
}
func RegisterMainStringConstant(name string, str string, flags int) {
	ZendRegisterStringlConstant(name, str, flags, 0)
}

func FreeZendConstant(zv *types.Zval) {
	var c *ZendConstant = zv.GetPtr()
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
		ZvalPtrDtorNogc(c.Value())
		if c.GetName() != nil {
			types.ZendStringReleaseEx(c.GetName(), 0)
		}
		Efree(c)
	} else {
		ZvalInternalPtrDtor(c.Value())
		if c.GetName() != nil {
			types.ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}
func CleanModuleConstant(el *types.Zval, arg any) int {
	var c *ZendConstant = (*ZendConstant)(el.GetPtr())
	var module_number int = *((*int)(arg))
	if ZEND_CONSTANT_MODULE_NUMBER(c) == module_number {
		return types.ArrayApplyRemove
	} else {
		return types.ArrayApplyKeep
	}
}
func CleanModuleConstants(module_number int) {
	types.ZendHashApplyWithArgument(EG__().GetZendConstants(), CleanModuleConstant, any(&module_number))
}
func ZendRegisterStandardConstants() {
	RegisterMainLongConstant("E_ERROR", faults.E_ERROR, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_RECOVERABLE_ERROR", faults.E_RECOVERABLE_ERROR, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_WARNING", faults.E_WARNING, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_PARSE", faults.E_PARSE, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_NOTICE", faults.E_NOTICE, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_STRICT", faults.E_STRICT, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_DEPRECATED", faults.E_DEPRECATED, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_CORE_ERROR", faults.E_CORE_ERROR, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_CORE_WARNING", faults.E_CORE_WARNING, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_COMPILE_ERROR", faults.E_COMPILE_ERROR, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_COMPILE_WARNING", faults.E_COMPILE_WARNING, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_USER_ERROR", faults.E_USER_ERROR, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_USER_WARNING", faults.E_USER_WARNING, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_USER_NOTICE", faults.E_USER_NOTICE, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_USER_DEPRECATED", faults.E_USER_DEPRECATED, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("E_ALL", faults.E_ALL, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("DEBUG_BACKTRACE_PROVIDE_OBJECT", DEBUG_BACKTRACE_PROVIDE_OBJECT, CONST_PERSISTENT|CONST_CS)
	RegisterMainLongConstant("DEBUG_BACKTRACE_IGNORE_ARGS", DEBUG_BACKTRACE_IGNORE_ARGS, CONST_PERSISTENT|CONST_CS)

	/* true/false constants */

	RegisterMainBoolConstant("TRUE", 1, CONST_PERSISTENT|CONST_CT_SUBST)
	RegisterMainBoolConstant("FALSE", 0, CONST_PERSISTENT|CONST_CT_SUBST)
	RegisterMainBoolConstant("ZEND_THREAD_SAFE", ZTS_V, CONST_PERSISTENT|CONST_CS)
	RegisterMainBoolConstant("ZEND_DEBUG_BUILD", 0, CONST_PERSISTENT|CONST_CS)
	RegisterMainNullConstant("NULL", CONST_PERSISTENT|CONST_CT_SUBST)
}
func ZendRegisterNullConstant(name string, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetNull()
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(types.ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterBoolConstant(name string, bval types.ZendBool, flags int, module_number int) {
	var c ZendConstant
	types.ZVAL_BOOL(c.Value(), bval != 0)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(types.ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterLongConstant(name string, lval ZendLong, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetLong(lval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(types.ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterDoubleConstant(name string, dval float64, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetDouble(dval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetName(types.ZendStringInitInterned(name, name_len, flags&CONST_PERSISTENT))
	ZendRegisterConstant(&c)
}
func ZendRegisterStringlConstant(name string, str string, flags int, module_number int) {
	c := NewZendConstant(name)
	c.Value().SetRawString(str)

	ZEND_CONSTANT_SET_FLAGS(c, flags, module_number)
	ZendRegisterConstant(c)
}
func ZendRegisterStringConstant(name string, strval *byte, flags int, module_number int) {
	ZendRegisterStringlConstant(name, strval, flags, module_number)
}
func ZendGetSpecialConstant(name string) *ZendConstant {
	var c *ZendConstant
	var haltoff = "__COMPILER_HALT_OFFSET__"
	if CurrEX() == nil {
		return nil
	} else if name_len == b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1 && !(memcmp(name, "__COMPILER_HALT_OFFSET__", b.SizeOf("\"__COMPILER_HALT_OFFSET__\"")-1)) {
		cfilename := ZendGetExecutedFilename()

		/* check for __COMPILER_HALT_OFFSET__ */

		haltname := ZendManglePropertyName_Ex(haltoff, cfilename)
		c = types.ZendHashFindPtr(EG__().GetZendConstants(), haltname.GetStr())
		return c
	} else {
		return nil
	}
}
func ZendVerifyConstAccess(c *ZendClassConstant, scope *types.ClassEntry) int {
	if (c.GetValue().GetAccessFlags() & ZEND_ACC_PUBLIC) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & ZEND_ACC_PRIVATE) != 0 {
		return c.GetCe() == scope
	} else {
		b.Assert((c.GetValue().GetAccessFlags() & ZEND_ACC_PROTECTED) != 0)
		return ZendCheckProtected(c.GetCe(), scope)
	}
}
func ZendGetConstantStrImpl(name string) *ZendConstant {
	var c *ZendConstant
	if b.Assign(&c, types.ZendHashStrFindPtr(EG__().GetZendConstants(), b.CastStr(name, name_len))) == nil {
		var lcname *byte = DoAlloca(name_len+1, use_heap)
		ZendStrTolowerCopy(lcname, name, name_len)
		if b.Assign(&c, types.ZendHashStrFindPtr(EG__().GetZendConstants(), b.CastStr(lcname, name_len))) != nil {
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
func ZendGetConstantImpl(name *types.String) *ZendConstant {
	var zv *types.Zval
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
func ZendGetConstant(name *types.String) *types.Zval {
	var c *ZendConstant = ZendGetConstantImpl(name)
	if c != nil {
		return c.Value()
	} else {
		return nil
	}
}
func IsAccessDeprecated(c *ZendConstant, access_name *byte) types.ZendBool {
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
func ZendGetConstantEx(cname *types.String, scope *types.ClassEntry, flags uint32) *types.Zval {
	var c *ZendConstant
	var colon *byte
	var ce *types.ClassEntry = nil
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
		var constant_name *types.String = types.NewString(b.CastStr(colon+1, const_name_len))
		var class_name *types.String = types.NewString(b.CastStr(name, class_name_len))
		var c *ZendClassConstant = nil
		var ret_constant *types.Zval = nil
		if types.ZendStringEqualsLiteralCi(class_name, "self") {
			if scope == nil {
				faults.ThrowError(nil, "Cannot access self:: when no class scope is active")
				goto failure
			}
			ce = scope
		} else if types.ZendStringEqualsLiteralCi(class_name, "parent") {
			if scope == nil {
				faults.ThrowError(nil, "Cannot access parent:: when no class scope is active")
				goto failure
			} else if !(scope.GetParent()) {
				faults.ThrowError(nil, "Cannot access parent:: when current class scope has no parent")
				goto failure
			} else {
				ce = scope.GetParent()
			}
		} else if types.ZendStringEqualsLiteralCi(class_name, "static") {
			ce = ZendGetCalledScope(CurrEX())
			if ce == nil {
				faults.ThrowError(nil, "Cannot access static:: when no class scope is active")
				goto failure
			}
		} else {
			ce = ZendFetchClass(class_name, flags)
		}
		if ce != nil {
			c = types.ZendHashFindPtr(ce.GetConstantsTable(), constant_name.GetStr())
			if c == nil {
				if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
					faults.ThrowError(nil, "Undefined class constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
					goto failure
				}
				ret_constant = nil
			} else {
				if ZendVerifyConstAccess(c, scope) == 0 {
					if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
						faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), class_name.GetVal(), constant_name.GetVal())
					}
					goto failure
				}
				ret_constant = c.GetValue()
			}
		}
		if ret_constant != nil && ret_constant.IsConstant() {
			var ret int
			if IS_CONSTANT_VISITED(ret_constant) {
				faults.ThrowError(nil, "Cannot declare self-referencing constant '%s::%s'", class_name.GetVal(), constant_name.GetVal())
				ret_constant = nil
				goto failure
			}
			MARK_CONSTANT_VISITED(ret_constant)
			ret = ZvalUpdateConstantEx(ret_constant, c.GetCe())
			RESET_CONSTANT_VISITED(ret_constant)
			if ret != types.SUCCESS {
				ret_constant = nil
				goto failure
			}
		}
	failure:
		types.ZendStringReleaseEx(class_name, 0)
		types.ZendStringEfree(constant_name)
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
		if b.Assign(&c, types.ZendHashStrFindPtr(EG__().GetZendConstants(), b.CastStr(lcname, lcname_len))) == nil {

			/* try lowercase */

			ZendStrTolower(lcname+prefix_len+1, const_name_len)
			if b.Assign(&c, types.ZendHashStrFindPtr(EG__().GetZendConstants(), b.CastStr(lcname, lcname_len))) != nil {
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
			faults.Error(faults.E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
		}
	}
	return c.Value()
}
func ZendHashAddConstant(ht *types.Array, key *types.String, c *ZendConstant) any {
	var ret any
	var copy *ZendConstant = Pemalloc(b.SizeOf("zend_constant"), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
	memcpy(copy, c, b.SizeOf("zend_constant"))
	ret = types.ZendHashAddPtr(ht, key.GetStr(), copy)
	if !ret {
		Pefree(copy, ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
	}
	return ret
}

func ZendRegisterConstant(c *ZendConstant) int {
	var lowercase_name *types.String = nil
	var name *types.String
	var ret int = types.SUCCESS
	if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) == 0 {
		lowercase_name = ZendStringTolowerEx(c.GetName(), ZEND_CONSTANT_FLAGS(c)&CONST_PERSISTENT)
		lowercase_name = types.ZendNewInternedString(lowercase_name)
		name = lowercase_name
	} else {
		var slash *byte = strrchr(c.GetName().GetVal(), '\\')
		if slash != nil {
			lowercase_name = types.NewString(c.GetName().GetStr())
			ZendStrTolower(lowercase_name.GetVal(), slash-c.GetName().GetVal())
			lowercase_name = types.ZendNewInternedString(lowercase_name)
			name = lowercase_name
		} else {
			name = c.GetName()
		}
	}

	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */

	if types.ZendStringEqualsLiteral(name, "__COMPILER_HALT_OFFSET__") || ZendHashAddConstant(EG__().GetZendConstants(), name, c) == nil {

		/* The internal __COMPILER_HALT_OFFSET__ is prefixed by NULL byte */

		if c.GetName().GetVal()[0] == '0' && c.GetName().GetLen() > b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")-1 && memcmp(name.GetVal(), "0__COMPILER_HALT_OFFSET__", b.SizeOf("\"\\0__COMPILER_HALT_OFFSET__\"")) == 0 {

		}
		faults.Error(faults.E_NOTICE, "Constant %s already defined", name.GetVal())
		types.ZendStringRelease(c.GetName())
		if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
			ZvalPtrDtorNogc(c.Value())
		}
		ret = types.FAILURE
	}
	if lowercase_name != nil {
		types.ZendStringRelease(lowercase_name)
	}
	return ret
}
