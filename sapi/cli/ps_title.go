// <<generate>>

package cli

import (
	b "sik/builtin"
	"sik/zend"
)

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

// #define PS_TITLE_HEADER

const PS_TITLE_SUCCESS = 0
const PS_TITLE_NOT_AVAILABLE = 1
const PS_TITLE_NOT_INITIALIZED = 2
const PS_TITLE_BUFFER_NOT_AVAILABLE = 3
const PS_TITLE_WINDOWS_ERROR = 4

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

// # include "php_config.h"

var Environ **byte

// # include "ps_title.h"

// # include < stdio . h >

// # include < sys / types . h >

// # include < unistd . h >

// # include < string . h >

// # include < stdlib . h >

// # include < crt_externs . h >

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

// #define PS_USE_CLOBBER_ARGV

/* Different systems want the buffer padded differently */

const PS_PADDING = '0'

var PsBuffer *byte
var PsBufferSize int
var EmptyEnviron []*byte = []*byte{0}
var PsBufferCurLen int

/* save the original argv[] location here */

var SaveArgc int
var SaveArgv **byte

/*
 * This holds the 'locally' allocated environ from the save_ps_args method.
 * This is subsequently free'd at exit.
 */

var FrozenEnviron **byte
var NewEnviron ***byte

/*
 * Call this method early, before any code has used the original argv passed in
 * from main().
 * If needed, this code will make deep copies of argv and environ and return
 * these to the caller for further use. The original argv is then 'clobbered'
 * to store the process title.
 */

func SavePsArgs(argc int, argv **byte) **byte {
	SaveArgc = argc
	SaveArgv = argv

	/*
	 * If we're going to overwrite the argv area, count the available space.
	 * Also move the environment to make additional room.
	 */

	var end_of_area *byte = nil
	var non_contiguous_area int = 0
	var i int

	/*
	 * check for contiguous argv strings
	 */

	for i = 0; non_contiguous_area == 0 && i < argc; i++ {
		if i != 0 && end_of_area+1 != argv[i] {
			non_contiguous_area = 1
		}
		end_of_area = argv[i] + strlen(argv[i])
	}

	/*
	 * check for contiguous environ strings following argv
	 */

	for i = 0; non_contiguous_area == 0 && Environ[i] != nil; i++ {
		if end_of_area+1 != Environ[i] {
			non_contiguous_area = 1
		}
		end_of_area = Environ[i] + strlen(Environ[i])
	}
	if non_contiguous_area != 0 {
		goto clobber_error
	}
	PsBuffer = argv[0]
	PsBufferSize = end_of_area - argv[0]

	/*
	 * move the environment out of the way
	 */

	NewEnviron = (**byte)(zend.Malloc((i + 1) * b.SizeOf("char *")))
	FrozenEnviron = (**byte)(zend.Malloc((i + 1) * b.SizeOf("char *")))
	if NewEnviron == nil || FrozenEnviron == nil {
		goto clobber_error
	}
	for i = 0; Environ[i] != nil; i++ {
		NewEnviron[i] = strdup(Environ[i])
		if NewEnviron[i] == nil {
			goto clobber_error
		}
	}
	NewEnviron[i] = nil
	Environ = NewEnviron
	memcpy((*byte)(FrozenEnviron), (*byte)(NewEnviron), b.SizeOf("char *")*(i+1))

	/*
	 * If we're going to change the original argv[] then make a copy for
	 * argument parsing purposes.
	 *
	 * (NB: do NOT think to remove the copying of argv[]!
	 * On some platforms, getopt() keeps pointers into the argv array, and
	 * will get horribly confused when it is re-called to analyze a subprocess'
	 * argument string if the argv storage has been clobbered meanwhile.
	 * Other platforms have other dependencies on argv[].)
	 */

	var new_argv **byte
	var i int
	new_argv = (**byte)(zend.Malloc((argc + 1) * b.SizeOf("char *")))
	if new_argv == nil {
		goto clobber_error
	}
	for i = 0; i < argc; i++ {
		new_argv[i] = strdup(argv[i])
		if new_argv[i] == nil {
			zend.Free(new_argv)
			goto clobber_error
		}
	}
	new_argv[argc] = nil

	/*
	 * Darwin (and perhaps other NeXT-derived platforms?) has a static
	 * copy of the argv pointer, which we may fix like so:
	 */

	(*_NSGetArgv)() = new_argv
	argv = new_argv

	/* make extra argv slots point at end_of_area (a NUL) */

	var i int
	for i = 1; i < SaveArgc; i++ {
		SaveArgv[i] = PsBuffer + PsBufferSize
	}
	return argv
clobber_error:

	/* probably can't happen?!
	 * if we ever get here, argv still points to originally passed
	 * in argument
	 */

	SaveArgv = nil
	SaveArgc = 0
	PsBuffer = nil
	PsBufferSize = 0
	return argv
}

/*
 * Returns PS_TITLE_SUCCESS if the OS supports this functionality
 * and the init function was called.
 * Otherwise returns NOT_AVAILABLE or NOT_INITIALIZED
 */

func IsPsTitleAvailable() int {
	if SaveArgv == nil {
		return PS_TITLE_NOT_INITIALIZED
	}
	if PsBuffer == nil {
		return PS_TITLE_BUFFER_NOT_AVAILABLE
	}
	return PS_TITLE_SUCCESS
}

/*
 * Convert error codes into error strings
 */

func PsTitleErrno(rc int) *byte {
	switch rc {
	case PS_TITLE_SUCCESS:
		return "Success"
	case PS_TITLE_NOT_AVAILABLE:
		return "Not available on this OS"
	case PS_TITLE_NOT_INITIALIZED:
		return "Not initialized correctly"
	case PS_TITLE_BUFFER_NOT_AVAILABLE:
		return "Buffer not contiguous"
	}
	return "Unknown error code"
}

/*
 * Set a new process title.
 * Returns the appropriate error code if if there's an error
 * (like the functionality is compile time disabled, or the
 * save_ps_args() was not called.
 * Else returns 0 on success.
 */

func SetPsTitle(title *byte) int {
	var rc int = IsPsTitleAvailable()
	if rc != PS_TITLE_SUCCESS {
		return rc
	}
	strncpy(PsBuffer, title, PsBufferSize)
	PsBuffer[PsBufferSize-1] = '0'
	PsBufferCurLen = strlen(PsBuffer)

	/* pad unused memory */

	if PsBufferCurLen < PsBufferSize {
		memset(PsBuffer+PsBufferCurLen, PS_PADDING, PsBufferSize-PsBufferCurLen)
	}
	return PS_TITLE_SUCCESS
}

/*
 * Returns the current ps_buffer value into string.  On some platforms
 * the string will not be null-terminated, so return the effective
 * length into *displen.
 * The return code indicates the error.
 */

func GetPsTitle(displen *int, string **byte) int {
	var rc int = IsPsTitleAvailable()
	if rc != PS_TITLE_SUCCESS {
		return rc
	}
	*displen = int(PsBufferCurLen)
	*string = PsBuffer
	return PS_TITLE_SUCCESS
}

/*
 * Clean up the allocated argv and environ if applicable. Only call
 * this right before exiting.
 * This isn't needed per-se because the OS will clean-up anyway, but
 * having and calling this will ensure Valgrind doesn't output 'false
 * positives'.
 */

func CleanupPsArgs(argv **byte) {
	if SaveArgv != nil {
		SaveArgv = nil
		SaveArgc = 0
		var i int
		for i = 0; FrozenEnviron[i] != nil; i++ {
			zend.Free(FrozenEnviron[i])
		}
		zend.Free(FrozenEnviron)
		zend.Free(NewEnviron)

		/* leave a sane environment behind since some atexit() handlers
		   call getenv(). */

		Environ = EmptyEnviron

		/* leave a sane environment behind since some atexit() handlers
		   call getenv(). */

		var i int
		for i = 0; argv[i] != nil; i++ {
			zend.Free(argv[i])
		}
		zend.Free(argv)
	}
	return
}
