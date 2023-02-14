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

var SplFilesystemDirItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemDirItDtor, SplFilesystemDirItValid, SplFilesystemDirItCurrentData, SplFilesystemDirItCurrentKey, SplFilesystemDirItMoveForward, SplFilesystemDirItRewind, nil)
var SplFilesystemTreeItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemTreeItDtor, SplFilesystemDirItValid, SplFilesystemTreeItCurrentData, SplFilesystemTreeItCurrentKey, SplFilesystemTreeItMoveForward, SplFilesystemTreeItRewind, nil)
var ArginfoInfoConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("file_name"),
}
var arginfo_info_openFile []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("open_mode"),
	zend.MakeArgInfo("use_include_path"),
	zend.MakeArgInfo("context"),
}
var arginfo_info_optinalFileClass []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("class_name"),
}
var arginfo_optinalSuffix []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("suffix"),
}
var ArginfoSplfileinfoVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}

/* the method table */

var spl_SplFileInfo_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_SplFileInfo___construct, ArginfoInfoConstruct, uint32(b.SizeOf("arginfo_info___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getPath", zim_spl_SplFileInfo_getPath, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getFilename", zim_spl_SplFileInfo_getFilename, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getExtension", zim_spl_SplFileInfo_getExtension, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getBasename", zim_spl_SplFileInfo_getBasename, arginfo_optinalSuffix, uint32(b.SizeOf("arginfo_optinalSuffix")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getPathname", zim_spl_SplFileInfo_getPathname, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getPerms", zim_spl_SplFileInfo_getPerms, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getInode", zim_spl_SplFileInfo_getInode, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getSize", zim_spl_SplFileInfo_getSize, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getOwner", zim_spl_SplFileInfo_getOwner, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getGroup", zim_spl_SplFileInfo_getGroup, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getATime", zim_spl_SplFileInfo_getATime, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getMTime", zim_spl_SplFileInfo_getMTime, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getCTime", zim_spl_SplFileInfo_getCTime, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getType", zim_spl_SplFileInfo_getType, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isWritable", zim_spl_SplFileInfo_isWritable, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isReadable", zim_spl_SplFileInfo_isReadable, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isExecutable", zim_spl_SplFileInfo_isExecutable, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isFile", zim_spl_SplFileInfo_isFile, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isDir", zim_spl_SplFileInfo_isDir, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isLink", zim_spl_SplFileInfo_isLink, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getLinkTarget", zim_spl_SplFileInfo_getLinkTarget, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getRealPath", zim_spl_SplFileInfo_getRealPath, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getFileInfo", zim_spl_SplFileInfo_getFileInfo, arginfo_info_optinalFileClass, uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getPathInfo", zim_spl_SplFileInfo_getPathInfo, arginfo_info_optinalFileClass, uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("openFile", zim_spl_SplFileInfo_openFile, arginfo_info_openFile, uint32(b.SizeOf("arginfo_info_openFile")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setFileClass", zim_spl_SplFileInfo_setFileClass, arginfo_info_optinalFileClass, uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setInfoClass", zim_spl_SplFileInfo_setInfoClass, arginfo_info_optinalFileClass, uint32(b.SizeOf("arginfo_info_optinalFileClass")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__debugInfo", zim_spl_SplFileInfo___debugInfo, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("_bad_state_ex", zim_spl_SplFileInfo__bad_state_ex, nil, uint32(b.SizeOf("NULL")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_FINAL),
	zend.MakeZendFunctionEntry("__toString", zim_spl_SplFileInfo_getPathname, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoDirConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("path"),
}
var ArginfoDirItSeek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("position"),
}

/* the method table */

var spl_DirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_DirectoryIterator___construct, ArginfoDirConstruct, uint32(b.SizeOf("arginfo_dir___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getFilename", zim_spl_DirectoryIterator_getFilename, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getExtension", zim_spl_DirectoryIterator_getExtension, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getBasename", zim_spl_DirectoryIterator_getBasename, arginfo_optinalSuffix, uint32(b.SizeOf("arginfo_optinalSuffix")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("isDot", zim_spl_DirectoryIterator_isDot, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("rewind", zim_spl_DirectoryIterator_rewind, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("valid", zim_spl_DirectoryIterator_valid, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("key", zim_spl_DirectoryIterator_key, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("current", zim_spl_DirectoryIterator_current, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("next", zim_spl_DirectoryIterator_next, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("seek", zim_spl_DirectoryIterator_seek, ArginfoDirItSeek, uint32(b.SizeOf("arginfo_dir_it_seek")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__toString", zim_spl_DirectoryIterator_getFilename, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoRDirConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("path"),
	zend.MakeArgInfo("flags"),
}
var arginfo_r_dir_hasChildren []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("allow_links"),
}
var arginfo_r_dir_setFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("flags"),
}
var spl_FilesystemIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_FilesystemIterator___construct, ArginfoRDirConstruct, uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("rewind", zim_spl_FilesystemIterator_rewind, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("next", zim_spl_DirectoryIterator_next, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("key", zim_spl_FilesystemIterator_key, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("current", zim_spl_FilesystemIterator_current, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getFlags", zim_spl_FilesystemIterator_getFlags, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setFlags", zim_spl_FilesystemIterator_setFlags, arginfo_r_dir_setFlags, uint32(b.SizeOf("arginfo_r_dir_setFlags")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var spl_RecursiveDirectoryIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_RecursiveDirectoryIterator___construct, ArginfoRDirConstruct, uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("hasChildren", zim_spl_RecursiveDirectoryIterator_hasChildren, arginfo_r_dir_hasChildren, uint32(b.SizeOf("arginfo_r_dir_hasChildren")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getChildren", zim_spl_RecursiveDirectoryIterator_getChildren, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getSubPath", zim_spl_RecursiveDirectoryIterator_getSubPath, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getSubPathname", zim_spl_RecursiveDirectoryIterator_getSubPathname, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var spl_GlobIterator_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_GlobIterator___construct, ArginfoRDirConstruct, uint32(b.SizeOf("arginfo_r_dir___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("count", zim_spl_GlobIterator_count, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
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

var ArginfoFileObjectConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("file_name"),
	zend.MakeArgInfo("open_mode"),
	zend.MakeArgInfo("use_include_path"),
	zend.MakeArgInfo("context"),
}
var arginfo_file_object_setFlags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("flags"),
}
var arginfo_file_object_setMaxLineLen []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("max_len"),
}
var ArginfoFileObjectFgetcsv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("delimiter"),
	zend.MakeArgInfo("enclosure"),
	zend.MakeArgInfo("escape"),
}
var ArginfoFileObjectFputcsv []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("fields"),
	zend.MakeArgInfo("delimiter"),
	zend.MakeArgInfo("enclosure"),
	zend.MakeArgInfo("escape"),
}
var ArginfoFileObjectFlock []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("operation"),
	zend.MakeArgInfo("wouldblock",ArgInfoByRef(1)),
}
var ArginfoFileObjectFseek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("pos"),
	zend.MakeArgInfo("whence"),
}
var ArginfoFileObjectFgetss []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("allowable_tags"),
}
var ArginfoFileObjectFscanf []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("format"),
	zend.MakeArgInfo("vars",ArgInfoByRef(1),ArgInfoVariadic()),
}
var ArginfoFileObjectFwrite []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("str"),
	zend.MakeArgInfo("length"),
}
var ArginfoFileObjectFread []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("length"),
}
var ArginfoFileObjectFtruncate []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("size"),
}
var ArginfoFileObjectSeek []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("line_pos"),
}
var spl_SplFileObject_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_SplFileObject___construct, ArginfoFileObjectConstruct, uint32(b.SizeOf("arginfo_file_object___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("rewind", zim_spl_SplFileObject_rewind, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("eof", zim_spl_SplFileObject_eof, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("valid", zim_spl_SplFileObject_valid, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fgets", zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fgetcsv", zim_spl_SplFileObject_fgetcsv, ArginfoFileObjectFgetcsv, uint32(b.SizeOf("arginfo_file_object_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fputcsv", zim_spl_SplFileObject_fputcsv, ArginfoFileObjectFputcsv, uint32(b.SizeOf("arginfo_file_object_fputcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setCsvControl", zim_spl_SplFileObject_setCsvControl, ArginfoFileObjectFgetcsv, uint32(b.SizeOf("arginfo_file_object_fgetcsv")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getCsvControl", zim_spl_SplFileObject_getCsvControl, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("flock", zim_spl_SplFileObject_flock, ArginfoFileObjectFlock, uint32(b.SizeOf("arginfo_file_object_flock")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fflush", zim_spl_SplFileObject_fflush, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("ftell", zim_spl_SplFileObject_ftell, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fseek", zim_spl_SplFileObject_fseek, ArginfoFileObjectFseek, uint32(b.SizeOf("arginfo_file_object_fseek")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fgetc", zim_spl_SplFileObject_fgetc, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fpassthru", zim_spl_SplFileObject_fpassthru, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fgetss", zim_spl_SplFileObject_fgetss, ArginfoFileObjectFgetss, uint32(b.SizeOf("arginfo_file_object_fgetss")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fscanf", zim_spl_SplFileObject_fscanf, ArginfoFileObjectFscanf, uint32(b.SizeOf("arginfo_file_object_fscanf")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fwrite", zim_spl_SplFileObject_fwrite, ArginfoFileObjectFwrite, uint32(b.SizeOf("arginfo_file_object_fwrite")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fread", zim_spl_SplFileObject_fread, ArginfoFileObjectFread, uint32(b.SizeOf("arginfo_file_object_fread")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("fstat", zim_spl_SplFileObject_fstat, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("ftruncate", zim_spl_SplFileObject_ftruncate, ArginfoFileObjectFtruncate, uint32(b.SizeOf("arginfo_file_object_ftruncate")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("current", zim_spl_SplFileObject_current, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("key", zim_spl_SplFileObject_key, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("next", zim_spl_SplFileObject_next, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setFlags", zim_spl_SplFileObject_setFlags, arginfo_file_object_setFlags, uint32(b.SizeOf("arginfo_file_object_setFlags")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getFlags", zim_spl_SplFileObject_getFlags, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("setMaxLineLen", zim_spl_SplFileObject_setMaxLineLen, arginfo_file_object_setMaxLineLen, uint32(b.SizeOf("arginfo_file_object_setMaxLineLen")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getMaxLineLen", zim_spl_SplFileObject_getMaxLineLen, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("hasChildren", zim_spl_SplFileObject_hasChildren, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getChildren", zim_spl_SplFileObject_getChildren, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("seek", zim_spl_SplFileObject_seek, ArginfoFileObjectSeek, uint32(b.SizeOf("arginfo_file_object_seek")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("getCurrentLine", zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry("__toString", zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid, uint32(b.SizeOf("arginfo_splfileinfo_void")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
var ArginfoTempFileObjectConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("max_memory"),
}
var spl_SplTempFileObject_functions []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntry("__construct", zim_spl_SplTempFileObject___construct, ArginfoTempFileObjectConstruct, uint32(b.SizeOf("arginfo_temp_file_object___construct")/b.SizeOf("struct _zend_internal_arg_info")-1), zend.ZEND_ACC_PUBLIC),
	zend.MakeZendFunctionEntry(nil, nil, nil, 0, 0),
}
