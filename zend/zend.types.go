// <<generate>>

package zend

import (
	r "sik/builtin/file"
	"sik/zend/types"
)

/**
 * ZendClassName
 */
type ZendClassName struct {
	name    *types.ZendString
	lc_name *types.ZendString
}

func (this *ZendClassName) GetName() *types.ZendString        { return this.name }
func (this *ZendClassName) SetName(value *types.ZendString)   { this.name = value }
func (this *ZendClassName) GetLcName() *types.ZendString      { return this.lc_name }
func (this *ZendClassName) SetLcName(value *types.ZendString) { this.lc_name = value }

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	method_name *types.ZendString
	class_name  *types.ZendString
}

func (this *ZendTraitMethodReference) GetMethodName() *types.ZendString { return this.method_name }
func (this *ZendTraitMethodReference) SetMethodName(value *types.ZendString) {
	this.method_name = value
}
func (this *ZendTraitMethodReference) GetClassName() *types.ZendString      { return this.class_name }
func (this *ZendTraitMethodReference) SetClassName(value *types.ZendString) { this.class_name = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	trait_method        ZendTraitMethodReference
	num_excludes        uint32
	exclude_class_names []*types.ZendString
}

func (this *ZendTraitPrecedence) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitPrecedence) GetNumExcludes() uint32                   { return this.num_excludes }
func (this *ZendTraitPrecedence) SetNumExcludes(value uint32)              { this.num_excludes = value }
func (this *ZendTraitPrecedence) GetExcludeClassNames() []*types.ZendString {
	return this.exclude_class_names
}

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	trait_method ZendTraitMethodReference
	alias        *types.ZendString
	modifiers    uint32
}

func (this *ZendTraitAlias) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitAlias) GetAlias() *types.ZendString              { return this.alias }
func (this *ZendTraitAlias) SetAlias(value *types.ZendString)         { this.alias = value }
func (this *ZendTraitAlias) GetModifiers() uint32                     { return this.modifiers }
func (this *ZendTraitAlias) SetModifiers(value uint32)                { this.modifiers = value }

type ZendUtilityFunctions struct {
	ErrorFunction             func(type_ int, error_filename string, error_lineno uint32, format string, args ...any)
	PrintfFunction            func(format string, args ...any) int
	WriteFunction             func(str string) int
	FopenFunction             func(filename string, opened_path *string) *r.FILE
	MessageHandler            func(message ZendLong, data any)
	GetConfigurationDirective func(name string) *types.Zval
	TicksFunction             func(ticks int)
	OnTimeout                 func(seconds int)
	StreamOpenFunction        func(filename string, handle *ZendFileHandle) int
	PrintfToSmartStrFunction  func(buf *SmartStr, format string, ap ...any)
	GetenvFunction            func(name string) *string
	ResolvePathFunction       func(filename string) *string
}

/**
 * ZendUtilityValues
 */
type ZendUtilityValues struct {
	html_errors types.ZendBool
}

func (this *ZendUtilityValues) GetHtmlErrors() types.ZendBool      { return this.html_errors }
func (this *ZendUtilityValues) SetHtmlErrors(value types.ZendBool) { this.html_errors = value }

/**
 * ZendErrorHandling
 */
type ZendErrorHandling struct {
	handling     ZendErrorHandlingT
	exception    *ZendClassEntry
	user_handler types.Zval
}

func (this *ZendErrorHandling) GetHandling() ZendErrorHandlingT      { return this.handling }
func (this *ZendErrorHandling) SetHandling(value ZendErrorHandlingT) { this.handling = value }
func (this *ZendErrorHandling) GetException() *ZendClassEntry        { return this.exception }
func (this *ZendErrorHandling) SetException(value *ZendClassEntry)   { this.exception = value }
func (this *ZendErrorHandling) GetUserHandler() types.Zval           { return this.user_handler }
