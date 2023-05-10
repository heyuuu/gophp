package types

import (
	"github.com/heyuuu/gophp/zend"
)

type FunctionTable = *Table[IFunction]
type PropertyTable = *Table[*zend.ZendPropertyInfo]
type ClassConstantTable = *Table[*zend.ZendClassConstant]

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
	ceFlags                      uint32
	default_properties_count     int
	default_static_members_count int
	default_properties_table     *Zval
	default_static_members_table *Zval
	static_members_table__ptr    **Zval

	functionTable FunctionTable
	propertyTable PropertyTable
	constantTable ClassConstantTable

	properties_info_table []*zend.ZendPropertyInfo

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

func NewUserClass(name string) *ClassEntry {
	ce := &ClassEntry{
		name:  name,
		type_: zend.ZEND_USER_CLASS,
	}
	// ZendInitializeClassData
	ce.SetCeFlags(zend.AccConstantsUpdated)
	if (zend.CG__().GetCompilerOptions() & zend.ZEND_COMPILE_GUARDS) != 0 {
		ce.SetIsUseGuards(true)
	}
	ce.InitTables()
	zend.ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())
	return ce
}

func NewInternalClass(origCe *ClassEntry) *ClassEntry {
	var ce = new(ClassEntry)
	*ce = *origCe
	ce.type_ = zend.ZEND_INTERNAL_CLASS

	// ZendInitializeClassData
	ce.SetCeFlags(zend.AccConstantsUpdated)
	if (zend.CG__().GetCompilerOptions() & zend.ZEND_COMPILE_GUARDS) != 0 {
		ce.SetIsUseGuards(true)
	}
	ce.SetDefaultPropertiesTable(nil)
	ce.SetDefaultStaticMembersTable(nil)
	ce.InitTables()
	zend.ZEND_MAP_PTR_INIT(ce.static_members_table, nil)
	ce.SetDefaultPropertiesCount(0)
	ce.SetDefaultStaticMembersCount(0)
	ce.SetPropertiesInfoTable(nil)

	return ce
}

func NewDisabledClass(origCe *ClassEntry, createObject func(*ClassEntry) *ZendObject) *ClassEntry {
	ce := &ClassEntry{
		type_: origCe.type_,
		name:  origCe.name,
	}
	ce.SetCreateObject(createObject)
	return ce
}

func (ce *ClassEntry) ClearMethods() {
	ce.functionTable.Destroy()
}

func NewClassEntry(name string, functions []FunctionEntry) *ClassEntry {
	class := &ClassEntry{
		name: name,
	}
	class.info.internal.builtin_functions = functions
	return class
}
func (ce *ClassEntry) InitTables() {
	ce.functionTable = NewLcTable[IFunction](nil)
	ce.propertyTable = NewTable[*zend.ZendPropertyInfo](nil)
	ce.constantTable = NewTable[*zend.ZendClassConstant](nil)
}

func (ce *ClassEntry) Name() string { return ce.name }

func (ce *ClassEntry) FunctionTable() FunctionTable       { return ce.functionTable }
func (ce *ClassEntry) PropertyTable() PropertyTable       { return ce.propertyTable }
func (ce *ClassEntry) ConstantsTable() ClassConstantTable { return ce.constantTable }

/**
 * Getter / Setter
 */
func (ce *ClassEntry) GetName() *String        { return NewString(ce.name) }
func (ce *ClassEntry) SetNameVal(value string) { ce.name = value }

func (ce *ClassEntry) GetType() byte                  { return ce.type_ }
func (ce *ClassEntry) SetType(value byte)             { ce.type_ = value }
func (ce *ClassEntry) GetParent() *ClassEntry         { return ce.__0.parent }
func (ce *ClassEntry) SetParent(value *ClassEntry)    { ce.__0.parent = value }
func (ce *ClassEntry) GetParentName() *String         { return ce.__0.parent_name }
func (ce *ClassEntry) SetParentName(value *String)    { ce.__0.parent_name = value }
func (ce *ClassEntry) GetCeFlags() uint32             { return ce.ceFlags }
func (ce *ClassEntry) SetCeFlags(value uint32)        { ce.ceFlags = value }
func (ce *ClassEntry) GetDefaultPropertiesCount() int { return ce.default_properties_count }
func (ce *ClassEntry) SetDefaultPropertiesCount(value int) {
	ce.default_properties_count = value
}
func (ce *ClassEntry) GetDefaultStaticMembersCount() int {
	return ce.default_static_members_count
}
func (ce *ClassEntry) SetDefaultStaticMembersCount(value int) {
	ce.default_static_members_count = value
}
func (ce *ClassEntry) GetDefaultPropertiesTable() *Zval {
	return ce.default_properties_table
}
func (ce *ClassEntry) SetDefaultPropertiesTable(value *Zval) {
	ce.default_properties_table = value
}
func (ce *ClassEntry) GetDefaultStaticMembersTable() *Zval {
	return ce.default_static_members_table
}
func (ce *ClassEntry) SetDefaultStaticMembersTable(value *Zval) {
	ce.default_static_members_table = value
}
func (ce *ClassEntry) GetStaticMembersTablePtr() **Zval {
	return ce.static_members_table__ptr
}
func (ce *ClassEntry) GetPropertiesInfoTable() []*zend.ZendPropertyInfo {
	return ce.properties_info_table
}
func (ce *ClassEntry) SetPropertiesInfoTable(value []*zend.ZendPropertyInfo) {
	ce.properties_info_table = value
}
func (ce *ClassEntry) GetConstructor() IFunction        { return ce.constructor }
func (ce *ClassEntry) SetConstructor(value IFunction)   { ce.constructor = value }
func (ce *ClassEntry) GetDestructor() IFunction         { return ce.destructor }
func (ce *ClassEntry) SetDestructor(value IFunction)    { ce.destructor = value }
func (ce *ClassEntry) GetClone() IFunction              { return ce.clone }
func (ce *ClassEntry) SetClone(value IFunction)         { ce.clone = value }
func (ce *ClassEntry) GetGet() IFunction                { return ce.__get }
func (ce *ClassEntry) SetGet(value IFunction)           { ce.__get = value }
func (ce *ClassEntry) GetSet() IFunction                { return ce.__set }
func (ce *ClassEntry) SetSet(value IFunction)           { ce.__set = value }
func (ce *ClassEntry) GetUnset() IFunction              { return ce.__unset }
func (ce *ClassEntry) SetUnset(value IFunction)         { ce.__unset = value }
func (ce *ClassEntry) GetIsset() IFunction              { return ce.__isset }
func (ce *ClassEntry) SetIsset(value IFunction)         { ce.__isset = value }
func (ce *ClassEntry) GetCall() IFunction               { return ce.__call }
func (ce *ClassEntry) SetCall(value IFunction)          { ce.__call = value }
func (ce *ClassEntry) GetCallstatic() IFunction         { return ce.__callstatic }
func (ce *ClassEntry) SetCallstatic(value IFunction)    { ce.__callstatic = value }
func (ce *ClassEntry) GetTostring() IFunction           { return ce.__tostring }
func (ce *ClassEntry) SetTostring(value IFunction)      { ce.__tostring = value }
func (ce *ClassEntry) GetDebugInfo() IFunction          { return ce.__debugInfo }
func (ce *ClassEntry) SetDebugInfo(value IFunction)     { ce.__debugInfo = value }
func (ce *ClassEntry) GetSerializeFunc() IFunction      { return ce.serialize_func }
func (ce *ClassEntry) SetSerializeFunc(value IFunction) { ce.serialize_func = value }
func (ce *ClassEntry) GetUnserializeFunc() IFunction    { return ce.unserialize_func }
func (ce *ClassEntry) SetUnserializeFunc(value IFunction) {
	ce.unserialize_func = value
}
func (ce *ClassEntry) GetIteratorFuncsPtr() *zend.ZendClassIteratorFuncs {
	return ce.iterator_funcs_ptr
}
func (ce *ClassEntry) SetIteratorFuncsPtr(value *zend.ZendClassIteratorFuncs) {
	ce.iterator_funcs_ptr = value
}
func (ce *ClassEntry) GetCreateObject() func(class_type *ClassEntry) *ZendObject {
	return ce.__1.create_object
}
func (ce *ClassEntry) SetCreateObject(value func(class_type *ClassEntry) *ZendObject) {
	ce.__1.create_object = value
}
func (ce *ClassEntry) GetInterfaceGetsImplemented() func(iface *ClassEntry, class_type *ClassEntry) int {
	return ce.__1.interface_gets_implemented
}
func (ce *ClassEntry) SetInterfaceGetsImplemented(value func(iface *ClassEntry, class_type *ClassEntry) int) {
	ce.__1.interface_gets_implemented = value
}
func (ce *ClassEntry) GetGetIterator() func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator {
	return ce.get_iterator
}
func (ce *ClassEntry) SetGetIterator(value func(ce *ClassEntry, object *Zval, by_ref int) *zend.ZendObjectIterator) {
	ce.get_iterator = value
}
func (ce *ClassEntry) GetGetStaticMethod() func(ce *ClassEntry, method *String) IFunction {
	return ce.get_static_method
}
func (ce *ClassEntry) SetGetStaticMethod(value func(ce *ClassEntry, method *String) IFunction) {
	ce.get_static_method = value
}
func (ce *ClassEntry) GetSerialize() func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int {
	return ce.serialize
}
func (ce *ClassEntry) SetSerialize(value func(object *Zval, buffer **uint8, buf_len *int, data *zend.ZendSerializeData) int) {
	ce.serialize = value
}
func (ce *ClassEntry) GetUnserialize() func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int {
	return ce.unserialize
}
func (ce *ClassEntry) SetUnserialize(value func(object *Zval, ce *ClassEntry, buf *uint8, buf_len int, data *zend.ZendUnserializeData) int) {
	ce.unserialize = value
}
func (ce *ClassEntry) GetNumInterfaces() uint32                { return ce.num_interfaces }
func (ce *ClassEntry) SetNumInterfaces(value uint32)           { ce.num_interfaces = value }
func (ce *ClassEntry) GetNumTraits() uint32                    { return ce.num_traits }
func (ce *ClassEntry) SetNumTraits(value uint32)               { ce.num_traits = value }
func (ce *ClassEntry) GetInterfaces() []*ClassEntry            { return ce.__2.interfaces }
func (ce *ClassEntry) SetInterfaces(value []*ClassEntry)       { ce.__2.interfaces = value }
func (ce *ClassEntry) GetInterfaceNames() []zend.ZendClassName { return ce.__2.interface_names }
func (ce *ClassEntry) SetInterfaceNames(value []zend.ZendClassName) {
	ce.__2.interface_names = value
}
func (ce *ClassEntry) GetTraitNames() *zend.ZendClassName           { return ce.trait_names }
func (ce *ClassEntry) SetTraitNames(value *zend.ZendClassName)      { ce.trait_names = value }
func (ce *ClassEntry) GetTraitAliases() []*zend.ZendTraitAlias      { return ce.trait_aliases }
func (ce *ClassEntry) SetTraitAliases(value []*zend.ZendTraitAlias) { ce.trait_aliases = value }
func (ce *ClassEntry) GetTraitPrecedences() []*zend.ZendTraitPrecedence {
	return ce.trait_precedences
}
func (ce *ClassEntry) SetTraitPrecedences(value []*zend.ZendTraitPrecedence) {
	ce.trait_precedences = value
}
func (ce *ClassEntry) GetFilename() *String      { return ce.info.user.filename }
func (ce *ClassEntry) SetFilename(value *String) { ce.info.user.filename = value }

// func (this *ClassEntry)  GetLineStart() uint32      { return this.info.user.line_start }
func (ce *ClassEntry) SetLineStart(value uint32) { ce.info.user.line_start = value }

// func (this *ClassEntry)  GetLineEnd() uint32      { return this.info.user.line_end }
func (ce *ClassEntry) SetLineEnd(value uint32) { ce.info.user.line_end = value }
func (ce *ClassEntry) GetDocComment() *String  { return ce.info.user.doc_comment }
func (ce *ClassEntry) SetDocComment(value *String) {
	ce.info.user.doc_comment = value
}
func (ce *ClassEntry) GetBuiltinFunctions() *FunctionEntry {
	return ce.info.internal.builtin_functions
}
func (ce *ClassEntry) SetBuiltinFunctions(value *FunctionEntry) {
	ce.info.internal.builtin_functions = value
}
func (ce *ClassEntry) GetModule() *zend.ModuleEntry      { return ce.info.internal.module }
func (ce *ClassEntry) SetModule(value *zend.ModuleEntry) { ce.info.internal.module = value }

/* ClassEntry.ceFlags */
func (ce *ClassEntry) AddCeFlags(value uint32)      { ce.ceFlags |= value }
func (ce *ClassEntry) HasCeFlags(value uint32) bool { return ce.ceFlags&value != 0 }
func (ce *ClassEntry) SwitchCeFlags(value uint32, cond bool) {
	if cond {
		ce.ceFlags |= value
	} else {
		ce.ceFlags &^= value
	}
}
func (ce ClassEntry) IsConstantsUpdated() bool {
	return ce.HasCeFlags(zend.AccConstantsUpdated)
}
func (ce ClassEntry) IsInterface() bool { return ce.HasCeFlags(zend.AccInterface) }
func (ce ClassEntry) IsTrait() bool     { return ce.HasCeFlags(zend.AccTrait) }
func (ce ClassEntry) IsImmutable() bool { return ce.HasCeFlags(zend.AccImmutable) }
func (ce ClassEntry) IsResolvedParent() bool {
	return ce.HasCeFlags(zend.AccResolvedParent)
}
func (ce ClassEntry) IsLinked() bool { return ce.HasCeFlags(zend.AccLinked) }
func (ce ClassEntry) IsImplementTraits() bool {
	return ce.HasCeFlags(zend.AccImplementTraits)
}
func (ce ClassEntry) IsHasStaticInMethods() bool {
	return ce.HasCeFlags(zend.AccHasStaticInMethods)
}
func (ce ClassEntry) IsNearlyLinked() bool { return ce.HasCeFlags(zend.AccNearlyLinked) }
func (ce ClassEntry) IsResolvedInterfaces() bool {
	return ce.HasCeFlags(zend.AccResolvedInterfaces)
}
func (ce ClassEntry) IsFinal() bool { return ce.HasCeFlags(zend.AccFinal) }
func (ce ClassEntry) IsImplementInterfaces() bool {
	return ce.HasCeFlags(zend.AccImplementInterfaces)
}
func (ce ClassEntry) IsImplicitAbstractClass() bool {
	return ce.HasCeFlags(zend.AccImplicitAbstractClass)
}
func (ce ClassEntry) IsUnresolvedVariance() bool {
	return ce.HasCeFlags(zend.AccUnresolvedVariance)
}
func (ce ClassEntry) IsHasUnlinkedUses() bool {
	return ce.HasCeFlags(zend.AccHasUnlinkedUses)
}
func (ce ClassEntry) IsUseGuards() bool { return ce.HasCeFlags(zend.AccUseGuards) }
func (ce ClassEntry) IsPropertyTypesResolved() bool {
	return ce.HasCeFlags(zend.AccPropertyTypesResolved)
}
func (ce ClassEntry) IsExplicitAbstractClass() bool {
	return ce.HasCeFlags(zend.AccExplicitAbstractClass)
}
func (ce ClassEntry) IsHasTypeHints() bool {
	return ce.HasCeFlags(zend.AccHasTypeHints)
}
func (ce ClassEntry) IsPreloaded() bool { return ce.HasCeFlags(zend.AccPreloaded) }
func (ce ClassEntry) IsInherited() bool { return ce.HasCeFlags(zend.AccInherited) }
func (ce ClassEntry) IsTopLevel() bool  { return ce.HasCeFlags(zend.AccTopLevel) }
func (ce ClassEntry) IsReuseGetIterator() bool {
	return ce.HasCeFlags(zend.AccReuseGetIterator)
}
func (ce *ClassEntry) SetIsConstantsUpdated(cond bool) {
	ce.SwitchCeFlags(zend.AccConstantsUpdated, cond)
}
func (ce *ClassEntry) SetIsInterface(cond bool) {
	ce.SwitchCeFlags(zend.AccInterface, cond)
}
func (ce *ClassEntry) SetIsTrait(cond bool) { ce.SwitchCeFlags(zend.AccTrait, cond) }
func (ce *ClassEntry) SetIsImmutable(cond bool) {
	ce.SwitchCeFlags(zend.AccImmutable, cond)
}
func (ce *ClassEntry) SetIsResolvedParent(cond bool) {
	ce.SwitchCeFlags(zend.AccResolvedParent, cond)
}
func (ce *ClassEntry) SetIsLinked(cond bool) { ce.SwitchCeFlags(zend.AccLinked, cond) }
func (ce *ClassEntry) SetIsImplementTraits(cond bool) {
	ce.SwitchCeFlags(zend.AccImplementTraits, cond)
}
func (ce *ClassEntry) SetIsHasStaticInMethods(cond bool) {
	ce.SwitchCeFlags(zend.AccHasStaticInMethods, cond)
}
func (ce *ClassEntry) SetIsNearlyLinked(cond bool) {
	ce.SwitchCeFlags(zend.AccNearlyLinked, cond)
}
func (ce *ClassEntry) SetIsResolvedInterfaces(cond bool) {
	ce.SwitchCeFlags(zend.AccResolvedInterfaces, cond)
}
func (ce *ClassEntry) SetIsFinal(cond bool) { ce.SwitchCeFlags(zend.AccFinal, cond) }
func (ce *ClassEntry) SetIsImplementInterfaces(cond bool) {
	ce.SwitchCeFlags(zend.AccImplementInterfaces, cond)
}
func (ce *ClassEntry) SetIsImplicitAbstractClass(cond bool) {
	ce.SwitchCeFlags(zend.AccImplicitAbstractClass, cond)
}
func (ce *ClassEntry) SetIsUnresolvedVariance(cond bool) {
	ce.SwitchCeFlags(zend.AccUnresolvedVariance, cond)
}
func (ce *ClassEntry) SetIsHasUnlinkedUses(cond bool) {
	ce.SwitchCeFlags(zend.AccHasUnlinkedUses, cond)
}
func (ce *ClassEntry) SetIsUseGuards(cond bool) {
	ce.SwitchCeFlags(zend.AccUseGuards, cond)
}
func (ce *ClassEntry) SetIsPropertyTypesResolved(cond bool) {
	ce.SwitchCeFlags(zend.AccPropertyTypesResolved, cond)
}
func (ce *ClassEntry) SetIsExplicitAbstractClass(cond bool) {
	ce.SwitchCeFlags(zend.AccExplicitAbstractClass, cond)
}
func (ce *ClassEntry) SetIsHasTypeHints(cond bool) {
	ce.SwitchCeFlags(zend.AccHasTypeHints, cond)
}
func (ce *ClassEntry) SetIsPreloaded(cond bool) {
	ce.SwitchCeFlags(zend.AccPreloaded, cond)
}
func (ce *ClassEntry) SetIsInherited(cond bool) {
	ce.SwitchCeFlags(zend.AccInherited, cond)
}
func (ce *ClassEntry) SetIsTopLevel(cond bool) {
	ce.SwitchCeFlags(zend.AccTopLevel, cond)
}
func (ce *ClassEntry) SetIsReuseGetIterator(cond bool) {
	ce.SwitchCeFlags(zend.AccReuseGetIterator, cond)
}
