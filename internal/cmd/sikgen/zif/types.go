package zif

import (
	f "github.com/heyuuu/gophp/internal/cmd/sikgen/astutil"
	"go/ast"
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
	typ    ZppRetType
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
	// nullable type
	ZppTypeBoolNullable
	ZppTypeLongNullable
	ZppTypeStrictLongNullable
	ZppTypeDoubleNullable
	ZppTypeStringNullable
	ZppTypePathNullable
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
	// nullable type
	{ZppTypeBoolNullable, "*bool", "ParseBoolValNullable", nil},
	{ZppTypeLongNullable, "*int", "ParseLongNullable", nil},
	{ZppTypeStrictLongNullable, "*zpp.StrictLong", "ParseStrictLongNullable", nil},
	{ZppTypeDoubleNullable, "*float64", "ParseDoubleNullable", nil},
	{ZppTypeStringNullable, "*string", "ParseStringValNullable", nil},
	{ZppTypePathNullable, "*zpp.Path", "ParsePathValNullable", nil},
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

//go:generate stringer -type=ZppRetType
type ZppRetType int

const (
	_ ZppRetType = iota
	RetTypeBool
	RetTypeLong
	RetTypeDouble
	RetTypeString
	RetTypeArray
	RetTypeObject
	RetTypeZval
	//
	RetTypeIntArray
	RetTypeStringArray
	RetTypeZvalArray
)

func toZppRetType(typCode string) (ZppRetType, bool) {
	switch typCode {
	case "bool":
		return RetTypeBool, true
	case "int":
		return RetTypeLong, true
	case "float64":
		return RetTypeDouble, true
	case "string":
		return RetTypeString, true
	case "*types.Array":
		return RetTypeArray, true
	case "*types.ZendObject":
		return RetTypeObject, true
	case "*types.Zval":
		return RetTypeZval, true
	// special array
	case "[]int":
		return RetTypeIntArray, true
	case "[]string":
		return RetTypeStringArray, true
	case "[]*types.Zval":
		return RetTypeZvalArray, true
	default:
		return 0, false
	}
}

func toZppSetMethod(typ ZppRetType, ret ast.Expr) (setter string, args []ast.Expr, ok bool) {
	// setter
	switch typ {
	case RetTypeBool:
		setter = "SetBool"
	case RetTypeLong:
		setter = "SetLong"
	case RetTypeDouble:
		setter = "SetDouble"
	case RetTypeString:
		setter = "SetStringVal"
	case RetTypeArray, RetTypeIntArray, RetTypeStringArray, RetTypeZvalArray:
		setter = "SetArray"
	default:
		return
	}

	// args
	switch typ {
	case RetTypeIntArray:
		arg := f.PkgCallExpr("types", "NewArrayOfInt", []ast.Expr{ret})
		args = []ast.Expr{arg}
	case RetTypeStringArray:
		arg := f.PkgCallExpr("types", "NewArrayOfString", []ast.Expr{ret})
		args = []ast.Expr{arg}
	case RetTypeZvalArray:
		arg := f.PkgCallExpr("types", "NewArrayOfZval", []ast.Expr{ret})
		args = []ast.Expr{arg}
	default:
		args = []ast.Expr{ret}
	}

	// return
	return setter, args, true
}
