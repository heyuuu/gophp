// <<generate>>

package standard

import (
	"sik/zend/types"
)

/**
 * Bucketindex
 */
type Bucketindex struct {
	b types.Bucket
	i uint
}

// func MakeBucketindex(b zend.Bucket, i uint) Bucketindex {
//     return Bucketindex{
//         b:b,
//         i:i,
//     }
// }
func (this *Bucketindex) GetB() types.Bucket      { return this.b }
func (this *Bucketindex) SetB(value types.Bucket) { this.b = value }
func (this *Bucketindex) GetI() uint              { return this.i }
func (this *Bucketindex) SetI(value uint)         { this.i = value }
