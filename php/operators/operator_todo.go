package operators

import (
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
)

func (op *Operator) ZvalToArrayKey(offset Val) types.ArrayKey {
	if offset.IsString() {
		// todo 字符串转数字
		return types.StrKey(offset.String())
	} else if offset.IsLong() {
		return types.IdxKey(offset.Long())
	} else {
		// todo 其他类型处理
		perr.Panic("此类型 key 转 ArrayKey 未完成: " + types.ZvalGetType(offset))
		return types.IdxKey(0)
	}
}
