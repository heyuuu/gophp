package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/slicekit"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
)

func ZendCallMethodWith0Params(obj *types.Zval, objCe *types.ClassEntry, fnProxy *types.IFunction, functionName string, retval *types.Zval) *types.Zval {
	return ZendCallMethod(obj, objCe, fnProxy, functionName, retval)
}
func ZendCallMethodWith1Params(obj *types.Zval, objCe *types.ClassEntry, fnProxy *types.IFunction, functionName string, retval *types.Zval, arg1 *types.Zval) *types.Zval {
	return ZendCallMethod(obj, objCe, fnProxy, functionName, retval, arg1)
}
func ZendCallMethodWith2Params(obj *types.Zval, objCe *types.ClassEntry, fnProxy *types.IFunction, functionName string, retval *types.Zval, arg1 *types.Zval, arg2 *types.Zval) *types.Zval {
	return ZendCallMethod(obj, objCe, fnProxy, functionName, retval, arg1, arg2)
}

func ZendCallMethod(object *types.Zval, objCe *types.ClassEntry, fnProxy *types.IFunction, functionName string, retvalPtr *types.Zval, args ...*types.Zval) *types.Zval {
	var objPtr *types.Object = nil
	if object != nil {
		objPtr = object.Object()
	}

	var fci = types.InitFCallInfo(objPtr, retvalPtr, args...)

	var result int
	if fnProxy == nil && objCe == nil {
		/* no interest in caching and no information already present that is
		 * needed later inside zend_call_function. */
		fci.SetFunctionName(functionName)
		result = ZendCallFunction(fci, nil)
	} else {
		var fcic types.ZendFcallInfoCache
		fci.ClearFunctionName()
		if objCe == nil {
			if object != nil {
				objCe = types.Z_OBJCE_P(object)
			} else {
				objCe = nil
			}
		}
		if fnProxy == nil || (*fnProxy) == nil {
			if objCe != nil {
				fcic.SetFunctionHandler(objCe.FunctionTable().Get(functionName))
				if fcic.GetFunctionHandler() == nil {
					/* error at c-level */
					faults.ErrorNoreturn(faults.E_CORE_ERROR, fmt.Sprintf("Couldn't find implementation for method %s::%s", objCe.Name(), functionName))
				}
			} else {
				fcic.SetFunctionHandler(ZendFetchFunctionStr(functionName))
				if fcic.GetFunctionHandler() == nil {
					/* error at c-level */
					faults.ErrorNoreturn(faults.E_CORE_ERROR, fmt.Sprintf("Couldn't find implementation for function %s", functionName))
				}
			}
			if fnProxy != nil {
				*fnProxy = fcic.GetFunctionHandler()
			}
		} else {
			fcic.SetFunctionHandler(*fnProxy)
		}
		if object != nil {
			fcic.SetCalledScope(types.Z_OBJCE_P(object))
		} else {
			var calledScope *types.ClassEntry = ZendGetCalledScope(CurrEX())
			if objCe != nil && (calledScope == nil || operators.InstanceofFunction(calledScope, objCe) == 0) {
				fcic.SetCalledScope(objCe)
			} else {
				fcic.SetCalledScope(calledScope)
			}
		}
		if object != nil {
			fcic.SetObject(object.Object())
		} else {
			fcic.SetObject(nil)
		}
		result = ZendCallFunction(fci, &fcic)
	}
	if result == types.FAILURE {

		/* error at c-level */

		if objCe == nil {
			if object != nil {
				objCe = types.Z_OBJCE_P(object)
			} else {
				objCe = nil
			}
		}
		if EG__().NoException() {
			faults.ErrorNoreturn(faults.E_CORE_ERROR, fmt.Sprintf("Couldn't execute method %s%s%s", lang.CondF1(objCe != nil, func() []byte { return objCe.Name() }, ""), lang.Cond(objCe != nil, "::", ""), functionName))
		}
	}
	return retvalPtr
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
		if EG__().NoException() {
			faults.Error(faults.E_WARNING, fmt.Sprintf("Nothing returned from %s::key()", iter.GetCe().Name()))
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
		if EG__().NoException() {
			faults.ThrowException(nil, fmt.Sprintf("Objects returned by %s::getIterator() must be traversable or implement interface Iterator", lang.CondF(ce != nil, func() []byte { return ce.Name() }, func() []byte { return types.Z_OBJCE_P(object).Name() })), 0)
		}
		// ZvalPtrDtor(&iterator)
		return nil
	}
	new_iterator = ce_it.GetGetIterator()(ce_it, &iterator, by_ref)
	// ZvalPtrDtor(&iterator)
	return new_iterator
}
func isTraversableRootInterface(iface *types.ClassEntry) bool {
	return iface == ZendCeAggregate || iface == ZendCeIterator
}

func ZendImplementTraversable(interface_ *types.ClassEntry, class_type *types.ClassEntry) int {
	/* check that class_type is traversable at c-level or implements at least one of 'aggregate' and 'Iterator' */
	if class_type.GetGetIterator() != nil || class_type.GetParent() != nil && class_type.GetParent().GetGetIterator() != nil {
		return types.SUCCESS
	}
	if class_type.HasInterfaces() {
		b.Assert(class_type.IsResolvedInterfaces())
		if slicekit.Any(class_type.GetInterfaces(), isTraversableRootInterface) {
			return types.SUCCESS
		}
	}
	faults.ErrorNoreturn(faults.E_CORE_ERROR, fmt.Sprintf("Class %s must implement interface %s as part of either %s or %s", class_type.Name(), ZendCeTraversable.Name(), ZendCeIterator.Name(), ZendCeAggregate.Name()))
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

			if class_type.HasInterfaces() {
				b.Assert(class_type.IsResolvedInterfaces())
				for i = 0; i < class_type.GetNumInterfaces(); i++ {
					if class_type.GetInterfaces()[i] == ZendCeIterator {
						faults.ErrorNoreturn(faults.E_ERROR, fmt.Sprintf("Class %s cannot implement both %s and %s at the same time", class_type.Name(), interface_.Name(), ZendCeIterator.Name()))
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
	if class_type.GetParent() != nil && (class_type.GetParent().GetCeFlags()&types.AccReuseGetIterator) != 0 {
		class_type.SetGetIterator(class_type.GetParent().GetGetIterator())
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetNewIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.IsInternalClass() {
		if funcs_ptr == nil {
			funcs_ptr = NewClassIteratorFuncs()
			class_type.SetIteratorFuncsPtr(funcs_ptr)
		}
		funcs_ptr.SetZfNewIterator(class_type.FunctionTable().Get("getiterator"))
	} else {
		if funcs_ptr == nil {
			funcs_ptr = NewClassIteratorFuncs()
			class_type.SetIteratorFuncsPtr(funcs_ptr)
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
				faults.ErrorNoreturn(faults.E_ERROR, fmt.Sprintf("Class %s cannot implement both %s and %s at the same time", class_type.Name(), interface_.Name(), ZendCeAggregate.Name()))
			}
			return types.FAILURE
		}
	}
	if class_type.GetParent() && (class_type.GetParent().GetCeFlags()&types.AccReuseGetIterator) != 0 {
		class_type.SetGetIterator(class_type.GetParent().get_iterator)
		class_type.SetIsReuseGetIterator(true)
	} else {
		class_type.SetGetIterator(ZendUserItGetIterator)
	}
	funcs_ptr = class_type.GetIteratorFuncsPtr()
	if class_type.IsInternalClass() {
		if funcs_ptr == nil {
			class_type.SetIteratorFuncsPtr(NewClassIteratorFuncs())
		} else {
			funcs_ptr.SetZfRewind(class_type.FunctionTable().Get("rewind"))
			funcs_ptr.SetZfValid(class_type.FunctionTable().Get("valid"))
			funcs_ptr.SetZfKey(class_type.FunctionTable().Get("key"))
			funcs_ptr.SetZfCurrent(class_type.FunctionTable().Get("current"))
			funcs_ptr.SetZfNext(class_type.FunctionTable().Get("next"))
		}
	} else {
		if funcs_ptr == nil {
			class_type.SetIteratorFuncsPtr(NewClassIteratorFuncs())
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
	var ce = object.Object().GetCe()
	var retval types.Zval
	var result int
	ZendCallMethodWith0Params(object, ce, ce.GetSerializeFunc(), "serialize", &retval)
	if retval.IsUndef() || EG__().HasException() {
		result = types.FAILURE
	} else {
		switch retval.Type() {
		case types.IsNull:

			/* we could also make this '*buf_len = 0' but this allows to skip variables */

			// ZvalPtrDtor(&retval)
			return types.FAILURE
		case types.IsString:
			*buffer = (*uint8)(Estrndup(retval.StringEx().GetVal(), retval.StringEx().GetLen()))
			*buf_len = retval.StringEx().GetLen()
			result = types.SUCCESS
		default:
			result = types.FAILURE
		}
		// ZvalPtrDtor(&retval)
	}
	if result == types.FAILURE && EG__().NoException() {
		faults.ThrowException(nil, fmt.Sprintf("%s::serialize() must return a string or NULL", ce.Name()), 0)
	}
	return result
}
func ZendUserUnserialize(object *types.Zval, ce *types.ClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	var zdata types.Zval
	if ObjectInitEx(object, ce) != types.SUCCESS {
		return types.FAILURE
	}
	zdata.SetString(b.CastStr((*byte)(buf), buf_len))
	ZendCallMethodWith1Params(object, ce, ce.GetUnserializeFunc(), "unserialize", nil, &zdata)
	// ZvalPtrDtor(&zdata)
	if EG__().HasException() {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}
func ZendClassSerializeDeny(object *types.Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(object)
	faults.ThrowException(nil, fmt.Sprintf("Serialization of '%s' is not allowed", ce.Name()), 0)
	return types.FAILURE
}
func ZendClassUnserializeDeny(object *types.Zval, ce *types.ClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	faults.ThrowException(nil, fmt.Sprintf("Unserialization of '%s' is not allowed", ce.Name()), 0)
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
	ZendCeTraversable = RegisterInterface(&types.InternalClassDecl{
		Name:                     "Traversable",
		InterfaceGetsImplemented: ZendImplementTraversable,
	})

	ZendCeAggregate = RegisterInterface(&types.InternalClassDecl{
		Name:                     "IteratorAggregate",
		Functions:                ZendFuncsAggregate,
		Interfaces:               []*types.ClassEntry{ZendCeTraversable},
		InterfaceGetsImplemented: ZendImplementAggregate,
	})

	ZendCeIterator = RegisterInterface(&types.InternalClassDecl{
		Name:                     "Iterator",
		Functions:                ZendFuncsIterator,
		Interfaces:               []*types.ClassEntry{ZendCeTraversable},
		InterfaceGetsImplemented: ZendImplementIterator,
	})

	ZendCeArrayaccess = RegisterInterface(&types.InternalClassDecl{
		Name:                     "ArrayAccess",
		Functions:                ZendFuncsArrayaccess,
		InterfaceGetsImplemented: ZendImplementArrayaccess,
	})

	ZendCeSerializable = RegisterInterface(&types.InternalClassDecl{
		Name:                     "Serializable",
		Functions:                ZendFuncsSerializable,
		InterfaceGetsImplemented: ZendImplementSerializable,
	})

	ZendCeCountable = RegisterInterface(&types.InternalClassDecl{
		Name:                     "Countable",
		Functions:                ZendFuncsCountable,
		InterfaceGetsImplemented: ZendImplementCountable,
	})
}
