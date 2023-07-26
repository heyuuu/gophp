package standard

import (
	"crypto/subtle"
	"fmt"
	"github.com/heyuuu/gophp/core"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
)

func PhpPasswordAlgoRegister(ident string, algo *PhpPasswordAlgo) bool {
	zalgo := types.NewZvalPtr(algo)
	ret := PhpPasswordAlgos.KeyAdd(ident, zalgo)
	return ret != nil
}
func PhpPasswordSaltIsAlphabet(str string) bool {
	for _, c := range []byte(str) {
		if ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') || c == '.' || c == '/' {
			continue
		}
		return false
	}
	return true
}
func PhpPasswordSaltTo64(str string, outLen int) (string, bool) {
	buffer := PhpBase64Encode(str)
	if len(buffer) < outLen {
		/* Too short of an encoded string generated */
		return "", false
	}

	ret := strings.ReplaceAll("+", ".", buffer[:outLen])
	if pos := strings.IndexByte(ret, '='); pos >= 0 {
		return ret, false
	}
	return ret, true
}
func PhpPasswordMakeSalt(length int) *types.String {
	if length > core.INT_MAX/3 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length is too large to safely generate")
		return nil
	}
	buffer, ok := PhpRandomStringSafe(length*3/4 + 1)
	if !ok {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to generate salt")
		return nil
	}

	salt, ok := PhpPasswordSaltTo64(buffer, length)
	if !ok {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Generated salt too short")
		return nil
	}
	return types.NewString(salt)
}
func PhpPasswordGetSalt(required_salt_len int, options *types.Array) *types.String {
	var optionBuffer *types.Zval = nil
	if options != nil {
		optionBuffer = options.KeyFind("salt")
	}
	if optionBuffer == nil {
		return PhpPasswordMakeSalt(required_salt_len)
	}

	var buffer *types.String
	core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Use of the 'salt' option to password_hash is deprecated")
	switch optionBuffer.GetType() {
	case types.IS_STRING:
		buffer = optionBuffer.String().Copy()
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	case types.IS_OBJECT:
		buffer = operators.ZvalTryGetString(optionBuffer)
		if buffer == nil {
			return nil
		}
	case types.IS_FALSE:
		fallthrough
	case types.IS_TRUE:
		fallthrough
	case types.IS_NULL:
		fallthrough
	case types.IS_RESOURCE:
		fallthrough
	case types.IS_ARRAY:
		fallthrough
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Non-string salt parameter supplied")
		return nil
	}

	/* XXX all the crypt related APIs work with int for string length.
	   That should be revised for size_t and then we maybe don't require
	   the > INT_MAX check. */

	if zend.ZEND_SIZE_T_INT_OVFL(buffer.GetLen()) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Supplied salt is too long")
		return nil
	}
	if buffer.GetLen() < required_salt_len {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Provided salt is too short: %zd expecting %zd", buffer.GetLen(), required_salt_len)
		return nil
	}
	if !PhpPasswordSaltIsAlphabet(buffer.GetStr()) {
		salt, ok := PhpPasswordSaltTo64(buffer.GetStr(), required_salt_len)
		if !ok {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Provided salt is too short: %zd", buffer.GetLen())
			return nil
		}
		return types.NewString(salt)
	} else {
		salt := buffer.GetStr()[:required_salt_len]
		return types.NewString(salt)
	}
}
func PhpPasswordBcryptValid(hash string) bool {
	return len(hash) == 60 && hash[:3] == "$2y"
}
func PhpPasswordBcryptGetInfo(hash string) *types.Array {
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if !PhpPasswordBcryptValid(hash) {
		/* Should never get called this way. */
		return nil
	}

	sscanf(hash, "$2y$"+zend.ZEND_LONG_FMT+"$", &cost)
	arr := types.NewArray(0)
	arr.KeyAdd("cost", types.NewZvalLong(cost))
	return arr
}
func PhpPasswordBcryptNeedsRehash(hash string, options *types.Array) bool {
	var znew_cost *types.Zval
	var old_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	var new_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if !PhpPasswordBcryptValid(hash) {
		/* Should never get called this way. */
		return true
	}
	sscanf(hash, "$2y$"+zend.ZEND_LONG_FMT+"$", &old_cost)
	if options != nil && b.Assign(&znew_cost, options.KeyFind("cost")) != nil {
		new_cost = operators.ZvalGetLong(znew_cost)
	}
	return old_cost != new_cost
}
func PhpPasswordBcryptVerify(password string, hash string) bool {
	var ret *types.String = PhpCrypt(password, hash, true)
	if ret == nil {
		return false
	}
	retStr := ret.GetStr()
	if len(retStr) != len(hash) || len(hash) < 13 {
		return false
	}

	/* We're using this method instead of == in order to provide
	 * resistance towards timing attacks. This is a constant time
	 * equality check that will always check every byte of both
	 * values. */
	return subtle.ConstantTimeCompare([]byte(retStr), []byte(hash)) != 0
}
func PhpPasswordBcryptHash(password string, options *types.Array) (string, bool) {
	var result *types.String
	var zcost *types.Zval
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if options != nil && b.Assign(&zcost, options.KeyFind("cost")) != nil {
		cost = operators.ZvalGetLong(zcost)
	}
	if cost < 4 || cost > 31 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid bcrypt cost parameter specified: "+zend.ZEND_LONG_FMT, cost)
		return "", false
	}
	hashFormat := fmt.Sprintf("$2y$%02d$", cost)

	salt := PhpPasswordGetSalt(22, options)
	if salt == nil {
		return "", false
	}

	hash := hashFormat + salt.GetStr()

	/* This cast is safe, since both values are defined here in code and cannot overflow */
	result = PhpCrypt(password, hash, true)
	if result == nil {
		return "", false
	}
	if result.GetLen() < 13 {
		return "", false
	}
	return result.GetStr(), true
}
func ZmStartupPassword(type_ int, module_number int) int {
	PhpPasswordAlgos = types.NewArray(0)
	if !PhpPasswordAlgoRegister("2y", passwordAlgoBcrypt) {
		return types.FAILURE
	}

	zend.RegisterStringConstant("PASSWORD_DEFAULT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterStringConstant("PASSWORD_BCRYPT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PASSWORD_BCRYPT_DEFAULT_COST", PHP_PASSWORD_BCRYPT_COST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}
func ZmShutdownPassword(type_ int, module_number int) int {
	PhpPasswordAlgos.Destroy()
	return types.SUCCESS
}
func PhpPasswordAlgoDefault() *PhpPasswordAlgo          { return passwordAlgoBcrypt }
func PhpPasswordAlgoFind(ident string) *PhpPasswordAlgo { return passwordAlgos[ident] }
func PhpPasswordAlgoFindZvalEx(arg *types.Zval, defaultAlgo *PhpPasswordAlgo) *PhpPasswordAlgo {
	if arg == nil || arg.IsNull() {
		return defaultAlgo
	}
	if arg.IsLong() {
		switch arg.Long() {
		case 0:
			return defaultAlgo
		case 1:
			return passwordAlgoBcrypt
		case 2:
			return PhpPasswordAlgoFind("argon2i")
		case 3:
			return PhpPasswordAlgoFind("argon2id")
		}
		return nil
	}
	if !arg.IsString() {
		return nil
	}
	return PhpPasswordAlgoFind(arg.StringVal())
}
func PhpPasswordAlgoFindZval(arg *types.Zval) *PhpPasswordAlgo {
	return PhpPasswordAlgoFindZvalEx(arg, PhpPasswordAlgoDefault())
}
func PhpPasswordAlgoExtractIdent(hash string) *types.String {
	if len(hash) < 3 {
		/* Minimum prefix: "$x$" */
		return nil
	}

	ident := hash[1:]
	if pos := strings.IndexByte(ident, '$'); pos >= 0 {
		return types.NewString(ident[:pos])
	} else {
		/* No terminating '$' */
		return nil
	}
}

func PhpPasswordAlgoIdentify(hash string) *PhpPasswordAlgo {
	return PhpPasswordAlgoIdentifyEx(hash, PhpPasswordAlgoDefault())
}
func PhpPasswordAlgoIdentifyEx(hash string, defaultAlgo *PhpPasswordAlgo) *PhpPasswordAlgo {
	var algo *PhpPasswordAlgo
	ident := PhpPasswordAlgoExtractIdent(hash)
	if ident == nil {
		return defaultAlgo
	}
	algo = PhpPasswordAlgoFind(ident.GetStr())
	if algo == nil || !algo.Valid(hash) {
		return defaultAlgo
	} else {
		return algo
	}
}
func ZifPasswordGetInfo(hash_ string) *types.Zval {
	var algo *PhpPasswordAlgo = nil
	if ident := PhpPasswordAlgoExtractIdent(hash_); ident != nil {
		algo = PhpPasswordAlgoFind(ident.GetStr())
	}
	if algo == nil || !algo.Valid(hash_) {
		arr := types.NewArray(0)
		arr.KeyAdd("algo", types.NewZvalNull())
		arr.KeyAdd("algoName", types.NewZvalString("unknown"))
		arr.KeyAdd("options", types.NewZvalEmptyArray())
		return types.NewZvalArray(arr)
	} else if options_ := algo.GetInfo(hash_); options_ != nil {
		arr := types.NewArray(0)
		arr.KeyAdd("algo", types.NewZvalString(PhpPasswordAlgoExtractIdent(hash_).GetStr()))
		arr.KeyAdd("algoName", types.NewZvalString(algo.Name()))
		arr.KeyAdd("options", types.NewZvalArray(options_))
		return types.NewZvalArray(arr)
	} else {
		return types.NewZvalNull()
	}
}
func ZifPasswordNeedsRehash(hash_ string, algo_ *types.Zval, _ zpp.Opt, options zpp.ArrayOrObjectHt) bool {
	newAlgo := PhpPasswordAlgoFindZval(algo_)
	if newAlgo == nil {
		/* Unknown new algorithm, never prompt to rehash. */
		return false
	}

	oldAlgo := PhpPasswordAlgoIdentifyEx(hash_, nil)
	if oldAlgo != newAlgo {
		return true
	}

	return oldAlgo.NeedsRehash(hash_, options)
}
func ZifPasswordVerify(password string, hash string) bool {
	algo := PhpPasswordAlgoIdentify(hash)
	return algo != nil && algo.Verify(password, hash)
}
func ZifPasswordHash(password string, algo_ *types.Zval, _ zpp.Opt, options zpp.ArrayOrObjectHt) *types.Zval {
	var algo = PhpPasswordAlgoFindZval(algo_)
	if algo == nil {
		var algoStr = operators.ZvalGetStrVal(algo_)
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown password hashing algorithm: %s", algoStr)
		return types.NewZvalNull()
	}
	if digest, ok := algo.Hash(password, options); ok {
		return types.NewZvalString(digest)
	} else {
		/* algo->hash should have raised an error. */
		return types.NewZvalNull()
	}
}
func ZifPasswordAlgos() []string {
	var algoNames []string
	PhpPasswordAlgos.Foreach(func(key types.ArrayKey, _ *types.Zval) {
		algoNames = append(algoNames, key.StrKey())
	})
	return algoNames
}
