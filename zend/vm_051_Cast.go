package zend

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func ZEND_CAST_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types2.Zval
	var result *types2.Zval = opline.Result()
	var ht *types2.Array
	expr = opline.Const1()
	switch opline.GetExtendedValue() {
	case types2.IS_NULL:
		result.SetNull()
	case types2.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types2.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types2.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types2.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			{

				// result.TryAddRefcount()

			}

			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types2.IS_ARRAY {
			{
				if expr.GetType() != types2.IS_NULL {
					result.SetArray(types2.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)
					{

						// expr.TryAddRefcount()

					}

				} else {
					result.SetEmptyArray()
				}
			}

			/* fast copy */

		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types2.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types2.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types2.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types2.IS_NULL {
				ht = types2.NewArray(1)
				types2.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types2.STR_SCALAR, expr)
				{

					// expr.TryAddRefcount()

				}

			}
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CAST_SPEC_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types2.Zval
	var result *types2.Zval = opline.Result()
	var ht *types2.Array
	expr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	switch opline.GetExtendedValue() {
	case types2.IS_NULL:
		result.SetNull()
	case types2.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types2.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types2.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types2.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types2.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types2.IS_ARRAY {
			if expr.GetType() != types2.IS_OBJECT || types2.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types2.IS_NULL {
					result.SetArray(types2.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types2.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types2.ZendProptableToSymtable(obj_ht, types2.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types2.Z_OBJ_P(expr).GetHandlers() != StdObjectHandlersPtr || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types2.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types2.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types2.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types2.IS_NULL {
				ht = types2.NewArray(1)
				types2.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types2.STR_SCALAR, expr)

				{

					// expr.TryAddRefcount()

				}
			}
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CAST_SPEC_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var expr *types2.Zval
	var result *types2.Zval = opline.Result()
	var ht *types2.Array
	expr = opline.Op1()
	switch opline.GetExtendedValue() {
	case types2.IS_NULL:
		result.SetNull()
	case types2.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types2.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types2.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types2.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types2.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			// ZvalPtrDtorNogc(free_op1)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types2.IS_ARRAY {
			if expr.GetType() != types2.IS_OBJECT || types2.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types2.IS_NULL {
					result.SetArray(types2.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types2.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types2.ZendProptableToSymtable(obj_ht, types2.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types2.Z_OBJ_P(expr).GetHandlers() != StdObjectHandlersPtr || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types2.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types2.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types2.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types2.IS_NULL {
				ht = types2.NewArray(1)
				types2.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types2.STR_SCALAR, expr)

				{

					// expr.TryAddRefcount()

				}
			}
		}
	}
	// ZvalPtrDtorNogc(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_CAST_SPEC_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types2.Zval
	var result *types2.Zval = opline.Result()
	var ht *types2.Array
	expr = opline.Cv1OrUndef()
	switch opline.GetExtendedValue() {
	case types2.IS_NULL:
		result.SetNull()
	case types2.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types2.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types2.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types2.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types2.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types2.IS_ARRAY {
			if expr.GetType() != types2.IS_OBJECT || types2.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types2.IS_NULL {
					result.SetArray(types2.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types2.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types2.ZendProptableToSymtable(obj_ht, types2.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types2.Z_OBJ_P(expr).GetHandlers() != StdObjectHandlersPtr || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types2.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types2.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types2.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types2.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types2.IS_NULL {
				ht = types2.NewArray(1)
				types2.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types2.STR_SCALAR, expr)

				{

					// expr.TryAddRefcount()

				}
			}
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
