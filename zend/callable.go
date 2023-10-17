package zend

import (
	"errors"
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

type Callable struct {
	name   string
	object *types.Object
	zv     *types.Zval
	fcc    types.ZendFcallInfoCache
}

func TryAsCallable(callable *types.Zval, object *types.Object, checkFlags uint32) (call *Callable, err error) {
	var fcc *types.ZendFcallInfoCache
	var errStr **byte
	var ret bool = zendIsCallableImpl(callable, object, checkFlags, fcc, errStr)
	if errStr != nil {
		err = errors.New(b.CastStrAuto(*errStr))
	}
	if ret {
		call = &Callable{
			name:   GetCallableName(callable, object),
			object: object,
			zv:     callable,
		}
	}
	return
}

func createMethodString(className string, methodName string) string {
	return className + "::" + methodName
}

func GetCallableName(callable *types.Zval, object *types.Object) string {
	callable = callable.DeRef()
	switch callable.Type() {
	case types.IsString:
		if object != nil {
			return createMethodString(object.GetCe().Name(), callable.String())
		}
		return callable.String()
	case types.IsArray:
		var method *types.Zval = nil
		var obj *types.Zval = nil
		if callable.Array().Len() == 2 {
			obj = types.ZendHashIndexFindDeref(callable.Array(), 0)
			method = types.ZendHashIndexFindDeref(callable.Array(), 1)
		}
		if obj == nil || method == nil || !method.IsString() {
			return "Array"
		}
		if obj.IsString() {
			return createMethodString(obj.String(), method.String())
		} else if obj.IsObject() {
			return createMethodString(obj.Object().GetCe().Name(), method.String())
		} else {
			return "Array"
		}
	case types.IsObject:
		var calling_scope *types.ClassEntry
		var fptr types.IFunction
		var object *types.Object
		if callable.Object().CanGetClosure() && callable.Object().GetClosure(callable, &calling_scope, &fptr, &object) == types.SUCCESS {
			var ce *types.ClassEntry = callable.Object().GetCe()
			return ce.Name() + "::__invoke"
		}
		return operators.ZvalGetStrVal(callable)
	default:
		return operators.ZvalGetStrVal(callable)
	}
}

func IsCallable(callable *types.Zval, object *types.Object, checkFlags uint32) bool {
	return zendIsCallableImpl(callable, object, checkFlags, nil, nil)
}

func ZendIsCallableEx(callable *types.Zval, object *types.Object, checkFlags uint32, callableName **types.String, fcc *types.ZendFcallInfoCache, error **byte) bool {
	var ret bool = zendIsCallableImpl(callable, object, checkFlags, fcc, error)
	if callableName != nil {
		*callableName = types.NewString(GetCallableName(callable, object))
	}
	return ret
}

func zendIsCallableImpl(callable *types.Zval, object *types.Object, checkFlags uint32, fcc *types.ZendFcallInfoCache, error **byte) bool {
	call := &Callable{
		zv:     callable,
		object: object,
	}
	ret, err := isCallableImpl(call, checkFlags, error == nil)
	if err != nil && error != nil {
		*error = b.CastStrPtr(err.Error())
	}
	return ret
}

func isCallableImpl(call *Callable, checkFlags uint32, throw bool) (bool, error) {
	callable := call.zv
	object := call.object
	fcc := &call.fcc
	checkSyntaxOnly := (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0

	var fccLocal types.ZendFcallInfoCache
	defer func() {
		if fcc == &fccLocal {
			zendReleaseFcallInfoCache(fcc)
		}
	}()

	fcc.SetCallingScope(nil)
	fcc.SetCalledScope(nil)
	fcc.SetFunctionHandler(nil)
	fcc.SetObject(nil)

	callable = callable.DeRef()
	switch callable.Type() {
	case types.IsString:
		if object != nil {
			fcc.SetObject(object)
			fcc.SetCallingScope(object.GetCe())
		}
		if checkSyntaxOnly {
			fcc.SetCalledScope(fcc.GetCallingScope())
			return true, nil
		}
		return isCallableCheckFunc(checkFlags, callable.String(), fcc, false, throw)
	case types.IsArray:
		if callable.Array().Len() != 2 {
			return false, errors.New("array must have exactly two members")
		}

		var method = callable.Array().IndexFind(0).SafeDeRef()
		var obj = callable.Array().IndexFind(1).SafeDeRef()

		if obj != nil && method != nil && method.IsString() {
			if obj.IsString() {
				if checkSyntaxOnly {
					return true, nil
				}
				var strictClass bool
				if err := isCallableCheckClass(obj.String(), ZendGetExecutedScope(), fcc, &strictClass); err != nil {
					return isCallableCheckFunc(checkFlags, method.String(), fcc, strictClass, throw)
				} else {
					return false, err
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(types.Z_OBJCE_P(obj))
				fcc.SetObject(obj.Object())
				if checkSyntaxOnly {
					fcc.SetCalledScope(fcc.GetCallingScope())
					return true, nil
				}
				return isCallableCheckFunc(checkFlags, method.String(), fcc, false, throw)
			}
		}

		if obj == nil || (!obj.IsString() && !obj.IsObject()) {
			return false, errors.New("first array member is not a valid class name or object")
		} else {
			return false, errors.New("second array member is not a valid method")
		}
	case types.IsObject:
		if callable.Object().CanGetClosure() {
			if callable.Object().GetClosure(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == types.SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fccLocal {
					zendReleaseFcallInfoCache(fcc)
				}
				return true, nil
			} else {
				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */
				zend.EG__().ClearException()
			}
		}
		return false, errors.New("no array or string given")
	default:
		return false, errors.New("no array or string given")
	}
}

func isCallableCheckFunc(checkFlags uint32, callableName string, fcc *types.ZendFcallInfoCache, strictClass bool, throw bool) (bool, error) {
	var ceOrg *types.ClassEntry = fcc.GetCallingScope()
	fcc.SetCallingScope(nil)
	defer func() {
		if fcc.GetObject() != nil {
			fcc.SetCalledScope(fcc.GetObject().GetCe())
			if fcc.GetFunctionHandler() != nil && fcc.GetFunctionHandler().IsStatic() {
				fcc.SetObject(nil)
			}
		}
	}()

	// prepare flags
	checkNoAccess := checkFlags&IS_CALLABLE_CHECK_NO_ACCESS != 0
	checkIsStatic := checkFlags&IS_CALLABLE_CHECK_IS_STATIC != 0
	checkSilent := checkFlags&IS_CALLABLE_CHECK_SILENT != 0

	var retval int = 0
	var ftable FunctionTable
	var callViaHandler = false
	var scope *types.ClassEntry

	if ceOrg == nil {
		var func_ types.IFunction

		/* Check if function with given name exists.
		 * This may be a compound name that includes namespace name */
		if callableName != "" && callableName[0] == '\\' {
			/* Skip leading \ */
			lmname := ascii.StrToLower(callableName[1:])
			func_ = ZendFetchFunctionStr(lmname)
		} else {
			lmname := callableName
			func_ = ZendFetchFunctionStr(lmname)
			if func_ == nil {
				lmname = ascii.StrToLower(callableName)
				func_ = ZendFetchFunctionStr(lmname)
			}
		}
		if func_ != nil {
			fcc.SetFunctionHandler(func_)
			return true, nil
		}
	}

	/* Split name into class/namespace and method/function names */
	var methodName string
	if colonIdx := strings.LastIndexByte(callableName, ':'); colonIdx > 0 && callableName[colonIdx-1] == ':' {
		if colonIdx == 1 {
			return false, errors.New("invalid function name")
		}

		className := callableName[:colonIdx-1]
		methodName = callableName[colonIdx+1:]

		/* This is a compound name.
		 * Try to fetch class and then find static method. */
		if ceOrg != nil {
			scope = ceOrg
		} else {
			scope = ZendGetExecutedScope()
		}
		if err := isCallableCheckClass(className, scope, fcc, &strictClass); err != nil {
			return false, err
		}

		ftable = fcc.GetCallingScope().FunctionTable()
		if ceOrg != nil && !operators.InstanceofFunction(ceOrg, fcc.GetCallingScope()) {
			return false, fmt.Errorf("class '%s' is not a subclass of '%s'", ceOrg.Name(), fcc.GetCallingScope().Name())
		}
	} else if ceOrg != nil {
		/* Try to fetch find static method of given class. */
		methodName = callableName
		ftable = ceOrg.FunctionTable()
		fcc.SetCallingScope(ceOrg)
	} else {
		/* We already checked for plain function before. */
		return false, fmt.Errorf("function '%s' not found or invalid function name", callableName)
	}

	var gotoGetFunctionViaHandler bool = false

	lcMethodName := ascii.StrToLower(methodName)
	if strictClass && fcc.GetCallingScope() != nil && lcMethodName == ZEND_CONSTRUCTOR_FUNC_NAME {
		fcc.SetFunctionHandler(fcc.GetCallingScope().GetConstructor())
		if fcc.GetFunctionHandler() != nil {
			retval = 1
		}
	} else if zif := ftable.Get(lcMethodName); zif != nil {
		fcc.SetFunctionHandler(zif)
		retval = 1
		if fcc.GetFunctionHandler().GetOpArray().IsChanged() && !strictClass {
			scope = ZendGetExecutedScope()
			if scope != nil && operators.InstanceofFunction(fcc.GetFunctionHandler().GetScope(), scope) {
				if priv_fbc := scope.FunctionTable().Get(lcMethodName); priv_fbc != nil {
					if priv_fbc.IsPrivate() && priv_fbc.GetScope() == scope {
						fcc.SetFunctionHandler(priv_fbc)
					}
				}
			}
		}
		if !fcc.GetFunctionHandler().IsPublic() && !checkNoAccess && (fcc.GetCallingScope() != nil && (fcc.GetObject() != nil && fcc.GetCallingScope().GetCall() != nil || fcc.GetObject() == nil && fcc.GetCallingScope().GetCallstatic() != nil)) {
			scope = ZendGetExecutedScope()
			if fcc.GetFunctionHandler().GetScope() != scope {
				if fcc.GetFunctionHandler().IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) {
					retval = 0
					fcc.SetFunctionHandler(nil)
					gotoGetFunctionViaHandler = true
				}
			}
		}
	} else {
		gotoGetFunctionViaHandler = true
	}

	// label get_function_via_handler
	if gotoGetFunctionViaHandler {
		if fcc.GetObject() != nil && fcc.GetCallingScope() == ceOrg {
			if strictClass && ceOrg.GetCall() != nil {
				fcc.SetFunctionHandler(ZendGetCallTrampolineFunc(ceOrg, methodName, false))
				callViaHandler = true
				retval = 1
			} else {
				fcc.SetFunctionHandler(fcc.GetObject().GetMethod(methodName, nil))
				if fcc.GetFunctionHandler() != nil {
					if strictClass && (fcc.GetFunctionHandler().GetScope() == nil || !operators.InstanceofFunction(ceOrg, fcc.GetFunctionHandler().GetScope())) {
						zendReleaseFcallInfoCache(fcc)
					} else {
						retval = 1
						callViaHandler = fcc.GetFunctionHandler().IsCallViaTrampoline()
					}
				}
			}
		} else if fcc.GetCallingScope() != nil {
			if fcc.GetCallingScope().GetGetStaticMethod() != nil {
				fcc.SetFunctionHandler(fcc.GetCallingScope().GetGetStaticMethod()(fcc.GetCallingScope(), types.NewString(methodName)))
			} else {
				fcc.SetFunctionHandler(ZendStdGetStaticMethod(fcc.GetCallingScope(), types.NewString(methodName), nil))
			}
			if fcc.GetFunctionHandler() != nil {
				retval = 1
				callViaHandler = fcc.GetFunctionHandler().IsCallViaTrampoline()
				if callViaHandler && fcc.GetObject() == nil {
					var object *types.Object = ZendGetThisObject(CurrEX())
					if object != nil && operators.InstanceofFunction(object.GetCe(), fcc.GetCallingScope()) {
						fcc.SetObject(object)
					}
				}
			}
		}
	}

	var err error
	if retval != 0 {
		if fcc.GetCallingScope() != nil && !callViaHandler {
			if fcc.GetFunctionHandler().IsAbstract() {
				return false, fmt.Errorf("cannot call abstract method %s::%s()", fcc.GetCallingScope().Name(), fcc.GetFunctionHandler().FunctionName())
			} else if fcc.GetObject() == nil && !fcc.GetFunctionHandler().IsStatic() {
				var severity int
				var verb string
				if fcc.GetFunctionHandler().IsAllowStatic() {
					severity = faults.E_DEPRECATED
					verb = "should not"
				} else {
					/* An internal function assumes $this is present and won't check that. So PHP would crash by allowing the call. */
					severity = faults.E_ERROR
					verb = "cannot"
				}
				if checkIsStatic {
					retval = 0
				}
				if !throw {
					err = fmt.Errorf("non-static method %s::%s() %s be called statically", fcc.GetCallingScope().Name(), fcc.GetFunctionHandler().FunctionName(), verb)
					if severity != faults.E_DEPRECATED {
						retval = 0
					}
				} else if retval != 0 {
					if severity == faults.E_ERROR {
						faults.ThrowError(nil, "Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().Name(), fcc.GetFunctionHandler().FunctionName(), verb)
					} else {
						faults.Error(severity, fmt.Sprintf("Non-static method %s::%s() %s be called statically", fcc.GetCallingScope().Name(), fcc.GetFunctionHandler().FunctionName(), verb))
					}
				}
			}
			if retval != 0 && !fcc.GetFunctionHandler().IsPublic() && !checkNoAccess {
				scope = ZendGetExecutedScope()
				if fcc.GetFunctionHandler().GetScope() != scope {
					if fcc.GetFunctionHandler().IsPrivate() || !ZendCheckProtected(ZendGetFunctionRootClass(fcc.GetFunctionHandler()), scope) {
						return false, fmt.Errorf("cannot access %s method %s::%s()", ZendVisibilityString(fcc.GetFunctionHandler().GetFnFlags()), fcc.GetCallingScope().Name(), fcc.GetFunctionHandler().FunctionName())
					}
				}
			}
		}
	} else if !checkSilent {
		if fcc.GetCallingScope() != nil {
			return false, fmt.Errorf("class '%s' does not have a method '%s'", fcc.GetCallingScope().Name(), methodName)
		} else {
			return false, fmt.Errorf("function '%s' does not exist", methodName)
		}
	}
	return retval != 0, err
}

func isCallableCheckClass(name string, scope *types.ClassEntry, fcc *types.ZendFcallInfoCache, strictClassPtr *bool) error {
	// strictClass
	strictClass := false
	if strictClassPtr != nil {
		defer func() {
			*strictClassPtr = strictClass
		}()
	}

	lcName := ascii.StrToLower(name)
	if lcName == "self" {
		if scope == nil {
			return errors.New("cannot access self:: when no class scope is active")
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(CurrEX()))
			if fcc.GetCalledScope() == nil || !operators.InstanceofFunction(fcc.GetCalledScope(), scope) {
				fcc.SetCalledScope(scope)
			}
			fcc.SetCallingScope(scope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(CurrEX()))
			}
			return nil
		}
	} else if lcName == "parent" {
		if scope == nil {
			return errors.New("cannot access parent:: when no class scope is active")
		} else if scope.GetParent() == nil {
			return errors.New("cannot access parent:: when current class scope has no parent")
		} else {
			fcc.SetCalledScope(ZendGetCalledScope(CurrEX()))
			if fcc.GetCalledScope() == nil || !operators.InstanceofFunction(fcc.GetCalledScope(), scope.GetParent()) {
				fcc.SetCalledScope(scope.GetParent())
			}
			fcc.SetCallingScope(scope.GetParent())
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(CurrEX()))
			}
			strictClass = true
			return nil
		}
	} else if lcName == "static" {
		var calledScope *types.ClassEntry = ZendGetCalledScope(CurrEX())
		if calledScope == nil {
			return errors.New("cannot access static:: when no class scope is active")
		} else {
			fcc.SetCalledScope(calledScope)
			fcc.SetCallingScope(calledScope)
			if fcc.GetObject() == nil {
				fcc.SetObject(ZendGetThisObject(CurrEX()))
			}
			strictClass = true
			return nil
		}
	} else if ce := ZendLookupClass(name); ce != nil {
		var ceScope *types.ClassEntry
		var ex *ZendExecuteData = CurrEX()
		for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
			ex = ex.GetPrevExecuteData()
		}
		if ex != nil {
			ceScope = ex.GetFunc().GetScope()
		} else {
			ceScope = nil
		}
		fcc.SetCallingScope(ce)
		if ceScope != nil && fcc.GetObject() == nil {
			var object *types.Object = ZendGetThisObject(CurrEX())
			if object != nil && operators.InstanceofFunction(object.GetCe(), ceScope) && operators.InstanceofFunction(ceScope, ce) {
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
		strictClass = true
		return nil
	} else {
		return fmt.Errorf("class '%s' not found", name)
	}
}
func zendReleaseFcallInfoCache(fcc *types.ZendFcallInfoCache) {
	if fcc.GetFunctionHandler() != nil && (fcc.GetFunctionHandler().IsCallViaTrampoline() || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY || fcc.GetFunctionHandler().GetType() == ZEND_OVERLOADED_FUNCTION) {
		ZendFreeTrampoline(fcc.GetFunctionHandler())
	}
	fcc.SetFunctionHandler(nil)
}
