// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/password.c>

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
   | Authors: Anthony Ferrara <ircmaxell@php.net>                         |
   |          Charles R. Portwood II <charlesportwoodii@erianna.com>      |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include "php.h"

// # include "fcntl.h"

// # include "php_password.h"

// # include "php_rand.h"

// # include "php_crypt.h"

// # include "base64.h"

// # include "zend_interfaces.h"

// # include "info.h"

// # include "php_random.h"

var PhpPasswordAlgos zend.ZendArray

func PhpPasswordAlgoRegister(ident string, algo *PhpPasswordAlgo) int {
	var zalgo zend.Zval
	&zalgo.value.ptr = (*PhpPasswordAlgo)(algo)
	&zalgo.u1.type_info = 14
	if zend.ZendHashStrAdd(&PhpPasswordAlgos, ident, strlen(ident), &zalgo) != nil {
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

/* }}} */

func PhpPasswordSaltTo64(str *byte, str_len int, out_len int, ret *byte) int {
	var pos int = 0
	var buffer *zend.ZendString
	if int(str_len < 0) != 0 {
		return zend.FAILURE
	}
	buffer = PhpBase64Encode((*uint8)(str), str_len)
	if buffer.len_ < out_len {

		/* Too short of an encoded string generated */

		zend.ZendStringReleaseEx(buffer, 0)
		return zend.FAILURE
	}
	for pos = 0; pos < out_len; pos++ {
		if buffer.val[pos] == '+' {
			ret[pos] = '.'
		} else if buffer.val[pos] == '=' {
			zend.ZendStringFree(buffer)
			return zend.FAILURE
		} else {
			ret[pos] = buffer.val[pos]
		}
	}
	zend.ZendStringFree(buffer)
	return zend.SUCCESS
}

/* }}} */

func PhpPasswordMakeSalt(length int) *zend.ZendString {
	var ret *zend.ZendString
	var buffer *zend.ZendString
	if length > 2147483647/3 {
		core.PhpErrorDocref(nil, 1<<1, "Length is too large to safely generate")
		return nil
	}
	buffer = zend.ZendStringAlloc(length*3/4+1, 0)
	if zend.FAILURE == PhpRandomBytes(buffer.val, buffer.len_, 0) {
		core.PhpErrorDocref(nil, 1<<1, "Unable to generate salt")
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	ret = zend.ZendStringAlloc(length, 0)
	if PhpPasswordSaltTo64(buffer.val, buffer.len_, length, ret.val) == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "Generated salt too short")
		zend.ZendStringReleaseEx(buffer, 0)
		zend.ZendStringReleaseEx(ret, 0)
		return nil
	}
	zend.ZendStringReleaseEx(buffer, 0)
	ret.val[length] = 0
	return ret
}

/* }}} */

func PhpPasswordGetSalt(unused_ *zend.Zval, required_salt_len int, options *zend.HashTable) *zend.ZendString {
	var buffer *zend.ZendString
	var option_buffer *zend.Zval
	if options == nil || !(g.Assign(&option_buffer, zend.ZendHashStrFind(options, "salt", g.SizeOf("\"salt\"")-1))) {
		return PhpPasswordMakeSalt(required_salt_len)
	}
	core.PhpErrorDocref(nil, 1<<13, "Use of the 'salt' option to password_hash is deprecated")
	switch option_buffer.u1.v.type_ {
	case 6:
		buffer = zend.ZendStringCopy(option_buffer.value.str)
		break
	case 4:

	case 5:

	case 8:
		buffer = zend.ZvalTryGetString(option_buffer)
		if buffer == nil {
			return nil
		}
		break
	case 2:

	case 3:

	case 1:

	case 9:

	case 7:

	default:
		core.PhpErrorDocref(nil, 1<<1, "Non-string salt parameter supplied")
		return nil
	}

	/* XXX all the crypt related APIs work with int for string length.
	   That should be revised for size_t and then we maybe don't require
	   the > INT_MAX check. */

	if buffer.len_ > int(2147483647) {
		core.PhpErrorDocref(nil, 1<<1, "Supplied salt is too long")
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if buffer.len_ < required_salt_len {
		core.PhpErrorDocref(nil, 1<<1, "Provided salt is too short: %zd expecting %zd", buffer.len_, required_salt_len)
		zend.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if PhpPasswordSaltIsAlphabet(buffer.val, buffer.len_) == zend.FAILURE {
		var salt *zend.ZendString = zend.ZendStringAlloc(required_salt_len, 0)
		if PhpPasswordSaltTo64(buffer.val, buffer.len_, required_salt_len, salt.val) == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "Provided salt is too short: %zd", buffer.len_)
			zend.ZendStringReleaseEx(salt, 0)
			zend.ZendStringReleaseEx(buffer, 0)
			return nil
		}
		zend.ZendStringReleaseEx(buffer, 0)
		return salt
	} else {
		var salt *zend.ZendString = zend.ZendStringAlloc(required_salt_len, 0)
		memcpy(salt.val, buffer.val, required_salt_len)
		zend.ZendStringReleaseEx(buffer, 0)
		return salt
	}
}

/* bcrypt implementation */

func PhpPasswordBcryptValid(hash *zend.ZendString) zend.ZendBool {
	var h *byte = hash.val
	return hash.len_ == 60 && h[0] == '$' && h[1] == '2' && h[2] == 'y'
}
func PhpPasswordBcryptGetInfo(return_value *zend.Zval, hash *zend.ZendString) int {
	var cost zend.ZendLong = 10
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return zend.FAILURE

		/* Should never get called this way. */

	}
	sscanf(hash.val, "$2y$"+"%"+"lld"+"$", &cost)
	zend.AddAssocLongEx(return_value, "cost", strlen("cost"), cost)
	return zend.SUCCESS
}
func PhpPasswordBcryptNeedsRehash(hash *zend.ZendString, options *zend.ZendArray) zend.ZendBool {
	var znew_cost *zend.Zval
	var old_cost zend.ZendLong = 10
	var new_cost zend.ZendLong = 10
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return 1

		/* Should never get called this way. */

	}
	sscanf(hash.val, "$2y$"+"%"+"lld"+"$", &old_cost)
	if options != nil && g.Assign(&znew_cost, zend.ZendHashStrFind(options, "cost", g.SizeOf("\"cost\"")-1)) != nil {
		new_cost = zend.ZvalGetLong(znew_cost)
	}
	return old_cost != new_cost
}
func PhpPasswordBcryptVerify(password *zend.ZendString, hash *zend.ZendString) zend.ZendBool {
	var i int
	var status int = 0
	var ret *zend.ZendString = PhpCrypt(password.val, int(password.len_), hash.val, int(hash.len_), 1)
	if ret == nil {
		return 0
	}
	if ret.len_ != hash.len_ || hash.len_ < 13 {
		zend.ZendStringFree(ret)
		return 0
	}

	/* We're using this method instead of == in order to provide
	 * resistance towards timing attacks. This is a constant time
	 * equality check that will always check every byte of both
	 * values. */

	for i = 0; i < hash.len_; i++ {
		status |= ret.val[i] ^ hash.val[i]
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
	var cost zend.ZendLong = 10
	if options != nil && g.Assign(&zcost, zend.ZendHashStrFind(options, "cost", g.SizeOf("\"cost\"")-1)) != nil {
		cost = zend.ZvalGetLong(zcost)
	}
	if cost < 4 || cost > 31 {
		core.PhpErrorDocref(nil, 1<<1, "Invalid bcrypt cost parameter specified: "+"%"+"lld", cost)
		return nil
	}
	hash_format_len = core.ApPhpSnprintf(hash_format, g.SizeOf("hash_format"), "$2y$%02"+"lld"+"$", cost)
	if !(g.Assign(&salt, PhpPasswordGetSalt(nil, 22, options))) {
		return nil
	}
	salt.val[salt.len_] = 0
	hash = zend.ZendStringAlloc(salt.len_+hash_format_len, 0)
	sprintf(hash.val, "%s%s", hash_format, salt.val)
	hash.val[hash_format_len+salt.len_] = 0
	zend.ZendStringReleaseEx(salt, 0)

	/* This cast is safe, since both values are defined here in code and cannot overflow */

	result = PhpCrypt(password.val, int(password.len_), hash.val, int(hash.len_), 1)
	zend.ZendStringReleaseEx(hash, 0)
	if result == nil {
		return nil
	}
	if result.len_ < 13 {
		zend.ZendStringFree(result)
		return nil
	}
	return result
}

var PhpPasswordAlgoBcrypt PhpPasswordAlgo = PhpPasswordAlgo{"bcrypt", PhpPasswordBcryptHash, PhpPasswordBcryptVerify, PhpPasswordBcryptNeedsRehash, PhpPasswordBcryptGetInfo, PhpPasswordBcryptValid}

func ZmStartupPassword(type_ int, module_number int) int {
	zend._zendHashInit(&PhpPasswordAlgos, 4, zend.ZvalPtrDtor, 1)
	zend.ZendRegisterStringConstant("PASSWORD_DEFAULT", g.SizeOf("\"PASSWORD_DEFAULT\"")-1, "2y", 1<<0|1<<1, module_number)
	if zend.FAILURE == PhpPasswordAlgoRegister("2y", &PhpPasswordAlgoBcrypt) {
		return zend.FAILURE
	}
	zend.ZendRegisterStringConstant("PASSWORD_BCRYPT", g.SizeOf("\"PASSWORD_BCRYPT\"")-1, "2y", 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("PASSWORD_BCRYPT_DEFAULT_COST", g.SizeOf("\"PASSWORD_BCRYPT_DEFAULT_COST\"")-1, 10, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownPassword(type_ int, module_number int) int {
	zend.ZendHashDestroy(&PhpPasswordAlgos)
	return zend.SUCCESS
}

/* }}} */

func PhpPasswordAlgoDefault() *PhpPasswordAlgo { return &PhpPasswordAlgoBcrypt }
func PhpPasswordAlgoFind(ident *zend.ZendString) *PhpPasswordAlgo {
	var tmp *zend.Zval
	if ident == nil {
		return nil
	}
	tmp = zend.ZendHashFind(&PhpPasswordAlgos, (*zend.ZendString)(ident))
	if tmp == nil || tmp.u1.v.type_ != 14 {
		return nil
	}
	return tmp.value.ptr
}
func PhpPasswordAlgoFindZvalEx(arg *zend.Zval, default_algo *PhpPasswordAlgo) *PhpPasswordAlgo {
	if arg == nil || arg.u1.v.type_ == 1 {
		return default_algo
	}
	if arg.u1.v.type_ == 4 {
		switch arg.value.lval {
		case 0:
			return default_algo
		case 1:
			return &PhpPasswordAlgoBcrypt
		case 2:
			var n *zend.ZendString = zend.ZendStringInit("argon2i", g.SizeOf("\"argon2i\"")-1, 0)
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			zend.ZendStringRelease(n)
			return ret
		case 3:
			var n *zend.ZendString = zend.ZendStringInit("argon2id", g.SizeOf("\"argon2id\"")-1, 0)
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			zend.ZendStringRelease(n)
			return ret
		}
		return nil
	}
	if arg.u1.v.type_ != 6 {
		return nil
	}
	return PhpPasswordAlgoFind(arg.value.str)
}
func PhpPasswordAlgoFindZval(arg *zend.Zval) *PhpPasswordAlgo {
	return PhpPasswordAlgoFindZvalEx(arg, PhpPasswordAlgoDefault())
}
func PhpPasswordAlgoExtractIdent(hash *zend.ZendString) *zend.ZendString {
	var ident *byte
	var ident_end *byte
	if hash == nil || hash.len_ < 3 {

		/* Minimum prefix: "$x$" */

		return nil

		/* Minimum prefix: "$x$" */

	}
	ident = hash.val + 1
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

/* {{{ proto array password_get_info(string $hash)
Retrieves information about a given hash */

func ZifPasswordGetInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var algo *PhpPasswordAlgo
	var hash *zend.ZendString
	var ident *zend.ZendString
	var options zend.Zval
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

			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = &options
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	ident = PhpPasswordAlgoExtractIdent(hash)
	algo = PhpPasswordAlgoFind(ident)
	if algo == nil || algo.GetValid() != nil && algo.GetValid()(hash) == 0 {
		if ident != nil {
			zend.ZendStringRelease(ident)
		}
		zend.AddAssocNullEx(return_value, "algo", strlen("algo"))
		zend.AddAssocStringEx(return_value, "algoName", strlen("algoName"), "unknown")
		zend.AddAssocZvalEx(return_value, "options", strlen("options"), &options)
		return
	}
	zend.AddAssocStrEx(return_value, "algo", strlen("algo"), PhpPasswordAlgoExtractIdent(hash))
	zend.ZendStringRelease(ident)
	zend.AddAssocStringEx(return_value, "algoName", strlen("algoName"), algo.GetName())
	if algo.GetGetInfo() != nil && zend.FAILURE == algo.GetGetInfo()(&options, hash) {
		zend.ZvalPtrDtorNogc(&options)
		zend.ZvalPtrDtorNogc(return_value)
		return_value.u1.type_info = 1
		return
	}
	zend.AddAssocZvalEx(return_value, "options", strlen("options"), &options)
}

/** }}} */

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

			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			zend.ZendParseArgZvalDeref(_arg, &znew_algo, 0)
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

			if zend.ZendParseArgArrayHt(_arg, &options, 0, 1, 0) == 0 {
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
			return
		}
		break
	}
	new_algo = PhpPasswordAlgoFindZval(znew_algo)
	if new_algo == nil {

		/* Unknown new algorithm, never prompt to rehash. */

		return_value.u1.type_info = 2
		return
	}
	old_algo = PhpPasswordAlgoIdentifyEx(hash, nil)
	if old_algo != new_algo {

		/* Different algorithm preferred, always rehash. */

		return_value.u1.type_info = 3
		return
	}
	if old_algo.GetNeedsRehash()(hash, options) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func ZifPasswordVerify(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var password *zend.ZendString
	var hash *zend.ZendString
	var algo *PhpPasswordAlgo
	for {
		var _flags int = 0
		var _min_num_args int = 2
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

			if zend.ZendParseArgStr(_arg, &password, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			if zend.ZendParseArgStr(_arg, &hash, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
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
	algo = PhpPasswordAlgoIdentify(hash)
	if algo != nil && (algo.GetVerify() == nil || algo.GetVerify()(password, hash) != 0) {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

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

			if zend.ZendParseArgStr(_arg, &password, 0) == 0 {
				_expected_type = zend.Z_EXPECTED_STRING
				_error_code = 4
				break
			}
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

			zend.ZendParseArgZvalDeref(_arg, &zalgo, 0)
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

			if zend.ZendParseArgArrayHt(_arg, &options, 0, 1, 0) == 0 {
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
			return
		}
		break
	}
	algo = PhpPasswordAlgoFindZval(zalgo)
	if algo == nil {
		var algostr *zend.ZendString = zend.ZvalGetString(zalgo)
		core.PhpErrorDocref(nil, 1<<1, "Unknown password hashing algorithm: %s", algostr.val)
		zend.ZendStringRelease(algostr)
		return_value.u1.type_info = 1
		return
	}
	digest = algo.GetHash()(password, options)
	if digest == nil {

		/* algo->hash should have raised an error. */

		return_value.u1.type_info = 1
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = digest
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func ZifPasswordAlgos(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var algo *zend.ZendString
	if execute_data.This.u2.num_args != 0 {
		zend.ZendWrongParametersNoneError()
		return
	}
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	for {
		var __ht *zend.HashTable = &PhpPasswordAlgos
		var _p *zend.Bucket = __ht.arData
		var _end *zend.Bucket = _p + __ht.nNumUsed
		for ; _p != _end; _p++ {
			var _z *zend.Zval = &_p.val

			if _z.u1.v.type_ == 0 {
				continue
			}
			algo = _p.key
			zend.AddNextIndexStr(return_value, zend.ZendStringCopy(algo))
		}
		break
	}
}

/* }}} */
