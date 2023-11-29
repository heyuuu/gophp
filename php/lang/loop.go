package lang

const Break = false
const Continue = true

var BreakErr = breakError{}

type breakError struct{}

func (b breakError) Error() string { return "break" }
