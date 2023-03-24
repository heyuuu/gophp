package standard

import (
	"sik/zend/def"
	"sik/zend/types"
	"sik/zend/zpp"
)

// generate by ZifUserFilterNop
var DefZifUserFilterNop = def.DefFunc("user_filter_nop", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifUserFilterNop(executeData, returnValue)
})

// generate by ZifStreamBucketMakeWriteable
var DefZifStreamBucketMakeWriteable = def.DefFunc("stream_bucket_make_writeable", 1, 1, []def.ArgInfo{{Name: "brigade"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 1, 1, 0)
	brigade := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketMakeWriteable(executeData, returnValue, brigade)
})

// generate by ZifStreamBucketPrepend
var DefZifStreamBucketPrepend = def.DefFunc("stream_bucket_prepend", 2, 2, []def.ArgInfo{{Name: "brigade"}, {Name: "bucket"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	brigade := fp.ParseZval()
	bucket := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketPrepend(executeData, returnValue, brigade, bucket)
})

// generate by ZifStreamBucketAppend
var DefZifStreamBucketAppend = def.DefFunc("stream_bucket_append", 2, 2, []def.ArgInfo{{Name: "brigade"}, {Name: "bucket"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	brigade := fp.ParseZval()
	bucket := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketAppend(executeData, returnValue, brigade, bucket)
})

// generate by ZifStreamBucketNew
var DefZifStreamBucketNew = def.DefFunc("stream_bucket_new", 2, 2, []def.ArgInfo{{Name: "stream"}, {Name: "buffer"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	stream := fp.ParseZval()
	buffer := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamBucketNew(executeData, returnValue, stream, buffer)
})

// generate by ZifStreamGetFilters
var DefZifStreamGetFilters = def.DefFunc("stream_get_filters", 0, 0, []def.ArgInfo{}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	if !zpp.CheckNumArgsNoneError(executeData) {
		return
	}
	ZifStreamGetFilters(executeData, returnValue)
})

// generate by ZifStreamFilterRegister
var DefZifStreamFilterRegister = def.DefFunc("stream_filter_register", 2, 2, []def.ArgInfo{{Name: "filtername"}, {Name: "classname"}}, func(executeData zpp.DefEx, returnValue zpp.DefReturn) {
	fp := zpp.FastParseStart(executeData, 2, 2, 0)
	filtername := fp.ParseZval()
	classname := fp.ParseZval()
	if fp.HasError() {
		return
	}
	ZifStreamFilterRegister(executeData, returnValue, filtername, classname)
})
