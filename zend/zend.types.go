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

// func MakeZendClassName(name *ZendString, lc_name *ZendString) ZendClassName {
//     return ZendClassName{
//         name:name,
//         lc_name:lc_name,
//     }
// }
func (this *ZendClassName) GetName() *ZendString        { return this.name }
func (this *ZendClassName) SetName(value *ZendString)   { this.name = value }
func (this *ZendClassName) GetLcName() *ZendString      { return this.lc_name }
func (this *ZendClassName) SetLcName(value *ZendString) { this.lc_name = value }

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	method_name *ZendString
	class_name  *ZendString
}

// func MakeZendTraitMethodReference(method_name *ZendString, class_name *ZendString) ZendTraitMethodReference {
//     return ZendTraitMethodReference{
//         method_name:method_name,
//         class_name:class_name,
//     }
// }
func (this *ZendTraitMethodReference) GetMethodName() *ZendString      { return this.method_name }
func (this *ZendTraitMethodReference) SetMethodName(value *ZendString) { this.method_name = value }
func (this *ZendTraitMethodReference) GetClassName() *ZendString       { return this.class_name }
func (this *ZendTraitMethodReference) SetClassName(value *ZendString)  { this.class_name = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	trait_method        ZendTraitMethodReference
	num_excludes        uint32
	exclude_class_names []*ZendString
}

// func MakeZendTraitPrecedence(trait_method ZendTraitMethodReference, num_excludes uint32, exclude_class_names []*ZendString) ZendTraitPrecedence {
//     return ZendTraitPrecedence{
//         trait_method:trait_method,
//         num_excludes:num_excludes,
//         exclude_class_names:exclude_class_names,
//     }
// }
func (this *ZendTraitPrecedence) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }

// func (this *ZendTraitPrecedence) SetTraitMethod(value ZendTraitMethodReference) { this.trait_method = value }
func (this *ZendTraitPrecedence) GetNumExcludes() uint32      { return this.num_excludes }
func (this *ZendTraitPrecedence) SetNumExcludes(value uint32) { this.num_excludes = value }
func (this *ZendTraitPrecedence) GetExcludeClassNames() []*ZendString {
	return this.exclude_class_names
}

// func (this *ZendTraitPrecedence) SetExcludeClassNames(value []*ZendString) { this.exclude_class_names = value }

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	trait_method ZendTraitMethodReference
	alias        *ZendString
	modifiers    uint32
}

// func MakeZendTraitAlias(trait_method ZendTraitMethodReference, alias *ZendString, modifiers uint32) ZendTraitAlias {
//     return ZendTraitAlias{
//         trait_method:trait_method,
//         alias:alias,
//         modifiers:modifiers,
//     }
// }
func (this *ZendTraitAlias) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }

// func (this *ZendTraitAlias) SetTraitMethod(value ZendTraitMethodReference) { this.trait_method = value }
func (this *ZendTraitAlias) GetAlias() *ZendString      { return this.alias }
func (this *ZendTraitAlias) SetAlias(value *ZendString) { this.alias = value }
func (this *ZendTraitAlias) GetModifiers() uint32       { return this.modifiers }
func (this *ZendTraitAlias) SetModifiers(value uint32)  { this.modifiers = value }

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

func (this *ZendUtilityFunctions) GetErrorFunction() func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any) {
	return this.error_function
}
func (this *ZendUtilityFunctions) SetErrorFunction(value func(type_ int, error_filename *byte, error_lineno uint32, format *byte, args ...any)) {
	this.error_function = value
}
func (this *ZendUtilityFunctions) GetPrintfFunction() func(format *byte, _ ...any) int {
	return this.printf_function
}
func (this *ZendUtilityFunctions) SetPrintfFunction(value func(format *byte, _ ...any) int) {
	this.printf_function = value
}
func (this *ZendUtilityFunctions) GetWriteFunction() func(str *byte, str_length int) int {
	return this.write_function
}
func (this *ZendUtilityFunctions) SetWriteFunction(value func(str *byte, str_length int) int) {
	this.write_function = value
}
func (this *ZendUtilityFunctions) GetFopenFunction() func(filename *byte, opened_path **ZendString) *r.FILE {
	return this.fopen_function
}
func (this *ZendUtilityFunctions) SetFopenFunction(value func(filename *byte, opened_path **ZendString) *r.FILE) {
	this.fopen_function = value
}
func (this *ZendUtilityFunctions) GetMessageHandler() func(message ZendLong, data any) {
	return this.message_handler
}
func (this *ZendUtilityFunctions) SetMessageHandler(value func(message ZendLong, data any)) {
	this.message_handler = value
}
func (this *ZendUtilityFunctions) GetGetConfigurationDirective() func(name *ZendString) *Zval {
	return this.get_configuration_directive
}
func (this *ZendUtilityFunctions) SetGetConfigurationDirective(value func(name *ZendString) *Zval) {
	this.get_configuration_directive = value
}
func (this *ZendUtilityFunctions) GetTicksFunction() func(ticks int) { return this.ticks_function }
func (this *ZendUtilityFunctions) SetTicksFunction(value func(ticks int)) {
	this.ticks_function = value
}
func (this *ZendUtilityFunctions) GetOnTimeout() func(seconds int)      { return this.on_timeout }
func (this *ZendUtilityFunctions) SetOnTimeout(value func(seconds int)) { this.on_timeout = value }
func (this *ZendUtilityFunctions) GetStreamOpenFunction() func(filename *byte, handle *ZendFileHandle) int {
	return this.stream_open_function
}
func (this *ZendUtilityFunctions) SetStreamOpenFunction(value func(filename *byte, handle *ZendFileHandle) int) {
	this.stream_open_function = value
}
func (this *ZendUtilityFunctions) GetPrintfToSmartStringFunction() func(buf *SmartString, format *byte, ap ...any) {
	return this.printf_to_smart_string_function
}
func (this *ZendUtilityFunctions) SetPrintfToSmartStringFunction(value func(buf *SmartString, format *byte, ap ...any)) {
	this.printf_to_smart_string_function = value
}
func (this *ZendUtilityFunctions) GetPrintfToSmartStrFunction() func(buf *SmartStr, format *byte, ap ...any) {
	return this.printf_to_smart_str_function
}
func (this *ZendUtilityFunctions) SetPrintfToSmartStrFunction(value func(buf *SmartStr, format *byte, ap ...any)) {
	this.printf_to_smart_str_function = value
}
func (this *ZendUtilityFunctions) GetGetenvFunction() func(name *byte, name_len int) *byte {
	return this.getenv_function
}
func (this *ZendUtilityFunctions) SetGetenvFunction(value func(name *byte, name_len int) *byte) {
	this.getenv_function = value
}
func (this *ZendUtilityFunctions) GetResolvePathFunction() func(filename *byte, filename_len int) *ZendString {
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

// func MakeZendUtilityValues(html_errors ZendBool) ZendUtilityValues {
//     return ZendUtilityValues{
//         html_errors:html_errors,
//     }
// }
func (this *ZendUtilityValues) GetHtmlErrors() ZendBool      { return this.html_errors }
func (this *ZendUtilityValues) SetHtmlErrors(value ZendBool) { this.html_errors = value }

/**
 * ZendErrorHandling
 */
type ZendErrorHandling struct {
	handling     ZendErrorHandlingT
	exception    *ZendClassEntry
	user_handler Zval
}

// func MakeZendErrorHandling(handling ZendErrorHandlingT, exception *ZendClassEntry, user_handler Zval) ZendErrorHandling {
//     return ZendErrorHandling{
//         handling:handling,
//         exception:exception,
//         user_handler:user_handler,
//     }
// }
func (this *ZendErrorHandling) GetHandling() ZendErrorHandlingT      { return this.handling }
func (this *ZendErrorHandling) SetHandling(value ZendErrorHandlingT) { this.handling = value }
func (this *ZendErrorHandling) GetException() *ZendClassEntry        { return this.exception }
func (this *ZendErrorHandling) SetException(value *ZendClassEntry)   { this.exception = value }
func (this *ZendErrorHandling) GetUserHandler() Zval                 { return this.user_handler }

// func (this *ZendErrorHandling) SetUserHandler(value Zval) { this.user_handler = value }
