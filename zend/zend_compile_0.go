package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
	"unsafe"
)

func OBJ_PROP(obj *types.Object, offset uint32) *types.Zval {
	return (*types.Zval)((*byte)(obj + offset))
}
func OBJ_PROP_TO_OFFSET(num int) int {
	var tmp *types.Object = nil

	return uintptr(unsafe.Pointer(tmp.GetPropertiesTable())) - uintptr(unsafe.Pointer(tmp)) + b.SizeOf("zval")*num
}
func OBJ_PROP_TO_NUM(offset uint32) int {
	return (offset - OBJ_PROP_TO_OFFSET(0)) / b.SizeOf("zval")
}
func ZEND_FN_SCOPE_NAME(function types.IFunction) string {
	if function != nil && function.GetScope() != nil {
		return function.GetScope().Name()
	} else {
		return ""
	}
}
func ZEND_CALL_INFO(call *ZendExecuteData) uint32         { return call.CallInfo() }
func ZEND_ADD_CALL_FLAG_EX(call_info uint32, flag uint32) { call_info |= flag }
func ZEND_DEL_CALL_FLAG_EX(call_info uint32, flag uint32) { call_info &= ^flag }
func ZEND_ADD_CALL_FLAG(call *ZendExecuteData, flag uint32) {
	ZEND_ADD_CALL_FLAG_EX(call.CallInfo(), flag)
}
func ZEND_DEL_CALL_FLAG(call *ZendExecuteData, flag uint32) {
	ZEND_DEL_CALL_FLAG_EX(call.CallInfo(), flag)
}
func ZEND_CALL_VAR(call *ZendExecuteData, n uint32) *types.Zval {
	return (*types.Zval)((*byte)(call) + int(n))
}
func EX_CALL_INFO() uint32 { return ZEND_CALL_INFO(executeData) }
func EX_VAR(executeData *ZendExecuteData, n uint32) *types.Zval {
	return ZEND_CALL_VAR(executeData, n)
}
func EX_VAR_TO_NUM(n uint32) __auto__ {
	return uint32(ZEND_CALL_VAR(nil, n) - nil.VarNum(0))
}
func ZEND_OPLINE_NUM_TO_OFFSET(op_array *types.ZendOpArray, opline *types.ZendOp, opline_num uint32) *byte {
	return (*byte)(op_array.GetOpcodes()[opline_num] - (*byte)(opline))
}
func ZEND_OFFSET_TO_OPLINE(base *types.ZendOp, offset uint32) *types.ZendOp {
	return (*types.ZendOp)((*byte)(base) + int(offset))
}
func OP_JMP_ADDR(opline *types.ZendOp, node types.ZnodeOp) *types.ZendOp {
	return ZEND_OFFSET_TO_OPLINE(opline, node.GetJmpOffset())
}
func ZEND_PASS_TWO_UPDATE_JMP_TARGET(op_array *types.ZendOpArray, opline *types.ZendOp, node types.ZnodeOp) {
	node.SetJmpOffset(ZEND_OPLINE_NUM_TO_OFFSET(op_array, opline, node.GetOplineNum()))
}
func CT_CONSTANT_EX(op_array *types.ZendOpArray, num uint32) *types.Zval {
	return op_array.GetLiteral(num)
}
func CT_CONSTANT(node types.ZnodeOp) *types.Zval {
	return CT_CONSTANT_EX(CG__().GetActiveOpArray(), node.GetConstant())
}
func RT_CONSTANT(opline *types.ZendOp, node types.ZnodeOp) *types.Zval {
	return (*types.Zval)((*byte)(opline) + int32(node.GetConstant()))
}
func ZEND_PASS_TWO_UPDATE_CONSTANT(op_array *types.ZendOpArray, opline *types.ZendOp, node types.ZnodeOp) {
	node.SetConstant((*byte)(CT_CONSTANT_EX(op_array, node.GetConstant())) - (*byte)(opline))
}
func RUN_TIME_CACHE(op_array *types.ZendOpArray) []any {
	return op_array.GetRunTimeCache()
}
func ZendGetUnmangledPropertyNameEx(mangledProp string) string {
	_, propName, _ := ZendUnmanglePropertyName_Ex(mangledProp)
	return propName
}
func ZEND_USER_CODE(type_ uint8) bool { return (type_ & 1) == 0 }
func ZendCheckArgSendType(zf types.IFunction, arg_num uint32, mask uint32) int {
	return types.IntBool(zf.CheckArgSendType(arg_num, uint8(mask)))
}
func ARG_MUST_BE_SENT_BY_REF(zf types.IFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF)
}
func ARG_SHOULD_BE_SENT_BY_REF(zf types.IFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func ARG_MAY_BE_SENT_BY_REF(zf types.IFunction, arg_num uint32) int {
	return ZendCheckArgSendType(zf, arg_num, ZEND_SEND_PREFER_REF)
}

func ZEND_CHECK_ARG_FLAG(zf types.IFunction, arg_num uint32, mask uint8) int {
	return types.IntBool(zf.CheckArgSendType(arg_num, mask))
}
func QUICK_ARG_MUST_BE_SENT_BY_REF(zf types.IFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF)
}
func QUICK_ARG_SHOULD_BE_SENT_BY_REF(zf types.IFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_BY_REF|ZEND_SEND_PREFER_REF)
}
func QUICK_ARG_MAY_BE_SENT_BY_REF(zf types.IFunction, arg_num int) int {
	return ZEND_CHECK_ARG_FLAG(zf, arg_num, ZEND_SEND_PREFER_REF)
}
func ZEND_IS_INCREMENT(opcode uint8) bool { return (opcode & 1) == 0 }
func ZendAllocCacheSlots(count unsigned) uint32 {
	var op_array *types.ZendOpArray = CG__().GetActiveOpArray()
	var ret uint32 = op_array.GetCacheSize()
	op_array.SetCacheSize(op_array.GetCacheSize() + count*b.SizeOf("void *"))
	return ret
}
func ZendAllocCacheSlot() uint32 { return ZendAllocCacheSlots(1) }

func GetNextOpNumber() uint32 {
	return CG__().GetActiveOpArray().GetLast()
}
func GetNextOp() *types.ZendOp {
	var op_array *types.ZendOpArray = CG__().GetActiveOpArray()
	var next_op_num uint32 = lang.PostInc(&(op_array.GetLast()))
	var next_op *types.ZendOp
	if next_op_num >= CG__().GetContext().GetOpcodesSize() {
		CG__().GetContext().SetOpcodesSize(CG__().GetContext().GetOpcodesSize() * 4)
		op_array.SetOpcodes(Erealloc(op_array.GetOpcodes(), CG__().GetContext().GetOpcodesSize()*b.SizeOf("zend_op")))
	}
	next_op = &op_array.GetOpcodes()[next_op_num]
	return CurrCompiler().initOp(next_op)
}
func GetNextBrkContElement() *ZendBrkContElement {
	return CG__().GetContext().AddBrkCont()
}
func ZendBuildRuntimeDefinitionKey(name *types.String, start_lineno uint32) *types.String {
	var filename = CG__().GetActiveOpArray().GetFilename()
	var result = ZendSprintf("%c%s%s:%u$%d", '\000', name.GetStr(), filename, start_lineno, lang.PostInc(&(CG__().GetRtdKeyCounter())))
	return types.NewString(result)
}

func ZendGetUnqualifiedNameEx(name string) (string, bool) {
	if pos := strings.LastIndexByte(name, '\\'); pos >= 0 {
		return name[pos+1:], true
	}
	return name, false
}
func ZendIsReservedClassName(name string) bool {
	name, _ = ZendGetUnqualifiedNameEx(name)
	lcName := ascii.StrToLower(name)
	_, reserved := reservedClassNames[lcName]
	return reserved
}
func ZendAssertValidClassName(name string) {
	if ZendIsReservedClassName(name) {
		faults.ErrorNoreturn(faults.E_COMPILE_ERROR, "Cannot use '%s' as class name as it is reserved", name)
	}
}

func ZendLookupBuiltinTypeByName(name string) uint8 {
	lcName := ascii.StrToLower(name)
	if typ, ok := builtinTypes[lcName]; ok {
		return typ
	}
	return 0
}
func ZendOparrayContextBegin(prev_context *ZendOparrayContext) {
	*prev_context = *CG__().GetContext()
	*CG__().GetContext() = *NewOpArrayContext()
}
func ZendOparrayContextEnd(prev_context *ZendOparrayContext) {
	//CG__().GetContext().End()
	CG__().SetContext(*prev_context)
}
func ZendResetImportTables() {
	FC__().ResetImportTables()
}
func ZendEndNamespace() {
	FC__().SetInNamespace(0)
	ZendResetImportTables()
	if FC__().GetCurrentNamespace() != nil {
		// types.ZendStringReleaseEx(FC__().GetCurrentNamespace(), 0)
		FC__().SetCurrentNamespace(nil)
	}
}
func ZendFileContextBegin(prevContext *ZendFileContext) {
	*prevContext = *CG__().GetFileContext()
	FC__().SetImports(nil)
	FC__().SetImportsFunction(nil)
	FC__().SetImportsConst(nil)
	FC__().SetCurrentNamespace(nil)
	FC__().SetInNamespace(0)
	FC__().SetHasBracketedNamespaces(0)
	FC__().InitSeenSymbols()
}
func ZendFileContextEnd(prevContext *ZendFileContext) {
	ZendEndNamespace()
	FC__().GetSeenSymbols().Destroy()
	CG__().SetFileContext(*prevContext)
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
func ZendRegisterSeenSymbol(name *types.String, kind uint32) {
	var zv *types.Zval = FC__().GetSeenSymbols().KeyFind(name.GetStr())
	if zv != nil {
		zv.SetLong(zv.Long() | kind)
	} else {
		var tmp types.Zval
		tmp.SetLong(kind)
		FC__().GetSeenSymbols().KeyAddNew(name.GetStr(), &tmp)
	}
}
func ZendHaveSeenSymbol(name string, kind uint32) bool {
	var zv *types.Zval = FC__().GetSeenSymbols().KeyFind(name)
	return zv != nil && (zv.Long()&kind) != 0
}
func FileHandleDtor(fh *FileHandle) { fh.Destroy() }
func InitCompiler() {
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
	if CG__().GetDelayedVarianceObligations() != nil {
		CG__().GetDelayedVarianceObligations().Destroy()
		CG__().SetDelayedVarianceObligations(nil)
	}
	if CG__().GetDelayedAutoloads() != nil {
		CG__().GetDelayedAutoloads().Destroy()
		CG__().SetDelayedAutoloads(nil)
	}
}
func ZendSetCompiledFilename(new_compiled_filename string) {
	if _, ok := CG__().filenamesTable[new_compiled_filename]; !ok {
		CG__().filenamesTable[new_compiled_filename] = new_compiled_filename
	}
}
func ZendRestoreCompiledFilename(original_compiled_filename string) {
	CG__().SetCompiledFilename(original_compiled_filename)
}
func ZendGetCompiledFilename() string { return CG__().GetCompiledFilename() }
func ZendGetCompiledLineno() int      { return CG__().GetZendLineno() }
func ZendIsCompiling() bool           { return CG__().GetInCompilation() }
func GetTemporaryVariable() uint32 {
	return lang.PostInc(&(CG__().GetActiveOpArray().GetT()))
}
func LookupCv(name string) int {
	var opArray *types.ZendOpArray = CG__().GetActiveOpArray()

	i := opArray.FindOrAddVarName(name)

	var ex *ZendExecuteData = nil
	return int(types.ZendIntptrT(ex.VarNum(i)))
}
func ZendAddLiteral(zv *types.Zval) int {
	var opArray = CG__().GetActiveOpArray()
	return opArray.AddLiteral(zv)
}
func ZendAddLiteralStringEx(str string) int {
	zv := types.NewZvalString(str)
	return ZendAddLiteral(zv)
}
func ZendAddLiteralString(str **types.String) int {
	zv := types.NewZvalString((*str).GetStr())
	var ret int
	ret = ZendAddLiteral(zv)
	*str = zv.StringEx()
	return ret
}
func ZendAddFuncNameLiteral(name *types.String) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *types.String = operators.ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)
	return ret
}
func ZendAddNsFuncNameLiteral(name *types.String) int {
	/* Original name */
	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */
	lcName := types.NewString(ascii.StrToLower(name.GetStr()))
	ZendAddLiteralString(&lcName)

	/* Lowercased unqualfied name */
	if unqualifiedName, ok := ZendGetUnqualifiedNameEx(name.GetStr()); ok {
		uqLcName := types.NewString(ascii.StrToLower(unqualifiedName))
		ZendAddLiteralString(&uqLcName)
	}
	return ret
}
func ZendAddClassNameLiteral(name *types.String) int {
	/* Original name */

	var ret int = ZendAddLiteralString(&name)

	/* Lowercased name */

	var lc_name *types.String = operators.ZendStringTolower(name)
	ZendAddLiteralString(&lc_name)
	return ret
}
func ZendAddConstNameLiteral(name string, unqualified bool) int {
	var ret int = ZendAddLiteralStringEx(name)

	var afterNs string
	if pos := strings.IndexByte(name, '\\'); pos >= 0 {
		ns := name[:pos]
		afterNs = name[pos+1:]

		/* lowercased namespace name & original constant name */
		ZendAddLiteralStringEx(ascii.StrToLower(ns))

		/* lowercased namespace name & lowercased constant name */
		ZendAddLiteralStringEx(ascii.StrToLower(name))
		if unqualified == 0 {
			return ret
		}
	} else {
		afterNs = name
	}

	/* original unqualified constant name */
	ZendAddLiteralStringEx(afterNs)

	/* lowercased unqualified constant name */
	ZendAddLiteralStringEx(ascii.StrToLower(afterNs))

	return ret
}
func LITERAL_STR(op types.ZnodeOp, str *types.String) {
	var _c types.Zval
	_c.SetStringEx(str)
	op.SetConstant(ZendAddLiteral(&_c))
}
func ZendBeginLoop(free_opcode uint8, loop_var *Znode, is_switch bool) {
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
		var opline *types.ZendOp = CG__().GetActiveOpArray().GetOpcodes()[CG__().GetActiveOpArray().GetLast()-1]
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
		var opline *types.ZendOp = CG__().GetActiveOpArray().GetOpcodes()[CG__().GetActiveOpArray().GetLast()-1]
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

		// ZvalPtrDtorNogc(op1.GetConstant())

		/* Destroy value without using GC: When opcache moves arrays into SHM it will
		 * free the zend_array structure, so references to it from outside the op array
		 * become invalid. GC would cause such a reference in the root buffer. */

	}
}
func ZendConcatNames(name1 string, name2 string) string {
	return name1 + "\\" + name2
}
func ZendPrefixWithNs(name *types.String) *types.String {
	str := ZendPrefixWithNsEx(name.GetStr())
	return types.NewString(str)
}
func ZendPrefixWithNsEx(name string) string {
	if FC__().GetCurrentNamespace() != nil {
		var ns = FC__().GetCurrentNamespace()
		return ZendConcatNames(ns.GetStr(), name)
	} else {
		return name
	}
}
func ZendHashFindPtrLc(ht *types.Array, str *byte, len_ int) any {
	name := b.CastStr(str, len_)
	lcName := ascii.StrToLower(name)
	return types.ZendHashFindPtr(ht, lcName)
}
