// <<generate>>

package zend

/**
 * ZendRefcountedH
 */
type ZendRefcountedH struct {
	refcount uint32
	u        struct /* union */ {
		type_info uint32
	}
}

func (this *ZendRefcountedH) GetRefcount() uint32      { return this.refcount }
func (this *ZendRefcountedH) SetRefcount(value uint32) { this.refcount = value }
func (this *ZendRefcountedH) GetTypeInfo() uint32      { return this.u.type_info }
func (this *ZendRefcountedH) SetTypeInfo(value uint32) { this.u.type_info = value }

/**
 * ZendRefcounted
 */
type ZendRefcounted struct {
	gc ZendRefcountedH
}

func (this *ZendRefcounted) GetGc() ZendRefcountedH      { return this.gc }
func (this *ZendRefcounted) SetGc(value ZendRefcountedH) { this.gc = value }
