package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func ZendWeakrefFrom(o *types.ZendObject) *ZendWeakref {
	return (*ZendWeakref)((*byte)(o) - zend_long((*byte)(&((*ZendWeakref)(nil).GetStd()))-(*byte)(nil)))
}
func ZendWeakrefFetch(z *types.Zval) *ZendWeakref { return ZendWeakrefFrom(z.Object()) }
func ZendWeakrefUnref(zv *types.Zval) {
	var wr = (*ZendWeakref)(zv.Ptr())
	wr.GetReferent().DelGcFlags(types.IS_OBJ_WEAKLY_REFERENCED)
	wr.SetReferent(nil)
}
func ZendWeakrefsInit() {
	EG__().GetWeakrefs().Init(8, ZendWeakrefUnref)
}
func ZendWeakrefsNotify(object *types.ZendObject) {
	types.ZendHashIndexDel(EG__().GetWeakrefs(), ZendUlong(object))
}
func ZendWeakrefsShutdown() { EG__().GetWeakrefs().Destroy() }
func ZendWeakrefNew(ce *types.ClassEntry) *types.ZendObject {
	var wr *ZendWeakref = ZendObjectAlloc(b.SizeOf("zend_weakref"), ZendCeWeakref)
	ZendObjectStdInit(wr.GetStd(), ZendCeWeakref)
	wr.GetStd().SetHandlers(&ZendWeakrefHandlers)
	return wr.GetStd()
}
func ZendWeakrefFind(referent *types.Zval, return_value *types.Zval) types.ZendBool {
	var wr *ZendWeakref = types.ZendHashIndexFindPtr(EG__().GetWeakrefs(), ZendUlong(referent.Object()))
	if wr == nil {
		return 0
	}
	wr.GetStd().AddRefcount()
	return_value.SetObject(wr.GetStd())
	return 1
}
func ZendWeakrefCreate(referent *types.Zval, return_value *types.Zval) {
	var wr *ZendWeakref
	ObjectInitEx(return_value, ZendCeWeakref)
	wr = ZendWeakrefFetch(return_value)
	wr.SetReferent(referent.Object())
	types.ZendHashIndexAddPtr(EG__().GetWeakrefs(), ZendUlong(wr.GetReferent()), wr)
	wr.GetReferent().AddGcFlags(types.IS_OBJ_WEAKLY_REFERENCED)
}
func ZendWeakrefGet(weakref *types.Zval, return_value *types.Zval) {
	var wr *ZendWeakref = ZendWeakrefFetch(weakref)
	if wr.GetReferent() != nil {
		return_value.SetObject(wr.GetReferent())
		// 		return_value.AddRefcount()
	}
}
func ZendWeakrefFree(zo *types.ZendObject) {
	var wr *ZendWeakref = ZendWeakrefFrom(zo)
	if wr.GetReferent() != nil {
		types.ZendHashIndexDel(EG__().GetWeakrefs(), ZendUlong(wr.GetReferent()))
	}
	ZendObjectStdDtor(wr.GetStd())
}
func ZendWeakrefUnsupported(thing string) {
	faults.ThrowError(nil, "WeakReference objects do not support "+thing)
}
func ZendWeakrefNoWrite(object *types.Zval, member *types.Zval, value *types.Zval, rtc *any) *types.Zval {
	ZendWeakrefUnsupported("properties")
	return EG__().GetUninitializedZval()
}
func ZendWeakrefNoRead(object *types.Zval, member *types.Zval, type_ int, rtc *any, rv *types.Zval) *types.Zval {
	if EG__().GetException() == nil {
		ZendWeakrefUnsupported("properties")
	}
	return EG__().GetUninitializedZval()
}
func ZendWeakrefNoReadPtr(object *types.Zval, member *types.Zval, type_ int, rtc *any) *types.Zval {
	ZendWeakrefUnsupported("property references")
	return nil
}
func ZendWeakrefNoIsset(object *types.Zval, member *types.Zval, hse int, rtc *any) int {
	if hse != 2 {
		ZendWeakrefUnsupported("properties")
	}
	return 0
}
func ZendWeakrefNoUnset(object *types.Zval, member *types.Zval, rtc *any) {
	ZendWeakrefUnsupported("properties")
}
func zim_WeakReference___construct(executeData *ZendExecuteData, return_value *types.Zval) {
	faults.ThrowError(nil, "Direct instantiation of 'WeakReference' is not allowed, "+"use WeakReference::create instead")
}
func zim_WeakReference_create(executeData *ZendExecuteData, return_value *types.Zval) {
	var referent *types.Zval
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
func zim_WeakReference_get(executeData *ZendExecuteData, return_value *types.Zval) {
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
	var ce types.ClassEntry
	memset(&ce, 0, b.SizeOf("zend_class_entry"))
	ce.SetName(types.NewString("WeakReference"))
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
