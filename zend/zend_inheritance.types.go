package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * ZendAbstractInfo
 */
type ZendAbstractInfo struct {
	afn  []types2.IFunction
	cnt  int
	ctor int
}

func (this *ZendAbstractInfo) GetAfn() []types2.IFunction { return this.afn }
func (this *ZendAbstractInfo) GetCnt() int                { return this.cnt }
func (this *ZendAbstractInfo) GetCtor() int               { return this.ctor }
func (this *ZendAbstractInfo) SetCtor(value int)          { this.ctor = value }

/**
 * VarianceObligation
 */
type VarianceObligation struct {
	type_ VarianceObligationType
	__0   struct /* union */ {
		dependency_ce *types2.ClassEntry
		__0           struct {
			parent_fn    types2.IFunction
			child_fn     types2.IFunction
			always_error types2.ZendBool
		}
		__1 struct {
			parent_prop *ZendPropertyInfo
			child_prop  *ZendPropertyInfo
		}
	}
}

func (this *VarianceObligation) GetType() VarianceObligationType      { return this.type_ }
func (this *VarianceObligation) SetType(value VarianceObligationType) { this.type_ = value }
func (this *VarianceObligation) GetDependencyCe() *types2.ClassEntry  { return this.__0.dependency_ce }
func (this *VarianceObligation) SetDependencyCe(value *types2.ClassEntry) {
	this.__0.dependency_ce = value
}
func (this *VarianceObligation) GetParentFn() types2.IFunction   { return this.__0.__0.parent_fn }
func (this *VarianceObligation) GetChildFn() types2.IFunction    { return this.__0.__0.child_fn }
func (this *VarianceObligation) GetAlwaysError() types2.ZendBool { return this.__0.__0.always_error }
func (this *VarianceObligation) SetAlwaysError(value types2.ZendBool) {
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
