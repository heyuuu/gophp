// <<generate>>

package core

import (
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/php_variables.h>

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
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define PHP_VARIABLES_H

// # include "php.h"

// # include "SAPI.h"

// #define PARSE_POST       0

// #define PARSE_GET       1

// #define PARSE_COOKIE       2

// #define PARSE_STRING       3

// #define PARSE_ENV       4

// #define PARSE_SERVER       5

// #define PARSE_SESSION       6

/* binary-safe version */

// #define NUM_TRACK_VARS       6

// Source: <main/php_variables.c>

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
   | Authors: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                       |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// # include < stdio . h >

// # include "php.h"

// # include "ext/standard/php_standard.h"

// # include "ext/standard/credits.h"

// # include "zend_smart_str.h"

// # include "php_variables.h"

// # include "php_globals.h"

// # include "php_content_types.h"

// # include "SAPI.h"

// # include "zend_globals.h"

/* for systems that need to override reading of environment variables */

var PhpImportEnvironmentVariables func(array_ptr *zend.Zval) = _phpImportEnvironmentVariables

func PhpRegisterVariable(var_ string, strval *byte, track_vars_array *zend.Zval) {
	PhpRegisterVariableSafe(var_, strval, strlen(strval), track_vars_array)
}

/* binary-safe version */

func PhpRegisterVariableSafe(var_ *byte, strval *byte, str_len int, track_vars_array *zend.Zval) {
	var new_entry zend.Zval
	r.Assert(strval != nil)

	/* Prepare value */

	if str_len == 0 {
		var __z *zend.Zval = &new_entry
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
	} else if str_len == 1 {
		var __z *zend.Zval = &new_entry
		var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar*strval]
		__z.value.str = __s
		__z.u1.type_info = 6
	} else {
		var __z *zend.Zval = &new_entry
		var __s *zend.ZendString = zend.ZendStringInit(strval, str_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
	PhpRegisterVariableEx(var_, &new_entry, track_vars_array)
}
func PhpRegisterVariableQuick(name *byte, name_len int, val *zend.Zval, ht *zend.HashTable) {
	var key *zend.ZendString = zend.ZendStringInitInterned(name, name_len, 0)
	zend.ZendHashUpdateInd(ht, key, val)
	zend.ZendStringReleaseEx(key, 0)
}
func PhpRegisterVariableEx(var_name *byte, val *zend.Zval, track_vars_array *zend.Zval) {
	var p *byte = nil
	var ip *byte = nil
	var index *byte
	var var_ *byte
	var var_orig *byte
	var var_len int
	var index_len int
	var gpc_element zend.Zval
	var gpc_element_p *zend.Zval
	var is_array zend.ZendBool = 0
	var symtable1 *zend.HashTable = nil
	r.Assert(var_name != nil)
	if track_vars_array != nil && track_vars_array.u1.v.type_ == 7 {
		symtable1 = track_vars_array.value.arr
	}
	if symtable1 == nil {

		/* Nothing to do */

		zend.ZvalPtrDtorNogc(val)
		return
	}

	/* ignore leading spaces in the variable name */

	for (*var_name) == ' ' {
		var_name++
	}

	/*
	 * Prepare variable name
	 */

	var_len = strlen(var_name)
	var_orig = zend._emalloc(var_len + 1)
	var_ = var_orig
	memcpy(var_orig, var_name, var_len+1)

	/* ensure that we don't have spaces or dots in the variable name (not binary safe) */

	for p = var_; *p; p++ {
		if (*p) == ' ' || (*p) == '.' {
			*p = '_'
		} else if (*p) == '[' {
			is_array = 1
			ip = p
			*p = 0
			break
		}
	}
	var_len = p - var_

	/* Discard variable if mangling made it start with __Host-, where pre-mangling it did not start with __Host- */

	if strncmp(var_, "__Host-", g.SizeOf("\"__Host-\"")-1) == 0 && strncmp(var_name, "__Host-", g.SizeOf("\"__Host-\"")-1) != 0 {
		zend.ZvalPtrDtorNogc(val)
		zend._efree(var_orig)
		return
	}

	/* Discard variable if mangling made it start with __Secure-, where pre-mangling it did not start with __Secure- */

	if strncmp(var_, "__Secure-", g.SizeOf("\"__Secure-\"")-1) == 0 && strncmp(var_name, "__Secure-", g.SizeOf("\"__Secure-\"")-1) != 0 {
		zend.ZvalPtrDtorNogc(val)
		zend._efree(var_orig)
		return
	}
	if var_len == 0 {
		zend.ZvalPtrDtorNogc(val)
		zend._efree(var_orig)
		return
	}
	if var_len == g.SizeOf("\"this\"")-1 && zend.EG.current_execute_data != nil {
		var ex *zend.ZendExecuteData = zend.EG.current_execute_data
		for ex != nil {
			if ex.func_ != nil && (ex.func_.common.type_&1) == 0 {
				if (ex.This.u1.type_info&1<<20) != 0 && ex.symbol_table == symtable1 {
					if memcmp(var_, "this", g.SizeOf("\"this\"")-1) == 0 {
						zend.ZendThrowError(nil, "Cannot re-assign $this")
						zend.ZvalPtrDtorNogc(val)
						zend._efree(var_orig)
						return
					}
				}
				break
			}
			ex = ex.prev_execute_data
		}
	}

	/* GLOBALS hijack attempt, reject parameter */

	if symtable1 == &zend.EG.symbol_table && var_len == g.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1)) {
		zend.ZvalPtrDtorNogc(val)
		zend._efree(var_orig)
		return
	}
	index = var_
	index_len = var_len
	if is_array != 0 {
		var nest_level int = 0
		for true {
			var index_s *byte
			var new_idx_len int = 0
			if g.PreInc(&nest_level) > CoreGlobals.GetMaxInputNestingLevel() {
				var ht *zend.HashTable

				/* too many levels of nesting */

				if track_vars_array != nil {
					ht = track_vars_array.value.arr
					zend.ZendSymtableStrDel(ht, var_, var_len)
				}
				zend.ZvalPtrDtorNogc(val)

				/* do not output the error message to the screen,
				   this helps us to to avoid "information disclosure" */

				if CoreGlobals.GetDisplayErrors() == 0 {
					PhpErrorDocref(nil, 1<<1, "Input variable nesting level exceeded "+"%"+"lld"+". To increase the limit change max_input_nesting_level in php.ini.", CoreGlobals.GetMaxInputNestingLevel())
				}
				zend._efree(var_orig)
				return
			}
			ip++
			index_s = ip
			if isspace(*ip) {
				ip++
			}
			if (*ip) == ']' {
				index_s = nil
			} else {
				ip = strchr(ip, ']')
				if ip == nil {

					/* PHP variables cannot contain '[' in their names, so we replace the character with a '_' */

					*(index_s - 1) = '_'
					index_len = 0
					if index != nil {
						index_len = strlen(index)
					}
					goto plain_var
					return
				}
				*ip = 0
				new_idx_len = strlen(index_s)
			}
			if index == nil {
				var __arr *zend.ZendArray = zend._zendNewArray(0)
				var __z *zend.Zval = &gpc_element
				__z.value.arr = __arr
				__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				if g.Assign(&gpc_element_p, zend.ZendHashNextIndexInsert(symtable1, &gpc_element)) == nil {
					zend.ZendArrayDestroy(gpc_element.value.arr)
					zend.ZvalPtrDtorNogc(val)
					zend._efree(var_orig)
					return
				}
			} else {
				gpc_element_p = zend.ZendSymtableStrFind(symtable1, index, index_len)
				if gpc_element_p == nil {
					var tmp zend.Zval
					var __arr *zend.ZendArray = zend._zendNewArray(0)
					var __z *zend.Zval = &tmp
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
					gpc_element_p = zend.ZendSymtableStrUpdateInd(symtable1, index, index_len, &tmp)
				} else {
					if gpc_element_p.u1.v.type_ == 13 {
						gpc_element_p = gpc_element_p.value.zv
					}
					if gpc_element_p.u1.v.type_ != 7 {
						zend.ZvalPtrDtorNogc(gpc_element_p)
						var __arr *zend.ZendArray = zend._zendNewArray(0)
						var __z *zend.Zval = gpc_element_p
						__z.value.arr = __arr
						__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
					} else {
						var _zv *zend.Zval = gpc_element_p
						var _arr *zend.ZendArray = _zv.value.arr
						if zend.ZendGcRefcount(&_arr.gc) > 1 {
							if _zv.u1.v.type_flags != 0 {
								zend.ZendGcDelref(&_arr.gc)
							}
							var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
							var __z *zend.Zval = _zv
							__z.value.arr = __arr
							__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
						}
					}
				}
			}
			symtable1 = gpc_element_p.value.arr

			/* ip pointed to the '[' character, now obtain the key */

			index = index_s
			index_len = new_idx_len
			ip++
			if (*ip) == '[' {
				is_array = 1
				*ip = 0
			} else {
				goto plain_var
			}
		}
	} else {
	plain_var:
		if index == nil {
			if zend.ZendHashNextIndexInsert(symtable1, val) == nil {
				zend.ZvalPtrDtorNogc(val)
			}
		} else {
			var idx zend.ZendUlong

			/*
			 * According to rfc2965, more specific paths are listed above the less specific ones.
			 * If we encounter a duplicate cookie name, we should skip it, since it is not possible
			 * to have the same (plain text) cookie name for the same path and we should not overwrite
			 * more specific cookies with the less specific ones.
			 */

			if CoreGlobals.GetHttpGlobals()[2].u1.v.type_ != 0 && symtable1 == CoreGlobals.GetHttpGlobals()[2].value.arr && zend.ZendSymtableStrExists(symtable1, index, index_len) != 0 {
				zend.ZvalPtrDtorNogc(val)
			} else if zend._zendHandleNumericStr(index, index_len, &idx) != 0 {
				zend.ZendHashIndexUpdate(symtable1, idx, val)
			} else {
				PhpRegisterVariableQuick(index, index_len, val, symtable1)
			}

			/*
			 * According to rfc2965, more specific paths are listed above the less specific ones.
			 * If we encounter a duplicate cookie name, we should skip it, since it is not possible
			 * to have the same (plain text) cookie name for the same path and we should not overwrite
			 * more specific cookies with the less specific ones.
			 */

		}
	}
	zend._efree(var_orig)
}

type PostVarData = PostVarDataT

func AddPostVar(arr *zend.Zval, var_ *PostVarDataT, eof zend.ZendBool) zend.ZendBool {
	var start *byte
	var ksep *byte
	var vsep *byte
	var val *byte
	var klen int
	var vlen int
	var new_vlen int
	if var_.GetPtr() >= var_.GetEnd() {
		return 0
	}
	start = var_.GetPtr() + var_.GetAlreadyScanned()
	vsep = memchr(start, '&', var_.GetEnd()-start)
	if vsep == nil {
		if eof == 0 {
			var_.SetAlreadyScanned(var_.GetEnd() - var_.GetPtr())
			return 0
		} else {
			vsep = var_.GetEnd()
		}
	}
	ksep = memchr(var_.GetPtr(), '=', vsep-var_.GetPtr())
	if ksep != nil {
		*ksep = '0'

		/* "foo=bar&" or "foo=&" */

		klen = ksep - var_.GetPtr()
		vlen = vsep - g.PreInc(&ksep)
	} else {
		ksep = ""

		/* "foo&" */

		klen = vsep - var_.GetPtr()
		vlen = 0
	}
	streams.PhpUrlDecode(var_.GetPtr(), klen)
	val = zend._estrndup(ksep, vlen)
	if vlen != 0 {
		vlen = streams.PhpUrlDecode(val, vlen)
	}
	if sapi_module.GetInputFilter()(0, var_.GetPtr(), &val, vlen, &new_vlen) != 0 {
		PhpRegisterVariableSafe(var_.GetPtr(), val, new_vlen, arr)
	}
	zend._efree(val)
	var_.SetPtr(vsep + (vsep != var_.GetEnd()))
	var_.SetAlreadyScanned(0)
	return 1
}
func AddPostVars(arr *zend.Zval, vars *PostVarDataT, eof zend.ZendBool) int {
	var max_vars uint64 = CoreGlobals.GetMaxInputVars()
	vars.SetPtr(vars.str.s.val)
	vars.SetEnd(vars.str.s.val + vars.str.s.len_)
	for AddPostVar(arr, vars, eof) != 0 {
		if g.PreInc(&(vars.GetCnt())) > max_vars {
			PhpErrorDocref(nil, 1<<1, "Input variables exceeded %"+"llu"+". "+"To increase the limit change max_input_vars in php.ini.", max_vars)
			return zend.FAILURE
		}
	}
	if eof == 0 && vars.str.s.val != vars.GetPtr() {
		memmove(vars.str.s.val, vars.GetPtr(), g.Assign(&(vars.str.s.len_), vars.GetEnd()-vars.GetPtr()))
	}
	return zend.SUCCESS
}

// #define SAPI_POST_HANDLER_BUFSIZ       BUFSIZ

func PhpStdPostHandler(content_type_dup *byte, arg any) {
	var arr *zend.Zval = (*zend.Zval)(arg)
	var s *PhpStream = sapi_globals.GetRequestInfo().GetRequestBody()
	var post_data PostVarDataT
	if s != nil && zend.SUCCESS == _phpStreamSeek(s, 0, 0) {
		memset(&post_data, 0, g.SizeOf("post_data"))
		for _phpStreamEof(s) == 0 {
			var buf []byte = []byte{0}
			var len_ ssize_t = _phpStreamRead(s, buf, 1024)
			if len_ > 0 {
				zend.SmartStrAppendlEx(&post_data.str, buf, len_, 0)
				if zend.SUCCESS != AddPostVars(arr, &post_data, 0) {
					zend.SmartStrFreeEx(&post_data.str, 0)
					return
				}
			}
			if len_ != 1024 {
				break
			}
		}
		if post_data.str.s != nil {
			AddPostVars(arr, &post_data, 1)
			zend.SmartStrFreeEx(&post_data.str, 0)
		}
	}
}
func PhpDefaultInputFilter(arg int, var_ *byte, val **byte, val_len int, new_val_len *int) uint {
	/* TODO: check .ini setting here and apply user-defined input filter */

	if new_val_len != nil {
		*new_val_len = val_len
	}
	return 1
}
func PhpDefaultTreatData(arg int, str *byte, destArray *zend.Zval) {
	var res *byte = nil
	var var_ *byte
	var val *byte
	var separator *byte = nil
	var c_var *byte
	var array zend.Zval
	var free_buffer int = 0
	var strtok_buf *byte = nil
	var count zend.ZendLong = 0
	&array.u1.type_info = 0
	switch arg {
	case 0:

	case 1:

	case 2:
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &array
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		switch arg {
		case 0:
			zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[0])
			var _z1 *zend.Zval = &CoreGlobals.GetHttpGlobals()[0]
			var _z2 *zend.Zval = &array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			break
		case 1:
			zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[1])
			var _z1 *zend.Zval = &CoreGlobals.GetHttpGlobals()[1]
			var _z2 *zend.Zval = &array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			break
		case 2:
			zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[2])
			var _z1 *zend.Zval = &CoreGlobals.GetHttpGlobals()[2]
			var _z2 *zend.Zval = &array
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			break
		}
		break
	default:
		var _z1 *zend.Zval = &array
		var _z2 *zend.Zval = destArray
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		break
	}
	if arg == 0 {
		SapiHandlePost(&array)
		return
	}
	if arg == 1 {
		c_var = sapi_globals.GetRequestInfo().GetQueryString()
		if c_var != nil && (*c_var) {
			res = (*byte)(zend._estrdup(c_var))
			free_buffer = 1
		} else {
			free_buffer = 0
		}
	} else if arg == 2 {
		c_var = sapi_globals.GetRequestInfo().GetCookieData()
		if c_var != nil && (*c_var) {
			res = (*byte)(zend._estrdup(c_var))
			free_buffer = 1
		} else {
			free_buffer = 0
		}
	} else if arg == 3 {
		res = str
		free_buffer = 1
	}
	if res == nil {
		return
	}
	switch arg {
	case 1:

	case 3:
		separator = CoreGlobals.GetArgSeparator().GetInput()
		break
	case 2:
		separator = ";0"
		break
	}
	var_ = strtok_r(res, separator, &strtok_buf)
	for var_ != nil {
		var val_len int
		var new_val_len int
		val = strchr(var_, '=')
		if arg == 2 {

			/* Remove leading spaces from cookie names, needed for multi-cookie header where ; can be followed by a space */

			for isspace(*var_) {
				var_++
			}
			if var_ == val || (*var_) == '0' {
				goto next_cookie
			}
		}
		if g.PreInc(&count) > CoreGlobals.GetMaxInputVars() {
			PhpErrorDocref(nil, 1<<1, "Input variables exceeded "+"%"+"lld"+". To increase the limit change max_input_vars in php.ini.", CoreGlobals.GetMaxInputVars())
			break
		}
		if val != nil {
			g.PostInc(&(*val)) = '0'
			if arg == 2 {
				val_len = standard.PhpRawUrlDecode(val, strlen(val))
			} else {
				val_len = streams.PhpUrlDecode(val, strlen(val))
			}
		} else {
			val = ""
			val_len = 0
		}
		val = zend._estrndup(val, val_len)
		if arg != 2 {
			streams.PhpUrlDecode(var_, strlen(var_))
		}
		if sapi_module.GetInputFilter()(arg, var_, &val, val_len, &new_val_len) != 0 {
			PhpRegisterVariableSafe(var_, val, new_val_len, &array)
		}
		zend._efree(val)
	next_cookie:
		var_ = strtok_r(nil, separator, &strtok_buf)
	}
	if free_buffer != 0 {
		zend._efree(res)
	}
}
func ValidEnvironmentName(name *byte, end *byte) int {
	var s *byte
	for s = name; s < end; s++ {
		if (*s) == ' ' || (*s) == '.' || (*s) == '[' {
			return 0
		}
	}
	return 1
}
func ImportEnvironmentVariable(ht *zend.HashTable, env *byte) {
	var p *byte
	var name_len int
	var len_ int
	var val zend.Zval
	var idx zend.ZendUlong
	p = strchr(env, '=')
	if p == nil || p == env || ValidEnvironmentName(env, p) == 0 {

		/* malformed entry? */

		return

		/* malformed entry? */

	}
	name_len = p - env
	p++
	len_ = strlen(p)
	if len_ == 0 {
		var __z *zend.Zval = &val
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
	} else if len_ == 1 {
		var __z *zend.Zval = &val
		var __s *zend.ZendString = zend.ZendOneCharString[zend_uchar*p]
		__z.value.str = __s
		__z.u1.type_info = 6
	} else {
		var __z *zend.Zval = &val
		var __s *zend.ZendString = zend.ZendStringInit(p, len_, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
	if zend._zendHandleNumericStr(env, name_len, &idx) != 0 {
		zend.ZendHashIndexUpdate(ht, idx, &val)
	} else {
		PhpRegisterVariableQuick(env, name_len, &val, ht)
	}
}
func _phpImportEnvironmentVariables(array_ptr *zend.Zval) {
	var env **byte
	tsrm_env_lock()
	for env = Environ; env != nil && (*env) != nil; env++ {
		ImportEnvironmentVariable(array_ptr.value.arr, *env)
	}
	tsrm_env_unlock()
}
func PhpStdAutoGlobalCallback(name *byte, name_len uint32) zend.ZendBool {
	zend.ZendPrintf("%s\n", name)
	return 0
}

/* {{{ php_build_argv
 */

func PhpBuildArgv(s *byte, track_vars_array *zend.Zval) {
	var arr zend.Zval
	var argc zend.Zval
	var tmp zend.Zval
	var count int = 0
	var ss *byte
	var space *byte
	if !(sapi_globals.GetRequestInfo().GetArgc() != 0 || track_vars_array != nil) {
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &arr
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* Prepare argv */

	if sapi_globals.GetRequestInfo().GetArgc() != 0 {
		var i int
		for i = 0; i < sapi_globals.GetRequestInfo().GetArgc(); i++ {
			var _s *byte = sapi_globals.GetRequestInfo().GetArgv()[i]
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			if zend.ZendHashNextIndexInsert(arr.value.arr, &tmp) == nil {
				zend.ZendStringEfree(tmp.value.str)
			}
		}
	} else if s != nil && (*s) {
		ss = s
		for ss != nil {
			space = strchr(ss, '+')
			if space != nil {
				*space = '0'
			}

			/* auto-type */

			var _s *byte = ss
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			count++
			if zend.ZendHashNextIndexInsert(arr.value.arr, &tmp) == nil {
				zend.ZendStringEfree(tmp.value.str)
			}
			if space != nil {
				*space = '+'
				ss = space + 1
			} else {
				ss = space
			}
		}
	}

	/* prepare argc */

	if sapi_globals.GetRequestInfo().GetArgc() != 0 {
		var __z *zend.Zval = &argc
		__z.value.lval = sapi_globals.GetRequestInfo().GetArgc()
		__z.u1.type_info = 4
	} else {
		var __z *zend.Zval = &argc
		__z.value.lval = count
		__z.u1.type_info = 4
	}
	if sapi_globals.GetRequestInfo().GetArgc() != 0 {
		zend.ZvalAddrefP(&arr)
		zend.ZendHashUpdate(&zend.EG.symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], &arr)
		zend.ZendHashUpdate(&zend.EG.symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_ARGC], &argc)
	}
	if track_vars_array != nil && track_vars_array.u1.v.type_ == 7 {
		zend.ZvalAddrefP(&arr)
		zend.ZendHashUpdate(track_vars_array.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], &arr)
		zend.ZendHashUpdate(track_vars_array.value.arr, zend.ZendKnownStrings[zend.ZEND_STR_ARGC], &argc)
	}
	zend.ZvalPtrDtorNogc(&arr)
}

/* }}} */

func PhpRegisterServerVariables() {
	var tmp zend.Zval
	var arr *zend.Zval = &CoreGlobals.GetHttpGlobals()[3]
	var ht *zend.HashTable
	zend.ZvalPtrDtorNogc(arr)
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = arr
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8

	/* Server variables */

	if sapi_module.GetRegisterServerVariables() != nil {
		sapi_module.GetRegisterServerVariables()(arr)
	}
	ht = arr.value.arr

	/* PHP Authentication support */

	if sapi_globals.GetRequestInfo().GetAuthUser() != nil {
		var _s *byte = sapi_globals.GetRequestInfo().GetAuthUser()
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		PhpRegisterVariableQuick("PHP_AUTH_USER", g.SizeOf("\"PHP_AUTH_USER\"")-1, &tmp, ht)
	}
	if sapi_globals.GetRequestInfo().GetAuthPassword() != nil {
		var _s *byte = sapi_globals.GetRequestInfo().GetAuthPassword()
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		PhpRegisterVariableQuick("PHP_AUTH_PW", g.SizeOf("\"PHP_AUTH_PW\"")-1, &tmp, ht)
	}
	if sapi_globals.GetRequestInfo().GetAuthDigest() != nil {
		var _s *byte = sapi_globals.GetRequestInfo().GetAuthDigest()
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		PhpRegisterVariableQuick("PHP_AUTH_DIGEST", g.SizeOf("\"PHP_AUTH_DIGEST\"")-1, &tmp, ht)
	}

	/* store request init time */

	var __z *zval = &tmp
	__z.value.dval = SapiGetRequestTime()
	__z.u1.type_info = 5
	PhpRegisterVariableQuick("REQUEST_TIME_FLOAT", g.SizeOf("\"REQUEST_TIME_FLOAT\"")-1, &tmp, ht)
	var __z *zend.Zval = &tmp
	__z.value.lval = zend.ZendDvalToLval(tmp.value.dval)
	__z.u1.type_info = 4
	PhpRegisterVariableQuick("REQUEST_TIME", g.SizeOf("\"REQUEST_TIME\"")-1, &tmp, ht)
}

/* }}} */

func PhpAutoglobalMerge(dest *zend.HashTable, src *zend.HashTable) {
	var src_entry *zend.Zval
	var dest_entry *zend.Zval
	var string_key *zend.ZendString
	var num_key zend.ZendUlong
	var globals_check int = dest == &zend.EG.symbol_table
	for {
		var __ht *zend.HashTable = src
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			num_key = _p.h
			string_key = _p.key
			src_entry = _z
			if src_entry.u1.v.type_ != 7 || string_key != nil && g.Assign(&dest_entry, zend.ZendHashFind(dest, string_key)) == nil || string_key == nil && g.Assign(&dest_entry, zend.ZendHashIndexFind(dest, num_key)) == nil || dest_entry.u1.v.type_ != 7 {
				if src_entry.u1.v.type_flags != 0 {
					zend.ZvalAddrefP(src_entry)
				}
				if string_key != nil {
					if globals_check == 0 || string_key.len_ != g.SizeOf("\"GLOBALS\"")-1 || memcmp(string_key.val, "GLOBALS", g.SizeOf("\"GLOBALS\"")-1) {
						zend.ZendHashUpdate(dest, string_key, src_entry)
					} else {
						if src_entry.u1.v.type_flags != 0 {
							zend.ZvalDelrefP(src_entry)
						}
					}
				} else {
					zend.ZendHashIndexUpdate(dest, num_key, src_entry)
				}
			} else {
				var _zv *zend.Zval = dest_entry
				var _arr *zend.ZendArray = _zv.value.arr
				if zend.ZendGcRefcount(&_arr.gc) > 1 {
					if _zv.u1.v.type_flags != 0 {
						zend.ZendGcDelref(&_arr.gc)
					}
					var __arr *zend.ZendArray = zend.ZendArrayDup(_arr)
					var __z *zend.Zval = _zv
					__z.value.arr = __arr
					__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
				}
				PhpAutoglobalMerge(dest_entry.value.arr, src_entry.value.arr)
			}
		}
		break
	}
}

/* }}} */

func PhpHashEnvironment() int {
	memset(CoreGlobals.GetHttpGlobals(), 0, g.SizeOf("PG ( http_globals )"))
	zend.ZendActivateAutoGlobals()
	if CoreGlobals.GetRegisterArgcArgv() != 0 {
		PhpBuildArgv(sapi_globals.GetRequestInfo().GetQueryString(), &CoreGlobals.GetHttpGlobals()[3])
	}
	return zend.SUCCESS
}

/* }}} */

func PhpAutoGlobalsCreateGet(name *zend.ZendString) zend.ZendBool {
	if CoreGlobals.GetVariablesOrder() != nil && (strchr(CoreGlobals.GetVariablesOrder(), 'G') || strchr(CoreGlobals.GetVariablesOrder(), 'g')) {
		sapi_module.GetTreatData()(1, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[1])
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[1]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[1])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[1])
	return 0
}
func PhpAutoGlobalsCreatePost(name *zend.ZendString) zend.ZendBool {
	if CoreGlobals.GetVariablesOrder() != nil && (strchr(CoreGlobals.GetVariablesOrder(), 'P') || strchr(CoreGlobals.GetVariablesOrder(), 'p')) && sapi_globals.GetHeadersSent() == 0 && sapi_globals.GetRequestInfo().GetRequestMethod() != nil && !(strcasecmp(sapi_globals.GetRequestInfo().GetRequestMethod(), "POST")) {
		sapi_module.GetTreatData()(0, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[0])
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[0]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[0])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[0])
	return 0
}
func PhpAutoGlobalsCreateCookie(name *zend.ZendString) zend.ZendBool {
	if CoreGlobals.GetVariablesOrder() != nil && (strchr(CoreGlobals.GetVariablesOrder(), 'C') || strchr(CoreGlobals.GetVariablesOrder(), 'c')) {
		sapi_module.GetTreatData()(2, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[2])
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[2]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[2])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[2])
	return 0
}
func PhpAutoGlobalsCreateFiles(name *zend.ZendString) zend.ZendBool {
	if CoreGlobals.GetHttpGlobals()[5].u1.v.type_ == 0 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[5]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[5])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[5])
	return 0
}

/* Upgly hack to fix HTTP_PROXY issue, see bug #72573 */

func CheckHttpProxy(var_table *zend.HashTable) {
	if zend.ZendHashStrExists(var_table, "HTTP_PROXY", g.SizeOf("\"HTTP_PROXY\"")-1) != 0 {
		var local_proxy *byte = getenv("HTTP_PROXY")
		if local_proxy == nil {
			zend.ZendHashStrDel(var_table, "HTTP_PROXY", g.SizeOf("\"HTTP_PROXY\"")-1)
		} else {
			var local_zval zend.Zval
			var _s *byte = local_proxy
			var __z *zend.Zval = &local_zval
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendHashStrUpdate(var_table, "HTTP_PROXY", g.SizeOf("\"HTTP_PROXY\"")-1, &local_zval)
		}
	}
}
func PhpAutoGlobalsCreateServer(name *zend.ZendString) zend.ZendBool {
	if CoreGlobals.GetVariablesOrder() != nil && (strchr(CoreGlobals.GetVariablesOrder(), 'S') || strchr(CoreGlobals.GetVariablesOrder(), 's')) {
		PhpRegisterServerVariables()
		if CoreGlobals.GetRegisterArgcArgv() != 0 {
			if sapi_globals.GetRequestInfo().GetArgc() != 0 {
				var argc *zend.Zval
				var argv *zend.Zval
				if g.Assign(&argc, zend.ZendHashFindExInd(&zend.EG.symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_ARGC], 1)) != nil && g.Assign(&argv, zend.ZendHashFindExInd(&zend.EG.symbol_table, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], 1)) != nil {
					zend.ZvalAddrefP(argv)
					zend.ZendHashUpdate(CoreGlobals.GetHttpGlobals()[3].value.arr, zend.ZendKnownStrings[zend.ZEND_STR_ARGV], argv)
					zend.ZendHashUpdate(CoreGlobals.GetHttpGlobals()[3].value.arr, zend.ZendKnownStrings[zend.ZEND_STR_ARGC], argc)
				}
			} else {
				PhpBuildArgv(sapi_globals.GetRequestInfo().GetQueryString(), &CoreGlobals.GetHttpGlobals()[3])
			}
		}
	} else {
		zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[3])
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[3]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}
	CheckHttpProxy(CoreGlobals.GetHttpGlobals()[3].value.arr)
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[3])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[3])

	/* TODO: TRACK_VARS_SERVER is modified in a number of places (e.g. phar) past this point,
	 * where rc>1 due to the $_SERVER global. Ideally this shouldn't happen, but for now we
	 * ignore this issue, as it would probably require larger changes. */

	return 0
}
func PhpAutoGlobalsCreateEnv(name *zend.ZendString) zend.ZendBool {
	zend.ZvalPtrDtorNogc(&CoreGlobals.GetHttpGlobals()[4])
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &CoreGlobals.GetHttpGlobals()[4]
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if CoreGlobals.GetVariablesOrder() != nil && (strchr(CoreGlobals.GetVariablesOrder(), 'E') || strchr(CoreGlobals.GetVariablesOrder(), 'e')) {
		PhpImportEnvironmentVariables(&CoreGlobals.GetHttpGlobals()[4])
	}
	CheckHttpProxy(CoreGlobals.GetHttpGlobals()[4].value.arr)
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &CoreGlobals.GetHttpGlobals()[4])
	zend.ZvalAddrefP(&CoreGlobals.GetHttpGlobals()[4])
	return 0
}
func PhpAutoGlobalsCreateRequest(name *zend.ZendString) zend.ZendBool {
	var form_variables zend.Zval
	var _gpc_flags []uint8 = []uint8{0, 0, 0}
	var p *byte
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &form_variables
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	if CoreGlobals.GetRequestOrder() != nil {
		p = CoreGlobals.GetRequestOrder()
	} else {
		p = CoreGlobals.GetVariablesOrder()
	}
	for ; p != nil && (*p); p++ {
		switch *p {
		case 'g':

		case 'G':
			if _gpc_flags[0] == 0 {
				PhpAutoglobalMerge(form_variables.value.arr, CoreGlobals.GetHttpGlobals()[1].value.arr)
				_gpc_flags[0] = 1
			}
			break
		case 'p':

		case 'P':
			if _gpc_flags[1] == 0 {
				PhpAutoglobalMerge(form_variables.value.arr, CoreGlobals.GetHttpGlobals()[0].value.arr)
				_gpc_flags[1] = 1
			}
			break
		case 'c':

		case 'C':
			if _gpc_flags[2] == 0 {
				PhpAutoglobalMerge(form_variables.value.arr, CoreGlobals.GetHttpGlobals()[2].value.arr)
				_gpc_flags[2] = 1
			}
			break
		}
	}
	zend.ZendHashUpdate(&zend.EG.symbol_table, name, &form_variables)
	return 0
}
func PhpStartupAutoGlobals() {
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_GET", g.SizeOf("\"_GET\"")-1, 1), 0, PhpAutoGlobalsCreateGet)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_POST", g.SizeOf("\"_POST\"")-1, 1), 0, PhpAutoGlobalsCreatePost)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_COOKIE", g.SizeOf("\"_COOKIE\"")-1, 1), 0, PhpAutoGlobalsCreateCookie)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_SERVER", g.SizeOf("\"_SERVER\"")-1, 1), CoreGlobals.GetAutoGlobalsJit(), PhpAutoGlobalsCreateServer)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_ENV", g.SizeOf("\"_ENV\"")-1, 1), CoreGlobals.GetAutoGlobalsJit(), PhpAutoGlobalsCreateEnv)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_REQUEST", g.SizeOf("\"_REQUEST\"")-1, 1), CoreGlobals.GetAutoGlobalsJit(), PhpAutoGlobalsCreateRequest)
	zend.ZendRegisterAutoGlobal(zend.ZendStringInitInterned("_FILES", g.SizeOf("\"_FILES\"")-1, 1), 0, PhpAutoGlobalsCreateFiles)
}
