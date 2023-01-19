// <<generate>>

package core

import (
	"sik/zend"
)

// Source: <main/php_open_temporary_file.h>

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

// #define PHP_OPEN_TEMPORARY_FILE_H

// #define PHP_TMP_FILE_DEFAULT       0

// #define PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK       ( 1 << 0 )

// #define PHP_TMP_FILE_SILENT       ( 1 << 1 )

// #define PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_EXPLICIT_DIR       ( 1 << 2 )

// #define PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ALWAYS       ( PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK | PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_EXPLICIT_DIR )

/* for compatibility purpose */

// #define PHP_TMP_FILE_OPEN_BASEDIR_CHECK       PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK

// Source: <main/php_open_temporary_file.c>

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

// # include "php.h"

// # include "php_open_temporary_file.h"

// # include < errno . h >

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < fcntl . h >

// # include < sys / param . h >

// # include < sys / socket . h >

// # include < netinet / in . h >

// # include < netdb . h >

// # include < arpa / inet . h >

// # include < sys / time . h >

// # include < sys / file . h >

// #define P_tmpdir       ""

/* {{{ php_do_open_temporary_file */

func PhpDoOpenTemporaryFile(path *byte, pfx *byte, opened_path_p **zend.ZendString) int {
	var opened_path []byte
	var trailing_slash *byte
	var cwd []byte
	var new_state zend.CwdState
	var fd int = -1
	if path == nil || !(path[0]) {
		return -1
	}
	if !(getcwd(cwd, 256)) {
		cwd[0] = '0'
	}
	new_state.cwd = zend._estrdup(cwd)
	new_state.cwd_length = strlen(cwd)
	if zend.VirtualFileEx(&new_state, path, nil, 2) != 0 {
		zend._efree(new_state.cwd)
		return -1
	}
	if new_state.cwd[new_state.cwd_length-1] == '/' {
		trailing_slash = ""
	} else {
		trailing_slash = "/"
	}
	if ApPhpSnprintf(opened_path, 256, "%s%s%sXXXXXX", new_state.cwd, trailing_slash, pfx) >= 256 {
		zend._efree(new_state.cwd)
		return -1
	}
	fd = mkstemp(opened_path)
	if fd != -1 && opened_path_p != nil {
		*opened_path_p = zend.ZendStringInit(opened_path, strlen(opened_path), 0)
	}
	zend._efree(new_state.cwd)
	return fd
}

/* }}} */

func PhpGetTemporaryDirectory() *byte {
	/* Did we determine the temporary directory already? */

	if CoreGlobals.GetPhpSysTempDir() != nil {
		return CoreGlobals.GetPhpSysTempDir()
	}

	/* Is there a temporary directory "sys_temp_dir" in .ini defined? */

	var sys_temp_dir *byte = CoreGlobals.GetSysTempDir()
	if sys_temp_dir != nil {
		var len_ int = strlen(sys_temp_dir)
		if len_ >= 2 && sys_temp_dir[len_-1] == '/' {
			CoreGlobals.SetPhpSysTempDir(zend._estrndup(sys_temp_dir, len_-1))
			return CoreGlobals.GetPhpSysTempDir()
		} else if len_ >= 1 && sys_temp_dir[len_-1] != '/' {
			CoreGlobals.SetPhpSysTempDir(zend._estrndup(sys_temp_dir, len_))
			return CoreGlobals.GetPhpSysTempDir()
		}
	}

	/* On Unix use the (usual) TMPDIR environment variable. */

	var s *byte = getenv("TMPDIR")
	if s != nil && (*s) {
		var len_ int = strlen(s)
		if s[len_-1] == '/' {
			CoreGlobals.SetPhpSysTempDir(zend._estrndup(s, len_-1))
		} else {
			CoreGlobals.SetPhpSysTempDir(zend._estrndup(s, len_))
		}
		return CoreGlobals.GetPhpSysTempDir()
	}

	/* Use the standard default temporary directory. */

	if "" {
		CoreGlobals.SetPhpSysTempDir(zend._estrdup(""))
		return CoreGlobals.GetPhpSysTempDir()
	}

	/* Shouldn't ever(!) end up here ... last ditch default. */

	CoreGlobals.SetPhpSysTempDir(zend._estrdup("/tmp"))
	return CoreGlobals.GetPhpSysTempDir()
}

/* {{{ php_open_temporary_file
 *
 * Unlike tempnam(), the supplied dir argument takes precedence
 * over the TMPDIR environment variable
 * This function should do its best to return a file pointer to a newly created
 * unique file, on every platform.
 */

func PhpOpenTemporaryFdEx(dir *byte, pfx *byte, opened_path_p **zend.ZendString, flags uint32) int {
	var fd int
	var temp_dir *byte
	if pfx == nil {
		pfx = "tmp."
	}
	if opened_path_p != nil {
		*opened_path_p = nil
	}
	if dir == nil || (*dir) == '0' {
	def_tmp:
		temp_dir = PhpGetTemporaryDirectory()
		if temp_dir != nil && (*temp_dir) != '0' && ((flags&1<<0) == 0 || PhpCheckOpenBasedir(temp_dir) == 0) {
			return PhpDoOpenTemporaryFile(temp_dir, pfx, opened_path_p)
		} else {
			return -1
		}
	}
	if (flags&1<<2) != 0 && PhpCheckOpenBasedir(dir) != 0 {
		return -1
	}

	/* Try the directory given as parameter. */

	fd = PhpDoOpenTemporaryFile(dir, pfx, opened_path_p)
	if fd == -1 {

		/* Use default temporary directory. */

		if (flags & 1 << 1) == 0 {
			PhpErrorDocref(nil, 1<<3, "file created in the system's temporary directory")
		}
		goto def_tmp
	}
	return fd
}
func PhpOpenTemporaryFd(dir *byte, pfx string, opened_path_p **zend.ZendString) int {
	return PhpOpenTemporaryFdEx(dir, pfx, opened_path_p, 0)
}
func PhpOpenTemporaryFile(dir *byte, pfx *byte, opened_path_p **zend.ZendString) *FILE {
	var fp *FILE
	var fd int = PhpOpenTemporaryFd(dir, pfx, opened_path_p)
	if fd == -1 {
		return nil
	}
	fp = fdopen(fd, "r+b")
	if fp == nil {
		close(fd)
	}
	return fp
}

/* }}} */
