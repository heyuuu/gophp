package standard

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpSerializeData
 */
type PhpSerializeData struct {
	ht *types.Array
	n  uint32
}

func (this *PhpSerializeData) GetHt() *types.Array      { return this.ht }
func (this *PhpSerializeData) SetHt(value *types.Array) { this.ht = value }
func (this *PhpSerializeData) GetN() uint32             { return this.n }
func (this *PhpSerializeData) SetN(value uint32)        { this.n = value }
