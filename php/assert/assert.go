package assert

import "github.com/heyuuu/gophp/php/perr"

func Assert(cond bool) {
	if !cond {
		panic(perr.Internal("Internal Assert Fail"))
	}
}
func AssertEx(cond bool, message string) {
	if !cond {
		panic(perr.Internal("Internal Assert Fail:" + message))
	}
}
