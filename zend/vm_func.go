package zend

import "github.com/heyuuu/gophp/zend/types"

func (ex *ZendExecuteData) getOp(opline *ZendOp, typ uint8, node ZnodeOp) *types.Zval {
	switch typ {
	case IS_CONST:
		return RT_CONSTANT(opline, node)
	case IS_TMP_VAR:
		return EX_VAR(ex, node.GetVar())
	case IS_VAR:
		return EX_VAR(ex, node.GetVar()).DeRef()
	case IS_CV:
		zv := EX_VAR(ex, node.GetVar())
		if zv.IsUndef() {
			return ZvalUndefinedCv(node.GetVar(), ex)
		}
		return zv
	}
	return nil
}

func (ex *ZendExecuteData) GetOp1(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, opline.op1Type, opline.op1)
}

func (ex *ZendExecuteData) GetOp2(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, opline.op2Type, opline.op2)
}

func (ex *ZendExecuteData) GetVarOp1(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, IS_TMP_VAR, opline.op1)
}

func (ex *ZendExecuteData) GetVarOp2(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, IS_TMP_VAR, opline.op2)
}

func (ex *ZendExecuteData) GetCvOp1(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, IS_CV, opline.op1)
}

func (ex *ZendExecuteData) GetCvOp2(opline *ZendOp) *types.Zval {
	return ex.getOp(opline, IS_CV, opline.op2)
}
