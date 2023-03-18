// <<generate>>

package core

import (
	"sik/zend"
	"sik/zend/types"
)

/**
 * SapiHeader
 */
type SapiHeader struct {
	header     *byte
	header_len int
}

func (this *SapiHeader) GetHeader() *byte       { return this.header }
func (this *SapiHeader) SetHeader(value *byte)  { this.header = value }
func (this *SapiHeader) GetHeaderLen() int      { return this.header_len }
func (this *SapiHeader) SetHeaderLen(value int) { this.header_len = value }

/**
 * SapiHeaders
 */
type SapiHeaders struct {
	headers                   zend.ZendLlist
	http_response_code        int
	send_default_content_type uint8
	mimetype                  *byte
	http_status_line          *byte
}

func (this *SapiHeaders) GetHeaders() zend.ZendLlist { return this.headers }

/**
 * SapiRequestInfo
 */
type SapiRequestInfo struct {
	request_method      *byte
	query_string        *byte
	cookie_data         *byte
	content_length      zend.ZendLong
	path_translated     *byte
	request_uri         *byte
	request_body        *PhpStream
	content_type        *byte
	headers_only        types.ZendBool
	no_headers          types.ZendBool
	headers_read        types.ZendBool
	post_entry          *SapiPostEntry
	content_type_dup    *byte
	auth_user           *byte
	auth_password       *byte
	auth_digest         *byte
	argv0               *byte
	current_user        *byte
	current_user_length int
	argc                int
	argv                **byte
	proto_num           int
}

func (this *SapiRequestInfo) SetRequestMethod(value *byte)         { this.request_method = value }
func (this *SapiRequestInfo) SetQueryString(value *byte)           { this.query_string = value }
func (this *SapiRequestInfo) SetContentLength(value zend.ZendLong) { this.content_length = value }
func (this *SapiRequestInfo) SetPathTranslated(value *byte)        { this.path_translated = value }
func (this *SapiRequestInfo) SetRequestUri(value *byte)            { this.request_uri = value }
func (this *SapiRequestInfo) SetContentType(value *byte)           { this.content_type = value }
func (this *SapiRequestInfo) SetAuthUser(value *byte)              { this.auth_user = value }
func (this *SapiRequestInfo) GetAuthPassword() *byte               { return this.auth_password }
func (this *SapiRequestInfo) SetAuthPassword(value *byte)          { this.auth_password = value }
func (this *SapiRequestInfo) GetAuthDigest() *byte                 { return this.auth_digest }
func (this *SapiRequestInfo) SetAuthDigest(value *byte)            { this.auth_digest = value }
func (this *SapiRequestInfo) SetProtoNum(value int)                { this.proto_num = value }

/**
 * SapiGlobals
 */
type SapiGlobals struct {
	server_context           any
	request_info             SapiRequestInfo
	sapi_headers             SapiHeaders
	read_post_bytes          int64
	post_read                uint8
	headers_sent             uint8
	global_stat              zend.ZendStatT
	default_mimetype         *byte
	default_charset          *byte
	rfc1867_uploaded_files   *types.HashTable
	post_max_size            zend.ZendLong
	options                  int
	sapi_started             types.ZendBool
	global_request_time      float64
	known_post_content_types types.HashTable
	callback_func            types.Zval
	fci_cache                zend.ZendFcallInfoCache
}

func (this *SapiGlobals) Init() {
	this.known_post_content_types = *types.NewZendArrayEx(8, nil, true)
	PhpSetupSapiContentTypes()
}

func (this *SapiGlobals) Destroy() {
	this.known_post_content_types.Destroy()
}

/**
 * generate
 */
func (this *SapiGlobals) GetDefaultMimetype() *byte     { return this.default_mimetype }
func (this *SapiGlobals) GetDefaultCharset() *byte      { return this.default_charset }
func (this *SapiGlobals) GetPostMaxSize() zend.ZendLong { return this.post_max_size }
func (this *SapiGlobals) GetKnownPostContentTypes() types.HashTable {
	return this.known_post_content_types
}

/**
 * SapiHeaderLine
 */
type SapiHeaderLine struct {
	line          *byte
	line_len      int
	response_code zend.ZendLong
}

func MakeSapiHeaderLineEx(line string) SapiHeaderLine {
	return SapiHeaderLine{}
}

func MakeSapiHeaderLine(line *byte, line_len int, response_code zend.ZendLong) SapiHeaderLine {
	return SapiHeaderLine{
		line:          line,
		line_len:      line_len,
		response_code: response_code,
	}
}
func (this *SapiHeaderLine) GetLine() *byte                 { return this.line }
func (this *SapiHeaderLine) SetLine(value *byte)            { this.line = value }
func (this *SapiHeaderLine) GetLineLen() int                { return this.line_len }
func (this *SapiHeaderLine) SetLineLen(value int)           { this.line_len = value }
func (this *SapiHeaderLine) GetResponseCode() zend.ZendLong { return this.response_code }

/**
 * SapiPostEntry
 */
type SapiPostEntry struct {
	content_type     *byte
	content_type_len uint32
	post_reader      func()
	post_handler     func(content_type_dup *byte, arg any)
}

func MakeSapiPostEntry(content_type *byte, content_type_len uint32, post_reader func(), post_handler func(content_type_dup *byte, arg any)) SapiPostEntry {
	return SapiPostEntry{
		content_type:     content_type,
		content_type_len: content_type_len,
		post_reader:      post_reader,
		post_handler:     post_handler,
	}
}
func (this *SapiPostEntry) GetContentType() *byte     { return this.content_type }
func (this *SapiPostEntry) GetContentTypeLen() uint32 { return this.content_type_len }
func (this *SapiPostEntry) GetPostReader() func()     { return this.post_reader }
