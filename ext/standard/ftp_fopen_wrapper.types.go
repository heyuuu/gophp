// <<generate>>

package standard

import (
	"sik/core"
)

/**
 * PhpFtpDirstreamData
 */
type PhpFtpDirstreamData struct {
	datastream    *core.PhpStream
	controlstream *core.PhpStream
	dirstream     *core.PhpStream
}

// func NewPhpFtpDirstreamData(datastream *core.PhpStream, controlstream *core.PhpStream, dirstream *core.PhpStream) *PhpFtpDirstreamData {
//     return &PhpFtpDirstreamData{
//         datastream:datastream,
//         controlstream:controlstream,
//         dirstream:dirstream,
//     }
// }
// func MakePhpFtpDirstreamData(datastream *core.PhpStream, controlstream *core.PhpStream, dirstream *core.PhpStream) PhpFtpDirstreamData {
//     return PhpFtpDirstreamData{
//         datastream:datastream,
//         controlstream:controlstream,
//         dirstream:dirstream,
//     }
// }
func (this *PhpFtpDirstreamData) GetDatastream() *core.PhpStream         { return this.datastream }
func (this *PhpFtpDirstreamData) SetDatastream(value *core.PhpStream)    { this.datastream = value }
func (this *PhpFtpDirstreamData) GetControlstream() *core.PhpStream      { return this.controlstream }
func (this *PhpFtpDirstreamData) SetControlstream(value *core.PhpStream) { this.controlstream = value }
func (this *PhpFtpDirstreamData) GetDirstream() *core.PhpStream          { return this.dirstream }
func (this *PhpFtpDirstreamData) SetDirstream(value *core.PhpStream)     { this.dirstream = value }
