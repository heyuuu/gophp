// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func PhpArrayElementDump(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	if key == nil {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', index)
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PHPWRITE(key.GetVal(), key.GetLen())
		core.PhpPrintf("\"]=>\n")
	}
	PhpVarDump(zv, level+2)
}
func PhpObjectPropertyDump(prop_info *zend.ZendPropertyInfo, zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	var prop_name *byte
	var class_name *byte
	if key == nil {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', index)
	} else {
		var unmangle int = zend.ZendUnmanglePropertyName(key, &class_name, &prop_name)
		core.PhpPrintf("%*c[", level+1, ' ')
		if class_name != nil && unmangle == zend.SUCCESS {
			if class_name[0] == '*' {
				core.PhpPrintf("\"%s\":protected", prop_name)
			} else {
				core.PhpPrintf("\"%s\":\"%s\":private", prop_name, class_name)
			}
		} else {
			core.PhpPrintf("\"")
			core.PHPWRITE(key.GetVal(), key.GetLen())
			core.PhpPrintf("\"")
		}
		zend.ZEND_PUTS("]=>\n")
	}
	if zv.IsType(zend.IS_UNDEF) {
		zend.ZEND_ASSERT(prop_info.GetType() != 0)
		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', b.Cond(prop_info.GetType().AllowNull(), "?", ""), b.CondF(prop_info.GetType().IsClass(), func() []byte {
			return b.CondF(prop_info.GetType().IsCe(), func() *zend.ZendString { return zend.ZEND_TYPE_CE(prop_info.GetType()).GetName() }, func() *zend.ZendString { return prop_info.GetType().Name() }).GetVal()
		}, func() *byte { return zend.ZendGetTypeByConst(prop_info.GetType().Code()) }))
	} else {
		PhpVarDump(zv, level+2)
	}
}
func PhpVarDump(struc *zend.Zval, level int) {
	var myht *zend.HashTable
	var class_name *zend.ZendString
	var is_ref int = 0
	var num zend.ZendUlong
	var key *zend.ZendString
	var val *zend.Zval
	var count uint32
	if level > 1 {
		core.PhpPrintf("%*c", level-1, ' ')
	}
again:
	switch struc.GetType() {
	case zend.IS_FALSE:
		core.PhpPrintf("%sbool(false)\n", COMMON)
		break
	case zend.IS_TRUE:
		core.PhpPrintf("%sbool(true)\n", COMMON)
		break
	case zend.IS_NULL:
		core.PhpPrintf("%sNULL\n", COMMON)
		break
	case zend.IS_LONG:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", COMMON, struc.GetLval())
		break
	case zend.IS_DOUBLE:
		core.PhpPrintf("%sfloat(%.*G)\n", COMMON, int(zend.__EG().GetPrecision()), struc.GetDval())
		break
	case zend.IS_STRING:
		core.PhpPrintf("%sstring(%zd) \"", COMMON, zend.Z_STRLEN_P(struc))
		core.PHPWRITE(zend.Z_STRVAL_P(struc), zend.Z_STRLEN_P(struc))
		core.PUTS("\"\n")
		break
	case zend.IS_ARRAY:
		myht = struc.GetArr()
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			if level > 1 {
				if zend.GC_IS_RECURSIVE(myht) != 0 {
					core.PUTS("*RECURSION*\n")
					return
				}
				zend.GC_PROTECT_RECURSION(myht)
			}
			myht.AddRefcount()
		}
		count = myht.Count()
		core.PhpPrintf("%sarray(%d) {\n", COMMON, count)
		var __ht *zend.HashTable = myht
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()
			if _z.IsType(zend.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
			}
			num = _p.GetH()
			key = _p.GetKey()
			val = _z
			PhpArrayElementDump(val, num, key, level)
		}
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			if level > 1 {
				zend.GC_UNPROTECT_RECURSION(myht)
			}
			myht.DelRefcount()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
		break
	case zend.IS_OBJECT:
		if zend.Z_IS_RECURSIVE_P(struc) != 0 {
			core.PUTS("*RECURSION*\n")
			return
		}
		zend.Z_PROTECT_RECURSION_P(struc)
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		class_name = zend.Z_OBJ_HT(*struc).GetGetClassName()(struc.GetObj())
		core.PhpPrintf("%sobject(%s)#%d (%d) {\n", COMMON, class_name.GetVal(), zend.Z_OBJ_HANDLE_P(struc), b.CondF1(myht != nil, func() uint32 { return myht.Count() }, 0))
		zend.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			var num zend.ZendUlong
			var key *zend.ZendString
			var val *zend.Zval
			var __ht *zend.HashTable = myht
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()

				num = _p.GetH()
				key = _p.GetKey()
				val = _z
				var prop_info *zend.ZendPropertyInfo = nil
				if val.IsType(zend.IS_INDIRECT) {
					val = val.GetZv()
					if key != nil {
						prop_info = zend.ZendGetTypedPropertyInfoForSlot(struc.GetObj(), val)
					}
				}
				if !(zend.Z_ISUNDEF_P(val)) || prop_info != nil {
					PhpObjectPropertyDump(prop_info, val, num, key, level)
				}
			}
			zend.ZendReleaseProperties(myht)
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
		zend.Z_UNPROTECT_RECURSION_P(struc)
		break
	case zend.IS_RESOURCE:
		var type_name *byte = zend.ZendRsrcListGetRsrcType(struc.GetRes())
		core.PhpPrintf("%sresource(%d) of type (%s)\n", COMMON, zend.Z_RES_P(struc).GetHandle(), b.Cond(type_name != nil, type_name, "Unknown"))
		break
	case zend.IS_REFERENCE:

		//??? hide references with refcount==1 (for compatibility)

		if zend.Z_REFCOUNT_P(struc) > 1 {
			is_ref = 1
		}
		struc = zend.Z_REFVAL_P(struc)
		goto again
		break
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", COMMON)
		break
	}
}
func ZifVarDump(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var argc int
	var i int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	for i = 0; i < argc; i++ {
		PhpVarDump(&args[i], 1)
	}
}
func ZvalArrayElementDump(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	if key == nil {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', index)
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PHPWRITE(key.GetVal(), key.GetLen())
		core.PhpPrintf("\"]=>\n")
	}
	PhpDebugZvalDump(zv, level+2)
}
func ZvalObjectPropertyDump(prop_info *zend.ZendPropertyInfo, zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	var prop_name *byte
	var class_name *byte
	if key == nil {
		core.PhpPrintf("%*c["+zend.ZEND_LONG_FMT+"]=>\n", level+1, ' ', index)
	} else {
		zend.ZendUnmanglePropertyName(key, &class_name, &prop_name)
		core.PhpPrintf("%*c[", level+1, ' ')
		if class_name != nil {
			if class_name[0] == '*' {
				core.PhpPrintf("\"%s\":protected", prop_name)
			} else {
				core.PhpPrintf("\"%s\":\"%s\":private", prop_name, class_name)
			}
		} else {
			core.PhpPrintf("\"%s\"", prop_name)
		}
		zend.ZEND_PUTS("]=>\n")
	}
	if prop_info != nil && zv.IsType(zend.IS_UNDEF) {
		zend.ZEND_ASSERT(prop_info.GetType() != 0)
		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', b.Cond(prop_info.GetType().AllowNull(), "?", ""), b.CondF(prop_info.GetType().IsClass(), func() []byte {
			return b.CondF(prop_info.GetType().IsCe(), func() *zend.ZendString { return zend.ZEND_TYPE_CE(prop_info.GetType()).GetName() }, func() *zend.ZendString { return prop_info.GetType().Name() }).GetVal()
		}, func() *byte { return zend.ZendGetTypeByConst(prop_info.GetType().Code()) }))
	} else {
		PhpDebugZvalDump(zv, level+2)
	}
}
func PhpDebugZvalDump(struc *zend.Zval, level int) {
	var myht *zend.HashTable = nil
	var class_name *zend.ZendString
	var is_ref int = 0
	var index zend.ZendUlong
	var key *zend.ZendString
	var val *zend.Zval
	var count uint32
	if level > 1 {
		core.PhpPrintf("%*c", level-1, ' ')
	}
again:
	switch struc.GetType() {
	case zend.IS_FALSE:
		core.PhpPrintf("%sbool(false)\n", COMMON)
		break
	case zend.IS_TRUE:
		core.PhpPrintf("%sbool(true)\n", COMMON)
		break
	case zend.IS_NULL:
		core.PhpPrintf("%sNULL\n", COMMON)
		break
	case zend.IS_LONG:
		core.PhpPrintf("%sint("+zend.ZEND_LONG_FMT+")\n", COMMON, struc.GetLval())
		break
	case zend.IS_DOUBLE:
		core.PhpPrintf("%sfloat(%.*G)\n", COMMON, int(zend.__EG().GetPrecision()), struc.GetDval())
		break
	case zend.IS_STRING:
		core.PhpPrintf("%sstring(%zd) \"", COMMON, zend.Z_STRLEN_P(struc))
		core.PHPWRITE(zend.Z_STRVAL_P(struc), zend.Z_STRLEN_P(struc))
		core.PhpPrintf("\" refcount(%u)\n", b.CondF1(zend.Z_REFCOUNTED_P(struc), func() uint32 { return zend.Z_REFCOUNT_P(struc) }, 1))
		break
	case zend.IS_ARRAY:
		myht = struc.GetArr()
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			if level > 1 {
				if zend.GC_IS_RECURSIVE(myht) != 0 {
					core.PUTS("*RECURSION*\n")
					return
				}
				zend.GC_PROTECT_RECURSION(myht)
			}
			myht.AddRefcount()
		}
		count = myht.Count()
		core.PhpPrintf("%sarray(%d) refcount(%u){\n", COMMON, count, b.CondF1(zend.Z_REFCOUNTED_P(struc), func() int { return zend.Z_REFCOUNT_P(struc) - 1 }, 1))
		var __ht *zend.HashTable = myht
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()
			if _z.IsType(zend.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
			}
			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			ZvalArrayElementDump(val, index, key, level)
		}
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			if level > 1 {
				zend.GC_UNPROTECT_RECURSION(myht)
			}
			myht.DelRefcount()
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
		break
	case zend.IS_OBJECT:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		if myht != nil {
			if zend.GC_IS_RECURSIVE(myht) != 0 {
				core.PUTS("*RECURSION*\n")
				zend.ZendReleaseProperties(myht)
				return
			}
			zend.GC_PROTECT_RECURSION(myht)
		}
		class_name = zend.Z_OBJ_HT(*struc).GetGetClassName()(struc.GetObj())
		core.PhpPrintf("%sobject(%s)#%d (%d) refcount(%u){\n", COMMON, class_name.GetVal(), zend.Z_OBJ_HANDLE_P(struc), b.CondF1(myht != nil, func() uint32 { return myht.Count() }, 0), zend.Z_REFCOUNT_P(struc))
		zend.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			var __ht *zend.HashTable = myht
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()

				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				var prop_info *zend.ZendPropertyInfo = nil
				if val.IsType(zend.IS_INDIRECT) {
					val = val.GetZv()
					if key != nil {
						prop_info = zend.ZendGetTypedPropertyInfoForSlot(struc.GetObj(), val)
					}
				}
				if !(zend.Z_ISUNDEF_P(val)) || prop_info != nil {
					ZvalObjectPropertyDump(prop_info, val, index, key, level)
				}
			}
			zend.GC_UNPROTECT_RECURSION(myht)
			zend.ZendReleaseProperties(myht)
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		core.PUTS("}\n")
		break
	case zend.IS_RESOURCE:
		var type_name *byte = zend.ZendRsrcListGetRsrcType(struc.GetRes())
		core.PhpPrintf("%sresource(%d) of type (%s) refcount(%u)\n", COMMON, zend.Z_RES_P(struc).GetHandle(), b.Cond(type_name != nil, type_name, "Unknown"), zend.Z_REFCOUNT_P(struc))
		break
	case zend.IS_REFERENCE:

		//??? hide references with refcount==1 (for compatibility)

		if zend.Z_REFCOUNT_P(struc) > 1 {
			is_ref = 1
		}
		struc = zend.Z_REFVAL_P(struc)
		goto again
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", COMMON)
		break
	}
}
func ZifDebugZvalDump(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var argc int
	var i int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			var _num_varargs int = _num_args - _i - 0
			if _num_varargs > 0 {
				args = _real_arg + 1
				argc = _num_varargs
				_i += _num_varargs
				_real_arg += _num_varargs
			} else {
				args = nil
				argc = 0
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	for i = 0; i < argc; i++ {
		PhpDebugZvalDump(&args[i], 1)
	}
}
func BufferAppendSpaces(buf *zend.SmartStr, num_spaces int) {
	var tmp_spaces *byte
	var tmp_spaces_len int
	tmp_spaces_len = core.Spprintf(&tmp_spaces, 0, "%*c", num_spaces, ' ')
	zend.SmartStrAppendl(buf, tmp_spaces, tmp_spaces_len)
	zend.Efree(tmp_spaces)
}
func PhpArrayElementExport(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int, buf *zend.SmartStr) {
	if key == nil {
		BufferAppendSpaces(buf, level+1)
		zend.SmartStrAppendLong(buf, zend.ZendLong(index))
		zend.SmartStrAppendl(buf, " => ", 4)
	} else {
		var tmp_str *zend.ZendString
		var ckey *zend.ZendString = PhpAddcslashes(key, "'\\", 2)
		tmp_str = PhpStrToStr(ckey.GetVal(), ckey.GetLen(), "0", 1, "' . \"\\0\" . '", 12)
		BufferAppendSpaces(buf, level+1)
		zend.SmartStrAppendc(buf, '\'')
		zend.SmartStrAppend(buf, tmp_str)
		zend.SmartStrAppendl(buf, "' => ", 5)
		zend.ZendStringFree(ckey)
		zend.ZendStringFree(tmp_str)
	}
	PhpVarExportEx(zv, level+2, buf)
	zend.SmartStrAppendc(buf, ',')
	zend.SmartStrAppendc(buf, '\n')
}
func PhpObjectElementExport(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int, buf *zend.SmartStr) {
	BufferAppendSpaces(buf, level+2)
	if key != nil {
		var class_name *byte
		var prop_name *byte
		var prop_name_len int
		var pname_esc *zend.ZendString
		zend.ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len)
		pname_esc = PhpAddcslashesStr(prop_name, prop_name_len, "'\\", 2)
		zend.SmartStrAppendc(buf, '\'')
		zend.SmartStrAppend(buf, pname_esc)
		zend.SmartStrAppendc(buf, '\'')
		zend.ZendStringReleaseEx(pname_esc, 0)
	} else {
		zend.SmartStrAppendLong(buf, zend.ZendLong(index))
	}
	zend.SmartStrAppendl(buf, " => ", 4)
	PhpVarExportEx(zv, level+2, buf)
	zend.SmartStrAppendc(buf, ',')
	zend.SmartStrAppendc(buf, '\n')
}
func PhpVarExportEx(struc *zend.Zval, level int, buf *zend.SmartStr) {
	var myht *zend.HashTable
	var tmp_str []byte
	var ztmp *zend.ZendString
	var ztmp2 *zend.ZendString
	var index zend.ZendUlong
	var key *zend.ZendString
	var val *zend.Zval
again:
	switch struc.GetType() {
	case zend.IS_FALSE:
		zend.SmartStrAppendl(buf, "false", 5)
		break
	case zend.IS_TRUE:
		zend.SmartStrAppendl(buf, "true", 4)
		break
	case zend.IS_NULL:
		zend.SmartStrAppendl(buf, "NULL", 4)
		break
	case zend.IS_LONG:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if struc.GetLval() == zend.ZEND_LONG_MIN {
			zend.SmartStrAppendLong(buf, zend.ZEND_LONG_MIN+1)
			zend.SmartStrAppends(buf, "-1")
			break
		}
		zend.SmartStrAppendLong(buf, struc.GetLval())
		break
	case zend.IS_DOUBLE:
		core.PhpGcvt(struc.GetDval(), int(core.PG(serialize_precision)), '.', 'E', tmp_str)
		zend.SmartStrAppends(buf, tmp_str)

		/* Without a decimal point, PHP treats a number literal as an int.
		 * This check even works for scientific notation, because the
		 * mantissa always contains a decimal point.
		 * We need to check for finiteness, because INF, -INF and NAN
		 * must not have a decimal point added.
		 */

		if core.ZendFinite(struc.GetDval()) && nil == strchr(tmp_str, '.') {
			zend.SmartStrAppendl(buf, ".0", 2)
		}
		break
	case zend.IS_STRING:
		ztmp = PhpAddcslashes(struc.GetStr(), "'\\", 2)
		ztmp2 = PhpStrToStr(ztmp.GetVal(), ztmp.GetLen(), "0", 1, "' . \"\\0\" . '", 12)
		zend.SmartStrAppendc(buf, '\'')
		zend.SmartStrAppend(buf, ztmp2)
		zend.SmartStrAppendc(buf, '\'')
		zend.ZendStringFree(ztmp)
		zend.ZendStringFree(ztmp2)
		break
	case zend.IS_ARRAY:
		myht = struc.GetArr()
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			if zend.GC_IS_RECURSIVE(myht) != 0 {
				zend.SmartStrAppendl(buf, "NULL", 4)
				zend.ZendError(zend.E_WARNING, "var_export does not handle circular references")
				return
			}
			myht.AddRefcount()
			zend.GC_PROTECT_RECURSION(myht)
		}
		if level > 1 {
			zend.SmartStrAppendc(buf, '\n')
			BufferAppendSpaces(buf, level-1)
		}
		zend.SmartStrAppendl(buf, "array (\n", 8)
		var __ht *zend.HashTable = myht
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()
			if _z.IsType(zend.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(zend.IS_UNDEF) {
					continue
				}
			}
			index = _p.GetH()
			key = _p.GetKey()
			val = _z
			PhpArrayElementExport(val, index, key, level, buf)
		}
		if (myht.GetGcFlags() & zend.GC_IMMUTABLE) == 0 {
			zend.GC_UNPROTECT_RECURSION(myht)
			myht.DelRefcount()
		}
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		zend.SmartStrAppendc(buf, ')')
		break
	case zend.IS_OBJECT:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_VAR_EXPORT)
		if myht != nil {
			if zend.GC_IS_RECURSIVE(myht) != 0 {
				zend.SmartStrAppendl(buf, "NULL", 4)
				zend.ZendError(zend.E_WARNING, "var_export does not handle circular references")
				zend.ZendReleaseProperties(myht)
				return
			} else {
				zend.GC_TRY_PROTECT_RECURSION(myht)
			}
		}
		if level > 1 {
			zend.SmartStrAppendc(buf, '\n')
			BufferAppendSpaces(buf, level-1)
		}

		/* stdClass has no __set_state method, but can be casted to */

		if zend.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			zend.SmartStrAppendl(buf, "(object) array(\n", 16)
		} else {
			zend.SmartStrAppend(buf, zend.Z_OBJCE_P(struc).GetName())
			zend.SmartStrAppendl(buf, "::__set_state(array(\n", 21)
		}
		if myht != nil {
			var __ht *zend.HashTable = myht
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()
				if _z.IsType(zend.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(zend.IS_UNDEF) {
						continue
					}
				}
				index = _p.GetH()
				key = _p.GetKey()
				val = _z
				PhpObjectElementExport(val, index, key, level, buf)
			}
			zend.GC_TRY_UNPROTECT_RECURSION(myht)
			zend.ZendReleaseProperties(myht)
		}
		if level > 1 {
			BufferAppendSpaces(buf, level-1)
		}
		if zend.Z_OBJCE_P(struc) == zend.ZendStandardClassDef {
			zend.SmartStrAppendc(buf, ')')
		} else {
			zend.SmartStrAppendl(buf, "))", 2)
		}
		break
	case zend.IS_REFERENCE:
		struc = zend.Z_REFVAL_P(struc)
		goto again
		break
	default:
		zend.SmartStrAppendl(buf, "NULL", 4)
		break
	}
}
func PhpVarExport(struc *zend.Zval, level int) {
	var buf zend.SmartStr = zend.SmartStr{0}
	PhpVarExportEx(struc, level, &buf)
	zend.SmartStr0(&buf)
	core.PHPWRITE(buf.GetS().GetVal(), buf.GetS().GetLen())
	zend.SmartStrFree(&buf)
}
func ZifVarExport(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var return_output zend.ZendBool = 0
	var buf zend.SmartStr = zend.SmartStr{0}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &return_output, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	PhpVarExportEx(var_, 1, &buf)
	zend.SmartStr0(&buf)
	if return_output != 0 {
		zend.RETVAL_NEW_STR(buf.GetS())
		return
	} else {
		core.PHPWRITE(buf.GetS().GetVal(), buf.GetS().GetLen())
		zend.SmartStrFree(&buf)
	}
}
func PhpAddVarHash(data PhpSerializeDataT, var_ *zend.Zval) zend.ZendLong {
	var zv *zend.Zval
	var key zend.ZendUlong
	var is_ref zend.ZendBool = zend.Z_ISREF_P(var_)
	data.SetN(data.GetN() + 1)
	if is_ref == 0 && var_.GetType() != zend.IS_OBJECT {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */

	if is_ref != 0 && zend.Z_REFVAL_P(var_).IsType(zend.IS_OBJECT) {
		var_ = zend.Z_REFVAL_P(var_)
	}

	/* Index for the variable is stored using the numeric value of the pointer to
	 * the zend_refcounted struct */

	key = zend.ZendUlong(zend.ZendUintptrT(var_.GetCounted()))
	zv = zend.ZendHashIndexFind(data.GetHt(), key)
	if zv != nil {

		/* References are only counted once, undo the data->n increment above */

		if is_ref != 0 && zv.GetLval() != -1 {
			data.SetN(data.GetN() - 1)
		}
		return zv.GetLval()
	} else {
		var zv_n zend.Zval
		zend.ZVAL_LONG(&zv_n, data.GetN())
		data.GetHt().IndexAddNewH(key, &zv_n)

		/* Additionally to the index, we also store the variable, to ensure that it is
		 * not destroyed during serialization and its pointer reused. The variable is
		 * stored at the numeric value of the pointer + 1, which cannot be the location
		 * of another zend_refcounted structure. */

		data.GetHt().IndexAddNewH(key+1, var_)
		zend.Z_ADDREF_P(var_)
		return 0
	}
}
func PhpVarSerializeLong(buf *zend.SmartStr, val zend.ZendLong) {
	zend.SmartStrAppendl(buf, "i:", 2)
	zend.SmartStrAppendLong(buf, val)
	zend.SmartStrAppendc(buf, ';')
}
func PhpVarSerializeString(buf *zend.SmartStr, str *byte, len_ int) {
	zend.SmartStrAppendl(buf, "s:", 2)
	zend.SmartStrAppendUnsigned(buf, len_)
	zend.SmartStrAppendl(buf, ":\"", 2)
	zend.SmartStrAppendl(buf, str, len_)
	zend.SmartStrAppendl(buf, "\";", 2)
}
func PhpVarSerializeClassName(buf *zend.SmartStr, struc *zend.Zval) zend.ZendBool {
	var class_name *zend.ZendString
	var incomplete_class zend.ZendBool = 0
	PHP_SET_CLASS_ATTRIBUTES(struc)
	zend.SmartStrAppendl(buf, "O:", 2)
	zend.SmartStrAppendUnsigned(buf, class_name.GetLen())
	zend.SmartStrAppendl(buf, ":\"", 2)
	zend.SmartStrAppend(buf, class_name)
	zend.SmartStrAppendl(buf, "\":", 2)
	PHP_CLEANUP_CLASS_ATTRIBUTES()
	return incomplete_class
}
func PhpVarSerializeCallSleep(retval *zend.Zval, struc *zend.Zval) int {
	var fname zend.Zval
	var res int
	zend.ZVAL_STRINGL(&fname, "__sleep", b.SizeOf("\"__sleep\"")-1)
	BG(serialize_lock)++
	res = zend.CallUserFunction(nil, struc, &fname, retval, 0, 0)
	BG(serialize_lock)--
	zend.ZvalPtrDtorStr(&fname)
	if res == zend.FAILURE || zend.Z_ISUNDEF_P(retval) {
		zend.ZvalPtrDtor(retval)
		return zend.FAILURE
	}
	if !(zend.HASH_OF(retval)) {
		zend.ZvalPtrDtor(retval)
		core.PhpErrorDocref(nil, zend.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize")
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func PhpVarSerializeCallMagicSerialize(retval *zend.Zval, obj *zend.Zval) int {
	var fname zend.Zval
	var res int
	zend.ZVAL_STRINGL(&fname, "__serialize", b.SizeOf("\"__serialize\"")-1)
	BG(serialize_lock)++
	res = zend.CallUserFunction(zend.__CG().GetFunctionTable(), obj, &fname, retval, 0, 0)
	BG(serialize_lock)--
	zend.ZvalPtrDtorStr(&fname)
	if res == zend.FAILURE || zend.Z_ISUNDEF_P(retval) {
		zend.ZvalPtrDtor(retval)
		return zend.FAILURE
	}
	if retval.GetType() != zend.IS_ARRAY {
		zend.ZvalPtrDtor(retval)
		zend.ZendTypeError("%s::__serialize() must return an array", zend.Z_OBJCE_P(obj).GetName().GetVal())
		return zend.FAILURE
	}
	return zend.SUCCESS
}
func PhpVarSerializeTryAddSleepProp(ht *zend.HashTable, props *zend.HashTable, name *zend.ZendString, error_name *zend.ZendString, struc *zend.Zval) int {
	var val *zend.Zval = props.FindByZendString(name)
	if val == nil {
		return zend.FAILURE
	}
	if val.IsType(zend.IS_INDIRECT) {
		val = val.GetZv()
		if val.IsType(zend.IS_UNDEF) {
			var info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(struc.GetObj(), val)
			if info != nil {
				return zend.SUCCESS
			}
			return zend.FAILURE
		}
	}
	if ht.KeyAdd(name.GetStr(), val) == nil {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "\"%s\" is returned from __sleep multiple times", error_name.GetVal())
		return zend.SUCCESS
	}
	zend.Z_TRY_ADDREF_P(val)
	return zend.SUCCESS
}
func PhpVarSerializeGetSleepProps(ht *zend.HashTable, struc *zend.Zval, sleep_retval *zend.HashTable) int {
	var ce *zend.ZendClassEntry = zend.Z_OBJCE_P(struc)
	var props *zend.HashTable = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)
	var name_val *zend.Zval
	var retval int = zend.SUCCESS
	zend.ZendHashInit(ht, sleep_retval.GetNNumOfElements(), nil, zend.ZVAL_PTR_DTOR, 0)

	/* TODO: Rewrite this by fetching the property info instead of trying out different
	 * name manglings? */

	var __ht *zend.HashTable = sleep_retval
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()
		if _z.IsType(zend.IS_INDIRECT) {
			_z = _z.GetZv()
			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
		}
		name_val = _z
		var name *zend.ZendString
		var tmp_name *zend.ZendString
		var priv_name *zend.ZendString
		var prot_name *zend.ZendString
		zend.ZVAL_DEREF(name_val)
		if name_val.GetType() != zend.IS_STRING {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "__sleep should return an array only containing the names of instance-variables to serialize.")
		}
		name = zend.ZvalGetTmpString(name_val, &tmp_name)
		if PhpVarSerializeTryAddSleepProp(ht, props, name, name, struc) == zend.SUCCESS {
			zend.ZendTmpStringRelease(tmp_name)
			continue
		}
		if zend.__EG().GetException() != nil {
			zend.ZendTmpStringRelease(tmp_name)
			retval = zend.FAILURE
			break
		}
		priv_name = zend.ZendManglePropertyName(ce.GetName().GetVal(), ce.GetName().GetLen(), name.GetVal(), name.GetLen(), ce.GetType()&zend.ZEND_INTERNAL_CLASS)
		if PhpVarSerializeTryAddSleepProp(ht, props, priv_name, name, struc) == zend.SUCCESS {
			zend.ZendTmpStringRelease(tmp_name)
			zend.ZendStringRelease(priv_name)
			continue
		}
		zend.ZendStringRelease(priv_name)
		if zend.__EG().GetException() != nil {
			zend.ZendTmpStringRelease(tmp_name)
			retval = zend.FAILURE
			break
		}
		prot_name = zend.ZendManglePropertyName("*", 1, name.GetVal(), name.GetLen(), ce.GetType()&zend.ZEND_INTERNAL_CLASS)
		if PhpVarSerializeTryAddSleepProp(ht, props, prot_name, name, struc) == zend.SUCCESS {
			zend.ZendTmpStringRelease(tmp_name)
			zend.ZendStringRelease(prot_name)
			continue
		}
		zend.ZendStringRelease(prot_name)
		if zend.__EG().GetException() != nil {
			zend.ZendTmpStringRelease(tmp_name)
			retval = zend.FAILURE
			break
		}
		core.PhpErrorDocref(nil, zend.E_NOTICE, "\"%s\" returned as member variable from __sleep() but does not exist", name.GetVal())
		ht.KeyAdd(name.GetStr(), zend.__EG().GetUninitializedZval())
		zend.ZendTmpStringRelease(tmp_name)
	}
	zend.ZendReleaseProperties(props)
	return retval
}
func PhpVarSerializeNestedData(buf *zend.SmartStr, struc *zend.Zval, ht *zend.HashTable, count uint32, incomplete_class zend.ZendBool, var_hash PhpSerializeDataT) {
	zend.SmartStrAppendUnsigned(buf, count)
	zend.SmartStrAppendl(buf, ":{", 2)
	if count > 0 {
		var key *zend.ZendString
		var data *zend.Zval
		var index zend.ZendUlong
		var __ht *zend.HashTable = ht
		for _, _p := range __ht.foreachData() {
			var _z *zend.Zval = _p.GetVal()
			if _z.IsType(zend.IS_INDIRECT) {
				_z = _z.GetZv()
				if _z.IsType(zend.IS_UNDEF) {
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
			if zend.Z_ISREF_P(data) && zend.Z_REFCOUNT_P(data) == 1 {
				data = zend.Z_REFVAL_P(data)
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */

			if data.IsType(zend.IS_ARRAY) {
				if zend.Z_IS_RECURSIVE_P(data) != 0 || struc.IsType(zend.IS_ARRAY) && data.GetArr() == struc.GetArr() {
					PhpAddVarHash(var_hash, struc)
					zend.SmartStrAppendl(buf, "N;", 2)
				} else {
					if zend.Z_REFCOUNTED_P(data) {
						zend.Z_PROTECT_RECURSION_P(data)
					}
					PhpVarSerializeIntern(buf, data, var_hash)
					if zend.Z_REFCOUNTED_P(data) {
						zend.Z_UNPROTECT_RECURSION_P(data)
					}
				}
			} else {
				PhpVarSerializeIntern(buf, data, var_hash)
			}

			/* we should still add element even if it's not OK,
			 * since we already wrote the length of the array before */

		}
	}
	zend.SmartStrAppendc(buf, '}')
}
func PhpVarSerializeClass(buf *zend.SmartStr, struc *zend.Zval, retval_ptr *zend.Zval, var_hash PhpSerializeDataT) {
	var props zend.HashTable
	if PhpVarSerializeGetSleepProps(&props, struc, zend.HASH_OF(retval_ptr)) == zend.SUCCESS {
		PhpVarSerializeClassName(buf, struc)
		PhpVarSerializeNestedData(buf, struc, &props, props.GetNNumOfElements(), 0, var_hash)
	}
	zend.ZendHashDestroy(&props)
}
func PhpVarSerializeIntern(buf *zend.SmartStr, struc *zend.Zval, var_hash PhpSerializeDataT) {
	var var_already zend.ZendLong
	var myht *zend.HashTable
	if zend.__EG().GetException() != nil {
		return
	}
	if var_hash != nil && b.Assign(&var_already, PhpAddVarHash(var_hash, struc)) {
		if var_already == -1 {

			/* Reference to an object that failed to serialize, replace with null. */

			zend.SmartStrAppendl(buf, "N;", 2)
			return
		} else if zend.Z_ISREF_P(struc) {
			zend.SmartStrAppendl(buf, "R:", 2)
			zend.SmartStrAppendLong(buf, var_already)
			zend.SmartStrAppendc(buf, ';')
			return
		} else if struc.IsType(zend.IS_OBJECT) {
			zend.SmartStrAppendl(buf, "r:", 2)
			zend.SmartStrAppendLong(buf, var_already)
			zend.SmartStrAppendc(buf, ';')
			return
		}
	}
again:
	switch struc.GetType() {
	case zend.IS_FALSE:
		zend.SmartStrAppendl(buf, "b:0;", 4)
		return
	case zend.IS_TRUE:
		zend.SmartStrAppendl(buf, "b:1;", 4)
		return
	case zend.IS_NULL:
		zend.SmartStrAppendl(buf, "N;", 2)
		return
	case zend.IS_LONG:
		PhpVarSerializeLong(buf, struc.GetLval())
		return
	case zend.IS_DOUBLE:
		var tmp_str []byte
		zend.SmartStrAppendl(buf, "d:", 2)
		core.PhpGcvt(struc.GetDval(), int(core.PG(serialize_precision)), '.', 'E', tmp_str)
		zend.SmartStrAppends(buf, tmp_str)
		zend.SmartStrAppendc(buf, ';')
		return
	case zend.IS_STRING:
		PhpVarSerializeString(buf, zend.Z_STRVAL_P(struc), zend.Z_STRLEN_P(struc))
		return
	case zend.IS_OBJECT:
		var ce *zend.ZendClassEntry = zend.Z_OBJCE_P(struc)
		var incomplete_class zend.ZendBool
		var count uint32
		if zend.ZendHashStrExists(ce.GetFunctionTable(), "__serialize", b.SizeOf("\"__serialize\"")-1) != 0 {
			var retval zend.Zval
			var obj zend.Zval
			var key *zend.ZendString
			var data *zend.Zval
			var index zend.ZendUlong
			zend.Z_ADDREF_P(struc)
			zend.ZVAL_OBJ(&obj, struc.GetObj())
			if PhpVarSerializeCallMagicSerialize(&retval, &obj) == zend.FAILURE {
				if zend.__EG().GetException() == nil {
					zend.SmartStrAppendl(buf, "N;", 2)
				}
				zend.ZvalPtrDtor(&obj)
				return
			}
			PhpVarSerializeClassName(buf, &obj)
			zend.SmartStrAppendUnsigned(buf, retval.GetArr().Count())
			zend.SmartStrAppendl(buf, ":{", 2)
			var __ht *zend.HashTable = retval.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()
				if _z.IsType(zend.IS_INDIRECT) {
					_z = _z.GetZv()
					if _z.IsType(zend.IS_UNDEF) {
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
				if zend.Z_ISREF_P(data) && zend.Z_REFCOUNT_P(data) == 1 {
					data = zend.Z_REFVAL_P(data)
				}
				PhpVarSerializeIntern(buf, data, var_hash)
			}
			zend.SmartStrAppendc(buf, '}')
			zend.ZvalPtrDtor(&obj)
			zend.ZvalPtrDtor(&retval)
			return
		}
		if ce.GetSerialize() != nil {

			/* has custom handler */

			var serialized_data *uint8 = nil
			var serialized_length int
			if ce.GetSerialize()(struc, &serialized_data, &serialized_length, (*zend.ZendSerializeData)(var_hash)) == zend.SUCCESS {
				zend.SmartStrAppendl(buf, "C:", 2)
				zend.SmartStrAppendUnsigned(buf, zend.Z_OBJCE_P(struc).GetName().GetLen())
				zend.SmartStrAppendl(buf, ":\"", 2)
				zend.SmartStrAppend(buf, zend.Z_OBJCE_P(struc).GetName())
				zend.SmartStrAppendl(buf, "\":", 2)
				zend.SmartStrAppendUnsigned(buf, serialized_length)
				zend.SmartStrAppendl(buf, ":{", 2)
				zend.SmartStrAppendl(buf, (*byte)(serialized_data), serialized_length)
				zend.SmartStrAppendc(buf, '}')
			} else {

				/* Mark this value in the var_hash, to avoid creating references to it. */

				var var_idx *zend.Zval = zend.ZendHashIndexFind(var_hash.GetHt(), zend.ZendUlong(zend.ZendUintptrT(struc.GetCounted())))
				zend.ZVAL_LONG(var_idx, -1)
				zend.SmartStrAppendl(buf, "N;", 2)
			}
			if serialized_data != nil {
				zend.Efree(serialized_data)
			}
			return
		}
		if ce != PHP_IC_ENTRY && zend.ZendHashStrExists(ce.GetFunctionTable(), "__sleep", b.SizeOf("\"__sleep\"")-1) != 0 {
			var retval zend.Zval
			var tmp zend.Zval
			zend.Z_ADDREF_P(struc)
			zend.ZVAL_OBJ(&tmp, struc.GetObj())
			if PhpVarSerializeCallSleep(&retval, &tmp) == zend.FAILURE {
				if zend.__EG().GetException() == nil {

					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */

					zend.SmartStrAppendl(buf, "N;", 2)

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
	case zend.IS_ARRAY:
		zend.SmartStrAppendl(buf, "a:", 2)
		myht = struc.GetArr()
		PhpVarSerializeNestedData(buf, struc, myht, myht.Count(), 0, var_hash)
		return
	case zend.IS_REFERENCE:
		struc = zend.Z_REFVAL_P(struc)
		goto again
	default:
		zend.SmartStrAppendl(buf, "i:0;", 4)
		return
	}
}
func PhpVarSerialize(buf *zend.SmartStr, struc *zend.Zval, data *PhpSerializeDataT) {
	PhpVarSerializeIntern(buf, struc, *data)
	zend.SmartStr0(buf)
}
func PhpVarSerializeInit() PhpSerializeDataT {
	var d *PhpSerializeData

	/* fprintf(stderr, "SERIALIZE_INIT      == lock: %u, level: %u\n", BG(serialize_lock), BG(serialize).level); */

	if BG(serialize_lock) || !(BG(serialize).level) {
		d = zend.Emalloc(b.SizeOf("struct php_serialize_data"))
		zend.ZendHashInit(d.GetHt(), 16, nil, zend.ZVAL_PTR_DTOR, 0)
		d.SetN(0)
		if !(BG(serialize_lock)) {
			BG(serialize).data = d
			BG(serialize).level = 1
		}
	} else {
		d = BG(serialize).data
		BG(serialize).level++
	}
	return d
}
func PhpVarSerializeDestroy(d PhpSerializeDataT) {
	/* fprintf(stderr, "SERIALIZE_DESTROY   == lock: %u, level: %u\n", BG(serialize_lock), BG(serialize).level); */

	if BG(serialize_lock) || BG(serialize).level == 1 {
		zend.ZendHashDestroy(d.GetHt())
		zend.Efree(d)
	}
	if !(BG(serialize_lock)) && !(b.PreDec(&(BG(serialize).level))) {
		BG(serialize).data = nil
	}
}
func ZifSerialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var struc *zend.Zval
	var var_hash PhpSerializeDataT
	var buf zend.SmartStr = zend.SmartStr{0}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &struc, 0)
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return
		}
		break
	}
	PHP_VAR_SERIALIZE_INIT(var_hash)
	PhpVarSerialize(&buf, struc, &var_hash)
	PHP_VAR_SERIALIZE_DESTROY(var_hash)
	if zend.__EG().GetException() != nil {
		zend.SmartStrFree(&buf)
		zend.RETVAL_FALSE
		return
	}
	if buf.GetS() != nil {
		zend.RETVAL_NEW_STR(buf.GetS())
		return
	} else {
		zend.RETVAL_NULL()
		return
	}
}
func ZifUnserialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var buf *byte = nil
	var buf_len int
	var p *uint8
	var var_hash PhpUnserializeDataT
	var options *zend.Zval = nil
	var retval *zend.Zval
	var class_hash *zend.HashTable = nil
	var prev_class_hash *zend.HashTable
	var prev_max_depth zend.ZendLong
	var prev_cur_depth zend.ZendLong
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &buf, &buf_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	if buf_len == 0 {
		zend.RETVAL_FALSE
		return
	}
	p = (*uint8)(buf)
	PHP_VAR_UNSERIALIZE_INIT(var_hash)
	prev_class_hash = PhpVarUnserializeGetAllowedClasses(var_hash)
	prev_max_depth = PhpVarUnserializeGetMaxDepth(var_hash)
	prev_cur_depth = PhpVarUnserializeGetCurDepth(var_hash)
	if options != nil {
		var classes *zend.Zval
		var max_depth *zend.Zval
		classes = zend.ZendHashStrFindDeref(options.GetArr(), "allowed_classes", b.SizeOf("\"allowed_classes\"")-1)
		if classes != nil && classes.GetType() != zend.IS_ARRAY && classes.GetType() != zend.IS_TRUE && classes.GetType() != zend.IS_FALSE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "allowed_classes option should be array or boolean")
			zend.RETVAL_FALSE
			goto cleanup
		}
		if classes != nil && (classes.IsType(zend.IS_ARRAY) || zend.ZendIsTrue(classes) == 0) {
			zend.ALLOC_HASHTABLE(class_hash)
			zend.ZendHashInit(class_hash, b.CondF1(classes.IsType(zend.IS_ARRAY), func() __auto__ { return zend.Z_ARRVAL_P(classes).GetNNumOfElements() }, 0), nil, nil, 0)
		}
		if class_hash != nil && classes.IsType(zend.IS_ARRAY) {
			var entry *zend.Zval
			var lcname *zend.ZendString
			var __ht *zend.HashTable = classes.GetArr()
			for _, _p := range __ht.foreachData() {
				var _z *zend.Zval = _p.GetVal()

				entry = _z
				zend.ConvertToStringEx(entry)
				lcname = zend.ZendStringTolower(entry.GetStr())
				zend.ZendHashAddEmptyElement(class_hash, lcname)
				zend.ZendStringReleaseEx(lcname, 0)
			}

			/* Exception during string conversion. */

			if zend.__EG().GetException() != nil {
				goto cleanup
			}

			/* Exception during string conversion. */

		}
		PhpVarUnserializeSetAllowedClasses(var_hash, class_hash)
		max_depth = zend.ZendHashStrFindDeref(options.GetArr(), "max_depth", b.SizeOf("\"max_depth\"")-1)
		if max_depth != nil {
			if max_depth.GetType() != zend.IS_LONG {
				core.PhpErrorDocref(nil, zend.E_WARNING, "max_depth should be int")
				zend.RETVAL_FALSE
				goto cleanup
			}
			if max_depth.GetLval() < 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "max_depth cannot be negative")
				zend.RETVAL_FALSE
				goto cleanup
			}
			PhpVarUnserializeSetMaxDepth(var_hash, max_depth.GetLval())

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

			PhpVarUnserializeSetCurDepth(var_hash, 0)

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

		}
	}
	if BG(unserialize).level > 1 {
		retval = VarTmpVar(&var_hash)
	} else {
		retval = return_value
	}
	if PhpVarUnserialize(retval, &p, p+buf_len, &var_hash) == 0 {
		if zend.__EG().GetException() == nil {
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Error at offset "+zend.ZEND_LONG_FMT+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
		}
		if BG(unserialize).level <= 1 {
			zend.ZvalPtrDtor(return_value)
		}
		zend.RETVAL_FALSE
	} else if BG(unserialize).level > 1 {
		zend.ZVAL_COPY(return_value, retval)
	} else if zend.Z_REFCOUNTED_P(return_value) {
		var ref *zend.ZendRefcounted = return_value.GetCounted()
		zend.GcCheckPossibleRoot(ref)
	}
cleanup:
	if class_hash != nil {
		zend.ZendHashDestroy(class_hash)
		zend.FREE_HASHTABLE(class_hash)
	}

	/* Reset to previous options in case this is a nested call */

	PhpVarUnserializeSetAllowedClasses(var_hash, prev_class_hash)
	PhpVarUnserializeSetMaxDepth(var_hash, prev_max_depth)
	PhpVarUnserializeSetCurDepth(var_hash, prev_cur_depth)
	PHP_VAR_UNSERIALIZE_DESTROY(var_hash)

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */

	if zend.Z_ISREF_P(return_value) {
		zend.ZendUnwrapReference(return_value)
	}

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */
}
func ZifMemoryGetUsage(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var real_usage zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &real_usage, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	zend.RETVAL_LONG(zend.ZendMemoryUsage(real_usage))
	return
}
func ZifMemoryGetPeakUsage(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var real_usage zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = zend.EX_NUM_ARGS()
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = zend.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = zend.ZPP_ERROR_FAILURE
				break
			}
			_real_arg = zend.ZEND_CALL_ARG(execute_data, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgBool(_arg, &real_usage, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			break
		}
		if _error_code != zend.ZPP_ERROR_OK {
			if (_flags & zend.ZEND_PARSE_PARAMS_QUIET) == 0 {
				if _error_code == zend.ZPP_ERROR_WRONG_CALLBACK {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_CLASS {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == zend.ZPP_ERROR_WRONG_ARG {
					if (_flags & zend.ZEND_PARSE_PARAMS_THROW) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			zend.RETVAL_FALSE
			return
		}
		break
	}
	zend.RETVAL_LONG(zend.ZendMemoryPeakUsage(real_usage))
	return
}
func ZmStartupVar(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
