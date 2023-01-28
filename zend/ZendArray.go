// <<generate>>

package zend

/**
 * ZendHashKey
 */
type ZendHashKey struct {
	h   ZendUlong
	key *ZendString
}

func (this *ZendHashKey) GetH() ZendUlong          { return this.h }
func (this *ZendHashKey) SetH(value ZendUlong)     { this.h = value }
func (this *ZendHashKey) GetKey() *ZendString      { return this.key }
func (this *ZendHashKey) SetKey(value *ZendString) { this.key = value }

/**
 * Bucket
 */
type Bucket struct {
	val Zval
	h   ZendUlong
	key *ZendString
}

func (this *Bucket) GetVal() Zval             { return this.val }
func (this *Bucket) SetVal(value Zval)        { this.val = value }
func (this *Bucket) GetH() ZendUlong          { return this.h }
func (this *Bucket) SetH(value ZendUlong)     { this.h = value }
func (this *Bucket) GetKey() *ZendString      { return this.key }
func (this *Bucket) SetKey(value *ZendString) { this.key = value }

/**
 * ZendArray
 * HashTable Data Layout
 * =====================
 *
 *                 +=============================+
 *                 | HT_HASH(ht, ht->nTableMask) |
 *                 | ...                         |
 *                 | HT_HASH(ht, -1)             |
 *                 +-----------------------------+
 * ht->arData ---> | Bucket[0]                   |
 *                 | ...                         |
 *                 | Bucket[ht->nTableSize-1]    |
 *                 +=============================+
 */
type HashTable = ZendArray
type ZendArray struct {
	ZendRefcounted
	u struct /* union */ {
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

var _ IRefcounted = &ZendArray{}

func (this *ZendArray) GetNIteratorsCount() ZendUchar      { return this.u.v.nIteratorsCount }
func (this *ZendArray) SetNIteratorsCount(value ZendUchar) { this.u.v.nIteratorsCount = value }
func (this *ZendArray) GetNTableMask() uint32              { return this.nTableMask }
func (this *ZendArray) SetNTableMask(value uint32)         { this.nTableMask = value }
func (this *ZendArray) GetArData() *Bucket                 { return this.arData }
func (this *ZendArray) SetArData(value *Bucket)            { this.arData = value }
func (this *ZendArray) GetNNumUsed() uint32                { return this.nNumUsed }
func (this *ZendArray) SetNNumUsed(value uint32)           { this.nNumUsed = value }
func (this *ZendArray) GetNNumOfElements() uint32          { return this.nNumOfElements }
func (this *ZendArray) SetNNumOfElements(value uint32)     { this.nNumOfElements = value }
func (this *ZendArray) GetNTableSize() uint32              { return this.nTableSize }
func (this *ZendArray) SetNTableSize(value uint32)         { this.nTableSize = value }
func (this *ZendArray) GetNInternalPointer() uint32        { return this.nInternalPointer }
func (this *ZendArray) SetNInternalPointer(value uint32)   { this.nInternalPointer = value }
func (this *ZendArray) GetNNextFreeElement() ZendLong      { return this.nNextFreeElement }
func (this *ZendArray) SetNNextFreeElement(value ZendLong) { this.nNextFreeElement = value }
func (this *ZendArray) GetPDestructor() DtorFuncT          { return this.pDestructor }
func (this *ZendArray) SetPDestructor(value DtorFuncT)     { this.pDestructor = value }

/* ZendArray.u.v.flags */
func (this *ZendArray) GetFlags() ZendUchar           { return this.u.v.flags }
func (this *ZendArray) SetFlags(value ZendUchar)      { this.u.v.flags = value }
func (this *ZendArray) AddFlags(value ZendUchar)      { this.u.v.flags |= value }
func (this *ZendArray) SubFlags(value ZendUchar)      { this.u.v.flags &^= value }
func (this *ZendArray) HasFlags(value ZendUchar) bool { return this.u.v.flags&value != 0 }
func (this *ZendArray) SwitchFlags(value ZendUchar, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}

const HASH_FLAG_PACKED = 1 << 2
const HASH_FLAG_UNINITIALIZED = 1 << 3
const HASH_FLAG_STATIC_KEYS = 1 << 4
const HASH_FLAG_HAS_EMPTY_IND = 1 << 5

func (this *ZendArray) IsPacked() bool        { return this.HasFlags(HASH_FLAG_PACKED) }
func (this *ZendArray) IsUninitialized() bool { return this.HasFlags(HASH_FLAG_UNINITIALIZED) }
func (this *ZendArray) IsStaticKeys() bool    { return this.HasFlags(HASH_FLAG_STATIC_KEYS) }
func (this *ZendArray) IsHasEmptyInd() bool   { return this.HasFlags(HASH_FLAG_HAS_EMPTY_IND) }
func (this *ZendArray) SetIsPacked()          { this.AddFlags(HASH_FLAG_PACKED) }
func (this *ZendArray) SetIsUninitialized()   { this.AddFlags(HASH_FLAG_UNINITIALIZED) }
func (this *ZendArray) SetIsStaticKeys()      { this.AddFlags(HASH_FLAG_STATIC_KEYS) }
func (this *ZendArray) SetIsHasEmptyInd()     { this.AddFlags(HASH_FLAG_HAS_EMPTY_IND) }

/* ZendArray.u.flags */
func (this *ZendArray) GetUFlags() uint32           { return this.u.flags }
func (this *ZendArray) SetUFlags(value uint32)      { this.u.flags = value }
func (this *ZendArray) AddUFlags(value uint32)      { this.u.flags |= value }
func (this *ZendArray) SubUFlags(value uint32)      { this.u.flags &^= value }
func (this *ZendArray) HasUFlags(value uint32) bool { return this.u.flags&value != 0 }
func (this *ZendArray) SwitchUFlags(value uint32, cond bool) {
	if cond {
		this.AddUFlags(value)
	} else {
		this.SubUFlags(value)
	}
}
