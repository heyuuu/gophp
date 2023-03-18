// <<generate>>

package zend

import "sik/zend/types"

/**
 * ZnodeOp
 */
type ZnodeOp struct /* union */ {
	constant   uint32
	var_       uint32
	num        uint32
	opline_num uint32
	jmp_offset uint32
}

func (this *ZnodeOp) GetConstant() uint32       { return this.constant }
func (this *ZnodeOp) SetConstant(value uint32)  { this.constant = value }
func (this *ZnodeOp) GetVar() uint32            { return this.var_ }
func (this *ZnodeOp) SetVar(value uint32)       { this.var_ = value }
func (this *ZnodeOp) GetNum() uint32            { return this.num }
func (this *ZnodeOp) SetNum(value uint32)       { this.num = value }
func (this *ZnodeOp) GetOplineNum() uint32      { return this.opline_num }
func (this *ZnodeOp) SetOplineNum(value uint32) { this.opline_num = value }
func (this *ZnodeOp) GetJmpOffset() uint32      { return this.jmp_offset }
func (this *ZnodeOp) SetJmpOffset(value uint32) { this.jmp_offset = value }

/**
 * Znode
 */
type Znode struct {
	op_type types.ZendUchar
	flag    types.ZendUchar
	u       struct /* union */ {
		op       ZnodeOp
		constant types.Zval
	}
}

func (this *Znode) GetOpType() types.ZendUchar      { return this.op_type }
func (this *Znode) SetOpType(value types.ZendUchar) { this.op_type = value }
func (this *Znode) GetOp() ZnodeOp                  { return this.u.op }
func (this *Znode) SetOp(value ZnodeOp)             { this.u.op = value }
func (this *Znode) GetConstant() types.Zval         { return this.u.constant }

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
	current_namespace        *types.ZendString
	in_namespace             types.ZendBool
	has_bracketed_namespaces types.ZendBool
	imports                  *types.HashTable
	imports_function         *types.HashTable
	imports_const            *types.HashTable
	seen_symbols             types.HashTable
}

func (this *ZendFileContext) GetDeclarables() ZendDeclarables        { return this.declarables }
func (this *ZendFileContext) SetDeclarables(value ZendDeclarables)   { this.declarables = value }
func (this *ZendFileContext) GetCurrentNamespace() *types.ZendString { return this.current_namespace }
func (this *ZendFileContext) SetCurrentNamespace(value *types.ZendString) {
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
func (this *ZendFileContext) GetImports() *types.HashTable         { return this.imports }
func (this *ZendFileContext) SetImports(value *types.HashTable)    { this.imports = value }
func (this *ZendFileContext) GetImportsFunction() *types.HashTable { return this.imports_function }
func (this *ZendFileContext) SetImportsFunction(value *types.HashTable) {
	this.imports_function = value
}
func (this *ZendFileContext) GetImportsConst() *types.HashTable      { return this.imports_const }
func (this *ZendFileContext) SetImportsConst(value *types.HashTable) { this.imports_const = value }
func (this *ZendFileContext) GetSeenSymbols() types.HashTable        { return this.seen_symbols }

// func (this *ZendFileContext) SetSeenSymbols(value HashTable) { this.seen_symbols = value }

/**
 * ZendParserStackElem
 */
type ZendParserStackElem struct /* union */ {
	ast *ZendAst
	str *types.ZendString
	num ZendUlong
	ptr *uint8
}

func (this *ZendParserStackElem) GetAst() *ZendAst          { return this.ast }
func (this *ZendParserStackElem) SetAst(value *ZendAst)     { this.ast = value }
func (this *ZendParserStackElem) GetStr() *types.ZendString { return this.str }

/**
 * ZendOp
 */
type ZendOp struct {
	handler        any
	op1            ZnodeOp
	op2            ZnodeOp
	result         ZnodeOp
	extended_value uint32
	lineno         uint32
	opcode         types.ZendUchar
	op1_type       types.ZendUchar
	op2_type       types.ZendUchar
	result_type    types.ZendUchar
}

func (this *ZendOp) GetHandler() any                     { return this.handler }
func (this *ZendOp) SetHandler(value any)                { this.handler = value }
func (this *ZendOp) GetOp1() ZnodeOp                     { return this.op1 }
func (this *ZendOp) SetOp1(value ZnodeOp)                { this.op1 = value }
func (this *ZendOp) GetOp2() ZnodeOp                     { return this.op2 }
func (this *ZendOp) SetOp2(value ZnodeOp)                { this.op2 = value }
func (this *ZendOp) GetResult() ZnodeOp                  { return this.result }
func (this *ZendOp) SetResult(value ZnodeOp)             { this.result = value }
func (this *ZendOp) GetExtendedValue() uint32            { return this.extended_value }
func (this *ZendOp) SetExtendedValue(value uint32)       { this.extended_value = value }
func (this *ZendOp) GetLineno() uint32                   { return this.lineno }
func (this *ZendOp) SetLineno(value uint32)              { this.lineno = value }
func (this *ZendOp) GetOpcode() types.ZendUchar          { return this.opcode }
func (this *ZendOp) SetOpcode(value types.ZendUchar)     { this.opcode = value }
func (this *ZendOp) GetOp1Type() types.ZendUchar         { return this.op1_type }
func (this *ZendOp) SetOp1Type(value types.ZendUchar)    { this.op1_type = value }
func (this *ZendOp) GetOp2Type() types.ZendUchar         { return this.op2_type }
func (this *ZendOp) SetOp2Type(value types.ZendUchar)    { this.op2_type = value }
func (this *ZendOp) GetResultType() types.ZendUchar      { return this.result_type }
func (this *ZendOp) SetResultType(value types.ZendUchar) { this.result_type = value }

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

// func MakeZendBrkContElement(start int, cont int, brk int, parent int, is_switch ZendBool) ZendBrkContElement {
//     return ZendBrkContElement{
//         start:start,
//         cont:cont,
//         brk:brk,
//         parent:parent,
//         is_switch:is_switch,
//     }
// }
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

// func MakeZendLabel(brk_cont int, opline_num uint32) ZendLabel {
//     return ZendLabel{
//         brk_cont:brk_cont,
//         opline_num:opline_num,
//     }
// }
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
	labels           *types.HashTable
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
func (this *ZendOparrayContext) GetLabels() *types.HashTable      { return this.labels }
func (this *ZendOparrayContext) SetLabels(value *types.HashTable) { this.labels = value }

/**
 * ZendPropertyInfo
 */
type ZendPropertyInfo struct {
	offset      uint32
	flags       uint32
	name        *types.ZendString
	doc_comment *types.ZendString
	ce          *ZendClassEntry
	type_       types.ZendType
}

func (this *ZendPropertyInfo) GetOffset() uint32                     { return this.offset }
func (this *ZendPropertyInfo) SetOffset(value uint32)                { this.offset = value }
func (this *ZendPropertyInfo) GetFlags() uint32                      { return this.flags }
func (this *ZendPropertyInfo) SetFlags(value uint32)                 { this.flags = value }
func (this *ZendPropertyInfo) GetName() *types.ZendString            { return this.name }
func (this *ZendPropertyInfo) SetName(value *types.ZendString)       { this.name = value }
func (this *ZendPropertyInfo) GetDocComment() *types.ZendString      { return this.doc_comment }
func (this *ZendPropertyInfo) SetDocComment(value *types.ZendString) { this.doc_comment = value }
func (this *ZendPropertyInfo) GetCe() *ZendClassEntry                { return this.ce }
func (this *ZendPropertyInfo) SetCe(value *ZendClassEntry)           { this.ce = value }
func (this *ZendPropertyInfo) GetType() types.ZendType               { return this.type_ }
func (this *ZendPropertyInfo) SetType(value types.ZendType)          { this.type_ = value }

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
func (this ZendPropertyInfo) IsStatic() bool            { return this.HasFlags(ZEND_ACC_STATIC) }
func (this ZendPropertyInfo) IsProtected() bool         { return this.HasFlags(ZEND_ACC_PROTECTED) }
func (this ZendPropertyInfo) IsPrivate() bool           { return this.HasFlags(ZEND_ACC_PRIVATE) }
func (this ZendPropertyInfo) IsPublic() bool            { return this.HasFlags(ZEND_ACC_PUBLIC) }
func (this ZendPropertyInfo) IsChanged() bool           { return this.HasFlags(ZEND_ACC_CHANGED) }
func (this *ZendPropertyInfo) SetIsStatic(cond bool)    { this.SwitchFlags(ZEND_ACC_STATIC, cond) }
func (this *ZendPropertyInfo) SetIsProtected(cond bool) { this.SwitchFlags(ZEND_ACC_PROTECTED, cond) }
func (this *ZendPropertyInfo) SetIsPrivate(cond bool)   { this.SwitchFlags(ZEND_ACC_PRIVATE, cond) }
func (this *ZendPropertyInfo) SetIsPublic(cond bool)    { this.SwitchFlags(ZEND_ACC_PUBLIC, cond) }
func (this *ZendPropertyInfo) SetIsChanged(cond bool)   { this.SwitchFlags(ZEND_ACC_CHANGED, cond) }

/**
 * ZendClassConstant
 */
type ZendClassConstant struct {
	value       types.Zval
	doc_comment *types.ZendString
	ce          *ZendClassEntry
}

// func MakeZendClassConstant(value Zval, doc_comment *ZendString, ce *ZendClassEntry) ZendClassConstant {
//     return ZendClassConstant{
//         value:value,
//         doc_comment:doc_comment,
//         ce:ce,
//     }
// }
func (this *ZendClassConstant) GetValue() types.Zval { return this.value }

// func (this *ZendClassConstant) SetValue(value Zval) { this.value = value }
func (this *ZendClassConstant) GetDocComment() *types.ZendString      { return this.doc_comment }
func (this *ZendClassConstant) SetDocComment(value *types.ZendString) { this.doc_comment = value }
func (this *ZendClassConstant) GetCe() *ZendClassEntry                { return this.ce }
func (this *ZendClassConstant) SetCe(value *ZendClassEntry)           { this.ce = value }

/**
 * ZendArgInfo
 */
type ZendArgInfo struct {
	name              *types.ZendString
	type_             types.ZendType
	pass_by_reference types.ZendUchar
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
	name *types.ZendString,
	type_ types.ZendType,
	pass_by_reference types.ZendUchar,
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

func (this *ZendArgInfo) GetName() *types.ZendString          { return this.name }
func (this *ZendArgInfo) GetType() types.ZendType             { return this.type_ }
func (this *ZendArgInfo) GetPassByReference() types.ZendUchar { return this.pass_by_reference }
func (this *ZendArgInfo) GetIsVariadic() types.ZendBool       { return this.is_variadic }

/**
 * ZendInternalFunctionInfo
 */
type ZendInternalFunctionInfo struct {
	required_num_args types.ZendUintptrT
	type_             types.ZendType
	return_reference  types.ZendBool
	_is_variadic      types.ZendBool
}

// func MakeZendInternalFunctionInfo(required_num_args ZendUintptrT, type_ ZendType, return_reference ZendBool, _is_variadic ZendBool) ZendInternalFunctionInfo {
//     return ZendInternalFunctionInfo{
//         required_num_args:required_num_args,
//         type_:type_,
//         return_reference:return_reference,
//         _is_variadic:_is_variadic,
//     }
// }
func (this *ZendInternalFunctionInfo) GetRequiredNumArgs() types.ZendUintptrT {
	return this.required_num_args
}

func (this *ZendInternalFunctionInfo) GetType() types.ZendType { return this.type_ }
func (this *ZendInternalFunctionInfo) GetReturnReference() types.ZendBool {
	return this.return_reference
}

/**
 * ZendOpArray
 */
type ZendOpArray struct {
	type_                     types.ZendUchar
	arg_flags                 []types.ZendUchar
	fn_flags                  uint32
	function_name             *types.ZendString
	scope                     *ZendClassEntry
	prototype                 *ZendFunction
	num_args                  uint32
	required_num_args         uint32
	arg_info                  *ZendArgInfo
	cache_size                int
	last_var                  int
	T                         uint32
	last                      uint32
	opcodes                   *ZendOp
	run_time_cache__ptr       **any
	static_variables_ptr__ptr **types.HashTable
	static_variables          *types.HashTable
	vars                      **types.ZendString
	refcount                  *uint32
	last_live_range           int
	last_try_catch            int
	live_range                *ZendLiveRange
	try_catch_array           *ZendTryCatchElement
	filename                  *types.ZendString
	line_start                uint32
	line_end                  uint32
	doc_comment               *types.ZendString
	last_literal              int
	literals                  *types.Zval
	reserved                  []any
}

//             func MakeZendOpArray(
// type_ ZendUchar,
// arg_flags []ZendUchar,
// fn_flags uint32,
// function_name *ZendString,
// scope *ZendClassEntry,
// prototype *ZendFunction,
// num_args uint32,
// required_num_args uint32,
// arg_info *ZendArgInfo,
// cache_size int,
// last_var int,
// T uint32,
// last uint32,
// opcodes *ZendOp,
// run_time_cache__ptr **any,
// static_variables_ptr__ptr **HashTable,
// static_variables *HashTable,
// vars **ZendString,
// refcount *uint32,
// last_live_range int,
// last_try_catch int,
// live_range *ZendLiveRange,
// try_catch_array *ZendTryCatchElement,
// filename *ZendString,
// line_start uint32,
// line_end uint32,
// doc_comment *ZendString,
// last_literal int,
// literals *Zval,
// reserved []any,
// ) ZendOpArray {
//                 return ZendOpArray{
//                     type_:type_,
//                     arg_flags:arg_flags,
//                     fn_flags:fn_flags,
//                     function_name:function_name,
//                     scope:scope,
//                     prototype:prototype,
//                     num_args:num_args,
//                     required_num_args:required_num_args,
//                     arg_info:arg_info,
//                     cache_size:cache_size,
//                     last_var:last_var,
//                     T:T,
//                     last:last,
//                     opcodes:opcodes,
//                     run_time_cache__ptr:run_time_cache__ptr,
//                     static_variables_ptr__ptr:static_variables_ptr__ptr,
//                     static_variables:static_variables,
//                     vars:vars,
//                     refcount:refcount,
//                     last_live_range:last_live_range,
//                     last_try_catch:last_try_catch,
//                     live_range:live_range,
//                     try_catch_array:try_catch_array,
//                     filename:filename,
//                     line_start:line_start,
//                     line_end:line_end,
//                     doc_comment:doc_comment,
//                     last_literal:last_literal,
//                     literals:literals,
//                     reserved:reserved,
//                 }
//             }
func (this *ZendOpArray) GetType() types.ZendUchar       { return this.type_ }
func (this *ZendOpArray) SetType(value types.ZendUchar)  { this.type_ = value }
func (this *ZendOpArray) GetArgFlags() []types.ZendUchar { return this.arg_flags }

// func (this *ZendOpArray) SetArgFlags(value []ZendUchar) { this.arg_flags = value }
func (this *ZendOpArray) GetFnFlags() uint32                      { return this.fn_flags }
func (this *ZendOpArray) SetFnFlags(value uint32)                 { this.fn_flags = value }
func (this *ZendOpArray) GetFunctionName() *types.ZendString      { return this.function_name }
func (this *ZendOpArray) SetFunctionName(value *types.ZendString) { this.function_name = value }
func (this *ZendOpArray) GetScope() *ZendClassEntry               { return this.scope }
func (this *ZendOpArray) SetScope(value *ZendClassEntry)          { this.scope = value }

// func (this *ZendOpArray)  GetPrototype() *ZendFunction      { return this.prototype }
func (this *ZendOpArray) SetPrototype(value *ZendFunction) { this.prototype = value }
func (this *ZendOpArray) GetNumArgs() uint32               { return this.num_args }
func (this *ZendOpArray) SetNumArgs(value uint32)          { this.num_args = value }

// func (this *ZendOpArray)  GetRequiredNumArgs() uint32      { return this.required_num_args }
func (this *ZendOpArray) SetRequiredNumArgs(value uint32) { this.required_num_args = value }
func (this *ZendOpArray) GetArgInfo() *ZendArgInfo        { return this.arg_info }
func (this *ZendOpArray) SetArgInfo(value *ZendArgInfo)   { this.arg_info = value }
func (this *ZendOpArray) GetCacheSize() int               { return this.cache_size }
func (this *ZendOpArray) SetCacheSize(value int)          { this.cache_size = value }
func (this *ZendOpArray) GetLastVar() int                 { return this.last_var }
func (this *ZendOpArray) SetLastVar(value int)            { this.last_var = value }
func (this *ZendOpArray) GetT() uint32                    { return this.T }
func (this *ZendOpArray) SetT(value uint32)               { this.T = value }
func (this *ZendOpArray) GetLast() uint32                 { return this.last }
func (this *ZendOpArray) SetLast(value uint32)            { this.last = value }
func (this *ZendOpArray) GetOpcodes() *ZendOp             { return this.opcodes }
func (this *ZendOpArray) SetOpcodes(value *ZendOp)        { this.opcodes = value }
func (this *ZendOpArray) GetRunTimeCachePtr() **any       { return this.run_time_cache__ptr }

// func (this *ZendOpArray) SetRunTimeCachePtr(value **any) { this.run_time_cache__ptr = value }
// func (this *ZendOpArray)  GetStaticVariablesPtrPtr() **HashTable      { return this.static_variables_ptr__ptr }
// func (this *ZendOpArray) SetStaticVariablesPtrPtr(value **HashTable) { this.static_variables_ptr__ptr = value }
func (this *ZendOpArray) GetStaticVariables() *types.HashTable        { return this.static_variables }
func (this *ZendOpArray) SetStaticVariables(value *types.HashTable)   { this.static_variables = value }
func (this *ZendOpArray) GetVars() **types.ZendString                 { return this.vars }
func (this *ZendOpArray) SetVars(value **types.ZendString)            { this.vars = value }
func (this *ZendOpArray) GetRefcount() *uint32                        { return this.refcount }
func (this *ZendOpArray) SetRefcount(value *uint32)                   { this.refcount = value }
func (this *ZendOpArray) GetLastLiveRange() int                       { return this.last_live_range }
func (this *ZendOpArray) SetLastLiveRange(value int)                  { this.last_live_range = value }
func (this *ZendOpArray) GetLastTryCatch() int                        { return this.last_try_catch }
func (this *ZendOpArray) SetLastTryCatch(value int)                   { this.last_try_catch = value }
func (this *ZendOpArray) GetLiveRange() *ZendLiveRange                { return this.live_range }
func (this *ZendOpArray) SetLiveRange(value *ZendLiveRange)           { this.live_range = value }
func (this *ZendOpArray) GetTryCatchArray() *ZendTryCatchElement      { return this.try_catch_array }
func (this *ZendOpArray) SetTryCatchArray(value *ZendTryCatchElement) { this.try_catch_array = value }
func (this *ZendOpArray) GetFilename() *types.ZendString              { return this.filename }
func (this *ZendOpArray) SetFilename(value *types.ZendString)         { this.filename = value }
func (this *ZendOpArray) GetLineStart() uint32                        { return this.line_start }
func (this *ZendOpArray) SetLineStart(value uint32)                   { this.line_start = value }
func (this *ZendOpArray) GetLineEnd() uint32                          { return this.line_end }
func (this *ZendOpArray) SetLineEnd(value uint32)                     { this.line_end = value }
func (this *ZendOpArray) GetDocComment() *types.ZendString            { return this.doc_comment }
func (this *ZendOpArray) SetDocComment(value *types.ZendString)       { this.doc_comment = value }
func (this *ZendOpArray) GetLastLiteral() int                         { return this.last_literal }
func (this *ZendOpArray) SetLastLiteral(value int)                    { this.last_literal = value }
func (this *ZendOpArray) GetLiterals() *types.Zval                    { return this.literals }
func (this *ZendOpArray) SetLiterals(value *types.Zval)               { this.literals = value }
func (this *ZendOpArray) GetReserved() []any                          { return this.reserved }

// func (this *ZendOpArray) SetReserved(value []any) { this.reserved = value }

/* ZendOpArray.fn_flags */
func (this *ZendOpArray) AddFnFlags(value uint32)      { this.fn_flags |= value }
func (this *ZendOpArray) SubFnFlags(value uint32)      { this.fn_flags &^= value }
func (this *ZendOpArray) HasFnFlags(value uint32) bool { return this.fn_flags&value != 0 }
func (this *ZendOpArray) SwitchFnFlags(value uint32, cond bool) {
	if cond {
		this.AddFnFlags(value)
	} else {
		this.SubFnFlags(value)
	}
}
func (this ZendOpArray) IsChanged() bool         { return this.HasFnFlags(ZEND_ACC_CHANGED) }
func (this ZendOpArray) IsPreloaded() bool       { return this.HasFnFlags(ZEND_ACC_PRELOADED) }
func (this ZendOpArray) IsHasReturnType() bool   { return this.HasFnFlags(ZEND_ACC_HAS_RETURN_TYPE) }
func (this ZendOpArray) IsEarlyBinding() bool    { return this.HasFnFlags(ZEND_ACC_EARLY_BINDING) }
func (this ZendOpArray) IsHeapRtCache() bool     { return this.HasFnFlags(ZEND_ACC_HEAP_RT_CACHE) }
func (this ZendOpArray) IsClosure() bool         { return this.HasFnFlags(ZEND_ACC_CLOSURE) }
func (this ZendOpArray) IsReturnReference() bool { return this.HasFnFlags(ZEND_ACC_RETURN_REFERENCE) }
func (this ZendOpArray) IsGenerator() bool       { return this.HasFnFlags(ZEND_ACC_GENERATOR) }
func (this ZendOpArray) IsHasFinallyBlock() bool { return this.HasFnFlags(ZEND_ACC_HAS_FINALLY_BLOCK) }
func (this ZendOpArray) IsVariadic() bool        { return this.HasFnFlags(ZEND_ACC_VARIADIC) }
func (this ZendOpArray) IsPublic() bool          { return this.HasFnFlags(ZEND_ACC_PUBLIC) }
func (this ZendOpArray) IsStatic() bool          { return this.HasFnFlags(ZEND_ACC_STATIC) }
func (this ZendOpArray) IsAbstract() bool        { return this.HasFnFlags(ZEND_ACC_ABSTRACT) }
func (this ZendOpArray) IsPrivate() bool         { return this.HasFnFlags(ZEND_ACC_PRIVATE) }
func (this ZendOpArray) IsHasTypeHints() bool    { return this.HasFnFlags(ZEND_ACC_HAS_TYPE_HINTS) }
func (this ZendOpArray) IsCallViaTrampoline() bool {
	return this.HasFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE)
}
func (this ZendOpArray) IsDonePassTwo() bool       { return this.HasFnFlags(ZEND_ACC_DONE_PASS_TWO) }
func (this ZendOpArray) IsUsesThis() bool          { return this.HasFnFlags(ZEND_ACC_USES_THIS) }
func (this ZendOpArray) IsStrictTypes() bool       { return this.HasFnFlags(ZEND_ACC_STRICT_TYPES) }
func (this ZendOpArray) IsAllowStatic() bool       { return this.HasFnFlags(ZEND_ACC_ALLOW_STATIC) }
func (this ZendOpArray) IsTopLevel() bool          { return this.HasFnFlags(ZEND_ACC_TOP_LEVEL) }
func (this ZendOpArray) IsTraitClone() bool        { return this.HasFnFlags(ZEND_ACC_TRAIT_CLONE) }
func (this ZendOpArray) IsImmutable() bool         { return this.HasFnFlags(ZEND_ACC_IMMUTABLE) }
func (this *ZendOpArray) SetIsChanged(cond bool)   { this.SwitchFnFlags(ZEND_ACC_CHANGED, cond) }
func (this *ZendOpArray) SetIsPreloaded(cond bool) { this.SwitchFnFlags(ZEND_ACC_PRELOADED, cond) }
func (this *ZendOpArray) SetIsHasReturnType(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_RETURN_TYPE, cond)
}
func (this *ZendOpArray) SetIsEarlyBinding(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_EARLY_BINDING, cond)
}
func (this *ZendOpArray) SetIsHeapRtCache(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HEAP_RT_CACHE, cond)
}
func (this *ZendOpArray) SetIsClosure(cond bool) { this.SwitchFnFlags(ZEND_ACC_CLOSURE, cond) }
func (this *ZendOpArray) SetIsReturnReference(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_RETURN_REFERENCE, cond)
}
func (this *ZendOpArray) SetIsGenerator(cond bool) { this.SwitchFnFlags(ZEND_ACC_GENERATOR, cond) }
func (this *ZendOpArray) SetIsHasFinallyBlock(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_FINALLY_BLOCK, cond)
}
func (this *ZendOpArray) SetIsVariadic(cond bool) { this.SwitchFnFlags(ZEND_ACC_VARIADIC, cond) }
func (this *ZendOpArray) SetIsPublic(cond bool)   { this.SwitchFnFlags(ZEND_ACC_PUBLIC, cond) }
func (this *ZendOpArray) SetIsStatic(cond bool)   { this.SwitchFnFlags(ZEND_ACC_STATIC, cond) }
func (this *ZendOpArray) SetIsAbstract(cond bool) { this.SwitchFnFlags(ZEND_ACC_ABSTRACT, cond) }
func (this *ZendOpArray) SetIsPrivate(cond bool)  { this.SwitchFnFlags(ZEND_ACC_PRIVATE, cond) }
func (this *ZendOpArray) SetIsHasTypeHints(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_HAS_TYPE_HINTS, cond)
}
func (this *ZendOpArray) SetIsCallViaTrampoline(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_CALL_VIA_TRAMPOLINE, cond)
}
func (this *ZendOpArray) SetIsDonePassTwo(cond bool) {
	this.SwitchFnFlags(ZEND_ACC_DONE_PASS_TWO, cond)
}
func (this *ZendOpArray) SetIsUsesThis(cond bool)    { this.SwitchFnFlags(ZEND_ACC_USES_THIS, cond) }
func (this *ZendOpArray) SetIsStrictTypes(cond bool) { this.SwitchFnFlags(ZEND_ACC_STRICT_TYPES, cond) }
func (this *ZendOpArray) SetIsAllowStatic(cond bool) { this.SwitchFnFlags(ZEND_ACC_ALLOW_STATIC, cond) }
func (this *ZendOpArray) SetIsTopLevel(cond bool)    { this.SwitchFnFlags(ZEND_ACC_TOP_LEVEL, cond) }
func (this *ZendOpArray) SetIsTraitClone(cond bool)  { this.SwitchFnFlags(ZEND_ACC_TRAIT_CLONE, cond) }
func (this *ZendOpArray) SetIsImmutable(cond bool)   { this.SwitchFnFlags(ZEND_ACC_IMMUTABLE, cond) }

/**
 * ZendAutoGlobal
 */
type ZendAutoGlobal struct {
	name                 *types.ZendString
	auto_global_callback ZendAutoGlobalCallback
	jit                  types.ZendBool
	armed                types.ZendBool
}

// func MakeZendAutoGlobal(name *ZendString, auto_global_callback ZendAutoGlobalCallback, jit ZendBool, armed ZendBool) ZendAutoGlobal {
//     return ZendAutoGlobal{
//         name:name,
//         auto_global_callback:auto_global_callback,
//         jit:jit,
//         armed:armed,
//     }
// }
func (this *ZendAutoGlobal) GetName() *types.ZendString      { return this.name }
func (this *ZendAutoGlobal) SetName(value *types.ZendString) { this.name = value }
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
	opcode           types.ZendUchar
	var_type         types.ZendUchar
	var_num          uint32
	try_catch_offset uint32
}

func MakeZendLoopVar(opcode types.ZendUchar, var_type types.ZendUchar, var_num uint32, try_catch_offset uint32) ZendLoopVar {
	return ZendLoopVar{
		opcode:           opcode,
		var_type:         var_type,
		var_num:          var_num,
		try_catch_offset: try_catch_offset,
	}
}
func (this *ZendLoopVar) GetOpcode() types.ZendUchar       { return this.opcode }
func (this *ZendLoopVar) SetOpcode(value types.ZendUchar)  { this.opcode = value }
func (this *ZendLoopVar) GetVarType() types.ZendUchar      { return this.var_type }
func (this *ZendLoopVar) SetVarType(value types.ZendUchar) { this.var_type = value }
func (this *ZendLoopVar) GetVarNum() uint32                { return this.var_num }
func (this *ZendLoopVar) SetVarNum(value uint32)           { this.var_num = value }
func (this *ZendLoopVar) GetTryCatchOffset() uint32        { return this.try_catch_offset }
func (this *ZendLoopVar) SetTryCatchOffset(value uint32)   { this.try_catch_offset = value }

/**
 * ReservedClassName
 */
type ReservedClassName struct {
	name *byte
	len_ int
}

func MakeReservedClassName(name *byte, len_ int) ReservedClassName {
	return ReservedClassName{
		name: name,
		len_: len_,
	}
}
func (this *ReservedClassName) GetName() *byte { return this.name }
func (this *ReservedClassName) GetLen() int    { return this.len_ }

/**
 * BuiltinTypeInfo
 */
type BuiltinTypeInfo struct {
	name     *byte
	name_len int
	type_    types.ZendUchar
}

func MakeBuiltinTypeInfo(name *byte, name_len int, type_ types.ZendUchar) BuiltinTypeInfo {
	return BuiltinTypeInfo{
		name:     name,
		name_len: name_len,
		type_:    type_,
	}
}
func (this *BuiltinTypeInfo) GetName() *byte { return this.name }

// func (this *BuiltinTypeInfo) SetName(value *byte) { this.name = value }
func (this *BuiltinTypeInfo) GetNameLen() int { return this.name_len }

// func (this *BuiltinTypeInfo) SetNameLen(value int) { this.name_len = value }
func (this *BuiltinTypeInfo) GetType() types.ZendUchar { return this.type_ }

// func (this *BuiltinTypeInfo) SetType(value ZendUchar) { this.type_ = value }

/**
 * ClosureInfo
 */
type ClosureInfo struct {
	uses         types.HashTable
	varvars_used types.ZendBool
}

// func MakeClosureInfo(uses HashTable, varvars_used ZendBool) ClosureInfo {
//     return ClosureInfo{
//         uses:uses,
//         varvars_used:varvars_used,
//     }
// }
func (this *ClosureInfo) GetUses() types.HashTable { return this.uses }

// func (this *ClosureInfo) SetUses(value HashTable) { this.uses = value }
// func (this *ClosureInfo)  GetVarvarsUsed() ZendBool      { return this.varvars_used }
func (this *ClosureInfo) SetVarvarsUsed(value types.ZendBool) { this.varvars_used = value }
