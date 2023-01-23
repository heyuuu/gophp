// <<generate>>

package zend

import (
	r "sik/runtime"
)

/**
 * ZendClassName
 */
type ZendClassName struct {
	name    *ZendString
	lc_name *ZendString
}

func (this ZendClassName) GetName() *ZendString         { return this.name }
func (this *ZendClassName) SetName(value *ZendString)   { this.name = value }
func (this ZendClassName) GetLcName() *ZendString       { return this.lc_name }
func (this *ZendClassName) SetLcName(value *ZendString) { this.lc_name = value }

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	method_name *ZendString
	class_name  *ZendString
}

func (this ZendTraitMethodReference) GetMethodName() *ZendString       { return this.method_name }
func (this *ZendTraitMethodReference) SetMethodName(value *ZendString) { this.method_name = value }
func (this ZendTraitMethodReference) GetClassName() *ZendString        { return this.class_name }
func (this *ZendTraitMethodReference) SetClassName(value *ZendString)  { this.class_name = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	trait_method        ZendTraitMethodReference
	num_excludes        uint32
	exclude_class_names []*ZendString
}

func (this ZendTraitPrecedence) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitPrecedence) SetTraitMethod(value ZendTraitMethodReference) {
	this.trait_method = value
}
func (this ZendTraitPrecedence) GetNumExcludes() uint32              { return this.num_excludes }
func (this *ZendTraitPrecedence) SetNumExcludes(value uint32)        { this.num_excludes = value }
func (this ZendTraitPrecedence) GetExcludeClassNames() []*ZendString { return this.exclude_class_names }
func (this *ZendTraitPrecedence) SetExcludeClassNames(value []*ZendString) {
	this.exclude_class_names = value
}

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	trait_method ZendTraitMethodReference
	alias        *ZendString
	modifiers    uint32
}

func (this ZendTraitAlias) GetTraitMethod() ZendTraitMethodReference       { return this.trait_method }
func (this *ZendTraitAlias) SetTraitMethod(value ZendTraitMethodReference) { this.trait_method = value }
func (this ZendTraitAlias) GetAlias() *ZendString                          { return this.alias }
func (this *ZendTraitAlias) SetAlias(value *ZendString)                    { this.alias = value }
func (this ZendTraitAlias) GetModifiers() uint32                           { return this.modifiers }
func (this *ZendTraitAlias) SetModifiers(value uint32)                     { this.modifiers = value }

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

func (this ZendClassEntry) GetType() byte                    { return this.type_ }
func (this *ZendClassEntry) SetType(value byte)              { this.type_ = value }
func (this ZendClassEntry) GetName() *ZendString             { return this.name }
func (this *ZendClassEntry) SetName(value *ZendString)       { this.name = value }
func (this ZendClassEntry) GetParent() *ZendClassEntry       { return this.__0.parent }
func (this *ZendClassEntry) SetParent(value *ZendClassEntry) { this.__0.parent = value }
func (this ZendClassEntry) GetParentName() *ZendString       { return this.__0.parent_name }
func (this *ZendClassEntry) SetParentName(value *ZendString) { this.__0.parent_name = value }
func (this ZendClassEntry) GetRefcount() int                 { return this.refcount }
func (this *ZendClassEntry) SetRefcount(value int)           { this.refcount = value }
func (this ZendClassEntry) GetCeFlags() uint32               { return this.ce_flags }
func (this *ZendClassEntry) SetCeFlags(value uint32)         { this.ce_flags = value }
func (this ZendClassEntry) GetDefaultPropertiesCount() int   { return this.default_properties_count }
func (this *ZendClassEntry) SetDefaultPropertiesCount(value int) {
	this.default_properties_count = value
}
func (this ZendClassEntry) GetDefaultStaticMembersCount() int {
	return this.default_static_members_count
}
func (this *ZendClassEntry) SetDefaultStaticMembersCount(value int) {
	this.default_static_members_count = value
}
func (this ZendClassEntry) GetDefaultPropertiesTable() *Zval { return this.default_properties_table }
func (this *ZendClassEntry) SetDefaultPropertiesTable(value *Zval) {
	this.default_properties_table = value
}
func (this ZendClassEntry) GetDefaultStaticMembersTable() *Zval {
	return this.default_static_members_table
}
func (this *ZendClassEntry) SetDefaultStaticMembersTable(value *Zval) {
	this.default_static_members_table = value
}
func (this ZendClassEntry) GetStaticMembersTablePtr() **Zval { return this.static_members_table__ptr }
func (this *ZendClassEntry) SetStaticMembersTablePtr(value **Zval) {
	this.static_members_table__ptr = value
}
func (this ZendClassEntry) GetFunctionTable() HashTable        { return this.function_table }
func (this *ZendClassEntry) SetFunctionTable(value HashTable)  { this.function_table = value }
func (this ZendClassEntry) GetPropertiesInfo() HashTable       { return this.properties_info }
func (this *ZendClassEntry) SetPropertiesInfo(value HashTable) { this.properties_info = value }
func (this ZendClassEntry) GetConstantsTable() HashTable       { return this.constants_table }
func (this *ZendClassEntry) SetConstantsTable(value HashTable) { this.constants_table = value }
func (this ZendClassEntry) GetPropertiesInfoTable() **ZendPropertyInfo {
	return this.properties_info_table
}
func (this *ZendClassEntry) SetPropertiesInfoTable(value **ZendPropertyInfo) {
	this.properties_info_table = value
}
func (this ZendClassEntry) GetConstructor() *ZendFunction           { return this.constructor }
func (this *ZendClassEntry) SetConstructor(value *ZendFunction)     { this.constructor = value }
func (this ZendClassEntry) GetDestructor() *ZendFunction            { return this.destructor }
func (this *ZendClassEntry) SetDestructor(value *ZendFunction)      { this.destructor = value }
func (this ZendClassEntry) GetClone() *ZendFunction                 { return this.clone }
func (this *ZendClassEntry) SetClone(value *ZendFunction)           { this.clone = value }
func (this ZendClassEntry) GetGet() *ZendFunction                   { return this.__get }
func (this *ZendClassEntry) SetGet(value *ZendFunction)             { this.__get = value }
func (this ZendClassEntry) GetSet() *ZendFunction                   { return this.__set }
func (this *ZendClassEntry) SetSet(value *ZendFunction)             { this.__set = value }
func (this ZendClassEntry) GetUnset() *ZendFunction                 { return this.__unset }
func (this *ZendClassEntry) SetUnset(value *ZendFunction)           { this.__unset = value }
func (this ZendClassEntry) GetIsset() *ZendFunction                 { return this.__isset }
func (this *ZendClassEntry) SetIsset(value *ZendFunction)           { this.__isset = value }
func (this ZendClassEntry) GetCall() *ZendFunction                  { return this.__call }
func (this *ZendClassEntry) SetCall(value *ZendFunction)            { this.__call = value }
func (this ZendClassEntry) GetCallstatic() *ZendFunction            { return this.__callstatic }
func (this *ZendClassEntry) SetCallstatic(value *ZendFunction)      { this.__callstatic = value }
func (this ZendClassEntry) GetTostring() *ZendFunction              { return this.__tostring }
func (this *ZendClassEntry) SetTostring(value *ZendFunction)        { this.__tostring = value }
func (this ZendClassEntry) GetDebugInfo() *ZendFunction             { return this.__debugInfo }
func (this *ZendClassEntry) SetDebugInfo(value *ZendFunction)       { this.__debugInfo = value }
func (this ZendClassEntry) GetSerializeFunc() *ZendFunction         { return this.serialize_func }
func (this *ZendClassEntry) SetSerializeFunc(value *ZendFunction)   { this.serialize_func = value }
func (this ZendClassEntry) GetUnserializeFunc() *ZendFunction       { return this.unserialize_func }
func (this *ZendClassEntry) SetUnserializeFunc(value *ZendFunction) { this.unserialize_func = value }
func (this ZendClassEntry) GetIteratorFuncsPtr() *ZendClassIteratorFuncs {
	return this.iterator_funcs_ptr
}
func (this *ZendClassEntry) SetIteratorFuncsPtr(value *ZendClassIteratorFuncs) {
	this.iterator_funcs_ptr = value
}
func (this ZendClassEntry) GetCreateObject() func(class_type *ZendClassEntry) *ZendObject {
	return this.__1.create_object
}
func (this *ZendClassEntry) SetCreateObject(value func(class_type *ZendClassEntry) *ZendObject) {
	this.__1.create_object = value
}
func (this ZendClassEntry) GetInterfaceGetsImplemented() func(iface *ZendClassEntry, class_type *ZendClassEntry) int {
	return this.__1.interface_gets_implemented
}
func (this *ZendClassEntry) SetInterfaceGetsImplemented(value func(iface *ZendClassEntry, class_type *ZendClassEntry) int) {
	this.__1.interface_gets_implemented = value
}
func (this ZendClassEntry) GetGetIterator() func(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator {
	return this.get_iterator
}
func (this *ZendClassEntry) SetGetIterator(value func(ce *ZendClassEntry, object *Zval, by_ref int) *ZendObjectIterator) {
	this.get_iterator = value
}
func (this ZendClassEntry) GetGetStaticMethod() func(ce *ZendClassEntry, method *ZendString) *ZendFunction {
	return this.get_static_method
}
func (this *ZendClassEntry) SetGetStaticMethod(value func(ce *ZendClassEntry, method *ZendString) *ZendFunction) {
	this.get_static_method = value
}
func (this ZendClassEntry) GetSerialize() func(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int {
	return this.serialize
}
func (this *ZendClassEntry) SetSerialize(value func(object *Zval, buffer **uint8, buf_len *int, data *ZendSerializeData) int) {
	this.serialize = value
}
func (this ZendClassEntry) GetUnserialize() func(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int {
	return this.unserialize
}
func (this *ZendClassEntry) SetUnserialize(value func(object *Zval, ce *ZendClassEntry, buf *uint8, buf_len int, data *ZendUnserializeData) int) {
	this.unserialize = value
}
func (this ZendClassEntry) GetNumInterfaces() uint32                   { return this.num_interfaces }
func (this *ZendClassEntry) SetNumInterfaces(value uint32)             { this.num_interfaces = value }
func (this ZendClassEntry) GetNumTraits() uint32                       { return this.num_traits }
func (this *ZendClassEntry) SetNumTraits(value uint32)                 { this.num_traits = value }
func (this ZendClassEntry) GetInterfaces() **ZendClassEntry            { return this.__2.interfaces }
func (this *ZendClassEntry) SetInterfaces(value **ZendClassEntry)      { this.__2.interfaces = value }
func (this ZendClassEntry) GetInterfaceNames() *ZendClassName          { return this.__2.interface_names }
func (this *ZendClassEntry) SetInterfaceNames(value *ZendClassName)    { this.__2.interface_names = value }
func (this ZendClassEntry) GetTraitNames() *ZendClassName              { return this.trait_names }
func (this *ZendClassEntry) SetTraitNames(value *ZendClassName)        { this.trait_names = value }
func (this ZendClassEntry) GetTraitAliases() **ZendTraitAlias          { return this.trait_aliases }
func (this *ZendClassEntry) SetTraitAliases(value **ZendTraitAlias)    { this.trait_aliases = value }
func (this ZendClassEntry) GetTraitPrecedences() **ZendTraitPrecedence { return this.trait_precedences }
func (this *ZendClassEntry) SetTraitPrecedences(value **ZendTraitPrecedence) {
	this.trait_precedences = value
}
func (this ZendClassEntry) GetFilename() *ZendString         { return this.info.user.filename }
func (this *ZendClassEntry) SetFilename(value *ZendString)   { this.info.user.filename = value }
func (this ZendClassEntry) GetLineStart() uint32             { return this.info.user.line_start }
func (this *ZendClassEntry) SetLineStart(value uint32)       { this.info.user.line_start = value }
func (this ZendClassEntry) GetLineEnd() uint32               { return this.info.user.line_end }
func (this *ZendClassEntry) SetLineEnd(value uint32)         { this.info.user.line_end = value }
func (this ZendClassEntry) GetDocComment() *ZendString       { return this.info.user.doc_comment }
func (this *ZendClassEntry) SetDocComment(value *ZendString) { this.info.user.doc_comment = value }
func (this ZendClassEntry) GetBuiltinFunctions() *ZendFunctionEntry {
	return this.info.internal.builtin_functions
}
func (this *ZendClassEntry) SetBuiltinFunctions(value *ZendFunctionEntry) {
	this.info.internal.builtin_functions = value
}
func (this ZendClassEntry) GetModule() *ZendModuleEntry       { return this.info.internal.module }
func (this *ZendClassEntry) SetModule(value *ZendModuleEntry) { this.info.internal.module = value }

/**
 * ZendUtilityFunctions
 */
type ZendUtilityFunctions struct {
	error_function                  func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any)
	printf_function                 func(format *byte, _ ...any) int
	write_function                  func(str *byte, str_length int) int
	fopen_function                  func(filename *byte, opened_path **ZendString) *r.FILE
	message_handler                 func(message ZendLong, data any)
	get_configuration_directive     func(name *ZendString) *Zval
	ticks_function                  func(ticks int)
	on_timeout                      func(seconds int)
	stream_open_function            func(filename *byte, handle *ZendFileHandle) int
	printf_to_smart_string_function func(buf *SmartString, format *byte, ap ...any)
	printf_to_smart_str_function    func(buf *SmartStr, format *byte, ap ...any)
	getenv_function                 func(name *byte, name_len int) *byte
	resolve_path_function           func(filename *byte, filename_len int) *ZendString
}

func (this ZendUtilityFunctions) GetErrorFunction() func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any) {
	return this.error_function
}
func (this *ZendUtilityFunctions) SetErrorFunction(value func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any)) {
	this.error_function = value
}
func (this ZendUtilityFunctions) GetPrintfFunction() func(format *byte, _ ...any) int {
	return this.printf_function
}
func (this *ZendUtilityFunctions) SetPrintfFunction(value func(format *byte, _ ...any) int) {
	this.printf_function = value
}
func (this ZendUtilityFunctions) GetWriteFunction() func(str *byte, str_length int) int {
	return this.write_function
}
func (this *ZendUtilityFunctions) SetWriteFunction(value func(str *byte, str_length int) int) {
	this.write_function = value
}
func (this ZendUtilityFunctions) GetFopenFunction() func(filename *byte, opened_path **ZendString) *r.FILE {
	return this.fopen_function
}
func (this *ZendUtilityFunctions) SetFopenFunction(value func(filename *byte, opened_path **ZendString) *r.FILE) {
	this.fopen_function = value
}
func (this ZendUtilityFunctions) GetMessageHandler() func(message ZendLong, data any) {
	return this.message_handler
}
func (this *ZendUtilityFunctions) SetMessageHandler(value func(message ZendLong, data any)) {
	this.message_handler = value
}
func (this ZendUtilityFunctions) GetGetConfigurationDirective() func(name *ZendString) *Zval {
	return this.get_configuration_directive
}
func (this *ZendUtilityFunctions) SetGetConfigurationDirective(value func(name *ZendString) *Zval) {
	this.get_configuration_directive = value
}
func (this ZendUtilityFunctions) GetTicksFunction() func(ticks int) { return this.ticks_function }
func (this *ZendUtilityFunctions) SetTicksFunction(value func(ticks int)) {
	this.ticks_function = value
}
func (this ZendUtilityFunctions) GetOnTimeout() func(seconds int)       { return this.on_timeout }
func (this *ZendUtilityFunctions) SetOnTimeout(value func(seconds int)) { this.on_timeout = value }
func (this ZendUtilityFunctions) GetStreamOpenFunction() func(filename *byte, handle *ZendFileHandle) int {
	return this.stream_open_function
}
func (this *ZendUtilityFunctions) SetStreamOpenFunction(value func(filename *byte, handle *ZendFileHandle) int) {
	this.stream_open_function = value
}
func (this ZendUtilityFunctions) GetPrintfToSmartStringFunction() func(buf *SmartString, format *byte, ap ...any) {
	return this.printf_to_smart_string_function
}
func (this *ZendUtilityFunctions) SetPrintfToSmartStringFunction(value func(buf *SmartString, format *byte, ap ...any)) {
	this.printf_to_smart_string_function = value
}
func (this ZendUtilityFunctions) GetPrintfToSmartStrFunction() func(buf *SmartStr, format *byte, ap ...any) {
	return this.printf_to_smart_str_function
}
func (this *ZendUtilityFunctions) SetPrintfToSmartStrFunction(value func(buf *SmartStr, format *byte, ap ...any)) {
	this.printf_to_smart_str_function = value
}
func (this ZendUtilityFunctions) GetGetenvFunction() func(name *byte, name_len int) *byte {
	return this.getenv_function
}
func (this *ZendUtilityFunctions) SetGetenvFunction(value func(name *byte, name_len int) *byte) {
	this.getenv_function = value
}
func (this ZendUtilityFunctions) GetResolvePathFunction() func(filename *byte, filename_len int) *ZendString {
	return this.resolve_path_function
}
func (this *ZendUtilityFunctions) SetResolvePathFunction(value func(filename *byte, filename_len int) *ZendString) {
	this.resolve_path_function = value
}

/**
 * ZendUtilityValues
 */
type ZendUtilityValues struct {
	html_errors ZendBool
}

func (this ZendUtilityValues) GetHtmlErrors() ZendBool       { return this.html_errors }
func (this *ZendUtilityValues) SetHtmlErrors(value ZendBool) { this.html_errors = value }

/**
 * ZendErrorHandling
 */
type ZendErrorHandling struct {
	handling     ZendErrorHandlingT
	exception    *ZendClassEntry
	user_handler Zval
}

func (this ZendErrorHandling) GetHandling() ZendErrorHandlingT       { return this.handling }
func (this *ZendErrorHandling) SetHandling(value ZendErrorHandlingT) { this.handling = value }
func (this ZendErrorHandling) GetException() *ZendClassEntry         { return this.exception }
func (this *ZendErrorHandling) SetException(value *ZendClassEntry)   { this.exception = value }
func (this ZendErrorHandling) GetUserHandler() Zval                  { return this.user_handler }
func (this *ZendErrorHandling) SetUserHandler(value Zval)            { this.user_handler = value }
