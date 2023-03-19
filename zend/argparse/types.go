package argparse

import "log"

type strReader struct {
	str string
}

func (r *strReader) curr() byte {
	if r.str != "" {
		return r.str[0]
	}
	return 0
}

func (r *strReader) inc() {
	if r.str != "" {
		r.str = r.str[1:]
	}
}

func vaArg[T any](va *[]any) *T {
	if len(*va) == 0 {
		log.Fatal("解析参数异常，超过获取长度")
	}

	ptr, ok := (*va)[0].(*T)
	if !ok {
		log.Fatalf("解析参数异常: 类型不匹配")
	}

	*va = (*va)[1:]
	return ptr
}

func vaArg_[T any](va *[]any, ptr *T) {
	if len(*va) == 0 {
		log.Fatal("解析参数异常，超过获取长度")
	}

	if val, ok := (*va)[0].(T); ok {
		*ptr = val
		*va = (*va)[1:]
	} else {
		log.Fatalf("解析参数异常: 类型不匹配，pos=%d", r.pos)
	}
}
