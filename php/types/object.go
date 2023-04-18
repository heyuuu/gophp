package types

import (
	"github.com/heyuuu/gophp/zend"
	"runtime"
)

/**
 * ZendObject
 */
type ZendObject struct {
	ZendRefcounted
	handle          uint32
	ce              *ClassEntry
	handlers        *zend.ZendObjectHandlers
	properties      *Array
	propertiesTable []Zval
}

var _ IRefcounted = &ZendObject{}

func NewObject(ce *ClassEntry, handle uint32, handlers *zend.ZendObjectHandlers) *ZendObject {
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
	o.SetRefcount(1)
	o.SetGcTypeInfo(uint32(IS_OBJECT) | GC_COLLECTABLE<<GC_FLAGS_SHIFT)

	o.handle = handle
	o.ce = ce
	o.properties = nil

	if ce.IsUseGuards() {
		o.propertiesTable[ce.GetDefaultPropertiesCount()].SetUndef()
	}

	runtime.SetFinalizer(o, ObjectAutoFree)
}

func (o *ZendObject) GetHandle() uint32                          { return o.handle }
func (o *ZendObject) SetHandle(value uint32)                     { o.handle = value }
func (o *ZendObject) GetCe() *ClassEntry                         { return o.ce }
func (o *ZendObject) SetCe(value *ClassEntry)                    { o.ce = value }
func (o *ZendObject) GetHandlers() *zend.ZendObjectHandlers      { return o.handlers }
func (o *ZendObject) SetHandlers(value *zend.ZendObjectHandlers) { o.handlers = value }
func (o *ZendObject) GetProperties() *Array                      { return o.properties }
func (o *ZendObject) SetProperties(value *Array)                 { o.properties = value }
func (o *ZendObject) GetPropertiesTable() []Zval                 { return o.propertiesTable }
func (o *ZendObject) SetPropertiesTable(value []Zval)            { o.propertiesTable = value }
