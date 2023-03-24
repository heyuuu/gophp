package zif

import (
	"go/ast"
	f "sik/internal/cmd/sikgen/astutil"
)

type ZifInfo struct {
	funcName   string
	defName    string
	name       string
	minNumArgs int
	maxNumArgs int
	argInfos   []ArgInfo
	returnInfo *ReturnInfo
	quiet      bool
	strict     bool
	oldMode    bool
}

type ArgInfo struct {
	name       string
	typ        ZppType
	isVariadic bool
}

type ReturnInfo struct {
	typ    ZppType
	withOk bool
}

const zifAnnoName = "@zif"

type zifAnnoFlags struct {
	name       string
	strNumArgs string
	minNumArgs int
	maxNumArgs int
	quiet      bool
	strict     bool
	oldMode    bool
}

//go:generate stringer -type=ZppType
type ZppType int

const (
	_ ZppType = iota
	ZppTypeBool
	ZppTypeLong
	ZppTypeDouble
	ZppTypeString
	ZppTypeZendBool
	ZppTypeZendString
	ZppTypeZendArray
	ZppTypeZval
	// ref
	ZppTypeZendArrayRef
	ZppTypeZvalRef
	// special
	ZppTypeEx
	ZppTypeRet
	ZppTypeOpt
	ZppTypeVariadic
)

func toZppType(typ string) (ZppType, bool) {
	switch typ {
	case "bool":
		return ZppTypeBool, true
	case "int":
		return ZppTypeLong, true
	case "float64":
		return ZppTypeDouble, true
	case "string":
		return ZppTypeString, true
	case "types.ZendBool":
		return ZppTypeZendBool, true
	case "*types.Array":
		return ZppTypeZendArray, true
	case "*types.String":
		return ZppTypeZendString, true
	case "*types.Zval":
		return ZppTypeZval, true
	// ref
	case "zpp.DefRefArray":
		return ZppTypeZendArrayRef, true
	case "zpp.DefRef":
		return ZppTypeZvalRef, true
	// special
	case "zpp.DefEx":
		return ZppTypeEx, true
	case "zpp.DefReturn":
		return ZppTypeRet, true
	case "zpp.DefOpt":
		return ZppTypeOpt, true
	case "[]*types.Zval":
		return ZppTypeVariadic, true

	default:
		return 0, false
	}
}

func toZppParseMethod(typ ZppType) (string, bool) {
	switch typ {
	case ZppTypeBool:
		return "ParseBoolVal", true
	case ZppTypeLong:
		return "ParseLong", true
	case ZppTypeDouble:
		return "ParseDouble", true
	case ZppTypeString:
		return "ParseStringVal", true
	case ZppTypeZendBool:
		return "ParseBool", true
	case ZppTypeZendString:
		return "ParseStr", true
	case ZppTypeZendArray:
		return "ParseArray", true
	case ZppTypeZval:
		return "ParseZval", true
	case ZppTypeVariadic:
		return "ParseVariadic", true
	default:
		return "", false
	}
}

func toZppParseMethodEx(typ ZppType) (string, []ast.Expr, bool) {
	switch typ {
	case ZppTypeZvalRef:
		return "ParseZvalEx", []ast.Expr{f.BoolLit(false), f.BoolLit(true)}, true
	case ZppTypeZendArrayRef:
		return "ParseArrayEx", []ast.Expr{f.BoolLit(false), f.BoolLit(true)}, true
	default:
		method, ok := toZppParseMethod(typ)
		return method, nil, ok
	}
}

func toZppSetMethod(typ ZppType) (string, bool) {
	switch typ {
	case ZppTypeBool:
		return "SetBool", true
	case ZppTypeLong:
		return "SetLong", true
	case ZppTypeDouble:
		return "SetDouble", true
	case ZppTypeString:
		return "SetStringVal", true
	case ZppTypeZendString:
		return "SetString", true
	case ZppTypeZendArray:
		return "SetArray", true
	default:
		return "", false
	}
}
