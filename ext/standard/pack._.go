package standard

/* Whether machine is little endian */

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
