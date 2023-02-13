// <<generate>>

package cli

import (
	b "sik/builtin"
	r "sik/runtime"
)

func MIN(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func PARSING_HEADER(state State) bool {
	return state <= SHeadersAlmostDone && 0 == (parser.flags&F_TRAILING)
}
func LOWER(c byte) uint8    { return uint8(c | 0x20) }
func TOKEN(c byte) byte     { return Tokens[uint8(c)] }
func NEW_MESSAGE() __auto__ { return StartState }
func PhpHttpParserExecute(parser *PhpHttpParser, settings *PhpHttpParserSettings, data *byte, len_ int) int {
	var ch byte
	var c signed__char
	var p *byte = data
	var pe *byte
	var to_read int
	var state State = State(parser.GetState())
	var header_state HeaderStates = HeaderStates(parser.GetHeaderState())
	var index uint32 = parser.GetIndex()
	var nread uint32 = parser.GetNread()

	var header_field_mark *byte = 0
	var header_value_mark *byte = 0
	var fragment_mark *byte = 0
	var query_string_mark *byte = 0
	var path_mark *byte = 0
	var url_mark *byte = 0
	if len_ == 0 {
		if state == SBodyIdentityEof {
			if settings.GetOnMessageComplete() != nil {
				if 0 != settings.GetOnMessageComplete()(parser) {
					return p - data
				}
			}
		}
		return 0
	}
	if state == SHeaderField {
		header_field_mark = data
	}
	if state == SHeaderValue {
		header_value_mark = data
	}
	if state == SReqFragment {
		fragment_mark = data
	}
	if state == SReqQueryString {
		query_string_mark = data
	}
	if state == SReqPath {
		path_mark = data
	}
	if state == SReqPath || state == SReqSchema || state == SReqSchemaSlash || state == SReqSchemaSlashSlash || state == SReqPort || state == SReqQueryStringStart || state == SReqQueryString || state == SReqHost || state == SReqFragmentStart || state == SReqFragment {
		url_mark = data
	}
	p = data
	pe = data + len_
	for ; p != pe; p++ {
		ch = *p
		if PARSING_HEADER(state) {
			nread++

			/* Buffer overflow attack */

			if nread > PHP_HTTP_MAX_HEADER_SIZE {
				goto error
			}

			/* Buffer overflow attack */

		}
		switch state {
		case SDead:

			/* this state is used after a 'Connection: close' message
			 * the parser will error out if it reads another message
			 */

			goto error
		case SStartReqOrRes:
			if ch == CR || ch == LF {
				break
			}
			parser.SetFlags(0)
			parser.SetContentLength(-1)
			if ch == 'H' {
				state = s_res_or_resp_H
			} else {
				parser.SetType(PHP_HTTP_REQUEST)
				goto start_req_method_assign
			}
			break
		case s_res_or_resp_H:
			if ch == 'T' {
				parser.SetType(PHP_HTTP_RESPONSE)
				state = s_res_HT
			} else {
				if ch != 'E' {
					goto error
				}
				parser.SetType(PHP_HTTP_REQUEST)
				parser.SetMethod(PHP_HTTP_HEAD)
				index = 2
				state = SReqMethod
			}
			break
		case SStartRes:
			parser.SetFlags(0)
			parser.SetContentLength(-1)
			switch ch {
			case 'H':
				state = s_res_H
				break
			case CR:

			case LF:
				break
			default:
				goto error
			}
			break
		case s_res_H:
			state = s_res_HT
			break
		case s_res_HT:
			state = s_res_HTT
			break
		case s_res_HTT:
			state = s_res_HTTP
			break
		case s_res_HTTP:
			state = SResFirstHttpMajor
			break
		case SResFirstHttpMajor:
			if ch < '1' || ch > '9' {
				goto error
			}
			parser.SetHttpMajor(ch - '0')
			state = SResHttpMajor
			break
		case SResHttpMajor:
			if ch == '.' {
				state = SResFirstHttpMinor
				break
			}
			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMajor(parser.GetHttpMajor() * 10)
			parser.SetHttpMajor(parser.GetHttpMajor() + ch - '0')
			if parser.GetHttpMajor() > 999 {
				goto error
			}
			break
		case SResFirstHttpMinor:
			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMinor(ch - '0')
			state = SResHttpMinor
			break
		case SResHttpMinor:
			if ch == ' ' {
				state = SResFirstStatusCode
				break
			}
			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMinor(parser.GetHttpMinor() * 10)
			parser.SetHttpMinor(parser.GetHttpMinor() + ch - '0')
			if parser.GetHttpMinor() > 999 {
				goto error
			}
			break
		case SResFirstStatusCode:
			if ch < '0' || ch > '9' {
				if ch == ' ' {
					break
				}
				goto error
			}
			parser.SetStatusCode(ch - '0')
			state = SResStatusCode
			break
		case SResStatusCode:
			if ch < '0' || ch > '9' {
				switch ch {
				case ' ':
					state = SResStatus
					break
				case CR:
					state = SResLineAlmostDone
					break
				case LF:
					state = SHeaderFieldStart
					break
				default:
					goto error
				}
				break
			}
			parser.SetStatusCode(parser.GetStatusCode() * 10)
			parser.SetStatusCode(parser.GetStatusCode() + ch - '0')
			if parser.GetStatusCode() > 999 {
				goto error
			}
			break
		case SResStatus:

			/* the human readable status. e.g. "NOT FOUND"
			 * we are not humans so just ignore this */

			if ch == CR {
				state = SResLineAlmostDone
				break
			}
			if ch == LF {
				state = SHeaderFieldStart
				break
			}
			break
		case SResLineAlmostDone:
			state = SHeaderFieldStart
			break
		case SStartReq:
			if ch == CR || ch == LF {
				break
			}
			parser.SetFlags(0)
			parser.SetContentLength(-1)
			if ch < 'A' || 'Z' < ch {
				goto error
			}
		start_req_method_assign:
			parser.SetMethod(PhpHttpMethod(0))
			index = 1
			switch ch {
			case 'C':
				parser.SetMethod(PHP_HTTP_CONNECT)
				break
			case 'D':
				parser.SetMethod(PHP_HTTP_DELETE)
				break
			case 'G':
				parser.SetMethod(PHP_HTTP_GET)
				break
			case 'H':
				parser.SetMethod(PHP_HTTP_HEAD)
				break
			case 'L':
				parser.SetMethod(PHP_HTTP_LOCK)
				break
			case 'M':
				parser.SetMethod(PHP_HTTP_MKCOL)
				break
			case 'N':
				parser.SetMethod(PHP_HTTP_NOTIFY)
				break
			case 'O':
				parser.SetMethod(PHP_HTTP_OPTIONS)
				break
			case 'P':
				parser.SetMethod(PHP_HTTP_POST)
				break
			case 'R':
				parser.SetMethod(PHP_HTTP_REPORT)
				break
			case 'S':
				parser.SetMethod(PHP_HTTP_SUBSCRIBE)
				break
			case 'T':
				parser.SetMethod(PHP_HTTP_TRACE)
				break
			case 'U':
				parser.SetMethod(PHP_HTTP_UNLOCK)
				break
			default:
				parser.SetMethod(PHP_HTTP_NOT_IMPLEMENTED)
				break
			}
			state = SReqMethod
			break
		case SReqMethod:
			var matcher *byte
			if ch == '0' {
				goto error
			}
			matcher = MethodStrings[parser.GetMethod()]
			if ch == ' ' {
				if parser.GetMethod() != PHP_HTTP_NOT_IMPLEMENTED && matcher[index] != '0' {
					parser.SetMethod(PHP_HTTP_NOT_IMPLEMENTED)
				}
				state = SReqSpacesBeforeUrl
			} else if parser.GetMethod() == PHP_HTTP_NOT_IMPLEMENTED || ch == matcher[index] {

			} else if parser.GetMethod() == PHP_HTTP_CONNECT {
				if index == 1 && ch == 'H' {
					parser.SetMethod(PHP_HTTP_CHECKOUT)
				} else if index == 2 && ch == 'P' {
					parser.SetMethod(PHP_HTTP_COPY)
				} else {
					parser.SetMethod(PHP_HTTP_NOT_IMPLEMENTED)
				}
			} else if parser.GetMethod() == PHP_HTTP_MKCOL {
				if index == 1 && ch == 'O' {
					parser.SetMethod(PHP_HTTP_MOVE)
				} else if index == 3 && ch == 'A' {
					parser.SetMethod(PHP_HTTP_MKCALENDAR)
				} else if index == 1 && ch == 'E' {
					parser.SetMethod(PHP_HTTP_MERGE)
				} else if index == 1 && ch == '-' {
					parser.SetMethod(PHP_HTTP_MSEARCH)
				} else if index == 2 && ch == 'A' {
					parser.SetMethod(PHP_HTTP_MKACTIVITY)
				} else {
					parser.SetMethod(PHP_HTTP_NOT_IMPLEMENTED)
				}
			} else if index == 1 && parser.GetMethod() == PHP_HTTP_POST && ch == 'R' {
				parser.SetMethod(PHP_HTTP_PROPFIND)
			} else if index == 1 && parser.GetMethod() == PHP_HTTP_POST && ch == 'U' {
				parser.SetMethod(PHP_HTTP_PUT)
			} else if index == 1 && parser.GetMethod() == PHP_HTTP_POST && ch == 'A' {
				parser.SetMethod(PHP_HTTP_PATCH)
			} else if index == 1 && parser.GetMethod() == PHP_HTTP_SUBSCRIBE && ch == 'E' {
				parser.SetMethod(PHP_HTTP_SEARCH)
			} else if index == 2 && parser.GetMethod() == PHP_HTTP_UNLOCK && ch == 'S' {
				parser.SetMethod(PHP_HTTP_UNSUBSCRIBE)
			} else if index == 4 && parser.GetMethod() == PHP_HTTP_PROPFIND && ch == 'P' {
				parser.SetMethod(PHP_HTTP_PROPPATCH)
			} else {
				parser.SetMethod(PHP_HTTP_NOT_IMPLEMENTED)
			}
			index++
			break
		case SReqSpacesBeforeUrl:
			if ch == ' ' {
				break
			}
			if ch == '/' || ch == '*' {
				url_mark = p
				path_mark = p
				state = SReqPath
				break
			}
			c = LOWER(ch)
			if c >= 'a' && c <= 'z' {
				url_mark = p
				state = SReqSchema
				break
			}
			goto error
		case SReqSchema:
			c = LOWER(ch)
			if c >= 'a' && c <= 'z' {
				break
			}
			if ch == ':' {
				state = SReqSchemaSlash
				break
			} else if ch == '.' {
				state = SReqHost
				break
			} else if '0' <= ch && ch <= '9' {
				state = SReqHost
				break
			}
			goto error
		case SReqSchemaSlash:
			state = SReqSchemaSlashSlash
			break
		case SReqSchemaSlashSlash:
			state = SReqHost
			break
		case SReqHost:
			c = LOWER(ch)
			if c >= 'a' && c <= 'z' {
				break
			}
			if ch >= '0' && ch <= '9' || ch == '.' || ch == '-' {
				break
			}
			switch ch {
			case ':':
				state = SReqPort
				break
			case '/':
				path_mark = p
				state = SReqPath
				break
			case ' ':

				/* The request line looks like:
				 *   "GET http://foo.bar.com HTTP/1.1"
				 * That is, there is no path.
				 */

				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				state = SReqHttpStart
				break
			default:
				goto error
			}
			break
		case SReqPort:
			if ch >= '0' && ch <= '9' {
				break
			}
			switch ch {
			case '/':
				path_mark = p
				state = SReqPath
				break
			case ' ':

				/* The request line looks like:
				 *   "GET http://foo.bar.com:1234 HTTP/1.1"
				 * That is, there is no path.
				 */

				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				state = SReqHttpStart
				break
			default:
				goto error
			}
			break
		case SReqPath:
			if NormalUrlChar[uint8(ch)] != 0 {
				break
			}
			switch ch {
			case ' ':
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if path_mark != nil {
					if settings.GetOnPath() != nil {
						if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
							return p - data
						}
					}
				}
				path_mark = nil
				state = SReqHttpStart
				break
			case CR:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if path_mark != nil {
					if settings.GetOnPath() != nil {
						if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
							return p - data
						}
					}
				}
				path_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SReqLineAlmostDone
				break
			case LF:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if path_mark != nil {
					if settings.GetOnPath() != nil {
						if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
							return p - data
						}
					}
				}
				path_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SHeaderFieldStart
				break
			case '?':
				if path_mark != nil {
					if settings.GetOnPath() != nil {
						if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
							return p - data
						}
					}
				}
				path_mark = nil
				state = SReqQueryStringStart
				break
			case '#':
				if path_mark != nil {
					if settings.GetOnPath() != nil {
						if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
							return p - data
						}
					}
				}
				path_mark = nil
				state = SReqFragmentStart
				break
			default:
				goto error
			}
			break
		case SReqQueryStringStart:
			if NormalUrlChar[uint8(ch)] != 0 {
				query_string_mark = p
				state = SReqQueryString
				break
			}
			switch ch {
			case '?':
				break
			case ' ':
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				state = SReqHttpStart
				break
			case CR:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SReqLineAlmostDone
				break
			case LF:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SHeaderFieldStart
				break
			case '#':
				state = SReqFragmentStart
				break
			default:
				goto error
			}
			break
		case SReqQueryString:
			if NormalUrlChar[uint8(ch)] != 0 {
				break
			}
			switch ch {
			case '?':

				/* allow extra '?' in query string */

				break
			case ' ':
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if query_string_mark != nil {
					if settings.GetOnQueryString() != nil {
						if 0 != settings.GetOnQueryString()(parser, query_string_mark, p-query_string_mark) {
							return p - data
						}
					}
				}
				query_string_mark = nil
				state = SReqHttpStart
				break
			case CR:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if query_string_mark != nil {
					if settings.GetOnQueryString() != nil {
						if 0 != settings.GetOnQueryString()(parser, query_string_mark, p-query_string_mark) {
							return p - data
						}
					}
				}
				query_string_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SReqLineAlmostDone
				break
			case LF:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if query_string_mark != nil {
					if settings.GetOnQueryString() != nil {
						if 0 != settings.GetOnQueryString()(parser, query_string_mark, p-query_string_mark) {
							return p - data
						}
					}
				}
				query_string_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SHeaderFieldStart
				break
			case '#':
				if query_string_mark != nil {
					if settings.GetOnQueryString() != nil {
						if 0 != settings.GetOnQueryString()(parser, query_string_mark, p-query_string_mark) {
							return p - data
						}
					}
				}
				query_string_mark = nil
				state = SReqFragmentStart
				break
			default:
				goto error
			}
			break
		case SReqFragmentStart:
			if NormalUrlChar[uint8(ch)] != 0 {
				fragment_mark = p
				state = SReqFragment
				break
			}
			switch ch {
			case ' ':
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				state = SReqHttpStart
				break
			case CR:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SReqLineAlmostDone
				break
			case LF:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SHeaderFieldStart
				break
			case '?':
				fragment_mark = p
				state = SReqFragment
				break
			case '#':
				break
			default:
				goto error
			}
			break
		case SReqFragment:
			if NormalUrlChar[uint8(ch)] != 0 {
				break
			}
			switch ch {
			case ' ':
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if fragment_mark != nil {
					if settings.GetOnFragment() != nil {
						if 0 != settings.GetOnFragment()(parser, fragment_mark, p-fragment_mark) {
							return p - data
						}
					}
				}
				fragment_mark = nil
				state = SReqHttpStart
				break
			case CR:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if fragment_mark != nil {
					if settings.GetOnFragment() != nil {
						if 0 != settings.GetOnFragment()(parser, fragment_mark, p-fragment_mark) {
							return p - data
						}
					}
				}
				fragment_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SReqLineAlmostDone
				break
			case LF:
				if url_mark != nil {
					if settings.GetOnUrl() != nil {
						if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
							return p - data
						}
					}
				}
				url_mark = nil
				if fragment_mark != nil {
					if settings.GetOnFragment() != nil {
						if 0 != settings.GetOnFragment()(parser, fragment_mark, p-fragment_mark) {
							return p - data
						}
					}
				}
				fragment_mark = nil
				parser.SetHttpMajor(0)
				parser.SetHttpMinor(9)
				state = SHeaderFieldStart
				break
			case '?':

			case '#':
				break
			default:
				goto error
			}
			break
		case SReqHttpStart:
			switch ch {
			case 'H':
				state = s_req_http_H
				break
			case ' ':
				break
			default:
				goto error
			}
			break
		case s_req_http_H:
			state = s_req_http_HT
			break
		case s_req_http_HT:
			state = s_req_http_HTT
			break
		case s_req_http_HTT:
			state = s_req_http_HTTP
			break
		case s_req_http_HTTP:
			state = SReqFirstHttpMajor
			break
		case SReqFirstHttpMajor:
			if ch < '1' || ch > '9' {
				goto error
			}
			parser.SetHttpMajor(ch - '0')
			state = SReqHttpMajor
			break
		case SReqHttpMajor:
			if ch == '.' {
				state = SReqFirstHttpMinor
				break
			}
			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMajor(parser.GetHttpMajor() * 10)
			parser.SetHttpMajor(parser.GetHttpMajor() + ch - '0')
			if parser.GetHttpMajor() > 999 {
				goto error
			}
			break
		case SReqFirstHttpMinor:
			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMinor(ch - '0')
			state = SReqHttpMinor
			break
		case SReqHttpMinor:
			if ch == CR {
				state = SReqLineAlmostDone
				break
			}
			if ch == LF {
				state = SHeaderFieldStart
				break
			}

			/* XXX allow spaces after digit? */

			if ch < '0' || ch > '9' {
				goto error
			}
			parser.SetHttpMinor(parser.GetHttpMinor() * 10)
			parser.SetHttpMinor(parser.GetHttpMinor() + ch - '0')
			if parser.GetHttpMinor() > 999 {
				goto error
			}
			break
		case SReqLineAlmostDone:
			if ch != LF {
				goto error
			}
			state = SHeaderFieldStart
			break
		case SHeaderFieldStart:
			if ch == CR {
				state = SHeadersAlmostDone
				break
			}
			if ch == LF {

				/* they might be just sending \n instead of \r\n so this would be
				 * the second \n to denote the end of headers*/

				state = SHeadersAlmostDone
				goto headers_almost_done
			}
			c = TOKEN(ch)
			if !c {
				goto error
			}
			header_field_mark = p
			index = 0
			state = SHeaderField
			switch c {
			case 'c':
				header_state = h_C
				break
			case 'p':
				header_state = HMatchingProxyConnection
				break
			case 't':
				header_state = HMatchingTransferEncoding
				break
			case 'u':
				header_state = HMatchingUpgrade
				break
			default:
				header_state = HGeneral
				break
			}
			break
		case SHeaderField:
			c = TOKEN(ch)
			if c {
				switch header_state {
				case HGeneral:
					break
				case h_C:
					index++
					if c == 'o' {
						header_state = h_CO
					} else {
						header_state = HGeneral
					}
					break
				case h_CO:
					index++
					if c == 'n' {
						header_state = h_CON
					} else {
						header_state = HGeneral
					}
					break
				case h_CON:
					index++
					switch c {
					case 'n':
						header_state = HMatchingConnection
						break
					case 't':
						header_state = HMatchingContentLength
						break
					default:
						header_state = HGeneral
						break
					}
					break
				case HMatchingConnection:
					index++
					if index > b.SizeOf("CONNECTION")-1 || c != CONNECTION[index] {
						header_state = HGeneral
					} else if index == b.SizeOf("CONNECTION")-2 {
						header_state = HConnection
					}
					break
				case HMatchingProxyConnection:
					index++
					if index > b.SizeOf("PROXY_CONNECTION")-1 || c != PROXY_CONNECTION[index] {
						header_state = HGeneral
					} else if index == b.SizeOf("PROXY_CONNECTION")-2 {
						header_state = HConnection
					}
					break
				case HMatchingContentLength:
					index++
					if index > b.SizeOf("CONTENT_LENGTH")-1 || c != CONTENT_LENGTH[index] {
						header_state = HGeneral
					} else if index == b.SizeOf("CONTENT_LENGTH")-2 {
						header_state = HContentLength
					}
					break
				case HMatchingTransferEncoding:
					index++
					if index > b.SizeOf("TRANSFER_ENCODING")-1 || c != TRANSFER_ENCODING[index] {
						header_state = HGeneral
					} else if index == b.SizeOf("TRANSFER_ENCODING")-2 {
						header_state = HTransferEncoding
					}
					break
				case HMatchingUpgrade:
					index++
					if index > b.SizeOf("UPGRADE")-1 || c != UPGRADE[index] {
						header_state = HGeneral
					} else if index == b.SizeOf("UPGRADE")-2 {
						header_state = HUpgrade
					}
					break
				case HConnection:

				case HContentLength:

				case HTransferEncoding:

				case HUpgrade:
					if ch != ' ' {
						header_state = HGeneral
					}
					break
				default:
					r.Assert(false)
					break
				}
				break
			}
			if ch == ':' {
				if header_field_mark != nil {
					if settings.GetOnHeaderField() != nil {
						if 0 != settings.GetOnHeaderField()(parser, header_field_mark, p-header_field_mark) {
							return p - data
						}
					}
				}
				header_field_mark = nil
				state = SHeaderValueStart
				break
			}
			if ch == CR {
				state = SHeaderAlmostDone
				if header_field_mark != nil {
					if settings.GetOnHeaderField() != nil {
						if 0 != settings.GetOnHeaderField()(parser, header_field_mark, p-header_field_mark) {
							return p - data
						}
					}
				}
				header_field_mark = nil
				break
			}
			if ch == LF {
				if header_field_mark != nil {
					if settings.GetOnHeaderField() != nil {
						if 0 != settings.GetOnHeaderField()(parser, header_field_mark, p-header_field_mark) {
							return p - data
						}
					}
				}
				header_field_mark = nil
				state = SHeaderFieldStart
				break
			}
			goto error
		case SHeaderValueStart:
			if ch == ' ' {
				break
			}
			header_value_mark = p
			state = SHeaderValue
			index = 0
			c = LOWER(ch)
			if ch == CR {
				if header_value_mark != nil {
					if settings.GetOnHeaderValue() != nil {
						if 0 != settings.GetOnHeaderValue()(parser, header_value_mark, p-header_value_mark) {
							return p - data
						}
					}
				}
				header_value_mark = nil
				header_state = HGeneral
				state = SHeaderAlmostDone
				break
			}
			if ch == LF {
				if header_value_mark != nil {
					if settings.GetOnHeaderValue() != nil {
						if 0 != settings.GetOnHeaderValue()(parser, header_value_mark, p-header_value_mark) {
							return p - data
						}
					}
				}
				header_value_mark = nil
				state = SHeaderFieldStart
				break
			}
			switch header_state {
			case HUpgrade:
				parser.SetIsUpgrade(true)
				header_state = HGeneral
				break
			case HTransferEncoding:

				/* looking for 'Transfer-Encoding: chunked' */

				if 'c' == c {
					header_state = HMatchingTransferEncodingChunked
				} else {
					header_state = HGeneral
				}
				break
			case HContentLength:
				if ch < '0' || ch > '9' {
					goto error
				}
				parser.SetContentLength(ch - '0')
				break
			case HConnection:

				/* looking for 'Connection: keep-alive' */

				if c == 'k' {
					header_state = HMatchingConnectionKeepAlive
				} else if c == 'c' {
					header_state = HMatchingConnectionClose
				} else {
					header_state = HGeneral
				}
				break
			default:
				header_state = HGeneral
				break
			}
			break
		case SHeaderValue:
			c = LOWER(ch)
			if ch == CR {
				if header_value_mark != nil {
					if settings.GetOnHeaderValue() != nil {
						if 0 != settings.GetOnHeaderValue()(parser, header_value_mark, p-header_value_mark) {
							return p - data
						}
					}
				}
				header_value_mark = nil
				state = SHeaderAlmostDone
				break
			}
			if ch == LF {
				if header_value_mark != nil {
					if settings.GetOnHeaderValue() != nil {
						if 0 != settings.GetOnHeaderValue()(parser, header_value_mark, p-header_value_mark) {
							return p - data
						}
					}
				}
				header_value_mark = nil
				goto header_almost_done
			}
			switch header_state {
			case HGeneral:
				break
			case HConnection:

			case HTransferEncoding:
				r.Assert(false)
				break
			case HContentLength:
				if ch == ' ' {
					break
				}
				if ch < '0' || ch > '9' {
					goto error
				}
				parser.SetContentLength(parser.GetContentLength() * 10)
				parser.SetContentLength(parser.GetContentLength() + ch - '0')
				break
			case HMatchingTransferEncodingChunked:
				index++
				if index > b.SizeOf("CHUNKED")-1 || c != CHUNKED[index] {
					header_state = HGeneral
				} else if index == b.SizeOf("CHUNKED")-2 {
					header_state = HTransferEncodingChunked
				}
				break
			case HMatchingConnectionKeepAlive:
				index++
				if index > b.SizeOf("KEEP_ALIVE")-1 || c != KEEP_ALIVE[index] {
					header_state = HGeneral
				} else if index == b.SizeOf("KEEP_ALIVE")-2 {
					header_state = HConnectionKeepAlive
				}
				break
			case HMatchingConnectionClose:
				index++
				if index > b.SizeOf("CLOSE")-1 || c != CLOSE[index] {
					header_state = HGeneral
				} else if index == b.SizeOf("CLOSE")-2 {
					header_state = HConnectionClose
				}
				break
			case HTransferEncodingChunked:

			case HConnectionKeepAlive:

			case HConnectionClose:
				if ch != ' ' {
					header_state = HGeneral
				}
				break
			default:
				state = SHeaderValue
				header_state = HGeneral
				break
			}
			break
		case SHeaderAlmostDone:
		header_almost_done:
			state = SHeaderFieldStart
			switch header_state {
			case HConnectionKeepAlive:
				parser.SetIsConnectionKeepAlive(true)
				break
			case HConnectionClose:
				parser.SetIsConnectionClose(true)
				break
			case HTransferEncodingChunked:
				parser.SetIsChunked(true)
				break
			default:
				break
			}
			break
		case SHeadersAlmostDone:
		headers_almost_done:
			if parser.IsTrailing() {

				/* End of a chunked request */

				if settings.GetOnMessageComplete() != nil {
					if 0 != settings.GetOnMessageComplete()(parser) {
						return p - data
					}
				}
				state = NEW_MESSAGE()
				break
			}
			nread = 0
			if parser.IsUpgrade() || parser.GetMethod() == PHP_HTTP_CONNECT {
				parser.SetUpgrade(1)
			}

			/* Here we call the headers_complete callback. This is somewhat
			 * different than other callbacks because if the user returns 1, we
			 * will interpret that as saying that this message has no body. This
			 * is needed for the annoying case of receiving a response to a HEAD
			 * request.
			 */

			if settings.GetOnHeadersComplete() != nil {
				switch settings.GetOnHeadersComplete()(parser) {
				case 0:
					break
				case 1:
					parser.SetIsSkipbody(true)
					break
				default:
					return p - data
				}
			}

			/* We cannot meaningfully support upgrade requests, since we only
			 * support HTTP/1 for now.
			 */

			if parser.IsSkipbody() {
				if settings.GetOnMessageComplete() != nil {
					if 0 != settings.GetOnMessageComplete()(parser) {
						return p - data
					}
				}
				state = NEW_MESSAGE()
			} else if parser.IsChunked() {

				/* chunked encoding - ignore Content-Length header */

				state = SChunkSizeStart

				/* chunked encoding - ignore Content-Length header */

			} else {
				if parser.GetContentLength() == 0 {

					/* Content-Length header given but zero: Content-Length: 0\r\n */

					if settings.GetOnMessageComplete() != nil {
						if 0 != settings.GetOnMessageComplete()(parser) {
							return p - data
						}
					}
					state = NEW_MESSAGE()
				} else if parser.GetContentLength() > 0 {

					/* Content-Length header given and non-zero */

					state = SBodyIdentity

					/* Content-Length header given and non-zero */

				} else {
					if parser.GetType() == PHP_HTTP_REQUEST || PhpHttpShouldKeepAlive(parser) != 0 {

						/* Assume content-length 0 - read the next */

						if settings.GetOnMessageComplete() != nil {
							if 0 != settings.GetOnMessageComplete()(parser) {
								return p - data
							}
						}
						state = NEW_MESSAGE()
					} else {

						/* Read body until EOF */

						state = SBodyIdentityEof

						/* Read body until EOF */

					}
				}
			}
			break
		case SBodyIdentity:
			r.Assert(pe >= p)
			to_read = MIN(size_t(pe-p), int(parser.GetContentLength()))
			if to_read > 0 {
				if settings.GetOnBody() != nil {
					settings.GetOnBody()(parser, p, to_read)
				}
				p += to_read - 1
				parser.SetContentLength(parser.GetContentLength() - to_read)
				if parser.GetContentLength() == 0 {
					if settings.GetOnMessageComplete() != nil {
						if 0 != settings.GetOnMessageComplete()(parser) {
							return p - data
						}
					}
					state = NEW_MESSAGE()
				}
			}
			break
		case SBodyIdentityEof:
			to_read = pe - p
			if to_read > 0 {
				if settings.GetOnBody() != nil {
					settings.GetOnBody()(parser, p, to_read)
				}
				p += to_read - 1
			}
			break
		case SChunkSizeStart:
			r.Assert(parser.IsChunked())
			c = Unhex[uint8(ch)]
			if c == -1 {
				goto error
			}
			parser.SetContentLength(c)
			state = SChunkSize
			break
		case SChunkSize:
			r.Assert(parser.IsChunked())
			if ch == CR {
				state = SChunkSizeAlmostDone
				break
			}
			c = Unhex[uint8(ch)]
			if c == -1 {
				if ch == ';' || ch == ' ' {
					state = SChunkParameters
					break
				}
				goto error
			}
			parser.SetContentLength(parser.GetContentLength() * 16)
			parser.SetContentLength(parser.GetContentLength() + c)
			break
		case SChunkParameters:
			r.Assert(parser.IsChunked())

			/* just ignore this shit. TODO check for overflow */

			if ch == CR {
				state = SChunkSizeAlmostDone
				break
			}
			break
		case SChunkSizeAlmostDone:
			r.Assert(parser.IsChunked())
			if parser.GetContentLength() == 0 {
				parser.SetIsTrailing(true)
				state = SHeaderFieldStart
			} else {
				state = SChunkData
			}
			break
		case SChunkData:
			r.Assert(parser.IsChunked())
			r.Assert(pe >= p)
			to_read = MIN(size_t(pe-p), size_t(parser.GetContentLength()))
			if to_read > 0 {
				if settings.GetOnBody() != nil {
					settings.GetOnBody()(parser, p, to_read)
				}
				p += to_read - 1
			}
			if to_read == int(parser.GetContentLength()) {
				state = SChunkDataAlmostDone
			}
			parser.SetContentLength(parser.GetContentLength() - to_read)
			break
		case SChunkDataAlmostDone:
			r.Assert(parser.IsChunked())
			state = SChunkDataDone
			break
		case SChunkDataDone:
			r.Assert(parser.IsChunked())
			state = SChunkSizeStart
			break
		default:
			r.Assert(false)
			goto error
		}
	}
	if header_field_mark != nil {
		if settings.GetOnHeaderField() != nil {
			if 0 != settings.GetOnHeaderField()(parser, header_field_mark, p-header_field_mark) {
				return p - data
			}
		}
	}
	if header_value_mark != nil {
		if settings.GetOnHeaderValue() != nil {
			if 0 != settings.GetOnHeaderValue()(parser, header_value_mark, p-header_value_mark) {
				return p - data
			}
		}
	}
	if fragment_mark != nil {
		if settings.GetOnFragment() != nil {
			if 0 != settings.GetOnFragment()(parser, fragment_mark, p-fragment_mark) {
				return p - data
			}
		}
	}
	if query_string_mark != nil {
		if settings.GetOnQueryString() != nil {
			if 0 != settings.GetOnQueryString()(parser, query_string_mark, p-query_string_mark) {
				return p - data
			}
		}
	}
	if path_mark != nil {
		if settings.GetOnPath() != nil {
			if 0 != settings.GetOnPath()(parser, path_mark, p-path_mark) {
				return p - data
			}
		}
	}
	if url_mark != nil {
		if settings.GetOnUrl() != nil {
			if 0 != settings.GetOnUrl()(parser, url_mark, p-url_mark) {
				return p - data
			}
		}
	}
	parser.SetState(state)
	parser.SetHeaderState(header_state)
	parser.SetIndex(index)
	parser.SetNread(nread)
	return len_
error:
	parser.SetState(SDead)
	return p - data
}
func PhpHttpShouldKeepAlive(parser *PhpHttpParser) int {
	if parser.GetHttpMajor() > 0 && parser.GetHttpMinor() > 0 {

		/* HTTP/1.1 */

		if parser.IsConnectionClose() {
			return 0
		} else {
			return 1
		}

		/* HTTP/1.1 */

	} else {

		/* HTTP/1.0 or earlier */

		if parser.IsConnectionKeepAlive() {
			return 1
		} else {
			return 0
		}

		/* HTTP/1.0 or earlier */

	}
}
func PhpHttpMethodStr(m PhpHttpMethod) *byte { return MethodStrings[m] }
func PhpHttpParserInit(parser *PhpHttpParser, t PhpHttpParserType) {
	parser.SetType(t)
	if t == PHP_HTTP_REQUEST {
		parser.SetState(SStartReq)
	} else {
		if t == PHP_HTTP_RESPONSE {
			parser.SetState(SStartRes)
		} else {
			parser.SetState(SStartReqOrRes)
		}
	}
	parser.SetNread(0)
	parser.SetUpgrade(0)
	parser.SetFlags(0)
	parser.SetMethod(0)
}
