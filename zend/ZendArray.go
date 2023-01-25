// <<generate>>

package zend

/**
 * ZendHashKey
 */
type ZendHashKey struct {
	h   ZendUlong
	key *ZendString
}

func (this ZendHashKey) GetH() ZendUlong           { return this.h }
func (this *ZendHashKey) SetH(value ZendUlong)     { this.h = value }
func (this ZendHashKey) GetKey() *ZendString       { return this.key }
func (this *ZendHashKey) SetKey(value *ZendString) { this.key = value }

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
func (this ZendArray) GetFlags() ZendUchar                 { return this.u.v.flags }
func (this *ZendArray) SetFlags(value ZendUchar)           { this.u.v.flags = value }
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

/* ZendArray.u.v.flags */
func (this *ZendArray) AddFlags(value ZendUchar)     { this.u.v.flags |= value }
func (this *ZendArray) SubFlags(value ZendUchar)     { this.u.v.flags &^= value }
func (this ZendArray) HasFlags(value ZendUchar) bool { return this.u.v.flags&value != 0 }
func (this *ZendArray) SwitchFlags(value ZendUchar, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendArray) isHasEmptyInd() bool         { return this.HasFlags(HASH_FLAG_HAS_EMPTY_IND) }
func (this *ZendArray) setIsHasEmptyInd(cond bool) { this.SwitchFlags(HASH_FLAG_HAS_EMPTY_IND, cond) }

/* ZendArray.u.flags */
func (this *ZendArray) AddUFlags(value uint32)     { this.u.flags |= value }
func (this *ZendArray) SubUFlags(value uint32)     { this.u.flags &^= value }
func (this ZendArray) HasUFlags(value uint32) bool { return this.u.flags&value != 0 }
func (this *ZendArray) SwitchUFlags(value uint32, cond bool) {
	if cond {
		this.AddUFlags(value)
	} else {
		this.SubUFlags(value)
	}
}
func (this ZendArray) isApplyProtection() bool { return this.HasUFlags(HASH_FLAG_APPLY_PROTECTION) }
func (this ZendArray) isPacked() bool          { return this.HasUFlags(HASH_FLAG_PACKED) }
func (this ZendArray) isInitialized() bool     { return this.HasUFlags(HASH_FLAG_INITIALIZED) }
func (this *ZendArray) setIsApplyProtection(cond bool) {
	this.SwitchUFlags(HASH_FLAG_APPLY_PROTECTION, cond)
}
func (this *ZendArray) setIsPacked(cond bool)      { this.SwitchUFlags(HASH_FLAG_PACKED, cond) }
func (this *ZendArray) setIsInitialized(cond bool) { this.SwitchUFlags(HASH_FLAG_INITIALIZED, cond) }
