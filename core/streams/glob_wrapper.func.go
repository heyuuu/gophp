// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

func _phpGlobStreamGetPath(stream *core.PhpStream, plen *int) *byte {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
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
func _phpGlobStreamGetPattern(stream *core.PhpStream, plen *int) *byte {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
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
func _phpGlobStreamGetCount(stream *core.PhpStream, pflags *int) int {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
	if pglob != nil {
		if pflags != nil {
			*pflags = pglob.GetFlags()
		}
		return pglob.GetGlob().gl_pathc
	} else {
		if pflags != nil {
			*pflags = 0
		}
		return 0
	}
}
func PhpGlobStreamPathSplit(pglob *GlobST, path *byte, get_path int, p_file **byte) {
	var pos *byte
	var gpath *byte = path
	if b.Assign(&pos, strrchr(path, '/')) != nil {
		path = pos + 1
	}
	*p_file = path
	if get_path != 0 {
		if pglob.GetPath() != nil {
			zend.Efree(pglob.GetPath())
		}
		if path-gpath > 1 {
			path--
		}
		pglob.SetPathLen(path - gpath)
		pglob.SetPath(zend.Estrndup(gpath, pglob.GetPathLen()))
	}
}
func PhpGlobStreamRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)
	var path *byte

	/* avoid problems if someone mis-uses the stream */

	if count == b.SizeOf("php_stream_dirent") && pglob != nil {
		if pglob.GetIndex() < int(pglob.GetGlob().gl_pathc) {
			PhpGlobStreamPathSplit(pglob, pglob.GetGlob().gl_pathv[b.PostInc(&(pglob.GetIndex()))], pglob.GetFlags()&GLOB_APPEND, &path)
			core.PHP_STRLCPY(ent.GetDName(), path, b.SizeOf("ent -> d_name"), strlen(path))
			return b.SizeOf("php_stream_dirent")
		}
		pglob.SetIndex(pglob.GetGlob().gl_pathc)
		if pglob.GetPath() != nil {
			zend.Efree(pglob.GetPath())
			pglob.SetPath(nil)
		}
	}
	return -1
}
func PhpGlobStreamClose(stream *core.PhpStream, close_handle int) int {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
	if pglob != nil {
		pglob.SetIndex(0)
		globfree(&pglob.GetGlob())
		if pglob.GetPath() != nil {
			zend.Efree(pglob.GetPath())
		}
		if pglob.GetPattern() != nil {
			zend.Efree(pglob.GetPattern())
		}
	}
	zend.Efree(stream.GetAbstract())
	return 0
}
func PhpGlobStreamRewind(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var pglob *GlobST = (*GlobST)(stream.GetAbstract())
	if pglob != nil {
		pglob.SetIndex(0)
		if pglob.GetPath() != nil {
			zend.Efree(pglob.GetPath())
			pglob.SetPath(nil)
		}
	}
	return 0
}
func PhpGlobStreamOpener(wrapper *core.PhpStreamWrapper, path *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var pglob *GlobST
	var ret int
	var tmp *byte
	var pos *byte
	if !(strncmp(path, "glob://", b.SizeOf("\"glob://\"")-1)) {
		path += b.SizeOf("\"glob://\"") - 1
		if opened_path != nil {
			*opened_path = zend.ZendStringInit(path, strlen(path), 0)
		}
	}
	if (options&core.STREAM_DISABLE_OPEN_BASEDIR) == 0 && core.PhpCheckOpenBasedir(path) != 0 {
		return nil
	}
	pglob = zend.Ecalloc(b.SizeOf("* pglob"), 1)
	if 0 != b.Assign(&ret, glob(path, pglob.GetFlags()&GLOB_FLAGMASK, nil, &pglob.GetGlob())) {
		zend.Efree(pglob)
		return nil
	}
	pos = path
	if b.Assign(&tmp, strrchr(pos, '/')) != nil {
		pos = tmp + 1
	}
	pglob.SetPatternLen(strlen(pos))
	pglob.SetPattern(zend.Estrndup(pos, pglob.GetPatternLen()))
	pglob.SetIsAppend(true)
	if pglob.GetGlob().gl_pathc {
		PhpGlobStreamPathSplit(pglob, pglob.GetGlob().gl_pathv[0], 1, &tmp)
	} else {
		PhpGlobStreamPathSplit(pglob, path, 1, &tmp)
	}
	return core.PhpStreamAlloc(&PhpGlobStreamOps, pglob, 0, mode)
}
