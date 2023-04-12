package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendDoInheritance(ce *types.ClassEntry, parent_ce *types.ClassEntry) {
	ZendDoInheritanceEx(ce, parent_ce, false)
}
func OverriddenPtrDtor(zv *types.Zval) {
	EfreeSize(zv.GetPtr(), b.SizeOf("zend_function"))
}
func ZendDuplicatePropertyInfoInternal(property_info *ZendPropertyInfo) *ZendPropertyInfo {
	var new_property_info *ZendPropertyInfo = Pemalloc(b.SizeOf("zend_property_info"), 1)
	memcpy(new_property_info, property_info, b.SizeOf("zend_property_info"))
	//new_property_info.GetName().AddRefcount()
	if new_property_info.GetType().IsName() {
		//new_property_info.GetType().Name().AddRefcount()
	}
	return new_property_info
}
func ZendDuplicateInternalFunction(func_ types.IFunction, ce *types.ClassEntry) types.IFunction {
	var new_function types.IFunction
	if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
		new_function = Pemalloc(b.SizeOf("zend_internal_function"), 1)
		memcpy(new_function, func_, b.SizeOf("zend_internal_function"))
	} else {
		new_function = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_internal_function"))
		memcpy(new_function, func_, b.SizeOf("zend_internal_function"))
		new_function.SetIsArenaAllocated(true)
	}
	if new_function.GetFunctionName() != nil {
		//new_function.GetFunctionName().AddRefcount()
	}
	return new_function
}
func ZendDuplicateUserFunction(func_ types.IFunction) types.IFunction {
	var new_function types.IFunction
	new_function = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_op_array"))
	memcpy(new_function, func_, b.SizeOf("zend_op_array"))
	if ZEND_MAP_PTR_GET(func_.GetOpArray().static_variables_ptr) {

		/* See: Zend/tests/method_static_var.phpt */

		new_function.GetOpArray().SetStaticVariables(ZEND_MAP_PTR_GET(func_.GetOpArray().static_variables_ptr))

		/* See: Zend/tests/method_static_var.phpt */

	}
	if (new_function.GetOpArray().GetStaticVariables().GetGcFlags() & types.IS_ARRAY_IMMUTABLE) == 0 {
		new_function.GetOpArray().GetStaticVariables().AddRefcount()
	}
	if (CG__().GetCompilerOptions() & ZEND_COMPILE_PRELOAD) != 0 {
		b.Assert(new_function.GetOpArray().IsPreloaded())
		ZEND_MAP_PTR_NEW(new_function.GetOpArray().static_variables_ptr)
	} else {
		ZEND_MAP_PTR_INIT(new_function.GetOpArray().static_variables_ptr, new_function.GetOpArray().GetStaticVariables())
	}
	return new_function
}
func ZendDuplicateFunction(func_ types.IFunction, ce *types.ClassEntry, is_interface types.ZendBool) types.IFunction {
	if func_.GetType() == ZEND_INTERNAL_FUNCTION {
		return ZendDuplicateInternalFunction(func_, ce)
	} else {
		if func_.GetOpArray().GetRefcount() != nil {
			func_.op_array.refcount++
		}
		if is_interface != 0 || func_.GetOpArray().GetStaticVariables() == nil {

			/* reuse the same op_array structure */

			return func_

			/* reuse the same op_array structure */

		}
		return ZendDuplicateUserFunction(func_)
	}
}
func DoInheritParentConstructor(ce *types.ClassEntry) {
	var parent *types.ClassEntry = ce.GetParent()
	b.Assert(parent != nil)

	/* You cannot change create_object */

	ce.SetCreateObject(parent.GetCreateObject())

	/* Inherit special functions if needed */

	if ce.GetGetIterator() == nil {
		ce.SetGetIterator(parent.GetGetIterator())
	}
	if parent.GetIteratorFuncsPtr() != nil {

		/* Must be initialized through iface->interface_gets_implemented() */

		b.Assert(ce.GetIteratorFuncsPtr() != nil)

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
		if parent.GetConstructor() != nil && parent.GetConstructor().IsFinal() {
			faults.ErrorNoreturn(faults.E_ERROR, "Cannot override final %s::%s() with %s::%s()", parent.GetName().GetVal(), parent.GetConstructor().GetFunctionName().GetVal(), ce.GetName().GetVal(), ce.GetConstructor().GetFunctionName().GetVal())
		}
		return
	}
	ce.SetConstructor(parent.GetConstructor())
}
func ZendVisibilityString(fn_flags uint32) *byte {
	if (fn_flags & AccPublic) != 0 {
		return "public"
	} else if (fn_flags & AccPrivate) != 0 {
		return "private"
	} else {
		b.Assert((fn_flags & AccProtected) != 0)
		return "protected"
	}
}
func ResolveClassName(scope *types.ClassEntry, name *types.String) *types.String {
	b.Assert(scope != nil)
	if ascii.StrCaseEquals(name.GetStr(), "parent") && scope.GetParent() {
		if scope.IsResolvedParent() {
			return scope.GetParent().name
		} else {
			return scope.GetParentName()
		}
	} else if ascii.StrCaseEquals(name.GetStr(), "self") {
		return scope.GetName()
	} else {
		return name
	}
}
func ClassVisible(ce *types.ClassEntry) types.ZendBool {
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		return !(CG__().GetCompilerOptions() & ZEND_COMPILE_IGNORE_INTERNAL_CLASSES)
	} else {
		b.Assert(ce.GetType() == ZEND_USER_CLASS)
		return (CG__().GetCompilerOptions()&ZEND_COMPILE_IGNORE_OTHER_FILES) == 0 || ce.GetFilename() == CG__().GetCompiledFilename()
	}
}
func LookupClass(scope *types.ClassEntry, name *types.String) *types.ClassEntry {
	var ce *types.ClassEntry
	if CG__().GetInCompilation() == 0 {
		var flags uint32 = ZEND_FETCH_CLASS_ALLOW_UNLINKED | ZEND_FETCH_CLASS_NO_AUTOLOAD
		ce = ZendLookupClassEx(name, nil, flags)
		if ce != nil {
			return ce
		}

		/* We'll autoload this class and process delayed variance obligations later. */

		if CG__().GetDelayedAutoloads() == nil {
			ALLOC_HASHTABLE(CG__().GetDelayedAutoloads())
			CG__().GetDelayedAutoloads() = types.MakeArrayEx(0, nil, 0)
		}
		types.ZendHashAddEmptyElement(CG__().GetDelayedAutoloads(), name.GetStr())
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce != nil && ClassVisible(ce) != 0 {
			return ce
		}

		/* The current class may not be registered yet, so check for it explicitly. */

		if ascii.StrCaseEquals(scope.GetName().GetStr(), name.GetStr()) {
			return scope
		}

		/* The current class may not be registered yet, so check for it explicitly. */

	}
	return nil
}
func UnlinkedInstanceof(ce1 *types.ClassEntry, ce2 *types.ClassEntry) types.ZendBool {
	if ce1 == ce2 {
		return 1
	}
	if ce1.IsLinked() {
		return InstanceofFunction(ce1, ce2)
	}
	if ce1.GetParent() {
		var parent_ce *types.ClassEntry
		if ce1.IsResolvedParent() {
			parent_ce = ce1.GetParent()
		} else {
			parent_ce = ZendLookupClassEx(ce1.GetParentName(), nil, ZEND_FETCH_CLASS_ALLOW_UNLINKED|ZEND_FETCH_CLASS_NO_AUTOLOAD)
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
		if ce1.IsResolvedInterfaces() {

			/* Unlike the normal instanceof_function(), we have to perform a recursive
			 * check here, as the parent interfaces might not have been fully copied yet. */

			for i = 0; i < ce1.GetNumInterfaces(); i++ {
				if UnlinkedInstanceof(ce1.GetInterfaces()[i], ce2) != 0 {
					return 1
				}
			}

			/* Unlike the normal instanceof_function(), we have to perform a recursive
			 * check here, as the parent interfaces might not have been fully copied yet. */

		} else {
			for i = 0; i < ce1.GetNumInterfaces(); i++ {
				var ce *types.ClassEntry = ZendLookupClassEx(ce1.GetInterfaceNames()[i].name, ce1.GetInterfaceNames()[i].lcName, ZEND_FETCH_CLASS_ALLOW_UNLINKED|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil && UnlinkedInstanceof(ce, ce2) != 0 {
					return 1
				}
			}
		}
	}
	return 0
}
func ZendPerformCovariantTypeCheck(unresolved_class **types.String, fe types.IFunction, fe_arg_info *ZendArgInfo, proto types.IFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	var fe_type types.ZendType = fe_arg_info.GetType()
	var proto_type types.ZendType = proto_arg_info.GetType()
	b.Assert(fe_type.IsSet() && proto_type.IsSet())
	if fe_type.AllowNull() && !(proto_type.AllowNull()) {
		return INHERITANCE_ERROR
	}
	if proto_type.IsClass() {
		var fe_class_name *types.String
		var proto_class_name *types.String
		var fe_ce *types.ClassEntry
		var proto_ce *types.ClassEntry
		if !(fe_type.IsClass()) {
			return INHERITANCE_ERROR
		}
		fe_class_name = ResolveClassName(fe.GetScope(), fe_type.Name())
		proto_class_name = ResolveClassName(proto.GetScope(), proto_type.Name())
		if ascii.StrCaseEquals(fe_class_name.GetStr(), proto_class_name.GetStr()) {
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
	} else if proto_type.Code() == types.IS_ITERABLE {
		if fe_type.IsClass() {
			var fe_class_name *types.String = ResolveClassName(fe.GetScope(), fe_type.Name())
			var fe_ce *types.ClassEntry = LookupClass(fe.GetScope(), fe_class_name)
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
		if fe_type.Code() == types.IS_ITERABLE || fe_type.Code() == types.IS_ARRAY {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else if proto_type.Code() == types.IS_OBJECT {
		if fe_type.IsClass() {

			/* Currently, any class name would be allowed here. We still perform a class lookup
			 * for forward-compatibility reasons, as we may have named types in the future that
			 * are not classes (such as enums or typedefs). */

			var fe_class_name *types.String = ResolveClassName(fe.GetScope(), fe_type.Name())
			var fe_ce *types.ClassEntry = LookupClass(fe.GetScope(), fe_class_name)
			if fe_ce == nil {
				*unresolved_class = fe_class_name
				return INHERITANCE_UNRESOLVED
			}
			return INHERITANCE_SUCCESS
		}
		if fe_type.Code() == types.IS_OBJECT {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else {
		if fe_type.Code() == proto_type.Code() {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	}
}
func ZendDoPerformArgTypeHintCheck(unresolved_class **types.String, fe types.IFunction, fe_arg_info *ZendArgInfo, proto types.IFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	if !(fe_arg_info.GetType().IsSet()) {

		/* Child with no type is always compatible */

		return INHERITANCE_SUCCESS

		/* Child with no type is always compatible */

	}
	if !(proto_arg_info.GetType().IsSet()) {

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
func ZendDoPerformImplementationCheck(unresolved_class **types.String, fe types.IFunction, proto types.IFunction) InheritanceStatus {
	var i uint32
	var num_args uint32
	var status InheritanceStatus
	var local_status InheritanceStatus

	/* If it's a user function then arg_info == NULL means we don't have any parameters but
	 * we still need to do the arg number checks.  We are only willing to ignore this for internal
	 * functions because extensions don't always define arg_info.
	 */

	if proto.GetArgInfo() == nil && proto.GetType() != ZEND_USER_FUNCTION {
		return INHERITANCE_SUCCESS
	}

	/* Checks for constructors only if they are declared in an interface,
	 * or explicitly marked as abstract
	 */

	b.Assert(!(fe.IsCtor() && (!proto.GetScope().IsInterface() && !proto.IsAbstract())))

	/* If the prototype method is private do not enforce a signature */

	b.Assert(!proto.IsPrivate())

	/* check number of arguments */

	if proto.GetRequiredNumArgs() < fe.GetRequiredNumArgs() || proto.GetNumArgs() > fe.GetNumArgs() {
		return INHERITANCE_ERROR
	}

	/* by-ref constraints on return values are covariant */

	if proto.IsReturnReference() && !fe.IsReturnReference() {
		return INHERITANCE_ERROR
	}
	if proto.IsVariadic() && !fe.IsVariadic() {
		return INHERITANCE_ERROR
	}

	/* For variadic functions any additional (optional) arguments that were added must be
	 * checked against the signature of the variadic argument, so in this case we have to
	 * go through all the parameters of the function and not just those present in the
	 * prototype. */

	num_args = proto.GetNumArgs()
	if proto.IsVariadic() {
		num_args++
		if fe.GetNumArgs() >= proto.GetNumArgs() {
			num_args = fe.GetNumArgs()
			if fe.IsVariadic() {
				num_args++
			}
		}
	}
	status = INHERITANCE_SUCCESS
	for i = 0; i < num_args; i++ {
		var fe_arg_info *ZendArgInfo = fe.GetArgInfo()[i]
		var proto_arg_info *ZendArgInfo
		if i < proto.GetNumArgs() {
			proto_arg_info = proto.GetArgInfo()[i]
		} else {
			proto_arg_info = proto.GetArgInfo()[proto.GetNumArgs()]
		}
		local_status = ZendDoPerformArgTypeHintCheck(unresolved_class, fe, fe_arg_info, proto, proto_arg_info)
		if local_status != INHERITANCE_SUCCESS {
			if local_status == INHERITANCE_ERROR {
				return INHERITANCE_ERROR
			}
			b.Assert(local_status == INHERITANCE_UNRESOLVED)
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

	if proto.IsHasReturnType() {

		/* Removing a return type is not valid. */

		if !fe.IsHasReturnType() {
			return INHERITANCE_ERROR
		}
		local_status = ZendPerformCovariantTypeCheck(unresolved_class, fe, fe.GetArgInfo()-1, proto, proto.GetArgInfo()-1)
		if local_status != INHERITANCE_SUCCESS {
			if local_status == INHERITANCE_ERROR {
				return INHERITANCE_ERROR
			}
			b.Assert(local_status == INHERITANCE_UNRESOLVED)
			status = INHERITANCE_UNRESOLVED
		}
	}
	return status
}
func ZendAppendTypeHint(str *SmartStr, fptr types.IFunction, arg_info *ZendArgInfo, return_hint int) {
	if arg_info.GetType().IsSet() && arg_info.GetType().AllowNull() {
		str.AppendByte('?')
	}
	if arg_info.GetType().IsClass() {
		var class_name *byte
		var class_name_len int
		class_name = types.ZEND_TYPE_NAME(arg_info.GetType()).GetVal()
		class_name_len = types.ZEND_TYPE_NAME(arg_info.GetType()).GetLen()
		if !(strcasecmp(class_name, "self")) && fptr.GetScope() != nil {
			class_name = fptr.GetScope().GetName().GetVal()
			class_name_len = fptr.GetScope().GetName().GetLen()
		} else if !(strcasecmp(class_name, "parent")) && fptr.GetScope() != nil && fptr.GetScope().GetParent() {
			class_name = fptr.GetScope().GetParent().name.GetVal()
			class_name_len = fptr.GetScope().GetParent().name.GetLen()
		}
		str.AppendString(b.CastStr(class_name, class_name_len))
		if return_hint == 0 {
			str.AppendByte(' ')
		}
	} else if arg_info.GetType().IsCode() {
		var type_name *byte = types.ZendGetTypeByConst(arg_info.GetType().Code())
		str.AppendString(b.CastStrAuto(type_name))
		if return_hint == 0 {
			str.AppendByte(' ')
		}
	}
}
func ZendGetFunctionDeclaration(fptr types.IFunction) *types.String {
	var str SmartStr = MakeSmartStr(0)
	if fptr.GetOpArray().IsReturnReference() {
		str.AppendString("& ")
	}
	if fptr.GetScope() != nil {

		/* cut off on NULL byte ... class@anonymous */

		str.AppendString(b.CastStr(fptr.GetScope().GetName().GetVal(), strlen(fptr.GetScope().GetName().GetVal())))
		str.AppendString("::")
	}
	str.AppendString(fptr.GetFunctionName().GetStr())
	str.AppendByte('(')
	if fptr.GetArgInfo() != nil {
		var i uint32
		var num_args uint32
		var required uint32
		var arg_info *ZendArgInfo = fptr.GetArgInfo()
		required = fptr.GetRequiredNumArgs()
		num_args = fptr.GetNumArgs()
		if fptr.IsVariadic() {
			num_args++
		}
		for i = 0; i < num_args; {
			ZendAppendTypeHint(&str, fptr, arg_info, 0)
			if arg_info.GetPassByReference() != 0 {
				str.AppendByte('&')
			}
			if arg_info.GetIsVariadic() != 0 {
				str.AppendString("...")
			}
			str.AppendByte('$')
			if arg_info.GetName() != nil {
				if fptr.GetType() == ZEND_INTERNAL_FUNCTION {
					str.AppendString((*ArgInfo)(arg_info).Name())
				} else {
					str.AppendString(arg_info.GetName().GetStr())
				}
			} else {
				str.AppendString("param")
				str.AppendUlong(i)
			}
			if i >= required && arg_info.GetIsVariadic() == 0 {
				str.AppendString(" = ")
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
						var zv *types.Zval = RT_CONSTANT(precv, precv.GetOp2())
						if zv.IsFalse() {
							str.AppendString("false")
						} else if zv.IsTrue() {
							str.AppendString("true")
						} else if zv.IsNull() {
							str.AppendString("NULL")
						} else if zv.IsString() {
							str.AppendByte('\'')
							str.AppendString(b.CastStr(zv.GetStr().GetVal(), b.Min(zv.GetStr().GetLen(), 10)))
							if zv.GetStr().GetLen() > 10 {
								str.AppendString("...")
							}
							str.AppendByte('\'')
						} else if zv.IsArray() {
							str.AppendString("Array")
						} else if zv.IsConstant() {
							var ast *ZendAst = types.Z_ASTVAL_P(zv)
							if ast.GetKind() == ZEND_AST_CONSTANT {
								str.AppendString(ZendAstGetConstantName(ast).GetStr())
							} else {
								str.AppendString("<expression>")
							}
						} else {
							var tmp_zv_str *types.String
							var zv_str *types.String = ZvalGetTmpString(zv, &tmp_zv_str)
							str.AppendString(zv_str.GetStr())
							ZendTmpStringRelease(tmp_zv_str)
						}
					}
				} else {
					str.AppendString("NULL")
				}
			}
			if b.PreInc(&i) < num_args {
				str.AppendString(", ")
			}
			arg_info++
		}
	}
	str.AppendByte(')')
	if fptr.IsHasReturnType() {
		str.AppendString(": ")
		ZendAppendTypeHint(&str, fptr, fptr.GetArgInfo()-1, 1)
	}
	str.ZeroTail()
	return str.GetS()
}
func FuncLineno(fn types.IFunction) uint32 {
	if fn.GetType() == ZEND_USER_FUNCTION {
		return fn.GetOpArray().GetLineStart()
	} else {
		return 0
	}
}
func EmitIncompatibleMethodError(
	error_level int,
	error_verb *byte,
	child types.IFunction,
	parent types.IFunction,
	status InheritanceStatus,
	unresolved_class *types.String,
) {
	var parent_prototype *types.String = ZendGetFunctionDeclaration(parent)
	var child_prototype *types.String = ZendGetFunctionDeclaration(child)
	if status == INHERITANCE_UNRESOLVED {
		faults.ErrorAt(error_level, nil, FuncLineno(child), "Could not check compatibility between %s and %s, because class %s is not available", child_prototype.GetVal(), parent_prototype.GetVal(), unresolved_class.GetVal())
	} else {
		faults.ErrorAt(error_level, nil, FuncLineno(child), "Declaration of %s %s be compatible with %s", child_prototype.GetVal(), error_verb, parent_prototype.GetVal())
	}
	// types.ZendStringEfree(child_prototype)
	// types.ZendStringEfree(parent_prototype)
}
func EmitIncompatibleMethodErrorOrWarning(child types.IFunction, parent types.IFunction, status InheritanceStatus, unresolved_class *types.String, always_error types.ZendBool) {
	var error_level int
	var error_verb *byte
	if always_error != 0 || child.GetPrototype() != nil && child.GetPrototype().IsAbstract() || parent.IsHasReturnType() && (!child.IsHasReturnType() || ZendPerformCovariantTypeCheck(&unresolved_class, child, child.GetArgInfo()-1, parent, parent.GetArgInfo()-1) != INHERITANCE_SUCCESS) {
		error_level = faults.E_COMPILE_ERROR
		error_verb = "must"
	} else {
		error_level = faults.E_WARNING
		error_verb = "should"
	}
	EmitIncompatibleMethodError(error_level, error_verb, child, parent, status, unresolved_class)
}
func PerformDelayableImplementationCheck(ce *types.ClassEntry, fe types.IFunction, proto types.IFunction, always_error types.ZendBool) {
	var unresolved_class *types.String
	var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, fe, proto)
	if status != INHERITANCE_SUCCESS {
		if status == INHERITANCE_UNRESOLVED {
			AddCompatibilityObligation(ce, fe, proto, always_error)
		} else {
			b.Assert(status == INHERITANCE_ERROR)
			if always_error != 0 {
				EmitIncompatibleMethodError(faults.E_COMPILE_ERROR, "must", fe, proto, status, unresolved_class)
			} else {
				EmitIncompatibleMethodErrorOrWarning(fe, proto, status, unresolved_class, always_error)
			}
		}
	}
}
func DoInheritanceCheckOnMethodEx(
	child types.IFunction,
	parent types.IFunction,
	ce *types.ClassEntry,
	child_dup_callback func(types.IFunction),
	check_only types.ZendBool,
	checked types.ZendBool,
) InheritanceStatus {
	var child_flags uint32
	var parent_flags uint32 = parent.GetFnFlags()
	var proto types.IFunction
	if checked == 0 && (parent_flags&AccFinal) != 0 {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot override final method %s::%s()", ZEND_FN_SCOPE_NAME(parent), child.GetFunctionName().GetVal())
	}
	child_flags = child.GetFnFlags()

	/* You cannot change from static to non static and vice versa.
	 */

	if checked == 0 && (child_flags&AccStatic) != (parent_flags&AccStatic) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		if (child_flags & AccStatic) != 0 {
			faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non static method %s::%s() static in class %s", ZEND_FN_SCOPE_NAME(parent), child.GetFunctionName().GetVal(), ZEND_FN_SCOPE_NAME(child))
		} else {
			faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make static method %s::%s() non static in class %s", ZEND_FN_SCOPE_NAME(parent), child.GetFunctionName().GetVal(), ZEND_FN_SCOPE_NAME(child))
		}
	}

	/* Disallow making an inherited method abstract. */

	if checked == 0 && (child_flags&AccAbstract) > (parent_flags&AccAbstract) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non abstract method %s::%s() abstract in class %s", ZEND_FN_SCOPE_NAME(parent), child.GetFunctionName().GetVal(), ZEND_FN_SCOPE_NAME(child))
	}
	if check_only == 0 && (parent_flags&(AccPrivate|AccChanged)) != 0 {
		child.SetIsChanged(true)
	}
	if (parent_flags & AccPrivate) != 0 {
		return INHERITANCE_SUCCESS
	}
	if parent.GetPrototype() != nil {
		proto = parent.GetPrototype()
	} else {
		proto = parent
	}
	if (parent_flags & AccCtor) != 0 {

		/* ctors only have a prototype if is abstract (or comes from an interface) */

		if !proto.IsAbstract() {
			return INHERITANCE_SUCCESS
		}
		parent = proto
	}
	if check_only == 0 && child.GetPrototype() != proto {
		for {
			if child.GetScope() != ce && child.GetType() == ZEND_USER_FUNCTION && child.GetOpArray().GetStaticVariables() == nil {
				if ce.IsInterface() {
					/* Few parent interfaces contain the same method */
					break
				} else if child_dup_callback != nil {

					/* op_array wasn't duplicated yet */

					var new_function types.IFunction = types.CopyOpArray(child.GetOpArray())
					child = new_function
					child_dup_callback(child)
				}
			}
			child.SetPrototype(proto)
			break
		}
	}

	/* Prevent derived classes from restricting access that was available in parent classes (except deriving from non-abstract ctors) */

	if checked == 0 && (child_flags&AccPppMask) > (parent_flags&AccPppMask) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Access level to %s::%s() must be %s (as in class %s)%s", ZEND_FN_SCOPE_NAME(child), child.GetFunctionName().GetVal(), ZendVisibilityString(parent_flags), ZEND_FN_SCOPE_NAME(parent), b.Cond((parent_flags&AccPublic) != 0, "", " or weaker"))
	}
	if checked == 0 {
		if check_only != 0 {
			var unresolved_class *types.String
			return ZendDoPerformImplementationCheck(&unresolved_class, child, parent)
		}
		PerformDelayableImplementationCheck(ce, child, parent, 0)
	}
	return INHERITANCE_SUCCESS
}
func DoInheritMethod(key string, parent types.IFunction, ce *types.ClassEntry, is_interface types.ZendBool, checked types.ZendBool) {
	var func_ = ce.FunctionTable().Get(key)
	if func_ != nil {
		if is_interface != 0 && func_ == parent {
			/* The same method in interface may be inherited few times */
			return
		}
		dupCallback := func(f types.IFunction) { ce.FunctionTable().UpdateDirect(key, f) }
		if checked != 0 {
			DoInheritanceCheckOnMethodEx(func_, parent, ce, dupCallback, 0, checked)
		} else {
			DoInheritanceCheckOnMethodEx(func_, parent, ce, dupCallback, 0, 0)
		}
	} else {
		if is_interface != 0 || parent.IsAbstract() {
			ce.SetIsImplicitAbstractClass(true)
		}
		parent = ZendDuplicateFunction(parent, ce, is_interface)
		ce.FunctionTable().Add(key, parent)
		// todo 考虑下 interface 是否会有多个同名函数同时存在
		//if is_interface == 0 {
		//	types._zendHashAppendPtr(ce.GetFunctionTable(), key, parent)
		//} else {
		//	types.ZendHashAddNewPtr(ce.GetFunctionTable(), key.GetStr(), parent)
		//}
	}
}
func PropertyTypesCompatible(parent_info *ZendPropertyInfo, child_info *ZendPropertyInfo) InheritanceStatus {
	var parent_name *types.String
	var child_name *types.String
	var parent_type_ce *types.ClassEntry
	var child_type_ce *types.ClassEntry
	if parent_info.GetType() == child_info.GetType() {
		return INHERITANCE_SUCCESS
	}
	if !(parent_info.GetType().IsClass()) || !(child_info.GetType().IsClass()) || parent_info.GetType().AllowNull() != child_info.GetType().AllowNull() {
		return INHERITANCE_ERROR
	}
	if parent_info.GetType().IsCe() {
		parent_name = types.ZEND_TYPE_CE(parent_info.GetType()).GetName()
	} else {
		parent_name = ResolveClassName(parent_info.GetCe(), parent_info.GetType().Name())
	}
	if child_info.GetType().IsCe() {
		child_name = types.ZEND_TYPE_CE(child_info.GetType()).GetName()
	} else {
		child_name = ResolveClassName(child_info.GetCe(), child_info.GetType().Name())
	}
	if ascii.StrCaseEquals(parent_name.GetStr(), child_name.GetStr()) {
		return INHERITANCE_SUCCESS
	}

	/* Check for class aliases */

	if parent_info.GetType().IsCe() {
		parent_type_ce = parent_info.GetType().Ce()
	} else {
		parent_type_ce = LookupClass(parent_info.GetCe(), parent_name)
	}
	if child_info.GetType().IsCe() {
		child_type_ce = child_info.GetType().Ce()
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
	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type of %s::$%s must be %s%s (as in class %s)", child.GetCe().GetName().GetVal(), ZendGetUnmangledPropertyName(child.GetName()), b.Cond(parent.GetType().AllowNull(), "?", ""), b.CondF(parent.GetType().IsClass(), func() []byte {
		return b.CondF(parent.GetType().IsCe(), func() *types.String { return types.ZEND_TYPE_CE(parent.GetType()).GetName() }, func() *types.String { return ResolveClassName(parent.GetCe(), parent.GetType().Name()) }).GetVal()
	}, func() *byte { return types.ZendGetTypeByConst(parent.GetType().Code()) }), parent.GetCe().GetName().GetVal())
}
func DoInheritProperty(parent_info *ZendPropertyInfo, key *types.String, ce *types.ClassEntry) {
	var child *types.Zval = ce.GetPropertiesInfo().KeyFind(key.GetStr())
	var child_info *ZendPropertyInfo
	if child != nil {
		child_info = child.GetPtr()
		if parent_info.HasFlags(AccPrivate | AccChanged) {
			child_info.SetIsChanged(true)
		}
		if !parent_info.IsPrivate() {
			if (parent_info.GetFlags() & AccStatic) != (child_info.GetFlags() & AccStatic) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot redeclare %s%s::$%s as %s%s::$%s", b.Cond(parent_info.IsStatic(), "static ", "non static "), ce.GetParent().name.GetVal(), key.GetVal(), b.Cond(child_info.IsStatic(), "static ", "non static "), ce.GetName().GetVal(), key.GetVal())
			}
			if (child_info.GetFlags() & AccPppMask) > (parent_info.GetFlags() & AccPppMask) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access level to %s::$%s must be %s (as in class %s)%s", ce.GetName().GetVal(), key.GetVal(), ZendVisibilityString(parent_info.GetFlags()), ce.GetParent().name.GetVal(), b.Cond(parent_info.IsPublic(), "", " or weaker"))
			} else if !child_info.IsStatic() {
				var parent_num int = OBJ_PROP_TO_NUM(parent_info.GetOffset())
				var child_num int = OBJ_PROP_TO_NUM(child_info.GetOffset())

				/* Don't keep default properties in GC (they may be freed by opcache) */

				ZvalPtrDtorNogc(&ce.GetDefaultPropertiesTable()[parent_num])
				ce.GetDefaultPropertiesTable()[parent_num] = ce.GetDefaultPropertiesTable()[child_num]
				ce.GetDefaultPropertiesTable()[child_num].SetUndef()
				child_info.SetOffset(parent_info.GetOffset())
			}
			if parent_info.GetType().IsSet() {
				var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
				if status == INHERITANCE_ERROR {
					EmitIncompatiblePropertyError(child_info, parent_info)
				}
				if status == INHERITANCE_UNRESOLVED {
					AddPropertyCompatibilityObligation(ce, child_info, parent_info)
				}
			} else if child_info.GetType().IsSet() && !(parent_info.GetType().IsSet()) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type of %s::$%s must not be defined (as in class %s)", ce.GetName().GetVal(), key.GetVal(), ce.GetParent().name.GetVal())
			}
		}
	} else {
		if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
			child_info = ZendDuplicatePropertyInfoInternal(parent_info)
		} else {
			child_info = parent_info
		}
		types._zendHashAppendPtr(ce.GetPropertiesInfo(), key, child_info)
	}
}
func DoImplementInterface(ce *types.ClassEntry, iface *types.ClassEntry) {
	if !ce.IsInterface() && iface.GetInterfaceGetsImplemented() && iface.GetInterfaceGetsImplemented()(iface, ce) == types.FAILURE {
		faults.ErrorNoreturn(faults.E_CORE_ERROR, "Class %s could not implement interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
	}

	/* This should be prevented by the class lookup logic. */

	b.Assert(ce != iface)

	/* This should be prevented by the class lookup logic. */
}
func ZendDoInheritInterfaces(ce *types.ClassEntry, iface *types.ClassEntry) {
	/* expects interface to be contained in ce's interface list already */

	var i uint32
	var ce_num uint32
	var if_num uint32 = iface.GetNumInterfaces()
	var entry *types.ClassEntry
	ce_num = ce.GetNumInterfaces()
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		ce.SetInterfaces((**types.ClassEntry)(realloc(ce.GetInterfaces(), b.SizeOf("zend_class_entry *")*(ce_num+if_num))))
	} else {
		ce.SetInterfaces((**types.ClassEntry)(Erealloc(ce.GetInterfaces(), b.SizeOf("zend_class_entry *")*(ce_num+if_num))))
	}

	/* Inherit the interfaces, only if they're not already inherited by the class */

	for b.PostDec(&if_num) {
		entry = iface.GetInterfaces()[if_num]
		for i = 0; i < ce_num; i++ {
			if ce.GetInterfaces()[i] == entry {
				break
			}
		}
		if i == ce_num {
			ce.GetInterfaces()[b.PostInc(&(ce.GetNumInterfaces()))] = entry
		}
	}
	ce.SetIsResolvedInterfaces(true)

	/* and now call the implementing handlers */

	for ce_num < ce.GetNumInterfaces() {
		DoImplementInterface(ce, ce.GetInterfaces()[b.PostInc(&ce_num)])
	}

	/* and now call the implementing handlers */
}
func DoInheritClassConstant(name *types.String, parentConst *ZendClassConstant, ce *types.ClassEntry) {
	var c = ce.ConstantsTable().Get(name.GetStr())
	if c != nil {
		if (c.GetValue().GetAccessFlags() & AccPppMask) > (parentConst.GetValue().GetAccessFlags() & AccPppMask) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access level to %s::%s must be %s (as in class %s)%s", ce.Name(), name.GetStr(), ZendVisibilityString(parentConst.GetValue().GetAccessFlags()), ce.GetParent().Name(), b.Cond((parentConst.GetValue().GetAccessFlags()&AccPublic) != 0, "", " or weaker"))
		}
	} else if (parentConst.GetValue().GetAccessFlags() & AccPrivate) == 0 {
		if parentConst.GetValue().IsConstant() {
			ce.SetIsConstantsUpdated(false)
		}
		if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
			parentConst = CopyClassConstant(parentConst)
		}
		ce.ConstantsTable().Add(name.GetStr(), parentConst)
	}
}
func ZendBuildPropertiesInfoTable(ce *types.ClassEntry) {
	var table **ZendPropertyInfo
	var prop **ZendPropertyInfo
	var size int
	if ce.GetDefaultPropertiesCount() == 0 {
		return
	}
	b.Assert(ce.GetPropertiesInfoTable() == nil)
	size = b.SizeOf("zend_property_info *") * ce.GetDefaultPropertiesCount()
	if ce.GetType() == ZEND_USER_CLASS {
		table = ZendArenaAlloc(CG__().GetArena(), size)
		ce.SetPropertiesInfoTable(table)
	} else {
		table = Pemalloc(size, 1)
		ce.SetPropertiesInfoTable(table)
	}

	/* Dead slots may be left behind during inheritance. Make sure these are NULLed out. */

	memset(table, 0, size)
	if ce.GetParent() && ce.GetParent().default_properties_count != 0 {
		var parent_table **ZendPropertyInfo = ce.GetParent().properties_info_table
		memcpy(table, parent_table, b.SizeOf("zend_property_info *")*ce.GetParent().default_properties_count)

		/* Child did not add any new properties, we are done */

		if ce.GetDefaultPropertiesCount() == ce.GetParent().default_properties_count {
			return
		}

		/* Child did not add any new properties, we are done */

	}
	var __ht *types.Array = ce.GetPropertiesInfo()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		prop = _z.GetPtr()
		if prop.GetCe() == ce && !prop.IsStatic() {
			table[OBJ_PROP_TO_NUM(prop.GetOffset())] = prop
		}
	}
}
func ZendDoInheritanceEx(ce *types.ClassEntry, parent_ce *types.ClassEntry, checked bool) {
	var property_info *ZendPropertyInfo
	var func_ types.IFunction
	var key *types.String
	if ce.IsInterface() {

		/* Interface can only inherit other interfaces */

		if !parent_ce.IsInterface() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Interface %s may not inherit from class (%s)", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Interface can only inherit other interfaces */

	} else if parent_ce.HasCeFlags(AccInterface | AccTrait | AccFinal) {

		/* Class declaration must not extend traits or interfaces */

		if parent_ce.IsInterface() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot extend from interface %s", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		} else if parent_ce.IsTrait() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot extend from trait %s", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Class must not extend a final class */

		if parent_ce.IsFinal() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s may not inherit from final class (%s)", ce.GetName().GetVal(), parent_ce.GetName().GetVal())
		}

		/* Class must not extend a final class */

	}
	if ce.GetParentName() {
		// types.ZendStringReleaseEx(ce.GetParentName(), 0)
	}
	ce.SetParent(parent_ce)
	ce.SetIsResolvedParent(true)

	/* Inherit interfaces */

	if parent_ce.GetNumInterfaces() != 0 {
		if !ce.IsImplementInterfaces() {
			ZendDoInheritInterfaces(ce, parent_ce)
		} else {
			var i uint32
			for i = 0; i < parent_ce.GetNumInterfaces(); i++ {
				DoImplementInterface(ce, parent_ce.GetInterfaces()[i])
			}
		}
	}

	/* Inherit properties */

	if parent_ce.GetDefaultPropertiesCount() != 0 {
		var src *types.Zval
		var dst *types.Zval
		var end *types.Zval
		if ce.GetDefaultPropertiesCount() != 0 {
			var table *types.Zval = Pemalloc(b.SizeOf("zval")*(ce.GetDefaultPropertiesCount()+parent_ce.GetDefaultPropertiesCount()), ce.GetType() == ZEND_INTERNAL_CLASS)
			src = ce.GetDefaultPropertiesTable() + ce.GetDefaultPropertiesCount()
			end = table + parent_ce.GetDefaultPropertiesCount()
			dst = end + ce.GetDefaultPropertiesCount()
			ce.SetDefaultPropertiesTable(table)
			for {
				dst--
				src--
				types.ZVAL_COPY_VALUE_PROP(dst, src)
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
		if parent_ce.GetType() != ce.GetType() {

			/* User class extends internal */

			for {
				dst--
				src--
				types.ZVAL_COPY_OR_DUP_PROP(dst, src)
				if dst.IsConstant() {
					ce.SetIsConstantsUpdated(false)
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
				types.ZVAL_COPY_PROP(dst, src)
				if dst.IsConstant() {
					ce.SetIsConstantsUpdated(false)
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
		var src *types.Zval
		var dst *types.Zval
		var end *types.Zval
		if ce.GetDefaultStaticMembersCount() != 0 {
			var table *types.Zval = Pemalloc(b.SizeOf("zval")*(ce.GetDefaultStaticMembersCount()+parent_ce.GetDefaultStaticMembersCount()), ce.GetType() == ZEND_INTERNAL_CLASS)
			src = ce.GetDefaultStaticMembersTable() + ce.GetDefaultStaticMembersCount()
			end = table + parent_ce.GetDefaultStaticMembersCount()
			dst = end + ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(table)
			for {
				dst--
				src--
				types.ZVAL_COPY_VALUE(dst, src)
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
		if parent_ce.GetType() != ce.GetType() {

			/* User class extends internal */

			if CE_STATIC_MEMBERS(parent_ce) == nil {
				ZendClassInitStatics(parent_ce)
			}
			if ZendUpdateClassConstants(parent_ce) != types.SUCCESS {
				b.Assert(false)
			}
			src = CE_STATIC_MEMBERS(parent_ce) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.IsIndirect() {
					dst.SetIndirect(src.GetZv())
				} else {
					dst.SetIndirect(src)
				}
				if dst == end {
					break
				}
			}
		} else if ce.GetType() == ZEND_USER_CLASS {
			if CE_STATIC_MEMBERS(parent_ce) == nil {
				b.Assert(parent_ce.HasCeFlags(AccImmutable | AccPreloaded))
				ZendClassInitStatics(parent_ce)
			}
			src = CE_STATIC_MEMBERS(parent_ce) + parent_ce.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.IsIndirect() {
					dst.SetIndirect(src.GetZv())
				} else {
					dst.SetIndirect(src)
				}
				if types.Z_INDIRECT_P(dst).IsConstant() {
					ce.SetIsConstantsUpdated(false)
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
				if src.IsIndirect() {
					dst.SetIndirect(src.GetZv())
				} else {
					dst.SetIndirect(src)
				}
				if dst == end {
					break
				}
			}
		}
		ce.SetDefaultStaticMembersCount(ce.GetDefaultStaticMembersCount() + parent_ce.GetDefaultStaticMembersCount())
		if ce.GetStaticMembersTablePtr() == nil {
			b.Assert(ce.GetType() == ZEND_INTERNAL_CLASS)
			if CurrEX() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {

				/* internal class loaded by dl() */

				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())

				/* internal class loaded by dl() */

			}
		}
	}
	var __ht *types.Array = ce.GetPropertiesInfo()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		property_info = _z.GetPtr()
		if property_info.GetCe() == ce {
			if property_info.IsStatic() {
				property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultStaticMembersCount())
			} else {
				property_info.SetOffset(property_info.GetOffset() + parent_ce.GetDefaultPropertiesCount()*b.SizeOf("zval"))
			}
		}
	}
	if parent_ce.GetPropertiesInfo().Len() {
		ce.GetPropertiesInfo().Extend(ce.GetPropertiesInfo().Len() + parent_ce.GetPropertiesInfo().Len())
		var __ht *types.Array = parent_ce.GetPropertiesInfo()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			property_info = _z.GetPtr()
			DoInheritProperty(property_info, key, ce)
		}
	}
	if parent_ce.GetConstantsTable().Len() {
		var c *ZendClassConstant
		ce.GetConstantsTable().Extend(ce.GetConstantsTable().Len() + parent_ce.GetConstantsTable().Len())
		var __ht *types.Array = parent_ce.GetConstantsTable()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			c = _z.GetPtr()
			DoInheritClassConstant(key, c, ce)
		}
	}
	if parent_ce.FunctionTable().Len() != 0 {
		if checked {
			parent_ce.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
				DoInheritMethod(key, func_, ce, 0, 1)
			})
		} else {
			parent_ce.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
				DoInheritMethod(key, func_, ce, 0, 0)
			})
		}
	}
	DoInheritParentConstructor(ce)
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		if ce.IsImplicitAbstractClass() {
			ce.SetIsExplicitAbstractClass(true)
		}
	}
	ce.AddCeFlags(parent_ce.GetCeFlags() & (AccHasStaticInMethods | AccHasTypeHints | AccUseGuards))
}
func DoInheritConstantCheck(child_constants_table *types.Array, parent_constant *ZendClassConstant, name *types.String, iface *types.ClassEntry) types.ZendBool {
	var zv *types.Zval = child_constants_table.KeyFind(name.GetStr())
	var old_constant *ZendClassConstant
	if zv != nil {
		old_constant = (*ZendClassConstant)(zv.GetPtr())
		if old_constant.GetCe() != parent_constant.GetCe() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot inherit previously-inherited or override constant %s from interface %s", name.GetVal(), iface.GetName().GetVal())
		}
		return 0
	}
	return 1
}
func DoInheritIfaceConstant(name *types.String, c *ZendClassConstant, ce *types.ClassEntry, iface *types.ClassEntry) {
	if DoInheritConstantCheck(ce.GetConstantsTable(), c, name, iface) != 0 {
		var ct *ZendClassConstant
		if c.GetValue().IsConstant() {
			ce.SetIsConstantsUpdated(false)
		}
		if (ce.GetType() & ZEND_INTERNAL_CLASS) != 0 {
			ct = Pemalloc(b.SizeOf("zend_class_constant"), 1)
			memcpy(ct, c, b.SizeOf("zend_class_constant"))
			c = ct
		}
		types.ZendHashUpdatePtr(ce.GetConstantsTable(), name.GetStr(), c)
	}
}
func DoInterfaceImplementation(ce *types.ClassEntry, iface *types.ClassEntry) {
	var func_ types.IFunction
	var key *types.String
	var c *ZendClassConstant
	var __ht *types.Array = iface.GetConstantsTable()
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		c = _z.GetPtr()
		DoInheritIfaceConstant(key, c, ce, iface)
	}
	iface.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
		DoInheritMethod(key, func_, ce, 1, 0)
	})
	DoImplementInterface(ce, iface)
	if iface.GetNumInterfaces() != 0 {
		ZendDoInheritInterfaces(ce, iface)
	}
}
func ZendDoImplementInterface(ce *types.ClassEntry, iface *types.ClassEntry) {
	var i uint32
	var ignore uint32 = 0
	var current_iface_num uint32 = ce.GetNumInterfaces()
	var parent_iface_num uint32 = b.CondF1(ce.GetParent(), func() __auto__ { return ce.GetParent().num_interfaces }, 0)
	var key *types.String
	var c *ZendClassConstant
	b.Assert(ce.IsLinked())
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		if ce.GetInterfaces()[i] == nil {
			memmove(ce.GetInterfaces()+i, ce.GetInterfaces()+i+1, b.SizeOf("zend_class_entry *")*(b.PreDec(&(ce.GetNumInterfaces()))-i))
			i--
		} else if ce.GetInterfaces()[i] == iface {
			if i < parent_iface_num {
				ignore = 1
			} else {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
			}
		}
	}
	if ignore != 0 {

		/* Check for attempt to redeclare interface constants */

		var __ht *types.Array = ce.GetConstantsTable()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			c = _z.GetPtr()
			DoInheritConstantCheck(iface.GetConstantsTable(), c, key, iface)
		}

		/* Check for attempt to redeclare interface constants */

	} else {
		if ce.GetNumInterfaces() >= current_iface_num {
			if ce.GetType() == ZEND_INTERNAL_CLASS {
				ce.SetInterfaces((**types.ClassEntry)(realloc(ce.GetInterfaces(), b.SizeOf("zend_class_entry *")*b.PreInc(&current_iface_num))))
			} else {
				ce.SetInterfaces((**types.ClassEntry)(Erealloc(ce.GetInterfaces(), b.SizeOf("zend_class_entry *")*b.PreInc(&current_iface_num))))
			}
		}
		ce.GetInterfaces()[b.PostInc(&(ce.GetNumInterfaces()))] = iface
		DoInterfaceImplementation(ce, iface)
	}
}
func ZendDoImplementInterfaces(ce *types.ClassEntry, interfaces **types.ClassEntry) {
	var iface *types.ClassEntry
	var num_parent_interfaces uint32 = b.CondF1(ce.GetParent(), func() __auto__ { return ce.GetParent().num_interfaces }, 0)
	var num_interfaces uint32 = num_parent_interfaces
	var key *types.String
	var c *ZendClassConstant
	var i uint32
	var j uint32
	for i = 0; i < ce.GetNumInterfaces(); i++ {
		iface = interfaces[num_parent_interfaces+i]
		if !iface.IsLinked() {
			AddDependencyObligation(ce, iface)
		}
		if !iface.IsInterface() {
			Efree(interfaces)
			faults.ErrorNoreturn(faults.E_ERROR, "%s cannot implement %s - it is not an interface", ce.GetName().GetVal(), iface.GetName().GetVal())
			return
		}
		for j = 0; j < num_interfaces; j++ {
			if interfaces[j] == iface {
				if j >= num_parent_interfaces {
					Efree(interfaces)
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ce.GetName().GetVal(), iface.GetName().GetVal())
					return
				}

				/* skip duplications */

				var __ht *types.Array = ce.GetConstantsTable()
				for _, _p := range __ht.ForeachData() {
					var _z *types.Zval = _p.GetVal()

					key = _p.GetKey()
					c = _z.GetPtr()
					DoInheritConstantCheck(iface.GetConstantsTable(), c, key, iface)
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
		// types.ZendStringReleaseEx(ce.GetInterfaceNames()[i].name, 0)
		// types.ZendStringReleaseEx(ce.GetInterfaceNames()[i].lc_name, 0)
	}
	Efree(ce.GetInterfaceNames())
	ce.SetNumInterfaces(num_interfaces)
	ce.SetInterfaces(interfaces)
	ce.SetIsResolvedInterfaces(true)
	i = num_parent_interfaces
	for ; i < ce.GetNumInterfaces(); i++ {
		DoInterfaceImplementation(ce, ce.GetInterfaces()[i])
	}
}
func ZendAddMagicMethods(ce *types.ClassEntry, mname *types.String, fe types.IFunction) {
	if mname.GetStr() == "serialize" {
		ce.SetSerializeFunc(fe)
	} else if mname.GetStr() == "unserialize" {
		ce.SetUnserializeFunc(fe)
	} else if ce.GetName().GetLen() != mname.GetLen() && (mname.GetVal()[0] != '_' || mname.GetVal()[1] != '_') {

	} else if mname.GetStr() == ZEND_CLONE_FUNC_NAME {
		ce.SetClone(fe)
	} else if mname.GetStr() == ZEND_CONSTRUCTOR_FUNC_NAME {
		if ce.GetConstructor() != nil && (!(ce.GetParent()) || ce.GetConstructor() != ce.GetParent().constructor) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ce.GetName().GetVal())
		}
		ce.SetConstructor(fe)
	} else if mname.GetStr() == ZEND_DESTRUCTOR_FUNC_NAME {
		ce.SetDestructor(fe)
	} else if mname.GetStr() == ZEND_GET_FUNC_NAME {
		ce.SetGet(fe)
		ce.SetIsUseGuards(true)
	} else if mname.GetStr() == ZEND_SET_FUNC_NAME {
		ce.SetSet(fe)
		ce.SetIsUseGuards(true)
	} else if mname.GetStr() == ZEND_CALL_FUNC_NAME {
		ce.SetCall(fe)
	} else if mname.GetStr() == ZEND_UNSET_FUNC_NAME {
		ce.SetUnset(fe)
		ce.SetIsUseGuards(true)
	} else if mname.GetStr() == ZEND_ISSET_FUNC_NAME {
		ce.SetIsset(fe)
		ce.SetIsUseGuards(true)
	} else if mname.GetStr() == ZEND_CALLSTATIC_FUNC_NAME {
		ce.SetCallstatic(fe)
	} else if mname.GetStr() == ZEND_TOSTRING_FUNC_NAME {
		ce.SetTostring(fe)
	} else if mname.GetStr() == ZEND_DEBUGINFO_FUNC_NAME {
		ce.SetDebugInfo(fe)
	} else if ce.GetName().GetLen() == mname.GetLen() {
		var lowercase_name *types.String = ZendStringTolower(ce.GetName())
		// lowercase_name = types.ZendNewInternedString(lowercase_name)
		if !(memcmp(mname.GetVal(), lowercase_name.GetVal(), mname.GetLen())) {
			if ce.GetConstructor() != nil && (!(ce.GetParent()) || ce.GetConstructor() != ce.GetParent().constructor) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ce.GetName().GetVal())
			}
			ce.SetConstructor(fe)
			fe.SetIsCtor(true)
		}
		// types.ZendStringReleaseEx(lowercase_name, 0)
	}
}
func ZendAddTraitMethod(ce *types.ClassEntry, name *byte, key *types.String, fn types.IFunction, overridden **types.Array) {
	var existing_fn types.IFunction = nil
	var new_fn types.IFunction
	if b.Assign(&existing_fn, ce.FunctionTable().Get(key.GetStr())) != nil {

		/* if it is the same function with the same visibility and has not been assigned a class scope yet, regardless
		 * of where it is coming from there is no conflict and we do not need to add it again */

		if existing_fn.GetOpArray().GetOpcodes() == fn.GetOpArray().GetOpcodes() && (existing_fn.GetFnFlags()&AccPppMask) == (fn.GetFnFlags()&AccPppMask) && (existing_fn.GetScope().GetCeFlags()&AccTrait) == AccTrait {
			return
		}
		if existing_fn.GetScope() == ce {

			/* members from the current class override trait methods */

			if (*overridden) != nil {
				if b.Assign(&existing_fn, types.ZendHashFindPtr(*overridden, key.GetStr())) != nil {
					if existing_fn.IsAbstract() {

						/* Make sure the trait method is compatible with previosly declared abstract method */

						PerformDelayableImplementationCheck(ce, fn, existing_fn, 1)

						/* Make sure the trait method is compatible with previosly declared abstract method */

					}
					if fn.IsAbstract() {

						/* Make sure the abstract declaration is compatible with previous declaration */

						PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
						return
					}
				}
			} else {
				ALLOC_HASHTABLE(*overridden)
				*overridden = types.MakeArrayEx(8, OverriddenPtrDtor, 0)
			}
			types.ZendHashUpdateMem(*overridden, key.GetStr(), fn, b.SizeOf("zend_function"))
			return
		} else if fn.IsAbstract() && !existing_fn.IsAbstract() {

			/* Make sure the abstract declaration is compatible with previous declaration */

			PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
			return
		} else if existing_fn.GetScope().IsTrait() && !existing_fn.IsAbstract() {

			/* two traits can't define the __special__  same non-abstract method */

			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Trait method %s has not been applied, because there are collisions with other trait methods on %s", name, ce.GetName().GetVal())

			/* two traits can't define the __special__  same non-abstract method */

		} else {

			/* inherited members are overridden by members inserted by traits */

			DoInheritanceCheckOnMethodEx(fn, existing_fn, ce, nil, 0, 0)
			fn.SetPrototype(nil)
		}
	}
	if fn.GetType() == ZEND_INTERNAL_FUNCTION {
		new_fn = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_internal_function"))
		memcpy(new_fn, fn, b.SizeOf("zend_internal_function"))
		new_fn.SetIsArenaAllocated(true)
	} else {
		new_fn = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_op_array"))
		memcpy(new_fn, fn, b.SizeOf("zend_op_array"))
		new_fn.GetOpArray().SetIsTraitClone(true)
		new_fn.GetOpArray().SetIsImmutable(false)
	}
	FunctionAddRef(new_fn)
	ce.FunctionTable().Update(key.GetStr(), new_fn)
	ZendAddMagicMethods(ce, key, new_fn)
}
func ZendFixupTraitMethod(fn types.IFunction, ce *types.ClassEntry) {
	if (fn.GetScope().GetCeFlags() & AccTrait) == AccTrait {
		fn.SetScope(ce)
		if fn.IsAbstract() {
			ce.SetIsImplicitAbstractClass(true)
		}
		if fn.GetType() == ZEND_USER_FUNCTION && fn.GetOpArray().GetStaticVariables() != nil {
			ce.SetIsHasStaticInMethods(true)
		}
	}
}
func ZendTraitsCopyFunctions(
	fnname string,
	fn types.IFunction,
	ce *types.ClassEntry,
	overridden **types.Array,
	exclude_table *types.Array,
	aliases **types.ClassEntry,
) {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	var lcname *types.String
	var fn_copy types.IFunction
	var i int

	/* apply aliases which are qualified with a class name, there should not be any ambiguity */

	if ce.GetTraitAliases() != nil {
		alias_ptr = ce.GetTraitAliases()
		alias = *alias_ptr
		i = 0
		for alias != nil {

			/* Scope unset or equal to the function we compare to, and the alias applies to fn */

			if alias.GetAlias() != nil && (aliases[i] == nil || fn.GetScope() == aliases[i]) && alias.GetTraitMethod().GetMethodName().GetLen() == fnname.GetLen() && ZendBinaryStrcasecmp(alias.GetTraitMethod().GetMethodName().GetStr(), fnname.GetStr()) == 0 {
				fn_copy = types.CopyFunction(fn)

				/* if it is 0, no modifieres has been changed */

				if alias.GetModifiers() != 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&AccPppMask)
				}
				lcname = ZendStringTolower(alias.GetAlias())
				ZendAddTraitMethod(ce, alias.GetAlias().GetVal(), lcname, &fn_copy, overridden)
				// types.ZendStringReleaseEx(lcname, 0)

				/* Record the trait from which this alias was resolved. */

				if aliases[i] == nil {
					aliases[i] = fn.GetScope()
				}
				if alias.GetTraitMethod().GetClassName() == nil {

					/* TODO: try to avoid this assignment (it's necessary only for reflection) */

					alias.GetTraitMethod().SetClassName(fn.GetScope().GetName().Copy())

					/* TODO: try to avoid this assignment (it's necessary only for reflection) */

				}
			}
			alias_ptr++
			alias = *alias_ptr
			i++
		}
	}
	if exclude_table == nil || exclude_table.KeyFind(fnname.GetStr()) == nil {

		/* is not in hashtable, thus, function is not to be excluded */

		memcpy(&fn_copy, fn, b.CondF(fn.GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))

		/* apply aliases which have not alias name, just setting visibility */

		if ce.GetTraitAliases() != nil {
			alias_ptr = ce.GetTraitAliases()
			alias = *alias_ptr
			i = 0
			for alias != nil {

				/* Scope unset or equal to the function we compare to, and the alias applies to fn */

				if alias.GetAlias() == nil && alias.GetModifiers() != 0 && (aliases[i] == nil || fn.GetScope() == aliases[i]) && alias.GetTraitMethod().GetMethodName().GetLen() == fnname.GetLen() && ZendBinaryStrcasecmp(alias.GetTraitMethod().GetMethodName().GetStr(), fnname.GetStr()) == 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&AccPppMask)

					/** Record the trait from which this alias was resolved. */

					if aliases[i] == nil {
						aliases[i] = fn.GetScope()
					}
					if alias.GetTraitMethod().GetClassName() == nil {

						/* TODO: try to avoid this assignment (it's necessary only for reflection) */

						alias.GetTraitMethod().SetClassName(fn.GetScope().GetName().Copy())

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
func ZendCheckTraitUsage(ce *types.ClassEntry, trait *types.ClassEntry, traits **types.ClassEntry) uint32 {
	var i uint32
	if (trait.GetCeFlags() & AccTrait) != AccTrait {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s is not a trait, Only traits may be used in 'as' and 'insteadof' statements", trait.GetName().GetVal())
		return 0
	}
	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == trait {
			return i
		}
	}
	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Required Trait %s wasn't added to %s", trait.GetName().GetVal(), ce.GetName().GetVal())
	return 0
}
func ZendTraitsInitTraitStructures(ce *types.ClassEntry, traits **types.ClassEntry, exclude_tables_ptr ***types.Array, aliases_ptr ***types.ClassEntry) {
	var i int
	var j int = 0
	var precedences **ZendTraitPrecedence
	var cur_precedence *ZendTraitPrecedence
	var cur_method_ref *ZendTraitMethodReference
	var lcname *types.String
	var exclude_tables **types.Array = nil
	var aliases **types.ClassEntry = nil
	var trait *types.ClassEntry

	/* resolve class references */

	if ce.GetTraitPrecedences() != nil {
		exclude_tables = Ecalloc(ce.GetNumTraits(), b.SizeOf("HashTable *"))
		i = 0
		precedences = ce.GetTraitPrecedences()
		ce.SetTraitPrecedences(nil)
		for b.Assign(&cur_precedence, precedences[i]) {

			/** Resolve classes for all precedence operations. */

			cur_method_ref = cur_precedence.GetTraitMethod()
			trait = ZendFetchClass(cur_method_ref.GetClassName().GetStr(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if trait == nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", cur_method_ref.GetClassName().GetVal())
			}
			ZendCheckTraitUsage(ce, trait, traits)

			/** Ensure that the preferred method is actually available. */

			lcname = ZendStringTolower(cur_method_ref.GetMethodName())
			if !trait.FunctionTable().Exists(lcname.GetStr()) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "A precedence rule was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.GetMethodName().GetVal())
			}

			/** With the other traits, we are more permissive.
			  We do not give errors for those. This allows to be more
			  defensive in such definitions.
			  However, we want to make sure that the insteadof declaration
			  is consistent in itself.
			*/

			for j = 0; j < cur_precedence.GetNumExcludes(); j++ {
				var class_name *types.String = cur_precedence.GetExcludeClassNames()[j]
				var exclude_ce *types.ClassEntry = ZendFetchClass(class_name.GetStr(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				var trait_num uint32
				if exclude_ce == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", class_name.GetVal())
				}
				trait_num = ZendCheckTraitUsage(ce, exclude_ce, traits)
				if exclude_tables[trait_num] == nil {
					ALLOC_HASHTABLE(exclude_tables[trait_num])
					exclude_tables[trait_num] = types.MakeArrayEx(0, nil, 0)
				}
				if types.ZendHashAddEmptyElement(exclude_tables[trait_num], lcname.GetStr()) == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Failed to evaluate a trait precedence (%s). Method of trait %s was defined to be excluded multiple times", precedences[i].GetTraitMethod().GetMethodName().GetVal(), exclude_ce.GetName().GetVal())
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

				if trait == exclude_ce {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Inconsistent insteadof definition. "+"The method %s is to be used from %s, but %s is also on the exclude list", cur_method_ref.GetMethodName().GetVal(), trait.GetName().GetVal(), trait.GetName().GetVal())
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

			}
			// types.ZendStringReleaseEx(lcname, 0)
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
				cur_method_ref = ce.GetTraitAliases()[i].GetTraitMethod()
				trait = ZendFetchClass(cur_method_ref.GetClassName().GetStr(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if trait == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", cur_method_ref.GetClassName().GetVal())
				}
				ZendCheckTraitUsage(ce, trait, traits)
				aliases[i] = trait

				/** And, ensure that the referenced method is resolvable, too. */

				lcname = ZendStringTolower(cur_method_ref.GetMethodName())
				if !trait.FunctionTable().Exists(lcname.GetStr()) {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "An alias was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.GetMethodName().GetVal())
				}
				// types.ZendStringReleaseEx(lcname, 0)
			}
			i++
		}
	}
	*exclude_tables_ptr = exclude_tables
	*aliases_ptr = aliases
}
func ZendDoTraitsMethodBinding(ce *types.ClassEntry, traits []*types.ClassEntry, exclude_tables []*types.Array, aliases **types.ClassEntry) {
	var i uint32
	var overridden *types.Array = nil
	if exclude_tables != nil {
		for i = 0; i < ce.GetNumTraits(); i++ {
			if traits[i] != nil {

				/* copies functions, applies defined aliasing, and excludes unused trait methods */
				traits[i].FunctionTable().Foreach(func(key string, fn types.IFunction) {
					ZendTraitsCopyFunctions(key, fn, ce, &overridden, exclude_tables[i], aliases)
				})
				if exclude_tables[i] != nil {
					exclude_tables[i].Destroy()
					FREE_HASHTABLE(exclude_tables[i])
					exclude_tables[i] = nil
				}
			}
		}
	} else {
		for i = 0; i < ce.GetNumTraits(); i++ {
			if traits[i] != nil {
				traits[i].FunctionTable().Foreach(func(key string, fn types.IFunction) {
					ZendTraitsCopyFunctions(key, fn, ce, &overridden, nil, aliases)
				})
			}
		}
	}
	ce.FunctionTable().Foreach(func(_ string, fn types.IFunction) {
		ZendFixupTraitMethod(fn, ce)
	})
	if overridden != nil {
		overridden.Destroy()
		FREE_HASHTABLE(overridden)
	}
}
func FindFirstDefinition(ce *types.ClassEntry, traits **types.ClassEntry, current_trait int, prop_name *types.String, coliding_ce *types.ClassEntry) *types.ClassEntry {
	var i int
	if coliding_ce == ce {
		for i = 0; i < current_trait; i++ {
			if traits[i] != nil && traits[i].GetPropertiesInfo().KeyExists(prop_name.GetStr()) {
				return traits[i]
			}
		}
	}
	return coliding_ce
}
func ZendDoTraitsPropertyBinding(ce *types.ClassEntry, traits **types.ClassEntry) {
	var i int
	var property_info *ZendPropertyInfo
	var coliding_prop *ZendPropertyInfo
	var prop_name *types.String
	var class_name_unused *byte
	var not_compatible types.ZendBool
	var prop_value *types.Zval
	var flags uint32
	var doc_comment *types.String

	/* In the following steps the properties are inserted into the property table
	 * for that, a very strict approach is applied:
	 * - check for compatibility, if not compatible with any property in class -> fatal
	 * - if compatible, then strict notice
	 */

	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == nil {
			continue
		}
		var __ht *types.Array = traits[i].GetPropertiesInfo()
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			property_info = _z.GetPtr()

			/* first get the unmangeld name if necessary,
			 * then check whether the property is already there
			 */

			flags = property_info.GetFlags()
			if (flags & AccPublic) != 0 {
				prop_name = property_info.GetName().Copy()
			} else {
				var pname *byte
				var pname_len int

				/* for private and protected we need to unmangle the names */

				ZendUnmanglePropertyNameEx(property_info.GetName(), &class_name_unused, &pname, &pname_len)
				prop_name = types.NewString(b.CastStr(pname, pname_len))
			}

			/* next: check for conflicts with current class */

			if b.Assign(&coliding_prop, types.ZendHashFindPtr(ce.GetPropertiesInfo(), prop_name.GetStr())) != nil {
				if coliding_prop.IsPrivate() && coliding_prop.GetCe() != ce {
					types.ZendHashDel(ce.GetPropertiesInfo(), prop_name.GetStr())
					flags |= AccChanged
				} else {
					not_compatible = 1
					if (coliding_prop.GetFlags()&(AccPppMask|AccStatic)) == (flags&(AccPppMask|AccStatic)) && PropertyTypesCompatible(property_info, coliding_prop) == INHERITANCE_SUCCESS {

						/* the flags are identical, thus, the properties may be compatible */

						var op1 *types.Zval
						var op2 *types.Zval
						var op1_tmp types.Zval
						var op2_tmp types.Zval
						if (flags & AccStatic) != 0 {
							op1 = ce.GetDefaultStaticMembersTable()[coliding_prop.GetOffset()]
							op2 = traits[i].GetDefaultStaticMembersTable()[property_info.GetOffset()]
							op1 = types.ZVAL_DEINDIRECT(op1)
							op2 = types.ZVAL_DEINDIRECT(op2)
						} else {
							op1 = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(coliding_prop.GetOffset())]
							op2 = traits[i].GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
						}

						/* if any of the values is a constant, we try to resolve it */

						if op1.IsConstant() {
							types.ZVAL_COPY_OR_DUP(&op1_tmp, op1)
							ZvalUpdateConstantEx(&op1_tmp, ce)
							op1 = &op1_tmp
						}
						if op2.IsConstant() {
							types.ZVAL_COPY_OR_DUP(&op2_tmp, op2)
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
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s and %s define the __special__  same property ($%s) in the composition of %s. However, the definition differs and is considered incompatible. Class was composed", FindFirstDefinition(ce, traits, i, prop_name, coliding_prop.GetCe()).GetName().GetVal(), property_info.GetCe().GetName().GetVal(), prop_name.GetVal(), ce.GetName().GetVal())
					}
					// types.ZendStringReleaseEx(prop_name, 0)
					continue
				}
			}

			/* property not found, so lets add it */

			if (flags & AccStatic) != 0 {
				prop_value = traits[i].GetDefaultStaticMembersTable()[property_info.GetOffset()]
				b.Assert(prop_value.GetType() != types.IS_INDIRECT)
			} else {
				prop_value = traits[i].GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
			}
			prop_value.TryAddRefcount()
			if property_info.GetDocComment() != nil {
				doc_comment = property_info.GetDocComment().Copy()
			} else {
				doc_comment = nil
			}
			if property_info.GetType().IsName() {
				//property_info.GetType().Name().AddRefcount()
			}
			ZendDeclareTypedProperty(ce, prop_name, prop_value, flags, doc_comment, property_info.GetType())
			// types.ZendStringReleaseEx(prop_name, 0)
		}
	}

	/* In the following steps the properties are inserted into the property table
	 * for that, a very strict approach is applied:
	 * - check for compatibility, if not compatible with any property in class -> fatal
	 * - if compatible, then strict notice
	 */
}
func ZendDoCheckForInconsistentTraitsAliasing(ce *types.ClassEntry, aliases **types.ClassEntry) {
	var i int = 0
	var cur_alias *ZendTraitAlias
	var lc_method_name *types.String
	if ce.GetTraitAliases() != nil {
		for ce.GetTraitAliases()[i] != nil {
			cur_alias = ce.GetTraitAliases()[i]

			/** The trait for this alias has not been resolved, this means, this
			  alias was not applied. Abort with an error. */

			if aliases[i] == nil {
				if cur_alias.GetAlias() != nil {

					/** Plain old inconsistency/typo/bug */

					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "An alias (%s) was defined for method %s(), but this method does not exist", cur_alias.GetAlias().GetVal(), cur_alias.GetTraitMethod().GetMethodName().GetVal())

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
					if ce.FunctionTable().Exists(lc_method_name.GetStr()) {
						// types.ZendStringReleaseEx(lc_method_name, 0)
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The modifiers for the trait alias %s() need to be changed in the same statement in which the alias is defined. Error", cur_alias.GetTraitMethod().GetMethodName().GetVal())
					} else {
						// types.ZendStringReleaseEx(lc_method_name, 0)
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The modifiers of the trait method %s() are changed, but this method does not exist. Error", cur_alias.GetTraitMethod().GetMethodName().GetVal())
					}
				}
			}
			i++
		}
	}
}
func ZendDoBindTraits(ce *types.ClassEntry) {
	var exclude_tables **types.Array
	var aliases **types.ClassEntry
	var traits **types.ClassEntry
	var trait **types.ClassEntry
	var i uint32
	var j uint32
	b.Assert(ce.GetNumTraits() > 0)
	traits = Emalloc(b.SizeOf("zend_class_entry *") * ce.GetNumTraits())
	for i = 0; i < ce.GetNumTraits(); i++ {
		trait = ZendFetchClassByName(ce.GetTraitNames()[i].GetName(), ce.GetTraitNames()[i].GetLcName(), ZEND_FETCH_CLASS_TRAIT)
		if trait == nil {
			return
		}
		if !trait.IsTrait() {
			faults.ErrorNoreturn(faults.E_ERROR, "%s cannot use %s - it is not a trait", ce.GetName().GetVal(), trait.GetName().GetVal())
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
func ZendHasDeprecatedConstructor(ce *types.ClassEntry) types.ZendBool {
	var constructor_name *types.String
	if ce.GetConstructor() == nil {
		return 0
	}
	constructor_name = ce.GetConstructor().GetFunctionName()
	return !(ZendBinaryStrcasecmp(ce.GetName().GetStr(), constructor_name.GetStr()))
}
func ZendCheckDeprecatedConstructor(ce *types.ClassEntry) {
	if ZendHasDeprecatedConstructor(ce) != 0 {
		faults.Error(faults.E_DEPRECATED, "Methods with the same name as their class will not be constructors in a future version of PHP; %s has a deprecated constructor", ce.GetName().GetVal())
	}
}
func DISPLAY_ABSTRACT_FN(idx int) {
	b.CondF1(ai.afn[idx], func() string { return ZEND_FN_SCOPE_NAME(ai.afn[idx]) }, "")
	b.Cond(ai.afn[idx], "::", "")
	b.CondF1(ai.afn[idx], func() []byte { return ai.afn[idx].common.function_name.GetVal() }, "")
	b.CondF2(ai.afn[idx] && ai.afn[idx+1], ", ", func() string {
		if ai.afn[idx] && ai.cnt > MAX_ABSTRACT_INFO_CNT {
			return ", ..."
		} else {
			return ""
		}
	})
}
func ZendVerifyAbstractClassFunction(fn types.IFunction, ai *ZendAbstractInfo) {
	if fn.IsAbstract() {
		if ai.GetCnt() < MAX_ABSTRACT_INFO_CNT {
			ai.GetAfn()[ai.GetCnt()] = fn
		}
		if fn.IsCtor() {
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
func ZendVerifyAbstractClass(ce *types.ClassEntry) {
	var ai ZendAbstractInfo
	b.Assert((ce.GetCeFlags() & (AccImplicitAbstractClass | AccInterface | AccTrait | AccExplicitAbstractClass)) == AccImplicitAbstractClass)
	memset(&ai, 0, b.SizeOf("ai"))
	ce.FunctionTable().Foreach(func(_ string, func_ types.IFunction) {
		ZendVerifyAbstractClassFunction(func_, &ai)
	})
	if ai.GetCnt() != 0 {
		faults.ErrorNoreturn(faults.E_ERROR, "Class %s contains %d abstract method%s and must therefore be declared abstract or implement the remaining methods ("+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+")", ce.GetName().GetVal(), ai.GetCnt(), b.Cond(ai.GetCnt() > 1, "s", ""), DISPLAY_ABSTRACT_FN(0), DISPLAY_ABSTRACT_FN(1), DISPLAY_ABSTRACT_FN(2))
	} else {

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

		ce.SetIsImplicitAbstractClass(false)

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

	}
}
func VarianceObligationDtor(zv *types.Zval) { Efree(zv.GetPtr()) }
func VarianceObligationHtDtor(zv *types.Zval) {
	zv.GetPtr().Destroy()
	FREE_HASHTABLE(zv.GetPtr())
}
func GetOrInitObligationsForClass(ce *types.ClassEntry) *types.Array {
	var ht *types.Array
	var key ZendUlong
	if CG__().GetDelayedVarianceObligations() == nil {
		ALLOC_HASHTABLE(CG__().GetDelayedVarianceObligations())
		CG__().GetDelayedVarianceObligations() = types.MakeArrayEx(0, VarianceObligationHtDtor, 0)
	}
	key = ZendUlong(uintPtr(ce))
	ht = types.ZendHashIndexFindPtr(CG__().GetDelayedVarianceObligations(), key)
	if ht != nil {
		return ht
	}
	ALLOC_HASHTABLE(ht)
	ht = types.MakeArrayEx(0, VarianceObligationDtor, 0)
	types.ZendHashIndexAddNewPtr(CG__().GetDelayedVarianceObligations(), key, ht)
	ce.SetIsUnresolvedVariance(true)
	return ht
}
func AddDependencyObligation(ce *types.ClassEntry, dependency_ce *types.ClassEntry) {
	var obligations *types.Array = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_DEPENDENCY)
	obligation.SetDependencyCe(dependency_ce)
	types.ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddCompatibilityObligation(ce *types.ClassEntry, child_fn types.IFunction, parent_fn types.IFunction, always_error types.ZendBool) {
	var obligations *types.Array = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_COMPATIBILITY)

	/* Copy functions, because they may be stack-allocated in the case of traits. */

	if child_fn.GetType() == ZEND_INTERNAL_FUNCTION {
		memcpy(obligation.GetChildFn(), child_fn, b.SizeOf("zend_internal_function"))
	} else {
		memcpy(obligation.GetChildFn(), child_fn, b.SizeOf("zend_op_array"))
	}
	if parent_fn.GetType() == ZEND_INTERNAL_FUNCTION {
		memcpy(obligation.GetParentFn(), parent_fn, b.SizeOf("zend_internal_function"))
	} else {
		memcpy(obligation.GetParentFn(), parent_fn, b.SizeOf("zend_op_array"))
	}
	obligation.SetAlwaysError(always_error)
	types.ZendHashNextIndexInsertPtr(obligations, obligation)
}
func AddPropertyCompatibilityObligation(ce *types.ClassEntry, child_prop *ZendPropertyInfo, parent_prop *ZendPropertyInfo) {
	var obligations *types.Array = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_PROPERTY_COMPATIBILITY)
	obligation.SetChildProp(child_prop)
	obligation.SetParentProp(parent_prop)
	types.ZendHashNextIndexInsertPtr(obligations, obligation)
}
func CheckVarianceObligation(zv *types.Zval) int {
	var obligation *VarianceObligation = zv.GetPtr()
	if obligation.GetType() == OBLIGATION_DEPENDENCY {
		var dependency_ce *types.ClassEntry = obligation.GetDependencyCe()
		if dependency_ce.IsUnresolvedVariance() {
			ResolveDelayedVarianceObligations(dependency_ce)
		}
		if !dependency_ce.IsLinked() {
			return types.ArrayApplyKeep
		}
	} else if obligation.GetType() == OBLIGATION_COMPATIBILITY {
		var unresolved_class *types.String
		var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, obligation.GetChildFn(), obligation.GetParentFn())
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return types.ArrayApplyKeep
			}
			b.Assert(status == INHERITANCE_ERROR)
			EmitIncompatibleMethodErrorOrWarning(obligation.GetChildFn(), obligation.GetParentFn(), status, unresolved_class, obligation.GetAlwaysError())
		}
	} else {
		b.Assert(obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY)
		var status InheritanceStatus = PropertyTypesCompatible(obligation.GetParentProp(), obligation.GetChildProp())
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return types.ArrayApplyKeep
			}
			b.Assert(status == INHERITANCE_ERROR)
			EmitIncompatiblePropertyError(obligation.GetChildProp(), obligation.GetParentProp())
		}
	}
	return types.ArrayApplyRemove
}
func LoadDelayedClasses() {
	var delayed_autoloads *types.Array = CG__().GetDelayedAutoloads()
	var name *types.String
	if delayed_autoloads == nil {
		return
	}

	/* Take ownership of this HT, to avoid concurrent modification during autoloading. */

	CG__().SetDelayedAutoloads(nil)
	var __ht *types.Array = delayed_autoloads
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		name = _p.GetKey()
		ZendLookupClass(name)
	}
	delayed_autoloads.Destroy()
	FREE_HASHTABLE(delayed_autoloads)
}
func ResolveDelayedVarianceObligations(ce *types.ClassEntry) {
	var all_obligations *types.Array = CG__().GetDelayedVarianceObligations()
	var obligations *types.Array
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	b.Assert(all_obligations != nil)
	obligations = types.ZendHashIndexFindPtr(all_obligations, num_key)
	b.Assert(obligations != nil)
	types.ZendHashApply(obligations, CheckVarianceObligation)
	if obligations.Len() == 0 {
		ce.SetIsUnresolvedVariance(false)
		ce.SetIsLinked(true)
		types.ZendHashIndexDel(all_obligations, num_key)
	}
}
func ReportVarianceErrors(ce *types.ClassEntry) {
	var all_obligations *types.Array = CG__().GetDelayedVarianceObligations()
	var obligations *types.Array
	var obligation *VarianceObligation
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	b.Assert(all_obligations != nil)
	obligations = types.ZendHashIndexFindPtr(all_obligations, num_key)
	b.Assert(obligations != nil)
	var __ht *types.Array = obligations
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		obligation = _z.GetPtr()
		var status InheritanceStatus
		var unresolved_class *types.String
		if obligation.GetType() == OBLIGATION_COMPATIBILITY {

			/* Just used to fetch the unresolved_class in this case. */

			status = ZendDoPerformImplementationCheck(&unresolved_class, obligation.GetChildFn(), obligation.GetParentFn())
			b.Assert(status == INHERITANCE_UNRESOLVED)
			EmitIncompatibleMethodErrorOrWarning(obligation.GetChildFn(), obligation.GetParentFn(), status, unresolved_class, obligation.GetAlwaysError())
		} else if obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY {
			EmitIncompatiblePropertyError(obligation.GetChildProp(), obligation.GetParentProp())
		} else {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Bug #78647")
		}
	}

	/* Only warnings were thrown above -- that means that there are incompatibilities, but only
	 * ones that we permit. Mark all classes with open obligations as fully linked. */

	ce.SetIsUnresolvedVariance(false)
	ce.SetIsLinked(true)
	types.ZendHashIndexDel(all_obligations, num_key)
}
func CheckUnrecoverableLoadFailure(ce *types.ClassEntry) {
	/* If this class has been used while unlinked through a variance obligation, it is not legal
	 * to remove the class from the class table and throw an exception, because there is already
	 * a dependence on the inheritance hierarchy of this specific class. Instead we fall back to
	 * a fatal error, as would happen if we did not allow exceptions in the first place. */

	if ce.IsHasUnlinkedUses() {
		var exception_str *types.String
		var exception_zv types.Zval
		b.Assert(EG__().GetException() != nil && "Exception must have been thrown")
		exception_zv.SetObject(EG__().GetException())
		exception_zv.AddRefcount()
		faults.ClearException()
		exception_str = ZvalGetString(&exception_zv)
		faults.ErrorNoreturn(faults.E_ERROR, "During inheritance of %s with variance dependencies: Uncaught %s", ce.GetName().GetVal(), exception_str.GetVal())
	}

	/* If this class has been used while unlinked through a variance obligation, it is not legal
	 * to remove the class from the class table and throw an exception, because there is already
	 * a dependence on the inheritance hierarchy of this specific class. Instead we fall back to
	 * a fatal error, as would happen if we did not allow exceptions in the first place. */
}
func ZendDoLinkClass(ce *types.ClassEntry, lc_parent_name *types.String) int {
	/* Load parent/interface dependencies first, so we can still gracefully abort linking
	 * with an exception and remove the class from the class table. This is only possible
	 * if no variance obligations on the current class have been added during autoloading. */

	var parent *types.ClassEntry = nil
	var interfaces **types.ClassEntry = nil
	if ce.GetParentName() {
		parent = ZendFetchClassByName(ce.GetParentName(), lc_parent_name, ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED|ZEND_FETCH_CLASS_EXCEPTION)
		if parent == nil {
			CheckUnrecoverableLoadFailure(ce)
			return types.FAILURE
		}
	}
	if ce.GetNumInterfaces() != 0 {

		/* Also copy the parent interfaces here, so we don't need to reallocate later. */

		var i uint32
		var num_parent_interfaces uint32 = b.CondF1(parent != nil, func() uint32 { return parent.GetNumInterfaces() }, 0)
		interfaces = Emalloc(b.SizeOf("zend_class_entry *") * (ce.GetNumInterfaces() + num_parent_interfaces))
		if num_parent_interfaces != 0 {
			memcpy(interfaces, parent.GetInterfaces(), b.SizeOf("zend_class_entry *")*num_parent_interfaces)
		}
		for i = 0; i < ce.GetNumInterfaces(); i++ {
			var iface *types.ClassEntry = ZendFetchClassByName(ce.GetInterfaceNames()[i].name, ce.GetInterfaceNames()[i].lcName, ZEND_FETCH_CLASS_INTERFACE|ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED|ZEND_FETCH_CLASS_EXCEPTION)
			if iface == nil {
				CheckUnrecoverableLoadFailure(ce)
				Efree(interfaces)
				return types.FAILURE
			}
			interfaces[num_parent_interfaces+i] = iface
		}
	}
	if parent != nil {
		if !parent.IsLinked() {
			AddDependencyObligation(ce, parent)
		}
		ZendDoInheritance(ce, parent)
	}
	if ce.IsImplementTraits() {
		ZendDoBindTraits(ce)
	}
	if ce.IsImplementInterfaces() {
		ZendDoImplementInterfaces(ce, interfaces)
	}
	if (ce.GetCeFlags() & (AccImplicitAbstractClass | AccInterface | AccTrait | AccExplicitAbstractClass)) == AccImplicitAbstractClass {
		ZendVerifyAbstractClass(ce)
	}
	ZendBuildPropertiesInfoTable(ce)
	if !ce.IsUnresolvedVariance() {
		ce.SetIsLinked(true)
		return types.SUCCESS
	}
	ce.SetIsNearlyLinked(true)
	LoadDelayedClasses()
	if ce.IsUnresolvedVariance() {
		ResolveDelayedVarianceObligations(ce)
		if !ce.IsLinked() {
			ReportVarianceErrors(ce)
		}
	}
	return types.SUCCESS
}
func ZendCanEarlyBind(ce *types.ClassEntry, parent_ce *types.ClassEntry) InheritanceStatus {
	var ret InheritanceStatus = INHERITANCE_SUCCESS
	var key *types.String
	var parent_info *ZendPropertyInfo

	parent_ce.FunctionTable().ForeachEx(func(key string, parent_func types.IFunction) bool {
		var child_func types.IFunction = ce.FunctionTable().Get(key)
		if child_func != nil {
			var status InheritanceStatus = DoInheritanceCheckOnMethodEx(child_func, parent_func, ce, nil, 1, 0)
			if status != INHERITANCE_SUCCESS {
				b.Assert(status == INHERITANCE_UNRESOLVED || status == INHERITANCE_ERROR)
				ret = status
				return false
			}
		}
		return true
	})
	if ret == INHERITANCE_UNRESOLVED {
		return ret
	}

	var __ht__1 *types.Array = parent_ce.GetPropertiesInfo()
	for _, _p := range __ht__1.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		key = _p.GetKey()
		parent_info = _z.GetPtr()
		var zv *types.Zval
		if parent_info.IsPrivate() || !(parent_info.GetType().IsSet()) {
			continue
		}
		zv = ce.GetPropertiesInfo().KeyFind(key.GetStr())
		if zv != nil {
			var child_info *ZendPropertyInfo = zv.GetPtr()
			if child_info.GetType().IsSet() {
				var status InheritanceStatus = PropertyTypesCompatible(parent_info, child_info)
				if status != INHERITANCE_SUCCESS {
					if status == INHERITANCE_UNRESOLVED {
						return INHERITANCE_UNRESOLVED
					}
					b.Assert(status == INHERITANCE_ERROR)
					ret = INHERITANCE_ERROR
				}
			}
		}
	}
	return ret
}
func ZendTryEarlyBind(ce *types.ClassEntry, parent_ce *types.ClassEntry, lcname *types.String) bool {
	var status InheritanceStatus = ZendCanEarlyBind(ce, parent_ce)
	if status != INHERITANCE_UNRESOLVED {
		if !CG__().ClassTable().Add(lcname.GetStr(), ce) {
			return false
		}
		ZendDoInheritanceEx(ce, parent_ce, status == INHERITANCE_SUCCESS)
		ZendBuildPropertiesInfoTable(ce)
		if (ce.GetCeFlags() & (AccImplicitAbstractClass | AccInterface | AccTrait | AccExplicitAbstractClass)) == AccImplicitAbstractClass {
			ZendVerifyAbstractClass(ce)
		}
		b.Assert(!ce.IsUnresolvedVariance())
		ce.SetIsLinked(true)
		return true
	}
	return false
}
