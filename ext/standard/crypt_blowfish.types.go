// <<generate>>

package standard

/**
 * BF_ctx
 */
type BF_ctx struct {
	S [][]BF_word
	P BF_key
}

// func NewBF_ctx(S [][]BF_word, P BF_key) *BF_ctx {
//     return &BF_ctx{
//         S:S,
//         P:P,
//     }
// }
func MakeBF_ctx(S [][]BF_word, P BF_key) BF_ctx {
	return BF_ctx{
		S: S,
		P: P,
	}
}
func (this *BF_ctx) GetS() [][]BF_word { return this.S }

// func (this *BF_ctx) SetS(value [][]BF_word) { this.S = value }
func (this *BF_ctx) GetP() BF_key { return this.P }

// func (this *BF_ctx) SetP(value BF_key) { this.P = value }
