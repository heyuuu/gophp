package types

import (
	b "github.com/heyuuu/gophp/builtin"
)

/**
 * ZendRefcountedH
 */
type ZendRefcountedH struct {
	refcount uint32
	u        struct /* union */ {
		type_info uint32
	}
}

func (this *ZendRefcountedH) GetRefcount() uint32      { return this.refcount }
func (this *ZendRefcountedH) SetRefcount(value uint32) { this.refcount = value }
func (this *ZendRefcountedH) GetTypeInfo() uint32      { return this.u.type_info }
func (this *ZendRefcountedH) SetTypeInfo(value uint32) { this.u.type_info = value }

/**
 * ZendRefcounted
 */
type IRefcounted interface {
	GetGc() *ZendRefcountedH
	SetGc(value ZendRefcountedH)

	// refcount
	GetRefcount() uint32
	SetRefcount(rc uint32) uint32
	AddRefcount() uint32
	DelRefcount() uint32
	AddRefcountEx(rc uint32) uint32
	DelRefcountEx(rc uint32) uint32

	// type_info
	GetGcTypeInfo() uint32
	SetGcTypeInfo(typeInfo uint32)

	GetGcType() uint8
	GetGcFlags() uint32

	AddGcFlags(flags uint32)
	DelGcFlags(flags uint32)
	HasGcFlags(flags uint32) bool

	// flags
	SetCollectable()
	DelCollectable()

	IsImmutable() bool
	SetImmutable()
	DelImmutable()

	IsPersistent() bool
	SetPersistent()
	DelPersistent()

	IsRecursive() bool
	ProtectRecursive()
	UnprotectRecursive()
	TryProtectRecursive()
	TryUnProtectRecursive()
}

type ZendRefcounted struct {
	gc ZendRefcountedH
}

var _ IRefcounted = &ZendRefcounted{}

func (this *ZendRefcounted) GetGc() *ZendRefcountedH     { return &this.gc }
func (this *ZendRefcounted) SetGc(value ZendRefcountedH) { this.gc = value }

// Refcount

func (this *ZendRefcounted) GetRefcount() uint32 {
	return this.gc.refcount
}
func (this *ZendRefcounted) SetRefcount(rc uint32) uint32 {
	this.gc.refcount = rc
	return this.gc.refcount
}

func (this *ZendRefcounted) AddRefcount() uint32 {
	this.gc.refcount++
	return this.gc.refcount
}

func (this *ZendRefcounted) DelRefcount() uint32 {
	b.Assert(this.gc.refcount > 0)
	this.gc.refcount--
	return this.gc.refcount
}

func (this *ZendRefcounted) AddRefcountEx(rc uint32) uint32 {
	this.gc.refcount += rc
	return this.gc.refcount
}

func (this *ZendRefcounted) DelRefcountEx(rc uint32) uint32 {
	this.gc.refcount -= rc
	return this.gc.refcount
}

/**
 *  type_info 保存三个 flag 标识:
 *	type(4) + flags(6) + info(22)
 */
const GC_TYPE_MASK = 0xf
const GC_FLAGS_MASK = 0x3f0
const GC_FLAGS_SHIFT = 0

func (this *ZendRefcounted) GetGcTypeInfo() uint32 {
	return this.gc.u.type_info
}

func (this *ZendRefcounted) SetGcTypeInfo(typeInfo uint32) {
	this.gc.u.type_info = typeInfo
}

func (this *ZendRefcounted) GetGcType() uint8 {
	var typeInfo = this.GetGcTypeInfo()
	return uint8(typeInfo & GC_TYPE_MASK)
}

func (this *ZendRefcounted) GetGcFlags() uint32 {
	var typeInfo = this.GetGcTypeInfo()
	return (typeInfo & GC_FLAGS_MASK) >> GC_FLAGS_SHIFT
}

func (this *ZendRefcounted) AddGcFlags(flags uint32) {
	this.gc.u.type_info |= flags << GC_FLAGS_SHIFT
}

func (this *ZendRefcounted) DelGcFlags(flags uint32) {
	this.gc.u.type_info &^= flags << GC_FLAGS_SHIFT
}

func (this *ZendRefcounted) HasGcFlags(flags uint32) bool {
	var gcFlags = this.GetGcFlags()
	return b.FlagMatch(gcFlags, flags)
}

func (this *ZendRefcounted) SetCollectable() { this.AddGcFlags(GC_COLLECTABLE) }
func (this *ZendRefcounted) DelCollectable() { this.DelGcFlags(GC_COLLECTABLE) }

func (this *ZendRefcounted) IsImmutable() bool { return this.HasGcFlags(GC_IMMUTABLE) }
func (this *ZendRefcounted) SetImmutable()     { this.AddGcFlags(GC_IMMUTABLE) }
func (this *ZendRefcounted) DelImmutable()     { this.DelGcFlags(GC_IMMUTABLE) }

func (this *ZendRefcounted) IsPersistent() bool { return this.HasGcFlags(GC_PERSISTENT) }
func (this *ZendRefcounted) SetPersistent()     { this.AddGcFlags(GC_PERSISTENT) }
func (this *ZendRefcounted) DelPersistent()     { this.DelGcFlags(GC_PERSISTENT) }

func (this *ZendRefcounted) IsRecursive() bool   { return this.HasGcFlags(GC_PROTECTED) }
func (this *ZendRefcounted) ProtectRecursive()   { this.AddGcFlags(GC_PROTECTED) }
func (this *ZendRefcounted) UnprotectRecursive() { this.DelGcFlags(GC_PROTECTED) }
func (this *ZendRefcounted) TryProtectRecursive() {
	if !this.HasGcFlags(GC_IMMUTABLE) {
		this.AddGcFlags(GC_PROTECTED)
	}
}
func (this *ZendRefcounted) TryUnProtectRecursive() {
	if !this.HasGcFlags(GC_IMMUTABLE) {
		this.DelGcFlags(GC_PROTECTED)
	}
}
