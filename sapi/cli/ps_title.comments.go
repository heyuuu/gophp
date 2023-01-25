// <<generate>>

package cli

// Source: <sapi/cli/ps_title.h>

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
   | Authors: Keyur Govande <kgovande@gmail.com>                          |
   +----------------------------------------------------------------------+
*/

// Source: <sapi/cli/ps_title.c>

/*
 * PostgreSQL is released under the PostgreSQL License, a liberal Open Source
 * license, similar to the BSD or MIT licenses.
 * PostgreSQL Database Management System (formerly known as Postgres, then as
 * Postgres95)
 *
 * Portions Copyright (c) 1996-2015, The PostgreSQL Global Development Group
 *
 * Portions Copyright (c) 1994, The Regents of the University of California
 *
 * Permission to use, copy, modify, and distribute this software and its
 * documentation for any purpose, without fee, and without a written
 * agreement is hereby granted, provided that the above copyright notice
 * and this paragraph and the following two paragraphs appear in all copies.
 *
 * IN NO EVENT SHALL THE UNIVERSITY OF CALIFORNIA BE LIABLE TO ANY PARTY FOR
 * DIRECT, INDIRECT, SPECIAL, INCIDENTAL, OR CONSEQUENTIAL DAMAGES, INCLUDING
 * LOST PROFITS, ARISING OUT OF THE USE OF THIS SOFTWARE AND ITS DOCUMENTATION,
 * EVEN IF THE UNIVERSITY OF CALIFORNIA HAS BEEN ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 *
 * THE UNIVERSITY OF CALIFORNIA SPECIFICALLY DISCLAIMS ANY WARRANTIES,
 * INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND
 * FITNESS FOR A PARTICULAR PURPOSE. THE SOFTWARE PROVIDED HEREUNDER IS ON AN
 * "AS IS" BASIS, AND THE UNIVERSITY OF CALIFORNIA HAS NO OBLIGATIONS TO
 * PROVIDE MAINTENANCE, SUPPORT, UPDATES, ENHANCEMENTS, OR MODIFICATIONS.
 *
 * The following code is adopted from the PostgreSQL's ps_status(.h/.c).
 */

/*
 * Ways of updating ps display:
 *
 * PS_USE_SETPROCTITLE
 *         use the function setproctitle(const char *, ...)
 *         (newer BSD systems)
 * PS_USE_PSTAT
 *         use the pstat(PSTAT_SETCMD, )
 *         (HPUX)
 * PS_USE_PS_STRINGS
 *         assign PS_STRINGS->ps_argvstr = "string"
 *         (some BSD systems)
 * PS_USE_CHANGE_ARGV
 *         assign argv[0] = "string"
 *         (some other BSD systems)
 * PS_USE_CLOBBER_ARGV
 *         write over the argv and environment area
 *         (Linux and most SysV-like systems)
 * PS_USE_WIN32
 *         push the string out as the name of a Windows event
 * PS_USE_NONE
 *         don't update ps display
 *         (This is the default, as it is safest.)
 */

/* Different systems want the buffer padded differently */

/* save the original argv[] location here */

/*
 * This holds the 'locally' allocated environ from the save_ps_args method.
 * This is subsequently free'd at exit.
 */

/*
 * Call this method early, before any code has used the original argv passed in
 * from main().
 * If needed, this code will make deep copies of argv and environ and return
 * these to the caller for further use. The original argv is then 'clobbered'
 * to store the process title.
 */

/*
 * Returns PS_TITLE_SUCCESS if the OS supports this functionality
 * and the init function was called.
 * Otherwise returns NOT_AVAILABLE or NOT_INITIALIZED
 */

/*
 * Convert error codes into error strings
 */

/*
 * Set a new process title.
 * Returns the appropriate error code if if there's an error
 * (like the functionality is compile time disabled, or the
 * save_ps_args() was not called.
 * Else returns 0 on success.
 */

/*
 * Returns the current ps_buffer value into string.  On some platforms
 * the string will not be null-terminated, so return the effective
 * length into *displen.
 * The return code indicates the error.
 */

/*
 * Clean up the allocated argv and environ if applicable. Only call
 * this right before exiting.
 * This isn't needed per-se because the OS will clean-up anyway, but
 * having and calling this will ensure Valgrind doesn't output 'false
 * positives'.
 */
