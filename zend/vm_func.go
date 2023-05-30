package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

type opMode int

const (
	opModeVarDeRef opMode = 1 << iota
	opModeCvCheckUndef
	opModeCvDeRef

	opMode0 = opModeVarDeRef | opModeCvCheckUndef | opModeCvDeRef
	opMode1 = 0
	opMode2 = opModeCvCheckUndef
)

func (ex *ZendExecuteData) getOp(opline *types.ZendOp, typ uint8, node types.ZnodeOp, mode opMode) *types.Zval {
	switch typ {
	case IS_CONST:
		return RT_CONSTANT(opline, node)
	case IS_TMP_VAR:
		return EX_VAR(ex, node.GetVar())
	case IS_VAR:
		zv := EX_VAR(ex, node.GetVar())
		if mode&opModeVarDeRef != 0 {
			zv = zv.DeRef()
		}
		return zv
	case IS_CV:
		zv := EX_VAR(ex, node.GetVar())
		if mode&opModeCvCheckUndef != 0 {
			if zv.IsUndef() {
				return ZvalUndefinedCv(node.GetVar(), ex)
			}
		}
		if mode&opModeCvDeRef != 0 {
			zv = zv.DeRef()
		}
		return zv
	}
	return nil
}

func (ex *ZendExecuteData) Op1(opline *types.ZendOp, mode opMode) *types.Zval {
	return ex.getOp(opline, opline.op1Type, opline.op1, mode)
}
func (ex *ZendExecuteData) Op2(opline *types.ZendOp, mode opMode) *types.Zval {
	return ex.getOp(opline, opline.op2Type, opline.op2, mode)
}

func (ex *ZendExecuteData) GetOp1(opline *types.ZendOp) *types.Zval { return ex.Op1(opline, opMode0) }
func (ex *ZendExecuteData) GetOp2(opline *types.ZendOp) *types.Zval { return ex.Op2(opline, opMode0) }

func (ex *ZendExecuteData) GetVarOp1(opline *types.ZendOp) *types.Zval {
	return ex.getOp(opline, IS_TMP_VAR, opline.op1, opMode0)
}

func (ex *ZendExecuteData) GetVarOp2(opline *types.ZendOp) *types.Zval {
	return ex.getOp(opline, IS_TMP_VAR, opline.op2, opMode0)
}

func (ex *ZendExecuteData) GetCvOp1(opline *types.ZendOp) *types.Zval {
	return ex.getOp(opline, IS_CV, opline.op1, opMode0)
}

func (ex *ZendExecuteData) GetCvOp2(opline *types.ZendOp) *types.Zval {
	return ex.getOp(opline, IS_CV, opline.op2, opMode0)
}
