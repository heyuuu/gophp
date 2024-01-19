package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

func CheckNumArgs(executeData *ExecuteData, minNumArgs int, maxNumArgs int, flags int) bool {
	// todo
	return true
}

func NewParser(executeData *ExecuteData, minNumArgs int, maxNumArgs int, flags int) zpp.IParser {
	return NewFastParamParser(executeData, minNumArgs, maxNumArgs, flags)
}

// @see Micro CHECK_NULL_PATH
func CheckNullPath(s string) bool {
	// 确认字符串是二进制安全的(即不包含 \0 字符)
	return strings.ContainsRune(s, 0)
}

/* Fast parameter parsing API */
const (
	Z_EXPECTED_LONG     = "int"
	Z_EXPECTED_BOOL     = "bool"
	Z_EXPECTED_STRING   = "string"
	Z_EXPECTED_ARRAY    = "array"
	Z_EXPECTED_FUNC     = "valid callback"
	Z_EXPECTED_RESOURCE = "resource"
	Z_EXPECTED_PATH     = "a valid path"
	Z_EXPECTED_OBJECT   = "object"
	Z_EXPECTED_DOUBLE   = "float"
)

const ZPP_ERROR_OK = 0
const ZPP_ERROR_FAILURE = 1
const ZPP_ERROR_WRONG_CALLBACK = 2
const ZPP_ERROR_WRONG_CLASS = 3
const ZPP_ERROR_WRONG_ARG = 4
const ZPP_ERROR_WRONG_COUNT = 5

/**
 * FAST_ZPP: PHP7之后新增的参数处理方式
 * @link: https://wiki.php.net/rfc/fast_zpp
 * FAST_ZPP 宏与原描述符对应表 (@see README.md):
 */
type FastParamParser struct {
	ctx        *Context
	ex         *ExecuteData
	numArgs    int // 所需参数个数 (注意: zpp.FlagOldMode 模式下， numArgs 与 executeData.NumArgs() 不一定相等)
	minNumArgs int
	maxNumArgs int
	flags      int
	err        error
	errorCode  int
	finish     bool // 解析已终止 (可能是已解析完成或出现错误)
	argIndex   int
	arg        types.Zval
	optional   bool
}

var _ zpp.IParser = (*FastParamParser)(nil)

func NewParamParser(ex *ExecuteData, minNumArgs int, maxNumArgs int, flags int) zpp.IParser {
	return NewFastParamParser(ex, minNumArgs, maxNumArgs, flags)
}

func NewFastParamParser(ex *ExecuteData, minNumArgs int, maxNumArgs int, flags int) *FastParamParser {
	return &FastParamParser{
		ex:         ex,
		numArgs:    ex.NumArgs(),
		minNumArgs: minNumArgs,
		maxNumArgs: maxNumArgs,
		flags:      flags,
	}
}

func (p *FastParamParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParamParser) StartOptional() {
	p.optional = true
}

func (p *FastParamParser) isQuiet() bool { return p.flags&zpp.FlagQuiet != 0 }
func (p *FastParamParser) isThrow() bool {
	return p.flags&zpp.FlagThrow != 0 || p.ex.IsArgUseStrictTypes()
}
func (p *FastParamParser) isOldMode() bool { return p.flags&zpp.FlagOldMode != 0 }

func (p *FastParamParser) useStrictTypes() bool { return p.ex.IsArgUseStrictTypes() }

func (p *FastParamParser) isFinish() bool { return p.finish }

func (p *FastParamParser) triggerError(errorCode int, err string) {
	// 记录错误信息
	p.errorCode = errorCode
	if errorCode != ZPP_ERROR_OK {
		p.finish = true
	}

	// 若已有异常，不做error报错
	if p.ctx.EG().HasException() {
		return
	}

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_FAILURE:
			// pass
		case ZPP_ERROR_WRONG_CALLBACK:
			message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.ex.CalleeName(), p.argIndex, err)
			//faults.InternalTypeError(p.ctx, p.isThrow(), message)
			panic(perr.Internalf(message))
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.ex.CalleeName(), p.argIndex, name, types.ZendZvalTypeName(p.arg))
			//faults.InternalTypeError(p.ctx, p.isThrow(), message)
			panic(perr.Internalf(message))
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.ex.CalleeName(), p.argIndex, expectedType, types.ZendZvalTypeName(p.arg))
			//faults.InternalTypeError(p.ctx, p.isThrow(), message)
			panic(perr.Internalf(message))
		}
	}
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParamParser) nextArg(deref bool, separate bool) (arg types.Zval, ok bool) {
	if p.isFinish() {
		return
	}

	if p.err != nil || p.argIndex >= p.ex.NumArgs() {
		return
	}

	arg = p.ex.Arg(p.argIndex)
	p.argIndex++
	p.arg = arg
	return arg, true
}

// @see Micro: Z_PARAM_BOOL，Old: 'b'
func (p *FastParamParser) ParseBool() (dest bool) {
	dest, _ = p.parseBoolEx(false, false)
	return
}
func (p *FastParamParser) ParseBoolNullable() *bool {
	dest, isNull := p.parseBoolEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseBoolEx(checkNull bool, separate bool) (dest bool, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return false, true
	}

	// parse
	if p.useStrictTypes() { // strict type
		if arg.IsBool() {
			dest, ok = arg.Bool(), true
		}
	} else { // weak type
		dest, ok = zppParseBoolWeak(p.ctx, arg)
	}
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_BOOL)
	}

	return
}

// @see Micro: Z_PARAM_LONG，Old: 'l'
func (p *FastParamParser) ParseLong() (dest int) {
	dest, _ = p.parseLongEx(false, false, false)
	return
}
func (p *FastParamParser) ParseLongNullable() *int {
	dest, isNull := p.parseLongEx(true, false, false)
	if isNull {
		return nil
	}
	return &dest
}

// @see Micro: Z_PARAM_STRICT_LONG，Old: 'L'
func (p *FastParamParser) ParseStrictLong() (dest int) {
	dest, _ = p.parseLongEx(false, false, true)
	return
}
func (p *FastParamParser) ParseStrictLongNullable() *int {
	dest, isNull := p.parseLongEx(true, false, true)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseLongEx(checkNull bool, separate bool, strict bool) (dest int, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return 0, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return 0, true
	}

	// parse
	if p.useStrictTypes() { // strict type
		if arg.IsBool() {
			dest, ok = arg.Long(), true
		}
	} else { // weak type
		dest, ok = zppParseLongWeak(p.ctx, arg, strict)
	}
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_LONG)
	}

	return
}

// @see Micro: Z_PARAM_DOUBLE，Old: 'd'
func (p *FastParamParser) ParseDouble() (dest float64) {
	dest, _ = p.parseDoubleEx(false, false)
	return
}
func (p *FastParamParser) ParseDoubleNullable() *float64 {
	dest, isNull := p.parseDoubleEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseDoubleEx(checkNull bool, separate bool) (dest float64, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return 0, true
	}

	// parse
	if p.useStrictTypes() { // strict type
		if arg.IsLong() {
			dest, ok = float64(arg.Long()), true
		} else if arg.IsDouble() {
			dest, ok = arg.Double(), true
		}
	} else { // weak type
		dest, ok = zppParseDoubleWeak(p.ctx, arg)
	}
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_DOUBLE)
	}

	return
}

// @see Micro: Z_PARAM_STRING，Old: 's' (稍有区别，返回 string 而不是 *byte+len)
// @see Micro: Z_PARAM_STR，Old: 'S' (稍有区别，返回 string 而不是 *String)
func (p *FastParamParser) ParseString() (dest string) {
	dest, _ = p.parseStringEx(false, false)
	return
}
func (p *FastParamParser) ParseStringNullable() *string {
	dest, isNull := p.parseStringEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseStringEx(checkNull bool, separate bool) (dest string, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return "", true
	}

	// parse
	if p.useStrictTypes() { // strict type
		if arg.IsString() {
			dest, ok = arg.String(), true
		}
	} else { // weak type
		dest, ok = zppParseStrWeak(p.ctx, arg)
	}
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_STRING)
	}

	return
}

// @see Micro: Z_PARAM_PATH，Old: 'p'
// @see Micro: Z_PARAM_PATH_STR，Old: 'P'
func (p *FastParamParser) ParsePath() (dest string) {
	dest, _ = p.parsePathEx(false, false)
	return
}
func (p *FastParamParser) ParsePathNullable() *string {
	dest, isNull := p.parsePathEx(false, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parsePathEx(checkNull bool, separate bool) (dest string, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return "", true
	}

	// parse
	if p.useStrictTypes() { // strict type
		if arg.IsString() {
			dest, ok = arg.String(), true
		}
	} else { // weak type
		dest, ok = zppParseStrWeak(p.ctx, arg)
	}
	if !ok || !CheckNullPath(dest) {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_PATH)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_HT，Old: 'h'
func (p *FastParamParser) ParseArray() (dest *types.Array) {
	return p.parseArrayEx(false, false)
}
func (p *FastParamParser) ParseArrayNullable() (dest *types.Array) {
	return p.parseArrayEx(true, false)
}
func (p *FastParamParser) parseArrayEx(checkNull bool, separate bool) (dest *types.Array) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	dest, ok = zppParseArrayHt(arg, checkNull, false, separate)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_HT，Old: 'H'
func (p *FastParamParser) ParseArrayOrObjectHt() (dest *types.Array) {
	return p.parseArrayOrObjectHtEx(false, false)
}
func (p *FastParamParser) parseArrayOrObjectHtEx(checkNull bool, separate bool) (dest *types.Array) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	dest, ok = zppParseArrayHt(arg, checkNull, true, separate)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT，Old: 'A'
func (p *FastParamParser) ParseArrayOrObjectZval() (dest types.Zval) {
	return p.parseArrayOrObjectEx(false, false)
}

// @see Micro: Z_PARAM_ARRAY_OR_OBJECT_EX
func (p *FastParamParser) parseArrayOrObjectEx(checkNull bool, separate bool) (dest types.Zval) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	dest, ok = zppParseArray(arg, checkNull, true)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
	}

	return
}

// @see Micro: Z_PARAM_OBJECT，Old: 'o' (类似但不相同，此处直接返回 *types.Object 而不是 *types.Zval)
func (p *FastParamParser) ParseObject() *types.Object {
	return p.parseObjectEx(false, false)
}
func (p *FastParamParser) ParseObjectNullable() *types.Object {
	return p.parseObjectEx(true, false)
}
func (p *FastParamParser) parseObjectEx(checkNull bool, separate bool) *types.Object {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return nil
	}

	dest, ok := zppParseObject(arg, nil, checkNull)
	if !ok {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_OBJECT)
	}

	return dest
}

// @see Micro: Z_PARAM_RESOURCE，Old: 'r'
func (p *FastParamParser) ParseResource() (dest types.Zval) {
	dest, _ = p.parseResourceEx(false, false)
	return
}
func (p *FastParamParser) ParseResourceNullable() *types.Zval {
	dest, isNull := p.parseResourceEx(true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseResourceEx(checkNull bool, separate bool) (dest types.Zval, isNull bool) {
	arg, ok := p.nextArg(separate, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return types.Undef, true
	}

	// parse
	if arg.IsResource() {
		return arg, true
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_RESOURCE)
	}

	return
}

func (p *FastParamParser) ParseCallable() *types.UserCallable {
	//TODO implement me
	panic(perr.Todo())
}

// @see Micro: Z_PARAM_ZVAL，Old: 'z'
func (p *FastParamParser) ParseZval() types.Zval {
	return *p.parseZvalEx(false, false)
}
func (p *FastParamParser) ParseZvalPtr() *types.Zval {
	return p.parseZvalEx(false, false)
}
func (p *FastParamParser) ParseZvalNullable() (dest *types.Zval) {
	return p.parseZvalEx(true, false)
}
func (p *FastParamParser) parseZvalEx(checkNull bool, separate bool) (dest *types.Zval) {
	return p.parseZvalEx2(checkNull, separate, separate)
}
func (p *FastParamParser) parseZvalEx2(checkNull bool, deref bool, separate bool) (dest *types.Zval) {
	arg, ok := p.nextArg(deref, separate)
	if !ok {
		return
	}

	// check null
	if checkNull && arg.IsNull() {
		return nil
	}

	return &arg
}

// @see Micro: Z_PARAM_VARIADIC | Z_PARAM_VARIADIC_EX, Old: '+ and '*'
func (p *FastParamParser) ParseVariadic(postVarargs uint) []types.Zval {
	var args []types.Zval
	p.eachVariadic(postVarargs, func(arg types.Zval) {
		args = append(args, arg)
	})
	return args
}
func (p *FastParamParser) ParseVariadicPtr(postVarargs uint) []*types.Zval {
	var args []*types.Zval
	p.eachVariadic(postVarargs, func(arg types.Zval) {
		args = append(args, &arg)
	})
	return args
}
func (p *FastParamParser) eachVariadic(postVarargs uint, h func(arg types.Zval)) {
	if p.isFinish() {
		return
	}

	numVarargs := p.numArgs - p.argIndex - int(postVarargs)
	for i := 0; i < numVarargs; i++ {
		arg, ok := p.nextArg(false, false)
		if !ok {
			break
		}
		h(arg)
	}
}

// ref

func (p *FastParamParser) ParseRefZval() types.RefZval {
	// todo ref 处理
	return p.parseZvalEx(false, true)
}
func (p *FastParamParser) ParseRefArrayOrObject() types.RefZval {
	// todo ref 处理
	dest := p.parseArrayOrObjectEx(false, true)
	return &dest
}
func (p *FastParamParser) ParseRefArray() *types.Array {
	// todo ref 处理
	return p.parseArrayEx(false, false)
}
func (p *FastParamParser) ParseRefVariadic(postVarargs uint) []types.RefZval {
	// todo ref 处理
	var args []types.RefZval
	p.eachVariadic(postVarargs, func(arg types.Zval) {
		args = append(args, &arg)
	})
	return args
}
