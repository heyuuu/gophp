package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_ASSIGN_STATIC_PROP_OP_SPEC_HANDLER(executeData *ZendExecuteData) int {
	/* This helper actually never will receive IS_VAR as second op, and has the same handling for VAR and TMP in the first op, but for interoperability with the other binary_assign_op helpers, it is necessary to "include" it */

	var opline *ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var prop *types.Zval
	var value *types.Zval
	var prop_info *ZendPropertyInfo
	var ref *types.ZendReference
	if ZendFetchStaticPropertyAddress(&prop, &prop_info, (opline+1).GetExtendedValue(), BP_VAR_RW, 0, opline, executeData) != types.SUCCESS {
		b.Assert(EG__().GetException() != nil)
		UNDEF_RESULT()
		FREE_UNFETCHED_OP((opline + 1).GetOp1Type(), (opline + 1).GetOp1().GetVar())
		return 0
	}
	value = GetOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data)
	for {
		if prop.IsReference() {
			ref = prop.Reference()
			prop = types.Z_REFVAL_P(prop)
			if ZEND_REF_HAS_TYPE_SOURCES(ref) {
				ZendBinaryAssignOpTypedRef(ref, value, opline, executeData)
				break
			}
		}
		if prop_info.GetType() != 0 {

			/* special case for typed properties */

			ZendBinaryAssignOpTypedProp(prop_info, prop, value, opline, executeData)

			/* special case for typed properties */

		} else {
			ZendBinaryOp(prop, prop, value, opline)
		}
		break
	}
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), prop)
	}
	// 	FREE_OP(free_op_data)

	/* assign_static_prop has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
