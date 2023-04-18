package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZendWeakrefFrom(o *types2.ZendObject) *ZendWeakref {
	return (*ZendWeakref)((*byte)(o) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
}
func ZendWeakrefFetch(z *types2.Zval) *ZendWeakref { return ZendWeakrefFrom(z.Object()) }
func ZendWeakrefUnref(zv *types2.Zval) {
	var wr = (*ZendWeakref)(zv.Ptr())
	wr.GetReferent().DelGcFlags(types2.IS_OBJ_WEAKLY_REFERENCED)
	wr.SetReferent(nil)
}
func ZendWeakrefsInit() {
	EG__().GetWeakrefs().Init(8, ZendWeakrefUnref)
}
func ZendWeakrefsNotify(object *types2.ZendObject) {
	types2.ZendHashIndexDel(EG__().GetWeakrefs(), ZendUlong(object))
}
func ZendWeakrefsShutdown() { EG__().GetWeakrefs().Destroy() }
func ZendWeakrefNew(ce *types2.ClassEntry) *types2.ZendObject {
	var wr *ZendWeakref = ZendObjectAlloc(b.SizeOf("zend_weakref"), ZendCeWeakref)
	ZendObjectStdInit(wr.GetStd(), ZendCeWeakref)
	wr.GetStd().SetHandlers(&ZendWeakrefHandlers)
	return wr.GetStd()
}
func ZendWeakrefFind(referent *types2.Zval, return_value *types2.Zval) types2.ZendBool {
	var wr *ZendWeakref = types2.ZendHashIndexFindPtr(EG__().GetWeakrefs(), ZendUlong(referent.Object()))
	if wr == nil {
		return 0
	}
	wr.GetStd().AddRefcount()
	return_value.SetObject(wr.GetStd())
	return 1
}
func ZendWeakrefCreate(referent *types2.Zval, return_value *types2.Zval) {
	var wr *ZendWeakref
	ObjectInitEx(return_value, ZendCeWeakref)
	wr = ZendWeakrefFetch(return_value)
	wr.SetReferent(referent.Object())
	types2.ZendHashIndexAddPtr(EG__().GetWeakrefs(), ZendUlong(wr.GetReferent()), wr)
	wr.GetReferent().AddGcFlags(types2.IS_OBJ_WEAKLY_REFERENCED)
}
func ZendWeakrefGet(weakref *types2.Zval, return_value *types2.Zval) {
	var wr *ZendWeakref = ZendWeakrefFetch(weakref)
	if wr.GetReferent() != nil {
		return_value.SetObject(wr.GetReferent())
		// 		return_value.AddRefcount()
	}
}
func ZendWeakrefFree(zo *types2.ZendObject) {
	var wr *ZendWeakref = ZendWeakrefFrom(zo)
	if wr.GetReferent() != nil {
		types2.ZendHashIndexDel(EG__().GetWeakrefs(), ZendUlong(wr.GetReferent()))
	}
	ZendObjectStdDtor(wr.GetStd())
}
func ZendWeakrefUnsupported(thing string) {
	faults.ThrowError(nil, "WeakReference objects do not support "+thing)
}
func ZendWeakrefNoWrite(object *types2.Zval, member *types2.Zval, value *types2.Zval, rtc *any) *types2.Zval {
	ZendWeakrefUnsupported("properties")
	return EG__().GetUninitializedZval()
}
func ZendWeakrefNoRead(object *types2.Zval, member *types2.Zval, type_ int, rtc *any, rv *types2.Zval) *types2.Zval {
	if EG__().GetException() == nil {
		ZendWeakrefUnsupported("properties")
	}
	return EG__().GetUninitializedZval()
}
func ZendWeakrefNoReadPtr(object *types2.Zval, member *types2.Zval, type_ int, rtc *any) *types2.Zval {
	ZendWeakrefUnsupported("property references")
	return nil
}
func ZendWeakrefNoIsset(object *types2.Zval, member *types2.Zval, hse int, rtc *any) int {
	if hse != 2 {
		ZendWeakrefUnsupported("properties")
	}
	return 0
}
func ZendWeakrefNoUnset(object *types2.Zval, member *types2.Zval, rtc *any) {
	ZendWeakrefUnsupported("properties")
}
func zim_WeakReference___construct(executeData *ZendExecuteData, return_value *types2.Zval) {
	faults.ThrowError(nil, "Direct instantiation of 'WeakReference' is not allowed, "+"use WeakReference::create instead")
}
func zim_WeakReference_create(executeData *ZendExecuteData, return_value *types2.Zval) {
	var referent *types2.Zval
	for {
		var _flags int = zpp.FlagThrow
		var _min_num_args int = 1
		var _max_num_args int = 1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			referent = fp.ParseObject()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if ZendWeakrefFind(referent, return_value) != 0 {
		return
	}
	ZendWeakrefCreate(referent, return_value)
}
func zim_WeakReference_get(executeData *ZendExecuteData, return_value *types2.Zval) {
	for {
		var _flags int = zpp.FlagThrow
		var _min_num_args int = 0
		var _max_num_args int = 0

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	ZendWeakrefGet(getThis(), return_value)
}
func ZendRegisterWeakrefCe() {
	var ce types2.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types2.NewString("WeakReference"))
	ce.SetBuiltinFunctions(ZendWeakrefMethods)
	ZendCeWeakref = ZendRegisterInternalClass(&ce)
	ZendCeWeakref.SetIsFinal(true)
	ZendCeWeakref.SetCreateObject(ZendWeakrefNew)
	ZendCeWeakref.SetSerialize(ZendClassSerializeDeny)
	ZendCeWeakref.SetUnserialize(ZendClassUnserializeDeny)
	memcpy(&ZendWeakrefHandlers, ZendGetStdObjectHandlers(), b.SizeOf("zend_object_handlers"))
	ZendWeakrefHandlers.SetOffset(zend_long((*byte)(&((*ZendWeakref)(nil).GetStd())) - (*byte)(nil)))
	ZendWeakrefHandlers.SetFreeObj(ZendWeakrefFree)
	ZendWeakrefHandlers.SetReadProperty(ZendWeakrefNoRead)
	ZendWeakrefHandlers.SetWriteProperty(ZendWeakrefNoWrite)
	ZendWeakrefHandlers.SetHasProperty(ZendWeakrefNoIsset)
	ZendWeakrefHandlers.SetUnsetProperty(ZendWeakrefNoUnset)
	ZendWeakrefHandlers.SetGetPropertyPtrPtr(ZendWeakrefNoReadPtr)
	ZendWeakrefHandlers.SetCloneObj(nil)
}
