package zend

type ArgInfos struct {
	requiredArgs int       // -1 表示需要所有参数
	returns      ArgInfo   // 返回值信息. notice: returnArg.name 无意义
	args         []ArgInfo // 参数信息列表
}

func MakeArgInfos(requiredArgs int, returns ArgInfo, args ...ArgInfo) ArgInfos {
	return ArgInfos{
		requiredArgs: requiredArgs,
		returns:      returns,
		args:         args,
	}
}
