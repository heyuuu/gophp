// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/core/streams"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/faults"
	"sik/zend/types"
)

func PhpFsockopenStream(executeData *zend.ZendExecuteData, return_value *types.Zval, persistent int) {
	var host *byte
	var host_len int
	var port zend.ZendLong = -1
	var zerrno *types.Zval = nil
	var zerrstr *types.Zval = nil
	var timeout float64 = float64(FG(default_socket_timeout))
	var conv int64
	var tv __struct__timeval
	var hashkey *byte = nil
	var stream *core.PhpStream = nil
	var err int
	var hostname *byte = nil
	var hostname_len int
	var errstr *types.String = nil
	return_value.SetFalse()
	for {
		var _flags int = 0
		var _min_num_args int = 1
		var _max_num_args int = 5

		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			host, host_len = fp.ParseString()
			fp.StartOptional()
			port = fp.ParseLong()
			zerrno = fp.ParseZval()
			zerrstr = fp.ParseZval()
			timeout = fp.ParseDouble()
			if fp.HasError() {
				fp.HandleError()
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if persistent != 0 {
		core.Spprintf(&hashkey, 0, "pfsockopen__%s:"+zend.ZEND_LONG_FMT, host, port)
	}
	if port > 0 {
		hostname_len = core.Spprintf(&hostname, 0, "%s:"+zend.ZEND_LONG_FMT, host, port)
	} else {
		hostname_len = host_len
		hostname = host
	}

	/* prepare the timeout value for use */

	conv = time_t(timeout * 1000000.0)
	tv.tv_sec = conv / 1000000
	tv.tv_usec = conv % 1000000
	stream = streams.PhpStreamXportCreate(hostname, hostname_len, core.REPORT_ERRORS, streams.STREAM_XPORT_CLIENT|streams.STREAM_XPORT_CONNECT, hashkey, &tv, nil, &errstr, &err)
	if port > 0 {
		zend.Efree(hostname)
	}
	if stream == nil {
		core.PhpErrorDocref(nil, faults.E_WARNING, "unable to connect to %s:"+zend.ZEND_LONG_FMT+" (%s)", host, port, b.CondF2(errstr == nil, "Unknown error", func() []byte { return errstr.GetVal() }))
	}
	if hashkey != nil {
		zend.Efree(hashkey)
	}
	if stream == nil {
		if zerrno != nil {
			zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, err)
		}
		if errstr != nil {
			if zerrstr != nil {
				zend.ZEND_TRY_ASSIGN_REF_STR(zerrstr, errstr)
			} else {
				types.ZendStringRelease(errstr)
			}
		}
		return_value.SetFalse()
		return
	}
	if zerrno != nil {
		zend.ZEND_TRY_ASSIGN_REF_LONG(zerrno, 0)
	}
	if zerrstr != nil {
		zend.ZEND_TRY_ASSIGN_REF_EMPTY_STRING(zerrstr)
	}
	if errstr != nil {
		types.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifFsockopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 0)
}
func ZifPfsockopen(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 1)
}
