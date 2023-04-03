package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/zend/types"
)

func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	var offset ZendLong
	var ht *types.Array
	container = opline.GetOp1Zval()
	dim = RT_CONSTANT(opline, opline.GetOp2())
	if container.IsArray() {
	fetch_dim_r_index_array:
		if dim.IsLong() {
			offset = dim.GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetArr()
		value = ht.IndexFind(offset)
		if value == nil {
			goto fetch_dim_r_index_undef
		}
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), value)
		{
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}

	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		if dim.GetU2Extra() == ZEND_EXTRA_VALUE {
			dim++
		}
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
fetch_dim_r_index_undef:
	opline.GetResultZval().SetNull()
	ZendUndefinedOffset(offset)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var container *types.Zval
	var dim *types.Zval
	var value *types.Zval
	var offset ZendLong
	var ht *types.Array
	container = opline.GetOp1Zval()
	dim = opline.GetOp2Zval()
	if container.IsArray() {
	fetch_dim_r_index_array:
		if dim.IsLong() {
			offset = dim.GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetArr()
		value = ht.IndexFind(offset)
		if value == nil {
			goto fetch_dim_r_index_undef
		}
		types.ZVAL_COPY_DEREF(opline.GetResultZval(), value)
		{
			return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
		}

	} else if container.IsReference() {
		container = types.Z_REFVAL_P(container)
		if container.IsArray() {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, executeData)
		return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
	}
fetch_dim_r_index_undef:
	opline.GetResultZval().SetNull()
	ZendUndefinedOffset(offset)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
