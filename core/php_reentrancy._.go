package core

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

/* currently, PHP does not check for these functions, but assumes
   that they are available on all systems. */

const HAVE_LOCALTIME = 1
const HAVE_GMTIME = 1
const HAVE_ASCTIME = 1
const HAVE_CTIME = 1
const PhpLocaltimeR = localtime_r
const PhpCtimeR = ctime_r
const PhpAsctimeR = asctime_r
const PhpGmtimeR = gmtime_r
const PhpStrtokR = strtok_r
const PhpRandR = rand_r
