package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
)

/**
 * constants and global variables
 */
var ZendIteratorClassEntry *types2.ClassEntry = types2.NewClassEntry("__iterator_wrapper", nil)

var IteratorObjectHandlers *ZendObjectHandlers = NewZendObjectHandlers(ObjectHandlersSetting{
	FreeObj: func(object *types2.ZendObject) {
		var iter *ZendObjectIterator = (*ZendObjectIterator)(object)
		iter.GetFuncs().GetDtor()(iter)
	},
	GetGc: func(object *types2.Zval, table **types2.Zval, n *int) *types2.Array {
		*table = nil
		*n = 0
		return nil
	},
})

/**
 * functions
 */
func ZendIteratorInit(iter *ZendObjectIterator) {
	ZendObjectStdInit(&iter.std, ZendIteratorClassEntry)
	iter.std.SetHandlers(IteratorObjectHandlers)
}
func ZendIteratorUnwrap(arrayPtr *types2.Zval) *ZendObjectIterator {
	b.Assert(arrayPtr.IsObject())
	if arrayPtr.Object().GetHandlers() == IteratorObjectHandlers {
		return (*ZendObjectIterator)(arrayPtr.Object())
	}
	return nil
}

/**
 * types
 */

/**
 * ZendObjectIteratorFuncs
 */
type ZendObjectIteratorFuncs struct {
	dtor               func(iter *ZendObjectIterator)
	valid              func(iter *ZendObjectIterator) int
	get_current_data   func(iter *ZendObjectIterator) *types2.Zval
	get_current_key    func(iter *ZendObjectIterator, key *types2.Zval)
	move_forward       func(iter *ZendObjectIterator)
	rewind             func(iter *ZendObjectIterator)
	invalidate_current func(iter *ZendObjectIterator)
}

func MakeZendObjectIteratorFuncs(
	dtor func(iter *ZendObjectIterator),
	valid func(iter *ZendObjectIterator) int,
	get_current_data func(iter *ZendObjectIterator) *types2.Zval,
	get_current_key func(iter *ZendObjectIterator, key *types2.Zval),
	move_forward func(iter *ZendObjectIterator),
	rewind func(iter *ZendObjectIterator),
	invalidate_current func(iter *ZendObjectIterator),
) ZendObjectIteratorFuncs {
	return ZendObjectIteratorFuncs{
		dtor:               dtor,
		valid:              valid,
		get_current_data:   get_current_data,
		get_current_key:    get_current_key,
		move_forward:       move_forward,
		rewind:             rewind,
		invalidate_current: invalidate_current,
	}
}
func (this *ZendObjectIteratorFuncs) GetDtor() func(iter *ZendObjectIterator)      { return this.dtor }
func (this *ZendObjectIteratorFuncs) GetValid() func(iter *ZendObjectIterator) int { return this.valid }
func (this *ZendObjectIteratorFuncs) GetGetCurrentData() func(iter *ZendObjectIterator) *types2.Zval {
	return this.get_current_data
}
func (this *ZendObjectIteratorFuncs) GetGetCurrentKey() func(iter *ZendObjectIterator, key *types2.Zval) {
	return this.get_current_key
}
func (this *ZendObjectIteratorFuncs) GetMoveForward() func(iter *ZendObjectIterator) {
	return this.move_forward
}
func (this *ZendObjectIteratorFuncs) GetRewind() func(iter *ZendObjectIterator) { return this.rewind }
func (this *ZendObjectIteratorFuncs) GetInvalidateCurrent() func(iter *ZendObjectIterator) {
	return this.invalidate_current
}

/**
 * ZendObjectIterator
 */
type ZendObjectIterator struct {
	std   types2.ZendObject
	data  types2.Zval
	funcs *ZendObjectIteratorFuncs
	index ZendUlong
}

func (this *ZendObjectIterator) GetStd() types2.ZendObject               { return this.std }
func (this *ZendObjectIterator) GetData() types2.Zval                    { return this.data }
func (this *ZendObjectIterator) GetFuncs() *ZendObjectIteratorFuncs      { return this.funcs }
func (this *ZendObjectIterator) SetFuncs(value *ZendObjectIteratorFuncs) { this.funcs = value }
func (this *ZendObjectIterator) GetIndex() ZendUlong                     { return this.index }
func (this *ZendObjectIterator) SetIndex(value ZendUlong)                { this.index = value }

/**
 * ZendClassIteratorFuncs
 */
type ZendClassIteratorFuncs struct {
	zf_new_iterator types2.IFunction
	zf_valid        types2.IFunction
	zf_current      types2.IFunction
	zf_key          types2.IFunction
	zf_next         types2.IFunction
	zf_rewind       types2.IFunction
}

func (this *ZendClassIteratorFuncs) GetZfNewIterator() types2.IFunction {
	return this.zf_new_iterator
}
func (this *ZendClassIteratorFuncs) SetZfNewIterator(value types2.IFunction) {
	this.zf_new_iterator = value
}
func (this *ZendClassIteratorFuncs) GetZfValid() types2.IFunction        { return this.zf_valid }
func (this *ZendClassIteratorFuncs) SetZfValid(value types2.IFunction)   { this.zf_valid = value }
func (this *ZendClassIteratorFuncs) GetZfCurrent() types2.IFunction      { return this.zf_current }
func (this *ZendClassIteratorFuncs) SetZfCurrent(value types2.IFunction) { this.zf_current = value }
func (this *ZendClassIteratorFuncs) GetZfKey() types2.IFunction          { return this.zf_key }
func (this *ZendClassIteratorFuncs) SetZfKey(value types2.IFunction)     { this.zf_key = value }
func (this *ZendClassIteratorFuncs) GetZfNext() types2.IFunction         { return this.zf_next }
func (this *ZendClassIteratorFuncs) SetZfNext(value types2.IFunction)    { this.zf_next = value }
func (this *ZendClassIteratorFuncs) GetZfRewind() types2.IFunction       { return this.zf_rewind }
func (this *ZendClassIteratorFuncs) SetZfRewind(value types2.IFunction)  { this.zf_rewind = value }
