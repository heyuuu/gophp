// <<generate>>

package zend

/**
 * ZendCpuInfo
 */
type ZendCpuInfo struct {
	eax         uint32
	ebx         uint32
	ecx         uint32
	edx         uint32
	initialized uint32
}

func (this *ZendCpuInfo) GetEax() uint32              { return this.eax }
func (this *ZendCpuInfo) SetEax(value uint32)         { this.eax = value }
func (this *ZendCpuInfo) GetEbx() uint32              { return this.ebx }
func (this *ZendCpuInfo) SetEbx(value uint32)         { this.ebx = value }
func (this *ZendCpuInfo) GetEcx() uint32              { return this.ecx }
func (this *ZendCpuInfo) SetEcx(value uint32)         { this.ecx = value }
func (this *ZendCpuInfo) GetEdx() uint32              { return this.edx }
func (this *ZendCpuInfo) SetEdx(value uint32)         { this.edx = value }
func (this *ZendCpuInfo) GetInitialized() uint32      { return this.initialized }
func (this *ZendCpuInfo) SetInitialized(value uint32) { this.initialized = value }
