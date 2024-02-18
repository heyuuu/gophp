package standard

type EncToUniStage2 = []uint16
type EncToUni = []EncToUniStage2

type UniToEnc struct {
	unCodePoint uint16 `get:""`
	csCode      uint8  `get:""`
}

type EntityMulticodepointRow struct /* union */ {
	entity string
	value  uint
	//leadingEntry struct {
	//	defaultEntity    *byte
	//	size             uint /* number of remaining entries in the table */
	//	defaultEntityLen uint16
	//}
	//normalEntry struct {
	//	entity    *byte
	//	secondCp  uint /* second code point */
	//	entityLen uint16
	//}
}

func (t *EntityMulticodepointRow) Entity() string        { return t.entity }
func (t *EntityMulticodepointRow) DefaultEntity() string { return t.entity }
func (t *EntityMulticodepointRow) Size() uint            { return t.value }
func (t *EntityMulticodepointRow) SecondCp() uint        { return t.value }

// EntityStage3Row
type EntityStage3Row struct {
	ambiguous bool `get:""`
	/* union */
	entity              string                    `get:""`
	multiCodepointTable []EntityMulticodepointRow `get:""`
}

func NewEntityStage3Row(entity string) *EntityStage3Row {
	return &EntityStage3Row{
		ambiguous: false,
		entity:    entity,
	}
}

func NewEntityStage3RowAmbiguous(multiCodepointTable []EntityMulticodepointRow) *EntityStage3Row {
	return &EntityStage3Row{
		ambiguous:           true,
		multiCodepointTable: multiCodepointTable,
	}
}

// EntityTableOpt
type EntityTableOpt struct {
	msTable []EntityStage1Row  `get:""`
	table   []*EntityStage3Row `get:""`
}

func MakeEntityTableOpt(msTable []EntityStage1Row, table []*EntityStage3Row) EntityTableOpt {
	return EntityTableOpt{msTable: msTable, table: table}
}

type EntityMap = map[string][2]uint
