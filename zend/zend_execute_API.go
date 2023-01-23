// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
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
	var c *ZendConstant = zv.GetValue().GetPtr()
	if (c.GetValue().GetConstantFlags() & 0xff & 1 << 1) != 0 {
		return 0
	} else {
		return 1 << 0
	}
}

/* }}} */

func CleanNonPersistentFunctionFull(zv *Zval) int {
	var function *ZendFunction = zv.GetValue().GetPtr()
	if function.GetType() == 1 {
		return 0
	} else {
		return 1 << 0
	}
}

/* }}} */

func CleanNonPersistentClassFull(zv *Zval) int {
	var ce *ZendClassEntry = zv.GetValue().GetPtr()
	if ce.GetType() == 1 {
		return 0
	} else {
		return 1 << 0
	}
}

/* }}} */

func InitExecutor() {
	ZendInitFpu()
	&EG.uninitialized_zval.u1.type_info = 1
	&EG.error_zval.u1.type_info = 15

	/* destroys stack frame, therefore makes core dumps worthless */

	EG.SetSymtableCachePtr(EG.GetSymtableCache())
	EG.SetSymtableCacheLimit(EG.GetSymtableCache() + 32)
	EG.SetNoExtensions(0)
	EG.SetFunctionTable(CG.GetFunctionTable())
	EG.SetClassTable(CG.GetClassTable())
	EG.SetInAutoload(nil)
	EG.SetAutoloadFunc(nil)
	EG.SetErrorHandling(EH_NORMAL)
	EG.SetFlags(0)
	ZendVmStackInit()
	_zendHashInit(&EG.symbol_table, 64, ZvalPtrDtor, 0)
	ZendLlistApply(&ZendExtensions, LlistApplyFuncT(ZendExtensionActivator))
	_zendHashInit(&EG.included_files, 8, nil, 0)
	EG.SetTicksCount(0)
	&EG.user_error_handler.u1.type_info = 0
	&EG.user_exception_handler.u1.type_info = 0
	EG.SetCurrentExecuteData(nil)
	ZendStackInit(&EG.user_error_handlers_error_reporting, g.SizeOf("int"))
	ZendStackInit(&EG.user_error_handlers, g.SizeOf("zval"))
	ZendStackInit(&EG.user_exception_handlers, g.SizeOf("zval"))
	ZendObjectsStoreInit(&EG.objects_store, 1024)
	EG.SetFullTablesCleanup(0)
	EG.SetVmInterrupt(0)
	EG.SetTimedOut(0)
	EG.SetException(nil)
	EG.SetPrevException(nil)
	EG.SetFakeScope(nil)
	EG.GetTrampoline().SetFunctionName(nil)
	EG.SetHtIteratorsCount(g.SizeOf("EG ( ht_iterators_slots )") / g.SizeOf("HashTableIterator"))
	EG.SetHtIteratorsUsed(0)
	EG.SetHtIterators(EG.GetHtIteratorsSlots())
	memset(EG.GetHtIterators(), 0, g.SizeOf("EG ( ht_iterators_slots )"))
	EG.SetEachDeprecationThrown(0)
	EG.SetPersistentConstantsCount(EG.GetZendConstants().GetNNumUsed())
	EG.SetPersistentFunctionsCount(EG.GetFunctionTable().GetNNumUsed())
	EG.SetPersistentClassesCount(EG.GetClassTable().GetNNumUsed())
	ZendWeakrefsInit()
	EG.SetActive(1)
}

/* }}} */

func ZvalCallDestructor(zv *Zval) int {
	if zv.GetType() == 13 {
		zv = zv.GetValue().GetZv()
	}
	if zv.GetType() == 8 && ZvalRefcountP(zv) == 1 {
		return 1 << 0
	} else {
		return 0
	}
}

/* }}} */

func ZendUncleanZvalPtrDtor(zv *Zval) {
	if zv.GetType() == 13 {
		zv = zv.GetValue().GetZv()
	}
	IZvalPtrDtor(zv)
}

/* }}} */

func ZendThrowOrError(fetch_type int, exception_ce *ZendClassEntry, format string, _ ...any) {
	var va va_list
	var message *byte = nil
	va_start(va, format)
	ZendVspprintf(&message, 0, format, va)
	if (fetch_type & 0x200) != 0 {
		ZendThrowError(exception_ce, "%s", message)
	} else {
		ZendError(1<<0, "%s", message)
	}
	_efree(message)
	va_end(va)
}

/* }}} */

func ShutdownDestructors() {
	if CG.GetUncleanShutdown() != 0 {
		EG.GetSymbolTable().SetPDestructor(ZendUncleanZvalPtrDtor)
	}
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		var symbols uint32
		for {
			symbols = &EG.symbol_table.nNumOfElements
			ZendHashReverseApply(&EG.symbol_table, ApplyFuncT(ZvalCallDestructor))
			if symbols == &EG.symbol_table.nNumOfElements {
				break
			}
		}
		ZendObjectsStoreCallDestructors(&EG.objects_store)
	} else {
		EG.SetBailout(__orig_bailout)

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

		ZendObjectsStoreMarkDestructed(&EG.objects_store)

		/* if we couldn't destruct cleanly, mark all objects as destructed anyway */

	}
	EG.SetBailout(__orig_bailout)
}

/* }}} */

func ShutdownExecutor() {
	var key *ZendString
	var zv *Zval
	var fast_shutdown ZendBool = IsZendMm() != 0 && EG.GetFullTablesCleanup() == 0
	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ZendLlistDestroy(&CG.open_files)
	}
	EG.SetBailout(__orig_bailout)
	EG.SetFlags(EG.GetFlags() | 1<<2)
	var __orig_bailout *sigjmp_buf = EG.bailout
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ZendCloseRsrcList(&EG.regular_list)
	}
	EG.SetBailout(__orig_bailout)

	/* No PHP callback functions should be called after this point. */

	EG.SetActive(0)
	if fast_shutdown == 0 {
		ZendHashGracefulReverseDestroy(&EG.symbol_table)

		/* Release static properties and static variables prior to the final GC run,
		 * as they may hold GC roots. */

		for {
			var __ht *HashTable = EG.GetFunctionTable()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				zv = _z
				var op_array *ZendOpArray = zv.GetValue().GetPtr()
				if op_array.GetType() == 1 {
					break
				}
				if op_array.GetStaticVariables() != nil {
					var ht *HashTable = g.CondF((uintPtr(op_array.GetStaticVariablesPtrPtr())&1) != 0, func() any {
						return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetStaticVariablesPtrPtr()-1))))
					}, func() any { return any(*(op_array.GetStaticVariablesPtrPtr())) })
					if ht != nil {
						if (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
							ZendArrayDestroy(ht)
						}
						if (uintPtr(op_array.GetStaticVariablesPtrPtr()) & 1) != 0 {
							*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetStaticVariablesPtrPtr()-1)))) = nil
						} else {
							*(op_array.GetStaticVariablesPtrPtr()) = nil
						}
					}
				}
			}
			break
		}
		for {
			var __ht *HashTable = EG.GetClassTable()
			var _idx uint32 = __ht.GetNNumUsed()
			var _p *Bucket = __ht.GetArData() + _idx
			var _z *Zval
			for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
				_p--
				_z = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				zv = _z
				var ce *ZendClassEntry = zv.GetValue().GetPtr()
				if ce.GetDefaultStaticMembersCount() != 0 {
					ZendCleanupInternalClassData(ce)
				}
				if (ce.GetCeFlags() & 1 << 16) != 0 {
					var op_array *ZendOpArray
					for {
						var __ht *HashTable = &ce.function_table
						var _p *Bucket = __ht.GetArData()
						var _end *Bucket = _p + __ht.GetNNumUsed()
						for ; _p != _end; _p++ {
							var _z *Zval = &_p.val

							if _z.GetType() == 0 {
								continue
							}
							op_array = _z.GetValue().GetPtr()
							if op_array.GetType() == 2 {
								if op_array.GetStaticVariables() != nil {
									var ht *HashTable = g.CondF((uintPtr(op_array.GetStaticVariablesPtrPtr())&1) != 0, func() any {
										return *((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetStaticVariablesPtrPtr()-1))))
									}, func() any { return any(*(op_array.GetStaticVariablesPtrPtr())) })
									if ht != nil {
										if (ZvalGcFlags(ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&ht.gc) == 0 {
											ZendArrayDestroy(ht)
										}
										if (uintPtr(op_array.GetStaticVariablesPtrPtr()) & 1) != 0 {
											*((*any)((*byte)(CG.GetMapPtrBase() + uintPtr(op_array.GetStaticVariablesPtrPtr()-1)))) = nil
										} else {
											*(op_array.GetStaticVariablesPtrPtr()) = nil
										}
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

		if EG.GetUserErrorHandler().GetType() != 0 {
			ZvalPtrDtor(&EG.user_error_handler)
			&EG.user_error_handler.u1.type_info = 0
		}
		if EG.GetUserExceptionHandler().GetType() != 0 {
			ZvalPtrDtor(&EG.user_exception_handler)
			&EG.user_exception_handler.u1.type_info = 0
		}
		ZendStackClean(&EG.user_error_handlers_error_reporting, nil, 1)
		ZendStackClean(&EG.user_error_handlers, (func(any))(ZvalPtrDtor), 1)
		ZendStackClean(&EG.user_exception_handlers, (func(any))(ZvalPtrDtor), 1)
	}
	ZendObjectsStoreFreeObjectStorage(&EG.objects_store, fast_shutdown)
	ZendWeakrefsShutdown()
	var __orig_bailout *sigjmp_buf = EG.GetBailout()
	var __bailout sigjmp_buf
	EG.SetBailout(&__bailout)
	if sigsetjmp(__bailout, 0) == 0 {
		ZendLlistApply(&ZendExtensions, LlistApplyFuncT(ZendExtensionDeactivator))
	}
	EG.SetBailout(__orig_bailout)
	if fast_shutdown != 0 {

		/* Fast Request Shutdown
		 * =====================
		 * Zend Memory Manager frees memory by its own. We don't have to free
		 * each allocated block separately.
		 */

		ZendHashDiscard(EG.GetZendConstants(), EG.GetPersistentConstantsCount())
		ZendHashDiscard(EG.GetFunctionTable(), EG.GetPersistentFunctionsCount())
		ZendHashDiscard(EG.GetClassTable(), EG.GetPersistentClassesCount())
		ZendCleanupInternalClasses()
	} else {
		ZendVmStackDestroy()
		if EG.GetFullTablesCleanup() != 0 {
			ZendHashReverseApply(EG.GetZendConstants(), CleanNonPersistentConstantFull)
			ZendHashReverseApply(EG.GetFunctionTable(), CleanNonPersistentFunctionFull)
			ZendHashReverseApply(EG.GetClassTable(), CleanNonPersistentClassFull)
		} else {
			for {
				var __ht *HashTable = EG.GetZendConstants()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					key = _p.GetKey()
					zv = _z
					var c *ZendConstant = zv.GetValue().GetPtr()
					if _idx == EG.GetPersistentConstantsCount() {
						break
					}
					ZvalPtrDtorNogc(&c.value)
					if c.GetName() != nil {
						ZendStringReleaseEx(c.GetName(), 0)
					}
					_efree(c)
					ZendStringReleaseEx(key, 0)
					ZEND_HASH_FOREACH_END_DEL_ITEM(__ht, _idx, _p)
				}
				__ht.SetNNumUsed(_idx)
				break
			}
			for {
				var __ht *HashTable = EG.GetFunctionTable()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					key = _p.GetKey()
					zv = _z
					var func_ *ZendFunction = zv.GetValue().GetPtr()
					if _idx == EG.GetPersistentFunctionsCount() {
						break
					}
					DestroyOpArray(&func_.op_array)
					ZendStringReleaseEx(key, 0)
					ZEND_HASH_FOREACH_END_DEL_ITEM(__ht, _idx, _p)
				}
				__ht.SetNNumUsed(_idx)
				break
			}
			for {
				var __ht *HashTable = EG.GetClassTable()
				var _idx uint32 = __ht.GetNNumUsed()
				var _p *Bucket = __ht.GetArData() + _idx
				var _z *Zval
				for _idx = __ht.GetNNumUsed(); _idx > 0; _idx-- {
					_p--
					_z = &_p.val

					if _z.GetType() == 0 {
						continue
					}
					key = _p.GetKey()
					zv = _z
					if _idx == EG.GetPersistentClassesCount() {
						break
					}
					DestroyZendClass(zv)
					ZendStringReleaseEx(key, 0)
					ZEND_HASH_FOREACH_END_DEL_ITEM(__ht, _idx, _p)
				}
				__ht.SetNNumUsed(_idx)
				break
			}
		}
		for EG.GetSymtableCachePtr() > EG.GetSymtableCache() {
			EG.GetSymtableCachePtr()--
			ZendHashDestroy(EG.symtable_cache_ptr)
			_efree(EG.symtable_cache_ptr)
		}
		ZendHashDestroy(&EG.included_files)
		ZendStackDestroy(&EG.user_error_handlers_error_reporting)
		ZendStackDestroy(&EG.user_error_handlers)
		ZendStackDestroy(&EG.user_exception_handlers)
		ZendObjectsStoreDestroy(&EG.objects_store)
		if EG.GetInAutoload() != nil {
			ZendHashDestroy(EG.GetInAutoload())
			_efree(EG.GetInAutoload())
		}
		if EG.GetHtIterators() != EG.GetHtIteratorsSlots() {
			_efree(EG.GetHtIterators())
		}
	}
	EG.SetHtIteratorsUsed(0)
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
	func_ = EG.GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case 2:

	case 1:
		var ce *ZendClassEntry = func_.GetScope()
		if space != nil {
			if ce != nil {
				*space = "::"
			} else {
				*space = ""
			}
		}
		if ce != nil {
			return ce.GetName().GetVal()
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
	func_ = EG.GetCurrentExecuteData().GetFunc()
	switch func_.GetType() {
	case 2:
		var function_name *ZendString = func_.GetFunctionName()
		if function_name != nil {
			return function_name.GetVal()
		} else {
			return "main"
		}
		break
	case 1:
		return func_.GetFunctionName().GetVal()
		break
	default:
		return nil
	}
}

/* }}} */

func ZendGetExecutedFilename() *byte {
	var ex *ZendExecuteData = EG.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetType()&1) != 0) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		return ex.GetFunc().GetOpArray().GetFilename().GetVal()
	} else {
		return "[no active file]"
	}
}

/* }}} */

func ZendGetExecutedFilenameEx() *ZendString {
	var ex *ZendExecuteData = EG.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetType()&1) != 0) {
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
	var ex *ZendExecuteData = EG.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetType()&1) != 0) {
		ex = ex.GetPrevExecuteData()
	}
	if ex != nil {
		if EG.GetException() != nil && ex.GetOpline().GetOpcode() == 149 && ex.GetOpline().GetLineno() == 0 && EG.GetOplineBeforeException() != nil {
			return EG.GetOplineBeforeException().GetLineno()
		}
		return ex.GetOpline().GetLineno()
	} else {
		return 0
	}
}

/* }}} */

func ZendGetExecutedScope() *ZendClassEntry {
	var ex *ZendExecuteData = EG.GetCurrentExecuteData()
	for true {
		if ex == nil {
			return nil
		} else if ex.GetFunc() != nil && ((ex.GetFunc().GetType()&1) == 0 || ex.GetFunc().GetScope() != nil) {
			return ex.GetFunc().GetScope()
		}
		ex = ex.GetPrevExecuteData()
	}
}

/* }}} */

func ZendIsExecuting() ZendBool { return EG.GetCurrentExecuteData() != 0 }

/* }}} */

func ZendUseUndefinedConstant(name *ZendString, attr ZendAstAttr, result *Zval) int {
	var colon *byte
	if EG.GetException() != nil {
		return FAILURE
	} else if g.Assign(&colon, (*byte)(ZendMemrchr(name.GetVal(), ':', name.GetLen()))) {
		ZendThrowError(nil, "Undefined class constant '%s'", name.GetVal())
		return FAILURE
	} else if (attr & 0x10) == 0 {
		ZendThrowError(nil, "Undefined constant '%s'", name.GetVal())
		return FAILURE
	} else {
		var actual *byte = name.GetVal()
		var actual_len int = name.GetLen()
		var slash *byte = (*byte)(ZendMemrchr(actual, '\\', actual_len))
		if slash != nil {
			actual = slash + 1
			actual_len -= actual - name.GetVal()
		}
		ZendError(1<<1, "Use of undefined constant %s - assumed '%s' (this will throw an Error in a future version of PHP)", actual, actual)
		if EG.GetException() != nil {
			return FAILURE
		} else {
			var result_str *ZendString = ZendStringInit(actual, actual_len, 0)
			ZvalPtrDtorNogc(result)
			var __z *Zval = result
			var __s *ZendString = result_str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	}
	return SUCCESS
}

/* }}} */

func ZvalUpdateConstantEx(p *Zval, scope *ZendClassEntry) int {
	if p.GetType() == 11 {
		var ast *ZendAst = (*ZendAst)((*byte)(p.GetValue().GetAst()) + g.SizeOf("zend_ast_ref"))
		if ast.GetKind() == ZEND_AST_CONSTANT {
			var name *ZendString = ZendAstGetConstantName(ast)
			var zv *Zval = ZendGetConstantEx(name, scope, ast.GetAttr())
			if zv == nil {
				return ZendUseUndefinedConstant(name, ast.GetAttr(), p)
			}
			ZvalPtrDtorNogc(p)
			var _z1 *Zval = p
			var _z2 *Zval = zv
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
					ZendGcAddref(&_gc.gc)
				} else {
					ZvalCopyCtorFunc(_z1)
				}
			}
		} else {
			var tmp Zval
			if ZendAstEvaluate(&tmp, ast, scope) != SUCCESS {
				return FAILURE
			}
			ZvalPtrDtorNogc(p)
			var _z1 *Zval = p
			var _z2 *Zval = &tmp
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	return SUCCESS
}

/* }}} */

func ZvalUpdateConstant(pp *Zval) int {
	return ZvalUpdateConstantEx(pp, g.CondF(EG.GetCurrentExecuteData() != nil, func() *ZendClassEntry { return ZendGetExecutedScope() }, func() *ZendClassEntry { return CG.GetActiveClassEntry() }))
}

/* }}} */

func _callUserFunctionEx(object *Zval, function_name *Zval, retval_ptr *Zval, param_count uint32, params []Zval, no_separation int) int {
	var fci ZendFcallInfo
	fci.SetSize(g.SizeOf("fci"))
	if object != nil {
		fci.SetObject(object.GetValue().GetObj())
	} else {
		fci.SetObject(nil)
	}
	var _z1 *Zval = &fci.function_name
	var _z2 *Zval = function_name
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
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
	fci.GetRetval().SetTypeInfo(0)
	if EG.GetActive() == 0 {
		return FAILURE
	}
	if EG.GetException() != nil {
		return FAILURE
	}
	r.Assert(fci.GetSize() == g.SizeOf("zend_fcall_info"))

	/* Initialize execute_data */

	if EG.GetCurrentExecuteData() == nil {

		/* This only happens when we're called outside any execute()'s
		 * It shouldn't be strictly necessary to NULL execute_data out,
		 * but it may make bugs easier to spot
		 */

		memset(&dummy_execute_data, 0, g.SizeOf("zend_execute_data"))
		EG.SetCurrentExecuteData(&dummy_execute_data)
	} else if EG.GetCurrentExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetFunc().GetCommonType()&1) == 0 && EG.GetCurrentExecuteData().GetOpline().GetOpcode() != 60 && EG.GetCurrentExecuteData().GetOpline().GetOpcode() != 129 && EG.GetCurrentExecuteData().GetOpline().GetOpcode() != 130 && EG.GetCurrentExecuteData().GetOpline().GetOpcode() != 131 {

		/* Insert fake frame in case of include or magic calls */

		dummy_execute_data = EG.current_execute_data
		dummy_execute_data.SetPrevExecuteData(EG.GetCurrentExecuteData())
		dummy_execute_data.SetCall(nil)
		dummy_execute_data.SetOpline(nil)
		dummy_execute_data.SetFunc(nil)
		EG.SetCurrentExecuteData(&dummy_execute_data)
	}
	if fci_cache == nil || fci_cache.GetFunctionHandler() == nil {
		var error *byte = nil
		if fci_cache == nil {
			fci_cache = &fci_cache_local
		}
		if ZendIsCallableEx(&fci.function_name, fci.GetObject(), 1<<3, nil, fci_cache, &error) == 0 {
			if error != nil {
				var callable_name *ZendString = ZendGetCallableNameEx(&fci.function_name, fci.GetObject())
				ZendError(1<<1, "Invalid callback %s, %s", callable_name.GetVal(), error)
				_efree(error)
				ZendStringReleaseEx(callable_name, 0)
			}
			if EG.GetCurrentExecuteData() == &dummy_execute_data {
				EG.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
			}
			return FAILURE
		} else if error != nil {

			/* Capitalize the first latter of the error message */

			if error[0] >= 'a' && error[0] <= 'z' {
				error[0] += 'A' - 'a'
			}
			ZendError(1<<13, "%s", error)
			_efree(error)
			if EG.GetException() != nil {
				if EG.GetCurrentExecuteData() == &dummy_execute_data {
					EG.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				}
				return FAILURE
			}
		}
	}
	func_ = fci_cache.GetFunctionHandler()
	if (func_.GetFnFlags()&1<<4) != 0 || fci_cache.GetObject() == nil {
		fci.SetObject(nil)
		object_or_called_scope = fci_cache.GetCalledScope()
		call_info = 1<<17 | 0<<16 | 1<<25
	} else {
		fci.SetObject(fci_cache.GetObject())
		object_or_called_scope = fci.GetObject()
		call_info = 1<<17 | 0<<16 | 1<<25 | (8 | 1<<0<<8 | 1<<1<<8)
	}
	call = ZendVmStackPushCallFrame(call_info, func_, fci.GetParamCount(), object_or_called_scope)
	if (func_.GetFnFlags() & 1 << 11) != 0 {
		ZendError(1<<13, "Function %s%s%s() is deprecated", g.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), g.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
		if EG.GetException() != nil {
			ZendVmStackFreeCallFrame(call)
			if EG.GetCurrentExecuteData() == &dummy_execute_data {
				EG.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
				ZendRethrowException(EG.GetCurrentExecuteData())
			}
			return FAILURE
		}
	}
	for i = 0; i < fci.GetParamCount(); i++ {
		var param *Zval
		var arg *Zval = &fci.params[i]
		var must_wrap ZendBool = 0
		if ZendCheckArgSendType(func_, i+1, 1|2) != 0 {
			if arg.GetType() != 10 {
				if fci.GetNoSeparation() == 0 {

					/* Separation is enabled -- create a ref */

					var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
					ZendGcSetRefcount(&_ref.gc, 1)
					_ref.GetGc().SetTypeInfo(10)
					var _z1 *Zval = &_ref.val
					var _z2 *Zval = arg
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					_ref.GetSources().SetPtr(nil)
					arg.GetValue().SetRef(_ref)
					arg.SetTypeInfo(10 | 1<<0<<8)

					/* Separation is enabled -- create a ref */

				} else if ZendCheckArgSendType(func_, i+1, 2) == 0 {

					/* By-value send is not allowed -- emit a warning,
					 * and perform the call with the value wrapped in a reference. */

					ZendError(1<<1, "Parameter %d to %s%s%s() expected to be a reference, value given", i+1, g.CondF1(func_.GetScope() != nil, func() []byte { return func_.GetScope().GetName().GetVal() }, ""), g.Cond(func_.GetScope() != nil, "::", ""), func_.GetFunctionName().GetVal())
					must_wrap = 1
					if EG.GetException() != nil {
						call.GetThis().SetNumArgs(i)
						ZendVmStackFreeArgs(call)
						ZendVmStackFreeCallFrame(call)
						if EG.GetCurrentExecuteData() == &dummy_execute_data {
							EG.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
						}
						return FAILURE
					}
				}
			}
		} else {
			if arg.GetType() == 10 && (func_.GetFnFlags()&1<<18) == 0 {

				/* don't separate references for __call */

				arg = &(*arg).value.GetRef().GetVal()

				/* don't separate references for __call */

			}
		}
		param = (*Zval)(call) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(i+1)-1))
		if must_wrap == 0 {
			var _z1 *Zval = param
			var _z2 *Zval = arg
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			if arg.GetTypeFlags() != 0 {
				ZvalAddrefP(arg)
			}
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = arg
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			param.GetValue().SetRef(_ref)
			param.SetTypeInfo(10 | 1<<0<<8)
		}
	}
	if (func_.GetOpArray().GetFnFlags() & 1 << 20) != 0 {
		var call_info uint32
		ZendGcAddref(&((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object")))).gc)
		call_info = 1 << 22
		if (func_.GetFnFlags() & 1 << 21) != 0 {
			call_info |= 1 << 23
		}
		call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | call_info)
	}
	if func_.GetType() == 2 {
		var call_via_handler int = (func_.GetFnFlags() & 1 << 18) != 0
		var current_opline_before_exception *ZendOp = EG.GetOplineBeforeException()
		ZendInitFuncExecuteData(call, &func_.op_array, fci.GetRetval())
		ZendExecuteEx(call)
		EG.SetOplineBeforeException(current_opline_before_exception)
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else if func_.GetType() == 1 {
		var call_via_handler int = (func_.GetFnFlags() & 1 << 18) != 0
		fci.GetRetval().SetTypeInfo(1)
		call.SetPrevExecuteData(EG.GetCurrentExecuteData())
		EG.SetCurrentExecuteData(call)
		if ZendExecuteInternal == nil {

			/* saves one function call if zend_execute_internal is not used */

			func_.GetInternalFunction().GetHandler()(call, fci.GetRetval())

			/* saves one function call if zend_execute_internal is not used */

		} else {
			ZendExecuteInternal(call, fci.GetRetval())
		}
		EG.SetCurrentExecuteData(call.GetPrevExecuteData())
		ZendVmStackFreeArgs(call)
		if EG.GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetTypeInfo(0)
		}
		if call_via_handler != 0 {

			/* We must re-initialize function again */

			fci_cache.SetFunctionHandler(nil)

			/* We must re-initialize function again */

		}
	} else {
		fci.GetRetval().SetTypeInfo(1)

		/* Not sure what should be done here if it's a static method */

		if fci.GetObject() != nil {
			call.SetPrevExecuteData(EG.GetCurrentExecuteData())
			EG.SetCurrentExecuteData(call)
			fci.GetObject().GetHandlers().GetCallMethod()(func_.GetFunctionName(), fci.GetObject(), call, fci.GetRetval())
			EG.SetCurrentExecuteData(call.GetPrevExecuteData())
		} else {
			ZendThrowError(nil, "Cannot call overloaded function for non-object")
		}
		ZendVmStackFreeArgs(call)
		if func_.GetType() == 5 {
			ZendStringReleaseEx(func_.GetFunctionName(), 0)
		}
		_efree(func_)
		if EG.GetException() != nil {
			ZvalPtrDtor(fci.GetRetval())
			fci.GetRetval().SetTypeInfo(0)
		}
	}
	ZendVmStackFreeCallFrame(call)
	if EG.GetCurrentExecuteData() == &dummy_execute_data {
		EG.SetCurrentExecuteData(dummy_execute_data.GetPrevExecuteData())
	}
	if EG.GetException() != nil {
		if EG.GetCurrentExecuteData() == nil {
			ZendThrowExceptionInternal(nil)
		} else if EG.GetCurrentExecuteData().GetFunc() != nil && (EG.GetCurrentExecuteData().GetFunc().GetCommonType()&1) == 0 {
			ZendRethrowException(EG.GetCurrentExecuteData())
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
		if name == nil || name.GetLen() == 0 {
			return nil
		}
		if name.GetVal()[0] == '\\' {
			lc_name = ZendStringAlloc(name.GetLen()-1, 0)
			ZendStrTolowerCopy(lc_name.GetVal(), name.GetVal()+1, name.GetLen()-1)
		} else {
			lc_name = ZendStringTolowerEx(name, 0)
		}
	}
	zv = ZendHashFind(EG.GetClassTable(), lc_name)
	if zv != nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		ce = (*ZendClassEntry)(zv.GetValue().GetPtr())
		if (ce.GetCeFlags() & 1 << 3) == 0 {
			if (flags&0x400) != 0 || (flags&0x800) != 0 && (ce.GetCeFlags()&1<<22) != 0 {
				ce.SetCeFlags(ce.GetCeFlags() | 1<<23)
				return ce
			}
			return nil
		}
		return ce
	}

	/* The compiler is not-reentrant. Make sure we __autoload() only during run-time
	 * (doesn't impact functionality of __autoload()
	 */

	if (flags&0x80) != 0 || ZendIsCompiling() != 0 {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	if EG.GetAutoloadFunc() == nil {
		var func_ *ZendFunction = ZendFetchFunction(ZendKnownStrings[ZEND_STR_MAGIC_AUTOLOAD])
		if func_ != nil {
			EG.SetAutoloadFunc(func_)
		} else {
			if key == nil {
				ZendStringReleaseEx(lc_name, 0)
			}
			return nil
		}
	}

	/* Verify class name before passing it to __autoload() */

	if key == nil && strspn(name.GetVal(), "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377\\") != name.GetLen() {
		ZendStringReleaseEx(lc_name, 0)
		return nil
	}
	if EG.GetInAutoload() == nil {
		EG.SetInAutoload((*HashTable)(_emalloc(g.SizeOf("HashTable"))))
		_zendHashInit(EG.GetInAutoload(), 8, nil, 0)
	}
	if ZendHashAddEmptyElement(EG.GetInAutoload(), lc_name) == nil {
		if key == nil {
			ZendStringReleaseEx(lc_name, 0)
		}
		return nil
	}
	&local_retval.SetTypeInfo(0)
	if name.GetVal()[0] == '\\' {
		var __z *Zval = &args[0]
		var __s *ZendString = ZendStringInit(name.GetVal()+1, name.GetLen()-1, 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	} else {
		var __z *Zval = &args[0]
		var __s *ZendString = name
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	}
	fcall_info.SetSize(g.SizeOf("fcall_info"))
	var __z *Zval = &fcall_info.function_name
	var __s *ZendString = EG.GetAutoloadFunc().GetFunctionName()
	__z.GetValue().SetStr(__s)
	if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
		__z.SetTypeInfo(6)
	} else {
		ZendGcAddref(&__s.gc)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	fcall_info.SetRetval(&local_retval)
	fcall_info.SetParamCount(1)
	fcall_info.SetParams(args)
	fcall_info.SetObject(nil)
	fcall_info.SetNoSeparation(1)
	fcall_cache.SetFunctionHandler(EG.GetAutoloadFunc())
	fcall_cache.SetCalledScope(nil)
	fcall_cache.SetObject(nil)
	orig_fake_scope = EG.GetFakeScope()
	EG.SetFakeScope(nil)
	ZendExceptionSave()
	if ZendCallFunction(&fcall_info, &fcall_cache) == SUCCESS && EG.GetException() == nil {
		ce = ZendHashFindPtr(EG.GetClassTable(), lc_name)
	}
	ZendExceptionRestore()
	EG.SetFakeScope(orig_fake_scope)
	ZvalPtrDtor(&args[0])
	ZvalPtrDtorStr(&fcall_info.function_name)
	ZendHashDel(EG.GetInAutoload(), lc_name)
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
		if ex.GetThis().GetType() == 8 {
			return ex.GetThis().GetValue().GetObj().GetCe()
		} else if ex.GetThis().GetValue().GetCe() != nil {
			return ex.GetThis().GetValue().GetCe()
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != 1 || ex.GetFunc().GetScope() != nil {
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
		if ex.GetThis().GetType() == 8 {
			return ex.GetThis().GetValue().GetObj()
		} else if ex.GetFunc() != nil {
			if ex.GetFunc().GetType() != 1 || ex.GetFunc().GetScope() != nil {
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
		var __z *Zval = &pv
		var __s *ZendString = ZendStringAlloc(str_len+g.SizeOf("\"return ;\"")-1, 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		memcpy(pv.GetValue().GetStr().GetVal(), "return ", g.SizeOf("\"return \"")-1)
		memcpy(pv.GetValue().GetStr().GetVal()+g.SizeOf("\"return \"")-1, str, str_len)
		pv.GetValue().GetStr().GetVal()[pv.GetValue().GetStr().GetLen()-1] = ';'
		pv.GetValue().GetStr().GetVal()[pv.GetValue().GetStr().GetLen()] = '0'
	} else {
		var __z *Zval = &pv
		var __s *ZendString = ZendStringInit(str, str_len, 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}

	/*printf("Evaluating '%s'\n", pv.value.str.val);*/

	original_compiler_options = CG.GetCompilerOptions()
	CG.SetCompilerOptions(0)
	new_op_array = ZendCompileString(&pv, string_name)
	CG.SetCompilerOptions(original_compiler_options)
	if new_op_array != nil {
		var local_retval Zval
		EG.SetNoExtensions(1)
		new_op_array.SetScope(ZendGetExecutedScope())
		var __orig_bailout *sigjmp_buf = EG.GetBailout()
		var __bailout sigjmp_buf
		EG.SetBailout(&__bailout)
		if sigsetjmp(__bailout, 0) == 0 {
			&local_retval.SetTypeInfo(0)
			ZendExecute(new_op_array, &local_retval)
		} else {
			EG.SetBailout(__orig_bailout)
			DestroyOpArray(new_op_array)
			_efree(new_op_array)
			_zendBailout(__FILE__, __LINE__)
		}
		EG.SetBailout(__orig_bailout)
		if local_retval.GetType() != 0 {
			if retval_ptr != nil {
				var _z1 *Zval = retval_ptr
				var _z2 *Zval = &local_retval
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				ZvalPtrDtor(&local_retval)
			}
		} else {
			if retval_ptr != nil {
				retval_ptr.SetTypeInfo(1)
			}
		}
		EG.SetNoExtensions(0)
		DestroyOpArray(new_op_array)
		_efree(new_op_array)
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
	if handle_exceptions != 0 && EG.GetException() != nil {
		ZendExceptionError(EG.GetException(), 1<<0)
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
	EG.SetTimedOut(0)
	ZendSetTimeoutEx(0, 1)
	ZendErrorNoreturn(1<<0, "Maximum execution time of "+"%"+"lld"+" second%s exceeded", EG.GetTimeoutSeconds(), g.Cond(EG.GetTimeoutSeconds() == 1, "", "s"))
}

/* }}} */

func ZendTimeoutHandler(dummy int) {
	if EG.GetTimedOut() != 0 {

		/* Die on hard timeout */

		var error_filename *byte = nil
		var error_lineno uint32 = 0
		var log_buffer []byte
		var output_len int = 0
		if ZendIsCompiling() != 0 {
			error_filename = ZendGetCompiledFilename().GetVal()
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
		output_len = snprintf(log_buffer, g.SizeOf("log_buffer"), "\nFatal error: Maximum execution time of "+"%"+"lld"+"+"+"%"+"lld"+" seconds exceeded (terminated) in %s on line %d\n", EG.GetTimeoutSeconds(), EG.GetHardTimeout(), error_filename, error_lineno)
		if output_len > 0 {
			void(write(2, log_buffer, g.CondF2(output_len < g.SizeOf("log_buffer"), output_len, func() __auto__ { return g.SizeOf("log_buffer") })))
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

		ZendSignalGlobals.SetRunning(0)
		ZendOnTimeout(EG.GetTimeoutSeconds())
	}
	EG.SetTimedOut(1)
	EG.SetVmInterrupt(1)
	if EG.GetHardTimeout() > 0 {

		/* Set hard timeout */

		ZendSetTimeoutEx(EG.GetHardTimeout(), 1)

		/* Set hard timeout */

	}
}

/* }}} */

/* This one doesn't exists on QNX */

// #define SIGPROF       27

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
	signo = 27
	if reset_signals != 0 {
		ZendSignal(signo, ZendTimeoutHandler)
	}
}

/* }}} */

func ZendSetTimeout(seconds ZendLong, reset_signals int) {
	EG.SetTimeoutSeconds(seconds)
	ZendSetTimeoutEx(seconds, reset_signals)
	EG.SetTimedOut(0)
}

/* }}} */

func ZendUnsetTimeout() {
	if EG.GetTimeoutSeconds() != 0 {
		var no_timeout __struct__itimerval
		no_timeout.it_interval.tv_usec = 0
		no_timeout.it_interval.tv_sec = no_timeout.it_interval.tv_usec
		no_timeout.it_value.tv_usec = no_timeout.it_interval.tv_sec
		no_timeout.it_value.tv_sec = no_timeout.it_value.tv_usec
		setitimer(ITIMER_PROF, &no_timeout, nil)
	}
	EG.SetTimedOut(0)
}

/* }}} */

func ZendFetchClass(class_name *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var fetch_sub_type int = fetch_type & 0xf
check_fetch_type:
	switch fetch_sub_type {
	case 1:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access self:: when no class scope is active")
		}
		return scope
	case 2:
		scope = ZendGetExecutedScope()
		if scope == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when no class scope is active")
			return nil
		}
		if !(scope.parent) {
			ZendThrowOrError(fetch_type, nil, "Cannot access parent:: when current class scope has no parent")
		}
		return scope.parent
	case 3:
		ce = ZendGetCalledScope(EG.GetCurrentExecuteData())
		if ce == nil {
			ZendThrowOrError(fetch_type, nil, "Cannot access static:: when no class scope is active")
			return nil
		}
		return ce
	case 4:
		fetch_sub_type = ZendGetClassFetchType(class_name)
		if fetch_sub_type != 0 {
			goto check_fetch_type
		}
		break
	}
	if (fetch_type & 0x80) != 0 {
		return ZendLookupClassEx(class_name, nil, fetch_type)
	} else if g.Assign(&ce, ZendLookupClassEx(class_name, nil, fetch_type)) == nil {
		if (fetch_type&0x100) == 0 && EG.GetException() == nil {
			if fetch_sub_type == 5 {
				ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", class_name.GetVal())
			} else if fetch_sub_type == 6 {
				ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", class_name.GetVal())
			} else {
				ZendThrowOrError(fetch_type, nil, "Class '%s' not found", class_name.GetVal())
			}
		}
		return nil
	}
	return ce
}

/* }}} */

func ZendFetchClassByName(class_name *ZendString, key *ZendString, fetch_type int) *ZendClassEntry {
	var ce *ZendClassEntry
	if (fetch_type & 0x80) != 0 {
		return ZendLookupClassEx(class_name, key, fetch_type)
	} else if g.Assign(&ce, ZendLookupClassEx(class_name, key, fetch_type)) == nil {
		if (fetch_type & 0x100) != 0 {
			return nil
		}
		if EG.GetException() != nil {
			if (fetch_type & 0x200) == 0 {
				var exception_str *ZendString
				var exception_zv Zval
				var __z *Zval = &exception_zv
				__z.GetValue().SetObj(EG.GetException())
				__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
				ZvalAddrefP(&exception_zv)
				ZendClearException()
				exception_str = ZvalGetString(&exception_zv)
				ZendErrorNoreturn(1<<0, "During class fetch: Uncaught %s", exception_str.GetVal())
			}
			return nil
		}
		if (fetch_type & 0xf) == 5 {
			ZendThrowOrError(fetch_type, nil, "Interface '%s' not found", class_name.GetVal())
		} else if (fetch_type & 0xf) == 6 {
			ZendThrowOrError(fetch_type, nil, "Trait '%s' not found", class_name.GetVal())
		} else {
			ZendThrowOrError(fetch_type, nil, "Class '%s' not found", class_name.GetVal())
		}
		return nil
	}
	return ce
}

/* }}} */

func ZendDeleteGlobalVariable(name *ZendString) int {
	return ZendHashDelInd(&EG.symbol_table, name)
}

/* }}} */

func ZendRebuildSymbolTable() *ZendArray {
	var ex *ZendExecuteData
	var symbol_table *ZendArray

	/* Search for last called user function */

	ex = EG.GetCurrentExecuteData()
	for ex != nil && (ex.GetFunc() == nil || (ex.GetFunc().GetCommonType()&1) != 0) {
		ex = ex.GetPrevExecuteData()
	}
	if ex == nil {
		return nil
	}
	if (ex.GetThis().GetTypeInfo() & 1 << 20) != 0 {
		return ex.GetSymbolTable()
	}
	ex.GetThis().SetTypeInfo(ex.GetThis().GetTypeInfo() | 1<<20)
	if EG.GetSymtableCachePtr() > EG.GetSymtableCache() {
		ex.SetSymbolTable(*(g.PreDec(&(EG.GetSymtableCachePtr()))))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		ZendHashExtend(symbol_table, ex.GetFunc().GetOpArray().GetLastVar(), 0)
	} else {
		ex.SetSymbolTable(_zendNewArray(ex.GetFunc().GetOpArray().GetLastVar()))
		symbol_table = ex.GetSymbolTable()
		if ex.GetFunc().GetOpArray().GetLastVar() == 0 {
			return symbol_table
		}
		ZendHashRealInitMixed(symbol_table)
	}
	if ex.GetFunc().GetOpArray().GetLastVar() != 0 {
		var str **ZendString = ex.GetFunc().GetOpArray().GetVars()
		var end **ZendString = str + ex.GetFunc().GetOpArray().GetLastVar()
		var var_ *Zval = (*Zval)(ex) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(0))
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

	if op_array.GetLastVar() != 0 {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(0))
		for {
			var zv *Zval = ZendHashFindEx(ht, *str, 1)
			if zv != nil {
				if zv.GetType() == 13 {
					var val *Zval = zv.GetValue().GetZv()
					var _z1 *Zval = var_
					var _z2 *Zval = val
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				} else {
					var _z1 *Zval = var_
					var _z2 *Zval = zv
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				}
			} else {
				var_.SetTypeInfo(0)
				zv = ZendHashAddNew(ht, *str, var_)
			}
			zv.GetValue().SetZv(var_)
			zv.SetTypeInfo(13)
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

	if op_array.GetLastVar() != 0 {
		var str **ZendString = op_array.GetVars()
		var end **ZendString = str + op_array.GetLastVar()
		var var_ *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(0))
		for {
			if var_.GetType() == 0 {
				ZendHashDel(ht, *str)
			} else {
				ZendHashUpdate(ht, *str, var_)
				var_.SetTypeInfo(0)
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
	var execute_data *ZendExecuteData = EG.GetCurrentExecuteData()
	for execute_data != nil && (execute_data.GetFunc() == nil || (execute_data.GetFunc().GetCommonType()&1) != 0) {
		execute_data = execute_data.GetPrevExecuteData()
	}
	if execute_data != nil {
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) == 0 {
			var h ZendUlong = ZendStringHashVal(name)
			var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
			if op_array.GetLastVar() != 0 {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if (*str).GetH() == h && ZendStringEqualContent(*str, name) != 0 {
						var var_ *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(str-op_array.GetVars()))
						var _z1 *Zval = var_
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
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
	var execute_data *ZendExecuteData = EG.GetCurrentExecuteData()
	for execute_data != nil && (execute_data.GetFunc() == nil || (execute_data.GetFunc().GetCommonType()&1) != 0) {
		execute_data = execute_data.GetPrevExecuteData()
	}
	if execute_data != nil {
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) == 0 {
			var h ZendUlong = ZendHashFunc(name, len_)
			var op_array *ZendOpArray = &execute_data.func_.GetOpArray()
			if op_array.GetLastVar() != 0 {
				var str **ZendString = op_array.GetVars()
				var end **ZendString = str + op_array.GetLastVar()
				for {
					if (*str).GetH() == h && (*str).GetLen() == len_ && memcmp((*str).GetVal(), name, len_) == 0 {
						var var_ *Zval = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(str-op_array.GetVars()))
						ZvalPtrDtor(var_)
						var _z1 *Zval = var_
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
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
