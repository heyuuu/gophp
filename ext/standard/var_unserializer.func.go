package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"math"
)

func PhpVarUnserializeInit() PhpUnserializeDataT {
	var d PhpUnserializeDataT

	/* fprintf(stderr, "UNSERIALIZE_INIT    == lock: %u, level: %u\n", BG__().serialize_lock, BG__().unserialize.level); */

	if BG__().serialize_lock || !(BG__().unserialize.level) {
		d = zend.Emalloc(b.SizeOf("struct php_unserialize_data"))
		d.SetLast(d.GetEntries())
		d.SetLastDtor(nil)
		d.SetFirstDtor(d.GetLastDtor())
		d.SetAllowedClasses(nil)
		d.SetRefProps(nil)
		d.SetCurDepth(0)
		d.SetMaxDepth(BG__().unserialize_max_depth)
		d.GetEntries().SetUsedSlots(0)
		d.GetEntries().SetNext(nil)
		if !(BG__().serialize_lock) {
			BG__().unserialize.data = d
			BG__().unserialize.level = 1
		}
	} else {
		d = BG__().unserialize.data
		BG__().unserialize.level++
	}
	return d
}
func PhpVarUnserializeDestroy(d PhpUnserializeDataT) {
	if BG__().serialize_lock || BG__().unserialize.level == 1 {
		VarDestroy(&d)
		zend.Efree(d)
	}
	if !(BG__().serialize_lock) && !(lang.PreDec(&(BG__().unserialize.level))) {
		BG__().unserialize.data = nil
	}
}
func VarPush(var_hashx *PhpUnserializeDataT, rval *types.Zval) {
	var var_hash *VarEntries = var_hashx.GetLast()
	if var_hash.GetUsedSlots() == VAR_ENTRIES_MAX {
		var_hash = zend.Emalloc(b.SizeOf("var_entries"))
		var_hash.SetUsedSlots(0)
		var_hash.SetNext(0)
		var_hashx.GetLast().SetNext(var_hash)
		var_hashx.SetLast(var_hash)
	}
	var_hash.GetData()[lang.PostInc(&(var_hash.GetUsedSlots()))] = rval
}
func VarPushDtor(var_hashx *PhpUnserializeDataT, rval *types.Zval) {
	if rval.IsRefcounted() {
		var tmp_var *types.Zval = VarTmpVar(var_hashx)
		if tmp_var == nil {
			return
		}
		types.ZVAL_COPY(tmp_var, rval)
	}
}
func TmpVar(var_hashx *PhpUnserializeDataT, num zend.ZendLong) *types.Zval {
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
		var_hash.GetData()[var_hash.GetUsedSlots()].SetUndef()
		var_hash.GetData()[var_hash.GetUsedSlots()].GetU2Extra() = 0
	}
	return var_hash.GetData()[used_slots]
}
func VarTmpVar(var_hashx *PhpUnserializeDataT) *types.Zval { return TmpVar(var_hashx, 1) }
func VarReplace(var_hashx *PhpUnserializeDataT, ozval *types.Zval, nzval *types.Zval) {
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
func VarAccess(var_hashx *PhpUnserializeDataT, id zend.ZendLong) *types.Zval {
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
	var delayed_call_failed bool = 0
	var wakeup_name types.Zval
	var unserialize_name types.Zval
	wakeup_name.SetUndef()
	unserialize_name.SetUndef()
	for var_hash != nil {
		next = var_hash.GetNext()
		zend.EfreeSize(var_hash, b.SizeOf("var_entries"))
		var_hash = next
	}
	for var_dtor_hash != nil {
		for i = 0; i < var_dtor_hash.GetUsedSlots(); i++ {
			var zv *types.Zval = var_dtor_hash.GetData()[i]
			if zv.GetU2Extra() == VAR_WAKEUP_FLAG {

				/* Perform delayed __wakeup calls */

				if delayed_call_failed == 0 {
					var retval types.Zval
					if wakeup_name.IsUndef() {
						wakeup_name.SetString("__wakeup")
					}
					BG__().serialize_lock++
					if zend.CallUserFunction(zv, &wakeup_name, &retval, 0, 0) == types.FAILURE || retval.IsUndef() {
						delayed_call_failed = 1
						zv.Object().MarkObjDtorCalled()
					}
					BG__().serialize_lock--
				} else {
					zv.Object().MarkObjDtorCalled()
				}

				/* Perform delayed __wakeup calls */

			} else if zv.GetU2Extra() == VAR_UNSERIALIZE_FLAG {

				/* Perform delayed __unserialize calls */

				if delayed_call_failed == 0 {
					var retval types.Zval
					var param types.Zval
					types.ZVAL_COPY(&param, var_dtor_hash.GetData()[i+1])
					if unserialize_name.IsUndef() {
						unserialize_name.SetString("__unserialize")
					}
					BG__().serialize_lock++
					if zend.CallUserFunction(zv, &unserialize_name, &retval, 1, &param) == types.FAILURE || retval.IsUndef() {
						delayed_call_failed = 1
						zv.Object().MarkObjDtorCalled()
					}
					BG__().serialize_lock--
					// zend.ZvalPtrDtor(&param)
					// zend.ZvalPtrDtor(&retval)
				} else {
					zv.Object().MarkObjDtorCalled()
				}

				/* Perform delayed __unserialize calls */

			}
			// zend.IZvalPtrDtor(zv)
		}
		next = var_dtor_hash.GetNext()
		zend.EfreeSize(var_dtor_hash, b.SizeOf("var_dtor_entries"))
		var_dtor_hash = next
	}
	// zend.ZvalPtrDtorNogc(&wakeup_name)
	// zend.ZvalPtrDtorNogc(&unserialize_name)
	if var_hashx.GetRefProps() != nil {
		var_hashx.GetRefProps().Destroy()
	}
}
func UnserializeStr(p **uint8, len_ int, maxlen int) *types.String {
	var i int
	var j int
	var str *types.String = types.ZendStringAlloc(len_, 0)
	var end *uint8 = *((**uint8)(p + maxlen))
	if end < (*p) {
		// types.ZendStringEfree(str)
		return nil
	}
	for i = 0; i < len_; i++ {
		if (*p) >= end {
			// types.ZendStringEfree(str)
			return nil
		}
		if (*(*p)) != '\\' {
			str.GetStr()[i] = byte(*(*p))
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
					// types.ZendStringEfree(str)
					return nil
				}
			}
			str.GetStr()[i] = byte(ch)
		}
		*p++
	}
	return str.Cutoff(i)
}
func UnserializeAllowedClass(class_name *types.String, var_hashx *PhpUnserializeDataT) bool {
	var classes *types.Array = (**var_hashx).GetAllowedClasses()
	if classes == nil {
		return true
	}
	if classes.Len() == 0 {
		return false
	}
	lcname := ascii.StrToLower(class_name.GetStr())
	return classes.KeyExists(lcname)
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
		core.PhpErrorDocref(nil, faults.E_WARNING, "Numerical result out of range")
		if neg == 0 {
			return zend.ZEND_LONG_MAX
		} else {
			return zend.ZEND_LONG_MIN
		}
	}
	return zend_long(lang.Cond(neg == 0, result, -result))
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
func ProcessNestedData(
	rval *types.Zval,
	p **uint8,
	max *uint8,
	var_hash *PhpUnserializeDataT,
	ht *types.Array,
	elements zend.ZendLong,
	obj *types.Object,
) int {
	if var_hash != nil {
		if var_hash.GetMaxDepth() > 0 && var_hash.GetCurDepth() >= var_hash.GetMaxDepth() {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Maximum depth of "+zend.ZEND_LONG_FMT+" exceeded. "+"The depth limit can be changed using the max_depth unserialize() option "+"or the unserialize_max_depth ini setting", var_hash.GetMaxDepth())
			return 0
		}
		var_hash.GetCurDepth()++
	}
	for lang.PostDec(&elements) > 0 {
		var key types.Zval
		var data *types.Zval
		var d types.Zval
		var old_data *types.Zval
		var idx zend.ZendUlong
		var info *types.PropertyInfo = nil
		key.SetUndef()
		if PhpVarUnserializeInternal(&key, p, max, nil, 1) == 0 {
			// zend.ZvalPtrDtor(&key)
			goto failure
		}
		data = nil
		d.SetUndef()
		if obj == nil {
			if key.IsType(types.IsLong) {
				idx = key.Long()
			numeric_key:
				if lang.Assign(&old_data, ht.IndexFind(idx)) != nil {

					//??? update hash

					VarPushDtor(var_hash, old_data)
					data = ht.IndexUpdate(idx, &d)
				} else {
					data = ht.IndexAddNew(idx, &d)
				}
			} else if key.IsString() {
				if types.HandleNumericStr(key.String().GetStr(), &idx) {
					goto numeric_key
				}
				if lang.Assign(&old_data, ht.KeyFind(key.String().GetStr())) != nil {

					//??? update hash

					VarPushDtor(var_hash, old_data)
					data = ht.KeyUpdate(key.String().GetStr(), &d)
				} else {
					data = ht.KeyAddNew(key.String().GetStr(), &d)
				}
			} else {
				// zend.ZvalPtrDtor(&key)
				goto failure
			}
		} else {
			if key.IsString() {
			string_key:
				if obj != nil && obj.GetCe().PropertyTable().Len() > 0 {
					var existing_propinfo *types.PropertyInfo
					var new_key *types.String
					var unmangled_class *byte = nil
					var unmangled_prop *byte
					var unmangled_prop_len int
					var unmangled *types.String
					if zend.ZendUnmanglePropertyNameEx(key.String(), &unmangled_class, &unmangled_prop, &unmangled_prop_len) == types.FAILURE {
						// zend.ZvalPtrDtor(&key)
						goto failure
					}
					unmangled = types.NewString(b.CastStr(unmangled_prop, unmangled_prop_len))
					existing_propinfo = obj.GetCe().PropertyTable().Get(unmangled.GetStr())
					if (unmangled_class == nil || !(strcmp(unmangled_class, "*")) || !(strcasecmp(unmangled_class, obj.GetCe().Name()))) && existing_propinfo != nil && existing_propinfo.HasFlags(types.AccPppMask) {
						if existing_propinfo.HasFlags(types.AccProtected) {
							new_key = zend.ZendManglePropertyName_ZStr("*", unmangled.GetStr())
							// types.ZendStringReleaseEx(unmangled, 0)
						} else if existing_propinfo.HasFlags(types.AccPrivate) {
							if unmangled_class != nil && strcmp(unmangled_class, "*") != 0 {
								new_key = zend.ZendManglePropertyName_ZStr(unmangled_class, unmangled.GetStr())
							} else {
								new_key = zend.ZendManglePropertyName_ZStr(existing_propinfo.GetCe().GetName().GetStr(), unmangled.GetStr())
							}
							// types.ZendStringReleaseEx(unmangled, 0)
						} else {
							b.Assert(existing_propinfo.HasFlags(types.AccPublic))
							new_key = unmangled
						}

						key.SetStringEx(new_key)
					} else {
						// types.ZendStringReleaseEx(unmangled, 0)
					}
				}
				if lang.Assign(&old_data, ht.KeyFind(key.String().GetStr())) != nil {
					if old_data.IsIndirect() {
						old_data = old_data.Indirect()
						info = zend.ZendGetTypedPropertyInfoForSlot(obj, old_data)
						VarPushDtor(var_hash, old_data)
						data = ht.KeyUpdateIndirect(key.String().GetStr(), &d)
						if info != nil {

							/* Remember to which property this slot belongs, so we can add a
							 * type source if it is turned into a reference lateron. */

							if (**var_hash).GetRefProps() == nil {
								(**var_hash).SetRefProps(types.NewArray(8))
							}
							types.ZendHashIndexUpdatePtr((**var_hash).GetRefProps(), types.ZendUintptrT(data), info)
						}
					} else {
						VarPushDtor(var_hash, old_data)
						data = ht.KeyUpdateIndirect(key.String().GetStr(), &d)
					}
				} else {
					data = ht.KeyAddNew(key.String().GetStr(), &d)
				}
			} else if key.IsType(types.IsLong) {

				/* object properties should include no integers */

				operators.ConvertToString(&key)
				goto string_key
			} else {
				// zend.ZvalPtrDtor(&key)
				goto failure
			}
		}
		if PhpVarUnserializeInternal(data, p, max, var_hash, 0) == 0 {
			// zend.ZvalPtrDtor(&key)
			goto failure
		}
		if info != nil {
			if zend.ZendVerifyPropAssignableByRef(info, data, 1) == 0 {
				// zend.ZvalPtrDtor(data)
				data.SetUndef()
				//zend.ZvalDtor(&key)
				goto failure
			}
			if data.IsRef() {
				zend.ZEND_REF_ADD_TYPE_SOURCE(data.Reference(), info)
			}
		}
		if BG__().unserialize.level > 1 {
			VarPushDtor(var_hash, data)
		}

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
func FinishNestedData(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int {
	if (*p) >= max || (*(*p)) != '}' {
		return 0
	}
	*p++
	return 1
}
func ObjectCustom(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, ce *types.ClassEntry) int {
	var datalen zend.ZendLong
	datalen = ParseIv2((*p)+2, p)
	*p += 2
	if datalen < 0 || max-(*p) <= datalen {
		faults.Error(faults.E_WARNING, "Insufficient data for unserializing - "+zend.ZEND_LONG_FMT+" required, "+zend.ZEND_LONG_FMT+" present", datalen, zend_long(max-(*p)))
		return 0
	}

	/* Check that '}' is present before calling ce->unserialize() to mitigate issues
	 * with unserialize reading past the end of the passed buffer if the string is not
	 * appropriately terminated (usually NUL terminated, but '}' is also sufficient.) */

	if (*p)[datalen] != '}' {
		return 0
	}
	if ce.GetUnserialize() == nil {
		faults.Error(faults.E_WARNING, "Class %s has no unserializer", ce.Name())
		zend.ObjectInitEx(rval, ce)
	} else if ce.GetUnserialize()(rval, ce, (*uint8)(*p), datalen, (*zend.ZendUnserializeData)(var_hash)) != types.SUCCESS {
		return 0
	}
	*p += datalen + 1
	return 1
}
func ObjectCommon(
	rval *types.Zval,
	p **uint8,
	max *uint8,
	var_hash *PhpUnserializeDataT,
	elements zend.ZendLong,
	has_unserialize bool,
) int {
	var ht *types.Array
	var has_wakeup bool
	if has_unserialize != 0 {
		var ary types.Zval
		var tmp *types.Zval
		if elements >= types.MaxArraySize {
			return 0
		}
		zend.ArrayInitSize(&ary, elements)

		/* Avoid reallocation due to packed -> mixed conversion. */

		//types.ZendHashRealInitMixed(ary.Array())
		if ProcessNestedData(rval, p, max, var_hash, ary.Array(), elements, nil) == 0 {
			rval = types.ZVAL_DEREF(rval)
			rval.Object().MarkObjDtorCalled()
			// zend.ZvalPtrDtor(&ary)
			return 0
		}

		/* Delay __unserialize() call until end of serialization. We use two slots here to
		 * store both the object and the unserialized data array. */

		rval = types.ZVAL_DEREF(rval)
		tmp = TmpVar(var_hash, 2)
		types.ZVAL_COPY(tmp, rval)
		tmp.SetU2Extra(VAR_UNSERIALIZE_FLAG)
		tmp++
		types.ZVAL_COPY_VALUE(tmp, &ary)
		return FinishNestedData(rval, p, max, var_hash)
	}
	has_wakeup = types.Z_OBJCE_P(rval) != PHP_IC_ENTRY && types.Z_OBJCE_P(rval).FunctionTable().Exists("__wakeup")
	ht = types.Z_OBJPROP_P(rval)
	if elements >= zend_long(types.MaxArraySize-ht.Len()) {
		return 0
	}
	if ProcessNestedData(rval, p, max, var_hash, ht, elements, rval.Object()) == 0 {
		if has_wakeup != 0 {
			rval = types.ZVAL_DEREF(rval)
			rval.Object().MarkObjDtorCalled()
		}
		return 0
	}
	rval = types.ZVAL_DEREF(rval)
	if has_wakeup != 0 {

		/* Delay __wakeup call until end of serialization */

		var wakeup_var *types.Zval = VarTmpVar(var_hash)
		types.ZVAL_COPY(wakeup_var, rval)
		wakeup_var.SetU2Extra(VAR_WAKEUP_FLAG)
	}
	return FinishNestedData(rval, p, max, var_hash)
}
func PhpVarUnserialize(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT) int {
	var orig_var_entries *VarEntries = var_hash.GetLast()
	var orig_used_slots zend.ZendLong = lang.CondF1(orig_var_entries != nil, func() zend.ZendLong { return orig_var_entries.GetUsedSlots() }, 0)
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
func PhpVarUnserializeInternal(rval *types.Zval, p **uint8, max *uint8, var_hash *PhpUnserializeDataT, as_key int) int {
	var cursor *uint8
	var limit *uint8
	var marker *uint8
	var start *uint8
	var rval_ref *types.Zval
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
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy84
	}
yy3:
	return 0
yy4:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy79
	}
	goto yy3
yy5:
	yych = *(lang.PreInc(&YYCURSOR))
	if yych == ';' {
		goto yy77
	}
	goto yy3
yy6:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy70
	}
	goto yy3
yy7:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy64
	}
	goto yy3
yy8:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy41
	}
	goto yy3
yy9:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy35
	}
	goto yy3
yy10:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy29
	}
	goto yy3
yy11:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy23
	}
	goto yy3
yy12:
	yych = *(lang.Assign(&YYMARKER, lang.PreInc(&YYCURSOR)))
	if yych == ':' {
		goto yy16
	}
	goto yy3
yy13:
	YYCURSOR++

	/* this is the case where we have less data than planned */

	core.PhpErrorDocref(nil, faults.E_NOTICE, "Unexpected end of serialized data")
	return 0
yy15:
	yych = *(lang.PreInc(&YYCURSOR))
	goto yy3
yy16:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	var class_name *types.String
	var ce *types.ClassEntry
	var incomplete_class bool = 0
	var custom_object bool = 0
	var has_unserialize bool = 0
	var user_func types.Zval
	var retval types.Zval
	var args []types.Zval
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
	class_name = types.NewString(b.CastStr(str, len_))
	for {
		if !UnserializeAllowedClass(class_name, var_hash) {
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			break
		}

		/* Try to find class directly */

		BG__().serialize_lock++
		ce = zend.ZendLookupClass(class_name)
		if ce != nil {
			BG__().serialize_lock--
			if zend.EG__().GetException() != nil {
				// types.ZendStringReleaseEx(class_name, 0)
				return 0
			}
			break
		}
		BG__().serialize_lock--
		if zend.EG__().GetException() != nil {
			// types.ZendStringReleaseEx(class_name, 0)
			return 0
		}

		/* Check for unserialize callback */

		if core.PG__().unserialize_callback_func == nil || core.PG__().unserialize_callback_func[0] == '0' {
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			break
		}

		/* Call unserialize callback */

		user_func.SetString(b.CastStrAuto(core.PG__().unserialize_callback_func))
		args[0].SetString(class_name.GetStr())
		BG__().serialize_lock++
		if zend.CallUserFunctionEx(nil, &user_func, &retval, 1, args, 0) != types.SUCCESS {
			BG__().serialize_lock--
			if zend.EG__().GetException() != nil {
				// types.ZendStringReleaseEx(class_name, 0)
				// zend.ZvalPtrDtor(&user_func)
				// zend.ZvalPtrDtor(&args[0])
				return 0
			}
			core.PhpErrorDocref(nil, faults.E_WARNING, "defined (%s) but not found", user_func.String().GetVal())
			incomplete_class = 1
			ce = PHP_IC_ENTRY
			// zend.ZvalPtrDtor(&user_func)
			// zend.ZvalPtrDtor(&args[0])
			break
		}
		BG__().serialize_lock--
		// zend.ZvalPtrDtor(&retval)
		if zend.EG__().GetException() != nil {
			// types.ZendStringReleaseEx(class_name, 0)
			// zend.ZvalPtrDtor(&user_func)
			// zend.ZvalPtrDtor(&args[0])
			return 0
		}

		/* The callback function may have defined the class */

		BG__().serialize_lock++
		if lang.Assign(&ce, zend.ZendLookupClass(class_name)) == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Function %s() hasn't defined the class it was called for", user_func.String().GetVal())
			incomplete_class = 1
			ce = PHP_IC_ENTRY
		}
		BG__().serialize_lock--
		// zend.ZvalPtrDtor(&user_func)
		// zend.ZvalPtrDtor(&args[0])
		break

	}
	*p = YYCURSOR
	if custom_object != 0 {
		var ret int
		ret = ObjectCustom(rval, p, max, var_hash, ce)
		if ret != 0 && incomplete_class != 0 {
			PhpStoreClassName(rval, class_name.GetVal(), len2)
		}
		// types.ZendStringReleaseEx(class_name, 0)
		return ret
	}
	if (*p) >= max-2 {
		faults.Error(faults.E_WARNING, "Bad unserialize data")
		// types.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	elements = ParseIv2((*p)+2, p)
	if elements < 0 || elements > max-YYCURSOR {
		// types.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	*p += 2
	has_unserialize = incomplete_class == 0 && ce.FunctionTable().Exists("__unserialize")

	/* If this class implements Serializable, it should not land here but in object_custom().
	 * The passed string obviously doesn't descend from the regular serializer. However, if
	 * there is both Serializable::unserialize() and __unserialize(), then both may be used,
	 * depending on the serialization format. */

	if ce.GetSerialize() != nil && has_unserialize == 0 {
		faults.Error(faults.E_WARNING, "Erroneous data format for unserializing '%s'", ce.Name())
		// types.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	if zend.ObjectInitEx(rval, ce) == types.FAILURE {
		// types.ZendStringReleaseEx(class_name, 0)
		return 0
	}
	if incomplete_class != 0 {
		PhpStoreClassName(rval, class_name.GetVal(), len2)
	}
	// types.ZendStringReleaseEx(class_name, 0)
	return ObjectCommon(rval, p, max, var_hash, elements, has_unserialize)
yy23:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	if elements < 0 || elements >= types.MaxArraySize || elements > max-YYCURSOR {
		return 0
	}
	if elements != 0 {
		zend.ArrayInitSize(rval, elements)

		/* we can't convert from packed to hash during unserialization, because
		   reference to some zvals might be keept in var_hash (to support references) */

		//types.ZendHashRealInitMixed(rval.Array())

		/* we can't convert from packed to hash during unserialization, because
		   reference to some zvals might be keept in var_hash (to support references) */

	} else {
		rval.SetEmptyArray()
		return FinishNestedData(rval, p, max, var_hash)
	}

	/* The array may contain references to itself, in which case we'll be modifying an
	 * rc>1 array. This is okay, since the array is, ostensibly, only visible to
	 * unserialize (in practice unserialization handlers also see it). Ideally we should
	 * prohibit "r:" references to non-objects, as we only generate them for objects. */

	if ProcessNestedData(rval, p, max, var_hash, rval.Array(), elements, nil) == 0 {
		return 0
	}
	return FinishNestedData(rval, p, max, var_hash)
yy29:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	if lang.Assign(&str, UnserializeStr(&YYCURSOR, len_, maxlen)) == nil {
		return 0
	}
	if (*YYCURSOR) != '"' {
		// types.ZendStringEfree(str)
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
	rval.SetStringEx(str)
	return 1
yy35:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	rval.SetString(b.CastStr(str, len_))
	return 1
yy41:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
	if yych == 'A' {
		goto yy63
	}
	goto yy17
yy43:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
	if yych == 'N' {
		goto yy59
	}
	goto yy17
yy45:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	rval.SetDouble(zend.ZendStrtod((*byte)(start+2), nil))
	return 1
yy53:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
	if yych != 'F' {
		goto yy17
	}
yy60:
	yych = *(lang.PreInc(&YYCURSOR))
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	*p = YYCURSOR
	if !(strncmp((*byte)(start+2), "NAN", 3)) {
		rval.SetDouble(math.NaN())
	} else if !(strncmp((*byte)(start+2), "INF", 3)) {
		rval.SetDouble(math.Inf(1))
	} else if !(strncmp((*byte)(start+2), "-INF", 4)) {
		rval.SetDouble(math.Inf(-1))
	} else {
		rval.SetNull()
	}
	return 1
yy63:
	yych = *(lang.PreInc(&YYCURSOR))
	if yych == 'N' {
		goto yy60
	}
	goto yy17
yy64:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
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
	rval.SetLong(ParseIv(start + 2))
	return 1
yy70:
	yych = *(lang.PreInc(&YYCURSOR))
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
	yych = *(lang.PreInc(&YYCURSOR))
	if yych == ';' {
		goto yy75
	}
	goto yy17
yy72:
	yych = *(lang.PreInc(&YYCURSOR))
	if yych != ';' {
		goto yy17
	}
	YYCURSOR++
	*p = YYCURSOR
	rval.SetTrue()
	return 1
yy75:
	YYCURSOR++
	*p = YYCURSOR
	rval.SetFalse()
	return 1
yy77:
	YYCURSOR++
	*p = YYCURSOR
	rval.SetNull()
	return 1
yy79:
	yych = *(lang.PreInc(&YYCURSOR))
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
	if id == -1 || lang.Assign(&rval_ref, VarAccess(var_hash, id)) == nil {
		return 0
	}
	if rval_ref == rval {
		return 0
	}
	rval_ref = types.ZVAL_DEREF(rval_ref)
	if !rval_ref.IsObject() {
		return 0
	}
	types.ZVAL_COPY(rval, rval_ref)
	return 1
yy84:
	yych = *(lang.PreInc(&YYCURSOR))
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
	if id == -1 || lang.Assign(&rval_ref, VarAccess(var_hash, id)) == nil {
		return 0
	}
	if rval_ref.IsUndef() || rval_ref.IsRef() && types.Z_REFVAL_P(rval_ref).IsUndef() {
		return 0
	}
	if !(rval_ref.IsRef()) {
		var info *types.PropertyInfo = nil
		if var_hash.GetRefProps() != nil {
			info = types.ZendHashIndexFindPtr(var_hash.GetRefProps(), types.ZendUintptrT(rval_ref))
		}
		rval_ref.SetNewRef(rval_ref)
		if info != nil {
			zend.ZEND_REF_ADD_TYPE_SOURCE(rval_ref.Reference(), info)
		}
	}
	types.ZVAL_COPY(rval, rval_ref)
	return 1
	return 0
}
