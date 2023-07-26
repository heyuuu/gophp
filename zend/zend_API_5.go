package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/globals"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

func ZendCollectModuleHandlers() {
	var class_count int = 0

	/* Collect internal classes with static members */
	CG__().ClassTable().Foreach(func(_ string, ce *types.ClassEntry) {
		if ce.IsInternalClass() && ce.GetDefaultStaticMembersCount() > 0 {
			class_count++
		}
	})

	ClassCleanupHandlers = (**types.ClassEntry)(Malloc(b.SizeOf("zend_class_entry *") * (class_count + 1)))
	ClassCleanupHandlers[class_count] = nil
	CG__().ClassTable().Foreach(func(_ string, ce *types.ClassEntry) {
		if ce.IsInternalClass() && ce.GetDefaultStaticMembersCount() > 0 {
			ClassCleanupHandlers[lang.PreDec(&class_count)] = ce
		}
	})
}
func ZendStartupModules() int {
	for _, module := range globals.G().GetSortedModules() {
		if !ZendStartupModuleEx(module) {
			globals.G().DelModule(module.GetName())
		}
	}
	return types.SUCCESS
}
func ZendDestroyModules() {
	Free(ClassCleanupHandlers)
	globals.G().DestroyModules()
}
func ZendRegisterModuleEx(module *ModuleEntry) *ModuleEntry {
	if module == nil {
		return nil
	}

	/* Check module dependencies */
	module = globals.G().RegisterModule(module)
	if module == nil {
		faults.Error(faults.E_CORE_WARNING, "Module '%s' already loaded", module.GetName())
		return nil
	}

	EG__().SetCurrentModule(module)
	if module.GetFunctions() != nil && ZendRegisterFunctions(nil, module.GetFunctions(), nil, module.GetType()) == types.FAILURE {
		globals.G().DelModule(module.GetName())
		EG__().SetCurrentModule(nil)
		faults.Error(faults.E_CORE_WARNING, "%s: Unable to register functions, unable to load", module.GetName())
		return nil
	}
	EG__().SetCurrentModule(nil)
	return module
}
func ZendRegisterInternalModule(module *ModuleEntry) *ModuleEntry {
	module.SetModuleNumber(ZendNextFreeModule())
	return ZendRegisterModuleEx(module)
}
func ZendCheckMagicMethodImplementation(ce *types.ClassEntry, fptr types.IFunction, error_type int) {
	functionName := fptr.FunctionName()
	if !strings.HasPrefix(functionName, "__") {
		return
	}

	lcname := ascii.StrToLower(functionName)
	if lcname == ZEND_DESTRUCTOR_FUNC_NAME && fptr.GetNumArgs() != 0 {
		faults.Error(error_type, "Destructor %s::%s() cannot take arguments", ce.Name(), ZEND_DESTRUCTOR_FUNC_NAME)
	} else if lcname == ZEND_CLONE_FUNC_NAME && fptr.GetNumArgs() != 0 {
		faults.Error(error_type, "Method %s::%s() cannot accept any arguments", ce.Name(), ZEND_CLONE_FUNC_NAME)
	} else if lcname == ZEND_GET_FUNC_NAME {
		if fptr.GetNumArgs() != 1 {
			faults.Error(error_type, "Method %s::%s() must take exactly 1 argument", ce.Name(), ZEND_GET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			faults.Error(error_type, "Method %s::%s() cannot take arguments by reference", ce.Name(), ZEND_GET_FUNC_NAME)
		}
	} else if lcname == ZEND_SET_FUNC_NAME {
		if fptr.GetNumArgs() != 2 {
			faults.Error(error_type, "Method %s::%s() must take exactly 2 arguments", ce.Name(), ZEND_SET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			faults.Error(error_type, "Method %s::%s() cannot take arguments by reference", ce.Name(), ZEND_SET_FUNC_NAME)
		}
	} else if lcname == ZEND_UNSET_FUNC_NAME {
		if fptr.GetNumArgs() != 1 {
			faults.Error(error_type, "Method %s::%s() must take exactly 1 argument", ce.Name(), ZEND_UNSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			faults.Error(error_type, "Method %s::%s() cannot take arguments by reference", ce.Name(), ZEND_UNSET_FUNC_NAME)
		}
	} else if lcname == ZEND_ISSET_FUNC_NAME {
		if fptr.GetNumArgs() != 1 {
			faults.Error(error_type, "Method %s::%s() must take exactly 1 argument", ce.Name(), ZEND_ISSET_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 {
			faults.Error(error_type, "Method %s::%s() cannot take arguments by reference", ce.Name(), ZEND_ISSET_FUNC_NAME)
		}
	} else if lcname == ZEND_CALL_FUNC_NAME {
		if fptr.GetNumArgs() != 2 {
			faults.Error(error_type, "Method %s::%s() must take exactly 2 arguments", ce.Name(), ZEND_CALL_FUNC_NAME)
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			faults.Error(error_type, "Method %s::%s() cannot take arguments by reference", ce.Name(), ZEND_CALL_FUNC_NAME)
		}
	} else if lcname == ZEND_CALLSTATIC_FUNC_NAME {
		if fptr.GetNumArgs() != 2 {
			faults.Error(error_type, "Method %s::__callStatic() must take exactly 2 arguments", ce.Name())
		} else if QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 1) != 0 || QUICK_ARG_SHOULD_BE_SENT_BY_REF(fptr, 2) != 0 {
			faults.Error(error_type, "Method %s::__callStatic() cannot take arguments by reference", ce.Name())
		}
	} else if lcname == ZEND_TOSTRING_FUNC_NAME && fptr.GetNumArgs() != 0 {
		faults.Error(error_type, "Method %s::%s() cannot take arguments", ce.Name(), ZEND_TOSTRING_FUNC_NAME)
	} else if lcname == ZEND_DEBUGINFO_FUNC_NAME && fptr.GetNumArgs() != 0 {
		faults.Error(error_type, "Method %s::%s() cannot take arguments", ce.Name(), ZEND_DEBUGINFO_FUNC_NAME)
	}
}
func ZendRegisterFunctions(scope *types.ClassEntry, functions *types.FunctionEntry, functionTable FunctionTable, type_ int) int {
	var ptr *types.FunctionEntry = functions
	var count int = 0
	var unload int = 0
	var targetFunctionTable FunctionTable = functionTable
	var error_type int
	var ctor types.IFunction = nil
	var dtor types.IFunction = nil
	var clone types.IFunction = nil
	var __get types.IFunction = nil
	var __set types.IFunction = nil
	var __unset types.IFunction = nil
	var __isset types.IFunction = nil
	var __call types.IFunction = nil
	var __callstatic types.IFunction = nil
	var __tostring types.IFunction = nil
	var __debugInfo types.IFunction = nil
	var serialize_func types.IFunction = nil
	var unserialize_func types.IFunction = nil
	var lowercase_name *types.String
	var fname_len int
	var lc_class_name *byte = nil
	var class_name_len int = 0
	if type_ == MODULE_PERSISTENT {
		error_type = faults.E_CORE_WARNING
	} else {
		error_type = faults.E_WARNING
	}
	if targetFunctionTable == nil {
		targetFunctionTable = CG__().FunctionTable()
	}
	var reg_function types.IFunction

	var internal_function = types.NewInternalFunction()
	var function types.IFunction = internal_function

	internal_function.SetModule(EG__().GetCurrentModule())
	if scope != nil {
		class_name_len = scope.GetName().GetLen()
		if lang.Assign(&lc_class_name, operators.ZendMemrchr(scope.Name(), '\\', class_name_len)) {
			lc_class_name++
			class_name_len -= lc_class_name - scope.Name()
			lc_class_name = ascii.StrToLower(b.CastStr(lc_class_name, class_name_len))
		} else {
			lc_class_name = ascii.StrToLower(b.CastStr(scope.Name(), class_name_len))
		}
	}
	for ptr.GetFname() != nil {
		fname_len = strlen(ptr.GetFname())

		internal_function.InitByEntry(ptr)
		internal_function.SetScope(scope)
		if ptr.GetFlags() != 0 {
			if !ptr.IsPppMask() {
				if ptr.GetFlags() != types.AccDeprecated && scope != nil {
					faults.Error(error_type, "Invalid access level for %s%s%s() - access must be exactly one of public, protected or private", lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""), lang.Cond(scope != nil, "::", ""), ptr.GetFname())
				}
				internal_function.SetFnFlags(types.AccPublic | ptr.GetFlags())
			} else {
				internal_function.SetFnFlags(ptr.GetFlags())
			}
		} else {
			internal_function.SetFnFlags(types.AccPublic)
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
						faults.ErrorNoreturn(faults.E_CORE_ERROR, "Cannot declare a return type of %s outside of a class scope", type_name)
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
				faults.Error(error_type, "Static function %s%s%s() cannot be abstract", lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""), lang.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
		} else {
			if scope != nil && scope.IsInterface() {
				Efree((*byte)(lc_class_name))
				faults.Error(error_type, "Interface %s cannot contain non abstract method %s()", scope.Name(), ptr.GetFname())
				return types.FAILURE
			}
			if internal_function.GetHandler() == nil {
				if scope != nil {
					Efree((*byte)(lc_class_name))
				}
				faults.Error(error_type, "Method %s%s%s() cannot be a NULL function", lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""), lang.Cond(scope != nil, "::", ""), ptr.GetFname())
				ZendUnregisterFunctions(functions, count, targetFunctionTable)
				return types.FAILURE
			}
		}
		lowercase_name = types.NewString(ascii.StrToLower(internal_function.FunctionName()))
		// lowercase_name = types.ZendNewInternedString(lowercase_name)
		reg_function = types.CopyFunction(function)
		if !targetFunctionTable.Add(lowercase_name.GetStr(), reg_function) {
			unload = 1
			Free(reg_function)
			// types.ZendStringRelease(lowercase_name)
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
		if reg_function.GetArgInfo() != nil && reg_function.HasFnFlags(types.AccHasReturnType|types.AccHasTypeHints) {

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
					var allow_null bool = 0
					var str *types.String
					if class_name[0] == '?' {
						class_name++
						allow_null = 1
					}
					str = types.NewString(b.CastStrAuto(class_name))
					new_arg_info[i].SetType(types.TypeHintClassName(str, allow_null != 0))
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
			} else if lowercase_name.GetStr() == "serialize" {
				serialize_func = reg_function
			} else if lowercase_name.GetStr() == "unserialize" {
				unserialize_func = reg_function
			} else if lowercase_name.GetStr()[0] != '_' || lowercase_name.GetStr()[1] != '_' {
				reg_function = nil
			} else if lowercase_name.GetStr() == ZEND_CONSTRUCTOR_FUNC_NAME {
				ctor = reg_function
			} else if lowercase_name.GetStr() == ZEND_DESTRUCTOR_FUNC_NAME {
				dtor = reg_function
				if internal_function.GetNumArgs() != 0 {
					faults.Error(error_type, "Destructor %s::%s() cannot take arguments", scope.Name(), ptr.GetFname())
				}
			} else if lowercase_name.GetStr() == ZEND_CLONE_FUNC_NAME {
				clone = reg_function
			} else if lowercase_name.GetStr() == ZEND_CALL_FUNC_NAME {
				__call = reg_function
			} else if lowercase_name.GetStr() == ZEND_CALLSTATIC_FUNC_NAME {
				__callstatic = reg_function
			} else if lowercase_name.GetStr() == ZEND_TOSTRING_FUNC_NAME {
				__tostring = reg_function
			} else if lowercase_name.GetStr() == ZEND_GET_FUNC_NAME {
				__get = reg_function
				scope.SetIsUseGuards(true)
			} else if lowercase_name.GetStr() == ZEND_SET_FUNC_NAME {
				__set = reg_function
				scope.SetIsUseGuards(true)
			} else if lowercase_name.GetStr() == ZEND_UNSET_FUNC_NAME {
				__unset = reg_function
				scope.SetIsUseGuards(true)
			} else if lowercase_name.GetStr() == ZEND_ISSET_FUNC_NAME {
				__isset = reg_function
				scope.SetIsUseGuards(true)
			} else if lowercase_name.GetStr() == ZEND_DEBUGINFO_FUNC_NAME {
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
		// types.ZendStringRelease(lowercase_name)
	}
	if unload != 0 {
		if scope != nil {
			Efree((*byte)(lc_class_name))
		}
		for ptr.GetFname() != nil {
			fname_len = strlen(ptr.GetFname())
			lowercaseName := ascii.StrToLower(b.CastStrAuto(ptr.GetFname()))
			if targetFunctionTable.Exists(lowercaseName) {
				faults.Error(error_type, "Function registration failed - duplicate name - %s%s%s", lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""), lang.Cond(scope != nil, "::", ""), ptr.GetFname())
			}
			ptr++
		}
		ZendUnregisterFunctions(functions, count, targetFunctionTable)
		return types.FAILURE
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
				faults.Error(error_type, "Constructor %s::%s() cannot be static", scope.Name(), ctor.FunctionName())
			}
			ctor.SetIsAllowStatic(false)
		}
		if dtor != nil {
			dtor.SetIsDtor(true)
			if dtor.IsStatic() {
				faults.Error(error_type, "Destructor %s::%s() cannot be static", scope.Name(), dtor.FunctionName())
			}
			dtor.SetIsAllowStatic(false)
		}
		if clone != nil {
			if clone.IsStatic() {
				faults.Error(error_type, "%s::%s() cannot be static", scope.Name(), clone.FunctionName())
			}
			clone.SetIsAllowStatic(false)
		}
		if __call != nil {
			if __call.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __call.FunctionName())
			}
			__call.SetIsAllowStatic(false)
		}
		if __callstatic != nil {
			if !__callstatic.IsStatic() {
				faults.Error(error_type, "Method %s::%s() must be static", scope.Name(), __callstatic.FunctionName())
			}
			__callstatic.SetIsStatic(true)
		}
		if __tostring != nil {
			if __tostring.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __tostring.FunctionName())
			}
			__tostring.SetIsAllowStatic(false)
		}
		if __get != nil {
			if __get.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __get.FunctionName())
			}
			__get.SetIsAllowStatic(false)
		}
		if __set != nil {
			if __set.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __set.FunctionName())
			}
			__set.SetIsAllowStatic(false)
		}
		if __unset != nil {
			if __unset.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __unset.FunctionName())
			}
			__unset.SetIsAllowStatic(false)
		}
		if __isset != nil {
			if __isset.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __isset.FunctionName())
			}
			__isset.SetIsAllowStatic(false)
		}
		if __debugInfo != nil {
			if __debugInfo.IsStatic() {
				faults.Error(error_type, "Method %s::%s() cannot be static", scope.Name(), __debugInfo.FunctionName())
			}
		}
		if ctor != nil && ctor.IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Constructor %s::%s() cannot declare a return type", scope.Name(), ctor.FunctionName())
		}
		if dtor != nil && dtor.IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Destructor %s::%s() cannot declare a return type", scope.Name(), dtor.FunctionName())
		}
		if clone != nil && clone.IsHasReturnType() {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "%s::%s() cannot declare a return type", scope.Name(), clone.FunctionName())
		}
		Efree((*byte)(lc_class_name))
	}
	return types.SUCCESS
}
