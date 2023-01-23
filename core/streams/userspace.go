// <<generate>>

package streams

import (
	"sik/core"
	"sik/ext/standard"
	r "sik/runtime"
	g "sik/runtime/grammar"
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

// @type PhpUserStreamWrapper struct
var UserStreamWops core.PhpStreamWrapperOps = core.PhpStreamWrapperOps{UserWrapperOpener, nil, nil, UserWrapperStatUrl, UserWrapperOpendir, "user-space", UserWrapperUnlink, UserWrapperRename, UserWrapperMkdir, UserWrapperRmdir, UserWrapperMetadata}

func StreamWrapperDtor(rsrc *zend.ZendResource) {
	var uwrap *PhpUserStreamWrapper = (*PhpUserStreamWrapper)(rsrc.ptr)
	zend._efree(uwrap.GetProtoname())
	zend._efree(uwrap.GetClassname())
	zend._efree(uwrap)
}
func ZmStartupUserStreams(type_ int, module_number int) int {
	LeProtocols = zend.ZendRegisterListDestructorsEx(StreamWrapperDtor, nil, "stream factory", 0)
	if LeProtocols == zend.FAILURE {
		return zend.FAILURE
	}
	zend.ZendRegisterLongConstant("STREAM_USE_PATH", g.SizeOf("\"STREAM_USE_PATH\"")-1, 0x1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_IGNORE_URL", g.SizeOf("\"STREAM_IGNORE_URL\"")-1, 0x2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_REPORT_ERRORS", g.SizeOf("\"STREAM_REPORT_ERRORS\"")-1, 0x8, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_MUST_SEEK", g.SizeOf("\"STREAM_MUST_SEEK\"")-1, 0x10, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_URL_STAT_LINK", g.SizeOf("\"STREAM_URL_STAT_LINK\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_URL_STAT_QUIET", g.SizeOf("\"STREAM_URL_STAT_QUIET\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_MKDIR_RECURSIVE", g.SizeOf("\"STREAM_MKDIR_RECURSIVE\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_IS_URL", g.SizeOf("\"STREAM_IS_URL\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_OPTION_BLOCKING", g.SizeOf("\"STREAM_OPTION_BLOCKING\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_OPTION_READ_TIMEOUT", g.SizeOf("\"STREAM_OPTION_READ_TIMEOUT\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_OPTION_READ_BUFFER", g.SizeOf("\"STREAM_OPTION_READ_BUFFER\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_OPTION_WRITE_BUFFER", g.SizeOf("\"STREAM_OPTION_WRITE_BUFFER\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_BUFFER_NONE", g.SizeOf("\"STREAM_BUFFER_NONE\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_BUFFER_LINE", g.SizeOf("\"STREAM_BUFFER_LINE\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_BUFFER_FULL", g.SizeOf("\"STREAM_BUFFER_FULL\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CAST_AS_STREAM", g.SizeOf("\"STREAM_CAST_AS_STREAM\"")-1, 0, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_CAST_FOR_SELECT", g.SizeOf("\"STREAM_CAST_FOR_SELECT\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_TOUCH", g.SizeOf("\"STREAM_META_TOUCH\"")-1, 1, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_OWNER", g.SizeOf("\"STREAM_META_OWNER\"")-1, 3, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_OWNER_NAME", g.SizeOf("\"STREAM_META_OWNER_NAME\"")-1, 2, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_GROUP", g.SizeOf("\"STREAM_META_GROUP\"")-1, 5, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_GROUP_NAME", g.SizeOf("\"STREAM_META_GROUP_NAME\"")-1, 4, 1<<0|1<<1, module_number)
	zend.ZendRegisterLongConstant("STREAM_META_ACCESS", g.SizeOf("\"STREAM_META_ACCESS\"")-1, 6, 1<<0|1<<1, module_number)
	return zend.SUCCESS
}

// @type _phpUserstreamData struct
type PhpUserstreamDataT = _phpUserstreamData

/* names of methods */

// #define USERSTREAM_OPEN       "stream_open"

// #define USERSTREAM_CLOSE       "stream_close"

// #define USERSTREAM_READ       "stream_read"

// #define USERSTREAM_WRITE       "stream_write"

// #define USERSTREAM_FLUSH       "stream_flush"

// #define USERSTREAM_SEEK       "stream_seek"

// #define USERSTREAM_TELL       "stream_tell"

// #define USERSTREAM_EOF       "stream_eof"

// #define USERSTREAM_STAT       "stream_stat"

// #define USERSTREAM_STATURL       "url_stat"

// #define USERSTREAM_UNLINK       "unlink"

// #define USERSTREAM_RENAME       "rename"

// #define USERSTREAM_MKDIR       "mkdir"

// #define USERSTREAM_RMDIR       "rmdir"

// #define USERSTREAM_DIR_OPEN       "dir_opendir"

// #define USERSTREAM_DIR_READ       "dir_readdir"

// #define USERSTREAM_DIR_REWIND       "dir_rewinddir"

// #define USERSTREAM_DIR_CLOSE       "dir_closedir"

// #define USERSTREAM_LOCK       "stream_lock"

// #define USERSTREAM_CAST       "stream_cast"

// #define USERSTREAM_SET_OPTION       "stream_set_option"

// #define USERSTREAM_TRUNCATE       "stream_truncate"

// #define USERSTREAM_METADATA       "stream_metadata"

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
	if (uwrap.GetCe().ce_flags & (1<<0 | 1<<1 | 1<<4 | 1<<6)) != 0 {
		object.u1.type_info = 0
		return
	}

	/* create an instance of our class */

	if zend.ObjectInitEx(object, uwrap.GetCe()) == zend.FAILURE {
		object.u1.type_info = 0
		return
	}
	if context != nil {
		zend.AddPropertyResourceEx(object, "context", strlen("context"), context.GetRes())
		zend.ZendGcAddref(&(context.GetRes()).gc)
	} else {
		zend.AddPropertyNullEx(object, "context", strlen("context"))
	}
	if uwrap.GetCe().constructor != nil {
		var fci zend.ZendFcallInfo
		var fcc zend.ZendFcallInfoCache
		var retval zend.Zval
		fci.size = g.SizeOf("fci")
		&fci.function_name.u1.type_info = 0
		fci.object = object.value.obj
		fci.retval = &retval
		fci.param_count = 0
		fci.params = nil
		fci.no_separation = 1
		fcc.function_handler = uwrap.GetCe().constructor
		fcc.called_scope = object.value.obj.ce
		fcc.object = object.value.obj
		if zend.ZendCallFunction(&fci, &fcc) == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "Could not execute %s::%s()", uwrap.GetCe().name.val, uwrap.GetCe().constructor.common.function_name.val)
			zend.ZvalPtrDtor(object)
			object.u1.type_info = 0
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

	if standard.FileGlobals.user_stream_current_filename != nil && strcmp(filename, standard.FileGlobals.user_stream_current_filename) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FileGlobals.user_stream_current_filename = filename

	/* if the user stream was registered as local and we are in include context,
	   we add allow_url_include restrictions to allow_url_fopen ones */

	old_in_user_include = core.CoreGlobals.in_user_include
	if uwrap.wrapper.is_url == 0 && (options&0x80) != 0 && core.CoreGlobals.allow_url_include == 0 {
		core.CoreGlobals.in_user_include = 1
	}
	us = zend._emalloc(g.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, &us.object)
	if us.object.u1.v.type_ == 0 {
		standard.FileGlobals.user_stream_current_filename = nil
		core.CoreGlobals.in_user_include = old_in_user_include
		zend._efree(us)
		return nil
	}

	/* call it's stream_open method - set up params first */

	var _s *byte = filename
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var _s *byte = mode
	var __z *zend.Zval = &args[1]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[2]
	__z.value.lval = options
	__z.u1.type_info = 4
	var _ref *zend.ZendReference = (*zend.ZendReference)(zend._emalloc(g.SizeOf("zend_reference")))
	zend.ZendGcSetRefcount(&_ref.gc, 1)
	_ref.gc.u.type_info = 10
	var _z1 *zend.Zval = &_ref.val
	var _z2 *zend.Zval = &zend.EG.uninitialized_zval
	var _gc *zend.ZendRefcounted = _z2.value.counted
	var _t uint32 = _z2.u1.type_info
	_z1.value.counted = _gc
	_z1.u1.type_info = _t
	_ref.sources.ptr = nil
	&args[3].value.ref = _ref
	&args[3].u1.type_info = 10 | 1<<0<<8
	var _s *byte = "stream_open"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __orig_bailout *sigjmp_buf = zend.EG.bailout
	var __bailout sigjmp_buf
	zend.EG.bailout = &__bailout
	if sigsetjmp(__bailout, 0) == 0 {
		call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &zfuncname, &zretval, 4, args, 0)
	} else {
		zend.EG.bailout = __orig_bailout
		standard.FileGlobals.user_stream_current_filename = nil
		zend._zendBailout(__FILE__, __LINE__)
	}
	zend.EG.bailout = __orig_bailout
	if call_result == zend.SUCCESS && zretval.u1.v.type_ != 0 && zend.ZendIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = _phpStreamAlloc(&PhpStreamUserspaceOps, us, 0, mode)

		/* if the opened path is set, copy it out */

		if args[3].u1.v.type_ == 10 && &args[3].value.ref.val.u1.v.type_ == 6 && opened_path != nil {
			*opened_path = zend.ZendStringCopy(&args[3].value.ref.val.value.str)
		}

		/* set wrapper data to be a reference to our object */

		var _z1 *zend.Zval = &stream.wrapperdata
		var _z2 *zend.Zval = &us.object
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+"stream_open"+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(&us.object)
		&us.object.u1.type_info = 0
		zend._efree(us)
	}
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[3])
	zend.ZvalPtrDtor(&args[2])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	standard.FileGlobals.user_stream_current_filename = nil
	core.CoreGlobals.in_user_include = old_in_user_include
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

	if standard.FileGlobals.user_stream_current_filename != nil && strcmp(filename, standard.FileGlobals.user_stream_current_filename) == 0 {
		PhpStreamWrapperLogError(wrapper, options, "infinite recursion prevented")
		return nil
	}
	standard.FileGlobals.user_stream_current_filename = filename
	us = zend._emalloc(g.SizeOf("* us"))
	us.SetWrapper(uwrap)
	UserStreamCreateObject(uwrap, context, &us.object)
	if us.object.u1.v.type_ == 0 {
		standard.FileGlobals.user_stream_current_filename = nil
		zend._efree(us)
		return nil
	}

	/* call it's dir_open method - set up params first */

	var _s *byte = filename
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[1]
	__z.value.lval = options
	__z.u1.type_info = 4
	var _s *byte = "dir_opendir"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &zfuncname, &zretval, 2, args, 0)
	if call_result == zend.SUCCESS && zretval.u1.v.type_ != 0 && zend.ZendIsTrue(&zretval) != 0 {

		/* the stream is now open! */

		stream = _phpStreamAlloc(&PhpStreamUserspaceDirOps, us, 0, mode)

		/* set wrapper data to be a reference to our object */

		var _z1 *zend.Zval = &stream.wrapperdata
		var _z2 *zend.Zval = &us.object
		var _gc *zend.ZendRefcounted = _z2.value.counted
		var _t uint32 = _z2.u1.type_info
		_z1.value.counted = _gc
		_z1.u1.type_info = _t
		if (_t & 0xff00) != 0 {
			zend.ZendGcAddref(&_gc.gc)
		}

		/* set wrapper data to be a reference to our object */

	} else {
		PhpStreamWrapperLogError(wrapper, options, "\"%s::"+"dir_opendir"+"\" call failed", us.GetWrapper().GetClassname())
	}

	/* destroy everything else */

	if stream == nil {
		zend.ZvalPtrDtor(&us.object)
		&us.object.u1.type_info = 0
		zend._efree(us)
	}
	zend.ZvalPtrDtor(&zretval)
	zend.ZvalPtrDtor(&zfuncname)
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&args[0])
	standard.FileGlobals.user_stream_current_filename = nil
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
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "SS|l", &protocol, &classname, &flags) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	uwrap = (*PhpUserStreamWrapper)(zend._ecalloc(1, g.SizeOf("* uwrap")))
	uwrap.SetProtoname(zend._estrndup(protocol.val, protocol.len_))
	uwrap.SetClassname(zend._estrndup(classname.val, classname.len_))
	uwrap.wrapper.wops = &UserStreamWops
	uwrap.wrapper.abstract = uwrap
	uwrap.wrapper.is_url = (flags & 1) != 0
	rsrc = zend.ZendRegisterResource(uwrap, LeProtocols)
	if g.Assign(&(uwrap.GetCe()), zend.ZendLookupClass(classname)) != nil {
		if PhpRegisterUrlStreamWrapperVolatile(protocol, &uwrap.wrapper) == zend.SUCCESS {
			return_value.u1.type_info = 3
			return
		} else {

			/* We failed.  But why? */

			if zend.ZendHashExists(_phpStreamGetUrlStreamWrappersHash(), protocol) != 0 {
				core.PhpErrorDocref(nil, 1<<1, "Protocol %s:// is already defined.", protocol.val)
			} else {

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

				core.PhpErrorDocref(nil, 1<<1, "Invalid protocol scheme specified. Unable to register wrapper class %s to %s://", classname.val, protocol.val)

				/* Hash doesn't exist so it must have been an invalid protocol scheme */

			}

			/* We failed.  But why? */

		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "class '%s' is undefined", classname.val)
	}
	zend.ZendListDelete(rsrc)
	return_value.u1.type_info = 2
	return
}

/* }}} */

func ZifStreamWrapperUnregister(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S", &protocol) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	if PhpUnregisterUrlStreamWrapperVolatile(protocol) == zend.FAILURE {

		/* We failed */

		core.PhpErrorDocref(nil, 1<<1, "Unable to unregister protocol %s://", protocol.val)
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
	return
}

/* }}} */

func ZifStreamWrapperRestore(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var protocol *zend.ZendString
	var wrapper *core.PhpStreamWrapper
	var global_wrapper_hash *zend.HashTable
	var wrapper_hash *zend.HashTable
	if zend.ZendParseParameters(execute_data.This.u2.num_args, "S", &protocol) == zend.FAILURE {
		return_value.u1.type_info = 2
		return
	}
	global_wrapper_hash = PhpStreamGetUrlStreamWrappersHashGlobal()
	if g.Assign(&wrapper, zend.ZendHashFindPtr(global_wrapper_hash, protocol)) == nil {
		core.PhpErrorDocref(nil, 1<<1, "%s:// never existed, nothing to restore", protocol.val)
		return_value.u1.type_info = 2
		return
	}
	wrapper_hash = _phpStreamGetUrlStreamWrappersHash()
	if wrapper_hash == global_wrapper_hash || zend.ZendHashFindPtr(wrapper_hash, protocol) == wrapper {
		core.PhpErrorDocref(nil, 1<<3, "%s:// was never changed, nothing to restore", protocol.val)
		return_value.u1.type_info = 3
		return
	}

	/* A failure here could be okay given that the protocol might have been merely unregistered */

	PhpUnregisterUrlStreamWrapperVolatile(protocol)
	if PhpRegisterUrlStreamWrapperVolatile(protocol, wrapper) == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "Unable to restore original %s:// wrapper", protocol.val)
		return_value.u1.type_info = 2
		return
	}
	return_value.u1.type_info = 3
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_write", g.SizeOf("USERSTREAM_WRITE")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit((*byte)(buf), count, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG.exception != nil {
		return -1
	}
	if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 {
		if retval.u1.v.type_ == 2 {
			didwrite = -1
		} else {
			zend.ConvertToLong(&retval)
			didwrite = retval.value.lval
		}
	} else {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_write"+" is not implemented!", us.GetWrapper().GetClassname())
		didwrite = -1
	}

	/* don't allow strange buffer overruns due to bogus return */

	if didwrite > 0 && didwrite > count {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_write"+" wrote "+"%"+"lld"+" bytes more data than requested ("+"%"+"lld"+" written, "+"%"+"lld"+" max)", us.GetWrapper().GetClassname(), zend_long(didwrite-count), zend.ZendLong(didwrite), zend.ZendLong(count))
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_read", g.SizeOf("USERSTREAM_READ")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[0]
	__z.value.lval = count
	__z.u1.type_info = 4
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&func_name)
	if zend.EG.exception != nil {
		return -1
	}
	if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_read"+" is not implemented!", us.GetWrapper().GetClassname())
		return -1
	}
	if retval.u1.v.type_ == 2 {
		return -1
	}
	if zend.TryConvertToString(&retval) == 0 {
		return -1
	}
	didread = retval.value.str.len_
	if didread > 0 {
		if didread > count {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_read"+" - read "+"%"+"lld"+" bytes more data than requested ("+"%"+"lld"+" read, "+"%"+"lld"+" max) - excess data will be lost", us.GetWrapper().GetClassname(), zend_long(didread-count), zend.ZendLong(didread), zend.ZendLong(count))
			didread = count
		}
		memcpy(buf, retval.value.str.val, didread)
	}
	zend.ZvalPtrDtor(&retval)
	&retval.u1.type_info = 0

	/* since the user stream has no way of setting the eof flag directly, we need to ask it if we hit eof */

	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_eof", g.SizeOf("USERSTREAM_EOF")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	zend.ZvalPtrDtor(&func_name)
	if zend.EG.exception != nil {
		stream.eof = 1
		return -1
	}
	if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 && zend.ZendIsTrue(&retval) != 0 {
		stream.eof = 1
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_eof"+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_close", g.SizeOf("USERSTREAM_CLOSE")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&us.object)
	&us.object.u1.type_info = 0
	zend._efree(us)
	return 0
}
func PhpUserstreamopFlush(stream *core.PhpStream) int {
	var func_name zend.Zval
	var retval zend.Zval
	var call_result int
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	r.Assert(us != nil)
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_flush", g.SizeOf("USERSTREAM_FLUSH")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 && zend.ZendIsTrue(&retval) != 0 {
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_seek", g.SizeOf("USERSTREAM_SEEK")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zval = &args[0]
	__z.value.lval = offset
	__z.u1.type_info = 4
	var __z *zend.Zval = &args[1]
	__z.value.lval = whence
	__z.u1.type_info = 4
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 2, args, 0)
	zend.ZvalPtrDtor(&args[0])
	zend.ZvalPtrDtor(&args[1])
	zend.ZvalPtrDtor(&func_name)
	if call_result == zend.FAILURE {

		/* stream_seek is not implemented, so disable seeks for this stream */

		stream.flags |= 0x1

		/* there should be no retval to clean up */

		zend.ZvalPtrDtor(&retval)
		return -1
	} else if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 && zend.ZendIsTrue(&retval) != 0 {
		ret = 0
	} else {
		ret = -1
	}
	zend.ZvalPtrDtor(&retval)
	&retval.u1.type_info = 0
	if ret != 0 {
		return ret
	}

	/* now determine where we are */

	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_tell", g.SizeOf("USERSTREAM_TELL")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	if call_result == zend.SUCCESS && retval.u1.v.type_ == 4 {
		*newoffs = retval.value.lval
		ret = 0
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_tell"+" is not implemented!", us.GetWrapper().GetClassname())
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

	memset(ssb, 0, g.SizeOf("php_stream_statbuf"))
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "dev", g.SizeOf("\"dev\"")-1)) {
		ssb.sb.st_dev = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "ino", g.SizeOf("\"ino\"")-1)) {
		ssb.sb.st_ino = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "mode", g.SizeOf("\"mode\"")-1)) {
		ssb.sb.st_mode = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "nlink", g.SizeOf("\"nlink\"")-1)) {
		ssb.sb.st_nlink = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "uid", g.SizeOf("\"uid\"")-1)) {
		ssb.sb.st_uid = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "gid", g.SizeOf("\"gid\"")-1)) {
		ssb.sb.st_gid = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "rdev", g.SizeOf("\"rdev\"")-1)) {
		ssb.sb.st_rdev = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "size", g.SizeOf("\"size\"")-1)) {
		ssb.sb.st_size = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "atime", g.SizeOf("\"atime\"")-1)) {
		ssb.sb.st_atime = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "mtime", g.SizeOf("\"mtime\"")-1)) {
		ssb.sb.st_mtime = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "ctime", g.SizeOf("\"ctime\"")-1)) {
		ssb.sb.st_ctime = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "blksize", g.SizeOf("\"blksize\"")-1)) {
		ssb.sb.st_blksize = zend.ZvalGetLong(elem)
	}
	if nil != g.Assign(&elem, zend.ZendHashStrFind(array.value.arr, "blocks", g.SizeOf("\"blocks\"")-1)) {
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_stat", g.SizeOf("USERSTREAM_STAT")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	if call_result == zend.SUCCESS && retval.u1.v.type_ == 7 {
		if zend.SUCCESS == StatbufFromArray(&retval, ssb) {
			ret = 0
		}
	} else {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_stat"+" is not implemented!", us.GetWrapper().GetClassname())
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
	var ret int = -2
	var args []zend.Zval
	switch option {
	case 12:
		var __z *zend.Zval = &func_name
		var __s *zend.ZendString = zend.ZendStringInit("stream_eof", g.SizeOf("USERSTREAM_EOF")-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
		if call_result == zend.SUCCESS && (retval.u1.v.type_ == 2 || retval.u1.v.type_ == 3) {
			if zend.ZendIsTrue(&retval) != 0 {
				ret = -1
			} else {
				ret = 0
			}
		} else {
			ret = -1
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_eof"+" is not implemented! Assuming EOF", us.GetWrapper().GetClassname())
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
		break
	case 6:
		var __z *zend.Zval = &args[0]
		__z.value.lval = 0
		__z.u1.type_info = 4
		if (value & LOCK_NB) != 0 {
			&args[0].value.lval |= 4
		}
		switch value & ^LOCK_NB {
		case LOCK_SH:
			&args[0].value.lval |= 1
			break
		case LOCK_EX:
			&args[0].value.lval |= 2
			break
		case LOCK_UN:
			&args[0].value.lval |= 3
			break
		}

		/* TODO wouldblock */

		var __z *zend.Zval = &func_name
		var __s *zend.ZendString = zend.ZendStringInit("stream_lock", g.SizeOf("USERSTREAM_LOCK")-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0)
		if call_result == zend.SUCCESS && (retval.u1.v.type_ == 2 || retval.u1.v.type_ == 3) {
			ret = retval.u1.v.type_ == 2
		} else if call_result == zend.FAILURE {
			if value == 0 {

				/* lock support test (TODO: more check) */

				ret = 0

				/* lock support test (TODO: more check) */

			} else {
				core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_lock"+" is not implemented!", us.GetWrapper().GetClassname())
				ret = -1
			}
		}
		zend.ZvalPtrDtor(&retval)
		zend.ZvalPtrDtor(&func_name)
		zend.ZvalPtrDtor(&args[0])
		break
	case 10:
		var __z *zend.Zval = &func_name
		var __s *zend.ZendString = zend.ZendStringInit("stream_truncate", g.SizeOf("USERSTREAM_TRUNCATE")-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		switch value {
		case 0:
			if zend.ZendIsCallableEx(&func_name, g.CondF2(us.object.u1.v.type_ == 0, nil, func() *zend.ZendObject { return us.object.value.obj }), 1<<3, nil, nil, nil) != 0 {
				ret = 0
			} else {
				ret = -1
			}
			break
		case 1:
			var new_size ptrdiff_t = *((*ptrdiff_t)(ptrparam))
			if new_size >= 0 && new_size <= ptrdiff_t(2147483647) {
				var __z *zend.Zval = &args[0]
				__z.value.lval = zend.ZendLong(new_size)
				__z.u1.type_info = 4
				call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0)
				if call_result == zend.SUCCESS && retval.u1.v.type_ != 0 {
					if retval.u1.v.type_ == 2 || retval.u1.v.type_ == 3 {
						if retval.u1.v.type_ == 3 {
							ret = 0
						} else {
							ret = -1
						}
					} else {
						core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_truncate"+" did not return a boolean!", us.GetWrapper().GetClassname())
					}
				} else {
					core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_truncate"+" is not implemented!", us.GetWrapper().GetClassname())
				}
				zend.ZvalPtrDtor(&retval)
				zend.ZvalPtrDtor(&args[0])
			} else {
				ret = -1
			}
			break
		}
		zend.ZvalPtrDtor(&func_name)
		break
	case 2:

	case 3:

	case 4:

	case 1:
		var __z *zend.Zval = &func_name
		var __s *zend.ZendString = zend.ZendStringInit("stream_set_option", g.SizeOf("USERSTREAM_SET_OPTION")-1, 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		var __z *zend.Zval = &args[0]
		__z.value.lval = option
		__z.u1.type_info = 4
		&args[1].u1.type_info = 1
		&args[2].u1.type_info = 1
		switch option {
		case 2:

		case 3:
			var __z *zend.Zval = &args[1]
			__z.value.lval = value
			__z.u1.type_info = 4
			if ptrparam {
				var __z *zend.Zval = &args[2]
				__z.value.lval = *((*long)(ptrparam))
				__z.u1.type_info = 4
			} else {
				var __z *zend.Zval = &args[2]
				__z.value.lval = 1024
				__z.u1.type_info = 4
			}
			break
		case 4:
			var tv __struct__timeval = *((*__struct__timeval)(ptrparam))
			var __z *zval = &args[1]
			__z.value.lval = tv.tv_sec
			__z.u1.type_info = 4
			var __z *zend.Zval = &args[2]
			__z.value.lval = tv.tv_usec
			__z.u1.type_info = 4
			break
		case 1:
			var __z *zend.Zval = &args[1]
			__z.value.lval = value
			__z.u1.type_info = 4
			break
		default:
			break
		}
		call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 3, args, 0)
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_set_option"+" is not implemented!", us.GetWrapper().GetClassname())
			ret = -1
		} else if zend.ZendIsTrue(&retval) != 0 {
			ret = 0
		} else {
			ret = -1
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
	if object.u1.v.type_ == 0 {
		return ret
	}

	/* call the unlink method */

	var _s *byte = url
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var _s *byte = "unlink"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 1, args, 0)
	if call_result == zend.SUCCESS && (zretval.u1.v.type_ == 2 || zretval.u1.v.type_ == 3) {
		ret = zretval.u1.v.type_ == 3
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"unlink"+" is not implemented!", uwrap.GetClassname())
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
	if object.u1.v.type_ == 0 {
		return ret
	}

	/* call the rename method */

	var _s *byte = url_from
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var _s *byte = url_to
	var __z *zend.Zval = &args[1]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var _s *byte = "rename"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == zend.SUCCESS && (zretval.u1.v.type_ == 2 || zretval.u1.v.type_ == 3) {
		ret = zretval.u1.v.type_ == 3
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"rename"+" is not implemented!", uwrap.GetClassname())
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
	if object.u1.v.type_ == 0 {
		return ret
	}

	/* call the mkdir method */

	var _s *byte = url
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zval = &args[1]
	__z.value.lval = mode
	__z.u1.type_info = 4
	var __z *zend.Zval = &args[2]
	__z.value.lval = options
	__z.u1.type_info = 4
	var _s *byte = "mkdir"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == zend.SUCCESS && (zretval.u1.v.type_ == 2 || zretval.u1.v.type_ == 3) {
		ret = zretval.u1.v.type_ == 3
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"mkdir"+" is not implemented!", uwrap.GetClassname())
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
	if object.u1.v.type_ == 0 {
		return ret
	}

	/* call the rmdir method */

	var _s *byte = url
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[1]
	__z.value.lval = options
	__z.u1.type_info = 4
	var _s *byte = "rmdir"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == zend.SUCCESS && (zretval.u1.v.type_ == 2 || zretval.u1.v.type_ == 3) {
		ret = zretval.u1.v.type_ == 3
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"rmdir"+" is not implemented!", uwrap.GetClassname())
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
	case 1:
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = &args[2]
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		if value {
			var newtime *__struct__utimbuf = (*__struct__utimbuf)(value)
			zend.AddIndexLong(&args[2], 0, newtime.modtime)
			zend.AddIndexLong(&args[2], 1, newtime.actime)
		}
		break
	case 5:

	case 3:

	case 6:
		var __z *zend.Zval = &args[2]
		__z.value.lval = *((*long)(value))
		__z.u1.type_info = 4
		break
	case 4:

	case 2:
		var _s *byte = value
		var __z *zend.Zval = &args[2]
		var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
		__z.value.str = __s
		__z.u1.type_info = 6 | 1<<0<<8
		break
	default:
		core.PhpErrorDocref(nil, 1<<1, "Unknown option %d for "+"stream_metadata", option)
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* create an instance of our class */

	UserStreamCreateObject(uwrap, context, &object)
	if object.u1.v.type_ == 0 {
		zend.ZvalPtrDtor(&args[2])
		return ret
	}

	/* call the mkdir method */

	var _s *byte = url
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[1]
	__z.value.lval = option
	__z.u1.type_info = 4
	var _s *byte = "stream_metadata"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 3, args, 0)
	if call_result == zend.SUCCESS && (zretval.u1.v.type_ == 2 || zretval.u1.v.type_ == 3) {
		ret = zretval.u1.v.type_ == 3
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_metadata"+" is not implemented!", uwrap.GetClassname())
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
	if object.u1.v.type_ == 0 {
		return ret
	}

	/* call it's stat_url method - set up params first */

	var _s *byte = url
	var __z *zend.Zval = &args[0]
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	var __z *zend.Zval = &args[1]
	__z.value.lval = flags
	__z.u1.type_info = 4
	var _s *byte = "url_stat"
	var __z *zend.Zval = &zfuncname
	var __s *zend.ZendString = zend.ZendStringInit(_s, strlen(_s), 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(&object, &zfuncname, &zretval, 2, args, 0)
	if call_result == zend.SUCCESS && zretval.u1.v.type_ == 7 {

		/* We got the info we needed */

		if zend.SUCCESS == StatbufFromArray(&zretval, ssb) {
			ret = 0
		}

		/* We got the info we needed */

	} else {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"url_stat"+" is not implemented!", uwrap.GetClassname())
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

	if count != g.SizeOf("php_stream_dirent") {
		return -1
	}
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("dir_readdir", g.SizeOf("USERSTREAM_DIR_READ")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	if call_result == zend.SUCCESS && retval.u1.v.type_ != 2 && retval.u1.v.type_ != 3 {
		if &retval.u1.v.type_ != 6 {
			zend._convertToString(&retval)
		}
		var php_str_len int
		if retval.value.str.len_ >= g.SizeOf("ent -> d_name") {
			php_str_len = g.SizeOf("ent -> d_name") - 1
		} else {
			php_str_len = retval.value.str.len_
		}
		memcpy(ent.d_name, retval.value.str.val, php_str_len)
		ent.d_name[php_str_len] = '0'
		didread = g.SizeOf("php_stream_dirent")
	} else if call_result == zend.FAILURE {
		core.PhpErrorDocref(nil, 1<<1, "%s::"+"dir_readdir"+" is not implemented!", us.GetWrapper().GetClassname())
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("dir_closedir", g.SizeOf("USERSTREAM_DIR_CLOSE")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&us.object)
	&us.object.u1.type_info = 0
	zend._efree(us)
	return 0
}
func PhpUserstreamopRewinddir(stream *core.PhpStream, offset zend.ZendOffT, whence int, newoffs *zend.ZendOffT) int {
	var func_name zend.Zval
	var retval zend.Zval
	var us *PhpUserstreamDataT = (*PhpUserstreamDataT)(stream.abstract)
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("dir_rewinddir", g.SizeOf("USERSTREAM_DIR_REWIND")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 0, nil, 1)
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
	var __z *zend.Zval = &func_name
	var __s *zend.ZendString = zend.ZendStringInit("stream_cast", g.SizeOf("USERSTREAM_CAST")-1, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	switch castas {
	case 3:
		var __z *zend.Zval = &args[0]
		__z.value.lval = 3
		__z.u1.type_info = 4
		break
	default:
		var __z *zend.Zval = &args[0]
		__z.value.lval = 0
		__z.u1.type_info = 4
		break
	}
	call_result = zend._callUserFunctionEx(g.CondF2(us.object.u1.v.type_ == 0, nil, func() zend.Zval { return &us.object }), &func_name, &retval, 1, args, 0)
	for {
		if call_result == zend.FAILURE {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_cast"+" is not implemented!", us.GetWrapper().GetClassname())
			break
		}
		if zend.ZendIsTrue(&retval) == 0 {
			break
		}
		intstream = (*core.PhpStream)(zend.ZendFetchResource2Ex(&retval, "stream", PhpFileLeStream(), PhpFileLePstream()))
		if intstream == nil {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_cast"+" must return a stream resource", us.GetWrapper().GetClassname())
			break
		}
		if intstream == stream {
			core.PhpErrorDocref(nil, 1<<1, "%s::"+"stream_cast"+" must not return itself", us.GetWrapper().GetClassname())
			intstream = nil
			break
		}
		ret = _phpStreamCast(intstream, castas, retptr, 1)
		break
	}
	zend.ZvalPtrDtor(&retval)
	zend.ZvalPtrDtor(&func_name)
	zend.ZvalPtrDtor(&args[0])
	return ret
}

var PhpStreamUserspaceOps core.PhpStreamOps = core.PhpStreamOps{PhpUserstreamopWrite, PhpUserstreamopRead, PhpUserstreamopClose, PhpUserstreamopFlush, "user-space", PhpUserstreamopSeek, PhpUserstreamopCast, PhpUserstreamopStat, PhpUserstreamopSetOption}
var PhpStreamUserspaceDirOps core.PhpStreamOps = core.PhpStreamOps{nil, PhpUserstreamopReaddir, PhpUserstreamopClosedir, nil, "user-space-dir", PhpUserstreamopRewinddir, nil, nil, nil}
