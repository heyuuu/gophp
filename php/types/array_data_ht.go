package types

var _ ArrayData = (*ArrayDataHt)(nil)

type ArrayDataHt struct {
	data     []Bucket          // 实际存储数据的地方
	indexMap map[int]uint32    // 数字索引到具体位置的映射
	keyMap   map[string]uint32 // 字符串索引到具体位置的映射
}

func (a ArrayDataHt) Len() int {
	//TODO implement me
	panic("implement me")
}

func (a ArrayDataHt) Exists(key ArrayKey) bool {
	//TODO implement me
	panic("implement me")
}

func (a ArrayDataHt) Find(key ArrayKey) *Zval {
	//TODO implement me
	panic("implement me")
}

func (a ArrayDataHt) Add(key ArrayKey, data *Zval) bool {
	//TODO implement me
	panic("implement me")
}

func (a ArrayDataHt) Update(key ArrayKey, data *Zval) {
	//TODO implement me
	panic("implement me")
}

func (a ArrayDataHt) Delete(key ArrayKey) bool {
	//TODO implement me
	panic("implement me")
}
