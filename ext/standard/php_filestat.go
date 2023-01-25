// <<generate>>

package standard

// Source: <ext/standard/php_filestat.h>

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
   | Author:  Jim Winstead <jimw@php.net>                                 |
   +----------------------------------------------------------------------+
*/

// #define PHP_FILESTAT_H

/* Compatibility. */

type PhpStatLen = int

/* Switches for various filestat functions: */

const FS_PERMS = 0
const FS_INODE = 1
const FS_SIZE = 2
const FS_OWNER = 3
const FS_GROUP = 4
const FS_ATIME = 5
const FS_MTIME = 6
const FS_CTIME = 7
const FS_TYPE = 8
const FS_IS_W = 9
const FS_IS_R = 10
const FS_IS_X = 11
const FS_IS_FILE = 12
const FS_IS_DIR = 13
const FS_IS_LINK = 14
const FS_EXISTS = 15
const FS_LSTAT = 16
const FS_STAT = 17
