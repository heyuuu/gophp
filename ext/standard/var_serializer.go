package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

var COMMON = b.Cond(is_ref, "&", "")

/**
 * types
 */
type PhpSerializeDataT = *PhpSerializeData

/**
 * PhpSerializeData
 */
type PhpSerializeData struct {
	ht *types.Array
	n  uint32
}

func NewPhpSerializeData() *PhpSerializeData {
	return &PhpSerializeData{
		ht: types.NewArray(16),
		n:  0,
	}
}

func (d *PhpSerializeData) Destroy() {
	d.ht.Destroy()
}

func (d *PhpSerializeData) IncN() { d.n++ }
func (d *PhpSerializeData) DecN() { d.n-- }

func (d *PhpSerializeData) Mark(zv *types.Zval) {
	key := int(types.ZendUintptrT(zv.RefCounted()))

	d.ht.IndexAdd(key, types.NewZvalLong(int(d.n)))

	/* Additionally to the index, we also store the variable, to ensure that it is
	 * not destroyed during serialization and its pointer reused. The variable is
	 * stored at the numeric value of the pointer + 1, which cannot be the location
	 * of another zend_refcounted structure. */
	d.ht.IndexAdd(key+1, zv)
}

func (d *PhpSerializeData) FindMark(zv *types.Zval) *types.Zval {
	key := int(types.ZendUintptrT(zv.RefCounted()))

	return d.ht.IndexFind(key)
}
