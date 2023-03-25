package zif

import (
	"go/ast"
	f "sik/internal/cmd/sikgen/astutil"
)

type ZifInfo struct {
	funcName   string
	defName    string
	name       string
	aliasNames []string
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
	aliasNames []string
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
	// special
	ZppTypeEx
	ZppTypeRet
	ZppTypeOpt
	// base type
	ZppTypeBool
	ZppTypeLong
	ZppTypeStrictLong
	ZppTypeDouble
	ZppTypeString
	ZppTypePath
	ZppTypeArrayHt
	ZppTypeArrayOrObjectHt
	ZppTypeArray
	ZppTypeArrayOrObject
	ZppTypeClass
	ZppTypeObject
	ZppTypeZval
	ZppTypeZvalDeref
	ZppTypeVariadic
	// ref
	ZppTypeRefZval
	ZppTypeRefArray
)

var zppInfos = []struct {
	typ    ZppType
	def    string
	parser string
	args   []ast.Expr
}{
	// special
	{ZppTypeEx, "zpp.Ex", "", nil},
	{ZppTypeRet, "zpp.Ret", "", nil},
	{ZppTypeOpt, "zpp.Opt", "", nil},
	// base type
	{ZppTypeBool, "bool", "ParseBoolVal", nil},
	{ZppTypeLong, "int", "ParseLong", nil},
	{ZppTypeStrictLong, "zpp.StrictLong", "ParseStrictLong", nil},
	{ZppTypeDouble, "float64", "ParseDouble", nil},
	{ZppTypeString, "string", "ParseStringVal", nil},
	{ZppTypePath, "zpp.Path", "ParsePathVal", nil},
	{ZppTypeArrayHt, "*types.Array", "ParseArrayHt", nil},
	{ZppTypeArrayOrObjectHt, "zpp.ArrayOrObjectHt", "ParseArrayOrObjectHt", nil},
	{ZppTypeArray, "zpp.Array", "ParseArray", nil},
	{ZppTypeArrayOrObject, "zpp.ArrayOrObject", "ParseArrayOrObject", nil},
	{ZppTypeClass, "zpp.Class", "ParseClass", nil},
	{ZppTypeObject, "zpp.Object", "ParseObject", nil},
	{ZppTypeZval, "*types.Zval", "ParseZval", nil},
	{ZppTypeZvalDeref, "zpp.ZvalDeref", "ParseZvalDeref", nil},
	{ZppTypeVariadic, "[]*types.Zval", "ParseVariadic", nil},
	// ref type
	{ZppTypeRefZval, "zpp.RefZval", "ParseZvalEx", []ast.Expr{f.False(), f.True()}},
	{ZppTypeRefArray, "zpp.RefArray", "ParseArrayEx", []ast.Expr{f.False(), f.True()}},
}

var toZppTypeMap map[string]ZppType

func toZppType(typCode string) (ZppType, bool) {
	if toZppTypeMap == nil {
		toZppTypeMap = make(map[string]ZppType)
		for _, info := range zppInfos {
			if len(info.def) != 0 {
				toZppTypeMap[info.def] = info.typ
			}
		}
	}
	if typ, ok := toZppTypeMap[typCode]; ok {
		return typ, true
	}
	return 0, false
}

type parseMethod struct {
	parser string
	args   []ast.Expr
}

var toZppParseMap map[ZppType]parseMethod

func toZppParseMethodEx(typ ZppType) (string, []ast.Expr, bool) {
	if toZppParseMap == nil {
		toZppParseMap = make(map[ZppType]parseMethod)
		for _, info := range zppInfos {
			if len(info.parser) != 0 {
				toZppParseMap[info.typ] = parseMethod{parser: info.parser, args: info.args}
			}
		}
	}
	if method, ok := toZppParseMap[typ]; ok {
		return method.parser, method.args, true
	}
	return "", nil, false
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
	case ZppTypeArrayHt:
		return "SetArray", true
	default:
		return "", false
	}
}
