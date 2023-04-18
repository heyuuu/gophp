package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/ext/standard/str"
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpSelect(m core.PhpSocketT, r *fd_set, w *fd_set, e *fd_set, t *__struct__timeval) __auto__ {
	return select_(m, r, w, e, t)
}
func GET_CTX_OPT(stream *core.PhpStream, wrapper string, name string, val *types2.Zval) bool {
	return core.PHP_STREAM_CONTEXT(stream) != nil && nil != b.Assign(&val, streams.PhpStreamContextGetOption(core.PHP_STREAM_CONTEXT(stream), wrapper, name))
}
func ZifStreamSocketPair(executeData zpp.Ex, return_value zpp.Ret, domain *types2.Zval, type_ *types2.Zval, protocol *types2.Zval) {
	var domain zend.ZendLong
	var type_ zend.ZendLong
	var protocol zend.ZendLong
	var s1 *core.PhpStream
	var s2 *core.PhpStream
	var pair []core.PhpSocketT
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			domain = fp.ParseLong()
			type_ = fp.ParseLong()
			protocol = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if 0 != socketpair(int(domain), int(type_), int(protocol), pair) {
		var errbuf []byte
		core.PhpErrorDocref(nil, faults.E_WARNING, "failed to create sockets: [%d]: %s", core.PhpSocketErrno(), core.PhpSocketStrerror(core.PhpSocketErrno(), errbuf, b.SizeOf("errbuf")))
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	s1 = core.PhpStreamSockOpenFromSocket(pair[0], 0)
	s2 = core.PhpStreamSockOpenFromSocket(pair[1], 0)

	/* set the __exposed flag.
	 * php_stream_to_zval() does, add_next_index_resource() does not */

	core.PhpStreamAutoCleanup(s1)
	core.PhpStreamAutoCleanup(s2)
	zend.AddNextIndexResource(return_value, s1.GetRes())
	zend.AddNextIndexResource(return_value, s2.GetRes())
}
func ZifStreamSocketClient(executeData zpp.Ex, return_value zpp.Ret, remoteaddress *types2.Zval, _ zpp.Opt, errcode zpp.RefZval, errstring zpp.RefZval, timeout *types2.Zval, flags *types2.Zval, context *types2.Zval) {
	var host *types2.String
	var zerrno *types2.Zval = nil
	var zerrstr *types2.Zval = nil
	var zcontext *types2.Zval = nil
	var timeout float64 = float64(FG__().default_socket_timeout)
	var conv PhpTimeoutUll
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var flags zend.ZendLong = PHP_STREAM_CLIENT_CONNECT
	var errstr *types2.String = nil
	var context *core.PhpStreamContext = nil
	return_value.SetFalse()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 6, 0)
			host = fp.ParseStr()
			fp.StartOptional()
			zerrno = fp.ParseZval()
			zerrstr = fp.ParseZval()
			timeout = fp.ParseDouble()
			flags = fp.ParseLong()
			zcontext = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if (flags & PHP_STREAM_CLIENT_PERSISTENT) != 0 {
		core.Spprintf(&hashkey, 0, "stream_socket_client__%s", host.GetVal())
	}

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	stream = streams.PhpStreamXportCreate(host.GetVal(), host.GetLen(), core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|b.Cond((flags&PHP_STREAM_CLIENT_CONNECT) != 0, streams.STREAM_XPORT_CONNECT, 0)|b.Cond((flags&PHP_STREAM_CLIENT_ASYNC_CONNECT) != 0, streams.STREAM_XPORT_CONNECT_ASYNC, 0), hashkey, &tv, context, &errstr, &err)
	if stream == nil {

		/* host might contain binary characters */

		var quoted_host *types2.String = types2.NewString(str.PhpAddslashes(host.GetStr()))
		core.PhpErrorDocref(nil, faults.E_WARNING, "unable to connect to %s (%s)", quoted_host.GetVal(), b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
		// types.ZendStringReleaseEx(quoted_host, 0)
	}
	if hashkey != nil {
		zend.Efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if zerrstr != nil && errstr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
		} else if errstr != nil {
			// types.ZendStringReleaseEx(errstr, 0)
		}
		return_value.SetFalse()
		return
	}
	if errstr != nil {
		// types.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifStreamSocketServer(executeData zpp.Ex, return_value zpp.Ret, localaddress *types2.Zval, _ zpp.Opt, errcode zpp.RefZval, errstring zpp.RefZval, flags *types2.Zval, context *types2.Zval) {
	var host *byte
	var host_len int
	var zerrno *types2.Zval = nil
	var zerrstr *types2.Zval = nil
	var zcontext *types2.Zval = nil
	var stream *core.PhpStream = nil
	var err int = 0
	var flags zend.ZendLong = streams.STREAM_XPORT_BIND | streams.STREAM_XPORT_LISTEN
	var errstr *types2.String = nil
	var context *core.PhpStreamContext = nil
	return_value.SetFalse()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 5, 0)
			host, host_len = fp.ParseString()
			fp.StartOptional()
			zerrno = fp.ParseZval()
			zerrstr = fp.ParseZval()
			flags = fp.ParseLong()
			zcontext = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, flags&PHP_FILE_NO_DEFAULT_CONTEXT)
	if context != nil {
		//context.Resource().AddRefcount()
	}
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	stream = streams.PhpStreamXportCreate(host, host_len, core.REPORT_ERRORS, streams.STREAM_XPORT_SERVER|int(flags), nil, nil, context, &errstr, &err)
	if stream == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "unable to connect to %s (%s)", host, b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if zerrstr != nil && errstr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
		} else if errstr != nil {
			// types.ZendStringReleaseEx(errstr, 0)
		}
		return_value.SetFalse()
		return
	}
	if errstr != nil {
		// types.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifStreamSocketAccept(executeData zpp.Ex, return_value zpp.Ret, serverstream *types2.Zval, _ zpp.Opt, timeout *types2.Zval, peername zpp.RefZval) {
	var timeout float64 = float64(FG__().default_socket_timeout)
	var zpeername *types2.Zval = nil
	var peername *types2.String = nil
	var conv PhpTimeoutUll
	var tv __struct__timeval
	var stream *core.PhpStream = nil
	var clistream *core.PhpStream = nil
	var zstream *types2.Zval
	var errstr *types2.String = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			zstream = fp.ParseResource()
			fp.StartOptional()
			timeout = fp.ParseDouble()
			zpeername = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)

	/* prepare the timeout value for use */

	conv = php_timeout_ull(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	if 0 == streams.PhpStreamXportAccept(stream, &clistream, b.Cond(zpeername != nil, &peername, nil), nil, nil, &tv, &errstr) && clistream != nil {
		if peername != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zpeername, peername)
		}
		core.PhpStreamToZval(clistream, return_value)
	} else {
		if peername != nil {
			// types.ZendStringRelease(peername)
		}
		core.PhpErrorDocref(nil, faults.E_WARNING, "accept failed: %s", b.CondF1(errstr != nil, func() []byte { return errstr.GetVal() }, "Unknown error"))
		return_value.SetFalse()
	}
	if errstr != nil {
		// types.ZendStringReleaseEx(errstr, 0)
	}
}
func ZifStreamSocketGetName(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, wantPeer *types2.Zval) {
	var stream *core.PhpStream
	var zstream *types2.Zval
	var want_peer types2.ZendBool
	var name *types2.String = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zstream = fp.ParseResource()
			want_peer = fp.ParseBool()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if 0 != streams.PhpStreamXportGetName(stream, want_peer, &name, nil, nil) || name == nil {
		return_value.SetFalse()
		return
	}
	if name.GetLen() == 0 || name.GetStr()[0] == 0 {
		// types.ZendStringReleaseEx(name, 0)
		return_value.SetFalse()
		return
	}
	return_value.SetString(name)
}
func ZifStreamSocketSendto(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, data *types2.Zval, _ zpp.Opt, flags *types2.Zval, targetAddr *types2.Zval) {
	var stream *core.PhpStream
	var zstream *types2.Zval
	var flags zend.ZendLong = 0
	var data *byte
	var target_addr *byte = nil
	var datalen int
	var target_addr_len int = 0
	var sa core.PhpSockaddrStorage
	var sl socklen_t = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			zstream = fp.ParseResource()
			data, datalen = fp.ParseString()
			fp.StartOptional()
			flags = fp.ParseLong()
			target_addr, target_addr_len = fp.ParseString()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if target_addr_len != 0 {

		/* parse the address */

		if types2.FAILURE == core.PhpNetworkParseNetworkAddressWithPort(target_addr, target_addr_len, (*__struct__sockaddr)(&sa), &sl) {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to parse `%s' into a valid network address", target_addr)
			return_value.SetFalse()
			return
		}

		/* parse the address */

	}
	return_value.SetLong(streams.PhpStreamXportSendto(stream, data, datalen, int(flags), b.Cond(target_addr_len != 0, &sa, nil), sl))
	return
}
func ZifStreamSocketRecvfrom(executeData zpp.Ex, return_value zpp.Ret, stream_ zpp.Resource, amount int, _ zpp.Opt, flags int, remoteAddr zpp.RefZval) (string, bool) {
	var zstream *types2.Zval = stream_
	var to_read zend.ZendLong = amount
	var zremote *types2.Zval = remoteAddr
	var stream *core.PhpStream
	var remote_addr *types2.String = nil
	var read_buf *types2.String
	var recvd int

	core.PhpStreamFromZval(stream, zstream)
	if zremote != nil {
		zend.ZEND_TRY_ASSIGN_REF_NULL(zremote)
	}

	if to_read <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter must be greater than 0")
		return "", false
	}
	read_buf = types2.ZendStringAlloc(to_read, 0)
	recvd = streams.PhpStreamXportRecvfrom(stream, read_buf.GetVal(), to_read, int(flags), nil, nil, b.Cond(zremote != nil, &remote_addr, nil))
	if recvd >= 0 {
		if zremote != nil && remote_addr != nil {
			zend.ZEND_TRY_ASSIGN_REF_STR(zremote, remote_addr)
		}
		return read_buf.Cutoff(recvd).GetStr()
	}
	return "", false
}
func ZifStreamGetContents(executeData zpp.Ex, return_value zpp.Ret, source *types2.Zval, _ zpp.Opt, maxlen *types2.Zval, offset *types2.Zval) {
	var stream *core.PhpStream
	var zsrc *types2.Zval
	var maxlen zend.ZendLong = ssize_t(core.PHP_STREAM_COPY_ALL)
	var desiredpos zend.ZendLong = -1
	var contents *types2.String
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 3, 0)
			zsrc = fp.ParseResource()
			fp.StartOptional()
			maxlen = fp.ParseLong()
			desiredpos = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if maxlen < 0 && maxlen != ssize_t(core.PHP_STREAM_COPY_ALL) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length must be greater than or equal to zero, or -1")
		return_value.SetFalse()
		return
	}
	core.PhpStreamFromZval(stream, zsrc)
	if desiredpos >= 0 {
		var seek_res int = 0
		var position zend.ZendOffT
		position = stream.GetPosition()
		if position >= 0 && desiredpos > position {

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

			seek_res = core.PhpStreamSeek(stream, desiredpos-position, r.SEEK_CUR)

			/* use SEEK_CUR to allow emulation in streams that don't support seeking */

		} else if desiredpos < position {

			/* desired position before position or error on tell */

			seek_res = core.PhpStreamSeek(stream, desiredpos, r.SEEK_SET)

			/* desired position before position or error on tell */

		}
		if seek_res != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", desiredpos)
			return_value.SetFalse()
			return
		}
	}
	if b.Assign(&contents, core.PhpStreamCopyToMem(stream, maxlen, 0)) {
		return_value.SetString(contents)
		return
	} else {
		return_value.SetStringVal("")
		return
	}
}
func ZifStreamCopyToStream(executeData zpp.Ex, return_value zpp.Ret, source *types2.Zval, dest *types2.Zval, _ zpp.Opt, maxlen *types2.Zval, pos *types2.Zval) {
	var src *core.PhpStream
	var dest *core.PhpStream
	var zsrc *types2.Zval
	var zdest *types2.Zval
	var maxlen zend.ZendLong = core.PHP_STREAM_COPY_ALL
	var pos zend.ZendLong = 0
	var len_ int
	var ret int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			zsrc = fp.ParseResource()
			zdest = fp.ParseResource()
			fp.StartOptional()
			maxlen = fp.ParseLong()
			pos = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(src, zsrc)
	core.PhpStreamFromZval(dest, zdest)
	if pos > 0 && core.PhpStreamSeek(src, pos, r.SEEK_SET) < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Failed to seek to position "+zend.ZEND_LONG_FMT+" in the stream", pos)
		return_value.SetFalse()
		return
	}
	ret = core.PhpStreamCopyToStreamEx(src, dest, maxlen, &len_)
	if ret != types2.SUCCESS {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(len_)
	return
}

//@zif -alias socket_get_status
func ZifStreamGetMetaData(executeData zpp.Ex, return_value zpp.Ret, fp *types2.Zval) {
	var zstream *types2.Zval
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zstream = fp.ParseResource()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	zend.ArrayInit(return_value)
	if core.PhpStreamPopulateMetaData(stream, return_value) == 0 {
		zend.AddAssocBool(return_value, "timed_out", 0)
		zend.AddAssocBool(return_value, "blocked", 1)
		zend.AddAssocBool(return_value, "eof", core.PhpStreamEof(stream))
	}
	if !(stream.GetWrapperdata().IsUndef()) {
		//stream.GetWrapperdata().AddRefcount()
		zend.AddAssocZval(return_value, "wrapper_data", stream.GetWrapperdata())
	}
	if stream.GetWrapper() != nil {
		zend.AddAssocString(return_value, "wrapper_type", (*byte)(stream.GetWrapper().GetWops().GetLabel()))
	}
	zend.AddAssocString(return_value, "stream_type", (*byte)(stream.GetOps().GetLabel()))
	zend.AddAssocString(return_value, "mode", stream.GetMode())
	zend.AddAssocLong(return_value, "unread_bytes", stream.GetWritepos()-stream.GetReadpos())
	zend.AddAssocBool(return_value, "seekable", stream.GetOps().GetSeek() != nil && !stream.HasFlags(core.PHP_STREAM_FLAG_NO_SEEK))
	if stream.GetOrigPath() != nil {
		zend.AddAssocString(return_value, "uri", stream.GetOrigPath())
	}
}
func ZifStreamGetTransports(executeData zpp.Ex, return_value zpp.Ret) {
	var stream_xport_hash *types2.Array
	var stream_xport *types2.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&stream_xport_hash, streams.PhpStreamXportGetHash()) {
		zend.ArrayInit(return_value)
		var __ht *types2.Array = stream_xport_hash
		for _, _p := range __ht.ForeachData() {
			var _z *types2.Zval = _p.GetVal()

			stream_xport = _p.GetKey()
			zend.AddNextIndexStr(return_value, stream_xport.Copy())
		}
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifStreamGetWrappers(executeData zpp.Ex, return_value zpp.Ret) {
	var url_stream_wrappers_hash *types2.Array
	var stream_protocol *types2.String
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	if b.Assign(&url_stream_wrappers_hash, core.PhpStreamGetUrlStreamWrappersHash()) {
		zend.ArrayInit(return_value)
		var __ht *types2.Array = url_stream_wrappers_hash
		for _, _p := range __ht.ForeachData() {
			var _z *types2.Zval = _p.GetVal()

			stream_protocol = _p.GetKey()
			if stream_protocol != nil {
				zend.AddNextIndexStr(return_value, stream_protocol.Copy())
			}
		}
	} else {
		return_value.SetFalse()
		return
	}
}
func StreamArrayToFdSet(stream_array *types2.Zval, fds *fd_set, max_fd *core.PhpSocketT) int {
	var elem *types2.Zval
	var stream *core.PhpStream
	var cnt int = 0
	if stream_array.GetType() != types2.IS_ARRAY {
		return 0
	}
	var __ht *types2.Array = stream_array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		elem = _z

		/* Temporary int fd is needed for the STREAM data type on windows, passing this_fd directly to php_stream_cast()
		   would eventually bring a wrong result on x64. php_stream_cast() casts to int internally, and this will leave
		   the higher bits of a SOCKET variable uninitialized on systems with little endian. */

		var this_fd core.PhpSocketT
		elem = types2.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}

		/* get the fd.
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 * */

		if types2.SUCCESS == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT|core.PHP_STREAM_CAST_INTERNAL, any(&this_fd), 1) && this_fd != -1 {
			core.PHP_SAFE_FD_SET(this_fd, fds)
			if this_fd > (*max_fd) {
				*max_fd = this_fd
			}
			cnt++
		}

		/* get the fd.
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 * */

	}
	if cnt != 0 {
		return 1
	} else {
		return 0
	}
}
func StreamArrayFromFdSet(stream_array *types2.Zval, fds *fd_set) int {
	var elem *types2.Zval
	var dest_elem *types2.Zval
	var ht *types2.Array
	var stream *core.PhpStream
	var ret int = 0
	var key *types2.String
	var num_ind zend.ZendUlong
	if stream_array.GetType() != types2.IS_ARRAY {
		return 0
	}
	ht = types2.NewArray(stream_array.Array().Len())
	var __ht *types2.Array = stream_array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		num_ind = _p.GetH()
		key = _p.GetKey()
		elem = _z
		var this_fd core.PhpSocketT
		elem = types2.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}

		/* get the fd
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 */

		if types2.SUCCESS == core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT|core.PHP_STREAM_CAST_INTERNAL, any(&this_fd), 1) && this_fd != core.SOCK_ERR {
			if core.PHP_SAFE_FD_ISSET(this_fd, fds) {
				if key == nil {
					dest_elem = ht.IndexUpdate(num_ind, elem)
				} else {
					dest_elem = ht.KeyUpdate(key.GetStr(), elem)
				}
				zend.ZvalAddRef(dest_elem)
				ret++
				continue
			}
		}

		/* get the fd
		 * NB: Most other code will NOT use the PHP_STREAM_CAST_INTERNAL flag
		 * when casting.  It is only used here so that the buffered data warning
		 * is not displayed.
		 */

	}

	/* destroy old array and add new one */

	// zend.ZvalPtrDtor(stream_array)
	stream_array.SetArray(ht)
	return ret
}
func StreamArrayEmulateReadFdSet(stream_array *types2.Zval) int {
	var elem *types2.Zval
	var dest_elem *types2.Zval
	var ht *types2.Array
	var stream *core.PhpStream
	var ret int = 0
	var num_ind zend.ZendUlong
	var key *types2.String
	if stream_array.GetType() != types2.IS_ARRAY {
		return 0
	}
	ht = types2.NewArray(stream_array.Array().Len())
	var __ht *types2.Array = stream_array.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		num_ind = _p.GetH()
		key = _p.GetKey()
		elem = _z
		elem = types2.ZVAL_DEREF(elem)
		core.PhpStreamFromZvalNoVerify(stream, elem)
		if stream == nil {
			continue
		}
		if stream.GetWritepos()-stream.GetReadpos() > 0 {

			/* allow readable non-descriptor based streams to participate in stream_select.
			 * Non-descriptor streams will only "work" if they have previously buffered the
			 * data.  Not ideal, but better than nothing.
			 * This branch of code also allows blocking streams with buffered data to
			 * operate correctly in stream_select.
			 * */

			if key == nil {
				dest_elem = ht.IndexUpdate(num_ind, elem)
			} else {
				dest_elem = ht.KeyUpdate(key.GetStr(), elem)
			}
			zend.ZvalAddRef(dest_elem)
			ret++
			continue
		}
	}
	if ret > 0 {

		/* destroy old array and add new one */

		// zend.ZvalPtrDtor(stream_array)
		stream_array.SetArray(ht)
	} else {
		ht.DestroyEx()
	}
	return ret
}
func ZifStreamSelect(executeData zpp.Ex, return_value zpp.Ret, readStreams zpp.RefZval, writeStreams zpp.RefZval, exceptStreams zpp.RefZval, tvSec *types2.Zval, _ zpp.Opt, tvUsec *types2.Zval) {
	var r_array *types2.Zval
	var w_array *types2.Zval
	var e_array *types2.Zval
	var tv __struct__timeval
	var tv_p *__struct__timeval = nil
	var rfds fd_set
	var wfds fd_set
	var efds fd_set
	var max_fd core.PhpSocketT = 0
	var retval int
	var sets int = 0
	var sec zend.ZendLong
	var usec zend.ZendLong = 0
	var secnull types2.ZendBool
	var set_count int
	var max_set_count int = 0
	for {
		for {
			fp := zpp.FastParseStart(executeData, 4, 5, 0)
			r_array = fp.ParseArrayEx2(true, true, false)
			w_array = fp.ParseArrayEx2(true, true, false)
			e_array = fp.ParseArrayEx2(true, true, false)
			sec, secnull = fp.ParseLongEx(true, false)
			fp.StartOptional()
			usec = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	FD_ZERO(&rfds)
	FD_ZERO(&wfds)
	FD_ZERO(&efds)
	if r_array != nil {
		set_count = StreamArrayToFdSet(r_array, &rfds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if w_array != nil {
		set_count = StreamArrayToFdSet(w_array, &wfds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if e_array != nil {
		set_count = StreamArrayToFdSet(e_array, &efds, &max_fd)
		if set_count > max_set_count {
			max_set_count = set_count
		}
		sets += set_count
	}
	if sets == 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "No stream arrays were passed")
		return_value.SetFalse()
		return
	}
	core.PHP_SAFE_MAX_FD(max_fd, max_set_count)

	/* If seconds is not set to null, build the timeval, else we wait indefinitely */

	if secnull == 0 {
		if sec < 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "The seconds parameter must be greater than 0")
			return_value.SetFalse()
			return
		} else if usec < 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "The microseconds parameter must be greater than 0")
			return_value.SetFalse()
			return
		}

		/* Windows, Solaris and BSD do not like microsecond values which are >= 1 sec */

		tv.tv_sec = long(sec + usec/1000000)
		tv.tv_usec = long(usec % 1000000)
		tv_p = &tv
	}

	/* slight hack to support buffered data; if there is data sitting in the
	 * read buffer of any of the streams in the read array, let's pretend
	 * that we selected, but return only the readable sockets */

	if r_array != nil {
		retval = StreamArrayEmulateReadFdSet(r_array)
		if retval > 0 {
			if w_array != nil {
				// zend.ZvalPtrDtor(w_array)
				w_array.SetEmptyArray()
			}
			if e_array != nil {
				// zend.ZvalPtrDtor(e_array)
				e_array.SetEmptyArray()
			}
			return_value.SetLong(retval)
			return
		}
	}
	retval = PhpSelect(max_fd+1, &rfds, &wfds, &efds, tv_p)
	if retval == -1 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "unable to select [%d]: %s (max_fd=%d)", errno, strerror(errno), max_fd)
		return_value.SetFalse()
		return
	}
	if r_array != nil {
		StreamArrayFromFdSet(r_array, &rfds)
	}
	if w_array != nil {
		StreamArrayFromFdSet(w_array, &wfds)
	}
	if e_array != nil {
		StreamArrayFromFdSet(e_array, &efds)
	}
	return_value.SetLong(retval)
	return
}
func UserSpaceStreamNotifier(
	context *core.PhpStreamContext,
	notifycode int,
	severity int,
	xmsg *byte,
	xcode int,
	bytes_sofar int,
	bytes_max int,
	ptr any,
) {
	var callback *types2.Zval = context.GetNotifier().GetPtr()
	var retval types2.Zval
	var zvs []types2.Zval
	var i int
	zvs[0].SetLong(notifycode)
	zvs[1].SetLong(severity)
	if xmsg != nil {
		zvs[2].SetStringVal(b.CastStrAuto(xmsg))
	} else {
		zvs[2].SetNull()
	}
	zvs[3].SetLong(xcode)
	zvs[4].SetLong(bytes_sofar)
	zvs[5].SetLong(bytes_max)
	if types2.FAILURE == zend.CallUserFunctionEx(nil, callback, &retval, 6, zvs, 0) {
		core.PhpErrorDocref(nil, faults.E_WARNING, "failed to call user notifier")
	}
	for i = 0; i < 6; i++ {
		// zend.ZvalPtrDtor(&zvs[i])
	}
	// zend.ZvalPtrDtor(&retval)
}
func UserSpaceStreamNotifierDtor(notifier *streams.PhpStreamNotifier) {
	if notifier != nil && notifier.GetPtr().IsNotUndef() {
		// zend.ZvalPtrDtor(notifier.GetPtr())
		notifier.GetPtr().SetUndef()
	}
}
func ParseContextOptions(context *core.PhpStreamContext, options *types2.Zval) int {
	var wval *types2.Zval
	var oval *types2.Zval
	var wkey *types2.String
	var okey *types2.String
	var ret int = types2.SUCCESS
	var __ht *types2.Array = options.Array()
	for _, _p := range __ht.ForeachData() {
		var _z *types2.Zval = _p.GetVal()

		wkey = _p.GetKey()
		wval = _z
		wval = types2.ZVAL_DEREF(wval)
		if wkey != nil && wval.IsType(types2.IS_ARRAY) {
			var __ht *types2.Array = wval.Array()
			for _, _p := range __ht.ForeachData() {
				var _z *types2.Zval = _p.GetVal()

				okey = _p.GetKey()
				oval = _z
				if okey != nil {
					streams.PhpStreamContextSetOption(context, wkey.GetVal(), okey.GetVal(), oval)
				}
			}
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "options should have the form [\"wrappername\"][\"optionname\"] = $value")
		}
	}
	return ret
}
func ParseContextParams(context *core.PhpStreamContext, params *types2.Zval) int {
	var ret int = types2.SUCCESS
	var tmp *types2.Zval
	if nil != b.Assign(&tmp, params.Array().KeyFind("notification")) {
		if context.GetNotifier() != nil {
			streams.PhpStreamNotificationFree(context.GetNotifier())
			context.SetNotifier(nil)
		}
		context.SetNotifier(streams.PhpStreamNotificationAlloc())
		context.GetNotifier().SetFunc(UserSpaceStreamNotifier)
		types2.ZVAL_COPY(context.GetNotifier().GetPtr(), tmp)
		context.GetNotifier().SetDtor(UserSpaceStreamNotifierDtor)
	}
	if nil != b.Assign(&tmp, params.Array().KeyFind("options")) {
		if tmp.IsType(types2.IS_ARRAY) {
			ParseContextOptions(context, tmp)
		} else {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
		}
	}
	return ret
}
func DecodeContextParam(contextresource *types2.Zval) *core.PhpStreamContext {
	var context *core.PhpStreamContext = nil
	context = zend.ZendFetchResourceEx(contextresource, nil, PhpLeStreamContext())
	if context == nil {
		var stream *core.PhpStream
		stream = zend.ZendFetchResource2Ex(contextresource, nil, streams.PhpFileLeStream(), streams.PhpFileLePstream())
		if stream != nil {
			context = core.PHP_STREAM_CONTEXT(stream)
			if context == nil {

				/* Only way this happens is if file is opened with NO_DEFAULT_CONTEXT
				   param, but then something is called which requires a context.
				   Don't give them the default one though since they already said they
				    didn't want it. */

				context = streams.PhpStreamContextAlloc()
				stream.SetCtx(context.GetRes())
			}
		}
	}
	return context
}
func ZifStreamContextGetOptions(executeData zpp.Ex, return_value zpp.Ret, streamOrContext *types2.Zval) {
	var zcontext *types2.Zval
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zcontext = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
		return_value.SetFalse()
		return
	}
	types2.ZVAL_COPY(return_value, context.GetOptions())
}
func ZifStreamContextSetOption(executeData zpp.Ex, return_value zpp.Ret, streamOrContext *types2.Zval, wrappername *types2.Zval, _ zpp.Opt, optionname *types2.Zval, value *types2.Zval) {
	var zcontext *types2.Zval = nil
	var context *core.PhpStreamContext
	if executeData.NumArgs() == 2 {
		var options *types2.Zval
		for {
			var _flags int = 0
			var _min_num_args int = 2
			var _max_num_args int = 2
			for {
				fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
				zcontext = fp.ParseResource()
				options = fp.ParseArray()
				if fp.HasError() {
					return
				}
				break
			}
			break
		}

		/* figure out where the context is coming from exactly */

		if !(b.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
			return_value.SetFalse()
			return
		}
		return_value.SetBool(ParseContextOptions(context, options) == types2.SUCCESS)
		return
	} else {
		var zvalue *types2.Zval
		var wrappername *byte
		var optionname *byte
		var wrapperlen int
		var optionlen int
		for {
			var _flags int = 0
			var _min_num_args int = 4
			var _max_num_args int = 4
			for {
				fp := zpp.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
				zcontext = fp.ParseResource()
				wrappername, wrapperlen = fp.ParseString()
				optionname, optionlen = fp.ParseString()
				zvalue = fp.ParseZval()
				if fp.HasError() {
					return_value.SetFalse()
					return
				}
				break
			}
			break
		}

		/* figure out where the context is coming from exactly */

		if !(b.Assign(&context, DecodeContextParam(zcontext))) {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
			return_value.SetFalse()
			return
		}
		return_value.SetBool(streams.PhpStreamContextSetOption(context, wrappername, optionname, zvalue) == types2.SUCCESS)
		return
	}
}
func ZifStreamContextSetParams(executeData zpp.Ex, return_value zpp.Ret, streamOrContext *types2.Zval, options *types2.Zval) {
	var params *types2.Zval
	var zcontext *types2.Zval
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zcontext = fp.ParseResource()
			params = fp.ParseArray()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
		return_value.SetFalse()
		return
	}
	return_value.SetBool(ParseContextParams(context, params) == types2.SUCCESS)
}
func ZifStreamContextGetParams(executeData zpp.Ex, return_value zpp.Ret, streamOrContext *types2.Zval) {
	var zcontext *types2.Zval
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zcontext = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = DecodeContextParam(zcontext)
	if context == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid stream/context parameter")
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	if context.GetNotifier() != nil && context.GetNotifier().GetPtr().IsNotUndef() && context.GetNotifier().GetFunc() == UserSpaceStreamNotifier {
		context.GetNotifier().GetPtr().TryAddRefcount()
		zend.AddAssocZvalEx(return_value, "notification", context.GetNotifier().GetPtr())
	}
	context.GetOptions().TryAddRefcount()
	zend.AddAssocZvalEx(return_value, "options", context.GetOptions())
}
func ZifStreamContextGetDefault(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, options *types2.Zval) {
	var params *types2.Zval = nil
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			params = fp.ParseArray()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if FG__().default_context == nil {
		FG__().default_context = streams.PhpStreamContextAlloc()
	}
	context = FG__().default_context
	if params != nil {
		ParseContextOptions(context, params)
	}
	streams.PhpStreamContextToZval(context, return_value)
}
func ZifStreamContextSetDefault(executeData zpp.Ex, return_value zpp.Ret, options *types2.Zval) {
	var options *types2.Zval = nil
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			options = fp.ParseArray()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if FG__().default_context == nil {
		FG__().default_context = streams.PhpStreamContextAlloc()
	}
	context = FG__().default_context
	ParseContextOptions(context, options)
	streams.PhpStreamContextToZval(context, return_value)
}
func ZifStreamContextCreate(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, options *types2.Zval, params *types2.Zval) {
	var options *types2.Zval = nil
	var params *types2.Zval = nil
	var context *core.PhpStreamContext
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			options = fp.ParseArrayEx(true, false)
			params = fp.ParseArrayEx(true, false)
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextAlloc()
	if options != nil {
		ParseContextOptions(context, options)
	}
	if params != nil {
		ParseContextParams(context, params)
	}
	return_value.SetResource(context.GetRes())
	return
}
func ApplyFilterToStream(append int, executeData *zend.ZendExecuteData, return_value *types2.Zval) {
	var zstream *types2.Zval
	var stream *core.PhpStream
	var filtername *byte
	var filternamelen int
	var read_write zend.ZendLong = 0
	var filterparams *types2.Zval = nil
	var filter *core.PhpStreamFilter = nil
	var ret int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			zstream = fp.ParseResource()
			filtername, filternamelen = fp.ParseString()
			fp.StartOptional()
			read_write = fp.ParseLong()
			filterparams = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if (read_write & streams.PHP_STREAM_FILTER_ALL) == 0 {

		/* Chain not specified.
		 * Examine stream->mode to determine which filters are needed
		 * There's no harm in attaching a filter to an unused chain,
		 * but why waste the memory and clock cycles?
		 */

		if strchr(stream.GetMode(), 'r') || strchr(stream.GetMode(), '+') {
			read_write |= streams.PHP_STREAM_FILTER_READ
		}
		if strchr(stream.GetMode(), 'w') || strchr(stream.GetMode(), '+') || strchr(stream.GetMode(), 'a') {
			read_write |= streams.PHP_STREAM_FILTER_WRITE
		}
	}
	if (read_write & streams.PHP_STREAM_FILTER_READ) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.GetIsPersistent())
		if filter == nil {
			return_value.SetFalse()
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(stream.GetReadfilters(), filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(stream.GetReadfilters(), filter)
		}
		if ret != types2.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			return_value.SetFalse()
			return
		}
	}
	if (read_write & streams.PHP_STREAM_FILTER_WRITE) != 0 {
		filter = streams.PhpStreamFilterCreate(filtername, filterparams, stream.GetIsPersistent())
		if filter == nil {
			return_value.SetFalse()
			return
		}
		if append != 0 {
			ret = streams.PhpStreamFilterAppendEx(stream.GetWritefilters(), filter)
		} else {
			ret = streams.PhpStreamFilterPrependEx(stream.GetWritefilters(), filter)
		}
		if ret != types2.SUCCESS {
			streams.PhpStreamFilterRemove(filter, 1)
			return_value.SetFalse()
			return
		}
	}
	if filter != nil {
		filter.SetRes(zend.ZendRegisterResource(filter, streams.PhpFileLeStreamFilter()))
		filter.GetRes().AddRefcount()
		return_value.SetResource(filter.GetRes())
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifStreamFilterPrepend(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, filtername *types2.Zval, _ zpp.Opt, readWrite *types2.Zval, filterparams *types2.Zval) {
	ApplyFilterToStream(0, executeData, return_value)
}
func ZifStreamFilterAppend(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, filtername *types2.Zval, _ zpp.Opt, readWrite *types2.Zval, filterparams *types2.Zval) {
	ApplyFilterToStream(1, executeData, return_value)
}
func ZifStreamFilterRemove(executeData zpp.Ex, return_value zpp.Ret, streamFilter *types2.Zval) {
	var zfilter *types2.Zval
	var filter *core.PhpStreamFilter
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zfilter = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	filter = zend.ZendFetchResource(zfilter.Resource(), nil, streams.PhpFileLeStreamFilter())
	if filter == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid resource given, not a stream filter")
		return_value.SetFalse()
		return
	}
	if streams.PhpStreamFilterFlush(filter, 1) == types2.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to flush filter, not removing")
		return_value.SetFalse()
		return
	}
	if zend.ZendListClose(zfilter.Resource()) == types2.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Could not invalidate filter, not removing")
		return_value.SetFalse()
		return
	} else {
		streams.PhpStreamFilterRemove(filter, 1)
		return_value.SetTrue()
		return
	}
}
func ZifStreamGetLine(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, maxlen *types2.Zval, _ zpp.Opt, ending *types2.Zval) {
	var str *byte = nil
	var str_len int = 0
	var max_length zend.ZendLong
	var zstream *types2.Zval
	var buf *types2.String
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			zstream = fp.ParseResource()
			max_length = fp.ParseLong()
			fp.StartOptional()
			str, str_len = fp.ParseString()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if max_length < 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The maximum allowed length must be greater than or equal to zero")
		return_value.SetFalse()
		return
	}
	if max_length == 0 {
		max_length = core.PHP_SOCK_CHUNK_SIZE
	}
	core.PhpStreamFromZval(stream, zstream)
	if b.Assign(&buf, streams.PhpStreamGetRecord(stream, max_length, str, str_len)) {
		return_value.SetString(buf)
		return
	} else {
		return_value.SetFalse()
		return
	}
}

//@zif -alias socket_set_blocking
func ZifStreamSetBlocking(executeData zpp.Ex, return_value zpp.Ret, socket *types2.Zval, mode *types2.Zval) {
	var zstream *types2.Zval
	var block types2.ZendBool
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zstream = fp.ParseResource()
			block = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_BLOCKING, block, nil) == -1 {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}

//@zif -alias socket_set_timeout
func ZifStreamSetTimeout(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, seconds *types2.Zval, _ zpp.Opt, microseconds *types2.Zval) {
	var socket *types2.Zval
	var seconds zend.ZendLong
	var microseconds zend.ZendLong = 0
	var t __struct__timeval
	var stream *core.PhpStream
	var argc int = executeData.NumArgs()
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 3, 0)
			socket = fp.ParseResource()
			seconds = fp.ParseLong()
			fp.StartOptional()
			microseconds = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, socket)
	t.tv_sec = seconds
	if argc == 3 {
		t.tv_usec = microseconds % 1000000
		t.tv_sec += microseconds / 1000000
	} else {
		t.tv_usec = 0
	}
	if core.PHP_STREAM_OPTION_RETURN_OK == core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_TIMEOUT, 0, &t) {
		return_value.SetTrue()
		return
	}
	return_value.SetFalse()
	return
}

//@zif -alias set_file_buffer
func ZifStreamSetWriteBuffer(executeData zpp.Ex, return_value zpp.Ret, fp *types2.Zval, buffer *types2.Zval) {
	var arg1 *types2.Zval
	var ret int
	var arg2 zend.ZendLong
	var buff int
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			arg1 = fp.ParseResource()
			arg2 = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_WRITE_BUFFER, core.PHP_STREAM_BUFFER_NONE, nil)
	} else {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_WRITE_BUFFER, core.PHP_STREAM_BUFFER_FULL, &buff)
	}
	return_value.SetLong(b.Cond(ret == 0, 0, r.EOF))
	return
}
func ZifStreamSetChunkSize(executeData zpp.Ex, return_value zpp.Ret, fp *types2.Zval, chunkSize *types2.Zval) {
	var ret int
	var csize zend.ZendLong
	var zstream *types2.Zval
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zstream = fp.ParseResource()
			csize = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if csize <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The chunk size must be a positive integer, given "+zend.ZEND_LONG_FMT, csize)
		return_value.SetFalse()
		return
	}

	/* stream.chunk_size is actually a size_t, but php_stream_set_option
	 * can only use an int to accept the new value and return the old one.
	 * In any case, values larger than INT_MAX for a chunk size make no sense.
	 */

	if csize > core.INT_MAX {
		core.PhpErrorDocref(nil, faults.E_WARNING, "The chunk size cannot be larger than %d", core.INT_MAX)
		return_value.SetFalse()
		return
	}
	core.PhpStreamFromZval(stream, zstream)
	ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_SET_CHUNK_SIZE, int(csize), nil)
	return_value.SetLong(b.CondF(ret > 0, func() zend.ZendLong { return zend.ZendLong(ret) }, func() zend.ZendLong { return zend.ZendLong(r.EOF) }))
	return
}
func ZifStreamSetReadBuffer(executeData zpp.Ex, return_value zpp.Ret, fp *types2.Zval, buffer *types2.Zval) {
	var arg1 *types2.Zval
	var ret int
	var arg2 zend.ZendLong
	var buff int
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			arg1 = fp.ParseResource()
			arg2 = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, arg1)
	buff = arg2

	/* if buff is 0 then set to non-buffered */

	if buff == 0 {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_BUFFER, core.PHP_STREAM_BUFFER_NONE, nil)
	} else {
		ret = core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_BUFFER, core.PHP_STREAM_BUFFER_FULL, &buff)
	}
	return_value.SetLong(b.Cond(ret == 0, 0, r.EOF))
	return
}
func ZifStreamSocketEnableCrypto(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, enable *types2.Zval, _ zpp.Opt, cryptokind *types2.Zval, sessionstream *types2.Zval) {
	var cryptokind zend.ZendLong = 0
	var zstream *types2.Zval
	var zsessstream *types2.Zval = nil
	var stream *core.PhpStream
	var sessstream *core.PhpStream = nil
	var enable types2.ZendBool
	var cryptokindnull types2.ZendBool = 1
	var ret int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 4, 0)
			zstream = fp.ParseResource()
			enable = fp.ParseBool()
			fp.StartOptional()
			cryptokind, cryptokindnull = fp.ParseLongEx(true, false)
			zsessstream = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zstream)
	if enable != 0 {
		if cryptokindnull != 0 {
			var val *types2.Zval
			if !(GET_CTX_OPT(stream, "ssl", "crypto_method", val)) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "When enabling encryption you must specify the crypto type")
				return_value.SetFalse()
				return
			}
			cryptokind = val.Long()
		}
		if zsessstream != nil {
			core.PhpStreamFromZval(sessstream, zsessstream)
		}
		if streams.PhpStreamXportCryptoSetup(stream, cryptokind, sessstream) < 0 {
			return_value.SetFalse()
			return
		}
	}
	ret = streams.PhpStreamXportCryptoEnable(stream, enable)
	switch ret {
	case -1:
		return_value.SetFalse()
		return
	case 0:
		return_value.SetLong(0)
		return
	default:
		return_value.SetTrue()
		return
	}
}
func ZifStreamResolveIncludePath(executeData zpp.Ex, return_value zpp.Ret, filename *types2.Zval) {
	var filename *byte
	var filename_len int
	var resolved_path *string
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			filename, filename_len = fp.ParsePath()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	resolved_path = core.PhpResolvePathForZend(b.CastStr(filename, filename_len))
	if resolved_path != nil {
		return_value.SetStringVal(*resolved_path)
		return
	}
	return_value.SetFalse()
	return
}
func ZifStreamIsLocal(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval) {
	var zstream *types2.Zval
	var stream *core.PhpStream = nil
	var wrapper *core.PhpStreamWrapper = nil
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zstream = fp.ParseZval()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if zstream.IsType(types2.IS_RESOURCE) {
		core.PhpStreamFromZval(stream, zstream)
		if stream == nil {
			return_value.SetFalse()
			return
		}
		wrapper = stream.GetWrapper()
	} else {
		if zend.TryConvertToString(zstream) == 0 {
			return
		}
		wrapper = streams.PhpStreamLocateUrlWrapper(zstream.String().GetVal(), nil, 0)
	}
	if wrapper == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetBool(wrapper.GetIsUrl() == 0)
	return
}
func ZifStreamSupportsLock(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval) {
	var stream *core.PhpStream
	var zsrc *types2.Zval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zsrc = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zsrc)
	if core.PhpStreamSupportsLock(stream) == 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifStreamIsatty(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval) {
	var zsrc *types2.Zval
	var stream *core.PhpStream
	var fileno core.PhpSocketT
	for {
		for {
			fp := zpp.FastParseStart(executeData, 1, 1, 0)
			zsrc = fp.ParseResource()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	core.PhpStreamFromZval(stream, zsrc)
	if core.PhpStreamCanCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT) == types2.SUCCESS {
		core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD_FOR_SELECT, any(&fileno), 0)
	} else if core.PhpStreamCanCast(stream, core.PHP_STREAM_AS_FD) == types2.SUCCESS {
		core.PhpStreamCast(stream, core.PHP_STREAM_AS_FD, any(&fileno), 0)
	} else {
		return_value.SetFalse()
		return
	}

	/* Check if the file descriptor identifier is a terminal */

	return_value.SetBool(zend.Isatty(fileno) != 0)

	/* Check if the file descriptor identifier is a terminal */
}
func ZifStreamSocketShutdown(executeData zpp.Ex, return_value zpp.Ret, stream *types2.Zval, how *types2.Zval) {
	var how zend.ZendLong
	var zstream *types2.Zval
	var stream *core.PhpStream
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			zstream = fp.ParseResource()
			how = fp.ParseLong()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if how != streams.STREAM_SHUT_RD && how != streams.STREAM_SHUT_WR && how != streams.STREAM_SHUT_RDWR {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Second parameter $how needs to be one of STREAM_SHUT_RD, STREAM_SHUT_WR or STREAM_SHUT_RDWR")
		return_value.SetFalse()
		return
	}
	core.PhpStreamFromZval(stream, zstream)
	return_value.SetBool(streams.PhpStreamXportShutdown(stream, streams.StreamShutdownT(how)) == 0)
	return
}
