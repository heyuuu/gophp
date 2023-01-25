// <<generate>>

package zend

// Source: <Zend/zend_globals_macros.h>

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

// #define ZEND_GLOBALS_MACROS_H

/* Compiler */

// #define CG(v) ( compiler_globals . v )

var Zendparse func() int

/* Executor */

// #define EG(v) ( executor_globals . v )

/* Language Scanner */

// #define LANG_SCNG(v) ( language_scanner_globals . v )

var LanguageScannerGlobals ZendPhpScannerGlobals

/* INI Scanner */

// #define INI_SCNG(v) ( ini_scanner_globals . v )

var IniScannerGlobals ZendIniScannerGlobals
