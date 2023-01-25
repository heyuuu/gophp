// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendDoInheritance(ce *ZendClassEntry, parent_ce *ZendClassEntry) {
	ZendDoInheritanceEx(ce, parent_ce, 0)
}
func OverriddenPtrDtor(zv *Zval) {
	EfreeSize(Z_PTR_P(zv), b.SizeOf("zend_function"))
}
func ZendDuplicatePropertyInfoInternal(property_info *ZendPropertyInfo) *ZendPropertyInfo {
	var new_property_info *ZendPropertyInfo = Pemalloc(b.SizeOf("zend_property_info"), 1)
	memcpy(new_property_info, property_info, b.SizeOf("zend_property_info"))
	ZendStringAddref(new_property_info.GetName())
	if ZEND_TYPE_IS_NAME(new_property_info.GetType()) {
		ZendStringAddref(ZEND_TYPE_NAME(new_property_info.GetType()))
	}
	return new_property_info
}
func ZendDuplicateInternalFunction(func_ *ZendFunction, ce *ZendClassEntry) *ZendFunction {
	var new_function *ZendFunction
	if UNEXPECTED((ce.GetType() & ZEND_INTERNAL_CLASS) != 0) {
		new_function = Pemalloc(b.SizeOf("zend_internal_function"), 1)
		memcpy(new_function, func_, b.SizeOf("zend_internal_function"))
	} else {
		new_function = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_internal_function"))
		memcpy(new_function, func_, b.SizeOf("zend_internal_function"))
		new_function.SetFnFlags(new_function.GetFnFlags() | ZEND_ACC_ARENA_ALLOCATED)
	}
	if EXPECTED(new_function.GetFunctionName() != nil) {
		ZendStringAddref(new_function.GetFunctionName())
	}
	return new_function
}
func ZendDuplicateUserFunction(func_ *ZendFunction) *ZendFunction {
	var new_function *ZendFunction
	new_function = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_op_array"))
	memcpy(new_function, func_, b.SizeOf("zend_op_array"))
	if ZEND_MAP_PTR_GET(func_.op_array.static_variables_ptr) {

		/* See: Zend/tests/method_static_var.phpt */

		new_function.GetOpArray().SetStaticVariables(ZEND_MAP_PTR_GET(func_.op_array.static_variables_ptr))

		/* See: Zend/tests/method_static_var.phpt */

	}
	if (GC_FLAGS(new_function.GetOpArray().GetStaticVariables()) & IS_ARRAY_IMMUTABLE) == 0 {
		GC_ADDREF(new_function.GetOpArray().GetStaticVariables())
	}
	if (CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		ZEND_ASSERT((new_function.GetOpArray().GetFnFlags() & ZEND_ACC_PRELOADED) != 0)
		ZEND_MAP_PTR_NEW(new_function.op_array.static_variables_ptr)
	} else {
		ZEND_MAP_PTR_INIT(new_function.op_array.static_variables_ptr, &new_function.op_array.GetStaticVariables())
	}
	return new_function
}
func ZendDuplicateFunction(func_ *ZendFunction, ce *ZendClassEntry, is_interface ZendBool) *ZendFunction {
	if UNEXPECTED(func_.GetType() == ZEND_INTERNAL_FUNCTION) {
		return ZendDuplicateInternalFunction(func_, ce)
	} else {
		if func_.GetOpArray().GetRefcount() != nil {
			(*func_).op_array.refcount++
		}
		if is_interface != 0 || EXPECTED(func_.GetOpArray().GetStaticVariables() == nil) {

			/* reuse the same op_array structure */

			return func_

			/* reuse the same op_array structure */

		}
		return ZendDuplicateUserFunction(func_)
	}
}
func DoInheritParentConstructor(ce *ZendClassEntry) {
	var parent *ZendClassEntry = ce.parent
	ZEND_ASSERT(parent != nil)

	/* You cannot change create_object */

	ce.create_object = parent.create_object

	/* Inherit special functions if needed */

	if EXPECTED(ce.GetGetIterator() == nil) {
		ce.SetGetIterator(parent.GetGetIterator())
	}
	if parent.GetIteratorFuncsPtr() != nil {

		/* Must be initialized through iface->interface_gets_implemented() */

		ZEND_ASSERT(ce.GetIteratorFuncsPtr() != nil)

		/* Must be initialized through iface->interface_gets_implemented() */

	}
	if EXPECTED(ce.GetGet() == nil) {
		ce.SetGet(parent.GetGet())
	}
	if EXPECTED(ce.GetSet() == nil) {
		ce.SetSet(parent.GetSet())
	}
	if EXPECTED(ce.GetUnset() == nil) {
		ce.SetUnset(parent.GetUnset())
	}
	if EXPECTED(ce.GetIsset() == nil) {
		ce.SetIsset(parent.GetIsset())
	}
	if EXPECTED(ce.GetCall() == nil) {
		ce.SetCall(parent.GetCall())
	}
	if EXPECTED(ce.GetCallstatic() == nil) {
		ce.SetCallstatic(parent.GetCallstatic())
	}
	if EXPECTED(ce.GetTostring() == nil) {
		ce.SetTostring(parent.GetTostring())
	}
	if EXPECTED(ce.GetClone() == nil) {
		ce.SetClone(parent.GetClone())
	}
	if EXPECTED(ce.GetSerializeFunc() == nil) {
		ce.SetSerializeFunc(parent.GetSerializeFunc())
	}
	if EXPECTED(ce.GetSerialize() == nil) {
		ce.SetSerialize(parent.GetSerialize())
	}
	if EXPECTED(ce.GetUnserializeFunc() == nil) {
		ce.SetUnserializeFunc(parent.GetUnserializeFunc())
	}
	if EXPECTED(ce.GetUnserialize() == nil) {
		ce.SetUnserialize(parent.GetUnserialize())
	}
	if ce.GetDestructor() == nil {
		ce.SetDestructor(parent.GetDestructor())
	}
	if EXPECTED(ce.GetDebugInfo() == nil) {
		ce.SetDebugInfo(parent.GetDebugInfo())
	}
	if ce.GetConstructor() != nil {
		if parent.GetConstructor() != nil && UNEXPECTED((parent.GetConstructor().GetFnFlags()&ZEND_ACC_FINAL) != 0) {
			ZendErrorNoreturn(E_ERROR, "Cannot override final %s::%s() with %s::%s()", ZSTR_VAL(parent.GetName()), ZSTR_VAL(parent.GetConstructor().GetFunctionName()), ZSTR_VAL(ce.GetName()), ZSTR_VAL(ce.GetConstructor().GetFunctionName()))
		}
		return
	}
	ce.SetConstructor(parent.GetConstructor())
}
func ZendVisibilityString(fn_flags uint32) *byte {
	if (fn_flags & ZEND_ACC_PUBLIC) != 0 {
		return "public"
	} else if (fn_flags & ZEND_ACC_PRIVATE) != 0 {
		return "private"
	} else {
		ZEND_ASSERT((fn_flags & ZEND_ACC_PROTECTED) != 0)
		return "protected"
	}
}
func ResolveClassName(scope *ZendClassEntry, name *ZendString) *ZendString {
	ZEND_ASSERT(scope != nil)
	if ZendStringEqualsLiteralCi(name, "parent") && scope.parent {
		if (scope.GetCeFlags() & ZEND_ACC_RESOLVED_PARENT) != 0 {
			return scope.parent.name
		} else {
			return scope.parent_name
		}
	} else if ZendStringEqualsLiteralCi(name, "self") {
		return scope.GetName()
	} else {
		return name
	}
}
func ClassVisible(ce *ZendClassEntry) ZendBool {
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		return !(CompilerGlobals.GetCompilerOptions() & ZEND_COMPILE_IGNORE_INTERNAL_CLASSES)
	} else {
		ZEND_ASSERT(ce.GetType() == ZEND_USER_CLASS)
		return (CompilerGlobals.GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) == 0 || ce.GetFilename() == CompilerGlobals.GetCompiledFilename()
	}
}
func LookupClass(scope *ZendClassEntry, name *ZendString) *ZendClassEntry {
	var ce *ZendClassEntry
	if CompilerGlobals.GetInCompilation() == 0 {
		var flags uint32 = ZEND_FETCH_CLASS_ALLOW_UNLINKED | ZEND_FETCH_CLASS_NO_AUTOLOAD
		ce = ZendLookupClassEx(name, nil, flags)
		if ce != nil {
			return ce
		}

		/* We'll autoload this class and process delayed variance obligations later. */

		if CompilerGlobals.GetDelayedAutoloads() == nil {
			ALLOC_HASHTABLE(CompilerGlobals.GetDelayedAutoloads())
			ZendHashInit(CompilerGlobals.GetDelayedAutoloads(), 0, nil, nil, 0)
		}
		ZendHashAddEmptyElement(CompilerGlobals.GetDelayedAutoloads(), name)
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce != nil && ClassVisible(ce) != 0 {
			return ce
		}

		/* The current class may not be registered yet, so check for it explicitly. */

		if ZendStringEqualsCi(scope.GetName(), name) {
			return scope
		}

		/* The current class may not be registered yet, so check for it explicitly. */

	}
	return nil
}
func UnlinkedInstanceof(ce1 *ZendClassEntry, ce2 *ZendClassEntry) ZendBool {
	if ce1 == ce2 {
		return 1
	}
	if (ce1.GetCeFlags() & ZEND_ACC_LINKED) != 0 {
		return InstanceofFunction(ce1, ce2)
	}
	if ce1.parent {
		var parent_ce *ZendClassEntry
		if (ce1.GetCeFlags() & ZEND_ACC_RESOLVED_PARENT) != 0 {
			parent_ce = ce1.parent
		} else {
			parent_ce = ZendLookupClassEx(ce1.parent_name, nil, ZEND_FETCH_CLASS_ALLOW_UNLINKED|ZEND_FETCH_CLASS_NO_AUTOLOAD)
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
		if (ce1.GetCeFlags() & ZEND_ACC_RESOLVED_INTERFACES) != 0 {

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
				var ce *ZendClassEntry = ZendLookupClassEx(ce1.interface_names[i].name, ce1.interface_names[i].lc_name, ZEND_FETCH_CLASS_ALLOW_UNLINKED|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil && UnlinkedInstanceof(ce, ce2) != 0 {
					return 1
				}
			}
		}
	}
	return 0
}
func ZendPerformCovariantTypeCheck(unresolved_class **ZendString, fe *ZendFunction, fe_arg_info *ZendArgInfo, proto *ZendFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	var fe_type ZendType = fe_arg_info.GetType()
	var proto_type ZendType = proto_arg_info.GetType()
	ZEND_ASSERT(ZEND_TYPE_IS_SET(fe_type) && ZEND_TYPE_IS_SET(proto_type))
	if ZEND_TYPE_ALLOW_NULL(fe_type) && !(ZEND_TYPE_ALLOW_NULL(proto_type)) {
		return INHERITANCE_ERROR
	}
	if ZEND_TYPE_IS_CLASS(proto_type) {
		var fe_class_name *ZendString
		var proto_class_name *ZendString
		var fe_ce *ZendClassEntry
		var proto_ce *ZendClassEntry
		if !(ZEND_TYPE_IS_CLASS(fe_type)) {
			return INHERITANCE_ERROR
		}
		fe_class_name = ResolveClassName(fe.GetScope(), ZEND_TYPE_NAME(fe_type))
		proto_class_name = ResolveClassName(proto.GetScope(), ZEND_TYPE_NAME(proto_type))
		if ZendStringEqualsCi(fe_class_name, proto_class_name) {
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
	} else if ZEND_TYPE_CODE(proto_type) == IS_ITERABLE {
		if ZEND_TYPE_IS_CLASS(fe_type) {
			var fe_class_name *ZendString = ResolveClassName(fe.GetScope(), ZEND_TYPE_NAME(fe_type))
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
		if ZEND_TYPE_CODE(fe_type) == IS_ITERABLE || ZEND_TYPE_CODE(fe_type) == IS_ARRAY {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else if ZEND_TYPE_CODE(proto_type) == IS_OBJECT {
		if ZEND_TYPE_IS_CLASS(fe_type) {

			/* Currently, any class name would be allowed here. We still perform a class lookup
			 * for forward-compatibility reasons, as we may have named types in the future that
			 * are not classes (such as enums or typedefs). */

			var fe_class_name *ZendString = ResolveClassName(fe.GetScope(), ZEND_TYPE_NAME(fe_type))
			var fe_ce *ZendClassEntry = LookupClass(fe.GetScope(), fe_class_name)
			if fe_ce == nil {
				*unresolved_class = fe_class_name
				return INHERITANCE_UNRESOLVED
			}
			return INHERITANCE_SUCCESS
		}
		if ZEND_TYPE_CODE(fe_type) == IS_OBJECT {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else {
		if ZEND_TYPE_CODE(fe_type) == ZEND_TYPE_CODE(proto_type) {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	}
}
func ZendDoPerformArgTypeHintCheck(unresolved_class **ZendString, fe *ZendFunction, fe_arg_info *ZendArgInfo, proto *ZendFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	if !(ZEND_TYPE_IS_SET(fe_arg_info.GetType())) {

		/* Child with no type is always compatible */

		return INHERITANCE_SUCCESS

		/* Child with no type is always compatible */

	}
	if !(ZEND_TYPE_IS_SET(proto_arg_info.GetType())) {

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
func ZendDoPerformImplementationCheck(unresolved_class **ZendString, fe *ZendFunction, proto *ZendFunction) InheritanceStatus {
	var i uint32
	var num_args uint32
	var status InheritanceStatus
	var local_status InheritanceStatus

	/* If it's a user function then arg_info == NULL means we don't have any parameters but
	 * we still need to do the arg number checks.  We are only willing to ignore this for internal
	 * functions because extensions don't always define arg_info.
	 */

	if proto.GetArgInfo() == nil && proto.GetCommonType() != ZEND_USER_FUNCTION {
		return INHERITANCE_SUCCESS
	}

	/* Checks for constructors only if they are declared in an interface,
	 * or explicitly marked as abstract
	 */

	ZEND_ASSERT(!((fe.GetFnFlags()&ZEND_ACC_CTOR) != 0 && ((proto.GetScope().GetCeFlags()&ZEND_ACC_INTERFACE) == 0 && (proto.GetFnFlags()&ZEND_ACC_ABSTRACT) == 0)))

	/* If the prototype method is private do not enforce a signature */

	ZEND_ASSERT((proto.GetFnFlags() & ZEND_ACC_PRIVATE) == 0)

	/* check number of arguments */

	if proto.GetRequiredNumArgs() < fe.GetRequiredNumArgs() || proto.GetNumArgs() > fe.GetNumArgs() {
		return INHERITANCE_ERROR
	}

	/* by-ref constraints on return values are covariant */

	if (proto.GetFnFlags()&ZEND_ACC_RETURN_REFERENCE) != 0 && (fe.GetFnFlags()&ZEND_ACC_RETURN_REFERENCE) == 0 {
		return INHERITANCE_ERROR
	}
	if (proto.GetFnFlags()&ZEND_ACC_VARIADIC) != 0 && (fe.GetFnFlags()&ZEND_ACC_VARIADIC) == 0 {
		return INHERITANCE_ERROR
	}

	/* For variadic functions any additional (optional) arguments that were added must be
	 * checked against the signature of the variadic argument, so in this case we have to
	 * go through all the parameters of the function and not just those present in the
	 * prototype. */

	num_args = proto.GetNumArgs()
	if (proto.GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
		num_args++
		if fe.GetNumArgs() >= proto.GetNumArgs() {
			num_args = fe.GetNumArgs()
			if (fe.GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
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
		if UNEXPECTED(local_status != INHERITANCE_SUCCESS) {
			if UNEXPECTED(local_status == INHERITANCE_ERROR) {
				return INHERITANCE_ERROR
			}
			ZEND_ASSERT(local_status == INHERITANCE_UNRESOLVED)
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

	if (proto.GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {

		/* Removing a return type is not valid. */

		if (fe.GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) == 0 {
			return INHERITANCE_ERROR
		}
		local_status = ZendPerformCovariantTypeCheck(unresolved_class, fe, fe.GetArgInfo()-1, proto, proto.GetArgInfo()-1)
		if UNEXPECTED(local_status != INHERITANCE_SUCCESS) {
			if UNEXPECTED(local_status == INHERITANCE_ERROR) {
				return INHERITANCE_ERROR
			}
			ZEND_ASSERT(local_status == INHERITANCE_UNRESOLVED)
			status = INHERITANCE_UNRESOLVED
		}
	}
	return status
}
func ZendAppendTypeHint(str *SmartStr, fptr *ZendFunction, arg_info *ZendArgInfo, return_hint int) {
	if ZEND_TYPE_IS_SET(arg_info.GetType()) && ZEND_TYPE_ALLOW_NULL(arg_info.GetType()) {
		SmartStrAppendc(str, '?')
	}
	if ZEND_TYPE_IS_CLASS(arg_info.GetType()) {
		var class_name *byte
		var class_name_len int
		class_name = ZSTR_VAL(ZEND_TYPE_NAME(arg_info.GetType()))
		class_name_len = ZSTR_LEN(ZEND_TYPE_NAME(arg_info.GetType()))
		if !(strcasecmp(class_name, "self")) && fptr.GetScope() != nil {
			class_name = ZSTR_VAL(fptr.GetScope().GetName())
			class_name_len = ZSTR_LEN(fptr.GetScope().GetName())
		} else if !(strcasecmp(class_name, "parent")) && fptr.GetScope() != nil && fptr.GetScope().parent {
			class_name = ZSTR_VAL(fptr.GetScope().parent.name)
			class_name_len = ZSTR_LEN(fptr.GetScope().parent.name)
		}
		SmartStrAppendl(str, class_name, class_name_len)
		if return_hint == 0 {
			SmartStrAppendc(str, ' ')
		}
	} else if ZEND_TYPE_IS_CODE(arg_info.GetType()) {
		var type_name *byte = ZendGetTypeByConst(ZEND_TYPE_CODE(arg_info.GetType()))
		SmartStrAppends(str, type_name)
		if return_hint == 0 {
			SmartStrAppendc(str, ' ')
		}
	}
}
func ZendGetFunctionDeclaration(fptr *ZendFunction) *ZendString {
	var str SmartStr = SmartStr{0}
	if (fptr.GetOpArray().GetFnFlags() & ZEND_ACC_RETURN_REFERENCE) != 0 {
		SmartStrAppends(&str, "& ")
	}
	if fptr.GetScope() != nil {

		/* cut off on NULL byte ... class@anonymous */

		SmartStrAppendl(&str, ZSTR_VAL(fptr.GetScope().GetName()), strlen(ZSTR_VAL(fptr.GetScope().GetName())))
		SmartStrAppends(&str, "::")
	}
	SmartStrAppend(&str, fptr.GetFunctionName())
	SmartStrAppendc(&str, '(')
	if fptr.GetArgInfo() != nil {
		var i uint32
		var num_args uint32
		var required uint32
		var arg_info *ZendArgInfo = fptr.GetArgInfo()
		required = fptr.GetRequiredNumArgs()
		num_args = fptr.GetNumArgs()
		if (fptr.GetFnFlags() & ZEND_ACC_VARIADIC) != 0 {
			num_args++
		}
		for i = 0; i < num_args; {
			ZendAppendTypeHint(&str, fptr, arg_info, 0)
			if arg_info.GetPassByReference() != 0 {
				SmartStrAppendc(&str, '&')
			}
			if arg_info.GetIsVariadic() != 0 {
				SmartStrAppends(&str, "...")
			}
			SmartStrAppendc(&str, '$')
			if arg_info.GetName() != nil {
				if fptr.GetType() == ZEND_INTERNAL_FUNCTION {
					SmartStrAppends(&str, (*ZendInternalArgInfo)(arg_info).GetName())
				} else {
					SmartStrAppendl(&str, ZSTR_VAL(arg_info.GetName()), ZSTR_LEN(arg_info.GetName()))
				}
			} else {
				SmartStrAppends(&str, "param")
				SmartStrAppendUnsigned(&str, i)
			}
			if i >= required && arg_info.GetIsVariadic() == 0 {
				SmartStrAppends(&str, " = ")
				if fptr.GetType() == ZEND_USER_FUNCTION {
					var precv *ZendOp = nil
					var idx uint32 = i
					var op *ZendOp = fptr.GetOpArray().GetOpcodes()
					var end *ZendOp = op + fptr.GetOpArray().GetLast()
					idx++
					for op < end {
						if (op.GetOpcode() == ZEND_RECV || op.GetOpcode() == ZEND_RECV_INIT) && op.GetOp1().GetNum() == ZendUlong(idx) {
							precv = op
						}
						op++
					}
					if precv != nil && precv.GetOpcode() == ZEND_RECV_INIT && precv.GetOp2Type() != IS_UNUSED {
						var zv *Zval = RT_CONSTANT(precv, precv.GetOp2())
						if Z_TYPE_P(zv) == IS_FALSE {
							SmartStrAppends(&str, "false")
						} else if Z_TYPE_P(zv) == IS_TRUE {
							SmartStrAppends(&str, "true")
						} else if Z_TYPE_P(zv) == IS_NULL {
							SmartStrAppends(&str, "NULL")
						} else if Z_TYPE_P(zv) == IS_STRING {
							SmartStrAppendc(&str, '\'')
							SmartStrAppendl(&str, Z_STRVAL_P(zv), MIN(Z_STRLEN_P(zv), 10))
							if Z_STRLEN_P(zv) > 10 {
								SmartStrAppends(&str, "...")
							}
							SmartStrAppendc(&str, '\'')
						} else if Z_TYPE_P(zv) == IS_ARRAY {
							SmartStrAppends(&str, "Array")
						} else if Z_TYPE_P(zv) == IS_CONSTANT_AST {
							var ast *ZendAst = Z_ASTVAL_P(zv)
							if ast.GetKind() == ZEND_AST_CONSTANT {
								SmartStrAppend(&str, ZendAstGetConstantName(ast))
							} else {
								SmartStrAppends(&str, "<expression>")
							}
						} else {
							var tmp_zv_str *ZendString
							var zv_str *ZendString = ZvalGetTmpString(zv, &tmp_zv_str)
							SmartStrAppend(&str, zv_str)
							ZendTmpStringRelease(tmp_zv_str)
						}
					}
				} else {
					SmartStrAppends(&str, "NULL")
				}
			}
			if b.PreInc(&i) < num_args {
				SmartStrAppends(&str, ", ")
			}
			arg_info++
		}
	}
	SmartStrAppendc(&str, ')')
	if (fptr.GetFnFlags() & ZEND_ACC_HAS_RETURN_TYPE) != 0 {
		SmartStrAppends(&str, ": ")
		ZendAppendTypeHint(&str, fptr, fptr.GetArgInfo()-1, 1)
	}
	SmartStr0(&str)
	return str.GetS()
}
func FuncLineno(fn *ZendFunction) uint32 {
	if fn.GetCommonType() == ZEND_USER_FUNCTION {
		return fn.GetOpArray().GetLineStart()
	} else {
		return 0
	}
}
func EmitIncompatibleMethodError(error_level int, error_verb *byte, child *ZendFunction, parent *ZendFunction, status InheritanceStatus, unresolved_class *ZendString) {
	var parent_prototype *ZendString = ZendGetFunctionDeclaration(parent)
	var child_prototype *ZendString = ZendGetFunctionDeclaration(child)
	if status == INHERITANCE_UNRESOLVED {
		ZendErrorAt(error_level, nil, FuncLineno(child), "Could not check compatibility between %s and %s, because class %s is not available", ZSTR_VAL(child_prototype), ZSTR_VAL(parent_prototype), ZSTR_VAL(unresolved_class))
	} else {
		ZendErrorAt(error_level, nil, FuncLineno(child), "Declaration of %s %s be compatible with %s", ZSTR_VAL(child_prototype), error_verb, ZSTR_VAL(parent_prototype))
	}
	ZendStringEfree(child_prototype)
	ZendStringEfree(parent_prototype)
}
func EmitIncompatibleMethodErrorOrWarning(child *ZendFunction, parent *ZendFunction, status InheritanceStatus, unresolved_class *ZendString, always_error ZendBool) {
	var error_level int
	var error_verb *byte
	if always_error != 0 || child.GetPrototype() != nil && (child.GetPrototype().GetFnFlags()&ZEND_ACC_ABSTRACT) != 0 || (parent.GetFnFlags()&ZEND_ACC_HAS_RETURN_TYPE) != 0 && ((child.GetFnFlags()&ZEND_ACC_HAS_RETURN_TYPE) == 0 || ZendPerformCovariantTypeCheck(&unresolved_class, child, child.GetArgInfo()-1, parent, parent.GetArgInfo()-1) != INHERITANCE_SUCCESS) {
		error_level = E_COMPILE_ERROR
		error_verb = "must"
	} else {
		error_level = E_WARNING
		error_verb = "should"
	}
	EmitIncompatibleMethodError(error_level, error_verb, child, parent, status, unresolved_class)
}
func PerformDelayableImplementationCheck(ce *ZendClassEntry, fe *ZendFunction, proto *ZendFunction, always_error ZendBool) {
	var unresolved_class *ZendString
	var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, fe, proto)
	if UNEXPECTED(status != INHERITANCE_SUCCESS) {
		if EXPECTED(status == INHERITANCE_UNRESOLVED) {
			AddCompatibilityObligation(ce, fe, proto, always_error)
		} else {
			ZEND_ASSERT(status == INHERITANCE_ERROR)
			if always_error != 0 {
				EmitIncompatibleMethodError(E_COMPILE_ERROR, "must", fe, proto, status, unresolved_class)
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
	if checked == 0 && UNEXPECTED((parent_flags&ZEND_ACC_FINAL) != 0) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot override final method %s::%s()", ZEND_FN_SCOPE_NAME(parent), ZSTR_VAL(child.GetFunctionName()))
	}
	child_flags = child.GetFnFlags()

	/* You cannot change from static to non static and vice versa.
	 */

	if checked == 0 && UNEXPECTED((child_flags&ZEND_ACC_STATIC) != (parent_flags&ZEND_ACC_STATIC)) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		if (child_flags & ZEND_ACC_STATIC) != 0 {
			ZendErrorAtNoreturn(E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non static method %s::%s() static in class %s", ZEND_FN_SCOPE_NAME(parent), ZSTR_VAL(child.GetFunctionName()), ZEND_FN_SCOPE_NAME(child))
		} else {
			ZendErrorAtNoreturn(E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make static method %s::%s() non static in class %s", ZEND_FN_SCOPE_NAME(parent), ZSTR_VAL(child.GetFunctionName()), ZEND_FN_SCOPE_NAME(child))
		}
	}

	/* Disallow making an inherited method abstract. */

	if checked == 0 && UNEXPECTED((child_flags&ZEND_ACC_ABSTRACT) > (parent_flags&ZEND_ACC_ABSTRACT)) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non abstract method %s::%s() abstract in class %s", ZEND_FN_SCOPE_NAME(parent), ZSTR_VAL(child.GetFunctionName()), ZEND_FN_SCOPE_NAME(child))
	}
	if check_only == 0 && (parent_flags&(ZEND_ACC_PRIVATE|ZEND_ACC_CHANGED)) != 0 {
		child.SetFnFlags(child.GetFnFlags() | ZEND_ACC_CHANGED)
	}
	if (parent_flags & ZEND_ACC_PRIVATE) != 0 {
		return INHERITANCE_SUCCESS
	}
	if parent.GetPrototype() != nil {
		proto = parent.GetPrototype()
	} else {
		proto = parent
	}
	if (parent_flags & ZEND_ACC_CTOR) != 0 {

		/* ctors only have a prototype if is abstract (or comes from an interface) */

		if (proto.GetFnFlags() & ZEND_ACC_ABSTRACT) == 0 {
			return INHERITANCE_SUCCESS
		}
		parent = proto
	}
	if check_only == 0 && child.GetPrototype() != proto {
		for {
			if child.GetScope() != ce && child.GetType() == ZEND_USER_FUNCTION && child.GetOpArray().GetStaticVariables() == nil {
				if (ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0 {

					/* Few parent interfaces contain the same method */

					break

					/* Few parent interfaces contain the same method */

				} else if child_zv != nil {

					/* op_array wasn't duplicated yet */

					var new_function *ZendFunction = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_op_array"))
					memcpy(new_function, child, b.SizeOf("zend_op_array"))
					child = new_function
					Z_PTR_P(child_zv) = child
				}
			}
			child.SetPrototype(proto)
			break
		}
	}

	/* Prevent derived classes from restricting access that was available in parent classes (except deriving from non-abstract ctors) */

	if checked == 0 && (child_flags&ZEND_ACC_PPP_MASK) > (parent_flags&ZEND_ACC_PPP_MASK) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		ZendErrorAtNoreturn(E_COMPILE_ERROR, nil, FuncLineno(child), "Access level to %s::%s() must be %s (as in class %s)%s", ZEND_FN_SCOPE_NAME(child), ZSTR_VAL(child.GetFunctionName()), ZendVisibilityString(parent_flags), ZEND_FN_SCOPE_NAME(parent), b.Cond((parent_flags&ZEND_ACC_PUBLIC) != 0, "", " or weaker"))
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
func DoInheritanceCheckOnMethod(child *ZendFunction, parent *ZendFunction, ce *ZendClassEntry, child_zv *Zval) {
	DoInheritanceCheckOnMethodEx(child, parent, ce, child_zv, 0, 0)
}
func DoInheritMethod(key *ZendString, parent *ZendFunction, ce *ZendClassEntry, is_interface ZendBool, checked ZendBool) {
	var child *Zval = ZendHashFindEx(&ce.function_table, key, 1)
	if child != nil {
		var func_ *ZendFunction = (*ZendFunction)(Z_PTR_P(child))
		if is_interface != 0 && UNEXPECTED(func_ == parent) {

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
		if is_interface != 0 || (parent.GetFnFlags()&ZEND_ACC_ABSTRACT) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
		}
		parent = ZendDuplicateFunction(parent, ce, is_interface)
		if is_interface == 0 {
			_zendHashAppendPtr(&ce.function_table, key, parent)
		} else {
			ZendHashAddNewPtr(&ce.function_table, key, parent)
		}
	}
}
func PropertyTypesCompatible(parent_info *ZendPropertyInfo, child_info *ZendPropertyInfo) InheritanceStatus {
	var parent_name *ZendString
	var child_name *ZendString
	var parent_type_ce *ZendClassEntry
	var child_type_ce *ZendClassEntry
	if parent_info.GetType() == child_info.GetType() {
		return INHERITANCE_SUCCESS
	}
	if !(ZEND_TYPE_IS_CLASS(parent_info.GetType())) || !(ZEND_TYPE_IS_CLASS(child_info.GetType())) || ZEND_TYPE_ALLOW_NULL(parent_info.GetType()) != ZEND_TYPE_ALLOW_NULL(child_info.GetType()) {
		return INHERITANCE_ERROR
	}
	if ZEND_TYPE_IS_CE(parent_info.GetType()) {
		parent_name = ZEND_TYPE_CE(parent_info.GetType()).GetName()
	} else {
		parent_name = ResolveClassName(parent_info.GetCe(), ZEND_TYPE_NAME(parent_info.GetType()))
	}
	if ZEND_TYPE_IS_CE(child_info.GetType()) {
		child_name = ZEND_TYPE_CE(child_info.GetType()).GetName()
	} else {
		child_name = ResolveClassName(child_info.GetCe(), ZEND_TYPE_NAME(child_info.GetType()))
	}
	if ZendStringEqualsCi(parent_name, child_name) {
		return INHERITANCE_SUCCESS
	}

	/* Check for class aliases */

	if ZEND_TYPE_IS_CE(parent_info.GetType()) {
		parent_type_ce = ZEND_TYPE_CE(parent_info.GetType())
	} else {
		parent_type_ce = LookupClass(parent_info.GetCe(), parent_name)
	}
	if ZEND_TYPE_IS_CE(child_info.GetType()) {
		child_type_ce = ZEND_TYPE_CE(child_info.GetType())
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
	ZendErrorNoreturn(E_COMPILE_ERROR, "Type of %s::$%s must be %s%s (as in class %s)", ZSTR_VAL(child.GetCe().GetName()), ZendGetUnmangledPropertyName(child.GetName()), b.Cond(ZEND_TYPE_ALLOW_NULL(parent.GetType()), "?", ""), b.CondF(ZEND_TYPE_IS_CLASS(parent.GetType()), func() []byte {
		return ZSTR_VAL(b.CondF(ZEND_TYPE_IS_CE(parent.GetType()), func() *ZendString { return ZEND_TYPE_CE(parent.GetType()).GetName() }, func() *ZendString { return ResolveClassName(parent.GetCe(), ZEND_TYPE_NAME(parent.GetType())) }))
	}, func() *byte { return ZendGetTypeByConst(ZEND_TYPE_CODE(parent.GetType())) }), ZSTR_VAL(parent.GetCe().GetName()))
}
func DoInheritProperty(parent_info *ZendPropertyInfo, key *ZendString, ce *ZendClassEntry) {
	var child *Zval = ZendHashFindEx(&ce.properties_info, key, 1)
	var child_info *ZendPropertyInfo
	if UNEXPECTED(child != nil) {
		child_info = Z_PTR_P(child)
		if (parent_info.GetFlags() & (ZEND_ACC_PRIVATE | ZEND_ACC_CHANGED)) != 0 {
			child_info.SetFlags(child_info.GetFlags() | ZEND_ACC_CHANGED)
		}
		if (parent_info.GetFlags() & ZEND_ACC_PRIVATE) == 0 {
			if UNEXPECTED((parent_info.GetFlags() & ZEND_ACC_STATIC) != (child_info.GetFlags() & ZEND_ACC_STATIC)) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot redeclare %s%s::$%s as %s%s::$%s", b.Cond((parent_info.GetFlags()&ZEND_ACC_STATIC) != 0, "static ", "non static "), ZSTR_VAL(ce.parent.name), ZSTR_VAL(key), b.Cond((child_info.GetFlags()&ZEND_ACC_STATIC) != 0, "static ", "non static "), ZSTR_VAL(ce.GetName()), ZSTR_VAL(key))
			}
			if UNEXPECTED((child_info.GetFlags() & ZEND_ACC_PPP_MASK) > (parent_info.GetFlags() & ZEND_ACC_PPP_MASK)) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Access level to %s::$%s must be %s (as in class %s)%s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(key), ZendVisibilityString(parent_info.GetFlags()), ZSTR_VAL(ce.parent.name), b.Cond((parent_info.GetFlags()&ZEND_ACC_PUBLIC) != 0, "", " or weaker"))
			} else if (child_info.GetFlags() & ZEND_ACC_STATIC) == 0 {
				var parent_num int = OBJ_PROP_TO_NUM(parent_info.GetOffset())
				var child_num int = OBJ_PROP_TO_NUM(child_info.GetOffset())

				/* Don't keep default properties in GC (they may be freed by opcache) */

				ZvalPtrDtorNogc(&ce.GetDefaultPropertiesTable()[parent_num])
				ce.GetDefaultPropertiesTable()[parent_num] = ce.GetDefaultPropertiesTable()[child_num]
				ZVAL_UNDEF(&ce.default_properties_table[child_num])
				child_info.SetOffset(parent_info.GetOffset())
			}
			if UNEXPECTED(ZEND_TYPE_IS_SET(parent_info.GetType())) {
				var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
				if status == INHERITANCE_ERROR {
					EmitIncompatiblePropertyError(child_info, parent_info)
				}
				if status == INHERITANCE_UNRESOLVED {
					AddPropertyCompatibilityObligation(ce, child_info, parent_info)
				}
			} else if UNEXPECTED(ZEND_TYPE_IS_SET(child_info.GetType()) && !(ZEND_TYPE_IS_SET(parent_info.GetType()))) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Type of %s::$%s must not be defined (as in class %s)", ZSTR_VAL(ce.GetName()), ZSTR_VAL(key), ZSTR_VAL(ce.parent.name))
			}
		}
	} else {
		if UNEXPECTED((ce.GetType() & ZEND_INTERNAL_CLASS) != 0) {
			child_info = ZendDuplicatePropertyInfoInternal(parent_info)
		} else {
			child_info = parent_info
		}
		_zendHashAppendPtr(&ce.properties_info, key, child_info)
	}
}
func DoImplementInterface(ce *ZendClassEntry, iface *ZendClassEntry) {
	if (ce.GetCeFlags()&ZEND_ACC_INTERFACE) == 0 && iface.interface_gets_implemented && iface.interface_gets_implemented(iface, ce) == FAILURE {
		ZendErrorNoreturn(E_CORE_ERROR, "Class %s could not implement interface %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(iface.GetName()))
	}

	/* This should be prevented by the class lookup logic. */

	ZEND_ASSERT(ce != iface)

	/* This should be prevented by the class lookup logic. */
}
func ZendDoInheritInterfaces(ce *ZendClassEntry, iface *ZendClassEntry) {
	/* expects interface to be contained in ce's interface list already */

	var i uint32
	var ce_num uint32
	var if_num uint32 = iface.GetNumInterfaces()
	var entry *ZendClassEntry
	ce_num = ce.GetNumInterfaces()
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		ce.interfaces = (**ZendClassEntry)(realloc(ce.interfaces, b.SizeOf("zend_class_entry *")*(ce_num+if_num)))
	} else {
		ce.interfaces = (**ZendClassEntry)(Erealloc(ce.interfaces, b.SizeOf("zend_class_entry *")*(ce_num+if_num)))
	}

	/* Inherit the interfaces, only if they're not already inherited by the class */

	for b.PostDec(&if_num) {
		entry = iface.interfaces[if_num]
		for i = 0; i < ce_num; i++ {
			if ce.interfaces[i] == entry {
				break
			}
		}
		if i == ce_num {
			ce.interfaces[b.PostInc(&(ce.GetNumInterfaces()))] = entry
		}
	}
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_RESOLVED_INTERFACES)

	/* and now call the implementing handlers */

	for ce_num < ce.GetNumInterfaces() {
		DoImplementInterface(ce, ce.interfaces[b.PostInc(&ce_num)])
	}

	/* and now call the implementing handlers */
}
func DoInheritClassConstant(name *ZendString, parent_const *ZendClassConstant, ce *ZendClassEntry) {
	var zv *Zval = ZendHashFindEx(&ce.constants_table, name, 1)
	var c *ZendClassConstant
	if zv != nil {
		c = (*ZendClassConstant)(Z_PTR_P(zv))
		if UNEXPECTED((Z_ACCESS_FLAGS(c.GetValue()) & ZEND_ACC_PPP_MASK) > (Z_ACCESS_FLAGS(parent_const.GetValue()) & ZEND_ACC_PPP_MASK)) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Access level to %s::%s must be %s (as in class %s)%s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(name), ZendVisibilityString(Z_ACCESS_FLAGS(parent_const.GetValue())), ZSTR_VAL(ce.parent.name), b.Cond((Z_ACCESS_FLAGS(parent_const.GetValue())&ZEND_ACC_PUBLIC) != 0, "", " or weaker"))
		}
	} else if (Z_ACCESS_FLAGS(parent_const.GetValue()) & ZEND_ACC_PRIVATE) == 0 {
		if Z_TYPE(parent_const.GetValue()) == IS_CONSTANT_AST {
			ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_CONSTANTS_UPDATED)
		}
		if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
			c = Pemalloc(b.SizeOf("zend_class_constant"), 1)
			memcpy(c, parent_const, b.SizeOf("zend_class_constant"))
			parent_const = c
		}
		_zendHashAppendPtr(&ce.constants_table, name, parent_const)
	}
}
func ZendBuildPropertiesInfoTable(ce *ZendClassEntry) {
	var table **ZendPropertyInfo
	var prop **ZendPropertyInfo
	var size int
	if ce.GetDefaultPropertiesCount() == 0 {
		return
	}
	ZEND_ASSERT(ce.GetPropertiesInfoTable() == nil)
	size = b.SizeOf("zend_property_info *") * ce.GetDefaultPropertiesCount()
	if ce.GetType() == ZEND_USER_CLASS {
		table = ZendArenaAlloc(&(CompilerGlobals.GetArena()), size)
		ce.SetPropertiesInfoTable(table)
	} else {
		table = Pemalloc(size, 1)
		ce.SetPropertiesInfoTable(table)
	}

	/* Dead slots may be left behind during inheritance. Make sure these are NULLed out. */

	memset(table, 0, size)
	if ce.parent && ce.parent.default_properties_count != 0 {
		var parent_table **ZendPropertyInfo = ce.parent.properties_info_table
		memcpy(table, parent_table, b.SizeOf("zend_property_info *")*ce.parent.default_properties_count)

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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			prop = Z_PTR_P(_z)
			if prop.ce == ce && (prop.flags&ZEND_ACC_STATIC) == 0 {
				table[OBJ_PROP_TO_NUM(prop.offset)] = prop
			}
		}
		break
	}
}
func ZendDoInheritanceEx(ce *ZendClassEntry, parent_ce *ZendClassEntry, checked ZendBool) {
	var property_info *ZendPropertyInfo
	var func_ *ZendFunction
	var key *ZendString
	if UNEXPECTED((ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0) {

		/* Interface can only inherit other interfaces */

		if UNEXPECTED((parent_ce.GetCeFlags() & ZEND_ACC_INTERFACE) == 0) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Interface %s may not inherit from class (%s)", ZSTR_VAL(ce.GetName()), ZSTR_VAL(parent_ce.GetName()))
		}

		/* Interface can only inherit other interfaces */

	} else if UNEXPECTED((parent_ce.GetCeFlags() & (ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_FINAL)) != 0) {

		/* Class declaration must not extend traits or interfaces */

		if (parent_ce.GetCeFlags() & ZEND_ACC_INTERFACE) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s cannot extend from interface %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(parent_ce.GetName()))
		} else if (parent_ce.GetCeFlags() & ZEND_ACC_TRAIT) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s cannot extend from trait %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(parent_ce.GetName()))
		}

		/* Class must not extend a final class */

		if (parent_ce.GetCeFlags() & ZEND_ACC_FINAL) != 0 {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s may not inherit from final class (%s)", ZSTR_VAL(ce.GetName()), ZSTR_VAL(parent_ce.GetName()))
		}

		/* Class must not extend a final class */

	}
	if ce.parent_name {
		ZendStringReleaseEx(ce.parent_name, 0)
	}
	ce.parent = parent_ce
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_RESOLVED_PARENT)

	/* Inherit interfaces */

	if parent_ce.GetNumInterfaces() != 0 {
		if (ce.GetCeFlags() & ZEND_ACC_IMPLEMENT_INTERFACES) == 0 {
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
			var table *Zval = Pemalloc(b.SizeOf("zval")*(ce.GetDefaultPropertiesCount()+parent_ce.GetDefaultPropertiesCount()), ce.GetType() == ZEND_INTERNAL_CLASS)
			src = ce.GetDefaultPropertiesTable() + ce.GetDefaultPropertiesCount()
			end = table + parent_ce.GetDefaultPropertiesCount()
			dst = end + ce.GetDefaultPropertiesCount()
			ce.SetDefaultPropertiesTable(table)
			for {
				dst--
				src--
				ZVAL_COPY_VALUE_PROP(dst, src)
				if dst == end {
					break
				}
			}
			Pefree(src, ce.GetType() == ZEND_INTERNAL_CLASS)
			end = ce.GetDefaultPropertiesTable()
		} else {
			end = Pemalloc(b.SizeOf("zval")*parent_ce.GetDefaultPropertiesCount(), ce.GetType() == ZEND_INTERNAL_CLASS)
			dst = end + parent_ce.GetDefaultPropertiesCount()
			ce.SetDefaultPropertiesTable(end)
		}
		src = parent_ce.GetDefaultPropertiesTable() + parent_ce.GetDefaultPropertiesCount()
		if UNEXPECTED(parent_ce.GetType() != ce.GetType()) {

			/* User class extends internal */

			for {
				dst--
				src--
				ZVAL_COPY_OR_DUP_PROP(dst, src)
				if Z_OPT_TYPE_P(dst) == IS_CONSTANT_AST {
					ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_CONSTANTS_UPDATED)
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
				ZVAL_COPY_PROP(dst, src)
				if Z_OPT_TYPE_P(dst) == IS_CONSTANT_AST {
					ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_CONSTANTS_UPDATED)
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
			var table *Zval = Pemalloc(b.SizeOf("zval")*(ce.GetDefaultStaticMembersCount()+parent_ce.GetDefaultStaticMembersCount()), ce.GetType() == ZEND_INTERNAL_CLASS)
			src = ce.GetDefaultStaticMembersTable() + ce.GetDefaultStaticMembersCount()
			end = table + parent_ce.GetDefaultStaticMembersCount()
			dst = end + ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(table)
			for {
				dst--
				src--
				ZVAL_COPY_VALUE(dst, src)
				if dst == end {
					break
				}
			}
			Pefree(src, ce.GetType() == ZEND_INTERNAL_CLASS)
			end = ce.GetDefaultStaticMembersTable()
		} else {
			end = Pemalloc(b.SizeOf("zval")*parent_ce.GetDefaultStaticMembersCount(), ce.GetType() == ZEND_INTERNAL_CLASS)
			dst = end + parent_ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(end)
		}
		if UNEXPECTED(parent_ce.GetType() != ce.GetType()) {

			/* User class extends internal */

			if CE_STATIC_MEMBERS(parent_ce) == nil {
				ZendClassInitStatics(parent_ce)
			}
			if UNEXPECTED(ZendUpdateClassConstants(parent_ce) != SUCCESS) {
				ZEND_ASSERT(false)
			}
			src = CE_STATIC_MEMBERS(parent_ce) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if Z_TYPE_P(src) == IS_INDIRECT {
					ZVAL_INDIRECT(dst, Z_INDIRECT_P(src))
				} else {
					ZVAL_INDIRECT(dst, src)
				}
				if dst == end {
					break
				}
			}
		} else if ce.GetType() == ZEND_USER_CLASS {
			if CE_STATIC_MEMBERS(parent_ce) == nil {
				ZEND_ASSERT((parent_ce.GetCeFlags() & (ZEND_ACC_IMMUTABLE | ZEND_ACC_PRELOADED)) != 0)
				ZendClassInitStatics(parent_ce)
			}
			src = CE_STATIC_MEMBERS(parent_ce) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if Z_TYPE_P(src) == IS_INDIRECT {
					ZVAL_INDIRECT(dst, Z_INDIRECT_P(src))
				} else {
					ZVAL_INDIRECT(dst, src)
				}
				if Z_TYPE_P(Z_INDIRECT_P(dst)) == IS_CONSTANT_AST {
					ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_CONSTANTS_UPDATED)
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
				if Z_TYPE_P(src) == IS_INDIRECT {
					ZVAL_INDIRECT(dst, Z_INDIRECT_P(src))
				} else {
					ZVAL_INDIRECT(dst, src)
				}
				if dst == end {
					break
				}
			}
		}
		ce.SetDefaultStaticMembersCount(ce.GetDefaultStaticMembersCount() + parent_ce.GetDefaultStaticMembersCount())
		if ce.GetStaticMembersTablePtr() == nil {
			ZEND_ASSERT(ce.GetType() == ZEND_INTERNAL_CLASS)
			if ExecutorGlobals.GetCurrentExecuteData() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {

				/* internal class loaded by dl() */

				ZEND_MAP_PTR_INIT(ce.static_members_table, &ce.default_static_members_table)

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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			property_info = Z_PTR_P(_z)
			if property_info.GetCe() == ce {
				if (property_info.GetFlags() & ZEND_ACC_STATIC) != 0 {
					property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultStaticMembersCount())
				} else {
					property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultPropertiesCount()*b.SizeOf("zval"))
				}
			}
		}
		break
	}
	if ZendHashNumElements(&parent_ce.properties_info) {
		ZendHashExtend(&ce.properties_info, ZendHashNumElements(&ce.properties_info)+ZendHashNumElements(&parent_ce.properties_info), 0)
		for {
			var __ht *HashTable = &parent_ce.properties_info
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				property_info = Z_PTR_P(_z)
				DoInheritProperty(property_info, key, ce)
			}
			break
		}
	}
	if ZendHashNumElements(&parent_ce.constants_table) {
		var c *ZendClassConstant
		ZendHashExtend(&ce.constants_table, ZendHashNumElements(&ce.constants_table)+ZendHashNumElements(&parent_ce.constants_table), 0)
		for {
			var __ht *HashTable = &parent_ce.constants_table
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				c = Z_PTR_P(_z)
				DoInheritClassConstant(key, c, ce)
			}
			break
		}
	}
	if ZendHashNumElements(&parent_ce.function_table) {
		ZendHashExtend(&ce.function_table, ZendHashNumElements(&ce.function_table)+ZendHashNumElements(&parent_ce.function_table), 0)
		if checked != 0 {
			for {
				var __ht *HashTable = &parent_ce.function_table
				var _p *Bucket = __ht.GetArData()
				var _end *Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *Zval = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					key = _p.GetKey()
					func_ = Z_PTR_P(_z)
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

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					key = _p.GetKey()
					func_ = Z_PTR_P(_z)
					DoInheritMethod(key, func_, ce, 0, 0)
				}
				break
			}
		}
	}
	DoInheritParentConstructor(ce)
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		if (ce.GetCeFlags() & ZEND_ACC_IMPLICIT_ABSTRACT_CLASS) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)
		}
	}
	ce.SetCeFlags(ce.GetCeFlags() | parent_ce.GetCeFlags()&(ZEND_HAS_STATIC_IN_METHODS|ZEND_ACC_HAS_TYPE_HINTS|ZEND_ACC_USE_GUARDS))
}
func DoInheritConstantCheck(child_constants_table *HashTable, parent_constant *ZendClassConstant, name *ZendString, iface *ZendClassEntry) ZendBool {
	var zv *Zval = ZendHashFindEx(child_constants_table, name, 1)
	var old_constant *ZendClassConstant
	if zv != nil {
		old_constant = (*ZendClassConstant)(Z_PTR_P(zv))
		if old_constant.GetCe() != parent_constant.GetCe() {
			ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot inherit previously-inherited or override constant %s from interface %s", ZSTR_VAL(name), ZSTR_VAL(iface.GetName()))
		}
		return 0
	}
	return 1
}
func DoInheritIfaceConstant(name *ZendString, c *ZendClassConstant, ce *ZendClassEntry, iface *ZendClassEntry) {
	if DoInheritConstantCheck(&ce.constants_table, c, name, iface) != 0 {
		var ct *ZendClassConstant
		if Z_TYPE(c.GetValue()) == IS_CONSTANT_AST {
			ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_CONSTANTS_UPDATED)
		}
		if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
			ct = Pemalloc(b.SizeOf("zend_class_constant"), 1)
			memcpy(ct, c, b.SizeOf("zend_class_constant"))
			c = ct
		}
		ZendHashUpdatePtr(&ce.constants_table, name, c)
	}
}
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			c = Z_PTR_P(_z)
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			func_ = Z_PTR_P(_z)
			DoInheritMethod(key, func_, ce, 1, 0)
		}
		break
	}
	DoImplementInterface(ce, iface)
	if iface.GetNumInterfaces() != 0 {
		ZendDoInheritInterfaces(ce, iface)
	}
}
func ZendDoImplementInterface(ce *ZendClassEntry, iface *ZendClassEntry) {
	var i uint32
	var ignore uint32 = 0
	var current_iface_num uint32 = ce.GetNumInterfaces()
	var parent_iface_num uint32 = b.CondF1(ce.parent, func() __auto__ { return ce.parent.num_interfaces }, 0)
	var key *ZendString
	var c *ZendClassConstant
	ZEND_ASSERT((ce.GetCeFlags() & ZEND_ACC_LINKED) != 0)
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		if ce.interfaces[i] == nil {
			memmove(ce.interfaces+i, ce.interfaces+i+1, b.SizeOf("zend_class_entry *")*(b.PreDec(&(ce.GetNumInterfaces()))-i))
			i--
		} else if ce.interfaces[i] == iface {
			if EXPECTED(i < parent_iface_num) {
				ignore = 1
			} else {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(iface.GetName()))
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

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				key = _p.GetKey()
				c = Z_PTR_P(_z)
				DoInheritConstantCheck(&iface.constants_table, c, key, iface)
			}
			break
		}

		/* Check for attempt to redeclare interface constants */

	} else {
		if ce.GetNumInterfaces() >= current_iface_num {
			if ce.GetType() == ZEND_INTERNAL_CLASS {
				ce.interfaces = (**ZendClassEntry)(realloc(ce.interfaces, b.SizeOf("zend_class_entry *")*b.PreInc(&current_iface_num)))
			} else {
				ce.interfaces = (**ZendClassEntry)(Erealloc(ce.interfaces, b.SizeOf("zend_class_entry *")*b.PreInc(&current_iface_num)))
			}
		}
		ce.interfaces[b.PostInc(&(ce.GetNumInterfaces()))] = iface
		DoInterfaceImplementation(ce, iface)
	}
}
func ZendDoImplementInterfaces(ce *ZendClassEntry, interfaces **ZendClassEntry) {
	var iface *ZendClassEntry
	var num_parent_interfaces uint32 = b.CondF1(ce.parent, func() __auto__ { return ce.parent.num_interfaces }, 0)
	var num_interfaces uint32 = num_parent_interfaces
	var key *ZendString
	var c *ZendClassConstant
	var i uint32
	var j uint32
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		iface = interfaces[num_parent_interfaces+i]
		if (iface.GetCeFlags() & ZEND_ACC_LINKED) == 0 {
			AddDependencyObligation(ce, iface)
		}
		if UNEXPECTED((iface.GetCeFlags() & ZEND_ACC_INTERFACE) == 0) {
			Efree(interfaces)
			ZendErrorNoreturn(E_ERROR, "%s cannot implement %s - it is not an interface", ZSTR_VAL(ce.GetName()), ZSTR_VAL(iface.GetName()))
			return
		}
		for j = 0; j < num_interfaces; j++ {
			if interfaces[j] == iface {
				if j >= num_parent_interfaces {
					Efree(interfaces)
					ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(iface.GetName()))
					return
				}

				/* skip duplications */

				for {
					var __ht *HashTable = &ce.constants_table
					var _p *Bucket = __ht.GetArData()
					var _end *Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *Zval = &_p.val

						if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
							continue
						}
						key = _p.GetKey()
						c = Z_PTR_P(_z)
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
	Efree(ce.interface_names)
	ce.SetNumInterfaces(num_interfaces)
	ce.interfaces = interfaces
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_RESOLVED_INTERFACES)
	i = num_parent_interfaces
	for ; i < ce.GetNumInterfaces(); i++ {
		DoInterfaceImplementation(ce, ce.interfaces[i])
	}
}
func ZendAddMagicMethods(ce *ZendClassEntry, mname *ZendString, fe *ZendFunction) {
	if ZendStringEqualsLiteral(mname, "serialize") {
		ce.SetSerializeFunc(fe)
	} else if ZendStringEqualsLiteral(mname, "unserialize") {
		ce.SetUnserializeFunc(fe)
	} else if ZSTR_LEN(ce.GetName()) != ZSTR_LEN(mname) && (ZSTR_VAL(mname)[0] != '_' || ZSTR_VAL(mname)[1] != '_') {

	} else if ZendStringEqualsLiteral(mname, ZEND_CLONE_FUNC_NAME) {
		ce.SetClone(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_CONSTRUCTOR_FUNC_NAME) {
		if ce.GetConstructor() != nil && (!(ce.parent) || ce.GetConstructor() != ce.parent.constructor) {
			ZendErrorNoreturn(E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ZSTR_VAL(ce.GetName()))
		}
		ce.SetConstructor(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_DESTRUCTOR_FUNC_NAME) {
		ce.SetDestructor(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_GET_FUNC_NAME) {
		ce.SetGet(fe)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
	} else if ZendStringEqualsLiteral(mname, ZEND_SET_FUNC_NAME) {
		ce.SetSet(fe)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
	} else if ZendStringEqualsLiteral(mname, ZEND_CALL_FUNC_NAME) {
		ce.SetCall(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_UNSET_FUNC_NAME) {
		ce.SetUnset(fe)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
	} else if ZendStringEqualsLiteral(mname, ZEND_ISSET_FUNC_NAME) {
		ce.SetIsset(fe)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_USE_GUARDS)
	} else if ZendStringEqualsLiteral(mname, ZEND_CALLSTATIC_FUNC_NAME) {
		ce.SetCallstatic(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_TOSTRING_FUNC_NAME) {
		ce.SetTostring(fe)
	} else if ZendStringEqualsLiteral(mname, ZEND_DEBUGINFO_FUNC_NAME) {
		ce.SetDebugInfo(fe)
	} else if ZSTR_LEN(ce.GetName()) == ZSTR_LEN(mname) {
		var lowercase_name *ZendString = ZendStringTolower(ce.GetName())
		lowercase_name = ZendNewInternedString(lowercase_name)
		if !(memcmp(ZSTR_VAL(mname), ZSTR_VAL(lowercase_name), ZSTR_LEN(mname))) {
			if ce.GetConstructor() != nil && (!(ce.parent) || ce.GetConstructor() != ce.parent.constructor) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ZSTR_VAL(ce.GetName()))
			}
			ce.SetConstructor(fe)
			fe.SetFnFlags(fe.GetFnFlags() | ZEND_ACC_CTOR)
		}
		ZendStringReleaseEx(lowercase_name, 0)
	}
}
func ZendAddTraitMethod(ce *ZendClassEntry, name *byte, key *ZendString, fn *ZendFunction, overridden **HashTable) {
	var existing_fn *ZendFunction = nil
	var new_fn *ZendFunction
	if b.Assign(&existing_fn, ZendHashFindPtr(&ce.function_table, key)) != nil {

		/* if it is the same function with the same visibility and has not been assigned a class scope yet, regardless
		 * of where it is coming from there is no conflict and we do not need to add it again */

		if existing_fn.GetOpArray().GetOpcodes() == fn.GetOpArray().GetOpcodes() && (existing_fn.GetFnFlags()&ZEND_ACC_PPP_MASK) == (fn.GetFnFlags()&ZEND_ACC_PPP_MASK) && (existing_fn.GetScope().GetCeFlags()&ZEND_ACC_TRAIT) == ZEND_ACC_TRAIT {
			return
		}
		if existing_fn.GetScope() == ce {

			/* members from the current class override trait methods */

			if (*overridden) != nil {
				if b.Assign(&existing_fn, ZendHashFindPtr(*overridden, key)) != nil {
					if (existing_fn.GetFnFlags() & ZEND_ACC_ABSTRACT) != 0 {

						/* Make sure the trait method is compatible with previosly declared abstract method */

						PerformDelayableImplementationCheck(ce, fn, existing_fn, 1)

						/* Make sure the trait method is compatible with previosly declared abstract method */

					}
					if (fn.GetFnFlags() & ZEND_ACC_ABSTRACT) != 0 {

						/* Make sure the abstract declaration is compatible with previous declaration */

						PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
						return
					}
				}
			} else {
				ALLOC_HASHTABLE(*overridden)
				ZendHashInitEx(*overridden, 8, nil, OverriddenPtrDtor, 0, 0)
			}
			ZendHashUpdateMem(*overridden, key, fn, b.SizeOf("zend_function"))
			return
		} else if (fn.GetFnFlags()&ZEND_ACC_ABSTRACT) != 0 && (existing_fn.GetFnFlags()&ZEND_ACC_ABSTRACT) == 0 {

			/* Make sure the abstract declaration is compatible with previous declaration */

			PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
			return
		} else if UNEXPECTED((existing_fn.GetScope().GetCeFlags()&ZEND_ACC_TRAIT) != 0 && (existing_fn.GetFnFlags()&ZEND_ACC_ABSTRACT) == 0) {

			/* two traits can't define the __special__  same non-abstract method */

			ZendErrorNoreturn(E_COMPILE_ERROR, "Trait method %s has not been applied, because there are collisions with other trait methods on %s", name, ZSTR_VAL(ce.GetName()))

			/* two traits can't define the __special__  same non-abstract method */

		} else {

			/* inherited members are overridden by members inserted by traits */

			DoInheritanceCheckOnMethod(fn, existing_fn, ce, nil)
			fn.SetPrototype(nil)
		}
	}
	if UNEXPECTED(fn.GetType() == ZEND_INTERNAL_FUNCTION) {
		new_fn = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_internal_function"))
		memcpy(new_fn, fn, b.SizeOf("zend_internal_function"))
		new_fn.SetFnFlags(new_fn.GetFnFlags() | ZEND_ACC_ARENA_ALLOCATED)
	} else {
		new_fn = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_op_array"))
		memcpy(new_fn, fn, b.SizeOf("zend_op_array"))
		new_fn.GetOpArray().SetFnFlags(new_fn.GetOpArray().GetFnFlags() | ZEND_ACC_TRAIT_CLONE)
		new_fn.GetOpArray().SetFnFlags(new_fn.GetOpArray().GetFnFlags() &^ ZEND_ACC_IMMUTABLE)
	}
	FunctionAddRef(new_fn)
	fn = ZendHashUpdatePtr(&ce.function_table, key, new_fn)
	ZendAddMagicMethods(ce, key, fn)
}
func ZendFixupTraitMethod(fn *ZendFunction, ce *ZendClassEntry) {
	if (fn.GetScope().GetCeFlags() & ZEND_ACC_TRAIT) == ZEND_ACC_TRAIT {
		fn.SetScope(ce)
		if (fn.GetFnFlags() & ZEND_ACC_ABSTRACT) != 0 {
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
		}
		if fn.GetType() == ZEND_USER_FUNCTION && fn.GetOpArray().GetStaticVariables() != nil {
			ce.SetCeFlags(ce.GetCeFlags() | ZEND_HAS_STATIC_IN_METHODS)
		}
	}
}
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

			if alias.GetAlias() != nil && (aliases[i] == nil || fn.GetScope() == aliases[i]) && ZSTR_LEN(alias.GetTraitMethod().GetMethodName()) == ZSTR_LEN(fnname) && ZendBinaryStrcasecmp(ZSTR_VAL(alias.GetTraitMethod().GetMethodName()), ZSTR_LEN(alias.GetTraitMethod().GetMethodName()), ZSTR_VAL(fnname), ZSTR_LEN(fnname)) == 0 {
				fn_copy = *fn

				/* if it is 0, no modifieres has been changed */

				if alias.GetModifiers() != 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&ZEND_ACC_PPP_MASK)
				}
				lcname = ZendStringTolower(alias.GetAlias())
				ZendAddTraitMethod(ce, ZSTR_VAL(alias.GetAlias()), lcname, &fn_copy, overridden)
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

		memcpy(&fn_copy, fn, b.CondF(fn.GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))

		/* apply aliases which have not alias name, just setting visibility */

		if ce.GetTraitAliases() != nil {
			alias_ptr = ce.GetTraitAliases()
			alias = *alias_ptr
			i = 0
			for alias != nil {

				/* Scope unset or equal to the function we compare to, and the alias applies to fn */

				if alias.GetAlias() == nil && alias.GetModifiers() != 0 && (aliases[i] == nil || fn.GetScope() == aliases[i]) && ZSTR_LEN(alias.GetTraitMethod().GetMethodName()) == ZSTR_LEN(fnname) && ZendBinaryStrcasecmp(ZSTR_VAL(alias.GetTraitMethod().GetMethodName()), ZSTR_LEN(alias.GetTraitMethod().GetMethodName()), ZSTR_VAL(fnname), ZSTR_LEN(fnname)) == 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&ZEND_ACC_PPP_MASK)

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
		ZendAddTraitMethod(ce, ZSTR_VAL(fn.GetFunctionName()), fnname, &fn_copy, overridden)
	}
}
func ZendCheckTraitUsage(ce *ZendClassEntry, trait *ZendClassEntry, traits **ZendClassEntry) uint32 {
	var i uint32
	if UNEXPECTED((trait.GetCeFlags() & ZEND_ACC_TRAIT) != ZEND_ACC_TRAIT) {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Class %s is not a trait, Only traits may be used in 'as' and 'insteadof' statements", ZSTR_VAL(trait.GetName()))
		return 0
	}
	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == trait {
			return i
		}
	}
	ZendErrorNoreturn(E_COMPILE_ERROR, "Required Trait %s wasn't added to %s", ZSTR_VAL(trait.GetName()), ZSTR_VAL(ce.GetName()))
	return 0
}
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
		exclude_tables = Ecalloc(ce.GetNumTraits(), b.SizeOf("HashTable *"))
		i = 0
		precedences = ce.GetTraitPrecedences()
		ce.SetTraitPrecedences(nil)
		for b.Assign(&cur_precedence, precedences[i]) {

			/** Resolve classes for all precedence operations. */

			cur_method_ref = &cur_precedence.trait_method
			trait = ZendFetchClass(cur_method_ref.GetClassName(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if trait == nil {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Could not find trait %s", ZSTR_VAL(cur_method_ref.GetClassName()))
			}
			ZendCheckTraitUsage(ce, trait, traits)

			/** Ensure that the preferred method is actually available. */

			lcname = ZendStringTolower(cur_method_ref.GetMethodName())
			if ZendHashExists(&trait.function_table, lcname) == 0 {
				ZendErrorNoreturn(E_COMPILE_ERROR, "A precedence rule was defined for %s::%s but this method does not exist", ZSTR_VAL(trait.GetName()), ZSTR_VAL(cur_method_ref.GetMethodName()))
			}

			/** With the other traits, we are more permissive.
			  We do not give errors for those. This allows to be more
			  defensive in such definitions.
			  However, we want to make sure that the insteadof declaration
			  is consistent in itself.
			*/

			for j = 0; j < cur_precedence.GetNumExcludes(); j++ {
				var class_name *ZendString = cur_precedence.GetExcludeClassNames()[j]
				var exclude_ce *ZendClassEntry = ZendFetchClass(class_name, ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				var trait_num uint32
				if exclude_ce == nil {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Could not find trait %s", ZSTR_VAL(class_name))
				}
				trait_num = ZendCheckTraitUsage(ce, exclude_ce, traits)
				if exclude_tables[trait_num] == nil {
					ALLOC_HASHTABLE(exclude_tables[trait_num])
					ZendHashInit(exclude_tables[trait_num], 0, nil, nil, 0)
				}
				if ZendHashAddEmptyElement(exclude_tables[trait_num], lcname) == nil {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Failed to evaluate a trait precedence (%s). Method of trait %s was defined to be excluded multiple times", ZSTR_VAL(precedences[i].GetTraitMethod().GetMethodName()), ZSTR_VAL(exclude_ce.GetName()))
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

				if trait == exclude_ce {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Inconsistent insteadof definition. "+"The method %s is to be used from %s, but %s is also on the exclude list", ZSTR_VAL(cur_method_ref.GetMethodName()), ZSTR_VAL(trait.GetName()), ZSTR_VAL(trait.GetName()))
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
		aliases = Ecalloc(i, b.SizeOf("zend_class_entry *"))
		i = 0
		for ce.GetTraitAliases()[i] != nil {

			/** For all aliases with an explicit class name, resolve the class now. */

			if ce.GetTraitAliases()[i].GetTraitMethod().GetClassName() != nil {
				cur_method_ref = &ce.trait_aliases[i].GetTraitMethod()
				trait = ZendFetchClass(cur_method_ref.GetClassName(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if trait == nil {
					ZendErrorNoreturn(E_COMPILE_ERROR, "Could not find trait %s", ZSTR_VAL(cur_method_ref.GetClassName()))
				}
				ZendCheckTraitUsage(ce, trait, traits)
				aliases[i] = trait

				/** And, ensure that the referenced method is resolvable, too. */

				lcname = ZendStringTolower(cur_method_ref.GetMethodName())
				if ZendHashExists(&trait.function_table, lcname) == 0 {
					ZendErrorNoreturn(E_COMPILE_ERROR, "An alias was defined for %s::%s but this method does not exist", ZSTR_VAL(trait.GetName()), ZSTR_VAL(cur_method_ref.GetMethodName()))
				}
				ZendStringReleaseEx(lcname, 0)
			}
			i++
		}
	}
	*exclude_tables_ptr = exclude_tables
	*aliases_ptr = aliases
}
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

						if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
							continue
						}
						key = _p.GetKey()
						fn = Z_PTR_P(_z)
						ZendTraitsCopyFunctions(key, fn, ce, &overridden, exclude_tables[i], aliases)
					}
					break
				}
				if exclude_tables[i] != nil {
					ZendHashDestroy(exclude_tables[i])
					FREE_HASHTABLE(exclude_tables[i])
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

						if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
							continue
						}
						key = _p.GetKey()
						fn = Z_PTR_P(_z)
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			fn = Z_PTR_P(_z)
			ZendFixupTraitMethod(fn, ce)
		}
		break
	}
	if overridden != nil {
		ZendHashDestroy(overridden)
		FREE_HASHTABLE(overridden)
	}
}
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

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				property_info = Z_PTR_P(_z)

				/* first get the unmangeld name if necessary,
				 * then check whether the property is already there
				 */

				flags = property_info.GetFlags()
				if (flags & ZEND_ACC_PUBLIC) != 0 {
					prop_name = ZendStringCopy(property_info.GetName())
				} else {
					var pname *byte
					var pname_len int

					/* for private and protected we need to unmangle the names */

					ZendUnmanglePropertyNameEx(property_info.GetName(), &class_name_unused, &pname, &pname_len)
					prop_name = ZendStringInit(pname, pname_len, 0)
				}

				/* next: check for conflicts with current class */

				if b.Assign(&coliding_prop, ZendHashFindPtr(&ce.properties_info, prop_name)) != nil {
					if (coliding_prop.GetFlags()&ZEND_ACC_PRIVATE) != 0 && coliding_prop.GetCe() != ce {
						ZendHashDel(&ce.properties_info, prop_name)
						flags |= ZEND_ACC_CHANGED
					} else {
						not_compatible = 1
						if (coliding_prop.GetFlags()&(ZEND_ACC_PPP_MASK|ZEND_ACC_STATIC)) == (flags&(ZEND_ACC_PPP_MASK|ZEND_ACC_STATIC)) && PropertyTypesCompatible(property_info, coliding_prop) == INHERITANCE_SUCCESS {

							/* the flags are identical, thus, the properties may be compatible */

							var op1 *Zval
							var op2 *Zval
							var op1_tmp Zval
							var op2_tmp Zval
							if (flags & ZEND_ACC_STATIC) != 0 {
								op1 = &ce.default_static_members_table[coliding_prop.GetOffset()]
								op2 = &traits[i].default_static_members_table[property_info.GetOffset()]
								ZVAL_DEINDIRECT(op1)
								ZVAL_DEINDIRECT(op2)
							} else {
								op1 = &ce.default_properties_table[OBJ_PROP_TO_NUM(coliding_prop.GetOffset())]
								op2 = &traits[i].default_properties_table[OBJ_PROP_TO_NUM(property_info.GetOffset())]
							}

							/* if any of the values is a constant, we try to resolve it */

							if UNEXPECTED(Z_TYPE_P(op1) == IS_CONSTANT_AST) {
								ZVAL_COPY_OR_DUP(&op1_tmp, op1)
								ZvalUpdateConstantEx(&op1_tmp, ce)
								op1 = &op1_tmp
							}
							if UNEXPECTED(Z_TYPE_P(op2) == IS_CONSTANT_AST) {
								ZVAL_COPY_OR_DUP(&op2_tmp, op2)
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
							ZendErrorNoreturn(E_COMPILE_ERROR, "%s and %s define the __special__  same property ($%s) in the composition of %s. However, the definition differs and is considered incompatible. Class was composed", ZSTR_VAL(FindFirstDefinition(ce, traits, i, prop_name, coliding_prop.GetCe()).GetName()), ZSTR_VAL(property_info.GetCe().GetName()), ZSTR_VAL(prop_name), ZSTR_VAL(ce.GetName()))
						}
						ZendStringReleaseEx(prop_name, 0)
						continue
					}
				}

				/* property not found, so lets add it */

				if (flags & ZEND_ACC_STATIC) != 0 {
					prop_value = &traits[i].default_static_members_table[property_info.GetOffset()]
					ZEND_ASSERT(Z_TYPE_P(prop_value) != IS_INDIRECT)
				} else {
					prop_value = &traits[i].default_properties_table[OBJ_PROP_TO_NUM(property_info.GetOffset())]
				}
				Z_TRY_ADDREF_P(prop_value)
				if property_info.GetDocComment() != nil {
					doc_comment = ZendStringCopy(property_info.GetDocComment())
				} else {
					doc_comment = nil
				}
				if ZEND_TYPE_IS_NAME(property_info.GetType()) {
					ZendStringAddref(ZEND_TYPE_NAME(property_info.GetType()))
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

					ZendErrorNoreturn(E_COMPILE_ERROR, "An alias (%s) was defined for method %s(), but this method does not exist", ZSTR_VAL(cur_alias.GetAlias()), ZSTR_VAL(cur_alias.GetTraitMethod().GetMethodName()))

					/** Plain old inconsistency/typo/bug */

				} else {

					/** Here are two possible cases:
					  1) this is an attempt to modify the visibility
					     of a method introduce as part of another alias.
					     Since that seems to violate the DRY principle,
					     we check against it and abort.
					  2) it is just a plain old inconsitency/typo/bug
					     as in the case where alias is set. */

					lc_method_name = ZendStringTolower(cur_alias.GetTraitMethod().GetMethodName())
					if ZendHashExists(&ce.function_table, lc_method_name) != 0 {
						ZendStringReleaseEx(lc_method_name, 0)
						ZendErrorNoreturn(E_COMPILE_ERROR, "The modifiers for the trait alias %s() need to be changed in the same statement in which the alias is defined. Error", ZSTR_VAL(cur_alias.GetTraitMethod().GetMethodName()))
					} else {
						ZendStringReleaseEx(lc_method_name, 0)
						ZendErrorNoreturn(E_COMPILE_ERROR, "The modifiers of the trait method %s() are changed, but this method does not exist. Error", ZSTR_VAL(cur_alias.GetTraitMethod().GetMethodName()))
					}
				}
			}
			i++
		}
	}
}
func ZendDoBindTraits(ce *ZendClassEntry) {
	var exclude_tables **HashTable
	var aliases **ZendClassEntry
	var traits **ZendClassEntry
	var trait **ZendClassEntry
	var i uint32
	var j uint32
	ZEND_ASSERT(ce.GetNumTraits() > 0)
	traits = Emalloc(b.SizeOf("zend_class_entry *") * ce.GetNumTraits())
	for i = 0; i < ce.GetNumTraits(); i++ {
		trait = ZendFetchClassByName(ce.GetTraitNames()[i].GetName(), ce.GetTraitNames()[i].GetLcName(), ZEND_FETCH_CLASS_TRAIT)
		if UNEXPECTED(trait == nil) {
			return
		}
		if UNEXPECTED((trait.ce_flags & ZEND_ACC_TRAIT) == 0) {
			ZendErrorNoreturn(E_ERROR, "%s cannot use %s - it is not a trait", ZSTR_VAL(ce.GetName()), ZSTR_VAL(trait.name))
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
		Efree(aliases)
	}
	if exclude_tables != nil {
		Efree(exclude_tables)
	}

	/* then flatten the properties into it, to, mostly to notfiy developer about problems */

	ZendDoTraitsPropertyBinding(ce, traits)
	Efree(traits)

	/* Emit E_DEPRECATED for PHP 4 constructors */

	ZendCheckDeprecatedConstructor(ce)

	/* Emit E_DEPRECATED for PHP 4 constructors */
}
func ZendHasDeprecatedConstructor(ce *ZendClassEntry) ZendBool {
	var constructor_name *ZendString
	if ce.GetConstructor() == nil {
		return 0
	}
	constructor_name = ce.GetConstructor().GetFunctionName()
	return !(ZendBinaryStrcasecmp(ZSTR_VAL(ce.GetName()), ZSTR_LEN(ce.GetName()), ZSTR_VAL(constructor_name), ZSTR_LEN(constructor_name)))
}
func ZendCheckDeprecatedConstructor(ce *ZendClassEntry) {
	if ZendHasDeprecatedConstructor(ce) != 0 {
		ZendError(E_DEPRECATED, "Methods with the same name as their class will not be constructors in a future version of PHP; %s has a deprecated constructor", ZSTR_VAL(ce.GetName()))
	}
}
func DISPLAY_ABSTRACT_FN(idx int) {
	b.CondF1(ai.afn[idx], func() string { return ZEND_FN_SCOPE_NAME(ai.afn[idx]) }, "")
	b.Cond(ai.afn[idx], "::", "")
	b.CondF1(ai.afn[idx], func() []byte { return ZSTR_VAL(ai.afn[idx].common.function_name) }, "")
	b.CondF2(ai.afn[idx] && ai.afn[idx+1], ", ", func() string {
		if ai.afn[idx] && ai.cnt > MAX_ABSTRACT_INFO_CNT {
			return ", ..."
		} else {
			return ""
		}
	})
}
func ZendVerifyAbstractClassFunction(fn *ZendFunction, ai *ZendAbstractInfo) {
	if (fn.GetFnFlags() & ZEND_ACC_ABSTRACT) != 0 {
		if ai.GetCnt() < MAX_ABSTRACT_INFO_CNT {
			ai.GetAfn()[ai.GetCnt()] = fn
		}
		if (fn.GetFnFlags() & ZEND_ACC_CTOR) != 0 {
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
func ZendVerifyAbstractClass(ce *ZendClassEntry) {
	var func_ *ZendFunction
	var ai ZendAbstractInfo
	ZEND_ASSERT((ce.GetCeFlags() & (ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) == ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
	memset(&ai, 0, b.SizeOf("ai"))
	for {
		var __ht *HashTable = &ce.function_table
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			func_ = Z_PTR_P(_z)
			ZendVerifyAbstractClassFunction(func_, &ai)
		}
		break
	}
	if ai.GetCnt() != 0 {
		ZendErrorNoreturn(E_ERROR, "Class %s contains %d abstract method%s and must therefore be declared abstract or implement the remaining methods ("+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+")", ZSTR_VAL(ce.GetName()), ai.GetCnt(), b.Cond(ai.GetCnt() > 1, "s", ""), DISPLAY_ABSTRACT_FN(0), DISPLAY_ABSTRACT_FN(1), DISPLAY_ABSTRACT_FN(2))
	} else {

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

		ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

	}
}
func VarianceObligationDtor(zv *Zval) { Efree(Z_PTR_P(zv)) }
func VarianceObligationHtDtor(zv *Zval) {
	ZendHashDestroy(Z_PTR_P(zv))
	FREE_HASHTABLE(Z_PTR_P(zv))
}
func GetOrInitObligationsForClass(ce *ZendClassEntry) *HashTable {
	var ht *HashTable
	var key ZendUlong
	if CompilerGlobals.GetDelayedVarianceObligations() == nil {
		ALLOC_HASHTABLE(CompilerGlobals.GetDelayedVarianceObligations())
		ZendHashInit(CompilerGlobals.GetDelayedVarianceObligations(), 0, nil, VarianceObligationHtDtor, 0)
	}
	key = ZendUlong(uintPtr(ce))
	ht = ZendHashIndexFindPtr(CompilerGlobals.GetDelayedVarianceObligations(), key)
	if ht != nil {
		return ht
	}
	ALLOC_HASHTABLE(ht)
	ZendHashInit(ht, 0, nil, VarianceObligationDtor, 0)
	ZendHashIndexAddNewPtr(CompilerGlobals.GetDelayedVarianceObligations(), key, ht)
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_UNRESOLVED_VARIANCE)
	return ht
}
func AddDependencyObligation(ce *ZendClassEntry, dependency_ce *ZendClassEntry) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_DEPENDENCY)
	obligation.dependency_ce = dependency_ce
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddCompatibilityObligation(ce *ZendClassEntry, child_fn *ZendFunction, parent_fn *ZendFunction, always_error ZendBool) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_COMPATIBILITY)

	/* Copy functions, because they may be stack-allocated in the case of traits. */

	if child_fn.GetCommonType() == ZEND_INTERNAL_FUNCTION {
		memcpy(&obligation.child_fn, child_fn, b.SizeOf("zend_internal_function"))
	} else {
		memcpy(&obligation.child_fn, child_fn, b.SizeOf("zend_op_array"))
	}
	if parent_fn.GetCommonType() == ZEND_INTERNAL_FUNCTION {
		memcpy(&obligation.parent_fn, parent_fn, b.SizeOf("zend_internal_function"))
	} else {
		memcpy(&obligation.parent_fn, parent_fn, b.SizeOf("zend_op_array"))
	}
	obligation.always_error = always_error
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddPropertyCompatibilityObligation(ce *ZendClassEntry, child_prop *ZendPropertyInfo, parent_prop *ZendPropertyInfo) {
	var obligations *HashTable = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_PROPERTY_COMPATIBILITY)
	obligation.child_prop = child_prop
	obligation.parent_prop = parent_prop
	ZendHashNextIndexInsertPtr(obligations, obligation)
}
func CheckVarianceObligation(zv *Zval) int {
	var obligation *VarianceObligation = Z_PTR_P(zv)
	if obligation.GetType() == OBLIGATION_DEPENDENCY {
		var dependency_ce *ZendClassEntry = obligation.dependency_ce
		if (dependency_ce.GetCeFlags() & ZEND_ACC_UNRESOLVED_VARIANCE) != 0 {
			ResolveDelayedVarianceObligations(dependency_ce)
		}
		if (dependency_ce.GetCeFlags() & ZEND_ACC_LINKED) == 0 {
			return ZEND_HASH_APPLY_KEEP
		}
	} else if obligation.GetType() == OBLIGATION_COMPATIBILITY {
		var unresolved_class *ZendString
		var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, &obligation.child_fn, &obligation.parent_fn)
		if UNEXPECTED(status != INHERITANCE_SUCCESS) {
			if EXPECTED(status == INHERITANCE_UNRESOLVED) {
				return ZEND_HASH_APPLY_KEEP
			}
			ZEND_ASSERT(status == INHERITANCE_ERROR)
			EmitIncompatibleMethodErrorOrWarning(&obligation.child_fn, &obligation.parent_fn, status, unresolved_class, obligation.always_error)
		}
	} else {
		ZEND_ASSERT(obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY)
		var status InheritanceStatus = PropertyTypesCompatible(obligation.parent_prop, obligation.child_prop)
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return ZEND_HASH_APPLY_KEEP
			}
			ZEND_ASSERT(status == INHERITANCE_ERROR)
			EmitIncompatiblePropertyError(obligation.child_prop, obligation.parent_prop)
		}
	}
	return ZEND_HASH_APPLY_REMOVE
}
func LoadDelayedClasses() {
	var delayed_autoloads *HashTable = CompilerGlobals.GetDelayedAutoloads()
	var name *ZendString
	if delayed_autoloads == nil {
		return
	}

	/* Take ownership of this HT, to avoid concurrent modification during autoloading. */

	CompilerGlobals.SetDelayedAutoloads(nil)
	for {
		var __ht *HashTable = delayed_autoloads
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			name = _p.GetKey()
			ZendLookupClass(name)
		}
		break
	}
	ZendHashDestroy(delayed_autoloads)
	FREE_HASHTABLE(delayed_autoloads)
}
func ResolveDelayedVarianceObligations(ce *ZendClassEntry) {
	var all_obligations *HashTable = CompilerGlobals.GetDelayedVarianceObligations()
	var obligations *HashTable
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	ZEND_ASSERT(all_obligations != nil)
	obligations = ZendHashIndexFindPtr(all_obligations, num_key)
	ZEND_ASSERT(obligations != nil)
	ZendHashApply(obligations, CheckVarianceObligation)
	if ZendHashNumElements(obligations) == 0 {
		ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_UNRESOLVED_VARIANCE)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_LINKED)
		ZendHashIndexDel(all_obligations, num_key)
	}
}
func ReportVarianceErrors(ce *ZendClassEntry) {
	var all_obligations *HashTable = CompilerGlobals.GetDelayedVarianceObligations()
	var obligations *HashTable
	var obligation *VarianceObligation
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	ZEND_ASSERT(all_obligations != nil)
	obligations = ZendHashIndexFindPtr(all_obligations, num_key)
	ZEND_ASSERT(obligations != nil)
	for {
		var __ht *HashTable = obligations
		var _p *Bucket = __ht.GetArData()
		var _end *Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *Zval = &_p.val

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			obligation = Z_PTR_P(_z)
			var status InheritanceStatus
			var unresolved_class *ZendString
			if obligation.GetType() == OBLIGATION_COMPATIBILITY {

				/* Just used to fetch the unresolved_class in this case. */

				status = ZendDoPerformImplementationCheck(&unresolved_class, &obligation.child_fn, &obligation.parent_fn)
				ZEND_ASSERT(status == INHERITANCE_UNRESOLVED)
				EmitIncompatibleMethodErrorOrWarning(&obligation.child_fn, &obligation.parent_fn, status, unresolved_class, obligation.always_error)
			} else if obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY {
				EmitIncompatiblePropertyError(obligation.child_prop, obligation.parent_prop)
			} else {
				ZendErrorNoreturn(E_CORE_ERROR, "Bug #78647")
			}
		}
		break
	}

	/* Only warnings were thrown above -- that means that there are incompatibilities, but only
	 * ones that we permit. Mark all classes with open obligations as fully linked. */

	ce.SetCeFlags(ce.GetCeFlags() &^ ZEND_ACC_UNRESOLVED_VARIANCE)
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_LINKED)
	ZendHashIndexDel(all_obligations, num_key)
}
func CheckUnrecoverableLoadFailure(ce *ZendClassEntry) {
	/* If this class has been used while unlinked through a variance obligation, it is not legal
	 * to remove the class from the class table and throw an exception, because there is already
	 * a dependence on the inheritance hierarchy of this specific class. Instead we fall back to
	 * a fatal error, as would happen if we did not allow exceptions in the first place. */

	if (ce.GetCeFlags() & ZEND_ACC_HAS_UNLINKED_USES) != 0 {
		var exception_str *ZendString
		var exception_zv Zval
		ZEND_ASSERT(ExecutorGlobals.GetException() != nil && "Exception must have been thrown")
		ZVAL_OBJ(&exception_zv, ExecutorGlobals.GetException())
		Z_ADDREF(exception_zv)
		ZendClearException()
		exception_str = ZvalGetString(&exception_zv)
		ZendErrorNoreturn(E_ERROR, "During inheritance of %s with variance dependencies: Uncaught %s", ZSTR_VAL(ce.GetName()), ZSTR_VAL(exception_str))
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
		parent = ZendFetchClassByName(ce.parent_name, lc_parent_name, ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED|ZEND_FETCH_CLASS_EXCEPTION)
		if parent == nil {
			CheckUnrecoverableLoadFailure(ce)
			return FAILURE
		}
	}
	if ce.GetNumInterfaces() != 0 {

		/* Also copy the parent interfaces here, so we don't need to reallocate later. */

		var i uint32
		var num_parent_interfaces uint32 = b.CondF1(parent != nil, func() uint32 { return parent.GetNumInterfaces() }, 0)
		interfaces = Emalloc(b.SizeOf("zend_class_entry *") * (ce.GetNumInterfaces() + num_parent_interfaces))
		if num_parent_interfaces != 0 {
			memcpy(interfaces, parent.interfaces, b.SizeOf("zend_class_entry *")*num_parent_interfaces)
		}
		for i = 0; i < ce.GetNumInterfaces(); i++ {
			var iface *ZendClassEntry = ZendFetchClassByName(ce.interface_names[i].name, ce.interface_names[i].lc_name, ZEND_FETCH_CLASS_INTERFACE|ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED|ZEND_FETCH_CLASS_EXCEPTION)
			if iface == nil {
				CheckUnrecoverableLoadFailure(ce)
				Efree(interfaces)
				return FAILURE
			}
			interfaces[num_parent_interfaces+i] = iface
		}
	}
	if parent != nil {
		if (parent.GetCeFlags() & ZEND_ACC_LINKED) == 0 {
			AddDependencyObligation(ce, parent)
		}
		ZendDoInheritance(ce, parent)
	}
	if (ce.GetCeFlags() & ZEND_ACC_IMPLEMENT_TRAITS) != 0 {
		ZendDoBindTraits(ce)
	}
	if (ce.GetCeFlags() & ZEND_ACC_IMPLEMENT_INTERFACES) != 0 {
		ZendDoImplementInterfaces(ce, interfaces)
	}
	if (ce.GetCeFlags() & (ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) == ZEND_ACC_IMPLICIT_ABSTRACT_CLASS {
		ZendVerifyAbstractClass(ce)
	}
	ZendBuildPropertiesInfoTable(ce)
	if (ce.GetCeFlags() & ZEND_ACC_UNRESOLVED_VARIANCE) == 0 {
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_LINKED)
		return SUCCESS
	}
	ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_NEARLY_LINKED)
	LoadDelayedClasses()
	if (ce.GetCeFlags() & ZEND_ACC_UNRESOLVED_VARIANCE) != 0 {
		ResolveDelayedVarianceObligations(ce)
		if (ce.GetCeFlags() & ZEND_ACC_LINKED) == 0 {
			ReportVarianceErrors(ce)
		}
	}
	return SUCCESS
}
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			parent_func = Z_PTR_P(_z)
			var zv *Zval = ZendHashFindEx(&ce.function_table, key, 1)
			if zv != nil {
				var child_func *ZendFunction = Z_FUNC_P(zv)
				var status InheritanceStatus = DoInheritanceCheckOnMethodEx(child_func, parent_func, ce, nil, 1, 0)
				if UNEXPECTED(status != INHERITANCE_SUCCESS) {
					if EXPECTED(status == INHERITANCE_UNRESOLVED) {
						return INHERITANCE_UNRESOLVED
					}
					ZEND_ASSERT(status == INHERITANCE_ERROR)
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

			if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			parent_info = Z_PTR_P(_z)
			var zv *Zval
			if (parent_info.GetFlags()&ZEND_ACC_PRIVATE) != 0 || !(ZEND_TYPE_IS_SET(parent_info.GetType())) {
				continue
			}
			zv = ZendHashFindEx(&ce.properties_info, key, 1)
			if zv != nil {
				var child_info *ZendPropertyInfo = Z_PTR_P(zv)
				if ZEND_TYPE_IS_SET(child_info.GetType()) {
					var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
					if UNEXPECTED(status != INHERITANCE_SUCCESS) {
						if EXPECTED(status == INHERITANCE_UNRESOLVED) {
							return INHERITANCE_UNRESOLVED
						}
						ZEND_ASSERT(status == INHERITANCE_ERROR)
						ret = INHERITANCE_ERROR
					}
				}
			}
		}
		break
	}
	return ret
}
func ZendTryEarlyBind(ce *ZendClassEntry, parent_ce *ZendClassEntry, lcname *ZendString, delayed_early_binding *Zval) ZendBool {
	var status InheritanceStatus = ZendCanEarlyBind(ce, parent_ce)
	if EXPECTED(status != INHERITANCE_UNRESOLVED) {
		if delayed_early_binding != nil {
			if UNEXPECTED(ZendHashSetBucketKey(ExecutorGlobals.GetClassTable(), (*Bucket)(delayed_early_binding), lcname) == nil) {
				ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ZSTR_VAL(ce.GetName()))
				return 0
			}
		} else {
			if UNEXPECTED(ZendHashAddPtr(CompilerGlobals.GetClassTable(), lcname, ce) == nil) {
				return 0
			}
		}
		ZendDoInheritanceEx(ce, parent_ce, status == INHERITANCE_SUCCESS)
		ZendBuildPropertiesInfoTable(ce)
		if (ce.GetCeFlags() & (ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | ZEND_ACC_INTERFACE | ZEND_ACC_TRAIT | ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) == ZEND_ACC_IMPLICIT_ABSTRACT_CLASS {
			ZendVerifyAbstractClass(ce)
		}
		ZEND_ASSERT((ce.GetCeFlags() & ZEND_ACC_UNRESOLVED_VARIANCE) == 0)
		ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_LINKED)
		return 1
	}
	return 0
}
