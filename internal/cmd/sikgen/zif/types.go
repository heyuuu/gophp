package zif

import "go/ast"

type ZifInfo struct {
	funcName      string
	defName       string
	name          string
	handler       ast.Expr
	minNumArgs    int
	maxNumArgs    int
	useArgNames   bool
	argNames      []string
	argInfos      []ArgInfo
	returnArgInfo *ArgInfo
	strict        bool
}

type ArgInfo struct {
	name       string
	typ        ZvalType
	isVariadic bool
}

type ZvalType string

const (
	ZvalTypeBool   ZvalType = "bool"
	ZvalTypeInt    ZvalType = "int"
	ZvalTypeDouble ZvalType = "double"
	ZvalTypeString ZvalType = "string"
)

func asZvalType(typ ast.Expr) (ZvalType, bool) {
	typName := printNode(typ)
	switch typName {
	case "bool":
		return ZvalTypeBool, true
	case "int":
		return ZvalTypeInt, true
	case "float64":
		return ZvalTypeDouble, true
	case "string":
		return ZvalTypeString, true
	default:
		return "", false
	}
}
