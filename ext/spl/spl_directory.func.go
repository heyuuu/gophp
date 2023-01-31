// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/sapi/cli"
	"sik/zend"
)

func SplFilesystemFromObj(obj *zend.ZendObject) *SplFilesystemObject {
	return (*SplFilesystemObject)((*byte)(obj - zend_long((*byte)(&((*SplFilesystemObject)(nil).GetStd()))-(*byte)(nil))))
}
func Z_SPLFILESYSTEM_P(zv *zend.Zval) *SplFilesystemObject { return SplFilesystemFromObj(zv.GetObj()) }
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
func SplFilesystemObjectDestroyObject(object *zend.ZendObject) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(object)
	zend.ZendObjectsDestroyObject(object)
	switch intern.GetType() {
	case SPL_FS_DIR:
		if intern.GetDirp() != nil {
			core.PhpStreamClose(intern.GetDirp())
			intern.SetDirp(nil)
		}
		break
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
		break
	default:
		break
	}
}
func SplFilesystemObjectFreeStorage(object *zend.ZendObject) {
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
		break
	case SPL_FS_DIR:
		if intern.GetSubPath() != nil {
			zend.Efree(intern.GetSubPath())
		}
		break
	case SPL_FS_FILE:
		if intern.GetOpenMode() != nil {
			zend.Efree(intern.GetOpenMode())
		}
		if intern.GetOrigPath() != nil {
			zend.Efree(intern.GetOrigPath())
		}
		SplFilesystemFileFreeLine(intern)
		break
	}
}
func SplFilesystemObjectNewEx(class_type *zend.ZendClassEntry) *zend.ZendObject {
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
func SplFilesystemObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplFilesystemObjectNewEx(class_type)
}
func SplFilesystemObjectNewCheck(class_type *zend.ZendClassEntry) *zend.ZendObject {
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

	case SPL_FS_FILE:
		if intern.GetFileName() == nil {
			core.PhpErrorDocref(nil, zend.E_ERROR, "Object not initialized")
		}
		break
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

		break
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

			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Failed to open directory \"%s\"", path)

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
	var tmp zend.Zval
	intern.SetType(SPL_FS_FILE)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, &tmp)
	if tmp.IsType(zend.IS_TRUE) {
		intern.SetOpenMode(nil)
		intern.SetFileName(nil)
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Cannot use SplFileObject with directories")
		return zend.FAILURE
	}
	intern.SetContext(streams.PhpStreamContextFromZval(intern.GetZcontext(), 0))
	intern.SetStream(core.PhpStreamOpenWrapperEx(intern.GetFileName(), intern.GetOpenMode(), b.Cond(use_include_path != 0, core.USE_PATH, 0)|core.REPORT_ERRORS, nil, intern.GetContext()))
	if intern.GetFileNameLen() == 0 || intern.GetStream() == nil {
		if zend.EG__().GetException() == nil {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot open file '%s'", b.CondF1(intern.GetFileNameLen() != 0, func() *byte { return intern.GetFileName() }, ""))
		}
		intern.SetFileName(nil)
		intern.SetOpenMode(nil)
		return zend.FAILURE
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

	zend.ZVAL_RES(intern.GetZresource(), intern.GetStream().GetRes())

	/*!!! TODO: maybe bug?
	  Z_SET_REFCOUNT(intern->u.file.zresource, 1);
	*/

	intern.SetDelimiter(',')
	intern.SetEnclosure('"')
	intern.SetEscape(uint8('\\'))
	intern.SetFuncGetCurr(zend.ZendHashStrFindPtr(intern.GetStd().GetCe().GetFunctionTable(), "getcurrentline", b.SizeOf("\"getcurrentline\"")-1))
	return zend.SUCCESS
}
func SplFilesystemObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
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
		break
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
		break
	case SPL_FS_FILE:
		zend.ZEND_ASSERT(false)
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
func SplFilesystemObjectCreateInfo(source *SplFilesystemObject, file_path *byte, file_path_len int, use_copy int, ce *zend.ZendClassEntry, return_value *zend.Zval) *SplFilesystemObject {
	var intern *SplFilesystemObject
	var arg1 zend.Zval
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
	zend.ZVAL_OBJ(return_value, intern.GetStd())
	if ce.GetConstructor().GetScope() != spl_ce_SplFileInfo {
		zend.ZVAL_STRINGL(&arg1, file_path, file_path_len)
		zend.ZendCallMethodWith1Params(return_value, ce, ce.GetConstructor(), "__construct", nil, &arg1)
		zend.ZvalPtrDtor(&arg1)
	} else {
		SplFilesystemInfoSetFilename(intern, file_path, file_path_len, use_copy)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
	return intern
}
func SplFilesystemObjectCreateType(ht int, source *SplFilesystemObject, type_ int, ce *zend.ZendClassEntry, return_value *zend.Zval) *SplFilesystemObject {
	var intern *SplFilesystemObject
	var use_include_path zend.ZendBool = 0
	var arg1 zend.Zval
	var arg2 zend.Zval
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	switch source.GetType() {
	case SPL_FS_INFO:

	case SPL_FS_FILE:
		break
	case SPL_FS_DIR:
		if !(source.GetEntry().GetDName()[0]) {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Could not open file")
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
		if zend.ZendUpdateClassConstants(ce) != zend.SUCCESS {
			break
		}
		intern = SplFilesystemFromObj(SplFilesystemObjectNewEx(ce))
		zend.ZVAL_OBJ(return_value, intern.GetStd())
		SplFilesystemObjectGetFileName(source)
		if ce.GetConstructor().GetScope() != spl_ce_SplFileInfo {
			zend.ZVAL_STRINGL(&arg1, source.GetFileName(), source.GetFileNameLen())
			zend.ZendCallMethodWith1Params(return_value, ce, ce.GetConstructor(), "__construct", nil, &arg1)
			zend.ZvalPtrDtor(&arg1)
		} else {
			intern.SetFileName(zend.Estrndup(source.GetFileName(), source.GetFileNameLen()))
			intern.SetFileNameLen(source.GetFileNameLen())
			intern.SetPath(SplFilesystemObjectGetPath(source, intern.GetPathLen()))
			intern.SetPath(zend.Estrndup(intern.GetPath(), intern.GetPathLen()))
		}
		break
	case SPL_FS_FILE:
		if ce != nil {
			ce = ce
		} else {
			ce = source.GetFileClass()
		}
		if zend.ZendUpdateClassConstants(ce) != zend.SUCCESS {
			break
		}
		intern = SplFilesystemFromObj(SplFilesystemObjectNewEx(ce))
		zend.ZVAL_OBJ(return_value, intern.GetStd())
		SplFilesystemObjectGetFileName(source)
		if ce.GetConstructor().GetScope() != spl_ce_SplFileObject {
			zend.ZVAL_STRINGL(&arg1, source.GetFileName(), source.GetFileNameLen())
			zend.ZVAL_STRINGL(&arg2, "r", 1)
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
			if ht != 0 && zend.ZendParseParameters(ht, "|sbr", intern.GetOpenMode(), intern.GetOpenModeLen(), &use_include_path, intern.GetZcontext()) == zend.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				intern.SetOpenMode(nil)
				intern.SetFileName(nil)
				zend.ZvalPtrDtor(return_value)
				return_value.SetNull()
				return nil
			}
			if SplFilesystemFileOpen(intern, use_include_path, 0) == zend.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				zend.ZvalPtrDtor(return_value)
				return_value.SetNull()
				return nil
			}
		}
		break
	case SPL_FS_DIR:
		zend.ZendRestoreErrorHandling(&error_handling)
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Operation not supported")
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
func SplFilesystemObjectGetDebugInfo(object *zend.Zval) *zend.HashTable {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(object)
	var tmp zend.Zval
	var rv *zend.HashTable
	var pnstr *zend.ZendString
	var path *byte
	var path_len int
	var stmp []byte
	if intern.GetStd().GetProperties() == nil {
		zend.RebuildObjectProperties(intern.GetStd())
	}
	rv = zend.ZendArrayDup(intern.GetStd().GetProperties())
	pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "pathName", b.SizeOf("\"pathName\"")-1)
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	zend.ZVAL_STRINGL(&tmp, b.Cond(path != nil, path, ""), path_len)
	zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	if intern.GetFileName() != nil {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "fileName", b.SizeOf("\"fileName\"")-1)
		SplFilesystemObjectGetPath(intern, &path_len)
		if path_len != 0 && path_len < intern.GetFileNameLen() {
			zend.ZVAL_STRINGL(&tmp, intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1))
		} else {
			zend.ZVAL_STRINGL(&tmp, intern.GetFileName(), intern.GetFileNameLen())
		}
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_DIR {
		pnstr = SplGenPrivatePropName(spl_ce_DirectoryIterator, "glob", b.SizeOf("\"glob\"")-1)
		if core.PhpStreamIs(intern.GetDirp(), &streams.PhpGlobStreamOps) {
			zend.ZVAL_STRINGL(&tmp, intern.GetPath(), intern.GetPathLen())
		} else {
			tmp.SetFalse()
		}
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		pnstr = SplGenPrivatePropName(spl_ce_RecursiveDirectoryIterator, "subPathName", b.SizeOf("\"subPathName\"")-1)
		if intern.GetSubPath() != nil {
			zend.ZVAL_STRINGL(&tmp, intern.GetSubPath(), intern.GetSubPathLen())
		} else {
			zend.ZVAL_EMPTY_STRING(&tmp)
		}
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_FILE {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "openMode", b.SizeOf("\"openMode\"")-1)
		zend.ZVAL_STRINGL(&tmp, intern.GetOpenMode(), intern.GetOpenModeLen())
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		stmp[1] = '0'
		stmp[0] = intern.GetDelimiter()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "delimiter", b.SizeOf("\"delimiter\"")-1)
		zend.ZVAL_STRINGL(&tmp, stmp, 1)
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		stmp[0] = intern.GetEnclosure()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "enclosure", b.SizeOf("\"enclosure\"")-1)
		zend.ZVAL_STRINGL(&tmp, stmp, 1)
		zend.ZendSymtableUpdate(rv, pnstr.GetStr(), &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	return rv
}
func SplFilesystemObjectGetMethodCheck(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var fsobj *SplFilesystemObject = SplFilesystemFromObj(*object)
	if fsobj.GetDirp() == nil && fsobj.GetOrigPath() == nil {
		var func_ *zend.ZendFunction
		var tmp *zend.ZendString = zend.ZendStringInit("_bad_state_ex", b.SizeOf("\"_bad_state_ex\"")-1, 0)
		func_ = zend.ZendStdGetMethod(object, tmp, nil)
		zend.ZendStringReleaseEx(tmp, 0)
		return func_
	}
	return zend.ZendStdGetMethod(object, method, key)
}
func SplFilesystemObjectConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ctor_flags zend.ZendLong) {
	var intern *SplFilesystemObject
	var path *byte
	var parsed int
	var len_ int
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if SPL_HAS_FLAG(ctor_flags, DIT_CTOR_FLAGS) != 0 {
		flags = SPL_FILE_DIR_KEY_AS_PATHNAME | SPL_FILE_DIR_CURRENT_AS_FILEINFO
		parsed = zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "p|l", &path, &len_, &flags)
	} else {
		flags = SPL_FILE_DIR_KEY_AS_PATHNAME | SPL_FILE_DIR_CURRENT_AS_SELF
		parsed = zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "p", &path, &len_)
	}
	if SPL_HAS_FLAG(ctor_flags, SPL_FILE_DIR_SKIPDOTS) != 0 {
		flags |= SPL_FILE_DIR_SKIPDOTS
	}
	if SPL_HAS_FLAG(ctor_flags, SPL_FILE_DIR_UNIXPATHS) != 0 {
		flags |= SPL_FILE_DIR_UNIXPATHS
	}
	if parsed == zend.FAILURE {
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	if len_ == 0 {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Directory name must not be empty.")
		zend.ZendRestoreErrorHandling(&error_handling)
		return
	}
	intern = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if intern.GetPath() != nil {

		/* object is already initialized */

		zend.ZendRestoreErrorHandling(&error_handling)
		core.PhpErrorDocref(nil, zend.E_WARNING, "Directory object is already initialized")
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
func zim_spl_DirectoryIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, 0)
}
func zim_spl_DirectoryIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	intern.SetIndex(0)
	if intern.GetDirp() != nil {
		core.PhpStreamRewinddir(intern.GetDirp())
	}
	SplFilesystemDirRead(intern)
}
func zim_spl_DirectoryIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetDirp() != nil {
		zend.RETVAL_LONG(intern.GetIndex())
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func zim_spl_DirectoryIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZVAL_OBJ(return_value, zend.ZEND_THIS.GetObj())
	return_value.AddRefcount()
}
func zim_spl_DirectoryIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var skip_dots int = SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
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
func zim_spl_DirectoryIterator_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var retval zend.Zval
	var pos zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &pos) == zend.FAILURE {
		return
	}
	if intern.GetIndex() > pos {

		/* we first rewind */

		zend.ZendCallMethodWith0Params(zend.ZEND_THIS, zend.Z_OBJCE_P(zend.ZEND_THIS), intern.GetFuncRewind(), "rewind", nil)

		/* we first rewind */

	}
	for intern.GetIndex() < pos {
		var valid int = 0
		zend.ZendCallMethodWith0Params(zend.ZEND_THIS, zend.Z_OBJCE_P(zend.ZEND_THIS), intern.GetFuncValid(), "valid", &retval)
		valid = zend.ZendIsTrue(&retval)
		zend.ZvalPtrDtor(&retval)
		if valid == 0 {
			zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+zend.ZEND_LONG_FMT+" is out of range", pos)
			return
		}
		zend.ZendCallMethodWith0Params(zend.ZEND_THIS, zend.Z_OBJCE_P(zend.ZEND_THIS), intern.GetFuncNext(), "next", nil)
	}
}
func zim_spl_DirectoryIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(intern.GetEntry().GetDName()[0] != '0')
	return
}
func zim_spl_SplFileInfo_getPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var path *byte
	var path_len int
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPath(intern, &path_len)
	if path != nil {
		zend.RETVAL_STRINGL(path, path_len)
		return
	} else {
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func zim_spl_SplFileInfo_getFilename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var path_len int
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplFilesystemObjectGetPath(intern, &path_len)
	if path_len != 0 && path_len < intern.GetFileNameLen() {
		zend.RETVAL_STRINGL(intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1))
		return
	} else {
		zend.RETVAL_STRINGL(intern.GetFileName(), intern.GetFileNameLen())
		return
	}
}
func zim_spl_DirectoryIterator_getFilename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_STRING(intern.GetEntry().GetDName())
	return
}
func zim_spl_SplFileInfo_getExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var fname *byte = nil
	var p *byte
	var flen int
	var path_len int
	var idx int
	var ret *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
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
		zend.RETVAL_STRINGL(ret.GetVal()+idx+1, ret.GetLen()-idx-1)
		zend.ZendStringReleaseEx(ret, 0)
		return
	} else {
		zend.ZendStringReleaseEx(ret, 0)
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func zim_spl_DirectoryIterator_getExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var p *byte
	var idx int
	var fname *zend.ZendString
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.GetEntry().GetDName(), strlen(intern.GetEntry().GetDName()), nil, 0)
	p = zend.ZendMemrchr(fname.GetVal(), '.', fname.GetLen())
	if p != nil {
		idx = p - fname.GetVal()
		zend.RETVAL_STRINGL(fname.GetVal()+idx+1, fname.GetLen()-idx-1)
		zend.ZendStringReleaseEx(fname, 0)
	} else {
		zend.ZendStringReleaseEx(fname, 0)
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func zim_spl_SplFileInfo_getBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var fname *byte
	var suffix *byte = 0
	var flen int
	var slen int = 0
	var path_len int
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|s", &suffix, &slen) == zend.FAILURE {
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
	zend.RETVAL_STR(standard.PhpBasename(fname, flen, suffix, slen))
	return
}
func zim_spl_DirectoryIterator_getBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var suffix *byte = 0
	var slen int = 0
	var fname *zend.ZendString
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|s", &suffix, &slen) == zend.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.GetEntry().GetDName(), strlen(intern.GetEntry().GetDName()), suffix, slen)
	zend.RETVAL_STR(fname)
}
func zim_spl_SplFileInfo_getPathname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var path *byte
	var path_len int
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	if path != nil {
		zend.RETVAL_STRINGL(path, path_len)
		return
	} else {
		zend.RETVAL_FALSE
		return
	}
}
func zim_spl_FilesystemIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if SPL_FILE_DIR_KEY(intern, SPL_FILE_DIR_KEY_AS_FILENAME) {
		zend.RETVAL_STRING(intern.GetEntry().GetDName())
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		zend.RETVAL_STRINGL(intern.GetFileName(), intern.GetFileNameLen())
		return
	}
}
func zim_spl_FilesystemIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if SPL_FILE_DIR_CURRENT(intern, SPL_FILE_DIR_CURRENT_AS_PATHNAME) {
		SplFilesystemObjectGetFileName(intern)
		zend.RETVAL_STRINGL(intern.GetFileName(), intern.GetFileNameLen())
		return
	} else if SPL_FILE_DIR_CURRENT(intern, SPL_FILE_DIR_CURRENT_AS_FILEINFO) {
		SplFilesystemObjectGetFileName(intern)
		SplFilesystemObjectCreateType(0, intern, SPL_FS_INFO, nil, return_value)
	} else {
		zend.ZVAL_OBJ(return_value, zend.ZEND_THIS.GetObj())
		return_value.AddRefcount()
	}
}
func zim_spl_DirectoryIterator_isDot(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_BOOL(SplFilesystemIsDot(intern.GetEntry().GetDName()) != 0)
	return
}
func zim_spl_SplFileInfo___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject
	var path *byte
	var len_ int
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "p", &path, &len_) == zend.FAILURE {
		return
	}
	intern = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	SplFilesystemInfoSetFilename(intern, path, len_, 1)
}
func zim_spl_SplFileInfo_getPerms(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_PERMS, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getInode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_INODE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_SIZE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getOwner(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_OWNER, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getGroup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_GROUP, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getATime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_ATIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getMTime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_MTIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getCTime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_CTIME, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getType(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_TYPE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isWritable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_W, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isReadable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_R, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isExecutable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_X, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_FILE, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isDir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_isLink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_LINK, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getLinkTarget(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ret ssize_t
	var buff []byte
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if intern.GetFileName() == nil {
		SplFilesystemObjectGetFileName(intern)
	}
	if intern.GetFileName() == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Empty filename")
		zend.RETVAL_FALSE
		return
	} else if !(zend.IS_ABSOLUTE_PATH(intern.GetFileName(), intern.GetFileNameLen())) {
		var expanded_path []byte
		if core.ExpandFilepathWithMode(intern.GetFileName(), expanded_path, nil, 0, zend.CWD_EXPAND) == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "No such file or directory")
			zend.RETVAL_FALSE
			return
		}
		ret = zend.PhpSysReadlink(expanded_path, buff, core.MAXPATHLEN-1)
	} else {
		ret = zend.PhpSysReadlink(intern.GetFileName(), buff, core.MAXPATHLEN-1)
	}
	if ret == -1 {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Unable to read link %s, error: %s", intern.GetFileName(), strerror(errno))
		zend.RETVAL_FALSE
	} else {

		/* Append NULL to the end of the string */

		buff[ret] = '0'
		zend.RETVAL_STRINGL(buff, ret)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getRealPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var buff []byte
	var filename *byte
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersNone() == zend.FAILURE {
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
		zend.RETVAL_STRING(buff)
	} else {
		zend.RETVAL_FALSE
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_openFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	SplFilesystemObjectCreateType(zend.ZEND_NUM_ARGS(), intern, SPL_FS_FILE, nil, return_value)
}
func zim_spl_SplFileInfo_setFileClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry = spl_ce_SplFileObject
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|C", &ce) == zend.SUCCESS {
		intern.SetFileClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_setInfoClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry = spl_ce_SplFileInfo
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|C", &ce) == zend.SUCCESS {
		intern.SetInfoClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getFileInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|C", &ce) == zend.SUCCESS {
		SplFilesystemObjectCreateType(zend.ZEND_NUM_ARGS(), intern, SPL_FS_INFO, ce, return_value)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileInfo_getPathInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ce *zend.ZendClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|C", &ce) == zend.SUCCESS {
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
func zim_spl_SplFileInfo___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_ARR(SplFilesystemObjectGetDebugInfo(zend.getThis()))
	return
}
func zim_spl_SplFileInfo__bad_state_ex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The parent constructor was not called: the object is in an "+"invalid state ")
}
func zim_spl_FilesystemIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, DIT_CTOR_FLAGS|SPL_FILE_DIR_SKIPDOTS)
}
func zim_spl_FilesystemIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var skip_dots int = SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_SKIPDOTS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
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
func zim_spl_FilesystemIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetFlags() & (SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK))
	return
}
func zim_spl_FilesystemIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var flags zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &flags) == zend.FAILURE {
		return
	}
	intern.SubFlags(SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK)
	intern.AddFlags((SPL_FILE_DIR_KEY_MODE_MASK | SPL_FILE_DIR_CURRENT_MODE_MASK | SPL_FILE_DIR_OTHERS_MASK) & flags)
}
func zim_spl_RecursiveDirectoryIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var allow_links zend.ZendBool = 0
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|b", &allow_links) == zend.FAILURE {
		return
	}
	if SplFilesystemIsInvalidOrDot(intern.GetEntry().GetDName()) != 0 {
		zend.RETVAL_FALSE
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		if allow_links == 0 && !intern.IsDirFollowSymlinks() {
			standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_LINK, return_value)
			if zend.ZendIsTrue(return_value) != 0 {
				zend.RETVAL_FALSE
				return
			}
		}
		standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), standard.FS_IS_DIR, return_value)
	}
}
func zim_spl_RecursiveDirectoryIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zpath zend.Zval
	var zflags zend.Zval
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var subdir *SplFilesystemObject
	var slash byte = b.Cond(SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_UNIXPATHS) != 0, '/', zend.DEFAULT_SLASH)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplFilesystemObjectGetFileName(intern)
	zflags.SetLong(intern.GetFlags())
	zend.ZVAL_STRINGL(&zpath, intern.GetFileName(), intern.GetFileNameLen())
	SplInstantiateArgEx2(zend.Z_OBJCE_P(zend.ZEND_THIS), return_value, &zpath, &zflags)
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
func zim_spl_RecursiveDirectoryIterator_getSubPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		zend.RETVAL_STRINGL(intern.GetSubPath(), intern.GetSubPathLen())
		return
	} else {
		zend.RETVAL_EMPTY_STRING()
		return
	}
}
func zim_spl_RecursiveDirectoryIterator_getSubPathname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var slash byte = b.Cond(SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_DIR_UNIXPATHS) != 0, '/', zend.DEFAULT_SLASH)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		zend.RETVAL_NEW_STR(core.Strpprintf(0, "%s%c%s", intern.GetSubPath(), slash, intern.GetEntry().GetDName()))
		return
	} else {
		zend.RETVAL_STRING(intern.GetEntry().GetDName())
		return
	}
}
func zim_spl_RecursiveDirectoryIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, DIT_CTOR_FLAGS)
}
func zim_spl_GlobIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, DIT_CTOR_FLAGS|DIT_CTOR_GLOB)
}
func zim_spl_GlobIterator_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetDirp() != nil && core.PhpStreamIs(intern.GetDirp(), &streams.PhpGlobStreamOps) {
		zend.RETVAL_LONG(streams.PhpGlobStreamGetCount(intern.GetDirp(), nil))
		return
	} else {

		/* should not happen */

		core.PhpErrorDocref(nil, zend.E_ERROR, "GlobIterator lost glob state")

		/* should not happen */

	}
}
func SplFilesystemDirGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = Z_SPLFILESYSTEM_P(object)
	iterator = SplFilesystemObjectToIterator(dir_object)
	object.AddRefcount()
	zend.ZVAL_OBJ(iterator.GetIntern().GetData(), object.GetObj())
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
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}
func SplFilesystemDirItCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	return iterator.GetCurrent()
}
func SplFilesystemDirItCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
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
func SplFilesystemTreeItCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	if SPL_FILE_DIR_CURRENT(object, SPL_FILE_DIR_CURRENT_AS_PATHNAME) {
		if iterator.GetCurrent().IsUndef() {
			SplFilesystemObjectGetFileName(object)
			zend.ZVAL_STRINGL(iterator.GetCurrent(), object.GetFileName(), object.GetFileNameLen())
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
func SplFilesystemTreeItCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	if SPL_FILE_DIR_KEY(object, SPL_FILE_DIR_KEY_AS_FILENAME) {
		zend.ZVAL_STRING(key, object.GetEntry().GetDName())
	} else {
		SplFilesystemObjectGetFileName(object)
		zend.ZVAL_STRINGL(key, object.GetFileName(), object.GetFileNameLen())
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
func SplFilesystemTreeGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = Z_SPLFILESYSTEM_P(object)
	iterator = SplFilesystemObjectToIterator(dir_object)
	object.AddRefcount()
	zend.ZVAL_OBJ(iterator.GetIntern().GetData(), object.GetObj())
	iterator.GetIntern().SetFuncs(&SplFilesystemTreeItFuncs)
	return iterator.GetIntern()
}
func SplFilesystemObjectCast(readobj *zend.Zval, writeobj *zend.Zval, type_ int) int {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(readobj)
	if type_ == zend.IS_STRING {
		if zend.Z_OBJCE_P(readobj).GetTostring() != nil {
			return zend.ZendStdCastObjectTostring(readobj, writeobj, type_)
		}
		switch intern.GetType() {
		case SPL_FS_INFO:

		case SPL_FS_FILE:
			zend.ZVAL_STRINGL(writeobj, intern.GetFileName(), intern.GetFileNameLen())
			return zend.SUCCESS
		case SPL_FS_DIR:
			zend.ZVAL_STRING(writeobj, intern.GetEntry().GetDName())
			return zend.SUCCESS
		}
	} else if type_ == zend._IS_BOOL {
		writeobj.SetTrue()
		return zend.SUCCESS
	}
	writeobj.SetNull()
	return zend.FAILURE
}
func SplFilesystemFileRead(intern *SplFilesystemObject, silent int) int {
	var buf *byte
	var line_len int = 0
	var line_add zend.ZendLong = b.Cond(intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()), 1, 0)
	SplFilesystemFileFreeLine(intern)
	if core.PhpStreamEof(intern.GetStream()) != 0 {
		if silent == 0 {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
		}
		return zend.FAILURE
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
	return zend.SUCCESS
}
func SplFilesystemFileCall(intern *SplFilesystemObject, func_ptr *zend.ZendFunction, pass_num_args int, return_value *zend.Zval, arg2 *zend.Zval) int {
	var fci zend.ZendFcallInfo
	var fcic zend.ZendFcallInfoCache
	var zresource_ptr *zend.Zval = intern.GetZresource()
	var params *zend.Zval
	var retval zend.Zval
	var result int
	var num_args int = pass_num_args + b.Cond(arg2 != nil, 2, 1)
	if zresource_ptr.IsUndef() {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return zend.FAILURE
	}
	params = (*zend.Zval)(zend.SafeEmalloc(num_args, b.SizeOf("zval"), 0))
	params[0] = *zresource_ptr
	if arg2 != nil {
		params[1] = *arg2
	}
	if zend.ZendGetParametersArrayEx(pass_num_args, params+b.Cond(arg2 != nil, 2, 1)) != zend.SUCCESS {
		zend.Efree(params)
		zend.WRONG_PARAM_COUNT_WITH_RETVAL(zend.FAILURE)
	}
	retval.SetUndef()
	fci.SetSize(b.SizeOf("fci"))
	fci.SetObject(nil)
	fci.SetRetval(&retval)
	fci.SetParamCount(num_args)
	fci.SetParams(params)
	fci.SetNoSeparation(1)
	zend.ZVAL_STR(fci.GetFunctionName(), func_ptr.GetFunctionName())
	fcic.SetFunctionHandler(func_ptr)
	fcic.SetCalledScope(nil)
	fcic.SetObject(nil)
	result = zend.ZendCallFunction(&fci, &fcic)
	if result == zend.FAILURE || retval.IsUndef() {
		zend.RETVAL_FALSE
	} else {
		zend.ZVAL_ZVAL(return_value, &retval, 0, 0)
	}
	zend.Efree(params)
	return result
}
func SplFilesystemFileReadCsv(intern *SplFilesystemObject, delimiter byte, enclosure byte, escape int, return_value *zend.Zval) int {
	var ret int = zend.SUCCESS
	var value *zend.Zval
	for {
		ret = SplFilesystemFileRead(intern, 1)
		if !(ret == zend.SUCCESS && intern.GetCurrentLineLen() == 0 && SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_SKIP_EMPTY) != 0) {
			break
		}
	}
	if ret == zend.SUCCESS {
		var buf_len int = intern.GetCurrentLineLen()
		var buf *byte = zend.Estrndup(intern.GetCurrentLine(), buf_len)
		if !(intern.GetCurrentZval().IsUndef()) {
			zend.ZvalPtrDtor(intern.GetCurrentZval())
			intern.GetCurrentZval().SetUndef()
		}
		standard.PhpFgetcsv(intern.GetStream(), delimiter, enclosure, escape, buf_len, buf, intern.GetCurrentZval())
		if return_value != nil {
			value = intern.GetCurrentZval()
			zend.ZVAL_COPY_DEREF(return_value, value)
		}
	}
	return ret
}
func SplFilesystemFileReadLineEx(this_ptr *zend.Zval, intern *SplFilesystemObject, silent int) int {
	var retval zend.Zval

	/* 1) use fgetcsv? 2) overloaded call the function, 3) do it directly */

	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 || intern.GetFuncGetCurr().GetScope() != spl_ce_SplFileObject {
		if core.PhpStreamEof(intern.GetStream()) != 0 {
			if silent == 0 {
				zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
			}
			return zend.FAILURE
		}
		if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 {
			return SplFilesystemFileReadCsv(intern, intern.GetDelimiter(), intern.GetEnclosure(), intern.GetEscape(), nil)
		} else {
			var execute_data *zend.ZendExecuteData = zend.EG__().GetCurrentExecuteData()
			zend.ZendCallMethodWith0Params(this_ptr, zend.Z_OBJCE_P(zend.ZEND_THIS), intern.GetFuncGetCurr(), "getCurrentLine", &retval)
		}
		if !(retval.IsUndef()) {
			if intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()) {
				intern.GetCurrentLineNum()++
			}
			SplFilesystemFileFreeLine(intern)
			if retval.IsType(zend.IS_STRING) {
				intern.SetCurrentLine(zend.Estrndup(zend.Z_STRVAL(retval), zend.Z_STRLEN(retval)))
				intern.SetCurrentLineLen(zend.Z_STRLEN(retval))
			} else {
				var value *zend.Zval = &retval
				zend.ZVAL_COPY_DEREF(intern.GetCurrentZval(), value)
			}
			zend.ZvalPtrDtor(&retval)
			return zend.SUCCESS
		} else {
			return zend.FAILURE
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
		case zend.IS_STRING:
			return zend.Z_STRLEN(intern.GetCurrentZval()) == 0
		case zend.IS_ARRAY:
			if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) != 0 && zend.Z_ARRVAL(intern.GetCurrentZval()).GetNNumOfElements() == 1 {
				var idx uint32 = 0
				var first *zend.Zval
				for zend.Z_ARRVAL(intern.GetCurrentZval()).GetArData()[idx].GetVal().IsUndef() {
					idx++
				}
				first = zend.Z_ARRVAL(intern.GetCurrentZval()).GetArData()[idx].GetVal()
				return first.IsType(zend.IS_STRING) && zend.Z_STRLEN_P(first) == 0
			}
			return zend.Z_ARRVAL(intern.GetCurrentZval()).GetNNumOfElements() == 0
		case zend.IS_NULL:
			return 1
		default:
			return 0
		}
	} else {
		return 1
	}
}
func SplFilesystemFileReadLine(this_ptr *zend.Zval, intern *SplFilesystemObject, silent int) int {
	var ret int = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	for SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_SKIP_EMPTY) != 0 && ret == zend.SUCCESS && SplFilesystemFileIsEmptyLine(intern) != 0 {
		SplFilesystemFileFreeLine(intern)
		ret = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	}
	return ret
}
func SplFilesystemFileRewind(this_ptr *zend.Zval, intern *SplFilesystemObject) {
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if -1 == core.PhpStreamRewind(intern.GetStream()) {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot rewind file %s", intern.GetFileName())
	} else {
		SplFilesystemFileFreeLine(intern)
		intern.SetCurrentLineNum(0)
	}
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		SplFilesystemFileReadLine(this_ptr, intern, 1)
	}
}
func zim_spl_SplFileObject___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var use_include_path zend.ZendBool = 0
	var p1 *byte
	var p2 *byte
	var tmp_path *byte
	var tmp_path_len int
	var error_handling zend.ZendErrorHandling
	intern.SetOpenMode(nil)
	intern.SetOpenModeLen(0)
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "p|sbr!", intern.GetFileName(), intern.GetFileNameLen(), intern.GetOpenMode(), intern.GetOpenModeLen(), &use_include_path, intern.GetZcontext()) == zend.FAILURE {
		intern.SetOpenMode(nil)
		intern.SetFileName(nil)
		return
	}
	if intern.GetOpenMode() == nil {
		intern.SetOpenMode("r")
		intern.SetOpenModeLen(1)
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if SplFilesystemFileOpen(intern, use_include_path, 0) == zend.SUCCESS {
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
func zim_spl_SplTempFileObject___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var max_memory zend.ZendLong = core.PHP_STREAM_MAX_MEM
	var tmp_fname []byte
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersThrow(zend.ZEND_NUM_ARGS(), "|l", &max_memory) == zend.FAILURE {
		return
	}
	if max_memory < 0 {
		intern.SetFileName("php://memory")
		intern.SetFileNameLen(12)
	} else if zend.ZEND_NUM_ARGS() != 0 {
		intern.SetFileNameLen(core.Slprintf(tmp_fname, b.SizeOf("tmp_fname"), "php://temp/maxmemory:"+zend.ZEND_LONG_FMT, max_memory))
		intern.SetFileName(tmp_fname)
	} else {
		intern.SetFileName("php://temp")
		intern.SetFileNameLen(10)
	}
	intern.SetOpenMode("wb")
	intern.SetOpenModeLen(1)
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if SplFilesystemFileOpen(intern, 0, 0) == zend.SUCCESS {
		intern.SetPathLen(0)
		intern.SetPath(zend.Estrndup("", 0))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}
func zim_spl_SplFileObject_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplFilesystemFileRewind(zend.ZEND_THIS, intern)
}
func zim_spl_SplFileObject_eof(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	zend.RETVAL_BOOL(core.PhpStreamEof(intern.GetStream()) != 0)
	return
}
func zim_spl_SplFileObject_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		zend.RETVAL_BOOL(intern.GetCurrentLine() != nil || !(intern.GetCurrentZval().IsUndef()))
		return
	} else {
		if intern.GetStream() == nil {
			zend.RETVAL_FALSE
			return
		}
		zend.RETVAL_BOOL(core.PhpStreamEof(intern.GetStream()) == 0)
	}
}
func zim_spl_SplFileObject_fgets(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if SplFilesystemFileRead(intern, 0) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STRINGL(intern.GetCurrentLine(), intern.GetCurrentLineLen())
	return
}
func zim_spl_SplFileObject_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if intern.GetCurrentLine() == nil && intern.GetCurrentZval().IsUndef() {
		SplFilesystemFileReadLine(zend.ZEND_THIS, intern, 1)
	}
	if intern.GetCurrentLine() != nil && (SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_CSV) == 0 || intern.GetCurrentZval().IsUndef()) {
		zend.RETVAL_STRINGL(intern.GetCurrentLine(), intern.GetCurrentLineLen())
		return
	} else if !(intern.GetCurrentZval().IsUndef()) {
		var value *zend.Zval = intern.GetCurrentZval()
		zend.ZVAL_COPY_DEREF(return_value, value)
		return
	}
	zend.RETVAL_FALSE
	return
}
func zim_spl_SplFileObject_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}

	/*    Do not read the next line to support correct counting with fgetc()
	      if (!intern->current_line) {
	          spl_filesystem_file_read_line(ZEND_THIS, intern, 1);
	      } */

	zend.RETVAL_LONG(intern.GetCurrentLineNum())
	return
}
func zim_spl_SplFileObject_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	SplFilesystemFileFreeLine(intern)
	if SPL_HAS_FLAG(intern.GetFlags(), SPL_FILE_OBJECT_READ_AHEAD) != 0 {
		SplFilesystemFileReadLine(zend.ZEND_THIS, intern, 1)
	}
	intern.GetCurrentLineNum()++
}
func zim_spl_SplFileObject_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", intern.GetFlags()) == zend.FAILURE {
		return
	}
}
func zim_spl_SplFileObject_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(intern.GetFlags() & SPL_FILE_OBJECT_MASK)
	return
}
func zim_spl_SplFileObject_setMaxLineLen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var max_len zend.ZendLong
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &max_len) == zend.FAILURE {
		return
	}
	if max_len < 0 {
		zend.ZendThrowExceptionEx(spl_ce_DomainException, 0, "Maximum line length must be greater than or equal zero")
		return
	}
	intern.SetMaxLineLen(max_len)
}
func zim_spl_SplFileObject_getMaxLineLen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_LONG(zend.ZendLong(intern.GetMaxLineLen()))
	return
}
func zim_spl_SplFileObject_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	zend.RETVAL_FALSE
	return
}
func zim_spl_SplFileObject_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
}
func zim_spl_SplFileObject_fgetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var delimiter byte = intern.GetDelimiter()
	var enclosure byte = intern.GetEnclosure()
	var escape int = intern.GetEscape()
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		if intern.GetStream() == nil {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
			return
		}
		switch zend.ZEND_NUM_ARGS() {
		case 3:
			if esc_len > 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "escape must be empty or a single character")
				zend.RETVAL_FALSE
				return
			}
			if esc_len == 0 {
				escape = standard.PHP_CSV_NO_ESCAPE
			} else {
				escape = uint8(esc[0])
			}
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "enclosure must be a character")
				zend.RETVAL_FALSE
				return
			}
			enclosure = enclo[0]
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "delimiter must be a character")
				zend.RETVAL_FALSE
				return
			}
			delimiter = delim[0]
		case 0:
			break
		}
		SplFilesystemFileReadCsv(intern, delimiter, enclosure, escape, return_value)
	}
}
func zim_spl_SplFileObject_fputcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
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
	var fields *zend.Zval = nil
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "a|sss", &fields, &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		switch zend.ZEND_NUM_ARGS() {
		case 4:
			switch esc_len {
			case 0:
				escape = standard.PHP_CSV_NO_ESCAPE
				break
			case 1:
				escape = uint8(esc[0])
				break
			default:
				core.PhpErrorDocref(nil, zend.E_WARNING, "escape must be empty or a single character")
				zend.RETVAL_FALSE
				return
			}
		case 3:
			if e_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "enclosure must be a character")
				zend.RETVAL_FALSE
				return
			}
			enclosure = enclo[0]
		case 2:
			if d_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "delimiter must be a character")
				zend.RETVAL_FALSE
				return
			}
			delimiter = delim[0]
		case 1:

		case 0:
			break
		}
		ret = standard.PhpFputcsv(intern.GetStream(), fields, delimiter, enclosure, escape)
		if ret < 0 {
			zend.RETVAL_FALSE
			return
		}
		zend.RETVAL_LONG(ret)
		return
	}
}
func zim_spl_SplFileObject_setCsvControl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		switch zend.ZEND_NUM_ARGS() {
		case 3:
			switch esc_len {
			case 0:
				escape = standard.PHP_CSV_NO_ESCAPE
				break
			case 1:
				escape = uint8(esc[0])
				break
			default:
				core.PhpErrorDocref(nil, zend.E_WARNING, "escape must be empty or a single character")
				zend.RETVAL_FALSE
				return
			}
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "enclosure must be a character")
				zend.RETVAL_FALSE
				return
			}
			enclosure = enclo[0]
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "delimiter must be a character")
				zend.RETVAL_FALSE
				return
			}
			delimiter = delim[0]
		case 0:
			break
		}
		intern.SetDelimiter(delimiter)
		intern.SetEnclosure(enclosure)
		intern.SetEscape(escape)
	}
}
func zim_spl_SplFileObject_getCsvControl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
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
func zim_spl_SplFileObject_flock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "flock", b.SizeOf("\"flock\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "flock")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, zend.ZEND_NUM_ARGS(), return_value, nil)
}
func zim_spl_SplFileObject_fflush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	zend.RETVAL_BOOL(core.PhpStreamFlush(intern.GetStream()) == 0)
	return
}
func zim_spl_SplFileObject_ftell(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var ret zend.ZendLong
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	ret = core.PhpStreamTell(intern.GetStream())
	if ret == -1 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(ret)
		return
	}
}
func zim_spl_SplFileObject_fseek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var pos zend.ZendLong
	var whence zend.ZendLong = r.SEEK_SET
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l|l", &pos, &whence) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	zend.RETVAL_LONG(core.PhpStreamSeek(intern.GetStream(), pos, int(whence)))
	return
}
func zim_spl_SplFileObject_fgetc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var buf []byte
	var result int
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	result = core.PhpStreamGetc(intern.GetStream())
	if result == r.EOF {
		zend.RETVAL_FALSE
	} else {
		if result == '\n' {
			intern.GetCurrentLineNum()++
		}
		buf[0] = result
		buf[1] = '0'
		zend.RETVAL_STRINGL(buf, 1)
		return
	}
}
func zim_spl_SplFileObject_fgetss(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var arg2 zend.Zval
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
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
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fgetss", b.SizeOf("\"fgetss\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fgetss")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, zend.ZEND_NUM_ARGS(), return_value, &arg2)
}
func zim_spl_SplFileObject_fpassthru(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	zend.RETVAL_LONG(core.PhpStreamPassthru(intern.GetStream()))
	return
}
func zim_spl_SplFileObject_fscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	intern.GetCurrentLineNum()++
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fscanf", b.SizeOf("\"fscanf\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fscanf")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, zend.ZEND_NUM_ARGS(), return_value, nil)
}
func zim_spl_SplFileObject_fwrite(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var str *byte
	var str_len int
	var length zend.ZendLong = 0
	var written ssize_t
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "s|l", &str, &str_len, &length) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if zend.ZEND_NUM_ARGS() > 1 {
		if length >= 0 {
			str_len = cli.MIN(int(length), str_len)
		} else {

			/* Negative length given, nothing to write */

			str_len = 0

			/* Negative length given, nothing to write */

		}
	}
	if str_len == 0 {
		zend.RETVAL_LONG(0)
		return
	}
	written = core.PhpStreamWrite(intern.GetStream(), str, str_len)
	if written < 0 {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_LONG(written)
	return
}
func zim_spl_SplFileObject_fread(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var length zend.ZendLong = 0
	var str *zend.ZendString
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &length) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if length <= 0 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Length parameter must be greater than 0")
		zend.RETVAL_FALSE
		return
	}
	str = streams.PhpStreamReadToStr(intern.GetStream(), length)
	if str == nil {
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_STR(str)
	return
}
func zim_spl_SplFileObject_fstat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG__().GetFunctionTable(), "fstat", b.SizeOf("\"fstat\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fstat")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, zend.ZEND_NUM_ARGS(), return_value, nil)
}
func zim_spl_SplFileObject_ftruncate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var size zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &size) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if core.PhpStreamTruncateSupported(intern.GetStream()) == 0 {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Can't truncate file %s", intern.GetFileName())
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_BOOL(0 == core.PhpStreamTruncateSetSize(intern.GetStream(), size))
	return
}
func zim_spl_SplFileObject_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = Z_SPLFILESYSTEM_P(zend.ZEND_THIS)
	var line_pos zend.ZendLong
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "l", &line_pos) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if line_pos < 0 {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Can't seek file %s to negative line "+zend.ZEND_LONG_FMT, intern.GetFileName(), line_pos)
		zend.RETVAL_FALSE
		return
	}
	SplFilesystemFileRewind(zend.ZEND_THIS, intern)
	for intern.GetCurrentLineNum() < line_pos {
		if SplFilesystemFileReadLine(zend.ZEND_THIS, intern, 1) == zend.FAILURE {
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
	return zend.SUCCESS
}
