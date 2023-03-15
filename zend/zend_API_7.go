// <<generate>>

package zend

import (
	b "sik/builtin"
)

func DisplayDisabledClass(class_type *ZendClassEntry) *ZendObject {
	var intern *ZendObject
	intern = ZendObjectsNew(class_type)

	/* Initialize default properties */

	if class_type.GetDefaultPropertiesCount() != 0 {
		var p *Zval = intern.GetPropertiesTable()
		var end *Zval = p + class_type.GetDefaultPropertiesCount()
		for {
			p.SetUndef()
			p++
			if p == end {
				break
			}
		}
	}
	ZendError(E_WARNING, "%s() has been disabled for security reasons", class_type.GetName().GetVal())
	return intern
}
func ZendDisableClass(class_name *byte, class_name_length int) int {
	var disabled_class *ZendClassEntry
	var key *ZendString
	var fn *ZendFunction
	key = ZendStringAlloc(class_name_length, 0)
	ZendStrTolowerCopy(key.GetVal(), class_name, class_name_length)
	disabled_class = ZendHashFindPtr(CG__().GetClassTable(), key)
	ZendStringReleaseEx(key, 0)
	if disabled_class == nil {
		return FAILURE
	}
	disabled_class.InitMethods(DisabledClassNew)
	disabled_class.SetCreateObject(DisplayDisabledClass)
	var __ht *HashTable = disabled_class.GetFunctionTable()
	for _, _p := range __ht.foreachData() {
		var _z *Zval = _p.GetVal()

		fn = _z.GetPtr()
		if fn.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE|ZEND_ACC_HAS_TYPE_HINTS) && fn.GetScope() == disabled_class {
			ZendFreeInternalArgInfo(fn.GetInternalFunction())
		}
	}
	disabled_class.GetFunctionTable().Clean()
	return SUCCESS
}
func ZendIsCallableCheckClass(name *ZendString, scope *ZendClassEntry, fcc *ZendFcallInfoCache, strict_class *int, error **byte) int {
	var ret int = 0
	var ce *ZendClassEntry
	var name_len int = name.GetLen()
	var lcname *ZendString
	ZSTR_ALLOCA_ALLOC(lcname, name_len, use_heap)
	ZendStrTolowerCopy(lcname.GetVal(), name.GetVal(), name_len)
	*strict_class = 0
	if ZendStringEqualsLiteral(lcname, "self") {
		if scope == nil {
			if error != nil {
				*error = Estrdup("cannot access self:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(EG__().GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope) == 0 {
				fcc.SetCalledScope(scope)
			}
			fcc.SetCallingScope(scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG__().GetCurrentExecuteData()))
			}
			ret = 1
		}
	} else if ZendStringEqualsLiteral(lcname, "parent") {
		if scope == nil {
			if error != nil {
				*error = Estrdup("cannot access parent:: when no class scope is active")
			}
		} else if !(scope.GetParent()) {
			if error != nil {
				*error = Estrdup("cannot access parent:: when current class scope has no parent")
			}
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(EG__().GetCurrentExecuteData()))
			if fcc.GetCalledScope() == nil || InstanceofFunction(fcc.GetCalledScope(), scope.GetParent()) == 0 {
				fcc.SetCalledScope(scope.GetParent())
			}
			fcc.SetCallingScope(scope.GetParent())
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG__().GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if ZendStringEqualsLiteral(lcname, "static") {
		var called_scope *ZendClassEntry = ZendGetCalledScope(EG__().GetCurrentExecuteData())
		if called_scope == nil {
			if error != nil {
				*error = Estrdup("cannot access static:: when no class scope is active")
			}
		} else {
			fcc.SetCalledScope(called_scope)
			fcc.SetCallingScope(called_scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(EG__().GetCurrentExecuteData()))
			}
			*strict_class = 1
			ret = 1
		}
	} else if b.Assign(&ce, ZendLookupClass(name)) != nil {
		var scope *ZendClassEntry
		var ex *ZendExecuteData = EG__().GetCurrentExecuteData()
		for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
			ex = ex.GetPrevExecuteData()
		}
		if ex != nil {
			scope = ex.GetFunc().GetScope()
		} else {
			scope = nil
		}
		fcc.SetCallingScope(ce)
		if scope != nil && fcc.GetObject() == nil {
			var object *ZendObject = ZendGetThisObject(EG__().GetCurrentExecuteData())
			if object != nil && InstanceofFunction(object.GetCe(), scope) != 0 && InstanceofFunction(scope, ce) != 0 {
				fcc.SetObject(object)
				fcc.SetCalledScope(object.GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		} else {
			if fcc.GetObject() != nil {
				fcc.SetCalledScope(fcc.GetObject().GetCe())
			} else {
				fcc.SetCalledScope(ce)
			}
		}
		*strict_class = 1
		ret = 1
	} else {
		if error != nil {
			ZendSpprintf(error, 0, "class '%.*s' not found", int(name_len), name.GetVal())
		}
	}
	lcname.Free()
	return ret
}
func ZendReleaseFcallInfoCache(fcc *ZendFcallInfoCache) {
	if fcc.GetFunctionHandler() != nil && (fcc.GetFunctionHandler().IsCallViaTrampoline() || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION) {
		if fcc.GetFunctionHandler().GetType() != ZEND_OVERLOADED_FUNCTION && fcc.GetFunctionHandler().GetFunctionName() != nil {
			ZendStringReleaseEx(fcc.GetFunctionHandler().GetFunctionName(), 0)
		}
		ZendFreeTrampoline(fcc.GetFunctionHandler())
	}
	fcc.SetFunctionHandler(nil)
}
func ZendIsCallableCheckFunc(check_flags int, callable *Zval, fcc *ZendFcallInfoCache, strict_class int, error **byte) int {
	var ce_org *ZendClassEntry = fcc.GetCallingScope()
	var retval int = 0
	var mname *ZendString
	var cname *ZendString
	var lmname *ZendString
	var colon *byte
	var clen int
	var ftable *HashTable
	var call_via_handler int = 0
	var scope *ZendClassEntry
	var zv *Zval
	fcc.SetCallingScope(nil)
	if ce_org == nil {
		var func_ *ZendFunction
		var lmname *ZendString

		/* Check if function with given name exists.
		 * This may be a compound name that includes namespace name */

		if Z_STRVAL_P(callable)[0] == '\\' {

			/* Skip leading \ */

			ZSTR_ALLOCA_ALLOC(lmname, Z_STRLEN_P(callable)-1, use_heap)
			ZendStrTolowerCopy(lmname.GetVal(), Z_STRVAL_P(callable)+1, Z_STRLEN_P(callable)-1)
			func_ = ZendFetchFunction(lmname)
			lmname.Free()
		} else {
			lmname = callable.GetStr()
			func_ = ZendFetchFunction(lmname)
			if func_ == nil {
				ZSTR_ALLOCA_ALLOC(lmname, Z_STRLEN_P(callable), use_heap)
				ZendStrTolowerCopy(lmname.GetVal(), Z_STRVAL_P(callable), Z_STRLEN_P(callable))
				func_ = ZendFetchFunction(lmname)
				lmname.Free()
			}
		}
		if func_ != nil {
			fcc.SetFunctionHandler(func_)
			return 1
		}
	}

	/* Split name into class/namespace and method/function names */

	if b.Assign(&colon, ZendMemrchr(Z_STRVAL_P(callable), ':', Z_STRLEN_P(callable))) != nil && colon > Z_STRVAL_P(callable) && (*(colon - 1)) == ':' {
		var mlen int
		colon--
		clen = colon - Z_STRVAL_P(callable)
		mlen = Z_STRLEN_P(callable) - clen - 2
		if colon == Z_STRVAL_P(callable) {
			if error != nil {
				*error = Estrdup("invalid function name")
			}
			return 0
		}

		/* This is a compound name.
		 * Try to fetch class and then find static method. */

		if ce_org != nil {
			scope = ce_org
		} else {
			scope = ZendGetExecutedScope()
		}
		cname = ZendStringInit(Z_STRVAL_P(callable), clen, 0)
		if ZendIsCallableCheckClass(cname, scope, fcc, &strict_class, error) == 0 {
			ZendStringReleaseEx(cname, 0)
			return 0
		}
		ZendStringReleaseEx(cname, 0)
		ftable = fcc.GetCallingScope().GetFunctionTable()
		if ce_org != nil && InstanceofFunction(ce_org, fcc.GetCallingScope()) == 0 {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' is not a subclass of '%s'", ce_org.GetName().GetVal(), fcc.GetCallingScope().GetName().GetVal())
			}
			return 0
		}
		mname = ZendStringInit(Z_STRVAL_P(callable)+clen+2, mlen, 0)
	} else if ce_org != nil {

		/* Try to fetch find static method of given class. */

		mname = callable.GetStr()
		mname.AddRefcount()
		ftable = ce_org.GetFunctionTable()
		fcc.SetCallingScope(ce_org)
	} else {

		/* We already checked for plain function before. */

		if error != nil && (check_flags&IS_CALLABLE_CHECK_SILENT) == 0 {
			ZendSpprintf(error, 0, "function '%s' not found or invalid function name", Z_STRVAL_P(callable))
		}
		return 0
	}
	lmname = ZendStringTolower(mname)
	if strict_class != 0 && fcc.GetCallingScope() != nil && ZendStringEqualsLiteral(lmname, ZEND_CONSTRUCTOR_FUNC_NAME) {
		fcc.SetFunctionHandler(fcc.GetCallingScope().GetConstructor())
		if fcc.GetFunctionHandler() != nil {
			retval = 1
		}
	} else if b.Assign(&zv, ftable.KeyFind(lmname.GetStr())) != nil {
		fcc.SetFunctionHandler(zv.GetPtr())
		retval = 1
		if fcc.GetFunctionHandler().GetOpArray().IsChanged() && strict_class == 0 {
			scope = ZendGetExecutedScope()
			if scope != nil && InstanceofFunction(fcc.GetFunctionHandler().GetScope(), scope) != 0 {
				zv = scope.GetFunctionTable().KeyFind(lmname.GetStr())
				if zv != nil {
					var priv_fbc *ZendFunction = zv.GetPtr()
					if priv_fbc.IsPrivate() && priv_fbc.GetScope() == scope {
						fcc.SetFunctionHandler(priv_fbc)
					}
				}
			}
		}
		if !fcc.GetFunctionHandler().IsPublic() && (check_flags&IS_CALLABLE_CHECK_NO_ACCESS) == 0 && (fcc.GetCallingScope() != nil && (fcc.GetObject() != nil && fcc.GetCallingScope().GetCall() != nil || fcc.GetObject() == nil && fcc.GetCallingScope().GetCallstatic() != nil)) {
			scope = ZendGetExecutedScope()
			if fcc.GetFunctionHandler().GetScope() != scope {
				if fcc.GetFunctionHandler().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) == 0 {
					retval = 0
					fcc.SetFunctionHandler(nil)
					goto get_function_via_handler
				}
			}
		}
	} else {
	get_function_via_handler:
		if fcc.GetObject() != nil && fcc.GetCallingScope() == ce_org {
			if strict_class != 0 && ce_org.GetCall() != nil {
				fcc.SetFunctionHandler(ZendGetCallTrampolineFunc(ce_org, mname, 0))
				call_via_handler = 1
				retval = 1
			} else {
				fcc.SetFunctionHandler(fcc.GetObject().GetHandlers().GetGetMethod()(fcc.GetObject(), mname, nil))
				if fcc.GetFunctionHandler() != nil {
					if strict_class != 0 && (fcc.GetFunctionHandler().GetScope() == nil || InstanceofFunction(ce_org, fcc.GetFunctionHandler().GetScope()) == 0) {
						ZendReleaseFcallInfoCache(fcc)
					} else {
						retval = 1
						call_via_handler = fcc.GetFunctionHandler().IsCallViaTrampoline()
					}
				}
			}
		} else if fcc.GetCallingScope() != nil {
			if fcc.GetCallingScope().GetGetStaticMethod() != nil {
				fcc.SetFunctionHandler(fcc.GetCallingScope().GetGetStaticMethod()(fcc.GetCallingScope(), mname))
			} else {
				fcc.SetFunctionHandler(ZendStdGetStaticMethod(fcc.GetCallingScope(), mname, nil))
			}
			if fcc.GetFunctionHandler() != nil {
				retval = 1
				call_via_handler = fcc.GetFunctionHandler().IsCallViaTrampoline()
				if call_via_handler != 0 && fcc.GetObject() == nil {
					var object *ZendObject = ZendGetThisObject(EG__().GetCurrentExecuteData())
					if object != nil && InstanceofFunction(object.GetCe(), fcc.GetCallingScope()) != 0 {
						fcc.SetObject(object)
					}
				}
			}
		}
	}
	if retval != 0 {
		if fcc.GetCallingScope() != nil && call_via_handler == 0 {
			if fcc.GetFunctionHandler().IsAbstract() {
				retval = 0
				if error != nil {
					ZendSpprintf(error, 0, "cannot call abstract method %s::%s()", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
				}
			} else if fcc.GetObject() == nil && !fcc.GetFunctionHandler().IsStatic() {
				var severity int
				var verb *byte
				if fcc.GetFunctionHandler().IsAllowStatic() {
					severity = E_DEPRECATED
					verb = "should not"
				} else {

					/* An internal function assumes $this is present and won't check that. So PHP would crash by allowing the call. */

					severity = E_ERROR
					verb = "cannot"
				}
				if (check_flags & IS_CALLABLE_CHECK_IS_STATIC) != 0 {
					retval = 0
				}
				if error != nil {
					ZendSpprintf(error, 0, "non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					if severity != E_DEPRECATED {
						retval = 0
					}
				} else if retval != 0 {
					if severity == E_ERROR {
						ZendThrowError(nil, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					} else {
						ZendError(severity, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal(), verb)
					}
				}
			}
			if retval != 0 && !fcc.GetFunctionHandler().IsPublic() && (check_flags&IS_CALLABLE_CHECK_NO_ACCESS) == 0 {
				scope = ZendGetExecutedScope()
				if fcc.GetFunctionHandler().GetScope() != scope {
					if fcc.GetFunctionHandler().IsPrivate() || ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) == 0 {
						if error != nil {
							if (*error) != nil {
								Efree(*error)
							}
							ZendSpprintf(error, 0, "cannot access %s method %s::%s()", ZendVisibilityString(fcc.GetFunctionHandler().GetFnFlags()), fcc.GetCallingScope().GetName().GetVal(), fcc.GetFunctionHandler().GetFunctionName().GetVal())
						}
						retval = 0
					}
				}
			}
		}
	} else if error != nil && (check_flags&IS_CALLABLE_CHECK_SILENT) == 0 {
		if fcc.GetCallingScope() != nil {
			if error != nil {
				ZendSpprintf(error, 0, "class '%s' does not have a method '%s'", fcc.GetCallingScope().GetName().GetVal(), mname.GetVal())
			}
		} else {
			if error != nil {
				ZendSpprintf(error, 0, "function '%s' does not exist", mname.GetVal())
			}
		}
	}
	ZendStringReleaseEx(lmname, 0)
	ZendStringReleaseEx(mname, 0)
	if fcc.GetObject() != nil {
		fcc.SetCalledScope(fcc.GetObject().GetCe())
		if fcc.GetFunctionHandler() != nil && fcc.GetFunctionHandler().IsStatic() {
			fcc.SetObject(nil)
		}
	}
	return retval
}
func ZendCreateMethodString(class_name *ZendString, method_name *ZendString) *ZendString {
	var callable_name *ZendString = ZendStringAlloc(class_name.GetLen()+method_name.GetLen()+b.SizeOf("\"::\"")-1, 0)
	var ptr *byte = callable_name.GetVal()
	memcpy(ptr, class_name.GetVal(), class_name.GetLen())
	ptr += class_name.GetLen()
	memcpy(ptr, "::", b.SizeOf("\"::\"")-1)
	ptr += b.SizeOf("\"::\"") - 1
	memcpy(ptr, method_name.GetVal(), method_name.GetLen()+1)
	return callable_name
}
func ZendGetCallableNameEx(callable *Zval, object *ZendObject) *ZendString {
try_again:
	switch callable.GetType() {
	case IS_STRING:
		if object != nil {
			return ZendCreateMethodString(object.GetCe().GetName(), callable.GetStr())
		}
		return callable.GetStr().Copy()
	case IS_ARRAY:
		var method *Zval = nil
		var obj *Zval = nil
		if Z_ARRVAL_P(callable).GetNNumOfElements() == 2 {
			obj = ZendHashIndexFindDeref(callable.GetArr(), 0)
			method = ZendHashIndexFindDeref(callable.GetArr(), 1)
		}
		if obj == nil || method == nil || method.GetType() != IS_STRING {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
		if obj.IsString() {
			return ZendCreateMethodString(obj.GetStr(), method.GetStr())
		} else if obj.IsObject() {
			return ZendCreateMethodString(Z_OBJCE_P(obj).GetName(), method.GetStr())
		} else {
			return ZSTR_KNOWN(ZEND_STR_ARRAY_CAPITALIZED)
		}
	case IS_OBJECT:
		var calling_scope *ZendClassEntry
		var fptr *ZendFunction
		var object *ZendObject
		if Z_OBJ_HT(*callable).GetGetClosure() != nil && Z_OBJ_HT(*callable).GetGetClosure()(callable, &calling_scope, &fptr, &object) == SUCCESS {
			var ce *ZendClassEntry = Z_OBJCE_P(callable)
			var callable_name *ZendString = ZendStringAlloc(ce.GetName().GetLen()+b.SizeOf("\"::__invoke\"")-1, 0)
			memcpy(callable_name.GetVal(), ce.GetName().GetVal(), ce.GetName().GetLen())
			memcpy(callable_name.GetVal()+ce.GetName().GetLen(), "::__invoke", b.SizeOf("\"::__invoke\""))
			return callable_name
		}
		return ZvalGetString(callable)
	case IS_REFERENCE:
		callable = Z_REFVAL_P(callable)
		goto try_again
	default:
		return ZvalGetStringFunc(callable)
	}
}
func ZendGetCallableName(callable *Zval) *ZendString {
	return ZendGetCallableNameEx(callable, nil)
}
