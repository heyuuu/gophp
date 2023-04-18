package spl

import (
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

/**
 * _spl_cbfilter_it_intern
 */
type _spl_cbfilter_it_intern struct {
	fci    types.ZendFcallInfo
	fcc    types.ZendFcallInfoCache
	object *types.ZendObject
}

// func Make_spl_cbfilter_it_intern(fci zend.ZendFcallInfo, fcc zend.ZendFcallInfoCache, object *zend.ZendObject) _spl_cbfilter_it_intern {
//     return _spl_cbfilter_it_intern{
//         fci:fci,
//         fcc:fcc,
//         object:object,
//     }
// }
func (this *_spl_cbfilter_it_intern) GetFci() types.ZendFcallInfo { return this.fci }

// func (this *_spl_cbfilter_it_intern) SetFci(value zend.ZendFcallInfo) { this.fci = value }
func (this *_spl_cbfilter_it_intern) GetFcc() types.ZendFcallInfoCache { return this.fcc }

// func (this *_spl_cbfilter_it_intern) SetFcc(value zend.ZendFcallInfoCache) { this.fcc = value }
func (this *_spl_cbfilter_it_intern) GetObject() *types.ZendObject      { return this.object }
func (this *_spl_cbfilter_it_intern) SetObject(value *types.ZendObject) { this.object = value }

/**
 * SplDualItObject
 */
type SplDualItObject struct {
	inner struct {
		zobject  types.Zval
		ce       *types.ClassEntry
		object   *types.ZendObject
		iterator *zend.ZendObjectIterator
	}
	current struct {
		data types.Zval
		key  types.Zval
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
			zstr      types.Zval
			zchildren types.Zval
			zcache    types.Zval
		}
		append struct {
			zarrayit types.Zval
			iterator *zend.ZendObjectIterator
		}
		regex struct {
			flags      zend.ZendLong
			preg_flags zend.ZendLong
			pce        *pcre_cache_entry
			regex      *types.String
			mode       RegexMode
			use_flags  int
		}
		cbfilter *_spl_cbfilter_it_intern
	}
	std types.ZendObject
}

func (this *SplDualItObject) GetZobject() types.Zval { return this.inner.zobject }

// func (this *SplDualItObject) SetZobject(value zend.Zval) { this.inner.zobject = value }
func (this *SplDualItObject) GetCe() *types.ClassEntry      { return this.inner.ce }
func (this *SplDualItObject) SetCe(value *types.ClassEntry) { this.inner.ce = value }

// func (this *SplDualItObject)  GetObject() *zend.ZendObject      { return this.inner.object }
func (this *SplDualItObject) SetObject(value *types.ZendObject)          { this.inner.object = value }
func (this *SplDualItObject) GetInnerIterator() *zend.ZendObjectIterator { return this.inner.iterator }
func (this *SplDualItObject) SetInnerIterator(value *zend.ZendObjectIterator) {
	this.inner.iterator = value
}
func (this *SplDualItObject) GetData() types.Zval { return this.current.data }

// func (this *SplDualItObject) SetData(value zend.Zval) { this.current.data = value }
func (this *SplDualItObject) GetKey() types.Zval { return this.current.key }

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
func (this *SplDualItObject) GetZstr() types.Zval                  { return this.u.caching.zstr }

// func (this *SplDualItObject) SetZstr(value zend.Zval) { this.u.caching.zstr = value }
func (this *SplDualItObject) GetZchildren() types.Zval { return this.u.caching.zchildren }

// func (this *SplDualItObject) SetZchildren(value zend.Zval) { this.u.caching.zchildren = value }
func (this *SplDualItObject) GetZcache() types.Zval { return this.u.caching.zcache }

// func (this *SplDualItObject) SetZcache(value zend.Zval) { this.u.caching.zcache = value }
func (this *SplDualItObject) GetZarrayit() types.Zval { return this.u.append.zarrayit }

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
func (this *SplDualItObject) GetURegexRegex() *types.String              { return this.u.regex.regex }
func (this *SplDualItObject) SetURegexRegex(value *types.String)         { this.u.regex.regex = value }
func (this *SplDualItObject) GetMode() RegexMode                         { return this.u.regex.mode }
func (this *SplDualItObject) SetMode(value RegexMode)                    { this.u.regex.mode = value }
func (this *SplDualItObject) GetUseFlags() int                           { return this.u.regex.use_flags }
func (this *SplDualItObject) SetUseFlags(value int)                      { this.u.regex.use_flags = value }
func (this *SplDualItObject) GetCbfilter() *_spl_cbfilter_it_intern      { return this.u.cbfilter }
func (this *SplDualItObject) SetCbfilter(value *_spl_cbfilter_it_intern) { this.u.cbfilter = value }
func (this *SplDualItObject) GetStd() types.ZendObject                   { return this.std }

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
	zobject  types.Zval
	ce       *types.ClassEntry
	state    RecursiveIteratorState
}

// func MakeSplSubIterator(iterator *zend.ZendObjectIterator, zobject zend.Zval, ce *zend.ClassEntry, state RecursiveIteratorState) SplSubIterator {
//     return SplSubIterator{
//         iterator:iterator,
//         zobject:zobject,
//         ce:ce,
//         state:state,
//     }
// }
func (this *SplSubIterator) GetIterator() *zend.ZendObjectIterator      { return this.iterator }
func (this *SplSubIterator) SetIterator(value *zend.ZendObjectIterator) { this.iterator = value }
func (this *SplSubIterator) GetZobject() types.Zval                     { return this.zobject }

// func (this *SplSubIterator) SetZobject(value zend.Zval) { this.zobject = value }
func (this *SplSubIterator) GetCe() *types.ClassEntry              { return this.ce }
func (this *SplSubIterator) SetCe(value *types.ClassEntry)         { this.ce = value }
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
	in_iteration    types.ZendBool
	beginIteration  types.IFunction
	endIteration    types.IFunction
	callHasChildren types.IFunction
	callGetChildren types.IFunction
	beginChildren   types.IFunction
	endChildren     types.IFunction
	nextElement     types.IFunction
	ce              *types.ClassEntry
	prefix          []zend.SmartStr
	postfix         []zend.SmartStr
	std             types.ZendObject
}

//             func MakeSplRecursiveItObject(
// iterators *SplSubIterator,
// level int,
// mode RecursiveIteratorMode,
// flags int,
// max_depth int,
// in_iteration zend.ZendBool,
// beginIteration *zend.ZendFunction,
// endIteration *zend.ZendFunction,
// callHasChildren *zend.ZendFunction,
// callGetChildren *zend.ZendFunction,
// beginChildren *zend.ZendFunction,
// endChildren *zend.ZendFunction,
// nextElement *zend.ZendFunction,
// ce *zend.ClassEntry,
// prefix []zend.SmartStr,
// postfix []zend.SmartStr,
// std zend.ZendObject,
// ) SplRecursiveItObject {
//                 return SplRecursiveItObject{
//                     iterators:iterators,
//                     level:level,
//                     mode:mode,
//                     flags:flags,
//                     max_depth:max_depth,
//                     in_iteration:in_iteration,
//                     beginIteration:beginIteration,
//                     endIteration:endIteration,
//                     callHasChildren:callHasChildren,
//                     callGetChildren:callGetChildren,
//                     beginChildren:beginChildren,
//                     endChildren:endChildren,
//                     nextElement:nextElement,
//                     ce:ce,
//                     prefix:prefix,
//                     postfix:postfix,
//                     std:std,
//                 }
//             }
func (this *SplRecursiveItObject) GetIterators() []SplSubIterator      { return this.iterators }
func (this *SplRecursiveItObject) SetIterators(value *SplSubIterator)  { this.iterators = value }
func (this *SplRecursiveItObject) GetLevel() int                       { return this.level }
func (this *SplRecursiveItObject) SetLevel(value int)                  { this.level = value }
func (this *SplRecursiveItObject) GetMode() RecursiveIteratorMode      { return this.mode }
func (this *SplRecursiveItObject) SetMode(value RecursiveIteratorMode) { this.mode = value }

// func (this *SplRecursiveItObject)  GetFlags() int      { return this.flags }
func (this *SplRecursiveItObject) SetFlags(value int)                  { this.flags = value }
func (this *SplRecursiveItObject) GetMaxDepth() int                    { return this.max_depth }
func (this *SplRecursiveItObject) SetMaxDepth(value int)               { this.max_depth = value }
func (this *SplRecursiveItObject) GetInIteration() types.ZendBool      { return this.in_iteration }
func (this *SplRecursiveItObject) SetInIteration(value types.ZendBool) { this.in_iteration = value }
func (this *SplRecursiveItObject) GetBeginIteration() types.IFunction  { return this.beginIteration }
func (this *SplRecursiveItObject) SetBeginIteration(value types.IFunction) {
	this.beginIteration = value
}
func (this *SplRecursiveItObject) GetEndIteration() types.IFunction { return this.endIteration }
func (this *SplRecursiveItObject) SetEndIteration(value types.IFunction) {
	this.endIteration = value
}
func (this *SplRecursiveItObject) GetCallHasChildren() types.IFunction {
	return this.callHasChildren
}
func (this *SplRecursiveItObject) SetCallHasChildren(value types.IFunction) {
	this.callHasChildren = value
}
func (this *SplRecursiveItObject) GetCallGetChildren() types.IFunction {
	return this.callGetChildren
}
func (this *SplRecursiveItObject) SetCallGetChildren(value types.IFunction) {
	this.callGetChildren = value
}
func (this *SplRecursiveItObject) GetBeginChildren() types.IFunction { return this.beginChildren }
func (this *SplRecursiveItObject) SetBeginChildren(value types.IFunction) {
	this.beginChildren = value
}
func (this *SplRecursiveItObject) GetEndChildren() types.IFunction      { return this.endChildren }
func (this *SplRecursiveItObject) SetEndChildren(value types.IFunction) { this.endChildren = value }
func (this *SplRecursiveItObject) GetNextElement() types.IFunction      { return this.nextElement }
func (this *SplRecursiveItObject) SetNextElement(value types.IFunction) { this.nextElement = value }
func (this *SplRecursiveItObject) GetCe() *types.ClassEntry             { return this.ce }
func (this *SplRecursiveItObject) SetCe(value *types.ClassEntry)        { this.ce = value }
func (this *SplRecursiveItObject) GetPrefix() []zend.SmartStr           { return this.prefix }

// func (this *SplRecursiveItObject) SetPrefix(value []zend.SmartStr) { this.prefix = value }
func (this *SplRecursiveItObject) GetPostfix() []zend.SmartStr { return this.postfix }

// func (this *SplRecursiveItObject) SetPostfix(value []zend.SmartStr) { this.postfix = value }
func (this *SplRecursiveItObject) GetStd() types.ZendObject { return this.std }

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

func (this *SplRecursiveItIterator) GetIntern() zend.ZendObjectIterator { return this.intern }

/**
 * SplIteratorApplyInfo
 */
type SplIteratorApplyInfo struct {
	obj   *types.Zval
	args  *types.Zval
	count zend.ZendLong
	fci   types.ZendFcallInfo
	fcc   types.ZendFcallInfoCache
}

func (this *SplIteratorApplyInfo) GetObj() *types.Zval              { return this.obj }
func (this *SplIteratorApplyInfo) GetArgs() *types.Zval             { return this.args }
func (this *SplIteratorApplyInfo) SetArgs(value *types.Zval)        { this.args = value }
func (this *SplIteratorApplyInfo) GetCount() zend.ZendLong          { return this.count }
func (this *SplIteratorApplyInfo) SetCount(value zend.ZendLong)     { this.count = value }
func (this *SplIteratorApplyInfo) GetFci() types.ZendFcallInfo      { return this.fci }
func (this *SplIteratorApplyInfo) GetFcc() types.ZendFcallInfoCache { return this.fcc }
