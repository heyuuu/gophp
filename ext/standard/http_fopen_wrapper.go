// <<generate>>

package standard

import (
	"sik/core"
	"sik/core/streams"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/http_fopen_wrapper.c>

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
   |          Wez Furlong <wez@thebrainroom.com>                          |
   |          Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "php_streams.h"

// # include "php_network.h"

// # include "php_ini.h"

// # include "ext/standard/basic_functions.h"

// # include "zend_smart_str.h"

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

// #define HTTP_HEADER_BLOCK_SIZE       1024

// #define PHP_URL_REDIRECT_MAX       20

// #define HTTP_HEADER_USER_AGENT       1

// #define HTTP_HEADER_HOST       2

// #define HTTP_HEADER_AUTH       4

// #define HTTP_HEADER_FROM       8

// #define HTTP_HEADER_CONTENT_LENGTH       16

// #define HTTP_HEADER_TYPE       32

// #define HTTP_HEADER_CONNECTION       64

// #define HTTP_WRAPPER_HEADER_INIT       1

// #define HTTP_WRAPPER_REDIRECTED       2

func StripHeader(header_bag *byte, lc_header_bag *byte, lc_header_name string) {
	var lc_header_start *byte = strstr(lc_header_bag, lc_header_name)
	if lc_header_start != nil && (lc_header_start == lc_header_bag || (*(lc_header_start - 1)) == '\n') {
		var header_start *byte = header_bag + (lc_header_start - lc_header_bag)
		var lc_eol *byte = strchr(lc_header_start, '\n')
		if lc_eol != nil {
			var eol *byte = header_start + (lc_eol - lc_header_start)
			var eollen int = strlen(lc_eol)
			memmove(lc_header_start, lc_eol+1, eollen)
			memmove(header_start, eol+1, eollen)
		} else {
			*lc_header_start = '0'
			*header_start = '0'
		}
	}
}
func CheckHasHeader(headers *byte, header string) zend.ZendBool {
	var s *byte = headers
	for g.Assign(&s, strstr(s, header)) {
		if s == headers || (*(s - 1)) == '\n' {
			return 1
		}
		s++
	}
	return 0
}
func PhpStreamUrlWrapHttpEx(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext, redirect_max int, flags int, response_header *zend.Zval) *core.PhpStream {
	var stream *core.PhpStream = nil
	var resource *PhpUrl = nil
	var use_ssl int
	var use_proxy int = 0
	var tmp *zend.ZendString = nil
	var ua_str *byte = nil
	var ua_zval *zend.Zval = nil
	var tmpzval *zend.Zval = nil
	var ssl_proxy_peer_name zend.Zval
	var location []byte
	var reqok int = 0
	var http_header_line *byte = nil
	var tmp_line []byte
	var chunk_size int = 0
	var file_size int = 0
	var eol_detect int = 0
	var transport_string *byte
	var errstr *zend.ZendString = nil
	var transport_len int
	var have_header int = 0
	var request_fulluri zend.ZendBool = 0
	var ignore_errors zend.ZendBool = 0
	var timeout __struct__timeval
	var user_headers *byte = nil
	var header_init int = (flags & 1) != 0
	var redirected int = (flags & 2) != 0
	var follow_location zend.ZendBool = 1
	var transfer_encoding *core.PhpStreamFilter = nil
	var response_code int
	var req_buf zend.SmartStr = zend.SmartStr{0}
	var custom_request_method zend.ZendBool
	tmp_line[0] = '0'
	if redirect_max < 1 {
		streams.PhpStreamWrapperLogError(wrapper, options, "Redirection limit reached, aborting")
		return nil
	}
	resource = PhpUrlParse(path)
	if resource == nil {
		return nil
	}
	if !(resource.GetScheme().len_ == g.SizeOf("\"http\"")-1 && zend.ZendBinaryStrcasecmp(resource.GetScheme().val, resource.GetScheme().len_, "http", g.SizeOf("\"http\"")-1) == 0) && !(resource.GetScheme().len_ == g.SizeOf("\"https\"")-1 && zend.ZendBinaryStrcasecmp(resource.GetScheme().val, resource.GetScheme().len_, "https", g.SizeOf("\"https\"")-1) == 0) {
		if context == nil || g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.wops.label, "proxy")) == nil || tmpzval.u1.v.type_ != 6 || tmpzval.value.str.len_ == 0 {
			PhpUrlFree(resource)
			return streams._phpStreamOpenWrapperEx(path, mode, 0x8, nil, context)
		}

		/* Called from a non-http wrapper with http proxying requested (i.e. ftp) */

		request_fulluri = 1
		use_ssl = 0
		use_proxy = 1
		transport_len = tmpzval.value.str.len_
		transport_string = zend._estrndup(tmpzval.value.str.val, tmpzval.value.str.len_)
	} else {

		/* Normal http request (possibly with proxy) */

		if strpbrk(mode, "awx+") {
			streams.PhpStreamWrapperLogError(wrapper, options, "HTTP wrapper does not support writeable connections")
			PhpUrlFree(resource)
			return nil
		}
		use_ssl = resource.GetScheme() != nil && resource.GetScheme().len_ > 4 && resource.GetScheme().val[4] == 's'

		/* choose default ports */

		if use_ssl != 0 && resource.GetPort() == 0 {
			resource.SetPort(443)
		} else if resource.GetPort() == 0 {
			resource.SetPort(80)
		}
		if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.wops.label, "proxy")) != nil && tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ > 0 {
			use_proxy = 1
			transport_len = tmpzval.value.str.len_
			transport_string = zend._estrndup(tmpzval.value.str.val, tmpzval.value.str.len_)
		} else {
			transport_len = zend.ZendSpprintf(&transport_string, 0, "%s://%s:%d", g.Cond(use_ssl != 0, "ssl", "tcp"), resource.GetHost().val, resource.GetPort())
		}
	}
	if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.wops.label, "timeout")) != nil {
		var d float64 = zend.ZvalGetDouble(tmpzval)
		timeout.tv_sec = int64(d)
		timeout.tv_usec = size_t((d - timeout.tv_sec) * 1000000)
	} else {
		timeout.tv_sec = FileGlobals.GetDefaultSocketTimeout()
		timeout.tv_usec = 0
	}
	stream = streams._phpStreamXportCreate(transport_string, transport_len, options, 0|2, nil, &timeout, context, &errstr, nil)
	if stream != nil {
		streams._phpStreamSetOption(stream, 4, 0, &timeout)
	}
	if errstr != nil {
		streams.PhpStreamWrapperLogError(wrapper, options, "%s", errstr.val)
		zend.ZendStringReleaseEx(errstr, 0)
		errstr = nil
	}
	zend._efree(transport_string)
	if stream != nil && use_proxy != 0 && use_ssl != 0 {
		var header zend.SmartStr = zend.SmartStr{0}

		/* Set peer_name or name verification will try to use the proxy server name */

		if context == nil || g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ssl", "peer_name")) == nil {
			var __z *zend.Zval = &ssl_proxy_peer_name
			var __s *zend.ZendString = resource.GetHost()
			__z.value.str = __s
			if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
				__z.u1.type_info = 6
			} else {
				zend.ZendGcAddref(&__s.gc)
				__z.u1.type_info = 6 | 1<<0<<8
			}
			streams.PhpStreamContextSetOption((*core.PhpStreamContext)(g.CondF1(stream.ctx != nil, func() any { return stream.ctx.ptr }, nil)), "ssl", "peer_name", &ssl_proxy_peer_name)
			zend.ZvalPtrDtor(&ssl_proxy_peer_name)
		}
		zend.SmartStrAppendlEx(&header, "CONNECT ", g.SizeOf("\"CONNECT \"")-1, 0)
		zend.SmartStrAppendlEx(&header, resource.GetHost().val, strlen(resource.GetHost().val), 0)
		zend.SmartStrAppendcEx(&header, ':', 0)
		zend.SmartStrAppendUnsignedEx(&header, resource.GetPort(), 0)
		zend.SmartStrAppendlEx(&header, " HTTP/1.0\r\n", g.SizeOf("\" HTTP/1.0\\r\\n\"")-1, 0)

		/* check if we have Proxy-Authorization header */

		if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "header")) != nil {
			var s *byte
			var p *byte
			if tmpzval.u1.v.type_ == 7 {
				var tmpheader *zend.Zval = nil
				for {
					var __ht *zend.HashTable = tmpzval.value.arr
					var _p *zend.Bucket = __ht.arData
					var _end *zend.Bucket = _p + __ht.nNumUsed
					for ; _p != _end; _p++ {
						var _z *zend.Zval = &_p.val

						if _z.u1.v.type_ == 0 {
							continue
						}
						tmpheader = _z
						if tmpheader.u1.v.type_ == 6 {
							s = tmpheader.value.str.val
							for {
								for (*s) == ' ' || (*s) == '\t' {
									s++
								}
								p = s
								for (*p) != 0 && (*p) != ':' && (*p) != '\r' && (*p) != '\n' {
									p++
								}
								if (*p) == ':' {
									p++
									if p-s == g.SizeOf("\"Proxy-Authorization:\"")-1 && zend.ZendBinaryStrcasecmp(s, g.SizeOf("\"Proxy-Authorization:\"")-1, "Proxy-Authorization:", g.SizeOf("\"Proxy-Authorization:\"")-1) == 0 {
										for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
											p++
										}
										zend.SmartStrAppendlEx(&header, s, p-s, 0)
										zend.SmartStrAppendlEx(&header, "\r\n", g.SizeOf("\"\\r\\n\"")-1, 0)
										goto finish
									} else {
										for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
											p++
										}
									}
								}
								s = p
								for (*s) == '\r' || (*s) == '\n' {
									s++
								}
								if (*s) == 0 {
									break
								}
							}
						}
					}
					break
				}
			} else if tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ != 0 {
				s = tmpzval.value.str.val
				for {
					for (*s) == ' ' || (*s) == '\t' {
						s++
					}
					p = s
					for (*p) != 0 && (*p) != ':' && (*p) != '\r' && (*p) != '\n' {
						p++
					}
					if (*p) == ':' {
						p++
						if p-s == g.SizeOf("\"Proxy-Authorization:\"")-1 && zend.ZendBinaryStrcasecmp(s, g.SizeOf("\"Proxy-Authorization:\"")-1, "Proxy-Authorization:", g.SizeOf("\"Proxy-Authorization:\"")-1) == 0 {
							for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
								p++
							}
							zend.SmartStrAppendlEx(&header, s, p-s, 0)
							zend.SmartStrAppendlEx(&header, "\r\n", g.SizeOf("\"\\r\\n\"")-1, 0)
							goto finish
						} else {
							for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
								p++
							}
						}
					}
					s = p
					for (*s) == '\r' || (*s) == '\n' {
						s++
					}
					if (*s) == 0 {
						break
					}
				}
			}
		}
	finish:
		zend.SmartStrAppendlEx(&header, "\r\n", g.SizeOf("\"\\r\\n\"")-1, 0)
		if streams._phpStreamWrite(stream, header.s.val, header.s.len_) != header.s.len_ {
			streams.PhpStreamWrapperLogError(wrapper, options, "Cannot connect to HTTPS server through proxy")
			streams._phpStreamFree(stream, 1|2)
			stream = nil
		}
		zend.SmartStrFreeEx(&header, 0)
		if stream != nil {
			var header_line []byte

			/* get response header */

			for streams._phpStreamGetLine(stream, header_line, 1024-1, nil) != nil {
				if header_line[0] == '\n' || header_line[0] == '\r' || header_line[0] == '0' {
					break
				}
			}

			/* get response header */

		}

		/* enable SSL transport layer */

		if stream != nil {
			if streams.PhpStreamXportCryptoSetup(stream, streams.STREAM_CRYPTO_METHOD_SSLv23_CLIENT, nil) < 0 || streams.PhpStreamXportCryptoEnable(stream, 1) < 0 {
				streams.PhpStreamWrapperLogError(wrapper, options, "Cannot connect to HTTPS server through proxy")
				streams._phpStreamFree(stream, 1|2)
				stream = nil
			}
		}

		/* enable SSL transport layer */

	}
	if stream == nil {
		goto out
	}

	/* avoid buffering issues while reading header */

	if (options & 0x20) != 0 {
		chunk_size = streams._phpStreamSetOption(stream, 5, 1, nil)
	}

	/* avoid problems with auto-detecting when reading the headers -> the headers
	 * are always in canonical \r\n format */

	eol_detect = stream.flags & (0x4 | 0x8)
	stream.flags &= ^(0x4 | 0x8)
	streams.PhpStreamContextSet(stream, context)
	if context != nil && context.notifier != nil {
		streams.PhpStreamNotificationNotify(context, 2, 0, nil, 0, 0, 0, nil)
	}
	if header_init != 0 && context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "max_redirects")) != nil {
		redirect_max = int(zend.ZvalGetLong(tmpzval))
	}
	custom_request_method = 0
	if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "method")) != nil {
		if tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ > 0 {

			/* As per the RFC, automatically redirected requests MUST NOT use other methods than
			 * GET and HEAD unless it can be confirmed by the user */

			if redirected == 0 || tmpzval.value.str.len_ == 3 && memcmp("GET", tmpzval.value.str.val, 3) == 0 || tmpzval.value.str.len_ == 4 && memcmp("HEAD", tmpzval.value.str.val, 4) == 0 {
				custom_request_method = 1
				zend.SmartStrAppendEx(&req_buf, tmpzval.value.str, 0)
				zend.SmartStrAppendcEx(&req_buf, ' ', 0)
			}

			/* As per the RFC, automatically redirected requests MUST NOT use other methods than
			 * GET and HEAD unless it can be confirmed by the user */

		}
	}
	if custom_request_method == 0 {
		zend.SmartStrAppendlEx(&req_buf, "GET ", strlen("GET "), 0)
	}

	/* Should we send the entire path in the request line, default to no. */

	if request_fulluri == 0 && context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "request_fulluri")) != nil {
		request_fulluri = zend.ZendIsTrue(tmpzval)
	}
	if request_fulluri != 0 {

		/* Ask for everything */

		zend.SmartStrAppendlEx(&req_buf, path, strlen(path), 0)

		/* Ask for everything */

	} else {

		/* Send the traditional /path/to/file?query_string */

		if resource.GetPath() != nil && resource.GetPath().len_ != 0 {
			zend.SmartStrAppendlEx(&req_buf, resource.GetPath().val, strlen(resource.GetPath().val), 0)
		} else {
			zend.SmartStrAppendcEx(&req_buf, '/', 0)
		}

		/* query string */

		if resource.GetQuery() != nil {
			zend.SmartStrAppendcEx(&req_buf, '?', 0)
			zend.SmartStrAppendlEx(&req_buf, resource.GetQuery().val, strlen(resource.GetQuery().val), 0)
		}

		/* query string */

	}

	/* protocol version we are speaking */

	if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "protocol_version")) != nil {
		var protocol_version *byte
		zend.ZendSpprintf(&protocol_version, 0, "%.1F", zend.ZvalGetDouble(tmpzval))
		zend.SmartStrAppendlEx(&req_buf, " HTTP/", strlen(" HTTP/"), 0)
		zend.SmartStrAppendlEx(&req_buf, protocol_version, strlen(protocol_version), 0)
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
		zend._efree(protocol_version)
	} else {
		zend.SmartStrAppendlEx(&req_buf, " HTTP/1.0\r\n", strlen(" HTTP/1.0\r\n"), 0)
	}
	if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "header")) != nil {
		tmp = nil
		if tmpzval.u1.v.type_ == 7 {
			var tmpheader *zend.Zval = nil
			var tmpstr zend.SmartStr = zend.SmartStr{0}
			for {
				var __ht *zend.HashTable = tmpzval.value.arr
				var _p *zend.Bucket = __ht.arData
				var _end *zend.Bucket = _p + __ht.nNumUsed
				for ; _p != _end; _p++ {
					var _z *zend.Zval = &_p.val

					if _z.u1.v.type_ == 0 {
						continue
					}
					tmpheader = _z
					if tmpheader.u1.v.type_ == 6 {
						zend.SmartStrAppendEx(&tmpstr, tmpheader.value.str, 0)
						zend.SmartStrAppendlEx(&tmpstr, "\r\n", g.SizeOf("\"\\r\\n\"")-1, 0)
					}
				}
				break
			}
			zend.SmartStr0(&tmpstr)

			/* Remove newlines and spaces from start and end. there's at least one extra \r\n at the end that needs to go. */

			if tmpstr.s != nil {
				tmp = PhpTrim(tmpstr.s, nil, 0, 3)
				zend.SmartStrFreeEx(&tmpstr, 0)
			}

			/* Remove newlines and spaces from start and end. there's at least one extra \r\n at the end that needs to go. */

		} else if tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ != 0 {

			/* Remove newlines and spaces from start and end php_trim will estrndup() */

			tmp = PhpTrim(tmpzval.value.str, nil, 0, 3)

			/* Remove newlines and spaces from start and end php_trim will estrndup() */

		}
		if tmp != nil && tmp.len_ != 0 {
			var s *byte
			var t *byte
			user_headers = zend._estrndup(tmp.val, tmp.len_)
			if (zend.ZvalGcFlags(tmp.gc.u.type_info) & 1 << 6) != 0 {
				tmp = zend.ZendStringInit(tmp.val, tmp.len_, 0)
			} else if zend.ZendGcRefcount(&tmp.gc) > 1 {
				zend.ZendGcDelref(&tmp.gc)
				tmp = zend.ZendStringInit(tmp.val, tmp.len_, 0)
			}

			/* Make lowercase for easy comparison against 'standard' headers */

			PhpStrtolower(tmp.val, tmp.len_)
			t = tmp.val
			if header_init == 0 {

				/* strip POST headers on redirect */

				StripHeader(user_headers, t, "content-length:")
				StripHeader(user_headers, t, "content-type:")
			}
			if CheckHasHeader(t, "user-agent:") != 0 {
				have_header |= 1
			}
			if CheckHasHeader(t, "host:") != 0 {
				have_header |= 2
			}
			if CheckHasHeader(t, "from:") != 0 {
				have_header |= 8
			}
			if CheckHasHeader(t, "authorization:") != 0 {
				have_header |= 4
			}
			if CheckHasHeader(t, "content-length:") != 0 {
				have_header |= 16
			}
			if CheckHasHeader(t, "content-type:") != 0 {
				have_header |= 32
			}
			if CheckHasHeader(t, "connection:") != 0 {
				have_header |= 64
			}

			/* remove Proxy-Authorization header */

			if use_proxy != 0 && use_ssl != 0 && g.Assign(&s, strstr(t, "proxy-authorization:")) && (s == t || (*(s - 1)) == '\n') {
				var p *byte = s + g.SizeOf("\"proxy-authorization:\"") - 1
				for s > t && ((*(s - 1)) == ' ' || (*(s - 1)) == '\t') {
					s--
				}
				for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
					p++
				}
				for (*p) == '\r' || (*p) == '\n' {
					p++
				}
				if (*p) == 0 {
					if s == t {
						zend._efree(user_headers)
						user_headers = nil
					} else {
						for s > t && ((*(s - 1)) == '\r' || (*(s - 1)) == '\n') {
							s--
						}
						user_headers[s-t] = 0
					}
				} else {
					memmove(user_headers+(s-t), user_headers+(p-t), strlen(p)+1)
				}
			}

			/* remove Proxy-Authorization header */

		}
		if tmp != nil {
			zend.ZendStringReleaseEx(tmp, 0)
		}
	}

	/* auth header if it was specified */

	if (have_header&4) == 0 && resource.GetUser() != nil {

		/* make scratch large enough to hold the whole URL (over-estimate) */

		var scratch_len int = strlen(path) + 1
		var scratch *byte = zend._emalloc(scratch_len)
		var stmp *zend.ZendString

		/* decode the strings first */

		PhpUrlDecode(resource.GetUser().val, resource.GetUser().len_)
		strcpy(scratch, resource.GetUser().val)
		strcat(scratch, ":")

		/* Note: password is optional! */

		if resource.GetPass() != nil {
			PhpUrlDecode(resource.GetPass().val, resource.GetPass().len_)
			strcat(scratch, resource.GetPass().val)
		}
		stmp = PhpBase64Encode((*uint8)(scratch), strlen(scratch))
		zend.SmartStrAppendlEx(&req_buf, "Authorization: Basic ", strlen("Authorization: Basic "), 0)
		zend.SmartStrAppendlEx(&req_buf, stmp.val, strlen(stmp.val), 0)
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
		if context != nil && context.notifier != nil {
			streams.PhpStreamNotificationNotify(context, 3, 0, nil, 0, 0, 0, nil)
		}
		zend.ZendStringFree(stmp)
		zend._efree(scratch)
	}

	/* if the user has configured who they are, send a From: line */

	if (have_header&8) == 0 && FileGlobals.GetFromAddress() != nil {
		zend.SmartStrAppendlEx(&req_buf, "From: ", strlen("From: "), 0)
		zend.SmartStrAppendlEx(&req_buf, FileGlobals.GetFromAddress(), strlen(FileGlobals.GetFromAddress()), 0)
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
	}

	/* Send Host: header so name-based virtual hosts work */

	if (have_header & 2) == 0 {
		zend.SmartStrAppendlEx(&req_buf, "Host: ", strlen("Host: "), 0)
		zend.SmartStrAppendlEx(&req_buf, resource.GetHost().val, strlen(resource.GetHost().val), 0)
		if use_ssl != 0 && resource.GetPort() != 443 && resource.GetPort() != 0 || use_ssl == 0 && resource.GetPort() != 80 && resource.GetPort() != 0 {
			zend.SmartStrAppendcEx(&req_buf, ':', 0)
			zend.SmartStrAppendUnsignedEx(&req_buf, resource.GetPort(), 0)
		}
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
	}

	/* Send a Connection: close header to avoid hanging when the server
	 * interprets the RFC literally and establishes a keep-alive connection,
	 * unless the user specifically requests something else by specifying a
	 * Connection header in the context options. Send that header even for
	 * HTTP/1.0 to avoid issues when the server respond with a HTTP/1.1
	 * keep-alive response, which is the preferred response type. */

	if (have_header & 64) == 0 {
		zend.SmartStrAppendlEx(&req_buf, "Connection: close\r\n", strlen("Connection: close\r\n"), 0)
	}
	if context != nil && g.Assign(&ua_zval, streams.PhpStreamContextGetOption(context, "http", "user_agent")) != nil && ua_zval.u1.v.type_ == 6 {
		ua_str = ua_zval.value.str.val
	} else if FileGlobals.GetUserAgent() != nil {
		ua_str = FileGlobals.GetUserAgent()
	}
	if (have_header&1) == 0 && ua_str != nil {

		// #define _UA_HEADER       "User-Agent: %s\r\n"

		var ua *byte
		var ua_len int
		ua_len = g.SizeOf("_UA_HEADER") + strlen(ua_str)

		/* ensure the header is only sent if user_agent is not blank */

		if ua_len > g.SizeOf("_UA_HEADER") {
			ua = zend._emalloc(ua_len + 1)
			if g.Assign(&ua_len, core.ApPhpSlprintf(ua, ua_len, "User-Agent: %s\r\n", ua_str)) > 0 {
				ua[ua_len] = 0
				zend.SmartStrAppendlEx(&req_buf, ua, ua_len, 0)
			} else {
				core.PhpErrorDocref(nil, 1<<1, "Cannot construct User-agent header")
			}
			zend._efree(ua)
		}

		/* ensure the header is only sent if user_agent is not blank */

	}
	if user_headers != nil {

		/* A bit weird, but some servers require that Content-Length be sent prior to Content-Type for POST
		 * see bug #44603 for details. Since Content-Type maybe part of user's headers we need to do this check first.
		 */

		if header_init != 0 && context != nil && (have_header&16) == 0 && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "content")) != nil && tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ > 0 {
			zend.SmartStrAppendlEx(&req_buf, "Content-Length: ", strlen("Content-Length: "), 0)
			zend.SmartStrAppendUnsignedEx(&req_buf, tmpzval.value.str.len_, 0)
			zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
			have_header |= 16
		}
		zend.SmartStrAppendlEx(&req_buf, user_headers, strlen(user_headers), 0)
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
		zend._efree(user_headers)
	}

	/* Request content, such as for POST requests */

	if header_init != 0 && context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "content")) != nil && tmpzval.u1.v.type_ == 6 && tmpzval.value.str.len_ > 0 {
		if (have_header & 16) == 0 {
			zend.SmartStrAppendlEx(&req_buf, "Content-Length: ", strlen("Content-Length: "), 0)
			zend.SmartStrAppendUnsignedEx(&req_buf, tmpzval.value.str.len_, 0)
			zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
		}
		if (have_header & 32) == 0 {
			zend.SmartStrAppendlEx(&req_buf, "Content-Type: application/x-www-form-urlencoded\r\n", strlen("Content-Type: application/x-www-form-urlencoded\r\n"), 0)
			core.PhpErrorDocref(nil, 1<<3, "Content-type not specified assuming application/x-www-form-urlencoded")
		}
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
		zend.SmartStrAppendlEx(&req_buf, tmpzval.value.str.val, tmpzval.value.str.len_, 0)
	} else {
		zend.SmartStrAppendlEx(&req_buf, "\r\n", strlen("\r\n"), 0)
	}

	/* send it */

	streams._phpStreamWrite(stream, req_buf.s.val, req_buf.s.len_)
	location[0] = '0'
	if response_header.u1.v.type_ == 0 {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = response_header
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	}

	/* get response header */

	var tmp_line_len int
	if streams._phpStreamEof(stream) == 0 && streams._phpStreamGetLine(stream, tmp_line, g.SizeOf("tmp_line")-1, &tmp_line_len) != nil {
		var http_response zend.Zval
		if tmp_line_len > 9 {
			response_code = atoi(tmp_line + 9)
		} else {
			response_code = 0
		}
		if context != nil && nil != g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "ignore_errors")) {
			ignore_errors = zend.ZendIsTrue(tmpzval)
		}

		/* when we request only the header, don't fail even on error codes */

		if (options&0x200) != 0 || ignore_errors != 0 {
			reqok = 1
		}

		/* status codes of 1xx are "informational", and will be followed by a real response
		 * e.g "100 Continue". RFC 7231 states that unexpected 1xx status MUST be parsed,
		 * and MAY be ignored. As such, we need to skip ahead to the "real" status*/

		if response_code >= 100 && response_code < 200 && response_code != 101 {

			/* consume lines until we find a line starting 'HTTP/1' */

			for streams._phpStreamEof(stream) == 0 && streams._phpStreamGetLine(stream, tmp_line, g.SizeOf("tmp_line")-1, &tmp_line_len) != nil && (tmp_line_len < g.SizeOf("\"HTTP/1\"")-1 || strncasecmp(tmp_line, "HTTP/1", g.SizeOf("\"HTTP/1\"")-1)) {

			}
			if tmp_line_len > 9 {
				response_code = atoi(tmp_line + 9)
			} else {
				response_code = 0
			}
		}

		/* all status codes in the 2xx range are defined by the specification as successful;
		 * all status codes in the 3xx range are for redirection, and so also should never
		 * fail */

		if response_code >= 200 && response_code < 400 {
			reqok = 1
		} else {
			switch response_code {
			case 403:
				if context != nil && context.notifier != nil {
					streams.PhpStreamNotificationNotify(context, 10, 2, tmp_line, response_code, 0, 0, nil)
				}
				break
			default:

				/* safety net in the event tmp_line == NULL */

				if tmp_line_len == 0 {
					tmp_line[0] = '0'
				}
				if context != nil && context.notifier != nil {
					streams.PhpStreamNotificationNotify(context, 9, 2, tmp_line, response_code, 0, 0, nil)
				}
			}
		}
		if tmp_line_len >= 1 && tmp_line[tmp_line_len-1] == '\n' {
			tmp_line_len--
			if tmp_line_len >= 1 && tmp_line[tmp_line_len-1] == '\r' {
				tmp_line_len--
			}
		}
		var __z *zend.Zval = &http_response
		var __s *zend.ZendString = zend.ZendStringInit(tmp_line, tmp_line_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendHashNextIndexInsert(response_header.value.arr, &http_response)
	} else {
		streams._phpStreamFree(stream, 1|2)
		stream = nil
		streams.PhpStreamWrapperLogError(wrapper, options, "HTTP request failed!")
		goto out
	}

	/* read past HTTP headers */

	for streams._phpStreamEof(stream) == 0 {
		var http_header_line_length int
		if http_header_line != nil {
			zend._efree(http_header_line)
		}
		if g.Assign(&http_header_line, streams._phpStreamGetLine(stream, nil, 0, &http_header_line_length)) && (*http_header_line) != '\n' && (*http_header_line) != '\r' {
			var e *byte = http_header_line + http_header_line_length - 1
			var http_header_value *byte
			for e >= http_header_line && ((*e) == '\n' || (*e) == '\r') {
				e--
			}

			/* The primary definition of an HTTP header in RFC 7230 states:
			 * > Each header field consists of a case-insensitive field name followed
			 * > by a colon (":"), optional leading whitespace, the field value, and
			 * > optional trailing whitespace. */

			for e >= http_header_line && ((*e) == ' ' || (*e) == '\t') {
				e--
			}

			/* Terminate header line */

			e++
			*e = '0'
			http_header_line_length = e - http_header_line
			http_header_value = memchr(http_header_line, ':', http_header_line_length)
			if http_header_value != nil {
				http_header_value++

				/* Strip leading whitespace */

				for http_header_value < e && ((*http_header_value) == ' ' || (*http_header_value) == '\t') {
					http_header_value++
				}

				/* Strip leading whitespace */

			} else {

				/* There is no colon. Set the value to the end of the header line, which is
				 * effectively an empty string. */

				http_header_value = e

				/* There is no colon. Set the value to the end of the header line, which is
				 * effectively an empty string. */

			}
			if !(strncasecmp(http_header_line, "Location:", g.SizeOf("\"Location:\"")-1)) {
				if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "follow_location")) != nil {
					follow_location = zend.ZendIsTrue(tmpzval)
				} else if !(response_code >= 300 && response_code < 304 || 307 == response_code || 308 == response_code) {

					/* we shouldn't redirect automatically
					   if follow_location isn't set and response_code not in (300, 301, 302, 303 and 307)
					   see http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.3.1
					   RFC 7238 defines 308: http://tools.ietf.org/html/rfc7238 */

					follow_location = 0

					/* we shouldn't redirect automatically
					   if follow_location isn't set and response_code not in (300, 301, 302, 303 and 307)
					   see http://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.3.1
					   RFC 7238 defines 308: http://tools.ietf.org/html/rfc7238 */

				}
				strlcpy(location, http_header_value, g.SizeOf("location"))
			} else if !(strncasecmp(http_header_line, "Content-Type:", g.SizeOf("\"Content-Type:\"")-1)) {
				if context != nil && context.notifier != nil {
					streams.PhpStreamNotificationNotify(context, 4, 0, http_header_value, 0, 0, 0, nil)
				}
			} else if !(strncasecmp(http_header_line, "Content-Length:", g.SizeOf("\"Content-Length:\"")-1)) {
				file_size = atoi(http_header_value)
				if context != nil && context.notifier != nil {
					streams.PhpStreamNotificationNotify(context, 5, 0, http_header_line, 0, 0, file_size, nil)
				}
			} else if !(strncasecmp(http_header_line, "Transfer-Encoding:", g.SizeOf("\"Transfer-Encoding:\"")-1)) && !(strncasecmp(http_header_value, "Chunked", g.SizeOf("\"Chunked\"")-1)) {

				/* create filter to decode response body */

				if (options & 0x200) == 0 {
					var decode zend.ZendLong = 1
					if context != nil && g.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "auto_decode")) != nil {
						decode = zend.ZendIsTrue(tmpzval)
					}
					if decode != 0 {
						transfer_encoding = streams.PhpStreamFilterCreate("dechunk", nil, stream.is_persistent)
						if transfer_encoding != nil {

							/* don't store transfer-encodeing header */

							continue

							/* don't store transfer-encodeing header */

						}
					}
				}

				/* create filter to decode response body */

			}
			var http_header zend.Zval
			var __z *zend.Zval = &http_header
			var __s *zend.ZendString = zend.ZendStringInit(http_header_line, http_header_line_length, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendHashNextIndexInsert(response_header.value.arr, &http_header)
		} else {
			break
		}
	}
	if reqok == 0 || location[0] != '0' && follow_location != 0 {
		if follow_location == 0 || ((options&0x200) != 0 || ignore_errors != 0) && redirect_max <= 1 {
			goto out
		}
		if location[0] != '0' {
			if context != nil && context.notifier != nil {
				streams.PhpStreamNotificationNotify(context, 6, 0, location, 0, 0, 0, nil)
			}
		}
		streams._phpStreamFree(stream, 1|2)
		stream = nil
		if location[0] != '0' {
			var new_path []byte
			var loc_path []byte
			*new_path = '0'
			if strlen(location) < 8 || strncasecmp(location, "http://", g.SizeOf("\"http://\"")-1) && strncasecmp(location, "https://", g.SizeOf("\"https://\"")-1) && strncasecmp(location, "ftp://", g.SizeOf("\"ftp://\"")-1) && strncasecmp(location, "ftps://", g.SizeOf("\"ftps://\"")-1) {
				if (*location) != '/' {
					if (*(location + 1)) != '0' && resource.GetPath() != nil {
						var s *byte = strrchr(resource.GetPath().val, '/')
						if s == nil {
							s = resource.GetPath().val
							if resource.GetPath().len_ == 0 {
								zend.ZendStringReleaseEx(resource.GetPath(), 0)
								resource.SetPath(zend.ZendStringInit("/", 1, 0))
								s = resource.GetPath().val
							} else {
								*s = '/'
							}
						}
						s[1] = '0'
						if resource.GetPath() != nil && resource.GetPath().val[0] == '/' && resource.GetPath().val[1] == '0' {
							core.ApPhpSnprintf(loc_path, g.SizeOf("loc_path")-1, "%s%s", resource.GetPath().val, location)
						} else {
							core.ApPhpSnprintf(loc_path, g.SizeOf("loc_path")-1, "%s/%s", resource.GetPath().val, location)
						}
					} else {
						core.ApPhpSnprintf(loc_path, g.SizeOf("loc_path")-1, "/%s", location)
					}
				} else {
					strlcpy(loc_path, location, g.SizeOf("loc_path"))
				}
				if use_ssl != 0 && resource.GetPort() != 443 || use_ssl == 0 && resource.GetPort() != 80 {
					core.ApPhpSnprintf(new_path, g.SizeOf("new_path")-1, "%s://%s:%d%s", resource.GetScheme().val, resource.GetHost().val, resource.GetPort(), loc_path)
				} else {
					core.ApPhpSnprintf(new_path, g.SizeOf("new_path")-1, "%s://%s%s", resource.GetScheme().val, resource.GetHost().val, loc_path)
				}
			} else {
				strlcpy(new_path, location, g.SizeOf("new_path"))
			}
			PhpUrlFree(resource)

			/* check for invalid redirection URLs */

			if g.Assign(&resource, PhpUrlParse(new_path)) == nil {
				streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
				goto out
			}

			// #define CHECK_FOR_CNTRL_CHARS(val) { if ( val ) { unsigned char * s , * e ; ZSTR_LEN ( val ) = php_url_decode ( ZSTR_VAL ( val ) , ZSTR_LEN ( val ) ) ; s = ( unsigned char * ) ZSTR_VAL ( val ) ; e = s + ZSTR_LEN ( val ) ; while ( s < e ) { if ( iscntrl ( * s ) ) { php_stream_wrapper_log_error ( wrapper , options , "Invalid redirect URL! %s" , new_path ) ; goto out ; } s ++ ; } } }

			/* check for control characters in login, password & path */

			if strncasecmp(new_path, "http://", g.SizeOf("\"http://\"")-1) || strncasecmp(new_path, "https://", g.SizeOf("\"https://\"")-1) {
				if resource.GetUser() != nil {
					var s *uint8
					var e *uint8
					resource.GetUser().len_ = PhpUrlDecode(resource.GetUser().val, resource.GetUser().len_)
					s = (*uint8)(resource.GetUser().val)
					e = s + resource.GetUser().len_
					for s < e {
						if iscntrl(*s) {
							streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
							goto out
						}
						s++
					}
				}
				if resource.GetPass() != nil {
					var s *uint8
					var e *uint8
					resource.GetPass().len_ = PhpUrlDecode(resource.GetPass().val, resource.GetPass().len_)
					s = (*uint8)(resource.GetPass().val)
					e = s + resource.GetPass().len_
					for s < e {
						if iscntrl(*s) {
							streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
							goto out
						}
						s++
					}
				}
				if resource.GetPath() != nil {
					var s *uint8
					var e *uint8
					resource.GetPath().len_ = PhpUrlDecode(resource.GetPath().val, resource.GetPath().len_)
					s = (*uint8)(resource.GetPath().val)
					e = s + resource.GetPath().len_
					for s < e {
						if iscntrl(*s) {
							streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
							goto out
						}
						s++
					}
				}
			}
			stream = PhpStreamUrlWrapHttpEx(wrapper, new_path, mode, options, opened_path, context, g.PreDec(&redirect_max), 2, response_header)
		} else {
			streams.PhpStreamWrapperLogError(wrapper, options, "HTTP request failed! %s", tmp_line)
		}
	}
out:
	zend.SmartStrFreeEx(&req_buf, 0)
	if http_header_line != nil {
		zend._efree(http_header_line)
	}
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		if header_init != 0 {
			var _z1 *zend.Zval = &stream.wrapperdata
			var _z2 *zend.Zval = response_header
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
		}
		if context != nil && context.notifier != nil {
			context.notifier.progress = 0
			context.notifier.progress_max = file_size
			context.notifier.mask |= 1
			if context != nil && context.notifier != nil {
				streams.PhpStreamNotificationNotify(context, 7, 0, nil, 0, 0, file_size, nil)
			}
		}

		/* Restore original chunk size now that we're done with headers */

		if (options & 0x20) != 0 {
			streams._phpStreamSetOption(stream, 5, int(chunk_size), nil)
		}

		/* restore the users auto-detect-line-endings setting */

		stream.flags |= eol_detect

		/* as far as streams are concerned, we are now at the start of
		 * the stream */

		stream.position = 0

		/* restore mode */

		strlcpy(stream.mode, mode, g.SizeOf("stream -> mode"))
		if transfer_encoding != nil {
			streams._phpStreamFilterAppend(&stream.readfilters, transfer_encoding)
		}
	} else {
		if transfer_encoding != nil {
			streams.PhpStreamFilterFree(transfer_encoding)
		}
	}
	return stream
}

/* }}} */

func PhpStreamUrlWrapHttp(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var stream *core.PhpStream
	var headers zend.Zval
	&headers.u1.type_info = 0
	stream = PhpStreamUrlWrapHttpEx(wrapper, path, mode, options, opened_path, context, 20, 1, &headers)
	if headers.u1.v.type_ != 0 {
		if zend.FAILURE == zend.ZendSetLocalVarStr("http_response_header", g.SizeOf("\"http_response_header\"")-1, &headers, 1) {
			zend.ZvalPtrDtor(&headers)
		}
	}
	return stream
}

/* }}} */

func PhpStreamHttpStreamStat(wrapper *core.PhpStreamWrapper, stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	/* one day, we could fill in the details based on Date: and Content-Length:
	 * headers.  For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */

	return -1

	/* one day, we could fill in the details based on Date: and Content-Length:
	 * headers.  For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */
}

/* }}} */

var HttpStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapHttp, nil, PhpStreamHttpStreamStat, nil, nil, "http", nil, nil, nil, nil, nil}
var PhpStreamHttpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&HttpStreamWops, nil, 1}
