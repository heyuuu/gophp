// <<generate>>

package zend

// Source: <Zend/zend_smart_str.h>

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
   | Author: Sascha Schumann <sascha@schumann.cx>                         |
   +----------------------------------------------------------------------+
*/

// Source: <Zend/zend_smart_str.c>

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
   | Author: Dmitry Stogov <dmitry@php.net>                               |
   +----------------------------------------------------------------------+
*/

const SMART_STR_OVERHEAD = ZEND_MM_OVERHEAD + _ZSTR_HEADER_SIZE + 1
const SMART_STR_START_SIZE = 256
const SMART_STR_START_LEN = SMART_STR_START_SIZE - SMART_STR_OVERHEAD
const SMART_STR_PAGE = 4096

/* Windows uses VK_ESCAPE instead of \e */

const VK_ESCAPE = 'e'
const SMART_STRING_OVERHEAD = ZEND_MM_OVERHEAD + 1
const SMART_STRING_START_SIZE = 256
const SMART_STRING_START_LEN = SMART_STRING_START_SIZE - SMART_STRING_OVERHEAD
const SMART_STRING_PAGE = 4096
