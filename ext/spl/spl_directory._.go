package spl

import (
	types2 "github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var spl_ce_SplFileInfo *types2.ClassEntry
var spl_ce_DirectoryIterator *types2.ClassEntry
var spl_ce_FilesystemIterator *types2.ClassEntry
var spl_ce_RecursiveDirectoryIterator *types2.ClassEntry
var spl_ce_GlobIterator *types2.ClassEntry
var spl_ce_SplFileObject *types2.ClassEntry
var spl_ce_SplTempFileObject *types2.ClassEntry

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

/* declare the class handlers */

var SplFilesystemObjectHandlers zend.ZendObjectHandlers

/* includes handler to validate object state when retrieving methods */

var SplFilesystemObjectCheckHandlers zend.ZendObjectHandlers

/* decalre the class entry */

const DIT_CTOR_FLAGS = 0x1
const DIT_CTOR_GLOB = 0x2

/* iterator handler table */

var SplFilesystemDirItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemDirItDtor, SplFilesystemDirItValid, SplFilesystemDirItCurrentData, SplFilesystemDirItCurrentKey, SplFilesystemDirItMoveForward, SplFilesystemDirItRewind, nil)
var SplFilesystemTreeItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemTreeItDtor, SplFilesystemDirItValid, SplFilesystemTreeItCurrentData, SplFilesystemTreeItCurrentKey, SplFilesystemTreeItMoveForward, SplFilesystemTreeItRewind, nil)

var spl_SplFileInfo_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_SplFileInfo___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("file_name"),
	}),
	types2.MakeZendFunctionEntryEx("getPath", zend.AccPublic, zim_spl_SplFileInfo_getPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getFilename", zend.AccPublic, zim_spl_SplFileInfo_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getExtension", zend.AccPublic, zim_spl_SplFileInfo_getExtension, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getBasename", zend.AccPublic, zim_spl_SplFileInfo_getBasename, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("suffix"),
	}),
	types2.MakeZendFunctionEntryEx("getPathname", zend.AccPublic, zim_spl_SplFileInfo_getPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getPerms", zend.AccPublic, zim_spl_SplFileInfo_getPerms, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getInode", zend.AccPublic, zim_spl_SplFileInfo_getInode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getSize", zend.AccPublic, zim_spl_SplFileInfo_getSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getOwner", zend.AccPublic, zim_spl_SplFileInfo_getOwner, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getGroup", zend.AccPublic, zim_spl_SplFileInfo_getGroup, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getATime", zend.AccPublic, zim_spl_SplFileInfo_getATime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getMTime", zend.AccPublic, zim_spl_SplFileInfo_getMTime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getCTime", zend.AccPublic, zim_spl_SplFileInfo_getCTime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getType", zend.AccPublic, zim_spl_SplFileInfo_getType, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isWritable", zend.AccPublic, zim_spl_SplFileInfo_isWritable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isReadable", zend.AccPublic, zim_spl_SplFileInfo_isReadable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isExecutable", zend.AccPublic, zim_spl_SplFileInfo_isExecutable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isFile", zend.AccPublic, zim_spl_SplFileInfo_isFile, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isDir", zend.AccPublic, zim_spl_SplFileInfo_isDir, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("isLink", zend.AccPublic, zim_spl_SplFileInfo_isLink, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getLinkTarget", zend.AccPublic, zim_spl_SplFileInfo_getLinkTarget, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getRealPath", zend.AccPublic, zim_spl_SplFileInfo_getRealPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getFileInfo", zend.AccPublic, zim_spl_SplFileInfo_getFileInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types2.MakeZendFunctionEntryEx("getPathInfo", zend.AccPublic, zim_spl_SplFileInfo_getPathInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types2.MakeZendFunctionEntryEx("openFile", zend.AccPublic, zim_spl_SplFileInfo_openFile, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("open_mode"),
		zend.MakeArgName("use_include_path"),
		zend.MakeArgName("context"),
	}),
	types2.MakeZendFunctionEntryEx("setFileClass", zend.AccPublic, zim_spl_SplFileInfo_setFileClass, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types2.MakeZendFunctionEntryEx("setInfoClass", zend.AccPublic, zim_spl_SplFileInfo_setInfoClass, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types2.MakeZendFunctionEntryEx("__debugInfo", zend.AccPublic, zim_spl_SplFileInfo___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("_bad_state_ex", zend.AccPublic|zend.AccFinal, zim_spl_SplFileInfo__bad_state_ex, nil),
	types2.MakeZendFunctionEntryEx("__toString", zend.AccPublic, zim_spl_SplFileInfo_getPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_DirectoryIterator_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_DirectoryIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("path"),
	}),
	types2.MakeZendFunctionEntryEx("getFilename", zend.AccPublic, zim_spl_DirectoryIterator_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getExtension", zend.AccPublic, zim_spl_DirectoryIterator_getExtension, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getBasename", zend.AccPublic, zim_spl_DirectoryIterator_getBasename, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("suffix"),
	}),
	types2.MakeZendFunctionEntryEx("isDot", zend.AccPublic, zim_spl_DirectoryIterator_isDot, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_DirectoryIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_DirectoryIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_DirectoryIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_DirectoryIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_DirectoryIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("seek", zend.AccPublic, zim_spl_DirectoryIterator_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("position"),
	}),
	types2.MakeZendFunctionEntryEx("__toString", zend.AccPublic, zim_spl_DirectoryIterator_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_FilesystemIterator_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_FilesystemIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_FilesystemIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_DirectoryIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_FilesystemIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_FilesystemIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_FilesystemIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_FilesystemIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("flags"),
	}),
}
var spl_RecursiveDirectoryIterator_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_RecursiveDirectoryIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_RecursiveDirectoryIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("allow_links"),
	}),
	types2.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_RecursiveDirectoryIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getSubPath", zend.AccPublic, zim_spl_RecursiveDirectoryIterator_getSubPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getSubPathname", zend.AccPublic, zim_spl_RecursiveDirectoryIterator_getSubPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_GlobIterator_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_GlobIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("count", zend.AccPublic, zim_spl_GlobIterator_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_SplFileObject_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_SplFileObject___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("file_name"),
		zend.MakeArgName("open_mode"),
		zend.MakeArgName("use_include_path"),
		zend.MakeArgName("context"),
	}),
	types2.MakeZendFunctionEntryEx("rewind", zend.AccPublic, zim_spl_SplFileObject_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("eof", zend.AccPublic, zim_spl_SplFileObject_eof, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("valid", zend.AccPublic, zim_spl_SplFileObject_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fgets", zend.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fgetcsv", zend.AccPublic, zim_spl_SplFileObject_fgetcsv, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types2.MakeZendFunctionEntryEx("fputcsv", zend.AccPublic, zim_spl_SplFileObject_fputcsv, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("fields"),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types2.MakeZendFunctionEntryEx("setCsvControl", zend.AccPublic, zim_spl_SplFileObject_setCsvControl, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types2.MakeZendFunctionEntryEx("getCsvControl", zend.AccPublic, zim_spl_SplFileObject_getCsvControl, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("flock", zend.AccPublic, zim_spl_SplFileObject_flock, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("operation"),
		zend.MakeArgByRef("wouldblock"),
	}),
	types2.MakeZendFunctionEntryEx("fflush", zend.AccPublic, zim_spl_SplFileObject_fflush, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("ftell", zend.AccPublic, zim_spl_SplFileObject_ftell, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fseek", zend.AccPublic, zim_spl_SplFileObject_fseek, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("pos"),
		zend.MakeArgName("whence"),
	}),
	types2.MakeZendFunctionEntryEx("fgetc", zend.AccPublic, zim_spl_SplFileObject_fgetc, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fpassthru", zend.AccPublic, zim_spl_SplFileObject_fpassthru, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("fgetss", zend.AccPublic, zim_spl_SplFileObject_fgetss, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("allowable_tags"),
	}),
	types2.MakeZendFunctionEntryEx("fscanf", zend.AccPublic, zim_spl_SplFileObject_fscanf, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("format"),
		zend.MakeArgInfo("vars", zend.ArgInfoByRef(1), zend.ArgInfoVariadic()),
	}),
	types2.MakeZendFunctionEntryEx("fwrite", zend.AccPublic, zim_spl_SplFileObject_fwrite, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("str"),
		zend.MakeArgName("length"),
	}),
	types2.MakeZendFunctionEntryEx("fread", zend.AccPublic, zim_spl_SplFileObject_fread, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("length"),
	}),
	types2.MakeZendFunctionEntryEx("fstat", zend.AccPublic, zim_spl_SplFileObject_fstat, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("ftruncate", zend.AccPublic, zim_spl_SplFileObject_ftruncate, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("size"),
	}),
	types2.MakeZendFunctionEntryEx("current", zend.AccPublic, zim_spl_SplFileObject_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("key", zend.AccPublic, zim_spl_SplFileObject_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("next", zend.AccPublic, zim_spl_SplFileObject_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setFlags", zend.AccPublic, zim_spl_SplFileObject_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types2.MakeZendFunctionEntryEx("getFlags", zend.AccPublic, zim_spl_SplFileObject_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("setMaxLineLen", zend.AccPublic, zim_spl_SplFileObject_setMaxLineLen, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("max_len"),
	}),
	types2.MakeZendFunctionEntryEx("getMaxLineLen", zend.AccPublic, zim_spl_SplFileObject_getMaxLineLen, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("hasChildren", zend.AccPublic, zim_spl_SplFileObject_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("getChildren", zend.AccPublic, zim_spl_SplFileObject_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("seek", zend.AccPublic, zim_spl_SplFileObject_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("line_pos"),
	}),
	types2.MakeZendFunctionEntryEx("getCurrentLine", zend.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types2.MakeZendFunctionEntryEx("__toString", zend.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_SplTempFileObject_functions []types2.FunctionEntry = []types2.FunctionEntry{
	types2.MakeZendFunctionEntryEx("__construct", zend.AccPublic, zim_spl_SplTempFileObject___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("max_memory"),
	}),
}
