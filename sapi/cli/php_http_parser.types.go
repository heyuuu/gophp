// <<generate>>

package cli

/**
 * PhpHttpParser
 */
type PhpHttpParser struct {
	type_          uint8
	flags          uint8
	state          uint8
	header_state   uint8
	index          uint8
	nread          uint32
	content_length ssize_t
	http_major     uint16
	http_minor     uint16
	status_code    uint16
	method         uint8
	upgrade        byte
	data           any
}

// func NewPhpHttpParser(type_ uint8, flags uint8, state uint8, header_state uint8, index uint8, nread uint32, content_length ssize_t, http_major uint16, http_minor uint16, status_code uint16, method uint8, upgrade byte, data any) *PhpHttpParser {
//     return &PhpHttpParser{
//         type_:type_,
//         flags:flags,
//         state:state,
//         header_state:header_state,
//         index:index,
//         nread:nread,
//         content_length:content_length,
//         http_major:http_major,
//         http_minor:http_minor,
//         status_code:status_code,
//         method:method,
//         upgrade:upgrade,
//         data:data,
//     }
// }
// func MakePhpHttpParser(type_ uint8, flags uint8, state uint8, header_state uint8, index uint8, nread uint32, content_length ssize_t, http_major uint16, http_minor uint16, status_code uint16, method uint8, upgrade byte, data any) PhpHttpParser {
//     return PhpHttpParser{
//         type_:type_,
//         flags:flags,
//         state:state,
//         header_state:header_state,
//         index:index,
//         nread:nread,
//         content_length:content_length,
//         http_major:http_major,
//         http_minor:http_minor,
//         status_code:status_code,
//         method:method,
//         upgrade:upgrade,
//         data:data,
//     }
// }
func (this *PhpHttpParser) GetType() uint8      { return this.type_ }
func (this *PhpHttpParser) SetType(value uint8) { this.type_ = value }

// func (this *PhpHttpParser)  GetFlags() uint8      { return this.flags }
func (this *PhpHttpParser) SetFlags(value uint8)           { this.flags = value }
func (this *PhpHttpParser) GetState() uint8                { return this.state }
func (this *PhpHttpParser) SetState(value uint8)           { this.state = value }
func (this *PhpHttpParser) GetHeaderState() uint8          { return this.header_state }
func (this *PhpHttpParser) SetHeaderState(value uint8)     { this.header_state = value }
func (this *PhpHttpParser) GetIndex() uint8                { return this.index }
func (this *PhpHttpParser) SetIndex(value uint8)           { this.index = value }
func (this *PhpHttpParser) GetNread() uint32               { return this.nread }
func (this *PhpHttpParser) SetNread(value uint32)          { this.nread = value }
func (this *PhpHttpParser) GetContentLength() ssize_t      { return this.content_length }
func (this *PhpHttpParser) SetContentLength(value ssize_t) { this.content_length = value }
func (this *PhpHttpParser) GetHttpMajor() uint16           { return this.http_major }
func (this *PhpHttpParser) SetHttpMajor(value uint16)      { this.http_major = value }
func (this *PhpHttpParser) GetHttpMinor() uint16           { return this.http_minor }
func (this *PhpHttpParser) SetHttpMinor(value uint16)      { this.http_minor = value }
func (this *PhpHttpParser) GetStatusCode() uint16          { return this.status_code }
func (this *PhpHttpParser) SetStatusCode(value uint16)     { this.status_code = value }
func (this *PhpHttpParser) GetMethod() uint8               { return this.method }
func (this *PhpHttpParser) SetMethod(value uint8)          { this.method = value }

// func (this *PhpHttpParser)  GetUpgrade() byte      { return this.upgrade }
func (this *PhpHttpParser) SetUpgrade(value byte) { this.upgrade = value }
func (this *PhpHttpParser) GetData() any          { return this.data }
func (this *PhpHttpParser) SetData(value any)     { this.data = value }

/* PhpHttpParser.flags */
func (this *PhpHttpParser) AddFlags(value uint8)      { this.flags |= value }
func (this *PhpHttpParser) SubFlags(value uint8)      { this.flags &^= value }
func (this *PhpHttpParser) HasFlags(value uint8) bool { return this.flags&value != 0 }
func (this *PhpHttpParser) SwitchFlags(value uint8, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this PhpHttpParser) IsTrailing() bool            { return this.HasFlags(F_TRAILING) }
func (this PhpHttpParser) IsUpgrade() bool             { return this.HasFlags(F_UPGRADE) }
func (this PhpHttpParser) IsSkipbody() bool            { return this.HasFlags(F_SKIPBODY) }
func (this PhpHttpParser) IsChunked() bool             { return this.HasFlags(F_CHUNKED) }
func (this PhpHttpParser) IsConnectionClose() bool     { return this.HasFlags(F_CONNECTION_CLOSE) }
func (this PhpHttpParser) IsConnectionKeepAlive() bool { return this.HasFlags(F_CONNECTION_KEEP_ALIVE) }
func (this *PhpHttpParser) SetIsTrailing(cond bool)    { this.SwitchFlags(F_TRAILING, cond) }
func (this *PhpHttpParser) SetIsUpgrade(cond bool)     { this.SwitchFlags(F_UPGRADE, cond) }
func (this *PhpHttpParser) SetIsSkipbody(cond bool)    { this.SwitchFlags(F_SKIPBODY, cond) }
func (this *PhpHttpParser) SetIsChunked(cond bool)     { this.SwitchFlags(F_CHUNKED, cond) }
func (this *PhpHttpParser) SetIsConnectionClose(cond bool) {
	this.SwitchFlags(F_CONNECTION_CLOSE, cond)
}
func (this *PhpHttpParser) SetIsConnectionKeepAlive(cond bool) {
	this.SwitchFlags(F_CONNECTION_KEEP_ALIVE, cond)
}

/**
 * PhpHttpParserSettings
 */
type PhpHttpParserSettings struct {
	on_message_begin    PhpHttpCb
	on_path             PhpHttpDataCb
	on_query_string     PhpHttpDataCb
	on_url              PhpHttpDataCb
	on_fragment         PhpHttpDataCb
	on_header_field     PhpHttpDataCb
	on_header_value     PhpHttpDataCb
	on_headers_complete PhpHttpCb
	on_body             PhpHttpDataCb
	on_message_complete PhpHttpCb
}

// func NewPhpHttpParserSettings(on_message_begin PhpHttpCb, on_path PhpHttpDataCb, on_query_string PhpHttpDataCb, on_url PhpHttpDataCb, on_fragment PhpHttpDataCb, on_header_field PhpHttpDataCb, on_header_value PhpHttpDataCb, on_headers_complete PhpHttpCb, on_body PhpHttpDataCb, on_message_complete PhpHttpCb) *PhpHttpParserSettings {
//     return &PhpHttpParserSettings{
//         on_message_begin:on_message_begin,
//         on_path:on_path,
//         on_query_string:on_query_string,
//         on_url:on_url,
//         on_fragment:on_fragment,
//         on_header_field:on_header_field,
//         on_header_value:on_header_value,
//         on_headers_complete:on_headers_complete,
//         on_body:on_body,
//         on_message_complete:on_message_complete,
//     }
// }
func MakePhpHttpParserSettings(on_message_begin PhpHttpCb, on_path PhpHttpDataCb, on_query_string PhpHttpDataCb, on_url PhpHttpDataCb, on_fragment PhpHttpDataCb, on_header_field PhpHttpDataCb, on_header_value PhpHttpDataCb, on_headers_complete PhpHttpCb, on_body PhpHttpDataCb, on_message_complete PhpHttpCb) PhpHttpParserSettings {
	return PhpHttpParserSettings{
		on_message_begin:    on_message_begin,
		on_path:             on_path,
		on_query_string:     on_query_string,
		on_url:              on_url,
		on_fragment:         on_fragment,
		on_header_field:     on_header_field,
		on_header_value:     on_header_value,
		on_headers_complete: on_headers_complete,
		on_body:             on_body,
		on_message_complete: on_message_complete,
	}
}
func (this *PhpHttpParserSettings) GetOnMessageBegin() PhpHttpCb { return this.on_message_begin }

// func (this *PhpHttpParserSettings) SetOnMessageBegin(value PhpHttpCb) { this.on_message_begin = value }
func (this *PhpHttpParserSettings) GetOnPath() PhpHttpDataCb { return this.on_path }

// func (this *PhpHttpParserSettings) SetOnPath(value PhpHttpDataCb) { this.on_path = value }
func (this *PhpHttpParserSettings) GetOnQueryString() PhpHttpDataCb { return this.on_query_string }

// func (this *PhpHttpParserSettings) SetOnQueryString(value PhpHttpDataCb) { this.on_query_string = value }
func (this *PhpHttpParserSettings) GetOnUrl() PhpHttpDataCb { return this.on_url }

// func (this *PhpHttpParserSettings) SetOnUrl(value PhpHttpDataCb) { this.on_url = value }
func (this *PhpHttpParserSettings) GetOnFragment() PhpHttpDataCb { return this.on_fragment }

// func (this *PhpHttpParserSettings) SetOnFragment(value PhpHttpDataCb) { this.on_fragment = value }
func (this *PhpHttpParserSettings) GetOnHeaderField() PhpHttpDataCb { return this.on_header_field }

// func (this *PhpHttpParserSettings) SetOnHeaderField(value PhpHttpDataCb) { this.on_header_field = value }
func (this *PhpHttpParserSettings) GetOnHeaderValue() PhpHttpDataCb { return this.on_header_value }

// func (this *PhpHttpParserSettings) SetOnHeaderValue(value PhpHttpDataCb) { this.on_header_value = value }
func (this *PhpHttpParserSettings) GetOnHeadersComplete() PhpHttpCb { return this.on_headers_complete }

// func (this *PhpHttpParserSettings) SetOnHeadersComplete(value PhpHttpCb) { this.on_headers_complete = value }
func (this *PhpHttpParserSettings) GetOnBody() PhpHttpDataCb { return this.on_body }

// func (this *PhpHttpParserSettings) SetOnBody(value PhpHttpDataCb) { this.on_body = value }
func (this *PhpHttpParserSettings) GetOnMessageComplete() PhpHttpCb { return this.on_message_complete }

// func (this *PhpHttpParserSettings) SetOnMessageComplete(value PhpHttpCb) { this.on_message_complete = value }
