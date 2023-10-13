package zend

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendDoInheritance(ce *types.ClassEntry, parentCe *types.ClassEntry) {
	ZendDoInheritanceEx(ce, parentCe, false)
}
func ZendDuplicatePropertyInfoInternal(property_info *types.PropertyInfo) *types.PropertyInfo {
	var new_property_info *types.PropertyInfo = Pemalloc(b.SizeOf("zend_property_info"))
	memcpy(new_property_info, property_info, b.SizeOf("zend_property_info"))
	return new_property_info
}
func ZendDuplicateInternalFunction(func_ *types.InternalFunction, ce *types.ClassEntry) *types.InternalFunction {
	var newFunction *types.InternalFunction
	if ce.IsInternalClass() {
		newFunction = types.CopyInternalFunction(func_)
	} else {
		newFunction = types.CopyInternalFunction(func_)
		newFunction.SetIsArenaAllocated(true)
	}
	return newFunction
}
func ZendDuplicateUserFunction(func_ *types.UserFunction) *types.UserFunction {
	var newFunction = types.CopyOpArray(func_)
	if func_.GetOpArray().GetStaticVariablesPtr() != nil {
		/* See: Zend/tests/method_static_var.phpt */
		newFunction.GetOpArray().SetStaticVariables(func_.GetOpArray().GetStaticVariablesPtr())
	}
	if CG__().IsCompilePreload() {
		b.Assert(newFunction.GetOpArray().IsPreloaded())
		ZEND_MAP_PTR_NEW(newFunction.GetOpArray().static_variables_ptr)
	} else {
		ZEND_MAP_PTR_INIT(newFunction.GetOpArray().static_variables_ptr, newFunction.GetOpArray().GetStaticVariables())
	}
	return newFunction
}
func ZendDuplicateFunction(func_ types.IFunction, ce *types.ClassEntry, isInterface bool) types.IFunction {
	if func_.GetType() == ZEND_INTERNAL_FUNCTION {
		return ZendDuplicateInternalFunction(func_.GetInternalFunction(), ce)
	} else {
		func_.GetOpArray().TryIncRefCount()
		if isInterface || func_.GetOpArray().GetStaticVariables() == nil {
			/* reuse the same op_array structure */
			return func_
		}
		return ZendDuplicateUserFunction(func_.GetOpArray())
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
			faults.ErrorNoreturn(faults.E_ERROR, "Cannot override final %s::%s() with %s::%s()", parent.Name(), parent.GetConstructor().FunctionName(), ce.Name(), ce.GetConstructor().FunctionName())
		}
		return
	}
	ce.SetConstructor(parent.GetConstructor())
}
func ZendVisibilityString(fn_flags uint32) string {
	if (fn_flags & types.AccPublic) != 0 {
		return "public"
	} else if (fn_flags & types.AccPrivate) != 0 {
		return "private"
	} else {
		b.Assert((fn_flags & types.AccProtected) != 0)
		return "protected"
	}
}
func ResolveClassName(scope *types.ClassEntry, name string) string {
	b.Assert(scope != nil)

	lcName := ascii.StrToLower(name)
	if lcName == "parent" && scope.GetParent() != nil {
		if scope.IsResolvedParent() {
			return scope.GetParent().Name()
		} else {
			return scope.GetParentName().GetStr()
		}
	} else if lcName == "self" {
		return scope.Name()
	} else {
		return name
	}
}
func ClassVisible(ce *types.ClassEntry) bool {
	if ce.IsInternalClass() {
		return (CG__().GetCompilerOptions() & ZEND_COMPILE_IGNORE_INTERNAL_CLASSES) != 0
	} else {
		b.Assert(ce.IsUserClass())
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
			CG__().SetDelayedAutoloads(types.NewArray())
		}
		types.ZendHashAddEmptyElement(CG__().GetDelayedAutoloads(), name.GetStr())
	} else {
		ce = ZendLookupClassEx(name, nil, ZEND_FETCH_CLASS_NO_AUTOLOAD)
		if ce != nil && ClassVisible(ce) {
			return ce
		}

		/* The current class may not be registered yet, so check for it explicitly. */

		if ascii.StrCaseEquals(scope.Name(), name.GetStr()) {
			return scope
		}

		/* The current class may not be registered yet, so check for it explicitly. */

	}
	return nil
}
func UnlinkedInstanceof(ce1 *types.ClassEntry, ce2 *types.ClassEntry) bool {
	if ce1 == ce2 {
		return 1
	}
	if ce1.IsLinked() {
		return operators.InstanceofFunction(ce1, ce2)
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
				var ce *types.ClassEntry = ZendLookupClassEx_Ex(ce1.GetInterfaceNames()[i].GetName(), ce1.GetInterfaceNames()[i].GetLcName(), ZEND_FETCH_CLASS_ALLOW_UNLINKED|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if ce != nil && UnlinkedInstanceof(ce, ce2) != 0 {
					return 1
				}
			}
		}
	}
	return 0
}
func ZendPerformCovariantTypeCheck(unresolved_class **types.String, fe types.IFunction, fe_arg_info *ZendArgInfo, proto types.IFunction, proto_arg_info *ZendArgInfo) InheritanceStatus {
	var fe_type types.TypeHint = fe_arg_info.GetType()
	var proto_type types.TypeHint = proto_arg_info.GetType()
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
	} else if proto_type.Code() == types.IsIterable {
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
		if fe_type.Code() == types.IsIterable || fe_type.Code() == types.IsArray {
			return INHERITANCE_SUCCESS
		} else {
			return INHERITANCE_ERROR
		}
	} else if proto_type.Code() == types.IsObject {
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
		if fe_type.Code() == types.IsObject {
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
		str.WriteByte('?')
	}
	if arg_info.GetType().IsClass() {
		className := arg_info.GetType().Name()
		lcClassName := ascii.StrToLower(className)

		if lcClassName == "self" && fptr.GetScope() != nil {
			className = fptr.GetScope().Name()
		} else if lcClassName == "parent" && fptr.GetScope() != nil && fptr.GetScope().GetParent() != nil {
			className = fptr.GetScope().GetParent().Name()
		}
		str.WriteString(className)
		if return_hint == 0 {
			str.WriteByte(' ')
		}
	} else if arg_info.GetType().IsCode() {
		var typeName = types.ZendGetTypeByConst(arg_info.GetType().Code())
		str.WriteString(typeName)
		if return_hint == 0 {
			str.WriteByte(' ')
		}
	}
}
func ZendGetFunctionDeclaration(fptr types.IFunction) string {
	var str SmartStr
	if fptr.GetOpArray().IsReturnReference() {
		str.WriteString("& ")
	}
	if fptr.GetScope() != nil {

		/* cut off on NULL byte ... class@anonymous */

		str.WriteString(b.CastStr(fptr.GetScope().Name(), strlen(fptr.GetScope().Name())))
		str.WriteString("::")
	}
	str.WriteString(fptr.FunctionName())
	str.WriteByte('(')
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
				str.WriteByte('&')
			}
			if arg_info.GetIsVariadic() != 0 {
				str.WriteString("...")
			}
			str.WriteByte('$')
			if arg_info.GetName() != nil {
				if fptr.GetType() == ZEND_INTERNAL_FUNCTION {
					str.WriteString((*ArgInfo)(arg_info).Name())
				} else {
					str.WriteString(arg_info.GetName().GetStr())
				}
			} else {
				str.WriteString("param")
				str.WriteUlong(i)
			}
			if i >= required && arg_info.GetIsVariadic() == 0 {
				str.WriteString(" = ")
				if fptr.GetType() == ZEND_USER_FUNCTION {
					var precv *types.ZendOp = nil
					var idx uint32 = i
					var op *types.ZendOp = fptr.GetOpArray().GetOpcodes()
					var end *types.ZendOp = op + fptr.GetOpArray().GetLast()
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
							str.WriteString("false")
						} else if zv.IsTrue() {
							str.WriteString("true")
						} else if zv.IsNull() {
							str.WriteString("NULL")
						} else if zv.IsString() {
							str.WriteByte('\'')
							str.WriteString(b.CastStr(zv.StringEx().GetVal(), b.Min(zv.StringEx().GetLen(), 10)))
							if zv.StringEx().GetLen() > 10 {
								str.WriteString("...")
							}
							str.WriteByte('\'')
						} else if zv.IsArray() {
							str.WriteString("Array")
						} else if zv.IsConstantAst() {
							var ast *ZendAst = types.Z_ASTVAL_P(zv)
							if ast.Kind() == ZEND_AST_CONSTANT {
								str.WriteString(ZendAstGetConstantName(ast).GetStr())
							} else {
								str.WriteString("<expression>")
							}
						} else {
							var zv_str *types.String = operators.ZvalGetString(zv)
							str.WriteString(zv_str.GetStr())
							//ZendTmpStringRelease(tmp_zv_str)
						}
					}
				} else {
					str.WriteString("NULL")
				}
			}
			if lang.PreInc(&i) < num_args {
				str.WriteString(", ")
			}
			arg_info++
		}
	}
	str.WriteByte(')')
	if fptr.IsHasReturnType() {
		str.WriteString(": ")
		ZendAppendTypeHint(&str, fptr, fptr.GetArgInfo()-1, 1)
	}
	str.ZeroTail()
	return str.GetStr()
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
	var parentPrototype = ZendGetFunctionDeclaration(parent)
	var childPrototype = ZendGetFunctionDeclaration(child)
	var errorMsg string
	if status == INHERITANCE_UNRESOLVED {
		errorMsg = fmt.Sprintf("Could not check compatibility between %s and %s, because class %s is not available", childPrototype, parentPrototype, unresolved_class.GetVal())
	} else {
		errorMsg = fmt.Sprintf("Declaration of %s %s be compatible with %s", childPrototype, error_verb, parentPrototype)
	}
	faults.ErrorAt(error_level, nil, FuncLineno(child), errorMsg)
}
func EmitIncompatibleMethodErrorOrWarning(child types.IFunction, parent types.IFunction, status InheritanceStatus, unresolved_class *types.String, always_error bool) {
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
func PerformDelayableImplementationCheck(ce *types.ClassEntry, fe types.IFunction, proto types.IFunction, always_error bool) {
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
	check_only bool,
	checked bool,
) InheritanceStatus {
	var child_flags uint32
	var parent_flags uint32 = parent.GetFnFlags()
	var proto types.IFunction
	if checked == 0 && (parent_flags&types.AccFinal) != 0 {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot override final method %s::%s()", ZEND_FN_SCOPE_NAME(parent), child.FunctionName())
	}
	child_flags = child.GetFnFlags()

	/* You cannot change from static to non static and vice versa.
	 */

	if checked == 0 && (child_flags&types.AccStatic) != (parent_flags&types.AccStatic) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		if (child_flags & types.AccStatic) != 0 {
			faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non static method %s::%s() static in class %s", ZEND_FN_SCOPE_NAME(parent), child.FunctionName(), ZEND_FN_SCOPE_NAME(child))
		} else {
			faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make static method %s::%s() non static in class %s", ZEND_FN_SCOPE_NAME(parent), child.FunctionName(), ZEND_FN_SCOPE_NAME(child))
		}
	}

	/* Disallow making an inherited method abstract. */

	if checked == 0 && (child_flags&types.AccAbstract) > (parent_flags&types.AccAbstract) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Cannot make non abstract method %s::%s() abstract in class %s", ZEND_FN_SCOPE_NAME(parent), child.FunctionName(), ZEND_FN_SCOPE_NAME(child))
	}
	if check_only == 0 && (parent_flags&(types.AccPrivate|types.AccChanged)) != 0 {
		child.SetIsChanged(true)
	}
	if (parent_flags & types.AccPrivate) != 0 {
		return INHERITANCE_SUCCESS
	}
	if parent.GetPrototype() != nil {
		proto = parent.GetPrototype()
	} else {
		proto = parent
	}
	if (parent_flags & types.AccCtor) != 0 {

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

	if checked == 0 && (child_flags&types.AccPppMask) > (parent_flags&types.AccPppMask) {
		if check_only != 0 {
			return INHERITANCE_ERROR
		}
		faults.ErrorAtNoreturn(faults.E_COMPILE_ERROR, nil, FuncLineno(child), "Access level to %s::%s() must be %s (as in class %s)%s", ZEND_FN_SCOPE_NAME(child), child.FunctionName(), ZendVisibilityString(parent_flags), ZEND_FN_SCOPE_NAME(parent), lang.Cond((parent_flags&types.AccPublic) != 0, "", " or weaker"))
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
func DoInheritMethod(key string, parent types.IFunction, ce *types.ClassEntry, is_interface bool, checked bool) {
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
		//	types.ZendHashAddNewPtr(ce.GetFunctionTable(), key.String(), parent)
		//}
	}
}
func PropertyTypesCompatible(parent_info *types.PropertyInfo, child_info *types.PropertyInfo) InheritanceStatus {
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
		parent_name = parent_info.GetType().Ce().GetName()
	} else {
		parent_name = ResolveClassName(parent_info.GetCe(), parent_info.GetType().Name())
	}
	if child_info.GetType().IsCe() {
		child_name = child_info.GetType().Name()
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
func EmitIncompatiblePropertyError(child *types.PropertyInfo, parent *types.PropertyInfo) {
	var typ string
	if parent.GetType().IsClass() {
		if parent.GetType().IsCe() {
			typ = parent.GetType().Name()
		} else {
			typ = ResolveClassName(parent.GetCe(), parent.GetType().Name())
		}
	} else {
		typ = types.ZendGetTypeByConst(parent.GetType().Code())
	}
	if parent.GetType().AllowNull() {
		typ = "?" + typ
	}

	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type of %s::$%s must be %s (as in class %s)", child.GetCe().Name(), ZendGetUnmangledPropertyNameEx(child.GetName()), typ, parent.GetCe().Name())
}
func DoInheritProperty(parent_info *types.PropertyInfo, key string, ce *types.ClassEntry) {
	var child_info *types.PropertyInfo = ce.PropertyTable().Get(key)
	if child_info != nil {
		if parent_info.HasFlags(types.AccPrivate | types.AccChanged) {
			child_info.MarkIsChanged()
		}
		if !parent_info.IsPrivate() {
			if (parent_info.GetFlags() & types.AccStatic) != (child_info.GetFlags() & types.AccStatic) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot redeclare %s%s::$%s as %s%s::$%s", lang.Cond(parent_info.IsStatic(), "static ", "non static "), ce.GetParent().Name(), key, lang.Cond(child_info.IsStatic(), "static ", "non static "), ce.Name(), key)
			}
			if (child_info.GetFlags() & types.AccPppMask) > (parent_info.GetFlags() & types.AccPppMask) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access level to %s::$%s must be %s (as in class %s)%s", ce.Name(), key, ZendVisibilityString(parent_info.GetFlags()), ce.GetParent().Name(), lang.Cond(parent_info.IsPublic(), "", " or weaker"))
			} else if !child_info.IsStatic() {
				var parent_num int = OBJ_PROP_TO_NUM(parent_info.GetOffset())
				var child_num int = OBJ_PROP_TO_NUM(child_info.GetOffset())

				/* Don't keep default properties in GC (they may be freed by opcache) */

				// ZvalPtrDtorNogc(&ce.GetDefaultPropertiesTable()[parent_num])
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
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Type of %s::$%s must not be defined (as in class %s)", ce.Name(), key, ce.GetParent().Name())
			}
		}
	} else {
		if ce.IsInternalClass() {
			child_info = ZendDuplicatePropertyInfoInternal(parent_info)
		} else {
			child_info = parent_info
		}
		ce.PropertyTable().Add(key, child_info)
	}
}
func DoImplementInterface(ce *types.ClassEntry, iface *types.ClassEntry) {
	if !ce.IsInterface() && iface.GetInterfaceGetsImplemented() && iface.GetInterfaceGetsImplemented()(iface, ce) == types.FAILURE {
		faults.ErrorNoreturn(faults.E_CORE_ERROR, "Class %s could not implement interface %s", ce.Name(), iface.Name())
	}

	/* This should be prevented by the class lookup logic. */
	b.Assert(ce != iface)
}
func ZendDoInheritInterfaces(ce *types.ClassEntry, iface *types.ClassEntry) {
	/* expects interface to be contained in ce's interface list already */
	/* Inherit the interfaces, only if they're not already inherited by the class */
	interfaces := ce.GetInterfaces()
	rawNum := len(interfaces)
	for _, newInterface := range iface.GetInterfaces() {
		i := 0
		for ; i < rawNum; i++ {
			if interfaces[i] == newInterface {
				break
			}
		}
		if i == rawNum {
			interfaces = append(interfaces, newInterface)
		}
	}
	ce.ResolvedInterfaces(interfaces)

	/* and now call the implementing handlers */
	for _, newInterface := range interfaces[rawNum:] {
		DoImplementInterface(ce, newInterface)
	}
}
func DoInheritClassConstant(name string, parentConst *types.ClassConstant, ce *types.ClassEntry) {
	var c = ce.ConstantsTable().Get(name)
	if c != nil {
		if c.PriorLevel() > parentConst.PriorLevel() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Access level to %s::%s must be %s (as in class %s)%s", ce.Name(), name, ZendVisibilityString(parentConst.GetAccessFlags()), ce.GetParent().Name(), lang.Cond(parentConst.IsPublic(), "", " or weaker"))
		}
	} else if !parentConst.IsPrivate() {
		if parentConst.GetValue().IsConstantAst() {
			ce.SetIsConstantsUpdated(false)
		}
		if ce.IsInternalClass() {
			parentConst = types.CopyClassConstant(parentConst)
		}
		ce.ConstantsTable().Add(name, parentConst)
	}
}
func ZendBuildPropertiesInfoTable(ce *types.ClassEntry) {
	if ce.GetDefaultPropertiesCount() == 0 {
		return
	}
	b.Assert(ce.GetPropertiesInfoTable() == nil)

	var table []*types.PropertyInfo = make([]*types.PropertyInfo, ce.GetDefaultPropertiesCount())
	ce.SetPropertiesInfoTable(table)

	/* Dead slots may be left behind during inheritance. Make sure these are NULLed out. */
	if ce.GetParent() != nil && ce.GetParent().GetDefaultPropertiesCount() != 0 {
		var parentTable []*types.PropertyInfo = ce.GetParent().GetPropertiesInfoTable()
		copy(table, parentTable[:ce.GetParent().GetDefaultPropertiesCount()])

		/* Child did not add any new properties, we are done */
		if ce.GetDefaultPropertiesCount() == ce.GetParent().GetDefaultPropertiesCount() {
			return
		}
	}

	ce.PropertyTable().Foreach(func(key string, propInfo *types.PropertyInfo) {
		if propInfo.GetCe() == ce && !propInfo.IsStatic() {
			table[OBJ_PROP_TO_NUM(propInfo.GetOffset())] = propInfo
		}
	})
}
func ZendDoInheritanceEx(ce *types.ClassEntry, parentCe *types.ClassEntry, checked bool) {
	if ce.IsInterface() {
		/* Interface can only inherit other interfaces */
		if !parentCe.IsInterface() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Interface %s may not inherit from class (%s)", ce.Name(), parentCe.Name())
		}
	} else if parentCe.HasCeFlags(types.AccInterface | types.AccTrait | types.AccFinal) {
		/* Class declaration must not extend traits or interfaces */
		if parentCe.IsInterface() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot extend from interface %s", ce.Name(), parentCe.Name())
		} else if parentCe.IsTrait() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot extend from trait %s", ce.Name(), parentCe.Name())
		}

		/* Class must not extend a final class */
		if parentCe.IsFinal() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s may not inherit from final class (%s)", ce.Name(), parentCe.Name())
		}
	}

	ce.SetParent(parentCe)
	ce.SetIsResolvedParent(true)

	/* Inherit interfaces */
	if parentCe.GetNumInterfaces() != 0 {
		if !ce.IsImplementInterfaces() {
			ZendDoInheritInterfaces(ce, parentCe)
		} else {
			for _, iface := range parentCe.GetInterfaces() {
				DoImplementInterface(ce, iface)
			}
		}
	}

	/* Inherit properties */
	if parentCe.GetDefaultPropertiesCount() != 0 {
		parentTable := parentCe.GetDefaultPropertiesTable()
		oldTable := ce.GetDefaultPropertiesTable()
		newTable := make([]types.Zval, len(parentTable)+len(oldTable))
		// 复制 parent 默认属性表
		extendsInternal := parentCe.GetType() == ce.GetType() /* User class extends internal */
		for i := range parentTable {
			src := &parentTable[i]
			dst := &newTable[i]
			if extendsInternal {
				types.ZVAL_COPY_OR_DUP_PROP(dst, src)
			} else {
				types.ZVAL_COPY_PROP(dst, src)
			}
			if dst.IsConstantAst() {
				ce.SetIsConstantsUpdated(false)
			}
		}
		// 保留 oldTable 默认属性表
		for i := range oldTable {
			src := &oldTable[i]
			dst := &newTable[len(parentTable)+i]
			types.ZVAL_COPY_VALUE_PROP(dst, src)
		}

		ce.SetDefaultPropertiesTableAndCount(newTable)
	}

	if parentCe.GetDefaultStaticMembersCount() != 0 {
		var src *types.Zval
		var dst *types.Zval
		var end *types.Zval
		if ce.GetDefaultStaticMembersCount() != 0 {
			var table *types.Zval = Pemalloc(b.SizeOf("zval") * (ce.GetDefaultStaticMembersCount() + parentCe.GetDefaultStaticMembersCount()))
			src = ce.GetDefaultStaticMembersTable() + ce.GetDefaultStaticMembersCount()
			end = table + parentCe.GetDefaultStaticMembersCount()
			dst = end + ce.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(table)
			for {
				dst--
				src--
				dst.CopyValueFrom(src)
				if dst == end {
					break
				}
			}
			Pefree(src)
			end = ce.GetDefaultStaticMembersTable()
		} else {
			end = Pemalloc(b.SizeOf("zval") * parentCe.GetDefaultStaticMembersCount())
			dst = end + parentCe.GetDefaultStaticMembersCount()
			ce.SetDefaultStaticMembersTable(end)
		}
		if parentCe.GetType() != ce.GetType() {

			/* User class extends internal */

			if CE_STATIC_MEMBERS(parentCe) == nil {
				ZendClassInitStatics(parentCe)
			}
			if ZendUpdateClassConstants(parentCe) != types.SUCCESS {
				b.Assert(false)
			}
			src = CE_STATIC_MEMBERS(parentCe) + parentCe.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.IsIndirect() {
					dst.SetIndirect(src.Indirect())
				} else {
					dst.SetIndirect(src)
				}
				if dst == end {
					break
				}
			}
		} else if ce.IsUserClass() {
			if CE_STATIC_MEMBERS(parentCe) == nil {
				b.Assert(parentCe.HasCeFlags(types.AccImmutable | types.AccPreloaded))
				ZendClassInitStatics(parentCe)
			}
			src = CE_STATIC_MEMBERS(parentCe) + parentCe.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.IsIndirect() {
					dst.SetIndirect(src.Indirect())
				} else {
					dst.SetIndirect(src)
				}
				if types.Z_INDIRECT_P(dst).IsConstantAst() {
					ce.SetIsConstantsUpdated(false)
				}
				if dst == end {
					break
				}
			}
		} else {
			src = parentCe.GetDefaultStaticMembersTable() + parentCe.GetDefaultStaticMembersCount()
			for {
				dst--
				src--
				if src.IsIndirect() {
					dst.SetIndirect(src.Indirect())
				} else {
					dst.SetIndirect(src)
				}
				if dst == end {
					break
				}
			}
		}
		ce.SetDefaultStaticMembersCount(ce.GetDefaultStaticMembersCount() + parentCe.GetDefaultStaticMembersCount())
		if ce.GetStaticMembersTablePtr() == nil {
			b.Assert(ce.IsInternalClass())
			if CurrEX() == nil {
				ZEND_MAP_PTR_NEW(ce.static_members_table)
			} else {
				/* internal class loaded by dl() */
				ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())
			}
		}
	}

	ce.PropertyTable().Foreach(func(key string, propertyInfo *types.PropertyInfo) {
		if propertyInfo.GetCe() == ce {
			if propertyInfo.IsStatic() {
				propertyInfo.SetOffset(propertyInfo.GetOffset() + uint32(parentCe.GetDefaultStaticMembersCount()))
			} else {
				propertyInfo.SetOffset(propertyInfo.GetOffset() + uint32(parentCe.GetDefaultPropertiesCount()*b.SizeOf("zval")))
			}
		}
	})

	if parentCe.PropertyTable().Len() != 0 {
		parentCe.PropertyTable().Foreach(func(key string, property_info *types.PropertyInfo) {
			DoInheritProperty(property_info, key, ce)
		})
	}
	if parentCe.ConstantsTable().Len() != 0 {
		parentCe.ConstantsTable().Foreach(func(key string, c *types.ClassConstant) {
			DoInheritClassConstant(key, c, ce)
		})

	}
	if parentCe.FunctionTable().Len() != 0 {
		if checked {
			parentCe.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
				DoInheritMethod(key, func_, ce, 0, 1)
			})
		} else {
			parentCe.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
				DoInheritMethod(key, func_, ce, 0, 0)
			})
		}
	}
	DoInheritParentConstructor(ce)
	if ce.IsInternalClass() {
		if ce.IsImplicitAbstractClass() {
			ce.SetIsExplicitAbstractClass(true)
		}
	}
	ce.AddCeFlags(parentCe.GetCeFlags() & (types.AccHasStaticInMethods | types.AccHasTypeHints | types.AccUseGuards))
}
func DoInheritConstantCheck(childConstantsTable types.ClassConstantTable, parentConstant *types.ClassConstant, name string, iface *types.ClassEntry) bool {
	var oldConstant *types.ClassConstant = childConstantsTable.Get(name)
	if oldConstant != nil {
		if oldConstant.GetCe() != parentConstant.GetCe() {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot inherit previously-inherited or override constant %s from interface %s", name, iface.Name())
		}
		return 0
	}
	return 1
}
func DoInheritIfaceConstant(name string, c *types.ClassConstant, ce *types.ClassEntry, iface *types.ClassEntry) {
	if DoInheritConstantCheck(ce.ConstantsTable(), c, name, iface) != 0 {
		var ct *types.ClassConstant
		if c.GetValue().IsConstantAst() {
			ce.SetIsConstantsUpdated(false)
		}
		if ce.IsInternalClass() {
			ct = Pemalloc(b.SizeOf("zend_class_constant"))
			memcpy(ct, c, b.SizeOf("zend_class_constant"))
			c = ct
		}
		ce.ConstantsTable().Update(name, c)
	}
}
func DoInterfaceImplementation(ce *types.ClassEntry, iface *types.ClassEntry) {
	iface.ConstantsTable().Foreach(func(key string, c *types.ClassConstant) {
		DoInheritIfaceConstant(key, c, ce, iface)
	})
	iface.FunctionTable().Foreach(func(key string, func_ types.IFunction) {
		DoInheritMethod(key, func_, ce, 1, 0)
	})
	DoImplementInterface(ce, iface)
	if iface.GetNumInterfaces() != 0 {
		ZendDoInheritInterfaces(ce, iface)
	}
}
func ZendDoImplementInterface(ce *types.ClassEntry, iface *types.ClassEntry) {
	b.Assert(ce.IsLinked())

	var ignore uint32 = 0
	var currentIfaceNum = ce.GetNumInterfaces()
	var parentIfaceNum = 0
	if ce.GetParent() != nil {
		parentIfaceNum = ce.GetParent().GetNumInterfaces()
	}
	for i := 0; i < ce.GetNumInterfaces(); i++ {
		if ce.GetInterfaces()[i] == nil {
			memmove(ce.GetInterfaces()+i, ce.GetInterfaces()+i+1, b.SizeOf("zend_class_entry *")*(lang.PreDec(&(ce.GetNumInterfaces()))-i))
			i--
		} else if ce.GetInterfaces()[i] == iface {
			if i < parentIfaceNum {
				ignore = 1
			} else {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ce.Name(), iface.Name())
			}
		}
	}
	if ignore != 0 {
		/* Check for attempt to redeclare interface constants */
		ce.ConstantsTable().Foreach(func(key string, c *types.ClassConstant) {
			DoInheritConstantCheck(iface.ConstantsTable(), c, key, iface)
		})
	} else {
		ce.AppendResolvedInterfaces(iface)
		DoInterfaceImplementation(ce, iface)
	}
}
func ZendDoImplementInterfaces(ce *types.ClassEntry, interfaces []*types.ClassEntry) {
	var numParentInterfaces = 0
	if ce.GetParent() != nil {
		numParentInterfaces = ce.GetParent().GetNumInterfaces()
	}
	for i := 0; i < ce.GetNumInterfaces(); i++ {
		iface := interfaces[numParentInterfaces+i]
		if !iface.IsLinked() {
			AddDependencyObligation(ce, iface)
		}
		if !iface.IsInterface() {
			faults.ErrorNoreturn(faults.E_ERROR, "%s cannot implement %s - it is not an interface", ce.Name(), iface.Name())
			return
		}
		for j := range interfaces {
			if interfaces[j] == iface {
				if j >= numParentInterfaces {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s cannot implement previously implemented interface %s", ce.Name(), iface.Name())
					return
				}

				/* skip duplications */
				ce.ConstantsTable().Foreach(func(key string, c *types.ClassConstant) {
					DoInheritConstantCheck(iface.ConstantsTable(), c, key, iface)
				})

				iface = nil
				break
			}
		}
		if iface != nil {
			interfaces = append(interfaces, iface)
		}
	}
	ce.ResolvedInterfaces(interfaces)

	for i := numParentInterfaces; i < ce.GetNumInterfaces(); i++ {
		DoInterfaceImplementation(ce, ce.GetInterfaces()[i])
	}
}
func ZendAddMagicMethods(ce *types.ClassEntry, mname *types.String, fe types.IFunction) {
	if mname.GetStr() == "serialize" {
		ce.SetSerializeFunc(fe)
	} else if mname.GetStr() == "unserialize" {
		ce.SetUnserializeFunc(fe)
	} else if ce.GetName().GetLen() != mname.GetLen() && (mname.GetStr()[0] != '_' || mname.GetStr()[1] != '_') {

	} else if mname.GetStr() == ZEND_CLONE_FUNC_NAME {
		ce.SetClone(fe)
	} else if mname.GetStr() == ZEND_CONSTRUCTOR_FUNC_NAME {
		if ce.GetConstructor() != nil && (!(ce.GetParent()) || ce.GetConstructor() != ce.GetParent().constructor) {
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ce.Name())
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
		var lowercase_name *types.String = operators.ZendStringTolower(ce.GetName())
		// lowercase_name = types.ZendNewInternedString(lowercase_name)
		if !(memcmp(mname.GetVal(), lowercase_name.GetVal(), mname.GetLen())) {
			if ce.GetConstructor() != nil && (!(ce.GetParent()) || ce.GetConstructor() != ce.GetParent().constructor) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s has colliding constructor definitions coming from traits", ce.Name())
			}
			ce.SetConstructor(fe)
			fe.SetIsCtor(true)
		}
		// types.ZendStringReleaseEx(lowercase_name, 0)
	}
}
func ZendAddTraitMethod(ce *types.ClassEntry, name string, key string, fn types.IFunction, overridden **types.Array) {
	var existing_fn types.IFunction = nil
	var new_fn types.IFunction
	if lang.Assign(&existing_fn, ce.FunctionTable().Get(key)) != nil {

		/* if it is the same function with the same visibility and has not been assigned a class scope yet, regardless
		 * of where it is coming from there is no conflict and we do not need to add it again */

		if existing_fn.GetOpArray().GetOpcodes() == fn.GetOpArray().GetOpcodes() && (existing_fn.GetFnFlags()&types.AccPppMask) == (fn.GetFnFlags()&types.AccPppMask) && (existing_fn.GetScope().GetCeFlags()&types.AccTrait) == types.AccTrait {
			return
		}
		if existing_fn.GetScope() == ce {

			/* members from the current class override trait methods */

			if (*overridden) != nil {
				if lang.Assign(&existing_fn, types.ZendHashFindPtr(*overridden, key)) != nil {
					if existing_fn.IsAbstract() {
						/* Make sure the trait method is compatible with previosly declared abstract method */
						PerformDelayableImplementationCheck(ce, fn, existing_fn, 1)
					}
					if fn.IsAbstract() {
						/* Make sure the abstract declaration is compatible with previous declaration */
						PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
						return
					}
				}
			} else {
				*overridden = types.NewArrayCap(8)
			}
			types.ZendHashUpdateMem(*overridden, key, fn, b.SizeOf("zend_function"))
			return
		} else if fn.IsAbstract() && !existing_fn.IsAbstract() {
			/* Make sure the abstract declaration is compatible with previous declaration */
			PerformDelayableImplementationCheck(ce, existing_fn, fn, 1)
			return
		} else if existing_fn.GetScope().IsTrait() && !existing_fn.IsAbstract() {
			/* two traits can't define the __special__  same non-abstract method */
			faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Trait method %s has not been applied, because there are collisions with other trait methods on %s", name, ce.Name())
		} else {
			/* inherited members are overridden by members inserted by traits */
			DoInheritanceCheckOnMethodEx(fn, existing_fn, ce, nil, 0, 0)
			fn.SetPrototype(nil)
		}
	}
	if fn.IsInternalFunction() {
		new_fn = types.CopyInternalFunction(fn.GetInternalFunction())
		new_fn.SetIsArenaAllocated(true)
	} else {
		new_fn = types.CopyOpArray(fn.GetOpArray())
		new_fn.GetOpArray().SetIsTraitClone(true)
		new_fn.GetOpArray().SetIsImmutable(false)
	}
	FunctionAddRef(new_fn)
	ce.FunctionTable().Update(key, new_fn)
	ZendAddMagicMethods(ce, key, new_fn)
}
func ZendFixupTraitMethod(fn types.IFunction, ce *types.ClassEntry) {
	if (fn.GetScope().GetCeFlags() & types.AccTrait) == types.AccTrait {
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
	excludeTable *types.Array,
	aliases []*types.ClassEntry,
) {
	var alias *ZendTraitAlias
	var alias_ptr **ZendTraitAlias
	var fn_copy types.IFunction
	var i int

	/* apply aliases which are qualified with a class name, there should not be any ambiguity */

	if ce.GetTraitAliases() != nil {
		alias_ptr = ce.GetTraitAliases()
		alias = *alias_ptr
		i = 0
		for alias != nil {

			/* Scope unset or equal to the function we compare to, and the alias applies to fn */

			if alias.GetAlias() != "" && (aliases[i] == nil || fn.GetScope() == aliases[i]) && ascii.StrCaseEquals(alias.GetTraitMethod().MethodName(), fnname) {
				fn_copy = types.CopyFunction(fn)

				/* if it is 0, no modifieres has been changed */

				if alias.GetModifiers() != 0 {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&types.AccPppMask)
				}
				lcname := ascii.StrToLower(alias.GetAlias())
				ZendAddTraitMethod(ce, alias.GetAlias(), lcname, &fn_copy, overridden)

				/* Record the trait from which this alias was resolved. */
				if aliases[i] == nil {
					aliases[i] = fn.GetScope()
				}
				if alias.GetTraitMethod().ClassName() == "" {

					/* TODO: try to avoid this assignment (it's necessary only for reflection) */

					alias.GetTraitMethod().SetClassName(fn.GetScope().Name())
				}
			}
			alias_ptr++
			alias = *alias_ptr
			i++
		}
	}
	if excludeTable == nil || excludeTable.KeyFind(fnname.GetStr()) == nil {

		/* is not in hashtable, thus, function is not to be excluded */

		memcpy(&fn_copy, fn, lang.CondF(fn.GetType() == ZEND_USER_FUNCTION, func() __auto__ { return b.SizeOf("zend_op_array") }, func() __auto__ { return b.SizeOf("zend_internal_function") }))

		/* apply aliases which have not alias name, just setting visibility */

		if ce.GetTraitAliases() != nil {
			alias_ptr = ce.GetTraitAliases()
			alias = *alias_ptr
			i = 0
			for alias != nil {

				/* Scope unset or equal to the function we compare to, and the alias applies to fn */

				if alias.GetAlias() == "" && alias.GetModifiers() != 0 && (aliases[i] == nil || fn.GetScope() == aliases[i]) && ascii.StrCaseEquals(alias.GetTraitMethod().MethodName(), fnname) {
					fn_copy.SetFnFlags(alias.GetModifiers() | fn.GetFnFlags() ^ fn.GetFnFlags()&types.AccPppMask)

					/** Record the trait from which this alias was resolved. */
					if aliases[i] == nil {
						aliases[i] = fn.GetScope()
					}
					if alias.GetTraitMethod().ClassName() == "" {
						/* TODO: try to avoid this assignment (it's necessary only for reflection) */
						alias.GetTraitMethod().SetClassName(fn.GetScope().Name())
					}
				}
				alias_ptr++
				alias = *alias_ptr
				i++
			}
		}
		ZendAddTraitMethod(ce, fn.FunctionName(), fnname, &fn_copy, overridden)
	}
}
func ZendCheckTraitUsage(ce *types.ClassEntry, trait *types.ClassEntry, traits **types.ClassEntry) uint32 {
	var i uint32
	if (trait.GetCeFlags() & types.AccTrait) != types.AccTrait {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Class %s is not a trait, Only traits may be used in 'as' and 'insteadof' statements", trait.Name())
		return 0
	}
	for i = 0; i < ce.GetNumTraits(); i++ {
		if traits[i] == trait {
			return i
		}
	}
	faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Required Trait %s wasn't added to %s", trait.Name(), ce.Name())
	return 0
}
func ZendTraitsInitTraitStructures(ce *types.ClassEntry, traits **types.ClassEntry, exclude_tables_ptr ***types.Array, aliases_ptr ***types.ClassEntry) {
	var i int
	var j int = 0
	var precedences []*ZendTraitPrecedence
	var cur_precedence *ZendTraitPrecedence
	var cur_method_ref *ZendTraitMethodReference
	var exclude_tables []*types.Array = nil
	var aliases **types.ClassEntry = nil
	var trait *types.ClassEntry

	/* resolve class references */

	if ce.GetTraitPrecedences() != nil {
		exclude_tables = Ecalloc(ce.GetNumTraits(), b.SizeOf("HashTable *"))
		i = 0
		precedences = ce.GetTraitPrecedences()
		ce.SetTraitPrecedences(nil)
		for lang.Assign(&cur_precedence, precedences[i]) {

			/** Resolve classes for all precedence operations. */

			cur_method_ref = cur_precedence.GetTraitMethod()
			trait = ZendFetchClass(cur_method_ref.ClassName(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
			if trait == nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", cur_method_ref.ClassName())
			}
			ZendCheckTraitUsage(ce, trait, traits)

			/** Ensure that the preferred method is actually available. */

			lcname := ascii.StrToLower(cur_method_ref.MethodName())
			if !trait.FunctionTable().Exists(lcname) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "A precedence rule was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.MethodName())
			}

			/** With the other traits, we are more permissive.
			  We do not give errors for those. This allows to be more
			  defensive in such definitions.
			  However, we want to make sure that the insteadof declaration
			  is consistent in itself.
			*/
			for _, class_name := range cur_precedence.GetExcludeClassNames() {
				var exclude_ce *types.ClassEntry = ZendFetchClass(class_name, ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				var trait_num uint32
				if exclude_ce == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", class_name)
				}
				trait_num = ZendCheckTraitUsage(ce, exclude_ce, traits)
				if exclude_tables[trait_num] == nil {
					exclude_tables[trait_num] = types.NewArray()
				}
				if types.ZendHashAddEmptyElement(exclude_tables[trait_num], lcname) == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Failed to evaluate a trait precedence (%s). Method of trait %s was defined to be excluded multiple times", precedences[i].GetTraitMethod().MethodName(), exclude_ce.Name())
				}

				/* make sure that the trait method is not from a class mentioned in
				   exclude_from_classes, for consistency */

				if trait == exclude_ce {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Inconsistent insteadof definition. "+"The method %s is to be used from %s, but %s is also on the exclude list", cur_method_ref.MethodName(), trait.GetName().GetVal(), trait.GetName().GetVal())
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

			if ce.GetTraitAliases()[i].GetTraitMethod().ClassName() != "" {
				cur_method_ref = ce.GetTraitAliases()[i].GetTraitMethod()
				trait = ZendFetchClass(cur_method_ref.ClassName(), ZEND_FETCH_CLASS_TRAIT|ZEND_FETCH_CLASS_NO_AUTOLOAD)
				if trait == nil {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Could not find trait %s", cur_method_ref.ClassName())
				}
				ZendCheckTraitUsage(ce, trait, traits)
				aliases[i] = trait

				/** And, ensure that the referenced method is resolvable, too. */

				lcname := ascii.StrToLower(cur_method_ref.MethodName())
				if !trait.FunctionTable().Exists(lcname) {
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "An alias was defined for %s::%s but this method does not exist", trait.GetName().GetVal(), cur_method_ref.MethodName())
				}
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
func ZendDoTraitsPropertyBinding(ce *types.ClassEntry, traits []*types.ClassEntry) {
	var coliding_prop *types.PropertyInfo
	var prop_name *types.String
	var class_name_unused *byte
	var not_compatible bool
	var prop_value *types.Zval
	var flags uint32
	var doc_comment *string

	/* In the following steps the properties are inserted into the property table
	 * for that, a very strict approach is applied:
	 * - check for compatibility, if not compatible with any property in class -> fatal
	 * - if compatible, then strict notice
	 */
	for i, trait := range traits {
		if trait == nil {
			continue
		}
		trait.PropertyTable().Foreach(func(_ string, property_info *types.PropertyInfo) {
			/* first get the unmangeld name if necessary,
			 * then check whether the property is already there
			 */
			flags = property_info.GetFlags()
			if (flags & types.AccPublic) != 0 {
				prop_name = types.NewString(property_info.GetName())
			} else {
				/* for private and protected we need to unmangle the names */
				_, propNameStr, _ := ZendUnmanglePropertyName_Ex(property_info.GetName())
				prop_name = types.NewString(propNameStr)
			}

			/* next: check for conflicts with current class */

			coliding_prop = ce.PropertyTable().Get(prop_name.GetStr())
			if coliding_prop != nil {
				if coliding_prop.IsPrivate() && coliding_prop.GetCe() != ce {
					ce.PropertyTable().Del(prop_name.GetStr())
					flags |= types.AccChanged
				} else {
					not_compatible = 1
					if (coliding_prop.GetFlags()&(types.AccPppMask|types.AccStatic)) == (flags&(types.AccPppMask|types.AccStatic)) && PropertyTypesCompatible(property_info, coliding_prop) == INHERITANCE_SUCCESS {

						/* the flags are identical, thus, the properties may be compatible */

						var op1 *types.Zval
						var op2 *types.Zval
						var op1_tmp types.Zval
						var op2_tmp types.Zval
						if (flags & types.AccStatic) != 0 {
							op1 = ce.GetDefaultStaticMembersTable()[coliding_prop.GetOffset()]
							op2 = trait.GetDefaultStaticMembersTable()[property_info.GetOffset()]
							op1 = types.ZVAL_DEINDIRECT(op1)
							op2 = types.ZVAL_DEINDIRECT(op2)
						} else {
							op1 = ce.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(coliding_prop.GetOffset())]
							op2 = trait.GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
						}

						/* if any of the values is a constant, we try to resolve it */

						if op1.IsConstantAst() {
							types.ZVAL_COPY_OR_DUP(&op1_tmp, op1)
							ZvalUpdateConstantEx(&op1_tmp, ce)
							op1 = &op1_tmp
						}
						if op2.IsConstantAst() {
							types.ZVAL_COPY_OR_DUP(&op2_tmp, op2)
							ZvalUpdateConstantEx(&op2_tmp, ce)
							op2 = &op2_tmp
						}
						not_compatible = operators.FastIsNotIdenticalFunction(op1, op2)
						if op1 == &op1_tmp {
							// ZvalPtrDtorNogc(&op1_tmp)
						}
						if op2 == &op2_tmp {
							// ZvalPtrDtorNogc(&op2_tmp)
						}
					}
					if not_compatible != 0 {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "%s and %s define the __special__  same property ($%s) in the composition of %s. However, the definition differs and is considered incompatible. Class was composed", FindFirstDefinition(ce, traits, i, prop_name, coliding_prop.GetCe()).GetName().GetVal(), property_info.GetCe().Name(), prop_name.GetVal(), ce.Name())
					}
					// types.ZendStringReleaseEx(prop_name, 0)
					return
				}
			}

			/* property not found, so lets add it */
			if (flags & types.AccStatic) != 0 {
				prop_value = traits[i].GetDefaultStaticMembersTable()[property_info.GetOffset()]
				b.Assert(!prop_value.IsIndirect())
			} else {
				prop_value = traits[i].GetDefaultPropertiesTable()[OBJ_PROP_TO_NUM(property_info.GetOffset())]
			}
			doc_comment = property_info.GetDocComment()
			ZendDeclareTypedProperty(ce, prop_name, prop_value, flags, doc_comment, property_info.GetType())
		})
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
	if ce.GetTraitAliases() != nil {
		for ce.GetTraitAliases()[i] != nil {
			cur_alias = ce.GetTraitAliases()[i]

			/** The trait for this alias has not been resolved, this means, this
			  alias was not applied. Abort with an error. */

			if aliases[i] == nil {
				if cur_alias.GetAlias() != "" {
					/** Plain old inconsistency/typo/bug */
					faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "An alias (%s) was defined for method %s(), but this method does not exist", cur_alias.GetAlias(), cur_alias.GetTraitMethod().MethodName())
				} else {

					/** Here are two possible cases:
					  1) this is an attempt to modify the visibility
					     of a method introduce as part of another alias.
					     Since that seems to violate the DRY principle,
					     we check against it and abort.
					  2) it is just a plain old inconsitency/typo/bug
					     as in the case where alias is set. */

					lcMethodName := ascii.StrToLower(cur_alias.GetTraitMethod().MethodName())
					if ce.FunctionTable().Exists(lcMethodName) {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The modifiers for the trait alias %s() need to be changed in the same statement in which the alias is defined. Error", cur_alias.GetTraitMethod().MethodName())
					} else {
						faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "The modifiers of the trait method %s() are changed, but this method does not exist. Error", cur_alias.GetTraitMethod().MethodName())
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
			faults.ErrorNoreturn(faults.E_ERROR, "%s cannot use %s - it is not a trait", ce.Name(), trait.GetName().GetVal())
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
func ZendHasDeprecatedConstructor(ce *types.ClassEntry) bool {
	if ce.GetConstructor() == nil {
		return false
	}
	return ascii.StrCaseEquals(ce.Name(), ce.GetConstructor().FunctionName())
}
func ZendCheckDeprecatedConstructor(ce *types.ClassEntry) {
	if ZendHasDeprecatedConstructor(ce) {
		faults.Error(faults.E_DEPRECATED, "Methods with the same name as their class will not be constructors in a future version of PHP; %s has a deprecated constructor", ce.Name())
	}
}
func DISPLAY_ABSTRACT_FN(idx int) {
	lang.CondF1(ai.afn[idx], func() string { return ZEND_FN_SCOPE_NAME(ai.afn[idx]) }, "")
	lang.Cond(ai.afn[idx], "::", "")
	lang.CondF1(ai.afn[idx], func() []byte { return ai.afn[idx].common.function_name.GetVal() }, "")
	lang.CondF2(ai.afn[idx] && ai.afn[idx+1], ", ", func() string {
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
	b.Assert((ce.GetCeFlags() & (types.AccImplicitAbstractClass | types.AccInterface | types.AccTrait | types.AccExplicitAbstractClass)) == types.AccImplicitAbstractClass)
	memset(&ai, 0, b.SizeOf("ai"))
	ce.FunctionTable().Foreach(func(_ string, func_ types.IFunction) {
		ZendVerifyAbstractClassFunction(func_, &ai)
	})
	if ai.GetCnt() != 0 {
		faults.ErrorNoreturn(faults.E_ERROR, "Class %s contains %d abstract method%s and must therefore be declared abstract or implement the remaining methods ("+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+MAX_ABSTRACT_INFO_FMT+")", ce.Name(), ai.GetCnt(), lang.Cond(ai.GetCnt() > 1, "s", ""), DISPLAY_ABSTRACT_FN(0), DISPLAY_ABSTRACT_FN(1), DISPLAY_ABSTRACT_FN(2))
	} else {

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

		ce.SetIsImplicitAbstractClass(false)

		/* now everything should be fine and an added ZEND_ACC_IMPLICIT_ABSTRACT_CLASS should be removed */

	}
}
func GetOrInitObligationsForClass(ce *types.ClassEntry) *types.Array {
	var ht *types.Array
	var key ZendUlong
	if CG__().GetDelayedVarianceObligations() == nil {
		CG__().SetDelayedVarianceObligations(types.NewArray())
	}
	key = ZendUlong(uintPtr(ce))
	ht = types.ZendHashIndexFindPtr(CG__().GetDelayedVarianceObligations(), key)
	if ht != nil {
		return ht
	}
	ht = types.NewArray()
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
func AddCompatibilityObligation(ce *types.ClassEntry, child_fn types.IFunction, parent_fn types.IFunction, always_error bool) {
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
func AddPropertyCompatibilityObligation(ce *types.ClassEntry, child_prop *types.PropertyInfo, parent_prop *types.PropertyInfo) {
	var obligations *types.Array = GetOrInitObligationsForClass(ce)
	var obligation *VarianceObligation = Emalloc(b.SizeOf("variance_obligation"))
	obligation.SetType(OBLIGATION_PROPERTY_COMPATIBILITY)
	obligation.SetChildProp(child_prop)
	obligation.SetParentProp(parent_prop)
	types.ZendHashNextIndexInsertPtr(obligations, obligation)
}
func CheckVarianceObligation(_ types.ArrayKey, zv *types.Zval) bool {
	var obligation *VarianceObligation = zv.Ptr()
	if obligation.GetType() == OBLIGATION_DEPENDENCY {
		var dependency_ce *types.ClassEntry = obligation.GetDependencyCe()
		if dependency_ce.IsUnresolvedVariance() {
			ResolveDelayedVarianceObligations(dependency_ce)
		}
		if !dependency_ce.IsLinked() {
			return true
		}
	} else if obligation.GetType() == OBLIGATION_COMPATIBILITY {
		var unresolved_class *types.String
		var status InheritanceStatus = ZendDoPerformImplementationCheck(&unresolved_class, obligation.GetChildFn(), obligation.GetParentFn())
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return true
			}
			b.Assert(status == INHERITANCE_ERROR)
			EmitIncompatibleMethodErrorOrWarning(obligation.GetChildFn(), obligation.GetParentFn(), status, unresolved_class, obligation.GetAlwaysError())
		}
	} else {
		b.Assert(obligation.GetType() == OBLIGATION_PROPERTY_COMPATIBILITY)
		var status InheritanceStatus = PropertyTypesCompatible(obligation.GetParentProp(), obligation.GetChildProp())
		if status != INHERITANCE_SUCCESS {
			if status == INHERITANCE_UNRESOLVED {
				return true
			}
			b.Assert(status == INHERITANCE_ERROR)
			EmitIncompatiblePropertyError(obligation.GetChildProp(), obligation.GetParentProp())
		}
	}
	return false
}
func LoadDelayedClasses() {
	var delayed_autoloads *types.Array = CG__().GetDelayedAutoloads()
	if delayed_autoloads == nil {
		return
	}

	/* Take ownership of this HT, to avoid concurrent modification during autoloading. */

	CG__().SetDelayedAutoloads(nil)
	delayed_autoloads.Foreach(func(key types.ArrayKey, value *types.Zval) {
		name := key.StrKey()
		ZendLookupClassString(name)
	})
	delayed_autoloads.Destroy()
}
func ResolveDelayedVarianceObligations(ce *types.ClassEntry) {
	var all_obligations *types.Array = CG__().GetDelayedVarianceObligations()
	var obligations *types.Array
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	b.Assert(all_obligations != nil)
	obligations = types.ZendHashIndexFindPtr(all_obligations, num_key)
	b.Assert(obligations != nil)
	obligations.Filter(CheckVarianceObligation)
	if obligations.Len() == 0 {
		ce.SetIsUnresolvedVariance(false)
		ce.SetIsLinked(true)
		all_obligations.IndexDelete(num_key)
	}
}
func ReportVarianceErrors(ce *types.ClassEntry) {
	var all_obligations *types.Array = CG__().GetDelayedVarianceObligations()
	var obligations *types.Array
	var num_key ZendUlong = ZendUlong(uintPtr(ce))
	b.Assert(all_obligations != nil)
	obligations = types.ZendHashIndexFindPtr(all_obligations, num_key)
	b.Assert(obligations != nil)
	obligations.Foreach(func(_ types.ArrayKey, value *types.Zval) {
		var obligation *VarianceObligation = value.Ptr()
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
	})

	/* Only warnings were thrown above -- that means that there are incompatibilities, but only
	 * ones that we permit. Mark all classes with open obligations as fully linked. */

	ce.SetIsUnresolvedVariance(false)
	ce.SetIsLinked(true)
	all_obligations.IndexDelete(num_key)
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
		// 		exception_zv.AddRefcount()
		EG__().ClearException()
		exception_str = operators.ZvalGetString(&exception_zv)
		faults.ErrorNoreturn(faults.E_ERROR, "During inheritance of %s with variance dependencies: Uncaught %s", ce.Name(), exception_str.GetVal())
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
		var num_parent_interfaces uint32 = lang.CondF1(parent != nil, func() uint32 { return parent.GetNumInterfaces() }, 0)
		interfaces = Emalloc(b.SizeOf("zend_class_entry *") * (ce.GetNumInterfaces() + num_parent_interfaces))
		if num_parent_interfaces != 0 {
			memcpy(interfaces, parent.GetInterfaces(), b.SizeOf("zend_class_entry *")*num_parent_interfaces)
		}
		for i = 0; i < ce.GetNumInterfaces(); i++ {
			var iface *types.ClassEntry = ZendFetchClassByName_Ex(ce.GetInterfaceNames()[i].GetName(), ce.GetInterfaceNames()[i].GetLcName(), ZEND_FETCH_CLASS_INTERFACE|ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED|ZEND_FETCH_CLASS_EXCEPTION)
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
	if (ce.GetCeFlags() & (types.AccImplicitAbstractClass | types.AccInterface | types.AccTrait | types.AccExplicitAbstractClass)) == types.AccImplicitAbstractClass {
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
	var parent_info *types.PropertyInfo

	parent_ce.FunctionTable().ForeachEx(func(key string, parent_func types.IFunction) bool {
		var child_func types.IFunction = ce.FunctionTable().Get(key)
		if child_func != nil {
			var status InheritanceStatus = DoInheritanceCheckOnMethodEx(child_func, parent_func, ce, nil, 1, 0)
			if status != INHERITANCE_SUCCESS {
				b.Assert(status == INHERITANCE_UNRESOLVED || status == INHERITANCE_ERROR)
				ret = status
				if status == INHERITANCE_UNRESOLVED {
					return false
				}
			}
		}
		return true
	})
	if ret == INHERITANCE_UNRESOLVED {
		return ret
	}

	parent_ce.PropertyTable().ForeachEx(func(key string, parent_info *types.PropertyInfo) bool {
		if parent_info.IsPrivate() || !(parent_info.GetType().IsSet()) {
			return true
		}
		var childInfo *types.PropertyInfo = ce.PropertyTable().Get(key)
		if childInfo != nil {
			if childInfo.GetType().IsSet() {
				var status InheritanceStatus = PropertyTypesCompatible(parent_info, childInfo)
				if status != INHERITANCE_SUCCESS {
					b.Assert(status == INHERITANCE_UNRESOLVED || status == INHERITANCE_ERROR)
					ret = status
					if status == INHERITANCE_UNRESOLVED {
						return false
					}
				}
			}
		}
		return true
	})

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
		if (ce.GetCeFlags() & (types.AccImplicitAbstractClass | types.AccInterface | types.AccTrait | types.AccExplicitAbstractClass)) == types.AccImplicitAbstractClass {
			ZendVerifyAbstractClass(ce)
		}
		b.Assert(!ce.IsUnresolvedVariance())
		ce.SetIsLinked(true)
		return true
	}
	return false
}
