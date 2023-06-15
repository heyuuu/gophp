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
	std                   *types.ZendObject
	iterator              *ZendObjectIterator
	executeData           *ZendExecuteData
	frozenCallStack       *ZendExecuteData
	value                 types.Zval
	key                   types.Zval
	retval                types.Zval
	sendTarget            *types.Zval
	largestUsedIntegerKey ZendLong
	values                types.Zval
	node                  ZendGeneratorNode
	executeFake           ZendExecuteData
	flags                 uint8
	gcBuffer              *types.Zval
	gcBufferSize          uint32
}

func NewZendGenerator(ce *types.ClassEntry) *ZendGenerator {
	generator := &ZendGenerator{
		std: types.NewObject(ce, &ZendGeneratorHandlers),
		/* The key will be incremented on first use, so it'll start at 0 */
		largestUsedIntegerKey: -1,
	}
	generator.retval.SetUndef()
	generator.values.SetUndef()

	/* By default we have a tree of only one node */
	generator.node.SetParent(nil)
	generator.node.SetChildren(0)
	generator.node.SetRoot(generator)

	return generator
}

func (gen *ZendGenerator) GetValues() *types.Zval  { return &gen.values }
func (gen *ZendGenerator) GetValuesFePos() uint32  { return gen.values.GetFePos() }
func (gen *ZendGenerator) SetValuesFePos(v uint32) { gen.values.SetFePos(v) }

func (gen *ZendGenerator) GetStd() *types.ZendObject                 { return gen.std }
func (gen *ZendGenerator) GetIterator() *ZendObjectIterator          { return gen.iterator }
func (gen *ZendGenerator) SetIterator(value *ZendObjectIterator)     { gen.iterator = value }
func (gen *ZendGenerator) GetExecuteData() *ZendExecuteData          { return gen.executeData }
func (gen *ZendGenerator) SetExecuteData(value *ZendExecuteData)     { gen.executeData = value }
func (gen *ZendGenerator) GetFrozenCallStack() *ZendExecuteData      { return gen.frozenCallStack }
func (gen *ZendGenerator) SetFrozenCallStack(value *ZendExecuteData) { gen.frozenCallStack = value }
func (gen *ZendGenerator) GetValue() *types.Zval                     { return &gen.value }
func (gen *ZendGenerator) GetKey() *types.Zval                       { return &gen.key }
func (gen *ZendGenerator) GetRetval() types.Zval                     { return gen.retval }
func (gen *ZendGenerator) GetSendTarget() *types.Zval                { return gen.sendTarget }
func (gen *ZendGenerator) SetSendTarget(value *types.Zval)           { gen.sendTarget = value }
func (gen *ZendGenerator) GetLargestUsedIntegerKey() ZendLong        { return gen.largestUsedIntegerKey }
func (gen *ZendGenerator) SetLargestUsedIntegerKey(value ZendLong) {
	gen.largestUsedIntegerKey = value
}
func (gen *ZendGenerator) GetNode() *ZendGeneratorNode      { return &gen.node }
func (gen *ZendGenerator) GetExecuteFake() *ZendExecuteData { return &gen.executeFake }
func (gen *ZendGenerator) GetGcBuffer() *types.Zval         { return gen.gcBuffer }
func (gen *ZendGenerator) SetGcBuffer(value *types.Zval)    { gen.gcBuffer = value }
func (gen *ZendGenerator) GetGcBufferSize() uint32          { return gen.gcBufferSize }
func (gen *ZendGenerator) SetGcBufferSize(value uint32)     { gen.gcBufferSize = value }

/* ZendGenerator.flags */
func (gen *ZendGenerator) AddFlags(value uint8)      { gen.flags |= value }
func (gen *ZendGenerator) SubFlags(value uint8)      { gen.flags &^= value }
func (gen *ZendGenerator) HasFlags(value uint8) bool { return gen.flags&value != 0 }
func (gen *ZendGenerator) SwitchFlags(value uint8, cond bool) {
	if cond {
		gen.AddFlags(value)
	} else {
		gen.SubFlags(value)
	}
}
func (gen ZendGenerator) IsCurrentlyRunning() bool {
	return gen.HasFlags(ZEND_GENERATOR_CURRENTLY_RUNNING)
}
func (gen ZendGenerator) IsDoInit() bool       { return gen.HasFlags(ZEND_GENERATOR_DO_INIT) }
func (gen ZendGenerator) IsAtFirstYield() bool { return gen.HasFlags(ZEND_GENERATOR_AT_FIRST_YIELD) }
func (gen ZendGenerator) IsForcedClose() bool  { return gen.HasFlags(ZEND_GENERATOR_FORCED_CLOSE) }
func (gen *ZendGenerator) SetIsCurrentlyRunning(cond bool) {
	gen.SwitchFlags(ZEND_GENERATOR_CURRENTLY_RUNNING, cond)
}
func (gen *ZendGenerator) SetIsDoInit(cond bool) { gen.SwitchFlags(ZEND_GENERATOR_DO_INIT, cond) }
func (gen *ZendGenerator) SetIsAtFirstYield(cond bool) {
	gen.SwitchFlags(ZEND_GENERATOR_AT_FIRST_YIELD, cond)
}
func (gen *ZendGenerator) SetIsForcedClose(cond bool) {
	gen.SwitchFlags(ZEND_GENERATOR_FORCED_CLOSE, cond)
}
