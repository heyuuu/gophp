// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
	"sik/zend/zpp"
)

func PhpPasswordAlgoRegister(ident string, algo *PhpPasswordAlgo) int {
	var zalgo types.Zval
	zalgo.SetAsPtr((*PhpPasswordAlgo)(algo))
	if PhpPasswordAlgos.KeyAdd(b.CastStrAuto(ident), &zalgo) != nil {
		return types.SUCCESS
	}
	return types.FAILURE
}
func PhpPasswordAlgoUnregister(ident *byte) {
	types.ZendHashStrDel(&PhpPasswordAlgos, ident)
}
func PhpPasswordSaltIsAlphabet(str *byte, len_ int) int {
	var i int = 0
	for i = 0; i < len_; i++ {
		if !(str[i] >= 'A' && str[i] <= 'Z' || str[i] >= 'a' && str[i] <= 'z' || str[i] >= '0' && str[i] <= '9' || str[i] == '.' || str[i] == '/') {
			return types.FAILURE
		}
	}
	return types.SUCCESS
}
func PhpPasswordSaltTo64(str *byte, str_len int, out_len int, ret *byte) int {
	var pos int = 0
	var buffer *types.String
	if int(str_len < 0) != 0 {
		return types.FAILURE
	}
	buffer = PhpBase64Encode((*uint8)(str), str_len)
	if buffer.GetLen() < out_len {

		/* Too short of an encoded string generated */

		types.ZendStringReleaseEx(buffer, 0)
		return types.FAILURE
	}
	for pos = 0; pos < out_len; pos++ {
		if buffer.GetVal()[pos] == '+' {
			ret[pos] = '.'
		} else if buffer.GetVal()[pos] == '=' {
			types.ZendStringFree(buffer)
			return types.FAILURE
		} else {
			ret[pos] = buffer.GetVal()[pos]
		}
	}
	types.ZendStringFree(buffer)
	return types.SUCCESS
}
func PhpPasswordMakeSalt(length int) *types.String {
	var ret *types.String
	var buffer *types.String
	if length > core.INT_MAX/3 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length is too large to safely generate")
		return nil
	}
	buffer = types.ZendStringAlloc(length*3/4+1, 0)
	if types.FAILURE == PhpRandomBytesSilent(buffer.GetVal(), buffer.GetLen()) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to generate salt")
		types.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	ret = types.ZendStringAlloc(length, 0)
	if PhpPasswordSaltTo64(buffer.GetVal(), buffer.GetLen(), length, ret.GetVal()) == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Generated salt too short")
		types.ZendStringReleaseEx(buffer, 0)
		types.ZendStringReleaseEx(ret, 0)
		return nil
	}
	types.ZendStringReleaseEx(buffer, 0)
	ret.GetVal()[length] = 0
	return ret
}
func PhpPasswordGetSalt(unused_ *types.Zval, required_salt_len int, options *types.Array) *types.String {
	var buffer *types.String
	var option_buffer *types.Zval
	if options == nil || !(b.Assign(&option_buffer, options.KeyFind("salt"))) {
		return PhpPasswordMakeSalt(required_salt_len)
	}
	core.PhpErrorDocref(nil, faults.E_DEPRECATED, "Use of the 'salt' option to password_hash is deprecated")
	switch option_buffer.GetType() {
	case types.IS_STRING:
		buffer = option_buffer.GetStr().Copy()
	case types.IS_LONG:
		fallthrough
	case types.IS_DOUBLE:
		fallthrough
	case types.IS_OBJECT:
		buffer = zend.ZvalTryGetString(option_buffer)
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
		types.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if buffer.GetLen() < required_salt_len {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Provided salt is too short: %zd expecting %zd", buffer.GetLen(), required_salt_len)
		types.ZendStringReleaseEx(buffer, 0)
		return nil
	}
	if PhpPasswordSaltIsAlphabet(buffer.GetVal(), buffer.GetLen()) == types.FAILURE {
		var salt *types.String = types.ZendStringAlloc(required_salt_len, 0)
		if PhpPasswordSaltTo64(buffer.GetVal(), buffer.GetLen(), required_salt_len, salt.GetVal()) == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Provided salt is too short: %zd", buffer.GetLen())
			types.ZendStringReleaseEx(salt, 0)
			types.ZendStringReleaseEx(buffer, 0)
			return nil
		}
		types.ZendStringReleaseEx(buffer, 0)
		return salt
	} else {
		var salt *types.String = types.ZendStringAlloc(required_salt_len, 0)
		memcpy(salt.GetVal(), buffer.GetVal(), required_salt_len)
		types.ZendStringReleaseEx(buffer, 0)
		return salt
	}
}
func PhpPasswordBcryptValid(hash *types.String) types.ZendBool {
	var h *byte = hash.GetVal()
	return hash.GetLen() == 60 && h[0] == '$' && h[1] == '2' && h[2] == 'y'
}
func PhpPasswordBcryptGetInfo(return_value *types.Zval, hash *types.String) int {
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return types.FAILURE

		/* Should never get called this way. */

	}
	sscanf(hash.GetVal(), "$2y$"+zend.ZEND_LONG_FMT+"$", &cost)
	zend.AddAssocLong(return_value, "cost", cost)
	return types.SUCCESS
}
func PhpPasswordBcryptNeedsRehash(hash *types.String, options *types.Array) types.ZendBool {
	var znew_cost *types.Zval
	var old_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	var new_cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if PhpPasswordBcryptValid(hash) == 0 {

		/* Should never get called this way. */

		return 1

		/* Should never get called this way. */

	}
	sscanf(hash.GetVal(), "$2y$"+zend.ZEND_LONG_FMT+"$", &old_cost)
	if options != nil && b.Assign(&znew_cost, options.KeyFind("cost")) != nil {
		new_cost = zend.ZvalGetLong(znew_cost)
	}
	return old_cost != new_cost
}
func PhpPasswordBcryptVerify(password *types.String, hash *types.String) types.ZendBool {
	var i int
	var status int = 0
	var ret *types.String = PhpCrypt(password.GetVal(), int(password.GetLen()), hash.GetVal(), int(hash.GetLen()), 1)
	if ret == nil {
		return 0
	}
	if ret.GetLen() != hash.GetLen() || hash.GetLen() < 13 {
		types.ZendStringFree(ret)
		return 0
	}

	/* We're using this method instead of == in order to provide
	 * resistance towards timing attacks. This is a constant time
	 * equality check that will always check every byte of both
	 * values. */

	for i = 0; i < hash.GetLen(); i++ {
		status |= ret.GetVal()[i] ^ hash.GetVal()[i]
	}
	types.ZendStringFree(ret)
	return status == 0
}
func PhpPasswordBcryptHash(password *types.String, options *types.Array) *types.String {
	var hash_format []byte
	var hash_format_len int
	var result *types.String
	var hash *types.String
	var salt *types.String
	var zcost *types.Zval
	var cost zend.ZendLong = PHP_PASSWORD_BCRYPT_COST
	if options != nil && b.Assign(&zcost, options.KeyFind("cost")) != nil {
		cost = zend.ZvalGetLong(zcost)
	}
	if cost < 4 || cost > 31 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid bcrypt cost parameter specified: "+zend.ZEND_LONG_FMT, cost)
		return nil
	}
	hash_format_len = core.Snprintf(hash_format, b.SizeOf("hash_format"), "$2y$%02"+zend.ZEND_LONG_FMT_SPEC+"$", cost)
	if !(b.Assign(&salt, PhpPasswordGetSalt(nil, uint64(22), options))) {
		return nil
	}
	salt.GetVal()[salt.GetLen()] = 0
	hash = types.ZendStringAlloc(salt.GetLen()+hash_format_len, 0)
	sprintf(hash.GetVal(), "%s%s", hash_format, salt.GetVal())
	hash.GetVal()[hash_format_len+salt.GetLen()] = 0
	types.ZendStringReleaseEx(salt, 0)

	/* This cast is safe, since both values are defined here in code and cannot overflow */

	result = PhpCrypt(password.GetVal(), int(password.GetLen()), hash.GetVal(), int(hash.GetLen()), 1)
	types.ZendStringReleaseEx(hash, 0)
	if result == nil {
		return nil
	}
	if result.GetLen() < 13 {
		types.ZendStringFree(result)
		return nil
	}
	return result
}
func ZmStartupPassword(type_ int, module_number int) int {
	&PhpPasswordAlgos = types.MakeArrayEx(4, zend.ZVAL_PTR_DTOR, 1)
	zend.RegisterStringConstant("PASSWORD_DEFAULT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	if types.FAILURE == PhpPasswordAlgoRegister("2y", &PhpPasswordAlgoBcrypt) {
		return types.FAILURE
	}
	zend.RegisterStringConstant("PASSWORD_BCRYPT", "2y", zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("PASSWORD_BCRYPT_DEFAULT_COST", PHP_PASSWORD_BCRYPT_COST, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	return types.SUCCESS
}
func ZmShutdownPassword(type_ int, module_number int) int {
	PhpPasswordAlgos.Destroy()
	return types.SUCCESS
}
func PhpPasswordAlgoDefault() *PhpPasswordAlgo { return &PhpPasswordAlgoBcrypt }
func PhpPasswordAlgoFind(ident *types.String) *PhpPasswordAlgo {
	var tmp *types.Zval
	if ident == nil {
		return nil
	}
	tmp = PhpPasswordAlgos.KeyFind((*types.String)(ident).GetStr())
	if tmp == nil || tmp.GetType() != types.IS_PTR {
		return nil
	}
	return tmp.GetPtr()
}
func PhpPasswordAlgoFindZvalEx(arg *types.Zval, default_algo *PhpPasswordAlgo) *PhpPasswordAlgo {
	if arg == nil || arg.IsType(types.IS_NULL) {
		return default_algo
	}
	if arg.IsType(types.IS_LONG) {
		switch arg.GetLval() {
		case 0:
			return default_algo
		case 1:
			return &PhpPasswordAlgoBcrypt
		case 2:
			var n *types.String = types.NewString("argon2i")
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			types.ZendStringRelease(n)
			return ret
		case 3:
			var n *types.String = types.NewString("argon2id")
			var ret *PhpPasswordAlgo = PhpPasswordAlgoFind(n)
			types.ZendStringRelease(n)
			return ret
		}
		return nil
	}
	if arg.GetType() != types.IS_STRING {
		return nil
	}
	return PhpPasswordAlgoFind(arg.GetStr())
}
func PhpPasswordAlgoFindZval(arg *types.Zval) *PhpPasswordAlgo {
	return PhpPasswordAlgoFindZvalEx(arg, PhpPasswordAlgoDefault())
}
func PhpPasswordAlgoExtractIdent(hash *types.String) *types.String {
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
	return types.NewString(b.CastStr(ident, ident_end-ident))
}
func PhpPasswordAlgoIdentifyEx(hash *types.String, default_algo *PhpPasswordAlgo) *PhpPasswordAlgo {
	var algo *PhpPasswordAlgo
	var ident *types.String = PhpPasswordAlgoExtractIdent(hash)
	if ident == nil {
		return default_algo
	}
	algo = PhpPasswordAlgoFind(ident)
	types.ZendStringRelease(ident)
	if algo == nil || algo.GetValid() != nil && algo.GetValid()(hash) == 0 {
		return default_algo
	} else {
		return algo
	}
}
func ZifPasswordGetInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var algo *PhpPasswordAlgo
	var hash *types.String
	var ident *types.String
	var options types.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			hash = fp.ParseStr()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	zend.ArrayInit(return_value)
	zend.ArrayInit(&options)
	ident = PhpPasswordAlgoExtractIdent(hash)
	algo = PhpPasswordAlgoFind(ident)
	if algo == nil || algo.GetValid() != nil && algo.GetValid()(hash) == 0 {
		if ident != nil {
			types.ZendStringRelease(ident)
		}
		zend.AddAssocNull(return_value, "algo")
		zend.AddAssocString(return_value, "algoName", "unknown")
		zend.AddAssocZval(return_value, "options", &options)
		return
	}
	zend.AddAssocStr(return_value, "algo", PhpPasswordAlgoExtractIdent(hash).GetStr())
	types.ZendStringRelease(ident)
	zend.AddAssocString(return_value, "algoName", algo.GetName())
	if algo.GetGetInfo() != nil && types.FAILURE == algo.GetGetInfo()(&options, hash) {
		zend.ZvalDtor(&options)
		zend.ZvalDtor(return_value)
		return_value.SetNull()
		return
	}
	zend.AddAssocZval(return_value, "options", &options)
}
func ZifPasswordNeedsRehash(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var old_algo *PhpPasswordAlgo
	var new_algo *PhpPasswordAlgo
	var hash *types.String
	var znew_algo *types.Zval
	var options *types.Array = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			hash = fp.ParseStr()
			znew_algo = fp.ParseZval()
			fp.StartOptional()
			options = fp.ParseArrayOrObjectHt()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	new_algo = PhpPasswordAlgoFindZval(znew_algo)
	if new_algo == nil {

		/* Unknown new algorithm, never prompt to rehash. */

		return_value.SetFalse()
		return
	}
	old_algo = PhpPasswordAlgoIdentifyEx(hash, nil)
	if old_algo != new_algo {

		/* Different algorithm preferred, always rehash. */

		return_value.SetTrue()
		return
	}
	types.ZVAL_BOOL(return_value, old_algo.GetNeedsRehash()(hash, options) != 0)
	return
}
func ZifPasswordVerify(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var password *types.String
	var hash *types.String
	var algo *PhpPasswordAlgo
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			password = fp.ParseStr()
			hash = fp.ParseStr()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	algo = PhpPasswordAlgoIdentify(hash)
	types.ZVAL_BOOL(return_value, algo != nil && (algo.GetVerify() == nil || algo.GetVerify()(password, hash) != 0))
	return
}
func ZifPasswordHash(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var password *types.String
	var digest *types.String = nil
	var zalgo *types.Zval
	var algo *PhpPasswordAlgo
	var options *types.Array = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			password = fp.ParseStr()
			zalgo = fp.ParseZval()
			fp.StartOptional()
			options = fp.ParseArrayOrObjectHt()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	algo = PhpPasswordAlgoFindZval(zalgo)
	if algo == nil {
		var algostr *types.String = zend.ZvalGetString(zalgo)
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown password hashing algorithm: %s", algostr.GetVal())
		types.ZendStringRelease(algostr)
		return_value.SetNull()
		return
	}
	digest = algo.GetHash()(password, options)
	if digest == nil {

		/* algo->hash should have raised an error. */

		return_value.SetNull()
		return
	}
	return_value.SetString(digest)
	return
}
func ZifPasswordAlgos(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var algo *types.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	zend.ArrayInit(return_value)
	var __ht *types.Array = &PhpPasswordAlgos
	for _, _p := range __ht.foreachData() {
		var _z *types.Zval = _p.GetVal()

		algo = _p.GetKey()
		zend.AddNextIndexStr(return_value, algo.Copy())
	}
}
