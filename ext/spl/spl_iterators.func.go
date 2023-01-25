// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func SplDualItFromObj(obj *zend.ZendObject) *SplDualItObject {
	return (*SplDualItObject)((*byte)(obj - zend_long((*byte)(&((*SplDualItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLDUAL_IT_P(zv *zend.Zval) *SplDualItObject {
	return SplDualItFromObj(zend.Z_OBJ_P(zv))
}
func SplRecursiveItFromObj(obj *zend.ZendObject) *SplRecursiveItObject {
	return (*SplRecursiveItObject)((*byte)(obj - zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLRECURSIVE_IT_P(zv *zend.Zval) *SplRecursiveItObject {
	return SplRecursiveItFromObj(zend.Z_OBJ_P(zv))
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
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(&iter.intern.data)
	var sub_iter *zend.ZendObjectIterator
	for object.GetLevel() > 0 {
		if !(zend.Z_ISUNDEF(object.GetIterators()[object.GetLevel()].GetZobject())) {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&object.iterators[object.GetLevel()].GetZobject())
		}
		object.GetLevel()--
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.SetLevel(0)
	zend.ZvalPtrDtor(&iter.intern.data)
}
func SplRecursiveItValidEx(object *SplRecursiveItObject, zthis *zend.Zval) int {
	var sub_iter *zend.ZendObjectIterator
	var level int = object.GetLevel()
	if object.GetIterators() == nil {
		return zend.FAILURE
	}
	for level >= 0 {
		sub_iter = object.GetIterators()[level].GetIterator()
		if sub_iter.funcs.valid(sub_iter) == zend.SUCCESS {
			return zend.SUCCESS
		}
		level--
	}
	if object.GetEndIteration() != nil && object.GetInIteration() != 0 {
		zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.endIteration, "endIteration", nil)
	}
	object.SetInIteration(0)
	return zend.FAILURE
}
func SplRecursiveItValid(iter *zend.ZendObjectIterator) int {
	return SplRecursiveItValidEx(Z_SPLRECURSIVE_IT_P(&iter.data), &iter.data)
}
func SplRecursiveItGetCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(&iter.data)
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	return sub_iter.funcs.get_current_data(sub_iter)
}
func SplRecursiveItGetCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(&iter.data)
	var sub_iter *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	if sub_iter.funcs.get_current_key != nil {
		sub_iter.funcs.get_current_key(sub_iter, key)
	} else {
		zend.ZVAL_LONG(key, iter.index)
	}
}
func SplRecursiveItMoveForwardEx(object *SplRecursiveItObject, zthis *zend.Zval) {
	var iterator *zend.ZendObjectIterator
	var zobject *zend.Zval
	var ce *zend.ZendClassEntry
	var retval zend.Zval
	var child zend.Zval
	var sub_iter *zend.ZendObjectIterator
	var has_children int
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	for zend.ExecutorGlobals.exception == nil {
	next_step:
		iterator = object.GetIterators()[object.GetLevel()].GetIterator()
		switch object.GetIterators()[object.GetLevel()].GetState() {
		case RS_NEXT:
			iterator.funcs.move_forward(iterator)
			if zend.ExecutorGlobals.exception != nil {
				if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
				}
			}
		case RS_START:
			if iterator.funcs.valid(iterator) == zend.FAILURE {
				break
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_TEST)
		case RS_TEST:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = &object.iterators[object.GetLevel()].GetZobject()
			if object.GetCallHasChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.callHasChildren, "callHasChildren", &retval)
			} else {
				zend.ZendCallMethodWith0Params(zobject, ce, nil, "haschildren", &retval)
			}
			if zend.ExecutorGlobals.exception != nil {
				if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					return
				} else {
					zend.ZendClearException()
				}
			}
			if zend.Z_TYPE(retval) != zend.IS_UNDEF {
				has_children = zend.ZendIsTrue(&retval)
				zend.ZvalPtrDtor(&retval)
				if has_children != 0 {
					if object.GetMaxDepth() == -1 || object.GetMaxDepth() > object.GetLevel() {
						switch object.GetMode() {
						case RIT_LEAVES_ONLY:

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
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.nextElement, "nextelement", nil)
			}
			object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			if zend.ExecutorGlobals.exception != nil {
				if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
				}
			}
			return
		case RS_SELF:
			if object.GetNextElement() != nil && (object.GetMode() == RIT_SELF_FIRST || object.GetMode() == RIT_CHILD_FIRST) {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.nextElement, "nextelement", nil)
			}
			if object.GetMode() == RIT_SELF_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_CHILD)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			return
		case RS_CHILD:
			ce = object.GetIterators()[object.GetLevel()].GetCe()
			zobject = &object.iterators[object.GetLevel()].GetZobject()
			if object.GetCallGetChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.callGetChildren, "callGetChildren", &child)
			} else {
				zend.ZendCallMethodWith0Params(zobject, ce, nil, "getchildren", &child)
			}
			if zend.ExecutorGlobals.exception != nil {
				if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
					return
				} else {
					zend.ZendClearException()
					zend.ZvalPtrDtor(&child)
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					goto next_step
				}
			}
			if zend.Z_TYPE(child) == zend.IS_UNDEF || zend.Z_TYPE(child) != zend.IS_OBJECT || !(b.Assign(&ce, zend.Z_OBJCE(child)) && zend.InstanceofFunction(ce, spl_ce_RecursiveIterator) != 0) {
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
			sub_iter = ce.get_iterator(ce, &child, 0)
			zend.ZVAL_COPY_VALUE(&object.iterators[object.GetLevel()].GetZobject(), &child)
			object.GetIterators()[object.GetLevel()].SetIterator(sub_iter)
			object.GetIterators()[object.GetLevel()].SetCe(ce)
			object.GetIterators()[object.GetLevel()].SetState(RS_START)
			if sub_iter.funcs.rewind != nil {
				sub_iter.funcs.rewind(sub_iter)
			}
			if object.GetBeginChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.beginChildren, "beginchildren", nil)
				if zend.ExecutorGlobals.exception != nil {
					if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
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
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.endChildren, "endchildren", nil)
				if zend.ExecutorGlobals.exception != nil {
					if (object.GetFlags() & RIT_CATCH_GET_CHILD) == 0 {
						return
					} else {
						zend.ZendClearException()
					}
				}
			}
			if object.GetLevel() > 0 {
				var garbage zend.Zval
				zend.ZVAL_COPY_VALUE(&garbage, &object.iterators[object.GetLevel()].GetZobject())
				zend.ZVAL_UNDEF(&object.iterators[object.GetLevel()].GetZobject())
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
func SplRecursiveItRewindEx(object *SplRecursiveItObject, zthis *zend.Zval) {
	var sub_iter *zend.ZendObjectIterator
	SPL_FETCH_SUB_ITERATOR(sub_iter, object)
	for object.GetLevel() != 0 {
		sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
		zend.ZendIteratorDtor(sub_iter)
		zend.ZvalPtrDtor(&object.iterators[b.PostDec(&(object.GetLevel()))].GetZobject())
		if zend.ExecutorGlobals.exception == nil && (object.GetEndChildren() == nil || object.GetEndChildren().common.scope != spl_ce_RecursiveIteratorIterator) {
			zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.endChildren, "endchildren", nil)
		}
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.GetIterators()[0].SetState(RS_START)
	sub_iter = object.GetIterators()[0].GetIterator()
	if sub_iter.funcs.rewind != nil {
		sub_iter.funcs.rewind(sub_iter)
	}
	if zend.ExecutorGlobals.exception == nil && object.GetBeginIteration() != nil && object.GetInIteration() == 0 {
		zend.ZendCallMethodWith0Params(zthis, object.GetCe(), &object.beginIteration, "beginIteration", nil)
	}
	object.SetInIteration(1)
	SplRecursiveItMoveForwardEx(object, zthis)
}
func SplRecursiveItMoveForward(iter *zend.ZendObjectIterator) {
	SplRecursiveItMoveForwardEx(Z_SPLRECURSIVE_IT_P(&iter.data), &iter.data)
}
func SplRecursiveItRewind(iter *zend.ZendObjectIterator) {
	SplRecursiveItRewindEx(Z_SPLRECURSIVE_IT_P(&iter.data), &iter.data)
}
func SplRecursiveItGetIterator(ce *zend.ZendClassEntry, zobject *zend.Zval, by_ref int) *zend.ZendObjectIterator {
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
	zend.Z_ADDREF_P(zobject)
	zend.ZVAL_OBJ(&iterator.intern.data, zend.Z_OBJ_P(zobject))
	iterator.intern.funcs = &SplRecursiveItIteratorFuncs
	return (*zend.ZendObjectIterator)(iterator)
}
func SplRecursiveItItConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, rit_type RecursiveItItType) {
	var object *zend.Zval = zend.ZEND_THIS
	var intern *SplRecursiveItObject
	var iterator *zend.Zval
	var ce_iterator *zend.ZendClassEntry
	var mode zend.ZendLong
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	var caching_it zend.Zval
	var aggregate_retval zend.Zval
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
	switch rit_type {
	case RIT_RecursiveTreeIterator:
		var caching_it_flags zend.Zval
		var user_caching_it_flags *zend.Zval = nil
		mode = RIT_SELF_FIRST
		flags = RTIT_BYPASS_KEY
		if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "o|lzl", &iterator, &flags, &user_caching_it_flags, &mode) == zend.SUCCESS {
			if zend.InstanceofFunction(zend.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, zend.Z_OBJCE_P(iterator), &zend.Z_OBJCE_P(iterator).iterator_funcs_ptr.zf_new_iterator, "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				zend.Z_ADDREF_P(iterator)
			}
			if user_caching_it_flags != nil {
				zend.ZVAL_COPY(&caching_it_flags, user_caching_it_flags)
			} else {
				zend.ZVAL_LONG(&caching_it_flags, CIT_CATCH_GET_CHILD)
			}
			SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &caching_it, iterator, &caching_it_flags)
			zend.ZvalPtrDtor(&caching_it_flags)
			zend.ZvalPtrDtor(iterator)
			iterator = &caching_it
		} else {
			iterator = nil
		}
		break
	case RIT_RecursiveIteratorIterator:

	default:
		mode = RIT_LEAVES_ONLY
		flags = 0
		if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "o|ll", &iterator, &mode, &flags) == zend.SUCCESS {
			if zend.InstanceofFunction(zend.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, zend.Z_OBJCE_P(iterator), &zend.Z_OBJCE_P(iterator).iterator_funcs_ptr.zf_new_iterator, "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				zend.Z_ADDREF_P(iterator)
			}
		} else {
			iterator = nil
		}
		break
	}
	if iterator == nil || zend.InstanceofFunction(zend.Z_OBJCE_P(iterator), spl_ce_RecursiveIterator) == 0 {
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
	intern.SetCe(zend.Z_OBJCE_P(object))
	intern.SetBeginIteration(zend.ZendHashStrFindPtr(&intern.ce.function_table, "beginiteration", b.SizeOf("\"beginiteration\"")-1))
	if intern.GetBeginIteration().common.scope == ce_base {
		intern.SetBeginIteration(nil)
	}
	intern.SetEndIteration(zend.ZendHashStrFindPtr(&intern.ce.function_table, "enditeration", b.SizeOf("\"enditeration\"")-1))
	if intern.GetEndIteration().common.scope == ce_base {
		intern.SetEndIteration(nil)
	}
	intern.SetCallHasChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "callhaschildren", b.SizeOf("\"callHasChildren\"")-1))
	if intern.GetCallHasChildren().common.scope == ce_base {
		intern.SetCallHasChildren(nil)
	}
	intern.SetCallGetChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "callgetchildren", b.SizeOf("\"callGetChildren\"")-1))
	if intern.GetCallGetChildren().common.scope == ce_base {
		intern.SetCallGetChildren(nil)
	}
	intern.SetBeginChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "beginchildren", b.SizeOf("\"beginchildren\"")-1))
	if intern.GetBeginChildren().common.scope == ce_base {
		intern.SetBeginChildren(nil)
	}
	intern.SetEndChildren(zend.ZendHashStrFindPtr(&intern.ce.function_table, "endchildren", b.SizeOf("\"endchildren\"")-1))
	if intern.GetEndChildren().common.scope == ce_base {
		intern.SetEndChildren(nil)
	}
	intern.SetNextElement(zend.ZendHashStrFindPtr(&intern.ce.function_table, "nextelement", b.SizeOf("\"nextElement\"")-1))
	if intern.GetNextElement().common.scope == ce_base {
		intern.SetNextElement(nil)
	}
	ce_iterator = zend.Z_OBJCE_P(iterator)
	intern.GetIterators()[0].SetIterator(ce_iterator.get_iterator(ce_iterator, iterator, 0))
	zend.ZVAL_OBJ(&intern.iterators[0].GetZobject(), zend.Z_OBJ_P(iterator))
	intern.GetIterators()[0].SetCe(ce_iterator)
	intern.GetIterators()[0].SetState(RS_START)
	zend.ZendRestoreErrorHandling(&error_handling)
	if zend.ExecutorGlobals.exception != nil {
		var sub_iter *zend.ZendObjectIterator
		for intern.GetLevel() >= 0 {
			sub_iter = intern.GetIterators()[intern.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&intern.iterators[b.PostDec(&(intern.GetLevel()))].GetZobject())
		}
		zend.Efree(intern.GetIterators())
		intern.SetIterators(nil)
	}
}
func zim_spl_RecursiveIteratorIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplRecursiveItItConstruct(execute_data, return_value, spl_ce_RecursiveIteratorIterator, zend.ZendCeIterator, RIT_RecursiveIteratorIterator)
}
func zim_spl_RecursiveIteratorIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplRecursiveItRewindEx(object, zend.ZEND_THIS)
}
func zim_spl_RecursiveIteratorIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(SplRecursiveItValidEx(object, zend.ZEND_THIS) == zend.SUCCESS)
	return
}
func zim_spl_RecursiveIteratorIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var iterator *zend.ZendObjectIterator
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	if iterator.funcs.get_current_key != nil {
		iterator.funcs.get_current_key(iterator, return_value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func zim_spl_RecursiveIteratorIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var iterator *zend.ZendObjectIterator
	var data *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	data = iterator.funcs.get_current_data(iterator)
	if data != nil {
		zend.ZVAL_COPY_DEREF(return_value, data)
	}
}
func zim_spl_RecursiveIteratorIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplRecursiveItMoveForwardEx(object, zend.ZEND_THIS)
}
func zim_spl_RecursiveIteratorIterator_getDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(object.GetLevel())
	return
}
func zim_spl_RecursiveIteratorIterator_getSubIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var level zend.ZendLong = object.GetLevel()
	var value *zend.Zval
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|l", &level) == zend.FAILURE {
		return
	}
	if level < 0 || level > object.GetLevel() {
		zend.RETVAL_NULL()
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	value = &object.iterators[level].GetZobject()
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_RecursiveIteratorIterator_getInnerIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var zobject *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	zend.ZVAL_COPY_DEREF(return_value, zobject)
}
func zim_spl_RecursiveIteratorIterator_beginIteration(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endIteration(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_callHasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry
	var zobject *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.RETVAL_NULL()
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	if zend.Z_TYPE_P(zobject) == zend.IS_UNDEF {
		zend.RETVAL_FALSE
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "haschildren", return_value)
		if zend.Z_TYPE_P(return_value) == zend.IS_UNDEF {
			zend.RETVAL_FALSE
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_callGetChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry
	var zobject *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = &object.iterators[object.GetLevel()].GetZobject()
	if zend.Z_TYPE_P(zobject) == zend.IS_UNDEF {
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "getchildren", return_value)
		if zend.Z_TYPE_P(return_value) == zend.IS_UNDEF {
			zend.RETVAL_NULL()
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_beginChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_nextElement(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_setMaxDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var max_depth zend.ZendLong = -1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|l", &max_depth) == zend.FAILURE {
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
func zim_spl_RecursiveIteratorIterator_getMaxDepth(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetMaxDepth() == -1 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(object.GetMaxDepth())
		return
	}
}
func SplRecursiveItGetMethod(zobject **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var object *SplRecursiveItObject = SplRecursiveItFromObj(*zobject)
	var level zend.ZendLong = object.GetLevel()
	var zobj *zend.Zval
	if object.GetIterators() == nil {
		core.PhpErrorDocref(nil, zend.E_ERROR, "The %s instance wasn't initialized properly", zend.ZSTR_VAL((*zobject).ce.name))
	}
	zobj = &object.iterators[level].GetZobject()
	function_handler = zend.ZendStdGetMethod(zobject, method, key)
	if function_handler == nil {
		if b.Assign(&function_handler, zend.ZendHashFindPtr(&zend.Z_OBJCE_P(zobj).function_table, method)) == nil {
			*zobject = zend.Z_OBJ_P(zobj)
			function_handler = (*zobject).handlers.get_method(zobject, method, key)
		} else {
			*zobject = zend.Z_OBJ_P(zobj)
		}
	}
	return function_handler
}
func spl_RecursiveIteratorIterator_dtor(_object *zend.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	var sub_iter *zend.ZendObjectIterator

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	if object.GetIterators() != nil {
		for object.GetLevel() >= 0 {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			zend.ZendIteratorDtor(sub_iter)
			zend.ZvalPtrDtor(&object.iterators[b.PostDec(&(object.GetLevel()))].GetZobject())
		}
		zend.Efree(object.GetIterators())
		object.SetIterators(nil)
	}
}
func spl_RecursiveIteratorIterator_free_storage(_object *zend.ZendObject) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	if object.GetIterators() != nil {
		zend.Efree(object.GetIterators())
		object.SetIterators(nil)
		object.SetLevel(0)
	}
	zend.ZendObjectStdDtor(&object.std)
	zend.SmartStrFree(&object.prefix[0])
	zend.SmartStrFree(&object.prefix[1])
	zend.SmartStrFree(&object.prefix[2])
	zend.SmartStrFree(&object.prefix[3])
	zend.SmartStrFree(&object.prefix[4])
	zend.SmartStrFree(&object.prefix[5])
	zend.SmartStrFree(&object.postfix[0])
}
func spl_RecursiveIteratorIterator_new_ex(class_type *zend.ZendClassEntry, init_prefix int) *zend.ZendObject {
	var intern *SplRecursiveItObject
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_recursive_it_object"), class_type)
	if init_prefix != 0 {
		zend.SmartStrAppendl(&intern.prefix[0], "", 0)
		zend.SmartStrAppendl(&intern.prefix[1], "| ", 2)
		zend.SmartStrAppendl(&intern.prefix[2], "  ", 2)
		zend.SmartStrAppendl(&intern.prefix[3], "|-", 2)
		zend.SmartStrAppendl(&intern.prefix[4], "\\-", 2)
		zend.SmartStrAppendl(&intern.prefix[5], "", 0)
		zend.SmartStrAppendl(&intern.postfix[0], "", 0)
	}
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.std.handlers = &SplHandlersRecItIt
	return &intern.std
}
func spl_RecursiveIteratorIterator_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 0)
}
func spl_RecursiveTreeIterator_new(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 1)
}
func SplRecursiveTreeIteratorGetPrefix(object *SplRecursiveItObject, return_value *zend.Zval) {
	var str zend.SmartStr = zend.SmartStr{0}
	var has_next zend.Zval
	var level int
	zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[0].s), zend.ZSTR_LEN(object.GetPrefix()[0].s))
	for level = 0; level < object.GetLevel(); level++ {
		zend.ZendCallMethodWith0Params(&object.iterators[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
		if zend.Z_TYPE(has_next) != zend.IS_UNDEF {
			if zend.Z_TYPE(has_next) == zend.IS_TRUE {
				zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[1].s), zend.ZSTR_LEN(object.GetPrefix()[1].s))
			} else {
				zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[2].s), zend.ZSTR_LEN(object.GetPrefix()[2].s))
			}
			zend.ZvalPtrDtor(&has_next)
		}
	}
	zend.ZendCallMethodWith0Params(&object.iterators[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
	if zend.Z_TYPE(has_next) != zend.IS_UNDEF {
		if zend.Z_TYPE(has_next) == zend.IS_TRUE {
			zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[3].s), zend.ZSTR_LEN(object.GetPrefix()[3].s))
		} else {
			zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[4].s), zend.ZSTR_LEN(object.GetPrefix()[4].s))
		}
		zend.ZvalPtrDtor(&has_next)
	}
	zend.SmartStrAppendl(&str, zend.ZSTR_VAL(object.GetPrefix()[5].s), zend.ZSTR_LEN(object.GetPrefix()[5].s))
	zend.SmartStr0(&str)
	zend.RETVAL_NEW_STR(str.s)
	return
}
func SplRecursiveTreeIteratorGetEntry(object *SplRecursiveItObject, return_value *zend.Zval) {
	var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	var data *zend.Zval
	data = iterator.funcs.get_current_data(iterator)
	if data != nil {
		zend.ZVAL_DEREF(data)

		/* TODO: Remove this special case? */

		if zend.Z_TYPE_P(data) == zend.IS_ARRAY {
			zend.RETVAL_INTERNED_STR(zend.ZSTR_KNOWN(zend.ZEND_STR_ARRAY_CAPITALIZED))
		} else {
			zend.ZVAL_COPY(return_value, data)
			zend.ConvertToString(return_value)
		}

		/* TODO: Remove this special case? */

	}
}
func SplRecursiveTreeIteratorGetPostfix(object *SplRecursiveItObject, return_value *zend.Zval) {
	zend.RETVAL_STR(object.GetPostfix()[0].s)
	zend.Z_ADDREF_P(return_value)
}
func zim_spl_RecursiveTreeIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplRecursiveItItConstruct(execute_data, return_value, spl_ce_RecursiveTreeIterator, zend.ZendCeIterator, RIT_RecursiveTreeIterator)
}
func zim_spl_RecursiveTreeIterator_setPrefixPart(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var part zend.ZendLong
	var prefix *byte
	var prefix_len int
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "ls", &part, &prefix, &prefix_len) == zend.FAILURE {
		return
	}
	if 0 > part || part > 5 {
		zend.ZendThrowExceptionEx(spl_ce_OutOfRangeException, 0, "Use RecursiveTreeIterator::PREFIX_* constant")
		return
	}
	zend.SmartStrFree(&object.prefix[part])
	zend.SmartStrAppendl(&object.prefix[part], prefix, prefix_len)
}
func zim_spl_RecursiveTreeIterator_getPrefix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPrefix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_setPostfix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var postfix *byte
	var postfix_len int
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s", &postfix, &postfix_len) == zend.FAILURE {
		return
	}
	zend.SmartStrFree(&object.postfix[0])
	zend.SmartStrAppendl(&object.postfix[0], postfix, postfix_len)
}
func zim_spl_RecursiveTreeIterator_getEntry(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetEntry(object, return_value)
}
func zim_spl_RecursiveTreeIterator_getPostfix(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var prefix zend.Zval
	var entry zend.Zval
	var postfix zend.Zval
	var ptr *byte
	var str *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if object.GetIterators() == nil {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	if (object.GetFlags() & RTIT_BYPASS_CURRENT) != 0 {
		var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
		var data *zend.Zval
		SPL_FETCH_SUB_ITERATOR(iterator, object)
		data = iterator.funcs.get_current_data(iterator)
		if data != nil {
			zend.ZVAL_COPY_DEREF(return_value, data)
			return
		} else {
			zend.RETVAL_NULL()
			return
		}
	}
	zend.ZVAL_NULL(&prefix)
	zend.ZVAL_NULL(&entry)
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetEntry(object, &entry)
	if zend.Z_TYPE(entry) != zend.IS_STRING {
		zend.ZvalPtrDtor(&prefix)
		zend.ZvalPtrDtor(&entry)
		zend.RETVAL_NULL()
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = zend.ZendStringAlloc(zend.Z_STRLEN(prefix)+zend.Z_STRLEN(entry)+zend.Z_STRLEN(postfix), 0)
	ptr = zend.ZSTR_VAL(str)
	memcpy(ptr, zend.Z_STRVAL(prefix), zend.Z_STRLEN(prefix))
	ptr += zend.Z_STRLEN(prefix)
	memcpy(ptr, zend.Z_STRVAL(entry), zend.Z_STRLEN(entry))
	ptr += zend.Z_STRLEN(entry)
	memcpy(ptr, zend.Z_STRVAL(postfix), zend.Z_STRLEN(postfix))
	ptr += zend.Z_STRLEN(postfix)
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&entry)
	zend.ZvalPtrDtor(&postfix)
	zend.RETVAL_NEW_STR(str)
	return
}
func zim_spl_RecursiveTreeIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(zend.ZEND_THIS)
	var iterator *zend.ZendObjectIterator
	var prefix zend.Zval
	var key zend.Zval
	var postfix zend.Zval
	var key_copy zend.Zval
	var ptr *byte
	var str *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	if iterator.funcs.get_current_key != nil {
		iterator.funcs.get_current_key(iterator, &key)
	} else {
		zend.ZVAL_NULL(&key)
	}
	if (object.GetFlags() & RTIT_BYPASS_KEY) != 0 {
		zend.RETVAL_ZVAL(&key, 1, 1)
		return
	}
	if zend.Z_TYPE(key) != zend.IS_STRING {
		if zend.ZendMakePrintableZval(&key, &key_copy) != 0 {
			key = key_copy
		}
	}
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)
	str = zend.ZendStringAlloc(zend.Z_STRLEN(prefix)+zend.Z_STRLEN(key)+zend.Z_STRLEN(postfix), 0)
	ptr = zend.ZSTR_VAL(str)
	memcpy(ptr, zend.Z_STRVAL(prefix), zend.Z_STRLEN(prefix))
	ptr += zend.Z_STRLEN(prefix)
	memcpy(ptr, zend.Z_STRVAL(key), zend.Z_STRLEN(key))
	ptr += zend.Z_STRLEN(key)
	memcpy(ptr, zend.Z_STRVAL(postfix), zend.Z_STRLEN(postfix))
	ptr += zend.Z_STRLEN(postfix)
	*ptr = 0
	zend.ZvalPtrDtor(&prefix)
	zend.ZvalPtrDtor(&key)
	zend.ZvalPtrDtor(&postfix)
	zend.RETVAL_NEW_STR(str)
	return
}
func SplDualItGetMethod(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var function_handler *zend.ZendFunction
	var intern *SplDualItObject
	intern = SplDualItFromObj(*object)
	function_handler = zend.ZendStdGetMethod(object, method, key)
	if function_handler == nil && intern.GetCe() != nil {
		if b.Assign(&function_handler, zend.ZendHashFindPtr(&intern.inner.ce.function_table, method)) == nil {
			if zend.Z_OBJ_HT(intern.GetZobject()).get_method != nil {
				*object = zend.Z_OBJ(intern.GetZobject())
				function_handler = (*object).handlers.get_method(object, method, key)
			}
		} else {
			*object = zend.Z_OBJ(intern.GetZobject())
		}
	}
	return function_handler
}
func APPENDIT_CHECK_CTOR(intern *SplDualItObject) {
	if intern.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "Classes derived from %s must call %s::__construct()", zend.ZSTR_VAL(spl_ce_AppendIterator.name), zend.ZSTR_VAL(spl_ce_AppendIterator.name))
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
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplDualItConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ce_base *zend.ZendClassEntry, ce_inner *zend.ZendClassEntry, dit_type DualItType) *SplDualItObject {
	var zobject *zend.Zval
	var retval zend.Zval
	var intern *SplDualItObject
	var ce *zend.ZendClassEntry = nil
	var inc_refcount int = 1
	var error_handling zend.ZendErrorHandling
	intern = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if intern.GetDitType() != DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s::getIterator() must be called exactly once per instance", zend.ZSTR_VAL(ce_base.name))
		return nil
	}
	intern.SetDitType(dit_type)
	switch dit_type {
	case DIT_LimitIterator:
		intern.SetOffset(0)
		intern.SetCount(-1)
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "O|ll", &zobject, ce_inner, &intern.u.limit.offset, &intern.u.limit.count) == zend.FAILURE {
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
		break
	case DIT_CachingIterator:

	case DIT_RecursiveCachingIterator:
		var flags zend.ZendLong = CIT_CALL_TOSTRING
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "O|l", &zobject, ce_inner, &flags) == zend.FAILURE {
			return nil
		}
		if SplCitCheckFlags(flags) != zend.SUCCESS {
			zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
			return nil
		}
		intern.SetUCachingFlags(intern.GetUCachingFlags() | flags&CIT_PUBLIC)
		zend.ArrayInit(&intern.u.caching.zcache)
		break
	case DIT_IteratorIterator:
		var ce_cast *zend.ZendClassEntry
		var class_name *zend.ZendString
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "O|S", &zobject, ce_inner, &class_name) == zend.FAILURE {
			return nil
		}
		ce = zend.Z_OBJCE_P(zobject)
		if zend.InstanceofFunction(ce, zend.ZendCeIterator) == 0 {
			if zend.ZEND_NUM_ARGS() > 1 {
				if !(b.Assign(&ce_cast, zend.ZendLookupClass(class_name))) || zend.InstanceofFunction(ce, ce_cast) == 0 || ce_cast.get_iterator == nil {
					zend.ZendThrowException(spl_ce_LogicException, "Class to downcast to not found or not base class or does not implement Traversable", 0)
					return nil
				}
				ce = ce_cast
			}
			if zend.InstanceofFunction(ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(zobject, ce, &ce.iterator_funcs_ptr.zf_new_iterator, "getiterator", &retval)
				if zend.ExecutorGlobals.exception != nil {
					zend.ZvalPtrDtor(&retval)
					return nil
				}
				if zend.Z_TYPE(retval) != zend.IS_OBJECT || zend.InstanceofFunction(zend.Z_OBJCE(retval), zend.ZendCeTraversable) == 0 {
					zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "%s::getIterator() must return an object that implements Traversable", zend.ZSTR_VAL(ce.name))
					return nil
				}
				zobject = &retval
				ce = zend.Z_OBJCE_P(zobject)
				inc_refcount = 0
			}
		}
		break
	case DIT_AppendIterator:
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		SplInstantiate(spl_ce_ArrayIterator, &intern.u.append.zarrayit)
		zend.ZendCallMethodWith0Params(&intern.u.append.zarrayit, spl_ce_ArrayIterator, &spl_ce_ArrayIterator.constructor, "__construct", nil)
		intern.SetUAppendIterator(spl_ce_ArrayIterator.get_iterator(spl_ce_ArrayIterator, &intern.u.append.zarrayit, 0))
		zend.ZendRestoreErrorHandling(&error_handling)
		return intern
	case DIT_RegexIterator:

	case DIT_RecursiveRegexIterator:
		var regex *zend.ZendString
		var mode zend.ZendLong = REGIT_MODE_MATCH
		intern.SetUseFlags(zend.ZEND_NUM_ARGS() >= 5)
		intern.SetURegexFlags(0)
		intern.SetPregFlags(0)
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "OS|lll", &zobject, ce_inner, &regex, &mode, &intern.u.regex.flags, &intern.u.regex.preg_flags) == zend.FAILURE {
			return nil
		}
		if mode < 0 || mode >= REGIT_MODE_MAX {
			zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+zend.ZEND_LONG_FMT, mode)
			return nil
		}
		intern.SetMode(mode)
		intern.SetURegexRegex(zend.ZendStringCopy(regex))
		zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_InvalidArgumentException, &error_handling)
		intern.SetPce(pcre_get_compiled_regex_cache(regex))
		zend.ZendRestoreErrorHandling(&error_handling)
		if intern.GetPce() == nil {

			/* pcre_get_compiled_regex_cache has already sent error */

			return nil

			/* pcre_get_compiled_regex_cache has already sent error */

		}
		php_pcre_pce_incref(intern.GetPce())
		break
	case DIT_CallbackFilterIterator:

	case DIT_RecursiveCallbackFilterIterator:
		var cfi *_spl_cbfilter_it_intern = zend.Emalloc(b.SizeOf("* cfi"))
		cfi.fci.object = nil
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "Of", &zobject, ce_inner, &cfi.fci, &cfi.fcc) == zend.FAILURE {
			zend.Efree(cfi)
			return nil
		}
		zend.Z_TRY_ADDREF(cfi.fci.function_name)
		cfi.SetObject(cfi.fcc.object)
		if cfi.GetObject() != nil {
			zend.GC_ADDREF(cfi.GetObject())
		}
		intern.SetCbfilter(cfi)
		break
	default:
		if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "O", &zobject, ce_inner) == zend.FAILURE {
			return nil
		}
		break
	}
	if inc_refcount != 0 {
		zend.Z_ADDREF_P(zobject)
	}
	zend.ZVAL_OBJ(&intern.inner.zobject, zend.Z_OBJ_P(zobject))
	if dit_type == DIT_IteratorIterator {
		intern.SetCe(ce)
	} else {
		intern.SetCe(zend.Z_OBJCE_P(zobject))
	}
	intern.SetObject(zend.Z_OBJ_P(zobject))
	intern.SetInnerIterator(intern.GetCe().get_iterator(intern.GetCe(), zobject, 0))
	return intern
}
func zim_spl_FilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_FilterIterator, zend.ZendCeIterator, DIT_FilterIterator)
}
func zim_spl_CallbackFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_CallbackFilterIterator, zend.ZendCeIterator, DIT_CallbackFilterIterator)
}
func zim_spl_dual_it_getInnerIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if !(zend.Z_ISUNDEF(intern.GetZobject())) {
		var value *zend.Zval = &intern.inner.zobject
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func SplDualItFree(intern *SplDualItObject) {
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().funcs.invalidate_current != nil {
		intern.GetInnerIterator().funcs.invalidate_current(intern.GetInnerIterator())
	}
	if zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&intern.current.data)
		zend.ZVAL_UNDEF(&intern.current.data)
	}
	if zend.Z_TYPE(intern.GetKey()) != zend.IS_UNDEF {
		zend.ZvalPtrDtor(&intern.current.key)
		zend.ZVAL_UNDEF(&intern.current.key)
	}
	if intern.GetDitType() == DIT_CachingIterator || intern.GetDitType() == DIT_RecursiveCachingIterator {
		if zend.Z_TYPE(intern.GetZstr()) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&intern.u.caching.zstr)
			zend.ZVAL_UNDEF(&intern.u.caching.zstr)
		}
		if zend.Z_TYPE(intern.GetZchildren()) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&intern.u.caching.zchildren)
			zend.ZVAL_UNDEF(&intern.u.caching.zchildren)
		}
	}
}
func SplDualItRewind(intern *SplDualItObject) {
	SplDualItFree(intern)
	intern.SetPos(0)
	if intern.GetInnerIterator() != nil && intern.GetInnerIterator().funcs.rewind != nil {
		intern.GetInnerIterator().funcs.rewind(intern.GetInnerIterator())
	}
}
func SplDualItValid(intern *SplDualItObject) int {
	if intern.GetInnerIterator() == nil {
		return zend.FAILURE
	}

	/* FAILURE / SUCCESS */

	return intern.GetInnerIterator().funcs.valid(intern.GetInnerIterator())

	/* FAILURE / SUCCESS */
}
func SplDualItFetch(intern *SplDualItObject, check_more int) int {
	var data *zend.Zval
	SplDualItFree(intern)
	if check_more == 0 || SplDualItValid(intern) == zend.SUCCESS {
		data = intern.GetInnerIterator().funcs.get_current_data(intern.GetInnerIterator())
		if data != nil {
			zend.ZVAL_COPY(&intern.current.data, data)
		}
		if intern.GetInnerIterator().funcs.get_current_key != nil {
			intern.GetInnerIterator().funcs.get_current_key(intern.GetInnerIterator(), &intern.current.key)
			if zend.ExecutorGlobals.exception != nil {
				zend.ZvalPtrDtor(&intern.current.key)
				zend.ZVAL_UNDEF(&intern.current.key)
			}
		} else {
			zend.ZVAL_LONG(&intern.current.key, intern.GetPos())
		}
		if zend.ExecutorGlobals.exception != nil {
			return zend.FAILURE
		} else {
			return zend.SUCCESS
		}
	}
	return zend.FAILURE
}
func SplDualItNext(intern *SplDualItObject, do_free int) {
	if do_free != 0 {
		SplDualItFree(intern)
	} else if intern.GetInnerIterator() == nil {
		zend.ZendThrowError(nil, "The inner constructor wasn't initialized with an iterator instance")
		return
	}
	intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
	intern.GetPos()++
}
func ZimSplDualItRewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplDualItFetch(intern, 1)
}
func ZimSplDualItValid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF)
	return
}
func ZimSplDualItKey(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.Z_TYPE(intern.GetKey()) != zend.IS_UNDEF {
		var value *zend.Zval = &intern.current.key
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func ZimSplDualItCurrent(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF {
		var value *zend.Zval = &intern.current.data
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func ZimSplDualItNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	SplDualItFetch(intern, 1)
}
func SplFilterItFetch(zthis *zend.Zval, intern *SplDualItObject) {
	var retval zend.Zval
	for SplDualItFetch(intern, 1) == zend.SUCCESS {
		zend.ZendCallMethodWith0Params(zthis, intern.std.ce, nil, "accept", &retval)
		if zend.Z_TYPE(retval) != zend.IS_UNDEF {
			if zend.ZendIsTrue(&retval) != 0 {
				zend.ZvalPtrDtor(&retval)
				return
			}
			zend.ZvalPtrDtor(&retval)
		}
		if zend.ExecutorGlobals.exception != nil {
			return
		}
		intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
	}
	SplDualItFree(intern)
}
func SplFilterItRewind(zthis *zend.Zval, intern *SplDualItObject) {
	SplDualItRewind(intern)
	SplFilterItFetch(zthis, intern)
}
func SplFilterItNext(zthis *zend.Zval, intern *SplDualItObject) {
	SplDualItNext(intern, 1)
	SplFilterItFetch(zthis, intern)
}
func zim_spl_FilterIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItRewind(zend.ZEND_THIS, intern)
}
func zim_spl_FilterIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplFilterItNext(zend.ZEND_THIS, intern)
}
func zim_spl_RecursiveCallbackFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveCallbackFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveCallbackFilterIterator)
}
func zim_spl_RecursiveFilterIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveFilterIterator, spl_ce_RecursiveIterator, DIT_RecursiveFilterIterator)
}
func zim_spl_RecursiveFilterIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "haschildren", &retval)
	if zend.Z_TYPE(retval) != zend.IS_UNDEF {
		zend.RETVAL_ZVAL(&retval, 0, 1)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func zim_spl_RecursiveFilterIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", &retval)
	if zend.ExecutorGlobals.exception == nil && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		SplInstantiateArgEx1(zend.Z_OBJCE_P(zend.ZEND_THIS), return_value, &retval)
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveCallbackFilterIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", &retval)
	if zend.ExecutorGlobals.exception == nil && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		SplInstantiateArgEx2(zend.Z_OBJCE_P(zend.ZEND_THIS), return_value, &retval, &intern.u.cbfilter.fci.function_name)
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_ParentIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_ParentIterator, spl_ce_RecursiveIterator, DIT_ParentIterator)
}
func zim_spl_RegexIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RegexIterator, zend.ZendCeIterator, DIT_RegexIterator)
}
func zim_spl_CallbackFilterIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	var fci *zend.ZendFcallInfo = &intern.u.cbfilter.GetFci()
	var fcc *zend.ZendFcallInfoCache = &intern.u.cbfilter.GetFcc()
	var params []zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if zend.Z_TYPE(intern.GetData()) == zend.IS_UNDEF || zend.Z_TYPE(intern.GetKey()) == zend.IS_UNDEF {
		zend.RETVAL_FALSE
		return
	}
	zend.ZVAL_COPY_VALUE(&params[0], &intern.current.data)
	zend.ZVAL_COPY_VALUE(&params[1], &intern.current.key)
	zend.ZVAL_COPY_VALUE(&params[2], &intern.inner.zobject)
	fci.retval = return_value
	fci.param_count = 3
	fci.params = params
	fci.no_separation = 0
	if zend.ZendCallFunction(fci, fcc) != zend.SUCCESS || zend.Z_ISUNDEF_P(return_value) {
		zend.RETVAL_FALSE
		return
	}
	if zend.ExecutorGlobals.exception != nil {
		zend.RETVAL_NULL()
		return
	}

	/* zend_call_function may change args to IS_REF */

	zend.ZVAL_COPY_VALUE(&intern.current.data, &params[0])
	zend.ZVAL_COPY_VALUE(&intern.current.key, &params[1])
}
func zim_spl_RegexIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var result *zend.ZendString
	var subject *zend.ZendString
	var count int = 0
	var zcount zend.Zval
	var rv zend.Zval
	var match_data *pcre2_match_data
	var re *pcre2_code
	var rc int
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.Z_TYPE(intern.GetData()) == zend.IS_UNDEF {
		zend.RETVAL_FALSE
		return
	}
	if (intern.GetURegexFlags() & REGIT_USE_KEY) != 0 {
		subject = zend.ZvalGetString(&intern.current.key)
	} else {
		if zend.Z_TYPE(intern.GetData()) == zend.IS_ARRAY {
			zend.RETVAL_FALSE
			return
		}
		subject = zend.ZvalGetString(&intern.current.data)
	}

	/* Exception during string conversion. */

	if zend.ExecutorGlobals.exception != nil {
		return
	}
	switch intern.GetMode() {
	case REGIT_MODE_MAX:

	case REGIT_MODE_MATCH:
		re = php_pcre_pce_re(intern.GetPce())
		match_data = php_pcre_create_match_data(0, re)
		if match_data == nil {
			zend.RETVAL_FALSE
			return
		}
		rc = pcre2_match(re, PCRE2_SPTR(zend.ZSTR_VAL(subject)), zend.ZSTR_LEN(subject), 0, 0, match_data, php_pcre_mctx())
		zend.RETVAL_BOOL(rc >= 0)
		php_pcre_free_match_data(match_data)
		break
	case REGIT_MODE_ALL_MATCHES:

	case REGIT_MODE_GET_MATCH:
		zend.ZvalPtrDtor(&intern.current.data)
		zend.ZVAL_UNDEF(&intern.current.data)
		php_pcre_match_impl(intern.GetPce(), subject, &zcount, &intern.current.data, intern.GetMode() == REGIT_MODE_ALL_MATCHES, intern.GetUseFlags(), intern.GetPregFlags(), 0)
		zend.RETVAL_BOOL(zend.Z_LVAL(zcount) > 0)
		break
	case REGIT_MODE_SPLIT:
		zend.ZvalPtrDtor(&intern.current.data)
		zend.ZVAL_UNDEF(&intern.current.data)
		php_pcre_split_impl(intern.GetPce(), subject, &intern.current.data, -1, intern.GetPregFlags())
		count = zend.ZendHashNumElements(zend.Z_ARRVAL(intern.GetData()))
		zend.RETVAL_BOOL(count > 1)
		break
	case REGIT_MODE_REPLACE:
		var replacement *zend.Zval = zend.ZendReadProperty(intern.std.ce, zend.ZEND_THIS, "replacement", b.SizeOf("\"replacement\"")-1, 1, &rv)
		var replacement_str *zend.ZendString = zend.ZvalTryGetString(replacement)
		if zend.UNEXPECTED(replacement_str == nil) {
			return
		}
		result = php_pcre_replace_impl(intern.GetPce(), subject, zend.ZSTR_VAL(subject), zend.ZSTR_LEN(subject), replacement_str, -1, &count)
		if (intern.GetURegexFlags() & REGIT_USE_KEY) != 0 {
			zend.ZvalPtrDtor(&intern.current.key)
			zend.ZVAL_STR(&intern.current.key, result)
		} else {
			zend.ZvalPtrDtor(&intern.current.data)
			zend.ZVAL_STR(&intern.current.data, result)
		}
		zend.ZendStringRelease(replacement_str)
		zend.RETVAL_BOOL(count > 0)
	}
	if (intern.GetURegexFlags() & REGIT_INVERTED) != 0 {
		zend.RETVAL_BOOL(zend.Z_TYPE_P(return_value) != zend.IS_TRUE)
	}
	zend.ZendStringReleaseEx(subject, 0)
}
func zim_spl_RegexIterator_getRegex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_STR_COPY(intern.GetURegexRegex())
	return
}
func zim_spl_RegexIterator_getMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_LONG(intern.GetMode())
	return
}
func zim_spl_RegexIterator_setMode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var mode zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &mode) == zend.FAILURE {
		return
	}
	if mode < 0 || mode >= REGIT_MODE_MAX {
		zend.ZendThrowExceptionEx(spl_ce_InvalidArgumentException, 0, "Illegal mode "+zend.ZEND_LONG_FMT, mode)
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetMode(mode)
}
func zim_spl_RegexIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_LONG(intern.GetURegexFlags())
	return
}
func zim_spl_RegexIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &flags) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetURegexFlags(flags)
}
func zim_spl_RegexIterator_getPregFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetUseFlags() != 0 {
		zend.RETVAL_LONG(intern.GetPregFlags())
		return
	} else {
		zend.RETVAL_LONG(0)
		return
	}
}
func zim_spl_RegexIterator_setPregFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var preg_flags zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &preg_flags) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.SetPregFlags(preg_flags)
	intern.SetUseFlags(1)
}
func zim_spl_RecursiveRegexIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveRegexIterator, spl_ce_RecursiveIterator, DIT_RecursiveRegexIterator)
}
func zim_spl_RecursiveRegexIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var retval zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", &retval)
	if zend.ExecutorGlobals.exception == nil {
		var args []zend.Zval
		zend.ZVAL_COPY(&args[0], &retval)
		zend.ZVAL_STR_COPY(&args[1], intern.GetURegexRegex())
		zend.ZVAL_LONG(&args[2], intern.GetMode())
		zend.ZVAL_LONG(&args[3], intern.GetURegexFlags())
		zend.ZVAL_LONG(&args[4], intern.GetPregFlags())
		SplInstantiateArgN(zend.Z_OBJCE_P(zend.ZEND_THIS), return_value, 5, args)
		zend.ZvalPtrDtor(&args[0])
		zend.ZvalPtrDtor(&args[1])
	}
	zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveRegexIterator_accept(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.Z_TYPE(intern.GetData()) == zend.IS_UNDEF {
		zend.RETVAL_FALSE
		return
	} else if zend.Z_TYPE(intern.GetData()) == zend.IS_ARRAY {
		zend.RETVAL_BOOL(zend.ZendHashNumElements(zend.Z_ARRVAL(intern.GetData())) > 0)
		return
	}
	zend.ZendCallMethodWith0Params(zend.ZEND_THIS, spl_ce_RegexIterator, nil, "accept", return_value)
}
func SplDualItDtor(_object *zend.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	SplDualItFree(object)
	if object.GetInnerIterator() != nil {
		zend.ZendIteratorDtor(object.GetInnerIterator())
	}
}
func SplDualItFreeStorage(_object *zend.ZendObject) {
	var object *SplDualItObject = SplDualItFromObj(_object)
	if !(zend.Z_ISUNDEF(object.GetZobject())) {
		zend.ZvalPtrDtor(&object.inner.zobject)
	}
	if object.GetDitType() == DIT_AppendIterator {
		zend.ZendIteratorDtor(object.GetUAppendIterator())
		if zend.Z_TYPE(object.GetZarrayit()) != zend.IS_UNDEF {
			zend.ZvalPtrDtor(&object.u.append.zarrayit)
		}
	}
	if object.GetDitType() == DIT_CachingIterator || object.GetDitType() == DIT_RecursiveCachingIterator {
		zend.ZvalPtrDtor(&object.u.caching.zcache)
	}
	if object.GetDitType() == DIT_RegexIterator || object.GetDitType() == DIT_RecursiveRegexIterator {
		if object.GetPce() != nil {
			php_pcre_pce_decref(object.GetPce())
		}
		if object.GetURegexRegex() != nil {
			zend.ZendStringReleaseEx(object.GetURegexRegex(), 0)
		}
	}
	if object.GetDitType() == DIT_CallbackFilterIterator || object.GetDitType() == DIT_RecursiveCallbackFilterIterator {
		if object.GetCbfilter() != nil {
			var cbfilter *_spl_cbfilter_it_intern = object.GetCbfilter()
			object.SetCbfilter(nil)
			zend.ZvalPtrDtor(&cbfilter.fci.function_name)
			if cbfilter.fci.object != nil {
				zend.OBJ_RELEASE(cbfilter.fci.object)
			}
			zend.Efree(cbfilter)
		}
	}
	zend.ZendObjectStdDtor(&object.std)
}
func SplDualItNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var intern *SplDualItObject
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_dual_it_object"), class_type)
	intern.SetDitType(DIT_Unknown)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.std.handlers = &SplHandlersDualIt
	return &intern.std
}
func SplLimitItValid(intern *SplDualItObject) int {
	/* FAILURE / SUCCESS */

	if intern.GetCount() != -1 && intern.GetPos() >= intern.GetOffset()+intern.GetCount() {
		return zend.FAILURE
	} else {
		return SplDualItValid(intern)
	}

	/* FAILURE / SUCCESS */
}
func SplLimitItSeek(intern *SplDualItObject, pos zend.ZendLong) {
	var zpos zend.Zval
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
		zend.ZVAL_LONG(&zpos, pos)
		SplDualItFree(intern)
		zend.ZendCallMethodWith1Params(&intern.inner.zobject, intern.GetCe(), nil, "seek", nil, &zpos)
		if zend.ExecutorGlobals.exception == nil {
			intern.SetPos(pos)
			if SplLimitItValid(intern) == zend.SUCCESS {
				SplDualItFetch(intern, 0)
			}
		}
	} else {

		/* emulate the forward seek, by next() calls */

		if pos < intern.GetPos() {
			SplDualItRewind(intern)
		}
		for pos > intern.GetPos() && SplDualItValid(intern) == zend.SUCCESS {
			SplDualItNext(intern, 1)
		}
		if SplDualItValid(intern) == zend.SUCCESS {
			SplDualItFetch(intern, 1)
		}
	}
}
func zim_spl_LimitIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_LimitIterator, zend.ZendCeIterator, DIT_LimitIterator)
}
func zim_spl_LimitIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplLimitItSeek(intern, intern.GetOffset())
}
func zim_spl_LimitIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it

	/*    RETURN_BOOL(spl_limit_it_valid(intern) == SUCCESS);*/

	zend.RETVAL_BOOL((intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount()) && zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF)
	return
}
func zim_spl_LimitIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
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
func zim_spl_LimitIterator_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var pos zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &pos) == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplLimitItSeek(intern, pos)
	zend.RETVAL_LONG(intern.GetPos())
	return
}
func zim_spl_LimitIterator_getPosition(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_LONG(intern.GetPos())
	return
}
func SplCachingItValid(intern *SplDualItObject) int {
	if (intern.GetUCachingFlags() & CIT_VALID) != 0 {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplCachingItHasNext(intern *SplDualItObject) int { return SplDualItValid(intern) }
func SplCachingItNext(intern *SplDualItObject) {
	if SplDualItFetch(intern, 1) == zend.SUCCESS {
		intern.SetUCachingFlags(intern.GetUCachingFlags() | CIT_VALID)

		/* Full cache ? */

		if (intern.GetUCachingFlags() & CIT_FULL_CACHE) != 0 {
			var key *zend.Zval = &intern.current.key
			var data *zend.Zval = &intern.current.data
			zend.ZVAL_DEREF(data)
			zend.Z_TRY_ADDREF_P(data)
			zend.ArraySetZvalKey(zend.Z_ARRVAL(intern.GetZcache()), key, data)
			zend.ZvalPtrDtor(data)
		}

		/* Recursion ? */

		if intern.GetDitType() == DIT_RecursiveCachingIterator {
			var retval zend.Zval
			var zchildren zend.Zval
			var zflags zend.Zval
			zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "haschildren", &retval)
			if zend.ExecutorGlobals.exception != nil {
				zend.ZvalPtrDtor(&retval)
				if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
					zend.ZendClearException()
				} else {
					return
				}
			} else {
				if zend.ZendIsTrue(&retval) != 0 {
					zend.ZendCallMethodWith0Params(&intern.inner.zobject, intern.GetCe(), nil, "getchildren", &zchildren)
					if zend.ExecutorGlobals.exception != nil {
						zend.ZvalPtrDtor(&zchildren)
						if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
							zend.ZendClearException()
						} else {
							zend.ZvalPtrDtor(&retval)
							return
						}
					} else {
						zend.ZVAL_LONG(&zflags, intern.GetUCachingFlags()&CIT_PUBLIC)
						SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &intern.u.caching.zchildren, &zchildren, &zflags)
						zend.ZvalPtrDtor(&zchildren)
					}
				}
				zend.ZvalPtrDtor(&retval)
				if zend.ExecutorGlobals.exception != nil {
					if (intern.GetUCachingFlags() & CIT_CATCH_GET_CHILD) != 0 {
						zend.ZendClearException()
					} else {
						return
					}
				}
			}
		}
		if (intern.GetUCachingFlags() & (CIT_TOSTRING_USE_INNER | CIT_CALL_TOSTRING)) != 0 {
			var use_copy int
			var expr_copy zend.Zval
			if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_INNER) != 0 {
				zend.ZVAL_COPY_VALUE(&intern.u.caching.zstr, &intern.inner.zobject)
			} else {
				zend.ZVAL_COPY_VALUE(&intern.u.caching.zstr, &intern.current.data)
			}
			use_copy = zend.ZendMakePrintableZval(&intern.u.caching.zstr, &expr_copy)
			if use_copy != 0 {
				zend.ZVAL_COPY_VALUE(&intern.u.caching.zstr, &expr_copy)
			} else {
				zend.Z_TRY_ADDREF(intern.GetZstr())
			}
		}
		SplDualItNext(intern, 0)
	} else {
		intern.SetUCachingFlags(intern.GetUCachingFlags() &^ CIT_VALID)
	}
}
func SplCachingItRewind(intern *SplDualItObject) {
	SplDualItRewind(intern)
	zend.ZendHashClean(zend.Z_ARRVAL(intern.GetZcache()))
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_CachingIterator, zend.ZendCeIterator, DIT_CachingIterator)
}
func zim_spl_CachingIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItRewind(intern)
}
func zim_spl_CachingIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(SplCachingItValid(intern) == zend.SUCCESS)
	return
}
func zim_spl_CachingIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator_hasNext(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(SplCachingItHasNext(intern) == zend.SUCCESS)
	return
}
func zim_spl_CachingIterator___toString(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & (CIT_CALL_TOSTRING | CIT_TOSTRING_USE_KEY | CIT_TOSTRING_USE_CURRENT | CIT_TOSTRING_USE_INNER)) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not fetch string value (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_KEY) != 0 {
		zend.ZVAL_COPY(return_value, &intern.current.key)
		zend.ConvertToString(return_value)
		return
	} else if (intern.GetUCachingFlags() & CIT_TOSTRING_USE_CURRENT) != 0 {
		zend.ZVAL_COPY(return_value, &intern.current.data)
		zend.ConvertToString(return_value)
		return
	}
	if zend.Z_TYPE(intern.GetZstr()) == zend.IS_STRING {
		zend.RETVAL_STR_COPY(zend.Z_STR_P(&intern.u.caching.zstr))
		return
	} else {
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func zim_spl_CachingIterator_offsetSet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var value *zend.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "Sz", &key, &value) == zend.FAILURE {
		return
	}
	zend.Z_TRY_ADDREF_P(value)
	zend.ZendSymtableUpdate(zend.Z_ARRVAL(intern.GetZcache()), key, value)
}
func zim_spl_CachingIterator_offsetGet(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var value *zend.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &key) == zend.FAILURE {
		return
	}
	if b.Assign(&value, zend.ZendSymtableFind(zend.Z_ARRVAL(intern.GetZcache()), key)) == nil {
		zend.ZendError(zend.E_NOTICE, "Undefined index: %s", zend.ZSTR_VAL(key))
		return
	}
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_CachingIterator_offsetUnset(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &key) == zend.FAILURE {
		return
	}
	zend.ZendSymtableDel(zend.Z_ARRVAL(intern.GetZcache()), key)
}
func zim_spl_CachingIterator_offsetExists(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var key *zend.ZendString
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &key) == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(zend.ZendSymtableExists(zend.Z_ARRVAL(intern.GetZcache()), key) != 0)
	return
}
func zim_spl_CachingIterator_getCache(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	zend.ZVAL_COPY(return_value, &intern.u.caching.zcache)
}
func zim_spl_CachingIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_LONG(intern.GetUCachingFlags())
	return
}
func zim_spl_CachingIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &flags) == zend.FAILURE {
		return
	}
	if SplCitCheckFlags(flags) != zend.SUCCESS {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
		return
	}
	if (intern.GetUCachingFlags()&CIT_CALL_TOSTRING) != 0 && (flags&CIT_CALL_TOSTRING) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag CALL_TO_STRING is not possible", 0)
		return
	}
	if (intern.GetUCachingFlags()&CIT_TOSTRING_USE_INNER) != 0 && (flags&CIT_TOSTRING_USE_INNER) == 0 {
		zend.ZendThrowException(spl_ce_InvalidArgumentException, "Unsetting flag TOSTRING_USE_INNER is not possible", 0)
		return
	}
	if (flags&CIT_FULL_CACHE) != 0 && (intern.GetUCachingFlags()&CIT_FULL_CACHE) == 0 {

		/* clear on (re)enable */

		zend.ZendHashClean(zend.Z_ARRVAL(intern.GetZcache()))

		/* clear on (re)enable */

	}
	intern.SetUCachingFlags(intern.GetUCachingFlags() & ^CIT_PUBLIC | flags&CIT_PUBLIC)
}
func zim_spl_CachingIterator_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if (intern.GetUCachingFlags() & CIT_FULL_CACHE) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_BadMethodCallException, 0, "%s does not use a full cache (see CachingIterator::__construct)", zend.ZSTR_VAL(zend.Z_OBJCE_P(zend.ZEND_THIS).name))
		return
	}
	zend.RETVAL_LONG(zend.ZendHashNumElements(zend.Z_ARRVAL(intern.GetZcache())))
	return
}
func zim_spl_RecursiveCachingIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_RecursiveCachingIterator, spl_ce_RecursiveIterator, DIT_RecursiveCachingIterator)
}
func zim_spl_RecursiveCachingIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(zend.Z_TYPE(intern.GetZchildren()) != zend.IS_UNDEF)
	return
}
func zim_spl_RecursiveCachingIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.Z_TYPE(intern.GetZchildren()) != zend.IS_UNDEF {
		var value *zend.Zval = &intern.u.caching.zchildren
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func zim_spl_IteratorIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_IteratorIterator, zend.ZendCeTraversable, DIT_IteratorIterator)
}
func zim_spl_NoRewindIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_NoRewindIterator, zend.ZendCeIterator, DIT_NoRewindIterator)
}
func zim_spl_NoRewindIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_NoRewindIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(intern.GetInnerIterator().funcs.valid(intern.GetInnerIterator()) == zend.SUCCESS)
	return
}
func zim_spl_NoRewindIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if intern.GetInnerIterator().funcs.get_current_key != nil {
		intern.GetInnerIterator().funcs.get_current_key(intern.GetInnerIterator(), return_value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func zim_spl_NoRewindIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var data *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	data = intern.GetInnerIterator().funcs.get_current_data(intern.GetInnerIterator())
	if data != nil {
		zend.ZVAL_COPY_DEREF(return_value, data)
	}
}
func zim_spl_NoRewindIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetInnerIterator().funcs.move_forward(intern.GetInnerIterator())
}
func zim_spl_InfiniteIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_InfiniteIterator, zend.ZendCeIterator, DIT_InfiniteIterator)
}
func zim_spl_InfiniteIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItNext(intern, 1)
	if SplDualItValid(intern) == zend.SUCCESS {
		SplDualItFetch(intern, 0)
	} else {
		SplDualItRewind(intern)
		if SplDualItValid(intern) == zend.SUCCESS {
			SplDualItFetch(intern, 0)
		}
	}
}
func zim_spl_EmptyIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_EmptyIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}
func zim_spl_EmptyIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the key of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendThrowException(spl_ce_BadMethodCallException, "Accessing the value of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func SplAppendItNextIterator(intern *SplDualItObject) int {
	SplDualItFree(intern)
	if !(zend.Z_ISUNDEF(intern.GetZobject())) {
		zend.ZvalPtrDtor(&intern.inner.zobject)
		zend.ZVAL_UNDEF(&intern.inner.zobject)
		intern.SetCe(nil)
		if intern.GetInnerIterator() != nil {
			zend.ZendIteratorDtor(intern.GetInnerIterator())
			intern.SetInnerIterator(nil)
		}
	}
	if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) == zend.SUCCESS {
		var it *zend.Zval
		it = intern.GetUAppendIterator().funcs.get_current_data(intern.GetUAppendIterator())
		zend.ZVAL_COPY(&intern.inner.zobject, it)
		intern.SetCe(zend.Z_OBJCE_P(it))
		intern.SetInnerIterator(intern.GetCe().get_iterator(intern.GetCe(), it, 0))
		SplDualItRewind(intern)
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplAppendItFetch(intern *SplDualItObject) {
	for SplDualItValid(intern) != zend.SUCCESS {
		intern.GetUAppendIterator().funcs.move_forward(intern.GetUAppendIterator())
		if SplAppendItNextIterator(intern) != zend.SUCCESS {
			return
		}
	}
	SplDualItFetch(intern, 0)
}
func SplAppendItNext(intern *SplDualItObject) {
	if SplDualItValid(intern) == zend.SUCCESS {
		SplDualItNext(intern, 1)
	}
	SplAppendItFetch(intern)
}
func zim_spl_AppendIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplDualItConstruct(execute_data, return_value, spl_ce_AppendIterator, zend.ZendCeIterator, DIT_AppendIterator)
}
func zim_spl_AppendIterator_append(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var it *zend.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	if zend.ZendParseParametersEx(zend.ZEND_PARSE_PARAMS_QUIET, zend.ZEND_NUM_ARGS(), "O", &it, zend.ZendCeIterator) == zend.FAILURE {
		return
	}
	if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) == zend.SUCCESS && SplDualItValid(intern) != zend.SUCCESS {
		SplArrayIteratorAppend(&intern.u.append.zarrayit, it)
		intern.GetUAppendIterator().funcs.move_forward(intern.GetUAppendIterator())
	} else {
		SplArrayIteratorAppend(&intern.u.append.zarrayit, it)
	}
	if intern.GetInnerIterator() == nil || SplDualItValid(intern) != zend.SUCCESS {
		if intern.GetUAppendIterator().funcs.valid(intern.GetUAppendIterator()) != zend.SUCCESS {
			intern.GetUAppendIterator().funcs.rewind(intern.GetUAppendIterator())
		}
		for {
			SplAppendItNextIterator(intern)
			if zend.Z_OBJ(intern.GetZobject()) == zend.Z_OBJ_P(it) {
				break
			}
		}
		SplAppendItFetch(intern)
	}
}
func zim_spl_AppendIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplDualItFetch(intern, 1)
	if zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF {
		var value *zend.Zval = &intern.current.data
		zend.ZVAL_COPY_DEREF(return_value, value)
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func zim_spl_AppendIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	intern.GetUAppendIterator().funcs.rewind(intern.GetUAppendIterator())
	if SplAppendItNextIterator(intern) == zend.SUCCESS {
		SplAppendItFetch(intern)
	}
}
func zim_spl_AppendIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	zend.RETVAL_BOOL(zend.Z_TYPE(intern.GetData()) != zend.IS_UNDEF)
	return
}
func zim_spl_AppendIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	SplAppendItNext(intern)
}
func zim_spl_AppendIterator_getIteratorIndex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	APPENDIT_CHECK_CTOR(intern)
	SplArrayIteratorKey(&intern.u.append.zarrayit, return_value)
}
func zim_spl_AppendIterator_getArrayIterator(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplDualItObject
	var value *zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(zend.ZEND_THIS)
	if it.GetDitType() == DIT_Unknown {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The object is in an invalid state as the parent constructor was not called")
		return
	}
	intern = it
	value = &intern.u.append.zarrayit
	zend.ZVAL_COPY_DEREF(return_value, value)
}
func SplIteratorApply(obj *zend.Zval, apply_func SplIteratorApplyFuncT, puser any) int {
	var iter *zend.ZendObjectIterator
	var ce *zend.ZendClassEntry = zend.Z_OBJCE_P(obj)
	iter = ce.get_iterator(ce, obj, 0)
	if zend.ExecutorGlobals.exception != nil {
		goto done
	}
	iter.index = 0
	if iter.funcs.rewind != nil {
		iter.funcs.rewind(iter)
		if zend.ExecutorGlobals.exception != nil {
			goto done
		}
	}
	for iter.funcs.valid(iter) == zend.SUCCESS {
		if zend.ExecutorGlobals.exception != nil {
			goto done
		}
		if apply_func(iter, puser) == zend.ZEND_HASH_APPLY_STOP || zend.ExecutorGlobals.exception != nil {
			goto done
		}
		iter.index++
		iter.funcs.move_forward(iter)
		if zend.ExecutorGlobals.exception != nil {
			goto done
		}
	}
done:
	if iter != nil {
		zend.ZendIteratorDtor(iter)
	}
	if zend.ExecutorGlobals.exception != nil {
		return zend.FAILURE
	} else {
		return zend.SUCCESS
	}
}
func SplIteratorToArrayApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *zend.Zval
	var return_value *zend.Zval = (*zend.Zval)(puser)
	data = iter.funcs.get_current_data(iter)
	if zend.ExecutorGlobals.exception != nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if data == nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if iter.funcs.get_current_key != nil {
		var key zend.Zval
		iter.funcs.get_current_key(iter, &key)
		if zend.ExecutorGlobals.exception != nil {
			return zend.ZEND_HASH_APPLY_STOP
		}
		zend.ArraySetZvalKey(zend.Z_ARRVAL_P(return_value), &key, data)
		zend.ZvalPtrDtor(&key)
	} else {
		zend.Z_TRY_ADDREF_P(data)
		zend.AddNextIndexZval(return_value, data)
	}
	return zend.ZEND_HASH_APPLY_KEEP
}
func SplIteratorToValuesApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *zend.Zval
	var return_value *zend.Zval = (*zend.Zval)(puser)
	data = iter.funcs.get_current_data(iter)
	if zend.ExecutorGlobals.exception != nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	if data == nil {
		return zend.ZEND_HASH_APPLY_STOP
	}
	zend.Z_TRY_ADDREF_P(data)
	zend.AddNextIndexZval(return_value, data)
	return zend.ZEND_HASH_APPLY_KEEP
}
func ZifIteratorToArray(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var use_keys zend.ZendBool = 1
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O|b", &obj, zend.ZendCeTraversable, &use_keys) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.ArrayInit(return_value)
	SplIteratorApply(obj, b.Cond(use_keys != 0, SplIteratorToArrayApply, SplIteratorToValuesApply), any(return_value))
}
func SplIteratorCountApply(iter *zend.ZendObjectIterator, puser any) int {
	*((*zend.ZendLong)(puser))++
	return zend.ZEND_HASH_APPLY_KEEP
}
func ZifIteratorCount(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var obj *zend.Zval
	var count zend.ZendLong = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "O", &obj, zend.ZendCeTraversable) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	if SplIteratorApply(obj, SplIteratorCountApply, any(&count)) == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(count)
	return
}
func SplIteratorFuncApply(iter *zend.ZendObjectIterator, puser any) int {
	var retval zend.Zval
	var apply_info *SplIteratorApplyInfo = (*SplIteratorApplyInfo)(puser)
	var result int
	apply_info.GetCount()++
	zend.ZendFcallInfoCall(&apply_info.fci, &apply_info.fcc, &retval, nil)
	if zend.ZendIsTrue(&retval) != 0 {
		result = zend.ZEND_HASH_APPLY_KEEP
	} else {
		result = zend.ZEND_HASH_APPLY_STOP
	}
	zend.ZvalPtrDtor(&retval)
	return result
}
func ZifIteratorApply(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var apply_info SplIteratorApplyInfo
	apply_info.SetArgs(nil)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "Of|a!", &apply_info.obj, zend.ZendCeTraversable, &apply_info.fci, &apply_info.fcc, &apply_info.args) == zend.FAILURE {
		return
	}
	apply_info.SetCount(0)
	zend.ZendFcallInfoArgs(&apply_info.fci, apply_info.GetArgs())
	if SplIteratorApply(apply_info.GetObj(), SplIteratorFuncApply, any(&apply_info)) == zend.FAILURE {
		zend.ZendFcallInfoArgs(&apply_info.fci, nil)
		return
	}
	zend.ZendFcallInfoArgs(&apply_info.fci, nil)
	zend.RETVAL_LONG(apply_info.GetCount())
	return
}
func ZmStartupSplIterators(type_ int, module_number int) int {
	SplRegisterInterface(&spl_ce_RecursiveIterator, "RecursiveIterator", spl_funcs_RecursiveIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIterator, 1, zend.ZendCeIterator)
	SplRegisterStdClass(&spl_ce_RecursiveIteratorIterator, "RecursiveIteratorIterator", spl_RecursiveIteratorIterator_new, spl_funcs_RecursiveIteratorIterator)
	zend.ZendClassImplements(spl_ce_RecursiveIteratorIterator, 1, zend.ZendCeIterator)
	memcpy(&SplHandlersRecItIt, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplHandlersRecItIt.offset = zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd())) - (*byte)(nil))
	SplHandlersRecItIt.get_method = SplRecursiveItGetMethod
	SplHandlersRecItIt.clone_obj = nil
	SplHandlersRecItIt.dtor_obj = spl_RecursiveIteratorIterator_dtor
	SplHandlersRecItIt.free_obj = spl_RecursiveIteratorIterator_free_storage
	memcpy(&SplHandlersDualIt, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplHandlersDualIt.offset = zend_long((*byte)(&((*SplDualItObject)(nil).GetStd())) - (*byte)(nil))
	SplHandlersDualIt.get_method = SplDualItGetMethod

	/*spl_handlers_dual_it.call_method = spl_dual_it_call_method;*/

	SplHandlersDualIt.clone_obj = nil
	SplHandlersDualIt.dtor_obj = SplDualItDtor
	SplHandlersDualIt.free_obj = SplDualItFreeStorage
	spl_ce_RecursiveIteratorIterator.get_iterator = SplRecursiveItGetIterator
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
	spl_ce_FilterIterator.ce_flags |= zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS
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
	return zend.SUCCESS
}
