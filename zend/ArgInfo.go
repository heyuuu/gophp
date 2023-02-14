package zend

type ArgInfoOpt func(*ArgInfo)

func ArgInfoType(typ ZendType) ArgInfoOpt {
	return func(info *ArgInfo) { info.typ = typ }
}
func ArgInfoByRef(byRef ZendUchar) ArgInfoOpt {
	return func(info *ArgInfo) { info.passByReference = byRef }
}
func ArgInfoVariadic() ArgInfoOpt {
	return func(info *ArgInfo) { info.isVariadic = true }
}

/**
 * ArgInfo
 * 用于代替 ArgInfo / ZendArgInfo
 */
type ArgInfo struct {
	name            string
	typ             ZendType
	passByReference ZendUchar
	isVariadic      bool
	// 为 returnArg 临时使用，后续需替换
	requiredArgs int // -1 表示需要所有参数
}

func (this *ArgInfo) Name() string      { return this.name }
func (this *ArgInfo) Type() ZendType    { return this.typ }
func (this *ArgInfo) RequiredArgs() int { return this.requiredArgs }

func MakeArgInfo(name string, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{name: name}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}
func MakeReturnArgInfo(requiredArgs int, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{requiredArgs: requiredArgs}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}
