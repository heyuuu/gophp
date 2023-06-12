package core

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"strings"
)

func PhpRegisterVariable(var_ string, strval string, track_vars_array *types.Zval) {
	PhpRegisterVariableSafe(var_, strval, track_vars_array)
}
func PhpRegisterVariableSafe(var_ string, strval string, track_vars_array *types.Zval) {
	/* Prepare value */
	tmp := types.NewZvalString(strval)
	PhpRegisterVariableEx(var_, tmp, track_vars_array)
}
func PhpRegisterVariableQuick(name string, val *types.Zval, ht *types.Array) {
	ht.KeyUpdateIndirect(name, val)
}
func PhpRegisterVariableEx(varName string, val *types.Zval, trackVarsArray *types.Zval) {
	var p *byte = nil
	var ip *byte = nil
	var index *byte
	var var_ *byte
	var varOrig *byte
	var varLen int
	var indexLen int
	var gpcElement types.Zval
	var gpcElementP *types.Zval
	var isArray bool = false
	var symtable1 *types.Array = nil
	b.Assert(varName != "")
	if trackVarsArray != nil && trackVarsArray.IsType(types.IS_ARRAY) {
		symtable1 = trackVarsArray.Array()
	}
	if symtable1 == nil {

		/* Nothing to do */

		// zend.ZvalPtrDtorNogc(val)
		return
	}

	/* ignore leading spaces in the variable name */
	varName = strings.TrimLeft(varName, " ")

	/*
	 * Prepare variable name
	 */
	varLen = strlen(varName)
	varOrig = zend.DoAlloca(varLen+1, use_heap)
	var_ = varOrig
	memcpy(varOrig, varName, varLen+1)

	/* ensure that we don't have spaces or dots in the variable name (not binary safe) */

	for p = var_; *p; p++ {
		if (*p) == ' ' || (*p) == '.' {
			*p = '_'
		} else if (*p) == '[' {
			isArray = true
			ip = p
			*p = 0
			break
		}
	}
	varLen = p - var_

	/* Discard variable if mangling made it start with __Host-, where pre-mangling it did not start with __Host- */

	if strncmp(var_, "__Host-", b.SizeOf("\"__Host-\"")-1) == 0 && strncmp(varName, "__Host-", b.SizeOf("\"__Host-\"")-1) != 0 {
		// zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(varOrig, use_heap)
		return
	}

	/* Discard variable if mangling made it start with __Secure-, where pre-mangling it did not start with __Secure- */

	if strncmp(var_, "__Secure-", b.SizeOf("\"__Secure-\"")-1) == 0 && strncmp(varName, "__Secure-", b.SizeOf("\"__Secure-\"")-1) != 0 {
		// zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(varOrig, use_heap)
		return
	}
	if varLen == 0 {
		// zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(varOrig, use_heap)
		return
	}
	if varLen == b.SizeOf("\"this\"")-1 && zend.CurrEX() != nil {
		var ex *zend.ZendExecuteData = zend.CurrEX()
		for ex != nil {
			if ex.GetFunc() != nil && zend.ZEND_USER_CODE(ex.GetFunc().GetType()) {
				if (zend.ZEND_CALL_INFO(ex)&zend.ZEND_CALL_HAS_SYMBOL_TABLE) != 0 && ex.GetSymbolTable() == symtable1 {
					if memcmp(var_, "this", b.SizeOf("\"this\"")-1) == 0 {
						faults.ThrowError(nil, "Cannot re-assign $this")
						// zend.ZvalPtrDtorNogc(val)
						zend.FreeAlloca(varOrig, use_heap)
						return
					}
				}
				break
			}
			ex = ex.GetPrevExecuteData()
		}
	}

	/* GLOBALS hijack attempt, reject parameter */

	if symtable1 == zend.EG__().GetSymbolTable() && varLen == b.SizeOf("\"GLOBALS\"")-1 && !(memcmp(var_, "GLOBALS", b.SizeOf("\"GLOBALS\"")-1)) {
		// zend.ZvalPtrDtorNogc(val)
		zend.FreeAlloca(varOrig, use_heap)
		return
	}
	index = var_
	indexLen = varLen
	if isArray {
		var nest_level int = 0
		for true {
			var index_s *byte
			var new_idx_len int = 0
			if b.PreInc(&nest_level) > PG__().max_input_nesting_level {
				var ht *types.Array

				/* too many levels of nesting */

				if trackVarsArray != nil {
					ht = trackVarsArray.Array()
					ht.SymtableDel(b.CastStr(var_, varLen))
				}
				// zend.ZvalPtrDtorNogc(val)

				/* do not output the error message to the screen,
				   this helps us to to avoid "information disclosure" */

				if !(PG__().display_errors) {
					PhpErrorDocref(nil, faults.E_WARNING, "Input variable nesting level exceeded "+zend.ZEND_LONG_FMT+". To increase the limit change max_input_nesting_level in php.ini.", PG__().max_input_nesting_level)
				}
				zend.FreeAlloca(varOrig, use_heap)
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
					indexLen = 0
					if index != nil {
						indexLen = strlen(index)
					}
					goto plain_var
					return
				}
				*ip = 0
				new_idx_len = strlen(index_s)
			}
			if index == nil {
				zend.ArrayInit(&gpcElement)
				if b.Assign(&gpcElementP, symtable1.Append(&gpcElement)) == nil {
					gpcElement.Array().Destroy()
					// zend.ZvalPtrDtorNogc(val)
					zend.FreeAlloca(varOrig, use_heap)
					return
				}
			} else {
				gpcElementP = symtable1.SymtableFind(b.CastStr(index, indexLen))
				if gpcElementP == nil {
					var tmp types.Zval
					zend.ArrayInit(&tmp)
					gpcElementP = symtable1.SymtableUpdateInd(b.CastStr(index, indexLen), &tmp)
				} else {
					if gpcElementP.IsIndirect() {
						gpcElementP = gpcElementP.Indirect()
					}
					if !gpcElementP.IsArray() {
						// zend.ZvalPtrDtorNogc(gpc_element_p)
						zend.ArrayInit(gpcElementP)
					} else {
						types.SeparateArray(gpcElementP)
					}
				}
			}
			symtable1 = gpcElementP.Array()

			/* ip pointed to the '[' character, now obtain the key */

			index = index_s
			indexLen = new_idx_len
			ip++
			if (*ip) == '[' {
				isArray = true
				*ip = 0
			} else {
				goto plain_var
			}
		}
	} else {
	plain_var:
		if index == nil {
			if symtable1.Append(val) == nil {
				// zend.ZvalPtrDtorNogc(val)
			}
		} else {
			var idx zend.ZendUlong

			/*
			 * According to rfc2965, more specific paths are listed above the less specific ones.
			 * If we encounter a duplicate cookie name, we should skip it, since it is not possible
			 * to have the same (plain text) cookie name for the same path and we should not overwrite
			 * more specific cookies with the less specific ones.
			 */

			if PG__().http_globals[TRACK_VARS_COOKIE].IsNotUndef() && symtable1 == PG__().http_globals[TRACK_VARS_COOKIE].Array() && symtable1.SymtableExists(b.CastStr(index, indexLen)) {
				// zend.ZvalPtrDtorNogc(val)
			} else if types.HandleNumericStr(b.CastStr(index, indexLen), &idx) {
				symtable1.IndexUpdate(idx, val)
			} else {
				PhpRegisterVariableQuick(b.CastStr(index, indexLen), val, symtable1)
			}

			/*
			 * According to rfc2965, more specific paths are listed above the less specific ones.
			 * If we encounter a duplicate cookie name, we should skip it, since it is not possible
			 * to have the same (plain text) cookie name for the same path and we should not overwrite
			 * more specific cookies with the less specific ones.
			 */

		}
	}
	zend.FreeAlloca(varOrig, use_heap)
}
func AddPostVar(arr *types.Zval, var_ *PostVarDataT, eof bool) bool {
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
		PhpRegisterVariableSafe(b.CastStrAuto(var_.GetPtr()), b.CastStr(val, new_vlen), arr)
	}
	zend.Efree(val)
	var_.SetPtr(vsep + (vsep != var_.GetEnd()))
	var_.SetAlreadyScanned(0)
	return 1
}
func AddPostVars(arr *types.Zval, vars *PostVarDataT, eof bool) int {
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
	var s *PhpStream = SG__().RequestInfo.request_body
	var post_data PostVarDataT
	if s != nil && types.SUCCESS == PhpStreamRewind(s) {
		memset(&post_data, 0, b.SizeOf("post_data"))
		for PhpStreamEof(s) == 0 {
			var buf []byte = []byte{0}
			var len_ ssize_t = PhpStreamRead(s, buf, r.BUFSIZ)
			if len_ > 0 {
				post_data.GetStr().WriteString(b.CastStr(buf, len_))
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
			// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_POST])
			types.ZVAL_COPY_VALUE(&PG__().http_globals[TRACK_VARS_POST], &array)
		case PARSE_GET:
			// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_GET])
			types.ZVAL_COPY_VALUE(&PG__().http_globals[TRACK_VARS_GET], &array)
		case PARSE_COOKIE:
			// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_COOKIE])
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
		c_var = SG__().RequestInfo.query_string
		if c_var != nil && (*c_var) {
			res = (*byte)(zend.Estrdup(c_var))
			free_buffer = 1
		} else {
			free_buffer = 0
		}
	} else if arg == PARSE_COOKIE {
		c_var = SG__().RequestInfo.cookie_data
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
			PhpRegisterVariableSafe(b.CastStrAuto(var_), b.CastStr(val, new_val_len), &array)
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
	var idx zend.ZendUlong
	p = strchr(env, '=')
	if p == nil || p == env || ValidEnvironmentName(env, p) == 0 {
		/* malformed entry? */
		return
	}
	name_len = p - env
	p++

	val := types.NewZvalString(b.CastStrAuto(p))
	if types.HandleNumericStr(b.CastStr(env, name_len), &idx) {
		ht.IndexUpdate(idx, val)
	} else {
		PhpRegisterVariableQuick(b.CastStr(env, name_len), val, ht)
	}
}
func DupEnvVariables() *types.Array {
	return Env__().DupArray()
}

func _phpImportEnvironmentVariables(array_ptr *types.Zval) {
	var env **byte
	tsrm_env_lock()
	for env = Environ; env != nil && (*env) != nil; env++ {
		ImportEnvironmentVariable(array_ptr.Array(), *env)
	}
	tsrm_env_unlock()
}
func PhpBuildArgv(s *byte, track_vars_array *types.Zval) {
	var arr types.Zval
	var argc types.Zval
	var tmp types.Zval
	var count int = 0
	var ss *byte
	var space *byte
	if !(SG__().RequestInfo.argc || track_vars_array != nil) {
		return
	}
	zend.ArrayInit(&arr)

	/* Prepare argv */

	if SG__().RequestInfo.argc {
		var i int
		for i = 0; i < SG__().RequestInfo.argc; i++ {
			tmp.SetStringVal(b.CastStrAuto(SG__().RequestInfo.argv[i]))
			if arr.Array().Append(&tmp) == nil {
				// types.ZendStringEfree(tmp.String())
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
			if arr.Array().Append(&tmp) == nil {
				// types.ZendStringEfree(tmp.String())
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

	if SG__().RequestInfo.argc {
		argc.SetLong(SG__().RequestInfo.argc)
	} else {
		argc.SetLong(count)
	}
	if SG__().RequestInfo.argc {
		// 		arr.AddRefcount()
		zend.EG__().GetSymbolTable().KeyUpdate(types.STR_ARGV, &arr)
		zend.EG__().GetSymbolTable().KeyUpdate(types.STR_ARGC, &argc)
	}
	if track_vars_array != nil && track_vars_array.IsType(types.IS_ARRAY) {
		// 		arr.AddRefcount()
		track_vars_array.Array().KeyUpdate(types.STR_ARGV, &arr)
		track_vars_array.Array().KeyUpdate(types.STR_ARGC, &argc)
	}
	// zend.ZvalPtrDtorNogc(&arr)
}
func PhpRegisterServerVariables() {
	var tmp types.Zval
	var arr *types.Zval = &PG__().http_globals[TRACK_VARS_SERVER]
	var ht *types.Array
	// zend.ZvalPtrDtorNogc(arr)
	zend.ArrayInit(arr)

	/* Server variables */

	if SM__().GetRegisterServerVariables() != nil {
		SM__().GetRegisterServerVariables()(arr)
	}
	ht = arr.Array()

	/* PHP Authentication support */

	if SG__().RequestInfo.auth_user {
		tmp.SetStringVal(b.CastStrAuto(SG__().RequestInfo.auth_user))
		PhpRegisterVariableQuick("PHP_AUTH_USER", &tmp, ht)
	}
	if SG__().RequestInfo.auth_password {
		tmp.SetStringVal(b.CastStrAuto(SG__().RequestInfo.auth_password))
		PhpRegisterVariableQuick("PHP_AUTH_PW", &tmp, ht)
	}
	if SG__().RequestInfo.auth_digest {
		tmp.SetStringVal(b.CastStrAuto(SG__().RequestInfo.auth_digest))
		PhpRegisterVariableQuick("PHP_AUTH_DIGEST", &tmp, ht)
	}

	/* store request init time */

	tmp.SetDouble(SapiGetRequestTime())
	PhpRegisterVariableQuick("REQUEST_TIME_FLOAT", &tmp, ht)
	tmp.SetLong(operators.DvalToLval(tmp.Double()))
	PhpRegisterVariableQuick("REQUEST_TIME", &tmp, ht)
}
func PhpAutoglobalMerge(dest *types.Array, src *types.Array) {
	var dest_entry *types.Zval
	var globalsCheck = dest == zend.EG__().GetSymbolTable()

	src.Foreach(func(key types.ArrayKey, value *types.Zval) {
		if !value.IsArray() || key.IsStrKey() && b.Assign(&dest_entry, dest.KeyFind(key.StrKey())) == nil || !key.IsStrKey() && b.Assign(&dest_entry, dest.IndexFind(key.IdxKey())) == nil || !dest_entry.IsArray() {
			if key.IsStrKey() {
				if !globalsCheck || key.StrKey() != "GLOBALS" {
					dest.KeyUpdate(key.StrKey(), value)
				}
			} else {
				dest.IndexUpdate(key.IdxKey(), value)
			}
		} else {
			types.SeparateArray(dest_entry)
			PhpAutoglobalMerge(dest_entry.Array(), value.Array())
		}
	})
}
func PhpHashEnvironment() int {
	memset(PG__().http_globals, 0, b.SizeOf("PG ( http_globals )"))
	zend.ZendActivateAutoGlobals()
	if PG__().register_argc_argv {
		PhpBuildArgv(SG__().RequestInfo.query_string, &PG__().http_globals[TRACK_VARS_SERVER])
	}
	return types.SUCCESS
}
func PhpAutoGlobalsCreateGet(name *types.String) bool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'G') || strchr(PG__().variables_order, 'g')) {
		SM__().GetTreatData()(PARSE_GET, nil, nil)
	} else {
		// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_GET])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_GET])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_GET])
	//PG__().http_globals[TRACK_VARS_GET].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreatePost(name *types.String) bool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'P') || strchr(PG__().variables_order, 'p')) && !(SG__().headers_sent) && SG__().RequestInfo.request_method && !(strcasecmp(SG__().RequestInfo.request_method, "POST")) {
		SM__().GetTreatData()(PARSE_POST, nil, nil)
	} else {
		// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_POST])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_POST])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_POST])
	//PG__().http_globals[TRACK_VARS_POST].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreateCookie(name *types.String) bool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'C') || strchr(PG__().variables_order, 'c')) {
		SM__().GetTreatData()(PARSE_COOKIE, nil, nil)
	} else {
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_COOKIE])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_COOKIE])
	return 0
}
func PhpAutoGlobalsCreateFiles(name *types.String) bool {
	if PG__().http_globals[TRACK_VARS_FILES].IsUndef() {
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_FILES])
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_FILES])
	return 0
}
func CheckHttpProxy(var_table *types.Array) {
	if var_table.KeyExists("HTTP_PROXY") {
		var local_proxy *byte = getenv("HTTP_PROXY")
		if local_proxy == nil {
			var_table.KeyDelete("HTTP_PROXY")
		} else {
			var local_zval types.Zval
			local_zval.SetStringVal(b.CastStrAuto(local_proxy))
			var_table.KeyUpdate("HTTP_PROXY", &local_zval)
		}
	}
}
func PhpAutoGlobalsCreateServer(name *types.String) bool {
	if PG__().variables_order && (strchr(PG__().variables_order, 'S') || strchr(PG__().variables_order, 's')) {
		PhpRegisterServerVariables()
		if PG__().register_argc_argv {
			if SG__().RequestInfo.argc {
				var argc *types.Zval
				var argv *types.Zval
				if b.Assign(&argc, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.STR_ARGC)) != nil && b.Assign(&argv, types.ZendHashFindInd(zend.EG__().GetSymbolTable(), types.STR_ARGV)) != nil {
					// 					argv.AddRefcount()
					PG__().http_globals[TRACK_VARS_SERVER].Array().KeyUpdate(types.STR_ARGV, argv)
					PG__().http_globals[TRACK_VARS_SERVER].Array().KeyUpdate(types.STR_ARGC, argc)
				}
			} else {
				PhpBuildArgv(SG__().RequestInfo.query_string, &PG__().http_globals[TRACK_VARS_SERVER])
			}
		}
	} else {
		// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_SERVER])
		zend.ArrayInit(&PG__().http_globals[TRACK_VARS_SERVER])
	}
	CheckHttpProxy(PG__().http_globals[TRACK_VARS_SERVER].Array())
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_SERVER])
	//PG__().http_globals[TRACK_VARS_SERVER].AddRefcount()

	/* TODO: TRACK_VARS_SERVER is modified in a number of places (e.g. phar) past this point,
	 * where rc>1 due to the $_SERVER global. Ideally this shouldn't happen, but for now we
	 * ignore this issue, as it would probably require larger changes. */

	return 0
}
func PhpAutoGlobalsCreateEnv(name *types.String) bool {
	// zend.ZvalPtrDtorNogc(&PG__().http_globals[TRACK_VARS_ENV])
	zend.ArrayInit(&PG__().http_globals[TRACK_VARS_ENV])
	if PG__().variables_order && (strchr(PG__().variables_order, 'E') || strchr(PG__().variables_order, 'e')) {
		PhpImportEnvironmentVariables(&PG__().http_globals[TRACK_VARS_ENV])
	}
	CheckHttpProxy(PG__().http_globals[TRACK_VARS_ENV].Array())
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &PG__().http_globals[TRACK_VARS_ENV])
	//PG__().http_globals[TRACK_VARS_ENV].AddRefcount()
	return 0
}
func PhpAutoGlobalsCreateRequest(name *types.String) bool {
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
				PhpAutoglobalMerge(form_variables.Array(), PG__().http_globals[TRACK_VARS_GET].Array())
				_gpc_flags[0] = 1
			}
		case 'p':
			fallthrough
		case 'P':
			if _gpc_flags[1] == 0 {
				PhpAutoglobalMerge(form_variables.Array(), PG__().http_globals[TRACK_VARS_POST].Array())
				_gpc_flags[1] = 1
			}
		case 'c':
			fallthrough
		case 'C':
			if _gpc_flags[2] == 0 {
				PhpAutoglobalMerge(form_variables.Array(), PG__().http_globals[TRACK_VARS_COOKIE].Array())
				_gpc_flags[2] = 1
			}
		}
	}
	zend.EG__().GetSymbolTable().KeyUpdate(name.GetStr(), &form_variables)
	return 0
}
func PhpStartupAutoGlobals() {
	zend.ZendRegisterAutoGlobal(types.NewString("_GET"), 0, PhpAutoGlobalsCreateGet)
	zend.ZendRegisterAutoGlobal(types.NewString("_POST"), 0, PhpAutoGlobalsCreatePost)
	zend.ZendRegisterAutoGlobal(types.NewString("_COOKIE"), 0, PhpAutoGlobalsCreateCookie)
	zend.ZendRegisterAutoGlobal(types.NewString("_SERVER"), PG__().auto_globals_jit, PhpAutoGlobalsCreateServer)
	zend.ZendRegisterAutoGlobal(types.NewString("_ENV"), PG__().auto_globals_jit, PhpAutoGlobalsCreateEnv)
	zend.ZendRegisterAutoGlobal(types.NewString("_REQUEST"), PG__().auto_globals_jit, PhpAutoGlobalsCreateRequest)
	zend.ZendRegisterAutoGlobal(types.NewString("_FILES"), 0, PhpAutoGlobalsCreateFiles)
}
