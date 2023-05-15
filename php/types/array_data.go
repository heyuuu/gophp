package types

type ArrayData interface {
	/**
	 * 读操作
	 */
	Len() int
	Cap() int
	Exists(key ArrayKey) bool
	Find(key ArrayKey) *Zval
	FindEx(key ArrayKey) (value *Zval, pos ArrayPosition)

	// 获取 Pos 对应位置的数据，必须用 FindEx() 返回的精确 Pos 值，不可用 Next/Prev 返回的 newPos
	Pos(pos ArrayPosition) *ArrayPair

	/**
	 * 写操作
	 */
	Add(key ArrayKey, data *Zval) bool
	Update(key ArrayKey, data *Zval)
	Delete(key ArrayKey) bool
	Push(data *Zval) int
	Clean()

	/**
	 * 获取指定 Pos 对应的数据和下一个指针位置
	 * pos 具体值由具体实现确定，需要满足如下条件
	 * - 未进行写操作时，pos 对应数据不会发生改变
	 */
	Next(pos ArrayPosition) (pair *ArrayPair, newPos ArrayPosition)
	Prev(pos ArrayPosition) (pair *ArrayPair, newPos ArrayPosition)
}
