package php

import (
	"fmt"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

// @see Micro CHECK_NULL_PATH
func CheckNullPath(s string) bool {
	// 确认字符串是二进制安全的(即不包含 \0 字符)
	return strings.IndexByte(s, 0) < 0
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

	Z_EXPECTED_REFERENCE = "reference"
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
	ctx *Context
	//ex         *ExecuteData
	callee     string // 调用方名称
	args       []types.Zval
	minNumArgs int
	maxNumArgs int
	flags      int
	strictType bool
	errorCode  int
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
		ctx: ex.ctx,
		//ex:         ex,
		callee:     ex.CalleeName(),
		args:       ex.args,
		minNumArgs: minNumArgs,
		maxNumArgs: maxNumArgs,
		flags:      flags,
		strictType: ex.IsArgUseStrictTypes(),
	}
}

func (p *FastParamParser) HasError() bool {
	return p.errorCode != ZPP_ERROR_OK
}

func (p *FastParamParser) CheckNumArgs() {
	// 检查参数个数，若检查通过直接返回
	numArgs, minNumArgs, maxNumArgs := len(p.args), p.minNumArgs, p.maxNumArgs
	if numArgs >= minNumArgs && (numArgs <= maxNumArgs || maxNumArgs < 0) {
		return
	}

	// 触发错误
	p.triggerError(ZPP_ERROR_WRONG_COUNT, "")
}

// @see Micro: Z_PARAM_OPTIONAL
func (p *FastParamParser) StartOptional() {
	p.optional = true
}

func (p *FastParamParser) isQuiet() bool { return p.flags&zpp.FlagQuiet != 0 }
func (p *FastParamParser) isThrow() bool {
	return p.flags&zpp.FlagThrow != 0 || p.strictType
}
func (p *FastParamParser) isOldMode() bool { return p.flags&zpp.FlagOldMode != 0 }

func (p *FastParamParser) useStrictTypes() bool { return p.strictType }

func (p *FastParamParser) isFinish() bool { return p.HasError() || p.argIndex >= len(p.args) }

func (p *FastParamParser) triggerError(errorCode int, err string) {
	if errorCode == ZPP_ERROR_OK {
		return
	}

	// 记录错误信息
	p.errorCode = errorCode

	// 触发错误或异常
	if !p.isQuiet() {
		switch errorCode {
		case ZPP_ERROR_FAILURE:
			// pass
		case ZPP_ERROR_WRONG_CALLBACK:
			message := fmt.Sprintf("%s() expects parameter %d to be a valid callback, %s", p.callee, p.argIndex, err)
			InternalTypeError(p.ctx, p.isThrow(), message)
		case ZPP_ERROR_WRONG_CLASS:
			name := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.callee, p.argIndex, name, types.ZendZvalTypeName(p.arg))
			InternalTypeError(p.ctx, p.isThrow(), message)
		case ZPP_ERROR_WRONG_ARG:
			expectedType := err
			message := fmt.Sprintf("%s() expects parameter %d to be %s, %s given", p.callee, p.argIndex, expectedType, types.ZendZvalTypeName(p.arg))
			InternalTypeError(p.ctx, p.isThrow(), message)
		case ZPP_ERROR_WRONG_COUNT:
			numArgs, minNumArgs, maxNumArgs := len(p.args), p.minNumArgs, p.maxNumArgs
			if minNumArgs == maxNumArgs {
				InternalArgumentCountError(p.ctx, p.isThrow(), fmt.Sprintf("%s() expects exactly %d parameter%s, %d given", p.callee, minNumArgs, lang.Cond(minNumArgs == 1, "", "s"), numArgs))
			} else if numArgs < minNumArgs {
				InternalArgumentCountError(p.ctx, p.isThrow(), fmt.Sprintf("%s() expects at least %d parameter%s, %d given", p.callee, minNumArgs, lang.Cond(minNumArgs == 1, "", "s"), numArgs))
			} else { // numArgs > maxNumArgs
				InternalArgumentCountError(p.ctx, p.isThrow(), fmt.Sprintf("%s() expects at most %d parameter%s, %d given", p.callee, maxNumArgs, lang.Cond(maxNumArgs == 1, "", "s"), numArgs))
			}
		}
	}
}

// Micro: Z_PARAM_PROLOGUE
func (p *FastParamParser) nextArg(deref bool, separate bool) (arg types.Zval, ok bool) {
	if p.isFinish() {
		return
	}

	arg = p.args[p.argIndex]
	if deref {
		arg = arg.DeRef()
	}
	if arg.IsUndef() {
		arg = types.ZvalNull() // 正常流程中传入 args 应没有
	}
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
	dest, isNull := p.parseBoolEx(true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseBoolEx(checkNull bool, separate bool) (dest bool, isNull bool) {
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return false, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return false, true
	}

	// parse
	ok = false
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
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return 0, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return 0, true
	}

	// parse
	ok = false
	if p.useStrictTypes() { // strict type
		if arg.IsLong() {
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
	dest, isNull := p.parseDoubleEx(true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseDoubleEx(checkNull bool, separate bool) (dest float64, isNull bool) {
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return 0, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return 0, true
	}

	// parse
	ok = false
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
	dest, isNull := p.parseStringEx(true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseStringEx(checkNull bool, separate bool) (dest string, isNull bool) {
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return "", true
	}

	// check null
	if checkNull && arg.IsNull() {
		return "", true
	}

	// parse
	ok = false
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
	dest, isNull := p.parsePathEx(true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parsePathEx(checkNull bool, separate bool) (dest string, isNull bool) {
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return "", true
	}

	// check null
	if checkNull && arg.IsNull() {
		return "", true
	}

	// parse
	ok = false
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
	arg, ok := p.nextArg(true, separate)
	if !ok {
		return nil
	}

	if arg.IsArray() {
		return arg.Array()
	} else if checkNull && arg.IsNull() {
		return nil
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		return nil
	}
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
		return types.Undef, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return types.Undef, true
	}

	// parse
	if arg.IsResource() {
		return arg, false
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_RESOURCE)
	}

	return
}

func (p *FastParamParser) ParseCallable() *types.UserCallable {
	//TODO implement me
	panic(perr.Todof("ParseCallable"))
}

// @see Micro: Z_PARAM_ZVAL，Old: 'z'
func (p *FastParamParser) ParseZval() (dest types.Zval) {
	dest, _ = p.parseZvalEx2(false, true, false)
	return
}
func (p *FastParamParser) ParseZvalNullable() *types.Zval {
	dest, isNull := p.parseZvalEx2(true, true, false)
	if isNull {
		return nil
	}
	return &dest
}
func (p *FastParamParser) parseZvalEx2(checkNull bool, deref bool, separate bool) (dest types.Zval, isNull bool) {
	arg, ok := p.nextArg(deref, separate)
	if !ok {
		return types.Undef, true
	}

	// check null
	if checkNull && arg.IsNull() {
		return types.Undef, true
	}

	return arg, false
}

// @see Micro: Z_PARAM_VARIADIC | Z_PARAM_VARIADIC_EX, Old: '+ and '*'
func (p *FastParamParser) ParseVariadic(postVarargs uint) []types.Zval {
	var args []types.Zval

	numVarargs := len(p.args) - p.argIndex - int(postVarargs)
	for i := 0; i < numVarargs; i++ {
		arg, ok := p.nextArg(false, false)
		if !ok {
			return nil
		}
		args = append(args, arg)
	}

	return args
}

// ref

func (p *FastParamParser) ParseRefZval() types.RefZval {
	return p.parseRefZval(false)
}
func (p *FastParamParser) ParseRefZvalNullable() types.RefZval {
	return p.parseRefZval(true)
}
func (p *FastParamParser) ParseRefArrayOrObject() types.RefZval {
	ref := p.parseRefZval(false)
	if ref == nil {
		return nil
	}
	if ref.Val().IsArray() || ref.Val().IsObject() {
		return ref
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		return nil
	}
}
func (p *FastParamParser) ParseRefArray() *types.Array {
	ref := p.parseRefZval(false)
	if ref == nil {
		return nil
	}
	if ref.Val().IsArray() {
		return ref.Val().Array()
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		return nil
	}
}
func (p *FastParamParser) ParseRefArrayNullable() *types.Array {
	ref := p.parseRefZval(true)
	if ref == nil {
		return nil
	}
	if ref.Val().IsArray() {
		return ref.Val().Array()
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, Z_EXPECTED_ARRAY)
		return nil
	}
}
func (p *FastParamParser) ParseRefVariadic(postVarargs uint) []types.RefZval {
	var args []types.RefZval

	numVarargs := len(p.args) - p.argIndex - int(postVarargs)
	for i := 0; i < numVarargs; i++ {
		ref := p.parseRefZval(false)
		if ref == nil {
			return nil
		}
		args = append(args, ref)
	}

	return args
}

func (p *FastParamParser) parseRefZval(checkNull bool) types.RefZval {
	arg, ok := p.nextArg(false, false)
	if !ok {
		return nil
	}

	if arg.IsRef() {
		return arg.Ref()
	} else if checkNull && arg.IsNull() {
		return nil
	} else {
		p.triggerError(ZPP_ERROR_WRONG_ARG, "reference")
		return nil
	}
}
