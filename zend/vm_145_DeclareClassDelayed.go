package zend

import (
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var lcname *types.Zval
	var zv *types.Zval
	var ce *types.ClassEntry
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		lcname = opline.Const1()
		zv = EG__().GetClassTable().KeyFind((lcname + 1).GetStr().GetStr())
		if zv != nil {
			ce = zv.GetCe()
			zv = types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), lcname.GetStr().GetStr())
			if zv == nil {
				faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot declare %s %s, because the name is already in use", ZendGetObjectType(ce), ce.GetName().GetVal())
			} else {
				if ZendDoLinkClass(ce, opline.Const2().GetStr()) == types.FAILURE {

					/* Reload bucket pointer, the hash table may have been reallocated */

					zv = EG__().GetClassTable().KeyFind(lcname.GetStr().GetStr())
					types.ZendHashSetBucketKey(EG__().GetClassTable(), (*types.Bucket)(zv), (lcname + 1).GetStr().GetStr())
					return 0
				}
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
