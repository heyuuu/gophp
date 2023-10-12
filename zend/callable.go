package zend

import (
	"errors"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
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
	ret, err := isCallableImpl(call, checkFlags)
	if err != nil && error != nil {
		*error = b.CastStrPtr(err.Error())
	}
	return ret
}

func isCallableImpl(call *Callable, checkFlags uint32) (bool, error) {
	zendIsCallableImplEx(call, checkFlags, nil)
}

func zendIsCallableImplEx(call *Callable, checkFlags uint32, error **byte) bool {
	callable := call.zv
	object := call.object
	fcc := &call.fcc

	var ret bool
	var fcc_local types.ZendFcallInfoCache
	var strict_class int = 0
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
		if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
			fcc.SetCalledScope(fcc.GetCallingScope())
			return true
		}
	check_func:
		ret = ZendIsCallableCheckFunc(checkFlags, callable, fcc, strict_class, error)
		if fcc == &fcc_local {
			ZendReleaseFcallInfoCache(fcc)
		}
		return ret
	case types.IsArray:
		var method *types.Zval = nil
		var obj *types.Zval = nil
		if callable.Array().Len() == 2 {
			obj = callable.Array().IndexFind(0)
			method = callable.Array().IndexFind(1)
		}
		for {
			if obj == nil || method == nil {
				break
			}
			method = types.ZVAL_DEREF(method)
			if !method.IsString() {
				break
			}
			obj = types.ZVAL_DEREF(obj)
			if obj.IsString() {
				if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					return 1
				}
				if ZendIsCallableCheckClass(obj.StringEx(), ZendGetExecutedScope(), fcc, &strict_class, error) == 0 {
					return 0
				}
			} else if obj.IsObject() {
				fcc.SetCallingScope(types.Z_OBJCE_P(obj))
				fcc.SetObject(obj.Object())
				if (checkFlags & IS_CALLABLE_CHECK_SYNTAX_ONLY) != 0 {
					fcc.SetCalledScope(fcc.GetCallingScope())
					return 1
				}
			} else {
				break
			}
			callable = method
			goto check_func
			break
		}
		if callable.Array().Len() == 2 {
			if obj == nil || lang.CondF(!(obj.IsRef()), func() bool { return !obj.IsString() && !obj.IsObject() }, func() bool {
				return types.Z_REFVAL_P(obj).Type() != types.IsString && types.Z_REFVAL_P(obj).Type() != types.IsObject
			}) {
				if error != nil {
					*error = Estrdup("first array member is not a valid class name or object")
				}
			} else {
				if error != nil {
					*error = Estrdup("second array member is not a valid method")
				}
			}
		} else {
			if error != nil {
				*error = Estrdup("array must have exactly two members")
			}
		}
		return 0
	case types.IsObject:
		if callable.Object().CanGetClosure() {
			if callable.Object().GetClosure(callable, fcc.GetCallingScope(), fcc.GetFunctionHandler(), fcc.GetObject()) == types.SUCCESS {
				fcc.SetCalledScope(fcc.GetCallingScope())
				if fcc == &fcc_local {
					ZendReleaseFcallInfoCache(fcc)
				}
				return 1
			} else {

				/* Discard exceptions thrown from Z_OBJ_HANDLER_P(callable, get_closure) */

				faults.ClearException()
			}
		}
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return false
	default:
		if error != nil {
			*error = Estrdup("no array or string given")
		}
		return false
	}
}
