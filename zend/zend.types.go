package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * ZendClassName
 */
type ZendClassName struct {
	name   *types2.String
	lcName *types2.String
}

func (this *ZendClassName) GetName() *types2.String        { return this.name }
func (this *ZendClassName) SetName(value *types2.String)   { this.name = value }
func (this *ZendClassName) GetLcName() *types2.String      { return this.lcName }
func (this *ZendClassName) SetLcName(value *types2.String) { this.lcName = value }

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	method_name *types2.String
	class_name  *types2.String
}

func (this *ZendTraitMethodReference) GetMethodName() *types2.String { return this.method_name }
func (this *ZendTraitMethodReference) SetMethodName(value *types2.String) {
	this.method_name = value
}
func (this *ZendTraitMethodReference) GetClassName() *types2.String      { return this.class_name }
func (this *ZendTraitMethodReference) SetClassName(value *types2.String) { this.class_name = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	trait_method        ZendTraitMethodReference
	num_excludes        uint32
	exclude_class_names []*types2.String
}

func (this *ZendTraitPrecedence) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitPrecedence) GetNumExcludes() uint32                   { return this.num_excludes }
func (this *ZendTraitPrecedence) SetNumExcludes(value uint32)              { this.num_excludes = value }
func (this *ZendTraitPrecedence) GetExcludeClassNames() []*types2.String {
	return this.exclude_class_names
}

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	trait_method ZendTraitMethodReference
	alias        *types2.String
	modifiers    uint32
}

func (this *ZendTraitAlias) GetTraitMethod() ZendTraitMethodReference { return this.trait_method }
func (this *ZendTraitAlias) GetAlias() *types2.String                 { return this.alias }
func (this *ZendTraitAlias) SetAlias(value *types2.String)            { this.alias = value }
func (this *ZendTraitAlias) GetModifiers() uint32                     { return this.modifiers }
func (this *ZendTraitAlias) SetModifiers(value uint32)                { this.modifiers = value }

/**
 * ZendUtilityValues
 */
type ZendUtilityValues struct {
	html_errors types2.ZendBool
}

func (this *ZendUtilityValues) GetHtmlErrors() types2.ZendBool      { return this.html_errors }
func (this *ZendUtilityValues) SetHtmlErrors(value types2.ZendBool) { this.html_errors = value }

/**
 * ZendErrorHandling
 */
type ZendErrorHandling struct {
	handling     ZendErrorHandlingT
	exception    *types2.ClassEntry
	user_handler types2.Zval
}

func (this *ZendErrorHandling) GetHandling() ZendErrorHandlingT       { return this.handling }
func (this *ZendErrorHandling) SetHandling(value ZendErrorHandlingT)  { this.handling = value }
func (this *ZendErrorHandling) GetException() *types2.ClassEntry      { return this.exception }
func (this *ZendErrorHandling) SetException(value *types2.ClassEntry) { this.exception = value }
func (this *ZendErrorHandling) GetUserHandler() types2.Zval           { return this.user_handler }
