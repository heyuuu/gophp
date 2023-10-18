package spl

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func SplDualItFromObj(obj *types.Object) *SplDualItObject {
	return (*SplDualItObject)((*byte)(obj - zend_long((*byte)(&((*SplDualItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLDUAL_IT_P(zv *types.Zval) *SplDualItObject { return SplDualItFromObj(zv.Object()) }
func SplRecursiveItFromObj(obj *types.Object) *SplRecursiveItObject {
	return (*SplRecursiveItObject)((*byte)(obj - zend_long((*byte)(&((*SplRecursiveItObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLRECURSIVE_IT_P(zv *types.Zval) *SplRecursiveItObject {
	return SplRecursiveItFromObj(zv.Object())
}
func SPL_FETCH_SUB_ITERATOR(var_ *zend.ZendObjectIterator, object *SplRecursiveItObject) {
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
			//zend.ZendIteratorDtor(sub_iter)
			// zend.ZvalPtrDtor(object.GetIterators()[object.GetLevel()].GetZobject())
		}
		object.GetLevel()--
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.SetLevel(0)
	// zend.ZvalPtrDtor(iter.GetIntern().GetData())
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
	var ce *types.ClassEntry
	var retval types.Zval
	var child types.Zval
	var sub_iter *zend.ZendObjectIterator
	var has_children int
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	for zend.EG__().NoException() {
	next_step:
		iterator = object.GetIterators()[object.GetLevel()].GetIterator()
		switch object.GetIterators()[object.GetLevel()].GetState() {
		case RS_NEXT:
			iterator.GetFuncs().GetMoveForward()(iterator)
			if zend.EG__().HasException() {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.EG__().ClearException()
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
			if zend.EG__().HasException() {
				if !object.IsRitCatchGetChild() {
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					return
				} else {
					zend.EG__().ClearException()
				}
			}
			if retval.IsNotUndef() {
				has_children = operators.IZendIsTrue(&retval)
				// zend.ZvalPtrDtor(&retval)
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
			if zend.EG__().HasException() {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.EG__().ClearException()
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
			if zend.EG__().HasException() {
				if !object.IsRitCatchGetChild() {
					return
				} else {
					zend.EG__().ClearException()
					// zend.ZvalPtrDtor(&child)
					object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
					goto next_step
				}
			}
			if child.IsUndef() || !child.IsObject() || !(lang.Assign(&ce, types.Z_OBJCE(child)) && operators.InstanceofFunction(ce, spl_ce_RecursiveIterator) != 0) {
				// zend.ZvalPtrDtor(&child)
				faults.ThrowException(spl_ce_UnexpectedValueException, "Objects returned by RecursiveIterator::getChildren() must implement RecursiveIterator", 0)
				return
			}
			if object.GetMode() == RIT_CHILD_FIRST {
				object.GetIterators()[object.GetLevel()].SetState(RS_SELF)
			} else {
				object.GetIterators()[object.GetLevel()].SetState(RS_NEXT)
			}
			object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")*(lang.PreInc(&(object.GetLevel()))+1)))
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
				if zend.EG__().HasException() {
					if !object.IsRitCatchGetChild() {
						return
					} else {
						zend.EG__().ClearException()
					}
				}
			}
			goto next_step
		}

		/* no more elements */

		if object.GetLevel() > 0 {
			if object.GetEndChildren() != nil {
				zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetEndChildren(), "endchildren", nil)
				if zend.EG__().HasException() {
					if !object.IsRitCatchGetChild() {
						return
					} else {
						zend.EG__().ClearException()
					}
				}
			}
			if object.GetLevel() > 0 {
				var garbage types.Zval
				types.ZVAL_COPY_VALUE(&garbage, object.GetIterators()[object.GetLevel()].GetZobject())
				object.GetIterators()[object.GetLevel()].GetZobject().SetUndef()
				// zend.ZvalPtrDtor(&garbage)
				//zend.ZendIteratorDtor(iterator)
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
		//zend.ZendIteratorDtor(sub_iter)
		// zend.ZvalPtrDtor(object.GetIterators()[b.PostDec(&(object.GetLevel()))].GetZobject())
		if zend.EG__().NoException() && (object.GetEndChildren() == nil || object.GetEndChildren().GetScope() != spl_ce_RecursiveIteratorIterator) {
			zend.ZendCallMethodWith0Params(zthis, object.GetCe(), object.GetEndChildren(), "endchildren", nil)
		}
	}
	object.SetIterators(zend.Erealloc(object.GetIterators(), b.SizeOf("spl_sub_iterator")))
	object.GetIterators()[0].SetState(RS_START)
	sub_iter = object.GetIterators()[0].GetIterator()
	if sub_iter.GetFuncs().GetRewind() != nil {
		sub_iter.GetFuncs().GetRewind()(sub_iter)
	}
	if zend.EG__().NoException() && object.GetBeginIteration() != nil && object.GetInIteration() == 0 {
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
func SplRecursiveItGetIterator(ce *types.ClassEntry, zobject *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplRecursiveItIterator
	var object *SplRecursiveItObject
	if by_ref != 0 {
		faults.ThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	iterator = zend.Emalloc(b.SizeOf("spl_recursive_it_iterator"))
	object = Z_SPLRECURSIVE_IT_P(zobject)
	if object.GetIterators() == nil {
		faults.Error(faults.E_ERROR, "The object to be iterated is in an invalid state: the parent constructor has not been called")
	}
	zend.ZendIteratorInit((*zend.ZendObjectIterator)(iterator))
	// 	zobject.AddRefcount()
	iterator.GetIntern().GetData().SetObject(zobject.Object())
	iterator.GetIntern().SetFuncs(&SplRecursiveItIteratorFuncs)
	return (*zend.ZendObjectIterator)(iterator)
}
func SplRecursiveItItConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval, ce_base *types.ClassEntry, ce_inner *types.ClassEntry, rit_type RecursiveItItType) {
	var object *types.Zval = executeData.ThisObjectZval()
	var intern *SplRecursiveItObject
	var iterator *types.Zval
	var ce_iterator *types.ClassEntry
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
		if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "o|lzl", &iterator, &flags, &user_caching_it_flags, &mode) == types.SUCCESS {
			if operators.InstanceofFunction(types.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, types.Z_OBJCE_P(iterator), types.Z_OBJCE_P(iterator).GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				// 				iterator.AddRefcount()
			}
			if user_caching_it_flags != nil {
				types.ZVAL_COPY(&caching_it_flags, user_caching_it_flags)
			} else {
				caching_it_flags.SetLong(CIT_CATCH_GET_CHILD)
			}
			SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, &caching_it, iterator, &caching_it_flags)
			// zend.ZvalPtrDtor(&caching_it_flags)
			// zend.ZvalPtrDtor(iterator)
			iterator = &caching_it
		} else {
			iterator = nil
		}
	case RIT_RecursiveIteratorIterator:
		fallthrough
	default:
		mode = RIT_LEAVES_ONLY
		flags = 0
		if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "o|ll", &iterator, &mode, &flags) == types.SUCCESS {
			if operators.InstanceofFunction(types.Z_OBJCE_P(iterator), zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(iterator, types.Z_OBJCE_P(iterator), types.Z_OBJCE_P(iterator).GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &aggregate_retval)
				iterator = &aggregate_retval
			} else {
				// 				iterator.AddRefcount()
			}
		} else {
			iterator = nil
		}
	}
	if iterator == nil || operators.InstanceofFunction(types.Z_OBJCE_P(iterator), spl_ce_RecursiveIterator) == 0 {
		if iterator != nil {
			// zend.ZvalPtrDtor(iterator)
		}
		faults.ThrowException(spl_ce_InvalidArgumentException, "An instance of RecursiveIterator or IteratorAggregate creating it is required", 0)
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
	intern.SetBeginIteration(intern.GetCe().FunctionTable().Get("beginiteration"))
	if intern.GetBeginIteration().GetScope() == ce_base {
		intern.SetBeginIteration(nil)
	}
	intern.SetEndIteration(intern.GetCe().FunctionTable().Get("enditeration"))
	if intern.GetEndIteration().GetScope() == ce_base {
		intern.SetEndIteration(nil)
	}
	intern.SetCallHasChildren(intern.GetCe().FunctionTable().Get("callhaschildren"))
	if intern.GetCallHasChildren().GetScope() == ce_base {
		intern.SetCallHasChildren(nil)
	}
	intern.SetCallGetChildren(intern.GetCe().FunctionTable().Get("callgetchildren"))
	if intern.GetCallGetChildren().GetScope() == ce_base {
		intern.SetCallGetChildren(nil)
	}
	intern.SetBeginChildren(intern.GetCe().FunctionTable().Get("beginchildren"))
	if intern.GetBeginChildren().GetScope() == ce_base {
		intern.SetBeginChildren(nil)
	}
	intern.SetEndChildren(intern.GetCe().FunctionTable().Get("endchildren"))
	if intern.GetEndChildren().GetScope() == ce_base {
		intern.SetEndChildren(nil)
	}
	intern.SetNextElement(intern.GetCe().FunctionTable().Get("nextelement"))
	if intern.GetNextElement().GetScope() == ce_base {
		intern.SetNextElement(nil)
	}
	ce_iterator = types.Z_OBJCE_P(iterator)
	intern.GetIterators()[0].SetIterator(ce_iterator.GetGetIterator()(ce_iterator, iterator, 0))
	intern.GetIterators()[0].GetZobject().SetObject(iterator.Object())
	intern.GetIterators()[0].SetCe(ce_iterator)
	intern.GetIterators()[0].SetState(RS_START)
	zend.ZendRestoreErrorHandling(&error_handling)
	if zend.EG__().HasException() {
		var sub_iter *zend.ZendObjectIterator
		for intern.GetLevel() >= 0 {
			sub_iter = intern.GetIterators()[intern.GetLevel()].GetIterator()
			//zend.ZendIteratorDtor(sub_iter)
			// zend.ZvalPtrDtor(intern.GetIterators()[b.PostDec(&(intern.GetLevel()))].GetZobject())
		}
		zend.Efree(intern.GetIterators())
		intern.SetIterators(nil)
	}
}
func zim_spl_RecursiveIteratorIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplRecursiveItItConstruct(executeData, return_value, spl_ce_RecursiveIteratorIterator, zend.ZendCeIterator, RIT_RecursiveIteratorIterator)
}
func zim_spl_RecursiveIteratorIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplRecursiveItRewindEx(object, executeData.ThisObjectZval())
}
func zim_spl_RecursiveIteratorIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetBool(SplRecursiveItValidEx(object, executeData.ThisObjectZval()) == types.SUCCESS)
	return
}
func zim_spl_RecursiveIteratorIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var iterator *zend.ZendObjectIterator
	if !executeData.CheckNumArgsNone(false) {
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
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var iterator *zend.ZendObjectIterator
	var data *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SPL_FETCH_SUB_ITERATOR(iterator, object)
	data = iterator.GetFuncs().GetGetCurrentData()(iterator)
	if data != nil {
		types.ZVAL_COPY_DEREF(return_value, data)
	}
}
func zim_spl_RecursiveIteratorIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	SplRecursiveItMoveForwardEx(object, executeData.ThisObjectZval())
}
func zim_spl_RecursiveIteratorIterator_getDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetLong(object.GetLevel())
	return
}
func zim_spl_RecursiveIteratorIterator_getSubIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
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
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	value = object.GetIterators()[level].GetZobject()
	types.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_RecursiveIteratorIterator_getInnerIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var zobject *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	types.ZVAL_COPY_DEREF(return_value, zobject)
}
func zim_spl_RecursiveIteratorIterator_beginIteration(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endIteration(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_callHasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var ce *types.ClassEntry
	var zobject *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		return_value.SetNull()
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	if zobject.IsUndef() {
		return_value.SetFalse()
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "haschildren", return_value)
		if return_value.IsUndef() {
			return_value.SetFalse()
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_callGetChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var ce *types.ClassEntry
	var zobject *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	ce = object.GetIterators()[object.GetLevel()].GetCe()
	zobject = object.GetIterators()[object.GetLevel()].GetZobject()
	if zobject.IsUndef() {
		return
	} else {
		zend.ZendCallMethodWith0Params(zobject, ce, nil, "getchildren", return_value)
		if return_value.IsUndef() {
			return_value.SetNull()
			return
		}
	}
}
func zim_spl_RecursiveIteratorIterator_beginChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_endChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_nextElement(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_RecursiveIteratorIterator_setMaxDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var max_depth zend.ZendLong = -1
	if zend.ZendParseParameters(executeData.NumArgs(), "|l", &max_depth) == types.FAILURE {
		return
	}
	if max_depth < -1 {
		faults.ThrowException(spl_ce_OutOfRangeException, "Parameter max_depth must be >= -1", 0)
		return
	} else if max_depth > core.INT_MAX {
		max_depth = core.INT_MAX
	}
	object.SetMaxDepth(int(max_depth))
}
func zim_spl_RecursiveIteratorIterator_getMaxDepth(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
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
func SplRecursiveItGetMethod(zobject **types.Object, method *types.String, key *types.Zval) types.IFunction {
	var function_handler types.IFunction
	var object *SplRecursiveItObject = SplRecursiveItFromObj(*zobject)
	var level zend.ZendLong = object.GetLevel()
	var zobj *types.Zval
	if object.GetIterators() == nil {
		core.PhpErrorDocref("", faults.E_ERROR, fmt.Sprintf("The %s instance wasn't initialized properly", zobject.GetCe().Name()))
	}
	zobj = object.GetIterators()[level].GetZobject()
	function_handler = zend.ZendStdGetMethod(zobject, method, key)
	if function_handler == nil {
		if lang.Assign(&function_handler, types.Z_OBJCE_P(zobj).FunctionTable().Get(method.GetStr())) == nil {
			*zobject = zobj.Object()
			function_handler = (*zobject).GetMethod(method.GetStr(), key)
		} else {
			*zobject = zobj.Object()
		}
	}
	return function_handler
}
func spl_RecursiveIteratorIterator_dtor(_object *types.Object) {
	var object *SplRecursiveItObject = SplRecursiveItFromObj(_object)
	var sub_iter *zend.ZendObjectIterator

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	if object.GetIterators() != nil {
		for object.GetLevel() >= 0 {
			sub_iter = object.GetIterators()[object.GetLevel()].GetIterator()
			//zend.ZendIteratorDtor(sub_iter)
			// zend.ZvalPtrDtor(object.GetIterators()[b.PostDec(&(object.GetLevel()))].GetZobject())
		}
		zend.Efree(object.GetIterators())
		object.SetIterators(nil)
	}
}
func spl_RecursiveIteratorIterator_free_storage(_object *types.Object) {
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
func spl_RecursiveIteratorIterator_new_ex(class_type *types.ClassEntry, init_prefix int) *types.Object {
	var intern = NewSplRecursiveItObject(class_type)
	if init_prefix != 0 {
		intern.GetPrefix()[0].WriteString("")
		intern.GetPrefix()[1].WriteString("| ")
		intern.GetPrefix()[2].WriteString("  ")
		intern.GetPrefix()[3].WriteString("|-")
		intern.GetPrefix()[4].WriteString("\\-")
		intern.GetPrefix()[5].WriteString("")
		intern.GetPostfix()[0].WriteString("")
	}
	return intern.GetStd()
}
func spl_RecursiveIteratorIterator_new(class_type *types.ClassEntry) *types.Object {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 0)
}
func spl_RecursiveTreeIterator_new(class_type *types.ClassEntry) *types.Object {
	return spl_RecursiveIteratorIterator_new_ex(class_type, 1)
}
func SplRecursiveTreeIteratorGetPrefix(object *SplRecursiveItObject, return_value *types.Zval) {
	var str zend.SmartStr
	var has_next types.Zval
	var level int
	str.WriteString(object.GetPrefix()[0].GetStr())
	for level = 0; level < object.GetLevel(); level++ {
		zend.ZendCallMethodWith0Params(object.GetIterators()[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
		if has_next.IsNotUndef() {
			if has_next.IsTrue() {
				str.WriteString(object.GetPrefix()[1].GetStr())
			} else {
				str.WriteString(object.GetPrefix()[2].GetStr())
			}
		}
	}
	zend.ZendCallMethodWith0Params(object.GetIterators()[level].GetZobject(), object.GetIterators()[level].GetCe(), nil, "hasnext", &has_next)
	if has_next.IsNotUndef() {
		if has_next.IsTrue() {
			str.WriteString(object.GetPrefix()[3].GetStr())
		} else {
			str.WriteString(object.GetPrefix()[4].GetStr())
		}
	}
	str.WriteString(object.GetPrefix()[5].GetStr())
	return_value.SetString(str.String())
	return
}
func SplRecursiveTreeIteratorGetEntry(object *SplRecursiveItObject, return_value *types.Zval) {
	var iterator *zend.ZendObjectIterator = object.GetIterators()[object.GetLevel()].GetIterator()
	var data *types.Zval
	data = iterator.GetFuncs().GetGetCurrentData()(iterator)
	if data != nil {
		data = types.ZVAL_DEREF(data)

		/* TODO: Remove this special case? */
		if data.IsType(types.IsArray) {
			return_value.SetString("Array")
		} else {
			types.ZVAL_COPY(return_value, data)
			operators.ConvertToString(return_value)
		}
	}
}
func SplRecursiveTreeIteratorGetPostfix(object *SplRecursiveItObject, return_value *types.Zval) {
	return_value.SetStringEx(object.GetPostfix()[0].GetS())
	// 	return_value.AddRefcount()
}
func zim_spl_RecursiveTreeIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplRecursiveItItConstruct(executeData, return_value, spl_ce_RecursiveTreeIterator, zend.ZendCeIterator, RIT_RecursiveTreeIterator)
}
func zim_spl_RecursiveTreeIterator_setPrefixPart(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var part zend.ZendLong
	var prefix *byte
	var prefix_len int
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if zend.ZendParseParameters(executeData.NumArgs(), "ls", &part, &prefix, &prefix_len) == types.FAILURE {
		return
	}
	if 0 > part || part > 5 {
		faults.ThrowException(spl_ce_OutOfRangeException, "Use RecursiveTreeIterator::PREFIX_* constant", 0)
		return
	}
	object.GetPrefix()[part].Free()
	object.GetPrefix()[part].WriteString(b.CastStr(prefix, prefix_len))
}
func zim_spl_RecursiveTreeIterator_getPrefix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	SplRecursiveTreeIteratorGetPrefix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_setPostfix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var postfix *byte
	var postfix_len int
	if zend.ZendParseParameters(executeData.NumArgs(), "s", &postfix, &postfix_len) == types.FAILURE {
		return
	}
	object.GetPostfix()[0].Free()
	object.GetPostfix()[0].WriteString(b.CastStr(postfix, postfix_len))
}
func zim_spl_RecursiveTreeIterator_getEntry(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	SplRecursiveTreeIteratorGetEntry(object, return_value)
}
func zim_spl_RecursiveTreeIterator_getPostfix(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, return_value)
}
func zim_spl_RecursiveTreeIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var prefix types.Zval
	var entry types.Zval
	var postfix types.Zval
	var ptr *byte
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if object.GetIterators() == nil {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !entry.IsString() {
		// zend.ZvalPtrDtor(&prefix)
		// zend.ZvalPtrDtor(&entry)
		return_value.SetNull()
		return
	}
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)

	str := prefix.String() + entry.String() + postfix.String()
	return_value.SetString(str)
	return
}
func zim_spl_RecursiveTreeIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var object *SplRecursiveItObject = Z_SPLRECURSIVE_IT_P(executeData.ThisObjectZval())
	var iterator *zend.ZendObjectIterator
	var prefix types.Zval
	var key types.Zval
	var postfix types.Zval
	var key_copy types.Zval
	var ptr *byte
	if !executeData.CheckNumArgsNone(false) {
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
	if !key.IsString() {
		if zend.ZendMakePrintableZval(&key, &key_copy) != 0 {
			key = key_copy
		}
	}
	SplRecursiveTreeIteratorGetPrefix(object, &prefix)
	SplRecursiveTreeIteratorGetPostfix(object, &postfix)

	str := prefix.String() + key.String() + postfix.String()
	return_value.SetString(str)
	return
}
func SplDualItGetMethod(object **types.Object, method *types.String, key *types.Zval) types.IFunction {
	var function_handler types.IFunction
	var intern *SplDualItObject
	intern = SplDualItFromObj(*object)
	function_handler = zend.ZendStdGetMethod(object, method, key)
	if function_handler == nil && intern.GetCe() != nil {
		if lang.Assign(&function_handler, intern.GetCe().FunctionTable().Get(method.GetStr())) == nil {
			if intern.GetZobject().Object().CanGetMethod() {
				*object = intern.GetZobject().Object()
				function_handler = (*object).GetMethod(method.GetStr(), key)
			}
		} else {
			*object = intern.GetZobject().Object()
		}
	}
	return function_handler
}
func APPENDIT_CHECK_CTOR(intern *SplDualItObject) {
	if intern.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("Classes derived from %s must call %s::__construct()", spl_ce_AppendIterator.Name(), spl_ce_AppendIterator.Name()), 0)
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
func SplDualItConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval, ce_base *types.ClassEntry, ce_inner *types.ClassEntry, dit_type DualItType) *SplDualItObject {
	var zobject *types.Zval
	var retval types.Zval
	var intern *SplDualItObject
	var ce *types.ClassEntry = nil
	var inc_refcount int = 1
	var error_handling zend.ZendErrorHandling
	intern = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if intern.GetDitType() != DIT_Unknown {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s::getIterator() must be called exactly once per instance", ce_base.Name()), 0)
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
			faults.ThrowException(spl_ce_OutOfRangeException, "Parameter offset must be >= 0", 0)
			return nil
		}
		if intern.GetCount() < 0 && intern.GetCount() != -1 {
			faults.ThrowException(spl_ce_OutOfRangeException, "Parameter count must either be -1 or a value greater than or equal 0", 0)
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
			faults.ThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
			return nil
		}
		intern.AddUCachingFlags(flags & CIT_PUBLIC)
		zend.ArrayInit(intern.GetZcache())
	case DIT_IteratorIterator:
		var ce_cast *types.ClassEntry
		var class_name *types.String
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O|S", &zobject, ce_inner, &class_name) == types.FAILURE {
			return nil
		}
		ce = types.Z_OBJCE_P(zobject)
		if operators.InstanceofFunction(ce, zend.ZendCeIterator) == 0 {
			if executeData.NumArgs() > 1 {
				if !(lang.Assign(&ce_cast, zend.ZendLookupClassEx(class_name.GetStr()))) || operators.InstanceofFunction(ce, ce_cast) == 0 || ce_cast.GetGetIterator() == nil {
					faults.ThrowException(spl_ce_LogicException, "Class to downcast to not found or not base class or does not implement Traversable", 0)
					return nil
				}
				ce = ce_cast
			}
			if operators.InstanceofFunction(ce, zend.ZendCeAggregate) != 0 {
				zend.ZendCallMethodWith0Params(zobject, ce, ce.GetIteratorFuncsPtr().GetZfNewIterator(), "getiterator", &retval)
				if zend.EG__().HasException() {
					// zend.ZvalPtrDtor(&retval)
					return nil
				}
				if !retval.IsObject() || operators.InstanceofFunction(types.Z_OBJCE(retval), zend.ZendCeTraversable) == 0 {
					faults.ThrowException(spl_ce_LogicException, fmt.Sprintf("%s::getIterator() must return an object that implements Traversable", ce.Name()), 0)
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
		var regex *types.String
		var mode zend.ZendLong = REGIT_MODE_MATCH
		intern.SetUseFlags(executeData.NumArgs() >= 5)
		intern.SetURegexFlags(0)
		intern.SetPregFlags(0)
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "OS|lll", &zobject, ce_inner, &regex, &mode, intern.GetURegexFlags(), intern.GetPregFlags()) == types.FAILURE {
			return nil
		}
		if mode < 0 || mode >= REGIT_MODE_MAX {
			faults.ThrowException(spl_ce_InvalidArgumentException, fmt.Sprintf("Illegal mode %d", mode), 0)
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
		//cfi.GetFci().GetFunctionName().TryAddRefcount()
		cfi.SetObject(cfi.GetFcc().GetObject())
		if cfi.GetObject() != nil {
			//cfi.GetObject().AddRefcount()
		}
		intern.SetCbfilter(cfi)
	default:
		if zend.ZendParseParametersThrow(executeData.NumArgs(), "O", &zobject, ce_inner) == types.FAILURE {
			return nil
		}
	}
	if inc_refcount != 0 {
		// 		zobject.AddRefcount()
	}
	intern.GetZobject().SetObject(zobject.Object())
	if dit_type == DIT_IteratorIterator {
		intern.SetCe(ce)
	} else {
		intern.SetCe(types.Z_OBJCE_P(zobject))
	}
	intern.SetObject(zobject.Object())
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if intern.GetData().IsNotUndef() {
		// zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
	}
	if intern.GetKey().IsNotUndef() {
		// zend.ZvalPtrDtor(intern.GetKey())
		intern.GetKey().SetUndef()
	}
	if intern.GetDitType() == DIT_CachingIterator || intern.GetDitType() == DIT_RecursiveCachingIterator {
		if intern.GetZstr().IsNotUndef() {
			// zend.ZvalPtrDtor(intern.GetZstr())
			intern.GetZstr().SetUndef()
		}
		if intern.GetZchildren().IsNotUndef() {
			// zend.ZvalPtrDtor(intern.GetZchildren())
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
			if zend.EG__().HasException() {
				// zend.ZvalPtrDtor(intern.GetKey())
				intern.GetKey().SetUndef()
			}
		} else {
			intern.GetKey().SetLong(intern.GetPos())
		}
		if zend.EG__().HasException() {
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
		faults.ThrowError(nil, "The inner constructor wasn't initialized with an iterator instance")
		return
	}
	intern.GetInnerIterator().GetFuncs().GetMoveForward()(intern.GetInnerIterator())
	intern.GetPos()++
}
func ZimSplDualItRewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplDualItFetch(intern, 1)
}
func ZimSplDualItValid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(intern.GetData().IsNotUndef())
	return
}
func ZimSplDualItKey(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if intern.GetKey().IsNotUndef() {
		var value *types.Zval = intern.GetKey()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func ZimSplDualItCurrent(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if intern.GetData().IsNotUndef() {
		var value *types.Zval = intern.GetData()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func ZimSplDualItNext(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
		if retval.IsNotUndef() {
			if operators.ZvalIsTrue(&retval) {
				// zend.ZvalPtrDtor(&retval)
				return
			}
			// zend.ZvalPtrDtor(&retval)
		}
		if zend.EG__().HasException() {
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplFilterItRewind(executeData.ThisObjectZval(), intern)
}
func zim_spl_FilterIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplFilterItNext(executeData.ThisObjectZval(), intern)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "haschildren", &retval)
	if retval.IsNotUndef() {
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().NoException() && retval.IsNotUndef() {
		SplInstantiateArgEx1(types.Z_OBJCE_P(executeData.ThisObjectZval()), return_value, &retval)
	}
	// zend.ZvalPtrDtor(&retval)
}
func zim_spl_RecursiveCallbackFilterIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var retval types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().NoException() && retval.IsNotUndef() {
		SplInstantiateArgEx2(types.Z_OBJCE_P(executeData.ThisObjectZval()), return_value, &retval, intern.GetCbfilter().GetFci().GetFunctionName())
	}
	// zend.ZvalPtrDtor(&retval)
}
func zim_spl_ParentIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_ParentIterator, spl_ce_RecursiveIterator, DIT_ParentIterator)
}
func zim_spl_RegexIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RegexIterator, zend.ZendCeIterator, DIT_RegexIterator)
}
func zim_spl_CallbackFilterIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	var fci *types.ZendFcallInfo = intern.GetCbfilter().GetFci()
	var fcc *types.ZendFcallInfoCache = intern.GetCbfilter().GetFcc()
	var params []types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if intern.GetData().IsUndef() || intern.GetKey().IsUndef() {
		return_value.SetFalse()
		return
	}
	types.ZVAL_COPY_VALUE(&params[0], intern.GetData())
	types.ZVAL_COPY_VALUE(&params[1], intern.GetKey())
	types.ZVAL_COPY_VALUE(&params[2], intern.GetZobject())
	fci.SetRetval(return_value)
	fci.SetParamCount(3)
	fci.SetParams(params)
	fci.SetNoSeparation(false)
	if zend.ZendCallFunction(fci, fcc) != types.SUCCESS || return_value.IsUndef() {
		return_value.SetFalse()
		return
	}
	if zend.EG__().HasException() {
		return_value.SetNull()
		return
	}

	/* zend_call_function may change args to IS_REF */

	types.ZVAL_COPY_VALUE(intern.GetData(), &params[0])
	types.ZVAL_COPY_VALUE(intern.GetKey(), &params[1])
}
func zim_spl_RegexIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var result *types.String
	var subject *types.String
	var count int = 0
	var zcount types.Zval
	var rv types.Zval
	var match_data *pcre2_match_data
	var re *pcre2_code
	var rc int
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if intern.GetData().IsUndef() {
		return_value.SetFalse()
		return
	}
	if intern.IsUseKey() {
		subject = operators.ZvalGetString(intern.GetKey())
	} else {
		if intern.GetData().IsType(types.IsArray) {
			return_value.SetFalse()
			return
		}
		subject = operators.ZvalGetString(intern.GetData())
	}

	/* Exception during string conversion. */

	if zend.EG__().HasException() {
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
		return_value.SetBool(rc >= 0)
		php_pcre_free_match_data(match_data)
	case REGIT_MODE_ALL_MATCHES:
		fallthrough
	case REGIT_MODE_GET_MATCH:
		// zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
		php_pcre_match_impl(intern.GetPce(), subject, &zcount, intern.GetData(), intern.GetMode() == REGIT_MODE_ALL_MATCHES, intern.GetUseFlags(), intern.GetPregFlags(), 0)
		return_value.SetBool(zcount.Long() > 0)
	case REGIT_MODE_SPLIT:
		// zend.ZvalPtrDtor(intern.GetData())
		intern.GetData().SetUndef()
		php_pcre_split_impl(intern.GetPce(), subject, intern.GetData(), -1, intern.GetPregFlags())
		count = intern.GetData().Array().Len()
		return_value.SetBool(count > 1)
	case REGIT_MODE_REPLACE:
		var replacement *types.Zval = zend.ZendReadProperty(intern.GetStd().GetCe(), executeData.ThisObjectZval(), "replacement", 1, &rv)
		var replacement_str *types.String = operators.ZvalTryGetString(replacement)
		if replacement_str == nil {
			return
		}
		result = php_pcre_replace_impl(intern.GetPce(), subject, subject.GetVal(), subject.GetLen(), replacement_str, -1, &count)
		if intern.IsUseKey() {
			// zend.ZvalPtrDtor(intern.GetKey())
			intern.GetKey().SetStringEx(result)
		} else {
			// zend.ZvalPtrDtor(intern.GetData())
			intern.GetData().SetStringEx(result)
		}
		// types.ZendStringRelease(replacement_str)
		return_value.SetBool(count > 0)
	}
	if intern.IsInverted() {
		return_value.SetBool(!return_value.IsTrue())
	}
	// types.ZendStringReleaseEx(subject, 0)
}
func zim_spl_RegexIterator_getRegex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetString(intern.GetURegexRegex().GetStr())
	return
}
func zim_spl_RegexIterator_getMode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
		faults.ThrowException(spl_ce_InvalidArgumentException, fmt.Sprintf("Illegal mode %d", mode), 0)
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	intern.SetMode(mode)
}
func zim_spl_RegexIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	intern.SetURegexFlags(flags)
}
func zim_spl_RegexIterator_getPregFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &retval)
	if zend.EG__().NoException() {
		var args = []*types.Zval{
			&retval,
			types.NewZvalString(intern.GetURegexRegex().GetStr()),
			types.NewZvalLong(intern.GetMode()),
			types.NewZvalLong(intern.GetURegexFlags()),
			types.NewZvalLong(intern.GetPregFlags()),
		}
		SplInstantiateArgN(types.Z_OBJCE_P(executeData.ThisObjectZval()), return_value, args)
	}
}
func zim_spl_RecursiveRegexIterator_accept(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if intern.GetData().IsUndef() {
		return_value.SetFalse()
		return
	} else if intern.GetData().IsType(types.IsArray) {
		return_value.SetBool(intern.GetData().Array().Len() > 0)
		return
	}
	zend.ZendCallMethodWith0Params(executeData.ThisObjectZval(), spl_ce_RegexIterator, nil, "accept", return_value)
}
func SplDualItDtor(_object *types.Object) {
	var object *SplDualItObject = SplDualItFromObj(_object)

	/* call standard dtor */

	zend.ZendObjectsDestroyObject(_object)
	SplDualItFree(object)
	if object.GetInnerIterator() != nil {
		//zend.ZendIteratorDtor(object.GetInnerIterator())
	}
}
func SplDualItFreeStorage(_object *types.Object) {
	var object *SplDualItObject = SplDualItFromObj(_object)
	if !(object.GetZobject().IsUndef()) {
		// zend.ZvalPtrDtor(object.GetZobject())
	}
	if object.GetDitType() == DIT_AppendIterator {
		//zend.ZendIteratorDtor(object.GetUAppendIterator())
		if object.GetZarrayit().IsNotUndef() {
			// zend.ZvalPtrDtor(object.GetZarrayit())
		}
	}
	if object.GetDitType() == DIT_CachingIterator || object.GetDitType() == DIT_RecursiveCachingIterator {
		// zend.ZvalPtrDtor(object.GetZcache())
	}
	if object.GetDitType() == DIT_RegexIterator || object.GetDitType() == DIT_RecursiveRegexIterator {
		if object.GetPce() != nil {
			php_pcre_pce_decref(object.GetPce())
		}
		if object.GetURegexRegex() != nil {
			// types.ZendStringReleaseEx(object.GetURegexRegex(), 0)
		}
	}
	if object.GetDitType() == DIT_CallbackFilterIterator || object.GetDitType() == DIT_RecursiveCallbackFilterIterator {
		if object.GetCbfilter() != nil {
			var cbfilter *_spl_cbfilter_it_intern = object.GetCbfilter()
			object.SetCbfilter(nil)
			// zend.ZvalPtrDtor(cbfilter.GetFci().GetFunctionName())
			if cbfilter.GetFci().GetObject() != nil {
				// zend.OBJ_RELEASE(cbfilter.GetFci().GetObject())
			}
			zend.Efree(cbfilter)
		}
	}
	zend.ZendObjectStdDtor(object.GetStd())
}
func SplDualItNew(class_type *types.ClassEntry) *types.Object {
	var intern *SplDualItObject = NewSplDualItObject(class_type)
	return intern.GetStd()
}
func SplLimitItValid(intern *SplDualItObject) int {
	/* FAILURE / SUCCESS */

	if intern.GetCount() != -1 && intern.GetPos() >= intern.GetOffset()+intern.GetCount() {
		return types.FAILURE
	} else {
		return SplDualItValid(intern)
	}
}
func SplLimitItSeek(intern *SplDualItObject, pos zend.ZendLong) {
	var zpos types.Zval
	SplDualItFree(intern)
	if pos < intern.GetOffset() {
		faults.ThrowException(spl_ce_OutOfBoundsException, fmt.Sprintf("Cannot seek to %d which is below the offset %d", pos, intern.GetOffset()), 0)
		return
	}
	if pos >= intern.GetOffset()+intern.GetCount() && intern.GetCount() != -1 {
		faults.ThrowException(spl_ce_OutOfBoundsException, fmt.Sprintf("Cannot seek to %d which is behind offset %d plus count %d", pos, intern.GetOffset(), intern.GetCount()), 0)
		return
	}
	if pos != intern.GetPos() && operators.InstanceofFunction(intern.GetCe(), spl_ce_SeekableIterator) != 0 {
		zpos.SetLong(pos)
		SplDualItFree(intern)
		zend.ZendCallMethodWith1Params(intern.GetZobject(), intern.GetCe(), nil, "seek", nil, &zpos)
		if zend.EG__().NoException() {
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
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplDualItRewind(intern)
	SplLimitItSeek(intern, intern.GetOffset())
}
func zim_spl_LimitIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it

	/*    RETURN_BOOL(spl_limit_it_valid(intern) == SUCCESS);*/

	return_value.SetBool((intern.GetCount() == -1 || intern.GetPos() < intern.GetOffset()+intern.GetCount()) && intern.GetData().IsNotUndef())
	return
}
func zim_spl_LimitIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplLimitItSeek(intern, pos)
	return_value.SetLong(intern.GetPos())
	return
}
func zim_spl_LimitIterator_getPosition(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
			//data.TryAddRefcount()
			zend.ArraySetZvalKey(intern.GetZcache().Array(), key, data)
			// zend.ZvalPtrDtor(data)
		}

		/* Recursion ? */

		if intern.GetDitType() == DIT_RecursiveCachingIterator {
			var retval types.Zval
			var zchildren types.Zval
			var zflags types.Zval
			zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "haschildren", &retval)
			if zend.EG__().HasException() {
				// zend.ZvalPtrDtor(&retval)
				if intern.IsCatchGetChild() {
					zend.EG__().ClearException()
				} else {
					return
				}
			} else {
				if operators.ZvalIsTrue(&retval) {
					zend.ZendCallMethodWith0Params(intern.GetZobject(), intern.GetCe(), nil, "getchildren", &zchildren)
					if zend.EG__().HasException() {
						// zend.ZvalPtrDtor(&zchildren)
						if intern.IsCatchGetChild() {
							zend.EG__().ClearException()
						} else {
							// zend.ZvalPtrDtor(&retval)
							return
						}
					} else {
						zflags.SetLong(intern.GetUCachingFlags() & CIT_PUBLIC)
						SplInstantiateArgEx2(spl_ce_RecursiveCachingIterator, intern.GetZchildren(), &zchildren, &zflags)
						// zend.ZvalPtrDtor(&zchildren)
					}
				}
				// zend.ZvalPtrDtor(&retval)
				if zend.EG__().HasException() {
					if intern.IsCatchGetChild() {
						zend.EG__().ClearException()
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
				//intern.GetZstr().TryAddRefcount()
			}
		}
		SplDualItNext(intern, 0)
	} else {
		intern.SetIsValid(false)
	}
}
func SplCachingItRewind(intern *SplDualItObject) {
	SplDualItRewind(intern)
	intern.GetZcache().Array().Clean()
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_CachingIterator, zend.ZendCeIterator, DIT_CachingIterator)
}
func zim_spl_CachingIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplCachingItRewind(intern)
}
func zim_spl_CachingIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(SplCachingItValid(intern) == types.SUCCESS)
	return
}
func zim_spl_CachingIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplCachingItNext(intern)
}
func zim_spl_CachingIterator_hasNext(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(SplCachingItHasNext(intern) == types.SUCCESS)
	return
}
func zim_spl_CachingIterator___toString(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.HasUCachingFlags(CIT_CALL_TOSTRING | CIT_TOSTRING_USE_KEY | CIT_TOSTRING_USE_CURRENT | CIT_TOSTRING_USE_INNER) {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not fetch string value (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	if intern.IsTostringUseKey() {
		types.ZVAL_COPY(return_value, intern.GetKey())
		operators.ConvertToString(return_value)
		return
	} else if intern.IsTostringUseCurrent() {
		types.ZVAL_COPY(return_value, intern.GetData())
		operators.ConvertToString(return_value)
		return
	}
	if intern.GetZstr().IsString() {
		return_value.SetString(intern.GetZstr().String())
		return
	} else {
		return_value.SetString("")
		return
	}
}
func zim_spl_CachingIterator_offsetSet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.String
	var value *types.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "Sz", &key, &value) == types.FAILURE {
		return
	}
	//value.TryAddRefcount()
	intern.GetZcache().Array().SymtableUpdate(key.GetStr(), value)
}
func zim_spl_CachingIterator_offsetGet(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.String
	var value *types.Zval
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	if lang.Assign(&value, intern.GetZcache().Array().SymtableFind(key.GetStr())) == nil {
		faults.Error(faults.E_NOTICE, fmt.Sprintf("Undefined index: %s", key.GetVal()))
		return
	}
	types.ZVAL_COPY_DEREF(return_value, value)
}
func zim_spl_CachingIterator_offsetUnset(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.String
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	intern.GetZcache().Array().SymtableDel(key.GetStr())
}
func zim_spl_CachingIterator_offsetExists(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var key *types.String
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &key) == types.FAILURE {
		return
	}
	return_value.SetBool(intern.GetZcache().Array().SymtableExists(key.GetStr()))
	return
}
func zim_spl_CachingIterator_getCache(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	types.ZVAL_COPY(return_value, intern.GetZcache())
}
func zim_spl_CachingIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetLong(intern.GetUCachingFlags())
	return
}
func zim_spl_CachingIterator_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var flags zend.ZendLong
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &flags) == types.FAILURE {
		return
	}
	if SplCitCheckFlags(flags) != types.SUCCESS {
		faults.ThrowException(spl_ce_InvalidArgumentException, "Flags must contain only one of CALL_TOSTRING, TOSTRING_USE_KEY, TOSTRING_USE_CURRENT, TOSTRING_USE_INNER", 0)
		return
	}
	if intern.IsCallTostring() && (flags&CIT_CALL_TOSTRING) == 0 {
		faults.ThrowException(spl_ce_InvalidArgumentException, "Unsetting flag CALL_TO_STRING is not possible", 0)
		return
	}
	if intern.IsTostringUseInner() && (flags&CIT_TOSTRING_USE_INNER) == 0 {
		faults.ThrowException(spl_ce_InvalidArgumentException, "Unsetting flag TOSTRING_USE_INNER is not possible", 0)
		return
	}
	if (flags&CIT_FULL_CACHE) != 0 && !intern.IsFullCache() {

		/* clear on (re)enable */

		intern.GetZcache().Array().Clean()

		/* clear on (re)enable */

	}
	intern.SetUCachingFlags(intern.GetUCachingFlags() & ^CIT_PUBLIC | flags&CIT_PUBLIC)
}
func zim_spl_CachingIterator_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if !intern.IsFullCache() {
		faults.ThrowException(spl_ce_BadMethodCallException, fmt.Sprintf("%s does not use a full cache (see CachingIterator::__construct)", types.Z_OBJCE_P(executeData.ThisObjectZval()).Name()), 0)
		return
	}
	return_value.SetLong(intern.GetZcache().Array().Len())
	return
}
func zim_spl_RecursiveCachingIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplDualItConstruct(executeData, return_value, spl_ce_RecursiveCachingIterator, spl_ce_RecursiveIterator, DIT_RecursiveCachingIterator)
}
func zim_spl_RecursiveCachingIterator_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(intern.GetZchildren().IsNotUndef())
	return
}
func zim_spl_RecursiveCachingIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	if intern.GetZchildren().IsNotUndef() {
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_NoRewindIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(intern.GetInnerIterator().GetFuncs().GetValid()(intern.GetInnerIterator()) == types.SUCCESS)
	return
}
func zim_spl_NoRewindIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func zim_spl_EmptyIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	return_value.SetFalse()
	return
}
func zim_spl_EmptyIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	faults.ThrowException(spl_ce_BadMethodCallException, "Accessing the key of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	faults.ThrowException(spl_ce_BadMethodCallException, "Accessing the value of an EmptyIterator", 0)
}
func zim_spl_EmptyIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
}
func SplAppendItNextIterator(intern *SplDualItObject) int {
	SplDualItFree(intern)
	if !(intern.GetZobject().IsUndef()) {
		// zend.ZvalPtrDtor(intern.GetZobject())
		intern.GetZobject().SetUndef()
		intern.SetCe(nil)
		if intern.GetInnerIterator() != nil {
			//zend.ZendIteratorDtor(intern.GetInnerIterator())
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
	var it__1 *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it__1.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it__1
	if zend.ZendParseParametersEx(zpp.FlagQuiet, executeData.NumArgs(), "O", &it, zend.ZendCeIterator) == types.FAILURE {
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
			if intern.GetZobject().Object() == it.Object() {
				break
			}
		}
		SplAppendItFetch(intern)
	}
}
func zim_spl_AppendIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplDualItFetch(intern, 1)
	if intern.GetData().IsNotUndef() {
		var value *types.Zval = intern.GetData()
		types.ZVAL_COPY_DEREF(return_value, value)
	} else {
		return_value.SetNull()
		return
	}
}
func zim_spl_AppendIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
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
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	return_value.SetBool(intern.GetData().IsNotUndef())
	return
}
func zim_spl_AppendIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	SplAppendItNext(intern)
}
func zim_spl_AppendIterator_getIteratorIndex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	APPENDIT_CHECK_CTOR(intern)
	*return_value = *SplArrayIteratorKey(intern.GetZarrayit())
}
func zim_spl_AppendIterator_getArrayIterator(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplDualItObject
	var value *types.Zval
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	var it *SplDualItObject = Z_SPLDUAL_IT_P(executeData.ThisObjectZval())
	if it.GetDitType() == DIT_Unknown {
		faults.ThrowException(spl_ce_LogicException, "The object is in an invalid state as the parent constructor was not called", 0)
		return
	}
	intern = it
	value = intern.GetZarrayit()
	types.ZVAL_COPY_DEREF(return_value, value)
}
func SplIteratorApply(obj *types.Zval, apply_func SplIteratorApplyFuncT, puser any) int {
	var iter *zend.ZendObjectIterator
	var ce *types.ClassEntry = types.Z_OBJCE_P(obj)
	iter = ce.GetGetIterator()(ce, obj, 0)
	if zend.EG__().HasException() {
		goto done
	}
	iter.SetIndex(0)
	if iter.GetFuncs().GetRewind() != nil {
		iter.GetFuncs().GetRewind()(iter)
		if zend.EG__().HasException() {
			goto done
		}
	}
	for iter.GetFuncs().GetValid()(iter) == types.SUCCESS {
		if zend.EG__().HasException() {
			goto done
		}
		if apply_func(iter, puser) == types.ArrayApplyStop || zend.EG__().HasException() {
			goto done
		}
		iter.GetIndex()++
		iter.GetFuncs().GetMoveForward()(iter)
		if zend.EG__().HasException() {
			goto done
		}
	}
done:
	if iter != nil {
		//zend.ZendIteratorDtor(iter)
	}
	if zend.EG__().HasException() {
		return types.FAILURE
	} else {
		return types.SUCCESS
	}
}
func SplIteratorToArrayApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *types.Zval
	var return_value *types.Zval = (*types.Zval)(puser)
	data = iter.GetFuncs().GetGetCurrentData()(iter)
	if zend.EG__().HasException() {
		return types.ArrayApplyStop
	}
	if data == nil {
		return types.ArrayApplyStop
	}
	if iter.GetFuncs().GetGetCurrentKey() != nil {
		var key types.Zval
		iter.GetFuncs().GetGetCurrentKey()(iter, &key)
		if zend.EG__().HasException() {
			return types.ArrayApplyStop
		}
		zend.ArraySetZvalKey(return_value.Array(), &key, data)
		// zend.ZvalPtrDtor(&key)
	} else {
		//data.TryAddRefcount()
		zend.AddNextIndexZval(return_value, data)
	}
	return types.ArrayApplyKeep
}
func SplIteratorToValuesApply(iter *zend.ZendObjectIterator, puser any) int {
	var data *types.Zval
	var return_value *types.Zval = (*types.Zval)(puser)
	data = iter.GetFuncs().GetGetCurrentData()(iter)
	if zend.EG__().HasException() {
		return types.ArrayApplyStop
	}
	if data == nil {
		return types.ArrayApplyStop
	}
	//data.TryAddRefcount()
	zend.AddNextIndexZval(return_value, data)
	return types.ArrayApplyKeep
}
func ZifIteratorToArray(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var obj *types.Zval
	var use_keys bool = 1
	if zend.ZendParseParameters(executeData.NumArgs(), "O|b", &obj, zend.ZendCeTraversable, &use_keys) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	SplIteratorApply(obj, lang.Cond(use_keys != 0, SplIteratorToArrayApply, SplIteratorToValuesApply), any(return_value))
}
func SplIteratorCountApply(iter *zend.ZendObjectIterator, puser any) int {
	*((*zend.ZendLong)(puser))++
	return types.ArrayApplyKeep
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
	if operators.ZvalIsTrue(&retval) {
		result = types.ArrayApplyKeep
	} else {
		result = types.ArrayApplyStop
	}
	// zend.ZvalPtrDtor(&retval)
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
func ZmStartupSplIterators() int {
	spl_ce_RecursiveIterator = zend.RegisterInterface(&types.InternalClassDecl{
		Name:       "RecursiveIterator",
		Interfaces: []*types.ClassEntry{zend.ZendCeIterator},
	})

	spl_ce_RecursiveIteratorIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveIteratorIterator",
		Functions:    spl_funcs_RecursiveIteratorIterator,
		Interfaces:   []*types.ClassEntry{zend.ZendCeIterator, spl_ce_OuterIterator},
		GetIterator:  SplRecursiveItGetIterator,
		CreateObject: spl_RecursiveIteratorIterator_new,
	})
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "LEAVES_ONLY", zend.ZendLong(RIT_LEAVES_ONLY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "SELF_FIRST", zend.ZendLong(RIT_SELF_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CHILD_FIRST", zend.ZendLong(RIT_CHILD_FIRST))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveIteratorIterator, "CATCH_GET_CHILD", zend.ZendLong(RIT_CATCH_GET_CHILD))

	SplHandlersRecItIt = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		Offset:    int((*byte)(&((*SplRecursiveItObject)(nil).GetStd())) - (*byte)(nil)),
		GetMethod: SplRecursiveItGetMethod,
		CloneObj:  nil,
		DtorObj:   spl_RecursiveIteratorIterator_dtor,
		FreeObj:   spl_RecursiveIteratorIterator_free_storage,
	})
	SplHandlersDualIt = *types.NewObjectHandlersEx(zend.StdObjectHandlersPtr, types.ObjectHandlersSetting{
		Offset:    int((*byte)(&((*SplDualItObject)(nil).GetStd())) - (*byte)(nil)),
		GetMethod: SplDualItGetMethod,
		CloneObj:  nil,
		DtorObj:   SplDualItDtor,
		FreeObj:   SplDualItFreeStorage,
	})

	spl_ce_OuterIterator = zend.RegisterInterface(&types.InternalClassDecl{
		Name:       "OuterIterator",
		Functions:  spl_funcs_OuterIterator,
		Interfaces: []*types.ClassEntry{zend.ZendCeIterator},
	})

	spl_ce_IteratorIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "IteratorIterator",
		Interfaces:   []*types.ClassEntry{zend.ZendCeIterator, spl_ce_OuterIterator},
		Functions:    spl_funcs_IteratorIterator,
		CreateObject: SplDualItNew,
	})

	spl_ce_FilterIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "FilterIterator",
		Parent:       spl_ce_IteratorIterator,
		Functions:    spl_funcs_FilterIterator,
		CreateObject: SplDualItNew,
		CeFlags:      types.AccExplicitAbstractClass,
	})

	spl_ce_RecursiveFilterIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveFilterIterator",
		Parent:       spl_ce_FilterIterator,
		Interfaces:   []*types.ClassEntry{spl_ce_RecursiveIterator},
		Functions:    spl_funcs_RecursiveFilterIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_CallbackFilterIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "CallbackFilterIterator",
		Parent:       spl_ce_FilterIterator,
		Functions:    spl_funcs_CallbackFilterIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_RecursiveCallbackFilterIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveCallbackFilterIterator",
		Parent:       spl_ce_CallbackFilterIterator,
		Interfaces:   []*types.ClassEntry{spl_ce_RecursiveIterator},
		Functions:    spl_funcs_RecursiveCallbackFilterIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_ParentIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "ParentIterator",
		Parent:       spl_ce_RecursiveFilterIterator,
		Functions:    spl_funcs_ParentIterator,
		CreateObject: SplDualItNew,
	})

	spl_ce_SeekableIterator = zend.RegisterInterface(&types.InternalClassDecl{
		Name:       "SeekableIterator",
		Functions:  spl_funcs_SeekableIterator,
		Interfaces: []*types.ClassEntry{zend.ZendCeIterator},
	})

	spl_ce_LimitIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "LimitIterator",
		Parent:       spl_ce_IteratorIterator,
		Functions:    spl_funcs_LimitIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_CachingIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "CachingIterator",
		Parent:       spl_ce_IteratorIterator,
		Interfaces:   []*types.ClassEntry{spl_ce_ArrayAccess, spl_ce_Countable},
		Functions:    spl_funcs_CachingIterator,
		CreateObject: SplDualItNew,
	})
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CALL_TOSTRING", zend.ZendLong(CIT_CALL_TOSTRING))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "CATCH_GET_CHILD", zend.ZendLong(CIT_CATCH_GET_CHILD))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_KEY", zend.ZendLong(CIT_TOSTRING_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_CURRENT", zend.ZendLong(CIT_TOSTRING_USE_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "TOSTRING_USE_INNER", zend.ZendLong(CIT_TOSTRING_USE_INNER))
	zend.ZendDeclareClassConstantLong(spl_ce_CachingIterator, "FULL_CACHE", zend.ZendLong(CIT_FULL_CACHE))
	spl_ce_RecursiveCachingIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveCachingIterator",
		Parent:       spl_ce_CachingIterator,
		Interfaces:   []*types.ClassEntry{spl_ce_RecursiveIterator},
		Functions:    spl_funcs_RecursiveCachingIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_NoRewindIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "NoRewindIterator",
		Parent:       spl_ce_IteratorIterator,
		Functions:    spl_funcs_NoRewindIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_AppendIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "AppendIterator",
		Parent:       spl_ce_IteratorIterator,
		Functions:    spl_funcs_AppendIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_InfiniteIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "InfiniteIterator",
		Parent:       spl_ce_IteratorIterator,
		Functions:    spl_funcs_InfiniteIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_RegexIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RegexIterator",
		Parent:       spl_ce_FilterIterator,
		Functions:    spl_funcs_RegexIterator,
		CreateObject: SplDualItNew,
	})
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "USE_KEY", zend.ZendLong(REGIT_USE_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "INVERT_MATCH", zend.ZendLong(REGIT_INVERTED))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "MATCH", zend.ZendLong(REGIT_MODE_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "GET_MATCH", zend.ZendLong(REGIT_MODE_GET_MATCH))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "ALL_MATCHES", zend.ZendLong(REGIT_MODE_ALL_MATCHES))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "SPLIT", zend.ZendLong(REGIT_MODE_SPLIT))
	zend.ZendDeclareClassConstantLong(spl_ce_RegexIterator, "REPLACE", zend.ZendLong(REGIT_MODE_REPLACE))
	zend.ZendDeclarePropertyNull(spl_ce_RegexIterator, "replacement", 0)
	spl_ce_RecursiveRegexIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveRegexIterator",
		Parent:       spl_ce_RegexIterator,
		Interfaces:   []*types.ClassEntry{spl_ce_RecursiveIterator},
		Functions:    spl_funcs_RecursiveRegexIterator,
		CreateObject: SplDualItNew,
	})
	spl_ce_EmptyIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "EmptyIterator",
		Interfaces:   []*types.ClassEntry{zend.ZendCeIterator},
		Functions:    spl_funcs_EmptyIterator,
		CreateObject: nil,
	})
	spl_ce_RecursiveTreeIterator = zend.RegisterClass(&types.InternalClassDecl{
		Name:         "RecursiveTreeIterator",
		Parent:       spl_ce_RecursiveIteratorIterator,
		Functions:    spl_funcs_RecursiveTreeIterator,
		CreateObject: spl_RecursiveTreeIterator_new,
	})
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_CURRENT", zend.ZendLong(RTIT_BYPASS_CURRENT))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "BYPASS_KEY", zend.ZendLong(RTIT_BYPASS_KEY))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_LEFT", zend.ZendLong(0))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_HAS_NEXT", zend.ZendLong(1))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_MID_LAST", zend.ZendLong(2))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_HAS_NEXT", zend.ZendLong(3))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_END_LAST", zend.ZendLong(4))
	zend.ZendDeclareClassConstantLong(spl_ce_RecursiveTreeIterator, "PREFIX_RIGHT", zend.ZendLong(5))
	return types.SUCCESS
}
