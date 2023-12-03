package oskit

import (
	"io/fs"
	"syscall"
)

func internalParseSysStat(s fs.FileInfo) (SysStat, bool) {
	switch sysStat := s.Sys().(type) {
	case SysStat:
		return sysStat, true
	case syscall.Stat_t:
		return SysStat{
			Dev:     int(sysStat.Dev),
			Mode:    int(sysStat.Mode),
			Nlink:   int(sysStat.Nlink),
			Ino:     int(sysStat.Ino),
			Uid:     int(sysStat.Uid),
			Gid:     int(sysStat.Gid),
			Rdev:    int(sysStat.Rdev),
			Atime:   int(sysStat.Atimespec.Sec),
			Mtime:   int(sysStat.Mtimespec.Sec),
			Ctime:   int(sysStat.Ctimespec.Sec),
			Size:    int(sysStat.Size),
			Blocks:  int(sysStat.Blocks),
			Blksize: int(sysStat.Blksize),
			Flags:   int(sysStat.Flags),
			Gen:     int(sysStat.Gen),
		}, true
	default:
		return SysStat{}, false
	}
}
