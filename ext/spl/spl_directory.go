// <<generate>>

package spl

import (
	"sik/core"
	"sik/core/streams"
	"sik/ext/standard"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/spl/spl_directory.h>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define SPL_DIRECTORY_H

// # include "php.h"

// # include "php_spl.h"

var spl_ce_SplFileInfo *zend.ZendClassEntry
var spl_ce_DirectoryIterator *zend.ZendClassEntry
var spl_ce_FilesystemIterator *zend.ZendClassEntry
var spl_ce_RecursiveDirectoryIterator *zend.ZendClassEntry
var spl_ce_GlobIterator *zend.ZendClassEntry
var spl_ce_SplFileObject *zend.ZendClassEntry
var spl_ce_SplTempFileObject *zend.ZendClassEntry

type SPL_FS_OBJ_TYPE = int

const (
	SPL_FS_INFO = iota
	SPL_FS_DIR
	SPL_FS_FILE
)

type SplForeignDtorT func(object *SplFilesystemObject)
type SplForeignCloneT func(src *SplFilesystemObject, dst *SplFilesystemObject)

// @type SplOtherHandler struct

/* define an __special__  overloaded iterator structure */

// @type SplFilesystemIterator struct
// @type SplFilesystemObject struct
func SplFilesystemFromObj(obj *zend.ZendObject) *SplFilesystemObject {
	return (*SplFilesystemObject)((*byte)(obj - zend_long((*byte)(&((*SplFilesystemObject)(nil).GetStd()))-(*byte)(nil))))
}

/* }}} */

// #define Z_SPLFILESYSTEM_P(zv) spl_filesystem_from_obj ( Z_OBJ_P ( ( zv ) ) )

func SplFilesystemObjectToIterator(obj *SplFilesystemObject) *SplFilesystemIterator {
	var it *SplFilesystemIterator
	it = zend._ecalloc(1, g.SizeOf("spl_filesystem_iterator"))
	it.SetObject(any(obj))
	zend.ZendIteratorInit(&it.intern)
	return it
}
func SplFilesystemIteratorToObject(it *SplFilesystemIterator) *SplFilesystemObject {
	return (*SplFilesystemObject)(it.GetObject())
}

// #define SPL_FILE_OBJECT_DROP_NEW_LINE       0x00000001

// #define SPL_FILE_OBJECT_READ_AHEAD       0x00000002

// #define SPL_FILE_OBJECT_SKIP_EMPTY       0x00000004

// #define SPL_FILE_OBJECT_READ_CSV       0x00000008

// #define SPL_FILE_OBJECT_MASK       0x0000000F

// #define SPL_FILE_DIR_CURRENT_AS_FILEINFO       0x00000000

// #define SPL_FILE_DIR_CURRENT_AS_SELF       0x00000010

// #define SPL_FILE_DIR_CURRENT_AS_PATHNAME       0x00000020

// #define SPL_FILE_DIR_CURRENT_MODE_MASK       0x000000F0

// #define SPL_FILE_DIR_CURRENT(intern,mode) ( ( intern -> flags & SPL_FILE_DIR_CURRENT_MODE_MASK ) == mode )

// #define SPL_FILE_DIR_KEY_AS_PATHNAME       0x00000000

// #define SPL_FILE_DIR_KEY_AS_FILENAME       0x00000100

// #define SPL_FILE_DIR_FOLLOW_SYMLINKS       0x00000200

// #define SPL_FILE_DIR_KEY_MODE_MASK       0x00000F00

// #define SPL_FILE_DIR_KEY(intern,mode) ( ( intern -> flags & SPL_FILE_DIR_KEY_MODE_MASK ) == mode )

// #define SPL_FILE_DIR_SKIPDOTS       0x00001000

// #define SPL_FILE_DIR_UNIXPATHS       0x00002000

// #define SPL_FILE_DIR_OTHERS_MASK       0x00003000

// Source: <ext/spl/spl_directory.c>

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
   | Author: Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_ini.h"

// # include "ext/standard/info.h"

// # include "ext/standard/file.h"

// # include "ext/standard/php_string.h"

// # include "zend_compile.h"

// # include "zend_exceptions.h"

// # include "zend_interfaces.h"

// # include "php_spl.h"

// # include "spl_functions.h"

// # include "spl_engine.h"

// # include "spl_iterators.h"

// # include "spl_directory.h"

// # include "spl_exceptions.h"

// # include "php.h"

// # include "fopen_wrappers.h"

// # include "ext/standard/basic_functions.h"

// # include "ext/standard/php_filestat.h"

// #define SPL_HAS_FLAG(flags,test_flag) ( ( flags & test_flag ) ? 1 : 0 )

/* declare the class handlers */

var SplFilesystemObjectHandlers zend.ZendObjectHandlers

/* includes handler to validate object state when retrieving methods */

var SplFilesystemObjectCheckHandlers zend.ZendObjectHandlers

/* decalre the class entry */

func SplFilesystemFileFreeLine(intern *SplFilesystemObject) {
	if intern.GetCurrentLine() != nil {
		zend._efree(intern.GetCurrentLine())
		intern.SetCurrentLine(nil)
	}
	if intern.u.file.current_zval.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&intern.u.file.current_zval)
		&intern.u.file.current_zval.u1.type_info = 0
	}
}
func SplFilesystemObjectDestroyObject(object *zend.ZendObject) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(object)
	zend.ZendObjectsDestroyObject(object)
	switch intern.GetType() {
	case SPL_FS_DIR:
		if intern.GetDirp() != nil {
			streams._phpStreamFree(intern.GetDirp(), 1|2)
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

			if intern.GetStream().is_persistent == 0 {
				streams._phpStreamFree(intern.GetStream(), 1|2)
			} else {
				streams._phpStreamFree(intern.GetStream(), 1|2|16)
			}
			intern.SetStream(nil)
			&intern.u.file.zresource.u1.type_info = 0
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
	zend.ZendObjectStdDtor(&intern.std)
	if intern.GetPath() != nil {
		zend._efree(intern.GetPath())
	}
	if intern.GetFileName() != nil {
		zend._efree(intern.GetFileName())
	}
	switch intern.GetType() {
	case SPL_FS_INFO:
		break
	case SPL_FS_DIR:
		if intern.GetSubPath() != nil {
			zend._efree(intern.GetSubPath())
		}
		break
	case SPL_FS_FILE:
		if intern.GetOpenMode() != nil {
			zend._efree(intern.GetOpenMode())
		}
		if intern.GetOrigPath() != nil {
			zend._efree(intern.GetOrigPath())
		}
		SplFilesystemFileFreeLine(intern)
		break
	}
}

/* {{{ spl_ce_dir_object_new */

func SplFilesystemObjectNewEx(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var intern *SplFilesystemObject
	intern = zend.ZendObjectAlloc(g.SizeOf("spl_filesystem_object"), class_type)

	/* intern->type = SPL_FS_INFO; done by set 0 */

	intern.SetFileClass(spl_ce_SplFileObject)
	intern.SetInfoClass(spl_ce_SplFileInfo)
	zend.ZendObjectStdInit(&intern.std, class_type)
	zend.ObjectPropertiesInit(&intern.std, class_type)
	intern.std.handlers = &SplFilesystemObjectHandlers
	return &intern.std
}

/* }}} */

func SplFilesystemObjectNew(class_type *zend.ZendClassEntry) *zend.ZendObject {
	return SplFilesystemObjectNewEx(class_type)
}

/* }}} */

func SplFilesystemObjectNewCheck(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var ret *SplFilesystemObject = SplFilesystemFromObj(SplFilesystemObjectNewEx(class_type))
	ret.std.handlers = &SplFilesystemObjectCheckHandlers
	return &ret.std
}

/* }}} */

func SplFilesystemObjectGetPath(intern *SplFilesystemObject, len_ *int) *byte {
	if intern.GetType() == SPL_FS_DIR {
		if intern.GetDirp().ops == &streams.PhpGlobStreamOps {
			return streams._phpGlobStreamGetPath(intern.GetDirp(), len_)
		}
	}
	if len_ != nil {
		*len_ = intern.GetPathLen()
	}
	return intern.GetPath()
}
func SplFilesystemObjectGetFileName(intern *SplFilesystemObject) {
	var slash byte = g.Cond(g.Cond((intern.GetFlags()&0x2000) != 0, 1, 0), '/', '/')
	switch intern.GetType() {
	case SPL_FS_INFO:

	case SPL_FS_FILE:
		if intern.GetFileName() == nil {
			core.PhpErrorDocref(nil, 1<<0, "Object not initialized")
		}
		break
	case SPL_FS_DIR:
		var path_len int = 0
		var path *byte = SplFilesystemObjectGetPath(intern, &path_len)
		if intern.GetFileName() != nil {
			zend._efree(intern.GetFileName())
		}

		/* if there is parent path, ammend it, otherwise just use the given path as is */

		if path_len == 0 {
			intern.SetFileNameLen(zend.ZendSpprintf(&intern.file_name, 0, "%s", intern.u.dir.entry.d_name))
		} else {
			intern.SetFileNameLen(zend.ZendSpprintf(&intern.file_name, 0, "%s%c%s", path, slash, intern.u.dir.entry.d_name))
		}

		/* if there is parent path, ammend it, otherwise just use the given path as is */

		break
	}
}
func SplFilesystemDirRead(intern *SplFilesystemObject) int {
	if intern.GetDirp() == nil || streams._phpStreamReaddir(intern.GetDirp(), &intern.u.dir.entry) == nil {
		intern.u.dir.entry.d_name[0] = '0'
		return 0
	} else {
		return 1
	}
}

/* }}} */

// #define IS_SLASH_AT(zs,pos) ( IS_SLASH ( zs [ pos ] ) )

func SplFilesystemIsDot(d_name *byte) int {
	return !(strcmp(d_name, ".")) || !(strcmp(d_name, ".."))
}

/* }}} */

func SplFilesystemDirOpen(intern *SplFilesystemObject, path *byte) {
	var skip_dots int = g.Cond((intern.GetFlags()&0x1000) != 0, 1, 0)
	intern.SetType(SPL_FS_DIR)
	intern.SetPathLen(strlen(path))
	intern.SetDirp(streams._phpStreamOpendir(path, 0x8, standard.FileGlobals.default_context))
	if intern.GetPathLen() > 1 && path[intern.GetPathLen()-1] == '/' {
		intern.SetPath(zend._estrndup(path, g.PreDec(&(intern.GetPathLen()))))
	} else {
		intern.SetPath(zend._estrndup(path, intern.GetPathLen()))
	}
	intern.SetIndex(0)
	if zend.EG.exception != nil || intern.GetDirp() == nil {
		intern.u.dir.entry.d_name[0] = '0'
		if zend.EG.exception == nil {

			/* open failed w/out notice (turned to exception due to EH_THROW) */

			zend.ZendThrowExceptionEx(spl_ce_UnexpectedValueException, 0, "Failed to open directory \"%s\"", path)

			/* open failed w/out notice (turned to exception due to EH_THROW) */

		}
	} else {
		for {
			SplFilesystemDirRead(intern)
			if !(skip_dots != 0 && SplFilesystemIsDot(intern.u.dir.entry.d_name) != 0) {
				break
			}
		}
	}
}

/* }}} */

func SplFilesystemFileOpen(intern *SplFilesystemObject, use_include_path int, silent int) int {
	var tmp zend.Zval
	intern.SetType(SPL_FS_FILE)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 13, &tmp)
	if tmp.u1.v.type_ == 3 {
		intern.SetOpenMode(nil)
		intern.SetFileName(nil)
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Cannot use SplFileObject with directories")
		return zend.FAILURE
	}
	if g.CondF2(g.CondF1(intern.GetZcontext() != nil, func() any {
		return zend.ZendFetchResourceEx(intern.GetZcontext(), "Stream-Context", standard.PhpLeStreamContext())
	}, 0), nil, func() *core.PhpStreamContext { return standard.FileGlobals.default_context }) {
		intern.SetContext(standard.FileGlobals.default_context)
	} else {
		standard.FileGlobals.default_context = streams.PhpStreamContextAlloc()
		intern.SetContext(standard.FileGlobals.default_context)
	}
	intern.SetStream(streams._phpStreamOpenWrapperEx(intern.GetFileName(), intern.GetOpenMode(), g.Cond(use_include_path != 0, 0x1, 0)|0x8, nil, intern.GetContext()))
	if intern.GetFileNameLen() == 0 || intern.GetStream() == nil {
		if zend.EG.exception == nil {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot open file '%s'", g.CondF1(intern.GetFileNameLen() != 0, func() *byte { return intern.GetFileName() }, ""))
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

	if intern.GetFileNameLen() > 1 && intern.GetFileName()[intern.GetFileNameLen()-1] == '/' {
		intern.GetFileNameLen()--
	}
	intern.SetOrigPath(zend._estrndup(intern.GetStream().orig_path, strlen(intern.GetStream().orig_path)))
	intern.SetFileName(zend._estrndup(intern.GetFileName(), intern.GetFileNameLen()))
	intern.SetOpenMode(zend._estrndup(intern.GetOpenMode(), intern.GetOpenModeLen()))

	/* avoid reference counting in debug mode, thus do it manually */

	var __z *zend.Zval = &intern.u.file.zresource
	__z.value.res = intern.GetStream().res
	__z.u1.type_info = 9 | 1<<0<<8

	/*!!! TODO: maybe bug?
	  Z_SET_REFCOUNT(intern->u.file.zresource, 1);
	*/

	intern.SetDelimiter(',')
	intern.SetEnclosure('"')
	intern.SetEscape(uint8('\\'))
	intern.SetFuncGetCurr(zend.ZendHashStrFindPtr(&intern.std.ce.function_table, "getcurrentline", g.SizeOf("\"getcurrentline\"")-1))
	return zend.SUCCESS
}

/* {{{ spl_filesystem_object_clone */

func SplFilesystemObjectClone(zobject *zend.Zval) *zend.ZendObject {
	var old_object *zend.ZendObject
	var new_object *zend.ZendObject
	var intern *SplFilesystemObject
	var source *SplFilesystemObject
	var index int
	var skip_dots int
	old_object = zobject.value.obj
	source = SplFilesystemFromObj(old_object)
	new_object = SplFilesystemObjectNewEx(old_object.ce)
	intern = SplFilesystemFromObj(new_object)
	intern.SetFlags(source.GetFlags())
	switch source.GetType() {
	case SPL_FS_INFO:
		intern.SetPathLen(source.GetPathLen())
		intern.SetPath(zend._estrndup(source.GetPath(), source.GetPathLen()))
		intern.SetFileNameLen(source.GetFileNameLen())
		intern.SetFileName(zend._estrndup(source.GetFileName(), intern.GetFileNameLen()))
		break
	case SPL_FS_DIR:
		SplFilesystemDirOpen(intern, source.GetPath())

		/* read until we hit the position in which we were before */

		if (source.GetFlags() & 0x1000) != 0 {
			skip_dots = 1
		} else {
			skip_dots = 0
		}
		for index = 0; index < source.GetIndex(); index++ {
			for {
				SplFilesystemDirRead(intern)
				if !(skip_dots != 0 && SplFilesystemIsDot(intern.u.dir.entry.d_name) != 0) {
					break
				}
			}
		}
		intern.SetIndex(index)
		break
	case SPL_FS_FILE:
		assert(false)
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

/* }}} */

func SplFilesystemInfoSetFilename(intern *SplFilesystemObject, path *byte, len_ int, use_copy int) {
	var p1 *byte
	var p2 *byte
	if intern.GetFileName() != nil {
		zend._efree(intern.GetFileName())
	}
	if use_copy != 0 {
		intern.SetFileName(zend._estrndup(path, len_))
	} else {
		intern.SetFileName(path)
	}
	intern.SetFileNameLen(len_)
	for intern.GetFileNameLen() > 1 && intern.GetFileName()[intern.GetFileNameLen()-1] == '/' {
		intern.GetFileName()[intern.GetFileNameLen()-1] = 0
		intern.GetFileNameLen()--
	}
	p1 = strrchr(intern.GetFileName(), '/')
	p2 = 0
	if p1 != nil || p2 != nil {
		intern.SetPathLen(g.Cond(p1 > p2, p1, p2) - intern.GetFileName())
	} else {
		intern.SetPathLen(0)
	}
	if intern.GetPath() != nil {
		zend._efree(intern.GetPath())
	}
	intern.SetPath(zend._estrndup(path, intern.GetPathLen()))
}
func SplFilesystemObjectCreateInfo(source *SplFilesystemObject, file_path *byte, file_path_len int, use_copy int, ce *zend.ZendClassEntry, return_value *zend.Zval) *SplFilesystemObject {
	var intern *SplFilesystemObject
	var arg1 zend.Zval
	var error_handling zend.ZendErrorHandling
	if file_path == nil || file_path_len == 0 {
		if file_path != nil && use_copy == 0 {
			zend._efree(file_path)
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
	var __z *zend.Zval = return_value
	__z.value.obj = &intern.std
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	if ce.constructor.common.scope != spl_ce_SplFileInfo {
		var __z *zend.Zval = &arg1
		var __s *zend.ZendString = zend.ZendStringInit(file_path, file_path_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendCallMethod(return_value, ce, &ce.constructor, "__construct", g.SizeOf("\"__construct\"")-1, nil, 1, &arg1, nil)
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
		if !(source.u.dir.entry.d_name[0]) {
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
		var __z *zend.Zval = return_value
		__z.value.obj = &intern.std
		__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
		SplFilesystemObjectGetFileName(source)
		if ce.constructor.common.scope != spl_ce_SplFileInfo {
			var __z *zend.Zval = &arg1
			var __s *zend.ZendString = zend.ZendStringInit(source.GetFileName(), source.GetFileNameLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendCallMethod(return_value, ce, &ce.constructor, "__construct", g.SizeOf("\"__construct\"")-1, nil, 1, &arg1, nil)
			zend.ZvalPtrDtor(&arg1)
		} else {
			intern.SetFileName(zend._estrndup(source.GetFileName(), source.GetFileNameLen()))
			intern.SetFileNameLen(source.GetFileNameLen())
			intern.SetPath(SplFilesystemObjectGetPath(source, &intern._path_len))
			intern.SetPath(zend._estrndup(intern.GetPath(), intern.GetPathLen()))
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
		var __z *zend.Zval = return_value
		__z.value.obj = &intern.std
		__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
		SplFilesystemObjectGetFileName(source)
		if ce.constructor.common.scope != spl_ce_SplFileObject {
			var __z *zend.Zval = &arg1
			var __s *zend.ZendString = zend.ZendStringInit(source.GetFileName(), source.GetFileNameLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			var __z *zend.Zval = &arg2
			var __s *zend.ZendString = zend.ZendStringInit("r", 1, 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			zend.ZendCallMethod(return_value, ce, &ce.constructor, "__construct", g.SizeOf("\"__construct\"")-1, nil, 2, &arg1, &arg2)
			zend.ZvalPtrDtor(&arg1)
			zend.ZvalPtrDtor(&arg2)
		} else {
			intern.SetFileName(source.GetFileName())
			intern.SetFileNameLen(source.GetFileNameLen())
			intern.SetPath(SplFilesystemObjectGetPath(source, &intern._path_len))
			intern.SetPath(zend._estrndup(intern.GetPath(), intern.GetPathLen()))
			intern.SetOpenMode("r")
			intern.SetOpenModeLen(1)
			if ht != 0 && zend.ZendParseParameters(ht, "|sbr", &intern.u.file.open_mode, &intern.u.file.open_mode_len, &use_include_path, &intern.u.file.zcontext) == zend.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				intern.SetOpenMode(nil)
				intern.SetFileName(nil)
				zend.ZvalPtrDtor(return_value)
				return_value.u1.type_info = 1
				return nil
			}
			if SplFilesystemFileOpen(intern, use_include_path, 0) == zend.FAILURE {
				zend.ZendRestoreErrorHandling(&error_handling)
				zend.ZvalPtrDtor(return_value)
				return_value.u1.type_info = 1
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

/* }}} */

func SplFilesystemObjectGetPathname(intern *SplFilesystemObject, len_ *int) *byte {
	switch intern.GetType() {
	case SPL_FS_INFO:

	case SPL_FS_FILE:
		*len_ = intern.GetFileNameLen()
		return intern.GetFileName()
	case SPL_FS_DIR:
		if intern.u.dir.entry.d_name[0] {
			SplFilesystemObjectGetFileName(intern)
			*len_ = intern.GetFileNameLen()
			return intern.GetFileName()
		}
	}
	*len_ = 0
	return nil
}

/* }}} */

func SplFilesystemObjectGetDebugInfo(object *zend.Zval) *zend.HashTable {
	var intern *SplFilesystemObject = SplFilesystemFromObj(object.value.obj)
	var tmp zend.Zval
	var rv *zend.HashTable
	var pnstr *zend.ZendString
	var path *byte
	var path_len int
	var stmp []byte
	if intern.std.properties == nil {
		zend.RebuildObjectProperties(&intern.std)
	}
	rv = zend.ZendArrayDup(intern.std.properties)
	pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "pathName", g.SizeOf("\"pathName\"")-1)
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	var __z *zend.Zval = &tmp
	var __s *zend.ZendString = zend.ZendStringInit(g.Cond(path != nil, path, ""), path_len, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend.ZendSymtableUpdate(rv, pnstr, &tmp)
	zend.ZendStringReleaseEx(pnstr, 0)
	if intern.GetFileName() != nil {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileInfo, "fileName", g.SizeOf("\"fileName\"")-1)
		SplFilesystemObjectGetPath(intern, &path_len)
		if path_len != 0 && path_len < intern.GetFileNameLen() {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		}
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_DIR {
		pnstr = SplGenPrivatePropName(spl_ce_DirectoryIterator, "glob", g.SizeOf("\"glob\"")-1)
		if intern.GetDirp().ops == &streams.PhpGlobStreamOps {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(intern.GetPath(), intern.GetPathLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			&tmp.u1.type_info = 2
		}
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		pnstr = SplGenPrivatePropName(spl_ce_RecursiveDirectoryIterator, "subPathName", g.SizeOf("\"subPathName\"")-1)
		if intern.GetSubPath() != nil {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendStringInit(intern.GetSubPath(), intern.GetSubPathLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		} else {
			var __z *zend.Zval = &tmp
			var __s *zend.ZendString = zend.ZendEmptyString
			__z.value.str = __s
			__z.u1.type_info = 6
		}
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	if intern.GetType() == SPL_FS_FILE {
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "openMode", g.SizeOf("\"openMode\"")-1)
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetOpenMode(), intern.GetOpenModeLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		stmp[1] = '0'
		stmp[0] = intern.GetDelimiter()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "delimiter", g.SizeOf("\"delimiter\"")-1)
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(stmp, 1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
		stmp[0] = intern.GetEnclosure()
		pnstr = SplGenPrivatePropName(spl_ce_SplFileObject, "enclosure", g.SizeOf("\"enclosure\"")-1)
		var __z *zend.Zval = &tmp
		var __s *zend.ZendString = zend.ZendStringInit(stmp, 1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendSymtableUpdate(rv, pnstr, &tmp)
		zend.ZendStringReleaseEx(pnstr, 0)
	}
	return rv
}

/* }}} */

func SplFilesystemObjectGetMethodCheck(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var fsobj *SplFilesystemObject = SplFilesystemFromObj(*object)
	if fsobj.GetDirp() == nil && fsobj.GetOrigPath() == nil {
		var func_ *zend.ZendFunction
		var tmp *zend.ZendString = zend.ZendStringInit("_bad_state_ex", g.SizeOf("\"_bad_state_ex\"")-1, 0)
		func_ = zend.ZendStdGetMethod(object, tmp, nil)
		zend.ZendStringReleaseEx(tmp, 0)
		return func_
	}
	return zend.ZendStdGetMethod(object, method, key)
}

/* }}} */

// #define DIT_CTOR_FLAGS       0x00000001

// #define DIT_CTOR_GLOB       0x00000002

func SplFilesystemObjectConstruct(execute_data *zend.ZendExecuteData, return_value *zend.Zval, ctor_flags zend.ZendLong) {
	var intern *SplFilesystemObject
	var path *byte
	var parsed int
	var len_ int
	var flags zend.ZendLong
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if g.Cond((ctor_flags&0x1) != 0, 1, 0) {
		flags = 0x0 | 0x0
		parsed = zend.ZendParseParameters(execute_data.This.u2.num_args, "p|l", &path, &len_, &flags)
	} else {
		flags = 0x0 | 0x10
		parsed = zend.ZendParseParameters(execute_data.This.u2.num_args, "p", &path, &len_)
	}
	if g.Cond((ctor_flags&0x1000) != 0, 1, 0) {
		flags |= 0x1000
	}
	if g.Cond((ctor_flags&0x2000) != 0, 1, 0) {
		flags |= 0x2000
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
	intern = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if intern.GetPath() != nil {

		/* object is already initialized */

		zend.ZendRestoreErrorHandling(&error_handling)
		core.PhpErrorDocref(nil, 1<<1, "Directory object is already initialized")
		return
	}
	intern.SetFlags(flags)
	if g.Cond((ctor_flags&0x2) != 0, 1, 0) && strstr(path, "glob://") != path {
		zend.ZendSpprintf(&path, 0, "glob://%s", path)
		SplFilesystemDirOpen(intern, path)
		zend._efree(path)
	} else {
		SplFilesystemDirOpen(intern, path)
	}
	if zend.InstanceofFunction(intern.std.ce, spl_ce_RecursiveDirectoryIterator) != 0 {
		intern.SetIsRecursive(1)
	} else {
		intern.SetIsRecursive(0)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_DirectoryIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, 0)
}

/* }}} */

func zim_spl_DirectoryIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern.SetIndex(0)
	if intern.GetDirp() != nil {
		streams._phpStreamSeek(intern.GetDirp(), 0, SEEK_SET)
	}
	SplFilesystemDirRead(intern)
}

/* }}} */

func zim_spl_DirectoryIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetDirp() != nil {
		var __z *zend.Zval = return_value
		__z.value.lval = intern.GetIndex()
		__z.u1.type_info = 4
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func zim_spl_DirectoryIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.obj = &(execute_data.This).value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	zend.ZvalAddrefP(return_value)
}

/* }}} */

func zim_spl_DirectoryIterator_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var skip_dots int = g.Cond((intern.GetFlags()&0x1000) != 0, 1, 0)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern.GetIndex()++
	for {
		SplFilesystemDirRead(intern)
		if !(skip_dots != 0 && SplFilesystemIsDot(intern.u.dir.entry.d_name) != 0) {
			break
		}
	}
	if intern.GetFileName() != nil {
		zend._efree(intern.GetFileName())
		intern.SetFileName(nil)
	}
}

/* }}} */

func zim_spl_DirectoryIterator_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var retval zend.Zval
	var pos zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &pos) == zend.FAILURE {
		return
	}
	if intern.GetIndex() > pos {

		/* we first rewind */

		zend.ZendCallMethod(&(execute_data.This), &(execute_data.This).value.obj.ce, &intern.u.dir.func_rewind, "rewind", g.SizeOf("\"rewind\"")-1, nil, 0, nil, nil)

		/* we first rewind */

	}
	for intern.GetIndex() < pos {
		var valid int = 0
		zend.ZendCallMethod(&(execute_data.This), &(execute_data.This).value.obj.ce, &intern.u.dir.func_valid, "valid", g.SizeOf("\"valid\"")-1, &retval, 0, nil, nil)
		valid = zend.ZendIsTrue(&retval)
		zend.ZvalPtrDtor(&retval)
		if valid == 0 {
			zend.ZendThrowExceptionEx(spl_ce_OutOfBoundsException, 0, "Seek position "+"%"+"lld"+" is out of range", pos)
			return
		}
		zend.ZendCallMethod(&(execute_data.This), &(execute_data.This).value.obj.ce, &intern.u.dir.func_next, "next", g.SizeOf("\"next\"")-1, nil, 0, nil, nil)
	}
}

/* {{{ proto string DirectoryIterator::valid()
   Check whether dir contains more entries */

func zim_spl_DirectoryIterator_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.u.dir.entry.d_name[0] != '0' {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplFileInfo_getPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var path *byte
	var path_len int
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPath(intern, &path_len)
	if path != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(path, path_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* }}} */

func zim_spl_SplFileInfo_getFilename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var path_len int
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplFilesystemObjectGetPath(intern, &path_len)
	if path_len != 0 && path_len < intern.GetFileNameLen() {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName()+path_len+1, intern.GetFileNameLen()-(path_len+1), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func zim_spl_DirectoryIterator_getFilename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var _s *byte = intern.u.dir.entry.d_name
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* }}} */

func zim_spl_SplFileInfo_getExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var fname *byte = nil
	var p *byte
	var flen int
	var path_len int
	var idx int
	var ret *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
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
	p = zend.ZendMemrchr(ret.val, '.', ret.len_)
	if p != nil {
		idx = p - ret.val
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(ret.val+idx+1, ret.len_-idx-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendStringReleaseEx(ret, 0)
		return
	} else {
		zend.ZendStringReleaseEx(ret, 0)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* }}}*/

func zim_spl_DirectoryIterator_getExtension(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var p *byte
	var idx int
	var fname *zend.ZendString
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.u.dir.entry.d_name, strlen(intern.u.dir.entry.d_name), nil, 0)
	p = zend.ZendMemrchr(fname.val, '.', fname.len_)
	if p != nil {
		idx = p - fname.val
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(fname.val+idx+1, fname.len_-idx-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		zend.ZendStringReleaseEx(fname, 0)
	} else {
		zend.ZendStringReleaseEx(fname, 0)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* }}} */

func zim_spl_SplFileInfo_getBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var fname *byte
	var suffix *byte = 0
	var flen int
	var slen int = 0
	var path_len int
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|s", &suffix, &slen) == zend.FAILURE {
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
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = standard.PhpBasename(fname, flen, suffix, slen)
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* }}}*/

func zim_spl_DirectoryIterator_getBasename(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var suffix *byte = 0
	var slen int = 0
	var fname *zend.ZendString
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|s", &suffix, &slen) == zend.FAILURE {
		return
	}
	fname = standard.PhpBasename(intern.u.dir.entry.d_name, strlen(intern.u.dir.entry.d_name), suffix, slen)
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = fname
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func zim_spl_SplFileInfo_getPathname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var path *byte
	var path_len int
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	path = SplFilesystemObjectGetPathname(intern, &path_len)
	if path != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(path, path_len, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		return_value.u1.type_info = 2
		return
	}
}

/* }}} */

func zim_spl_FilesystemIterator_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if (intern.GetFlags() & 0xf00) == 0x100 {
		var _s *byte = intern.u.dir.entry.d_name
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func zim_spl_FilesystemIterator_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if (intern.GetFlags() & 0xf0) == 0x20 {
		SplFilesystemObjectGetFileName(intern)
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else if (intern.GetFlags() & 0xf0) == 0x0 {
		SplFilesystemObjectGetFileName(intern)
		SplFilesystemObjectCreateType(0, intern, SPL_FS_INFO, nil, return_value)
	} else {
		var __z *zend.Zval = return_value
		__z.value.obj = &(execute_data.This).value.obj
		__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
		zend.ZvalAddrefP(return_value)
	}
}

/* }}} */

func zim_spl_DirectoryIterator_isDot(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if SplFilesystemIsDot(intern.u.dir.entry.d_name) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* }}} */

func zim_spl_SplFileInfo___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject
	var path *byte
	var len_ int
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "p", &path, &len_) == zend.FAILURE {
		return
	}
	intern = SplFilesystemFromObj(&(execute_data.This).value.obj)
	SplFilesystemInfoSetFilename(intern, path, len_, 1)
}

/* }}} */

// #define FileInfoFunction(func_name,func_num) SPL_METHOD ( SplFileInfo , func_name ) { spl_filesystem_object * intern = Z_SPLFILESYSTEM_P ( ZEND_THIS ) ; zend_error_handling error_handling ; if ( zend_parse_parameters_none ( ) == FAILURE ) { return ; } zend_replace_error_handling ( EH_THROW , spl_ce_RuntimeException , & error_handling ) ; spl_filesystem_object_get_file_name ( intern ) ; php_stat ( intern -> file_name , intern -> file_name_len , func_num , return_value ) ; zend_restore_error_handling ( & error_handling ) ; }

/* }}} */

func zim_spl_SplFileInfo_getPerms(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 0, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getInode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 1, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getSize(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 2, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getOwner(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 3, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getGroup(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 4, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getATime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 5, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getMTime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 6, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getCTime(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 7, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getType(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 8, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isWritable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 9, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isReadable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 10, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isExecutable(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 11, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 12, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isDir(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 13, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_isLink(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	SplFilesystemObjectGetFileName(intern)
	standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 14, return_value)
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getLinkTarget(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ret ssize_t
	var buff []byte
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if intern.GetFileName() == nil {
		SplFilesystemObjectGetFileName(intern)
	}
	if intern.GetFileName() == nil {
		core.PhpErrorDocref(nil, 1<<1, "Empty filename")
		return_value.u1.type_info = 2
		return
	} else if intern.GetFileName()[0] != '/' {
		var expanded_path []byte
		if core.ExpandFilepathWithMode(intern.GetFileName(), expanded_path, nil, 0, 0) == nil {
			core.PhpErrorDocref(nil, 1<<1, "No such file or directory")
			return_value.u1.type_info = 2
			return
		}
		ret = readlink(expanded_path, buff, 256-1)
	} else {
		ret = readlink(intern.GetFileName(), buff, 256-1)
	}
	if ret == -1 {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Unable to read link %s, error: %s", intern.GetFileName(), strerror(errno))
		return_value.u1.type_info = 2
	} else {

		/* Append NULL to the end of the string */

		buff[ret] = '0'
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(buff, ret, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

/* {{{ proto string SplFileInfo::getRealPath()
   Return the resolved path */

func zim_spl_SplFileInfo_getRealPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var buff []byte
	var filename *byte
	var error_handling zend.ZendErrorHandling
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_RuntimeException, &error_handling)
	if intern.GetType() == SPL_FS_DIR && intern.GetFileName() == nil && intern.u.dir.entry.d_name[0] {
		SplFilesystemObjectGetFileName(intern)
	}
	if intern.GetOrigPath() != nil {
		filename = intern.GetOrigPath()
	} else {
		filename = intern.GetFileName()
	}
	if filename != nil && zend.TsrmRealpath(filename, buff) != nil {
		var _s *byte = buff
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		return_value.u1.type_info = 2
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

/* {{{ proto SplFileObject SplFileInfo::openFile([string mode = 'r' [, bool use_include_path  [, resource context]]])
   Open the current file */

func zim_spl_SplFileInfo_openFile(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	SplFilesystemObjectCreateType(execute_data.This.u2.num_args, intern, SPL_FS_FILE, nil, return_value)
}

/* }}} */

func zim_spl_SplFileInfo_setFileClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry = spl_ce_SplFileObject
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|C", &ce) == zend.SUCCESS {
		intern.SetFileClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_setInfoClass(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry = spl_ce_SplFileInfo
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|C", &ce) == zend.SUCCESS {
		intern.SetInfoClass(ce)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getFileInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|C", &ce) == zend.SUCCESS {
		SplFilesystemObjectCreateType(execute_data.This.u2.num_args, intern, SPL_FS_INFO, ce, return_value)
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo_getPathInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ce *zend.ZendClassEntry = intern.GetInfoClass()
	var error_handling zend.ZendErrorHandling
	zend.ZendReplaceErrorHandling(zend.EH_THROW, spl_ce_UnexpectedValueException, &error_handling)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|C", &ce) == zend.SUCCESS {
		var path_len int
		var path *byte = SplFilesystemObjectGetPathname(intern, &path_len)
		if path != nil {
			var dpath *byte = zend._estrndup(path, path_len)
			path_len = standard.PhpDirname(dpath, path_len)
			SplFilesystemObjectCreateInfo(intern, dpath, path_len, 1, ce, return_value)
			zend._efree(dpath)
		}
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* }}} */

func zim_spl_SplFileInfo___debugInfo(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __arr *zend.ZendArray = SplFilesystemObjectGetDebugInfo(g.CondF1(&(execute_data.This).u1.v.type_ == 8, func() *zend.Zval { return &(execute_data.This) }, nil))
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	return
}

/* {{{  proto SplFileInfo::_bad_state_ex(void) */

func zim_spl_SplFileInfo__bad_state_ex(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "The parent constructor was not called: the object is in an "+"invalid state ")
}

/* }}} */

func zim_spl_FilesystemIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, 0x1|0x1000)
}

/* }}} */

func zim_spl_FilesystemIterator_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var skip_dots int = g.Cond((intern.GetFlags()&0x1000) != 0, 1, 0)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	intern.SetIndex(0)
	if intern.GetDirp() != nil {
		streams._phpStreamSeek(intern.GetDirp(), 0, SEEK_SET)
	}
	for {
		SplFilesystemDirRead(intern)
		if !(skip_dots != 0 && SplFilesystemIsDot(intern.u.dir.entry.d_name) != 0) {
			break
		}
	}
}

/* }}} */

func zim_spl_FilesystemIterator_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags() & (0xf00 | 0xf0 | 0x3000)
	__z.u1.type_info = 4
	return
}

/* {{{ proto void FilesystemIterator::setFlags(long $flags)
   Set handling flags */

func zim_spl_FilesystemIterator_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var flags zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &flags) == zend.FAILURE {
		return
	}
	intern.SetFlags(intern.GetFlags() &^ (0xf00 | 0xf0 | 0x3000))
	intern.SetFlags(intern.GetFlags() | (0xf00|0xf0|0x3000)&flags)
}

/* {{{ proto bool RecursiveDirectoryIterator::hasChildren([bool $allow_links = false])
   Returns whether current entry is a directory and not '.' or '..' */

func zim_spl_RecursiveDirectoryIterator_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var allow_links zend.ZendBool = 0
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|b", &allow_links) == zend.FAILURE {
		return
	}
	if SplFilesystemIsInvalidOrDot(intern.u.dir.entry.d_name) != 0 {
		return_value.u1.type_info = 2
		return
	} else {
		SplFilesystemObjectGetFileName(intern)
		if allow_links == 0 && (intern.GetFlags()&0x200) == 0 {
			standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 14, return_value)
			if zend.ZendIsTrue(return_value) != 0 {
				return_value.u1.type_info = 2
				return
			}
		}
		standard.PhpStat(intern.GetFileName(), intern.GetFileNameLen(), 13, return_value)
	}
}

/* }}} */

func zim_spl_RecursiveDirectoryIterator_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var zpath zend.Zval
	var zflags zend.Zval
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var subdir *SplFilesystemObject
	var slash byte = g.Cond(g.Cond((intern.GetFlags()&0x2000) != 0, 1, 0), '/', '/')
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplFilesystemObjectGetFileName(intern)
	var __z *zend.Zval = &zflags
	__z.value.lval = intern.GetFlags()
	__z.u1.type_info = 4
	var __z *zend.Zval = &zpath
	var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	SplInstantiateArgEx2(&(execute_data.This).value.obj.ce, return_value, &zpath, &zflags)
	zend.ZvalPtrDtor(&zpath)
	subdir = SplFilesystemFromObj(return_value.value.obj)
	if subdir != nil {
		if intern.GetSubPath() != nil && intern.GetSubPath()[0] {
			subdir.SetSubPathLen(zend.ZendSpprintf(&subdir.u.dir.sub_path, 0, "%s%c%s", intern.GetSubPath(), slash, intern.u.dir.entry.d_name))
		} else {
			subdir.SetSubPathLen(strlen(intern.u.dir.entry.d_name))
			subdir.SetSubPath(zend._estrndup(intern.u.dir.entry.d_name, subdir.GetSubPathLen()))
		}
		subdir.SetInfoClass(intern.GetInfoClass())
		subdir.SetFileClass(intern.GetFileClass())
		subdir.SetOth(intern.GetOth())
	}
}

/* }}} */

func zim_spl_RecursiveDirectoryIterator_getSubPath(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetSubPath(), intern.GetSubPathLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendEmptyString
		__z.value.str = __s
		__z.u1.type_info = 6
		return
	}
}

/* }}} */

func zim_spl_RecursiveDirectoryIterator_getSubPathname(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var slash byte = g.Cond(g.Cond((intern.GetFlags()&0x2000) != 0, 1, 0), '/', '/')
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetSubPath() != nil {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStrpprintf(0, "%s%c%s", intern.GetSubPath(), slash, intern.u.dir.entry.d_name)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else {
		var _s *byte = intern.u.dir.entry.d_name
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* }}} */

func zim_spl_RecursiveDirectoryIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, 0x1)
}

/* }}} */

/* {{{ proto GlobIterator::__construct(string path [, int flags])
Cronstructs a new dir iterator from a glob expression (no glob:// needed). */

func zim_spl_GlobIterator___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	SplFilesystemObjectConstruct(execute_data, return_value, 0x1|0x2)
}

/* }}} */

func zim_spl_GlobIterator_count(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetDirp() != nil && intern.GetDirp().ops == &streams.PhpGlobStreamOps {
		var __z *zend.Zval = return_value
		__z.value.lval = streams._phpGlobStreamGetCount(intern.GetDirp(), nil)
		__z.u1.type_info = 4
		return
	} else {

		/* should not happen */

		core.PhpErrorDocref(nil, 1<<0, "GlobIterator lost glob state")

		/* should not happen */

	}
}

/* }}} */

/* {{{ forward declarations to the iterator handlers */

/* iterator handler table */

var SplFilesystemDirItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplFilesystemDirItDtor, SplFilesystemDirItValid, SplFilesystemDirItCurrentData, SplFilesystemDirItCurrentKey, SplFilesystemDirItMoveForward, SplFilesystemDirItRewind, nil}

/* }}} */

func SplFilesystemDirGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = SplFilesystemFromObj(object.value.obj)
	iterator = SplFilesystemObjectToIterator(dir_object)
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.funcs = &SplFilesystemDirItFuncs

	/* ->current must be initialized; rewind doesn't set it and valid
	 * doesn't check whether it's set */

	iterator.SetCurrent(*object)
	return &iterator.intern
}

/* }}} */

func SplFilesystemDirItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	zend.ZvalPtrDtor(&iterator.intern.data)
}

/* }}} */

func SplFilesystemDirItValid(iter *zend.ZendObjectIterator) int {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	if object.u.dir.entry.d_name[0] != '0' {
		return zend.SUCCESS
	} else {
		return zend.FAILURE
	}
}

/* }}} */

func SplFilesystemDirItCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	return &iterator.current
}

/* }}} */

func SplFilesystemDirItCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	var __z *zend.Zval = key
	__z.value.lval = object.GetIndex()
	__z.u1.type_info = 4
}

/* }}} */

func SplFilesystemDirItMoveForward(iter *zend.ZendObjectIterator) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	object.GetIndex()++
	SplFilesystemDirRead(object)
	if object.GetFileName() != nil {
		zend._efree(object.GetFileName())
		object.SetFileName(nil)
	}
}

/* }}} */

func SplFilesystemDirItRewind(iter *zend.ZendObjectIterator) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	object.SetIndex(0)
	if object.GetDirp() != nil {
		streams._phpStreamSeek(object.GetDirp(), 0, SEEK_SET)
	}
	SplFilesystemDirRead(object)
}

/* }}} */

func SplFilesystemTreeItDtor(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	zend.ZvalPtrDtor(&iterator.intern.data)
	zend.ZvalPtrDtor(&iterator.current)
}

/* }}} */

func SplFilesystemTreeItCurrentData(iter *zend.ZendObjectIterator) *zend.Zval {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	if (object.GetFlags() & 0xf0) == 0x20 {
		if iterator.current.u1.v.type_ == 0 {
			SplFilesystemObjectGetFileName(object)
			var __z *zend.Zval = &iterator.current
			var __s *zend.ZendString = zend.ZendStringInit(object.GetFileName(), object.GetFileNameLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
		}
		return &iterator.current
	} else if (object.GetFlags() & 0xf0) == 0x0 {
		if iterator.current.u1.v.type_ == 0 {
			SplFilesystemObjectGetFileName(object)
			SplFilesystemObjectCreateType(0, object, SPL_FS_INFO, nil, &iterator.current)
		}
		return &iterator.current
	} else {
		return &iterator.intern.data
	}
}

/* }}} */

func SplFilesystemTreeItCurrentKey(iter *zend.ZendObjectIterator, key *zend.Zval) {
	var object *SplFilesystemObject = SplFilesystemIteratorToObject((*SplFilesystemIterator)(iter))
	if (object.GetFlags() & 0xf00) == 0x100 {
		var _s *byte = object.u.dir.entry.d_name
		var __z *zend.Zval = key
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	} else {
		SplFilesystemObjectGetFileName(object)
		var __z *zend.Zval = key
		var __s *zend.ZendString = zend.ZendStringInit(object.GetFileName(), object.GetFileNameLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
	}
}

/* }}} */

func SplFilesystemTreeItMoveForward(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	object.GetIndex()++
	for {
		SplFilesystemDirRead(object)
		if SplFilesystemIsDot(object.u.dir.entry.d_name) == 0 {
			break
		}
	}
	if object.GetFileName() != nil {
		zend._efree(object.GetFileName())
		object.SetFileName(nil)
	}
	if iterator.current.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&iterator.current)
		&iterator.current.u1.type_info = 0
	}
}

/* }}} */

func SplFilesystemTreeItRewind(iter *zend.ZendObjectIterator) {
	var iterator *SplFilesystemIterator = (*SplFilesystemIterator)(iter)
	var object *SplFilesystemObject = SplFilesystemIteratorToObject(iterator)
	object.SetIndex(0)
	if object.GetDirp() != nil {
		streams._phpStreamSeek(object.GetDirp(), 0, SEEK_SET)
	}
	for {
		SplFilesystemDirRead(object)
		if SplFilesystemIsDot(object.u.dir.entry.d_name) == 0 {
			break
		}
	}
	if iterator.current.u1.v.type_ != 0 {
		zend.ZvalPtrDtor(&iterator.current)
		&iterator.current.u1.type_info = 0
	}
}

/* }}} */

var SplFilesystemTreeItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplFilesystemTreeItDtor, SplFilesystemDirItValid, SplFilesystemTreeItCurrentData, SplFilesystemTreeItCurrentKey, SplFilesystemTreeItMoveForward, SplFilesystemTreeItRewind, nil}

/* }}} */

func SplFilesystemTreeGetIterator(ce *zend.ZendClassEntry, object *zend.Zval, by_ref int) *zend.ZendObjectIterator {
	var iterator *SplFilesystemIterator
	var dir_object *SplFilesystemObject
	if by_ref != 0 {
		zend.ZendThrowException(spl_ce_RuntimeException, "An iterator cannot be used with foreach by reference", 0)
		return nil
	}
	dir_object = SplFilesystemFromObj(object.value.obj)
	iterator = SplFilesystemObjectToIterator(dir_object)
	zend.ZvalAddrefP(object)
	var __z *zend.Zval = &iterator.intern.data
	__z.value.obj = object.value.obj
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	iterator.intern.funcs = &SplFilesystemTreeItFuncs
	return &iterator.intern
}

/* }}} */

func SplFilesystemObjectCast(readobj *zend.Zval, writeobj *zend.Zval, type_ int) int {
	var intern *SplFilesystemObject = SplFilesystemFromObj(readobj.value.obj)
	if type_ == 6 {
		if readobj.value.obj.ce.__tostring != nil {
			return zend.ZendStdCastObjectTostring(readobj, writeobj, type_)
		}
		switch intern.GetType() {
		case SPL_FS_INFO:

		case SPL_FS_FILE:
			var __z *zend.Zval = writeobj
			var __s *zend.ZendString = zend.ZendStringInit(intern.GetFileName(), intern.GetFileNameLen(), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return zend.SUCCESS
		case SPL_FS_DIR:
			var _s *byte = intern.u.dir.entry.d_name
			var __z *zend.Zval = writeobj
			var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
			__z.value.str = __s
			__z.u1.type_info = 6 | 1<<0<<8
			return zend.SUCCESS
		}
	} else if type_ == 16 {
		writeobj.u1.type_info = 3
		return zend.SUCCESS
	}
	writeobj.u1.type_info = 1
	return zend.FAILURE
}

/* }}} */

var ArginfoInfoConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"file_name", 0, 0, 0}}
var arginfo_info_openFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"open_mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var arginfo_info_optinalFileClass []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"class_name", 0, 0, 0}}
var arginfo_optinalSuffix []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"suffix", 0, 0, 0}}
var ArginfoSplfileinfoVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}}

/* the method table */

var spl_SplFileInfo_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplFileInfo___construct,
		ArginfoInfoConstruct,
		uint32(g.SizeOf("arginfo_info___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPath",
		zim_spl_SplFileInfo_getPath,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFilename",
		zim_spl_SplFileInfo_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getExtension",
		zim_spl_SplFileInfo_getExtension,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getBasename",
		zim_spl_SplFileInfo_getBasename,
		arginfo_optinalSuffix,
		uint32(g.SizeOf("arginfo_optinalSuffix")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPathname",
		zim_spl_SplFileInfo_getPathname,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPerms",
		zim_spl_SplFileInfo_getPerms,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getInode",
		zim_spl_SplFileInfo_getInode,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getSize",
		zim_spl_SplFileInfo_getSize,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getOwner",
		zim_spl_SplFileInfo_getOwner,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getGroup",
		zim_spl_SplFileInfo_getGroup,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getATime",
		zim_spl_SplFileInfo_getATime,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getMTime",
		zim_spl_SplFileInfo_getMTime,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getCTime",
		zim_spl_SplFileInfo_getCTime,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getType",
		zim_spl_SplFileInfo_getType,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isWritable",
		zim_spl_SplFileInfo_isWritable,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isReadable",
		zim_spl_SplFileInfo_isReadable,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isExecutable",
		zim_spl_SplFileInfo_isExecutable,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isFile",
		zim_spl_SplFileInfo_isFile,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isDir",
		zim_spl_SplFileInfo_isDir,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isLink",
		zim_spl_SplFileInfo_isLink,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getLinkTarget",
		zim_spl_SplFileInfo_getLinkTarget,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getRealPath",
		zim_spl_SplFileInfo_getRealPath,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFileInfo",
		zim_spl_SplFileInfo_getFileInfo,
		arginfo_info_optinalFileClass,
		uint32(g.SizeOf("arginfo_info_optinalFileClass")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getPathInfo",
		zim_spl_SplFileInfo_getPathInfo,
		arginfo_info_optinalFileClass,
		uint32(g.SizeOf("arginfo_info_optinalFileClass")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"openFile",
		zim_spl_SplFileInfo_openFile,
		arginfo_info_openFile,
		uint32(g.SizeOf("arginfo_info_openFile")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFileClass",
		zim_spl_SplFileInfo_setFileClass,
		arginfo_info_optinalFileClass,
		uint32(g.SizeOf("arginfo_info_optinalFileClass")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setInfoClass",
		zim_spl_SplFileInfo_setInfoClass,
		arginfo_info_optinalFileClass,
		uint32(g.SizeOf("arginfo_info_optinalFileClass")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__debugInfo",
		zim_spl_SplFileInfo___debugInfo,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"_bad_state_ex",
		zim_spl_SplFileInfo__bad_state_ex,
		nil,
		uint32(g.SizeOf("NULL")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1<<0 | 1<<5,
	},
	{
		"__toString",
		zim_spl_SplFileInfo_getPathname,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoDirConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"path", 0, 0, 0}}
var ArginfoDirItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"position", 0, 0, 0}}

/* the method table */

var spl_DirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_DirectoryIterator___construct,
		ArginfoDirConstruct,
		uint32(g.SizeOf("arginfo_dir___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFilename",
		zim_spl_DirectoryIterator_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getExtension",
		zim_spl_DirectoryIterator_getExtension,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getBasename",
		zim_spl_DirectoryIterator_getBasename,
		arginfo_optinalSuffix,
		uint32(g.SizeOf("arginfo_optinalSuffix")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"isDot",
		zim_spl_DirectoryIterator_isDot,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_DirectoryIterator_rewind,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_DirectoryIterator_valid,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_DirectoryIterator_key,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_DirectoryIterator_current,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_DirectoryIterator_next,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"seek",
		zim_spl_DirectoryIterator_seek,
		ArginfoDirItSeek,
		uint32(g.SizeOf("arginfo_dir_it_seek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__toString",
		zim_spl_DirectoryIterator_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoRDirConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"path", 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_r_dir_hasChildren []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"allow_links", 0, 0, 0}}
var arginfo_r_dir_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"flags", 0, 0, 0}}
var spl_FilesystemIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_FilesystemIterator___construct,
		ArginfoRDirConstruct,
		uint32(g.SizeOf("arginfo_r_dir___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_FilesystemIterator_rewind,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_DirectoryIterator_next,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_FilesystemIterator_key,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_FilesystemIterator_current,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_FilesystemIterator_getFlags,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_FilesystemIterator_setFlags,
		arginfo_r_dir_setFlags,
		uint32(g.SizeOf("arginfo_r_dir_setFlags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_RecursiveDirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveDirectoryIterator___construct,
		ArginfoRDirConstruct,
		uint32(g.SizeOf("arginfo_r_dir___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_RecursiveDirectoryIterator_hasChildren,
		arginfo_r_dir_hasChildren,
		uint32(g.SizeOf("arginfo_r_dir_hasChildren")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_RecursiveDirectoryIterator_getChildren,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getSubPath",
		zim_spl_RecursiveDirectoryIterator_getSubPath,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getSubPathname",
		zim_spl_RecursiveDirectoryIterator_getSubPathname,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var spl_GlobIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_GlobIterator___construct,
		ArginfoRDirConstruct,
		uint32(g.SizeOf("arginfo_r_dir___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"count",
		zim_spl_GlobIterator_count,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func SplFilesystemFileRead(intern *SplFilesystemObject, silent int) int {
	var buf *byte
	var line_len int = 0
	var line_add zend.ZendLong = g.Cond(intern.GetCurrentLine() != nil || intern.u.file.current_zval.u1.v.type_ != 0, 1, 0)
	SplFilesystemFileFreeLine(intern)
	if streams._phpStreamEof(intern.GetStream()) != 0 {
		if silent == 0 {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
		}
		return zend.FAILURE
	}
	if intern.GetMaxLineLen() > 0 {
		buf = zend._safeEmalloc(intern.GetMaxLineLen()+1, g.SizeOf("char"), 0)
		if streams._phpStreamGetLine(intern.GetStream(), buf, intern.GetMaxLineLen()+1, &line_len) == nil {
			zend._efree(buf)
			buf = nil
		} else {
			buf[line_len] = '0'
		}
	} else {
		buf = streams._phpStreamGetLine(intern.GetStream(), nil, 0, &line_len)
	}
	if buf == nil {
		intern.SetCurrentLine(zend._estrdup(""))
		intern.SetCurrentLineLen(0)
	} else {
		if g.Cond((intern.GetFlags()&0x1) != 0, 1, 0) {
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
	var zresource_ptr *zend.Zval = &intern.u.file.zresource
	var params *zend.Zval
	var retval zend.Zval
	var result int
	var num_args int = pass_num_args + g.Cond(arg2 != nil, 2, 1)
	if zresource_ptr.u1.v.type_ == 0 {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return zend.FAILURE
	}
	params = (*zend.Zval)(zend._safeEmalloc(num_args, g.SizeOf("zval"), 0))
	params[0] = *zresource_ptr
	if arg2 != nil {
		params[1] = *arg2
	}
	if zend._zendGetParametersArrayEx(pass_num_args, params+g.Cond(arg2 != nil, 2, 1)) != zend.SUCCESS {
		zend._efree(params)
		zend.ZendWrongParamCount()
		return zend.FAILURE
	}
	&retval.u1.type_info = 0
	fci.size = g.SizeOf("fci")
	fci.object = nil
	fci.retval = &retval
	fci.param_count = num_args
	fci.params = params
	fci.no_separation = 1
	var __z *zend.Zval = &fci.function_name
	var __s *zend.ZendString = func_ptr.common.function_name
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	fcic.function_handler = func_ptr
	fcic.called_scope = nil
	fcic.object = nil
	result = zend.ZendCallFunction(&fci, &fcic)
	if result == zend.FAILURE || retval.u1.v.type_ == 0 {
		return_value.u1.type_info = 2
	} else {
		var __z *zend.Zval = return_value
		var __zv *zend.Zval = &retval
		if __zv.u1.v.type_ != 10 {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = __zv
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		} else {
			var _z1 *zend.Zval = __z
			var _z2 *zend.Zval = &(*__zv).value.ref.val
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
			if (_t & 0xff00) != 0 {
				zend.ZendGcAddref(&_gc.gc)
			}
			zend.ZvalPtrDtor(__zv)
		}
	}
	zend._efree(params)
	return result
}

// #define FileFunctionCall(func_name,pass_num_args,arg2) { zend_function * func_ptr ; func_ptr = ( zend_function * ) zend_hash_str_find_ptr ( EG ( function_table ) , # func_name , sizeof ( # func_name ) - 1 ) ; if ( func_ptr == NULL ) { zend_throw_exception_ex ( spl_ce_RuntimeException , 0 , "Internal error, function '%s' not found. Please report" , # func_name ) ; return ; } spl_filesystem_file_call ( intern , func_ptr , pass_num_args , return_value , arg2 ) ; }

func SplFilesystemFileReadCsv(intern *SplFilesystemObject, delimiter byte, enclosure byte, escape int, return_value *zend.Zval) int {
	var ret int = zend.SUCCESS
	var value *zend.Zval
	for {
		ret = SplFilesystemFileRead(intern, 1)
		if !(ret == zend.SUCCESS && intern.GetCurrentLineLen() == 0 && g.Cond((intern.GetFlags()&0x4) != 0, 1, 0)) {
			break
		}
	}
	if ret == zend.SUCCESS {
		var buf_len int = intern.GetCurrentLineLen()
		var buf *byte = zend._estrndup(intern.GetCurrentLine(), buf_len)
		if intern.u.file.current_zval.u1.v.type_ != 0 {
			zend.ZvalPtrDtor(&intern.u.file.current_zval)
			&intern.u.file.current_zval.u1.type_info = 0
		}
		standard.PhpFgetcsv(intern.GetStream(), delimiter, enclosure, escape, buf_len, buf, &intern.u.file.current_zval)
		if return_value != nil {
			value = &intern.u.file.current_zval
			var _z3 *zend.Zval = value
			if (_z3.u1.type_info & 0xff00) != 0 {
				if (_z3.u1.type_info & 0xff) == 10 {
					_z3 = &(*_z3).value.ref.val
					if (_z3.u1.type_info & 0xff00) != 0 {
						zend.ZvalAddrefP(_z3)
					}
				} else {
					zend.ZvalAddrefP(_z3)
				}
			}
			var _z1 *zend.Zval = return_value
			var _z2 *zend.Zval = _z3
			var _gc *zend.ZendRefcounted = _z2.value.counted
			var _t uint32 = _z2.u1.type_info
			_z1.value.counted = _gc
			_z1.u1.type_info = _t
		}
	}
	return ret
}

/* }}} */

func SplFilesystemFileReadLineEx(this_ptr *zend.Zval, intern *SplFilesystemObject, silent int) int {
	var retval zend.Zval

	/* 1) use fgetcsv? 2) overloaded call the function, 3) do it directly */

	if g.Cond((intern.GetFlags()&0x8) != 0, 1, 0) || intern.GetFuncGetCurr().common.scope != spl_ce_SplFileObject {
		if streams._phpStreamEof(intern.GetStream()) != 0 {
			if silent == 0 {
				zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot read from file %s", intern.GetFileName())
			}
			return zend.FAILURE
		}
		if g.Cond((intern.GetFlags()&0x8) != 0, 1, 0) {
			return SplFilesystemFileReadCsv(intern, intern.GetDelimiter(), intern.GetEnclosure(), intern.GetEscape(), nil)
		} else {
			var execute_data *zend.ZendExecuteData = zend.EG.current_execute_data
			zend.ZendCallMethod(this_ptr, &(execute_data.This).value.obj.ce, &intern.u.file.func_getCurr, "getCurrentLine", g.SizeOf("\"getCurrentLine\"")-1, &retval, 0, nil, nil)
		}
		if retval.u1.v.type_ != 0 {
			if intern.GetCurrentLine() != nil || intern.u.file.current_zval.u1.v.type_ != 0 {
				intern.GetCurrentLineNum()++
			}
			SplFilesystemFileFreeLine(intern)
			if retval.u1.v.type_ == 6 {
				intern.SetCurrentLine(zend._estrndup(retval.value.str.val, retval.value.str.len_))
				intern.SetCurrentLineLen(retval.value.str.len_)
			} else {
				var value *zend.Zval = &retval
				var _z3 *zend.Zval = value
				if (_z3.u1.type_info & 0xff00) != 0 {
					if (_z3.u1.type_info & 0xff) == 10 {
						_z3 = &(*_z3).value.ref.val
						if (_z3.u1.type_info & 0xff00) != 0 {
							zend.ZvalAddrefP(_z3)
						}
					} else {
						zend.ZvalAddrefP(_z3)
					}
				}
				var _z1 *zend.Zval = &intern.u.file.current_zval
				var _z2 *zend.Zval = _z3
				var _gc *zend.ZendRefcounted = _z2.value.counted
				var _t uint32 = _z2.u1.type_info
				_z1.value.counted = _gc
				_z1.u1.type_info = _t
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
	} else if intern.u.file.current_zval.u1.v.type_ != 0 {
		switch intern.u.file.current_zval.u1.v.type_ {
		case 6:
			return intern.u.file.current_zval.value.str.len_ == 0
		case 7:
			if g.Cond((intern.GetFlags()&0x8) != 0, 1, 0) && intern.u.file.current_zval.value.arr.nNumOfElements == 1 {
				var idx uint32 = 0
				var first *zend.Zval
				for intern.u.file.current_zval.value.arr.arData[idx].val.u1.v.type_ == 0 {
					idx++
				}
				first = &(intern.GetCurrentZval()).value.arr.arData[idx].val
				return first.u1.v.type_ == 6 && first.value.str.len_ == 0
			}
			return intern.u.file.current_zval.value.arr.nNumOfElements == 0
		case 1:
			return 1
		default:
			return 0
		}
	} else {
		return 1
	}
}

/* }}} */

func SplFilesystemFileReadLine(this_ptr *zend.Zval, intern *SplFilesystemObject, silent int) int {
	var ret int = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	for g.Cond((intern.GetFlags()&0x4) != 0, 1, 0) && ret == zend.SUCCESS && SplFilesystemFileIsEmptyLine(intern) != 0 {
		SplFilesystemFileFreeLine(intern)
		ret = SplFilesystemFileReadLineEx(this_ptr, intern, silent)
	}
	return ret
}

/* }}} */

func SplFilesystemFileRewind(this_ptr *zend.Zval, intern *SplFilesystemObject) {
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if -1 == streams._phpStreamSeek(intern.GetStream(), 0, SEEK_SET) {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Cannot rewind file %s", intern.GetFileName())
	} else {
		SplFilesystemFileFreeLine(intern)
		intern.SetCurrentLineNum(0)
	}
	if g.Cond((intern.GetFlags()&0x2) != 0, 1, 0) {
		SplFilesystemFileReadLine(this_ptr, intern, 1)
	}
}

/* {{{ proto SplFileObject::__construct(string filename [, string mode = 'r' [, bool use_include_path  [, resource context]]]])
   Construct a new file object */

func zim_spl_SplFileObject___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var use_include_path zend.ZendBool = 0
	var p1 *byte
	var p2 *byte
	var tmp_path *byte
	var tmp_path_len int
	var error_handling zend.ZendErrorHandling
	intern.SetOpenMode(nil)
	intern.SetOpenModeLen(0)
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "p|sbr!", &intern.file_name, &intern.file_name_len, &intern.u.file.open_mode, &intern.u.file.open_mode_len, &use_include_path, &intern.u.file.zcontext) == zend.FAILURE {
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
		tmp_path_len = strlen(intern.GetStream().orig_path)
		if tmp_path_len > 1 && intern.GetStream().orig_path[tmp_path_len-1] == '/' {
			tmp_path_len--
		}
		tmp_path = zend._estrndup(intern.GetStream().orig_path, tmp_path_len)
		p1 = strrchr(tmp_path, '/')
		p2 = 0
		if p1 != nil || p2 != nil {
			intern.SetPathLen(g.Cond(p1 > p2, p1, p2) - tmp_path)
		} else {
			intern.SetPathLen(0)
		}
		zend._efree(tmp_path)
		intern.SetPath(zend._estrndup(intern.GetStream().orig_path, intern.GetPathLen()))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* {{{ proto SplTempFileObject::__construct([int max_memory])
   Construct a new temp file object */

func zim_spl_SplTempFileObject___construct(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var max_memory zend.ZendLong = 2 * 1024 * 1024
	var tmp_fname []byte
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var error_handling zend.ZendErrorHandling
	if zend.ZendParseParametersThrow(execute_data.This.u2.num_args, "|l", &max_memory) == zend.FAILURE {
		return
	}
	if max_memory < 0 {
		intern.SetFileName("php://memory")
		intern.SetFileNameLen(12)
	} else if execute_data.This.u2.num_args != 0 {
		intern.SetFileNameLen(core.ApPhpSlprintf(tmp_fname, g.SizeOf("tmp_fname"), "php://temp/maxmemory:"+"%"+"lld", max_memory))
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
		intern.SetPath(zend._estrndup("", 0))
	}
	zend.ZendRestoreErrorHandling(&error_handling)
}

/* {{{ proto void SplFileObject::rewind()
   Rewind the file and read the first line */

func zim_spl_SplFileObject_rewind(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplFilesystemFileRewind(&(execute_data.This), intern)
}

/* {{{ proto void SplFileObject::eof()
   Return whether end of file is reached */

func zim_spl_SplFileObject_eof(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if streams._phpStreamEof(intern.GetStream()) != 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto void SplFileObject::valid()
   Return !eof() */

func zim_spl_SplFileObject_valid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if g.Cond((intern.GetFlags()&0x2) != 0, 1, 0) {
		if intern.GetCurrentLine() != nil || intern.u.file.current_zval.u1.v.type_ != 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
		return
	} else {
		if intern.GetStream() == nil {
			return_value.u1.type_info = 2
			return
		}
		if streams._phpStreamEof(intern.GetStream()) == 0 {
			return_value.u1.type_info = 3
		} else {
			return_value.u1.type_info = 2
		}
	}
}

/* {{{ proto string SplFileObject::fgets()
   Rturn next line from file */

func zim_spl_SplFileObject_fgets(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if SplFilesystemFileRead(intern, 0) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = zend.ZendStringInit(intern.GetCurrentLine(), intern.GetCurrentLineLen(), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	return
}

/* {{{ proto string SplFileObject::current()
   Return current line from file */

func zim_spl_SplFileObject_current(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if intern.GetCurrentLine() == nil && intern.u.file.current_zval.u1.v.type_ == 0 {
		SplFilesystemFileReadLine(&(execute_data.This), intern, 1)
	}
	if intern.GetCurrentLine() != nil && (!(g.Cond((intern.GetFlags()&0x8) != 0, 1, 0)) || intern.u.file.current_zval.u1.v.type_ == 0) {
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(intern.GetCurrentLine(), intern.GetCurrentLineLen(), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	} else if intern.u.file.current_zval.u1.v.type_ != 0 {
		var value *zend.Zval = &intern.u.file.current_zval
		var _z3 *zend.Zval = value
		if (_z3.u1.type_info & 0xff00) != 0 {
			if (_z3.u1.type_info & 0xff) == 10 {
				_z3 = &(*_z3).value.ref.val
				if (_z3.u1.type_info & 0xff00) != 0 {
					zend.ZvalAddrefP(_z3)
				}
			} else {
				zend.ZvalAddrefP(_z3)
			}
		}
		var _z1 *zend.Zval = return_value
		var _z2 *zend.Zval = _z3
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		return
	}
	return_value.u1.type_info = 2
	return
}

/* {{{ proto int SplFileObject::key()
   Return line number */

func zim_spl_SplFileObject_key(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}

	/*    Do not read the next line to support correct counting with fgetc()
	      if (!intern->current_line) {
	          spl_filesystem_file_read_line(ZEND_THIS, intern, 1);
	      } */

	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetCurrentLineNum()
	__z.u1.type_info = 4
	return
}

/* {{{ proto void SplFileObject::next()
   Read next line */

func zim_spl_SplFileObject_next(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	SplFilesystemFileFreeLine(intern)
	if g.Cond((intern.GetFlags()&0x2) != 0, 1, 0) {
		SplFilesystemFileReadLine(&(execute_data.This), intern, 1)
	}
	intern.GetCurrentLineNum()++
}

/* {{{ proto void SplFileObject::setFlags(int flags)
   Set file handling flags */

func zim_spl_SplFileObject_setFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &intern.flags) == zend.FAILURE {
		return
	}
}

/* {{{ proto int SplFileObject::getFlags()
   Get file handling flags */

func zim_spl_SplFileObject_getFlags(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = intern.GetFlags() & 0xf
	__z.u1.type_info = 4
	return
}

/* {{{ proto void SplFileObject::setMaxLineLen(int max_len)
   Set maximum line length */

func zim_spl_SplFileObject_setMaxLineLen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var max_len zend.ZendLong
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &max_len) == zend.FAILURE {
		return
	}
	if max_len < 0 {
		zend.ZendThrowExceptionEx(spl_ce_DomainException, 0, "Maximum line length must be greater than or equal zero")
		return
	}
	intern.SetMaxLineLen(max_len)
}

/* {{{ proto int SplFileObject::getMaxLineLen()
   Get maximum line length */

func zim_spl_SplFileObject_getMaxLineLen(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = zend.ZendLong(intern.GetMaxLineLen())
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool SplFileObject::hasChildren()
   Return false */

func zim_spl_SplFileObject_hasChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	return_value.u1.type_info = 2
	return
}

/* {{{ proto bool SplFileObject::getChildren()
   Read NULL */

func zim_spl_SplFileObject_getChildren(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
}

/* {{{ FileFunction */

// #define FileFunction(func_name) SPL_METHOD ( SplFileObject , func_name ) { spl_filesystem_object * intern = Z_SPLFILESYSTEM_P ( ZEND_THIS ) ; FileFunctionCall ( func_name , ZEND_NUM_ARGS ( ) , NULL ) ; }

/* }}} */

func zim_spl_SplFileObject_fgetcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var delimiter byte = intern.GetDelimiter()
	var enclosure byte = intern.GetEnclosure()
	var escape int = intern.GetEscape()
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		if intern.GetStream() == nil {
			zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
			return
		}
		switch execute_data.This.u2.num_args {
		case 3:
			if esc_len > 1 {
				core.PhpErrorDocref(nil, 1<<1, "escape must be empty or a single character")
				return_value.u1.type_info = 2
				return
			}
			if esc_len == 0 {
				escape = EOF
			} else {
				escape = uint8(esc[0])
			}
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "enclosure must be a character")
				return_value.u1.type_info = 2
				return
			}
			enclosure = enclo[0]
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "delimiter must be a character")
				return_value.u1.type_info = 2
				return
			}
			delimiter = delim[0]
		case 0:
			break
		}
		SplFilesystemFileReadCsv(intern, delimiter, enclosure, escape, return_value)
	}
}

/* }}} */

func zim_spl_SplFileObject_fputcsv(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
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
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "a|sss", &fields, &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		switch execute_data.This.u2.num_args {
		case 4:
			switch esc_len {
			case 0:
				escape = EOF
				break
			case 1:
				escape = uint8(esc[0])
				break
			default:
				core.PhpErrorDocref(nil, 1<<1, "escape must be empty or a single character")
				return_value.u1.type_info = 2
				return
			}
		case 3:
			if e_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "enclosure must be a character")
				return_value.u1.type_info = 2
				return
			}
			enclosure = enclo[0]
		case 2:
			if d_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "delimiter must be a character")
				return_value.u1.type_info = 2
				return
			}
			delimiter = delim[0]
		case 1:

		case 0:
			break
		}
		ret = standard.PhpFputcsv(intern.GetStream(), fields, delimiter, enclosure, escape)
		if ret < 0 {
			return_value.u1.type_info = 2
			return
		}
		var __z *zend.Zval = return_value
		__z.value.lval = ret
		__z.u1.type_info = 4
		return
	}
}

/* }}} */

func zim_spl_SplFileObject_setCsvControl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var delimiter byte = ','
	var enclosure byte = '"'
	var escape int = uint8('\\')
	var delim *byte = nil
	var enclo *byte = nil
	var esc *byte = nil
	var d_len int = 0
	var e_len int = 0
	var esc_len int = 0
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "|sss", &delim, &d_len, &enclo, &e_len, &esc, &esc_len) == zend.SUCCESS {
		switch execute_data.This.u2.num_args {
		case 3:
			switch esc_len {
			case 0:
				escape = EOF
				break
			case 1:
				escape = uint8(esc[0])
				break
			default:
				core.PhpErrorDocref(nil, 1<<1, "escape must be empty or a single character")
				return_value.u1.type_info = 2
				return
			}
		case 2:
			if e_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "enclosure must be a character")
				return_value.u1.type_info = 2
				return
			}
			enclosure = enclo[0]
		case 1:
			if d_len != 1 {
				core.PhpErrorDocref(nil, 1<<1, "delimiter must be a character")
				return_value.u1.type_info = 2
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

/* }}} */

func zim_spl_SplFileObject_getCsvControl(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var delimiter []byte
	var enclosure []byte
	var escape []byte
	var __arr *zend.ZendArray = zend._zendNewArray(0)
	var __z *zend.Zval = return_value
	__z.value.arr = __arr
	__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
	delimiter[0] = intern.GetDelimiter()
	delimiter[1] = '0'
	enclosure[0] = intern.GetEnclosure()
	enclosure[1] = '0'
	if intern.GetEscape() == EOF {
		escape[0] = '0'
	} else {
		escape[0] = uint8(intern.GetEscape())
		escape[1] = '0'
	}
	zend.AddNextIndexString(return_value, delimiter)
	zend.AddNextIndexString(return_value, enclosure)
	zend.AddNextIndexString(return_value, escape)
}

/* }}} */

func zim_spl_SplFileObject_flock(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG.function_table, "flock", g.SizeOf("\"flock\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "flock")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, execute_data.This.u2.num_args, return_value, nil)
}

/* }}} */

func zim_spl_SplFileObject_fflush(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if streams._phpStreamFlush(intern.GetStream(), 0) == 0 {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto int SplFileObject::ftell()
   Return current file position */

func zim_spl_SplFileObject_ftell(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var ret zend.ZendLong
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	ret = streams._phpStreamTell(intern.GetStream())
	if ret == -1 {
		return_value.u1.type_info = 2
		return
	} else {
		var __z *zend.Zval = return_value
		__z.value.lval = ret
		__z.u1.type_info = 4
		return
	}
}

/* {{{ proto int SplFileObject::fseek(int pos [, int whence = SEEK_SET])
   Return current file position */

func zim_spl_SplFileObject_fseek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var pos zend.ZendLong
	var whence zend.ZendLong = SEEK_SET
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l|l", &pos, &whence) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	var __z *zend.Zval = return_value
	__z.value.lval = streams._phpStreamSeek(intern.GetStream(), pos, int(whence))
	__z.u1.type_info = 4
	return
}

/* {{{ proto int SplFileObject::fgetc()
   Get a character form the file */

func zim_spl_SplFileObject_fgetc(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var buf []byte
	var result int
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	result = streams._phpStreamGetc(intern.GetStream())
	if result == EOF {
		return_value.u1.type_info = 2
	} else {
		if result == '\n' {
			intern.GetCurrentLineNum()++
		}
		buf[0] = result
		buf[1] = '0'
		var __z *zend.Zval = return_value
		var __s *zend.ZendString = zend.ZendStringInit(buf, 1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		return
	}
}

/* {{{ proto string SplFileObject::fgetss([string allowable_tags])
   Get a line from file pointer and strip HTML tags */

func zim_spl_SplFileObject_fgetss(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var arg2 zend.Zval
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if intern.GetMaxLineLen() > 0 {
		var __z *zend.Zval = &arg2
		__z.value.lval = intern.GetMaxLineLen()
		__z.u1.type_info = 4
	} else {
		var __z *zend.Zval = &arg2
		__z.value.lval = 1024
		__z.u1.type_info = 4
	}
	SplFilesystemFileFreeLine(intern)
	intern.GetCurrentLineNum()++
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG.function_table, "fgetss", g.SizeOf("\"fgetss\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fgetss")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, execute_data.This.u2.num_args, return_value, &arg2)
}

/* {{{ proto int SplFileObject::fpassthru()
   Output all remaining data from a file pointer */

func zim_spl_SplFileObject_fpassthru(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = streams._phpStreamPassthru(intern.GetStream())
	__z.u1.type_info = 4
	return
}

/* {{{ proto bool SplFileObject::fscanf(string format [, string ...])
   Implements a mostly ANSI compatible fscanf() */

func zim_spl_SplFileObject_fscanf(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	SplFilesystemFileFreeLine(intern)
	intern.GetCurrentLineNum()++
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG.function_table, "fscanf", g.SizeOf("\"fscanf\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fscanf")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, execute_data.This.u2.num_args, return_value, nil)
}

/* }}} */

func zim_spl_SplFileObject_fwrite(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var str *byte
	var str_len int
	var length zend.ZendLong = 0
	var written ssize_t
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "s|l", &str, &str_len, &length) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if execute_data.This.u2.num_args > 1 {
		if length >= 0 {
			if int(length) < str_len {
				str_len = int(length)
			} else {
				str_len = str_len
			}
		} else {

			/* Negative length given, nothing to write */

			str_len = 0

			/* Negative length given, nothing to write */

		}
	}
	if str_len == 0 {
		var __z *zend.Zval = return_value
		__z.value.lval = 0
		__z.u1.type_info = 4
		return
	}
	written = streams._phpStreamWrite(intern.GetStream(), str, str_len)
	if written < 0 {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	__z.value.lval = written
	__z.u1.type_info = 4
	return
}
func zim_spl_SplFileObject_fread(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var length zend.ZendLong = 0
	var str *zend.ZendString
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &length) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if length <= 0 {
		core.PhpErrorDocref(nil, 1<<1, "Length parameter must be greater than 0")
		return_value.u1.type_info = 2
		return
	}
	str = streams.PhpStreamReadToStr(intern.GetStream(), length)
	if str == nil {
		return_value.u1.type_info = 2
		return
	}
	var __z *zend.Zval = return_value
	var __s *zend.ZendString = str
	__z.value.str = __s
	if (zend.ZvalGcFlags(__s.gc.u.type_info) & 1 << 6) != 0 {
		__z.u1.type_info = 6
	} else {
		__z.u1.type_info = 6 | 1<<0<<8
	}
	return
}

/* {{{ proto bool SplFileObject::fstat()
   Stat() on a filehandle */

func zim_spl_SplFileObject_fstat(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var func_ptr *zend.ZendFunction
	func_ptr = (*zend.ZendFunction)(zend.ZendHashStrFindPtr(zend.EG.function_table, "fstat", g.SizeOf("\"fstat\"")-1))
	if func_ptr == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Internal error, function '%s' not found. Please report", "fstat")
		return
	}
	SplFilesystemFileCall(intern, func_ptr, execute_data.This.u2.num_args, return_value, nil)
}

/* }}} */

func zim_spl_SplFileObject_ftruncate(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var size zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &size) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if !(g.Cond(streams._phpStreamSetOption(intern.GetStream(), 10, 0, nil) == 0, 1, 0)) {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Can't truncate file %s", intern.GetFileName())
		return_value.u1.type_info = 2
		return
	}
	if 0 == streams._phpStreamTruncateSetSize(intern.GetStream(), size) {
		return_value.u1.type_info = 3
	} else {
		return_value.u1.type_info = 2
	}
	return
}

/* {{{ proto void SplFileObject::seek(int line_pos)
   Seek to specified line */

func zim_spl_SplFileObject_seek(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var intern *SplFilesystemObject = SplFilesystemFromObj(&(execute_data.This).value.obj)
	var line_pos zend.ZendLong
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "l", &line_pos) == zend.FAILURE {
		return
	}
	if intern.GetStream() == nil {
		zend.ZendThrowExceptionEx(spl_ce_RuntimeException, 0, "Object not initialized")
		return
	}
	if line_pos < 0 {
		zend.ZendThrowExceptionEx(spl_ce_LogicException, 0, "Can't seek file %s to negative line "+"%"+"lld", intern.GetFileName(), line_pos)
		return_value.u1.type_info = 2
		return
	}
	SplFilesystemFileRewind(&(execute_data.This), intern)
	for intern.GetCurrentLineNum() < line_pos {
		if SplFilesystemFileReadLine(&(execute_data.This), intern, 1) == zend.FAILURE {
			break
		}
	}
}

/* {{{ Function/Class/Method definitions */

var ArginfoFileObjectConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"file_name", 0, 0, 0}, {"open_mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var arginfo_file_object_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var arginfo_file_object_setMaxLineLen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(-1)), 0, 0, 0}, {"max_len", 0, 0, 0}}
var ArginfoFileObjectFgetcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoFileObjectFputcsv []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"fields", 0, 0, 0}, {"delimiter", 0, 0, 0}, {"enclosure", 0, 0, 0}, {"escape", 0, 0, 0}}
var ArginfoFileObjectFlock []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"operation", 0, 0, 0}, {"wouldblock", 0, 1, 0}}
var ArginfoFileObjectFseek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"pos", 0, 0, 0}, {"whence", 0, 0, 0}}
var ArginfoFileObjectFgetss []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"allowable_tags", 0, 0, 0}}
var ArginfoFileObjectFscanf []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"format", 0, 0, 0}, {"vars", 0, 1, 1}}
var ArginfoFileObjectFwrite []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"str", 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFileObjectFread []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"length", 0, 0, 0}}
var ArginfoFileObjectFtruncate []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"size", 0, 0, 0}}
var ArginfoFileObjectSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"line_pos", 0, 0, 0}}
var spl_SplFileObject_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplFileObject___construct,
		ArginfoFileObjectConstruct,
		uint32(g.SizeOf("arginfo_file_object___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"rewind",
		zim_spl_SplFileObject_rewind,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"eof",
		zim_spl_SplFileObject_eof,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"valid",
		zim_spl_SplFileObject_valid,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fgets",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fgetcsv",
		zim_spl_SplFileObject_fgetcsv,
		ArginfoFileObjectFgetcsv,
		uint32(g.SizeOf("arginfo_file_object_fgetcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fputcsv",
		zim_spl_SplFileObject_fputcsv,
		ArginfoFileObjectFputcsv,
		uint32(g.SizeOf("arginfo_file_object_fputcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setCsvControl",
		zim_spl_SplFileObject_setCsvControl,
		ArginfoFileObjectFgetcsv,
		uint32(g.SizeOf("arginfo_file_object_fgetcsv")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getCsvControl",
		zim_spl_SplFileObject_getCsvControl,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"flock",
		zim_spl_SplFileObject_flock,
		ArginfoFileObjectFlock,
		uint32(g.SizeOf("arginfo_file_object_flock")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fflush",
		zim_spl_SplFileObject_fflush,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"ftell",
		zim_spl_SplFileObject_ftell,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fseek",
		zim_spl_SplFileObject_fseek,
		ArginfoFileObjectFseek,
		uint32(g.SizeOf("arginfo_file_object_fseek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fgetc",
		zim_spl_SplFileObject_fgetc,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fpassthru",
		zim_spl_SplFileObject_fpassthru,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fgetss",
		zim_spl_SplFileObject_fgetss,
		ArginfoFileObjectFgetss,
		uint32(g.SizeOf("arginfo_file_object_fgetss")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fscanf",
		zim_spl_SplFileObject_fscanf,
		ArginfoFileObjectFscanf,
		uint32(g.SizeOf("arginfo_file_object_fscanf")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fwrite",
		zim_spl_SplFileObject_fwrite,
		ArginfoFileObjectFwrite,
		uint32(g.SizeOf("arginfo_file_object_fwrite")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fread",
		zim_spl_SplFileObject_fread,
		ArginfoFileObjectFread,
		uint32(g.SizeOf("arginfo_file_object_fread")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"fstat",
		zim_spl_SplFileObject_fstat,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"ftruncate",
		zim_spl_SplFileObject_ftruncate,
		ArginfoFileObjectFtruncate,
		uint32(g.SizeOf("arginfo_file_object_ftruncate")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"current",
		zim_spl_SplFileObject_current,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"key",
		zim_spl_SplFileObject_key,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"next",
		zim_spl_SplFileObject_next,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setFlags",
		zim_spl_SplFileObject_setFlags,
		arginfo_file_object_setFlags,
		uint32(g.SizeOf("arginfo_file_object_setFlags")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getFlags",
		zim_spl_SplFileObject_getFlags,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"setMaxLineLen",
		zim_spl_SplFileObject_setMaxLineLen,
		arginfo_file_object_setMaxLineLen,
		uint32(g.SizeOf("arginfo_file_object_setMaxLineLen")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getMaxLineLen",
		zim_spl_SplFileObject_getMaxLineLen,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"hasChildren",
		zim_spl_SplFileObject_hasChildren,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getChildren",
		zim_spl_SplFileObject_getChildren,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"seek",
		zim_spl_SplFileObject_seek,
		ArginfoFileObjectSeek,
		uint32(g.SizeOf("arginfo_file_object_seek")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"getCurrentLine",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{
		"__toString",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(g.SizeOf("arginfo_splfileinfo_void")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoTempFileObjectConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"max_memory", 0, 0, 0}}
var spl_SplTempFileObject_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplTempFileObject___construct,
		ArginfoTempFileObjectConstruct,
		uint32(g.SizeOf("arginfo_temp_file_object___construct")/g.SizeOf("struct _zend_internal_arg_info") - 1),
		1 << 0,
	},
	{nil, nil, nil, 0, 0},
}

/* }}} */

func ZmStartupSplDirectory(type_ int, module_number int) int {
	SplRegisterStdClass(&spl_ce_SplFileInfo, "SplFileInfo", SplFilesystemObjectNew, spl_SplFileInfo_functions)
	memcpy(&SplFilesystemObjectHandlers, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
	SplFilesystemObjectHandlers.offset = zend_long((*byte)(&((*SplFilesystemObject)(nil).GetStd())) - (*byte)(nil))
	SplFilesystemObjectHandlers.clone_obj = SplFilesystemObjectClone
	SplFilesystemObjectHandlers.cast_object = SplFilesystemObjectCast
	SplFilesystemObjectHandlers.dtor_obj = SplFilesystemObjectDestroyObject
	SplFilesystemObjectHandlers.free_obj = SplFilesystemObjectFreeStorage
	spl_ce_SplFileInfo.serialize = zend.ZendClassSerializeDeny
	spl_ce_SplFileInfo.unserialize = zend.ZendClassUnserializeDeny
	SplRegisterSubClass(&spl_ce_DirectoryIterator, spl_ce_SplFileInfo, "DirectoryIterator", SplFilesystemObjectNew, spl_DirectoryIterator_functions)
	zend.ZendClassImplements(spl_ce_DirectoryIterator, 1, zend.ZendCeIterator)
	zend.ZendClassImplements(spl_ce_DirectoryIterator, 1, spl_ce_SeekableIterator)
	spl_ce_DirectoryIterator.get_iterator = SplFilesystemDirGetIterator
	SplRegisterSubClass(&spl_ce_FilesystemIterator, spl_ce_DirectoryIterator, "FilesystemIterator", SplFilesystemObjectNew, spl_FilesystemIterator_functions)
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_MODE_MASK", g.SizeOf("\"CURRENT_MODE_MASK\"")-1, zend.ZendLong(0xf0))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_PATHNAME", g.SizeOf("\"CURRENT_AS_PATHNAME\"")-1, zend.ZendLong(0x20))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_FILEINFO", g.SizeOf("\"CURRENT_AS_FILEINFO\"")-1, zend.ZendLong(0x0))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "CURRENT_AS_SELF", g.SizeOf("\"CURRENT_AS_SELF\"")-1, zend.ZendLong(0x10))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_MODE_MASK", g.SizeOf("\"KEY_MODE_MASK\"")-1, zend.ZendLong(0xf00))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_AS_PATHNAME", g.SizeOf("\"KEY_AS_PATHNAME\"")-1, zend.ZendLong(0x0))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "FOLLOW_SYMLINKS", g.SizeOf("\"FOLLOW_SYMLINKS\"")-1, zend.ZendLong(0x200))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "KEY_AS_FILENAME", g.SizeOf("\"KEY_AS_FILENAME\"")-1, zend.ZendLong(0x100))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "NEW_CURRENT_AND_KEY", g.SizeOf("\"NEW_CURRENT_AND_KEY\"")-1, zend.ZendLong(0x100|0x0))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "OTHER_MODE_MASK", g.SizeOf("\"OTHER_MODE_MASK\"")-1, zend.ZendLong(0x3000))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "SKIP_DOTS", g.SizeOf("\"SKIP_DOTS\"")-1, zend.ZendLong(0x1000))
	zend.ZendDeclareClassConstantLong(spl_ce_FilesystemIterator, "UNIX_PATHS", g.SizeOf("\"UNIX_PATHS\"")-1, zend.ZendLong(0x2000))
	spl_ce_FilesystemIterator.get_iterator = SplFilesystemTreeGetIterator
	SplRegisterSubClass(&spl_ce_RecursiveDirectoryIterator, spl_ce_FilesystemIterator, "RecursiveDirectoryIterator", SplFilesystemObjectNew, spl_RecursiveDirectoryIterator_functions)
	zend.ZendClassImplements(spl_ce_RecursiveDirectoryIterator, 1, spl_ce_RecursiveIterator)
	memcpy(&SplFilesystemObjectCheckHandlers, &SplFilesystemObjectHandlers, g.SizeOf("zend_object_handlers"))
	SplFilesystemObjectCheckHandlers.clone_obj = nil
	SplFilesystemObjectCheckHandlers.get_method = SplFilesystemObjectGetMethodCheck
	SplRegisterSubClass(&spl_ce_GlobIterator, spl_ce_FilesystemIterator, "GlobIterator", SplFilesystemObjectNewCheck, spl_GlobIterator_functions)
	zend.ZendClassImplements(spl_ce_GlobIterator, 1, zend.ZendCeCountable)
	SplRegisterSubClass(&spl_ce_SplFileObject, spl_ce_SplFileInfo, "SplFileObject", SplFilesystemObjectNewCheck, spl_SplFileObject_functions)
	zend.ZendClassImplements(spl_ce_SplFileObject, 1, spl_ce_RecursiveIterator)
	zend.ZendClassImplements(spl_ce_SplFileObject, 1, spl_ce_SeekableIterator)
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "DROP_NEW_LINE", g.SizeOf("\"DROP_NEW_LINE\"")-1, zend.ZendLong(0x1))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "READ_AHEAD", g.SizeOf("\"READ_AHEAD\"")-1, zend.ZendLong(0x2))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "SKIP_EMPTY", g.SizeOf("\"SKIP_EMPTY\"")-1, zend.ZendLong(0x4))
	zend.ZendDeclareClassConstantLong(spl_ce_SplFileObject, "READ_CSV", g.SizeOf("\"READ_CSV\"")-1, zend.ZendLong(0x8))
	SplRegisterSubClass(&spl_ce_SplTempFileObject, spl_ce_SplFileObject, "SplTempFileObject", SplFilesystemObjectNewCheck, spl_SplTempFileObject_functions)
	return zend.SUCCESS
}

/* }}} */
