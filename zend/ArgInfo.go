package zend

import b "sik/builtin"

/**
 * ArgInfo
 * 用于代替 ArgInfo / ZendArgInfo
 */
type ArgInfo struct {
	name              *byte
	type_             ZendType
	pass_by_reference ZendUchar
	is_variadic       ZendBool
}

func (this *ArgInfo) Name() string { return b.CastStrAuto(this.name) }

func MakeZendInternalArgInfo(name *byte, type_ ZendType, pass_by_reference ZendUchar, is_variadic ZendBool) ArgInfo {
	return ArgInfo{
		name:              name,
		type_:             type_,
		pass_by_reference: pass_by_reference,
		is_variadic:       is_variadic,
	}
}
