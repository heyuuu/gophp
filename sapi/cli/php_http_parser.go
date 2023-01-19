// <<generate>>

package cli

import g "sik/runtime/grammar"

// Source: <sapi/cli/php_http_parser.h>

/* Copyright 2009,2010 Ryan Dahl <ry@tinyclouds.org>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 */

// #define php_http_parser_h

// # include < sys / types . h >

// # include "php_config.h"

// # include "php_stdint.h"

/* Compile with -DPHP_HTTP_PARSER_STRICT=0 to make less checks, but run
 * faster
 */

// #define PHP_HTTP_PARSER_STRICT       1

/* Maximium header size allowed */

// #define PHP_HTTP_MAX_HEADER_SIZE       ( 80 * 1024 )

/* Callbacks should return non-zero to indicate an error. The parser will
 * then halt execution.
 *
 * The one exception is on_headers_complete. In a PHP_HTTP_RESPONSE parser
 * returning '1' from on_headers_complete will tell the parser that it
 * should not expect a body. This is used when receiving a response to a
 * HEAD request which may contain 'Content-Length' or 'Transfer-Encoding:
 * chunked' headers that indicate the presence of a body.
 *
 * http_data_cb does not return data chunks. It will be call arbitrarally
 * many times for each string. E.G. you might get 10 callbacks for "on_path"
 * each providing just a few characters more data.
 */

type PhpHttpDataCb func(_ *PhpHttpParser, at *byte, length int) int
type PhpHttpCb func(*PhpHttpParser) int

/* Request Methods */

type PhpHttpMethod = int

const (
	PHP_HTTP_DELETE = 0
	PHP_HTTP_GET
	PHP_HTTP_HEAD
	PHP_HTTP_POST
	PHP_HTTP_PUT
	PHP_HTTP_PATCH
	PHP_HTTP_CONNECT
	PHP_HTTP_OPTIONS
	PHP_HTTP_TRACE
	PHP_HTTP_COPY
	PHP_HTTP_LOCK
	PHP_HTTP_MKCOL
	PHP_HTTP_MOVE
	PHP_HTTP_MKCALENDAR
	PHP_HTTP_PROPFIND
	PHP_HTTP_PROPPATCH
	PHP_HTTP_SEARCH
	PHP_HTTP_UNLOCK
	PHP_HTTP_REPORT
	PHP_HTTP_MKACTIVITY
	PHP_HTTP_CHECKOUT
	PHP_HTTP_MERGE
	PHP_HTTP_MSEARCH
	PHP_HTTP_NOTIFY
	PHP_HTTP_SUBSCRIBE
	PHP_HTTP_UNSUBSCRIBE
	PHP_HTTP_NOT_IMPLEMENTED
)

type PhpHttpParserType = int

const (
	PHP_HTTP_REQUEST = iota
	PHP_HTTP_RESPONSE
	PHP_HTTP_BOTH
)

type State = int

const (
	SDead = 1
	SStartReqOrRes
	s_res_or_resp_H
	SStartRes
	s_res_H
	s_res_HT
	s_res_HTT
	s_res_HTTP
	SResFirstHttpMajor
	SResHttpMajor
	SResFirstHttpMinor
	SResHttpMinor
	SResFirstStatusCode
	SResStatusCode
	SResStatus
	SResLineAlmostDone
	SStartReq
	SReqMethod
	SReqSpacesBeforeUrl
	SReqSchema
	SReqSchemaSlash
	SReqSchemaSlashSlash
	SReqHost
	SReqPort
	SReqPath
	SReqQueryStringStart
	SReqQueryString
	SReqFragmentStart
	SReqFragment
	SReqHttpStart
	s_req_http_H
	s_req_http_HT
	s_req_http_HTT
	s_req_http_HTTP
	SReqFirstHttpMajor
	SReqHttpMajor
	SReqFirstHttpMinor
	SReqHttpMinor
	SReqLineAlmostDone
	SHeaderFieldStart
	SHeaderField
	SHeaderValueStart
	SHeaderValue
	SHeaderAlmostDone
	SHeadersAlmostDone
	SChunkSizeStart
	SChunkSize
	SChunkSizeAlmostDone
	SChunkParameters
	SChunkData
	SChunkDataAlmostDone
	SChunkDataDone
	SBodyIdentity
	SBodyIdentityEof
)

// @type PhpHttpParser struct
// @type PhpHttpParserSettings struct

/* If php_http_should_keep_alive() in the on_headers_complete or
 * on_message_complete callback returns true, then this will be should be
 * the last message on the connection.
 * If you are the server, respond with the "Connection: close" header.
 * If you are the client, close the connection.
 */

/* Returns a string version of the HTTP method. */

// Source: <sapi/cli/php_http_parser.c>

/* Copyright 2009,2010 Ryan Dahl <ry@tinyclouds.org>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to
 * deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
 * sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 */

// # include < assert . h >

// # include < stddef . h >

// # include "php_http_parser.h"

// #define MIN(a,b) ( ( a ) < ( b ) ? ( a ) : ( b ) )

// #define CALLBACK2(FOR) do { if ( settings -> on_ ## FOR ) { if ( 0 != settings -> on_ ## FOR ( parser ) ) return ( p - data ) ; } } while ( 0 )

// #define MARK(FOR) do { FOR ## _mark = p ; } while ( 0 )

// #define CALLBACK_NOCLEAR(FOR) do { if ( FOR ## _mark ) { if ( settings -> on_ ## FOR ) { if ( 0 != settings -> on_ ## FOR ( parser , FOR ## _mark , p - FOR ## _mark ) ) { return ( p - data ) ; } } } } while ( 0 )

// #define CALLBACK(FOR) do { CALLBACK_NOCLEAR ( FOR ) ; FOR ## _mark = NULL ; } while ( 0 )

// #define PROXY_CONNECTION       "proxy-connection"

// #define CONNECTION       "connection"

// #define CONTENT_LENGTH       "content-length"

// #define TRANSFER_ENCODING       "transfer-encoding"

// #define UPGRADE       "upgrade"

// #define CHUNKED       "chunked"

// #define KEEP_ALIVE       "keep-alive"

// #define CLOSE       "close"

var MethodStrings []*byte = []*byte{"DELETE", "GET", "HEAD", "POST", "PUT", "PATCH", "CONNECT", "OPTIONS", "TRACE", "COPY", "LOCK", "MKCOL", "MOVE", "MKCALENDAR", "PROPFIND", "PROPPATCH", "SEARCH", "UNLOCK", "REPORT", "MKACTIVITY", "CHECKOUT", "MERGE", "M-SEARCH", "NOTIFY", "SUBSCRIBE", "UNSUBSCRIBE", "NOTIMPLEMENTED"}

/* Tokens as defined by rfc 2616. Also lowercases them.
 *        token       = 1*<any CHAR except CTLs or separators>
 *     separators     = "(" | ")" | "<" | ">" | "@"
 *                    | "," | ";" | ":" | "\" | <">
 *                    | "/" | "[" | "]" | "?" | "="
 *                    | "{" | "}" | SP | HT
 */

var Tokens []byte = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, ' ', '!', '"', '#', '$', '%', '&', '\'', 0, 0, '*', '+', 0, '-', '.', '/', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 0, 0, 0, 0, 0, 0, 0, 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 0, 0, 0, '^', '_', '`', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 0, '|', '}', '~', 0}
var Unhex []int8_t = []int8_t{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, -1, -1, -1, -1, -1, -1, -1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 10, 11, 12, 13, 14, 15, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
var NormalUrlChar []uint8 = []uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0}

// #define PARSING_HEADER(state) ( state <= s_headers_almost_done && 0 == ( parser -> flags & F_TRAILING ) )

type HeaderStates = int

const (
	HGeneral = 0
	h_C
	h_CO
	h_CON
	HMatchingConnection
	HMatchingProxyConnection
	HMatchingContentLength
	HMatchingTransferEncoding
	HMatchingUpgrade
	HConnection
	HContentLength
	HTransferEncoding
	HUpgrade
	HMatchingTransferEncodingChunked
	HMatchingConnectionKeepAlive
	HMatchingConnectionClose
	HTransferEncodingChunked
	HConnectionKeepAlive
	HConnectionClose
)

type Flags = int

const (
	F_CHUNKED               Flags = 1 << 0
	F_CONNECTION_KEEP_ALIVE Flags = 1 << 1
	F_CONNECTION_CLOSE      Flags = 1 << 2
	F_TRAILING              Flags = 1 << 3
	F_UPGRADE               Flags = 1 << 4
	F_SKIPBODY              Flags = 1 << 5
)

// #define CR       '\r'

// #define LF       '\n'

// #define LOWER(c) ( unsigned char ) ( c | 0x20 )

// #define TOKEN(c) tokens [ ( unsigned char ) c ]

// #define start_state       ( parser -> type == PHP_HTTP_REQUEST ? s_start_req : s_start_res )

// #define STRICT_CHECK(cond)

// #define NEW_MESSAGE() start_state

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

	/* technically we could combine all of these (except for url_mark) into one
	   variable, saving stack space, but it seems more clear to have them
	   separated. */

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
		if state <= SHeadersAlmostDone && 0 == (parser.GetFlags()&F_TRAILING) {
			nread++

			/* Buffer overflow attack */

			if nread > 80*1024 {
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
			if ch == '\r' || ch == '\n' {
				break
			}
			parser.SetFlags(0)
			parser.SetContentLength(-1)
			if settings.GetOnMessageBegin() != nil {
				if 0 != settings.GetOnMessageBegin()(parser) {
					return p - data
				}
			}
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
			if settings.GetOnMessageBegin() != nil {
				if 0 != settings.GetOnMessageBegin()(parser) {
					return p - data
				}
			}
			switch ch {
			case 'H':
				state = s_res_H
				break
			case '\r':

			case '\n':
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
				case '\r':
					state = SResLineAlmostDone
					break
				case '\n':
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

			if ch == '\r' {
				state = SResLineAlmostDone
				break
			}
			if ch == '\n' {
				state = SHeaderFieldStart
				break
			}
			break
		case SResLineAlmostDone:
			state = SHeaderFieldStart
			break
		case SStartReq:
			if ch == '\r' || ch == '\n' {
				break
			}
			parser.SetFlags(0)
			parser.SetContentLength(-1)
			if settings.GetOnMessageBegin() != nil {
				if 0 != settings.GetOnMessageBegin()(parser) {
					return p - data
				}
			}
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
			c = uint8(ch | 0x20)
			if c >= 'a' && c <= 'z' {
				url_mark = p
				state = SReqSchema
				break
			}
			goto error
		case SReqSchema:
			c = uint8(ch | 0x20)
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
			c = uint8(ch | 0x20)
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
			case '\r':
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
			case '\n':
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
			case '\r':
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
			case '\n':
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
			case '\r':
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
			case '\n':
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
			case '\r':
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
			case '\n':
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
			case '\r':
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
			case '\n':
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
			if ch == '\r' {
				state = SReqLineAlmostDone
				break
			}
			if ch == '\n' {
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
			if ch != '\n' {
				goto error
			}
			state = SHeaderFieldStart
			break
		case SHeaderFieldStart:
			if ch == '\r' {
				state = SHeadersAlmostDone
				break
			}
			if ch == '\n' {

				/* they might be just sending \n instead of \r\n so this would be
				 * the second \n to denote the end of headers*/

				state = SHeadersAlmostDone
				goto headers_almost_done
			}
			c = Tokens[uint8(ch)]
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
			c = Tokens[uint8(ch)]
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
					if index > g.SizeOf("CONNECTION")-1 || c != "connection"[index] {
						header_state = HGeneral
					} else if index == g.SizeOf("CONNECTION")-2 {
						header_state = HConnection
					}
					break
				case HMatchingProxyConnection:
					index++
					if index > g.SizeOf("PROXY_CONNECTION")-1 || c != "proxy-connection"[index] {
						header_state = HGeneral
					} else if index == g.SizeOf("PROXY_CONNECTION")-2 {
						header_state = HConnection
					}
					break
				case HMatchingContentLength:
					index++
					if index > g.SizeOf("CONTENT_LENGTH")-1 || c != "content-length"[index] {
						header_state = HGeneral
					} else if index == g.SizeOf("CONTENT_LENGTH")-2 {
						header_state = HContentLength
					}
					break
				case HMatchingTransferEncoding:
					index++
					if index > g.SizeOf("TRANSFER_ENCODING")-1 || c != "transfer-encoding"[index] {
						header_state = HGeneral
					} else if index == g.SizeOf("TRANSFER_ENCODING")-2 {
						header_state = HTransferEncoding
					}
					break
				case HMatchingUpgrade:
					index++
					if index > g.SizeOf("UPGRADE")-1 || c != "upgrade"[index] {
						header_state = HGeneral
					} else if index == g.SizeOf("UPGRADE")-2 {
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
					assert(false)
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
			if ch == '\r' {
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
			if ch == '\n' {
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
			c = uint8(ch | 0x20)
			if ch == '\r' {
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
			if ch == '\n' {
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
				parser.SetFlags(parser.GetFlags() | F_UPGRADE)
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
			c = uint8(ch | 0x20)
			if ch == '\r' {
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
			if ch == '\n' {
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
				assert(false)
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
				if index > g.SizeOf("CHUNKED")-1 || c != "chunked"[index] {
					header_state = HGeneral
				} else if index == g.SizeOf("CHUNKED")-2 {
					header_state = HTransferEncodingChunked
				}
				break
			case HMatchingConnectionKeepAlive:
				index++
				if index > g.SizeOf("KEEP_ALIVE")-1 || c != "keep-alive"[index] {
					header_state = HGeneral
				} else if index == g.SizeOf("KEEP_ALIVE")-2 {
					header_state = HConnectionKeepAlive
				}
				break
			case HMatchingConnectionClose:
				index++
				if index > g.SizeOf("CLOSE")-1 || c != "close"[index] {
					header_state = HGeneral
				} else if index == g.SizeOf("CLOSE")-2 {
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
				parser.SetFlags(parser.GetFlags() | F_CONNECTION_KEEP_ALIVE)
				break
			case HConnectionClose:
				parser.SetFlags(parser.GetFlags() | F_CONNECTION_CLOSE)
				break
			case HTransferEncodingChunked:
				parser.SetFlags(parser.GetFlags() | F_CHUNKED)
				break
			default:
				break
			}
			break
		case SHeadersAlmostDone:
		headers_almost_done:
			if (parser.GetFlags() & F_TRAILING) != 0 {

				/* End of a chunked request */

				if settings.GetOnMessageComplete() != nil {
					if 0 != settings.GetOnMessageComplete()(parser) {
						return p - data
					}
				}
				if parser.GetType() == PHP_HTTP_REQUEST {
					state = SStartReq
				} else {
					state = SStartRes
				}
				break
			}
			nread = 0
			if (parser.GetFlags()&F_UPGRADE) != 0 || parser.GetMethod() == PHP_HTTP_CONNECT {
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
					parser.SetFlags(parser.GetFlags() | F_SKIPBODY)
					break
				default:
					return p - data
				}
			}

			/* We cannot meaningfully support upgrade requests, since we only
			 * support HTTP/1 for now.
			 */

			if (parser.GetFlags() & F_SKIPBODY) != 0 {
				if settings.GetOnMessageComplete() != nil {
					if 0 != settings.GetOnMessageComplete()(parser) {
						return p - data
					}
				}
				if parser.GetType() == PHP_HTTP_REQUEST {
					state = SStartReq
				} else {
					state = SStartRes
				}
			} else if (parser.GetFlags() & F_CHUNKED) != 0 {

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
					if parser.GetType() == PHP_HTTP_REQUEST {
						state = SStartReq
					} else {
						state = SStartRes
					}
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
						if parser.GetType() == PHP_HTTP_REQUEST {
							state = SStartReq
						} else {
							state = SStartRes
						}
					} else {

						/* Read body until EOF */

						state = SBodyIdentityEof

						/* Read body until EOF */

					}
				}
			}
			break
		case SBodyIdentity:
			assert(pe >= p)
			if size_t(pe-p) < int(parser.GetContentLength()) {
				to_read = size_t(pe - p)
			} else {
				to_read = int(parser.GetContentLength())
			}
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
					if parser.GetType() == PHP_HTTP_REQUEST {
						state = SStartReq
					} else {
						state = SStartRes
					}
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
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			c = Unhex[uint8(ch)]
			if c == -1 {
				goto error
			}
			parser.SetContentLength(c)
			state = SChunkSize
			break
		case SChunkSize:
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			if ch == '\r' {
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
			assert((parser.GetFlags() & F_CHUNKED) != 0)

			/* just ignore this shit. TODO check for overflow */

			if ch == '\r' {
				state = SChunkSizeAlmostDone
				break
			}
			break
		case SChunkSizeAlmostDone:
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			if parser.GetContentLength() == 0 {
				parser.SetFlags(parser.GetFlags() | F_TRAILING)
				state = SHeaderFieldStart
			} else {
				state = SChunkData
			}
			break
		case SChunkData:
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			assert(pe >= p)
			if size_t(pe-p) < size_t(parser.GetContentLength()) {
				to_read = size_t(pe - p)
			} else {
				to_read = size_t(parser.GetContentLength())
			}
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
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			state = SChunkDataDone
			break
		case SChunkDataDone:
			assert((parser.GetFlags() & F_CHUNKED) != 0)
			state = SChunkSizeStart
			break
		default:
			assert(false)
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

		if (parser.GetFlags() & F_CONNECTION_CLOSE) != 0 {
			return 0
		} else {
			return 1
		}

		/* HTTP/1.1 */

	} else {

		/* HTTP/1.0 or earlier */

		if (parser.GetFlags() & F_CONNECTION_KEEP_ALIVE) != 0 {
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
