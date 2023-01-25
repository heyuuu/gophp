// <<generate>>

package standard

import (
	"sik/zend"
)

// Source: <ext/standard/password.c>

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
   | Authors: Anthony Ferrara <ircmaxell@php.net>                         |
   |          Charles R. Portwood II <charlesportwoodii@erianna.com>      |
   +----------------------------------------------------------------------+
*/

var PhpPasswordAlgos zend.ZendArray

/* bcrypt implementation */

var PhpPasswordAlgoBcrypt PhpPasswordAlgo = PhpPasswordAlgo{"bcrypt", PhpPasswordBcryptHash, PhpPasswordBcryptVerify, PhpPasswordBcryptNeedsRehash, PhpPasswordBcryptGetInfo, PhpPasswordBcryptValid}

/* {{{ proto array password_get_info(string $hash)
Retrieves information about a given hash */

/** }}} */
