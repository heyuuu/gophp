package types

type ArgInfo struct {
}

type Function interface {
	Name() string
	ArgInfos() []ArgInfo
}
