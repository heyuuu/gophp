package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func PhpArrayElementDump(zv *types.Zval, key types.ArrayKey, level int) {
	if key.IsStrKey() {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.OG__().WriteString(key.StrKey())
		core.PhpPrintf("\"]=>\n")
	} else {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key.IdxKey())
	}
	PhpVarDump(zv, level+2)
}
func PhpObjectPropertyDump(propInfo *types.PropertyInfo, zv *types.Zval, key_ types.ArrayKey, level int) {
	if !key_.IsStrKey() {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key_.IdxKey())
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
			core.OG__().WriteString(key_.StrKey())
			core.PhpPrintf("\"")
		}
		zend.ZEND_PUTS("]=>\n")
	}
	if zv.IsUndef() {
		b.Assert(propInfo.GetType() != nil)

		typ := propInfo.GetType().FormatType()
		core.PhpPrintf("%*cuninitialized(%s)\n", level+1, ' ', typ)
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
	switch struc.Type() {
	case types.IsFalse:
		core.PhpPrintf("%sbool(false)\n", common)
	case types.IsTrue:
		core.PhpPrintf("%sbool(true)\n", common)
	case types.IsNull:
		core.PhpPrintf("%sNULL\n", common)
	case types.IsLong:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", common, struc.Long())
	case types.IsDouble:
		core.PhpPrintf("%sfloat(%.*G)\n", common, int(zend.EG__().GetPrecision()), struc.Double())
	case types.IsString:
		core.PhpPrintf("%sstring(%zd) \"", common, struc.StringEx().GetLen())
		core.OG__().WriteString(struc.String())
		core.PUTS("\"\n")
	case types.IsArray:
		myht := struc.Array()
		if level > 1 {
			if myht.IsRecursive() {
				core.PUTS("*RECURSION*\n")
				return
			}
			myht.ProtectRecursive()
		}
		count := myht.Count()
		core.PhpPrintf("%sarray(%d) {\n", common, count)
		myht.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			PhpArrayElementDump(value, key, level)
		})
		if level > 1 {
			myht.UnprotectRecursive()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IsObject:
		if struc.Object().IsRecursive() {
			core.PUTS("*RECURSION*\n")
			return
		}
		struc.Object().ProtectRecursive()
		myht := zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		className := struc.Object().ClassName()
		core.PhpPrintf("%sobject(%s)#%d (%d) {\n", common, className, struc.Object().GetHandle(), lang.CondF1(myht != nil, func() int { return myht.Count() }, 0))
		if myht != nil {
			myht.Foreach(func(key types.ArrayKey, value *types.Zval) {
				var prop_info *types.PropertyInfo = nil
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
		struc.Object().UnprotectRecursive()
	case types.IsResource:
		typeName := b.Option(zend.ZendRsrcListGetRsrcTypeEx(struc.Resource()), "Unknown")
		core.PhpPrintf("%sresource(%d) of type (%s)\n", common, struc.ResourceHandle(), typeName)
	case types.IsRef:
		isRef = true
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
func ZvalArrayElementDump(zv *types.Zval, key types.ArrayKey, level int) {
	if !key.IsStrKey() {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key.IdxKey())
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PUTS(key.StrKey())
		core.PUTS("\"]=>\n")
	}
	PhpDebugZvalDump(zv, level+2)
}
func ZvalObjectPropertyDump(propInfo *types.PropertyInfo, zv *types.Zval, key types.ArrayKey, level int) {
	if !key.IsStrKey() {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', key.IdxKey())
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
		b.Assert(propInfo.GetType() != nil)

		typ := propInfo.GetType().FormatType()
		core.PhpPrintf("%*cuninitialized(%s)\n", level+1, ' ', typ)
	} else {
		PhpDebugZvalDump(zv, level+2)
	}
}
func PhpDebugZvalDump(struc *types.Zval, level int) {
	if level > 1 {
		core.PhpPrintf("%*c", level-1, ' ')
	}

	// deref
	common := ""
	if struc.IsRef() {
		common = "&"
		struc = struc.DeRef()
	}

	switch struc.Type() {
	case types.IsFalse:
		core.PhpPrintf("%sbool(false)\n", common)
	case types.IsTrue:
		core.PhpPrintf("%sbool(true)\n", common)
	case types.IsNull:
		core.PhpPrintf("%sNULL\n", common)
	case types.IsLong:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", common, struc.Long())
	case types.IsDouble:
		core.PhpPrintf("%sfloat(%.*G)\n", common, int(zend.EG__().GetPrecision()), struc.Double())
	case types.IsString:
		core.PhpPrintf("%sstring(%zd) \"", common, struc.StringEx().GetLen())
		core.PUTS(struc.String())
		core.PhpPrintf("\"\n")
	case types.IsArray:
		myht := struc.Array()
		if level > 1 {
			if myht.IsRecursive() {
				core.PUTS("*RECURSION*\n")
				return
			}
			myht.ProtectRecursive()
		}
		count := myht.Count()
		core.PhpPrintf("%sarray(%d){\n", common, count)
		myht.ForeachIndirect(func(key types.ArrayKey, value *types.Zval) {
			ZvalArrayElementDump(value, key, level)
		})
		if level > 1 {
			myht.UnprotectRecursive()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IsObject:
		myht := zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		if myht != nil {
			if myht.IsRecursive() {
				core.PUTS("*RECURSION*\n")
				//zend.ZendReleaseProperties(myht)
				return
			}
			myht.ProtectRecursive()
		}
		className := struc.Object().ClassName()
		core.PhpPrintf("%sobject(%s)#%d (%d) {\n", common, className, struc.Object().GetHandle(), lang.CondF1(myht != nil, func() uint32 { return myht.Count() }, 0))
		// types.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			myht.Foreach(func(key types.ArrayKey, value *types.Zval) {
				var propInfo *types.PropertyInfo = nil
				if value.IsIndirect() {
					value = value.Indirect()
					if key.IsStrKey() {
						propInfo = zend.ZendGetTypedPropertyInfoForSlot(struc.Object(), value)
					}
				}
				if !value.IsUndef() || propInfo != nil {
					ZvalObjectPropertyDump(propInfo, value, key, level)
				}
			})
			myht.UnprotectRecursive()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
	case types.IsResource:
		typeName := b.Option(zend.ZendRsrcListGetRsrcTypeEx(struc.Resource()), "Unknown")
		core.PhpPrintf("%sresource(%d) of type (%s)\n", common, struc.ResourceHandle(), typeName)
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", common)
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
	buf.WriteString(b.CastStr(tmp_spaces, tmp_spaces_len))
	zend.Efree(tmp_spaces)
}
func PhpArrayElementExport(zv *types.Zval, index zend.ZendUlong, key *types.String, level int, buf *zend.SmartStr) {
	if key == nil {
		BufferAppendSpaces(buf, level+1)
		buf.WriteLong(zend.ZendLong(index))
		buf.WriteString(" => ")
	} else {
		var ckey = types.NewString(str.PhpAddcslashes(key.GetStr(), "'\\"))
		tmp_str := strings.ReplaceAll(ckey.GetStr(), "0", "' . \"\\0\" . '")
		BufferAppendSpaces(buf, level+1)
		buf.WriteByte('\'')
		buf.WriteString(tmp_str)
		buf.WriteString("' => ")
		//types.ZendStringFree(ckey)
		//types.ZendStringFree(tmp_str)
	}
	PhpVarExportEx(zv, level+2, buf)
	buf.WriteByte(',')
	buf.WriteByte('\n')
}
func PhpObjectElementExport(zv *types.Zval, index zend.ZendUlong, key *types.String, level int, buf *zend.SmartStr) {
	BufferAppendSpaces(buf, level+2)
	if key != nil {
		_, propName, _ := zend.ZendUnmanglePropertyName_Ex(key.GetStr())

		propNameEscaped := str.PhpAddcslashes(propName, "'\\")
		buf.WriteByte('\'')
		buf.WriteString(propNameEscaped)
		buf.WriteByte('\'')
	} else {
		buf.WriteLong(zend.ZendLong(index))
	}
	buf.WriteString(" => ")
	PhpVarExportEx(zv, level+2, buf)
	buf.WriteByte(',')
	buf.WriteByte('\n')
}
func PhpVarExportEx(struc *types.Zval, level int, buf *zend.SmartStr) {
	var myht *types.Array
	var tmp_str []byte
	var index zend.ZendUlong
	var key *types.String
	var val *types.Zval
again:
	switch struc.Type() {
	case types.IsFalse:
		buf.WriteString("false")
	case types.IsTrue:
		buf.WriteString("true")
	case types.IsNull:
		buf.WriteString("NULL")
	case types.IsLong:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if struc.Long() == zend.ZEND_LONG_MIN {
			buf.WriteLong(zend.ZEND_LONG_MIN + 1)
			buf.WriteString("-1")
			break
		}
		buf.WriteLong(struc.Long())
	case types.IsDouble:
		core.PhpGcvt(struc.Double(), int(core.PG__().serialize_precision), '.', 'E', tmp_str)
		buf.WriteString(b.CastStrAuto(tmp_str))

		/* Without a decimal point, PHP treats a number literal as an int.
		 * This check even works for scientific notation, because the
		 * mantissa always contains a decimal point.
		 * We need to check for finiteness, because INF, -INF and NAN
		 * must not have a decimal point added.
		 */

		if core.ZendFinite(struc.Double()) && nil == strchr(tmp_str, '.') {
			buf.WriteString(".0")
		}
	case types.IsString:
		ztmp := str.PhpAddcslashes(struc.String(), "'\\")
		ztmp2 := strings.ReplaceAll(ztmp, "0", "' . \"\\0\" . '")
		buf.WriteByte('\'')
		buf.WriteString(ztmp2)
		buf.WriteByte('\'')
		//types.ZendStringFree(ztmp)
		//types.ZendStringFree(ztmp2)
	case types.IsArray:
		myht = struc.Array()
		if myht.IsRecursive() {
			buf.WriteString("NULL")
			faults.Error(faults.E_WARNING, "var_export does not handle circular references")
			return
		}
		myht.ProtectRecursive()
		if level > 1 {
			buf.WriteByte('\n')
			BufferAppendSpaces(buf, level-1)
		}
		buf.WriteString("array (\n")
		var __ht = myht
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
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
		myht.UnprotectRecursive()
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		buf.WriteByte(')')
	case types.IsObject:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_VAR_EXPORT)
		if myht != nil {
			if myht.IsRecursive() {
				buf.WriteString("NULL")
				faults.Error(faults.E_WARNING, "var_export does not handle circular references")
				//zend.ZendReleaseProperties(myht)
				return
			} else {
				myht.ProtectRecursive()
			}
		}
		if level > 1 {
			buf.WriteByte('\n')
			BufferAppendSpaces(buf, level-1)
		}

		/* stdClass has no __set_state method, but can be casted to */

		if types.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			buf.WriteString("(object) array(\n")
		} else {
			buf.WriteString(types.Z_OBJCE_P(struc).Name())
			buf.WriteString("::__set_state(array(\n")
		}
		if myht != nil {
			var __ht = myht
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()
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
			myht.UnprotectRecursive()
			//zend.ZendReleaseProperties(myht)
		}
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		if types.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			buf.WriteByte(')')
		} else {
			buf.WriteString("))")
		}
	case types.IsRef:
		struc = types.Z_REFVAL_P(struc)
		goto again
	default:
		buf.WriteString("NULL")
	}
}
func ZifVarExport(executeData zpp.Ex, return_value zpp.Ret, var__ *types.Zval, _ zpp.Opt, return_ *types.Zval) {
	var var_ *types.Zval
	var return_output = 0
	var buf zend.SmartStr
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
	//buf.ZeroTail()
	if return_output != 0 {
		return_value.SetStringEx(buf.GetS())
		return
	} else {
		core.PUTS(buf.GetS().GetStr())
		buf.Free()
	}
}

func PhpVarSerializeInit() PhpSerializeDataT {
	var d *PhpSerializeData

	if BG__().serialize_lock != 0 {
		d = NewPhpSerializeData()
	} else if BG__().serialize.level == 0 {
		d = NewPhpSerializeData()
		BG__().serialize.data = d
		BG__().serialize.level = 1
	} else {
		d = BG__().serialize.data
		BG__().serialize.level++
	}
	return d
}
func PhpVarSerializeDestroy(d PhpSerializeDataT) {
	if BG__().serialize_lock != 0 || BG__().serialize.level == 1 {
		d.Destroy()
	}
	if BG__().serialize_lock == 0 {
		BG__().serialize.level--
		if BG__().serialize.level == 0 {
			BG__().serialize.data = nil
		}
	}
}
func ZifSerialize(var_ *types.Zval) *types.Zval {
	serializer := InitSerializer()
	serializer.Serialize(var_)
	serializer.DestroyData()

	if zend.EG__().HasException() {
		return types.NewZvalFalse()
	}

	serializerStr := serializer.String()
	if serializerStr != "" {
		return types.NewZvalString(serializerStr)
	} else {
		return types.NewZvalNull()
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
		if classes != nil && !classes.IsArray() && !classes.IsTrue() && !classes.IsFalse() {
			core.PhpErrorDocref("", faults.E_WARNING, "allowed_classes option should be array or boolean")
			return_value.SetFalse()
			goto cleanup
		}
		if classes != nil && (classes.IsType(types.IsArray) || !operators.ZvalIsTrue(classes)) {
			if classes.IsArray() {
				class_hash = types.NewArrayCap(classes.Array().Len())
			} else {
				class_hash = types.NewArray()
			}
		}
		if class_hash != nil && classes.IsType(types.IsArray) {
			var entry *types.Zval
			var lcname *types.String
			var __ht = classes.Array()
			for _, _p := range __ht.ForeachData() {
				var _z = _p.GetVal()

				entry = _z
				operators.ConvertToStringEx(entry)
				lcname = operators.ZendStringTolower(entry.StringEx())
				types.ZendHashAddEmptyElement(class_hash, lcname.GetStr())
				// types.ZendStringReleaseEx(lcname, 0)
			}

			/* Exception during string conversion. */

			if zend.EG__().HasException() {
				goto cleanup
			}
		}
		var_hash.SetAllowedClasses(class_hash)
		max_depth = types.ZendHashStrFindDeref(options.Array(), "max_depth")
		if max_depth != nil {
			if !max_depth.IsLong() {
				core.PhpErrorDocref("", faults.E_WARNING, "max_depth should be int")
				return_value.SetFalse()
				goto cleanup
			}
			if max_depth.Long() < 0 {
				core.PhpErrorDocref("", faults.E_WARNING, "max_depth cannot be negative")
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
		if zend.EG__().NoException() {
			core.PhpErrorDocref("", faults.E_NOTICE, "Error at offset "+zend.ZEND_LONG_FMT+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
		}
		if BG__().unserialize.level <= 1 {
			// zend.ZvalPtrDtor(return_value)
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
	}

	/* Reset to previous options in case this is a nested call */

	var_hash.SetAllowedClasses(prev_class_hash)
	var_hash.SetMaxDepth(prev_max_depth)
	var_hash.SetCurDepth(prev_cur_depth)
	PHP_VAR_UNSERIALIZE_DESTROY(var_hash)

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */

	if return_value.IsRef() {
		operators.ZendUnwrapReference(return_value)
	}

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */
}
func ZifMemoryGetUsage(_ zpp.Opt, realUsage bool) int {
	return zend.ZendMemoryUsage(realUsage)
}
func ZifMemoryGetPeakUsage(_ zpp.Opt, realUsage bool) int {
	return zend.ZendMemoryPeakUsage(realUsage)
}
func ZmStartupVar(moduleNumber int) int {
	zend.REGISTER_INI_ENTRIES(moduleNumber)
	return types.SUCCESS
}
