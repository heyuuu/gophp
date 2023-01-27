// <<generate>>

package zend

/**
 * ZendRefcountedH
 */
type ZendRefcountedH struct {
	refcount uint32
	u        struct /* union */ {
		type_info uint32
	}
}

func (this *ZendRefcountedH) GetRefcount() uint32 { return this.refcount }
func (this *ZendRefcountedH) SetRefcount(value uint32) uint32 {
	this.refcount = value
	return this.refcount
}

func (this *ZendRefcountedH) IncRefcount() uint32 {
	this.refcount++
	return this.refcount
}

func (this *ZendRefcountedH) IncRefcountEx(rc uint32) uint32 {
	this.refcount += rc
	return this.refcount
}
func (this *ZendRefcountedH) DecRefcount() uint32 {
	ZEND_ASSERT(this.refcount > 0)
	this.refcount--
	return this.refcount
}

/**
 *  type_info 保存三个 flag 标识:
 *	type(4) + flags(6) + info(22)
 */
const GC_TYPE_MASK = 0xf
const GC_FLAGS_MASK = 0x3f0
const GC_INFO_MASK = 0xfffffc00
const GC_FLAGS_SHIFT = 0
const GC_INFO_SHIFT = 10

func (this *ZendRefcountedH) GetTypeInfo() uint32      { return this.u.type_info }
func (this *ZendRefcountedH) SetTypeInfo(value uint32) { this.u.type_info = value }

func (this *ZendRefcountedH) GetType() uint8 {
	var typeInfo = this.GetTypeInfo()
	return uint8(typeInfo & GC_TYPE_MASK)
}

func (this *ZendRefcountedH) GetFlags() uint32 {
	var typeInfo = this.GetTypeInfo()
	return typeInfo >> GC_FLAGS_SHIFT & GC_FLAGS_MASK >> GC_FLAGS_SHIFT
}

func (this *ZendRefcountedH) GetInfo() uint32 {
	var typeInfo = this.GetTypeInfo()
	return typeInfo >> GC_INFO_SHIFT
}

func (this *ZendRefcountedH) AddFlags(flags uint32) {
	this.u.type_info |= flags << GC_FLAGS_SHIFT
}

func (this *ZendRefcountedH) DelFlags(flags uint32) {
	this.u.type_info &^= flags << GC_FLAGS_SHIFT
}

/**
 * ZendRefcounted
 */
type ZendRefcounted interface {
	GetGc() *ZendRefcountedH
	SetGc(value ZendRefcountedH)

	GetGcRefcount() uint32
	SetGcRefcount(value uint32) uint32
	IncGcRefcount() uint32
	IncGcRefcountEx(rc uint32) uint32
	DecGcRefcount() uint32

	GetGcTypeInfo() uint32
	SetGcTypeInfo(typeInfo uint32)
	GetGcType() uint8
	GetGcFlags() uint32
	GetGcInfo() uint32

	AddGcFlags(flags uint32)
	DelGcFlags(flags uint32)
}

type baseZendRefcounted struct {
	gc ZendRefcountedH
}

var _ ZendRefcounted = &baseZendRefcounted{}

func (this *baseZendRefcounted) GetGc() *ZendRefcountedH     { return &this.gc }
func (this *baseZendRefcounted) SetGc(value ZendRefcountedH) { this.gc = value }

func (this *baseZendRefcounted) GetGcRefcount() uint32             { return this.gc.GetRefcount() }
func (this *baseZendRefcounted) SetGcRefcount(value uint32) uint32 { return this.gc.SetRefcount(value) }
func (this *baseZendRefcounted) IncGcRefcount() uint32             { return this.gc.IncRefcount() }
func (this *baseZendRefcounted) IncGcRefcountEx(rc uint32) uint32  { return this.gc.IncRefcountEx(rc) }
func (this *baseZendRefcounted) DecGcRefcount() uint32             { return this.gc.DecRefcount() }

func (this *baseZendRefcounted) GetGcTypeInfo() uint32      { return this.gc.GetTypeInfo() }
func (this *baseZendRefcounted) SetGcTypeInfo(value uint32) { this.gc.SetTypeInfo(value) }

func (this *baseZendRefcounted) GetGcType() uint8   { return this.gc.GetType() }
func (this *baseZendRefcounted) GetGcFlags() uint32 { return this.gc.GetFlags() }
func (this *baseZendRefcounted) GetGcInfo() uint32  { return this.gc.GetInfo() }

func (this *baseZendRefcounted) AddGcFlags(flags uint32) { this.gc.AddFlags(flags) }
func (this *baseZendRefcounted) DelGcFlags(flags uint32) { this.gc.DelFlags(flags) }
