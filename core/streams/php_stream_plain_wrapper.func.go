package streams

import (
	r "sik/builtin/file"
	"sik/core"
	"sik/zend/types"
)

func PhpStreamFopenFromFile(file *r.FILE, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromFile(file, mode)
}
func PhpStreamFopenFromFd(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	return _phpStreamFopenFromFd(fd, mode, persistent_id)
}
func PhpStreamFopenFromPipe(file *r.FILE, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromPipe(file, mode)
}
func PhpStreamFopenTemporaryFile(dir *byte, pfx string, opened_path **types.String) *core.PhpStream {
	return _phpStreamFopenTemporaryFile(dir, pfx, opened_path)
}
func PhpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **types.String) *r.FILE {
	return _phpStreamOpenWrapperAsFile(path, mode, options, opened_path)
}
