package perr

func Unreachable() error {
	return NewInternal("unreachable")
}
