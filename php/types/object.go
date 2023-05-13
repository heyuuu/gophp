package types

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"runtime"
)

/**
 * ZendObject
 */
type ZendObject struct {
	ZendRefcounted
	handle          uint32
	ce              *ClassEntry
	handlers        *ObjectHandlers
	properties      *Array
	propertiesTable []Zval
}

func NewObject(ce *ClassEntry, handle uint32, handlers *ObjectHandlers) *ZendObject {
	propertyCount := ce.GetDefaultPropertiesCount()
	if ce.IsUseGuards() {
		propertyCount++
	}

	o := &ZendObject{}
	o.handlers = handlers
	o.propertiesTable = make([]Zval, propertyCount)

	o.Init(ce, handle)
	return o
}

func (o *ZendObject) Init(ce *ClassEntry, handle uint32) {
	o.SetGcTypeInfo(uint32(IS_OBJECT) | GC_COLLECTABLE<<GC_FLAGS_SHIFT)

	o.handle = handle
	o.ce = ce
	o.properties = nil

	if ce.IsUseGuards() {
		o.propertiesTable[ce.GetDefaultPropertiesCount()].SetUndef()
	}

	runtime.SetFinalizer(o, ObjectAutoFree)
}

func (o *ZendObject) GetHandle() uint32                 { return o.handle }
func (o *ZendObject) GetCe() *ClassEntry                { return o.ce }
func (o *ZendObject) GetHandlers() *ObjectHandlers      { return o.handlers }
func (o *ZendObject) SetHandlers(value *ObjectHandlers) { o.handlers = value }
func (o *ZendObject) GetProperties() *Array             { return o.properties }
func (o *ZendObject) SetProperties(value *Array)        { o.properties = value }
func (o *ZendObject) GetPropertiesTable() []Zval        { return o.propertiesTable }

// object handlers
func (o *ZendObject) Free() { o.handlers.freeObj(o) }
func (o *ZendObject) Dtor() { o.handlers.dtorObj(o) }

func (o *ZendObject) CanClone() bool { return o.handlers.cloneObj != nil }
func (o *ZendObject) Clone(zv *types.Zval) *ZendObject {
	b.Assert(o.handlers.cloneObj != nil)
	return o.handlers.cloneObj(zv)
}
