package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Const2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Const2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Const2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Const2()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Op2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Op2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Op2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Op2()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Cv2OrUndef()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Cv2OrUndef()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Cv2OrUndef()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, executeData)
	property = opline.Cv2OrUndef()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Const2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	value = object.Object().WritePropertyEx(property, value)
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Const2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Const2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Const2()
	value = opline.Offset(1).Cv1OrUndef()

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Op2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())

	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Op2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Op2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Op2()
	value = opline.Offset(1).Cv1OrUndef()

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Cv2OrUndef()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())

	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Cv2OrUndef()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Cv2OrUndef()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = executeData.ThisObjectZval()
	if object == nil {
		return zend_this_not_in_object_context_helper_SPEC(executeData)
	}
	property = opline.Cv2OrUndef()
	value = opline.Offset(1).Cv1OrUndef()

	value = types.ZVAL_DEREF(value)
	value = object.Object().WritePropertyEx(property, value)

	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */
	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Const2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Const2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	// ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.Op1()
	property = opline.Const2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	// ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	var tmp types.Zval
	object = opline.Op1()
	property = opline.Const2()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Op2()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Op2()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Op2()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Op2()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Cv2OrUndef()
	value = RT_CONSTANT(opline+1, (opline + 1).GetOp1())
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_TMP_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Cv2OrUndef()
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_VAR_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var free_op_data ZendFreeOp
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Cv2OrUndef()
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, executeData)
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}

	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CV_HANDLER(executeData *ZendExecuteData) int {
	var opline *types.ZendOp = executeData.GetOpline()
	var object *types.Zval
	var property *types.Zval
	var value *types.Zval
	object = opline.Op1()
	property = opline.Cv2OrUndef()
	value = opline.Offset(1).Cv1OrUndef()
	if !object.IsObject() {
		if object.IsReference() && types.Z_REFVAL_P(object).IsObject() {
			object = types.Z_REFVAL_P(object)
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, executeData)
		if object == nil {
			value = UninitializedZval()
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	{
		value = types.ZVAL_DEREF(value)
	}
	value = object.Object().WritePropertyEx(property, value)
free_and_exit_assign_obj:
	if RETURN_VALUE_USED(opline) {
		types.ZVAL_COPY(opline.Result(), value)
	}
	/* assign_obj has two opcodes! */

	OPLINE = executeData.GetOpline() + 2
	return 0
}
