package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendClassName
 */
type ZendClassName struct {
	name   *types.String
	lcName *types.String
}

func (this *ZendClassName) GetName() *types.String        { return this.name }
func (this *ZendClassName) SetName(value *types.String)   { this.name = value }
func (this *ZendClassName) GetLcName() *types.String      { return this.lcName }
func (this *ZendClassName) SetLcName(value *types.String) { this.lcName = value }

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	method_name *types.String
	class_name  *types.String
}

func (this *ZendTraitMethodReference) GetMethodName() *types.String { return this.method_name }
func (this *ZendTraitMethodReference) SetMethodName(value *types.String) {
	this.method_name = value
}
func (this *ZendTraitMethodReference) GetClassName() *types.String      { return this.class_name }
func (this *ZendTraitMethodReference) SetClassName(value *types.String) { this.class_name = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	trait_method        ZendTraitMethodReference
	num_excludes        uint32
	exclude_class_names []*types.String
}

func (this *ZendTraitPrecedence) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitPrecedence) GetNumExcludes() uint32                   { return this.num_excludes }
func (this *ZendTraitPrecedence) SetNumExcludes(value uint32)              { this.num_excludes = value }
func (this *ZendTraitPrecedence) GetExcludeClassNames() []*types.String {
	return this.exclude_class_names
}

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	trait_method ZendTraitMethodReference
	alias        *types.String
	modifiers    uint32
}

func (this *ZendTraitAlias) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitAlias) GetAlias() *types.String                  { return this.alias }
func (this *ZendTraitAlias) SetAlias(value *types.String)             { this.alias = value }
func (this *ZendTraitAlias) GetModifiers() uint32                     { return this.modifiers }
func (this *ZendTraitAlias) SetModifiers(value uint32)                { this.modifiers = value }

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
	exception    *types.ClassEntry
	user_handler types.Zval
}

func (this *ZendErrorHandling) GetHandling() ZendErrorHandlingT      { return this.handling }
func (this *ZendErrorHandling) SetHandling(value ZendErrorHandlingT) { this.handling = value }
func (this *ZendErrorHandling) GetException() *types.ClassEntry      { return this.exception }
func (this *ZendErrorHandling) SetException(value *types.ClassEntry) { this.exception = value }
func (this *ZendErrorHandling) GetUserHandler() types.Zval           { return this.user_handler }
