package argparse

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
)

/**
 * FAST_ZPP: PHP7之后新增的参数处理方式
 * @link: https://wiki.php.net/rfc/fast_zpp
 * 涉及对应以下宏
 * - ZEND_PARSE_PARAMETERS_START
 * - ZEND_PARSE_PARAMETERS_END
 * - ZEND_PARAM_OPTION
 * - ZEND_PARSE_*
 * - ...
 *
 * FAST_ZPP 宏与原描述符对应表 (@se):
 * | 原描述符 |  		FAST_ZPP 宏 			| 					参数类型 							|
 * | ------ | --------------------------------- | ----------------------------------------------------- |
 * | 	| 	| Z_PARAM_OPTIONAL 					| 														|
 * | 	a 	| Z_PARAM_ARRAY(dest) 				| dest - zval* 											|
 * | 	A 	| Z_PARAM_ARRAY_OR_OBJECT(dest) 	| dest - zval* 											|
 * | 	b 	| Z_PARAM_BOOL(dest) 				| dest - zend_bool 										|
 * | 	C 	| Z_PARAM_CLASS(dest) 				| dest - zend_class_entry* 								|
 * | 	d 	| Z_PARAM_DOUBLE(dest) 				| dest - double 										|
 * | 	f 	| Z_PARAM_FUNC(fci, fcc) 			| fci - zend_fcall_info, fcc - zend_fcall_info_cache 	|
 * | 	h 	| Z_PARAM_ARRAY_HT(dest) 			| dest - HashTable* 									|
 * | 	H 	| Z_PARAM_ARRAY_OR_OBJECT_HT(dest) 	| dest - HashTable* 									|
 * | 	l 	| Z_PARAM_LONG(dest) 				| dest - long 											|
 * | 	L 	| Z_PARAM_STRICT_LONG(dest) 		| dest - long 											|
 * | 	o 	| Z_PARAM_OBJECT(dest) 				| dest - zval* 											|
 * | 	O 	| Z_PARAM_OBJECT_OF_CLASS(dest, ce) | dest - zval* 											|
 * | 	p 	| Z_PARAM_PATH(dest, dest_len) 		| dest - char*, dest_len - int 							|
 * | 	P 	| Z_PARAM_PATH_STR(dest) 			| dest - zend_string* 									|
 * | 	r 	| Z_PARAM_RESOURCE(dest) 			| dest - zval* 											|
 * | 	s 	| Z_PARAM_STRING(dest, dest_len) 	| dest - char*, dest_len - int 							|
 * | 	S 	| Z_PARAM_STR(dest) 				| dest - zend_string* 									|
 * | 	z 	| Z_PARAM_ZVAL(dest) 				| dest - zval* 											|
 * | 	  	| Z_PARAM_ZVAL_DEREF(dest) 			| dest - zval* 											|
 * | 	+ 	| Z_PARAM_VARIADIC('+', dest, num) 	| dest - zval*, num int 								|
 * | 	* 	| Z_PARAM_VARIADIC('*', dest, num) 	| dest - zval*, num int 								|
 */
type FastParser struct {
	executeData ExecuteData
	numArgs     int
	minNumArgs  int
	maxNumArgs  int
	flags       int
	errorCode   int
	finish      bool // 解析已终止 (可能是已解析完成或出现错误)
	optional    bool
	idx         int
	arg         *types.Zval
}

// @see Micro: ZEND_PARSE_PARAMETERS_START | ZEND_PARSE_PARAMETERS_START_EX
func FastParseStart(executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParser {
	// new
	p := &FastParser{
		executeData: executeData,
		numArgs:     executeData.NumArgs(),
		minNumArgs:  minNumArgs,
		maxNumArgs:  maxNumArgs,
		flags:       flags,
	}

	// check num args
	if p.numArgs < minNumArgs || p.numArgs > maxNumArgs && maxNumArgs >= 0 {
		if (flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
			if (flags & ZEND_PARSE_PARAMS_THROW) != 0 {
				zend.CheckNumArgsException(minNumArgs, maxNumArgs)
			} else {
				zend.CheckNumArgsError(minNumArgs, maxNumArgs)
			}
		}
		p.triggerError(ZPP_ERROR_FAILURE, "")
	}

	return p
}

func (p *FastParser) currArg() *types.Zval {
	return p.executeData.Arg(p.idx)
}

func (p *FastParser) isQuiet() bool { return p.flags&ZEND_PARSE_PARAMS_QUIET != 0 }
func (p *FastParser) isThrow() bool { return p.flags&ZEND_PARSE_PARAMS_THROW != 0 }

func (p *FastParser) triggerError(errorCode int, err string) {
	// 记录错误信息
	p.errorCode = errorCode
	if errorCode != ZPP_ERROR_OK {
		p.finish = true
	}

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_WRONG_CALLBACK:
			zend.WrongCallbackError(p.idx, err, p.isThrow())
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			zend.WrongParamClassError(p.idx, name, p.arg, p.isThrow())
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			zend.WrongParamTypeError(p.idx, expectedType, p.arg, p.isThrow())
		}
	}
}

func (p *FastParser) HandleError() {}

func (p *FastParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

func (p *FastParser) IsFinish() bool {
	return p.finish
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParser) StartOptional() {
	p.optional = true
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParser) parsePrologue(deref bool, separate bool) {
	if p.IsFinish() {
		return
	}
	p.idx++
	b.Assert(p.idx <= p.minNumArgs || p.optional)
	b.Assert(p.idx > p.minNumArgs || !p.optional)
	if p.optional {
		if p.idx > p.numArgs {
			p.finish = true
			return
		}
	}

	p.arg = p.currArg()
	if deref {
		if p.arg.IsReference() {
			p.arg = types.Z_REFVAL_P(p.arg)
		}
	}
	if separate {
		types.SEPARATE_ZVAL_NOREF(p.arg)
	}
}

// @see Micro: Z_PARAM_ARRAY
func (p *FastParser) ParseArray() (dest *types.Zval) {
	return p.ParseArrayEx(false, false)
}

// @see Micro: Z_PARAM_ARRAY_EX
func (p *FastParser) ParseArrayEx(checkNull bool, separate bool) (dest *types.Zval) {
	return p.ParseArrayEx2(checkNull, separate, separate)
}

// @see Micro: Z_PARAM_ARRAY_EX2
func (p *FastParser) ParseArrayEx2(checkNull bool, deref bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(deref, separate)
	if p.IsFinish() {
		return
	}

	if ZendParseArgArray(p.arg, &dest, types.IntBool(checkNull), 0) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT
func (p *FastParser) ParseArrayOrObject() (dest *types.Zval) {
	return p.ParseArrayOrObjectEx(false, false)
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_EX
func (p *FastParser) ParseArrayOrObjectEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	if ZendParseArgArray(p.arg, &dest, types.IntBool(checkNull), 1) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_BOOL
func (p *FastParser) ParseBool() (dest types.ZendBool) {
	dest, _ = p.ParseBoolEx(false)
	return
}

// @see Micro: Z_PARAM_BOOL_EX
func (p *FastParser) ParseBoolEx(checkNull bool) (dest types.ZendBool, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgBool(p.arg, &dest, &isNull, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_BOOL)
	}

	return
}

// @see Micro: Z_PARAM_CLASS
func (p *FastParser) ParseClass() (dest *zend.ZendClassEntry) {
	return p.ParseClassEx(false)
}
func (p *FastParser) ParseClassEx(checkNull bool) (dest *zend.ZendClassEntry) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgClass(p.arg, &dest, p.idx, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_FAILURE, "")
	}

	return
}

// @see Micro: Z_PARAM_DOUBLE
func (p *FastParser) ParseDouble() (dest float64) {
	dest, _ = p.ParseDoubleEx(false)
	return
}
func (p *FastParser) ParseDoubleEx(checkNull bool) (dest float64, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgDouble(p.arg, &dest, &isNull, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
	}

	return
}

// @see Micro: Z_PARAM_FUNC
func (p *FastParser) ParseFunc(fci *zend.ZendFcallInfo, fcc *zend.ZendFcallInfoCache) {
	p.ParseFuncEx(fci, fcc, false)
}
func (p *FastParser) ParseFuncEx(fci *zend.ZendFcallInfo, fcc *zend.ZendFcallInfoCache, checkNull bool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	var err *string
	if ZendParseArgFunc(p.arg, fci, fcc, types.IntBool(checkNull), &err) == 0 {
		if err == nil {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_FUNC)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_CALLBACK, *err)
		}
	} else if err != nil {
		zend.ZendWrongCallbackDeprecated(p.idx, *err)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_HT
func (p *FastParser) ParseArrayHt() (dest *types.ZendArray) {
	return p.ParseArrayHtEx(false, false)
}
func (p *FastParser) ParseArrayHtEx(checkNull bool, separate bool) (dest *types.ZendArray) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}
	if ZendParseArgArrayHt(p.arg, &dest, types.IntBool(checkNull), 0, types.IntBool(separate)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}
	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_HT
func (p *FastParser) ParseArrayOrObjectHt() (dest *types.ZendArray) {
	dest, _ = p.ParseArrayOrObjectHtEx(false, false)
	return
}
func (p *FastParser) ParseArrayOrObjectHtEx(checkNull bool, separate bool) (dest *types.ZendArray, isNull types.ZendBool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	if ZendParseArgArrayHt(p.arg, &dest, types.IntBool(checkNull), 1, types.IntBool(separate)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_LONG
func (p *FastParser) ParseLong() (dest int) {
	dest, _ = p.ParseLongEx(false)
	return
}
func (p *FastParser) ParseLongEx(checkNull bool) (dest int, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, isNullBool, ok := ParseLong(p.arg, checkNull, false)
	isNull = types.IntBool(isNullBool)

	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}

	return
}

// @see Micro: Z_PARAM_STRICT_LONG
func (p *FastParser) ParseStrictLong() (dest int) {
	dest, _ = p.ParseStrictLongEx(false)
	return
}
func (p *FastParser) ParseStrictLongEx(checkNull bool) (dest int, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, isNullBool, ok := ParseLong(p.arg, checkNull, true)
	isNull = types.IntBool(isNullBool)

	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}

	return
}

// @see Micro: Z_PARAM_OBJECT
func (p *FastParser) ParseObject() (dest *types.Zval) {
	dest, _ = p.ParseObjectEx(false)
	return
}
func (p *FastParser) ParseObjectEx(checkNull bool) (dest *types.Zval, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgObject(p.arg, &dest, nil, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
	}

	return
}

// @see Micro: Z_PARAM_OBJECT_OF_CLASS
func (p *FastParser) ParseObjectOfClass(ce *zend.ZendClassEntry) (dest *types.Zval) {
	dest, _ = p.ParseObjectOfClassEx(ce, false)
	return
}
func (p *FastParser) ParseObjectOfClassEx(ce *zend.ZendClassEntry, checkNull bool) (dest *types.Zval, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgObject(p.arg, &dest, ce, types.IntBool(checkNull)) == 0 {
		if ce != nil {
			p.triggerError(ZPP_ERROR_FAILURE, ce.Name())
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
		}
	}

	return
}

// @see Micro: Z_PARAM_PATH
func (p *FastParser) ParsePath() (dest *byte, destLen int) {
	return p.ParsePathEx(false)
}
func (p *FastParser) ParsePathEx(checkNull bool) (dest *byte, destLen int) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgPath(p.arg, &dest, &destLen, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_PATH_STR
func (p *FastParser) ParsePathStr() (dest *types.ZendString) {
	return p.ParsePathStrEx(false)
}
func (p *FastParser) ParsePathStrEx(checkNull bool) (dest *types.ZendString) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgPathStr(p.arg, &dest, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_RESOURCE
func (p *FastParser) ParseResource() (dest *types.Zval) {
	return p.ParseResourceEx(false)
}
func (p *FastParser) ParseResourceEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgResource(p.arg, &dest, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_RESOURCE)
	}

	return
}

// @see Micro: Z_PARAM_STRING
func (p *FastParser) ParseString() (dest *byte, destLen int) {
	return p.ParseStringEx(false)
}
func (p *FastParser) ParseStringEx(checkNull bool) (dest *byte, destLen int) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgString(p.arg, &dest, &destLen, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return
}

// @see Micro: Z_PARAM_STR
func (p *FastParser) ParseStr() (dest *types.ZendString) {
	return p.ParseStrEx(false)
}
func (p *FastParser) ParseStrEx(checkNull bool) (dest *types.ZendString) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	if ZendParseArgStr(p.arg, &dest, types.IntBool(checkNull)) == 0 {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return
}

// @see Micro: Z_PARAM_ZVAL
func (p *FastParser) ParseZval() (dest *types.Zval) {
	return p.ParseZvalEx(false)
}
func (p *FastParser) ParseZvalEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	ZendParseArgZvalDeref(p.arg, &dest, types.IntBool(checkNull))

	return
}

// @see Micro: Z_PARAM_ZVAL_DEREF
func (p *FastParser) ParseZvalDeref() (dest *types.Zval) {
	dest, _ = p.ParseZvalDerefEx(false)
	return
}
func (p *FastParser) ParseZvalDerefEx(checkNull bool) (dest *types.Zval, isNull types.ZendBool) {
	p.parsePrologue(true, false)
	if p.IsFinish() {
		return
	}

	ZendParseArgZvalDeref(p.arg, &dest, types.IntBool(checkNull))

	return
}

// @see Micro: Z_PARAM_VARIADIC | Z_PARAM_VARIADIC_EX, Old: "+" and "*"
func (p *FastParser) ParseVariadic() []*types.Zval {
	return p.ParseVariadicEx(0)
}
func (p *FastParser) ParseVariadicEx(postVarargs int) []*types.Zval {
	if p.IsFinish() {
		return nil
	}

	numVarargs := p.numArgs - p.idx - postVarargs
	var args []*types.Zval
	for i := 0; i < numVarargs; i++ {
		p.idx++
		args = append(args, p.currArg())
	}

	return args
}
func (p *FastParser) ParseVariadic0() (dest *types.Zval, num int) {
	if p.IsFinish() {
		return
	}

	args := p.ParseVariadic()
	return args[0], len(args)
}
