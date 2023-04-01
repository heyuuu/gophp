package zend

const SPEC_START_MASK = 0xffff
const SPEC_EXTRA_MASK = 0xfffc0000
const SPEC_RULE_OP1 = 0x10000
const SPEC_RULE_OP2 = 0x20000
const SPEC_RULE_OP_DATA = 0x40000
const SPEC_RULE_RETVAL = 0x80000
const SPEC_RULE_QUICK_ARG = 0x100000
const SPEC_RULE_SMART_BRANCH = 0x200000
const SPEC_RULE_COMMUTATIVE = 0x800000
const SPEC_RULE_ISSET = 0x1000000

var ZendSpecHandlers *uint32
var ZendOpcodeHandlers []OpcodeHandlerT

type OpcodeHandlerT func(executeData *ZendExecuteData) int

const OPLINE = executeData.GetOpline()
