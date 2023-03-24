package zend

/**
 * ZendObjectHandlers
 */
type ZendObjectHandlers struct {
	offset               int
	free_obj             ZendObjectFreeObjT
	dtor_obj             ZendObjectDtorObjT
	clone_obj            ZendObjectCloneObjT
	read_property        ZendObjectReadPropertyT
	write_property       ZendObjectWritePropertyT
	read_dimension       ZendObjectReadDimensionT
	write_dimension      ZendObjectWriteDimensionT
	get_property_ptr_ptr ZendObjectGetPropertyPtrPtrT
	get                  ZendObjectGetT
	set                  ZendObjectSetT
	has_property         ZendObjectHasPropertyT
	unset_property       ZendObjectUnsetPropertyT
	has_dimension        ZendObjectHasDimensionT
	unset_dimension      ZendObjectUnsetDimensionT
	get_properties       ZendObjectGetPropertiesT
	get_method           ZendObjectGetMethodT
	call_method          ZendObjectCallMethodT
	get_constructor      ZendObjectGetConstructorT
	get_class_name       ZendObjectGetClassNameT
	compare_objects      ZendObjectCompareT
	cast_object          ZendObjectCastT
	count_elements       ZendObjectCountElementsT
	get_debug_info       ZendObjectGetDebugInfoT
	get_closure          ZendObjectGetClosureT
	get_gc               ZendObjectGetGcT
	do_operation         ZendObjectDoOperationT
	compare              ZendObjectCompareZvalsT
	get_properties_for   ZendObjectGetPropertiesForT
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
		offset:               offset,
		free_obj:             free_obj,
		dtor_obj:             dtor_obj,
		clone_obj:            clone_obj,
		read_property:        read_property,
		write_property:       write_property,
		read_dimension:       read_dimension,
		write_dimension:      write_dimension,
		get_property_ptr_ptr: get_property_ptr_ptr,
		get:                  get,
		set:                  set,
		has_property:         has_property,
		unset_property:       unset_property,
		has_dimension:        has_dimension,
		unset_dimension:      unset_dimension,
		get_properties:       get_properties,
		get_method:           get_method,
		call_method:          call_method,
		get_constructor:      get_constructor,
		get_class_name:       get_class_name,
		compare_objects:      compare_objects,
		cast_object:          cast_object,
		count_elements:       count_elements,
		get_debug_info:       get_debug_info,
		get_closure:          get_closure,
		get_gc:               get_gc,
		do_operation:         do_operation,
		compare:              compare,
		get_properties_for:   get_properties_for,
	}
}
func (this *ZendObjectHandlers) GetOffset() int                           { return this.offset }
func (this *ZendObjectHandlers) SetOffset(value int)                      { this.offset = value }
func (this *ZendObjectHandlers) GetFreeObj() ZendObjectFreeObjT           { return this.free_obj }
func (this *ZendObjectHandlers) SetFreeObj(value ZendObjectFreeObjT)      { this.free_obj = value }
func (this *ZendObjectHandlers) GetDtorObj() ZendObjectDtorObjT           { return this.dtor_obj }
func (this *ZendObjectHandlers) SetDtorObj(value ZendObjectDtorObjT)      { this.dtor_obj = value }
func (this *ZendObjectHandlers) GetCloneObj() ZendObjectCloneObjT         { return this.clone_obj }
func (this *ZendObjectHandlers) SetCloneObj(value ZendObjectCloneObjT)    { this.clone_obj = value }
func (this *ZendObjectHandlers) GetReadProperty() ZendObjectReadPropertyT { return this.read_property }
func (this *ZendObjectHandlers) SetReadProperty(value ZendObjectReadPropertyT) {
	this.read_property = value
}
func (this *ZendObjectHandlers) GetWriteProperty() ZendObjectWritePropertyT {
	return this.write_property
}
func (this *ZendObjectHandlers) SetWriteProperty(value ZendObjectWritePropertyT) {
	this.write_property = value
}
func (this *ZendObjectHandlers) GetReadDimension() ZendObjectReadDimensionT {
	return this.read_dimension
}
func (this *ZendObjectHandlers) SetReadDimension(value ZendObjectReadDimensionT) {
	this.read_dimension = value
}
func (this *ZendObjectHandlers) GetWriteDimension() ZendObjectWriteDimensionT {
	return this.write_dimension
}
func (this *ZendObjectHandlers) SetWriteDimension(value ZendObjectWriteDimensionT) {
	this.write_dimension = value
}
func (this *ZendObjectHandlers) GetGetPropertyPtrPtr() ZendObjectGetPropertyPtrPtrT {
	return this.get_property_ptr_ptr
}
func (this *ZendObjectHandlers) SetGetPropertyPtrPtr(value ZendObjectGetPropertyPtrPtrT) {
	this.get_property_ptr_ptr = value
}
func (this *ZendObjectHandlers) GetGet() ZendObjectGetT { return this.get }

// func (this *ZendObjectHandlers) SetGet(value ZendObjectGetT) { this.get = value }
func (this *ZendObjectHandlers) GetSet() ZendObjectSetT { return this.set }

// func (this *ZendObjectHandlers) SetSet(value ZendObjectSetT) { this.set = value }
func (this *ZendObjectHandlers) GetHasProperty() ZendObjectHasPropertyT { return this.has_property }
func (this *ZendObjectHandlers) SetHasProperty(value ZendObjectHasPropertyT) {
	this.has_property = value
}
func (this *ZendObjectHandlers) GetUnsetProperty() ZendObjectUnsetPropertyT {
	return this.unset_property
}
func (this *ZendObjectHandlers) SetUnsetProperty(value ZendObjectUnsetPropertyT) {
	this.unset_property = value
}
func (this *ZendObjectHandlers) GetHasDimension() ZendObjectHasDimensionT { return this.has_dimension }
func (this *ZendObjectHandlers) SetHasDimension(value ZendObjectHasDimensionT) {
	this.has_dimension = value
}
func (this *ZendObjectHandlers) GetUnsetDimension() ZendObjectUnsetDimensionT {
	return this.unset_dimension
}
func (this *ZendObjectHandlers) SetUnsetDimension(value ZendObjectUnsetDimensionT) {
	this.unset_dimension = value
}
func (this *ZendObjectHandlers) GetGetProperties() ZendObjectGetPropertiesT {
	return this.get_properties
}
func (this *ZendObjectHandlers) SetGetProperties(value ZendObjectGetPropertiesT) {
	this.get_properties = value
}
func (this *ZendObjectHandlers) GetGetMethod() ZendObjectGetMethodT      { return this.get_method }
func (this *ZendObjectHandlers) SetGetMethod(value ZendObjectGetMethodT) { this.get_method = value }
func (this *ZendObjectHandlers) GetCallMethod() ZendObjectCallMethodT    { return this.call_method }

// func (this *ZendObjectHandlers) SetCallMethod(value ZendObjectCallMethodT) { this.call_method = value }
func (this *ZendObjectHandlers) GetGetConstructor() ZendObjectGetConstructorT {
	return this.get_constructor
}
func (this *ZendObjectHandlers) SetGetConstructor(value ZendObjectGetConstructorT) {
	this.get_constructor = value
}
func (this *ZendObjectHandlers) GetGetClassName() ZendObjectGetClassNameT { return this.get_class_name }

// func (this *ZendObjectHandlers) SetGetClassName(value ZendObjectGetClassNameT) { this.get_class_name = value }
func (this *ZendObjectHandlers) GetCompareObjects() ZendObjectCompareT { return this.compare_objects }
func (this *ZendObjectHandlers) SetCompareObjects(value ZendObjectCompareT) {
	this.compare_objects = value
}
func (this *ZendObjectHandlers) GetCastObject() ZendObjectCastT      { return this.cast_object }
func (this *ZendObjectHandlers) SetCastObject(value ZendObjectCastT) { this.cast_object = value }
func (this *ZendObjectHandlers) GetCountElements() ZendObjectCountElementsT {
	return this.count_elements
}
func (this *ZendObjectHandlers) SetCountElements(value ZendObjectCountElementsT) {
	this.count_elements = value
}
func (this *ZendObjectHandlers) GetGetDebugInfo() ZendObjectGetDebugInfoT { return this.get_debug_info }
func (this *ZendObjectHandlers) SetGetDebugInfo(value ZendObjectGetDebugInfoT) {
	this.get_debug_info = value
}
func (this *ZendObjectHandlers) GetGetClosure() ZendObjectGetClosureT      { return this.get_closure }
func (this *ZendObjectHandlers) SetGetClosure(value ZendObjectGetClosureT) { this.get_closure = value }
func (this *ZendObjectHandlers) GetGetGc() ZendObjectGetGcT                { return this.get_gc }
func (this *ZendObjectHandlers) SetGetGc(value ZendObjectGetGcT)           { this.get_gc = value }
func (this *ZendObjectHandlers) GetDoOperation() ZendObjectDoOperationT    { return this.do_operation }

// func (this *ZendObjectHandlers) SetDoOperation(value ZendObjectDoOperationT) { this.do_operation = value }
func (this *ZendObjectHandlers) GetCompare() ZendObjectCompareZvalsT { return this.compare }

// func (this *ZendObjectHandlers) SetCompare(value ZendObjectCompareZvalsT) { this.compare = value }
func (this *ZendObjectHandlers) GetGetPropertiesFor() ZendObjectGetPropertiesForT {
	return this.get_properties_for
}
func (this *ZendObjectHandlers) SetGetPropertiesFor(value ZendObjectGetPropertiesForT) {
	this.get_properties_for = value
}
