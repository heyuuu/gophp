// <<generate>>

package zend

/**
 * HashTableIterator
 */
type HashTableIterator struct {
	ht  *HashTable
	pos HashPosition
}

func (this *HashTableIterator) GetHt() *HashTable         { return this.ht }
func (this *HashTableIterator) SetHt(value *HashTable)    { this.ht = value }
func (this *HashTableIterator) GetPos() HashPosition      { return this.pos }
func (this *HashTableIterator) SetPos(value HashPosition) { this.pos = value }

/**
 * ZendObject
 */
type ZendObject struct {
	baseZendRefcounted
	handle           uint32
	ce               *ZendClassEntry
	handlers         *ZendObjectHandlers
	properties       *HashTable
	properties_table []Zval
}

var _ ZendRefcounted = &ZendObject{}

func (this *ZendObject) GetHandle() uint32                     { return this.handle }
func (this *ZendObject) SetHandle(value uint32)                { this.handle = value }
func (this *ZendObject) GetCe() *ZendClassEntry                { return this.ce }
func (this *ZendObject) SetCe(value *ZendClassEntry)           { this.ce = value }
func (this *ZendObject) GetHandlers() *ZendObjectHandlers      { return this.handlers }
func (this *ZendObject) SetHandlers(value *ZendObjectHandlers) { this.handlers = value }
func (this *ZendObject) GetProperties() *HashTable             { return this.properties }
func (this *ZendObject) SetProperties(value *HashTable)        { this.properties = value }
func (this *ZendObject) GetPropertiesTable() []Zval            { return this.properties_table }
func (this *ZendObject) SetPropertiesTable(value []Zval)       { this.properties_table = value }

/**
 * ZendResource
 */
type ZendResource struct {
	baseZendRefcounted
	handle int
	type_  int
	ptr    any
}

var _ ZendRefcounted = &ZendResource{}

func (this *ZendResource) GetHandle() int      { return this.handle }
func (this *ZendResource) SetHandle(value int) { this.handle = value }
func (this *ZendResource) GetType() int        { return this.type_ }
func (this *ZendResource) SetType(value int)   { this.type_ = value }
func (this *ZendResource) GetPtr() any         { return this.ptr }
func (this *ZendResource) SetPtr(value any)    { this.ptr = value }

/**
 * ZendPropertyInfoList
 */
type ZendPropertyInfoList struct {
	num           int
	num_allocated int
	ptr           []*ZendPropertyInfo
}

func (this *ZendPropertyInfoList) GetNum() int                      { return this.num }
func (this *ZendPropertyInfoList) SetNum(value int)                 { this.num = value }
func (this *ZendPropertyInfoList) GetNumAllocated() int             { return this.num_allocated }
func (this *ZendPropertyInfoList) SetNumAllocated(value int)        { this.num_allocated = value }
func (this *ZendPropertyInfoList) GetPtr() []*ZendPropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoList) SetPtr(value []*ZendPropertyInfo) { this.ptr = value }

/**
 * ZendPropertyInfoSourceList
 */
type ZendPropertyInfoSourceList struct /* union */ {
	ptr  *ZendPropertyInfo
	list uintptr
}

func (this *ZendPropertyInfoSourceList) GetPtr() *ZendPropertyInfo      { return this.ptr }
func (this *ZendPropertyInfoSourceList) SetPtr(value *ZendPropertyInfo) { this.ptr = value }
func (this *ZendPropertyInfoSourceList) GetList() uintptr               { return this.list }
func (this *ZendPropertyInfoSourceList) SetList(value uintptr)          { this.list = value }

/**
 * ZendReference
 */
type ZendReference struct {
	baseZendRefcounted
	val     Zval
	sources ZendPropertyInfoSourceList
}

var _ ZendRefcounted = &ZendReference{}

func (this *ZendReference) GetVal() Zval                                { return this.val }
func (this *ZendReference) SetVal(value Zval)                           { this.val = value }
func (this *ZendReference) GetSources() ZendPropertyInfoSourceList      { return this.sources }
func (this *ZendReference) SetSources(value ZendPropertyInfoSourceList) { this.sources = value }

/**
 * ZendAstRef
 */
type ZendAstRef struct {
	baseZendRefcounted
}

var _ ZendRefcounted = &ZendReference{}
