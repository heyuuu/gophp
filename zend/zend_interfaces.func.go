package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendCallMethodWith0Params(obj *types.Zval, obj_ce *types.ClassEntry, fn_proxy *types.IFunction, function_name string, retval *types.Zval) *types.Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, retval, 0, nil, nil)
}
func ZendCallMethodWith1Params(
	obj *types.Zval,
	obj_ce *types.ClassEntry,
	fn_proxy *types.IFunction,
	function_name string,
	retval *types.Zval,
	arg1 *types.Zval,
) *types.Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, retval, 1, arg1, nil)
}
func ZendCallMethodWith2Params(
	obj *types.Zval,
	obj_ce *types.ClassEntry,
	fn_proxy *types.IFunction,
	function_name string,
	retval *types.Zval,
	arg1 *types.Zval,
	arg2 *types.Zval,
) *types.Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, retval, 2, arg1, arg2)
}
func ZendCallMethod(
	object *types.Zval,
	obj_ce *types.ClassEntry,
	fn_proxy *types.IFunction,
	function_name string,
	retval_ptr *types.Zval,
	param_count int,
	arg1 *types.Zval,
	arg2 *types.Zval,
) *types.Zval {
	var result int
	var fci types.ZendFcallInfo
	var retval types.Zval
	var params []types.Zval
	if param_count > 0 {
		types.ZVAL_COPY_VALUE(&params[0], arg1)
	}
	if param_count > 1 {
		types.ZVAL_COPY_VALUE(&params[1], arg2)
	}
	fci.SetSize(b.SizeOf("fci"))
	if object != nil {
		fci.SetObject(object.Object())
	} else {
		fci.SetObject(nil)
	}
	if retval_ptr != nil {
		fci.SetRetval(retval_ptr)
	} else {
		fci.SetRetval(&retval)
	}
	fci.SetParamCount(param_count)
	fci.SetParams(params)
	fci.SetNoSeparation(1)
	if fn_proxy == nil && obj_ce == nil {

		/* no interest in caching and no information already present that is
		 * needed later inside zend_call_function. */

		fci.GetFunctionName().SetStringVal(function_name)
		result = ZendCallFunction(&fci, nil)
		// ZvalPtrDtor(fci.GetFunctionName())
	} else {
		var fcic types.ZendFcallInfoCache
		fci.GetFunctionName().SetUndef()
		if obj_ce == nil {
			if object != nil {
				obj_ce = types.Z_OBJCE_P(object)
			} else {
				obj_ce = nil
			}
		}
		if fn_proxy == nil || (*fn_proxy) == nil {
			if obj_ce != nil {
				fcic.SetFunctionHandler(obj_ce.FunctionTable().Get(function_name))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					faults.ErrorNoreturn(faults.E_CORE_ERROR, "Couldn't find implementation for method %s::%s", obj_ce.Name(), function_name)

					/* error at c-level */

				}
			} else {
				fcic.SetFunctionHandler(ZendFetchFunctionStr(function_name))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					faults.ErrorNoreturn(faults.E_CORE_ERROR, "Couldn't find implementation for function %s", function_name)

					/* error at c-level */

				}
			}
			if fn_proxy != nil {
				*fn_proxy = fcic.GetFunctionHandler()
			}
		} else {
			fcic.SetFunctionHandler(*fn_proxy)
		}
		if object != nil {
			fcic.SetCalledScope(types.Z_OBJCE_P(object))
		} else {
			var called_scope *types.ClassEntry = ZendGetCalledScope(CurrEX())
			if obj_ce != nil && (called_scope == nil || operators.InstanceofFunction(called_scope, obj_ce) == 0) {
				fcic.SetCalledScope(obj_ce)
			} else {
				fcic.SetCalledScope(called_scope)
			}
		}
		if object != nil {
			fcic.SetObject(object.Object())
		} else {
			fcic.SetObject(nil)
		}
		result = ZendCallFunction(&fci, &fcic)
	}
	if result == types.FAILURE {

		/* error at c-level */

		if obj_ce == nil {
			if object != nil {
				obj_ce = types.Z_OBJCE_P(object)
			} else {
				obj_ce = nil
			}
		}
		if EG__().GetException() == nil {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, "Couldn't execute method %s%s%s", b.CondF1(obj_ce != nil, func() []byte { return obj_ce.Name() }, ""), b.Cond(obj_ce != nil, "::", ""), function_name)
		}
	}
	if retval_ptr == nil {
		// ZvalPtrDtor(&retval)
		return nil
	}
	return retval_ptr
}
func ZendUserItNewIterator(ce *types.ClassEntry, object *types.Zval, retval *types.Zval) {
	ZendCallMethodWith0Params(object, ce, ce.GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", retval)
}
func ZendUserItInvalidateCurrent(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	if !(iter.GetValue().IsUndef()) {
		// ZvalPtrDtor(iter.GetValue())
		iter.GetValue().SetUndef()
	}
}
func ZendUserItDtor(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *types.Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	// ZvalPtrDtor(object)
}
func ZendUserItValid(_iter *ZendObjectIterator) int {
	if _iter != nil {
		var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
		var object *types.Zval = iter.GetIt().GetData()
		var more types.Zval
		var result int
		ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfValid(), "valid", &more)
		result = operators.IZendIsTrue(&more)
		// ZvalPtrDtor(&more)
		if result != 0 {
			return types.SUCCESS
		} else {
			return types.FAILURE
		}
	}
	return types.FAILURE
}
func ZendUserItGetCurrentData(_iter *ZendObjectIterator) *types.Zval {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *types.Zval = iter.GetIt().GetData()
	if iter.GetValue().IsUndef() {
		ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfCurrent(), "current", iter.GetValue())
	}
	return iter.GetValue()
}
func ZendUserItGetCurrentKey(_iter *ZendObjectIterator, key *types.Zval) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *types.Zval = iter.GetIt().GetData()
	var retval types.Zval
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfKey(), "key", &retval)
	if retval.IsNotUndef() {
		ZVAL_ZVAL(key, &retval, 1, 1)
	} else {
		if EG__().GetException() == nil {
			faults.Error(faults.E_WARNING, "Nothing returned from %s::key()", iter.GetCe().Name())
		}
		key.SetLong(0)
	}
}
func ZendUserItMoveForward(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *types.Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfNext(), "next", nil)
}
func ZendUserItRewind(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *types.Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfRewind(), "rewind", nil)
}
func ZendUserItGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendUserIterator
	if by_ref != 0 {
		faults.ThrowError(nil, "An iterator cannot be used with foreach by reference")
		return nil
	}
	iterator = Emalloc(b.SizeOf("zend_user_iterator"))
	ZendIteratorInit((*ZendObjectIterator)(iterator))
	// 	object.AddRefcount()
	iterator.GetIt().GetData().SetObject(object.Object())
	iterator.GetIt().SetFuncs(&ZendInterfaceIteratorFuncsIterator)
	iterator.SetCe(types.Z_OBJCE_P(object))
	iterator.GetValue().SetUndef()
	return (*ZendObjectIterator)(iterator)
}
func ZendUserItGetNewIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *ZendObjectIterator {
	var iterator types.Zval
	var new_iterator *ZendObjectIterator
	var ce_it *types.ClassEntry
	ZendUserItNewIterator(ce, object, &iterator)
	if iterator.IsObject() {
		ce_it = types.Z_OBJCE(iterator)
	} else {
		ce_it = nil
	}
	if ce_it == nil || ce_it.GetGetIterator() == nil || ce_it.GetGetIterator() == ZendUserItGetNewIterator && iterator.Object() == object.Object() {
		if EG__().GetException() == nil {
			faults.ThrowExceptionEx(nil, 0, "Objects returned by %s::getIterator() must be traversable or implement interface Iterator", b.CondF(ce != nil, func() []byte { return ce.Name() }, func() []byte { return types.Z_OBJCE_P(object).Name() }))
		}
		// ZvalPtrDtor(&iterator)
		return nil
	}
	new_iterator = ce_it.GetGetIterator()(ce_it, &iterator, by_ref)
	// ZvalPtrDtor(&iterator)
	return new_iterator
}
func ZendImplementTraversable(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	/* check that class_type is traversable at c-level or implements at least one of 'aggregate' and 'Iterator' */

	var i uint32
	if class_type.GetGetIterator() != nil || class_type.GetParent() != nil && class_type.GetParent().GetGetIterator() != nil {
		return types.SUCCESS
	}
	if class_type.GetNumInterfaces() != 0 {
		b.Assert(class_type.IsResolvedInterfaces())
		for i = 0; i < class_type.GetNumInterfaces(); i++ {
			if class_type.GetInterfaces()[i] == ZendCeAggregate || class_type.GetInterfaces()[i] == ZendCeIterator {
				return types.SUCCESS
			}
		}
	}
	faults.ErrorNoreturn(faults.E_CORE_ERROR, "Class %s must implement interface %s as part of either %s or %s", class_type.Name(), ZendCeTraversable.Name(), ZendCeIterator.Name(), ZendCeAggregate.Name())
	return types.FAILURE
}
func ZendImplementAggregate(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	var i uint32
	var t int = -1
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil {
		if class_type.IsInternalClass() {

			/* inheritance ensures the class has necessary userland methods */

			return types.SUCCESS

			/* inheritance ensures the class has necessary userland methods */

		} else if class_type.GetGetIterator() != ZendUserItGetNewIterator {

			/* c-level get_iterator cannot be changed (exception being only Traversable is implemented) */

			if class_type.GetNumInterfaces() != 0 {
				b.Assert(class_type.IsResolvedInterfaces())
				for i = 0; i < class_type.GetNumInterfaces(); i++ {
					if class_type.GetInterfaces()[i] == ZendCeIterator {
						faults.ErrorNoreturn(faults.E_ERROR, "Class %s cannot implement both %s and %s at the same time", class_type.Name(), interface_.Name(), ZendCeIterator.Name())
						return types.FAILURE
					}
					if class_type.GetInterfaces()[i] == ZendCeTraversable {
						t = i
					}
				}
			}
			if t == -1 {
				return types.FAILURE
			}
		}
	}
	if class_type.GetParent() && (class_type.GetParent().ce_flags&types.AccReuseGetIterator) != 0 {
		class_type.SetGetIterator(class_type.GetParent().get_iterator)
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetNewIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.IsInternalClass() {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		}
		funcs_ptr.SetZfNewIterator(class_type.FunctionTable().Get("getiterator"))
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
			memset(funcs_ptr, 0, b.SizeOf("zend_class_iterator_funcs"))
		} else {
			funcs_ptr.SetZfNewIterator(nil)
		}
	}
	return types.SUCCESS
}
func ZendImplementIterator(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil && class_type.GetGetIterator() != ZendUserItGetIterator {
		if class_type.IsInternalClass() {

			/* inheritance ensures the class has the necessary userland methods */

			return types.SUCCESS

			/* inheritance ensures the class has the necessary userland methods */

		} else {

			/* c-level get_iterator cannot be changed */

			if class_type.GetGetIterator() == ZendUserItGetNewIterator {
				faults.ErrorNoreturn(faults.E_ERROR, "Class %s cannot implement both %s and %s at the same time", class_type.Name(), interface_.GetName().GetVal(), ZendCeAggregate.GetName().GetVal())
			}
			return types.FAILURE
		}
	}
	if class_type.GetParent() && (class_type.GetParent().ce_flags&types.AccReuseGetIterator) != 0 {
		class_type.SetGetIterator(class_type.GetParent().get_iterator)
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.IsInternalClass() {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		} else {
			funcs_ptr.SetZfRewind(class_type.FunctionTable().Get("rewind"))
			funcs_ptr.SetZfValid(class_type.FunctionTable().Get("valid"))
			funcs_ptr.SetZfKey(class_type.FunctionTable().Get("key"))
			funcs_ptr.SetZfCurrent(class_type.FunctionTable().Get("current"))
			funcs_ptr.SetZfNext(class_type.FunctionTable().Get("next"))
		}
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(CG__().GetArena(), b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
			memset(funcs_ptr, 0, b.SizeOf("zend_class_iterator_funcs"))
		} else {
			funcs_ptr.SetZfValid(nil)
			funcs_ptr.SetZfCurrent(nil)
			funcs_ptr.SetZfKey(nil)
			funcs_ptr.SetZfNext(nil)
			funcs_ptr.SetZfRewind(nil)
		}
	}
	return types.SUCCESS
}
func ZendImplementArrayaccess(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	return types.SUCCESS
}
func ZendUserSerialize(object *types.Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	var retval types.Zval
	var result int
	ZendCallMethodWith0Params(object, ce, ce.GetSerializeFunc(), "serialize", &retval)
	if retval.IsUndef() || EG__().GetException() != nil {
		result = types.FAILURE
	} else {
		switch retval.GetType() {
		case types.IS_NULL:

			/* we could also make this '*buf_len = 0' but this allows to skip variables */

			// ZvalPtrDtor(&retval)
			return types.FAILURE
		case types.IS_STRING:
			*buffer = (*uint8)(Estrndup(retval.String().GetVal(), retval.String().GetLen()))
			*buf_len = retval.String().GetLen()
			result = types.SUCCESS
		default:
			result = types.FAILURE
		}
		// ZvalPtrDtor(&retval)
	}
	if result == types.FAILURE && EG__().GetException() == nil {
		faults.ThrowExceptionEx(nil, 0, "%s::serialize() must return a string or NULL", ce.Name())
	}
	return result
}
func ZendUserUnserialize(object *types.Zval, ce *types.ClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	var zdata types.Zval
	if ObjectInitEx(object, ce) != types.SUCCESS {
		return types.FAILURE
	}
	zdata.SetStringVal(b.CastStr((*byte)(buf), buf_len))
	ZendCallMethodWith1Params(object, ce, ce.GetUnserializeFunc(), "unserialize", nil, &zdata)
	// ZvalPtrDtor(&zdata)
	if EG__().GetException() != nil {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}
func ZendClassSerializeDeny(object *types.Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	faults.ThrowExceptionEx(nil, 0, "Serialization of '%s' is not allowed", ce.Name())
	return types.FAILURE
}
func ZendClassUnserializeDeny(object *types.Zval, ce *types.ClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	faults.ThrowExceptionEx(nil, 0, "Unserialization of '%s' is not allowed", ce.Name())
	return types.FAILURE
}
func ZendImplementSerializable(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	if class_type.GetParent() && (class_type.GetParent().serialize || class_type.GetParent().unserialize) && operators.InstanceofFunctionEx(class_type.GetParent(), ZendCeSerializable, 1) == 0 {
		return types.FAILURE
	}
	if class_type.GetSerialize() == nil {
		class_type.SetSerialize(ZendUserSerialize)
	}
	if class_type.GetUnserialize() == nil {
		class_type.SetUnserialize(ZendUserUnserialize)
	}
	return types.SUCCESS
}
func ZendImplementCountable(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	return types.SUCCESS
}
func ZendRegisterInterfaces() {
	ZendCeTraversable = RegisterInternalInterface("Traversable", nil)
	ZendCeTraversable.SetInterfaceGetsImplemented(ZendImplementTraversable)

	ZendCeAggregate = RegisterInternalInterface("IteratorAggregate", ZendFuncsAggregate)
	ZendCeAggregate.SetInterfaceGetsImplemented(ZendImplementAggregate)
	ZendClassImplements(ZendCeAggregate, 1, ZendCeTraversable)

	ZendCeIterator = RegisterInternalInterface("Iterator", ZendFuncsIterator)
	ZendCeIterator.SetInterfaceGetsImplemented(ZendImplementIterator)
	ZendClassImplements(ZendCeIterator, 1, ZendCeTraversable)

	ZendCeArrayaccess = RegisterInternalInterface("ArrayAccess", ZendFuncsArrayaccess)
	ZendCeArrayaccess.SetInterfaceGetsImplemented(ZendImplementArrayaccess)

	ZendCeSerializable = RegisterInternalInterface("Serializable", ZendFuncsSerializable)
	ZendCeSerializable.SetInterfaceGetsImplemented(ZendImplementSerializable)

	ZendCeCountable = RegisterInternalInterface("Countable", ZendFuncsCountable)
	ZendCeCountable.SetInterfaceGetsImplemented(ZendImplementCountable)
}
