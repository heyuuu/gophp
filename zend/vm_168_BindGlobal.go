package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_BIND_GLOBAL_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var varname *types.String
	var value *types.Zval
	var variable_ptr *types.Zval
	var idx uintPtr
	var ref *types.ZendReference
	for {
		varname = opline.Const2().String()

		/* We store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

		idx = uintPtr(CACHED_PTR(opline.GetExtendedValue()) - 1)
		if idx < EG__().GetSymbolTable().GetNNumUsed()*b.SizeOf("Bucket") {
			var p *types.Bucket = (*types.Bucket)((*byte)(EG__().GetSymbolTable().Bucket(idx)))
			if p.GetVal().IsNotUndef() && (p.GetKey() == varname || p.GetH() == varname.GetH() && p.GetKey() != nil && types.ZendStringEqualContent(p.GetKey(), varname) != 0) {
				value = (*types.Zval)(p)
				goto check_indirect
			}
		}
		value = EG__().GetSymbolTable().KeyFind(varname.GetStr())
		if value == nil {
			value = EG__().GetSymbolTable().KeyAddNew(varname.GetStr(), EG__().GetUninitializedZval())

			SymbolFindAndCache(EG__().GetSymbolTable(), varname.GetStr(), executeData)
		} else {
			SymbolFindAndCache(EG__().GetSymbolTable(), varname.GetStr(), executeData)
		check_indirect:

			/* GLOBAL variable may be an INDIRECT pointer to CV */

			if value.IsIndirect() {
				value = value.Indirect()
				if value.IsUndef() {
					value.SetNull()
				}
			}

			/* GLOBAL variable may be an INDIRECT pointer to CV */

		}
		if !(value.IsReference()) {
			types.ZVAL_MAKE_REF_EX(value, 2)
			ref = value.Reference()
		} else {
			ref = value.Reference()
			// 			ref.AddRefcount()
		}
		variable_ptr = opline.Op1()
		variable_ptr.SetReference(ref)
		if b.PreInc(&opline).opcode != ZEND_BIND_GLOBAL {
			break
		}
	}
	OPLINE = opline
	return 0
	return ZEND_VM_NEXT_OPCODE(executeData, opline)
}
