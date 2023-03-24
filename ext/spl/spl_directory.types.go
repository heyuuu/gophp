package spl

import (
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

/**
 * SplOtherHandler
 */
type SplOtherHandler struct {
	dtor  SplForeignDtorT
	clone SplForeignCloneT
}

// func MakeSplOtherHandler(dtor SplForeignDtorT, clone SplForeignCloneT) SplOtherHandler {
//     return SplOtherHandler{
//         dtor:dtor,
//         clone:clone,
//     }
// }
func (this *SplOtherHandler) GetDtor() SplForeignDtorT { return this.dtor }

// func (this *SplOtherHandler) SetDtor(value SplForeignDtorT) { this.dtor = value }
func (this *SplOtherHandler) GetClone() SplForeignCloneT { return this.clone }

// func (this *SplOtherHandler) SetClone(value SplForeignCloneT) { this.clone = value }

/**
 * SplFilesystemIterator
 */
type SplFilesystemIterator struct {
	intern  zend.ZendObjectIterator
	current types.Zval
	object  any
}

// func MakeSplFilesystemIterator(intern zend.ZendObjectIterator, current zend.Zval, object any) SplFilesystemIterator {
//     return SplFilesystemIterator{
//         intern:intern,
//         current:current,
//         object:object,
//     }
// }
func (this *SplFilesystemIterator) GetIntern() zend.ZendObjectIterator { return this.intern }

// func (this *SplFilesystemIterator) SetIntern(value zend.ZendObjectIterator) { this.intern = value }
func (this *SplFilesystemIterator) GetCurrent() types.Zval      { return this.current }
func (this *SplFilesystemIterator) SetCurrent(value types.Zval) { this.current = value }
func (this *SplFilesystemIterator) GetObject() any              { return this.object }
func (this *SplFilesystemIterator) SetObject(value any)         { this.object = value }

/**
 * SplFilesystemObject
 */
type SplFilesystemObject struct {
	oth           any
	oth_handler   *SplOtherHandler
	_path         *byte
	_path_len     int
	orig_path     *byte
	file_name     *byte
	file_name_len int
	type_         SPL_FS_OBJ_TYPE
	flags         zend.ZendLong
	file_class    *types.ClassEntry
	info_class    *types.ClassEntry
	u             struct /* union */ {
		dir struct {
			dirp         *core.PhpStream
			entry        core.PhpStreamDirent
			sub_path     *byte
			sub_path_len int
			index        int
			is_recursive int
			func_rewind  *zend.ZendFunction
			func_next    *zend.ZendFunction
			func_valid   *zend.ZendFunction
		}
		file struct {
			stream           *core.PhpStream
			context          *core.PhpStreamContext
			zcontext         *types.Zval
			open_mode        *byte
			open_mode_len    int
			current_zval     types.Zval
			current_line     *byte
			current_line_len int
			max_line_len     int
			current_line_num zend.ZendLong
			zresource        types.Zval
			func_getCurr     *zend.ZendFunction
			delimiter        byte
			enclosure        byte
			escape           int
		}
	}
	std types.ZendObject
}

func (this *SplFilesystemObject) GetOth() any                          { return this.oth }
func (this *SplFilesystemObject) SetOth(value any)                     { this.oth = value }
func (this *SplFilesystemObject) GetOthHandler() *SplOtherHandler      { return this.oth_handler }
func (this *SplFilesystemObject) SetOthHandler(value *SplOtherHandler) { this.oth_handler = value }
func (this *SplFilesystemObject) GetPath() *byte                       { return this._path }
func (this *SplFilesystemObject) SetPath(value *byte)                  { this._path = value }
func (this *SplFilesystemObject) GetPathLen() int                      { return this._path_len }
func (this *SplFilesystemObject) SetPathLen(value int)                 { this._path_len = value }
func (this *SplFilesystemObject) GetOrigPath() *byte                   { return this.orig_path }
func (this *SplFilesystemObject) SetOrigPath(value *byte)              { this.orig_path = value }
func (this *SplFilesystemObject) GetFileName() *byte                   { return this.file_name }
func (this *SplFilesystemObject) SetFileName(value *byte)              { this.file_name = value }
func (this *SplFilesystemObject) GetFileNameLen() int                  { return this.file_name_len }
func (this *SplFilesystemObject) SetFileNameLen(value int)             { this.file_name_len = value }
func (this *SplFilesystemObject) GetType() SPL_FS_OBJ_TYPE             { return this.type_ }
func (this *SplFilesystemObject) SetType(value SPL_FS_OBJ_TYPE)        { this.type_ = value }
func (this *SplFilesystemObject) GetFlags() zend.ZendLong              { return this.flags }
func (this *SplFilesystemObject) SetFlags(value zend.ZendLong)         { this.flags = value }
func (this *SplFilesystemObject) GetFileClass() *types.ClassEntry      { return this.file_class }
func (this *SplFilesystemObject) SetFileClass(value *types.ClassEntry) { this.file_class = value }
func (this *SplFilesystemObject) GetInfoClass() *types.ClassEntry      { return this.info_class }
func (this *SplFilesystemObject) SetInfoClass(value *types.ClassEntry) { this.info_class = value }
func (this *SplFilesystemObject) GetDirp() *core.PhpStream             { return this.u.dir.dirp }
func (this *SplFilesystemObject) SetDirp(value *core.PhpStream)        { this.u.dir.dirp = value }
func (this *SplFilesystemObject) GetEntry() core.PhpStreamDirent       { return this.u.dir.entry }

// func (this *SplFilesystemObject) SetEntry(value core.PhpStreamDirent) { this.u.dir.entry = value }
func (this *SplFilesystemObject) GetSubPath() *byte       { return this.u.dir.sub_path }
func (this *SplFilesystemObject) SetSubPath(value *byte)  { this.u.dir.sub_path = value }
func (this *SplFilesystemObject) GetSubPathLen() int      { return this.u.dir.sub_path_len }
func (this *SplFilesystemObject) SetSubPathLen(value int) { this.u.dir.sub_path_len = value }
func (this *SplFilesystemObject) GetIndex() int           { return this.u.dir.index }
func (this *SplFilesystemObject) SetIndex(value int)      { this.u.dir.index = value }

// func (this *SplFilesystemObject)  GetIsRecursive() int      { return this.u.dir.is_recursive }
func (this *SplFilesystemObject) SetIsRecursive(value int)          { this.u.dir.is_recursive = value }
func (this *SplFilesystemObject) GetFuncRewind() *zend.ZendFunction { return this.u.dir.func_rewind }

// func (this *SplFilesystemObject) SetFuncRewind(value *zend.ZendFunction) { this.u.dir.func_rewind = value }
func (this *SplFilesystemObject) GetFuncNext() *zend.ZendFunction { return this.u.dir.func_next }

// func (this *SplFilesystemObject) SetFuncNext(value *zend.ZendFunction) { this.u.dir.func_next = value }
func (this *SplFilesystemObject) GetFuncValid() *zend.ZendFunction { return this.u.dir.func_valid }

// func (this *SplFilesystemObject) SetFuncValid(value *zend.ZendFunction) { this.u.dir.func_valid = value }
func (this *SplFilesystemObject) GetStream() *core.PhpStream         { return this.u.file.stream }
func (this *SplFilesystemObject) SetStream(value *core.PhpStream)    { this.u.file.stream = value }
func (this *SplFilesystemObject) GetContext() *core.PhpStreamContext { return this.u.file.context }
func (this *SplFilesystemObject) SetContext(value *core.PhpStreamContext) {
	this.u.file.context = value
}
func (this *SplFilesystemObject) GetZcontext() *types.Zval { return this.u.file.zcontext }

// func (this *SplFilesystemObject) SetZcontext(value *zend.Zval) { this.u.file.zcontext = value }
func (this *SplFilesystemObject) GetOpenMode() *byte         { return this.u.file.open_mode }
func (this *SplFilesystemObject) SetOpenMode(value *byte)    { this.u.file.open_mode = value }
func (this *SplFilesystemObject) GetOpenModeLen() int        { return this.u.file.open_mode_len }
func (this *SplFilesystemObject) SetOpenModeLen(value int)   { this.u.file.open_mode_len = value }
func (this *SplFilesystemObject) GetCurrentZval() types.Zval { return this.u.file.current_zval }

// func (this *SplFilesystemObject) SetCurrentZval(value zend.Zval) { this.u.file.current_zval = value }
func (this *SplFilesystemObject) GetCurrentLine() *byte       { return this.u.file.current_line }
func (this *SplFilesystemObject) SetCurrentLine(value *byte)  { this.u.file.current_line = value }
func (this *SplFilesystemObject) GetCurrentLineLen() int      { return this.u.file.current_line_len }
func (this *SplFilesystemObject) SetCurrentLineLen(value int) { this.u.file.current_line_len = value }
func (this *SplFilesystemObject) GetMaxLineLen() int          { return this.u.file.max_line_len }
func (this *SplFilesystemObject) SetMaxLineLen(value int)     { this.u.file.max_line_len = value }
func (this *SplFilesystemObject) GetCurrentLineNum() zend.ZendLong {
	return this.u.file.current_line_num
}
func (this *SplFilesystemObject) SetCurrentLineNum(value zend.ZendLong) {
	this.u.file.current_line_num = value
}
func (this *SplFilesystemObject) GetZresource() types.Zval { return this.u.file.zresource }

// func (this *SplFilesystemObject) SetZresource(value zend.Zval) { this.u.file.zresource = value }
func (this *SplFilesystemObject) GetFuncGetCurr() *zend.ZendFunction { return this.u.file.func_getCurr }
func (this *SplFilesystemObject) SetFuncGetCurr(value *zend.ZendFunction) {
	this.u.file.func_getCurr = value
}
func (this *SplFilesystemObject) GetDelimiter() byte       { return this.u.file.delimiter }
func (this *SplFilesystemObject) SetDelimiter(value byte)  { this.u.file.delimiter = value }
func (this *SplFilesystemObject) GetEnclosure() byte       { return this.u.file.enclosure }
func (this *SplFilesystemObject) SetEnclosure(value byte)  { this.u.file.enclosure = value }
func (this *SplFilesystemObject) GetEscape() int           { return this.u.file.escape }
func (this *SplFilesystemObject) SetEscape(value int)      { this.u.file.escape = value }
func (this *SplFilesystemObject) GetStd() types.ZendObject { return this.std }

// func (this *SplFilesystemObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplFilesystemObject.flags */
func (this *SplFilesystemObject) AddFlags(value zend.ZendLong)      { this.flags |= value }
func (this *SplFilesystemObject) SubFlags(value zend.ZendLong)      { this.flags &^= value }
func (this *SplFilesystemObject) HasFlags(value zend.ZendLong) bool { return this.flags&value != 0 }
func (this *SplFilesystemObject) SwitchFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplFilesystemObject) IsDirFollowSymlinks() bool {
	return this.HasFlags(SPL_FILE_DIR_FOLLOW_SYMLINKS)
}
func (this *SplFilesystemObject) SetIsDirFollowSymlinks(cond bool) {
	this.SwitchFlags(SPL_FILE_DIR_FOLLOW_SYMLINKS, cond)
}
