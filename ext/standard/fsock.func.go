package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/core/streams"
	b "github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/faults"
	"github.com/heyuuu/gophp/zend/zpp"
)

func PhpFsockopenStream(executeData *zend.ZendExecuteData, return_value *types.Zval, persistent int) {
	var host *byte
	var host_len int
	var port zend.ZendLong = -1
	var zerrno *types.Zval = nil
	var zerrstr *types.Zval = nil
	var timeout float64 = float64(FG__().default_socket_timeout)
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
		for {
			fp := zpp.FastParseStart(executeData, 1, 5, 0)
			host, host_len = fp.ParseString()
			fp.StartOptional()
			port = fp.ParseLong()
			zerrno = fp.ParseZval()
			zerrstr = fp.ParseZval()
			timeout = fp.ParseDouble()
			if fp.HasError() {
				return_value.SetFalse()
				return
			}
			break
		}
		break
	}
	if persistent != 0 {
		hashkey = fmt.Sprintf("pfsockopen__%s:%d", host, port)
	}
	if port > 0 {
		hostname = fmt.Sprintf("%s:%d", host, port)
		hostname_len = len(hostname)
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
		core.PhpErrorDocref("", faults.E_WARNING, fmt.Sprintf("unable to connect to %s:%d (%s)", host, port, b.CondF2(errstr == nil, "Unknown error", func() string { return errstr.GetStr() })))
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
				// types.ZendStringRelease(errstr)
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
		// types.ZendStringReleaseEx(errstr, 0)
	}
	core.PhpStreamToZval(stream, return_value)
}
func ZifFsockopen(executeData zpp.Ex, return_value zpp.Ret, hostname *types.Zval, _ zpp.Opt, port *types.Zval, errno zpp.RefZval, errstr zpp.RefZval, timeout *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 0)
}
func ZifPfsockopen(executeData zpp.Ex, return_value zpp.Ret, hostname *types.Zval, _ zpp.Opt, port *types.Zval, errno zpp.RefZval, errstr zpp.RefZval, timeout *types.Zval) {
	PhpFsockopenStream(executeData, return_value, 1)
}
