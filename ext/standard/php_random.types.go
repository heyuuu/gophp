package standard

/**
 * PhpRandomGlobals
 */
type PhpRandomGlobals struct {
	fd int
}

// func MakePhpRandomGlobals(fd int) PhpRandomGlobals {
//     return PhpRandomGlobals{
//         fd:fd,
//     }
// }
func (this *PhpRandomGlobals) GetFd() int      { return this.fd }
func (this *PhpRandomGlobals) SetFd(value int) { this.fd = value }
