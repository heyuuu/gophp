package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
)

func PhpFlock(fd int, operation int) int {
	var flck __struct__flock
	var ret int
	flck.l_len = 0
	flck.l_start = flck.l_len
	flck.l_whence = r.SEEK_SET
	if (operation & LOCK_SH) != 0 {
		flck.l_type = F_RDLCK
	} else if (operation & LOCK_EX) != 0 {
		flck.l_type = F_WRLCK
	} else if (operation & LOCK_UN) != 0 {
		flck.l_type = F_UNLCK
	} else {
		errno = EINVAL
		return -1
	}
	ret = fcntl(fd, b.Cond((operation&LOCK_NB) != 0, F_SETLK, F_SETLKW), &flck)
	if (operation&LOCK_NB) != 0 && ret == -1 && (errno == EACCES || errno == EAGAIN) {
		errno = core.EWOULDBLOCK
	}
	if ret != -1 {
		ret = 0
	}
	return ret
}
