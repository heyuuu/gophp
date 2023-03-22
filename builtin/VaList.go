package builtin

import "log"

func VaArg(va *[]any) any {
	if len(*va) == 0 {
		log.Fatal("解析参数异常，超过获取长度")
	}

	arg := (*va)[0]
	*va = (*va)[1:]
	return arg
}
