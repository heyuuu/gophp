// <<generate>>

package zend

import (
	b "sik/builtin"
	"sik/core"
)

// Source: <Zend/zend_execute_API.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < signal . h >

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_execute.h"

// # include "zend_API.h"

// # include "zend_stack.h"

// # include "zend_constants.h"

// # include "zend_extensions.h"

// # include "zend_exceptions.h"

// # include "zend_closures.h"

// # include "zend_generators.h"

// # include "zend_vm.h"

// # include "zend_float.h"

// # include "zend_weakrefs.h"

// # include "zend_inheritance.h"

// # include < sys / time . h >

// # include < unistd . h >

/* true globals */

var EmptyFcallInfo ZendFcallInfo = ZendFcallInfo{0, {{0}, {{0}}, {0}}, nil, nil, nil, 0, 0}
var EmptyFcallInfoCache ZendFcallInfoCache = ZendFcallInfoCache{nil, nil, nil, nil}

func ZendExtensionActivator(extension *ZendExtension) {
	if extension.GetActivate() != nil {
		extension.GetActivate()()
	}
}

/* }}} */

func ZendExtensionDeactivator(extension *ZendExtension) {
	if extension.GetDeactivate() != nil {
		extension.GetDeactivate()()
	}
}

/* }}} */

func CleanNonPersistentConstantFull(zv *Zval) int {
	var c *ZendConstant = Z_PTR_P(zv)
	if (ZEND_CONSTANT_FLAGS(c) & CONST_PERSISTENT) != 0 {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}

/* }}} */

func CleanNonPersistentFunctionFull(zv *Zval) int {
	var function *ZendFunction = Z_PTR_P(zv)
	if function.GetType() == ZEND_INTERNAL_FUNCTION {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}

/* }}} */

func CleanNonPersistentClassFull(zv *Zval) int {
	var ce *ZendClassEntry = Z_PTR_P(zv)
	if ce.GetType() == ZEND_INTERNAL_CLASS {
		return ZEND_HASH_APPLY_KEEP
	} else {
		return ZEND_HASH_APPLY_REMOVE
	}
}

/* }}} */

func InitExecutor() {
	ZendInitFpu()
	ZVAL_NULL(&(ExecutorGlobals.GetUninitializedZval()))
	ZVAL_ERROR(&(ExecutorGlobals.GetErrorZval()))

	/* destroys stack frame, therefore makes core dumps worthless */

	ExecutorGlobals.SetSymtableCachePtr(ExecutorGlobals.GetSymtableCache())
	ExecutorGlobals.SetSymtableCacheLimit(ExecutorGlobals.GetSymtableCache() + SYMTABLE_CACHE_SIZE)
	ExecutorGlobals.SetNoExtensions(0)
	ExecutorGlobals.SetFunctionTable(CompilerGlobals.GetFunctionTable())
	ExecutorGlobals.SetClassTable(CompilerGlobals.GetClassTable())
	ExecutorGlobals.SetInAutoload(nil)
	ExecutorGlobals.SetAutoloadFunc(nil)
	ExecutorGlobals.SetErrorHandling(EH_NORMAL)
	ExecutorGlobals.SetFlags(EG_FLAGS_INITIAL)
	ZendVmStackInit()
	ZendHashInit(&(ExecutorGlobals.GetSymbolTable()), 64, nil, ZVAL_PTR_DTOR, 0)
	ZendLlistApply(&ZendExtensions, LlistApplyFuncT(ZendExtensionActivator))
	ZendHashInit(&(ExecutorGlobals.GetIncludedFiles()), 8, nil, nil, 0)
	ExecutorGlobals.SetTicksCount(0)
	ZVAL_UNDEF(&(ExecutorGlobals.GetUserErrorHandler()))
	ZVAL_UNDEF(&(ExecutorGlobals.GetUserExceptionHandler()))
	ExecutorGlobals.SetCurrentExecuteData(nil)
	ZendStackInit(&(ExecutorGlobals.GetUserErrorHandlersErrorReporting()), b.SizeOf("int"))
	ZendStackInit(&(ExecutorGlobals.GetUserErrorHandlers()), b.SizeOf("zval"))
	ZendStackInit(&(ExecutorGlobals.GetUserExceptionHandlers()), b.SizeOf("zval"))
	ZendObjectsStoreInit(&(ExecutorGlobals.GetObjectsStore()), 1024)
	ExecutorGlobals.SetFullTablesCleanup(0)
	ExecutorGlobals.SetVmInterrupt(0)
	ExecutorGlobals.SetTimedOut(0)
	ExecutorGlobals.SetException(nil)
	ExecutorGlobals.SetPrevException(nil)
	ExecutorGlobals.SetFakeScope(nil)
	ExecutorGlobals.GetTrampoline().SetFunctionName(nil)
	ExecutorGlobals.SetHtIteratorsCount(b.SizeOf("EG ( ht_iterators_slots )") / b.SizeOf("HashTableIterator"))
	ExecutorGlobals.SetHtIteratorsUsed(0)
	ExecutorGlobals.SetHtIterators(ExecutorGlobals.GetHtIteratorsSlots())
	memset(ExecutorGlobals.GetHtIterators(), 0, b.SizeOf("EG ( ht_iterators_slots )"))
	ExecutorGlobals.SetEachDeprecationThrown(0)
	ExecutorGlobals.SetPersistentConstantsCount(ExecutorGlobals.GetZendConstants().GetNNumUsed())
	ExecutorGlobals.SetPersistentFunctionsCount(ExecutorGlobals.GetFunctionTable().GetNNumUsed())
	ExecutorGlobals.SetPersistentClassesCount(ExecutorGlobals.GetClassTable().GetNNumUsed())
	ZendWeakrefsInit()
	ExecutorGlobals.SetActive(1)
}

/* }}} */

func ZvalCallDestructor(zv *Zval) int {
	if Z_TYPE_P(zv) == IS_INDIRECT {
		zv = Z_INDIRECT_P(zv)
	}
	if Z_TYPE_P(zv) == IS_OBJECT && Z_REFCOUNT_P(zv) == 1 {
		return ZEND_HASH_APPLY_REMOVE
	} else {
		return ZEND_HASH_APPLY_KEEP
	}
}

/* }}} */

func ZendUncleanZvalPtrDtor(zv *Zval) {
	if Z_TYPE_P(zv) == IS_INDIRECT {
		zv = Z_INDIRECT_P(zv)
	}
	IZvalPtrDtor(zv)
}

/* }}} */

func ZendThrowOrError(fetch_type int, exception_ce *ZendClassEntry, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) != 0 {
		ZendThrowError(exception_ce, "%s", message)
	} else {
		ZendError(E_ERROR, "%s", message)
	}
	Efree(message)
	va_end(va)
}

/* }}} */

func ShutdownDestructors() {
	if CompilerGlobals.GetUncleanShutdown() != 0 {
		ExecutorGlobals.GetSymbolTable().SetPDestructor(ZendUncleanZvalPtrDtor)
	}
	var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
	var __bailout JMP_BUF
	ExecutorGlobals.SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		var symbols uint32
		for {
			symbols = ZendHashNumElements(&(ExecutorGlobals.GetSymbolTable()))
			ZendHashReverseApply(&(ExecutorGlobals.GetSymbolTable()), ApplyFuncT(ZvalCallDestructor))
			if symbols == ZendHashNumElements(&(ExecutorGlobals.GetSymbolTable())) {
				break
			}
		}
		ZendObjectsStoreCallDestructors(&(ExecutorGlobals.GetObjectsStore()))
	} else {
		ExecutorGlobals.SetBailout(__orig_bailout)

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

		ZendObjectsStoreMarkDestructed(&(ExecutorGlobals.GetObjectsStore()))

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

	}
	ExecutorGlobals.SetBailout(__orig_bailout)
}

/* }}} */

func ShutdownExecutor() {
	var key *ZendString
	var zv *Zval
	var fast_shutdown ZendBool = IsZendMm() != 0 && ExecutorGlobals.GetFullTablesCleanup() == 0
	var __orig_bailout *JMP_BUF = executor_globals.bailout
	var __bailout JMP_BUF
	ExecutorGlobals.SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendLlistDestroy(&(CompilerGlobals.GetOpenFiles()))
	}
	ExecutorGlobals.SetBailout(__orig_bailout)
	ExecutorGlobals.SetFlags(ExecutorGlobals.GetFlags() | EG_FLAGS_IN_RESOURCE_SHUTDOWN)
	var __orig_bailout *JMP_BUF = executor_globals.bailout
	var __bailout JMP_BUF
	ExecutorGlobals.SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendCloseRsrcList(&(ExecutorGlobals.GetRegularList()))
	}
	ExecutorGlobals.SetBailout(__orig_bailout)

	/* No PHP callback functions should be called after this point. */

	ExecutorGlobals.SetActive(0)
	if fast_shutdown == 0 {
		ZendHashGracefulReverseDestroy(&(ExecutorGlobals.GetSymbolTable()))

		/* Release static properties and static variables prior to the final GC run,
		 * as they may hold GC roots. */

		for {
			var __ht *HashTable = ExecutorGlobals.GetFunctionTable()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				zv = _z
				var op_array *ZendOpArray = Z_PTR_P(zv)
				if op_array.GetType() == ZEND_INTERNAL_FUNCTION {
					break
				}
				if op_array.GetStaticVariables() != nil {
					var ht *HashTable = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
					if ht != nil {
						if (GC_FLAGS(ht)&IS_ARRAY_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
							ZendArrayDestroy(ht)
						}
						ZEND_MAP_PTR_SET(op_array.static_variables_ptr, nil)
					}
				}
			}
			break
		}
		for {
			var __ht *HashTable = ExecutorGlobals.GetClassTable()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
					continue
				}
				zv = _z
				var ce *ZendClassEntry = Z_PTR_P(zv)
				if ce.GetDefaultStaticMembersCount() != 0 {
					ZendCleanupInternalClassData(ce)
				}
				if (ce.GetCeFlags() & ZEND_HAS_STATIC_IN_METHODS) != 0 {
					var op_array *ZendOpArray
					for {
						var __ht *HashTable = &ce.function_table
						var _p *Bucket = __ht.GetArData()
						var _end *Bucket = _p + __ht.GetNNumUsed()
						for ; _p != _end; _p++ {
							var _z *Zval = &_p.val

							if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
								continue
							}
							op_array = Z_PTR_P(_z)
							if op_array.GetType() == ZEND_USER_FUNCTION {
								if op_array.GetStaticVariables() != nil {
									var ht *HashTable = ZEND_MAP_PTR_GET(op_array.static_variables_ptr)
									if ht != nil {
										if (GC_FLAGS(ht)&IS_ARRAY_IMMUTABLE) == 0 && GC_DELREF(ht) == 0 {
											ZendArrayDestroy(ht)
										}
										ZEND_MAP_PTR_SET(op_array.static_variables_ptr, nil)
									}
								}
							}
						}
						break
					}
				}
			}
			break
		}

		/* Also release error and exception handlers, which may hold objects. */

		if Z_TYPE(ExecutorGlobals.GetUserErrorHandler()) != IS_UNDEF {
			ZvalPtrDtor(&(ExecutorGlobals.GetUserErrorHandler()))
			ZVAL_UNDEF(&(ExecutorGlobals.GetUserErrorHandler()))
		}
		if Z_TYPE(ExecutorGlobals.GetUserExceptionHandler()) != IS_UNDEF {
			ZvalPtrDtor(&(ExecutorGlobals.GetUserExceptionHandler()))
			ZVAL_UNDEF(&(ExecutorGlobals.GetUserExceptionHandler()))
		}
		ZendStackClean(&(ExecutorGlobals.GetUserErrorHandlersErrorReporting()), nil, 1)
		ZendStackClean(&(ExecutorGlobals.GetUserErrorHandlers()), (func(any))(ZVAL_PTR_DTOR), 1)
		ZendStackClean(&(ExecutorGlobals.GetUserExceptionHandlers()), (func(any))(ZVAL_PTR_DTOR), 1)
	}
	ZendObjectsStoreFreeObjectStorage(&(ExecutorGlobals.GetObjectsStore()), fast_shutdown)
	ZendWeakrefsShutdown()
	var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
	var __bailout JMP_BUF
	ExecutorGlobals.SetBailout(&__bailout)
	if SETJMP(__bailout) == 0 {
		ZendLlistApply(&ZendExtensions, LlistApplyFuncT(ZendExtensionDeactivator))
	}
	ExecutorGlobals.SetBailout(__orig_bailout)
	if fast_shutdown != 0 {

		/* Fast Request Shutdown
		 * =====================
		 * Zend Memory Manager frees memory by its own. We don't have to free
		 * each allocated block separately.
		 */

		ZendHashDiscard(ExecutorGlobals.GetZendConstants(), ExecutorGlobals.GetPersistentConstantsCount())
		ZendHashDiscard(ExecutorGlobals.GetFunctionTable(), ExecutorGlobals.GetPersistentFunctionsCount())
		ZendHashDiscard(ExecutorGlobals.GetClassTable(), ExecutorGlobals.GetPersistentClassesCount())
		ZendCleanupInternalClasses()
	} else {
		ZendVmStackDestroy()
		if ExecutorGlobals.GetFullTablesCleanup() != 0 {
			ZendHashReverseApply(ExecutorGlobals.GetZendConstants(), CleanNonPersistentConstantFull)
			ZendHashReverseApply(ExecutorGlobals.GetFunctionTable(), CleanNonPersistentFunctionFull)
			ZendHashReverseApply(ExecutorGlobals.GetClassTable(), CleanNonPersistentClassFull)
		} else {
			for {
				var __ht *HashTable = ExecutorGlobals.GetZendConstants()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					key = _p.GetKey()
					zv = _z
					var c *ZendConstant = Z_PTR_P(zv)
					if _idx == ExecutorGlobals.GetPersistentConstantsCount() {
						break
					}
					ZvalPtrDtorNogc(&c.value)
					if c.GetName() != nil {
						ZendStringReleaseEx(c.GetName(), 0)
					}
					Efree(c)
					ZendStringReleaseEx(key, 0)
					__ht.GetNNumOfElements()--
					var j uint32 = HT_IDX_TO_HASH(_idx - 1)
					var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
					var i uint32 = HT_HASH(__ht, nIndex)
					if UNEXPECTED(j != i) {
						var prev *Bucket = HT_HASH_TO_BUCKET(__ht, i)
						for Z_NEXT(prev.GetVal()) != j {
							i = Z_NEXT(prev.GetVal())
							prev = HT_HASH_TO_BUCKET(__ht, i)
						}
						Z_NEXT(prev.GetVal()) = Z_NEXT(_p.GetVal())
					} else {
						HT_HASH(__ht, nIndex) = Z_NEXT(_p.GetVal())
					}
				}
				__ht.SetNNumUsed(_idx)
				break
			}
			for {
				var __ht *HashTable = ExecutorGlobals.GetFunctionTable()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					key = _p.GetKey()
					zv = _z
					var func_ *ZendFunction = Z_PTR_P(zv)
					if _idx == ExecutorGlobals.GetPersistentFunctionsCount() {
						break
					}
					DestroyOpArray(&func_.op_array)
					ZendStringReleaseEx(key, 0)
					__ht.GetNNumOfElements()--
					var j uint32 = HT_IDX_TO_HASH(_idx - 1)
					var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
					var i uint32 = HT_HASH(__ht, nIndex)
					if UNEXPECTED(j != i) {
						var prev *Bucket = HT_HASH_TO_BUCKET(__ht, i)
						for Z_NEXT(prev.GetVal()) != j {
							i = Z_NEXT(prev.GetVal())
							prev = HT_HASH_TO_BUCKET(__ht, i)
						}
						Z_NEXT(prev.GetVal()) = Z_NEXT(_p.GetVal())
					} else {
						HT_HASH(__ht, nIndex) = Z_NEXT(_p.GetVal())
					}
				}
				__ht.SetNNumUsed(_idx)
				break
			}
			for {
				var __ht *HashTable = ExecutorGlobals.GetClassTable()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if UNEXPECTED(Z_TYPE_P(_z) == IS_UNDEF) {
						continue
					}
					key = _p.GetKey()
					zv = _z
					if _idx == ExecutorGlobals.GetPersistentClassesCount() {
						break
					}
					DestroyZendClass(zv)
					ZendStringReleaseEx(key, 0)
					__ht.GetNNumOfElements()--
					var j uint32 = HT_IDX_TO_HASH(_idx - 1)
					var nIndex uint32 = _p.GetH() | __ht.GetNTableMask()
					var i uint32 = HT_HASH(__ht, nIndex)
					if UNEXPECTED(j != i) {
						var prev *Bucket = HT_HASH_TO_BUCKET(__ht, i)
						for Z_NEXT(prev.GetVal()) != j {
							i = Z_NEXT(prev.GetVal())
							prev = HT_HASH_TO_BUCKET(__ht, i)
						}
						Z_NEXT(prev.GetVal()) = Z_NEXT(_p.GetVal())
					} else {
						HT_HASH(__ht, nIndex) = Z_NEXT(_p.GetVal())
					}
				}
				__ht.SetNNumUsed(_idx)
				break
			}
		}
		for ExecutorGlobals.GetSymtableCachePtr() > ExecutorGlobals.GetSymtableCache() {
			ExecutorGlobals.GetSymtableCachePtr()--
			ZendHashDestroy(*(ExecutorGlobals.GetSymtableCachePtr()))
			FREE_HASHTABLE(*(ExecutorGlobals.GetSymtableCachePtr()))
		}
		ZendHashDestroy(&(ExecutorGlobals.GetIncludedFiles()))
		ZendStackDestroy(&(ExecutorGlobals.GetUserErrorHandlersErrorReporting()))
		ZendStackDestroy(&(ExecutorGlobals.GetUserErrorHandlers()))
		ZendStackDestroy(&(ExecutorGlobals.GetUserExceptionHandlers()))
		ZendObjectsStoreDestroy(&(ExecutorGlobals.GetObjectsStore()))
		if ExecutorGlobals.GetInAutoload() != nil {
			ZendHashDestroy(ExecutorGlobals.GetInAutoload())
			FREE_HASHTABLE(ExecutorGlobals.GetInAutoload())
		}
		if ExecutorGlobals.GetHtIterators() != ExecutorGlobals.GetHtIteratorsSlots() {
			Efree(ExecutorGlobals.GetHtIterators())
		}
	}
	ExecutorGlobals.SetHtIteratorsUsed(0)
	ZendShutdownFpu()
}

/* }}} */

func GetActiveClassName(space **byte) *byte {
	var func_ *ZendFunction
	if ZendIsExecuting() == 0 {
		if space != nil {
			*space = ""
		}
		return ""
	}
	func_ = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case ZEND_USER_FUNCTION:

	case ZEND_INTERNAL_FUNCTION:
		var ce *ZendClassEntry = func_.GetScope()
		if space != nil {
			if ce != nil {
				*space = "::"
			} else {
				*space = ""
			}
		}
		if ce != nil {
			return ZSTR_VAL(ce.GetName())
		} else {
			return ""
		}
	default:
		if space != nil {
			*space = ""
		}
		return ""
	}
}

/* }}} */

func GetActiveFunctionName() *byte {
	var func_ *ZendFunction
	if ZendIsExecuting() == 0 {
		return nil
	}
	func_ = ExecutorGlobals.GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case ZEND_USER_FUNCTION:
		var function_name *ZendString = func_.GetFunctionName()
		if function_name != nil {
			return ZSTR_VAL(function_name)
		} else {
			return "main"
		}
		break
	case ZEND_INTERNAL_FUNCTION:
		return ZSTR_VAL(func_.GetFunctionName())
		break
	default:
		return nil
	}
}

/* }}} */

func ZendGetExecutedFilename() *byte {
	var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ZSTR_VAL(ex.GetFunc().GetOpArray().GetFilename())
	} else {
		return "[no active file]"
	}
}

/* }}} */

func ZendGetExecutedFilenameEx() *ZendString {
	var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ex.GetFunc().GetOpArray().GetFilename()
	} else {
		return nil
	}
}

/* }}} */

func ZendGetExecutedLineno() uint32 {
	var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		if ExecutorGlobals.GetException() != nil && ex.GetOpline().GetOpcode() == ZEND_HANDLE_EXCEPTION && ex.GetOpline().GetLineno() == 0 && ExecutorGlobals.GetOplineBeforeException() != nil {
			return ExecutorGlobals.GetOplineBeforeException().GetLineno()
		}
		return ex.GetOpline().GetLineno()
	} else {
		return 0
	}
}

/* }}} */

func ZendGetExecutedScope() *ZendClassEntry {
	var ex *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for true {
		if ex == nil {
			return nil
		} else if ex.GetFunc() != nil && (ZEND_USER_CODE(ex.GetFunc().GetType()) || ex.GetFunc().GetScope() != nil) {
			return ex.GetFunc().GetScope()
		}
		ex = ex.GetPrevExecuteData()
	}
}

/* }}} */

func ZendIsExecuting() ZendBool {
	return ExecutorGlobals.GetCurrentExecuteData() != 0
}

/* }}} */

func ZendUseUndefinedConstant(name *ZendString, attr ZendAstAttr, result *Zval) int {
	var colon *byte
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		return FAILURE
	} else if b.Assign(&colon, (*byte)(ZendMemrchr(ZSTR_VAL(name), ':', ZSTR_LEN(name)))) {
		ZendThrowError(nil, "Undefined class constant '%s'", ZSTR_VAL(name))
		return FAILURE
	} else if (attr & IS_CONSTANT_UNQUALIFIED) == 0 {
		ZendThrowError(nil, "Undefined constant '%s'", ZSTR_VAL(name))
		return FAILURE
	} else {
		var actual *byte = ZSTR_VAL(name)
		var actual_len int = ZSTR_LEN(name)
		var slash *byte = (*byte)(ZendMemrchr(actual, '\\', actual_len))
		if slash != nil {
			actual = slash + 1
			actual_len -= actual - ZSTR_VAL(name)
		}
		ZendError(E_WARNING, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", actual, actual)
		if ExecutorGlobals.GetException() != nil {
			return FAILURE
		} else {
			var result_str *ZendString = ZendStringInit(actual, actual_len, 0)
			ZvalPtrDtorNogc(result)
			ZVAL_NEW_STR(result, result_str)
		}
	}
	return SUCCESS
}

/* }}} */

func ZvalUpdateConstantEx(p *Zval, scope *ZendClassEntry) int {
	if Z_TYPE_P(p) == IS_CONSTANT_AST {
		var ast *ZendAst = Z_ASTVAL_P(p)
		if ast.GetKind() == ZEND_AST_CONSTANT {
			var name *ZendString = ZendAstGetConstantName(ast)
			var zv *Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
			if UNEXPECTED(zv == nil) {
				return ZendUseUndefinedConstant(name, ast.GetAttr(), p)
			}
			ZvalPtrDtorNogc(p)
			ZVAL_COPY_OR_DUP(p, zv)
		} else {
			var tmp Zval
			if UNEXPECTED(ZendAstEvaluate(&tmp, ast, scope) != SUCCESS) {
				return FAILURE
			}
			ZvalPtrDtorNogc(p)
			ZVAL_COPY_VALUE(p, &tmp)
		}
	}
	return SUCCESS
}

/* }}} */

func ZvalUpdateConstant(pp *Zval) int {
	return ZvalUpdateConstantEx(pp, b.CondF(ExecutorGlobals.GetCurrentExecuteData() != nil, func() *ZendClassEntry { return ZendGetExecutedScope() }, func() *ZendClassEntry { return CompilerGlobals.GetActiveClassEntry() }))
}

/* }}} */

func _callUserFunctionEx(object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval, no_separation int) int {
	var fci ZendFcallInfo
	fci.SetSize(b.SizeOf("fci"))
	if object != nil {
		fci.SetObject(Z_OBJ_P(object))
	} else {
		fci.SetObject(nil)
	}
	ZVAL_COPY_VALUE(&fci.function_name, function_name)
	fci.SetRetval(retval_ptr)
	fci.SetParamCount(param_count)
	fci.SetParams(params)
	fci.SetNoSeparation(ZendBool(no_separation))
	return ZendCallFunction(&fci, nil)
}

/* }}} */

func ZendCallFunction(fci *ZendFcallInfo, fci_cache *ZendFcallInfoCache) int {
	var i uint32
	var call *ZendExecuteData
	var dummy_execute_data ZendExecuteData
	var fci_cache_local ZendFcallInfoCache
	var func_ *ZendFunction
	var call_info uint32
	var object_or_called_scope any
	ZVAL_UNDEF(fci.GetRetval())
	if ExecutorGlobals.GetActive() == 0 {
		return FAILURE
	}
	if ExecutorGlobals.GetException() != nil {
		return FAILURE
	}
	ZEND_ASSERT(fci.GetSize() == b.SizeOf("zend_fcall_info"))

	/* Initialize execute_data */

	if ExecutorGlobals.GetCurrentExecuteData() == nil {

		/* This only happens when we're called outside any execute()'s
		 * It shouldn't be strictly necessary to NULL execute_data out,
		 * but it may make bugs easier to spot
		 */

		memset(&dummy_execute_data, 0, b.SizeOf("zend_execute_data"))
		ExecutorGlobals.SetCurrentExecuteData(&dummy_execute_data)
	} else if ExecutorGlobals.GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetCommonType()) && ExecutorGlobals.GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL && ExecutorGlobals.GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_ICALL && ExecutorGlobals.GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_UCALL && ExecutorGlobals.GetCurrentExecuteData().GetOpline().GetOpcode() != ZEND_DO_FCALL_BY_NAME {

		/* Insert fake frame in case of include or magic calls */

		dummy_execute_data = *(ExecutorGlobals.GetCurrentExecuteData())
		dummy_execute_data.SetPrevExecuteData(ExecutorGlobals.GetCurrentExecuteData())
		dummy_execute_data.SetCall(nil)
		dummy_execute_data.SetOpline(nil)
		dummy_execute_data.SetFunc(nil)
		ExecutorGlobals.SetCurrentExecuteData(&dummy_execute_data)
	}
	if fci_cache == nil || fci_cache.GetFunctionHandler() == nil {
		var error *byte = nil
		if fci_cache == nil {
			fci_cache = &fci_cache_local
		}
		if ZendIsCallableEx(&fci.function_name, fci.GetObject(), IS_CALLABLE_CHECK_SILENT, nil, fci_cache, &error) == 0 {
			if error != nil {
				var callable_name *ZendString = ZendGetCallableNameEx(&fci.function_name, fci.GetObject())
				ZendError(E_WARNING, "Invalid callback %s, %s", ZSTR_VAL(callable_name), error)
				Efree(error)
				ZendStringReleaseEx(callable_name, 0)
			}
			if ExecutorGlobals.GetCurrentExecuteData() == &dummy_execute_data {
				ExecutorGlobals.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
			}
			return FAILURE
		} else if error != nil {

			/* Capitalize the first latter of the error message */

			if error[0] >= 'a' && error[0] <= 'z' {
				error[0] += 'A' - 'a'
			}
			ZendError(E_DEPRECATED, "%s", error)
			Efree(error)
			if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
				if ExecutorGlobals.GetCurrentExecuteData() == &dummy_execute_data {
					ExecutorGlobals.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				}
				return FAILURE
			}
		}
	}
	func_ = fci_cache.GetFunctionHandler()
	if (func_.GetFnFlags()&ZEND_ACC_STATIC) != 0 || fci_cache.GetObject() == nil {
		fci.SetObject(nil)
		object_or_called_scope = fci_cache.GetCalledScope()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC
	} else {
		fci.SetObject(fci_cache.GetObject())
		object_or_called_scope = fci.GetObject()
		call_info = ZEND_CALL_TOP_FUNCTION | ZEND_CALL_DYNAMIC | ZEND_CALL_HAS_THIS
	}
	call = ZendVmStackPushCallFrame(call_info, func_, fci.GetParamCount(), object_or_called_scope)
	if UNEXPECTED((func_.GetFnFlags() & ZEND_ACC_DEPRECATED) != 0) {
		ZendError(E_DEPRECATED, "Function %s%s%s() is deprecated", b.CondF1(func_.GetScope() != nil, func() []byte { return ZSTR_VAL(func_.GetScope().GetName()) }, ""), b.Cond(func_.GetScope() != nil, "::", ""), ZSTR_VAL(func_.GetFunctionName()))
		if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
			ZendVmStackFreeCallFrame(call)
			if ExecutorGlobals.GetCurrentExecuteData() == &dummy_execute_data {
				ExecutorGlobals.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				ZendRethrowException(ExecutorGlobals.GetCurrentExecuteData())
			}
			return FAILURE
		}
	}
	for i = 0; i < fci.GetParamCount(); i++ {
		var param *Zval
		var arg *Zval = &fci.params[i]
		var must_wrap ZendBool = 0
		if ARG_SHOULD_BE_SENT_BY_REF(func_, i+1) != 0 {
			if UNEXPECTED(!(Z_ISREF_P(arg))) {
				if fci.GetNoSeparation() == 0 {

					/* Separation is enabled -- create a ref */

					ZVAL_NEW_REF(arg, arg)

					/* Separation is enabled -- create a ref */

				} else if ARG_MAY_BE_SENT_BY_REF(func_, i+1) == 0 {

					/* By-value send is not allowed -- emit a warning,
					 * and perform the call with the value wrapped in a reference. */

					ZendError(E_WARNING, "Parameter %d to %s%s%s() expected to be a reference, value given", i+1, b.CondF1(func_.GetScope() != nil, func() []byte { return ZSTR_VAL(func_.GetScope().GetName()) }, ""), b.Cond(func_.GetScope() != nil, "::", ""), ZSTR_VAL(func_.GetFunctionName()))
					must_wrap = 1
					if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
						ZEND_CALL_NUM_ARGS(call) = i
						ZendVmStackFreeArgs(call)
						ZendVmStackFreeCallFrame(call)
						if ExecutorGlobals.GetCurrentExecuteData() == &dummy_execute_data {
							ExecutorGlobals.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
						}
						return FAILURE
					}
				}
			}
		} else {
			if Z_ISREF_P(arg) && (func_.GetFnFlags()&ZEND_ACC_CALL_VIA_TRAMPOLINE) == 0 {

				/* don't separate references for __call */

				arg = Z_REFVAL_P(arg)

				/* don't separate references for __call */

			}
		}
		param = ZEND_CALL_ARG(call, i+1)
		if EXPECTED(must_wrap == 0) {
			ZVAL_COPY(param, arg)
		} else {
			Z_TRY_ADDREF_P(arg)
			ZVAL_NEW_REF(param, arg)
		}
	}
	if UNEXPECTED((func_.GetOpArray().GetFnFlags() & ZEND_ACC_CLOSURE) != 0) {
		var call_info uint32
		GC_ADDREF(ZEND_CLOSURE_OBJECT(func_))
		call_info = ZEND_CALL_CLOSURE
		if (func_.GetFnFlags() & ZEND_ACC_FAKE_CLOSURE) != 0 {
			call_info |= ZEND_CALL_FAKE_CLOSURE
		}
		ZEND_ADD_CALL_FLAG(call, call_info)
	}
	if func_.GetType() == ZEND_USER_FUNCTION {
		var call_via_handler int = (func_.GetFnFlags() & ZEND_ACC_CALL_VIA_TRAMPOLINE) != 0
		var current_opline_before_exception *ZendOp = ExecutorGlobals.GetOplineBeforeException()
		ZendInitFuncExecuteData(call, &func_.op_array, fci.GetRetval())
		ZendExecuteEx(call)
		ExecutorGlobals.SetOplineBeforeException(current_opline_before_exception)
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else if func_.GetType() == ZEND_INTERNAL_FUNCTION {
		var call_via_handler int = (func_.GetFnFlags() & ZEND_ACC_CALL_VIA_TRAMPOLINE) != 0
		ZVAL_NULL(fci.GetRetval())
		call.SetPrevExecuteData(ExecutorGlobals.GetCurrentExecuteData())
		ExecutorGlobals.SetCurrentExecuteData(call)
		if EXPECTED(ZendExecuteInternal == nil) {

			/* saves one function call if zend_execute_internal is not used */

			func_.GetInternalFunction().GetHandler()(call, fci.GetRetval())

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, fci.GetRetval())
		}
		ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
		ZendVmStackFreeArgs(call)
		if ExecutorGlobals.GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			ZVAL_UNDEF(fci.GetRetval())
		}
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else {
		ZVAL_NULL(fci.GetRetval())

		/* Not sure what should be done here if it's a static method */

		if fci.GetObject() != nil {
			call.SetPrevExecuteData(ExecutorGlobals.GetCurrentExecuteData())
			ExecutorGlobals.SetCurrentExecuteData(call)
			fci.GetObject().GetHandlers().GetCallMethod()(func_.GetFunctionName(), fci.GetObject(), call, fci.GetRetval())
			ExecutorGlobals.SetCurrentExecuteData(call.GetPrevExecuteData())
		} else {
			ZendThrowError(nil, "Cannot call overloaded function for non-object")
		}
		ZendVmStackFreeArgs(call)
		if func_.GetType() == ZEND_OVERLOADED_FUNCTION_TEMPORARY {
			ZendStringReleaseEx(func_.GetFunctionName(), 0)
		}
		Efree(func_)
		if ExecutorGlobals.GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			ZVAL_UNDEF(fci.GetRetval())
		}
	}
	ZendVmStackFreeCallFrame(call)
	if ExecutorGlobals.GetCurrentExecuteData() == &dummy_execute_data {
		ExecutorGlobals.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
	}
	if UNEXPECTED(ExecutorGlobals.GetException() != nil) {
		if UNEXPECTED(ExecutorGlobals.GetCurrentExecuteData() == nil) {
			ZendThrowExceptionInternal(nil)
		} else if ExecutorGlobals.GetCurrentExecuteData().GetFunc() != nil && ZEND_USER_CODE(ExecutorGlobals.GetCurrentExecuteData().GetFunc().GetCommonType()) {
			ZendRethrowException(ExecutorGlobals.GetCurrentExecuteData())
		}
	}
	return SUCCESS
}

/* }}} */

func ZendLookupClassEx(name *ZendString, key *ZendString, flags uint32) *ZendClassEntry {
	var ce *ZendClassEntry = nil
	var args []Zval
	var zv *Zval
	var local_retval Zval
	var lc_name *ZendString
	var fcall_info ZendFcallInfo
	var fcall_cache ZendFcallInfoCache
	var orig_fake_scope *ZendClassEntry
	if key != nil {
		lc_name = key
	} else {
		if name == nil || ZSTR_LEN(name) == 0 {
			return nil
		}
		if ZSTR_VAL(name)[0] == '\\' {
			lc_name = ZendStringAlloc(ZSTR_LEN(name)-1, 0)
			ZendStrTolowerCopy(ZSTR_VAL(lc_name), ZSTR_VAL(name)+1, ZSTR_LEN(name)-1)
		} else {
			lc_name = ZendStringTolower(name)
		}
	}
	zv = ZendHashFind(ExecutorGlobals.GetClassTable(), lc_name)
	if zv != nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		ce = (*ZendClassEntry)(Z_PTR_P(zv))
		if UNEXPECTED((ce.GetCeFlags() & ZEND_ACC_LINKED) == 0) {
			if (flags&ZEND_FETCH_CLASS_ALLOW_UNLINKED) != 0 || (flags&ZEND_FETCH_CLASS_ALLOW_NEARLY_LINKED) != 0 && (ce.GetCeFlags()&ZEND_ACC_NEARLY_LINKED) != 0 {
				ce.SetCeFlags(ce.GetCeFlags() | ZEND_ACC_HAS_UNLINKED_USES)
				return ce
			}
			return nil
		}
		return ce
	}

	/* The compiler is not-reentrant. Make sure we __autoload() only during run-time
	 * (doesn't impact functionality of __autoload()
	 */

	if (flags&ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 || ZendIsCompiling() != 0 {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	if ExecutorGlobals.GetAutoloadFunc() == nil {
		var func_ *ZendFunction = ZendFetchFunction(ZSTR_KNOWN(ZEND_STR_MAGIC_AUTOLOAD))
		if func_ != nil {
			ExecutorGlobals.SetAutoloadFunc(func_)
		} else {
			if key == nil {
				ZendStringReleaseEx(lc_name, 0)
			}
			return nil
		}
	}

	/* Verify class name before passing it to __autoload() */

	if key == nil && strspn(ZSTR_VAL(name), "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377\\") != ZSTR_LEN(name) {
		ZendStringReleaseEx(lc_name, 0)
		return nil
	}
	if ExecutorGlobals.GetInAutoload() == nil {
		ALLOC_HASHTABLE(ExecutorGlobals.GetInAutoload())
		ZendHashInit(ExecutorGlobals.GetInAutoload(), 8, nil, nil, 0)
	}
	if ZendHashAddEmptyElement(ExecutorGlobals.GetInAutoload(), lc_name) == nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	ZVAL_UNDEF(&local_retval)
	if ZSTR_VAL(name)[0] == '\\' {
		ZVAL_STRINGL(&args[0], ZSTR_VAL(name)+1, ZSTR_LEN(name)-1)
	} else {
		ZVAL_STR_COPY(&args[0], name)
	}
	fcall_info.SetSize(b.SizeOf("fcall_info"))
	ZVAL_STR_COPY(&fcall_info.function_name, ExecutorGlobals.GetAutoloadFunc().GetFunctionName())
	fcall_info.SetRetval(&local_retval)
	fcall_info.SetParamCount(1)
	fcall_info.SetParams(args)
	fcall_info.SetObject(nil)
	fcall_info.SetNoSeparation(1)
	fcall_cache.SetFunctionHandler(ExecutorGlobals.GetAutoloadFunc())
	fcall_cache.SetCalledScope(nil)
	fcall_cache.SetObject(nil)
	orig_fake_scope = ExecutorGlobals.GetFakeScope()
	ExecutorGlobals.SetFakeScope(nil)
	ZendExceptionSave()
	if ZendCallFunction(&fcall_info, &fcall_cache) == SUCCESS && ExecutorGlobals.GetException() == nil {
		ce = ZendHashFindPtr(ExecutorGlobals.GetClassTable(), lc_name)
	}
	ZendExceptionRestore()
	ExecutorGlobals.SetFakeScope(orig_fake_scope)
	ZvalPtrDtor(&args[0])
	ZvalPtrDtorStr(&fcall_info.function_name)
	ZendHashDel(ExecutorGlobals.GetInAutoload(), lc_name)
	ZvalPtrDtor(&local_retval)
	if key == nil {
		ZendStringReleaseEx(lc_name, 0)
	}
	return ce
}

/* }}} */

func ZendLookupClass(name *ZendString) *ZendClassEntry { return ZendLookupClassEx(name, nil, 0) }

/* }}} */

func ZendGetCalledScope(ex *ZendExecuteData) *ZendClassEntry {
	for ex != nil {
		if Z_TYPE(ex.GetThis()) == IS_OBJECT {
			return Z_OBJCE(ex.GetThis())
		} else if Z_CE(ex.GetThis()) != nil {
			return Z_CE(ex.GetThis())
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}

/* }}} */

func ZendGetThisObject(ex *ZendExecuteData) *ZendObject {
	for ex != nil {
		if Z_TYPE(ex.GetThis()) == IS_OBJECT {
			return Z_OBJ(ex.GetThis())
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != ZEND_INTERNAL_FUNCTION || ex.GetFunc().GetScope() != nil {
				return nil
			}
		}
		ex = ex.GetPrevExecuteData()
	}
	return nil
}

/* }}} */

func ZendEvalStringl(str *byte, str_len int, retval_ptr *Zval, string_name *byte) int {
	var pv Zval
	var new_op_array *ZendOpArray
	var original_compiler_options uint32
	var retval int
	if retval_ptr != nil {
		ZVAL_NEW_STR(&pv, ZendStringAlloc(str_len+b.SizeOf("\"return ;\"")-1, 0))
		memcpy(Z_STRVAL(pv), "return ", b.SizeOf("\"return \"")-1)
		memcpy(Z_STRVAL(pv)+b.SizeOf("\"return \"")-1, str, str_len)
		Z_STRVAL(pv)[Z_STRLEN(pv)-1] = ';'
		Z_STRVAL(pv)[Z_STRLEN(pv)] = '0'
	} else {
		ZVAL_STRINGL(&pv, str, str_len)
	}

	/*printf("Evaluating '%s'\n", pv.value.str.val);*/

	original_compiler_options = CompilerGlobals.GetCompilerOptions()
	CompilerGlobals.SetCompilerOptions(ZEND_COMPILE_DEFAULT_FOR_EVAL)
	new_op_array = ZendCompileString(&pv, string_name)
	CompilerGlobals.SetCompilerOptions(original_compiler_options)
	if new_op_array != nil {
		var local_retval Zval
		ExecutorGlobals.SetNoExtensions(1)
		new_op_array.SetScope(ZendGetExecutedScope())
		var __orig_bailout *JMP_BUF = ExecutorGlobals.GetBailout()
		var __bailout JMP_BUF
		ExecutorGlobals.SetBailout(&__bailout)
		if SETJMP(__bailout) == 0 {
			ZVAL_UNDEF(&local_retval)
			ZendExecute(new_op_array, &local_retval)
		} else {
			ExecutorGlobals.SetBailout(__orig_bailout)
			DestroyOpArray(new_op_array)
			EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
			ZendBailout()
		}
		ExecutorGlobals.SetBailout(__orig_bailout)
		if Z_TYPE(local_retval) != IS_UNDEF {
			if retval_ptr != nil {
				ZVAL_COPY_VALUE(retval_ptr, &local_retval)
			} else {
				ZvalPtrDtor(&local_retval)
			}
		} else {
			if retval_ptr != nil {
				ZVAL_NULL(retval_ptr)
			}
		}
		ExecutorGlobals.SetNoExtensions(0)
		DestroyOpArray(new_op_array)
		EfreeSize(new_op_array, b.SizeOf("zend_op_array"))
		retval = SUCCESS
	} else {
		retval = FAILURE
	}
	ZvalPtrDtorStr(&pv)
	return retval
}

/* }}} */

func ZendEvalString(str *byte, retval_ptr *Zval, string_name *byte) int {
	return ZendEvalStringl(str, strlen(str), retval_ptr, string_name)
}

/* }}} */

func ZendEvalStringlEx(str *byte, str_len int, retval_ptr *Zval, string_name *byte, handle_exceptions int) int {
	var result int
	result = ZendEvalStringl(str, str_len, retval_ptr, string_name)
	if handle_exceptions != 0 && ExecutorGlobals.GetException() != nil {
		ZendExceptionError(ExecutorGlobals.GetException(), E_ERROR)
		result = FAILURE
	}
	return result
}

/* }}} */

func ZendEvalStringEx(str *byte, retval_ptr *Zval, string_name string, handle_exceptions int) int {
	return ZendEvalStringlEx(str, strlen(str), retval_ptr, string_name, handle_exceptions)
}

/* }}} */

func ZendTimeout(dummy int) {
	ExecutorGlobals.SetTimedOut(0)
	ZendSetTimeoutEx(0, 1)
	ZendErrorNoreturn(E_ERROR, "Maximum execution time of "+ZEND_LONG_FMT+" second%s exceeded", ExecutorGlobals.GetTimeoutSeconds(), b.Cond(ExecutorGlobals.GetTimeoutSeconds() == 1, "", "s"))
}

/* }}} */

func ZendTimeoutHandler(dummy int) {
	if ExecutorGlobals.GetTimedOut() != 0 {

		/* Die on hard timeout */

		var error_filename *byte = nil
		var error_lineno uint32 = 0
		var log_buffer []byte
		var output_len int = 0
		if ZendIsCompiling() != 0 {
			error_filename = ZSTR_VAL(ZendGetCompiledFilename())
			error_lineno = ZendGetCompiledLineno()
		} else if ZendIsExecuting() != 0 {
			error_filename = ZendGetExecutedFilename()
			if error_filename[0] == '[' {
				error_filename = nil
				error_lineno = 0
			} else {
				error_lineno = ZendGetExecutedLineno()
			}
		}
		if error_filename == nil {
			error_filename = "Unknown"
		}
		output_len = core.Snprintf(log_buffer, b.SizeOf("log_buffer"), "\nFatal error: Maximum execution time of "+ZEND_LONG_FMT+"+"+ZEND_LONG_FMT+" seconds exceeded (terminated) in %s on line %d\n", ExecutorGlobals.GetTimeoutSeconds(), ExecutorGlobals.GetHardTimeout(), error_filename, error_lineno)
		if output_len > 0 {
			ZendQuietWrite(2, log_buffer, MIN(output_len, b.SizeOf("log_buffer")))
		}
		_exit(124)
	}
	if ZendOnTimeout != nil {

		/*
		   We got here because we got a timeout signal, so we are in a signal handler
		   at this point. However, we want to be able to timeout any user-supplied
		   shutdown functions, so pretend we are not in a signal handler while we are
		   calling these
		*/

		SIGG(running) = 0
		ZendOnTimeout(ExecutorGlobals.GetTimeoutSeconds())
	}
	ExecutorGlobals.SetTimedOut(1)
	ExecutorGlobals.SetVmInterrupt(1)
	if ExecutorGlobals.GetHardTimeout() > 0 {

		/* Set hard timeout */

		ZendSetTimeoutEx(ExecutorGlobals.GetHardTimeout(), 1)

		/* Set hard timeout */

	}
}

/* }}} */

/* This one doesn't exists on QNX */

const SIGPROF = 27

func ZendSetTimeoutEx(seconds ZendLong, reset_signals int) {
	var t_r __struct__itimerval
	var signo int
	if seconds != 0 {
		t_r.it_value.tv_sec = seconds
		t_r.it_interval.tv_usec = 0
		t_r.it_interval.tv_sec = t_r.it_interval.tv_usec
		t_r.it_value.tv_usec = t_r.it_interval.tv_sec
		setitimer(ITIMER_PROF, &t_r, nil)
	}
	signo = SIGPROF
	if reset_signals != 0 {
		ZendSignal(signo, ZendTimeoutHandler)
	}
}

/* }}} */

func ZendSetTimeout(seconds ZendLong, reset_signals int) {
	ExecutorGlobals.SetTimeoutSeconds(seconds)
	ZendSetTimeoutEx(seconds, reset_signals)
	ExecutorGlobals.SetTimedOut(0)
}

/* }}} */

func ZendUnsetTimeout() {
	if ExecutorGlobals.GetTimeoutSeconds() != 0 {
		var no_timeout __struct__itimerval
		no_timeout.it_interval.tv_usec = 0
		no_timeout.it_interval.tv_sec = no_timeout.it_interval.tv_usec
		no_timeout.it_value.tv_usec = no_timeout.it_interval.tv_sec
		no_timeout.it_value.tv_sec = no_timeout.it_value.tv_usec
		setitimer(ITIMER_PROF, &no_timeout, nil)
	}
	ExecutorGlobals.SetTimedOut(0)
}

/* }}} */

func ZendFetchClass(class_name *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var fetch_sub_type int = fetch_type & ZEND_FETCH_CLASS_MASK
check_fetch_type:
	switch fetch_sub_type {
	case ZEND_FETCH_CLASS_SELF:
		scope = ZendGetExecutedScope()
		if UNEXPECTED(scope == nil) {
			ZendThrowOrError(fetch_type, nil, "Cannot access self:: when no class scope is active")
		}
		return scope
	case ZEND_FETCH_CLASS_PARENT:
		scope = ZendGetExecutedScope()
		if UNEXPECTED(scope == nil) {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when no class scope is active")
			return nil
		}
		if UNEXPECTED(!(scope.parent)) {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when current class scope has no parent")
		}
		return scope.parent
	case ZEND_FETCH_CLASS_STATIC:
		ce = ZendGetCalledScope(ExecutorGlobals.GetCurrentExecuteData())
		if UNEXPECTED(ce == nil) {
			ZendThrowOrError(fetch_type, nil, "Cannot access static:: when no class scope is active")
			return nil
		}
		return ce
	case ZEND_FETCH_CLASS_AUTO:
		fetch_sub_type = ZendGetClassFetchType(class_name)
		if UNEXPECTED(fetch_sub_type != ZEND_FETCH_CLASS_DEFAULT) {
			goto check_fetch_type
		}
		break
	}
	if (fetch_type & ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 {
		return ZendLookupClassEx(class_name, nil, fetch_type)
	} else if b.Assign(&ce, ZendLookupClassEx(class_name, nil, fetch_type)) == nil {
		if (fetch_type&ZEND_FETCH_CLASS_SILENT) == 0 && ExecutorGlobals.GetException() == nil {
			if fetch_sub_type == ZEND_FETCH_CLASS_INTERFACE {
				ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", ZSTR_VAL(class_name))
			} else if fetch_sub_type == ZEND_FETCH_CLASS_TRAIT {
				ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", ZSTR_VAL(class_name))
			} else {
				ZendThrowOrError(fetch_type, nil, "Class '%s' not found", ZSTR_VAL(class_name))
			}
		}
		return nil
	}
	return ce
}

/* }}} */

func ZendFetchClassByName(class_name *ZendString, key *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	if (fetch_type & ZEND_FETCH_CLASS_NO_AUTOLOAD) != 0 {
		return ZendLookupClassEx(class_name, key, fetch_type)
	} else if b.Assign(&ce, ZendLookupClassEx(class_name, key, fetch_type)) == nil {
		if (fetch_type & ZEND_FETCH_CLASS_SILENT) != 0 {
			return nil
		}
		if ExecutorGlobals.GetException() != nil {
			if (fetch_type & ZEND_FETCH_CLASS_EXCEPTION) == 0 {
				var exception_str *ZendString
				var exception_zv Zval
				ZVAL_OBJ(&exception_zv, ExecutorGlobals.GetException())
				Z_ADDREF(exception_zv)
				ZendClearException()
				exception_str = ZvalGetString(&exception_zv)
				ZendErrorNoreturn(E_ERROR, "During class fetch: Uncaught %s", ZSTR_VAL(exception_str))
			}
			return nil
		}
		if (fetch_type & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_INTERFACE {
			ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", ZSTR_VAL(class_name))
		} else if (fetch_type & ZEND_FETCH_CLASS_MASK) == ZEND_FETCH_CLASS_TRAIT {
			ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", ZSTR_VAL(class_name))
		} else {
			ZendThrowOrError(fetch_type, nil, "Class '%s' not found", ZSTR_VAL(class_name))
		}
		return nil
	}
	return ce
}

/* }}} */

func ZendDeleteGlobalVariable(name *ZendString) int {
	return ZendHashDelInd(&(ExecutorGlobals.GetSymbolTable()), name)
}

/* }}} */

func ZendRebuildSymbolTable() *ZendArray {
	var ex *ZendExecuteData
	var symbol_table *ZendArray

	/* Search for last called user function */

	ex = ExecutorGlobals.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || !(ZEND_USER_CODE(ex.GetFunc().GetCommonType()))) {
		ex = ex.GetPrevExecuteData()
	}
	if ex == nil {
		return nil
	}
	if (ZEND_CALL_INFO(ex) & ZEND_CALL_HAS_SYMBOL_TABLE) != 0 {
		return ex.GetSymbolTable()
	}
	ZEND_ADD_CALL_FLAG(ex, ZEND_CALL_HAS_SYMBOL_TABLE)
	if ExecutorGlobals.GetSymtableCachePtr() > ExecutorGlobals.GetSymtableCache() {
		ex.SetSymbolTable(*(b.PreDec(&(ExecutorGlobals.GetSymtableCachePtr()))))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		ZendHashExtend(symbol_table, ex.GetFunc().GetOpArray().GetLastVar(), 0)
	} else {
		ex.SetSymbolTable(ZendNewArray(ex.GetFunc().GetOpArray().GetLastVar()))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		ZendHashRealInitMixed(symbol_table)
	}
	if EXPECTED(ex.GetFunc().GetOpArray().GetLastVar() != 0) {
		var str **ZendString = ex.GetFunc().GetOpArray().GetVars()
		var end **ZendString = str + ex.GetFunc().GetOpArray().GetLastVar()
		var var_ *Zval = ZEND_CALL_VAR_NUM(ex, 0)
		for {
			_zendHashAppendInd(symbol_table, *str, var_)
			str++
			var_++
			if str == end {
				break
			}
		}
	}
	return symbol_table
}

/* }}} */

func ZendAttachSymbolTable(execute_data *ZendExecuteData) {
	var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
	var ht *HashTable = execute_data.GetSymbolTable()

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */

	if EXPECTED(op_array.GetLastVar() != 0) {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = EX_VAR_NUM(0)
		for {
			var zv *Zval = ZendHashFindEx(ht, *str, 1)
			if zv != nil {
				if Z_TYPE_P(zv) == IS_INDIRECT {
					var val *Zval = Z_INDIRECT_P(zv)
					ZVAL_COPY_VALUE(var_, val)
				} else {
					ZVAL_COPY_VALUE(var_, zv)
				}
			} else {
				ZVAL_UNDEF(var_)
				zv = ZendHashAddNew(ht, *str, var_)
			}
			ZVAL_INDIRECT(zv, var_)
			str++
			var_++
			if str == end {
				break
			}
		}
	}

	/* copy real values from symbol table into CV slots and create
	   INDIRECT references to CV in symbol table  */
}

/* }}} */

func ZendDetachSymbolTable(execute_data *ZendExecuteData) {
	var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
	var ht *HashTable = execute_data.GetSymbolTable()

	/* copy real values from CV slots into symbol table */

	if EXPECTED(op_array.GetLastVar() != 0) {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = EX_VAR_NUM(0)
		for {
			if Z_TYPE_P(var_) == IS_UNDEF {
				ZendHashDel(ht, *str)
			} else {
				ZendHashUpdate(ht, *str, var_)
				ZVAL_UNDEF(var_)
			}
			str++
			var_++
			if str == end {
				break
			}
		}
	}

	/* copy real values from CV slots into symbol table */
}

/* }}} */

func ZendSetLocalVar(name *ZendString, value *Zval, force int) int {
	var execute_data *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for execute_data != nil && (execute_data.GetFunc() == nil || !(ZEND_USER_CODE(execute_data.GetFunc().GetCommonType()))) {
		execute_data = execute_data.GetPrevExecuteData()
	}
	if execute_data != nil {
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			var h ZendUlong = ZendStringHashVal(name)
			var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
			if EXPECTED(op_array.GetLastVar() != 0) {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if ZSTR_H(*str) == h && ZendStringEqualContent(*str, name) != 0 {
						var var_ *Zval = EX_VAR_NUM(str - op_array.GetVars())
						ZVAL_COPY_VALUE(var_, value)
						return SUCCESS
					}
					str++
					if str == end {
						break
					}
				}
			}
			if force != 0 {
				var symbol_table *ZendArray = ZendRebuildSymbolTable()
				if symbol_table != nil {
					ZendHashUpdate(symbol_table, name, value)
					return SUCCESS
				}
			}
		} else {
			ZendHashUpdateInd(execute_data.GetSymbolTable(), name, value)
			return SUCCESS
		}
	}
	return FAILURE
}

/* }}} */

func ZendSetLocalVarStr(name string, len_ int, value *Zval, force int) int {
	var execute_data *ZendExecuteData = ExecutorGlobals.GetCurrentExecuteData()
	for execute_data != nil && (execute_data.GetFunc() == nil || !(ZEND_USER_CODE(execute_data.GetFunc().GetCommonType()))) {
		execute_data = execute_data.GetPrevExecuteData()
	}
	if execute_data != nil {
		if (EX_CALL_INFO() & ZEND_CALL_HAS_SYMBOL_TABLE) == 0 {
			var h ZendUlong = ZendHashFunc(name, len_)
			var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
			if EXPECTED(op_array.GetLastVar() != 0) {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if ZSTR_H(*str) == h && ZSTR_LEN(*str) == len_ && memcmp(ZSTR_VAL(*str), name, len_) == 0 {
						var var_ *Zval = EX_VAR_NUM(str - op_array.GetVars())
						ZvalPtrDtor(var_)
						ZVAL_COPY_VALUE(var_, value)
						return SUCCESS
					}
					str++
					if str == end {
						break
					}
				}
			}
			if force != 0 {
				var symbol_table *ZendArray = ZendRebuildSymbolTable()
				if symbol_table != nil {
					ZendHashStrUpdate(symbol_table, name, len_, value)
					return SUCCESS
				}
			}
		} else {
			ZendHashStrUpdateInd(execute_data.GetSymbolTable(), name, len_, value)
			return SUCCESS
		}
	}
	return FAILURE
}

/* }}} */
