package types

import (
	b "sik/builtin"
	"sik/zend"
)

func (this *HashTable) SymtableClean() {
	// todo 这里可能不会严格对等，需要处理一下
	b.Assert(this.pDestructor == zend.ZVAL_PTR_DTOR)

	this.Clean()
}

func (ht *HashTable) SymtableAddNew(key string, pData *Zval) *Zval {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexAddNew(idx, pData)
	} else {
		return ht.KeyAddNew(key, pData)
	}
}
func (ht *HashTable) SymtableUpdate(key string, pData *Zval) *Zval {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.KeyUpdate(key, pData)
	}
}
func (ht *HashTable) SymtableUpdateInd(key string, pData *Zval) *Zval {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexUpdate(idx, pData)
	} else {
		return ht.KeyUpdateIndirect(key, pData)
	}
}
func (ht *HashTable) SymtableDel(key string) int {
	var result bool
	if idx, ok := zend.zendParseNumericStr(key); ok {
		result = ht.IndexDelete(idx)
	} else {
		result = ht.KeyDelete(key)
	}
	return resultCode(result)
}
func (ht *HashTable) SymtableFind(key string) *Zval {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexFind(idx)
	} else {
		return ht.KeyFind(key)
	}
}
func (ht *HashTable) SymtableExists(key string) bool {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexExists(idx)
	} else {
		return ht.KeyExists(key)
	}
}

func (ht *HashTable) SymtableExistsInd(key string) bool {
	if idx, ok := zend.zendParseNumericStr(key); ok {
		return ht.IndexExists(idx)
	} else {
		return ht.KeyExistsInd(key)
	}
}
