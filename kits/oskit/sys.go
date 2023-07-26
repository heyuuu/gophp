package oskit

import (
	"golang.org/x/sys/unix"
	"os"
)

// 判断是否为 tty
func Isatty(f *os.File) bool {
	_, err := unix.IoctlGetWinsize(int(f.Fd()), unix.TIOCGWINSZ)
	return err == nil
}
