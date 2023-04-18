package zpp

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * FAST_ZPP: PHP7之后新增的参数处理方式
 * @link: https://wiki.php.net/rfc/fast_zpp
 * FAST_ZPP 宏与原描述符对应表 (@see README.md):
 */
type FastParser struct {
	executeData ExecuteData
	numArgs     int // 所需参数个数 (注意: numArgs 与 executeData.NumArgs() 不一定相等)
	minNumArgs  int
	maxNumArgs  int
	flags       int
	errorCode   int
	finish      bool // 解析已终止 (可能是已解析完成或出现错误)
	idx         int  // 已读取的 Arg 位置，需注意它的值是从 1 开始的，范围是 [1, numArgs]
	arg         *types.Zval
	optional    bool
}

// @see Micro: ZEND_PARSE_PARAMETERS_START | ZEND_PARSE_PARAMETERS_START_EX
func FastParseStart(executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParser {
	return FastParseStartEx(executeData.NumArgs(), executeData, minNumArgs, maxNumArgs, flags)
}

func FastParseStartEx(numArgs int, executeData ExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParser {
	// new
	p := &FastParser{
		executeData: executeData,
		numArgs:     numArgs,
		minNumArgs:  minNumArgs,
		maxNumArgs:  maxNumArgs,
		flags:       flags,
	}

	// check num args
	if !CheckNumArgsEx(p.numArgs, p.executeData, p.minNumArgs, p.maxNumArgs, p.flags) {
		p.triggerError(ZPP_ERROR_FAILURE, "")
		return p
	}

	return p
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParser) StartOptional() {
	p.optional = true
}

func (p *FastParser) currArg() *types.Zval { return p.executeData.Arg(p.idx) }

func (p *FastParser) isQuiet() bool { return p.flags&FlagQuiet != 0 }
func (p *FastParser) isThrow() bool {
	return p.flags&FlagThrow != 0 || p.executeData.IsArgUseStrictTypes()
}
func (p *FastParser) isOldMode() bool { return p.flags&FlagOldMode != 0 }

func (p *FastParser) useWeakTypes() bool { return !p.executeData.IsArgUseStrictTypes() }

func (p *FastParser) IsFinish() bool { return p.finish }

func (p *FastParser) HasError() bool { return p.errorCode != ZPP_ERROR_OK }

func (p *FastParser) triggerError(errorCode int, err string) {
	// 记录错误信息
	p.errorCode = errorCode
	if errorCode != ZPP_ERROR_OK {
		p.finish = true
	}

	// 若已有异常，不做error报错
	if existException() {
		return
	}

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_FAILURE:
			// pass
		case ZPP_ERROR_WRONG_CALLBACK:
			message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.executeData.CalleeName(), p.idx, err)
			faults.InternalTypeErrorEx(p.isThrow(), message)
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.executeData.CalleeName(), p.idx, name, types.ZendZvalTypeName(p.arg))
			faults.InternalTypeErrorEx(p.isThrow(), message)
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.executeData.CalleeName(), p.idx, expectedType, types.ZendZvalTypeName(p.arg))
			faults.InternalTypeErrorEx(p.isThrow(), message)
		}
	}
}

func (p *FastParser) triggerDeprecated(errorCode int, err string) {
	switch errorCode {
	case ZPP_ERROR_WRONG_CALLBACK:
		message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.executeData.CalleeName(), p.idx, err)
		faults.Error(faults.E_DEPRECATED, message)
	}
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParser) parsePrologue(deref bool, separate bool) {
	if p.IsFinish() {
		return
	}

	if p.isOldMode() {
		p.parsePrologueOldZpp(deref, separate)
	} else {
		p.parsePrologueFastZpp(deref, separate)
	}
}

func (p *FastParser) parsePrologueFastZpp(deref bool, separate bool) {
	p.idx++
	b.Assert(p.idx <= p.minNumArgs || p.optional)
	b.Assert(p.idx > p.minNumArgs || !p.optional)
	if p.optional {
		if p.idx > p.numArgs {
			p.finish = true
			return
		}
	}

	p.readArg(deref, separate)
}

func (p *FastParser) parsePrologueOldZpp(deref bool, separate bool) {
	// 在 old zpp 模式下，未强制设置 separate 时， deref 默认值是反的
	// todo 待确认是否可优化到 Parser 外
	if !separate {
		deref = !deref
	}

	p.idx++
	p.readArg(deref, separate)
}

func (p *FastParser) readArg(deref bool, separate bool) {
	p.arg = p.currArg()
	if separate {
		// separate 为 true 时，必然需要会执行 DeRef，所以无需再判断 deref
		types.SeparateZval(p.arg)
	} else if deref {
		p.arg = p.arg.DeRef()
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
	dest, _ = p.ParseBoolEx(false, false)
	return
}
func (p *FastParser) ParseBoolEx(checkNull bool, separate bool) (dest types.ZendBool, isNull types.ZendBool) {
	val, valIsNull := p.ParseBoolValEx(checkNull, separate)
	return types.IntBool(val), types.IntBool(valIsNull)
}

//
func (p *FastParser) ParseBoolVal() (dest bool) {
	dest, _ = p.ParseBoolValEx(false, false)
	return
}
func (p *FastParser) ParseBoolValNullable() *bool {
	dest, isNull := p.ParseBoolValEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParser) ParseBoolValEx(checkNull bool, separate bool) (dest bool, isNull bool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, isNull, ok := ParseBool(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_BOOL)
	}

	return
}

// @see Micro: Z_PARAM_CLASS，Old: 'C'
func (p *FastParser) ParseClass(baseCe *types.ClassEntry) (dest *types.ClassEntry) {
	return p.ParseClassEx(baseCe, false, false)
}

func (p *FastParser) ParseClassEx(baseCe *types.ClassEntry, checkNull bool, separate bool) (dest *types.ClassEntry) {
	p.parsePrologue(separate, separate)
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
	dest, _ = p.ParseDoubleEx(false, false)
	return
}
func (p *FastParser) ParseDoubleNullable() *float64 {
	dest, isNull := p.ParseDoubleEx(false, false)
	if isNull == types.SUCCESS {
		return nil
	}
	return &dest
}
func (p *FastParser) ParseDoubleEx(checkNull bool, separate bool) (dest float64, isNull types.ZendBool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, isNullBool, ok := ParseDouble(p.arg, checkNull, p.useWeakTypes())
	isNull = types.IntBool(isNullBool)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
	}

	return
}

// @see Micro: Z_PARAM_FUNC，Old: 'f'
func (p *FastParser) ParseFunc(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache) {
	p.ParseFuncEx(fci, fcc, false, false)
}
func (p *FastParser) ParseFuncEx(fci *types.ZendFcallInfo, fcc *types.ZendFcallInfoCache, checkNull bool, separate bool) {
	p.parsePrologue(separate, separate)
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
func (p *FastParser) ParseArrayHt() (dest *types.Array) {
	return p.ParseArrayHtEx(false, false)
}
func (p *FastParser) ParseArrayHtEx(checkNull bool, separate bool) (dest *types.Array) {
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
func (p *FastParser) ParseArrayOrObjectHt() (dest *types.Array) {
	return p.ParseArrayOrObjectHtEx(false, false)
}
func (p *FastParser) ParseArrayOrObjectHtEx(checkNull bool, separate bool) (dest *types.Array) {
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
	dest, _ = p.ParseLongEx(false, false)
	return
}
func (p *FastParser) ParseLongNullable() *int {
	dest, isNull := p.ParseLongEx(false, false)
	if isNull == types.SUCCESS {
		return nil
	}
	return &dest
}
func (p *FastParser) ParseLongEx(checkNull bool, separate bool) (dest int, isNull types.ZendBool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, isNullBool, ok := ParseLong(p.arg, checkNull, false, p.useWeakTypes())
	isNull = types.IntBool(isNullBool)

	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}

	return
}

// @see Micro: Z_PARAM_STRICT_LONG，Old: 'L'
func (p *FastParser) ParseStrictLong() (dest int) {
	dest, _ = p.ParseStrictLongEx(false, false)
	return
}
func (p *FastParser) ParseStrictLongNullable() *int {
	dest, isNull := p.ParseStrictLongEx(false, false)
	if isNull == types.SUCCESS {
		return nil
	}
	return &dest
}
func (p *FastParser) ParseStrictLongEx(checkNull bool, separate bool) (dest int, isNull types.ZendBool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, isNullBool, ok := ParseLong(p.arg, checkNull, true, p.useWeakTypes())
	isNull = types.IntBool(isNullBool)

	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}

	return
}

// @see Micro: Z_PARAM_OBJECT，Old: 'o'
func (p *FastParser) ParseObject() (dest *types.Zval) {
	return p.ParseObjectEx(false, false)
}
func (p *FastParser) ParseObjectEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
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
	return p.ParseObjectOfClassEx(ce, false, false)
}
func (p *FastParser) ParseObjectOfClassEx(ce *types.ClassEntry, checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
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
func (p *FastParser) ParsePath() (strPtr *byte, strLen int) {
	return p.ParsePathEx(false, false)
}
func (p *FastParser) ParsePathEx(checkNull bool, separate bool) (strPtr *byte, strLen int) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	val, ok := ParsePathStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}
	if checkNull && val == nil {
		return nil, 0
	} else {
		return b.CastStrPtr(val.GetStr()), val.GetLen()
	}
}

func (p *FastParser) ParsePathVal() string {
	str := p.ParsePathValEx(false, false)
	if str == nil {
		return ""
	} else {
		return *str
	}
}
func (p *FastParser) ParsePathValNullable() *string {
	return p.ParsePathValEx(false, false)
}
func (p *FastParser) ParsePathValEx(checkNull bool, separate bool) *string {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return nil
	}

	zs, ok := ParsePathStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	if checkNull && zs == nil {
		return nil
	} else {
		str := zs.GetStr()
		return &str
	}
}

// @see Micro: Z_PARAM_PATH_STR，Old: 'P'
func (p *FastParser) ParsePathStr() (dest *types.String) {
	return p.ParsePathStrEx(false, false)
}
func (p *FastParser) ParsePathStrEx(checkNull bool, separate bool) (dest *types.String) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParsePathStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_RESOURCE，Old: 'r'
func (p *FastParser) ParseResource() (dest *types.Zval) {
	return p.ParseResourceEx(false, false)
}
func (p *FastParser) ParseResourceEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
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
func (p *FastParser) ParseString() (strPtr *byte, strLen int) {
	return p.ParseStringEx(false, false)
}
func (p *FastParser) ParseStringEx(checkNull bool, separate bool) (strPtr *byte, strLen int) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	val, ok := ParseZStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}
	if checkNull && val == nil {
		return nil, 0
	} else {
		return b.CastStrPtr(val.GetStr()), val.GetLen()
	}
}

// 替代 ParseString, 使用 string 代替 *byte+len
func (p *FastParser) ParseStringVal() (dest string) {
	dest, _ = p.ParseStringValEx(false, false)
	return
}
func (p *FastParser) ParseStringValNullable() *string {
	dest, isNull := p.ParseStringValEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParser) ParseStringValEx(checkNull bool, separate bool) (dest string, isNull bool) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	zs, ok := ParseZStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	if checkNull && zs == nil {
		return "", true
	} else {
		return zs.GetStr(), false
	}
}

// @see Micro: Z_PARAM_STR，Old: 'S'
func (p *FastParser) ParseStr() (dest *types.String) {
	return p.ParseStrEx(false, false)
}
func (p *FastParser) ParseStrEx(checkNull bool, separate bool) (dest *types.String) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	dest, ok := ParseZStr(p.arg, checkNull, p.useWeakTypes())
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return
}

// @see Micro: Z_PARAM_ZVAL，Old: 'z'
func (p *FastParser) ParseZval() (dest *types.Zval) {
	return p.ParseZvalEx(false, false)
}
func (p *FastParser) ParseZvalEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(separate, separate)
	if p.IsFinish() {
		return
	}

	return ParseZvalDeref(p.arg, checkNull)
}

// @see Micro: Z_PARAM_ZVAL_DEREF, Old: ""
func (p *FastParser) ParseZvalDeref() (dest *types.Zval) {
	return p.ParseZvalDerefEx(false, false)
}
func (p *FastParser) ParseZvalDerefEx(checkNull bool, separate bool) (dest *types.Zval) {
	p.parsePrologue(true, separate)
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
