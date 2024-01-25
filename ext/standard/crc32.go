package standard

import (
	"hash/crc32"
)

func ZifCrc32(str string) int {
	return int(crc32.ChecksumIEEE([]byte(str)))
}
