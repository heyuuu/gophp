package zend

import (
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_PRE_INC_STATIC_PROP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var prop *types.Zval
	var prop_info *types.PropertyInfo
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, opline.GetExtendedValue(), BP_VAR_RW, 0, opline, executeData) != types.SUCCESS {
		UNDEF_RESULT()
		return 0
	}
	ZendPreIncdecPropertyZval(prop, b.Cond(prop_info.GetType() != 0, prop_info, nil), opline, executeData)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
