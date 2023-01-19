// <<generate>>

package zend

/**
 * ZendValue
 */
type ZendValue struct /* union */ {
	lval    ZendLong
	dval    float64
	counted *ZendRefcounted
	str     *ZendString
	arr     *ZendArray
	obj     *ZendObject
	res     *ZendResource
	ref     *ZendReference
	ast     *ZendAstRef
	zv      *Zval
	ptr     any
	ce      *ZendClassEntry
	func_   *ZendFunction
	ww      struct {
		w1 uint32
		w2 uint32
	}
}

func (this ZendValue) GetLval() ZendLong                 { return this.lval }
func (this *ZendValue) SetLval(value ZendLong)           { this.lval = value }
func (this ZendValue) GetDval() float64                  { return this.dval }
func (this *ZendValue) SetDval(value float64)            { this.dval = value }
func (this ZendValue) GetCounted() *ZendRefcounted       { return this.counted }
func (this *ZendValue) SetCounted(value *ZendRefcounted) { this.counted = value }
func (this ZendValue) GetStr() *ZendString               { return this.str }
func (this *ZendValue) SetStr(value *ZendString)         { this.str = value }
func (this ZendValue) GetArr() *ZendArray                { return this.arr }
func (this *ZendValue) SetArr(value *ZendArray)          { this.arr = value }
func (this ZendValue) GetObj() *ZendObject               { return this.obj }
func (this *ZendValue) SetObj(value *ZendObject)         { this.obj = value }
func (this ZendValue) GetRes() *ZendResource             { return this.res }
func (this *ZendValue) SetRes(value *ZendResource)       { this.res = value }
func (this ZendValue) GetRef() *ZendReference            { return this.ref }
func (this *ZendValue) SetRef(value *ZendReference)      { this.ref = value }
func (this ZendValue) GetAst() *ZendAstRef               { return this.ast }
func (this *ZendValue) SetAst(value *ZendAstRef)         { this.ast = value }
func (this ZendValue) GetZv() *Zval                      { return this.zv }
func (this *ZendValue) SetZv(value *Zval)                { this.zv = value }
func (this ZendValue) GetPtr() any                       { return this.ptr }
func (this *ZendValue) SetPtr(value any)                 { this.ptr = value }
func (this ZendValue) GetCe() *ZendClassEntry            { return this.ce }
func (this *ZendValue) SetCe(value *ZendClassEntry)      { this.ce = value }
func (this ZendValue) GetFunc() *ZendFunction            { return this.func_ }
func (this *ZendValue) SetFunc(value *ZendFunction)      { this.func_ = value }
func (this ZendValue) GetW1() uint32                     { return this.ww.w1 }
func (this *ZendValue) SetW1(value uint32)               { this.ww.w1 = value }
func (this ZendValue) GetW2() uint32                     { return this.ww.w2 }
func (this *ZendValue) SetW2(value uint32)               { this.ww.w2 = value }

/**
 * Zval
 */
type Zval struct {
	value ZendValue
	u1    struct /* union */ {
		v struct {
			type_      ZendUchar
			type_flags ZendUchar
			u          struct /* union */ {
				extra uint16
			}
		}
		type_info uint32
	}
	u2 struct /* union */ {
		next           uint32
		cache_slot     uint32
		opline_num     uint32
		lineno         uint32
		num_args       uint32
		fe_pos         uint32
		fe_iter_idx    uint32
		access_flags   uint32
		property_guard uint32
		constant_flags uint32
		extra          uint32
	}
}

func (this Zval) GetValue() ZendValue            { return this.value }
func (this *Zval) SetValue(value ZendValue)      { this.value = value }
func (this Zval) GetType() ZendUchar             { return this.u1.v.type_ }
func (this *Zval) SetType(value ZendUchar)       { this.u1.v.type_ = value }
func (this Zval) GetTypeFlags() ZendUchar        { return this.u1.v.type_flags }
func (this *Zval) SetTypeFlags(value ZendUchar)  { this.u1.v.type_flags = value }
func (this Zval) GetU1VUExtra() uint16           { return this.u1.v.u.extra }
func (this *Zval) SetU1VUExtra(value uint16)     { this.u1.v.u.extra = value }
func (this Zval) GetTypeInfo() uint32            { return this.u1.type_info }
func (this *Zval) SetTypeInfo(value uint32)      { this.u1.type_info = value }
func (this Zval) GetNext() uint32                { return this.u2.next }
func (this *Zval) SetNext(value uint32)          { this.u2.next = value }
func (this Zval) GetCacheSlot() uint32           { return this.u2.cache_slot }
func (this *Zval) SetCacheSlot(value uint32)     { this.u2.cache_slot = value }
func (this Zval) GetOplineNum() uint32           { return this.u2.opline_num }
func (this *Zval) SetOplineNum(value uint32)     { this.u2.opline_num = value }
func (this Zval) GetLineno() uint32              { return this.u2.lineno }
func (this *Zval) SetLineno(value uint32)        { this.u2.lineno = value }
func (this Zval) GetNumArgs() uint32             { return this.u2.num_args }
func (this *Zval) SetNumArgs(value uint32)       { this.u2.num_args = value }
func (this Zval) GetFePos() uint32               { return this.u2.fe_pos }
func (this *Zval) SetFePos(value uint32)         { this.u2.fe_pos = value }
func (this Zval) GetFeIterIdx() uint32           { return this.u2.fe_iter_idx }
func (this *Zval) SetFeIterIdx(value uint32)     { this.u2.fe_iter_idx = value }
func (this Zval) GetAccessFlags() uint32         { return this.u2.access_flags }
func (this *Zval) SetAccessFlags(value uint32)   { this.u2.access_flags = value }
func (this Zval) GetPropertyGuard() uint32       { return this.u2.property_guard }
func (this *Zval) SetPropertyGuard(value uint32) { this.u2.property_guard = value }
func (this Zval) GetConstantFlags() uint32       { return this.u2.constant_flags }
func (this *Zval) SetConstantFlags(value uint32) { this.u2.constant_flags = value }
func (this Zval) GetU2Extra() uint32             { return this.u2.extra }
func (this *Zval) SetU2Extra(value uint32)       { this.u2.extra = value }

/**
 * ZendRefcountedH
 */
type ZendRefcountedH struct {
	refcount uint32
	u        struct /* union */ {
		type_info uint32
	}
}

func (this ZendRefcountedH) GetRefcount() uint32       { return this.refcount }
func (this *ZendRefcountedH) SetRefcount(value uint32) { this.refcount = value }
func (this ZendRefcountedH) GetTypeInfo() uint32       { return this.u.type_info }
func (this *ZendRefcountedH) SetTypeInfo(value uint32) { this.u.type_info = value }

/**
 * ZendRefcounted
 */
type ZendRefcounted struct {
	gc ZendRefcountedH
}

func (this ZendRefcounted) GetGc() ZendRefcountedH       { return this.gc }
func (this *ZendRefcounted) SetGc(value ZendRefcountedH) { this.gc = value }

/**
 * ZendString
 */
type ZendString struct {
	gc   ZendRefcountedH
	h    ZendUlong
	len_ int
	val  []byte
}

func (this ZendString) GetGc() ZendRefcountedH       { return this.gc }
func (this *ZendString) SetGc(value ZendRefcountedH) { this.gc = value }
func (this ZendString) GetH() ZendUlong              { return this.h }
func (this *ZendString) SetH(value ZendUlong)        { this.h = value }
func (this ZendString) GetLen() int                  { return this.len_ }
func (this *ZendString) SetLen(value int)            { this.len_ = value }
func (this ZendString) GetVal() []byte               { return this.val }
func (this *ZendString) SetVal(value []byte)         { this.val = value }

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	h   ZendUlong
	key *ZendString
}

func (this Bucket) GetVal() Zval              { return this.val }
func (this *Bucket) SetVal(value Zval)        { this.val = value }
func (this Bucket) GetH() ZendUlong           { return this.h }
func (this *Bucket) SetH(value ZendUlong)     { this.h = value }
func (this Bucket) GetKey() *ZendString       { return this.key }
func (this *Bucket) SetKey(value *ZendString) { this.key = value }

/**
 * ZendArray
 */
type ZendArray struct {
	gc ZendRefcountedH
	u  struct /* union */ {
		v struct {
			flags           ZendUchar
			_unused         ZendUchar
			nIteratorsCount ZendUchar
			_unused2        ZendUchar
		}
		flags uint32
	}
	nTableMask       uint32
	arData           *Bucket
	nNumUsed         uint32
	nNumOfElements   uint32
	nTableSize       uint32
	nInternalPointer uint32
	nNextFreeElement ZendLong
	pDestructor      DtorFuncT
}

func (this ZendArray) GetGc() ZendRefcountedH              { return this.gc }
func (this *ZendArray) SetGc(value ZendRefcountedH)        { this.gc = value }
func (this ZendArray) GetUVFlags() ZendUchar               { return this.u.v.flags }
func (this *ZendArray) SetUVFlags(value ZendUchar)         { this.u.v.flags = value }
func (this ZendArray) GetUnused() ZendUchar                { return this.u.v._unused }
func (this *ZendArray) SetUnused(value ZendUchar)          { this.u.v._unused = value }
func (this ZendArray) GetNIteratorsCount() ZendUchar       { return this.u.v.nIteratorsCount }
func (this *ZendArray) SetNIteratorsCount(value ZendUchar) { this.u.v.nIteratorsCount = value }
func (this ZendArray) GetUnused2() ZendUchar               { return this.u.v._unused2 }
func (this *ZendArray) SetUnused2(value ZendUchar)         { this.u.v._unused2 = value }
func (this ZendArray) GetUFlags() uint32                   { return this.u.flags }
func (this *ZendArray) SetUFlags(value uint32)             { this.u.flags = value }
func (this ZendArray) GetNTableMask() uint32               { return this.nTableMask }
func (this *ZendArray) SetNTableMask(value uint32)         { this.nTableMask = value }
func (this ZendArray) GetArData() *Bucket                  { return this.arData }
func (this *ZendArray) SetArData(value *Bucket)            { this.arData = value }
func (this ZendArray) GetNNumUsed() uint32                 { return this.nNumUsed }
func (this *ZendArray) SetNNumUsed(value uint32)           { this.nNumUsed = value }
func (this ZendArray) GetNNumOfElements() uint32           { return this.nNumOfElements }
func (this *ZendArray) SetNNumOfElements(value uint32)     { this.nNumOfElements = value }
func (this ZendArray) GetNTableSize() uint32               { return this.nTableSize }
func (this *ZendArray) SetNTableSize(value uint32)         { this.nTableSize = value }
func (this ZendArray) GetNInternalPointer() uint32         { return this.nInternalPointer }
func (this *ZendArray) SetNInternalPointer(value uint32)   { this.nInternalPointer = value }
func (this ZendArray) GetNNextFreeElement() ZendLong       { return this.nNextFreeElement }
func (this *ZendArray) SetNNextFreeElement(value ZendLong) { this.nNextFreeElement = value }
func (this ZendArray) GetPDestructor() DtorFuncT           { return this.pDestructor }
func (this *ZendArray) SetPDestructor(value DtorFuncT)     { this.pDestructor = value }

/**
 * HashTableIterator
 */
type HashTableIterator struct {
	ht  *HashTable
	pos HashPosition
}

func (this HashTableIterator) GetHt() *HashTable          { return this.ht }
func (this *HashTableIterator) SetHt(value *HashTable)    { this.ht = value }
func (this HashTableIterator) GetPos() HashPosition       { return this.pos }
func (this *HashTableIterator) SetPos(value HashPosition) { this.pos = value }

/**
 * ZendObject
 */
type ZendObject struct {
	gc               ZendRefcountedH
	handle           uint32
	ce               *ZendClassEntry
	handlers         *ZendObjectHandlers
	properties       *HashTable
	properties_table []Zval
}

func (this ZendObject) GetGc() ZendRefcountedH                 { return this.gc }
func (this *ZendObject) SetGc(value ZendRefcountedH)           { this.gc = value }
func (this ZendObject) GetHandle() uint32                      { return this.handle }
func (this *ZendObject) SetHandle(value uint32)                { this.handle = value }
func (this ZendObject) GetCe() *ZendClassEntry                 { return this.ce }
func (this *ZendObject) SetCe(value *ZendClassEntry)           { this.ce = value }
func (this ZendObject) GetHandlers() *ZendObjectHandlers       { return this.handlers }
func (this *ZendObject) SetHandlers(value *ZendObjectHandlers) { this.handlers = value }
func (this ZendObject) GetProperties() *HashTable              { return this.properties }
func (this *ZendObject) SetProperties(value *HashTable)        { this.properties = value }
func (this ZendObject) GetPropertiesTable() []Zval             { return this.properties_table }
func (this *ZendObject) SetPropertiesTable(value []Zval)       { this.properties_table = value }

/**
 * ZendResource
 */
type ZendResource struct {
	gc     ZendRefcountedH
	handle int
	type_  int
	ptr    any
}

func (this ZendResource) GetGc() ZendRefcountedH       { return this.gc }
func (this *ZendResource) SetGc(value ZendRefcountedH) { this.gc = value }
func (this ZendResource) GetHandle() int               { return this.handle }
func (this *ZendResource) SetHandle(value int)         { this.handle = value }
func (this ZendResource) GetType() int                 { return this.type_ }
func (this *ZendResource) SetType(value int)           { this.type_ = value }
func (this ZendResource) GetPtr() any                  { return this.ptr }
func (this *ZendResource) SetPtr(value any)            { this.ptr = value }

/**
 * ZendPropertyInfoList
 */
type ZendPropertyInfoList struct {
	num           int
	num_allocated int
	ptr           []*ZendPropertyInfo
}

func (this ZendPropertyInfoList) GetNum() int                       { return this.num }
func (this *ZendPropertyInfoList) SetNum(value int)                 { this.num = value }
func (this ZendPropertyInfoList) GetNumAllocated() int              { return this.num_allocated }
func (this *ZendPropertyInfoList) SetNumAllocated(value int)        { this.num_allocated = value }
func (this ZendPropertyInfoList) GetPtr() []*ZendPropertyInfo       { return this.ptr }
func (this *ZendPropertyInfoList) SetPtr(value []*ZendPropertyInfo) { this.ptr = value }

/**
 * ZendPropertyInfoSourceList
 */
type ZendPropertyInfoSourceList struct /* union */ {
	ptr  *ZendPropertyInfo
	list uintPtr
}

func (this ZendPropertyInfoSourceList) GetPtr() *ZendPropertyInfo       { return this.ptr }
func (this *ZendPropertyInfoSourceList) SetPtr(value *ZendPropertyInfo) { this.ptr = value }
func (this ZendPropertyInfoSourceList) GetList() uintPtr                { return this.list }
func (this *ZendPropertyInfoSourceList) SetList(value uintPtr)          { this.list = value }

/**
 * ZendReference
 */
type ZendReference struct {
	gc      ZendRefcountedH
	val     Zval
	sources ZendPropertyInfoSourceList
}

func (this ZendReference) GetGc() ZendRefcountedH                       { return this.gc }
func (this *ZendReference) SetGc(value ZendRefcountedH)                 { this.gc = value }
func (this ZendReference) GetVal() Zval                                 { return this.val }
func (this *ZendReference) SetVal(value Zval)                           { this.val = value }
func (this ZendReference) GetSources() ZendPropertyInfoSourceList       { return this.sources }
func (this *ZendReference) SetSources(value ZendPropertyInfoSourceList) { this.sources = value }

/**
 * ZendAstRef
 */
type ZendAstRef struct {
	gc ZendRefcountedH
}

func (this ZendAstRef) GetGc() ZendRefcountedH       { return this.gc }
func (this *ZendAstRef) SetGc(value ZendRefcountedH) { this.gc = value }
