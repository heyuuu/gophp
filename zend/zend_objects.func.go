// <<generate>>

package zend

import (
	b "sik/builtin"
)

func _zendObjectStdInit(object *ZendObject, ce *ZendClassEntry) {
	object.SetRefcount(1)
	object.GetGcTypeInfo() = IS_OBJECT | GC_COLLECTABLE<<GC_FLAGS_SHIFT
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
		if (object.GetProperties().GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
			if object.GetProperties().DelRefcount() == 0 && object.GetProperties().GetGcType() != IS_NULL {
				object.GetProperties().DestroyEx()
			}
		}
	}
	p = object.GetPropertiesTable()
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		end = p + object.GetCe().GetDefaultPropertiesCount()
		for {
			if Z_REFCOUNTED_P(p) {
				if Z_ISREF_P(p) && ZEND_REF_HAS_TYPE_SOURCES(p.GetRef()) {
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
			guards.Destroy()
			FREE_HASHTABLE(guards)
		}
	}
	if (object.GetGcFlags() & IS_OBJ_WEAKLY_REFERENCED) != 0 {
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

				if EG__().GetCurrentExecuteData() != nil {
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

				if EG__().GetCurrentExecuteData() != nil {
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
		object.AddRefcount()

		/* Make sure that destructors are protected from previously thrown exceptions.
		 * For example, if an exception was thrown in a function and when the function's
		 * local variable destruction results in a destructor being called.
		 */

		old_exception = nil
		if EG__().GetException() != nil {
			if EG__().GetException() == object {
				ZendErrorNoreturn(E_CORE_ERROR, "Attempt to destruct pending exception")
			} else {
				old_exception = EG__().GetException()
				EG__().SetException(nil)
			}
		}
		orig_fake_scope = EG__().GetFakeScope()
		EG__().SetFakeScope(nil)
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
			if EG__().GetException() != nil {
				ZendExceptionSetPrevious(EG__().GetException(), old_exception)
			} else {
				EG__().SetException(old_exception)
			}
		}
		OBJ_RELEASE(object)
		EG__().SetFakeScope(orig_fake_scope)
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
			if Z_ISREF_P(dst) && ZEND_REF_HAS_TYPE_SOURCES(dst.GetRef()) {
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
			if (old_object.GetProperties().GetGcFlags() & IS_ARRAY_IMMUTABLE) == 0 {
				old_object.GetProperties().AddRefcount()
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
			new_object.GetProperties().Extend(new_object.GetProperties().GetNNumUsed() + old_object.GetProperties().GetNNumOfElements())
		}
		new_object.GetProperties().GetUFlags() |= old_object.GetProperties().GetUFlags() & HASH_FLAG_HAS_EMPTY_IND
		var __ht *HashTable = old_object.GetProperties()
		for _, _p := range __ht.foreachData() {
			var _z *Zval = _p.GetVal()

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
				new_object.GetProperties().IndexAddNewH(num_key, &new_prop)
			}
		}
	}
	if old_object.GetCe().GetClone() != nil {
		var fci ZendFcallInfo
		var fcic ZendFcallInfoCache
		var ret Zval
		new_object.AddRefcount()
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
