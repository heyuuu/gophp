// <<generate>>

package standard

/**
 * EncToUniStage2
 */
type EncToUniStage2 struct {
	uni_cp []uint16
}

func MakeEncToUniStage2(uni_cp []uint16) EncToUniStage2 {
	return EncToUniStage2{
		uni_cp: uni_cp,
	}
}
func (this *EncToUniStage2) GetUniCp() []uint16 { return this.uni_cp }

// func (this *EncToUniStage2) SetUniCp(value []uint16) { this.uni_cp = value }

/**
 * EncToUni
 */
type EncToUni struct {
	inner []*EncToUniStage2
}

func MakeEncToUni(inner []*EncToUniStage2) EncToUni {
	return EncToUni{
		inner: inner,
	}
}
func (this *EncToUni) GetInner() []*EncToUniStage2 { return this.inner }

// func (this *EncToUni) SetInner(value []*EncToUniStage2) { this.inner = value }

/**
 * UniToEnc
 */
type UniToEnc struct {
	un_code_point uint16
	cs_code       uint8
}

func MakeUniToEnc(un_code_point uint16, cs_code uint8) UniToEnc {
	return UniToEnc{
		un_code_point: un_code_point,
		cs_code:       cs_code,
	}
}
func (this *UniToEnc) GetUnCodePoint() uint16 { return this.un_code_point }

// func (this *UniToEnc) SetUnCodePoint(value uint16) { this.un_code_point = value }
func (this *UniToEnc) GetCsCode() uint8 { return this.cs_code }

// func (this *UniToEnc) SetCsCode(value uint8) { this.cs_code = value }

/**
 * EntityMulticodepointRow
 */
type EntityMulticodepointRow struct /* union */ {
	leading_entry struct {
		default_entity     *byte
		size               unsigned
		default_entity_len uint16
	}
	normal_entry struct {
		entity     *byte
		second_cp  unsigned
		entity_len uint16
	}
}

func (this *EntityMulticodepointRow) GetDefaultEntity() *byte {
	return this.leading_entry.default_entity
}

// func (this *EntityMulticodepointRow) SetDefaultEntity(value *byte) { this.leading_entry.default_entity = value }
func (this *EntityMulticodepointRow) GetSize() unsigned { return this.leading_entry.size }

// func (this *EntityMulticodepointRow) SetSize(value unsigned) { this.leading_entry.size = value }
func (this *EntityMulticodepointRow) GetDefaultEntityLen() uint16 {
	return this.leading_entry.default_entity_len
}

// func (this *EntityMulticodepointRow) SetDefaultEntityLen(value uint16) { this.leading_entry.default_entity_len = value }
func (this *EntityMulticodepointRow) GetEntity() *byte { return this.normal_entry.entity }

// func (this *EntityMulticodepointRow) SetEntity(value *byte) { this.normal_entry.entity = value }
func (this *EntityMulticodepointRow) GetSecondCp() unsigned { return this.normal_entry.second_cp }

// func (this *EntityMulticodepointRow) SetSecondCp(value unsigned) { this.normal_entry.second_cp = value }
func (this *EntityMulticodepointRow) GetEntityLen() uint16 { return this.normal_entry.entity_len }

// func (this *EntityMulticodepointRow) SetEntityLen(value uint16) { this.normal_entry.entity_len = value }

/**
 * EntityStage3Row
 */
type EntityStage3Row struct {
	ambiguous byte
	data      struct /* union */ {
		ent struct {
			entity     *byte
			entity_len uint16
		}
		multicodepoint_table *EntityMulticodepointRow
	}
}

func (this *EntityStage3Row) GetAmbiguous() byte { return this.ambiguous }

// func (this *EntityStage3Row) SetAmbiguous(value byte) { this.ambiguous = value }
func (this *EntityStage3Row) GetEntity() *byte { return this.data.ent.entity }

// func (this *EntityStage3Row) SetEntity(value *byte) { this.data.ent.entity = value }
func (this *EntityStage3Row) GetEntityLen() uint16 { return this.data.ent.entity_len }

// func (this *EntityStage3Row) SetEntityLen(value uint16) { this.data.ent.entity_len = value }
func (this *EntityStage3Row) GetMulticodepointTable() *EntityMulticodepointRow {
	return this.data.multicodepoint_table
}

// func (this *EntityStage3Row) SetMulticodepointTable(value *EntityMulticodepointRow) { this.data.multicodepoint_table = value }

/**
 * EntityTableOpt
 */
type EntityTableOpt struct {
	ms_table *EntityStage1Row
	table    *EntityStage3Row
}

func MakeEntityTableOpt(ms_table *EntityStage1Row, table *EntityStage3Row) EntityTableOpt {
	return EntityTableOpt{
		ms_table: ms_table,
		table:    table,
	}
}
func (this *EntityTableOpt) GetMsTable() *EntityStage1Row      { return this.ms_table }
func (this *EntityTableOpt) SetMsTable(value *EntityStage1Row) { this.ms_table = value }
func (this *EntityTableOpt) GetTable() *EntityStage3Row        { return this.table }
func (this *EntityTableOpt) SetTable(value *EntityStage3Row)   { this.table = value }

/**
 * EntityCpMap
 */
type EntityCpMap struct {
	entity     *byte
	entity_len uint16
	codepoint1 uint
	codepoint2 uint
}

func MakeEntityCpMap(entity *byte, entity_len uint16, codepoint1 uint, codepoint2 uint) EntityCpMap {
	return EntityCpMap{
		entity:     entity,
		entity_len: entity_len,
		codepoint1: codepoint1,
		codepoint2: codepoint2,
	}
}
func (this *EntityCpMap) GetEntity() *byte { return this.entity }

// func (this *EntityCpMap) SetEntity(value *byte) { this.entity = value }
func (this *EntityCpMap) GetEntityLen() uint16 { return this.entity_len }

// func (this *EntityCpMap) SetEntityLen(value uint16) { this.entity_len = value }
func (this *EntityCpMap) GetCodepoint1() uint { return this.codepoint1 }

// func (this *EntityCpMap) SetCodepoint1(value uint) { this.codepoint1 = value }
func (this *EntityCpMap) GetCodepoint2() uint { return this.codepoint2 }

// func (this *EntityCpMap) SetCodepoint2(value uint) { this.codepoint2 = value }

/**
 * EntityHt
 */
type EntityHt struct {
	num_elems unsigned
	buckets   *EntityHtBucket
}

func MakeEntityHt(num_elems unsigned, buckets *EntityHtBucket) EntityHt {
	return EntityHt{
		num_elems: num_elems,
		buckets:   buckets,
	}
}
func (this *EntityHt) GetNumElems() unsigned { return this.num_elems }

// func (this *EntityHt) SetNumElems(value unsigned) { this.num_elems = value }
func (this *EntityHt) GetBuckets() *EntityHtBucket { return this.buckets }

// func (this *EntityHt) SetBuckets(value *EntityHtBucket) { this.buckets = value }
