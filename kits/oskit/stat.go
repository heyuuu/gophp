package oskit

import (
	"io/fs"
)

func IsRegularFile(s fs.FileInfo) bool {
	return s.Mode().IsRegular()
}

func IsLink(s fs.FileInfo) bool {
	return s.Mode()&fs.ModeSymlink != 0
}

type SysStat struct {
	Dev     int
	Mode    int
	Nlink   int
	Ino     int
	Uid     int
	Gid     int
	Rdev    int
	Atime   int
	Mtime   int
	Ctime   int
	Size    int
	Blocks  int
	Blksize int
	Flags   int
	Gen     int
}

func ParseSysStat(s fs.FileInfo) (SysStat, bool) {
	if sysStat, ok := s.Sys().(SysStat); ok {
		return sysStat, true
	}
	return internalParseSysStat(s)
}
