// <<generate>>

package standard

/**
 * PhpRandomGlobals
 */
type PhpRandomGlobals struct {
	fd int
}

func (this *PhpRandomGlobals) GetFd() int      { return this.fd }
func (this *PhpRandomGlobals) SetFd(value int) { this.fd = value }
