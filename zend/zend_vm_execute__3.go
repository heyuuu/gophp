// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

func ZEND_UNSET_OBJ_SPEC_VAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var offset *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		if 1<<2 != 0 && container.GetType() != 8 {
			if container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() != 8 {
					if 1<<2 == 1<<3 && container.GetType() == 0 {
						_zvalUndefinedOp1(execute_data)
					}
					break
				}
			} else {
				break
			}
		}
		container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_VAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<2 != 0 {
		var free_op1 ZendFreeOp
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 2 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<2 == 1<<2 {
						r.Assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<2 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<2 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<2&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				ZvalPtrDtorNogc(free_op1)
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	if 1<<0 != 0 {
		var key *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

		/* Consts, temporary variables and references need copying */

		if 1<<0 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<0 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<0&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_IN_ARRAY_SPEC_VAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var ht *HashTable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	var result *Zval
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if op1.GetType() == 6 {
		result = ZendHashFindEx(ht, op1.GetValue().GetStr(), 1<<2 == 1<<0)
	} else if opline.GetExtendedValue() != 0 {
		if op1.GetType() == 4 {
			result = ZendHashIndexFind(ht, op1.GetValue().GetLval())
		} else {
			result = nil
		}
	} else if op1.GetType() <= 2 {
		result = ZendHashFindEx(ht, ZendEmptyString, 1)
	} else {
		var key *ZendString
		var key_tmp Zval
		var result_tmp Zval
		var val *Zval
		result = nil
		for {
			var __ht *HashTable = ht
			var _p *Bucket = __ht.GetArData()
			var _end *Bucket = _p + __ht.GetNNumUsed()
			for ; _p != _end; _p++ {
				var _z *Zval = &_p.val

				if _z.GetType() == 0 {
					continue
				}
				key = _p.GetKey()
				val = _z
				var __z *Zval = &key_tmp
				var __s *ZendString = key
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				CompareFunction(&result_tmp, op1, &key_tmp)
				if result_tmp.GetValue().GetLval() == 0 {
					result = val
					break
				}
			}
			break
		}
	}
	ZvalPtrDtorNogc(free_op1)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != nil {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == nil {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != nil {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto assign_op_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	assign_op_object:

		/* here we are sure we are dealing with an object */

		if (1<<1 | 1<<2) == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline + 1).GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				var orig_zptr *Zval = zptr
				var ref *ZendReference
				for {
					if zptr.GetType() == 10 {
						ref = zptr.GetValue().GetRef()
						zptr = &(*zptr).value.GetRef().GetVal()
						if ref.GetSources().GetPtr() != nil {
							ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
							break
						}
					}
					if (1<<1 | 1<<2) == 1<<0 {
						prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
					} else {
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, execute_data)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = zptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, execute_data)
		}
		break
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMP|VAR|CV, UNUSED|CONST|TMPVAR) */

func ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	var container *Zval
	var dim *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if container.GetType() == 7 {
	assign_dim_op_array:
		var _zv *Zval = container
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
	assign_dim_op_new_array:
		dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		if (1<<1 | 1<<2) == 0 {
			var_ptr = ZendHashNextIndexInsert(container.GetValue().GetArr(), &EG.uninitialized_zval)
			if var_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_op_ret_null
			}
		} else {
			if (1<<1 | 1<<2) == 1<<0 {
				var_ptr = zend_fetch_dimension_address_inner_RW_CONST(container.GetValue().GetArr(), dim, execute_data)
			} else {
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetValue().GetArr(), dim, execute_data)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1, execute_data, opline)
		for {
			if (1<<1|1<<2) != 0 && var_ptr.GetType() == 10 {
				var ref *ZendReference = var_ptr.GetValue().GetRef()
				var_ptr = &(*var_ptr).value.GetRef().GetVal()
				if ref.GetSources().GetPtr() != nil {
					ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = var_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
		if free_op_data1 != nil {
			ZvalPtrDtorNogc(free_op_data1)
		}
	} else {
		if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto assign_dim_op_array
			}
		}
		dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		if container.GetType() == 8 {
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, execute_data)
		} else if container.GetType() <= 2 {
			if 1<<2 == 1<<3 && container.GetTypeInfo() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			var __arr *ZendArray = _zendNewArray(8)
			var __z *Zval = container
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, execute_data)
		assign_dim_op_ret_null:
			if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			}
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OP_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && var_ptr.GetType() == 15 {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
	} else {
		for {
			if var_ptr.GetType() == 10 {
				var ref *ZendReference = var_ptr.GetValue().GetRef()
				var_ptr = &(*var_ptr).value.GetRef().GetVal()
				if ref.GetSources().GetPtr() != nil {
					ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = var_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto pre_incdec_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		if (1<<1 | 1<<2) == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				if (1<<1 | 1<<2) == 1<<0 {
					prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
				} else {
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
				}
				ZendPreIncdecPropertyZval(zptr, prop_info, opline, execute_data)
			}
		} else {
			ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto post_incdec_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		if (1<<1 | 1<<2) == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			} else {
				if (1<<1 | 1<<2) == 1<<0 {
					prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
				} else {
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
				}
				ZendPostIncdecPropertyZval(zptr, prop_info, opline, execute_data)
			}
		} else {
			ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_W(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_RW(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 2 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER(execute_data)
	} else {
		if (1<<1 | 1<<2) == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_UNSET(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_RW_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 2 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var dim *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<2 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 && container.GetType() != 10 {
		ZendError(1<<3, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, 1<<1|1<<2, opline, execute_data)
	} else {
		zend_fetch_dimension_address_W(container, dim, 1<<1|1<<2, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if (1<<1|1<<2) == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<0 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<0 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<0 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<0 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if (1<<1|1<<2) == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<1 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<1 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<1 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<1 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if (1<<1|1<<2) == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<2 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<2 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<2 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<2 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if (1<<1|1<<2) == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<3 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<3 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<3 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<3 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<1 | 1<<2) == 0 {
			if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<0 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<0 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
				}
			} else if 1<<0 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			if (1<<1 | 1<<2) == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if (1<<1 | 1<<2) == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<1 | 1<<2) == 0 {
			if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<1 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<1 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
					ZvalPtrDtorNogc(free_op_data)
				}
			} else if 1<<1 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			if (1<<1 | 1<<2) == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if (1<<1 | 1<<2) == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<1 | 1<<2) == 0 {
			if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<2 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<2 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
					ZvalPtrDtorNogc(free_op_data)
				}
			} else if 1<<2 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			if (1<<1 | 1<<2) == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if (1<<1 | 1<<2) == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<1 | 1<<2) == 0 {
			if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<3 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<3 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
				}
			} else if 1<<3 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			if (1<<1 | 1<<2) == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if (1<<1 | 1<<2) == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 == 0 {
		if (1<<1 | 1<<2) == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if (1<<1 | 1<<2) == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<2 == 0 {
		if (1<<1 | 1<<2) == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if (1<<1 | 1<<2) == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<2 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
				return 0
			}
			if (1<<1 | 1<<2) != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else if 1<<2 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<2 == 1<<0 && (1<<1|1<<2) == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<2 != 1<<0 && (1<<1|1<<2) == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else if (1<<1 | 1<<2) != 0 {
		var free_op2 ZendFreeOp
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		if (1<<1 | 1<<2) != 1<<0 {
			if function_name.GetType() != 6 {
				for {
					if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
						function_name = &(*function_name).value.GetRef().GetVal()
						if function_name.GetType() == 6 {
							break
						}
					} else if (1<<1|1<<2) == 1<<3 && function_name.GetType() == 0 {
						_zvalUndefinedOp2(execute_data)
						if EG.GetException() != nil {
							return 0
						}
					}
					ZendThrowError(nil, "Function name must be a string")
					ZvalPtrDtorNogc(free_op2)
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetValue().GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetValue().GetStr(), g.CondF1((1<<1|1<<2) == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetValue().GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = ce
			slot[1] = fbc
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		if (1<<1 | 1<<2) != 1<<0 {
			ZvalPtrDtorNogc(free_op2)
		}
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if 1<<2 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<2 == 1<<2 || 1<<2 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if expr_ptr.GetType() == 10 {
			ZvalAddrefP(expr_ptr)
		} else {
			var _z *Zval = expr_ptr
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if 1<<2 == 1<<1 {

		} else if 1<<2 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<2 == 1<<3 {
			if expr_ptr.GetType() == 10 {
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
			}
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else {
			if expr_ptr.GetType() == 10 {
				var ref *ZendRefcounted = expr_ptr.GetValue().GetCounted()
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
				if ZendGcDelref(&ref.gc) == 0 {
					var _z1 *Zval = &new_expr
					var _z2 *Zval = expr_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					expr_ptr = &new_expr
					_efree(ref)
				} else if (expr_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(expr_ptr)
				}
			}
		}
	}
	if (1<<1 | 1<<2) != 0 {
		var free_op2 ZendFreeOp
		var offset *Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		var str *ZendString
		var hval ZendUlong
	add_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if (1<<1 | 1<<2) != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index
				}
			}
		str_index:
			ZendHashUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), str, expr_ptr)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index:
			ZendHashIndexUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), hval, expr_ptr)
		} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto add_again
		} else if offset.GetType() == 1 {
			str = ZendEmptyString
			goto str_index
		} else if offset.GetType() == 5 {
			hval = ZendDvalToLval(offset.GetValue().GetDval())
			goto num_index
		} else if offset.GetType() == 2 {
			hval = 0
			goto num_index
		} else if offset.GetType() == 3 {
			hval = 1
			goto num_index
		} else if offset.GetType() == 9 {
			ZendUseResourceAsOffset(offset)
			hval = offset.GetValue().GetRes().GetHandle()
			goto num_index
		} else if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			str = ZendEmptyString
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
		ZvalPtrDtorNogc(free_op2)
	} else {
		if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<2 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER(execute_data)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var offset *Zval
	var hval ZendUlong
	var key *ZendString
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		if container.GetType() == 7 {
			var ht *HashTable
		unset_dim_array:
			var _zv *Zval = container
			var _arr *ZendArray = _zv.GetValue().GetArr()
			if ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.GetTypeFlags() != 0 {
					ZendGcDelref(&_arr.gc)
				}
				var __arr *ZendArray = ZendArrayDup(_arr)
				var __z *Zval = _zv
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			}
			ht = container.GetValue().GetArr()
		offset_again:
			if offset.GetType() == 6 {
				key = offset.GetValue().GetStr()
				if (1<<1 | 1<<2) != 1<<0 {
					if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &hval) != 0 {
						goto num_index_dim
					}
				}
			str_index_dim:
				if ht == &EG.symbol_table {
					ZendDeleteGlobalVariable(key)
				} else {
					ZendHashDel(ht, key)
				}
			} else if offset.GetType() == 4 {
				hval = offset.GetValue().GetLval()
			num_index_dim:
				ZendHashIndexDel(ht, hval)
			} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
				offset = &(*offset).value.GetRef().GetVal()
				goto offset_again
			} else if offset.GetType() == 5 {
				hval = ZendDvalToLval(offset.GetValue().GetDval())
				goto num_index_dim
			} else if offset.GetType() == 1 {
				key = ZendEmptyString
				goto str_index_dim
			} else if offset.GetType() == 2 {
				hval = 0
				goto num_index_dim
			} else if offset.GetType() == 3 {
				hval = 1
				goto num_index_dim
			} else if offset.GetType() == 9 {
				hval = offset.GetValue().GetRes().GetHandle()
				goto num_index_dim
			} else if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				key = ZendEmptyString
				goto str_index_dim
			} else {
				ZendError(1<<1, "Illegal offset type in unset")
			}
			break
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto unset_dim_array
			}
		}
		if 1<<2 == 1<<3 && container.GetType() == 0 {
			container = _zvalUndefinedOp1(execute_data)
		}
		if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
			offset = _zvalUndefinedOp2(execute_data)
		}
		if container.GetType() == 8 {
			if (1<<1|1<<2) == 1<<0 && offset.GetU2Extra() == 1 {
				offset++
			}
			container.GetValue().GetObj().GetHandlers().GetUnsetDimension()(container, offset)
		} else if 1<<2 != 0 && container.GetType() == 6 {
			ZendThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var offset *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		if 1<<2 != 0 && container.GetType() != 8 {
			if container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() != 8 {
					if 1<<2 == 1<<3 && container.GetType() == 0 {
						_zvalUndefinedOp1(execute_data)
					}
					break
				}
			} else {
				break
			}
		}
		container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
		break
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_IDENTICAL_SPEC_VAR_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_VAR_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<2 != 0 {
		var free_op1 ZendFreeOp
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 2 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<2 == 1<<2 {
						r.Assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<2 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<2 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<2&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				ZvalPtrDtorNogc(free_op1)
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	if 1<<1 != 0 {
		var free_op2 ZendFreeOp
		var key *Zval = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)

		/* Consts, temporary variables and references need copying */

		if 1<<1 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<1 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<1&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<1 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_IS_IDENTICAL_SPEC_VAR_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_IS_NOT_IDENTICAL_SPEC_VAR_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsNotIdenticalFunction(op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_REF_SPEC_VAR_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var variable_ptr *Zval
	var value_ptr *Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<2 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 {
		ZendThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<2 == 1<<2 && value_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<2 == 1<<2 && opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, execute_data)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = variable_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	if free_op2 != nil {
		ZvalPtrDtorNogc(free_op2)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_VAR_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<2 != 0 {
		var free_op1 ZendFreeOp
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 2 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<2 == 1<<2 {
						r.Assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<2 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<2 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<2&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				ZvalPtrDtorNogc(free_op1)
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	if 1<<2 != 0 {
		var free_op2 ZendFreeOp
		var key *Zval = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)

		/* Consts, temporary variables and references need copying */

		if 1<<2 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<2 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<2&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
			ZvalPtrDtorNogc(free_op2)
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<2 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_ASSIGN_DIM_OP_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	var container *Zval
	var dim *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if container.GetType() == 7 {
	assign_dim_op_array:
		var _zv *Zval = container
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
	assign_dim_op_new_array:
		dim = nil
		var_ptr = ZendHashNextIndexInsert(container.GetValue().GetArr(), &EG.uninitialized_zval)
		if var_ptr == nil {
			ZendCannotAddElement()
			goto assign_dim_op_ret_null
		}
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1, execute_data, opline)

		ZendBinaryOp(var_ptr, var_ptr, value, opline)
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = var_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
		if free_op_data1 != nil {
			ZvalPtrDtorNogc(free_op_data1)
		}
	} else {
		if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto assign_dim_op_array
			}
		}
		dim = nil
		if container.GetType() == 8 {
			if 0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, execute_data)
		} else if container.GetType() <= 2 {
			if 1<<2 == 1<<3 && container.GetTypeInfo() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			var __arr *ZendArray = _zendNewArray(8)
			var __z *Zval = container
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, execute_data)
		assign_dim_op_ret_null:
			if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			}
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_FETCH_DIM_W_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_W(container, nil, 0, opline, execute_data)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_RW(container, nil, 0, opline, execute_data)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 2 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_UNUSED_HANDLER(execute_data)
	} else {
		return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		return ZEND_NULL_HANDLER(execute_data)
	}
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
			if value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
			}
		}
		variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
		if variable_ptr == nil {
			ZendCannotAddElement()
			goto assign_dim_error
		} else if 1<<0 == 1<<3 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		} else if 1<<0 == 1<<2 {
			if value != free_op_data {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else if 1<<0 == 1<<0 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = nil
			value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
			if 0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			ZendUseNewElementForString()
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = nil
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
			if value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
			}
		}
		variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
		if variable_ptr == nil {
			ZendCannotAddElement()
			goto assign_dim_error
		} else if 1<<1 == 1<<3 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		} else if 1<<1 == 1<<2 {
			if value != free_op_data {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if 1<<1 == 1<<0 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = nil
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			ZendUseNewElementForString()
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
			if value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
			}
		}
		variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
		if variable_ptr == nil {
			ZendCannotAddElement()
			goto assign_dim_error
		} else if 1<<2 == 1<<3 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		} else if 1<<2 == 1<<2 {
			if value != free_op_data {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if 1<<2 == 1<<0 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = nil
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			ZendUseNewElementForString()
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = nil
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
			if value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
			}
		}
		variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
		if variable_ptr == nil {
			ZendCannotAddElement()
			goto assign_dim_error
		} else if 1<<3 == 1<<3 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		} else if 1<<3 == 1<<2 {
			if value != free_op_data {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else if 1<<3 == 1<<0 {
			if value.GetTypeFlags() != 0 {
				ZvalAddrefP(value)
			}
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = nil
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
			if 0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			ZendUseNewElementForString()
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = nil
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<2 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				return 0
			}
			if 0 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else if 1<<2 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<2 == 1<<0 && 0 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<2 != 1<<0 && 0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if 1<<2 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_VERIFY_RETURN_TYPE_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<2 == 0 {
		ZendVerifyMissingReturnType(execute_data.GetFunc(), (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum())))
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_NEW_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var result *Zval
	var constructor *ZendFunction
	var ce *ZendClassEntry
	var call *ZendExecuteData
	if 1<<2 == 1<<0 {
		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetOp2().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetOp2().GetNum()))[0] = ce
		}
	} else if 1<<2 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if ObjectInitEx(result, ce) != SUCCESS {
		result.SetTypeInfo(0)
		return 0
	}
	constructor = result.GetValue().GetObj().GetHandlers().GetGetConstructor()(result.GetValue().GetObj())
	if constructor == nil {
		if EG.GetException() != nil {
			return 0
		}

		/* If there are no arguments, skip over the DO_FCALL opcode. We check if the next
		 * opcode is DO_FCALL in case EXT instructions are used. */

		if opline.GetExtendedValue() == 0 && (opline+1).GetOpcode() == 60 {
			execute_data.SetOpline(execute_data.GetOpline() + 2)
			return 0
		}

		/* Perform a dummy function call */

		call = ZendVmStackPushCallFrame(0<<16, (*ZendFunction)(&ZendPassFunction), opline.GetExtendedValue(), nil)

		/* Perform a dummy function call */

	} else {
		if constructor.GetType() == 2 && !(g.CondF((uintptr_t(&constructor.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&constructor.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&constructor.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&constructor.op_array)
		}

		/* We are not handling overloaded classes right now */

		call = ZendVmStackPushCallFrame(0<<16|1<<21|(8|1<<0<<8|1<<1<<8), constructor, opline.GetExtendedValue(), result.GetValue().GetObj())
		ZvalAddrefP(result)
	}
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<2 == 1<<2 || 1<<2 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if expr_ptr.GetType() == 10 {
			ZvalAddrefP(expr_ptr)
		} else {
			var _z *Zval = expr_ptr
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if 1<<2 == 1<<1 {

		} else if 1<<2 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<2 == 1<<3 {
			if expr_ptr.GetType() == 10 {
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
			}
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else {
			if expr_ptr.GetType() == 10 {
				var ref *ZendRefcounted = expr_ptr.GetValue().GetCounted()
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
				if ZendGcDelref(&ref.gc) == 0 {
					var _z1 *Zval = &new_expr
					var _z2 *Zval = expr_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					expr_ptr = &new_expr
					_efree(ref)
				} else if (expr_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(expr_ptr)
				}
			}
		}
	}
	if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
		ZendCannotAddElement()
		ZvalPtrDtorNogc(expr_ptr)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<2 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER(execute_data)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 10 {
		if ZvalRefcountP(var_ptr) == 1 {
			var _z *Zval = var_ptr
			var ref *ZendReference
			r.Assert(_z.GetType() == 10)
			ref = _z.GetValue().GetRef()
			var _z1 *Zval = _z
			var _z2 *Zval = &ref.val
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_efree(ref)
		}
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_YIELD_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<2 != 0 {
		var free_op1 ZendFreeOp
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 2 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<2 == 1<<2 {
						r.Assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<2 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<2 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<2&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				ZvalPtrDtorNogc(free_op1)
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	/* If no key was specified we use auto-increment keys */

	generator.GetLargestUsedIntegerKey()++
	var __z *Zval = &generator.key
	__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
	__z.SetTypeInfo(4)
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_MAKE_REF_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<2 == 1<<3 {
		if op1.GetType() == 0 {
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			_ref.GetSources().SetPtr(nil)
			op1.GetValue().SetRef(_ref)
			op1.SetTypeInfo(10 | 1<<0<<8)
			ZvalSetRefcountP(op1, 2)
			&(*op1).value.GetRef().GetVal().u1.type_info = 1
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			__z.GetValue().SetRef(op1.GetValue().GetRef())
			__z.SetTypeInfo(10 | 1<<0<<8)
		} else {
			if op1.GetType() == 10 {
				ZvalAddrefP(op1)
			} else {
				var _z *Zval = op1
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 2)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = _z
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				_z.GetValue().SetRef(_ref)
				_z.SetTypeInfo(10 | 1<<0<<8)
			}
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			__z.GetValue().SetRef(op1.GetValue().GetRef())
			__z.SetTypeInfo(10 | 1<<0<<8)
		}
	} else if op1.GetType() == 13 {
		op1 = op1.GetValue().GetZv()
		if op1.GetType() != 10 {
			var _z *Zval = op1
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		} else {
			ZendGcAddref(&(op1.GetValue().GetRef()).gc)
		}
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetRef(op1.GetValue().GetRef())
		__z.SetTypeInfo(10 | 1<<0<<8)
	} else {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = op1
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_GET_TYPE_SPEC_VAR_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var type_ *ZendString
	op1 = _getZvalPtrVarDeref(opline.GetOp1().GetVar(), &free_op1, execute_data)
	type_ = ZendZvalGetType(op1)
	if type_ != nil {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = type_
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6)
	} else {
		var _s *byte = "unknown type"
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = ZendStringInit(_s, strlen(_s), 0)
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
	}
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto assign_op_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	assign_op_object:

		/* here we are sure we are dealing with an object */

		if 1<<3 == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline + 1).GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				var orig_zptr *Zval = zptr
				var ref *ZendReference
				for {
					if zptr.GetType() == 10 {
						ref = zptr.GetValue().GetRef()
						zptr = &(*zptr).value.GetRef().GetVal()
						if ref.GetSources().GetPtr() != nil {
							ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
							break
						}
					}
					if 1<<3 == 1<<0 {
						prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
					} else {
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, execute_data)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = zptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, execute_data)
		}
		break
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMP|VAR|CV, UNUSED|CONST|TMPVAR) */

func ZEND_ASSIGN_DIM_OP_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data1 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	var container *Zval
	var dim *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if container.GetType() == 7 {
	assign_dim_op_array:
		var _zv *Zval = container
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
	assign_dim_op_new_array:
		dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
		if 1<<3 == 0 {
			var_ptr = ZendHashNextIndexInsert(container.GetValue().GetArr(), &EG.uninitialized_zval)
			if var_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_op_ret_null
			}
		} else {
			if 1<<3 == 1<<0 {
				var_ptr = zend_fetch_dimension_address_inner_RW_CONST(container.GetValue().GetArr(), dim, execute_data)
			} else {
				var_ptr = zend_fetch_dimension_address_inner_RW(container.GetValue().GetArr(), dim, execute_data)
			}
			if var_ptr == nil {
				goto assign_dim_op_ret_null
			}
		}
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data1, execute_data, opline)
		for {
			if 1<<3 != 0 && var_ptr.GetType() == 10 {
				var ref *ZendReference = var_ptr.GetValue().GetRef()
				var_ptr = &(*var_ptr).value.GetRef().GetVal()
				if ref.GetSources().GetPtr() != nil {
					ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = var_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
		if free_op_data1 != nil {
			ZvalPtrDtorNogc(free_op_data1)
		}
	} else {
		if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto assign_dim_op_array
			}
		}
		dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		if container.GetType() == 8 {
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, execute_data)
		} else if container.GetType() <= 2 {
			if 1<<2 == 1<<3 && container.GetTypeInfo() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			var __arr *ZendArray = _zendNewArray(8)
			var __z *Zval = container
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			goto assign_dim_op_new_array
		} else {
			ZendBinaryAssignOpDimSlow(container, dim, opline, execute_data)
		assign_dim_op_ret_null:
			if ((opline + 1).GetOp1Type() & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			}
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OP_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	var_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && var_ptr.GetType() == 15 {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
	} else {
		for {
			if var_ptr.GetType() == 10 {
				var ref *ZendReference = var_ptr.GetValue().GetRef()
				var_ptr = &(*var_ptr).value.GetRef().GetVal()
				if ref.GetSources().GetPtr() != nil {
					ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
					break
				}
			}
			ZendBinaryOp(var_ptr, var_ptr, value, opline)
			break
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = var_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	for {
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto pre_incdec_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	pre_incdec_object:

		/* here we are sure we are dealing with an object */

		if 1<<3 == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				if 1<<3 == 1<<0 {
					prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
				} else {
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
				}
				ZendPreIncdecPropertyZval(zptr, prop_info, opline, execute_data)
			}
		} else {
			ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
		}
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	for {
		if 1<<2 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto post_incdec_object
			}
			if 1<<2 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
	post_incdec_object:

		/* here we are sure we are dealing with an object */

		if 1<<3 == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			} else {
				if 1<<3 == 1<<0 {
					prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
				} else {
					prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
				}
				ZendPostIncdecPropertyZval(zptr, prop_info, opline, execute_data)
			}
		} else {
			ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
		}
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_W(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_RW_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_RW(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 2 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER(execute_data)
	} else {
		if 1<<3 == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_UNSET(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	if 1<<2 == 1<<2 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_RW_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 2 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<2, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
	if 1<<2 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_LIST_W_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var dim *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<2 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 && container.GetType() != 10 {
		ZendError(1<<3, "Attempting to set reference to non referenceable value")
		zend_fetch_dimension_address_LIST_r(container, dim, 1<<3, opline, execute_data)
	} else {
		zend_fetch_dimension_address_W(container, dim, 1<<3, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if 1<<3 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<0 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<0 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<0 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<0 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if 1<<3 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<1 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<1 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<1 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<1 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if 1<<3 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<2 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<2 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<2 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<2 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<2 != 0 && object.GetType() != 8 {
		if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
			object = &(*object).value.GetRef().GetVal()
			goto assign_object
		}
		object = MakeRealObject(object, property, opline, execute_data)
		if object == nil {
			value = &EG.uninitialized_zval
			goto free_and_exit_assign_obj
		}
	}
assign_object:
	if 1<<3 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<3 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<3 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<3 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<3 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<3 == 0 {
			if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<0 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<0 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
				}
			} else if 1<<0 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
			if 1<<3 == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
			value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<3 == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<3 != 0 {

	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<3 == 0 {
			if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<1 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<1 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
					ZvalPtrDtorNogc(free_op_data)
				}
			} else if 1<<1 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
			if 1<<3 == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<3 == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<3 != 0 {

	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<3 == 0 {
			if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<2 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<2 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
					ZvalPtrDtorNogc(free_op_data)
				}
			} else if 1<<2 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
			if 1<<3 == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<3 == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<3 != 0 {

	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	orig_object_ptr = object_ptr
	if object_ptr.GetType() == 7 {
	try_assign_dim_array:
		value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
		var _zv *Zval = object_ptr
		var _arr *ZendArray = _zv.GetValue().GetArr()
		if ZendGcRefcount(&_arr.gc) > 1 {
			if _zv.GetTypeFlags() != 0 {
				ZendGcDelref(&_arr.gc)
			}
			var __arr *ZendArray = ZendArrayDup(_arr)
			var __z *Zval = _zv
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		}
		if 1<<3 == 0 {
			if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
				if value.GetType() == 10 {
					value = &(*value).value.GetRef().GetVal()
				}
			}
			variable_ptr = ZendHashNextIndexInsert(object_ptr.GetValue().GetArr(), value)
			if variable_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_error
			} else if 1<<3 == 1<<3 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			} else if 1<<3 == 1<<2 {
				if value != free_op_data {
					if value.GetTypeFlags() != 0 {
						ZvalAddrefP(value)
					}
				}
			} else if 1<<3 == 1<<0 {
				if value.GetTypeFlags() != 0 {
					ZvalAddrefP(value)
				}
			}
		} else {
			dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
			if 1<<3 == 1<<0 {
				variable_ptr = zend_fetch_dimension_address_inner_W_CONST(object_ptr.GetValue().GetArr(), dim, execute_data)
			} else {
				variable_ptr = zend_fetch_dimension_address_inner_W(object_ptr.GetValue().GetArr(), dim, execute_data)
			}
			if variable_ptr == nil {
				goto assign_dim_error
			}
			value = ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		}
		if opline.GetResultType() != 0 {
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		}
	} else {
		if object_ptr.GetType() == 10 {
			object_ptr = &(*object_ptr).value.GetRef().GetVal()
			if object_ptr.GetType() == 7 {
				goto try_assign_dim_array
			}
		}
		if object_ptr.GetType() == 8 {
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<3 == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				var __arr *ZendArray = _zendNewArray(8)
				var __z *Zval = object_ptr
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
				goto try_assign_dim_array
			}
		} else {
			if 1<<2 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<3 != 0 {

	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_REF_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var variable_ptr *Zval
	var value_ptr *Zval
	value_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 1<<2 && variable_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<2 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 {
		ZendThrowError(nil, "Cannot assign by reference to an array dimension of an object")
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<3 == 1<<2 && value_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<3 == 1<<2 && opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
		variable_ptr = ZendWrongAssignToVariableReference(variable_ptr, value_ptr, opline, execute_data)
	} else {
		ZendAssignToVariableReference(variable_ptr, value_ptr)
	}
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = variable_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<2 == 0 {
		if 1<<3 == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if 1<<3 == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<2 == 0 {
		if 1<<3 == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if 1<<3 == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<2 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				return 0
			}
			if 1<<3 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else if 1<<2 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<2 == 1<<0 && 1<<3 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<2 != 1<<0 && 1<<3 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else if 1<<3 != 0 {
		function_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
		if 1<<3 != 1<<0 {
			if function_name.GetType() != 6 {
				for {
					if (1<<3&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
						function_name = &(*function_name).value.GetRef().GetVal()
						if function_name.GetType() == 6 {
							break
						}
					} else if 1<<3 == 1<<3 && function_name.GetType() == 0 {
						_zvalUndefinedOp2(execute_data)
						if EG.GetException() != nil {
							return 0
						}
					}
					ZendThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetValue().GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetValue().GetStr(), g.CondF1(1<<3 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<3 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = ce
			slot[1] = fbc
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		if 1<<3 != 1<<0 {

		}
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if 1<<2 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<2 == 1<<2 || 1<<2 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if expr_ptr.GetType() == 10 {
			ZvalAddrefP(expr_ptr)
		} else {
			var _z *Zval = expr_ptr
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 2)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = _z
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			_z.GetValue().SetRef(_ref)
			_z.SetTypeInfo(10 | 1<<0<<8)
		}
		if free_op1 != nil {
			ZvalPtrDtorNogc(free_op1)
		}
	} else {
		expr_ptr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
		if 1<<2 == 1<<1 {

		} else if 1<<2 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<2 == 1<<3 {
			if expr_ptr.GetType() == 10 {
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
			}
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else {
			if expr_ptr.GetType() == 10 {
				var ref *ZendRefcounted = expr_ptr.GetValue().GetCounted()
				expr_ptr = &(*expr_ptr).value.GetRef().GetVal()
				if ZendGcDelref(&ref.gc) == 0 {
					var _z1 *Zval = &new_expr
					var _z2 *Zval = expr_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					expr_ptr = &new_expr
					_efree(ref)
				} else if (expr_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(expr_ptr)
				}
			}
		}
	}
	if 1<<3 != 0 {
		var offset *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
		var str *ZendString
		var hval ZendUlong
	add_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<3 != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index
				}
			}
		str_index:
			ZendHashUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), str, expr_ptr)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index:
			ZendHashIndexUpdate((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), hval, expr_ptr)
		} else if (1<<3&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto add_again
		} else if offset.GetType() == 1 {
			str = ZendEmptyString
			goto str_index
		} else if offset.GetType() == 5 {
			hval = ZendDvalToLval(offset.GetValue().GetDval())
			goto num_index
		} else if offset.GetType() == 2 {
			hval = 0
			goto num_index
		} else if offset.GetType() == 3 {
			hval = 1
			goto num_index
		} else if offset.GetType() == 9 {
			ZendUseResourceAsOffset(offset)
			hval = offset.GetValue().GetRes().GetHandle()
			goto num_index
		} else if 1<<3 == 1<<3 && offset.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
			str = ZendEmptyString
			goto str_index
		} else {
			ZendIllegalOffset()
			ZvalPtrDtorNogc(expr_ptr)
		}
	} else {
		if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
			ZendCannotAddElement()
			ZvalPtrDtorNogc(expr_ptr)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<2 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER(execute_data)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_UNSET_DIM_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var offset *Zval
	var hval ZendUlong
	var key *ZendString
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	offset = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	for {
		if container.GetType() == 7 {
			var ht *HashTable
		unset_dim_array:
			var _zv *Zval = container
			var _arr *ZendArray = _zv.GetValue().GetArr()
			if ZendGcRefcount(&_arr.gc) > 1 {
				if _zv.GetTypeFlags() != 0 {
					ZendGcDelref(&_arr.gc)
				}
				var __arr *ZendArray = ZendArrayDup(_arr)
				var __z *Zval = _zv
				__z.GetValue().SetArr(__arr)
				__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
			}
			ht = container.GetValue().GetArr()
		offset_again:
			if offset.GetType() == 6 {
				key = offset.GetValue().GetStr()
				if 1<<3 != 1<<0 {
					if _zendHandleNumericStr(key.GetVal(), key.GetLen(), &hval) != 0 {
						goto num_index_dim
					}
				}
			str_index_dim:
				if ht == &EG.symbol_table {
					ZendDeleteGlobalVariable(key)
				} else {
					ZendHashDel(ht, key)
				}
			} else if offset.GetType() == 4 {
				hval = offset.GetValue().GetLval()
			num_index_dim:
				ZendHashIndexDel(ht, hval)
			} else if (1<<3&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
				offset = &(*offset).value.GetRef().GetVal()
				goto offset_again
			} else if offset.GetType() == 5 {
				hval = ZendDvalToLval(offset.GetValue().GetDval())
				goto num_index_dim
			} else if offset.GetType() == 1 {
				key = ZendEmptyString
				goto str_index_dim
			} else if offset.GetType() == 2 {
				hval = 0
				goto num_index_dim
			} else if offset.GetType() == 3 {
				hval = 1
				goto num_index_dim
			} else if offset.GetType() == 9 {
				hval = offset.GetValue().GetRes().GetHandle()
				goto num_index_dim
			} else if 1<<3 == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				key = ZendEmptyString
				goto str_index_dim
			} else {
				ZendError(1<<1, "Illegal offset type in unset")
			}
			break
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto unset_dim_array
			}
		}
		if 1<<2 == 1<<3 && container.GetType() == 0 {
			container = _zvalUndefinedOp1(execute_data)
		}
		if 1<<3 == 1<<3 && offset.GetType() == 0 {
			offset = _zvalUndefinedOp2(execute_data)
		}
		if container.GetType() == 8 {
			if 1<<3 == 1<<0 && offset.GetU2Extra() == 1 {
				offset++
			}
			container.GetValue().GetObj().GetHandlers().GetUnsetDimension()(container, offset)
		} else if 1<<2 != 0 && container.GetType() == 6 {
			ZendThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var offset *Zval
	container = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if 1<<2 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	for {
		if 1<<2 != 0 && container.GetType() != 8 {
			if container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() != 8 {
					if 1<<2 == 1<<3 && container.GetType() == 0 {
						_zvalUndefinedOp1(execute_data)
					}
					break
				}
			} else {
				break
			}
		}
		container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
		break
	}
	if free_op1 != nil {
		ZvalPtrDtorNogc(free_op1)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_VAR_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	if 1<<2 != 0 {
		var free_op1 ZendFreeOp
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 2 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _getZvalPtrPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<2 == 1<<2 {
						r.Assert(value_ptr != &EG.uninitialized_zval)
						if opline.GetExtendedValue() == 1<<0 && value_ptr.GetType() != 10 {
							ZendError(1<<3, "Only variable references should be yielded by reference")
							var _z1 *Zval = &generator.value
							var _z2 *Zval = value_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (_t & 0xff00) != 0 {
								ZendGcAddref(&_gc.gc)
							}
							break
						}
					}
					if value_ptr.GetType() == 10 {
						ZvalAddrefP(value_ptr)
					} else {
						var _z *Zval = value_ptr
						var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
						ZendGcSetRefcount(&_ref.gc, 2)
						_ref.GetGc().SetTypeInfo(10)
						var _z1 *Zval = &_ref.val
						var _z2 *Zval = _z
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						_ref.GetSources().SetPtr(nil)
						_z.GetValue().SetRef(_ref)
						_z.SetTypeInfo(10 | 1<<0<<8)
					}
					var __z *Zval = &generator.value
					__z.GetValue().SetRef(value_ptr.GetValue().GetRef())
					__z.SetTypeInfo(10 | 1<<0<<8)
					break
				}
				if free_op1 != nil {
					ZvalPtrDtorNogc(free_op1)
				}
			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<2 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<2 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<2&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				ZvalPtrDtorNogc(free_op1)
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<2 == 1<<3 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				}
			}

			/* Consts, temporary variables and references need copying */

		}
	} else {

		/* If no value was specified yield null */

		&generator.value.u1.type_info = 1

		/* If no value was specified yield null */

	}

	/* Set the new yielded key */

	if 1<<3 != 0 {
		var key *Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)

		/* Consts, temporary variables and references need copying */

		if 1<<3 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<3 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<3&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<3 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array *Zval
	var value *Zval
	var variable_ptr *Zval
	var value_type uint32
	var fe_ht *HashTable
	var pos HashPosition
	var p *Bucket
	array = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	fe_ht = array.GetValue().GetArr()
	pos = array.GetFePos()
	p = fe_ht.GetArData() + pos
	for true {
		if pos >= fe_ht.GetNNumUsed() {

			/* reached end of iteration */

			execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		value = &p.val
		value_type = value.GetTypeInfo()
		if value_type != 0 {
			if value_type == 13 {
				value = value.GetValue().GetZv()
				value_type = value.GetTypeInfo()
				if value_type != 0 {
					break
				}
			} else {
				break
			}
		}
		pos++
		p++
	}
	array.SetFePos(pos + 1)

	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array *Zval
	var value *Zval
	var variable_ptr *Zval
	var value_type uint32
	var fe_ht *HashTable
	var pos HashPosition
	var p *Bucket
	array = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	fe_ht = array.GetValue().GetArr()
	pos = array.GetFePos()
	p = fe_ht.GetArData() + pos
	for true {
		if pos >= fe_ht.GetNNumUsed() {

			/* reached end of iteration */

			execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
			if EG.GetVmInterrupt() != 0 {
				return zend_interrupt_helper_SPEC(execute_data)
			}
			return 0
		}
		value = &p.val
		value_type = value.GetTypeInfo()
		if value_type != 0 {
			if value_type == 13 {
				value = value.GetValue().GetZv()
				value_type = value.GetTypeInfo()
				if value_type != 0 {
					break
				}
			} else {
				break
			}
		}
		pos++
		p++
	}
	array.SetFePos(pos + 1)
	if p.GetKey() == nil {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(p.GetH())
		__z.SetTypeInfo(4)
	} else {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = p.GetKey()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	}
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1|2) != 0 {
		execute_data.GetCall().GetThis().SetTypeInfo(execute_data.GetCall().GetThis().GetTypeInfo() | 1<<31)
	} else {
		execute_data.GetCall().GetThis().SetTypeInfo(execute_data.GetCall().GetThis().GetTypeInfo() &^ (1 << 31))
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_QUICK_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg_num uint32 = opline.GetOp2().GetNum()
	if (execute_data.GetCall().GetFunc().GetQuickArgFlags() >> (arg_num + 3) * 2 & (1 | 2)) != 0 {
		execute_data.GetCall().GetThis().SetTypeInfo(execute_data.GetCall().GetThis().GetTypeInfo() | 1<<31)
	} else {
		execute_data.GetCall().GetThis().SetTypeInfo(execute_data.GetCall().GetThis().GetTypeInfo() &^ (1 << 31))
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_CLONE_SPEC_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var obj *Zval
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = &(execute_data.GetThis())
	if obj.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	for {
		if 0 == 1<<0 {
			if (0&(1<<2|1<<3)) != 0 && obj.GetType() == 10 {
				obj = &(*obj).value.GetRef().GetVal()
				if obj.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			if 0 == 1<<3 && obj.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "__clone method called on non-object")
			return 0
		}
		break
	}
	ce = obj.GetValue().GetObj().GetCe()
	clone = ce.GetClone()
	clone_call = obj.GetValue().GetObj().GetHandlers().GetCloneObj()
	if clone_call == nil {
		ZendThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	if clone != nil && (clone.GetFnFlags()&1<<0) == 0 {
		scope = execute_data.GetFunc().GetOpArray().GetScope()
		if clone.GetScope() != scope {
			if (clone.GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(clone.GetPrototype() != nil, func() *ZendClassEntry { return clone.GetPrototype().GetScope() }, func() *ZendClassEntry { return clone.GetScope() }), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
		}
	}
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetObj(clone_call(obj))
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var fetch_type uint32
	var called_scope *ZendClassEntry
	var scope *ZendClassEntry
	var opline *ZendOp = execute_data.GetOpline()
	fetch_type = opline.GetOp1().GetNum()
	scope = execute_data.GetFunc().GetOpArray().GetScope()
	if scope == nil {
		ZendThrowError(nil, "Cannot use \"%s\" when no class scope is active", g.Cond(g.Cond(fetch_type == 1, "self", fetch_type == 2), "parent", "static"))
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	switch fetch_type {
	case 1:
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = scope.GetName()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	case 2:
		if scope.parent == nil {
			ZendThrowError(nil, "Cannot use \"parent\" when current class scope has no parent")
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = scope.parent.name
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	case 3:
		if execute_data.GetThis().GetType() == 8 {
			called_scope = execute_data.GetThis().GetValue().GetObj().GetCe()
		} else {
			called_scope = execute_data.GetThis().GetValue().GetCe()
		}
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = called_scope.GetName()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	default:
		break
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)

	assign_op_object:

		/* here we are sure we are dealing with an object */

		if 1<<0 == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline + 1).GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				var orig_zptr *Zval = zptr
				var ref *ZendReference
				for {
					if zptr.GetType() == 10 {
						ref = zptr.GetValue().GetRef()
						zptr = &(*zptr).value.GetRef().GetVal()
						if ref.GetSources().GetPtr() != nil {
							ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
							break
						}
					}
					if 1<<0 == 1<<0 {
						prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
					} else {
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, execute_data)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = zptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, execute_data)
		}
		break
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMP|VAR|CV, UNUSED|CONST|TMPVAR) */

func ZEND_PRE_INC_OBJ_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

pre_incdec_object:

	/* here we are sure we are dealing with an object */

	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
	} else {
		cache_slot = nil
	}
	if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
		if zptr.GetType() == 15 {
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		} else {
			if 1<<0 == 1<<0 {
				prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
			} else {
				prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
			}
			ZendPreIncdecPropertyZval(zptr, prop_info, opline, execute_data)
		}
	} else {
		ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

post_incdec_object:

	/* here we are sure we are dealing with an object */

	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
	} else {
		cache_slot = nil
	}
	if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
		if zptr.GetType() == 15 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		} else {
			if 1<<0 == 1<<0 {
				prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
			} else {
				prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
			}
			ZendPostIncdecPropertyZval(zptr, prop_info, opline, execute_data)
		}
	} else {
		ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 0 == 1<<0 {
		for {
			if (0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 0 == 1<<3 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if 1<<0 == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			ZendWrongPropertyRead(offset)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^1)))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetTypeInfo() != 0 {
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
					fetch_obj_r_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						r.Assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_r_copy
							} else {
								goto fetch_obj_r_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	} else if 1<<0 == 1<<3 && offset.GetTypeInfo() == 0 {
		_zvalUndefinedOp2(execute_data)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 0, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_r_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_INLINE_HANDLER(execute_data)
}
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	if 0 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	if 0 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 0 == 1<<0 {
		for {
			if (0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if 1<<0 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetType() != 0 {
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
					fetch_obj_is_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						r.Assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_is_copy
							} else {
								goto fetch_obj_is_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
						goto fetch_obj_is_fast_copy
					}
				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 3, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_is_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
	if 0 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)

assign_object:
	if 1<<0 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<0 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<0 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<0 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<0 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<0 == 1<<3 || 1<<0 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

assign_object:
	if 1<<0 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<1 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<1 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<1 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<1 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<1 == 1<<3 || 1<<1 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

assign_object:
	if 1<<0 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<2 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<2 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<2 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<2 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<2 == 1<<3 || 1<<2 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
	ZvalPtrDtorNogc(free_op_data)
exit_assign_obj:

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)

assign_object:
	if 1<<0 == 1<<0 && object.GetValue().GetObj().GetCe() == (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
		var cache_slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
		var zobj *ZendObject = object.GetValue().GetObj()
		var property_val *Zval
		if intptr_t(prop_offset) > 0 {
			property_val = (*Zval)((*byte)(zobj + prop_offset))
			if property_val.GetType() != 0 {
				var prop_info *ZendPropertyInfo = (*ZendPropertyInfo)((cache_slot + 2)[0])
				if prop_info != nil {
					value = ZendAssignToTypedProp(prop_info, property_val, value, execute_data)
					goto free_and_exit_assign_obj
				} else {
				fast_assign_obj:
					value = ZendAssignToVariable(property_val, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
					if opline.GetResultType() != 0 {
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = value
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						if (_t & 0xff00) != 0 {
							ZendGcAddref(&_gc.gc)
						}
					}
					goto exit_assign_obj
				}
			}
		} else {
			if zobj.GetProperties() != nil {
				if ZendGcRefcount(&(zobj.GetProperties()).gc) > 1 {
					if (ZvalGcFlags(zobj.GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
						ZendGcDelref(&(zobj.GetProperties()).gc)
					}
					zobj.SetProperties(ZendArrayDup(zobj.GetProperties()))
				}
				property_val = ZendHashFindEx(zobj.GetProperties(), property.GetValue().GetStr(), 1)
				if property_val != nil {
					goto fast_assign_obj
				}
			}
			if zobj.GetCe().GetSet() == nil {
				if zobj.GetProperties() == nil {
					RebuildObjectProperties(zobj)
				}
				if 1<<3 == 1<<0 {
					if (value.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(value)
					}
				} else if 1<<3 != 1<<1 {
					if value.GetType() == 10 {
						if 1<<3 == 1<<2 {
							var ref *ZendReference = value.GetValue().GetRef()
							if ZendGcDelref(&ref.gc) == 0 {
								var _z1 *Zval = &tmp
								var _z2 *Zval = &(*value).value.GetRef().GetVal()
								var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
								var _t uint32 = _z2.GetTypeInfo()
								_z1.GetValue().SetCounted(_gc)
								_z1.SetTypeInfo(_t)
								_efree(ref)
								value = &tmp
							} else {
								value = &(*value).value.GetRef().GetVal()
								if value.GetTypeFlags() != 0 {
									ZvalAddrefP(value)
								}
							}
						} else {
							value = &(*value).value.GetRef().GetVal()
							if value.GetTypeFlags() != 0 {
								ZvalAddrefP(value)
							}
						}
					} else if 1<<3 == 1<<3 {
						if value.GetTypeFlags() != 0 {
							ZvalAddrefP(value)
						}
					}
				}
				ZendHashAddNew(zobj.GetProperties(), property.GetValue().GetStr(), value)
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = value
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				goto exit_assign_obj
			}
		}
	}
	if 1<<3 == 1<<3 || 1<<3 == 1<<2 {
		if value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
		}
	}
	value = object.GetValue().GetObj().GetHandlers().GetWriteProperty()(object, property, value, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
free_and_exit_assign_obj:
	if opline.GetResultType() != 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	}
exit_assign_obj:

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<0 == 1<<0 {
		ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
	} else {
		ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<0 == 1<<0 {
		ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
	} else {
		ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ROPE_INIT_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var rope **ZendString
	var var_ *Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**ZendString)((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
	if 1<<0 == 1<<0 {
		var_ = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		rope[0] = var_.GetValue().GetStr()
		if var_.GetTypeFlags() != 0 {
			ZvalAddrefP(var_)
		}
	} else {
		var_ = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		if var_.GetType() == 6 {
			if 1<<0 == 1<<3 {
				rope[0] = ZendStringCopy(var_.GetValue().GetStr())
			} else {
				rope[0] = var_.GetValue().GetStr()
			}
		} else {
			if 1<<0 == 1<<3 && var_.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			rope[0] = ZvalGetStringFunc(var_)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var class_name *Zval
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<0 == 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	} else if 1<<0 == 1<<0 {
		var ce *ZendClassEntry = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
		if ce == nil {
			class_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			ce = ZendFetchClassByName(class_name.GetValue().GetStr(), (class_name + 1).GetValue().GetStr(), opline.GetOp1().GetNum())
			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ce)
	} else {
		class_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	try_class_name:
		if class_name.GetType() == 8 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(class_name.GetValue().GetObj().GetCe())
		} else if class_name.GetType() == 6 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(class_name.GetValue().GetStr(), opline.GetOp1().GetNum()))
		} else if (1<<0&(1<<2|1<<3)) != 0 && class_name.GetType() == 10 {
			class_name = &(*class_name).value.GetRef().GetVal()
			goto try_class_name
		} else {
			if 1<<0 == 1<<3 && class_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	if 1<<0 != 1<<0 {
		function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	}
	if 1<<0 != 1<<0 && function_name.GetType() != 6 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
				function_name = &(*function_name).value.GetRef().GetVal()
				if function_name.GetType() == 6 {
					break
				}
			} else if 1<<0 == 1<<3 && function_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}

	obj = object.GetValue().GetObj()
	called_scope = obj.GetCe()
	if 1<<0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == called_scope {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		var orig_obj *ZendObject = obj
		if 1<<0 == 1<<0 {
			function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetValue().GetStr(), g.CondF1(1<<0 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if (0&(1<<2|1<<1)) != 0 && obj != orig_obj {

			/* Reset "object" to trigger reference counting */

			object = nil

			/* Reset "object" to trigger reference counting */

		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if 1<<0 != 1<<0 {

	}
	call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
	if (fbc.GetFnFlags() & 1 << 4) != 0 {
		if (0&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if (0 & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if 0 == 1<<3 {
			ZendGcAddref(&obj.gc)
		} else if free_op1 != object {
			ZendGcAddref(&obj.gc)
		}

		/* CV may be changed indirectly (e.g. when it's a reference) */

		call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8) | 1<<21

		/* CV may be changed indirectly (e.g. when it's a reference) */

	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), obj)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 0 == 1<<0 {

		/* no function found. try a static method in class */

		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				return 0
			}
			if 1<<0 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			return 0
		}
	}
	if 0 == 1<<0 && 1<<0 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 0 != 1<<0 && 1<<0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else if 1<<0 != 0 {
		function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		if 1<<0 != 1<<0 {
			if function_name.GetType() != 6 {
				for {
					if (1<<0&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
						function_name = &(*function_name).value.GetRef().GetVal()
						if function_name.GetType() == 6 {
							break
						}
					} else if 1<<0 == 1<<3 && function_name.GetType() == 0 {
						_zvalUndefinedOp2(execute_data)
						if EG.GetException() != nil {
							return 0
						}
					}
					ZendThrowError(nil, "Function name must be a string")
					return 0
					break
				}
			}
		}
		if ce.GetGetStaticMethod() != nil {
			fbc = ce.GetGetStaticMethod()(ce, function_name.GetValue().GetStr())
		} else {
			fbc = ZendStdGetStaticMethod(ce, function_name.GetValue().GetStr(), g.CondF1(1<<0 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		}
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(ce, function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = ce
			slot[1] = fbc
		}
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
		if 1<<0 != 1<<0 {

		}
	} else {
		if ce.GetConstructor() == nil {
			ZendThrowError(nil, "Cannot call constructor")
			return 0
		}
		if execute_data.GetThis().GetType() == 8 && execute_data.GetThis().GetValue().GetObj().GetCe() != ce.GetConstructor().GetScope() && (ce.GetConstructor().GetFnFlags()&1<<2) != 0 {
			ZendThrowError(nil, "Cannot call private %s::__construct()", ce.GetName().GetVal())
			return 0
		}
		fbc = ce.GetConstructor()
		if fbc.GetType() == 2 && !(g.CondF((uintptr_t(&fbc.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&fbc.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&fbc.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&fbc.op_array)
		}
	}
	if (fbc.GetFnFlags() & 1 << 4) == 0 {
		if execute_data.GetThis().GetType() == 8 && InstanceofFunction(execute_data.GetThis().GetValue().GetObj().GetCe(), ce) != 0 {
			ce = (*ZendClassEntry)(execute_data.GetThis().GetValue().GetObj())
			call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
		} else {
			ZendNonStaticMethodCall(fbc)
			if EG.GetException() != nil {
				return 0
			}
			goto check_parent_and_self
		}
	} else {
	check_parent_and_self:

		/* previous opcode is ZEND_FETCH_CLASS */

		if (opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1 {
			if execute_data.GetThis().GetType() == 8 {
				ce = execute_data.GetThis().GetValue().GetObj().GetCe()
			} else {
				ce = execute_data.GetThis().GetValue().GetCe()
			}
		}
		call_info = 0<<16 | 0<<17
	}
	call = ZendVmStackPushCallFrame(call_info, fbc, opline.GetExtendedValue(), ce)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_CONSTANT_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var c *ZendConstant
	c = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
	if c != nil && (uintptr_t(c)&1<<0) == 0 {
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = &c.value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
				ZendGcAddref(&_gc.gc)
			} else {
				ZvalCopyCtorFunc(_z1)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	ZendQuickGetConstant((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant)+1, opline.GetOp1().GetNum(), opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_CLASS_CONSTANT_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var c *ZendClassConstant
	var value *Zval
	var zv *Zval
	var opline *ZendOp = execute_data.GetOpline()
	for {
		if 0 == 1<<0 {
			if (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0] {
				value = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0]
				break
			} else if (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] {
				ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
			} else {
				ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
				if ce == nil {
					r.Assert(EG.GetException() != nil)
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					return 0
				}
			}
		} else {
			ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
			if (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] == ce {
				value = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() + g.SizeOf("void *"))))[0]
				break
			}
		}
		zv = ZendHashFindEx(&ce.constants_table, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), 1)
		if zv != nil {
			c = zv.GetValue().GetPtr()
			scope = execute_data.GetFunc().GetOpArray().GetScope()
			if ZendVerifyConstAccess(c, scope) == 0 {
				ZendThrowError(nil, "Cannot access %s const %s::%s", ZendVisibilityString(c.GetValue().GetAccessFlags()), ce.GetName().GetVal(), (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal())
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
			value = &c.value
			if value.GetType() == 11 {
				ZvalUpdateConstantEx(value, c.GetCe())
				if EG.GetException() != nil {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					return 0
				}
			}
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
			slot[0] = ce
			slot[1] = value
		} else {
			ZendThrowError(nil, "Undefined class constant '%s'", (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr().GetVal())
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
		break
	}
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		if (ZvalGcFlags(_gc.GetGc().GetTypeInfo()) & 1 << 7) == 0 {
			ZendGcAddref(&_gc.gc)
		} else {
			ZvalCopyCtorFunc(_z1)
		}
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = array
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

	container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 0 == 1<<0 {
		if (0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() != 8 {
				result = opline.GetExtendedValue() & 1 << 0
				goto isset_object_finish
			}
		} else {
			result = opline.GetExtendedValue() & 1 << 0
			goto isset_object_finish
		}
	}
	result = opline.GetExtendedValue()&1<<0 ^ container.GetValue().GetObj().GetHandlers().GetHasProperty()(container, offset, opline.GetExtendedValue()&1<<0, g.CondF1(1<<0 == 1<<0, func() *any {
		return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))
	}, nil))
isset_object_finish:
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result != 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result == 0 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else {
			break
		}
		return 0
		break
	}
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_UNUSED_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		return zend_yield_in_closed_generator_helper_SPEC(execute_data)
	}

	/* Destroy the previously yielded value */

	ZvalPtrDtor(&generator.value)

	/* Destroy the previously yielded key */

	ZvalPtrDtor(&generator.key)

	/* Set the new yielded value */

	/* If no value was specified yield null */

	&generator.value.u1.type_info = 1

	/* If no value was specified yield null */

	/* Set the new yielded key */

	if 1<<0 != 0 {
		var key *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)

		/* Consts, temporary variables and references need copying */

		if 1<<0 == 1<<0 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (generator.GetKey().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetKey()))
			}
		} else if 1<<0 == 1<<1 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else if (1<<0&(1<<2|1<<3)) != 0 && key.GetType() == 10 {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = &(*key).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if (_t & 0xff00) != 0 {
				ZendGcAddref(&_gc.gc)
			}
		} else {
			var _z1 *Zval = &generator.key
			var _z2 *Zval = key
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<0 == 1<<3 {
				if (key.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(key)
				}
			}
		}
		if generator.GetKey().GetType() == 4 && generator.GetKey().GetValue().GetLval() > generator.GetLargestUsedIntegerKey() {
			generator.SetLargestUsedIntegerKey(generator.GetKey().GetValue().GetLval())
		}
	} else {

		/* If no key was specified we use auto-increment keys */

		generator.GetLargestUsedIntegerKey()++
		var __z *Zval = &generator.key
		__z.GetValue().SetLval(generator.GetLargestUsedIntegerKey())
		__z.SetTypeInfo(4)
	}
	if opline.GetResultType() != 0 {

		/* If the return value of yield is used set the send
		 * target and initialize it to NULL */

		generator.SetSendTarget((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
		generator.GetSendTarget().SetTypeInfo(1)
	} else {
		generator.SetSendTarget(nil)
	}

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)

	assign_op_object:

		/* here we are sure we are dealing with an object */

		if (1<<1 | 1<<2) == 1<<0 {
			cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline + 1).GetExtendedValue()))
		} else {
			cache_slot = nil
		}
		if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
			if zptr.GetType() == 15 {
				if opline.GetResultType() != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
				}
			} else {
				var orig_zptr *Zval = zptr
				var ref *ZendReference
				for {
					if zptr.GetType() == 10 {
						ref = zptr.GetValue().GetRef()
						zptr = &(*zptr).value.GetRef().GetVal()
						if ref.GetSources().GetPtr() != nil {
							ZendBinaryAssignOpTypedRef(ref, value, opline, execute_data)
							break
						}
					}
					if (1<<1 | 1<<2) == 1<<0 {
						prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
					} else {
						prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), orig_zptr)
					}
					if prop_info != nil {

						/* special case for typed properties */

						ZendBinaryAssignOpTypedProp(prop_info, zptr, value, opline, execute_data)

						/* special case for typed properties */

					} else {
						ZendBinaryOp(zptr, zptr, value, opline)
					}
					break
				}
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = zptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
			}
		} else {
			ZendAssignOpOverloadedProperty(object, property, cache_slot, value, opline, execute_data)
		}
		break
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	ZvalPtrDtorNogc(free_op2)

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMP|VAR|CV, UNUSED|CONST|TMPVAR) */

func ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)

pre_incdec_object:

	/* here we are sure we are dealing with an object */

	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
	} else {
		cache_slot = nil
	}
	if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
		if zptr.GetType() == 15 {
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		} else {
			if (1<<1 | 1<<2) == 1<<0 {
				prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
			} else {
				prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
			}
			ZendPreIncdecPropertyZval(zptr, prop_info, opline, execute_data)
		}
	} else {
		ZendPreIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)

post_incdec_object:

	/* here we are sure we are dealing with an object */

	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
	} else {
		cache_slot = nil
	}
	if g.Assign(&zptr, object.GetValue().GetObj().GetHandlers().GetGetPropertyPtrPtr()(object, property, 2, cache_slot)) != nil {
		if zptr.GetType() == 15 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		} else {
			if (1<<1 | 1<<2) == 1<<0 {
				prop_info = (*ZendPropertyInfo)((cache_slot + 2)[0])
			} else {
				prop_info = ZendObjectFetchPropertyTypeInfo(object.GetValue().GetObj(), zptr)
			}
			ZendPostIncdecPropertyZval(zptr, prop_info, opline, execute_data)
		}
	} else {
		ZendPostIncdecOverloadedProperty(object, property, cache_slot, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 0 == 1<<0 {
		for {
			if (0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 0 == 1<<3 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if (1<<1|1<<2) == 1<<3 && offset.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			ZendWrongPropertyRead(offset)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_r_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^1)))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetTypeInfo() != 0 {
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
					fetch_obj_r_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						r.Assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_r_copy
							} else {
								goto fetch_obj_r_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	} else if (1<<1|1<<2) == 1<<3 && offset.GetTypeInfo() == 0 {
		_zvalUndefinedOp2(execute_data)
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 0, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_r_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_r_finish:
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 0 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	if 0 == 1<<2 {
		var __container_to_free *Zval = free_op1
		if __container_to_free != nil && __container_to_free.GetTypeFlags() != 0 {
			var __ref *ZendRefcounted = __container_to_free.GetValue().GetCounted()
			if ZendGcDelref(&__ref.gc) == 0 {
				var __zv *Zval = result
				if __zv.GetType() == 13 {
					var _z1 *Zval = __zv
					var _z2 *Zval = __zv.GetValue().GetZv()
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				RcDtorFunc(__ref)
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 0 == 1<<0 {
		for {
			if (0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			goto fetch_obj_is_finish
			break
		}
	}

	/* here we are sure we are dealing with an object */

	var zobj *ZendObject = container.GetValue().GetObj()
	var retval *Zval
	if (1<<1 | 1<<2) == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetType() != 0 {
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
					fetch_obj_is_fast_copy:
						var _z3 *Zval = retval
						if (_z3.GetTypeInfo() & 0xff00) != 0 {
							if (_z3.GetTypeInfo() & 0xff) == 10 {
								_z3 = &(*_z3).value.GetRef().GetVal()
								if (_z3.GetTypeInfo() & 0xff00) != 0 {
									ZvalAddrefP(_z3)
								}
							} else {
								ZvalAddrefP(_z3)
							}
						}
						var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
						var _z2 *Zval = _z3
						var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
						var _t uint32 = _z2.GetTypeInfo()
						_z1.GetValue().SetCounted(_gc)
						_z1.SetTypeInfo(_t)
						r.Assert(EG.GetException() == nil)
						execute_data.SetOpline(opline + 1)
						return 0
					}
				}
			} else if zobj.GetProperties() != nil {
				if prop_offset != uintptr_t(intptr_t)(-1) {
					var idx uintPtr = uintptr_t(-(intptr_t(prop_offset)) - 2)
					if idx < zobj.GetProperties().GetNNumUsed()*g.SizeOf("Bucket") {
						var p *Bucket = (*Bucket)((*byte)(zobj.GetProperties().GetArData() + idx))
						if p.GetVal().GetType() != 0 && (p.GetKey() == offset.GetValue().GetStr() || p.GetH() == offset.GetValue().GetStr().GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), offset.GetValue().GetStr()) != 0) {
							retval = &p.val
							if (0 & (1<<1 | 1<<2)) != 0 {
								goto fetch_obj_is_copy
							} else {
								goto fetch_obj_is_fast_copy
							}
						}
					}
					(cache_slot + 1)[0] = any(uintptr_t(intptr_t)(-1))
				}
				retval = ZendHashFindEx(zobj.GetProperties(), offset.GetValue().GetStr(), 1)
				if retval != nil {
					var idx uintPtr = (*byte)(retval - (*byte)(zobj.GetProperties().GetArData()))
					(cache_slot + 1)[0] = any(uintptr_t(-(intptr_t(idx) + 2)))
					if (0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_is_copy
					} else {
						goto fetch_obj_is_fast_copy
					}
				}
			}
		}
	}
	retval = zobj.GetHandlers().GetReadProperty()(container, offset, 3, cache_slot, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())))
	if retval != (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())) {
	fetch_obj_is_copy:
		var _z3 *Zval = retval
		if (_z3.GetTypeInfo() & 0xff00) != 0 {
			if (_z3.GetTypeInfo() & 0xff) == 10 {
				_z3 = &(*_z3).value.GetRef().GetVal()
				if (_z3.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(_z3)
				}
			} else {
				ZvalAddrefP(_z3)
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if retval.GetType() == 10 {
		ZendUnwrapReference(retval)
	}
fetch_obj_is_finish:
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
