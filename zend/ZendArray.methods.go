package zend

// constructor

func NewZendArray(size uint32) *ZendArray {
	return NewZendArrayEx(size, ZVAL_PTR_DTOR, false)
}

func NewZendArrayEx(size uint32, pDestructor DtorFuncT, persistent bool) *ZendArray {
	var ht = &ZendArray{
		nNumUsed:         0,
		nNumOfElements:   0,
		nTableSize:       ZendHashCheckSize(size),
		nInternalPointer: 0,
		nNextFreeElement: 0,
		pDestructor:      pDestructor,

		// 数据存储
		data:     nil,
		indexMap: make(map[int]uint32),
		keyMap:   make(map[string]uint32),
	}

	// GC 信息
	ht.SetRefcount(1)
	ht.SetGcTypeInfo(IS_ARRAY)
	if persistent {
		ht.AddGcFlags(GC_PERSISTENT)
	} else {
		ht.AddGcFlags(GC_COLLECTABLE)
	}

	return ht
}

func (this *ZendArray) assertRc1() {
	ZEND_ASSERT(this.GetRefcount() == 1)
}

func (this *ZendArray) realInit() {
	this.assertRc1()

	this.nNumUsed = 0
	this.nNumOfElements = 0
	this.data = nil
	this.indexMap = make(map[int]uint32)
	this.keyMap = make(map[string]uint32)

	this.SetIsStaticKeys()
}
