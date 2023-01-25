// <<generate>>

package standard

import (
	"sik/core"
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

const HTTP_HEADER_BLOCK_SIZE = 1024
const PHP_URL_REDIRECT_MAX = 20
const HTTP_HEADER_USER_AGENT = 1
const HTTP_HEADER_HOST = 2
const HTTP_HEADER_AUTH = 4
const HTTP_HEADER_FROM = 8
const HTTP_HEADER_CONTENT_LENGTH = 16
const HTTP_HEADER_TYPE = 32
const HTTP_HEADER_CONNECTION = 64
const HTTP_WRAPPER_HEADER_INIT = 1
const HTTP_WRAPPER_REDIRECTED = 2

var HttpStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapHttp, nil, PhpStreamHttpStreamStat, nil, nil, "http", nil, nil, nil, nil, nil}
var PhpStreamHttpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&HttpStreamWops, nil, 1}
