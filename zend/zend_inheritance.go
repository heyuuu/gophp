// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_inheritance.h>

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

// #define ZEND_INHERITANCE_H

// # include "zend.h"

// #define zend_do_inheritance(ce,parent_ce) zend_do_inheritance_ex ( ce , parent_ce , 0 )

// Source: <Zend/zend_inheritance.c>

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

// # include "zend_API.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_inheritance.h"

// # include "zend_interfaces.h"

// # include "zend_smart_str.h"

// # include "zend_operators.h"

// # include "zend_exceptions.h"

func OverriddenPtrDtor(zv *Zval) { _efree(zv.GetValue().GetPtr()) }

/* }}} */

func ZendDuplicatePropertyInfoInternal(property_info *ZendPropertyInfo) *ZendPropertyInfo {
	var new_property_info *ZendPropertyInfo = g.CondF(true, func() any { return __zendMalloc(g.SizeOf("zend_property_info")) }, func() any { return _emalloc(g.SizeOf("zend_property_info")) })
	memcpy(new_property_info, property_info, g.SizeOf("zend_property_info"))
	ZendStringAddref(new_property_info.GetName())
	if new_property_info.GetType() > 0x3ff && (new_property_info.GetType()&0x2) == 0 {
		ZendStringAddref((*ZendString)(new_property_info.GetType() & ^0x3))
	}
	return new_property_info
}

/* }}} */

func ZendDuplicateInternalFunction(func_ *ZendFunction, ce *ZendClassEntry) *ZendFunction {
	var new_function *ZendFunction
	if (ce.GetType() & 1) != 0 {
		new_function = __zendMalloc(g.SizeOf("zend_internal_function"))
		memcpy(new_function, func_, g.SizeOf("zend_internal_function"))
	} else {
		new_function = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_internal_function"))
		memcpy(new_function, func_, g.SizeOf("zend_internal_function"))
		new_function.SetFnFlags(new_function.GetFnFlags() | 1<<25)
	}
	if new_function.GetFunctionName() != nil {
		ZendStringAddref(new_function.GetFunctionName())
	}
	return new_function
}

/* }}} */

func ZendDuplicateUserFunction(func_ *ZendFunction) *ZendFunction {
	var new_function *ZendFunction
	new_function = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_op_array"))
	memcpy(new_function, func_, g.SizeOf("zend_op_array"))
	if g.CondF((uintPtr(func_.GetOpArray().GetStaticVariablesPtrPtr())&1) != 0, func() any {
		return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(func_.GetOpArray().GetStaticVariablesPtrPtr()-1))))
	}, func() any { return any(*(func_.GetOpArray().GetStaticVariablesPtrPtr())) }) {

		/* See: Zend/tests/method_static_var.phpt */

		if (uintPtr(func_.GetOpArray().GetStaticVariablesPtrPtr()) & 1) != 0 {
			new_function.GetOpArray().SetStaticVariables(*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(func_.GetOpArray().GetStaticVariablesPtrPtr()-1)))))
		} else {
			new_function.GetOpArray().SetStaticVariables(any(*(func_.GetOpArray().GetStaticVariablesPtrPtr())))
		}

		/* See: Zend/tests/method_static_var.phpt */

	}
	if (ZvalGcFlags(new_function.GetOpArray().GetStaticVariables().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
		ZendGcAddref(&(new_function.GetOpArray().GetStaticVariables()).gc)
	}
	if (CG.GetCompilerOptions() & 1 << 15) != 0 {
		assert((new_function.GetOpArray().GetFnFlags() & 1 << 10) != 0)
		new_function.GetOpArray().SetStaticVariablesPtrPtr(ZendMapPtrNew())
	} else {
		new_function.GetOpArray().SetStaticVariablesPtrPtr(&new_function.op_array.GetStaticVariables())
	}
	return new_function
}

/* }}} */

func ZendDuplicateFunction(func_ *ZendFunction, ce *ZendClassEntry, is_interface ZendBool) *ZendFunction {
	if func_.GetType() == 1 {
		return ZendDuplicateInternalFunction(func_, ce)
	} else {
		if func_.GetOpArray().GetRefcount() != nil {
			(*func_).op_array.refcount++
		}
		if is_interface != 0 || func_.GetOpArray().GetStaticVariables() == nil {

			/* reuse the same op_array structure */

			return func_

			/* reuse the same op_array structure */

		}
		return ZendDuplicateUserFunction(func_)
	}
}

/* }}} */

func DoInheritParentConstructor(ce *ZendClassEntry) {
	var parent *ZendClassEntry = ce.parent
	assert(parent != nil)

	/* You cannot change create_object */

	ce.create_object = parent.create_object

	/* Inherit special functions if needed */

	if ce.GetGetIterator() == nil {
		ce.SetGetIterator(parent.GetGetIterator())
	}
	if parent.GetIteratorFuncsPtr() != nil {

		/* Must be initialized through iface->interface_gets_implemented() */

		assert(ce.GetIteratorFuncsPtr() != nil)

		/* Must be initialized through iface->interface_gets_implemented() */

	}
	if ce.GetGet() == nil {
		ce.SetGet(parent.GetGet())
	}
	if ce.GetSet() == nil {
		ce.SetSet(parent.GetSet())
	}
	if ce.GetUnset() == nil {
		ce.SetUnset(parent.GetUnset())
	}
	if ce.GetIsset() == nil {
		ce.SetIsset(parent.GetIsset())
	}
	if ce.GetCall() == nil {
		ce.SetCall(parent.GetCall())
	}
	if ce.GetCallstatic() == nil {
		ce.SetCallstatic(parent.GetCallstatic())
	}
	if ce.GetTostring() == nil {
		ce.SetTostring(parent.GetTostring())
	}
	if ce.GetClone() == nil {
		ce.SetClone(parent.GetClone())
	}
	if ce.GetSerializeFunc() == nil {
		ce.SetSerializeFunc(parent.GetSerializeFunc())
	}
	if ce.GetSerialize() == nil {
		ce.SetSerialize(parent.GetSerialize())
	}
	if ce.GetUnserializeFunc() == nil {
		ce.SetUnserializeFunc(parent.GetUnserializeFunc())
	}
	if ce.GetUnserialize() == nil {
		ce.SetUnserialize(parent.GetUnserialize())
	}
	if ce.GetDestructor() == nil {
		ce.SetDestructor(parent.GetDestructor())
	}
	if ce.GetDebugInfo() == nil {
		ce.SetDebugInfo(parent.GetDebugInfo())
	}
	if ce.GetConstructor() != nil {
		if parent.GetConstructor() != nil && (parent.GetConstructor().GetFnFlags()&1<<5) != 0 {
			ZendErrorNoreturn(1<<0, "Cannot override final %s::%s() with %s::%s()", parent.GetName().GetVal(), parent.GetConstructor().GetFunctionName().GetVal(), ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
		return
	}
	ce.SetConstructor(parent.GetConstructor())
}

/* }}} */

func ZendVisibilityString(fn_flags uint32) *byte {
	if (fn_flags & 1 << 0) != 0 {
		return "public"
	} else if (fn_flags & 1 << 2) != 0 {
		return "private"
	} else {
		assert((fn_flags & 1 << 1) != 0)
		return "protected"
	}
}

/* }}} */

func ResolveClassName(scope *ZendClassEntry, name *ZendString) *ZendString {
	assert(scope != nil)
	if name.GetLen() == g.SizeOf("\"parent\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "parent", g.SizeOf("\"parent\"")-1) == 0 && scope.parent {
		if (scope.GetCeFlags() & 1 << 19) != 0 {
			return scope.parent.name
		} else {
			return scope.parent_name
		}
	} else if name.GetLen() == g.SizeOf("\"self\"")-1 && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), "self", g.SizeOf("\"self\"")-1) == 0 {
		return scope.GetName()
	} else {
		return name
	}
}
func ClassVisible(ce *ZendClassEntry) ZendBool {
	if ce.GetType() == 1 {
		return !(CG.GetCompilerOptions() & 1 << 4)
	} else {
		assert(ce.GetType() == 2)
		return (CG.GetCompilerOptions()&1<<13) == 0 || ce.GetFilename() == CG.GetCompiledFilename()
	}
}
func LookupClass(scope *ZendClassEntry, name *ZendString) *ZendClassEntry {
	var ce *ZendClassEntry
	if CG.GetInCompilation() == 0 {
		var flags uint32 = 0x400 | 0x80
		ce = ZendLookupClassEx(name, nil, flags)
		if ce != nil {
			return ce
		}

		/* We'll autoload this class and process delayed variance obligations later. */

		if CG.GetDelayedAutoloads() == nil {
			CG.SetDelayedAutoloads((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
			_zendHashInit(CG.GetDelayedAutoloads(), 0, nil, 0)
		}
		ZendHashAddEmptyElement(CG.GetDelayedAutoloads(), name)
	} else {
		ce = ZendLookupClassEx(name, nil, 0x80)
		if ce != nil && ClassVisible(ce) != 0 {
			return ce
		}

		/* The current class may not be registered yet, so check for it explicitly. */

		if scope.GetName().GetLen() == name.GetLen() && ZendBinaryStrcasecmp(scope.GetName().GetVal(), scope.GetName().GetLen(), name.GetVal(), name.GetLen()) == 0 {
			return scope
		}

		/* The current class may not be registered yet, so check for it explicitly. */

	}
	return nil
}

/* Instanceof that's safe to use on unlinked classes. */

func UnlinkedInstanceof(ce1 *ZendClassEntry, ce2 *ZendClassEntry) ZendBool {
	if ce1 == ce2 {
		return 1
	}
	if (ce1.GetCeFlags() & 1 << 3) != 0 {
		return InstanceofFunction(ce1, ce2)
	}
	if ce1.parent {
		var parent_ce *ZendClassEntry
		if (ce1.GetCeFlags() & 1 << 19) != 0 {
			parent_ce = ce1.parent
		} else {
			parent_ce = ZendLookupClassEx(ce1.parent_name, nil, 0x400|0x80)
		}

		/* It's not sufficient to only check the parent chain itself, as need to do a full
		 * recursive instanceof in case the parent interfaces haven't been copied yet. */

		if parent_ce != nil && UnlinkedInstanceof(parent_ce, ce2) != 0 {
			return 1
		}

		/* It's not sufficient to only check the parent chain itself, as need to do a full
		 * recursive instanceof in case the parent interfaces haven't been copied yet. */

	}
	if ce1.GetNumInterfaces() != 0 {
		var i uint32
		if (ce1.GetCeFlags() & 1 << 20) != 0 {

			/* Unlike the normal instanceof_function(), we have to perform a recursive
			 * check here, as the parent interfaces might not have been fully copied yet. */

			for i = 0; i < ce1.GetNumInterfaces(); i++ {
				if UnlinkedInstanceof(ce1.interfaces[i], ce2) != 0 {
					return 1
				}
			}

			/* Unlike the normal instanceof_function(), we have to perform a recursive
			 * check here, as the parent interfaces might not have been fully copied yet. */

		} else {
			for i = 0; i < ce1.GetNumInterfaces(); i++ {
				var ce *ZendClassEntry = ZendLookupClassEx(ce1.interface_names[i].name, ce1.interface_names[i].lc_name, 0x400|0x80)
				if ce != nil && UnlinkedInstanceof(ce, ce2) != 0 {
					return 1
				}
			}
		}
	}
	return 0
}

/* Unresolved means that class declarations that are currently not available are needed to
 * determine whether the inheritance is valid or not. At runtime UNRESOLVED should be treated
 * as an ERROR. */

type InheritanceStatus = int

const (
	INHERITANCE_UNRESOLVED InheritanceStatus = -1
	INHERITANCE_ERROR                        = 0
	INHERITANCE_SUCCESS                      = 1
)

func ZendPerformCovariantTypeCheck(unresolved_class **ZendString, fe *ZendFunction, fe_arg_info *ZendArgInfo, proto *ZendFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	var fe_type ZendType = fe_arg_info.GetType()
	var proto_type ZendType = proto_arg_info.GetType()
	assert(fe_type > 0x3 && proto_type > 0x3)
	if (fe_type&0x1) != 0 && (proto_type&0x1) == 0 {
		return INHERITANCE_ERROR
	}
	if proto_type > 0x3ff {
		var fe_class_name *ZendString
		var proto_class_name *ZendString
		var fe_ce *ZendClassEntry
		var proto_ce *ZendClassEntry
		if fe_type <= 0x3ff {
			return INHERITANCE_ERROR
		}
		fe_class_name = ResolveClassName(fe.GetScope(), (*ZendString)(fe_type & ^0x3))
		proto_class_name = ResolveClassName(proto.GetScope(), (*ZendString)(proto_type & ^0x3))
		if fe_class_name.GetLen() == proto_class_name.GetLen() && ZendBinaryStrcasecmp(fe_class_name.GetVal(), fe_class_name.GetLen(), proto_class_name.GetVal(), proto_class_name.GetLen()) == 0 {
			return INHERITANCE_SUCCESS
		}

		/* Make sure to always load both classes, to avoid only registering one of them as
		 * a delayed autoload. */

		fe_ce = LookupClass(fe.GetScope(), fe_class_name)
		proto_ce = LookupClass(proto.GetScope(), proto_class_name)
		if fe_ce == nil {
			*unresolved_class = fe_class_name
			return INHERITANCE_UNRESOLVED
		}
		if proto_ce == nil {
			*unresolved_class = proto_class_name
			return INHERITANCE_UNRESOLVED
		}
		if UnlinkedInstanceof(fe_ce, proto_ce) != 0 {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else if proto_type>>2 == 18 {
		if fe_type > 0x3ff {
			var fe_class_name *ZendString = ResolveClassName(fe.GetScope(), (*ZendString)(fe_type & ^0x3))
			var fe_ce *ZendClassEntry = LookupClass(fe.GetScope(), fe_class_name)
			if fe_ce == nil {
				*unresolved_class = fe_class_name
				return INHERITANCE_UNRESOLVED
			}
			if UnlinkedInstanceof(fe_ce, ZendCeTraversable) != 0 {
				return INHERITANCE_SUCCESS
			} else {
				return INHERITANCE_ERROR
			}
		}
		if fe_type>>2 == 18 || fe_type>>2 == 7 {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else if proto_type>>2 == 8 {
		if fe_type > 0x3ff {

			/* Currently, any class name would be allowed here. We still perform a class lookup
			 * for forward-compatibility reasons, as we may have named types in the future that
			 * are not classes (such as enums or typedefs). */

			var fe_class_name *ZendString = ResolveClassName(fe.GetScope(), (*ZendString)(fe_type & ^0x3))
			var fe_ce *ZendClassEntry = LookupClass(fe.GetScope(), fe_class_name)
			if fe_ce == nil {
				*unresolved_class = fe_class_name
				return INHERITANCE_UNRESOLVED
			}
			return INHERITANCE_SUCCESS
		}
		if fe_type>>2 == 8 {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else {
		if fe_type>>2 == proto_type>>2 {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	}
}

/* }}} */

func ZendDoPerformArgTypeHintCheck(unresolved_class **ZendString, fe *ZendFunction, fe_arg_info *ZendArgInfo, proto *ZendFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	if fe_arg_info.GetType() <= 0x3 {

		/* Child with no type is always compatible */

		return INHERITANCE_SUCCESS

		/* Child with no type is always compatible */

	}
	if proto_arg_info.GetType() <= 0x3 {

		/* Child defines a type, but parent doesn't, violates LSP */

		return INHERITANCE_ERROR

		/* Child defines a type, but parent doesn't, violates LSP */

	}

	/* Contravariant type check is performed as a covariant type check with swapped
	 * argument order. */

	return ZendPerformCovariantTypeCheck(unresolved_class, proto, proto_arg_info, fe, fe_arg_info)

	/* Contravariant type check is performed as a covariant type check with swapped
	 * argument order. */
}

/* }}} */

func ZendDoPerformImplementationCheck(unresolved_class **ZendString, fe *ZendFunction, proto *ZendFunction) InheritanceStatus {
	var i uint32
	var num_args uint32
	var status InheritanceStatus
	var local_status InheritanceStatus

	/* If it's a user function then arg_info == NULL means we don't have any parameters but
	 * we still need to do the arg number checks.  We are only willing to ignore this for internal
	 * functions because extensions don't always define arg_info.
	 */

	if proto.GetArgInfo() == nil && proto.GetCommonType() != 2 {
		return INHERITANCE_SUCCESS
	}

	/* Checks for constructors only if they are declared in an interface,
	 * or explicitly marked as abstract
	 */

	assert(!((fe.GetFnFlags()&1<<28) != 0 && ((proto.GetScope().GetCeFlags()&1<<0) == 0 && (proto.GetFnFlags()&1<<6) == 0)))

	/* If the prototype method is private do not enforce a signature */

	assert((proto.GetFnFlags() & 1 << 2) == 0)

	/* check number of arguments */

	if proto.GetRequiredNumArgs() < fe.GetRequiredNumArgs() || proto.GetNumArgs() > fe.GetNumArgs() {
		return INHERITANCE_ERROR
	}

	/* by-ref constraints on return values are covariant */

	if (proto.GetFnFlags()&1<<12) != 0 && (fe.GetFnFlags()&1<<12) == 0 {
		return INHERITANCE_ERROR
	}
	if (proto.GetFnFlags()&1<<14) != 0 && (fe.GetFnFlags()&1<<14) == 0 {
		return INHERITANCE_ERROR
	}

	/* For variadic functions any additional (optional) arguments that were added must be
	 * checked against the signature of the variadic argument, so in this case we have to
	 * go through all the parameters of the function and not just those present in the
	 * prototype. */

	num_args = proto.GetNumArgs()
	if (proto.GetFnFlags() & 1 << 14) != 0 {
		num_args++
		if fe.GetNumArgs() >= proto.GetNumArgs() {
			num_args = fe.GetNumArgs()
			if (fe.GetFnFlags() & 1 << 14) != 0 {
				num_args++
			}
		}
	}
	status = INHERITANCE_SUCCESS
	for i = 0; i < num_args; i++ {
		var fe_arg_info *ZendArgInfo = &fe.common.arg_info[i]
		var proto_arg_info *ZendArgInfo
		if i < proto.GetNumArgs() {
			proto_arg_info = &proto.common.arg_info[i]
		} else {
			proto_arg_info = &proto.common.arg_info[proto.GetNumArgs()]
		}
		local_status = ZendDoPerformArgTypeHintCheck(unresolved_class, fe, fe_arg_info, proto, proto_arg_info)
		if local_status != INHERITANCE_SUCCESS {
			if local_status == INHERITANCE_ERROR {
				return INHERITANCE_ERROR
			}
			assert(local_status == INHERITANCE_UNRESOLVED)
			status = INHERITANCE_UNRESOLVED
		}

		/* by-ref constraints on arguments are invariant */

		if fe_arg_info.GetPassByReference() != proto_arg_info.GetPassByReference() {
			return INHERITANCE_ERROR
		}

		/* by-ref constraints on arguments are invariant */

	}

	/* Check return type compatibility, but only if the prototype already specifies
	 * a return type. Adding a new return type is always valid. */

	if (proto.GetFnFlags() & 1 << 13) != 0 {

		/* Removing a return type is not valid. */

		if (fe.GetFnFlags() & 1 << 13) == 0 {
			return INHERITANCE_ERROR
		}
		local_status = ZendPerformCovariantTypeCheck(unresolved_class, fe, fe.GetArgInfo()-1, proto, proto.GetArgInfo()-1)
		if local_status != INHERITANCE_SUCCESS {
			if local_status == INHERITANCE_ERROR {
				return INHERITANCE_ERROR
			}
			assert(local_status == INHERITANCE_UNRESOLVED)
			status = INHERITANCE_UNRESOLVED
		}
	}
	return status
}

/* }}} */

func ZendAppendTypeHint(str *SmartStr, fptr *ZendFunction, arg_info *ZendArgInfo, return_hint int) {
	if arg_info.GetType() > 0x3 && (arg_info.GetType()&0x1) != 0 {
		SmartStrAppendcEx(str, '?', 0)
	}
	if arg_info.GetType() > 0x3ff {
		var class_name *byte
		var class_name_len int
		class_name = (*ZendString)(arg_info.GetType() & ^0x3).GetVal()
		class_name_len = (*ZendString)(arg_info.GetType() & ^0x3).GetLen()
		if !(strcasecmp(class_name, "self")) && fptr.GetScope() != nil {
			class_name = fptr.GetScope().GetName().GetVal()
			class_name_len = fptr.GetScope().GetName().GetLen()
		} else if !(strcasecmp(class_name, "parent")) && fptr.GetScope() != nil && fptr.GetScope().parent {
			class_name = fptr.GetScope().parent.name.val
			class_name_len = fptr.GetScope().parent.name.len_
		}
		SmartStrAppendlEx(str, class_name, class_name_len, 0)
		if return_hint == 0 {
			SmartStrAppendcEx(str, ' ', 0)
		}
	} else if arg_info.GetType() > 0x3 && arg_info.GetType() <= 0x3ff {
		var type_name *byte = ZendGetTypeByConst(arg_info.GetType() >> 2)
		SmartStrAppendlEx(str, type_name, strlen(type_name), 0)
		if return_hint == 0 {
			SmartStrAppendcEx(str, ' ', 0)
		}
	}
}

/* }}} */

func ZendGetFunctionDeclaration(fptr *ZendFunction) *ZendString {
	var str SmartStr = SmartStr{0}
	if (fptr.GetOpArray().GetFnFlags() & 1 << 12) != 0 {
		SmartStrAppendlEx(&str, "& ", strlen("& "), 0)
	}
	if fptr.GetScope() != nil {

		/* cut off on NULL byte ... class@anonymous */

		SmartStrAppendlEx(&str, fptr.GetScope().GetName().GetVal(), strlen(fptr.GetScope().GetName().GetVal()), 0)
		SmartStrAppendlEx(&str, "::", strlen("::"), 0)
	}
	SmartStrAppendEx(&str, fptr.GetFunctionName(), 0)
	SmartStrAppendcEx(&str, '(', 0)
	if fptr.GetArgInfo() != nil {
		var i uint32
		var num_args uint32
		var required uint32
		var arg_info *ZendArgInfo = fptr.GetArgInfo()
		required = fptr.GetRequiredNumArgs()
		num_args = fptr.GetNumArgs()
		if (fptr.GetFnFlags() & 1 << 14) != 0 {
			num_args++
		}
		for i = 0; i < num_args; {
			ZendAppendTypeHint(&str, fptr, arg_info, 0)
			if arg_info.GetPassByReference() != 0 {
				SmartStrAppendcEx(&str, '&', 0)
			}
			if arg_info.GetIsVariadic() != 0 {
				SmartStrAppendlEx(&str, "...", strlen("..."), 0)
			}
			SmartStrAppendcEx(&str, '$', 0)
			if arg_info.GetName() != nil {
				if fptr.GetType() == 1 {
					SmartStrAppendlEx(&str, (*ZendInternalArgInfo)(arg_info).GetName(), strlen((*ZendInternalArgInfo)(arg_info).GetName()), 0)
				} else {
					SmartStrAppendlEx(&str, arg_info.GetName().GetVal(), arg_info.GetName().GetLen(), 0)
				}
			} else {
				SmartStrAppendlEx(&str, "param", strlen("param"), 0)
				SmartStrAppendUnsignedEx(&str, i, 0)
			}
			if i >= required && arg_info.GetIsVariadic() == 0 {
				SmartStrAppendlEx(&str, " = ", strlen(" = "), 0)
				if fptr.GetType() == 2 {
					var precv *ZendOp = nil
					var idx uint32 = i
					var op *ZendOp = fptr.GetOpArray().GetOpcodes()
					var end *ZendOp = op + fptr.GetOpArray().GetLast()
					idx++
					for op < end {
						if (op.GetOpcode() == 63 || op.GetOpcode() == 64) && op.GetOp1().GetNum() == ZendUlong(idx) {
							precv = op
						}
						op++
					}
					if precv != nil && precv.GetOpcode() == 64 && precv.GetOp2Type() != 0 {
						var zv *Zval = (*Zval)((*byte)(precv) + int32(precv.GetOp2()).constant)
						if zv.GetType() == 2 {
							SmartStrAppendlEx(&str, "false", strlen("false"), 0)
						} else if zv.GetType() == 3 {
							SmartStrAppendlEx(&str, "true", strlen("true"), 0)
						} else if zv.GetType() == 1 {
							SmartStrAppendlEx(&str, "NULL", strlen("NULL"), 0)
						} else if zv.GetType() == 6 {
							SmartStrAppendcEx(&str, '\'', 0)
							SmartStrAppendlEx(&str, zv.GetValue().GetStr().GetVal(), g.CondF1(zv.GetValue().GetStr().GetLen() < 10, func() int { return zv.GetValue().GetStr().GetLen() }, 10), 0)
							if zv.GetValue().GetStr().GetLen() > 10 {
								SmartStrAppendlEx(&str, "...", strlen("..."), 0)
							}
							SmartStrAppendcEx(&str, '\'', 0)
						} else if zv.GetType() == 7 {
							SmartStrAppendlEx(&str, "Array", strlen("Array"), 0)
						} else if zv.GetType() == 11 {
							var ast *ZendAst = (*ZendAst)((*byte)(zv.GetValue().GetAst()) + g.SizeOf("zend_ast_ref"))
							if ast.GetKind() == ZEND_AST_CONSTANT {
								SmartStrAppendEx(&str, ZendAstGetConstantName(ast), 0)
							} else {
								SmartStrAppendlEx(&str, "<expression>", strlen("<expression>"), 0)
							}
						} else {
							var tmp_zv_str *ZendString
							var zv_str *ZendString = ZvalGetTmpString(zv, &tmp_zv_str)
							SmartStrAppendEx(&str, zv_str, 0)
							ZendTmpStringRelease(tmp_zv_str)
						}
					}
				} else {
					SmartStrAppendlEx(&str, "NULL", strlen("NULL"), 0)
				}
			}
			if g.PreInc(&i) < num_args {
				SmartStrAppendlEx(&str, ", ", strlen(", "), 0)
			}
			arg_info++
		}
	}
	SmartStrAppendcEx(&str, ')', 0)
	if (fptr.GetFnFlags() & 1 << 13) != 0 {
		SmartStrAppendlEx(&str, ": ", strlen(": "), 0)
		ZendAppendTypeHint(&str, fptr, fptr.GetArgInfo()-1, 1)
	}
	SmartStr0(&str)
	return str.GetS()
}

/* }}} */

func FuncLineno(fn *ZendFunction) uint32 {
	if fn.GetCommonType() == 2 {
		return fn.GetOpArray().GetLineStart()
	} else {
		return 0
	}
}
func EmitIncompatibleMethodError(error_level int, error_verb *byte, child *ZendFunction, parent *ZendFunction, status InheritanceStatus, unresolved_class *ZendString) {
	var parent_prototype *ZendString = ZendGetFunctionDeclaration(parent)
	var child_prototype *ZendString = ZendGetFunctionDeclaration(child)
	if status == INHERITANCE_UNRESOLVED {
		ZendErrorAt(error_level, nil, FuncLineno(child), "Could not check compatibility between %s and %s, because class %s is not available", child_prototype.GetVal(), parent_prototype.GetVal(), unresolved_class.GetVal())
	} else {
		ZendErrorAt(error_level, nil, FuncLineno(child), "Declaration of %s %s be compatible with %s", child_prototype.GetVal(), error_verb, parent_prototype.GetVal())
	}
	ZendStringEfree(child_prototype)
	ZendStringEfree(parent_prototype)
}
func EmitIncompatibleMethodErrorOrWarning(child *ZendFunction, parent *ZendFunction, status InheritanceStatus, unresolved_class *ZendString, always_error ZendBool) {
	var error_level int
	var error_verb *byte
	if always_error != 0 || child.GetPrototype() != nil && (child.GetPrototype().GetFnFlags()&1<<6) != 0 || (parent.GetFnFlags()&1<<13) != 0 && ((child.GetFnFlags()&1<<13) == 0 || ZendPerformCovariantTypeCheck(&unresolved_class, child, child.GetArgInfo()-1, parent, parent.GetArgInfo()-1) != INHERITANCE_SUCCESS) {
		error_level = 1 << 6
		error_verb = "must"
	} else {
		error_level = 1 << 1
		error_verb = "should"
	}
	EmitIncompatibleMethodError(error_level, error_verb, child, parent, status, unresolved_class)
}
func PerformDelayableImplementationCheck(ce *ZendClassEntry, fe *ZendFunction, proto *ZendFunction, always_error ZendBool) {
	var unresolved_class *ZendString
	var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, fe, proto)
	if status != INHERITANCE_SUCCESS {
		if status == INHERITANCE_UNRESOLVED {
			AddCompatibilityObligation(ce, fe, proto, always_error)
		} else {
			assert(status == INHERITANCE_ERROR)
			if always_error != 0 {
				EmitIncompatibleMethodError(1<<6, "must", fe, proto, status, unresolved_class)
			} else {
				EmitIncompatibleMethodErrorOrWarning(fe, proto, status, unresolved_class, always_error)
			}
		}
	}
}
func DoInheritanceCheckOnMethodEx(child *ZendFunction, parent *ZendFunction, ce *ZendClassEntry, child_zv *Zval, check_only ZendBool, checked ZendBool) InheritanceStatus {
	var child_flags uint32
	var parent_flags uint32 = parent.GetFnFlags()
	var proto *ZendFunction
	if checked == 0 && (parent_flags&1<<5) != 0 {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(1<<6, nil, FuncLineno(child), "Cannot override final method %s::%s()", g.CondF1(parent != nil && parent.GetScope() != nil, func() []byte { return parent.GetScope().GetName().GetVal() }, ""), child.GetFunctionName().GetVal())
	}
	child_flags = child.GetFnFlags()

	/* You cannot change from static to non static and vice versa.
	 */

	if checked == 0 && (child_flags&1<<4) != (parent_flags&1<<4) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		if (child_flags & 1 << 4) != 0 {
			ZendErrorAtNoreturn(1<<6, nil, FuncLineno(child), "Cannot make non static method %s::%s() static in class %s", g.CondF1(parent != nil && parent.GetScope() != nil, func() []byte { return parent.GetScope().GetName().GetVal() }, ""), child.GetFunctionName().GetVal(), g.CondF1(child != nil && child.GetScope() != nil, func() []byte { return child.GetScope().GetName().GetVal() }, ""))
		} else {
			ZendErrorAtNoreturn(1<<6, nil, FuncLineno(child), "Cannot make static method %s::%s() non static in class %s", g.CondF1(parent != nil && parent.GetScope() != nil, func() []byte { return parent.GetScope().GetName().GetVal() }, ""), child.GetFunctionName().GetVal(), g.CondF1(child != nil && child.GetScope() != nil, func() []byte { return child.GetScope().GetName().GetVal() }, ""))
		}
	}

	/* Disallow making an inherited method abstract. */

	if checked == 0 && (child_flags&1<<6) > (parent_flags&1<<6) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(1<<6, nil, FuncLineno(child), "Cannot make non abstract method %s::%s() abstract in class %s", g.CondF1(parent != nil && parent.GetScope() != nil, func() []byte { return parent.GetScope().GetName().GetVal() }, ""), child.GetFunctionName().GetVal(), g.CondF1(child != nil && child.GetScope() != nil, func() []byte { return child.GetScope().GetName().GetVal() }, ""))
	}
	if check_only == 0 && (parent_flags&(1<<2|1<<3)) != 0 {
		child.SetFnFlags(child.GetFnFlags() | 1<<3)
	}
	if (parent_flags & 1 << 2) != 0 {
		return INHERITANCE_SUCCESS
	}
	if parent.GetPrototype() != nil {
		proto = parent.GetPrototype()
	} else {
		proto = parent
	}
	if (parent_flags & 1 << 28) != 0 {

		/* ctors only have a prototype if is abstract (or comes from an interface) */

		if (proto.GetFnFlags() & 1 << 6) == 0 {
			return INHERITANCE_SUCCESS
		}
		parent = proto
	}
	if check_only == 0 && child.GetPrototype() != proto {
		for {
			if child.GetScope() != ce && child.GetType() == 2 && child.GetOpArray().GetStaticVariables() == nil {
				if (ce.GetCeFlags() & 1 << 0) != 0 {

					/* Few parent interfaces contain the same method */

					break

					/* Few parent interfaces contain the same method */

				} else if child_zv != nil {

					/* op_array wasn't duplicated yet */

					var new_function *ZendFunction = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_op_array"))
					memcpy(new_function, child, g.SizeOf("zend_op_array"))
					child = new_function
					child_zv.GetValue().SetPtr(child)
				}
			}
			child.SetPrototype(proto)
			break
		}
	}

	/* Prevent derived classes from restricting access that was available in parent classes (except deriving from non-abstract ctors) */

	if checked == 0 && (child_flags&(1<<0|1<<1|1<<2)) > (parent_flags&(1<<0|1<<1|1<<2)) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(1<<6, nil, FuncLineno(child), "Access level to %s::%s() must be %s (as in class %s)%s", g.CondF1(child != nil && child.GetScope() != nil, func() []byte { return child.GetScope().GetName().GetVal() }, ""), child.GetFunctionName().GetVal(), ZendVisibilityString(parent_flags), g.CondF1(parent != nil && parent.GetScope() != nil, func() []byte { return parent.GetScope().GetName().GetVal() }, ""), g.Cond((parent_flags&1<<0) != 0, "", " or weaker"))
	}
	if checked == 0 {
		if check_only != 0 {
			var unresolved_class *ZendString
			return ZendDoPerformImplementationCheck(&unresolved_class, child, parent)
		}
		PerformDelayableImplementationCheck(ce, child, parent, 0)
	}
	return INHERITANCE_SUCCESS
}

/* }}} */

func DoInheritanceCheckOnMethod(child *ZendFunction, parent *ZendFunction, ce *ZendClassEntry, child_zv *Zval) {
	DoInheritanceCheckOnMethodEx(child, parent, ce, child_zv, 0, 0)
}

/* }}} */

func DoInheritMethod(key *ZendString, parent *ZendFunction, ce *ZendClassEntry, is_interface ZendBool, checked ZendBool) {
	var child *Zval = ZendHashFindEx(&ce.function_table, key, 1)
	if child != nil {
		var func_ *ZendFunction = (*ZendFunction)(child.GetValue().GetPtr())
		if is_interface != 0 && func_ == parent {

			/* The same method in interface may be inherited few times */

			return

			/* The same method in interface may be inherited few times */

		}
		if checked != 0 {
			DoInheritanceCheckOnMethodEx(func_, parent, ce, child, 0, checked)
		} else {
			DoInheritanceCheckOnMethod(func_, parent, ce, child)
		}
	} else {
		if is_interface != 0 || (parent.GetFnFlags()&1<<6) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | 1<<4)
		}
		parent = ZendDuplicateFunction(parent, ce, is_interface)
		if is_interface == 0 {
			_zendHashAppendPtr(&ce.function_table, key, parent)
		} else {
			ZendHashAddNewPtr(&ce.function_table, key, parent)
		}
	}
}

/* }}} */

func PropertyTypesCompatible(parent_info *ZendPropertyInfo, child_info *ZendPropertyInfo) InheritanceStatus {
	var parent_name *ZendString
	var child_name *ZendString
	var parent_type_ce *ZendClassEntry
	var child_type_ce *ZendClassEntry
	if parent_info.GetType() == child_info.GetType() {
		return INHERITANCE_SUCCESS
	}
	if parent_info.GetType() <= 0x3ff || child_info.GetType() <= 0x3ff || (parent_info.GetType()&0x1) != 0 != ((child_info.GetType()&0x1) != 0) {
		return INHERITANCE_ERROR
	}
	if (parent_info.GetType() & 0x2) != 0 {
		parent_name = (*ZendClassEntry)(parent_info.GetType() & ^0x3).GetName()
	} else {
		parent_name = ResolveClassName(parent_info.GetCe(), (*ZendString)(parent_info.GetType() & ^0x3))
	}
	if (child_info.GetType() & 0x2) != 0 {
		child_name = (*ZendClassEntry)(child_info.GetType() & ^0x3).GetName()
	} else {
		child_name = ResolveClassName(child_info.GetCe(), (*ZendString)(child_info.GetType() & ^0x3))
	}
	if parent_name.GetLen() == child_name.GetLen() && ZendBinaryStrcasecmp(parent_name.GetVal(), parent_name.GetLen(), child_name.GetVal(), child_name.GetLen()) == 0 {
		return INHERITANCE_SUCCESS
	}

	/* Check for class aliases */

	if (parent_info.GetType() & 0x2) != 0 {
		parent_type_ce = (*ZendClassEntry)(parent_info.GetType() & ^0x3)
	} else {
		parent_type_ce = LookupClass(parent_info.GetCe(), parent_name)
	}
	if (child_info.GetType() & 0x2) != 0 {
		child_type_ce = (*ZendClassEntry)(child_info.GetType() & ^0x3)
	} else {
		child_type_ce = LookupClass(child_info.GetCe(), child_name)
	}
	if parent_type_ce == nil || child_type_ce == nil {
		return INHERITANCE_UNRESOLVED
	}
	if parent_type_ce == child_type_ce {
		return INHERITANCE_SUCCESS
	} else {
		return INHERITANCE_ERROR
	}
}
func EmitIncompatiblePropertyError(child *ZendPropertyInfo, parent *ZendPropertyInfo) {
	ZendErrorNoreturn(1<<6, "Type of %s::$%s must be %s%s (as in class %s)", child.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(child.GetName()), g.Cond((parent.GetType()&0x1) != 0, "?", ""), g.CondF(parent.GetType() > 0x3ff, func() __auto__ {
		return g.CondF((parent.GetType()&0x2) != 0, func() *ZendString { return (*ZendClassEntry)(parent.GetType() & ^0x3).GetName() }, func() *ZendString { return ResolveClassName(parent.GetCe(), (*ZendString)(parent.GetType() & ^0x3)) }).val
	}, func() *byte { return ZendGetTypeByConst(parent.GetType() >> 2) }), parent.GetCe().GetName().GetVal())
}
func DoInheritProperty(parent_info *ZendPropertyInfo, key *ZendString, ce *ZendClassEntry) {
	var child *Zval = ZendHashFindEx(&ce.properties_info, key, 1)
	var child_info *ZendPropertyInfo
	if child != nil {
		child_info = child.GetValue().GetPtr()
		if (parent_info.GetFlags() & (1<<2 | 1<<3)) != 0 {
			child_info.SetFlags(child_info.GetFlags() | 1<<3)
		}
		if (parent_info.GetFlags() & 1 << 2) == 0 {
			if (parent_info.GetFlags() & 1 << 4) != (child_info.GetFlags() & 1 << 4) {
				ZendErrorNoreturn(1<<6, "Cannot redeclare %s%s::$%s as %s%s::$%s", g.Cond((parent_info.GetFlags()&1<<4) != 0, "static ", "non static "), ce.parent.name.val, key.GetVal(), g.Cond((child_info.GetFlags()&1<<4) != 0, "static ", "non static "), ce.GetName().GetVal(), key.GetVal())
			}
			if (child_info.GetFlags() & (1<<0 | 1<<1 | 1<<2)) > (parent_info.GetFlags() & (1<<0 | 1<<1 | 1<<2)) {
				ZendErrorNoreturn(1<<6, "Access level to %s::$%s must be %s (as in class %s)%s", ce.GetName().GetVal(), key.GetVal(), ZendVisibilityString(parent_info.GetFlags()), ce.parent.name.val, g.Cond((parent_info.GetFlags()&1<<0) != 0, "", " or weaker"))
			} else if (child_info.GetFlags() & 1 << 4) == 0 {
				var parent_num int = (parent_info.GetOffset() - uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0)) / g.SizeOf("zval")
				var child_num int = (child_info.GetOffset() - uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0)) / g.SizeOf("zval")

				/* Don't keep default properties in GC (they may be freed by opcache) */

				ZvalPtrDtorNogc(&ce.GetDefaultPropertiesTable()[parent_num])
				ce.GetDefaultPropertiesTable()[parent_num] = ce.GetDefaultPropertiesTable()[child_num]
				&ce.default_properties_table[child_num].u1.type_info = 0
				child_info.SetOffset(parent_info.GetOffset())
			}
			if parent_info.GetType() > 0x3 {
				var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
				if status == INHERITANCE_ERROR {
					EmitIncompatiblePropertyError(child_info, parent_info)
				}
				if status == INHERITANCE_UNRESOLVED {
					AddPropertyCompatibilityObligation(ce, child_info, parent_info)
				}
			} else if child_info.GetType() > 0x3 && parent_info.GetType() <= 0x3 {
				ZendErrorNoreturn(1<<6, "Type of %s::$%s must not be defined (as in class %s)", ce.GetName().GetVal(), key.GetVal(), ce.parent.name.val)
			}
		}
	} else {
		if (ce.GetType() & 1) != 0 {
			child_info = ZendDuplicatePropertyInfoInternal(parent_info)
		} else {
			child_info = parent_info
		}
		_zendHashAppendPtr(&ce.properties_info, key, child_info)
	}
}

/* }}} */

func DoImplementInterface(ce *ZendClassEntry, iface *ZendClassEntry) {
	if (ce.GetCeFlags()&1<<0) == 0 && iface.interface_gets_implemented && iface.interface_gets_implemented(iface, ce) == FAILURE {
		ZendErrorNoreturn(1<<4, "Class %s could not implement interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
	}

	/* This should be prevented by the class lookup logic. */

	assert(ce != iface)

	/* This should be prevented by the class lookup logic. */
}

/* }}} */

func ZendDoInheritInterfaces(ce *ZendClassEntry, iface *ZendClassEntry) {
	/* expects interface to be contained in ce's interface list already */

	var i uint32
	var ce_num uint32
	var if_num uint32 = iface.GetNumInterfaces()
	var entry *ZendClassEntry
	ce_num = ce.GetNumInterfaces()
	if ce.GetType() == 1 {
		ce.interfaces = (**ZendClassEntry)(realloc(ce.interfaces, g.SizeOf("zend_class_entry *")*(ce_num+if_num)))
	} else {
		ce.interfaces = (**ZendClassEntry)(_erealloc(ce.interfaces, g.SizeOf("zend_class_entry *")*(ce_num+if_num)))
	}

	/* Inherit the interfaces, only if they're not already inherited by the class */

	for g.PostDec(&if_num) {
		entry = iface.interfaces[if_num]
		for i = 0; i < ce_num; i++ {
			if ce.interfaces[i] == entry {
				break
			}
		}
		if i == ce_num {
			ce.interfaces[g.PostInc(&(ce.GetNumInterfaces()))] = entry
		}
	}
	ce.SetCeFlags(ce.GetCeFlags() | 1<<20)

	/* and now call the implementing handlers */

	for ce_num < ce.GetNumInterfaces() {
		DoImplementInterface(ce, ce.interfaces[g.PostInc(&ce_num)])
	}

	/* and now call the implementing handlers */
}

/* }}} */

func DoInheritClassConstant(name *ZendString, parent_const *ZendClassConstant, ce *ZendClassEntry) {
	var zv *Zval = ZendHashFindEx(&ce.constants_table, name, 1)
	var c *ZendClassConstant
	if zv != nil {
		c = (*ZendClassConstant)(zv.GetValue().GetPtr())
		if (c.GetValue().GetAccessFlags() & (1<<0 | 1<<1 | 1<<2)) > (parent_const.GetValue().GetAccessFlags() & (1<<0 | 1<<1 | 1<<2)) {
			ZendErrorNoreturn(1<<6, "Access level to %s::%s must be %s (as in class %s)%s", ce.GetName().GetVal(), name.GetVal(), ZendVisibilityString(parent_const.GetValue().GetAccessFlags()), ce.parent.name.val, g.Cond((parent_const.GetValue().GetAccessFlags()&1<<0) != 0, "", " or weaker"))
		}
	} else if (parent_const.GetValue().GetAccessFlags() & 1 << 2) == 0 {
		if parent_const.GetValue().GetType() == 11 {
			ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
		}
		if (ce.GetType() & 1) != 0 {
			c = __zendMalloc(g.SizeOf("zend_class_constant"))
			memcpy(c, parent_const, g.SizeOf("zend_class_constant"))
			parent_const = c
		}
		_zendHashAppendPtr(&ce.constants_table, name, parent_const)
	}
}

/* }}} */

func ZendBuildPropertiesInfoTable(ce *ZendClassEntry) {
	var table **ZendPropertyInfo
	var prop **ZendPropertyInfo
	var size int
	if ce.GetDefaultPropertiesCount() == 0 {
		return
	}
	assert(ce.GetPropertiesInfoTable() == nil)
	size = g.SizeOf("zend_property_info *") * ce.GetDefaultPropertiesCount()
	if ce.GetType() == 2 {
		table = ZendArenaAlloc(&CG.arena, size)
		ce.SetPropertiesInfoTable(table)
	} else {
		table = __zendMalloc(size)
		ce.SetPropertiesInfoTable(table)
	}

	/* Dead slots may be left behind during inheritance. Make sure these are NULLed out. */

	memset(table, 0, size)
	if ce.parent && ce.parent.default_properties_count != 0 {
		var parent_table **ZendPropertyInfo = ce.parent.properties_info_table
		memcpy(table, parent_table, g.SizeOf("zend_property_info *")*ce.parent.default_properties_count)

		/* Child did not add any new properties, we are done */

		if ce.GetDefaultPropertiesCount() == ce.parent.default_properties_count {
			return
		}

		/* Child did not add any new properties, we are done */

	}
	for {
		var __ht *HashTable = &ce.properties_info
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			prop = _z.GetValue().GetPtr()
			if prop.ce == ce && (prop.flags&1<<4) == 0 {
				table[(prop.offset-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")] = prop
			}
		}
		break
	}
}
func ZendDoInheritanceEx(ce *ZendClassEntry, parent_ce *ZendClassEntry, checked ZendBool) {
	var property_info *ZendPropertyInfo
	var func_ *ZendFunction
	var key *ZendString
	if (ce.GetCeFlags() & 1 << 0) != 0 {

		/* Interface can only inherit other interfaces */

		if (parent_ce.GetCeFlags() & 1 << 0) == 0 {
			ZendErrorNoreturn(1<<6, "Interface %s may not inherit from class (%s)", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Interface can only inherit other interfaces */

	} else if (parent_ce.GetCeFlags() & (1<<0 | 1<<1 | 1<<5)) != 0 {

		/* Class declaration must not extend traits or interfaces */

		if (parent_ce.GetCeFlags() & 1 << 0) != 0 {
			ZendErrorNoreturn(1<<6, "Class %s cannot extend from interface %s", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		} else if (parent_ce.GetCeFlags() & 1 << 1) != 0 {
			ZendErrorNoreturn(1<<6, "Class %s cannot extend from trait %s", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Class must not extend a final class */

		if (parent_ce.GetCeFlags() & 1 << 5) != 0 {
			ZendErrorNoreturn(1<<6, "Class %s may not inherit from final class (%s)", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Class must not extend a final class */

	}
	if ce.parent_name {
		ZendStringReleaseEx(ce.parent_name, 0)
	}
	ce.parent = parent_ce
	ce.SetCeFlags(ce.GetCeFlags() | 1<<19)

	/* Inherit interfaces */

	if parent_ce.GetNumInterfaces() != 0 {
		if (ce.GetCeFlags() & 1 << 14) == 0 {
			ZendDoInheritInterfaces(ce, parent_ce)
		} else {
			var i uint32
			for i = 0; i < parent_ce.GetNumInterfaces(); i++ {
				DoImplementInterface(ce, parent_ce.interfaces[i])
			}
		}
	}

	/* Inherit properties */

	if parent_ce.GetDefaultPropertiesCount() != 0 {
		var src *Zval
		var dst *Zval
		var end *Zval
		if ce.GetDefaultPropertiesCount() != 0 {
			var table *Zval = g.CondF(ce.GetType() == 1, func() any {
				return __zendMalloc(g.SizeOf("zval") * (ce.GetDefaultPropertiesCount() + parent_ce.GetDefaultPropertiesCount()))
			}, func() any {
				return _emalloc(g.SizeOf("zval") * (ce.GetDefaultPropertiesCount() + parent_ce.GetDefaultPropertiesCount()))
			})
			src = ce.GetDefaultPropertiesTable() + ce.GetDefaultPropertiesCount()
			end = table + parent_ce.GetDefaultPropertiesCount()
			dst = end + ce.GetDefaultPropertiesCount()
			ce.SetDefaultPropertiesTable(table)
			for {
				dst--
				src--
				*dst = *src
				if dst == end {
					break
				}
			}
			g.CondF(ce.GetType() == 1, func() { return Free(src) }, func() { return _efree(src) })
			end = ce.GetDefaultPropertiesTable()
		} else {
			if ce.GetType() == 1 {
				end = __zendMalloc(g.SizeOf("zval") * parent_ce.GetDefaultPropertiesCount())
			} else {
				end = _emalloc(g.SizeOf("zval") * parent_ce.GetDefaultPropertiesCount())
			}
			dst = end + parent_ce.GetDefaultPropertiesCount()
			ce.SetDefaultPropertiesTable(end)
		}
		src = parent_ce.GetDefaultPropertiesTable() + parent_ce.GetDefaultPropertiesCount()
		if parent_ce.GetType() != ce.GetType() {

			/* User class extends internal */

			for {
				dst--
				src--
				var _z1 *Zval = dst
				var _z2 *Zval = src
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
						ZendGcAddref(&_gc.gc)
					} else {
						ZvalCopyCtorFunc(_z1)
					}
				}
				dst.SetU2Extra(src.GetU2Extra())
				if (dst.GetTypeInfo() & 0xff) == 11 {
					ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
				}
				continue
				if dst == end {
					break
				}
			}

			/* User class extends internal */

		} else {
			for {
				dst--
				src--
				var _z1 *Zval = dst
				var _z2 *Zval = src
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				dst.SetU2Extra(src.GetU2Extra())
				if (dst.GetTypeInfo() & 0xff) == 11 {
					ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
				}
				continue
				if dst == end {
					break
				}
			}
		}
		ce.SetDefaultPropertiesCount(ce.GetDefaultPropertiesCount() + parent_ce.GetDefaultPropertiesCount())
	}
	if parent_ce.GetDefaultStaticMembersCount() != 0 {
		var src *Zval
		var dst *Zval
		var end *Zval
		if ce.GetDefaultStaticMembersCount() != 0 {
			var table *Zval = g.CondF(ce.GetType() == 1, func() any {
				return __zendMalloc(g.SizeOf("zval") * (ce.GetDefaultStaticMembersCount() + parent_ce.GetDefaultStaticMembersCount()))
			}, func() any {
				return _emalloc(g.SizeOf("zval") * (ce.GetDefaultStaticMembersCount() + parent_ce.GetDefaultStaticMembersCount()))
			})
			src = ce.GetDefaultStaticMembersTable() + ce.GetDefaultStaticMembersCount()
			end = table + parent_ce.GetDefaultStaticMembersCount()
			dst = end + ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(table)
			for {
				dst--
				src--
				var _z1 *Zval = dst
				var _z2 *Zval = src
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if dst == end {
					break
				}
			}
			g.CondF(ce.GetType() == 1, func() { return Free(src) }, func() { return _efree(src) })
			end = ce.GetDefaultStaticMembersTable()
		} else {
			if ce.GetType() == 1 {
				end = __zendMalloc(g.SizeOf("zval") * parent_ce.GetDefaultStaticMembersCount())
			} else {
				end = _emalloc(g.SizeOf("zval") * parent_ce.GetDefaultStaticMembersCount())
			}
			dst = end + parent_ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(end)
		}
		if parent_ce.GetType() != ce.GetType() {

			/* User class extends internal */

			if (*Zval)(g.CondF((uintptr_t(parent_ce).static_members_table__ptr&1) != 0, func() any {
				return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(parent_ce).static_members_table__ptr - 1)))
			}, func() any { return any(*(parent_ce.GetStaticMembersTablePtr())) })) == nil {
				ZendClassInitStatics(parent_ce)
			}
			if ZendUpdateClassConstants(parent_ce) != SUCCESS {
				assert(false)
			}
			src = (*Zval)(g.CondF((uintptr_t(parent_ce).static_members_table__ptr&1) != 0, func() any {
				return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(parent_ce).static_members_table__ptr - 1)))
			}, func() any { return any(*(parent_ce.GetStaticMembersTablePtr())) })) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.GetType() == 13 {
					dst.GetValue().SetZv(src.GetValue().GetZv())
					dst.SetTypeInfo(13)
				} else {
					dst.GetValue().SetZv(src)
					dst.SetTypeInfo(13)
				}
				if dst == end {
					break
				}
			}
		} else if ce.GetType() == 2 {
			if (*Zval)(g.CondF((uintptr_t(parent_ce).static_members_table__ptr&1) != 0, func() any {
				return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(parent_ce).static_members_table__ptr - 1)))
			}, func() any { return any(*(parent_ce.GetStaticMembersTablePtr())) })) == nil {
				assert((parent_ce.GetCeFlags() & (1<<7 | 1<<10)) != 0)
				ZendClassInitStatics(parent_ce)
			}
			src = (*Zval)(g.CondF((uintptr_t(parent_ce).static_members_table__ptr&1) != 0, func() any {
				return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(parent_ce).static_members_table__ptr - 1)))
			}, func() any { return any(*(parent_ce.GetStaticMembersTablePtr())) })) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.GetType() == 13 {
					dst.GetValue().SetZv(src.GetValue().GetZv())
					dst.SetTypeInfo(13)
				} else {
					dst.GetValue().SetZv(src)
					dst.SetTypeInfo(13)
				}
				if dst.GetValue().GetZv().GetType() == 11 {
					ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
				}
				if dst == end {
					break
				}
			}
		} else {
			src = parent_ce.GetDefaultStaticMembersTable() + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.GetType() == 13 {
					dst.GetValue().SetZv(src.GetValue().GetZv())
					dst.SetTypeInfo(13)
				} else {
					dst.GetValue().SetZv(src)
					dst.SetTypeInfo(13)
				}
				if dst == end {
					break
				}
			}
		}
		ce.SetDefaultStaticMembersCount(ce.GetDefaultStaticMembersCount() + parent_ce.GetDefaultStaticMembersCount())
		if ce.GetStaticMembersTablePtr() == nil {
			assert(ce.GetType() == 1)
			if EG.GetCurrentExecuteData() == nil {
				ce.SetStaticMembersTablePtr(ZendMapPtrNew())
			} else {

				/* internal class loaded by dl() */

				ce.SetStaticMembersTablePtr(&ce.default_static_members_table)

				/* internal class loaded by dl() */

			}
		}
	}
	for {
		var __ht *HashTable = &ce.properties_info
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			property_info = _z.GetValue().GetPtr()
			if property_info.GetCe() == ce {
				if (property_info.GetFlags() & 1 << 4) != 0 {
					property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultStaticMembersCount())
				} else {
					property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultPropertiesCount()*g.SizeOf("zval"))
				}
			}
		}
		break
	}
	if &parent_ce.properties_info.nNumOfElements {
		ZendHashExtend(&ce.properties_info, &ce.properties_info.nNumOfElements+&parent_ce.properties_info.nNumOfElements, 0)
		for {
			var __ht *HashTable = &parent_ce.properties_info
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				property_info = _z.GetValue().GetPtr()
				DoInheritProperty(property_info, key, ce)
			}
			break
		}
	}
	if &parent_ce.constants_table.nNumOfElements {
		var c *ZendClassConstant
		ZendHashExtend(&ce.constants_table, &ce.constants_table.nNumOfElements+&parent_ce.constants_table.nNumOfElements, 0)
		for {
			var __ht *HashTable = &parent_ce.constants_table
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				c = _z.GetValue().GetPtr()
				DoInheritClassConstant(key, c, ce)
			}
			break
		}
	}
	if &parent_ce.function_table.nNumOfElements {
		ZendHashExtend(&ce.function_table, &ce.function_table.nNumOfElements+&parent_ce.function_table.nNumOfElements, 0)
		if checked != 0 {
			for {
				var __ht *HashTable = &parent_ce.function_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					key = _p.GetKey()
					func_ = _z.GetValue().GetPtr()
					DoInheritMethod(key, func_, ce, 0, 1)
				}
				break
			}
		} else {
			for {
				var __ht *HashTable = &parent_ce.function_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					key = _p.GetKey()
					func_ = _z.GetValue().GetPtr()
					DoInheritMethod(key, func_, ce, 0, 0)
				}
				break
			}
		}
	}
	DoInheritParentConstructor(ce)
	if ce.GetType() == 1 {
		if (ce.GetCeFlags() & 1 << 4) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | 1<<6)
		}
	}
	ce.SetCeFlags(ce.GetCeFlags() | parent_ce.GetCeFlags()&(1<<16|1<<8|1<<11))
}

/* }}} */

func DoInheritConstantCheck(child_constants_table *HashTable, parent_constant *ZendClassConstant, name *ZendString, iface *ZendClassEntry) ZendBool {
	var zv *Zval = ZendHashFindEx(child_constants_table, name, 1)
	var old_constant *ZendClassConstant
	if zv != nil {
		old_constant = (*ZendClassConstant)(zv.GetValue().GetPtr())
		if old_constant.GetCe() != parent_constant.GetCe() {
			ZendErrorNoreturn(1<<6, "Cannot inherit previously-inherited or override constant %s from interface %s", name.GetVal(), iface.GetName().GetVal())
		}
		return 0
	}
	return 1
}

/* }}} */

func DoInheritIfaceConstant(name *ZendString, c *ZendClassConstant, ce *ZendClassEntry, iface *ZendClassEntry) {
	if DoInheritConstantCheck(&ce.constants_table, c, name, iface) != 0 {
		var ct *ZendClassConstant
		if c.GetValue().GetType() == 11 {
			ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 12))
		}
		if (ce.GetType() & 1) != 0 {
			ct = __zendMalloc(g.SizeOf("zend_class_constant"))
			memcpy(ct, c, g.SizeOf("zend_class_constant"))
			c = ct
		}
		ZendHashUpdatePtr(&ce.constants_table, name, c)
	}
}

/* }}} */

func DoInterfaceImplementation(ce *ZendClassEntry, iface *ZendClassEntry) {
	var func_ *ZendFunction
	var key *ZendString
	var c *ZendClassConstant
	for {
		var __ht *HashTable = &iface.constants_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			c = _z.GetValue().GetPtr()
			DoInheritIfaceConstant(key, c, ce, iface)
		}
		break
	}
	for {
		var __ht *HashTable = &iface.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			func_ = _z.GetValue().GetPtr()
			DoInheritMethod(key, func_, ce, 1, 0)
		}
		break
	}
	DoImplementInterface(ce, iface)
	if iface.GetNumInterfaces() != 0 {
		ZendDoInheritInterfaces(ce, iface)
	}
}

/* }}} */

func ZendDoImplementInterface(ce *ZendClassEntry, iface *ZendClassEntry) {
	var i uint32
	var ignore uint32 = 0
	var current_iface_num uint32 = ce.GetNumInterfaces()
	var parent_iface_num uint32 = g.CondF1(ce.parent, func() __auto__ { return ce.parent.num_interfaces }, 0)
	var key *ZendString
	var c *ZendClassConstant
	assert((ce.GetCeFlags() & 1 << 3) != 0)
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		if ce.interfaces[i] == nil {
			memmove(ce.interfaces+i, ce.interfaces+i+1, g.SizeOf("zend_class_entry *")*(g.PreDec(&(ce.GetNumInterfaces()))-i))
			i--
		} else if ce.interfaces[i] == iface {
			if i < parent_iface_num {
				ignore = 1
			} else {
				ZendErrorNoreturn(1<<6, "Class %s cannot implement previously implemented interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
			}
		}
	}
	if ignore != 0 {

		/* Check for attempt to redeclare interface constants */

		for {
			var __ht *HashTable = &ce.constants_table
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				c = _z.GetValue().GetPtr()
				DoInheritConstantCheck(&iface.constants_table, c, key, iface)
			}
			break
		}

		/* Check for attempt to redeclare interface constants */

	} else {
		if ce.GetNumInterfaces() >= current_iface_num {
			if ce.GetType() == 1 {
				ce.interfaces = (**ZendClassEntry)(realloc(ce.interfaces, g.SizeOf("zend_class_entry *")*g.PreInc(&current_iface_num)))
			} else {
				ce.interfaces = (**ZendClassEntry)(_erealloc(ce.interfaces, g.SizeOf("zend_class_entry *")*g.PreInc(&current_iface_num)))
			}
		}
		ce.interfaces[g.PostInc(&(ce.GetNumInterfaces()))] = iface
		DoInterfaceImplementation(ce, iface)
	}
}

/* }}} */

func ZendDoImplementInterfaces(ce *ZendClassEntry, interfaces **ZendClassEntry) {
	var iface *ZendClassEntry
	var num_parent_interfaces uint32 = g.CondF1(ce.parent, func() __auto__ { return ce.parent.num_interfaces }, 0)
	var num_interfaces uint32 = num_parent_interfaces
	var key *ZendString
	var c *ZendClassConstant
	var i uint32
	var j uint32
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		iface = interfaces[num_parent_interfaces+i]
		if (iface.GetCeFlags() & 1 << 3) == 0 {
			AddDependencyObligation(ce, iface)
		}
		if (iface.GetCeFlags() & 1 << 0) == 0 {
			_efree(interfaces)
			ZendErrorNoreturn(1<<0, "%s cannot implement %s - it is not an interface", ce.GetName().GetVal(), iface.GetName().GetVal())
			return
		}
		for j = 0; j < num_interfaces; j++ {
			if interfaces[j] == iface {
				if j >= num_parent_interfaces {
					_efree(interfaces)
					ZendErrorNoreturn(1<<6, "Class %s cannot implement previously implemented interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
					return
				}

				/* skip duplications */

				for {
					var __ht *HashTable = &ce.constants_table
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						key = _p.GetKey()
						c = _z.GetValue().GetPtr()
						DoInheritConstantCheck(&iface.constants_table, c, key, iface)
					}
					break
				}
				iface = nil
				break
			}
		}
		if iface != nil {
			interfaces[num_interfaces] = iface
			num_interfaces++
		}
	}
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		ZendStringReleaseEx(ce.interface_names[i].name, 0)
		ZendStringReleaseEx(ce.interface_names[i].lc_name, 0)
	}
	_efree(ce.interface_names)
	ce.SetNumInterfaces(num_interfaces)
	ce.interfaces = interfaces
	ce.SetCeFlags(ce.GetCeFlags() | 1<<20)
	i = num_parent_interfaces
	for ; i < ce.GetNumInterfaces(); i++ {
		DoInterfaceImplementation(ce, ce.interfaces[i])
	}
}

/* }}} */

func ZendAddMagicMethods(ce *ZendClassEntry, mname *ZendString, fe *ZendFunction) {
	if mname.GetLen() == g.SizeOf("\"serialize\"")-1 && !(memcmp(mname.GetVal(), "serialize", g.SizeOf("\"serialize\"")-1)) {
		ce.SetSerializeFunc(fe)
	} else if mname.GetLen() == g.SizeOf("\"unserialize\"")-1 && !(memcmp(mname.GetVal(), "unserialize", g.SizeOf("\"unserialize\"")-1)) {
		ce.SetUnserializeFunc(fe)
	} else if ce.GetName().GetLen() != mname.GetLen() && (mname.GetVal()[0] != '_' || mname.GetVal()[1] != '_') {

	} else if mname.GetLen() == g.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__clone", g.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) {
		ce.SetClone(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__construct", g.SizeOf("ZEND_CONSTRUCTOR_FUNC_NAME")-1)) {
		if ce.GetConstructor() != nil && (!(ce.parent) || ce.GetConstructor() != ce.parent.constructor) {
			ZendErrorNoreturn(1<<6, "%s has colliding constructor definitions coming from traits", ce.GetName().GetVal())
		}
		ce.SetConstructor(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__destruct", g.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) {
		ce.SetDestructor(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__get", g.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
		ce.SetGet(fe)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
	} else if mname.GetLen() == g.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__set", g.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
		ce.SetSet(fe)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
	} else if mname.GetLen() == g.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__call", g.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
		ce.SetCall(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__unset", g.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
		ce.SetUnset(fe)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
	} else if mname.GetLen() == g.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__isset", g.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
		ce.SetIsset(fe)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<11)
	} else if mname.GetLen() == g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__callstatic", g.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
		ce.SetCallstatic(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__tostring", g.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) {
		ce.SetTostring(fe)
	} else if mname.GetLen() == g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(mname.GetVal(), "__debuginfo", g.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) {
		ce.SetDebugInfo(fe)
	} else if ce.GetName().GetLen() == mname.GetLen() {
		var lowercase_name *ZendString = ZendStringTolowerEx(ce.GetName(), 0)
		lowercase_name = ZendNewInternedString(lowercase_name)
		if !(memcmp(mname.GetVal(), lowercase_name.GetVal(), mname.GetLen())) {
			if ce.GetConstructor() != nil && (!(ce.parent) || ce.GetConstructor() != ce.parent.constructor) {
				ZendErrorNoreturn(1<<6, "%s has colliding constructor definitions coming from traits", ce.GetName().GetVal())
			}
			ce.SetConstructor(fe)
			fe.SetFnFlags(fe.GetFnFlags() | 1<<28)
		}
		ZendStringReleaseEx(lowercase_name, 0)
	}
}

/* }}} */

func ZendAddTraitMethod(ce *ZendClassEntry, name *byte, key *ZendString, fn *ZendFunction, overridden **HashTable) {
	var existing_fn *ZendFunction = nil
	var new_fn *ZendFunction
	if g.Assign(&existing_fn, ZendHashFindPtr(&ce.function_table, key)) != nil {

		/* if it is the same function with the same visibility and has not been assigned a class scope yet, regardless
		 * of where it is coming from there is no conflict and we do not need to add it again */

		if existing_fn.GetOpArray().GetOpcodes() == fn.GetOpArray().GetOpcodes() && (existing_fn.GetFnFlags()&(1<<0|1<<1|1<<2)) == (fn.GetFnFlags()&(1<<0|1<<1|1<<2)) && (existing_fn.GetScope().GetCeFlags()&1<<1) == 1<<1 {
			return
		}
		if existing_fn.GetScope() == ce {

			/* members from the current class override trait methods */

			if (*overridden) != nil {
				if g.Assign(&existing_fn, ZendHashFindPtr(*overridden, key)) != nil {
					if (existing_fn.GetFnFlags() & 1 << 6) != 0 {

						/* Make sure the trait method is compatible with previosly declared abstract method */

						PerformDelayableImplementationCheck(ce, fn, existing_fn, 1)

						/* Make sure the trait method is compatible with previosly declared abstract method */

					}
					if (fn.GetFnFlags() & 1 << 6) != 0 {

						/* Make sure the abstract declaration is compatible with previous declaration */

						PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
						return
					}
				}
			} else {
				*overridden = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
				_zendHashInit(*overridden, 8, OverriddenPtrDtor, 0)
			}
			ZendHashUpdateMem(*overridden, key, fn, g.SizeOf("zend_function"))
			return
		} else if (fn.GetFnFlags()&1<<6) != 0 && (existing_fn.GetFnFlags()&1<<6) == 0 {

			/* Make sure the abstract declaration is compatible with previous declaration */

			PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
			return
		} else if (existing_fn.GetScope().GetCeFlags()&1<<1) != 0 && (existing_fn.GetFnFlags()&1<<6) == 0 {

			/* two traits can't define the __special__  same non-abstract method */

			ZendErrorNoreturn(1<<6, "Trait method %s has not been applied, because there are collisions with other trait methods on %s", name, ce.GetName().GetVal())

			/* two traits can't define the __special__  same non-abstract method */

		} else {

			/* inherited members are overridden by members inserted by traits */

			DoInheritanceCheckOnMethod(fn, existing_fn, ce, nil)
			fn.SetPrototype(nil)
		}
	}
	if fn.GetType() == 1 {
		new_fn = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_internal_function"))
		memcpy(new_fn, fn, g.SizeOf("zend_internal_function"))
		new_fn.SetFnFlags(new_fn.GetFnFlags() | 1<<25)
	} else {
		new_fn = ZendArenaAlloc(&CG.arena, g.SizeOf("zend_op_array"))
		memcpy(new_fn, fn, g.SizeOf("zend_op_array"))
		new_fn.GetOpArray().SetFnFlags(new_fn.GetOpArray().GetFnFlags() | 1<<27)
		new_fn.GetOpArray().SetFnFlags(new_fn.GetOpArray().GetFnFlags() &^ (1 << 7))
	}
	FunctionAddRef(new_fn)
	fn = ZendHashUpdatePtr(&ce.function_table, key, new_fn)
	ZendAddMagicMethods(ce, key, fn)
}

/* }}} */

func ZendFixupTraitMethod(fn *ZendFunction, ce *ZendClassEntry) {
	if (fn.GetScope().GetCeFlags() & 1 << 1) == 1<<1 {
		fn.SetScope(ce)
		if (fn.GetFnFlags() & 1 << 6) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | 1<<4)
		}
		if fn.GetType() == 2 && fn.GetOpArray().GetStaticVariables() != nil {
			ce.SetCeFlags(ce.GetCeFlags() | 1<<16)
		}
	}
}

/* }}} */

func ZendTraitsCopyFunctions(fnname *ZendString, fn *ZendFunction, ce *ZendClassEntry, overridden **HashTable, exclude_table *HashTable, aliases **ZendClassEntry) {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	var lcname *ZendString
	var fn_copy ZendFunction
	var i int

	/* apply aliases which are qualified with a class name, there should not be any ambiguity */

	if ce.GetTraitAliases() != nil {
		alias_ptr = ce.GetTraitAliases()
		alias = *alias_ptr
		i = 0
		for alias != nil {

			/* Scope unset or equal to the function we compare to, and the alias applies to fn */

			if alias.GetAlias() != nil && (aliases[i] == nil || fn.GetScope() == aliases[i]) && alias.GetTraitMethod().GetMethodName().GetLen() == fnname.GetLen() && ZendBinaryStrcasecmp(alias.GetTraitMethod().GetMethodName().GetVal(), alias.GetTraitMethod().GetMethodName().GetLen(), fnname.GetVal(), fnname.GetLen()) == 0 {
				fn_copy = *fn

				/* if it is 0, no modifieres has been changed */

				if alias.GetModifiers() != 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&(1<<0|1<<1|1<<2))
				}
				lcname = ZendStringTolowerEx(alias.GetAlias(), 0)
				ZendAddTraitMethod(ce, alias.GetAlias().GetVal(), lcname, &fn_copy, overridden)
				ZendStringReleaseEx(lcname, 0)

				/* Record the trait from which this alias was resolved. */

				if aliases[i] == nil {
					aliases[i] = fn.GetScope()
				}
				if alias.GetTraitMethod().GetClassName() == nil {

					/* TODO: try to avoid this assignment (it's necessary only for reflection) */

					alias.GetTraitMethod().SetClassName(ZendStringCopy(fn.GetScope().GetName()))

					/* TODO: try to avoid this assignment (it's necessary only for reflection) */

				}
			}
			alias_ptr++
			alias = *alias_ptr
			i++
		}
	}
	if exclude_table == nil || ZendHashFind(exclude_table, fnname) == nil {

		/* is not in hashtable, thus, function is not to be excluded */

		memcpy(&fn_copy, fn, g.CondF(fn.GetType() == 2, func() __auto__ { return g.SizeOf("zend_op_array") }, func() __auto__ { return g.SizeOf("zend_internal_function") }))

		/* apply aliases which have not alias name, just setting visibility */

		if ce.GetTraitAliases() != nil {
			alias_ptr = ce.GetTraitAliases()
			alias = *alias_ptr
			i = 0
			for alias != nil {

				/* Scope unset or equal to the function we compare to, and the alias applies to fn */

				if alias.GetAlias() == nil && alias.GetModifiers() != 0 && (aliases[i] == nil || fn.GetScope() == aliases[i]) && alias.GetTraitMethod().GetMethodName().GetLen() == fnname.GetLen() && ZendBinaryStrcasecmp(alias.GetTraitMethod().GetMethodName().GetVal(), alias.GetTraitMethod().GetMethodName().GetLen(), fnname.GetVal(), fnname.GetLen()) == 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&(1<<0|1<<1|1<<2))

					/** Record the trait from which this alias was resolved. */

					if aliases[i] == nil {
						aliases[i] = fn.GetScope()
					}
					if alias.GetTraitMethod().GetClassName() == nil {

						/* TODO: try to avoid this assignment (it's necessary only for reflection) */

						alias.GetTraitMethod().SetClassName(ZendStringCopy(fn.GetScope().GetName()))

						/* TODO: try to avoid this assignment (it's necessary only for reflection) */

					}
				}
				alias_ptr++
				alias = *alias_ptr
				i++
			}
		}
		ZendAddTraitMethod(ce, fn.GetFunctionName().GetVal(), fnname, &fn_copy, overridden)
	}
}

/* }}} */

func ZendCheckTraitUsage(ce *ZendClassEntry, trait *ZendClassEntry, traits **ZendClassEntry) uint32 {
	var i uint32
	if (trait.GetCeFlags() & 1 << 1) != 1<<1 {
		ZendErrorNoreturn(1<<6, "Class %s is not a trait, Only traits may be used in 'as' and 'insteadof' statements", trait.GetName().GetVal())
		return 0
	}
	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == trait {
			return i
		}
	}
	ZendErrorNoreturn(1<<6, "Required Trait %s wasn't added to %s", trait.GetName().GetVal(), ce.GetName().GetVal())
	return 0
}

/* }}} */

func ZendTraitsInitTraitStructures(ce *ZendClassEntry, traits **ZendClassEntry, exclude_tables_ptr ***HashTable, aliases_ptr ***ZendClassEntry) {
	var i int
	var j int = 0
	var precedences **ZendTraitPrecedence
	var cur_precedence *ZendTraitPrecedence
	var cur_method_ref *ZendTraitMethodReference
	var lcname *ZendString
	var exclude_tables **HashTable = nil
	var aliases **ZendClassEntry = nil
	var trait *ZendClassEntry

	/* resolve class references */

	if ce.GetTraitPrecedences() != nil {
		exclude_tables = _ecalloc(ce.GetNumTraits(), g.SizeOf("HashTable *"))
		i = 0
		precedences = ce.GetTraitPrecedences()
		ce.SetTraitPrecedences(nil)
		for g.Assign(&cur_precedence, precedences[i]) {

			/** Resolve classes for all precedence operations. */

			cur_method_ref = &cur_precedence.trait_method
			trait = ZendFetchClass(cur_method_ref.GetClassName(), 6|0x80)
			if trait == nil {
				ZendErrorNoreturn(1<<6, "Could not find trait %s", cur_method_ref.GetClassName().GetVal())
			}
			ZendCheckTraitUsage(ce, trait, traits)

			/** Ensure that the preferred method is actually available. */

			lcname = ZendStringTolowerEx(cur_method_ref.GetMethodName(), 0)
			if ZendHashExists(&trait.function_table, lcname) == 0 {
				ZendErrorNoreturn(1<<6, "A precedence rule was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.GetMethodName().GetVal())
			}

			/** With the other traits, we are more permissive.
			  We do not give errors for those. This allows to be more
			  defensive in such definitions.
			  However, we want to make sure that the insteadof declaration
			  is consistent in itself.
			*/

			for j = 0; j < cur_precedence.GetNumExcludes(); j++ {
				var class_name *ZendString = cur_precedence.GetExcludeClassNames()[j]
				var exclude_ce *ZendClassEntry = ZendFetchClass(class_name, 6|0x80)
				var trait_num uint32
				if exclude_ce == nil {
					ZendErrorNoreturn(1<<6, "Could not find trait %s", class_name.GetVal())
				}
				trait_num = ZendCheckTraitUsage(ce, exclude_ce, traits)
				if exclude_tables[trait_num] == nil {
					exclude_tables[trait_num] = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
					_zendHashInit(exclude_tables[trait_num], 0, nil, 0)
				}
				if ZendHashAddEmptyElement(exclude_tables[trait_num], lcname) == nil {
					ZendErrorNoreturn(1<<6, "Failed to evaluate a trait precedence (%s). Method of trait %s was defined to be excluded multiple times", precedences[i].GetTraitMethod().GetMethodName().GetVal(), exclude_ce.GetName().GetVal())
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

				if trait == exclude_ce {
					ZendErrorNoreturn(1<<6, "Inconsistent insteadof definition. "+"The method %s is to be used from %s, but %s is also on the exclude list", cur_method_ref.GetMethodName().GetVal(), trait.GetName().GetVal(), trait.GetName().GetVal())
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

			}
			ZendStringReleaseEx(lcname, 0)
			i++
		}
		ce.SetTraitPrecedences(precedences)
	}
	if ce.GetTraitAliases() != nil {
		i = 0
		for ce.GetTraitAliases()[i] != nil {
			i++
		}
		aliases = _ecalloc(i, g.SizeOf("zend_class_entry *"))
		i = 0
		for ce.GetTraitAliases()[i] != nil {

			/** For all aliases with an explicit class name, resolve the class now. */

			if ce.GetTraitAliases()[i].GetTraitMethod().GetClassName() != nil {
				cur_method_ref = &ce.trait_aliases[i].GetTraitMethod()
				trait = ZendFetchClass(cur_method_ref.GetClassName(), 6|0x80)
				if trait == nil {
					ZendErrorNoreturn(1<<6, "Could not find trait %s", cur_method_ref.GetClassName().GetVal())
				}
				ZendCheckTraitUsage(ce, trait, traits)
				aliases[i] = trait

				/** And, ensure that the referenced method is resolvable, too. */

				lcname = ZendStringTolowerEx(cur_method_ref.GetMethodName(), 0)
				if ZendHashExists(&trait.function_table, lcname) == 0 {
					ZendErrorNoreturn(1<<6, "An alias was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.GetMethodName().GetVal())
				}
				ZendStringReleaseEx(lcname, 0)
			}
			i++
		}
	}
	*exclude_tables_ptr = exclude_tables
	*aliases_ptr = aliases
}

/* }}} */

func ZendDoTraitsMethodBinding(ce *ZendClassEntry, traits **ZendClassEntry, exclude_tables **HashTable, aliases **ZendClassEntry) {
	var i uint32
	var overridden *HashTable = nil
	var key *ZendString
	var fn *ZendFunction
	if exclude_tables != nil {
		for i = 0; i < ce.GetNumTraits(); i++ {
			if traits[i] != nil {

				/* copies functions, applies defined aliasing, and excludes unused trait methods */

				for {
					var __ht *HashTable = &traits[i].function_table
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						key = _p.GetKey()
						fn = _z.GetValue().GetPtr()
						ZendTraitsCopyFunctions(key, fn, ce, &overridden, exclude_tables[i], aliases)
					}
					break
				}
				if exclude_tables[i] != nil {
					ZendHashDestroy(exclude_tables[i])
					_efree(exclude_tables[i])
					exclude_tables[i] = nil
				}
			}
		}
	} else {
		for i = 0; i < ce.GetNumTraits(); i++ {
			if traits[i] != nil {
				for {
					var __ht *HashTable = &traits[i].function_table
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if _z.GetType() == 0 {
							continue
						}
						key = _p.GetKey()
						fn = _z.GetValue().GetPtr()
						ZendTraitsCopyFunctions(key, fn, ce, &overridden, nil, aliases)
					}
					break
				}
			}
		}
	}
	for {
		var __ht *HashTable = &ce.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			fn = _z.GetValue().GetPtr()
			ZendFixupTraitMethod(fn, ce)
		}
		break
	}
	if overridden != nil {
		ZendHashDestroy(overridden)
		_efree(overridden)
	}
}

/* }}} */

func FindFirstDefinition(ce *ZendClassEntry, traits **ZendClassEntry, current_trait int, prop_name *ZendString, coliding_ce *ZendClassEntry) *ZendClassEntry {
	var i int
	if coliding_ce == ce {
		for i = 0; i < current_trait; i++ {
			if traits[i] != nil && ZendHashExists(&traits[i].properties_info, prop_name) != 0 {
				return traits[i]
			}
		}
	}
	return coliding_ce
}

/* }}} */

func ZendDoTraitsPropertyBinding(ce *ZendClassEntry, traits **ZendClassEntry) {
	var i int
	var property_info *ZendPropertyInfo
	var coliding_prop *ZendPropertyInfo
	var prop_name *ZendString
	var class_name_unused *byte
	var not_compatible ZendBool
	var prop_value *Zval
	var flags uint32
	var doc_comment *ZendString

	/* In the following steps the properties are inserted into the property table
	 * for that, a very strict approach is applied:
	 * - check for compatibility, if not compatible with any property in class -> fatal
	 * - if compatible, then strict notice
	 */

	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == nil {
			continue
		}
		for {
			var __ht *HashTable = &traits[i].properties_info
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				property_info = _z.GetValue().GetPtr()

				/* first get the unmangeld name if necessary,
				 * then check whether the property is already there
				 */

				flags = property_info.GetFlags()
				if (flags & 1 << 0) != 0 {
					prop_name = ZendStringCopy(property_info.GetName())
				} else {
					var pname *byte
					var pname_len int

					/* for private and protected we need to unmangle the names */

					ZendUnmanglePropertyNameEx(property_info.GetName(), &class_name_unused, &pname, &pname_len)
					prop_name = ZendStringInit(pname, pname_len, 0)
				}

				/* next: check for conflicts with current class */

				if g.Assign(&coliding_prop, ZendHashFindPtr(&ce.properties_info, prop_name)) != nil {
					if (coliding_prop.GetFlags()&1<<2) != 0 && coliding_prop.GetCe() != ce {
						ZendHashDel(&ce.properties_info, prop_name)
						flags |= 1 << 3
					} else {
						not_compatible = 1
						if (coliding_prop.GetFlags()&(1<<0|1<<1|1<<2|1<<4)) == (flags&(1<<0|1<<1|1<<2|1<<4)) && PropertyTypesCompatible(property_info, coliding_prop) == INHERITANCE_SUCCESS {

							/* the flags are identical, thus, the properties may be compatible */

							var op1 *Zval
							var op2 *Zval
							var op1_tmp Zval
							var op2_tmp Zval
							if (flags & 1 << 4) != 0 {
								op1 = &ce.default_static_members_table[coliding_prop.GetOffset()]
								op2 = &traits[i].default_static_members_table[property_info.GetOffset()]
								if op1.GetType() == 13 {
									op1 = op1.GetValue().GetZv()
								}
								if op2.GetType() == 13 {
									op2 = op2.GetValue().GetZv()
								}
							} else {
								op1 = &ce.default_properties_table[(coliding_prop.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")]
								op2 = &traits[i].default_properties_table[(property_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")]
							}

							/* if any of the values is a constant, we try to resolve it */

							if op1.GetType() == 11 {
								var _z1 *Zval = &op1_tmp
								var _z2 *Zval = op1
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								if (_t & 0xff00) != 0 {
									if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
										ZendGcAddref(&_gc.gc)
									} else {
										ZvalCopyCtorFunc(_z1)
									}
								}
								ZvalUpdateConstantEx(&op1_tmp, ce)
								op1 = &op1_tmp
							}
							if op2.GetType() == 11 {
								var _z1 *Zval = &op2_tmp
								var _z2 *Zval = op2
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								if (_t & 0xff00) != 0 {
									if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
										ZendGcAddref(&_gc.gc)
									} else {
										ZvalCopyCtorFunc(_z1)
									}
								}
								ZvalUpdateConstantEx(&op2_tmp, ce)
								op2 = &op2_tmp
							}
							not_compatible = FastIsNotIdenticalFunction(op1, op2)
							if op1 == &op1_tmp {
								ZvalPtrDtorNogc(&op1_tmp)
							}
							if op2 == &op2_tmp {
								ZvalPtrDtorNogc(&op2_tmp)
							}
						}
						if not_compatible != 0 {
							ZendErrorNoreturn(1<<6, "%s and %s define the __special__  same property ($%s) in the composition of %s. However, the definition differs and is considered incompatible. Class was composed", FindFirstDefinition(ce, traits, i, prop_name, coliding_prop.GetCe()).GetName().GetVal(), property_info.GetCe().GetName().GetVal(), prop_name.GetVal(), ce.GetName().GetVal())
						}
						ZendStringReleaseEx(prop_name, 0)
						continue
					}
				}

				/* property not found, so lets add it */

				if (flags & 1 << 4) != 0 {
					prop_value = &traits[i].default_static_members_table[property_info.GetOffset()]
					assert(prop_value.GetType() != 13)
				} else {
					prop_value = &traits[i].default_properties_table[(property_info.GetOffset()-uint32(zend_long((*byte)(&((*ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil))+g.SizeOf("zval")*0))/g.SizeOf("zval")]
				}
				if prop_value.GetTypeFlags() != 0 {
					ZvalAddrefP(prop_value)
				}
				if property_info.GetDocComment() != nil {
					doc_comment = ZendStringCopy(property_info.GetDocComment())
				} else {
					doc_comment = nil
				}
				if property_info.GetType() > 0x3ff && (property_info.GetType()&0x2) == 0 {
					ZendStringAddref((*ZendString)(property_info.GetType() & ^0x3))
				}
				ZendDeclareTypedProperty(ce, prop_name, prop_value, flags, doc_comment, property_info.GetType())
				ZendStringReleaseEx(prop_name, 0)
			}
			break
		}
	}

	/* In the following steps the properties are inserted into the property table
	 * for that, a very strict approach is applied:
	 * - check for compatibility, if not compatible with any property in class -> fatal
	 * - if compatible, then strict notice
	 */
}

/* }}} */

func ZendDoCheckForInconsistentTraitsAliasing(ce *ZendClassEntry, aliases **ZendClassEntry) {
	var i int = 0
	var cur_alias *ZendTraitAlias
	var lc_method_name *ZendString
	if ce.GetTraitAliases() != nil {
		for ce.GetTraitAliases()[i] != nil {
			cur_alias = ce.GetTraitAliases()[i]

			/** The trait for this alias has not been resolved, this means, this
			  alias was not applied. Abort with an error. */

			if aliases[i] == nil {
				if cur_alias.GetAlias() != nil {

					/** Plain old inconsistency/typo/bug */

					ZendErrorNoreturn(1<<6, "An alias (%s) was defined for method %s(), but this method does not exist", cur_alias.GetAlias().GetVal(), cur_alias.GetTraitMethod().GetMethodName().GetVal())

					/** Plain old inconsistency/typo/bug */

				} else {

					/** Here are two possible cases:
					  1) this is an attempt to modify the visibility
					     of a method introduce as part of another alias.
					     Since that seems to violate the DRY principle,
					     we check against it and abort.
					  2) it is just a plain old inconsitency/typo/bug
					     as in the case where alias is set. */

					lc_method_name = ZendStringTolowerEx(cur_alias.GetTraitMethod().GetMethodName(), 0)
					if ZendHashExists(&ce.function_table, lc_method_name) != 0 {
						ZendStringReleaseEx(lc_method_name, 0)
						ZendErrorNoreturn(1<<6, "The modifiers for the trait alias %s() need to be changed in the same statement in which the alias is defined. Error", cur_alias.GetTraitMethod().GetMethodName().GetVal())
					} else {
						ZendStringReleaseEx(lc_method_name, 0)
						ZendErrorNoreturn(1<<6, "The modifiers of the trait method %s() are changed, but this method does not exist. Error", cur_alias.GetTraitMethod().GetMethodName().GetVal())
					}
				}
			}
			i++
		}
	}
}

/* }}} */

func ZendDoBindTraits(ce *ZendClassEntry) {
	var exclude_tables **HashTable
	var aliases **ZendClassEntry
	var traits **ZendClassEntry
	var trait **ZendClassEntry
	var i uint32
	var j uint32
	assert(ce.GetNumTraits() > 0)
	traits = _emalloc(g.SizeOf("zend_class_entry *") * ce.GetNumTraits())
	for i = 0; i < ce.GetNumTraits(); i++ {
		trait = ZendFetchClassByName(ce.GetTraitNames()[i].GetName(), ce.GetTraitNames()[i].GetLcName(), 6)
		if trait == nil {
			return
		}
		if (trait.ce_flags & 1 << 1) == 0 {
			ZendErrorNoreturn(1<<0, "%s cannot use %s - it is not a trait", ce.GetName().GetVal(), trait.name.GetVal())
			return
		}
		for j = 0; j < i; j++ {
			if traits[j] == trait {

				/* skip duplications */

				trait = nil
				break
			}
		}
		traits[i] = trait
	}

	/* complete initialization of trait strutures in ce */

	ZendTraitsInitTraitStructures(ce, traits, &exclude_tables, &aliases)

	/* first care about all methods to be flattened into the class */

	ZendDoTraitsMethodBinding(ce, traits, exclude_tables, aliases)

	/* Aliases which have not been applied indicate typos/bugs. */

	ZendDoCheckForInconsistentTraitsAliasing(ce, aliases)
	if aliases != nil {
		_efree(aliases)
	}
	if exclude_tables != nil {
		_efree(exclude_tables)
	}

	/* then flatten the properties into it, to, mostly to notfiy developer about problems */

	ZendDoTraitsPropertyBinding(ce, traits)
	_efree(traits)

	/* Emit E_DEPRECATED for PHP 4 constructors */

	ZendCheckDeprecatedConstructor(ce)

	/* Emit E_DEPRECATED for PHP 4 constructors */
}

/* }}} */

func ZendHasDeprecatedConstructor(ce *ZendClassEntry) ZendBool {
	var constructor_name *ZendString
	if ce.GetConstructor() == nil {
		return 0
	}
	constructor_name = ce.GetConstructor().GetFunctionName()
	return !(ZendBinaryStrcasecmp(ce.GetName().GetVal(), ce.GetName().GetLen(), constructor_name.GetVal(), constructor_name.GetLen()))
}

/* }}} */

func ZendCheckDeprecatedConstructor(ce *ZendClassEntry) {
	if ZendHasDeprecatedConstructor(ce) != 0 {
		ZendError(1<<13, "Methods with the same name as their class will not be constructors in a future version of PHP; %s has a deprecated constructor", ce.GetName().GetVal())
	}
}

/* }}} */

// #define MAX_ABSTRACT_INFO_CNT       3

// #define MAX_ABSTRACT_INFO_FMT       "%s%s%s%s"

// #define DISPLAY_ABSTRACT_FN(idx) ai . afn [ idx ] ? ZEND_FN_SCOPE_NAME ( ai . afn [ idx ] ) : "" , ai . afn [ idx ] ? "::" : "" , ai . afn [ idx ] ? ZSTR_VAL ( ai . afn [ idx ] -> common . function_name ) : "" , ai . afn [ idx ] && ai . afn [ idx + 1 ] ? ", " : ( ai . afn [ idx ] && ai . cnt > MAX_ABSTRACT_INFO_CNT ? ", ..." : "" )

// @type ZendAbstractInfo struct

func ZendVerifyAbstractClassFunction(fn *ZendFunction, ai *ZendAbstractInfo) {
	if (fn.GetFnFlags() & 1 << 6) != 0 {
		if ai.GetCnt() < 3 {
			ai.GetAfn()[ai.GetCnt()] = fn
		}
		if (fn.GetFnFlags() & 1 << 28) != 0 {
			if ai.GetCtor() == 0 {
				ai.GetCnt()++
				ai.SetCtor(1)
			} else {
				ai.GetAfn()[ai.GetCnt()] = nil
			}
		} else {
			ai.GetCnt()++
		}
	}
}

/* }}} */

func ZendVerifyAbstractClass(ce *ZendClassEntry) {
	var func_ *ZendFunction
	var ai ZendAbstractInfo
	assert((ce.GetCeFlags() & (1<<4 | 1<<0 | 1<<1 | 1<<6)) == 1<<4)
	memset(&ai, 0, g.SizeOf("ai"))
	for {
		var __ht *HashTable = &ce.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			func_ = _z.GetValue().GetPtr()
			ZendVerifyAbstractClassFunction(func_, &ai)
		}
		break
	}
	if ai.GetCnt() != 0 {
		ZendErrorNoreturn(1<<0, "Class %s contains %d abstract method%s and must therefore be declared abstract or implement the remaining methods ("+"%s%s%s%s"+"%s%s%s%s"+"%s%s%s%s"+")", ce.GetName().GetVal(), ai.GetCnt(), g.Cond(ai.GetCnt() > 1, "s", ""), g.CondF1(ai.GetAfn()[0] != nil, func() __auto__ {
			if ai.GetAfn()[0] != nil && ai.GetAfn()[0].GetScope() != nil {
				return ai.GetAfn()[0].GetScope().GetName().GetVal()
			} else {
				return ""
			}
		}, ""), g.Cond(ai.GetAfn()[0] != nil, "::", ""), g.CondF1(ai.GetAfn()[0] != nil, func() []byte { return ai.GetAfn()[0].GetFunctionName().GetVal() }, ""), g.CondF2(ai.GetAfn()[0] != nil && ai.GetAfn()[0+1] != nil, ", ", func() string {
			if ai.GetAfn()[0] != nil && ai.GetCnt() > 3 {
				return ", ..."
			} else {
				return ""
			}
		}), g.CondF1(ai.GetAfn()[1] != nil, func() __auto__ {
			if ai.GetAfn()[1] != nil && ai.GetAfn()[1].GetScope() != nil {
				return ai.GetAfn()[1].GetScope().GetName().GetVal()
			} else {
				return ""
			}
		}, ""), g.Cond(ai.GetAfn()[1] != nil, "::", ""), g.CondF1(ai.GetAfn()[1] != nil, func() []byte { return ai.GetAfn()[1].GetFunctionName().GetVal() }, ""), g.CondF2(ai.GetAfn()[1] != nil && ai.GetAfn()[1+1] != nil, ", ", func() string {
			if ai.GetAfn()[1] != nil && ai.GetCnt() > 3 {
				return ", ..."
			} else {
				return ""
			}
		}), g.CondF1(ai.GetAfn()[2] != nil, func() __auto__ {
			if ai.GetAfn()[2] != nil && ai.GetAfn()[2].GetScope() != nil {
				return ai.GetAfn()[2].GetScope().GetName().GetVal()
			} else {
				return ""
			}
		}, ""), g.Cond(ai.GetAfn()[2] != nil, "::", ""), g.CondF1(ai.GetAfn()[2] != nil, func() []byte { return ai.GetAfn()[2].GetFunctionName().GetVal() }, ""), g.CondF2(ai.GetAfn()[2] != nil && ai.GetAfn()[2+1] != nil, ", ", func() string {
			if ai.GetAfn()[2] != nil && ai.GetCnt() > 3 {
				return ", ..."
			} else {
				return ""
			}
		}))
	} else {

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

		ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 4))

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

	}
}

/* }}} */

type VarianceObligationType = int

const (
	OBLIGATION_DEPENDENCY = iota
	OBLIGATION_COMPATIBILITY
	OBLIGATION_PROPERTY_COMPATIBILITY
)

// @type VarianceObligation struct
func VarianceObligationDtor(zv *Zval) { _efree(zv.GetValue().GetPtr()) }
func VarianceObligationHtDtor(zv *Zval) {
	ZendHashDestroy(zv.GetValue().GetPtr())
	_efree(zv.GetValue().GetPtr())
}
func GetOrInitObligationsForClass(ce *ZendClassEntry) *HashTable {
	var ht *HashTable
	var key ZendUlong
	if CG.GetDelayedVarianceObligations() == nil {
		CG.SetDelayedVarianceObligations((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
		_zendHashInit(CG.GetDelayedVarianceObligations(), 0, VarianceObligationHtDtor, 0)
	}
	key = ZendUlong(uintPtr(ce))
	ht = ZendHashIndexFindPtr(CG.GetDelayedVarianceObligations(), key)
	if ht != nil {
		return ht
	}
	ht = (*HashTable)(_emalloc(g.SizeOf("HashTable")))
	_zendHashInit(ht, 0, VarianceObligationDtor, 0)
	ZendHashIndexAddNewPtr(CG.GetDelayedVarianceObligations(), key, ht)
	ce.SetCeFlags(ce.GetCeFlags() | 1<<21)
	return ht
}
func AddDependencyObligation(ce *ZendClassEntry, dependency_ce *ZendClassEntry) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = _emalloc(g.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_DEPENDENCY)
	obligation.dependency_ce = dependency_ce
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddCompatibilityObligation(ce *ZendClassEntry, child_fn *ZendFunction, parent_fn *ZendFunction, always_error ZendBool) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = _emalloc(g.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_COMPATIBILITY)

	/* Copy functions, because they may be stack-allocated in the case of traits. */

	if child_fn.GetCommonType() == 1 {
		memcpy(&obligation.child_fn, child_fn, g.SizeOf("zend_internal_function"))
	} else {
		memcpy(&obligation.child_fn, child_fn, g.SizeOf("zend_op_array"))
	}
	if parent_fn.GetCommonType() == 1 {
		memcpy(&obligation.parent_fn, parent_fn, g.SizeOf("zend_internal_function"))
	} else {
		memcpy(&obligation.parent_fn, parent_fn, g.SizeOf("zend_op_array"))
	}
	obligation.always_error = always_error
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddPropertyCompatibilityObligation(ce *ZendClassEntry, child_prop *ZendPropertyInfo, parent_prop *ZendPropertyInfo) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = _emalloc(g.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_PROPERTY_COMPATIBILITY)
	obligation.child_prop = child_prop
	obligation.parent_prop = parent_prop
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func CheckVarianceObligation(zv *Zval) int {
	var obligation *VarianceObligation = zv.GetValue().GetPtr()
	if obligation.GetType() == OBLIGATION_DEPENDENCY {
		var dependency_ce *ZendClassEntry = obligation.dependency_ce
		if (dependency_ce.GetCeFlags() & 1 << 21) != 0 {
			ResolveDelayedVarianceObligations(dependency_ce)
		}
		if (dependency_ce.GetCeFlags() & 1 << 3) == 0 {
			return 0
		}
	} else if obligation.GetType() == OBLIGATION_COMPATIBILITY {
		var unresolved_class *ZendString
		var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, &obligation.child_fn, &obligation.parent_fn)
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return 0
			}
			assert(status == INHERITANCE_ERROR)
			EmitIncompatibleMethodErrorOrWarning(&obligation.child_fn, &obligation.parent_fn, status, unresolved_class, obligation.always_error)
		}
	} else {
		assert(obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY)
		var status InheritanceStatus = PropertyTypesCompatible(obligation.parent_prop, obligation.child_prop)
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return 0
			}
			assert(status == INHERITANCE_ERROR)
			EmitIncompatiblePropertyError(obligation.child_prop, obligation.parent_prop)
		}
	}
	return 1 << 0
}
func LoadDelayedClasses() {
	var delayed_autoloads *HashTable = CG.GetDelayedAutoloads()
	var name *ZendString
	if delayed_autoloads == nil {
		return
	}

	/* Take ownership of this HT, to avoid concurrent modification during autoloading. */

	CG.SetDelayedAutoloads(nil)
	for {
		var __ht *HashTable = delayed_autoloads
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			name = _p.GetKey()
			ZendLookupClass(name)
		}
		break
	}
	ZendHashDestroy(delayed_autoloads)
	_efree(delayed_autoloads)
}
func ResolveDelayedVarianceObligations(ce *ZendClassEntry) {
	var all_obligations *HashTable = CG.GetDelayedVarianceObligations()
	var obligations *HashTable
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	assert(all_obligations != nil)
	obligations = ZendHashIndexFindPtr(all_obligations, num_key)
	assert(obligations != nil)
	ZendHashApply(obligations, CheckVarianceObligation)
	if obligations.GetNNumOfElements() == 0 {
		ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 21))
		ce.SetCeFlags(ce.GetCeFlags() | 1<<3)
		ZendHashIndexDel(all_obligations, num_key)
	}
}
func ReportVarianceErrors(ce *ZendClassEntry) {
	var all_obligations *HashTable = CG.GetDelayedVarianceObligations()
	var obligations *HashTable
	var obligation *VarianceObligation
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	assert(all_obligations != nil)
	obligations = ZendHashIndexFindPtr(all_obligations, num_key)
	assert(obligations != nil)
	for {
		var __ht *HashTable = obligations
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			obligation = _z.GetValue().GetPtr()
			var status InheritanceStatus
			var unresolved_class *ZendString
			if obligation.GetType() == OBLIGATION_COMPATIBILITY {

				/* Just used to fetch the unresolved_class in this case. */

				status = ZendDoPerformImplementationCheck(&unresolved_class, &obligation.child_fn, &obligation.parent_fn)
				assert(status == INHERITANCE_UNRESOLVED)
				EmitIncompatibleMethodErrorOrWarning(&obligation.child_fn, &obligation.parent_fn, status, unresolved_class, obligation.always_error)
			} else if obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY {
				EmitIncompatiblePropertyError(obligation.child_prop, obligation.parent_prop)
			} else {
				ZendErrorNoreturn(1<<4, "Bug #78647")
			}
		}
		break
	}

	/* Only warnings were thrown above -- that means that there are incompatibilities, but only
	 * ones that we permit. Mark all classes with open obligations as fully linked. */

	ce.SetCeFlags(ce.GetCeFlags() &^ (1 << 21))
	ce.SetCeFlags(ce.GetCeFlags() | 1<<3)
	ZendHashIndexDel(all_obligations, num_key)
}
func CheckUnrecoverableLoadFailure(ce *ZendClassEntry) {
	/* If this class has been used while unlinked through a variance obligation, it is not legal
	 * to remove the class from the class table and throw an exception, because there is already
	 * a dependence on the inheritance hierarchy of this specific class. Instead we fall back to
	 * a fatal error, as would happen if we did not allow exceptions in the first place. */

	if (ce.GetCeFlags() & 1 << 23) != 0 {
		var exception_str *ZendString
		var exception_zv Zval
		assert(EG.GetException() != nil && "Exception must have been thrown")
		var __z *Zval = &exception_zv
		__z.GetValue().SetObj(EG.GetException())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZvalAddrefP(&exception_zv)
		ZendClearException()
		exception_str = ZvalGetString(&exception_zv)
		ZendErrorNoreturn(1<<0, "During inheritance of %s with variance dependencies: Uncaught %s", ce.GetName().GetVal(), exception_str.GetVal())
	}

	/* If this class has been used while unlinked through a variance obligation, it is not legal
	 * to remove the class from the class table and throw an exception, because there is already
	 * a dependence on the inheritance hierarchy of this specific class. Instead we fall back to
	 * a fatal error, as would happen if we did not allow exceptions in the first place. */
}
func ZendDoLinkClass(ce *ZendClassEntry, lc_parent_name *ZendString) int {
	/* Load parent/interface dependencies first, so we can still gracefully abort linking
	 * with an exception and remove the class from the class table. This is only possible
	 * if no variance obligations on the current class have been added during autoloading. */

	var parent *ZendClassEntry = nil
	var interfaces **ZendClassEntry = nil
	if ce.parent_name {
		parent = ZendFetchClassByName(ce.parent_name, lc_parent_name, 0x800|0x200)
		if parent == nil {
			CheckUnrecoverableLoadFailure(ce)
			return FAILURE
		}
	}
	if ce.GetNumInterfaces() != 0 {

		/* Also copy the parent interfaces here, so we don't need to reallocate later. */

		var i uint32
		var num_parent_interfaces uint32 = g.CondF1(parent != nil, func() uint32 { return parent.GetNumInterfaces() }, 0)
		interfaces = _emalloc(g.SizeOf("zend_class_entry *") * (ce.GetNumInterfaces() + num_parent_interfaces))
		if num_parent_interfaces != 0 {
			memcpy(interfaces, parent.interfaces, g.SizeOf("zend_class_entry *")*num_parent_interfaces)
		}
		for i = 0; i < ce.GetNumInterfaces(); i++ {
			var iface *ZendClassEntry = ZendFetchClassByName(ce.interface_names[i].name, ce.interface_names[i].lc_name, 5|0x800|0x200)
			if iface == nil {
				CheckUnrecoverableLoadFailure(ce)
				_efree(interfaces)
				return FAILURE
			}
			interfaces[num_parent_interfaces+i] = iface
		}
	}
	if parent != nil {
		if (parent.GetCeFlags() & 1 << 3) == 0 {
			AddDependencyObligation(ce, parent)
		}
		ZendDoInheritanceEx(ce, parent, 0)
	}
	if (ce.GetCeFlags() & 1 << 15) != 0 {
		ZendDoBindTraits(ce)
	}
	if (ce.GetCeFlags() & 1 << 14) != 0 {
		ZendDoImplementInterfaces(ce, interfaces)
	}
	if (ce.GetCeFlags() & (1<<4 | 1<<0 | 1<<1 | 1<<6)) == 1<<4 {
		ZendVerifyAbstractClass(ce)
	}
	ZendBuildPropertiesInfoTable(ce)
	if (ce.GetCeFlags() & 1 << 21) == 0 {
		ce.SetCeFlags(ce.GetCeFlags() | 1<<3)
		return SUCCESS
	}
	ce.SetCeFlags(ce.GetCeFlags() | 1<<22)
	LoadDelayedClasses()
	if (ce.GetCeFlags() & 1 << 21) != 0 {
		ResolveDelayedVarianceObligations(ce)
		if (ce.GetCeFlags() & 1 << 3) == 0 {
			ReportVarianceErrors(ce)
		}
	}
	return SUCCESS
}

/* }}} */

func ZendCanEarlyBind(ce *ZendClassEntry, parent_ce *ZendClassEntry) InheritanceStatus {
	var ret InheritanceStatus = INHERITANCE_SUCCESS
	var key *ZendString
	var parent_func *ZendFunction
	var parent_info *ZendPropertyInfo
	for {
		var __ht *HashTable = &parent_ce.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			parent_func = _z.GetValue().GetPtr()
			var zv *Zval = ZendHashFindEx(&ce.function_table, key, 1)
			if zv != nil {
				var child_func *ZendFunction = zv.GetValue().GetFunc()
				var status InheritanceStatus = DoInheritanceCheckOnMethodEx(child_func, parent_func, ce, nil, 1, 0)
				if status != INHERITANCE_SUCCESS {
					if status == INHERITANCE_UNRESOLVED {
						return INHERITANCE_UNRESOLVED
					}
					assert(status == INHERITANCE_ERROR)
					ret = INHERITANCE_ERROR
				}
			}
		}
		break
	}
	for {
		var __ht *HashTable = &parent_ce.properties_info
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if _z.GetType() == 0 {
				continue
			}
			key = _p.GetKey()
			parent_info = _z.GetValue().GetPtr()
			var zv *Zval
			if (parent_info.GetFlags()&1<<2) != 0 || parent_info.GetType() <= 0x3 {
				continue
			}
			zv = ZendHashFindEx(&ce.properties_info, key, 1)
			if zv != nil {
				var child_info *ZendPropertyInfo = zv.GetValue().GetPtr()
				if child_info.GetType() > 0x3 {
					var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
					if status != INHERITANCE_SUCCESS {
						if status == INHERITANCE_UNRESOLVED {
							return INHERITANCE_UNRESOLVED
						}
						assert(status == INHERITANCE_ERROR)
						ret = INHERITANCE_ERROR
					}
				}
			}
		}
		break
	}
	return ret
}

/* }}} */

func ZendTryEarlyBind(ce *ZendClassEntry, parent_ce *ZendClassEntry, lcname *ZendString, delayed_early_binding *Zval) ZendBool {
	var status InheritanceStatus = ZendCanEarlyBind(ce, parent_ce)
	if status != INHERITANCE_UNRESOLVED {
		if delayed_early_binding != nil {
			if ZendHashSetBucketKey(EG.GetClassTable(), (*Bucket)(delayed_early_binding), lcname) == nil {
				ZendErrorNoreturn(1<<6, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
				return 0
			}
		} else {
			if ZendHashAddPtr(CG.GetClassTable(), lcname, ce) == nil {
				return 0
			}
		}
		ZendDoInheritanceEx(ce, parent_ce, status == INHERITANCE_SUCCESS)
		ZendBuildPropertiesInfoTable(ce)
		if (ce.GetCeFlags() & (1<<4 | 1<<0 | 1<<1 | 1<<6)) == 1<<4 {
			ZendVerifyAbstractClass(ce)
		}
		assert((ce.GetCeFlags() & 1 << 21) == 0)
		ce.SetCeFlags(ce.GetCeFlags() | 1<<3)
		return 1
	}
	return 0
}

/* }}} */
