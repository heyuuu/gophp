// <<generate>>

package zend

/**
 * ZendClassEntry
 */
type ZendClassEntry struct {
	type_ byte
	name  *ZendString
	__0   struct /* union */ {
		parent      *ZendClassEntry
		parent_name *ZendString
	}
	refcount                     int
	ce_flags                     uint32
	default_properties_count     int
	default_static_members_count int
	default_properties_table     *Zval
	default_static_members_table *Zval
	static_members_table__ptr    **Zval
	function_table               HashTable
	properties_info              HashTable
	constants_table              HashTable
	properties_info_table        **ZendPropertyInfo
	constructor                  *ZendFunction
	destructor                   *ZendFunction
	clone                        *ZendFunction
	__get                        *ZendFunction
	__set                        *ZendFunction
	__unset                      *ZendFunction
	__isset                      *ZendFunction
	__call                       *ZendFunction
	__callstatic                 *ZendFunction
	__tostring                   *ZendFunction
	__debugInfo                  *ZendFunction
	serialize_func               *ZendFunction
	unserialize_func             *ZendFunction
	iterator_funcs_ptr           *ZendClassIteratorFuncs
	__1                          struct /* union */ {
		create_object              func(class_type *ZendClassEntry) *ZendObject
		interface_gets_implemented func(iface *ZendClassEntry, class_type *ZendClassEntry) int
	}
	get_iterator      func(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator
	get_static_method func(ce *ZendClassEntry, method *ZendString) *ZendFunction
	serialize         func(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int
	unserialize       func(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int
	num_interfaces    uint32
	num_traits        uint32
	__2               struct /* union */ {
		interfaces      **ZendClassEntry
		interface_names *ZendClassName
	}
	trait_names       *ZendClassName
	trait_aliases     **ZendTraitAlias
	trait_precedences **ZendTraitPrecedence
	info              struct /* union */ {
		user struct {
			filename    *ZendString
			line_start  uint32
			line_end    uint32
			doc_comment *ZendString
		}
		internal struct {
			builtin_functions *ZendFunctionEntry
			module            *ZendModuleEntry
		}
	}
}

func (this *ZendClassEntry) InitMethods(functions []ZendFunctionEntry) {
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

func (this *ZendClassEntry) Name() string { return this.name.GetStr() }

/**
 * Getter / Setter
 */
func (this *ZendClassEntry) GetType() byte                   { return this.type_ }
func (this *ZendClassEntry) SetType(value byte)              { this.type_ = value }
func (this *ZendClassEntry) GetName() *ZendString            { return this.name }
func (this *ZendClassEntry) SetName(value *ZendString)       { this.name = value }
func (this *ZendClassEntry) GetParent() *ZendClassEntry      { return this.__0.parent }
func (this *ZendClassEntry) SetParent(value *ZendClassEntry) { this.__0.parent = value }
func (this *ZendClassEntry) GetParentName() *ZendString      { return this.__0.parent_name }
func (this *ZendClassEntry) SetParentName(value *ZendString) { this.__0.parent_name = value }
func (this *ZendClassEntry) GetRefcount() int                { return this.refcount }
func (this *ZendClassEntry) SetRefcount(value int)           { this.refcount = value }
func (this *ZendClassEntry) GetCeFlags() uint32              { return this.ce_flags }
func (this *ZendClassEntry) SetCeFlags(value uint32)         { this.ce_flags = value }
func (this *ZendClassEntry) GetDefaultPropertiesCount() int  { return this.default_properties_count }
func (this *ZendClassEntry) SetDefaultPropertiesCount(value int) {
	this.default_properties_count = value
}
func (this *ZendClassEntry) GetDefaultStaticMembersCount() int {
	return this.default_static_members_count
}
func (this *ZendClassEntry) SetDefaultStaticMembersCount(value int) {
	this.default_static_members_count = value
}
func (this *ZendClassEntry) GetDefaultPropertiesTable() *Zval { return this.default_properties_table }
func (this *ZendClassEntry) SetDefaultPropertiesTable(value *Zval) {
	this.default_properties_table = value
}
func (this *ZendClassEntry) GetDefaultStaticMembersTable() *Zval {
	return this.default_static_members_table
}
func (this *ZendClassEntry) SetDefaultStaticMembersTable(value *Zval) {
	this.default_static_members_table = value
}
func (this *ZendClassEntry) GetStaticMembersTablePtr() **Zval { return this.static_members_table__ptr }
func (this *ZendClassEntry) GetFunctionTable() HashTable      { return this.function_table }
func (this *ZendClassEntry) GetPropertiesInfo() HashTable     { return this.properties_info }
func (this *ZendClassEntry) GetConstantsTable() HashTable     { return this.constants_table }
func (this *ZendClassEntry) GetPropertiesInfoTable() **ZendPropertyInfo {
	return this.properties_info_table
}
func (this *ZendClassEntry) SetPropertiesInfoTable(value **ZendPropertyInfo) {
	this.properties_info_table = value
}
func (this *ZendClassEntry) GetConstructor() *ZendFunction          { return this.constructor }
func (this *ZendClassEntry) SetConstructor(value *ZendFunction)     { this.constructor = value }
func (this *ZendClassEntry) GetDestructor() *ZendFunction           { return this.destructor }
func (this *ZendClassEntry) SetDestructor(value *ZendFunction)      { this.destructor = value }
func (this *ZendClassEntry) GetClone() *ZendFunction                { return this.clone }
func (this *ZendClassEntry) SetClone(value *ZendFunction)           { this.clone = value }
func (this *ZendClassEntry) GetGet() *ZendFunction                  { return this.__get }
func (this *ZendClassEntry) SetGet(value *ZendFunction)             { this.__get = value }
func (this *ZendClassEntry) GetSet() *ZendFunction                  { return this.__set }
func (this *ZendClassEntry) SetSet(value *ZendFunction)             { this.__set = value }
func (this *ZendClassEntry) GetUnset() *ZendFunction                { return this.__unset }
func (this *ZendClassEntry) SetUnset(value *ZendFunction)           { this.__unset = value }
func (this *ZendClassEntry) GetIsset() *ZendFunction                { return this.__isset }
func (this *ZendClassEntry) SetIsset(value *ZendFunction)           { this.__isset = value }
func (this *ZendClassEntry) GetCall() *ZendFunction                 { return this.__call }
func (this *ZendClassEntry) SetCall(value *ZendFunction)            { this.__call = value }
func (this *ZendClassEntry) GetCallstatic() *ZendFunction           { return this.__callstatic }
func (this *ZendClassEntry) SetCallstatic(value *ZendFunction)      { this.__callstatic = value }
func (this *ZendClassEntry) GetTostring() *ZendFunction             { return this.__tostring }
func (this *ZendClassEntry) SetTostring(value *ZendFunction)        { this.__tostring = value }
func (this *ZendClassEntry) GetDebugInfo() *ZendFunction            { return this.__debugInfo }
func (this *ZendClassEntry) SetDebugInfo(value *ZendFunction)       { this.__debugInfo = value }
func (this *ZendClassEntry) GetSerializeFunc() *ZendFunction        { return this.serialize_func }
func (this *ZendClassEntry) SetSerializeFunc(value *ZendFunction)   { this.serialize_func = value }
func (this *ZendClassEntry) GetUnserializeFunc() *ZendFunction      { return this.unserialize_func }
func (this *ZendClassEntry) SetUnserializeFunc(value *ZendFunction) { this.unserialize_func = value }
func (this *ZendClassEntry) GetIteratorFuncsPtr() *ZendClassIteratorFuncs {
	return this.iterator_funcs_ptr
}
func (this *ZendClassEntry) SetIteratorFuncsPtr(value *ZendClassIteratorFuncs) {
	this.iterator_funcs_ptr = value
}
func (this *ZendClassEntry) GetCreateObject() func(class_type *ZendClassEntry) *ZendObject {
	return this.__1.create_object
}
func (this *ZendClassEntry) SetCreateObject(value func(class_type *ZendClassEntry) *ZendObject) {
	this.__1.create_object = value
}
func (this *ZendClassEntry) GetInterfaceGetsImplemented() func(iface *ZendClassEntry, class_type *ZendClassEntry) int {
	return this.__1.interface_gets_implemented
}
func (this *ZendClassEntry) SetInterfaceGetsImplemented(value func(iface *ZendClassEntry, class_type *ZendClassEntry) int) {
	this.__1.interface_gets_implemented = value
}
func (this *ZendClassEntry) GetGetIterator() func(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	return this.get_iterator
}
func (this *ZendClassEntry) SetGetIterator(value func(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator) {
	this.get_iterator = value
}
func (this *ZendClassEntry) GetGetStaticMethod() func(ce *ZendClassEntry, method *ZendString) *ZendFunction {
	return this.get_static_method
}
func (this *ZendClassEntry) SetGetStaticMethod(value func(ce *ZendClassEntry, method *ZendString) *ZendFunction) {
	this.get_static_method = value
}
func (this *ZendClassEntry) GetSerialize() func(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	return this.serialize
}
func (this *ZendClassEntry) SetSerialize(value func(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int) {
	this.serialize = value
}
func (this *ZendClassEntry) GetUnserialize() func(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	return this.unserialize
}
func (this *ZendClassEntry) SetUnserialize(value func(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int) {
	this.unserialize = value
}
func (this *ZendClassEntry) GetNumInterfaces() uint32               { return this.num_interfaces }
func (this *ZendClassEntry) SetNumInterfaces(value uint32)          { this.num_interfaces = value }
func (this *ZendClassEntry) GetNumTraits() uint32                   { return this.num_traits }
func (this *ZendClassEntry) SetNumTraits(value uint32)              { this.num_traits = value }
func (this *ZendClassEntry) GetInterfaces() **ZendClassEntry        { return this.__2.interfaces }
func (this *ZendClassEntry) SetInterfaces(value **ZendClassEntry)   { this.__2.interfaces = value }
func (this *ZendClassEntry) GetInterfaceNames() *ZendClassName      { return this.__2.interface_names }
func (this *ZendClassEntry) SetInterfaceNames(value *ZendClassName) { this.__2.interface_names = value }
func (this *ZendClassEntry) GetTraitNames() *ZendClassName          { return this.trait_names }
func (this *ZendClassEntry) SetTraitNames(value *ZendClassName)     { this.trait_names = value }
func (this *ZendClassEntry) GetTraitAliases() **ZendTraitAlias      { return this.trait_aliases }
func (this *ZendClassEntry) SetTraitAliases(value **ZendTraitAlias) { this.trait_aliases = value }
func (this *ZendClassEntry) GetTraitPrecedences() **ZendTraitPrecedence {
	return this.trait_precedences
}
func (this *ZendClassEntry) SetTraitPrecedences(value **ZendTraitPrecedence) {
	this.trait_precedences = value
}
func (this *ZendClassEntry) GetFilename() *ZendString      { return this.info.user.filename }
func (this *ZendClassEntry) SetFilename(value *ZendString) { this.info.user.filename = value }

// func (this *ZendClassEntry)  GetLineStart() uint32      { return this.info.user.line_start }
func (this *ZendClassEntry) SetLineStart(value uint32) { this.info.user.line_start = value }

// func (this *ZendClassEntry)  GetLineEnd() uint32      { return this.info.user.line_end }
func (this *ZendClassEntry) SetLineEnd(value uint32)         { this.info.user.line_end = value }
func (this *ZendClassEntry) GetDocComment() *ZendString      { return this.info.user.doc_comment }
func (this *ZendClassEntry) SetDocComment(value *ZendString) { this.info.user.doc_comment = value }
func (this *ZendClassEntry) GetBuiltinFunctions() *ZendFunctionEntry {
	return this.info.internal.builtin_functions
}
func (this *ZendClassEntry) SetBuiltinFunctions(value *ZendFunctionEntry) {
	this.info.internal.builtin_functions = value
}
func (this *ZendClassEntry) GetModule() *ZendModuleEntry      { return this.info.internal.module }
func (this *ZendClassEntry) SetModule(value *ZendModuleEntry) { this.info.internal.module = value }

/* ZendClassEntry.ce_flags */
func (this *ZendClassEntry) AddCeFlags(value uint32)      { this.ce_flags |= value }
func (this *ZendClassEntry) SubCeFlags(value uint32)      { this.ce_flags &^= value }
func (this *ZendClassEntry) HasCeFlags(value uint32) bool { return this.ce_flags&value != 0 }
func (this *ZendClassEntry) SwitchCeFlags(value uint32, cond bool) {
	if cond {
		this.AddCeFlags(value)
	} else {
		this.SubCeFlags(value)
	}
}
func (this ZendClassEntry) IsConstantsUpdated() bool {
	return this.HasCeFlags(ZEND_ACC_CONSTANTS_UPDATED)
}
func (this ZendClassEntry) IsInterface() bool      { return this.HasCeFlags(ZEND_ACC_INTERFACE) }
func (this ZendClassEntry) IsTrait() bool          { return this.HasCeFlags(ZEND_ACC_TRAIT) }
func (this ZendClassEntry) IsImmutable() bool      { return this.HasCeFlags(ZEND_ACC_IMMUTABLE) }
func (this ZendClassEntry) IsResolvedParent() bool { return this.HasCeFlags(ZEND_ACC_RESOLVED_PARENT) }
func (this ZendClassEntry) IsLinked() bool         { return this.HasCeFlags(ZEND_ACC_LINKED) }
func (this ZendClassEntry) IsImplementTraits() bool {
	return this.HasCeFlags(ZEND_ACC_IMPLEMENT_TRAITS)
}
func (this ZendClassEntry) IsHasStaticInMethods() bool {
	return this.HasCeFlags(ZEND_HAS_STATIC_IN_METHODS)
}
func (this ZendClassEntry) IsNearlyLinked() bool { return this.HasCeFlags(ZEND_ACC_NEARLY_LINKED) }
func (this ZendClassEntry) IsResolvedInterfaces() bool {
	return this.HasCeFlags(ZEND_ACC_RESOLVED_INTERFACES)
}
func (this ZendClassEntry) IsFinal() bool { return this.HasCeFlags(ZEND_ACC_FINAL) }
func (this ZendClassEntry) IsImplementInterfaces() bool {
	return this.HasCeFlags(ZEND_ACC_IMPLEMENT_INTERFACES)
}
func (this ZendClassEntry) IsImplicitAbstractClass() bool {
	return this.HasCeFlags(ZEND_ACC_IMPLICIT_ABSTRACT_CLASS)
}
func (this ZendClassEntry) IsUnresolvedVariance() bool {
	return this.HasCeFlags(ZEND_ACC_UNRESOLVED_VARIANCE)
}
func (this ZendClassEntry) IsHasUnlinkedUses() bool {
	return this.HasCeFlags(ZEND_ACC_HAS_UNLINKED_USES)
}
func (this ZendClassEntry) IsUseGuards() bool { return this.HasCeFlags(ZEND_ACC_USE_GUARDS) }
func (this ZendClassEntry) IsPropertyTypesResolved() bool {
	return this.HasCeFlags(ZEND_ACC_PROPERTY_TYPES_RESOLVED)
}
func (this ZendClassEntry) IsExplicitAbstractClass() bool {
	return this.HasCeFlags(ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)
}
func (this ZendClassEntry) IsHasTypeHints() bool { return this.HasCeFlags(ZEND_ACC_HAS_TYPE_HINTS) }
func (this ZendClassEntry) IsPreloaded() bool    { return this.HasCeFlags(ZEND_ACC_PRELOADED) }
func (this ZendClassEntry) IsInherited() bool    { return this.HasCeFlags(ZEND_ACC_INHERITED) }
func (this ZendClassEntry) IsTopLevel() bool     { return this.HasCeFlags(ZEND_ACC_TOP_LEVEL) }
func (this ZendClassEntry) IsReuseGetIterator() bool {
	return this.HasCeFlags(ZEND_ACC_REUSE_GET_ITERATOR)
}
func (this *ZendClassEntry) SetIsConstantsUpdated(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_CONSTANTS_UPDATED, cond)
}
func (this *ZendClassEntry) SetIsInterface(cond bool) { this.SwitchCeFlags(ZEND_ACC_INTERFACE, cond) }
func (this *ZendClassEntry) SetIsTrait(cond bool)     { this.SwitchCeFlags(ZEND_ACC_TRAIT, cond) }
func (this *ZendClassEntry) SetIsImmutable(cond bool) { this.SwitchCeFlags(ZEND_ACC_IMMUTABLE, cond) }
func (this *ZendClassEntry) SetIsResolvedParent(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_RESOLVED_PARENT, cond)
}
func (this *ZendClassEntry) SetIsLinked(cond bool) { this.SwitchCeFlags(ZEND_ACC_LINKED, cond) }
func (this *ZendClassEntry) SetIsImplementTraits(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_IMPLEMENT_TRAITS, cond)
}
func (this *ZendClassEntry) SetIsHasStaticInMethods(cond bool) {
	this.SwitchCeFlags(ZEND_HAS_STATIC_IN_METHODS, cond)
}
func (this *ZendClassEntry) SetIsNearlyLinked(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_NEARLY_LINKED, cond)
}
func (this *ZendClassEntry) SetIsResolvedInterfaces(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_RESOLVED_INTERFACES, cond)
}
func (this *ZendClassEntry) SetIsFinal(cond bool) { this.SwitchCeFlags(ZEND_ACC_FINAL, cond) }
func (this *ZendClassEntry) SetIsImplementInterfaces(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_IMPLEMENT_INTERFACES, cond)
}
func (this *ZendClassEntry) SetIsImplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_IMPLICIT_ABSTRACT_CLASS, cond)
}
func (this *ZendClassEntry) SetIsUnresolvedVariance(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_UNRESOLVED_VARIANCE, cond)
}
func (this *ZendClassEntry) SetIsHasUnlinkedUses(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_HAS_UNLINKED_USES, cond)
}
func (this *ZendClassEntry) SetIsUseGuards(cond bool) { this.SwitchCeFlags(ZEND_ACC_USE_GUARDS, cond) }
func (this *ZendClassEntry) SetIsPropertyTypesResolved(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_PROPERTY_TYPES_RESOLVED, cond)
}
func (this *ZendClassEntry) SetIsExplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_EXPLICIT_ABSTRACT_CLASS, cond)
}
func (this *ZendClassEntry) SetIsHasTypeHints(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_HAS_TYPE_HINTS, cond)
}
func (this *ZendClassEntry) SetIsPreloaded(cond bool) { this.SwitchCeFlags(ZEND_ACC_PRELOADED, cond) }
func (this *ZendClassEntry) SetIsInherited(cond bool) { this.SwitchCeFlags(ZEND_ACC_INHERITED, cond) }
func (this *ZendClassEntry) SetIsTopLevel(cond bool)  { this.SwitchCeFlags(ZEND_ACC_TOP_LEVEL, cond) }
func (this *ZendClassEntry) SetIsReuseGetIterator(cond bool) {
	this.SwitchCeFlags(ZEND_ACC_REUSE_GET_ITERATOR, cond)
}
