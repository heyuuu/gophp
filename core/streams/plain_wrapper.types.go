// <<generate>>

package streams

import (
	r "sik/runtime"
	"sik/zend"
)

/**
 * PhpStdioStreamData
 */
type PhpStdioStreamData struct {
	file             *r.FILE
	fd               int
	is_process_pipe  unsigned
	is_pipe          unsigned
	cached_fstat     unsigned
	is_pipe_blocking unsigned
	no_forced_fstat  unsigned
	is_seekable      unsigned
	_reserved        unsigned
	lock_flag        int
	temp_name        *zend.ZendString
	last_op          byte
	last_mapped_addr *byte
	last_mapped_len  int
	sb               zend.ZendStatT
}

func (this *PhpStdioStreamData) GetFile() *r.FILE                   { return this.file }
func (this *PhpStdioStreamData) SetFile(value *r.FILE)              { this.file = value }
func (this *PhpStdioStreamData) GetFd() int                         { return this.fd }
func (this *PhpStdioStreamData) SetFd(value int)                    { this.fd = value }
func (this *PhpStdioStreamData) GetIsProcessPipe() unsigned         { return this.is_process_pipe }
func (this *PhpStdioStreamData) SetIsProcessPipe(value unsigned)    { this.is_process_pipe = value }
func (this *PhpStdioStreamData) GetIsPipe() unsigned                { return this.is_pipe }
func (this *PhpStdioStreamData) SetIsPipe(value unsigned)           { this.is_pipe = value }
func (this *PhpStdioStreamData) GetCachedFstat() unsigned           { return this.cached_fstat }
func (this *PhpStdioStreamData) SetCachedFstat(value unsigned)      { this.cached_fstat = value }
func (this *PhpStdioStreamData) GetIsPipeBlocking() unsigned        { return this.is_pipe_blocking }
func (this *PhpStdioStreamData) SetIsPipeBlocking(value unsigned)   { this.is_pipe_blocking = value }
func (this *PhpStdioStreamData) GetNoForcedFstat() unsigned         { return this.no_forced_fstat }
func (this *PhpStdioStreamData) SetNoForcedFstat(value unsigned)    { this.no_forced_fstat = value }
func (this *PhpStdioStreamData) GetIsSeekable() unsigned            { return this.is_seekable }
func (this *PhpStdioStreamData) SetIsSeekable(value unsigned)       { this.is_seekable = value }
func (this *PhpStdioStreamData) GetReserved() unsigned              { return this._reserved }
func (this *PhpStdioStreamData) SetReserved(value unsigned)         { this._reserved = value }
func (this *PhpStdioStreamData) GetLockFlag() int                   { return this.lock_flag }
func (this *PhpStdioStreamData) SetLockFlag(value int)              { this.lock_flag = value }
func (this *PhpStdioStreamData) GetTempName() *zend.ZendString      { return this.temp_name }
func (this *PhpStdioStreamData) SetTempName(value *zend.ZendString) { this.temp_name = value }
func (this *PhpStdioStreamData) GetLastOp() byte                    { return this.last_op }
func (this *PhpStdioStreamData) SetLastOp(value byte)               { this.last_op = value }
func (this *PhpStdioStreamData) GetLastMappedAddr() *byte           { return this.last_mapped_addr }
func (this *PhpStdioStreamData) SetLastMappedAddr(value *byte)      { this.last_mapped_addr = value }
func (this *PhpStdioStreamData) GetLastMappedLen() int              { return this.last_mapped_len }
func (this *PhpStdioStreamData) SetLastMappedLen(value int)         { this.last_mapped_len = value }
func (this *PhpStdioStreamData) GetSb() zend.ZendStatT              { return this.sb }
func (this *PhpStdioStreamData) SetSb(value zend.ZendStatT)         { this.sb = value }
