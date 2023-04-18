package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func IS_VALID_SALT_CHARACTER(c byte) bool {
	return c >= '.' && c <= '9' || c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z'
}
func ZmStartupCrypt(type_ int, module_number int) int {
	zend.RegisterLongConstant("CRYPT_SALT_LENGTH", PHP_MAX_SALT_LEN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_STD_DES", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_EXT_DES", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_MD5", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_BLOWFISH", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_SHA256", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("CRYPT_SHA512", 1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	PhpInitCryptR()
	return types2.SUCCESS
}
func ZmShutdownCrypt(type_ int, module_number int) int {
	PhpShutdownCryptR()
	return types2.SUCCESS
}
func PhpTo64(s *byte, n int) {
	for b.PreDec(&n) >= 0 {
		*s = Itoa64[(*s)&0x3f]
		s++
	}
}

func PhpCrypt(password string, salt string, quiet bool) *types2.String {
	var crypt_res *byte
	var result *types2.String
	if salt[0] == '*' && (salt[1] == '0' || salt[1] == '1') {
		return nil
	}

	/* Windows (win32/crypt) has a stripped down version of libxcrypt and
	   a CryptoApi md5_crypt implementation */

	var buffer PhpCryptExtendedData
	if salt[0] == '$' && salt[1] == '1' && salt[2] == '$' {
		out := PhpMd5CryptR(password, salt)
		return types2.NewString(out)
	} else if salt[0] == '$' && salt[1] == '6' && salt[2] == '$' {
		var output *byte
		output = zend.Emalloc(PHP_MAX_SALT_LEN)
		crypt_res = PhpSha512CryptR(password, salt, output, PHP_MAX_SALT_LEN)
		if crypt_res == nil {
			zend.ZEND_SECURE_ZERO(utput, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return nil
		} else {
			result = types2.NewString(output)
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
			result = types2.NewString(output)
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
			result = types2.NewString(output)
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
					core.PhpErrorDocref(nil, faults.E_DEPRECATED, DES_INVALID_SALT_ERROR)
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
			result = types2.NewString(crypt_res)
			return result
		}
	}
	if crypt_res == nil || salt[0] == '*' && salt[1] == '0' {
		return nil
	} else {
		result = types2.NewString(crypt_res)
		return result
	}
}
func ZifCrypt(executeData zpp.Ex, return_value zpp.Ret, str_ string, _ zpp.Opt, salt_ string) string {
	var salt []byte
	var str *byte
	var salt_in *byte = nil
	var str_len int
	var salt_in_len int = 0
	var result *types2.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 2, 0)
			str, str_len = fp.ParseString()
			fp.StartOptional()
			salt_in, salt_in_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	salt[PHP_MAX_SALT_LEN] = '\000'
	salt[0] = '\000'

	/* This will produce suitable results if people depend on DES-encryption
	 * available (passing always 2-character salt). At least for glibc6.1 */

	memset(&salt[1], '$', PHP_MAX_SALT_LEN-1)
	if salt_in != nil {
		memcpy(salt, salt_in, cli.MIN(PHP_MAX_SALT_LEN, salt_in_len))
	} else {
		core.PhpErrorDocref(nil, faults.E_NOTICE, "No salt parameter was specified. You must use a randomly generated salt and a strong hash function to produce a secure hash.")
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

	result = PhpCrypt(b.CastStr(str, str_len), b.CastStr(salt, salt_in_len), false)
	if result == nil {
		if salt[0] == '*' && salt[1] == '0' {
			return "*1"
		} else {
			return "*0"
		}
	}
	return result.GetStr()
}
