// <<generate>>

package core

import (
	"sik/zend"
)

/**
 * MultipartEventStart
 */
type MultipartEventStart struct {
	content_length int
}

// func NewMultipartEventStart(content_length int) *MultipartEventStart {
//     return &MultipartEventStart{
//         content_length:content_length,
//     }
// }
// func MakeMultipartEventStart(content_length int) MultipartEventStart {
//     return MultipartEventStart{
//         content_length:content_length,
//     }
// }
// func (this *MultipartEventStart)  GetContentLength() int      { return this.content_length }
func (this *MultipartEventStart) SetContentLength(value int) { this.content_length = value }

/**
 * MultipartEventFormdata
 */
type MultipartEventFormdata struct {
	post_bytes_processed int
	name                 *byte
	value                **byte
	length               int
	newlength            *int
}

// func NewMultipartEventFormdata(post_bytes_processed int, name *byte, value **byte, length int, newlength *int) *MultipartEventFormdata {
//     return &MultipartEventFormdata{
//         post_bytes_processed:post_bytes_processed,
//         name:name,
//         value:value,
//         length:length,
//         newlength:newlength,
//     }
// }
// func MakeMultipartEventFormdata(post_bytes_processed int, name *byte, value **byte, length int, newlength *int) MultipartEventFormdata {
//     return MultipartEventFormdata{
//         post_bytes_processed:post_bytes_processed,
//         name:name,
//         value:value,
//         length:length,
//         newlength:newlength,
//     }
// }
// func (this *MultipartEventFormdata)  GetPostBytesProcessed() int      { return this.post_bytes_processed }
func (this *MultipartEventFormdata) SetPostBytesProcessed(value int) {
	this.post_bytes_processed = value
}

// func (this *MultipartEventFormdata)  GetName() *byte      { return this.name }
func (this *MultipartEventFormdata) SetName(value *byte) { this.name = value }

// func (this *MultipartEventFormdata)  GetValue() **byte      { return this.value }
func (this *MultipartEventFormdata) SetValue(value **byte) { this.value = value }

// func (this *MultipartEventFormdata)  GetLength() int      { return this.length }
func (this *MultipartEventFormdata) SetLength(value int) { this.length = value }

// func (this *MultipartEventFormdata)  GetNewlength() *int      { return this.newlength }
func (this *MultipartEventFormdata) SetNewlength(value *int) { this.newlength = value }

/**
 * MultipartEventFileStart
 */
type MultipartEventFileStart struct {
	post_bytes_processed int
	name                 *byte
	filename             **byte
}

// func NewMultipartEventFileStart(post_bytes_processed int, name *byte, filename **byte) *MultipartEventFileStart {
//     return &MultipartEventFileStart{
//         post_bytes_processed:post_bytes_processed,
//         name:name,
//         filename:filename,
//     }
// }
// func MakeMultipartEventFileStart(post_bytes_processed int, name *byte, filename **byte) MultipartEventFileStart {
//     return MultipartEventFileStart{
//         post_bytes_processed:post_bytes_processed,
//         name:name,
//         filename:filename,
//     }
// }
// func (this *MultipartEventFileStart)  GetPostBytesProcessed() int      { return this.post_bytes_processed }
func (this *MultipartEventFileStart) SetPostBytesProcessed(value int) {
	this.post_bytes_processed = value
}

// func (this *MultipartEventFileStart)  GetName() *byte      { return this.name }
func (this *MultipartEventFileStart) SetName(value *byte) { this.name = value }

// func (this *MultipartEventFileStart)  GetFilename() **byte      { return this.filename }
func (this *MultipartEventFileStart) SetFilename(value **byte) { this.filename = value }

/**
 * MultipartEventFileData
 */
type MultipartEventFileData struct {
	post_bytes_processed int
	offset               zend.ZendOffT
	data                 *byte
	length               int
	newlength            *int
}

// func NewMultipartEventFileData(post_bytes_processed int, offset zend.ZendOffT, data *byte, length int, newlength *int) *MultipartEventFileData {
//     return &MultipartEventFileData{
//         post_bytes_processed:post_bytes_processed,
//         offset:offset,
//         data:data,
//         length:length,
//         newlength:newlength,
//     }
// }
// func MakeMultipartEventFileData(post_bytes_processed int, offset zend.ZendOffT, data *byte, length int, newlength *int) MultipartEventFileData {
//     return MultipartEventFileData{
//         post_bytes_processed:post_bytes_processed,
//         offset:offset,
//         data:data,
//         length:length,
//         newlength:newlength,
//     }
// }
// func (this *MultipartEventFileData)  GetPostBytesProcessed() int      { return this.post_bytes_processed }
func (this *MultipartEventFileData) SetPostBytesProcessed(value int) {
	this.post_bytes_processed = value
}

// func (this *MultipartEventFileData)  GetOffset() zend.ZendOffT      { return this.offset }
func (this *MultipartEventFileData) SetOffset(value zend.ZendOffT) { this.offset = value }

// func (this *MultipartEventFileData)  GetData() *byte      { return this.data }
func (this *MultipartEventFileData) SetData(value *byte) { this.data = value }

// func (this *MultipartEventFileData)  GetLength() int      { return this.length }
func (this *MultipartEventFileData) SetLength(value int) { this.length = value }

// func (this *MultipartEventFileData)  GetNewlength() *int      { return this.newlength }
func (this *MultipartEventFileData) SetNewlength(value *int) { this.newlength = value }

/**
 * MultipartEventFileEnd
 */
type MultipartEventFileEnd struct {
	post_bytes_processed int
	temp_filename        *byte
	cancel_upload        int
}

// func NewMultipartEventFileEnd(post_bytes_processed int, temp_filename *byte, cancel_upload int) *MultipartEventFileEnd {
//     return &MultipartEventFileEnd{
//         post_bytes_processed:post_bytes_processed,
//         temp_filename:temp_filename,
//         cancel_upload:cancel_upload,
//     }
// }
// func MakeMultipartEventFileEnd(post_bytes_processed int, temp_filename *byte, cancel_upload int) MultipartEventFileEnd {
//     return MultipartEventFileEnd{
//         post_bytes_processed:post_bytes_processed,
//         temp_filename:temp_filename,
//         cancel_upload:cancel_upload,
//     }
// }
// func (this *MultipartEventFileEnd)  GetPostBytesProcessed() int      { return this.post_bytes_processed }
func (this *MultipartEventFileEnd) SetPostBytesProcessed(value int) {
	this.post_bytes_processed = value
}

// func (this *MultipartEventFileEnd)  GetTempFilename() *byte      { return this.temp_filename }
func (this *MultipartEventFileEnd) SetTempFilename(value *byte) { this.temp_filename = value }

// func (this *MultipartEventFileEnd)  GetCancelUpload() int      { return this.cancel_upload }
func (this *MultipartEventFileEnd) SetCancelUpload(value int) { this.cancel_upload = value }

/**
 * MultipartEventEnd
 */
type MultipartEventEnd struct {
	post_bytes_processed int
}

// func NewMultipartEventEnd(post_bytes_processed int) *MultipartEventEnd {
//     return &MultipartEventEnd{
//         post_bytes_processed:post_bytes_processed,
//     }
// }
// func MakeMultipartEventEnd(post_bytes_processed int) MultipartEventEnd {
//     return MultipartEventEnd{
//         post_bytes_processed:post_bytes_processed,
//     }
// }
// func (this *MultipartEventEnd)  GetPostBytesProcessed() int      { return this.post_bytes_processed }
func (this *MultipartEventEnd) SetPostBytesProcessed(value int) { this.post_bytes_processed = value }

/**
 * MultipartBuffer
 */
type MultipartBuffer struct {
	buffer            *byte
	buf_begin         *byte
	bufsize           int
	bytes_in_buffer   int
	boundary          *byte
	boundary_next     *byte
	boundary_next_len int
	input_encoding    *zend.ZendEncoding
	detect_order      **zend.ZendEncoding
	detect_order_size int
}

// func NewMultipartBuffer(buffer *byte, buf_begin *byte, bufsize int, bytes_in_buffer int, boundary *byte, boundary_next *byte, boundary_next_len int, input_encoding *zend.ZendEncoding, detect_order **zend.ZendEncoding, detect_order_size int) *MultipartBuffer {
//     return &MultipartBuffer{
//         buffer:buffer,
//         buf_begin:buf_begin,
//         bufsize:bufsize,
//         bytes_in_buffer:bytes_in_buffer,
//         boundary:boundary,
//         boundary_next:boundary_next,
//         boundary_next_len:boundary_next_len,
//         input_encoding:input_encoding,
//         detect_order:detect_order,
//         detect_order_size:detect_order_size,
//     }
// }
// func MakeMultipartBuffer(buffer *byte, buf_begin *byte, bufsize int, bytes_in_buffer int, boundary *byte, boundary_next *byte, boundary_next_len int, input_encoding *zend.ZendEncoding, detect_order **zend.ZendEncoding, detect_order_size int) MultipartBuffer {
//     return MultipartBuffer{
//         buffer:buffer,
//         buf_begin:buf_begin,
//         bufsize:bufsize,
//         bytes_in_buffer:bytes_in_buffer,
//         boundary:boundary,
//         boundary_next:boundary_next,
//         boundary_next_len:boundary_next_len,
//         input_encoding:input_encoding,
//         detect_order:detect_order,
//         detect_order_size:detect_order_size,
//     }
// }
func (this *MultipartBuffer) GetBuffer() *byte           { return this.buffer }
func (this *MultipartBuffer) SetBuffer(value *byte)      { this.buffer = value }
func (this *MultipartBuffer) GetBufBegin() *byte         { return this.buf_begin }
func (this *MultipartBuffer) SetBufBegin(value *byte)    { this.buf_begin = value }
func (this *MultipartBuffer) GetBufsize() int            { return this.bufsize }
func (this *MultipartBuffer) SetBufsize(value int)       { this.bufsize = value }
func (this *MultipartBuffer) GetBytesInBuffer() int      { return this.bytes_in_buffer }
func (this *MultipartBuffer) SetBytesInBuffer(value int) { this.bytes_in_buffer = value }
func (this *MultipartBuffer) GetBoundary() *byte         { return this.boundary }

// func (this *MultipartBuffer) SetBoundary(value *byte) { this.boundary = value }
func (this *MultipartBuffer) GetBoundaryNext() *byte { return this.boundary_next }

// func (this *MultipartBuffer) SetBoundaryNext(value *byte) { this.boundary_next = value }
func (this *MultipartBuffer) GetBoundaryNextLen() int                   { return this.boundary_next_len }
func (this *MultipartBuffer) SetBoundaryNextLen(value int)              { this.boundary_next_len = value }
func (this *MultipartBuffer) GetInputEncoding() *zend.ZendEncoding      { return this.input_encoding }
func (this *MultipartBuffer) SetInputEncoding(value *zend.ZendEncoding) { this.input_encoding = value }
func (this *MultipartBuffer) GetDetectOrder() **zend.ZendEncoding       { return this.detect_order }
func (this *MultipartBuffer) SetDetectOrder(value **zend.ZendEncoding)  { this.detect_order = value }
func (this *MultipartBuffer) GetDetectOrderSize() int                   { return this.detect_order_size }
func (this *MultipartBuffer) SetDetectOrderSize(value int)              { this.detect_order_size = value }

/**
 * MimeHeaderEntry
 */
type MimeHeaderEntry struct {
	key   *byte
	value *byte
}

// func NewMimeHeaderEntry(key *byte, value *byte) *MimeHeaderEntry {
//     return &MimeHeaderEntry{
//         key:key,
//         value:value,
//     }
// }
// func MakeMimeHeaderEntry(key *byte, value *byte) MimeHeaderEntry {
//     return MimeHeaderEntry{
//         key:key,
//         value:value,
//     }
// }
func (this *MimeHeaderEntry) GetKey() *byte        { return this.key }
func (this *MimeHeaderEntry) SetKey(value *byte)   { this.key = value }
func (this *MimeHeaderEntry) GetValue() *byte      { return this.value }
func (this *MimeHeaderEntry) SetValue(value *byte) { this.value = value }
