package types

import "math"

const MaxLong = math.MaxInt
const MinLong = math.MinInt

/* Used to get hash of the properties of the object, as hash of zval's */
type PropPurposeType int

const (
	PropPurposeDebug PropPurposeType = iota
	PropPurposeArrayCast
	PropPurposeSerialize
	PropPurposeVarExport
	PropPurposeJson
	PropPurposeArrayKeyExists
)
