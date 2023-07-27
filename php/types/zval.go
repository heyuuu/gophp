package types

/** Zval.u2 */
func (zv *Zval) GetCacheSlot() uint32      { return zv.u2 }
func (zv *Zval) SetCacheSlot(value uint32) { zv.u2 = value }
func (zv *Zval) GetOplineNum() uint32      { return zv.u2 }
func (zv *Zval) SetOplineNum(value uint32) { zv.u2 = value }
func (zv *Zval) GetFePos() uint32          { return zv.u2 }
func (zv *Zval) SetFePos(value uint32)     { zv.u2 = value }
func (zv *Zval) GetFeIterIdx() uint32      { return zv.u2 }
func (zv *Zval) SetFeIterIdx(value uint32) { zv.u2 = value }
func (zv *Zval) GetU2Extra() uint32        { return zv.u2 }
func (zv *Zval) SetU2Extra(value uint32)   { zv.u2 = value }
