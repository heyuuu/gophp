package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

/**
 * Znode
 */
type Znode struct {
	op_type uint8
	flag    uint8
	u       struct /* union */ {
		op       types.ZnodeOp
		constant types.Zval
	}
}

func (this *Znode) GetOpType() uint8          { return this.op_type }
func (this *Znode) SetOpType(value uint8)     { this.op_type = value }
func (this *Znode) GetOp() types.ZnodeOp      { return this.u.op }
func (this *Znode) SetOp(value types.ZnodeOp) { this.u.op = value }
func (this *Znode) GetConstant() *types.Zval  { return &this.u.constant }

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
 * ZendFileContext
 */
type ImportNames = *types.Table[string]
type ZendFileContext struct {
	current_namespace        *types.String
	in_namespace             types.ZendBool
	has_bracketed_namespaces types.ZendBool
	imports                  ImportNames
	imports_function         ImportNames
	imports_const            ImportNames
	seen_symbols             types.Array
}

func (this *ZendFileContext) ResetImportTables() {
	this.imports = nil
	this.imports_function = nil
	this.imports_const = nil
}

func (this *ZendFileContext) Imports() ImportNames {
	if this.imports == nil {
		this.imports = types.NewLcTable[string](nil)
	}
	return this.imports
}
func (this *ZendFileContext) ImportsFunction() ImportNames {
	if this.imports_function == nil {
		this.imports_function = types.NewLcTable[string](nil)
	}
	return this.imports_function
}
func (this *ZendFileContext) ImportsConst() ImportNames {
	if this.imports_const == nil {
		this.imports_const = types.NewLcTable[string](nil)
	}
	return this.imports_const
}

func (this *ZendFileContext) GetImports() ImportNames          { return this.imports }
func (this *ZendFileContext) SetImports(value *types.Array)    { this.imports = value }
func (this *ZendFileContext) GetImportsFunction() *types.Array { return this.imports_function }
func (this *ZendFileContext) SetImportsFunction(value *types.Array) {
	this.imports_function = value
}
func (this *ZendFileContext) GetImportsConst() *types.Array      { return this.imports_const }
func (this *ZendFileContext) SetImportsConst(value *types.Array) { this.imports_const = value }

func (this *ZendFileContext) GetCurrentNamespace() *types.String { return this.current_namespace }
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
func (this *ZendFileContext) GetSeenSymbols() types.Array { return this.seen_symbols }

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
	fast_call_var    uint32
	try_catch_offset uint32
	current_brk_cont int
	last_brk_cont    int
	brk_cont_array   *ZendBrkContElement
	labels           *types.Array
}

func NewOpArrayContext() *ZendOparrayContext {
	return &ZendOparrayContext{
		opcodes_size:     INITIAL_OP_ARRAY_SIZE,
		fast_call_var:    -1,
		try_catch_offset: -1,
		current_brk_cont: -1,
		last_brk_cont:    0,
		brk_cont_array:   nil,
		labels:           nil,
	}
}

func (this *ZendOparrayContext) GetOpcodesSize() uint32               { return this.opcodes_size }
func (this *ZendOparrayContext) SetOpcodesSize(value uint32)          { this.opcodes_size = value }
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
 * ZendArgInfo
 */
type ZendArgInfo struct {
	name              *types.String
	type_             types.TypeHint
	pass_by_reference uint8
	is_variadic       types.ZendBool
}

func MakeZendReturnArgInfo(type_ types.TypeHint, pass_by_reference bool) ZendArgInfo {
	return ZendArgInfo{
		name:              nil,
		type_:             type_,
		pass_by_reference: types.IntBool(pass_by_reference),
		is_variadic:       0,
	}
}

func MakeZendArgInfo(
	name *types.String,
	type_ types.TypeHint,
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

func (this *ZendArgInfo) SetType(value types.TypeHint) { this.type_ = value }

func (this *ZendArgInfo) GetName() *types.String        { return this.name }
func (this *ZendArgInfo) GetType() types.TypeHint       { return this.type_ }
func (this *ZendArgInfo) GetPassByReference() uint8     { return this.pass_by_reference }
func (this *ZendArgInfo) GetIsVariadic() types.ZendBool { return this.is_variadic }

/**
 * ZendInternalFunctionInfo
 */
type ZendInternalFunctionInfo struct {
	required_num_args types.ZendUintptrT
	type_             types.TypeHint
	return_reference  types.ZendBool
	_is_variadic      types.ZendBool
}

func (this *ZendInternalFunctionInfo) GetRequiredNumArgs() types.ZendUintptrT {
	return this.required_num_args
}
func (this *ZendInternalFunctionInfo) GetType() types.TypeHint { return this.type_ }
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
 * ClosureInfo
 */
type ClosureInfo struct {
	uses         *types.Array
	varvars_used types.ZendBool
}

func (this *ClosureInfo) GetUses() *types.Array               { return this.uses }
func (this *ClosureInfo) SetUses(value *types.Array)          { this.uses = value }
func (this *ClosureInfo) SetVarvarsUsed(value types.ZendBool) { this.varvars_used = value }
