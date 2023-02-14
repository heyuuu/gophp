package zend

import b "sik/builtin"

type ArgInfoOpt func(*ArgInfo)

func ArgInfoType(typ ZendType) ArgInfoOpt {
	return func(info *ArgInfo) { info.type_ = typ }
}
func ArgInfoByRef(byRef ZendUchar) ArgInfoOpt {
	return func(info *ArgInfo) { info.pass_by_reference = byRef }
}
func ArgInfoVariadic() ArgInfoOpt {
	return func(info *ArgInfo) { info.is_variadic = 1 }
}

/**
 * ArgInfo
 * 用于代替 ArgInfo / ZendArgInfo
 */
type ArgInfo struct {
	// 为 returnArg 临时使用，后续需替换
	requiredArgs int // -1 表示需要所有参数
	//
	name              *byte
	type_             ZendType
	pass_by_reference ZendUchar
	is_variadic       ZendBool
}

func (this *ArgInfo) RequiredArgs() int { return this.requiredArgs }
func (this *ArgInfo) Name() string      { return b.CastStrAuto(this.name) }

func MakeArgInfo(name string, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{name: name}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}
func MakeArgInfoSpecial(name int, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{name: name}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}

func MakeZendInternalArgInfo(name *byte, type_ ZendType, pass_by_reference ZendUchar, is_variadic ZendBool) ArgInfo {
	return ArgInfo{
		name:              name,
		type_:             type_,
		pass_by_reference: pass_by_reference,
		is_variadic:       is_variadic,
	}
}
