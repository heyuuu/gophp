// <<generate>>

package zend

// Source: <Zend/zend_errors.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

const E_ERROR = 1 << 0
const E_WARNING = 1 << 1
const E_PARSE ZendLong = 1 << 2
const E_NOTICE = 1 << 3
const E_CORE_ERROR = 1 << 4
const E_CORE_WARNING = 1 << 5
const E_COMPILE_ERROR = 1 << 6
const E_COMPILE_WARNING = 1 << 7
const E_USER_ERROR ZendLong = 1 << 8
const E_USER_WARNING ZendLong = 1 << 9
const E_USER_NOTICE ZendLong = 1 << 10
const E_STRICT ZendLong = 1 << 11
const E_RECOVERABLE_ERROR ZendLong = 1 << 12
const E_DEPRECATED = 1 << 13
const E_USER_DEPRECATED ZendLong = 1 << 14
const E_ALL ZendLong = E_ERROR | E_WARNING | E_PARSE | E_NOTICE | E_CORE_ERROR | E_CORE_WARNING | E_COMPILE_ERROR | E_COMPILE_WARNING | E_USER_ERROR | E_USER_WARNING | E_USER_NOTICE | E_RECOVERABLE_ERROR | E_DEPRECATED | E_USER_DEPRECATED | E_STRICT
const E_CORE = E_CORE_ERROR | E_CORE_WARNING
