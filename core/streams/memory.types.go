package streams

import (
	"sik/core"
	"sik/zend/types"
)

/**
 * PhpStreamMemoryData
 */
type PhpStreamMemoryData struct {
	data  *byte
	fpos  int
	fsize int
	smax  int
	mode  int
}

// func MakePhpStreamMemoryData(data *byte, fpos int, fsize int, smax int, mode int) PhpStreamMemoryData {
//     return PhpStreamMemoryData{
//         data:data,
//         fpos:fpos,
//         fsize:fsize,
//         smax:smax,
//         mode:mode,
//     }
// }
func (this *PhpStreamMemoryData) GetData() *byte      { return this.data }
func (this *PhpStreamMemoryData) SetData(value *byte) { this.data = value }
func (this *PhpStreamMemoryData) GetFpos() int        { return this.fpos }
func (this *PhpStreamMemoryData) SetFpos(value int)   { this.fpos = value }
func (this *PhpStreamMemoryData) GetFsize() int       { return this.fsize }
func (this *PhpStreamMemoryData) SetFsize(value int)  { this.fsize = value }

// func (this *PhpStreamMemoryData)  GetSmax() int      { return this.smax }
func (this *PhpStreamMemoryData) SetSmax(value int) { this.smax = value }
func (this *PhpStreamMemoryData) GetMode() int      { return this.mode }
func (this *PhpStreamMemoryData) SetMode(value int) { this.mode = value }

/**
 * PhpStreamTempData
 */
type PhpStreamTempData struct {
	innerstream *core.PhpStream
	smax        int
	mode        int
	meta        types.Zval
	tmpdir      *byte
}

// func MakePhpStreamTempData(innerstream *core.PhpStream, smax int, mode int, meta zend.Zval, tmpdir *byte) PhpStreamTempData {
//     return PhpStreamTempData{
//         innerstream:innerstream,
//         smax:smax,
//         mode:mode,
//         meta:meta,
//         tmpdir:tmpdir,
//     }
// }
func (this *PhpStreamTempData) GetInnerstream() *core.PhpStream      { return this.innerstream }
func (this *PhpStreamTempData) SetInnerstream(value *core.PhpStream) { this.innerstream = value }
func (this *PhpStreamTempData) GetSmax() int                         { return this.smax }
func (this *PhpStreamTempData) SetSmax(value int)                    { this.smax = value }

// func (this *PhpStreamTempData)  GetMode() int      { return this.mode }
func (this *PhpStreamTempData) SetMode(value int)   { this.mode = value }
func (this *PhpStreamTempData) GetMeta() types.Zval { return this.meta }

// func (this *PhpStreamTempData) SetMeta(value zend.Zval) { this.meta = value }
func (this *PhpStreamTempData) GetTmpdir() *byte      { return this.tmpdir }
func (this *PhpStreamTempData) SetTmpdir(value *byte) { this.tmpdir = value }
