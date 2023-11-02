package operators

func ZvalGetLong(op Val) int      { return ZvalGetLongEx(op, true) }
func zvalGetLongNoisy(op Val) int { return ZvalGetLongEx(op, false) }
func ZvalGetLongEx(op Val, silent bool) int {
	// todo
}

func ZvalGetStrVal(op Val) string {
	// todo
}
