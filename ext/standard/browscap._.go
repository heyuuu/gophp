// <<generate>>

package standard

// Source: <ext/standard/browscap.c>

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

// failed # include "ext/pcre/php_pcre.h"

const BROWSCAP_NUM_CONTAINS = 5

/* browser data defined in startup phase, eagerly loaded in MINIT */

var GlobalBdata BrowserData = BrowserData{0}

/* browser data defined in activation phase, lazily loaded in get_browser.
 * Per request and per thread, if applicable */

var BrowscapGlobals ZendBrowscapGlobals

const DEFAULT_SECTION_NAME = "Default Browser Capability Settings"

/* OBJECTS_FIXME: This whole extension needs going through. The use of objects looks pretty broken here */

/* Length of prefix not containing any wildcards */

/* Length of regex, including escapes, anchors, etc. */
