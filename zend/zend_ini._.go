// <<generate>>

package zend

import "sik/zend/types"

// Source: <Zend/zend_ini.h>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

const ZEND_INI_USER = 1 << 0
const ZEND_INI_PERDIR = 1 << 1
const ZEND_INI_SYSTEM = 1 << 2
const ZEND_INI_ALL = ZEND_INI_USER | ZEND_INI_PERDIR | ZEND_INI_SYSTEM

var DisplayIniEntries func(module *ZendModuleEntry)

/* Standard message handlers */

const ZEND_INI_DISPLAY_ORIG = 1
const ZEND_INI_DISPLAY_ACTIVE = 2
const ZEND_INI_STAGE_STARTUP = 1 << 0
const ZEND_INI_STAGE_SHUTDOWN = 1 << 1
const ZEND_INI_STAGE_ACTIVATE = 1 << 2
const ZEND_INI_STAGE_DEACTIVATE = 1 << 3
const ZEND_INI_STAGE_RUNTIME = 1 << 4
const ZEND_INI_STAGE_HTACCESS = 1 << 5
const ZEND_INI_STAGE_IN_REQUEST = ZEND_INI_STAGE_ACTIVATE | ZEND_INI_STAGE_DEACTIVATE | ZEND_INI_STAGE_RUNTIME | ZEND_INI_STAGE_HTACCESS

/* INI parsing engine */

type ZendIniParserCbT func(arg1 *types.Zval, arg2 *types.Zval, arg3 *types.Zval, callback_type int, arg any)

/* INI entries */

const ZEND_INI_PARSER_ENTRY = 1
const ZEND_INI_PARSER_SECTION = 2
const ZEND_INI_PARSER_POP_ENTRY = 3

var RegisteredZendIniDirectives *types.HashTable

const NO_VALUE_PLAINTEXT = "no value"
const NO_VALUE_HTML = "<i>no value</i>"
