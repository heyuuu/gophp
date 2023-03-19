// <<generate>>

package core

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/types"
)

func OnUpdateBaseDir(
	entry *zend.ZendIniEntry,
	new_value *types.ZendString,
	mh_arg1 any,
	mh_arg2 any,
	mh_arg3 any,
	stage int,
) int {
	var p **byte
	var pathbuf **byte
	var ptr **byte
	var end **byte
	var base *byte = (*byte)(mh_arg2)
	p = (**byte)(base + int(mh_arg1))
	if stage == PHP_INI_STAGE_STARTUP || stage == PHP_INI_STAGE_SHUTDOWN || stage == PHP_INI_STAGE_ACTIVATE || stage == PHP_INI_STAGE_DEACTIVATE {

		/* We're in a PHP_INI_SYSTEM context, no restrictions */

		if new_value != nil {
			*p = new_value.GetVal()
		} else {
			*p = nil
		}
		return types.SUCCESS
	}

	/* Otherwise we're in runtime */

	if (*p) == nil || !(*(*p)) {

		/* open_basedir not set yet, go ahead and give it a value */

		*p = new_value.GetVal()
		return types.SUCCESS
	}

	/* Shortcut: When we have a open_basedir and someone tries to unset, we know it'll fail */

	if new_value == nil || !new_value.GetVal() {
		return types.FAILURE
	}

	/* Is the proposed open_basedir at least as restrictive as the current setting? */

	pathbuf = zend.Estrdup(new_value.GetVal())
	ptr = pathbuf
	for ptr != nil && (*ptr) != nil {
		end = strchr(ptr, zend.DEFAULT_DIR_SEPARATOR)
		if end != nil {
			*end = '0'
			end++
		}
		if ptr[0] == '.' && ptr[1] == '.' && (ptr[2] == '0' || zend.IS_SLASH(ptr[2])) {

			/* Don't allow paths with a leading .. path component to be set at runtime */

			zend.Efree(pathbuf)
			return types.FAILURE
		}
		if PhpCheckOpenBasedirEx(ptr, 0) != 0 {

			/* At least one portion of this open_basedir is less restrictive than the prior one, FAIL */

			zend.Efree(pathbuf)
			return types.FAILURE
		}
		ptr = end
	}
	zend.Efree(pathbuf)

	/* Everything checks out, set it */

	*p = new_value.GetVal()
	return types.SUCCESS
}
func PhpCheckSpecificOpenBasedir(basedir *byte, path *byte) int {
	var resolved_name []byte
	var resolved_basedir []byte
	var local_open_basedir []byte
	var path_tmp []byte
	var path_file *byte
	var resolved_basedir_len int
	var resolved_name_len int
	var path_len int
	var nesting_level int = 0

	/* Special case basedir==".": Use script-directory */

	if strcmp(basedir, ".") || !(zend.VCWD_GETCWD(local_open_basedir, MAXPATHLEN)) {

		/* Else use the unmodified path */

		strlcpy(local_open_basedir, basedir, b.SizeOf("local_open_basedir"))

		/* Else use the unmodified path */

	}
	path_len = strlen(path)
	if path_len > MAXPATHLEN-1 {

		/* empty and too long paths are invalid */

		return -1

		/* empty and too long paths are invalid */

	}

	/* normalize and expand path */

	if ExpandFilepath(path, resolved_name) == nil {
		return -1
	}
	path_len = strlen(resolved_name)
	memcpy(path_tmp, resolved_name, path_len+1)
	for zend.VCWD_REALPATH(path_tmp, resolved_name) == nil {
		if nesting_level == 0 {
			var ret ssize_t
			var buf []byte
			ret = zend.PhpSysReadlink(path_tmp, buf, MAXPATHLEN-1)
			if ret == -1 {

			} else {

				/* put the real path into the path buffer */

				memcpy(path_tmp, buf, ret)
				path_tmp[ret] = '0'
			}
		}
		path_file = strrchr(path_tmp, zend.DEFAULT_SLASH)
		if path_file == nil {

			/* none of the path components exist. definitely not in open_basedir.. */

			return -1

			/* none of the path components exist. definitely not in open_basedir.. */

		} else {
			path_len = path_file - path_tmp + 1
			path_tmp[path_len-1] = '0'
		}
		if (*path_tmp) == '0' {

			/* Do not pass an empty string to realpath(), as this will resolve to CWD. */

			break

			/* Do not pass an empty string to realpath(), as this will resolve to CWD. */

		}
		nesting_level++
	}

	/* Resolve open_basedir to resolved_basedir */

	if ExpandFilepath(local_open_basedir, resolved_basedir) != nil {
		var basedir_len int = strlen(basedir)

		/* Handler for basedirs that end with a / */

		resolved_basedir_len = strlen(resolved_basedir)
		if basedir[basedir_len-1] == PHP_DIR_SEPARATOR {
			if resolved_basedir[resolved_basedir_len-1] != PHP_DIR_SEPARATOR {
				resolved_basedir[resolved_basedir_len] = PHP_DIR_SEPARATOR
				resolved_basedir[b.PreInc(&resolved_basedir_len)] = '0'
			}
		} else {
			resolved_basedir[b.PostInc(&resolved_basedir_len)] = PHP_DIR_SEPARATOR
			resolved_basedir[resolved_basedir_len] = '0'
		}
		resolved_name_len = strlen(resolved_name)
		if path_tmp[path_len-1] == PHP_DIR_SEPARATOR {
			if resolved_name[resolved_name_len-1] != PHP_DIR_SEPARATOR {
				resolved_name[resolved_name_len] = PHP_DIR_SEPARATOR
				resolved_name[b.PreInc(&resolved_name_len)] = '0'
			}
		}

		/* Check the path */

		if strncmp(resolved_basedir, resolved_name, resolved_basedir_len) == 0 {
			if resolved_name_len > resolved_basedir_len && resolved_name[resolved_basedir_len-1] != PHP_DIR_SEPARATOR {
				return -1
			} else {

				/* File is in the right directory */

				return 0

				/* File is in the right directory */

			}
		} else {

			/* /openbasedir/ and /openbasedir are the same directory */

			if resolved_basedir_len == resolved_name_len+1 && resolved_basedir[resolved_basedir_len-1] == PHP_DIR_SEPARATOR {
				if strncmp(resolved_basedir, resolved_name, resolved_name_len) == 0 {
					return 0
				}
			}
			return -1
		}

		/* Check the path */

	} else {

		/* Unable to resolve the real path, return -1 */

		return -1

		/* Unable to resolve the real path, return -1 */

	}

	/* Resolve open_basedir to resolved_basedir */
}
func PhpCheckOpenBasedir(path *byte) int { return PhpCheckOpenBasedirEx(path, 1) }
func PhpCheckOpenBasedirEx(path *byte, warn int) int {
	/* Only check when open_basedir is available */

	if PG__().open_basedir && *PG__().open_basedir {
		var pathbuf *byte
		var ptr *byte
		var end *byte

		/* Check if the path is too long so we can give a more useful error
		 * message. */

		if strlen(path) > MAXPATHLEN-1 {
			PhpErrorDocref(nil, zend.E_WARNING, "File name is longer than the maximum allowed path length on this platform (%d): %s", MAXPATHLEN, path)
			errno = EINVAL
			return -1
		}
		pathbuf = zend.Estrdup(PG__().open_basedir)
		ptr = pathbuf
		for ptr != nil && (*ptr) {
			end = strchr(ptr, zend.DEFAULT_DIR_SEPARATOR)
			if end != nil {
				*end = '0'
				end++
			}
			if PhpCheckSpecificOpenBasedir(ptr, path) == 0 {
				zend.Efree(pathbuf)
				return 0
			}
			ptr = end
		}
		if warn != 0 {
			PhpErrorDocref(nil, zend.E_WARNING, "open_basedir restriction in effect. File(%s) is not within the allowed path(s): (%s)", path, PG__().open_basedir)
		}
		zend.Efree(pathbuf)
		errno = EPERM
		return -1
	}

	/* Nothing to check... */

	return 0

	/* Nothing to check... */
}
func PhpFopenAndSetOpenedPath(path *byte, mode string, opened_path **types.ZendString) *r.FILE {
	var fp *r.FILE
	if PhpCheckOpenBasedir((*byte)(path)) != 0 {
		return nil
	}
	fp = zend.VCWD_FOPEN(path, mode)
	if fp != nil && opened_path != nil {

		//TODO :avoid reallocation

		var tmp *byte = ExpandFilepathWithMode(path, nil, nil, 0, zend.CWD_EXPAND)
		if tmp != nil {
			*opened_path = types.ZendStringInit(tmp, strlen(tmp), 0)
			zend.Efree(tmp)
		}
	}
	return fp
}
func PhpFopenPrimaryScript(file_handle *zend.ZendFileHandle) int {
	var path_info *byte
	var filename *byte = nil
	var resolved_path *string = nil
	var length int
	var orig_display_errors types.ZendBool
	path_info = SG__().request_info.request_uri
	if PG__().user_dir && *PG__().user_dir && path_info != nil && '/' == path_info[0] && '~' == path_info[1] {
		var s *byte = strchr(path_info+2, '/')
		if s != nil {
			var user []byte
			var pw *__struct__passwd
			length = s - (path_info + 2)
			if length > b.SizeOf("user")-1 {
				length = b.SizeOf("user") - 1
			}
			memcpy(user, path_info+2, length)
			user[length] = '0'
			pw = getpwnam(user)
			if pw != nil && pw.pw_dir {
				Spprintf(&filename, 0, "%s%c%s%c%s", pw.pw_dir, PHP_DIR_SEPARATOR, PG__().user_dir, PHP_DIR_SEPARATOR, s+1)
			} else {
				filename = SG__().request_info.path_translated
			}
		}
	} else if PG__().doc_root && path_info != nil && b.Assign(&length, strlen(PG__().doc_root)) && zend.IS_ABSOLUTE_PATH(PG__().doc_root, length) {
		var path_len int = strlen(path_info)
		filename = zend.Emalloc(length + path_len + 2)
		memcpy(filename, PG__().doc_root, length)
		if !(zend.IS_SLASH(filename[length-1])) {
			filename[b.PostInc(&length)] = PHP_DIR_SEPARATOR
		}
		if zend.IS_SLASH(path_info[0]) {
			length--
		}
		strncpy(filename+length, path_info, path_len+1)
	} else {
		filename = SG__().request_info.path_translated
	}
	if filename != nil {
		resolved_path = zend.ZendResolvePath(filename)
	}
	if resolved_path == nil {
		if SG__().request_info.path_translated != filename {
			if filename != nil {
				zend.Efree(filename)
			}
		}

		/* we have to free SG__().request_info.path_translated here because
		 * php_destroy_request_info assumes that it will get
		 * freed when the include_names hash is emptied, but
		 * we're not adding it in this case */

		if SG__().request_info.path_translated {
			zend.Efree(SG__().request_info.path_translated)
			SG__().request_info.path_translated = nil
		}
		return types.FAILURE
	}
	orig_display_errors = PG__().display_errors
	PG__().display_errors = 0
	if zend.ZendStreamOpen(filename, file_handle) == types.FAILURE {
		PG__().display_errors = orig_display_errors
		if SG__().request_info.path_translated != filename {
			if filename != nil {
				zend.Efree(filename)
			}
		}
		if SG__().request_info.path_translated {
			zend.Efree(SG__().request_info.path_translated)
			SG__().request_info.path_translated = nil
		}
		return types.FAILURE
	}
	PG__().display_errors = orig_display_errors
	if SG__().request_info.path_translated != filename {
		if SG__().request_info.path_translated {
			zend.Efree(SG__().request_info.path_translated)
		}
		SG__().request_info.path_translated = filename
	}
	return types.SUCCESS
}
func PhpResolvePath(fileName string, filenamePtr *byte, filename_length int, path *byte) *types.ZendString {
	var resolved_path []byte
	var trypath []byte
	var ptr *byte
	var end *byte
	var p *byte
	var actual_path *byte
	var wrapper *PhpStreamWrapper
	var exec_filename *types.ZendString

	/* Don't resolve paths which contain protocol (except of file://) */

	for p = filenamePtr; isalnum(int(*p)) || (*p) == '+' || (*p) == '-' || (*p) == '.'; p++ {

	}
	if (*p) == ':' && p-filenamePtr > 1 && p[1] == '/' && p[2] == '/' {
		wrapper = PhpStreamLocateUrlWrapper(filenamePtr, &actual_path, STREAM_OPEN_FOR_INCLUDE)
		if wrapper == &streams.PhpPlainFilesWrapper {
			if zend.TsrmRealpath(actual_path, resolved_path) != nil {
				return types.ZendStringInit(resolved_path, strlen(resolved_path), 0)
			}
		}
		return nil
	}
	if (*filenamePtr) == '.' && (zend.IS_SLASH(filenamePtr[1]) || filenamePtr[1] == '.' && zend.IS_SLASH(filenamePtr[2])) || zend.IS_ABSOLUTE_PATH(filenamePtr, filename_length) || path == nil || !(*path) {
		if zend.TsrmRealpath(filenamePtr, resolved_path) != nil {
			return types.ZendStringInit(resolved_path, strlen(resolved_path), 0)
		} else {
			return nil
		}
	}
	ptr = path
	for ptr != nil && (*ptr) {

		/* Check for stream wrapper */

		var is_stream_wrapper int = 0
		for p = ptr; isalnum(int(*p)) || (*p) == '+' || (*p) == '-' || (*p) == '.'; p++ {

		}
		if (*p) == ':' && p-ptr > 1 && p[1] == '/' && p[2] == '/' {

			/* .:// or ..:// is not a stream wrapper */

			if p[-1] != '.' || p[-2] != '.' || p-2 != ptr {
				p += 3
				is_stream_wrapper = 1
			}

			/* .:// or ..:// is not a stream wrapper */

		}
		end = strchr(p, zend.DEFAULT_DIR_SEPARATOR)
		if end != nil {
			if filename_length > MAXPATHLEN-2 || end-ptr > MAXPATHLEN || end-ptr+1+filename_length+1 >= MAXPATHLEN {
				ptr = end + 1
				continue
			}
			memcpy(trypath, ptr, end-ptr)
			trypath[end-ptr] = '/'
			memcpy(trypath+(end-ptr)+1, filenamePtr, filename_length+1)
			ptr = end + 1
		} else {
			var len_ int = strlen(ptr)
			if filename_length > MAXPATHLEN-2 || len_ > MAXPATHLEN || len_+1+filename_length+1 >= MAXPATHLEN {
				break
			}
			memcpy(trypath, ptr, len_)
			trypath[len_] = '/'
			memcpy(trypath+len_+1, filenamePtr, filename_length+1)
			ptr = nil
		}
		actual_path = trypath
		if is_stream_wrapper != 0 {
			wrapper = PhpStreamLocateUrlWrapper(trypath, &actual_path, STREAM_OPEN_FOR_INCLUDE)
			if wrapper == nil {
				continue
			} else if wrapper != &streams.PhpPlainFilesWrapper {
				if wrapper.GetWops().GetUrlStat() != nil {
					var ssb PhpStreamStatbuf
					if types.SUCCESS == wrapper.GetWops().GetUrlStat()(wrapper, trypath, PHP_STREAM_URL_STAT_QUIET, &ssb, nil) {
						return types.ZendStringInit(trypath, strlen(trypath), 0)
					}
					if zend.EG__().GetException() != nil {
						return nil
					}
				}
				continue
			}
		}
		if zend.TsrmRealpath(actual_path, resolved_path) != nil {
			return types.ZendStringInit(resolved_path, strlen(resolved_path), 0)
		}
	}

	/* check in calling scripts' current working directory as a fall back case
	 */

	if zend.ZendIsExecuting() != 0 && b.Assign(&exec_filename, zend.ZendGetExecutedFilenameEx()) != nil {
		var exec_fname *byte = exec_filename.GetVal()
		var exec_fname_length int = exec_filename.GetLen()
		for b.PreDec(&exec_fname_length) < SIZE_MAX && !(zend.IS_SLASH(exec_fname[exec_fname_length])) {

		}
		if exec_fname_length > 0 && filename_length < MAXPATHLEN-2 && exec_fname_length+1+filename_length+1 < MAXPATHLEN {
			memcpy(trypath, exec_fname, exec_fname_length+1)
			memcpy(trypath+exec_fname_length+1, filenamePtr, filename_length+1)
			actual_path = trypath

			/* Check for stream wrapper */

			for p = trypath; isalnum(int(*p)) || (*p) == '+' || (*p) == '-' || (*p) == '.'; p++ {

			}
			if (*p) == ':' && p-trypath > 1 && p[1] == '/' && p[2] == '/' {
				wrapper = PhpStreamLocateUrlWrapper(trypath, &actual_path, STREAM_OPEN_FOR_INCLUDE)
				if wrapper == nil {
					return nil
				} else if wrapper != &streams.PhpPlainFilesWrapper {
					if wrapper.GetWops().GetUrlStat() != nil {
						var ssb PhpStreamStatbuf
						if types.SUCCESS == wrapper.GetWops().GetUrlStat()(wrapper, trypath, PHP_STREAM_URL_STAT_QUIET, &ssb, nil) {
							return types.ZendStringInit(trypath, strlen(trypath), 0)
						}
						if zend.EG__().GetException() != nil {
							return nil
						}
					}
					return nil
				}
			}
			if zend.TsrmRealpath(actual_path, resolved_path) != nil {
				return types.ZendStringInit(resolved_path, strlen(resolved_path), 0)
			}
		}
	}
	return nil
}
func PhpFopenWithPath(filename *byte, mode string, path *byte, opened_path **types.ZendString) *r.FILE {
	var pathbuf *byte
	var ptr *byte
	var end *byte
	var trypath []byte
	var fp *r.FILE
	var filename_length int
	var exec_filename *types.ZendString
	if opened_path != nil {
		*opened_path = nil
	}
	if filename == nil {
		return nil
	}
	filename_length = strlen(filename)
	void(filename_length)

	/* Relative path open */

	if (*filename) == '.' || zend.IS_ABSOLUTE_PATH(filename, filename_length) || (path == nil || !(*path)) {
		return PhpFopenAndSetOpenedPath(filename, mode, opened_path)
	}

	/* check in provided path */

	if zend.ZendIsExecuting() != 0 && b.Assign(&exec_filename, zend.ZendGetExecutedFilenameEx()) != nil {
		var exec_fname *byte = exec_filename.GetVal()
		var exec_fname_length int = exec_filename.GetLen()
		for b.PreDec(&exec_fname_length) < SIZE_MAX && !(zend.IS_SLASH(exec_fname[exec_fname_length])) {

		}
		if exec_fname != nil && exec_fname[0] == '[' || exec_fname_length <= 0 {

			/* [no active file] or no path */

			pathbuf = zend.Estrdup(path)

			/* [no active file] or no path */

		} else {
			var path_length int = strlen(path)
			pathbuf = (*byte)(zend.Emalloc(exec_fname_length + path_length + 1 + 1))
			memcpy(pathbuf, path, path_length)
			pathbuf[path_length] = zend.DEFAULT_DIR_SEPARATOR
			memcpy(pathbuf+path_length+1, exec_fname, exec_fname_length)
			pathbuf[path_length+exec_fname_length+1] = '0'
		}
	} else {
		pathbuf = zend.Estrdup(path)
	}
	ptr = pathbuf
	for ptr != nil && (*ptr) {
		end = strchr(ptr, zend.DEFAULT_DIR_SEPARATOR)
		if end != nil {
			*end = '0'
			end++
		}
		if Snprintf(trypath, MAXPATHLEN, "%s/%s", ptr, filename) >= MAXPATHLEN {
			PhpErrorDocref(nil, zend.E_NOTICE, "%s/%s path was truncated to %d", ptr, filename, MAXPATHLEN)
		}
		fp = PhpFopenAndSetOpenedPath(trypath, mode, opened_path)
		if fp != nil {
			zend.Efree(pathbuf)
			return fp
		}
		ptr = end
	}
	zend.Efree(pathbuf)
	return nil
}
func PhpStripUrlPasswd(url *byte) *byte {
	var p *byte
	var url_start *byte
	if url == nil {
		return ""
	}
	p = url
	for *p {
		if (*p) == ':' && (*(p + 1)) == '/' && (*(p + 2)) == '/' {

			/* found protocol */

			p = p + 3
			url_start = p
			for *p {
				if (*p) == '@' {
					var i int
					for i = 0; i < 3 && url_start < p; {
						*url_start = '.'
						i++
						url_start++
					}
					for ; *p; p++ {
						b.PostInc(&(*url_start)) = *p
					}
					*url_start = 0
					break
				}
				p++
			}
			return url
		}
		p++
	}
	return url
}
func ExpandFilepath(filepath *byte, real_path *byte) *byte {
	return ExpandFilepathEx(filepath, real_path, nil, 0)
}
func ExpandFilepathEx(filepath *byte, real_path *byte, relative_to *byte, relative_to_len int) *byte {
	return ExpandFilepathWithMode(filepath, real_path, relative_to, relative_to_len, zend.CWD_FILEPATH)
}
func ExpandFilepathWithMode(filepath *byte, real_path *byte, relative_to *byte, relative_to_len int, realpath_mode int) *byte {
	var new_state zend.CwdState
	var cwd []byte
	var copy_len int
	var path_len int
	if !(filepath[0]) {
		return nil
	}
	path_len = strlen(filepath)
	if zend.IS_ABSOLUTE_PATH(filepath, path_len) {
		cwd[0] = '0'
	} else {
		var iam *byte = SG__().request_info.path_translated
		var result *byte
		if relative_to != nil {
			if relative_to_len > MAXPATHLEN-1 {
				return nil
			}
			result = relative_to
			memcpy(cwd, relative_to, relative_to_len+1)
		} else {
			result = zend.VCWD_GETCWD(cwd, MAXPATHLEN)
		}
		if result == nil && iam != filepath {
			var fdtest int = -1
			fdtest = zend.VCWD_OPEN(filepath, O_RDONLY)
			if fdtest != -1 {

				/* return a relative file path if for any reason
				 * we cannot cannot getcwd() and the requested,
				 * relatively referenced file is accessible */

				if path_len > MAXPATHLEN-1 {
					copy_len = MAXPATHLEN - 1
				} else {
					copy_len = path_len
				}
				if real_path != nil {
					memcpy(real_path, filepath, copy_len)
					real_path[copy_len] = '0'
				} else {
					real_path = zend.Estrndup(filepath, copy_len)
				}
				close(fdtest)
				return real_path
			} else {
				cwd[0] = '0'
			}
		} else if result == nil {
			cwd[0] = '0'
		}
	}
	new_state.SetCwd(zend.Estrdup(cwd))
	new_state.SetCwdLength(strlen(cwd))
	if zend.VirtualFileEx(&new_state, filepath, nil, realpath_mode) != 0 {
		zend.Efree(new_state.GetCwd())
		return nil
	}
	if real_path != nil {
		if new_state.GetCwdLength() > MAXPATHLEN-1 {
			copy_len = MAXPATHLEN - 1
		} else {
			copy_len = new_state.GetCwdLength()
		}
		memcpy(real_path, new_state.GetCwd(), copy_len)
		real_path[copy_len] = '0'
	} else {
		real_path = zend.Estrndup(new_state.GetCwd(), new_state.GetCwdLength())
	}
	zend.Efree(new_state.GetCwd())
	return real_path
}
