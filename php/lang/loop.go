package lang

var BreakErr = breakError{}

type breakError struct{}

func (b breakError) Error() string { return "break" }
