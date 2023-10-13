package array

import (
	"github.com/heyuuu/gophp/core"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
	"strconv"
)

func bitTest(bits []uint32, c byte) bool {
	return bits[c/32]&(1<<c%32) != 0
}
func isValidVarName(varName string) bool {
	if varName == "" {
		return false
	}

	/* first 256 bits for first character, and second 256 bits for the next */
	var charset = []uint32{0x0, 0x0, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}
	var charset2 = []uint32{0x0, 0x3ff0000, 0x87fffffe, 0x7fffffe, 0xffffffff, 0xffffffff, 0xffffffff, 0xffffffff}

	/* These are allowed as first char: [a-zA-Z_\x7f-\xff] */
	if !bitTest(charset, varName[0]) {
		return false
	}

	/* And these as the rest: [a-zA-Z0-9_\x7f-\xff] */
	for i := 1; i < len(varName); i++ {
		if !bitTest(charset2, varName[i]) {
			return false
		}
	}

	return true
}

func prefixVarName(prefix string, varName string) string {
	return prefix + "_" + varName
}

func phpExtractEx(arr *types.Array, handler func(key types.ArrayKey, entry *types.Zval) int) int {
	var count = 0
	arr.ForeachIndirectEx(func(key types.ArrayKey, value *types.Zval) bool {
		ret := handler(key, value)
		if ret < 0 {
			count = -1
			return false
		} else {
			count += ret
			return true
		}
	})
	return count
}
func phpExtract(arr *types.Array, handler func(varName string, entry *types.Zval) int) int {
	return phpExtractEx(arr, func(key types.ArrayKey, entry *types.Zval) int {
		if !key.IsStrKey() {
			return 0
		}
		return handler(key.StrKey(), entry)
	})
}

func PhpExtractRefIfExists(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					return 0
				}
			}
			if !isValidVarName(varName) {
				return 0
			}
			if varName == "GLOBALS" {
				return 0
			}
			if varName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			origVar.SetReference(entry.Ref())
			return 1
		}
		return 0
	})
}
func PhpExtractIfExists(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					return 0
				}
			}
			if !isValidVarName(varName) {
				return 0
			}
			if varName == "GLOBALS" {
				return 0
			}
			if varName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(origVar, entry, 0)
			if zend.EG__().HasException() {
				return -1
			}
			return 1
		}
		return 0
	})
}
func PhpExtractRefOverwrite(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		if !isValidVarName(varName) {
			return 0
		}
		if varName == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
			}
			if varName == "GLOBALS" {
				return 0
			}
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			origVar.SetReference(entry.Ref())
		} else {
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbolTable.KeyAddNew(varName, entry)
		}
		return 1
	})
}
func PhpExtractOverwrite(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		if varName == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		}
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
			}
			if varName == "GLOBALS" {
				return 0
			}
			entry = types.ZVAL_DEREF(entry)
			zend.ZEND_TRY_ASSIGN_COPY_EX(origVar, entry, 0)
			if zend.EG__().HasException() {
				return -1
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			symbolTable.KeyAddNew(varName, entry)
		}
		return 1
	})
}
func PhpExtractRefPrefixIfExists(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	arr.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
		if !key.IsStrKey() {
			return true
		}
		varName := key.StrKey()
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					if entry.IsRef() {
						// entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					origVar.SetReference(entry.Ref())
					count++
					return true
				}
			}
			finalName := prefixVarName(prefix.String(), varName)
			if isValidVarName(finalName) {
				if finalName == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					count = -1
					return false
				} else {
					if entry.IsRef() {
						// entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}

					if origVar2 := symbolTable.KeyFind(finalName); origVar2 != nil {
						if origVar2.IsIndirect() {
							origVar2 = origVar2.Indirect()
						}
						origVar2.SetReference(entry.Ref())
					} else {
						symbolTable.KeyAddNew(finalName, entry)
					}
					count++
				}
			}

		}
		return true
	})
	return count
}
func PhpExtractPrefixIfExists(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	arr.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
		if !key.IsStrKey() {
			return true
		}
		varName := key.StrKey()
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					types.ZVAL_COPY_DEREF(origVar, entry)
					count++
					return true
				}
			}
			finalName := prefixVarName(prefix.String(), varName)
			if isValidVarName(finalName) {
				if finalName == "this" {
					faults.ThrowError(nil, "Cannot re-assign $this")
					count = -1
					return false
				} else {
					entry = types.ZVAL_DEREF(entry)
					if origVar2 := symbolTable.KeyFind(finalName); origVar2 != nil {
						if origVar2.IsIndirect() {
							origVar2 = origVar2.Indirect()
						}
						zend.ZEND_TRY_ASSIGN_COPY_EX(origVar2, entry, 0)
						if zend.EG__().HasException() {
							count = -1
							return false
						}
					} else {
						symbolTable.KeyAddNew(finalName, entry)
					}
					count++
				}
			}
		}
		return true
	})
	return count
}
func PhpExtractRefPrefixSame(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	var count = 0
	arr.ForeachIndirectEx(func(key types.ArrayKey, entry *types.Zval) bool {
		if !key.IsStrKey() {
			return true
		}

		varName := key.StrKey()
		if varName == "" {
			return true
		}

		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					if entry.IsRef() {
						// entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					origVar.SetReference(entry.Ref())
					count++
					return true
				}
			}
			goto prefix
		} else {
			if !isValidVarName(varName) {
				return true
			}
			if varName == "this" {
				goto prefix
			}
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbolTable.KeyAddNew(varName, entry)
			count++
			return true
		}
	prefix:
		finalName := prefixVarName(prefix.String(), varName)
		if isValidVarName(finalName) {
			if finalName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				count = -1
				return false
			} else {
				if entry.IsRef() {
					// entry.AddRefcount()
				} else {
					types.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if origVar2 := symbolTable.KeyFind(finalName); origVar2 != nil {
					if origVar2.IsIndirect() {
						origVar2 = origVar2.Indirect()
					}
					origVar.SetReference(entry.Ref())
				} else {
					symbolTable.KeyAddNew(finalName, entry)
				}
				count++
			}
		}
		return true
	})
	return count
}
func PhpExtractPrefixSame(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		if varName == "" {
			return 0
		}
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					types.ZVAL_COPY_DEREF(origVar, entry)
					return 1
				}
			}
			goto prefix
		} else {
			if !isValidVarName(varName) {
				return 0
			}
			if varName == "this" {
				goto prefix
			}
			entry = types.ZVAL_DEREF(entry)
			symbolTable.KeyAddNew(varName, entry)
			return 1
		}
	prefix:
		finalName := prefixVarName(prefix.String(), varName)
		if isValidVarName(finalName) {
			if finalName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				entry = types.ZVAL_DEREF(entry)
				if b.Assign(&origVar, symbolTable.KeyFind(finalName)) != nil {
					if origVar.IsIndirect() {
						origVar = origVar.Indirect()
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(origVar, entry, 0)
					if zend.EG__().HasException() {
						return -1
					}
				} else {
					symbolTable.KeyAddNew(finalName, entry)
				}
				return 1
			}
		}
		return 0
	})
}
func PhpExtractRefPrefixAll(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	return phpExtractEx(arr, func(key types.ArrayKey, entry *types.Zval) int {
		var finalName string
		if key.IsStrKey() {
			finalName = prefixVarName(prefix.String(), key.StrKey())
		} else {
			finalName = prefixVarName(prefix.String(), strconv.Itoa(key.IdxKey()))
		}

		if isValidVarName(finalName) {
			if finalName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				if entry.IsRef() {
					// entry.AddRefcount()
				} else {
					types.ZVAL_MAKE_REF_EX(entry, 2)
				}
				if origVar := symbolTable.KeyFind(finalName); origVar != nil {
					if origVar.IsIndirect() {
						origVar = origVar.Indirect()
					}
					origVar.SetReference(entry.Ref())
				} else {
					symbolTable.KeyAddNew(finalName, entry)
				}
				return 1
			}
		}
		return 0
	})
}
func PhpExtractPrefixAll(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	return phpExtractEx(arr, func(key types.ArrayKey, entry *types.Zval) int {
		var finalName string
		if key.IsStrKey() {
			finalName = prefixVarName(prefix.String(), key.StrKey())
		} else {
			finalName = prefixVarName(prefix.String(), strconv.Itoa(key.IdxKey()))
		}

		if isValidVarName(finalName) {
			if finalName == "this" {
				faults.ThrowError(nil, "Cannot re-assign $this")
				return -1
			} else {
				entry = types.ZVAL_DEREF(entry)
				if origVar := symbolTable.KeyFind(finalName); origVar != nil {
					if origVar.IsIndirect() {
						origVar = origVar.Indirect()
					}
					zend.ZEND_TRY_ASSIGN_COPY_EX(origVar, entry, 0)
					if zend.EG__().HasException() {
						return -1
					}
				} else {
					symbolTable.KeyAddNew(finalName, entry)
				}
				return 1
			}
		}

		return 0
	})
}
func PhpExtractRefPrefixInvalid(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	return phpExtractEx(arr, func(key types.ArrayKey, entry *types.Zval) int {
		var finalName string
		if key.IsStrKey() {
			varName := key.StrKey()
			if !isValidVarName(varName) || varName == "this" {
				finalName = prefixVarName(prefix.String(), varName)
			} else {
				finalName = varName
			}
		} else {
			finalName = prefixVarName(prefix.String(), strconv.Itoa(key.IdxKey()))
		}
		if !isValidVarName(finalName) {
			return 0
		}

		if finalName == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			if origVar := symbolTable.KeyFind(finalName); origVar != nil {
				if origVar.IsIndirect() {
					origVar = origVar.Indirect()
				}
				origVar.SetReference(entry.Ref())
			} else {
				symbolTable.KeyAddNew(finalName, entry)
			}
			return 1
		}
	})
}
func PhpExtractPrefixInvalid(arr *types.Array, symbolTable *types.Array, prefix *types.Zval) zend.ZendLong {
	return phpExtractEx(arr, func(key types.ArrayKey, entry *types.Zval) int {
		var finalName string
		if key.IsStrKey() {
			varName := key.StrKey()
			if !isValidVarName(varName) || varName == "this" {
				finalName = prefixVarName(prefix.String(), varName)
			} else {
				finalName = varName
			}
		} else {
			finalName = prefixVarName(prefix.String(), strconv.Itoa(key.IdxKey()))
		}
		if !isValidVarName(finalName) {
			return 0
		}
		if finalName == "this" {
			faults.ThrowError(nil, "Cannot re-assign $this")
			return -1
		} else {
			entry = types.ZVAL_DEREF(entry)
			if origVar := symbolTable.KeyFind(finalName); origVar != nil {
				if origVar.IsIndirect() {
					origVar = origVar.Indirect()
				}
				zend.ZEND_TRY_ASSIGN_COPY_EX(origVar, entry, 0)
				if zend.EG__().HasException() {
					return -1
				}
			} else {
				symbolTable.KeyAddNew(finalName, entry)
			}
			return 1
		}
	})
}
func PhpExtractRefSkip(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		if !isValidVarName(varName) || varName == "this" {
			return 0
		}
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					if entry.IsRef() {
						// entry.AddRefcount()
					} else {
						types.ZVAL_MAKE_REF_EX(entry, 2)
					}
					origVar.SetReference(entry.Ref())
					return 1
				}
			}
		} else {
			if entry.IsRef() {
				// entry.AddRefcount()
			} else {
				types.ZVAL_MAKE_REF_EX(entry, 2)
			}
			symbolTable.KeyAddNew(varName, entry)
			return 1
		}
		return 0
	})
}
func PhpExtractSkip(arr *types.Array, symbolTable *types.Array) zend.ZendLong {
	return phpExtract(arr, func(varName string, entry *types.Zval) int {
		if !isValidVarName(varName) || varName == "this" {
			return 0
		}
		origVar := symbolTable.KeyFind(varName)
		if origVar != nil {
			if origVar.IsIndirect() {
				origVar = origVar.Indirect()
				if origVar.IsUndef() {
					types.ZVAL_COPY_DEREF(origVar, entry)
					return 1
				}
			}
		} else {
			entry = types.ZVAL_DEREF(entry)
			symbolTable.KeyAddNew(varName, entry)
			return 1
		}
		return 0
	})
}
func ZifExtract(array zpp.DerefArray, _ zpp.Opt, flags int, prefix *types.Zval) int {
	extractRefs := flags&EXTR_REFS != 0
	if extractRefs {
		types.SeparateArray(array)
	}

	extractType := flags & 0xff
	if extractType < EXTR_OVERWRITE || extractType > EXTR_IF_EXISTS {
		core.PhpErrorDocref("", faults.E_WARNING, "Invalid extract type")
		return 0
	}
	if extractType > EXTR_SKIP && extractType <= EXTR_PREFIX_IF_EXISTS && prefix == nil {
		core.PhpErrorDocref("", faults.E_WARNING, "specified extract type requires the prefix parameter")
		return 0
	}
	if prefix != nil {
		if operators.TryConvertToString(prefix) == 0 {
			return 0
		}
		if prefix.String() != "" && !isValidVarName(prefix.String()) {
			core.PhpErrorDocref("", faults.E_WARNING, "prefix is not a valid identifier")
			return 0
		}
	}
	if zend.ZendForbidDynamicCall("extract()") == types.FAILURE {
		return 0
	}

	var count int
	var symbolTable = zend.ZendRebuildSymbolTable()
	if extractRefs {
		switch extractType {
		case EXTR_IF_EXISTS:
			count = PhpExtractRefIfExists(array.Array(), symbolTable)
		case EXTR_OVERWRITE:
			count = PhpExtractRefOverwrite(array.Array(), symbolTable)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractRefPrefixIfExists(array.Array(), symbolTable, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractRefPrefixSame(array.Array(), symbolTable, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractRefPrefixAll(array.Array(), symbolTable, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractRefPrefixInvalid(array.Array(), symbolTable, prefix)
		default:
			count = PhpExtractRefSkip(array.Array(), symbolTable)
		}
	} else {
		/* The array might be stored in a local variable that will be overwritten */
		var arrayCopy types.Zval
		types.ZVAL_COPY(&arrayCopy, array)
		switch extractType {
		case EXTR_IF_EXISTS:
			count = PhpExtractIfExists(arrayCopy.Array(), symbolTable)
		case EXTR_OVERWRITE:
			count = PhpExtractOverwrite(arrayCopy.Array(), symbolTable)
		case EXTR_PREFIX_IF_EXISTS:
			count = PhpExtractPrefixIfExists(arrayCopy.Array(), symbolTable, prefix)
		case EXTR_PREFIX_SAME:
			count = PhpExtractPrefixSame(arrayCopy.Array(), symbolTable, prefix)
		case EXTR_PREFIX_ALL:
			count = PhpExtractPrefixAll(arrayCopy.Array(), symbolTable, prefix)
		case EXTR_PREFIX_INVALID:
			count = PhpExtractPrefixInvalid(arrayCopy.Array(), symbolTable, prefix)
		default:
			count = PhpExtractSkip(arrayCopy.Array(), symbolTable)
		}
	}
	return count
}

func PhpCompactVar(activeSymbolTable *types.Array, resultArr *types.Array, entry *types.Zval) {
	entry = types.ZVAL_DEREF(entry)
	if entry.IsString() {
		if valuePtr := types.ZendHashFindInd(activeSymbolTable, entry.StringEx().GetStr()); valuePtr != nil {
			valuePtr = types.ZVAL_DEREF(valuePtr)
			resultArr.KeyUpdate(entry.StringEx().GetStr(), valuePtr)
		} else if entry.String() == "this" {
			var object = zend.ZendGetThisObject(zend.CurrEX())
			if object != nil {
				resultArr.KeyUpdate(entry.StringEx().GetStr(), types.NewZvalObject(object))
			}
		} else {
			core.PhpErrorDocref("", faults.E_NOTICE, "Undefined variable: %s", entry.StringEx().GetVal())
		}
	} else if entry.IsArray() {
		if entry.Array().IsRecursive() {
			core.PhpErrorDocref("", faults.E_WARNING, "recursion detected")
			return
		}
		entry.Array().ProtectRecursive()
		entry.Array().ForeachIndirect(func(key types.ArrayKey, valuePtr *types.Zval) {
			PhpCompactVar(activeSymbolTable, resultArr, valuePtr)
		})
		entry.Array().UnprotectRecursive()
	}
}

//@zif -c=1,
func ZifCompact(varNames []*types.Zval) *types.Array {
	if zend.ZendForbidDynamicCall("compact()") == types.FAILURE {
		return nil
	}

	symbolTable := zend.ZendRebuildSymbolTable()
	if symbolTable == nil {
		return nil
	}

	retArr := types.NewArrayCap(len(varNames))
	for _, varName := range varNames {
		PhpCompactVar(symbolTable, retArr, varName)
	}
	return retArr
}
