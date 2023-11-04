package lang

func Assert(cond bool) {
	if !cond {
		panic("Internal Assert Fail")
	}
}
func AssertEx(cond bool, message string) {
	if !cond {
		panic("Internal Assert Fail:" + message)
	}
}
