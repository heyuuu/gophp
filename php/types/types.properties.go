package types

// properties for Class
func (ce *Class) Type() byte {
	return ce.typ
}
func (ce *Class) Flags() uint32 {
	return ce.flags
}
func (ce *Class) SetFlags(v uint32) {
	ce.flags = v
}
func (ce *Class) ParentName() string {
	return ce.parentName
}
func (ce *Class) SetParentName(v string) {
	ce.parentName = v
}
func (ce *Class) Parent() *Class {
	return ce.parent
}
func (ce *Class) SetParent(v *Class) {
	ce.parent = v
}
func (ce *Class) InterfaceNames() []ClassName {
	return ce.interfaceNames
}
func (ce *Class) Interfaces() []*Class {
	return ce.interfaces
}
func (ce *Class) FunctionTable() *FunctionTable {
	return ce.functionTable
}
func (ce *Class) PropertyTable() *PropertyInfoTable {
	return ce.propertyTable
}
func (ce *Class) ConstantTable() *ClassConstantTable {
	return ce.constantTable
}
func (ce *Class) GetConstructor() *Function {
	return ce.constructor
}
func (ce *Class) SetConstructor(v *Function) {
	ce.constructor = v
}
func (ce *Class) GetDestructor() *Function {
	return ce.destructor
}
func (ce *Class) SetDestructor(v *Function) {
	ce.destructor = v
}
func (ce *Class) GetClone() *Function {
	return ce.clone
}
func (ce *Class) SetClone(v *Function) {
	ce.clone = v
}
func (ce *Class) GetGet() *Function {
	return ce.__get
}
func (ce *Class) SetGet(v *Function) {
	ce.__get = v
}
func (ce *Class) GetSet() *Function {
	return ce.__set
}
func (ce *Class) SetSet(v *Function) {
	ce.__set = v
}
func (ce *Class) GetUnset() *Function {
	return ce.__unset
}
func (ce *Class) SetUnset(v *Function) {
	ce.__unset = v
}
func (ce *Class) GetIsset() *Function {
	return ce.__isset
}
func (ce *Class) SetIsset(v *Function) {
	ce.__isset = v
}
func (ce *Class) GetCall() *Function {
	return ce.__call
}
func (ce *Class) SetCall(v *Function) {
	ce.__call = v
}
func (ce *Class) GetCallstatic() *Function {
	return ce.__callstatic
}
func (ce *Class) SetCallstatic(v *Function) {
	ce.__callstatic = v
}
func (ce *Class) GetTostring() *Function {
	return ce.__tostring
}
func (ce *Class) SetTostring(v *Function) {
	ce.__tostring = v
}
func (ce *Class) GetDebugInfo() *Function {
	return ce.__debugInfo
}
func (ce *Class) SetDebugInfo(v *Function) {
	ce.__debugInfo = v
}
func (ce *Class) GetSerializeFunc() *Function {
	return ce.serializeFunc
}
func (ce *Class) SetSerializeFunc(v *Function) {
	ce.serializeFunc = v
}
func (ce *Class) GetUnserializeFunc() *Function {
	return ce.unserializeFunc
}
func (ce *Class) SetUnserializeFunc(v *Function) {
	ce.unserializeFunc = v
}
func (ce *Class) ModuleNumber() int {
	return ce.moduleNumber
}

// properties for ClassConstant
func (c *ClassConstant) Name() string {
	return c.name
}
func (c *ClassConstant) Value() Zval {
	return c.value
}
func (c *ClassConstant) DocComment() string {
	return c.docComment
}
func (c *ClassConstant) Ce() *Class {
	return c.ce
}
func (c *ClassConstant) Flags() uint32 {
	return c.flags
}
