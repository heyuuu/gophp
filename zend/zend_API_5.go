// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendStartupModuleZval(zv *Zval) int {
	var module *ZendModuleEntry = zv.GetPtr()
	if ZendStartupModuleEx(module) == SUCCESS {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}
func ZendSortModules(base any, count int, siz int, compare CompareFuncT, swp SwapFuncT) {
	var b1 *Bucket = base
	var b2 *Bucket
	var end *Bucket = b1 + count
	var tmp Bucket
	var m *ZendModuleEntry
	var r *ZendModuleEntry
	for b1 < end {
	try_again:
		m = (*ZendModuleEntry)(b1.GetVal().GetPtr())
		if m.GetModuleStarted() == 0 && m.GetDeps() != nil {
			var dep *ZendModuleDep = m.GetDeps()
			for dep.GetName() != nil {
				if dep.GetType() == MODULE_DEP_REQUIRED || dep.GetType() == MODULE_DEP_OPTIONAL {
					b2 = b1 + 1
					for b2 < end {
						r = (*ZendModuleEntry)(b2.GetVal().GetPtr())
						if strcasecmp(dep.GetName(), r.GetName()) == 0 {
							tmp = *b1
							*b1 = *b2
							*b2 = tmp
							goto try_again
						}
						b2++
					}
				}
				dep++
			}
		}
		b1++
	}
}
func ZendCollectModuleHandlers() {
	var module *ZendModuleEntry
	var startup_count int = 0
	var shutdown_count int = 0
	var post_deactivate_count int = 0
	var ce *ZendClassEntry
	var class_count int = 0

	/* Collect extensions with request startup/shutdown handlers */

	var __ht *HashTable = &ModuleRegistry
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		module = _z.GetPtr()
		if module.GetRequestStartupFunc() != nil {
			startup_count++
		}
		if module.GetRequestShutdownFunc() != nil {
			shutdown_count++
		}
		if module.GetPostDeactivateFunc() != nil {
			post_deactivate_count++
		}
	}
	ModuleRequestStartupHandlers = (**ZendModuleEntry)(Malloc(b.SizeOf("zend_module_entry *") * (startup_count + 1 + shutdown_count + 1 + post_deactivate_count + 1)))
	ModuleRequestStartupHandlers[startup_count] = nil
	ModuleRequestShutdownHandlers = ModuleRequestStartupHandlers + startup_count + 1
	ModuleRequestShutdownHandlers[shutdown_count] = nil
	ModulePostDeactivateHandlers = ModuleRequestShutdownHandlers + shutdown_count + 1
	ModulePostDeactivateHandlers[post_deactivate_count] = nil
	startup_count = 0
	var __ht__1 *HashTable = &ModuleRegistry
	for _, _p := range __ht__1.foreachData() {
		var _z *Zval = _p.GetVal()

		module = _z.GetPtr()
		if module.GetRequestStartupFunc() != nil {
			ModuleRequestStartupHandlers[b.PostInc(&startup_count)] = module
		}
		if module.GetRequestShutdownFunc() != nil {
			ModuleRequestShutdownHandlers[b.PreDec(&shutdown_count)] = module
		}
		if module.GetPostDeactivateFunc() != nil {
			ModulePostDeactivateHandlers[b.PreDec(&post_deactivate_count)] = module
		}
	}

	/* Collect internal classes with static members */

	var __ht__2 *HashTable = CG__().GetClassTable()
	for _, _p := range __ht__2.foreachData() {
		var _z *Zval = _p.GetVal()

		ce = _z.GetPtr()
		if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetDefaultStaticMembersCount() > 0 {
			class_count++
		}
	}
	ClassCleanupHandlers = (**ZendClassEntry)(Malloc(b.SizeOf("zend_class_entry *") * (class_count + 1)))
	ClassCleanupHandlers[class_count] = nil
	if class_count != 0 {
		var __ht *HashTable = CG__().GetClassTable()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

			ce = _z.GetPtr()
			if ce.GetType() == ZEND_INTERNAL_CLASS && ce.GetDefaultStaticMembersCount() > 0 {
				ClassCleanupHandlers[b.PreDec(&class_count)] = ce
			}
		}
	}
}
func ZendStartupModules() int {
	ModuleRegistry.SortCompatibleEx(ZendSortModules)
	ZendHashApply(&ModuleRegistry, ZendStartupModuleZval)
	return SUCCESS
}
func ZendDestroyModules() {
	Free(ClassCleanupHandlers)
	Free(ModuleRequestStartupHandlers)
	ModuleRegistry.GracefulReverseDestroy()
}
func ZendRegisterModuleEx(module *ZendModuleEntry) *ZendModuleEntry {
	var name_len int
	var lcname *ZendString
	var module_ptr *ZendModuleEntry
	if module == nil {
		return nil
	}

	/* Check module dependencies */

	if module.GetDeps() != nil {
		var dep *ZendModuleDep = module.GetDeps()
		for dep.GetName() != nil {
			if dep.GetType() == MODULE_DEP_CONFLICTS {
				name_len = strlen(dep.GetName())
				lcname = ZendStringAlloc(name_len, 0)
				ZendStrTolowerCopy(lcname.GetVal(), dep.GetName(), name_len)
				if ZendHashExists(&ModuleRegistry, lcname) != 0 || ZendGetExtension(dep.GetName()) != nil {
					ZendStringEfree(lcname)

					/* TODO: Check version relationship */

					ZendError(E_CORE_WARNING, "Cannot load module '%s' because conflicting module '%s' is already loaded", module.GetName(), dep.GetName())
					return nil
				}
				ZendStringEfree(lcname)
			}
			dep++
		}
	}
	name_len = strlen(module.GetName())
	lcname = ZendStringAlloc(name_len, module.GetType() == MODULE_PERSISTENT)
	ZendStrTolowerCopy(lcname.GetVal(), module.GetName(), name_len)
	lcname = ZendNewInternedString(lcname)
	if b.Assign(&module_ptr, ZendHashAddMem(&ModuleRegistry, lcname, module, b.SizeOf("zend_module_entry"))) == nil {
		ZendError(E_CORE_WARNING, "Module '%s' already loaded", module.GetName())
		ZendStringRelease(lcname)
		return nil
	}
	module = module_ptr
	EG__().SetCurrentModule(module)
	if module.GetFunctions() != nil && ZendRegisterFunctions(nil, module.GetFunctions(), nil, module.GetType()) == FAILURE {
		ZendHashDel(&ModuleRegistry, lcname)
		ZendStringRelease(lcname)
		EG__().SetCurrentModule(nil)
		ZendError(E_CORE_WARNING, "%s: Unable to register functions, unable to load", module.GetName())
		return nil
	}
	EG__().SetCurrentModule(nil)
	ZendStringRelease(lcname)
	return module
}
func ZendRegisterInternalModule(module *ZendModuleEntry) *ZendModuleEntry {
	module.SetModuleNumber(ZendNextFreeModule())
	module.SetType(MODULE_PERSISTENT)
	return ZendRegisterModuleEx(module)
}
func ZendCheckMagicMethodImplementation(ce *ZendClassEntry, fptr *ZendFunction, error_type int) {
	var lcname []byte
	var name_len int
	if fptr.GetFunctionName().GetVal()[0] != '_' || fptr.GetFunctionName().GetVal()[1] != '_' {
		return
	}

	/* we don't care if the function name is longer, in fact lowercasing only
	 * the beginning of the name speeds up the check process */

	name_len = fptr.GetFunctionName().GetLen()
	ZendStrTolowerCopy(lcname, fptr.GetFunctionName().GetVal(), MIN(name_len, b.SizeOf("lcname")-1))
	lcname[b.SizeOf("lcname")-1] = '0'
	if name_len == b.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_DESTRUCTOR_FUNC_NAME, b.SizeOf("ZEND_DESTRUCTOR_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Destructor %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_DESTRUCTOR_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_CLONE_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CLONE_FUNC_NAME, b.SizeOf("ZEND_CLONE_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot accept any arguments", ce.GetName().GetVal(), ZEND_CLONE_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_GET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_GET_FUNC_NAME, b.SizeOf("ZEND_GET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_GET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_GET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_SET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_SET_FUNC_NAME, b.SizeOf("ZEND_SET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), ZEND_SET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_SET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_UNSET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_UNSET_FUNC_NAME, b.SizeOf("ZEND_UNSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_UNSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_UNSET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_ISSET_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_ISSET_FUNC_NAME, b.SizeOf("ZEND_ISSET_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 1 {
			ZendError(error_type, "Method %s::%s() must take exactly 1 argument", ce.GetName().GetVal(), ZEND_ISSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_ISSET_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_CALL_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CALL_FUNC_NAME, b.SizeOf("ZEND_CALL_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::%s() must take exactly 2 arguments", ce.GetName().GetVal(), ZEND_CALL_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::%s() cannot take arguments by reference", ce.GetName().GetVal(), ZEND_CALL_FUNC_NAME)
		}
	} else if name_len == b.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_CALLSTATIC_FUNC_NAME, b.SizeOf("ZEND_CALLSTATIC_FUNC_NAME")-1)) {
		if fptr.GetNumArgs() != 2 {
			ZendError(error_type, "Method %s::__callStatic() must take exactly 2 arguments", ce.GetName().GetVal())
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			ZendError(error_type, "Method %s::__callStatic() cannot take arguments by reference", ce.GetName().GetVal())
		}
	} else if name_len == b.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_TOSTRING_FUNC_NAME, b.SizeOf("ZEND_TOSTRING_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_TOSTRING_FUNC_NAME)
	} else if name_len == b.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1 && !(memcmp(lcname, ZEND_DEBUGINFO_FUNC_NAME, b.SizeOf("ZEND_DEBUGINFO_FUNC_NAME")-1)) && fptr.GetNumArgs() != 0 {
		ZendError(error_type, "Method %s::%s() cannot take arguments", ce.GetName().GetVal(), ZEND_DEBUGINFO_FUNC_NAME)
	}
}
func ZendRegisterFunctions(scope *ZendClassEntry, functions *ZendFunctionEntry, function_table *HashTable, type_ int) int {
	var ptr *ZendFunctionEntry = functions
	var count int = 0
	var unload int = 0
	var target_function_table *HashTable = function_table
	var error_type int
	var ctor *ZendFunction = nil
	var dtor *ZendFunction = nil
	var clone *ZendFunction = nil
	var __get *ZendFunction = nil
	var __set *ZendFunction = nil
	var __unset *ZendFunction = nil
	var __isset *ZendFunction = nil
	var __call *ZendFunction = nil
	var __callstatic *ZendFunction = nil
	var __tostring *ZendFunction = nil
	var __debugInfo *ZendFunction = nil
	var serialize_func *ZendFunction = nil
	var unserialize_func *ZendFunction = nil
	var lowercase_name *ZendString
	var fname_len int
	var lc_class_name *byte = nil
	var class_name_len int = 0
	if type_ == MODULE_PERSISTENT {
		error_type = E_CORE_WARNING
	} else {
		error_type = E_WARNING
	}
	if target_function_table == nil {
		target_function_table = CG__().GetFunctionTable()
	}
	var reg_function *ZendFunction

	var internal_function = NewInternalFunction()
	var function ZendFunction = MakeZendFunctionInternal(internal_function)

	internal_function.SetModule(EG__().GetCurrentModule())
	if scope != nil {
		class_name_len = scope.GetName().GetLen()
		if b.Assign(&lc_class_name, ZendMemrchr(scope.GetName().GetVal(), '\\', class_name_len)) {
			lc_class_name++
			class_name_len -= lc_class_name - scope.GetName().GetVal()
			lc_class_name = ZendStrTolowerDup(lc_class_name, class_name_len)
		} else {
			lc_class_name = ZendStrTolowerDup(scope.GetName().GetVal(), class_name_len)
		}
	}
	for ptr.GetFname() != nil {
		fname_len = strlen(ptr.GetFname())

		internal_function.InitByEntry(ptr)
		internal_function.SetScope(scope)
		if ptr.GetFlags() != 0 {
			if !ptr.IsPppMask() {
				if ptr.GetFlags() != ZEND_ACC_DEPRECATED && scope != nil {
					ZendError(error_type, "Invalid access level for %s%s%s() - access must be exactly one of public, protected or private", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
				}
				internal_function.SetFnFlags(ZEND_ACC_PUBLIC | ptr.GetFlags())
			} else {
				internal_function.SetFnFlags(ptr.GetFlags())
			}
		} else {
			internal_function.SetFnFlags(ZEND_ACC_PUBLIC)
		}
		if ptr.GetArgInfo() != nil {
			var info *ZendInternalFunctionInfo = (*ZendInternalFunctionInfo)(ptr.GetArgInfo())
			internal_function.SetArgInfo((*ArgInfo)(ptr.GetArgInfo() + 1))
			internal_function.SetNumArgs(ptr.GetNumArgs())

			/* Currently you cannot denote that the function can accept less arguments than num_args */

			if info.GetRequiredNumArgs() == -1 {
				internal_function.SetRequiredNumArgs(ptr.GetNumArgs())
			} else {
				internal_function.SetRequiredNumArgs(info.GetRequiredNumArgs())
			}
			if info.GetReturnReference() != 0 {
				internal_function.SetIsReturnReference(true)
			}
			if ptr.GetArgInfo()[ptr.GetNumArgs()].GetIsVariadic() != 0 {
				internal_function.SetIsVariadic(true)

				/* Don't count the variadic argument */

				internal_function.GetNumArgs()--

				/* Don't count the variadic argument */

			}
			if info.GetType().IsSet() {
				if info.GetType().IsClass() {
					var type_name *byte = (*byte)(info.GetType())
					if type_name[0] == '?' {
						type_name++
					}
					if scope == nil && (!(strcasecmp(type_name, "self")) || !(strcasecmp(type_name, "parent"))) {
						ZendErrorNoreturn(E_CORE_ERROR, "Cannot declare a return type of %s outside of a class scope", type_name)
					}
				}
				internal_function.SetIsHasReturnType(true)
			}
		} else {
			internal_function.SetArgInfo(nil)
			internal_function.SetNumArgs(0)
			internal_function.SetRequiredNumArgs(0)
		}
		if ptr.IsAbstract() {
			if scope != nil {

				/* This is a class that must be abstract itself. Here we set the check info. */

				scope.SetIsImplicitAbstractClass(true)
				if !scope.IsInterface() {

					/* Since the class is not an interface it needs to be declared as a abstract class. */

					scope.SetIsExplicitAbstractClass(true)

					/* Since the class is not an interface it needs to be declared as a abstract class. */

				}
			}
			if ptr.IsStatic() && (scope == nil || !scope.IsInterface()) {
				ZendError(error_type, "Static function %s%s%s() cannot be abstract", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
		} else {
			if scope != nil && scope.IsInterface() {
				Efree((*byte)(lc_class_name))
				ZendError(error_type, "Interface %s cannot contain non abstract method %s()", scope.GetName().GetVal(), ptr.GetFname())
				return FAILURE
			}
			if internal_function.GetHandler() == nil {
				if scope != nil {
					Efree((*byte)(lc_class_name))
				}
				ZendError(error_type, "Method %s%s%s() cannot be a NULL function", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
				ZendUnregisterFunctions(functions, count, target_function_table)
				return FAILURE
			}
		}
		lowercase_name = ZendStringTolowerEx(internal_function.GetFunctionName(), type_ == MODULE_PERSISTENT)
		lowercase_name = ZendNewInternedString(lowercase_name)
		reg_function = Malloc(b.SizeOf("zend_internal_function"))
		memcpy(reg_function, &function, b.SizeOf("zend_internal_function"))
		if ZendHashAddPtr(target_function_table, lowercase_name, reg_function) == nil {
			unload = 1
			Free(reg_function)
			ZendStringRelease(lowercase_name)
			break
		}

		/* If types of arguments have to be checked */

		if reg_function.GetArgInfo() != nil && reg_function.GetNumArgs() != 0 {
			var i uint32
			for i = 0; i < reg_function.GetNumArgs(); i++ {
				if reg_function.GetArgInfo()[i].GetType().IsSet() {
					reg_function.SetIsHasTypeHints(true)
					break
				}
			}
		}
		if reg_function.GetArgInfo() != nil && reg_function.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) {

			/* convert "const char*" class type names into "zend_string*" */

			var i uint32
			var num_args uint32 = reg_function.GetNumArgs() + 1
			var arg_info *ZendArgInfo = reg_function.GetArgInfo() - 1
			var new_arg_info *ZendArgInfo
			if reg_function.IsVariadic() {
				num_args++
			}
			new_arg_info = Malloc(b.SizeOf("zend_arg_info") * num_args)
			memcpy(new_arg_info, arg_info, b.SizeOf("zend_arg_info")*num_args)
			reg_function.SetArgInfo(new_arg_info + 1)
			for i = 0; i < num_args; i++ {
				if new_arg_info[i].GetType().IsClass() {
					var class_name *byte = (*byte)(new_arg_info[i].GetType())
					var allow_null ZendBool = 0
					var str *ZendString
					if class_name[0] == '?' {
						class_name++
						allow_null = 1
					}
					str = ZendStringInitInterned(class_name, strlen(class_name), 1)
					new_arg_info[i].SetType(ZEND_TYPE_ENCODE_CLASS(str, allow_null))
				}
			}
		}
		if scope != nil {

			/* Look for ctor, dtor, clone
			 * If it's an old-style constructor, store it only if we don't have
			 * a constructor already.
			 */

			if fname_len == class_name_len && ctor == nil && !(memcmp(lowercase_name.GetVal(), lc_class_name, class_name_len+1)) {
				ctor = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, "serialize") {
				serialize_func = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, "unserialize") {
				unserialize_func = reg_function
			} else if lowercase_name.GetVal()[0] != '_' || lowercase_name.GetVal()[1] != '_' {
				reg_function = nil
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CONSTRUCTOR_FUNC_NAME) {
				ctor = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_DESTRUCTOR_FUNC_NAME) {
				dtor = reg_function
				if internal_function.GetNumArgs() != 0 {
					ZendError(error_type, "Destructor %s::%s() cannot take arguments", scope.GetName().GetVal(), ptr.GetFname())
				}
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CLONE_FUNC_NAME) {
				clone = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CALL_FUNC_NAME) {
				__call = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_CALLSTATIC_FUNC_NAME) {
				__callstatic = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_TOSTRING_FUNC_NAME) {
				__tostring = reg_function
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_GET_FUNC_NAME) {
				__get = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_SET_FUNC_NAME) {
				__set = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_UNSET_FUNC_NAME) {
				__unset = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_ISSET_FUNC_NAME) {
				__isset = reg_function
				scope.SetIsUseGuards(true)
			} else if ZendStringEqualsLiteral(lowercase_name, ZEND_DEBUGINFO_FUNC_NAME) {
				__debugInfo = reg_function
			} else {
				reg_function = nil
			}
			if reg_function != nil {
				ZendCheckMagicMethodImplementation(scope, reg_function, error_type)
			}
		}
		ptr++
		count++
		ZendStringRelease(lowercase_name)
	}
	if unload != 0 {
		if scope != nil {
			Efree((*byte)(lc_class_name))
		}
		for ptr.GetFname() != nil {
			fname_len = strlen(ptr.GetFname())
			lowercase_name = ZendStringAlloc(fname_len, 0)
			ZendStrTolowerCopy(lowercase_name.GetVal(), ptr.GetFname(), fname_len)
			if ZendHashExists(target_function_table, lowercase_name) != 0 {
				ZendError(error_type, "Function registration failed - duplicate name - %s%s%s", b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""), b.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
			ZendStringEfree(lowercase_name)
			ptr++
		}
		ZendUnregisterFunctions(functions, count, target_function_table)
		return FAILURE
	}
	if scope != nil {
		scope.SetConstructor(ctor)
		scope.SetDestructor(dtor)
		scope.SetClone(clone)
		scope.SetCall(__call)
		scope.SetCallstatic(__callstatic)
		scope.SetTostring(__tostring)
		scope.SetGet(__get)
		scope.SetSet(__set)
		scope.SetUnset(__unset)
		scope.SetIsset(__isset)
		scope.SetDebugInfo(__debugInfo)
		scope.SetSerializeFunc(serialize_func)
		scope.SetUnserializeFunc(unserialize_func)
		if ctor != nil {
			ctor.SetIsCtor(true)
			if ctor.IsStatic() {
				ZendError(error_type, "Constructor %s::%s() cannot be static", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
			}
			ctor.SetIsAllowStatic(false)
		}
		if dtor != nil {
			dtor.SetIsDtor(true)
			if dtor.IsStatic() {
				ZendError(error_type, "Destructor %s::%s() cannot be static", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
			}
			dtor.SetIsAllowStatic(false)
		}
		if clone != nil {
			if clone.IsStatic() {
				ZendError(error_type, "%s::%s() cannot be static", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
			}
			clone.SetIsAllowStatic(false)
		}
		if __call != nil {
			if __call.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __call.GetFunctionName().GetVal())
			}
			__call.SetIsAllowStatic(false)
		}
		if __callstatic != nil {
			if !__callstatic.IsStatic() {
				ZendError(error_type, "Method %s::%s() must be static", scope.GetName().GetVal(), __callstatic.GetFunctionName().GetVal())
			}
			__callstatic.SetIsStatic(true)
		}
		if __tostring != nil {
			if __tostring.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __tostring.GetFunctionName().GetVal())
			}
			__tostring.SetIsAllowStatic(false)
		}
		if __get != nil {
			if __get.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __get.GetFunctionName().GetVal())
			}
			__get.SetIsAllowStatic(false)
		}
		if __set != nil {
			if __set.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __set.GetFunctionName().GetVal())
			}
			__set.SetIsAllowStatic(false)
		}
		if __unset != nil {
			if __unset.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __unset.GetFunctionName().GetVal())
			}
			__unset.SetIsAllowStatic(false)
		}
		if __isset != nil {
			if __isset.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __isset.GetFunctionName().GetVal())
			}
			__isset.SetIsAllowStatic(false)
		}
		if __debugInfo != nil {
			if __debugInfo.IsStatic() {
				ZendError(error_type, "Method %s::%s() cannot be static", scope.GetName().GetVal(), __debugInfo.GetFunctionName().GetVal())
			}
		}
		if ctor != nil && ctor.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "Constructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), ctor.GetFunctionName().GetVal())
		}
		if dtor != nil && dtor.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "Destructor %s::%s() cannot declare a return type", scope.GetName().GetVal(), dtor.GetFunctionName().GetVal())
		}
		if clone != nil && clone.IsHasReturnType() {
			ZendErrorNoreturn(E_CORE_ERROR, "%s::%s() cannot declare a return type", scope.GetName().GetVal(), clone.GetFunctionName().GetVal())
		}
		Efree((*byte)(lc_class_name))
	}
	return SUCCESS
}
