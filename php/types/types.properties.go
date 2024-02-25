package types

import "github.com/heyuuu/gophp/compile/ast"

// properties for ArgInfo
func (t *ArgInfo) Name() string {
	return t.name
}
func (t *ArgInfo) ByRef() bool {
	return t.byRef
}
func (t *ArgInfo) Variadic() bool {
	return t.variadic
}

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
func (ce *Class) DefaultPropertiesCount() int {
	return ce.defaultPropertiesCount
}
func (ce *Class) DefaultPropertiesTable() []Zval {
	return ce.defaultPropertiesTable
}
func (ce *Class) DefaultStaticMembersCount() int {
	return ce.defaultStaticMembersCount
}
func (ce *Class) DefaultStaticMembersTable() []Zval {
	return ce.defaultStaticMembersTable
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
func (ce *Class) GetCreateObject() func(classType *Class) *Object {
	return ce.createObject
}
func (ce *Class) SetCreateObject(v func(classType *Class) *Object) {
	ce.createObject = v
}
func (ce *Class) GetInterfaceGetsImplemented() func(iface *Class, classType *Class) int {
	return ce.interfaceGetsImplemented
}
func (ce *Class) SetInterfaceGetsImplemented(v func(iface *Class, classType *Class) int) {
	ce.interfaceGetsImplemented = v
}
func (ce *Class) SetSerialize(v func(object *Zval) (string, bool)) {
	ce.serialize = v
}
func (ce *Class) SetUnserialize(v func(ce *Class, data string) (*Zval, bool)) {
	ce.unserialize = v
}
func (ce *Class) ModuleNumber() int {
	return ce.moduleNumber
}

// properties for ClassConstant
func (t *ClassConstant) Name() string {
	return t.name
}
func (t *ClassConstant) Value() Zval {
	return t.value
}
func (t *ClassConstant) DocComment() string {
	return t.docComment
}
func (t *ClassConstant) Ce() *Class {
	return t.ce
}
func (t *ClassConstant) Flags() uint32 {
	return t.flags
}

// properties for ClosureData
func (t *ClosureData) Ce() *Class {
	return t.ce
}
func (t *ClosureData) SetCe(v *Class) {
	t.ce = v
}
func (t *ClosureData) Fn() *Function {
	return t.fn
}
func (t *ClosureData) SetFn(v *Function) {
	t.fn = v
}
func (t *ClosureData) Obj() *Object {
	return t.obj
}
func (t *ClosureData) SetObj(v *Object) {
	t.obj = v
}

// properties for Function
func (f *Function) Type() FunctionType {
	return f.typ
}
func (f *Function) Flags() uint32 {
	return f.flags
}
func (f *Function) SetFlags(v uint32) {
	f.flags = v
}
func (f *Function) Name() string {
	return f.name
}
func (f *Function) SetName(v string) {
	f.name = v
}
func (f *Function) Scope() *Class {
	return f.scope
}
func (f *Function) SetScope(v *Class) {
	f.scope = v
}
func (f *Function) ArgInfos() []ArgInfo {
	return f.argInfos
}
func (f *Function) Handler() any {
	return f.handler
}
func (f *Function) Stmts() []ast.Stmt {
	return f.stmts
}

// properties for PropertyInfo
func (prop *PropertyInfo) Offset() uint32 {
	return prop.offset
}
func (prop *PropertyInfo) SetOffset(v uint32) {
	prop.offset = v
}
func (prop *PropertyInfo) Flags() uint32 {
	return prop.flags
}
func (prop *PropertyInfo) Name() string {
	return prop.name
}
func (prop *PropertyInfo) DocComment() string {
	return prop.docComment
}
func (prop *PropertyInfo) Ce() *Class {
	return prop.ce
}
func (prop *PropertyInfo) Type() *TypeHint {
	return prop.typ
}

// properties for Resource
func (res *Resource) Handle() int {
	return res.handle
}
func (res *Resource) Ptr() any {
	return res.ptr
}
func (res *Resource) Type() ResourceType {
	return res.typ
}

// properties for blockInfo
func (t *blockInfo) Filename() string {
	return t.filename
}
func (t *blockInfo) SetFilename(v string) {
	t.filename = v
}
func (t *blockInfo) LineStart() uint32 {
	return t.lineStart
}
func (t *blockInfo) SetLineStart(v uint32) {
	t.lineStart = v
}
func (t *blockInfo) LineEnd() uint32 {
	return t.lineEnd
}
func (t *blockInfo) SetLineEnd(v uint32) {
	t.lineEnd = v
}
func (t *blockInfo) DocComment() string {
	return t.docComment
}
func (t *blockInfo) SetDocComment(v string) {
	t.docComment = v
}
