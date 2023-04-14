package zend

import (
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var lcname *types.Zval
	var ce *types.ClassEntry
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		lcname = opline.Const1()
		var lcname1 *types.Zval = lcname + 1

		lcnameStr := lcname.StringVal()
		lcname1Str := lcname1.StringVal()

		if EG__().ClassTable().Exists(lcname1Str) {
			ce = EG__().ClassTable().Get(lcname1Str)

			// 判断新名称是否已存在
			if EG__().ClassTable().Exists(lcnameStr) {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.Name())
			}

			//
			if ZendDoLinkClass(ce, opline.Const2().String()) == types.FAILURE {
				return 0
			}

			// 更新 key
			EG__().ClassTable().Del(lcname1Str)
			EG__().ClassTable().Update(lcnameStr, ce)
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
