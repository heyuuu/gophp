// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/types"
)

func PhpStatpage() {
	var pstat *zend.ZendStatT
	pstat = core.SapiGetStat()
	if BG__().page_uid == -1 || BG__().page_gid == -1 {
		if pstat != nil {
			BG__().page_uid = pstat.st_uid
			BG__().page_gid = pstat.st_gid
			BG__().page_inode = pstat.st_ino
			BG__().page_mtime = pstat.st_mtime
		} else {
			BG__().page_uid = getuid()
			BG__().page_gid = getgid()
		}
	}
}
func PhpGetuid() zend.ZendLong {
	PhpStatpage()
	return BG__().page_uid
}
func PhpGetgid() zend.ZendLong {
	PhpStatpage()
	return BG__().page_gid
}
func ZifGetmyuid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var uid zend.ZendLong
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	uid = PhpGetuid()
	if uid < 0 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(uid)
		return
	}
}
func ZifGetmygid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var gid zend.ZendLong
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	gid = PhpGetgid()
	if gid < 0 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(gid)
		return
	}
}
func ZifGetmypid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var pid zend.ZendLong
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	pid = getpid()
	if pid < 0 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(pid)
		return
	}
}
func ZifGetmyinode(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	PhpStatpage()
	if BG__().page_inode < 0 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(BG__().page_inode)
		return
	}
}
func PhpGetlastmod() int64 {
	PhpStatpage()
	return BG__().page_mtime
}
func ZifGetlastmod(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var lm zend.ZendLong
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	lm = PhpGetlastmod()
	if lm < 0 {
		return_value.SetFalse()
		return
	} else {
		return_value.SetLong(lm)
		return
	}
}
