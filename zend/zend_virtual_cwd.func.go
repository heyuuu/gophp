package zend

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
)

func IS_SLASH(c byte) bool                       { return c == '/' }
func COPY_WHEN_ABSOLUTE(path *byte) int          { return 0 }
func IS_ABSOLUTE_PATH(path *byte, len_ int) bool { return IS_SLASH(path[0]) }
func PhpSysReadlink(link *byte, target *byte, target_len int) __auto__ {
	return readlink(link, target, target_len)
}
func CWDG__() *VirtualCwdGlobals                              { return &CwdGlobals }
func VCWD_FOPEN(path *byte, mode string) *r.FILE              { return r.Fopen(path, mode) }
func VCWD_RENAME(oldname *byte, newname *byte) int            { return r.Rename(oldname, newname) }
func VCWD_MKDIR(pathname *byte, mode mode_t) __auto__         { return mkdir(pathname, mode) }
func VCWD_RMDIR(pathname *byte) __auto__                      { return rmdir(pathname) }
func VCWD_UNLINK(path *byte) __auto__                         { return unlink(path) }
func VCWD_CHDIR(path *byte) __auto__                          { return chdir(path) }
func VCWD_ACCESS(pathname *byte, mode __auto__) __auto__      { return access(pathname, mode) }
func VCWD_GETCWD(buff []byte, size int) __auto__              { return getcwd(buff, size) }
func VCWD_CHMOD(path *byte, mode mode_t) __auto__             { return chmod(path, mode) }
func VCWD_CHDIR_FILE(path *byte) int                          { return VirtualChdirFile(path, chdir) }
func VCWD_STAT(path __auto__, buff *ZendStatT) __auto__       { return PhpSysStat(path, buff) }
func VCWD_LSTAT(path *byte, buff *ZendStatT) __auto__         { return lstat(path, buff) }
func VCWD_OPENDIR(pathname *byte) __auto__                    { return opendir(pathname) }
func VCWD_POPEN(command *byte, type_ __auto__) __auto__       { return popen(command, type_) }
func VCWD_REALPATH(path *byte, real_path *byte) *byte         { return TsrmRealpath(path, real_path) }
func VCWD_UTIME(path *byte, time *__struct__utimbuf) __auto__ { return utime(path, time) }
func VCWD_CHOWN(path *byte, owner __auto__, group __auto__) __auto__ {
	return chown(path, owner, group)
}
func VCWD_LCHOWN(path *byte, owner __auto__, group __auto__) __auto__ {
	return lchown(path, owner, group)
}
func S_ISDIR(mode __auto__) bool { return (mode & S_IFMT) == S_IFDIR }
func S_ISREG(mode __auto__) bool { return (mode & S_IFMT) == S_IFREG }
func S_ISLNK(mode __auto__) bool { return (mode & S_IFMT) == S_IFLNK }
func CWD_STATE_COPY(d __auto__, s *CwdState) {
	d.cwd_length = s.GetCwdLength()
	d.cwd = (*byte)(Emalloc(s.GetCwdLength() + 1))
	memcpy(d.cwd, s.GetCwd(), s.GetCwdLength()+1)
}
func CWD_STATE_FREE(s *CwdState) {
	Efree(s.GetCwd())
	s.SetCwdLength(0)
}
func CwdGlobalsCtor(cwd_g *VirtualCwdGlobals) {
	CWD_STATE_COPY(cwd_g.GetCwd(), &MainCwdState)
	cwd_g.SetRealpathCacheSize(0)
	cwd_g.SetRealpathCacheSizeLimit(REALPATH_CACHE_SIZE)
	cwd_g.SetRealpathCacheTtl(REALPATH_CACHE_TTL)
	memset(cwd_g.GetRealpathCache(), 0, b.SizeOf("cwd_g -> realpath_cache"))
}
func RealpathCacheCleanHelper(max_entries uint32, cache **RealpathCacheBucket, cache_size *ZendLong) {
	var i uint32
	for i = 0; i < max_entries; i++ {
		var p *RealpathCacheBucket = cache[i]
		for p != nil {
			var r *RealpathCacheBucket = p
			p = p.GetNext()
			Free(r)
		}
		cache[i] = nil
	}
	*cache_size = 0
}
func CwdGlobalsDtor(cwd_g *VirtualCwdGlobals) {
	RealpathCacheCleanHelper(b.SizeOf("cwd_g -> realpath_cache")/b.SizeOf("cwd_g -> realpath_cache [ 0 ]"), cwd_g.GetRealpathCache(), cwd_g.GetRealpathCacheSize())
}
func VirtualCwdMainCwdInit(reinit uint8) {
	var cwd []byte
	var result *byte
	if reinit != 0 {
		Free(MainCwdState.GetCwd())
	}
	result = getcwd(cwd, b.SizeOf("cwd"))
	if result == nil {
		cwd[0] = '0'
	}
	MainCwdState.SetCwdLength(strlen(cwd))
	MainCwdState.SetCwd(strdup(cwd))
}
func VirtualCwdStartup() {
	VirtualCwdMainCwdInit(0)
	CwdGlobalsCtor(&CwdGlobals)
}
func VirtualCwdShutdown() {
	CwdGlobalsDtor(&CwdGlobals)
	Free(MainCwdState.GetCwd())
}
func VirtualCwdActivate() int {
	if CWDG__().cwd.cwd == nil {
		CWD_STATE_COPY(&(CWDG__().cwd), &MainCwdState)
	}
	return 0
}
func VirtualCwdDeactivate() int {
	if CWDG__().cwd.cwd != nil {
		CWD_STATE_FREE(&(CWDG__().cwd))
		CWDG__().cwd.cwd = nil
	}
	return 0
}
func RealpathCacheKey(path *byte, path_len int) ZendUlong {
	var h ZendUlong
	var e *byte = path + path_len
	for h = uint64(2166136261); path < e; {
		h *= uint64(16777619)
		*path++
		h ^= (*path) - 1
	}
	return h
}
func RealpathCacheClean() {
	RealpathCacheCleanHelper(b.SizeOf("CWDG ( realpath_cache )")/b.SizeOf("CWDG ( realpath_cache ) [ 0 ]"), CWDG__().realpath_cache, &(CWDG__().RealpathCacheSize))
}
func RealpathCacheDel(path *byte, path_len int) {
	var key ZendUlong = RealpathCacheKey(path, path_len)
	var n ZendUlong = key % (b.SizeOf("CWDG ( realpath_cache )") / b.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
	var bucket **RealpathCacheBucket = &CWDG__().realpath_cache[n]
	for (*bucket) != nil {
		if key == bucket.GetKey() && path_len == bucket.GetPathLen() && memcmp(path, bucket.GetPath(), path_len) == 0 {
			var r *RealpathCacheBucket = *bucket
			*bucket = bucket.GetNext()

			/* if the pointers match then only subtract the length of the path */

			if r.GetPath() == r.GetRealpath() {
				CWDG__().RealpathCacheSize -= b.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1
			} else {
				CWDG__().RealpathCacheSize -= b.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1 + r.GetRealpathLen() + 1
			}
			Free(r)
			return
		} else {
			bucket = bucket.GetNext()
		}
	}
}
func RealpathCacheAdd(
	path *byte,
	path_len int,
	realpath *byte,
	realpath_len int,
	is_dir int,
	t int64,
) {
	var size ZendLong = b.SizeOf("realpath_cache_bucket") + path_len + 1
	var same int = 1
	if realpath_len != path_len || memcmp(path, realpath, path_len) != 0 {
		size += realpath_len + 1
		same = 0
	}
	if CWDG__().RealpathCacheSize+size <= CWDG__().realpath_cache_size_limit {
		var bucket *RealpathCacheBucket = Malloc(size)
		var n ZendUlong
		if bucket == nil {
			return
		}
		bucket.SetKey(RealpathCacheKey(path, path_len))
		bucket.SetPath((*byte)(bucket + b.SizeOf("realpath_cache_bucket")))
		memcpy(bucket.GetPath(), path, path_len+1)
		bucket.SetPathLen(path_len)
		if same != 0 {
			bucket.SetRealpath(bucket.GetPath())
		} else {
			bucket.SetRealpath(bucket.GetPath() + (path_len + 1))
			memcpy(bucket.GetRealpath(), realpath, realpath_len+1)
		}
		bucket.SetRealpathLen(realpath_len)
		bucket.SetIsDir(is_dir > 0)
		bucket.SetExpires(t + CWDG__().realpath_cache_ttl)
		n = bucket.GetKey() % (b.SizeOf("CWDG ( realpath_cache )") / b.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
		bucket.SetNext(CWDG__().realpath_cache[n])
		CWDG__().realpath_cache[n] = bucket
		CWDG__().RealpathCacheSize += size
	}
}
func RealpathCacheFind(path *byte, path_len int, t int64) *RealpathCacheBucket {
	var key ZendUlong = RealpathCacheKey(path, path_len)
	var n ZendUlong = key % (b.SizeOf("CWDG ( realpath_cache )") / b.SizeOf("CWDG ( realpath_cache ) [ 0 ]"))
	var bucket **RealpathCacheBucket = &CWDG__().realpath_cache[n]
	for (*bucket) != nil {
		if CWDG__().realpath_cache_ttl && bucket.GetExpires() < t {
			var r *RealpathCacheBucket = *bucket
			*bucket = bucket.GetNext()

			/* if the pointers match then only subtract the length of the path */

			if r.GetPath() == r.GetRealpath() {
				CWDG__().RealpathCacheSize -= b.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1
			} else {
				CWDG__().RealpathCacheSize -= b.SizeOf("realpath_cache_bucket") + r.GetPathLen() + 1 + r.GetRealpathLen() + 1
			}
			Free(r)
		} else if key == bucket.GetKey() && path_len == bucket.GetPathLen() && memcmp(path, bucket.GetPath(), path_len) == 0 {
			return *bucket
		} else {
			bucket = bucket.GetNext()
		}
	}
	return nil
}
func RealpathCacheSize() ZendLong { return CWDG__().RealpathCacheSize }
func RealpathCacheMaxBuckets() ZendLong {
	return b.SizeOf("CWDG ( realpath_cache )") / b.SizeOf("CWDG ( realpath_cache ) [ 0 ]")
}
func RealpathCacheGetBuckets() **RealpathCacheBucket { return CWDG__().realpath_cache }
func TsrmRealpathR(
	path *byte,
	start int,
	len_ int,
	ll *int,
	t *int64,
	use_realpath int,
	is_dir int,
	link_is_dir *int,
) int {
	var i int
	var j int
	var directory int = 0
	var save int
	var st ZendStatT
	var bucket *RealpathCacheBucket
	var tmp *byte
	for true {
		if len_ <= start {
			if link_is_dir != nil {
				*link_is_dir = 1
			}
			return start
		}
		i = len_
		for i > start && !(IS_SLASH(path[i-1])) {
			i--
		}
		b.Assert(i < MAXPATHLEN)
		if i == len_ || i+1 == len_ && path[i] == '.' {

			/* remove double slashes and '.' */

			if i > 0 {
				len_ = i - 1
			} else {
				len_ = 0
			}
			is_dir = 1
			continue
		} else if i+2 == len_ && path[i] == '.' && path[i+1] == '.' {

			/* remove '..' and previous directory */

			is_dir = 1
			if link_is_dir != nil {
				*link_is_dir = 1
			}
			if i <= start+1 {
				if start != 0 {
					return start
				} else {
					return len_
				}
			}
			j = TsrmRealpathR(path, start, i-1, ll, t, use_realpath, 1, nil)
			if j > start && j != size_t-1 {
				j--
				b.Assert(i < MAXPATHLEN)
				for j > start && !(IS_SLASH(path[j])) {
					j--
				}
				b.Assert(i < MAXPATHLEN)
				if start == 0 {

					/* leading '..' must not be removed in case of relative path */

					if j == 0 && path[0] == '.' && path[1] == '.' && IS_SLASH(path[2]) {
						path[3] = '.'
						path[4] = '.'
						path[5] = DEFAULT_SLASH
						j = 5
					} else if j > 0 && path[j+1] == '.' && path[j+2] == '.' && IS_SLASH(path[j+3]) {
						j += 4
						path[b.PostInc(&j)] = '.'
						path[b.PostInc(&j)] = '.'
						path[j] = DEFAULT_SLASH
					}

					/* leading '..' must not be removed in case of relative path */

				}
			} else if start == 0 && j == 0 {

				/* leading '..' must not be removed in case of relative path */

				path[0] = '.'
				path[1] = '.'
				path[2] = DEFAULT_SLASH
				j = 2
			}
			return j
		}
		path[len_] = 0
		save = use_realpath != CWD_EXPAND
		if start != 0 && save != 0 && CWDG__().realpath_cache_size_limit {

			/* cache lookup for absolute path */

			if (*t) == 0 {
				*t = time(0)
			}
			if b.Assign(&bucket, RealpathCacheFind(path, len_, *t)) != nil {
				if is_dir != 0 && bucket.GetIsDir() == 0 {

					/* not a directory */

					return size_t - 1

					/* not a directory */

				} else {
					if link_is_dir != nil {
						*link_is_dir = bucket.GetIsDir()
					}
					memcpy(path, bucket.GetRealpath(), bucket.GetRealpathLen()+1)
					return bucket.GetRealpathLen()
				}
			}
		}
		if save != 0 && PhpSysLstat(path, &st) < 0 {
			if use_realpath == CWD_REALPATH {

				/* file not found */

				return size_t - 1

				/* file not found */

			}

			/* continue resolution anyway but don't save result in the cache */

			save = 0

			/* continue resolution anyway but don't save result in the cache */

		}
		tmp = DoAlloca(len_+1, use_heap)
		memcpy(tmp, path, len_+1)
		if save != 0 && S_ISLNK(st.st_mode) {
			if b.PreInc(&(*ll)) > LINK_MAX || b.Assign(&j, int(PhpSysReadlink(tmp, path, MAXPATHLEN))) == size_t-1 {

				/* too many links or broken symlinks */

				FreeAlloca(tmp, use_heap)
				return size_t - 1
			}
			path[j] = 0
			if IS_ABSOLUTE_PATH(path, j) {
				j = TsrmRealpathR(path, 1, j, ll, t, use_realpath, is_dir, &directory)
				if j == size_t-1 {
					FreeAlloca(tmp, use_heap)
					return size_t - 1
				}
			} else {
				if i+j >= MAXPATHLEN-1 {
					FreeAlloca(tmp, use_heap)
					return size_t - 1
				}
				memmove(path+i, path, j+1)
				memcpy(path, tmp, i-1)
				path[i-1] = DEFAULT_SLASH
				j = TsrmRealpathR(path, start, i+j, ll, t, use_realpath, is_dir, &directory)
				if j == size_t-1 {
					FreeAlloca(tmp, use_heap)
					return size_t - 1
				}
			}
			if link_is_dir != nil {
				*link_is_dir = directory
			}
		} else {
			if save != 0 {
				directory = S_ISDIR(st.st_mode)
				if link_is_dir != nil {
					*link_is_dir = directory
				}
				if is_dir != 0 && directory == 0 {

					/* not a directory */

					FreeAlloca(tmp, use_heap)
					return size_t - 1
				}
			}
			if i <= start+1 {
				j = start
			} else {

				/* some leading directories may be unaccessable */

				j = TsrmRealpathR(path, start, i-1, ll, t, b.Cond(save != 0, CWD_FILEPATH, use_realpath), 1, nil)
				if j > start && j != size_t-1 {
					path[b.PostInc(&j)] = DEFAULT_SLASH
				}
			}
			if j == size_t-1 || j+len_ >= MAXPATHLEN-1+i {
				FreeAlloca(tmp, use_heap)
				return size_t - 1
			}
			memcpy(path+j, tmp+i, len_-i+1)
			j += len_ - i
		}
		if save != 0 && start != 0 && CWDG__().realpath_cache_size_limit {

			/* save absolute path in the cache */

			RealpathCacheAdd(tmp, len_, path, j, directory, *t)

			/* save absolute path in the cache */

		}
		FreeAlloca(tmp, use_heap)
		return j
	}
}
func VirtualFileEx(state *CwdState, path *byte, verify_path VerifyPathFunc, use_realpath int) int {
	var path_length int = strlen(path)
	var resolved_path []byte = []byte{0}
	var start int = 1
	var ll int = 0
	var t int64
	var ret int
	var add_slash int
	var tmp any
	if path_length == 0 || path_length >= MAXPATHLEN-1 {
		errno = EINVAL
		return 1
	}

	/* cwd_length can be 0 when getcwd() fails.
	 * This can happen under solaris when a dir does not have read permissions
	 * but *does* have execute permissions */

	if !(IS_ABSOLUTE_PATH(path, path_length)) {
		if state.GetCwdLength() == 0 {

			/* resolve relative path */

			start = 0
			memcpy(resolved_path, path, path_length+1)
		} else {
			var state_cwd_length int = state.GetCwdLength()
			if path_length+state_cwd_length+1 >= MAXPATHLEN-1 {
				errno = ENAMETOOLONG
				return 1
			}
			memcpy(resolved_path, state.GetCwd(), state_cwd_length)
			if resolved_path[state_cwd_length-1] == DEFAULT_SLASH {
				memcpy(resolved_path+state_cwd_length, path, path_length+1)
				path_length += state_cwd_length
			} else {
				resolved_path[state_cwd_length] = DEFAULT_SLASH
				memcpy(resolved_path+state_cwd_length+1, path, path_length+1)
				path_length += state_cwd_length + 1
			}
		}
	} else {
		memcpy(resolved_path, path, path_length+1)
	}
	add_slash = use_realpath != CWD_REALPATH && path_length > 0 && IS_SLASH(resolved_path[path_length-1])
	if CWDG__().realpath_cache_ttl {
		t = 0
	} else {
		t = -1
	}
	path_length = TsrmRealpathR(resolved_path, start, path_length, &ll, &t, use_realpath, 0, nil)
	if path_length == size_t-1 {
		errno = ENOENT
		return 1
	}
	if start == 0 && path_length == 0 {
		resolved_path[b.PostInc(&path_length)] = '.'
	}
	if add_slash != 0 && path_length != 0 && !(IS_SLASH(resolved_path[path_length-1])) {
		if path_length >= MAXPATHLEN-1 {
			return -1
		}
		resolved_path[b.PostInc(&path_length)] = DEFAULT_SLASH
	}
	resolved_path[path_length] = 0
	if verify_path != nil {
		var old_state CwdState
		CWD_STATE_COPY(&old_state, state)
		state.SetCwdLength(path_length)
		tmp = Erealloc(state.GetCwd(), state.GetCwdLength()+1)
		state.SetCwd((*byte)(tmp))
		memcpy(state.GetCwd(), resolved_path, state.GetCwdLength()+1)
		if verify_path(state) != 0 {
			CWD_STATE_FREE(state)
			*state = old_state
			ret = 1
		} else {
			CWD_STATE_FREE(&old_state)
			ret = 0
		}
	} else {
		state.SetCwdLength(path_length)
		tmp = Erealloc(state.GetCwd(), state.GetCwdLength()+1)
		state.SetCwd((*byte)(tmp))
		memcpy(state.GetCwd(), resolved_path, state.GetCwdLength()+1)
		ret = 0
	}
	return ret
}
func VirtualChdirFile(path *byte, p_chdir func(path *byte) int) int {
	var length int = strlen(path)
	var temp *byte
	var retval int
	if length == 0 {
		return 1
	}
	for b.PreDec(&length) < SIZE_MAX && !(IS_SLASH(path[length])) {

	}
	if length == SIZE_MAX {

		/* No directory only file name */

		errno = ENOENT
		return -1
	}
	if length == COPY_WHEN_ABSOLUTE(path) && IS_ABSOLUTE_PATH(path, length+1) {
		length++
	}
	temp = (*byte)(DoAlloca(length+1, use_heap))
	memcpy(temp, path, length)
	temp[length] = 0
	retval = p_chdir(temp)
	FreeAlloca(temp, use_heap)
	return retval
}
func TsrmRealpath(path *byte, real_path *byte) *byte {
	var new_state CwdState
	var cwd []byte

	/* realpath("") returns CWD */

	if !(*path) {
		new_state.SetCwd((*byte)(Emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
		if VCWD_GETCWD(cwd, MAXPATHLEN) {
			path = cwd
		}
	} else if !(IS_ABSOLUTE_PATH(path, strlen(path))) && VCWD_GETCWD(cwd, MAXPATHLEN) {
		new_state.SetCwd(Estrdup(cwd))
		new_state.SetCwdLength(strlen(cwd))
	} else {
		new_state.SetCwd((*byte)(Emalloc(1)))
		new_state.GetCwd()[0] = '0'
		new_state.SetCwdLength(0)
	}
	if VirtualFileEx(&new_state, path, nil, CWD_REALPATH) != 0 {
		Efree(new_state.GetCwd())
		return nil
	}
	if real_path != nil {
		var copy_len int = b.CondF2(new_state.GetCwdLength() > MAXPATHLEN-1, MAXPATHLEN-1, func() int { return new_state.GetCwdLength() })
		memcpy(real_path, new_state.GetCwd(), copy_len)
		real_path[copy_len] = '0'
		Efree(new_state.GetCwd())
		return real_path
	} else {
		return new_state.GetCwd()
	}
}
