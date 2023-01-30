// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func PhpVarUnserializeInit() PhpUnserializeDataT {
	var d PhpUnserializeDataT

	/* fprintf(stderr, "UNSERIALIZE_INIT    == lock: %u, level: %u\n", BG(serialize_lock), BG(unserialize).level); */

	if BG(serialize_lock) || !(BG(unserialize).level) {
		d = zend.Emalloc(b.SizeOf("struct php_unserialize_data"))
		d.SetLast(d.GetEntries())
		d.SetLastDtor(nil)
		d.SetFirstDtor(d.GetLastDtor())
		d.SetAllowedClasses(nil)
		d.SetRefProps(nil)
		d.SetCurDepth(0)
		d.SetMaxDepth(BG(unserialize_max_depth))
		d.GetEntries().SetUsedSlots(0)
		d.GetEntries().SetNext(nil)
		if !(BG(serialize_lock)) {
			BG(unserialize).data = d
			BG(unserialize).level = 1
		}
	} else {
		d = BG(unserialize).data
		BG(unserialize).level++
	}
	return d
}
func PhpVarUnserializeDestroy(d PhpUnserializeDataT) {
	/* fprintf(stderr, "UNSERIALIZE_DESTROY == lock: %u, level: %u\n", BG(serialize_lock), BG(unserialize).level); */

	if BG(serialize_lock) || BG(unserialize).level == 1 {
		VarDestroy(&d)
		zend.Efree(d)
	}
	if !(BG(serialize_lock)) && !(b.PreDec(&(BG(unserialize).level))) {
		BG(unserialize).data = nil
	}
}
func PhpVarUnserializeGetAllowedClasses(d PhpUnserializeDataT) *zend.HashTable {
	return d.GetAllowedClasses()
}
func PhpVarUnserializeSetAllowedClasses(d PhpUnserializeDataT, classes *zend.HashTable) {
	d.SetAllowedClasses(classes)
}
func PhpVarUnserializeSetMaxDepth(d PhpUnserializeDataT, max_depth zend.ZendLong) {
	d.SetMaxDepth(max_depth)
}
func PhpVarUnserializeGetMaxDepth(d PhpUnserializeDataT) zend.ZendLong { return d.GetMaxDepth() }
func PhpVarUnserializeSetCurDepth(d PhpUnserializeDataT, cur_depth zend.ZendLong) {
	d.SetCurDepth(cur_depth)
}
func PhpVarUnserializeGetCurDepth(d PhpUnserializeDataT) zend.ZendLong { return d.GetCurDepth() }
func VarPush(var_hashx *PhpUnserializeDataT, rval *zend.Zval) {
	var var_hash *VarEntries = var_hashx.GetLast()
	if var_hash.GetUsedSlots() == VAR_ENTRIES_MAX {
		var_hash = zend.Emalloc(b.SizeOf("var_entries"))
		var_hash.SetUsedSlots(0)
		var_hash.SetNext(0)
		var_hashx.GetLast().SetNext(var_hash)
		var_hashx.SetLast(var_hash)
	}
	var_hash.GetData()[b.PostInc(&(var_hash.GetUsedSlots()))] = rval
}
func VarPushDtor(var_hashx *PhpUnserializeDataT, rval *zend.Zval) {
	if zend.Z_REFCOUNTED_P(rval) {
		var tmp_var *zend.Zval = VarTmpVar(var_hashx)
		if tmp_var == nil {
			return
		}
		zend.ZVAL_COPY(tmp_var, rval)
	}
}
func TmpVar(var_hashx *PhpUnserializeDataT, num zend.ZendLong) *zend.Zval {
	var var_hash *VarDtorEntries
	var used_slots zend.ZendLong
	if var_hashx == nil || (*var_hashx) == nil || num < 1 {
		return nil
	}
	var_hash = var_hashx.GetLastDtor()
	if var_hash == nil || var_hash.GetUsedSlots()+num > VAR_DTOR_ENTRIES_MAX {
		var_hash = zend.Emalloc(b.SizeOf("var_dtor_entries"))
		var_hash.SetUsedSlots(0)
		var_hash.SetNext(0)
		if var_hashx.GetFirstDtor() == nil {
			var_hashx.SetFirstDtor(var_hash)
		} else {
			var_hashx.GetLastDtor().SetNext(var_hash)
		}
		var_hashx.SetLastDtor(var_hash)
	}
	for used_slots = var_hash.GetUsedSlots(); var_hash.GetUsedSlots() < used_slots+num; var_hash.GetUsedSlots()++ {
		zend.ZVAL_UNDEF(var_hash.GetData()[var_hash.GetUsedSlots()])
		var_hash.GetData()[var_hash.GetUsedSlots()].GetU2Extra() = 0
	}
	return var_hash.GetData()[used_slots]
}
func VarTmpVar(var_hashx *PhpUnserializeDataT) *zend.Zval { return TmpVar(var_hashx, 1) }
func VarReplace(var_hashx *PhpUnserializeDataT, ozval *zend.Zval, nzval *zend.Zval) {
	var i zend.ZendLong
	var var_hash *VarEntries = var_hashx.GetEntries()
	for var_hash != nil {
		for i = 0; i < var_hash.GetUsedSlots(); i++ {
			if var_hash.GetData()[i] == ozval {
				var_hash.GetData()[i] = nzval
			}
		}
		var_hash = var_hash.GetNext()
	}
}
func VarAccess(var_hashx *PhpUnserializeDataT, id zend.ZendLong) *zend.Zval {
	var var_hash *VarEntries = var_hashx.GetEntries()
	for id >= VAR_ENTRIES_MAX && var_hash != nil && var_hash.GetUsedSlots() == VAR_ENTRIES_MAX {
		var_hash = var_hash.GetNext()
		id -= VAR_ENTRIES_MAX
	}
	if var_hash == nil {
		return nil
	}
	if id < 0 || id >= var_hash.GetUsedSlots() {
		return nil
	}
	return var_hash.GetData()[id]
}
func VarDestroy(var_hashx *PhpUnserializeDataT) {
	var next any
	var i zend.ZendLong
	var var_hash *VarEntries = var_hashx.GetEntries().GetNext()
	var var_dtor_hash *VarDtorEntries = var_hashx.GetFirstDtor()
	var delayed_call_failed zend.ZendBool = 0
	var wakeup_name zend.Zval
	var unserialize_name zend.Zval
	zend.ZVAL_UNDEF(&wakeup_name)
	zend.ZVAL_UNDEF(&unserialize_name)
	for var_hash != nil {
		next = var_hash.GetNext()
		zend.EfreeSize(var_hash, b.SizeOf("var_entries"))
		var_hash = next
	}
	for var_dtor_hash != nil {
		for i = 0; i < var_dtor_hash.GetUsedSlots(); i++ {
			var zv *zend.Zval = var_dtor_hash.GetData()[i]
			if zv.GetU2Extra() == VAR_WAKEUP_FLAG {

				/* Perform delayed __wakeup calls */

				if delayed_call_failed == 0 {
					var retval zend.Zval
					if zend.Z_ISUNDEF(wakeup_name) {
						zend.ZVAL_STRINGL(&wakeup_name, "__wakeup", b.SizeOf("\"__wakeup\"")-1)
					}
					BG(serialize_lock)++
					if zend.CallUserFunction(nil, zv, &wakeup_name, &retval, 0, 0) == zend.FAILURE || zend.Z_ISUNDEF(retval) {
						delayed_call_failed = 1
						zv.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
					}
					BG(serialize_lock)--
					zend.ZvalPtrDtor(&retval)
				} else {
					zv.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
				}

				/* Perform delayed __wakeup calls */

			} else if zv.GetU2Extra() == VAR_UNSERIALIZE_FLAG {

				/* Perform delayed __unserialize calls */

				if delayed_call_failed == 0 {
					var retval zend.Zval
					var param zend.Zval
					zend.ZVAL_COPY(&param, var_dtor_hash.GetData()[i+1])
					if zend.Z_ISUNDEF(unserialize_name) {
						zend.ZVAL_STRINGL(&unserialize_name, "__unserialize", b.SizeOf("\"__unserialize\"")-1)
					}
					BG(serialize_lock)++
					if zend.CallUserFunction(zend.__CG().GetFunctionTable(), zv, &unserialize_name, &retval, 1, &param) == zend.FAILURE || zend.Z_ISUNDEF(retval) {
						delayed_call_failed = 1
						zv.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
					}
					BG(serialize_lock)--
					zend.ZvalPtrDtor(&param)
					zend.ZvalPtrDtor(&retval)
				} else {
					zv.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
				}

				/* Perform delayed __unserialize calls */

			}
			zend.IZvalPtrDtor(zv)
		}
		next = var_dtor_hash.GetNext()
		zend.EfreeSize(var_dtor_hash, b.SizeOf("var_dtor_entries"))
		var_dtor_hash = next
	}
	zend.ZvalPtrDtorNogc(&wakeup_name)
	zend.ZvalPtrDtorNogc(&unserialize_name)
	if var_hashx.GetRefProps() != nil {
		zend.ZendHashDestroy(var_hashx.GetRefProps())
		zend.FREE_HASHTABLE(var_hashx.GetRefProps())
	}
}
func UnserializeStr(p **uint8, len_ int, maxlen int) *zend.ZendString {
	var i int
	var j int
	var str *zend.ZendString = zend.ZendStringSafeAlloc(1, len_, 0, 0)
	var end *uint8 = *((**uint8)(p + maxlen))
	if end < (*p) {
		zend.ZendStringEfree(str)
		return nil
	}
	for i = 0; i < len_; i++ {
		if (*p) >= end {
			zend.ZendStringEfree(str)
			return nil
		}
		if (*(*p)) != '\\' {
			str.GetVal()[i] = byte(*(*p))
		} else {
			var ch uint8 = 0
			for j = 0; j < 2; j++ {
				*p++
				if (*(*p)) >= '0' && (*(*p)) <= '9' {
					ch = (ch << 4) + ((*(*p)) - '0')
				} else if (*(*p)) >= 'a' && (*(*p)) <= 'f' {
					ch = (ch << 4) + ((*(*p)) - 'a' + 10)
				} else if (*(*p)) >= 'A' && (*(*p)) <= 'F' {
					ch = (ch << 4) + ((*(*p)) - 'A' + 10)
				} else {
					zend.ZendStringEfree(str)
					return nil
				}
			}
			str.GetVal()[i] = byte(ch)
		}
		*p++
	}
	str.GetVal()[i] = 0
	str.SetLen(i)
	return str
}
func UnserializeAllowedClass(class_name *zend.ZendString, var_hashx *PhpUnserializeDataT) int {
	var classes *zend.HashTable = var_hashx.GetAllowedClasses()
	var lcname *zend.ZendString
	var res int
	if classes == nil {
		return 1
	}
	if !(classes.GetNNumOfElements()) {
		return 0
	}
	zend.ZSTR_ALLOCA_ALLOC(lcname, class_name.GetLen(), use_heap)
	zend.ZendStrTolowerCopy(lcname.GetVal(), class_name.GetVal(), class_name.GetLen())
	res = zend.ZendHashExists(classes, lcname)
	zend.ZSTR_ALLOCA_FREE(lcname, use_heap)
	return res
}
func ParseIv2(p *uint8, q **uint8) zend.ZendLong {
	var result zend.ZendUlong = 0
	var neg zend.ZendUlong = 0
	var start *uint8
	if (*p) == '-' {
		neg = 1
		p++
	} else if (*p) == '+' {
		p++
	}
	for (*p) == '0' {
		p++
	}
	start = p
	for (*p) >= '0' && (*p) <= '9' {
		result = result*10 + (zend_ulong(*p) - '0')
		p++
	}
	if q != nil {
		*q = p
	}

	/* number too long or overflow */

	if p-start > zend.MAX_LENGTH_OF_LONG-1 || zend.SIZEOF_ZEND_LONG == 4 && p-start == zend.MAX_LENGTH_OF_LONG-1 && (*start) > '2' || result > zend.ZEND_LONG_MAX+neg {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Numerical result out of range")
		if neg == 0 {
			return zend.ZEND_LONG_MAX
		} else {
			return zend.ZEND_LONG_MIN
		}
	}
	return zend_long(b.Cond(neg == 0, result, -result))
}
func ParseIv(p *uint8) zend.ZendLong { return ParseIv2(p, nil) }
func ParseUiv(p *uint8) int {
	var cursor uint8
	var result int = 0
	for true {
		cursor = *p
		if cursor >= '0' && cursor <= '9' {
			result = result*10 + size_t(cursor-uint8('0'))
		} else {
			break
		}
		p++
	}
	return result
}
func ProcessNestedData(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, ht *zend.HashTable, elements zend.ZendLong, obj *zend.ZendObject) int {
	if var_hash != nil {
		if var_hash.GetMaxDepth() > 0 && var_hash.GetCurDepth() >= var_hash.GetMaxDepth() {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Maximum depth of "+zend.ZEND_LONG_FMT+" exceeded. "+"The depth limit can be changed using the max_depth unserialize() option "+"or the unserialize_max_depth ini setting", var_hash.GetMaxDepth())
			return 0
		}
		var_hash.GetCurDepth()++
	}
	for b.PostDec(&elements) > 0 {
		var key zend.Zval
		var data *zend.Zval
		var d zend.Zval
		var old_data *zend.Zval
		var idx zend.ZendUlong
		var info *zend.ZendPropertyInfo = nil
		zend.ZVAL_UNDEF(&key)
		if PhpVarUnserializeInternal(&key, p, max, nil, 1) == 0 {
			zend.ZvalPtrDtor(&key)
			goto failure
		}
		data = nil
		zend.ZVAL_UNDEF(&d)
		if obj == nil {
			if key.IsType(zend.IS_LONG) {
				idx = key.GetLval()
			numeric_key:
				if b.Assign(&old_data, zend.ZendHashIndexFind(ht, idx)) != nil {

					//??? update hash

					VarPushDtor(var_hash, old_data)
					data = ht.IndexUpdateH(idx, &d)
				} else {
					data = ht.IndexAddNewH(idx, &d)
				}
			} else if key.IsType(zend.IS_STRING) {
				if zend.ZEND_HANDLE_NUMERIC(key.GetStr(), &idx) {
					goto numeric_key
				}
				if b.Assign(&old_data, ht.FindByZendString(key.GetStr())) != nil {

					//??? update hash

					VarPushDtor(var_hash, old_data)
					data = zend.ZendHashUpdate(ht, key.GetStr().GetStr(), &d)
				} else {
					data = zend.ZendHashAddNew(ht, key.GetStr().GetStr(), &d)
				}
			} else {
				zend.ZvalPtrDtor(&key)
				goto failure
			}
		} else {
			if key.IsType(zend.IS_STRING) {
			string_key:
				if obj != nil && obj.GetCe().GetPropertiesInfo().GetNNumOfElements() > 0 {
					var existing_propinfo *zend.ZendPropertyInfo
					var new_key *zend.ZendString
					var unmangled_class *byte = nil
					var unmangled_prop *byte
					var unmangled_prop_len int
					var unmangled *zend.ZendString
					if zend.ZendUnmanglePropertyNameEx(key.GetStr(), &unmangled_class, &unmangled_prop, &unmangled_prop_len) == zend.FAILURE {
						zend.ZvalPtrDtor(&key)
						goto failure
					}
					unmangled = zend.ZendStringInit(unmangled_prop, unmangled_prop_len, 0)
					existing_propinfo = zend.ZendHashFindPtr(obj.GetCe().GetPropertiesInfo(), unmangled)
					if (unmangled_class == nil || !(strcmp(unmangled_class, "*")) || !(strcasecmp(unmangled_class, obj.GetCe().GetName().GetVal()))) && existing_propinfo != nil && existing_propinfo.HasFlags(zend.ZEND_ACC_PPP_MASK) {
						if existing_propinfo.HasFlags(zend.ZEND_ACC_PROTECTED) {
							new_key = zend.ZendManglePropertyName("*", 1, unmangled.GetVal(), unmangled.GetLen(), 0)
							zend.ZendStringReleaseEx(unmangled, 0)
						} else if existing_propinfo.HasFlags(zend.ZEND_ACC_PRIVATE) {
							if unmangled_class != nil && strcmp(unmangled_class, "*") != 0 {
								new_key = zend.ZendManglePropertyName(unmangled_class, strlen(unmangled_class), unmangled.GetVal(), unmangled.GetLen(), 0)
							} else {
								new_key = zend.ZendManglePropertyName(existing_propinfo.GetCe().GetName().GetVal(), existing_propinfo.GetCe().GetName().GetLen(), unmangled.GetVal(), unmangled.GetLen(), 0)
							}
							zend.ZendStringReleaseEx(unmangled, 0)
						} else {
							zend.ZEND_ASSERT(existing_propinfo.HasFlags(zend.ZEND_ACC_PUBLIC))
							new_key = unmangled
						}
						zend.ZvalPtrDtorStr(&key)
						zend.ZVAL_STR(&key, new_key)
					} else {
						zend.ZendStringReleaseEx(unmangled, 0)
					}
				}
				if b.Assign(&old_data, ht.FindByZendString(key.GetStr())) != nil {
					if old_data.IsType(zend.IS_INDIRECT) {
						old_data = old_data.GetZv()
						info = zend.ZendGetTypedPropertyInfoForSlot(obj, old_data)
						VarPushDtor(var_hash, old_data)
						data = zend.ZendHashUpdateInd(ht, key.GetStr().GetStr(), &d)
						if info != nil {

							/* Remember to which property this slot belongs, so we can add a
							 * type source if it is turned into a reference lateron. */

							if var_hash.GetRefProps() == nil {
								var_hash.SetRefProps(zend.Emalloc(b.SizeOf("HashTable")))
								zend.ZendHashInit(var_hash.GetRefProps(), 8, nil, nil, 0)
							}
							zend.ZendHashIndexUpdatePtr(var_hash.GetRefProps(), zend.ZendUintptrT(data), info)
						}
					} else {
						VarPushDtor(var_hash, old_data)
						data = zend.ZendHashUpdateInd(ht, key.GetStr().GetStr(), &d)
					}
				} else {
					data = zend.ZendHashAddNew(ht, key.GetStr().GetStr(), &d)
				}
			} else if key.IsType(zend.IS_LONG) {

				/* object properties should include no integers */

				zend.ConvertToString(&key)
				goto string_key
			} else {
				zend.ZvalPtrDtor(&key)
				goto failure
			}
		}
		if PhpVarUnserializeInternal(data, p, max, var_hash, 0) == 0 {
			zend.ZvalPtrDtor(&key)
			goto failure
		}
		if info != nil {
			if zend.ZendVerifyPropAssignableByRef(info, data, 1) == 0 {
				zend.ZvalPtrDtor(data)
				zend.ZVAL_UNDEF(data)
				zend.ZvalDtor(&key)
				goto failure
			}
			if zend.Z_ISREF_P(data) {
				zend.ZEND_REF_ADD_TYPE_SOURCE(data.GetRef(), info)
			}
		}
		if BG(unserialize).level > 1 {
			VarPushDtor(var_hash, data)
		}
		zend.ZvalPtrDtorStr(&key)
		if elements != 0 && (*((*p) - 1)) != ';' && (*((*p) - 1)) != '}' {
			*p--
			goto failure
		}
	}
	if var_hash != nil {
		var_hash.GetCurDepth()--
	}
	return 1
failure:
	if var_hash != nil {
		var_hash.GetCurDepth()--
	}
	return 0
}
func FinishNestedData(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int {
	if (*p) >= max || (*(*p)) != '}' {
		return 0
	}
	*p++
	return 1
}
func ObjectCustom(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, ce *zend.ZendClassEntry) int {
	var datalen zend.ZendLong
	datalen = ParseIv2((*p)+2, p)
	*p += 2
	if datalen < 0 || max-(*p) <= datalen {
		zend.ZendError(zend.E_WARNING, "Insufficient data for unserializing - "+zend.ZEND_LONG_FMT+" required, "+zend.ZEND_LONG_FMT+" present", datalen, zend_long(max-(*p)))
		return 0
	}

	/* Check that '}' is present before calling ce->unserialize() to mitigate issues
	 * with unserialize reading past the end of the passed buffer if the string is not
	 * appropriately terminated (usually NUL terminated, but '}' is also sufficient.) */

	if (*p)[datalen] != '}' {
		return 0
	}
	if ce.GetUnserialize() == nil {
		zend.ZendError(zend.E_WARNING, "Class %s has no unserializer", ce.GetName().GetVal())
		zend.ObjectInitEx(rval, ce)
	} else if ce.GetUnserialize()(rval, ce, (*uint8)(*p), datalen, (*zend.ZendUnserializeData)(var_hash)) != zend.SUCCESS {
		return 0
	}
	*p += datalen + 1
	return 1
}
func ObjectCommon(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, elements zend.ZendLong, has_unserialize zend.ZendBool) int {
	var ht *zend.HashTable
	var has_wakeup zend.ZendBool
	if has_unserialize != 0 {
		var ary zend.Zval
		var tmp *zend.Zval
		if elements >= zend.HT_MAX_SIZE {
			return 0
		}
		zend.ArrayInitSize(&ary, elements)

		/* Avoid reallocation due to packed -> mixed conversion. */

		zend.ZendHashRealInitMixed(ary.GetArr())
		if ProcessNestedData(rval, p, max, var_hash, ary.GetArr(), elements, nil) == 0 {
			zend.ZVAL_DEREF(rval)
			rval.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
			zend.ZvalPtrDtor(&ary)
			return 0
		}

		/* Delay __unserialize() call until end of serialization. We use two slots here to
		 * store both the object and the unserialized data array. */

		zend.ZVAL_DEREF(rval)
		tmp = TmpVar(var_hash, 2)
		zend.ZVAL_COPY(tmp, rval)
		tmp.SetU2Extra(VAR_UNSERIALIZE_FLAG)
		tmp++
		zend.ZVAL_COPY_VALUE(tmp, &ary)
		return FinishNestedData(rval, p, max, var_hash)
	}
	has_wakeup = zend.Z_OBJCE_P(rval) != PHP_IC_ENTRY && zend.ZendHashStrExists(zend.Z_OBJCE_P(rval).GetFunctionTable(), "__wakeup", b.SizeOf("\"__wakeup\"")-1) != 0
	ht = zend.Z_OBJPROP_P(rval)
	if elements >= zend_long(zend.HT_MAX_SIZE-ht.GetNNumOfElements()) {
		return 0
	}
	ht.Extend(ht.GetNNumOfElements() + elements)
	if ProcessNestedData(rval, p, max, var_hash, ht, elements, rval.GetObj()) == 0 {
		if has_wakeup != 0 {
			zend.ZVAL_DEREF(rval)
			rval.GetObj().AddGcFlags(zend.IS_OBJ_DESTRUCTOR_CALLED)
		}
		return 0
	}
	zend.ZVAL_DEREF(rval)
	if has_wakeup != 0 {

		/* Delay __wakeup call until end of serialization */

		var wakeup_var *zend.Zval = VarTmpVar(var_hash)
		zend.ZVAL_COPY(wakeup_var, rval)
		wakeup_var.SetU2Extra(VAR_WAKEUP_FLAG)
	}
	return FinishNestedData(rval, p, max, var_hash)
}
func PhpVarUnserialize(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int {
	var orig_var_entries *VarEntries = var_hash.GetLast()
	var orig_used_slots zend.ZendLong = b.CondF1(orig_var_entries != nil, func() zend.ZendLong { return orig_var_entries.GetUsedSlots() }, 0)
	var result int
	result = PhpVarUnserializeInternal(rval, p, max, var_hash, 0)
	if result == 0 {

		/* If the unserialization failed, mark all elements that have been added to var_hash
		 * as NULL. This will forbid their use by other unserialize() calls in the same
		 * unserialization context. */

		var e *VarEntries = orig_var_entries
		var s zend.ZendLong = orig_used_slots
		for e != nil {
			for ; s < e.GetUsedSlots(); s++ {
				e.GetData()[s] = nil
			}
			e = e.GetNext()
			s = 0
		}
	}
	return result
}
func PhpVarUnserializeInternal(rval *zend.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, as_key int) int {
	var cursor *uint8
	var limit *uint8
	var marker *uint8
	var start *uint8
	var rval_ref *zend.Zval
	limit = max
	cursor = *p
	if YYCURSOR >= YYLIMIT {
		return 0
	}
	if var_hash != nil && (*p)[0] != 'R' {
		VarPush(var_hash, rval)
	}
	start = cursor
	var yych uint8
	var yybm []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 128, 128, 128, 128, 128, 128, 128, 128, 128, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	if YYLIMIT-YYCURSOR < 7 {

	}
	yych = *YYCURSOR
	if yych <= 'a' {
		if yych <= 'O' {
			if yych <= 'C' {
				if yych <= 'B' {
					goto yy15
				}
				goto yy12
			} else {
				if yych <= 'M' {
					goto yy15
				}
				if yych <= 'N' {
					goto yy5
				}
				goto yy12
			}
		} else {
			if yych <= 'R' {
				if yych <= 'Q' {
					goto yy15
				}
			} else {
				if yych <= 'S' {
					goto yy10
				}
				if yych <= '`' {
					goto yy15
				}
				goto yy11
			}
		}
	} else {
		if yych <= 'i' {
			if yych <= 'c' {
				if yych <= 'b' {
					goto yy6
				}
				goto yy15
			} else {
				if yych <= 'd' {
					goto yy8
				}
				if yych <= 'h' {
					goto yy15
				}
				goto yy7
			}
		} else {
			if yych <= 's' {
				if yych <= 'q' {
					goto yy15
				}
				if yych <= 'r' {
					goto yy4
				}
				goto yy9
			} else {
				if yych == '}' {
					goto yy13
				}
				goto yy15
			}
		}
	}
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy84
	}
yy3:
	return 0
yy4:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy79
	}
	goto yy3
yy5:
	yych = *(b.PreInc(&YYCURSOR))
	if yych == ';' {
		goto yy77
	}
	goto yy3
yy6:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy70
	}
	goto yy3
yy7:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy64
	}
	goto yy3
yy8:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy41
	}
	goto yy3
yy9:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy35
	}
	goto yy3
yy10:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy29
	}
	goto yy3
yy11:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy23
	}
	goto yy3
yy12:
	yych = *(b.Assign(&YYMARKER, b.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy16
	}
	goto yy3
yy13:
	YYCURSOR++

	/* this is the case where we have less data than planned */

	core.PhpErrorDocref(nil, zend.E_NOTICE, "Unexpected end of serialized data")
	return 0
yy15:
	yych = *(b.PreInc(&YYCURSOR))
	goto yy3
yy16:
	yych = *(b.PreInc(&YYCURSOR))
	if (yybm[0+yych] & 128) != 0 {
		goto yy18
	}
yy17:
	YYCURSOR = YYMARKER
	goto yy3
yy18:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 2 {

	}
	yych = *YYCURSOR
	if (yybm[0+yych] & 128) != 0 {
		goto yy18
	}
	if yych != ':' {
		goto yy17
	}
	yych = *(b.PreInc(&YYCURSOR))
	if yych != '"' {
		goto yy17
	}
	YYCURSOR++
	var len int
	var len2 int
	var len3 int
	var maxlen int
	var elements zend_long
	var str *byte
	var class_name *zend.ZendString
	var ce *zend.ZendClassEntry
	var incomplete_class zend.ZendBool = 0
	var custom_object zend.ZendBool = 0
	var has_unserialize zend.ZendBool = 0
	var user_func zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	if var_hash == nil {
		return 0
	}
	if (*start) == 'C' {
		custom_object = 1
	}
	len_ = ParseUiv(start + 2)
	len2 = len_
	maxlen = max - YYCURSOR
	if maxlen < len_ || len_ == 0 {
		*p = start + 2
		return 0
	}
	str = (*byte)(YYCURSOR)
	YYCURSOR += len_
	if (*YYCURSOR) != '"' {
		*p = YYCURSOR
		return 0
	}
	if (*(YYCURSOR + 1)) != ':' {
		*p = YYCURSOR + 1
		return 0
	}
	len3 = strspn(str, "0123456789_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ177200201202203204205206207210211212213214215216217220221222223224225226227230231232233234235236237240241242243244245246247250251252253254255256257260261262263264265266267270271272273274275276277300301302303304305306307310311312313314315316317320321322323324325326327330331332333334335336337340341342343344345346347350351352353354355356357360361362363364365366367370371372373374375376377\\")
	if len3 != len_ {
		*p = YYCURSOR + len3 - len_
		return 0
	}
	class_name = zend.ZendStringInit(str, len_, 0)
	for {
		if UnserializeAllowedClass(class_name, var_hash) == 0 {
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			break
		}

		/* Try to find class directly */

		BG(serialize_lock)++
		ce = zend.ZendLookupClass(class_name)
		if ce != nil {
			BG(serialize_lock)--
			if zend.__EG().GetException() != nil {
				zend.ZendStringReleaseEx(class_name, 0)
				return 0
			}
			break
		}
		BG(serialize_lock)--
		if zend.__EG().GetException() != nil {
			zend.ZendStringReleaseEx(class_name, 0)
			return 0
		}

		/* Check for unserialize callback */

		if core.PG(unserialize_callback_func) == nil || core.PG(unserialize_callback_func)[0] == '0' {
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			break
		}

		/* Call unserialize callback */

		zend.ZVAL_STRING(&user_func, core.PG(unserialize_callback_func))
		zend.ZVAL_STR_COPY(&args[0], class_name)
		BG(serialize_lock)++
		if zend.CallUserFunctionEx(nil, nil, &user_func, &retval, 1, args, 0, nil) != zend.SUCCESS {
			BG(serialize_lock)--
			if zend.__EG().GetException() != nil {
				zend.ZendStringReleaseEx(class_name, 0)
				zend.ZvalPtrDtor(&user_func)
				zend.ZvalPtrDtor(&args[0])
				return 0
			}
			core.PhpErrorDocref(nil, zend.E_WARNING, "defined (%s) but not found", zend.Z_STRVAL(user_func))
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			zend.ZvalPtrDtor(&user_func)
			zend.ZvalPtrDtor(&args[0])
			break
		}
		BG(serialize_lock)--
		zend.ZvalPtrDtor(&retval)
		if zend.__EG().GetException() != nil {
			zend.ZendStringReleaseEx(class_name, 0)
			zend.ZvalPtrDtor(&user_func)
			zend.ZvalPtrDtor(&args[0])
			return 0
		}

		/* The callback function may have defined the class */

		BG(serialize_lock)++
		if b.Assign(&ce, zend.ZendLookupClass(class_name)) == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Function %s() hasn't defined the class it was called for", zend.Z_STRVAL(user_func))
			incomplete_class = 1
			ce = PHP_IC_ENTRY
		}
		BG(serialize_lock)--
		zend.ZvalPtrDtor(&user_func)
		zend.ZvalPtrDtor(&args[0])
		break

	}
	*p = YYCURSOR
	if custom_object != 0 {
		var ret int
		ret = ObjectCustom(rval, p, max, var_hash, ce)
		if ret != 0 && incomplete_class != 0 {
			PhpStoreClassName(rval, class_name.GetVal(), len2)
		}
		zend.ZendStringReleaseEx(class_name, 0)
		return ret
	}
	if (*p) >= max-2 {
		zend.ZendError(zend.E_WARNING, "Bad unserialize data")
		zend.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	elements = ParseIv2((*p)+2, p)
	if elements < 0 || elements > max-YYCURSOR {
		zend.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	*p += 2
	has_unserialize = incomplete_class == 0 && zend.ZendHashStrExists(ce.GetFunctionTable(), "__unserialize", b.SizeOf("\"__unserialize\"")-1) != 0

	/* If this class implements Serializable, it should not land here but in object_custom().
	 * The passed string obviously doesn't descend from the regular serializer. However, if
	 * there is both Serializable::unserialize() and __unserialize(), then both may be used,
	 * depending on the serialization format. */

	if ce.GetSerialize() != nil && has_unserialize == 0 {
		zend.ZendError(zend.E_WARNING, "Erroneous data format for unserializing '%s'", ce.GetName().GetVal())
		zend.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	if zend.ObjectInitEx(rval, ce) == zend.FAILURE {
		zend.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	if incomplete_class != 0 {
		PhpStoreClassName(rval, class_name.GetVal(), len2)
	}
	zend.ZendStringReleaseEx(class_name, 0)
	return ObjectCommon(rval, p, max, var_hash, elements, has_unserialize)
yy23:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy24:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 2 {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy24
	}
	if yych >= ';' {
		goto yy17
	}
	yych = *(b.PreInc(&YYCURSOR))
	if yych != '{' {
		goto yy17
	}
	YYCURSOR++
	var elements zend.ZendLong = ParseIv(start + 2)

	/* use iv() not uiv() in order to check data range */

	*p = YYCURSOR
	if var_hash == nil {
		return 0
	}
	if elements < 0 || elements >= zend.HT_MAX_SIZE || elements > max-YYCURSOR {
		return 0
	}
	if elements != 0 {
		zend.ArrayInitSize(rval, elements)

		/* we can't convert from packed to hash during unserialization, because
		   reference to some zvals might be keept in var_hash (to support references) */

		zend.ZendHashRealInitMixed(rval.GetArr())

		/* we can't convert from packed to hash during unserialization, because
		   reference to some zvals might be keept in var_hash (to support references) */

	} else {
		zend.ZVAL_EMPTY_ARRAY(rval)
		return FinishNestedData(rval, p, max, var_hash)
	}

	/* The array may contain references to itself, in which case we'll be modifying an
	 * rc>1 array. This is okay, since the array is, ostensibly, only visible to
	 * unserialize (in practice unserialization handlers also see it). Ideally we should
	 * prohibit "r:" references to non-objects, as we only generate them for objects. */

	if ProcessNestedData(rval, p, max, var_hash, rval.GetArr(), elements, nil) == 0 {
		return 0
	}
	return FinishNestedData(rval, p, max, var_hash)
yy29:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy30:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 2 {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy30
	}
	if yych >= ';' {
		goto yy17
	}
	yych = *(b.PreInc(&YYCURSOR))
	if yych != '"' {
		goto yy17
	}
	YYCURSOR++
	var len int
	var maxlen int
	var str *zend_string
	len_ = ParseUiv(start + 2)
	maxlen = max - YYCURSOR
	if maxlen < len_ {
		*p = start + 2
		return 0
	}
	if b.Assign(&str, UnserializeStr(&YYCURSOR, len_, maxlen)) == nil {
		return 0
	}
	if (*YYCURSOR) != '"' {
		zend.ZendStringEfree(str)
		*p = YYCURSOR
		return 0
	}
	if (*(YYCURSOR + 1)) != ';' {
		zend.Efree(str)
		*p = YYCURSOR + 1
		return 0
	}
	YYCURSOR += 2
	*p = YYCURSOR
	zend.ZVAL_STR(rval, str)
	return 1
yy35:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy36:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 2 {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy36
	}
	if yych >= ';' {
		goto yy17
	}
	yych = *(b.PreInc(&YYCURSOR))
	if yych != '"' {
		goto yy17
	}
	YYCURSOR++
	var len_ int
	var maxlen int
	var str *byte
	len_ = ParseUiv(start + 2)
	maxlen = max - YYCURSOR
	if maxlen < len_ {
		*p = start + 2
		return 0
	}
	str = (*byte)(YYCURSOR)
	YYCURSOR += len_
	if (*YYCURSOR) != '"' {
		*p = YYCURSOR
		return 0
	}
	if (*(YYCURSOR + 1)) != ';' {
		*p = YYCURSOR + 1
		return 0
	}
	YYCURSOR += 2
	*p = YYCURSOR
	if len_ == 0 {
		zend.ZVAL_EMPTY_STRING(rval)
	} else if len_ == 1 {
		zend.ZVAL_INTERNED_STR(rval, zend.ZSTR_CHAR(zend_uchar*str))
	} else if as_key != 0 {
		zend.ZVAL_STR(rval, zend.ZendStringInitInterned(str, len_, 0))
	} else {
		zend.ZVAL_STRINGL(rval, str, len_)
	}
	return 1
yy41:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		if yych <= ',' {
			if yych == '+' {
				goto yy45
			}
			goto yy17
		} else {
			if yych <= '-' {
				goto yy43
			}
			if yych <= '.' {
				goto yy48
			}
			goto yy17
		}
	} else {
		if yych <= 'I' {
			if yych <= '9' {
				goto yy46
			}
			if yych <= 'H' {
				goto yy17
			}
			goto yy44
		} else {
			if yych != 'N' {
				goto yy17
			}
		}
	}
	yych = *(b.PreInc(&YYCURSOR))
	if yych == 'A' {
		goto yy63
	}
	goto yy17
yy43:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		if yych == '.' {
			goto yy48
		}
		goto yy17
	} else {
		if yych <= '9' {
			goto yy46
		}
		if yych != 'I' {
			goto yy17
		}
	}
yy44:
	yych = *(b.PreInc(&YYCURSOR))
	if yych == 'N' {
		goto yy59
	}
	goto yy17
yy45:
	yych = *(b.PreInc(&YYCURSOR))
	if yych == '.' {
		goto yy48
	}
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy46:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 3 {

	}
	yych = *YYCURSOR
	if yych <= ':' {
		if yych <= '.' {
			if yych <= '-' {
				goto yy17
			}
			goto yy57
		} else {
			if yych <= '/' {
				goto yy17
			}
			if yych <= '9' {
				goto yy46
			}
			goto yy17
		}
	} else {
		if yych <= 'E' {
			if yych <= ';' {
				goto yy51
			}
			if yych <= 'D' {
				goto yy17
			}
			goto yy53
		} else {
			if yych == 'e' {
				goto yy53
			}
			goto yy17
		}
	}
yy48:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy49:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 3 {

	}
	yych = *YYCURSOR
	if yych <= ';' {
		if yych <= '/' {
			goto yy17
		}
		if yych <= '9' {
			goto yy49
		}
		if yych <= ':' {
			goto yy17
		}
	} else {
		if yych <= 'E' {
			if yych <= 'D' {
				goto yy17
			}
			goto yy53
		} else {
			if yych == 'e' {
				goto yy53
			}
			goto yy17
		}
	}
yy51:
	YYCURSOR++
	*p = YYCURSOR
	zend.ZVAL_DOUBLE(rval, zend.ZendStrtod((*byte)(start+2), nil))
	return 1
yy53:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= ',' {
		if yych != '+' {
			goto yy17
		}
	} else {
		if yych <= '-' {
			goto yy54
		}
		if yych <= '/' {
			goto yy17
		}
		if yych <= '9' {
			goto yy55
		}
		goto yy17
	}
yy54:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy55:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy55
	}
	if yych == ';' {
		goto yy51
	}
	goto yy17
yy57:
	YYCURSOR++
	if YYLIMIT-YYCURSOR < 3 {

	}
	yych = *YYCURSOR
	if yych <= ';' {
		if yych <= '/' {
			goto yy17
		}
		if yych <= '9' {
			goto yy57
		}
		if yych <= ':' {
			goto yy17
		}
		goto yy51
	} else {
		if yych <= 'E' {
			if yych <= 'D' {
				goto yy17
			}
			goto yy53
		} else {
			if yych == 'e' {
				goto yy53
			}
			goto yy17
		}
	}
yy59:
	yych = *(b.PreInc(&YYCURSOR))
	if yych != 'F' {
		goto yy17
	}
yy60:
	yych = *(b.PreInc(&YYCURSOR))
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	*p = YYCURSOR
	if !(strncmp((*byte)(start+2), "NAN", 3)) {
		zend.ZVAL_DOUBLE(rval, zend.ZEND_NAN)
	} else if !(strncmp((*byte)(start+2), "INF", 3)) {
		zend.ZVAL_DOUBLE(rval, zend.ZEND_INFINITY)
	} else if !(strncmp((*byte)(start+2), "-INF", 4)) {
		zend.ZVAL_DOUBLE(rval, -zend.ZEND_INFINITY)
	} else {
		zend.ZVAL_NULL(rval)
	}
	return 1
yy63:
	yych = *(b.PreInc(&YYCURSOR))
	if yych == 'N' {
		goto yy60
	}
	goto yy17
yy64:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= ',' {
		if yych != '+' {
			goto yy17
		}
	} else {
		if yych <= '-' {
			goto yy65
		}
		if yych <= '/' {
			goto yy17
		}
		if yych <= '9' {
			goto yy66
		}
		goto yy17
	}
yy65:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy66:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy66
	}
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	*p = YYCURSOR
	zend.ZVAL_LONG(rval, ParseIv(start+2))
	return 1
yy70:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych <= '0' {
		goto yy71
	}
	if yych <= '1' {
		goto yy72
	}
	goto yy17
yy71:
	yych = *(b.PreInc(&YYCURSOR))
	if yych == ';' {
		goto yy75
	}
	goto yy17
yy72:
	yych = *(b.PreInc(&YYCURSOR))
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	*p = YYCURSOR
	zend.ZVAL_TRUE(rval)
	return 1
yy75:
	YYCURSOR++
	*p = YYCURSOR
	zend.ZVAL_FALSE(rval)
	return 1
yy77:
	YYCURSOR++
	*p = YYCURSOR
	zend.ZVAL_NULL(rval)
	return 1
yy79:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy80:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy80
	}
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	var id zend_long
	*p = YYCURSOR
	if var_hash == nil {
		return 0
	}
	id = ParseUiv(start+2) - 1
	if id == -1 || b.Assign(&rval_ref, VarAccess(var_hash, id)) == nil {
		return 0
	}
	if rval_ref == rval {
		return 0
	}
	zend.ZVAL_DEREF(rval_ref)
	if rval_ref.GetType() != zend.IS_OBJECT {
		return 0
	}
	zend.ZVAL_COPY(rval, rval_ref)
	return 1
yy84:
	yych = *(b.PreInc(&YYCURSOR))
	if yych <= '/' {
		goto yy17
	}
	if yych >= ':' {
		goto yy17
	}
yy85:
	YYCURSOR++
	if YYLIMIT <= YYCURSOR {

	}
	yych = *YYCURSOR
	if yych <= '/' {
		goto yy17
	}
	if yych <= '9' {
		goto yy85
	}
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	var id zend.ZendLong
	*p = YYCURSOR
	if var_hash == nil {
		return 0
	}
	id = ParseUiv(start+2) - 1
	if id == -1 || b.Assign(&rval_ref, VarAccess(var_hash, id)) == nil {
		return 0
	}
	if zend.Z_ISUNDEF_P(rval_ref) || zend.Z_ISREF_P(rval_ref) && zend.Z_ISUNDEF_P(zend.Z_REFVAL_P(rval_ref)) {
		return 0
	}
	if !(zend.Z_ISREF_P(rval_ref)) {
		var info *zend.ZendPropertyInfo = nil
		if var_hash.GetRefProps() != nil {
			info = zend.ZendHashIndexFindPtr(var_hash.GetRefProps(), zend.ZendUintptrT(rval_ref))
		}
		zend.ZVAL_NEW_REF(rval_ref, rval_ref)
		if info != nil {
			zend.ZEND_REF_ADD_TYPE_SOURCE(rval_ref.GetRef(), info)
		}
	}
	zend.ZVAL_COPY(rval, rval_ref)
	return 1
	return 0
}
