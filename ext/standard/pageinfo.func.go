// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
)

func PhpStatpage() {
	var pstat *zend.ZendStatT
	pstat = core.SapiGetStat()
	if BG(page_uid) == -1 || BG(page_gid) == -1 {
		if pstat != nil {
			BG(page_uid) = pstat.st_uid
			BG(page_gid) = pstat.st_gid
			BG(page_inode) = pstat.st_ino
			BG(page_mtime) = pstat.st_mtime
		} else {
			BG(page_uid) = getuid()
			BG(page_gid) = getgid()
		}
	}
}
func PhpGetuid() zend.ZendLong {
	PhpStatpage()
	return BG(page_uid)
}
func PhpGetgid() zend.ZendLong {
	PhpStatpage()
	return BG(page_gid)
}
func ZifGetmyuid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var uid zend.ZendLong
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	uid = PhpGetuid()
	if uid < 0 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(uid)
		return
	}
}
func ZifGetmygid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var gid zend.ZendLong
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	gid = PhpGetgid()
	if gid < 0 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(gid)
		return
	}
}
func ZifGetmypid(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var pid zend.ZendLong
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	pid = getpid()
	if pid < 0 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(pid)
		return
	}
}
func ZifGetmyinode(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	PhpStatpage()
	if BG(page_inode) < 0 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(BG(page_inode))
		return
	}
}
func PhpGetlastmod() int64 {
	PhpStatpage()
	return BG(page_mtime)
}
func ZifGetlastmod(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var lm zend.ZendLong
	if zend.ZendParseParametersNone() == zend.FAILURE {
		return
	}
	lm = PhpGetlastmod()
	if lm < 0 {
		zend.RETVAL_FALSE
		return
	} else {
		zend.RETVAL_LONG(lm)
		return
	}
}
