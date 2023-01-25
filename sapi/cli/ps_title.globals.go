// <<generate>>

package cli

const PS_TITLE_SUCCESS = 0
const PS_TITLE_NOT_AVAILABLE = 1
const PS_TITLE_NOT_INITIALIZED = 2
const PS_TITLE_BUFFER_NOT_AVAILABLE = 3
const PS_TITLE_WINDOWS_ERROR = 4

var Environ **byte

const PS_PADDING = '0'

var PsBuffer *byte
var PsBufferSize int
var EmptyEnviron []*byte = []*byte{0}
var PsBufferCurLen int
var SaveArgc int
var SaveArgv **byte
var FrozenEnviron **byte
var NewEnviron ***byte
