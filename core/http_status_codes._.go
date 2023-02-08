// <<generate>>

package core

import (
	b "sik/builtin"
)

// Source: <main/http_status_codes.h>

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
   | Author: Andrea Faulds     <ajf@ajf.me>                               |
   +----------------------------------------------------------------------+
*/

var HttpStatusMap []HttpResponseStatusCodePair = []HttpResponseStatusCodePair{
	MakeHttpResponseStatusCodePair(100, "Continue"),
	MakeHttpResponseStatusCodePair(101, "Switching Protocols"),
	MakeHttpResponseStatusCodePair(200, "OK"),
	MakeHttpResponseStatusCodePair(201, "Created"),
	MakeHttpResponseStatusCodePair(202, "Accepted"),
	MakeHttpResponseStatusCodePair(203, "Non-Authoritative Information"),
	MakeHttpResponseStatusCodePair(204, "No Content"),
	MakeHttpResponseStatusCodePair(205, "Reset Content"),
	MakeHttpResponseStatusCodePair(206, "Partial Content"),
	MakeHttpResponseStatusCodePair(300, "Multiple Choices"),
	MakeHttpResponseStatusCodePair(301, "Moved Permanently"),
	MakeHttpResponseStatusCodePair(302, "Found"),
	MakeHttpResponseStatusCodePair(303, "See Other"),
	MakeHttpResponseStatusCodePair(304, "Not Modified"),
	MakeHttpResponseStatusCodePair(305, "Use Proxy"),
	MakeHttpResponseStatusCodePair(307, "Temporary Redirect"),
	MakeHttpResponseStatusCodePair(308, "Permanent Redirect"),
	MakeHttpResponseStatusCodePair(400, "Bad Request"),
	MakeHttpResponseStatusCodePair(401, "Unauthorized"),
	MakeHttpResponseStatusCodePair(402, "Payment Required"),
	MakeHttpResponseStatusCodePair(403, "Forbidden"),
	MakeHttpResponseStatusCodePair(404, "Not Found"),
	MakeHttpResponseStatusCodePair(405, "Method Not Allowed"),
	MakeHttpResponseStatusCodePair(406, "Not Acceptable"),
	MakeHttpResponseStatusCodePair(407, "Proxy Authentication Required"),
	MakeHttpResponseStatusCodePair(408, "Request Timeout"),
	MakeHttpResponseStatusCodePair(409, "Conflict"),
	MakeHttpResponseStatusCodePair(410, "Gone"),
	MakeHttpResponseStatusCodePair(411, "Length Required"),
	MakeHttpResponseStatusCodePair(412, "Precondition Failed"),
	MakeHttpResponseStatusCodePair(413, "Request Entity Too Large"),
	MakeHttpResponseStatusCodePair(414, "Request-URI Too Long"),
	MakeHttpResponseStatusCodePair(415, "Unsupported Media Type"),
	MakeHttpResponseStatusCodePair(416, "Requested Range Not Satisfiable"),
	MakeHttpResponseStatusCodePair(417, "Expectation Failed"),
	MakeHttpResponseStatusCodePair(426, "Upgrade Required"),
	MakeHttpResponseStatusCodePair(428, "Precondition Required"),
	MakeHttpResponseStatusCodePair(429, "Too Many Requests"),
	MakeHttpResponseStatusCodePair(431, "Request Header Fields Too Large"),
	MakeHttpResponseStatusCodePair(451, "Unavailable For Legal Reasons"),
	MakeHttpResponseStatusCodePair(500, "Internal Server Error"),
	MakeHttpResponseStatusCodePair(501, "Not Implemented"),
	MakeHttpResponseStatusCodePair(502, "Bad Gateway"),
	MakeHttpResponseStatusCodePair(503, "Service Unavailable"),
	MakeHttpResponseStatusCodePair(504, "Gateway Timeout"),
	MakeHttpResponseStatusCodePair(505, "HTTP Version Not Supported"),
	MakeHttpResponseStatusCodePair(506, "Variant Also Negotiates"),
	MakeHttpResponseStatusCodePair(511, "Network Authentication Required"),
	MakeHttpResponseStatusCodePair(0, nil),
}
var HttpStatusMapLen int = b.SizeOf("http_status_map")/b.SizeOf("http_response_status_code_pair") - 1
