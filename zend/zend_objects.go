// <<generate>>

package zend

import g "sik/runtime/grammar"

// Source: <Zend/zend_objects.h>

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
   +----------------------------------------------------------------------+
*/

// #define ZEND_OBJECTS_H

// # include "zend.h"

// Source: <Zend/zend_objects.c>

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

// # include "zend.h"

// # include "zend_globals.h"

// # include "zend_variables.h"

// # include "zend_API.h"

// # include "zend_interfaces.h"

// # include "zend_exceptions.h"

// # include "zend_weakrefs.h"

func _zendObjectStdInit(object *ZendObject, ce *ZendClassEntry) {
	ZendGcSetRefcount(&object.gc, 1)
	object.GetGc().SetTypeInfo(8 | 1<<4<<0)
	object.SetCe(ce)
	object.SetProperties(nil)
	ZendObjectsStorePut(object)
	if (ce.GetCeFlags() & 1 << 11) != 0 {
		(object.GetPropertiesTable() + object.GetCe().GetDefaultPropertiesCount()).u1.type_info = 0
	}
}
func ZendObjectStdInit(object *ZendObject, ce *ZendClassEntry) { _zendObjectStdInit(object, ce) }
func ZendObjectStdDtor(object *ZendObject) {
	var p *Zval
	var end *Zval
	if object.GetProperties() != nil {
		if (ZvalGcFlags(object.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			if ZendGcDelref(&(object.GetProperties()).gc) == 0 && ZvalGcType(object.GetProperties().GetGc().GetTypeInfo()) != 1 {
				ZendArrayDestroy(object.GetProperties())
			}
		}
	}
	p = object.GetPropertiesTable()
	if object.GetCe().GetDefaultPropertiesCount() != 0 {
		end = p + object.GetCe().GetDefaultPropertiesCount()
		for {
			if p.GetTypeFlags() != 0 {
				if p.GetType() == 10 && p.GetValue().GetRef().GetSources().GetPtr() != nil {
					var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(object, p)
					if prop_info.GetType() != 0 {
						ZendRefDelTypeSource(&(p.GetValue().GetRef()).sources, prop_info)
					}
				}
				IZvalPtrDtor(p)
			}
			p++
			if p == end {
				break
			}
		}
	}
	if (object.GetCe().GetCeFlags() & 1 << 11) != 0 {
		if p.GetType() == 6 {
			ZvalPtrDtorStr(p)
		} else if p.GetType() == 7 {
			var guards *HashTable
			guards = p.GetValue().GetArr()
			assert(guards != nil)
			ZendHashDestroy(guards)
			_efree(guards)
		}
	}
	if (ZvalGcFlags(object.GetGc().GetTypeInfo()) & 1 << 7) != 0 {
		ZendWeakrefsNotify(object)
	}
}
func ZendObjectsDestroyObject(object *ZendObject) {
	var destructor *ZendFunction = object.GetCe().GetDestructor()
	if destructor != nil {
		var old_exception *ZendObject
		var orig_fake_scope *ZendClassEntry
		var fci ZendFcallInfo
		var fcic ZendFcallInfoCache
		var ret Zval
		if (destructor.GetOpArray().GetFnFlags() & (1<<2 | 1<<1)) != 0 {
			if (destructor.GetOpArray().GetFnFlags() & 1 << 2) != 0 {

				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */

				if EG.GetCurrentExecuteData() != nil {
					var scope *ZendClassEntry = ZendGetExecutedScope()
					if object.GetCe() != scope {
						ZendThrowError(nil, "Call to private %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					ZendError(1<<1, "Call to private %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}

				/* Ensure that if we're calling a private function, we're allowed to do so.
				 */

			} else {

				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */

				if EG.GetCurrentExecuteData() != nil {
					var scope *ZendClassEntry = ZendGetExecutedScope()
					if ZendCheckProtected(g.CondF(destructor.GetPrototype() != nil, func() *ZendClassEntry { return destructor.GetPrototype().GetScope() }, func() *ZendClassEntry { return destructor.GetScope() }), scope) == 0 {
						ZendThrowError(nil, "Call to protected %s::__destruct() from context '%s'", object.GetCe().GetName().GetVal(), g.CondF1(scope != nil, func() []byte { return scope.GetName().GetVal() }, ""))
						return
					}
				} else {
					ZendError(1<<1, "Call to protected %s::__destruct() from context '' during shutdown ignored", object.GetCe().GetName().GetVal())
					return
				}

				/* Ensure that if we're calling a protected function, we're allowed to do so.
				 */

			}
		}
		ZendGcAddref(&object.gc)

		/* Make sure that destructors are protected from previously thrown exceptions.
		 * For example, if an exception was thrown in a function and when the function's
		 * local variable destruction results in a destructor being called.
		 */

		old_exception = nil
		if EG.GetException() != nil {
			if EG.GetException() == object {
				ZendErrorNoreturn(1<<4, "Attempt to destruct pending exception")
			} else {
				old_exception = EG.GetException()
				EG.SetException(nil)
			}
		}
		orig_fake_scope = EG.GetFakeScope()
		EG.SetFakeScope(nil)
		&ret.SetTypeInfo(0)
		fci.SetSize(g.SizeOf("fci"))
		fci.SetObject(object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		&fci.function_name.u1.type_info = 0
		fcic.SetFunctionHandler(destructor)
		fcic.SetCalledScope(object.GetCe())
		fcic.SetObject(object)
		ZendCallFunction(&fci, &fcic)
		ZvalPtrDtor(&ret)
		if old_exception != nil {
			if EG.GetException() != nil {
				ZendExceptionSetPrevious(EG.GetException(), old_exception)
			} else {
				EG.SetException(old_exception)
			}
		}
		ZendObjectRelease(object)
		EG.SetFakeScope(orig_fake_scope)
	}
}
func ZendObjectsNew(ce *ZendClassEntry) *ZendObject {
	var object *ZendObject = _emalloc(g.SizeOf("zend_object") + ZendObjectPropertiesSize(ce))
	_zendObjectStdInit(object, ce)
	object.SetHandlers(&StdObjectHandlers)
	return object
}
func ZendObjectsCloneMembers(new_object *ZendObject, old_object *ZendObject) {
	if old_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var src *Zval = old_object.GetPropertiesTable()
		var dst *Zval = new_object.GetPropertiesTable()
		var end *Zval = src + old_object.GetCe().GetDefaultPropertiesCount()
		for {
			IZvalPtrDtor(dst)
			*dst = *src
			ZvalAddRef(dst)
			if dst.GetType() == 10 && dst.GetValue().GetRef().GetSources().GetPtr() != nil {
				var prop_info *ZendPropertyInfo = ZendGetPropertyInfoForSlot(new_object, dst)
				if prop_info.GetType() != 0 {
					ZendRefAddTypeSource(&(dst.GetValue().GetRef()).sources, prop_info)
				}
			}
			src++
			dst++
			if src == end {
				break
			}
		}
	} else if old_object.GetProperties() != nil && old_object.GetCe().GetClone() == nil {

		/* fast copy */

		if old_object.GetHandlers() == &StdObjectHandlers {
			if (ZvalGcFlags(old_object.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
				ZendGcAddref(&(old_object.GetProperties()).gc)
			}
			new_object.SetProperties(old_object.GetProperties())
			return
		}

		/* fast copy */

	}
	if old_object.GetProperties() != nil && old_object.GetProperties().GetNNumOfElements() != 0 {
		var prop *Zval
		var new_prop Zval
		var num_key ZendUlong
		var key *ZendString
		if new_object.GetProperties() == nil {
			new_object.SetProperties(_zendNewArray(old_object.GetProperties().GetNNumOfElements()))
			ZendHashRealInitMixed(new_object.GetProperties())
		} else {
			ZendHashExtend(new_object.GetProperties(), new_object.GetProperties().GetNNumUsed()+old_object.GetProperties().GetNNumOfElements(), 0)
		}
		new_object.GetProperties().SetUFlags(new_object.GetProperties().GetUFlags() | old_object.GetProperties().GetUFlags()&1<<5)
		for {
			var __ht *HashTable = old_object.GetProperties()
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				num_key = _p.GetH()
				key = _p.GetKey()
				prop = _z
				if prop.GetType() == 13 {
					&new_prop.GetValue().SetZv(new_object.GetPropertiesTable() + (prop.GetValue().GetZv() - old_object.GetPropertiesTable()))
					&new_prop.SetTypeInfo(13)
				} else {
					var _z1 *Zval = &new_prop
					var _z2 *Zval = prop
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					ZvalAddRef(&new_prop)
				}
				if key != nil {
					_zendHashAppend(new_object.GetProperties(), key, &new_prop)
				} else {
					ZendHashIndexAddNew(new_object.GetProperties(), num_key, &new_prop)
				}
			}
			break
		}
	}
	if old_object.GetCe().GetClone() != nil {
		var fci ZendFcallInfo
		var fcic ZendFcallInfoCache
		var ret Zval
		ZendGcAddref(&new_object.gc)
		&ret.SetTypeInfo(0)
		fci.SetSize(g.SizeOf("fci"))
		fci.SetObject(new_object)
		fci.SetRetval(&ret)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		&fci.function_name.u1.type_info = 0
		fcic.SetFunctionHandler(new_object.GetCe().GetClone())
		fcic.SetCalledScope(new_object.GetCe())
		fcic.SetObject(new_object)
		ZendCallFunction(&fci, &fcic)
		ZvalPtrDtor(&ret)
		ZendObjectRelease(new_object)
	}
}
func ZendObjectsCloneObj(zobject *Zval) *ZendObject {
	var old_object *ZendObject
	var new_object *ZendObject

	/* assume that create isn't overwritten, so when clone depends on the
	 * overwritten one then it must itself be overwritten */

	old_object = zobject.GetValue().GetObj()
	new_object = ZendObjectsNew(old_object.GetCe())

	/* zend_objects_clone_members() expect the properties to be initialized. */

	if new_object.GetCe().GetDefaultPropertiesCount() != 0 {
		var p *Zval = new_object.GetPropertiesTable()
		var end *Zval = p + new_object.GetCe().GetDefaultPropertiesCount()
		for {
			p.SetTypeInfo(0)
			p++
			if p == end {
				break
			}
		}
	}
	ZendObjectsCloneMembers(new_object, old_object)
	return new_object
}
