package zend

// Source: <Zend/zend_string.h>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

type ZendStringCopyStorageFuncT func()
type ZendNewInternedStringFuncT func(str *ZendString) *ZendString
type ZendStringInitInternedFuncT func(str *byte, size int, permanent int) *ZendString

var ZendNewInternedString ZendNewInternedStringFuncT
var ZendStringInitInterned ZendStringInitInternedFuncT
var ZendOneCharString []*ZendString

/* Shortcuts */

/* Compatibility macros */

const _ZSTR_HEADER_SIZE = zend_long((*byte)(&((*ZendString)(nil).GetVal())) - (*byte)(nil))
const ZSTR_MAX_OVERHEAD = ZEND_MM_ALIGNED_SIZE(_ZSTR_HEADER_SIZE + 1)
const ZSTR_MAX_LEN = SIZE_MAX - ZSTR_MAX_OVERHEAD

/*---*/

/*
 * DJBX33A (Daniel J. Bernstein, Times 33 with Addition)
 *
 * This is Daniel J. Bernstein's popular `times 33' hash function as
 * posted by him years ago on comp.lang.c. It basically uses a function
 * like ``hash(i) = hash(i-1) * 33 + str[i]''. This is one of the best
 * known hash functions for strings. Because it is both computed very
 * fast and distributes very well.
 *
 * The magic of number 33, i.e. why it works better than many other
 * constants, prime or not, has never been adequately explained by
 * anyone. So I try an explanation: if one experimentally tests all
 * multipliers between 1 and 256 (as RSE did now) one detects that even
 * numbers are not usable at all. The remaining 128 odd numbers
 * (except for the number 1) work more or less all equally well. They
 * all distribute in an acceptable way and this way fill a hash table
 * with an average percent of approx. 86%.
 *
 * If one compares the Chi^2 values of the variants, the number 33 not
 * even has the best value. But the number 33 and a few other equally
 * good numbers like 17, 31, 63, 127 and 129 have nevertheless a great
 * advantage to the remaining numbers in the large set of possible
 * multipliers: their multiply operation can be replaced by a faster
 * operation based on just one shift plus either a single addition
 * or subtraction operation. And because a hash function has to both
 * distribute good _and_ has to be very fast to compute, those few
 * numbers should be preferred and seems to be the reason why Daniel J.
 * Bernstein also preferred it.
 *
 *
 *                  -- Ralf S. Engelschall <rse@engelschall.com>
 */

type ZendKnownStringId = int

const (
	ZEND_STR_FILE = iota
	ZEND_STR_LINE
	ZEND_STR_FUNCTION
	ZEND_STR_CLASS
	ZEND_STR_OBJECT
	ZEND_STR_TYPE
	ZEND_STR_OBJECT_OPERATOR
	ZEND_STR_PAAMAYIM_NEKUDOTAYIM
	ZEND_STR_ARGS
	ZEND_STR_UNKNOWN
	ZEND_STR_EVAL
	ZEND_STR_INCLUDE
	ZEND_STR_REQUIRE
	ZEND_STR_INCLUDE_ONCE
	ZEND_STR_REQUIRE_ONCE
	ZEND_STR_SCALAR
	ZEND_STR_ERROR_REPORTING
	ZEND_STR_STATIC
	ZEND_STR_THIS
	ZEND_STR_VALUE
	ZEND_STR_KEY
	ZEND_STR_MAGIC_AUTOLOAD
	ZEND_STR_MAGIC_INVOKE
	ZEND_STR_PREVIOUS
	ZEND_STR_CODE
	ZEND_STR_MESSAGE
	ZEND_STR_SEVERITY
	ZEND_STR_STRING
	ZEND_STR_TRACE
	ZEND_STR_SCHEME
	ZEND_STR_HOST
	ZEND_STR_PORT
	ZEND_STR_USER
	ZEND_STR_PASS
	ZEND_STR_PATH
	ZEND_STR_QUERY
	ZEND_STR_FRAGMENT
	ZEND_STR_NULL
	ZEND_STR_BOOLEAN
	ZEND_STR_INTEGER
	ZEND_STR_DOUBLE
	ZEND_STR_ARRAY
	ZEND_STR_RESOURCE
	ZEND_STR_CLOSED_RESOURCE
	ZEND_STR_NAME
	ZEND_STR_ARGV
	ZEND_STR_ARGC
	ZEND_STR_ARRAY_CAPITALIZED
	ZEND_STR_LAST_KNOWN
)

// Source: <Zend/zend_string.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* Any strings interned in the startup phase. Common to all the threads,
   won't be free'd until process exit. If we want an ability to
   add permanent strings even after startup, it would be still
   possible on costs of locking in the thread safe builds. */

var InternedStringsPermanent HashTable
var InternedStringRequestHandler ZendNewInternedStringFuncT = ZendNewInternedStringRequest
var InternedStringInitRequestHandler ZendStringInitInternedFuncT = ZendStringInitInternedRequest
var ZendEmptyString *ZendString = nil
var ZendKnownStrings **ZendString = nil
var KnownStrings []*byte = []*byte{"file", "line", "function", "class", "object", "type", "->", "::", "args", "unknown", "eval", "include", "require", "include_once", "require_once", "scalar", "error_reporting", "static", "this", "value", "key", "__autoload", "__invoke", "previous", "code", "message", "severity", "string", "trace", "scheme", "host", "port", "user", "pass", "path", "query", "fragment", "NULL", "boolean", "integer", "double", "array", "resource", "resource (closed)", "name", "argv", "argc", "Array", nil}

/* This function might be not thread safe at least because it would update the
   hash val in the passed string. Be sure it is called in the appropriate context. */
