// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/ftp_fopen_wrapper.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Jim Winstead <jimw@php.net>                                 |
   |          Hartmut Holzgraefe <hholzgra@php.net>                       |
   |          Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "php_network.h"

// # include "php_ini.h"

// # include < stdio . h >

// # include < stdlib . h >

// # include < errno . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < sys / param . h >

// # include "php_standard.h"

// # include < sys / types . h >

// # include < sys / socket . h >

// # include < netinet / in . h >

// # include < netdb . h >

// # include < arpa / inet . h >

// # include "php_fopen_wrappers.h"

// #define FTPS_ENCRYPT_DATA       1

// #define GET_FTP_RESULT(stream) get_ftp_result ( ( stream ) , tmp_line , sizeof ( tmp_line ) )

// @type PhpFtpDirstreamData struct

/* {{{ get_ftp_result
 */

func GetFtpResult(stream *core.PhpStream, buffer *byte, buffer_size int) int {
	buffer[0] = '0'
	for streams._phpStreamGetLine(stream, buffer, buffer_size-1, nil) != nil && !(isdigit(int(buffer[0])) && isdigit(int(buffer[1])) && isdigit(int(buffer[2])) && buffer[3] == ' ') {

	}
	return strtol(buffer, nil, 10)
}

/* }}} */

func PhpStreamFtpStreamStat(wrapper *core.PhpStreamWrapper, stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	/* For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */

	return -1

	/* For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */
}

/* }}} */

func PhpStreamFtpStreamClose(wrapper *core.PhpStreamWrapper, stream *core.PhpStream) int {
	var controlstream *core.PhpStream = stream.wrapperthis
	var ret int = 0
	if controlstream != nil {
		if strpbrk(stream.mode, "wa+") {
			var tmp_line []byte
			var result int

			/* For write modes close data stream first to signal EOF to server */

			result = GetFtpResult(controlstream, tmp_line, g.SizeOf("tmp_line"))
			if result != 226 && result != 250 {
				core.PhpErrorDocref(nil, 1<<1, "FTP server error %d:%s", result, tmp_line)
				ret = -1
			}
		}
		streams._phpStreamWrite(controlstream, "QUIT\r\n", strlen("QUIT\r\n"))
		streams._phpStreamFree(controlstream, 1|2)
		stream.wrapperthis = nil
	}
	return ret
}

/* }}} */

func PhpFtpFopenConnect(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext, preuseid **core.PhpStream, presource **PhpUrl, puse_ssl *int, puse_ssl_on_data *int) *core.PhpStream {
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
	use_ssl = resource.GetScheme() != nil && resource.GetScheme().len_ > 3 && resource.GetScheme().val[3] == 's'

	/* use port 21 if one wasn't specified */

	if resource.GetPort() == 0 {
		resource.SetPort(21)
	}
	transport_len = int(zend.ZendSpprintf(&transport, 0, "tcp://%s:%d", resource.GetHost().val, resource.GetPort()))
	stream = streams._phpStreamXportCreate(transport, transport_len, 0x8, 0|2, nil, nil, context, nil, nil)
	zend._efree(transport)
	if stream == nil {
		result = 0
		goto connect_errexit
	}
	streams.PhpStreamContextSet(stream, context)
	if context != nil && context.notifier != nil {
		streams.PhpStreamNotificationNotify(context, 2, 0, nil, 0, 0, 0, nil)
	}

	/* Start talking to ftp server */

	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result > 299 || result < 200 {
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 9, 2, tmp_line, result, 0, 0, nil)
		}
		goto connect_errexit
	}
	if use_ssl != 0 {

		/* send the AUTH TLS request name */

		streams._phpStreamWrite(stream, "AUTH TLS\r\n", strlen("AUTH TLS\r\n"))

		/* get the response */

		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
		if result != 234 {

			/* AUTH TLS not supported try AUTH SSL */

			streams._phpStreamWrite(stream, "AUTH SSL\r\n", strlen("AUTH SSL\r\n"))

			/* get the response */

			result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
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
			streams._phpStreamFree(stream, 1|2)
			stream = nil
			goto connect_errexit
		}

		/* set PBSZ to 0 */

		streams._phpStreamWrite(stream, "PBSZ 0\r\n", strlen("PBSZ 0\r\n"))

		/* ignore the response */

		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))

		/* set data connection protection level */

		streams._phpStreamWrite(stream, "PROT P\r\n", strlen("PROT P\r\n"))

		/* get the response */

		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
		use_ssl_on_data = result >= 200 && result <= 299 || reuseid != nil
	}

	// #define PHP_FTP_CNTRL_CHK(val,val_len,err_msg) { unsigned char * s = ( unsigned char * ) val , * e = ( unsigned char * ) s + val_len ; while ( s < e ) { if ( iscntrl ( * s ) ) { php_stream_wrapper_log_error ( wrapper , options , err_msg , val ) ; goto connect_errexit ; } s ++ ; } }

	/* send the user name */

	if resource.GetUser() != nil {
		resource.GetUser().len_ = PhpRawUrlDecode(resource.GetUser().val, resource.GetUser().len_)
		var s *uint8 = (*uint8)(resource.GetUser().val)
		var e *uint8 = (*uint8)(s + resource.GetUser().len_)
		for s < e {
			if iscntrl(*s) {
				streams.PhpStreamWrapperLogError(wrapper, options, "Invalid login %s", resource.GetUser().val)
				goto connect_errexit
			}
			s++
		}
		streams._phpStreamPrintf(stream, "USER %s\r\n", resource.GetUser().val)
	} else {
		streams._phpStreamWrite(stream, "USER anonymous\r\n", strlen("USER anonymous\r\n"))
	}

	/* get the response */

	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))

	/* if a password is required, send it */

	if result >= 300 && result <= 399 {
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 3, 0, tmp_line, 0, 0, 0, nil)
		}
		if resource.GetPass() != nil {
			resource.GetPass().len_ = PhpRawUrlDecode(resource.GetPass().val, resource.GetPass().len_)
			var s *uint8 = (*uint8)(resource.GetPass().val)
			var e *uint8 = (*uint8)(s + resource.GetPass().len_)
			for s < e {
				if iscntrl(*s) {
					streams.PhpStreamWrapperLogError(wrapper, options, "Invalid password %s", resource.GetPass().val)
					goto connect_errexit
				}
				s++
			}
			streams._phpStreamPrintf(stream, "PASS %s\r\n", resource.GetPass().val)
		} else {

			/* if the user has configured who they are,
			   send that as the password */

			if FileGlobals.GetFromAddress() != nil {
				streams._phpStreamPrintf(stream, "PASS %s\r\n", FileGlobals.GetFromAddress())
			} else {
				streams._phpStreamWrite(stream, "PASS anonymous\r\n", strlen("PASS anonymous\r\n"))
			}

			/* if the user has configured who they are,
			   send that as the password */

		}

		/* read the response */

		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
		if result > 299 || result < 200 {
			if context != nil && context.notifier != nil {
				streams.PhpStreamNotificationNotify(context, 10, 2, tmp_line, result, 0, 0, nil)
			}
		} else {
			if context != nil && context.notifier != nil {
				streams.PhpStreamNotificationNotify(context, 10, 0, tmp_line, result, 0, 0, nil)
			}
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
		streams._phpStreamFree(stream, 1|2)
	}
	return nil
}

/* }}} */

func PhpFopenDoPasv(stream *core.PhpStream, ip *byte, ip_size int, phoststart **byte) uint16 {
	var tmp_line []byte
	var result int
	var i int
	var portno uint16
	var tpath *byte
	var ttpath *byte
	var hoststart *byte = nil

	/* We try EPSV first, needed for IPv6 and works on some IPv4 servers */

	streams._phpStreamWrite(stream, "EPSV\r\n", strlen("EPSV\r\n"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))

	/* check if we got a 229 response */

	if result != 229 {

		/* EPSV failed, let's try PASV */

		streams._phpStreamWrite(stream, "PASV\r\n", strlen("PASV\r\n"))
		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))

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

/* }}} */

func PhpStreamUrlWrapFtp(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
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
	var tmpzval *zend.Zval
	var allow_overwrite zend.ZendBool = 0
	var read_write int8_t = 0
	var transport *byte
	var transport_len int
	var error_message *zend.ZendString = nil
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
	if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "proxy")) != nil {
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

	streams._phpStreamWrite(stream, "TYPE I\r\n", strlen("TYPE I\r\n"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result > 299 || result < 200 {
		goto errexit
	}

	/* find out the size of the file (verifying it exists) */

	streams._phpStreamPrintf(stream, "SIZE %s\r\n", resource.GetPath().val)

	/* read the response */

	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
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
			if context != nil && context.notifier != nil {
				streams.PhpStreamNotificationNotify(context, 5, 0, tmp_line, result, 0, file_size, nil)
			}
		}
	} else if read_write == 2 {

		/* when writing file (but not appending), it must NOT exist, unless a context option exists which allows it */

		if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "overwrite")) != nil {
			if tmpzval.value.lval != 0 {
				allow_overwrite = 1
			} else {
				allow_overwrite = 0
			}
		}
		if result <= 299 && result >= 200 {
			if allow_overwrite != 0 {

				/* Context permits overwriting file,
				   so we just delete whatever's there in preparation */

				streams._phpStreamPrintf(stream, "DELE %s\r\n", resource.GetPath().val)
				result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
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

	portno = PhpFopenDoPasv(stream, ip, g.SizeOf("ip"), &hoststart)
	if portno == 0 {
		goto errexit
	}

	/* Send RETR/STOR command */

	if read_write == 1 {

		/* set resume position if applicable */

		if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ftp", "resume_pos")) != nil && tmpzval.u1.v.type_ == 4 && tmpzval.value.lval > 0 {
			streams._phpStreamPrintf(stream, "REST "+"%"+"lld"+"\r\n", tmpzval.value.lval)
			result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
			if result < 300 || result > 399 {
				streams.PhpStreamWrapperLogError(wrapper, options, "Unable to resume from offset "+"%"+"lld", tmpzval.value.lval)
				goto errexit
			}
		}

		/* retrieve file */

		memcpy(tmp_line, "RETR", g.SizeOf("\"RETR\""))

		/* retrieve file */

	} else if read_write == 2 {

		/* Write new file */

		memcpy(tmp_line, "STOR", g.SizeOf("\"STOR\""))

		/* Write new file */

	} else {

		/* Append */

		memcpy(tmp_line, "APPE", g.SizeOf("\"APPE\""))

		/* Append */

	}
	streams._phpStreamPrintf(stream, "%s %s\r\n", tmp_line, g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))

	/* open the data channel */

	if hoststart == nil {
		hoststart = resource.GetHost().val
	}
	transport_len = int(zend.ZendSpprintf(&transport, 0, "tcp://%s:%d", hoststart, portno))
	datastream = streams._phpStreamXportCreate(transport, transport_len, 0x8, 0|2, nil, nil, context, &error_message, nil)
	zend._efree(transport)
	if datastream == nil {
		tmp_line[0] = '0'
		goto errexit
	}
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result != 150 && result != 125 {

		/* Could not retrieve or send the file
		 * this data will only be sent to us after connection on the data port was initiated.
		 */

		streams._phpStreamFree(datastream, 1|2)
		datastream = nil
		goto errexit
	}
	streams.PhpStreamContextSet(datastream, context)
	if context != nil && context.notifier != nil {
		context.notifier.progress = 0
		context.notifier.progress_max = file_size
		context.notifier.mask |= 1
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 7, 0, nil, 0, 0, file_size, nil)
		}
	}
	if use_ssl_on_data != 0 && (streams.PhpStreamXportCryptoSetup(datastream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(datastream, 1) < 0) {
		streams.PhpStreamWrapperLogError(wrapper, options, "Unable to activate SSL mode")
		streams._phpStreamFree(datastream, 1|2)
		datastream = nil
		tmp_line[0] = '0'
		goto errexit
	}

	/* remember control stream */

	datastream.wrapperthis = stream
	PhpUrlFree(resource)
	return datastream
errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 9, 2, tmp_line, result, 0, 0, nil)
		}
		streams._phpStreamFree(stream, 1|2)
	}
	if tmp_line[0] != '0' {
		streams.PhpStreamWrapperLogError(wrapper, options, "FTP server reports %s", tmp_line)
	}
	if error_message != nil {
		streams.PhpStreamWrapperLogError(wrapper, options, "Failed to set up data channel: %s", error_message.val)
		zend.ZendStringRelease(error_message)
	}
	return nil
}

/* }}} */

func PhpFtpDirstreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)
	var innerstream *core.PhpStream
	var tmp_len int
	var basename *zend.ZendString
	innerstream = (*PhpFtpDirstreamData)(stream.abstract).GetDatastream()
	if count != g.SizeOf("php_stream_dirent") {
		return -1
	}
	if streams._phpStreamEof(innerstream) != 0 {
		return 0
	}
	if streams._phpStreamGetLine(innerstream, ent.d_name, g.SizeOf("ent -> d_name"), &tmp_len) == nil {
		return -1
	}
	basename = PhpBasename(ent.d_name, tmp_len, nil, 0)
	if g.SizeOf("ent -> d_name") < basename.len_-1 {
		tmp_len = g.SizeOf("ent -> d_name")
	} else {
		tmp_len = basename.len_ - 1
	}
	memcpy(ent.d_name, basename.val, tmp_len)
	ent.d_name[tmp_len-1] = '0'
	zend.ZendStringReleaseEx(basename, 0)

	/* Trim off trailing whitespace characters */

	for tmp_len > 0 && (ent.d_name[tmp_len-1] == '\n' || ent.d_name[tmp_len-1] == '\r' || ent.d_name[tmp_len-1] == '\t' || ent.d_name[tmp_len-1] == ' ') {
		ent.d_name[g.PreDec(&tmp_len)] = '0'
	}
	return g.SizeOf("php_stream_dirent")
}

/* }}} */

func PhpFtpDirstreamClose(stream *core.PhpStream, close_handle int) int {
	var data *PhpFtpDirstreamData = stream.abstract

	/* close control connection */

	if data.GetControlstream() != nil {
		streams._phpStreamFree(data.GetControlstream(), 1|2)
		data.SetControlstream(nil)
	}

	/* close data connection */

	streams._phpStreamFree(data.GetDatastream(), 1|2)
	data.SetDatastream(nil)
	zend._efree(data)
	stream.abstract = nil
	return 0
}

/* }}} */

var PhpFtpDirstreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpFtpDirstreamRead, PhpFtpDirstreamClose, nil, "ftpdir", nil, nil, nil, nil}

/* {{{ php_stream_ftp_opendir
 */

func PhpStreamFtpOpendir(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
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

	streams._phpStreamWrite(stream, "TYPE A\r\n", strlen("TYPE A\r\n"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result > 299 || result < 200 {
		goto opendir_errexit
	}

	// tmp_line isn't relevant after the php_fopen_do_pasv().

	tmp_line[0] = '0'

	/* set up the passive connection */

	portno = PhpFopenDoPasv(stream, ip, g.SizeOf("ip"), &hoststart)
	if portno == 0 {
		goto opendir_errexit
	}

	/* open the data channel */

	if hoststart == nil {
		hoststart = resource.GetHost().val
	}
	datastream = core._phpStreamSockOpenHost(hoststart, portno, SOCK_STREAM, 0, 0)
	if datastream == nil {
		goto opendir_errexit
	}
	streams._phpStreamPrintf(stream, "NLST %s\r\n", g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result != 150 && result != 125 {

		/* Could not retrieve or send the file
		 * this data will only be sent to us after connection on the data port was initiated.
		 */

		streams._phpStreamFree(datastream, 1|2)
		datastream = nil
		goto opendir_errexit
	}
	streams.PhpStreamContextSet(datastream, context)
	if use_ssl_on_data != 0 && (streams.PhpStreamXportCryptoSetup(datastream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(datastream, 1) < 0) {
		streams.PhpStreamWrapperLogError(wrapper, options, "Unable to activate SSL mode")
		streams._phpStreamFree(datastream, 1|2)
		datastream = nil
		goto opendir_errexit
	}
	PhpUrlFree(resource)
	dirsdata = zend._emalloc(sizeof * dirsdata)
	dirsdata.SetDatastream(datastream)
	dirsdata.SetControlstream(stream)
	dirsdata.SetDirstream(streams._phpStreamAlloc(&PhpFtpDirstreamOps, dirsdata, 0, mode))
	return dirsdata.GetDirstream()
opendir_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 9, 2, tmp_line, result, 0, 0, nil)
		}
		streams._phpStreamFree(stream, 1|2)
	}
	if tmp_line[0] != '0' {
		streams.PhpStreamWrapperLogError(wrapper, options, "FTP server reports %s", tmp_line)
	}
	return nil
}

/* }}} */

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
	ssb.sb.st_mode = 0644
	streams._phpStreamPrintf(stream, "CWD %s\r\n", g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {
		ssb.sb.st_mode |= S_IFREG
	} else {
		ssb.sb.st_mode |= S_IFDIR | S_IXUSR | S_IXGRP | S_IXOTH
	}
	streams._phpStreamWrite(stream, "TYPE I\r\n", strlen("TYPE I\r\n"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {
		goto stat_errexit
	}
	streams._phpStreamPrintf(stream, "SIZE %s\r\n", g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {

		/* Failure either means it doesn't exist
		   or it's a directory and this server
		   fails on listing directory sizes */

		if (ssb.sb.st_mode & S_IFDIR) != 0 {
			ssb.sb.st_size = 0
		} else {
			goto stat_errexit
		}

		/* Failure either means it doesn't exist
		   or it's a directory and this server
		   fails on listing directory sizes */

	} else {
		ssb.sb.st_size = atoi(tmp_line + 4)
	}
	streams._phpStreamPrintf(stream, "MDTM %s\r\n", g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result == 213 {
		var p *byte = tmp_line + 4
		var n int
		var tm __struct__tm
		var tmbuf __struct__tm
		var gmt *__struct__tm
		var stamp int64
		for size_t(p-tmp_line) < g.SizeOf("tmp_line") && !(isdigit(*p)) {
			p++
		}
		if size_t(p-tmp_line) > g.SizeOf("tmp_line") {
			goto mdtm_error
		}
		n = sscanf(p, "%4u%2u%2u%2u%2u%2u", &tm.tm_year, &tm.tm_mon, &tm.tm_mday, &tm.tm_hour, &tm.tm_min, &tm.tm_sec)
		if n != 6 {
			goto mdtm_error
		}
		tm.tm_year -= 1900
		tm.tm_mon--
		tm.tm_isdst = -1

		/* figure out the GMT offset */

		stamp = time(nil)
		gmt = gmtime_r(&stamp, &tmbuf)
		if gmt == nil {
			goto mdtm_error
		}
		gmt.tm_isdst = -1

		/* apply the GMT offset */

		tm.tm_sec += long(stamp - mktime(gmt))
		tm.tm_isdst = gmt.tm_isdst
		ssb.sb.st_mtime = mktime(&tm)
	} else {

		/* error or unsupported command */

	mdtm_error:
		ssb.sb.st_mtime = -1
	}
	ssb.sb.st_ino = 0
	ssb.sb.st_dev = 0
	ssb.sb.st_uid = 0
	ssb.sb.st_gid = 0
	ssb.sb.st_atime = -1
	ssb.sb.st_ctime = -1
	ssb.sb.st_nlink = 1
	ssb.sb.st_rdev = -1
	ssb.sb.st_blksize = 4096
	ssb.sb.st_blocks = int((4095 + ssb.sb.st_size) / ssb.sb.st_blksize)
	streams._phpStreamFree(stream, 1|2)
	PhpUrlFree(resource)
	return 0
stat_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		streams._phpStreamFree(stream, 1|2)
	}
	return -1
}

/* }}} */

func PhpStreamFtpUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Unable to connect to %s", url)
		}
		goto unlink_errexit
	}
	if resource.GetPath() == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Invalid path provided in %s", url)
		}
		goto unlink_errexit
	}

	/* Attempt to delete the file */

	streams._phpStreamPrintf(stream, "DELE %s\r\n", g.CondF1(resource.GetPath() != nil, func() []byte { return resource.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Error Deleting file: %s", tmp_line)
		}
		goto unlink_errexit
	}
	PhpUrlFree(resource)
	streams._phpStreamFree(stream, 1|2)
	return 1
unlink_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		streams._phpStreamFree(stream, 1|2)
	}
	return 0
}

/* }}} */

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

	if resource_from == nil || resource_to == nil || resource_from.GetScheme() == nil || resource_to.GetScheme() == nil || zend.ZendStringEquals(resource_from.GetScheme(), resource_to.GetScheme()) == 0 || resource_from.GetHost() == nil || resource_to.GetHost() == nil || zend.ZendStringEquals(resource_from.GetHost(), resource_to.GetHost()) == 0 || resource_from.GetPort() != resource_to.GetPort() && resource_from.GetPort()*resource_to.GetPort() != 0 && resource_from.GetPort()+resource_to.GetPort() != 21 || resource_from.GetPath() == nil || resource_to.GetPath() == nil {
		goto rename_errexit
	}
	stream = PhpFtpFopenConnect(wrapper, url_from, "r", 0, nil, context, nil, nil, nil, nil)
	if stream == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Unable to connect to %s", resource_from.GetHost().val)
		}
		goto rename_errexit
	}

	/* Rename FROM */

	streams._phpStreamPrintf(stream, "RNFR %s\r\n", g.CondF1(resource_from.GetPath() != nil, func() []byte { return resource_from.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 300 || result > 399 {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Error Renaming file: %s", tmp_line)
		}
		goto rename_errexit
	}

	/* Rename TO */

	streams._phpStreamPrintf(stream, "RNTO %s\r\n", g.CondF1(resource_to.GetPath() != nil, func() []byte { return resource_to.GetPath().val }, "/"))
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Error Renaming file: %s", tmp_line)
		}
		goto rename_errexit
	}
	PhpUrlFree(resource_from)
	PhpUrlFree(resource_to)
	streams._phpStreamFree(stream, 1|2)
	return 1
rename_errexit:
	if resource_from != nil {
		PhpUrlFree(resource_from)
	}
	if resource_to != nil {
		PhpUrlFree(resource_to)
	}
	if stream != nil {
		streams._phpStreamFree(stream, 1|2)
	}
	return 0
}

/* }}} */

func PhpStreamFtpMkdir(wrapper *core.PhpStreamWrapper, url *byte, mode int, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var recursive int = options & 1
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Unable to connect to %s", url)
		}
		goto mkdir_errexit
	}
	if resource.GetPath() == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Invalid path provided in %s", url)
		}
		goto mkdir_errexit
	}
	if recursive == 0 {
		streams._phpStreamPrintf(stream, "MKD %s\r\n", resource.GetPath().val)
		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	} else {

		/* we look for directory separator from the end of string, thus hopefully reducing our work load */

		var p *byte
		var e *byte
		var buf *byte
		buf = zend._estrndup(resource.GetPath().val, resource.GetPath().len_)
		e = buf + resource.GetPath().len_

		/* find a top level directory we need to create */

		for g.Assign(&p, strrchr(buf, '/')) {
			*p = '0'
			streams._phpStreamPrintf(stream, "CWD %s\r\n", g.Cond(strlen(buf), buf, "/"))
			result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
			if result >= 200 && result <= 299 {
				*p = '/'
				break
			}
		}
		streams._phpStreamPrintf(stream, "MKD %s\r\n", g.Cond(strlen(buf), buf, "/"))
		result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
		if result >= 200 && result <= 299 {
			if p == nil {
				p = buf
			}

			/* create any needed directories if the creation of the 1st directory worked */

			for p != e {
				if (*p) == '0' && (*(p + 1)) != '0' {
					*p = '/'
					streams._phpStreamPrintf(stream, "MKD %s\r\n", buf)
					result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
					if result < 200 || result > 299 {
						if (options & 0x8) != 0 {
							core.PhpErrorDocref(nil, 1<<1, "%s", tmp_line)
						}
						break
					}
				}
				p++
			}

			/* create any needed directories if the creation of the 1st directory worked */

		}
		zend._efree(buf)
	}
	PhpUrlFree(resource)
	streams._phpStreamFree(stream, 1|2)
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
		streams._phpStreamFree(stream, 1|2)
	}
	return 0
}

/* }}} */

func PhpStreamFtpRmdir(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var result int
	var tmp_line []byte
	stream = PhpFtpFopenConnect(wrapper, url, "r", 0, nil, context, nil, &resource, nil, nil)
	if stream == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Unable to connect to %s", url)
		}
		goto rmdir_errexit
	}
	if resource.GetPath() == nil {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "Invalid path provided in %s", url)
		}
		goto rmdir_errexit
	}
	streams._phpStreamPrintf(stream, "RMD %s\r\n", resource.GetPath().val)
	result = GetFtpResult(stream, tmp_line, g.SizeOf("tmp_line"))
	if result < 200 || result > 299 {
		if (options & 0x8) != 0 {
			core.PhpErrorDocref(nil, 1<<1, "%s", tmp_line)
		}
		goto rmdir_errexit
	}
	PhpUrlFree(resource)
	streams._phpStreamFree(stream, 1|2)
	return 1
rmdir_errexit:
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		streams._phpStreamFree(stream, 1|2)
	}
	return 0
}

/* }}} */

var FtpStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapFtp, PhpStreamFtpStreamClose, PhpStreamFtpStreamStat, PhpStreamFtpUrlStat, PhpStreamFtpOpendir, "ftp", PhpStreamFtpUnlink, PhpStreamFtpRename, PhpStreamFtpMkdir, PhpStreamFtpRmdir, nil}
var PhpStreamFtpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&FtpStreamWops, nil, 1}
