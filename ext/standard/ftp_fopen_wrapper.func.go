package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/sapi/cli"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
)

func GET_FTP_RESULT(stream *core.PhpStream) int {
	return GetFtpResult(stream, tmp_line, b.SizeOf("tmp_line"))
}
func GetFtpResult(stream *core.PhpStream, buffer *byte, buffer_size int) int {
	buffer[0] = '0'
	for core.PhpStreamGets(stream, buffer, buffer_size-1) != nil && !(isdigit(int(buffer[0])) && isdigit(int(buffer[1])) && isdigit(int(buffer[2])) && buffer[3] == ' ') {

	}
	return strtol(buffer, nil, 10)
}
func PhpStreamFtpStreamStat(wrapper *core.PhpStreamWrapper, stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	/* For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */

	return -1

	/* For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */
}
func PhpStreamFtpStreamClose(wrapper *core.PhpStreamWrapper, stream *core.PhpStream) int {
	var controlstream *core.PhpStream = stream.GetWrapperthis()
	var ret int = 0
	if controlstream != nil {
		if strpbrk(stream.GetMode(), "wa+") {
			var tmp_line []byte
			var result int

			/* For write modes close data stream first to signal EOF to server */

			result = GET_FTP_RESULT(controlstream)
			if result != 226 && result != 250 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "FTP server error %d:%s", result, tmp_line)
				ret = r.EOF
			}
		}
		core.PhpStreamWriteString(controlstream, "QUIT\r\n")
		core.PhpStreamClose(controlstream)
		stream.SetWrapperthis(nil)
	}
	return ret
}
func PhpFtpFopenConnect(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
	preuseid **core.PhpStream,
	presource **PhpUrl,
	puse_ssl *int,
	puse_ssl_on_data *int,
) *core.PhpStream {
	var stream *core.PhpStream = nil
	var reuseid *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var use_ssl int
	var use_ssl_on_data int = 0
	var tmp_line []byte
	var transport *byte
	var transport_len int
	resource = PhpUrlParse(path)
	if resource == nil || resource.GetPath() == nil {
		if resource != nil && presource != nil {
			*presource = resource
		}
		return nil
	}
	use_ssl = resource.GetScheme() != nil && resource.GetScheme().GetLen() > 3 && resource.GetScheme().GetVal()[3] == 's'

	/* use port 21 if one wasn't specified */

	if resource.GetPort() == 0 {
		resource.SetPort(21)
	}
	transport_len = int(core.Spprintf(&transport, 0, "tcp://%s:%d", resource.GetHost().GetVal(), resource.GetPort()))
	stream = streams.PhpStreamXportCreate(transport, transport_len, core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, nil, nil, context, nil, nil)
	zend.Efree(transport)
	if stream == nil {
		result = 0
		goto connect_errexit
	}
	streams.PhpStreamContextSet(stream, context)
	streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_CONNECT, nil, 0)

	/* Start talking to ftp server */

	result = GET_FTP_RESULT(stream)
	if result > 299 || result < 200 {
		streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_FAILURE, tmp_line, result)
		goto connect_errexit
	}
	if use_ssl != 0 {

		/* send the AUTH TLS request name */

		core.PhpStreamWriteString(stream, "AUTH TLS\r\n")

		/* get the response */

		result = GET_FTP_RESULT(stream)
		if result != 234 {

			/* AUTH TLS not supported try AUTH SSL */

			core.PhpStreamWriteString(stream, "AUTH SSL\r\n")

			/* get the response */

			result = GET_FTP_RESULT(stream)
			if result != 334 {
				streams.PhpStreamWrapperLogError(wrapper, options, "Server doesn't support FTPS.")
				goto connect_errexit
			} else {

				/* we must reuse the old SSL session id */

				reuseid = stream

				/* we must reuse the old SSL session id */

			}
		}
	}
	if use_ssl != 0 {
		if streams.PhpStreamXportCryptoSetup(stream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(stream, 1) < 0 {
			streams.PhpStreamWrapperLogError(wrapper, options, "Unable to activate SSL mode")
			core.PhpStreamClose(stream)
			stream = nil
			goto connect_errexit
		}

		/* set PBSZ to 0 */

		core.PhpStreamWriteString(stream, "PBSZ 0\r\n")

		/* ignore the response */

		result = GET_FTP_RESULT(stream)

		/* set data connection protection level */

		core.PhpStreamWriteString(stream, "PROT P\r\n")

		/* get the response */

		result = GET_FTP_RESULT(stream)
		use_ssl_on_data = result >= 200 && result <= 299 || reuseid != nil
	}

	// #define PHP_FTP_CNTRL_CHK(val,val_len,err_msg) { unsigned char * s = ( unsigned char * ) val , * e = ( unsigned char * ) s + val_len ; while ( s < e ) { if ( iscntrl ( * s ) ) { php_stream_wrapper_log_error ( wrapper , options , err_msg , val ) ; goto connect_errexit ; } s ++ ; } }

	/* send the user name */

	if resource.GetUser() != nil {
		resource.GetUser().GetLen() = PhpRawUrlDecode(resource.GetUser().GetVal(), resource.GetUser().GetLen())
		var s *uint8 = (*uint8)(resource.GetUser().GetVal())
		var e *uint8 = (*uint8)(s + resource.GetUser().GetLen())
		for s < e {
			if iscntrl(*s) {
				streams.PhpStreamWrapperLogError(wrapper, options, "Invalid login %s", resource.GetUser().GetVal())
				goto connect_errexit
			}
			s++
		}
		core.PhpStreamPrintf(stream, "USER %s\r\n", resource.GetUser().GetVal())
	} else {
		core.PhpStreamWriteString(stream, "USER anonymous\r\n")
	}

	/* get the response */

	result = GET_FTP_RESULT(stream)

	/* if a password is required, send it */

	if result >= 300 && result <= 399 {
		streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_AUTH_REQUIRED, tmp_line, 0)
		if resource.GetPass() != nil {
			resource.GetPass().GetLen() = PhpRawUrlDecode(resource.GetPass().GetVal(), resource.GetPass().GetLen())
			var s *uint8 = (*uint8)(resource.GetPass().GetVal())
			var e *uint8 = (*uint8)(s + resource.GetPass().GetLen())
			for s < e {
				if iscntrl(*s) {
					streams.PhpStreamWrapperLogError(wrapper, options, "Invalid password %s", resource.GetPass().GetVal())
					goto connect_errexit
				}
				s++
			}
			core.PhpStreamPrintf(stream, "PASS %s\r\n", resource.GetPass().GetVal())
		} else {

			/* if the user has configured who they are,
			   send that as the password */

			if FG(from_address) {
				core.PhpStreamPrintf(stream, "PASS %s\r\n", FG(from_address))
			} else {
				core.PhpStreamWriteString(stream, "PASS anonymous\r\n")
			}

			/* if the user has configured who they are,
			   send that as the password */

		}

		/* read the response */

		result = GET_FTP_RESULT(stream)
		if result > 299 || result < 200 {
			streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_AUTH_RESULT, tmp_line, result)
		} else {
			streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_AUTH_RESULT, tmp_line, result)
		}
	}
	if result > 299 || result < 200 {
		goto connect_errexit
	}
	if puse_ssl != nil {
		*puse_ssl = use_ssl
	}
	if puse_ssl_on_data != nil {
		*puse_ssl_on_data = use_ssl_on_data
	}
	if preuseid != nil {
		*preuseid = reuseid
	}
	if presource != nil {
		*presource = resource
	}
	return stream
connect_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return nil
}
func PhpFopenDoPasv(stream *core.PhpStream, ip *byte, ip_size int, phoststart **byte) uint16 {
	var tmp_line []byte
	var result int
	var i int
	var portno uint16
	var tpath *byte
	var ttpath *byte
	var hoststart *byte = nil

	/* We try EPSV first, needed for IPv6 and works on some IPv4 servers */

	core.PhpStreamWriteString(stream, "EPSV\r\n")
	result = GET_FTP_RESULT(stream)

	/* check if we got a 229 response */

	if result != 229 {

		/* EPSV failed, let's try PASV */

		core.PhpStreamWriteString(stream, "PASV\r\n")
		result = GET_FTP_RESULT(stream)

		/* make sure we got a 227 response */

		if result != 227 {
			return 0
		}

		/* parse pasv command (129, 80, 95, 25, 13, 221) */

		tpath = tmp_line

		/* skip over the "227 Some message " part */

		for tpath += 4; (*tpath) && !(isdigit(int(*tpath))); tpath++ {

		}
		if !(*tpath) {
			return 0
		}

		/* skip over the host ip, to get the port */

		hoststart = tpath
		for i = 0; i < 4; i++ {
			for ; isdigit(int(*tpath)); tpath++ {

			}
			if (*tpath) != ',' {
				return 0
			}
			*tpath = '.'
			tpath++
		}
		tpath[-1] = '0'
		memcpy(ip, hoststart, ip_size)
		ip[ip_size-1] = '0'
		hoststart = ip

		/* pull out the MSB of the port */

		portno = uint16(strtoul(tpath, &ttpath, 10) * 256)
		if ttpath == nil {

			/* didn't get correct response from PASV */

			return 0

			/* didn't get correct response from PASV */

		}
		tpath = ttpath
		if (*tpath) != ',' {
			return 0
		}
		tpath++

		/* pull out the LSB of the port */

		portno += uint16(strtoul(tpath, &ttpath, 10))

		/* pull out the LSB of the port */

	} else {

		/* parse epsv command (|||6446|) */

		i = 0
		tpath = tmp_line + 4
		for ; *tpath; tpath++ {
			if (*tpath) == '|' {
				i++
				if i == 3 {
					break
				}
			}
		}
		if i < 3 {
			return 0
		}

		/* pull out the port */

		portno = uint16(strtoul(tpath+1, &ttpath, 10))

		/* pull out the port */

	}
	if ttpath == nil {

		/* didn't get correct response from EPSV/PASV */

		return 0

		/* didn't get correct response from EPSV/PASV */

	}
	if phoststart != nil {
		*phoststart = hoststart
	}
	return portno
}
func PhpStreamUrlWrapFtp(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var stream *core.PhpStream = nil
	var datastream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var tmp_line []byte
	var ip []byte
	var portno uint16
	var hoststart *byte = nil
	var result int = 0
	var use_ssl int
	var use_ssl_on_data int = 0
	var reuseid *core.PhpStream = nil
	var file_size int = 0
	var tmpzval *types.Zval
	var allow_overwrite types.ZendBool = 0
	var read_write int8_t = 0
	var transport *byte
	var transport_len int
	var error_message *types.String = nil
	tmp_line[0] = '0'
	if strpbrk(mode, "r+") {
		read_write = 1
	}
	if strpbrk(mode, "wa+") {
		if read_write {
			streams.PhpStreamWrapperLogError(wrapper, options, "FTP does not support simultaneous read/write connections")
			return nil
		}
		if strchr(mode, 'a') {
			read_write = 3
		} else {
			read_write = 2
		}
	}
	if !read_write {

		/* No mode specified? */

		streams.PhpStreamWrapperLogError(wrapper, options, "Unknown file open mode")
		return nil
	}
	if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "proxy")) != nil {
		if read_write == 1 {

			/* Use http wrapper to proxy ftp request */

			return PhpStreamUrlWrapHttp(wrapper, path, mode, options, opened_path, context)

			/* Use http wrapper to proxy ftp request */

		} else {

			/* ftp proxy is read-only */

			streams.PhpStreamWrapperLogError(wrapper, options, "FTP proxy may only be used in read mode")
			return nil
		}
	}
	stream = PhpFtpFopenConnect(wrapper, path, mode, options, opened_path, context, &reuseid, &resource, &use_ssl, &use_ssl_on_data)
	if stream == nil {
		goto errexit
	}

	/* set the connection to be binary */

	core.PhpStreamWriteString(stream, "TYPE I\r\n")
	result = GET_FTP_RESULT(stream)
	if result > 299 || result < 200 {
		goto errexit
	}

	/* find out the size of the file (verifying it exists) */

	core.PhpStreamPrintf(stream, "SIZE %s\r\n", resource.GetPath().GetVal())

	/* read the response */

	result = GET_FTP_RESULT(stream)
	if read_write == 1 {

		/* Read Mode */

		var sizestr *byte

		/* when reading file, it must exist */

		if result > 299 || result < 200 {
			errno = ENOENT
			goto errexit
		}
		sizestr = strchr(tmp_line, ' ')
		if sizestr != nil {
			sizestr++
			file_size = atoi(sizestr)
			streams.PhpStreamNotifyFileSize(context, file_size, tmp_line, result)
		}
	} else if read_write == 2 {

		/* when writing file (but not appending), it must NOT exist, unless a context option exists which allows it */

		if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "overwrite")) != nil {
			if tmpzval.GetLval() != 0 {
				allow_overwrite = 1
			} else {
				allow_overwrite = 0
			}
		}
		if result <= 299 && result >= 200 {
			if allow_overwrite != 0 {

				/* Context permits overwriting file,
				   so we just delete whatever's there in preparation */

				core.PhpStreamPrintf(stream, "DELE %s\r\n", resource.GetPath().GetVal())
				result = GET_FTP_RESULT(stream)
				if result >= 300 || result <= 199 {
					goto errexit
				}
			} else {
				streams.PhpStreamWrapperLogError(wrapper, options, "Remote file already exists and overwrite context option not specified")
				errno = EEXIST
				goto errexit
			}
		}
	}

	/* set up the passive connection */

	portno = PhpFopenDoPasv(stream, ip, b.SizeOf("ip"), &hoststart)
	if portno == 0 {
		goto errexit
	}

	/* Send RETR/STOR command */

	if read_write == 1 {

		/* set resume position if applicable */

		if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "resume_pos")) != nil && tmpzval.IsType(types.IS_LONG) && tmpzval.GetLval() > 0 {
			core.PhpStreamPrintf(stream, "REST "+zend.ZEND_LONG_FMT+"\r\n", tmpzval.GetLval())
			result = GET_FTP_RESULT(stream)
			if result < 300 || result > 399 {
				streams.PhpStreamWrapperLogError(wrapper, options, "Unable to resume from offset "+zend.ZEND_LONG_FMT, tmpzval.GetLval())
				goto errexit
			}
		}

		/* retrieve file */

		memcpy(tmp_line, "RETR", b.SizeOf("\"RETR\""))

		/* retrieve file */

	} else if read_write == 2 {

		/* Write new file */

		memcpy(tmp_line, "STOR", b.SizeOf("\"STOR\""))

		/* Write new file */

	} else {

		/* Append */

		memcpy(tmp_line, "APPE", b.SizeOf("\"APPE\""))

		/* Append */

	}
	core.PhpStreamPrintf(stream, "%s %s\r\n", tmp_line, b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))

	/* open the data channel */

	if hoststart == nil {
		hoststart = resource.GetHost().GetVal()
	}
	transport_len = int(core.Spprintf(&transport, 0, "tcp://%s:%d", hoststart, portno))
	datastream = streams.PhpStreamXportCreate(transport, transport_len, core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, nil, nil, context, &error_message, nil)
	zend.Efree(transport)
	if datastream == nil {
		tmp_line[0] = '0'
		goto errexit
	}
	result = GET_FTP_RESULT(stream)
	if result != 150 && result != 125 {

		/* Could not retrieve or send the file
		 * this data will only be sent to us after connection on the data port was initiated.
		 */

		core.PhpStreamClose(datastream)
		datastream = nil
		goto errexit
	}
	streams.PhpStreamContextSet(datastream, context)
	streams.PhpStreamNotifyProgressInit(context, 0, file_size)
	if use_ssl_on_data != 0 && (streams.PhpStreamXportCryptoSetup(datastream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(datastream, 1) < 0) {
		streams.PhpStreamWrapperLogError(wrapper, options, "Unable to activate SSL mode")
		core.PhpStreamClose(datastream)
		datastream = nil
		tmp_line[0] = '0'
		goto errexit
	}

	/* remember control stream */

	datastream.SetWrapperthis(stream)
	PhpUrlFree(resource)
	return datastream
errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_FAILURE, tmp_line, result)
		core.PhpStreamClose(stream)
	}
	if tmp_line[0] != '0' {
		streams.PhpStreamWrapperLogError(wrapper, options, "FTP server reports %s", tmp_line)
	}
	if error_message != nil {
		streams.PhpStreamWrapperLogError(wrapper, options, "Failed to set up data channel: %s", error_message.GetVal())
		types.ZendStringRelease(error_message)
	}
	return nil
}
func PhpFtpDirstreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)
	var innerstream *core.PhpStream
	var tmp_len int
	var basename *types.String
	innerstream = (*PhpFtpDirstreamData)(stream.GetAbstract()).GetDatastream()
	if count != b.SizeOf("php_stream_dirent") {
		return -1
	}
	if core.PhpStreamEof(innerstream) != 0 {
		return 0
	}
	if core.PhpStreamGetLine(innerstream, ent.GetDName(), b.SizeOf("ent -> d_name"), &tmp_len) == nil {
		return -1
	}
	basename = PhpBasenameZStr(b.CastStr(ent.GetDName(), tmp_len), "")
	tmp_len = cli.MIN(b.SizeOf("ent -> d_name"), basename.GetLen()-1)
	memcpy(ent.GetDName(), basename.GetVal(), tmp_len)
	ent.GetDName()[tmp_len-1] = '0'
	types.ZendStringReleaseEx(basename, 0)

	/* Trim off trailing whitespace characters */

	for tmp_len > 0 && (ent.GetDName()[tmp_len-1] == '\n' || ent.GetDName()[tmp_len-1] == '\r' || ent.GetDName()[tmp_len-1] == '\t' || ent.GetDName()[tmp_len-1] == ' ') {
		ent.GetDName()[b.PreDec(&tmp_len)] = '0'
	}
	return b.SizeOf("php_stream_dirent")
}
func PhpFtpDirstreamClose(stream *core.PhpStream, close_handle int) int {
	var data *PhpFtpDirstreamData = stream.GetAbstract()

	/* close control connection */

	if data.GetControlstream() != nil {
		core.PhpStreamClose(data.GetControlstream())
		data.SetControlstream(nil)
	}

	/* close data connection */

	core.PhpStreamClose(data.GetDatastream())
	data.SetDatastream(nil)
	zend.Efree(data)
	stream.SetAbstract(nil)
	return 0
}
func PhpStreamFtpOpendir(
	wrapper *core.PhpStreamWrapper,
	path *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var stream *core.PhpStream
	var reuseid *core.PhpStream
	var datastream *core.PhpStream = nil
	var dirsdata *PhpFtpDirstreamData
	var resource *PhpUrl = nil
	var result int = 0
	var use_ssl int
	var use_ssl_on_data int = 0
	var hoststart *byte = nil
	var tmp_line []*byte
	var ip []byte
	var portno uint16
	tmp_line[0] = '0'
	stream = PhpFtpFopenConnect(wrapper, path, mode, options, opened_path, context, &reuseid, &resource, &use_ssl, &use_ssl_on_data)
	if stream == nil {
		goto opendir_errexit
	}

	/* set the connection to be ascii */

	core.PhpStreamWriteString(stream, "TYPE A\r\n")
	result = GET_FTP_RESULT(stream)
	if result > 299 || result < 200 {
		goto opendir_errexit
	}

	// tmp_line isn't relevant after the php_fopen_do_pasv().

	tmp_line[0] = '0'

	/* set up the passive connection */

	portno = PhpFopenDoPasv(stream, ip, b.SizeOf("ip"), &hoststart)
	if portno == 0 {
		goto opendir_errexit
	}

	/* open the data channel */

	if hoststart == nil {
		hoststart = resource.GetHost().GetVal()
	}
	datastream = core.PhpStreamSockOpenHost(hoststart, portno, SOCK_STREAM, 0, 0)
	if datastream == nil {
		goto opendir_errexit
	}
	core.PhpStreamPrintf(stream, "NLST %s\r\n", b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result != 150 && result != 125 {

		/* Could not retrieve or send the file
		 * this data will only be sent to us after connection on the data port was initiated.
		 */

		core.PhpStreamClose(datastream)
		datastream = nil
		goto opendir_errexit
	}
	streams.PhpStreamContextSet(datastream, context)
	if use_ssl_on_data != 0 && (streams.PhpStreamXportCryptoSetup(datastream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(datastream, 1) < 0) {
		streams.PhpStreamWrapperLogError(wrapper, options, "Unable to activate SSL mode")
		core.PhpStreamClose(datastream)
		datastream = nil
		goto opendir_errexit
	}
	PhpUrlFree(resource)
	dirsdata = zend.Emalloc(sizeof * dirsdata)
	dirsdata.SetDatastream(datastream)
	dirsdata.SetControlstream(stream)
	dirsdata.SetDirstream(core.PhpStreamAlloc(&PhpFtpDirstreamOps, dirsdata, 0, mode))
	return dirsdata.GetDirstream()
opendir_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_FAILURE, tmp_line, result)
		core.PhpStreamClose(stream)
	}
	if tmp_line[0] != '0' {
		streams.PhpStreamWrapperLogError(wrapper, options, "FTP server reports %s", tmp_line)
	}
	return nil
}
func PhpStreamFtpUrlStat(wrapper *core.PhpStreamWrapper, url *byte, flags int, ssb *core.PhpStreamStatbuf, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var tmp_line []byte

	/* If ssb is NULL then someone is misbehaving */

	if ssb == nil {
		return -1
	}
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		goto stat_errexit
	}
	ssb.GetSb().st_mode = 0644
	core.PhpStreamPrintf(stream, "CWD %s\r\n", b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {
		ssb.GetSb().st_mode |= S_IFREG
	} else {
		ssb.GetSb().st_mode |= S_IFDIR | S_IXUSR | S_IXGRP | S_IXOTH
	}
	core.PhpStreamWriteString(stream, "TYPE I\r\n")
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {
		goto stat_errexit
	}
	core.PhpStreamPrintf(stream, "SIZE %s\r\n", b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {

		/* Failure either means it doesn't exist
		   or it's a directory and this server
		   fails on listing directory sizes */

		if (ssb.GetSb().st_mode & S_IFDIR) != 0 {
			ssb.GetSb().st_size = 0
		} else {
			goto stat_errexit
		}

		/* Failure either means it doesn't exist
		   or it's a directory and this server
		   fails on listing directory sizes */

	} else {
		ssb.GetSb().st_size = atoi(tmp_line + 4)
	}
	core.PhpStreamPrintf(stream, "MDTM %s\r\n", b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result == 213 {
		var p *byte = tmp_line + 4
		var n int
		var tm __struct__tm
		var tmbuf __struct__tm
		var gmt *__struct__tm
		var stamp int64
		for size_t(p-tmp_line) < b.SizeOf("tmp_line") && !(isdigit(*p)) {
			p++
		}
		if size_t(p-tmp_line) > b.SizeOf("tmp_line") {
			goto mdtm_error
		}
		n = sscanf(p, "%4u%2u%2u%2u%2u%2u", tm.tm_year, tm.tm_mon, tm.tm_mday, tm.tm_hour, tm.tm_min, tm.tm_sec)
		if n != 6 {
			goto mdtm_error
		}
		tm.tm_year -= 1900
		tm.tm_mon--
		tm.tm_isdst = -1

		/* figure out the GMT offset */

		stamp = time(nil)
		gmt = core.PhpGmtimeR(&stamp, &tmbuf)
		if gmt == nil {
			goto mdtm_error
		}
		gmt.tm_isdst = -1

		/* apply the GMT offset */

		tm.tm_sec += long(stamp - mktime(gmt))
		tm.tm_isdst = gmt.tm_isdst
		ssb.GetSb().st_mtime = mktime(&tm)
	} else {

		/* error or unsupported command */

	mdtm_error:
		ssb.GetSb().st_mtime = -1
	}
	ssb.GetSb().st_ino = 0
	ssb.GetSb().st_dev = 0
	ssb.GetSb().st_uid = 0
	ssb.GetSb().st_gid = 0
	ssb.GetSb().st_atime = -1
	ssb.GetSb().st_ctime = -1
	ssb.GetSb().st_nlink = 1
	ssb.GetSb().st_rdev = -1
	ssb.GetSb().st_blksize = 4096
	ssb.GetSb().st_blocks = int((4095 + ssb.GetSb().st_size) / ssb.GetSb().st_blksize)
	core.PhpStreamClose(stream)
	PhpUrlFree(resource)
	return 0
stat_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return -1
}
func PhpStreamFtpUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to connect to %s", url)
		}
		goto unlink_errexit
	}
	if resource.GetPath() == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid path provided in %s", url)
		}
		goto unlink_errexit
	}

	/* Attempt to delete the file */

	core.PhpStreamPrintf(stream, "DELE %s\r\n", b.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Error Deleting file: %s", tmp_line)
		}
		goto unlink_errexit
	}
	PhpUrlFree(resource)
	core.PhpStreamClose(stream)
	return 1
unlink_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return 0
}
func PhpStreamFtpRename(wrapper *core.PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource_from *PhpUrl = nil
	var resource_to *PhpUrl = nil
	var result int
	var tmp_line []byte
	resource_from = PhpUrlParse(url_from)
	resource_to = PhpUrlParse(url_to)

	/* Must be same scheme (ftp/ftp or ftps/ftps), same host, and same port
	   (or a 21/0 0/21 combination which is also "same")
	  Also require paths to/from */

	if resource_from == nil || resource_to == nil || resource_from.GetScheme() == nil || resource_to.GetScheme() == nil || types.ZendStringEquals(resource_from.GetScheme(), resource_to.GetScheme()) == 0 || resource_from.GetHost() == nil || resource_to.GetHost() == nil || types.ZendStringEquals(resource_from.GetHost(), resource_to.GetHost()) == 0 || resource_from.GetPort() != resource_to.GetPort() && resource_from.GetPort()*resource_to.GetPort() != 0 && resource_from.GetPort()+resource_to.GetPort() != 21 || resource_from.GetPath() == nil || resource_to.GetPath() == nil {
		goto rename_errexit
	}
	stream = PhpFtpFopenConnect(wrapper, url_from, "r", 0, nil, context, nil, nil, nil, nil)
	if stream == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to connect to %s", resource_from.GetHost().GetVal())
		}
		goto rename_errexit
	}

	/* Rename FROM */

	core.PhpStreamPrintf(stream, "RNFR %s\r\n", b.CondF1(resource_from.GetPath() != nil, func() []byte { return resource_from.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result < 300 || result > 399 {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Error Renaming file: %s", tmp_line)
		}
		goto rename_errexit
	}

	/* Rename TO */

	core.PhpStreamPrintf(stream, "RNTO %s\r\n", b.CondF1(resource_to.GetPath() != nil, func() []byte { return resource_to.GetPath().GetVal() }, "/"))
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Error Renaming file: %s", tmp_line)
		}
		goto rename_errexit
	}
	PhpUrlFree(resource_from)
	PhpUrlFree(resource_to)
	core.PhpStreamClose(stream)
	return 1
rename_errexit:
	if resource_from != nil {
		PhpUrlFree(resource_from)
	}
	if resource_to != nil {
		PhpUrlFree(resource_to)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return 0
}
func PhpStreamFtpMkdir(wrapper *core.PhpStreamWrapper, url *byte, mode int, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var recursive int = options & core.PHP_STREAM_MKDIR_RECURSIVE
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to connect to %s", url)
		}
		goto mkdir_errexit
	}
	if resource.GetPath() == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid path provided in %s", url)
		}
		goto mkdir_errexit
	}
	if recursive == 0 {
		core.PhpStreamPrintf(stream, "MKD %s\r\n", resource.GetPath().GetVal())
		result = GET_FTP_RESULT(stream)
	} else {

		/* we look for directory separator from the end of string, thus hopefully reducing our work load */

		var p *byte
		var e *byte
		var buf *byte
		buf = zend.Estrndup(resource.GetPath().GetVal(), resource.GetPath().GetLen())
		e = buf + resource.GetPath().GetLen()

		/* find a top level directory we need to create */

		for b.Assign(&p, strrchr(buf, '/')) {
			*p = '0'
			core.PhpStreamPrintf(stream, "CWD %s\r\n", b.Cond(strlen(buf), buf, "/"))
			result = GET_FTP_RESULT(stream)
			if result >= 200 && result <= 299 {
				*p = '/'
				break
			}
		}
		core.PhpStreamPrintf(stream, "MKD %s\r\n", b.Cond(strlen(buf), buf, "/"))
		result = GET_FTP_RESULT(stream)
		if result >= 200 && result <= 299 {
			if p == nil {
				p = buf
			}

			/* create any needed directories if the creation of the 1st directory worked */

			for p != e {
				if (*p) == '0' && (*(p + 1)) != '0' {
					*p = '/'
					core.PhpStreamPrintf(stream, "MKD %s\r\n", buf)
					result = GET_FTP_RESULT(stream)
					if result < 200 || result > 299 {
						if (options & core.REPORT_ERRORS) != 0 {
							core.PhpErrorDocref(nil, faults.E_WARNING, "%s", tmp_line)
						}
						break
					}
				}
				p++
			}

			/* create any needed directories if the creation of the 1st directory worked */

		}
		zend.Efree(buf)
	}
	PhpUrlFree(resource)
	core.PhpStreamClose(stream)
	if result < 200 || result > 299 {

		/* Failure */

		return 0

		/* Failure */

	}
	return 1
mkdir_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return 0
}
func PhpStreamFtpRmdir(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to connect to %s", url)
		}
		goto rmdir_errexit
	}
	if resource.GetPath() == nil {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid path provided in %s", url)
		}
		goto rmdir_errexit
	}
	core.PhpStreamPrintf(stream, "RMD %s\r\n", resource.GetPath().GetVal())
	result = GET_FTP_RESULT(stream)
	if result < 200 || result > 299 {
		if (options & core.REPORT_ERRORS) != 0 {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s", tmp_line)
		}
		goto rmdir_errexit
	}
	PhpUrlFree(resource)
	core.PhpStreamClose(stream)
	return 1
rmdir_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		core.PhpStreamClose(stream)
	}
	return 0
}
