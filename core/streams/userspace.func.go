package streams

import (
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/types"
	"github.com/heyuuu/gophp/zend/zpp"
)

func StreamWrapperDtor(rsrc *types.ZendResource) {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(rsrc.GetPtr())
	zend.Efree(uwrap.GetProtoname())
	zend.Efree(uwrap.GetClassname())
	zend.Efree(uwrap)
}
func UserStreamCreateObject(uwrap *PhpUserStreamWrapper, context *core.PhpStreamContext, object *types.Zval) {
	if uwrap.GetCe().HasCeFlags(zend.AccInterface | zend.AccTrait | zend.AccImplicitAbstractClass | zend.AccExplicitAbstractClass) {
		object.SetUndef()
		return
	}

	/* create an instance of our class */

	if zend.ObjectInitEx(object, uwrap.GetCe()) == types.FAILURE {
		object.SetUndef()
		return
	}
	if context != nil {
		zend.AddPropertyResource(object, "context", context.GetRes())
		context.GetRes().AddRefcount()
	} else {
		zend.AddPropertyNull(object, "context")
	}
	if uwrap.GetCe().GetConstructor() != nil {
		var fci types.ZendFcallInfo
		var fcc types.ZendFcallInfoCache
		var retval types.Zval
		fci.SetSize(b.SizeOf("fci"))
		fci.GetFunctionName().SetUndef()
		fci.SetObject(object.GetObj())
		fci.SetRetval(&retval)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		fcc.SetFunctionHandler(uwrap.GetCe().GetConstructor())
		fcc.SetCalledScope(types.Z_OBJCE_P(object))
		fcc.SetObject(object.GetObj())
		if zend.ZendCallFunction(&fci, &fcc) == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "Could not execute %s::%s()", uwrap.GetCe().GetName().GetVal(), uwrap.GetCe().GetConstructor().GetFunctionName().GetVal())
			zend.ZvalPtrDtor(object)
			object.SetUndef()
		} else {
			zend.ZvalPtrDtor(&retval)
		}
	}
}
func UserWrapperOpener(
	wrapper *core.PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var us *PhpUserstreamDataT
	var zretval types.Zval
	var zfuncname types.Zval
	var args []types.Zval
	var call_result int
	var stream *core.PhpStream = nil
	var old_in_user_include types.ZendBool

	/* Try to catch bad usage without preventing flexibility */

	if standard.FG__().user_stream_current_filename != nil && strcmp(filename, standard.FG__().user_stream_current_filename) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FG__().user_stream_current_filename = filename

	/* if the user stream was registered as local and we are in include context,
	   we add allow_url_include restrictions to allow_url_fopen ones */

	old_in_user_include = core.PG__().in_user_include
	if uwrap.GetWrapper().GetIsUrl() == 0 && (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG__().allow_url_include) {
		core.PG__().in_user_include = 1
	}
	us = zend.Emalloc(b.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, us.GetObject())
	if us.GetObject().IsUndef() {
		standard.FG__().user_stream_current_filename = nil
		core.PG__().in_user_include = old_in_user_include
		zend.Efree(us)
		return nil
	}

	/* call it's stream_open method - set up params first */

	args[0].SetStringVal(b.CastStrAuto(filename))
	args[1].SetStringVal(b.CastStrAuto(mode))
	args[2].SetLong(options)
	args[3].SetNewRef(zend.EG__().GetUninitializedZval())
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_OPEN))

	faults.TryCatch(func() {
		call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &zfuncname, &zretval, 4, args, 0)
	}, func() {
		standard.FG__().user_stream_current_filename = nil
		faults.Bailout()
	})

	if call_result == types.SUCCESS && zretval.IsNotUndef() && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceOps, us, 0, mode)

		/* if the opened path is set, copy it out */

		if args[3].IsReference() && types.Z_REFVAL(args[3]).IsType(types.IS_STRING) && opened_path != nil {
			*opened_path = types.Z_REFVAL(args[3]).GetStr().Copy()
		}

		/* set wrapper data to be a reference to our object */

		types.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+USERSTREAM_OPEN+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(us.GetObject())
		us.GetObject().SetUndef()
		zend.Efree(us)
	}
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[3])
	zend.ZvalPtrDtor(&args[2])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	standard.FG__().user_stream_current_filename = nil
	core.PG__().in_user_include = old_in_user_include
	return stream
}
func UserWrapperOpendir(
	wrapper *core.PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **types.String,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var us *PhpUserstreamDataT
	var zretval types.Zval
	var zfuncname types.Zval
	var args []types.Zval
	var call_result int
	var stream *core.PhpStream = nil

	/* Try to catch bad usage without preventing flexibility */

	if standard.FG__().user_stream_current_filename != nil && strcmp(filename, standard.FG__().user_stream_current_filename) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FG__().user_stream_current_filename = filename
	us = zend.Emalloc(b.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, us.GetObject())
	if us.GetObject().IsUndef() {
		standard.FG__().user_stream_current_filename = nil
		zend.Efree(us)
		return nil
	}

	/* call it's dir_open method - set up params first */

	args[0].SetStringVal(b.CastStrAuto(filename))
	args[1].SetLong(options)
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_DIR_OPEN))
	call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && zretval.IsNotUndef() && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceDirOps, us, 0, mode)

		/* set wrapper data to be a reference to our object */

		types.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+USERSTREAM_DIR_OPEN+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(us.GetObject())
		us.GetObject().SetUndef()
		zend.Efree(us)
	}
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	standard.FG__().user_stream_current_filename = nil
	return stream
}

//@zif -alias stream_register_wrapper
func ZifStreamWrapperRegister(executeData zpp.Ex, return_value zpp.Ret, protocol *types.Zval, classname *types.Zval, _ zpp.Opt, flags *types.Zval) {
	var protocol *types.String
	var classname *types.String
	var uwrap *PhpUserStreamWrapper
	var rsrc *types.ZendResource
	var flags zend.ZendLong = 0
	if zend.ZendParseParameters(executeData.NumArgs(), "SS|l", &protocol, &classname, &flags) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	uwrap = (*PhpUserStreamWrapper)(zend.Ecalloc(1, b.SizeOf("* uwrap")))
	uwrap.SetProtoname(zend.Estrndup(protocol.GetVal(), protocol.GetLen()))
	uwrap.SetClassname(zend.Estrndup(classname.GetVal(), classname.GetLen()))
	uwrap.GetWrapper().SetWops(&UserStreamWops)
	uwrap.GetWrapper().SetAbstract(uwrap)
	uwrap.GetWrapper().SetIsUrl((flags & core.PHP_STREAM_IS_URL) != 0)
	rsrc = zend.ZendRegisterResource(uwrap, LeProtocols)
	if b.Assign(&(uwrap.GetCe()), zend.ZendLookupClass(classname)) != nil {
		if PhpRegisterUrlStreamWrapperVolatile(protocol, uwrap.GetWrapper()) == types.SUCCESS {
			return_value.SetTrue()
			return
		} else {

			/* We failed.  But why? */

			if core.PhpStreamGetUrlStreamWrappersHash().KeyExists(protocol.GetStr()) {
				core.PhpErrorDocref(nil, faults.E_WARNING, "Protocol %s:// is already defined.", protocol.GetVal())
			} else {

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

				core.PhpErrorDocref(nil, faults.E_WARNING, "Invalid protocol scheme specified. Unable to register wrapper class %s to %s://", classname.GetVal(), protocol.GetVal())

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

			}

			/* We failed.  But why? */

		}
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "class '%s' is undefined", classname.GetVal())
	}
	zend.ZendListDelete(rsrc)
	return_value.SetFalse()
	return
}
func ZifStreamWrapperUnregister(executeData zpp.Ex, return_value zpp.Ret, protocol *types.Zval) {
	var protocol *types.String
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &protocol) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	if PhpUnregisterUrlStreamWrapperVolatile(protocol) == types.FAILURE {

		/* We failed */

		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to unregister protocol %s://", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifStreamWrapperRestore(executeData zpp.Ex, return_value zpp.Ret, protocol *types.Zval) {
	var protocol *types.String
	var wrapper *core.PhpStreamWrapper
	var global_wrapper_hash *types.Array
	var wrapper_hash *types.Array
	if zend.ZendParseParameters(executeData.NumArgs(), "S", &protocol) == types.FAILURE {
		return_value.SetFalse()
		return
	}
	global_wrapper_hash = PhpStreamGetUrlStreamWrappersHashGlobal()
	if b.Assign(&wrapper, types.ZendHashFindPtr(global_wrapper_hash, protocol.GetStr())) == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s:// never existed, nothing to restore", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	wrapper_hash = core.PhpStreamGetUrlStreamWrappersHash()
	if wrapper_hash == global_wrapper_hash || types.ZendHashFindPtr(wrapper_hash, protocol.GetStr()) == wrapper {
		core.PhpErrorDocref(nil, faults.E_NOTICE, "%s:// was never changed, nothing to restore", protocol.GetVal())
		return_value.SetTrue()
		return
	}

	/* A failure here could be okay given that the protocol might have been merely unregistered */

	PhpUnregisterUrlStreamWrapperVolatile(protocol)
	if PhpRegisterUrlStreamWrapperVolatile(protocol, wrapper) == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unable to restore original %s:// wrapper", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func PhpUserstreamopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var args []types.Zval
	var didwrite ssize_t
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_WRITE)
	args[0].SetStringVal(b.CastStr((*byte)(buf), count))
	call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		return -1
	}
	if call_result == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsType(types.IS_FALSE) {
			didwrite = -1
		} else {
			zend.ConvertToLong(&retval)
			didwrite = retval.GetLval()
		}
	} else {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_WRITE+" is not implemented!", us.GetWrapper().GetClassname())
		didwrite = -1
	}

	/* don't allow strange buffer overruns due to bogus return */

	if didwrite > 0 && didwrite > count {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_WRITE+" wrote "+zend.ZEND_LONG_FMT+" bytes more data than requested ("+zend.ZEND_LONG_FMT+" written, "+zend.ZEND_LONG_FMT+" max)", us.GetWrapper().GetClassname(), zend_long(didwrite-count), zend.ZendLong(didwrite), zend.ZendLong(count))
		didwrite = count
	}
	zend.ZvalPtrDtor(&retval)
	return didwrite
}
func PhpUserstreamopRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name types.Zval
	var retval types.Zval
	var args []types.Zval
	var call_result int
	var didread int = 0
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_READ)
	args[0].SetLong(count)
	call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		return -1
	}
	if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_READ+" is not implemented!", us.GetWrapper().GetClassname())
		return -1
	}
	if retval.IsType(types.IS_FALSE) {
		return -1
	}
	if zend.TryConvertToString(&retval) == 0 {
		return -1
	}
	didread = retval.GetStr().GetLen()
	if didread > 0 {
		if didread > count {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_READ+" - read "+zend.ZEND_LONG_FMT+" bytes more data than requested ("+zend.ZEND_LONG_FMT+" read, "+zend.ZEND_LONG_FMT+" max) - excess data will be lost", us.GetWrapper().GetClassname(), zend_long(didread-count), zend.ZendLong(didread), zend.ZendLong(count))
			didread = count
		}
		memcpy(buf, retval.GetStr().GetVal(), didread)
	}
	zend.ZvalPtrDtor(&retval)
	retval.SetUndef()

	/* since the user stream has no way of setting the eof flag directly, we need to ask it if we hit eof */

	func_name.SetStringVal(USERSTREAM_EOF)
	call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		stream.SetEof(1)
		return -1
	}
	if call_result == types.SUCCESS && retval.IsNotUndef() && zend.ZvalIsTrue(&retval) != 0 {
		stream.SetEof(1)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		stream.SetEof(1)
	}
	zend.ZvalPtrDtor(&retval)
	return didread
}
func PhpUserstreamopClose(stream *core.PhpStream, close_handle int) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_CLOSE)
	zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(us.GetObject())
	us.GetObject().SetUndef()
	zend.Efree(us)
	return 0
}
func PhpUserstreamopFlush(stream *core.PhpStream) int {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_FLUSH)
	call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsNotUndef() && zend.ZvalIsTrue(&retval) != 0 {
		call_result = 0
	} else {
		call_result = -1
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return call_result
}
func PhpUserstreamopSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var ret int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var args []types.Zval
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_SEEK)
	args[0].SetLong(offset)
	args[1].SetLong(whence)
	call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 2, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&func_name)
	if call_result == types.FAILURE {

		/* stream_seek is not implemented, so disable seeks for this stream */

		stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)

		/* there should be no retval to clean up */

		zend.ZvalPtrDtor(&retval)
		return -1
	} else if call_result == types.SUCCESS && retval.IsNotUndef() && zend.ZvalIsTrue(&retval) != 0 {
		ret = 0
	} else {
		ret = -1
	}
	zend.ZvalPtrDtor(&retval)
	retval.SetUndef()
	if ret != 0 {
		return ret
	}

	/* now determine where we are */

	func_name.SetStringVal(USERSTREAM_TELL)
	call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsType(types.IS_LONG) {
		*newoffs = retval.GetLval()
		ret = 0
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_TELL+" is not implemented!", us.GetWrapper().GetClassname())
		ret = -1
	} else {
		ret = -1
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return ret
}
func StatbufFromArray(array *types.Zval, ssb *core.PhpStreamStatbuf) int {
	var elem *types.Zval

	// #define STAT_PROP_ENTRY_EX(name,name2) if ( NULL != ( elem = zend_hash_str_find ( Z_ARRVAL_P ( array ) , # name , sizeof ( # name ) - 1 ) ) ) { ssb -> sb . st_ ## name2 = zval_get_long ( elem ) ; }

	// #define STAT_PROP_ENTRY(name) STAT_PROP_ENTRY_EX ( name , name )

	memset(ssb, 0, b.SizeOf("php_stream_statbuf"))
	if nil != b.Assign(&elem, array.GetArr().KeyFind("dev")) {
		ssb.GetSb().st_dev = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("ino")) {
		ssb.GetSb().st_ino = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("mode")) {
		ssb.GetSb().st_mode = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("nlink")) {
		ssb.GetSb().st_nlink = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("uid")) {
		ssb.GetSb().st_uid = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("gid")) {
		ssb.GetSb().st_gid = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("rdev")) {
		ssb.GetSb().st_rdev = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("size")) {
		ssb.GetSb().st_size = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("atime")) {
		ssb.GetSb().st_atime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("mtime")) {
		ssb.GetSb().st_mtime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("ctime")) {
		ssb.GetSb().st_ctime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("blksize")) {
		ssb.GetSb().st_blksize = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, array.GetArr().KeyFind("blocks")) {
		ssb.GetSb().st_blocks = zend.ZvalGetLong(elem)
	}
	return types.SUCCESS
}
func PhpUserstreamopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ret int = -1
	func_name.SetStringVal(USERSTREAM_STAT)
	call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsType(types.IS_ARRAY) {
		if types.SUCCESS == StatbufFromArray(&retval, ssb) {
			ret = 0
		}
	} else {
		if call_result == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_STAT+" is not implemented!", us.GetWrapper().GetClassname())
		}
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return ret
}
func PhpUserstreamopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ret int = core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	var args []types.Zval
	switch option {
	case core.PHP_STREAM_OPTION_CHECK_LIVENESS:
		func_name.SetStringVal(USERSTREAM_EOF)
		call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
		if call_result == types.SUCCESS && (retval.IsType(types.IS_FALSE) || retval.IsType(types.IS_TRUE)) {
			if zend.ZvalIsTrue(&retval) != 0 {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			}
		} else {
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
	case core.PHP_STREAM_OPTION_LOCKING:
		args[0].SetLong(0)
		if (value & LOCK_NB) != 0 {
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_NB)
		}
		switch value & ^LOCK_NB {
		case LOCK_SH:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_SH)
		case LOCK_EX:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_EX)
		case LOCK_UN:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_UN)
		}

		/* TODO wouldblock */

		func_name.SetStringVal(USERSTREAM_LOCK)
		call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
		if call_result == types.SUCCESS && (retval.IsType(types.IS_FALSE) || retval.IsType(types.IS_TRUE)) {
			ret = retval.IsType(types.IS_FALSE)
		} else if call_result == types.FAILURE {
			if value == 0 {

				/* lock support test (TODO: more check) */

				ret = core.PHP_STREAM_OPTION_RETURN_OK

				/* lock support test (TODO: more check) */

			} else {
				core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_LOCK+" is not implemented!", us.GetWrapper().GetClassname())
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
		zend.ZvalPtrDtor(&args[0])
	case core.PHP_STREAM_OPTION_TRUNCATE_API:
		func_name.SetStringVal(USERSTREAM_TRUNCATE)
		switch value {
		case core.PHP_STREAM_TRUNCATE_SUPPORTED:
			if zend.ZendIsCallableEx(&func_name, b.CondF2(us.GetObject().IsUndef(), nil, func() *types.ZendObject { return us.GetObject().GetObj() }), zend.IS_CALLABLE_CHECK_SILENT, nil, nil, nil) != 0 {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size >= 0 && new_size <= ptrdiff_t(zend.LONG_MAX) {
				args[0].SetLong(zend.ZendLong(new_size))
				call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
				if call_result == types.SUCCESS && retval.IsNotUndef() {
					if retval.IsType(types.IS_FALSE) || retval.IsType(types.IS_TRUE) {
						if retval.IsType(types.IS_TRUE) {
							ret = core.PHP_STREAM_OPTION_RETURN_OK
						} else {
							ret = core.PHP_STREAM_OPTION_RETURN_ERR
						}
					} else {
						core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" did not return a boolean!", us.GetWrapper().GetClassname())
					}
				} else {
					core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" is not implemented!", us.GetWrapper().GetClassname())
				}
				zend.ZvalPtrDtor(&retval)
				zend.ZvalPtrDtor(&args[0])
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		zend.ZvalPtrDtor(&func_name)
	case core.PHP_STREAM_OPTION_READ_BUFFER:
		fallthrough
	case core.PHP_STREAM_OPTION_WRITE_BUFFER:
		fallthrough
	case core.PHP_STREAM_OPTION_READ_TIMEOUT:
		fallthrough
	case core.PHP_STREAM_OPTION_BLOCKING:
		func_name.SetStringVal(USERSTREAM_SET_OPTION)
		args[0].SetLong(option)
		args[1].SetNull()
		args[2].SetNull()
		switch option {
		case core.PHP_STREAM_OPTION_READ_BUFFER:
			fallthrough
		case core.PHP_STREAM_OPTION_WRITE_BUFFER:
			args[1].SetLong(value)
			if ptrparam {
				args[2].SetLong(*((*long)(ptrparam)))
			} else {
				args[2].SetLong(r.BUFSIZ)
			}
		case core.PHP_STREAM_OPTION_READ_TIMEOUT:
			var tv __struct__timeval = *((*__struct__timeval)(ptrparam))
			args[1].SetLong(tv.tv_sec)
			args[2].SetLong(tv.tv_usec)
		case core.PHP_STREAM_OPTION_BLOCKING:
			args[1].SetLong(value)
		default:

		}
		call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 3, args, 0)
		if call_result == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_SET_OPTION+" is not implemented!", us.GetWrapper().GetClassname())
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
		} else if zend.ZendIsTrue(&retval) != 0 {
			ret = core.PHP_STREAM_OPTION_RETURN_OK
		} else {
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&args[2])
		zend.ZvalPtrDtor(&args[1])
		zend.ZvalPtrDtor(&args[0])
		zend.ZvalPtrDtor(&func_name)
	}
	return ret
}
func UserWrapperUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		return ret
	}

	/* call the unlink method */

	args[0].SetStringVal(b.CastStrAuto(url))
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_UNLINK))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 1, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IS_FALSE) || zretval.IsType(types.IS_TRUE)) {
		ret = zretval.IsType(types.IS_TRUE)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_UNLINK+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[0])
	return ret
}
func UserWrapperRename(wrapper *core.PhpStreamWrapper, url_from *byte, url_to *byte, options int, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		return ret
	}

	/* call the rename method */

	args[0].SetStringVal(b.CastStrAuto(url_from))
	args[1].SetStringVal(b.CastStrAuto(url_to))
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_RENAME))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IS_FALSE) || zretval.IsType(types.IS_TRUE)) {
		ret = zretval.IsType(types.IS_TRUE)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_RENAME+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	return ret
}
func UserWrapperMkdir(wrapper *core.PhpStreamWrapper, url *byte, mode int, options int, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		return ret
	}

	/* call the mkdir method */

	args[0].SetStringVal(b.CastStrAuto(url))
	args[1].SetLong(mode)
	args[2].SetLong(options)
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_MKDIR))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IS_FALSE) || zretval.IsType(types.IS_TRUE)) {
		ret = zretval.IsType(types.IS_TRUE)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_MKDIR+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[2])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	return ret
}
func UserWrapperRmdir(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		return ret
	}

	/* call the rmdir method */

	args[0].SetStringVal(b.CastStrAuto(url))
	args[1].SetLong(options)
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_RMDIR))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IS_FALSE) || zretval.IsType(types.IS_TRUE)) {
		ret = zretval.IsType(types.IS_TRUE)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_RMDIR+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	return ret
}
func UserWrapperMetadata(wrapper *core.PhpStreamWrapper, url *byte, option int, value any, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = 0
	switch option {
	case core.PHP_STREAM_META_TOUCH:
		zend.ArrayInit(&args[2])
		if value {
			var newtime *__struct__utimbuf = (*__struct__utimbuf)(value)
			zend.AddIndexLong(&args[2], 0, newtime.modtime)
			zend.AddIndexLong(&args[2], 1, newtime.actime)
		}
	case core.PHP_STREAM_META_GROUP:
		fallthrough
	case core.PHP_STREAM_META_OWNER:
		fallthrough
	case core.PHP_STREAM_META_ACCESS:
		args[2].SetLong(*((*long)(value)))
	case core.PHP_STREAM_META_GROUP_NAME:
		fallthrough
	case core.PHP_STREAM_META_OWNER_NAME:
		args[2].SetStringVal(b.CastStrAuto(value))
	default:
		core.PhpErrorDocref(nil, faults.E_WARNING, "Unknown option %d for "+USERSTREAM_METADATA, option)
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* call the mkdir method */

	args[0].SetStringVal(b.CastStrAuto(url))
	args[1].SetLong(option)
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_METADATA))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IS_FALSE) || zretval.IsType(types.IS_TRUE)) {
		ret = zretval.IsType(types.IS_TRUE)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_METADATA+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[2])
	return ret
}
func UserWrapperStatUrl(wrapper *core.PhpStreamWrapper, url *byte, flags int, ssb *core.PhpStreamStatbuf, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname types.Zval
	var zretval types.Zval
	var args []types.Zval
	var call_result int
	var object types.Zval
	var ret int = -1

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		return ret
	}

	/* call it's stat_url method - set up params first */

	args[0].SetStringVal(b.CastStrAuto(url))
	args[1].SetLong(flags)
	zfuncname.SetStringVal(b.CastStrAuto(USERSTREAM_STATURL))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && zretval.IsType(types.IS_ARRAY) {

		/* We got the info we needed */

		if types.SUCCESS == StatbufFromArray(&zretval, ssb) {
			ret = 0
		}

		/* We got the info we needed */

	} else {
		if call_result == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_STATURL+" is not implemented!", uwrap.GetClassname())
		}
	}

	/* clean up */

	zend.ZvalPtrDtor(&object)
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	return ret
}
func PhpUserstreamopReaddir(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var didread int = 0
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)

	/* avoid problems if someone mis-uses the stream */

	if count != b.SizeOf("php_stream_dirent") {
		return -1
	}
	func_name.SetStringVal(USERSTREAM_DIR_READ)
	call_result = zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.GetType() != types.IS_FALSE && retval.GetType() != types.IS_TRUE {
		zend.ConvertToString(&retval)
		core.PHP_STRLCPY(ent.GetDName(), retval.GetStr().GetVal(), b.SizeOf("ent -> d_name"), retval.GetStr().GetLen())
		didread = b.SizeOf("php_stream_dirent")
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_DIR_READ+" is not implemented!", us.GetWrapper().GetClassname())
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return didread
}
func PhpUserstreamopClosedir(stream *core.PhpStream, close_handle int) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetStringVal(USERSTREAM_DIR_CLOSE)
	zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(us.GetObject())
	us.GetObject().SetUndef()
	zend.Efree(us)
	return 0
}
func PhpUserstreamopRewinddir(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	func_name.SetStringVal(USERSTREAM_DIR_REWIND)
	zend.CallUserFunction(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return 0
}
func PhpUserstreamopCast(stream *core.PhpStream, castas int, retptr *any) int {
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var func_name types.Zval
	var retval types.Zval
	var args []types.Zval
	var intstream *core.PhpStream = nil
	var call_result int
	var ret int = types.FAILURE
	func_name.SetStringVal(USERSTREAM_CAST)
	switch castas {
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		args[0].SetLong(core.PHP_STREAM_AS_FD_FOR_SELECT)
	default:
		args[0].SetLong(core.PHP_STREAM_AS_STDIO)
	}
	call_result = zend.CallUserFunctionEx(b.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	for {
		if call_result == types.FAILURE {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_CAST+" is not implemented!", us.GetWrapper().GetClassname())
			break
		}
		if zend.ZendIsTrue(&retval) == 0 {
			break
		}
		core.PhpStreamFromZvalNoVerify(intstream, &retval)
		if intstream == nil {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_CAST+" must return a stream resource", us.GetWrapper().GetClassname())
			break
		}
		if intstream == stream {
			core.PhpErrorDocref(nil, faults.E_WARNING, "%s::"+USERSTREAM_CAST+" must not return itself", us.GetWrapper().GetClassname())
			intstream = nil
			break
		}
		ret = core.PhpStreamCast(intstream, castas, retptr, 1)
		break
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&args[0])
	return ret
}
