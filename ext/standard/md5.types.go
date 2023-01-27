// <<generate>>

package standard

/**
 * PHP_MD5_CTX
 */
type PHP_MD5_CTX struct {
	lo     uint32
	hi     uint32
	a      uint32
	b      uint32
	c      uint32
	d      uint32
	buffer []uint8
	block  []uint32
}

func (this *PHP_MD5_CTX) GetLo() uint32           { return this.lo }
func (this *PHP_MD5_CTX) SetLo(value uint32)      { this.lo = value }
func (this *PHP_MD5_CTX) GetHi() uint32           { return this.hi }
func (this *PHP_MD5_CTX) SetHi(value uint32)      { this.hi = value }
func (this *PHP_MD5_CTX) GetA() uint32            { return this.a }
func (this *PHP_MD5_CTX) SetA(value uint32)       { this.a = value }
func (this *PHP_MD5_CTX) GetB() uint32            { return this.b }
func (this *PHP_MD5_CTX) SetB(value uint32)       { this.b = value }
func (this *PHP_MD5_CTX) GetC() uint32            { return this.c }
func (this *PHP_MD5_CTX) SetC(value uint32)       { this.c = value }
func (this *PHP_MD5_CTX) GetD() uint32            { return this.d }
func (this *PHP_MD5_CTX) SetD(value uint32)       { this.d = value }
func (this *PHP_MD5_CTX) GetBuffer() []uint8      { return this.buffer }
func (this *PHP_MD5_CTX) SetBuffer(value []uint8) { this.buffer = value }
func (this *PHP_MD5_CTX) GetBlock() []uint32      { return this.block }
func (this *PHP_MD5_CTX) SetBlock(value []uint32) { this.block = value }
