package standard

// properties for EntityStage3Row
func (t *EntityStage3Row) Ambiguous() bool {
	return t.ambiguous
}
func (t *EntityStage3Row) Entity() string {
	return t.entity
}
func (t *EntityStage3Row) MultiCodepointTable() []EntityMulticodepointRow {
	return t.multiCodepointTable
}

// properties for EntityTableOpt
func (t *EntityTableOpt) MsTable() []EntityStage1Row {
	return t.msTable
}
func (t *EntityTableOpt) Table() []*EntityStage3Row {
	return t.table
}

// properties for UniToEnc
func (t *UniToEnc) UnCodePoint() uint16 {
	return t.unCodePoint
}
func (t *UniToEnc) CsCode() uint8 {
	return t.csCode
}
