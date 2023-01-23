// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
)

func ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<0 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER(execute_data)
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
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if container.GetType() == 7 {
		var ht *HashTable
		var value *Zval
		var str *ZendString
	isset_dim_obj_array:
		ht = container.GetValue().GetArr()
	isset_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if (1<<1 | 1<<2) != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index_prop
				}
			}
			value = ZendHashFindExInd(ht, str, (1<<1|1<<2) == 1<<0)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index_prop:
			value = ZendHashIndexFind(ht, hval)
		} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto isset_again
		} else {
			value = ZendFindArrayDimSlow(ht, offset, execute_data)
			if EG.GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & 1 << 0) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1)
			if (1 << 0 & (1<<0 | 1<<3)) != 0 {

				/* avoid exception check */

				ZvalPtrDtorNogc(free_op2)
				for {

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
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto isset_dim_obj_array
		}
	}
	if (1<<1|1<<2) == 1<<0 && offset.GetU2Extra() == 1 {
		offset++
	}
	if (opline.GetExtendedValue() & 1 << 0) == 0 {
		result = ZendIssetDimSlow(container, offset, execute_data)
	} else {
		result = ZendIsemptyDimSlow(container, offset, execute_data)
	}
isset_dim_obj_exit:
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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var result int
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
	result = opline.GetExtendedValue()&1<<0 ^ container.GetValue().GetObj().GetHandlers().GetHasProperty()(container, offset, opline.GetExtendedValue()&1<<0, g.CondF1((1<<1|1<<2) == 1<<0, func() *any {
		return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^(1 << 0))))
	}, nil))
isset_object_finish:
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	subject = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if subject.GetType() == 7 {
	array_key_exists_array:
		ht = subject.GetValue().GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, execute_data)
	} else {
		if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && subject.GetType() == 10 {
			subject = &(*subject).value.GetRef().GetVal()
			if subject.GetType() == 7 {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result == 3 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result != 3 {
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
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(result)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|UNUSED|CV, ANY) */

func ZEND_YIELD_SPEC_CONST_TMP_HANDLER(execute_data *ZendExecuteData) int {
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

	if 1<<0 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 0 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = nil

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<0 == 1<<2 {
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

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

			/* Consts, temporary variables and references need copying */

			if 1<<0 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<0 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<3 {
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
func ZEND_YIELD_SPEC_CONST_VAR_HANDLER(execute_data *ZendExecuteData) int {
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

	if 1<<0 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 0 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = nil

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<0 == 1<<2 {
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

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

			/* Consts, temporary variables and references need copying */

			if 1<<0 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<0 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<3 {
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
func zend_fetch_var_address_helper_SPEC_CONST_UNUSED(type_ int, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *Zval
	var retval *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else if varname.GetType() == 6 {
		name = varname.GetValue().GetStr()
		tmp_name = nil
	} else {
		if 1<<0 == 1<<3 && varname.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	retval = ZendHashFindEx(target_symbol_table, name, 1<<0 == 1<<0)
	if retval == nil {
		if ZendStringEquals(name, ZendKnownStrings[ZEND_STR_THIS]) != 0 {
		fetch_this:
			ZendFetchThisVar(type_, opline, execute_data)
			if 1<<0 != 1<<0 {
				ZendTmpStringRelease(tmp_name)
			}
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
		if type_ == 1 {
			retval = ZendHashAddNew(target_symbol_table, name, &EG.uninitialized_zval)
		} else if type_ == 3 {
			retval = &EG.uninitialized_zval
		} else {
			ZendError(1<<3, "Undefined variable: %s", name.GetVal())
			if type_ == 2 {
				retval = ZendHashUpdate(target_symbol_table, name, &EG.uninitialized_zval)
			} else {
				retval = &EG.uninitialized_zval
			}
		}
	} else if retval.GetType() == 13 {
		retval = retval.GetValue().GetZv()
		if retval.GetType() == 0 {
			if ZendStringEquals(name, ZendKnownStrings[ZEND_STR_THIS]) != 0 {
				goto fetch_this
			}
			if type_ == 1 {
				retval.SetTypeInfo(1)
			} else if type_ == 3 {
				retval = &EG.uninitialized_zval
			} else {
				ZendError(1<<3, "Undefined variable: %s", name.GetVal())
				if type_ == 2 {
					retval.SetTypeInfo(1)
				} else {
					retval = &EG.uninitialized_zval
				}
			}
		}
	}
	if (opline.GetExtendedValue() & 1 << 3) == 0 {

	}
	if 1<<0 != 1<<0 {
		ZendTmpStringRelease(tmp_name)
	}
	r.Assert(retval != nil)
	if type_ == 0 || type_ == 3 {
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
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetZv(retval)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(13)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_R_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(0, execute_data)
}
func ZEND_FETCH_W_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(1, execute_data)
}
func ZEND_FETCH_RW_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(2, execute_data)
}
func ZEND_FETCH_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var fetch_type int = g.Cond((execute_data.GetCall().GetThis().GetTypeInfo()&1<<31) != 0, 1, 0)
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(fetch_type, execute_data)
}
func ZEND_FETCH_UNSET_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(5, execute_data)
}
func ZEND_FETCH_IS_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CONST_UNUSED(3, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		return ZEND_NULL_HANDLER(execute_data)
	}
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<0 == 1<<0 {

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
	} else if 1<<0 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<0 == 1<<0 && 0 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<0 != 1<<0 && 0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
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

		if 1<<0 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
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
func ZEND_VERIFY_RETURN_TYPE_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<0 == 0 {
		ZendVerifyMissingReturnType(execute_data.GetFunc(), (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum())))
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_NEW_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var result *Zval
	var constructor *ZendFunction
	var ce *ZendClassEntry
	var call *ZendExecuteData
	if 1<<0 == 1<<0 {
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
	} else if 1<<0 == 0 {
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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<0 == 1<<2 || 1<<0 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = nil
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
	} else {
		expr_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		if 1<<0 == 1<<1 {

		} else if 1<<0 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<0 == 1<<3 {
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
func ZEND_INIT_ARRAY_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<0 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER(execute_data)
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
func ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else if varname.GetType() == 6 {
		name = varname.GetValue().GetStr()
		tmp_name = nil
	} else {
		if 1<<0 == 1<<3 && varname.GetType() == 0 {
			varname = _zvalUndefinedOp1(execute_data)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	ZendHashDelInd(target_symbol_table, name)
	if 1<<0 != 1<<0 {
		ZendTmpStringRelease(tmp_name)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int
	var varname *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else {
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	value = ZendHashFindEx(target_symbol_table, name, 1<<0 == 1<<0)
	if 1<<0 != 1<<0 {
		ZendTmpStringRelease(tmp_name)
	}
	if value == nil {
		result = opline.GetExtendedValue() & 1 << 0
	} else {
		if value.GetType() == 13 {
			value = value.GetValue().GetZv()
		}
		if (opline.GetExtendedValue() & 1 << 0) == 0 {
			if value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
			}
			result = value.GetType() > 1
		} else {
			result = !(IZendIsTrue(value))
		}
	}
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

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var func_ *ZendFunction
	var zfunc *Zval
	var object *Zval
	var called_scope *ZendClassEntry
	func_ = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
	if func_ == nil {
		zfunc = ZendHashFindEx(EG.GetFunctionTable(), (*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), 1)
		r.Assert(zfunc != nil)
		func_ = zfunc.GetValue().GetFunc()
		r.Assert(func_.GetType() == 2)
		(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = func_
	}
	if execute_data.GetThis().GetType() == 8 {
		called_scope = execute_data.GetThis().GetValue().GetObj().GetCe()
		if (func_.GetFnFlags()&1<<4) != 0 || (execute_data.GetFunc().GetFnFlags()&1<<4) != 0 {
			object = nil
		} else {
			object = &(execute_data.GetThis())
		}
	} else {
		called_scope = execute_data.GetThis().GetValue().GetCe()
		object = nil
	}
	ZendCreateClosure((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), func_, execute_data.GetFunc().GetOpArray().GetScope(), called_scope, object)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_YIELD_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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

	if 1<<0 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 0 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = nil

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<0 == 1<<2 {
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

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

			/* Consts, temporary variables and references need copying */

			if 1<<0 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<0 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<3 {
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
func ZEND_COUNT_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var count ZendLong
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	for true {
		if op1.GetType() == 7 {
			count = ZendArrayCount(op1.GetValue().GetArr())
			break
		} else if op1.GetType() == 8 {

			/* first, we check if the handler is defined */

			if op1.GetValue().GetObj().GetHandlers().GetCountElements() != nil {
				if SUCCESS == op1.GetValue().GetObj().GetHandlers().GetCountElements()(op1, &count) {
					break
				}
				if EG.GetException() != nil {
					count = 0
					break
				}
			}

			/* if not and the object implements Countable we call its count() method */

			if InstanceofFunction(op1.GetValue().GetObj().GetCe(), ZendCeCountable) != 0 {
				var retval Zval
				ZendCallMethod(op1, nil, nil, "count", g.SizeOf("\"count\"")-1, &retval, 0, nil, nil)
				count = ZvalGetLong(&retval)
				ZvalPtrDtor(&retval)
				break
			}

			/* If There's no handler and it doesn't implement Countable then add a warning */

			count = 1

			/* If There's no handler and it doesn't implement Countable then add a warning */

		} else if (1<<0&(1<<2|1<<3)) != 0 && op1.GetType() == 10 {
			op1 = &(*op1).value.GetRef().GetVal()
			continue
		} else if op1.GetType() <= 1 {
			if 1<<0 == 1<<3 && op1.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			count = 0
		} else {
			count = 1
		}
		ZendError(1<<1, "%s(): Parameter must be an array or an object that implements Countable", g.Cond(opline.GetExtendedValue() != 0, "sizeof", "count"))
		break
	}
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(count)
	__z.SetTypeInfo(4)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<0 == 0 {
		if execute_data.GetFunc().GetScope() == nil {
			ZendError(1<<1, "get_class() called without object from outside a class")
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = execute_data.GetFunc().GetScope().GetName()
			__z.GetValue().SetStr(__s)
			if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
				__z.SetTypeInfo(6)
			} else {
				ZendGcAddref(&__s.gc)
				__z.SetTypeInfo(6 | 1<<0<<8)
			}
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else {
		var op1 *Zval
		op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		for true {
			if op1.GetType() == 8 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1.GetValue().GetObj().GetCe().GetName()
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else if (1<<0&(1<<2|1<<3)) != 0 && op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				continue
			} else {
				if 1<<0 == 1<<3 && op1.GetType() == 0 {
					_zvalUndefinedOp1(execute_data)
				}
				ZendError(1<<1, "get_class() expects parameter 1 to be object, %s given", ZendGetTypeByConst(op1.GetType()))
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
			}
			break
		}
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var type_ *ZendString
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FUNC_GET_ARGS_SPEC_CONST_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var ht *ZendArray
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = execute_data.GetThis().GetNumArgs()
	if 1<<0 == 1<<0 {
		skip = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant).GetValue().GetLval()
		if arg_count < skip {
			result_size = 0
		} else {
			result_size = arg_count - skip
		}
	} else {
		skip = 0
		result_size = arg_count
	}
	if result_size != 0 {
		var first_extra_arg uint32 = execute_data.GetFunc().GetOpArray().GetNumArgs()
		ht = _zendNewArray(result_size)
		var __arr *ZendArray = ht
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		ZendHashRealInitPacked(ht)
		var __fill_ht *HashTable = ht
		var __fill_bkt *Bucket = __fill_ht.GetArData() + __fill_ht.GetNNumUsed()
		var __fill_idx uint32 = __fill_ht.GetNNumUsed()
		r.Assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
		var p *Zval
		var q *Zval
		var i uint32 = skip
		p = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(i))
		if arg_count > first_extra_arg {
			for i < first_extra_arg {
				q = p
				if q.GetTypeInfo() != 0 {
					if q.GetType() == 10 {
						q = &(*q).value.GetRef().GetVal()
					}
					if (q.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(q)
					}
					var _z1 *Zval = &__fill_bkt.val
					var _z2 *Zval = q
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
				} else {
					&__fill_bkt.val.u1.type_info = 1
				}
				__fill_bkt.SetH(__fill_idx)
				__fill_bkt.SetKey(nil)
				__fill_bkt++
				__fill_idx++
				p++
				i++
			}
			if skip < first_extra_arg {
				skip = 0
			} else {
				skip -= first_extra_arg
			}
			p = (*Zval)(execute_data) + (int(((g.SizeOf("zend_execute_data")+8 - 1 & ^(8-1))+(g.SizeOf("zval")+8 - 1 & ^(8-1))-1)/(g.SizeOf("zval")+8 - 1 & ^(8-1))) + int(execute_data.GetFunc().GetOpArray().GetLastVar()+execute_data.GetFunc().GetOpArray().GetT()+skip))
		}
		for i < arg_count {
			q = p
			if q.GetTypeInfo() != 0 {
				if q.GetType() == 10 {
					q = &(*q).value.GetRef().GetVal()
				}
				if (q.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(q)
				}
				var _z1 *Zval = &__fill_bkt.val
				var _z2 *Zval = q
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				&__fill_bkt.val.u1.type_info = 1
			}
			__fill_bkt.SetH(__fill_idx)
			__fill_bkt.SetKey(nil)
			__fill_bkt++
			__fill_idx++
			p++
			i++
		}
		__fill_ht.SetNNumUsed(__fill_idx)
		__fill_ht.SetNNumOfElements(__fill_idx)
		__fill_ht.SetNNextFreeElement(__fill_idx)
		__fill_ht.SetNInternalPointer(0)
		ht.SetNNumOfElements(result_size)
	} else {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
		__z.SetTypeInfo(7)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DIV_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<0 == 1<<0 || op1.GetType() == 6) && (1<<3 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<0 != 1<<0 && op1_str.GetLen() == 0 {
			if 1<<3 == 1<<0 || 1<<3 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<3 != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<0 != 1<<0 && 1<<0 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if 1<<0 == 1<<3 && op1.GetType() == 0 {
			op1 = _zvalUndefinedOp1(execute_data)
		}
		if 1<<3 == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_SPACESHIP_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 != 1<<0 {
		if container.GetType() == 7 {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, 1<<3, 0, execute_data)
			var _z3 *Zval = value
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
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if 1<<3 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		}
	} else {
		zend_fetch_dimension_address_read_R(container, dim, 1<<3, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_IS_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_read_IS(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		if 1<<3 == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 1<<0 == 1<<3 && container.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			if 1<<3 == 1<<3 && offset.GetType() == 0 {
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
	if 1<<3 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^1)))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetTypeInfo() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
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
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
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
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
						goto fetch_obj_r_copy
					} else {
						goto fetch_obj_r_fast_copy
					}
				}
			}
		}
	} else if 1<<3 == 1<<3 && offset.GetTypeInfo() == 0 {
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
func ZEND_FETCH_OBJ_IS_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		for {
			if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
	if 1<<3 == 1<<0 {
		cache_slot = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))
		if zobj.GetCe() == cache_slot[0] {
			var prop_offset uintPtr = uintptr_t(cache_slot + 1)[0]
			if intptr_t(prop_offset) > 0 {
				retval = (*Zval)((*byte)(zobj + prop_offset))
				if retval.GetType() != 0 {
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
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
							if (1 << 0 & (1<<1 | 1<<2)) != 0 {
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
					if (1 << 0 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_NULL_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_LIST_R_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	zend_fetch_dimension_address_LIST_r(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<0 == 1<<0 || op1.GetType() == 6) && (1<<3 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<0 != 1<<0 && op1_str.GetLen() == 0 {
			if 1<<3 == 1<<0 || 1<<3 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<3 != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<0 != 1<<0 && 1<<0 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			if len_ > SIZE_MAX-(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1+8 - 1 & ^(8-1))-op2_str.GetLen() {
				ZendErrorNoreturn(1<<0, "Integer overflow in memory allocation")
			}
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<0 == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if 1<<0 == 1<<3 && op1.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	if 1<<3 == 1<<0 {
		op2_str = op2.GetValue().GetStr()
	} else if op2.GetType() == 6 {
		op2_str = ZendStringCopy(op2.GetValue().GetStr())
	} else {
		if 1<<3 == 1<<3 && op2.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		if 1<<0 != 1<<0 {
			if op1_str.GetLen() == 0 {
				if 1<<3 == 1<<0 {
					if op2.GetTypeFlags() != 0 {
						ZendGcAddref(&op2_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		if 1<<3 != 1<<0 {
			if op2_str.GetLen() == 0 {
				if 1<<0 == 1<<0 {
					if op1.GetTypeFlags() != 0 {
						ZendGcAddref(&op1_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if 1<<3 != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	if 1<<3 != 1<<0 {
		function_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	}
	if 1<<3 != 1<<0 && function_name.GetType() != 6 {
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
			ZendThrowError(nil, "Method name must be a string")
			return 0
			break
		}
	}
	if 1<<0 != 0 {
		for {
			if 1<<0 == 1<<0 || object.GetType() != 8 {
				if (1<<0&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if 1<<0 == 1<<3 && object.GetType() == 0 {
					object = _zvalUndefinedOp1(execute_data)
					if EG.GetException() != nil {
						if 1<<3 != 1<<0 {

						}
						return 0
					}
				}
				if 1<<3 == 1<<0 {
					function_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
				}
				ZendInvalidMethodCall(object, function_name)
				return 0
			}
			break
		}
	}
	obj = object.GetValue().GetObj()
	called_scope = obj.GetCe()
	if 1<<3 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == called_scope {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		var orig_obj *ZendObject = obj
		if 1<<3 == 1<<0 {
			function_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetValue().GetStr(), g.CondF1(1<<3 == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetValue().GetStr())
			}
			return 0
		}
		if 1<<3 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if (1<<0&(1<<2|1<<1)) != 0 && obj != orig_obj {

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
	if 1<<3 != 1<<0 {

	}
	call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
	if (fbc.GetFnFlags() & 1 << 4) != 0 {
		if (1<<0&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if (1 << 0 & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if 1<<0 == 1<<3 {
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
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var ce *ZendClassEntry
	var call_info uint32
	var fbc *ZendFunction
	var call *ZendExecuteData
	if 1<<0 == 1<<0 {

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
	} else if 1<<0 == 0 {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			r.Assert(EG.GetException() != nil)
			return 0
		}
	} else {
		ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())).GetValue().GetCe()
	}
	if 1<<0 == 1<<0 && 1<<3 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 1<<0 != 1<<0 && 1<<3 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
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

		if 1<<0 == 0 && ((opline.GetOp1().GetNum()&0xf) == 2 || (opline.GetOp1().GetNum()&0xf) == 1) {
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
func ZEND_INIT_USER_CALL_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var fcc ZendFcallInfoCache
	var error *byte = nil
	var func_ *ZendFunction
	var object_or_called_scope any
	var call *ZendExecuteData
	var call_info uint32 = 0<<16 | 0<<17 | 1<<25
	function_name = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	if ZendIsCallableEx(function_name, nil, 0, nil, &fcc, &error) != 0 {
		func_ = fcc.GetFunctionHandler()
		if error != nil {
			_efree(error)

			/* This is the only soft error is_callable() can generate */

			ZendNonStaticMethodCall(func_)
			if EG.GetException() != nil {
				return 0
			}
		}
		object_or_called_scope = fcc.GetCalledScope()
		if (func_.GetFnFlags() & 1 << 20) != 0 {

			/* Delay closure destruction until its invocation */

			ZendGcAddref(&((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object")))).gc)
			call_info |= 1 << 22
			if (func_.GetFnFlags() & 1 << 21) != 0 {
				call_info |= 1 << 23
			}
			if fcc.GetObject() != nil {
				object_or_called_scope = fcc.GetObject()
				call_info |= 8 | 1<<0<<8 | 1<<1<<8
			}
		} else if fcc.GetObject() != nil {
			ZendGcAddref(&(fcc.GetObject()).gc)
			object_or_called_scope = fcc.GetObject()
			call_info |= 1<<21 | (8 | 1<<0<<8 | 1<<1<<8)
		}
		if (1<<3&(1<<1|1<<2)) != 0 && EG.GetException() != nil {
			if (call_info & 1 << 22) != 0 {
				ZendObjectRelease((*ZendObject)((*byte)(func_ - g.SizeOf("zend_object"))))
			} else if (call_info & 1 << 21) != 0 {
				ZendObjectRelease(fcc.GetObject())
			}
			return 0
		}
		if func_.GetType() == 2 && !(g.CondF((uintptr_t(&func_.op_array).run_time_cache__ptr&1) != 0, func() any {
			return *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(&func_.op_array).run_time_cache__ptr - 1)))
		}, func() any { return any(*(&func_.op_array.run_time_cache__ptr)) })) {
			InitFuncRunTimeCache(&func_.op_array)
		}
	} else {
		ZendInternalTypeError((execute_data.GetFunc().GetFnFlags()&1<<31) != 0, "%s() expects parameter 1 to be a valid callback, %s", (*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr().GetVal(), error)
		_efree(error)
		if EG.GetException() != nil {
			return 0
		}
		func_ = (*ZendFunction)(&ZendPassFunction)
		object_or_called_scope = nil
	}
	call = ZendVmStackPushCallFrame(call_info, func_, opline.GetExtendedValue(), object_or_called_scope)
	call.SetPrevExecuteData(execute_data.GetCall())
	execute_data.SetCall(call)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<0 == 1<<2 || 1<<0 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = nil
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
	} else {
		expr_ptr = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
		if 1<<0 == 1<<1 {

		} else if 1<<0 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<0 == 1<<3 {
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
func ZEND_INIT_ARRAY_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<0 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER(execute_data)
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
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	offset = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if container.GetType() == 7 {
		var ht *HashTable
		var value *Zval
		var str *ZendString
	isset_dim_obj_array:
		ht = container.GetValue().GetArr()
	isset_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<3 != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index_prop
				}
			}
			value = ZendHashFindExInd(ht, str, 1<<3 == 1<<0)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index_prop:
			value = ZendHashIndexFind(ht, hval)
		} else if (1<<3&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto isset_again
		} else {
			value = ZendFindArrayDimSlow(ht, offset, execute_data)
			if EG.GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & 1 << 0) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1)
			if (1 << 0 & (1<<0 | 1<<3)) != 0 {

				/* avoid exception check */

				for {

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
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto isset_dim_obj_array
		}
	}
	if 1<<3 == 1<<0 && offset.GetU2Extra() == 1 {
		offset++
	}
	if (opline.GetExtendedValue() & 1 << 0) == 0 {
		result = ZendIssetDimSlow(container, offset, execute_data)
	} else {
		result = ZendIsemptyDimSlow(container, offset, execute_data)
	}
isset_dim_obj_exit:
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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	if 1<<0 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	if 1<<0 == 1<<0 || 1<<0 != 0 && container.GetType() != 8 {
		if (1<<0&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
	result = opline.GetExtendedValue()&1<<0 ^ container.GetValue().GetObj().GetHandlers().GetHasProperty()(container, offset, opline.GetExtendedValue()&1<<0, g.CondF1(1<<3 == 1<<0, func() *any {
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
	subject = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if subject.GetType() == 7 {
	array_key_exists_array:
		ht = subject.GetValue().GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, execute_data)
	} else {
		if (1<<3&(1<<2|1<<3)) != 0 && subject.GetType() == 10 {
			subject = &(*subject).value.GetRef().GetVal()
			if subject.GetType() == 7 {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, execute_data)
	}
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result == 3 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result != 3 {
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
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(result)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|UNUSED|CV, ANY) */

func ZEND_YIELD_SPEC_CONST_CV_HANDLER(execute_data *ZendExecuteData) int {
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

	if 1<<0 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 0 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = nil

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<0 == 1<<2 {
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

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

			}

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

		} else {
			var value *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp1()).constant)

			/* Consts, temporary variables and references need copying */

			if 1<<0 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<0 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<0&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = &(*value).value.GetRef().GetVal()
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
			} else {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<0 == 1<<3 {
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
func ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(value.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetDval(value.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongAddFunction(result, op1, op2)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto add_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		add_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 + d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongSubFunction(result, op1, op2)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto sub_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		sub_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 - d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			var overflow ZendLong
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __lres long = op1.GetValue().GetLval() * op2.GetValue().GetLval()
			var __dres long__double = long__double(op1.GetValue().GetLval() * long__double(op2.GetValue().GetLval()))
			var __delta long__double = long__double(__lres - __dres)
			if g.Assign(&overflow, __dres+__delta != __dres) {
				result.GetValue().SetDval(__dres)
			} else {
				result.GetValue().SetLval(__lres)
			}
			if overflow != 0 {
				result.SetTypeInfo(5)
			} else {
				result.SetTypeInfo(4)
			}
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto mul_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		mul_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 * d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto mul_double
		}
	}
	return zend_mul_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			if op2.GetValue().GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(execute_data)
			} else if op2.GetValue().GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetLval(op1.GetValue().GetLval() % op2.GetValue().GetLval())
				__z.SetTypeInfo(4)
			}
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	return zend_mod_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(zend_long(zend_ulong(*op1).value.lval << op2.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_left_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() >> op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_right_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() | op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_or_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() & op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_and_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2|1<<3) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() ^ op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_xor_helper_SPEC(op1, op2, execute_data)
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_LIST_r(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op *Zval
	var jump_zv *Zval
	var jumptable *HashTable
	op = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	jumptable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	if op.GetType() != 4 {
		if op.GetType() == 10 {
			op = &(*op).value.GetRef().GetVal()
		}
		if op.GetType() != 4 {

			/* Wrong type, fall back to ZEND_CASE chain */

			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	jump_zv = ZendHashIndexFind(jumptable, op.GetValue().GetLval())
	if jump_zv != nil {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(jump_zv.GetValue().GetLval())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else {

		/* default */

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op *Zval
	var jump_zv *Zval
	var jumptable *HashTable
	op = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	jumptable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	if op.GetType() != 6 {
		if (1<<1 | 1<<2 | 1<<3) == 1<<0 {

			/* Wrong type, fall back to ZEND_CASE chain */

			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else {
			if op.GetType() == 10 {
				op = &(*op).value.GetRef().GetVal()
			}
			if op.GetType() != 6 {

				/* Wrong type, fall back to ZEND_CASE chain */

				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	}
	jump_zv = ZendHashFindEx(jumptable, op.GetValue().GetStr(), (1<<1|1<<2|1<<3) == 1<<0)
	if jump_zv != nil {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(jump_zv.GetValue().GetLval())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else {

		/* default */

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() + op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	FastLongAddFunction(result, op1, op2)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() + op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() - op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	FastLongSubFunction(result, op1, op2)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() - op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() * op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var overflow ZendLong
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __lres long = op1.GetValue().GetLval() * op2.GetValue().GetLval()
	var __dres long__double = long__double(op1.GetValue().GetLval() * long__double(op2.GetValue().GetLval()))
	var __delta long__double = long__double(__lres - __dres)
	if g.Assign(&overflow, __dres+__delta != __dres) {
		result.GetValue().SetDval(__dres)
	} else {
		result.GetValue().SetLval(__lres)
	}
	if overflow != 0 {
		result.SetTypeInfo(5)
	} else {
		result.SetTypeInfo(4)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() * op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongAddFunction(result, op1, op2)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto add_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		add_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 + d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto add_double
		}
	}
	return zend_add_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			FastLongSubFunction(result, op1, op2)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto sub_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		sub_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 - d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto sub_double
		}
	}
	return zend_sub_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			var overflow ZendLong
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __lres long = op1.GetValue().GetLval() * op2.GetValue().GetLval()
			var __dres long__double = long__double(op1.GetValue().GetLval() * long__double(op2.GetValue().GetLval()))
			var __delta long__double = long__double(__lres - __dres)
			if g.Assign(&overflow, __dres+__delta != __dres) {
				result.GetValue().SetDval(__dres)
			} else {
				result.GetValue().SetLval(__lres)
			}
			if overflow != 0 {
				result.SetTypeInfo(5)
			} else {
				result.SetTypeInfo(4)
			}
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto mul_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		mul_double:
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __z *Zval = result
			__z.GetValue().SetDval(d1 * d2)
			__z.SetTypeInfo(5)
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto mul_double
		}
	}
	return zend_mul_helper_SPEC(op1, op2, execute_data)
}
func ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			if op2.GetValue().GetLval() == 0 {
				return zend_mod_by_zero_helper_SPEC(execute_data)
			} else if op2.GetValue().GetLval() == -1 {

				/* Prevent overflow error/crash if op1==ZEND_LONG_MIN */

				var __z *Zval = result
				__z.GetValue().SetLval(0)
				__z.SetTypeInfo(4)
			} else {
				var __z *Zval = result
				__z.GetValue().SetLval(op1.GetValue().GetLval() % op2.GetValue().GetLval())
				__z.SetTypeInfo(4)
			}
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	return zend_mod_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {

		/* Perform shift on unsigned numbers to get well-defined wrap behavior. */

		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(zend_long(zend_ulong(*op1).value.lval << op2.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_left_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 && zend_ulong(*op2).value.lval < 8*8 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() >> op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_shift_right_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() < op2.GetValue().GetLval() {
			is_smaller_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_double:
			if d1 < d2 {
				goto is_smaller_true
			} else {
				goto is_smaller_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_double
		}
	}
	return zend_is_smaller_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 {
		if op2.GetTypeInfo() == 4 {
			if op1.GetValue().GetLval() <= op2.GetValue().GetLval() {
			is_smaller_or_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_smaller_or_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetTypeInfo() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_smaller_or_equal_double
		}
	} else if op1.GetTypeInfo() == 5 {
		if op2.GetTypeInfo() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_smaller_or_equal_double:
			if d1 <= d2 {
				goto is_smaller_or_equal_true
			} else {
				goto is_smaller_or_equal_false
			}
		} else if op2.GetTypeInfo() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_smaller_or_equal_double
		}
	}
	return zend_is_smaller_or_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() | op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_or_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() & op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_and_helper_SPEC(op1, op2, execute_data)
}
func ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<1|1<<2|1<<3) == 1<<0 && (1<<1|1<<2|1<<3) == 1<<0 {

	} else if op1.GetTypeInfo() == 4 && op2.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(op1.GetValue().GetLval() ^ op2.GetValue().GetLval())
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_bw_xor_helper_SPEC(op1, op2, execute_data)
}
func ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() + op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	FastLongAddFunction(result, op1, op2)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() + op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() - op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	FastLongSubFunction(result, op1, op2)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() - op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetLval(op1.GetValue().GetLval() * op2.GetValue().GetLval())
	__z.SetTypeInfo(4)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	var overflow ZendLong
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __lres long = op1.GetValue().GetLval() * op2.GetValue().GetLval()
	var __dres long__double = long__double(op1.GetValue().GetLval() * long__double(op2.GetValue().GetLval()))
	var __delta long__double = long__double(__lres - __dres)
	if g.Assign(&overflow, __dres+__delta != __dres) {
		result.GetValue().SetDval(__dres)
	} else {
		result.GetValue().SetLval(__lres)
	}
	if overflow != 0 {
		result.SetTypeInfo(5)
	} else {
		result.SetTypeInfo(4)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __z *Zval = result
	__z.GetValue().SetDval(op1.GetValue().GetDval() * op2.GetValue().GetDval())
	__z.SetTypeInfo(5)
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() == op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() == op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() != op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() != op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() < op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() < op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetLval() <= op2.GetValue().GetLval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result != 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result int
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	result = op1.GetValue().GetDval() <= op2.GetValue().GetDval()

	if result == 0 {
		execute_data.SetOpline(opline + 2)
	} else {
		execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
	}
	return 0
	if result != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_LIST_r(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_LIST_r(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BW_NOT_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if op1.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(^(op1.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if (1<<1|1<<2) == 1<<3 && op1.GetType() == 0 {
		op1 = _zvalUndefinedOp1(execute_data)
	}
	BitwiseNotFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		if (1<<1|1<<2) == 1<<3 && orig_val_type == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	} else {
		if IZendIsTrue(val) == 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		ZvalPtrDtorNogc(free_op1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ECHO_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var z *Zval
	z = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if z.GetType() == 6 {
		var str *ZendString = z.GetValue().GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		}
	} else {
		var str *ZendString = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		} else if (1<<1|1<<2) == 1<<3 && z.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		ZendStringReleaseEx(str, 0)
	}
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_JMPZ_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var val *Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if (1<<1|1<<2) == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline++
	} else {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	ZvalPtrDtorNogc(free_op1)
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPNZ_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var val *Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if (1<<1|1<<2) == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	} else {
		opline++
	}
	ZvalPtrDtorNogc(free_op1)
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var val *Zval
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if (1<<1|1<<2) == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue()))
	} else {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	ZvalPtrDtorNogc(free_op1)
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var val *Zval
	var ret int
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if (1<<1|1<<2) == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	ret = IZendIsTrue(val)
	ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		opline++
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var val *Zval
	var ret int
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if (1<<1|1<<2) == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	ret = IZendIsTrue(val)
	ZvalPtrDtorNogc(free_op1)
	if ret != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		opline++
	}
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_FREE_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar())))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FE_FREE_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var var_ *Zval
	var opline *ZendOp = execute_data.GetOpline()
	var_ = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_.GetType() != 7 && var_.GetFeIterIdx() != uint32-1 {
		ZendHashIteratorDel(var_.GetFeIterIdx())
	}
	ZvalPtrDtorNogc(var_)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var arg *Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = arg
	var _z2 *Zval = value
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (1<<1 | 1<<2) == 1<<0 {
		if (arg.GetTypeInfo() & 0xff00) != 0 {
			ZvalAddrefP(arg)
		}
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_BOOL_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var free_op1 ZendFreeOp
	val = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if (1<<1|1<<2) == 1<<3 && orig_val_type == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	} else {
		if IZendIsTrue(val) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		} else {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		}
		ZvalPtrDtorNogc(free_op1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_CLONE_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var obj *Zval
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && obj.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	for {
		if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && obj.GetType() != 8 {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && obj.GetType() == 10 {
				obj = &(*obj).value.GetRef().GetVal()
				if obj.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			if (1<<1|1<<2) == 1<<3 && obj.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "__clone method called on non-object")
			ZvalPtrDtorNogc(free_op1)
			return 0
		}
		break
	}
	ce = obj.GetValue().GetObj().GetCe()
	clone = ce.GetClone()
	clone_call = obj.GetValue().GetObj().GetHandlers().GetCloneObj()
	if clone_call == nil {
		ZendThrowError(nil, "Trying to clone an uncloneable object of class %s", ce.GetName().GetVal())
		ZvalPtrDtorNogc(free_op1)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	if clone != nil && (clone.GetFnFlags()&1<<0) == 0 {
		scope = execute_data.GetFunc().GetOpArray().GetScope()
		if clone.GetScope() != scope {
			if (clone.GetFnFlags()&1<<2) != 0 || ZendCheckProtected(g.CondF(clone.GetPrototype() != nil, func() *ZendClassEntry { return clone.GetPrototype().GetScope() }, func() *ZendClassEntry { return clone.GetScope() }), scope) == 0 {
				ZendWrongCloneCall(clone, scope)
				ZvalPtrDtorNogc(free_op1)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
		}
	}
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetObj(clone_call(obj))
	__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var new_op_array *ZendOpArray
	var free_op1 ZendFreeOp
	var inc_filename *Zval
	inc_filename = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
	ZvalPtrDtorNogc(free_op1)
	if EG.GetException() != nil {
		if new_op_array != (*ZendOpArray)(zend_intptr_t-1) && new_op_array != nil {
			DestroyOpArray(new_op_array)
			_efree(new_op_array)
		}
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	} else if new_op_array == (*ZendOpArray)(zend_intptr_t-1) {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		}
	} else if new_op_array != nil {
		var return_value *Zval = nil
		var call *ZendExecuteData
		if opline.GetResultType() != 0 {
			return_value = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		}
		new_op_array.SetScope(execute_data.GetFunc().GetOpArray().GetScope())
		call = ZendVmStackPushCallFrame(execute_data.GetThis().GetTypeInfo()&(8|1<<0<<8|1<<1<<8)|(1<<16|0<<17)|1<<20, (*ZendFunction)(new_op_array), 0, execute_data.GetThis().GetValue().GetPtr())
		if (execute_data.GetThis().GetTypeInfo() & 1 << 20) != 0 {
			call.SetSymbolTable(execute_data.GetSymbolTable())
		} else {
			call.SetSymbolTable(ZendRebuildSymbolTable())
		}
		call.SetPrevExecuteData(execute_data)
		IInitCodeExecuteData(call, new_op_array, return_value)
		if ZendExecuteEx == ExecuteEx {
			return 1
		} else {
			call.GetThis().SetTypeInfo(call.GetThis().GetTypeInfo() | 1<<17)
			ZendExecuteEx(call)
			ZendVmStackFreeCallFrame(call)
		}
		DestroyOpArray(new_op_array)
		_efree(new_op_array)
		if EG.GetException() != nil {
			ZendRethrowException(execute_data)
			if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			}
			return 0
		}
	} else if opline.GetResultType() != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	r.Assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_STRLEN_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if value.GetType() == 6 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
		__z.SetTypeInfo(4)
		ZvalPtrDtorNogc(free_op1)
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		var strict ZendBool
		if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
			if value.GetType() == 6 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
				__z.SetTypeInfo(4)
				ZvalPtrDtorNogc(free_op1)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
		if (1<<1|1<<2) == 1<<3 && value.GetType() == 0 {
			value = _zvalUndefinedOp1(execute_data)
		}
		strict = (execute_data.GetFunc().GetFnFlags() & 1 << 31) != 0
		for {
			if strict == 0 {
				var str *ZendString
				var tmp Zval
				var _z1 *Zval = &tmp
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (_t & 0xff00) != 0 {
					ZendGcAddref(&_gc.gc)
				}
				if ZendParseArgStrWeak(&tmp, &str) != 0 {
					var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					__z.GetValue().SetLval(str.GetLen())
					__z.SetTypeInfo(4)
					ZvalPtrDtor(&tmp)
					break
				}
				ZvalPtrDtor(&tmp)
			}
			if EG.GetException() == nil {
				ZendInternalTypeError(strict, "strlen() expects parameter 1 to be string, %s given", ZendGetTypeByConst(value.GetType()))
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			break
		}
	}
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int = 0
	var free_op1 ZendFreeOp
	value = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
	type_check_resource:
		if value.GetType() != 9 || nil != ZendRsrcListGetRsrcType(value.GetValue().GetRes()) {
			result = 1
		}
	} else if ((1<<1|1<<2)&(1<<3|1<<2)) != 0 && value.GetType() == 10 {
		value = &(*value).value.GetRef().GetVal()
		if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
			goto type_check_resource
		}
	} else if (1<<1|1<<2) == 1<<3 && value.GetType() == 0 {
		result = (1 << 1 & opline.GetExtendedValue()) != 0
		_zvalUndefinedOp1(execute_data)
		if EG.GetException() != nil {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	}
	if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
		ZvalPtrDtorNogc(free_op1)
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
	} else {
		for {

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
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if ((1<<1|1<<2) == 1<<0 || op1.GetType() == 6) && (1<<0 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if (1<<1|1<<2) != 1<<0 && op1_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<0 != 1<<0 && op2_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && (1<<1|1<<2) != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (1<<1|1<<2) == 1<<3 && op1.GetType() == 0 {
			op1 = _zvalUndefinedOp1(execute_data)
		}
		if 1<<0 == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		ZvalPtrDtorNogc(free_op1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1 | 1<<2) != 1<<0 {
		if container.GetType() == 7 {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, 1<<0, 0, execute_data)
			var _z3 *Zval = value
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
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		}
	} else {
		zend_fetch_dimension_address_read_R(container, dim, 1<<0, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_read_IS(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && container.GetType() != 8 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if (1<<1|1<<2) == 1<<3 && container.GetType() == 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
							if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && container.GetType() != 8 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
							if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if ((1<<1|1<<2) == 1<<0 || op1.GetType() == 6) && (1<<0 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if (1<<1|1<<2) != 1<<0 && op1_str.GetLen() == 0 {
			if 1<<0 == 1<<0 || 1<<0 == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<0 != 1<<0 && op2_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && (1<<1|1<<2) != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			if len_ > SIZE_MAX-(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1+8 - 1 & ^(8-1))-op2_str.GetLen() {
				ZendErrorNoreturn(1<<0, "Integer overflow in memory allocation")
			}
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if (1<<1 | 1<<2) == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if (1<<1|1<<2) == 1<<3 && op1.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	if 1<<0 == 1<<0 {
		op2_str = op2.GetValue().GetStr()
	} else if op2.GetType() == 6 {
		op2_str = ZendStringCopy(op2.GetValue().GetStr())
	} else {
		if 1<<0 == 1<<3 && op2.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		if (1<<1 | 1<<2) != 1<<0 {
			if op1_str.GetLen() == 0 {
				if 1<<0 == 1<<0 {
					if op2.GetTypeFlags() != 0 {
						ZendGcAddref(&op2_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		if 1<<0 != 1<<0 {
			if op2_str.GetLen() == 0 {
				if (1<<1 | 1<<2) == 1<<0 {
					if op1.GetTypeFlags() != 0 {
						ZendGcAddref(&op1_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if (1<<1 | 1<<2) != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && object.GetType() == 0 {
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
					ZvalPtrDtorNogc(free_op1)
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op1)
			return 0
			break
		}
	}
	if (1<<1 | 1<<2) != 0 {
		for {
			if (1<<1|1<<2) == 1<<0 || object.GetType() != 8 {
				if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if (1<<1|1<<2) == 1<<3 && object.GetType() == 0 {
					object = _zvalUndefinedOp1(execute_data)
					if EG.GetException() != nil {
						if 1<<0 != 1<<0 {

						}
						return 0
					}
				}
				if 1<<0 == 1<<0 {
					function_name = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				}
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op1)
				return 0
			}
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
			ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if ((1<<1|1<<2)&(1<<2|1<<1)) != 0 && obj != orig_obj {

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
		ZvalPtrDtorNogc(free_op1)
		if ((1<<1|1<<2)&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if ((1<<1 | 1<<2) & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if (1<<1 | 1<<2) == 1<<3 {
			ZendGcAddref(&obj.gc)
		} else if free_op1 != object {
			ZendGcAddref(&obj.gc)
			ZvalPtrDtorNogc(free_op1)
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
func ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			case_true:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline(opline + 2)
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			case_false:
				if (opline + 1).GetOpcode() == 44 {
					execute_data.SetOpline(opline + 2)
					return 0
				} else if (opline + 1).GetOpcode() == 43 {
					execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
					if EG.GetVmInterrupt() != 0 {
						return zend_interrupt_helper_SPEC(execute_data)
					}
					return 0
				}
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto case_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		case_double:
			if d1 == d2 {
				goto case_true
			} else {
				goto case_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto case_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if result != 0 {
				goto case_true
			} else {
				goto case_false
			}
		}
	}
	return zend_case_helper_SPEC(op1, op2, execute_data)
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if container.GetType() == 7 {
		var ht *HashTable
		var value *Zval
		var str *ZendString
	isset_dim_obj_array:
		ht = container.GetValue().GetArr()
	isset_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<0 != 1<<0 {
				if _zendHandleNumericStr(str.GetVal(), str.GetLen(), &hval) != 0 {
					goto num_index_prop
				}
			}
			value = ZendHashFindExInd(ht, str, 1<<0 == 1<<0)
		} else if offset.GetType() == 4 {
			hval = offset.GetValue().GetLval()
		num_index_prop:
			value = ZendHashIndexFind(ht, hval)
		} else if (1<<0&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
			offset = &(*offset).value.GetRef().GetVal()
			goto isset_again
		} else {
			value = ZendFindArrayDimSlow(ht, offset, execute_data)
			if EG.GetException() != nil {
				result = 0
				goto isset_dim_obj_exit
			}
		}
		if (opline.GetExtendedValue() & 1 << 0) == 0 {

			/* > IS_NULL means not IS_UNDEF and not IS_NULL */

			result = value != nil && value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1)
			if ((1<<1 | 1<<2) & (1<<0 | 1<<3)) != 0 {

				/* avoid exception check */

				for {

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
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto isset_dim_obj_array
		}
	}
	if 1<<0 == 1<<0 && offset.GetU2Extra() == 1 {
		offset++
	}
	if (opline.GetExtendedValue() & 1 << 0) == 0 {
		result = ZendIssetDimSlow(container, offset, execute_data)
	} else {
		result = ZendIsemptyDimSlow(container, offset, execute_data)
	}
isset_dim_obj_exit:
	ZvalPtrDtorNogc(free_op1)
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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var result int
	var offset *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && container.GetType() != 8 {
		if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
	ZvalPtrDtorNogc(free_op1)
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	subject = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if subject.GetType() == 7 {
	array_key_exists_array:
		ht = subject.GetValue().GetArr()
		result = ZendArrayKeyExistsFast(ht, key, opline, execute_data)
	} else {
		if (1<<0&(1<<2|1<<3)) != 0 && subject.GetType() == 10 {
			subject = &(*subject).value.GetRef().GetVal()
			if subject.GetType() == 7 {
				goto array_key_exists_array
			}
		}
		result = ZendArrayKeyExistsSlow(subject, key, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op1)
	for {
		if EG.GetException() != nil {
			break
		}
		if (opline + 1).GetOpcode() == 43 {
			if result == 3 {
				execute_data.SetOpline(opline + 2)
			} else {
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
			}
		} else if (opline + 1).GetOpcode() == 44 {
			if result != 3 {
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
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(result)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|UNUSED|CV, ANY) */

func ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var expr *Zval
	var result ZendBool
	expr = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
try_instanceof:
	if expr.GetType() == 8 {
		var ce *ZendClassEntry
		if 1<<0 == 1<<0 {
			ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
			if ce == nil {
				ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1).GetValue().GetStr(), 0x80)
				if ce != nil {
					(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
				}
			}
		} else if 1<<0 == 0 {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				r.Assert(EG.GetException() != nil)
				ZvalPtrDtorNogc(free_op1)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
		} else {
			ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())).GetValue().GetCe()
		}
		result = ce != nil && InstanceofFunction(expr.GetValue().GetObj().GetCe(), ce) != 0
	} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && expr.GetType() == 10 {
		expr = &(*expr).value.GetRef().GetVal()
		goto try_instanceof
	} else {
		if (1<<1|1<<2) == 1<<3 && expr.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		result = 0
	}
	ZvalPtrDtorNogc(free_op1)
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
func ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	var offset ZendLong
	var ht *HashTable
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if container.GetType() == 7 {
	fetch_dim_r_index_array:
		if dim.GetType() == 4 {
			offset = dim.GetValue().GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetValue().GetArr()
		if (ht.GetUFlags() & 1 << 2) != 0 {
			if zend_ulong(offset) < zend_ulong(ht).nNumUsed {
				value = &ht.arData[offset].GetVal()
				if value.GetType() == 0 {
					goto fetch_dim_r_index_undef
				}
			} else {
				goto fetch_dim_r_index_undef
			}
		} else {
			value = _zendHashIndexFind(ht, offset)
			if value == nil {
				goto fetch_dim_r_index_undef
			}
		}
		var _z3 *Zval = value
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
		if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc(free_op1)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else if (1<<1|1<<2) != 1<<0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
			dim++
		}
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		ZvalPtrDtorNogc(free_op1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
fetch_dim_r_index_undef:
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	ZendUndefinedOffset(offset)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	var offset ZendLong
	var ht *HashTable
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if container.GetType() == 7 {
	fetch_dim_r_index_array:
		if dim.GetType() == 4 {
			offset = dim.GetValue().GetLval()
		} else {
			offset = ZvalGetLong(dim)
		}
		ht = container.GetValue().GetArr()
		if (ht.GetUFlags() & 1 << 2) != 0 {
			if zend_ulong(offset) < zend_ulong(ht).nNumUsed {
				value = &ht.arData[offset].GetVal()
				if value.GetType() == 0 {
					goto fetch_dim_r_index_undef
				}
			} else {
				goto fetch_dim_r_index_undef
			}
		} else {
			value = _zendHashIndexFind(ht, offset)
			if value == nil {
				goto fetch_dim_r_index_undef
			}
		}
		var _z3 *Zval = value
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
		if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
			ZvalPtrDtorNogc(free_op1)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			r.Assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else if (1<<1|1<<2) != 1<<0 && container.GetType() == 10 {
		container = &(*container).value.GetRef().GetVal()
		if container.GetType() == 7 {
			goto fetch_dim_r_index_array
		} else {
			goto fetch_dim_r_index_slow
		}
	} else {
	fetch_dim_r_index_slow:
		if (1<<1|1<<2|1<<3) == 1<<0 && dim.GetU2Extra() == 1 {
			dim++
		}
		zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		ZvalPtrDtorNogc(free_op1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
fetch_dim_r_index_undef:
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	ZendUndefinedOffset(offset)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if ((1<<1|1<<2) == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if (1<<1|1<<2) != 1<<0 && op1_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && (1<<1|1<<2) != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			if len_ > SIZE_MAX-(zend_long((*byte)(&((*ZendString)(nil).GetVal()))-(*byte)(nil))+1+8 - 1 & ^(8-1))-op2_str.GetLen() {
				ZendErrorNoreturn(1<<0, "Integer overflow in memory allocation")
			}
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if (1<<1|1<<2) == 1<<3 && op1.GetType() == 0 {
			op1 = _zvalUndefinedOp1(execute_data)
		}
		if (1<<1|1<<2) == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		ZvalPtrDtorNogc(free_op1)
		ZvalPtrDtorNogc(free_op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_equal_double:
			if d1 == d2 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result != 0 {
				goto is_equal_true
			} else {
				goto is_equal_false
			}
		}
	}
	return zend_is_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 && (1<<1|1<<2) == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				execute_data.SetOpline((*ZendOp)((*byte)(opline+1) + int((opline + 1).GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				r.Assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else if op2.GetType() == 5 {
			d1 = float64(op1.GetValue().GetLval())
			d2 = op2.GetValue().GetDval()
			goto is_not_equal_double
		}
	} else if op1.GetType() == 5 {
		if op2.GetType() == 5 {
			d1 = op1.GetValue().GetDval()
			d2 = op2.GetValue().GetDval()
		is_not_equal_double:
			if d1 != d2 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		} else if op2.GetType() == 4 {
			d1 = op1.GetValue().GetDval()
			d2 = float64(op2.GetValue().GetLval())
			goto is_not_equal_double
		}
	} else if op1.GetType() == 6 {
		if op2.GetType() == 6 {
			var result int = ZendFastEqualStrings(op1.GetValue().GetStr(), op2.GetValue().GetStr())
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op1)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZvalPtrDtorStr(op2)
			}
			if result == 0 {
				goto is_not_equal_true
			} else {
				goto is_not_equal_false
			}
		}
	}
	return zend_is_not_equal_helper_SPEC(op1, op2, execute_data)
}
func ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1 | 1<<2) != 1<<0 {
		if container.GetType() == 7 {
		fetch_dim_r_array:
			value = ZendFetchDimensionAddressInner(container.GetValue().GetArr(), dim, 1<<1|1<<2, 0, execute_data)
			var _z3 *Zval = value
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
		} else if container.GetType() == 10 {
			container = &(*container).value.GetRef().GetVal()
			if container.GetType() == 7 {
				goto fetch_dim_r_array
			} else {
				goto fetch_dim_r_slow
			}
		} else {
		fetch_dim_r_slow:
			if (1<<1|1<<2) == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			zend_fetch_dimension_address_read_R_slow(container, dim, opline, execute_data)
		}
	} else {
		zend_fetch_dimension_address_read_R(container, dim, 1<<1|1<<2, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && container.GetType() != 8 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if (1<<1|1<<2) == 1<<3 && container.GetType() == 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
							if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) != 0 && container.GetType() != 8 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
							if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
					if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
	ZvalPtrDtorNogc(free_op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if ((1<<1|1<<2) == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if (1<<1|1<<2) != 1<<0 && op1_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
			if (1<<1|1<<2) == 1<<0 || (1<<1|1<<2) == 1<<3 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					ZendGcAddref(&__s.gc)
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			} else {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && (1<<1|1<<2) != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
			var len_ int = op1_str.GetLen()
			str = ZendStringExtend(op1_str, len_+op2_str.GetLen(), 0)
			memcpy(str.GetVal()+len_, op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else {
			str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
			memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
			memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
			var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var __s *ZendString = str
			__z.GetValue().SetStr(__s)
			__z.SetTypeInfo(6 | 1<<0<<8)
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		r.Assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if (1<<1 | 1<<2) == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if (1<<1|1<<2) == 1<<3 && op1.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		op1_str = ZvalGetStringFunc(op1)
	}
	if (1<<1 | 1<<2) == 1<<0 {
		op2_str = op2.GetValue().GetStr()
	} else if op2.GetType() == 6 {
		op2_str = ZendStringCopy(op2.GetValue().GetStr())
	} else {
		if (1<<1|1<<2) == 1<<3 && op2.GetType() == 0 {
			_zvalUndefinedOp2(execute_data)
		}
		op2_str = ZvalGetStringFunc(op2)
	}
	for {
		if (1<<1 | 1<<2) != 1<<0 {
			if op1_str.GetLen() == 0 {
				if (1<<1 | 1<<2) == 1<<0 {
					if op2.GetTypeFlags() != 0 {
						ZendGcAddref(&op2_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op2_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op1_str, 0)
				break
			}
		}
		if (1<<1 | 1<<2) != 1<<0 {
			if op2_str.GetLen() == 0 {
				if (1<<1 | 1<<2) == 1<<0 {
					if op1.GetTypeFlags() != 0 {
						ZendGcAddref(&op1_str.gc)
					}
				}
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var __s *ZendString = op1_str
				__z.GetValue().SetStr(__s)
				if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
					__z.SetTypeInfo(6)
				} else {
					__z.SetTypeInfo(6 | 1<<0<<8)
				}
				ZendStringReleaseEx(op2_str, 0)
				break
			}
		}
		str = ZendStringAlloc(op1_str.GetLen()+op2_str.GetLen(), 0)
		memcpy(str.GetVal(), op1_str.GetVal(), op1_str.GetLen())
		memcpy(str.GetVal()+op1_str.GetLen(), op2_str.GetVal(), op2_str.GetLen()+1)
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = str
		__z.GetValue().SetStr(__s)
		__z.SetTypeInfo(6 | 1<<0<<8)
		if (1<<1 | 1<<2) != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if (1<<1 | 1<<2) != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op1)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = _getZvalPtrVar(opline.GetOp1().GetVar(), &free_op1, execute_data)
	if (1<<1|1<<2) == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	if (1<<1 | 1<<2) != 1<<0 {
		function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	}
	if (1<<1|1<<2) != 1<<0 && function_name.GetType() != 6 {
		for {
			if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && function_name.GetType() == 10 {
				function_name = &(*function_name).value.GetRef().GetVal()
				if function_name.GetType() == 6 {
					break
				}
			} else if (1<<1|1<<2) == 1<<3 && function_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					ZvalPtrDtorNogc(free_op1)
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			ZvalPtrDtorNogc(free_op1)
			return 0
			break
		}
	}
	if (1<<1 | 1<<2) != 0 {
		for {
			if (1<<1|1<<2) == 1<<0 || object.GetType() != 8 {
				if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if (1<<1|1<<2) == 1<<3 && object.GetType() == 0 {
					object = _zvalUndefinedOp1(execute_data)
					if EG.GetException() != nil {
						if (1<<1 | 1<<2) != 1<<0 {
							ZvalPtrDtorNogc(free_op2)
						}
						return 0
					}
				}
				if (1<<1 | 1<<2) == 1<<0 {
					function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
				}
				ZendInvalidMethodCall(object, function_name)
				ZvalPtrDtorNogc(free_op2)
				ZvalPtrDtorNogc(free_op1)
				return 0
			}
			break
		}
	}
	obj = object.GetValue().GetObj()
	called_scope = obj.GetCe()
	if (1<<1|1<<2) == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == called_scope {
		fbc = (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]
	} else {
		var orig_obj *ZendObject = obj
		if (1<<1 | 1<<2) == 1<<0 {
			function_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		}

		/* First, locate the function. */

		fbc = obj.GetHandlers().GetGetMethod()(&obj, function_name.GetValue().GetStr(), g.CondF1((1<<1|1<<2) == 1<<0, func() *Zval { return (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1 }, nil))
		if fbc == nil {
			if EG.GetException() == nil {
				ZendUndefinedMethod(obj.GetCe(), function_name.GetValue().GetStr())
			}
			ZvalPtrDtorNogc(free_op2)
			ZvalPtrDtorNogc(free_op1)
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if ((1<<1|1<<2)&(1<<2|1<<1)) != 0 && obj != orig_obj {

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
	if (1<<1 | 1<<2) != 1<<0 {
		ZvalPtrDtorNogc(free_op2)
	}
	call_info = 0<<16 | 0<<17 | (8 | 1<<0<<8 | 1<<1<<8)
	if (fbc.GetFnFlags() & 1 << 4) != 0 {
		ZvalPtrDtorNogc(free_op1)
		if ((1<<1|1<<2)&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if ((1<<1 | 1<<2) & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if (1<<1 | 1<<2) == 1<<3 {
			ZendGcAddref(&obj.gc)
		} else if free_op1 != object {
			ZendGcAddref(&obj.gc)
			ZvalPtrDtorNogc(free_op1)
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
