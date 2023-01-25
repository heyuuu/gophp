// <<generate>>

package core

// Source: <main/php_getopt.h>

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

// #define PHP_GETOPT_H

// # include "php.h"

/* Define structure for one recognized option (both single char and long name).
 * If short_open is '-' this is the last option. */

type _opt = Opt

/* holds the index of the latest fetched element from the opts array */

/* php_getopt will return this value if there is an error in arguments */

const PHP_GETOPT_INVALID_ARG = -2
