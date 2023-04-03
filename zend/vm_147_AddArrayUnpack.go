package zend

func ZEND_ADD_ARRAY_UNPACK_SPEC_HANDLER(executeData *ZendExecuteData) int {
	var opline *ZendOp = executeData.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *types.Zval
	op1 = GetZvalPtr(opline.GetOp1Type(), opline.GetOp1(), &free_op1, BP_VAR_R)
add_unpack_again:
	if op1.IsArray() {
		var ht *types.Array = op1.GetArr()
		var val *types.Zval
		var key *types.String
		var __ht *types.Array = ht
		for _, _p := range __ht.ForeachData() {
			var _z *types.Zval = _p.GetVal()

			key = _p.GetKey()
			val = _z
			if key != nil {
				faults.ThrowError(nil, "Cannot unpack array with string keys")
				FREE_OP(free_op1)
				return 0
			} else {
				if val.IsReference() && val.GetRefcount() == 1 {
					val = types.Z_REFVAL_P(val)
				}
				val.TryAddRefcount()
				if opline.GetResultZval().GetArr().NextIndexInsert(val) == nil {
					ZendCannotAddElement()
					ZvalPtrDtorNogc(val)
					break
				}
			}
		}
	} else if op1.IsObject() {
		var ce *types.ClassEntry = types.Z_OBJCE_P(op1)
		var iter *ZendObjectIterator
		if ce == nil || ce.GetGetIterator() == nil {
			faults.ThrowError(nil, "Only arrays and Traversables can be unpacked")
		} else {
			iter = ce.GetGetIterator()(ce, op1, 0)
			if iter == nil {
				FREE_OP(free_op1)
				if EG__().GetException() == nil {
					faults.ThrowExceptionEx(nil, 0, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				return 0
			}
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
			}
			for iter.GetFuncs().GetValid()(iter) == types.SUCCESS {
				var val *types.Zval
				if EG__().GetException() != nil {
					break
				}
				val = iter.GetFuncs().GetGetCurrentData()(iter)
				if EG__().GetException() != nil {
					break
				}
				if iter.GetFuncs().GetGetCurrentKey() != nil {
					var key types.Zval
					iter.GetFuncs().GetGetCurrentKey()(iter, &key)
					if EG__().GetException() != nil {
						break
					}
					if key.GetType() != types.IS_LONG {
						faults.ThrowError(nil, b.Cond(key.IsString(), "Cannot unpack Traversable with string keys", "Cannot unpack Traversable with non-integer keys"))
						ZvalPtrDtor(&key)
						break
					}
				}
				val = types.ZVAL_DEREF(val)
				val.TryAddRefcount()
				if opline.GetResultZval().GetArr().NextIndexInsert(val) == nil {
					ZendCannotAddElement()
					ZvalPtrDtorNogc(val)
				}
				iter.GetFuncs().GetMoveForward()(iter)
			}
			ZendIteratorDtor(iter)
		}
	} else if op1.IsReference() {
		op1 = types.Z_REFVAL_P(op1)
		goto add_unpack_again
	} else {
		faults.ThrowError(nil, "Only arrays and Traversables can be unpacked")
	}
	FREE_OP(free_op1)
	return ZEND_VM_NEXT_OPCODE_CHECK_EXCEPTION(executeData)
}
