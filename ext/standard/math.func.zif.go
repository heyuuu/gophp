package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifAbs
var DefZifAbs = def.DefFunc("abs", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAbs(executeData, returnValue, number)
})

// generate by ZifCeil
var DefZifCeil = def.DefFunc("ceil", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCeil(executeData, returnValue, number)
})

// generate by ZifFloor
var DefZifFloor = def.DefFunc("floor", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFloor(executeData, returnValue, number)
})

// generate by ZifRound
var DefZifRound = def.DefFunc("round", 1, 3, []def.ArgInfo{{name: "number"}, {name: "precision"}, {name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	precision := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRound(executeData, returnValue, number, nil, precision, mode)
})

// generate by ZifSin
var DefZifSin = def.DefFunc("sin", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSin(executeData, returnValue, number)
})

// generate by ZifCos
var DefZifCos = def.DefFunc("cos", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCos(executeData, returnValue, number)
})

// generate by ZifTan
var DefZifTan = def.DefFunc("tan", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTan(executeData, returnValue, number)
})

// generate by ZifAsin
var DefZifAsin = def.DefFunc("asin", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAsin(executeData, returnValue, number)
})

// generate by ZifAcos
var DefZifAcos = def.DefFunc("acos", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAcos(executeData, returnValue, number)
})

// generate by ZifAtan
var DefZifAtan = def.DefFunc("atan", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtan(executeData, returnValue, number)
})

// generate by ZifAtan2
var DefZifAtan2 = def.DefFunc("atan2", 2, 2, []def.ArgInfo{{name: "y"}, {name: "x"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	y := fp.ParseZval()
	x := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtan2(executeData, returnValue, y, x)
})

// generate by ZifSinh
var DefZifSinh = def.DefFunc("sinh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSinh(executeData, returnValue, number)
})

// generate by ZifCosh
var DefZifCosh = def.DefFunc("cosh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCosh(executeData, returnValue, number)
})

// generate by ZifTanh
var DefZifTanh = def.DefFunc("tanh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTanh(executeData, returnValue, number)
})

// generate by ZifAsinh
var DefZifAsinh = def.DefFunc("asinh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAsinh(executeData, returnValue, number)
})

// generate by ZifAcosh
var DefZifAcosh = def.DefFunc("acosh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAcosh(executeData, returnValue, number)
})

// generate by ZifAtanh
var DefZifAtanh = def.DefFunc("atanh", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtanh(executeData, returnValue, number)
})

// generate by ZifPi
var DefZifPi = def.DefFunc("pi", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPi(executeData, returnValue)
})

// generate by ZifIsFinite
var DefZifIsFinite = def.DefFunc("is_finite", 1, 1, []def.ArgInfo{{name: "val"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFinite(executeData, returnValue, val)
})

// generate by ZifIsInfinite
var DefZifIsInfinite = def.DefFunc("is_infinite", 1, 1, []def.ArgInfo{{name: "val"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInfinite(executeData, returnValue, val)
})

// generate by ZifIsNan
var DefZifIsNan = def.DefFunc("is_nan", 1, 1, []def.ArgInfo{{name: "val"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNan(executeData, returnValue, val)
})

// generate by ZifPow
var DefZifPow = def.DefFunc("pow", 2, 2, []def.ArgInfo{{name: "base"}, {name: "exponent"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	base := fp.ParseZval()
	exponent := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPow(executeData, returnValue, base, exponent)
})

// generate by ZifExp
var DefZifExp = def.DefFunc("exp", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifExp(executeData, returnValue, number)
})

// generate by ZifExpm1
var DefZifExpm1 = def.DefFunc("expm1", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifExpm1(executeData, returnValue, number)
})

// generate by ZifLog1p
var DefZifLog1p = def.DefFunc("log1p", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog1p(executeData, returnValue, number)
})

// generate by ZifLog
var DefZifLog = def.DefFunc("log", 1, 2, []def.ArgInfo{{name: "number"}, {name: "base"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	base := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog(executeData, returnValue, number, nil, base)
})

// generate by ZifLog10
var DefZifLog10 = def.DefFunc("log10", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog10(executeData, returnValue, number)
})

// generate by ZifSqrt
var DefZifSqrt = def.DefFunc("sqrt", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSqrt(executeData, returnValue, number)
})

// generate by ZifHypot
var DefZifHypot = def.DefFunc("hypot", 2, 2, []def.ArgInfo{{name: "num1"}, {name: "num2"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	num1 := fp.ParseZval()
	num2 := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHypot(executeData, returnValue, num1, num2)
})

// generate by ZifDeg2rad
var DefZifDeg2rad = def.DefFunc("deg2rad", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDeg2rad(executeData, returnValue, number)
})

// generate by ZifRad2deg
var DefZifRad2deg = def.DefFunc("rad2deg", 1, 1, []def.ArgInfo{{name: "number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRad2deg(executeData, returnValue, number)
})

// generate by ZifBindec
var DefZifBindec = def.DefFunc("bindec", 1, 1, []def.ArgInfo{{name: "binary_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	binary_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBindec(executeData, returnValue, binary_number)
})

// generate by ZifHexdec
var DefZifHexdec = def.DefFunc("hexdec", 1, 1, []def.ArgInfo{{name: "hexadecimal_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hexadecimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHexdec(executeData, returnValue, hexadecimal_number)
})

// generate by ZifOctdec
var DefZifOctdec = def.DefFunc("octdec", 1, 1, []def.ArgInfo{{name: "octal_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	octal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOctdec(executeData, returnValue, octal_number)
})

// generate by ZifDecbin
var DefZifDecbin = def.DefFunc("decbin", 1, 1, []def.ArgInfo{{name: "decimal_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDecbin(executeData, returnValue, decimal_number)
})

// generate by ZifDecoct
var DefZifDecoct = def.DefFunc("decoct", 1, 1, []def.ArgInfo{{name: "decimal_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDecoct(executeData, returnValue, decimal_number)
})

// generate by ZifDechex
var DefZifDechex = def.DefFunc("dechex", 1, 1, []def.ArgInfo{{name: "decimal_number"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDechex(executeData, returnValue, decimal_number)
})

// generate by ZifBaseConvert
var DefZifBaseConvert = def.DefFunc("base_convert", 3, 3, []def.ArgInfo{{name: "number"}, {name: "frombase"}, {name: "tobase"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	number := fp.ParseZval()
	frombase := fp.ParseZval()
	tobase := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBaseConvert(executeData, returnValue, number, frombase, tobase)
})

// generate by ZifNumberFormat
var DefZifNumberFormat = def.DefFunc("number_format", 1, 4, []def.ArgInfo{{name: "number"}, {name: "num_decimal_places"}, {name: "dec_separator"}, {name: "thousands_separator"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	num_decimal_places := fp.ParseZval()
	dec_separator := fp.ParseZval()
	thousands_separator := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifNumberFormat(executeData, returnValue, number, nil, num_decimal_places, dec_separator, thousands_separator)
})

// generate by ZifFmod
var DefZifFmod = def.DefFunc("fmod", 2, 2, []def.ArgInfo{{name: "x"}, {name: "y"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	x := fp.ParseZval()
	y := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFmod(executeData, returnValue, x, y)
})

// generate by ZifIntdiv
var DefZifIntdiv = def.DefFunc("intdiv", 2, 2, []def.ArgInfo{{name: "dividend"}, {name: "divisor"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	dividend := fp.ParseZval()
	divisor := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIntdiv(executeData, returnValue, dividend, divisor)
})
