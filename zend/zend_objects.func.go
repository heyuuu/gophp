package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZendObjectStdDtorEx(properties []types.Zval, ce *types.ClassEntry) {
	defaultPropertiesCount := ce.GetDefaultPropertiesCount()
	if defaultPropertiesCount != 0 {
		for propNum := range properties[:defaultPropertiesCount] {
			prop := &properties[propNum]
			if prop.IsRef() && ZEND_REF_HAS_TYPE_SOURCES(prop.Ref()) {
				var propInfo = ce.GetPropertyInfo(propNum)
				if propInfo.GetType() != nil {
					ZEND_REF_DEL_TYPE_SOURCE(prop.Ref(), propInfo)
				}
			}
		}
	}

	if ce.IsUseGuards() {
		p := &properties[defaultPropertiesCount]
		if p.IsArray() {
			guards := p.Array()
			b.Assert(guards != nil)
			guards.Destroy()
		}
	}
}

func ZendObjectStdDtor(object *types.Object) {
	ZendObjectStdDtorEx(object.GetPropertiesTable(), object.GetCe())
}
func ZendObjectsDestroyObject(object *types.Object) {
	var destructor types.IFunction = object.GetCe().GetDestructor()
	if destructor != nil {
		var old_exception *types.Object
		var orig_fake_scope *types.ClassEntry
		if destructor.GetOpArray().HasFnFlags(types.AccPrivate | types.AccProtected) {
			if destructor.GetOpArray().IsPrivate() {
				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */
				if CurrEX() != nil {
					var scope *types.ClassEntry = ZendGetExecutedScope()
					if object.GetCe() != scope {
						faults.ThrowError(nil, "Call to private %s::__destruct() from context '%s'", object.GetCe().Name(), lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""))
						return
					}
				} else {
					faults.Error(faults.E_WARNING, "Call to private %s::__destruct() from context '' during shutdown ignored", object.GetCe().Name())
					return
				}
			} else {
				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */
				if CurrEX() != nil {
					var scope *types.ClassEntry = ZendGetExecutedScope()
					if !ZendCheckProtected(ZendGetFunctionRootClass(destructor), scope) {
						faults.ThrowError(nil, "Call to protected %s::__destruct() from context '%s'", object.GetCe().Name(), lang.CondF1(scope != nil, func() []byte { return scope.Name() }, ""))
						return
					}
				} else {
					faults.Error(faults.E_WARNING, "Call to protected %s::__destruct() from context '' during shutdown ignored", object.GetCe().Name())
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
		if EG__().HasException() {
			if EG__().GetException() == object {
				faults.ErrorNoreturn(faults.E_CORE_ERROR, "Attempt to destruct pending exception")
			} else {
				old_exception = EG__().GetException()
				EG__().SetException(nil)
			}
		}
		orig_fake_scope = EG__().GetFakeScope()
		EG__().SetFakeScope(nil)

		var fci = types.InitFCallInfo(object, nil)
		var fcic types.ZendFcallInfoCache
		fcic.SetFunctionHandler(destructor)
		fcic.SetCalledScope(object.GetCe())
		fcic.SetObject(object)

		ZendCallFunction(fci, &fcic)
		if old_exception != nil {
			if EG__().HasException() {
				faults.ExceptionSetPrevious(EG__().GetException(), old_exception)
			} else {
				EG__().SetException(old_exception)
			}
		}
		EG__().SetFakeScope(orig_fake_scope)
	}
}
func ZendObjectsCloneMembers(new_object *types.Object, old_object *types.Object) {
	if old_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var src *types.Zval = old_object.GetPropertiesTable()
		var dst *types.Zval = new_object.GetPropertiesTable()
		var end *types.Zval = src + old_object.GetCe().GetDefaultPropertiesCount()
		for {
			// IZvalPtrDtor(dst)
			types.ZVAL_COPY_VALUE_PROP(dst, src)
			//ZvalAddRef(dst)
			if dst.IsRef() && ZEND_REF_HAS_TYPE_SOURCES(dst.Ref()) {
				var prop_info *types.PropertyInfo = ZendGetPropertyInfoForSlot(new_object, dst)
				if prop_info.GetType() != 0 {
					ZEND_REF_ADD_TYPE_SOURCE(dst.Ref(), prop_info)
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
			new_object.SetProperties(old_object.GetProperties())
			return
		}

		/* fast copy */

	}
	if old_object.GetProperties() != nil && old_object.GetProperties().Len() {
		if new_object.GetProperties() == nil {
			new_object.SetProperties(types.NewArrayCap(old_object.GetProperties().Len()))
		}
		new_object.GetProperties().CopyFlags(old_object.GetProperties())
		new_object.GetProperties().MarkHasEmptyIndex()
		old_object.GetProperties().Foreach(func(key types.ArrayKey, prop *types.Zval) {
			var new_prop types.Zval
			if prop.IsIndirect() {
				new_prop.SetIndirect(new_object.GetPropertiesTable() + (prop.Indirect() - old_object.GetPropertiesTable()))
			} else {
				types.ZVAL_COPY_VALUE(&new_prop, prop)
				//ZvalAddRef(&new_prop)
			}
			if key.IsStrKey() {
				new_object.GetProperties().KeyAddNew(key.StrKey(), &new_prop)
			} else {
				new_object.GetProperties().IndexAddNew(key.IdxKey(), &new_prop)
			}
		})
	}
	if old_object.GetCe().GetClone() != nil {
		var fci *types.ZendFcallInfo = types.InitFCallInfo(new_object, nil)

		var fcic types.ZendFcallInfoCache
		fcic.SetFunctionHandler(new_object.GetCe().GetClone())
		fcic.SetCalledScope(new_object.GetCe())
		fcic.SetObject(new_object)

		ZendCallFunction(fci, &fcic)
	}
}
func ZendObjectsCloneObjEx(oldObject *types.Object) *types.Object {
	/* assume that create isn't overwritten, so when clone depends on the
	 * overwritten one then it must itself be overwritten */
	newObject := types.NewStdObjectSkipPropertiesInit(oldObject.GetCe())
	ZendObjectsCloneMembers(newObject, oldObject)
	return newObject
}
