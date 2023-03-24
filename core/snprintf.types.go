package core

/**
 * BufArea
 */
type BufArea struct {
	buf_end *byte
	nextb   *byte
}

func (this *BufArea) GetBufEnd() *byte      { return this.buf_end }
func (this *BufArea) SetBufEnd(value *byte) { this.buf_end = value }
func (this *BufArea) GetNextb() *byte       { return this.nextb }
func (this *BufArea) SetNextb(value *byte)  { this.nextb = value }
