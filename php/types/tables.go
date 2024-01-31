package types

// ConstantTable
type ConstantTable = Table[*Constant]

func NewConstantTable() *ConstantTable {
	return NewTable[*Constant]()
}

// FunctionTable
type FunctionTable = Table[*Function]

func NewFunctionTable() *FunctionTable {
	return NewTable[*Function]()
}

// ClassTable
type ClassTable = Table[*Class]

func NewClassTable() *ClassTable {
	return NewTable[*Class]()
}

// ClassConstantTable
type ClassConstantTable = Table[*ClassConstant]

func NewClassConstantTable() *ClassConstantTable {
	return NewTable[*ClassConstant]()
}

// PropertyInfoTable
type PropertyInfoTable = Table[*PropertyInfo]

func NewPropertyInfoTable() *PropertyInfoTable {
	return NewTable[*PropertyInfo]()
}
