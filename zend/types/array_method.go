package types

/**
 * Open methods
 */
func (ht *Array) RealInit() { ht.clearData() } // todo remove 无需init操作

/**
 * Methods by index key
 */
func (ht *Array) indexFindBucket(index int) *Bucket {
	if pos, ok := ht.indexMap[index]; ok {
		return &ht.data[pos]
	}
	return nil
}
func (ht *Array) IndexFind(index int) *Zval {
	if pos, ok := ht.indexMap[index]; ok {
		return &ht.data[pos].val
	}
	return nil
}
func (ht *Array) IndexExists(index int) bool {
	_, ok := ht.indexMap[index]
	return ok
}

/**
 * Internal methods
 */
func (ht *Array) copyDataAndHash(source *Array) {
	ht.data = make([]Bucket, len(source.data))
	copy(ht.data, source.data)

	ht.indexMap = make(map[int]uint32)
	for i, pos := range source.indexMap {
		ht.indexMap[i] = pos
	}

	ht.keyMap = make(map[string]uint32)
	for i, pos := range source.keyMap {
		ht.keyMap[i] = pos
	}
}
