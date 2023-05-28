package zend

import b "github.com/heyuuu/gophp/builtin"

func ZEND_MAP_PTR_IS_OFFSET(ptr __auto__) int { return uintPtr(ptr__ptr) & 1 }
func ZEND_MAP_PTR_OFFSET2PTR(ptr __auto__) *any {
	return (*any)((*byte)(CG__().GetMapPtrBase() + uintPtr(ptr__ptr-1)))
}
func ZEND_MAP_PTR_PTR2OFFSET(ptr *any) any {
	return any(uintPtr((*byte)(ptr)-(*byte)(CG__().GetMapPtrBase())) | 1)
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

func ZendMapPtrNew() any {
	var ptr *any
	if CG__().GetMapPtrLast() >= CG__().GetMapPtrSize() {
		/* Grow map_ptr table */
		CG__().SetMapPtrSize(ZEND_MM_ALIGNED_SIZE_EX(CG__().GetMapPtrLast()+1, 4096))
		CG__().SetMapPtrBase(Perealloc(CG__().GetMapPtrBase(), CG__().GetMapPtrSize()*b.SizeOf("void *")))
	}
	ptr = (*any)(CG__().GetMapPtrBase() + CG__().GetMapPtrLast())
	*ptr = nil
	CG__().GetMapPtrLast()++
	return ZEND_MAP_PTR_PTR2OFFSET(ptr)
}

func ZendMapPtrStartup() {
	CG__().SetMapPtrBase(nil)
	CG__().SetMapPtrSize(0)
	CG__().SetMapPtrLast(0)
}

func ZendMapPtrPostStartup() {
	GlobalMapPtrLast = CG__().GetMapPtrLast()
}

func ZendMapPtrActivate() {
	if CG__().GetMapPtrLast() != 0 {
		memset(CG__().GetMapPtrBase(), 0, CG__().GetMapPtrLast()*b.SizeOf("void *"))
	}
}

func ZendMapPtrShutdown() {
	if CG__().GetMapPtrBase() {
		Free(CG__().GetMapPtrBase())
		CG__().SetMapPtrBase(nil)
		CG__().SetMapPtrSize(0)
	}
}
