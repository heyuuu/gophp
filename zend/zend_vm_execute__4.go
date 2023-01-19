// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if (1<<1 | 1<<2) == 1<<0 {
		ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
	} else {
		ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if (1<<1 | 1<<2) == 1<<0 {
		ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
	} else {
		ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var rope **ZendString
	var var_ *Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**ZendString)((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
	if (1<<1 | 1<<2) == 1<<0 {
		var_ = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		rope[0] = var_.GetValue().GetStr()
		if var_.GetTypeFlags() != 0 {
			ZvalAddrefP(var_)
		}
	} else {
		var_ = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
		if var_.GetType() == 6 {
			if (1<<1 | 1<<2) == 1<<3 {
				rope[0] = ZendStringCopy(var_.GetValue().GetStr())
			} else {
				rope[0] = var_.GetValue().GetStr()
			}
		} else {
			if (1<<1|1<<2) == 1<<3 && var_.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			rope[0] = ZvalGetStringFunc(var_)
			ZvalPtrDtorNogc(free_op2)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var free_op2 ZendFreeOp
	var class_name *Zval
	var opline *ZendOp = execute_data.GetOpline()
	if (1<<1 | 1<<2) == 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	} else if (1<<1 | 1<<2) == 1<<0 {
		var ce *ZendClassEntry = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
		if ce == nil {
			class_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
			ce = ZendFetchClassByName(class_name.GetValue().GetStr(), (class_name + 1).GetValue().GetStr(), opline.GetOp1().GetNum())
			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ce)
	} else {
		class_name = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	try_class_name:
		if class_name.GetType() == 8 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(class_name.GetValue().GetObj().GetCe())
		} else if class_name.GetType() == 6 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(class_name.GetValue().GetStr(), opline.GetOp1().GetNum()))
		} else if ((1<<1|1<<2)&(1<<2|1<<3)) != 0 && class_name.GetType() == 10 {
			class_name = &(*class_name).value.GetRef().GetVal()
			goto try_class_name
		} else {
			if (1<<1|1<<2) == 1<<3 && class_name.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Class name must be a valid object or a string")
		}
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
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
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
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
					return 0
				}
			}
			ZendThrowError(nil, "Method name must be a string")
			ZvalPtrDtorNogc(free_op2)
			return 0
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
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
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
	if (1<<1 | 1<<2) != 1<<0 {
		ZvalPtrDtorNogc(free_op2)
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
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
				assert(EG.GetException() != nil)
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
				return 0
			}
			if (1<<1 | 1<<2) != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())))
			return 0
		}
	}
	if 0 == 1<<0 && (1<<1|1<<2) == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 0 != 1<<0 && (1<<1|1<<2) == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = array
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)

	container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var result int
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
func ZEND_YIELD_SPEC_UNUSED_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_YIELD_SPEC_UNUSED_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_FETCH_CLASS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var class_name *Zval
	var opline *ZendOp = execute_data.GetOpline()
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
				assert(EG.GetException() != nil)
				return 0
			}
			if 0 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			return 0
		}
	}
	if 0 == 1<<0 && 0 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 0 != 1<<0 && 0 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_VERIFY_RETURN_TYPE_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendVerifyMissingReturnType(execute_data.GetFunc(), (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum())))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_NEW_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var result *Zval
	var constructor *ZendFunction
	var ce *ZendClassEntry
	var call *ZendExecuteData
	if 0 == 1<<0 {
		ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetOp2().GetNum()))[0]
		if ce == nil {
			ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp1()).constant) + 1).GetValue().GetStr(), 0|0x200)
			if ce == nil {
				assert(EG.GetException() != nil)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetOp2().GetNum()))[0] = ce
		}
	} else {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = array
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_YIELD_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_FETCH_THIS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if execute_data.GetThis().GetType() == 8 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __z *Zval = result
		__z.GetValue().SetObj(execute_data.GetThis().GetValue().GetObj())
		__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		ZvalAddrefP(result)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
}
func ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if (opline.GetExtendedValue()&1<<0^execute_data.GetThis().GetType() == 8) != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
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
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if execute_data.GetThis().GetType() == 8 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = execute_data.GetThis().GetValue().GetObj().GetCe().GetName()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	} else if execute_data.GetThis().GetValue().GetCe() != nil {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var __s *ZendString = execute_data.GetThis().GetValue().GetCe().GetName()
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			ZendGcAddref(&__s.gc)
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
	} else {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if execute_data.GetFunc().GetScope() == nil {
			ZendError(1<<1, "get_called_class() called from outside a class")
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(execute_data.GetThis().GetNumArgs())
	__z.SetTypeInfo(4)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FUNC_GET_ARGS_SPEC_UNUSED_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var ht *ZendArray
	var arg_count uint32
	var result_size uint32
	var skip uint32
	arg_count = execute_data.GetThis().GetNumArgs()
	if 0 == 1<<0 {
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
		assert((__fill_ht.GetUFlags() & 1 << 2) != 0)
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMP|VAR|CV, UNUSED|CONST|TMPVAR) */

func ZEND_PRE_INC_OBJ_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)

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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)

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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
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
						assert(EG.GetException() == nil)
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
func ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
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
func ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
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
func ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
	if 1<<3 == 1<<0 {
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
						assert(EG.GetException() == nil)
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (0 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 0, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
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
func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = &(execute_data.GetThis())
	if object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)

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

	/* assign_obj has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 == 1<<0 {
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

func ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<3 == 1<<0 {
		ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
	} else {
		ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ROPE_INIT_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var rope **ZendString
	var var_ *Zval

	/* Compiler allocates the necessary number of zval slots to keep the rope */

	rope = (**ZendString)((*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())))
	if 1<<3 == 1<<0 {
		var_ = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
		rope[0] = var_.GetValue().GetStr()
		if var_.GetTypeFlags() != 0 {
			ZvalAddrefP(var_)
		}
	} else {
		var_ = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
		if var_.GetType() == 6 {
			if 1<<3 == 1<<3 {
				rope[0] = ZendStringCopy(var_.GetValue().GetStr())
			} else {
				rope[0] = var_.GetValue().GetStr()
			}
		} else {
			if 1<<3 == 1<<3 && var_.GetType() == 0 {
				_zvalUndefinedOp2(execute_data)
			}
			rope[0] = ZvalGetStringFunc(var_)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FETCH_CLASS_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var class_name *Zval
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<3 == 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(nil, opline.GetOp1().GetNum()))
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	} else if 1<<3 == 1<<0 {
		var ce *ZendClassEntry = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
		if ce == nil {
			class_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
			ce = ZendFetchClassByName(class_name.GetValue().GetStr(), (class_name + 1).GetValue().GetStr(), opline.GetOp1().GetNum())
			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ce)
	} else {
		class_name = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	try_class_name:
		if class_name.GetType() == 8 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(class_name.GetValue().GetObj().GetCe())
		} else if class_name.GetType() == 6 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).GetValue().SetCe(ZendFetchClass(class_name.GetValue().GetStr(), opline.GetOp1().GetNum()))
		} else if (1<<3&(1<<2|1<<3)) != 0 && class_name.GetType() == 10 {
			class_name = &(*class_name).value.GetRef().GetVal()
			goto try_class_name
		} else {
			if 1<<3 == 1<<3 && class_name.GetType() == 0 {
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
func ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	if 1<<3 != 1<<0 {

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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
				assert(EG.GetException() != nil)
				return 0
			}
			if 1<<3 != 1<<0 {
				(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] = ce
			}
		}
	} else {
		ce = ZendFetchClass(nil, opline.GetOp1().GetNum())
		if ce == nil {
			assert(EG.GetException() != nil)
			return 0
		}
	}
	if 0 == 1<<0 && 1<<3 == 1<<0 && g.Assign(&fbc, (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetResult().GetNum() + g.SizeOf("void *"))))[0]) != nil {

	} else if 0 != 1<<0 && 1<<3 == 1<<0 && (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))[0] == ce {
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var __arr *ZendArray = _zendNewArray(0)
	var __z *Zval = array
	__z.GetValue().SetArr(__arr)
	__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)

	container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = &(execute_data.GetThis())
	if container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
func ZEND_YIELD_SPEC_UNUSED_CV_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_BW_NOT_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if op1.GetTypeInfo() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(^(op1.GetValue().GetLval()))
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<3 && op1.GetType() == 0 {
		op1 = _zvalUndefinedOp1(execute_data)
	}
	BitwiseNotFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_NOT_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		if 1<<3 == 1<<3 && orig_val_type == 0 {
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
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func zend_pre_inc_helper_SPEC_CV(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && var_ptr.GetType() == 15 {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<3 && var_ptr.GetType() == 0 {
		var_ptr.SetTypeInfo(1)
		_zvalUndefinedOp1(execute_data)
	}
	for {
		if var_ptr.GetType() == 10 {
			var ref *ZendReference = var_ptr.GetValue().GetRef()
			var_ptr = &(*var_ptr).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendIncdecTypedRef(ref, nil, opline, execute_data)
				break
			}
		}
		IncrementFunction(var_ptr)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		FastLongIncrementFunction(var_ptr)

		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_pre_inc_helper_SPEC_CV(execute_data)
}
func ZEND_PRE_INC_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		FastLongIncrementFunction(var_ptr)
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = var_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_pre_inc_helper_SPEC_CV(execute_data)
}
func zend_pre_dec_helper_SPEC_CV(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && var_ptr.GetType() == 15 {
		if opline.GetResultType() != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<3 && var_ptr.GetType() == 0 {
		var_ptr.SetTypeInfo(1)
		_zvalUndefinedOp1(execute_data)
	}
	for {
		if var_ptr.GetType() == 10 {
			var ref *ZendReference = var_ptr.GetValue().GetRef()
			var_ptr = &(*var_ptr).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendIncdecTypedRef(ref, nil, opline, execute_data)
				break
			}
		}
		DecrementFunction(var_ptr)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_DEC_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		FastLongDecrementFunction(var_ptr)

		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_pre_dec_helper_SPEC_CV(execute_data)
}
func ZEND_PRE_DEC_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		FastLongDecrementFunction(var_ptr)
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = var_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_pre_dec_helper_SPEC_CV(execute_data)
}
func zend_post_inc_helper_SPEC_CV(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && var_ptr.GetType() == 15 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<3 && var_ptr.GetType() == 0 {
		var_ptr.SetTypeInfo(1)
		_zvalUndefinedOp1(execute_data)
	}
	for {
		if var_ptr.GetType() == 10 {
			var ref *ZendReference = var_ptr.GetValue().GetRef()
			var_ptr = &(*var_ptr).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendIncdecTypedRef(ref, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), opline, execute_data)
				break
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = var_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		IncrementFunction(var_ptr)
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
		__z.SetTypeInfo(4)
		FastLongIncrementFunction(var_ptr)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_post_inc_helper_SPEC_CV(execute_data)
}
func zend_post_dec_helper_SPEC_CV(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && var_ptr.GetType() == 15 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<3 && var_ptr.GetType() == 0 {
		var_ptr.SetTypeInfo(1)
		_zvalUndefinedOp1(execute_data)
	}
	for {
		if var_ptr.GetType() == 10 {
			var ref *ZendReference = var_ptr.GetValue().GetRef()
			var_ptr = &(*var_ptr).value.GetRef().GetVal()
			if ref.GetSources().GetPtr() != nil {
				ZendIncdecTypedRef(ref, (*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), opline, execute_data)
				break
			}
		}
		var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z2 *Zval = var_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
		DecrementFunction(var_ptr)
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_DEC_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_ptr.GetType() == 4 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
		__z.SetTypeInfo(4)
		FastLongDecrementFunction(var_ptr)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	return zend_post_dec_helper_SPEC_CV(execute_data)
}
func ZEND_ECHO_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var z *Zval
	z = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if z.GetType() == 6 {
		var str *ZendString = z.GetValue().GetStr()
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		}
	} else {
		var str *ZendString = ZvalGetStringFunc(z)
		if str.GetLen() != 0 {
			ZendWrite(str.GetVal(), str.GetLen())
		} else if 1<<3 == 1<<3 && z.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		ZendStringReleaseEx(str, 0)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_JMPZ_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<3 == 1<<3 && val.GetTypeInfo() == 0 {
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
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPNZ_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<3 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			if EG.GetException() != nil {
				return 0
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if IZendIsTrue(val) != 0 {
		opline = (*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset()))
	} else {
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
func ZEND_JMPZNZ_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetExtendedValue())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		if 1<<3 == 1<<3 && val.GetTypeInfo() == 0 {
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
	if EG.GetException() != nil {
		return 0
	}
	execute_data.SetOpline(opline)
	if EG.GetVmInterrupt() != 0 {
		return zend_interrupt_helper_SPEC(execute_data)
	}
	return 0
}
func ZEND_JMPZ_EX_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var ret int
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<3 == 1<<3 && val.GetTypeInfo() == 0 {
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
func ZEND_JMPNZ_EX_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	var ret int
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	} else if val.GetTypeInfo() <= 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<3 == 1<<3 && val.GetTypeInfo() == 0 {
			_zvalUndefinedOp1(execute_data)
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	}
	ret = IZendIsTrue(val)
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
func ZEND_RETURN_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval_ptr *Zval
	var return_value *Zval
	var free_op1 ZendFreeOp
	retval_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	return_value = execute_data.GetReturnValue()
	if 1<<3 == 1<<3 && retval_ptr.GetTypeInfo() == 0 {
		retval_ptr = _zvalUndefinedOp1(execute_data)
		if return_value != nil {
			return_value.SetTypeInfo(1)
		}
	} else if return_value == nil {
		if (1 << 3 & (1<<2 | 1<<1)) != 0 {
			if free_op1.GetTypeFlags() != 0 && ZvalDelrefP(free_op1) == 0 {
				RcDtorFunc(free_op1.GetValue().GetCounted())
			}
		}
	} else {
		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			var _z1 *Zval = return_value
			var _z2 *Zval = retval_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<3 == 1<<0 {
				if (return_value.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(return_value)
				}
			}
		} else if 1<<3 == 1<<3 {
			for {
				if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
					if (retval_ptr.GetTypeInfo() & 0xff) != 10 {
						if (execute_data.GetThis().GetTypeInfo() & 1 << 16) == 0 {
							var ref *ZendRefcounted = retval_ptr.GetValue().GetCounted()
							var _z1 *Zval = return_value
							var _z2 *Zval = retval_ptr
							var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
							var _t uint32 = _z2.GetTypeInfo()
							_z1.GetValue().SetCounted(_gc)
							_z1.SetTypeInfo(_t)
							if (ref.GetGc().GetTypeInfo() & (0xfffffc00 | 1<<4<<0)) == 1<<4<<0 {
								GcPossibleRoot(ref)
							}
							retval_ptr.SetTypeInfo(1)
							break
						} else {
							ZvalAddrefP(retval_ptr)
						}
					} else {
						retval_ptr = &(*retval_ptr).value.GetRef().GetVal()
						if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(retval_ptr)
						}
					}
				}
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				break
			}
		} else {
			if retval_ptr.GetType() == 10 {
				var ref *ZendRefcounted = retval_ptr.GetValue().GetCounted()
				retval_ptr = &(*retval_ptr).value.GetRef().GetVal()
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if ZendGcDelref(&ref.gc) == 0 {
					_efree(ref)
				} else if (retval_ptr.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(retval_ptr)
				}
			} else {
				var _z1 *Zval = return_value
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			}
		}
	}
	return zend_leave_helper_SPEC(execute_data)
}
func ZEND_RETURN_BY_REF_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval_ptr *Zval
	for {
		if (1<<3&(1<<0|1<<1)) != 0 || 1<<3 == 1<<2 && opline.GetExtendedValue() == 1<<1 {

			/* Not supposed to happen, but we'll allow it */

			ZendError(1<<3, "Only variable references should be returned by reference")
			retval_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
			if execute_data.GetReturnValue() == nil {

			} else {
				if 1<<3 == 1<<2 && retval_ptr.GetType() == 10 {
					var _z1 *Zval = execute_data.GetReturnValue()
					var _z2 *Zval = retval_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					break
				}
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 1)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = retval_ptr
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				execute_data.GetReturnValue().GetValue().SetRef(_ref)
				execute_data.GetReturnValue().SetTypeInfo(10 | 1<<0<<8)
				if 1<<3 == 1<<0 {
					if retval_ptr.GetTypeFlags() != 0 {
						ZvalAddrefP(retval_ptr)
					}
				}
			}
			break
		}
		retval_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), execute_data)
		if 1<<3 == 1<<2 {
			assert(retval_ptr != &EG.uninitialized_zval)
			if opline.GetExtendedValue() == 1<<0 && retval_ptr.GetType() != 10 {
				ZendError(1<<3, "Only variable references should be returned by reference")
				if execute_data.GetReturnValue() != nil {
					var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
					ZendGcSetRefcount(&_ref.gc, 1)
					_ref.GetGc().SetTypeInfo(10)
					var _z1 *Zval = &_ref.val
					var _z2 *Zval = retval_ptr
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					_ref.GetSources().SetPtr(nil)
					execute_data.GetReturnValue().GetValue().SetRef(_ref)
					execute_data.GetReturnValue().SetTypeInfo(10 | 1<<0<<8)
				}
				break
			}
		}
		if execute_data.GetReturnValue() != nil {
			if retval_ptr.GetType() == 10 {
				ZvalAddrefP(retval_ptr)
			} else {
				var _z *Zval = retval_ptr
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
			var __z *Zval = execute_data.GetReturnValue()
			__z.GetValue().SetRef(retval_ptr.GetValue().GetRef())
			__z.SetTypeInfo(10 | 1<<0<<8)
		}
		break
	}
	return zend_leave_helper_SPEC(execute_data)
}
func ZEND_GENERATOR_RETURN_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var retval *Zval
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	retval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)

	/* Copy return value into generator->retval */

	if (1 << 3 & (1<<0 | 1<<1)) != 0 {
		var _z1 *Zval = &generator.retval
		var _z2 *Zval = retval
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 == 1<<0 {
			if (generator.GetRetval().GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(&(generator.GetRetval()))
			}
		}
	} else if 1<<3 == 1<<3 {
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
		var _z1 *Zval = &generator.retval
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		if retval.GetType() == 10 {
			var ref *ZendRefcounted = retval.GetValue().GetCounted()
			retval = &(*retval).value.GetRef().GetVal()
			var _z1 *Zval = &generator.retval
			var _z2 *Zval = retval
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZendGcDelref(&ref.gc) == 0 {
				_efree(ref)
			} else if (retval.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(retval)
			}
		} else {
			var _z1 *Zval = &generator.retval
			var _z2 *Zval = retval
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}

	/* Close the generator to free up resources */

	ZendGeneratorClose(generator, 1)

	/* Pass execution back to handling code */

	return -1

	/* Pass execution back to handling code */
}
func ZEND_THROW_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	for {
		if 1<<3 == 1<<0 || value.GetType() != 8 {
			if (1<<3&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
				value = &(*value).value.GetRef().GetVal()
				if value.GetType() == 8 {
					break
				}
			}
			if 1<<3 == 1<<3 && value.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
				if EG.GetException() != nil {
					return 0
				}
			}
			ZendThrowError(nil, "Can only throw objects")
			return 0
		}
		break
	}
	ZendExceptionSave()
	if 1<<3 != 1<<1 {
		if value.GetTypeFlags() != 0 {
			ZvalAddrefP(value)
		}
	}
	ZendThrowExceptionObject(value)
	ZendExceptionRestore()
	return 0
}
func ZEND_SEND_VAR_SPEC_CV_INLINE_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	varptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<3 && varptr.GetTypeInfo() == 0 {
		_zvalUndefinedOp1(execute_data)
		arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
		arg.SetTypeInfo(1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<3 {
		var _z3 *Zval = varptr
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
		var _z1 *Zval = arg
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		if varptr.GetType() == 10 {
			var ref *ZendRefcounted = varptr.GetValue().GetCounted()
			varptr = &(*varptr).value.GetRef().GetVal()
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZendGcDelref(&ref.gc) == 0 {
				_efree(ref)
			} else if (arg.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(arg)
			}
		} else {
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAR_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	return ZEND_SEND_VAR_SPEC_CV_INLINE_HANDLER(execute_data)
}
func ZEND_SEND_REF_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	varptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), execute_data)
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<2 && varptr.GetType() == 15 {
		var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
		ZendGcSetRefcount(&_ref.gc, 1)
		_ref.GetGc().SetTypeInfo(10)
		_ref.GetSources().SetPtr(nil)
		arg.GetValue().SetRef(_ref)
		arg.SetTypeInfo(10 | 1<<0<<8)
		&(*arg).value.GetRef().GetVal().u1.type_info = 1
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if varptr.GetType() == 10 {
		ZvalAddrefP(varptr)
	} else {
		var _z *Zval = varptr
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
	var __z *Zval = arg
	__z.GetValue().SetRef(varptr.GetValue().GetRef())
	__z.SetTypeInfo(10 | 1<<0<<8)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAR_EX_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), arg_num, 1|2) != 0 {
	send_var_by_ref:
		return ZEND_SEND_REF_SPEC_CV_HANDLER(execute_data)
	}
	varptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<3 && varptr.GetTypeInfo() == 0 {
		_zvalUndefinedOp1(execute_data)
		arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
		arg.SetTypeInfo(1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<3 {
		var _z3 *Zval = varptr
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
		var _z1 *Zval = arg
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		if varptr.GetType() == 10 {
			var ref *ZendRefcounted = varptr.GetValue().GetCounted()
			varptr = &(*varptr).value.GetRef().GetVal()
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZendGcDelref(&ref.gc) == 0 {
				_efree(ref)
			} else if (arg.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(arg)
			}
		} else {
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAR_EX_SPEC_CV_QUICK_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if (execute_data.GetCall().GetFunc().GetQuickArgFlags() >> (arg_num + 3) * 2 & (1 | 2)) != 0 {
		goto send_var_by_ref
	}
	varptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<3 && varptr.GetTypeInfo() == 0 {
		_zvalUndefinedOp1(execute_data)
		arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
		arg.SetTypeInfo(1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<3 {
		var _z3 *Zval = varptr
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
		var _z1 *Zval = arg
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else {
		if varptr.GetType() == 10 {
			var ref *ZendRefcounted = varptr.GetValue().GetCounted()
			varptr = &(*varptr).value.GetRef().GetVal()
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZendGcDelref(&ref.gc) == 0 {
				_efree(ref)
			} else if (arg.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(arg)
			}
		} else {
			var _z1 *Zval = arg
			var _z2 *Zval = varptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_USER_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var arg *Zval
	var param *Zval
	if ZendCheckArgSendType(execute_data.GetCall().GetFunc(), opline.GetOp2().GetNum(), 1) != 0 {
		ZendParamMustBeRef(execute_data.GetCall().GetFunc(), opline.GetOp2().GetNum())
	}
	arg = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	param = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	var _z1 *Zval = param
	var _z2 *Zval = arg
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	if (_t & 0xff00) != 0 {
		ZendGcAddref(&_gc.gc)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var val *Zval
	val = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if val.GetTypeInfo() == 3 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
	} else if val.GetTypeInfo() <= 3 {

		/* The result and op1 can be the same cv zval */

		var orig_val_type uint32 = val.GetTypeInfo()
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
		if 1<<3 == 1<<3 && orig_val_type == 0 {
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
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_CLONE_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var obj *Zval
	var ce *ZendClassEntry
	var scope *ZendClassEntry
	var clone *ZendFunction
	var clone_call ZendObjectCloneObjT
	obj = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && obj.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	for {
		if 1<<3 == 1<<0 || 1<<3 != 0 && obj.GetType() != 8 {
			if (1<<3&(1<<2|1<<3)) != 0 && obj.GetType() == 10 {
				obj = &(*obj).value.GetRef().GetVal()
				if obj.GetType() == 8 {
					break
				}
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			if 1<<3 == 1<<3 && obj.GetType() == 0 {
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
func ZEND_CAST_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr *Zval
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var ht *HashTable
	expr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	switch opline.GetExtendedValue() {
	case 1:
		result.SetTypeInfo(1)
		break
	case 16:
		if ZendIsTrue(expr) != 0 {
			result.SetTypeInfo(3)
		} else {
			result.SetTypeInfo(2)
		}
		break
	case 4:
		var __z *Zval = result
		__z.GetValue().SetLval(ZvalGetLong(expr))
		__z.SetTypeInfo(4)
		break
	case 5:
		var __z *Zval = result
		__z.GetValue().SetDval(ZvalGetDouble(expr))
		__z.SetTypeInfo(5)
		break
	case 6:
		var __z *Zval = result
		var __s *ZendString = ZvalGetString(expr)
		__z.GetValue().SetStr(__s)
		if (ZvalGcFlags(__s.GetGc().GetTypeInfo()) & 1 << 6) != 0 {
			__z.SetTypeInfo(6)
		} else {
			__z.SetTypeInfo(6 | 1<<0<<8)
		}
		break
	default:
		if (1 << 3 & (1<<2 | 1<<3)) != 0 {
			if expr.GetType() == 10 {
				expr = &(*expr).value.GetRef().GetVal()
			}
		}

		/* If value is already of correct type, return it directly */

		if expr.GetType() == opline.GetExtendedValue() {
			var _z1 *Zval = result
			var _z2 *Zval = expr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<3 == 1<<0 {
				if (result.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(result)
				}
			} else if 1<<3 != 1<<1 {
				if (result.GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(result)
				}
			}
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		}
		if opline.GetExtendedValue() == 7 {
			if 1<<3 == 1<<0 || expr.GetType() != 8 || expr.GetValue().GetObj().GetCe() == ZendCeClosure {
				if expr.GetType() != 1 {
					var __arr *ZendArray = _zendNewArray(1)
					var __z *Zval = result
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					expr = ZendHashIndexAddNew(result.GetValue().GetArr(), 0, expr)
					if 1<<3 == 1<<0 {
						if (expr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(expr)
						}
					} else {
						if (expr.GetTypeInfo() & 0xff00) != 0 {
							ZvalAddrefP(expr)
						}
					}
				} else {
					var __z *Zval = result
					__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
					__z.SetTypeInfo(7)
				}
			} else {
				var obj_ht *HashTable = ZendGetPropertiesFor(expr, ZEND_PROP_PURPOSE_ARRAY_CAST)
				if obj_ht != nil {

					/* fast copy */

					var __arr *ZendArray = ZendProptableToSymtable(obj_ht, expr.GetValue().GetObj().GetCe().GetDefaultPropertiesCount() != 0 || expr.GetValue().GetObj().GetHandlers() != &StdObjectHandlers || (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<5) != 0)
					var __z *Zval = result
					__z.GetValue().SetArr(__arr)
					__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
					if obj_ht != nil && (ZvalGcFlags(obj_ht.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcDelref(&obj_ht.gc) == 0 {
						ZendArrayDestroy(obj_ht)
					}
				} else {
					var __z *Zval = result
					__z.GetValue().SetArr((*ZendArray)(&ZendEmptyArray))
					__z.SetTypeInfo(7)
				}
			}
		} else {
			var __z *Zval = result
			__z.GetValue().SetObj(ZendObjectsNew(ZendStandardClassDef))
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
			if expr.GetType() == 7 {
				ht = ZendSymtableToProptable(expr.GetValue().GetArr())
				if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) != 0 {

					/* TODO: try not to duplicate immutable arrays as well ??? */

					ht = ZendArrayDup(ht)

					/* TODO: try not to duplicate immutable arrays as well ??? */

				}
				result.GetValue().GetObj().SetProperties(ht)
			} else if expr.GetType() != 1 {
				ht = _zendNewArray(1)
				result.GetValue().GetObj().SetProperties(ht)
				expr = ZendHashAddNew(ht, ZendKnownStrings[ZEND_STR_SCALAR], expr)
				if 1<<3 == 1<<0 {
					if (expr.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(expr)
					}
				} else {
					if (expr.GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(expr)
					}
				}
			}
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INCLUDE_OR_EVAL_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var new_op_array *ZendOpArray
	var inc_filename *Zval
	inc_filename = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	new_op_array = ZendIncludeOrEval(inc_filename, opline.GetExtendedValue())
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_FE_RESET_R_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array_ptr *Zval
	var result *Zval
	array_ptr = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	if array_ptr.GetType() == 7 {
		result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = array_ptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 != 1<<1 && (result.GetTypeInfo()&0xff00) != 0 {
			ZvalAddrefP(array_ptr)
		}
		result.SetFePos(0)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if 1<<3 != 1<<0 && array_ptr.GetType() == 8 {
		if array_ptr.GetValue().GetObj().GetCe().GetGetIterator() == nil {
			var properties *HashTable
			if array_ptr.GetValue().GetObj().GetProperties() != nil && ZendGcRefcount(&(array_ptr.GetValue().GetObj().GetProperties()).gc) > 1 {
				if (ZvalGcFlags(array_ptr.GetValue().GetObj().GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(array_ptr.GetValue().GetObj().GetProperties()).gc)
				}
				array_ptr.GetValue().GetObj().SetProperties(ZendArrayDup(array_ptr.GetValue().GetObj().GetProperties()))
			}
			properties = array_ptr.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*array_ptr))
			result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z1 *Zval = result
			var _z2 *Zval = array_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if 1<<3 != 1<<1 {
				ZvalAddrefP(array_ptr)
			}
			if properties.GetNNumOfElements() == 0 {
				result.SetFeIterIdx(uint32 - 1)
				if EG.GetException() != nil {
					return 0
				}
				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			}
			result.SetFeIterIdx(ZendHashIteratorAdd(properties, 0))
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			var is_empty ZendBool = ZendFeResetIterator(array_ptr, 0, opline, execute_data)
			if EG.GetException() != nil {
				return 0
			} else if is_empty != 0 {

				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			} else {
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	} else {
		ZendError(1<<1, "Invalid argument supplied for foreach()")
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
		if EG.GetException() != nil {
			return 0
		}
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_FE_RESET_RW_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var array_ptr *Zval
	var array_ref *Zval
	if 1<<3 == 1<<2 || 1<<3 == 1<<3 {
		array_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
		array_ref = array_ptr
		if array_ref.GetType() == 10 {
			array_ptr = &(*array_ref).value.GetRef().GetVal()
		}
	} else {
		array_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
		array_ref = array_ptr
	}
	if array_ptr.GetType() == 7 {
		if 1<<3 == 1<<2 || 1<<3 == 1<<3 {
			if array_ptr == array_ref {
				var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
				ZendGcSetRefcount(&_ref.gc, 1)
				_ref.GetGc().SetTypeInfo(10)
				var _z1 *Zval = &_ref.val
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				_ref.GetSources().SetPtr(nil)
				array_ref.GetValue().SetRef(_ref)
				array_ref.SetTypeInfo(10 | 1<<0<<8)
				array_ptr = &(*array_ref).value.GetRef().GetVal()
			}
			ZvalAddrefP(array_ref)
			var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _z2 *Zval = array_ref
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		} else {
			array_ref = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
			var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&_ref.gc, 1)
			_ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &_ref.val
			var _z2 *Zval = array_ptr
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			_ref.GetSources().SetPtr(nil)
			array_ref.GetValue().SetRef(_ref)
			array_ref.SetTypeInfo(10 | 1<<0<<8)
			array_ptr = &(*array_ref).value.GetRef().GetVal()
		}
		if 1<<3 == 1<<0 {
			var __arr *ZendArray = ZendArrayDup(array_ptr.GetValue().GetArr())
			var __z *Zval = array_ptr
			__z.GetValue().SetArr(__arr)
			__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		} else {
			var _zv *Zval = array_ptr
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
		}
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(ZendHashIteratorAdd(array_ptr.GetValue().GetArr(), 0))
		if 1<<3 == 1<<2 {

		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else if 1<<3 != 1<<0 && array_ptr.GetType() == 8 {
		if array_ptr.GetValue().GetObj().GetCe().GetGetIterator() == nil {
			var properties *HashTable
			if 1<<3 == 1<<2 || 1<<3 == 1<<3 {
				if array_ptr == array_ref {
					var _ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
					ZendGcSetRefcount(&_ref.gc, 1)
					_ref.GetGc().SetTypeInfo(10)
					var _z1 *Zval = &_ref.val
					var _z2 *Zval = array_ref
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					_ref.GetSources().SetPtr(nil)
					array_ref.GetValue().SetRef(_ref)
					array_ref.SetTypeInfo(10 | 1<<0<<8)
					array_ptr = &(*array_ref).value.GetRef().GetVal()
				}
				ZvalAddrefP(array_ref)
				var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else {
				array_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				var _z1 *Zval = array_ptr
				var _z2 *Zval = array_ref
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			}
			if array_ptr.GetValue().GetObj().GetProperties() != nil && ZendGcRefcount(&(array_ptr.GetValue().GetObj().GetProperties()).gc) > 1 {
				if (ZvalGcFlags(array_ptr.GetValue().GetObj().GetProperties().GetGc().GetTypeInfo()) & 1 << 6) == 0 {
					ZendGcDelref(&(array_ptr.GetValue().GetObj().GetProperties()).gc)
				}
				array_ptr.GetValue().GetObj().SetProperties(ZendArrayDup(array_ptr.GetValue().GetObj().GetProperties()))
			}
			properties = array_ptr.GetValue().GetObj().GetHandlers().GetGetProperties()(&(*array_ptr))
			if properties.GetNNumOfElements() == 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
				if EG.GetException() != nil {
					return 0
				}
				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			}
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(ZendHashIteratorAdd(properties, 0))
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			var is_empty ZendBool = ZendFeResetIterator(array_ptr, 1, opline, execute_data)
			if 1<<3 == 1<<2 {

			}
			if EG.GetException() != nil {
				return 0
			} else if is_empty != 0 {

				execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
				if EG.GetVmInterrupt() != 0 {
					return zend_interrupt_helper_SPEC(execute_data)
				}
				return 0
			} else {
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
	} else {
		ZendError(1<<1, "Invalid argument supplied for foreach()")
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetFeIterIdx(uint32 - 1)
		if 1<<3 == 1<<2 {

		}
		if EG.GetException() != nil {
			return 0
		}
		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
}
func ZEND_JMP_SET_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var ref *Zval = nil
	var ret int
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	if (1<<3 == 1<<2 || 1<<3 == 1<<3) && value.GetType() == 10 {
		if 1<<3 == 1<<2 {
			ref = value
		}
		value = &(*value).value.GetRef().GetVal()
	}
	ret = IZendIsTrue(value)
	if EG.GetException() != nil {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
	if ret != 0 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<3 == 1<<3 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<3 == 1<<2 && ref != nil {
			var r *ZendReference = ref.GetValue().GetRef()
			if ZendGcDelref(&r.gc) == 0 {
				_efree(r)
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_COALESCE_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var ref *Zval = nil
	value = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if (1<<3&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
		if (1 << 3 & 1 << 2) != 0 {
			ref = value
		}
		value = &(*value).value.GetRef().GetVal()
	}
	if value.GetType() > 1 {
		var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if 1<<3 == 1<<3 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else if (1<<3&1<<2) != 0 && ref != nil {
			var r *ZendReference = ref.GetValue().GetRef()
			if ZendGcDelref(&r.gc) == 0 {
				_efree(r)
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}

		execute_data.SetOpline((*ZendOp)((*byte)(opline) + int(opline.GetOp2().GetJmpOffset())))
		if EG.GetVmInterrupt() != 0 {
			return zend_interrupt_helper_SPEC(execute_data)
		}
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_QM_ASSIGN_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<3 && value.GetType() == 0 {
		_zvalUndefinedOp1(execute_data)
		result.SetTypeInfo(1)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	if 1<<3 == 1<<3 {
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
		var _z1 *Zval = result
		var _z2 *Zval = _z3
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	} else if 1<<3 == 1<<2 {
		if value.GetType() == 10 {
			var _z1 *Zval = result
			var _z2 *Zval = &(*value).value.GetRef().GetVal()
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			if ZvalDelrefP(value) == 0 {
				_efree(value.GetValue().GetRef())
			} else if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		} else {
			var _z1 *Zval = result
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
		}
	} else {
		var _z1 *Zval = result
		var _z2 *Zval = value
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 == 1<<0 {
			if (result.GetTypeInfo() & 0xff00) != 0 {
				ZvalAddrefP(result)
			}
		}
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_YIELD_FROM_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var generator *ZendGenerator = ZendGetRunningGenerator(execute_data)
	var val *Zval
	val = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	if (generator.GetFlags() & ZEND_GENERATOR_FORCED_CLOSE) != 0 {
		ZendThrowError(nil, "Cannot use \"yield from\" in a force-closed generator")
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}
	if val.GetType() == 7 {
		var _z1 *Zval = &generator.values
		var _z2 *Zval = val
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if 1<<3 != 1<<1 && (val.GetTypeInfo()&0xff00) != 0 {
			ZvalAddrefP(val)
		}
		generator.GetValues().SetFePos(0)
	} else if 1<<3 != 1<<0 && val.GetType() == 8 && val.GetValue().GetObj().GetCe().GetGetIterator() != nil {
		var ce *ZendClassEntry = val.GetValue().GetObj().GetCe()
		if ce == ZendCeGenerator {
			var new_gen *ZendGenerator = (*ZendGenerator)(val.GetValue().GetObj())
			if 1<<3 != 1<<1 {
				ZvalAddrefP(val)
			}
			if new_gen.GetRetval().GetType() == 0 {
				if ZendGeneratorGetCurrent(new_gen) == generator {
					ZendThrowError(nil, "Impossible to yield from the Generator being currently run")
					ZvalPtrDtor(val)
					if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
						(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					}
					return 0
				} else {
					ZendGeneratorYieldFrom(generator, new_gen)
				}
			} else if new_gen.GetExecuteData() == nil {
				ZendThrowError(nil, "Generator passed to yield from was aborted without proper return and is unable to continue")
				ZvalPtrDtor(val)
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
				return 0
			} else {
				if opline.GetResultType() != 0 {
					var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
					var _z2 *Zval = &new_gen.retval
					var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
					var _t uint32 = _z2.GetTypeInfo()
					_z1.GetValue().SetCounted(_gc)
					_z1.SetTypeInfo(_t)
					if (_t & 0xff00) != 0 {
						ZendGcAddref(&_gc.gc)
					}
				}
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			var iter *ZendObjectIterator = ce.GetGetIterator()(ce, val, 0)
			if iter == nil || EG.GetException() != nil {
				if EG.GetException() == nil {
					ZendThrowError(nil, "Object of type %s did not create an Iterator", ce.GetName().GetVal())
				}
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
				return 0
			}
			iter.SetIndex(0)
			if iter.GetFuncs().GetRewind() != nil {
				iter.GetFuncs().GetRewind()(iter)
				if EG.GetException() != nil {
					ZendObjectRelease(&iter.std)
					if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
						(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
					}
					return 0
				}
			}
			var __z *Zval = &generator.values
			__z.GetValue().SetObj(&iter.std)
			__z.SetTypeInfo(8 | 1<<0<<8 | 1<<1<<8)
		}
	} else {
		ZendThrowError(nil, "Can use \"yield from\" only with arrays and Traversables")
		if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		}
		return 0
	}

	/* This is the default return value
	 * when the expression is a Generator, it will be overwritten in zend_generator_resume() */

	if opline.GetResultType() != 0 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	}

	/* This generator has no send target (though the generator we delegate to might have one) */

	generator.SetSendTarget(nil)

	/* We increment to the next op, so we are at the correct position when the
	 * generator is resumed. */

	execute_data.GetOpline()++

	/* The GOTO VM uses a local opline variable. We need to set the opline
	 * variable in execute_data so we don't resume at an old position. */

	return -1
}
func ZEND_STRLEN_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if value.GetType() == 6 {
		var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
		__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
		__z.SetTypeInfo(4)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		var strict ZendBool
		if (1<<3&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
			value = &(*value).value.GetRef().GetVal()
			if value.GetType() == 6 {
				var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
				__z.GetValue().SetLval(value.GetValue().GetStr().GetLen())
				__z.SetTypeInfo(4)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		}
		if 1<<3 == 1<<3 && value.GetType() == 0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_TYPE_CHECK_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int = 0
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
	type_check_resource:
		if value.GetType() != 9 || nil != ZendRsrcListGetRsrcType(value.GetValue().GetRes()) {
			result = 1
		}
	} else if (1<<3&(1<<3|1<<2)) != 0 && value.GetType() == 10 {
		value = &(*value).value.GetRef().GetVal()
		if (opline.GetExtendedValue() >> uint32(*value).u1.v.type_ & 1) != 0 {
			goto type_check_resource
		}
	} else if 1<<3 == 1<<3 && value.GetType() == 0 {
		result = (1 << 1 & opline.GetExtendedValue()) != 0
		_zvalUndefinedOp1(execute_data)
		if EG.GetException() != nil {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	}
	if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var_ptr.GetValue().GetLval()++

	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var_ptr.GetValue().GetLval()++
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	FastLongIncrementFunction(var_ptr)

	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	FastLongIncrementFunction(var_ptr)
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = var_ptr
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var_ptr.GetValue().GetLval()--

	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var_ptr.GetValue().GetLval()--
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	FastLongDecrementFunction(var_ptr)

	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	FastLongDecrementFunction(var_ptr)
	var _z1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	var _z2 *Zval = var_ptr
	var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
	var _t uint32 = _z2.GetTypeInfo()
	_z1.GetValue().SetCounted(_gc)
	_z1.SetTypeInfo(_t)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_POST_INC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	var_ptr.GetValue().GetLval()++
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_POST_INC_LONG_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	FastLongIncrementFunction(var_ptr)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_POST_DEC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	var_ptr.GetValue().GetLval()--
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_POST_DEC_LONG_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var __z *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	__z.GetValue().SetLval(var_ptr.GetValue().GetLval())
	__z.SetTypeInfo(4)
	FastLongDecrementFunction(var_ptr)
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAR_SIMPLE_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	varptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<3 {
		var _z1 *Zval = arg
		var _z2 *Zval = varptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	} else {
		var _z1 *Zval = arg
		var _z2 *Zval = varptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_SEND_VAR_EX_SIMPLE_SPEC_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varptr *Zval
	var arg *Zval
	var arg_num uint32 = opline.GetOp2().GetNum()
	if (execute_data.GetCall().GetFunc().GetQuickArgFlags() >> (arg_num + 3) * 2 & (1 | 2)) != 0 {
		return ZEND_SEND_REF_SPEC_CV_HANDLER(execute_data)
	}
	varptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	arg = (*Zval)((*byte)(execute_data.GetCall()) + int(opline.GetResult().GetVar()))
	if 1<<3 == 1<<3 {
		var _z1 *Zval = arg
		var _z2 *Zval = varptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
		if (_t & 0xff00) != 0 {
			ZendGcAddref(&_gc.gc)
		}
	} else {
		var _z1 *Zval = arg
		var _z2 *Zval = varptr
		var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
		var _t uint32 = _z2.GetTypeInfo()
		_z1.GetValue().SetCounted(_gc)
		_z1.SetTypeInfo(_t)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_DIV_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<3 == 1<<0 || op1.GetType() == 6) && (1<<0 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<0 != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<3 == 1<<0 || 1<<3 == 1<<3 {
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
		} else if 1<<3 != 1<<0 && 1<<3 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
		if 1<<3 == 1<<3 && op1.GetType() == 0 {
			op1 = _zvalUndefinedOp1(execute_data)
		}
		if 1<<0 == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_IS_IDENTICAL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = FastIsIdenticalFunction(op1, op2)
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
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = FastIsNotIdenticalFunction(op1, op2)
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
func ZEND_IS_EQUAL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() == op2.GetValue().GetLval() {
			is_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
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
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

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
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

	} else if op1.GetType() == 4 {
		if op2.GetType() == 4 {
			if op1.GetValue().GetLval() != op2.GetValue().GetLval() {
			is_not_equal_true:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(3)
				assert(EG.GetException() == nil)
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
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 && 1<<0 == 1<<0 {

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
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			} else {
			is_not_equal_false:
				execute_data.SetOpline(opline + 2)
				return 0
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(2)
				assert(EG.GetException() == nil)
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_SPACESHIP_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_XOR_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		value = _getOpDataZvalPtrR((opline + 1).GetOp1Type(), (opline + 1).GetOp1(), &free_op_data, execute_data, opline)
		if 1<<3 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto assign_op_object
			}
			if 1<<3 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
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

func ZEND_ASSIGN_DIM_OP_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data1 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	var container *Zval
	var dim *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		if 1<<0 == 0 {
			var_ptr = ZendHashNextIndexInsert(container.GetValue().GetArr(), &EG.uninitialized_zval)
			if var_ptr == nil {
				ZendCannotAddElement()
				goto assign_dim_op_ret_null
			}
		} else {
			if 1<<0 == 1<<0 {
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
			if 1<<0 != 0 && var_ptr.GetType() == 10 {
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
		dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		if container.GetType() == 8 {
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendBinaryAssignOpObjDim(container, dim, opline, execute_data)
		} else if container.GetType() <= 2 {
			if 1<<3 == 1<<3 && container.GetTypeInfo() == 0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OP_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var value *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	var_ptr = _get_zval_ptr_cv_BP_VAR_RW(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 1<<2 && var_ptr.GetType() == 15 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		if 1<<3 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto pre_incdec_object
			}
			if 1<<3 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
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
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		if 1<<3 != 0 && object.GetType() != 8 {
			if object.GetType() == 10 && &(*object).value.GetRef().GetVal().u1.v.type_ == 8 {
				object = &(*object).value.GetRef().GetVal()
				goto post_incdec_object
			}
			if 1<<3 == 1<<3 && object.GetType() == 0 {
				_zvalUndefinedOp1(execute_data)
			}
			object = MakeRealObject(object, property, opline, execute_data)
			if object == nil {
				break
			}
		}
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
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 != 1<<0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_W(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_FETCH_DIM_RW_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_RW(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_FETCH_DIM_IS_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_read_IS(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER(execute_data)
	} else {
		if 1<<0 == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER(execute_data)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_UNSET(container, (*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant), 1<<0, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 || 1<<3 != 0 && container.GetType() != 8 {
		for {
			if (1<<3&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() == 8 {
					break
				}
			}
			if 1<<3 == 1<<3 && container.GetType() == 0 {
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
					if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
						assert(EG.GetException() == nil)
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
							if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
					if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_INLINE_HANDLER(execute_data)
}
func ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_FETCH_OBJ_RW_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_FETCH_OBJ_IS_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 || 1<<3 != 0 && container.GetType() != 8 {
		for {
			if (1<<3&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
					if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
						assert(EG.GetException() == nil)
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
							if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
					if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<0, g.CondF1(1<<0 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
	if 1<<3 == 1<<2 {
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
func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
	if 1<<3 != 0 && object.GetType() != 8 {
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 != 0 && object.GetType() != 8 {
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _getZvalPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 != 0 && object.GetType() != 8 {
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value = _get_zval_ptr_cv_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<3 != 0 && object.GetType() != 8 {
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

func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<0 == 0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			if 1<<0 == 1<<0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<0 == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				value = (*Zval)((*byte)(opline+1) + int32((opline + 1).GetOp1()).constant)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<0 != 0 {

	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<0 == 0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			if 1<<0 == 1<<0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<0 == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				value = _getZvalPtrTmp((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<0 != 0 {

	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<0 == 0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			if 1<<0 == 1<<0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
			ZvalPtrDtorNogc(free_op_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<0 == 0 {
				ZendUseNewElementForString()
				ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				value = _getZvalPtrVarDeref((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
				ZvalPtrDtorNogc(free_op_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		assign_dim_error:
			ZvalPtrDtorNogc((*Zval)((*byte)(execute_data) + int((opline + 1).GetOp1().GetVar())))
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<0 != 0 {

	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op_data ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	var dim *Zval
	object_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<0 == 0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			if 1<<0 == 1<<0 {
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
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
			value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
			if 1<<0 == 1<<0 && dim.GetU2Extra() == 1 {
				dim++
			}
			ZendAssignToObjectDim(object_ptr, dim, value, opline, execute_data)
		} else if object_ptr.GetType() == 6 {
			if 1<<0 == 0 {
				ZendUseNewElementForString()
				if (opline.GetResultType() & (1<<2 | 1<<1)) != 0 {
					(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				}
			} else {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
				value = _get_zval_ptr_cv_deref_BP_VAR_R((opline + 1).GetOp1().GetVar(), execute_data)
				ZendAssignToStringOffset(object_ptr, dim, value, opline, execute_data)
			}
		} else if object_ptr.GetType() <= 2 {
			if orig_object_ptr.GetType() == 10 && orig_object_ptr.GetValue().GetRef().GetSources().GetPtr() != nil && ZendVerifyRefArrayAssignable(orig_object_ptr.GetValue().GetRef()) == 0 {
				dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}
	if 1<<0 != 0 {

	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var variable_ptr *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var variable_ptr *Zval
	value = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<0, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 == 0 {
		if 1<<0 == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if 1<<0 == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 {
		if 1<<0 == 1<<0 {
			ZendAssignToPropertyReferenceThisConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceThisVar(container, property, value_ptr, opline, execute_data)
		}
	} else {
		if 1<<0 == 1<<0 {
			ZendAssignToPropertyReferenceVarConst(container, property, value_ptr, opline, execute_data)
		} else {
			ZendAssignToPropertyReferenceVarVar(container, property, value_ptr, opline, execute_data)
		}
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_FAST_CONCAT_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if (1<<3 == 1<<0 || op1.GetType() == 6) && (1<<0 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<0 != 1<<0 && op2_str.GetLen() == 0 {
			if 1<<3 == 1<<0 || 1<<3 == 1<<3 {
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
		} else if 1<<3 != 1<<0 && 1<<3 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 0 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		}
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
	if 1<<3 == 1<<0 {
		op1_str = op1.GetValue().GetStr()
	} else if op1.GetType() == 6 {
		op1_str = ZendStringCopy(op1.GetValue().GetStr())
	} else {
		if 1<<3 == 1<<3 && op1.GetType() == 0 {
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
		if 1<<3 != 1<<0 {
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
				if 1<<3 == 1<<0 {
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
		if 1<<3 != 1<<0 {
			ZendStringReleaseEx(op1_str, 0)
		}
		if 1<<0 != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var function_name *Zval
	var free_op1 ZendFreeOp
	var object *Zval
	var fbc *ZendFunction
	var called_scope *ZendClassEntry
	var obj *ZendObject
	var call *ZendExecuteData
	var call_info uint32
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
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
	if 1<<3 != 0 {
		for {
			if 1<<3 == 1<<0 || object.GetType() != 8 {
				if (1<<3&(1<<2|1<<3)) != 0 && object.GetType() == 10 {
					object = &(*object).value.GetRef().GetVal()
					if object.GetType() == 8 {
						break
					}
				}
				if 1<<3 == 1<<3 && object.GetType() == 0 {
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
			return 0
		}
		if 1<<0 == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
			var slot *any = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetResult().GetNum()))
			slot[0] = called_scope
			slot[1] = fbc
		}
		if (1<<3&(1<<2|1<<1)) != 0 && obj != orig_obj {

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
		if (1<<3&(1<<2|1<<1)) != 0 && EG.GetException() != nil {
			return 0
		}

		/* call static method */

		obj = (*ZendObject)(called_scope)
		call_info = 0<<16 | 0<<17
	} else if (1 << 3 & (1<<2 | 1<<1 | 1<<3)) != 0 {
		if 1<<3 == 1<<3 {
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr_ptr *Zval
	var new_expr Zval
	if (1<<3 == 1<<2 || 1<<3 == 1<<3) && (opline.GetExtendedValue()&1<<0) != 0 {
		expr_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), execute_data)
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
		expr_ptr = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
		if 1<<3 == 1<<1 {

		} else if 1<<3 == 1<<0 {
			if expr_ptr.GetTypeFlags() != 0 {
				ZvalAddrefP(expr_ptr)
			}
		} else if 1<<3 == 1<<3 {
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
	if 1<<0 != 0 {
		var offset *Zval = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
		var str *ZendString
		var hval ZendUlong
	add_again:
		if offset.GetType() == 6 {
			str = offset.GetValue().GetStr()
			if 1<<0 != 1<<0 {
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
		} else if (1<<0&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
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
		} else if 1<<0 == 1<<3 && offset.GetType() == 0 {
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
func ZEND_INIT_ARRAY_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var array *Zval
	var size uint32
	var opline *ZendOp = execute_data.GetOpline()
	array = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	if 1<<3 != 0 {
		size = opline.GetExtendedValue() >> 2
		var __arr *ZendArray = _zendNewArray(size)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)

		/* Explicitly initialize array as not-packed if flag is set */

		if (opline.GetExtendedValue() & 1 << 1) != 0 {
			ZendHashRealInitMixed(array.GetValue().GetArr())
		}
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER(execute_data)
	} else {
		var __arr *ZendArray = _zendNewArray(0)
		var __z *Zval = array
		__z.GetValue().SetArr(__arr)
		__z.SetTypeInfo(7 | 1<<0<<8 | 1<<1<<8)
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_UNSET_DIM_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var hval ZendUlong
	var key *ZendString
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
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
				if 1<<0 != 1<<0 {
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
			} else if (1<<0&(1<<2|1<<3)) != 0 && offset.GetType() == 10 {
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
			} else if 1<<0 == 1<<3 && offset.GetType() == 0 {
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
		if 1<<3 == 1<<3 && container.GetType() == 0 {
			container = _zvalUndefinedOp1(execute_data)
		}
		if 1<<0 == 1<<3 && offset.GetType() == 0 {
			offset = _zvalUndefinedOp2(execute_data)
		}
		if container.GetType() == 8 {
			if 1<<0 == 1<<0 && offset.GetU2Extra() == 1 {
				offset++
			}
			container.GetValue().GetObj().GetHandlers().GetUnsetDimension()(container, offset)
		} else if 1<<3 != 0 && container.GetType() == 6 {
			ZendThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	for {
		if 1<<3 != 0 && container.GetType() != 8 {
			if container.GetType() == 10 {
				container = &(*container).value.GetRef().GetVal()
				if container.GetType() != 8 {
					if 1<<3 == 1<<3 && container.GetType() == 0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
			if (1 << 3 & (1<<0 | 1<<3)) != 0 {

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
				assert(EG.GetException() == nil)
				execute_data.SetOpline(opline + 1)
				return 0
			}
		} else {
			result = value == nil || IZendIsTrue(value) == 0
		}
		goto isset_dim_obj_exit
	} else if (1<<3&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant)
	if 1<<3 == 1<<0 || 1<<3 != 0 && container.GetType() != 8 {
		if (1<<3&(1<<2|1<<3)) != 0 && container.GetType() == 10 {
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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

func ZEND_INSTANCEOF_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr *Zval
	var result ZendBool
	expr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
				assert(EG.GetException() != nil)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
		} else {
			ce = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar())).GetValue().GetCe()
		}
		result = ce != nil && InstanceofFunction(expr.GetValue().GetObj().GetCe(), ce) != 0
	} else if (1<<3&(1<<2|1<<3)) != 0 && expr.GetType() == 10 {
		expr = &(*expr).value.GetRef().GetVal()
		goto try_instanceof
	} else {
		if 1<<3 == 1<<3 && expr.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		result = 0
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
func ZEND_YIELD_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
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

	if 1<<3 != 0 {
		if (execute_data.GetFunc().GetOpArray().GetFnFlags() & 1 << 12) != 0 {

			/* Constants and temporary variables aren't yieldable by reference,
			 * but we still allow them with a notice. */

			if (1 << 3 & (1<<0 | 1<<1)) != 0 {
				var value *Zval
				ZendError(1<<3, "Only variable references should be yielded by reference")
				value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if 1<<3 == 1<<0 {
					if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
						ZvalAddrefP(&(generator.GetValue()))
					}
				}
			} else {
				var value_ptr *Zval = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp1().GetVar(), execute_data)

				/* If a function call result is yielded and the function did
				 * not return by reference we throw a notice. */

				for {
					if 1<<3 == 1<<2 {
						assert(value_ptr != &EG.uninitialized_zval)
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
			var value *Zval = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)

			/* Consts, temporary variables and references need copying */

			if 1<<3 == 1<<0 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
				if (generator.GetValue().GetTypeInfo() & 0xff00) != 0 {
					ZvalAddrefP(&(generator.GetValue()))
				}
			} else if 1<<3 == 1<<1 {
				var _z1 *Zval = &generator.value
				var _z2 *Zval = value
				var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
				var _t uint32 = _z2.GetTypeInfo()
				_z1.GetValue().SetCounted(_gc)
				_z1.SetTypeInfo(_t)
			} else if (1<<3&(1<<2|1<<3)) != 0 && value.GetType() == 10 {
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
				if 1<<3 == 1<<3 {
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
func ZEND_BIND_GLOBAL_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *ZendString
	var value *Zval
	var variable_ptr *Zval
	var idx uintPtr
	var ref *ZendReference
	for {
		varname = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetStr()

		/* We store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

		idx = uintptr_t((*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())))[0] - 1
		if idx < EG.GetSymbolTable().GetNNumUsed()*g.SizeOf("Bucket") {
			var p *Bucket = (*Bucket)((*byte)(EG.GetSymbolTable().GetArData() + idx))
			if p.GetVal().GetType() != 0 && (p.GetKey() == varname || p.GetH() == varname.GetH() && p.GetKey() != nil && ZendStringEqualContent(p.GetKey(), varname) != 0) {
				value = (*Zval)(p)
				goto check_indirect
			}
		}
		value = ZendHashFindEx(&EG.symbol_table, varname, 1)
		if value == nil {
			value = ZendHashAddNew(&EG.symbol_table, varname, &EG.uninitialized_zval)
			idx = (*byte)(value - (*byte)(EG.GetSymbolTable().GetArData()))

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = any(idx + 1)

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

		} else {
			idx = (*byte)(value - (*byte)(EG.GetSymbolTable().GetArData()))

			/* Store "hash slot index" + 1 (NULL is a mark of uninitialized cache slot) */

			(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = any(idx + 1)
		check_indirect:

			/* GLOBAL variable may be an INDIRECT pointer to CV */

			if value.GetType() == 13 {
				value = value.GetValue().GetZv()
				if value.GetType() == 0 {
					value.SetTypeInfo(1)
				}
			}

			/* GLOBAL variable may be an INDIRECT pointer to CV */

		}
		if value.GetType() != 10 {
			var _z *Zval = value
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
			ref = value.GetValue().GetRef()
		} else {
			ref = value.GetValue().GetRef()
			ZendGcAddref(&ref.gc)
		}
		variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
		if variable_ptr.GetTypeFlags() != 0 {
			var ref *ZendRefcounted = variable_ptr.GetValue().GetCounted()
			var refcnt uint32 = ZendGcDelref(&ref.gc)
			if variable_ptr != value {
				if refcnt == 0 {
					RcDtorFunc(ref)
					if EG.GetException() != nil {
						variable_ptr.SetTypeInfo(1)
						return 0
					}
				} else {
					GcCheckPossibleRoot(ref)
				}
			}
		}
		var __z *Zval = variable_ptr
		__z.GetValue().SetRef(ref)
		__z.SetTypeInfo(10 | 1<<0<<8)
		if g.PreInc(&opline).opcode != 168 {
			break
		}
	}
	execute_data.SetOpline(opline)
	return 0
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var ht *HashTable = (*Zval)((*byte)(opline) + int32(opline.GetOp2()).constant).GetValue().GetArr()
	var result *Zval
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	if op1.GetType() == 6 {
		result = ZendHashFindEx(ht, op1.GetValue().GetStr(), 1<<3 == 1<<0)
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
func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	var offset ZendLong
	var ht *HashTable
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if (1 << 3 & (1<<1 | 1<<2)) != 0 {
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else if 1<<3 != 1<<0 && container.GetType() == 10 {
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
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
fetch_dim_r_index_undef:
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	ZendUndefinedOffset(offset)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	var offset ZendLong
	var ht *HashTable
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if (1 << 3 & (1<<1 | 1<<2)) != 0 {
			execute_data.SetOpline(execute_data.GetOpline() + 1)
			return 0
		} else {
			assert(EG.GetException() == nil)
			execute_data.SetOpline(opline + 1)
			return 0
		}
	} else if 1<<3 != 1<<0 && container.GetType() == 10 {
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
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
fetch_dim_r_index_undef:
	(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
	ZendUndefinedOffset(offset)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
