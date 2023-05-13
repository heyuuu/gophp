package types

import "github.com/heyuuu/gophp/zend"

type ObjectHandlersSetting struct {
	Offset            int                     // 指向 Object 的偏移量
	FreeObj           zend.ZendObjectFreeObjT // todo free函数,在释放时若无free_obj则调用
	DtorObj           zend.ZendObjectDtorObjT // todo 析构函数,在释放时调用，优先级高于 freeObj
	CloneObj          zend.ZendObjectCloneObjT
	ReadProperty      zend.ZendObjectReadPropertyT
	WriteProperty     zend.ZendObjectWritePropertyT
	ReadDimension     zend.ZendObjectReadDimensionT
	WriteDimension    zend.ZendObjectWriteDimensionT
	GetPropertyPtrPtr zend.ZendObjectGetPropertyPtrPtrT
	Get               zend.ZendObjectGetT
	Set               zend.ZendObjectSetT
	HasProperty       zend.ZendObjectHasPropertyT
	UnsetProperty     zend.ZendObjectUnsetPropertyT
	HasDimension      zend.ZendObjectHasDimensionT
	UnsetDimension    zend.ZendObjectUnsetDimensionT
	GetProperties     zend.ZendObjectGetPropertiesT
	GetMethod         zend.ZendObjectGetMethodT
	CallMethod        zend.ZendObjectCallMethodT
	GetConstructor    zend.ZendObjectGetConstructorT
	GetClassName      zend.ZendObjectGetClassNameT
	CompareObjects    zend.ZendObjectCompareT
	CastObject        zend.ZendObjectCastT
	CountElements     zend.ZendObjectCountElementsT
	GetDebugInfo      zend.ZendObjectGetDebugInfoT
	GetClosure        zend.ZendObjectGetClosureT
	GetGc             zend.ZendObjectGetGcT
	DoOperation       zend.ZendObjectDoOperationT
	Compare           zend.ZendObjectCompareZvalsT
	GetPropertiesFor  zend.ZendObjectGetPropertiesForT
}

/**
 * ObjectHandlers
 */
type ObjectHandlers struct {
	offset            int                     // 指向 Object 的偏移量
	freeObj           zend.ZendObjectFreeObjT // todo free函数,在释放时若无free_obj则调用
	dtorObj           zend.ZendObjectDtorObjT // todo 析构函数,在释放时调用，优先级高于 freeObj
	cloneObj          zend.ZendObjectCloneObjT
	readProperty      zend.ZendObjectReadPropertyT
	writeProperty     zend.ZendObjectWritePropertyT
	readDimension     zend.ZendObjectReadDimensionT
	writeDimension    zend.ZendObjectWriteDimensionT
	getPropertyPtrPtr zend.ZendObjectGetPropertyPtrPtrT
	get               zend.ZendObjectGetT
	set               zend.ZendObjectSetT
	hasProperty       zend.ZendObjectHasPropertyT
	unsetProperty     zend.ZendObjectUnsetPropertyT
	hasDimension      zend.ZendObjectHasDimensionT
	unsetDimension    zend.ZendObjectUnsetDimensionT
	getProperties     zend.ZendObjectGetPropertiesT
	getMethod         zend.ZendObjectGetMethodT
	callMethod        zend.ZendObjectCallMethodT
	getConstructor    zend.ZendObjectGetConstructorT
	getClassName      zend.ZendObjectGetClassNameT
	compareObjects    zend.ZendObjectCompareT
	castObject        zend.ZendObjectCastT
	countElements     zend.ZendObjectCountElementsT
	getClosure        zend.ZendObjectGetClosureT
	doOperation       zend.ZendObjectDoOperationT
	compare           zend.ZendObjectCompareZvalsT
	getPropertiesFor  zend.ZendObjectGetPropertiesForT
}

func NewObjectHandlers(s ObjectHandlersSetting) *ObjectHandlers {
	return &ObjectHandlers{
		offset:            s.Offset,
		freeObj:           s.FreeObj,
		dtorObj:           s.DtorObj,
		cloneObj:          s.CloneObj,
		readProperty:      s.ReadProperty,
		writeProperty:     s.WriteProperty,
		readDimension:     s.ReadDimension,
		writeDimension:    s.WriteDimension,
		getPropertyPtrPtr: s.GetPropertyPtrPtr,
		get:               s.Get,
		set:               s.Set,
		hasProperty:       s.HasProperty,
		unsetProperty:     s.UnsetProperty,
		hasDimension:      s.HasDimension,
		unsetDimension:    s.UnsetDimension,
		getProperties:     s.GetProperties,
		getMethod:         s.GetMethod,
		callMethod:        s.CallMethod,
		getConstructor:    s.GetConstructor,
		getClassName:      s.GetClassName,
		compareObjects:    s.CompareObjects,
		castObject:        s.CastObject,
		countElements:     s.CountElements,
		getClosure:        s.GetClosure,
		doOperation:       s.DoOperation,
		compare:           s.Compare,
		getPropertiesFor:  s.GetPropertiesFor,
	}
}
func NewObjectHandlersEx(base *ObjectHandlers, s ObjectHandlersSetting) *ObjectHandlers {
	// todo settings 覆盖 base 产生新 handlers，后续用接口替换
	panic("todo")
}
