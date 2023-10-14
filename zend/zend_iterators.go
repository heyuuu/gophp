package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

/**
 * constants and global variables
 */
var ZendIteratorClassEntry *types.ClassEntry = types.NewInternalClassSimple("__iterator_wrapper")

var IteratorObjectHandlers *types.ObjectHandlers = types.NewObjectHandlers(types.ObjectHandlersSetting{
	FreeObj: func(object *types.Object) {
		var iter *ZendObjectIterator = (*ZendObjectIterator)(object)
		iter.GetFuncs().GetDtor()(iter)
	},
	DtorObj: func(object *types.Object) {},
})

/**
 * functions
 */
func ZendIteratorInit(iter *ZendObjectIterator) {
	iter.Init()
}
func ZendIteratorUnwrap(arrayPtr *types.Zval) *ZendObjectIterator {
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
	get_current_data   func(iter *ZendObjectIterator) *types.Zval
	get_current_key    func(iter *ZendObjectIterator, key *types.Zval)
	move_forward       func(iter *ZendObjectIterator)
	rewind             func(iter *ZendObjectIterator)
	invalidate_current func(iter *ZendObjectIterator)
}

func MakeZendObjectIteratorFuncs(
	dtor func(iter *ZendObjectIterator),
	valid func(iter *ZendObjectIterator) int,
	get_current_data func(iter *ZendObjectIterator) *types.Zval,
	get_current_key func(iter *ZendObjectIterator, key *types.Zval),
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
func (this *ZendObjectIteratorFuncs) GetGetCurrentData() func(iter *ZendObjectIterator) *types.Zval {
	return this.get_current_data
}
func (this *ZendObjectIteratorFuncs) GetGetCurrentKey() func(iter *ZendObjectIterator, key *types.Zval) {
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
	std   *types.Object
	data  types.Zval
	funcs *ZendObjectIteratorFuncs
	index ZendUlong
}

func (this *ZendObjectIterator) Init() {
	this.std = types.NewObject(ZendIteratorClassEntry, IteratorObjectHandlers)
}

func (this *ZendObjectIterator) GetStd() *types.Object                   { return this.std }
func (this *ZendObjectIterator) GetData() types.Zval                     { return this.data }
func (this *ZendObjectIterator) GetFuncs() *ZendObjectIteratorFuncs      { return this.funcs }
func (this *ZendObjectIterator) SetFuncs(value *ZendObjectIteratorFuncs) { this.funcs = value }
func (this *ZendObjectIterator) GetIndex() ZendUlong                     { return this.index }
func (this *ZendObjectIterator) SetIndex(value ZendUlong)                { this.index = value }

/**
 * ZendClassIteratorFuncs
 */
type ZendClassIteratorFuncs struct {
	zf_new_iterator types.IFunction
	zf_valid        types.IFunction
	zf_current      types.IFunction
	zf_key          types.IFunction
	zf_next         types.IFunction
	zf_rewind       types.IFunction
}

func NewClassIteratorFuncs() *ZendClassIteratorFuncs {
	return &ZendClassIteratorFuncs{}
}

func (this *ZendClassIteratorFuncs) GetZfNewIterator() types.IFunction {
	return this.zf_new_iterator
}
func (this *ZendClassIteratorFuncs) SetZfNewIterator(value types.IFunction) {
	this.zf_new_iterator = value
}
func (this *ZendClassIteratorFuncs) GetZfValid() types.IFunction        { return this.zf_valid }
func (this *ZendClassIteratorFuncs) SetZfValid(value types.IFunction)   { this.zf_valid = value }
func (this *ZendClassIteratorFuncs) GetZfCurrent() types.IFunction      { return this.zf_current }
func (this *ZendClassIteratorFuncs) SetZfCurrent(value types.IFunction) { this.zf_current = value }
func (this *ZendClassIteratorFuncs) GetZfKey() types.IFunction          { return this.zf_key }
func (this *ZendClassIteratorFuncs) SetZfKey(value types.IFunction)     { this.zf_key = value }
func (this *ZendClassIteratorFuncs) GetZfNext() types.IFunction         { return this.zf_next }
func (this *ZendClassIteratorFuncs) SetZfNext(value types.IFunction)    { this.zf_next = value }
func (this *ZendClassIteratorFuncs) GetZfRewind() types.IFunction       { return this.zf_rewind }
func (this *ZendClassIteratorFuncs) SetZfRewind(value types.IFunction)  { this.zf_rewind = value }
