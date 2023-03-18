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
	executeData *zend.ZendExecuteData
	numArgs     int
	minNumArgs  int
	maxNumArgs  int
	flags       int
	errorCode   int
	finish      bool

	optional        bool
	_i              int
	_real_arg_index int
	_real_arg       *types.Zval
	_arg            *types.Zval
	_expected_type  ZendExpectedType
	_error          *byte
	_dummy          types.ZendBool
}

// @see Micro: ZEND_PARSE_PARAMETERS_START | ZEND_PARSE_PARAMETERS_START_EX
func FastParseStart(executeData *zend.ZendExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParser {
	// new
	p := &FastParser{
		executeData: executeData,
		numArgs:     executeData.NumArgs(),
		minNumArgs:  minNumArgs,
		maxNumArgs:  maxNumArgs,
		flags:       flags,
		//
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
		p.errorCode = ZPP_ERROR_FAILURE
	}

	// init
	if !p.IsFinish() {
		p._real_arg_index = 0
		p._real_arg = executeData.Arg(p._real_arg_index)
	}

	return p
}

func (p *FastParser) HandleError() {
	if (p.flags & ZEND_PARSE_PARAMS_QUIET) == 0 {
		if p.errorCode == ZPP_ERROR_WRONG_CALLBACK {
			if (p.flags & ZEND_PARSE_PARAMS_THROW) != 0 {
				zend.ZendWrongCallbackException(p._i, p._error)
			} else {
				zend.ZendWrongCallbackError(p._i, p._error)
			}
		} else if p.errorCode == ZPP_ERROR_WRONG_CLASS {
			if (p.flags & ZEND_PARSE_PARAMS_THROW) != 0 {
				zend.ZendWrongParameterClassException(p._i, p._error, p._arg)
			} else {
				zend.ZendWrongParameterClassError(p._i, p._error, p._arg)
			}
		} else if p.errorCode == ZPP_ERROR_WRONG_ARG {
			if (p.flags & ZEND_PARSE_PARAMS_THROW) != 0 {
				zend.ZendWrongParameterTypeException(p._i, p._expected_type, p._arg)
			} else {
				zend.ZendWrongParameterTypeError(p._i, p._expected_type, p._arg)
			}
		}
	}
}

func (p *FastParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

func (p *FastParser) IsFinish() bool {
	return p.finish || p.errorCode != ZPP_ERROR_OK
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParser) StartOptional() {
	p.optional = true
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParser) paramPrologue(deref bool, separate bool) {
	if p.IsFinish() {
		return
	}
	p._i++
	b.Assert(p._i <= p.minNumArgs || p.optional)
	b.Assert(p._i > p.minNumArgs || !p.optional)
	if p.optional {
		if p._i > p.numArgs {
			p.finish = true
			return
		}
	}

	p._real_arg_index++
	p._real_arg = p.executeData.Arg(p._real_arg_index)
	p._arg = p._real_arg
	if deref {
		if p._arg.IsReference() {
			p._arg = types.Z_REFVAL_P(p._arg)
		}
	}
	if separate {
		types.SEPARATE_ZVAL_NOREF(p._arg)
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
	p.paramPrologue(deref, separate)
	if p.IsFinish() {
		return
	}

	if zend.ZendParseArgArray(p._arg, &dest, types.IntBool(checkNull), 0) == 0 {
		p._expected_type = Z_EXPECTED_ARRAY
		p.errorCode = ZPP_ERROR_WRONG_ARG
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT
func (p *FastParser) ParseArrayOrObject() (dest *types.Zval) {
	return p.ParseArrayOrObjectEx(false, false)
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_EX
func (p *FastParser) ParseArrayOrObjectEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.paramPrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	if zend.ZendParseArgArray(p._arg, &dest, types.IntBool(checkNull), 1) == 0 {
		p._expected_type = Z_EXPECTED_ARRAY
		p.errorCode = ZPP_ERROR_WRONG_ARG
	}

	return
}

// @see Micro: Z_PARAM_BOOL
func (p *FastParser) ParseBool() (dest types.ZendBool) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_CLASS
func (p *FastParser) ParseClass() (dest *zend.ZendClassEntry) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_DOUBLE
func (p *FastParser) ParseDouble() (dest float64) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_FUNC
func (p *FastParser) ParseFunc() (fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ARRAY_HT
func (p *FastParser) ParseArrayHt() (dest *types.ZendArray) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_HT
func (p *FastParser) ParseArrayOrObjectHt() (dest *types.ZendArray) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_LONG
func (p *FastParser) ParseLong() (dest int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STRICT_LONG
func (p *FastParser) ParseStrictLong() (dest int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_OBJECT
func (p *FastParser) ParseObject() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_OBJECT_OF_CLASS
func (p *FastParser) ParseObjectOfClass(ce *zend.ZendClassEntry) (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_PATH
func (p *FastParser) ParsePath() (dest *byte, dest_len int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_PATH_STR
func (p *FastParser) ParsePathStr() (dest *types.ZendString) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_RESOURCE
func (p *FastParser) ParseResource() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STRING
func (p *FastParser) ParseString() (dest *byte, dest_len int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STR
func (p *FastParser) ParseStr() (dest *types.ZendString) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ZVAL
func (p *FastParser) ParseZval() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ZVAL_DEREF
func (p *FastParser) ParseZvalDeref() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_VARIADIC_1
func (p *FastParser) ParseVariadic1() (dest *types.Zval, num int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_VARIADIC_0
func (p *FastParser) ParseVariadic0() (dest *types.Zval, num int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}
