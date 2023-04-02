package zend

const SPEC_START_MASK = 0xffff
const SPEC_EXTRA_MASK = 0xfffc0000
const SPEC_RULE_OP1 = 0x10000           // 1 << 16
const SPEC_RULE_OP2 = 0x20000           // 1 << 17
const SPEC_RULE_OP_DATA = 0x40000       // 1 << 18
const SPEC_RULE_RETVAL = 0x80000        // 1 << 19
const SPEC_RULE_QUICK_ARG = 0x100000    // 1 << 20
const SPEC_RULE_SMART_BRANCH = 0x200000 // 1 << 21
const SPEC_RULE_COMMUTATIVE = 0x800000  // 1 << 23
const SPEC_RULE_ISSET = 0x1000000       // 1 << 24

var ZendSpecHandlers *uint32
var ZendOpcodeHandlers []OpcodeHandlerT

type OpcodeHandlerT func(executeData *ZendExecuteData) int

const OPLINE = executeData.GetOpline()
