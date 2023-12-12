package perr

func Unreachable() error {
	return NewInternal("unreachable")
}

func Assert(cond bool) {
	if !cond {
		panic(NewInternal("Internal Assert Fail"))
	}
}
func AssertEx(cond bool, message string) {
	if !cond {
		panic(NewInternal("Internal Assert Fail:" + message))
	}
}
