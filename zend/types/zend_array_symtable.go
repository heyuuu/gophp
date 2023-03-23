package types

import (
	b "sik/builtin"
	"sik/zend"
)

func (ht *Array) SymtableClean() {
	// todo 这里可能不会严格对等，需要处理一下
	b.Assert(ht.destructor == zend.ZVAL_PTR_DTOR)

	ht.Clean()
}

func (ht *Array) SymtableAddNew(key string, pData *Zval) *Zval {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexAddNew(idx, pData)
	} else {
		return ht.KeyAddNew(key, pData)
	}
}
func (ht *Array) SymtableUpdate(key string, pData *Zval) *Zval {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.KeyUpdate(key, pData)
	}
}
func (ht *Array) SymtableUpdateInd(key string, pData *Zval) *Zval {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.KeyUpdateIndirect(key, pData)
	}
}
func (ht *Array) SymtableDel(key string) int {
	var result bool
	if idx, ok := parseNumericStr(key); ok {
		result = ht.IndexDelete(idx)
	} else {
		result = ht.KeyDelete(key)
	}
	return ResultCode(result)
}
func (ht *Array) SymtableFind(key string) *Zval {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexFind(idx)
	} else {
		return ht.KeyFind(key)
	}
}
func (ht *Array) SymtableExists(key string) bool {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexExists(idx)
	} else {
		return ht.KeyExists(key)
	}
}

func (ht *Array) SymtableExistsInd(key string) bool {
	if idx, ok := parseNumericStr(key); ok {
		return ht.IndexExists(idx)
	} else {
		return ht.KeyExistsInd(key)
	}
}
