// <<generate>>

package zend

import (
	g "sik/runtime/grammar"
)

// Source: <Zend/zend_virtual_cwd.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Sascha Schumann <sascha@schumann.cx>                        |
   |          Pierre Joye <pierre@php.net>                                |
   +----------------------------------------------------------------------+
*/

// #define VIRTUAL_CWD_H

// failed # include "TSRM.h"

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < ctype . h >

// # include < utime . h >

// # include < stdarg . h >

// # include < limits . h >

// # include < sys / param . h >

// #define MAXPATHLEN       256

// # include < unistd . h >

// # include < dirent . h >

// #define DEFAULT_SLASH       '/'

// #define DEFAULT_DIR_SEPARATOR       ':'

// #define IS_SLASH(c) ( ( c ) == '/' )

// #define IS_SLASH_P(c) ( * ( c ) == '/' )

// #define COPY_WHEN_ABSOLUTE(path) 0

// #define IS_ABSOLUTE_PATH(path,len) ( IS_SLASH ( path [ 0 ] ) )

// #define CWD_API

// #define php_sys_stat       stat

// #define php_sys_lstat       lstat

// #define php_sys_fstat       fstat

// #define php_sys_readlink(link,target,target_len) readlink ( link , target , target_len )

// #define php_sys_symlink       symlink

// #define php_sys_link       link

// @type CwdState struct

type VerifyPathFunc func(*CwdState) int

/* One of the following constants must be used as the last argument
   in virtual_file_ex() call. */

// #define CWD_EXPAND       0

// #define CWD_FILEPATH       1

// #define CWD_REALPATH       2

// #define REALPATH_CACHE_TTL       ( 2 * 60 )

// #define REALPATH_CACHE_SIZE       0

// @type RealpathCacheBucket struct

// @type VirtualCwdGlobals struct

var CwdGlobals VirtualCwdGlobals

// #define CWDG(v) ( cwd_globals . v )

/* The actual macros to be used in programs using TSRM
 * If the program defines VIRTUAL_DIR it will use the
 * virtual_* functions
 */

// #define VCWD_CREAT(path,mode) creat ( path , mode )

/* rename on windows will fail if newname already exists.
   MoveFileEx has to be used */

// #define VCWD_FOPEN(path,mode) fopen ( path , mode )

// #define VCWD_OPEN(path,flags) open ( path , flags )

// #define VCWD_OPEN_MODE(path,flags,mode) open ( path , flags , mode )

// #define VCWD_RENAME(oldname,newname) rename ( oldname , newname )

// #define VCWD_MKDIR(pathname,mode) mkdir ( pathname , mode )

// #define VCWD_RMDIR(pathname) rmdir ( pathname )

// #define VCWD_UNLINK(path) unlink ( path )

// #define VCWD_CHDIR(path) chdir ( path )

// #define VCWD_ACCESS(pathname,mode) access ( pathname , mode )

// #define VCWD_GETCWD(buff,size) getcwd ( buff , size )

// #define VCWD_CHMOD(path,mode) chmod ( path , mode )

// #define VCWD_CHDIR_FILE(path) virtual_chdir_file ( path , chdir )

// #define VCWD_GETWD(buf) getwd ( buf )

// #define VCWD_STAT(path,buff) php_sys_stat ( path , buff )

// #define VCWD_LSTAT(path,buff) lstat ( path , buff )

// #define VCWD_OPENDIR(pathname) opendir ( pathname )

// #define VCWD_POPEN(command,type) popen ( command , type )

// #define VCWD_REALPATH(path,real_path) tsrm_realpath ( path , real_path )

// #define VCWD_UTIME(path,time) utime ( path , time )

// #define VCWD_CHOWN(path,owner,group) chown ( path , owner , group )

// #define VCWD_LCHOWN(path,owner,group) lchown ( path , owner , group )

/* Global stat declarations */

// #define _S_IFDIR       S_IFDIR

// #define _S_IFREG       S_IFREG

// #define _IFLNK       0120000

// #define S_IFLNK       _IFLNK

// #define S_ISDIR(mode) ( ( ( mode ) & S_IFMT ) == S_IFDIR )

// #define S_ISREG(mode) ( ( ( mode ) & S_IFMT ) == S_IFREG )

// #define S_ISLNK(mode) ( ( ( mode ) & S_IFMT ) == S_IFLNK )

// #define S_IXROOT       ( S_IXUSR | S_IXGRP | S_IXOTH )

/* XXX should be _S_IFIFO? */

// #define _IFIFO       0010000

// #define S_IFIFO       _IFIFO

// #define _IFBLK       0060000

// #define S_IFBLK       _IFBLK

// Source: <Zend/zend_virtual_cwd.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Sascha Schumann <sascha@schumann.cx>                        |
   |          Pierre Joye <pierre@php.net>                                |
   +----------------------------------------------------------------------+
*/

// # include < sys / types . h >

// # include < sys / stat . h >

// # include < string . h >

// # include < stdio . h >

// # include < limits . h >

// # include < errno . h >

// # include < stdlib . h >

// # include < fcntl . h >

// # include < time . h >

// # include "zend.h"

// # include "zend_virtual_cwd.h"

// #define VIRTUAL_CWD_DEBUG       0

// failed # include "TSRM.h"

/* Only need mutex for popen() in Windows because it doesn't chdir() on UNIX */

var MainCwdState CwdState

// # include < unistd . h >

// #define CWD_STATE_COPY(d,s) ( d ) -> cwd_length = ( s ) -> cwd_length ; ( d ) -> cwd = ( char * ) emalloc ( ( s ) -> cwd_length + 1 ) ; memcpy ( ( d ) -> cwd , ( s ) -> cwd , ( s ) -> cwd_length + 1 ) ;

// #define CWD_STATE_FREE(s) efree ( ( s ) -> cwd ) ; ( s ) -> cwd_length = 0 ;

// #define CWD_STATE_FREE_ERR(state) CWD_STATE_FREE ( state )

func PhpIsDirOk(state *CwdState) int {
	var buf ZendStatT
	if stat(state.GetCwd(), &buf) == 0 && (buf.st_mode&S_IFMT) == S_IFDIR {
		return 0
	}
	return 1
}

/* }}} */

func PhpIsFileOk(state *CwdState) int {
	var buf ZendStatT
	if stat(state.GetCwd(), &buf) == 0 && (buf.st_mode&S_IFMT) == S_IFREG {
		return 0
	}
	return 1
}

/* }}} */

func CwdGlobalsCtor(cwd_g *VirtualCwdGlobals) {
	&cwd_g.cwd.cwd_length = &MainCwdState.GetCwdLength()
	&cwd_g.cwd.cwd = (*byte)(_emalloc(&MainCwdState.GetCwdLength() + 1))
	memcpy(&cwd_g.cwd.cwd, &MainCwdState.GetCwd(), &MainCwdState.GetCwdLength()+1)
	cwd_g.SetRealpathCacheSize(0)
	cwd_g.SetRealpathCacheSizeLimit(0)
	cwd_g.SetRealpathCacheTtl(2 * 60)
	memset(cwd_g.GetRealpathCache(), 0, g.SizeOf("cwd_g -> realpath_cache"))
}

/* }}} */

func RealpathCacheCleanHelper(max_entries uint32, cache **RealpathCacheBucket, cache_size *ZendLong) {
	var i uint32
	for i = 0; i < max_entries; i++ {
		var p *RealpathCacheBucket = cache[i]
		for p != nil {
			var r *RealpathCacheBucket = p
			p = p.GetNext()
			Free(r)
		}
		cache[i] = nil
	}
	*cache_size = 0
}
func CwdGlobalsDtor(cwd_g *VirtualCwdGlobals) {
	RealpathCacheCleanHelper(g.SizeOf("cwd_g -> realpath_cache")/g.SizeOf("cwd_g -> realpath_cache [ 0 ]"), cwd_g.GetRealpathCache(), &cwd_g.realpath_cache_size)
}

/* }}} */

func VirtualCwdMainCwdInit(reinit uint8) {
	var cwd []byte
	var result *byte
	if reinit != 0 {
		Free(MainCwdState.GetCwd())
	}
	result = getcwd(cwd, g.SizeOf("cwd"))
	if result == nil {
		cwd[0] = '0'
	}
	MainCwdState.SetCwdLength(strlen(cwd))
	MainCwdState.SetCwd(strdup(cwd))
}

/* }}} */

func VirtualCwdStartup() {
	VirtualCwdMainCwdInit(0)
	CwdGlobalsCtor(&CwdGlobals)
}

/* }}} */

func VirtualCwdShutdown() {
	CwdGlobalsDtor(&CwdGlobals)
	Free(MainCwdState.GetCwd())
}

/* }}} */

func VirtualCwdActivate() int {
	if CwdGlobals.GetCwd().GetCwd() == nil {
		&(CwdGlobals.GetCwd()).SetCwdLength(&MainCwdState.GetCwdLength())
		&(CwdGlobals.GetCwd()).SetCwd((*byte)(_emalloc(&MainCwdState.GetCwdLength() + 1)))
		memcpy(&(CwdGlobals.GetCwd()).GetCwd(), &MainCwdState.GetCwd(), &MainCwdState.GetCwdLength()+1)
	}
	return 0
}

/* }}} */

func VirtualCwdDeactivate() int {
	if CwdGlobals.GetCwd().GetCwd() != nil {
		_efree(&(CwdGlobals.GetCwd()).GetCwd())
		&(CwdGlobals.GetCwd()).SetCwdLength(0)
		CwdGlobals.GetCwd().SetCwd(nil)
	}
	return 0
}

/* }}} */

func VirtualGetcwdEx(length *int) *byte {
	var state *CwdState
	state = &(CwdGlobals.GetCwd())
	if state.GetCwdLength() == 0 {
		var retval *byte
		*length = 1
		retval = (*byte)(_emalloc(2))
		retval[0] = '/'
		retval[1] = '0'
		return retval
	}
	if state.GetCwd() == nil {
		*length = 0
		return nil
	}
	*length = state.GetCwdLength()
	return _estrdup(state.GetCwd())
}

/* }}} */

func VirtualGetcwd(buf *byte, size int) *byte {
	var length int
	var cwd *byte
	cwd = VirtualGetcwdEx(&length)
	if buf == nil {
		return cwd
	}
	if length > size-1 {
		_efree(cwd)
		errno = ERANGE
		return nil
	}
	if cwd == nil {
		return nil
	}
	memcpy(buf, cwd, length+1)
	_efree(cwd)
	return buf
}

/* }}} */

func RealpathCacheKey(path *byte, path_len int) ZendUlong {
	var h ZendUlong
	var e *byte = path + path_len
	for h = 2166136261; path < e; {
		h *= 16777619
		*path++
		h ^= (*path) - 1
	}
	return h
}

/* }}} */

func RealpathCacheClean() {
	RealpathCacheCleanHelper(g.SizeOf("CWDG ( realpath_cache )")/g.SizeOf("CWDG ( realpath_cache ) [ 0 ]"), CwdGlobals.GetRealpathCache(), &(CwdGlobals.GetRealpathCacheSize()))
}

/* }}} */

func RealpathCacheDel(path *byte, path_len int) {
	var key ZendUlong = RealpathCacheKey(path, path_len)
	var n ZendUlong = key % (g.SizeOf("CWDG ( realpath_cache )") / g.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
	var bucket **RealpathCacheBucket = &CwdGlobals.GetRealpathCache()[n]
	for (*bucket) != nil {
		if key == (*bucket).GetKey() && path_len == (*bucket).GetPathLen() && memcmp(path, (*bucket).GetPath(), path_len) == 0 {
			var r *RealpathCacheBucket = *bucket
			*bucket = (*bucket).GetNext()

			/* if the pointers match then only subtract the length of the path */

			if r.GetPath() == r.GetRealpath() {
				CwdGlobals.SetRealpathCacheSize(CwdGlobals.GetRealpathCacheSize() - g.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1)
			} else {
				CwdGlobals.SetRealpathCacheSize(CwdGlobals.GetRealpathCacheSize() - g.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1 + r.GetRealpathLen() + 1)
			}
			Free(r)
			return
		} else {
			bucket = &(*bucket).next
		}
	}
}

/* }}} */

func RealpathCacheAdd(path *byte, path_len int, realpath *byte, realpath_len int, is_dir int, t int64) {
	var size ZendLong = g.SizeOf("realpath_cache_bucket") + path_len + 1
	var same int = 1
	if realpath_len != path_len || memcmp(path, realpath, path_len) != 0 {
		size += realpath_len + 1
		same = 0
	}
	if CwdGlobals.GetRealpathCacheSize()+size <= CwdGlobals.GetRealpathCacheSizeLimit() {
		var bucket *RealpathCacheBucket = Malloc(size)
		var n ZendUlong
		if bucket == nil {
			return
		}
		bucket.SetKey(RealpathCacheKey(path, path_len))
		bucket.SetPath((*byte)(bucket + g.SizeOf("realpath_cache_bucket")))
		memcpy(bucket.GetPath(), path, path_len+1)
		bucket.SetPathLen(path_len)
		if same != 0 {
			bucket.SetRealpath(bucket.GetPath())
		} else {
			bucket.SetRealpath(bucket.GetPath() + (path_len + 1))
			memcpy(bucket.GetRealpath(), realpath, realpath_len+1)
		}
		bucket.SetRealpathLen(realpath_len)
		bucket.SetIsDir(is_dir > 0)
		bucket.SetExpires(t + CwdGlobals.GetRealpathCacheTtl())
		n = bucket.GetKey() % (g.SizeOf("CWDG ( realpath_cache )") / g.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
		bucket.SetNext(CwdGlobals.GetRealpathCache()[n])
		CwdGlobals.GetRealpathCache()[n] = bucket
		CwdGlobals.SetRealpathCacheSize(CwdGlobals.GetRealpathCacheSize() + size)
	}
}

/* }}} */

func RealpathCacheFind(path *byte, path_len int, t int64) *RealpathCacheBucket {
	var key ZendUlong = RealpathCacheKey(path, path_len)
	var n ZendUlong = key % (g.SizeOf("CWDG ( realpath_cache )") / g.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
	var bucket **RealpathCacheBucket = &CwdGlobals.GetRealpathCache()[n]
	for (*bucket) != nil {
		if CwdGlobals.GetRealpathCacheTtl() != 0 && (*bucket).GetExpires() < t {
			var r *RealpathCacheBucket = *bucket
			*bucket = (*bucket).GetNext()

			/* if the pointers match then only subtract the length of the path */

			if r.GetPath() == r.GetRealpath() {
				CwdGlobals.SetRealpathCacheSize(CwdGlobals.GetRealpathCacheSize() - g.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1)
			} else {
				CwdGlobals.SetRealpathCacheSize(CwdGlobals.GetRealpathCacheSize() - g.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1 + r.GetRealpathLen() + 1)
			}
			Free(r)
		} else if key == (*bucket).GetKey() && path_len == (*bucket).GetPathLen() && memcmp(path, (*bucket).GetPath(), path_len) == 0 {
			return *bucket
		} else {
			bucket = &(*bucket).next
		}
	}
	return nil
}

/* }}} */

func RealpathCacheLookup(path *byte, path_len int, t int64) *RealpathCacheBucket {
	return RealpathCacheFind(path, path_len, t)
}

/* }}} */

func RealpathCacheSize() ZendLong { return CwdGlobals.GetRealpathCacheSize() }
func RealpathCacheMaxBuckets() ZendLong {
	return g.SizeOf("CWDG ( realpath_cache )") / g.SizeOf("CWDG ( realpath_cache ) [ 0 ]")
}
func RealpathCacheGetBuckets() **RealpathCacheBucket { return CwdGlobals.GetRealpathCache() }

// #define LINK_MAX       32

func TsrmRealpathR(path *byte, start int, len_ int, ll *int, t *int64, use_realpath int, is_dir int, link_is_dir *int) int {
	var i int
	var j int
	var directory int = 0
	var save int
	var st ZendStatT
	var bucket *RealpathCacheBucket
	var tmp *byte
	for true {
		if len_ <= start {
			if link_is_dir != nil {
				*link_is_dir = 1
			}
			return start
		}
		i = len_
		for i > start && path[i-1] != '/' {
			i--
		}
		assert(i < 256)
		if i == len_ || i+1 == len_ && path[i] == '.' {

			/* remove double slashes and '.' */

			if i > 0 {
				len_ = i - 1
			} else {
				len_ = 0
			}
			is_dir = 1
			continue
		} else if i+2 == len_ && path[i] == '.' && path[i+1] == '.' {

			/* remove '..' and previous directory */

			is_dir = 1
			if link_is_dir != nil {
				*link_is_dir = 1
			}
			if i <= start+1 {
				if start != 0 {
					return start
				} else {
					return len_
				}
			}
			j = TsrmRealpathR(path, start, i-1, ll, t, use_realpath, 1, nil)
			if j > start && j != size_t-1 {
				j--
				assert(i < 256)
				for j > start && path[j] != '/' {
					j--
				}
				assert(i < 256)
				if start == 0 {

					/* leading '..' must not be removed in case of relative path */

					if j == 0 && path[0] == '.' && path[1] == '.' && path[2] == '/' {
						path[3] = '.'
						path[4] = '.'
						path[5] = '/'
						j = 5
					} else if j > 0 && path[j+1] == '.' && path[j+2] == '.' && path[j+3] == '/' {
						j += 4
						path[g.PostInc(&j)] = '.'
						path[g.PostInc(&j)] = '.'
						path[j] = '/'
					}

					/* leading '..' must not be removed in case of relative path */

				}
			} else if start == 0 && j == 0 {

				/* leading '..' must not be removed in case of relative path */

				path[0] = '.'
				path[1] = '.'
				path[2] = '/'
				j = 2
			}
			return j
		}
		path[len_] = 0
		save = use_realpath != 0
		if start != 0 && save != 0 && CwdGlobals.GetRealpathCacheSizeLimit() != 0 {

			/* cache lookup for absolute path */

			if (*t) == 0 {
				*t = time(0)
			}
			if g.Assign(&bucket, RealpathCacheFind(path, len_, *t)) != nil {
				if is_dir != 0 && bucket.GetIsDir() == 0 {

					/* not a directory */

					return size_t - 1

					/* not a directory */

				} else {
					if link_is_dir != nil {
						*link_is_dir = bucket.GetIsDir()
					}
					memcpy(path, bucket.GetRealpath(), bucket.GetRealpathLen()+1)
					return bucket.GetRealpathLen()
				}
			}
		}
		if save != 0 && lstat(path, &st) < 0 {
			if use_realpath == 2 {

				/* file not found */

				return size_t - 1

				/* file not found */

			}

			/* continue resolution anyway but don't save result in the cache */

			save = 0

			/* continue resolution anyway but don't save result in the cache */

		}
		tmp = _emalloc(len_ + 1)
		memcpy(tmp, path, len_+1)
		if save != 0 && (st.st_mode&S_IFMT) == 0120000 {
			if g.PreInc(&(*ll)) > 32 || g.Assign(&j, int(readlink(tmp, path, 256))) == size_t-1 {

				/* too many links or broken symlinks */

				_efree(tmp)
				return size_t - 1
			}
			path[j] = 0
			if path[0] == '/' {
				j = TsrmRealpathR(path, 1, j, ll, t, use_realpath, is_dir, &directory)
				if j == size_t-1 {
					_efree(tmp)
					return size_t - 1
				}
			} else {
				if i+j >= 256-1 {
					_efree(tmp)
					return size_t - 1
				}
				memmove(path+i, path, j+1)
				memcpy(path, tmp, i-1)
				path[i-1] = '/'
				j = TsrmRealpathR(path, start, i+j, ll, t, use_realpath, is_dir, &directory)
				if j == size_t-1 {
					_efree(tmp)
					return size_t - 1
				}
			}
			if link_is_dir != nil {
				*link_is_dir = directory
			}
		} else {
			if save != 0 {
				directory = (st.st_mode & S_IFMT) == S_IFDIR
				if link_is_dir != nil {
					*link_is_dir = directory
				}
				if is_dir != 0 && directory == 0 {

					/* not a directory */

					_efree(tmp)
					return size_t - 1
				}
			}
			if i <= start+1 {
				j = start
			} else {

				/* some leading directories may be unaccessable */

				j = TsrmRealpathR(path, start, i-1, ll, t, g.Cond(save != 0, 1, use_realpath), 1, nil)
				if j > start && j != size_t-1 {
					path[g.PostInc(&j)] = '/'
				}
			}
			if j == size_t-1 || j+len_ >= 256-1+i {
				_efree(tmp)
				return size_t - 1
			}
			memcpy(path+j, tmp+i, len_-i+1)
			j += len_ - i
		}
		if save != 0 && start != 0 && CwdGlobals.GetRealpathCacheSizeLimit() != 0 {

			/* save absolute path in the cache */

			RealpathCacheAdd(tmp, len_, path, j, directory, *t)

			/* save absolute path in the cache */

		}
		_efree(tmp)
		return j
	}
}

/* }}} */

func VirtualFileEx(state *CwdState, path *byte, verify_path VerifyPathFunc, use_realpath int) int {
	var path_length int = strlen(path)
	var resolved_path []byte = []byte{0}
	var start int = 1
	var ll int = 0
	var t int64
	var ret int
	var add_slash int
	var tmp any
	if path_length == 0 || path_length >= 256-1 {
		errno = EINVAL
		return 1
	}

	/* cwd_length can be 0 when getcwd() fails.
	 * This can happen under solaris when a dir does not have read permissions
	 * but *does* have execute permissions */

	if path[0] != '/' {
		if state.GetCwdLength() == 0 {

			/* resolve relative path */

			start = 0
			memcpy(resolved_path, path, path_length+1)
		} else {
			var state_cwd_length int = state.GetCwdLength()
			if path_length+state_cwd_length+1 >= 256-1 {
				errno = ENAMETOOLONG
				return 1
			}
			memcpy(resolved_path, state.GetCwd(), state_cwd_length)
			if resolved_path[state_cwd_length-1] == '/' {
				memcpy(resolved_path+state_cwd_length, path, path_length+1)
				path_length += state_cwd_length
			} else {
				resolved_path[state_cwd_length] = '/'
				memcpy(resolved_path+state_cwd_length+1, path, path_length+1)
				path_length += state_cwd_length + 1
			}
		}
	} else {
		memcpy(resolved_path, path, path_length+1)
	}
	add_slash = use_realpath != 2 && path_length > 0 && resolved_path[path_length-1] == '/'
	if CwdGlobals.GetRealpathCacheTtl() != 0 {
		t = 0
	} else {
		t = -1
	}
	path_length = TsrmRealpathR(resolved_path, start, path_length, &ll, &t, use_realpath, 0, nil)
	if path_length == size_t-1 {
		errno = ENOENT
		return 1
	}
	if start == 0 && path_length == 0 {
		resolved_path[g.PostInc(&path_length)] = '.'
	}
	if add_slash != 0 && path_length != 0 && resolved_path[path_length-1] != '/' {
		if path_length >= 256-1 {
			return -1
		}
		resolved_path[g.PostInc(&path_length)] = '/'
	}
	resolved_path[path_length] = 0
	if verify_path != nil {
		var old_state CwdState
		&old_state.SetCwdLength(state.GetCwdLength())
		&old_state.SetCwd((*byte)(_emalloc(state.GetCwdLength() + 1)))
		memcpy(&old_state.GetCwd(), state.GetCwd(), state.GetCwdLength()+1)
		state.SetCwdLength(path_length)
		tmp = _erealloc(state.GetCwd(), state.GetCwdLength()+1)
		state.SetCwd((*byte)(tmp))
		memcpy(state.GetCwd(), resolved_path, state.GetCwdLength()+1)
		if verify_path(state) != 0 {
			_efree(state.GetCwd())
			state.SetCwdLength(0)
			*state = old_state
			ret = 1
		} else {
			_efree(&old_state.GetCwd())
			&old_state.SetCwdLength(0)
			ret = 0
		}
	} else {
		state.SetCwdLength(path_length)
		tmp = _erealloc(state.GetCwd(), state.GetCwdLength()+1)
		state.SetCwd((*byte)(tmp))
		memcpy(state.GetCwd(), resolved_path, state.GetCwdLength()+1)
		ret = 0
	}
	return ret
}

/* }}} */

func VirtualChdir(path *byte) int {
	if VirtualFileEx(&(CwdGlobals.GetCwd()), path, PhpIsDirOk, 2) != 0 {
		return -1
	} else {
		return 0
	}
}

/* }}} */

func VirtualChdirFile(path *byte, p_chdir func(path *byte) int) int {
	var length int = strlen(path)
	var temp *byte
	var retval int
	if length == 0 {
		return 1
	}
	for g.PreDec(&length) < SIZE_MAX && path[length] != '/' {

	}
	if length == SIZE_MAX {

		/* No directory only file name */

		errno = ENOENT
		return -1
	}
	if length == 0 && path[0] == '/' {
		length++
	}
	temp = (*byte)(_emalloc(length + 1))
	memcpy(temp, path, length)
	temp[length] = 0
	retval = p_chdir(temp)
	_efree(temp)
	return retval
}

/* }}} */

func VirtualRealpath(path *byte, real_path *byte) *byte {
	var new_state CwdState
	var retval *byte
	var cwd []byte

	/* realpath("") returns CWD */

	if !(*path) {
		new_state.SetCwd((*byte)(_emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
		if getcwd(cwd, 256) {
			path = cwd
		}
	} else if path[0] != '/' {
		&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
		&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
		memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	} else {
		new_state.SetCwd((*byte)(_emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
	}
	if VirtualFileEx(&new_state, path, nil, 2) == 0 {
		var len_ int = g.CondF2(new_state.GetCwdLength() > 256-1, 256-1, func() int { return new_state.GetCwdLength() })
		memcpy(real_path, new_state.GetCwd(), len_)
		real_path[len_] = '0'
		retval = real_path
	} else {
		retval = nil
	}
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualFilepathEx(path *byte, filepath **byte, verify_path VerifyPathFunc) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	retval = VirtualFileEx(&new_state, path, verify_path, 1)
	*filepath = new_state.GetCwd()
	return retval
}

/* }}} */

func VirtualFilepath(path *byte, filepath **byte) int {
	return VirtualFilepathEx(path, filepath, PhpIsFileOk)
}

/* }}} */

func VirtualFopen(path *byte, mode *byte) *FILE {
	var new_state CwdState
	var f *FILE
	if path[0] == '0' {
		return nil
	}
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 0) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return nil
	}
	f = fopen(new_state.GetCwd(), mode)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return f
}

/* }}} */

func VirtualAccess(pathname *byte, mode int) int {
	var new_state CwdState
	var ret int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, pathname, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	ret = access(new_state.GetCwd(), mode)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return ret
}

/* }}} */

func VirtualUtime(filename *byte, buf *__struct__utimbuf) int {
	var new_state CwdState
	var ret int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, filename, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	ret = utime(new_state.GetCwd(), buf)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return ret
}

/* }}} */

func VirtualChmod(filename *byte, mode mode_t) int {
	var new_state CwdState
	var ret int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, filename, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	ret = chmod(new_state.GetCwd(), mode)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return ret
}

/* }}} */

func VirtualChown(filename *byte, owner uid_t, group gid_t, link int) int {
	var new_state CwdState
	var ret int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, filename, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	if link != 0 {
		ret = lchown(new_state.GetCwd(), owner, group)
	} else {
		ret = chown(new_state.GetCwd(), owner, group)
	}
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return ret
}

/* }}} */

func VirtualOpen(path *byte, flags int, _ ...any) int {
	var new_state CwdState
	var f int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 1) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	if (flags & O_CREAT) != 0 {
		var mode mode_t
		var arg va_list
		va_start(arg, flags)
		mode = mode_t(__va_arg(arg, int(_)))
		va_end(arg)
		f = open(new_state.GetCwd(), flags, mode)
	} else {
		f = open(new_state.GetCwd(), flags)
	}
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return f
}

/* }}} */

func VirtualCreat(path *byte, mode mode_t) int {
	var new_state CwdState
	var f int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 1) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	f = creat(new_state.GetCwd(), mode)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return f
}

/* }}} */

func VirtualRename(oldname *byte, newname *byte) int {
	var old_state CwdState
	var new_state CwdState
	var retval int
	&old_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&old_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&old_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&old_state, oldname, nil, 0) != 0 {
		_efree(&old_state.GetCwd())
		&old_state.SetCwdLength(0)
		return -1
	}
	oldname = old_state.GetCwd()
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, newname, nil, 0) != 0 {
		_efree(&old_state.GetCwd())
		&old_state.SetCwdLength(0)
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	newname = new_state.GetCwd()

	/* rename on windows will fail if newname already exists.
	   MoveFileEx has to be used */

	retval = rename(oldname, newname)
	_efree(&old_state.GetCwd())
	&old_state.SetCwdLength(0)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualStat(path *byte, buf *ZendStatT) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	retval = stat(new_state.GetCwd(), buf)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualLstat(path *byte, buf *ZendStatT) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 0) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	retval = lstat(new_state.GetCwd(), buf)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualUnlink(path *byte) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, path, nil, 0) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	retval = unlink(new_state.GetCwd())
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualMkdir(pathname *byte, mode mode_t) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, pathname, nil, 1) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	retval = mkdir(new_state.GetCwd(), mode)
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualRmdir(pathname *byte) int {
	var new_state CwdState
	var retval int
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, pathname, nil, 0) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return -1
	}
	retval = rmdir(new_state.GetCwd())
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualOpendir(pathname *byte) *DIR {
	var new_state CwdState
	var retval *DIR
	&new_state.SetCwdLength(&(CwdGlobals.GetCwd()).GetCwdLength())
	&new_state.SetCwd((*byte)(_emalloc(&(CwdGlobals.GetCwd()).GetCwdLength() + 1)))
	memcpy(&new_state.GetCwd(), &(CwdGlobals.GetCwd()).GetCwd(), &(CwdGlobals.GetCwd()).GetCwdLength()+1)
	if VirtualFileEx(&new_state, pathname, nil, 2) != 0 {
		_efree(&new_state.GetCwd())
		&new_state.SetCwdLength(0)
		return nil
	}
	retval = opendir(new_state.GetCwd())
	_efree(&new_state.GetCwd())
	&new_state.SetCwdLength(0)
	return retval
}

/* }}} */

func VirtualPopen(command *byte, type_ *byte) *FILE {
	var command_length int
	var dir_length int
	var extra int = 0
	var command_line *byte
	var ptr *byte
	var dir *byte
	var retval *FILE
	command_length = strlen(command)
	dir_length = CwdGlobals.GetCwd().GetCwdLength()
	dir = CwdGlobals.GetCwd().GetCwd()
	for dir_length > 0 {
		if (*dir) == '\'' {
			extra += 3
		}
		dir++
		dir_length--
	}
	dir_length = CwdGlobals.GetCwd().GetCwdLength()
	dir = CwdGlobals.GetCwd().GetCwd()
	command_line = (*byte)(_emalloc(command_length + g.SizeOf("\"cd '' ; \"") + dir_length + extra + 1 + 1))
	ptr = command_line
	memcpy(ptr, "cd ", g.SizeOf("\"cd \"")-1)
	ptr += g.SizeOf("\"cd \"") - 1
	if CwdGlobals.GetCwd().GetCwdLength() == 0 {
		g.PostInc(&(*ptr)) = '/'
	} else {
		g.PostInc(&(*ptr)) = '\''
		for dir_length > 0 {
			switch *dir {
			case '\'':
				g.PostInc(&(*ptr)) = '\''
				g.PostInc(&(*ptr)) = '\\'
				g.PostInc(&(*ptr)) = '\''
			default:
				g.PostInc(&(*ptr)) = *dir
			}
			dir++
			dir_length--
		}
		g.PostInc(&(*ptr)) = '\''
	}
	g.PostInc(&(*ptr)) = ' '
	g.PostInc(&(*ptr)) = ';'
	g.PostInc(&(*ptr)) = ' '
	memcpy(ptr, command, command_length+1)
	retval = popen(command_line, type_)
	_efree(command_line)
	return retval
}

/* }}} */

func TsrmRealpath(path *byte, real_path *byte) *byte {
	var new_state CwdState
	var cwd []byte

	/* realpath("") returns CWD */

	if !(*path) {
		new_state.SetCwd((*byte)(_emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
		if getcwd(cwd, 256) {
			path = cwd
		}
	} else if path[0] != '/' && getcwd(cwd, 256) {
		new_state.SetCwd(_estrdup(cwd))
		new_state.SetCwdLength(strlen(cwd))
	} else {
		new_state.SetCwd((*byte)(_emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
	}
	if VirtualFileEx(&new_state, path, nil, 2) != 0 {
		_efree(new_state.GetCwd())
		return nil
	}
	if real_path != nil {
		var copy_len int = g.CondF2(new_state.GetCwdLength() > 256-1, 256-1, func() int { return new_state.GetCwdLength() })
		memcpy(real_path, new_state.GetCwd(), copy_len)
		real_path[copy_len] = '0'
		_efree(new_state.GetCwd())
		return real_path
	} else {
		return new_state.GetCwd()
	}
}

/* }}} */
