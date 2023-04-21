package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

func ZEND_INIT_ARRAY_SPEC_CONST_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CONST_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CONST_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			//// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_TMP_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			//// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_TMP_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_TMP_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_VAR_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CONST_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_VAR_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_VAR_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()

	/* Explicitly initialize array as not-packed if flag is set */

	{
		array.SetArray(types.NewArray(0))
		return ZEND_VM_NEXT_OPCODE(executeData, opline)
	}
}
func ZEND_INIT_ARRAY_SPEC_CV_CONST_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CV_UNUSED_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(executeData)
	}

}
func ZEND_INIT_ARRAY_SPEC_CV_CV_HANDLER(executeData *ZendExecuteData) int {
	var array *types.Zval
	var size uint32
	var opline *ZendOp = executeData.GetOpline()
	array = opline.Result()
	{
		size = opline.GetExtendedValue() >> ZEND_ARRAY_SIZE_SHIFT
		array.SetArray(types.NewArray(size))

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & ZEND_ARRAY_NOT_PACKED) != 0 {
			// types.ZendHashRealInitMixed(array.Array())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER(executeData)
	}

}
