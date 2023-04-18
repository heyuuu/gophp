package standard

import (
	"github.com/heyuuu/gophp/zend/def"
	"github.com/heyuuu/gophp/zend/zpp"
)

// generate by ZifAbs
var DefZifAbs = def.DefFunc("abs", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifAbs(number)
	if ok {
		returnValue.SetBy(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifCeil
var DefZifCeil = def.DefFunc("ceil", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifCeil(number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifFloor
var DefZifFloor = def.DefFunc("floor", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ret, ok := ZifFloor(number)
	if ok {
		returnValue.SetDouble(ret)
	} else {
		returnValue.SetFalse()
	}
})

// generate by ZifRound
var DefZifRound = def.DefFunc("round", 1, 3, []def.ArgInfo{{Name: "number"}, {Name: "precision"}, {Name: "mode"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	precision := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRound(number, nil, precision, mode)
})

// generate by ZifSin
var DefZifSin = def.DefFunc("sin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSin(number)
})

// generate by ZifCos
var DefZifCos = def.DefFunc("cos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCos(executeData, returnValue, number)
})

// generate by ZifTan
var DefZifTan = def.DefFunc("tan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTan(executeData, returnValue, number)
})

// generate by ZifAsin
var DefZifAsin = def.DefFunc("asin", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAsin(executeData, returnValue, number)
})

// generate by ZifAcos
var DefZifAcos = def.DefFunc("acos", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAcos(executeData, returnValue, number)
})

// generate by ZifAtan
var DefZifAtan = def.DefFunc("atan", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtan(executeData, returnValue, number)
})

// generate by ZifAtan2
var DefZifAtan2 = def.DefFunc("atan2", 2, 2, []def.ArgInfo{{Name: "y"}, {Name: "x"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	y := fp.ParseZval()
	x := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtan2(y, x)
})

// generate by ZifSinh
var DefZifSinh = def.DefFunc("sinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSinh(number)
})

// generate by ZifCosh
var DefZifCosh = def.DefFunc("cosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifCosh(executeData, returnValue, number)
})

// generate by ZifTanh
var DefZifTanh = def.DefFunc("tanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifTanh(executeData, returnValue, number)
})

// generate by ZifAsinh
var DefZifAsinh = def.DefFunc("asinh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAsinh(executeData, returnValue, number)
})

// generate by ZifAcosh
var DefZifAcosh = def.DefFunc("acosh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAcosh(executeData, returnValue, number)
})

// generate by ZifAtanh
var DefZifAtanh = def.DefFunc("atanh", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifAtanh(executeData, returnValue, number)
})

// generate by ZifPi
var DefZifPi = def.DefFunc("pi", 0, 0, []def.ArgInfo{}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifPi()
})

// generate by ZifIsFinite
var DefZifIsFinite = def.DefFunc("is_finite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsFinite(val)
})

// generate by ZifIsInfinite
var DefZifIsInfinite = def.DefFunc("is_infinite", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsInfinite(val)
})

// generate by ZifIsNan
var DefZifIsNan = def.DefFunc("is_nan", 1, 1, []def.ArgInfo{{Name: "val"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	val := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIsNan(val)
})

// generate by ZifPow
var DefZifPow = def.DefFunc("pow", 2, 2, []def.ArgInfo{{Name: "base"}, {Name: "exponent"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	base := fp.ParseZval()
	exponent := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifPow(returnValue, base, exponent)
})

// generate by ZifExp
var DefZifExp = def.DefFunc("exp", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifExp(number)
})

// generate by ZifExpm1
var DefZifExpm1 = def.DefFunc("expm1", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifExpm1(number)
})

// generate by ZifLog1p
var DefZifLog1p = def.DefFunc("log1p", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog1p(number)
})

// generate by ZifLog
var DefZifLog = def.DefFunc("log", 1, 2, []def.ArgInfo{{Name: "number"}, {Name: "base"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 2, 0)
	number := fp.ParseZval()
	fp.StartOptional()
	base := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog(number, nil, base)
})

// generate by ZifLog10
var DefZifLog10 = def.DefFunc("log10", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifLog10(number)
})

// generate by ZifSqrt
var DefZifSqrt = def.DefFunc("sqrt", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifSqrt(number)
})

// generate by ZifHypot
var DefZifHypot = def.DefFunc("hypot", 2, 2, []def.ArgInfo{{Name: "num1"}, {Name: "num2"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	num1 := fp.ParseZval()
	num2 := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHypot(num1, num2)
})

// generate by ZifDeg2rad
var DefZifDeg2rad = def.DefFunc("deg2rad", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDeg2rad(number)
})

// generate by ZifRad2deg
var DefZifRad2deg = def.DefFunc("rad2deg", 1, 1, []def.ArgInfo{{Name: "number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifRad2deg(number)
})

// generate by ZifBindec
var DefZifBindec = def.DefFunc("bindec", 1, 1, []def.ArgInfo{{Name: "binary_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	binary_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBindec(binary_number)
})

// generate by ZifHexdec
var DefZifHexdec = def.DefFunc("hexdec", 1, 1, []def.ArgInfo{{Name: "hexadecimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	hexadecimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifHexdec(hexadecimal_number)
})

// generate by ZifOctdec
var DefZifOctdec = def.DefFunc("octdec", 1, 1, []def.ArgInfo{{Name: "octal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	octal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifOctdec(octal_number)
})

// generate by ZifDecbin
var DefZifDecbin = def.DefFunc("decbin", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDecbin(decimal_number)
})

// generate by ZifDecoct
var DefZifDecoct = def.DefFunc("decoct", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDecoct(decimal_number)
})

// generate by ZifDechex
var DefZifDechex = def.DefFunc("dechex", 1, 1, []def.ArgInfo{{Name: "decimal_number"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	decimal_number := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifDechex(decimal_number)
})

// generate by ZifBaseConvert
var DefZifBaseConvert = def.DefFunc("base_convert", 3, 3, []def.ArgInfo{{Name: "number"}, {Name: "frombase"}, {Name: "tobase"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	number := fp.ParseZval()
	frombase := fp.ParseZval()
	tobase := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifBaseConvert(number, frombase, tobase)
})

// generate by ZifNumberFormat
var DefZifNumberFormat = def.DefFunc("number_format", 1, 4, []def.ArgInfo{{Name: "number"}, {Name: "num_decimal_places"}, {Name: "dec_separator_"}, {Name: "thousands_separator"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 1, 4, 0)
	number := fp.ParseDouble()
	fp.StartOptional()
	num_decimal_places := fp.ParseLong()
	dec_separator_ := fp.ParseStringValNullable()
	thousands_separator := fp.ParseStringValNullable()
	if fp.HasError() {
		return
	}
	ret := ZifNumberFormat(number, nil, num_decimal_places, dec_separator_, thousands_separator)
	returnValue.SetStringVal(ret)
})

// generate by ZifFmod
var DefZifFmod = def.DefFunc("fmod", 2, 2, []def.ArgInfo{{Name: "x"}, {Name: "y"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	x := fp.ParseZval()
	y := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifFmod(x, y)
})

// generate by ZifIntdiv
var DefZifIntdiv = def.DefFunc("intdiv", 2, 2, []def.ArgInfo{{Name: "dividend"}, {Name: "divisor"}}, func(executeData zpp.Ex, returnValue zpp.Ret) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	dividend := fp.ParseZval()
	divisor := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifIntdiv(dividend, divisor)
})
