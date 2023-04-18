package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZendObjectsNew(ce *types2.ClassEntry) *types2.ZendObject {
	handle := EG__().NextObjectHandle()
	return types2.NewObject(ce, handle, StdObjectHandlersPtr)
}
func ZendObjectStdInit(object *types2.ZendObject, ce *types2.ClassEntry) {
	handle := EG__().NextObjectHandle()
	object.Init(ce, handle)
}
func ZendObjectStdDtor(object *types2.ZendObject) {
	var p *types2.Zval
	var end *types2.Zval
	if object.GetProperties() != nil {
		if (object.GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
			if object.GetProperties().DelRefcount() == 0 && object.GetProperties().GetGcType() != types2.IS_NULL {
				object.GetProperties().DestroyEx()
			}
		}
	}
	p = object.GetPropertiesTable()
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		end = p + object.GetCe().GetDefaultPropertiesCount()
		for {
			if p.IsRefcounted() {
				if p.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(p.Reference()) {
					var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(object, p)
					if prop_info.GetType() != 0 {
						ZEND_REF_DEL_TYPE_SOURCE(p.Reference(), prop_info)
					}
				}
				// IZvalPtrDtor(p)
			}
			p++
			if p == end {
				break
			}
		}
	}
	if object.GetCe().IsUseGuards() {
		if p.IsString() {

		} else if p.IsArray() {
			var guards *types2.Array
			guards = p.Array()
			b.Assert(guards != nil)
			guards.Destroy()
			FREE_HASHTABLE(guards)
		}
	}
	if (object.GetGcFlags() & types2.IS_OBJ_WEAKLY_REFERENCED) != 0 {
		ZendWeakrefsNotify(object)
	}
}
func ZendObjectsDestroyObject(object *types2.ZendObject) {
	var destructor types2.IFunction = object.GetCe().GetDestructor()
	if destructor != nil {
		var old_exception *types2.ZendObject
		var orig_fake_scope *types2.ClassEntry
		var fci types2.ZendFcallInfo
		var fcic types2.ZendFcallInfoCache
		var ret types2.Zval
		if destructor.GetOpArray().HasFnFlags(AccPrivate | AccProtected) {
			if destructor.GetOpArray().IsPrivate() {
				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */
				if CurrEX() != nil {
					var scope *types2.ClassEntry = ZendGetExecutedScope()
					if object.GetCe() != scope {
						faults.ThrowError(nil, "Call to private %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					faults.Error(faults.E_WARNING, "Call to private %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}
			} else {
				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */
				if CurrEX() != nil {
					var scope *types2.ClassEntry = ZendGetExecutedScope()
					if !ZendCheckProtected(ZendGetFunctionRootClass(destructor), scope) {
						faults.ThrowError(nil, "Call to protected %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), b.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					faults.Error(faults.E_WARNING, "Call to protected %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}
			}
		}
		// 		object.AddRefcount()

		/* Make sure that destructors are protected from previously thrown exceptions.
		 * For example, if an exception was thrown in a function and when the function's
		 * local variable destruction results in a destructor being called.
		 */

		old_exception = nil
		if EG__().GetException() != nil {
			if EG__().GetException() == object {
				faults.ErrorNoreturn(faults.E_CORE_ERROR, "Attempt to destruct pending exception")
			} else {
				old_exception = EG__().GetException()
				EG__().SetException(nil)
			}
		}
		orig_fake_scope = EG__().GetFakeScope()
		EG__().SetFakeScope(nil)
		ret.SetUndef()
		fci.SetSize(b.SizeOf("fci"))
		fci.SetObject(object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		fci.GetFunctionName().SetUndef()
		fcic.SetFunctionHandler(destructor)
		fcic.SetCalledScope(object.GetCe())
		fcic.SetObject(object)
		ZendCallFunction(&fci, &fcic)
		// ZvalPtrDtor(&ret)
		if old_exception != nil {
			if EG__().GetException() != nil {
				faults.ExceptionSetPrevious(EG__().GetException(), old_exception)
			} else {
				EG__().SetException(old_exception)
			}
		}
		// OBJ_RELEASE(object)
		EG__().SetFakeScope(orig_fake_scope)
	}
}
func ZendObjectsCloneMembers(new_object *types2.ZendObject, old_object *types2.ZendObject) {
	if old_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var src *types2.Zval = old_object.GetPropertiesTable()
		var dst *types2.Zval = new_object.GetPropertiesTable()
		var end *types2.Zval = src + old_object.GetCe().GetDefaultPropertiesCount()
		for {
			// IZvalPtrDtor(dst)
			types2.ZVAL_COPY_VALUE_PROP(dst, src)
			ZvalAddRef(dst)
			if dst.IsReference() && ZEND_REF_HAS_TYPE_SOURCES(dst.Reference()) {
				var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(new_object, dst)
				if prop_info.GetType() != 0 {
					ZEND_REF_ADD_TYPE_SOURCE(dst.Reference(), prop_info)
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

		if old_object.GetHandlers() == StdObjectHandlersPtr {
			if (old_object.GetProperties().GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) == 0 {
				old_object.GetProperties().AddRefcount()
			}
			new_object.SetProperties(old_object.GetProperties())
			return
		}

		/* fast copy */

	}
	if old_object.GetProperties() != nil && old_object.GetProperties().Len() {
		var prop *types2.Zval
		var new_prop types2.Zval
		var num_key ZendUlong
		var key *types2.String
		if new_object.GetProperties() == nil {
			new_object.SetProperties(types2.NewArray(old_object.GetProperties().Len()))
			types2.ZendHashRealInitMixed(new_object.GetProperties())
		} else {
			new_object.GetProperties().Extend(new_object.GetProperties().GetNNumUsed() + old_object.GetProperties().Len())
		}
		new_object.GetProperties().CopyFlags(old_object.GetProperties())
		new_object.GetProperties().SetIteratorsCount(old_object.GetProperties().GetIteratorsCount())
		new_object.GetProperties().MarkHasEmptyIndex()
		var __ht *types2.Array = old_object.GetProperties()
		for _, _p := range __ht.ForeachData() {
			var _z *types2.Zval = _p.GetVal()

			num_key = _p.GetH()
			key = _p.GetKey()
			prop = _z
			if prop.IsIndirect() {
				new_prop.SetIndirect(new_object.GetPropertiesTable() + (prop.Indirect() - old_object.GetPropertiesTable()))
			} else {
				types2.ZVAL_COPY_VALUE(&new_prop, prop)
				ZvalAddRef(&new_prop)
			}
			if key != nil {
				types._zendHashAppend(new_object.GetProperties(), key, &new_prop)
			} else {
				new_object.GetProperties().IndexAddNew(num_key, &new_prop)
			}
		}
	}
	if old_object.GetCe().GetClone() != nil {
		var fci types2.ZendFcallInfo
		var fcic types2.ZendFcallInfoCache
		var ret types2.Zval
		// 		new_object.AddRefcount()
		ret.SetUndef()
		fci.SetSize(b.SizeOf("fci"))
		fci.SetObject(new_object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		fci.GetFunctionName().SetUndef()
		fcic.SetFunctionHandler(new_object.GetCe().GetClone())
		fcic.SetCalledScope(new_object.GetCe())
		fcic.SetObject(new_object)
		ZendCallFunction(&fci, &fcic)
		// ZvalPtrDtor(&ret)
		// OBJ_RELEASE(new_object)
	}
}
func ZendObjectsCloneObj(zobject *types2.Zval) *types2.ZendObject {
	var old_object *types2.ZendObject
	var new_object *types2.ZendObject

	/* assume that create isn't overwritten, so when clone depends on the
	 * overwritten one then it must itself be overwritten */

	old_object = zobject.Object()
	new_object = ZendObjectsNew(old_object.GetCe())

	/* zend_objects_clone_members() expect the properties to be initialized. */

	if new_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var p *types2.Zval = new_object.GetPropertiesTable()
		var end *types2.Zval = p + new_object.GetCe().GetDefaultPropertiesCount()
		for {
			p.SetUndef()
			p++
			if p == end {
				break
			}
		}
	}
	ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
