package def

//go:generate stringer -type=Stage -output=enum_string.go

type Stage int

const (
	_ Stage = iota
	StageVariableName
	StageFuncName
	StageMethodName
	StageClassName
)
