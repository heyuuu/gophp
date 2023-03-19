package argparse

import (
	b "sik/builtin"
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
	baseParser
	optional bool
}

// @see Micro: ZEND_PARSE_PARAMETERS_START | ZEND_PARSE_PARAMETERS_START_EX
func FastParseStart(executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParser {
	// new
	p := &FastParser{
		baseParser: makeBaseParser(executeData, executeData.NumArgs(), minNumArgs, maxNumArgs, flags),
	}

	// check num args
	p.start()

	return p
}

// todo delete
func (p *FastParser) HandleError() {}

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
		p.arg = types.ZVAL_DEREF(p.arg)
	}
	if separate {
		types.SEPARATE_ZVAL_NOREF(p.arg)
	}
}

// @see Micro: Z_PARAM_ARRAY，Old: 'a'
func (p *FastParser) ParseArray() (dest *types.Zval) {
	return p.ParseArrayEx(false, false)
}
func (p *FastParser) ParseArrayEx(checkNull bool, separate bool) (dest *types.Zval) {
	return p.ParseArrayEx2(checkNull, separate, separate)
}
func (p *FastParser) ParseArrayEx2(checkNull bool, deref bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(deref, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseArray(p.arg, checkNull, false)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT，Old: 'A'
func (p *FastParser) ParseArrayOrObject() (dest *types.Zval) {
	return p.ParseArrayOrObjectEx(false, false)
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_EX
func (p *FastParser) ParseArrayOrObjectEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseArray(p.arg, checkNull, true)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_BOOL，Old: 'b'
func (p *FastParser) ParseBool() (dest types.ZendBool) {
	dest, _ = p.ParseBoolEx(false)
	return
}
func (p *FastParser) ParseBoolEx(checkNull bool) (dest types.ZendBool, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	val, valIsNull, ok := ParseBool(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_BOOL)
	}

	return types.IntBool(val), types.IntBool(valIsNull)
}

// @see Micro: Z_PARAM_CLASS，Old: 'C'
func (p *FastParser) ParseClass(baseCe *types.ClassEntry) (dest *types.ClassEntry) {
	return p.ParseClassEx(baseCe, false)
}

func (p *FastParser) ParseClassEx(baseCe *types.ClassEntry, checkNull bool) (dest *types.ClassEntry) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseClass(p.arg, baseCe, p.idx, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_FAILURE, "")
	}

	return
}

// @see Micro: Z_PARAM_DOUBLE，Old: 'd'
func (p *FastParser) ParseDouble() (dest float64) {
	dest, _ = p.ParseDoubleEx(false)
	return
}
func (p *FastParser) ParseDoubleEx(checkNull bool) (dest float64, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	val, valIsNull, ok := ParseDouble(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
	}

	return val, types.IntBool(valIsNull)
}

// @see Micro: Z_PARAM_FUNC，Old: 'f'
func (p *FastParser) ParseFunc(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache) {
	p.ParseFuncEx(fci, fcc, false)
}
func (p *FastParser) ParseFuncEx(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache, checkNull bool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	err, ok := ParseFunc(p.arg, fci, fcc, checkNull)
	if !ok {
		if err == nil {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_FUNC)
		} else {
			p.triggerError(ZPP_ERROR_WRONG_CALLBACK, *err)
		}
	} else if err != nil {
		p.triggerDeprecated(ZPP_ERROR_WRONG_CALLBACK, *err)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_HT，Old: 'h'
func (p *FastParser) ParseArrayHt() (dest *types.ZendArray) {
	return p.ParseArrayHtEx(false, false)
}
func (p *FastParser) ParseArrayHtEx(checkNull bool, separate bool) (dest *types.ZendArray) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseArrayHt(p.arg, checkNull, false, separate)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_HT，Old: 'H'
func (p *FastParser) ParseArrayOrObjectHt() (dest *types.ZendArray) {
	dest, _ = p.ParseArrayOrObjectHtEx(false, false)
	return
}
func (p *FastParser) ParseArrayOrObjectHtEx(checkNull bool, separate bool) (dest *types.ZendArray, isNull types.ZendBool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseArrayHt(p.arg, checkNull, true, separate)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_LONG，Old: 'l'
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

// @see Micro: Z_PARAM_STRICT_LONG，Old: 'L'
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

// @see Micro: Z_PARAM_OBJECT，Old: 'o'
func (p *FastParser) ParseObject() (dest *types.Zval) {
	return p.ParseObjectEx(false)
}
func (p *FastParser) ParseObjectEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseObject(p.arg, nil, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
	}

	return
}

// @see Micro: Z_PARAM_OBJECT_OF_CLASS，Old: 'O'
func (p *FastParser) ParseObjectOfClass(ce *types.ClassEntry) (dest *types.Zval) {
	dest, _ = p.ParseObjectOfClassEx(ce, false)
	return
}
func (p *FastParser) ParseObjectOfClassEx(ce *types.ClassEntry, checkNull bool) (dest *types.Zval, isNull types.ZendBool) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseObject(p.arg, ce, checkNull)
	if !ok {
		if ce != nil {
			p.triggerError(ZPP_ERROR_WRONG_CLASS, ce.Name())
		} else {
			p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
		}
	}

	return
}

// @see Micro: Z_PARAM_PATH，Old: 'p'
func (p *FastParser) ParsePath() (dest *byte, destLen int) {
	return p.ParsePathEx(false)
}
func (p *FastParser) ParsePathEx(checkNull bool) (dest *byte, destLen int) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, destLen, ok := ParsePathStrPtr(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_PATH_STR，Old: 'P'
func (p *FastParser) ParsePathStr() (dest *types.ZendString) {
	return p.ParsePathStrEx(false)
}
func (p *FastParser) ParsePathStrEx(checkNull bool) (dest *types.ZendString) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, ok := ParsePathStr(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_RESOURCE，Old: 'r'
func (p *FastParser) ParseResource() (dest *types.Zval) {
	return p.ParseResourceEx(false)
}
func (p *FastParser) ParseResourceEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseResource(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_RESOURCE)
	}

	return
}

// @see Micro: Z_PARAM_STRING，Old: 's'
func (p *FastParser) ParseString() (dest *byte, destLen int) {
	return p.ParseStringEx(false)
}
func (p *FastParser) ParseStringEx(checkNull bool) (dest *byte, destLen int) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	s, l, ok := ParseStrPtr(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return s, l
}

// @see Micro: Z_PARAM_STR，Old: 'S'
func (p *FastParser) ParseStr() (dest *types.ZendString) {
	return p.ParseStrEx(false)
}
func (p *FastParser) ParseStrEx(checkNull bool) (dest *types.ZendString) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	val, ok := ParseZStr(p.arg, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return val
}

// @see Micro: Z_PARAM_ZVAL，Old: 'z'
func (p *FastParser) ParseZval() (dest *types.Zval) {
	return p.ParseZvalEx(false)
}
func (p *FastParser) ParseZvalEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(false, false)
	if p.IsFinish() {
		return
	}

	return ParseZvalDeref(p.arg, checkNull)
}

// @see Micro: Z_PARAM_ZVAL_DEREF, Old: ""
func (p *FastParser) ParseZvalDeref() (dest *types.Zval) {
	return p.ParseZvalDerefEx(false)
}
func (p *FastParser) ParseZvalDerefEx(checkNull bool) (dest *types.Zval) {
	p.parsePrologue(true, false)
	if p.IsFinish() {
		return
	}

	return ParseZvalDeref(p.arg, checkNull)
}

// @see Micro: Z_PARAM_VARIADIC | Z_PARAM_VARIADIC_EX, Old: '+ and '*'
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
