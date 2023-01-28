// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

func _zendObjectStdInit(object *ZendObject, ce *ZendClassEntry) {
	GC_SET_REFCOUNT(object, 1)
	GC_TYPE_INFO(object) = IS_OBJECT | GC_COLLECTABLE<<GC_FLAGS_SHIFT
	object.SetCe(ce)
	object.SetProperties(nil)
	ZendObjectsStorePut(object)
	if ce.IsUseGuards() {
		ZVAL_UNDEF(object.GetPropertiesTable() + object.GetCe().GetDefaultPropertiesCount())
	}
}
func ZendObjectStdInit(object *ZendObject, ce *ZendClassEntry) { _zendObjectStdInit(object, ce) }
func ZendObjectStdDtor(object *ZendObject) {
	var p *Zval
	var end *Zval
	if object.GetProperties() != nil {
		if (GC_FLAGS(object.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0 {
			if GC_DELREF(object.GetProperties()) == 0 && GC_TYPE(object.GetProperties()) != IS_NULL {
				ZendArrayDestroy(object.GetProperties())
			}
		}
	}
	p = object.GetPropertiesTable()
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		end = p + object.GetCe().GetDefaultPropertiesCount()
		for {
			if Z_REFCOUNTED_P(p) {
				if Z_ISREF_P(p) && (core.ZEND_DEBUG != 0 || ZEND_REF_HAS_TYPE_SOURCES(p.GetRef())) {
					var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(object, p)
					if prop_info.GetType() != 0 {
						ZEND_REF_DEL_TYPE_SOURCE(p.GetRef(), prop_info)
					}
				}
				IZvalPtrDtor(p)
			}
			p++
			if p == end {
				break
			}
		}
	}
	if object.GetCe().IsUseGuards() {
		if p.IsType(IS_STRING) {
			ZvalPtrDtorStr(p)
		} else if p.IsType(IS_ARRAY) {
			var guards *HashTable
			guards = p.GetArr()
			ZEND_ASSERT(guards != nil)
			ZendHashDestroy(guards)
			FREE_HASHTABLE(guards)
		}
	}
	if (GC_FLAGS(object) & IS_OBJ_WEAKLY_REFERENCED) != 0 {
		ZendWeakrefsNotify(object)
	}
}
func ZendObjectsDestroyObject(object *ZendObject) {
	var destructor *ZendFunction = object.GetCe().GetDestructor()
	if destructor != nil {
		var old_exception *ZendObject
		var orig_fake_scope *ZendClassEntry
		var fci ZendFcallInfo
		var fcic ZendFcallInfoCache
		var ret Zval
		if destructor.GetOpArray().HasFnFlags(ZEND_ACC_PRIVATE | ZEND_ACC_PROTECTED) {
			if destructor.GetOpArray().IsPrivate() {

				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */

				if ExecutorGlobals.GetCurrentExecuteData() != nil {
					var scope *ZendClassEntry = ZendGetExecutedScope()
					if object.GetCe() != scope {
						ZendThrowError(nil, "Call to private %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					ZendError(E_WARNING, "Call to private %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}

				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */

			} else {

				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */

				if ExecutorGlobals.GetCurrentExecuteData() != nil {
					var scope *ZendClassEntry = ZendGetExecutedScope()
					if ZendCheckProtected(ZendGetFunctionRootClass(destructor), scope) == 0 {
						ZendThrowError(nil, "Call to protected %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					ZendError(E_WARNING, "Call to protected %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}

				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */

			}
		}
		GC_ADDREF(object)

		/* Make sure that destructors are protected from previously thrown exceptions.
		 * For example, if an exception was thrown in a function and when the function's
		 * local variable destruction results in a destructor being called.
		 */

		old_exception = nil
		if ExecutorGlobals.GetException() != nil {
			if ExecutorGlobals.GetException() == object {
				ZendErrorNoreturn(E_CORE_ERROR, "Attempt to destruct pending exception")
			} else {
				old_exception = ExecutorGlobals.GetException()
				ExecutorGlobals.SetException(nil)
			}
		}
		orig_fake_scope = ExecutorGlobals.GetFakeScope()
		ExecutorGlobals.SetFakeScope(nil)
		ZVAL_UNDEF(&ret)
		fci.SetSize(b.SizeOf("fci"))
		fci.SetObject(object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		ZVAL_UNDEF(fci.GetFunctionName())
		fcic.SetFunctionHandler(destructor)
		fcic.SetCalledScope(object.GetCe())
		fcic.SetObject(object)
		ZendCallFunction(&fci, &fcic)
		ZvalPtrDtor(&ret)
		if old_exception != nil {
			if ExecutorGlobals.GetException() != nil {
				ZendExceptionSetPrevious(ExecutorGlobals.GetException(), old_exception)
			} else {
				ExecutorGlobals.SetException(old_exception)
			}
		}
		OBJ_RELEASE(object)
		ExecutorGlobals.SetFakeScope(orig_fake_scope)
	}
}
func ZendObjectsNew(ce *ZendClassEntry) *ZendObject {
	var object *ZendObject = Emalloc(b.SizeOf("zend_object") + ZendObjectPropertiesSize(ce))
	_zendObjectStdInit(object, ce)
	object.SetHandlers(&StdObjectHandlers)
	return object
}
func ZendObjectsCloneMembers(new_object *ZendObject, old_object *ZendObject) {
	if old_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var src *Zval = old_object.GetPropertiesTable()
		var dst *Zval = new_object.GetPropertiesTable()
		var end *Zval = src + old_object.GetCe().GetDefaultPropertiesCount()
		for {
			IZvalPtrDtor(dst)
			ZVAL_COPY_VALUE_PROP(dst, src)
			ZvalAddRef(dst)
			if Z_ISREF_P(dst) && (core.ZEND_DEBUG != 0 || ZEND_REF_HAS_TYPE_SOURCES(dst.GetRef())) {
				var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(new_object, dst)
				if prop_info.GetType() != 0 {
					ZEND_REF_ADD_TYPE_SOURCE(dst.GetRef(), prop_info)
				}
			}
			src++
			dst++
			if src == end {
				break
			}
		}
	} else if old_object.GetProperties() != nil && old_object.GetCe().GetClone() == nil {

		/* fast copy */

		if old_object.GetHandlers() == &StdObjectHandlers {
			if (GC_FLAGS(old_object.GetProperties()) & IS_ARRAY_IMMUTABLE) == 0 {
				GC_ADDREF(old_object.GetProperties())
			}
			new_object.SetProperties(old_object.GetProperties())
			return
		}

		/* fast copy */

	}
	if old_object.GetProperties() != nil && old_object.GetProperties().GetNNumOfElements() {
		var prop *Zval
		var new_prop Zval
		var num_key ZendUlong
		var key *ZendString
		if new_object.GetProperties() == nil {
			new_object.SetProperties(ZendNewArray(old_object.GetProperties().GetNNumOfElements()))
			ZendHashRealInitMixed(new_object.GetProperties())
		} else {
			ZendHashExtend(new_object.GetProperties(), new_object.GetProperties().GetNNumUsed()+old_object.GetProperties().GetNNumOfElements(), 0)
		}
		new_object.GetProperties().GetUFlags() |= old_object.GetProperties().GetUFlags() & HASH_FLAG_HAS_EMPTY_IND
		for {
			var __ht *HashTable = old_object.GetProperties()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = _p.GetVal()

				if _z.IsType(IS_UNDEF) {
					continue
				}
				num_key = _p.GetH()
				key = _p.GetKey()
				prop = _z
				if prop.IsType(IS_INDIRECT) {
					ZVAL_INDIRECT(&new_prop, new_object.GetPropertiesTable()+(prop.GetZv()-old_object.GetPropertiesTable()))
				} else {
					ZVAL_COPY_VALUE(&new_prop, prop)
					ZvalAddRef(&new_prop)
				}
				if key != nil {
					_zendHashAppend(new_object.GetProperties(), key, &new_prop)
				} else {
					ZendHashIndexAddNew(new_object.GetProperties(), num_key, &new_prop)
				}
			}
			break
		}
	}
	if old_object.GetCe().GetClone() != nil {
		var fci ZendFcallInfo
		var fcic ZendFcallInfoCache
		var ret Zval
		GC_ADDREF(new_object)
		ZVAL_UNDEF(&ret)
		fci.SetSize(b.SizeOf("fci"))
		fci.SetObject(new_object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		ZVAL_UNDEF(fci.GetFunctionName())
		fcic.SetFunctionHandler(new_object.GetCe().GetClone())
		fcic.SetCalledScope(new_object.GetCe())
		fcic.SetObject(new_object)
		ZendCallFunction(&fci, &fcic)
		ZvalPtrDtor(&ret)
		OBJ_RELEASE(new_object)
	}
}
func ZendObjectsCloneObj(zobject *Zval) *ZendObject {
	var old_object *ZendObject
	var new_object *ZendObject

	/* assume that create isn't overwritten, so when clone depends on the
	 * overwritten one then it must itself be overwritten */

	old_object = zobject.GetObj()
	new_object = ZendObjectsNew(old_object.GetCe())

	/* zend_objects_clone_members() expect the properties to be initialized. */

	if new_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var p *Zval = new_object.GetPropertiesTable()
		var end *Zval = p + new_object.GetCe().GetDefaultPropertiesCount()
		for {
			ZVAL_UNDEF(p)
			p++
			if p == end {
				break
			}
		}
	}
	ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
