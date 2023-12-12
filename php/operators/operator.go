package operators

// Operator
type Operator struct {
	OperatorHandler
}

var _ iOperator = (*Operator)(nil)

func New(handler OperatorHandler) *Operator {
	return &Operator{OperatorHandler: handler}
}
