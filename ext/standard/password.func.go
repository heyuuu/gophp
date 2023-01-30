// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func PhpPasswordAlgoRegister(ident string, algo *PhpPasswordAlgo) int {
	var zalgo zend.Zval
	zend.ZVAL_PTR(&zalgo, (*PhpPasswordAlgo)(algo))
	if PhpPasswordAlgos.KeyAdd(b.CastStr(ident, strlen(ident)), &zalgo) != nil {
		return zend.SUCCESS
	}
	return zend.FAILURE
}
func PhpPasswordAlgoUnregister(ident *byte) {
	zend.ZendHashStrDel(&PhpPasswordAlgos, ident, strlen(ident))
}
func PhpPasswordSaltIsAlphabet(str *byte, len_ int) int {
	var i int = 0
	for i = 0; i < len_; i++ {
		if !(str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9' || str[i] == '.' || str[i] == '/') {
			return zend.FAILURE
		}
	}
	return zend.SUCCESS
}
func PhpPasswordSaltTo64(str *byte, str_len int, out_len int, ret *byte) int {
	var pos int = 0
	var buffer *zend.ZendString
	if int(str_len < 0) != 0 {
		return zend.FAILURE
	}
	buffer = PhpBase64Encode((*uint8)(str), str_len)
	if buffer.GetLen() < out_len {

		/* Too short of an encoded string generated */

		zend.ZendStringReleaseEx(buffer, 0)
		return zend.FAILURE
	}
	for pos = 0; pos < out_len; pos++ {
		if buffer.GetVal()[pos] == '+' {
			ret[pos] = '.'
		} else if buffer.GetVal()[pos] == '=' {
			zend.ZendStringFree(buffer)
			return zend.FAILURE
		} else {
			ret[pos] = buffer.GetVal()[pos]
		}
	}
	zend.ZendStringFree(buffer)
	return zend.SUCCESS
}
func PhpPasswordMakeSalt(length int) *zend.ZendString {
	var ret *zend.ZendString
	var buffer *zend.ZendString
	if length > core.INT_MAX/3 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Length is too large to safely generate")
		return nil
	}
	buffer = zend.ZendStringAlloc(length*3/4+1, 0)
	if zend.FAILURE == PhpRandomBytesSilent(buffer.GetVal(), buffer.GetLen()) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to generate salt")
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	ret = zend.ZendStringAlloc(length, 0)
	if PhpPasswordSaltTo64(buffer.GetVal(), buffer.GetLen(), length, ret.GetVal()) == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Generated salt too short")
		zend.ZendStringReleaseEx(buffer, 0)
		zend.ZendStringReleaseEx(ret, 0)
		return nil
	}
	zend.ZendStringReleaseEx(buffer, 0)
	ret.GetVal()[length] = 0
	return ret
}
func PhpPasswordGetSalt(unused_ *zend.Zval, required_salt_len int, options *zend.HashTable) *zend.ZendString {
	var buffer *zend.ZendString
	var option_buffer *zend.Zval
	if options == nil || !(b.Assign(&option_buffer, options.FindByStrPtr("salt", b.SizeOf("\"salt\"")-1))) {
		return PhpPasswordMakeSalt(required_salt_len)
	}
	core.PhpErrorDocref(nil, zend.E_DEPRECATED, "Use of the 'salt' option to password_hash is deprecated")
	switch option_buffer.GetType() {
	case zend.IS_STRING:
		buffer = option_buffer.GetStr().Copy()
		break
	case zend.IS_LONG:

	case zend.IS_DOUBLE:

	case zend.IS_OBJECT:
		buffer = zend.ZvalTryGetString(option_buffer)
		if buffer == nil {
			return nil
		}
		break
	case zend.IS_FALSE:

	case zend.IS_TRUE:

	case zend.IS_NULL:

	case zend.IS_RESOURCE:

	case zend.IS_ARRAY:

	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Non-string salt parameter supplied")
		return nil
	}

	/* XXX all the crypt related APIs work with int for string length.
	   That should be revised for size_t and then we maybe don't require
	   the > INT_MAX check. */

	if zend.ZEND_SIZE_T_INT_OVFL(buffer.GetLen()) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Supplied salt is too long")
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if buffer.GetLen() < required_salt_len {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Provided salt is too short: %zd expecting %zd", buffer.GetLen(), required_salt_len)
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if PhpPasswordSaltIsAlphabet(buffer.GetVal(), buffer.GetLen()) == zend.FAILURE {
		var salt *zend.ZendString = zend.ZendStringAlloc(required_salt_len, 0)
		if PhpPasswordSaltTo64(buffer.GetVal(), buffer.GetLen(), required_salt_len, salt.GetVal()) == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Provided salt is too short: %zd", buffer.GetLen())
			zend.ZendStringReleaseEx(salt, 0)
			zend.ZendStringReleaseEx(buffer, 0)
			return nil
		}
		zend.ZendStringReleaseEx(buffer, 0)
		return salt
	} else {
		var salt *zend.ZendString = zend.ZendStringAlloc(required_salt_len, 0)
		memcpy(salt.GetVal(), buffer.GetVal(), required_salt_len)
		zend.ZendStringReleaseEx(buffer, 0)
		return salt
	}
}
func PhpPasswordBcryptValid(hash *zend.ZendString) zend.ZendBool {
	var h *byte = hash.GetVal()
	return hash.GetLen() == 60 && h[0] == '$' && h[1] == '2' && h[2] == 'y'
}
func PhpPasswordBcryptGetInfo(return_value *zend.Zval, hash *zend.ZendString) int {
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return zend.FAILURE

		/* Should never get called this way. */

	}
	sscanf(hash.GetVal(), "$2y$"+zend.ZEND_LONG_FMT+"$", &cost)
	zend.AddAssocLong(return_value, "cost", cost)
	return zend.SUCCESS
}
func PhpPasswordBcryptNeedsRehash(hash *zend.ZendString, options *zend.ZendArray) zend.ZendBool {
	var znew_cost *zend.Zval
	var old_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	var new_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return 1

		/* Should never get called this way. */

	}
	sscanf(hash.GetVal(), "$2y$"+zend.ZEND_LONG_FMT+"$", &old_cost)
	if options != nil && b.Assign(&znew_cost, options.FindByStrPtr("cost", b.SizeOf("\"cost\"")-1)) != nil {
		new_cost = zend.ZvalGetLong(znew_cost)
	}
	return old_cost != new_cost
}
func PhpPasswordBcryptVerify(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool {
	var i int
	var status int = 0
	var ret *zend.ZendString = PhpCrypt(password.GetVal(), int(password.GetLen()), hash.GetVal(), int(hash.GetLen()), 1)
	if ret == nil {
		return 0
	}
	if ret.GetLen() != hash.GetLen() || hash.GetLen() < 13 {
		zend.ZendStringFree(ret)
		return 0
	}

	/* We're using this method instead of == in order to provide
	 * resistance towards timing attacks. This is a constant time
	 * equality check that will always check every byte of both
	 * values. */

	for i = 0; i < hash.GetLen(); i++ {
		status |= ret.GetVal()[i] ^ hash.GetVal()[i]
	}
	zend.ZendStringFree(ret)
	return status == 0
}
func PhpPasswordBcryptHash(password *zend.ZendString, options *zend.ZendArray) *zend.ZendString {
	var hash_format []byte
	var hash_format_len int
	var result *zend.ZendString
	var hash *zend.ZendString
	var salt *zend.ZendString
	var zcost *zend.Zval
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if options != nil && b.Assign(&zcost, options.FindByStrPtr("cost", b.SizeOf("\"cost\"")-1)) != nil {
		cost = zend.ZvalGetLong(zcost)
	}
	if cost < 4 || cost > 31 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid bcrypt cost parameter specified: "+zend.ZEND_LONG_FMT, cost)
		return nil
	}
	hash_format_len = core.Snprintf(hash_format, b.SizeOf("hash_format"), "$2y$%02"+zend.ZEND_LONG_FMT_SPEC+"$", cost)
	if !(b.Assign(&salt, PhpPasswordGetSalt(nil, uint64(22), options))) {
		return nil
	}
	salt.GetVal()[salt.GetLen()] = 0
	hash = zend.ZendStringAlloc(salt.GetLen()+hash_format_len, 0)
	sprintf(hash.GetVal(), "%s%s", hash_format, salt.GetVal())
	hash.GetVal()[hash_format_len+salt.GetLen()] = 0
	zend.ZendStringReleaseEx(salt, 0)

	/* This cast is safe, since both values are defined here in code and cannot overflow */

	result = PhpCrypt(password.GetVal(), int(password.GetLen()), hash.GetVal(), int(hash.GetLen()), 1)
	zend.ZendStringReleaseEx(hash, 0)
	if result == nil {
		return nil
	}
	if result.GetLen() < 13 {
		zend.ZendStringFree(result)
		return nil
	}
	return result
}
func ZmStartupPassword(type_ int, module_number int) int {
	zend.ZendHashInit(&PhpPasswordAlgos, 4, nil, zend.ZVAL_PTR_DTOR, 1)
	zend.REGISTER_STRING_CONSTANT("PASSWORD_DEFAULT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT)
	if zend.FAILURE == PhpPasswordAlgoRegister("2y", &PhpPasswordAlgoBcrypt) {
		return zend.FAILURE
	}
	zend.REGISTER_STRING_CONSTANT("PASSWORD_BCRYPT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("PASSWORD_BCRYPT_DEFAULT_COST", PHP_PASSWORD_BCRYPT_COST, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func ZmShutdownPassword(type_ int, module_number int) int {
	zend.ZendHashDestroy(&PhpPasswordAlgos)
	return zend.SUCCESS
}
func PhpPasswordAlgoDefault() *PhpPasswordAlgo { return &PhpPasswordAlgoBcrypt }
func PhpPasswordAlgoFind(ident *zend.ZendString) *PhpPasswordAlgo {
	var tmp *zend.Zval
	if ident == nil {
		return nil
	}
	tmp = PhpPasswordAlgos.FindByZendString((*zend.ZendString)(ident))
	if tmp == nil || tmp.GetType() != zend.IS_PTR {
		return nil
	}
	return tmp.GetPtr()
}
func PhpPasswordAlgoFindZvalEx(arg *zend.Zval, default_algo *PhpPasswordAlgo) *PhpPasswordAlgo {
	if arg == nil || arg.IsType(zend.IS_NULL) {
		return default_algo
	}
	if arg.IsType(zend.IS_LONG) {
		switch arg.GetLval() {
		case 0:
			return default_algo
		case 1:
			return &PhpPasswordAlgoBcrypt
		case 2:
			var n *zend.ZendString = zend.ZendStringInit("argon2i", b.SizeOf("\"argon2i\"")-1, 0)
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			zend.ZendStringRelease(n)
			return ret
		case 3:
			var n *zend.ZendString = zend.ZendStringInit("argon2id", b.SizeOf("\"argon2id\"")-1, 0)
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			zend.ZendStringRelease(n)
			return ret
		}
		return nil
	}
	if arg.GetType() != zend.IS_STRING {
		return nil
	}
	return PhpPasswordAlgoFind(arg.GetStr())
}
func PhpPasswordAlgoFindZval(arg *zend.Zval) *PhpPasswordAlgo {
	return PhpPasswordAlgoFindZvalEx(arg, PhpPasswordAlgoDefault())
}
func PhpPasswordAlgoExtractIdent(hash *zend.ZendString) *zend.ZendString {
	var ident *byte
	var ident_end *byte
	if hash == nil || hash.GetLen() < 3 {

		/* Minimum prefix: "$x$" */

		return nil

		/* Minimum prefix: "$x$" */

	}
	ident = hash.GetVal() + 1
	ident_end = strchr(ident, '$')
	if ident_end == nil {

		/* No terminating '$' */

		return nil

		/* No terminating '$' */

	}
	return zend.ZendStringInit(ident, ident_end-ident, 0)
}
func PhpPasswordAlgoIdentifyEx(hash *zend.ZendString, default_algo *PhpPasswordAlgo) *PhpPasswordAlgo {
	var algo *PhpPasswordAlgo
	var ident *zend.ZendString = PhpPasswordAlgoExtractIdent(hash)
	if ident == nil {
		return default_algo
	}
	algo = PhpPasswordAlgoFind(ident)
	zend.ZendStringRelease(ident)
	if algo == nil || algo.GetValid() != nil && algo.GetValid()(hash) == 0 {
		return default_algo
	} else {
		return algo
	}
}
func ZifPasswordGetInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var algo *PhpPasswordAlgo
	var hash *zend.ZendString
	var ident *zend.ZendString
	var options zend.Zval
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
			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	zend.ArrayInit(return_value)
	zend.ArrayInit(&options)
	ident = PhpPasswordAlgoExtractIdent(hash)
	algo = PhpPasswordAlgoFind(ident)
	if algo == nil || algo.GetValid() != nil && algo.GetValid()(hash) == 0 {
		if ident != nil {
			zend.ZendStringRelease(ident)
		}
		zend.AddAssocNull(return_value, "algo")
		zend.AddAssocString(return_value, "algoName", "unknown")
		zend.AddAssocZval(return_value, "options", &options)
		return
	}
	zend.AddAssocStr(return_value, "algo", PhpPasswordAlgoExtractIdent(hash))
	zend.ZendStringRelease(ident)
	zend.AddAssocString(return_value, "algoName", algo.GetName())
	if algo.GetGetInfo() != nil && zend.FAILURE == algo.GetGetInfo()(&options, hash) {
		zend.ZvalDtor(&options)
		zend.ZvalDtor(return_value)
		zend.RETVAL_NULL()
		return
	}
	zend.AddAssocZval(return_value, "options", &options)
}
func ZifPasswordNeedsRehash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var old_algo *PhpPasswordAlgo
	var new_algo *PhpPasswordAlgo
	var hash *zend.ZendString
	var znew_algo *zend.Zval
	var options *zend.ZendArray = 0
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &znew_algo, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArrayHt(_arg, &options, 0, 1, 0) == 0 {
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
			return
		}
		break
	}
	new_algo = PhpPasswordAlgoFindZval(znew_algo)
	if new_algo == nil {

		/* Unknown new algorithm, never prompt to rehash. */

		zend.RETVAL_FALSE
		return
	}
	old_algo = PhpPasswordAlgoIdentifyEx(hash, nil)
	if old_algo != new_algo {

		/* Different algorithm preferred, always rehash. */

		zend.RETVAL_TRUE
		return
	}
	zend.RETVAL_BOOL(old_algo.GetNeedsRehash()(hash, options) != 0)
	return
}
func ZifPasswordVerify(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var password *zend.ZendString
	var hash *zend.ZendString
	var algo *PhpPasswordAlgo
	for {
		var _flags int = 0
		var _min_num_args int = 2
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
			if zend.ZendParseArgStr(_arg, &password, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	algo = PhpPasswordAlgoIdentify(hash)
	zend.RETVAL_BOOL(algo != nil && (algo.GetVerify() == nil || algo.GetVerify()(password, hash) != 0))
	return
}
func ZifPasswordHash(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var password *zend.ZendString
	var digest *zend.ZendString = nil
	var zalgo *zend.Zval
	var algo *PhpPasswordAlgo
	var options *zend.ZendArray = nil
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
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
			if zend.ZendParseArgStr(_arg, &password, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = zend.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			zend.ZendParseArgZvalDeref(_arg, &zalgo, 0)
			_optional = 1
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgArrayHt(_arg, &options, 0, 1, 0) == 0 {
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
			return
		}
		break
	}
	algo = PhpPasswordAlgoFindZval(zalgo)
	if algo == nil {
		var algostr *zend.ZendString = zend.ZvalGetString(zalgo)
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown password hashing algorithm: %s", algostr.GetVal())
		zend.ZendStringRelease(algostr)
		zend.RETVAL_NULL()
		return
	}
	digest = algo.GetHash()(password, options)
	if digest == nil {

		/* algo->hash should have raised an error. */

		zend.RETVAL_NULL()
		return
	}
	zend.RETVAL_NEW_STR(digest)
	return
}
func ZifPasswordAlgos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var algo *zend.ZendString
	if zend.ZEND_NUM_ARGS() != 0 {
		zend.ZendWrongParametersNoneError()
		return
	}
	zend.ArrayInit(return_value)
	var __ht *zend.HashTable = &PhpPasswordAlgos
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		algo = _p.GetKey()
		zend.AddNextIndexStr(return_value, algo.Copy())
	}
}
