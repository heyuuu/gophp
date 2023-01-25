// <<generate>>

package streams

import (
	"sik/core"
)

// Source: <main/streams/memory.c>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

var PhpUrlDecode func(str *byte, len_ int) int

/* Memory streams use a dynamic memory buffer to emulate a stream.
 * You can use php_stream_memory_open to create a readonly stream
 * from an existing memory buffer.
 */

/* {{{ */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

var PhpStreamMemoryOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamMemoryWrite, PhpStreamMemoryRead, PhpStreamMemoryClose, PhpStreamMemoryFlush, "MEMORY", PhpStreamMemorySeek, PhpStreamMemoryCast, PhpStreamMemoryStat, PhpStreamMemorySetOption}

/* {{{ */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

var PhpStreamTempOps core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "TEMP", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}

/* }}} */

/* }}} */

/* }}} */

/* }}} */

var PhpStreamRfc2397Ops core.PhpStreamOps = core.PhpStreamOps{PhpStreamTempWrite, PhpStreamTempRead, PhpStreamTempClose, PhpStreamTempFlush, "RFC2397", PhpStreamTempSeek, PhpStreamTempCast, PhpStreamTempStat, PhpStreamTempSetOption}
var PhpStreamRfc2397Wops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapRfc2397, nil, nil, nil, nil, "RFC2397", nil, nil, nil, nil, nil}
var PhpStreamRfc2397Wrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpStreamRfc2397Wops, nil, 1}
