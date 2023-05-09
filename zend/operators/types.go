package operators

type number struct {
	isF bool
	i   int
	f   float64
}

func intNumber(i int) number       { return number{i: i} }
func floatNumber(f float64) number { return number{isF: true, f: f} }
