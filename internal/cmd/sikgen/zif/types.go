package zif

import (
	"strconv"
)

type ZifInfo struct {
	funcName      string
	defName       string
	name          string
	minNumArgs    int
	maxNumArgs    int
	argInfos      []ArgInfo
	returnArgInfo *ArgInfo
	strict        bool
}

type ArgInfo struct {
	name       string
	typ        ZppType
	isVariadic bool
}

const zifAnnoName = "@zif"

type zifAnnoFlags struct {
	name       string
	strNumArgs string
	minNumArgs int
	maxNumArgs int
	strict     bool
	typeSpec   string
}

type ZppType int

func (z ZppType) String() string { return strconv.Itoa(int(z)) }

const (
	_ ZppType = iota
	ZppTypeBool
	ZppTypeLong
	ZppTypeDouble
	ZppTypeString
	ZppTypeZendBool
	ZppTypeZendString
	ZppTypeZendArray
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
	// special
	case "zpp.DefEx":
		return ZppTypeEx, true
	case "zpp.DefRet":
		return ZppTypeRet, true
	case "Zpp.DefOpt":
		return ZppTypeOpt, true
	case "[]*type.Zval":
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
	case ZppTypeVariadic:
		return "ParseVariadic", true
	default:
		return "", false
	}
}
