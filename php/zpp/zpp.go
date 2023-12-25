package zpp

func NewParser(executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) IParser {
	return NewFastParser(executeData, minNumArgs, maxNumArgs, flags)
}
