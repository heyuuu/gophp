// <<generate>>

package standard

import (
	"sik/core"
)

// Source: <ext/standard/css.h>

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
   | Authors: Colin Viebrock <colin@viebrock.ca>                          |
   +----------------------------------------------------------------------+
*/

// #define CSS_H

// Source: <ext/standard/css.c>

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
   | Authors: Colin Viebrock <colin@viebrock.ca>                          |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "info.h"

func PhpInfoPrintCss() {
	var __str *byte = "body {background-color: #fff; color: #222; font-family: sans-serif;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "pre {margin: 0; font-family: monospace;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "a:link {color: #009; text-decoration: none; background-color: #fff;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "a:hover {text-decoration: underline;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "table {border-collapse: collapse; border: 0; width: 934px; box-shadow: 1px 2px 3px #ccc;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".center {text-align: center;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".center table {margin: 1em auto; text-align: left;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".center th {text-align: center !important;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "td, th {border: 1px solid #666; font-size: 75%; vertical-align: baseline; padding: 4px 5px;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "th {position: sticky; top: 0; background: inherit;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "h1 {font-size: 150%;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "h2 {font-size: 125%;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".p {text-align: left;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".e {background-color: #ccf; width: 300px; font-weight: bold;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".h {background-color: #99c; font-weight: bold;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".v {background-color: #ddd; max-width: 300px; overflow-x: auto; word-wrap: break-word;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = ".v i {color: #999;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "img {float: right; border: 0;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
	var __str *byte = "hr {width: 934px; background-color: #ccc; border: 0; height: 1px;}\n"
	core.PhpOutputWrite(__str, strlen(__str))
}

/* }}} */
