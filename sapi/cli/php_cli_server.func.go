// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

func CLI_SERVER_G(v int) __auto__ { return CliServerGlobals.v }
func PhpCliServerGetSystemTime(buf *byte) int {
	var tv __struct__timeval
	var tm __struct__tm
	gettimeofday(&tv, nil)

	/* TODO: should be checked for NULL tm/return vaue */

	core.PhpLocaltimeR(tv.tv_sec, &tm)
	core.PhpAsctimeR(&tm, buf)
	return 0
}
func CharPtrDtorP(zv *zend.Zval) { zend.Pefree(zv.GetPtr(), 1) }
func GetLastError() *byte        { return zend.Pestrdup(strerror(errno), 1) }
func StatusComp(a any, b any) int {
	var pa *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(a)
	var pb *core.HttpResponseStatusCodePair = (*core.HttpResponseStatusCodePair)(b)
	if pa.GetCode() < pb.GetCode() {
		return -1
	} else if pa.GetCode() > pb.GetCode() {
		return 1
	}
	return 0
}
func GetStatusString(code int) *byte {
	var needle core.HttpResponseStatusCodePair = core.HttpResponseStatusCodePair{code, nil}
	var result *core.HttpResponseStatusCodePair = nil
	result = bsearch(&needle, core.HttpStatusMap, core.HttpStatusMapLen, b.SizeOf("needle"), StatusComp)
	if result != nil {
		return result.GetStr()
	}

	/* Returning NULL would require complicating append_http_status_line() to
	 * not segfault in that case, so let's just return a placeholder, since RFC
	 * 2616 requires a reason phrase. This is basically what a lot of other Web
	 * servers do in this case anyway. */

	return "Unknown Status Code"

	/* Returning NULL would require complicating append_http_status_line() to
	 * not segfault in that case, so let's just return a placeholder, since RFC
	 * 2616 requires a reason phrase. This is basically what a lot of other Web
	 * servers do in this case anyway. */
}
func GetTemplateString(code int) *byte {
	var e int = b.SizeOf("template_map") / b.SizeOf("php_cli_server_http_response_status_code_pair")
	var s int = 0
	for e != s {
		var c int = MIN((e+s+1)/2, e-1)
		var d int = TemplateMap[c].GetCode()
		if d > code {
			e = c
		} else if d < code {
			s = c
		} else {
			return TemplateMap[c].GetStr()
		}
	}
	return nil
}
func AppendHttpStatusLine(buffer *zend.SmartStr, protocol_version int, response_code int, persistent int) {
	if response_code == 0 {
		response_code = 200
	}
	zend.SmartStrAppendlEx(buffer, "HTTP", 4, persistent)
	zend.SmartStrAppendcEx(buffer, '/', persistent)
	zend.SmartStrAppendLongEx(buffer, protocol_version/100, persistent)
	zend.SmartStrAppendcEx(buffer, '.', persistent)
	zend.SmartStrAppendLongEx(buffer, protocol_version%100, persistent)
	zend.SmartStrAppendcEx(buffer, ' ', persistent)
	zend.SmartStrAppendLongEx(buffer, response_code, persistent)
	zend.SmartStrAppendcEx(buffer, ' ', persistent)
	zend.SmartStrAppendsEx(buffer, GetStatusString(response_code), persistent)
	zend.SmartStrAppendlEx(buffer, "\r\n", 2, persistent)
}
func AppendEssentialHeaders(buffer *zend.SmartStr, client *PhpCliServerClient, persistent int) {
	var val *byte
	var tv __struct__timeval = __struct__timeval{0}
	if nil != b.Assign(&val, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "host", b.SizeOf("\"host\"")-1)) {
		zend.SmartStrAppendsEx(buffer, "Host: ", persistent)
		zend.SmartStrAppendsEx(buffer, val, persistent)
		zend.SmartStrAppendsEx(buffer, "\r\n", persistent)
	}
	if !(gettimeofday(&tv, nil)) {
		var dt *zend.ZendString = php_format_date("D, d M Y H:i:s", b.SizeOf("\"D, d M Y H:i:s\"")-1, tv.tv_sec, 0)
		zend.SmartStrAppendsEx(buffer, "Date: ", persistent)
		zend.SmartStrAppendsEx(buffer, dt.GetVal(), persistent)
		zend.SmartStrAppendsEx(buffer, " GMT\r\n", persistent)
		zend.ZendStringReleaseEx(dt, 0)
	}
	zend.SmartStrAppendlEx(buffer, "Connection: close\r\n", b.SizeOf("\"Connection: close\\r\\n\"")-1, persistent)
}
func GetMimeType(server *PhpCliServer, ext *byte, ext_len int) *byte {
	var ret *byte
	var ext_lower *byte = zend.DoAlloca(ext_len+1, use_heap)
	zend.ZendStrTolowerCopy(ext_lower, ext, ext_len)
	ret = zend.ZendHashStrFindPtr(server.GetExtensionMimeTypes(), ext_lower, ext_len)
	zend.FreeAlloca(ext_lower, use_heap)
	return (*byte)(ret)
}
func ZifApacheRequestHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var client *PhpCliServerClient
	var headers *zend.HashTable
	var key *zend.ZendString
	var value *byte
	var tmp zend.Zval
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	client = core.SG(server_context)
	headers = client.GetRequest().GetHeadersOriginalCase()
	zend.ArrayInitSize(return_value, headers.GetNNumOfElements())
	for {
		var __ht *zend.HashTable = headers
		var _p *zend.Bucket = __ht.GetArData()
		var _end *zend.Bucket = _p + __ht.GetNNumUsed()
		for ; _p != _end; _p++ {
			var _z *zend.Zval = _p.GetVal()

			if _z.IsType(zend.IS_UNDEF) {
				continue
			}
			key = _p.GetKey()
			value = _z.GetPtr()
			zend.ZVAL_STRING(&tmp, value)
			zend.ZendSymtableUpdate(return_value.GetArr(), key, &tmp)
		}
		break
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
func ZifApacheResponseHeaders(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ArrayInit(return_value)
	zend.ZendLlistApplyWithArgument(core.SG(sapi_headers).headers, zend.LlistApplyWithArgFuncT(AddResponseHeader), return_value)
}
func CliServerInitGlobals(cg *ZendCliServerGlobals) { cg.SetColor(0) }
func ZmStartupCliServer(type_ int, module_number int) int {
	CliServerInitGlobals(&CliServerGlobals)
	zend.REGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmShutdownCliServer(type_ int, module_number int) int {
	zend.UNREGISTER_INI_ENTRIES()
	return zend.SUCCESS
}
func ZmInfoCliServer(ZEND_MODULE_INFO_FUNC_ARGS) { zend.DISPLAY_INI_ENTRIES() }
func SapiCliServerStartup(sapi_module *core.sapi_module_struct) int {
	var workers *byte
	if core.PhpModuleStartup(sapi_module, &CliServerModuleEntry, 1) == zend.FAILURE {
		return zend.FAILURE
	}
	if b.Assign(&workers, getenv("PHP_CLI_SERVER_WORKERS")) {
		r.Fprintf(stderr, "platform does not support SO_REUSEPORT, cannot create workers\n")
	}
	return zend.SUCCESS
}
func SapiCliServerUbWrite(str *byte, str_length int) int {
	var client *PhpCliServerClient = core.SG(server_context)
	if client == nil {
		return 0
	}
	return PhpCliServerClientSendThrough(client, str, str_length)
}
func SapiCliServerFlush(server_context any) {
	var client *PhpCliServerClient = server_context
	if client == nil {
		return
	}
	if !(zend.ZEND_VALID_SOCKET(client.GetSock())) {
		core.PhpHandleAbortedConnection()
		return
	}
	if !(core.SG(headers_sent)) {
		core.SapiSendHeaders()
		core.SG(headers_sent) = 1
	}
}
func SapiCliServerDiscardHeaders(sapi_headers *core.SapiHeaders) int {
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}
func SapiCliServerSendHeaders(sapi_headers *core.SapiHeaders) int {
	var client *PhpCliServerClient = core.SG(server_context)
	var buffer zend.SmartStr = zend.SmartStr{0}
	var h *core.SapiHeader
	var pos zend.ZendLlistPosition
	if client == nil || core.SG(request_info).no_headers {
		return core.SAPI_HEADER_SENT_SUCCESSFULLY
	}
	if core.SG(sapi_headers).http_status_line {
		zend.SmartStrAppends(&buffer, core.SG(sapi_headers).http_status_line)
		zend.SmartStrAppendl(&buffer, "\r\n", 2)
	} else {
		AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), core.SG(sapi_headers).http_response_code, 0)
	}
	AppendEssentialHeaders(&buffer, client, 0)
	h = (*core.SapiHeader)(zend.ZendLlistGetFirstEx(sapi_headers.GetHeaders(), &pos))
	for h != nil {
		if h.GetHeaderLen() != 0 {
			zend.SmartStrAppendl(&buffer, h.GetHeader(), h.GetHeaderLen())
			zend.SmartStrAppendl(&buffer, "\r\n", 2)
		}
		h = (*core.SapiHeader)(zend.ZendLlistGetNextEx(sapi_headers.GetHeaders(), &pos))
	}
	zend.SmartStrAppendl(&buffer, "\r\n", 2)
	PhpCliServerClientSendThrough(client, buffer.GetS().GetVal(), buffer.GetS().GetLen())
	zend.SmartStrFree(&buffer)
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}
func SapiCliServerReadCookies() *byte {
	var client *PhpCliServerClient = core.SG(server_context)
	var val *byte
	if nil == b.Assign(&val, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "cookie", b.SizeOf("\"cookie\"")-1)) {
		return nil
	}
	return val
}
func SapiCliServerReadPost(buf *byte, count_bytes int) int {
	var client *PhpCliServerClient = core.SG(server_context)
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
	if core.sapi_module.GetInputFilter()(core.PARSE_SERVER, (*byte)(key), &new_val, strlen(val), &new_val_len) != 0 {
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
	var client *PhpCliServerClient = core.SG(server_context)
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
	SapiCliServerRegisterVariable(track_vars_array, "REQUEST_METHOD", core.SG(request_info).request_method)
	SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_NAME", client.GetRequest().GetVpath())
	if core.SG(request_info).path_translated {
		SapiCliServerRegisterVariable(track_vars_array, "SCRIPT_FILENAME", core.SG(request_info).path_translated)
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
		r.Fprintf(stderr, "[%ld] [%s] %s\n", long(getpid()), buf, msg)
	} else {
		r.Fprintf(stderr, "[%s] %s\n", buf, msg)
	}
}
func SapiCliServerLogMessage(msg *byte, syslog_type_int int) {
	SapiCliServerLogWrite(PHP_CLI_SERVER_LOG_MESSAGE, msg)
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
func PhpCliServerPollerRemove(poller *PhpCliServerPoller, mode int, fd core.PhpSocketT) {
	if (mode & POLLIN) != 0 {
		core.PHP_SAFE_FD_CLR(fd, poller.GetRfds())
	}
	if (mode & POLLOUT) != 0 {
		core.PHP_SAFE_FD_CLR(fd, poller.GetWfds())
	}
	if fd == poller.GetMaxFd() {
		for fd > 0 {
			fd--
			if core.PHP_SAFE_FD_ISSET(fd, poller.GetRfds()) || core.PHP_SAFE_FD_ISSET(fd, poller.GetWfds()) {
				break
			}
		}
		poller.SetMaxFd(fd)
	}
}
func PhpCliServerPollerPoll(poller *PhpCliServerPoller, tv *__struct__timeval) int {
	memmove(poller.GetActiveRfds(), poller.GetRfds(), b.SizeOf("poller -> rfds"))
	memmove(poller.GetActiveWfds(), poller.GetWfds(), b.SizeOf("poller -> wfds"))
	return PhpSelect(poller.GetMaxFd()+1, poller.GetActiveRfds(), poller.GetActiveWfds(), nil, tv)
}
func PhpCliServerPollerIterOnActive(poller *PhpCliServerPoller, opaque any, callback func(_ any, fd core.PhpSocketT, events int) int) int {
	var retval int = zend.SUCCESS
	var fd core.PhpSocketT
	var max_fd core.PhpSocketT = poller.GetMaxFd()
	for fd = 0; fd <= max_fd; fd++ {
		if core.PHP_SAFE_FD_ISSET(fd, poller.GetActiveRfds()) {
			if zend.SUCCESS != callback(opaque, fd, POLLIN) {
				retval = zend.FAILURE
			}
		}
		if core.PHP_SAFE_FD_ISSET(fd, poller.GetActiveWfds()) {
			if zend.SUCCESS != callback(opaque, fd, POLLOUT) {
				retval = zend.FAILURE
			}
		}
	}
	return retval
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
func PhpCliServerChunkImmortalNew(buf *byte, len_ int) *PhpCliServerChunk {
	var chunk *PhpCliServerChunk = zend.Pemalloc(b.SizeOf("php_cli_server_chunk"), 1)
	chunk.SetType(PHP_CLI_SERVER_CHUNK_IMMORTAL)
	chunk.SetNext(nil)
	chunk.SetDataImmortalP(buf)
	chunk.SetDataImmortalLen(len_)
	return chunk
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
func PhpCliIsOutputTty() int {
	if PhpCliOutputIsTty == OUTPUT_NOT_CHECKED {
		PhpCliOutputIsTty = zend.Isatty(STDOUT_FILENO)
	}
	return PhpCliOutputIsTty
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
	if CLI_SERVER_G(color) && PhpCliIsOutputTty() == OUTPUT_IS_TTY {
		if effective_status >= 500 {

			/* server error: red */

			color = 1

			/* server error: red */

		} else if effective_status >= 400 {

			/* client error: yellow */

			color = 3

			/* client error: yellow */

		} else if effective_status >= 200 {

			/* success: green */

			color = 2

			/* success: green */

		}
	}

	/* basic */

	core.Spprintf(&basic_buf, 0, "%s [%d]: %s %s", client.GetAddrStr(), status, core.SG(request_info).request_method, client.GetRequest().GetRequestUri())
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
func PhpCliServerRequestCtor(req *PhpCliServerRequest) int {
	req.SetProtocolVersion(0)
	req.SetRequestUri(nil)
	req.SetRequestUriLen(0)
	req.SetVpath(nil)
	req.SetVpathLen(0)
	req.SetPathTranslated(nil)
	req.SetPathTranslatedLen(0)
	req.SetPathInfo(nil)
	req.SetPathInfoLen(0)
	req.SetQueryString(nil)
	req.SetQueryStringLen(0)
	zend.ZendHashInit(req.GetHeaders(), 0, nil, CharPtrDtorP, 1)
	zend.ZendHashInit(req.GetHeadersOriginalCase(), 0, nil, nil, 1)
	req.SetContent(nil)
	req.SetContentLen(0)
	req.SetExt(nil)
	req.SetExtLen(0)
	return zend.SUCCESS
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
	zend.ZendHashDestroy(req.GetHeaders())
	zend.ZendHashDestroy(req.GetHeadersOriginalCase())
	if req.GetContent() != nil {
		zend.Pefree(req.GetContent(), 1)
	}
}
func PhpCliServerRequestTranslateVpath(request *PhpCliServerRequest, document_root *byte, document_root_len int) {
	var sb zend.ZendStatT
	var index_files []*byte = []*byte{"index.php", "index.html", nil}
	var buf *byte = zend.SafePemalloc(1, request.GetVpathLen(), 1+document_root_len+1+b.SizeOf("\"index.html\""), 1)
	var p *byte = buf
	var prev_path *byte = nil
	var q *byte
	var vpath *byte
	var prev_path_len int = 0
	var is_static_file int = 0
	memmove(p, document_root, document_root_len)
	p += document_root_len
	vpath = p
	if request.GetVpathLen() > 0 && request.GetVpath()[0] != '/' {
		b.PostInc(&(*p)) = zend.DEFAULT_SLASH
	}
	q = request.GetVpath() + request.GetVpathLen()
	for q > request.GetVpath() {
		if b.PostDec(&(*q)) == '.' {
			is_static_file = 1
			break
		}
	}
	memmove(p, request.GetVpath(), request.GetVpathLen())
	p += request.GetVpathLen()
	*p = '0'
	q = p
	for q > buf {
		if !(zend.PhpSysStat(buf, &sb)) {
			if (sb.st_mode & S_IFDIR) != 0 {
				var file **byte = index_files
				if q[-1] != zend.DEFAULT_SLASH {
					b.PostInc(&(*q)) = zend.DEFAULT_SLASH
				}
				for (*file) != nil {
					var l int = strlen(*file)
					memmove(q, *file, l+1)
					if !(zend.PhpSysStat(buf, &sb)) && (sb.st_mode&S_IFREG) != 0 {
						q += l
						break
					}
					file++
				}
				if (*file) == nil || is_static_file != 0 {
					if prev_path != nil {
						zend.Pefree(prev_path, 1)
					}
					zend.Pefree(buf, 1)
					return
				}
			}
			break
		}
		if prev_path != nil {
			zend.Pefree(prev_path, 1)
			*q = zend.DEFAULT_SLASH
		}
		for q > buf && (*(b.PreDec(&q))) != zend.DEFAULT_SLASH {

		}
		prev_path_len = p - q
		prev_path = zend.Pestrndup(q, prev_path_len, 1)
		*q = '0'
	}
	if prev_path != nil {
		request.SetPathInfoLen(prev_path_len)
		request.SetPathInfo(prev_path)
		zend.Pefree(request.GetVpath(), 1)
		request.SetVpath(zend.Pestrndup(vpath, q-vpath, 1))
		request.SetVpathLen(q - vpath)
		request.SetPathTranslated(buf)
		request.SetPathTranslatedLen(q - buf)
	} else {
		zend.Pefree(request.GetVpath(), 1)
		request.SetVpath(zend.Pestrndup(vpath, q-vpath, 1))
		request.SetVpathLen(q - vpath)
		request.SetPathTranslated(buf)
		request.SetPathTranslatedLen(q - buf)
	}
	request.SetSb(sb)
}
func NormalizeVpath(retval **byte, retval_len *int, vpath *byte, vpath_len int, persistent int) {
	var decoded_vpath *byte = nil
	var decoded_vpath_end *byte
	var p *byte
	*retval = nil
	*retval_len = 0
	decoded_vpath = zend.Pestrndup(vpath, vpath_len, persistent)
	if decoded_vpath == nil {
		return
	}
	decoded_vpath_end = decoded_vpath + standard.PhpRawUrlDecode(decoded_vpath, int(vpath_len))
	p = decoded_vpath
	if p < decoded_vpath_end && (*p) == '/' {
		var n *byte = p
		for n < decoded_vpath_end && (*n) == '/' {
			n++
		}
		memmove(b.PreInc(&p), n, decoded_vpath_end-n)
		decoded_vpath_end -= n - p
	}
	for p < decoded_vpath_end {
		var n *byte = p
		for n < decoded_vpath_end && (*n) != '/' {
			n++
		}
		if n-p == 2 && p[0] == '.' && p[1] == '.' {
			if p > decoded_vpath {
				p--
				for {
					if p == decoded_vpath {
						if (*p) == '/' {
							p++
						}
						break
					}
					if (*(b.PreDec(&p))) == '/' {
						p++
						break
					}
				}
			}
			for n < decoded_vpath_end && (*n) == '/' {
				n++
			}
			memmove(p, n, decoded_vpath_end-n)
			decoded_vpath_end -= n - p
		} else if n-p == 1 && p[0] == '.' {
			for n < decoded_vpath_end && (*n) == '/' {
				n++
			}
			memmove(p, n, decoded_vpath_end-n)
			decoded_vpath_end -= n - p
		} else {
			if n < decoded_vpath_end {
				var nn *byte = n
				for nn < decoded_vpath_end && (*nn) == '/' {
					nn++
				}
				p = n + 1
				memmove(p, nn, decoded_vpath_end-nn)
				decoded_vpath_end -= nn - p
			} else {
				p = n
			}
		}
	}
	*decoded_vpath_end = '0'
	*retval = decoded_vpath
	*retval_len = decoded_vpath_end - decoded_vpath
}
func PhpCliServerClientReadRequestOnMessageBegin(parser *PhpHttpParser) int { return 0 }
func PhpCliServerClientReadRequestOnPath(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	var vpath *byte
	var vpath_len int
	if client.GetRequest().GetVpath() != nil {
		return 1
	}
	NormalizeVpath(&vpath, &vpath_len, at, length, 1)
	client.GetRequest().SetVpath(vpath)
	client.GetRequest().SetVpathLen(vpath_len)
	return 0
}
func PhpCliServerClientReadRequestOnQueryString(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetQueryString() == nil {
		client.GetRequest().SetQueryString(zend.Pestrndup(at, length, 1))
		client.GetRequest().SetQueryStringLen(length)
	} else {
		zend.ZEND_ASSERT(length <= PHP_HTTP_MAX_HEADER_SIZE && PHP_HTTP_MAX_HEADER_SIZE-length >= client.GetRequest().GetQueryStringLen())
		client.GetRequest().SetQueryString(zend.Perealloc(client.GetRequest().GetQueryString(), client.GetRequest().GetQueryStringLen()+length+1, 1))
		memcpy(client.GetRequest().GetQueryString()+client.GetRequest().GetQueryStringLen(), at, length)
		client.GetRequest().SetQueryStringLen(client.GetRequest().GetQueryStringLen() + length)
		client.GetRequest().GetQueryString()[client.GetRequest().GetQueryStringLen()] = '0'
	}
	return 0
}
func PhpCliServerClientReadRequestOnUrl(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetRequestUri() == nil {
		client.GetRequest().SetRequestMethod(parser.GetMethod())
		client.GetRequest().SetRequestUri(zend.Pestrndup(at, length, 1))
		client.GetRequest().SetRequestUriLen(length)
	} else {
		zend.ZEND_ASSERT(client.GetRequest().GetRequestMethod() == parser.GetMethod())
		zend.ZEND_ASSERT(length <= PHP_HTTP_MAX_HEADER_SIZE && PHP_HTTP_MAX_HEADER_SIZE-length >= client.GetRequest().GetQueryStringLen())
		client.GetRequest().SetRequestUri(zend.Perealloc(client.GetRequest().GetRequestUri(), client.GetRequest().GetRequestUriLen()+length+1, 1))
		memcpy(client.GetRequest().GetRequestUri()+client.GetRequest().GetRequestUriLen(), at, length)
		client.GetRequest().SetRequestUriLen(client.GetRequest().GetRequestUriLen() + length)
		client.GetRequest().GetRequestUri()[client.GetRequest().GetRequestUriLen()] = '0'
	}
	return 0
}
func PhpCliServerClientReadRequestOnFragment(parser *PhpHttpParser, at *byte, length int) int {
	return 0
}
func PhpCliServerClientSaveHeader(client *PhpCliServerClient) {
	/* strip off the colon */

	var orig_header_name *zend.ZendString = zend.ZendStringInit(client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen(), 1)
	var lc_header_name *zend.ZendString = zend.ZendStringAlloc(client.GetCurrentHeaderNameLen(), 1)
	zend.ZendStrTolowerCopy(lc_header_name.GetVal(), client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())
	zend.GC_MAKE_PERSISTENT_LOCAL(orig_header_name)
	zend.GC_MAKE_PERSISTENT_LOCAL(lc_header_name)
	zend.ZendHashAddPtr(client.GetRequest().GetHeaders(), lc_header_name, client.GetCurrentHeaderValue())
	zend.ZendHashAddPtr(client.GetRequest().GetHeadersOriginalCase(), orig_header_name, client.GetCurrentHeaderValue())
	zend.ZendStringReleaseEx(lc_header_name, 1)
	zend.ZendStringReleaseEx(orig_header_name, 1)
	if client.GetCurrentHeaderNameAllocated() != 0 {
		zend.Pefree(client.GetCurrentHeaderName(), 1)
		client.SetCurrentHeaderNameAllocated(0)
	}
	client.SetCurrentHeaderName(nil)
	client.SetCurrentHeaderNameLen(0)
	client.SetCurrentHeaderValue(nil)
	client.SetCurrentHeaderValueLen(0)
}
func PhpCliServerClientReadRequestOnHeaderField(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_VALUE:
		PhpCliServerClientSaveHeader(client)
	case HEADER_NONE:
		client.SetCurrentHeaderName((*byte)(at))
		client.SetCurrentHeaderNameLen(length)
		break
	case HEADER_FIELD:
		if client.GetCurrentHeaderNameAllocated() != 0 {
			var new_length int = client.GetCurrentHeaderNameLen() + length
			client.SetCurrentHeaderName(zend.Perealloc(client.GetCurrentHeaderName(), new_length+1, 1))
			memcpy(client.GetCurrentHeaderName()+client.GetCurrentHeaderNameLen(), at, length)
			client.GetCurrentHeaderName()[new_length] = '0'
			client.SetCurrentHeaderNameLen(new_length)
		} else {
			var new_length int = client.GetCurrentHeaderNameLen() + length
			var field *byte = zend.Pemalloc(new_length+1, 1)
			memcpy(field, client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())
			memcpy(field+client.GetCurrentHeaderNameLen(), at, length)
			field[new_length] = '0'
			client.SetCurrentHeaderName(field)
			client.SetCurrentHeaderNameLen(new_length)
			client.SetCurrentHeaderNameAllocated(1)
		}
		break
	}
	client.SetLastHeaderElement(HEADER_FIELD)
	return 0
}
func PhpCliServerClientReadRequestOnHeaderValue(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_FIELD:
		client.SetCurrentHeaderValue(zend.Pestrndup(at, length, 1))
		client.SetCurrentHeaderValueLen(length)
		break
	case HEADER_VALUE:
		var new_length int = client.GetCurrentHeaderValueLen() + length
		client.SetCurrentHeaderValue(zend.Perealloc(client.GetCurrentHeaderValue(), new_length+1, 1))
		memcpy(client.GetCurrentHeaderValue()+client.GetCurrentHeaderValueLen(), at, length)
		client.GetCurrentHeaderValue()[new_length] = '0'
		client.SetCurrentHeaderValueLen(new_length)
		break
	case HEADER_NONE:

		// can't happen

		r.Assert(false)
		break
	}
	client.SetLastHeaderElement(HEADER_VALUE)
	return 0
}
func PhpCliServerClientReadRequestOnHeadersComplete(parser *PhpHttpParser) int {
	var client *PhpCliServerClient = parser.GetData()
	switch client.GetLastHeaderElement() {
	case HEADER_NONE:
		break
	case HEADER_FIELD:
		client.SetCurrentHeaderValue(zend.Pemalloc(1, 1))
		client.current_header_value = '0'
		client.SetCurrentHeaderValueLen(0)
	case HEADER_VALUE:
		PhpCliServerClientSaveHeader(client)
		break
	}
	client.SetLastHeaderElement(HEADER_NONE)
	return 0
}
func PhpCliServerClientReadRequestOnBody(parser *PhpHttpParser, at *byte, length int) int {
	var client *PhpCliServerClient = parser.GetData()
	if client.GetRequest().GetContent() == nil {
		client.GetRequest().SetContent(zend.Pemalloc(parser.GetContentLength(), 1))
		client.GetRequest().SetContentLen(0)
	}
	client.GetRequest().SetContent(zend.Perealloc(client.GetRequest().GetContent(), client.GetRequest().GetContentLen()+length, 1))
	memmove(client.GetRequest().GetContent()+client.GetRequest().GetContentLen(), at, length)
	client.GetRequest().SetContentLen(client.GetRequest().GetContentLen() + length)
	return 0
}
func PhpCliServerClientReadRequestOnMessageComplete(parser *PhpHttpParser) int {
	var client *PhpCliServerClient = parser.GetData()
	client.GetRequest().SetProtocolVersion(parser.GetHttpMajor()*100 + parser.GetHttpMinor())
	PhpCliServerRequestTranslateVpath(client.GetRequest(), client.GetServer().GetDocumentRoot(), client.GetServer().GetDocumentRootLen())
	var vpath *byte = client.GetRequest().GetVpath()
	var end *byte = vpath + client.GetRequest().GetVpathLen()
	var p *byte = end
	client.GetRequest().SetExt(end)
	client.GetRequest().SetExtLen(0)
	for p > vpath {
		p--
		if (*p) == '.' {
			p++
			client.GetRequest().SetExt(p)
			client.GetRequest().SetExtLen(end - p)
			break
		}
	}
	client.SetRequestRead(1)
	return 0
}
func PhpCliServerClientReadRequest(client *PhpCliServerClient, errstr **byte) int {
	var buf []byte
	var settings PhpHttpParserSettings = PhpHttpParserSettings{
		PhpCliServerClientReadRequestOnMessageBegin,
		PhpCliServerClientReadRequestOnPath,
		PhpCliServerClientReadRequestOnQueryString,
		PhpCliServerClientReadRequestOnUrl,
		PhpCliServerClientReadRequestOnFragment,
		PhpCliServerClientReadRequestOnHeaderField,
		PhpCliServerClientReadRequestOnHeaderValue,
		PhpCliServerClientReadRequestOnHeadersComplete,
		PhpCliServerClientReadRequestOnBody,
		PhpCliServerClientReadRequestOnMessageComplete,
	}
	var nbytes_consumed int
	var nbytes_read int
	if client.GetRequestRead() != 0 {
		return 1
	}
	nbytes_read = recv(client.GetSock(), buf, b.SizeOf("buf")-1, 0)
	if nbytes_read < 0 {
		var err int = core.PhpSocketErrno()
		if err == SOCK_EAGAIN {
			return 0
		}
		if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
			*errstr = core.PhpSocketStrerror(err, nil, 0)
		}
		return -1
	} else if nbytes_read == 0 {
		if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
			*errstr = zend.Estrdup(PhpCliServerRequestErrorUnexpectedEof)
		}
		return -1
	}
	client.GetParser().SetData(client)
	nbytes_consumed = PhpHttpParserExecute(client.GetParser(), &settings, buf, nbytes_read)
	if nbytes_consumed != int(nbytes_read) {
		if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
			if (buf[0]&0x80) != 0 || buf[0] == 0x16 {
				*errstr = zend.Estrdup("Unsupported SSL request")
			} else {
				*errstr = zend.Estrdup("Malformed HTTP request")
			}
		}
		return -1
	}
	if client.GetCurrentHeaderName() != nil {
		var header_name *byte = zend.SafePemalloc(client.GetCurrentHeaderNameLen(), 1, 1, 1)
		memmove(header_name, client.GetCurrentHeaderName(), client.GetCurrentHeaderNameLen())
		client.SetCurrentHeaderName(header_name)
		client.SetCurrentHeaderNameAllocated(1)
	}
	if client.GetRequestRead() != 0 {
		return 1
	} else {
		return 0
	}
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
func PhpCliServerClientCtor(client *PhpCliServerClient, server *PhpCliServer, client_sock core.PhpSocketT, addr *__struct__sockaddr, addr_len socklen_t) int {
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
	if zend.FAILURE == PhpCliServerRequestCtor(client.GetRequest()) {
		return zend.FAILURE
	}
	client.SetContentSenderInitialized(0)
	client.SetFileFd(-1)
	return zend.SUCCESS
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
	var content_template *byte = GetTemplateString(status)
	var errstr *byte = GetLastError()
	r.Assert(status_string != nil && content_template != nil)
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
	var buffer zend.SmartStr = zend.SmartStr{0}
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.GetS() == nil {

		/* out of memory */

		goto fail

		/* out of memory */

	}
	AppendEssentialHeaders(&buffer, client, 1)
	zend.SmartStrAppendsEx(&buffer, "Content-Type: text/html; charset=UTF-8\r\n", 1)
	zend.SmartStrAppendsEx(&buffer, "Content-Length: ", 1)
	zend.SmartStrAppendUnsignedEx(&buffer, PhpCliServerBufferSize(client.GetContentSender().GetBuffer()), 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	chunk = PhpCliServerChunkHeapNew(buffer.GetS(), buffer.GetS().GetVal(), buffer.GetS().GetLen())
	if chunk == nil {
		zend.SmartStrFreeEx(&buffer, 1)
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
	zend.ZendStreamInitFilename(&zfd, core.SG(request_info).path_translated)
	var __orig_bailout *JMP_BUF = zend.ExecutorGlobals.GetBailout()
	var __bailout JMP_BUF
	zend.ExecutorGlobals.SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		core.PhpExecuteScript(&zfd)
	}
	zend.ExecutorGlobals.SetBailout(__orig_bailout)
	PhpCliServerLogResponse(client, core.SG(sapi_headers).http_response_code, nil)
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
	var buffer zend.SmartStr = zend.SmartStr{0}
	var mime_type *byte = GetMimeType(server, client.GetRequest().GetExt(), client.GetRequest().GetExtLen())
	AppendHttpStatusLine(&buffer, client.GetRequest().GetProtocolVersion(), status, 1)
	if buffer.GetS() == nil {

		/* out of memory */

		PhpCliServerLogResponse(client, 500, nil)
		return zend.FAILURE
	}
	AppendEssentialHeaders(&buffer, client, 1)
	if mime_type != nil {
		zend.SmartStrAppendlEx(&buffer, "Content-Type: ", b.SizeOf("\"Content-Type: \"")-1, 1)
		zend.SmartStrAppendsEx(&buffer, mime_type, 1)
		if strncmp(mime_type, "text/", 5) == 0 {
			zend.SmartStrAppendsEx(&buffer, "; charset=UTF-8", 1)
		}
		zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	}
	zend.SmartStrAppendsEx(&buffer, "Content-Length: ", 1)
	zend.SmartStrAppendUnsignedEx(&buffer, client.GetRequest().GetSb().st_size, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	zend.SmartStrAppendlEx(&buffer, "\r\n", 2, 1)
	chunk = PhpCliServerChunkHeapNew(buffer.GetS(), buffer.GetS().GetVal(), buffer.GetS().GetLen())
	if chunk == nil {
		zend.SmartStrFreeEx(&buffer, 1)
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
	PhpCliServerClientPopulateRequestInfo(client, &(core.SG(request_info)))
	if nil != b.Assign(&auth, zend.ZendHashStrFindPtr(client.GetRequest().GetHeaders(), "authorization", b.SizeOf("\"authorization\"")-1)) {
		core.PhpHandleAuthData(auth)
	}
	core.SG(sapi_headers).http_response_code = 200
	if zend.FAILURE == core.PhpRequestStartup() {

		/* should never be happen */

		DestroyRequestInfo(&(core.SG(request_info)))
		return zend.FAILURE
	}
	core.PG(during_request_startup) = 0
	return zend.SUCCESS
}
func PhpCliServerRequestShutdown(server *PhpCliServer, client *PhpCliServerClient) int {
	core.PhpRequestShutdown(0)
	PhpCliServerCloseConnection(server, client)
	DestroyRequestInfo(&(core.SG(request_info)))
	core.SG(server_context) = nil
	core.SG(rfc1867_uploaded_files) = nil
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
	var __orig_bailout *JMP_BUF = zend.ExecutorGlobals.GetBailout()
	var __bailout JMP_BUF
	zend.ExecutorGlobals.SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		var retval zend.Zval
		zend.ZVAL_UNDEF(&retval)
		if zend.SUCCESS == zend.ZendExecuteScripts(zend.ZEND_REQUIRE, &retval, 1, &zfd) {
			if retval.GetType() != zend.IS_UNDEF {
				decline = retval.IsType(zend.IS_FALSE)
				zend.ZvalPtrDtor(&retval)
			}
		} else {
			decline = 1
		}
	}
	zend.ExecutorGlobals.SetBailout(__orig_bailout)
	if old_cwd[0] != '0' {
		core.PhpIgnoreValue(zend.VCWD_CHDIR(old_cwd))
	}
	zend.FreeAlloca(old_cwd, use_heap)
	return decline
}
func PhpCliServerDispatch(server *PhpCliServer, client *PhpCliServerClient) int {
	var is_static_file int = 0
	var ext *byte = client.GetRequest().GetExt()
	core.SG(server_context) = client
	if client.GetRequest().GetExtLen() != 3 || ext[0] != 'p' && ext[0] != 'P' || ext[1] != 'h' && ext[1] != 'H' || ext[2] != 'p' && ext[2] != 'P' || client.GetRequest().GetPathTranslated() == nil {
		is_static_file = 1
	}
	if server.GetRouter() != nil || is_static_file == 0 {
		if zend.FAILURE == PhpCliServerRequestStartup(server, client) {
			core.SG(server_context) = nil
			PhpCliServerCloseConnection(server, client)
			DestroyRequestInfo(&(core.SG(request_info)))
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
			if core.SG(sapi_headers).http_response_code == 304 {
				core.SG(sapi_headers).send_default_content_type = 0
			}
			PhpCliServerRequestShutdown(server, client)
			return zend.SUCCESS
		}
	} else {
		if server.GetRouter() != nil {
			var send_header_func func(*core.SapiHeaders) int
			send_header_func = core.sapi_module.GetSendHeaders()

			/* do not generate default content type header */

			core.SG(sapi_headers).send_default_content_type = 0

			/* we don't want headers to be sent */

			core.sapi_module.SetSendHeaders(SapiCliServerDiscardHeaders)
			core.PhpRequestShutdown(0)
			core.sapi_module.SetSendHeaders(send_header_func)
			core.SG(sapi_headers).send_default_content_type = 1
			core.SG(rfc1867_uploaded_files) = nil
		}
		if zend.SUCCESS != PhpCliServerBeginSendStatic(server, client) {
			PhpCliServerCloseConnection(server, client)
		}
		core.SG(server_context) = nil
		return zend.SUCCESS
	}
	core.SG(server_context) = nil
	DestroyRequestInfo(&(core.SG(request_info)))
	return zend.SUCCESS
}
func PhpCliServerMimeTypeCtor(server *PhpCliServer, mime_type_map *PhpCliServerExtMimeTypePair) int {
	var pair *PhpCliServerExtMimeTypePair
	zend.ZendHashInit(server.GetExtensionMimeTypes(), 0, nil, nil, 1)
	for pair = mime_type_map; pair.GetExt() != nil; pair++ {
		var ext_len int = strlen(pair.GetExt())
		zend.ZendHashStrAddPtr(server.GetExtensionMimeTypes(), pair.GetExt(), ext_len, any(pair.GetMimeType()))
	}
	return zend.SUCCESS
}
func PhpCliServerDtor(server *PhpCliServer) {
	zend.ZendHashDestroy(server.GetClients())
	zend.ZendHashDestroy(server.GetExtensionMimeTypes())
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
	PhpCliServerPollerRemove(p.GetServer().GetPoller(), POLLIN|POLLOUT, p.GetSock())
	PhpCliServerClientDtor(p)
	zend.Pefree(p, 1)
}
func PhpCliServerCtor(server *PhpCliServer, addr *byte, document_root *byte, router *byte) int {
	var retval int = zend.SUCCESS
	var host *byte = nil
	var errstr *zend.ZendString = nil
	var _document_root *byte = nil
	var _router *byte = nil
	var err int = 0
	var port int = 3000
	var server_sock core.PhpSocketT = core.SOCK_ERR
	var p *byte = nil
	if addr[0] == '[' {
		host = zend.Pestrdup(addr+1, 1)
		if host == nil {
			return zend.FAILURE
		}
		p = strchr(host, ']')
		if p != nil {
			b.PostInc(&(*p)) = '0'
			if (*p) == ':' {
				port = strtol(p+1, &p, 10)
				if port <= 0 || port > 65535 {
					p = nil
				}
			} else if (*p) != '0' {
				p = nil
			}
		}
	} else {
		host = zend.Pestrdup(addr, 1)
		if host == nil {
			return zend.FAILURE
		}
		p = strchr(host, ':')
		if p != nil {
			b.PostInc(&(*p)) = '0'
			port = strtol(p, &p, 10)
			if port <= 0 || port > 65535 {
				p = nil
			}
		}
	}
	if p == nil {
		r.Fprintf(stderr, "Invalid address: %s\n", addr)
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
	server.SetHost(host)
	server.SetPort(port)
	zend.ZendHashInit(server.GetClients(), 0, nil, PhpCliServerClientDtorWrapper, 1)
	var document_root_len int = strlen(document_root)
	_document_root = zend.Pestrndup(document_root, document_root_len, 1)
	if _document_root == nil {
		retval = zend.FAILURE
		goto out
	}
	server.SetDocumentRoot(_document_root)
	server.SetDocumentRootLen(document_root_len)
	if router != nil {
		var router_len int = strlen(router)
		_router = zend.Pestrndup(router, router_len, 1)
		if _router == nil {
			retval = zend.FAILURE
			goto out
		}
		server.SetRouter(_router)
		server.SetRouterLen(router_len)
	} else {
		server.SetRouter(nil)
		server.SetRouterLen(0)
	}
	if PhpCliServerMimeTypeCtor(server, MimeTypeMap) == zend.FAILURE {
		retval = zend.FAILURE
		goto out
	}
	server.SetIsRunning(1)
out:
	if retval != zend.SUCCESS {
		if host != nil {
			zend.Pefree(host, 1)
		}
		if _document_root != nil {
			zend.Pefree(_document_root, 1)
		}
		if _router != nil {
			zend.Pefree(_router, 1)
		}
		if server_sock > -1 {
			core.Closesocket(server_sock)
		}
	}
	return retval
}
func PhpCliServerRecvEventReadRequest(server *PhpCliServer, client *PhpCliServerClient) int {
	var errstr *byte = nil
	var status int = PhpCliServerClientReadRequest(client, &errstr)
	if status < 0 {
		if errstr != nil {
			if strcmp(errstr, PhpCliServerRequestErrorUnexpectedEof) == 0 && client.GetParser().GetState() == SStartReq {
				PhpCliServerLogf(PHP_CLI_SERVER_LOG_MESSAGE, "%s Closed without sending a request; it was probably just an unused speculative preconnection", client.GetAddrStr())
			} else {
				PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "%s Invalid request (%s)", client.GetAddrStr(), errstr)
			}
			zend.Efree(errstr)
		}
		PhpCliServerCloseConnection(server, client)
		return zend.FAILURE
	} else if status == 1 && client.GetRequest().GetRequestMethod() == PHP_HTTP_NOT_IMPLEMENTED {
		return PhpCliServerSendErrorPage(server, client, 501)
	} else if status == 1 {
		PhpCliServerPollerRemove(server.GetPoller(), POLLIN, client.GetSock())
		PhpCliServerDispatch(server, client)
	} else {
		PhpCliServerPollerAdd(server.GetPoller(), POLLIN, client.GetSock())
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
func PhpCliServerDoEventForEachFdCallback(_params any, fd core.PhpSocketT, event int) int {
	var params *PhpCliServerDoEventForEachFdCallbackParams = _params
	var server *PhpCliServer = params.GetServer()
	if server.GetServerSock() == fd {
		var client *PhpCliServerClient = nil
		var client_sock core.PhpSocketT
		var socklen socklen_t = server.GetSocklen()
		var sa *__struct__sockaddr = zend.Pemalloc(server.GetSocklen(), 1)
		client_sock = accept(server.GetServerSock(), sa, &socklen)
		if !(zend.ZEND_VALID_SOCKET(client_sock)) {
			if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
				var errstr *byte = core.PhpSocketStrerror(core.PhpSocketErrno(), nil, 0)
				PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "Failed to accept a client (reason: %s)", errstr)
				zend.Efree(errstr)
			}
			zend.Pefree(sa, 1)
			return zend.SUCCESS
		}
		if zend.SUCCESS != core.PhpSetSockBlocking(client_sock, 0) {
			zend.Pefree(sa, 1)
			core.Closesocket(client_sock)
			return zend.SUCCESS
		}
		client = zend.Pemalloc(b.SizeOf("php_cli_server_client"), 1)
		if zend.FAILURE == PhpCliServerClientCtor(client, server, client_sock, sa, socklen) {
			PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "Failed to create a new request object")
			zend.Pefree(sa, 1)
			core.Closesocket(client_sock)
			return zend.SUCCESS
		}
		PhpCliServerLogf(PHP_CLI_SERVER_LOG_MESSAGE, "%s Accepted", client.GetAddrStr())
		zend.ZendHashIndexUpdatePtr(server.GetClients(), client_sock, client)
		PhpCliServerPollerAdd(server.GetPoller(), POLLIN, client.GetSock())
	} else {
		var client *PhpCliServerClient
		if nil != b.Assign(&client, zend.ZendHashIndexFindPtr(server.GetClients(), fd)) {
			if (event & POLLIN) != 0 {
				params.GetRhandler()(server, client)
			}
			if (event & POLLOUT) != 0 {
				params.GetWhandler()(server, client)
			}
		}
	}
	return zend.SUCCESS
}
func PhpCliServerDoEventForEachFd(server *PhpCliServer, rhandler func(*PhpCliServer, *PhpCliServerClient) int, whandler func(*PhpCliServer, *PhpCliServerClient) int) {
	var params PhpCliServerDoEventForEachFdCallbackParams = PhpCliServerDoEventForEachFdCallbackParams{server, rhandler, whandler}
	PhpCliServerPollerIterOnActive(server.GetPoller(), &params, PhpCliServerDoEventForEachFdCallback)
}
func PhpCliServerDoEventLoop(server *PhpCliServer) int {
	var retval int = zend.SUCCESS
	for server.GetIsRunning() != 0 {
		var tv __struct__timeval = __struct__timeval{1, 0}
		var n int = PhpCliServerPollerPoll(server.GetPoller(), &tv)
		if n > 0 {
			PhpCliServerDoEventForEachFd(server, PhpCliServerRecvEventReadRequest, PhpCliServerSendEvent)
		} else if n == 0 {

		} else {
			var err int = core.PhpSocketErrno()
			if err != SOCK_EINTR {
				if PhpCliServerLogLevel >= PHP_CLI_SERVER_LOG_ERROR {
					var errstr *byte = core.PhpSocketStrerror(err, nil, 0)
					PhpCliServerLogf(PHP_CLI_SERVER_LOG_ERROR, "%s", errstr)
					zend.Efree(errstr)
				}
				retval = zend.FAILURE
				goto out
			}
		}
	}
out:
	return retval
}
func PhpCliServerSigintHandler(sig int) { Server.SetIsRunning(0) }
func DoCliServer(argc int, argv **byte) int {
	var php_optarg *byte = nil
	var php_optind int = 1
	var c int
	var server_bind_address *byte = nil
	var OPTIONS []core.Opt
	var document_root *byte = nil
	var router *byte = nil
	var document_root_buf []byte
	for b.Assign(&c, core.PhpGetopt(argc, argv, OPTIONS, &php_optarg, &php_optind, 0, 2)) != -1 {
		switch c {
		case 'S':
			server_bind_address = php_optarg
			break
		case 't':
			document_root = php_optarg
			break
		case 'q':
			if PhpCliServerLogLevel > 1 {
				PhpCliServerLogLevel--
			}
			break
		}
	}
	if document_root != nil {
		var sb zend.ZendStatT
		if zend.PhpSysStat(document_root, &sb) {
			r.Fprintf(stderr, "Directory %s does not exist.\n", document_root)
			return 1
		}
		if !(zend.S_ISDIR(sb.st_mode)) {
			r.Fprintf(stderr, "%s is not a directory.\n", document_root)
			return 1
		}
		if zend.VCWD_REALPATH(document_root, document_root_buf) != nil {
			document_root = document_root_buf
		}
	} else {
		var ret *byte = nil
		ret = zend.VCWD_GETCWD(document_root_buf, core.MAXPATHLEN)
		if ret != nil {
			document_root = document_root_buf
		} else {
			document_root = "."
		}
	}
	if argc > php_optind {
		router = argv[php_optind]
	}
	if zend.FAILURE == PhpCliServerCtor(&Server, server_bind_address, document_root, router) {
		return 1
	}
	core.sapi_module.SetPhpinfoAsText(0)
	PhpCliServerLogf(PHP_CLI_SERVER_LOG_PROCESS, "PHP %s Development Server (http://%s) started", core.PHP_VERSION, server_bind_address)
	zend.ZendSignalInit()
	PhpCliServerDoEventLoop(&Server)
	PhpCliServerDtor(&Server)
	return 0
}
