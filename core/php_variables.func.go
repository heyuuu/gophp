package core

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpRegisterVariable(var_ string, strval *byte, track_vars_array *types.Zval) {
	PhpRegisterVariableSafe(var_, strval, strlen(strval), track_vars_array)
}
func PhpRegisterVariableSafe(var_ *byte, strval *byte, str_len int, track_vars_array *types.Zval) {
	var new_entry types.Zval
	b.Assert(strval != nil)

	/* Prepare value */

	if str_len == 0 {
		zend.ZVAL_EMPTY_STRING(&new_entry)
	} else if str_len == 1 {
		new_entry.SetInternedString(types.ZSTR_CHAR(zend_uchar * strval))
	} else {
		new_entry.SetString(types.NewString(b.CastStr(strval, str_len)))
	}
	PhpRegisterVariableEx(var_, &new_entry, track_vars_array)
}
func PhpRegisterVariableQuick(name *byte, name_len int, val *types.Zval, ht *types.Array) {
	var key *types.String = types.ZendStringInitInterned(name, name_len, 0)
	ht.KeyUpdateIndirect(key.GetStr(), val)
	types.ZendStringReleaseEx(key, 0)
}
func PhpRegisterVariableEx(var_name *byte, val *types.Zval, track_vars_array *types.Zval) {
	var p *byte = nil
	var ip *byte = nil
	var index *byte
	var var_ *byte
	var var_orig *byte
	var var_len int
	var index_len int
	var gpc_element types.Zval
	var gpc_element_p *types.Zval
	var is_array types.ZendBool = 0
	var symtable1 *types.Array = nil
	b.Assert(var_name != nil)
	if track_vars_array != nil && track_vars_array.IsType(types.IS_ARRAY) {
		symtable1 = track_vars_array.GetArr()
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
	var_orig = zend.DoAlloca(var_len+1, use_heap)
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

	if strncmp(var_, "__Host-", b.SizeOf("\"__Host-\"")-1) == 0 && strncmp(var_name, "__Host-", b.SizeOf("\"__Host-\"")-1) != 0 {
		zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(var_orig, use_heap)
		return
	}

	/* Discard variable if mangling made it start with __Secure-, where pre-mangling it did not start with __Secure- */

	if strncmp(var_, "__Secure-", b.SizeOf("\"__Secure-\"")-1) == 0 && strncmp(var_name, "__Secure-", b.SizeOf("\"__Secure-\"")-1) != 0 {
		zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(var_orig, use_heap)
		return
	}
	if var_len == 0 {
		zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(var_orig, use_heap)
		return
	}
	if var_len == b.SizeOf("\"this\"")-1 && zend.CurrEX() != nil {
		var ex *zend.ZendExecuteData = zend.CurrEX()
		for ex != nil {
			if ex.GetFunc() != nil && zend.ZEND_USER_CODE(ex.GetFunc().GetCommonType()) {
				if (zend.ZEND_CALL_INFO(ex)&zend.ZEND_CALL_HAS_SYMBOL_TABLE) != 0 && ex.GetSymbolTable() == symtable1 {
					if memcmp(var_, "this", b.SizeOf("\"this\"")-1) == 0 {
						faults.ThrowError(nil, "Cannot re-assign $this")
						zend.ZvalPtrDtorNogc(val)
						zend.FreeAlloca(var_orig, use_heap)
						return
					}
				}
				break
			}
			ex = ex.GetPrevExecuteData()
		}
	}

	/* GLOBALS hijack attempt, reject parameter */

	if symtable1 == zend.EG__().GetSymbolTable() && var_len == b.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_, "GLOBALS", b.SizeOf("\"GLOBALS\"")-1)) {
		zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(var_orig, use_heap)
		return
	}
	index = var_
	index_len = var_len
	if is_array != 0 {
		var nest_level int = 0
		for true {
			var index_s *byte
			var new_idx_len int = 0
			if b.PreInc(&nest_level) > PG__().max_input_nesting_level {
				var ht *types.Array

				/* too many levels of nesting */

				if track_vars_array != nil {
					ht = track_vars_array.GetArr()
					ht.SymtableDel(b.CastStr(var_, var_len))
				}
				zend.ZvalPtrDtorNogc(val)

				/* do not output the error message to the screen,
				   this helps us to to avoid "information disclosure" */

				if !(PG__().display_errors) {
					PhpErrorDocref(nil, faults.E_WARNING, "Input variable nesting level exceeded "+zend.ZEND_LONG_FMT+". To increase the limit change max_input_nesting_level in php.ini.", PG__().max_input_nesting_level)
				}
				zend.FreeAlloca(var_orig, use_heap)
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
				zend.ArrayInit(&gpc_element)
				if b.Assign(&gpc_element_p, symtable1.NextIndexInsert(&gpc_element)) == nil {
					gpc_element.GetArr().DestroyEx()
					zend.ZvalPtrDtorNogc(val)
					zend.FreeAlloca(var_orig, use_heap)
					return
				}
			} else {
				gpc_element_p = symtable1.SymtableFind(b.CastStr(index, index_len))
				if gpc_element_p == nil {
					var tmp types.Zval
					zend.ArrayInit(&tmp)
					gpc_element_p = symtable1.SymtableUpdateInd(b.CastStr(index, index_len), &tmp)
				} else {
					if gpc_element_p.IsIndirect() {
						gpc_element_p = gpc_element_p.GetZv()
					}
					if gpc_element_p.GetType() != types.IS_ARRAY {
						zend.ZvalPtrDtorNogc(gpc_element_p)
						zend.ArrayInit(gpc_element_p)
					} else {
						types.SEPARATE_ARRAY(gpc_element_p)
					}
				}
			}
			symtable1 = gpc_element_p.GetArr()

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
			if symtable1.NextIndexInsert(val) == nil {
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

			if PG__().http_globals[TRACK_VARS_COOKIE].IsNotUndef() && symtable1 == PG__().http_globals[TRACK_VARS_COOKIE].GetArr() && symtable1.SymtableExists(b.CastStr(index, index_len)) {
				zend.ZvalPtrDtorNogc(val)
			} else if types.HandleNumericStr(b.CastStr(index, index_len), &idx) {
				symtable1.IndexUpdate(idx, val)
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
	zend.FreeAlloca(var_orig, use_heap)
}
func AddPostVar(arr *types.Zval, var_ *PostVarDataT, eof types.ZendBool) types.ZendBool {
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
		vlen = vsep - b.PreInc(&ksep)
	} else {
		ksep = ""

		/* "foo&" */

		klen = vsep - var_.GetPtr()
		vlen = 0
	}
	streams.PhpUrlDecode(var_.GetPtr(), klen)
	val = zend.Estrndup(ksep, vlen)
	if vlen != 0 {
		vlen = streams.PhpUrlDecode(val, vlen)
	}
	if SM__().GetInputFilter()(PARSE_POST, var_.GetPtr(), &val, vlen, &new_vlen) != 0 {
		PhpRegisterVariableSafe(var_.GetPtr(), val, new_vlen, arr)
	}
	zend.Efree(val)
	var_.SetPtr(vsep + (vsep != var_.GetEnd()))
	var_.SetAlreadyScanned(0)
	return 1
}
func AddPostVars(arr *types.Zval, vars *PostVarDataT, eof types.ZendBool) int {
	var max_vars uint64 = PG__().max_input_vars
	vars.SetPtr(vars.GetStr().GetS().GetVal())
	vars.SetEnd(vars.GetStr().GetS().GetVal() + vars.GetStr().GetS().GetLen())
	for AddPostVar(arr, vars, eof) != 0 {
		if b.PreInc(&(vars.GetCnt())) > max_vars {
			PhpErrorDocref(nil, faults.E_WARNING, "Input variables exceeded %"+"llu"+". "+"To increase the limit change max_input_vars in php.ini.", max_vars)
			return types.FAILURE
		}
	}
	if eof == 0 && vars.GetStr().GetS().GetVal() != vars.GetPtr() {
		memmove(vars.GetStr().GetS().GetVal(), vars.GetPtr(), b.Assign(&(vars.GetStr().GetS().GetLen()), vars.GetEnd()-vars.GetPtr()))
	}
	return types.SUCCESS
}
func PhpStdPostHandler(content_type_dup *byte, arg any) {
	var arr *types.Zval = (*types.Zval)(arg)
	var s *PhpStream = SG__().request_info.request_body
	var post_data PostVarDataT
	if s != nil && types.SUCCESS == PhpStreamRewind(s) {
		memset(&post_data, 0, b.SizeOf("post_data"))
		for PhpStreamEof(s) == 0 {
			var buf []byte = []byte{0}
			var len_ ssize_t = PhpStreamRead(s, buf, r.BUFSIZ)
			if len_ > 0 {
				post_data.GetStr().AppendString(b.CastStr(buf, len_))
				if types.SUCCESS != AddPostVars(arr, &post_data, 0) {
					post_data.GetStr().Free()
					return
				}
			}
			if len_ != r.BUFSIZ {
				break
			}
		}
		if post_data.GetStr().GetS() != nil {
			AddPostVars(arr, &post_data, 1)
			post_data.GetStr().Free()
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
func PhpDefaultTreatData(arg int, str *byte, destArray *types.Zval) {
	var res *byte = nil
	var var_ *byte
	var val *byte
	var separator *byte = nil
	var c_var *byte
	var array types.Zval
	var free_buffer int = 0
	var strtok_buf *byte = nil
	var count zend.ZendLong = 0
	array.SetUndef()
	switch arg {
	case PARSE_POST:
		fallthrough
	case PARSE_GET:
		fallthrough
	case PARSE_COOKIE:
		zend.ArrayInit(&array)
		switch arg {
		case PARSE_POST:
			zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_POST])
			types.ZVAL_COPY_VALUE(&PG__().http_globals[TRACK_VARS_POST], &array)
		case PARSE_GET:
			zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_GET])
			types.ZVAL_COPY_VALUE(&PG__().http_globals[TRACK_VARS_GET], &array)
		case PARSE_COOKIE:
			zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_COOKIE])
			types.ZVAL_COPY_VALUE(&PG__().http_globals[TRACK_VARS_COOKIE], &array)
		}
	default:
		types.ZVAL_COPY_VALUE(&array, destArray)
	}
	if arg == PARSE_POST {
		SapiHandlePost(&array)
		return
	}
	if arg == PARSE_GET {
		c_var = SG__().request_info.query_string
		if c_var != nil && (*c_var) {
			res = (*byte)(zend.Estrdup(c_var))
			free_buffer = 1
		} else {
			free_buffer = 0
		}
	} else if arg == PARSE_COOKIE {
		c_var = SG__().request_info.cookie_data
		if c_var != nil && (*c_var) {
			res = (*byte)(zend.Estrdup(c_var))
			free_buffer = 1
		} else {
			free_buffer = 0
		}
	} else if arg == PARSE_STRING {
		res = str
		free_buffer = 1
	}
	if res == nil {
		return
	}
	switch arg {
	case PARSE_GET:
		fallthrough
	case PARSE_STRING:
		separator = PG__().arg_separator.input
	case PARSE_COOKIE:
		separator = ";0"
	}
	var_ = PhpStrtokR(res, separator, &strtok_buf)
	for var_ != nil {
		var val_len int
		var new_val_len int
		val = strchr(var_, '=')
		if arg == PARSE_COOKIE {

			/* Remove leading spaces from cookie names, needed for multi-cookie header where ; can be followed by a space */

			for isspace(*var_) {
				var_++
			}
			if var_ == val || (*var_) == '0' {
				goto next_cookie
			}
		}
		if b.PreInc(&count) > PG__().max_input_vars {
			PhpErrorDocref(nil, faults.E_WARNING, "Input variables exceeded "+zend.ZEND_LONG_FMT+". To increase the limit change max_input_vars in php.ini.", PG__().max_input_vars)
			break
		}
		if val != nil {
			b.PostInc(&(*val)) = '0'
			if arg == PARSE_COOKIE {
				val_len = standard.PhpRawUrlDecode(val, strlen(val))
			} else {
				val_len = streams.PhpUrlDecode(val, strlen(val))
			}
		} else {
			val = ""
			val_len = 0
		}
		val = zend.Estrndup(val, val_len)
		if arg != PARSE_COOKIE {
			streams.PhpUrlDecode(var_, strlen(var_))
		}
		if SM__().GetInputFilter()(arg, var_, &val, val_len, &new_val_len) != 0 {
			PhpRegisterVariableSafe(var_, val, new_val_len, &array)
		}
		zend.Efree(val)
	next_cookie:
		var_ = PhpStrtokR(nil, separator, &strtok_buf)
	}
	if free_buffer != 0 {
		zend.Efree(res)
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
func ImportEnvironmentVariable(ht *types.Array, env *byte) {
	var p *byte
	var name_len int
	var len_ int
	var val types.Zval
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
		zend.ZVAL_EMPTY_STRING(&val)
	} else if len_ == 1 {
		val.SetInternedString(types.ZSTR_CHAR(zend_uchar * p))
	} else {
		val.SetString(types.NewString(b.CastStr(p, len_)))
	}
	if types.HandleNumericStr(b.CastStr(env, name_len), &idx) {
		ht.IndexUpdate(idx, &val)
	} else {
		PhpRegisterVariableQuick(env, name_len, &val, ht)
	}
}
func _phpImportEnvironmentVariables(array_ptr *types.Zval) {
	var env **byte
	tsrm_env_lock()
	for env = Environ; env != nil && (*env) != nil; env++ {
		ImportEnvironmentVariable(array_ptr.GetArr(), *env)
	}
	tsrm_env_unlock()
}
func PhpStdAutoGlobalCallback(name *byte, name_len uint32) types.ZendBool {
	zend.ZendPrintf("%s\n", name)
	return 0
}
func PhpBuildArgv(s *byte, track_vars_array *types.Zval) {
	var arr types.Zval
	var argc types.Zval
	var tmp types.Zval
	var count int = 0
	var ss *byte
	var space *byte
	if !(SG__().request_info.argc || track_vars_array != nil) {
		return
	}
	zend.ArrayInit(&arr)

	/* Prepare argv */

	if SG__().request_info.argc {
		var i int
		for i = 0; i < SG__().request_info.argc; i++ {
			tmp.SetStringVal(b.CastStrAuto(SG__().request_info.argv[i]))
			if arr.GetArr().NextIndexInsert(&tmp) == nil {
				types.ZendStringEfree(tmp.GetStr())
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

			tmp.SetStringVal(b.CastStrAuto(ss))
			count++
			if arr.GetArr().NextIndexInsert(&tmp) == nil {
				types.ZendStringEfree(tmp.GetStr())
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

	if SG__().request_info.argc {
		argc.SetLong(SG__().request_info.argc)
	} else {
		argc.SetLong(count)
	}
	if SG__().request_info.argc {
		arr.AddRefcount()
		zend.EG__().GetSymbolTable().KeyUpdate(types.ZSTR_ARGV.GetStr(), &arr)
		zend.EG__().GetSymbolTable().KeyUpdate(types.ZSTR_ARGC.GetStr(), &argc)
	}
	if track_vars_array != nil && track_vars_array.IsType(types.IS_ARRAY) {
		arr.AddRefcount()
		track_vars_array.GetArr().KeyUpdate(types.ZSTR_ARGV.GetStr(), &arr)
		track_vars_array.GetArr().KeyUpdate(types.ZSTR_ARGC.GetStr(), &argc)
	}
	zend.ZvalPtrDtorNogc(&arr)
}
func PhpRegisterServerVariables() {
	var tmp types.Zval
	var arr *types.Zval = &PG__().http_globals[TRACK_VARS_SERVER]
	var ht *types.Array
	zend.ZvalPtrDtorNogc(arr)
	zend.ArrayInit(arr)

	/* Server variables */

	if SM__().GetRegisterServerVariables() != nil {
		SM__().GetRegisterServerVariables()(arr)
	}
	ht = arr.GetArr()

	/* PHP Authentication support */

	if SG__().request_info.auth_user {
		tmp.SetStringVal(b.CastStrAuto(SG__().request_info.auth_user))
		PhpRegisterVariableQuick("PHP_AUTH_USER", b.SizeOf("\"PHP_AUTH_USER\"")-1, &tmp, ht)
	}
	if SG__().request_info.auth_password {
		tmp.SetStringVal(b.CastStrAuto(SG__().request_info.auth_password))
		PhpRegisterVariableQuick("PHP_AUTH_PW", b.SizeOf("\"PHP_AUTH_PW\"")-1, &tmp, ht)
	}
	if SG__().request_info.auth_digest {
		tmp.SetStringVal(b.CastStrAuto(SG__().request_info.auth_digest))
		PhpRegisterVariableQuick("PHP_AUTH_DIGEST", b.SizeOf("\"PHP_AUTH_DIGEST\"")-1, &tmp, ht)
	}

	/* store request init time */

	tmp.SetDouble(SapiGetRequestTime())
	PhpRegisterVariableQuick("REQUEST_TIME_FLOAT", b.SizeOf("\"REQUEST_TIME_FLOAT\"")-1, &tmp, ht)
	tmp.SetLong(zend.DvalToLval(tmp.GetDval()))
	PhpRegisterVariableQuick("REQUEST_TIME", b.SizeOf("\"REQUEST_TIME\"")-1, &tmp, ht)
}
func PhpAutoglobalMerge(dest *types.Array, src *types.Array) {
	var src_entry *types.Zval
	var dest_entry *types.Zval
	var string_key *types.String
	var num_key zend.ZendUlong
	var globals_check int = dest == zend.EG__().GetSymbolTable()
	var __ht *types.Array = src
	for _, _p := range __ht.ForeachData() {
		var _z *types.Zval = _p.GetVal()

		num_key = _p.GetH()
		string_key = _p.GetKey()
		src_entry = _z
		if src_entry.GetType() != types.IS_ARRAY || string_key != nil && b.Assign(&dest_entry, dest.KeyFind(string_key.GetStr())) == nil || string_key == nil && b.Assign(&dest_entry, dest.IndexFind(num_key)) == nil || dest_entry.GetType() != types.IS_ARRAY {
			src_entry.TryAddRefcount()
			if string_key != nil {
				if globals_check == 0 || string_key.GetLen() != b.SizeOf("\"GLOBALS\"")-1 || memcmp(string_key.GetVal(), "GLOBALS", b.SizeOf("\"GLOBALS\"")-1) {
					dest.KeyUpdate(string_key.GetStr(), src_entry)
				} else {
					src_entry.TryDelRefcount()
				}
			} else {
				dest.IndexUpdate(num_key, src_entry)
			}
		} else {
			types.SEPARATE_ARRAY(dest_entry)
			PhpAutoglobalMerge(dest_entry.GetArr(), src_entry.GetArr())
		}
	}
}
func PhpHashEnvironment() int {
	memset(PG__().http_globals, 0, b.SizeOf("PG ( http_globals )"))
	zend.ZendActivateAutoGlobals()
	if PG__().register_argc_argv {
		PhpBuildArgv(SG__().request_info.query_string, &PG__().http_globals[TRACK_VARS_SERVER])
	}
	return types.SUCCESS
}
func PhpAutoGlobalsCreateGet(name *types.String) types.ZendBool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'G') || strchr(PG__().variables_order, 'g')) {
		SM__().GetTreatData()(PARSE_GET, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_GET])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_GET])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_GET])
	PG__().http_globals[TRACK_VARS_GET].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreatePost(name *types.String) types.ZendBool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'P') || strchr(PG__().variables_order, 'p')) && !(SG__().headers_sent) && SG__().request_info.request_method && !(strcasecmp(SG__().request_info.request_method, "POST")) {
		SM__().GetTreatData()(PARSE_POST, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_POST])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_POST])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_POST])
	PG__().http_globals[TRACK_VARS_POST].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreateCookie(name *types.String) types.ZendBool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'C') || strchr(PG__().variables_order, 'c')) {
		SM__().GetTreatData()(PARSE_COOKIE, nil, nil)
	} else {
		zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_COOKIE])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_COOKIE])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_COOKIE])
	PG__().http_globals[TRACK_VARS_COOKIE].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreateFiles(name *types.String) types.ZendBool {
	if PG__().http_globals[TRACK_VARS_FILES].IsUndef() {
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_FILES])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_FILES])
	PG__().http_globals[TRACK_VARS_FILES].AddRefcount()
	return 0
}
func CheckHttpProxy(var_table *types.Array) {
	if var_table.KeyExists("HTTP_PROXY") {
		var local_proxy *byte = getenv("HTTP_PROXY")
		if local_proxy == nil {
			types.ZendHashStrDel(var_table, "HTTP_PROXY")
		} else {
			var local_zval types.Zval
			local_zval.SetStringVal(b.CastStrAuto(local_proxy))
			var_table.KeyUpdate("HTTP_PROXY", &local_zval)
		}
	}
}
func PhpAutoGlobalsCreateServer(name *types.String) types.ZendBool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'S') || strchr(PG__().variables_order, 's')) {
		PhpRegisterServerVariables()
		if PG__().register_argc_argv {
			if SG__().request_info.argc {
				var argc *types.Zval
				var argv *types.Zval
				if b.Assign(&argc, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.ZSTR_ARGC.GetStr())) != nil && b.Assign(&argv, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.ZSTR_ARGV.GetStr())) != nil {
					argv.AddRefcount()
					PG__().http_globals[TRACK_VARS_SERVER].GetArr().KeyUpdate(types.ZSTR_ARGV.GetStr(), argv)
					PG__().http_globals[TRACK_VARS_SERVER].GetArr().KeyUpdate(types.ZSTR_ARGC.GetStr(), argc)
				}
			} else {
				PhpBuildArgv(SG__().request_info.query_string, &PG__().http_globals[TRACK_VARS_SERVER])
			}
		}
	} else {
		zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_SERVER])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_SERVER])
	}
	CheckHttpProxy(PG__().http_globals[TRACK_VARS_SERVER].GetArr())
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_SERVER])
	PG__().http_globals[TRACK_VARS_SERVER].AddRefcount()

	/* TODO: TRACK_VARS_SERVER is modified in a number of places (e.g. phar) past this point,
	 * where rc>1 due to the $_SERVER global. Ideally this shouldn't happen, but for now we
	 * ignore this issue, as it would probably require larger changes. */

	return 0
}
func PhpAutoGlobalsCreateEnv(name *types.String) types.ZendBool {
	zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_ENV])
	zend.ArrayInit(&PG__().http_globals[TRACK_VARS_ENV])
	if PG__().variables_order && (strchr(PG__().variables_order, 'E') || strchr(PG__().variables_order, 'e')) {
		PhpImportEnvironmentVariables(&PG__().http_globals[TRACK_VARS_ENV])
	}
	CheckHttpProxy(PG__().http_globals[TRACK_VARS_ENV].GetArr())
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_ENV])
	PG__().http_globals[TRACK_VARS_ENV].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreateRequest(name *types.String) types.ZendBool {
	var form_variables types.Zval
	var _gpc_flags []uint8 = []uint8{0, 0, 0}
	var p *byte
	zend.ArrayInit(&form_variables)
	if PG__().request_order != nil {
		p = PG__().request_order
	} else {
		p = PG__().variables_order
	}
	for ; p != nil && (*p); p++ {
		switch *p {
		case 'g':
			fallthrough
		case 'G':
			if _gpc_flags[0] == 0 {
				PhpAutoglobalMerge(form_variables.GetArr(), PG__().http_globals[TRACK_VARS_GET].GetArr())
				_gpc_flags[0] = 1
			}
		case 'p':
			fallthrough
		case 'P':
			if _gpc_flags[1] == 0 {
				PhpAutoglobalMerge(form_variables.GetArr(), PG__().http_globals[TRACK_VARS_POST].GetArr())
				_gpc_flags[1] = 1
			}
		case 'c':
			fallthrough
		case 'C':
			if _gpc_flags[2] == 0 {
				PhpAutoglobalMerge(form_variables.GetArr(), PG__().http_globals[TRACK_VARS_COOKIE].GetArr())
				_gpc_flags[2] = 1
			}
		}
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &form_variables)
	return 0
}
func PhpStartupAutoGlobals() {
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_GET", b.SizeOf("\"_GET\"")-1, 1), 0, PhpAutoGlobalsCreateGet)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_POST", b.SizeOf("\"_POST\"")-1, 1), 0, PhpAutoGlobalsCreatePost)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_COOKIE", b.SizeOf("\"_COOKIE\"")-1, 1), 0, PhpAutoGlobalsCreateCookie)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_SERVER", b.SizeOf("\"_SERVER\"")-1, 1), PG__().auto_globals_jit, PhpAutoGlobalsCreateServer)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_ENV", b.SizeOf("\"_ENV\"")-1, 1), PG__().auto_globals_jit, PhpAutoGlobalsCreateEnv)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_REQUEST", b.SizeOf("\"_REQUEST\"")-1, 1), PG__().auto_globals_jit, PhpAutoGlobalsCreateRequest)
	zend.ZendRegisterAutoGlobal(types.ZendStringInitInterned("_FILES", b.SizeOf("\"_FILES\"")-1, 1), 0, PhpAutoGlobalsCreateFiles)
}
