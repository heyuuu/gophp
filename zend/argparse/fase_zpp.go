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

func (p *FastParser) IsFinish() bool {
	return p.finish || p.errorCode != ZPP_ERROR_OK
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParser) StartOptional() {
	p.optional = true
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParser) paramPrologue(deref bool, separate bool) {
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
func (p *FastParser) PARAM_ARRAY() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT
func (p *FastParser) PARAM_ARRAY_OR_OBJECT() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_BOOL
func (p *FastParser) PARAM_BOOL() (dest types.ZendBool) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_CLASS
func (p *FastParser) PARAM_CLASS() (dest *zend.ZendClassEntry) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_DOUBLE
func (p *FastParser) PARAM_DOUBLE() (dest float64) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_FUNC
func (p *FastParser) PARAM_FUNC() (fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ARRAY_HT
func (p *FastParser) PARAM_ARRAY_HT() (dest *types.ZendArray) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_HT
func (p *FastParser) PARAM_ARRAY_OR_OBJECT_HT() (dest *types.ZendArray) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_LONG
func (p *FastParser) PARAM_LONG() (dest int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STRICT_LONG
func (p *FastParser) PARAM_STRICT_LONG() (dest int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_OBJECT
func (p *FastParser) PARAM_OBJECT() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_OBJECT_OF_CLASS
func (p *FastParser) PARAM_OBJECT_OF_CLASS(ce *zend.ZendClassEntry) (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_PATH
func (p *FastParser) PARAM_PATH() (dest *byte, dest_len int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_PATH_STR
func (p *FastParser) PARAM_PATH_STR() (dest *types.ZendString) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_RESOURCE
func (p *FastParser) PARAM_RESOURCE() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STRING
func (p *FastParser) PARAM_STRING() (dest *byte, dest_len int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_STR
func (p *FastParser) PARAM_STR() (dest *types.ZendString) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ZVAL
func (p *FastParser) PARAM_ZVAL() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_ZVAL_DEREF
func (p *FastParser) PARAM_ZVAL_DEREF() (dest *types.Zval) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_VARIADIC_1
func (p *FastParser) PARAM_VARIADIC_1() (dest *types.Zval, num int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}

// @see Micro: Z_PARAM_VARIADIC_0
func (p *FastParser) PARAM_VARIADIC_0() (dest *types.Zval, num int) {
	if p.IsFinish() {
		return
	}

	// todo
	return
}
