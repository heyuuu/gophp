package standard

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard/str"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

/**
 * PhpBasicGlobals
 */
type PhpBasicGlobals struct {
	strTokState str.StrTokState

	user_shutdown_function_names *types.Array
	putenv_ht                    types.Array
	locale_string                *types.String
	locale_changed               types.ZendBool
	str_ebuf                     []byte
	array_walk_fci               types.ZendFcallInfo
	array_walk_fci_cache         types.ZendFcallInfoCache
	user_compare_fci             types.ZendFcallInfo
	user_compare_fci_cache       types.ZendFcallInfoCache
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
	state                        []uint32
	next                         *uint32
	left                         int
	mt_rand_is_seeded            types.ZendBool
	mt_rand_mode                 zend.ZendLong
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
	url_adapt_session_hosts_ht types.Array
	url_adapt_output_ex        UrlAdaptStateExT
	url_adapt_output_hosts_ht  types.Array
	mmap_file                  any
	mmap_len                   int
	user_filter_map            *types.Array
	umask                      int
	unserialize_max_depth      zend.ZendLong
}

func (this *PhpBasicGlobals) GetUrlAdaptSessionEx() UrlAdaptStateExT {
	return this.url_adapt_session_ex
}

func (this *PhpBasicGlobals) GetUrlAdaptSessionHostsHt() types.Array {
	return this.url_adapt_session_hosts_ht
}

func (this *PhpBasicGlobals) GetUrlAdaptOutputEx() UrlAdaptStateExT { return this.url_adapt_output_ex }
func (this *PhpBasicGlobals) GetUrlAdaptOutputHostsHt() types.Array {
	return this.url_adapt_output_hosts_ht
}

func (this *PhpBasicGlobals) GetStrTokState() *str.StrTokState {
	return &this.strTokState
}

/**
 * PutenvEntry
 */
type PutenvEntry struct {
	putenv_string  *byte
	previous_value *byte
	key            *byte
	key_len        int
}

func (this *PutenvEntry) GetPutenvString() *byte       { return this.putenv_string }
func (this *PutenvEntry) SetPutenvString(value *byte)  { this.putenv_string = value }
func (this *PutenvEntry) GetPreviousValue() *byte      { return this.previous_value }
func (this *PutenvEntry) SetPreviousValue(value *byte) { this.previous_value = value }
func (this *PutenvEntry) GetKey() *byte                { return this.key }
func (this *PutenvEntry) SetKey(value *byte)           { this.key = value }
func (this *PutenvEntry) GetKeyLen() int               { return this.key_len }
func (this *PutenvEntry) SetKeyLen(value int)          { this.key_len = value }

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
