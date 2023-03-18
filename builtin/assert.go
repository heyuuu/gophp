package builtin

var OpenAssert = true

func Assert(c bool) {
	if OpenAssert && !c {
		panic("Assert Fail")
	}
}
