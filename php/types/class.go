package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend"
)

const (
	typeInternalClass = 1
	typeUserClass     = 2
)

type FunctionTable = *Table[IFunction]
type PropertyTable = *Table[*PropertyInfo]
type ClassConstantTable = *Table[*ClassConstant]

/**
 * ClassName
 */
type ClassName struct {
	name   string
	lcName string
}

func MakeClassName(name string) ClassName {
	return ClassName{
		name:   name,
		lcName: ascii.StrToLower(name),
	}
}

func (n ClassName) GetName() string   { return n.name }
func (n ClassName) GetLcName() string { return n.lcName }

/**
 * ClassEntry
 */
type ClassEntry struct {
	typ     byte
	name    string // *String
	ceFlags uint32

	// 继承父类，在 link 前只有 parentName 可能有值，在 link 后只有 parent 可能有值(union)。
	parentName *String
	parent     *ClassEntry

	// 继承接口列表，在 link 前只有 interfaceNames 可能有值，在 link 后只有 interfaces 可能有值(union)。
	interfaceNames []ClassName
	interfaces     []*ClassEntry

	// 默认属性表
	defaultPropertiesCount int
	defaultPropertiesTable []Zval

	defaultStaticMembersCount int
	defaultStaticMembersTable *Zval
	static_members_table__ptr **Zval

	functionTable FunctionTable
	propertyTable PropertyTable
	constantTable ClassConstantTable

	properties_info_table []*PropertyInfo

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
	num_traits        uint32

	trait_names       *ClassName
	trait_aliases     []*zend.ZendTraitAlias
	trait_precedences []*zend.ZendTraitPrecedence

	// info.internal
	moduleNumber int
	// info.user
	infoUser struct {
		filename   string
		lineStart  uint32
		lineEnd    uint32
		docComment string
	}
}

func NewUserClass(name string) *ClassEntry {
	ce := &ClassEntry{
		typ:  typeUserClass,
		name: name,
	}
	ce.initData()
	return ce
}

func NewInternalClass(name string, moduleNumber int) *ClassEntry {
	var ce = &ClassEntry{
		typ:          typeInternalClass,
		name:         name,
		moduleNumber: moduleNumber,
	}
	ce.initData()
	return ce
}

func NewDisabledClass(origCe *ClassEntry, createObject func(*ClassEntry) *ZendObject) *ClassEntry {
	ce := &ClassEntry{
		typ:  origCe.typ,
		name: origCe.name,
	}
	ce.SetCreateObject(createObject)
	return ce
}

func (ce *ClassEntry) initData() {
	// ZendInitializeClassData
	ce.SetCeFlags(AccConstantsUpdated)
	if (zend.CG__().GetCompilerOptions() & zend.ZEND_COMPILE_GUARDS) != 0 {
		ce.SetIsUseGuards(true)
	}
	ce.initTables()
	if ce.typ == typeUserClass {
		zend.ZEND_MAP_PTR_INIT(ce.static_members_table, ce.GetDefaultStaticMembersTable())
	} else {
		zend.ZEND_MAP_PTR_INIT(ce.static_members_table, nil)
	}
}

func (ce *ClassEntry) initTables() {
	ce.functionTable = NewLcTable[IFunction](nil)
	ce.propertyTable = NewTable[*PropertyInfo](nil)
	ce.constantTable = NewTable[*ClassConstant](nil)
}

func (ce *ClassEntry) Name() string      { return ce.name }
func (ce *ClassEntry) ModuleNumber() int { return ce.moduleNumber }

func (ce *ClassEntry) FunctionTable() FunctionTable       { return ce.functionTable }
func (ce *ClassEntry) PropertyTable() PropertyTable       { return ce.propertyTable }
func (ce *ClassEntry) ConstantsTable() ClassConstantTable { return ce.constantTable }

// methods
func (ce *ClassEntry) GetPropertyInfo(propNum int) *PropertyInfo {
	b.Assert(0 <= propNum && propNum < ce.GetDefaultPropertiesCount())
	return ce.GetPropertiesInfoTable()[propNum]
}

// interfaces
func (ce *ClassEntry) GetNumInterfaces() int {
	if ce.IsResolvedInterfaces() {
		return len(ce.interfaces)
	} else {
		return len(ce.interfaceNames)
	}
}

func (ce *ClassEntry) GetInterfaceNames() []ClassName { return ce.interfaceNames }
func (ce *ClassEntry) ImplementInterfaceNames(names []string) {
	b.Assert(!ce.IsResolvedInterfaces())

	// string to ClassName
	var interfaceNames []ClassName
	if len(names) > 0 {
		interfaceNames = make([]ClassName, len(names))
		for i, name := range names {
			interfaceNames[i] = MakeClassName(name)
		}
	}

	// set field
	if len(interfaceNames) > 0 {
		ce.SetIsImplementInterfaces(true)
	}
	ce.interfaceNames = interfaceNames
}

func (ce *ClassEntry) GetInterfaces() []*ClassEntry { return ce.interfaces }
func (ce *ClassEntry) ResolvedInterfaces(interfaces []*ClassEntry) {
	ce.SetIsResolvedInterfaces(true)
	ce.interfaces = interfaces
	ce.interfaceNames = nil
}
func (ce *ClassEntry) AppendResolvedInterfaces(iface *ClassEntry) {
	b.Assert(ce.IsResolvedInterfaces())
	ce.interfaces = append(ce.interfaces, iface)
}

/**
 * Getter / Setter
 */
func (ce *ClassEntry) GetName() *String      { return NewString(ce.name) }
func (ce *ClassEntry) IsInternalClass() bool { return ce.typ == typeInternalClass }
func (ce *ClassEntry) IsUserClass() bool     { return ce.typ == typeUserClass }

func (ce *ClassEntry) GetType() byte                  { return ce.typ }
func (ce *ClassEntry) GetParent() *ClassEntry         { return ce.parent }
func (ce *ClassEntry) SetParent(value *ClassEntry)    { ce.parent = value }
func (ce *ClassEntry) GetParentName() *String         { return ce.parentName }
func (ce *ClassEntry) SetParentName(value string)     { ce.parentName = NewString(value) }
func (ce *ClassEntry) GetCeFlags() uint32             { return ce.ceFlags }
func (ce *ClassEntry) SetCeFlags(value uint32)        { ce.ceFlags = value }
func (ce *ClassEntry) GetDefaultPropertiesCount() int { return ce.defaultPropertiesCount }
func (ce *ClassEntry) SetDefaultPropertiesCount(value int) {
	ce.defaultPropertiesCount = value
}
func (ce *ClassEntry) GetDefaultStaticMembersCount() int {
	return ce.defaultStaticMembersCount
}
func (ce *ClassEntry) SetDefaultStaticMembersCount(value int) {
	ce.defaultStaticMembersCount = value
}
func (ce *ClassEntry) GetDefaultPropertiesTable() []Zval {
	return ce.defaultPropertiesTable
}
func (ce *ClassEntry) SetDefaultPropertiesTable(value []Zval) {
	ce.defaultPropertiesTable = value
}
func (ce *ClassEntry) GetDefaultStaticMembersTable() *Zval {
	return ce.defaultStaticMembersTable
}
func (ce *ClassEntry) SetDefaultStaticMembersTable(value *Zval) {
	ce.defaultStaticMembersTable = value
}
func (ce *ClassEntry) GetStaticMembersTablePtr() **Zval {
	return ce.static_members_table__ptr
}
func (ce *ClassEntry) GetPropertiesInfoTable() []*PropertyInfo {
	return ce.properties_info_table
}
func (ce *ClassEntry) SetPropertiesInfoTable(value []*PropertyInfo) {
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
func (ce *ClassEntry) GetNumTraits() uint32      { return ce.num_traits }
func (ce *ClassEntry) SetNumTraits(value uint32) { ce.num_traits = value }

func (ce *ClassEntry) GetTraitNames() *ClassName                    { return ce.trait_names }
func (ce *ClassEntry) SetTraitNames(value *ClassName)               { ce.trait_names = value }
func (ce *ClassEntry) GetTraitAliases() []*zend.ZendTraitAlias      { return ce.trait_aliases }
func (ce *ClassEntry) SetTraitAliases(value []*zend.ZendTraitAlias) { ce.trait_aliases = value }
func (ce *ClassEntry) GetTraitPrecedences() []*zend.ZendTraitPrecedence {
	return ce.trait_precedences
}
func (ce *ClassEntry) SetTraitPrecedences(value []*zend.ZendTraitPrecedence) {
	ce.trait_precedences = value
}
func (ce *ClassEntry) GetFilename() string       { return ce.infoUser.filename }
func (ce *ClassEntry) SetFilename(value string)  { ce.infoUser.filename = value }
func (ce *ClassEntry) SetLineStart(value uint32) { ce.infoUser.lineStart = value }
func (ce *ClassEntry) SetLineEnd(value uint32)   { ce.infoUser.lineEnd = value }
func (ce *ClassEntry) GetDocComment() string     { return ce.infoUser.docComment }
func (ce *ClassEntry) SetDocComment(value string) {
	ce.infoUser.docComment = value
}

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
	return ce.HasCeFlags(AccConstantsUpdated)
}
func (ce ClassEntry) IsInterface() bool { return ce.HasCeFlags(AccInterface) }
func (ce ClassEntry) IsTrait() bool     { return ce.HasCeFlags(AccTrait) }
func (ce ClassEntry) IsImmutable() bool { return ce.HasCeFlags(AccImmutable) }
func (ce ClassEntry) IsResolvedParent() bool {
	return ce.HasCeFlags(AccResolvedParent)
}
func (ce ClassEntry) IsLinked() bool { return ce.HasCeFlags(AccLinked) }
func (ce ClassEntry) IsImplementTraits() bool {
	return ce.HasCeFlags(AccImplementTraits)
}
func (ce ClassEntry) IsHasStaticInMethods() bool {
	return ce.HasCeFlags(AccHasStaticInMethods)
}
func (ce ClassEntry) IsNearlyLinked() bool { return ce.HasCeFlags(AccNearlyLinked) }
func (ce ClassEntry) IsResolvedInterfaces() bool {
	return ce.HasCeFlags(AccResolvedInterfaces)
}
func (ce ClassEntry) IsFinal() bool { return ce.HasCeFlags(AccFinal) }
func (ce ClassEntry) IsImplementInterfaces() bool {
	return ce.HasCeFlags(AccImplementInterfaces)
}
func (ce ClassEntry) IsImplicitAbstractClass() bool {
	return ce.HasCeFlags(AccImplicitAbstractClass)
}
func (ce ClassEntry) IsUnresolvedVariance() bool {
	return ce.HasCeFlags(AccUnresolvedVariance)
}
func (ce ClassEntry) IsHasUnlinkedUses() bool {
	return ce.HasCeFlags(AccHasUnlinkedUses)
}
func (ce ClassEntry) IsUseGuards() bool { return ce.HasCeFlags(AccUseGuards) }
func (ce ClassEntry) IsPropertyTypesResolved() bool {
	return ce.HasCeFlags(AccPropertyTypesResolved)
}
func (ce ClassEntry) IsExplicitAbstractClass() bool {
	return ce.HasCeFlags(AccExplicitAbstractClass)
}
func (ce ClassEntry) IsHasTypeHints() bool {
	return ce.HasCeFlags(AccHasTypeHints)
}
func (ce ClassEntry) IsPreloaded() bool { return ce.HasCeFlags(AccPreloaded) }
func (ce ClassEntry) IsInherited() bool { return ce.HasCeFlags(AccInherited) }
func (ce ClassEntry) IsTopLevel() bool  { return ce.HasCeFlags(AccTopLevel) }
func (ce ClassEntry) IsReuseGetIterator() bool {
	return ce.HasCeFlags(AccReuseGetIterator)
}
func (ce *ClassEntry) SetIsConstantsUpdated(cond bool) {
	ce.SwitchCeFlags(AccConstantsUpdated, cond)
}
func (ce *ClassEntry) SetIsInterface(cond bool) {
	ce.SwitchCeFlags(AccInterface, cond)
}
func (ce *ClassEntry) SetIsTrait(cond bool) { ce.SwitchCeFlags(AccTrait, cond) }
func (ce *ClassEntry) SetIsImmutable(cond bool) {
	ce.SwitchCeFlags(AccImmutable, cond)
}
func (ce *ClassEntry) SetIsResolvedParent(cond bool) {
	ce.SwitchCeFlags(AccResolvedParent, cond)
}
func (ce *ClassEntry) SetIsLinked(cond bool) { ce.SwitchCeFlags(AccLinked, cond) }
func (ce *ClassEntry) SetIsImplementTraits(cond bool) {
	ce.SwitchCeFlags(AccImplementTraits, cond)
}
func (ce *ClassEntry) SetIsHasStaticInMethods(cond bool) {
	ce.SwitchCeFlags(AccHasStaticInMethods, cond)
}
func (ce *ClassEntry) SetIsNearlyLinked(cond bool) {
	ce.SwitchCeFlags(AccNearlyLinked, cond)
}
func (ce *ClassEntry) SetIsResolvedInterfaces(cond bool) {
	ce.SwitchCeFlags(AccResolvedInterfaces, cond)
}
func (ce *ClassEntry) SetIsFinal(cond bool) { ce.SwitchCeFlags(AccFinal, cond) }
func (ce *ClassEntry) SetIsImplementInterfaces(cond bool) {
	ce.SwitchCeFlags(AccImplementInterfaces, cond)
}
func (ce *ClassEntry) SetIsImplicitAbstractClass(cond bool) {
	ce.SwitchCeFlags(AccImplicitAbstractClass, cond)
}
func (ce *ClassEntry) SetIsUnresolvedVariance(cond bool) {
	ce.SwitchCeFlags(AccUnresolvedVariance, cond)
}
func (ce *ClassEntry) SetIsHasUnlinkedUses(cond bool) {
	ce.SwitchCeFlags(AccHasUnlinkedUses, cond)
}
func (ce *ClassEntry) SetIsUseGuards(cond bool) {
	ce.SwitchCeFlags(AccUseGuards, cond)
}
func (ce *ClassEntry) SetIsPropertyTypesResolved(cond bool) {
	ce.SwitchCeFlags(AccPropertyTypesResolved, cond)
}
func (ce *ClassEntry) SetIsExplicitAbstractClass(cond bool) {
	ce.SwitchCeFlags(AccExplicitAbstractClass, cond)
}
func (ce *ClassEntry) SetIsHasTypeHints(cond bool) {
	ce.SwitchCeFlags(AccHasTypeHints, cond)
}
func (ce *ClassEntry) SetIsPreloaded(cond bool) {
	ce.SwitchCeFlags(AccPreloaded, cond)
}
func (ce *ClassEntry) SetIsInherited(cond bool) {
	ce.SwitchCeFlags(AccInherited, cond)
}
func (ce *ClassEntry) SetIsTopLevel(cond bool) {
	ce.SwitchCeFlags(AccTopLevel, cond)
}
func (ce *ClassEntry) SetIsReuseGetIterator(cond bool) {
	ce.SwitchCeFlags(AccReuseGetIterator, cond)
}
