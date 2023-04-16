package zend

const MAXPATHLEN = 256
const DEFAULT_SLASH = '/'
const DEFAULT_DIR_SEPARATOR = ':'
const PhpSysStat = stat
const PhpSysLstat = lstat
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

const _IFLNK = 0120000
const S_IFLNK = _IFLNK
const S_IXROOT = S_IXUSR | S_IXGRP | S_IXOTH
const _IFIFO = 010000
const S_IFIFO = _IFIFO
const _IFBLK = 060000
const S_IFBLK = _IFBLK

var MainCwdState CwdState

const LINK_MAX = 32
