// <<generate>>

package zend

func ZEND_MAP_PTR_IS_OFFSET(ptr __auto__) int { return uintPtr(ptr__ptr) & 1 }
func ZEND_MAP_PTR_OFFSET2PTR(ptr __auto__) *any {
	return (*any)((*byte)(CG(map_ptr_base) + uintPtr(ptr__ptr-1)))
}
func ZEND_MAP_PTR_PTR2OFFSET(ptr *any) any {
	return any(uintPtr((*byte)(ptr)-(*byte)(CG(map_ptr_base))) | 1)
}
func ZEND_MAP_PTR_GET(ptr __auto__) any {
	if ZEND_MAP_PTR_IS_OFFSET(ptr) != 0 {
		return *(ZEND_MAP_PTR_OFFSET2PTR(ptr))
	} else {
		return any(*ptr__ptr)
	}
}
func ZEND_MAP_PTR_SET(ptr __auto__, val any) {
	if ZEND_MAP_PTR_IS_OFFSET(ptr) != 0 {
		*(ZEND_MAP_PTR_OFFSET2PTR(ptr)) = val
	} else {
		*ptr__ptr = val
	}
}
func ZEND_MAP_PTR_INIT(ptr __auto__, val __auto__) { ptr__ptr = val }
func ZEND_MAP_PTR_NEW(ptr __auto__)                { ptr__ptr = ZendMapPtrNew() }
