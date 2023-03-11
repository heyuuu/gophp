// <<generate>>

package cli

import (
	"fmt"
	"log"
	"net/http"
	"os"
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
	"strconv"
	"strings"
)

func PhpCliServerGetSystemTime(buf *byte) int {
	var tv __struct__timeval
	var tm __struct__tm
	gettimeofday(&tv, nil)

	/* TODO: should be checked for NULL tm/return vaue */

	core.PhpLocaltimeR(tv.tv_sec, &tm)
	core.PhpAsctimeR(&tm, buf)
	return 0
}
func GetLastError() *byte { return zend.Pestrdup(strerror(errno), 1) }
func GetStatusString(code int) string {
	if result, ok := core.HttpStatusMap[code]; ok {
		return result
	}

	/* Returning NULL would require complicating append_http_status_line() to
	 * not segfault in that case, so let's just return a placeholder, since RFC
	 * 2616 requires a reason phrase. This is basically what a lot of other Web
	 * servers do in this case anyway. */
	return "Unknown Status Code"
}
func GetTemplateString(code int) string {
	if val, ok := TemplateMap[code]; ok {
		return val
	}
	panic(fmt.Sprintf("Template for code:%d is not found", code))
}
func AppendHttpStatusLine(buffer *zend.SmartStr, protocol_version int, response_code int, persistent int) {
	if response_code == 0 {
		response_code = 200
	}
	buffer.AppendString("HTTP")
	buffer.AppendByte('/')
	buffer.AppendLong(protocol_version / 100)
	buffer.AppendByte('.')
	buffer.AppendLong(protocol_version % 100)
	buffer.AppendByte(' ')
	buffer.AppendLong(response_code)
	buffer.AppendByte(' ')
	buffer.AppendString(b.CastStrAuto(GetStatusString(response_code)))
	buffer.AppendString("\r\n")
}
func AppendEssentialHeaders(buffer *zend.SmartStr, client *PhpCliServerClient, persistent int) {
	var val *byte
	var tv __struct__timeval = __struct__timeval{0}
	if nil != b.Assign(&val, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "host", b.SizeOf("\"host\"")-1)) {
		buffer.AppendString("Host: ")
		buffer.AppendString(b.CastStrAuto(val))
		buffer.AppendString("\r\n")
	}
	if !(gettimeofday(&tv, nil)) {
		var dt *zend.ZendString = php_format_date("D, d M Y H:i:s", b.SizeOf("\"D, d M Y H:i:s\"")-1, tv.tv_sec, 0)
		buffer.AppendString("Date: ")
		buffer.AppendString(b.CastStrAuto(dt.GetVal()))
		buffer.AppendString(" GMT\r\n")
		zend.ZendStringReleaseEx(dt, 0)
	}
	buffer.AppendString("Connection: close\r\n")
}
func GetMimeType(server *PhpCliServer, ext *byte, ext_len int) *byte {
	var s = b.CastStr(ext, ext_len)
	var sLower = strings.ToLower(s)
	if mimeType, ok := MimeTypeMap[sLower]; ok {
		return b.CastStrPtr(mimeType)
	}
	return nil
}
func ZifApacheRequestHeaders(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	var client *PhpCliServerClient
	var headers *zend.HashTable
	var key *zend.ZendString
	var value *byte
	var tmp zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	client = core.SG__().server_context
	headers = client.GetRequest().GetHeadersOriginalCase()
	zend.ArrayInitSize(return_value, headers.GetNNumOfElements())
	var __ht *zend.HashTable = headers
	for _, _p := range __ht.foreachData() {
		var _z *zend.Zval = _p.GetVal()

		key = _p.GetKey()
		value = _z.GetPtr()
		zend.ZVAL_STRING(&tmp, value)
		return_value.GetArr().SymtableUpdate(key.GetStr(), &tmp)
	}
}
func AddResponseHeader(h *core.SapiHeader, return_value *zend.Zval) {
	var s *byte
	var p *byte
	var len_ ptrdiff_t
	if h.GetHeaderLen() > 0 {
		p = strchr(h.GetHeader(), ':')
		len_ = p - h.GetHeader()
		if p != nil && len_ > 0 {
			for len_ > 0 && (h.GetHeader()[len_-1] == ' ' || h.GetHeader()[len_-1] == '\t') {
				len_--
			}
			if len_ {
				s = zend.DoAlloca(len_+1, use_heap)
				memcpy(s, h.GetHeader(), len_)
				s[len_] = 0
				for {
					p++
					if !((*p) == ' ' || (*p) == '\t') {
						break
					}
				}
				zend.AddAssocStringlEx(return_value, s, uint32(len_), p, h.GetHeaderLen()-(p-h.GetHeader()))
				zend.FreeAlloca(s, use_heap)
			}
		}
	}
}
func ZifApacheResponseHeaders(executeData *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	zend.ZendLlistApplyWithArgument(core.SG__().sapi_headers.headers, zend.LlistApplyWithArgFuncT(AddResponseHeader), return_value)
}
func ZmStartupCliServer(type_ int, module_number int) int {
	zend.REGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmShutdownCliServer(type_ int, module_number int) int {
	zend.UNREGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmInfoCliServer(ZEND_MODULE_INFO_FUNC_ARGS) { zend.DISPLAY_INI_ENTRIES() }
func SapiCliServerDiscardHeaders(sapi_headers *core.SapiHeaders) int {
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}
func SapiCliServerSendHeaders(sapi_headers *core.SapiHeaders) int {
	var client *PhpCliServerClient = core.SG__().server_context
	var buffer zend.SmartStr = zend.MakeSmartStr(0)
	var h *core.SapiHeader
	var pos zend.ZendLlistPosition
	if client == nil || core.SG__().request_info.no_headers {
		return core.SAPI_HEADER_SENT_SUCCESSFULLY
	}
	if core.SG__().sapi_headers.http_status_line {
		buffer.AppendString(b.CastStrAuto(core.SG__().sapi_headers.http_status_line))
		buffer.AppendString("\r\n")
	} else {
		AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), core.SG__().sapi_headers.http_response_code, 0)
	}
	AppendEssentialHeaders(&buffer, client, 0)
	h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(sapi_headers.GetHeaders(), &pos))
	for h != nil {
		if h.GetHeaderLen() != 0 {
			buffer.AppendString(b.CastStr(h.GetHeader(), h.GetHeaderLen()))
			buffer.AppendString("\r\n")
		}
		h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(sapi_headers.GetHeaders(), &pos))
	}
	buffer.AppendString("\r\n")
	PhpCliServerClientSendThrough(client, buffer.GetS().GetVal(), buffer.GetS().GetLen())
	buffer.Free()
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}
func SapiCliServerReadCookies() *byte {
	var client *PhpCliServerClient = core.SG__().server_context
	var val *byte
	if nil == b.Assign(&val, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "cookie", b.SizeOf("\"cookie\"")-1)) {
		return nil
	}
	return val
}
func SapiCliServerReadPost(buf *byte, count_bytes int) int {
	var client *PhpCliServerClient = core.SG__().server_context
	if client.GetRequest().GetContent() != nil {
		var content_len int = client.GetRequest().GetContentLen()
		var nbytes_copied int = MIN(client.GetPostReadOffset()+count_bytes, content_len) - client.GetPostReadOffset()
		memmove(buf, client.GetRequest().GetContent()+client.GetPostReadOffset(), nbytes_copied)
		client.SetPostReadOffset(client.GetPostReadOffset() + nbytes_copied)
		return nbytes_copied
	}
	return 0
}
func SapiCliServerRegisterVariable(track_vars_array *zend.Zval, key *byte, val *byte) {
	var new_val *byte = (*byte)(val)
	var new_val_len int
	if nil == val {
		return
	}
	if core.SM__().GetInputFilter()(core.PARSE_SERVER, (*byte)(key), &new_val, strlen(val), &new_val_len) != 0 {
		core.PhpRegisterVariableSafe((*byte)(key), new_val, new_val_len, track_vars_array)
	}
}
func SapiCliServerRegisterEntryCb(entry **byte, num_args int, args ...any, hash_key *zend.ZendHashKey) int {
	var track_vars_array *zend.Zval = __va_arg(args, (*zend.Zval)(_))
	if hash_key.GetKey() != nil {
		var real_key *byte
		var key *byte
		var i uint32
		key = zend.Estrndup(hash_key.GetKey().GetVal(), hash_key.GetKey().GetLen())
		for i = 0; i < hash_key.GetKey().GetLen(); i++ {
			if key[i] == '-' {
				key[i] = '_'
			} else {
				key[i] = toupper(key[i])
			}
		}
		core.Spprintf(&real_key, 0, "%s_%s", "HTTP", key)
		if strcmp(key, "CONTENT_TYPE") == 0 || strcmp(key, "CONTENT_LENGTH") == 0 {
			SapiCliServerRegisterVariable(track_vars_array, key, *entry)
		}
		SapiCliServerRegisterVariable(track_vars_array, real_key, *entry)
		zend.Efree(key)
		zend.Efree(real_key)
	}
	return zend.ZEND_HASH_APPLY_KEEP
}
func SapiCliServerRegisterVariables(track_vars_array *zend.Zval) {
	var client *PhpCliServerClient = core.SG__().server_context
	SapiCliServerRegisterVariable(track_vars_array, "DOCUMENT_ROOT", client.GetServer().GetDocumentRoot())
	var tmp *byte
	if b.Assign(&tmp, strrchr(client.GetAddrStr(), ':')) {
		var addr []byte
		var port []byte
		var addr_start *byte = client.GetAddrStr()
		var addr_end *byte = tmp
		if addr_start[0] == '[' {
			addr_start++
		}
		if addr_end[-1] == ']' {
			addr_end--
		}
		strncpy(port, tmp+1, 8)
		port[7] = '0'
		strncpy(addr, addr_start, addr_end-addr_start)
		addr[addr_end-addr_start] = '0'
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_ADDR", addr)
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_PORT", port)
	} else {
		SapiCliServerRegisterVariable(track_vars_array, "REMOTE_ADDR", client.GetAddrStr())
	}
	var tmp *byte
	core.Spprintf(&tmp, 0, "PHP %s Development Server", core.PHP_VERSION)
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_SOFTWARE", tmp)
	zend.Efree(tmp)
	var tmp *byte
	core.Spprintf(&tmp, 0, "HTTP/%d.%d", client.GetRequest().GetProtocolVersion()/100, client.GetRequest().GetProtocolVersion()%100)
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_PROTOCOL", tmp)
	zend.Efree(tmp)
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_NAME", client.GetServer().GetHost())
	var tmp *byte
	core.Spprintf(&tmp, 0, "%i", client.GetServer().GetPort())
	SapiCliServerRegisterVariable(track_vars_array, "SERVER_PORT", tmp)
	zend.Efree(tmp)
	SapiCliServerRegisterVariable(track_vars_array, "REQUEST_URI", client.GetRequest().GetRequestUri())
	SapiCliServerRegisterVariable(track_vars_array, "REQUEST_METHOD", core.SG__().request_info.request_method)
	SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_NAME", client.GetRequest().GetVpath())
	if core.SG__().request_info.path_translated {
		SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_FILENAME", core.SG__().request_info.path_translated)
	} else if client.GetServer().GetRouter() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_FILENAME", client.GetServer().GetRouter())
	}
	if client.GetRequest().GetPathInfo() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "PATH_INFO", client.GetRequest().GetPathInfo())
	}
	if client.GetRequest().GetPathInfoLen() != 0 {
		var tmp *byte
		core.Spprintf(&tmp, 0, "%s%s", client.GetRequest().GetVpath(), client.GetRequest().GetPathInfo())
		SapiCliServerRegisterVariable(track_vars_array, "PHP_SELF", tmp)
		zend.Efree(tmp)
	} else {
		SapiCliServerRegisterVariable(track_vars_array, "PHP_SELF", client.GetRequest().GetVpath())
	}
	if client.GetRequest().GetQueryString() != nil {
		SapiCliServerRegisterVariable(track_vars_array, "QUERY_STRING", client.GetRequest().GetQueryString())
	}
	zend.ZendHashApplyWithArguments(client.GetRequest().GetHeaders(), zend.ApplyFuncArgsT(SapiCliServerRegisterEntryCb), 1, track_vars_array)
}
func SapiCliServerLogWrite(type_ int, msg *byte) {
	var buf []byte
	if PhpCliServerLogLevel < type_ {
		return
	}
	if PhpCliServerGetSystemTime(buf) != 0 {
		memmove(buf, "unknown time, can't be fetched", b.SizeOf("\"unknown time, can't be fetched\""))
	} else {
		var l int = strlen(buf)
		if l > 0 {
			buf[l-1] = '0'
		} else {
			memmove(buf, "unknown", b.SizeOf("\"unknown\""))
		}
	}
	if PhpCliServerWorkersMax > 1 {
		log.Printf("[%ld] [%s] %s\n", long(getpid()), buf, msg)
	} else {
		log.Printf("[%s] %s\n", buf, msg)
	}
}
func PhpCliServerPollerCtor(poller *PhpCliServerPoller) int {
	FD_ZERO(poller.GetRfds())
	FD_ZERO(poller.GetWfds())
	poller.SetMaxFd(-1)
	return zend.SUCCESS
}
func PhpCliServerPollerAdd(poller *PhpCliServerPoller, mode int, fd core.PhpSocketT) {
	if (mode & POLLIN) != 0 {
		core.PHP_SAFE_FD_SET(fd, poller.GetRfds())
	}
	if (mode & POLLOUT) != 0 {
		core.PHP_SAFE_FD_SET(fd, poller.GetWfds())
	}
	if fd > poller.GetMaxFd() {
		poller.SetMaxFd(fd)
	}
}
func PhpCliServerChunkSize(chunk *PhpCliServerChunk) int {
	switch chunk.GetType() {
	case PHP_CLI_SERVER_CHUNK_HEAP:
		return chunk.GetDataHeapLen()
	case PHP_CLI_SERVER_CHUNK_IMMORTAL:
		return chunk.GetDataImmortalLen()
	}
	return 0
}
func PhpCliServerChunkDtor(chunk *PhpCliServerChunk) {
	switch chunk.GetType() {
	case PHP_CLI_SERVER_CHUNK_HEAP:
		if chunk.GetBlock() != chunk {
			zend.Pefree(chunk.GetBlock(), 1)
		}
		break
	case PHP_CLI_SERVER_CHUNK_IMMORTAL:
		break
	}
}
func PhpCliServerBufferDtor(buffer *PhpCliServerBuffer) {
	var chunk *PhpCliServerChunk
	var next *PhpCliServerChunk
	for chunk = buffer.GetFirst(); chunk != nil; chunk = next {
		next = chunk.GetNext()
		PhpCliServerChunkDtor(chunk)
		zend.Pefree(chunk, 1)
	}
}
func PhpCliServerBufferCtor(buffer *PhpCliServerBuffer) {
	buffer.SetFirst(nil)
	buffer.SetLast(nil)
}
func PhpCliServerBufferAppend(buffer *PhpCliServerBuffer, chunk *PhpCliServerChunk) {
	var last *PhpCliServerChunk
	for last = chunk; last.GetNext() != nil; last = last.GetNext() {

	}
	if buffer.GetLast() == nil {
		buffer.SetFirst(chunk)
	} else {
		buffer.GetLast().SetNext(chunk)
	}
	buffer.SetLast(last)
}
func PhpCliServerBufferPrepend(buffer *PhpCliServerBuffer, chunk *PhpCliServerChunk) {
	var last *PhpCliServerChunk
	for last = chunk; last.GetNext() != nil; last = last.GetNext() {

	}
	last.SetNext(buffer.GetFirst())
	if buffer.GetLast() == nil {
		buffer.SetLast(last)
	}
	buffer.SetFirst(chunk)
}
func PhpCliServerBufferSize(buffer *PhpCliServerBuffer) int {
	var chunk *PhpCliServerChunk
	var retval int = 0
	for chunk = buffer.GetFirst(); chunk != nil; chunk = chunk.GetNext() {
		retval += PhpCliServerChunkSize(chunk)
	}
	return retval
}
func PhpCliServerChunkHeapNew(block any, buf *byte, len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = zend.Pemalloc(b.SizeOf("php_cli_server_chunk"), 1)
	chunk.SetType(PHP_CLI_SERVER_CHUNK_HEAP)
	chunk.SetNext(nil)
	chunk.SetBlock(block)
	chunk.SetDataHeapP(buf)
	chunk.SetDataHeapLen(len_)
	return chunk
}
func PhpCliServerChunkHeapNewSelfContained(len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = zend.Pemalloc(b.SizeOf("php_cli_server_chunk")+len_, 1)
	chunk.SetType(PHP_CLI_SERVER_CHUNK_HEAP)
	chunk.SetNext(nil)
	chunk.SetBlock(chunk)
	chunk.SetDataHeapP((*byte)(chunk + 1))
	chunk.SetDataHeapLen(len_)
	return chunk
}
func PhpCliServerContentSenderDtor(sender *PhpCliServerContentSender) {
	PhpCliServerBufferDtor(sender.GetBuffer())
}
func PhpCliServerContentSenderCtor(sender *PhpCliServerContentSender) {
	PhpCliServerBufferCtor(sender.GetBuffer())
}
func PhpCliServerContentSenderSend(sender *PhpCliServerContentSender, fd core.PhpSocketT, nbytes_sent_total *int) int {
	var chunk *PhpCliServerChunk
	var next *PhpCliServerChunk
	var _nbytes_sent_total int = 0
	for chunk = sender.GetBuffer().GetFirst(); chunk != nil; chunk = next {
		var nbytes_sent ssize_t
		next = chunk.GetNext()
		switch chunk.GetType() {
		case PHP_CLI_SERVER_CHUNK_HEAP:
			nbytes_sent = send(fd, chunk.GetDataHeapP(), chunk.GetDataHeapLen(), 0)
			if nbytes_sent < 0 {
				*nbytes_sent_total = _nbytes_sent_total
				return core.PhpSocketErrno()
			} else if nbytes_sent == ssize_t(chunk.GetDataHeapLen()) {
				PhpCliServerChunkDtor(chunk)
				zend.Pefree(chunk, 1)
				sender.GetBuffer().SetFirst(next)
				if next == nil {
					sender.GetBuffer().SetLast(nil)
				}
			} else {
				chunk.SetDataHeapP(chunk.GetDataHeapP() + nbytes_sent)
				chunk.SetDataHeapLen(chunk.GetDataHeapLen() - nbytes_sent)
			}
			_nbytes_sent_total += nbytes_sent
			break
		case PHP_CLI_SERVER_CHUNK_IMMORTAL:
			nbytes_sent = send(fd, chunk.GetDataImmortalP(), chunk.GetDataImmortalLen(), 0)
			if nbytes_sent < 0 {
				*nbytes_sent_total = _nbytes_sent_total
				return core.PhpSocketErrno()
			} else if nbytes_sent == ssize_t(chunk.GetDataImmortalLen()) {
				PhpCliServerChunkDtor(chunk)
				zend.Pefree(chunk, 1)
				sender.GetBuffer().SetFirst(next)
				if next == nil {
					sender.GetBuffer().SetLast(nil)
				}
			} else {
				chunk.SetDataImmortalP(chunk.GetDataImmortalP() + nbytes_sent)
				chunk.SetDataImmortalLen(chunk.GetDataImmortalLen() - nbytes_sent)
			}
			_nbytes_sent_total += nbytes_sent
			break
		}
	}
	*nbytes_sent_total = _nbytes_sent_total
	return 0
}
func PhpCliServerContentSenderPull(sender *PhpCliServerContentSender, fd int, nbytes_read *int) int {
	var _nbytes_read ssize_t
	var chunk *PhpCliServerChunk = PhpCliServerChunkHeapNewSelfContained(131072)
	_nbytes_read = read(fd, chunk.GetDataHeapP(), chunk.GetDataHeapLen())
	if _nbytes_read < 0 {
		if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
			var errstr *byte = GetLastError()
			PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "%s", errstr)
			zend.Pefree(errstr, 1)
		}
		PhpCliServerChunkDtor(chunk)
		zend.Pefree(chunk, 1)
		return 1
	}
	chunk.SetDataHeapLen(_nbytes_read)
	PhpCliServerBufferAppend(sender.GetBuffer(), chunk)
	*nbytes_read = _nbytes_read
	return 0
}
func PhpCliServerLogResponse(client *PhpCliServerClient, status int, message *byte) {
	var color int = 0
	var effective_status int = status
	var basic_buf *byte
	var message_buf *byte = ""
	var error_buf *byte = ""
	var append_error_message zend.ZendBool = 0
	if core.PG(last_error_message) {
		switch core.PG(last_error_type) {
		case zend.E_ERROR:

		case zend.E_CORE_ERROR:

		case zend.E_COMPILE_ERROR:

		case zend.E_USER_ERROR:

		case zend.E_PARSE:
			if status == 200 {

				/* the status code isn't changed by a fatal error, so fake it */

				effective_status = 500

				/* the status code isn't changed by a fatal error, so fake it */

			}
			append_error_message = 1
			break
		}
	}

	/* basic */

	core.Spprintf(&basic_buf, 0, "%s [%d]: %s %s", client.GetAddrStr(), status, core.SG__().request_info.request_method, client.GetRequest().GetRequestUri())
	if basic_buf == nil {
		return
	}

	/* message */

	if message != nil {
		core.Spprintf(&message_buf, 0, " - %s", message)
		if message_buf == nil {
			zend.Efree(basic_buf)
			return
		}
	}

	/* error */

	if append_error_message != 0 {
		core.Spprintf(&error_buf, 0, " - %s in %s on line %d", core.PG(last_error_message), core.PG(last_error_file), core.PG(last_error_lineno))
		if error_buf == nil {
			zend.Efree(basic_buf)
			if message != nil {
				zend.Efree(message_buf)
			}
			return
		}
	}
	if color != 0 {
		PhpCliServerLogf(PHP_CLI_SERVER_LOG_MESSAGE, "x1b[3%dm%s%s%sx1b[0m", color, basic_buf, message_buf, error_buf)
	} else {
		PhpCliServerLogf(PHP_CLI_SERVER_LOG_MESSAGE, "%s%s%s", basic_buf, message_buf, error_buf)
	}
	zend.Efree(basic_buf)
	if message != nil {
		zend.Efree(message_buf)
	}
	if append_error_message != 0 {
		zend.Efree(error_buf)
	}
}
func PhpCliServerLogf(type_ int, format string, _ ...any) {
	var buf *byte = nil
	var ap va_list
	if PhpCliServerLogLevel < type_ {
		return
	}
	va_start(ap, format)
	core.Vspprintf(&buf, 0, format, ap)
	va_end(ap)
	if buf == nil {
		return
	}
	SapiCliServerLogWrite(type_, buf)
	zend.Efree(buf)
}
func PhpNetworkListenSocket(host *byte, port *int, socktype int, af *int, socklen *socklen_t, errstr **zend.ZendString) core.PhpSocketT {
	var retval core.PhpSocketT = core.SOCK_ERR
	var err int = 0
	var sa *__struct__sockaddr = nil
	var p **__struct__sockaddr
	var sal **__struct__sockaddr
	var num_addrs int = core.PhpNetworkGetaddresses(host, socktype, &sal, errstr)
	if num_addrs == 0 {
		return -1
	}
	for p = sal; (*p) != nil; p++ {
		if sa != nil {
			zend.Pefree(sa, 1)
			sa = nil
		}
		retval = socket(p.sa_family, socktype, 0)
		if retval == core.SOCK_ERR {
			continue
		}
		switch p.sa_family {
		case AF_INET6:
			sa = zend.Pemalloc(b.SizeOf("struct sockaddr_in6"), 1)
			*((*__struct__sockaddr_in6)(sa)) = *((*__struct__sockaddr_in6)(*p))
			(*__struct__sockaddr_in6)(sa).sin6_port = htons(*port)
			*socklen = b.SizeOf("struct sockaddr_in6")
			break
		case AF_INET:
			sa = zend.Pemalloc(b.SizeOf("struct sockaddr_in"), 1)
			*((*__struct__sockaddr_in)(sa)) = *((*__struct__sockaddr_in)(*p))
			(*__struct__sockaddr_in)(sa).sin_port = htons(*port)
			*socklen = b.SizeOf("struct sockaddr_in")
			break
		default:

			/* Unknown family */

			*socklen = 0
			core.Closesocket(retval)
			continue
		}
		if bind(retval, sa, *socklen) == core.SOCK_CONN_ERR {
			err = core.PhpSocketErrno()
			if err == SOCK_EINVAL || err == SOCK_EADDRINUSE {
				goto out
			}
			core.Closesocket(retval)
			retval = core.SOCK_ERR
			continue
		}
		err = 0
		*af = sa.sa_family
		if (*port) == 0 {
			if getsockname(retval, sa, socklen) {
				err = core.PhpSocketErrno()
				goto out
			}
			switch sa.sa_family {
			case AF_INET6:
				*port = ntohs((*__struct__sockaddr_in6)(sa).sin6_port)
				break
			case AF_INET:
				*port = ntohs((*__struct__sockaddr_in)(sa).sin_port)
				break
			}
		}
		break
	}
	if retval == core.SOCK_ERR {
		goto out
	}
	if listen(retval, SOMAXCONN) {
		err = core.PhpSocketErrno()
		goto out
	}
out:
	if sa != nil {
		zend.Pefree(sa, 1)
	}
	if sal != nil {
		core.PhpNetworkFreeaddresses(sal)
	}
	if err != 0 {
		if zend.ZEND_VALID_SOCKET(retval) {
			core.Closesocket(retval)
		}
		if errstr != nil {
			*errstr = core.PhpSocketErrorStr(err)
		}
		return core.SOCK_ERR
	}
	return retval
}
func PhpCliServerRequestDtor(req *PhpCliServerRequest) {
	if req.GetRequestUri() != nil {
		zend.Pefree(req.GetRequestUri(), 1)
	}
	if req.GetVpath() != nil {
		zend.Pefree(req.GetVpath(), 1)
	}
	if req.GetPathTranslated() != nil {
		zend.Pefree(req.GetPathTranslated(), 1)
	}
	if req.GetPathInfo() != nil {
		zend.Pefree(req.GetPathInfo(), 1)
	}
	if req.GetQueryString() != nil {
		zend.Pefree(req.GetQueryString(), 1)
	}
	req.GetHeaders().Destroy()
	req.GetHeadersOriginalCase().Destroy()
	if req.GetContent() != nil {
		zend.Pefree(req.GetContent(), 1)
	}
}
func PhpCliServerClientReadRequest(client *PhpCliServerClient, request *http.Request) int {
	client.request.ReadRequest(request)

}
func PhpCliServerClientSendThrough(client *PhpCliServerClient, str *byte, str_len int) int {
	var tv __struct__timeval = __struct__timeval{10, 0}
	var nbytes_left ssize_t = ssize_t(str_len)
	for {
		var nbytes_sent ssize_t
		nbytes_sent = send(client.GetSock(), str+str_len-nbytes_left, nbytes_left, 0)
		if nbytes_sent < 0 {
			var err int = core.PhpSocketErrno()
			if err == SOCK_EAGAIN {
				var nfds int = core.PhpPollfdFor(client.GetSock(), POLLOUT, &tv)
				if nfds > 0 {
					continue
				} else if nfds < 0 {

					/* error */

					core.PhpHandleAbortedConnection()
					return nbytes_left
				} else {

					/* timeout */

					core.PhpHandleAbortedConnection()
					return nbytes_left
				}
			} else {
				core.PhpHandleAbortedConnection()
				return nbytes_left
			}
		}
		nbytes_left -= nbytes_sent
		if nbytes_left <= 0 {
			break
		}
	}
	return str_len
}
func PhpCliServerClientPopulateRequestInfo(client *PhpCliServerClient, request_info *core.SapiRequestInfo) {
	var val *byte
	request_info.SetRequestMethod(PhpHttpMethodStr(client.GetRequest().GetRequestMethod()))
	request_info.SetProtoNum(client.GetRequest().GetProtocolVersion())
	request_info.SetRequestUri(client.GetRequest().GetRequestUri())
	request_info.SetPathTranslated(client.GetRequest().GetPathTranslated())
	request_info.SetQueryString(client.GetRequest().GetQueryString())
	request_info.SetContentLength(client.GetRequest().GetContentLen())
	request_info.SetAuthDigest(nil)
	request_info.SetAuthPassword(request_info.GetAuthDigest())
	request_info.SetAuthUser(request_info.GetAuthPassword())
	if nil != b.Assign(&val, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "content-type", b.SizeOf("\"content-type\"")-1)) {
		request_info.SetContentType(val)
	}
}
func DestroyRequestInfo(request_info *core.SapiRequestInfo) {}
func PhpCliServerClientCtor(client *PhpCliServerClient, server *PhpCliServer, client_sock core.PhpSocketT, addr *__struct__sockaddr, addr_len socklen_t) {
	client.SetServer(server)
	client.SetSock(client_sock)
	client.SetAddr(addr)
	client.SetAddrLen(addr_len)
	var addr_str *zend.ZendString = 0
	core.PhpNetworkPopulateNameFromSockaddr(addr, addr_len, &addr_str, nil, 0)
	client.SetAddrStr(zend.Pestrndup(addr_str.GetVal(), addr_str.GetLen(), 1))
	client.SetAddrStrLen(addr_str.GetLen())
	zend.ZendStringReleaseEx(addr_str, 0)
	PhpHttpParserInit(client.GetParser(), PHP_HTTP_REQUEST)
	client.SetRequestRead(0)
	client.SetLastHeaderElement(HEADER_NONE)
	client.SetCurrentHeaderName(nil)
	client.SetCurrentHeaderNameLen(0)
	client.SetCurrentHeaderNameAllocated(0)
	client.SetCurrentHeaderValue(nil)
	client.SetCurrentHeaderValueLen(0)
	client.SetPostReadOffset(0)
	client.request.Ctor()
	client.SetContentSenderInitialized(0)
	client.SetFileFd(-1)
}
func PhpCliServerClientDtor(client *PhpCliServerClient) {
	PhpCliServerRequestDtor(client.GetRequest())
	if client.GetFileFd() >= 0 {
		close(client.GetFileFd())
		client.SetFileFd(-1)
	}
	zend.Pefree(client.GetAddr(), 1)
	zend.Pefree(client.GetAddrStr(), 1)
	if client.GetContentSenderInitialized() != 0 {
		PhpCliServerContentSenderDtor(client.GetContentSender())
	}
}
func PhpCliServerCloseConnection(server *PhpCliServer, client *PhpCliServerClient) {
	PhpCliServerLogf(PHP_CLI_SERVER_LOG_MESSAGE, "%s Closing", client.GetAddrStr())
	zend.ZendHashIndexDel(server.GetClients(), client.GetSock())
}
func PhpCliServerSendErrorPage(server *PhpCliServer, client *PhpCliServerClient, status int) int {
	var escaped_request_uri *zend.ZendString = nil
	var status_string *byte = GetStatusString(status)
	var content_template string = GetTemplateString(status)
	var errstr *byte = GetLastError()
	r.Assert(status_string != nil)
	PhpCliServerContentSenderCtor(client.GetContentSender())
	client.SetContentSenderInitialized(1)
	escaped_request_uri = standard.PhpEscapeHtmlEntitiesEx((*uint8)(client.GetRequest().GetRequestUri()), client.GetRequest().GetRequestUriLen(), 0, standard.ENT_QUOTES, nil, 0)
	var prologue_template []byte = "<!doctype html><html><head><title>%d %s</title>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_heap_new_self_contained(strlen(prologue_template) + 3 + strlen(status_string) + 1)
	if chunk == nil {
		goto fail
	}
	core.Snprintf(chunk.GetDataHeapP(), chunk.GetDataHeapLen(), prologue_template, status, status_string)
	chunk.SetDataHeapLen(strlen(chunk.GetDataHeapP()))
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(php_cli_server_css, b.SizeOf("php_cli_server_css")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	var template []byte = "</head><body>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(template, b.SizeOf("template")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	var chunk *php_cli_server_chunk = php_cli_server_chunk_heap_new_self_contained(strlen(content_template) + ZSTR_LEN(escaped_request_uri) + 3 + strlen(status_string) + 1)
	if chunk == nil {
		goto fail
	}
	core.Snprintf(chunk.GetDataHeapP(), chunk.GetDataHeapLen(), content_template, status_string, escaped_request_uri.GetVal())
	chunk.SetDataHeapLen(strlen(chunk.GetDataHeapP()))
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	var epilogue_template []byte = "</body></html>"
	var chunk *php_cli_server_chunk = php_cli_server_chunk_immortal_new(epilogue_template, b.SizeOf("epilogue_template")-1)
	if chunk == nil {
		goto fail
	}
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	var chunk *PhpCliServerChunk
	var buffer zend.SmartStr
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.GetS() == nil {

		/* out of memory */

		goto fail

		/* out of memory */

	}
	AppendEssentialHeaders(&buffer, client, 1)
	buffer.AppendString("Content-Type: text/html; charset=UTF-8\r\n")
	buffer.AppendString("Content-Length: ")
	buffer.AppendUlong(PhpCliServerBufferSize(client.GetContentSender().GetBuffer()))
	buffer.AppendString("\r\n")
	buffer.AppendString("\r\n")
	chunk = PhpCliServerChunkHeapNew(buffer.GetS(), buffer.GetS().GetVal(), buffer.GetS().GetLen())
	if chunk == nil {
		buffer.Free()
		goto fail
	}
	PhpCliServerBufferPrepend(client.GetContentSender().GetBuffer(), chunk)
	PhpCliServerLogResponse(client, status, b.Cond(errstr != nil, errstr, "?"))
	PhpCliServerPollerAdd(server.GetPoller(), POLLOUT, client.GetSock())
	if errstr != nil {
		zend.Pefree(errstr, 1)
	}
	zend.ZendStringFree(escaped_request_uri)
	return zend.SUCCESS
fail:
	if errstr != nil {
		zend.Pefree(errstr, 1)
	}
	zend.ZendStringFree(escaped_request_uri)
	return zend.FAILURE
}
func PhpCliServerDispatchScript(server *PhpCliServer, client *PhpCliServerClient) int {
	if strlen(client.GetRequest().GetPathTranslated()) != client.GetRequest().GetPathTranslatedLen() {

		/* can't handle paths that contain nul bytes */

		return PhpCliServerSendErrorPage(server, client, 400)

		/* can't handle paths that contain nul bytes */

	}
	var zfd zend.ZendFileHandle
	zend.ZendStreamInitFilename(&zfd, core.SG__().request_info.path_translated)
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		core.PhpExecuteScript(&zfd)
	}
	zend.EG__().SetBailout(__orig_bailout)
	PhpCliServerLogResponse(client, core.SG__().sapi_headers.http_response_code, nil)
	return zend.SUCCESS
}
func PhpCliServerBeginSendStatic(server *PhpCliServer, client *PhpCliServerClient) int {
	var fd int
	var status int = 200
	if client.GetRequest().GetPathTranslated() != nil && strlen(client.GetRequest().GetPathTranslated()) != client.GetRequest().GetPathTranslatedLen() {

		/* can't handle paths that contain nul bytes */

		return PhpCliServerSendErrorPage(server, client, 400)

		/* can't handle paths that contain nul bytes */

	}
	if client.GetRequest().GetPathTranslated() != nil {
		fd = open(client.GetRequest().GetPathTranslated(), O_RDONLY)
	} else {
		fd = -1
	}
	if fd < 0 {
		return PhpCliServerSendErrorPage(server, client, 404)
	}
	PhpCliServerContentSenderCtor(client.GetContentSender())
	client.SetContentSenderInitialized(1)
	client.SetFileFd(fd)
	var chunk *PhpCliServerChunk
	var buffer zend.SmartStr = zend.MakeSmartStr(0)
	var mime_type *byte = GetMimeType(server, client.GetRequest().GetExt(), client.GetRequest().GetExtLen())
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.GetS() == nil {

		/* out of memory */

		PhpCliServerLogResponse(client, 500, nil)
		return zend.FAILURE
	}
	AppendEssentialHeaders(&buffer, client, 1)
	if mime_type != nil {
		buffer.AppendString("Content-Type: ")
		buffer.AppendString(b.CastStrAuto(mime_type))
		if strncmp(mime_type, "text/", 5) == 0 {
			buffer.AppendString("; charset=UTF-8")
		}
		buffer.AppendString("\r\n")
	}
	buffer.AppendString("Content-Length: ")
	buffer.AppendUlong(client.GetRequest().GetSb().st_size)
	buffer.AppendString("\r\n")
	buffer.AppendString("\r\n")
	chunk = PhpCliServerChunkHeapNew(buffer.GetS(), buffer.GetS().GetVal(), buffer.GetS().GetLen())
	if chunk == nil {
		buffer.Free()
		PhpCliServerLogResponse(client, 500, nil)
		return zend.FAILURE
	}
	PhpCliServerBufferAppend(client.GetContentSender().GetBuffer(), chunk)
	PhpCliServerLogResponse(client, 200, nil)
	PhpCliServerPollerAdd(server.GetPoller(), POLLOUT, client.GetSock())
	return zend.SUCCESS
}
func PhpCliServerRequestStartup(server *PhpCliServer, client *PhpCliServerClient) int {
	var auth *byte
	PhpCliServerClientPopulateRequestInfo(client, &(core.SG__().request_info))
	if nil != b.Assign(&auth, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "authorization", b.SizeOf("\"authorization\"")-1)) {
		core.PhpHandleAuthData(auth)
	}
	core.SG__().sapi_headers.http_response_code = 200
	if zend.FAILURE == core.PhpRequestStartup() {

		/* should never be happen */

		DestroyRequestInfo(&(core.SG__().request_info))
		return zend.FAILURE
	}
	core.PG(during_request_startup) = 0
	return zend.SUCCESS
}
func PhpCliServerRequestShutdown(server *PhpCliServer, client *PhpCliServerClient) int {
	core.PhpRequestShutdown(0)
	PhpCliServerCloseConnection(server, client)
	DestroyRequestInfo(&(core.SG__().request_info))
	core.SG__().server_context = nil
	core.SG__().rfc1867_uploaded_files = nil
	return zend.SUCCESS
}
func PhpCliServerDispatchRouter(server *PhpCliServer, client *PhpCliServerClient) int {
	var decline int = 0
	var zfd zend.ZendFileHandle
	var old_cwd *byte
	old_cwd = zend.DoAlloca(core.MAXPATHLEN, use_heap)
	old_cwd[0] = '0'
	core.PhpIgnoreValue(zend.VCWD_GETCWD(old_cwd, core.MAXPATHLEN-1))
	zend.ZendStreamInitFilename(&zfd, server.GetRouter())
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		var retval zend.Zval
		retval.SetUndef()
		if zend.SUCCESS == zend.ZendExecuteScripts(zend.ZEND_REQUIRE, &retval, 1, &zfd) {
			if retval.GetType() != zend.IS_UNDEF {
				decline = retval.IsType(zend.IS_FALSE)
				zend.ZvalPtrDtor(&retval)
			}
		} else {
			decline = 1
		}
	}
	zend.EG__().SetBailout(__orig_bailout)
	if old_cwd[0] != '0' {
		core.PhpIgnoreValue(zend.VCWD_CHDIR(old_cwd))
	}
	zend.FreeAlloca(old_cwd, use_heap)
	return decline
}
func PhpCliServerDispatch(server *PhpCliServer, client *PhpCliServerClient) int {
	var is_static_file int = 0
	var ext *byte = client.GetRequest().GetExt()
	core.SG__().server_context = client
	var request = client.request

	if strings.ToLower(request.ext) != "php" || request.path_translated == "" {
		is_static_file = 1
	}
	if server.GetRouter() != nil || is_static_file == 0 {
		if zend.FAILURE == PhpCliServerRequestStartup(server, client) {
			core.SG__().server_context = nil
			PhpCliServerCloseConnection(server, client)
			DestroyRequestInfo(&(core.SG__().request_info))
			return zend.SUCCESS
		}
	}
	if server.GetRouter() != nil {
		if PhpCliServerDispatchRouter(server, client) == 0 {
			PhpCliServerRequestShutdown(server, client)
			return zend.SUCCESS
		}
	}
	if is_static_file == 0 {
		if zend.SUCCESS == PhpCliServerDispatchScript(server, client) || zend.SUCCESS != PhpCliServerSendErrorPage(server, client, 500) {
			if core.SG__().sapi_headers.http_response_code == 304 {
				core.SG__().sapi_headers.send_default_content_type = 0
			}
			PhpCliServerRequestShutdown(server, client)
			return zend.SUCCESS
		}
	} else {
		if server.GetRouter() != nil {
			var send_header_func func(*core.SapiHeaders) int
			send_header_func = core.SM__().GetSendHeaders()

			/* do not generate default content type header */

			core.SG__().sapi_headers.send_default_content_type = 0

			/* we don't want headers to be sent */

			core.SM__().SetSendHeaders(SapiCliServerDiscardHeaders)
			core.PhpRequestShutdown(0)
			core.SM__().SetSendHeaders(send_header_func)
			core.SG__().sapi_headers.send_default_content_type = 1
			core.SG__().rfc1867_uploaded_files = nil
		}
		if zend.SUCCESS != PhpCliServerBeginSendStatic(server, client) {
			PhpCliServerCloseConnection(server, client)
		}
		core.SG__().server_context = nil
		return zend.SUCCESS
	}
	core.SG__().server_context = nil
	DestroyRequestInfo(&(core.SG__().request_info))
	return zend.SUCCESS
}
func PhpCliServerDtor(server *PhpCliServer) {
	server.GetClients().Destroy()
	if zend.ZEND_VALID_SOCKET(server.GetServerSock()) {
		core.Closesocket(server.GetServerSock())
	}
	if server.GetHost() != nil {
		zend.Pefree(server.GetHost(), 1)
	}
	if server.GetDocumentRoot() != nil {
		zend.Pefree(server.GetDocumentRoot(), 1)
	}
	if server.GetRouter() != nil {
		zend.Pefree(server.GetRouter(), 1)
	}
	if PhpCliServerWorkersMax > 1 && PhpCliServerWorkers != nil && getpid() == PhpCliServerMaster {
		var php_cli_server_worker zend.ZendLong
		for php_cli_server_worker = 0; php_cli_server_worker < PhpCliServerWorkersMax; php_cli_server_worker++ {
			var php_cli_server_worker_status int
			for {
				if waitpid(PhpCliServerWorkers[php_cli_server_worker], &php_cli_server_worker_status, 0) == zend.FAILURE {

					/* an extremely bad thing happened */

					break

					/* an extremely bad thing happened */

				}
				if !(!(WIFEXITED(php_cli_server_worker_status)) && !(WIFSIGNALED(php_cli_server_worker_status))) {
					break
				}
			}
		}
		zend.Free(PhpCliServerWorkers)
	}
}
func PhpCliServerClientDtorWrapper(zv *zend.Zval) {
	var p *PhpCliServerClient = zv.GetPtr()
	shutdown(p.GetSock(), core.SHUT_RDWR)
	core.Closesocket(p.GetSock())
	PhpCliServerClientDtor(p)
	zend.Pefree(p, 1)
}

// 解析域名地址为 host + port
func parseServerAddr(addr string) (host string, port int, ok bool) {
	addrPair := strings.SplitN(addr, ":", 2)
	if len(addrPair) != 2 {
		return
	}

	// 检查 host
	host = addrPair[0]
	if len(host) == 0 {
		return
	}
	if host[0] == '[' { // 去除 [host] 模式的括号
		if host[len(host)-1] != ']' || len(host) <= 2 {
			return
		}
		host = host[1 : len(host)-1]
	}

	// 检查 port
	port, err := strconv.Atoi(addrPair[1])
	if err != nil {
		return
	}
	if port <= 0 || port > 65535 {
		return
	}

	// 解析成功
	return host, port, true
}

func PhpCliServerCtor(server *PhpCliServer, addr string, document_root string, router string) bool {
	var retval int = zend.SUCCESS
	var errstr *zend.ZendString = nil
	var _router *byte = nil
	var err int = 0
	var server_sock core.PhpSocketT = core.SOCK_ERR

	// 从 addr 解析 host + port
	host, port, ok := parseServerAddr(addr)
	if !ok {
		log.Printf("Invalid address: %s\n", addr)
		retval = zend.FAILURE
		goto out
	}

	server_sock = PhpNetworkListenSocket(host, &port, SOCK_STREAM, server.GetAddressFamily(), server.GetSocklen(), &errstr)
	if server_sock == core.SOCK_ERR {
		PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "Failed to listen on %s:%d (reason: %s)", host, port, b.CondF1(errstr != nil, func() []byte { return errstr.GetVal() }, "?"))
		if errstr != nil {
			zend.ZendStringReleaseEx(errstr, 0)
		}
		retval = zend.FAILURE
		goto out
	}
	server.SetServerSock(server_sock)
	err = PhpCliServerPollerCtor(server.GetPoller())
	if zend.SUCCESS != err {
		goto out
	}
	PhpCliServerPollerAdd(server.GetPoller(), POLLIN, server_sock)
	server.SetHostStr(host)
	server.SetPort(port)
	server.clients = *zend.NewZendArrayEx(0, PhpCliServerClientDtorWrapper, true)
	server.SetDocumentRootStr(document_root)
	server.SetRouterStr(router)
	server.SetIsRunning(1)
out:
	if retval != zend.SUCCESS {
		if server_sock > -1 {
			core.Closesocket(server_sock)
		}
	}
	return retval == zend.SUCCESS
}
func PhpCliServerRecvEventReadRequest(server *PhpCliServer, client *PhpCliServerClient, request *http.Request) int {
	var status int = PhpCliServerClientReadRequest(client, request)
	if status == 1 && client.GetRequest().GetRequestMethod() == PHP_HTTP_NOT_IMPLEMENTED {
		return PhpCliServerSendErrorPage(server, client, 501)
	} else if status == 1 {
		PhpCliServerDispatch(server, client)
	}
	return zend.SUCCESS
}
func PhpCliServerSendEvent(server *PhpCliServer, client *PhpCliServerClient) int {
	if client.GetContentSenderInitialized() != 0 {
		if client.GetFileFd() >= 0 && client.GetContentSender().GetBuffer().GetFirst() == nil {
			var nbytes_read int
			if PhpCliServerContentSenderPull(client.GetContentSender(), client.GetFileFd(), &nbytes_read) != 0 {
				PhpCliServerCloseConnection(server, client)
				return zend.FAILURE
			}
			if nbytes_read == 0 {
				close(client.GetFileFd())
				client.SetFileFd(-1)
			}
		}
		var nbytes_sent int
		var err int = PhpCliServerContentSenderSend(client.GetContentSender(), client.GetSock(), &nbytes_sent)
		if err != 0 && err != SOCK_EAGAIN {
			PhpCliServerCloseConnection(server, client)
			return zend.FAILURE
		}
		if client.GetContentSender().GetBuffer().GetFirst() == nil && client.GetFileFd() < 0 {
			PhpCliServerCloseConnection(server, client)
		}
	}
	return zend.SUCCESS
}
func DoCliServer(optArgs core.OptArgs) int {
	var server_bind_address string
	var document_root string

	for _, optArg := range optArgs.OptionValues {
		switch optArg.Char {
		case 'S':
			server_bind_address = optArg.Value
			break
		case 't':
			document_root = optArg.Value
			break
		case 'q':
			if PhpCliServerLogLevel > 1 {
				PhpCliServerLogLevel--
			}
			break
		}
	}
	if document_root != "" {
		sb, err := os.Stat(document_root)
		if err != nil {
			log.Printf("Directory %s does not exist.\n", document_root)
			return 1
		}
		if !sb.IsDir() {
			log.Printf("%s is not a directory.\n", document_root)
			return 1
		}
	} else {
		if ret, err := os.Getwd(); err == nil {
			document_root = ret
		} else {
			document_root = "."
		}
	}

	var router string
	if len(optArgs.Arguments) > 0 {
		router = optArgs.Arguments[0]
	}

	var Server PhpCliServer
	if !PhpCliServerCtor(&Server, server_bind_address, document_root, router) {
		return 1
	}
	core.SM__().SetPhpinfoAsText(0)
	PhpCliServerLogf(PHP_CLI_SERVER_LOG_PROCESS, "PHP %s Development Server (http://%s) started", core.PHP_VERSION, server_bind_address)
	zend.ZendSignalInit()

	err := Server.Serve()
	if err != nil {
		log.Print(err)
	}
	PhpCliServerDtor(&Server)
	return 0
}
