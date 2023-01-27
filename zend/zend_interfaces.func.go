// <<generate>>

package zend

import (
	b "sik/builtin"
)

func ZendCallMethodWith0Params(obj *Zval, obj_ce *ZendClassEntry, fn_proxy **ZendFunction, function_name string, retval *Zval) *Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, b.SizeOf("function_name")-1, retval, 0, nil, nil)
}
func ZendCallMethodWith1Params(obj *Zval, obj_ce *ZendClassEntry, fn_proxy **ZendFunction, function_name string, retval *Zval, arg1 *Zval) *Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, b.SizeOf("function_name")-1, retval, 1, arg1, nil)
}
func ZendCallMethodWith2Params(obj *Zval, obj_ce *ZendClassEntry, fn_proxy **ZendFunction, function_name string, retval *Zval, arg1 *Zval, arg2 *Zval) *Zval {
	return ZendCallMethod(obj, obj_ce, fn_proxy, function_name, b.SizeOf("function_name")-1, retval, 2, arg1, arg2)
}
func ZendCallMethod(object *Zval, obj_ce *ZendClassEntry, fn_proxy **ZendFunction, function_name string, function_name_len int, retval_ptr *Zval, param_count int, arg1 *Zval, arg2 *Zval) *Zval {
	var result int
	var fci ZendFcallInfo
	var retval Zval
	var params []Zval
	if param_count > 0 {
		ZVAL_COPY_VALUE(&params[0], arg1)
	}
	if param_count > 1 {
		ZVAL_COPY_VALUE(&params[1], arg2)
	}
	fci.SetSize(b.SizeOf("fci"))
	if object != nil {
		fci.SetObject(object.GetObj())
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

		ZVAL_STRINGL(fci.GetFunctionName(), function_name, function_name_len)
		result = ZendCallFunction(&fci, nil)
		ZvalPtrDtor(fci.GetFunctionName())
	} else {
		var fcic ZendFcallInfoCache
		ZVAL_UNDEF(fci.GetFunctionName())
		if obj_ce == nil {
			if object != nil {
				obj_ce = Z_OBJCE_P(object)
			} else {
				obj_ce = nil
			}
		}
		if fn_proxy == nil || (*fn_proxy) == nil {
			if obj_ce != nil {
				fcic.SetFunctionHandler(obj_ce.GetFunctionTable().StrFindPtr(function_name, function_name_len))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					ZendErrorNoreturn(E_CORE_ERROR, "Couldn't find implementation for method %s::%s", obj_ce.GetName().GetVal(), function_name)

					/* error at c-level */

				}
			} else {
				fcic.SetFunctionHandler(ZendFetchFunctionStr(function_name, function_name_len))
				if fcic.GetFunctionHandler() == nil {

					/* error at c-level */

					ZendErrorNoreturn(E_CORE_ERROR, "Couldn't find implementation for function %s", function_name)

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
			fcic.SetCalledScope(Z_OBJCE_P(object))
		} else {
			var called_scope *ZendClassEntry = ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData())
			if obj_ce != nil && (called_scope == nil || InstanceofFunction(called_scope, obj_ce) == 0) {
				fcic.SetCalledScope(obj_ce)
			} else {
				fcic.SetCalledScope(called_scope)
			}
		}
		if object != nil {
			fcic.SetObject(object.GetObj())
		} else {
			fcic.SetObject(nil)
		}
		result = ZendCallFunction(&fci, &fcic)
	}
	if result == FAILURE {

		/* error at c-level */

		if obj_ce == nil {
			if object != nil {
				obj_ce = Z_OBJCE_P(object)
			} else {
				obj_ce = nil
			}
		}
		if ExecutorGlobals.GetException() == nil {
			ZendErrorNoreturn(E_CORE_ERROR, "Couldn't execute method %s%s%s", b.CondF1(obj_ce != nil, func() []byte { return obj_ce.GetName().GetVal() }, ""), b.Cond(obj_ce != nil, "::", ""), function_name)
		}
	}
	if retval_ptr == nil {
		ZvalPtrDtor(&retval)
		return nil
	}
	return retval_ptr
}
func ZendUserItNewIterator(ce *ZendClassEntry, object *Zval, retval *Zval) {
	ZendCallMethodWith0Params(object, ce, ce.GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", retval)
}
func ZendUserItInvalidateCurrent(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	if !(Z_ISUNDEF(iter.GetValue())) {
		ZvalPtrDtor(iter.GetValue())
		ZVAL_UNDEF(iter.GetValue())
	}
}
func ZendUserItDtor(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZvalPtrDtor(object)
}
func ZendUserItValid(_iter *ZendObjectIterator) int {
	if _iter != nil {
		var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
		var object *Zval = iter.GetIt().GetData()
		var more Zval
		var result int
		ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfValid(), "valid", &more)
		result = IZendIsTrue(&more)
		ZvalPtrDtor(&more)
		if result != 0 {
			return SUCCESS
		} else {
			return FAILURE
		}
	}
	return FAILURE
}
func ZendUserItGetCurrentData(_iter *ZendObjectIterator) *Zval {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = iter.GetIt().GetData()
	if Z_ISUNDEF(iter.GetValue()) {
		ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfCurrent(), "current", iter.GetValue())
	}
	return iter.GetValue()
}
func ZendUserItGetCurrentKey(_iter *ZendObjectIterator, key *Zval) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = iter.GetIt().GetData()
	var retval Zval
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfKey(), "key", &retval)
	if retval.GetType() != IS_UNDEF {
		ZVAL_ZVAL(key, &retval, 1, 1)
	} else {
		if ExecutorGlobals.GetException() == nil {
			ZendError(E_WARNING, "Nothing returned from %s::key()", iter.GetCe().GetName().GetVal())
		}
		ZVAL_LONG(key, 0)
	}
}
func ZendUserItMoveForward(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfNext(), "next", nil)
}
func ZendUserItRewind(_iter *ZendObjectIterator) {
	var iter *ZendUserIterator = (*ZendUserIterator)(_iter)
	var object *Zval = iter.GetIt().GetData()
	ZendUserItInvalidateCurrent(_iter)
	ZendCallMethodWith0Params(object, iter.GetCe(), iter.GetCe().GetIteratorFuncsPtr().GetZfRewind(), "rewind", nil)
}
func ZendUserItGetIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator *ZendUserIterator
	if by_ref != 0 {
		ZendThrowError(nil, "An iterator cannot be used with foreach by reference")
		return nil
	}
	iterator = Emalloc(b.SizeOf("zend_user_iterator"))
	ZendIteratorInit((*ZendObjectIterator)(iterator))
	Z_ADDREF_P(object)
	ZVAL_OBJ(iterator.GetIt().GetData(), object.GetObj())
	iterator.GetIt().SetFuncs(&ZendInterfaceIteratorFuncsIterator)
	iterator.SetCe(Z_OBJCE_P(object))
	ZVAL_UNDEF(iterator.GetValue())
	return (*ZendObjectIterator)(iterator)
}
func ZendUserItGetNewIterator(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	var iterator Zval
	var new_iterator *ZendObjectIterator
	var ce_it *ZendClassEntry
	ZendUserItNewIterator(ce, object, &iterator)
	if iterator.IsType(IS_OBJECT) {
		ce_it = Z_OBJCE(iterator)
	} else {
		ce_it = nil
	}
	if ce_it == nil || ce_it.GetGetIterator() == nil || ce_it.GetGetIterator() == ZendUserItGetNewIterator && iterator.GetObj() == object.GetObj() {
		if ExecutorGlobals.GetException() == nil {
			ZendThrowExceptionEx(nil, 0, "Objects returned by %s::getIterator() must be traversable or implement interface Iterator", b.CondF(ce != nil, func() []byte { return ce.GetName().GetVal() }, func() []byte { return Z_OBJCE_P(object).GetName().GetVal() }))
		}
		ZvalPtrDtor(&iterator)
		return nil
	}
	new_iterator = ce_it.GetGetIterator()(ce_it, &iterator, by_ref)
	ZvalPtrDtor(&iterator)
	return new_iterator
}
func ZendImplementTraversable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	/* check that class_type is traversable at c-level or implements at least one of 'aggregate' and 'Iterator' */

	var i uint32
	if class_type.GetGetIterator() != nil || class_type.parent && class_type.parent.get_iterator {
		return SUCCESS
	}
	if class_type.GetNumInterfaces() != 0 {
		ZEND_ASSERT(class_type.IsResolvedInterfaces())
		for i = 0; i < class_type.GetNumInterfaces(); i++ {
			if class_type.interfaces[i] == ZendCeAggregate || class_type.interfaces[i] == ZendCeIterator {
				return SUCCESS
			}
		}
	}
	ZendErrorNoreturn(E_CORE_ERROR, "Class %s must implement interface %s as part of either %s or %s", class_type.GetName().GetVal(), ZendCeTraversable.GetName().GetVal(), ZendCeIterator.GetName().GetVal(), ZendCeAggregate.GetName().GetVal())
	return FAILURE
}
func ZendImplementAggregate(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	var i uint32
	var t int = -1
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil {
		if class_type.GetType() == ZEND_INTERNAL_CLASS {

			/* inheritance ensures the class has necessary userland methods */

			return SUCCESS

			/* inheritance ensures the class has necessary userland methods */

		} else if class_type.GetGetIterator() != ZendUserItGetNewIterator {

			/* c-level get_iterator cannot be changed (exception being only Traversable is implemented) */

			if class_type.GetNumInterfaces() != 0 {
				ZEND_ASSERT(class_type.IsResolvedInterfaces())
				for i = 0; i < class_type.GetNumInterfaces(); i++ {
					if class_type.interfaces[i] == ZendCeIterator {
						ZendErrorNoreturn(E_ERROR, "Class %s cannot implement both %s and %s at the same time", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeIterator.GetName().GetVal())
						return FAILURE
					}
					if class_type.interfaces[i] == ZendCeTraversable {
						t = i
					}
				}
			}
			if t == -1 {
				return FAILURE
			}
		}
	}
	if class_type.parent && (class_type.parent.ce_flags&ZEND_ACC_REUSE_GET_ITERATOR) != 0 {
		class_type.SetGetIterator(class_type.parent.get_iterator)
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetNewIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.GetType() == ZEND_INTERNAL_CLASS {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		}
		funcs_ptr.SetZfNewIterator(class_type.GetFunctionTable().StrFindPtr("getiterator", b.SizeOf("\"getiterator\"")-1))
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
			memset(funcs_ptr, 0, b.SizeOf("zend_class_iterator_funcs"))
		} else {
			funcs_ptr.SetZfNewIterator(nil)
		}
	}
	return SUCCESS
}
func ZendImplementIterator(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	var funcs_ptr *ZendClassIteratorFuncs
	if class_type.GetGetIterator() != nil && class_type.GetGetIterator() != ZendUserItGetIterator {
		if class_type.GetType() == ZEND_INTERNAL_CLASS {

			/* inheritance ensures the class has the necessary userland methods */

			return SUCCESS

			/* inheritance ensures the class has the necessary userland methods */

		} else {

			/* c-level get_iterator cannot be changed */

			if class_type.GetGetIterator() == ZendUserItGetNewIterator {
				ZendErrorNoreturn(E_ERROR, "Class %s cannot implement both %s and %s at the same time", class_type.GetName().GetVal(), interface_.GetName().GetVal(), ZendCeAggregate.GetName().GetVal())
			}
			return FAILURE
		}
	}
	if class_type.parent && (class_type.parent.ce_flags&ZEND_ACC_REUSE_GET_ITERATOR) != 0 {
		class_type.SetGetIterator(class_type.parent.get_iterator)
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.GetType() == ZEND_INTERNAL_CLASS {
		if funcs_ptr == nil {
			funcs_ptr = calloc(1, b.SizeOf("zend_class_iterator_funcs"))
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		} else {
			funcs_ptr.SetZfRewind(class_type.GetFunctionTable().StrFindPtr("rewind", b.SizeOf("\"rewind\"")-1))
			funcs_ptr.SetZfValid(class_type.GetFunctionTable().StrFindPtr("valid", b.SizeOf("\"valid\"")-1))
			funcs_ptr.SetZfKey(class_type.GetFunctionTable().StrFindPtr("key", b.SizeOf("\"key\"")-1))
			funcs_ptr.SetZfCurrent(class_type.GetFunctionTable().StrFindPtr("current", b.SizeOf("\"current\"")-1))
			funcs_ptr.SetZfNext(class_type.GetFunctionTable().StrFindPtr("next", b.SizeOf("\"next\"")-1))
		}
	} else {
		if funcs_ptr == nil {
			funcs_ptr = ZendArenaAlloc(&(CompilerGlobals.GetArena()), b.SizeOf("zend_class_iterator_funcs"))
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
	return SUCCESS
}
func ZendImplementArrayaccess(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	return SUCCESS
}
func ZendUserSerialize(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	var retval Zval
	var result int
	ZendCallMethodWith0Params(object, ce, ce.GetSerializeFunc(), "serialize", &retval)
	if retval.IsType(IS_UNDEF) || ExecutorGlobals.GetException() != nil {
		result = FAILURE
	} else {
		switch retval.GetType() {
		case IS_NULL:

			/* we could also make this '*buf_len = 0' but this allows to skip variables */

			ZvalPtrDtor(&retval)
			return FAILURE
		case IS_STRING:
			*buffer = (*uint8)(Estrndup(Z_STRVAL(retval), Z_STRLEN(retval)))
			*buf_len = Z_STRLEN(retval)
			result = SUCCESS
			break
		default:
			result = FAILURE
			break
		}
		ZvalPtrDtor(&retval)
	}
	if result == FAILURE && ExecutorGlobals.GetException() == nil {
		ZendThrowExceptionEx(nil, 0, "%s::serialize() must return a string or NULL", ce.GetName().GetVal())
	}
	return result
}
func ZendUserUnserialize(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	var zdata Zval
	if ObjectInitEx(object, ce) != SUCCESS {
		return FAILURE
	}
	ZVAL_STRINGL(&zdata, (*byte)(buf), buf_len)
	ZendCallMethodWith1Params(object, ce, ce.GetUnserializeFunc(), "unserialize", nil, &zdata)
	ZvalPtrDtor(&zdata)
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	} else {
		return SUCCESS
	}
}
func ZendClassSerializeDeny(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *ZendClassEntry = Z_OBJCE_P(object)
	ZendThrowExceptionEx(nil, 0, "Serialization of '%s' is not allowed", ce.GetName().GetVal())
	return FAILURE
}
func ZendClassUnserializeDeny(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	ZendThrowExceptionEx(nil, 0, "Unserialization of '%s' is not allowed", ce.GetName().GetVal())
	return FAILURE
}
func ZendImplementSerializable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	if class_type.parent && (class_type.parent.serialize || class_type.parent.unserialize) && InstanceofFunctionEx(class_type.parent, ZendCeSerializable, 1) == 0 {
		return FAILURE
	}
	if class_type.GetSerialize() == nil {
		class_type.SetSerialize(ZendUserSerialize)
	}
	if class_type.GetUnserialize() == nil {
		class_type.SetUnserialize(ZendUserUnserialize)
	}
	return SUCCESS
}
func ZendImplementCountable(interface_ *ZendClassEntry, class_type *ZendClassEntry) int {
	return SUCCESS
}
func ZendRegisterInterfaces() {
	var ce zend_class_entry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Traversable", b.SizeOf("\"Traversable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsTraversable)
	ZendCeTraversable = ZendRegisterInternalInterface(&ce)
	ZendCeTraversable.interface_gets_implemented = ZendImplementTraversable
	var ce zend_class_entry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("IteratorAggregate", b.SizeOf("\"IteratorAggregate\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsAggregate)
	ZendCeAggregate = ZendRegisterInternalInterface(&ce)
	ZendCeAggregate.interface_gets_implemented = ZendImplementAggregate
	ZendClassImplements(ZendCeAggregate, 1, ZendCeTraversable)
	var ce zend_class_entry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Iterator", b.SizeOf("\"Iterator\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsIterator)
	ZendCeIterator = ZendRegisterInternalInterface(&ce)
	ZendCeIterator.interface_gets_implemented = ZendImplementIterator
	ZendClassImplements(ZendCeIterator, 1, ZendCeTraversable)
	var ce zend_class_entry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("ArrayAccess", b.SizeOf("\"ArrayAccess\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsArrayaccess)
	ZendCeArrayaccess = ZendRegisterInternalInterface(&ce)
	ZendCeArrayaccess.interface_gets_implemented = ZendImplementArrayaccess
	var ce zend_class_entry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Serializable", b.SizeOf("\"Serializable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsSerializable)
	ZendCeSerializable = ZendRegisterInternalInterface(&ce)
	ZendCeSerializable.interface_gets_implemented = ZendImplementSerializable
	var ce ZendClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(ZendStringInitInterned("Countable", b.SizeOf("\"Countable\"")-1, 1))
	ce.SetBuiltinFunctions(ZendFuncsCountable)
	ZendCeCountable = ZendRegisterInternalInterface(&ce)
	ZendCeCountable.interface_gets_implemented = ZendImplementCountable
}
