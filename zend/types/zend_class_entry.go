package types

import "sik/zend"

/**
 * ClassEntry
 */
type ClassEntry struct {
	type_ byte
	name  string // *String
	__0   struct /* union */ {
		parent      *ClassEntry
		parent_name *String
	}
	refcount                     int
	ce_flags                     uint32
	default_properties_count     int
	default_static_members_count int
	default_properties_table     *Zval
	default_static_members_table *Zval
	static_members_table__ptr    **Zval
	function_table               Array
	properties_info              Array
	constants_table              Array
	properties_info_table        **zend.ZendPropertyInfo
	constructor                  *zend.ZendFunction
	destructor                   *zend.ZendFunction
	clone                        *zend.ZendFunction
	__get                        *zend.ZendFunction
	__set                        *zend.ZendFunction
	__unset                      *zend.ZendFunction
	__isset                      *zend.ZendFunction
	__call                       *zend.ZendFunction
	__callstatic                 *zend.ZendFunction
	__tostring                   *zend.ZendFunction
	__debugInfo                  *zend.ZendFunction
	serialize_func               *zend.ZendFunction
	unserialize_func             *zend.ZendFunction
	iterator_funcs_ptr           *zend.ZendClassIteratorFuncs
	__1                          struct /* union */ {
		create_object              func(class_type *ClassEntry) *ZendObject
		interface_gets_implemented func(iface *ClassEntry, class_type *ClassEntry) int
	}
	get_iterator      func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator
	get_static_method func(ce *ClassEntry, method *String) *zend.ZendFunction
	serialize         func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int
	unserialize       func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int
	num_interfaces    uint32
	num_traits        uint32
	__2               struct /* union */ {
		interfaces      **ClassEntry
		interface_names *zend.ZendClassName
	}
	trait_names       *zend.ZendClassName
	trait_aliases     **zend.ZendTraitAlias
	trait_precedences **zend.ZendTraitPrecedence
	info              struct /* union */ {
		user struct {
			filename    *String
			line_start  uint32
			line_end    uint32
			doc_comment *String
		}
		internal struct {
			builtin_functions []ZendFunctionEntry
			module            *zend.ZendModuleEntry
		}
	}
}

func (this *ClassEntry) InitMethods(functions []ZendFunctionEntry) {
	this.SetConstructor(nil)
	this.SetDestructor(nil)
	this.SetClone(nil)
	this.SetSerialize(nil)
	this.SetUnserialize(nil)
	this.SetCreateObject(nil)
	this.SetGetStaticMethod(nil)
	this.SetCall(nil)
	this.SetCallstatic(nil)
	this.SetTostring(nil)
	this.SetGet(nil)
	this.SetSet(nil)
	this.SetUnset(nil)
	this.SetIsset(nil)
	this.SetDebugInfo(nil)
	this.SetSerializeFunc(nil)
	this.SetUnserializeFunc(nil)
	this.SetParent(nil)
	this.SetNumInterfaces(0)
	this.SetTraitNames(nil)
	this.SetNumTraits(0)
	this.SetTraitAliases(nil)
	this.SetTraitPrecedences(nil)
	this.SetInterfaces(nil)
	this.SetGetIterator(nil)
	this.SetIteratorFuncsPtr(nil)
	this.SetModule(nil)
	this.SetBuiltinFunctions(functions)
}

func NewClassEntry(name string, functions []ZendFunctionEntry) *ClassEntry {
	class := &ClassEntry{
		name: name,
	}
	class.info.internal.builtin_functions = functions
	return class
}

func (this *ClassEntry) Name() string { return this.name }

/**
 * Getter / Setter
 */
func (this *ClassEntry) GetName() *String      { return NewString(this.name) }
func (this *ClassEntry) SetName(value *String) { this.name = value.GetStr() }

func (this *ClassEntry) GetType() byte                  { return this.type_ }
func (this *ClassEntry) SetType(value byte)             { this.type_ = value }
func (this *ClassEntry) GetParent() *ClassEntry         { return this.__0.parent }
func (this *ClassEntry) SetParent(value *ClassEntry)    { this.__0.parent = value }
func (this *ClassEntry) GetParentName() *String         { return this.__0.parent_name }
func (this *ClassEntry) SetParentName(value *String)    { this.__0.parent_name = value }
func (this *ClassEntry) GetRefcount() int               { return this.refcount }
func (this *ClassEntry) SetRefcount(value int)          { this.refcount = value }
func (this *ClassEntry) GetCeFlags() uint32             { return this.ce_flags }
func (this *ClassEntry) SetCeFlags(value uint32)        { this.ce_flags = value }
func (this *ClassEntry) GetDefaultPropertiesCount() int { return this.default_properties_count }
func (this *ClassEntry) SetDefaultPropertiesCount(value int) {
	this.default_properties_count = value
}
func (this *ClassEntry) GetDefaultStaticMembersCount() int {
	return this.default_static_members_count
}
func (this *ClassEntry) SetDefaultStaticMembersCount(value int) {
	this.default_static_members_count = value
}
func (this *ClassEntry) GetDefaultPropertiesTable() *Zval {
	return this.default_properties_table
}
func (this *ClassEntry) SetDefaultPropertiesTable(value *Zval) {
	this.default_properties_table = value
}
func (this *ClassEntry) GetDefaultStaticMembersTable() *Zval {
	return this.default_static_members_table
}
func (this *ClassEntry) SetDefaultStaticMembersTable(value *Zval) {
	this.default_static_members_table = value
}
func (this *ClassEntry) GetStaticMembersTablePtr() **Zval {
	return this.static_members_table__ptr
}
func (this *ClassEntry) GetFunctionTable() Array  { return this.function_table }
func (this *ClassEntry) GetPropertiesInfo() Array { return this.properties_info }
func (this *ClassEntry) GetConstantsTable() Array { return this.constants_table }
func (this *ClassEntry) GetPropertiesInfoTable() **zend.ZendPropertyInfo {
	return this.properties_info_table
}
func (this *ClassEntry) SetPropertiesInfoTable(value **zend.ZendPropertyInfo) {
	this.properties_info_table = value
}
func (this *ClassEntry) GetConstructor() *zend.ZendFunction        { return this.constructor }
func (this *ClassEntry) SetConstructor(value *zend.ZendFunction)   { this.constructor = value }
func (this *ClassEntry) GetDestructor() *zend.ZendFunction         { return this.destructor }
func (this *ClassEntry) SetDestructor(value *zend.ZendFunction)    { this.destructor = value }
func (this *ClassEntry) GetClone() *zend.ZendFunction              { return this.clone }
func (this *ClassEntry) SetClone(value *zend.ZendFunction)         { this.clone = value }
func (this *ClassEntry) GetGet() *zend.ZendFunction                { return this.__get }
func (this *ClassEntry) SetGet(value *zend.ZendFunction)           { this.__get = value }
func (this *ClassEntry) GetSet() *zend.ZendFunction                { return this.__set }
func (this *ClassEntry) SetSet(value *zend.ZendFunction)           { this.__set = value }
func (this *ClassEntry) GetUnset() *zend.ZendFunction              { return this.__unset }
func (this *ClassEntry) SetUnset(value *zend.ZendFunction)         { this.__unset = value }
func (this *ClassEntry) GetIsset() *zend.ZendFunction              { return this.__isset }
func (this *ClassEntry) SetIsset(value *zend.ZendFunction)         { this.__isset = value }
func (this *ClassEntry) GetCall() *zend.ZendFunction               { return this.__call }
func (this *ClassEntry) SetCall(value *zend.ZendFunction)          { this.__call = value }
func (this *ClassEntry) GetCallstatic() *zend.ZendFunction         { return this.__callstatic }
func (this *ClassEntry) SetCallstatic(value *zend.ZendFunction)    { this.__callstatic = value }
func (this *ClassEntry) GetTostring() *zend.ZendFunction           { return this.__tostring }
func (this *ClassEntry) SetTostring(value *zend.ZendFunction)      { this.__tostring = value }
func (this *ClassEntry) GetDebugInfo() *zend.ZendFunction          { return this.__debugInfo }
func (this *ClassEntry) SetDebugInfo(value *zend.ZendFunction)     { this.__debugInfo = value }
func (this *ClassEntry) GetSerializeFunc() *zend.ZendFunction      { return this.serialize_func }
func (this *ClassEntry) SetSerializeFunc(value *zend.ZendFunction) { this.serialize_func = value }
func (this *ClassEntry) GetUnserializeFunc() *zend.ZendFunction    { return this.unserialize_func }
func (this *ClassEntry) SetUnserializeFunc(value *zend.ZendFunction) {
	this.unserialize_func = value
}
func (this *ClassEntry) GetIteratorFuncsPtr() *zend.ZendClassIteratorFuncs {
	return this.iterator_funcs_ptr
}
func (this *ClassEntry) SetIteratorFuncsPtr(value *zend.ZendClassIteratorFuncs) {
	this.iterator_funcs_ptr = value
}
func (this *ClassEntry) GetCreateObject() func(class_type *ClassEntry) *ZendObject {
	return this.__1.create_object
}
func (this *ClassEntry) SetCreateObject(value func(class_type *ClassEntry) *ZendObject) {
	this.__1.create_object = value
}
func (this *ClassEntry) GetInterfaceGetsImplemented() func(iface *ClassEntry, class_type *ClassEntry) int {
	return this.__1.interface_gets_implemented
}
func (this *ClassEntry) SetInterfaceGetsImplemented(value func(iface *ClassEntry, class_type *ClassEntry) int) {
	this.__1.interface_gets_implemented = value
}
func (this *ClassEntry) GetGetIterator() func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator {
	return this.get_iterator
}
func (this *ClassEntry) SetGetIterator(value func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator) {
	this.get_iterator = value
}
func (this *ClassEntry) GetGetStaticMethod() func(ce *ClassEntry, method *String) *zend.ZendFunction {
	return this.get_static_method
}
func (this *ClassEntry) SetGetStaticMethod(value func(ce *ClassEntry, method *String) *zend.ZendFunction) {
	this.get_static_method = value
}
func (this *ClassEntry) GetSerialize() func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int {
	return this.serialize
}
func (this *ClassEntry) SetSerialize(value func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int) {
	this.serialize = value
}
func (this *ClassEntry) GetUnserialize() func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int {
	return this.unserialize
}
func (this *ClassEntry) SetUnserialize(value func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int) {
	this.unserialize = value
}
func (this *ClassEntry) GetNumInterfaces() uint32               { return this.num_interfaces }
func (this *ClassEntry) SetNumInterfaces(value uint32)          { this.num_interfaces = value }
func (this *ClassEntry) GetNumTraits() uint32                   { return this.num_traits }
func (this *ClassEntry) SetNumTraits(value uint32)              { this.num_traits = value }
func (this *ClassEntry) GetInterfaces() **ClassEntry            { return this.__2.interfaces }
func (this *ClassEntry) SetInterfaces(value **ClassEntry)       { this.__2.interfaces = value }
func (this *ClassEntry) GetInterfaceNames() *zend.ZendClassName { return this.__2.interface_names }
func (this *ClassEntry) SetInterfaceNames(value *zend.ZendClassName) {
	this.__2.interface_names = value
}
func (this *ClassEntry) GetTraitNames() *zend.ZendClassName          { return this.trait_names }
func (this *ClassEntry) SetTraitNames(value *zend.ZendClassName)     { this.trait_names = value }
func (this *ClassEntry) GetTraitAliases() **zend.ZendTraitAlias      { return this.trait_aliases }
func (this *ClassEntry) SetTraitAliases(value **zend.ZendTraitAlias) { this.trait_aliases = value }
func (this *ClassEntry) GetTraitPrecedences() **zend.ZendTraitPrecedence {
	return this.trait_precedences
}
func (this *ClassEntry) SetTraitPrecedences(value **zend.ZendTraitPrecedence) {
	this.trait_precedences = value
}
func (this *ClassEntry) GetFilename() *String      { return this.info.user.filename }
func (this *ClassEntry) SetFilename(value *String) { this.info.user.filename = value }

// func (this *ClassEntry)  GetLineStart() uint32      { return this.info.user.line_start }
func (this *ClassEntry) SetLineStart(value uint32) { this.info.user.line_start = value }

// func (this *ClassEntry)  GetLineEnd() uint32      { return this.info.user.line_end }
func (this *ClassEntry) SetLineEnd(value uint32) { this.info.user.line_end = value }
func (this *ClassEntry) GetDocComment() *String  { return this.info.user.doc_comment }
func (this *ClassEntry) SetDocComment(value *String) {
	this.info.user.doc_comment = value
}
func (this *ClassEntry) GetBuiltinFunctions() *ZendFunctionEntry {
	return this.info.internal.builtin_functions
}
func (this *ClassEntry) SetBuiltinFunctions(value *ZendFunctionEntry) {
	this.info.internal.builtin_functions = value
}
func (this *ClassEntry) GetModule() *zend.ZendModuleEntry      { return this.info.internal.module }
func (this *ClassEntry) SetModule(value *zend.ZendModuleEntry) { this.info.internal.module = value }

/* ClassEntry.ce_flags */
func (this *ClassEntry) AddCeFlags(value uint32)      { this.ce_flags |= value }
func (this *ClassEntry) SubCeFlags(value uint32)      { this.ce_flags &^= value }
func (this *ClassEntry) HasCeFlags(value uint32) bool { return this.ce_flags&value != 0 }
func (this *ClassEntry) SwitchCeFlags(value uint32, cond bool) {
	if cond {
		this.AddCeFlags(value)
	} else {
		this.SubCeFlags(value)
	}
}
func (this ClassEntry) IsConstantsUpdated() bool {
	return this.HasCeFlags(zend.ZEND_ACC_CONSTANTS_UPDATED)
}
func (this ClassEntry) IsInterface() bool { return this.HasCeFlags(zend.ZEND_ACC_INTERFACE) }
func (this ClassEntry) IsTrait() bool     { return this.HasCeFlags(zend.ZEND_ACC_TRAIT) }
func (this ClassEntry) IsImmutable() bool { return this.HasCeFlags(zend.ZEND_ACC_IMMUTABLE) }
func (this ClassEntry) IsResolvedParent() bool {
	return this.HasCeFlags(zend.ZEND_ACC_RESOLVED_PARENT)
}
func (this ClassEntry) IsLinked() bool { return this.HasCeFlags(zend.ZEND_ACC_LINKED) }
func (this ClassEntry) IsImplementTraits() bool {
	return this.HasCeFlags(zend.ZEND_ACC_IMPLEMENT_TRAITS)
}
func (this ClassEntry) IsHasStaticInMethods() bool {
	return this.HasCeFlags(zend.ZEND_HAS_STATIC_IN_METHODS)
}
func (this ClassEntry) IsNearlyLinked() bool { return this.HasCeFlags(zend.ZEND_ACC_NEARLY_LINKED) }
func (this ClassEntry) IsResolvedInterfaces() bool {
	return this.HasCeFlags(zend.ZEND_ACC_RESOLVED_INTERFACES)
}
func (this ClassEntry) IsFinal() bool { return this.HasCeFlags(zend.ZEND_ACC_FINAL) }
func (this ClassEntry) IsImplementInterfaces() bool {
	return this.HasCeFlags(zend.ZEND_ACC_IMPLEMENT_INTERFACES)
}
func (this ClassEntry) IsImplicitAbstractClass() bool {
	return this.HasCeFlags(zend.ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
}
func (this ClassEntry) IsUnresolvedVariance() bool {
	return this.HasCeFlags(zend.ZEND_ACC_UNRESOLVED_VARIANCE)
}
func (this ClassEntry) IsHasUnlinkedUses() bool {
	return this.HasCeFlags(zend.ZEND_ACC_HAS_UNLINKED_USES)
}
func (this ClassEntry) IsUseGuards() bool { return this.HasCeFlags(zend.ZEND_ACC_USE_GUARDS) }
func (this ClassEntry) IsPropertyTypesResolved() bool {
	return this.HasCeFlags(zend.ZEND_ACC_PROPERTY_TYPES_RESOLVED)
}
func (this ClassEntry) IsExplicitAbstractClass() bool {
	return this.HasCeFlags(zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)
}
func (this ClassEntry) IsHasTypeHints() bool {
	return this.HasCeFlags(zend.ZEND_ACC_HAS_TYPE_HINTS)
}
func (this ClassEntry) IsPreloaded() bool { return this.HasCeFlags(zend.ZEND_ACC_PRELOADED) }
func (this ClassEntry) IsInherited() bool { return this.HasCeFlags(zend.ZEND_ACC_INHERITED) }
func (this ClassEntry) IsTopLevel() bool  { return this.HasCeFlags(zend.ZEND_ACC_TOP_LEVEL) }
func (this ClassEntry) IsReuseGetIterator() bool {
	return this.HasCeFlags(zend.ZEND_ACC_REUSE_GET_ITERATOR)
}
func (this *ClassEntry) SetIsConstantsUpdated(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_CONSTANTS_UPDATED, cond)
}
func (this *ClassEntry) SetIsInterface(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_INTERFACE, cond)
}
func (this *ClassEntry) SetIsTrait(cond bool) { this.SwitchCeFlags(zend.ZEND_ACC_TRAIT, cond) }
func (this *ClassEntry) SetIsImmutable(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_IMMUTABLE, cond)
}
func (this *ClassEntry) SetIsResolvedParent(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_RESOLVED_PARENT, cond)
}
func (this *ClassEntry) SetIsLinked(cond bool) { this.SwitchCeFlags(zend.ZEND_ACC_LINKED, cond) }
func (this *ClassEntry) SetIsImplementTraits(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_IMPLEMENT_TRAITS, cond)
}
func (this *ClassEntry) SetIsHasStaticInMethods(cond bool) {
	this.SwitchCeFlags(zend.ZEND_HAS_STATIC_IN_METHODS, cond)
}
func (this *ClassEntry) SetIsNearlyLinked(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_NEARLY_LINKED, cond)
}
func (this *ClassEntry) SetIsResolvedInterfaces(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_RESOLVED_INTERFACES, cond)
}
func (this *ClassEntry) SetIsFinal(cond bool) { this.SwitchCeFlags(zend.ZEND_ACC_FINAL, cond) }
func (this *ClassEntry) SetIsImplementInterfaces(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_IMPLEMENT_INTERFACES, cond)
}
func (this *ClassEntry) SetIsImplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_IMPLICIT_ABSTRACT_CLASS, cond)
}
func (this *ClassEntry) SetIsUnresolvedVariance(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_UNRESOLVED_VARIANCE, cond)
}
func (this *ClassEntry) SetIsHasUnlinkedUses(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_HAS_UNLINKED_USES, cond)
}
func (this *ClassEntry) SetIsUseGuards(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_USE_GUARDS, cond)
}
func (this *ClassEntry) SetIsPropertyTypesResolved(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_PROPERTY_TYPES_RESOLVED, cond)
}
func (this *ClassEntry) SetIsExplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS, cond)
}
func (this *ClassEntry) SetIsHasTypeHints(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_HAS_TYPE_HINTS, cond)
}
func (this *ClassEntry) SetIsPreloaded(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_PRELOADED, cond)
}
func (this *ClassEntry) SetIsInherited(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_INHERITED, cond)
}
func (this *ClassEntry) SetIsTopLevel(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_TOP_LEVEL, cond)
}
func (this *ClassEntry) SetIsReuseGetIterator(cond bool) {
	this.SwitchCeFlags(zend.ZEND_ACC_REUSE_GET_ITERATOR, cond)
}
