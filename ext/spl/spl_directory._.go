// <<generate>>

package spl

import (
	b "sik/builtin"
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

/* define an __special__  overloaded iterator structure */

const SPL_FILE_OBJECT_DROP_NEW_LINE = 0x1
const SPL_FILE_OBJECT_READ_AHEAD = 0x2
const SPL_FILE_OBJECT_SKIP_EMPTY = 0x4
const SPL_FILE_OBJECT_READ_CSV = 0x8
const SPL_FILE_OBJECT_MASK = 0xf
const SPL_FILE_DIR_CURRENT_AS_FILEINFO = 0x0
const SPL_FILE_DIR_CURRENT_AS_SELF = 0x10
const SPL_FILE_DIR_CURRENT_AS_PATHNAME = 0x20
const SPL_FILE_DIR_CURRENT_MODE_MASK = 0xf0
const SPL_FILE_DIR_KEY_AS_PATHNAME = 0x0
const SPL_FILE_DIR_KEY_AS_FILENAME = 0x100
const SPL_FILE_DIR_FOLLOW_SYMLINKS = 0x200
const SPL_FILE_DIR_KEY_MODE_MASK = 0xf00
const SPL_FILE_DIR_SKIPDOTS = 0x1000
const SPL_FILE_DIR_UNIXPATHS = 0x2000
const SPL_FILE_DIR_OTHERS_MASK = 0x3000

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

/* declare the class handlers */

var SplFilesystemObjectHandlers zend.ZendObjectHandlers

/* includes handler to validate object state when retrieving methods */

var SplFilesystemObjectCheckHandlers zend.ZendObjectHandlers

/* decalre the class entry */

/* {{{ spl_ce_dir_object_new */

/* {{{ spl_filesystem_object_clone */

const DIT_CTOR_FLAGS = 0x1
const DIT_CTOR_GLOB = 0x2

/* {{{ proto string DirectoryIterator::valid()
   Check whether dir contains more entries */

/* }}}*/

/* }}}*/

/* {{{ proto string SplFileInfo::getRealPath()
   Return the resolved path */

/* {{{ proto SplFileObject SplFileInfo::openFile([string mode = 'r' [, bool use_include_path  [, resource context]]])
   Open the current file */

/* {{{  proto SplFileInfo::_bad_state_ex(void) */

/* {{{ proto void FilesystemIterator::setFlags(long $flags)
   Set handling flags */

/* {{{ proto bool RecursiveDirectoryIterator::hasChildren([bool $allow_links = false])
   Returns whether current entry is a directory and not '.' or '..' */

/* {{{ proto GlobIterator::__construct(string path [, int flags])
Cronstructs a new dir iterator from a glob expression (no glob:// needed). */

/* {{{ forward declarations to the iterator handlers */

/* iterator handler table */

var SplFilesystemDirItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplFilesystemDirItDtor, SplFilesystemDirItValid, SplFilesystemDirItCurrentData, SplFilesystemDirItCurrentKey, SplFilesystemDirItMoveForward, SplFilesystemDirItRewind, nil}
var SplFilesystemTreeItFuncs zend.ZendObjectIteratorFuncs = zend.ZendObjectIteratorFuncs{SplFilesystemTreeItDtor, SplFilesystemDirItValid, SplFilesystemTreeItCurrentData, SplFilesystemTreeItCurrentKey, SplFilesystemTreeItMoveForward, SplFilesystemTreeItRewind, nil}
var ArginfoInfoConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"file_name", 0, 0, 0},
}
var arginfo_info_openFile []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"open_mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var arginfo_info_optinalFileClass []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"class_name", 0, 0, 0}}
var arginfo_optinalSuffix []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"suffix", 0, 0, 0}}
var ArginfoSplfileinfoVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/* the method table */

var spl_SplFileInfo_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplFileInfo___construct,
		ArginfoInfoConstruct,
		uint32(b.SizeOf("arginfo_info___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPath",
		zim_spl_SplFileInfo_getPath,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFilename",
		zim_spl_SplFileInfo_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getExtension",
		zim_spl_SplFileInfo_getExtension,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getBasename",
		zim_spl_SplFileInfo_getBasename,
		arginfo_optinalSuffix,
		uint32(b.SizeOf("arginfo_optinalSuffix")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPathname",
		zim_spl_SplFileInfo_getPathname,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPerms",
		zim_spl_SplFileInfo_getPerms,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getInode",
		zim_spl_SplFileInfo_getInode,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getSize",
		zim_spl_SplFileInfo_getSize,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getOwner",
		zim_spl_SplFileInfo_getOwner,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getGroup",
		zim_spl_SplFileInfo_getGroup,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getATime",
		zim_spl_SplFileInfo_getATime,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getMTime",
		zim_spl_SplFileInfo_getMTime,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getCTime",
		zim_spl_SplFileInfo_getCTime,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getType",
		zim_spl_SplFileInfo_getType,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isWritable",
		zim_spl_SplFileInfo_isWritable,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isReadable",
		zim_spl_SplFileInfo_isReadable,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isExecutable",
		zim_spl_SplFileInfo_isExecutable,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isFile",
		zim_spl_SplFileInfo_isFile,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isDir",
		zim_spl_SplFileInfo_isDir,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isLink",
		zim_spl_SplFileInfo_isLink,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getLinkTarget",
		zim_spl_SplFileInfo_getLinkTarget,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getRealPath",
		zim_spl_SplFileInfo_getRealPath,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFileInfo",
		zim_spl_SplFileInfo_getFileInfo,
		arginfo_info_optinalFileClass,
		uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getPathInfo",
		zim_spl_SplFileInfo_getPathInfo,
		arginfo_info_optinalFileClass,
		uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"openFile",
		zim_spl_SplFileInfo_openFile,
		arginfo_info_openFile,
		uint32(b.SizeOf("arginfo_info_openFile")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFileClass",
		zim_spl_SplFileInfo_setFileClass,
		arginfo_info_optinalFileClass,
		uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setInfoClass",
		zim_spl_SplFileInfo_setInfoClass,
		arginfo_info_optinalFileClass,
		uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__debugInfo",
		zim_spl_SplFileInfo___debugInfo,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"_bad_state_ex",
		zim_spl_SplFileInfo__bad_state_ex,
		nil,
		uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_FINAL,
	},
	{
		"__toString",
		zim_spl_SplFileInfo_getPathname,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoDirConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"path", 0, 0, 0},
}
var ArginfoDirItSeek []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"position", 0, 0, 0},
}

/* the method table */

var spl_DirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_DirectoryIterator___construct,
		ArginfoDirConstruct,
		uint32(b.SizeOf("arginfo_dir___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFilename",
		zim_spl_DirectoryIterator_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getExtension",
		zim_spl_DirectoryIterator_getExtension,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getBasename",
		zim_spl_DirectoryIterator_getBasename,
		arginfo_optinalSuffix,
		uint32(b.SizeOf("arginfo_optinalSuffix")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"isDot",
		zim_spl_DirectoryIterator_isDot,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_DirectoryIterator_rewind,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_DirectoryIterator_valid,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_DirectoryIterator_key,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_DirectoryIterator_current,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_DirectoryIterator_next,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"seek",
		zim_spl_DirectoryIterator_seek,
		ArginfoDirItSeek,
		uint32(b.SizeOf("arginfo_dir_it_seek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__toString",
		zim_spl_DirectoryIterator_getFilename,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
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
		uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_FilesystemIterator_rewind,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_DirectoryIterator_next,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_FilesystemIterator_key,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_FilesystemIterator_current,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_FilesystemIterator_getFlags,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_FilesystemIterator_setFlags,
		arginfo_r_dir_setFlags,
		uint32(b.SizeOf("arginfo_r_dir_setFlags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_RecursiveDirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_RecursiveDirectoryIterator___construct,
		ArginfoRDirConstruct,
		uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_RecursiveDirectoryIterator_hasChildren,
		arginfo_r_dir_hasChildren,
		uint32(b.SizeOf("arginfo_r_dir_hasChildren")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_RecursiveDirectoryIterator_getChildren,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getSubPath",
		zim_spl_RecursiveDirectoryIterator_getSubPath,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getSubPathname",
		zim_spl_RecursiveDirectoryIterator_getSubPathname,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var spl_GlobIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_GlobIterator___construct,
		ArginfoRDirConstruct,
		uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"count",
		zim_spl_GlobIterator_count,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ proto SplFileObject::__construct(string filename [, string mode = 'r' [, bool use_include_path  [, resource context]]]])
   Construct a new file object */

/* {{{ proto SplTempFileObject::__construct([int max_memory])
   Construct a new temp file object */

/* {{{ proto void SplFileObject::rewind()
   Rewind the file and read the first line */

/* {{{ proto void SplFileObject::eof()
   Return whether end of file is reached */

/* {{{ proto void SplFileObject::valid()
   Return !eof() */

/* {{{ proto string SplFileObject::fgets()
   Rturn next line from file */

/* {{{ proto string SplFileObject::current()
   Return current line from file */

/* {{{ proto int SplFileObject::key()
   Return line number */

/* {{{ proto void SplFileObject::next()
   Read next line */

/* {{{ proto void SplFileObject::setFlags(int flags)
   Set file handling flags */

/* {{{ proto int SplFileObject::getFlags()
   Get file handling flags */

/* {{{ proto void SplFileObject::setMaxLineLen(int max_len)
   Set maximum line length */

/* {{{ proto int SplFileObject::getMaxLineLen()
   Get maximum line length */

/* {{{ proto bool SplFileObject::hasChildren()
   Return false */

/* {{{ proto bool SplFileObject::getChildren()
   Read NULL */

/* {{{ FileFunction */

/* {{{ proto int SplFileObject::ftell()
   Return current file position */

/* {{{ proto int SplFileObject::fseek(int pos [, int whence = SEEK_SET])
   Return current file position */

/* {{{ proto int SplFileObject::fgetc()
   Get a character form the file */

/* {{{ proto string SplFileObject::fgetss([string allowable_tags])
   Get a line from file pointer and strip HTML tags */

/* {{{ proto int SplFileObject::fpassthru()
   Output all remaining data from a file pointer */

/* {{{ proto bool SplFileObject::fscanf(string format [, string ...])
   Implements a mostly ANSI compatible fscanf() */

/* {{{ proto bool SplFileObject::fstat()
   Stat() on a filehandle */

/* {{{ proto void SplFileObject::seek(int line_pos)
   Seek to specified line */

/* {{{ Function/Class/Method definitions */

var ArginfoFileObjectConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"file_name", 0, 0, 0}, {"open_mode", 0, 0, 0}, {"use_include_path", 0, 0, 0}, {"context", 0, 0, 0}}
var arginfo_file_object_setFlags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"flags", 0, 0, 0},
}
var arginfo_file_object_setMaxLineLen []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"max_len", 0, 0, 0},
}
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
		uint32(b.SizeOf("arginfo_file_object___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"rewind",
		zim_spl_SplFileObject_rewind,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"eof",
		zim_spl_SplFileObject_eof,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"valid",
		zim_spl_SplFileObject_valid,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fgets",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fgetcsv",
		zim_spl_SplFileObject_fgetcsv,
		ArginfoFileObjectFgetcsv,
		uint32(b.SizeOf("arginfo_file_object_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fputcsv",
		zim_spl_SplFileObject_fputcsv,
		ArginfoFileObjectFputcsv,
		uint32(b.SizeOf("arginfo_file_object_fputcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setCsvControl",
		zim_spl_SplFileObject_setCsvControl,
		ArginfoFileObjectFgetcsv,
		uint32(b.SizeOf("arginfo_file_object_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getCsvControl",
		zim_spl_SplFileObject_getCsvControl,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"flock",
		zim_spl_SplFileObject_flock,
		ArginfoFileObjectFlock,
		uint32(b.SizeOf("arginfo_file_object_flock")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fflush",
		zim_spl_SplFileObject_fflush,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"ftell",
		zim_spl_SplFileObject_ftell,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fseek",
		zim_spl_SplFileObject_fseek,
		ArginfoFileObjectFseek,
		uint32(b.SizeOf("arginfo_file_object_fseek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fgetc",
		zim_spl_SplFileObject_fgetc,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fpassthru",
		zim_spl_SplFileObject_fpassthru,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fgetss",
		zim_spl_SplFileObject_fgetss,
		ArginfoFileObjectFgetss,
		uint32(b.SizeOf("arginfo_file_object_fgetss")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fscanf",
		zim_spl_SplFileObject_fscanf,
		ArginfoFileObjectFscanf,
		uint32(b.SizeOf("arginfo_file_object_fscanf")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fwrite",
		zim_spl_SplFileObject_fwrite,
		ArginfoFileObjectFwrite,
		uint32(b.SizeOf("arginfo_file_object_fwrite")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fread",
		zim_spl_SplFileObject_fread,
		ArginfoFileObjectFread,
		uint32(b.SizeOf("arginfo_file_object_fread")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"fstat",
		zim_spl_SplFileObject_fstat,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"ftruncate",
		zim_spl_SplFileObject_ftruncate,
		ArginfoFileObjectFtruncate,
		uint32(b.SizeOf("arginfo_file_object_ftruncate")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"current",
		zim_spl_SplFileObject_current,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"key",
		zim_spl_SplFileObject_key,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"next",
		zim_spl_SplFileObject_next,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setFlags",
		zim_spl_SplFileObject_setFlags,
		arginfo_file_object_setFlags,
		uint32(b.SizeOf("arginfo_file_object_setFlags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getFlags",
		zim_spl_SplFileObject_getFlags,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"setMaxLineLen",
		zim_spl_SplFileObject_setMaxLineLen,
		arginfo_file_object_setMaxLineLen,
		uint32(b.SizeOf("arginfo_file_object_setMaxLineLen")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getMaxLineLen",
		zim_spl_SplFileObject_getMaxLineLen,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"hasChildren",
		zim_spl_SplFileObject_hasChildren,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getChildren",
		zim_spl_SplFileObject_getChildren,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"seek",
		zim_spl_SplFileObject_seek,
		ArginfoFileObjectSeek,
		uint32(b.SizeOf("arginfo_file_object_seek")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"getCurrentLine",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{
		"__toString",
		zim_spl_SplFileObject_fgets,
		ArginfoSplfileinfoVoid,
		uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
var ArginfoTempFileObjectConstruct []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(0)), 0, 0, 0}, {"max_memory", 0, 0, 0}}
var spl_SplTempFileObject_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_SplTempFileObject___construct,
		ArginfoTempFileObjectConstruct,
		uint32(b.SizeOf("arginfo_temp_file_object___construct")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC,
	},
	{nil, nil, nil, 0, 0},
}
