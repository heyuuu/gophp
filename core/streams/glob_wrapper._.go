// <<generate>>

package streams

import (
	"sik/core"
)

// Source: <main/streams/glob_wrapper.c>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

const GLOB_ONLYDIR = 1 << 30
const GLOB_FLAGMASK = ^GLOB_ONLYDIR

/* {{{ */

var PhpGlobStreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpGlobStreamRead, PhpGlobStreamClose, nil, "glob", PhpGlobStreamRewind, nil, nil, nil}

/* {{{ php_glob_stream_opener */

var PhpGlobStreamWrapperOps core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{nil, nil, nil, nil, PhpGlobStreamOpener, "glob", nil, nil, nil, nil, nil}
var PhpGlobStreamWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpGlobStreamWrapperOps, nil, 0}
