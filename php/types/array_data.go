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

	/**
	 * 写操作
	 */
	Add(key ArrayKey, data *Zval) bool
	Update(key ArrayKey, data *Zval)
	Delete(key ArrayKey) bool
	Push(data *Zval) int
	Clean()

	/**
	 * Pos 操作相关
	 * pos 表示数据在内部的具体位置，具体值由具体 ArrayData 实现确定，需要满足如下条件
	 * - 每个 pos 对应 0~1 具体合法值，在无写操作的情况下对应值不变
	 * - 取值范围在 [0, Cap() - 1] 之间
	 */
	// 获取 Pos 对应位置的数据，必须用 FindEx() 返回的精确 Pos 值。
	Pos(pos ArrayPosition) *ArrayPair
	// 从传入 Pos 开始向后查找合法 Pos 位置。
	FindPos(pos ArrayPosition) (pair *ArrayPair, realPos ArrayPosition)
	// 从传入 Pos 开始向前查找合法 Pos 位置。
	FindPosReserve(pos ArrayPosition) (pair *ArrayPair, newPos ArrayPosition)
}
