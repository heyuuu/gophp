// <<generate>>

package zend

/**
 * U
 */
type U struct /* union */ {
	d float64
	L []ULong
}

func (this U) GetD() float64       { return this.d }
func (this *U) SetD(value float64) { this.d = value }
func (this U) GetL() []ULong       { return this.L }
func (this *U) SetL(value []ULong) { this.L = value }

/**
 * BCinfo
 */
type BCinfo struct {
	dp0      int
	dp1      int
	dplen    int
	dsign    int
	e0       int
	inexact  int
	nd       int
	nd0      int
	rounding int
	scale    int
	uflchk   int
}

func (this BCinfo) GetDp0() int            { return this.dp0 }
func (this *BCinfo) SetDp0(value int)      { this.dp0 = value }
func (this BCinfo) GetDp1() int            { return this.dp1 }
func (this *BCinfo) SetDp1(value int)      { this.dp1 = value }
func (this BCinfo) GetDplen() int          { return this.dplen }
func (this *BCinfo) SetDplen(value int)    { this.dplen = value }
func (this BCinfo) GetDsign() int          { return this.dsign }
func (this *BCinfo) SetDsign(value int)    { this.dsign = value }
func (this BCinfo) GetE0() int             { return this.e0 }
func (this *BCinfo) SetE0(value int)       { this.e0 = value }
func (this BCinfo) GetInexact() int        { return this.inexact }
func (this *BCinfo) SetInexact(value int)  { this.inexact = value }
func (this BCinfo) GetNd() int             { return this.nd }
func (this *BCinfo) SetNd(value int)       { this.nd = value }
func (this BCinfo) GetNd0() int            { return this.nd0 }
func (this *BCinfo) SetNd0(value int)      { this.nd0 = value }
func (this BCinfo) GetRounding() int       { return this.rounding }
func (this *BCinfo) SetRounding(value int) { this.rounding = value }
func (this BCinfo) GetScale() int          { return this.scale }
func (this *BCinfo) SetScale(value int)    { this.scale = value }
func (this BCinfo) GetUflchk() int         { return this.uflchk }
func (this *BCinfo) SetUflchk(value int)   { this.uflchk = value }

/**
 * Bigint
 */
type Bigint struct {
	next   *Bigint
	k      int
	maxwds int
	sign   int
	wds    int
	x      []ULong
}

func (this Bigint) GetNext() *Bigint       { return this.next }
func (this *Bigint) SetNext(value *Bigint) { this.next = value }
func (this Bigint) GetK() int              { return this.k }
func (this *Bigint) SetK(value int)        { this.k = value }
func (this Bigint) GetMaxwds() int         { return this.maxwds }
func (this *Bigint) SetMaxwds(value int)   { this.maxwds = value }
func (this Bigint) GetSign() int           { return this.sign }
func (this *Bigint) SetSign(value int)     { this.sign = value }
func (this Bigint) GetWds() int            { return this.wds }
func (this *Bigint) SetWds(value int)      { this.wds = value }
func (this Bigint) GetX() []ULong          { return this.x }
func (this *Bigint) SetX(value []ULong)    { this.x = value }
