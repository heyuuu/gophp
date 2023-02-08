// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

func StreamWrapperDtor(rsrc *zend.ZendResource) {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(rsrc.GetPtr())
	zend.Efree(uwrap.GetProtoname())
	zend.Efree(uwrap.GetClassname())
	zend.Efree(uwrap)
}
func ZmStartupUserStreams(type_ int, module_number int) int {
	LeProtocols = zend.ZendRegisterListDestructorsEx(StreamWrapperDtor, nil, "stream factory", 0)
	if LeProtocols == zend.FAILURE {
		return zend.FAILURE
	}
	zend.REGISTER_LONG_CONSTANT("STREAM_USE_PATH", core.USE_PATH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_IGNORE_URL", core.IGNORE_URL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_REPORT_ERRORS", core.REPORT_ERRORS, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_MUST_SEEK", core.STREAM_MUST_SEEK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_URL_STAT_LINK", core.PHP_STREAM_URL_STAT_LINK, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_URL_STAT_QUIET", core.PHP_STREAM_URL_STAT_QUIET, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_MKDIR_RECURSIVE", core.PHP_STREAM_MKDIR_RECURSIVE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_IS_URL", core.PHP_STREAM_IS_URL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_OPTION_BLOCKING", core.PHP_STREAM_OPTION_BLOCKING, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_OPTION_READ_TIMEOUT", core.PHP_STREAM_OPTION_READ_TIMEOUT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_OPTION_READ_BUFFER", core.PHP_STREAM_OPTION_READ_BUFFER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_OPTION_WRITE_BUFFER", core.PHP_STREAM_OPTION_WRITE_BUFFER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_BUFFER_NONE", core.PHP_STREAM_BUFFER_NONE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_BUFFER_LINE", core.PHP_STREAM_BUFFER_LINE, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_BUFFER_FULL", core.PHP_STREAM_BUFFER_FULL, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CAST_AS_STREAM", core.PHP_STREAM_AS_STDIO, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_CAST_FOR_SELECT", core.PHP_STREAM_AS_FD_FOR_SELECT, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_TOUCH", core.PHP_STREAM_META_TOUCH, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_OWNER", core.PHP_STREAM_META_OWNER, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_OWNER_NAME", core.PHP_STREAM_META_OWNER_NAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_GROUP", core.PHP_STREAM_META_GROUP, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_GROUP_NAME", core.PHP_STREAM_META_GROUP_NAME, zend.CONST_CS|zend.CONST_PERSISTENT)
	zend.REGISTER_LONG_CONSTANT("STREAM_META_ACCESS", core.PHP_STREAM_META_ACCESS, zend.CONST_CS|zend.CONST_PERSISTENT)
	return zend.SUCCESS
}
func UserStreamCreateObject(uwrap *PhpUserStreamWrapper, context *core.PhpStreamContext, object *zend.Zval) {
	if uwrap.GetCe().HasCeFlags(zend.ZEND_ACC_INTERFACE | zend.ZEND_ACC_TRAIT | zend.ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS) {
		object.SetUndef()
		return
	}

	/* create an instance of our class */

	if zend.ObjectInitEx(object, uwrap.GetCe()) == zend.FAILURE {
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
		var fci zend.ZendFcallInfo
		var fcc zend.ZendFcallInfoCache
		var retval zend.Zval
		fci.SetSize(b.SizeOf("fci"))
		fci.GetFunctionName().SetUndef()
		fci.SetObject(object.GetObj())
		fci.SetRetval(&retval)
		fci.SetParamCount(0)
		fci.SetParams(nil)
		fci.SetNoSeparation(1)
		fcc.SetFunctionHandler(uwrap.GetCe().GetConstructor())
		fcc.SetCalledScope(zend.Z_OBJCE_P(object))
		fcc.SetObject(object.GetObj())
		if zend.ZendCallFunction(&fci, &fcc) == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Could not execute %s::%s()", uwrap.GetCe().GetName().GetVal(), uwrap.GetCe().GetConstructor().GetFunctionName().GetVal())
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
	opened_path **zend.ZendString,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var us *PhpUserstreamDataT
	var zretval zend.Zval
	var zfuncname zend.Zval
	var args []zend.Zval
	var call_result int
	var stream *core.PhpStream = nil
	var old_in_user_include zend.ZendBool

	/* Try to catch bad usage without preventing flexibility */

	if standard.FG(user_stream_current_filename) != nil && strcmp(filename, standard.FG(user_stream_current_filename)) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FG(user_stream_current_filename) = filename

	/* if the user stream was registered as local and we are in include context,
	   we add allow_url_include restrictions to allow_url_fopen ones */

	old_in_user_include = core.PG(in_user_include)
	if uwrap.GetWrapper().GetIsUrl() == 0 && (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG(allow_url_include)) {
		core.PG(in_user_include) = 1
	}
	us = zend.Emalloc(b.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, us.GetObject())
	if us.GetObject().IsType(zend.IS_UNDEF) {
		standard.FG(user_stream_current_filename) = nil
		core.PG(in_user_include) = old_in_user_include
		zend.Efree(us)
		return nil
	}

	/* call it's stream_open method - set up params first */

	zend.ZVAL_STRING(&args[0], filename)
	zend.ZVAL_STRING(&args[1], mode)
	args[2].SetLong(options)
	args[3].SetNewRef(zend.EG__().GetUninitializedZval())
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_OPEN)
	var __orig_bailout *JMP_BUF = zend.EG__().GetBailout()
	var __bailout JMP_BUF
	zend.EG__().SetBailout(&__bailout)
	if zend.SETJMP(__bailout) == 0 {
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &zfuncname, &zretval, 4, args, 0, nil)
	} else {
		zend.EG__().SetBailout(__orig_bailout)
		standard.FG(user_stream_current_filename) = nil
		zend.ZendBailout()
	}
	zend.EG__().SetBailout(__orig_bailout)
	if call_result == zend.SUCCESS && zretval.GetType() != zend.IS_UNDEF && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceOps, us, 0, mode)

		/* if the opened path is set, copy it out */

		if args[3].IsReference() && zend.Z_REFVAL(args[3]).IsType(zend.IS_STRING) && opened_path != nil {
			*opened_path = zend.Z_REFVAL(args[3]).GetStr().Copy()
		}

		/* set wrapper data to be a reference to our object */

		zend.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

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
	standard.FG(user_stream_current_filename) = nil
	core.PG(in_user_include) = old_in_user_include
	return stream
}
func UserWrapperOpendir(
	wrapper *core.PhpStreamWrapper,
	filename *byte,
	mode *byte,
	options int,
	opened_path **zend.ZendString,
	context *core.PhpStreamContext,
) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var us *PhpUserstreamDataT
	var zretval zend.Zval
	var zfuncname zend.Zval
	var args []zend.Zval
	var call_result int
	var stream *core.PhpStream = nil

	/* Try to catch bad usage without preventing flexibility */

	if standard.FG(user_stream_current_filename) != nil && strcmp(filename, standard.FG(user_stream_current_filename)) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FG(user_stream_current_filename) = filename
	us = zend.Emalloc(b.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, us.GetObject())
	if us.GetObject().IsType(zend.IS_UNDEF) {
		standard.FG(user_stream_current_filename) = nil
		zend.Efree(us)
		return nil
	}

	/* call it's dir_open method - set up params first */

	zend.ZVAL_STRING(&args[0], filename)
	args[1].SetLong(options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_DIR_OPEN)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && zretval.GetType() != zend.IS_UNDEF && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceDirOps, us, 0, mode)

		/* set wrapper data to be a reference to our object */

		zend.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

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
	standard.FG(user_stream_current_filename) = nil
	return stream
}
func ZifStreamWrapperRegister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	var classname *zend.ZendString
	var uwrap *PhpUserStreamWrapper
	var rsrc *zend.ZendResource
	var flags zend.ZendLong = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "SS|l", &protocol, &classname, &flags) == zend.FAILURE {
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
		if PhpRegisterUrlStreamWrapperVolatile(protocol, uwrap.GetWrapper()) == zend.SUCCESS {
			return_value.SetTrue()
			return
		} else {

			/* We failed.  But why? */

			if zend.ZendHashExists(core.PhpStreamGetUrlStreamWrappersHash(), protocol) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Protocol %s:// is already defined.", protocol.GetVal())
			} else {

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid protocol scheme specified. Unable to register wrapper class %s to %s://", classname.GetVal(), protocol.GetVal())

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

			}

			/* We failed.  But why? */

		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "class '%s' is undefined", classname.GetVal())
	}
	zend.ZendListDelete(rsrc)
	return_value.SetFalse()
	return
}
func ZifStreamWrapperUnregister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &protocol) == zend.FAILURE {
		return_value.SetFalse()
		return
	}
	if PhpUnregisterUrlStreamWrapperVolatile(protocol) == zend.FAILURE {

		/* We failed */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to unregister protocol %s://", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifStreamWrapperRestore(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	var wrapper *core.PhpStreamWrapper
	var global_wrapper_hash *zend.HashTable
	var wrapper_hash *zend.HashTable
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &protocol) == zend.FAILURE {
		return_value.SetFalse()
		return
	}
	global_wrapper_hash = PhpStreamGetUrlStreamWrappersHashGlobal()
	if b.Assign(&wrapper, zend.ZendHashFindPtr(global_wrapper_hash, protocol)) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s:// never existed, nothing to restore", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	wrapper_hash = core.PhpStreamGetUrlStreamWrappersHash()
	if wrapper_hash == global_wrapper_hash || zend.ZendHashFindPtr(wrapper_hash, protocol) == wrapper {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "%s:// was never changed, nothing to restore", protocol.GetVal())
		return_value.SetTrue()
		return
	}

	/* A failure here could be okay given that the protocol might have been merely unregistered */

	PhpUnregisterUrlStreamWrapperVolatile(protocol)
	if PhpRegisterUrlStreamWrapperVolatile(protocol, wrapper) == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to restore original %s:// wrapper", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func PhpUserstreamopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var args []zend.Zval
	var didwrite ssize_t
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_WRITE, b.SizeOf("USERSTREAM_WRITE")-1)
	zend.ZVAL_STRINGL(&args[0], (*byte)(buf), count)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		return -1
	}
	if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
		if retval.IsType(zend.IS_FALSE) {
			didwrite = -1
		} else {
			zend.ConvertToLong(&retval)
			didwrite = retval.GetLval()
		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_WRITE+" is not implemented!", us.GetWrapper().GetClassname())
		didwrite = -1
	}

	/* don't allow strange buffer overruns due to bogus return */

	if didwrite > 0 && didwrite > count {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_WRITE+" wrote "+zend.ZEND_LONG_FMT+" bytes more data than requested ("+zend.ZEND_LONG_FMT+" written, "+zend.ZEND_LONG_FMT+" max)", us.GetWrapper().GetClassname(), zend_long(didwrite-count), zend.ZendLong(didwrite), zend.ZendLong(count))
		didwrite = count
	}
	zend.ZvalPtrDtor(&retval)
	return didwrite
}
func PhpUserstreamopRead(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	var call_result int
	var didread int = 0
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_READ, b.SizeOf("USERSTREAM_READ")-1)
	args[0].SetLong(count)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		return -1
	}
	if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_READ+" is not implemented!", us.GetWrapper().GetClassname())
		return -1
	}
	if retval.IsType(zend.IS_FALSE) {
		return -1
	}
	if zend.TryConvertToString(&retval) == 0 {
		return -1
	}
	didread = zend.Z_STRLEN(retval)
	if didread > 0 {
		if didread > count {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_READ+" - read "+zend.ZEND_LONG_FMT+" bytes more data than requested ("+zend.ZEND_LONG_FMT+" read, "+zend.ZEND_LONG_FMT+" max) - excess data will be lost", us.GetWrapper().GetClassname(), zend_long(didread-count), zend.ZendLong(didread), zend.ZendLong(count))
			didread = count
		}
		memcpy(buf, zend.Z_STRVAL(retval), didread)
	}
	zend.ZvalPtrDtor(&retval)
	retval.SetUndef()

	/* since the user stream has no way of setting the eof flag directly, we need to ask it if we hit eof */

	zend.ZVAL_STRINGL(&func_name, USERSTREAM_EOF, b.SizeOf("USERSTREAM_EOF")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&func_name)
	if zend.EG__().GetException() != nil {
		stream.SetEof(1)
		return -1
	}
	if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
		stream.SetEof(1)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		stream.SetEof(1)
	}
	zend.ZvalPtrDtor(&retval)
	return didread
}
func PhpUserstreamopClose(stream *core.PhpStream, close_handle int) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_CLOSE, b.SizeOf("USERSTREAM_CLOSE")-1)
	zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(us.GetObject())
	us.GetObject().SetUndef()
	zend.Efree(us)
	return 0
}
func PhpUserstreamopFlush(stream *core.PhpStream) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_FLUSH, b.SizeOf("USERSTREAM_FLUSH")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
		call_result = 0
	} else {
		call_result = -1
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return call_result
}
func PhpUserstreamopSeek(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var ret int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var args []zend.Zval
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_SEEK, b.SizeOf("USERSTREAM_SEEK")-1)
	args[0].SetLong(offset)
	args[1].SetLong(whence)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 2, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&func_name)
	if call_result == zend.FAILURE {

		/* stream_seek is not implemented, so disable seeks for this stream */

		stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)

		/* there should be no retval to clean up */

		zend.ZvalPtrDtor(&retval)
		return -1
	} else if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
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

	zend.ZVAL_STRINGL(&func_name, USERSTREAM_TELL, b.SizeOf("USERSTREAM_TELL")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && retval.IsType(zend.IS_LONG) {
		*newoffs = retval.GetLval()
		ret = 0
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_TELL+" is not implemented!", us.GetWrapper().GetClassname())
		ret = -1
	} else {
		ret = -1
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return ret
}
func StatbufFromArray(array *zend.Zval, ssb *core.PhpStreamStatbuf) int {
	var elem *zend.Zval

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
	return zend.SUCCESS
}
func PhpUserstreamopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ret int = -1
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_STAT, b.SizeOf("USERSTREAM_STAT")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && retval.IsType(zend.IS_ARRAY) {
		if zend.SUCCESS == StatbufFromArray(&retval, ssb) {
			ret = 0
		}
	} else {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_STAT+" is not implemented!", us.GetWrapper().GetClassname())
		}
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return ret
}
func PhpUserstreamopSetOption(stream *core.PhpStream, option int, value int, ptrparam any) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ret int = core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	var args []zend.Zval
	switch option {
	case core.PHP_STREAM_OPTION_CHECK_LIVENESS:
		zend.ZVAL_STRINGL(&func_name, USERSTREAM_EOF, b.SizeOf("USERSTREAM_EOF")-1)
		call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
		if call_result == zend.SUCCESS && (retval.IsType(zend.IS_FALSE) || retval.IsType(zend.IS_TRUE)) {
			if zend.ZvalIsTrue(&retval) != 0 {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			}
		} else {
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
		break
	case core.PHP_STREAM_OPTION_LOCKING:
		args[0].SetLong(0)
		if (value & LOCK_NB) != 0 {
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_NB)
		}
		switch value & ^LOCK_NB {
		case LOCK_SH:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_SH)
			break
		case LOCK_EX:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_EX)
			break
		case LOCK_UN:
			args[0].SetLval(args[0].GetLval() | standard.PHP_LOCK_UN)
			break
		}

		/* TODO wouldblock */

		zend.ZVAL_STRINGL(&func_name, USERSTREAM_LOCK, b.SizeOf("USERSTREAM_LOCK")-1)
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0, nil)
		if call_result == zend.SUCCESS && (retval.IsType(zend.IS_FALSE) || retval.IsType(zend.IS_TRUE)) {
			ret = retval.IsType(zend.IS_FALSE)
		} else if call_result == zend.FAILURE {
			if value == 0 {

				/* lock support test (TODO: more check) */

				ret = core.PHP_STREAM_OPTION_RETURN_OK

				/* lock support test (TODO: more check) */

			} else {
				core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_LOCK+" is not implemented!", us.GetWrapper().GetClassname())
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
		zend.ZvalPtrDtor(&args[0])
		break
	case core.PHP_STREAM_OPTION_TRUNCATE_API:
		zend.ZVAL_STRINGL(&func_name, USERSTREAM_TRUNCATE, b.SizeOf("USERSTREAM_TRUNCATE")-1)
		switch value {
		case core.PHP_STREAM_TRUNCATE_SUPPORTED:
			if zend.ZendIsCallableEx(&func_name, b.CondF2(us.GetObject().IsUndef(), nil, func() *zend.ZendObject { return us.GetObject().GetObj() }), zend.IS_CALLABLE_CHECK_SILENT, nil, nil, nil) != 0 {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
			break
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size >= 0 && new_size <= ptrdiff_t(zend.LONG_MAX) {
				args[0].SetLong(zend.ZendLong(new_size))
				call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0, nil)
				if call_result == zend.SUCCESS && retval.GetType() != zend.IS_UNDEF {
					if retval.IsType(zend.IS_FALSE) || retval.IsType(zend.IS_TRUE) {
						if retval.IsType(zend.IS_TRUE) {
							ret = core.PHP_STREAM_OPTION_RETURN_OK
						} else {
							ret = core.PHP_STREAM_OPTION_RETURN_ERR
						}
					} else {
						core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" did not return a boolean!", us.GetWrapper().GetClassname())
					}
				} else {
					core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" is not implemented!", us.GetWrapper().GetClassname())
				}
				zend.ZvalPtrDtor(&retval)
				zend.ZvalPtrDtor(&args[0])
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
			break
		}
		zend.ZvalPtrDtor(&func_name)
		break
	case core.PHP_STREAM_OPTION_READ_BUFFER:

	case core.PHP_STREAM_OPTION_WRITE_BUFFER:

	case core.PHP_STREAM_OPTION_READ_TIMEOUT:

	case core.PHP_STREAM_OPTION_BLOCKING:
		zend.ZVAL_STRINGL(&func_name, USERSTREAM_SET_OPTION, b.SizeOf("USERSTREAM_SET_OPTION")-1)
		args[0].SetLong(option)
		args[1].SetNull()
		args[2].SetNull()
		switch option {
		case core.PHP_STREAM_OPTION_READ_BUFFER:

		case core.PHP_STREAM_OPTION_WRITE_BUFFER:
			args[1].SetLong(value)
			if ptrparam {
				args[2].SetLong(*((*long)(ptrparam)))
			} else {
				args[2].SetLong(r.BUFSIZ)
			}
			break
		case core.PHP_STREAM_OPTION_READ_TIMEOUT:
			var tv __struct__timeval = *((*__struct__timeval)(ptrparam))
			args[1].SetLong(tv.tv_sec)
			args[2].SetLong(tv.tv_usec)
			break
		case core.PHP_STREAM_OPTION_BLOCKING:
			args[1].SetLong(value)
			break
		default:
			break
		}
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 3, args, 0, nil)
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_SET_OPTION+" is not implemented!", us.GetWrapper().GetClassname())
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
		break
	}
	return ret
}
func UserWrapperUnlink(wrapper *core.PhpStreamWrapper, url *byte, options int, context *core.PhpStreamContext) int {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.GetAbstract())
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		return ret
	}

	/* call the unlink method */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_UNLINK)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 1, args, 0, nil)
	if call_result == zend.SUCCESS && (zretval.IsType(zend.IS_FALSE) || zretval.IsType(zend.IS_TRUE)) {
		ret = zretval.IsType(zend.IS_TRUE)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_UNLINK+" is not implemented!", uwrap.GetClassname())
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
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		return ret
	}

	/* call the rename method */

	zend.ZVAL_STRING(&args[0], url_from)
	zend.ZVAL_STRING(&args[1], url_to)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_RENAME)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && (zretval.IsType(zend.IS_FALSE) || zretval.IsType(zend.IS_TRUE)) {
		ret = zretval.IsType(zend.IS_TRUE)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_RENAME+" is not implemented!", uwrap.GetClassname())
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
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		return ret
	}

	/* call the mkdir method */

	zend.ZVAL_STRING(&args[0], url)
	args[1].SetLong(mode)
	args[2].SetLong(options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_MKDIR)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 3, args, 0, nil)
	if call_result == zend.SUCCESS && (zretval.IsType(zend.IS_FALSE) || zretval.IsType(zend.IS_TRUE)) {
		ret = zretval.IsType(zend.IS_TRUE)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_MKDIR+" is not implemented!", uwrap.GetClassname())
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
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		return ret
	}

	/* call the rmdir method */

	zend.ZVAL_STRING(&args[0], url)
	args[1].SetLong(options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_RMDIR)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && (zretval.IsType(zend.IS_FALSE) || zretval.IsType(zend.IS_TRUE)) {
		ret = zretval.IsType(zend.IS_TRUE)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_RMDIR+" is not implemented!", uwrap.GetClassname())
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
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0
	switch option {
	case core.PHP_STREAM_META_TOUCH:
		zend.ArrayInit(&args[2])
		if value {
			var newtime *__struct__utimbuf = (*__struct__utimbuf)(value)
			zend.AddIndexLong(&args[2], 0, newtime.modtime)
			zend.AddIndexLong(&args[2], 1, newtime.actime)
		}
		break
	case core.PHP_STREAM_META_GROUP:

	case core.PHP_STREAM_META_OWNER:

	case core.PHP_STREAM_META_ACCESS:
		args[2].SetLong(*((*long)(value)))
		break
	case core.PHP_STREAM_META_GROUP_NAME:

	case core.PHP_STREAM_META_OWNER_NAME:
		zend.ZVAL_STRING(&args[2], value)
		break
	default:
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unknown option %d for "+USERSTREAM_METADATA, option)
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* call the mkdir method */

	zend.ZVAL_STRING(&args[0], url)
	args[1].SetLong(option)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_METADATA)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 3, args, 0, nil)
	if call_result == zend.SUCCESS && (zretval.IsType(zend.IS_FALSE) || zretval.IsType(zend.IS_TRUE)) {
		ret = zretval.IsType(zend.IS_TRUE)
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_METADATA+" is not implemented!", uwrap.GetClassname())
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
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = -1

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsType(zend.IS_UNDEF) {
		return ret
	}

	/* call it's stat_url method - set up params first */

	zend.ZVAL_STRING(&args[0], url)
	args[1].SetLong(flags)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_STATURL)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && zretval.IsType(zend.IS_ARRAY) {

		/* We got the info we needed */

		if zend.SUCCESS == StatbufFromArray(&zretval, ssb) {
			ret = 0
		}

		/* We got the info we needed */

	} else {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_STATURL+" is not implemented!", uwrap.GetClassname())
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
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var didread int = 0
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)

	/* avoid problems if someone mis-uses the stream */

	if count != b.SizeOf("php_stream_dirent") {
		return -1
	}
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_READ, b.SizeOf("USERSTREAM_DIR_READ")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && retval.GetType() != zend.IS_FALSE && retval.GetType() != zend.IS_TRUE {
		zend.ConvertToString(&retval)
		core.PHP_STRLCPY(ent.GetDName(), zend.Z_STRVAL(retval), b.SizeOf("ent -> d_name"), zend.Z_STRLEN(retval))
		didread = b.SizeOf("php_stream_dirent")
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_DIR_READ+" is not implemented!", us.GetWrapper().GetClassname())
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return didread
}
func PhpUserstreamopClosedir(stream *core.PhpStream, close_handle int) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_CLOSE, b.SizeOf("USERSTREAM_DIR_CLOSE")-1)
	zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(us.GetObject())
	us.GetObject().SetUndef()
	zend.Efree(us)
	return 0
}
func PhpUserstreamopRewinddir(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_REWIND, b.SizeOf("USERSTREAM_DIR_REWIND")-1)
	zend.CallUserFunction(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return 0
}
func PhpUserstreamopCast(stream *core.PhpStream, castas int, retptr *any) int {
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var func_name zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	var intstream *core.PhpStream = nil
	var call_result int
	var ret int = zend.FAILURE
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_CAST, b.SizeOf("USERSTREAM_CAST")-1)
	switch castas {
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		args[0].SetLong(core.PHP_STREAM_AS_FD_FOR_SELECT)
		break
	default:
		args[0].SetLong(core.PHP_STREAM_AS_STDIO)
		break
	}
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(us.GetObject().IsUndef(), nil, func() zend.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0, nil)
	for {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_CAST+" is not implemented!", us.GetWrapper().GetClassname())
			break
		}
		if zend.ZendIsTrue(&retval) == 0 {
			break
		}
		core.PhpStreamFromZvalNoVerify(intstream, &retval)
		if intstream == nil {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_CAST+" must return a stream resource", us.GetWrapper().GetClassname())
			break
		}
		if intstream == stream {
			core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_CAST+" must not return itself", us.GetWrapper().GetClassname())
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
