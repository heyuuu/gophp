// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/var.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Jani Lehtimäki <jkl@njet.net>                               |
   |          Thies C. Arntzen <thies@thieso.net>                         |
   |          Sascha Schumann <sascha@schumann.cx>                        |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include < stdlib . h >

// # include < errno . h >

// # include "php.h"

// # include "php_string.h"

// # include "php_var.h"

// # include "zend_smart_str.h"

// # include "basic_functions.h"

// # include "php_incomplete_class.h"

/* }}} */

// @type PhpSerializeData struct

// #define COMMON       ( is_ref ? "&" : "" )

func PhpArrayElementDump(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	if key == nil {
		core.PhpPrintf("%*c["+"%"+"lld"+"]=>\n", level+1, ' ', index)
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PhpOutputWrite(key.val, key.len_)
		core.PhpPrintf("\"]=>\n")
	}
	PhpVarDump(zv, level+2)
}

/* }}} */

func PhpObjectPropertyDump(prop_info *zend.ZendPropertyInfo, zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	var prop_name *byte
	var class_name *byte
	if key == nil {
		core.PhpPrintf("%*c["+"%"+"lld"+"]=>\n", level+1, ' ', index)
	} else {
		var unmangle int = zend.ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, nil)
		core.PhpPrintf("%*c[", level+1, ' ')
		if class_name != nil && unmangle == zend.SUCCESS {
			if class_name[0] == '*' {
				core.PhpPrintf("\"%s\":protected", prop_name)
			} else {
				core.PhpPrintf("\"%s\":\"%s\":private", prop_name, class_name)
			}
		} else {
			core.PhpPrintf("\"")
			core.PhpOutputWrite(key.val, key.len_)
			core.PhpPrintf("\"")
		}
		zend.ZendWrite("]=>\n", strlen("]=>\n"))
	}
	if zv.u1.v.type_ == 0 {
		assert(prop_info.type_ != 0)
		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', g.Cond((prop_info.type_&0x1) != 0, "?", ""), g.CondF(prop_info.type_ > 0x3ff, func() __auto__ {
			return g.CondF((prop_info.type_&0x2) != 0, func() *zend.ZendString { return (*zend.ZendClassEntry)(prop_info.type_ & ^0x3).name }, func() *zend.ZendString { return (*zend.ZendString)(prop_info.type_ & ^0x3) }).val
		}, func() *byte { return zend.ZendGetTypeByConst(prop_info.type_ >> 2) }))
	} else {
		PhpVarDump(zv, level+2)
	}
}

/* }}} */

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
	switch struc.u1.v.type_ {
	case 2:
		core.PhpPrintf("%sbool(false)\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 3:
		core.PhpPrintf("%sbool(true)\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 1:
		core.PhpPrintf("%sNULL\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 4:
		core.PhpPrintf("%sint("+"%"+"lld"+")\n", g.Cond(is_ref != 0, "&", ""), struc.value.lval)
		break
	case 5:
		core.PhpPrintf("%sfloat(%.*G)\n", g.Cond(is_ref != 0, "&", ""), int(zend.EG.precision), struc.value.dval)
		break
	case 6:
		core.PhpPrintf("%sstring(%zd) \"", g.Cond(is_ref != 0, "&", ""), struc.value.str.len_)
		core.PhpOutputWrite(struc.value.str.val, struc.value.str.len_)
		var __str *byte = "\"\n"
		core.PhpOutputWrite(__str, strlen(__str))
		break
	case 7:
		myht = struc.value.arr
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			if level > 1 {
				if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 5) != 0 {
					var __str *byte = "*RECURSION*\n"
					core.PhpOutputWrite(__str, strlen(__str))
					return
				}
				myht.gc.u.type_info |= 1 << 5 << 0
			}
			zend.ZendGcAddref(&myht.gc)
		}
		count = zend.ZendArrayCount(myht)
		core.PhpPrintf("%sarray(%d) {\n", g.Cond(is_ref != 0, "&", ""), count)
		for {
			var __ht *zend.HashTable = myht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				num = _p.h
				key = _p.key
				val = _z
				PhpArrayElementDump(val, num, key, level)
			}
			break
		}
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			if level > 1 {
				myht.gc.u.type_info &= ^(1 << 5 << 0)
			}
			zend.ZendGcDelref(&myht.gc)
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		var __str *byte = "}\n"
		core.PhpOutputWrite(__str, strlen(__str))
		break
	case 8:
		if (zend.ZvalGcFlags(struc.value.counted.gc.u.type_info) & 1 << 5) != 0 {
			var __str *byte = "*RECURSION*\n"
			core.PhpOutputWrite(__str, strlen(__str))
			return
		}
		struc.value.counted.gc.u.type_info |= 1 << 5 << 0
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		class_name = struc.value.obj.handlers.get_class_name(struc.value.obj)
		core.PhpPrintf("%sobject(%s)#%d (%d) {\n", g.Cond(is_ref != 0, "&", ""), class_name.val, struc.value.obj.handle, g.CondF1(myht != nil, func() uint32 { return zend.ZendArrayCount(myht) }, 0))
		zend.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			var num zend.ZendUlong
			var key *zend.ZendString
			var val *zend.Zval
			for {
				var __ht *zend.HashTable = myht
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					num = _p.h
					key = _p.key
					val = _z
					var prop_info *zend.ZendPropertyInfo = nil
					if val.u1.v.type_ == 13 {
						val = val.value.zv
						if key != nil {
							prop_info = zend.ZendGetTypedPropertyInfoForSlot(struc.value.obj, val)
						}
					}
					if val.u1.v.type_ != 0 || prop_info != nil {
						PhpObjectPropertyDump(prop_info, val, num, key, level)
					}
				}
				break
			}
			if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
				zend.ZendArrayDestroy(myht)
			}
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		var __str *byte = "}\n"
		core.PhpOutputWrite(__str, strlen(__str))
		struc.value.counted.gc.u.type_info &= ^(1 << 5 << 0)
		break
	case 9:
		var type_name *byte = zend.ZendRsrcListGetRsrcType(struc.value.res)
		core.PhpPrintf("%sresource(%d) of type (%s)\n", g.Cond(is_ref != 0, "&", ""), struc.value.res.handle, g.Cond(type_name != nil, type_name, "Unknown"))
		break
	case 10:

		//??? hide references with refcount==1 (for compatibility)

		if zend.ZvalRefcountP(struc) > 1 {
			is_ref = 1
		}
		struc = &(*struc).value.ref.val
		goto again
		break
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", g.Cond(is_ref != 0, "&", ""))
		break
	}
}

/* }}} */

func ZifVarDump(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var argc int
	var i int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

/* }}} */

func ZvalArrayElementDump(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	if key == nil {
		core.PhpPrintf("%*c["+"%"+"lld"+"]=>\n", level+1, ' ', index)
	} else {
		core.PhpPrintf("%*c[\"", level+1, ' ')
		core.PhpOutputWrite(key.val, key.len_)
		core.PhpPrintf("\"]=>\n")
	}
	PhpDebugZvalDump(zv, level+2)
}

/* }}} */

func ZvalObjectPropertyDump(prop_info *zend.ZendPropertyInfo, zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int) {
	var prop_name *byte
	var class_name *byte
	if key == nil {
		core.PhpPrintf("%*c["+"%"+"lld"+"]=>\n", level+1, ' ', index)
	} else {
		zend.ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, nil)
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
		zend.ZendWrite("]=>\n", strlen("]=>\n"))
	}
	if prop_info != nil && zv.u1.v.type_ == 0 {
		assert(prop_info.type_ != 0)
		core.PhpPrintf("%*cuninitialized(%s%s)\n", level+1, ' ', g.Cond((prop_info.type_&0x1) != 0, "?", ""), g.CondF(prop_info.type_ > 0x3ff, func() __auto__ {
			return g.CondF((prop_info.type_&0x2) != 0, func() *zend.ZendString { return (*zend.ZendClassEntry)(prop_info.type_ & ^0x3).name }, func() *zend.ZendString { return (*zend.ZendString)(prop_info.type_ & ^0x3) }).val
		}, func() *byte { return zend.ZendGetTypeByConst(prop_info.type_ >> 2) }))
	} else {
		PhpDebugZvalDump(zv, level+2)
	}
}

/* }}} */

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
	switch struc.u1.v.type_ {
	case 2:
		core.PhpPrintf("%sbool(false)\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 3:
		core.PhpPrintf("%sbool(true)\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 1:
		core.PhpPrintf("%sNULL\n", g.Cond(is_ref != 0, "&", ""))
		break
	case 4:
		core.PhpPrintf("%sint("+"%"+"lld"+")\n", g.Cond(is_ref != 0, "&", ""), struc.value.lval)
		break
	case 5:
		core.PhpPrintf("%sfloat(%.*G)\n", g.Cond(is_ref != 0, "&", ""), int(zend.EG.precision), struc.value.dval)
		break
	case 6:
		core.PhpPrintf("%sstring(%zd) \"", g.Cond(is_ref != 0, "&", ""), struc.value.str.len_)
		core.PhpOutputWrite(struc.value.str.val, struc.value.str.len_)
		core.PhpPrintf("\" refcount(%u)\n", g.CondF1(struc.u1.v.type_flags != 0, func() uint32 { return zend.ZvalRefcountP(struc) }, 1))
		break
	case 7:
		myht = struc.value.arr
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			if level > 1 {
				if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 5) != 0 {
					var __str *byte = "*RECURSION*\n"
					core.PhpOutputWrite(__str, strlen(__str))
					return
				}
				myht.gc.u.type_info |= 1 << 5 << 0
			}
			zend.ZendGcAddref(&myht.gc)
		}
		count = zend.ZendArrayCount(myht)
		core.PhpPrintf("%sarray(%d) refcount(%u){\n", g.Cond(is_ref != 0, "&", ""), count, g.CondF1(struc.u1.v.type_flags != 0, func() int { return zend.ZvalRefcountP(struc) - 1 }, 1))
		for {
			var __ht *zend.HashTable = myht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				index = _p.h
				key = _p.key
				val = _z
				ZvalArrayElementDump(val, index, key, level)
			}
			break
		}
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			if level > 1 {
				myht.gc.u.type_info &= ^(1 << 5 << 0)
			}
			zend.ZendGcDelref(&myht.gc)
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		var __str *byte = "}\n"
		core.PhpOutputWrite(__str, strlen(__str))
		break
	case 8:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_DEBUG)
		if myht != nil {
			if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 5) != 0 {
				var __str *byte = "*RECURSION*\n"
				core.PhpOutputWrite(__str, strlen(__str))
				if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
					zend.ZendArrayDestroy(myht)
				}
				return
			}
			myht.gc.u.type_info |= 1 << 5 << 0
		}
		class_name = struc.value.obj.handlers.get_class_name(struc.value.obj)
		core.PhpPrintf("%sobject(%s)#%d (%d) refcount(%u){\n", g.Cond(is_ref != 0, "&", ""), class_name.val, struc.value.obj.handle, g.CondF1(myht != nil, func() uint32 { return zend.ZendArrayCount(myht) }, 0), zend.ZvalRefcountP(struc))
		zend.ZendStringReleaseEx(class_name, 0)
		if myht != nil {
			for {
				var __ht *zend.HashTable = myht
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					index = _p.h
					key = _p.key
					val = _z
					var prop_info *zend.ZendPropertyInfo = nil
					if val.u1.v.type_ == 13 {
						val = val.value.zv
						if key != nil {
							prop_info = zend.ZendGetTypedPropertyInfoForSlot(struc.value.obj, val)
						}
					}
					if val.u1.v.type_ != 0 || prop_info != nil {
						ZvalObjectPropertyDump(prop_info, val, index, key, level)
					}
				}
				break
			}
			myht.gc.u.type_info &= ^(1 << 5 << 0)
			if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
				zend.ZendArrayDestroy(myht)
			}
		}
		if level > 1 {
			core.PhpPrintf("%*c", level-1, ' ')
		}
		var __str *byte = "}\n"
		core.PhpOutputWrite(__str, strlen(__str))
		break
	case 9:
		var type_name *byte = zend.ZendRsrcListGetRsrcType(struc.value.res)
		core.PhpPrintf("%sresource(%d) of type (%s) refcount(%u)\n", g.Cond(is_ref != 0, "&", ""), struc.value.res.handle, g.Cond(type_name != nil, type_name, "Unknown"), zend.ZvalRefcountP(struc))
		break
	case 10:

		//??? hide references with refcount==1 (for compatibility)

		if zend.ZvalRefcountP(struc) > 1 {
			is_ref = 1
		}
		struc = &(*struc).value.ref.val
		goto again
	default:
		core.PhpPrintf("%sUNKNOWN:0\n", g.Cond(is_ref != 0, "&", ""))
		break
	}
}

/* }}} */

func ZifDebugZvalDump(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var args *zend.Zval
	var argc int
	var i int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = -1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
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
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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

/* }}} */

// #define buffer_append_spaces(buf,num_spaces) do { char * tmp_spaces ; size_t tmp_spaces_len ; tmp_spaces_len = spprintf ( & tmp_spaces , 0 , "%*c" , num_spaces , ' ' ) ; smart_str_appendl ( buf , tmp_spaces , tmp_spaces_len ) ; efree ( tmp_spaces ) ; } while ( 0 ) ;

func PhpArrayElementExport(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int, buf *zend.SmartStr) {
	if key == nil {
		var tmp_spaces *byte
		var tmp_spaces_len int
		tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level+1, ' ')
		zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
		zend._efree(tmp_spaces)
		zend.SmartStrAppendLongEx(buf, zend.ZendLong(index), 0)
		zend.SmartStrAppendlEx(buf, " => ", 4, 0)
	} else {
		var tmp_str *zend.ZendString
		var ckey *zend.ZendString = PhpAddcslashes(key, "'\\", 2)
		tmp_str = PhpStrToStr(ckey.val, ckey.len_, "0", 1, "' . \"\\0\" . '", 12)
		var tmp_spaces *byte
		var tmp_spaces_len int
		tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level+1, ' ')
		zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
		zend._efree(tmp_spaces)
		zend.SmartStrAppendcEx(buf, '\'', 0)
		zend.SmartStrAppendEx(buf, tmp_str, 0)
		zend.SmartStrAppendlEx(buf, "' => ", 5, 0)
		zend.ZendStringFree(ckey)
		zend.ZendStringFree(tmp_str)
	}
	PhpVarExportEx(zv, level+2, buf)
	zend.SmartStrAppendcEx(buf, ',', 0)
	zend.SmartStrAppendcEx(buf, '\n', 0)
}

/* }}} */

func PhpObjectElementExport(zv *zend.Zval, index zend.ZendUlong, key *zend.ZendString, level int, buf *zend.SmartStr) {
	var tmp_spaces *byte
	var tmp_spaces_len int
	tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level+2, ' ')
	zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
	zend._efree(tmp_spaces)
	if key != nil {
		var class_name *byte
		var prop_name *byte
		var prop_name_len int
		var pname_esc *zend.ZendString
		zend.ZendUnmanglePropertyNameEx(key, &class_name, &prop_name, &prop_name_len)
		pname_esc = PhpAddcslashesStr(prop_name, prop_name_len, "'\\", 2)
		zend.SmartStrAppendcEx(buf, '\'', 0)
		zend.SmartStrAppendEx(buf, pname_esc, 0)
		zend.SmartStrAppendcEx(buf, '\'', 0)
		zend.ZendStringReleaseEx(pname_esc, 0)
	} else {
		zend.SmartStrAppendLongEx(buf, zend.ZendLong(index), 0)
	}
	zend.SmartStrAppendlEx(buf, " => ", 4, 0)
	PhpVarExportEx(zv, level+2, buf)
	zend.SmartStrAppendcEx(buf, ',', 0)
	zend.SmartStrAppendcEx(buf, '\n', 0)
}

/* }}} */

func PhpVarExportEx(struc *zend.Zval, level int, buf *zend.SmartStr) {
	var myht *zend.HashTable
	var tmp_str []byte
	var ztmp *zend.ZendString
	var ztmp2 *zend.ZendString
	var index zend.ZendUlong
	var key *zend.ZendString
	var val *zend.Zval
again:
	switch struc.u1.v.type_ {
	case 2:
		zend.SmartStrAppendlEx(buf, "false", 5, 0)
		break
	case 3:
		zend.SmartStrAppendlEx(buf, "true", 4, 0)
		break
	case 1:
		zend.SmartStrAppendlEx(buf, "NULL", 4, 0)
		break
	case 4:

		/* INT_MIN as a literal will be parsed as a float. Emit something like
		 * -9223372036854775807-1 to avoid this. */

		if struc.value.lval == INT64_MIN {
			zend.SmartStrAppendLongEx(buf, INT64_MIN+1, 0)
			zend.SmartStrAppendlEx(buf, "-1", strlen("-1"), 0)
			break
		}
		zend.SmartStrAppendLongEx(buf, struc.value.lval, 0)
		break
	case 5:
		core.PhpGcvt(struc.value.dval, int(core.CoreGlobals.serialize_precision), '.', 'E', tmp_str)
		zend.SmartStrAppendlEx(buf, tmp_str, strlen(tmp_str), 0)

		/* Without a decimal point, PHP treats a number literal as an int.
		 * This check even works for scientific notation, because the
		 * mantissa always contains a decimal point.
		 * We need to check for finiteness, because INF, -INF and NAN
		 * must not have a decimal point added.
		 */

		if isfinite(struc.value.dval) && nil == strchr(tmp_str, '.') {
			zend.SmartStrAppendlEx(buf, ".0", 2, 0)
		}
		break
	case 6:
		ztmp = PhpAddcslashes(struc.value.str, "'\\", 2)
		ztmp2 = PhpStrToStr(ztmp.val, ztmp.len_, "0", 1, "' . \"\\0\" . '", 12)
		zend.SmartStrAppendcEx(buf, '\'', 0)
		zend.SmartStrAppendEx(buf, ztmp2, 0)
		zend.SmartStrAppendcEx(buf, '\'', 0)
		zend.ZendStringFree(ztmp)
		zend.ZendStringFree(ztmp2)
		break
	case 7:
		myht = struc.value.arr
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 5) != 0 {
				zend.SmartStrAppendlEx(buf, "NULL", 4, 0)
				zend.ZendError(1<<1, "var_export does not handle circular references")
				return
			}
			zend.ZendGcAddref(&myht.gc)
			myht.gc.u.type_info |= 1 << 5 << 0
		}
		if level > 1 {
			zend.SmartStrAppendcEx(buf, '\n', 0)
			var tmp_spaces *byte
			var tmp_spaces_len int
			tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level-1, ' ')
			zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
			zend._efree(tmp_spaces)
		}
		zend.SmartStrAppendlEx(buf, "array (\n", 8, 0)
		for {
			var __ht *zend.HashTable = myht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				index = _p.h
				key = _p.key
				val = _z
				PhpArrayElementExport(val, index, key, level, buf)
			}
			break
		}
		if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
			myht.gc.u.type_info &= ^(1 << 5 << 0)
			zend.ZendGcDelref(&myht.gc)
		}
		if level > 1 {
			var tmp_spaces *byte
			var tmp_spaces_len int
			tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level-1, ' ')
			zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
			zend._efree(tmp_spaces)
		}
		zend.SmartStrAppendcEx(buf, ')', 0)
		break
	case 8:
		myht = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_VAR_EXPORT)
		if myht != nil {
			if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 5) != 0 {
				zend.SmartStrAppendlEx(buf, "NULL", 4, 0)
				zend.ZendError(1<<1, "var_export does not handle circular references")
				if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
					zend.ZendArrayDestroy(myht)
				}
				return
			} else {
				if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
					myht.gc.u.type_info |= 1 << 5 << 0
				}
			}
		}
		if level > 1 {
			zend.SmartStrAppendcEx(buf, '\n', 0)
			var tmp_spaces *byte
			var tmp_spaces_len int
			tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level-1, ' ')
			zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
			zend._efree(tmp_spaces)
		}

		/* stdClass has no __set_state method, but can be casted to */

		if struc.value.obj.ce == zend.ZendStandardClassDef {
			zend.SmartStrAppendlEx(buf, "(object) array(\n", 16, 0)
		} else {
			zend.SmartStrAppendEx(buf, struc.value.obj.ce.name, 0)
			zend.SmartStrAppendlEx(buf, "::__set_state(array(\n", 21, 0)
		}
		if myht != nil {
			for {
				var __ht *zend.HashTable = myht
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					index = _p.h
					key = _p.key
					val = _z
					PhpObjectElementExport(val, index, key, level, buf)
				}
				break
			}
			if (zend.ZvalGcFlags(myht.gc.u.type_info) & 1 << 6) == 0 {
				myht.gc.u.type_info &= ^(1 << 5 << 0)
			}
			if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
				zend.ZendArrayDestroy(myht)
			}
		}
		if level > 1 {
			var tmp_spaces *byte
			var tmp_spaces_len int
			tmp_spaces_len = zend.ZendSpprintf(&tmp_spaces, 0, "%*c", level-1, ' ')
			zend.SmartStrAppendlEx(buf, tmp_spaces, tmp_spaces_len, 0)
			zend._efree(tmp_spaces)
		}
		if struc.value.obj.ce == zend.ZendStandardClassDef {
			zend.SmartStrAppendcEx(buf, ')', 0)
		} else {
			zend.SmartStrAppendlEx(buf, "))", 2, 0)
		}
		break
	case 10:
		struc = &(*struc).value.ref.val
		goto again
		break
	default:
		zend.SmartStrAppendlEx(buf, "NULL", 4, 0)
		break
	}
}

/* }}} */

func PhpVarExport(struc *zend.Zval, level int) {
	var buf zend.SmartStr = zend.SmartStr{0}
	PhpVarExportEx(struc, level, &buf)
	zend.SmartStr0(&buf)
	core.PhpOutputWrite(buf.s.val, buf.s.len_)
	zend.SmartStrFreeEx(&buf, 0)
}

/* }}} */

func ZifVarExport(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var var_ *zend.Zval
	var return_output zend.ZendBool = 0
	var buf zend.SmartStr = zend.SmartStr{0}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &var_, 0)
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &return_output, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = buf.s
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		core.PhpOutputWrite(buf.s.val, buf.s.len_)
		zend.SmartStrFreeEx(&buf, 0)
	}
}

/* }}} */

func PhpAddVarHash(data PhpSerializeDataT, var_ *zend.Zval) zend.ZendLong {
	var zv *zend.Zval
	var key zend.ZendUlong
	var is_ref zend.ZendBool = var_.u1.v.type_ == 10
	data.SetN(data.GetN() + 1)
	if is_ref == 0 && var_.u1.v.type_ != 8 {
		return 0
	}

	/* References to objects are treated as if the reference didn't exist */

	if is_ref != 0 && &(*var_).value.ref.val.u1.v.type_ == 8 {
		var_ = &(*var_).value.ref.val
	}

	/* Index for the variable is stored using the numeric value of the pointer to
	 * the zend_refcounted struct */

	key = zend_ulong(zend_uintptr_t)(*var_).value.counted
	zv = zend.ZendHashIndexFind(&data.ht, key)
	if zv != nil {

		/* References are only counted once, undo the data->n increment above */

		if is_ref != 0 && zv.value.lval != -1 {
			data.SetN(data.GetN() - 1)
		}
		return zv.value.lval
	} else {
		var zv_n zend.Zval
		var __z *zend.Zval = &zv_n
		__z.value.lval = data.GetN()
		__z.u1.type_info = 4
		zend.ZendHashIndexAddNew(&data.ht, key, &zv_n)

		/* Additionally to the index, we also store the variable, to ensure that it is
		 * not destroyed during serialization and its pointer reused. The variable is
		 * stored at the numeric value of the pointer + 1, which cannot be the location
		 * of another zend_refcounted structure. */

		zend.ZendHashIndexAddNew(&data.ht, key+1, var_)
		zend.ZvalAddrefP(var_)
		return 0
	}
}

/* }}} */

func PhpVarSerializeLong(buf *zend.SmartStr, val zend.ZendLong) {
	zend.SmartStrAppendlEx(buf, "i:", 2, 0)
	zend.SmartStrAppendLongEx(buf, val, 0)
	zend.SmartStrAppendcEx(buf, ';', 0)
}

/* }}} */

func PhpVarSerializeString(buf *zend.SmartStr, str *byte, len_ int) {
	zend.SmartStrAppendlEx(buf, "s:", 2, 0)
	zend.SmartStrAppendUnsignedEx(buf, len_, 0)
	zend.SmartStrAppendlEx(buf, ":\"", 2, 0)
	zend.SmartStrAppendlEx(buf, str, len_, 0)
	zend.SmartStrAppendlEx(buf, "\";", 2, 0)
}

/* }}} */

func PhpVarSerializeClassName(buf *zend.SmartStr, struc *zend.Zval) zend.ZendBool {
	var class_name *zend.ZendString
	var incomplete_class zend.ZendBool = 0
	if struc.value.obj.ce == BasicGlobals.GetIncompleteClass() {
		class_name = PhpLookupClassName(struc)
		if class_name == nil {
			class_name = zend.ZendStringInit("__PHP_Incomplete_Class", g.SizeOf("INCOMPLETE_CLASS")-1, 0)
		}
		incomplete_class = 1
	} else {
		class_name = zend.ZendStringCopy(struc.value.obj.ce.name)
	}
	zend.SmartStrAppendlEx(buf, "O:", 2, 0)
	zend.SmartStrAppendUnsignedEx(buf, class_name.len_, 0)
	zend.SmartStrAppendlEx(buf, ":\"", 2, 0)
	zend.SmartStrAppendEx(buf, class_name, 0)
	zend.SmartStrAppendlEx(buf, "\":", 2, 0)
	zend.ZendStringReleaseEx(class_name, 0)
	return incomplete_class
}

/* }}} */

func PhpVarSerializeCallSleep(retval *zend.Zval, struc *zend.Zval) int {
	var fname zend.Zval
	var res int
	var __z *zend.Zval = &fname
	var __s *zend.ZendString = zend.ZendStringInit("__sleep", g.SizeOf("\"__sleep\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	BasicGlobals.GetSerializeLock()++
	res = zend._callUserFunctionEx(struc, &fname, retval, 0, 0, 1)
	BasicGlobals.GetSerializeLock()--
	zend.ZvalPtrDtorStr(&fname)
	if res == zend.FAILURE || retval.u1.v.type_ == 0 {
		zend.ZvalPtrDtor(retval)
		return zend.FAILURE
	}
	if !(g.CondF(retval.u1.v.type_ == 7, func() *zend.ZendArray { return retval.value.arr }, func() __auto__ {
		if retval.u1.v.type_ == 8 {
			return retval.value.obj.handlers.get_properties(retval)
		} else {
			return nil
		}
	})) {
		zend.ZvalPtrDtor(retval)
		core.PhpErrorDocref(nil, 1<<3, "__sleep should return an array only containing the names of instance-variables to serialize")
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func PhpVarSerializeCallMagicSerialize(retval *zend.Zval, obj *zend.Zval) int {
	var fname zend.Zval
	var res int
	var __z *zend.Zval = &fname
	var __s *zend.ZendString = zend.ZendStringInit("__serialize", g.SizeOf("\"__serialize\"")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	BasicGlobals.GetSerializeLock()++
	res = zend._callUserFunctionEx(obj, &fname, retval, 0, 0, 1)
	BasicGlobals.GetSerializeLock()--
	zend.ZvalPtrDtorStr(&fname)
	if res == zend.FAILURE || retval.u1.v.type_ == 0 {
		zend.ZvalPtrDtor(retval)
		return zend.FAILURE
	}
	if retval.u1.v.type_ != 7 {
		zend.ZvalPtrDtor(retval)
		zend.ZendTypeError("%s::__serialize() must return an array", obj.value.obj.ce.name.val)
		return zend.FAILURE
	}
	return zend.SUCCESS
}

/* }}} */

func PhpVarSerializeTryAddSleepProp(ht *zend.HashTable, props *zend.HashTable, name *zend.ZendString, error_name *zend.ZendString, struc *zend.Zval) int {
	var val *zend.Zval = zend.ZendHashFind(props, name)
	if val == nil {
		return zend.FAILURE
	}
	if val.u1.v.type_ == 13 {
		val = val.value.zv
		if val.u1.v.type_ == 0 {
			var info *zend.ZendPropertyInfo = zend.ZendGetTypedPropertyInfoForSlot(struc.value.obj, val)
			if info != nil {
				return zend.SUCCESS
			}
			return zend.FAILURE
		}
	}
	if zend.ZendHashAdd(ht, name, val) == nil {
		core.PhpErrorDocref(nil, 1<<3, "\"%s\" is returned from __sleep multiple times", error_name.val)
		return zend.SUCCESS
	}
	if val.u1.v.type_flags != 0 {
		zend.ZvalAddrefP(val)
	}
	return zend.SUCCESS
}

/* }}} */

func PhpVarSerializeGetSleepProps(ht *zend.HashTable, struc *zend.Zval, sleep_retval *zend.HashTable) int {
	var ce *zend.ZendClassEntry = struc.value.obj.ce
	var props *zend.HashTable = zend.ZendGetPropertiesFor(struc, zend.ZEND_PROP_PURPOSE_SERIALIZE)
	var name_val *zend.Zval
	var retval int = zend.SUCCESS
	zend._zendHashInit(ht, sleep_retval.nNumOfElements, zend.ZvalPtrDtor, 0)

	/* TODO: Rewrite this by fetching the property info instead of trying out different
	 * name manglings? */

	for {
		var __ht *zend.HashTable = sleep_retval
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val
			if _z.u1.v.type_ == 13 {
				_z = _z.value.zv
			}
			if _z.u1.v.type_ == 0 {
				continue
			}
			name_val = _z
			var name *zend.ZendString
			var tmp_name *zend.ZendString
			var priv_name *zend.ZendString
			var prot_name *zend.ZendString
			if name_val.u1.v.type_ == 10 {
				name_val = &(*name_val).value.ref.val
			}
			if name_val.u1.v.type_ != 6 {
				core.PhpErrorDocref(nil, 1<<3, "__sleep should return an array only containing the names of instance-variables to serialize.")
			}
			name = zend.ZvalGetTmpString(name_val, &tmp_name)
			if PhpVarSerializeTryAddSleepProp(ht, props, name, name, struc) == zend.SUCCESS {
				zend.ZendTmpStringRelease(tmp_name)
				continue
			}
			if zend.EG.exception != nil {
				zend.ZendTmpStringRelease(tmp_name)
				retval = zend.FAILURE
				break
			}
			priv_name = zend.ZendManglePropertyName(ce.name.val, ce.name.len_, name.val, name.len_, ce.type_&1)
			if PhpVarSerializeTryAddSleepProp(ht, props, priv_name, name, struc) == zend.SUCCESS {
				zend.ZendTmpStringRelease(tmp_name)
				zend.ZendStringRelease(priv_name)
				continue
			}
			zend.ZendStringRelease(priv_name)
			if zend.EG.exception != nil {
				zend.ZendTmpStringRelease(tmp_name)
				retval = zend.FAILURE
				break
			}
			prot_name = zend.ZendManglePropertyName("*", 1, name.val, name.len_, ce.type_&1)
			if PhpVarSerializeTryAddSleepProp(ht, props, prot_name, name, struc) == zend.SUCCESS {
				zend.ZendTmpStringRelease(tmp_name)
				zend.ZendStringRelease(prot_name)
				continue
			}
			zend.ZendStringRelease(prot_name)
			if zend.EG.exception != nil {
				zend.ZendTmpStringRelease(tmp_name)
				retval = zend.FAILURE
				break
			}
			core.PhpErrorDocref(nil, 1<<3, "\"%s\" returned as member variable from __sleep() but does not exist", name.val)
			zend.ZendHashAdd(ht, name, &zend.EG.uninitialized_zval)
			zend.ZendTmpStringRelease(tmp_name)
		}
		break
	}
	if props != nil && (zend.ZvalGcFlags(props.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&props.gc) == 0 {
		zend.ZendArrayDestroy(props)
	}
	return retval
}

/* }}} */

func PhpVarSerializeNestedData(buf *zend.SmartStr, struc *zend.Zval, ht *zend.HashTable, count uint32, incomplete_class zend.ZendBool, var_hash PhpSerializeDataT) {
	zend.SmartStrAppendUnsignedEx(buf, count, 0)
	zend.SmartStrAppendlEx(buf, ":{", 2, 0)
	if count > 0 {
		var key *zend.ZendString
		var data *zend.Zval
		var index zend.ZendUlong
		for {
			var __ht *zend.HashTable = ht
			var _p *zend.Bucket = __ht.arData
			var _end *zend.Bucket = _p + __ht.nNumUsed
			for ; _p != _end; _p++ {
				var _z *zend.Zval = &_p.val
				if _z.u1.v.type_ == 13 {
					_z = _z.value.zv
				}
				if _z.u1.v.type_ == 0 {
					continue
				}
				index = _p.h
				key = _p.key
				data = _z
				if incomplete_class != 0 && strcmp(key.val, "__PHP_Incomplete_Class_Name") == 0 {
					continue
				}
				if key == nil {
					PhpVarSerializeLong(buf, index)
				} else {
					PhpVarSerializeString(buf, key.val, key.len_)
				}
				if data.u1.v.type_ == 10 && zend.ZvalRefcountP(data) == 1 {
					data = &(*data).value.ref.val
				}

				/* we should still add element even if it's not OK,
				 * since we already wrote the length of the array before */

				if data.u1.v.type_ == 7 {
					if (zend.ZvalGcFlags(data.value.counted.gc.u.type_info)&1<<5) != 0 || struc.u1.v.type_ == 7 && data.value.arr == struc.value.arr {
						PhpAddVarHash(var_hash, struc)
						zend.SmartStrAppendlEx(buf, "N;", 2, 0)
					} else {
						if data.u1.v.type_flags != 0 {
							data.value.counted.gc.u.type_info |= 1 << 5 << 0
						}
						PhpVarSerializeIntern(buf, data, var_hash)
						if data.u1.v.type_flags != 0 {
							data.value.counted.gc.u.type_info &= ^(1 << 5 << 0)
						}
					}
				} else {
					PhpVarSerializeIntern(buf, data, var_hash)
				}

				/* we should still add element even if it's not OK,
				 * since we already wrote the length of the array before */

			}
			break
		}
	}
	zend.SmartStrAppendcEx(buf, '}', 0)
}

/* }}} */

func PhpVarSerializeClass(buf *zend.SmartStr, struc *zend.Zval, retval_ptr *zend.Zval, var_hash PhpSerializeDataT) {
	var props zend.HashTable
	if PhpVarSerializeGetSleepProps(&props, struc, g.CondF(retval_ptr.u1.v.type_ == 7, func() *zend.ZendArray { return retval_ptr.value.arr }, func() __auto__ {
		if retval_ptr.u1.v.type_ == 8 {
			return retval_ptr.value.obj.handlers.get_properties(retval_ptr)
		} else {
			return nil
		}
	})) == zend.SUCCESS {
		PhpVarSerializeClassName(buf, struc)
		PhpVarSerializeNestedData(buf, struc, &props, &props.nNumOfElements, 0, var_hash)
	}
	zend.ZendHashDestroy(&props)
}

/* }}} */

func PhpVarSerializeIntern(buf *zend.SmartStr, struc *zend.Zval, var_hash PhpSerializeDataT) {
	var var_already zend.ZendLong
	var myht *zend.HashTable
	if zend.EG.exception != nil {
		return
	}
	if var_hash != nil && g.Assign(&var_already, PhpAddVarHash(var_hash, struc)) {
		if var_already == -1 {

			/* Reference to an object that failed to serialize, replace with null. */

			zend.SmartStrAppendlEx(buf, "N;", 2, 0)
			return
		} else if struc.u1.v.type_ == 10 {
			zend.SmartStrAppendlEx(buf, "R:", 2, 0)
			zend.SmartStrAppendLongEx(buf, var_already, 0)
			zend.SmartStrAppendcEx(buf, ';', 0)
			return
		} else if struc.u1.v.type_ == 8 {
			zend.SmartStrAppendlEx(buf, "r:", 2, 0)
			zend.SmartStrAppendLongEx(buf, var_already, 0)
			zend.SmartStrAppendcEx(buf, ';', 0)
			return
		}
	}
again:
	switch struc.u1.v.type_ {
	case 2:
		zend.SmartStrAppendlEx(buf, "b:0;", 4, 0)
		return
	case 3:
		zend.SmartStrAppendlEx(buf, "b:1;", 4, 0)
		return
	case 1:
		zend.SmartStrAppendlEx(buf, "N;", 2, 0)
		return
	case 4:
		PhpVarSerializeLong(buf, struc.value.lval)
		return
	case 5:
		var tmp_str []byte
		zend.SmartStrAppendlEx(buf, "d:", 2, 0)
		core.PhpGcvt(struc.value.dval, int(core.CoreGlobals.serialize_precision), '.', 'E', tmp_str)
		zend.SmartStrAppendlEx(buf, tmp_str, strlen(tmp_str), 0)
		zend.SmartStrAppendcEx(buf, ';', 0)
		return
	case 6:
		PhpVarSerializeString(buf, struc.value.str.val, struc.value.str.len_)
		return
	case 8:
		var ce *zend.ZendClassEntry = struc.value.obj.ce
		var incomplete_class zend.ZendBool
		var count uint32
		if zend.ZendHashStrExists(&ce.function_table, "__serialize", g.SizeOf("\"__serialize\"")-1) != 0 {
			var retval zend.Zval
			var obj zend.Zval
			var key *zend.ZendString
			var data *zend.Zval
			var index zend.ZendUlong
			zend.ZvalAddrefP(struc)
			var __z *zend.Zval = &obj
			__z.value.obj = struc.value.obj
			__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
			if PhpVarSerializeCallMagicSerialize(&retval, &obj) == zend.FAILURE {
				if zend.EG.exception == nil {
					zend.SmartStrAppendlEx(buf, "N;", 2, 0)
				}
				zend.ZvalPtrDtor(&obj)
				return
			}
			PhpVarSerializeClassName(buf, &obj)
			zend.SmartStrAppendUnsignedEx(buf, zend.ZendArrayCount(retval.value.arr), 0)
			zend.SmartStrAppendlEx(buf, ":{", 2, 0)
			for {
				var __ht *zend.HashTable = retval.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val
					if _z.u1.v.type_ == 13 {
						_z = _z.value.zv
					}
					if _z.u1.v.type_ == 0 {
						continue
					}
					index = _p.h
					key = _p.key
					data = _z
					if key == nil {
						PhpVarSerializeLong(buf, index)
					} else {
						PhpVarSerializeString(buf, key.val, key.len_)
					}
					if data.u1.v.type_ == 10 && zend.ZvalRefcountP(data) == 1 {
						data = &(*data).value.ref.val
					}
					PhpVarSerializeIntern(buf, data, var_hash)
				}
				break
			}
			zend.SmartStrAppendcEx(buf, '}', 0)
			zend.ZvalPtrDtor(&obj)
			zend.ZvalPtrDtor(&retval)
			return
		}
		if ce.serialize != nil {

			/* has custom handler */

			var serialized_data *uint8 = nil
			var serialized_length int
			if ce.serialize(struc, &serialized_data, &serialized_length, (*zend.ZendSerializeData)(var_hash)) == zend.SUCCESS {
				zend.SmartStrAppendlEx(buf, "C:", 2, 0)
				zend.SmartStrAppendUnsignedEx(buf, struc.value.obj.ce.name.len_, 0)
				zend.SmartStrAppendlEx(buf, ":\"", 2, 0)
				zend.SmartStrAppendEx(buf, struc.value.obj.ce.name, 0)
				zend.SmartStrAppendlEx(buf, "\":", 2, 0)
				zend.SmartStrAppendUnsignedEx(buf, serialized_length, 0)
				zend.SmartStrAppendlEx(buf, ":{", 2, 0)
				zend.SmartStrAppendlEx(buf, (*byte)(serialized_data), serialized_length, 0)
				zend.SmartStrAppendcEx(buf, '}', 0)
			} else {

				/* Mark this value in the var_hash, to avoid creating references to it. */

				var var_idx *zend.Zval = zend.ZendHashIndexFind(&var_hash.ht, zend_ulong(zend_uintptr_t)(*struc).value.counted)
				var __z *zend.Zval = var_idx
				__z.value.lval = -1
				__z.u1.type_info = 4
				zend.SmartStrAppendlEx(buf, "N;", 2, 0)
			}
			if serialized_data != nil {
				zend._efree(serialized_data)
			}
			return
		}
		if ce != BasicGlobals.GetIncompleteClass() && zend.ZendHashStrExists(&ce.function_table, "__sleep", g.SizeOf("\"__sleep\"")-1) != 0 {
			var retval zend.Zval
			var tmp zend.Zval
			zend.ZvalAddrefP(struc)
			var __z *zend.Zval = &tmp
			__z.value.obj = struc.value.obj
			__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
			if PhpVarSerializeCallSleep(&retval, &tmp) == zend.FAILURE {
				if zend.EG.exception == nil {

					/* we should still add element even if it's not OK,
					 * since we already wrote the length of the array before */

					zend.SmartStrAppendlEx(buf, "N;", 2, 0)

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

		count = zend.ZendArrayCount(myht)
		if count > 0 && incomplete_class != 0 {
			count--
		}
		PhpVarSerializeNestedData(buf, struc, myht, count, incomplete_class, var_hash)
		if myht != nil && (zend.ZvalGcFlags(myht.gc.u.type_info)&1<<6) == 0 && zend.ZendGcDelref(&myht.gc) == 0 {
			zend.ZendArrayDestroy(myht)
		}
		return
	case 7:
		zend.SmartStrAppendlEx(buf, "a:", 2, 0)
		myht = struc.value.arr
		PhpVarSerializeNestedData(buf, struc, myht, zend.ZendArrayCount(myht), 0, var_hash)
		return
	case 10:
		struc = &(*struc).value.ref.val
		goto again
	default:
		zend.SmartStrAppendlEx(buf, "i:0;", 4, 0)
		return
	}
}

/* }}} */

func PhpVarSerialize(buf *zend.SmartStr, struc *zend.Zval, data *PhpSerializeDataT) {
	PhpVarSerializeIntern(buf, struc, *data)
	zend.SmartStr0(buf)
}

/* }}} */

func PhpVarSerializeInit() PhpSerializeDataT {
	var d *PhpSerializeData

	/* fprintf(stderr, "SERIALIZE_INIT      == lock: %u, level: %u\n", BG(serialize_lock), BG(serialize).level); */

	if BasicGlobals.GetSerializeLock() || !(BasicGlobals.GetSerializeLevel()) {
		d = zend._emalloc(g.SizeOf("struct php_serialize_data"))
		zend._zendHashInit(&d.ht, 16, zend.ZvalPtrDtor, 0)
		d.SetN(0)
		if !(BasicGlobals.GetSerializeLock()) {
			BasicGlobals.SetSerializeData(d)
			BasicGlobals.SetSerializeLevel(1)
		}
	} else {
		d = BasicGlobals.GetSerializeData()
		BasicGlobals.GetSerializeLevel()++
	}
	return d
}
func PhpVarSerializeDestroy(d PhpSerializeDataT) {
	/* fprintf(stderr, "SERIALIZE_DESTROY   == lock: %u, level: %u\n", BG(serialize_lock), BG(serialize).level); */

	if BasicGlobals.GetSerializeLock() || BasicGlobals.GetSerializeLevel() == 1 {
		zend.ZendHashDestroy(&d.ht)
		zend._efree(d)
	}
	if !(BasicGlobals.GetSerializeLock()) && !(g.PreDec(&(BasicGlobals.GetSerializeLevel()))) {
		BasicGlobals.SetSerializeData(nil)
	}
}

/* {{{ proto string serialize(mixed variable)
   Returns a string representation of variable (which can later be unserialized) */

func ZifSerialize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var struc *zend.Zval
	var var_hash PhpSerializeDataT
	var buf zend.SmartStr = zend.SmartStr{0}
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			zend.ZendParseArgZvalDeref(_arg, &struc, 0)
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
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
	var_hash = PhpVarSerializeInit()
	PhpVarSerialize(&buf, struc, &var_hash)
	PhpVarSerializeDestroy(var_hash)
	if zend.EG.exception != nil {
		zend.SmartStrFreeEx(&buf, 0)
		return_value.u1.type_info = 2
		return
	}
	if buf.s != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = buf.s
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 1
		return
	}
}

/* }}} */

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
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgString(_arg, &buf, &buf_len, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgArray(_arg, &options, 0, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_ARRAY
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	if buf_len == 0 {
		return_value.u1.type_info = 2
		return
	}
	p = (*uint8)(buf)
	var_hash = PhpVarUnserializeInit()
	prev_class_hash = PhpVarUnserializeGetAllowedClasses(var_hash)
	prev_max_depth = PhpVarUnserializeGetMaxDepth(var_hash)
	prev_cur_depth = PhpVarUnserializeGetCurDepth(var_hash)
	if options != nil {
		var classes *zend.Zval
		var max_depth *zend.Zval
		classes = zend.ZendHashStrFindDeref(options.value.arr, "allowed_classes", g.SizeOf("\"allowed_classes\"")-1)
		if classes != nil && classes.u1.v.type_ != 7 && classes.u1.v.type_ != 3 && classes.u1.v.type_ != 2 {
			core.PhpErrorDocref(nil, 1<<1, "allowed_classes option should be array or boolean")
			return_value.u1.type_info = 2
			goto cleanup
		}
		if classes != nil && (classes.u1.v.type_ == 7 || zend.ZendIsTrue(classes) == 0) {
			class_hash = (*zend.HashTable)(zend._emalloc(g.SizeOf("HashTable")))
			zend._zendHashInit(class_hash, g.CondF1(classes.u1.v.type_ == 7, func() uint32 { return classes.value.arr.nNumOfElements }, 0), nil, 0)
		}
		if class_hash != nil && classes.u1.v.type_ == 7 {
			var entry *zend.Zval
			var lcname *zend.ZendString
			for {
				var __ht *zend.HashTable = classes.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					entry = _z
					if entry.u1.v.type_ != 6 {
						if entry.u1.v.type_ != 6 {
							zend._convertToString(entry)
						}
					}
					lcname = zend.ZendStringTolowerEx(entry.value.str, 0)
					zend.ZendHashAddEmptyElement(class_hash, lcname)
					zend.ZendStringReleaseEx(lcname, 0)
				}
				break
			}

			/* Exception during string conversion. */

			if zend.EG.exception != nil {
				goto cleanup
			}

			/* Exception during string conversion. */

		}
		PhpVarUnserializeSetAllowedClasses(var_hash, class_hash)
		max_depth = zend.ZendHashStrFindDeref(options.value.arr, "max_depth", g.SizeOf("\"max_depth\"")-1)
		if max_depth != nil {
			if max_depth.u1.v.type_ != 4 {
				core.PhpErrorDocref(nil, 1<<1, "max_depth should be int")
				return_value.u1.type_info = 2
				goto cleanup
			}
			if max_depth.value.lval < 0 {
				core.PhpErrorDocref(nil, 1<<1, "max_depth cannot be negative")
				return_value.u1.type_info = 2
				goto cleanup
			}
			PhpVarUnserializeSetMaxDepth(var_hash, max_depth.value.lval)

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

			PhpVarUnserializeSetCurDepth(var_hash, 0)

			/* If the max_depth for a nested unserialize() call has been overridden,
			 * start counting from zero again (for the nested call only). */

		}
	}
	if BasicGlobals.GetUnserializeLevel() > 1 {
		retval = VarTmpVar(&var_hash)
	} else {
		retval = return_value
	}
	if PhpVarUnserialize(retval, &p, p+buf_len, &var_hash) == 0 {
		if zend.EG.exception == nil {
			core.PhpErrorDocref(nil, 1<<3, "Error at offset "+"%"+"lld"+" of %zd bytes", zend_long((*byte)(p-buf)), buf_len)
		}
		if BasicGlobals.GetUnserializeLevel() <= 1 {
			zend.ZvalPtrDtor(return_value)
		}
		return_value.u1.type_info = 2
	} else if BasicGlobals.GetUnserializeLevel() > 1 {
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = retval
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}
	} else if return_value.u1.v.type_flags != 0 {
		var ref *zend.ZendRefcounted = return_value.value.counted
		zend.GcCheckPossibleRoot(ref)
	}
cleanup:
	if class_hash != nil {
		zend.ZendHashDestroy(class_hash)
		zend._efree(class_hash)
	}

	/* Reset to previous options in case this is a nested call */

	PhpVarUnserializeSetAllowedClasses(var_hash, prev_class_hash)
	PhpVarUnserializeSetMaxDepth(var_hash, prev_max_depth)
	PhpVarUnserializeSetCurDepth(var_hash, prev_cur_depth)
	PhpVarUnserializeDestroy(var_hash)

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */

	if return_value.u1.v.type_ == 10 {
		zend.ZendUnwrapReference(return_value)
	}

	/* Per calling convention we must not return a reference here, so unwrap. We're doing this at
	 * the very end, because __wakeup() calls performed during UNSERIALIZE_DESTROY might affect
	 * the value we unwrap here. This is compatible with behavior in PHP <=7.0. */
}

/* }}} */

func ZifMemoryGetUsage(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var real_usage zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &real_usage, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	__z.value.lval = zend.ZendMemoryUsage(real_usage)
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZifMemoryGetPeakUsage(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var real_usage zend.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = execute_data.This.u2.num_args
		var _i int = 0
		var _real_arg *zend.Zval
		var _arg *zend.Zval = nil
		var _expected_type zend.ZendExpectedType = zend.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy zend.ZendBool
		var _optional zend.ZendBool = 0
		var _error_code int = 0
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			if _num_args < _min_num_args || _num_args > _max_num_args && _max_num_args >= 0 {
				if (_flags & 1 << 1) == 0 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParametersCountException(_min_num_args, _max_num_args)
					} else {
						zend.ZendWrongParametersCountError(_min_num_args, _max_num_args)
					}
				}
				_error_code = 1
				break
			}
			_real_arg = (*zend.Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(int(0)-1))
			_optional = 1
			_i++
			assert(_i <= _min_num_args || _optional == 1)
			assert(_i > _min_num_args || _optional == 0)
			if _optional != 0 {
				if _i > _num_args {
					break
				}
			}
			_real_arg++
			_arg = _real_arg

			if zend.ZendParseArgBool(_arg, &real_usage, &_dummy, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_BOOL
				_error_code = 4
				break
			}
			break
		}
		if _error_code != 0 {
			if (_flags & 1 << 1) == 0 {
				if _error_code == 2 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongCallbackException(_i, _error)
					} else {
						zend.ZendWrongCallbackError(_i, _error)
					}
				} else if _error_code == 3 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterClassException(_i, _error, _arg)
					} else {
						zend.ZendWrongParameterClassError(_i, _error, _arg)
					}
				} else if _error_code == 4 {
					if (_flags & 1 << 2) != 0 {
						zend.ZendWrongParameterTypeException(_i, _expected_type, _arg)
					} else {
						zend.ZendWrongParameterTypeError(_i, _expected_type, _arg)
					}
				}
			}
			return_value.u1.type_info = 2
			return
		}
		break
	}
	var __z *zend.Zval = return_value
	__z.value.lval = zend.ZendMemoryPeakUsage(real_usage)
	__z.u1.type_info = 4
	return
}

/* }}} */

func ZmStartupVar(type_ int, module_number int) int {
	zend.ZendRegisterIniEntries(IniEntries, module_number)
	return zend.SUCCESS
}
