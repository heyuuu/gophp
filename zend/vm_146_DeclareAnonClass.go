package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var ce *types2.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		var rtd_key string = opline.Const1().StringVal()

		ce = EG__().ClassTable().Get(rtd_key)
		if ce == nil {
			for {
				b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & AccPreloaded) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(executeData.GetFunc().GetOpArray().GetFilename()) == types2.SUCCESS {
					ce = EG__().ClassTable().Get(rtd_key)
					if ce != nil {
						goto afterGetCe
					}
				}
				faults.ErrorNoreturn(faults.E_ERROR, "Anonymous class wasn't preloaded")
				break
			}
		}

	afterGetCe:
		b.Assert(ce != nil)
		if !ce.IsLinked() {
			if ZendDoLinkClass(ce, b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types2.String { return opline.Const2().String() }, nil)) == types2.FAILURE {
				return 0
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	opline.Result().SetCe(ce)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
