package types

import (
	"github.com/heyuuu/gophp/kits/ascii"
)

const (
	typeInternalClass = 1
	typeUserClass     = 2
)

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

func (n ClassName) Name() string   { return n.name }
func (n ClassName) LcName() string { return n.lcName }

// Class
type Class struct {
	typ   byte `get:"Type"`
	name  ClassName
	flags uint32 `prop:""`

	// 继承父类，在 link 前只有 parentName 可能有值，在 link 后只有 parent 可能有值(union)。
	parentName string `prop:""`
	parent     *Class `prop:""`

	// 继承接口列表，在 link 前只有 interfaceNames 可能有值，在 link 后只有 interfaces 可能有值(union)。
	interfaceNames []ClassName `get:""`
	interfaces     []*Class    `get:""`

	functionTable *FunctionTable      `get:""`
	propertyTable *PropertyInfoTable  `get:""`
	constantTable *ClassConstantTable `get:""`

	// 魔术方法
	constructor     *Function `prop:"@"`
	destructor      *Function `prop:"@"`
	clone           *Function `prop:"@"`
	__get           *Function `prop:"@"`
	__set           *Function `prop:"@"`
	__unset         *Function `prop:"@"`
	__isset         *Function `prop:"@"`
	__call          *Function `prop:"@"`
	__callstatic    *Function `prop:"@"`
	__tostring      *Function `prop:"@"`
	__debugInfo     *Function `prop:"@"`
	serializeFunc   *Function `prop:"@"`
	unserializeFunc *Function `prop:"@"`

	// info.internal
	moduleNumber int `get:""`
	// info.user
	blockInfo
}

func NewUserClass(entry *UserClassEntry) *Class {
	ce := initClass(typeUserClass, entry.Name)
	for _, constant := range entry.Constants {
		constant.ce = ce
		ce.ConstantTable().Add(ascii.StrToLower(constant.name), constant)
	}
	for _, property := range entry.Properties {
		property.ce = ce
		ce.PropertyTable().Add(ascii.StrToLower(property.name), property)
	}

	return ce
}

func NewInternalClass(name string, moduleNumber int, flags uint32) *Class {
	ce := initClass(typeUserClass, name)
	ce.moduleNumber = moduleNumber
	ce.flags = flags
	return ce
}

func initClass(typ byte, name string) *Class {
	ce := &Class{
		typ:           typ,
		name:          MakeClassName(name),
		functionTable: NewFunctionTable(),
		propertyTable: NewPropertyInfoTable(),
		constantTable: NewClassConstantTable(),
	}
	return ce
}

func (ce *Class) Name() string   { return ce.name.Name() }
func (ce *Class) LcName() string { return ce.name.LcName() }

/* ClassEntry.flags */
func (ce *Class) AddFlags(value uint32)      { ce.flags |= value }
func (ce *Class) HasFlags(value uint32) bool { return ce.flags&value == value }
func (ce *Class) SwitchFlags(value uint32, cond bool) {
	if cond {
		ce.flags |= value
	} else {
		ce.flags &^= value
	}
}
func (ce *Class) IsConstantsUpdated() bool      { return ce.HasFlags(AccConstantsUpdated) }
func (ce *Class) IsInterface() bool             { return ce.HasFlags(AccInterface) }
func (ce *Class) IsTrait() bool                 { return ce.HasFlags(AccTrait) }
func (ce *Class) IsImmutable() bool             { return ce.HasFlags(AccImmutable) }
func (ce *Class) IsResolvedParent() bool        { return ce.HasFlags(AccResolvedParent) }
func (ce *Class) IsLinked() bool                { return ce.HasFlags(AccLinked) }
func (ce *Class) IsImplementTraits() bool       { return ce.HasFlags(AccImplementTraits) }
func (ce *Class) IsHasStaticInMethods() bool    { return ce.HasFlags(AccHasStaticInMethods) }
func (ce *Class) IsNearlyLinked() bool          { return ce.HasFlags(AccNearlyLinked) }
func (ce *Class) IsResolvedInterfaces() bool    { return ce.HasFlags(AccResolvedInterfaces) }
func (ce *Class) IsFinal() bool                 { return ce.HasFlags(AccFinal) }
func (ce *Class) IsImplementInterfaces() bool   { return ce.HasFlags(AccImplementInterfaces) }
func (ce *Class) IsImplicitAbstractClass() bool { return ce.HasFlags(AccImplicitAbstractClass) }
func (ce *Class) IsUnresolvedVariance() bool    { return ce.HasFlags(AccUnresolvedVariance) }
func (ce *Class) IsHasUnlinkedUses() bool       { return ce.HasFlags(AccHasUnlinkedUses) }
func (ce *Class) IsUseGuards() bool             { return ce.HasFlags(AccUseGuards) }
func (ce *Class) IsPropertyTypesResolved() bool { return ce.HasFlags(AccPropertyTypesResolved) }
func (ce *Class) IsExplicitAbstractClass() bool { return ce.HasFlags(AccExplicitAbstractClass) }
func (ce *Class) IsHasTypeHints() bool          { return ce.HasFlags(AccHasTypeHints) }
func (ce *Class) IsPreloaded() bool             { return ce.HasFlags(AccPreloaded) }
func (ce *Class) IsInherited() bool             { return ce.HasFlags(AccInherited) }
func (ce *Class) IsTopLevel() bool              { return ce.HasFlags(AccTopLevel) }
func (ce *Class) IsReuseGetIterator() bool      { return ce.HasFlags(AccReuseGetIterator) }

func (ce *Class) SetIsConstantsUpdated(cond bool)    { ce.SwitchFlags(AccConstantsUpdated, cond) }
func (ce *Class) SetIsResolvedParent(cond bool)      { ce.SwitchFlags(AccResolvedParent, cond) }
func (ce *Class) SetIsLinked(cond bool)              { ce.SwitchFlags(AccLinked, cond) }
func (ce *Class) SetIsImplementTraits(cond bool)     { ce.SwitchFlags(AccImplementTraits, cond) }
func (ce *Class) SetIsHasStaticInMethods(cond bool)  { ce.SwitchFlags(AccHasStaticInMethods, cond) }
func (ce *Class) SetIsNearlyLinked(cond bool)        { ce.SwitchFlags(AccNearlyLinked, cond) }
func (ce *Class) SetIsResolvedInterfaces(cond bool)  { ce.SwitchFlags(AccResolvedInterfaces, cond) }
func (ce *Class) SetIsFinal(cond bool)               { ce.SwitchFlags(AccFinal, cond) }
func (ce *Class) SetIsImplementInterfaces(cond bool) { ce.SwitchFlags(AccImplementInterfaces, cond) }
func (ce *Class) SetIsImplicitAbstractClass(cond bool) {
	ce.SwitchFlags(AccImplicitAbstractClass, cond)
}
func (ce *Class) SetIsUnresolvedVariance(cond bool) { ce.SwitchFlags(AccUnresolvedVariance, cond) }
func (ce *Class) SetIsHasUnlinkedUses(cond bool)    { ce.SwitchFlags(AccHasUnlinkedUses, cond) }
func (ce *Class) SetIsUseGuards(cond bool)          { ce.SwitchFlags(AccUseGuards, cond) }
func (ce *Class) SetIsPropertyTypesResolved(cond bool) {
	ce.SwitchFlags(AccPropertyTypesResolved, cond)
}
func (ce *Class) SetIsExplicitAbstractClass(cond bool) {
	ce.SwitchFlags(AccExplicitAbstractClass, cond)
}
func (ce *Class) SetIsHasTypeHints(cond bool)     { ce.SwitchFlags(AccHasTypeHints, cond) }
func (ce *Class) SetIsPreloaded(cond bool)        { ce.SwitchFlags(AccPreloaded, cond) }
func (ce *Class) SetIsInherited(cond bool)        { ce.SwitchFlags(AccInherited, cond) }
func (ce *Class) SetIsTopLevel(cond bool)         { ce.SwitchFlags(AccTopLevel, cond) }
func (ce *Class) SetIsReuseGetIterator(cond bool) { ce.SwitchFlags(AccReuseGetIterator, cond) }
