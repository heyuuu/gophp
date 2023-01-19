// <<generate>>

package standard

import (
	"sik/zend"
)

/**
 * Bucketindex
 */
type Bucketindex struct {
	b zend.Bucket
	i uint
}

func (this Bucketindex) GetB() zend.Bucket       { return this.b }
func (this *Bucketindex) SetB(value zend.Bucket) { this.b = value }
func (this Bucketindex) GetI() uint              { return this.i }
func (this *Bucketindex) SetI(value uint)        { this.i = value }
