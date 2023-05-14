package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"unsafe"
)

/**
 * types
 */
type PhpSerializeDataT = *PhpSerializeData

/**
 * PhpSerializeData
 */
type PhpSerializeData struct {
	nMap map[unsafe.Pointer]int
	vMap map[unsafe.Pointer]*types.Zval
	n    uint32
}

func NewPhpSerializeData() *PhpSerializeData {
	return &PhpSerializeData{
		nMap: make(map[unsafe.Pointer]int),
		vMap: make(map[unsafe.Pointer]*types.Zval),
		n:    0,
	}
}

func (d *PhpSerializeData) Destroy() {
	d.nMap = nil
	d.vMap = nil
	d.n = 0
}

func (d *PhpSerializeData) IncN() { d.n++ }
func (d *PhpSerializeData) DecN() { d.n-- }

func (d *PhpSerializeData) zvalKey(zv *types.Zval) unsafe.Pointer {
	b.Assert(zv.IsReference() || zv.IsObject())
	if zv.IsReference() {
		return unsafe.Pointer(zv.Reference())
	} else if zv.IsObject() {
		return unsafe.Pointer(zv.Object())
	} else {
		panic("unreachable")
	}
}

func (d *PhpSerializeData) Mark(zv *types.Zval) {
	key := d.zvalKey(zv)

	//d.ht.IndexAdd(key, types.NewZvalLong(int(d.n)))
	d.nMap[key] = int(d.n)

	/* Additionally to the index, we also store the variable, to ensure that it is
	 * not destroyed during serialization and its pointer reused. The variable is
	 * stored at the numeric value of the pointer + 1, which cannot be the location
	 * of another zend_refcounted structure. */
	//d.ht.IndexAdd(key+1, zv)
	d.vMap[key] = zv
}

func (d *PhpSerializeData) MarkUsed(zv *types.Zval) {
	key := d.zvalKey(zv)
	d.nMap[key] = -1
}

func (d *PhpSerializeData) FindMark(zv *types.Zval) (n int, exists bool) {
	key := d.zvalKey(zv)
	n, ok := d.nMap[key]
	return n, ok
}
