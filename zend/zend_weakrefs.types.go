package zend

import "github.com/heyuuu/gophp/zend/types"

/**
 * ZendWeakref
 */
type ZendWeakref struct {
	referent *types.ZendObject
	std      types.ZendObject
}

func (this *ZendWeakref) GetReferent() *types.ZendObject      { return this.referent }
func (this *ZendWeakref) SetReferent(value *types.ZendObject) { this.referent = value }
func (this *ZendWeakref) GetStd() types.ZendObject            { return this.std }
