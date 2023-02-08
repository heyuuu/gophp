// <<generate>>

package zend

/**
 * ZendObjectIteratorFuncs
 */
type ZendObjectIteratorFuncs struct {
	dtor               func(iter *ZendObjectIterator)
	valid              func(iter *ZendObjectIterator) int
	get_current_data   func(iter *ZendObjectIterator) *Zval
	get_current_key    func(iter *ZendObjectIterator, key *Zval)
	move_forward       func(iter *ZendObjectIterator)
	rewind             func(iter *ZendObjectIterator)
	invalidate_current func(iter *ZendObjectIterator)
}

func MakeZendObjectIteratorFuncs(
	dtor func(iter *ZendObjectIterator),
	valid func(iter *ZendObjectIterator) int,
	get_current_data func(iter *ZendObjectIterator) *Zval,
	get_current_key func(iter *ZendObjectIterator, key *Zval),
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
func (this *ZendObjectIteratorFuncs) GetDtor() func(iter *ZendObjectIterator) { return this.dtor }

// func (this *ZendObjectIteratorFuncs) SetDtor(value func(iter *ZendObjectIterator)) { this.dtor = value }
func (this *ZendObjectIteratorFuncs) GetValid() func(iter *ZendObjectIterator) int { return this.valid }

// func (this *ZendObjectIteratorFuncs) SetValid(value func(iter *ZendObjectIterator) int) { this.valid = value }
func (this *ZendObjectIteratorFuncs) GetGetCurrentData() func(iter *ZendObjectIterator) *Zval {
	return this.get_current_data
}

// func (this *ZendObjectIteratorFuncs) SetGetCurrentData(value func(iter *ZendObjectIterator) *Zval) { this.get_current_data = value }
func (this *ZendObjectIteratorFuncs) GetGetCurrentKey() func(iter *ZendObjectIterator, key *Zval) {
	return this.get_current_key
}

// func (this *ZendObjectIteratorFuncs) SetGetCurrentKey(value func(iter *ZendObjectIterator, key *Zval)) { this.get_current_key = value }
func (this *ZendObjectIteratorFuncs) GetMoveForward() func(iter *ZendObjectIterator) {
	return this.move_forward
}

// func (this *ZendObjectIteratorFuncs) SetMoveForward(value func(iter *ZendObjectIterator)) { this.move_forward = value }
func (this *ZendObjectIteratorFuncs) GetRewind() func(iter *ZendObjectIterator) { return this.rewind }

// func (this *ZendObjectIteratorFuncs) SetRewind(value func(iter *ZendObjectIterator)) { this.rewind = value }
func (this *ZendObjectIteratorFuncs) GetInvalidateCurrent() func(iter *ZendObjectIterator) {
	return this.invalidate_current
}

// func (this *ZendObjectIteratorFuncs) SetInvalidateCurrent(value func(iter *ZendObjectIterator)) { this.invalidate_current = value }

/**
 * ZendObjectIterator
 */
type ZendObjectIterator struct {
	std   ZendObject
	data  Zval
	funcs *ZendObjectIteratorFuncs
	index ZendUlong
}

// func MakeZendObjectIterator(std ZendObject, data Zval, funcs *ZendObjectIteratorFuncs, index ZendUlong) ZendObjectIterator {
//     return ZendObjectIterator{
//         std:std,
//         data:data,
//         funcs:funcs,
//         index:index,
//     }
// }
func (this *ZendObjectIterator) GetStd() ZendObject { return this.std }

// func (this *ZendObjectIterator) SetStd(value ZendObject) { this.std = value }
func (this *ZendObjectIterator) GetData() Zval { return this.data }

// func (this *ZendObjectIterator) SetData(value Zval) { this.data = value }
func (this *ZendObjectIterator) GetFuncs() *ZendObjectIteratorFuncs      { return this.funcs }
func (this *ZendObjectIterator) SetFuncs(value *ZendObjectIteratorFuncs) { this.funcs = value }
func (this *ZendObjectIterator) GetIndex() ZendUlong                     { return this.index }
func (this *ZendObjectIterator) SetIndex(value ZendUlong)                { this.index = value }

/**
 * ZendClassIteratorFuncs
 */
type ZendClassIteratorFuncs struct {
	zf_new_iterator *ZendFunction
	zf_valid        *ZendFunction
	zf_current      *ZendFunction
	zf_key          *ZendFunction
	zf_next         *ZendFunction
	zf_rewind       *ZendFunction
}

//             func MakeZendClassIteratorFuncs(
// zf_new_iterator *ZendFunction,
// zf_valid *ZendFunction,
// zf_current *ZendFunction,
// zf_key *ZendFunction,
// zf_next *ZendFunction,
// zf_rewind *ZendFunction,
// ) ZendClassIteratorFuncs {
//                 return ZendClassIteratorFuncs{
//                     zf_new_iterator:zf_new_iterator,
//                     zf_valid:zf_valid,
//                     zf_current:zf_current,
//                     zf_key:zf_key,
//                     zf_next:zf_next,
//                     zf_rewind:zf_rewind,
//                 }
//             }
func (this *ZendClassIteratorFuncs) GetZfNewIterator() *ZendFunction { return this.zf_new_iterator }
func (this *ZendClassIteratorFuncs) SetZfNewIterator(value *ZendFunction) {
	this.zf_new_iterator = value
}
func (this *ZendClassIteratorFuncs) GetZfValid() *ZendFunction        { return this.zf_valid }
func (this *ZendClassIteratorFuncs) SetZfValid(value *ZendFunction)   { this.zf_valid = value }
func (this *ZendClassIteratorFuncs) GetZfCurrent() *ZendFunction      { return this.zf_current }
func (this *ZendClassIteratorFuncs) SetZfCurrent(value *ZendFunction) { this.zf_current = value }
func (this *ZendClassIteratorFuncs) GetZfKey() *ZendFunction          { return this.zf_key }
func (this *ZendClassIteratorFuncs) SetZfKey(value *ZendFunction)     { this.zf_key = value }
func (this *ZendClassIteratorFuncs) GetZfNext() *ZendFunction         { return this.zf_next }
func (this *ZendClassIteratorFuncs) SetZfNext(value *ZendFunction)    { this.zf_next = value }
func (this *ZendClassIteratorFuncs) GetZfRewind() *ZendFunction       { return this.zf_rewind }
func (this *ZendClassIteratorFuncs) SetZfRewind(value *ZendFunction)  { this.zf_rewind = value }
