// <<generate>>

package zend

import g "sik/runtime/grammar"

func ZEND_DIV_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<3 == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
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
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<3 != 1<<0 && 1<<3 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
		if (1<<1|1<<2) == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		ZvalPtrDtorNogc(free_op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 == 1<<0 && (1<<1|1<<2) == 1<<0 {

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
func ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
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
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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

func ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
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
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var var_ptr *Zval
	var value *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var zptr *Zval
	var cache_slot *any
	var prop_info *ZendPropertyInfo
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	dim = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if 1<<3 != 1<<0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_W(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_RW(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_read_IS(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER(execute_data)
	} else {
		if (1<<1 | 1<<2) == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_UNSET(container, _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data), 1<<1|1<<2, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
func ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var free_op2 ZendFreeOp
	var offset *Zval
	var cache_slot *any = nil
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
	if (1<<1 | 1<<2) == 1<<0 {
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
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var free_op2 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<1|1<<2, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
	ZvalPtrDtorNogc(free_op2)
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
func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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

func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object_ptr *Zval
	var orig_object_ptr *Zval
	var free_op2 ZendFreeOp
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 == 0 {
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
	ZvalPtrDtorNogc(free_op2)
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 {
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
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	if (1<<3 == 1<<0 || op1.GetType() == 6) && ((1<<1|1<<2) == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if (1<<1|1<<2) != 1<<0 && op2_str.GetLen() == 0 {
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
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<3 != 1<<0 && 1<<3 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if ((1<<1 | 1<<2) & (1<<1 | 1<<2)) != 0 {
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
		if 1<<3 != 1<<0 {
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
		if (1<<1 | 1<<2) != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
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
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
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
			return 0
		}
		if (1<<1|1<<2) == 1<<0 && fbc.GetType() <= 2 && (fbc.GetFnFlags()&(1<<18|1<<19)) == 0 && obj == orig_obj {
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
	if (1<<1 | 1<<2) != 1<<0 {
		ZvalPtrDtorNogc(free_op2)
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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER(execute_data)
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
func ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var offset *Zval
	var hval ZendUlong
	var key *ZendString
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<3 == 1<<3 && container.GetType() == 0 {
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
		} else if 1<<3 != 0 && container.GetType() == 6 {
			ZendThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
		container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1((1<<1|1<<2) == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
		break
	}
	ZvalPtrDtorNogc(free_op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
			if (1 << 3 & (1<<0 | 1<<3)) != 0 {

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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var container *Zval
	var result int
	var offset *Zval
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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

func ZEND_IS_IDENTICAL_SPEC_CV_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsIdenticalFunction(op1, op2)
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
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_TMP_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsNotIdenticalFunction(op1, op2)
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
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<1, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrTmp(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
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
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_YIELD_SPEC_CV_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_IS_IDENTICAL_SPEC_CV_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsIdenticalFunction(op1, op2)
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
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _getZvalPtrVarDeref(opline.GetOp2().GetVar(), &free_op2, execute_data)
	result = FastIsNotIdenticalFunction(op1, op2)
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
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
		ZvalPtrDtorNogc(free_op2)

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<2, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var value *Zval
	var variable_ptr *Zval
	value = _getZvalPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
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
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_REF_SPEC_CV_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op2 ZendFreeOp
	var variable_ptr *Zval
	var value_ptr *Zval
	value_ptr = _getZvalPtrPtrVar(opline.GetOp2().GetVar(), &free_op2, execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<3 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INSTANCEOF_SPEC_CV_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr *Zval
	var result ZendBool
	expr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
try_instanceof:
	if expr.GetType() == 8 {
		var ce *ZendClassEntry
		if 1<<2 == 1<<0 {
			ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
			if ce == nil {
				ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1).GetValue().GetStr(), 0x80)
				if ce != nil {
					(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
				}
			}
		} else if 1<<2 == 0 {
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
func ZEND_YIELD_SPEC_CV_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_ASSIGN_DIM_OP_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
func zend_fetch_var_address_helper_SPEC_CV_UNUSED(type_ int, execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *Zval
	var retval *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else if varname.GetType() == 6 {
		name = varname.GetValue().GetStr()
		tmp_name = nil
	} else {
		if 1<<3 == 1<<3 && varname.GetType() == 0 {
			_zvalUndefinedOp1(execute_data)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	retval = ZendHashFindEx(target_symbol_table, name, 1<<3 == 1<<0)
	if retval == nil {
		if ZendStringEquals(name, ZendKnownStrings[ZEND_STR_THIS]) != 0 {
		fetch_this:
			ZendFetchThisVar(type_, opline, execute_data)
			if 1<<3 != 1<<0 {
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
	if 1<<3 != 1<<0 {
		ZendTmpStringRelease(tmp_name)
	}
	assert(retval != nil)
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
func ZEND_FETCH_R_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(0, execute_data)
}
func ZEND_FETCH_W_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(1, execute_data)
}
func ZEND_FETCH_RW_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(2, execute_data)
}
func ZEND_FETCH_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var fetch_type int = g.Cond((execute_data.GetCall().GetThis().GetTypeInfo()&1<<31) != 0, 1, 0)
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(fetch_type, execute_data)
}
func ZEND_FETCH_UNSET_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(5, execute_data)
}
func ZEND_FETCH_IS_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	return zend_fetch_var_address_helper_SPEC_CV_UNUSED(3, execute_data)
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_W(container, nil, 0, opline, execute_data)
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
func ZEND_FETCH_DIM_RW_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_RW(container, nil, 0, opline, execute_data)
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
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER(execute_data)
	} else {
		return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		return ZEND_NULL_HANDLER(execute_data)
	}
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
				ZendUseScalarAsArray()
			}
			dim = nil
		assign_dim_error:
			if opline.GetResultType() != 0 {
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(1)
			}
		}
	}

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_VERIFY_RETURN_TYPE_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<3 == 0 {
		ZendVerifyMissingReturnType(execute_data.GetFunc(), (*any)((*byte)(execute_data.GetRunTimeCache()+opline.GetOp2().GetNum())))
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
	if ZendHashNextIndexInsert((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())).GetValue().GetArr(), expr_ptr) == nil {
		ZendCannotAddElement()
		ZvalPtrDtorNogc(expr_ptr)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_ARRAY_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER(execute_data)
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
func ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if var_.GetTypeFlags() != 0 {
		var garbage *ZendRefcounted = var_.GetValue().GetCounted()
		var_.SetTypeInfo(0)
		if ZendGcDelref(&garbage.gc) == 0 {
			RcDtorFunc(garbage)
		} else {
			GcCheckPossibleRoot(garbage)
		}
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	} else {
		var_.SetTypeInfo(0)
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var varname *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else if varname.GetType() == 6 {
		name = varname.GetValue().GetStr()
		tmp_name = nil
	} else {
		if 1<<3 == 1<<3 && varname.GetType() == 0 {
			varname = _zvalUndefinedOp1(execute_data)
		}
		name = ZvalTryGetTmpString(varname, &tmp_name)
		if name == nil {
			return 0
		}
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	ZendHashDelInd(target_symbol_table, name)
	if 1<<3 != 1<<0 {
		ZendTmpStringRelease(tmp_name)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CLASS_FETCH|CONST|VAR) */

func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_SET_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if value.GetType() > 1 && (value.GetType() != 10 || &(*value).value.GetRef().GetVal().u1.v.type_ != 1) {
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
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	} else {
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
		assert(EG.GetException() == nil)
		execute_data.SetOpline(opline + 1)
		return 0
	}
}
func ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_EMPTY_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	value = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	var result int
	result = !(IZendIsTrue(value))
	if EG.GetException() != nil {
		(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
		return 0
	}
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
func ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var result int
	var varname *Zval
	var name *ZendString
	var tmp_name *ZendString
	var target_symbol_table *HashTable
	varname = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 1<<0 {
		name = varname.GetValue().GetStr()
	} else {
		name = ZvalGetTmpString(varname, &tmp_name)
	}
	target_symbol_table = ZendGetTargetSymbolTable(opline.GetExtendedValue(), execute_data)
	value = ZendHashFindEx(target_symbol_table, name, 1<<3 == 1<<0)
	if 1<<3 != 1<<0 {
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

func ZEND_INSTANCEOF_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var expr *Zval
	var result ZendBool
	expr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
try_instanceof:
	if expr.GetType() == 8 {
		var ce *ZendClassEntry
		if 0 == 1<<0 {
			ce = (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0]
			if ce == nil {
				ce = ZendFetchClassByName((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant).GetValue().GetStr(), ((*Zval)((*byte)(opline)+int32(opline.GetOp2()).constant) + 1).GetValue().GetStr(), 0x80)
				if ce != nil {
					(*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue()))[0] = ce
				}
			}
		} else {
			ce = ZendFetchClass(nil, opline.GetOp2().GetNum())
			if ce == nil {
				assert(EG.GetException() != nil)
				(*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar())).SetTypeInfo(0)
				return 0
			}
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
func ZEND_YIELD_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var ht *HashTable
	var value *Zval
	var variable_ptr *Zval
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	IZvalPtrDtor(variable_ptr)
	if (uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr & 1) != 0 {
		ht = *((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr - 1)))
	} else {
		ht = any(*(execute_data.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()))
	}
	if ht == nil {
		assert((execute_data.GetFunc().GetOpArray().GetFnFlags() & (1<<7 | 1<<10)) != 0)
		ht = ZendArrayDup(execute_data.GetFunc().GetOpArray().GetStaticVariables())
		if (uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr & 1) != 0 {
			*((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr - 1))) = ht
		} else {
			*(execute_data.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()) = ht
		}
	} else if ZendGcRefcount(&ht.gc) > 1 {
		if (ZvalGcFlags(ht.GetGc().GetTypeInfo()) & 1 << 6) == 0 {
			ZendGcDelref(&ht.gc)
		}
		ht = ZendArrayDup(ht)
		if (uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr & 1) != 0 {
			*((*any)((*byte)(CG.GetMapPtrBase() + uintptr_t(execute_data.GetFunc()).op_array.static_variables_ptr__ptr - 1))) = ht
		} else {
			*(execute_data.GetFunc().GetOpArray().GetStaticVariablesPtrPtr()) = ht
		}
	}
	value = (*Zval)((*byte)(ht.GetArData() + (opline.GetExtendedValue() & ^(1 | 2))))
	if (opline.GetExtendedValue() & 1) != 0 {
		if value.GetType() == 11 {
			if ZvalUpdateConstantEx(value, execute_data.GetFunc().GetOpArray().GetScope()) != SUCCESS {
				variable_ptr.SetTypeInfo(1)
				return 0
			}
		}
		if value.GetType() != 10 {
			var ref *ZendReference = (*ZendReference)(_emalloc(g.SizeOf("zend_reference")))
			ZendGcSetRefcount(&ref.gc, 2)
			ref.GetGc().SetTypeInfo(10)
			var _z1 *Zval = &ref.val
			var _z2 *Zval = value
			var _gc *ZendRefcounted = _z2.GetValue().GetCounted()
			var _t uint32 = _z2.GetTypeInfo()
			_z1.GetValue().SetCounted(_gc)
			_z1.SetTypeInfo(_t)
			ref.GetSources().SetPtr(nil)
			value.GetValue().SetRef(ref)
			value.SetTypeInfo(10 | 1<<0<<8)
			var __z *Zval = variable_ptr
			__z.GetValue().SetRef(ref)
			__z.SetTypeInfo(10 | 1<<0<<8)
		} else {
			ZvalAddrefP(value)
			var __z *Zval = variable_ptr
			__z.GetValue().SetRef(value.GetValue().GetRef())
			__z.SetTypeInfo(10 | 1<<0<<8)
		}
	} else {
		var _z1 *Zval = variable_ptr
		var _z2 *Zval = value
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
func ZEND_CHECK_VAR_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if op1.GetTypeInfo() == 0 {
		_zvalUndefinedOp1(execute_data)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<3 {
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
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ZEND_COUNT_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var count ZendLong
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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

		} else if (1<<3&(1<<2|1<<3)) != 0 && op1.GetType() == 10 {
			op1 = &(*op1).value.GetRef().GetVal()
			continue
		} else if op1.GetType() <= 1 {
			if 1<<3 == 1<<3 && op1.GetType() == 0 {
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
func ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	if 1<<3 == 0 {
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
	} else {
		var op1 *Zval
		op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
			} else if (1<<3&(1<<2|1<<3)) != 0 && op1.GetType() == 10 {
				op1 = &(*op1).value.GetRef().GetVal()
				continue
			} else {
				if 1<<3 == 1<<3 && op1.GetType() == 0 {
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
func ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var type_ *ZendString
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
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
func ZEND_DIV_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	FastDivFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POW_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	PowFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_CONCAT_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<3 == 1<<0 || op1.GetType() == 6) && (1<<3 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<3 != 1<<0 && op2_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
		if 1<<3 == 1<<3 && op2.GetType() == 0 {
			op2 = _zvalUndefinedOp2(execute_data)
		}
		ConcatFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
		execute_data.SetOpline(execute_data.GetOpline() + 1)
		return 0
	}
}
func ZEND_IS_IDENTICAL_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
func ZEND_IS_NOT_IDENTICAL_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var result ZendBool
	op1 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_deref_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
func ZEND_IS_EQUAL_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_EQUAL_SPEC_CV_CV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var d1 float64
	var d2 float64
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 == 1<<0 && 1<<3 == 1<<0 {

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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
func ZEND_SPACESHIP_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	CompareFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_BOOL_XOR_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	op1 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp1().GetVar(), execute_data)
	op2 = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	BooleanXorFunction((*Zval)((*byte)(execute_data)+int(opline.GetResult().GetVar())), op1, op2)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_OP_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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

func ZEND_ASSIGN_DIM_OP_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_ASSIGN_OP_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var var_ptr *Zval
	var value *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
func ZEND_PRE_INC_OBJ_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_POST_INC_OBJ_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var dim *Zval
	var value *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	dim = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if 1<<3 != 1<<0 {
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
func ZEND_FETCH_DIM_W_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_W(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
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
func ZEND_FETCH_DIM_RW_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_RW(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
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
func ZEND_FETCH_DIM_IS_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_read_IS(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {
		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_W_SPEC_CV_CV_HANDLER(execute_data)
	} else {
		if 1<<3 == 0 {
			return zend_use_undef_in_read_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_DIM_UNSET_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	zend_fetch_dimension_address_UNSET(container, (*Zval)((*byte)(execute_data)+int(opline.GetOp2().GetVar())), 1<<3, opline, execute_data)
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
func ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
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
func ZEND_FETCH_OBJ_W_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + (opline.GetExtendedValue() & ^3))) }, nil), 1, opline.GetExtendedValue()&3, 1, opline, execute_data)
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
func ZEND_FETCH_OBJ_RW_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var property *Zval
	var container *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 2, 0, 1, opline, execute_data)
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
func ZEND_FETCH_OBJ_IS_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var cache_slot *any = nil
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
	if 1<<3 == 1<<0 {
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
func ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	if (execute_data.GetCall().GetThis().GetTypeInfo() & 1 << 31) != 0 {

		/* Behave like FETCH_OBJ_W */

		if (1 << 3 & (1<<0 | 1<<1)) != 0 {
			return zend_use_tmp_in_write_context_helper_SPEC(execute_data)
		}
		return ZEND_FETCH_OBJ_W_SPEC_CV_CV_HANDLER(execute_data)
	} else {
		return ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER(execute_data)
	}
}
func ZEND_FETCH_OBJ_UNSET_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op1 ZendFreeOp
	var container *Zval
	var property *Zval
	var result *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	result = (*Zval)((*byte)(execute_data) + int(opline.GetResult().GetVar()))
	ZendFetchPropertyAddress(result, container, 1<<3, property, 1<<3, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil), 5, 0, 1, opline, execute_data)
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
func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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

func ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var object *Zval
	var property *Zval
	var value *Zval
	var tmp Zval
	object = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && object.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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

func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CONST_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_TMP_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
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
			if 1<<3 != 1<<2 || object_ptr.GetType() != 15 {
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

	/* assign_dim has two opcodes! */

	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_CV_RETVAL_UNUSED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var variable_ptr *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {

	} else {
		value = ZendAssignToVariable(variable_ptr, value, 1<<3, (execute_data.GetFunc().GetFnFlags()&1<<31) != 0)

	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_SPEC_CV_CV_RETVAL_USED_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var value *Zval
	var variable_ptr *Zval
	value = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
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
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_REF_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var variable_ptr *Zval
	var value_ptr *Zval
	value_ptr = _get_zval_ptr_cv_BP_VAR_W(opline.GetOp2().GetVar(), execute_data)
	variable_ptr = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 1<<2 && variable_ptr.GetType() == 15 {
		variable_ptr = &EG.uninitialized_zval
	} else if 1<<3 == 1<<2 && (*Zval)((*byte)(execute_data)+int(opline.GetOp1().GetVar())).GetType() != 13 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_VAR_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var free_op_data ZendFreeOp
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _getZvalPtrPtrVar((opline + 1).GetOp1().GetVar(), &free_op_data, execute_data)
	if 1<<3 == 0 {
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
	if free_op_data != nil {
		ZvalPtrDtorNogc(free_op_data)
	}
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var property *Zval
	var container *Zval
	var value_ptr *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	property = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
	value_ptr = _get_zval_ptr_cv_BP_VAR_W((opline + 1).GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 {
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
	execute_data.SetOpline(execute_data.GetOpline() + 2)
	return 0
}

/* No specialization for op_types (CONST|TMPVAR|CV, UNUSED|CONST|VAR) */

func ZEND_FAST_CONCAT_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var op1 *Zval
	var op2 *Zval
	var op1_str *ZendString
	var op2_str *ZendString
	var str *ZendString
	op1 = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	op2 = (*Zval)((*byte)(execute_data) + int(opline.GetOp2().GetVar()))
	if (1<<3 == 1<<0 || op1.GetType() == 6) && (1<<3 == 1<<0 || op2.GetType() == 6) {
		var op1_str *ZendString = op1.GetValue().GetStr()
		var op2_str *ZendString = op2.GetValue().GetStr()
		var str *ZendString
		if 1<<3 != 1<<0 && op1_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
		} else if 1<<3 != 1<<0 && op2_str.GetLen() == 0 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op2_str, 0)
			}
		} else if 1<<3 != 1<<0 && 1<<3 != 1<<3 && (ZvalGcFlags(op1_str.GetGc().GetTypeInfo())&1<<6) == 0 && ZendGcRefcount(&op1_str.gc) == 1 {
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
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
				ZendStringReleaseEx(op1_str, 0)
			}
			if (1 << 3 & (1<<1 | 1<<2)) != 0 {
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
		if 1<<3 != 1<<0 {
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
		if 1<<3 != 1<<0 {
			ZendStringReleaseEx(op2_str, 0)
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_INIT_METHOD_CALL_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
	if 1<<3 != 1<<0 {

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
func ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_INIT_ARRAY_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
		return ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER(execute_data)
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
func ZEND_UNSET_DIM_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	var hval ZendUlong
	var key *ZendString
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
		if 1<<3 == 1<<3 && container.GetType() == 0 {
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
		} else if 1<<3 != 0 && container.GetType() == 6 {
			ZendThrowError(nil, "Cannot unset string offsets")
		}
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_UNSET_OBJ_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
		container.GetValue().GetObj().GetHandlers().GetUnsetProperty()(container, offset, g.CondF1(1<<3 == 1<<0, func() *any { return (*any)((*byte)(execute_data.GetRunTimeCache() + opline.GetExtendedValue())) }, nil))
		break
	}
	execute_data.SetOpline(execute_data.GetOpline() + 1)
	return 0
}
func ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var hval ZendUlong
	var offset *Zval
	container = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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
func ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var container *Zval
	var result int
	var offset *Zval
	container = _get_zval_ptr_cv_BP_VAR_IS(opline.GetOp1().GetVar(), execute_data)
	if 1<<3 == 0 && container.GetType() == 0 {
		return zend_this_not_in_object_context_helper_SPEC(execute_data)
	}
	offset = _get_zval_ptr_cv_BP_VAR_R(opline.GetOp2().GetVar(), execute_data)
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
func ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	var key *Zval
	var subject *Zval
	var ht *HashTable
	var result uint32
	key = (*Zval)((*byte)(execute_data) + int(opline.GetOp1().GetVar()))
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

func ZEND_YIELD_SPEC_CV_CV_HANDLER(execute_data *ZendExecuteData) int {
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
func ZEND_NULL_HANDLER(execute_data *ZendExecuteData) int {
	var opline *ZendOp = execute_data.GetOpline()
	ZendErrorNoreturn(1<<0, "Invalid opcode %d/%d/%d.", execute_data.GetOpline().GetOpcode(), execute_data.GetOpline().GetOp1Type(), execute_data.GetOpline().GetOp2Type())
	assert(EG.GetException() == nil)
	execute_data.SetOpline(opline + 1)
	return 0
}
func ExecuteEx(ex *ZendExecuteData) {
	var execute_data *ZendExecuteData = ex
	if EG.GetVmInterrupt() != 0 {
		zend_interrupt_helper_SPEC(execute_data)
	}
	for true {
		var ret int
		if g.Assign(&ret, opcode_handler_t(execute_data.GetOpline()).handler(execute_data)) != 0 {
			if ret > 0 {
				execute_data = EG.GetCurrentExecuteData()
				if EG.GetVmInterrupt() != 0 {
					zend_interrupt_helper_SPEC(execute_data)
				}
			} else {
				return
			}
		}
	}
	ZendErrorNoreturn(1<<4, "Arrived at end of main loop which shouldn't happen")
}
func ZendExecute(op_array *ZendOpArray, return_value *Zval) {
	var execute_data *ZendExecuteData
	var object_or_called_scope any
	var call_info uint32
	if EG.GetException() != nil {
		return
	}
	object_or_called_scope = ZendGetThisObject(EG.GetCurrentExecuteData())
	if !object_or_called_scope {
		object_or_called_scope = ZendGetCalledScope(EG.GetCurrentExecuteData())
		call_info = 1<<16 | 1<<17 | 1<<20
	} else {
		call_info = 1<<16 | 1<<17 | 1<<20 | (8 | 1<<0<<8 | 1<<1<<8)
	}
	execute_data = ZendVmStackPushCallFrame(call_info, (*ZendFunction)(op_array), 0, object_or_called_scope)
	if EG.GetCurrentExecuteData() != nil {
		execute_data.SetSymbolTable(ZendRebuildSymbolTable())
	} else {
		execute_data.SetSymbolTable(&EG.symbol_table)
	}
	execute_data.SetPrevExecuteData(EG.GetCurrentExecuteData())
	IInitCodeExecuteData(execute_data, op_array, return_value)
	ZendExecuteEx(execute_data)
	ZendVmStackFreeCallFrame(execute_data)
}
func ZendVmInit() {
	var labels []any = []any{ZEND_NOP_SPEC_HANDLER, ZEND_ADD_SPEC_CONST_CONST_HANDLER, ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_SPEC_CONST_CONST_HANDLER, ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_DIV_SPEC_CONST_CONST_HANDLER, ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER, ZEND_DIV_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_DIV_SPEC_CONST_CV_HANDLER, ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER, ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_DIV_SPEC_TMPVAR_CV_HANDLER, ZEND_DIV_SPEC_TMPVAR_CONST_HANDLER, ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_DIV_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_DIV_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_DIV_SPEC_CV_CONST_HANDLER, ZEND_DIV_SPEC_CV_TMPVAR_HANDLER, ZEND_DIV_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_DIV_SPEC_CV_CV_HANDLER, ZEND_MOD_SPEC_CONST_CONST_HANDLER, ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MOD_SPEC_CONST_TMPVARCV_HANDLER, ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MOD_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MOD_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SL_SPEC_CONST_CONST_HANDLER, ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SR_SPEC_CONST_CONST_HANDLER, ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SR_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER, ZEND_CONCAT_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_CONST_CV_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_CV_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_CONST_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_CV_CONST_HANDLER, ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER, ZEND_CONCAT_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CONCAT_SPEC_CV_CV_HANDLER, ZEND_BW_OR_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_OR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_AND_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_AND_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_XOR_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_CONST_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_XOR_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_POW_SPEC_CONST_CONST_HANDLER, ZEND_POW_SPEC_CONST_TMPVAR_HANDLER, ZEND_POW_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POW_SPEC_CONST_CV_HANDLER, ZEND_POW_SPEC_TMPVAR_CONST_HANDLER, ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POW_SPEC_TMPVAR_CV_HANDLER, ZEND_POW_SPEC_TMPVAR_CONST_HANDLER, ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_POW_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POW_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_POW_SPEC_CV_CONST_HANDLER, ZEND_POW_SPEC_CV_TMPVAR_HANDLER, ZEND_POW_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POW_SPEC_CV_CV_HANDLER, ZEND_BW_NOT_SPEC_CONST_HANDLER, ZEND_BW_NOT_SPEC_TMPVAR_HANDLER, ZEND_BW_NOT_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_BW_NOT_SPEC_CV_HANDLER, ZEND_BOOL_NOT_SPEC_CONST_HANDLER, ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER, ZEND_BOOL_NOT_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_NOT_SPEC_CV_HANDLER, ZEND_BOOL_XOR_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_CONST_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_BOOL_XOR_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_XOR_SPEC_CV_CONST_HANDLER, ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER, ZEND_BOOL_XOR_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_XOR_SPEC_CV_CV_HANDLER, ZEND_IS_IDENTICAL_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_IDENTICAL_SPEC_TMP_CONST_HANDLER, ZEND_IS_IDENTICAL_SPEC_TMP_TMP_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_IDENTICAL_SPEC_VAR_CONST_HANDLER, ZEND_IS_IDENTICAL_SPEC_VAR_TMP_HANDLER, ZEND_IS_IDENTICAL_SPEC_VAR_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_IDENTICAL_SPEC_CV_CONST_HANDLER, ZEND_IS_IDENTICAL_SPEC_CV_TMP_HANDLER, ZEND_IS_IDENTICAL_SPEC_CV_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_IDENTICAL_SPEC_CV_CV_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_TMP_CONST_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_TMP_TMP_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_VAR_CONST_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_VAR_TMP_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_VAR_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_CV_CONST_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_CV_TMP_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_CV_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_IDENTICAL_SPEC_CV_CV_HANDLER, ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CONST_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CV_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CV_JMPZ_HANDLER, ZEND_IS_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_TMPVAR_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_TMPVAR_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CV_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_SPEC_CV_CV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_VAR_CONST_RETVAL_USED_HANDLER, ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_VAR_TMP_RETVAL_USED_HANDLER, ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_VAR_VAR_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_VAR_CV_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_CV_CONST_RETVAL_USED_HANDLER, ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_CV_TMP_RETVAL_USED_HANDLER, ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_CV_VAR_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_SPEC_CV_CV_RETVAL_UNUSED_HANDLER, ZEND_ASSIGN_SPEC_CV_CV_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CONST_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_UNUSED_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_VAR_CV_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CONST_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_UNUSED_OP_DATA_CV_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_SPEC_CV_CV_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CONST_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_VAR_CV_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CONST_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_SPEC_CV_CV_OP_DATA_CV_HANDLER, ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CONST_HANDLER, ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_TMP_HANDLER, ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_STATIC_PROP_SPEC_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OP_SPEC_VAR_CONST_HANDLER, ZEND_ASSIGN_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_ASSIGN_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OP_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OP_SPEC_CV_CONST_HANDLER, ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_ASSIGN_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OP_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_VAR_CONST_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_VAR_UNUSED_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_CV_CONST_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_CV_UNUSED_HANDLER, ZEND_ASSIGN_DIM_OP_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CONST_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_VAR_CV_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CONST_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_UNUSED_CV_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_CV_CONST_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_OP_SPEC_CV_CV_HANDLER, ZEND_ASSIGN_STATIC_PROP_OP_SPEC_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_REF_SPEC_VAR_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_REF_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_REF_SPEC_CV_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_REF_SPEC_CV_CV_HANDLER, ZEND_QM_ASSIGN_SPEC_CONST_HANDLER, ZEND_QM_ASSIGN_SPEC_TMP_HANDLER, ZEND_QM_ASSIGN_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_QM_ASSIGN_SPEC_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CONST_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_VAR_CV_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CONST_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_UNUSED_CV_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_CONST_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_TMPVAR_OP_DATA_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ASSIGN_OBJ_REF_SPEC_CV_CV_OP_DATA_CV_HANDLER, ZEND_ASSIGN_STATIC_PROP_REF_SPEC_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_SPEC_VAR_RETVAL_UNUSED_HANDLER, ZEND_PRE_INC_SPEC_VAR_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_INC_SPEC_CV_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_DEC_SPEC_VAR_RETVAL_UNUSED_HANDLER, ZEND_PRE_DEC_SPEC_VAR_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_DEC_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_DEC_SPEC_CV_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_SPEC_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_DEC_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_DEC_SPEC_CV_HANDLER, ZEND_PRE_INC_STATIC_PROP_SPEC_HANDLER, ZEND_POST_INC_STATIC_PROP_SPEC_HANDLER, ZEND_JMP_SPEC_HANDLER, ZEND_JMPZ_SPEC_CONST_HANDLER, ZEND_JMPZ_SPEC_TMPVAR_HANDLER, ZEND_JMPZ_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMPZ_SPEC_CV_HANDLER, ZEND_JMPNZ_SPEC_CONST_HANDLER, ZEND_JMPNZ_SPEC_TMPVAR_HANDLER, ZEND_JMPNZ_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMPNZ_SPEC_CV_HANDLER, ZEND_JMPZNZ_SPEC_CONST_HANDLER, ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER, ZEND_JMPZNZ_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMPZNZ_SPEC_CV_HANDLER, ZEND_JMPZ_EX_SPEC_CONST_HANDLER, ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER, ZEND_JMPZ_EX_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMPZ_EX_SPEC_CV_HANDLER, ZEND_JMPNZ_EX_SPEC_CONST_HANDLER, ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER, ZEND_JMPNZ_EX_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMPNZ_EX_SPEC_CV_HANDLER, ZEND_CASE_SPEC_TMPVAR_CONST_HANDLER, ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_CASE_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CASE_SPEC_TMPVAR_CV_HANDLER, ZEND_CHECK_VAR_SPEC_CV_UNUSED_HANDLER, ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_HANDLER, ZEND_SEND_VAR_NO_REF_EX_SPEC_VAR_QUICK_HANDLER, ZEND_CAST_SPEC_CONST_HANDLER, ZEND_CAST_SPEC_TMP_HANDLER, ZEND_CAST_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_CAST_SPEC_CV_HANDLER, ZEND_BOOL_SPEC_CONST_HANDLER, ZEND_BOOL_SPEC_TMPVAR_HANDLER, ZEND_BOOL_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_BOOL_SPEC_CV_HANDLER, ZEND_FAST_CONCAT_SPEC_CONST_CONST_HANDLER, ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER, ZEND_FAST_CONCAT_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FAST_CONCAT_SPEC_CONST_CV_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_CONST_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FAST_CONCAT_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FAST_CONCAT_SPEC_CV_CONST_HANDLER, ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER, ZEND_FAST_CONCAT_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FAST_CONCAT_SPEC_CV_CV_HANDLER, ZEND_ROPE_INIT_SPEC_UNUSED_CONST_HANDLER, ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_ROPE_INIT_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ROPE_INIT_SPEC_UNUSED_CV_HANDLER, ZEND_ROPE_ADD_SPEC_TMP_CONST_HANDLER, ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER, ZEND_ROPE_ADD_SPEC_TMP_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ROPE_ADD_SPEC_TMP_CV_HANDLER, ZEND_ROPE_END_SPEC_TMP_CONST_HANDLER, ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER, ZEND_ROPE_END_SPEC_TMP_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ROPE_END_SPEC_TMP_CV_HANDLER, ZEND_BEGIN_SILENCE_SPEC_HANDLER, ZEND_END_SILENCE_SPEC_TMP_HANDLER, ZEND_INIT_FCALL_BY_NAME_SPEC_CONST_HANDLER, ZEND_DO_FCALL_SPEC_RETVAL_UNUSED_HANDLER, ZEND_DO_FCALL_SPEC_RETVAL_USED_HANDLER, ZEND_INIT_FCALL_SPEC_CONST_HANDLER, ZEND_RETURN_SPEC_CONST_HANDLER, ZEND_RETURN_SPEC_TMP_HANDLER, ZEND_RETURN_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_RETURN_SPEC_CV_HANDLER, ZEND_RECV_SPEC_UNUSED_HANDLER, ZEND_RECV_INIT_SPEC_CONST_HANDLER, ZEND_SEND_VAL_SPEC_CONST_HANDLER, ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER, ZEND_SEND_VAL_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_EX_SPEC_VAR_HANDLER, ZEND_SEND_VAR_EX_SPEC_VAR_QUICK_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_EX_SPEC_CV_HANDLER, ZEND_SEND_VAR_EX_SPEC_CV_QUICK_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_REF_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_REF_SPEC_CV_HANDLER, ZEND_NEW_SPEC_CONST_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NEW_SPEC_VAR_UNUSED_HANDLER, ZEND_NEW_SPEC_UNUSED_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_NS_FCALL_BY_NAME_SPEC_CONST_HANDLER, ZEND_FREE_SPEC_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_CONST_CONST_HANDLER, ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_CONST_UNUSED_HANDLER, ZEND_INIT_ARRAY_SPEC_CONST_CV_HANDLER, ZEND_INIT_ARRAY_SPEC_TMP_CONST_HANDLER, ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_TMP_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_TMP_UNUSED_HANDLER, ZEND_INIT_ARRAY_SPEC_TMP_CV_HANDLER, ZEND_INIT_ARRAY_SPEC_VAR_CONST_HANDLER, ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_VAR_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_VAR_UNUSED_HANDLER, ZEND_INIT_ARRAY_SPEC_VAR_CV_HANDLER, ZEND_INIT_ARRAY_SPEC_UNUSED_CONST_HANDLER, ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_UNUSED_UNUSED_HANDLER, ZEND_INIT_ARRAY_SPEC_UNUSED_CV_HANDLER, ZEND_INIT_ARRAY_SPEC_CV_CONST_HANDLER, ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_CV_TMPVAR_HANDLER, ZEND_INIT_ARRAY_SPEC_CV_UNUSED_HANDLER, ZEND_INIT_ARRAY_SPEC_CV_CV_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CONST_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_UNUSED_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CONST_CV_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CONST_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_UNUSED_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_TMP_CV_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CONST_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_UNUSED_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CONST_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_TMPVAR_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_UNUSED_HANDLER, ZEND_ADD_ARRAY_ELEMENT_SPEC_CV_CV_HANDLER, ZEND_INCLUDE_OR_EVAL_SPEC_CONST_HANDLER, ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER, ZEND_INCLUDE_OR_EVAL_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INCLUDE_OR_EVAL_SPEC_CV_HANDLER, ZEND_UNSET_VAR_SPEC_CONST_UNUSED_HANDLER, ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_UNSET_VAR_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_VAR_SPEC_CV_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_DIM_SPEC_VAR_CONST_HANDLER, ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER, ZEND_UNSET_DIM_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_DIM_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_DIM_SPEC_CV_CONST_HANDLER, ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER, ZEND_UNSET_DIM_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_DIM_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_OBJ_SPEC_VAR_CONST_HANDLER, ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_UNSET_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_OBJ_SPEC_VAR_CV_HANDLER, ZEND_UNSET_OBJ_SPEC_UNUSED_CONST_HANDLER, ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_UNSET_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_OBJ_SPEC_UNUSED_CV_HANDLER, ZEND_UNSET_OBJ_SPEC_CV_CONST_HANDLER, ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_UNSET_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_UNSET_OBJ_SPEC_CV_CV_HANDLER, ZEND_FE_RESET_R_SPEC_CONST_HANDLER, ZEND_FE_RESET_R_SPEC_TMP_HANDLER, ZEND_FE_RESET_R_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FE_RESET_R_SPEC_CV_HANDLER, ZEND_FE_FETCH_R_SPEC_VAR_HANDLER, ZEND_EXIT_SPEC_HANDLER, ZEND_FETCH_R_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_R_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_R_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_R_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_R_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_DIM_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_SPEC_CONST_CV_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_R_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_SPEC_CV_CV_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CONST_CV_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_R_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_OBJ_R_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_R_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_R_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_R_SPEC_CV_CV_HANDLER, ZEND_FETCH_W_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_W_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_W_SPEC_CV_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_W_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_W_SPEC_VAR_UNUSED_HANDLER, ZEND_FETCH_DIM_W_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_W_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_W_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_W_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_W_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_W_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_W_SPEC_VAR_CV_HANDLER, ZEND_FETCH_OBJ_W_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_W_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_W_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_W_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_W_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_W_SPEC_CV_CV_HANDLER, ZEND_FETCH_RW_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_RW_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_RW_SPEC_CV_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_RW_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_RW_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_RW_SPEC_VAR_UNUSED_HANDLER, ZEND_FETCH_DIM_RW_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_RW_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_RW_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_RW_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_RW_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_VAR_CV_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_RW_SPEC_CV_CV_HANDLER, ZEND_FETCH_IS_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_IS_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_IS_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_IS_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CONST_CV_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_IS_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_IS_SPEC_CV_CV_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CONST_CV_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_TMPVAR_CV_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_IS_SPEC_CV_CV_HANDLER, ZEND_FETCH_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_FUNC_ARG_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_FUNC_ARG_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CONST_CV_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CONST_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_UNUSED_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_TMP_CV_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_UNUSED_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_DIM_FUNC_ARG_SPEC_CV_CV_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CONST_CV_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CONST_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_TMP_CV_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_VAR_CV_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_FUNC_ARG_SPEC_CV_CV_HANDLER, ZEND_FETCH_UNSET_SPEC_CONST_UNUSED_HANDLER, ZEND_FETCH_UNSET_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_FETCH_UNSET_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_UNSET_SPEC_CV_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_VAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_UNSET_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_VAR_CV_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_UNUSED_CV_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_CV_CONST_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_OBJ_UNSET_SPEC_CV_CV_HANDLER, ZEND_FETCH_LIST_R_SPEC_CONST_CONST_HANDLER, ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_FETCH_LIST_R_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_R_SPEC_CONST_CV_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CONST_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_R_SPEC_TMPVARCV_CV_HANDLER, ZEND_FETCH_CONSTANT_SPEC_UNUSED_CONST_HANDLER, ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_HANDLER, ZEND_CHECK_FUNC_ARG_SPEC_UNUSED_QUICK_HANDLER, ZEND_EXT_STMT_SPEC_HANDLER, ZEND_EXT_FCALL_BEGIN_SPEC_HANDLER, ZEND_EXT_FCALL_END_SPEC_HANDLER, ZEND_EXT_NOP_SPEC_HANDLER, ZEND_TICKS_SPEC_HANDLER, ZEND_SEND_VAR_NO_REF_SPEC_VAR_HANDLER, ZEND_CATCH_SPEC_CONST_HANDLER, ZEND_THROW_SPEC_CONST_HANDLER, ZEND_THROW_SPEC_TMP_HANDLER, ZEND_THROW_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_THROW_SPEC_CV_HANDLER, ZEND_FETCH_CLASS_SPEC_UNUSED_CONST_HANDLER, ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_CLASS_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_FETCH_CLASS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_FETCH_CLASS_SPEC_UNUSED_CV_HANDLER, ZEND_CLONE_SPEC_CONST_HANDLER, ZEND_CLONE_SPEC_TMPVAR_HANDLER, ZEND_CLONE_SPEC_TMPVAR_HANDLER, ZEND_CLONE_SPEC_UNUSED_HANDLER, ZEND_CLONE_SPEC_CV_HANDLER, ZEND_RETURN_BY_REF_SPEC_CONST_HANDLER, ZEND_RETURN_BY_REF_SPEC_TMP_HANDLER, ZEND_RETURN_BY_REF_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_RETURN_BY_REF_SPEC_CV_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CONST_CONST_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CONST_CV_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CONST_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_TMPVAR_CV_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_UNUSED_CV_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CV_CONST_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_METHOD_CALL_SPEC_CV_CV_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CONST_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_UNUSED_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_CONST_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CONST_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_UNUSED_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_VAR_CV_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CONST_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_UNUSED_HANDLER, ZEND_INIT_STATIC_METHOD_CALL_SPEC_UNUSED_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_VAR_SPEC_CONST_UNUSED_HANDLER, ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_ISSET_ISEMPTY_VAR_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_VAR_SPEC_CV_UNUSED_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CONST_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CONST_CV_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CV_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CONST_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CONST_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_DIM_OBJ_SPEC_CV_CV_HANDLER, ZEND_SEND_VAL_EX_SPEC_CONST_HANDLER, ZEND_SEND_VAL_EX_SPEC_CONST_QUICK_HANDLER, ZEND_SEND_VAL_EX_SPEC_TMP_HANDLER, ZEND_SEND_VAL_EX_SPEC_TMP_QUICK_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_SPEC_CV_HANDLER, ZEND_INIT_USER_CALL_SPEC_CONST_CONST_HANDLER, ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_INIT_USER_CALL_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_USER_CALL_SPEC_CONST_CV_HANDLER, ZEND_SEND_ARRAY_SPEC_HANDLER, ZEND_SEND_USER_SPEC_CONST_HANDLER, ZEND_SEND_USER_SPEC_TMP_HANDLER, ZEND_SEND_USER_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_USER_SPEC_CV_HANDLER, ZEND_STRLEN_SPEC_CONST_HANDLER, ZEND_STRLEN_SPEC_TMPVAR_HANDLER, ZEND_STRLEN_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_STRLEN_SPEC_CV_HANDLER, ZEND_DEFINED_SPEC_CONST_HANDLER, ZEND_TYPE_CHECK_SPEC_CONST_HANDLER, ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER, ZEND_TYPE_CHECK_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_TYPE_CHECK_SPEC_CV_HANDLER, ZEND_VERIFY_RETURN_TYPE_SPEC_CONST_UNUSED_HANDLER, ZEND_VERIFY_RETURN_TYPE_SPEC_TMP_UNUSED_HANDLER, ZEND_VERIFY_RETURN_TYPE_SPEC_VAR_UNUSED_HANDLER, ZEND_VERIFY_RETURN_TYPE_SPEC_UNUSED_UNUSED_HANDLER, ZEND_VERIFY_RETURN_TYPE_SPEC_CV_UNUSED_HANDLER, ZEND_FE_RESET_RW_SPEC_CONST_HANDLER, ZEND_FE_RESET_RW_SPEC_TMP_HANDLER, ZEND_FE_RESET_RW_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FE_RESET_RW_SPEC_CV_HANDLER, ZEND_FE_FETCH_RW_SPEC_VAR_HANDLER, ZEND_FE_FREE_SPEC_TMPVAR_HANDLER, ZEND_INIT_DYNAMIC_CALL_SPEC_CONST_HANDLER, ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER, ZEND_INIT_DYNAMIC_CALL_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_INIT_DYNAMIC_CALL_SPEC_CV_HANDLER, ZEND_DO_ICALL_SPEC_RETVAL_UNUSED_HANDLER, ZEND_DO_ICALL_SPEC_RETVAL_USED_HANDLER, ZEND_DO_UCALL_SPEC_RETVAL_UNUSED_HANDLER, ZEND_DO_UCALL_SPEC_RETVAL_USED_HANDLER, ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_UNUSED_HANDLER, ZEND_DO_FCALL_BY_NAME_SPEC_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_OBJ_SPEC_VAR_CONST_HANDLER, ZEND_PRE_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_PRE_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_OBJ_SPEC_VAR_CV_HANDLER, ZEND_PRE_INC_OBJ_SPEC_UNUSED_CONST_HANDLER, ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_PRE_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_OBJ_SPEC_UNUSED_CV_HANDLER, ZEND_PRE_INC_OBJ_SPEC_CV_CONST_HANDLER, ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_PRE_INC_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_PRE_INC_OBJ_SPEC_CV_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_OBJ_SPEC_VAR_CONST_HANDLER, ZEND_POST_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_POST_INC_OBJ_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_OBJ_SPEC_VAR_CV_HANDLER, ZEND_POST_INC_OBJ_SPEC_UNUSED_CONST_HANDLER, ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_POST_INC_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_OBJ_SPEC_UNUSED_CV_HANDLER, ZEND_POST_INC_OBJ_SPEC_CV_CONST_HANDLER, ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_POST_INC_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_POST_INC_OBJ_SPEC_CV_CV_HANDLER, ZEND_ECHO_SPEC_CONST_HANDLER, ZEND_ECHO_SPEC_TMPVAR_HANDLER, ZEND_ECHO_SPEC_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ECHO_SPEC_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_VAR_HANDLER, ZEND_INSTANCEOF_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_CV_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_INSTANCEOF_SPEC_CV_VAR_HANDLER, ZEND_INSTANCEOF_SPEC_CV_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_GENERATOR_CREATE_SPEC_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MAKE_REF_SPEC_VAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_MAKE_REF_SPEC_CV_UNUSED_HANDLER, ZEND_DECLARE_FUNCTION_SPEC_HANDLER, ZEND_DECLARE_LAMBDA_FUNCTION_SPEC_CONST_UNUSED_HANDLER, ZEND_DECLARE_CONST_SPEC_CONST_CONST_HANDLER, ZEND_DECLARE_CLASS_SPEC_CONST_HANDLER, ZEND_DECLARE_CLASS_DELAYED_SPEC_CONST_CONST_HANDLER, ZEND_DECLARE_ANON_CLASS_SPEC_HANDLER, ZEND_ADD_ARRAY_UNPACK_SPEC_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CONST_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CONST_CV_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CONST_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_TMPVAR_CV_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CONST_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_UNUSED_CV_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CONST_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ISSET_ISEMPTY_PROP_OBJ_SPEC_CV_CV_HANDLER, ZEND_HANDLE_EXCEPTION_SPEC_HANDLER, ZEND_USER_OPCODE_SPEC_HANDLER, ZEND_ASSERT_CHECK_SPEC_HANDLER, ZEND_JMP_SET_SPEC_CONST_HANDLER, ZEND_JMP_SET_SPEC_TMP_HANDLER, ZEND_JMP_SET_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_JMP_SET_SPEC_CV_HANDLER, ZEND_UNSET_CV_SPEC_CV_UNUSED_HANDLER, ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_SET_HANDLER, ZEND_ISSET_ISEMPTY_CV_SPEC_CV_UNUSED_EMPTY_HANDLER, ZEND_FETCH_LIST_W_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_FETCH_LIST_W_SPEC_VAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_LIST_W_SPEC_VAR_CV_HANDLER, ZEND_SEPARATE_SPEC_VAR_UNUSED_HANDLER, ZEND_FETCH_CLASS_NAME_SPEC_UNUSED_HANDLER, ZEND_CALL_TRAMPOLINE_SPEC_HANDLER, ZEND_DISCARD_EXCEPTION_SPEC_HANDLER, ZEND_YIELD_SPEC_CONST_CONST_HANDLER, ZEND_YIELD_SPEC_CONST_TMP_HANDLER, ZEND_YIELD_SPEC_CONST_VAR_HANDLER, ZEND_YIELD_SPEC_CONST_UNUSED_HANDLER, ZEND_YIELD_SPEC_CONST_CV_HANDLER, ZEND_YIELD_SPEC_TMP_CONST_HANDLER, ZEND_YIELD_SPEC_TMP_TMP_HANDLER, ZEND_YIELD_SPEC_TMP_VAR_HANDLER, ZEND_YIELD_SPEC_TMP_UNUSED_HANDLER, ZEND_YIELD_SPEC_TMP_CV_HANDLER, ZEND_YIELD_SPEC_VAR_CONST_HANDLER, ZEND_YIELD_SPEC_VAR_TMP_HANDLER, ZEND_YIELD_SPEC_VAR_VAR_HANDLER, ZEND_YIELD_SPEC_VAR_UNUSED_HANDLER, ZEND_YIELD_SPEC_VAR_CV_HANDLER, ZEND_YIELD_SPEC_UNUSED_CONST_HANDLER, ZEND_YIELD_SPEC_UNUSED_TMP_HANDLER, ZEND_YIELD_SPEC_UNUSED_VAR_HANDLER, ZEND_YIELD_SPEC_UNUSED_UNUSED_HANDLER, ZEND_YIELD_SPEC_UNUSED_CV_HANDLER, ZEND_YIELD_SPEC_CV_CONST_HANDLER, ZEND_YIELD_SPEC_CV_TMP_HANDLER, ZEND_YIELD_SPEC_CV_VAR_HANDLER, ZEND_YIELD_SPEC_CV_UNUSED_HANDLER, ZEND_YIELD_SPEC_CV_CV_HANDLER, ZEND_GENERATOR_RETURN_SPEC_CONST_HANDLER, ZEND_GENERATOR_RETURN_SPEC_TMP_HANDLER, ZEND_GENERATOR_RETURN_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_GENERATOR_RETURN_SPEC_CV_HANDLER, ZEND_FAST_CALL_SPEC_HANDLER, ZEND_FAST_RET_SPEC_HANDLER, ZEND_RECV_VARIADIC_SPEC_UNUSED_HANDLER, ZEND_SEND_UNPACK_SPEC_HANDLER, ZEND_YIELD_FROM_SPEC_CONST_HANDLER, ZEND_YIELD_FROM_SPEC_TMP_HANDLER, ZEND_YIELD_FROM_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_YIELD_FROM_SPEC_CV_HANDLER, ZEND_COPY_TMP_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_BIND_GLOBAL_SPEC_CV_CONST_HANDLER, ZEND_COALESCE_SPEC_CONST_HANDLER, ZEND_COALESCE_SPEC_TMP_HANDLER, ZEND_COALESCE_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_COALESCE_SPEC_CV_HANDLER, ZEND_SPACESHIP_SPEC_CONST_CONST_HANDLER, ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER, ZEND_SPACESHIP_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SPACESHIP_SPEC_CONST_CV_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_CV_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_CONST_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SPACESHIP_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SPACESHIP_SPEC_CV_CONST_HANDLER, ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER, ZEND_SPACESHIP_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SPACESHIP_SPEC_CV_CV_HANDLER, ZEND_FUNC_NUM_ARGS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_FUNC_GET_ARGS_SPEC_CONST_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FUNC_GET_ARGS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_STATIC_PROP_R_SPEC_HANDLER, ZEND_FETCH_STATIC_PROP_W_SPEC_HANDLER, ZEND_FETCH_STATIC_PROP_RW_SPEC_HANDLER, ZEND_FETCH_STATIC_PROP_IS_SPEC_HANDLER, ZEND_FETCH_STATIC_PROP_FUNC_ARG_SPEC_HANDLER, ZEND_FETCH_STATIC_PROP_UNSET_SPEC_HANDLER, ZEND_UNSET_STATIC_PROP_SPEC_HANDLER, ZEND_ISSET_ISEMPTY_STATIC_PROP_SPEC_HANDLER, ZEND_FETCH_CLASS_CONSTANT_SPEC_CONST_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_CLASS_CONSTANT_SPEC_VAR_CONST_HANDLER, ZEND_FETCH_CLASS_CONSTANT_SPEC_UNUSED_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_BIND_LEXICAL_SPEC_TMP_CV_HANDLER, ZEND_BIND_STATIC_SPEC_CV_UNUSED_HANDLER, ZEND_FETCH_THIS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_SEND_FUNC_ARG_SPEC_VAR_HANDLER, ZEND_ISSET_ISEMPTY_THIS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_SWITCH_LONG_SPEC_CONST_CONST_HANDLER, ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_SWITCH_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SWITCH_STRING_SPEC_CONST_CONST_HANDLER, ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_SWITCH_STRING_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IN_ARRAY_SPEC_CONST_CONST_HANDLER, ZEND_IN_ARRAY_SPEC_TMP_CONST_HANDLER, ZEND_IN_ARRAY_SPEC_VAR_CONST_HANDLER, ZEND_NULL_HANDLER, ZEND_IN_ARRAY_SPEC_CV_CONST_HANDLER, ZEND_COUNT_SPEC_CONST_UNUSED_HANDLER, ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_COUNT_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_COUNT_SPEC_CV_UNUSED_HANDLER, ZEND_GET_CLASS_SPEC_CONST_UNUSED_HANDLER, ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_GET_CLASS_SPEC_TMPVAR_UNUSED_HANDLER, ZEND_GET_CLASS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_GET_CLASS_SPEC_CV_UNUSED_HANDLER, ZEND_GET_CALLED_CLASS_SPEC_UNUSED_UNUSED_HANDLER, ZEND_GET_TYPE_SPEC_CONST_UNUSED_HANDLER, ZEND_GET_TYPE_SPEC_TMP_UNUSED_HANDLER, ZEND_GET_TYPE_SPEC_VAR_UNUSED_HANDLER, ZEND_NULL_HANDLER, ZEND_GET_TYPE_SPEC_CV_UNUSED_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CONST_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CONST_CV_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CONST_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_TMPVAR_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CONST_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CV_TMPVAR_HANDLER, ZEND_NULL_HANDLER, ZEND_ARRAY_KEY_EXISTS_SPEC_CV_CV_HANDLER, ZEND_JMP_FORWARD_SPEC_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_ADD_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_SUB_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_NO_OVERFLOW_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_MUL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_NOT_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_LONG_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_CONST_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_CONST_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPZ_HANDLER, ZEND_IS_SMALLER_OR_EQUAL_DOUBLE_SPEC_TMPVARCV_TMPVARCV_JMPNZ_HANDLER, ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_INC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER, ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_INC_LONG_SPEC_CV_RETVAL_USED_HANDLER, ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_DEC_LONG_NO_OVERFLOW_SPEC_CV_RETVAL_USED_HANDLER, ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_UNUSED_HANDLER, ZEND_PRE_DEC_LONG_SPEC_CV_RETVAL_USED_HANDLER, ZEND_POST_INC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER, ZEND_POST_INC_LONG_SPEC_CV_HANDLER, ZEND_POST_DEC_LONG_NO_OVERFLOW_SPEC_CV_HANDLER, ZEND_POST_DEC_LONG_SPEC_CV_HANDLER, ZEND_QM_ASSIGN_LONG_SPEC_CONST_HANDLER, ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER, ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_QM_ASSIGN_LONG_SPEC_TMPVARCV_HANDLER, ZEND_QM_ASSIGN_DOUBLE_SPEC_CONST_HANDLER, ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER, ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_QM_ASSIGN_DOUBLE_SPEC_TMPVARCV_HANDLER, ZEND_QM_ASSIGN_NOREF_SPEC_CONST_HANDLER, ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER, ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_QM_ASSIGN_NOREF_SPEC_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CONST_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_CONST_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_TMPVAR_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CV_CONST_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_FETCH_DIM_R_INDEX_SPEC_CV_TMPVARCV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_SIMPLE_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_SIMPLE_SPEC_CV_HANDLER, ZEND_NULL_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_EX_SIMPLE_SPEC_VAR_HANDLER, ZEND_NULL_HANDLER, ZEND_SEND_VAR_EX_SIMPLE_SPEC_CV_HANDLER, ZEND_SEND_VAL_SIMPLE_SPEC_CONST_HANDLER, ZEND_SEND_VAL_EX_SIMPLE_SPEC_CONST_HANDLER, ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_UNUSED_HANDLER, ZEND_FE_FETCH_R_SIMPLE_SPEC_VAR_CV_RETVAL_USED_HANDLER, ZEND_NULL_HANDLER}
	var specs []uint32 = []uint32{0, 1 | 0x10000 | 0x20000, 26 | 0x10000 | 0x20000, 51 | 0x10000 | 0x20000 | 0x800000, 76 | 0x10000 | 0x20000, 101 | 0x10000 | 0x20000, 126 | 0x10000 | 0x20000, 151 | 0x10000 | 0x20000, 176 | 0x10000 | 0x20000, 201 | 0x10000 | 0x20000 | 0x800000, 226 | 0x10000 | 0x20000 | 0x800000, 251 | 0x10000 | 0x20000 | 0x800000, 276 | 0x10000 | 0x20000, 301 | 0x10000, 306 | 0x10000, 311 | 0x10000 | 0x20000 | 0x800000, 336 | 0x10000 | 0x20000 | 0x800000, 361 | 0x10000 | 0x20000 | 0x800000, 386 | 0x10000 | 0x20000 | 0x200000 | 0x800000, 461 | 0x10000 | 0x20000 | 0x200000 | 0x800000, 536 | 0x10000 | 0x20000 | 0x200000, 611 | 0x10000 | 0x20000 | 0x200000, 686 | 0x10000 | 0x20000 | 0x80000, 736 | 0x10000 | 0x20000 | 0x40000, 861 | 0x10000 | 0x20000 | 0x40000, 986 | 0x40000, 991 | 0x10000 | 0x20000, 1016 | 0x10000 | 0x20000, 1041 | 0x10000 | 0x20000, 1066, 1067 | 0x10000 | 0x20000, 1092 | 0x10000, 1097 | 0x10000 | 0x20000 | 0x40000, 1222, 1223 | 0x10000 | 0x80000, 1233 | 0x10000 | 0x80000, 1243 | 0x10000, 1248 | 0x10000, 1253, 1253, 1254, 1254, 1255, 1256 | 0x10000, 1261 | 0x10000, 1266 | 0x10000, 1271 | 0x10000, 1276 | 0x10000, 1281 | 0x20000, 1286, 1287 | 0x100000, 1289 | 0x10000, 1294 | 0x10000, 1299 | 0x10000 | 0x20000, 1324 | 0x20000, 1329 | 0x20000, 1334 | 0x20000, 1339, 1340, 1341, 1342 | 0x80000, 1344, 1345 | 0x10000, 1350, 1351, 1352 | 0x10000, 1357 | 0x10000 | 0x100000, 1367 | 0x10000, 1372 | 0x10000, 1377, 1378, 1379 | 0x10000 | 0x20000, 1404 | 0x10000 | 0x20000, 1429 | 0x10000, 1434 | 0x10000, 1439 | 0x10000 | 0x20000, 1464 | 0x10000 | 0x20000, 1489 | 0x10000, 1494, 1495, 1496 | 0x10000, 1501 | 0x10000 | 0x20000, 1526 | 0x10000 | 0x20000, 1551 | 0x10000, 1556 | 0x10000 | 0x20000, 1581 | 0x10000 | 0x20000, 1606 | 0x10000, 1611 | 0x10000 | 0x20000, 1636 | 0x10000 | 0x20000, 1661 | 0x10000, 1666 | 0x10000 | 0x20000, 1691 | 0x10000 | 0x20000, 1716 | 0x10000, 1721 | 0x10000 | 0x20000, 1746 | 0x10000 | 0x20000, 1771 | 0x10000, 1776 | 0x10000 | 0x20000, 1801 | 0x10000 | 0x20000, 1826 | 0x10000 | 0x20000, 1851, 1852 | 0x100000, 1854, 1855, 1856, 1857, 1858, 1859, 1860, 1861 | 0x10000, 1866 | 0x20000, 1871 | 0x10000, 1876 | 0x10000, 1881 | 0x10000 | 0x20000, 1906 | 0x10000 | 0x20000, 1931 | 0x10000, 1936 | 0x10000 | 0x20000, 1961 | 0x10000 | 0x100000, 1971 | 0x10000, 1976 | 0x20000, 1981, 1982 | 0x10000, 1987 | 0x10000, 1992, 1993 | 0x10000, 1998 | 0x10000, 2003 | 0x10000, 2008, 2009, 2010 | 0x20000, 2015 | 0x80000, 2017 | 0x80000, 2019 | 0x80000, 2021 | 0x10000 | 0x20000, 2021 | 0x10000 | 0x20000, 2046 | 0x10000 | 0x20000, 2046 | 0x10000 | 0x20000, 2071 | 0x10000, 2076, 2077 | 0x10000 | 0x20000, 2102, 2103 | 0x10000, 2108, 2109, 2110, 2111, 2112, 2113, 2114, 2115 | 0x10000 | 0x20000, 2140, 2141, 2142, 2143 | 0x10000, 2148, 2149 | 0x1000000, 2151 | 0x20000, 2156, 2157, 2158, 2159, 2160 | 0x10000 | 0x20000, 2185 | 0x10000, 2190, 2191, 2192, 2193, 2194 | 0x10000, 2199, 2200, 2201 | 0x10000, 2206 | 0x10000 | 0x20000, 2231, 2232 | 0x10000, 2237, 2238, 2239, 2240, 2241, 2242, 2243, 2244, 2245 | 0x10000, 2250, 2251, 2252, 2253, 2254, 2255 | 0x10000, 2260 | 0x10000, 2265 | 0x10000, 2270 | 0x10000, 2275 | 0x10000, 2280, 2281 | 0x10000, 2286 | 0x10000 | 0x20000, 3203}
	ZendOpcodeHandlers = labels
	ZendHandlersCount = g.SizeOf("labels") / g.SizeOf("void *")
	ZendSpecHandlers = specs
}
