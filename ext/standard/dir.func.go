// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func DIRG(v __auto__) __auto__ { return DirGlobals.v }
func PhpSetDefaultDir(res *types.ZendResource) {
	if DIRG(default_dir) {
		zend.ZendListDelete(DIRG(default_dir))
	}
	if res != nil {
		res.AddRefcount()
	}
	DIRG(default_dir) = res
}
func ZmActivateDir(type_ int, module_number int) int {
	DIRG(default_dir) = nil
	return types.SUCCESS
}
func ZmStartupDir(type_ int, module_number int) int {
	var dirsep_str []byte
	var pathsep_str []byte
	var dir_class_entry zend.ZendClassEntry
	memset(&dir_class_entry, 0, b.SizeOf("zend_class_entry"))
	dir_class_entry.SetName(types.ZendStringInitInterned("Directory", b.SizeOf("\"Directory\"")-1, 1))
	dir_class_entry.SetBuiltinFunctions(PhpDirClassFunctions)
	DirClassEntryPtr = zend.ZendRegisterInternalClass(&dir_class_entry)
	dirsep_str[0] = zend.DEFAULT_SLASH
	dirsep_str[1] = '0'
	zend.REGISTER_STRING_CONSTANT("DIRECTORY_SEPARATOR", dirsep_str, zend.CONST_CS|zend.CONST_PERSISTENT)
	pathsep_str[0] = zend.ZEND_PATHS_SEPARATOR
	pathsep_str[1] = '0'
	zend.REGISTER_STRING_CONSTANT("PATH_SEPARATOR", pathsep_str, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_ASCENDING", PHP_SCANDIR_SORT_ASCENDING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_DESCENDING", PHP_SCANDIR_SORT_DESCENDING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("SCANDIR_SORT_NONE", PHP_SCANDIR_SORT_NONE, zend.CONST_CS|zend.CONST_PERSISTENT)
	const GLOB_BRACE = 0
	const GLOB_MARK = 0
	const GLOB_NOSORT = 0
	const GLOB_NOCHECK = 0
	const GLOB_NOESCAPE = 0
	const GLOB_ERR = 0
	const GLOB_ONLYDIR zend.ZendLong = 1 << 30

	// #define GLOB_EMULATE_ONLYDIR

	const GLOB_FLAGMASK = ^GLOB_ONLYDIR

	/* This is used for checking validity of passed flags (passing invalid flags causes segfault in glob()!! */

	const GLOB_AVAILABLE_FLAGS zend.ZendLong = 0 | GLOB_BRACE | GLOB_MARK | GLOB_NOSORT | GLOB_NOCHECK | GLOB_NOESCAPE | GLOB_ERR | GLOB_ONLYDIR
	zend.REGISTER_LONG_CONSTANT("GLOB_ONLYDIR", GLOB_ONLYDIR, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("GLOB_AVAILABLE_FLAGS", GLOB_AVAILABLE_FLAGS, zend.CONST_CS|zend.CONST_PERSISTENT)
	return types.SUCCESS
}
func _phpDoOpendir(executeData *zend.ZendExecuteData, return_value *types.Zval, createobject int) {
	var dirname *byte
	var dir_len int
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	var dirp *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &dirname, &dir_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_PATH
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_RESOURCE
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	context = streams.PhpStreamContextFromZval(zcontext, 0)
	dirp = core.PhpStreamOpendir(dirname, core.REPORT_ERRORS, context)
	if dirp == nil {
		return_value.SetFalse()
		return
	}
	dirp.AddFlags(core.PHP_STREAM_FLAG_NO_FCLOSE)
	PhpSetDefaultDir(dirp.GetRes())
	if createobject != 0 {
		zend.ObjectInitEx(return_value, DirClassEntryPtr)
		zend.AddPropertyStringl(return_value, "path", b.CastStrPtr(dirname, dir_len))
		zend.AddPropertyResource(return_value, "handle", dirp.GetRes())
		core.PhpStreamAutoCleanup(dirp)
	} else {
		core.PhpStreamToZval(dirp, return_value)
	}
}
func ZifOpendir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	_phpDoOpendir(executeData, return_value, 0)
}
func ZifGetdir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	_phpDoOpendir(executeData, return_value, 1)
}
func ZifClosedir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var id *types.Zval = nil
	var tmp *types.Zval
	var myself *types.Zval
	var dirp *core.PhpStream
	var res *types.ZendResource
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_RESOURCE
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, types.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				return_value.SetFalse()
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.SetFalse()
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		return_value.SetFalse()
		return
	}
	res = dirp.GetRes()
	zend.ZendListClose(dirp.GetRes())
	if res == DIRG(default_dir) {
		PhpSetDefaultDir(nil)
	}
}
func ZifChroot(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *byte
	var ret int
	var str_len int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_PATH
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	ret = chroot(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		return_value.SetFalse()
		return
	}
	PhpClearStatCache(1, nil, 0)
	ret = chdir("/")
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifChdir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var str *byte
	var ret int
	var str_len int
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &str, &str_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_PATH
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if core.PhpCheckOpenBasedir(str) != 0 {
		return_value.SetFalse()
		return
	}
	ret = zend.VCWD_CHDIR(str)
	if ret != 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s (errno %d)", strerror(errno), errno)
		return_value.SetFalse()
		return
	}
	if BG(CurrentStatFile) && !(zend.IS_ABSOLUTE_PATH(BG(CurrentStatFile), strlen(BG(CurrentStatFile)))) {
		zend.Efree(BG(CurrentStatFile))
		BG(CurrentStatFile) = nil
	}
	if BG(CurrentLStatFile) && !(zend.IS_ABSOLUTE_PATH(BG(CurrentLStatFile), strlen(BG(CurrentLStatFile)))) {
		zend.Efree(BG(CurrentLStatFile))
		BG(CurrentLStatFile) = nil
	}
	return_value.SetTrue()
	return
}
func ZifGetcwd(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var path []byte
	var ret *byte = nil
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	ret = zend.VCWD_GETCWD(path, core.MAXPATHLEN)
	if ret != nil {
		return_value.SetRawString(b.CastStrAuto(path))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func ZifRewinddir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var id *types.Zval = nil
	var tmp *types.Zval
	var myself *types.Zval
	var dirp *core.PhpStream
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_RESOURCE
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, types.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				return_value.SetFalse()
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.SetFalse()
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		return_value.SetFalse()
		return
	}
	core.PhpStreamRewinddir(dirp)
}
func PhpIfReaddir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var id *types.Zval = nil
	var tmp *types.Zval
	var myself *types.Zval
	var dirp *core.PhpStream
	var entry core.PhpStreamDirent
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 1
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &id, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_RESOURCE
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if executeData.NumArgs() == 0 {
		myself = zend.getThis()
		if myself != nil {
			if b.Assign(&tmp, types.Z_OBJPROP_P(myself).KeyFind("handle")) == nil {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to find my handle property")
				return_value.SetFalse()
				return
			}
			if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResourceEx(tmp, "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		} else {
			if !(DIRG(default_dir)) || b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(DIRG(default_dir), "Directory", streams.PhpFileLeStream()))) == nil {
				return_value.SetFalse()
				return
			}
		}
	} else {
		if b.Assign(&dirp, (*core.PhpStream)(zend.ZendFetchResource(id.GetRes(), "Directory", streams.PhpFileLeStream()))) == nil {
			return_value.SetFalse()
			return
		}
	}
	if !dirp.HasFlags(core.PHP_STREAM_FLAG_IS_DIR) {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%d is not a valid Directory resource", dirp.GetRes().GetHandle())
		return_value.SetFalse()
		return
	}
	if core.PhpStreamReaddir(dirp, &entry) != nil {
		return_value.SetRawString(entry.GetDName())
		return
	}
	return_value.SetFalse()
	return
}
func ZifGlob(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var cwd_skip int = 0
	var pattern *byte = nil
	var pattern_len int
	var flags zend.ZendLong = 0
	var globbuf glob_t
	var n int
	var ret int
	var basedir_limit types.ZendBool = 0
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &pattern, &pattern_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_PATH
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &flags) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if pattern_len >= core.MAXPATHLEN {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Pattern exceeds the maximum allowed length of %d characters", core.MAXPATHLEN)
		return_value.SetFalse()
		return
	}
	if (GLOB_AVAILABLE_FLAGS & flags) != flags {
		core.PhpErrorDocref(nil, zend.E_WARNING, "At least one of the passed flags is invalid or not supported on this platform")
		return_value.SetFalse()
		return
	}
	memset(&globbuf, 0, b.SizeOf("glob_t"))
	globbuf.gl_offs = 0
	if 0 != b.Assign(&ret, glob(pattern, flags&streams.GLOB_FLAGMASK, nil, &globbuf)) {
		return_value.SetFalse()
		return
	}

	/* now catch the FreeBSD style of "no matches" */

	if !(globbuf.gl_pathc) || !(globbuf.gl_pathv) {

		/* Paths containing '*', '?' and some other chars are
		   illegal on Windows but legit on other platforms. For
		   this reason the direct basedir check against the glob
		   query is senseless on windows. For instance while *.txt
		   is a pretty valid filename on EXT3, it's invalid on NTFS. */

		if core.PG(open_basedir) && (*core.PG)(open_basedir) {
			if core.PhpCheckOpenBasedirEx(pattern, 0) != 0 {
				return_value.SetFalse()
				return
			}
		}
		zend.ArrayInit(return_value)
		return
	}
	zend.ArrayInit(return_value)
	for n = 0; n < int(globbuf.gl_pathc); n++ {
		if core.PG(open_basedir) && (*core.PG)(open_basedir) {
			if core.PhpCheckOpenBasedirEx(globbuf.gl_pathv[n], 0) != 0 {
				basedir_limit = 1
				continue
			}
		}

		/* we need to do this every time since GLOB_ONLYDIR does not guarantee that
		 * all directories will be filtered. GNU libc documentation states the
		 * following:
		 * If the information about the type of the file is easily available
		 * non-directories will be rejected but no extra work will be done to
		 * determine the information for each file. I.e., the caller must still be
		 * able to filter directories out.
		 */

		if (flags & streams.GLOB_ONLYDIR) != 0 {
			var s zend.ZendStatT
			if 0 != zend.VCWD_STAT(globbuf.gl_pathv[n], &s) {
				continue
			}
			if S_IFDIR != (s.st_mode & S_IFMT) {
				continue
			}
		}
		zend.AddNextIndexString(return_value, globbuf.gl_pathv[n]+cwd_skip)
	}
	globfree(&globbuf)
	if basedir_limit != 0 && !(types.Z_ARRVAL_P(return_value).GetNNumOfElements()) {
		return_value.GetArr().DestroyEx()
		return_value.SetFalse()
		return
	}
}
func ZifScandir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var dirn *byte
	var dirn_len int
	var flags zend.ZendLong = 0
	var namelist **types.ZendString
	var n int
	var i int
	var zcontext *types.Zval = nil
	var context *core.PhpStreamContext = nil
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgPath(_arg, &dirn, &dirn_len, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_PATH
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			fp.StartOptional()
			zend.Z_PARAM_PROLOGUE(0, 0)
			if !zend.ZendParseArgLong00(_arg, &flags) {
				_expected_type = argparse.Z_EXPECTED_LONG
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			zend.Z_PARAM_PROLOGUE(0, 0)
			if zend.ZendParseArgResource(_arg, &zcontext, 0) == 0 {
				_expected_type = argparse.Z_EXPECTED_RESOURCE
				_error_code = argparse.ZPP_ERROR_WRONG_ARG
				break
			}
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	if dirn_len < 1 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Directory name cannot be empty")
		return_value.SetFalse()
		return
	}
	if zcontext != nil {
		context = streams.PhpStreamContextFromZval(zcontext, 0)
	}
	if flags == PHP_SCANDIR_SORT_ASCENDING {
		n = core.PhpStreamScandir(dirn, &namelist, context, any(streams.PhpStreamDirentAlphasort))
	} else if flags == PHP_SCANDIR_SORT_NONE {
		n = core.PhpStreamScandir(dirn, &namelist, context, nil)
	} else {
		n = core.PhpStreamScandir(dirn, &namelist, context, any(streams.PhpStreamDirentAlphasortr))
	}
	if n < 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "(errno %d): %s", errno, strerror(errno))
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	for i = 0; i < n; i++ {
		zend.AddNextIndexStr(return_value, namelist[i])
	}
	if n != 0 {
		zend.Efree(namelist)
	}
}
