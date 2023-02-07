// <<generate>>

package zend

/**
 * CwdState
 */
type CwdState struct {
	cwd        *byte
	cwd_length int
}

// func NewCwdState(cwd *byte, cwd_length int) *CwdState {
//     return &CwdState{
//         cwd:cwd,
//         cwd_length:cwd_length,
//     }
// }
// func MakeCwdState(cwd *byte, cwd_length int) CwdState {
//     return CwdState{
//         cwd:cwd,
//         cwd_length:cwd_length,
//     }
// }
func (this *CwdState) GetCwd() *byte          { return this.cwd }
func (this *CwdState) SetCwd(value *byte)     { this.cwd = value }
func (this *CwdState) GetCwdLength() int      { return this.cwd_length }
func (this *CwdState) SetCwdLength(value int) { this.cwd_length = value }

/**
 * RealpathCacheBucket
 */
type RealpathCacheBucket struct {
	key          ZendUlong
	path         *byte
	realpath     *byte
	next         *RealpathCacheBucket
	expires      int64
	path_len     uint16
	realpath_len uint16
	is_dir       uint8
}

// func NewRealpathCacheBucket(key ZendUlong, path *byte, realpath *byte, next *RealpathCacheBucket, expires int64, path_len uint16, realpath_len uint16, is_dir uint8) *RealpathCacheBucket {
//     return &RealpathCacheBucket{
//         key:key,
//         path:path,
//         realpath:realpath,
//         next:next,
//         expires:expires,
//         path_len:path_len,
//         realpath_len:realpath_len,
//         is_dir:is_dir,
//     }
// }
// func MakeRealpathCacheBucket(key ZendUlong, path *byte, realpath *byte, next *RealpathCacheBucket, expires int64, path_len uint16, realpath_len uint16, is_dir uint8) RealpathCacheBucket {
//     return RealpathCacheBucket{
//         key:key,
//         path:path,
//         realpath:realpath,
//         next:next,
//         expires:expires,
//         path_len:path_len,
//         realpath_len:realpath_len,
//         is_dir:is_dir,
//     }
// }
func (this *RealpathCacheBucket) GetKey() ZendUlong                  { return this.key }
func (this *RealpathCacheBucket) SetKey(value ZendUlong)             { this.key = value }
func (this *RealpathCacheBucket) GetPath() *byte                     { return this.path }
func (this *RealpathCacheBucket) SetPath(value *byte)                { this.path = value }
func (this *RealpathCacheBucket) GetRealpath() *byte                 { return this.realpath }
func (this *RealpathCacheBucket) SetRealpath(value *byte)            { this.realpath = value }
func (this *RealpathCacheBucket) GetNext() *RealpathCacheBucket      { return this.next }
func (this *RealpathCacheBucket) SetNext(value *RealpathCacheBucket) { this.next = value }
func (this *RealpathCacheBucket) GetExpires() int64                  { return this.expires }
func (this *RealpathCacheBucket) SetExpires(value int64)             { this.expires = value }
func (this *RealpathCacheBucket) GetPathLen() uint16                 { return this.path_len }
func (this *RealpathCacheBucket) SetPathLen(value uint16)            { this.path_len = value }
func (this *RealpathCacheBucket) GetRealpathLen() uint16             { return this.realpath_len }
func (this *RealpathCacheBucket) SetRealpathLen(value uint16)        { this.realpath_len = value }
func (this *RealpathCacheBucket) GetIsDir() uint8                    { return this.is_dir }
func (this *RealpathCacheBucket) SetIsDir(value uint8)               { this.is_dir = value }

/**
 * VirtualCwdGlobals
 */
type VirtualCwdGlobals struct {
	cwd                       CwdState
	realpath_cache_size       ZendLong
	realpath_cache_size_limit ZendLong
	realpath_cache_ttl        ZendLong
	realpath_cache            []*RealpathCacheBucket
}

// func NewVirtualCwdGlobals(cwd CwdState, realpath_cache_size ZendLong, realpath_cache_size_limit ZendLong, realpath_cache_ttl ZendLong, realpath_cache []*RealpathCacheBucket) *VirtualCwdGlobals {
//     return &VirtualCwdGlobals{
//         cwd:cwd,
//         realpath_cache_size:realpath_cache_size,
//         realpath_cache_size_limit:realpath_cache_size_limit,
//         realpath_cache_ttl:realpath_cache_ttl,
//         realpath_cache:realpath_cache,
//     }
// }
// func MakeVirtualCwdGlobals(cwd CwdState, realpath_cache_size ZendLong, realpath_cache_size_limit ZendLong, realpath_cache_ttl ZendLong, realpath_cache []*RealpathCacheBucket) VirtualCwdGlobals {
//     return VirtualCwdGlobals{
//         cwd:cwd,
//         realpath_cache_size:realpath_cache_size,
//         realpath_cache_size_limit:realpath_cache_size_limit,
//         realpath_cache_ttl:realpath_cache_ttl,
//         realpath_cache:realpath_cache,
//     }
// }
func (this *VirtualCwdGlobals) GetCwd() CwdState { return this.cwd }

// func (this *VirtualCwdGlobals) SetCwd(value CwdState) { this.cwd = value }
func (this *VirtualCwdGlobals) GetRealpathCacheSize() ZendLong      { return this.realpath_cache_size }
func (this *VirtualCwdGlobals) SetRealpathCacheSize(value ZendLong) { this.realpath_cache_size = value }
func (this *VirtualCwdGlobals) GetRealpathCacheSizeLimit() ZendLong {
	return this.realpath_cache_size_limit
}
func (this *VirtualCwdGlobals) SetRealpathCacheSizeLimit(value ZendLong) {
	this.realpath_cache_size_limit = value
}
func (this *VirtualCwdGlobals) GetRealpathCacheTtl() ZendLong            { return this.realpath_cache_ttl }
func (this *VirtualCwdGlobals) SetRealpathCacheTtl(value ZendLong)       { this.realpath_cache_ttl = value }
func (this *VirtualCwdGlobals) GetRealpathCache() []*RealpathCacheBucket { return this.realpath_cache }

// func (this *VirtualCwdGlobals) SetRealpathCache(value []*RealpathCacheBucket) { this.realpath_cache = value }
