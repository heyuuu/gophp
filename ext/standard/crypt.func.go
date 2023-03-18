// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func IS_VALID_SALT_CHARACTER(c byte) bool {
	return c >= '.' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}
func ZmStartupCrypt(type_ int, module_number int) int {
	zend.REGISTER_LONG_CONSTANT("CRYPT_SALT_LENGTH", PHP_MAX_SALT_LEN, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_STD_DES", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_EXT_DES", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_MD5", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_BLOWFISH", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_SHA256", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("CRYPT_SHA512", 1, zend.CONST_CS|zend.CONST_PERSISTENT)
	PhpInitCryptR()
	return types.SUCCESS
}
func ZmShutdownCrypt(type_ int, module_number int) int {
	PhpShutdownCryptR()
	return types.SUCCESS
}
func PhpTo64(s *byte, n int) {
	for b.PreDec(&n) >= 0 {
		*s = Itoa64[(*s)&0x3f]
		s++
	}
}
func PhpCrypt(password *byte, pass_len int, salt *byte, salt_len int, quiet types.ZendBool) *types.ZendString {
	var crypt_res *byte
	var result *types.ZendString
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
			return types.ZendStringInit(out, strlen(out), 0)
		}
		return nil
	} else if salt[0] == '$' && salt[1] == '6' && salt[2] == '$' {
		var output *byte
		output = zend.Emalloc(PHP_MAX_SALT_LEN)
		crypt_res = PhpSha512CryptR(password, salt, output, PHP_MAX_SALT_LEN)
		if crypt_res == nil {
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return nil
		} else {
			result = types.ZendStringInit(output, strlen(output), 0)
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return result
		}
	} else if salt[0] == '$' && salt[1] == '5' && salt[2] == '$' {
		var output *byte
		output = zend.Emalloc(PHP_MAX_SALT_LEN)
		crypt_res = PhpSha256CryptR(password, salt, output, PHP_MAX_SALT_LEN)
		if crypt_res == nil {
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return nil
		} else {
			result = types.ZendStringInit(output, strlen(output), 0)
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return result
		}
	} else if salt[0] == '$' && salt[1] == '2' && salt[3] == '$' {
		var output []byte
		memset(output, 0, PHP_MAX_SALT_LEN+1)
		crypt_res = PhpCryptBlowfishRn(password, salt, output, b.SizeOf("output"))
		if crypt_res == nil {
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN+1)
			return nil
		} else {
			result = types.ZendStringInit(output, strlen(output), 0)
			zend.ZEND_SECURE_ZERO(output, PHP_MAX_SALT_LEN+1)
			return result
		}
	} else {

		/* DES Fallback */

		if salt[0] != '_' {

			/* DES style hashes */

			if !(IS_VALID_SALT_CHARACTER(salt[0])) || !(IS_VALID_SALT_CHARACTER(salt[1])) {
				if quiet == 0 {

					/* error consistently about invalid DES fallbacks */

					core.PhpErrorDocref(nil, zend.E_DEPRECATED, DES_INVALID_SALT_ERROR)

					/* error consistently about invalid DES fallbacks */

				}
			}

			/* DES style hashes */

		}
		memset(&buffer, 0, b.SizeOf("buffer"))
		_cryptExtendedInitR()
		crypt_res = _cryptExtendedR((*uint8)(password), salt, &buffer)
		if crypt_res == nil || salt[0] == '*' && salt[1] == '0' {
			return nil
		} else {
			result = types.ZendStringInit(crypt_res, strlen(crypt_res), 0)
			return result
		}
	}
	if crypt_res == nil || salt[0] == '*' && salt[1] == '0' {
		return nil
	} else {
		result = types.ZendStringInit(crypt_res, strlen(crypt_res), 0)
		return result
	}
}
func ZifCrypt(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var salt []byte
	var str *byte
	var salt_in *byte = nil
	var str_len int
	var salt_in_len int = 0
	var result *types.ZendString
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &str, &str_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgString(_arg, &salt_in, &salt_in_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_STRING
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	salt[PHP_MAX_SALT_LEN] = '0'
	salt[0] = salt[PHP_MAX_SALT_LEN]

	/* This will produce suitable results if people depend on DES-encryption
	 * available (passing always 2-character salt). At least for glibc6.1 */

	memset(&salt[1], '$', PHP_MAX_SALT_LEN-1)
	if salt_in != nil {
		memcpy(salt, salt_in, cli.MIN(PHP_MAX_SALT_LEN, salt_in_len))
	} else {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "No salt parameter was specified. You must use a randomly generated salt and a strong hash function to produce a secure hash.")
	}

	/* The automatic salt generation covers standard DES, md5-crypt and Blowfish (simple) */

	if !(*salt) {
		memcpy(salt, "$1$", 3)
		PhpRandomBytesThrow(&salt[3], 8)
		PhpTo64(&salt[3], 8)
		strncpy(&salt[11], "$", PHP_MAX_SALT_LEN-11)
		salt_in_len = strlen(salt)
	} else {
		salt_in_len = cli.MIN(PHP_MAX_SALT_LEN, salt_in_len)
	}
	salt[salt_in_len] = '0'
	if b.Assign(&result, PhpCrypt(str, int(str_len), salt, int(salt_in_len), 0)) == nil {
		if salt[0] == '*' && salt[1] == '0' {
			return_value.SetRawString("*1")
			return
		} else {
			return_value.SetRawString("*0")
			return
		}
	}
	return_value.SetString(result)
	return
}
