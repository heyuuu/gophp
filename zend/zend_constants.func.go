package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"strings"
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

func FreeZendConstantEx(c *ZendConstant) {
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
		ZvalPtrDtorNogc(c.Value())
		if c.GetName() != nil {
			// types.ZendStringReleaseEx(c.GetName(), 0)
		}
		Efree(c)
	} else {
		ZvalInternalPtrDtor(c.Value())
		if c.GetName() != nil {
			// types.ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}

func FreeZendConstant(zv *types.Zval) {
	var c *ZendConstant = zv.GetPtr()
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
		ZvalPtrDtorNogc(c.Value())
		if c.GetName() != nil {
			// types.ZendStringReleaseEx(c.GetName(), 0)
		}
		Efree(c)
	} else {
		ZvalInternalPtrDtor(c.Value())
		if c.GetName() != nil {
			// types.ZendStringReleaseEx(c.GetName(), 1)
		}
		Free(c)
	}
}
func CleanModuleConstants(moduleNumber int) {
	EG__().ConstantTable().Filter(func(_ string, c *ZendConstant) bool {
		return ZEND_CONSTANT_MODULE_NUMBER(c) != moduleNumber
	})
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
	RegisterMainBoolConstant("ZEND_THREAD_SAFE", 0, CONST_PERSISTENT|CONST_CS)
	RegisterMainBoolConstant("ZEND_DEBUG_BUILD", 0, CONST_PERSISTENT|CONST_CS)
	RegisterMainNullConstant("NULL", CONST_PERSISTENT|CONST_CT_SUBST)
}
func ZendRegisterNullConstant(name string, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetNull()
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetNameVal(name)
	ZendRegisterConstant(&c)
}
func ZendRegisterBoolConstant(name string, bval types.ZendBool, flags int, module_number int) {
	var c ZendConstant
	types.ZVAL_BOOL(c.Value(), bval != 0)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetNameVal(name)
	ZendRegisterConstant(&c)
}
func ZendRegisterLongConstant(name string, lval ZendLong, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetLong(lval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetNameVal(name)
	ZendRegisterConstant(&c)
}
func ZendRegisterDoubleConstant(name string, dval float64, flags int, module_number int) {
	var c ZendConstant
	c.Value().SetDouble(dval)
	ZEND_CONSTANT_SET_FLAGS(&c, flags, module_number)
	c.SetNameVal(name)
	ZendRegisterConstant(&c)
}
func ZendRegisterStringlConstant(name string, str string, flags int, module_number int) {
	c := NewConstant(name)
	c.Value().SetStringVal(str)

	ZEND_CONSTANT_SET_FLAGS(c, flags, module_number)
	ZendRegisterConstant(c)
}
func ZendRegisterStringConstant(name string, strval *byte, flags int, module_number int) {
	ZendRegisterStringlConstant(name, strval, flags, module_number)
}
func ZendGetSpecialConstant(name string) *ZendConstant {
	var haltoff = "__COMPILER_HALT_OFFSET__"
	if CurrEX() == nil {
		return nil
	} else if name == haltoff {
		cfilename := ZendGetExecutedFilename()

		/* check for __COMPILER_HALT_OFFSET__ */
		haltname := ZendManglePropertyName_Ex(haltoff, cfilename)
		return EG__().ConstantTable().Get(haltname)
	} else {
		return nil
	}
}
func ZendVerifyConstAccess(c *ZendClassConstant, scope *types.ClassEntry) int {
	if (c.GetValue().GetAccessFlags() & AccPublic) != 0 {
		return 1
	} else if (c.GetValue().GetAccessFlags() & AccPrivate) != 0 {
		return c.GetCe() == scope
	} else {
		b.Assert((c.GetValue().GetAccessFlags() & AccProtected) != 0)
		return ZendCheckProtected(c.GetCe(), scope)
	}
}
func ZendGetConstantImpl(name string) *ZendConstant {
	var c *ZendConstant
	c = EG__().ConstantTable().Get(name)
	if c == nil {
		c = EG__().ConstantTable().Get(ascii.StrToLower(name))
		if c != nil {
			if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
				c = nil
			}
		} else {
			c = ZendGetSpecialConstant(name)
		}
	}
	return c
}
func ZendGetConstant(name string) *types.Zval {
	var c *ZendConstant = ZendGetConstantImpl(name)
	if c != nil {
		return c.Value()
	} else {
		return nil
	}
}
func IsAccessDeprecated(c *ZendConstant, access_name *byte) types.ZendBool {
	var ns_sep *byte = ZendMemrchr(c.GetName().GetStr(), '\\', c.GetName().GetLen())
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
	var name_ string = cname.GetStr()
	var name *byte = cname.GetVal()
	var name_len int = cname.GetLen()

	/* Skip leading \\ */
	if name_ != "" && name_[0] == '\\' {
		name_ = name_[1:]
		cname = nil
	}
	if pos := strings.LastIndexByte(name_, ':'); pos > 0 && name_[pos-1] == ':' {
		var constantName = name_[pos+1:]
		var className = name_[:pos-1]
		var c *ZendClassConstant = nil
		var ret_constant *types.Zval = nil
		var ce *types.ClassEntry
		if ascii.StrCaseEquals(className, "self") {
			if scope == nil {
				faults.ThrowError(nil, "Cannot access self:: when no class scope is active")
				goto failure
			}
			ce = scope
		} else if ascii.StrCaseEquals(className, "parent") {
			if scope == nil {
				faults.ThrowError(nil, "Cannot access parent:: when no class scope is active")
				goto failure
			} else if !(scope.GetParent()) {
				faults.ThrowError(nil, "Cannot access parent:: when current class scope has no parent")
				goto failure
			} else {
				ce = scope.GetParent()
			}
		} else if ascii.StrCaseEquals(className, "static") {
			ce = ZendGetCalledScope(CurrEX())
			if ce == nil {
				faults.ThrowError(nil, "Cannot access static:: when no class scope is active")
				goto failure
			}
		} else {
			ce = ZendFetchClass(className, flags)
		}
		if ce != nil {
			c = types.ZendHashFindPtr(ce.GetConstantsTable(), constantName)
			if c == nil {
				if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
					faults.ThrowError(nil, "Undefined class constant '%s::%s'", className, constantName)
					goto failure
				}
				ret_constant = nil
			} else {
				if ZendVerifyConstAccess(c, scope) == 0 {
					if (flags & ZEND_FETCH_CLASS_SILENT) == 0 {
						faults.ThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), className, constantName)
					}
					goto failure
				}
				ret_constant = c.GetValue()
			}
		}
		if ret_constant != nil && ret_constant.IsConstant() {
			var ret int
			if IS_CONSTANT_VISITED(ret_constant) {
				faults.ThrowError(nil, "Cannot declare self-referencing constant '%s::%s'", className, constantName)
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
		// types.ZendStringReleaseEx(class_name, 0)
		// types.ZendStringEfree(constant_name)
		return ret_constant
	}

	/* non-class constant */
	if pos := strings.LastIndexByte(name_, '\\'); pos >= 0 {
		/* compound constant name */
		lcPrefix := ascii.StrToLower(name_[:pos]) + "\\"
		constName := name_[pos+1:]

		/* Check for namespace constant */
		// 查找常量顺序
		// - 不忽略大小写查找 (searchName=命名空间名小写+原常量名)
		// - 忽略大小写查找且判断常量属性非大小写敏感 (searchName=命名空间名小写+常量名小写)
		searchName := lcPrefix + constName
		if c = EG__().ConstantTable().Get(searchName); c == nil {
			/* try lowercase */
			searchName = lcPrefix + ascii.StrToLower(constName)
			if c = EG__().ConstantTable().Get(searchName); c != nil {
				if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) != 0 {
					c = nil
				}
			}
		}
		if c == nil {
			if (flags & IS_CONSTANT_UNQUALIFIED) == 0 {
				return nil
			}

			/* name requires runtime resolution, need to check non-namespaced name */
			c = ZendGetConstantImpl(constName)
			name_ = constName
		}
	} else {
		if cname != nil {
			c = ZendGetConstantImpl(cname.GetStr())
		} else {
			c = ZendGetConstantImpl(name_)
		}
	}
	if c == nil {
		return nil
	}
	if (flags & ZEND_GET_CONSTANT_NO_DEPRECATION_CHECK) == 0 {
		if (ZEND_CONSTANT_FLAGS(c)&(CONST_CS|CONST_CT_SUBST)) == 0 && IsAccessDeprecated(c, name_) != 0 {
			faults.Error(faults.E_DEPRECATED, "Case-insensitive constants are deprecated. "+"The correct casing for this constant is \"%s\"", c.GetName().GetVal())
		}
	}
	return c.Value()
}

func ZendRegisterConstant(c *ZendConstant) int {
	var lowercase_name *types.String = nil
	var name *types.String
	var ret int = types.SUCCESS
	if (ZEND_CONSTANT_FLAGS(c) & CONST_CS) == 0 {
		lowercase_name = ZendStringTolowerEx(c.GetName())
		// lowercase_name = types.ZendNewInternedString(lowercase_name)
		name = lowercase_name
	} else {
		var slash *byte = strrchr(c.GetName().GetVal(), '\\')
		if slash != nil {
			lowercase_name = types.NewString(c.GetName().GetStr())
			ZendStrTolower(lowercase_name.GetVal(), slash-c.GetName().GetVal())
			// lowercase_name = types.ZendNewInternedString(lowercase_name)
			name = lowercase_name
		} else {
			name = c.GetName()
		}
	}

	/* Check if the user is trying to define the __special__  internal pseudo constant name __COMPILER_HALT_OFFSET__ */

	if name.GetStr() == "__COMPILER_HALT_OFFSET__" || !EG__().ConstantTable().Add(name.GetStr(), CopyConstant(c)) {
		faults.Error(faults.E_NOTICE, "Constant %s already defined", name.GetVal())
		if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) == 0 {
			ZvalPtrDtorNogc(c.Value())
		}
		ret = types.FAILURE
	}
	return ret
}
