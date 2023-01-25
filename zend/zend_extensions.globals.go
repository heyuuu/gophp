// <<generate>>

package zend

const ZEND_EXTENSION_API_NO = 320190902

type StartupFuncT func(extension *ZendExtension) int
type ShutdownFuncT func(extension *ZendExtension)
type ActivateFuncT func()
type DeactivateFuncT func()
type MessageHandlerFuncT func(message int, arg any)
type OpArrayHandlerFuncT func(op_array *ZendOpArray)
type StatementHandlerFuncT func(frame *ZendExecuteData)
type FcallBeginHandlerFuncT func(frame *ZendExecuteData)
type FcallEndHandlerFuncT func(frame *ZendExecuteData)
type OpArrayCtorFuncT func(op_array *ZendOpArray)
type OpArrayDtorFuncT func(op_array *ZendOpArray)
type OpArrayPersistCalcFuncT func(op_array *ZendOpArray) int
type OpArrayPersistFuncT func(op_array *ZendOpArray, mem any) int

const ZEND_EXTMSG_NEW_EXTENSION = 1

var ZendExtensions ZendLlist

const ZEND_EXTENSIONS_HAVE_OP_ARRAY_CTOR = 1 << 0
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_DTOR = 1 << 1
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_HANDLER = 1 << 2
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST_CALC = 1 << 3
const ZEND_EXTENSIONS_HAVE_OP_ARRAY_PERSIST = 1 << 4

var ZendExtensionFlags uint32 = 0
var ZendOpArrayExtensionHandles int = 0
var LastResourceNumber int
