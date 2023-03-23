package zend

const SYMTABLE_CACHE_SIZE = 32

/* excpt.h on Digital Unix 4.0 defines function_table */

const EG_FLAGS_INITIAL = 0
const EG_FLAGS_IN_SHUTDOWN = 1 << 0
const EG_FLAGS_OBJECT_STORE_NO_REUSE = 1 << 1
const EG_FLAGS_IN_RESOURCE_SHUTDOWN = 1 << 2

type ZendPhpScannerEvent = int

const (
	ON_TOKEN = iota
	ON_FEEDBACK
	ON_STOP
)
