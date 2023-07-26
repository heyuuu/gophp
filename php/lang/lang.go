package lang

// 一元操作

type num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

func PreInc[T num](variable *T) T {
	*variable++
	return *variable
}

func PreDec[T num](variable *T) T {
	*variable--
	return *variable
}

func PostInc[T num](variable *T) T {
	var result = *variable
	*variable++
	return result
}

func PostDec[T num](variable *T) T {
	var result = *variable
	*variable--
	return result
}

// 赋值操作
func Assign[T any](variable *T, value T) T {
	*variable = value
	return *variable
}

// 三元操作 cond ? a : b

func Cond[T any](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	}
	return falseVal
}

func CondF[T any](cond bool, trueValue func() T, falseValue func() T) T {
	if cond {
		return trueValue()
	}
	return falseValue()
}

func CondF1[T any](cond bool, trueValue func() T, falseValue T) T {
	if cond {
		return trueValue()
	}
	return falseValue
}

func CondF2[T any](cond bool, trueValue T, falseValue func() T) T {
	if cond {
		return trueValue
	}
	return falseValue()
}
