package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func PhpArrayElementDump(zv *types.Zval, key types.ArrayKey, level int) {
	if key.IsStrKey() {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PhpOutputWrite(key.StrKey())
		core.PhpPrintf("\"]=>\n")
	} else {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key.IndexKey())
	}
	PhpVarDump(zv, level+2)
}
func PhpObjectPropertyDump(propInfo *zend.ZendPropertyInfo, zv *types.Zval, key_ types.ArrayKey, level int) {
	if !key_.IsStrKey() {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key_.IndexKey())
	} else {
		className, propName, ok := zend.ZendUnmanglePropertyName_Ex(key_.StrKey())
		core.PhpPrintf("%*c[", level+1, ' ')
		if ok {
			if className[0] == '*' {
				core.PhpPrintf("\"%s\":protected", propName)
			} else {
				core.PhpPrintf("\"%s\":\"%s\":private", propName, className)
			}
		} else {
			core.PhpPrintf("\"")
			core.PhpOutputWrite(key_.StrKey())
			core.PhpPrintf("\"")
		}
		zend.ZEND_PUTS("]=>\n")
	}
	if zv.IsUndef() {
		b.Assert(propInfo.GetType() != 0)
		var typ string
		if propInfo.GetType().IsClass() {
			if propInfo.GetType().IsCe() {
				typ = types.ZEND_TYPE_CE(propInfo.GetType()).Name()
			} else {
				typ = propInfo.GetType().Name().GetStr()
			}
		} else {
			typ = types.ZendGetTypeByConst(propInfo.GetType().Code())
		}
		if propInfo.GetType().AllowNull() {
			typ = "?" + typ
		}

		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', typ)
	} else {
		PhpVarDump(zv, level+2)
	}
}
func PhpVarDump(struc *types.Zval, level int) {
	isRef := false
	if level > 1 {
		core.PhpPrintf("%*c", level-1, ' ')
	}
again:
	common := ""
	if isRef {
		common = "&"
	}
	switch struc.GetType() {
	case types.IS_FALSE:
		core.PhpPrintf("%sbool(false)\n", common)
	case types.IS_TRUE:
		core.PhpPrintf("%sbool(true)\n", common)
	case types.IS_NULL:
		core.PhpPrintf("%sNULL\n", common)
	case types.IS_LONG:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", common, struc.Long())
	case types.IS_DOUBLE:
		core.PhpPrintf("%sfloat(%.*G)\n", common, int(zend.EG__().GetPrecision()), struc.Double())
	case types.IS_STRING:
		core.PhpPrintf("%sstring(%zd) \"", common, struc.String().GetLen())
		core.PhpOutputWrite(struc.StringVal())
		core.PUTS("\"\n")
	case types.IS_ARRAY:
		myht := struc.Array()
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if level > 1 {
				if myht.IsRecursive() {
					core.PUTS("*RECURSION*\n")
					return
				}
				myht.ProtectRecursive()
			}
			myht.AddRefcount()
		}
		count := myht.Count()
		core.PhpPrintf("%sarray(%d) {\n", common, count)
		myht.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			PhpArrayElementDump(value, key, level)
		})
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if level > 1 {
				myht.UnprotectRecursive()
			}
			myht.DelRefcount()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IS_OBJECT:
		if struc.IsRecursive() {
			core.PUTS("*RECURSION*\n")
			return
		}
		struc.ProtectRecursive()
		myht := zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		className := types.Z_OBJ_HT(*struc).GetGetClassName()(struc.Object())
		core.PhpPrintf("%sobject(%s)#%d (%d) {\n", common, className.GetVal(), zend.Z_OBJ_HANDLE_P(struc), b.CondF1(myht != nil, func() int { return myht.Count() }, 0))
		if myht != nil {
			myht.Foreach(func(key types.ArrayKey, value *types.Zval) {
				var prop_info *zend.ZendPropertyInfo = nil
				if value.IsIndirect() {
					value = value.Indirect()
					if key.IsStrKey() {
						prop_info = zend.ZendGetTypedPropertyInfoForSlot(struc.Object(), value)
					}
				}
				if !value.IsUndef() || prop_info != nil {
					PhpObjectPropertyDump(prop_info, value, key, level)
				}
			})
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
		struc.UnprotectRecursive()
	case types.IS_RESOURCE:
		typeName := b.Option(zend.ZendRsrcListGetRsrcTypeEx(types.Z_RES_P(struc)), "Unknown")
		core.PhpPrintf("%sresource(%d) of type (%s)\n", common, types.Z_RES_P(struc).GetHandle(), typeName)
	case types.IS_REFERENCE:
		//??? hide references with refcount==1 (for compatibility)
		if struc.GetRefcount() > 1 {
			isRef = true
		}
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", common)
	}
}
func ZifVarDump(vars []*types.Zval) {
	for _, zv := range vars {
		PhpVarDump(zv, 1)
	}
}
func ZvalArrayElementDump(zv *types.Zval, index zend.ZendUlong, key *types.String, level int) {
	if key == nil {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', index)
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PhpOutputWrite(key.GetStr())
		core.PhpPrintf("\"]=>\n")
	}
	PhpDebugZvalDump(zv, level+2)
}
func ZvalObjectPropertyDump(propInfo *zend.ZendPropertyInfo, zv *types.Zval, key types.ArrayKey, level int) {
	if !key.IsStrKey() {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key.IndexKey())
	} else {
		className, propName, ok := zend.ZendUnmanglePropertyName_Ex(key.StrKey())
		core.PhpPrintf("%*c[", level+1, ' ')
		if ok {
			if className[0] == '*' {
				core.PhpPrintf("\"%s\":protected", propName)
			} else {
				core.PhpPrintf("\"%s\":\"%s\":private", propName, className)
			}
		} else {
			core.PhpPrintf("\"%s\"", propName)
		}
		zend.ZEND_PUTS("]=>\n")
	}
	if propInfo != nil && zv.IsUndef() {
		b.Assert(propInfo.GetType() != 0)
		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', b.Cond(propInfo.GetType().AllowNull(), "?", ""), b.CondF(propInfo.GetType().IsClass(), func() []byte {
			return b.CondF(propInfo.GetType().IsCe(), func() *types.String { return types.ZEND_TYPE_CE(propInfo.GetType()).GetName() }, func() *types.String { return propInfo.GetType().Name() }).GetVal()
		}, func() *byte { return types.ZendGetTypeByConst(propInfo.GetType().Code()) }))
	} else {
		PhpDebugZvalDump(zv, level+2)
	}
}
func PhpDebugZvalDump(struc *types.Zval, level int) {
	var myht *types.Array = nil
	var class_name *types.String
	var is_ref int = 0
	var index zend.ZendUlong
	var key *types.String
	var val *types.Zval
	var count uint32
	if level > 1 {
		core.PhpPrintf("%*c", level-1, ' ')
	}
again:
	switch struc.GetType() {
	case types.IS_FALSE:
		core.PhpPrintf("%sbool(false)\n", COMMON)
	case types.IS_TRUE:
		core.PhpPrintf("%sbool(true)\n", COMMON)
	case types.IS_NULL:
		core.PhpPrintf("%sNULL\n", COMMON)
	case types.IS_LONG:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", COMMON, struc.Long())
	case types.IS_DOUBLE:
		core.PhpPrintf("%sfloat(%.*G)\n", COMMON, int(zend.EG__().GetPrecision()), struc.Double())
	case types.IS_STRING:
		core.PhpPrintf("%sstring(%zd) \"", COMMON, struc.String().GetLen())
		core.PUTS(struc.String().GetStr())
		core.PhpPrintf("\" refcount(%u)\n", b.CondF1(struc.IsRefcounted(), func() uint32 { return struc.GetRefcount() }, 1))
	case types.IS_ARRAY:
		myht = struc.Array()
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if level > 1 {
				if myht.IsRecursive() {
					core.PUTS("*RECURSION*\n")
					return
				}
				myht.ProtectRecursive()
			}
			myht.AddRefcount()
		}
		count = myht.Count()
		core.PhpPrintf("%sarray(%d) refcount(%u){\n", COMMON, count, b.CondF1(struc.IsRefcounted(), func() int { return struc.GetRefcount() - 1 }, 1))
		var __ht *types.Array = myht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			ZvalArrayElementDump(val, index, key, level)
		}
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if level > 1 {
				myht.UnprotectRecursive()
			}
			myht.DelRefcount()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IS_OBJECT:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		if myht != nil {
			if myht.IsRecursive() {
				core.PUTS("*RECURSION*\n")
				zend.ZendReleaseProperties(myht)
				return
			}
			myht.ProtectRecursive()
		}
		class_name = types.Z_OBJ_HT(*struc).GetGetClassName()(struc.Object())
		core.PhpPrintf("%sobject(%s)#%d (%d) refcount(%u){\n", COMMON, class_name.GetVal(), zend.Z_OBJ_HANDLE_P(struc), b.CondF1(myht != nil, func() uint32 { return myht.Count() }, 0), struc.GetRefcount())
		// types.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			myht.Foreach(func(key types.ArrayKey, value *types.Zval) {
				var propInfo *zend.ZendPropertyInfo = nil
				if value.IsIndirect() {
					value = value.Indirect()
					if key.IsStrKey() {
						propInfo = zend.ZendGetTypedPropertyInfoForSlot(struc.Object(), value)
					}
				}
				if !(val.IsUndef()) || propInfo != nil {
					ZvalObjectPropertyDump(propInfo, val, key, level)
				}
			})
			myht.UnprotectRecursive()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IS_RESOURCE:
		typeName := b.Option(zend.ZendRsrcListGetRsrcTypeEx(types.Z_RES_P(struc)), "Unknown")
		core.PhpPrintf("%sresource(%d) of type (%s) refcount(%u)\n", COMMON, types.Z_RES_P(struc).GetHandle(), typeName, struc.GetRefcount())
	case types.IS_REFERENCE:

		//??? hide references with refcount==1 (for compatibility)

		if struc.GetRefcount() > 1 {
			is_ref = 1
		}
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", COMMON)
	}
}
func ZifDebugZvalDump(vars []*types.Zval) {
	for _, zval := range vars {
		PhpDebugZvalDump(zval, 1)
	}
}
func BufferAppendSpaces(buf *zend.SmartStr, num_spaces int) {
	var tmp_spaces *byte
	var tmp_spaces_len int
	tmp_spaces_len = core.Spprintf(&tmp_spaces, 0, "%*c", num_spaces, ' ')
	buf.AppendString(b.CastStr(tmp_spaces, tmp_spaces_len))
	zend.Efree(tmp_spaces)
}
func PhpArrayElementExport(zv *types.Zval, index zend.ZendUlong, key *types.String, level int, buf *zend.SmartStr) {
	if key == nil {
		BufferAppendSpaces(buf, level+1)
		buf.AppendLong(zend.ZendLong(index))
		buf.AppendString(" => ")
	} else {
		var ckey *types.String = types.NewString(str.PhpAddcslashes(key.GetStr(), "'\\"))
		tmp_str := strings.ReplaceAll(ckey.GetStr(), "0", "' . \"\\0\" . '")
		BufferAppendSpaces(buf, level+1)
		buf.AppendByte('\'')
		buf.AppendString(tmp_str)
		buf.AppendString("' => ")
		//types.ZendStringFree(ckey)
		//types.ZendStringFree(tmp_str)
	}
	PhpVarExportEx(zv, level+2, buf)
	buf.AppendByte(',')
	buf.AppendByte('\n')
}
func PhpObjectElementExport(zv *types.Zval, index zend.ZendUlong, key *types.String, level int, buf *zend.SmartStr) {
	BufferAppendSpaces(buf, level+2)
	if key != nil {
		_, propName, _ := zend.ZendUnmanglePropertyName_Ex(key.GetStr())

		propNameEscaped := str.PhpAddcslashes(propName, "'\\")
		buf.AppendByte('\'')
		buf.AppendString(propNameEscaped)
		buf.AppendByte('\'')
	} else {
		buf.AppendLong(zend.ZendLong(index))
	}
	buf.AppendString(" => ")
	PhpVarExportEx(zv, level+2, buf)
	buf.AppendByte(',')
	buf.AppendByte('\n')
}
func PhpVarExportEx(struc *types.Zval, level int, buf *zend.SmartStr) {
	var myht *types.Array
	var tmp_str []byte
	var index zend.ZendUlong
	var key *types.String
	var val *types.Zval
again:
	switch struc.GetType() {
	case types.IS_FALSE:
		buf.AppendString("false")
	case types.IS_TRUE:
		buf.AppendString("true")
	case types.IS_NULL:
		buf.AppendString("NULL")
	case types.IS_LONG:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if struc.Long() == zend.ZEND_LONG_MIN {
			buf.AppendLong(zend.ZEND_LONG_MIN + 1)
			buf.AppendString("-1")
			break
		}
		buf.AppendLong(struc.Long())
	case types.IS_DOUBLE:
		core.PhpGcvt(struc.Double(), int(core.PG__().serialize_precision), '.', 'E', tmp_str)
		buf.AppendString(b.CastStrAuto(tmp_str))

		/* Without a decimal point, PHP treats a number literal as an int.
		 * This check even works for scientific notation, because the
		 * mantissa always contains a decimal point.
		 * We need to check for finiteness, because INF, -INF and NAN
		 * must not have a decimal point added.
		 */

		if core.ZendFinite(struc.Double()) && nil == strchr(tmp_str, '.') {
			buf.AppendString(".0")
		}
	case types.IS_STRING:
		ztmp := str.PhpAddcslashes(struc.StringVal(), "'\\")
		ztmp2 := strings.ReplaceAll(ztmp, "0", "' . \"\\0\" . '")
		buf.AppendByte('\'')
		buf.AppendString(ztmp2)
		buf.AppendByte('\'')
		//types.ZendStringFree(ztmp)
		//types.ZendStringFree(ztmp2)
	case types.IS_ARRAY:
		myht = struc.Array()
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			if myht.IsRecursive() {
				buf.AppendString("NULL")
				faults.Error(faults.E_WARNING, "var_export does not handle circular references")
				return
			}
			myht.AddRefcount()
			myht.ProtectRecursive()
		}
		if level > 1 {
			buf.AppendByte('\n')
			BufferAppendSpaces(buf, level-1)
		}
		buf.AppendString("array (\n")
		var __ht *types.Array = myht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			PhpArrayElementExport(val, index, key, level, buf)
		}
		if (myht.GetGcFlags() & types.GC_IMMUTABLE) == 0 {
			myht.UnprotectRecursive()
			myht.DelRefcount()
		}
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		buf.AppendByte(')')
	case types.IS_OBJECT:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_VAR_EXPORT)
		if myht != nil {
			if myht.IsRecursive() {
				buf.AppendString("NULL")
				faults.Error(faults.E_WARNING, "var_export does not handle circular references")
				zend.ZendReleaseProperties(myht)
				return
			} else {
				myht.TryProtectRecursive()
			}
		}
		if level > 1 {
			buf.AppendByte('\n')
			BufferAppendSpaces(buf, level-1)
		}

		/* stdClass has no __set_state method, but can be casted to */

		if types.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			buf.AppendString("(object) array(\n")
		} else {
			buf.AppendString(types.Z_OBJCE_P(struc).GetName().GetStr())
			buf.AppendString("::__set_state(array(\n")
		}
		if myht != nil {
			var __ht *types.Array = myht
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				PhpObjectElementExport(val, index, key, level, buf)
			}
			myht.TryUnProtectRecursive()
			zend.ZendReleaseProperties(myht)
		}
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		if types.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			buf.AppendByte(')')
		} else {
			buf.AppendString("))")
		}
	case types.IS_REFERENCE:
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		buf.AppendString("NULL")
	}
}
func PhpVarExport(struc *types.Zval, level int) {
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	PhpVarExportEx(struc, level, &buf)
	buf.ZeroTail()
	core.PUTS(buf.GetS().GetStr())
	buf.Free()
}
func ZifVarExport(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval, _ zpp.Opt, return_ *types.Zval) {
	var var_ *types.Zval
	var return_output types.ZendBool = 0
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			var_ = fp.ParseZval()
			fp.StartOptional()
			return_output = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	PhpVarExportEx(var_, 1, &buf)
	buf.ZeroTail()
	if return_output != 0 {
		return_value.SetString(buf.GetS())
		return
	} else {
		core.PUTS(buf.GetS().GetStr())
		buf.Free()
	}
}
func PhpAddVarHash(data PhpSerializeDataT, var_ *types.Zval) zend.ZendLong {
	var zv *types.Zval
	var key zend.ZendUlong
	var is_ref types.ZendBool = var_.IsReference()
	data.SetN(data.GetN() + 1)
	if is_ref == 0 && var_.GetType() != types.IS_OBJECT {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */

	if is_ref != 0 && types.Z_REFVAL_P(var_).IsType(types.IS_OBJECT) {
		var_ = types.Z_REFVAL_P(var_)
	}

	/* Index for the variable is stored using the numeric value of the pointer to
	 * the zend_refcounted struct */

	key = zend.ZendUlong(types.ZendUintptrT(var_.RefCounted()))
	zv = data.GetHt().IndexFindH(key)
	if zv != nil {

		/* References are only counted once, undo the data->n increment above */

		if is_ref != 0 && zv.Long() != -1 {
			data.SetN(data.GetN() - 1)
		}
		return zv.Long()
	} else {
		var zv_n types.Zval
		zv_n.SetLong(data.GetN())
		data.GetHt().IndexAddNewH(key, &zv_n)

		/* Additionally to the index, we also store the variable, to ensure that it is
		 * not destroyed during serialization and its pointer reused. The variable is
		 * stored at the numeric value of the pointer + 1, which cannot be the location
		 * of another zend_refcounted structure. */

		data.GetHt().IndexAddNewH(key+1, var_)
		var_.AddRefcount()
		return 0
	}
}
func PhpVarSerializeLong(buf *zend.SmartStr, val zend.ZendLong) {
	buf.AppendString("i:")
	buf.AppendLong(val)
	buf.AppendByte(';')
}
func PhpVarSerializeString(buf *zend.SmartStr, str *byte, len_ int) {
	buf.AppendString("s:")
	buf.AppendUlong(len_)
	buf.AppendString(":\"")
	buf.AppendString(b.CastStr(str, len_))
	buf.AppendString("\";")
}
func PhpVarSerializeClassName(buf *zend.SmartStr, struc *types.Zval) types.ZendBool {
	var class_name *types.String
	var incomplete_class types.ZendBool = 0
	PHP_SET_CLASS_ATTRIBUTES(struc)
	buf.AppendString("O:")
	buf.AppendUlong(class_name.GetLen())
	buf.AppendString(":\"")
	buf.AppendString(class_name.GetStr())
	buf.AppendString("\":")
	PHP_CLEANUP_CLASS_ATTRIBUTES()
	return incomplete_class
}
func PhpVarSerializeCallSleep(retval *types.Zval, struc *types.Zval) int {
	var fname types.Zval
	var res int
	fname.SetStringVal("__sleep")
	BG__().serialize_lock++
	res = zend.CallUserFunction(nil, struc, &fname, retval, 0, 0)
	BG__().serialize_lock--

	if res == types.FAILURE || retval.IsUndef() {
		zend.ZvalPtrDtor(retval)
		return types.FAILURE
	}
	if !(zend.HASH_OF(retval)) {
		zend.ZvalPtrDtor(retval)
		core.PhpErrorDocref(nil, faults.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize")
		return types.FAILURE
	}
	return types.SUCCESS
}
func PhpVarSerializeCallMagicSerialize(retval *types.Zval, obj *types.Zval) int {
	var fname types.Zval
	var res int
	fname.SetStringVal("__serialize")
	BG__().serialize_lock++
	res = zend.CallUserFunction(obj, &fname, retval, 0, 0)
	BG__().serialize_lock--

	if res == types.FAILURE || retval.IsUndef() {
		zend.ZvalPtrDtor(retval)
		return types.FAILURE
	}
	if retval.GetType() != types.IS_ARRAY {
		zend.ZvalPtrDtor(retval)
		faults.TypeError("%s::__serialize() must return an array", types.Z_OBJCE_P(obj).GetName().GetVal())
		return types.FAILURE
	}
	return types.SUCCESS
}
func PhpVarSerializeTryAddSleepProp(ht *types.Array, props *types.Array, name *types.String, error_name *types.String, struc *types.Zval) int {
	var val *types.Zval = props.KeyFind(name.GetStr())
	if val == nil {
		return types.FAILURE
	}
	if val.IsIndirect() {
		val = val.Indirect()
		if val.IsUndef() {
			var info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(struc.Object(), val)
			if info != nil {
				return types.SUCCESS
			}
			return types.FAILURE
		}
	}
	if ht.KeyAdd(name.GetStr(), val) == nil {
		core.PhpErrorDocref(nil, faults.E_NOTICE, "\"%s\" is returned from __sleep multiple times", error_name.GetVal())
		return types.SUCCESS
	}
	val.TryAddRefcount()
	return types.SUCCESS
}
func PhpVarSerializeGetSleepProps(ht *types.Array, struc *types.Zval, sleep_retval *types.Array) int {
	var ce *types.ClassEntry = types.Z_OBJCE_P(struc)
	var props *types.Array = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)
	var name_val *types.Zval
	var retval int = types.SUCCESS
	ht = types.MakeArrayEx(sleep_retval.Len(), zend.ZVAL_PTR_DTOR, 0)

	/* TODO: Rewrite this by fetching the property info instead of trying out different
	 * name manglings? */

	var __ht *types.Array = sleep_retval
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		name_val = _z
		var name *types.String
		var tmp_name *types.String
		var priv_name *types.String
		var prot_name *types.String
		name_val = types.ZVAL_DEREF(name_val)
		if name_val.GetType() != types.IS_STRING {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize.")
		}
		name = zend.ZvalGetTmpString(name_val, &tmp_name)
		if PhpVarSerializeTryAddSleepProp(ht, props, name, name, struc) == types.SUCCESS {
			// zend.ZendTmpStringRelease(tmp_name)
			continue
		}
		if zend.EG__().GetException() != nil {
			// zend.ZendTmpStringRelease(tmp_name)
			retval = types.FAILURE
			break
		}
		priv_name = zend.ZendManglePropertyName_ZStr(ce.GetName().GetStr(), name.GetStr())
		if PhpVarSerializeTryAddSleepProp(ht, props, priv_name, name, struc) == types.SUCCESS {
			// zend.ZendTmpStringRelease(tmp_name)
			// types.ZendStringRelease(priv_name)
			continue
		}
		// types.ZendStringRelease(priv_name)
		if zend.EG__().GetException() != nil {
			// zend.ZendTmpStringRelease(tmp_name)
			retval = types.FAILURE
			break
		}
		prot_name = zend.ZendManglePropertyName_ZStr("*", name.GetStr())
		if PhpVarSerializeTryAddSleepProp(ht, props, prot_name, name, struc) == types.SUCCESS {
			// zend.ZendTmpStringRelease(tmp_name)
			// types.ZendStringRelease(prot_name)
			continue
		}
		// types.ZendStringRelease(prot_name)
		if zend.EG__().GetException() != nil {
			// zend.ZendTmpStringRelease(tmp_name)
			retval = types.FAILURE
			break
		}
		core.PhpErrorDocref(nil, faults.E_NOTICE, "\"%s\" returned as member variable from __sleep() but does not exist", name.GetVal())
		ht.KeyAdd(name.GetStr(), zend.EG__().GetUninitializedZval())
		// zend.ZendTmpStringRelease(tmp_name)
	}
	zend.ZendReleaseProperties(props)
	return retval
}
func PhpVarSerializeNestedData(
	buf *zend.SmartStr,
	struc *types.Zval,
	ht *types.Array,
	count uint32,
	incomplete_class types.ZendBool,
	var_hash PhpSerializeDataT,
) {
	buf.AppendUlong(count)
	buf.AppendString(":{")
	if count > 0 {
		var key *types.String
		var data *types.Zval
		var index zend.ZendUlong
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			index = _p.GetH()
			key = _p.GetKey()
			data = _z
			if incomplete_class != 0 && strcmp(key.GetVal(), MAGIC_MEMBER) == 0 {
				continue
			}
			if key == nil {
				PhpVarSerializeLong(buf, index)
			} else {
				PhpVarSerializeString(buf, key.GetVal(), key.GetLen())
			}
			if data.IsReference() && data.GetRefcount() == 1 {
				data = types.Z_REFVAL_P(data)
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */

			if data.IsType(types.IS_ARRAY) {
				if data.IsRecursive() || struc.IsType(types.IS_ARRAY) && data.Array() == struc.Array() {
					PhpAddVarHash(var_hash, struc)
					buf.AppendString("N;")
				} else {
					if data.IsRefcounted() {
						data.ProtectRecursive()
					}
					PhpVarSerializeIntern(buf, data, var_hash)
					if data.IsRefcounted() {
						data.UnprotectRecursive()
					}
				}
			} else {
				PhpVarSerializeIntern(buf, data, var_hash)
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */

		}
	}
	buf.AppendByte('}')
}
func PhpVarSerializeClass(buf *zend.SmartStr, struc *types.Zval, retval_ptr *types.Zval, var_hash PhpSerializeDataT) {
	var props types.Array
	if PhpVarSerializeGetSleepProps(&props, struc, zend.HASH_OF(retval_ptr)) == types.SUCCESS {
		PhpVarSerializeClassName(buf, struc)
		PhpVarSerializeNestedData(buf, struc, &props, props.Len(), 0, var_hash)
	}
	props.Destroy()
}
func PhpVarSerializeIntern(buf *zend.SmartStr, struc *types.Zval, var_hash PhpSerializeDataT) {
	var var_already zend.ZendLong
	var myht *types.Array
	if zend.EG__().GetException() != nil {
		return
	}
	if var_hash != nil && b.Assign(&var_already, PhpAddVarHash(var_hash, struc)) {
		if var_already == -1 {

			/* Reference to an object that failed to serialize, replace with null. */

			buf.AppendString("N;")
			return
		} else if struc.IsReference() {
			buf.AppendString("R:")
			buf.AppendLong(var_already)
			buf.AppendByte(';')
			return
		} else if struc.IsType(types.IS_OBJECT) {
			buf.AppendString("r:")
			buf.AppendLong(var_already)
			buf.AppendByte(';')
			return
		}
	}
again:
	switch struc.GetType() {
	case types.IS_FALSE:
		buf.AppendString("b:0;")
		return
	case types.IS_TRUE:
		buf.AppendString("b:1;")
		return
	case types.IS_NULL:
		buf.AppendString("N;")
		return
	case types.IS_LONG:
		PhpVarSerializeLong(buf, struc.Long())
		return
	case types.IS_DOUBLE:
		var tmp_str []byte
		buf.AppendString("d:")
		core.PhpGcvt(struc.Double(), int(core.PG__().serialize_precision), '.', 'E', tmp_str)
		buf.AppendString(b.CastStrAuto(tmp_str))
		buf.AppendByte(';')
		return
	case types.IS_STRING:
		PhpVarSerializeString(buf, struc.String().GetVal(), struc.String().GetLen())
		return
	case types.IS_OBJECT:
		var ce *types.ClassEntry = types.Z_OBJCE_P(struc)
		var incomplete_class types.ZendBool
		var count uint32
		if ce.FunctionTable().Exists("__serialize") {
			var retval types.Zval
			var obj types.Zval
			var key *types.String
			var data *types.Zval
			var index zend.ZendUlong
			struc.AddRefcount()
			obj.SetObject(struc.Object())
			if PhpVarSerializeCallMagicSerialize(&retval, &obj) == types.FAILURE {
				if zend.EG__().GetException() == nil {
					buf.AppendString("N;")
				}
				zend.ZvalPtrDtor(&obj)
				return
			}
			PhpVarSerializeClassName(buf, &obj)
			buf.AppendUlong(retval.Array().Count())
			buf.AppendString(":{")
			var __ht *types.Array = retval.Array()
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()
				if _z.IsIndirect() {
					_z = _z.Indirect()
					if _z.IsUndef() {
						continue
					}
				}
				index = _p.GetH()
				key = _p.GetKey()
				data = _z
				if key == nil {
					PhpVarSerializeLong(buf, index)
				} else {
					PhpVarSerializeString(buf, key.GetVal(), key.GetLen())
				}
				if data.IsReference() && data.GetRefcount() == 1 {
					data = types.Z_REFVAL_P(data)
				}
				PhpVarSerializeIntern(buf, data, var_hash)
			}
			buf.AppendByte('}')
			zend.ZvalPtrDtor(&obj)
			zend.ZvalPtrDtor(&retval)
			return
		}
		if ce.GetSerialize() != nil {

			/* has custom handler */

			var serialized_data *uint8 = nil
			var serialized_length int
			if ce.GetSerialize()(struc, &serialized_data, &serialized_length, (*zend.ZendSerializeData)(var_hash)) == types.SUCCESS {
				buf.AppendString("C:")
				buf.AppendUlong(types.Z_OBJCE_P(struc).GetName().GetLen())
				buf.AppendString(":\"")
				buf.AppendString(types.Z_OBJCE_P(struc).GetName().GetStr())
				buf.AppendString("\":")
				buf.AppendUlong(serialized_length)
				buf.AppendString(":{")
				buf.AppendString(b.CastStr((*byte)(serialized_data), serialized_length))
				buf.AppendByte('}')
			} else {

				/* Mark this value in the var_hash, to avoid creating references to it. */

				var var_idx *types.Zval = var_hash.GetHt().IndexFindH(zend.ZendUlong(types.ZendUintptrT(struc.RefCounted())))
				var_idx.SetLong(-1)
				buf.AppendString("N;")
			}
			if serialized_data != nil {
				zend.Efree(serialized_data)
			}
			return
		}
		if ce != PHP_IC_ENTRY && ce.FunctionTable().Exists("__sleep") {
			var retval types.Zval
			var tmp types.Zval
			struc.AddRefcount()
			tmp.SetObject(struc.Object())
			if PhpVarSerializeCallSleep(&retval, &tmp) == types.FAILURE {
				if zend.EG__().GetException() == nil {

					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */

					buf.AppendString("N;")

					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */

				}
				zend.ZvalPtrDtor(&tmp)
				return
			}
			PhpVarSerializeClass(buf, &tmp, &retval, var_hash)
			zend.ZvalPtrDtor(&retval)
			zend.ZvalPtrDtor(&tmp)
			return
		}
		incomplete_class = PhpVarSerializeClassName(buf, struc)
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)

		/* count after serializing name, since php_var_serialize_class_name
		 * changes the count if the variable is incomplete class */

		count = myht.Count()
		if count > 0 && incomplete_class != 0 {
			count--
		}
		PhpVarSerializeNestedData(buf, struc, myht, count, incomplete_class, var_hash)
		zend.ZendReleaseProperties(myht)
		return
	case types.IS_ARRAY:
		buf.AppendString("a:")
		myht = struc.Array()
		PhpVarSerializeNestedData(buf, struc, myht, myht.Count(), 0, var_hash)
		return
	case types.IS_REFERENCE:
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		buf.AppendString("i:0;")
		return
	}
}
func PhpVarSerialize(buf *zend.SmartStr, struc *types.Zval, data *PhpSerializeDataT) {
	PhpVarSerializeIntern(buf, struc, *data)
	buf.ZeroTail()
}
func PhpVarSerializeInit() PhpSerializeDataT {
	var d *PhpSerializeData

	/* fprintf(stderr, "SERIALIZE_INIT      == lock: %u, level: %u\n", BG__().serialize_lock, BG__().serialize.level); */

	if BG__().serialize_lock || !(BG__().serialize.level) {
		d = zend.Emalloc(b.SizeOf("struct php_serialize_data"))
		d.GetHt() = types.MakeArrayEx(16, zend.ZVAL_PTR_DTOR, 0)
		d.SetN(0)
		if !(BG__().serialize_lock) {
			BG__().serialize.data = d
			BG__().serialize.level = 1
		}
	} else {
		d = BG__().serialize.data
		BG__().serialize.level++
	}
	return d
}
func PhpVarSerializeDestroy(d PhpSerializeDataT) {
	/* fprintf(stderr, "SERIALIZE_DESTROY   == lock: %u, level: %u\n", BG__().serialize_lock, BG__().serialize.level); */

	if BG__().serialize_lock || BG__().serialize.level == 1 {
		d.GetHt().Destroy()
		zend.Efree(d)
	}
	if !(BG__().serialize_lock) && !(b.PreDec(&(BG__().serialize.level))) {
		BG__().serialize.data = nil
	}
}
func ZifSerialize(executeData zpp.Ex, return_value zpp.Ret, var_ *types.Zval) {
	var struc *types.Zval
	var var_hash PhpSerializeDataT
	var buf zend.SmartStr = zend.MakeSmartStr(0)
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			struc = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	PHP_VAR_SERIALIZE_INIT(var_hash)
	PhpVarSerialize(&buf, struc, &var_hash)
	PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if zend.EG__().GetException() != nil {
		buf.Free()
		return_value.SetFalse()
		return
	}
	if buf.GetS() != nil {
		return_value.SetString(buf.GetS())
		return
	} else {
		return_value.SetNull()
		return
	}
}
func ZifUnserialize(executeData zpp.Ex, return_value zpp.Ret, variableRepresentation *types.Zval, _ zpp.Opt, allowedClasses *types.Zval) {
	var buf *byte = nil
	var buf_len int
	var p *uint8
	var var_hash PhpUnserializeDataT
	var options *types.Zval = nil
	var retval *types.Zval
	var class_hash *types.Array = nil
	var prev_class_hash *types.Array
	var prev_max_depth zend.ZendLong
	var prev_cur_depth zend.ZendLong
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			buf, buf_len = fp.ParseString()
			fp.StartOptional()
			options = fp.ParseArray()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if buf_len == 0 {
		return_value.SetFalse()
		return
	}
	p = (*uint8)(buf)
	PHP_VAR_UNSERIALIZE_INIT(var_hash)
	prev_class_hash = var_hash.GetAllowedClasses()
	prev_max_depth = var_hash.GetMaxDepth()
	prev_cur_depth = var_hash.GetCurDepth()
	if options != nil {
		var classes *types.Zval
		var max_depth *types.Zval
		classes = types.ZendHashStrFindDeref(options.Array(), "allowed_classes")
		if classes != nil && classes.GetType() != types.IS_ARRAY && classes.GetType() != types.IS_TRUE && classes.GetType() != types.IS_FALSE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "allowed_classes option should be array or boolean")
			return_value.SetFalse()
			goto cleanup
		}
		if classes != nil && (classes.IsType(types.IS_ARRAY) || zend.ZendIsTrue(classes) == 0) {
			zend.ALLOC_HASHTABLE(class_hash)
			class_hash = types.MakeArrayEx(b.CondF1(classes.IsType(types.IS_ARRAY), func() __auto__ { return classes.Array().Len() }, 0), nil, 0)
		}
		if class_hash != nil && classes.IsType(types.IS_ARRAY) {
			var entry *types.Zval
			var lcname *types.String
			var __ht *types.Array = classes.Array()
			for _, _p := range __ht.ForeachData() {
				var _z *types.Zval = _p.GetVal()

				entry = _z
				zend.ConvertToStringEx(entry)
				lcname = zend.ZendStringTolower(entry.String())
				types.ZendHashAddEmptyElement(class_hash, lcname.GetStr())
				// types.ZendStringReleaseEx(lcname, 0)
			}

			/* Exception during string conversion. */

			if zend.EG__().GetException() != nil {
				goto cleanup
			}

			/* Exception during string conversion. */

		}
		var_hash.SetAllowedClasses(class_hash)
		max_depth = types.ZendHashStrFindDeref(options.Array(), "max_depth")
		if max_depth != nil {
			if max_depth.GetType() != types.IS_LONG {
				core.PhpErrorDocref(nil, faults.E_WARNING, "max_depth should be int")
				return_value.SetFalse()
				goto cleanup
			}
			if max_depth.Long() < 0 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "max_depth cannot be negative")
				return_value.SetFalse()
				goto cleanup
			}
			var_hash.SetMaxDepth(max_depth.Long())

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

			var_hash.SetCurDepth(0)

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

		}
	}
	if BG__().unserialize.level > 1 {
		retval = VarTmpVar(&var_hash)
	} else {
		retval = return_value
	}
	if PhpVarUnserialize(retval, &p, p+buf_len, &var_hash) == 0 {
		if zend.EG__().GetException() == nil {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "Error at offset "+zend.ZEND_LONG_FMT+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
		}
		if BG__().unserialize.level <= 1 {
			zend.ZvalPtrDtor(return_value)
		}
		return_value.SetFalse()
	} else if BG__().unserialize.level > 1 {
		types.ZVAL_COPY(return_value, retval)
	} else if return_value.IsRefcounted() {
		//var ref *types.ZendRefcounted = return_value.GetCounted()
		//zend.GcCheckPossibleRoot(ref)
	}
cleanup:
	if class_hash != nil {
		class_hash.Destroy()
		zend.FREE_HASHTABLE(class_hash)
	}

	/* Reset to previous options in case this is a nested call */

	var_hash.SetAllowedClasses(prev_class_hash)
	var_hash.SetMaxDepth(prev_max_depth)
	var_hash.SetCurDepth(prev_cur_depth)
	PHP_VAR_UNSERIALIZE_DESTROY(var_hash)

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */

	if return_value.IsReference() {
		zend.ZendUnwrapReference(return_value)
	}

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */
}
func ZifMemoryGetUsage(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, realUsage *types.Zval) {
	var real_usage types.ZendBool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			real_usage = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(zend.ZendMemoryUsage(real_usage))
	return
}
func ZifMemoryGetPeakUsage(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, realUsage *types.Zval) {
	var real_usage types.ZendBool = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			real_usage = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	return_value.SetLong(zend.ZendMemoryPeakUsage(real_usage))
	return
}
func ZmStartupVar(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES(module_number)
	return types.SUCCESS
}
