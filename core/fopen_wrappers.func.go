package core

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core/streams"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"os"
	"strings"
)

func OnUpdateBaseDir(
	entry *zend.ZendIniEntry,
	new_value *types.String,
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
		if ptr[0] == '.' && ptr[1] == '.' && (ptr[2] == '0' || zend.IsSlash(ptr[2])) {

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
				resolved_basedir[lang.PreInc(&resolved_basedir_len)] = '0'
			}
		} else {
			resolved_basedir[lang.PostInc(&resolved_basedir_len)] = PHP_DIR_SEPARATOR
			resolved_basedir[resolved_basedir_len] = '0'
		}
		resolved_name_len = strlen(resolved_name)
		if path_tmp[path_len-1] == PHP_DIR_SEPARATOR {
			if resolved_name[resolved_name_len-1] != PHP_DIR_SEPARATOR {
				resolved_name[resolved_name_len] = PHP_DIR_SEPARATOR
				resolved_name[lang.PreInc(&resolved_name_len)] = '0'
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
			PhpErrorDocref("", faults.E_WARNING, "File name is longer than the maximum allowed path length on this platform (%d): %s", MAXPATHLEN, path)
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
			PhpErrorDocref("", faults.E_WARNING, "open_basedir restriction in effect. File(%s) is not within the allowed path(s): (%s)", path, PG__().open_basedir)
		}
		zend.Efree(pathbuf)
		errno = EPERM
		return -1
	}

	/* Nothing to check... */

	return 0

	/* Nothing to check... */
}
func PhpFopenAndSetOpenedPath(path string, mode string, opened_path **types.String) *r.File {
	var fp *r.File
	if PhpCheckOpenBasedir((*byte)(path)) != 0 {
		return nil
	}
	fp = zend.VCWD_FOPEN(path, mode)
	if fp != nil && opened_path != nil {
		//TODO :avoid reallocation
		var tmp *byte = ExpandFilepathWithMode(path, nil, nil, 0, zend.CWD_EXPAND)
		if tmp != nil {
			*opened_path = types.NewString(tmp)
			zend.Efree(tmp)
		}
	}
	return fp
}
func PhpFopenPrimaryScript() *zend.FileHandle {
	var path_info *byte
	var filename *byte = nil
	var resolved_path *string = nil
	var length int
	var orig_display_errors bool
	path_info = SG__().RequestInfo.requestUri
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
				filename = fmt.Sprintf("%s%c%s%c%s", pw.pw_dir, PHP_DIR_SEPARATOR, PG__().user_dir, PHP_DIR_SEPARATOR, s+1)
			} else {
				filename = SG__().RequestInfo.pathTranslated
			}
		}
	} else if PG__().doc_root && path_info != nil && lang.Assign(&length, strlen(PG__().doc_root)) && zend.IsAbsolutePathOld(PG__().doc_root, length) {
		var path_len int = strlen(path_info)
		filename = zend.Emalloc(length + path_len + 2)
		memcpy(filename, PG__().doc_root, length)
		if !(zend.IsSlash(filename[length-1])) {
			filename[lang.PostInc(&length)] = PHP_DIR_SEPARATOR
		}
		if zend.IsSlash(path_info[0]) {
			length--
		}
		strncpy(filename+length, path_info, path_len+1)
	} else {
		filename = SG__().RequestInfo.pathTranslated
	}
	if filename != nil {
		resolved_path = PhpResolvePathForZend(filename)
	}
	if resolved_path == nil {
		if SG__().RequestInfo.pathTranslated != filename {
			if filename != nil {
				zend.Efree(filename)
			}
		}

		/* we have to free SG__().RequestInfo.path_translated here because
		 * php_destroy_request_info assumes that it will get
		 * freed when the include_names hash is emptied, but
		 * we're not adding it in this case */

		if SG__().RequestInfo.pathTranslated {
			zend.Efree(SG__().RequestInfo.pathTranslated)
			SG__().RequestInfo.pathTranslated = nil
		}
		return types.FAILURE
	}
	orig_display_errors = PG__().display_errors
	PG__().display_errors = 0
	fileHandle := zend.NewFileHandleByOpenStream(filename)
	if fileHandle == nil {
		PG__().display_errors = orig_display_errors
		if SG__().RequestInfo.pathTranslated != filename {
			if filename != nil {
				zend.Efree(filename)
			}
		}
		if SG__().RequestInfo.pathTranslated {
			zend.Efree(SG__().RequestInfo.pathTranslated)
			SG__().RequestInfo.pathTranslated = nil
		}
		return nil
	}
	PG__().display_errors = orig_display_errors
	if SG__().RequestInfo.pathTranslated != filename {
		if SG__().RequestInfo.pathTranslated {
			zend.Efree(SG__().RequestInfo.pathTranslated)
		}
		SG__().RequestInfo.pathTranslated = filename
	}
	return fileHandle
}
func parseProtocol(name string) (protocol string, remain string, ok bool) {
	idx := 0
	for idx < len(name) && (ascii.IsAlphaNum(name[idx]) || name[idx] == '+' || name[idx] == '-' || name[idx] == '.') {
		idx++
	}
	if idx > 1 && strings.HasPrefix(name[idx:], "://") {
		return name[:idx], name[idx+3:], true
	}
	return "", name, false
}

func PhpResolvePath(filename string, filenamePtr *byte, filename_length int, path string) *types.String {
	var resolved_path []byte
	var exec_filename *types.String

	if filename == "" {
		return nil
	}

	/* Don't resolve paths which contain protocol (except of file://) */
	if _, _, ok := parseProtocol(filename); ok {
		var actual_path string
		wrapper := PhpStreamLocateUrlWrapper(filename, &actual_path, STREAM_OPEN_FOR_INCLUDE)
		if wrapper == &streams.PhpPlainFilesWrapper {
			if zend.TsrmRealpath(actual_path, resolved_path) != nil {
				return types.NewString(resolved_path)
			}
		}
		return nil
	}

	var direct = false
	if len(filename) > 2 && filename[0] == '.' && zend.IsSlash(filename[1]) {
		direct = true
	} else if len(filename) > 3 && filename[:2] == ".." && zend.IsSlash(filename[2]) {
		direct = true
	} else if zend.IsAbsolutePath(filename) {
		direct = true
	} else if path == "" {
		direct = true
	}
	if direct {
		if zend.TsrmRealpath(filenamePtr, resolved_path) != nil {
			return types.NewString(resolved_path)
		} else {
			return nil
		}
	}

	for _, onePath := range strings.Split(path, zend.DEFAULT_DIR_SEPARATOR) {
		/* Check for stream wrapper */
		streamWrapper := false
		if protocol, remain, ok := parseProtocol(path); ok {
			/* .:// or ..:// is not a stream wrapper */
			if protocol != ".." {
				streamWrapper = true
				onePath = remain
			}
		}

		if len(filename)+len(onePath)+2 >= MAXPATHLEN {
			continue
		}
		trypath := onePath + "/" + filename
		actual_path := trypath
		if streamWrapper {
			wrapper := PhpStreamLocateUrlWrapper(trypath, &actual_path, STREAM_OPEN_FOR_INCLUDE)
			if wrapper == nil {
				continue
			} else if wrapper != &streams.PhpPlainFilesWrapper {
				if wrapper.GetWops().GetUrlStat() != nil {
					var ssb PhpStreamStatbuf
					if types.SUCCESS == wrapper.GetWops().GetUrlStat()(wrapper, trypath, PHP_STREAM_URL_STAT_QUIET, &ssb, nil) {
						return types.NewString(trypath)
					}
					if zend.EG__().HasException() {
						return nil
					}
				}
				continue
			}
		}
		if zend.TsrmRealpath(actual_path, resolved_path) != nil {
			return types.NewString(resolved_path)
		}
	}

	/* check in calling scripts' current working directory as a fall back case */
	if zend.ZendIsExecuting() {
		execDirname, _ := zend.CutPath(zend.ZendGetExecutedFilenameEx())
		if execDirname != "" && len(execDirname)+len(filename)+2 < MAXPATHLEN {
			trypath := execDirname + "/" + filename
			actual_path := trypath

			/* Check for stream wrapper */
			if _, _, ok := parseProtocol(trypath); ok {
				wrapper := PhpStreamLocateUrlWrapper(trypath, &actual_path, STREAM_OPEN_FOR_INCLUDE)
				if wrapper == nil {
					return nil
				} else if wrapper != &streams.PhpPlainFilesWrapper {
					if wrapper.GetWops().GetUrlStat() != nil {
						var ssb PhpStreamStatbuf
						if types.SUCCESS == wrapper.GetWops().GetUrlStat()(wrapper, trypath, PHP_STREAM_URL_STAT_QUIET, &ssb, nil) {
							return types.NewString(trypath)
						}
						if zend.EG__().HasException() {
							return nil
						}
					}
					return nil
				}
			}
			if zend.TsrmRealpath(actual_path, resolved_path) != nil {
				return types.NewString(resolved_path)
			}
		}
	}
	return nil
}
func PhpFopenWithPath(filename string, mode string, path string, openedPath **types.String) *r.File {
	var fp *r.File
	if openedPath != nil {
		*openedPath = nil
	}
	if filename == "" {
		return nil
	}

	/* Relative path open */
	if filename[0] == '.' || zend.IsAbsolutePath(filename) || path == "" {
		return PhpFopenAndSetOpenedPath(filename, mode, openedPath)
	}

	/* check in provided path */
	if zend.ZendIsExecuting() {
		execDirname, _ := zend.CutPath(zend.ZendGetExecutedFilenameEx())
		if execDirname != "" {
			path = path + string(zend.DEFAULT_DIR_SEPARATOR) + execDirname
		}
	}

	for _, onePath := range strings.Split(path, string(zend.DEFAULT_DIR_SEPARATOR)) {
		trypath := onePath + "/" + filename
		if len(trypath) >= MAXPATHLEN {
			trypath = trypath[:MAXPATHLEN]
			PhpErrorDocref("", faults.E_NOTICE, "%s/%s path was truncated to %d", onePath, filename, MAXPATHLEN)
		}
		fp = PhpFopenAndSetOpenedPath(trypath, mode, openedPath)
		if fp != nil {
			return fp
		}
	}
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
						lang.PostInc(&(*url_start)) = *p
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
	if zend.IsAbsolutePathOld(filepath, path_len) {
		cwd[0] = '0'
	} else {
		var iam *byte = SG__().RequestInfo.pathTranslated
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
			if fp, err := os.OpenFile(filepath, os.O_RDONLY, 0); err != nil {
				defer fp.Close()

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
