// <<generate>>

package zend

/**
 * ZendGcStatus
 */
type ZendGcStatus struct {
	runs      uint32
	collected uint32
	threshold uint32
	num_roots uint32
}

func (this ZendGcStatus) GetRuns() uint32            { return this.runs }
func (this *ZendGcStatus) SetRuns(value uint32)      { this.runs = value }
func (this ZendGcStatus) GetCollected() uint32       { return this.collected }
func (this *ZendGcStatus) SetCollected(value uint32) { this.collected = value }
func (this ZendGcStatus) GetThreshold() uint32       { return this.threshold }
func (this *ZendGcStatus) SetThreshold(value uint32) { this.threshold = value }
func (this ZendGcStatus) GetNumRoots() uint32        { return this.num_roots }
func (this *ZendGcStatus) SetNumRoots(value uint32)  { this.num_roots = value }

/**
 * GcRootBuffer
 */
type GcRootBuffer struct {
	ref *ZendRefcounted
}

func (this GcRootBuffer) GetRef() *ZendRefcounted       { return this.ref }
func (this *GcRootBuffer) SetRef(value *ZendRefcounted) { this.ref = value }

/**
 * ZendGcGlobals
 */
type ZendGcGlobals struct {
	buf          *GcRootBuffer
	gc_enabled   ZendBool
	gc_active    ZendBool
	gc_protected ZendBool
	gc_full      ZendBool
	unused       uint32
	first_unused uint32
	gc_threshold uint32
	buf_size     uint32
	num_roots    uint32
	gc_runs      uint32
	collected    uint32
}

func (this ZendGcGlobals) GetBuf() *GcRootBuffer          { return this.buf }
func (this *ZendGcGlobals) SetBuf(value *GcRootBuffer)    { this.buf = value }
func (this ZendGcGlobals) GetGcEnabled() ZendBool         { return this.gc_enabled }
func (this *ZendGcGlobals) SetGcEnabled(value ZendBool)   { this.gc_enabled = value }
func (this ZendGcGlobals) GetGcActive() ZendBool          { return this.gc_active }
func (this *ZendGcGlobals) SetGcActive(value ZendBool)    { this.gc_active = value }
func (this ZendGcGlobals) GetGcProtected() ZendBool       { return this.gc_protected }
func (this *ZendGcGlobals) SetGcProtected(value ZendBool) { this.gc_protected = value }
func (this ZendGcGlobals) GetGcFull() ZendBool            { return this.gc_full }
func (this *ZendGcGlobals) SetGcFull(value ZendBool)      { this.gc_full = value }
func (this ZendGcGlobals) GetUnused() uint32              { return this.unused }
func (this *ZendGcGlobals) SetUnused(value uint32)        { this.unused = value }
func (this ZendGcGlobals) GetFirstUnused() uint32         { return this.first_unused }
func (this *ZendGcGlobals) SetFirstUnused(value uint32)   { this.first_unused = value }
func (this ZendGcGlobals) GetGcThreshold() uint32         { return this.gc_threshold }
func (this *ZendGcGlobals) SetGcThreshold(value uint32)   { this.gc_threshold = value }
func (this ZendGcGlobals) GetBufSize() uint32             { return this.buf_size }
func (this *ZendGcGlobals) SetBufSize(value uint32)       { this.buf_size = value }
func (this ZendGcGlobals) GetNumRoots() uint32            { return this.num_roots }
func (this *ZendGcGlobals) SetNumRoots(value uint32)      { this.num_roots = value }
func (this ZendGcGlobals) GetGcRuns() uint32              { return this.gc_runs }
func (this *ZendGcGlobals) SetGcRuns(value uint32)        { this.gc_runs = value }
func (this ZendGcGlobals) GetCollected() uint32           { return this.collected }
func (this *ZendGcGlobals) SetCollected(value uint32)     { this.collected = value }

/**
 * GcStack
 */
type GcStack struct {
	prev *GcStack
	next *GcStack
	data []*ZendRefcounted
}

func (this GcStack) GetPrev() *GcStack                { return this.prev }
func (this *GcStack) SetPrev(value *GcStack)          { this.prev = value }
func (this GcStack) GetNext() *GcStack                { return this.next }
func (this *GcStack) SetNext(value *GcStack)          { this.next = value }
func (this GcStack) GetData() []*ZendRefcounted       { return this.data }
func (this *GcStack) SetData(value []*ZendRefcounted) { this.data = value }
