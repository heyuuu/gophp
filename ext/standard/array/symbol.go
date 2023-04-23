package array

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpValidVarName(var_name *byte, var_name_len int) int {
	/* first 256 bits for first character, and second 256 bits for the next */

	var charset = []uint32{0x0, 0x0, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var charset2 = []uint32{0x0, 0x3ff0000, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var i int
	var ch uint32
	if var_name_len == 0 {
		return 0
	}

	/* These are allowed as first char: [a-zA-Z_\x7f-\xff] */

	ch = uint32((*uint8)(var_name))[0]
	if zend.ZEND_BIT_TEST(charset, ch) == 0 {
		return 0
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */

	if var_name_len > 1 {
		i = 1
		for {
			ch = uint32((*uint8)(var_name))[i]
			if zend.ZEND_BIT_TEST(charset2, ch) == 0 {
				return 0
			}
			if b.PreInc(&i) >= var_name_len {
				break
			}
		}
	}
	return 1
}
func PhpPrefixVarname(result *types.Zval, prefix *types.Zval, var_name *byte, var_name_len int, add_underscore types.ZendBool) int {
	result.SetString(types.ZendStringAlloc(prefix.String().GetLen()+b.Cond(add_underscore != 0, 1, 0)+var_name_len, 0))
	memcpy(result.String().GetVal(), prefix.String().GetVal(), prefix.String().GetLen())
	if add_underscore != 0 {
		result.String().GetStr()[prefix.String().GetLen()] = '_'
	}
	memcpy(result.String().GetVal()+prefix.String().GetLen()+b.Cond(add_underscore != 0, 1, 0), var_name, var_name_len+1)
	return types.SUCCESS
}
func PhpExtractRefIfExists(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if var_name.GetStr() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			// zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.Reference())
			count++
		}
	}
	return count
}
func PhpExtractIfExists(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					continue
				}
			}
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if var_name.GetStr() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
			count++
		}
	}
	return count
}
func PhpExtractRefOverwrite(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			// zend.ZvalPtrDtor(orig_var)
			orig_var.SetReference(entry.Reference())
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractOverwrite(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
			}
			if var_name.GetStr() == "GLOBALS" {
				continue
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
			if zend.EG__().GetException() != nil {
				return -1
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			//entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
		}
		count++
	}
	return count
}
func PhpExtractRefPrefixIfExists(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						// zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.Reference())
					} else {
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		}
	}
	return count
}
func PhpExtractPrefixIfExists(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							// types.ZendStringReleaseEx(final_name.String(), 0)
							return -1
						}
					} else {
						//entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		}
	}
	return count
}
func PhpExtractRefPrefixSame(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetStr(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.StringVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						// zend.ZvalPtrDtor(orig_var)
						orig_var.SetReference(entry.Reference())
					} else {
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "this" {
				goto prefix
			}
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractPrefixSame(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if var_name.GetLen() == 0 {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
					continue
				}
			}
		prefix:
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
				if final_name.StringVal() == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					return -1
				} else {
					entry = types.ZVAL_DEREF(entry)
					if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
						if orig_var.IsIndirect() {
							orig_var = orig_var.Indirect()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
						if zend.EG__().GetException() != nil {
							// types.ZendStringReleaseEx(final_name.String(), 0)
							return -1
						}
					} else {
						//entry.TryAddRefcount()
						symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
					}
					count++
				}
			}

		} else {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
				continue
			}
			if var_name.GetStr() == "this" {
				goto prefix
			}
			entry = types.ZVAL_DEREF(entry)
			//entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractRefPrefixAll(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
			if final_name.StringVal() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.IsReference() {
					// 					entry.AddRefcount()
				} else {
					types.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
					if orig_var.IsIndirect() {
						orig_var = orig_var.Indirect()
					}
					// zend.ZvalPtrDtor(orig_var)
					orig_var.SetReference(entry.Reference())
				} else {
					symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
				}
				count++
			}
		}

	}
	return count
}
func PhpExtractPrefixAll(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if var_name.GetLen() == 0 {
				continue
			}
			PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
		}
		if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) != 0 {
			if final_name.StringVal() == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				entry = types.ZVAL_DEREF(entry)
				if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
					if orig_var.IsIndirect() {
						orig_var = orig_var.Indirect()
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
					if zend.EG__().GetException() != nil {
						// types.ZendStringReleaseEx(final_name.String(), 0)
						return -1
					}
				} else {
					//entry.TryAddRefcount()
					symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
				}
				count++
			}
		}

	}
	return count
}
func PhpExtractRefPrefixInvalid(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || var_name.GetStr() == "this" {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

				continue
			}
		}
		if final_name.StringVal() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
				if orig_var.IsIndirect() {
					orig_var = orig_var.Indirect()
				}
				// zend.ZvalPtrDtor(orig_var)
				orig_var.SetReference(entry.Reference())
			} else {
				symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
			}
			count++
		}

	}
	return count
}
func PhpExtractPrefixInvalid(arr *types.Array, symbol_table *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var num_key zend.ZendUlong
	var entry *types.Zval
	var orig_var *types.Zval
	var final_name types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		num_key = _p.GetH()
		var_name = _p.GetKey()
		entry = _z
		if var_name != nil {
			if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 || var_name.GetStr() == "this" {
				PhpPrefixVarname(&final_name, prefix, var_name.GetVal(), var_name.GetLen(), 1)
				if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

					continue
				}
			} else {
				final_name.SetStringCopy(var_name)
			}
		} else {
			var str = zend.ZendLongToStr(num_key)
			PhpPrefixVarname(&final_name, prefix, str.GetVal(), str.GetLen(), 1)
			// types.ZendStringReleaseEx(str, 0)
			if PhpValidVarName(final_name.String().GetVal(), final_name.String().GetLen()) == 0 {

				continue
			}
		}
		if final_name.StringVal() == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			entry = types.ZVAL_DEREF(entry)
			if b.Assign(&orig_var, symbol_table.KeyFind(final_name.String().GetStr())) != nil {
				if orig_var.IsIndirect() {
					orig_var = orig_var.Indirect()
				}
				zend.ZEND_TRY_ASSIGN_COPY_EX(orig_var, entry, 0)
				if zend.EG__().GetException() != nil {
					// types.ZendStringReleaseEx(final_name.String(), 0)
					return -1
				}
			} else {
				//entry.TryAddRefcount()
				symbol_table.KeyAddNew(final_name.String().GetStr(), entry)
			}
			count++
		}

	}
	return count
}
func PhpExtractRefSkip(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					if entry.IsReference() {
						// 						entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					orig_var.SetReference(entry.Reference())
					count++
				}
			}
		} else {
			if entry.IsReference() {
				// 				entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func PhpExtractSkip(arr *types.Array, symbol_table *types.Array) zend.ZendLong {
	var count = 0
	var var_name *types.String
	var entry *types.Zval
	var orig_var *types.Zval
	var __ht = arr
	for _, _p := range __ht.ForeachData() {
		var _z = _p.GetVal()
		if _z.IsIndirect() {
			_z = _z.Indirect()
			if _z.IsUndef() {
				continue
			}
		}
		var_name = _p.GetKey()
		entry = _z
		if var_name == nil {
			continue
		}
		if PhpValidVarName(var_name.GetVal(), var_name.GetLen()) == 0 {
			continue
		}
		if var_name.GetStr() == "this" {
			continue
		}
		orig_var = symbol_table.KeyFind(var_name.GetStr())
		if orig_var != nil {
			if orig_var.IsIndirect() {
				orig_var = orig_var.Indirect()
				if orig_var.IsUndef() {
					types.ZVAL_COPY_DEREF(orig_var, entry)
					count++
				}
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			// entry.TryAddRefcount()
			symbol_table.KeyAddNew(var_name.GetStr(), entry)
			count++
		}
	}
	return count
}
func ZifExtract(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var var_array_param *types.Zval
	var prefix *types.Zval = nil
	var extract_refs zend.ZendLong
	var extract_type = EXTR_OVERWRITE
	var count zend.ZendLong
	var symbol_table *types.Array
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			var_array_param = fp.ParseArrayEx2(false, true, false)
			fp.StartOptional()
			extract_type = fp.ParseLong()
			prefix = fp.ParseZval()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	extract_refs = extract_type & EXTR_REFS
	if extract_refs != 0 {
		types.SeparateArray(var_array_param)
	}
	extract_type &= 0xff
	if extract_type < EXTR_OVERWRITE || extract_type > EXTR_IF_EXISTS {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid extract type")
		return
	}
	if extract_type > EXTR_SKIP && extract_type <= EXTR_PREFIX_IF_EXISTS && executeData.NumArgs() < 3 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "specified extract type requires the prefix parameter")
		return
	}
	if prefix != nil {
		if zend.TryConvertToString(prefix) == 0 {
			return
		}
		if prefix.String().GetLen() != 0 && PhpValidVarName(prefix.String().GetVal(), prefix.String().GetLen()) == 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "prefix is not a valid identifier")
			return
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if extract_refs != 0 {
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractRefIfExists(var_array_param.Array(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractRefOverwrite(var_array_param.Array(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractRefPrefixIfExists(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractRefPrefixSame(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractRefPrefixAll(var_array_param.Array(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractRefPrefixInvalid(var_array_param.Array(), symbol_table, prefix)
		default:
			count = PhpExtractRefSkip(var_array_param.Array(), symbol_table)
		}
	} else {

		/* The array might be stored in a local variable that will be overwritten */

		var array_copy types.Zval
		types.ZVAL_COPY(&array_copy, var_array_param)
		switch extract_type {
		case EXTR_IF_EXISTS:
			count = PhpExtractIfExists(array_copy.Array(), symbol_table)
		case EXTR_OVERWRITE:
			count = PhpExtractOverwrite(array_copy.Array(), symbol_table)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractPrefixIfExists(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractPrefixSame(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractPrefixAll(array_copy.Array(), symbol_table, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractPrefixInvalid(array_copy.Array(), symbol_table, prefix)
		default:
			count = PhpExtractSkip(array_copy.Array(), symbol_table)
		}
		// zend.ZvalPtrDtor(&array_copy)
	}
	return_value.SetLong(count)
	return
}

func PhpCompactVar(eg_active_symbol_table *types.Array, return_value *types.Zval, entry *types.Zval) {
	var value_ptr *types.Zval
	var data types.Zval
	entry = types.ZVAL_DEREF(entry)
	if entry.IsString() {
		if b.Assign(&value_ptr, types.ZendHashFindInd(eg_active_symbol_table, entry.String().GetStr())) != nil {
			value_ptr = types.ZVAL_DEREF(value_ptr)
			// value_ptr.TryAddRefcount()
			return_value.Array().KeyUpdate(entry.String().GetStr(), value_ptr)
		} else if entry.StringVal() == "this" {
			var object = zend.ZendGetThisObject(zend.CurrEX())
			if object != nil {
				// 				object.AddRefcount()
				data.SetObject(object)
				return_value.Array().KeyUpdate(entry.String().GetStr(), &data)
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_NOTICE, "Undefined variable: %s", entry.String().GetVal())
		}
	} else if entry.IsType(types.IS_ARRAY) {
		if entry.IsRefcounted() {
			if entry.IsRecursive() {
				core.PhpErrorDocref(nil, faults.E_WARNING, "recursion detected")
				return
			}
			entry.ProtectRecursive()
		}
		var __ht = entry.Array()
		for _, _p := range __ht.ForeachData() {
			var _z = _p.GetVal()
			if _z.IsIndirect() {
				_z = _z.Indirect()
				if _z.IsUndef() {
					continue
				}
			}
			value_ptr = _z
			PhpCompactVar(eg_active_symbol_table, return_value, value_ptr)
		}
		if entry.IsRefcounted() {
			entry.UnprotectRecursive()
		}
	}
}
func ZifCompact(executeData zpp.Ex, return_value zpp.Ret, varNames []*types.Zval) {
	var args *types.Zval = nil
	var num_args uint32
	var i uint32
	var symbol_table *types.Array
	for {
		var _flags = 0
		var _min_num_args = 1
		var _max_num_args = -1

		for {
			fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			args, num_args = fp.ParseVariadic0()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if zend.ZendForbidDynamicCall("compact()") == types.FAILURE {
		return
	}
	symbol_table = zend.ZendRebuildSymbolTable()
	if symbol_table == nil {
		return
	}

	/* compact() is probably most used with a single array of var_names
	   or multiple string names, rather than a combination of both.
	   So quickly guess a minimum result size based on that */

	if num_args != 0 && args[0].IsType(types.IS_ARRAY) {
		zend.ArrayInitSize(return_value, args[0].Array().Len())
	} else {
		zend.ArrayInitSize(return_value, num_args)
	}
	for i = 0; i < num_args; i++ {
		PhpCompactVar(symbol_table, return_value, &args[i])
	}
}
