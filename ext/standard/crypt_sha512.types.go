package standard

/**
 * Sha512Ctx
 */
type Sha512Ctx struct {
	H      []uint64
	total  []uint64
	buflen uint64
	buffer []byte
}

// func MakeSha512Ctx(H []uint64, total []uint64, buflen uint64, buffer []byte) Sha512Ctx {
//     return Sha512Ctx{
//         H:H,
//         total:total,
//         buflen:buflen,
//         buffer:buffer,
//     }
// }
func (this *Sha512Ctx) GetH() []uint64 { return this.H }

// func (this *Sha512Ctx) SetH(value []uint64) { this.H = value }
func (this *Sha512Ctx) GetTotal() []uint64 { return this.total }

// func (this *Sha512Ctx) SetTotal(value []uint64) { this.total = value }
func (this *Sha512Ctx) GetBuflen() uint64      { return this.buflen }
func (this *Sha512Ctx) SetBuflen(value uint64) { this.buflen = value }
func (this *Sha512Ctx) GetBuffer() []byte      { return this.buffer }

// func (this *Sha512Ctx) SetBuffer(value []byte) { this.buffer = value }
