package types

import "github.com/heyuuu/gophp/zend"

// zval 类型展示名
func ZvalGetType(v *Zval) string {
	switch v.GetType() {
	case IS_NULL:
		return "NULL"
	case IS_FALSE, IS_TRUE:
		return "boolean"
	case IS_LONG:
		return "integer"
	case IS_DOUBLE:
		return "double"
	case IS_STRING:
		return "string"
	case IS_ARRAY:
		return "array"
	case IS_OBJECT:
		return "object"
	case IS_RESOURCE:
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
	case IS_FALSE, IS_TRUE, IS_BOOL:
		return "bool"
	case IS_LONG:
		return "int"
	case IS_DOUBLE:
		return "float"
	case IS_STRING:
		return "string"
	case IS_OBJECT:
		return "object"
	case IS_RESOURCE:
		return "resource"
	case IS_NULL:
		return "null"
	case IS_CALLABLE:
		return "callable"
	case IS_ITERABLE:
		return "iterable"
	case IS_ARRAY:
		return "array"
	case IS_VOID:
		return "void"
	case IS_NUMBER:
		return "number"
	default:
		return "unknown"
	}
}
