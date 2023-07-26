package types

import "github.com/heyuuu/gophp/zend"

// zval 类型展示名
func ZvalGetType(v *Zval) string {
	switch v.GetType() {
	case IsNull:
		return "NULL"
	case IsFalse, IsTrue:
		return "boolean"
	case IsLong:
		return "integer"
	case IsDouble:
		return "double"
	case IsString:
		return "string"
	case IsArray:
		return "array"
	case IsObject:
		return "object"
	case IsResource:
		if zend.ZendRsrcListGetRsrcType(v.Resource()) != nil {
			return "resource"
		} else {
			return "resource (closed)"
		}
	default:
		return "unknown type"
	}
}

func ZendZvalTypeName(arg *Zval) string {
	arg = ZVAL_DEREF(arg)
	return ZendGetTypeByConst(arg.GetType())
}

func ZendGetTypeByConst(type_ uint8) string {
	switch type_ {
	case IsFalse, IsTrue, IsBool:
		return "bool"
	case IsLong:
		return "int"
	case IsDouble:
		return "float"
	case IsString:
		return "string"
	case IsObject:
		return "object"
	case IsResource:
		return "resource"
	case IsNull:
		return "null"
	case IsCallable:
		return "callable"
	case IsIterable:
		return "iterable"
	case IsArray:
		return "array"
	case IsVoid:
		return "void"
	case IsNumber:
		return "number"
	default:
		return "unknown"
	}
}
