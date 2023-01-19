// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/crypt.c>

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
   | Authors: Stig Bakken <ssb@php.net>                                   |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Rasmus Lerdorf <rasmus@php.net>                             |
   |          Pierre Joye <pierre@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include < stdlib . h >

// # include "php.h"

// # include < unistd . h >

// # include "php_crypt_r.h"

// # include "crypt_freesec.h"

// # include < time . h >

// # include < string . h >

// # include "php_crypt.h"

// # include "php_random.h"

/* sha512 crypt has the maximal salt length of 123 characters */

// #define PHP_MAX_SALT_LEN       123

/* Used to check DES salts to ensure that they contain only valid characters */

// #define IS_VALID_SALT_CHARACTER(c) ( ( ( c ) >= '.' && ( c ) <= '9' ) || ( ( c ) >= 'A' && ( c ) <= 'Z' ) || ( ( c ) >= 'a' && ( c ) <= 'z' ) )

// #define DES_INVALID_SALT_ERROR       "Supplied salt is not valid for DES. Possible bug in provided salt format."

func ZmStartupCrypt(type_ int, module_number int) int {
	zend.ZendRegisterLongConstant("CRYPT_SALT_LENGTH", g.SizeOf("\"CRYPT_SALT_LENGTH\"")-1, 123, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_STD_DES", g.SizeOf("\"CRYPT_STD_DES\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_EXT_DES", g.SizeOf("\"CRYPT_EXT_DES\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_MD5", g.SizeOf("\"CRYPT_MD5\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_BLOWFISH", g.SizeOf("\"CRYPT_BLOWFISH\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_SHA256", g.SizeOf("\"CRYPT_SHA256\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("CRYPT_SHA512", g.SizeOf("\"CRYPT_SHA512\"")-1, 1, 1<<0|1<<1, module_number)
	PhpInitCryptR()
	return zend.SUCCESS
}

/* }}} */

func ZmShutdownCrypt(type_ int, module_number int) int {
	PhpShutdownCryptR()
	return zend.SUCCESS
}

/* }}} */

var Itoa64 []uint8 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func PhpTo64(s *byte, n int) {
	for g.PreDec(&n) >= 0 {
		*s = Itoa64[(*s)&0x3f]
		s++
	}
}

/* }}} */

func PhpCrypt(password *byte, pass_len int, salt *byte, salt_len int, quiet zend.ZendBool) *zend.ZendString {
	var crypt_res *byte
	var result *zend.ZendString
	if salt[0] == '*' && (salt[1] == '0' || salt[1] == '1') {
		return nil
	}

	/* Windows (win32/crypt) has a stripped down version of libxcrypt and
	   a CryptoApi md5_crypt implementation */

	var buffer PhpCryptExtendedData
	if salt[0] == '$' && salt[1] == '1' && salt[2] == '$' {
		var output []byte
		var out *byte
		out = PhpMd5CryptR(password, salt, output)
		if out != nil {
			return zend.ZendStringInit(out, strlen(out), 0)
		}
		return nil
	} else if salt[0] == '$' && salt[1] == '6' && salt[2] == '$' {
		var output *byte
		output = zend._emalloc(123)
		crypt_res = PhpSha512CryptR(password, salt, output, 123)
		if crypt_res == nil {
			core.PhpExplicitBzero(output, 123)
			zend._efree(output)
			return nil
		} else {
			result = zend.ZendStringInit(output, strlen(output), 0)
			core.PhpExplicitBzero(output, 123)
			zend._efree(output)
			return result
		}
	} else if salt[0] == '$' && salt[1] == '5' && salt[2] == '$' {
		var output *byte
		output = zend._emalloc(123)
		crypt_res = PhpSha256CryptR(password, salt, output, 123)
		if crypt_res == nil {
			core.PhpExplicitBzero(output, 123)
			zend._efree(output)
			return nil
		} else {
			result = zend.ZendStringInit(output, strlen(output), 0)
			core.PhpExplicitBzero(output, 123)
			zend._efree(output)
			return result
		}
	} else if salt[0] == '$' && salt[1] == '2' && salt[3] == '$' {
		var output []byte
		memset(output, 0, 123+1)
		crypt_res = PhpCryptBlowfishRn(password, salt, output, g.SizeOf("output"))
		if crypt_res == nil {
			core.PhpExplicitBzero(output, 123+1)
			return nil
		} else {
			result = zend.ZendStringInit(output, strlen(output), 0)
			core.PhpExplicitBzero(output, 123+1)
			return result
		}
	} else {

		/* DES Fallback */

		if salt[0] != '_' {

			/* DES style hashes */

			if !(salt[0] >= '.' && salt[0] <= '9' || salt[0] >= 'A' && salt[0] <= 'Z' || salt[0] >= 'a' && salt[0] <= 'z') || !(salt[1] >= '.' && salt[1] <= '9' || salt[1] >= 'A' && salt[1] <= 'Z' || salt[1] >= 'a' && salt[1] <= 'z') {
				if quiet == 0 {

					/* error consistently about invalid DES fallbacks */

					core.PhpErrorDocref(nil, 1<<13, "Supplied salt is not valid for DES. Possible bug in provided salt format.")

					/* error consistently about invalid DES fallbacks */

				}
			}

			/* DES style hashes */

		}
		memset(&buffer, 0, g.SizeOf("buffer"))
		_cryptExtendedInitR()
		crypt_res = _cryptExtendedR((*uint8)(password), salt, &buffer)
		if crypt_res == nil || salt[0] == '*' && salt[1] == '0' {
			return nil
		} else {
			result = zend.ZendStringInit(crypt_res, strlen(crypt_res), 0)
			return result
		}
	}
	if crypt_res == nil || salt[0] == '*' && salt[1] == '0' {
		return nil
	} else {
		result = zend.ZendStringInit(crypt_res, strlen(crypt_res), 0)
		return result
	}
}

/* }}} */

func ZifCrypt(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var salt []byte
	var str *byte
	var salt_in *byte = nil
	var str_len int
	var salt_in_len int = 0
	var result *zend.ZendString
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

			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
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

			if zend.ZendParseArgString(_arg, &salt_in, &salt_in_len, 0) == 0 {
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
	salt[123] = '0'
	salt[0] = salt[123]

	/* This will produce suitable results if people depend on DES-encryption
	 * available (passing always 2-character salt). At least for glibc6.1 */

	memset(&salt[1], '$', 123-1)
	if salt_in != nil {
		memcpy(salt, salt_in, g.Cond(123 < salt_in_len, 123, salt_in_len))
	} else {
		core.PhpErrorDocref(nil, 1<<3, "No salt parameter was specified. You must use a randomly generated salt and a strong hash function to produce a secure hash.")
	}

	/* The automatic salt generation covers standard DES, md5-crypt and Blowfish (simple) */

	if !(*salt) {
		memcpy(salt, "$1$", 3)
		PhpRandomBytes(&salt[3], 8, 1)
		PhpTo64(&salt[3], 8)
		strncpy(&salt[11], "$", 123-11)
		salt_in_len = strlen(salt)
	} else {
		if 123 < salt_in_len {
			salt_in_len = 123
		} else {
			salt_in_len = salt_in_len
		}
	}
	salt[salt_in_len] = '0'
	if g.Assign(&result, PhpCrypt(str, int(str_len), salt, int(salt_in_len), 0)) == nil {
		if salt[0] == '*' && salt[1] == '0' {
			var _s *byte = "*1"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		} else {
			var _s *byte = "*0"
			var __z *zend.Zval = return_value
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return
		}
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = result
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}} */
