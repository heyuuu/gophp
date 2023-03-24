package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifStreamSocketPair
var DefZifStreamSocketPair = def.DefFunc("stream_socket_pair", 3, 3, []def.ArgInfo{{name: "domain"}, {name: "type_"}, {name: "protocol"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 3, 3, 0)
	domain := fp.ParseZval()
	type_ := fp.ParseZval()
	protocol := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketPair(executeData, returnValue, domain, type_, protocol)
})

// generate by ZifStreamSocketClient
var DefZifStreamSocketClient = def.DefFunc("stream_socket_client", 1, 6, []def.ArgInfo{{name: "remoteaddress"}, {name: "errcode"}, {name: "errstring"}, {name: "timeout"}, {name: "flags"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 6, 0)
	remoteaddress := fp.ParseZval()
	fp.StartOptional()
	errcode := fp.ParseZvalEx(false, true)
	errstring := fp.ParseZvalEx(false, true)
	timeout := fp.ParseZval()
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketClient(executeData, returnValue, remoteaddress, nil, errcode, errstring, timeout, flags, context)
})

// generate by ZifStreamSocketServer
var DefZifStreamSocketServer = def.DefFunc("stream_socket_server", 1, 5, []def.ArgInfo{{name: "localaddress"}, {name: "errcode"}, {name: "errstring"}, {name: "flags"}, {name: "context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 5, 0)
	localaddress := fp.ParseZval()
	fp.StartOptional()
	errcode := fp.ParseZvalEx(false, true)
	errstring := fp.ParseZvalEx(false, true)
	flags := fp.ParseZval()
	context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketServer(executeData, returnValue, localaddress, nil, errcode, errstring, flags, context)
})

// generate by ZifStreamSocketAccept
var DefZifStreamSocketAccept = def.DefFunc("stream_socket_accept", 1, 3, []def.ArgInfo{{name: "serverstream"}, {name: "timeout"}, {name: "peername"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	serverstream := fp.ParseZval()
	fp.StartOptional()
	timeout := fp.ParseZval()
	peername := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifStreamSocketAccept(executeData, returnValue, serverstream, nil, timeout, peername)
})

// generate by ZifStreamSocketGetName
var DefZifStreamSocketGetName = def.DefFunc("stream_socket_get_name", 2, 2, []def.ArgInfo{{name: "stream"}, {name: "want_peer"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	want_peer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketGetName(executeData, returnValue, stream, want_peer)
})

// generate by ZifStreamSocketSendto
var DefZifStreamSocketSendto = def.DefFunc("stream_socket_sendto", 2, 4, []def.ArgInfo{{name: "stream"}, {name: "data"}, {name: "flags"}, {name: "target_addr"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	data := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	target_addr := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketSendto(executeData, returnValue, stream, data, nil, flags, target_addr)
})

// generate by ZifStreamSocketRecvfrom
var DefZifStreamSocketRecvfrom = def.DefFunc("stream_socket_recvfrom", 2, 4, []def.ArgInfo{{name: "stream"}, {name: "amount"}, {name: "flags"}, {name: "remote_addr"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	amount := fp.ParseZval()
	fp.StartOptional()
	flags := fp.ParseZval()
	remote_addr := fp.ParseZvalEx(false, true)
	if fp.HasError() {
		return
	}
	ZifStreamSocketRecvfrom(executeData, returnValue, stream, amount, nil, flags, remote_addr)
})

// generate by ZifStreamGetContents
var DefZifStreamGetContents = def.DefFunc("stream_get_contents", 1, 3, []def.ArgInfo{{name: "source"}, {name: "maxlen"}, {name: "offset"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 3, 0)
	source := fp.ParseZval()
	fp.StartOptional()
	maxlen := fp.ParseZval()
	offset := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetContents(executeData, returnValue, source, nil, maxlen, offset)
})

// generate by ZifStreamCopyToStream
var DefZifStreamCopyToStream = def.DefFunc("stream_copy_to_stream", 2, 4, []def.ArgInfo{{name: "source"}, {name: "dest"}, {name: "maxlen"}, {name: "pos"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	source := fp.ParseZval()
	dest := fp.ParseZval()
	fp.StartOptional()
	maxlen := fp.ParseZval()
	pos := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamCopyToStream(executeData, returnValue, source, dest, nil, maxlen, pos)
})

// generate by ZifStreamGetMetaData
var DefZifStreamGetMetaData = def.DefFunc("stream_get_meta_data", 1, 1, []def.ArgInfo{{name: "fp"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	fp := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetMetaData(executeData, returnValue, fp)
})

// generate by ZifStreamGetTransports
var DefZifStreamGetTransports = def.DefFunc("stream_get_transports", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetTransports(executeData, returnValue)
})

// generate by ZifStreamGetWrappers
var DefZifStreamGetWrappers = def.DefFunc("stream_get_wrappers", 0, 0, []def.ArgInfo{}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetWrappers(executeData, returnValue)
})

// generate by ZifStreamSelect
var DefZifStreamSelect = def.DefFunc("stream_select", 4, 5, []def.ArgInfo{{name: "read_streams"}, {name: "write_streams"}, {name: "except_streams"}, {name: "tv_sec"}, {name: "tv_usec"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 4, 5, 0)
	read_streams := fp.ParseZvalEx(false, true)
	write_streams := fp.ParseZvalEx(false, true)
	except_streams := fp.ParseZvalEx(false, true)
	tv_sec := fp.ParseZval()
	fp.StartOptional()
	tv_usec := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSelect(executeData, returnValue, read_streams, write_streams, except_streams, tv_sec, nil, tv_usec)
})

// generate by ZifStreamContextGetOptions
var DefZifStreamContextGetOptions = def.DefFunc("stream_context_get_options", 1, 1, []def.ArgInfo{{name: "stream_or_context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_or_context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetOptions(executeData, returnValue, stream_or_context)
})

// generate by ZifStreamContextSetOption
var DefZifStreamContextSetOption = def.DefFunc("stream_context_set_option", 2, 4, []def.ArgInfo{{name: "stream_or_context"}, {name: "wrappername"}, {name: "optionname"}, {name: "value"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream_or_context := fp.ParseZval()
	wrappername := fp.ParseZval()
	fp.StartOptional()
	optionname := fp.ParseZval()
	value := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetOption(executeData, returnValue, stream_or_context, wrappername, nil, optionname, value)
})

// generate by ZifStreamContextSetParams
var DefZifStreamContextSetParams = def.DefFunc("stream_context_set_params", 2, 2, []def.ArgInfo{{name: "stream_or_context"}, {name: "options"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream_or_context := fp.ParseZval()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetParams(executeData, returnValue, stream_or_context, options)
})

// generate by ZifStreamContextGetParams
var DefZifStreamContextGetParams = def.DefFunc("stream_context_get_params", 1, 1, []def.ArgInfo{{name: "stream_or_context"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_or_context := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetParams(executeData, returnValue, stream_or_context)
})

// generate by ZifStreamContextGetDefault
var DefZifStreamContextGetDefault = def.DefFunc("stream_context_get_default", 0, 1, []def.ArgInfo{{name: "options"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 1, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextGetDefault(executeData, returnValue, nil, options)
})

// generate by ZifStreamContextSetDefault
var DefZifStreamContextSetDefault = def.DefFunc("stream_context_set_default", 1, 1, []def.ArgInfo{{name: "options"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	options := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextSetDefault(executeData, returnValue, options)
})

// generate by ZifStreamContextCreate
var DefZifStreamContextCreate = def.DefFunc("stream_context_create", 0, 2, []def.ArgInfo{{name: "options"}, {name: "params"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 0, 2, 0)
	fp.StartOptional()
	options := fp.ParseZval()
	params := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamContextCreate(executeData, returnValue, nil, options, params)
})

// generate by ZifStreamFilterPrepend
var DefZifStreamFilterPrepend = def.DefFunc("stream_filter_prepend", 2, 4, []def.ArgInfo{{name: "stream"}, {name: "filtername"}, {name: "read_write"}, {name: "filterparams"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	filtername := fp.ParseZval()
	fp.StartOptional()
	read_write := fp.ParseZval()
	filterparams := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterPrepend(executeData, returnValue, stream, filtername, nil, read_write, filterparams)
})

// generate by ZifStreamFilterAppend
var DefZifStreamFilterAppend = def.DefFunc("stream_filter_append", 2, 4, []def.ArgInfo{{name: "stream"}, {name: "filtername"}, {name: "read_write"}, {name: "filterparams"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	filtername := fp.ParseZval()
	fp.StartOptional()
	read_write := fp.ParseZval()
	filterparams := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterAppend(executeData, returnValue, stream, filtername, nil, read_write, filterparams)
})

// generate by ZifStreamFilterRemove
var DefZifStreamFilterRemove = def.DefFunc("stream_filter_remove", 1, 1, []def.ArgInfo{{name: "stream_filter"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream_filter := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterRemove(executeData, returnValue, stream_filter)
})

// generate by ZifStreamGetLine
var DefZifStreamGetLine = def.DefFunc("stream_get_line", 2, 3, []def.ArgInfo{{name: "stream"}, {name: "maxlen"}, {name: "ending"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	stream := fp.ParseZval()
	maxlen := fp.ParseZval()
	fp.StartOptional()
	ending := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamGetLine(executeData, returnValue, stream, maxlen, nil, ending)
})

// generate by ZifStreamSetBlocking
var DefZifStreamSetBlocking = def.DefFunc("stream_set_blocking", 2, 2, []def.ArgInfo{{name: "socket"}, {name: "mode"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	socket := fp.ParseZval()
	mode := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetBlocking(executeData, returnValue, socket, mode)
})

// generate by ZifStreamSetTimeout
var DefZifStreamSetTimeout = def.DefFunc("stream_set_timeout", 2, 3, []def.ArgInfo{{name: "stream"}, {name: "seconds"}, {name: "microseconds"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 3, 0)
	stream := fp.ParseZval()
	seconds := fp.ParseZval()
	fp.StartOptional()
	microseconds := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetTimeout(executeData, returnValue, stream, seconds, nil, microseconds)
})

// generate by ZifStreamSetWriteBuffer
var DefZifStreamSetWriteBuffer = def.DefFunc("stream_set_write_buffer", 2, 2, []def.ArgInfo{{name: "fp"}, {name: "buffer"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetWriteBuffer(executeData, returnValue, fp, buffer)
})

// generate by ZifStreamSetChunkSize
var DefZifStreamSetChunkSize = def.DefFunc("stream_set_chunk_size", 2, 2, []def.ArgInfo{{name: "fp"}, {name: "chunk_size"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	chunk_size := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetChunkSize(executeData, returnValue, fp, chunk_size)
})

// generate by ZifStreamSetReadBuffer
var DefZifStreamSetReadBuffer = def.DefFunc("stream_set_read_buffer", 2, 2, []def.ArgInfo{{name: "fp"}, {name: "buffer"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	fp := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSetReadBuffer(executeData, returnValue, fp, buffer)
})

// generate by ZifStreamSocketEnableCrypto
var DefZifStreamSocketEnableCrypto = def.DefFunc("stream_socket_enable_crypto", 2, 4, []def.ArgInfo{{name: "stream"}, {name: "enable"}, {name: "cryptokind"}, {name: "sessionstream"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 4, 0)
	stream := fp.ParseZval()
	enable := fp.ParseZval()
	fp.StartOptional()
	cryptokind := fp.ParseZval()
	sessionstream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketEnableCrypto(executeData, returnValue, stream, enable, nil, cryptokind, sessionstream)
})

// generate by ZifStreamResolveIncludePath
var DefZifStreamResolveIncludePath = def.DefFunc("stream_resolve_include_path", 1, 1, []def.ArgInfo{{name: "filename"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	filename := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamResolveIncludePath(executeData, returnValue, filename)
})

// generate by ZifStreamIsLocal
var DefZifStreamIsLocal = def.DefFunc("stream_is_local", 1, 1, []def.ArgInfo{{name: "stream"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamIsLocal(executeData, returnValue, stream)
})

// generate by ZifStreamSupportsLock
var DefZifStreamSupportsLock = def.DefFunc("stream_supports_lock", 1, 1, []def.ArgInfo{{name: "stream"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSupportsLock(executeData, returnValue, stream)
})

// generate by ZifStreamIsatty
var DefZifStreamIsatty = def.DefFunc("stream_isatty", 1, 1, []def.ArgInfo{{name: "stream"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	stream := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamIsatty(executeData, returnValue, stream)
})

// generate by ZifStreamSocketShutdown
var DefZifStreamSocketShutdown = def.DefFunc("stream_socket_shutdown", 2, 2, []def.ArgInfo{{name: "stream"}, {name: "how"}}, func(executeData *ZendExecuteData, returnValue *types.Zval) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	how := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamSocketShutdown(executeData, returnValue, stream, how)
})
