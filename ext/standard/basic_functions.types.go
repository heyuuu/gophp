// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

/**
 * PhpBasicGlobals
 */
type PhpBasicGlobals struct {
	user_shutdown_function_names *zend.HashTable
	putenv_ht                    zend.HashTable
	strtok_zval                  zend.Zval
	strtok_string                *byte
	locale_string                *zend.ZendString
	locale_changed               zend.ZendBool
	strtok_last                  *byte
	strtok_table                 []byte
	strtok_len                   zend.ZendUlong
	str_ebuf                     []byte
	array_walk_fci               zend.ZendFcallInfo
	array_walk_fci_cache         zend.ZendFcallInfoCache
	user_compare_fci             zend.ZendFcallInfo
	user_compare_fci_cache       zend.ZendFcallInfoCache
	user_tick_functions          *zend.ZendLlist
	active_ini_file_section      zend.Zval
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
	mt_rand_is_seeded            zend.ZendBool
	mt_rand_mode                 zend.ZendLong
	syslog_device                *byte
	incomplete_class             *zend.ZendClassEntry
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
	url_adapt_session_hosts_ht zend.HashTable
	url_adapt_output_ex        UrlAdaptStateExT
	url_adapt_output_hosts_ht  zend.HashTable
	mmap_file                  any
	mmap_len                   int
	user_filter_map            *zend.HashTable
	umask                      int
	unserialize_max_depth      zend.ZendLong
}

func (this *PhpBasicGlobals) GetUserShutdownFunctionNames() *zend.HashTable {
	return this.user_shutdown_function_names
}
func (this *PhpBasicGlobals) SetUserShutdownFunctionNames(value *zend.HashTable) {
	this.user_shutdown_function_names = value
}
func (this *PhpBasicGlobals) GetPutenvHt() zend.HashTable              { return this.putenv_ht }
func (this *PhpBasicGlobals) SetPutenvHt(value zend.HashTable)         { this.putenv_ht = value }
func (this *PhpBasicGlobals) GetStrtokZval() zend.Zval                 { return this.strtok_zval }
func (this *PhpBasicGlobals) SetStrtokZval(value zend.Zval)            { this.strtok_zval = value }
func (this *PhpBasicGlobals) GetStrtokString() *byte                   { return this.strtok_string }
func (this *PhpBasicGlobals) SetStrtokString(value *byte)              { this.strtok_string = value }
func (this *PhpBasicGlobals) GetLocaleString() *zend.ZendString        { return this.locale_string }
func (this *PhpBasicGlobals) SetLocaleString(value *zend.ZendString)   { this.locale_string = value }
func (this *PhpBasicGlobals) GetLocaleChanged() zend.ZendBool          { return this.locale_changed }
func (this *PhpBasicGlobals) SetLocaleChanged(value zend.ZendBool)     { this.locale_changed = value }
func (this *PhpBasicGlobals) GetStrtokLast() *byte                     { return this.strtok_last }
func (this *PhpBasicGlobals) SetStrtokLast(value *byte)                { this.strtok_last = value }
func (this *PhpBasicGlobals) GetStrtokTable() []byte                   { return this.strtok_table }
func (this *PhpBasicGlobals) SetStrtokTable(value []byte)              { this.strtok_table = value }
func (this *PhpBasicGlobals) GetStrtokLen() zend.ZendUlong             { return this.strtok_len }
func (this *PhpBasicGlobals) SetStrtokLen(value zend.ZendUlong)        { this.strtok_len = value }
func (this *PhpBasicGlobals) GetStrEbuf() []byte                       { return this.str_ebuf }
func (this *PhpBasicGlobals) SetStrEbuf(value []byte)                  { this.str_ebuf = value }
func (this *PhpBasicGlobals) GetArrayWalkFci() zend.ZendFcallInfo      { return this.array_walk_fci }
func (this *PhpBasicGlobals) SetArrayWalkFci(value zend.ZendFcallInfo) { this.array_walk_fci = value }
func (this *PhpBasicGlobals) GetArrayWalkFciCache() zend.ZendFcallInfoCache {
	return this.array_walk_fci_cache
}
func (this *PhpBasicGlobals) SetArrayWalkFciCache(value zend.ZendFcallInfoCache) {
	this.array_walk_fci_cache = value
}
func (this *PhpBasicGlobals) GetUserCompareFci() zend.ZendFcallInfo { return this.user_compare_fci }
func (this *PhpBasicGlobals) SetUserCompareFci(value zend.ZendFcallInfo) {
	this.user_compare_fci = value
}
func (this *PhpBasicGlobals) GetUserCompareFciCache() zend.ZendFcallInfoCache {
	return this.user_compare_fci_cache
}
func (this *PhpBasicGlobals) SetUserCompareFciCache(value zend.ZendFcallInfoCache) {
	this.user_compare_fci_cache = value
}
func (this *PhpBasicGlobals) GetUserTickFunctions() *zend.ZendLlist { return this.user_tick_functions }
func (this *PhpBasicGlobals) SetUserTickFunctions(value *zend.ZendLlist) {
	this.user_tick_functions = value
}
func (this *PhpBasicGlobals) GetActiveIniFileSection() zend.Zval { return this.active_ini_file_section }
func (this *PhpBasicGlobals) SetActiveIniFileSection(value zend.Zval) {
	this.active_ini_file_section = value
}
func (this *PhpBasicGlobals) GetPageUid() zend.ZendLong                { return this.page_uid }
func (this *PhpBasicGlobals) SetPageUid(value zend.ZendLong)           { this.page_uid = value }
func (this *PhpBasicGlobals) GetPageGid() zend.ZendLong                { return this.page_gid }
func (this *PhpBasicGlobals) SetPageGid(value zend.ZendLong)           { this.page_gid = value }
func (this *PhpBasicGlobals) GetPageInode() zend.ZendLong              { return this.page_inode }
func (this *PhpBasicGlobals) SetPageInode(value zend.ZendLong)         { this.page_inode = value }
func (this *PhpBasicGlobals) GetPageMtime() int64                      { return this.page_mtime }
func (this *PhpBasicGlobals) SetPageMtime(value int64)                 { this.page_mtime = value }
func (this *PhpBasicGlobals) GetCurrentStatFile() *byte                { return this.CurrentStatFile }
func (this *PhpBasicGlobals) SetCurrentStatFile(value *byte)           { this.CurrentStatFile = value }
func (this *PhpBasicGlobals) GetCurrentLStatFile() **byte              { return this.CurrentLStatFile }
func (this *PhpBasicGlobals) SetCurrentLStatFile(value **byte)         { this.CurrentLStatFile = value }
func (this *PhpBasicGlobals) GetSsb() core.PhpStreamStatbuf            { return this.ssb }
func (this *PhpBasicGlobals) SetSsb(value core.PhpStreamStatbuf)       { this.ssb = value }
func (this *PhpBasicGlobals) GetLssb() core.PhpStreamStatbuf           { return this.lssb }
func (this *PhpBasicGlobals) SetLssb(value core.PhpStreamStatbuf)      { this.lssb = value }
func (this *PhpBasicGlobals) GetState() []uint32                       { return this.state }
func (this *PhpBasicGlobals) SetState(value []uint32)                  { this.state = value }
func (this *PhpBasicGlobals) GetNext() *uint32                         { return this.next }
func (this *PhpBasicGlobals) SetNext(value *uint32)                    { this.next = value }
func (this *PhpBasicGlobals) GetLeft() int                             { return this.left }
func (this *PhpBasicGlobals) SetLeft(value int)                        { this.left = value }
func (this *PhpBasicGlobals) GetMtRandIsSeeded() zend.ZendBool         { return this.mt_rand_is_seeded }
func (this *PhpBasicGlobals) SetMtRandIsSeeded(value zend.ZendBool)    { this.mt_rand_is_seeded = value }
func (this *PhpBasicGlobals) GetMtRandMode() zend.ZendLong             { return this.mt_rand_mode }
func (this *PhpBasicGlobals) SetMtRandMode(value zend.ZendLong)        { this.mt_rand_mode = value }
func (this *PhpBasicGlobals) GetSyslogDevice() *byte                   { return this.syslog_device }
func (this *PhpBasicGlobals) SetSyslogDevice(value *byte)              { this.syslog_device = value }
func (this *PhpBasicGlobals) GetIncompleteClass() *zend.ZendClassEntry { return this.incomplete_class }
func (this *PhpBasicGlobals) SetIncompleteClass(value *zend.ZendClassEntry) {
	this.incomplete_class = value
}
func (this *PhpBasicGlobals) GetSerializeLock() unsigned               { return this.serialize_lock }
func (this *PhpBasicGlobals) SetSerializeLock(value unsigned)          { this.serialize_lock = value }
func (this *PhpBasicGlobals) GetSerializeData() *PhpSerializeData      { return this.serialize.data }
func (this *PhpBasicGlobals) SetSerializeData(value *PhpSerializeData) { this.serialize.data = value }
func (this *PhpBasicGlobals) GetSerializeLevel() unsigned              { return this.serialize.level }
func (this *PhpBasicGlobals) SetSerializeLevel(value unsigned)         { this.serialize.level = value }
func (this *PhpBasicGlobals) GetUnserializeData() *PhpUnserializeData  { return this.unserialize.data }
func (this *PhpBasicGlobals) SetUnserializeData(value *PhpUnserializeData) {
	this.unserialize.data = value
}
func (this *PhpBasicGlobals) GetUnserializeLevel() unsigned      { return this.unserialize.level }
func (this *PhpBasicGlobals) SetUnserializeLevel(value unsigned) { this.unserialize.level = value }
func (this *PhpBasicGlobals) GetUrlAdaptSessionEx() UrlAdaptStateExT {
	return this.url_adapt_session_ex
}
func (this *PhpBasicGlobals) SetUrlAdaptSessionEx(value UrlAdaptStateExT) {
	this.url_adapt_session_ex = value
}
func (this *PhpBasicGlobals) GetUrlAdaptSessionHostsHt() zend.HashTable {
	return this.url_adapt_session_hosts_ht
}
func (this *PhpBasicGlobals) SetUrlAdaptSessionHostsHt(value zend.HashTable) {
	this.url_adapt_session_hosts_ht = value
}
func (this *PhpBasicGlobals) GetUrlAdaptOutputEx() UrlAdaptStateExT { return this.url_adapt_output_ex }
func (this *PhpBasicGlobals) SetUrlAdaptOutputEx(value UrlAdaptStateExT) {
	this.url_adapt_output_ex = value
}
func (this *PhpBasicGlobals) GetUrlAdaptOutputHostsHt() zend.HashTable {
	return this.url_adapt_output_hosts_ht
}
func (this *PhpBasicGlobals) SetUrlAdaptOutputHostsHt(value zend.HashTable) {
	this.url_adapt_output_hosts_ht = value
}
func (this *PhpBasicGlobals) GetMmapFile() any                       { return this.mmap_file }
func (this *PhpBasicGlobals) SetMmapFile(value any)                  { this.mmap_file = value }
func (this *PhpBasicGlobals) GetMmapLen() int                        { return this.mmap_len }
func (this *PhpBasicGlobals) SetMmapLen(value int)                   { this.mmap_len = value }
func (this *PhpBasicGlobals) GetUserFilterMap() *zend.HashTable      { return this.user_filter_map }
func (this *PhpBasicGlobals) SetUserFilterMap(value *zend.HashTable) { this.user_filter_map = value }
func (this *PhpBasicGlobals) GetUmask() int                          { return this.umask }
func (this *PhpBasicGlobals) SetUmask(value int)                     { this.umask = value }
func (this *PhpBasicGlobals) GetUnserializeMaxDepth() zend.ZendLong {
	return this.unserialize_max_depth
}
func (this *PhpBasicGlobals) SetUnserializeMaxDepth(value zend.ZendLong) {
	this.unserialize_max_depth = value
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
	arguments *zend.Zval
	arg_count int
}

func (this *PhpShutdownFunctionEntry) GetArguments() *zend.Zval      { return this.arguments }
func (this *PhpShutdownFunctionEntry) SetArguments(value *zend.Zval) { this.arguments = value }
func (this *PhpShutdownFunctionEntry) GetArgCount() int              { return this.arg_count }
func (this *PhpShutdownFunctionEntry) SetArgCount(value int)         { this.arg_count = value }

/**
 * UserTickFunctionEntry
 */
type UserTickFunctionEntry struct {
	arguments *zend.Zval
	arg_count int
	calling   int
}

func (this *UserTickFunctionEntry) GetArguments() *zend.Zval      { return this.arguments }
func (this *UserTickFunctionEntry) SetArguments(value *zend.Zval) { this.arguments = value }
func (this *UserTickFunctionEntry) GetArgCount() int              { return this.arg_count }
func (this *UserTickFunctionEntry) SetArgCount(value int)         { this.arg_count = value }
func (this *UserTickFunctionEntry) GetCalling() int               { return this.calling }
func (this *UserTickFunctionEntry) SetCalling(value int)          { this.calling = value }
