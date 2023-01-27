// <<generate>>

package standard

/**
 * Gfxinfo
 */
type Gfxinfo struct {
	width    uint
	height   uint
	bits     uint
	channels uint
}

func (this *Gfxinfo) GetWidth() uint         { return this.width }
func (this *Gfxinfo) SetWidth(value uint)    { this.width = value }
func (this *Gfxinfo) GetHeight() uint        { return this.height }
func (this *Gfxinfo) SetHeight(value uint)   { this.height = value }
func (this *Gfxinfo) GetBits() uint          { return this.bits }
func (this *Gfxinfo) SetBits(value uint)     { this.bits = value }
func (this *Gfxinfo) GetChannels() uint      { return this.channels }
func (this *Gfxinfo) SetChannels(value uint) { this.channels = value }
