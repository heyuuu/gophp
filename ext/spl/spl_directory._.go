// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
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

var spl_ce_SplFileInfo *types.ClassEntry
var spl_ce_DirectoryIterator *types.ClassEntry
var spl_ce_FilesystemIterator *types.ClassEntry
var spl_ce_RecursiveDirectoryIterator *types.ClassEntry
var spl_ce_GlobIterator *types.ClassEntry
var spl_ce_SplFileObject *types.ClassEntry
var spl_ce_SplTempFileObject *types.ClassEntry

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

var spl_SplFileInfo_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo___construct, ArginfoInfoConstruct),
	types.MakeZendFunctionEntryEx("getPath", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getPath, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getFilename", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getFilename, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getExtension", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getExtension, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getBasename", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getBasename, arginfo_optinalSuffix),
	types.MakeZendFunctionEntryEx("getPathname", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getPathname, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getPerms", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getPerms, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getInode", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getInode, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getSize", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getSize, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getOwner", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getOwner, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getGroup", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getGroup, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getATime", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getATime, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getMTime", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getMTime, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getCTime", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getCTime, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getType", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getType, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isWritable", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isWritable, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isReadable", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isReadable, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isExecutable", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isExecutable, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isFile", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isFile, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isDir", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isDir, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("isLink", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_isLink, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getLinkTarget", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getLinkTarget, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getRealPath", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getRealPath, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getFileInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getFileInfo, arginfo_info_optinalFileClass),
	types.MakeZendFunctionEntryEx("getPathInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getPathInfo, arginfo_info_optinalFileClass),
	types.MakeZendFunctionEntryEx("openFile", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_openFile, arginfo_info_openFile),
	types.MakeZendFunctionEntryEx("setFileClass", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_setFileClass, arginfo_info_optinalFileClass),
	types.MakeZendFunctionEntryEx("setInfoClass", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_setInfoClass, arginfo_info_optinalFileClass),
	types.MakeZendFunctionEntryEx("__debugInfo", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo___debugInfo, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("_bad_state_ex", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_FINAL, zim_spl_SplFileInfo__bad_state_ex, nil),
	types.MakeZendFunctionEntryEx("__toString", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileInfo_getPathname, ArginfoSplfileinfoVoid),
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

var spl_DirectoryIterator_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator___construct, ArginfoDirConstruct),
	types.MakeZendFunctionEntryEx("getFilename", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_getFilename, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getExtension", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_getExtension, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getBasename", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_getBasename, arginfo_optinalSuffix),
	types.MakeZendFunctionEntryEx("isDot", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_isDot, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_rewind, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_valid, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_key, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_current, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_next, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("seek", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_seek, ArginfoDirItSeek),
	types.MakeZendFunctionEntryEx("__toString", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_getFilename, ArginfoSplfileinfoVoid),
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
var spl_FilesystemIterator_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator___construct, ArginfoRDirConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator_rewind, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_DirectoryIterator_next, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator_key, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator_current, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator_getFlags, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_FilesystemIterator_setFlags, arginfo_r_dir_setFlags),
}
var spl_RecursiveDirectoryIterator_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveDirectoryIterator___construct, ArginfoRDirConstruct),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveDirectoryIterator_hasChildren, arginfo_r_dir_hasChildren),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveDirectoryIterator_getChildren, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getSubPath", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveDirectoryIterator_getSubPath, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getSubPathname", zend.ZEND_ACC_PUBLIC, zim_spl_RecursiveDirectoryIterator_getSubPathname, ArginfoSplfileinfoVoid),
}
var spl_GlobIterator_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_GlobIterator___construct, ArginfoRDirConstruct),
	types.MakeZendFunctionEntryEx("count", zend.ZEND_ACC_PUBLIC, zim_spl_GlobIterator_count, ArginfoSplfileinfoVoid),
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
	zend.MakeArgInfo("wouldblock", ArgInfoByRef(1)),
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
	zend.MakeArgInfo("vars", ArgInfoByRef(1), ArgInfoVariadic()),
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
var spl_SplFileObject_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject___construct, ArginfoFileObjectConstruct),
	types.MakeZendFunctionEntryEx("rewind", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_rewind, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("eof", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_eof, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("valid", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_valid, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("fgets", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("fgetcsv", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgetcsv, ArginfoFileObjectFgetcsv),
	types.MakeZendFunctionEntryEx("fputcsv", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fputcsv, ArginfoFileObjectFputcsv),
	types.MakeZendFunctionEntryEx("setCsvControl", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_setCsvControl, ArginfoFileObjectFgetcsv),
	types.MakeZendFunctionEntryEx("getCsvControl", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_getCsvControl, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("flock", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_flock, ArginfoFileObjectFlock),
	types.MakeZendFunctionEntryEx("fflush", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fflush, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("ftell", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_ftell, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("fseek", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fseek, ArginfoFileObjectFseek),
	types.MakeZendFunctionEntryEx("fgetc", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgetc, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("fpassthru", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fpassthru, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("fgetss", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgetss, ArginfoFileObjectFgetss),
	types.MakeZendFunctionEntryEx("fscanf", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fscanf, ArginfoFileObjectFscanf),
	types.MakeZendFunctionEntryEx("fwrite", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fwrite, ArginfoFileObjectFwrite),
	types.MakeZendFunctionEntryEx("fread", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fread, ArginfoFileObjectFread),
	types.MakeZendFunctionEntryEx("fstat", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fstat, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("ftruncate", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_ftruncate, ArginfoFileObjectFtruncate),
	types.MakeZendFunctionEntryEx("current", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_current, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("key", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_key, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("next", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_next, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("setFlags", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_setFlags, arginfo_file_object_setFlags),
	types.MakeZendFunctionEntryEx("getFlags", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_getFlags, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("setMaxLineLen", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_setMaxLineLen, arginfo_file_object_setMaxLineLen),
	types.MakeZendFunctionEntryEx("getMaxLineLen", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_getMaxLineLen, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("hasChildren", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_hasChildren, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("getChildren", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_getChildren, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("seek", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_seek, ArginfoFileObjectSeek),
	types.MakeZendFunctionEntryEx("getCurrentLine", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid),
	types.MakeZendFunctionEntryEx("__toString", zend.ZEND_ACC_PUBLIC, zim_spl_SplFileObject_fgets, ArginfoSplfileinfoVoid),
}
var ArginfoTempFileObjectConstruct []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(0),
	zend.MakeArgInfo("max_memory"),
}
var spl_SplTempFileObject_functions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", zend.ZEND_ACC_PUBLIC, zim_spl_SplTempFileObject___construct, ArginfoTempFileObjectConstruct),
}
