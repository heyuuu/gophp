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
	core.PUTS("body {background-color: #fff; color: #222; font-family: sans-serif;}\n")
	core.PUTS("pre {margin: 0; font-family: monospace;}\n")
	core.PUTS("a:link {color: #009; text-decoration: none; background-color: #fff;}\n")
	core.PUTS("a:hover {text-decoration: underline;}\n")
	core.PUTS("table {border-collapse: collapse; border: 0; width: 934px; box-shadow: 1px 2px 3px #ccc;}\n")
	core.PUTS(".center {text-align: center;}\n")
	core.PUTS(".center table {margin: 1em auto; text-align: left;}\n")
	core.PUTS(".center th {text-align: center !important;}\n")
	core.PUTS("td, th {border: 1px solid #666; font-size: 75%; vertical-align: baseline; padding: 4px 5px;}\n")
	core.PUTS("th {position: sticky; top: 0; background: inherit;}\n")
	core.PUTS("h1 {font-size: 150%;}\n")
	core.PUTS("h2 {font-size: 125%;}\n")
	core.PUTS(".p {text-align: left;}\n")
	core.PUTS(".e {background-color: #ccf; width: 300px; font-weight: bold;}\n")
	core.PUTS(".h {background-color: #99c; font-weight: bold;}\n")
	core.PUTS(".v {background-color: #ddd; max-width: 300px; overflow-x: auto; word-wrap: break-word;}\n")
	core.PUTS(".v i {color: #999;}\n")
	core.PUTS("img {float: right; border: 0;}\n")
	core.PUTS("hr {width: 934px; background-color: #ccc; border: 0; height: 1px;}\n")
}

/* }}} */
