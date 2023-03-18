// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

/**
 * PhpMetaTagsData
 */
type PhpMetaTagsData struct {
	stream       *core.PhpStream
	ulc          int
	lc           int
	input_buffer *byte
	token_data   *byte
	token_len    int
	in_meta      int
}

//             func MakePhpMetaTagsData(
// stream *core.PhpStream,
// ulc int,
// lc int,
// input_buffer *byte,
// token_data *byte,
// token_len int,
// in_meta int,
// ) PhpMetaTagsData {
//                 return PhpMetaTagsData{
//                     stream:stream,
//                     ulc:ulc,
//                     lc:lc,
//                     input_buffer:input_buffer,
//                     token_data:token_data,
//                     token_len:token_len,
//                     in_meta:in_meta,
//                 }
//             }
func (this *PhpMetaTagsData) GetStream() *core.PhpStream      { return this.stream }
func (this *PhpMetaTagsData) SetStream(value *core.PhpStream) { this.stream = value }
func (this *PhpMetaTagsData) GetUlc() int                     { return this.ulc }
func (this *PhpMetaTagsData) SetUlc(value int)                { this.ulc = value }
func (this *PhpMetaTagsData) GetLc() int                      { return this.lc }
func (this *PhpMetaTagsData) SetLc(value int)                 { this.lc = value }

// func (this *PhpMetaTagsData)  GetInputBuffer() *byte      { return this.input_buffer }
// func (this *PhpMetaTagsData) SetInputBuffer(value *byte) { this.input_buffer = value }
func (this *PhpMetaTagsData) GetTokenData() *byte      { return this.token_data }
func (this *PhpMetaTagsData) SetTokenData(value *byte) { this.token_data = value }
func (this *PhpMetaTagsData) GetTokenLen() int         { return this.token_len }
func (this *PhpMetaTagsData) SetTokenLen(value int)    { this.token_len = value }
func (this *PhpMetaTagsData) GetInMeta() int           { return this.in_meta }
func (this *PhpMetaTagsData) SetInMeta(value int)      { this.in_meta = value }

/**
 * PhpFileGlobals
 */
type PhpFileGlobals struct {
	pclose_ret                   int
	def_chunk_size               int
	auto_detect_line_endings     types.ZendBool
	default_socket_timeout       zend.ZendLong
	user_agent                   *byte
	from_address                 *byte
	user_stream_current_filename *byte
	default_context              *core.PhpStreamContext
	stream_wrappers              *types.HashTable
	stream_filters               *types.HashTable
	wrapper_errors               *types.HashTable
	pclose_wait                  int
}

//             func MakePhpFileGlobals(
// pclose_ret int,
// def_chunk_size int,
// auto_detect_line_endings zend.ZendBool,
// default_socket_timeout zend.ZendLong,
// user_agent *byte,
// from_address *byte,
// user_stream_current_filename *byte,
// default_context *core.PhpStreamContext,
// stream_wrappers *zend.HashTable,
// stream_filters *zend.HashTable,
// wrapper_errors *zend.HashTable,
// pclose_wait int,
// ) PhpFileGlobals {
//                 return PhpFileGlobals{
//                     pclose_ret:pclose_ret,
//                     def_chunk_size:def_chunk_size,
//                     auto_detect_line_endings:auto_detect_line_endings,
//                     default_socket_timeout:default_socket_timeout,
//                     user_agent:user_agent,
//                     from_address:from_address,
//                     user_stream_current_filename:user_stream_current_filename,
//                     default_context:default_context,
//                     stream_wrappers:stream_wrappers,
//                     stream_filters:stream_filters,
//                     wrapper_errors:wrapper_errors,
//                     pclose_wait:pclose_wait,
//                 }
//             }
// func (this *PhpFileGlobals)  GetPcloseRet() int      { return this.pclose_ret }
// func (this *PhpFileGlobals) SetPcloseRet(value int) { this.pclose_ret = value }
// func (this *PhpFileGlobals)  GetDefChunkSize() int      { return this.def_chunk_size }
func (this *PhpFileGlobals) SetDefChunkSize(value int) { this.def_chunk_size = value }

// func (this *PhpFileGlobals)  GetAutoDetectLineEndings() zend.ZendBool      { return this.auto_detect_line_endings }
// func (this *PhpFileGlobals) SetAutoDetectLineEndings(value zend.ZendBool) { this.auto_detect_line_endings = value }
// func (this *PhpFileGlobals)  GetDefaultSocketTimeout() zend.ZendLong      { return this.default_socket_timeout }
// func (this *PhpFileGlobals) SetDefaultSocketTimeout(value zend.ZendLong) { this.default_socket_timeout = value }
// func (this *PhpFileGlobals)  GetUserAgent() *byte      { return this.user_agent }
// func (this *PhpFileGlobals) SetUserAgent(value *byte) { this.user_agent = value }
// func (this *PhpFileGlobals)  GetFromAddress() *byte      { return this.from_address }
// func (this *PhpFileGlobals) SetFromAddress(value *byte) { this.from_address = value }
// func (this *PhpFileGlobals)  GetUserStreamCurrentFilename() *byte      { return this.user_stream_current_filename }
// func (this *PhpFileGlobals) SetUserStreamCurrentFilename(value *byte) { this.user_stream_current_filename = value }
// func (this *PhpFileGlobals)  GetDefaultContext() *core.PhpStreamContext      { return this.default_context }
// func (this *PhpFileGlobals) SetDefaultContext(value *core.PhpStreamContext) { this.default_context = value }
// func (this *PhpFileGlobals)  GetStreamWrappers() *zend.HashTable      { return this.stream_wrappers }
// func (this *PhpFileGlobals) SetStreamWrappers(value *zend.HashTable) { this.stream_wrappers = value }
// func (this *PhpFileGlobals)  GetStreamFilters() *zend.HashTable      { return this.stream_filters }
// func (this *PhpFileGlobals) SetStreamFilters(value *zend.HashTable) { this.stream_filters = value }
// func (this *PhpFileGlobals)  GetWrapperErrors() *zend.HashTable      { return this.wrapper_errors }
// func (this *PhpFileGlobals) SetWrapperErrors(value *zend.HashTable) { this.wrapper_errors = value }
// func (this *PhpFileGlobals)  GetPcloseWait() int      { return this.pclose_wait }
// func (this *PhpFileGlobals) SetPcloseWait(value int) { this.pclose_wait = value }
