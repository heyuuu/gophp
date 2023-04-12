package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/internal"
)

type FunctionTable = *internal.Table[IFunction]
type PropertyTable = *internal.Table[*zend.ZendPropertyInfo]
type ClassConstantTable = *internal.Table[*zend.ZendClassConstant]

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

	functionTable FunctionTable
	propertyTable PropertyTable
	constantTable ClassConstantTable

	properties_info       Array
	constants_table       Array
	properties_info_table **zend.ZendPropertyInfo

	constructor      IFunction
	destructor       IFunction
	clone            IFunction
	__get            IFunction
	__set            IFunction
	__unset          IFunction
	__isset          IFunction
	__call           IFunction
	__callstatic     IFunction
	__tostring       IFunction
	__debugInfo      IFunction
	serialize_func   IFunction
	unserialize_func IFunction

	iterator_funcs_ptr *zend.ZendClassIteratorFuncs
	__1                struct /* union */ {
		create_object              func(class_type *ClassEntry) *ZendObject
		interface_gets_implemented func(iface *ClassEntry, class_type *ClassEntry) int
	}
	get_iterator      func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator
	get_static_method func(ce *ClassEntry, method *String) IFunction
	serialize         func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int
	unserialize       func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int
	num_interfaces    uint32
	num_traits        uint32
	__2               struct /* union */ {
		interfaces      []*ClassEntry
		interface_names []zend.ZendClassName
	}
	trait_names       *zend.ZendClassName
	trait_aliases     []*zend.ZendTraitAlias
	trait_precedences []*zend.ZendTraitPrecedence
	info              struct /* union */ {
		user struct {
			filename    *String
			line_start  uint32
			line_end    uint32
			doc_comment *String
		}
		internal struct {
			builtin_functions []FunctionEntry
			module            *zend.ModuleEntry
		}
	}
}

func (this *ClassEntry) InitMethods(functions []FunctionEntry) {
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

func NewClassEntry(name string, functions []FunctionEntry) *ClassEntry {
	class := &ClassEntry{
		name: name,
	}
	class.info.internal.builtin_functions = functions
	return class
}
func (this *ClassEntry) InitTables(persistentHashes bool) {
	this.properties_info = MakeArrayEx(8, b.Cond(persistentHashes, zend.ZendDestroyPropertyInfoInternal, nil), IntBool(persistentHashes))
	this.constants_table = MakeArrayEx(8, nil, IntBool(persistentHashes))
	this.functionTable = internal.NewLcTable[IFunction](zend.ZendFunctionDtorEx)
}

func (this *ClassEntry) Name() string { return this.name }

func (this *ClassEntry) FunctionTable() FunctionTable       { return this.functionTable }
func (this *ClassEntry) PropertyTable() PropertyTable       { return this.propertyTable }
func (this *ClassEntry) ConstantsTable() ClassConstantTable { return this.constantTable }

func (this *ClassEntry) GetPropertiesInfo() Array { return this.properties_info }
func (this *ClassEntry) GetConstantsTable() Array { return this.constants_table }

/**
 * Getter / Setter
 */
func (this *ClassEntry) GetName() *String        { return NewString(this.name) }
func (this *ClassEntry) SetName(value *String)   { this.name = value.GetStr() }
func (this *ClassEntry) SetNameVal(value string) { this.name = value }

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
func (this *ClassEntry) GetPropertiesInfoTable() **zend.ZendPropertyInfo {
	return this.properties_info_table
}
func (this *ClassEntry) SetPropertiesInfoTable(value **zend.ZendPropertyInfo) {
	this.properties_info_table = value
}
func (this *ClassEntry) GetConstructor() IFunction        { return this.constructor }
func (this *ClassEntry) SetConstructor(value IFunction)   { this.constructor = value }
func (this *ClassEntry) GetDestructor() IFunction         { return this.destructor }
func (this *ClassEntry) SetDestructor(value IFunction)    { this.destructor = value }
func (this *ClassEntry) GetClone() IFunction              { return this.clone }
func (this *ClassEntry) SetClone(value IFunction)         { this.clone = value }
func (this *ClassEntry) GetGet() IFunction                { return this.__get }
func (this *ClassEntry) SetGet(value IFunction)           { this.__get = value }
func (this *ClassEntry) GetSet() IFunction                { return this.__set }
func (this *ClassEntry) SetSet(value IFunction)           { this.__set = value }
func (this *ClassEntry) GetUnset() IFunction              { return this.__unset }
func (this *ClassEntry) SetUnset(value IFunction)         { this.__unset = value }
func (this *ClassEntry) GetIsset() IFunction              { return this.__isset }
func (this *ClassEntry) SetIsset(value IFunction)         { this.__isset = value }
func (this *ClassEntry) GetCall() IFunction               { return this.__call }
func (this *ClassEntry) SetCall(value IFunction)          { this.__call = value }
func (this *ClassEntry) GetCallstatic() IFunction         { return this.__callstatic }
func (this *ClassEntry) SetCallstatic(value IFunction)    { this.__callstatic = value }
func (this *ClassEntry) GetTostring() IFunction           { return this.__tostring }
func (this *ClassEntry) SetTostring(value IFunction)      { this.__tostring = value }
func (this *ClassEntry) GetDebugInfo() IFunction          { return this.__debugInfo }
func (this *ClassEntry) SetDebugInfo(value IFunction)     { this.__debugInfo = value }
func (this *ClassEntry) GetSerializeFunc() IFunction      { return this.serialize_func }
func (this *ClassEntry) SetSerializeFunc(value IFunction) { this.serialize_func = value }
func (this *ClassEntry) GetUnserializeFunc() IFunction    { return this.unserialize_func }
func (this *ClassEntry) SetUnserializeFunc(value IFunction) {
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
func (this *ClassEntry) GetGetStaticMethod() func(ce *ClassEntry, method *String) IFunction {
	return this.get_static_method
}
func (this *ClassEntry) SetGetStaticMethod(value func(ce *ClassEntry, method *String) IFunction) {
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
func (this *ClassEntry) GetNumInterfaces() uint32                { return this.num_interfaces }
func (this *ClassEntry) SetNumInterfaces(value uint32)           { this.num_interfaces = value }
func (this *ClassEntry) GetNumTraits() uint32                    { return this.num_traits }
func (this *ClassEntry) SetNumTraits(value uint32)               { this.num_traits = value }
func (this *ClassEntry) GetInterfaces() []*ClassEntry            { return this.__2.interfaces }
func (this *ClassEntry) SetInterfaces(value []*ClassEntry)       { this.__2.interfaces = value }
func (this *ClassEntry) GetInterfaceNames() []zend.ZendClassName { return this.__2.interface_names }
func (this *ClassEntry) SetInterfaceNames(value []zend.ZendClassName) {
	this.__2.interface_names = value
}
func (this *ClassEntry) GetTraitNames() *zend.ZendClassName           { return this.trait_names }
func (this *ClassEntry) SetTraitNames(value *zend.ZendClassName)      { this.trait_names = value }
func (this *ClassEntry) GetTraitAliases() []*zend.ZendTraitAlias      { return this.trait_aliases }
func (this *ClassEntry) SetTraitAliases(value []*zend.ZendTraitAlias) { this.trait_aliases = value }
func (this *ClassEntry) GetTraitPrecedences() []*zend.ZendTraitPrecedence {
	return this.trait_precedences
}
func (this *ClassEntry) SetTraitPrecedences(value []*zend.ZendTraitPrecedence) {
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
func (this *ClassEntry) GetBuiltinFunctions() *FunctionEntry {
	return this.info.internal.builtin_functions
}
func (this *ClassEntry) SetBuiltinFunctions(value *FunctionEntry) {
	this.info.internal.builtin_functions = value
}
func (this *ClassEntry) GetModule() *zend.ModuleEntry      { return this.info.internal.module }
func (this *ClassEntry) SetModule(value *zend.ModuleEntry) { this.info.internal.module = value }

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
	return this.HasCeFlags(zend.AccConstantsUpdated)
}
func (this ClassEntry) IsInterface() bool { return this.HasCeFlags(zend.AccInterface) }
func (this ClassEntry) IsTrait() bool     { return this.HasCeFlags(zend.AccTrait) }
func (this ClassEntry) IsImmutable() bool { return this.HasCeFlags(zend.AccImmutable) }
func (this ClassEntry) IsResolvedParent() bool {
	return this.HasCeFlags(zend.AccResolvedParent)
}
func (this ClassEntry) IsLinked() bool { return this.HasCeFlags(zend.AccLinked) }
func (this ClassEntry) IsImplementTraits() bool {
	return this.HasCeFlags(zend.AccImplementTraits)
}
func (this ClassEntry) IsHasStaticInMethods() bool {
	return this.HasCeFlags(zend.AccHasStaticInMethods)
}
func (this ClassEntry) IsNearlyLinked() bool { return this.HasCeFlags(zend.AccNearlyLinked) }
func (this ClassEntry) IsResolvedInterfaces() bool {
	return this.HasCeFlags(zend.AccResolvedInterfaces)
}
func (this ClassEntry) IsFinal() bool { return this.HasCeFlags(zend.AccFinal) }
func (this ClassEntry) IsImplementInterfaces() bool {
	return this.HasCeFlags(zend.AccImplementInterfaces)
}
func (this ClassEntry) IsImplicitAbstractClass() bool {
	return this.HasCeFlags(zend.AccImplicitAbstractClass)
}
func (this ClassEntry) IsUnresolvedVariance() bool {
	return this.HasCeFlags(zend.AccUnresolvedVariance)
}
func (this ClassEntry) IsHasUnlinkedUses() bool {
	return this.HasCeFlags(zend.AccHasUnlinkedUses)
}
func (this ClassEntry) IsUseGuards() bool { return this.HasCeFlags(zend.AccUseGuards) }
func (this ClassEntry) IsPropertyTypesResolved() bool {
	return this.HasCeFlags(zend.AccPropertyTypesResolved)
}
func (this ClassEntry) IsExplicitAbstractClass() bool {
	return this.HasCeFlags(zend.AccExplicitAbstractClass)
}
func (this ClassEntry) IsHasTypeHints() bool {
	return this.HasCeFlags(zend.AccHasTypeHints)
}
func (this ClassEntry) IsPreloaded() bool { return this.HasCeFlags(zend.AccPreloaded) }
func (this ClassEntry) IsInherited() bool { return this.HasCeFlags(zend.AccInherited) }
func (this ClassEntry) IsTopLevel() bool  { return this.HasCeFlags(zend.AccTopLevel) }
func (this ClassEntry) IsReuseGetIterator() bool {
	return this.HasCeFlags(zend.AccReuseGetIterator)
}
func (this *ClassEntry) SetIsConstantsUpdated(cond bool) {
	this.SwitchCeFlags(zend.AccConstantsUpdated, cond)
}
func (this *ClassEntry) SetIsInterface(cond bool) {
	this.SwitchCeFlags(zend.AccInterface, cond)
}
func (this *ClassEntry) SetIsTrait(cond bool) { this.SwitchCeFlags(zend.AccTrait, cond) }
func (this *ClassEntry) SetIsImmutable(cond bool) {
	this.SwitchCeFlags(zend.AccImmutable, cond)
}
func (this *ClassEntry) SetIsResolvedParent(cond bool) {
	this.SwitchCeFlags(zend.AccResolvedParent, cond)
}
func (this *ClassEntry) SetIsLinked(cond bool) { this.SwitchCeFlags(zend.AccLinked, cond) }
func (this *ClassEntry) SetIsImplementTraits(cond bool) {
	this.SwitchCeFlags(zend.AccImplementTraits, cond)
}
func (this *ClassEntry) SetIsHasStaticInMethods(cond bool) {
	this.SwitchCeFlags(zend.AccHasStaticInMethods, cond)
}
func (this *ClassEntry) SetIsNearlyLinked(cond bool) {
	this.SwitchCeFlags(zend.AccNearlyLinked, cond)
}
func (this *ClassEntry) SetIsResolvedInterfaces(cond bool) {
	this.SwitchCeFlags(zend.AccResolvedInterfaces, cond)
}
func (this *ClassEntry) SetIsFinal(cond bool) { this.SwitchCeFlags(zend.AccFinal, cond) }
func (this *ClassEntry) SetIsImplementInterfaces(cond bool) {
	this.SwitchCeFlags(zend.AccImplementInterfaces, cond)
}
func (this *ClassEntry) SetIsImplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(zend.AccImplicitAbstractClass, cond)
}
func (this *ClassEntry) SetIsUnresolvedVariance(cond bool) {
	this.SwitchCeFlags(zend.AccUnresolvedVariance, cond)
}
func (this *ClassEntry) SetIsHasUnlinkedUses(cond bool) {
	this.SwitchCeFlags(zend.AccHasUnlinkedUses, cond)
}
func (this *ClassEntry) SetIsUseGuards(cond bool) {
	this.SwitchCeFlags(zend.AccUseGuards, cond)
}
func (this *ClassEntry) SetIsPropertyTypesResolved(cond bool) {
	this.SwitchCeFlags(zend.AccPropertyTypesResolved, cond)
}
func (this *ClassEntry) SetIsExplicitAbstractClass(cond bool) {
	this.SwitchCeFlags(zend.AccExplicitAbstractClass, cond)
}
func (this *ClassEntry) SetIsHasTypeHints(cond bool) {
	this.SwitchCeFlags(zend.AccHasTypeHints, cond)
}
func (this *ClassEntry) SetIsPreloaded(cond bool) {
	this.SwitchCeFlags(zend.AccPreloaded, cond)
}
func (this *ClassEntry) SetIsInherited(cond bool) {
	this.SwitchCeFlags(zend.AccInherited, cond)
}
func (this *ClassEntry) SetIsTopLevel(cond bool) {
	this.SwitchCeFlags(zend.AccTopLevel, cond)
}
func (this *ClassEntry) SetIsReuseGetIterator(cond bool) {
	this.SwitchCeFlags(zend.AccReuseGetIterator, cond)
}
