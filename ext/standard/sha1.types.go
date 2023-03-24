package standard

/**
 * PHP_SHA1_CTX
 */
type PHP_SHA1_CTX struct {
	state  []uint32
	count  []uint32
	buffer []uint8
}

// func MakePHP_SHA1_CTX(state []uint32, count []uint32, buffer []uint8) PHP_SHA1_CTX {
//     return PHP_SHA1_CTX{
//         state:state,
//         count:count,
//         buffer:buffer,
//     }
// }
func (this *PHP_SHA1_CTX) GetState() []uint32 { return this.state }

// func (this *PHP_SHA1_CTX) SetState(value []uint32) { this.state = value }
func (this *PHP_SHA1_CTX) GetCount() []uint32 { return this.count }

// func (this *PHP_SHA1_CTX) SetCount(value []uint32) { this.count = value }
func (this *PHP_SHA1_CTX) GetBuffer() []uint8 { return this.buffer }

// func (this *PHP_SHA1_CTX) SetBuffer(value []uint8) { this.buffer = value }
