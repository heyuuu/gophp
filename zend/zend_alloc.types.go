// <<generate>>

package zend

/**
 * ZendLeakInfo
 */
type ZendLeakInfo struct {
	addr          any
	size          int
	filename      *byte
	orig_filename *byte
	lineno        uint32
	orig_lineno   uint32
}

func (this *ZendLeakInfo) GetAddr() any                { return this.addr }
func (this *ZendLeakInfo) SetAddr(value any)           { this.addr = value }
func (this *ZendLeakInfo) GetSize() int                { return this.size }
func (this *ZendLeakInfo) SetSize(value int)           { this.size = value }
func (this *ZendLeakInfo) GetFilename() *byte          { return this.filename }
func (this *ZendLeakInfo) SetFilename(value *byte)     { this.filename = value }
func (this *ZendLeakInfo) GetOrigFilename() *byte      { return this.orig_filename }
func (this *ZendLeakInfo) SetOrigFilename(value *byte) { this.orig_filename = value }
func (this *ZendLeakInfo) GetLineno() uint32           { return this.lineno }
func (this *ZendLeakInfo) SetLineno(value uint32)      { this.lineno = value }
func (this *ZendLeakInfo) GetOrigLineno() uint32       { return this.orig_lineno }
func (this *ZendLeakInfo) SetOrigLineno(value uint32)  { this.orig_lineno = value }

/**
 * ZendMmHandlers
 */
type ZendMmHandlers struct {
	chunk_alloc    ZendMmChunkAllocT
	chunk_free     ZendMmChunkFreeT
	chunk_truncate ZendMmChunkTruncateT
	chunk_extend   ZendMmChunkExtendT
}

func (this *ZendMmHandlers) GetChunkAlloc() ZendMmChunkAllocT            { return this.chunk_alloc }
func (this *ZendMmHandlers) SetChunkAlloc(value ZendMmChunkAllocT)       { this.chunk_alloc = value }
func (this *ZendMmHandlers) GetChunkFree() ZendMmChunkFreeT              { return this.chunk_free }
func (this *ZendMmHandlers) SetChunkFree(value ZendMmChunkFreeT)         { this.chunk_free = value }
func (this *ZendMmHandlers) GetChunkTruncate() ZendMmChunkTruncateT      { return this.chunk_truncate }
func (this *ZendMmHandlers) SetChunkTruncate(value ZendMmChunkTruncateT) { this.chunk_truncate = value }
func (this *ZendMmHandlers) GetChunkExtend() ZendMmChunkExtendT          { return this.chunk_extend }
func (this *ZendMmHandlers) SetChunkExtend(value ZendMmChunkExtendT)     { this.chunk_extend = value }

/**
 * ZendMmStorage
 */
type ZendMmStorage struct {
	handlers ZendMmHandlers
	data     any
}

func (this *ZendMmStorage) GetHandlers() ZendMmHandlers      { return this.handlers }
func (this *ZendMmStorage) SetHandlers(value ZendMmHandlers) { this.handlers = value }
func (this *ZendMmStorage) GetData() any                     { return this.data }
func (this *ZendMmStorage) SetData(value any)                { this.data = value }

/**
 * ZendMmHeap
 */
type ZendMmHeap struct {
	use_custom_heap             int
	storage                     *ZendMmStorage
	size                        int
	peak                        int
	free_slot                   []*ZendMmFreeSlot
	real_size                   int
	real_peak                   int
	limit                       int
	overflow                    int
	huge_list                   *ZendMmHugeList
	main_chunk                  *ZendMmChunk
	cached_chunks               *ZendMmChunk
	chunks_count                int
	peak_chunks_count           int
	cached_chunks_count         int
	avg_chunks_count            float64
	last_chunks_delete_boundary int
	last_chunks_delete_count    int
	custom_heap                 struct /* union */ {
		std struct {
			_malloc  func(int) any
			_free    func(any)
			_realloc func(any, int) any
		}
		debug struct {
			_malloc  func(int) any
			_free    func(any)
			_realloc func(any, int) any
		}
	}
	tracked_allocs *HashTable
}

func (this *ZendMmHeap) GetUseCustomHeap() int               { return this.use_custom_heap }
func (this *ZendMmHeap) SetUseCustomHeap(value int)          { this.use_custom_heap = value }
func (this *ZendMmHeap) GetStorage() *ZendMmStorage          { return this.storage }
func (this *ZendMmHeap) SetStorage(value *ZendMmStorage)     { this.storage = value }
func (this *ZendMmHeap) GetSize() int                        { return this.size }
func (this *ZendMmHeap) SetSize(value int)                   { this.size = value }
func (this *ZendMmHeap) GetPeak() int                        { return this.peak }
func (this *ZendMmHeap) SetPeak(value int)                   { this.peak = value }
func (this *ZendMmHeap) GetFreeSlot() []*ZendMmFreeSlot      { return this.free_slot }
func (this *ZendMmHeap) SetFreeSlot(value []*ZendMmFreeSlot) { this.free_slot = value }
func (this *ZendMmHeap) GetRealSize() int                    { return this.real_size }
func (this *ZendMmHeap) SetRealSize(value int)               { this.real_size = value }
func (this *ZendMmHeap) GetRealPeak() int                    { return this.real_peak }
func (this *ZendMmHeap) SetRealPeak(value int)               { this.real_peak = value }
func (this *ZendMmHeap) GetLimit() int                       { return this.limit }
func (this *ZendMmHeap) SetLimit(value int)                  { this.limit = value }
func (this *ZendMmHeap) GetOverflow() int                    { return this.overflow }
func (this *ZendMmHeap) SetOverflow(value int)               { this.overflow = value }
func (this *ZendMmHeap) GetHugeList() *ZendMmHugeList        { return this.huge_list }
func (this *ZendMmHeap) SetHugeList(value *ZendMmHugeList)   { this.huge_list = value }
func (this *ZendMmHeap) GetMainChunk() *ZendMmChunk          { return this.main_chunk }
func (this *ZendMmHeap) SetMainChunk(value *ZendMmChunk)     { this.main_chunk = value }
func (this *ZendMmHeap) GetCachedChunks() *ZendMmChunk       { return this.cached_chunks }
func (this *ZendMmHeap) SetCachedChunks(value *ZendMmChunk)  { this.cached_chunks = value }
func (this *ZendMmHeap) GetChunksCount() int                 { return this.chunks_count }
func (this *ZendMmHeap) SetChunksCount(value int)            { this.chunks_count = value }
func (this *ZendMmHeap) GetPeakChunksCount() int             { return this.peak_chunks_count }
func (this *ZendMmHeap) SetPeakChunksCount(value int)        { this.peak_chunks_count = value }
func (this *ZendMmHeap) GetCachedChunksCount() int           { return this.cached_chunks_count }
func (this *ZendMmHeap) SetCachedChunksCount(value int)      { this.cached_chunks_count = value }
func (this *ZendMmHeap) GetAvgChunksCount() float64          { return this.avg_chunks_count }
func (this *ZendMmHeap) SetAvgChunksCount(value float64)     { this.avg_chunks_count = value }
func (this *ZendMmHeap) GetLastChunksDeleteBoundary() int    { return this.last_chunks_delete_boundary }
func (this *ZendMmHeap) SetLastChunksDeleteBoundary(value int) {
	this.last_chunks_delete_boundary = value
}
func (this *ZendMmHeap) GetLastChunksDeleteCount() int         { return this.last_chunks_delete_count }
func (this *ZendMmHeap) SetLastChunksDeleteCount(value int)    { this.last_chunks_delete_count = value }
func (this *ZendMmHeap) GetCustomHeapStdMalloc() func(int) any { return this.custom_heap.std._malloc }
func (this *ZendMmHeap) SetCustomHeapStdMalloc(value func(int) any) {
	this.custom_heap.std._malloc = value
}
func (this *ZendMmHeap) GetCustomHeapStdFree() func(any)      { return this.custom_heap.std._free }
func (this *ZendMmHeap) SetCustomHeapStdFree(value func(any)) { this.custom_heap.std._free = value }
func (this *ZendMmHeap) GetCustomHeapStdRealloc() func(any, int) any {
	return this.custom_heap.std._realloc
}
func (this *ZendMmHeap) SetCustomHeapStdRealloc(value func(any, int) any) {
	this.custom_heap.std._realloc = value
}
func (this *ZendMmHeap) GetCustomHeapDebugMalloc() func(int) any {
	return this.custom_heap.debug._malloc
}
func (this *ZendMmHeap) SetCustomHeapDebugMalloc(value func(int) any) {
	this.custom_heap.debug._malloc = value
}
func (this *ZendMmHeap) GetCustomHeapDebugFree() func(any)      { return this.custom_heap.debug._free }
func (this *ZendMmHeap) SetCustomHeapDebugFree(value func(any)) { this.custom_heap.debug._free = value }
func (this *ZendMmHeap) GetCustomHeapDebugRealloc() func(any, int) any {
	return this.custom_heap.debug._realloc
}
func (this *ZendMmHeap) SetCustomHeapDebugRealloc(value func(any, int) any) {
	this.custom_heap.debug._realloc = value
}
func (this *ZendMmHeap) GetTrackedAllocs() *HashTable      { return this.tracked_allocs }
func (this *ZendMmHeap) SetTrackedAllocs(value *HashTable) { this.tracked_allocs = value }

/**
 * ZendMmChunk
 */
type ZendMmChunk struct {
	heap       *ZendMmHeap
	next       *ZendMmChunk
	prev       *ZendMmChunk
	free_pages uint32
	free_tail  uint32
	num        uint32
	reserve    []byte
	heap_slot  ZendMmHeap
	free_map   ZendMmPageMap
	map_       []ZendMmPageInfo
}

func (this *ZendMmChunk) GetHeap() *ZendMmHeap           { return this.heap }
func (this *ZendMmChunk) SetHeap(value *ZendMmHeap)      { this.heap = value }
func (this *ZendMmChunk) GetNext() *ZendMmChunk          { return this.next }
func (this *ZendMmChunk) SetNext(value *ZendMmChunk)     { this.next = value }
func (this *ZendMmChunk) GetPrev() *ZendMmChunk          { return this.prev }
func (this *ZendMmChunk) SetPrev(value *ZendMmChunk)     { this.prev = value }
func (this *ZendMmChunk) GetFreePages() uint32           { return this.free_pages }
func (this *ZendMmChunk) SetFreePages(value uint32)      { this.free_pages = value }
func (this *ZendMmChunk) GetFreeTail() uint32            { return this.free_tail }
func (this *ZendMmChunk) SetFreeTail(value uint32)       { this.free_tail = value }
func (this *ZendMmChunk) GetNum() uint32                 { return this.num }
func (this *ZendMmChunk) SetNum(value uint32)            { this.num = value }
func (this *ZendMmChunk) GetReserve() []byte             { return this.reserve }
func (this *ZendMmChunk) SetReserve(value []byte)        { this.reserve = value }
func (this *ZendMmChunk) GetHeapSlot() ZendMmHeap        { return this.heap_slot }
func (this *ZendMmChunk) SetHeapSlot(value ZendMmHeap)   { this.heap_slot = value }
func (this *ZendMmChunk) GetFreeMap() ZendMmPageMap      { return this.free_map }
func (this *ZendMmChunk) SetFreeMap(value ZendMmPageMap) { this.free_map = value }
func (this *ZendMmChunk) GetMap() []ZendMmPageInfo       { return this.map_ }
func (this *ZendMmChunk) SetMap(value []ZendMmPageInfo)  { this.map_ = value }

/**
 * ZendMmPage
 */
type ZendMmPage struct {
	bytes []byte
}

func (this *ZendMmPage) GetBytes() []byte      { return this.bytes }
func (this *ZendMmPage) SetBytes(value []byte) { this.bytes = value }

/**
 * ZendMmBin
 */
type ZendMmBin struct {
	bytes []byte
}

func (this *ZendMmBin) GetBytes() []byte      { return this.bytes }
func (this *ZendMmBin) SetBytes(value []byte) { this.bytes = value }

/**
 * ZendMmFreeSlot
 */
type ZendMmFreeSlot struct {
	next_free_slot *ZendMmFreeSlot
}

func (this *ZendMmFreeSlot) GetNextFreeSlot() *ZendMmFreeSlot      { return this.next_free_slot }
func (this *ZendMmFreeSlot) SetNextFreeSlot(value *ZendMmFreeSlot) { this.next_free_slot = value }

/**
 * ZendMmHugeList
 */
type ZendMmHugeList struct {
	ptr  any
	size int
	next *ZendMmHugeList
}

func (this *ZendMmHugeList) GetPtr() any                   { return this.ptr }
func (this *ZendMmHugeList) SetPtr(value any)              { this.ptr = value }
func (this *ZendMmHugeList) GetSize() int                  { return this.size }
func (this *ZendMmHugeList) SetSize(value int)             { this.size = value }
func (this *ZendMmHugeList) GetNext() *ZendMmHugeList      { return this.next }
func (this *ZendMmHugeList) SetNext(value *ZendMmHugeList) { this.next = value }

/**
 * ZendAllocGlobals
 */
type ZendAllocGlobals struct {
	mm_heap *ZendMmHeap
}

func (this *ZendAllocGlobals) GetMmHeap() *ZendMmHeap      { return this.mm_heap }
func (this *ZendAllocGlobals) SetMmHeap(value *ZendMmHeap) { this.mm_heap = value }
