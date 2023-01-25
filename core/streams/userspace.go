// <<generate>>

package streams

import (
	b "sik/builtin"
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	"sik/zend"
)

// Source: <main/streams/userspace.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   |          Sara Golemon <pollita@php.net>                              |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "php_globals.h"

// # include "ext/standard/file.h"

// # include "ext/standard/flock_compat.h"

// # include < sys / file . h >

// # include < stddef . h >

// # include < utime . h >

var LeProtocols int
var UserStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{UserWrapperOpener, nil, nil, UserWrapperStatUrl, UserWrapperOpendir, "user-space", UserWrapperUnlink, UserWrapperRename, UserWrapperMkdir, UserWrapperRmdir, UserWrapperMetadata}

func StreamWrapperDtor(rsrc *zend.ZendResource) {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(rsrc.ptr)
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

type PhpUserstreamDataT = _phpUserstreamData

/* names of methods */

const USERSTREAM_OPEN = "stream_open"
const USERSTREAM_CLOSE = "stream_close"
const USERSTREAM_READ = "stream_read"
const USERSTREAM_WRITE = "stream_write"
const USERSTREAM_FLUSH = "stream_flush"
const USERSTREAM_SEEK = "stream_seek"
const USERSTREAM_TELL = "stream_tell"
const USERSTREAM_EOF = "stream_eof"
const USERSTREAM_STAT = "stream_stat"
const USERSTREAM_STATURL = "url_stat"
const USERSTREAM_UNLINK = "unlink"
const USERSTREAM_RENAME = "rename"
const USERSTREAM_MKDIR = "mkdir"
const USERSTREAM_RMDIR = "rmdir"
const USERSTREAM_DIR_OPEN = "dir_opendir"
const USERSTREAM_DIR_READ = "dir_readdir"
const USERSTREAM_DIR_REWIND = "dir_rewinddir"
const USERSTREAM_DIR_CLOSE = "dir_closedir"
const USERSTREAM_LOCK = "stream_lock"
const USERSTREAM_CAST = "stream_cast"
const USERSTREAM_SET_OPTION = "stream_set_option"
const USERSTREAM_TRUNCATE = "stream_truncate"
const USERSTREAM_METADATA = "stream_metadata"

/* {{{ class should have methods like these:

   function stream_open($path, $mode, $options, &$opened_path)
   {
         return true/false;
   }

   function stream_read($count)
   {
          return false on error;
       else return string;
   }

   function stream_write($data)
   {
          return false on error;
       else return count written;
   }

   function stream_close()
   {
   }

   function stream_flush()
   {
       return true/false;
   }

   function stream_seek($offset, $whence)
   {
       return true/false;
   }

   function stream_tell()
   {
       return (int)$position;
   }

   function stream_eof()
   {
       return true/false;
   }

   function stream_stat()
   {
       return array( just like that returned by fstat() );
   }

   function stream_cast($castas)
   {
       if ($castas == STREAM_CAST_FOR_SELECT) {
           return $this->underlying_stream;
       }
       return false;
   }

   function stream_set_option($option, $arg1, $arg2)
   {
       switch($option) {
       case STREAM_OPTION_BLOCKING:
           $blocking = $arg1;
           ...
       case STREAM_OPTION_READ_TIMEOUT:
           $sec = $arg1;
           $usec = $arg2;
           ...
       case STREAM_OPTION_WRITE_BUFFER:
           $mode = $arg1;
           $size = $arg2;
           ...
       default:
           return false;
       }
   }

   function url_stat(string $url, int $flags)
   {
       return array( just like that returned by stat() );
   }

   function unlink(string $url)
   {
       return true / false;
   }

   function rename(string $from, string $to)
   {
       return true / false;
   }

   function mkdir($dir, $mode, $options)
   {
       return true / false;
   }

   function rmdir($dir, $options)
   {
       return true / false;
   }

   function dir_opendir(string $url, int $options)
   {
       return true / false;
   }

   function dir_readdir()
   {
       return string next filename in dir ;
   }

   function dir_closedir()
   {
       release dir related resources;
   }

   function dir_rewinddir()
   {
       reset to start of dir list;
   }

   function stream_lock($operation)
   {
       return true / false;
   }

    function stream_truncate($new_size)
   {
       return true / false;
   }

   }}} **/

func UserStreamCreateObject(uwrap *PhpUserStreamWrapper, context *core.PhpStreamContext, object *zend.Zval) {
	if (uwrap.GetCe().ce_flags & (zend.ZEND_ACC_INTERFACE | zend.ZEND_ACC_TRAIT | zend.ZEND_ACC_IMPLICIT_ABSTRACT_CLASS | zend.ZEND_ACC_EXPLICIT_ABSTRACT_CLASS)) != 0 {
		zend.ZVAL_UNDEF(object)
		return
	}

	/* create an instance of our class */

	if zend.ObjectInitEx(object, uwrap.GetCe()) == zend.FAILURE {
		zend.ZVAL_UNDEF(object)
		return
	}
	if context != nil {
		zend.AddPropertyResource(object, "context", context.GetRes())
		zend.GC_ADDREF(context.GetRes())
	} else {
		zend.AddPropertyNull(object, "context")
	}
	if uwrap.GetCe().constructor != nil {
		var fci zend.ZendFcallInfo
		var fcc zend.ZendFcallInfoCache
		var retval zend.Zval
		fci.size = b.SizeOf("fci")
		zend.ZVAL_UNDEF(&fci.function_name)
		fci.object = zend.Z_OBJ_P(object)
		fci.retval = &retval
		fci.param_count = 0
		fci.params = nil
		fci.no_separation = 1
		fcc.function_handler = uwrap.GetCe().constructor
		fcc.called_scope = zend.Z_OBJCE_P(object)
		fcc.object = zend.Z_OBJ_P(object)
		if zend.ZendCallFunction(&fci, &fcc) == zend.FAILURE {
			core.PhpErrorDocref(nil, zend.E_WARNING, "Could not execute %s::%s()", zend.ZSTR_VAL(uwrap.GetCe().name), zend.ZSTR_VAL(uwrap.GetCe().constructor.common.function_name))
			zend.ZvalPtrDtor(object)
			zend.ZVAL_UNDEF(object)
		} else {
			zend.ZvalPtrDtor(&retval)
		}
	}
}
func UserWrapperOpener(wrapper *core.PhpStreamWrapper, filename *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
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
	if uwrap.wrapper.is_url == 0 && (options&core.STREAM_OPEN_FOR_INCLUDE) != 0 && !(core.PG(allow_url_include)) {
		core.PG(in_user_include) = 1
	}
	us = zend.Emalloc(b.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, &us.object)
	if zend.Z_TYPE(us.GetObject()) == zend.IS_UNDEF {
		standard.FG(user_stream_current_filename) = nil
		core.PG(in_user_include) = old_in_user_include
		zend.Efree(us)
		return nil
	}

	/* call it's stream_open method - set up params first */

	zend.ZVAL_STRING(&args[0], filename)
	zend.ZVAL_STRING(&args[1], mode)
	zend.ZVAL_LONG(&args[2], options)
	zend.ZVAL_NEW_REF(&args[3], &(zend.ExecutorGlobals.uninitialized_zval))
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_OPEN)
	var __orig_bailout *JMP_BUF = zend.ExecutorGlobals.bailout
	var __bailout JMP_BUF
	zend.ExecutorGlobals.bailout = &__bailout
	if zend.SETJMP(__bailout) == 0 {
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &zfuncname, &zretval, 4, args, 0, nil)
	} else {
		zend.ExecutorGlobals.bailout = __orig_bailout
		standard.FG(user_stream_current_filename) = nil
		zend.ZendBailout()
	}
	zend.ExecutorGlobals.bailout = __orig_bailout
	if call_result == zend.SUCCESS && zend.Z_TYPE(zretval) != zend.IS_UNDEF && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceOps, us, 0, mode)

		/* if the opened path is set, copy it out */

		if zend.Z_ISREF(args[3]) && zend.Z_TYPE_P(zend.Z_REFVAL(args[3])) == zend.IS_STRING && opened_path != nil {
			*opened_path = zend.ZendStringCopy(zend.Z_STR_P(zend.Z_REFVAL(args[3])))
		}

		/* set wrapper data to be a reference to our object */

		zend.ZVAL_COPY(&stream.wrapperdata, &us.object)

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+USERSTREAM_OPEN+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(&us.object)
		zend.ZVAL_UNDEF(&us.object)
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
func UserWrapperOpendir(wrapper *core.PhpStreamWrapper, filename *byte, mode *byte, options int, opened_path **zend.ZendString, context *core.PhpStreamContext) *core.PhpStream {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
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
	UserStreamCreateObject(uwrap, context, &us.object)
	if zend.Z_TYPE(us.GetObject()) == zend.IS_UNDEF {
		standard.FG(user_stream_current_filename) = nil
		zend.Efree(us)
		return nil
	}

	/* call it's dir_open method - set up params first */

	zend.ZVAL_STRING(&args[0], filename)
	zend.ZVAL_LONG(&args[1], options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_DIR_OPEN)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(zretval) != zend.IS_UNDEF && zend.ZvalIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = core.PhpStreamAllocRel(&PhpStreamUserspaceDirOps, us, 0, mode)

		/* set wrapper data to be a reference to our object */

		zend.ZVAL_COPY(&stream.wrapperdata, &us.object)

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+USERSTREAM_DIR_OPEN+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(&us.object)
		zend.ZVAL_UNDEF(&us.object)
		zend.Efree(us)
	}
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	standard.FG(user_stream_current_filename) = nil
	return stream
}

/* {{{ proto bool stream_wrapper_register(string protocol, string classname[, int flags])
   Registers a custom URL protocol handler class */

func ZifStreamWrapperRegister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	var classname *zend.ZendString
	var uwrap *PhpUserStreamWrapper
	var rsrc *zend.ZendResource
	var flags zend.ZendLong = 0
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "SS|l", &protocol, &classname, &flags) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	uwrap = (*PhpUserStreamWrapper)(zend.Ecalloc(1, b.SizeOf("* uwrap")))
	uwrap.SetProtoname(zend.Estrndup(zend.ZSTR_VAL(protocol), zend.ZSTR_LEN(protocol)))
	uwrap.SetClassname(zend.Estrndup(zend.ZSTR_VAL(classname), zend.ZSTR_LEN(classname)))
	uwrap.wrapper.wops = &UserStreamWops
	uwrap.wrapper.abstract = uwrap
	uwrap.wrapper.is_url = (flags & core.PHP_STREAM_IS_URL) != 0
	rsrc = zend.ZendRegisterResource(uwrap, LeProtocols)
	if b.Assign(&(uwrap.GetCe()), zend.ZendLookupClass(classname)) != nil {
		if PhpRegisterUrlStreamWrapperVolatile(protocol, &uwrap.wrapper) == zend.SUCCESS {
			zend.RETVAL_TRUE
			return
		} else {

			/* We failed.  But why? */

			if zend.ZendHashExists(core.PhpStreamGetUrlStreamWrappersHash(), protocol) != 0 {
				core.PhpErrorDocref(nil, zend.E_WARNING, "Protocol %s:// is already defined.", zend.ZSTR_VAL(protocol))
			} else {

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

				core.PhpErrorDocref(nil, zend.E_WARNING, "Invalid protocol scheme specified. Unable to register wrapper class %s to %s://", zend.ZSTR_VAL(classname), zend.ZSTR_VAL(protocol))

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

			}

			/* We failed.  But why? */

		}
	} else {
		core.PhpErrorDocref(nil, zend.E_WARNING, "class '%s' is undefined", zend.ZSTR_VAL(classname))
	}
	zend.ZendListDelete(rsrc)
	zend.RETVAL_FALSE
	return
}

/* }}} */

func ZifStreamWrapperUnregister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &protocol) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	if PhpUnregisterUrlStreamWrapperVolatile(protocol) == zend.FAILURE {

		/* We failed */

		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to unregister protocol %s://", zend.ZSTR_VAL(protocol))
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */

func ZifStreamWrapperRestore(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	var wrapper *core.PhpStreamWrapper
	var global_wrapper_hash *zend.HashTable
	var wrapper_hash *zend.HashTable
	if zend.ZendParseParameters(zend.ZEND_NUM_ARGS(), "S", &protocol) == zend.FAILURE {
		zend.RETVAL_FALSE
		return
	}
	global_wrapper_hash = PhpStreamGetUrlStreamWrappersHashGlobal()
	if b.Assign(&wrapper, zend.ZendHashFindPtr(global_wrapper_hash, protocol)) == nil {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s:// never existed, nothing to restore", zend.ZSTR_VAL(protocol))
		zend.RETVAL_FALSE
		return
	}
	wrapper_hash = core.PhpStreamGetUrlStreamWrappersHash()
	if wrapper_hash == global_wrapper_hash || zend.ZendHashFindPtr(wrapper_hash, protocol) == wrapper {
		core.PhpErrorDocref(nil, zend.E_NOTICE, "%s:// was never changed, nothing to restore", zend.ZSTR_VAL(protocol))
		zend.RETVAL_TRUE
		return
	}

	/* A failure here could be okay given that the protocol might have been merely unregistered */

	PhpUnregisterUrlStreamWrapperVolatile(protocol)
	if PhpRegisterUrlStreamWrapperVolatile(protocol, wrapper) == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Unable to restore original %s:// wrapper", zend.ZSTR_VAL(protocol))
		zend.RETVAL_FALSE
		return
	}
	zend.RETVAL_TRUE
	return
}

/* }}} */

func PhpUserstreamopWrite(stream *core.PhpStream, buf *byte, count int) ssize_t {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var args []zend.Zval
	var didwrite ssize_t
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_WRITE, b.SizeOf("USERSTREAM_WRITE")-1)
	zend.ZVAL_STRINGL(&args[0], (*byte)(buf), count)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.ExecutorGlobals.exception != nil {
		return -1
	}
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
		if zend.Z_TYPE(retval) == zend.IS_FALSE {
			didwrite = -1
		} else {
			zend.ConvertToLong(&retval)
			didwrite = zend.Z_LVAL(retval)
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
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_READ, b.SizeOf("USERSTREAM_READ")-1)
	zend.ZVAL_LONG(&args[0], count)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.ExecutorGlobals.exception != nil {
		return -1
	}
	if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_READ+" is not implemented!", us.GetWrapper().GetClassname())
		return -1
	}
	if zend.Z_TYPE(retval) == zend.IS_FALSE {
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
	zend.ZVAL_UNDEF(&retval)

	/* since the user stream has no way of setting the eof flag directly, we need to ask it if we hit eof */

	zend.ZVAL_STRINGL(&func_name, USERSTREAM_EOF, b.SizeOf("USERSTREAM_EOF")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&func_name)
	if zend.ExecutorGlobals.exception != nil {
		stream.eof = 1
		return -1
	}
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
		stream.eof = 1
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, zend.E_WARNING, "%s::"+USERSTREAM_EOF+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		stream.eof = 1
	}
	zend.ZvalPtrDtor(&retval)
	return didread
}
func PhpUserstreamopClose(stream *core.PhpStream, close_handle int) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_CLOSE, b.SizeOf("USERSTREAM_CLOSE")-1)
	zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&us.object)
	zend.ZVAL_UNDEF(&us.object)
	zend.Efree(us)
	return 0
}
func PhpUserstreamopFlush(stream *core.PhpStream) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_FLUSH, b.SizeOf("USERSTREAM_FLUSH")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
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
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var args []zend.Zval
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_SEEK, b.SizeOf("USERSTREAM_SEEK")-1)
	zend.ZVAL_LONG(&args[0], offset)
	zend.ZVAL_LONG(&args[1], whence)
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 2, args, 0, nil)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&func_name)
	if call_result == zend.FAILURE {

		/* stream_seek is not implemented, so disable seeks for this stream */

		stream.flags |= core.PHP_STREAM_FLAG_NO_SEEK

		/* there should be no retval to clean up */

		zend.ZvalPtrDtor(&retval)
		return -1
	} else if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF && zend.ZvalIsTrue(&retval) != 0 {
		ret = 0
	} else {
		ret = -1
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZVAL_UNDEF(&retval)
	if ret != 0 {
		return ret
	}

	/* now determine where we are */

	zend.ZVAL_STRINGL(&func_name, USERSTREAM_TELL, b.SizeOf("USERSTREAM_TELL")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) == zend.IS_LONG {
		*newoffs = zend.Z_LVAL(retval)
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

/* parse the return value from one of the stat functions and store the
 * relevant fields into the statbuf provided */

func StatbufFromArray(array *zend.Zval, ssb *core.PhpStreamStatbuf) int {
	var elem *zend.Zval

	// #define STAT_PROP_ENTRY_EX(name,name2) if ( NULL != ( elem = zend_hash_str_find ( Z_ARRVAL_P ( array ) , # name , sizeof ( # name ) - 1 ) ) ) { ssb -> sb . st_ ## name2 = zval_get_long ( elem ) ; }

	// #define STAT_PROP_ENTRY(name) STAT_PROP_ENTRY_EX ( name , name )

	memset(ssb, 0, b.SizeOf("php_stream_statbuf"))
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "dev", b.SizeOf("\"dev\"")-1)) {
		ssb.sb.st_dev = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "ino", b.SizeOf("\"ino\"")-1)) {
		ssb.sb.st_ino = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "mode", b.SizeOf("\"mode\"")-1)) {
		ssb.sb.st_mode = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "nlink", b.SizeOf("\"nlink\"")-1)) {
		ssb.sb.st_nlink = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "uid", b.SizeOf("\"uid\"")-1)) {
		ssb.sb.st_uid = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "gid", b.SizeOf("\"gid\"")-1)) {
		ssb.sb.st_gid = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "rdev", b.SizeOf("\"rdev\"")-1)) {
		ssb.sb.st_rdev = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "size", b.SizeOf("\"size\"")-1)) {
		ssb.sb.st_size = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "atime", b.SizeOf("\"atime\"")-1)) {
		ssb.sb.st_atime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "mtime", b.SizeOf("\"mtime\"")-1)) {
		ssb.sb.st_mtime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "ctime", b.SizeOf("\"ctime\"")-1)) {
		ssb.sb.st_ctime = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "blksize", b.SizeOf("\"blksize\"")-1)) {
		ssb.sb.st_blksize = zend.ZvalGetLong(elem)
	}
	if nil != b.Assign(&elem, zend.ZendHashStrFind(zend.Z_ARRVAL_P(array), "blocks", b.SizeOf("\"blocks\"")-1)) {
		ssb.sb.st_blocks = zend.ZvalGetLong(elem)
	}
	return zend.SUCCESS
}
func PhpUserstreamopStat(stream *core.PhpStream, ssb *core.PhpStreamStatbuf) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var ret int = -1
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_STAT, b.SizeOf("USERSTREAM_STAT")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) == zend.IS_ARRAY {
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
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var ret int = core.PHP_STREAM_OPTION_RETURN_NOTIMPL
	var args []zend.Zval
	switch option {
	case core.PHP_STREAM_OPTION_CHECK_LIVENESS:
		zend.ZVAL_STRINGL(&func_name, USERSTREAM_EOF, b.SizeOf("USERSTREAM_EOF")-1)
		call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
		if call_result == zend.SUCCESS && (zend.Z_TYPE(retval) == zend.IS_FALSE || zend.Z_TYPE(retval) == zend.IS_TRUE) {
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
		zend.ZVAL_LONG(&args[0], 0)
		if (value & LOCK_NB) != 0 {
			zend.Z_LVAL_P(&args[0]) |= standard.PHP_LOCK_NB
		}
		switch value & ^LOCK_NB {
		case LOCK_SH:
			zend.Z_LVAL_P(&args[0]) |= standard.PHP_LOCK_SH
			break
		case LOCK_EX:
			zend.Z_LVAL_P(&args[0]) |= standard.PHP_LOCK_EX
			break
		case LOCK_UN:
			zend.Z_LVAL_P(&args[0]) |= standard.PHP_LOCK_UN
			break
		}

		/* TODO wouldblock */

		zend.ZVAL_STRINGL(&func_name, USERSTREAM_LOCK, b.SizeOf("USERSTREAM_LOCK")-1)
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0, nil)
		if call_result == zend.SUCCESS && (zend.Z_TYPE(retval) == zend.IS_FALSE || zend.Z_TYPE(retval) == zend.IS_TRUE) {
			ret = zend.Z_TYPE(retval) == zend.IS_FALSE
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
			if zend.ZendIsCallableEx(&func_name, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() *zend.ZendObject { return zend.Z_OBJ(us.GetObject()) }), zend.IS_CALLABLE_CHECK_SILENT, nil, nil, nil) != 0 {
				ret = core.PHP_STREAM_OPTION_RETURN_OK
			} else {
				ret = core.PHP_STREAM_OPTION_RETURN_ERR
			}
			break
		case core.PHP_STREAM_TRUNCATE_SET_SIZE:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size >= 0 && new_size <= ptrdiff_t(zend.LONG_MAX) {
				zend.ZVAL_LONG(&args[0], zend.ZendLong(new_size))
				call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0, nil)
				if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_UNDEF {
					if zend.Z_TYPE(retval) == zend.IS_FALSE || zend.Z_TYPE(retval) == zend.IS_TRUE {
						if zend.Z_TYPE(retval) == zend.IS_TRUE {
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
		zend.ZVAL_LONG(&args[0], option)
		zend.ZVAL_NULL(&args[1])
		zend.ZVAL_NULL(&args[2])
		switch option {
		case core.PHP_STREAM_OPTION_READ_BUFFER:

		case core.PHP_STREAM_OPTION_WRITE_BUFFER:
			zend.ZVAL_LONG(&args[1], value)
			if ptrparam {
				zend.ZVAL_LONG(&args[2], *((*long)(ptrparam)))
			} else {
				zend.ZVAL_LONG(&args[2], r.BUFSIZ)
			}
			break
		case core.PHP_STREAM_OPTION_READ_TIMEOUT:
			var tv __struct__timeval = *((*__struct__timeval)(ptrparam))
			zend.ZVAL_LONG(&args[1], tv.tv_sec)
			zend.ZVAL_LONG(&args[2], tv.tv_usec)
			break
		case core.PHP_STREAM_OPTION_BLOCKING:
			zend.ZVAL_LONG(&args[1], value)
			break
		default:
			break
		}
		call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 3, args, 0, nil)
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		return ret
	}

	/* call the unlink method */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_UNLINK)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 1, args, 0, nil)
	if call_result == zend.SUCCESS && (zend.Z_TYPE(zretval) == zend.IS_FALSE || zend.Z_TYPE(zretval) == zend.IS_TRUE) {
		ret = zend.Z_TYPE(zretval) == zend.IS_TRUE
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		return ret
	}

	/* call the rename method */

	zend.ZVAL_STRING(&args[0], url_from)
	zend.ZVAL_STRING(&args[1], url_to)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_RENAME)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && (zend.Z_TYPE(zretval) == zend.IS_FALSE || zend.Z_TYPE(zretval) == zend.IS_TRUE) {
		ret = zend.Z_TYPE(zretval) == zend.IS_TRUE
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		return ret
	}

	/* call the mkdir method */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_LONG(&args[1], mode)
	zend.ZVAL_LONG(&args[2], options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_MKDIR)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 3, args, 0, nil)
	if call_result == zend.SUCCESS && (zend.Z_TYPE(zretval) == zend.IS_FALSE || zend.Z_TYPE(zretval) == zend.IS_TRUE) {
		ret = zend.Z_TYPE(zretval) == zend.IS_TRUE
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = 0

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		return ret
	}

	/* call the rmdir method */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_LONG(&args[1], options)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_RMDIR)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && (zend.Z_TYPE(zretval) == zend.IS_FALSE || zend.Z_TYPE(zretval) == zend.IS_TRUE) {
		ret = zend.Z_TYPE(zretval) == zend.IS_TRUE
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
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
		zend.ZVAL_LONG(&args[2], *((*long)(value)))
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
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* call the mkdir method */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_LONG(&args[1], option)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_METADATA)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 3, args, 0, nil)
	if call_result == zend.SUCCESS && (zend.Z_TYPE(zretval) == zend.IS_FALSE || zend.Z_TYPE(zretval) == zend.IS_TRUE) {
		ret = zend.Z_TYPE(zretval) == zend.IS_TRUE
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
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(wrapper.abstract)
	var zfuncname zend.Zval
	var zretval zend.Zval
	var args []zend.Zval
	var call_result int
	var object zend.Zval
	var ret int = -1

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if zend.Z_TYPE(object) == zend.IS_UNDEF {
		return ret
	}

	/* call it's stat_url method - set up params first */

	zend.ZVAL_STRING(&args[0], url)
	zend.ZVAL_LONG(&args[1], flags)
	zend.ZVAL_STRING(&zfuncname, USERSTREAM_STATURL)
	call_result = zend.CallUserFunctionEx(nil, &object, &zfuncname, &zretval, 2, args, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(zretval) == zend.IS_ARRAY {

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
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var ent *core.PhpStreamDirent = (*core.PhpStreamDirent)(buf)

	/* avoid problems if someone mis-uses the stream */

	if count != b.SizeOf("php_stream_dirent") {
		return -1
	}
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_READ, b.SizeOf("USERSTREAM_DIR_READ")-1)
	call_result = zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	if call_result == zend.SUCCESS && zend.Z_TYPE(retval) != zend.IS_FALSE && zend.Z_TYPE(retval) != zend.IS_TRUE {
		zend.ConvertToString(&retval)
		core.PHP_STRLCPY(ent.d_name, zend.Z_STRVAL(retval), b.SizeOf("ent -> d_name"), zend.Z_STRLEN(retval))
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
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	r.Assert(us != nil)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_CLOSE, b.SizeOf("USERSTREAM_DIR_CLOSE")-1)
	zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&us.object)
	zend.ZVAL_UNDEF(&us.object)
	zend.Efree(us)
	return 0
}
func PhpUserstreamopRewinddir(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_DIR_REWIND, b.SizeOf("USERSTREAM_DIR_REWIND")-1)
	zend.CallUserFunction(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	return 0
}
func PhpUserstreamopCast(stream *core.PhpStream, castas int, retptr *any) int {
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var func_name zend.Zval
	var retval zend.Zval
	var args []zend.Zval
	var intstream *core.PhpStream = nil
	var call_result int
	var ret int = zend.FAILURE
	zend.ZVAL_STRINGL(&func_name, USERSTREAM_CAST, b.SizeOf("USERSTREAM_CAST")-1)
	switch castas {
	case core.PHP_STREAM_AS_FD_FOR_SELECT:
		zend.ZVAL_LONG(&args[0], core.PHP_STREAM_AS_FD_FOR_SELECT)
		break
	default:
		zend.ZVAL_LONG(&args[0], core.PHP_STREAM_AS_STDIO)
		break
	}
	call_result = zend.CallUserFunctionEx(nil, b.CondF2(zend.Z_ISUNDEF(us.GetObject()), nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0, nil)
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

var PhpStreamUserspaceOps core.PhpStreamOps = core.PhpStreamOps{PhpUserstreamopWrite, PhpUserstreamopRead, PhpUserstreamopClose, PhpUserstreamopFlush, "user-space", PhpUserstreamopSeek, PhpUserstreamopCast, PhpUserstreamopStat, PhpUserstreamopSetOption}
var PhpStreamUserspaceDirOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpUserstreamopReaddir, PhpUserstreamopClosedir, nil, "user-space-dir", PhpUserstreamopRewinddir, nil, nil, nil}
