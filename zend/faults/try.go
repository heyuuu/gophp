package faults

import "errors"

var bailoutError = errors.New("faults bailout")

func TryCatch(try func(), catch func()) (ok bool) {
	defer func() {
		if err := recover(); err != nil {
			if catch != nil {
				catch()
			}
			ok = false
		}
	}()

	if try != nil {
		try()
	}

	return true
}

func Try(try func()) bool {
	return TryCatch(try, nil)
}

func throw() {
	panic(bailoutError)
}
