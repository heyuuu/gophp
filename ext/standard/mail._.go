package standard

// Source: <ext/standard/mail.c>

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
   | Author: Rasmus Lerdorf <rasmus@php.net>                              |
   +----------------------------------------------------------------------+
*/

// failed # include "ext/date/php_date.h"

/* {{{ proto int ezmlm_hash(string addr)
   Calculate EZMLM list hash value. */

/* {{{ proto int mail(string to, string subject, string message [, string additional_headers [, string additional_parameters]])
   Send an email message */

/* {{{ php_mail
 */
