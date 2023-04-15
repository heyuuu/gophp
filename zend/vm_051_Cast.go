package zend

import "github.com/heyuuu/gophp/zend/types"

func ZEND_CAST_SPEC_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var expr *types.Zval
	var result *types.Zval = opline.Result()
	var ht *types.Array
	expr = opline.Const1()
	switch opline.GetExtendedValue() {
	case types.IS_NULL:
		result.SetNull()
	case types.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types.IS_STRING:
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
		if opline.GetExtendedValue() == types.IS_ARRAY {
			{
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
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
				ht = types.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types.IS_NULL {
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.STR_SCALAR, expr)
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
	var expr *types.Zval
	var result *types.Zval = opline.Result()
	var ht *types.Array
	expr = _getZvalPtrTmp(opline.GetOp1().GetVar(), &free_op1, executeData)
	switch opline.GetExtendedValue() {
	case types.IS_NULL:
		result.SetNull()
	case types.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			if expr.GetType() != types.IS_OBJECT || types.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(expr).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types.IS_NULL {
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.STR_SCALAR, expr)

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
	var expr *types.Zval
	var result *types.Zval = opline.Result()
	var ht *types.Array
	expr = opline.Op1()
	switch opline.GetExtendedValue() {
	case types.IS_NULL:
		result.SetNull()
	case types.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			// ZvalPtrDtorNogc(free_op1)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			if expr.GetType() != types.IS_OBJECT || types.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(expr).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types.IS_NULL {
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.STR_SCALAR, expr)

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
	var expr *types.Zval
	var result *types.Zval = opline.Result()
	var ht *types.Array
	expr = opline.Cv1OrUndef()
	switch opline.GetExtendedValue() {
	case types.IS_NULL:
		result.SetNull()
	case types.IS_BOOL:
		result.SetBool(ZendIsTrue(expr) != 0)
	case types.IS_LONG:
		result.SetLong(ZvalGetLong(expr))
	case types.IS_DOUBLE:
		result.SetDouble(ZvalGetDouble(expr))
	case types.IS_STRING:
		result.SetString(ZvalGetString(expr))
	default:
		{
			expr = types.ZVAL_DEREF(expr)
		}

		/* If value is already of correct type, return it directly */

		if expr.IsType(opline.GetExtendedValue()) {
			result.CopyValueFrom(expr)
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}
		if opline.GetExtendedValue() == types.IS_ARRAY {
			if expr.GetType() != types.IS_OBJECT || types.Z_OBJCE_P(expr) == ZendCeClosure {
				if expr.GetType() != types.IS_NULL {
					result.SetArray(types.NewArray(1))
					expr = result.Array().IndexAddNew(0, expr)

					{

						// expr.TryAddRefcount()

					}
				} else {
					result.SetEmptyArray()
				}
			} else {
				var obj_ht *types.Array = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					result.SetArray(types.ZendProptableToSymtable(obj_ht, types.Z_OBJCE_P(expr).GetDefaultPropertiesCount() != 0 || types.Z_OBJ_P(expr).GetHandlers() != &StdObjectHandlers || obj_ht.IsRecursive()))
					ZendReleaseProperties(obj_ht)
				} else {
					result.SetEmptyArray()
				}
			}
		} else {
			result.SetObject(ZendObjectsNew(ZendStandardClassDef))
			if expr.IsArray() {
				ht = types.ZendSymtableToProptable(expr.Array())
				if (ht.GetGcFlags() & types.IS_ARRAY_IMMUTABLE) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = types.ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				types.Z_OBJ_P(result).SetProperties(ht)
			} else if expr.GetType() != types.IS_NULL {
				ht = types.NewArray(1)
				types.Z_OBJ_P(result).SetProperties(ht)
				expr = ht.KeyAddNew(types.STR_SCALAR, expr)

				{

					// expr.TryAddRefcount()

				}
			}
		}
	}
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
