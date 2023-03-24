package zend

import (
	r "sik/builtin/file"
	"sik/zend/types"
)

func ZendStreamStdioReader(handle *r.FILE, buf *byte, len_ int) ssize_t {
	return r.Fread(buf, 1, len_, handle)
}
func ZendStreamStdioCloser(handle any) {
	if handle && (*r.FILE)(handle != stdin) != nil {
		r.Fclose((*r.FILE)(handle))
	}
}
func ZendStreamStdioFsizer(handle any) int {
	var buf ZendStatT
	if handle && ZendFstat(fileno((*r.FILE)(handle)), &buf) == 0 {
		return buf.st_size
	}
	return -1
}
func ZendStreamInitFp(handle *ZendFileHandle, fp *r.FILE, filename string) {
	handle.InitFp(fp, filename)
}
func ZendStreamInitFilename(handle *ZendFileHandle, filename string) {
	handle.InitFilename(filename)
}
func ZendStreamOpen(filename string, handle *ZendFileHandle) int {
	if handle.Open(filename) {
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendStreamFixup(file_handle *ZendFileHandle, buf **byte, len_ *int) int {
	if buf_, ok := file_handle.Fixup(); ok {
		*buf = (*byte)(buf_)
		*len_ = len(buf_)
		return types.SUCCESS
	} else {
		return types.FAILURE
	}
}
func ZendFileHandleDtor(fh *ZendFileHandle) { fh.Destroy() }
func ZendCompareFileHandles(fh1 *ZendFileHandle, fh2 *ZendFileHandle) int {
	result := IsFileHandlesEquals(fh1, fh2)
	if result {
		return 1
	} else {
		return 0
	}
}
