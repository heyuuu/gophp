package zif

import (
	"go/ast"
	"strconv"
)

type ZifInfo struct {
	funcName      string
	defName       string
	name          string
	minNumArgs    int
	maxNumArgs    int
	useArgNames   bool
	argNames      []string
	argInfos      []ArgInfo
	argNeedEx     bool
	argNeedRet    bool
	returnArgInfo *ArgInfo
	strict        bool
}

type ArgInfo struct {
	name       string
	typ        ZppType
	isVariadic bool
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
)

func toZppType(typ ast.Expr) (zpp ZppType, ok bool) {
	typName := printNode(typ)
	switch typName {
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
	default:
		return 0, false
	}
}
