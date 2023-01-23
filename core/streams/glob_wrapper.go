// <<generate>>

package streams

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <main/streams/glob_wrapper.c>

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

// # include "php.h"

// # include "php_streams_int.h"

// # include < glob . h >

// #define GLOB_ONLYDIR       ( 1 << 30 )

// #define GLOB_FLAGMASK       ( ~ GLOB_ONLYDIR )

func _phpGlobStreamGetPath(stream *core.PhpStream, plen *int) *byte {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	if pglob != nil && pglob.GetPath() != nil {
		if plen != nil {
			*plen = pglob.GetPathLen()
		}
		return pglob.GetPath()
	} else {
		if plen != nil {
			*plen = 0
		}
		return nil
	}
}

/* }}} */

func _phpGlobStreamGetPattern(stream *core.PhpStream, plen *int) *byte {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	if pglob != nil && pglob.GetPattern() != nil {
		if plen != nil {
			*plen = pglob.GetPatternLen()
		}
		return pglob.GetPattern()
	} else {
		if plen != nil {
			*plen = 0
		}
		return nil
	}
}

/* }}} */

func _phpGlobStreamGetCount(stream *core.PhpStream, pflags *int) int {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	if pglob != nil {
		if pflags != nil {
			*pflags = pglob.GetFlags()
		}
		return pglob.glob.gl_pathc
	} else {
		if pflags != nil {
			*pflags = 0
		}
		return 0
	}
}

/* }}} */

func PhpGlobStreamPathSplit(pglob *GlobST, path *byte, get_path int, p_file **byte) {
	var pos *byte
	var gpath *byte = path
	if g.Assign(&pos, strrchr(path, '/')) != nil {
		path = pos + 1
	}
	*p_file = path
	if get_path != 0 {
		if pglob.GetPath() != nil {
			zend._efree(pglob.GetPath())
		}
		if path-gpath > 1 {
			path--
		}
		pglob.SetPathLen(path - gpath)
		pglob.SetPath(zend._estrndup(gpath, pglob.GetPathLen()))
	}
}

/* }}} */

func PhpGlobStreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)
	var path *byte

	/* avoid problems if someone mis-uses the stream */

	if count == g.SizeOf("php_stream_dirent") && pglob != nil {
		if pglob.GetIndex() < int(pglob.glob.gl_pathc) {
			PhpGlobStreamPathSplit(pglob, pglob.glob.gl_pathv[g.PostInc(&(pglob.GetIndex()))], pglob.GetFlags()&GLOB_APPEND, &path)
			var php_str_len int
			if strlen(path) >= g.SizeOf("ent -> d_name") {
				php_str_len = g.SizeOf("ent -> d_name") - 1
			} else {
				php_str_len = strlen(path)
			}
			memcpy(ent.d_name, path, php_str_len)
			ent.d_name[php_str_len] = '0'
			return g.SizeOf("php_stream_dirent")
		}
		pglob.SetIndex(pglob.glob.gl_pathc)
		if pglob.GetPath() != nil {
			zend._efree(pglob.GetPath())
			pglob.SetPath(nil)
		}
	}
	return -1
}

/* }}} */

func PhpGlobStreamClose(stream *core.PhpStream, close_handle int) int {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	if pglob != nil {
		pglob.SetIndex(0)
		globfree(&pglob.glob)
		if pglob.GetPath() != nil {
			zend._efree(pglob.GetPath())
		}
		if pglob.GetPattern() != nil {
			zend._efree(pglob.GetPattern())
		}
	}
	zend._efree(stream.abstract)
	return 0
}

/* {{{ */

func PhpGlobStreamRewind(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var pglob *GlobST = (*GlobST)(stream.abstract)
	if pglob != nil {
		pglob.SetIndex(0)
		if pglob.GetPath() != nil {
			zend._efree(pglob.GetPath())
			pglob.SetPath(nil)
		}
	}
	return 0
}

/* }}} */

var PhpGlobStreamOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpGlobStreamRead, PhpGlobStreamClose, nil, "glob", PhpGlobStreamRewind, nil, nil, nil}

/* {{{ php_glob_stream_opener */

func PhpGlobStreamOpener(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var pglob *GlobST
	var ret int
	var tmp *byte
	var pos *byte
	if !(strncmp(path, "glob://", g.SizeOf("\"glob://\"")-1)) {
		path += g.SizeOf("\"glob://\"") - 1
		if opened_path != nil {
			*opened_path = zend.ZendStringInit(path, strlen(path), 0)
		}
	}
	if (options&0x400) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	pglob = zend._ecalloc(g.SizeOf("* pglob"), 1)
	if 0 != g.Assign(&ret, glob(path, pglob.GetFlags() & ^(1<<30), nil, &pglob.glob)) {
		zend._efree(pglob)
		return nil
	}
	pos = path
	if g.Assign(&tmp, strrchr(pos, '/')) != nil {
		pos = tmp + 1
	}
	pglob.SetPatternLen(strlen(pos))
	pglob.SetPattern(zend._estrndup(pos, pglob.GetPatternLen()))
	pglob.SetFlags(pglob.GetFlags() | GLOB_APPEND)
	if pglob.glob.gl_pathc {
		PhpGlobStreamPathSplit(pglob, pglob.glob.gl_pathv[0], 1, &tmp)
	} else {
		PhpGlobStreamPathSplit(pglob, path, 1, &tmp)
	}
	return _phpStreamAlloc(&PhpGlobStreamOps, pglob, 0, mode)
}

/* }}} */

var PhpGlobStreamWrapperOps core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{nil, nil, nil, nil, PhpGlobStreamOpener, "glob", nil, nil, nil, nil, nil}
var PhpGlobStreamWrapper core.PhpStreamWrapper = core.PhpStreamWrapper{&PhpGlobStreamWrapperOps, nil, 0}
