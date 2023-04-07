package standard

/**
 * Sha256Ctx
 */
type Sha256Ctx struct {
	H      []uint32
	total  []uint32
	buflen uint32
	buffer []byte
}

func (this *Sha256Ctx) GetH() []uint32 { return this.H }

// func (this *Sha256Ctx) SetH(value []uint32) { this.H = value }
func (this *Sha256Ctx) GetTotal() []uint32 { return this.total }

// func (this *Sha256Ctx) SetTotal(value []uint32) { this.total = value }
func (this *Sha256Ctx) GetBuflen() uint32      { return this.buflen }
func (this *Sha256Ctx) SetBuflen(value uint32) { this.buflen = value }
func (this *Sha256Ctx) GetBuffer() []byte      { return this.buffer }

// func (this *Sha256Ctx) SetBuffer(value []byte) { this.buffer = value }
