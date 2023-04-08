package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var opline *ZendOp = executeData.GetOpline()
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		var rtd_key string = opline.Const1().GetStrVal()

		ce = EG__().ClassTable().Get(rtd_key)
		if ce == nil {
			for {
				b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & AccPreloaded) != 0)
				if ZendPreloadAutoload != nil && ZendPreloadAutoload(executeData.GetFunc().GetOpArray().GetFilename()) == types.SUCCESS {
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
			if ZendDoLinkClass(ce, b.CondF1(opline.GetOp2Type() == IS_CONST, func() *types.String { return opline.Const2().GetStr() }, nil)) == types.FAILURE {
				return 0
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	opline.Result().SetCe(ce)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
