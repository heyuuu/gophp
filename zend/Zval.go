// <<generate>>

package zend

/**
 * ZendValue
 */
type ZendValue struct /* union */ {
	lval    ZendLong
	dval    float64
	counted *ZendRefcounted
	str     *ZendString
	arr     *ZendArray
	obj     *ZendObject
	res     *ZendResource
	ref     *ZendReference
	ast     *ZendAstRef
	zv      *Zval
	ptr     any
	ce      *ZendClassEntry
	func_   *ZendFunction
	ww      struct {
		w1 uint32
		w2 uint32
	}
}

func (this ZendValue) GetLval() ZendLong                 { return this.lval }
func (this *ZendValue) SetLval(value ZendLong)           { this.lval = value }
func (this ZendValue) GetDval() float64                  { return this.dval }
func (this *ZendValue) SetDval(value float64)            { this.dval = value }
func (this ZendValue) GetCounted() *ZendRefcounted       { return this.counted }
func (this *ZendValue) SetCounted(value *ZendRefcounted) { this.counted = value }
func (this ZendValue) GetStr() *ZendString               { return this.str }
func (this *ZendValue) SetStr(value *ZendString)         { this.str = value }
func (this ZendValue) GetArr() *ZendArray                { return this.arr }
func (this *ZendValue) SetArr(value *ZendArray)          { this.arr = value }
func (this ZendValue) GetObj() *ZendObject               { return this.obj }
func (this *ZendValue) SetObj(value *ZendObject)         { this.obj = value }
func (this ZendValue) GetRes() *ZendResource             { return this.res }
func (this *ZendValue) SetRes(value *ZendResource)       { this.res = value }
func (this ZendValue) GetRef() *ZendReference            { return this.ref }
func (this *ZendValue) SetRef(value *ZendReference)      { this.ref = value }
func (this ZendValue) GetAst() *ZendAstRef               { return this.ast }
func (this *ZendValue) SetAst(value *ZendAstRef)         { this.ast = value }
func (this ZendValue) GetZv() *Zval                      { return this.zv }
func (this *ZendValue) SetZv(value *Zval)                { this.zv = value }
func (this ZendValue) GetPtr() any                       { return this.ptr }
func (this *ZendValue) SetPtr(value any)                 { this.ptr = value }
func (this ZendValue) GetCe() *ZendClassEntry            { return this.ce }
func (this *ZendValue) SetCe(value *ZendClassEntry)      { this.ce = value }
func (this ZendValue) GetFunc() *ZendFunction            { return this.func_ }
func (this *ZendValue) SetFunc(value *ZendFunction)      { this.func_ = value }
func (this ZendValue) GetW1() uint32                     { return this.ww.w1 }
func (this *ZendValue) SetW1(value uint32)               { this.ww.w1 = value }
func (this ZendValue) GetW2() uint32                     { return this.ww.w2 }
func (this *ZendValue) SetW2(value uint32)               { this.ww.w2 = value }

/**
 * Zval
 */
type Zval struct {
	value ZendValue
	u1    struct /* union */ {
		v struct {
			type_      ZendUchar
			type_flags ZendUchar
			u          struct /* union */ {
				extra uint16
			}
		}
		type_info uint32
	}
	u2 struct /* union */ {
		next           uint32
		cache_slot     uint32
		opline_num     uint32
		lineno         uint32
		num_args       uint32
		fe_pos         uint32
		fe_iter_idx    uint32
		access_flags   uint32
		property_guard uint32
		constant_flags uint32
		extra          uint32
	}
}

func (this Zval) GetValue() ZendValue            { return this.value }
func (this *Zval) SetValue(value ZendValue)      { this.value = value }
func (this Zval) GetType() ZendUchar             { return this.u1.v.type_ }
func (this *Zval) SetType(value ZendUchar)       { this.u1.v.type_ = value }
func (this Zval) GetTypeFlags() ZendUchar        { return this.u1.v.type_flags }
func (this *Zval) SetTypeFlags(value ZendUchar)  { this.u1.v.type_flags = value }
func (this Zval) GetU1VUExtra() uint16           { return this.u1.v.u.extra }
func (this *Zval) SetU1VUExtra(value uint16)     { this.u1.v.u.extra = value }
func (this Zval) GetTypeInfo() uint32            { return this.u1.type_info }
func (this *Zval) SetTypeInfo(value uint32)      { this.u1.type_info = value }
func (this Zval) GetNext() uint32                { return this.u2.next }
func (this *Zval) SetNext(value uint32)          { this.u2.next = value }
func (this Zval) GetCacheSlot() uint32           { return this.u2.cache_slot }
func (this *Zval) SetCacheSlot(value uint32)     { this.u2.cache_slot = value }
func (this Zval) GetOplineNum() uint32           { return this.u2.opline_num }
func (this *Zval) SetOplineNum(value uint32)     { this.u2.opline_num = value }
func (this Zval) GetLineno() uint32              { return this.u2.lineno }
func (this *Zval) SetLineno(value uint32)        { this.u2.lineno = value }
func (this Zval) GetNumArgs() uint32             { return this.u2.num_args }
func (this *Zval) SetNumArgs(value uint32)       { this.u2.num_args = value }
func (this Zval) GetFePos() uint32               { return this.u2.fe_pos }
func (this *Zval) SetFePos(value uint32)         { this.u2.fe_pos = value }
func (this Zval) GetFeIterIdx() uint32           { return this.u2.fe_iter_idx }
func (this *Zval) SetFeIterIdx(value uint32)     { this.u2.fe_iter_idx = value }
func (this Zval) GetAccessFlags() uint32         { return this.u2.access_flags }
func (this *Zval) SetAccessFlags(value uint32)   { this.u2.access_flags = value }
func (this Zval) GetPropertyGuard() uint32       { return this.u2.property_guard }
func (this *Zval) SetPropertyGuard(value uint32) { this.u2.property_guard = value }
func (this Zval) GetConstantFlags() uint32       { return this.u2.constant_flags }
func (this *Zval) SetConstantFlags(value uint32) { this.u2.constant_flags = value }
func (this Zval) GetU2Extra() uint32             { return this.u2.extra }
func (this *Zval) SetU2Extra(value uint32)       { this.u2.extra = value }
