// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/zend/types"
)

func MAKE_NOP(opline *ZendOp) {
	opline.GetOp1().SetNum(0)
	opline.GetOp2().SetNum(0)
	opline.GetResult().SetNum(0)
	opline.SetOpcode(ZEND_NOP)
	opline.SetOp1Type(IS_UNUSED)
	opline.SetOp2Type(IS_UNUSED)
	opline.SetResultType(IS_UNUSED)
}
func RESET_DOC_COMMENT() {
	if CG__().GetDocComment() != nil {
		types.ZendStringReleaseEx(CG__().GetDocComment(), 0)
		CG__().SetDocComment(nil)
	}
}
func ZendAstGetZnode(ast *ZendAst) *Znode { return (*ZendAstZnode)(ast).GetNode() }
func OBJ_PROP(obj *types.ZendObject, offset *types.ZendObject) *types.Zval {
	return (*types.Zval)((*byte)(obj + offset))
}
func OBJ_PROP_NUM(obj __auto__, num __auto__) __auto__ { return obj.properties_table[num] }
func OBJ_PROP_TO_OFFSET(num int) __auto__ {
	return uint32(zend_long((*byte)(&((*types.ZendObject)(nil).GetPropertiesTable()))-(*byte)(nil)) + b.SizeOf("zval")*num)
}
func OBJ_PROP_TO_NUM(offset uint32) int {
	return (offset - OBJ_PROP_TO_OFFSET(0)) / b.SizeOf("zval")
}
func ZEND_FN_SCOPE_NAME(function *ZendFunction) string {
	if function != nil && function.GetScope() != nil {
		return function.GetScope().GetName().GetStr()
	} else {
		return ""
	}
}
func ZEND_CALL_INFO(call *ZendExecuteData) uint32         { return call.GetThis().GetTypeInfo() }
func ZEND_ADD_CALL_FLAG_EX(call_info uint32, flag uint32) { call_info |= flag }
func ZEND_DEL_CALL_FLAG_EX(call_info uint32, flag uint32) { call_info &= ^flag }
func ZEND_ADD_CALL_FLAG(call *ZendExecuteData, flag uint32) {
	ZEND_ADD_CALL_FLAG_EX(call.GetThis().GetTypeInfo(), flag)
}
func ZEND_DEL_CALL_FLAG(call *ZendExecuteData, flag uint32) {
	ZEND_DEL_CALL_FLAG_EX(call.This.GetTypeInfo(), flag)
}
func ZEND_CALL_VAR(call *ZendExecuteData, n uint32) *types.Zval {
	return (*types.Zval)((*byte)(call) + int(n))
}
func EX_CALL_INFO() uint32        { return ZEND_CALL_INFO(executeData) }
func EX_VAR(n uint32) *types.Zval { return ZEND_CALL_VAR(executeData, n) }
func EX_VAR_TO_NUM(n uint32) __auto__ {
	return uint32(ZEND_CALL_VAR(nil, n) - nil.VarNum(0))
}
func ZEND_OPLINE_NUM_TO_OFFSET(op_array *ZendOpArray, opline *ZendOp, opline_num uint32) *byte {
	return (*byte)(op_array.GetOpcodes()[opline_num] - (*byte)(opline))
}
func ZEND_OFFSET_TO_OPLINE(base *ZendOp, offset uint32) *ZendOp {
	return (*ZendOp)((*byte)(base) + int(offset))
}
func OP_JMP_ADDR(opline *ZendOp, node ZnodeOp) *ZendOp {
	return ZEND_OFFSET_TO_OPLINE(opline, node.GetJmpOffset())
}
func ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array *ZendOpArray, opline *ZendOp, node ZnodeOp) {
	node.SetJmpOffset(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, node.GetOplineNum()))
}
func CT_CONSTANT_EX(op_array *ZendOpArray, num *types.Zval) __auto__ {
	return op_array.GetLiterals() + num
}
func CT_CONSTANT(node ZnodeOp) __auto__ {
	return CT_CONSTANT_EX(CG__().GetActiveOpArray(), node.GetConstant())
}
func RT_CONSTANT(opline *ZendOp, node ZnodeOp) *types.Zval {
	return (*types.Zval)((*byte)(opline) + int32(node).constant)
}
func ZEND_PASS_TWO_UPDATE_CONSTANT(op_array *ZendOpArray, opline *ZendOp, node ZnodeOp) {
	node.SetConstant((*byte)(CT_CONSTANT_EX(op_array, node.GetConstant())) - (*byte)(opline))
}
func RUN_TIME_CACHE(op_array __auto__) any {
	return ZEND_MAP_PTR_GET(op_array.run_time_cache)
}
func ZendUnmanglePropertyName(mangled_property *types.ZendString, class_name **byte, prop_name **byte) int {
	return ZendUnmanglePropertyNameEx(mangled_property, class_name, prop_name, nil)
}
func ZendGetUnmangledPropertyName(mangled_prop *types.ZendString) *byte {
	var class_name *byte
	var prop_name *byte
	ZendUnmanglePropertyName(mangled_prop, &class_name, &prop_name)
	return prop_name
}
func ZEND_USER_CODE(type_ types.ZendUchar) bool { return (type_ & 1) == 0 }
func ZendCheckArgSendType(zf *ZendFunction, arg_num uint32, mask uint32) int {
	return types.IntBool(zf.CheckArgSendType(arg_num, uint8(mask)))
}
func ARG_MUST_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF)
}
func ARG_SHOULD_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func ARG_MAY_BE_SENT_BY_REF(zf *ZendFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_PREFER_REF)
}

func ZEND_CHECK_ARG_FLAG(zf *ZendFunction, arg_num uint32, mask uint8) int {
	return types.IntBool(zf.CheckArgSendType(arg_num, mask))
}
func QUICK_ARG_MUST_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF)
}
func QUICK_ARG_SHOULD_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func QUICK_ARG_MAY_BE_SENT_BY_REF(zf *ZendFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_PREFER_REF)
}
func ZEND_IS_INCREMENT(opcode types.ZendUchar) bool { return (opcode & 1) == 0 }
func ZEND_IS_BINARY_ASSIGN_OP_OPCODE(opcode __auto__) bool {
	return opcode >= ZEND_ADD && opcode <= ZEND_POW
}
func ZendAllocCacheSlots(count unsigned) uint32 {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var ret uint32 = op_array.GetCacheSize()
	op_array.SetCacheSize(op_array.GetCacheSize() + count*b.SizeOf("void *"))
	return ret
}
func ZendAllocCacheSlot() uint32 { return ZendAllocCacheSlots(1) }
func InitOp(op *ZendOp) {
	MAKE_NOP(op)
	op.SetExtendedValue(0)
	op.SetLineno(CG__().GetZendLineno())
}
func GetNextOpNumber() uint32 {
	return CG__().GetActiveOpArray().GetLast()
}
func GetNextOp() *ZendOp {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var next_op_num uint32 = b.PostInc(&(op_array.GetLast()))
	var next_op *ZendOp
	if next_op_num >= CG__().GetContext().GetOpcodesSize() {
		CG__().GetContext().SetOpcodesSize(CG__().GetContext().GetOpcodesSize() * 4)
		op_array.SetOpcodes(Erealloc(op_array.GetOpcodes(), CG__().GetContext().GetOpcodesSize()*b.SizeOf("zend_op")))
	}
	next_op = &op_array.GetOpcodes()[next_op_num]
	InitOp(next_op)
	return next_op
}
func GetNextBrkContElement() *ZendBrkContElement {
	CG__().GetContext().GetLastBrkCont()++
	CG__().GetContext().SetBrkContArray(Erealloc(CG__().GetContext().GetBrkContArray(), b.SizeOf("zend_brk_cont_element")*CG__().GetContext().GetLastBrkCont()))
	return CG__().GetContext().GetBrkContArray()[CG__().GetContext().GetLastBrkCont()-1]
}
func ZendDestroyPropertyInfoInternal(zv *types.Zval) {
	var property_info *ZendPropertyInfo = zv.GetPtr()
	types.ZendStringRelease(property_info.GetName())
	Free(property_info)
}
func ZendBuildRuntimeDefinitionKey(name *types.ZendString, start_lineno uint32) *types.ZendString {
	var filename *types.ZendString = CG__().GetActiveOpArray().GetFilename()
	var result *types.ZendString = ZendStrpprintf(0, "%c%s%s:%"+"u"+"$%"+PRIx32, '0', name.GetVal(), filename.GetVal(), start_lineno, b.PostInc(&(CG__().GetRtdKeyCounter())))
	return types.ZendNewInternedString(result)
}
func ZendGetUnqualifiedName(name *types.ZendString, result **byte, result_len *int) types.ZendBool {
	var ns_separator *byte = ZendMemrchr(name.GetVal(), '\\', name.GetLen())
	if ns_separator != nil {
		*result = ns_separator + 1
		*result_len = name.GetVal() + name.GetLen() - (*result)
		return 1
	}
	return 0
}
func ZendIsReservedClassName(name *types.ZendString) types.ZendBool {
	var reserved *ReservedClassName = ReservedClassNames
	var uqname *byte = name.GetVal()
	var uqname_len int = name.GetLen()
	ZendGetUnqualifiedName(name, &uqname, &uqname_len)
	for ; reserved.GetName() != nil; reserved++ {
		if uqname_len == reserved.GetLen() && ZendBinaryStrcasecmp(uqname, uqname_len, reserved.GetName(), reserved.GetLen()) == 0 {
			return 1
		}
	}
	return 0
}
func ZendAssertValidClassName(name *types.ZendString) {
	if ZendIsReservedClassName(name) != 0 {
		ZendErrorNoreturn(E_COMPILE_ERROR, "Cannot use '%s' as class name as it is reserved", name.GetVal())
	}
}
func ZendLookupBuiltinTypeByName(name *types.ZendString) types.ZendUchar {
	var info *BuiltinTypeInfo = &BuiltinTypes[0]
	for ; info.GetName() != nil; info++ {
		if name.GetLen() == info.GetNameLen() && ZendBinaryStrcasecmp(name.GetVal(), name.GetLen(), info.GetName(), info.GetNameLen()) == 0 {
			return info.GetType()
		}
	}
	return 0
}
func ZendOparrayContextBegin(prev_context *ZendOparrayContext) {
	*prev_context = CG__().GetContext()
	CG__().GetContext().SetOpcodesSize(INITIAL_OP_ARRAY_SIZE)
	CG__().GetContext().SetVarsSize(0)
	CG__().GetContext().SetLiteralsSize(0)
	CG__().GetContext().SetFastCallVar(-1)
	CG__().GetContext().SetTryCatchOffset(-1)
	CG__().GetContext().SetCurrentBrkCont(-1)
	CG__().GetContext().SetLastBrkCont(0)
	CG__().GetContext().SetBrkContArray(nil)
	CG__().GetContext().SetLabels(nil)
}
func ZendOparrayContextEnd(prev_context *ZendOparrayContext) {
	if CG__().GetContext().GetBrkContArray() != nil {
		Efree(CG__().GetContext().GetBrkContArray())
		CG__().GetContext().SetBrkContArray(nil)
	}
	if CG__().GetContext().GetLabels() != nil {
		CG__().GetContext().GetLabels().Destroy()
		FREE_HASHTABLE(CG__().GetContext().GetLabels())
		CG__().GetContext().SetLabels(nil)
	}
	CG__().SetContext(*prev_context)
}
func ZendResetImportTables() {
	if FC__().GetImports() != nil {
		FC__().GetImports().Destroy()
		Efree(FC__().GetImports())
		FC__().SetImports(nil)
	}
	if FC__().GetImportsFunction() != nil {
		FC__().GetImportsFunction().Destroy()
		Efree(FC__().GetImportsFunction())
		FC__().SetImportsFunction(nil)
	}
	if FC__().GetImportsConst() != nil {
		FC__().GetImportsConst().Destroy()
		Efree(FC__().GetImportsConst())
		FC__().SetImportsConst(nil)
	}
}
func ZendEndNamespace() {
	FC__().SetInNamespace(0)
	ZendResetImportTables()
	if FC__().GetCurrentNamespace() != nil {
		types.ZendStringReleaseEx(FC__().GetCurrentNamespace(), 0)
		FC__().SetCurrentNamespace(nil)
	}
}
func ZendFileContextBegin(prev_context *ZendFileContext) {
	*prev_context = CG__().GetFileContext()
	FC__().SetImports(nil)
	FC__().SetImportsFunction(nil)
	FC__().SetImportsConst(nil)
	FC__().SetCurrentNamespace(nil)
	FC__().SetInNamespace(0)
	FC__().SetHasBracketedNamespaces(0)
	FC__().GetDeclarables().SetTicks(0)
	ZendHashInit(&(FC__().GetSeenSymbols()), 8, nil, nil, 0)
}
func ZendFileContextEnd(prev_context *ZendFileContext) {
	ZendEndNamespace()
	FC__().GetSeenSymbols().Destroy()
	CG__().SetFileContext(*prev_context)
}
func ZendInitCompilerDataStructures() {
	CG__().GetLoopVarStack().Init()
	CG__().GetDelayedOplinesStack().Init()
	CG__().SetActiveClassEntry(nil)
	CG__().SetInCompilation(0)
	CG__().SetSkipShebang(0)
	CG__().SetMemoizedExprs(nil)
	CG__().SetMemoizeMode(0)
}
func ZendRegisterSeenSymbol(name *types.ZendString, kind uint32) {
	var zv *types.Zval = FC__().GetSeenSymbols().KeyFind(name.GetStr())
	if zv != nil {
		zv.SetLval(zv.GetLval() | kind)
	} else {
		var tmp types.Zval
		tmp.SetLong(kind)
		FC__().GetSeenSymbols().KeyAddNew(name.GetStr(), &tmp)
	}
}
func ZendHaveSeenSymbol(name *types.ZendString, kind uint32) types.ZendBool {
	var zv *types.Zval = FC__().GetSeenSymbols().KeyFind(name.GetStr())
	return zv != nil && (zv.GetLval()&kind) != 0
}
func FileHandleDtor(fh *ZendFileHandle) { fh.Destroy() }
func InitCompiler() {
	CG__().SetArena(ZendArenaCreate(64 * 1024))
	CG__().SetActiveOpArray(nil)
	memset(CG__().GetContext(), 0, b.SizeOf("CG ( context )"))
	ZendInitCompilerDataStructures()
	ZendInitRsrcList()
	CG__().filenamesTable = make(map[string]string)
	CG__().GetOpenFiles().Init(b.SizeOf("zend_file_handle"), (func(any))(FileHandleDtor), 0)
	CG__().SetUncleanShutdown(0)
	CG__().SetDelayedVarianceObligations(nil)
	CG__().SetDelayedAutoloads(nil)
}
func ShutdownCompiler() {
	CG__().GetLoopVarStack().Destroy()
	CG__().GetDelayedOplinesStack().Destroy()
	CG__().filenamesTable = nil
	ZendArenaDestroy(CG__().GetArena())
	if CG__().GetDelayedVarianceObligations() != nil {
		CG__().GetDelayedVarianceObligations().Destroy()
		FREE_HASHTABLE(CG__().GetDelayedVarianceObligations())
		CG__().SetDelayedVarianceObligations(nil)
	}
	if CG__().GetDelayedAutoloads() != nil {
		CG__().GetDelayedAutoloads().Destroy()
		FREE_HASHTABLE(CG__().GetDelayedAutoloads())
		CG__().SetDelayedAutoloads(nil)
	}
}
func ZendSetCompiledFilename(new_compiled_filename string) {
	if _, ok := CG__().filenamesTable[new_compiled_filename]; !ok {
		CG__().filenamesTable[new_compiled_filename] = new_compiled_filename
	}
}
func ZendRestoreCompiledFilename(original_compiled_filename *types.ZendString) {
	CG__().SetCompiledFilename(original_compiled_filename)
}
func ZendGetCompiledFilename() *types.ZendString { return CG__().GetCompiledFilename() }
func ZendGetCompiledLineno() int                 { return CG__().GetZendLineno() }
func ZendIsCompiling() types.ZendBool            { return CG__().GetInCompilation() }
func GetTemporaryVariable() uint32 {
	return b.PostInc(&(CG__().GetActiveOpArray().GetT()))
}
func LookupCv(name *types.ZendString) int {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var i int = 0
	var hash_value ZendUlong = name.GetHash()
	for i < op_array.GetLastVar() {
		if op_array.GetVars()[i].GetH() == hash_value && types.ZendStringEquals(op_array.GetVars()[i], name) != 0 {
			return int(types.ZendIntptrT(nil.VarNum(i)))
		}
		i++
	}
	i = op_array.GetLastVar()
	op_array.GetLastVar()++
	if op_array.GetLastVar() > CG__().GetContext().GetVarsSize() {
		CG__().GetContext().SetVarsSize(CG__().GetContext().GetVarsSize() + 16)
		op_array.SetVars(Erealloc(op_array.GetVars(), CG__().GetContext().GetVarsSize()*b.SizeOf("zend_string *")))
	}
	op_array.GetVars()[i] = name.Copy()
	return int(types.ZendIntptrT(nil.VarNum(i)))
}
func ZendDelLiteral(op_array *ZendOpArray, n int) {
	ZvalPtrDtorNogc(CT_CONSTANT_EX(op_array, n))
	if n+1 == op_array.GetLastLiteral() {
		op_array.GetLastLiteral()--
	} else {
		CT_CONSTANT_EX(op_array, n).SetUndef()
	}
}
func ZendInsertLiteral(op_array *ZendOpArray, zv *types.Zval, literal_position int) {
	var lit *types.Zval = CT_CONSTANT_EX(op_array, literal_position)
	if zv.IsString() {
		ZvalMakeInternedString(zv)
	}
	types.ZVAL_COPY_VALUE(lit, zv)
	lit.SetU2Extra(0)
}
func ZendAddLiteral(zv *types.Zval) int {
	var op_array *ZendOpArray = CG__().GetActiveOpArray()
	var i int = op_array.GetLastLiteral()
	op_array.GetLastLiteral()++
	if i >= CG__().GetContext().GetLiteralsSize() {
		for i >= CG__().GetContext().GetLiteralsSize() {
			CG__().GetContext().SetLiteralsSize(CG__().GetContext().GetLiteralsSize() + 16)
		}
		op_array.SetLiterals((*types.Zval)(Erealloc(op_array.GetLiterals(), CG__().GetContext().GetLiteralsSize()*b.SizeOf("zval"))))
	}
	ZendInsertLiteral(op_array, zv, i)
	return i
}
func ZendAddLiteralString(str **types.ZendString) int {
	var ret int
	var zv types.Zval
	zv.SetString(*str)
	ret = ZendAddLiteral(&zv)
	*str = zv.GetStr()
	return ret
}
func ZendAddFuncNameLiteral(name *types.ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *types.ZendString = ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)
	return ret
}
func ZendAddNsFuncNameLiteral(name *types.ZendString) int {
	var unqualified_name *byte
	var unqualified_name_len int

	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *types.ZendString = ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)

	/* Lowercased unqualfied name */

	if ZendGetUnqualifiedName(name, &unqualified_name, &unqualified_name_len) != 0 {
		lc_name = types.ZendStringAlloc(unqualified_name_len, 0)
		ZendStrTolowerCopy(lc_name.GetVal(), unqualified_name, unqualified_name_len)
		ZendAddLiteralString(&lc_name)
	}
	return ret
}
func ZendAddClassNameLiteral(name *types.ZendString) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *types.ZendString = ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)
	return ret
}
func ZendAddConstNameLiteral(name *types.ZendString, unqualified types.ZendBool) int {
	var tmp_name *types.ZendString
	var ret int = ZendAddLiteralString(&name)
	var ns_len int = 0
	var after_ns_len int = name.GetLen()
	var after_ns *byte = ZendMemrchr(name.GetVal(), '\\', name.GetLen())
	if after_ns != nil {
		after_ns += 1
		ns_len = after_ns - name.GetVal() - 1
		after_ns_len = name.GetLen() - ns_len - 1

		/* lowercased namespace name & original constant name */

		tmp_name = types.ZendStringInit(name.GetVal(), name.GetLen(), 0)
		ZendStrTolower(tmp_name.GetVal(), ns_len)
		ZendAddLiteralString(&tmp_name)

		/* lowercased namespace name & lowercased constant name */

		tmp_name = ZendStringTolower(name)
		ZendAddLiteralString(&tmp_name)
		if unqualified == 0 {
			return ret
		}
	} else {
		after_ns = name.GetVal()
	}

	/* original unqualified constant name */

	tmp_name = types.ZendStringInit(after_ns, after_ns_len, 0)
	ZendAddLiteralString(&tmp_name)

	/* lowercased unqualified constant name */

	tmp_name = types.ZendStringAlloc(after_ns_len, 0)
	ZendStrTolowerCopy(tmp_name.GetVal(), after_ns, after_ns_len)
	ZendAddLiteralString(&tmp_name)
	return ret
}
func LITERAL_STR(op ZnodeOp, str *types.ZendString) {
	var _c types.Zval
	_c.SetString(str)
	op.SetConstant(ZendAddLiteral(&_c))
}
func ZendStopLexing() {
	if INI_SCNG__().on_event {
		INI_SCNG__().on_event(ON_STOP, END, 0, INI_SCNG__().on_event_context)
	}
	INI_SCNG__().SetYyCursor(INI_SCNG__().GetYyLimit())
}
func ZendBeginLoop(free_opcode types.ZendUchar, loop_var *Znode, is_switch types.ZendBool) {
	var brk_cont_element *ZendBrkContElement
	var parent int = CG__().GetContext().GetCurrentBrkCont()
	var info ZendLoopVar = MakeZendLoopVar(0)
	CG__().GetContext().SetCurrentBrkCont(CG__().GetContext().GetLastBrkCont())
	brk_cont_element = GetNextBrkContElement()
	brk_cont_element.SetParent(parent)
	brk_cont_element.SetIsSwitch(is_switch)
	if loop_var != nil && (loop_var.GetOpType()&(IS_VAR|IS_TMP_VAR)) != 0 {
		var start uint32 = GetNextOpNumber()
		info.SetOpcode(free_opcode)
		info.SetVarType(loop_var.GetOpType())
		info.SetVarNum(loop_var.GetOp().GetVar())
		brk_cont_element.SetStart(start)
	} else {
		info.SetOpcode(ZEND_NOP)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

		brk_cont_element.SetStart(-1)

		/* The start field is used to free temporary variables in case of exceptions.
		 * We won't try to free something of we don't have loop variable.  */

	}
	CG__().GetLoopVarStack().Push(&info)
}
func ZendEndLoop(cont_addr int, var_node *Znode) {
	var end uint32 = GetNextOpNumber()
	var brk_cont_element *ZendBrkContElement = CG__().GetContext().GetBrkContArray()[CG__().GetContext().GetCurrentBrkCont()]
	brk_cont_element.SetCont(cont_addr)
	brk_cont_element.SetBrk(end)
	CG__().GetContext().SetCurrentBrkCont(brk_cont_element.GetParent())
	CG__().GetLoopVarStack().DelTop()
}
func ZendDoFree(op1 *Znode) {
	if op1.GetOpType() == IS_TMP_VAR {
		var opline *ZendOp = CG__().GetActiveOpArray().GetOpcodes()[CG__().GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == ZEND_END_SILENCE {
			opline--
		}
		if opline.GetResultType() == IS_TMP_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == ZEND_BOOL || opline.GetOpcode() == ZEND_BOOL_NOT {
				return
			}
		}
		ZendEmitOp(nil, ZEND_FREE, op1, nil)
	} else if op1.GetOpType() == IS_VAR {
		var opline *ZendOp = CG__().GetActiveOpArray().GetOpcodes()[CG__().GetActiveOpArray().GetLast()-1]
		for opline.GetOpcode() == ZEND_END_SILENCE || opline.GetOpcode() == ZEND_EXT_FCALL_END || opline.GetOpcode() == ZEND_OP_DATA {
			opline--
		}
		if opline.GetResultType() == IS_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
			if opline.GetOpcode() == ZEND_FETCH_THIS {
				opline.SetOpcode(ZEND_NOP)
				opline.SetResultType(IS_UNUSED)
			} else {
				opline.SetResultType(IS_UNUSED)
			}
		} else {
			for opline >= CG__().GetActiveOpArray().GetOpcodes() {
				if (opline.GetOpcode() == ZEND_FETCH_LIST_R || opline.GetOpcode() == ZEND_FETCH_LIST_W) && opline.GetOp1Type() == IS_VAR && opline.GetOp1().GetVar() == op1.GetOp().GetVar() {
					ZendEmitOp(nil, ZEND_FREE, op1, nil)
					return
				}
				if opline.GetResultType() == IS_VAR && opline.GetResult().GetVar() == op1.GetOp().GetVar() {
					if opline.GetOpcode() == ZEND_NEW {
						ZendEmitOp(nil, ZEND_FREE, op1, nil)
					}
					break
				}
				opline--
			}
		}
	} else if op1.GetOpType() == IS_CONST {

		/* Destroy value without using GC: When opcache moves arrays into SHM it will
		 * free the zend_array structure, so references to it from outside the op array
		 * become invalid. GC would cause such a reference in the root buffer. */

		ZvalPtrDtorNogc(op1.GetConstant())

		/* Destroy value without using GC: When opcache moves arrays into SHM it will
		 * free the zend_array structure, so references to it from outside the op array
		 * become invalid. GC would cause such a reference in the root buffer. */

	}
}
func ZendAddClassModifier(flags uint32, new_flag uint32) uint32 {
	var new_flags uint32 = flags | new_flag
	if (flags&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 && (new_flag&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_FINAL) != 0 && (new_flag&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) != 0 && (new_flags&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class", 0)
		return 0
	}
	return new_flags
}
func ZendAddMemberModifier(flags uint32, new_flag uint32) uint32 {
	var new_flags uint32 = flags | new_flag
	if (flags&ZEND_ACC_PPP_MASK) != 0 && (new_flag&ZEND_ACC_PPP_MASK) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple access type modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_ABSTRACT) != 0 && (new_flag&ZEND_ACC_ABSTRACT) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple abstract modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_STATIC) != 0 && (new_flag&ZEND_ACC_STATIC) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple static modifiers are not allowed", 0)
		return 0
	}
	if (flags&ZEND_ACC_FINAL) != 0 && (new_flag&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Multiple final modifiers are not allowed", 0)
		return 0
	}
	if (new_flags&ZEND_ACC_ABSTRACT) != 0 && (new_flags&ZEND_ACC_FINAL) != 0 {
		ZendThrowException(ZendCeCompileError, "Cannot use the final modifier on an abstract class member", 0)
		return 0
	}
	return new_flags
}
func ZendConcat3(
	str1 *byte,
	str1_len int,
	str2 string,
	str2_len int,
	str3 *byte,
	str3_len int,
) *types.ZendString {
	var len_ int = str1_len + str2_len + str3_len
	var res *types.ZendString = types.ZendStringAlloc(len_, 0)
	memcpy(res.GetVal(), str1, str1_len)
	memcpy(res.GetVal()+str1_len, str2, str2_len)
	memcpy(res.GetVal()+str1_len+str2_len, str3, str3_len)
	res.GetVal()[len_] = '0'
	return res
}
func ZendConcatNames(name1 *byte, name1_len int, name2 *byte, name2_len int) *types.ZendString {
	return ZendConcat3(name1, name1_len, "\\", 1, name2, name2_len)
}
func ZendPrefixWithNs(name *types.ZendString) *types.ZendString {
	if FC__().GetCurrentNamespace() != nil {
		var ns *types.ZendString = FC__().GetCurrentNamespace()
		return ZendConcatNames(ns.GetVal(), ns.GetLen(), name.GetVal(), name.GetLen())
	} else {
		return name.Copy()
	}
}
func ZendHashFindPtrLc(ht *types.HashTable, str *byte, len_ int) any {
	var result any
	var lcname *types.ZendString
	types.ZSTR_ALLOCA_ALLOC(lcname, len_)
	ZendStrTolowerCopy(lcname.GetVal(), str, len_)
	result = ZendHashFindPtr(ht, lcname)
	lcname.Free()
	return result
}
func ZendResolveNonClassName(name *types.ZendString, type_ uint32, is_fully_qualified *types.ZendBool, case_sensitive types.ZendBool, current_import_sub *types.HashTable) *types.ZendString {
	var compound *byte
	*is_fully_qualified = 0
	if name.GetVal()[0] == '\\' {

		/* Remove \ prefix (only relevant if this is a string rather than a label) */

		*is_fully_qualified = 1
		return types.ZendStringInit(name.GetVal()+1, name.GetLen()-1, 0)
	}
	if type_ == ZEND_NAME_FQ {
		*is_fully_qualified = 1
		return name.Copy()
	}
	if type_ == ZEND_NAME_RELATIVE {
		*is_fully_qualified = 1
		return ZendPrefixWithNs(name)
	}
	if current_import_sub != nil {

		/* If an unqualified name is a function/const alias, replace it. */

		var import_name *types.ZendString
		if case_sensitive != 0 {
			import_name = ZendHashFindPtr(current_import_sub, name)
		} else {
			import_name = ZendHashFindPtrLc(current_import_sub, name.GetVal(), name.GetLen())
		}
		if import_name != nil {
			*is_fully_qualified = 1
			return import_name.Copy()
		}
	}
	compound = memchr(name.GetVal(), '\\', name.GetLen())
	if compound != nil {
		*is_fully_qualified = 1
	}
	if compound != nil && FC__().GetImports() != nil {

		/* If the first part of a qualified name is an alias, substitute it. */

		var len_ int = compound - name.GetVal()
		var import_name *types.ZendString = ZendHashFindPtrLc(FC__().GetImports(), name.GetVal(), len_)
		if import_name != nil {
			return ZendConcatNames(import_name.GetVal(), import_name.GetLen(), name.GetVal()+len_+1, name.GetLen()-len_-1)
		}
	}
	return ZendPrefixWithNs(name)
}
