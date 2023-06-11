package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendTraitMethodReference
 */
type ZendTraitMethodReference struct {
	methodName string
	className  string
}

func NewTraitMethodRef(methodName string, className string) *ZendTraitMethodReference {
	return &ZendTraitMethodReference{methodName: methodName, className: className}
}

func (this *ZendTraitMethodReference) Init(methodName string, className string) {
	this.methodName = methodName
	this.className = className
}

func (this *ZendTraitMethodReference) MethodName() string        { return this.methodName }
func (this *ZendTraitMethodReference) ClassName() string         { return this.className }
func (this *ZendTraitMethodReference) SetClassName(value string) { this.className = value }

/**
 * ZendTraitPrecedence
 */
type ZendTraitPrecedence struct {
	traitMethod       ZendTraitMethodReference
	excludeClassNames []string
}

func NewTraitPrecedence(traitMethod *ZendTraitMethodReference, excludeClassNames []string) *ZendTraitPrecedence {
	return &ZendTraitPrecedence{traitMethod: *traitMethod, excludeClassNames: excludeClassNames}
}
func (this *ZendTraitPrecedence) GetTraitMethod() *ZendTraitMethodReference {
	return &this.traitMethod
}
func (this *ZendTraitPrecedence) GetNumExcludes() uint32         { return uint32(len(this.excludeClassNames)) }
func (this *ZendTraitPrecedence) GetExcludeClassNames() []string { return this.excludeClassNames }

/**
 * ZendTraitAlias
 */
type ZendTraitAlias struct {
	traitMethod ZendTraitMethodReference
	alias       string
	modifiers   uint32
}

func NewTraitAlias(traitMethod *ZendTraitMethodReference, alias string, modifiers uint32) *ZendTraitAlias {
	return &ZendTraitAlias{traitMethod: *traitMethod, alias: alias, modifiers: modifiers}
}

func (this *ZendTraitAlias) GetTraitMethod() *ZendTraitMethodReference { return &this.traitMethod }
func (this *ZendTraitAlias) GetAlias() string                          { return this.alias }
func (this *ZendTraitAlias) GetModifiers() uint32                      { return this.modifiers }

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
