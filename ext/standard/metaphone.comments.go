// <<generate>>

package standard

// Source: <ext/standard/metaphone.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

/* {{{ proto string metaphone(string text[, int phones])
   Break english phrases down into their phonemes */

/* }}} */

/*-----------------------------  */

/* These letters are passed through unchanged */

/* These form diphthongs when preceding H */

/* These make C and G soft */

/* These prevent GH from becoming F */

/*----------------------------- */

/* Look at the current letter in the word */

/* Go N letters back. */

/* Previous letter.  I dunno, should this return null on failure? */

/* Look two letters down.  It makes sure you don't walk off the string. */

/* Allows us to safely look ahead an arbitrary # of letters */

/* phonize one letter
 * We don't know the buffers size in advance. On way to solve this is to just
 * re-allocate the buffer size. We're using an extra of 2 characters (this
 * could be one though; or more too). */

/* Slap a null character on the end of the phoned word */

/* How long is the phoned word? */

/* Note is a letter is a 'break' in the word */

/* {{{ metaphone
 */

/* }}} */
