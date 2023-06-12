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
type PhpBasicGlobals struct {
	strTokState   str.StrTokState
	RandGenerator *rand.Rand

	user_shutdown_function_names *types.Array
	putenv_ht                    types.Array
	locale_string                *types.String
	locale_changed               bool
	str_ebuf                     []byte
	array_walk_fci               types.ZendFcallInfo
	array_walk_fci_cache         types.ZendFcallInfoCache
	user_tick_functions          *zend.ZendLlist
	active_ini_file_section      types.Zval
	page_uid                     zend.ZendLong
	page_gid                     zend.ZendLong
	page_inode                   zend.ZendLong
	page_mtime                   int64
	CurrentStatFile              *byte
	CurrentLStatFile             **byte
	ssb                          core.PhpStreamStatbuf
	lssb                         core.PhpStreamStatbuf
	syslog_device                *byte
	incomplete_class             *types.ClassEntry
	serialize_lock               unsigned
	serialize                    struct {
		data  *PhpSerializeData
		level unsigned
	}
	unserialize struct {
		data  *PhpUnserializeData
		level unsigned
	}
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

func (this *PhpBasicGlobals) GetUrlAdaptSessionEx() UrlAdaptStateExT {
	return this.url_adapt_session_ex
}

func (this *PhpBasicGlobals) GetUrlAdaptSessionHostsHt() *types.Array {
	return this.url_adapt_session_hosts_ht
}

func (this *PhpBasicGlobals) GetUrlAdaptOutputEx() UrlAdaptStateExT { return this.url_adapt_output_ex }
func (this *PhpBasicGlobals) GetUrlAdaptOutputHostsHt() *types.Array {
	return this.url_adapt_output_hosts_ht
}

func (this *PhpBasicGlobals) GetStrTokState() *str.StrTokState {
	return &this.strTokState
}

func (this *PhpBasicGlobals) ResetRandGenerator() {
	this.RandGenerator = nil
}
func (this *PhpBasicGlobals) InitRandGenerator(seed int64) {
	this.RandGenerator = rand.New(rand.NewSource(seed))
}
func (this *PhpBasicGlobals) GetRandGenerator() *rand.Rand {
	if this.RandGenerator == nil {
		seed := time.Now().UnixNano()
		this.InitRandGenerator(seed)
	}
	return this.RandGenerator
}

/**
 * PhpShutdownFunctionEntry
 */
type PhpShutdownFunctionEntry struct {
	arguments *types.Zval
	arg_count int
}

func (this *PhpShutdownFunctionEntry) GetArguments() *types.Zval      { return this.arguments }
func (this *PhpShutdownFunctionEntry) SetArguments(value *types.Zval) { this.arguments = value }
func (this *PhpShutdownFunctionEntry) GetArgCount() int               { return this.arg_count }
func (this *PhpShutdownFunctionEntry) SetArgCount(value int)          { this.arg_count = value }

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
