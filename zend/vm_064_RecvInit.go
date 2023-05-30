package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_RECV_INIT_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var arg_num uint32
	var param *types.Zval
	for {
		arg_num = opline.GetOp1().GetNum()
		param = opline.Result()
		if arg_num > executeData.NumArgs() {
			var default_value *types.Zval = opline.Const2()
			if default_value.IsConstantAst() {
				var cache_val *types.Zval = (*types.Zval)(CACHE_ADDR(default_value.GetCacheSlot()))

				/* we keep in cache only not refcounted values */

				if cache_val.IsNotUndef() {
					param.CopyValueFrom(cache_val)
				} else {
					types.ZVAL_COPY(param, default_value)
					if ZvalUpdateConstantEx(param, executeData.GetFunc().GetOpArray().scope) != types.SUCCESS {
						// ZvalPtrDtorNogc(param)
						param.SetUndef()
						return 0
					}
					if !(param.IsRefcounted()) {
						cache_val.CopyValueFrom(param)
					}
				}
				goto recv_init_check_type
			} else {
				types.ZVAL_COPY(param, default_value)
			}
		} else {
		recv_init_check_type:
			if (executeData.GetFunc().GetOpArray().GetFnFlags() & types.AccHasTypeHints) != 0 {
				var default_value *types.Zval = opline.Const2()
				if ZendVerifyRecvArgType(executeData.GetFunc(), arg_num, param, default_value, CACHE_ADDR(opline.GetExtendedValue())) == 0 {
					return 0
				}
			}
		}
		if b.PreInc(&opline).opcode != ZEND_RECV_INIT {
			break
		}
	}
	OPLINE = opline
	return 0
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
