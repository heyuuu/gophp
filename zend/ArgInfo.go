package zend

import "sik/zend/types"

type ArgInfoOpt func(*ArgInfo)

func ArgInfoType(typ types.ZendType) ArgInfoOpt {
	return func(info *ArgInfo) { info.typ = typ }
}
func ArgInfoByRef(byRef types.ZendUchar) ArgInfoOpt {
	return func(info *ArgInfo) { info.byReference = byRef }
}
func ArgInfoVariadic() ArgInfoOpt {
	return func(info *ArgInfo) { info.isVariadic = true }
}

/**
 * ArgInfo
 * 用于代替 ArgInfo / ZendArgInfo
 */
type ArgInfo struct {
	name        string
	typ         types.ZendType
	byReference uint8
	isVariadic  bool
	// 为 returnArg 临时使用，后续需替换
	requiredNumArgs int // -1 表示需要所有参数
}

func (this *ArgInfo) Name() string         { return this.name }
func (this *ArgInfo) Type() types.ZendType { return this.typ }
func (this *ArgInfo) ByReference() uint8   { return this.byReference }
func (this *ArgInfo) IsVariadic() bool     { return this.isVariadic }
func (this *ArgInfo) RequiredNumArgs() int { return this.requiredNumArgs }

func MakeArgName(name string) ArgInfo {
	return ArgInfo{name: name}
}

func MakeArgVariadic(name string) ArgInfo {
	return ArgInfo{name: name, isVariadic: true}
}

func MakeArgByRef(name string) ArgInfo {
	return ArgInfo{name: name, byReference: 1}
}

func MakeArgInfo(name string, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{name: name}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}
func MakeReturnArgInfo(requiredNumArgs int, opts ...ArgInfoOpt) ArgInfo {
	argInfo := ArgInfo{requiredNumArgs: requiredNumArgs}
	for _, opt := range opts {
		opt(&argInfo)
	}
	return argInfo
}
