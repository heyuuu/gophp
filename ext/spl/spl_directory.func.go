// <<generate>>

package spl

import (
	b "sik/builtin"
	r "sik/builtin/file"
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	"sik/sapi/cli"
	"sik/zend"
	"sik/zend/faults"
	"sik/zend/types"
)

func SplFilesystemFromObj(obj *types.ZendObject) *SplFilesystemObject {
	return (*SplFilesystemObject)((*byte)(obj - zend_long((*byte)(&((*SplFilesystemObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLFILESYSTEM_P(zv *types.Zval) *SplFilesystemObject { return SplFilesystemFromObj(zv.GetObj()) }
func SplFilesystemObjectToIterator(obj *SplFilesystemObject) *SplFilesystemIterator {
	var it *SplFilesystemIterator
	it = zend.Ecalloc(1, b.SizeOf("spl_filesystem_iterator"))
	it.SetObject(any(obj))
	zend.ZendIteratorInit(it.GetIntern())
	return it
}
func SplFilesystemIteratorToObject(it *SplFilesystemIterator) *SplFilesystemObject {
	return (*SplFilesystemObject)(it.GetObject())
}
func SPL_FILE_DIR_CURRENT(intern *SplFilesystemObject, mode __auto__) bool {
	return (intern.GetFlags() & SPL_FILE_DIR_CURRENT_MODE_MASK) == mode
}
func SPL_FILE_DIR_KEY(intern *SplFilesystemObject, mode __auto__) bool {
	return (intern.GetFlags() & SPL_FILE_DIR_KEY_MODE_MASK) == mode
}
func SPL_HAS_FLAG(flags zend.ZendLong, test_flag zend.ZendLong) int {
	if (flags & test_flag) != 0 {
		return 1
	} else {
		return 0
	}
}
func SplFilesystemFileFreeLine(intern *SplFilesystemObject) {
	if intern.GetCurrentLine() != nil {
		zend.Efree(intern.GetCurrentLine())
		intern.SetCurrentLine(nil)
	}
	if !(intern.GetCurrentZval().IsUndef()) {
		zend.ZvalPtrDtor(intern.GetCurrentZval())
		intern.GetCurrentZval().SetUndef()
	}
}
func SplFilesystemObjectDestroyObject(object *types.ZendObject) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(object)
	zend.ZendObjectsDestroyObject(object)
	switch intern.GetType() {
	case SPL_FS_DIR:
		if intern.GetDirp() != nil {
			core.PhpStreamClose(intern.GetDirp())
			intern.SetDirp(nil)
		}
	case SPL_FS_FILE:
		if intern.GetStream() != nil {

			/*
			   if (intern->u.file.zcontext) {
			      zend_list_delref(Z_RESVAL_P(intern->zcontext));
			   }
			*/

			if intern.GetStream().GetIsPersistent() == 0 {
				core.PhpStreamClose(intern.GetStream())
			} else {
				core.PhpStreamPclose(intern.GetStream())
			}
			intern.SetStream(nil)
			intern.GetZresource().SetUndef()
		}
	default:

	}
}
func SplFilesystemObjectFreeStorage(object *types.ZendObject) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(object)
	if intern.GetOthHandler() != nil && intern.GetOthHandler().GetDtor() != nil {
		intern.GetOthHandler().GetDtor()(intern)
	}
	zend.ZendObjectStdDtor(intern.GetStd())
	if intern.GetPath() != nil {
		zend.Efree(intern.GetPath())
	}
	if intern.GetFileName() != nil {
		zend.Efree(intern.GetFileName())
	}
	switch intern.GetType() {
	case SPL_FS_INFO:

	case SPL_FS_DIR:
		if intern.GetSubPath() != nil {
			zend.Efree(intern.GetSubPath())
		}
	case SPL_FS_FILE:
		if intern.GetOpenMode() != nil {
			zend.Efree(intern.GetOpenMode())
		}
		if intern.GetOrigPath() != nil {
			zend.Efree(intern.GetOrigPath())
		}
		SplFilesystemFileFreeLine(intern)
	}
}
func SplFilesystemObjectNewEx(class_type *types.ClassEntry) *types.ZendObject {
	var intern *SplFilesystemObject
	intern = zend.ZendObjectAlloc(b.SizeOf("spl_filesystem_object"), class_type)

	/* intern->type = SPL_FS_INFO; done by set 0 */

	intern.SetFileClass(spl_ce_SplFileObject)
	intern.SetInfoClass(spl_ce_SplFileInfo)
	zend.ZendObjectStdInit(intern.GetStd(), class_type)
	zend.ObjectPropertiesInit(intern.GetStd(), class_type)
	intern.GetStd().SetHandlers(&SplFilesystemObjectHandlers)
	return intern.GetStd()
}
func SplFilesystemObjectNew(class_type *types.ClassEntry) *types.ZendObject {
	return SplFilesystemObjectNewEx(class_type)
}
func SplFilesystemObjectNewCheck(class_type *types.ClassEntry) *types.ZendObject {
	var ret *SplFilesystemObject = SplFilesystemFromObj(SplFilesystemObjectNewEx(class_type))
	ret.GetStd().SetHandlers(&SplFilesystemObjectCheckHandlers)
	return ret.GetStd()
}
func SplFilesystemObjectGetPath(intern *SplFilesystemObject, len_ *int) *byte {
	if intern.GetType() == SPL_FS_DIR {
		if core.PhpStreamIs(intern.GetDirp(), &streams.PhpGlobStreamOps) {
			return streams.PhpGlobStreamGetPath(intern.GetDirp(), len_)
		}
	}
	if len_ != nil {
		*len_ = intern.GetPathLen()
	}
	return intern.GetPath()
}
func SplFilesystemObjectGetFileName(intern *SplFilesystemObject) {
	var slash byte = b.Cond(SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_UNIXPATHS) != 0, '/', zend.DEFAULT_SLASH)
	switch intern.GetType() {
	case SPL_FS_INFO:
		fallthrough
	case SPL_FS_FILE:
		if intern.GetFileName() == nil {
			core.PhpErrorDocref(nil, faults.E_ERROR, "Object not initialized")
		}
	case SPL_FS_DIR:
		var path_len int = 0
		var path *byte = SplFilesystemObjectGetPath(intern, &path_len)
		if intern.GetFileName() != nil {
			zend.Efree(intern.GetFileName())
		}

		/* if there is parent path, ammend it, otherwise just use the given path as is */

		if path_len == 0 {
			intern.SetFileNameLen(core.Spprintf(intern.GetFileName(), 0, "%s", intern.GetEntry().GetDName()))
		} else {
			intern.SetFileNameLen(core.Spprintf(intern.GetFileName(), 0, "%s%c%s", path, slash, intern.GetEntry().GetDName()))
		}

		/* if there is parent path, ammend it, otherwise just use the given path as is */

	}
}
func SplFilesystemDirRead(intern *SplFilesystemObject) int {
	if intern.GetDirp() == nil || core.PhpStreamReaddir(intern.GetDirp(), intern.GetEntry()) == nil {
		intern.GetEntry().GetDName()[0] = '0'
		return 0
	} else {
		return 1
	}
}
func IS_SLASH_AT(zs []byte, pos int) bool { return zend.IS_SLASH(zs[pos]) }
func SplFilesystemIsDot(d_name *byte) int {
	return !(strcmp(d_name, ".")) || !(strcmp(d_name, ".."))
}
func SplFilesystemDirOpen(intern *SplFilesystemObject, path *byte) {
	var skip_dots int = SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
	intern.SetType(SPL_FS_DIR)
	intern.SetPathLen(strlen(path))
	intern.SetDirp(core.PhpStreamOpendir(path, core.REPORT_ERRORS, standard.FG(default_context)))
	if intern.GetPathLen() > 1 && IS_SLASH_AT(path, intern.GetPathLen()-1) {
		intern.SetPath(zend.Estrndup(path, b.PreDec(&(intern.GetPathLen()))))
	} else {
		intern.SetPath(zend.Estrndup(path, intern.GetPathLen()))
	}
	intern.SetIndex(0)
	if zend.EG__().GetException() != nil || intern.GetDirp() == nil {
		intern.GetEntry().GetDName()[0] = '0'
		if zend.EG__().GetException() == nil {

			/* open failed w/out notice (turned to exception due to EH_THROW) */

			faults.ThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Failed to open directory \"%s\"", path)

			/* open failed w/out notice (turned to exception due to EH_THROW) */

		}
	} else {
		for {
			SplFilesystemDirRead(intern)
			if !(skip_dots != 0 && SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0) {
				break
			}
		}
	}
}
func SplFilesystemFileOpen(intern *SplFilesystemObject, use_include_path int, silent int) int {
	var tmp types.Zval
	intern.SetType(SPL_FS_FILE)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, &tmp)
	if tmp.IsType(types.IS_TRUE) {
		intern.SetOpenMode(nil)
		intern.SetFileName(nil)
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Cannot use SplFileObject with directories")
		return types.FAILURE
	}
	intern.SetContext(streams.PhpStreamContextFromZval(intern.GetZcontext(), 0))
	intern.SetStream(core.PhpStreamOpenWrapperEx(intern.GetFileName(), intern.GetOpenMode(), b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, intern.GetContext()))
	if intern.GetFileNameLen() == 0 || intern.GetStream() == nil {
		if zend.EG__().GetException() == nil {
			faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot open file '%s'", b.CondF1(intern.GetFileNameLen() != 0, func() *byte { return intern.GetFileName() }, ""))
		}
		intern.SetFileName(nil)
		intern.SetOpenMode(nil)
		return types.FAILURE
	}

	/*
	   if (intern->u.file.zcontext) {
	       //zend_list_addref(Z_RES_VAL(intern->u.file.zcontext));
	       Z_ADDREF_P(intern->u.file.zcontext);
	   }
	*/

	if intern.GetFileNameLen() > 1 && IS_SLASH_AT(intern.GetFileName(), intern.GetFileNameLen()-1) {
		intern.GetFileNameLen()--
	}
	intern.SetOrigPath(zend.Estrndup(intern.GetStream().GetOrigPath(), strlen(intern.GetStream().GetOrigPath())))
	intern.SetFileName(zend.Estrndup(intern.GetFileName(), intern.GetFileNameLen()))
	intern.SetOpenMode(zend.Estrndup(intern.GetOpenMode(), intern.GetOpenModeLen()))

	/* avoid reference counting in debug mode, thus do it manually */

	intern.GetZresource().SetResource(intern.GetStream().GetRes())

	/*!!! TODO: maybe bug?
	  Z_SET_REFCOUNT(intern->u.file.zresource, 1);
	*/

	intern.SetDelimiter(',')
	intern.SetEnclosure('"')
	intern.SetEscape(uint8('\\'))
	intern.SetFuncGetCurr(types.ZendHashStrFindPtr(intern.GetStd().GetCe().GetFunctionTable(), "getcurrentline", b.SizeOf("\"getcurrentline\"")-1))
	return types.SUCCESS
}
func SplFilesystemObjectClone(zobject *types.Zval) *types.ZendObject {
	var old_object *types.ZendObject
	var new_object *types.ZendObject
	var intern *SplFilesystemObject
	var source *SplFilesystemObject
	var index int
	var skip_dots int
	old_object = zobject.GetObj()
	source = SplFilesystemFromObj(old_object)
	new_object = SplFilesystemObjectNewEx(old_object.GetCe())
	intern = SplFilesystemFromObj(new_object)
	intern.SetFlags(source.GetFlags())
	switch source.GetType() {
	case SPL_FS_INFO:
		intern.SetPathLen(source.GetPathLen())
		intern.SetPath(zend.Estrndup(source.GetPath(), source.GetPathLen()))
		intern.SetFileNameLen(source.GetFileNameLen())
		intern.SetFileName(zend.Estrndup(source.GetFileName(), intern.GetFileNameLen()))
	case SPL_FS_DIR:
		SplFilesystemDirOpen(intern, source.GetPath())

		/* read until we hit the position in which we were before */

		skip_dots = SPL_HAS_FLAG(source.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
		for index = 0; index < source.GetIndex(); index++ {
			for {
				SplFilesystemDirRead(intern)
				if !(skip_dots != 0 && SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0) {
					break
				}
			}
		}
		intern.SetIndex(index)
	case SPL_FS_FILE:
		b.Assert(false)
	}
	intern.SetFileClass(source.GetFileClass())
	intern.SetInfoClass(source.GetInfoClass())
	intern.SetOth(source.GetOth())
	intern.SetOthHandler(source.GetOthHandler())
	zend.ZendObjectsCloneMembers(new_object, old_object)
	if intern.GetOthHandler() != nil && intern.GetOthHandler().GetClone() != nil {
		intern.GetOthHandler().GetClone()(source, intern)
	}
	return new_object
}
func SplFilesystemInfoSetFilename(intern *SplFilesystemObject, path *byte, len_ int, use_copy int) {
	var p1 *byte
	var p2 *byte
	if intern.GetFileName() != nil {
		zend.Efree(intern.GetFileName())
	}
	if use_copy != 0 {
		intern.SetFileName(zend.Estrndup(path, len_))
	} else {
		intern.SetFileName(path)
	}
	intern.SetFileNameLen(len_)
	for intern.GetFileNameLen() > 1 && IS_SLASH_AT(intern.GetFileName(), intern.GetFileNameLen()-1) {
		intern.GetFileName()[intern.GetFileNameLen()-1] = 0
		intern.GetFileNameLen()--
	}
	p1 = strrchr(intern.GetFileName(), '/')
	p2 = 0
	if p1 != nil || p2 != nil {
		intern.SetPathLen(b.Cond(p1 > p2, p1, p2) - intern.GetFileName())
	} else {
		intern.SetPathLen(0)
	}
	if intern.GetPath() != nil {
		zend.Efree(intern.GetPath())
	}
	intern.SetPath(zend.Estrndup(path, intern.GetPathLen()))
}
func SplFilesystemObjectCreateInfo(
	source *SplFilesystemObject,
	file_path *byte,
	file_path_len int,
	use_copy int,
	ce *types.ClassEntry,
	return_value *types.Zval,
) *SplFilesystemObject {
	var intern *SplFilesystemObject
	var arg1 types.Zval
	var error_handling zend.ZendErrorHandling
	if file_path == nil || file_path_len == 0 {
		if file_path != nil && use_copy == 0 {
			zend.Efree(file_path)
		}
		file_path_len = 1
		file_path = "/"
		return nil
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if ce != nil {
		ce = ce
	} else {
		ce = source.GetInfoClass()
	}
	zend.ZendUpdateClassConstants(ce)
	intern = SplFilesystemFromObj(SplFilesystemObjectNewEx(ce))
	return_value.SetObject(intern.GetStd())
	if ce.GetConstructor().GetScope() != spl_ce_SplFileInfo {
		arg1.SetRawString(b.CastStr(file_path, file_path_len))
		zend.ZendCallMethodWith1Params(return_value, ce, ce.GetConstructor(), "__construct", nil, &arg1)
		zend.ZvalPtrDtor(&arg1)
	} else {
		SplFilesystemInfoSetFilename(intern, file_path, file_path_len, use_copy)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
	return intern
}
func SplFilesystemObjectCreateType(ht int, source *SplFilesystemObject, type_ int, ce *types.ClassEntry, return_value *types.Zval) *SplFilesystemObject {
	var intern *SplFilesystemObject
	var use_include_path types.ZendBool = 0
	var arg1 types.Zval
	var arg2 types.Zval
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	switch source.GetType() {
	case SPL_FS_INFO:
		fallthrough
	case SPL_FS_FILE:

	case SPL_FS_DIR:
		if !(source.GetEntry().GetDName()[0]) {
			faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Could not open file")
			zend.ZendRestoreErrorHandling(&error_handling)
			return nil
		}
	}
	switch type_ {
	case SPL_FS_INFO:
		if ce != nil {
			ce = ce
		} else {
			ce = source.GetInfoClass()
		}
		if zend.ZendUpdateClassConstants(ce) != types.SUCCESS {
			break
		}
		intern = SplFilesystemFromObj(SplFilesystemObjectNewEx(ce))
		return_value.SetObject(intern.GetStd())
		SplFilesystemObjectGetFileName(source)
		if ce.GetConstructor().GetScope() != spl_ce_SplFileInfo {
			arg1.SetRawString(b.CastStr(source.GetFileName(), source.GetFileNameLen()))
			zend.ZendCallMethodWith1Params(return_value, ce, ce.GetConstructor(), "__construct", nil, &arg1)
			zend.ZvalPtrDtor(&arg1)
		} else {
			intern.SetFileName(zend.Estrndup(source.GetFileName(), source.GetFileNameLen()))
			intern.SetFileNameLen(source.GetFileNameLen())
			intern.SetPath(SplFilesystemObjectGetPath(source, intern.GetPathLen()))
			intern.SetPath(zend.Estrndup(intern.GetPath(), intern.GetPathLen()))
		}
	case SPL_FS_FILE:
		if ce != nil {
			ce = ce
		} else {
			ce = source.GetFileClass()
		}
		if zend.ZendUpdateClassConstants(ce) != types.SUCCESS {
			break
		}
		intern = SplFilesystemFromObj(SplFilesystemObjectNewEx(ce))
		return_value.SetObject(intern.GetStd())
		SplFilesystemObjectGetFileName(source)
		if ce.GetConstructor().GetScope() != spl_ce_SplFileObject {
			arg1.SetRawString(b.CastStr(source.GetFileName(), source.GetFileNameLen()))
			arg2.SetRawString("r")
			zend.ZendCallMethodWith2Params(return_value, ce, ce.GetConstructor(), "__construct", nil, &arg1, &arg2)
			zend.ZvalPtrDtor(&arg1)
			zend.ZvalPtrDtor(&arg2)
		} else {
			intern.SetFileName(source.GetFileName())
			intern.SetFileNameLen(source.GetFileNameLen())
			intern.SetPath(SplFilesystemObjectGetPath(source, intern.GetPathLen()))
			intern.SetPath(zend.Estrndup(intern.GetPath(), intern.GetPathLen()))
			intern.SetOpenMode("r")
			intern.SetOpenModeLen(1)
			if ht != 0 && zend.ZendParseParameters(ht, "|sbr", intern.GetOpenMode(), intern.GetOpenModeLen(), &use_include_path, intern.GetZcontext()) == types.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				intern.SetOpenMode(nil)
				intern.SetFileName(nil)
				zend.ZvalPtrDtor(return_value)
				return_value.SetNull()
				return nil
			}
			if SplFilesystemFileOpen(intern, use_include_path, 0) == types.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				zend.ZvalPtrDtor(return_value)
				return_value.SetNull()
				return nil
			}
		}
	case SPL_FS_DIR:
		zend.ZendRestoreErrorHandling(&error_handling)
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Operation not supported")
		return nil
	}
	zend.ZendRestoreErrorHandling(&error_handling)
	return nil
}
func SplFilesystemIsInvalidOrDot(d_name *byte) int {
	return d_name[0] == '0' || SplFilesystemIsDot(d_name) != 0
}
func SplFilesystemObjectGetPathname(intern *SplFilesystemObject, len_ *int) *byte {
	switch intern.GetType() {
	case SPL_FS_INFO:
		fallthrough
	case SPL_FS_FILE:
		*len_ = intern.GetFileNameLen()
		return intern.GetFileName()
	case SPL_FS_DIR:
		if intern.GetEntry().GetDName()[0] {
			SplFilesystemObjectGetFileName(intern)
			*len_ = intern.GetFileNameLen()
			return intern.GetFileName()
		}
	}
	*len_ = 0
	return nil
}
func SplFilesystemObjectGetDebugInfo(object *types.Zval) *types.HashTable {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(object)
	var tmp types.Zval
	var rv *types.HashTable
	var pnstr *types.ZendString
	var path *byte
	var path_len int
	var stmp []byte
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	rv = types.ZendArrayDup(intern.GetStd().GetProperties())
	pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "pathName")
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	tmp.SetRawString(b.CastStr(b.Cond(path != nil, path, ""), path_len))
	rv.SymtableUpdate(pnstr.GetStr(), &tmp)
	types.ZendStringReleaseEx(pnstr, 0)
	if intern.GetFileName() != nil {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "fileName")
		SplFilesystemObjectGetPath(intern, &path_len)
		if path_len != 0 && path_len < intern.GetFileNameLen() {
			tmp.SetRawString(b.CastStr(intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1)))
		} else {
			tmp.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
		}
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_DIR {
		pnstr = SplGenPrivatePropName(spl_ce_DirectoryIterator, "glob")
		if core.PhpStreamIs(intern.GetDirp(), &streams.PhpGlobStreamOps) {
			tmp.SetRawString(b.CastStr(intern.GetPath(), intern.GetPathLen()))
		} else {
			tmp.SetFalse()
		}
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
		pnstr = SplGenPrivatePropName(spl_ce_RecursiveDirectoryIterator, "subPathName")
		if intern.GetSubPath() != nil {
			tmp.SetRawString(b.CastStr(intern.GetSubPath(), intern.GetSubPathLen()))
		} else {
			zend.ZVAL_EMPTY_STRING(&tmp)
		}
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_FILE {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "openMode")
		tmp.SetRawString(b.CastStr(intern.GetOpenMode(), intern.GetOpenModeLen()))
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
		stmp[1] = '0'
		stmp[0] = intern.GetDelimiter()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "delimiter")
		tmp.SetRawString(b.CastStr(stmp, 1))
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
		stmp[0] = intern.GetEnclosure()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "enclosure")
		tmp.SetRawString(b.CastStr(stmp, 1))
		rv.SymtableUpdate(pnstr.GetStr(), &tmp)
		types.ZendStringReleaseEx(pnstr, 0)
	}
	return rv
}
func SplFilesystemObjectGetMethodCheck(object **types.ZendObject, method *types.ZendString, key *types.Zval) *zend.ZendFunction {
	var fsobj *SplFilesystemObject = SplFilesystemFromObj(*object)
	if fsobj.GetDirp() == nil && fsobj.GetOrigPath() == nil {
		var func_ *zend.ZendFunction
		var tmp *types.ZendString = types.ZendStringInit("_bad_state_ex", b.SizeOf("\"_bad_state_ex\"")-1, 0)
		func_ = zend.ZendStdGetMethod(object, tmp, nil)
		types.ZendStringReleaseEx(tmp, 0)
		return func_
	}
	return zend.ZendStdGetMethod(object, method, key)
}
func SplFilesystemObjectConstruct(executeData *zend.ZendExecuteData, return_value *types.Zval, ctor_flags zend.ZendLong) {
	var intern *SplFilesystemObject
	var path *byte
	var parsed int
	var len_ int
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if SPL_HAS_FLAG(ctor_flags, DIT_CTOR_FLAGS) != 0 {
		flags = SPL_FILE_DIR_KEY_AS_PATHNAME | SPL_FILE_DIR_CURRENT_AS_FILEINFO
		parsed = zend.ZendParseParameters(executeData.NumArgs(), "p|l", &path, &len_, &flags)
	} else {
		flags = SPL_FILE_DIR_KEY_AS_PATHNAME | SPL_FILE_DIR_CURRENT_AS_SELF
		parsed = zend.ZendParseParameters(executeData.NumArgs(), "p", &path, &len_)
	}
	if SPL_HAS_FLAG(ctor_flags, SPL_FILE_DIR_SKIPDOTS) != 0 {
		flags |= SPL_FILE_DIR_SKIPDOTS
	}
	if SPL_HAS_FLAG(ctor_flags, SPL_FILE_DIR_UNIXPATHS) != 0 {
		flags |= SPL_FILE_DIR_UNIXPATHS
	}
	if parsed == types.FAILURE {
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	if len_ == 0 {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Directory name must not be empty.")
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	intern = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if intern.GetPath() != nil {

		/* object is already initialized */

		zend.ZendRestoreErrorHandling(&error_handling)
		core.PhpErrorDocref(nil, faults.E_WARNING, "Directory object is already initialized")
		return
	}
	intern.SetFlags(flags)
	if SPL_HAS_FLAG(ctor_flags, DIT_CTOR_GLOB) != 0 && strstr(path, "glob://") != path {
		core.Spprintf(&path, 0, "glob://%s", path)
		SplFilesystemDirOpen(intern, path)
		zend.Efree(path)
	} else {
		SplFilesystemDirOpen(intern, path)
	}
	if zend.InstanceofFunction(intern.GetStd().GetCe(), spl_ce_RecursiveDirectoryIterator) != 0 {
		intern.SetIsRecursive(1)
	} else {
		intern.SetIsRecursive(0)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_DirectoryIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplFilesystemObjectConstruct(executeData, return_value, 0)
}
func zim_spl_DirectoryIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern.SetIndex(0)
	if intern.GetDirp() != nil {
		core.PhpStreamRewinddir(intern.GetDirp())
	}
	SplFilesystemDirRead(intern)
}
func zim_spl_DirectoryIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetDirp() != nil {
		return_value.SetLong(intern.GetIndex())
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func zim_spl_DirectoryIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetObject(zend.ZEND_THIS(executeData).GetObj())
	return_value.AddRefcount()
}
func zim_spl_DirectoryIterator_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var skip_dots int = SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern.GetIndex()++
	for {
		SplFilesystemDirRead(intern)
		if !(skip_dots != 0 && SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0) {
			break
		}
	}
	if intern.GetFileName() != nil {
		zend.Efree(intern.GetFileName())
		intern.SetFileName(nil)
	}
}
func zim_spl_DirectoryIterator_seek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var retval types.Zval
	var pos zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &pos) == types.FAILURE {
		return
	}
	if intern.GetIndex() > pos {

		/* we first rewind */

		zend.ZendCallMethodWith0Params(zend.ZEND_THIS(executeData), types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), intern.GetFuncRewind(), "rewind", nil)

		/* we first rewind */

	}
	for intern.GetIndex() < pos {
		var valid int = 0
		zend.ZendCallMethodWith0Params(zend.ZEND_THIS(executeData), types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), intern.GetFuncValid(), "valid", &retval)
		valid = zend.ZendIsTrue(&retval)
		zend.ZvalPtrDtor(&retval)
		if valid == 0 {
			faults.ThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+zend.ZEND_LONG_FMT+" is out of range", pos)
			return
		}
		zend.ZendCallMethodWith0Params(zend.ZEND_THIS(executeData), types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), intern.GetFuncNext(), "next", nil)
	}
}
func zim_spl_DirectoryIterator_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, intern.GetEntry().GetDName()[0] != '0')
	return
}
func zim_spl_SplFileInfo_getPath(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var path *byte
	var path_len int
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPath(intern, &path_len)
	if path != nil {
		return_value.SetRawString(b.CastStr(path, path_len))
		return
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
}
func zim_spl_SplFileInfo_getFilename(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var path_len int
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplFilesystemObjectGetPath(intern, &path_len)
	if path_len != 0 && path_len < intern.GetFileNameLen() {
		return_value.SetRawString(b.CastStr(intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1)))
		return
	} else {
		return_value.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
		return
	}
}
func zim_spl_DirectoryIterator_getFilename(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetRawString(b.CastStrAuto(intern.GetEntry().GetDName()))
	return
}
func zim_spl_SplFileInfo_getExtension(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var fname *byte = nil
	var p *byte
	var flen int
	var path_len int
	var idx int
	var ret *types.ZendString
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplFilesystemObjectGetPath(intern, &path_len)
	if path_len != 0 && path_len < intern.GetFileNameLen() {
		fname = intern.GetFileName() + path_len + 1
		flen = intern.GetFileNameLen() - (path_len + 1)
	} else {
		fname = intern.GetFileName()
		flen = intern.GetFileNameLen()
	}
	ret = standard.PhpBasename(fname, flen, nil, 0)
	p = zend.ZendMemrchr(ret.GetVal(), '.', ret.GetLen())
	if p != nil {
		idx = p - ret.GetVal()
		return_value.SetRawString(b.CastStr(ret.GetVal()+idx+1, ret.GetLen()-idx-1))
		types.ZendStringReleaseEx(ret, 0)
		return
	} else {
		types.ZendStringReleaseEx(ret, 0)
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
}
func zim_spl_DirectoryIterator_getExtension(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var p *byte
	var idx int
	var fname *types.ZendString
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.GetEntry().GetDName(), strlen(intern.GetEntry().GetDName()), nil, 0)
	p = zend.ZendMemrchr(fname.GetVal(), '.', fname.GetLen())
	if p != nil {
		idx = p - fname.GetVal()
		return_value.SetRawString(b.CastStr(fname.GetVal()+idx+1, fname.GetLen()-idx-1))
		types.ZendStringReleaseEx(fname, 0)
	} else {
		types.ZendStringReleaseEx(fname, 0)
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
}
func zim_spl_SplFileInfo_getBasename(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var fname *byte
	var suffix *byte = 0
	var flen int
	var slen int = 0
	var path_len int
	if zend.ZendParseParameters(executeData.NumArgs(), "|s", &suffix, &slen) == types.FAILURE {
		return
	}
	SplFilesystemObjectGetPath(intern, &path_len)
	if path_len != 0 && path_len < intern.GetFileNameLen() {
		fname = intern.GetFileName() + path_len + 1
		flen = intern.GetFileNameLen() - (path_len + 1)
	} else {
		fname = intern.GetFileName()
		flen = intern.GetFileNameLen()
	}
	return_value.SetString(standard.PhpBasename(fname, flen, suffix, slen))
	return
}
func zim_spl_DirectoryIterator_getBasename(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var suffix *byte = 0
	var slen int = 0
	var fname *types.ZendString
	if zend.ZendParseParameters(executeData.NumArgs(), "|s", &suffix, &slen) == types.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.GetEntry().GetDName(), strlen(intern.GetEntry().GetDName()), suffix, slen)
	return_value.SetString(fname)
}
func zim_spl_SplFileInfo_getPathname(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var path *byte
	var path_len int
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	if path != nil {
		return_value.SetRawString(b.CastStr(path, path_len))
		return
	} else {
		return_value.SetFalse()
		return
	}
}
func zim_spl_FilesystemIterator_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if SPL_FILE_DIR_KEY(intern, SPL_FILE_DIR_KEY_AS_FILENAME) {
		return_value.SetRawString(b.CastStrAuto(intern.GetEntry().GetDName()))
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		return_value.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
		return
	}
}
func zim_spl_FilesystemIterator_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if SPL_FILE_DIR_CURRENT(intern, SPL_FILE_DIR_CURRENT_AS_PATHNAME) {
		SplFilesystemObjectGetFileName(intern)
		return_value.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
		return
	} else if SPL_FILE_DIR_CURRENT(intern, SPL_FILE_DIR_CURRENT_AS_FILEINFO) {
		SplFilesystemObjectGetFileName(intern)
		SplFilesystemObjectCreateType(0, intern, SPL_FS_INFO, nil, return_value)
	} else {
		return_value.SetObject(zend.ZEND_THIS(executeData).GetObj())
		return_value.AddRefcount()
	}
}
func zim_spl_DirectoryIterator_isDot(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	types.ZVAL_BOOL(return_value, SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0)
	return
}
func zim_spl_SplFileInfo___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject
	var path *byte
	var len_ int
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "p", &path, &len_) == types.FAILURE {
		return
	}
	intern = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	SplFilesystemInfoSetFilename(intern, path, len_, 1)
}
func zim_spl_SplFileInfo_getPerms(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_PERMS, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getInode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_INODE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getSize(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_SIZE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getOwner(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_OWNER, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getGroup(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_GROUP, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getATime(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_ATIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getMTime(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_MTIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getCTime(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_CTIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getType(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_TYPE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isWritable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_W, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isReadable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_R, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isExecutable(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_X, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isFile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_FILE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isDir(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isLink(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_LINK, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getLinkTarget(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ret ssize_t
	var buff []byte
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if intern.GetFileName() == nil {
		SplFilesystemObjectGetFileName(intern)
	}
	if intern.GetFileName() == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Empty filename")
		return_value.SetFalse()
		return
	} else if !(zend.IS_ABSOLUTE_PATH(intern.GetFileName(), intern.GetFileNameLen())) {
		var expanded_path []byte
		if core.ExpandFilepathWithMode(intern.GetFileName(), expanded_path, nil, 0, zend.CWD_EXPAND) == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "No such file or directory")
			return_value.SetFalse()
			return
		}
		ret = zend.PhpSysReadlink(expanded_path, buff, core.MAXPATHLEN-1)
	} else {
		ret = zend.PhpSysReadlink(intern.GetFileName(), buff, core.MAXPATHLEN-1)
	}
	if ret == -1 {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Unable to read link %s, error: %s", intern.GetFileName(), strerror(errno))
		return_value.SetFalse()
	} else {

		/* Append NULL to the end of the string */

		buff[ret] = '0'
		return_value.SetRawString(b.CastStr(buff, ret))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getRealPath(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var buff []byte
	var filename *byte
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if intern.GetType() == SPL_FS_DIR && intern.GetFileName() == nil && intern.GetEntry().GetDName()[0] {
		SplFilesystemObjectGetFileName(intern)
	}
	if intern.GetOrigPath() != nil {
		filename = intern.GetOrigPath()
	} else {
		filename = intern.GetFileName()
	}
	if filename != nil && zend.VCWD_REALPATH(filename, buff) != nil {
		return_value.SetRawString(b.CastStrAuto(buff))
	} else {
		return_value.SetFalse()
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_openFile(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	SplFilesystemObjectCreateType(executeData.NumArgs(), intern, SPL_FS_FILE, nil, return_value)
}
func zim_spl_SplFileInfo_setFileClass(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ce *types.ClassEntry = spl_ce_SplFileObject
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(executeData.NumArgs(), "|C", &ce) == types.SUCCESS {
		intern.SetFileClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_setInfoClass(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ce *types.ClassEntry = spl_ce_SplFileInfo
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(executeData.NumArgs(), "|C", &ce) == types.SUCCESS {
		intern.SetInfoClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getFileInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ce *types.ClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(executeData.NumArgs(), "|C", &ce) == types.SUCCESS {
		SplFilesystemObjectCreateType(executeData.NumArgs(), intern, SPL_FS_INFO, ce, return_value)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getPathInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ce *types.ClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(executeData.NumArgs(), "|C", &ce) == types.SUCCESS {
		var path_len int
		var path *byte = SplFilesystemObjectGetPathname(intern, &path_len)
		if path != nil {
			var dpath *byte = zend.Estrndup(path, path_len)
			path_len = standard.PhpDirname(dpath, path_len)
			SplFilesystemObjectCreateInfo(intern, dpath, path_len, 1, ce, return_value)
			zend.Efree(dpath)
		}
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo___debugInfo(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetArray(SplFilesystemObjectGetDebugInfo(zend.getThis()))
	return
}
func zim_spl_SplFileInfo__bad_state_ex(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	faults.ThrowExceptionEx(spl_ce_LogicException, 0, "The parent constructor was not called: the object is in an "+"invalid state ")
}
func zim_spl_FilesystemIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplFilesystemObjectConstruct(executeData, return_value, DIT_CTOR_FLAGS|SPL_FILE_DIR_SKIPDOTS)
}
func zim_spl_FilesystemIterator_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var skip_dots int = SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	intern.SetIndex(0)
	if intern.GetDirp() != nil {
		core.PhpStreamRewinddir(intern.GetDirp())
	}
	for {
		SplFilesystemDirRead(intern)
		if !(skip_dots != 0 && SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0) {
			break
		}
	}
}
func zim_spl_FilesystemIterator_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetLong(intern.GetFlags() & (SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK))
	return
}
func zim_spl_FilesystemIterator_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var flags zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &flags) == types.FAILURE {
		return
	}
	intern.SubFlags(SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK)
	intern.AddFlags((SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK) & flags)
}
func zim_spl_RecursiveDirectoryIterator_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var allow_links types.ZendBool = 0
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParameters(executeData.NumArgs(), "|b", &allow_links) == types.FAILURE {
		return
	}
	if SplFilesystemIsInvalidOrDot(intern.GetEntry().GetDName()) != 0 {
		return_value.SetFalse()
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		if allow_links == 0 && !intern.IsDirFollowSymlinks() {
			standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_LINK, return_value)
			if zend.ZendIsTrue(return_value) != 0 {
				return_value.SetFalse()
				return
			}
		}
		standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, return_value)
	}
}
func zim_spl_RecursiveDirectoryIterator_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var zpath types.Zval
	var zflags types.Zval
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var subdir *SplFilesystemObject
	var slash byte = b.Cond(SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_UNIXPATHS) != 0, '/', zend.DEFAULT_SLASH)
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplFilesystemObjectGetFileName(intern)
	zflags.SetLong(intern.GetFlags())
	zpath.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
	SplInstantiateArgEx2(types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), return_value, &zpath, &zflags)
	zend.ZvalPtrDtor(&zpath)
	subdir = Z_SPLFILESYSTEM_P(return_value)
	if subdir != nil {
		if intern.GetSubPath() != nil && intern.GetSubPath()[0] {
			subdir.SetSubPathLen(core.Spprintf(subdir.GetSubPath(), 0, "%s%c%s", intern.GetSubPath(), slash, intern.GetEntry().GetDName()))
		} else {
			subdir.SetSubPathLen(strlen(intern.GetEntry().GetDName()))
			subdir.SetSubPath(zend.Estrndup(intern.GetEntry().GetDName(), subdir.GetSubPathLen()))
		}
		subdir.SetInfoClass(intern.GetInfoClass())
		subdir.SetFileClass(intern.GetFileClass())
		subdir.SetOth(intern.GetOth())
	}
}
func zim_spl_RecursiveDirectoryIterator_getSubPath(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		return_value.SetRawString(b.CastStr(intern.GetSubPath(), intern.GetSubPathLen()))
		return
	} else {
		zend.ZVAL_EMPTY_STRING(return_value)
		return
	}
}
func zim_spl_RecursiveDirectoryIterator_getSubPathname(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var slash byte = b.Cond(SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_UNIXPATHS) != 0, '/', zend.DEFAULT_SLASH)
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		return_value.SetString(core.Strpprintf(0, "%s%c%s", intern.GetSubPath(), slash, intern.GetEntry().GetDName()))
		return
	} else {
		return_value.SetRawString(b.CastStrAuto(intern.GetEntry().GetDName()))
		return
	}
}
func zim_spl_RecursiveDirectoryIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplFilesystemObjectConstruct(executeData, return_value, DIT_CTOR_FLAGS)
}
func zim_spl_GlobIterator___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	SplFilesystemObjectConstruct(executeData, return_value, DIT_CTOR_FLAGS|DIT_CTOR_GLOB)
}
func zim_spl_GlobIterator_count(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetDirp() != nil && core.PhpStreamIs(intern.GetDirp(), &streams.PhpGlobStreamOps) {
		return_value.SetLong(streams.PhpGlobStreamGetCount(intern.GetDirp(), nil))
		return
	} else {

		/* should not happen */

		core.PhpErrorDocref(nil, faults.E_ERROR, "GlobIterator lost glob state")

		/* should not happen */

	}
}
func SplFilesystemDirGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		faults.ThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = Z_SPLFILESYSTEM_P(object)
	iterator = SplFilesystemObjectToIterator(dir_object)
	object.AddRefcount()
	iterator.GetIntern().GetData().SetObject(object.GetObj())
	iterator.GetIntern().SetFuncs(&SplFilesystemDirItFuncs)

	/* ->current must be initialized; rewind doesn't set it and valid
	 * doesn't check whether it's set */

	iterator.SetCurrent(*object)
	return iterator.GetIntern()
}
func SplFilesystemDirItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	zend.ZvalPtrDtor(iterator.GetIntern().GetData())
}
func SplFilesystemDirItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	if object.GetEntry().GetDName()[0] != '0' {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func SplFilesystemDirItCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	return iterator.GetCurrent()
}
func SplFilesystemDirItCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	key.SetLong(object.GetIndex())
}
func SplFilesystemDirItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	object.GetIndex()++
	SplFilesystemDirRead(object)
	if object.GetFileName() != nil {
		zend.Efree(object.GetFileName())
		object.SetFileName(nil)
	}
}
func SplFilesystemDirItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	object.SetIndex(0)
	if object.GetDirp() != nil {
		core.PhpStreamRewinddir(object.GetDirp())
	}
	SplFilesystemDirRead(object)
}
func SplFilesystemTreeItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	zend.ZvalPtrDtor(iterator.GetIntern().GetData())
	zend.ZvalPtrDtor(iterator.GetCurrent())
}
func SplFilesystemTreeItCurrentData(iter *zend.ZendObjectIterator) *types.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	if SPL_FILE_DIR_CURRENT(object, SPL_FILE_DIR_CURRENT_AS_PATHNAME) {
		if iterator.GetCurrent().IsUndef() {
			SplFilesystemObjectGetFileName(object)
			iterator.GetCurrent().SetRawString(b.CastStr(object.GetFileName(), object.GetFileNameLen()))
		}
		return iterator.GetCurrent()
	} else if SPL_FILE_DIR_CURRENT(object, SPL_FILE_DIR_CURRENT_AS_FILEINFO) {
		if iterator.GetCurrent().IsUndef() {
			SplFilesystemObjectGetFileName(object)
			SplFilesystemObjectCreateType(0, object, SPL_FS_INFO, nil, iterator.GetCurrent())
		}
		return iterator.GetCurrent()
	} else {
		return iterator.GetIntern().GetData()
	}
}
func SplFilesystemTreeItCurrentKey(iter *zend.ZendObjectIterator, key *types.Zval) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	if SPL_FILE_DIR_KEY(object, SPL_FILE_DIR_KEY_AS_FILENAME) {
		key.SetRawString(b.CastStrAuto(object.GetEntry().GetDName()))
	} else {
		SplFilesystemObjectGetFileName(object)
		key.SetRawString(b.CastStr(object.GetFileName(), object.GetFileNameLen()))
	}
}
func SplFilesystemTreeItMoveForward(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	object.GetIndex()++
	for {
		SplFilesystemDirRead(object)
		if SplFilesystemIsDot(object.GetEntry().GetDName()) == 0 {
			break
		}
	}
	if object.GetFileName() != nil {
		zend.Efree(object.GetFileName())
		object.SetFileName(nil)
	}
	if !(iterator.GetCurrent().IsUndef()) {
		zend.ZvalPtrDtor(iterator.GetCurrent())
		iterator.GetCurrent().SetUndef()
	}
}
func SplFilesystemTreeItRewind(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	object.SetIndex(0)
	if object.GetDirp() != nil {
		core.PhpStreamRewinddir(object.GetDirp())
	}
	for {
		SplFilesystemDirRead(object)
		if SplFilesystemIsDot(object.GetEntry().GetDName()) == 0 {
			break
		}
	}
	if !(iterator.GetCurrent().IsUndef()) {
		zend.ZvalPtrDtor(iterator.GetCurrent())
		iterator.GetCurrent().SetUndef()
	}
}
func SplFilesystemTreeGetIterator(ce *types.ClassEntry, object *types.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		faults.ThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = Z_SPLFILESYSTEM_P(object)
	iterator = SplFilesystemObjectToIterator(dir_object)
	object.AddRefcount()
	iterator.GetIntern().GetData().SetObject(object.GetObj())
	iterator.GetIntern().SetFuncs(&SplFilesystemTreeItFuncs)
	return iterator.GetIntern()
}
func SplFilesystemObjectCast(readobj *types.Zval, writeobj *types.Zval, type_ int) int {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(readobj)
	if type_ == types.IS_STRING {
		if types.Z_OBJCE_P(readobj).GetTostring() != nil {
			return zend.ZendStdCastObjectTostring(readobj, writeobj, type_)
		}
		switch intern.GetType() {
		case SPL_FS_INFO:
			fallthrough
		case SPL_FS_FILE:
			writeobj.SetRawString(b.CastStr(intern.GetFileName(), intern.GetFileNameLen()))
			return types.SUCCESS
		case SPL_FS_DIR:
			writeobj.SetRawString(b.CastStrAuto(intern.GetEntry().GetDName()))
			return types.SUCCESS
		}
	} else if type_ == zend._IS_BOOL {
		writeobj.SetTrue()
		return types.SUCCESS
	}
	writeobj.SetNull()
	return types.FAILURE
}
func SplFilesystemFileRead(intern *SplFilesystemObject, silent int) int {
	var buf *byte
	var line_len int = 0
	var line_add zend.ZendLong = b.Cond(intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()), 1, 0)
	SplFilesystemFileFreeLine(intern)
	if core.PhpStreamEof(intern.GetStream()) != 0 {
		if silent == 0 {
			faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
		}
		return types.FAILURE
	}
	if intern.GetMaxLineLen() > 0 {
		buf = zend.SafeEmalloc(intern.GetMaxLineLen()+1, b.SizeOf("char"), 0)
		if core.PhpStreamGetLine(intern.GetStream(), buf, intern.GetMaxLineLen()+1, &line_len) == nil {
			zend.Efree(buf)
			buf = nil
		} else {
			buf[line_len] = '0'
		}
	} else {
		buf = core.PhpStreamGetLine(intern.GetStream(), nil, 0, &line_len)
	}
	if buf == nil {
		intern.SetCurrentLine(zend.Estrdup(""))
		intern.SetCurrentLineLen(0)
	} else {
		if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_DROP_NEW_LINE) != 0 {
			if line_len > 0 && buf[line_len-1] == '\n' {
				line_len--
				if line_len > 0 && buf[line_len-1] == '\r' {
					line_len--
				}
				buf[line_len] = '0'
			}
		}
		intern.SetCurrentLine(buf)
		intern.SetCurrentLineLen(line_len)
	}
	intern.SetCurrentLineNum(intern.GetCurrentLineNum() + line_add)
	return types.SUCCESS
}
func SplFilesystemFileCall(intern *SplFilesystemObject, func_ptr *zend.ZendFunction, pass_num_args int, return_value *types.Zval, arg2 *types.Zval) int {
	var fci types.ZendFcallInfo
	var fcic types.ZendFcallInfoCache
	var zresource_ptr *types.Zval = intern.GetZresource()
	var params *types.Zval
	var retval types.Zval
	var result int
	var num_args int = pass_num_args + b.Cond(arg2 != nil, 2, 1)
	if zresource_ptr.IsUndef() {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return types.FAILURE
	}
	params = (*types.Zval)(zend.SafeEmalloc(num_args, b.SizeOf("zval"), 0))
	params[0] = *zresource_ptr
	if arg2 != nil {
		params[1] = *arg2
	}
	if zend.ZendGetParametersArrayEx(pass_num_args, params+b.Cond(arg2 != nil, 2, 1)) != types.SUCCESS {
		zend.Efree(params)
		zend.ZendWrongParamCount()
		return types.FAILURE
	}
	retval.SetUndef()
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(nil)
	fci.SetRetval(&retval)
	fci.SetParamCount(num_args)
	fci.SetParams(params)
	fci.SetNoSeparation(1)
	fci.GetFunctionName().SetString(func_ptr.GetFunctionName())
	fcic.SetFunctionHandler(func_ptr)
	fcic.SetCalledScope(nil)
	fcic.SetObject(nil)
	result = zend.ZendCallFunction(&fci, &fcic)
	if result == types.FAILURE || retval.IsUndef() {
		return_value.SetFalse()
	} else {
		zend.ZVAL_ZVAL(return_value, &retval, 0, 0)
	}
	zend.Efree(params)
	return result
}
func SplFilesystemFileReadCsv(intern *SplFilesystemObject, delimiter byte, enclosure byte, escape int, return_value *types.Zval) int {
	var ret int = types.SUCCESS
	var value *types.Zval
	for {
		ret = SplFilesystemFileRead(intern, 1)
		if !(ret == types.SUCCESS && intern.GetCurrentLineLen() == 0 && SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_SKIP_EMPTY) != 0) {
			break
		}
	}
	if ret == types.SUCCESS {
		var buf_len int = intern.GetCurrentLineLen()
		var buf *byte = zend.Estrndup(intern.GetCurrentLine(), buf_len)
		if !(intern.GetCurrentZval().IsUndef()) {
			zend.ZvalPtrDtor(intern.GetCurrentZval())
			intern.GetCurrentZval().SetUndef()
		}
		standard.PhpFgetcsv(intern.GetStream(), delimiter, enclosure, escape, buf_len, buf, intern.GetCurrentZval())
		if return_value != nil {
			value = intern.GetCurrentZval()
			types.ZVAL_COPY_DEREF(return_value, value)
		}
	}
	return ret
}
func SplFilesystemFileReadLineEx(this_ptr *types.Zval, intern *SplFilesystemObject, silent int) int {
	var retval types.Zval

	/* 1) use fgetcsv? 2) overloaded call the function, 3) do it directly */

	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 || intern.GetFuncGetCurr().GetScope() != spl_ce_SplFileObject {
		if core.PhpStreamEof(intern.GetStream()) != 0 {
			if silent == 0 {
				faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
			}
			return types.FAILURE
		}
		if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 {
			return SplFilesystemFileReadCsv(intern, intern.GetDelimiter(), intern.GetEnclosure(), intern.GetEscape(), nil)
		} else {
			var executeData *zend.ZendExecuteData = zend.CurrEX()
			zend.ZendCallMethodWith0Params(this_ptr, types.Z_OBJCE_P(zend.ZEND_THIS(executeData)), intern.GetFuncGetCurr(), "getCurrentLine", &retval)
		}
		if !(retval.IsUndef()) {
			if intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()) {
				intern.GetCurrentLineNum()++
			}
			SplFilesystemFileFreeLine(intern)
			if retval.IsType(types.IS_STRING) {
				intern.SetCurrentLine(zend.Estrndup(retval.GetStr().GetVal(), retval.GetStr().GetLen()))
				intern.SetCurrentLineLen(retval.GetStr().GetLen())
			} else {
				var value *types.Zval = &retval
				types.ZVAL_COPY_DEREF(intern.GetCurrentZval(), value)
			}
			zend.ZvalPtrDtor(&retval)
			return types.SUCCESS
		} else {
			return types.FAILURE
		}
	} else {
		return SplFilesystemFileRead(intern, silent)
	}

	/* 1) use fgetcsv? 2) overloaded call the function, 3) do it directly */
}
func SplFilesystemFileIsEmptyLine(intern *SplFilesystemObject) int {
	if intern.GetCurrentLine() != nil {
		return intern.GetCurrentLineLen() == 0
	} else if !(intern.GetCurrentZval().IsUndef()) {
		switch intern.GetCurrentZval().GetType() {
		case types.IS_STRING:
			return intern.GetCurrentZval().GetStr().GetLen() == 0
		case types.IS_ARRAY:
			if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 && types.Z_ARRVAL(intern.GetCurrentZval()).GetNNumOfElements() == 1 {
				var idx uint32 = 0
				var first *types.Zval
				for types.Z_ARRVAL(intern.GetCurrentZval()).GetArData()[idx].GetVal().IsUndef() {
					idx++
				}
				first = types.Z_ARRVAL(intern.GetCurrentZval()).GetArData()[idx].GetVal()
				return first.IsType(types.IS_STRING) && first.GetStr().GetLen() == 0
			}
			return types.Z_ARRVAL(intern.GetCurrentZval()).GetNNumOfElements() == 0
		case types.IS_NULL:
			return 1
		default:
			return 0
		}
	} else {
		return 1
	}
}
func SplFilesystemFileReadLine(this_ptr *types.Zval, intern *SplFilesystemObject, silent int) int {
	var ret int = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	for SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_SKIP_EMPTY) != 0 && ret == types.SUCCESS && SplFilesystemFileIsEmptyLine(intern) != 0 {
		SplFilesystemFileFreeLine(intern)
		ret = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	}
	return ret
}
func SplFilesystemFileRewind(this_ptr *types.Zval, intern *SplFilesystemObject) {
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if -1 == core.PhpStreamRewind(intern.GetStream()) {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot rewind file %s", intern.GetFileName())
	} else {
		SplFilesystemFileFreeLine(intern)
		intern.SetCurrentLineNum(0)
	}
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		SplFilesystemFileReadLine(this_ptr, intern, 1)
	}
}
func zim_spl_SplFileObject___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var use_include_path types.ZendBool = 0
	var p1 *byte
	var p2 *byte
	var tmp_path *byte
	var tmp_path_len int
	var error_handling zend.ZendErrorHandling
	intern.SetOpenMode(nil)
	intern.SetOpenModeLen(0)
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "p|sbr!", intern.GetFileName(), intern.GetFileNameLen(), intern.GetOpenMode(), intern.GetOpenModeLen(), &use_include_path, intern.GetZcontext()) == types.FAILURE {
		intern.SetOpenMode(nil)
		intern.SetFileName(nil)
		return
	}
	if intern.GetOpenMode() == nil {
		intern.SetOpenMode("r")
		intern.SetOpenModeLen(1)
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if SplFilesystemFileOpen(intern, use_include_path, 0) == types.SUCCESS {
		tmp_path_len = strlen(intern.GetStream().GetOrigPath())
		if tmp_path_len > 1 && IS_SLASH_AT(intern.GetStream().GetOrigPath(), tmp_path_len-1) {
			tmp_path_len--
		}
		tmp_path = zend.Estrndup(intern.GetStream().GetOrigPath(), tmp_path_len)
		p1 = strrchr(tmp_path, '/')
		p2 = 0
		if p1 != nil || p2 != nil {
			intern.SetPathLen(b.Cond(p1 > p2, p1, p2) - tmp_path)
		} else {
			intern.SetPathLen(0)
		}
		zend.Efree(tmp_path)
		intern.SetPath(zend.Estrndup(intern.GetStream().GetOrigPath(), intern.GetPathLen()))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplTempFileObject___construct(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var max_memory zend.ZendLong = core.PHP_STREAM_MAX_MEM
	var tmp_fname []byte
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersThrow(executeData.NumArgs(), "|l", &max_memory) == types.FAILURE {
		return
	}
	if max_memory < 0 {
		intern.SetFileName("php://memory")
		intern.SetFileNameLen(12)
	} else if executeData.NumArgs() != 0 {
		intern.SetFileNameLen(core.Slprintf(tmp_fname, b.SizeOf("tmp_fname"), "php://temp/maxmemory:"+zend.ZEND_LONG_FMT, max_memory))
		intern.SetFileName(tmp_fname)
	} else {
		intern.SetFileName("php://temp")
		intern.SetFileNameLen(10)
	}
	intern.SetOpenMode("wb")
	intern.SetOpenModeLen(1)
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if SplFilesystemFileOpen(intern, 0, 0) == types.SUCCESS {
		intern.SetPathLen(0)
		intern.SetPath(zend.Estrndup("", 0))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileObject_rewind(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplFilesystemFileRewind(zend.ZEND_THIS(executeData), intern)
}
func zim_spl_SplFileObject_eof(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	types.ZVAL_BOOL(return_value, core.PhpStreamEof(intern.GetStream()) != 0)
	return
}
func zim_spl_SplFileObject_valid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		types.ZVAL_BOOL(return_value, intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()))
		return
	} else {
		if intern.GetStream() == nil {
			return_value.SetFalse()
			return
		}
		types.ZVAL_BOOL(return_value, core.PhpStreamEof(intern.GetStream()) == 0)
	}
}
func zim_spl_SplFileObject_fgets(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if SplFilesystemFileRead(intern, 0) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	return_value.SetRawString(b.CastStr(intern.GetCurrentLine(), intern.GetCurrentLineLen()))
	return
}
func zim_spl_SplFileObject_current(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if intern.GetCurrentLine() == nil && intern.GetCurrentZval().IsUndef() {
		SplFilesystemFileReadLine(zend.ZEND_THIS(executeData), intern, 1)
	}
	if intern.GetCurrentLine() != nil && (SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) == 0 || intern.GetCurrentZval().IsUndef()) {
		return_value.SetRawString(b.CastStr(intern.GetCurrentLine(), intern.GetCurrentLineLen()))
		return
	} else if !(intern.GetCurrentZval().IsUndef()) {
		var value *types.Zval = intern.GetCurrentZval()
		types.ZVAL_COPY_DEREF(return_value, value)
		return
	}
	return_value.SetFalse()
	return
}
func zim_spl_SplFileObject_key(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}

	/*    Do not read the next line to support correct counting with fgetc()
	      if (!intern->current_line) {
	          spl_filesystem_file_read_line(ZEND_THIS(executeData), intern, 1);
	      } */

	return_value.SetLong(intern.GetCurrentLineNum())
	return
}
func zim_spl_SplFileObject_next(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	SplFilesystemFileFreeLine(intern)
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		SplFilesystemFileReadLine(zend.ZEND_THIS(executeData), intern, 1)
	}
	intern.GetCurrentLineNum()++
}
func zim_spl_SplFileObject_setFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParameters(executeData.NumArgs(), "l", intern.GetFlags()) == types.FAILURE {
		return
	}
}
func zim_spl_SplFileObject_getFlags(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetLong(intern.GetFlags() & SPL_FILE_OBJECT_MASK)
	return
}
func zim_spl_SplFileObject_setMaxLineLen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var max_len zend.ZendLong
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &max_len) == types.FAILURE {
		return
	}
	if max_len < 0 {
		faults.ThrowExceptionEx(spl_ce_DomainException, 0, "Maximum line length must be greater than or equal zero")
		return
	}
	intern.SetMaxLineLen(max_len)
}
func zim_spl_SplFileObject_getMaxLineLen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetLong(zend.ZendLong(intern.GetMaxLineLen()))
	return
}
func zim_spl_SplFileObject_hasChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
	return_value.SetFalse()
	return
}
func zim_spl_SplFileObject_getChildren(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if zend.ZendParseParametersNone() == types.FAILURE {
		return
	}
}
func zim_spl_SplFileObject_fgetcsv(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var delimiter byte = intern.GetDelimiter()
	var enclosure byte = intern.GetEnclosure()
	var escape int = intern.GetEscape()
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == types.SUCCESS {
		if intern.GetStream() == nil {
			faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
			return
		}
		switch executeData.NumArgs() {
		case 3:
			if esc_len > 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "escape must be empty or a single character")
				return_value.SetFalse()
				return
			}
			if esc_len == 0 {
				escape = standard.PHP_CSV_NO_ESCAPE
			} else {
				escape = uint8(esc[0])
			}
			fallthrough
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "enclosure must be a character")
				return_value.SetFalse()
				return
			}
			enclosure = enclo[0]
			fallthrough
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "delimiter must be a character")
				return_value.SetFalse()
				return
			}
			delimiter = delim[0]
			fallthrough
		case 0:

		}
		SplFilesystemFileReadCsv(intern, delimiter, enclosure, escape, return_value)
	}
}
func zim_spl_SplFileObject_fputcsv(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var delimiter byte = intern.GetDelimiter()
	var enclosure byte = intern.GetEnclosure()
	var escape int = intern.GetEscape()
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	var ret zend.ZendLong
	var fields *types.Zval = nil
	if zend.ZendParseParameters(executeData.NumArgs(), "a|sss", &fields, &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == types.SUCCESS {
		switch executeData.NumArgs() {
		case 4:
			switch esc_len {
			case 0:
				escape = standard.PHP_CSV_NO_ESCAPE
			case 1:
				escape = uint8(esc[0])
			default:
				core.PhpErrorDocref(nil, faults.E_WARNING, "escape must be empty or a single character")
				return_value.SetFalse()
				return
			}
			fallthrough
		case 3:
			if e_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "enclosure must be a character")
				return_value.SetFalse()
				return
			}
			enclosure = enclo[0]
			fallthrough
		case 2:
			if d_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "delimiter must be a character")
				return_value.SetFalse()
				return
			}
			delimiter = delim[0]
			fallthrough
		case 1:
			fallthrough
		case 0:

		}
		ret = standard.PhpFputcsv(intern.GetStream(), fields, delimiter, enclosure, escape)
		if ret < 0 {
			return_value.SetFalse()
			return
		}
		return_value.SetLong(ret)
		return
	}
}
func zim_spl_SplFileObject_setCsvControl(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == types.SUCCESS {
		switch executeData.NumArgs() {
		case 3:
			switch esc_len {
			case 0:
				escape = standard.PHP_CSV_NO_ESCAPE
			case 1:
				escape = uint8(esc[0])
			default:
				core.PhpErrorDocref(nil, faults.E_WARNING, "escape must be empty or a single character")
				return_value.SetFalse()
				return
			}
			fallthrough
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "enclosure must be a character")
				return_value.SetFalse()
				return
			}
			enclosure = enclo[0]
			fallthrough
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, faults.E_WARNING, "delimiter must be a character")
				return_value.SetFalse()
				return
			}
			delimiter = delim[0]
			fallthrough
		case 0:

		}
		intern.SetDelimiter(delimiter)
		intern.SetEnclosure(enclosure)
		intern.SetEscape(escape)
	}
}
func zim_spl_SplFileObject_getCsvControl(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var delimiter []byte
	var enclosure []byte
	var escape []byte
	zend.ArrayInit(return_value)
	delimiter[0] = intern.GetDelimiter()
	delimiter[1] = '0'
	enclosure[0] = intern.GetEnclosure()
	enclosure[1] = '0'
	if intern.GetEscape() == standard.PHP_CSV_NO_ESCAPE {
		escape[0] = '0'
	} else {
		escape[0] = uint8(intern.GetEscape())
		escape[1] = '0'
	}
	zend.AddNextIndexString(return_value, delimiter)
	zend.AddNextIndexString(return_value, enclosure)
	zend.AddNextIndexString(return_value, escape)
}
func zim_spl_SplFileObject_flock(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(types.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "flock", b.SizeOf("\"flock\"")-1))
	if func_ptr == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "flock")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, executeData.NumArgs(), return_value, nil)
}
func zim_spl_SplFileObject_fflush(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	types.ZVAL_BOOL(return_value, core.PhpStreamFlush(intern.GetStream()) == 0)
	return
}
func zim_spl_SplFileObject_ftell(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var ret zend.ZendLong
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	ret = intern.GetStream().GetPosition()
	if ret == -1 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(ret)
		return
	}
}
func zim_spl_SplFileObject_fseek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var pos zend.ZendLong
	var whence zend.ZendLong = r.SEEK_SET
	if zend.ZendParseParameters(executeData.NumArgs(), "l|l", &pos, &whence) == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	return_value.SetLong(core.PhpStreamSeek(intern.GetStream(), pos, int(whence)))
	return
}
func zim_spl_SplFileObject_fgetc(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var buf []byte
	var result int
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	result = core.PhpStreamGetc(intern.GetStream())
	if result == r.EOF {
		return_value.SetFalse()
	} else {
		if result == '\n' {
			intern.GetCurrentLineNum()++
		}
		buf[0] = result
		buf[1] = '0'
		return_value.SetRawString(b.CastStr(buf, 1))
		return
	}
}
func zim_spl_SplFileObject_fgetss(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var arg2 types.Zval
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if intern.GetMaxLineLen() > 0 {
		arg2.SetLong(intern.GetMaxLineLen())
	} else {
		arg2.SetLong(1024)
	}
	SplFilesystemFileFreeLine(intern)
	intern.GetCurrentLineNum()++
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(types.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fgetss", b.SizeOf("\"fgetss\"")-1))
	if func_ptr == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fgetss")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, executeData.NumArgs(), return_value, &arg2)
}
func zim_spl_SplFileObject_fpassthru(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	return_value.SetLong(core.PhpStreamPassthru(intern.GetStream()))
	return
}
func zim_spl_SplFileObject_fscanf(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	intern.GetCurrentLineNum()++
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(types.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fscanf", b.SizeOf("\"fscanf\"")-1))
	if func_ptr == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fscanf")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, executeData.NumArgs(), return_value, nil)
}
func zim_spl_SplFileObject_fwrite(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var str *byte
	var str_len int
	var length zend.ZendLong = 0
	var written ssize_t
	if zend.ZendParseParameters(executeData.NumArgs(), "s|l", &str, &str_len, &length) == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if executeData.NumArgs() > 1 {
		if length >= 0 {
			str_len = cli.MIN(int(length), str_len)
		} else {

			/* Negative length given, nothing to write */

			str_len = 0

			/* Negative length given, nothing to write */

		}
	}
	if str_len == 0 {
		return_value.SetLong(0)
		return
	}
	written = core.PhpStreamWrite(intern.GetStream(), str, str_len)
	if written < 0 {
		return_value.SetFalse()
		return
	}
	return_value.SetLong(written)
	return
}
func zim_spl_SplFileObject_fread(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var length zend.ZendLong = 0
	var str *types.ZendString
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &length) == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if length <= 0 {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Length parameter must be greater than 0")
		return_value.SetFalse()
		return
	}
	str = streams.PhpStreamReadToStr(intern.GetStream(), length)
	if str == nil {
		return_value.SetFalse()
		return
	}
	return_value.SetString(str)
	return
}
func zim_spl_SplFileObject_fstat(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(types.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fstat", b.SizeOf("\"fstat\"")-1))
	if func_ptr == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fstat")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, executeData.NumArgs(), return_value, nil)
}
func zim_spl_SplFileObject_ftruncate(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var size zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &size) == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if core.PhpStreamTruncateSupported(intern.GetStream()) == 0 {
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Can't truncate file %s", intern.GetFileName())
		return_value.SetFalse()
		return
	}
	types.ZVAL_BOOL(return_value, 0 == core.PhpStreamTruncateSetSize(intern.GetStream(), size))
	return
}
func zim_spl_SplFileObject_seek(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS(executeData))
	var line_pos zend.ZendLong
	if zend.ZendParseParameters(executeData.NumArgs(), "l", &line_pos) == types.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		faults.ThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if line_pos < 0 {
		faults.ThrowExceptionEx(spl_ce_LogicException, 0, "Can't seek file %s to negative line "+zend.ZEND_LONG_FMT, intern.GetFileName(), line_pos)
		return_value.SetFalse()
		return
	}
	SplFilesystemFileRewind(zend.ZEND_THIS(executeData), intern)
	for intern.GetCurrentLineNum() < line_pos {
		if SplFilesystemFileReadLine(zend.ZEND_THIS(executeData), intern, 1) == types.FAILURE {
			break
		}
	}
}
func ZmStartupSplDirectory(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplFileInfo, "SplFileInfo", SplFilesystemObjectNew, spl_SplFileInfo_functions)
	memcpy(&SplFilesystemObjectHandlers, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplFilesystemObjectHandlers.SetOffset(zend_long((*byte)(&((*SplFilesystemObject)(nil).GetStd())) - (*byte)(nil)))
	SplFilesystemObjectHandlers.SetCloneObj(SplFilesystemObjectClone)
	SplFilesystemObjectHandlers.SetCastObject(SplFilesystemObjectCast)
	SplFilesystemObjectHandlers.SetDtorObj(SplFilesystemObjectDestroyObject)
	SplFilesystemObjectHandlers.SetFreeObj(SplFilesystemObjectFreeStorage)
	spl_ce_SplFileInfo.SetSerialize(zend.ZendClassSerializeDeny)
	spl_ce_SplFileInfo.SetUnserialize(zend.ZendClassUnserializeDeny)
	SplRegisterSubClass(&spl_ce_DirectoryIterator, spl_ce_SplFileInfo, "DirectoryIterator", SplFilesystemObjectNew, spl_DirectoryIterator_functions)
	zend.ZendClassImplements(spl_ce_DirectoryIterator, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_DirectoryIterator, 1, spl_ce_SeekableIterator)
	spl_ce_DirectoryIterator.SetGetIterator(SplFilesystemDirGetIterator)
	SplRegisterSubClass(&spl_ce_FilesystemIterator, spl_ce_DirectoryIterator, "FilesystemIterator", SplFilesystemObjectNew, spl_FilesystemIterator_functions)
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_MODE_MASK", b.SizeOf("\"CURRENT_MODE_MASK\"")-1, zend.ZendLong(SPL_FILE_DIR_CURRENT_MODE_MASK))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_PATHNAME", b.SizeOf("\"CURRENT_AS_PATHNAME\"")-1, zend.ZendLong(SPL_FILE_DIR_CURRENT_AS_PATHNAME))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_FILEINFO", b.SizeOf("\"CURRENT_AS_FILEINFO\"")-1, zend.ZendLong(SPL_FILE_DIR_CURRENT_AS_FILEINFO))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_SELF", b.SizeOf("\"CURRENT_AS_SELF\"")-1, zend.ZendLong(SPL_FILE_DIR_CURRENT_AS_SELF))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_MODE_MASK", b.SizeOf("\"KEY_MODE_MASK\"")-1, zend.ZendLong(SPL_FILE_DIR_KEY_MODE_MASK))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_AS_PATHNAME", b.SizeOf("\"KEY_AS_PATHNAME\"")-1, zend.ZendLong(SPL_FILE_DIR_KEY_AS_PATHNAME))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "FOLLOW_SYMLINKS", b.SizeOf("\"FOLLOW_SYMLINKS\"")-1, zend.ZendLong(SPL_FILE_DIR_FOLLOW_SYMLINKS))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_AS_FILENAME", b.SizeOf("\"KEY_AS_FILENAME\"")-1, zend.ZendLong(SPL_FILE_DIR_KEY_AS_FILENAME))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "NEW_CURRENT_AND_KEY", b.SizeOf("\"NEW_CURRENT_AND_KEY\"")-1, zend.ZendLong(SPL_FILE_DIR_KEY_AS_FILENAME|SPL_FILE_DIR_CURRENT_AS_FILEINFO))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "OTHER_MODE_MASK", b.SizeOf("\"OTHER_MODE_MASK\"")-1, zend.ZendLong(SPL_FILE_DIR_OTHERS_MASK))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "SKIP_DOTS", b.SizeOf("\"SKIP_DOTS\"")-1, zend.ZendLong(SPL_FILE_DIR_SKIPDOTS))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "UNIX_PATHS", b.SizeOf("\"UNIX_PATHS\"")-1, zend.ZendLong(SPL_FILE_DIR_UNIXPATHS))
	spl_ce_FilesystemIterator.SetGetIterator(SplFilesystemTreeGetIterator)
	SplRegisterSubClass(&spl_ce_RecursiveDirectoryIterator, spl_ce_FilesystemIterator, "RecursiveDirectoryIterator", SplFilesystemObjectNew, spl_RecursiveDirectoryIterator_functions)
	zend.ZendClassImplements(spl_ce_RecursiveDirectoryIterator, 1, spl_ce_RecursiveIterator)
	memcpy(&SplFilesystemObjectCheckHandlers, &SplFilesystemObjectHandlers, b.SizeOf("zend_object_handlers"))
	SplFilesystemObjectCheckHandlers.SetCloneObj(nil)
	SplFilesystemObjectCheckHandlers.SetGetMethod(SplFilesystemObjectGetMethodCheck)
	SplRegisterSubClass(&spl_ce_GlobIterator, spl_ce_FilesystemIterator, "GlobIterator", SplFilesystemObjectNewCheck, spl_GlobIterator_functions)
	zend.ZendClassImplements(spl_ce_GlobIterator, 1, spl_ce_Countable)
	SplRegisterSubClass(&spl_ce_SplFileObject, spl_ce_SplFileInfo, "SplFileObject", SplFilesystemObjectNewCheck, spl_SplFileObject_functions)
	zend.ZendClassImplements(spl_ce_SplFileObject, 1, spl_ce_RecursiveIterator)
	zend.ZendClassImplements(spl_ce_SplFileObject, 1, spl_ce_SeekableIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "DROP_NEW_LINE", b.SizeOf("\"DROP_NEW_LINE\"")-1, zend.ZendLong(SPL_FILE_OBJECT_DROP_NEW_LINE))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "READ_AHEAD", b.SizeOf("\"READ_AHEAD\"")-1, zend.ZendLong(SPL_FILE_OBJECT_READ_AHEAD))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "SKIP_EMPTY", b.SizeOf("\"SKIP_EMPTY\"")-1, zend.ZendLong(SPL_FILE_OBJECT_SKIP_EMPTY))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "READ_CSV", b.SizeOf("\"READ_CSV\"")-1, zend.ZendLong(SPL_FILE_OBJECT_READ_CSV))
	SplRegisterSubClass(&spl_ce_SplTempFileObject, spl_ce_SplFileObject, "SplTempFileObject", SplFilesystemObjectNewCheck, spl_SplTempFileObject_functions)
	return types.SUCCESS
}
