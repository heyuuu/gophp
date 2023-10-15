package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/kits/slicekit"
	"github.com/heyuuu/gophp/php/contracts"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend/faults"
)

var CompilerGlobals ZendCompilerGlobals

func CG__() *ZendCompilerGlobals { return &CompilerGlobals }

// ZendCompilerGlobals
type ZendCompilerGlobals struct {
	loopVarStack       []*ZendLoopVar
	active_class_entry *types.ClassEntry
	compiled_filename  string
	zend_lineno        int
	active_op_array    *types.ZendOpArray
	functionTable      FunctionTable
	classTable         ClassTable
	filenamesTable     map[string]string //filenames_table              HashTable
	autoGlobals        map[string]*ZendAutoGlobal

	parse_error                  bool
	in_compilation               bool
	short_tags                   bool
	unclean_shutdown             bool
	ini_parser_unbuffered_errors bool
	openFiles                    ZendLlist[*FileHandle]
	ini_parser_param             *ZendIniParserParam
	skip_shebang                 bool
	increment_lineno             bool
	doc_comment                  *types.String
	extra_fn_flags               uint32
	compiler_options             uint32
	context                      ZendOparrayContext
	file_context                 ZendFileContext
	ast                          *ZendAst
	delayedOplinesStack          []*types.ZendOp
	memoized_exprs               *types.Array
	memoize_mode                 int
	map_ptr_base                 any
	map_ptr_size                 int
	map_ptr_last                 int
	delayed_variance_obligations *types.Array
	delayed_autoloads            *types.Array
	rtd_key_counter              uint32
}

// life cycle
var _ contracts.ModuleLifeCycle = (*ZendCompilerGlobals)(nil)

func (cg *ZendCompilerGlobals) StartUp() {
	cg.functionTable = types.NewLcTable[types.IFunction]()
	cg.classTable = types.NewLcTable[*types.ClassEntry]()
	cg.autoGlobals = make(map[string]*ZendAutoGlobal)
}

func (cg *ZendCompilerGlobals) Shutdown() {
	cg.functionTable.Destroy()
	cg.classTable.Destroy()
	cg.autoGlobals = nil
}

func (cg *ZendCompilerGlobals) Activate() {
	cg.active_op_array = nil
	cg.context = ZendOparrayContext{}

	// ZendInitCompilerDataStructures
	cg.loopVarStack = nil
	cg.delayedOplinesStack = nil
	cg.active_class_entry = nil
	cg.in_compilation = false
	cg.skip_shebang = false
	cg.memoized_exprs = nil
	cg.memoize_mode = 0

	// ZendInitRsrcList
	EG__().InitRegularList()

	cg.filenamesTable = make(map[string]string)
	cg.openFiles.Init()
	cg.unclean_shutdown = false
	cg.delayed_variance_obligations = nil
	cg.delayed_autoloads = nil
}

func (cg *ZendCompilerGlobals) Deactivate() {
	faults.Try(func() {
		cg.loopVarStack = nil
		cg.delayedOplinesStack = nil
		cg.filenamesTable = nil
		if cg.delayed_variance_obligations != nil {
			cg.delayed_variance_obligations.Destroy()
			cg.delayed_variance_obligations = nil
		}
		if cg.delayed_autoloads != nil {
			cg.delayed_autoloads.Destroy()
			cg.delayed_autoloads = nil
		}
	})
}

// class table
func (cg *ZendCompilerGlobals) ClassTable() ClassTable {
	return cg.classTable
}

func (cg *ZendCompilerGlobals) FunctionTable() FunctionTable { return cg.functionTable }

// compiler_options
func (cg *ZendCompilerGlobals) GetCompilerOptions() uint32      { return cg.compiler_options }
func (cg *ZendCompilerGlobals) SetCompilerOptions(value uint32) { cg.compiler_options = value }
func (cg *ZendCompilerGlobals) IsCompilePreload() bool {
	return cg.compiler_options&ZEND_COMPILE_PRELOAD != 0
}

// auto globals
func (cg *ZendCompilerGlobals) FindAutoGlobal(name string) *ZendAutoGlobal {
	return cg.autoGlobals[name]
}
func (cg *ZendCompilerGlobals) AddAutoGlobal(autoGlobal ZendAutoGlobal) bool {
	name := autoGlobal.Name()
	if _, exists := cg.autoGlobals[name]; exists {
		return false
	}

	cg.autoGlobals[name] = &autoGlobal
	return true
}
func (cg *ZendCompilerGlobals) EachAutoGlobal(fn func(*ZendAutoGlobal)) {
	for _, autoGlobal := range cg.autoGlobals {
		fn(autoGlobal)
	}
}

// stack
func (cg *ZendCompilerGlobals) LoopVarStackPush(loopVar *ZendLoopVar) {
	slicekit.Push(&cg.loopVarStack, loopVar)
}
func (cg *ZendCompilerGlobals) LoopVarStackPop() {
	slicekit.Pop(&cg.loopVarStack)
}
func (cg *ZendCompilerGlobals) LoopVarStackTop() *ZendLoopVar {
	loopVar, _ := slicekit.Last(&cg.loopVarStack)
	return loopVar
}
func (cg *ZendCompilerGlobals) LoopVarStackEach(f func(*ZendLoopVar) bool) {
	slicekit.EachReserveEx(cg.loopVarStack, f)
}
func (cg *ZendCompilerGlobals) LoopVarStackDepth() int {
	return len(cg.loopVarStack)
}

func (cg *ZendCompilerGlobals) DelayedOplinesStackPush(opline *types.ZendOp) {
	slicekit.Push(&cg.delayedOplinesStack, opline)
}
func (cg *ZendCompilerGlobals) DelayedOplinesStackTop() *types.ZendOp {
	opline, _ := slicekit.Last(&cg.delayedOplinesStack)
	return opline
}
func (cg *ZendCompilerGlobals) DelayedOplinesStackDepth() int {
	return len(cg.delayedOplinesStack)
}
func (cg *ZendCompilerGlobals) DelayedOplinesStackCut(offset int) []*types.ZendOp {
	b.Assert(offset < len(cg.delayedOplinesStack))
	ret := cg.delayedOplinesStack[offset:]
	cg.delayedOplinesStack = cg.delayedOplinesStack[:offset]
	return ret
}

// fileContext
func (cg *ZendCompilerGlobals) GetFileContext() *ZendFileContext { return &cg.file_context }
func (cg *ZendCompilerGlobals) FileContextBegin() (prevContext ZendFileContext) {
	prevContext = cg.file_context
	cg.file_context = ZendFileContext{}
	return prevContext
}
func (cg *ZendCompilerGlobals) FileContextEnd(prevContext ZendFileContext) {
	cg.file_context = prevContext
}

// open_files
func (cg *ZendCompilerGlobals) AddOpenFile(fh *FileHandle) {
	cg.openFiles.AddLast(fh)
}
func (cg *ZendCompilerGlobals) DelOpenFile(fh *FileHandle) {
	cg.openFiles.Filter(func(fileHandle *FileHandle) bool {
		if IsFileHandlesEquals(fileHandle, fh) {
			fileHandle.Destroy()
			return false
		}
		return true
	})
}
func (cg *ZendCompilerGlobals) CleanOpenFiles() {
	cg.openFiles.Each(func(fileHandle *FileHandle) {
		fileHandle.Destroy()
	})
	cg.openFiles.Clean()
}

// getter/setter
func (cg *ZendCompilerGlobals) GetActiveClassEntry() *types.ClassEntry {
	return cg.active_class_entry
}
func (cg *ZendCompilerGlobals) SetActiveClassEntry(value *types.ClassEntry) {
	cg.active_class_entry = value
}
func (cg *ZendCompilerGlobals) GetCompiledFilename() string {
	return cg.compiled_filename
}
func (cg *ZendCompilerGlobals) SetCompiledFilename(value string) {
	cg.compiled_filename = value
}
func (cg *ZendCompilerGlobals) GetZendLineno() int                   { return cg.zend_lineno }
func (cg *ZendCompilerGlobals) SetZendLineno(value int)              { cg.zend_lineno = value }
func (cg *ZendCompilerGlobals) GetActiveOpArray() *types.ZendOpArray { return cg.active_op_array }
func (cg *ZendCompilerGlobals) SetActiveOpArray(value *types.ZendOpArray) {
	cg.active_op_array = value
}
func (cg *ZendCompilerGlobals) GetParseError() bool         { return cg.parse_error }
func (cg *ZendCompilerGlobals) SetParseError(value bool)    { cg.parse_error = value }
func (cg *ZendCompilerGlobals) GetInCompilation() bool      { return cg.in_compilation }
func (cg *ZendCompilerGlobals) SetInCompilation(value bool) { cg.in_compilation = value }
func (cg *ZendCompilerGlobals) GetShortTags() bool          { return cg.short_tags }
func (cg *ZendCompilerGlobals) SetShortTags(value bool)     { cg.short_tags = value }
func (cg *ZendCompilerGlobals) GetUncleanShutdown() bool    { return cg.unclean_shutdown }
func (cg *ZendCompilerGlobals) SetUncleanShutdown(value bool) {
	cg.unclean_shutdown = value
}
func (cg *ZendCompilerGlobals) GetIniParserUnbufferedErrors() bool {
	return cg.ini_parser_unbuffered_errors
}
func (cg *ZendCompilerGlobals) SetIniParserUnbufferedErrors(value bool) {
	cg.ini_parser_unbuffered_errors = value
}
func (cg *ZendCompilerGlobals) GetIniParserParam() *ZendIniParserParam {
	return cg.ini_parser_param
}
func (cg *ZendCompilerGlobals) SetIniParserParam(value *ZendIniParserParam) {
	cg.ini_parser_param = value
}
func (cg *ZendCompilerGlobals) GetSkipShebang() bool      { return cg.skip_shebang }
func (cg *ZendCompilerGlobals) SetSkipShebang(value bool) { cg.skip_shebang = value }
func (cg *ZendCompilerGlobals) GetIncrementLineno() bool  { return cg.increment_lineno }
func (cg *ZendCompilerGlobals) SetIncrementLineno(value bool) {
	cg.increment_lineno = value
}
func (cg *ZendCompilerGlobals) GetDocComment() *types.String        { return cg.doc_comment }
func (cg *ZendCompilerGlobals) SetDocComment(value *types.String)   { cg.doc_comment = value }
func (cg *ZendCompilerGlobals) GetExtraFnFlags() uint32             { return cg.extra_fn_flags }
func (cg *ZendCompilerGlobals) SetExtraFnFlags(value uint32)        { cg.extra_fn_flags = value }
func (cg *ZendCompilerGlobals) GetContext() *ZendOparrayContext     { return &cg.context }
func (cg *ZendCompilerGlobals) SetContext(value ZendOparrayContext) { cg.context = value }
func (cg *ZendCompilerGlobals) GetAst() *ZendAst                    { return cg.ast }
func (cg *ZendCompilerGlobals) SetAst(value *ZendAst)               { cg.ast = value }
func (cg *ZendCompilerGlobals) GetMemoizedExprs() *types.Array      { return cg.memoized_exprs }
func (cg *ZendCompilerGlobals) SetMemoizedExprs(value *types.Array) {
	cg.memoized_exprs = value
}
func (cg *ZendCompilerGlobals) GetMemoizeMode() int      { return cg.memoize_mode }
func (cg *ZendCompilerGlobals) SetMemoizeMode(value int) { cg.memoize_mode = value }

func (cg *ZendCompilerGlobals) GetMapPtrBase() any      { return cg.map_ptr_base }
func (cg *ZendCompilerGlobals) SetMapPtrBase(value any) { cg.map_ptr_base = value }
func (cg *ZendCompilerGlobals) GetMapPtrSize() int      { return cg.map_ptr_size }
func (cg *ZendCompilerGlobals) SetMapPtrSize(value int) { cg.map_ptr_size = value }
func (cg *ZendCompilerGlobals) GetMapPtrLast() int      { return cg.map_ptr_last }
func (cg *ZendCompilerGlobals) SetMapPtrLast(value int) { cg.map_ptr_last = value }

func (cg *ZendCompilerGlobals) GetDelayedVarianceObligations() *types.Array {
	return cg.delayed_variance_obligations
}
func (cg *ZendCompilerGlobals) SetDelayedVarianceObligations(value *types.Array) {
	cg.delayed_variance_obligations = value
}
func (cg *ZendCompilerGlobals) GetDelayedAutoloads() *types.Array {
	return cg.delayed_autoloads
}
func (cg *ZendCompilerGlobals) SetDelayedAutoloads(value *types.Array) {
	cg.delayed_autoloads = value
}
func (cg *ZendCompilerGlobals) GetRtdKeyCounter() uint32      { return cg.rtd_key_counter }
func (cg *ZendCompilerGlobals) SetRtdKeyCounter(value uint32) { cg.rtd_key_counter = value }

/* ZendCompilerGlobals.extra_fn_flags */
func (cg *ZendCompilerGlobals) AddExtraFnFlags(value uint32) { cg.extra_fn_flags |= value }
func (cg *ZendCompilerGlobals) SubExtraFnFlags(value uint32) { cg.extra_fn_flags &^= value }
func (cg *ZendCompilerGlobals) HasExtraFnFlags(value uint32) bool {
	return cg.extra_fn_flags&value != 0
}
func (cg *ZendCompilerGlobals) SwitchExtraFnFlags(value uint32, cond bool) {
	if cond {
		cg.AddExtraFnFlags(value)
	} else {
		cg.SubExtraFnFlags(value)
	}
}

//
type CGBackupStack struct {
	activeClassEntry    *types.ClassEntry
	loopVarStack        []*ZendLoopVar
	delayedOplinesStack []*types.ZendOp
}

func (cg *ZendCompilerGlobals) BackupStack() (backup *CGBackupStack) {
	if !cg.in_compilation {
		backup = &CGBackupStack{
			activeClassEntry:    cg.active_class_entry,
			loopVarStack:        cg.loopVarStack,
			delayedOplinesStack: cg.delayedOplinesStack,
		}

		cg.in_compilation = false
		cg.active_class_entry = nil
		cg.loopVarStack = nil
		cg.delayedOplinesStack = nil
	}
	return
}
func (cg *ZendCompilerGlobals) RestorePostError(backup *CGBackupStack) {
	if backup == nil {
		return
	}

	cg.in_compilation = true
	cg.active_class_entry = backup.activeClassEntry
	cg.loopVarStack = backup.loopVarStack
	cg.delayedOplinesStack = backup.delayedOplinesStack
}
