// <<generate>>

package core

import (
	"sik/zend"
)

// Source: <main/php_ini.h>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

const PHP_INI_USER = zend.ZEND_INI_USER
const PHP_INI_PERDIR = zend.ZEND_INI_PERDIR
const PHP_INI_SYSTEM = zend.ZEND_INI_SYSTEM
const PHP_INI_ALL = zend.ZEND_INI_ALL
const PhpIniEntry = zend_ini_entry
const PHP_INI_ENTRY3_EX = ZEND_INI_ENTRY3_EX
const PHP_INI_ENTRY3 = ZEND_INI_ENTRY3
const PHP_INI_ENTRY2_EX = ZEND_INI_ENTRY2_EX
const PHP_INI_ENTRY2 = ZEND_INI_ENTRY2
const PHP_INI_ENTRY1_EX = ZEND_INI_ENTRY1_EX
const PHP_INI_ENTRY1 = ZEND_INI_ENTRY1
const PHP_INI_DISPLAY_ORIG = zend.ZEND_INI_DISPLAY_ORIG
const PHP_INI_DISPLAY_ACTIVE = zend.ZEND_INI_DISPLAY_ACTIVE
const PHP_INI_STAGE_STARTUP = zend.ZEND_INI_STAGE_STARTUP
const PHP_INI_STAGE_SHUTDOWN = zend.ZEND_INI_STAGE_SHUTDOWN
const PHP_INI_STAGE_ACTIVATE = zend.ZEND_INI_STAGE_ACTIVATE
const PHP_INI_STAGE_DEACTIVATE = zend.ZEND_INI_STAGE_DEACTIVATE
const PHP_INI_STAGE_RUNTIME = zend.ZEND_INI_STAGE_RUNTIME
const PHP_INI_STAGE_HTACCESS = zend.ZEND_INI_STAGE_HTACCESS
const PhpIniBooleanDisplayerCb = zend.ZendIniBooleanDisplayerCb
const PhpIniColorDisplayerCb = zend.ZendIniColorDisplayerCb
const PhpAlterIniEntry = zend.ZendAlterIniEntry
const PhpIniLong = zend.ZendIniLong
const PhpIniDouble = zend.ZendIniDouble
const PhpIniString = zend.ZendIniString

// Source: <main/php_ini.c>

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
   | Author: Zeev Suraski <zeev@php.net>                                  |
   +----------------------------------------------------------------------+
*/

/* True globals */

var IsSpecialSection int = 0
var ActiveIniHash *zend.HashTable
var HasPerDirConfig int = 0
var HasPerHostConfig int = 0
var PhpIniOpenedPath *byte = nil
var ExtensionLists PhpExtensionLists
var PhpIniScannedPath *byte = nil
var PhpIniScannedFiles *byte = nil

/* {{{ php_ini_displayer_cb
 */

const PHP_EXTENSION_TOKEN = "extension"
const ZEND_EXTENSION_TOKEN = "zend_extension"

/* {{{ config_zval_dtor
 */

/* Reset / free active_ini_sectin global */
