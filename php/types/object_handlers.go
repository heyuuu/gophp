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
	getDebugInfo      zend.ZendObjectGetDebugInfoT
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
		getDebugInfo:      s.GetDebugInfo,
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

func (this *ObjectHandlers) GetCloneObj() zend.ZendObjectCloneObjT         { return this.cloneObj }
func (this *ObjectHandlers) GetReadProperty() zend.ZendObjectReadPropertyT { return this.readProperty }
func (this *ObjectHandlers) GetWriteProperty() zend.ZendObjectWritePropertyT {
	return this.writeProperty
}
func (this *ObjectHandlers) GetReadDimension() zend.ZendObjectReadDimensionT {
	return this.readDimension
}
func (this *ObjectHandlers) GetWriteDimension() zend.ZendObjectWriteDimensionT {
	return this.writeDimension
}
func (this *ObjectHandlers) GetGetPropertyPtrPtr() zend.ZendObjectGetPropertyPtrPtrT {
	return this.getPropertyPtrPtr
}
func (this *ObjectHandlers) GetGet() zend.ZendObjectGetT                 { return this.get }
func (this *ObjectHandlers) GetSet() zend.ZendObjectSetT                 { return this.set }
func (this *ObjectHandlers) GetHasProperty() zend.ZendObjectHasPropertyT { return this.hasProperty }
func (this *ObjectHandlers) GetUnsetProperty() zend.ZendObjectUnsetPropertyT {
	return this.unsetProperty
}
func (this *ObjectHandlers) GetHasDimension() zend.ZendObjectHasDimensionT { return this.hasDimension }
func (this *ObjectHandlers) GetUnsetDimension() zend.ZendObjectUnsetDimensionT {
	return this.unsetDimension
}
func (this *ObjectHandlers) GetGetProperties() zend.ZendObjectGetPropertiesT {
	return this.getProperties
}
func (this *ObjectHandlers) GetGetMethod() zend.ZendObjectGetMethodT   { return this.getMethod }
func (this *ObjectHandlers) GetCallMethod() zend.ZendObjectCallMethodT { return this.callMethod }
func (this *ObjectHandlers) GetGetConstructor() zend.ZendObjectGetConstructorT {
	return this.getConstructor
}
func (this *ObjectHandlers) GetGetClassName() zend.ZendObjectGetClassNameT { return this.getClassName }
func (this *ObjectHandlers) GetCompareObjects() zend.ZendObjectCompareT    { return this.compareObjects }
func (this *ObjectHandlers) GetCastObject() zend.ZendObjectCastT           { return this.castObject }
func (this *ObjectHandlers) GetCountElements() zend.ZendObjectCountElementsT {
	return this.countElements
}
func (this *ObjectHandlers) GetGetDebugInfo() zend.ZendObjectGetDebugInfoT { return this.getDebugInfo }
func (this *ObjectHandlers) GetGetClosure() zend.ZendObjectGetClosureT     { return this.getClosure }
func (this *ObjectHandlers) GetDoOperation() zend.ZendObjectDoOperationT   { return this.doOperation }
func (this *ObjectHandlers) GetCompare() zend.ZendObjectCompareZvalsT      { return this.compare }
func (this *ObjectHandlers) GetGetPropertiesFor() zend.ZendObjectGetPropertiesForT {
	return this.getPropertiesFor
}
