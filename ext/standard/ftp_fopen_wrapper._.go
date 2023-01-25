// <<generate>>

package standard

import (
	"sik/core"
)

// Source: <ext/standard/ftp_fopen_wrapper.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Jim Winstead <jimw@php.net>                                 |
   |          Hartmut Holzgraefe <hholzgra@php.net>                       |
   |          Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

const FTPS_ENCRYPT_DATA = 1

/* {{{ get_ftp_result
 */

var PhpFtpDirstreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpFtpDirstreamRead, PhpFtpDirstreamClose, nil, "ftpdir", nil, nil, nil, nil}

/* {{{ php_stream_ftp_opendir
 */

var FtpStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{PhpStreamUrlWrapFtp, PhpStreamFtpStreamClose, PhpStreamFtpStreamStat, PhpStreamFtpUrlStat, PhpStreamFtpOpendir, "ftp", PhpStreamFtpUnlink, PhpStreamFtpRename, PhpStreamFtpMkdir, PhpStreamFtpRmdir, nil}
var PhpStreamFtpWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&FtpStreamWops, nil, 1}
