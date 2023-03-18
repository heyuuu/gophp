// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

func SplDualItFromObj(obj *types.ZendObject) *SplDualItObject {
	return (*SplDualItObject)((*byte)(obj - zend_long((*byte)(&((*SplDualItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLDUAL_IT_P(zv *types.Zval) *SplDualItObject { return SplDualItFromObj(zv.GetObj()) }
func SplRecursiveItFromObj(obj *types.ZendObject) *SplRecursiveItObject {
	return (*SplRecursiveItObject)((*byte)(obj - zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLRECURSIVE_IT_P(zv *types.Zval) *SplRecursiveItObject {
	return SplRecursiveItFromObj(zv.GetObj())
}
func SPL_FETCH_SUB_ITERATOR(var_ *zend.ZendObjectIterator, object *SplRecursiveItObject) {
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	var_ = object.GetIterators()[object.GetLevel()].GetIterator()
}
func SplRecursiveItDtor(_iter *zend.ZendObjectIterator) {
	var iter *SplRecursiveItIterator = (*SplRecursiveItIterator)(_iter)
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(iter.GetIntern().GetData())
	var sub_iter *zend.ZendObjectIterator
	for object.GetLevel() > 0 {
		if !(object.GetIterators()[object.GetLevel()].GetZobject().IsUndef()) {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(object.GetIterators()[object.GetLevel()].GetZobject())
		}
		object.GetLevel()--
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.SetLevel(0)
	zend.ZvalPtrDtor(iter.GetIntern().GetData())
}
func SplRecursiveItValidEx(object *SplRecursiveItObject, zthis *types.Zval) int {
	var sub_iter *zend.ZendObjectIterator
	var level int = object.GetLevel()
	if object.GetIterators() == nil {
		return types.FAILURE
	}
	for level >= 0 {
		sub_iter = object.GetIterators()[level].GetIterator()
		if sub_iter.GetFuncs().GetValid()(sub_iter) == types.SUCCESS {
			return types.SUCCESS
		}
		level--
	}
	if object.GetEndIteration() != nil && object.GetInIteration() != 0 {
		zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetEndIteration(), "endIteration", nil)
	}
	object.SetInIteration(0)
	return types.FAILURE
}
func SplRecursiveItValid(iter *zend.ZendObjectIterator) int {
	return SplRecursiveItValidEx(Z_SPLRECURSIVE_IT_P(iter.GetData()), iter.GetData())
}
func SplRecursiveItGetCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(iter.GetData())
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	return sub_iter.GetFuncs().GetGetCurrentData()(sub_iter)
}
func SplRecursiveItGetCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(iter.GetData())
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	if sub_iter.GetFuncs().GetGetCurrentKey() != nil {
		sub_iter.GetFuncs().GetGetCurrentKey()(sub_iter, key)
	} else {
		key.SetLong(iter.GetIndex())
	}
}
func SplRecursiveItMoveForwardEx(object *SplRecursiveItObject, zthis *types.Zval) {
	var iterator *zend.ZendObjectIterator
	var zobject *types.Zval
	var ce *zend.ZendClassEntry
	var retval types.Zval
	var child types.Zval
	var sub_iter *zend.ZendObjectIterator
	var has_children int
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	for zend.EG__().GetException() == nil {
	next_step:
		iterator = object.GetIterators()[object.GetLevel()].GetIterator()
		switch object.GetIterators()[object.GetLevel()].GetState() {
		case RS_NEXT:
			iterator.GetFuncs().GetMoveForward()(iterator)
			if zend.EG__().GetException() != nil {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.ZendClearException()
				}
			}
			fallthrough
		case RS_START:
			if iterator.GetFuncs().GetValid()(iterator) == types.FAILURE {
				break
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_TEST)
			fallthrough
		case RS_TEST:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = object.GetIterators()[object.GetLevel()].GetZobject()
			if object.GetCallHasChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetCallHasChildren(), "callHasChildren", &retval)
			} else {
				zend.ZendCallMethodWith0Params(zobject, ce, nil, "haschildren", &retval)
			}
			if zend.EG__().GetException() != nil {
				if !object.IsRitCatchGetChild() {
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					return
				} else {
					zend.ZendClearException()
				}
			}
			if retval.GetType() != types.IS_UNDEF {
				has_children = zend.ZendIsTrue(&retval)
				zend.ZvalPtrDtor(&retval)
				if has_children != 0 {
					if object.GetMaxDepth() == -1 || object.GetMaxDepth() > object.GetLevel() {
						switch object.GetMode() {
						case RIT_LEAVES_ONLY:
							fallthrough
						case RIT_CHILD_FIRST:
							object.GetIterators()[object.GetLevel()].SetState(RS_CHILD)
							goto next_step
						case RIT_SELF_FIRST:
							object.GetIterators()[object.GetLevel()].SetState(RS_SELF)
							goto next_step
						}
					} else {

						/* do not recurse into */

						if object.GetMode() == RIT_LEAVES_ONLY {

							/* this is not a leave, so skip it */

							object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
							goto next_step
						}

						/* do not recurse into */

					}
				}
			}
			if object.GetNextElement() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetNextElement(), "nextelement", nil)
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			if zend.EG__().GetException() != nil {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.ZendClearException()
				}
			}
			return
		case RS_SELF:
			if object.GetNextElement() != nil && (object.GetMode() == RIT_SELF_FIRST || object.GetMode() == RIT_CHILD_FIRST) {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetNextElement(), "nextelement", nil)
			}
			if object.GetMode() == RIT_SELF_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_CHILD)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			return
		case RS_CHILD:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = object.GetIterators()[object.GetLevel()].GetZobject()
			if object.GetCallGetChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetCallGetChildren(), "callGetChildren", &child)
			} else {
				zend.ZendCallMethodWith0Params(zobject, ce, nil, "getchildren", &child)
			}
			if zend.EG__().GetException() != nil {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.ZendClearException()
					zend.ZvalPtrDtor(&child)
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					goto next_step
				}
			}
			if child.IsType(types.IS_UNDEF) || child.GetType() != types.IS_OBJECT || !(b.Assign(&ce, types.Z_OBJCE(child)) && zend.InstanceofFunction(ce, spl_ce_RecursiveIterator) != 0) {
				zend.ZvalPtrDtor(&child)
				zend.ZendThrowException(spl_ce_UnexpectedValueException, "Objects returned by RecursiveIterator::getChildren() must implement RecursiveIterator", 0)
				return
			}
			if object.GetMode() == RIT_CHILD_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_SELF)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")*(b.PreInc(&(object.GetLevel()))+1)))
			sub_iter = ce.GetGetIterator()(ce, &child, 0)
			types.ZVAL_COPY_VALUE(object.GetIterators()[object.GetLevel()].GetZobject(), &child)
			object.GetIterators()[object.GetLevel()].SetIterator(sub_iter)
			object.GetIterators()[object.GetLevel()].SetCe(ce)
			object.GetIterators()[object.GetLevel()].SetState(RS_START)
			if sub_iter.GetFuncs().GetRewind() != nil {
				sub_iter.GetFuncs().GetRewind()(sub_iter)
			}
			if object.GetBeginChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetBeginChildren(), "beginchildren", nil)
				if zend.EG__().GetException() != nil {
					if !object.IsRitCatchGetChild() {
						return
					} else {
						zend.ZendClearException()
					}
				}
			}
			goto next_step
		}

		/* no more elements */

		if object.GetLevel() > 0 {
			if object.GetEndChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetEndChildren(), "endchildren", nil)
				if zend.EG__().GetException() != nil {
					if !object.IsRitCatchGetChild() {
						return
					} else {
						zend.ZendClearException()
					}
				}
			}
			if object.GetLevel() > 0 {
				var garbage types.Zval
				types.ZVAL_COPY_VALUE(&garbage, object.GetIterators()[object.GetLevel()].GetZobject())
				object.GetIterators()[object.GetLevel()].GetZobject().SetUndef()
				zend.ZvalPtrDtor(&garbage)
				zend.ZendIteratorDtor(iterator)
				object.GetLevel()--
			}
		} else {
			return
		}

		/* no more elements */

	}
}
func SplRecursiveItRewindEx(object *SplRecursiveItObject, zthis *types.Zval) {
	var sub_iter *zend.ZendObjectIterator
	SPL_FETCH_SUB_ITERATOR(sub_iter, object)
	for object.GetLevel() != 0 {
		sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
		zend.ZendIteratorDtor(sub_iter)
		zend.ZvalPtrDtor(object.GetIterators()[b.PostDec(&(object.GetLevel()))].GetZobject())
		if zend.EG__().GetException() == nil && (object.GetEndChildren() == nil || object.GetEndChildren().GetScope() != spl_ce_RecursiveIteratorIterator) {
			zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetEndChildren(), "endchildren", nil)
		}
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.GetIterators()[0].SetState(RS_START)
	sub_iter = object.GetIterators()[0].GetIterator()
	if sub_iter.GetFuncs().GetRewind() != nil {
		sub_iter.GetFuncs().GetRewind()(sub_iter)
	}
	if zend.EG__().GetException() == nil && object.GetBeginIteration() != nil && object.GetInIteration() == 0 {
		zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetBeginIteration(), "beginIteration", nil)
	}
	object.SetInIteration(1)
	SplRecursiveItMoveForwardEx(object, zthis)
}
func SplRecursiveItMoveForward(iter *zend.ZendObjectIterator) {
	SplRecursiveItMoveForwardEx(Z_SPLRECURSIVE_IT_P(iter.GetData()), iter.GetData())
}
func SplRecursiveItRewind(iter *zend.ZendObjectIterator) {
	SplRecursiveItRewindEx(Z_SPLRECURSIVE_IT_P(iter.GetData()), iter.GetData())
}
func SplRecursiveItGetIterator(ce *zend.ZendClassEntry, zobject *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplRecursiveItIterator
	var object *SplRecursiveItObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_recursive_it_iterator"))
	object = Z_SPLRECURSIVE_IT_P(zobject)
	if object.GetIterators() == nil {
		zend.ZendError(zend.E_ERROR, "The object to be iterated is in an invalid state: "+"the parent constructor has not been called")
	}
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	zobject.AddRefcount()
	iterator.GetIntern().GetData().SetObject(zobject.GetObj())
	iterator.GetIntern().SetFuncs(&SplRecursiveItIteratorFuncs)
	return (*zend.ZendObjectIterator)(iterator)
}
func SplRecursiveItItConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, rit_type RecursiveItItType) {
	var object *types.Zval = zend.ZEND_THIS(executeData)
	var intern *SplRecursiveItObject
	var iterator *types.Zval
	var ce_iterator *zend.ZendClassEntry
	var mode zend.ZendLong
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	var caching_it types.Zval
	var aggregate_retval types.Zval
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
	switch rit_type {
	case RIT_RecursiveTreeIterator:
		var caching_it_flags types.Zval
		var user_caching_it_flags *types.Zval = nil
		mode = RIT_SELF_FIRST
		flags = RTIT_BYPASS_KEY
		if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, executeData.NumArgs(), "o|lzl", &iterator, &flags, &user_caching_it_flags, &mode) == types.SUCCESS {
			if zend.InstanceofFunction(types.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, types.Z_OBJCE_P(iterator), types.Z_OBJCE_P(iterator).GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				iterator.AddRefcount()
			}
			if user_caching_it_flags != nil {
				types.ZVAL_COPY(&caching_it_flags, user_caching_it_flags)
			} else {
				caching_it_flags.SetLong(CIT_CATCH_GET_CHILD)
			}
			SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &caching_it, iterator, &caching_it_flags)
			zend.ZvalPtrDtor(&caching_it_flags)
			zend.ZvalPtrDtor(iterator)
			iterator = &caching_it
		} else {
			iterator = nil
		}
	case RIT_RecursiveIteratorIterator:
		fallthrough
	default:
		mode = RIT_LEAVES_ONLY
		flags = 0
		if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, executeData.NumArgs(), "o|ll", &iterator, &mode, &flags) == types.SUCCESS {
			if zend.InstanceofFunction(types.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, types.Z_OBJCE_P(iterator), types.Z_OBJCE_P(iterator).GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				iterator.AddRefcount()
			}
		} else {
			iterator = nil
		}
	}
	if iterator == nil || zend.InstanceofFunction(types.Z_OBJCE_P(iterator), spl_ce_RecursiveIterator) == 0 {
		if iterator != nil {
			zend.ZvalPtrDtor(iterator)
		}
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "An instance of RecursiveIterator or IteratorAggregate creating it is required", 0)
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	intern = Z_SPLRECURSIVE_IT_P(object)
	intern.SetIterators(zend.Emalloc(b.SizeOf("spl_sub_iterator")))
	intern.SetLevel(0)
	intern.SetMode(mode)
	intern.SetFlags(int(flags))
	intern.SetMaxDepth(-1)
	intern.SetInIteration(0)
	intern.SetCe(types.Z_OBJCE_P(object))
	intern.SetBeginIteration(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "beginiteration", b.SizeOf("\"beginiteration\"")-1))
	if intern.GetBeginIteration().GetScope() == ce_base {
		intern.SetBeginIteration(nil)
	}
	intern.SetEndIteration(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "enditeration", b.SizeOf("\"enditeration\"")-1))
	if intern.GetEndIteration().GetScope() == ce_base {
		intern.SetEndIteration(nil)
	}
	intern.SetCallHasChildren(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "callhaschildren", b.SizeOf("\"callHasChildren\"")-1))
	if intern.GetCallHasChildren().GetScope() == ce_base {
		intern.SetCallHasChildren(nil)
	}
	intern.SetCallGetChildren(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "callgetchildren", b.SizeOf("\"callGetChildren\"")-1))
	if intern.GetCallGetChildren().GetScope() == ce_base {
		intern.SetCallGetChildren(nil)
	}
	intern.SetBeginChildren(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "beginchildren", b.SizeOf("\"beginchildren\"")-1))
	if intern.GetBeginChildren().GetScope() == ce_base {
		intern.SetBeginChildren(nil)
	}
	intern.SetEndChildren(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "endchildren", b.SizeOf("\"endchildren\"")-1))
	if intern.GetEndChildren().GetScope() == ce_base {
		intern.SetEndChildren(nil)
	}
	intern.SetNextElement(zend.ZendHashStrFindPtr(intern.GetCe().GetFunctionTable(), "nextelement", b.SizeOf("\"nextElement\"")-1))
	if intern.GetNextElement().GetScope() == ce_base {
		intern.SetNextElement(nil)
	}
	ce_iterator = types.Z_OBJCE_P(iterator)
	intern.GetIterators()[0].SetIterator(ce_iterator.GetGetIterator()(ce_iterator, iterator, 0))
	intern.GetIterators()[0].GetZobject().SetObject(iterator.GetObj())
	intern.GetIterators()[0].SetCe(ce_iterator)
	intern.GetIterators()[0].SetState(RS_START)
	zend.ZendRestoreErrorHandling(&error_handling)
	if zend.EG__().GetException() != nil {
		var sub_iter *zend.ZendObjectIterator
		for intern.GetLevel() >= 0 {
			sub_iter = intern.GetIterators()[intern.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(intern.GetIterators()[b.PostDec(&(intern.GetLevel()))].GetZobject())
		}
		zend.Efree(intern.GetIterators())
		intern.SetIterators(nil)
	}
}
func zim_spl_RecursiveIteratorIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplRecursiveItItConstruct(executeData, return_value, spl_ce_RecursiveIteratorIterator, zend.ZendCeIterator, RIT_RecursiveIteratorIterator)
}
func zim_spl_RecursiveIteratorIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplRecursiveItRewindEx(object, zend.ZEND_THIS(executeData))
}
func zim_spl_RecursiveIteratorIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, SplRecursiveItValidEx(object, zend.ZEND_THIS(executeData)) == types.SUCCESS)
	return
}
func zim_spl_RecursiveIteratorIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var iterator *zend.ZendObjectIterator
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	if iterator.GetFuncs().GetGetCurrentKey() != nil {
		iterator.GetFuncs().GetGetCurrentKey()(iterator, return_value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_RecursiveIteratorIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var iterator *zend.ZendObjectIterator
	var data *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	data = iterator.GetFuncs().GetGetCurrentData()(iterator)
	if data != nil {
		types.ZVAL_COPY_DEREF(return_value, data)
	}
}
func zim_spl_RecursiveIteratorIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplRecursiveItMoveForwardEx(object, zend.ZEND_THIS(executeData))
}
func zim_spl_RecursiveIteratorIterator_getDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetLong(object.GetLevel())
	return
}
func zim_spl_RecursiveIteratorIterator_getSubIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var level zend.ZendLong = object.GetLevel()
	var value *types.Zval
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &level) == types.FAILURE {
		return
	}
	if level < 0 || level > object.GetLevel() {
		return_value.SetNull()
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	value = object.GetIterators()[level].GetZobject()
	types.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_RecursiveIteratorIterator_getInnerIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var zobject *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	types.ZVAL_COPY_DEREF(return_value, zobject)
}
func zim_spl_RecursiveIteratorIterator_beginIteration(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endIteration(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_callHasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var ce *zend.ZendClassEntry
	var zobject *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		return_value.SetNull()
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	if zobject.IsType(types.IS_UNDEF) {
		return_value.SetFalse()
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "haschildren", return_value)
		if return_value.IsType(types.IS_UNDEF) {
			return_value.SetFalse()
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_callGetChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var ce *zend.ZendClassEntry
	var zobject *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	if zobject.IsType(types.IS_UNDEF) {
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "getchildren", return_value)
		if return_value.IsType(types.IS_UNDEF) {
			return_value.SetNull()
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_beginChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_nextElement(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_setMaxDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var max_depth zend.ZendLong = -1
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &max_depth) == types.FAILURE {
		return
	}
	if max_depth < -1 {
		zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter max_depth must be >= -1", 0)
		return
	} else if max_depth > core.INT_MAX {
		max_depth = core.INT_MAX
	}
	object.SetMaxDepth(int(max_depth))
}
func zim_spl_RecursiveIteratorIterator_getMaxDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetMaxDepth() == -1 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(object.GetMaxDepth())
		return
	}
}
func SplRecursiveItGetMethod(zobject **types.ZendObject, method *types.ZendString, key *types.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var object *SplRecursiveItObject = SplRecursiveItFromObj(*zobject)
	var level zend.ZendLong = object.GetLevel()
	var zobj *types.Zval
	if object.GetIterators() == nil {
		core.PhpErrorDocref(nil, zend.E_ERROR, "The %s instance wasn't initialized properly", zobject.GetCe().GetName().GetVal())
	}
	zobj = object.GetIterators()[level].GetZobject()
	function_handler = zend.ZendStdGetMethod(zobject, method, key)
	if function_handler == nil {
		if b.Assign(&function_handler, zend.ZendHashFindPtr(types.Z_OBJCE_P(zobj).GetFunctionTable(), method)) == nil {
			*zobject = zobj.GetObj()
			function_handler = zobject.GetHandlers().GetGetMethod()(zobject, method, key)
		} else {
			*zobject = zobj.GetObj()
		}
	}
	return function_handler
}
func spl_RecursiveIteratorIterator_dtor(_object *types.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	var sub_iter *zend.ZendObjectIterator

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	if object.GetIterators() != nil {
		for object.GetLevel() >= 0 {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(object.GetIterators()[b.PostDec(&(object.GetLevel()))].GetZobject())
		}
		zend.Efree(object.GetIterators())
		object.SetIterators(nil)
	}
}
func spl_RecursiveIteratorIterator_free_storage(_object *types.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	if object.GetIterators() != nil {
		zend.Efree(object.GetIterators())
		object.SetIterators(nil)
		object.SetLevel(0)
	}
	zend.ZendObjectStdDtor(object.GetStd())
	object.GetPrefix()[0].Free()
	object.GetPrefix()[1].Free()
	object.GetPrefix()[2].Free()
	object.GetPrefix()[3].Free()
	object.GetPrefix()[4].Free()
	object.GetPrefix()[5].Free()
	object.GetPostfix()[0].Free()
}
func spl_RecursiveIteratorIterator_new_ex(class_type *zend.ZendClassEntry, init_prefix int) *types.ZendObject {
	var intern *SplRecursiveItObject
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_recursive_it_object"), class_type)
	if init_prefix != 0 {
		intern.GetPrefix()[0].AppendString("")
		intern.GetPrefix()[1].AppendString("| ")
		intern.GetPrefix()[2].AppendString("  ")
		intern.GetPrefix()[3].AppendString("|-")
		intern.GetPrefix()[4].AppendString("\\-")
		intern.GetPrefix()[5].AppendString("")
		intern.GetPostfix()[0].AppendString("")
	}
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	intern.GetStd().SetHandlers(&SplHandlersRecItIt)
	return intern.GetStd()
}
func spl_RecursiveIteratorIterator_new(class_type *zend.ZendClassEntry) *types.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 0)
}
func spl_RecursiveTreeIterator_new(class_type *zend.ZendClassEntry) *types.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 1)
}
func SplRecursiveTreeIteratorGetPrefix(object *SplRecursiveItObject, return_value *types.Zval) {
	var str zend.SmartStr = zend.MakeSmartStr(0)
	var has_next types.Zval
	var level int
	str.AppendString(object.GetPrefix()[0].GetS().GetStr())
	for level = 0; level < object.GetLevel(); level++ {
		zend.ZendCallMethodWith0Params(object.GetIterators()[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
		if has_next.GetType() != types.IS_UNDEF {
			if has_next.IsType(types.IS_TRUE) {
				str.AppendString(object.GetPrefix()[1].GetS().GetStr())
			} else {
				str.AppendString(object.GetPrefix()[2].GetS().GetStr())
			}
			zend.ZvalPtrDtor(&has_next)
		}
	}
	zend.ZendCallMethodWith0Params(object.GetIterators()[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
	if has_next.GetType() != types.IS_UNDEF {
		if has_next.IsType(types.IS_TRUE) {
			str.AppendString(object.GetPrefix()[3].GetS().GetStr())
		} else {
			str.AppendString(object.GetPrefix()[4].GetS().GetStr())
		}
		zend.ZvalPtrDtor(&has_next)
	}
	str.AppendString(object.GetPrefix()[5].GetS().GetStr())
	str.ZeroTail()
	return_value.SetString(str.GetS())
	return
}
func SplRecursiveTreeIteratorGetEntry(object *SplRecursiveItObject, return_value *types.Zval) {
	var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	var data *types.Zval
	data = iterator.GetFuncs().GetGetCurrentData()(iterator)
	if data != nil {
		data = types.ZVAL_DEREF(data)

		/* TODO: Remove this special case? */

		if data.IsType(types.IS_ARRAY) {
			return_value.SetInternedString(types.ZSTR_ARRAY_CAPITALIZED)
		} else {
			types.ZVAL_COPY(return_value, data)
			zend.ConvertToString(return_value)
		}

		/* TODO: Remove this special case? */

	}
}
func SplRecursiveTreeIteratorGetPostfix(object *SplRecursiveItObject, return_value *types.Zval) {
	return_value.SetString(object.GetPostfix()[0].GetS())
	return_value.AddRefcount()
}
func zim_spl_RecursiveTreeIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplRecursiveItItConstruct(executeData, return_value, spl_ce_RecursiveTreeIterator, zend.ZendCeIterator, RIT_RecursiveTreeIterator)
}
func zim_spl_RecursiveTreeIterator_setPrefixPart(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var part zend.ZendLong
	var prefix *byte
	var prefix_len int
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParameters(executeData.NumArgs(), "ls", &part, &prefix, &prefix_len) == types.FAILURE {
		return
	}
	if 0 > part || part > 5 {
		zend.ZendThrowExceptionEx(spl_ce_OutOfRangeException, 0, "Use RecursiveTreeIterator::PREFIX_* constant")
		return
	}
	object.GetPrefix()[part].Free()
	object.GetPrefix()[part].AppendString(b.CastStr(prefix, prefix_len))
}
func zim_spl_RecursiveTreeIterator_getPrefix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPrefix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_setPostfix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var postfix *byte
	var postfix_len int
	if zend.ZendParseParameters(executeData.NumArgs(), "s", &postfix, &postfix_len) == types.FAILURE {
		return
	}
	object.GetPostfix()[0].Free()
	object.GetPostfix()[0].AppendString(b.CastStr(postfix, postfix_len))
}
func zim_spl_RecursiveTreeIterator_getEntry(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetEntry(object, return_value)
}
func zim_spl_RecursiveTreeIterator_getPostfix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var prefix types.Zval
	var entry types.Zval
	var postfix types.Zval
	var ptr *byte
	var str *types.ZendString
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	if object.IsRtitBypassCurrent() {
		var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
		var data *types.Zval
		SPL_FETCH_SUB_ITERATOR(iterator, object)
		data = iterator.GetFuncs().GetGetCurrentData()(iterator)
		if data != nil {
			types.ZVAL_COPY_DEREF(return_value, data)
			return
		} else {
			return_value.SetNull()
			return
		}
	}
	prefix.SetNull()
	entry.SetNull()
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetEntry(object, &entry)
	if entry.GetType() != types.IS_STRING {
		zend.ZvalPtrDtor(&prefix)
		zend.ZvalPtrDtor(&entry)
		return_value.SetNull()
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = types.ZendStringAlloc(prefix.GetStr().GetLen()+entry.GetStr().GetLen()+postfix.GetStr().GetLen(), 0)
	ptr = str.GetVal()
	memcpy(ptr, prefix.GetStr().GetVal(), prefix.GetStr().GetLen())
	ptr += prefix.GetStr().GetLen()
	memcpy(ptr, entry.GetStr().GetVal(), entry.GetStr().GetLen())
	ptr += entry.GetStr().GetLen()
	memcpy(ptr, postfix.GetStr().GetVal(), postfix.GetStr().GetLen())
	ptr += postfix.GetStr().GetLen()
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&entry)
	zend.ZvalPtrDtor(&postfix)
	return_value.SetString(str)
	return
}
func zim_spl_RecursiveTreeIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS(executeData))
	var iterator *zend.ZendObjectIterator
	var prefix types.Zval
	var key types.Zval
	var postfix types.Zval
	var key_copy types.Zval
	var ptr *byte
	var str *types.ZendString
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	if iterator.GetFuncs().GetGetCurrentKey() != nil {
		iterator.GetFuncs().GetGetCurrentKey()(iterator, &key)
	} else {
		key.SetNull()
	}
	if object.IsRtitBypassKey() {
		zend.ZVAL_ZVAL(return_value, &key, 1, 1)
		return
	}
	if key.GetType() != types.IS_STRING {
		if zend.ZendMakePrintableZval(&key, &key_copy) != 0 {
			key = key_copy
		}
	}
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = types.ZendStringAlloc(prefix.GetStr().GetLen()+key.GetStr().GetLen()+postfix.GetStr().GetLen(), 0)
	ptr = str.GetVal()
	memcpy(ptr, prefix.GetStr().GetVal(), prefix.GetStr().GetLen())
	ptr += prefix.GetStr().GetLen()
	memcpy(ptr, key.GetStr().GetVal(), key.GetStr().GetLen())
	ptr += key.GetStr().GetLen()
	memcpy(ptr, postfix.GetStr().GetVal(), postfix.GetStr().GetLen())
	ptr += postfix.GetStr().GetLen()
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&key)
	zend.ZvalPtrDtor(&postfix)
	return_value.SetString(str)
	return
}
func SplDualItGetMethod(object **types.ZendObject, method *types.ZendString, key *types.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var intern *SplDualItObject
	intern = SplDualItFromObj(*object)
	function_handler = zend.ZendStdGetMethod(object, method, key)
	if function_handler == nil && intern.GetCe() != nil {
		if b.Assign(&function_handler, zend.ZendHashFindPtr(intern.GetCe().GetFunctionTable(), method)) == nil {
			if types.Z_OBJ_HT(intern.GetZobject()).GetGetMethod() != nil {
				*object = intern.GetZobject().GetObj()
				function_handler = object.GetHandlers().GetGetMethod()(object, method, key)
			}
		} else {
			*object = intern.GetZobject().GetObj()
		}
	}
	return function_handler
}
func APPENDIT_CHECK_CTOR(intern *SplDualItObject) {
	if intern.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "Classes derived from %s must call %s::__construct()", spl_ce_AppendIterator.GetName().GetVal(), spl_ce_AppendIterator.GetName().GetVal())
		return
	}
}
func SplCitCheckFlags(flags zend.ZendLong) int {
	var cnt zend.ZendLong = 0
	if (flags & CIT_CALL_TOSTRING) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_KEY) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_CURRENT) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if (flags & CIT_TOSTRING_USE_INNER) != 0 {
		cnt += 1
	} else {
		cnt += 0
	}
	if cnt <= 1 {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SplDualItConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, dit_type DualItType) *SplDualItObject {
	var zobject *types.Zval
	var retval types.Zval
	var intern *SplDualItObject
	var ce *zend.ZendClassEntry = nil
	var inc_refcount int = 1
	var error_handling zend.ZendErrorHandling
	intern = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if intern.GetDitType() != DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s::getIterator() must be called exactly once per instance", ce_base.GetName().GetVal())
		return nil
	}
	intern.SetDitType(dit_type)
	switch dit_type {
	case DIT_LimitIterator:
		intern.SetOffset(0)
		intern.SetCount(-1)
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O|ll", &zobject, ce_inner, intern.GetOffset(), intern.GetCount()) == types.FAILURE {
			return nil
		}
		if intern.GetOffset() < 0 {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter offset must be >= 0", 0)
			return nil
		}
		if intern.GetCount() < 0 && intern.GetCount() != -1 {
			zend.ZendThrowException(spl_ce_OutOfRangeException, "Parameter count must either be -1 or a value greater than or equal 0", 0)
			return nil
		}
	case DIT_CachingIterator:
		fallthrough
	case DIT_RecursiveCachingIterator:
		var flags zend.ZendLong = CIT_CALL_TOSTRING
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O|l", &zobject, ce_inner, &flags) == types.FAILURE {
			return nil
		}
		if SplCitCheckFlags(flags) != types.SUCCESS {
			zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
			return nil
		}
		intern.AddUCachingFlags(flags & CIT_PUBLIC)
		zend.ArrayInit(intern.GetZcache())
	case DIT_IteratorIterator:
		var ce_cast *zend.ZendClassEntry
		var class_name *types.ZendString
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O|S", &zobject, ce_inner, &class_name) == types.FAILURE {
			return nil
		}
		ce = types.Z_OBJCE_P(zobject)
		if zend.InstanceofFunction(ce, zend.ZendCeIterator) == 0 {
			if executeData.NumArgs() > 1 {
				if !(b.Assign(&ce_cast, zend.ZendLookupClass(class_name))) || zend.InstanceofFunction(ce, ce_cast) == 0 || ce_cast.GetGetIterator() == nil {
					zend.ZendThrowException(spl_ce_LogicException, "Class to downcast to not found or not base class or does not implement Traversable", 0)
					return nil
				}
				ce = ce_cast
			}
			if zend.InstanceofFunction(ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(zobject, ce, ce.GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &retval)
				if zend.EG__().GetException() != nil {
					zend.ZvalPtrDtor(&retval)
					return nil
				}
				if retval.GetType() != types.IS_OBJECT || zend.InstanceofFunction(types.Z_OBJCE(retval), zend.ZendCeTraversable) == 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "%s::getIterator() must return an object that implements Traversable", ce.GetName().GetVal())
					return nil
				}
				zobject = &retval
				ce = types.Z_OBJCE_P(zobject)
				inc_refcount = 0
			}
		}
	case DIT_AppendIterator:
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		SplInstantiate(spl_ce_ArrayIterator, intern.GetZarrayit())
		zend.ZendCallMethodWith0Params(intern.GetZarrayit(), spl_ce_ArrayIterator, spl_ce_ArrayIterator.GetConstructor(), "__construct", nil)
		intern.SetUAppendIterator(spl_ce_ArrayIterator.GetGetIterator()(spl_ce_ArrayIterator, intern.GetZarrayit(), 0))
		zend.ZendRestoreErrorHandling(&error_handling)
		return intern
	case DIT_RegexIterator:
		fallthrough
	case DIT_RecursiveRegexIterator:
		var regex *types.ZendString
		var mode zend.ZendLong = REGIT_MODE_MATCH
		intern.SetUseFlags(executeData.NumArgs() >= 5)
		intern.SetURegexFlags(0)
		intern.SetPregFlags(0)
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "OS|lll", &zobject, ce_inner, &regex, &mode, intern.GetURegexFlags(), intern.GetPregFlags()) == types.FAILURE {
			return nil
		}
		if mode < 0 || mode >= REGIT_MODE_MAX {
			zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+zend.ZEND_LONG_FMT, mode)
			return nil
		}
		intern.SetMode(mode)
		intern.SetURegexRegex(regex.Copy())
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		intern.SetPce(pcre_get_compiled_regex_cache(regex))
		zend.ZendRestoreErrorHandling(&error_handling)
		if intern.GetPce() == nil {

			/* pcre_get_compiled_regex_cache has already sent error */

			return nil

			/* pcre_get_compiled_regex_cache has already sent error */

		}
		php_pcre_pce_incref(intern.GetPce())
	case DIT_CallbackFilterIterator:
		fallthrough
	case DIT_RecursiveCallbackFilterIterator:
		var cfi *_spl_cbfilter_it_intern = zend.Emalloc(b.SizeOf("* cfi"))
		cfi.GetFci().SetObject(nil)
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "Of", &zobject, ce_inner, cfi.GetFci(), cfi.GetFcc()) == types.FAILURE {
			zend.Efree(cfi)
			return nil
		}
		cfi.GetFci().GetFunctionName().TryAddRefcount()
		cfi.SetObject(cfi.GetFcc().GetObject())
		if cfi.GetObject() != nil {
			cfi.GetObject().AddRefcount()
		}
		intern.SetCbfilter(cfi)
	default:
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O", &zobject, ce_inner) == types.FAILURE {
			return nil
		}
	}
	if inc_refcount != 0 {
		zobject.AddRefcount()
	}
	intern.GetZobject().SetObject(zobject.GetObj())
	if dit_type == DIT_IteratorIterator {
		intern.SetCe(ce)
	} else {
		intern.SetCe(types.Z_OBJCE_P(zobject))
	}
	intern.SetObject(zobject.GetObj())
	intern.SetInnerIterator(intern.GetCe().GetGetIterator()(intern.GetCe(), zobject, 0))
	return intern
}
func zim_spl_FilterIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_FilterIterator, zend.ZendCeIterator, DIT_FilterIterator)
}
func zim_spl_CallbackFilterIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_CallbackFilterIterator, zend.ZendCeIterator, DIT_CallbackFilterIterator)
}
func zim_spl_dual_it_getInnerIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !(intern.GetZobject().IsUndef()) {
		var value *types.Zval = intern.GetZobject()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func SplDualItFree(intern *SplDualItObject) {
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().GetFuncs().GetInvalidateCurrent() != nil {
		intern.GetInnerIterator().GetFuncs().GetInvalidateCurrent()(intern.GetInnerIterator())
	}
	if intern.GetData().GetType() != types.IS_UNDEF {
		zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
	}
	if intern.GetKey().GetType() != types.IS_UNDEF {
		zend.ZvalPtrDtor(intern.GetKey())
		intern.GetKey().SetUndef()
	}
	if intern.GetDitType() == DIT_CachingIterator || intern.GetDitType() == DIT_RecursiveCachingIterator {
		if intern.GetZstr().GetType() != types.IS_UNDEF {
			zend.ZvalPtrDtor(intern.GetZstr())
			intern.GetZstr().SetUndef()
		}
		if intern.GetZchildren().GetType() != types.IS_UNDEF {
			zend.ZvalPtrDtor(intern.GetZchildren())
			intern.GetZchildren().SetUndef()
		}
	}
}
func SplDualItRewind(intern *SplDualItObject) {
	SplDualItFree(intern)
	intern.SetPos(0)
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().GetFuncs().GetRewind() != nil {
		intern.GetInnerIterator().GetFuncs().GetRewind()(intern.GetInnerIterator())
	}
}
func SplDualItValid(intern *SplDualItObject) int {
	if intern.GetInnerIterator() == nil {
		return types.FAILURE
	}

	/* FAILURE / SUCCESS */

	return intern.GetInnerIterator().GetFuncs().GetValid()(intern.GetInnerIterator())

	/* FAILURE / SUCCESS */
}
func SplDualItFetch(intern *SplDualItObject, check_more int) int {
	var data *types.Zval
	SplDualItFree(intern)
	if check_more == 0 || SplDualItValid(intern) == types.SUCCESS {
		data = intern.GetInnerIterator().GetFuncs().GetGetCurrentData()(intern.GetInnerIterator())
		if data != nil {
			types.ZVAL_COPY(intern.GetData(), data)
		}
		if intern.GetInnerIterator().GetFuncs().GetGetCurrentKey() != nil {
			intern.GetInnerIterator().GetFuncs().GetGetCurrentKey()(intern.GetInnerIterator(), intern.GetKey())
			if zend.EG__().GetException() != nil {
				zend.ZvalPtrDtor(intern.GetKey())
				intern.GetKey().SetUndef()
			}
		} else {
			intern.GetKey().SetLong(intern.GetPos())
		}
		if zend.EG__().GetException() != nil {
			return types.FAILURE
		} else {
			return types.SUCCESS
		}
	}
	return types.FAILURE
}
func SplDualItNext(intern *SplDualItObject, do_free int) {
	if do_free != 0 {
		SplDualItFree(intern)
	} else if intern.GetInnerIterator() == nil {
		zend.ZendThrowError(nil, "The inner constructor wasn't initialized with an iterator instance")
		return
	}
	intern.GetInnerIterator().GetFuncs().GetMoveForward()(intern.GetInnerIterator())
	intern.GetPos()++
}
func ZimSplDualItRewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplDualItFetch(intern, 1)
}
func ZimSplDualItValid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, intern.GetData().GetType() != types.IS_UNDEF)
	return
}
func ZimSplDualItKey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetKey().GetType() != types.IS_UNDEF {
		var value *types.Zval = intern.GetKey()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func ZimSplDualItCurrent(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetData().GetType() != types.IS_UNDEF {
		var value *types.Zval = intern.GetData()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func ZimSplDualItNext(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	SplDualItFetch(intern, 1)
}
func SplFilterItFetch(zthis *types.Zval, intern *SplDualItObject) {
	var retval types.Zval
	for SplDualItFetch(intern, 1) == types.SUCCESS {
		zend.ZendCallMethodWith0Params(zthis, intern.GetStd().GetCe(), nil, "accept", &retval)
		if retval.GetType() != types.IS_UNDEF {
			if zend.ZendIsTrue(&retval) != 0 {
				zend.ZvalPtrDtor(&retval)
				return
			}
			zend.ZvalPtrDtor(&retval)
		}
		if zend.EG__().GetException() != nil {
			return
		}
		intern.GetInnerIterator().GetFuncs().GetMoveForward()(intern.GetInnerIterator())
	}
	SplDualItFree(intern)
}
func SplFilterItRewind(zthis *types.Zval, intern *SplDualItObject) {
	SplDualItRewind(intern)
	SplFilterItFetch(zthis, intern)
}
func SplFilterItNext(zthis *types.Zval, intern *SplDualItObject) {
	SplDualItNext(intern, 1)
	SplFilterItFetch(zthis, intern)
}
func zim_spl_FilterIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItRewind(zend.ZEND_THIS(executeData), intern)
}
func zim_spl_FilterIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItNext(zend.ZEND_THIS(executeData), intern)
}
func zim_spl_RecursiveCallbackFilterIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RecursiveCallbackFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveCallbackFilterIterator)
}
func zim_spl_RecursiveFilterIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RecursiveFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveFilterIterator)
}
func zim_spl_RecursiveFilterIterator_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var retval types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "haschildren", &retval)
	if retval.GetType() != types.IS_UNDEF {
		zend.ZVAL_ZVAL(return_value, &retval, 0, 1)
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func zim_spl_RecursiveFilterIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var retval types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().GetException() == nil && retval.GetType() != types.IS_UNDEF {
		SplInstantiateArgEx1(types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), return_value, &retval)
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveCallbackFilterIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var retval types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().GetException() == nil && retval.GetType() != types.IS_UNDEF {
		SplInstantiateArgEx2(types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), return_value, &retval, intern.GetCbfilter().GetFci().GetFunctionName())
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_ParentIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_ParentIterator, spl_ce_RecursiveIterator, DIT_ParentIterator)
}
func zim_spl_RegexIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RegexIterator, zend.ZendCeIterator, DIT_RegexIterator)
}
func zim_spl_CallbackFilterIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	var fci *zend.ZendFcallInfo = intern.GetCbfilter().GetFci()
	var fcc *zend.ZendFcallInfoCache = intern.GetCbfilter().GetFcc()
	var params []types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetData().IsType(types.IS_UNDEF) || intern.GetKey().IsType(types.IS_UNDEF) {
		return_value.SetFalse()
		return
	}
	types.ZVAL_COPY_VALUE(&params[0], intern.GetData())
	types.ZVAL_COPY_VALUE(&params[1], intern.GetKey())
	types.ZVAL_COPY_VALUE(&params[2], intern.GetZobject())
	fci.SetRetval(return_value)
	fci.SetParamCount(3)
	fci.SetParams(params)
	fci.SetNoSeparation(0)
	if zend.ZendCallFunction(fci, fcc) != types.SUCCESS || return_value.IsUndef() {
		return_value.SetFalse()
		return
	}
	if zend.EG__().GetException() != nil {
		return_value.SetNull()
		return
	}

	/* zend_call_function may change args to IS_REF */

	types.ZVAL_COPY_VALUE(intern.GetData(), &params[0])
	types.ZVAL_COPY_VALUE(intern.GetKey(), &params[1])
}
func zim_spl_RegexIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var result *types.ZendString
	var subject *types.ZendString
	var count int = 0
	var zcount types.Zval
	var rv types.Zval
	var match_data *pcre2_match_data
	var re *pcre2_code
	var rc int
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetData().IsType(types.IS_UNDEF) {
		return_value.SetFalse()
		return
	}
	if intern.IsUseKey() {
		subject = zend.ZvalGetString(intern.GetKey())
	} else {
		if intern.GetData().IsType(types.IS_ARRAY) {
			return_value.SetFalse()
			return
		}
		subject = zend.ZvalGetString(intern.GetData())
	}

	/* Exception during string conversion. */

	if zend.EG__().GetException() != nil {
		return
	}
	switch intern.GetMode() {
	case REGIT_MODE_MAX:
		fallthrough
	case REGIT_MODE_MATCH:
		re = php_pcre_pce_re(intern.GetPce())
		match_data = php_pcre_create_match_data(0, re)
		if match_data == nil {
			return_value.SetFalse()
			return
		}
		rc = pcre2_match(re, PCRE2_SPTR(subject.GetVal()), subject.GetLen(), 0, 0, match_data, php_pcre_mctx())
		types.ZVAL_BOOL(return_value, rc >= 0)
		php_pcre_free_match_data(match_data)
	case REGIT_MODE_ALL_MATCHES:
		fallthrough
	case REGIT_MODE_GET_MATCH:
		zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
		php_pcre_match_impl(intern.GetPce(), subject, &zcount, intern.GetData(), intern.GetMode() == REGIT_MODE_ALL_MATCHES, intern.GetUseFlags(), intern.GetPregFlags(), 0)
		types.ZVAL_BOOL(return_value, zcount.GetLval() > 0)
	case REGIT_MODE_SPLIT:
		zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
		php_pcre_split_impl(intern.GetPce(), subject, intern.GetData(), -1, intern.GetPregFlags())
		count = types.Z_ARRVAL(intern.GetData()).GetNNumOfElements()
		types.ZVAL_BOOL(return_value, count > 1)
	case REGIT_MODE_REPLACE:
		var replacement *types.Zval = zend.ZendReadProperty(intern.GetStd().GetCe(), zend.ZEND_THIS(executeData), "replacement", b.SizeOf("\"replacement\"")-1, 1, &rv)
		var replacement_str *types.ZendString = zend.ZvalTryGetString(replacement)
		if replacement_str == nil {
			return
		}
		result = php_pcre_replace_impl(intern.GetPce(), subject, subject.GetVal(), subject.GetLen(), replacement_str, -1, &count)
		if intern.IsUseKey() {
			zend.ZvalPtrDtor(intern.GetKey())
			intern.GetKey().SetString(result)
		} else {
			zend.ZvalPtrDtor(intern.GetData())
			intern.GetData().SetString(result)
		}
		types.ZendStringRelease(replacement_str)
		types.ZVAL_BOOL(return_value, count > 0)
	}
	if intern.IsInverted() {
		types.ZVAL_BOOL(return_value, return_value.GetType() != types.IS_TRUE)
	}
	types.ZendStringReleaseEx(subject, 0)
}
func zim_spl_RegexIterator_getRegex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	return_value.SetStringCopy(intern.GetURegexRegex())
	return
}
func zim_spl_RegexIterator_getMode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	return_value.SetLong(intern.GetMode())
	return
}
func zim_spl_RegexIterator_setMode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var mode zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &mode) == types.FAILURE {
		return
	}
	if mode < 0 || mode >= REGIT_MODE_MAX {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+zend.ZEND_LONG_FMT, mode)
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetMode(mode)
}
func zim_spl_RegexIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	return_value.SetLong(intern.GetURegexFlags())
	return
}
func zim_spl_RegexIterator_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &flags) == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetURegexFlags(flags)
}
func zim_spl_RegexIterator_getPregFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetUseFlags() != 0 {
		return_value.SetLong(intern.GetPregFlags())
		return
	} else {
		return_value.SetLong(0)
		return
	}
}
func zim_spl_RegexIterator_setPregFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var preg_flags zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &preg_flags) == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetPregFlags(preg_flags)
	intern.SetUseFlags(1)
}
func zim_spl_RecursiveRegexIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RecursiveRegexIterator, spl_ce_RecursiveIterator, DIT_RecursiveRegexIterator)
}
func zim_spl_RecursiveRegexIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var retval types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().GetException() == nil {
		var args []types.Zval
		types.ZVAL_COPY(&args[0], &retval)
		args[1].SetStringCopy(intern.GetURegexRegex())
		args[2].SetLong(intern.GetMode())
		args[3].SetLong(intern.GetURegexFlags())
		args[4].SetLong(intern.GetPregFlags())
		SplInstantiateArgN(types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), return_value, 5, args)
		zend.ZvalPtrDtor(&args[0])
		zend.ZvalPtrDtor(&args[1])
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveRegexIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetData().IsType(types.IS_UNDEF) {
		return_value.SetFalse()
		return
	} else if intern.GetData().IsType(types.IS_ARRAY) {
		types.ZVAL_BOOL(return_value, types.Z_ARRVAL(intern.GetData()).GetNNumOfElements() > 0)
		return
	}
	zend.ZendCallMethodWith0Params(zend.ZEND_THIS(executeData), spl_ce_RegexIterator, nil, "accept", return_value)
}
func SplDualItDtor(_object *types.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	SplDualItFree(object)
	if object.GetInnerIterator() != nil {
		zend.ZendIteratorDtor(object.GetInnerIterator())
	}
}
func SplDualItFreeStorage(_object *types.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)
	if !(object.GetZobject().IsUndef()) {
		zend.ZvalPtrDtor(object.GetZobject())
	}
	if object.GetDitType() == DIT_AppendIterator {
		zend.ZendIteratorDtor(object.GetUAppendIterator())
		if object.GetZarrayit().GetType() != types.IS_UNDEF {
			zend.ZvalPtrDtor(object.GetZarrayit())
		}
	}
	if object.GetDitType() == DIT_CachingIterator || object.GetDitType() == DIT_RecursiveCachingIterator {
		zend.ZvalPtrDtor(object.GetZcache())
	}
	if object.GetDitType() == DIT_RegexIterator || object.GetDitType() == DIT_RecursiveRegexIterator {
		if object.GetPce() != nil {
			php_pcre_pce_decref(object.GetPce())
		}
		if object.GetURegexRegex() != nil {
			types.ZendStringReleaseEx(object.GetURegexRegex(), 0)
		}
	}
	if object.GetDitType() == DIT_CallbackFilterIterator || object.GetDitType() == DIT_RecursiveCallbackFilterIterator {
		if object.GetCbfilter() != nil {
			var cbfilter *_spl_cbfilter_it_intern = object.GetCbfilter()
			object.SetCbfilter(nil)
			zend.ZvalPtrDtor(cbfilter.GetFci().GetFunctionName())
			if cbfilter.GetFci().GetObject() != nil {
				zend.OBJ_RELEASE(cbfilter.GetFci().GetObject())
			}
			zend.Efree(cbfilter)
		}
	}
	zend.ZendObjectStdDtor(object.GetStd())
}
func SplDualItNew(class_type *zend.ZendClassEntry) *types.ZendObject {
	var intern *SplDualItObject
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_dual_it_object"), class_type)
	intern.SetDitType(DIT_Unknown)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	intern.GetStd().SetHandlers(&SplHandlersDualIt)
	return intern.GetStd()
}
func SplLimitItValid(intern *SplDualItObject) int {
	/* FAILURE / SUCCESS */

	if intern.GetCount() != -1 && intern.GetPos() >= intern.GetOffset()+intern.GetCount() {
		return types.FAILURE
	} else {
		return SplDualItValid(intern)
	}

	/* FAILURE / SUCCESS */
}
func SplLimitItSeek(intern *SplDualItObject, pos zend.ZendLong) {
	var zpos types.Zval
	SplDualItFree(intern)
	if pos < intern.GetOffset() {
		zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Cannot seek to "+zend.ZEND_LONG_FMT+" which is below the offset "+zend.ZEND_LONG_FMT, pos, intern.GetOffset())
		return
	}
	if pos >= intern.GetOffset()+intern.GetCount() && intern.GetCount() != -1 {
		zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Cannot seek to "+zend.ZEND_LONG_FMT+" which is behind offset "+zend.ZEND_LONG_FMT+" plus count "+zend.ZEND_LONG_FMT, pos, intern.GetOffset(), intern.GetCount())
		return
	}
	if pos != intern.GetPos() && zend.InstanceofFunction(intern.GetCe(), spl_ce_SeekableIterator) != 0 {
		zpos.SetLong(pos)
		SplDualItFree(intern)
		zend.ZendCallMethodWith1Params(intern.GetZobject(), intern.GetCe(), nil, "seek", nil, &zpos)
		if zend.EG__().GetException() == nil {
			intern.SetPos(pos)
			if SplLimitItValid(intern) == types.SUCCESS {
				SplDualItFetch(intern, 0)
			}
		}
	} else {

		/* emulate the forward seek, by next() calls */

		if pos < intern.GetPos() {
			SplDualItRewind(intern)
		}
		for pos > intern.GetPos() && SplDualItValid(intern) == types.SUCCESS {
			SplDualItNext(intern, 1)
		}
		if SplDualItValid(intern) == types.SUCCESS {
			SplDualItFetch(intern, 1)
		}
	}
}
func zim_spl_LimitIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_LimitIterator, zend.ZendCeIterator, DIT_LimitIterator)
}
func zim_spl_LimitIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplLimitItSeek(intern, intern.GetOffset())
}
func zim_spl_LimitIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it

	/*    RETURN_BOOL(spl_limit_it_valid(intern) == SUCCESS);*/

	types.ZVAL_BOOL(return_value, (intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount()) && intern.GetData().GetType() != types.IS_UNDEF)
	return
}
func zim_spl_LimitIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	if intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount() {
		SplDualItFetch(intern, 1)
	}
}
func zim_spl_LimitIterator_seek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var pos zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &pos) == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplLimitItSeek(intern, pos)
	return_value.SetLong(intern.GetPos())
	return
}
func zim_spl_LimitIterator_getPosition(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	return_value.SetLong(intern.GetPos())
	return
}
func SplCachingItValid(intern *SplDualItObject) int {
	if intern.IsValid() {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SplCachingItHasNext(intern *SplDualItObject) int { return SplDualItValid(intern) }
func SplCachingItNext(intern *SplDualItObject) {
	if SplDualItFetch(intern, 1) == types.SUCCESS {
		intern.SetIsValid(true)

		/* Full cache ? */

		if intern.IsFullCache() {
			var key *types.Zval = intern.GetKey()
			var data *types.Zval = intern.GetData()
			data = types.ZVAL_DEREF(data)
			data.TryAddRefcount()
			zend.ArraySetZvalKey(intern.GetZcache().GetArr(), key, data)
			zend.ZvalPtrDtor(data)
		}

		/* Recursion ? */

		if intern.GetDitType() == DIT_RecursiveCachingIterator {
			var retval types.Zval
			var zchildren types.Zval
			var zflags types.Zval
			zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "haschildren", &retval)
			if zend.EG__().GetException() != nil {
				zend.ZvalPtrDtor(&retval)
				if intern.IsCatchGetChild() {
					zend.ZendClearException()
				} else {
					return
				}
			} else {
				if zend.ZendIsTrue(&retval) != 0 {
					zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &zchildren)
					if zend.EG__().GetException() != nil {
						zend.ZvalPtrDtor(&zchildren)
						if intern.IsCatchGetChild() {
							zend.ZendClearException()
						} else {
							zend.ZvalPtrDtor(&retval)
							return
						}
					} else {
						zflags.SetLong(intern.GetUCachingFlags() & CIT_PUBLIC)
						SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, intern.GetZchildren(), &zchildren, &zflags)
						zend.ZvalPtrDtor(&zchildren)
					}
				}
				zend.ZvalPtrDtor(&retval)
				if zend.EG__().GetException() != nil {
					if intern.IsCatchGetChild() {
						zend.ZendClearException()
					} else {
						return
					}
				}
			}
		}
		if intern.HasUCachingFlags(CIT_TOSTRING_USE_INNER | CIT_CALL_TOSTRING) {
			var use_copy int
			var expr_copy types.Zval
			if intern.IsTostringUseInner() {
				types.ZVAL_COPY_VALUE(intern.GetZstr(), intern.GetZobject())
			} else {
				types.ZVAL_COPY_VALUE(intern.GetZstr(), intern.GetData())
			}
			use_copy = zend.ZendMakePrintableZval(intern.GetZstr(), &expr_copy)
			if use_copy != 0 {
				types.ZVAL_COPY_VALUE(intern.GetZstr(), &expr_copy)
			} else {
				intern.GetZstr().TryAddRefcount()
			}
		}
		SplDualItNext(intern, 0)
	} else {
		intern.SetIsValid(false)
	}
}
func SplCachingItRewind(intern *SplDualItObject) {
	SplDualItRewind(intern)
	intern.GetZcache().GetArr().Clean()
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_CachingIterator, zend.ZendCeIterator, DIT_CachingIterator)
}
func zim_spl_CachingIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItRewind(intern)
}
func zim_spl_CachingIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, SplCachingItValid(intern) == types.SUCCESS)
	return
}
func zim_spl_CachingIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator_hasNext(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, SplCachingItHasNext(intern) == types.SUCCESS)
	return
}
func zim_spl_CachingIterator___toString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.HasUCachingFlags(CIT_CALL_TOSTRING | CIT_TOSTRING_USE_KEY | CIT_TOSTRING_USE_CURRENT | CIT_TOSTRING_USE_INNER) {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not fetch string value (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	if intern.IsTostringUseKey() {
		types.ZVAL_COPY(return_value, intern.GetKey())
		zend.ConvertToString(return_value)
		return
	} else if intern.IsTostringUseCurrent() {
		types.ZVAL_COPY(return_value, intern.GetData())
		zend.ConvertToString(return_value)
		return
	}
	if intern.GetZstr().IsType(types.IS_STRING) {
		return_value.SetStringCopy(intern.GetZstr().GetStr())
		return
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
}
func zim_spl_CachingIterator_offsetSet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.ZendString
	var value *types.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "Sz", &key, &value) == types.FAILURE {
		return
	}
	value.TryAddRefcount()
	intern.GetZcache().GetArr().SymtableUpdate(key.GetStr(), value)
}
func zim_spl_CachingIterator_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.ZendString
	var value *types.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	if b.Assign(&value, intern.GetZcache().GetArr().SymtableFind(key.GetStr())) == nil {
		zend.ZendError(zend.E_NOTICE, "Undefined index: %s", key.GetVal())
		return
	}
	types.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_CachingIterator_offsetUnset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.ZendString
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	intern.GetZcache().GetArr().SymtableDel(key.GetStr())
}
func zim_spl_CachingIterator_offsetExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.ZendString
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, intern.GetZcache().GetArr().SymtableExists(key.GetStr()))
	return
}
func zim_spl_CachingIterator_getCache(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	types.ZVAL_COPY(return_value, intern.GetZcache())
}
func zim_spl_CachingIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	return_value.SetLong(intern.GetUCachingFlags())
	return
}
func zim_spl_CachingIterator_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &flags) == types.FAILURE {
		return
	}
	if SplCitCheckFlags(flags) != types.SUCCESS {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
		return
	}
	if intern.IsCallTostring() && (flags&CIT_CALL_TOSTRING) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag CALL_TO_STRING is not possible", 0)
		return
	}
	if intern.IsTostringUseInner() && (flags&CIT_TOSTRING_USE_INNER) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag TOSTRING_USE_INNER is not possible", 0)
		return
	}
	if (flags&CIT_FULL_CACHE) != 0 && !intern.IsFullCache() {

		/* clear on (re)enable */

		intern.GetZcache().GetArr().Clean()

		/* clear on (re)enable */

	}
	intern.SetUCachingFlags(intern.GetUCachingFlags() & ^CIT_PUBLIC | flags&CIT_PUBLIC)
}
func zim_spl_CachingIterator_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !intern.IsFullCache() {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(zend.ZEND_THIS(executeData)).GetName().GetVal())
		return
	}
	return_value.SetLong(types.Z_ARRVAL(intern.GetZcache()).GetNNumOfElements())
	return
}
func zim_spl_RecursiveCachingIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RecursiveCachingIterator, spl_ce_RecursiveIterator, DIT_RecursiveCachingIterator)
}
func zim_spl_RecursiveCachingIterator_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, intern.GetZchildren().GetType() != types.IS_UNDEF)
	return
}
func zim_spl_RecursiveCachingIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetZchildren().GetType() != types.IS_UNDEF {
		var value *types.Zval = intern.GetZchildren()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_IteratorIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_IteratorIterator, zend.ZendCeTraversable, DIT_IteratorIterator)
}
func zim_spl_NoRewindIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_NoRewindIterator, zend.ZendCeIterator, DIT_NoRewindIterator)
}
func zim_spl_NoRewindIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_NoRewindIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, intern.GetInnerIterator().GetFuncs().GetValid()(intern.GetInnerIterator()) == types.SUCCESS)
	return
}
func zim_spl_NoRewindIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetInnerIterator().GetFuncs().GetGetCurrentKey() != nil {
		intern.GetInnerIterator().GetFuncs().GetGetCurrentKey()(intern.GetInnerIterator(), return_value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_NoRewindIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var data *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	data = intern.GetInnerIterator().GetFuncs().GetGetCurrentData()(intern.GetInnerIterator())
	if data != nil {
		types.ZVAL_COPY_DEREF(return_value, data)
	}
}
func zim_spl_NoRewindIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetInnerIterator().GetFuncs().GetMoveForward()(intern.GetInnerIterator())
}
func zim_spl_InfiniteIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_InfiniteIterator, zend.ZendCeIterator, DIT_InfiniteIterator)
}
func zim_spl_InfiniteIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	if SplDualItValid(intern) == types.SUCCESS {
		SplDualItFetch(intern, 0)
	} else {
		SplDualItRewind(intern)
		if SplDualItValid(intern) == types.SUCCESS {
			SplDualItFetch(intern, 0)
		}
	}
}
func zim_spl_EmptyIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_EmptyIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetFalse()
	return
}
func zim_spl_EmptyIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the key of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the value of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func SplAppendItNextIterator(intern *SplDualItObject) int {
	SplDualItFree(intern)
	if !(intern.GetZobject().IsUndef()) {
		zend.ZvalPtrDtor(intern.GetZobject())
		intern.GetZobject().SetUndef()
		intern.SetCe(nil)
		if intern.GetInnerIterator() != nil {
			zend.ZendIteratorDtor(intern.GetInnerIterator())
			intern.SetInnerIterator(nil)
		}
	}
	if intern.GetUAppendIterator().GetFuncs().GetValid()(intern.GetUAppendIterator()) == types.SUCCESS {
		var it *types.Zval
		it = intern.GetUAppendIterator().GetFuncs().GetGetCurrentData()(intern.GetUAppendIterator())
		types.ZVAL_COPY(intern.GetZobject(), it)
		intern.SetCe(types.Z_OBJCE_P(it))
		intern.SetInnerIterator(intern.GetCe().GetGetIterator()(intern.GetCe(), it, 0))
		SplDualItRewind(intern)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SplAppendItFetch(intern *SplDualItObject) {
	for SplDualItValid(intern) != types.SUCCESS {
		intern.GetUAppendIterator().GetFuncs().GetMoveForward()(intern.GetUAppendIterator())
		if SplAppendItNextIterator(intern) != types.SUCCESS {
			return
		}
	}
	SplDualItFetch(intern, 0)
}
func SplAppendItNext(intern *SplDualItObject) {
	if SplDualItValid(intern) == types.SUCCESS {
		SplDualItNext(intern, 1)
	}
	SplAppendItFetch(intern)
}
func zim_spl_AppendIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_AppendIterator, zend.ZendCeIterator, DIT_AppendIterator)
}
func zim_spl_AppendIterator_append(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *types.Zval
	var it__1 *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it__1.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it__1
	if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, executeData.NumArgs(), "O", &it, zend.ZendCeIterator) == types.FAILURE {
		return
	}
	if intern.GetUAppendIterator().GetFuncs().GetValid()(intern.GetUAppendIterator()) == types.SUCCESS && SplDualItValid(intern) != types.SUCCESS {
		SplArrayIteratorAppend(intern.GetZarrayit(), it)
		intern.GetUAppendIterator().GetFuncs().GetMoveForward()(intern.GetUAppendIterator())
	} else {
		SplArrayIteratorAppend(intern.GetZarrayit(), it)
	}
	if intern.GetInnerIterator() == nil || SplDualItValid(intern) != types.SUCCESS {
		if intern.GetUAppendIterator().GetFuncs().GetValid()(intern.GetUAppendIterator()) != types.SUCCESS {
			intern.GetUAppendIterator().GetFuncs().GetRewind()(intern.GetUAppendIterator())
		}
		for {
			SplAppendItNextIterator(intern)
			if intern.GetZobject().GetObj() == it.GetObj() {
				break
			}
		}
		SplAppendItFetch(intern)
	}
}
func zim_spl_AppendIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItFetch(intern, 1)
	if intern.GetData().GetType() != types.IS_UNDEF {
		var value *types.Zval = intern.GetData()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_AppendIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetUAppendIterator().GetFuncs().GetRewind()(intern.GetUAppendIterator())
	if SplAppendItNextIterator(intern) == types.SUCCESS {
		SplAppendItFetch(intern)
	}
}
func zim_spl_AppendIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	types.ZVAL_BOOL(return_value, intern.GetData().GetType() != types.IS_UNDEF)
	return
}
func zim_spl_AppendIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplAppendItNext(intern)
}
func zim_spl_AppendIterator_getIteratorIndex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	APPENDIT_CHECK_CTOR(intern)
	SplArrayIteratorKey(intern.GetZarrayit(), return_value)
}
func zim_spl_AppendIterator_getArrayIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var value *types.Zval
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS(executeData))
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	value = intern.GetZarrayit()
	types.ZVAL_COPY_DEREF(return_value, value)
}
func SplIteratorApply(obj *types.Zval, apply_func SplIteratorApplyFuncT, puser any) int {
	var iter *zend.ZendObjectIterator
	var ce *zend.ZendClassEntry = types.Z_OBJCE_P(obj)
	iter = ce.GetGetIterator()(ce, obj, 0)
	if zend.EG__().GetException() != nil {
		goto done
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if zend.EG__().GetException() != nil {
			goto done
		}
	}
	for iter.GetFuncs().GetValid()(iter) == types.SUCCESS {
		if zend.EG__().GetException() != nil {
			goto done
		}
		if apply_func(iter, puser) == zend.ZEND_HASH_APPLY_STOP || zend.EG__().GetException() != nil {
			goto done
		}
		iter.GetIndex()++
		iter.GetFuncs().GetMoveForward()(iter)
		if zend.EG__().GetException() != nil {
			goto done
		}
	}
done:
	if iter != nil {
		zend.ZendIteratorDtor(iter)
	}
	if zend.EG__().GetException() != nil {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}
func SplIteratorToArrayApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *types.Zval
	var return_value *types.Zval = (*types.Zval)(puser)
	data = iter.GetFuncs().GetGetCurrentData()(iter)
	if zend.EG__().GetException() != nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if data == nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if iter.GetFuncs().GetGetCurrentKey() != nil {
		var key types.Zval
		iter.GetFuncs().GetGetCurrentKey()(iter, &key)
		if zend.EG__().GetException() != nil {
			return zend.ZEND_HASH_APPLY_STOP
		}
		zend.ArraySetZvalKey(return_value.GetArr(), &key, data)
		zend.ZvalPtrDtor(&key)
	} else {
		data.TryAddRefcount()
		zend.AddNextIndexZval(return_value, data)
	}
	return zend.ZEND_HASH_APPLY_KEEP
}
func SplIteratorToValuesApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *types.Zval
	var return_value *types.Zval = (*types.Zval)(puser)
	data = iter.GetFuncs().GetGetCurrentData()(iter)
	if zend.EG__().GetException() != nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if data == nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	data.TryAddRefcount()
	zend.AddNextIndexZval(return_value, data)
	return zend.ZEND_HASH_APPLY_KEEP
}
func ZifIteratorToArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var use_keys types.ZendBool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "O|b", &obj, zend.ZendCeTraversable, &use_keys) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	SplIteratorApply(obj, b.Cond(use_keys != 0, SplIteratorToArrayApply, SplIteratorToValuesApply), any(return_value))
}
func SplIteratorCountApply(iter *zend.ZendObjectIterator, puser any) int {
	*((*zend.ZendLong)(puser))++
	return zend.ZEND_HASH_APPLY_KEEP
}
func ZifIteratorCount(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var count zend.ZendLong = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "O", &obj, zend.ZendCeTraversable) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if SplIteratorApply(obj, SplIteratorCountApply, any(&count)) == types.FAILURE {
		return
	}
	return_value.SetLong(count)
	return
}
func SplIteratorFuncApply(iter *zend.ZendObjectIterator, puser any) int {
	var retval types.Zval
	var apply_info *SplIteratorApplyInfo = (*SplIteratorApplyInfo)(puser)
	var result int
	apply_info.GetCount()++
	zend.ZendFcallInfoCall(apply_info.GetFci(), apply_info.GetFcc(), &retval, nil)
	if zend.ZendIsTrue(&retval) != 0 {
		result = zend.ZEND_HASH_APPLY_KEEP
	} else {
		result = zend.ZEND_HASH_APPLY_STOP
	}
	zend.ZvalPtrDtor(&retval)
	return result
}
func ZifIteratorApply(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var apply_info SplIteratorApplyInfo
	apply_info.SetArgs(nil)
	if zend.ZendParseParameters(executeData.NumArgs(), "Of|a!", apply_info.GetObj(), zend.ZendCeTraversable, apply_info.GetFci(), apply_info.GetFcc(), apply_info.GetArgs()) == types.FAILURE {
		return
	}
	apply_info.SetCount(0)
	zend.ZendFcallInfoArgs(apply_info.GetFci(), apply_info.GetArgs())
	if SplIteratorApply(apply_info.GetObj(), SplIteratorFuncApply, any(&apply_info)) == types.FAILURE {
		zend.ZendFcallInfoArgs(apply_info.GetFci(), nil)
		return
	}
	zend.ZendFcallInfoArgs(apply_info.GetFci(), nil)
	return_value.SetLong(apply_info.GetCount())
	return
}
func ZmStartupSplIterators(type_ int, module_number int) int {
	SplRegisterInterface(&spl_ce_RecursiveIterator, "RecursiveIterator", spl_funcs_RecursiveIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIterator, 1, zend.ZendCeIterator)
	SplRegisterStdClass(&spl_ce_RecursiveIteratorIterator, "RecursiveIteratorIterator", spl_RecursiveIteratorIterator_new, spl_funcs_RecursiveIteratorIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIteratorIterator, 1, zend.ZendCeIterator)
	memcpy(&SplHandlersRecItIt, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplHandlersRecItIt.SetOffset(zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd())) - (*byte)(nil)))
	SplHandlersRecItIt.SetGetMethod(SplRecursiveItGetMethod)
	SplHandlersRecItIt.SetCloneObj(nil)
	SplHandlersRecItIt.SetDtorObj(spl_RecursiveIteratorIterator_dtor)
	SplHandlersRecItIt.SetFreeObj(spl_RecursiveIteratorIterator_free_storage)
	memcpy(&SplHandlersDualIt, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplHandlersDualIt.SetOffset(zend_long((*byte)(&((*SplDualItObject)(nil).GetStd())) - (*byte)(nil)))
	SplHandlersDualIt.SetGetMethod(SplDualItGetMethod)

	/*spl_handlers_dual_it.call_method = spl_dual_it_call_method;*/

	SplHandlersDualIt.SetCloneObj(nil)
	SplHandlersDualIt.SetDtorObj(SplDualItDtor)
	SplHandlersDualIt.SetFreeObj(SplDualItFreeStorage)
	spl_ce_RecursiveIteratorIterator.SetGetIterator(SplRecursiveItGetIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "LEAVES_ONLY", b.SizeOf("\"LEAVES_ONLY\"")-1, zend.ZendLong(RIT_LEAVES_ONLY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "SELF_FIRST", b.SizeOf("\"SELF_FIRST\"")-1, zend.ZendLong(RIT_SELF_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CHILD_FIRST", b.SizeOf("\"CHILD_FIRST\"")-1, zend.ZendLong(RIT_CHILD_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CATCH_GET_CHILD", b.SizeOf("\"CATCH_GET_CHILD\"")-1, zend.ZendLong(RIT_CATCH_GET_CHILD))
	SplRegisterInterface(&spl_ce_OuterIterator, "OuterIterator", spl_funcs_OuterIterator)
	zend.ZendClassImplements(spl_ce_OuterIterator, 1, zend.ZendCeIterator)
	SplRegisterStdClass(&spl_ce_IteratorIterator, "IteratorIterator", SplDualItNew, spl_funcs_IteratorIterator)
	zend.ZendClassImplements(spl_ce_IteratorIterator, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_IteratorIterator, 1, spl_ce_OuterIterator)
	SplRegisterSubClass(&spl_ce_FilterIterator, spl_ce_IteratorIterator, "FilterIterator", SplDualItNew, spl_funcs_FilterIterator)
	spl_ce_FilterIterator.AddCeFlags(zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)
	SplRegisterSubClass(&spl_ce_RecursiveFilterIterator, spl_ce_FilterIterator, "RecursiveFilterIterator", SplDualItNew, spl_funcs_RecursiveFilterIterator)
	zend.ZendClassImplements(spl_ce_RecursiveFilterIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_CallbackFilterIterator, spl_ce_FilterIterator, "CallbackFilterIterator", SplDualItNew, spl_funcs_CallbackFilterIterator)
	SplRegisterSubClass(&spl_ce_RecursiveCallbackFilterIterator, spl_ce_CallbackFilterIterator, "RecursiveCallbackFilterIterator", SplDualItNew, spl_funcs_RecursiveCallbackFilterIterator)
	zend.ZendClassImplements(spl_ce_RecursiveCallbackFilterIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_ParentIterator, spl_ce_RecursiveFilterIterator, "ParentIterator", SplDualItNew, spl_funcs_ParentIterator)
	SplRegisterInterface(&spl_ce_SeekableIterator, "SeekableIterator", spl_funcs_SeekableIterator)
	zend.ZendClassImplements(spl_ce_SeekableIterator, 1, zend.ZendCeIterator)
	SplRegisterSubClass(&spl_ce_LimitIterator, spl_ce_IteratorIterator, "LimitIterator", SplDualItNew, spl_funcs_LimitIterator)
	SplRegisterSubClass(&spl_ce_CachingIterator, spl_ce_IteratorIterator, "CachingIterator", SplDualItNew, spl_funcs_CachingIterator)
	zend.ZendClassImplements(spl_ce_CachingIterator, 1, spl_ce_ArrayAccess)
	zend.ZendClassImplements(spl_ce_CachingIterator, 1, spl_ce_Countable)
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CALL_TOSTRING", b.SizeOf("\"CALL_TOSTRING\"")-1, zend.ZendLong(CIT_CALL_TOSTRING))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CATCH_GET_CHILD", b.SizeOf("\"CATCH_GET_CHILD\"")-1, zend.ZendLong(CIT_CATCH_GET_CHILD))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_KEY", b.SizeOf("\"TOSTRING_USE_KEY\"")-1, zend.ZendLong(CIT_TOSTRING_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_CURRENT", b.SizeOf("\"TOSTRING_USE_CURRENT\"")-1, zend.ZendLong(CIT_TOSTRING_USE_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_INNER", b.SizeOf("\"TOSTRING_USE_INNER\"")-1, zend.ZendLong(CIT_TOSTRING_USE_INNER))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "FULL_CACHE", b.SizeOf("\"FULL_CACHE\"")-1, zend.ZendLong(CIT_FULL_CACHE))
	SplRegisterSubClass(&spl_ce_RecursiveCachingIterator, spl_ce_CachingIterator, "RecursiveCachingIterator", SplDualItNew, spl_funcs_RecursiveCachingIterator)
	zend.ZendClassImplements(spl_ce_RecursiveCachingIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterSubClass(&spl_ce_NoRewindIterator, spl_ce_IteratorIterator, "NoRewindIterator", SplDualItNew, spl_funcs_NoRewindIterator)
	SplRegisterSubClass(&spl_ce_AppendIterator, spl_ce_IteratorIterator, "AppendIterator", SplDualItNew, spl_funcs_AppendIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIteratorIterator, 1, spl_ce_OuterIterator)
	SplRegisterSubClass(&spl_ce_InfiniteIterator, spl_ce_IteratorIterator, "InfiniteIterator", SplDualItNew, spl_funcs_InfiniteIterator)
	SplRegisterSubClass(&spl_ce_RegexIterator, spl_ce_FilterIterator, "RegexIterator", SplDualItNew, spl_funcs_RegexIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "USE_KEY", b.SizeOf("\"USE_KEY\"")-1, zend.ZendLong(REGIT_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "INVERT_MATCH", b.SizeOf("\"INVERT_MATCH\"")-1, zend.ZendLong(REGIT_INVERTED))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "MATCH", b.SizeOf("\"MATCH\"")-1, zend.ZendLong(REGIT_MODE_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "GET_MATCH", b.SizeOf("\"GET_MATCH\"")-1, zend.ZendLong(REGIT_MODE_GET_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "ALL_MATCHES", b.SizeOf("\"ALL_MATCHES\"")-1, zend.ZendLong(REGIT_MODE_ALL_MATCHES))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "SPLIT", b.SizeOf("\"SPLIT\"")-1, zend.ZendLong(REGIT_MODE_SPLIT))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "REPLACE", b.SizeOf("\"REPLACE\"")-1, zend.ZendLong(REGIT_MODE_REPLACE))
	SplRegisterProperty(spl_ce_RegexIterator, "replacement", b.SizeOf("\"replacement\"")-1, 0)
	SplRegisterSubClass(&spl_ce_RecursiveRegexIterator, spl_ce_RegexIterator, "RecursiveRegexIterator", SplDualItNew, spl_funcs_RecursiveRegexIterator)
	zend.ZendClassImplements(spl_ce_RecursiveRegexIterator, 1, spl_ce_RecursiveIterator)
	SplRegisterStdClass(&spl_ce_EmptyIterator, "EmptyIterator", nil, spl_funcs_EmptyIterator)
	zend.ZendClassImplements(spl_ce_EmptyIterator, 1, zend.ZendCeIterator)
	SplRegisterSubClass(&spl_ce_RecursiveTreeIterator, spl_ce_RecursiveIteratorIterator, "RecursiveTreeIterator", spl_RecursiveTreeIterator_new, spl_funcs_RecursiveTreeIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_CURRENT", b.SizeOf("\"BYPASS_CURRENT\"")-1, zend.ZendLong(RTIT_BYPASS_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_KEY", b.SizeOf("\"BYPASS_KEY\"")-1, zend.ZendLong(RTIT_BYPASS_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_LEFT", b.SizeOf("\"PREFIX_LEFT\"")-1, zend.ZendLong(0))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_HAS_NEXT", b.SizeOf("\"PREFIX_MID_HAS_NEXT\"")-1, zend.ZendLong(1))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_LAST", b.SizeOf("\"PREFIX_MID_LAST\"")-1, zend.ZendLong(2))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_HAS_NEXT", b.SizeOf("\"PREFIX_END_HAS_NEXT\"")-1, zend.ZendLong(3))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_LAST", b.SizeOf("\"PREFIX_END_LAST\"")-1, zend.ZendLong(4))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_RIGHT", b.SizeOf("\"PREFIX_RIGHT\"")-1, zend.ZendLong(5))
	return types.SUCCESS
}
