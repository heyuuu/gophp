package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
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

func PhpTo64Ex(s string) string {
	var bytes = []byte(s)
	for i, c := range bytes {
		bytes[i] = Itoa64[c&0x3f]
	}
	return string(bytes)
}

func PhpCrypt(password string, salt string, quiet bool) *types.String {
	var crypt_res *byte
	var result *types.String
	if salt[0] == '*' && (salt[1] == '0' || salt[1] == '1') {
		return nil
	}

	/* Windows (win32/crypt) has a stripped down version of libxcrypt and
	   a CryptoApi md5_crypt implementation */

	var buffer PhpCryptExtendedData
	if salt[0] == '$' && salt[1] == '1' && salt[2] == '$' {
		out := PhpMd5CryptR(password, salt)
		return types.NewString(out)
	} else if salt[0] == '$' && salt[1] == '6' && salt[2] == '$' {
		var output *byte
		output = zend.Emalloc(PHP_MAX_SALT_LEN)
		crypt_res = PhpSha512CryptR(password, salt, output, PHP_MAX_SALT_LEN)
		if crypt_res == nil {
			zend.ZEND_SECURE_ZERO(utput, PHP_MAX_SALT_LEN)
			zend.Efree(output)
			return nil
		} else {
			result = types.NewString(output)
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
			result = types.NewString(output)
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
			result = types.NewString(output)
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
			result = types.NewString(crypt_res)
			return result
		}
	}
}
func ZifCrypt(str_ string, _ zpp.Opt, salt_ string) string {
	var result *types.String

	if len(salt_) > PHP_MAX_SALT_LEN {
		salt_ = salt_[:PHP_MAX_SALT_LEN]
	}

	var realSalt string
	/* This will produce suitable results if people depend on DES-encryption
	 * available (passing always 2-character salt). At least for glibc6.1 */
	if salt_ != "" {
		realSalt = salt_ + strings.Repeat("$", PHP_MAX_SALT_LEN-len(salt_))
	} else {
		core.PhpErrorDocref(nil, faults.E_NOTICE, "No salt parameter was specified. You must use a randomly generated salt and a strong hash function to produce a secure hash.")
	}

	/* The automatic salt generation covers standard DES, md5-crypt and Blowfish (simple) */
	if realSalt == "" || realSalt[0] == '\000' {
		randStr, _ := PhpRandomStringSafe(8)
		realSalt = "$1$" + PhpTo64Ex(randStr) + strings.Repeat("$", PHP_MAX_SALT_LEN-11)
	}

	result = PhpCrypt(str_, salt_, false)
	if result == nil {
		if realSalt[:2] == "*0" {
			return "*1"
		} else {
			return "*0"
		}
	}
	return result.GetStr()
}
