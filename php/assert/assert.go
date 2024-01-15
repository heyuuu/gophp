package assert

import "github.com/heyuuu/gophp/php/perr"

func Assert(cond bool) {
	if !cond {
		perr.Panic("Assert Fail")
	}
}
func AssertEx(cond bool, message string) {
	if !cond {
		perr.Panic(message)
	}
}
