package zend

const MAXPATHLEN = 256
const DEFAULT_SLASH = '/'
const DEFAULT_DIR_SEPARATOR = ':'
const PhpSysStat = stat
const PhpSysLstat = lstat
const PhpSysFstat = fstat
const PhpSysSymlink = symlink
const PhpSysLink = link

type VerifyPathFunc func(*CwdState) int

/* One of the following constants must be used as the last argument
   in virtual_file_ex() call. */

const CWD_EXPAND = 0
const CWD_FILEPATH = 1
const CWD_REALPATH = 2
const REALPATH_CACHE_TTL ZendLong = 2 * 60
const REALPATH_CACHE_SIZE = 0

var CwdGlobals VirtualCwdGlobals

/* The actual macros to be used in programs using TSRM
 * If the program defines VIRTUAL_DIR it will use the
 * virtual_* functions
 */

/* rename on windows will fail if newname already exists.
   MoveFileEx has to be used */

/* Global stat declarations */

const _S_IFDIR = S_IFDIR
const _S_IFREG = S_IFREG
const _IFLNK = 0120000
const S_IFLNK = _IFLNK
const S_IXROOT = S_IXUSR | S_IXGRP | S_IXOTH

/* XXX should be _S_IFIFO? */

const _IFIFO = 010000
const S_IFIFO = _IFIFO
const _IFBLK = 060000
const S_IFBLK = _IFBLK

var MainCwdState CwdState

const LINK_MAX = 32
