package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"math/rand"
	"time"
)

/**
 * PhpBasicGlobals
 */
type serializeType struct {
	data  *PhpSerializeData
	level uint
}
type unserializeType struct {
	data  *PhpUnserializeData
	level uint
}
type PhpBasicGlobals struct {
	strTokState   str.StrTokState
	RandGenerator *rand.Rand

	userShutdownFunctions      []PhpShutdownFunction
	putenv_ht                  types.Array
	localeString               *types.String
	localeChanged              bool
	str_ebuf                   []byte
	active_ini_file_section    types.Zval
	CurrentStatFile            *byte
	CurrentLStatFile           **byte
	ssb                        core.PhpStreamStatbuf
	lssb                       core.PhpStreamStatbuf
	syslog_device              *byte
	incomplete_class           *types.ClassEntry
	serialize_lock             uint
	serialize                  serializeType
	unserialize                unserializeType
	url_adapt_session_ex       UrlAdaptStateExT
	url_adapt_session_hosts_ht *types.Array
	url_adapt_output_ex        UrlAdaptStateExT
	url_adapt_output_hosts_ht  *types.Array
	mmap_file                  any
	mmap_len                   int
	UserFilterMap              map[string]*PhpUserFilterData
	umask                      int
	unserialize_max_depth      zend.ZendLong
}

func (bg *PhpBasicGlobals) Ctor() {
	bg.ResetRandGenerator()
	bg.umask = -1
	bg.UserFilterMap = nil
	bg.serialize_lock = 0
	bg.serialize = serializeType{}
	bg.unserialize = unserializeType{}
	bg.url_adapt_session_ex = UrlAdaptStateExT{type_: 1}
	bg.url_adapt_output_ex = UrlAdaptStateExT{type_: 0}
	bg.url_adapt_session_hosts_ht = types.NewArray(0)
	bg.url_adapt_output_hosts_ht = types.NewArray(0)
	bg.incomplete_class = IncompleteClassEntry
}
func (bg *PhpBasicGlobals) Activate() {
	bg.serialize_lock = 0
	bg.serialize = serializeType{}
	bg.unserialize = unserializeType{}
	bg.localeString = nil
	bg.localeChanged = false
	bg.CurrentStatFile = nil
	bg.CurrentLStatFile = nil
	bg.syslog_device = nil
}
func (bg *PhpBasicGlobals) Deactivate() {
	bg.ResetRandGenerator()
}
func (bg *PhpBasicGlobals) Dtor() {
	if bg.url_adapt_session_ex.tags != nil {
		bg.url_adapt_session_ex.tags.Destroy()
	}
	if bg.url_adapt_output_ex.tags != nil {
		bg.url_adapt_output_ex.tags.Destroy()
	}
	bg.url_adapt_session_hosts_ht.Destroy()
	bg.url_adapt_output_hosts_ht.Destroy()
}

func (bg *PhpBasicGlobals) GetUrlAdaptSessionEx() UrlAdaptStateExT {
	return bg.url_adapt_session_ex
}

func (bg *PhpBasicGlobals) GetUrlAdaptSessionHostsHt() *types.Array {
	return bg.url_adapt_session_hosts_ht
}

func (bg *PhpBasicGlobals) GetUrlAdaptOutputEx() UrlAdaptStateExT { return bg.url_adapt_output_ex }
func (bg *PhpBasicGlobals) GetUrlAdaptOutputHostsHt() *types.Array {
	return bg.url_adapt_output_hosts_ht
}

func (bg *PhpBasicGlobals) GetStrTokState() *str.StrTokState {
	return &bg.strTokState
}

func (bg *PhpBasicGlobals) ResetRandGenerator() {
	bg.RandGenerator = nil
}
func (bg *PhpBasicGlobals) InitRandGenerator(seed int64) {
	bg.RandGenerator = rand.New(rand.NewSource(seed))
}
func (bg *PhpBasicGlobals) GetRandGenerator() *rand.Rand {
	if bg.RandGenerator == nil {
		seed := time.Now().UnixNano()
		bg.InitRandGenerator(seed)
	}
	return bg.RandGenerator
}

func (bg *PhpBasicGlobals) ResetUserShutdownFunctions() {
	bg.userShutdownFunctions = nil
}
func (bg *PhpBasicGlobals) HasUserShutdownFunctions() bool {
	return len(bg.userShutdownFunctions) != 0
}
func (bg *PhpBasicGlobals) AddUserShutdownFunction(entry PhpShutdownFunction) {
	bg.userShutdownFunctions = append(bg.userShutdownFunctions, entry)
}
func (bg *PhpBasicGlobals) EachUserShutdownFunction(fn func(*PhpShutdownFunction)) {
	for i := range bg.userShutdownFunctions {
		entry := &bg.userShutdownFunctions[i]
		fn(entry)
	}
}

/**
 * PhpShutdownFunction
 */
type PhpShutdownFunction struct {
	fn   *types.Zval
	args []types.Zval
}

func NewShutdownFunction(fn *types.Zval, args []types.Zval) *PhpShutdownFunction {
	return &PhpShutdownFunction{fn: fn, args: args}
}

func (sfe *PhpShutdownFunction) Fn() *types.Zval    { return sfe.fn }
func (sfe *PhpShutdownFunction) Args() []types.Zval { return sfe.args }

/**
 * UserTickFunctionEntry
 */
type UserTickFunctionEntry struct {
	arguments *types.Zval
	arg_count int
	calling   int
}

func (this *UserTickFunctionEntry) GetArguments() *types.Zval      { return this.arguments }
func (this *UserTickFunctionEntry) SetArguments(value *types.Zval) { this.arguments = value }
func (this *UserTickFunctionEntry) GetArgCount() int               { return this.arg_count }
func (this *UserTickFunctionEntry) SetArgCount(value int)          { this.arg_count = value }
func (this *UserTickFunctionEntry) GetCalling() int                { return this.calling }
func (this *UserTickFunctionEntry) SetCalling(value int)           { this.calling = value }
