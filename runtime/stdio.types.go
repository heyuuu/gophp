// <<generate>>

package runtime

/**
 * FILE
 */
type FILE struct {
	_p       *uint8
	_r       int
	_w       int
	_flags   short
	_file    short
	_bf      __struct____sbuf
	_lbfsize int
	_cookie  any
	_close   func(any) int
	_read    func(any, *byte, int) int
	_seek    func(any, fpos_t, int) fpos_t
	_write   func(any, *byte, int) int
	_ub      __struct____sbuf
	_extra   *__struct____sFILEX
	_ur      int
	_ubuf    []uint8
	_nbuf    []uint8
	_lb      __struct____sbuf
	_blksize int
	_offset  fpos_t
}

// func NewFILE(_p *uint8, _r int, _w int, _flags short, _file short, _bf __struct____sbuf, _lbfsize int, _cookie any, _close func(any) int, _read func(any, *byte, int) int, _seek func(any, fpos_t, int) fpos_t, _write func(any, *byte, int) int, _ub __struct____sbuf, _extra *__struct____sFILEX, _ur int, _ubuf []uint8, _nbuf []uint8, _lb __struct____sbuf, _blksize int, _offset fpos_t) *FILE {
//     return &FILE{
//         _p:_p,
//         _r:_r,
//         _w:_w,
//         _flags:_flags,
//         _file:_file,
//         _bf:_bf,
//         _lbfsize:_lbfsize,
//         _cookie:_cookie,
//         _close:_close,
//         _read:_read,
//         _seek:_seek,
//         _write:_write,
//         _ub:_ub,
//         _extra:_extra,
//         _ur:_ur,
//         _ubuf:_ubuf,
//         _nbuf:_nbuf,
//         _lb:_lb,
//         _blksize:_blksize,
//         _offset:_offset,
//     }
// }
// func MakeFILE(_p *uint8, _r int, _w int, _flags short, _file short, _bf __struct____sbuf, _lbfsize int, _cookie any, _close func(any) int, _read func(any, *byte, int) int, _seek func(any, fpos_t, int) fpos_t, _write func(any, *byte, int) int, _ub __struct____sbuf, _extra *__struct____sFILEX, _ur int, _ubuf []uint8, _nbuf []uint8, _lb __struct____sbuf, _blksize int, _offset fpos_t) FILE {
//     return FILE{
//         _p:_p,
//         _r:_r,
//         _w:_w,
//         _flags:_flags,
//         _file:_file,
//         _bf:_bf,
//         _lbfsize:_lbfsize,
//         _cookie:_cookie,
//         _close:_close,
//         _read:_read,
//         _seek:_seek,
//         _write:_write,
//         _ub:_ub,
//         _extra:_extra,
//         _ur:_ur,
//         _ubuf:_ubuf,
//         _nbuf:_nbuf,
//         _lb:_lb,
//         _blksize:_blksize,
//         _offset:_offset,
//     }
// }
// func (this *FILE)  GetP() *uint8      { return this._p }
// func (this *FILE) SetP(value *uint8) { this._p = value }
// func (this *FILE)  GetR() int      { return this._r }
// func (this *FILE) SetR(value int) { this._r = value }
// func (this *FILE)  GetW() int      { return this._w }
// func (this *FILE) SetW(value int) { this._w = value }
// func (this *FILE)  GetFlags() short      { return this._flags }
// func (this *FILE) SetFlags(value short) { this._flags = value }
// func (this *FILE)  GetFile() short      { return this._file }
// func (this *FILE) SetFile(value short) { this._file = value }
// func (this *FILE)  GetBf() __struct____sbuf      { return this._bf }
// func (this *FILE) SetBf(value __struct____sbuf) { this._bf = value }
// func (this *FILE)  GetLbfsize() int      { return this._lbfsize }
// func (this *FILE) SetLbfsize(value int) { this._lbfsize = value }
// func (this *FILE)  GetCookie() any      { return this._cookie }
// func (this *FILE) SetCookie(value any) { this._cookie = value }
// func (this *FILE)  GetClose() func(any) int      { return this._close }
// func (this *FILE) SetClose(value func(any) int) { this._close = value }
// func (this *FILE)  GetRead() func(any, *byte, int) int      { return this._read }
// func (this *FILE) SetRead(value func(any, *byte, int) int) { this._read = value }
// func (this *FILE)  GetSeek() func(any, fpos_t, int) fpos_t      { return this._seek }
// func (this *FILE) SetSeek(value func(any, fpos_t, int) fpos_t) { this._seek = value }
// func (this *FILE)  GetWrite() func(any, *byte, int) int      { return this._write }
// func (this *FILE) SetWrite(value func(any, *byte, int) int) { this._write = value }
// func (this *FILE)  GetUb() __struct____sbuf      { return this._ub }
// func (this *FILE) SetUb(value __struct____sbuf) { this._ub = value }
// func (this *FILE)  GetExtra() *__struct____sFILEX      { return this._extra }
// func (this *FILE) SetExtra(value *__struct____sFILEX) { this._extra = value }
// func (this *FILE)  GetUr() int      { return this._ur }
// func (this *FILE) SetUr(value int) { this._ur = value }
// func (this *FILE)  GetUbuf() []uint8      { return this._ubuf }
// func (this *FILE) SetUbuf(value []uint8) { this._ubuf = value }
// func (this *FILE)  GetNbuf() []uint8      { return this._nbuf }
// func (this *FILE) SetNbuf(value []uint8) { this._nbuf = value }
// func (this *FILE)  GetLb() __struct____sbuf      { return this._lb }
// func (this *FILE) SetLb(value __struct____sbuf) { this._lb = value }
// func (this *FILE)  GetBlksize() int      { return this._blksize }
// func (this *FILE) SetBlksize(value int) { this._blksize = value }
// func (this *FILE)  GetOffset() fpos_t      { return this._offset }
// func (this *FILE) SetOffset(value fpos_t) { this._offset = value }
