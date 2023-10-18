package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

func ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var ce *types.ClassEntry
	var opline *types.ZendOp = executeData.GetOpline()
	ce = CACHED_PTR(opline.GetExtendedValue())
	if ce == nil {
		var rtd_key string = opline.Const1().String()

		ce = EG__().ClassTable().Get(rtd_key)
		if ce == nil {
			b.Assert((executeData.GetFunc().GetOpArray().GetFnFlags() & types.AccPreloaded) != 0)
			faults.ErrorNoreturn(faults.E_ERROR, "Anonymous class wasn't preloaded")
		}

		b.Assert(ce != nil)
		if !ce.IsLinked() {
			var lc_parent_name string
			if opline.GetOp2Type() == IS_CONST {
				lc_parent_name = opline.Const2().String()
			}
			if ZendDoLinkClass(ce, lc_parent_name) == types.FAILURE {
				return 0
			}
		}
		CACHE_PTR(opline.GetExtendedValue(), ce)
	}
	opline.Result().SetCe(ce)
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
