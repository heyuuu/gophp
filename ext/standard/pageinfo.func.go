package standard

import (
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/zpp"
	"os"
)

func PhpGetuid() int { return zend.CurrEntrance().Uid() }
func PhpGetgid() int { return zend.CurrEntrance().Gid() }
func ZifGetmyuid() (int, bool) {
	if uid := PhpGetuid(); uid >= 0 {
		return uid, true
	}
	return 0, false
}
func ZifGetmygid() (int, bool) {
	if gid := PhpGetgid(); gid >= 0 {
		return gid, true
	}
	return 0, false
}
func ZifGetmypid() int {
	return os.Getpid()
}
func ZifGetmyinode() (int, bool) {
	if inode := zend.CurrEntrance().Inode(); inode >= 0 {
		return inode, true
	}
	return 0, false
}
func ZifGetlastmod(executeData zpp.Ex, return_value zpp.Ret) (int, bool) {
	if mtime := zend.CurrEntrance().Mtime(); mtime >= 0 {
		return mtime, true
	}
	return 0, false
}
