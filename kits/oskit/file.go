package oskit

import (
	"os"
	"unsafe"
)

func ReadFileAsString(name string) (string, error) {
	bytes, err := os.ReadFile(name)
	s := unsafe.String(unsafe.SliceData(bytes), len(bytes))
	return s, err
}
