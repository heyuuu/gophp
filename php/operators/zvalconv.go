package operators

func ZvalIsTrue(op Val) bool {

}

func ZvalGetLong(op Val) int      { return ZvalGetLongEx(op, true) }
func zvalGetLongNoisy(op Val) int { return ZvalGetLongEx(op, false) }
func ZvalGetLongEx(op Val, silent bool) int {
	// todo
	return 0
}

func ZvalGetStrVal(op Val) string {
	// todo
	return ""
}
