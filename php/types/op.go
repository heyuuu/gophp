package types

import "github.com/heyuuu/gophp/zend"

type OpcodeHandlerT func(executeData *zend.ZendExecuteData) int

/**
 * ZnodeOp
 */
type ZnodeOp struct /* union */ {
	value uint32
	//constant  uint32
	//var_      uint32
	//num       uint32
	//oplineNum uint32
	//jmpOffset uint32
}

func (this *ZnodeOp) GetConstant() uint32       { return this.value }
func (this *ZnodeOp) SetConstant(value uint32)  { this.value = value }
func (this *ZnodeOp) GetVar() uint32            { return this.value }
func (this *ZnodeOp) SetVar(value uint32)       { this.value = value }
func (this *ZnodeOp) GetNum() uint32            { return this.value }
func (this *ZnodeOp) SetNum(value uint32)       { this.value = value }
func (this *ZnodeOp) GetOplineNum() uint32      { return this.value }
func (this *ZnodeOp) SetOplineNum(value uint32) { this.value = value }
func (this *ZnodeOp) GetJmpOffset() uint32      { return this.value }
func (this *ZnodeOp) SetJmpOffset(value uint32) { this.value = value }

/**
 * ZendOp
 */
type ZendOp struct {
	handler       OpcodeHandlerT // 指令执行 handler
	op1           ZnodeOp
	op2           ZnodeOp
	result        ZnodeOp
	extendedValue uint32
	lineno        uint32
	opcode        zend.OpCode
	op1Type       uint8
	op2Type       uint8
	resultType    uint8
}

func NewOp(lineno uint32) *ZendOp {
	return &ZendOp{
		op1:           ZnodeOp{value: 0},
		op2:           ZnodeOp{value: 0},
		result:        ZnodeOp{value: 0},
		extendedValue: 0,
		lineno:        lineno,
		opcode:        zend.ZEND_NOP,
		op1Type:       zend.IS_UNUSED,
		op2Type:       zend.IS_UNUSED,
		resultType:    zend.IS_UNUSED,
	}
}
func (op *ZendOp) SetNop() {
	// MAKE_NOP
	op.op1.value = 0
	op.op2.value = 0
	op.result.value = 0
	op.opcode = zend.ZEND_NOP
	op.op1Type = zend.IS_UNUSED
	op.op2Type = zend.IS_UNUSED
	op.resultType = zend.IS_UNUSED
}

func (op *ZendOp) GetHandler() OpcodeHandlerT      { return op.handler }
func (op *ZendOp) SetHandler(value OpcodeHandlerT) { op.handler = value }
func (op *ZendOp) GetOp1() *ZnodeOp                { return &op.op1 }
func (op *ZendOp) SetOp1(value ZnodeOp)            { op.op1 = value }
func (op *ZendOp) GetOp2() *ZnodeOp                { return &op.op2 }
func (op *ZendOp) SetOp2(value ZnodeOp)            { op.op2 = value }
func (op *ZendOp) GetResult() *ZnodeOp             { return &op.result }
func (op *ZendOp) SetResult(value ZnodeOp)         { op.result = value }
func (op *ZendOp) GetExtendedValue() uint32        { return op.extendedValue }
func (op *ZendOp) SetExtendedValue(value uint32)   { op.extendedValue = value }
func (op *ZendOp) GetLineno() uint32               { return op.lineno }
func (op *ZendOp) SetLineno(value uint32)          { op.lineno = value }
func (op *ZendOp) GetOpcode() zend.OpCode          { return op.opcode }
func (op *ZendOp) SetOpcode(value zend.OpCode)     { op.opcode = value }
func (op *ZendOp) GetOp1Type() uint8               { return op.op1Type }
func (op *ZendOp) SetOp1Type(value uint8)          { op.op1Type = value }
func (op *ZendOp) GetOp2Type() uint8               { return op.op2Type }
func (op *ZendOp) SetOp2Type(value uint8)          { op.op2Type = value }
func (op *ZendOp) GetResultType() uint8            { return op.resultType }
func (op *ZendOp) SetResultType(value uint8)       { op.resultType = value }

func (op *ZendOp) Offset(offset int) *ZendOp { return op + offset }

func (op *ZendOp) currEx() *zend.ZendExecuteData {
	return zend.CurrEX()
}

//
func (op *ZendOp) _const(node ZnodeOp) *Zval { return zend.RT_CONSTANT(op, node) }
func (op *ZendOp) _op(node ZnodeOp) *Zval    { return zend.EX_VAR(node.GetVar()) }
func (op *ZendOp) _cvOrUndef(node ZnodeOp) *Zval {
	ret := op._op(node)
	if ret.IsUndef() {
		return zend.ZvalUndefinedCv(node.GetVar(), op.currEx())
	}
	return ret
}

func (op *ZendOp) Const1() *Zval     { return op._const(op.op1) }
func (op *ZendOp) Const2() *Zval     { return op._const(op.op2) }
func (op *ZendOp) Op1() *Zval        { return op._op(op.op1) }
func (op *ZendOp) Op2() *Zval        { return op._op(op.op2) }
func (op *ZendOp) Result() *Zval     { return op._op(op.result) }
func (op *ZendOp) Cv1OrUndef() *Zval { return op._cvOrUndef(op.op1) }
func (op *ZendOp) Cv2OrUndef() *Zval { return op._cvOrUndef(op.op2) }
