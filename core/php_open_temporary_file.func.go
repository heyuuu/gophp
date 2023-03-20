// <<generate>>

package core

import (
	r "sik/builtin/file"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func PhpDoOpenTemporaryFile(path *byte, pfx *byte, opened_path_p **types.String) int {
	var opened_path []byte
	var trailing_slash *byte
	var cwd []byte
	var new_state zend.CwdState
	var fd int = -1
	if path == nil || !(path[0]) {
		return -1
	}
	if !(zend.VCWD_GETCWD(cwd, MAXPATHLEN)) {
		cwd[0] = '0'
	}
	new_state.SetCwd(zend.Estrdup(cwd))
	new_state.SetCwdLength(strlen(cwd))
	if zend.VirtualFileEx(&new_state, path, nil, zend.CWD_REALPATH) != 0 {
		zend.Efree(new_state.GetCwd())
		return -1
	}
	if zend.IS_SLASH(new_state.GetCwd()[new_state.GetCwdLength()-1]) {
		trailing_slash = ""
	} else {
		trailing_slash = "/"
	}
	if Snprintf(opened_path, MAXPATHLEN, "%s%s%sXXXXXX", new_state.GetCwd(), trailing_slash, pfx) >= MAXPATHLEN {
		zend.Efree(new_state.GetCwd())
		return -1
	}
	fd = mkstemp(opened_path)
	if fd != -1 && opened_path_p != nil {
		*opened_path_p = types.ZendStringInit(opened_path)
	}
	zend.Efree(new_state.GetCwd())
	return fd
}
func PhpGetTemporaryDirectory() *byte {
	/* Did we determine the temporary directory already? */

	if PG__().php_sys_temp_dir {
		return PG__().php_sys_temp_dir
	}

	/* Is there a temporary directory "sys_temp_dir" in .ini defined? */

	var sys_temp_dir *byte = PG__().sys_temp_dir
	if sys_temp_dir != nil {
		var len_ int = strlen(sys_temp_dir)
		if len_ >= 2 && sys_temp_dir[len_-1] == zend.DEFAULT_SLASH {
			PG__().php_sys_temp_dir = zend.Estrndup(sys_temp_dir, len_-1)
			return PG__().php_sys_temp_dir
		} else if len_ >= 1 && sys_temp_dir[len_-1] != zend.DEFAULT_SLASH {
			PG__().php_sys_temp_dir = zend.Estrndup(sys_temp_dir, len_)
			return PG__().php_sys_temp_dir
		}
	}

	/* On Unix use the (usual) TMPDIR environment variable. */

	var s *byte = getenv("TMPDIR")
	if s != nil && (*s) {
		var len_ int = strlen(s)
		if s[len_-1] == zend.DEFAULT_SLASH {
			PG__().php_sys_temp_dir = zend.Estrndup(s, len_-1)
		} else {
			PG__().php_sys_temp_dir = zend.Estrndup(s, len_)
		}
		return PG__().php_sys_temp_dir
	}

	/* Use the standard default temporary directory. */

	if P_tmpdir {
		PG__().php_sys_temp_dir = zend.Estrdup(P_tmpdir)
		return PG__().php_sys_temp_dir
	}

	/* Shouldn't ever(!) end up here ... last ditch default. */

	PG__().php_sys_temp_dir = zend.Estrdup("/tmp")
	return PG__().php_sys_temp_dir
}
func PhpOpenTemporaryFdEx(dir *byte, pfx *byte, opened_path_p **types.String, flags uint32) int {
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
		if temp_dir != nil && (*temp_dir) != '0' && ((flags&PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_FALLBACK) == 0 || PhpCheckOpenBasedir(temp_dir) == 0) {
			return PhpDoOpenTemporaryFile(temp_dir, pfx, opened_path_p)
		} else {
			return -1
		}
	}
	if (flags&PHP_TMP_FILE_OPEN_BASEDIR_CHECK_ON_EXPLICIT_DIR) != 0 && PhpCheckOpenBasedir(dir) != 0 {
		return -1
	}

	/* Try the directory given as parameter. */

	fd = PhpDoOpenTemporaryFile(dir, pfx, opened_path_p)
	if fd == -1 {

		/* Use default temporary directory. */

		if (flags & PHP_TMP_FILE_SILENT) == 0 {
			PhpErrorDocref(nil, faults.E_NOTICE, "file created in the system's temporary directory")
		}
		goto def_tmp
	}
	return fd
}
func PhpOpenTemporaryFd(dir *byte, pfx string, opened_path_p **types.String) int {
	return PhpOpenTemporaryFdEx(dir, pfx, opened_path_p, PHP_TMP_FILE_DEFAULT)
}
func PhpOpenTemporaryFile(dir *byte, pfx *byte, opened_path_p **types.String) *r.FILE {
	var fp *r.FILE
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
