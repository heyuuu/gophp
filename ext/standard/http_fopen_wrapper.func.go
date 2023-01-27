// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
)

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
	for b.Assign(&s, strstr(s, header)) {
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
	var header_init int = (flags & HTTP_WRAPPER_HEADER_INIT) != 0
	var redirected int = (flags & HTTP_WRAPPER_REDIRECTED) != 0
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
	if !(zend.ZendStringEqualsLiteralCi(resource.GetScheme(), "http")) && !(zend.ZendStringEqualsLiteralCi(resource.GetScheme(), "https")) {
		if context == nil || b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.GetWops().GetLabel(), "proxy")) == nil || tmpzval.GetType() != zend.IS_STRING || zend.Z_STRLEN_P(tmpzval) == 0 {
			PhpUrlFree(resource)
			return core.PhpStreamOpenWrapperEx(path, mode, core.REPORT_ERRORS, nil, context)
		}

		/* Called from a non-http wrapper with http proxying requested (i.e. ftp) */

		request_fulluri = 1
		use_ssl = 0
		use_proxy = 1
		transport_len = zend.Z_STRLEN_P(tmpzval)
		transport_string = zend.Estrndup(zend.Z_STRVAL_P(tmpzval), zend.Z_STRLEN_P(tmpzval))
	} else {

		/* Normal http request (possibly with proxy) */

		if strpbrk(mode, "awx+") {
			streams.PhpStreamWrapperLogError(wrapper, options, "HTTP wrapper does not support writeable connections")
			PhpUrlFree(resource)
			return nil
		}
		use_ssl = resource.GetScheme() != nil && resource.GetScheme().GetLen() > 4 && resource.GetScheme().GetVal()[4] == 's'

		/* choose default ports */

		if use_ssl != 0 && resource.GetPort() == 0 {
			resource.SetPort(443)
		} else if resource.GetPort() == 0 {
			resource.SetPort(80)
		}
		if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.GetWops().GetLabel(), "proxy")) != nil && tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) > 0 {
			use_proxy = 1
			transport_len = zend.Z_STRLEN_P(tmpzval)
			transport_string = zend.Estrndup(zend.Z_STRVAL_P(tmpzval), zend.Z_STRLEN_P(tmpzval))
		} else {
			transport_len = core.Spprintf(&transport_string, 0, "%s://%s:%d", b.Cond(use_ssl != 0, "ssl", "tcp"), resource.GetHost().GetVal(), resource.GetPort())
		}
	}
	if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, wrapper.GetWops().GetLabel(), "timeout")) != nil {
		var d float64 = zend.ZvalGetDouble(tmpzval)
		timeout.tv_sec = int64(d)
		timeout.tv_usec = size_t((d - timeout.tv_sec) * 1000000)
	} else {
		timeout.tv_sec = FG(default_socket_timeout)
		timeout.tv_usec = 0
	}
	stream = streams.PhpStreamXportCreate(transport_string, transport_len, options, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, nil, &timeout, context, &errstr, nil)
	if stream != nil {
		core.PhpStreamSetOption(stream, core.PHP_STREAM_OPTION_READ_TIMEOUT, 0, &timeout)
	}
	if errstr != nil {
		streams.PhpStreamWrapperLogError(wrapper, options, "%s", errstr.GetVal())
		zend.ZendStringReleaseEx(errstr, 0)
		errstr = nil
	}
	zend.Efree(transport_string)
	if stream != nil && use_proxy != 0 && use_ssl != 0 {
		var header zend.SmartStr = zend.SmartStr{0}

		/* Set peer_name or name verification will try to use the proxy server name */

		if context == nil || b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "ssl", "peer_name")) == nil {
			zend.ZVAL_STR_COPY(&ssl_proxy_peer_name, resource.GetHost())
			streams.PhpStreamContextSetOption(core.PHP_STREAM_CONTEXT(stream), "ssl", "peer_name", &ssl_proxy_peer_name)
			zend.ZvalPtrDtor(&ssl_proxy_peer_name)
		}
		zend.SmartStrAppendl(&header, "CONNECT ", b.SizeOf("\"CONNECT \"")-1)
		zend.SmartStrAppends(&header, resource.GetHost().GetVal())
		zend.SmartStrAppendc(&header, ':')
		zend.SmartStrAppendUnsigned(&header, resource.GetPort())
		zend.SmartStrAppendl(&header, " HTTP/1.0\r\n", b.SizeOf("\" HTTP/1.0\\r\\n\"")-1)

		/* check if we have Proxy-Authorization header */

		if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "header")) != nil {
			var s *byte
			var p *byte
			if tmpzval.IsType(zend.IS_ARRAY) {
				var tmpheader *zend.Zval = nil
				for {
					var __ht *zend.HashTable = tmpzval.GetArr()
					var _p *zend.Bucket = __ht.GetArData()
					var _end *zend.Bucket = _p + __ht.GetNNumUsed()
					for ; _p != _end; _p++ {
						var _z *zend.Zval = _p.GetVal()

						if _z.IsType(zend.IS_UNDEF) {
							continue
						}
						tmpheader = _z
						if tmpheader.IsType(zend.IS_STRING) {
							s = zend.Z_STRVAL_P(tmpheader)
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
									if p-s == b.SizeOf("\"Proxy-Authorization:\"")-1 && zend.ZendBinaryStrcasecmp(s, b.SizeOf("\"Proxy-Authorization:\"")-1, "Proxy-Authorization:", b.SizeOf("\"Proxy-Authorization:\"")-1) == 0 {
										for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
											p++
										}
										zend.SmartStrAppendl(&header, s, p-s)
										zend.SmartStrAppendl(&header, "\r\n", b.SizeOf("\"\\r\\n\"")-1)
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
			} else if tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) != 0 {
				s = zend.Z_STRVAL_P(tmpzval)
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
						if p-s == b.SizeOf("\"Proxy-Authorization:\"")-1 && zend.ZendBinaryStrcasecmp(s, b.SizeOf("\"Proxy-Authorization:\"")-1, "Proxy-Authorization:", b.SizeOf("\"Proxy-Authorization:\"")-1) == 0 {
							for (*p) != 0 && (*p) != '\r' && (*p) != '\n' {
								p++
							}
							zend.SmartStrAppendl(&header, s, p-s)
							zend.SmartStrAppendl(&header, "\r\n", b.SizeOf("\"\\r\\n\"")-1)
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
		zend.SmartStrAppendl(&header, "\r\n", b.SizeOf("\"\\r\\n\"")-1)
		if core.PhpStreamWrite(stream, header.GetS().GetVal(), header.GetS().GetLen()) != header.GetS().GetLen() {
			streams.PhpStreamWrapperLogError(wrapper, options, "Cannot connect to HTTPS server through proxy")
			core.PhpStreamClose(stream)
			stream = nil
		}
		zend.SmartStrFree(&header)
		if stream != nil {
			var header_line []byte

			/* get response header */

			for core.PhpStreamGets(stream, header_line, HTTP_HEADER_BLOCK_SIZE-1) != nil {
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
				core.PhpStreamClose(stream)
				stream = nil
			}
		}

		/* enable SSL transport layer */

	}
	if stream == nil {
		goto out
	}

	/* avoid buffering issues while reading header */

	if (options & core.STREAM_WILL_CAST) != 0 {
		chunk_size = core.PhpStreamSetChunkSize(stream, 1)
	}

	/* avoid problems with auto-detecting when reading the headers -> the headers
	 * are always in canonical \r\n format */

	eol_detect = stream.GetFlags() & (core.PHP_STREAM_FLAG_DETECT_EOL | core.PHP_STREAM_FLAG_EOL_MAC)
	stream.SubFlags(core.PHP_STREAM_FLAG_DETECT_EOL | core.PHP_STREAM_FLAG_EOL_MAC)
	streams.PhpStreamContextSet(stream, context)
	streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_CONNECT, nil, 0)
	if header_init != 0 && context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "max_redirects")) != nil {
		redirect_max = int(zend.ZvalGetLong(tmpzval))
	}
	custom_request_method = 0
	if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "method")) != nil {
		if tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) > 0 {

			/* As per the RFC, automatically redirected requests MUST NOT use other methods than
			 * GET and HEAD unless it can be confirmed by the user */

			if redirected == 0 || zend.Z_STRLEN_P(tmpzval) == 3 && memcmp("GET", zend.Z_STRVAL_P(tmpzval), 3) == 0 || zend.Z_STRLEN_P(tmpzval) == 4 && memcmp("HEAD", zend.Z_STRVAL_P(tmpzval), 4) == 0 {
				custom_request_method = 1
				zend.SmartStrAppend(&req_buf, tmpzval.GetStr())
				zend.SmartStrAppendc(&req_buf, ' ')
			}

			/* As per the RFC, automatically redirected requests MUST NOT use other methods than
			 * GET and HEAD unless it can be confirmed by the user */

		}
	}
	if custom_request_method == 0 {
		zend.SmartStrAppends(&req_buf, "GET ")
	}

	/* Should we send the entire path in the request line, default to no. */

	if request_fulluri == 0 && context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "request_fulluri")) != nil {
		request_fulluri = zend.ZendIsTrue(tmpzval)
	}
	if request_fulluri != 0 {

		/* Ask for everything */

		zend.SmartStrAppends(&req_buf, path)

		/* Ask for everything */

	} else {

		/* Send the traditional /path/to/file?query_string */

		if resource.GetPath() != nil && resource.GetPath().GetLen() != 0 {
			zend.SmartStrAppends(&req_buf, resource.GetPath().GetVal())
		} else {
			zend.SmartStrAppendc(&req_buf, '/')
		}

		/* query string */

		if resource.GetQuery() != nil {
			zend.SmartStrAppendc(&req_buf, '?')
			zend.SmartStrAppends(&req_buf, resource.GetQuery().GetVal())
		}

		/* query string */

	}

	/* protocol version we are speaking */

	if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "protocol_version")) != nil {
		var protocol_version *byte
		core.Spprintf(&protocol_version, 0, "%.1F", zend.ZvalGetDouble(tmpzval))
		zend.SmartStrAppends(&req_buf, " HTTP/")
		zend.SmartStrAppends(&req_buf, protocol_version)
		zend.SmartStrAppends(&req_buf, "\r\n")
		zend.Efree(protocol_version)
	} else {
		zend.SmartStrAppends(&req_buf, " HTTP/1.0\r\n")
	}
	if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "header")) != nil {
		tmp = nil
		if tmpzval.IsType(zend.IS_ARRAY) {
			var tmpheader *zend.Zval = nil
			var tmpstr zend.SmartStr = zend.SmartStr{0}
			for {
				var __ht *zend.HashTable = tmpzval.GetArr()
				var _p *zend.Bucket = __ht.GetArData()
				var _end *zend.Bucket = _p + __ht.GetNNumUsed()
				for ; _p != _end; _p++ {
					var _z *zend.Zval = _p.GetVal()

					if _z.IsType(zend.IS_UNDEF) {
						continue
					}
					tmpheader = _z
					if tmpheader.IsType(zend.IS_STRING) {
						zend.SmartStrAppend(&tmpstr, tmpheader.GetStr())
						zend.SmartStrAppendl(&tmpstr, "\r\n", b.SizeOf("\"\\r\\n\"")-1)
					}
				}
				break
			}
			zend.SmartStr0(&tmpstr)

			/* Remove newlines and spaces from start and end. there's at least one extra \r\n at the end that needs to go. */

			if tmpstr.GetS() != nil {
				tmp = PhpTrim(tmpstr.GetS(), nil, 0, 3)
				zend.SmartStrFree(&tmpstr)
			}

			/* Remove newlines and spaces from start and end. there's at least one extra \r\n at the end that needs to go. */

		} else if tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) != 0 {

			/* Remove newlines and spaces from start and end php_trim will estrndup() */

			tmp = PhpTrim(tmpzval.GetStr(), nil, 0, 3)

			/* Remove newlines and spaces from start and end php_trim will estrndup() */

		}
		if tmp != nil && tmp.GetLen() != 0 {
			var s *byte
			var t *byte
			user_headers = zend.Estrndup(tmp.GetVal(), tmp.GetLen())
			if tmp.GetGcRefcount() > 1 {
				tmp.DecGcRefcount()
				tmp = zend.ZendStringInit(tmp.GetVal(), tmp.GetLen(), 0)
			}

			/* Make lowercase for easy comparison against 'standard' headers */

			PhpStrtolower(tmp.GetVal(), tmp.GetLen())
			t = tmp.GetVal()
			if header_init == 0 {

				/* strip POST headers on redirect */

				StripHeader(user_headers, t, "content-length:")
				StripHeader(user_headers, t, "content-type:")
			}
			if CheckHasHeader(t, "user-agent:") != 0 {
				have_header |= HTTP_HEADER_USER_AGENT
			}
			if CheckHasHeader(t, "host:") != 0 {
				have_header |= HTTP_HEADER_HOST
			}
			if CheckHasHeader(t, "from:") != 0 {
				have_header |= HTTP_HEADER_FROM
			}
			if CheckHasHeader(t, "authorization:") != 0 {
				have_header |= HTTP_HEADER_AUTH
			}
			if CheckHasHeader(t, "content-length:") != 0 {
				have_header |= HTTP_HEADER_CONTENT_LENGTH
			}
			if CheckHasHeader(t, "content-type:") != 0 {
				have_header |= HTTP_HEADER_TYPE
			}
			if CheckHasHeader(t, "connection:") != 0 {
				have_header |= HTTP_HEADER_CONNECTION
			}

			/* remove Proxy-Authorization header */

			if use_proxy != 0 && use_ssl != 0 && b.Assign(&s, strstr(t, "proxy-authorization:")) && (s == t || (*(s - 1)) == '\n') {
				var p *byte = s + b.SizeOf("\"proxy-authorization:\"") - 1
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
						zend.Efree(user_headers)
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

	if (have_header&HTTP_HEADER_AUTH) == 0 && resource.GetUser() != nil {

		/* make scratch large enough to hold the whole URL (over-estimate) */

		var scratch_len int = strlen(path) + 1
		var scratch *byte = zend.Emalloc(scratch_len)
		var stmp *zend.ZendString

		/* decode the strings first */

		PhpUrlDecode(resource.GetUser().GetVal(), resource.GetUser().GetLen())
		strcpy(scratch, resource.GetUser().GetVal())
		strcat(scratch, ":")

		/* Note: password is optional! */

		if resource.GetPass() != nil {
			PhpUrlDecode(resource.GetPass().GetVal(), resource.GetPass().GetLen())
			strcat(scratch, resource.GetPass().GetVal())
		}
		stmp = PhpBase64Encode((*uint8)(scratch), strlen(scratch))
		zend.SmartStrAppends(&req_buf, "Authorization: Basic ")
		zend.SmartStrAppends(&req_buf, stmp.GetVal())
		zend.SmartStrAppends(&req_buf, "\r\n")
		streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_AUTH_REQUIRED, nil, 0)
		zend.ZendStringFree(stmp)
		zend.Efree(scratch)
	}

	/* if the user has configured who they are, send a From: line */

	if (have_header&HTTP_HEADER_FROM) == 0 && FG(from_address) {
		zend.SmartStrAppends(&req_buf, "From: ")
		zend.SmartStrAppends(&req_buf, FG(from_address))
		zend.SmartStrAppends(&req_buf, "\r\n")
	}

	/* Send Host: header so name-based virtual hosts work */

	if (have_header & HTTP_HEADER_HOST) == 0 {
		zend.SmartStrAppends(&req_buf, "Host: ")
		zend.SmartStrAppends(&req_buf, resource.GetHost().GetVal())
		if use_ssl != 0 && resource.GetPort() != 443 && resource.GetPort() != 0 || use_ssl == 0 && resource.GetPort() != 80 && resource.GetPort() != 0 {
			zend.SmartStrAppendc(&req_buf, ':')
			zend.SmartStrAppendUnsigned(&req_buf, resource.GetPort())
		}
		zend.SmartStrAppends(&req_buf, "\r\n")
	}

	/* Send a Connection: close header to avoid hanging when the server
	 * interprets the RFC literally and establishes a keep-alive connection,
	 * unless the user specifically requests something else by specifying a
	 * Connection header in the context options. Send that header even for
	 * HTTP/1.0 to avoid issues when the server respond with a HTTP/1.1
	 * keep-alive response, which is the preferred response type. */

	if (have_header & HTTP_HEADER_CONNECTION) == 0 {
		zend.SmartStrAppends(&req_buf, "Connection: close\r\n")
	}
	if context != nil && b.Assign(&ua_zval, streams.PhpStreamContextGetOption(context, "http", "user_agent")) != nil && ua_zval.IsType(zend.IS_STRING) {
		ua_str = zend.Z_STRVAL_P(ua_zval)
	} else if FG(user_agent) {
		ua_str = FG(user_agent)
	}
	if (have_header&HTTP_HEADER_USER_AGENT) == 0 && ua_str != nil {
		const _UA_HEADER = "User-Agent: %s\r\n"
		var ua *byte
		var ua_len int
		ua_len = b.SizeOf("_UA_HEADER") + strlen(ua_str)

		/* ensure the header is only sent if user_agent is not blank */

		if ua_len > b.SizeOf("_UA_HEADER") {
			ua = zend.Emalloc(ua_len + 1)
			if b.Assign(&ua_len, core.Slprintf(ua, ua_len, _UA_HEADER, ua_str)) > 0 {
				ua[ua_len] = 0
				zend.SmartStrAppendl(&req_buf, ua, ua_len)
			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Cannot construct User-agent header")
			}
			zend.Efree(ua)
		}

		/* ensure the header is only sent if user_agent is not blank */

	}
	if user_headers != nil {

		/* A bit weird, but some servers require that Content-Length be sent prior to Content-Type for POST
		 * see bug #44603 for details. Since Content-Type maybe part of user's headers we need to do this check first.
		 */

		if header_init != 0 && context != nil && (have_header&HTTP_HEADER_CONTENT_LENGTH) == 0 && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "content")) != nil && tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) > 0 {
			zend.SmartStrAppends(&req_buf, "Content-Length: ")
			zend.SmartStrAppendUnsigned(&req_buf, zend.Z_STRLEN_P(tmpzval))
			zend.SmartStrAppends(&req_buf, "\r\n")
			have_header |= HTTP_HEADER_CONTENT_LENGTH
		}
		zend.SmartStrAppends(&req_buf, user_headers)
		zend.SmartStrAppends(&req_buf, "\r\n")
		zend.Efree(user_headers)
	}

	/* Request content, such as for POST requests */

	if header_init != 0 && context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "content")) != nil && tmpzval.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(tmpzval) > 0 {
		if (have_header & HTTP_HEADER_CONTENT_LENGTH) == 0 {
			zend.SmartStrAppends(&req_buf, "Content-Length: ")
			zend.SmartStrAppendUnsigned(&req_buf, zend.Z_STRLEN_P(tmpzval))
			zend.SmartStrAppends(&req_buf, "\r\n")
		}
		if (have_header & HTTP_HEADER_TYPE) == 0 {
			zend.SmartStrAppends(&req_buf, "Content-Type: application/x-www-form-urlencoded\r\n")
			core.PhpErrorDocref(nil, zend.E_NOTICE, "Content-type not specified assuming application/x-www-form-urlencoded")
		}
		zend.SmartStrAppends(&req_buf, "\r\n")
		zend.SmartStrAppendl(&req_buf, zend.Z_STRVAL_P(tmpzval), zend.Z_STRLEN_P(tmpzval))
	} else {
		zend.SmartStrAppends(&req_buf, "\r\n")
	}

	/* send it */

	core.PhpStreamWrite(stream, req_buf.GetS().GetVal(), req_buf.GetS().GetLen())
	location[0] = '0'
	if zend.Z_ISUNDEF_P(response_header) {
		zend.ArrayInit(response_header)
	}

	/* get response header */

	var tmp_line_len int
	if core.PhpStreamEof(stream) == 0 && core.PhpStreamGetLine(stream, tmp_line, b.SizeOf("tmp_line")-1, &tmp_line_len) != nil {
		var http_response zend.Zval
		if tmp_line_len > 9 {
			response_code = atoi(tmp_line + 9)
		} else {
			response_code = 0
		}
		if context != nil && nil != b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "ignore_errors")) {
			ignore_errors = zend.ZendIsTrue(tmpzval)
		}

		/* when we request only the header, don't fail even on error codes */

		if (options&core.STREAM_ONLY_GET_HEADERS) != 0 || ignore_errors != 0 {
			reqok = 1
		}

		/* status codes of 1xx are "informational", and will be followed by a real response
		 * e.g "100 Continue". RFC 7231 states that unexpected 1xx status MUST be parsed,
		 * and MAY be ignored. As such, we need to skip ahead to the "real" status*/

		if response_code >= 100 && response_code < 200 && response_code != 101 {

			/* consume lines until we find a line starting 'HTTP/1' */

			for core.PhpStreamEof(stream) == 0 && core.PhpStreamGetLine(stream, tmp_line, b.SizeOf("tmp_line")-1, &tmp_line_len) != nil && (tmp_line_len < b.SizeOf("\"HTTP/1\"")-1 || strncasecmp(tmp_line, "HTTP/1", b.SizeOf("\"HTTP/1\"")-1)) {

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
				streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_AUTH_RESULT, tmp_line, response_code)
				break
			default:

				/* safety net in the event tmp_line == NULL */

				if tmp_line_len == 0 {
					tmp_line[0] = '0'
				}
				streams.PhpStreamNotifyError(context, streams.PHP_STREAM_NOTIFY_FAILURE, tmp_line, response_code)
			}
		}
		if tmp_line_len >= 1 && tmp_line[tmp_line_len-1] == '\n' {
			tmp_line_len--
			if tmp_line_len >= 1 && tmp_line[tmp_line_len-1] == '\r' {
				tmp_line_len--
			}
		}
		zend.ZVAL_STRINGL(&http_response, tmp_line, tmp_line_len)
		response_header.GetArr().NextIndexInsert(&http_response)
	} else {
		core.PhpStreamClose(stream)
		stream = nil
		streams.PhpStreamWrapperLogError(wrapper, options, "HTTP request failed!")
		goto out
	}

	/* read past HTTP headers */

	for core.PhpStreamEof(stream) == 0 {
		var http_header_line_length int
		if http_header_line != nil {
			zend.Efree(http_header_line)
		}
		if b.Assign(&http_header_line, core.PhpStreamGetLine(stream, nil, 0, &http_header_line_length)) && (*http_header_line) != '\n' && (*http_header_line) != '\r' {
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
			if !(strncasecmp(http_header_line, "Location:", b.SizeOf("\"Location:\"")-1)) {
				if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "follow_location")) != nil {
					follow_location = zend.ZvalIsTrue(tmpzval)
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
				strlcpy(location, http_header_value, b.SizeOf("location"))
			} else if !(strncasecmp(http_header_line, "Content-Type:", b.SizeOf("\"Content-Type:\"")-1)) {
				streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_MIME_TYPE_IS, http_header_value, 0)
			} else if !(strncasecmp(http_header_line, "Content-Length:", b.SizeOf("\"Content-Length:\"")-1)) {
				file_size = atoi(http_header_value)
				streams.PhpStreamNotifyFileSize(context, file_size, http_header_line, 0)
			} else if !(strncasecmp(http_header_line, "Transfer-Encoding:", b.SizeOf("\"Transfer-Encoding:\"")-1)) && !(strncasecmp(http_header_value, "Chunked", b.SizeOf("\"Chunked\"")-1)) {

				/* create filter to decode response body */

				if (options & core.STREAM_ONLY_GET_HEADERS) == 0 {
					var decode zend.ZendLong = 1
					if context != nil && b.Assign(&tmpzval, streams.PhpStreamContextGetOption(context, "http", "auto_decode")) != nil {
						decode = zend.ZendIsTrue(tmpzval)
					}
					if decode != 0 {
						transfer_encoding = streams.PhpStreamFilterCreate("dechunk", nil, stream.GetIsPersistent())
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
			zend.ZVAL_STRINGL(&http_header, http_header_line, http_header_line_length)
			response_header.GetArr().NextIndexInsert(&http_header)
		} else {
			break
		}
	}
	if reqok == 0 || location[0] != '0' && follow_location != 0 {
		if follow_location == 0 || ((options&core.STREAM_ONLY_GET_HEADERS) != 0 || ignore_errors != 0) && redirect_max <= 1 {
			goto out
		}
		if location[0] != '0' {
			streams.PhpStreamNotifyInfo(context, streams.PHP_STREAM_NOTIFY_REDIRECTED, location, 0)
		}
		core.PhpStreamClose(stream)
		stream = nil
		if location[0] != '0' {
			var new_path []byte
			var loc_path []byte
			*new_path = '0'
			if strlen(location) < 8 || strncasecmp(location, "http://", b.SizeOf("\"http://\"")-1) && strncasecmp(location, "https://", b.SizeOf("\"https://\"")-1) && strncasecmp(location, "ftp://", b.SizeOf("\"ftp://\"")-1) && strncasecmp(location, "ftps://", b.SizeOf("\"ftps://\"")-1) {
				if (*location) != '/' {
					if (*(location + 1)) != '0' && resource.GetPath() != nil {
						var s *byte = strrchr(resource.GetPath().GetVal(), '/')
						if s == nil {
							s = resource.GetPath().GetVal()
							if resource.GetPath().GetLen() == 0 {
								zend.ZendStringReleaseEx(resource.GetPath(), 0)
								resource.SetPath(zend.ZendStringInit("/", 1, 0))
								s = resource.GetPath().GetVal()
							} else {
								*s = '/'
							}
						}
						s[1] = '0'
						if resource.GetPath() != nil && resource.GetPath().GetVal()[0] == '/' && resource.GetPath().GetVal()[1] == '0' {
							core.Snprintf(loc_path, b.SizeOf("loc_path")-1, "%s%s", resource.GetPath().GetVal(), location)
						} else {
							core.Snprintf(loc_path, b.SizeOf("loc_path")-1, "%s/%s", resource.GetPath().GetVal(), location)
						}
					} else {
						core.Snprintf(loc_path, b.SizeOf("loc_path")-1, "/%s", location)
					}
				} else {
					strlcpy(loc_path, location, b.SizeOf("loc_path"))
				}
				if use_ssl != 0 && resource.GetPort() != 443 || use_ssl == 0 && resource.GetPort() != 80 {
					core.Snprintf(new_path, b.SizeOf("new_path")-1, "%s://%s:%d%s", resource.GetScheme().GetVal(), resource.GetHost().GetVal(), resource.GetPort(), loc_path)
				} else {
					core.Snprintf(new_path, b.SizeOf("new_path")-1, "%s://%s%s", resource.GetScheme().GetVal(), resource.GetHost().GetVal(), loc_path)
				}
			} else {
				strlcpy(new_path, location, b.SizeOf("new_path"))
			}
			PhpUrlFree(resource)

			/* check for invalid redirection URLs */

			if b.Assign(&resource, PhpUrlParse(new_path)) == nil {
				streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
				goto out
			}
			var CHECK_FOR_CNTRL_CHARS func(val *zend.ZendString) = func(val *zend.ZendString) {
				if val != nil {
					var s *uint8
					var e *uint8
					val.SetLen(PhpUrlDecode(val.GetVal(), val.GetLen()))
					s = (*uint8)(val.GetVal())
					e = s + val.GetLen()
					for s < e {
						if iscntrl(*s) {
							streams.PhpStreamWrapperLogError(wrapper, options, "Invalid redirect URL! %s", new_path)
							goto out
						}
						s++
					}
				}
			}

			/* check for control characters in login, password & path */

			if strncasecmp(new_path, "http://", b.SizeOf("\"http://\"")-1) || strncasecmp(new_path, "https://", b.SizeOf("\"https://\"")-1) {
				CHECK_FOR_CNTRL_CHARS(resource.GetUser())
				CHECK_FOR_CNTRL_CHARS(resource.GetPass())
				CHECK_FOR_CNTRL_CHARS(resource.GetPath())
			}
			stream = PhpStreamUrlWrapHttpEx(wrapper, new_path, mode, options, opened_path, context, b.PreDec(&redirect_max), HTTP_WRAPPER_REDIRECTED, response_header)
		} else {
			streams.PhpStreamWrapperLogError(wrapper, options, "HTTP request failed! %s", tmp_line)
		}
	}
out:
	zend.SmartStrFree(&req_buf)
	if http_header_line != nil {
		zend.Efree(http_header_line)
	}
	if resource != nil {
		PhpUrlFree(resource)
	}
	if stream != nil {
		if header_init != 0 {
			zend.ZVAL_COPY(stream.GetWrapperdata(), response_header)
		}
		streams.PhpStreamNotifyProgressInit(context, 0, file_size)

		/* Restore original chunk size now that we're done with headers */

		if (options & core.STREAM_WILL_CAST) != 0 {
			core.PhpStreamSetChunkSize(stream, int(chunk_size))
		}

		/* restore the users auto-detect-line-endings setting */

		stream.AddFlags(eol_detect)

		/* as far as streams are concerned, we are now at the start of
		 * the stream */

		stream.SetPosition(0)

		/* restore mode */

		strlcpy(stream.GetMode(), mode, b.SizeOf("stream -> mode"))
		if transfer_encoding != nil {
			streams.PhpStreamFilterAppend(stream.GetReadfilters(), transfer_encoding)
		}
	} else {
		if transfer_encoding != nil {
			streams.PhpStreamFilterFree(transfer_encoding)
		}
	}
	return stream
}
func PhpStreamUrlWrapHttp(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var stream *core.PhpStream
	var headers zend.Zval
	zend.ZVAL_UNDEF(&headers)
	stream = PhpStreamUrlWrapHttpEx(wrapper, path, mode, options, opened_path, context, PHP_URL_REDIRECT_MAX, HTTP_WRAPPER_HEADER_INIT, &headers)
	if !(zend.Z_ISUNDEF(headers)) {
		if zend.FAILURE == zend.ZendSetLocalVarStr("http_response_header", b.SizeOf("\"http_response_header\"")-1, &headers, 1) {
			zend.ZvalPtrDtor(&headers)
		}
	}
	return stream
}
func PhpStreamHttpStreamStat(wrapper *core.PhpStreamWrapper, stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	/* one day, we could fill in the details based on Date: and Content-Length:
	 * headers.  For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */

	return -1

	/* one day, we could fill in the details based on Date: and Content-Length:
	 * headers.  For now, we return with a failure code to prevent the underlying
	 * file's details from being used instead. */
}
