package zend

import (
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * ZnodeOp
 */
type ZnodeOp struct /* union */ {
	constant  uint32
	var_      uint32
	num       uint32
	oplineNum uint32
	jmpOffset uint32
}

func (this *ZnodeOp) GetConstant() uint32       { return this.constant }
func (this *ZnodeOp) SetConstant(value uint32)  { this.constant = value }
func (this *ZnodeOp) GetVar() uint32            { return this.var_ }
func (this *ZnodeOp) SetVar(value uint32)       { this.var_ = value }
func (this *ZnodeOp) GetNum() uint32            { return this.num }
func (this *ZnodeOp) SetNum(value uint32)       { this.num = value }
func (this *ZnodeOp) GetOplineNum() uint32      { return this.oplineNum }
func (this *ZnodeOp) SetOplineNum(value uint32) { this.oplineNum = value }
func (this *ZnodeOp) GetJmpOffset() uint32      { return this.jmpOffset }
func (this *ZnodeOp) SetJmpOffset(value uint32) { this.jmpOffset = value }

/**
 * Znode
 */
type Znode struct {
	op_type uint8
	flag    uint8
	u       struct /* union */ {
		op       ZnodeOp
		constant types.Zval
	}
}

func (this *Znode) GetOpType() uint8        { return this.op_type }
func (this *Znode) SetOpType(value uint8)   { this.op_type = value }
func (this *Znode) GetOp() ZnodeOp          { return this.u.op }
func (this *Znode) SetOp(value ZnodeOp)     { this.u.op = value }
func (this *Znode) GetConstant() types.Zval { return this.u.constant }

/**
 * ZendAstZnode
 */
type ZendAstZnode struct {
	kind   ZendAstKind
	attr   ZendAstAttr
	lineno uint32
	node   Znode
}

func (this *ZendAstZnode) SetKind(value ZendAstKind) { this.kind = value }
func (this *ZendAstZnode) SetAttr(value ZendAstAttr) { this.attr = value }
func (this *ZendAstZnode) SetLineno(value uint32)    { this.lineno = value }
func (this *ZendAstZnode) GetNode() Znode            { return this.node }
func (this *ZendAstZnode) SetNode(value Znode)       { this.node = value }

/**
 * ZendDeclarables
 */
type ZendDeclarables struct {
	ticks ZendLong
}

func (this *ZendDeclarables) GetTicks() ZendLong      { return this.ticks }
func (this *ZendDeclarables) SetTicks(value ZendLong) { this.ticks = value }

/**
 * ZendFileContext
 */
type ZendFileContext struct {
	declarables              ZendDeclarables
	current_namespace        *types.String
	in_namespace             types.ZendBool
	has_bracketed_namespaces types.ZendBool
	imports                  *types.Array
	imports_function         *types.Array
	imports_const            *types.Array
	seen_symbols             types.Array
}

func (this *ZendFileContext) GetDeclarables() ZendDeclarables      { return this.declarables }
func (this *ZendFileContext) SetDeclarables(value ZendDeclarables) { this.declarables = value }
func (this *ZendFileContext) GetCurrentNamespace() *types.String   { return this.current_namespace }
func (this *ZendFileContext) SetCurrentNamespace(value *types.String) {
	this.current_namespace = value
}
func (this *ZendFileContext) GetInNamespace() types.ZendBool      { return this.in_namespace }
func (this *ZendFileContext) SetInNamespace(value types.ZendBool) { this.in_namespace = value }
func (this *ZendFileContext) GetHasBracketedNamespaces() types.ZendBool {
	return this.has_bracketed_namespaces
}
func (this *ZendFileContext) SetHasBracketedNamespaces(value types.ZendBool) {
	this.has_bracketed_namespaces = value
}
func (this *ZendFileContext) GetImports() *types.Array         { return this.imports }
func (this *ZendFileContext) SetImports(value *types.Array)    { this.imports = value }
func (this *ZendFileContext) GetImportsFunction() *types.Array { return this.imports_function }
func (this *ZendFileContext) SetImportsFunction(value *types.Array) {
	this.imports_function = value
}
func (this *ZendFileContext) GetImportsConst() *types.Array      { return this.imports_const }
func (this *ZendFileContext) SetImportsConst(value *types.Array) { this.imports_const = value }
func (this *ZendFileContext) GetSeenSymbols() types.Array        { return this.seen_symbols }

// func (this *ZendFileContext) SetSeenSymbols(value HashTable) { this.seen_symbols = value }

/**
 * ZendParserStackElem
 */
type ZendParserStackElem struct /* union */ {
	ast *ZendAst
	str *types.String
	num ZendUlong
	ptr *uint8
}

func (this *ZendParserStackElem) GetAst() *ZendAst      { return this.ast }
func (this *ZendParserStackElem) SetAst(value *ZendAst) { this.ast = value }
func (this *ZendParserStackElem) GetStr() *types.String { return this.str }

/**
 * ZendOp
 */
type ZendOp struct {
	handler       any // 指令执行 handler
	op1           ZnodeOp
	op2           ZnodeOp
	result        ZnodeOp
	extendedValue uint32
	lineno        uint32
	opcode        OpCode
	op1Type       uint8
	op2Type       uint8
	resultType    uint8
}

func (op *ZendOp) GetHandler() any               { return op.handler }
func (op *ZendOp) SetHandler(value any)          { op.handler = value }
func (op *ZendOp) GetOp1() ZnodeOp               { return op.op1 }
func (op *ZendOp) SetOp1(value ZnodeOp)          { op.op1 = value }
func (op *ZendOp) GetOp2() ZnodeOp               { return op.op2 }
func (op *ZendOp) SetOp2(value ZnodeOp)          { op.op2 = value }
func (op *ZendOp) GetResult() ZnodeOp            { return op.result }
func (op *ZendOp) SetResult(value ZnodeOp)       { op.result = value }
func (op *ZendOp) GetExtendedValue() uint32      { return op.extendedValue }
func (op *ZendOp) SetExtendedValue(value uint32) { op.extendedValue = value }
func (op *ZendOp) GetLineno() uint32             { return op.lineno }
func (op *ZendOp) SetLineno(value uint32)        { op.lineno = value }
func (op *ZendOp) GetOpcode() OpCode             { return op.opcode }
func (op *ZendOp) SetOpcode(value OpCode)        { op.opcode = value }
func (op *ZendOp) GetOp1Type() uint8             { return op.op1Type }
func (op *ZendOp) SetOp1Type(value uint8)        { op.op1Type = value }
func (op *ZendOp) GetOp2Type() uint8             { return op.op2Type }
func (op *ZendOp) SetOp2Type(value uint8)        { op.op2Type = value }
func (op *ZendOp) GetResultType() uint8          { return op.resultType }
func (op *ZendOp) SetResultType(value uint8)     { op.resultType = value }

func (op *ZendOp) Offset(offset int) *ZendOp { return op + offset }

func (op *ZendOp) currEx() *ZendExecuteData {
	return CurrEX()
}

/**
 * opGetter
 */
type opGetter func(node ZnodeOp) *types.Zval
type opExGetter func(node ZnodeOp, shouldFree *ZendFreeOp) *types.Zval

func (op *ZendOp) _complexOp(opType uint8, node ZnodeOp, constGetter opGetter, varGetter opGetter, cvGetter opGetter) *types.Zval {
	switch opType {
	case IS_CONST:
		return constGetter(node)
	case IS_TMP_VAR, IS_VAR:
		return varGetter(node)
	case IS_CV:
		return cvGetter(node)
	}
	panic("unreachable")
}
func (op *ZendOp) _complexOpEx(opType uint8, node ZnodeOp, shouldFree *ZendFreeOp, constGetter opGetter, varGetter opExGetter, cvGetter opGetter) *types.Zval {
	switch opType {
	case IS_CONST:
		return constGetter(node)
	case IS_TMP_VAR, IS_VAR:
		return varGetter(node, shouldFree)
	case IS_CV:
		return cvGetter(node)
	}
	panic("unreachable")
}

//
func (op *ZendOp) _const(node ZnodeOp) *types.Zval { return RT_CONSTANT(op, node) }
func (op *ZendOp) _var(node ZnodeOp) *types.Zval   { return EX_VAR(node.GetVar()) }
func (op *ZendOp) _varAndPtr(node ZnodeOp, shouldFree *ZendFreeOp) *types.Zval {
	var ret = op._var(node)
	*shouldFree = ret
	return ret
}
func (op *ZendOp) _cvOrUndef(node ZnodeOp) *types.Zval {
	ret := op._var(node)
	if ret.IsUndef() {
		return ZvalUndefinedCv(node.var_, op.currEx())
	}
	return ret
}

func (op *ZendOp) Const1() *types.Zval { return op._const(op.op1) }
func (op *ZendOp) Const2() *types.Zval { return op._const(op.op2) }
func (op *ZendOp) Op1() *types.Zval    { return op._var(op.op1) }
func (op *ZendOp) Op2() *types.Zval    { return op._var(op.op2) }
func (op *ZendOp) Result() *types.Zval { return op._var(op.result) }

//
func (op *ZendOp) Op1Ptr(shouldFree *ZendFreeOp) *types.Zval {
	return op._varAndPtr(op.op1, shouldFree)
}
func (op *ZendOp) Op2Ptr(shouldFree *ZendFreeOp) *types.Zval {
	return op._varAndPtr(op.op2, shouldFree)
}

func (op *ZendOp) Cv1OrUndef() *types.Zval { return op._cvOrUndef(op.op1) }
func (op *ZendOp) Cv2OrUndef() *types.Zval { return op._cvOrUndef(op.op2) }

// VarEx
func (op *ZendOp) _varEx(opType uint8, node ZnodeOp) *types.Zval {
	return op._complexOp(opType, node, op._const, op._var, op._var)
}
func (op *ZendOp) Op1Ex() *types.Zval { return op._varEx(op.op1Type, op.op1) }
func (op *ZendOp) Op2Ex() *types.Zval { return op._varEx(op.op2Type, op.op2) }

// VarExEx
func (op *ZendOp) _varExEx(opType uint8, node ZnodeOp, shouldFree *ZendFreeOp) *types.Zval {
	return op._complexOpEx(opType, node, shouldFree, op._const, op._varAndPtr, op._cvOrUndef)
}
func (op *ZendOp) Op1ExEx(shouldFree *ZendFreeOp) *types.Zval {
	return op._varExEx(op.op1Type, op.op1, shouldFree)
}
func (op *ZendOp) Op2ExEx(shouldFree *ZendFreeOp) *types.Zval {
	return op._varExEx(op.op2Type, op.op2, shouldFree)
}

//
func (op *ZendOp) _concatOp(freeOp *ZendFreeOp, opType uint8, node ZnodeOp) *types.Zval {
	return op._complexOpEx(opType, node, freeOp, op._const, op._varAndPtr, op._var)
}
func (op *ZendOp) ConcatOp1(freeOp *ZendFreeOp) *types.Zval {
	return op._concatOp(freeOp, op.op1Type, op.op1)
}
func (op *ZendOp) ConcatOp2(freeOp *ZendFreeOp) *types.Zval {
	return op._concatOp(freeOp, op.op1Type, op.op1)
}

/**
 * ZendBrkContElement
 */
type ZendBrkContElement struct {
	start     int
	cont      int
	brk       int
	parent    int
	is_switch types.ZendBool
}

func (this *ZendBrkContElement) GetStart() int                    { return this.start }
func (this *ZendBrkContElement) SetStart(value int)               { this.start = value }
func (this *ZendBrkContElement) GetCont() int                     { return this.cont }
func (this *ZendBrkContElement) SetCont(value int)                { this.cont = value }
func (this *ZendBrkContElement) GetBrk() int                      { return this.brk }
func (this *ZendBrkContElement) SetBrk(value int)                 { this.brk = value }
func (this *ZendBrkContElement) GetParent() int                   { return this.parent }
func (this *ZendBrkContElement) SetParent(value int)              { this.parent = value }
func (this *ZendBrkContElement) GetIsSwitch() types.ZendBool      { return this.is_switch }
func (this *ZendBrkContElement) SetIsSwitch(value types.ZendBool) { this.is_switch = value }

/**
 * ZendLabel
 */
type ZendLabel struct {
	brk_cont   int
	opline_num uint32
}

func (this *ZendLabel) GetBrkCont() int           { return this.brk_cont }
func (this *ZendLabel) SetBrkCont(value int)      { this.brk_cont = value }
func (this *ZendLabel) GetOplineNum() uint32      { return this.opline_num }
func (this *ZendLabel) SetOplineNum(value uint32) { this.opline_num = value }

/**
 * ZendTryCatchElement
 */
type ZendTryCatchElement struct {
	try_op      uint32
	catch_op    uint32
	finally_op  uint32
	finally_end uint32
}

func (this *ZendTryCatchElement) GetTryOp() uint32           { return this.try_op }
func (this *ZendTryCatchElement) SetTryOp(value uint32)      { this.try_op = value }
func (this *ZendTryCatchElement) GetCatchOp() uint32         { return this.catch_op }
func (this *ZendTryCatchElement) SetCatchOp(value uint32)    { this.catch_op = value }
func (this *ZendTryCatchElement) GetFinallyOp() uint32       { return this.finally_op }
func (this *ZendTryCatchElement) SetFinallyOp(value uint32)  { this.finally_op = value }
func (this *ZendTryCatchElement) GetFinallyEnd() uint32      { return this.finally_end }
func (this *ZendTryCatchElement) SetFinallyEnd(value uint32) { this.finally_end = value }

/**
 * ZendLiveRange
 */
type ZendLiveRange struct {
	var_  uint32
	start uint32
	end   uint32
}

func (this *ZendLiveRange) GetVar() uint32        { return this.var_ }
func (this *ZendLiveRange) SetVar(value uint32)   { this.var_ = value }
func (this *ZendLiveRange) GetStart() uint32      { return this.start }
func (this *ZendLiveRange) SetStart(value uint32) { this.start = value }
func (this *ZendLiveRange) GetEnd() uint32        { return this.end }
func (this *ZendLiveRange) SetEnd(value uint32)   { this.end = value }

/**
 * ZendOparrayContext
 */
type ZendOparrayContext struct {
	opcodes_size     uint32
	vars_size        int
	literals_size    int
	fast_call_var    uint32
	try_catch_offset uint32
	current_brk_cont int
	last_brk_cont    int
	brk_cont_array   *ZendBrkContElement
	labels           *types.Array
}

//             func MakeZendOparrayContext(
// opcodes_size uint32,
// vars_size int,
// literals_size int,
// fast_call_var uint32,
// try_catch_offset uint32,
// current_brk_cont int,
// last_brk_cont int,
// brk_cont_array *ZendBrkContElement,
// labels *HashTable,
// ) ZendOparrayContext {
//                 return ZendOparrayContext{
//                     opcodes_size:opcodes_size,
//                     vars_size:vars_size,
//                     literals_size:literals_size,
//                     fast_call_var:fast_call_var,
//                     try_catch_offset:try_catch_offset,
//                     current_brk_cont:current_brk_cont,
//                     last_brk_cont:last_brk_cont,
//                     brk_cont_array:brk_cont_array,
//                     labels:labels,
//                 }
//             }
func (this *ZendOparrayContext) GetOpcodesSize() uint32               { return this.opcodes_size }
func (this *ZendOparrayContext) SetOpcodesSize(value uint32)          { this.opcodes_size = value }
func (this *ZendOparrayContext) GetVarsSize() int                     { return this.vars_size }
func (this *ZendOparrayContext) SetVarsSize(value int)                { this.vars_size = value }
func (this *ZendOparrayContext) GetLiteralsSize() int                 { return this.literals_size }
func (this *ZendOparrayContext) SetLiteralsSize(value int)            { this.literals_size = value }
func (this *ZendOparrayContext) GetFastCallVar() uint32               { return this.fast_call_var }
func (this *ZendOparrayContext) SetFastCallVar(value uint32)          { this.fast_call_var = value }
func (this *ZendOparrayContext) GetTryCatchOffset() uint32            { return this.try_catch_offset }
func (this *ZendOparrayContext) SetTryCatchOffset(value uint32)       { this.try_catch_offset = value }
func (this *ZendOparrayContext) GetCurrentBrkCont() int               { return this.current_brk_cont }
func (this *ZendOparrayContext) SetCurrentBrkCont(value int)          { this.current_brk_cont = value }
func (this *ZendOparrayContext) GetLastBrkCont() int                  { return this.last_brk_cont }
func (this *ZendOparrayContext) SetLastBrkCont(value int)             { this.last_brk_cont = value }
func (this *ZendOparrayContext) GetBrkContArray() *ZendBrkContElement { return this.brk_cont_array }
func (this *ZendOparrayContext) SetBrkContArray(value *ZendBrkContElement) {
	this.brk_cont_array = value
}
func (this *ZendOparrayContext) GetLabels() *types.Array      { return this.labels }
func (this *ZendOparrayContext) SetLabels(value *types.Array) { this.labels = value }

/**
 * ZendPropertyInfo
 */
type ZendPropertyInfo struct {
	offset      uint32
	flags       uint32
	name        *types.String
	doc_comment *types.String
	ce          *types.ClassEntry
	type_       types.ZendType
}

func (this *ZendPropertyInfo) GetOffset() uint32                 { return this.offset }
func (this *ZendPropertyInfo) SetOffset(value uint32)            { this.offset = value }
func (this *ZendPropertyInfo) GetFlags() uint32                  { return this.flags }
func (this *ZendPropertyInfo) SetFlags(value uint32)             { this.flags = value }
func (this *ZendPropertyInfo) GetName() *types.String            { return this.name }
func (this *ZendPropertyInfo) SetName(value *types.String)       { this.name = value }
func (this *ZendPropertyInfo) GetDocComment() *types.String      { return this.doc_comment }
func (this *ZendPropertyInfo) SetDocComment(value *types.String) { this.doc_comment = value }
func (this *ZendPropertyInfo) GetCe() *types.ClassEntry          { return this.ce }
func (this *ZendPropertyInfo) SetCe(value *types.ClassEntry)     { this.ce = value }
func (this *ZendPropertyInfo) GetType() types.ZendType           { return this.type_ }
func (this *ZendPropertyInfo) SetType(value types.ZendType)      { this.type_ = value }

/* ZendPropertyInfo.flags */
func (this *ZendPropertyInfo) AddFlags(value uint32)      { this.flags |= value }
func (this *ZendPropertyInfo) SubFlags(value uint32)      { this.flags &^= value }
func (this *ZendPropertyInfo) HasFlags(value uint32) bool { return this.flags&value != 0 }
func (this *ZendPropertyInfo) SwitchFlags(value uint32, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this ZendPropertyInfo) IsStatic() bool            { return this.HasFlags(AccStatic) }
func (this ZendPropertyInfo) IsProtected() bool         { return this.HasFlags(AccProtected) }
func (this ZendPropertyInfo) IsPrivate() bool           { return this.HasFlags(AccPrivate) }
func (this ZendPropertyInfo) IsPublic() bool            { return this.HasFlags(AccPublic) }
func (this ZendPropertyInfo) IsChanged() bool           { return this.HasFlags(AccChanged) }
func (this *ZendPropertyInfo) SetIsStatic(cond bool)    { this.SwitchFlags(AccStatic, cond) }
func (this *ZendPropertyInfo) SetIsProtected(cond bool) { this.SwitchFlags(AccProtected, cond) }
func (this *ZendPropertyInfo) SetIsPrivate(cond bool)   { this.SwitchFlags(AccPrivate, cond) }
func (this *ZendPropertyInfo) SetIsPublic(cond bool)    { this.SwitchFlags(AccPublic, cond) }
func (this *ZendPropertyInfo) SetIsChanged(cond bool)   { this.SwitchFlags(AccChanged, cond) }

/**
 * ZendClassConstant
 */
type ZendClassConstant struct {
	value       types.Zval
	doc_comment *types.String
	ce          *types.ClassEntry
}

// func MakeZendClassConstant(value Zval, doc_comment *String, ce *ClassEntry) ZendClassConstant {
//     return ZendClassConstant{
//         value:value,
//         doc_comment:doc_comment,
//         ce:ce,
//     }
// }
func (this *ZendClassConstant) GetValue() types.Zval { return this.value }

// func (this *ZendClassConstant) SetValue(value Zval) { this.value = value }
func (this *ZendClassConstant) GetDocComment() *types.String      { return this.doc_comment }
func (this *ZendClassConstant) SetDocComment(value *types.String) { this.doc_comment = value }
func (this *ZendClassConstant) GetCe() *types.ClassEntry          { return this.ce }
func (this *ZendClassConstant) SetCe(value *types.ClassEntry)     { this.ce = value }

/**
 * ZendArgInfo
 */
type ZendArgInfo struct {
	name              *types.String
	type_             types.ZendType
	pass_by_reference uint8
	is_variadic       types.ZendBool
}

func MakeZendReturnArgInfo(type_ types.ZendType, pass_by_reference bool) ZendArgInfo {
	return ZendArgInfo{
		name:              nil,
		type_:             type_,
		pass_by_reference: types.IntBool(pass_by_reference),
		is_variadic:       0,
	}
}

func MakeZendArgInfo(
	name *types.String,
	type_ types.ZendType,
	pass_by_reference uint8,
	is_variadic types.ZendBool,
) ZendArgInfo {
	return ZendArgInfo{
		name:              name,
		type_:             type_,
		pass_by_reference: pass_by_reference,
		is_variadic:       is_variadic,
	}
}

func (this *ZendArgInfo) SetType(value types.ZendType) { this.type_ = value }

func (this *ZendArgInfo) GetName() *types.String        { return this.name }
func (this *ZendArgInfo) GetType() types.ZendType       { return this.type_ }
func (this *ZendArgInfo) GetPassByReference() uint8     { return this.pass_by_reference }
func (this *ZendArgInfo) GetIsVariadic() types.ZendBool { return this.is_variadic }

/**
 * ZendInternalFunctionInfo
 */
type ZendInternalFunctionInfo struct {
	required_num_args types.ZendUintptrT
	type_             types.ZendType
	return_reference  types.ZendBool
	_is_variadic      types.ZendBool
}

func (this *ZendInternalFunctionInfo) GetRequiredNumArgs() types.ZendUintptrT {
	return this.required_num_args
}
func (this *ZendInternalFunctionInfo) GetType() types.ZendType { return this.type_ }
func (this *ZendInternalFunctionInfo) GetReturnReference() types.ZendBool {
	return this.return_reference
}

/**
 * ZendAutoGlobal
 */
type ZendAutoGlobal struct {
	name                 *types.String
	auto_global_callback ZendAutoGlobalCallback
	jit                  types.ZendBool
	armed                types.ZendBool
}

// func MakeZendAutoGlobal(name *String, auto_global_callback ZendAutoGlobalCallback, jit ZendBool, armed ZendBool) ZendAutoGlobal {
//     return ZendAutoGlobal{
//         name:name,
//         auto_global_callback:auto_global_callback,
//         jit:jit,
//         armed:armed,
//     }
// }
func (this *ZendAutoGlobal) GetName() *types.String      { return this.name }
func (this *ZendAutoGlobal) SetName(value *types.String) { this.name = value }
func (this *ZendAutoGlobal) GetAutoGlobalCallback() ZendAutoGlobalCallback {
	return this.auto_global_callback
}
func (this *ZendAutoGlobal) SetAutoGlobalCallback(value ZendAutoGlobalCallback) {
	this.auto_global_callback = value
}
func (this *ZendAutoGlobal) GetJit() types.ZendBool        { return this.jit }
func (this *ZendAutoGlobal) SetJit(value types.ZendBool)   { this.jit = value }
func (this *ZendAutoGlobal) GetArmed() types.ZendBool      { return this.armed }
func (this *ZendAutoGlobal) SetArmed(value types.ZendBool) { this.armed = value }

/**
 * ZendLoopVar
 */
type ZendLoopVar struct {
	opcode           uint8
	var_type         uint8
	var_num          uint32
	try_catch_offset uint32
}

func MakeZendLoopVar(opcode uint8, var_type uint8, var_num uint32, try_catch_offset uint32) ZendLoopVar {
	return ZendLoopVar{
		opcode:           opcode,
		var_type:         var_type,
		var_num:          var_num,
		try_catch_offset: try_catch_offset,
	}
}
func (this *ZendLoopVar) GetOpcode() uint8               { return this.opcode }
func (this *ZendLoopVar) SetOpcode(value uint8)          { this.opcode = value }
func (this *ZendLoopVar) GetVarType() uint8              { return this.var_type }
func (this *ZendLoopVar) SetVarType(value uint8)         { this.var_type = value }
func (this *ZendLoopVar) GetVarNum() uint32              { return this.var_num }
func (this *ZendLoopVar) SetVarNum(value uint32)         { this.var_num = value }
func (this *ZendLoopVar) GetTryCatchOffset() uint32      { return this.try_catch_offset }
func (this *ZendLoopVar) SetTryCatchOffset(value uint32) { this.try_catch_offset = value }

/**
 * ReservedClassName
 */
type ReservedClassName struct {
	name string
	len_ int
}

func MakeReservedClassName(name string) ReservedClassName {
	return ReservedClassName{name: name}
}
func (this *ReservedClassName) Name() string { return this.name }

func (this *ReservedClassName) GetName() *byte { return this.name }
func (this *ReservedClassName) GetLen() int    { return len(this.name) }

/**
 * BuiltinTypeInfo
 */
type BuiltinTypeInfo struct {
	name     *byte
	name_len int
	type_    uint8
}

func MakeBuiltinTypeInfo(name string, type_ uint8) BuiltinTypeInfo {
	return BuiltinTypeInfo{
		name:     name,
		name_len: name_len,
		type_:    type_,
	}
}
func (this *BuiltinTypeInfo) GetName() *byte  { return this.name }
func (this *BuiltinTypeInfo) GetNameLen() int { return this.name_len }
func (this *BuiltinTypeInfo) GetType() uint8  { return this.type_ }

/**
 * ClosureInfo
 */
type ClosureInfo struct {
	uses         types.Array
	varvars_used types.ZendBool
}

// func MakeClosureInfo(uses HashTable, varvars_used ZendBool) ClosureInfo {
//     return ClosureInfo{
//         uses:uses,
//         varvars_used:varvars_used,
//     }
// }
func (this *ClosureInfo) GetUses() types.Array { return this.uses }

// func (this *ClosureInfo) SetUses(value HashTable) { this.uses = value }
// func (this *ClosureInfo)  GetVarvarsUsed() ZendBool      { return this.varvars_used }
func (this *ClosureInfo) SetVarvarsUsed(value types.ZendBool) { this.varvars_used = value }
