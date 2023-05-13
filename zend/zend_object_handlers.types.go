package zend

type ObjectHandlersSetting struct {
	Offset            int                // 指向 Object 的偏移量
	FreeObj           ZendObjectFreeObjT // todo free函数,在释放时若无free_obj则调用
	DtorObj           ZendObjectDtorObjT // todo 析构函数,在释放时调用，优先级高于 freeObj
	CloneObj          ZendObjectCloneObjT
	ReadProperty      ZendObjectReadPropertyT
	WriteProperty     ZendObjectWritePropertyT
	ReadDimension     ZendObjectReadDimensionT
	WriteDimension    ZendObjectWriteDimensionT
	GetPropertyPtrPtr ZendObjectGetPropertyPtrPtrT
	Get               ZendObjectGetT
	Set               ZendObjectSetT
	HasProperty       ZendObjectHasPropertyT
	UnsetProperty     ZendObjectUnsetPropertyT
	HasDimension      ZendObjectHasDimensionT
	UnsetDimension    ZendObjectUnsetDimensionT
	GetProperties     ZendObjectGetPropertiesT
	GetMethod         ZendObjectGetMethodT
	CallMethod        ZendObjectCallMethodT
	GetConstructor    ZendObjectGetConstructorT
	GetClassName      ZendObjectGetClassNameT
	CompareObjects    ZendObjectCompareT
	CastObject        ZendObjectCastT
	CountElements     ZendObjectCountElementsT
	GetDebugInfo      ZendObjectGetDebugInfoT
	GetClosure        ZendObjectGetClosureT
	GetGc             ZendObjectGetGcT
	DoOperation       ZendObjectDoOperationT
	Compare           ZendObjectCompareZvalsT
	GetPropertiesFor  ZendObjectGetPropertiesForT
}

/**
 * ObjectHandlers
 */
type ObjectHandlers struct {
	offset            int                // 指向 Object 的偏移量
	freeObj           ZendObjectFreeObjT // todo free函数,在释放时若无free_obj则调用
	dtorObj           ZendObjectDtorObjT // todo 析构函数,在释放时调用，优先级高于 freeObj
	cloneObj          ZendObjectCloneObjT
	readProperty      ZendObjectReadPropertyT
	writeProperty     ZendObjectWritePropertyT
	readDimension     ZendObjectReadDimensionT
	writeDimension    ZendObjectWriteDimensionT
	getPropertyPtrPtr ZendObjectGetPropertyPtrPtrT
	get               ZendObjectGetT
	set               ZendObjectSetT
	hasProperty       ZendObjectHasPropertyT
	unsetProperty     ZendObjectUnsetPropertyT
	hasDimension      ZendObjectHasDimensionT
	unsetDimension    ZendObjectUnsetDimensionT
	getProperties     ZendObjectGetPropertiesT
	getMethod         ZendObjectGetMethodT
	callMethod        ZendObjectCallMethodT
	getConstructor    ZendObjectGetConstructorT
	getClassName      ZendObjectGetClassNameT
	compareObjects    ZendObjectCompareT
	castObject        ZendObjectCastT
	countElements     ZendObjectCountElementsT
	getDebugInfo      ZendObjectGetDebugInfoT
	getClosure        ZendObjectGetClosureT
	doOperation       ZendObjectDoOperationT
	compare           ZendObjectCompareZvalsT
	getPropertiesFor  ZendObjectGetPropertiesForT
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

func (this *ObjectHandlers) GetFreeObj() ZendObjectFreeObjT               { return this.freeObj }
func (this *ObjectHandlers) GetDtorObj() ZendObjectDtorObjT               { return this.dtorObj }
func (this *ObjectHandlers) GetCloneObj() ZendObjectCloneObjT             { return this.cloneObj }
func (this *ObjectHandlers) GetReadProperty() ZendObjectReadPropertyT     { return this.readProperty }
func (this *ObjectHandlers) GetWriteProperty() ZendObjectWritePropertyT   { return this.writeProperty }
func (this *ObjectHandlers) GetReadDimension() ZendObjectReadDimensionT   { return this.readDimension }
func (this *ObjectHandlers) GetWriteDimension() ZendObjectWriteDimensionT { return this.writeDimension }
func (this *ObjectHandlers) GetGetPropertyPtrPtr() ZendObjectGetPropertyPtrPtrT {
	return this.getPropertyPtrPtr
}
func (this *ObjectHandlers) GetGet() ZendObjectGetT                       { return this.get }
func (this *ObjectHandlers) GetSet() ZendObjectSetT                       { return this.set }
func (this *ObjectHandlers) GetHasProperty() ZendObjectHasPropertyT       { return this.hasProperty }
func (this *ObjectHandlers) GetUnsetProperty() ZendObjectUnsetPropertyT   { return this.unsetProperty }
func (this *ObjectHandlers) GetHasDimension() ZendObjectHasDimensionT     { return this.hasDimension }
func (this *ObjectHandlers) GetUnsetDimension() ZendObjectUnsetDimensionT { return this.unsetDimension }
func (this *ObjectHandlers) GetGetProperties() ZendObjectGetPropertiesT   { return this.getProperties }
func (this *ObjectHandlers) GetGetMethod() ZendObjectGetMethodT           { return this.getMethod }
func (this *ObjectHandlers) GetCallMethod() ZendObjectCallMethodT         { return this.callMethod }
func (this *ObjectHandlers) GetGetConstructor() ZendObjectGetConstructorT { return this.getConstructor }
func (this *ObjectHandlers) GetGetClassName() ZendObjectGetClassNameT     { return this.getClassName }
func (this *ObjectHandlers) GetCompareObjects() ZendObjectCompareT        { return this.compareObjects }
func (this *ObjectHandlers) GetCastObject() ZendObjectCastT               { return this.castObject }
func (this *ObjectHandlers) GetCountElements() ZendObjectCountElementsT   { return this.countElements }
func (this *ObjectHandlers) GetGetDebugInfo() ZendObjectGetDebugInfoT     { return this.getDebugInfo }
func (this *ObjectHandlers) GetGetClosure() ZendObjectGetClosureT         { return this.getClosure }
func (this *ObjectHandlers) GetDoOperation() ZendObjectDoOperationT       { return this.doOperation }
func (this *ObjectHandlers) GetCompare() ZendObjectCompareZvalsT          { return this.compare }
func (this *ObjectHandlers) GetGetPropertiesFor() ZendObjectGetPropertiesForT {
	return this.getPropertiesFor
}
