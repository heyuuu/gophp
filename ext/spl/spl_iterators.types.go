// <<generate>>

package spl

import (
	"sik/zend"
)

/**
 * _spl_cbfilter_it_intern
 */
type _spl_cbfilter_it_intern struct {
	fci    zend.ZendFcallInfo
	fcc    zend.ZendFcallInfoCache
	object *zend.ZendObject
}

// func New_spl_cbfilter_it_intern(fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache, object *zend.ZendObject) *_spl_cbfilter_it_intern {
//     return &_spl_cbfilter_it_intern{
//         fci:fci,
//         fcc:fcc,
//         object:object,
//     }
// }
// func Make_spl_cbfilter_it_intern(fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache, object *zend.ZendObject) _spl_cbfilter_it_intern {
//     return _spl_cbfilter_it_intern{
//         fci:fci,
//         fcc:fcc,
//         object:object,
//     }
// }
func (this *_spl_cbfilter_it_intern) GetFci() zend.ZendFcallInfo { return this.fci }

// func (this *_spl_cbfilter_it_intern) SetFci(value zend.ZendFcallInfo) { this.fci = value }
func (this *_spl_cbfilter_it_intern) GetFcc() zend.ZendFcallInfoCache { return this.fcc }

// func (this *_spl_cbfilter_it_intern) SetFcc(value zend.ZendFcallInfoCache) { this.fcc = value }
func (this *_spl_cbfilter_it_intern) GetObject() *zend.ZendObject      { return this.object }
func (this *_spl_cbfilter_it_intern) SetObject(value *zend.ZendObject) { this.object = value }

/**
 * SplDualItObject
 */
type SplDualItObject struct {
	inner struct {
		zobject  zend.Zval
		ce       *zend.ZendClassEntry
		object   *zend.ZendObject
		iterator *zend.ZendObjectIterator
	}
	current struct {
		data zend.Zval
		key  zend.Zval
		pos  zend.ZendLong
	}
	dit_type DualItType
	u        struct /* union */ {
		limit struct {
			offset zend.ZendLong
			count  zend.ZendLong
		}
		caching struct {
			flags     zend.ZendLong
			zstr      zend.Zval
			zchildren zend.Zval
			zcache    zend.Zval
		}
		append struct {
			zarrayit zend.Zval
			iterator *zend.ZendObjectIterator
		}
		regex struct {
			flags      zend.ZendLong
			preg_flags zend.ZendLong
			pce        *pcre_cache_entry
			regex      *zend.ZendString
			mode       RegexMode
			use_flags  int
		}
		cbfilter *_spl_cbfilter_it_intern
	}
	std zend.ZendObject
}

func (this *SplDualItObject) GetZobject() zend.Zval { return this.inner.zobject }

// func (this *SplDualItObject) SetZobject(value zend.Zval) { this.inner.zobject = value }
func (this *SplDualItObject) GetCe() *zend.ZendClassEntry      { return this.inner.ce }
func (this *SplDualItObject) SetCe(value *zend.ZendClassEntry) { this.inner.ce = value }

// func (this *SplDualItObject)  GetObject() *zend.ZendObject      { return this.inner.object }
func (this *SplDualItObject) SetObject(value *zend.ZendObject)           { this.inner.object = value }
func (this *SplDualItObject) GetInnerIterator() *zend.ZendObjectIterator { return this.inner.iterator }
func (this *SplDualItObject) SetInnerIterator(value *zend.ZendObjectIterator) {
	this.inner.iterator = value
}
func (this *SplDualItObject) GetData() zend.Zval { return this.current.data }

// func (this *SplDualItObject) SetData(value zend.Zval) { this.current.data = value }
func (this *SplDualItObject) GetKey() zend.Zval { return this.current.key }

// func (this *SplDualItObject) SetKey(value zend.Zval) { this.current.key = value }
func (this *SplDualItObject) GetPos() zend.ZendLong                { return this.current.pos }
func (this *SplDualItObject) SetPos(value zend.ZendLong)           { this.current.pos = value }
func (this *SplDualItObject) GetDitType() DualItType               { return this.dit_type }
func (this *SplDualItObject) SetDitType(value DualItType)          { this.dit_type = value }
func (this *SplDualItObject) GetOffset() zend.ZendLong             { return this.u.limit.offset }
func (this *SplDualItObject) SetOffset(value zend.ZendLong)        { this.u.limit.offset = value }
func (this *SplDualItObject) GetCount() zend.ZendLong              { return this.u.limit.count }
func (this *SplDualItObject) SetCount(value zend.ZendLong)         { this.u.limit.count = value }
func (this *SplDualItObject) GetUCachingFlags() zend.ZendLong      { return this.u.caching.flags }
func (this *SplDualItObject) SetUCachingFlags(value zend.ZendLong) { this.u.caching.flags = value }
func (this *SplDualItObject) GetZstr() zend.Zval                   { return this.u.caching.zstr }

// func (this *SplDualItObject) SetZstr(value zend.Zval) { this.u.caching.zstr = value }
func (this *SplDualItObject) GetZchildren() zend.Zval { return this.u.caching.zchildren }

// func (this *SplDualItObject) SetZchildren(value zend.Zval) { this.u.caching.zchildren = value }
func (this *SplDualItObject) GetZcache() zend.Zval { return this.u.caching.zcache }

// func (this *SplDualItObject) SetZcache(value zend.Zval) { this.u.caching.zcache = value }
func (this *SplDualItObject) GetZarrayit() zend.Zval { return this.u.append.zarrayit }

// func (this *SplDualItObject) SetZarrayit(value zend.Zval) { this.u.append.zarrayit = value }
func (this *SplDualItObject) GetUAppendIterator() *zend.ZendObjectIterator {
	return this.u.append.iterator
}
func (this *SplDualItObject) SetUAppendIterator(value *zend.ZendObjectIterator) {
	this.u.append.iterator = value
}
func (this *SplDualItObject) GetURegexFlags() zend.ZendLong              { return this.u.regex.flags }
func (this *SplDualItObject) SetURegexFlags(value zend.ZendLong)         { this.u.regex.flags = value }
func (this *SplDualItObject) GetPregFlags() zend.ZendLong                { return this.u.regex.preg_flags }
func (this *SplDualItObject) SetPregFlags(value zend.ZendLong)           { this.u.regex.preg_flags = value }
func (this *SplDualItObject) GetPce() *pcre_cache_entry                  { return this.u.regex.pce }
func (this *SplDualItObject) SetPce(value *pcre_cache_entry)             { this.u.regex.pce = value }
func (this *SplDualItObject) GetURegexRegex() *zend.ZendString           { return this.u.regex.regex }
func (this *SplDualItObject) SetURegexRegex(value *zend.ZendString)      { this.u.regex.regex = value }
func (this *SplDualItObject) GetMode() RegexMode                         { return this.u.regex.mode }
func (this *SplDualItObject) SetMode(value RegexMode)                    { this.u.regex.mode = value }
func (this *SplDualItObject) GetUseFlags() int                           { return this.u.regex.use_flags }
func (this *SplDualItObject) SetUseFlags(value int)                      { this.u.regex.use_flags = value }
func (this *SplDualItObject) GetCbfilter() *_spl_cbfilter_it_intern      { return this.u.cbfilter }
func (this *SplDualItObject) SetCbfilter(value *_spl_cbfilter_it_intern) { this.u.cbfilter = value }
func (this *SplDualItObject) GetStd() zend.ZendObject                    { return this.std }

// func (this *SplDualItObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplDualItObject.u.caching.flags */
func (this *SplDualItObject) AddUCachingFlags(value zend.ZendLong) { this.u.caching.flags |= value }
func (this *SplDualItObject) SubUCachingFlags(value zend.ZendLong) { this.u.caching.flags &^= value }
func (this *SplDualItObject) HasUCachingFlags(value zend.ZendLong) bool {
	return this.u.caching.flags&value != 0
}
func (this *SplDualItObject) SwitchUCachingFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddUCachingFlags(value)
	} else {
		this.SubUCachingFlags(value)
	}
}
func (this SplDualItObject) IsValid() bool         { return this.HasUCachingFlags(CIT_VALID) }
func (this SplDualItObject) IsFullCache() bool     { return this.HasUCachingFlags(CIT_FULL_CACHE) }
func (this SplDualItObject) IsCatchGetChild() bool { return this.HasUCachingFlags(CIT_CATCH_GET_CHILD) }
func (this SplDualItObject) IsTostringUseInner() bool {
	return this.HasUCachingFlags(CIT_TOSTRING_USE_INNER)
}
func (this SplDualItObject) IsTostringUseKey() bool {
	return this.HasUCachingFlags(CIT_TOSTRING_USE_KEY)
}
func (this SplDualItObject) IsTostringUseCurrent() bool {
	return this.HasUCachingFlags(CIT_TOSTRING_USE_CURRENT)
}
func (this SplDualItObject) IsCallTostring() bool  { return this.HasUCachingFlags(CIT_CALL_TOSTRING) }
func (this *SplDualItObject) SetIsValid(cond bool) { this.SwitchUCachingFlags(CIT_VALID, cond) }
func (this *SplDualItObject) SetIsFullCache(cond bool) {
	this.SwitchUCachingFlags(CIT_FULL_CACHE, cond)
}
func (this *SplDualItObject) SetIsCatchGetChild(cond bool) {
	this.SwitchUCachingFlags(CIT_CATCH_GET_CHILD, cond)
}
func (this *SplDualItObject) SetIsTostringUseInner(cond bool) {
	this.SwitchUCachingFlags(CIT_TOSTRING_USE_INNER, cond)
}
func (this *SplDualItObject) SetIsTostringUseKey(cond bool) {
	this.SwitchUCachingFlags(CIT_TOSTRING_USE_KEY, cond)
}
func (this *SplDualItObject) SetIsTostringUseCurrent(cond bool) {
	this.SwitchUCachingFlags(CIT_TOSTRING_USE_CURRENT, cond)
}
func (this *SplDualItObject) SetIsCallTostring(cond bool) {
	this.SwitchUCachingFlags(CIT_CALL_TOSTRING, cond)
}

/* SplDualItObject.u.regex.flags */
func (this *SplDualItObject) AddURegexFlags(value zend.ZendLong) { this.u.regex.flags |= value }
func (this *SplDualItObject) SubURegexFlags(value zend.ZendLong) { this.u.regex.flags &^= value }
func (this *SplDualItObject) HasURegexFlags(value zend.ZendLong) bool {
	return this.u.regex.flags&value != 0
}
func (this *SplDualItObject) SwitchURegexFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddURegexFlags(value)
	} else {
		this.SubURegexFlags(value)
	}
}
func (this SplDualItObject) IsUseKey() bool           { return this.HasURegexFlags(REGIT_USE_KEY) }
func (this SplDualItObject) IsInverted() bool         { return this.HasURegexFlags(REGIT_INVERTED) }
func (this *SplDualItObject) SetIsUseKey(cond bool)   { this.SwitchURegexFlags(REGIT_USE_KEY, cond) }
func (this *SplDualItObject) SetIsInverted(cond bool) { this.SwitchURegexFlags(REGIT_INVERTED, cond) }

/* SplDualItObject.u.regex.preg_flags */
func (this *SplDualItObject) AddPregFlags(value zend.ZendLong) { this.u.regex.preg_flags |= value }
func (this *SplDualItObject) SubPregFlags(value zend.ZendLong) { this.u.regex.preg_flags &^= value }
func (this *SplDualItObject) HasPregFlags(value zend.ZendLong) bool {
	return this.u.regex.preg_flags&value != 0
}
func (this *SplDualItObject) SwitchPregFlags(value zend.ZendLong, cond bool) {
	if cond {
		this.AddPregFlags(value)
	} else {
		this.SubPregFlags(value)
	}
}

/* SplDualItObject.u.regex.use_flags */
func (this *SplDualItObject) AddUseFlags(value int)      { this.u.regex.use_flags |= value }
func (this *SplDualItObject) SubUseFlags(value int)      { this.u.regex.use_flags &^= value }
func (this *SplDualItObject) HasUseFlags(value int) bool { return this.u.regex.use_flags&value != 0 }
func (this *SplDualItObject) SwitchUseFlags(value int, cond bool) {
	if cond {
		this.AddUseFlags(value)
	} else {
		this.SubUseFlags(value)
	}
}

/**
 * SplSubIterator
 */
type SplSubIterator struct {
	iterator *zend.ZendObjectIterator
	zobject  zend.Zval
	ce       *zend.ZendClassEntry
	state    RecursiveIteratorState
}

// func NewSplSubIterator(iterator *zend.ZendObjectIterator, zobject zend.Zval, ce *zend.ZendClassEntry, state RecursiveIteratorState) *SplSubIterator {
//     return &SplSubIterator{
//         iterator:iterator,
//         zobject:zobject,
//         ce:ce,
//         state:state,
//     }
// }
// func MakeSplSubIterator(iterator *zend.ZendObjectIterator, zobject zend.Zval, ce *zend.ZendClassEntry, state RecursiveIteratorState) SplSubIterator {
//     return SplSubIterator{
//         iterator:iterator,
//         zobject:zobject,
//         ce:ce,
//         state:state,
//     }
// }
func (this *SplSubIterator) GetIterator() *zend.ZendObjectIterator      { return this.iterator }
func (this *SplSubIterator) SetIterator(value *zend.ZendObjectIterator) { this.iterator = value }
func (this *SplSubIterator) GetZobject() zend.Zval                      { return this.zobject }

// func (this *SplSubIterator) SetZobject(value zend.Zval) { this.zobject = value }
func (this *SplSubIterator) GetCe() *zend.ZendClassEntry           { return this.ce }
func (this *SplSubIterator) SetCe(value *zend.ZendClassEntry)      { this.ce = value }
func (this *SplSubIterator) GetState() RecursiveIteratorState      { return this.state }
func (this *SplSubIterator) SetState(value RecursiveIteratorState) { this.state = value }

/**
 * SplRecursiveItObject
 */
type SplRecursiveItObject struct {
	iterators       *SplSubIterator
	level           int
	mode            RecursiveIteratorMode
	flags           int
	max_depth       int
	in_iteration    zend.ZendBool
	beginIteration  *zend.ZendFunction
	endIteration    *zend.ZendFunction
	callHasChildren *zend.ZendFunction
	callGetChildren *zend.ZendFunction
	beginChildren   *zend.ZendFunction
	endChildren     *zend.ZendFunction
	nextElement     *zend.ZendFunction
	ce              *zend.ZendClassEntry
	prefix          []zend.SmartStr
	postfix         []zend.SmartStr
	std             zend.ZendObject
}

// func NewSplRecursiveItObject(iterators *SplSubIterator, level int, mode RecursiveIteratorMode, flags int, max_depth int, in_iteration zend.ZendBool, beginIteration *zend.ZendFunction, endIteration *zend.ZendFunction, callHasChildren *zend.ZendFunction, callGetChildren *zend.ZendFunction, beginChildren *zend.ZendFunction, endChildren *zend.ZendFunction, nextElement *zend.ZendFunction, ce *zend.ZendClassEntry, prefix []zend.SmartStr, postfix []zend.SmartStr, std zend.ZendObject) *SplRecursiveItObject {
//     return &SplRecursiveItObject{
//         iterators:iterators,
//         level:level,
//         mode:mode,
//         flags:flags,
//         max_depth:max_depth,
//         in_iteration:in_iteration,
//         beginIteration:beginIteration,
//         endIteration:endIteration,
//         callHasChildren:callHasChildren,
//         callGetChildren:callGetChildren,
//         beginChildren:beginChildren,
//         endChildren:endChildren,
//         nextElement:nextElement,
//         ce:ce,
//         prefix:prefix,
//         postfix:postfix,
//         std:std,
//     }
// }
// func MakeSplRecursiveItObject(iterators *SplSubIterator, level int, mode RecursiveIteratorMode, flags int, max_depth int, in_iteration zend.ZendBool, beginIteration *zend.ZendFunction, endIteration *zend.ZendFunction, callHasChildren *zend.ZendFunction, callGetChildren *zend.ZendFunction, beginChildren *zend.ZendFunction, endChildren *zend.ZendFunction, nextElement *zend.ZendFunction, ce *zend.ZendClassEntry, prefix []zend.SmartStr, postfix []zend.SmartStr, std zend.ZendObject) SplRecursiveItObject {
//     return SplRecursiveItObject{
//         iterators:iterators,
//         level:level,
//         mode:mode,
//         flags:flags,
//         max_depth:max_depth,
//         in_iteration:in_iteration,
//         beginIteration:beginIteration,
//         endIteration:endIteration,
//         callHasChildren:callHasChildren,
//         callGetChildren:callGetChildren,
//         beginChildren:beginChildren,
//         endChildren:endChildren,
//         nextElement:nextElement,
//         ce:ce,
//         prefix:prefix,
//         postfix:postfix,
//         std:std,
//     }
// }
func (this *SplRecursiveItObject) GetIterators() *SplSubIterator       { return this.iterators }
func (this *SplRecursiveItObject) SetIterators(value *SplSubIterator)  { this.iterators = value }
func (this *SplRecursiveItObject) GetLevel() int                       { return this.level }
func (this *SplRecursiveItObject) SetLevel(value int)                  { this.level = value }
func (this *SplRecursiveItObject) GetMode() RecursiveIteratorMode      { return this.mode }
func (this *SplRecursiveItObject) SetMode(value RecursiveIteratorMode) { this.mode = value }

// func (this *SplRecursiveItObject)  GetFlags() int      { return this.flags }
func (this *SplRecursiveItObject) SetFlags(value int)                    { this.flags = value }
func (this *SplRecursiveItObject) GetMaxDepth() int                      { return this.max_depth }
func (this *SplRecursiveItObject) SetMaxDepth(value int)                 { this.max_depth = value }
func (this *SplRecursiveItObject) GetInIteration() zend.ZendBool         { return this.in_iteration }
func (this *SplRecursiveItObject) SetInIteration(value zend.ZendBool)    { this.in_iteration = value }
func (this *SplRecursiveItObject) GetBeginIteration() *zend.ZendFunction { return this.beginIteration }
func (this *SplRecursiveItObject) SetBeginIteration(value *zend.ZendFunction) {
	this.beginIteration = value
}
func (this *SplRecursiveItObject) GetEndIteration() *zend.ZendFunction { return this.endIteration }
func (this *SplRecursiveItObject) SetEndIteration(value *zend.ZendFunction) {
	this.endIteration = value
}
func (this *SplRecursiveItObject) GetCallHasChildren() *zend.ZendFunction {
	return this.callHasChildren
}
func (this *SplRecursiveItObject) SetCallHasChildren(value *zend.ZendFunction) {
	this.callHasChildren = value
}
func (this *SplRecursiveItObject) GetCallGetChildren() *zend.ZendFunction {
	return this.callGetChildren
}
func (this *SplRecursiveItObject) SetCallGetChildren(value *zend.ZendFunction) {
	this.callGetChildren = value
}
func (this *SplRecursiveItObject) GetBeginChildren() *zend.ZendFunction { return this.beginChildren }
func (this *SplRecursiveItObject) SetBeginChildren(value *zend.ZendFunction) {
	this.beginChildren = value
}
func (this *SplRecursiveItObject) GetEndChildren() *zend.ZendFunction      { return this.endChildren }
func (this *SplRecursiveItObject) SetEndChildren(value *zend.ZendFunction) { this.endChildren = value }
func (this *SplRecursiveItObject) GetNextElement() *zend.ZendFunction      { return this.nextElement }
func (this *SplRecursiveItObject) SetNextElement(value *zend.ZendFunction) { this.nextElement = value }
func (this *SplRecursiveItObject) GetCe() *zend.ZendClassEntry             { return this.ce }
func (this *SplRecursiveItObject) SetCe(value *zend.ZendClassEntry)        { this.ce = value }
func (this *SplRecursiveItObject) GetPrefix() []zend.SmartStr              { return this.prefix }

// func (this *SplRecursiveItObject) SetPrefix(value []zend.SmartStr) { this.prefix = value }
func (this *SplRecursiveItObject) GetPostfix() []zend.SmartStr { return this.postfix }

// func (this *SplRecursiveItObject) SetPostfix(value []zend.SmartStr) { this.postfix = value }
func (this *SplRecursiveItObject) GetStd() zend.ZendObject { return this.std }

// func (this *SplRecursiveItObject) SetStd(value zend.ZendObject) { this.std = value }

/* SplRecursiveItObject.flags */
func (this *SplRecursiveItObject) AddFlags(value int)      { this.flags |= value }
func (this *SplRecursiveItObject) SubFlags(value int)      { this.flags &^= value }
func (this *SplRecursiveItObject) HasFlags(value int) bool { return this.flags&value != 0 }
func (this *SplRecursiveItObject) SwitchFlags(value int, cond bool) {
	if cond {
		this.AddFlags(value)
	} else {
		this.SubFlags(value)
	}
}
func (this SplRecursiveItObject) IsRitCatchGetChild() bool { return this.HasFlags(RIT_CATCH_GET_CHILD) }
func (this SplRecursiveItObject) IsRtitBypassCurrent() bool {
	return this.HasFlags(RTIT_BYPASS_CURRENT)
}
func (this SplRecursiveItObject) IsRtitBypassKey() bool { return this.HasFlags(RTIT_BYPASS_KEY) }
func (this *SplRecursiveItObject) SetIsRitCatchGetChild(cond bool) {
	this.SwitchFlags(RIT_CATCH_GET_CHILD, cond)
}
func (this *SplRecursiveItObject) SetIsRtitBypassCurrent(cond bool) {
	this.SwitchFlags(RTIT_BYPASS_CURRENT, cond)
}
func (this *SplRecursiveItObject) SetIsRtitBypassKey(cond bool) {
	this.SwitchFlags(RTIT_BYPASS_KEY, cond)
}

/**
 * SplRecursiveItIterator
 */
type SplRecursiveItIterator struct {
	intern zend.ZendObjectIterator
}

// func NewSplRecursiveItIterator(intern zend.ZendObjectIterator) *SplRecursiveItIterator {
//     return &SplRecursiveItIterator{
//         intern:intern,
//     }
// }
// func MakeSplRecursiveItIterator(intern zend.ZendObjectIterator) SplRecursiveItIterator {
//     return SplRecursiveItIterator{
//         intern:intern,
//     }
// }
func (this *SplRecursiveItIterator) GetIntern() zend.ZendObjectIterator { return this.intern }

// func (this *SplRecursiveItIterator) SetIntern(value zend.ZendObjectIterator) { this.intern = value }

/**
 * SplIteratorApplyInfo
 */
type SplIteratorApplyInfo struct {
	obj   *zend.Zval
	args  *zend.Zval
	count zend.ZendLong
	fci   zend.ZendFcallInfo
	fcc   zend.ZendFcallInfoCache
}

// func NewSplIteratorApplyInfo(obj *zend.Zval, args *zend.Zval, count zend.ZendLong, fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache) *SplIteratorApplyInfo {
//     return &SplIteratorApplyInfo{
//         obj:obj,
//         args:args,
//         count:count,
//         fci:fci,
//         fcc:fcc,
//     }
// }
// func MakeSplIteratorApplyInfo(obj *zend.Zval, args *zend.Zval, count zend.ZendLong, fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache) SplIteratorApplyInfo {
//     return SplIteratorApplyInfo{
//         obj:obj,
//         args:args,
//         count:count,
//         fci:fci,
//         fcc:fcc,
//     }
// }
func (this *SplIteratorApplyInfo) GetObj() *zend.Zval { return this.obj }

// func (this *SplIteratorApplyInfo) SetObj(value *zend.Zval) { this.obj = value }
func (this *SplIteratorApplyInfo) GetArgs() *zend.Zval          { return this.args }
func (this *SplIteratorApplyInfo) SetArgs(value *zend.Zval)     { this.args = value }
func (this *SplIteratorApplyInfo) GetCount() zend.ZendLong      { return this.count }
func (this *SplIteratorApplyInfo) SetCount(value zend.ZendLong) { this.count = value }
func (this *SplIteratorApplyInfo) GetFci() zend.ZendFcallInfo   { return this.fci }

// func (this *SplIteratorApplyInfo) SetFci(value zend.ZendFcallInfo) { this.fci = value }
func (this *SplIteratorApplyInfo) GetFcc() zend.ZendFcallInfoCache { return this.fcc }

// func (this *SplIteratorApplyInfo) SetFcc(value zend.ZendFcallInfoCache) { this.fcc = value }
