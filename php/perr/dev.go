package perr

func Unreachable() error {
	return New("unreachable")
}
