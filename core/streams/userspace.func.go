package streams

import (
	"fmt"
	b "github.com/heyuuu/gophp/builtin"
	r "github.com/heyuuu/gophp/builtin/file"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/ext/standard"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/operators"
	"github.com/heyuuu/gophp/zend/zpp"
)

func StreamWrapperDtor(rsrc *types.Resource) {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(rsrc.GetPtr())
	zend.Efree(uwrap.GetProtoname())
	zend.Efree(uwrap.GetClassname())
	zend.Efree(uwrap)
}
func UserStreamCreateObject(uwrap *PhpUserStreamWrapper, context *core.PhpStreamContext, object *types.Zval) {
	if uwrap.GetCe().HasCeFlags(types.AccInterface | types.AccTrait | types.AccImplicitAbstractClass | types.AccExplicitAbstractClass) {
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
		//context.GetRes().AddRefcount()
	} else {
		zend.AddPropertyNull(object, "context")
	}
	if uwrap.GetCe().GetConstructor() != nil {
		var fci = types.InitFCallInfo(object.Object(), nil)

		var fcc types.ZendFcallInfoCache
		fcc.SetFunctionHandler(uwrap.GetCe().GetConstructor())
		fcc.SetCalledScope(types.Z_OBJCE_P(object))
		fcc.SetObject(object.Object())
		if zend.ZendCallFunction(fci, &fcc) == types.FAILURE {
			core.PhpErrorDocref("", faults.E_WARNING, "Could not execute %s::%s()", uwrap.GetCe().Name(), uwrap.GetCe().GetConstructor().FunctionName())
			object.SetUndef()
		} else {
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
	var old_in_user_include bool

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

	args[0].SetString(b.CastStrAuto(filename))
	args[1].SetString(b.CastStrAuto(mode))
	args[2].SetLong(options)
	args[3].SetNewRef(zend.UninitializedZval())
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_OPEN))

	faults.TryCatch(func() {
		call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &zfuncname, &zretval, 4, args, 0)
	}, func() {
		standard.FG__().user_stream_current_filename = nil
		faults.Bailout()
	})

	if call_result == types.SUCCESS && zretval.IsNotUndef() && operators.ZvalIsTrue(&zretval) {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceOps, us, 0, mode)

		/* if the opened path is set, copy it out */

		if args[3].IsRef() && types.Z_REFVAL(args[3]).IsString() && opened_path != nil {
			*opened_path = types.Z_REFVAL(args[3]).StringEx().Copy()
		}

		/* set wrapper data to be a reference to our object */

		types.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, fmt.Sprintf(`"%s::%s" call failed`, USERSTREAM_OPEN, us.GetWrapper().GetClassname()))
	}

	/* destroy everything else */

	if stream == nil {
		// zend.ZvalPtrDtor(us.GetObject())
		us.GetObject().SetUndef()
		zend.Efree(us)
	}
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[3])
	// zend.ZvalPtrDtor(&args[2])
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
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

	args[0].SetString(b.CastStrAuto(filename))
	args[1].SetLong(options)
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_DIR_OPEN))
	call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && zretval.IsNotUndef() && operators.ZvalIsTrue(&zretval) {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceDirOps, us, 0, mode)

		/* set wrapper data to be a reference to our object */

		types.ZVAL_COPY(stream.GetWrapperdata(), us.GetObject())

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, fmt.Sprintf(`"%s::%s" call failed`, USERSTREAM_DIR_OPEN, us.GetWrapper().GetClassname()))
	}

	/* destroy everything else */

	if stream == nil {
		// zend.ZvalPtrDtor(us.GetObject())
		us.GetObject().SetUndef()
		zend.Efree(us)
	}
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
	standard.FG__().user_stream_current_filename = nil
	return stream
}

//@zif -alias stream_register_wrapper
func ZifStreamWrapperRegister(executeData zpp.Ex, return_value zpp.Ret, protocol *types.Zval, classname *types.Zval, _ zpp.Opt, flags *types.Zval) {
	var protocol *types.String
	var classname *types.String
	var uwrap *PhpUserStreamWrapper
	var rsrc *types.Resource
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
	if lang.Assign(&(uwrap.GetCe()), zend.ZendLookupClass(classname.String())) != nil {
		if PhpRegisterUrlStreamWrapperVolatile(protocol, uwrap.GetWrapper()) == types.SUCCESS {
			return_value.SetTrue()
			return
		} else {

			/* We failed.  But why? */

			if core.PhpStreamGetUrlStreamWrappersHash().KeyExists(protocol.StringEx()) {
				core.PhpErrorDocref("", faults.E_WARNING, "Protocol %s:// is already defined.", protocol.GetVal())
			} else {

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

				core.PhpErrorDocref("", faults.E_WARNING, "Invalid protocol scheme specified. Unable to register wrapper class %s to %s://", classname.GetVal(), protocol.GetVal())

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

			}

			/* We failed.  But why? */

		}
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "class '%s' is undefined", classname.GetVal())
	}
	//zend.ZendListDelete(rsrc)
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

		core.PhpErrorDocref("", faults.E_WARNING, "Unable to unregister protocol %s://", protocol.GetVal())
		return_value.SetFalse()
		return
	}
	return_value.SetTrue()
	return
}
func ZifStreamWrapperRestore(protocol string) bool {
	var wrapper_hash map[string]*core.PhpStreamWrapper
	global_wrapper_hash := UrlStreamWrappersHash
	wrapper := global_wrapper_hash[protocol]
	if wrapper == nil {
		core.PhpErrorDocref("", faults.E_WARNING, "%s:// never existed, nothing to restore", protocol)
		return false
	}

	wrapper_hash = core.PhpStreamGetUrlStreamWrappersHash()
	if wrapper_hash == global_wrapper_hash || wrapper_hash[protocol] == wrapper {
		core.PhpErrorDocref("", faults.E_NOTICE, "%s:// was never changed, nothing to restore", protocol)
		return false
	}

	/* A failure here could be okay given that the protocol might have been merely unregistered */

	PhpUnregisterUrlStreamWrapperVolatile(protocol)
	if PhpRegisterUrlStreamWrapperVolatile(protocol, wrapper) == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "Unable to restore original %s:// wrapper", protocol)
		return false
	}
	return true
}
func PhpUserstreamopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var args []types.Zval
	var didwrite ssize_t
	b.Assert(us != nil)
	func_name.SetString(USERSTREAM_WRITE)
	args[0].SetString(b.CastStr((*byte)(buf), count))
	call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	// zend.ZvalPtrDtor(&args[0])
	// zend.ZvalPtrDtor(&func_name)
	if zend.EG__().HasException() {
		return -1
	}
	if call_result == types.SUCCESS && retval.IsNotUndef() {
		if retval.IsType(types.IsFalse) {
			didwrite = -1
		} else {
			operators.ConvertToLong(&retval)
			didwrite = retval.Long()
		}
	} else {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_WRITE+" is not implemented!", us.GetWrapper().GetClassname())
		didwrite = -1
	}

	/* don't allow strange buffer overruns due to bogus return */

	if didwrite > 0 && didwrite > count {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_WRITE+" wrote "+"%d"+" bytes more data than requested (%d written, %d max)", us.GetWrapper().GetClassname(), zend_long(didwrite-count), zend.ZendLong(didwrite), zend.ZendLong(count))
		didwrite = count
	}
	// zend.ZvalPtrDtor(&retval)
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
	func_name.SetString(USERSTREAM_READ)
	args[0].SetLong(count)
	call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	// zend.ZvalPtrDtor(&args[0])
	// zend.ZvalPtrDtor(&func_name)
	if zend.EG__().HasException() {
		return -1
	}
	if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_READ+" is not implemented!", us.GetWrapper().GetClassname())
		return -1
	}
	if retval.IsType(types.IsFalse) {
		return -1
	}
	if operators.TryConvertToString(&retval) == 0 {
		return -1
	}
	didread = retval.StringEx().GetLen()
	if didread > 0 {
		if didread > count {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_READ+" - read "+"%d"+" bytes more data than requested (%d read, %d max) - excess data will be lost", us.GetWrapper().GetClassname(), zend_long(didread-count), zend.ZendLong(didread), zend.ZendLong(count))
			didread = count
		}
		memcpy(buf, retval.StringEx().GetVal(), didread)
	}
	// zend.ZvalPtrDtor(&retval)
	retval.SetUndef()

	/* since the user stream has no way of setting the eof flag directly, we need to ask it if we hit eof */

	func_name.SetString(USERSTREAM_EOF)
	call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	// zend.ZvalPtrDtor(&func_name)
	if zend.EG__().HasException() {
		stream.SetEof(1)
		return -1
	}
	if call_result == types.SUCCESS && retval.IsNotUndef() && operators.ZvalIsTrue(&retval) {
		stream.SetEof(1)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		stream.SetEof(1)
	}
	// zend.ZvalPtrDtor(&retval)
	return didread
}
func PhpUserstreamopClose(stream *core.PhpStream, close_handle int) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetString(USERSTREAM_CLOSE)
	zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
	// zend.ZvalPtrDtor(us.GetObject())
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
	func_name.SetString(USERSTREAM_FLUSH)
	call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsNotUndef() && operators.ZvalIsTrue(&retval) {
		call_result = 0
	} else {
		call_result = -1
	}
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
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
	func_name.SetString(USERSTREAM_SEEK)
	args[0].SetLong(offset)
	args[1].SetLong(whence)
	call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 2, args, 0)
	// zend.ZvalPtrDtor(&args[0])
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&func_name)
	if call_result == types.FAILURE {

		/* stream_seek is not implemented, so disable seeks for this stream */

		stream.AddFlags(core.PHP_STREAM_FLAG_NO_SEEK)

		/* there should be no retval to clean up */

		// zend.ZvalPtrDtor(&retval)
		return -1
	} else if call_result == types.SUCCESS && retval.IsNotUndef() && operators.ZvalIsTrue(&retval) {
		ret = 0
	} else {
		ret = -1
	}
	// zend.ZvalPtrDtor(&retval)
	retval.SetUndef()
	if ret != 0 {
		return ret
	}

	/* now determine where we are */

	func_name.SetString(USERSTREAM_TELL)
	call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsType(types.IsLong) {
		*newoffs = retval.Long()
		ret = 0
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_TELL+" is not implemented!", us.GetWrapper().GetClassname())
		ret = -1
	} else {
		ret = -1
	}
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
	return ret
}
func StatbufFromArray(array *types.Zval, ssb *core.PhpStreamStatbuf) int {
	var elem *types.Zval

	// #define STAT_PROP_ENTRY_EX(name,name2) if ( NULL != ( elem = zend_hash_str_find ( Z_ARRVAL_P ( array ) , # name , sizeof ( # name ) - 1 ) ) ) { ssb -> sb . st_ ## name2 = zval_get_long ( elem ) ; }

	// #define STAT_PROP_ENTRY(name) STAT_PROP_ENTRY_EX ( name , name )

	memset(ssb, 0, b.SizeOf("php_stream_statbuf"))
	if nil != lang.Assign(&elem, array.Array().KeyFind("dev")) {
		ssb.GetSb().st_dev = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("ino")) {
		ssb.GetSb().st_ino = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("mode")) {
		ssb.GetSb().st_mode = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("nlink")) {
		ssb.GetSb().st_nlink = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("uid")) {
		ssb.GetSb().st_uid = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("gid")) {
		ssb.GetSb().st_gid = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("rdev")) {
		ssb.GetSb().st_rdev = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("size")) {
		ssb.GetSb().st_size = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("atime")) {
		ssb.GetSb().st_atime = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("mtime")) {
		ssb.GetSb().st_mtime = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("ctime")) {
		ssb.GetSb().st_ctime = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("blksize")) {
		ssb.GetSb().st_blksize = operators.ZvalGetLong(elem)
	}
	if nil != lang.Assign(&elem, array.Array().KeyFind("blocks")) {
		ssb.GetSb().st_blocks = operators.ZvalGetLong(elem)
	}
	return types.SUCCESS
}
func PhpUserstreamopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var func_name types.Zval
	var retval types.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	var ret int = -1
	func_name.SetString(USERSTREAM_STAT)
	call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && retval.IsType(types.IsArray) {
		if types.SUCCESS == StatbufFromArray(&retval, ssb) {
			ret = 0
		}
	} else {
		if call_result == types.FAILURE {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_STAT+" is not implemented!", us.GetWrapper().GetClassname())
		}
	}
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
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
		func_name.SetString(USERSTREAM_EOF)
		call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
		if call_result == types.SUCCESS && (retval.IsType(types.IsFalse) || retval.IsType(types.IsTrue)) {
			if operators.ZvalIsTrue(&retval) {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			}
		} else {
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		}
		// zend.ZvalPtrDtor(&retval)
		// zend.ZvalPtrDtor(&func_name)
	case core.PHP_STREAM_OPTION_LOCKING:
		args[0].SetLong(0)
		if (value & LOCK_NB) != 0 {
			args[0].SetLong(args[0].Long() | standard.PHP_LOCK_NB)
		}
		switch value & ^LOCK_NB {
		case LOCK_SH:
			args[0].SetLong(args[0].Long() | standard.PHP_LOCK_SH)
		case LOCK_EX:
			args[0].SetLong(args[0].Long() | standard.PHP_LOCK_EX)
		case LOCK_UN:
			args[0].SetLong(args[0].Long() | standard.PHP_LOCK_UN)
		}

		/* TODO wouldblock */

		func_name.SetString(USERSTREAM_LOCK)
		call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
		if call_result == types.SUCCESS && (retval.IsType(types.IsFalse) || retval.IsType(types.IsTrue)) {
			ret = retval.IsType(types.IsFalse)
		} else if call_result == types.FAILURE {
			if value == 0 {

				/* lock support test (TODO: more check) */

				ret = core.PHP_STREAM_OPTION_RETURN_OK

				/* lock support test (TODO: more check) */

			} else {
				core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_LOCK+" is not implemented!", us.GetWrapper().GetClassname())
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		// zend.ZvalPtrDtor(&retval)
		// zend.ZvalPtrDtor(&func_name)
		// zend.ZvalPtrDtor(&args[0])
	case core.PHP_STREAM_OPTION_TRUNCATE_API:
		func_name.SetString(USERSTREAM_TRUNCATE)
		switch value {
		case core.PHP_STREAM_TRUNCATE_SUPPORTED:
			var object *types.Object
			if !us.GetObject().IsUndef() {
				object = us.GetObject().Object()
			}

			if zend.IsCallable(&func_name, object, zend.IS_CALLABLE_CHECK_SILENT) {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size >= 0 && new_size <= ptrdiff_t(zend.LONG_MAX) {
				args[0].SetLong(zend.ZendLong(new_size))
				call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
				if call_result == types.SUCCESS && retval.IsNotUndef() {
					if retval.IsType(types.IsFalse) || retval.IsType(types.IsTrue) {
						if retval.IsType(types.IsTrue) {
							ret = core.PHP_STREAM_OPTION_RETURN_OK
						} else {
							ret = core.PHP_STREAM_OPTION_RETURN_ERR
						}
					} else {
						core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" did not return a boolean!", us.GetWrapper().GetClassname())
					}
				} else {
					core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_TRUNCATE+" is not implemented!", us.GetWrapper().GetClassname())
				}
				// zend.ZvalPtrDtor(&retval)
				// zend.ZvalPtrDtor(&args[0])
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
		}
		// zend.ZvalPtrDtor(&func_name)
	case core.PHP_STREAM_OPTION_READ_BUFFER:
		fallthrough
	case core.PHP_STREAM_OPTION_WRITE_BUFFER:
		fallthrough
	case core.PHP_STREAM_OPTION_READ_TIMEOUT:
		fallthrough
	case core.PHP_STREAM_OPTION_BLOCKING:
		func_name.SetString(USERSTREAM_SET_OPTION)
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
		call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 3, args, 0)
		if call_result == types.FAILURE {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_SET_OPTION+" is not implemented!", us.GetWrapper().GetClassname())
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
		} else if operators.ZvalIsTrue(&retval) {
			ret = core.PHP_STREAM_OPTION_RETURN_OK
		} else {
			ret = core.PHP_STREAM_OPTION_RETURN_ERR
		}
		// zend.ZvalPtrDtor(&retval)
		// zend.ZvalPtrDtor(&args[2])
		// zend.ZvalPtrDtor(&args[1])
		// zend.ZvalPtrDtor(&args[0])
		// zend.ZvalPtrDtor(&func_name)
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

	args[0].SetString(b.CastStrAuto(url))
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_UNLINK))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 1, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IsFalse) || zretval.IsType(types.IsTrue)) {
		ret = zretval.IsType(types.IsTrue)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_UNLINK+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[0])
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

	args[0].SetString(b.CastStrAuto(url_from))
	args[1].SetString(b.CastStrAuto(url_to))
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_RENAME))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IsFalse) || zretval.IsType(types.IsTrue)) {
		ret = zretval.IsType(types.IsTrue)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_RENAME+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
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

	args[0].SetString(b.CastStrAuto(url))
	args[1].SetLong(mode)
	args[2].SetLong(options)
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_MKDIR))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IsFalse) || zretval.IsType(types.IsTrue)) {
		ret = zretval.IsType(types.IsTrue)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_MKDIR+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[2])
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
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

	args[0].SetString(b.CastStrAuto(url))
	args[1].SetLong(options)
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_RMDIR))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IsFalse) || zretval.IsType(types.IsTrue)) {
		ret = zretval.IsType(types.IsTrue)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_RMDIR+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
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
		args[2].SetString(b.CastStrAuto(value))
	default:
		core.PhpErrorDocref("", faults.E_WARNING, "Unknown option %d for "+USERSTREAM_METADATA, option)
		// zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.IsUndef() {
		// zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* call the mkdir method */

	args[0].SetString(b.CastStrAuto(url))
	args[1].SetLong(option)
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_METADATA))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == types.SUCCESS && (zretval.IsType(types.IsFalse) || zretval.IsType(types.IsTrue)) {
		ret = zretval.IsType(types.IsTrue)
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_METADATA+" is not implemented!", uwrap.GetClassname())
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[0])
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[2])
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

	args[0].SetString(b.CastStrAuto(url))
	args[1].SetLong(flags)
	zfuncname.SetString(b.CastStrAuto(USERSTREAM_STATURL))
	call_result = zend.CallUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == types.SUCCESS && zretval.IsType(types.IsArray) {

		/* We got the info we needed */

		if types.SUCCESS == StatbufFromArray(&zretval, ssb) {
			ret = 0
		}

		/* We got the info we needed */

	} else {
		if call_result == types.FAILURE {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_STATURL+" is not implemented!", uwrap.GetClassname())
		}
	}

	/* clean up */

	// zend.ZvalPtrDtor(&object)
	// zend.ZvalPtrDtor(&zretval)
	// zend.ZvalPtrDtor(&zfuncname)
	// zend.ZvalPtrDtor(&args[1])
	// zend.ZvalPtrDtor(&args[0])
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
	func_name.SetString(USERSTREAM_DIR_READ)
	call_result = zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	if call_result == types.SUCCESS && !retval.IsFalse() && !retval.IsTrue() {
		operators.ConvertToString(&retval)
		core.PHP_STRLCPY(ent.GetDName(), retval.StringEx().GetVal(), b.SizeOf("ent -> d_name"), retval.StringEx().GetLen())
		didread = b.SizeOf("php_stream_dirent")
	} else if call_result == types.FAILURE {
		core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_DIR_READ+" is not implemented!", us.GetWrapper().GetClassname())
	}
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
	return didread
}
func PhpUserstreamopClosedir(stream *core.PhpStream, close_handle int) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	b.Assert(us != nil)
	func_name.SetString(USERSTREAM_DIR_CLOSE)
	zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
	// zend.ZvalPtrDtor(us.GetObject())
	us.GetObject().SetUndef()
	zend.Efree(us)
	return 0
}
func PhpUserstreamopRewinddir(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name types.Zval
	var retval types.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.GetAbstract())
	func_name.SetString(USERSTREAM_DIR_REWIND)
	zend.CallUserFunction(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 0, nil)
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
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
	func_name.SetString(USERSTREAM_CAST)
	switch castas {
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		args[0].SetLong(core.PHP_STREAM_AS_FD_FOR_SELECT)
	default:
		args[0].SetLong(core.PHP_STREAM_AS_STDIO)
	}
	call_result = zend.CallUserFunctionEx(lang.CondF2(us.GetObject().IsUndef(), nil, func() types.Zval { return us.GetObject() }), &func_name, &retval, 1, args, 0)
	for {
		if call_result == types.FAILURE {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_CAST+" is not implemented!", us.GetWrapper().GetClassname())
			break
		}
		if !operators.ZvalIsTrue(&retval) {
			break
		}
		core.PhpStreamFromZvalNoVerify(intstream, &retval)
		if intstream == nil {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_CAST+" must return a stream resource", us.GetWrapper().GetClassname())
			break
		}
		if intstream == stream {
			core.PhpErrorDocref("", faults.E_WARNING, "%s::"+USERSTREAM_CAST+" must not return itself", us.GetWrapper().GetClassname())
			intstream = nil
			break
		}
		ret = core.PhpStreamCast(intstream, castas, retptr, 1)
		break
	}
	// zend.ZvalPtrDtor(&retval)
	// zend.ZvalPtrDtor(&func_name)
	// zend.ZvalPtrDtor(&args[0])
	return ret
}
