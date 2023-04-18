package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendAbstractInfo
 */
type ZendAbstractInfo struct {
	afn  []types.IFunction
	cnt  int
	ctor int
}

func (this *ZendAbstractInfo) GetAfn() []types.IFunction { return this.afn }
func (this *ZendAbstractInfo) GetCnt() int               { return this.cnt }
func (this *ZendAbstractInfo) GetCtor() int              { return this.ctor }
func (this *ZendAbstractInfo) SetCtor(value int)         { this.ctor = value }

/**
 * VarianceObligation
 */
type VarianceObligation struct {
	type_ VarianceObligationType
	__0   struct /* union */ {
		dependency_ce *types.ClassEntry
		__0           struct {
			parent_fn    types.IFunction
			child_fn     types.IFunction
			always_error types.ZendBool
		}
		__1 struct {
			parent_prop *ZendPropertyInfo
			child_prop  *ZendPropertyInfo
		}
	}
}

func (this *VarianceObligation) GetType() VarianceObligationType      { return this.type_ }
func (this *VarianceObligation) SetType(value VarianceObligationType) { this.type_ = value }
func (this *VarianceObligation) GetDependencyCe() *types.ClassEntry   { return this.__0.dependency_ce }
func (this *VarianceObligation) SetDependencyCe(value *types.ClassEntry) {
	this.__0.dependency_ce = value
}
func (this *VarianceObligation) GetParentFn() types.IFunction   { return this.__0.__0.parent_fn }
func (this *VarianceObligation) GetChildFn() types.IFunction    { return this.__0.__0.child_fn }
func (this *VarianceObligation) GetAlwaysError() types.ZendBool { return this.__0.__0.always_error }
func (this *VarianceObligation) SetAlwaysError(value types.ZendBool) {
	this.__0.__0.always_error = value
}
func (this *VarianceObligation) GetParentProp() *ZendPropertyInfo { return this.__0.__1.parent_prop }
func (this *VarianceObligation) SetParentProp(value *ZendPropertyInfo) {
	this.__0.__1.parent_prop = value
}
func (this *VarianceObligation) GetChildProp() *ZendPropertyInfo { return this.__0.__1.child_prop }
func (this *VarianceObligation) SetChildProp(value *ZendPropertyInfo) {
	this.__0.__1.child_prop = value
}
