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
 * ZendObjectHandlers
 */
type ZendObjectHandlers struct {
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
	getGc             ZendObjectGetGcT
	doOperation       ZendObjectDoOperationT
	compare           ZendObjectCompareZvalsT
	getPropertiesFor  ZendObjectGetPropertiesForT
}

func NewZendObjectHandlers(s ObjectHandlersSetting) *ZendObjectHandlers {
	return &ZendObjectHandlers{
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
		getGc:             s.GetGc,
		doOperation:       s.DoOperation,
		compare:           s.Compare,
		getPropertiesFor:  s.GetPropertiesFor,
	}
}

func MakeZendObjectHandlers(
	offset int,
	free_obj ZendObjectFreeObjT,
	dtor_obj ZendObjectDtorObjT,
	clone_obj ZendObjectCloneObjT,
	read_property ZendObjectReadPropertyT,
	write_property ZendObjectWritePropertyT,
	read_dimension ZendObjectReadDimensionT,
	write_dimension ZendObjectWriteDimensionT,
	get_property_ptr_ptr ZendObjectGetPropertyPtrPtrT,
	get ZendObjectGetT,
	set ZendObjectSetT,
	has_property ZendObjectHasPropertyT,
	unset_property ZendObjectUnsetPropertyT,
	has_dimension ZendObjectHasDimensionT,
	unset_dimension ZendObjectUnsetDimensionT,
	get_properties ZendObjectGetPropertiesT,
	get_method ZendObjectGetMethodT,
	call_method ZendObjectCallMethodT,
	get_constructor ZendObjectGetConstructorT,
	get_class_name ZendObjectGetClassNameT,
	compare_objects ZendObjectCompareT,
	cast_object ZendObjectCastT,
	count_elements ZendObjectCountElementsT,
	get_debug_info ZendObjectGetDebugInfoT,
	get_closure ZendObjectGetClosureT,
	get_gc ZendObjectGetGcT,
	do_operation ZendObjectDoOperationT,
	compare ZendObjectCompareZvalsT,
	get_properties_for ZendObjectGetPropertiesForT,
) ZendObjectHandlers {
	return ZendObjectHandlers{
		offset:            offset,
		freeObj:           free_obj,
		dtorObj:           dtor_obj,
		cloneObj:          clone_obj,
		readProperty:      read_property,
		writeProperty:     write_property,
		readDimension:     read_dimension,
		writeDimension:    write_dimension,
		getPropertyPtrPtr: get_property_ptr_ptr,
		get:               get,
		set:               set,
		hasProperty:       has_property,
		unsetProperty:     unset_property,
		hasDimension:      has_dimension,
		unsetDimension:    unset_dimension,
		getProperties:     get_properties,
		getMethod:         get_method,
		callMethod:        call_method,
		getConstructor:    get_constructor,
		getClassName:      get_class_name,
		compareObjects:    compare_objects,
		castObject:        cast_object,
		countElements:     count_elements,
		getDebugInfo:      get_debug_info,
		getClosure:        get_closure,
		getGc:             get_gc,
		doOperation:       do_operation,
		compare:           compare,
		getPropertiesFor:  get_properties_for,
	}
}
func (this *ZendObjectHandlers) GetOffset() int                           { return this.offset }
func (this *ZendObjectHandlers) SetOffset(value int)                      { this.offset = value }
func (this *ZendObjectHandlers) GetFreeObj() ZendObjectFreeObjT           { return this.freeObj }
func (this *ZendObjectHandlers) SetFreeObj(value ZendObjectFreeObjT)      { this.freeObj = value }
func (this *ZendObjectHandlers) GetDtorObj() ZendObjectDtorObjT           { return this.dtorObj }
func (this *ZendObjectHandlers) SetDtorObj(value ZendObjectDtorObjT)      { this.dtorObj = value }
func (this *ZendObjectHandlers) GetCloneObj() ZendObjectCloneObjT         { return this.cloneObj }
func (this *ZendObjectHandlers) SetCloneObj(value ZendObjectCloneObjT)    { this.cloneObj = value }
func (this *ZendObjectHandlers) GetReadProperty() ZendObjectReadPropertyT { return this.readProperty }
func (this *ZendObjectHandlers) SetReadProperty(value ZendObjectReadPropertyT) {
	this.readProperty = value
}
func (this *ZendObjectHandlers) GetWriteProperty() ZendObjectWritePropertyT {
	return this.writeProperty
}
func (this *ZendObjectHandlers) SetWriteProperty(value ZendObjectWritePropertyT) {
	this.writeProperty = value
}
func (this *ZendObjectHandlers) GetReadDimension() ZendObjectReadDimensionT {
	return this.readDimension
}
func (this *ZendObjectHandlers) SetReadDimension(value ZendObjectReadDimensionT) {
	this.readDimension = value
}
func (this *ZendObjectHandlers) GetWriteDimension() ZendObjectWriteDimensionT {
	return this.writeDimension
}
func (this *ZendObjectHandlers) SetWriteDimension(value ZendObjectWriteDimensionT) {
	this.writeDimension = value
}
func (this *ZendObjectHandlers) GetGetPropertyPtrPtr() ZendObjectGetPropertyPtrPtrT {
	return this.getPropertyPtrPtr
}
func (this *ZendObjectHandlers) SetGetPropertyPtrPtr(value ZendObjectGetPropertyPtrPtrT) {
	this.getPropertyPtrPtr = value
}
func (this *ZendObjectHandlers) GetGet() ZendObjectGetT { return this.get }

// func (this *ZendObjectHandlers) SetGet(value ZendObjectGetT) { this.get = value }
func (this *ZendObjectHandlers) GetSet() ZendObjectSetT { return this.set }

// func (this *ZendObjectHandlers) SetSet(value ZendObjectSetT) { this.set = value }
func (this *ZendObjectHandlers) GetHasProperty() ZendObjectHasPropertyT { return this.hasProperty }
func (this *ZendObjectHandlers) SetHasProperty(value ZendObjectHasPropertyT) {
	this.hasProperty = value
}
func (this *ZendObjectHandlers) GetUnsetProperty() ZendObjectUnsetPropertyT {
	return this.unsetProperty
}
func (this *ZendObjectHandlers) SetUnsetProperty(value ZendObjectUnsetPropertyT) {
	this.unsetProperty = value
}
func (this *ZendObjectHandlers) GetHasDimension() ZendObjectHasDimensionT { return this.hasDimension }
func (this *ZendObjectHandlers) SetHasDimension(value ZendObjectHasDimensionT) {
	this.hasDimension = value
}
func (this *ZendObjectHandlers) GetUnsetDimension() ZendObjectUnsetDimensionT {
	return this.unsetDimension
}
func (this *ZendObjectHandlers) SetUnsetDimension(value ZendObjectUnsetDimensionT) {
	this.unsetDimension = value
}
func (this *ZendObjectHandlers) GetGetProperties() ZendObjectGetPropertiesT {
	return this.getProperties
}
func (this *ZendObjectHandlers) SetGetProperties(value ZendObjectGetPropertiesT) {
	this.getProperties = value
}
func (this *ZendObjectHandlers) GetGetMethod() ZendObjectGetMethodT      { return this.getMethod }
func (this *ZendObjectHandlers) SetGetMethod(value ZendObjectGetMethodT) { this.getMethod = value }
func (this *ZendObjectHandlers) GetCallMethod() ZendObjectCallMethodT    { return this.callMethod }

// func (this *ZendObjectHandlers) SetCallMethod(value ZendObjectCallMethodT) { this.call_method = value }
func (this *ZendObjectHandlers) GetGetConstructor() ZendObjectGetConstructorT {
	return this.getConstructor
}
func (this *ZendObjectHandlers) SetGetConstructor(value ZendObjectGetConstructorT) {
	this.getConstructor = value
}
func (this *ZendObjectHandlers) GetGetClassName() ZendObjectGetClassNameT { return this.getClassName }

// func (this *ZendObjectHandlers) SetGetClassName(value ZendObjectGetClassNameT) { this.get_class_name = value }
func (this *ZendObjectHandlers) GetCompareObjects() ZendObjectCompareT { return this.compareObjects }
func (this *ZendObjectHandlers) SetCompareObjects(value ZendObjectCompareT) {
	this.compareObjects = value
}
func (this *ZendObjectHandlers) GetCastObject() ZendObjectCastT      { return this.castObject }
func (this *ZendObjectHandlers) SetCastObject(value ZendObjectCastT) { this.castObject = value }
func (this *ZendObjectHandlers) GetCountElements() ZendObjectCountElementsT {
	return this.countElements
}
func (this *ZendObjectHandlers) SetCountElements(value ZendObjectCountElementsT) {
	this.countElements = value
}
func (this *ZendObjectHandlers) GetGetDebugInfo() ZendObjectGetDebugInfoT { return this.getDebugInfo }
func (this *ZendObjectHandlers) SetGetDebugInfo(value ZendObjectGetDebugInfoT) {
	this.getDebugInfo = value
}
func (this *ZendObjectHandlers) GetGetClosure() ZendObjectGetClosureT      { return this.getClosure }
func (this *ZendObjectHandlers) SetGetClosure(value ZendObjectGetClosureT) { this.getClosure = value }
func (this *ZendObjectHandlers) GetGetGc() ZendObjectGetGcT                { return this.getGc }
func (this *ZendObjectHandlers) SetGetGc(value ZendObjectGetGcT)           { this.getGc = value }
func (this *ZendObjectHandlers) GetDoOperation() ZendObjectDoOperationT    { return this.doOperation }

// func (this *ZendObjectHandlers) SetDoOperation(value ZendObjectDoOperationT) { this.do_operation = value }
func (this *ZendObjectHandlers) GetCompare() ZendObjectCompareZvalsT { return this.compare }

// func (this *ZendObjectHandlers) SetCompare(value ZendObjectCompareZvalsT) { this.compare = value }
func (this *ZendObjectHandlers) GetGetPropertiesFor() ZendObjectGetPropertiesForT {
	return this.getPropertiesFor
}
func (this *ZendObjectHandlers) SetGetPropertiesFor(value ZendObjectGetPropertiesForT) {
	this.getPropertiesFor = value
}
