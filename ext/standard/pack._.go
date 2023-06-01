package standard

import (
	"encoding/binary"
	"unsafe"
)

/* Whether machine is little endian */
var machineEndian binary.ByteOrder
var bigEndian binary.ByteOrder = binary.BigEndian
var litteEndian binary.ByteOrder = binary.LittleEndian

func init() {
	if isMachineLittleEndian() {
		machineEndian = litteEndian
	} else {
		machineEndian = bigEndian
	}
}

func isMachineLittleEndian() bool {
	n := 0x12345678
	return *(*byte)(unsafe.Pointer(&n)) == 0x78
}

var MachineLittleEndian byte

/* Mapping of byte from char (8bit) to long for machine endian */
var ByteMap []int

/* Mappings of bytes from int (machine dependent) to int for machine endian */

var IntMap []int

/* Mappings of bytes from shorts (16bit) for all endian environments */

var MachineEndianShortMap []int
var BigEndianShortMap []int
var LittleEndianShortMap []int

/* Mappings of bytes from longs (32bit) for all endian environments */

var MachineEndianLongMap []int
var BigEndianLongMap []int
var LittleEndianLongMap []int

/* Mappings of bytes from quads (64bit) for all endian environments */

var MachineEndianLonglongMap []int
var BigEndianLonglongMap []int
var LittleEndianLonglongMap []int

/* {{{ php_pack
 */
