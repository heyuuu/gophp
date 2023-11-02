package types

// zval 类型展示名
func ZvalGetType(v *Zval) string {
	switch v.Type() {
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
		return "resource"
	default:
		return "unknown type"
	}
}

func ZendZvalTypeName(arg *Zval) string {
	return ZendGetTypeByConst(arg.DeRef().Type())
}

func ZendGetTypeByConst(typ ZvalType) string {
	switch typ {
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
