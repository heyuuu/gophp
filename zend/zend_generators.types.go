package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * ZendGeneratorNode
 */
type ZendGeneratorNode struct {
	parent   *ZendGenerator
	children uint32
	child    struct /* union */ {
		ht     *types.Array
		single struct {
			leaf  *ZendGenerator
			child *ZendGenerator
		}
	}
	ptr struct /* union */ {
		leaf *ZendGenerator
		root *ZendGenerator
	}
}

func (this *ZendGeneratorNode) GetParent() *ZendGenerator          { return this.parent }
func (this *ZendGeneratorNode) SetParent(value *ZendGenerator)     { this.parent = value }
func (this *ZendGeneratorNode) GetChildren() uint32                { return this.children }
func (this *ZendGeneratorNode) SetChildren(value uint32)           { this.children = value }
func (this *ZendGeneratorNode) GetHt() *types.Array                { return this.child.ht }
func (this *ZendGeneratorNode) SetHt(value *types.Array)           { this.child.ht = value }
func (this *ZendGeneratorNode) GetChildSingleLeaf() *ZendGenerator { return this.child.single.leaf }
func (this *ZendGeneratorNode) SetChildSingleLeaf(value *ZendGenerator) {
	this.child.single.leaf = value
}
func (this *ZendGeneratorNode) GetChildSingleChild() *ZendGenerator { return this.child.single.child }
func (this *ZendGeneratorNode) SetChildSingleChild(value *ZendGenerator) {
	this.child.single.child = value
}
func (this *ZendGeneratorNode) GetPtrLeaf() *ZendGenerator      { return this.ptr.leaf }
func (this *ZendGeneratorNode) SetPtrLeaf(value *ZendGenerator) { this.ptr.leaf = value }
func (this *ZendGeneratorNode) GetRoot() *ZendGenerator         { return this.ptr.root }
func (this *ZendGeneratorNode) SetRoot(value *ZendGenerator)    { this.ptr.root = value }

/**
 * ZendGenerator
 */
type ZendGenerator struct {
	std                      types.ZendObject
	iterator                 *ZendObjectIterator
	executeData              *ZendExecuteData
	frozen_call_stack        *ZendExecuteData
	value                    types.Zval
	key                      types.Zval
	retval                   types.Zval
	send_target              *types.Zval
	largest_used_integer_key ZendLong
	values                   types.Zval
	node                     ZendGeneratorNode
	execute_fake             ZendExecuteData
	flags                    uint8
	gc_buffer                *types.Zval
	gc_buffer_size           uint32
}

//             func MakeZendGenerator(
// std ZendObject,
// iterator *ZendObjectIterator,
// executeData *ZendExecuteData,
// frozen_call_stack *ZendExecuteData,
// value Zval,
// key Zval,
// retval Zval,
// send_target *Zval,
// largest_used_integer_key ZendLong,
// values Zval,
// node ZendGeneratorNode,
// execute_fake ZendExecuteData,
// flags ZendUchar,
// gc_buffer *Zval,
// gc_buffer_size uint32,
// ) ZendGenerator {
//                 return ZendGenerator{
//                     std:std,
//                     iterator:iterator,
//                     executeData:execute_data,
//                     frozen_call_stack:frozen_call_stack,
//                     value:value,
//                     key:key,
//                     retval:retval,
//                     send_target:send_target,
//                     largest_used_integer_key:largest_used_integer_key,
//                     values:values,
//                     node:node,
//                     execute_fake:execute_fake,
//                     flags:flags,
//                     gc_buffer:gc_buffer,
//                     gc_buffer_size:gc_buffer_size,
//                 }
//             }
func (this *ZendGenerator) GetStd() types.ZendObject { return this.std }

// func (this *ZendGenerator) SetStd(value ZendObject) { this.std = value }
func (this *ZendGenerator) GetIterator() *ZendObjectIterator          { return this.iterator }
func (this *ZendGenerator) SetIterator(value *ZendObjectIterator)     { this.iterator = value }
func (this *ZendGenerator) GetExecuteData() *ZendExecuteData          { return this.executeData }
func (this *ZendGenerator) SetExecuteData(value *ZendExecuteData)     { this.executeData = value }
func (this *ZendGenerator) GetFrozenCallStack() *ZendExecuteData      { return this.frozen_call_stack }
func (this *ZendGenerator) SetFrozenCallStack(value *ZendExecuteData) { this.frozen_call_stack = value }
func (this *ZendGenerator) GetValue() types.Zval                      { return this.value }

// func (this *ZendGenerator) SetValue(value Zval) { this.value = value }
func (this *ZendGenerator) GetKey() types.Zval { return this.key }

// func (this *ZendGenerator) SetKey(value Zval) { this.key = value }
func (this *ZendGenerator) GetRetval() types.Zval { return this.retval }

// func (this *ZendGenerator) SetRetval(value Zval) { this.retval = value }
func (this *ZendGenerator) GetSendTarget() *types.Zval         { return this.send_target }
func (this *ZendGenerator) SetSendTarget(value *types.Zval)    { this.send_target = value }
func (this *ZendGenerator) GetLargestUsedIntegerKey() ZendLong { return this.largest_used_integer_key }
func (this *ZendGenerator) SetLargestUsedIntegerKey(value ZendLong) {
	this.largest_used_integer_key = value
}
func (this *ZendGenerator) GetValues() types.Zval { return this.values }

// func (this *ZendGenerator) SetValues(value Zval) { this.values = value }
func (this *ZendGenerator) GetNode() ZendGeneratorNode { return this.node }

// func (this *ZendGenerator) SetNode(value ZendGeneratorNode) { this.node = value }
func (this *ZendGenerator) GetExecuteFake() *ZendExecuteData { return &this.execute_fake }

// func (this *ZendGenerator) SetExecuteFake(value ZendExecuteData) { this.execute_fake = value }
// func (this *ZendGenerator)  GetFlags() ZendUchar      { return this.flags }
// func (this *ZendGenerator) SetFlags(value ZendUchar) { this.flags = value }
func (this *ZendGenerator) GetGcBuffer() *types.Zval      { return this.gc_buffer }
func (this *ZendGenerator) SetGcBuffer(value *types.Zval) { this.gc_buffer = value }
func (this *ZendGenerator) GetGcBufferSize() uint32       { return this.gc_buffer_size }
func (this *ZendGenerator) SetGcBufferSize(value uint32)  { this.gc_buffer_size = value }

/* ZendGenerator.flags */
func (this *ZendGenerator) AddFlags(value uint8)      { this.flags |= value }
func (this *ZendGenerator) SubFlags(value uint8)      { this.flags &^= value }
func (this *ZendGenerator) HasFlags(value uint8) bool { return this.flags&value != 0 }
func (this *ZendGenerator) SwitchFlags(value uint8, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendGenerator) IsCurrentlyRunning() bool {
	return this.HasFlags(ZEND_GENERATOR_CURRENTLY_RUNNING)
}
func (this ZendGenerator) IsDoInit() bool       { return this.HasFlags(ZEND_GENERATOR_DO_INIT) }
func (this ZendGenerator) IsAtFirstYield() bool { return this.HasFlags(ZEND_GENERATOR_AT_FIRST_YIELD) }
func (this ZendGenerator) IsForcedClose() bool  { return this.HasFlags(ZEND_GENERATOR_FORCED_CLOSE) }
func (this *ZendGenerator) SetIsCurrentlyRunning(cond bool) {
	this.SwitchFlags(ZEND_GENERATOR_CURRENTLY_RUNNING, cond)
}
func (this *ZendGenerator) SetIsDoInit(cond bool) { this.SwitchFlags(ZEND_GENERATOR_DO_INIT, cond) }
func (this *ZendGenerator) SetIsAtFirstYield(cond bool) {
	this.SwitchFlags(ZEND_GENERATOR_AT_FIRST_YIELD, cond)
}
func (this *ZendGenerator) SetIsForcedClose(cond bool) {
	this.SwitchFlags(ZEND_GENERATOR_FORCED_CLOSE, cond)
}
