package cli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

/**
 * PhpCliServerPoller
 */
type PhpCliServerPoller struct {
	rfds   fd_set
	wfds   fd_set
	active struct {
		rfds fd_set
		wfds fd_set
	}
	max_fd core.PhpSocketT
}

func (this *PhpCliServerPoller) GetRfds() fd_set                { return this.rfds }
func (this *PhpCliServerPoller) GetWfds() fd_set                { return this.wfds }
func (this *PhpCliServerPoller) GetActiveRfds() fd_set          { return this.active.rfds }
func (this *PhpCliServerPoller) GetActiveWfds() fd_set          { return this.active.wfds }
func (this *PhpCliServerPoller) GetMaxFd() core.PhpSocketT      { return this.max_fd }
func (this *PhpCliServerPoller) SetMaxFd(value core.PhpSocketT) { this.max_fd = value }

/**
 * PhpCliServerRequest
 */
type PhpCliServerRequest struct {
	request_method        PhpHttpMethod
	protocol_version      int
	request_uri           string
	vpath                 string
	path_translated       string
	path_info             string
	query_string          string
	headers               types.Array
	headers_original_case types.Array
	content               string
	ext                   string
	sb                    zend.ZendStatT
}

func (this *PhpCliServerRequest) Ctor() {
	//this.request_method = PHP_HTTP_GET
	this.protocol_version = 0
	this.request_uri = ""
	this.vpath = ""
	this.path_translated = ""
	this.path_info = ""
	this.query_string = ""
	this.headers.Clean()
	this.headers_original_case.Clean()
	this.content = ""
	this.ext = ""
	//this.sb = 0
}

func (this *PhpCliServerRequest) ReadRequest(req *http.Request) {
	this.request_method = req.Method // todo 字符串转数字
	this.protocol_version = req.ProtoMajor*100 + req.ProtoMinor
	this.request_uri = req.RequestURI
	this.vpath = "" // todo vpath 是什么？
	this.path_info = req.URL.Path
	this.path_translated = "" // todo
	this.query_string = req.URL.RawQuery
	this.headers = req.Header
	this.headers_original_case = req.Header
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		this.content = string(body)
	}
}

func (this *PhpCliServerRequest) GetRequestMethod() PhpHttpMethod { return this.request_method }
func (this *PhpCliServerRequest) GetProtocolVersion() int         { return this.protocol_version }
func (this *PhpCliServerRequest) GetRequestUri() *byte            { return this.request_uri }
func (this *PhpCliServerRequest) GetRequestUriLen() int           { return this.request_uri_len }
func (this *PhpCliServerRequest) GetVpath() *byte                 { return this.vpath }
func (this *PhpCliServerRequest) SetVpath(value *byte)            { this.vpath = value }
func (this *PhpCliServerRequest) GetVpathLen() int                { return this.vpath_len }
func (this *PhpCliServerRequest) SetVpathLen(value int)           { this.vpath_len = value }
func (this *PhpCliServerRequest) GetPathTranslated() *byte        { return this.path_translated }
func (this *PhpCliServerRequest) SetPathTranslated(value *byte)   { this.path_translated = value }
func (this *PhpCliServerRequest) GetPathTranslatedLen() int       { return this.path_translated_len }
func (this *PhpCliServerRequest) SetPathTranslatedLen(value int)  { this.path_translated_len = value }
func (this *PhpCliServerRequest) GetPathInfo() *byte              { return this.path_info }
func (this *PhpCliServerRequest) SetPathInfo(value *byte)         { this.path_info = value }
func (this *PhpCliServerRequest) GetPathInfoLen() int             { return this.path_info_len }
func (this *PhpCliServerRequest) SetPathInfoLen(value int)        { this.path_info_len = value }
func (this *PhpCliServerRequest) GetQueryString() *byte           { return this.query_string }
func (this *PhpCliServerRequest) GetQueryStringLen() int          { return this.query_string_len }
func (this *PhpCliServerRequest) GetHeaders() types.Array         { return this.headers }
func (this *PhpCliServerRequest) GetHeadersOriginalCase() types.Array {
	return this.headers_original_case
}
func (this *PhpCliServerRequest) GetContent() *byte          { return this.content }
func (this *PhpCliServerRequest) SetContent(value *byte)     { this.content = value }
func (this *PhpCliServerRequest) GetContentLen() int         { return this.content_len }
func (this *PhpCliServerRequest) SetContentLen(value int)    { this.content_len = value }
func (this *PhpCliServerRequest) GetExt() *byte              { return this.ext }
func (this *PhpCliServerRequest) SetExt(value *byte)         { this.ext = value }
func (this *PhpCliServerRequest) GetExtLen() int             { return this.ext_len }
func (this *PhpCliServerRequest) SetExtLen(value int)        { this.ext_len = value }
func (this *PhpCliServerRequest) GetSb() zend.ZendStatT      { return this.sb }
func (this *PhpCliServerRequest) SetSb(value zend.ZendStatT) { this.sb = value }

/**
 * PhpCliServerChunk
 */
type PhpCliServerChunk struct {
	next  *PhpCliServerChunk
	type_ int
	data  struct /* union */ {
		heap struct {
			block any
			p     *byte
			len_  int
		}
		immortal struct {
			p    *byte
			len_ int
		}
	}
}

func (this *PhpCliServerChunk) GetNext() *PhpCliServerChunk      { return this.next }
func (this *PhpCliServerChunk) SetNext(value *PhpCliServerChunk) { this.next = value }
func (this *PhpCliServerChunk) GetType() int                     { return this.type_ }
func (this *PhpCliServerChunk) SetType(value int)                { this.type_ = value }
func (this *PhpCliServerChunk) GetBlock() any                    { return this.data.heap.block }
func (this *PhpCliServerChunk) SetBlock(value any)               { this.data.heap.block = value }
func (this *PhpCliServerChunk) GetDataHeapP() *byte              { return this.data.heap.p }
func (this *PhpCliServerChunk) SetDataHeapP(value *byte)         { this.data.heap.p = value }
func (this *PhpCliServerChunk) GetDataHeapLen() int              { return this.data.heap.len_ }
func (this *PhpCliServerChunk) SetDataHeapLen(value int)         { this.data.heap.len_ = value }
func (this *PhpCliServerChunk) GetDataImmortalP() *byte          { return this.data.immortal.p }
func (this *PhpCliServerChunk) SetDataImmortalP(value *byte)     { this.data.immortal.p = value }
func (this *PhpCliServerChunk) GetDataImmortalLen() int          { return this.data.immortal.len_ }
func (this *PhpCliServerChunk) SetDataImmortalLen(value int)     { this.data.immortal.len_ = value }

/**
 * PhpCliServerBuffer
 */
type PhpCliServerBuffer struct {
	first *PhpCliServerChunk
	last  *PhpCliServerChunk
}

// func NewPhpCliServerBuffer(first *PhpCliServerChunk, last *PhpCliServerChunk) *PhpCliServerBuffer {
//     return &PhpCliServerBuffer{
//         first:first,
//         last:last,
//     }
// }
// func MakePhpCliServerBuffer(first *PhpCliServerChunk, last *PhpCliServerChunk) PhpCliServerBuffer {
//     return PhpCliServerBuffer{
//         first:first,
//         last:last,
//     }
// }
func (this *PhpCliServerBuffer) GetFirst() *PhpCliServerChunk      { return this.first }
func (this *PhpCliServerBuffer) SetFirst(value *PhpCliServerChunk) { this.first = value }
func (this *PhpCliServerBuffer) GetLast() *PhpCliServerChunk       { return this.last }
func (this *PhpCliServerBuffer) SetLast(value *PhpCliServerChunk)  { this.last = value }

/**
 * PhpCliServerContentSender
 */
type PhpCliServerContentSender struct {
	buffer PhpCliServerBuffer
}

func (this *PhpCliServerContentSender) GetBuffer() PhpCliServerBuffer { return this.buffer }

/**
 * PhpCliServerClient
 */
type PhpCliServerClient struct {
	server                        *PhpCliServer
	sock                          core.PhpSocketT
	addr                          *__struct__sockaddr
	addr_len                      socklen_t
	addr_str                      *byte
	addr_str_len                  int
	parser                        PhpHttpParser
	request_read                  uint
	current_header_name           *byte
	current_header_name_len       int
	current_header_name_allocated uint
	current_header_value          *byte
	current_header_value_len      int
	last_header_element           int
	post_read_offset              int
	request                       PhpCliServerRequest
	content_sender_initialized    uint
	content_sender                PhpCliServerContentSender
	file_fd                       int
}

func NewPhpCliServerClient(server *PhpCliServer) *PhpCliServerClient {
	client := &PhpCliServerClient{
		server: server,
		//sock:     clientSock,
		//addr:     addr,
		//addr_len: addr_len,
	}

	return client
}

func (this *PhpCliServerClient) GetServer() *PhpCliServer          { return this.server }
func (this *PhpCliServerClient) SetServer(value *PhpCliServer)     { this.server = value }
func (this *PhpCliServerClient) GetSock() core.PhpSocketT          { return this.sock }
func (this *PhpCliServerClient) SetSock(value core.PhpSocketT)     { this.sock = value }
func (this *PhpCliServerClient) GetAddr() *__struct__sockaddr      { return this.addr }
func (this *PhpCliServerClient) SetAddr(value *__struct__sockaddr) { this.addr = value }
func (this *PhpCliServerClient) SetAddrLen(value socklen_t)        { this.addr_len = value }
func (this *PhpCliServerClient) GetAddrStr() *byte                 { return this.addr_str }
func (this *PhpCliServerClient) SetAddrStr(value *byte)            { this.addr_str = value }
func (this *PhpCliServerClient) SetAddrStrLen(value int)           { this.addr_str_len = value }
func (this *PhpCliServerClient) GetParser() PhpHttpParser          { return this.parser }
func (this *PhpCliServerClient) GetRequestRead() uint              { return this.request_read }
func (this *PhpCliServerClient) SetRequestRead(value uint)         { this.request_read = value }
func (this *PhpCliServerClient) GetCurrentHeaderName() *byte       { return this.current_header_name }
func (this *PhpCliServerClient) SetCurrentHeaderName(value *byte)  { this.current_header_name = value }
func (this *PhpCliServerClient) GetCurrentHeaderNameLen() int      { return this.current_header_name_len }
func (this *PhpCliServerClient) SetCurrentHeaderNameLen(value int) {
	this.current_header_name_len = value
}
func (this *PhpCliServerClient) GetCurrentHeaderNameAllocated() uint {
	return this.current_header_name_allocated
}
func (this *PhpCliServerClient) SetCurrentHeaderNameAllocated(value uint) {
	this.current_header_name_allocated = value
}
func (this *PhpCliServerClient) GetCurrentHeaderValue() *byte      { return this.current_header_value }
func (this *PhpCliServerClient) SetCurrentHeaderValue(value *byte) { this.current_header_value = value }
func (this *PhpCliServerClient) GetCurrentHeaderValueLen() int     { return this.current_header_value_len }
func (this *PhpCliServerClient) SetCurrentHeaderValueLen(value int) {
	this.current_header_value_len = value
}
func (this *PhpCliServerClient) GetLastHeaderElement() int       { return this.last_header_element }
func (this *PhpCliServerClient) SetLastHeaderElement(value int)  { this.last_header_element = value }
func (this *PhpCliServerClient) GetPostReadOffset() int          { return this.post_read_offset }
func (this *PhpCliServerClient) SetPostReadOffset(value int)     { this.post_read_offset = value }
func (this *PhpCliServerClient) GetRequest() PhpCliServerRequest { return this.request }
func (this *PhpCliServerClient) GetContentSenderInitialized() uint {
	return this.content_sender_initialized
}
func (this *PhpCliServerClient) SetContentSenderInitialized(value uint) {
	this.content_sender_initialized = value
}
func (this *PhpCliServerClient) GetContentSender() PhpCliServerContentSender {
	return this.content_sender
}
func (this *PhpCliServerClient) GetFileFd() int      { return this.file_fd }
func (this *PhpCliServerClient) SetFileFd(value int) { this.file_fd = value }

/**
 * PhpCliServer
 */
type PhpCliServer struct {
	server_sock    core.PhpSocketT
	poller         PhpCliServerPoller
	is_running     int
	host           string
	port           int
	address_family int
	document_root  string
	router         string
	socklen        socklen_t
	clients        types.Array
}

func (this *PhpCliServer) Serve() error {
	addr := fmt.Sprintf("%s:%d", this.host, this.port)

	err := http.ListenAndServe(addr, http.HandlerFunc(this.handler))
	if err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (this *PhpCliServer) handler(writer http.ResponseWriter, request *http.Request) {
	client := this.NewClient(nil)
	PhpCliServerRecvEventReadRequest(this, client, request)
	PhpCliServerSendEvent(this, client)
}

func (this *PhpCliServer) NewClient() *PhpCliServerClient {
	client := NewPhpCliServerClient(this)
	PhpCliServerClientCtor(client, this)
	types.ZendHashIndexUpdatePtr(&this.clients, 0, client)
	return client
}

func (this *PhpCliServer) SetHostStr(value string)         { this.host = value }
func (this *PhpCliServer) SetDocumentRootStr(value string) { this.document_root = value }
func (this *PhpCliServer) SetRouterStr(value string)       { this.router = value }

func (this *PhpCliServer) GetServerSock() core.PhpSocketT      { return this.server_sock }
func (this *PhpCliServer) SetServerSock(value core.PhpSocketT) { this.server_sock = value }
func (this *PhpCliServer) GetPoller() PhpCliServerPoller       { return this.poller }
func (this *PhpCliServer) GetIsRunning() int                   { return this.is_running }
func (this *PhpCliServer) SetIsRunning(value int)              { this.is_running = value }
func (this *PhpCliServer) GetHost() *byte                      { return this.host }
func (this *PhpCliServer) SetHost(value *byte)                 { this.host = value }
func (this *PhpCliServer) GetPort() int                        { return this.port }
func (this *PhpCliServer) SetPort(value int)                   { this.port = value }
func (this *PhpCliServer) GetAddressFamily() int               { return this.address_family }
func (this *PhpCliServer) GetDocumentRoot() *byte              { return this.document_root }
func (this *PhpCliServer) SetDocumentRoot(value *byte)         { this.document_root = value }
func (this *PhpCliServer) GetDocumentRootLen() int             { return this.document_root_len }
func (this *PhpCliServer) GetRouter() *byte                    { return this.router }
func (this *PhpCliServer) SetRouter(value *byte)               { this.router = value }
func (this *PhpCliServer) GetSocklen() socklen_t               { return this.socklen }
func (this *PhpCliServer) GetClients() types.Array             { return this.clients }

/**
 * PhpCliServerDoEventForEachFdCallbackParams
 */
type PhpCliServerDoEventForEachFdCallbackParams struct {
	Server       *PhpCliServer
	ReadHandler  func(*PhpCliServer, *PhpCliServerClient) int
	WriteHandler func(*PhpCliServer, *PhpCliServerClient) int
}

func (this *PhpCliServerDoEventForEachFdCallbackParams) GetServer() *PhpCliServer { return this.Server }
func (this *PhpCliServerDoEventForEachFdCallbackParams) GetRhandler() func(*PhpCliServer, *PhpCliServerClient) int {
	return this.ReadHandler
}
func (this *PhpCliServerDoEventForEachFdCallbackParams) GetWhandler() func(*PhpCliServer, *PhpCliServerClient) int {
	return this.WriteHandler
}
