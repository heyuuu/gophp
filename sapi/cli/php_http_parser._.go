// <<generate>>

package cli

import (
	b "sik/builtin"
)

const PHP_HTTP_MAX_HEADER_SIZE = 80 * 1024

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

const PROXY_CONNECTION = "proxy-connection"
const CONNECTION = "connection"
const CONTENT_LENGTH = "content-length"
const TRANSFER_ENCODING = "transfer-encoding"
const UPGRADE = "upgrade"
const CHUNKED = "chunked"
const KEEP_ALIVE = "keep-alive"
const CLOSE = "close"

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
const CR = '\r'
const LF = '\n'
const StartState = b.Cond(parser.type_ == PHP_HTTP_REQUEST, SStartReq, SStartRes)
