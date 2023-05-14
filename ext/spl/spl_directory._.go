package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

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

/* declare the class handlers */

var SplFilesystemObjectHandlers types.ObjectHandlers

/* includes handler to validate object state when retrieving methods */

var SplFilesystemObjectCheckHandlers types.ObjectHandlers

/* decalre the class entry */

const DIT_CTOR_FLAGS = 0x1
const DIT_CTOR_GLOB = 0x2

/* iterator handler table */

var SplFilesystemDirItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemDirItDtor, SplFilesystemDirItValid, SplFilesystemDirItCurrentData, SplFilesystemDirItCurrentKey, SplFilesystemDirItMoveForward, SplFilesystemDirItRewind, nil)
var SplFilesystemTreeItFuncs zend.ZendObjectIteratorFuncs = zend.MakeZendObjectIteratorFuncs(SplFilesystemTreeItDtor, SplFilesystemDirItValid, SplFilesystemTreeItCurrentData, SplFilesystemTreeItCurrentKey, SplFilesystemTreeItMoveForward, SplFilesystemTreeItRewind, nil)

var spl_SplFileInfo_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_SplFileInfo___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("file_name"),
	}),
	types.MakeZendFunctionEntryEx("getPath", types.AccPublic, zim_spl_SplFileInfo_getPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFilename", types.AccPublic, zim_spl_SplFileInfo_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getExtension", types.AccPublic, zim_spl_SplFileInfo_getExtension, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getBasename", types.AccPublic, zim_spl_SplFileInfo_getBasename, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("suffix"),
	}),
	types.MakeZendFunctionEntryEx("getPathname", types.AccPublic, zim_spl_SplFileInfo_getPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getPerms", types.AccPublic, zim_spl_SplFileInfo_getPerms, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getInode", types.AccPublic, zim_spl_SplFileInfo_getInode, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getSize", types.AccPublic, zim_spl_SplFileInfo_getSize, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getOwner", types.AccPublic, zim_spl_SplFileInfo_getOwner, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getGroup", types.AccPublic, zim_spl_SplFileInfo_getGroup, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getATime", types.AccPublic, zim_spl_SplFileInfo_getATime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getMTime", types.AccPublic, zim_spl_SplFileInfo_getMTime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getCTime", types.AccPublic, zim_spl_SplFileInfo_getCTime, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getType", types.AccPublic, zim_spl_SplFileInfo_getType, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isWritable", types.AccPublic, zim_spl_SplFileInfo_isWritable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isReadable", types.AccPublic, zim_spl_SplFileInfo_isReadable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isExecutable", types.AccPublic, zim_spl_SplFileInfo_isExecutable, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isFile", types.AccPublic, zim_spl_SplFileInfo_isFile, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isDir", types.AccPublic, zim_spl_SplFileInfo_isDir, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("isLink", types.AccPublic, zim_spl_SplFileInfo_isLink, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getLinkTarget", types.AccPublic, zim_spl_SplFileInfo_getLinkTarget, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getRealPath", types.AccPublic, zim_spl_SplFileInfo_getRealPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFileInfo", types.AccPublic, zim_spl_SplFileInfo_getFileInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types.MakeZendFunctionEntryEx("getPathInfo", types.AccPublic, zim_spl_SplFileInfo_getPathInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types.MakeZendFunctionEntryEx("openFile", types.AccPublic, zim_spl_SplFileInfo_openFile, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("open_mode"),
		zend.MakeArgName("use_include_path"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("setFileClass", types.AccPublic, zim_spl_SplFileInfo_setFileClass, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types.MakeZendFunctionEntryEx("setInfoClass", types.AccPublic, zim_spl_SplFileInfo_setInfoClass, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("class_name"),
	}),
	types.MakeZendFunctionEntryEx("__debugInfo", types.AccPublic, zim_spl_SplFileInfo___debugInfo, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("_bad_state_ex", types.AccPublic|types.AccFinal, zim_spl_SplFileInfo__bad_state_ex, nil),
	types.MakeZendFunctionEntryEx("__toString", types.AccPublic, zim_spl_SplFileInfo_getPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_DirectoryIterator_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_DirectoryIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("path"),
	}),
	types.MakeZendFunctionEntryEx("getFilename", types.AccPublic, zim_spl_DirectoryIterator_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getExtension", types.AccPublic, zim_spl_DirectoryIterator_getExtension, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getBasename", types.AccPublic, zim_spl_DirectoryIterator_getBasename, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("suffix"),
	}),
	types.MakeZendFunctionEntryEx("isDot", types.AccPublic, zim_spl_DirectoryIterator_isDot, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_DirectoryIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic, zim_spl_DirectoryIterator_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_DirectoryIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_DirectoryIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_DirectoryIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("seek", types.AccPublic, zim_spl_DirectoryIterator_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("position"),
	}),
	types.MakeZendFunctionEntryEx("__toString", types.AccPublic, zim_spl_DirectoryIterator_getFilename, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_FilesystemIterator_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_FilesystemIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_FilesystemIterator_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_DirectoryIterator_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_FilesystemIterator_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_FilesystemIterator_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getFlags", types.AccPublic, zim_spl_FilesystemIterator_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", types.AccPublic, zim_spl_FilesystemIterator_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("flags"),
	}),
}
var spl_RecursiveDirectoryIterator_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_RecursiveDirectoryIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("hasChildren", types.AccPublic, zim_spl_RecursiveDirectoryIterator_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("allow_links"),
	}),
	types.MakeZendFunctionEntryEx("getChildren", types.AccPublic, zim_spl_RecursiveDirectoryIterator_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getSubPath", types.AccPublic, zim_spl_RecursiveDirectoryIterator_getSubPath, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getSubPathname", types.AccPublic, zim_spl_RecursiveDirectoryIterator_getSubPathname, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_GlobIterator_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_GlobIterator___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("path"),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("count", types.AccPublic, zim_spl_GlobIterator_count, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}

var spl_SplFileObject_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_SplFileObject___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("file_name"),
		zend.MakeArgName("open_mode"),
		zend.MakeArgName("use_include_path"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("rewind", types.AccPublic, zim_spl_SplFileObject_rewind, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("eof", types.AccPublic, zim_spl_SplFileObject_eof, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("valid", types.AccPublic, zim_spl_SplFileObject_valid, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fgets", types.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fgetcsv", types.AccPublic, zim_spl_SplFileObject_fgetcsv, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types.MakeZendFunctionEntryEx("fputcsv", types.AccPublic, zim_spl_SplFileObject_fputcsv, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("fields"),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types.MakeZendFunctionEntryEx("setCsvControl", types.AccPublic, zim_spl_SplFileObject_setCsvControl, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("delimiter"),
		zend.MakeArgName("enclosure"),
		zend.MakeArgName("escape"),
	}),
	types.MakeZendFunctionEntryEx("getCsvControl", types.AccPublic, zim_spl_SplFileObject_getCsvControl, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("flock", types.AccPublic, zim_spl_SplFileObject_flock, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("operation"),
		zend.MakeArgByRef("wouldblock"),
	}),
	types.MakeZendFunctionEntryEx("fflush", types.AccPublic, zim_spl_SplFileObject_fflush, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("ftell", types.AccPublic, zim_spl_SplFileObject_ftell, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fseek", types.AccPublic, zim_spl_SplFileObject_fseek, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("pos"),
		zend.MakeArgName("whence"),
	}),
	types.MakeZendFunctionEntryEx("fgetc", types.AccPublic, zim_spl_SplFileObject_fgetc, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fpassthru", types.AccPublic, zim_spl_SplFileObject_fpassthru, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("fgetss", types.AccPublic, zim_spl_SplFileObject_fgetss, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("allowable_tags"),
	}),
	types.MakeZendFunctionEntryEx("fscanf", types.AccPublic, zim_spl_SplFileObject_fscanf, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("format"),
		zend.MakeArgInfo("vars", zend.ArgInfoByRef(1), zend.ArgInfoVariadic()),
	}),
	types.MakeZendFunctionEntryEx("fwrite", types.AccPublic, zim_spl_SplFileObject_fwrite, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("str"),
		zend.MakeArgName("length"),
	}),
	types.MakeZendFunctionEntryEx("fread", types.AccPublic, zim_spl_SplFileObject_fread, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("length"),
	}),
	types.MakeZendFunctionEntryEx("fstat", types.AccPublic, zim_spl_SplFileObject_fstat, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("ftruncate", types.AccPublic, zim_spl_SplFileObject_ftruncate, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("size"),
	}),
	types.MakeZendFunctionEntryEx("current", types.AccPublic, zim_spl_SplFileObject_current, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("key", types.AccPublic, zim_spl_SplFileObject_key, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("next", types.AccPublic, zim_spl_SplFileObject_next, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setFlags", types.AccPublic, zim_spl_SplFileObject_setFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("flags"),
	}),
	types.MakeZendFunctionEntryEx("getFlags", types.AccPublic, zim_spl_SplFileObject_getFlags, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("setMaxLineLen", types.AccPublic, zim_spl_SplFileObject_setMaxLineLen, []zend.ArgInfo{zend.MakeReturnArgInfo(-1),
		zend.MakeArgName("max_len"),
	}),
	types.MakeZendFunctionEntryEx("getMaxLineLen", types.AccPublic, zim_spl_SplFileObject_getMaxLineLen, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("hasChildren", types.AccPublic, zim_spl_SplFileObject_hasChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("getChildren", types.AccPublic, zim_spl_SplFileObject_getChildren, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("seek", types.AccPublic, zim_spl_SplFileObject_seek, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("line_pos"),
	}),
	types.MakeZendFunctionEntryEx("getCurrentLine", types.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
	types.MakeZendFunctionEntryEx("__toString", types.AccPublic, zim_spl_SplFileObject_fgets, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var spl_SplTempFileObject_functions []types.FunctionEntry = []types.FunctionEntry{
	types.MakeZendFunctionEntryEx("__construct", types.AccPublic, zim_spl_SplTempFileObject___construct, []zend.ArgInfo{zend.MakeReturnArgInfo(0),
		zend.MakeArgName("max_memory"),
	}),
}
