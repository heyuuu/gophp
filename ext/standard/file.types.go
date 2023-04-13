package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
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

func (this *PhpMetaTagsData) GetStream() *core.PhpStream      { return this.stream }
func (this *PhpMetaTagsData) SetStream(value *core.PhpStream) { this.stream = value }
func (this *PhpMetaTagsData) GetUlc() int                     { return this.ulc }
func (this *PhpMetaTagsData) SetUlc(value int)                { this.ulc = value }
func (this *PhpMetaTagsData) GetLc() int                      { return this.lc }
func (this *PhpMetaTagsData) SetLc(value int)                 { this.lc = value }
func (this *PhpMetaTagsData) GetTokenData() *byte             { return this.token_data }
func (this *PhpMetaTagsData) SetTokenData(value *byte)        { this.token_data = value }
func (this *PhpMetaTagsData) GetTokenLen() int                { return this.token_len }
func (this *PhpMetaTagsData) SetTokenLen(value int)           { this.token_len = value }
func (this *PhpMetaTagsData) GetInMeta() int                  { return this.in_meta }
func (this *PhpMetaTagsData) SetInMeta(value int)             { this.in_meta = value }

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
	stream_wrappers              map[string]*core.PhpStreamWrapper
	stream_filters               map[string]*streams.PhpStreamFilterFactory
	wrapper_errors               *types.Array
	pclose_wait                  int
}

func (this *PhpFileGlobals) StreamWrappers() map[string]*core.PhpStreamWrapper {
	return this.stream_wrappers
}
func (this *PhpFileGlobals) GetStreamWrappers() map[string]*core.PhpStreamWrapper {
	return this.stream_wrappers
}
func (this *PhpFileGlobals) SetStreamWrappers(value map[string]*core.PhpStreamWrapper) {
	this.stream_wrappers = value
}

//
func (this *PhpFileGlobals) SetDefChunkSize(value int) { this.def_chunk_size = value }

func (this *PhpFileGlobals) GetAutoDetectLineEndings() types.ZendBool {
	return this.auto_detect_line_endings
}
func (this *PhpFileGlobals) SetAutoDetectLineEndings(value types.ZendBool) {
	this.auto_detect_line_endings = value
}
func (this *PhpFileGlobals) GetDefaultSocketTimeout() zend.ZendLong {
	return this.default_socket_timeout
}
func (this *PhpFileGlobals) SetDefaultSocketTimeout(value zend.ZendLong) {
	this.default_socket_timeout = value
}
func (this *PhpFileGlobals) GetUserAgent() *byte        { return this.user_agent }
func (this *PhpFileGlobals) SetUserAgent(value *byte)   { this.user_agent = value }
func (this *PhpFileGlobals) GetFromAddress() *byte      { return this.from_address }
func (this *PhpFileGlobals) SetFromAddress(value *byte) { this.from_address = value }
func (this *PhpFileGlobals) GetUserStreamCurrentFilename() *byte {
	return this.user_stream_current_filename
}
func (this *PhpFileGlobals) SetUserStreamCurrentFilename(value *byte) {
	this.user_stream_current_filename = value
}
func (this *PhpFileGlobals) GetDefaultContext() *core.PhpStreamContext { return this.default_context }
func (this *PhpFileGlobals) SetDefaultContext(value *core.PhpStreamContext) {
	this.default_context = value
}

func (this *PhpFileGlobals) GetStreamFilters() map[string]*streams.PhpStreamFilterFactory {
	return this.stream_filters
}
func (this *PhpFileGlobals) SetStreamFilters(value map[string]*streams.PhpStreamFilterFactory) {
	this.stream_filters = value
}

func (this *PhpFileGlobals) GetWrapperErrors() *types.Array      { return this.wrapper_errors }
func (this *PhpFileGlobals) SetWrapperErrors(value *types.Array) { this.wrapper_errors = value }
func (this *PhpFileGlobals) GetPcloseWait() int                  { return this.pclose_wait }
func (this *PhpFileGlobals) SetPcloseWait(value int)             { this.pclose_wait = value }
