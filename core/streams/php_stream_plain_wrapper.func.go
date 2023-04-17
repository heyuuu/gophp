package streams

import (
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStreamFopenFromFile(file *r.File, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromFile(file, mode)
}
func PhpStreamFopenFromFd(fd int, mode *byte, persistent_id *byte) *core.PhpStream {
	return _phpStreamFopenFromFd(fd, mode, persistent_id)
}
func PhpStreamFopenFromPipe(file *r.File, mode *byte) *core.PhpStream {
	return _phpStreamFopenFromPipe(file, mode)
}
func PhpStreamFopenTemporaryFile(dir *byte, pfx string, opened_path **types.String) *core.PhpStream {
	return _phpStreamFopenTemporaryFile(dir, pfx, opened_path)
}
func PhpStreamOpenWrapperAsFile(path *byte, mode string, options int, opened_path **types.String) *r.File {
	return _phpStreamOpenWrapperAsFile(path, mode, options, opened_path)
}
