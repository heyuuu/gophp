// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZmStartupSyslog(type_ int, module_number int) int {
	/* error levels */

	zend.RegisterLongConstant("LOG_EMERG", LOG_EMERG, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_ALERT", LOG_ALERT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_CRIT", LOG_CRIT, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_ERR", LOG_ERR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_WARNING", LOG_WARNING, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_NOTICE", LOG_NOTICE, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_INFO", LOG_INFO, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_DEBUG", LOG_DEBUG, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)

	/* facility: type of program logging the message */

	zend.RegisterLongConstant("LOG_KERN", LOG_KERN, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_USER", LOG_USER, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_MAIL", LOG_MAIL, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_DAEMON", LOG_DAEMON, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_AUTH", LOG_AUTH, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_SYSLOG", LOG_SYSLOG, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LPR", LOG_LPR, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL0", LOG_LOCAL0, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL1", LOG_LOCAL1, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL2", LOG_LOCAL2, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL3", LOG_LOCAL3, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL4", LOG_LOCAL4, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL5", LOG_LOCAL5, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL6", LOG_LOCAL6, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_LOCAL7", LOG_LOCAL7, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)

	/* options */

	zend.RegisterLongConstant("LOG_PID", LOG_PID, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_CONS", LOG_CONS, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_ODELAY", LOG_ODELAY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	zend.RegisterLongConstant("LOG_NDELAY", LOG_NDELAY, zend.CONST_CS|zend.CONST_PERSISTENT, module_number)
	BG__().syslog_device = nil
	return types.SUCCESS
}
func ZmActivateSyslog(type_ int, module_number int) int {
	BG__().syslog_device = nil
	return types.SUCCESS
}
func ZmShutdownSyslog(type_ int, module_number int) int {
	if BG__().syslog_device {
		zend.Free(BG__().syslog_device)
		BG__().syslog_device = nil
	}
	return types.SUCCESS
}
func PhpOpenlog(ident *byte, option int, facility int) {
	openlog(ident, option, facility)
	core.PG__().have_called_openlog = 1
}
func ZifOpenlog(executeData zpp.DefEx, return_value zpp.DefReturn, ident *types.Zval, option *types.Zval, facility *types.Zval) {
	var ident *byte
	var option zend.ZendLong
	var facility zend.ZendLong
	var ident_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 3, 3, 0)
			ident, ident_len = fp.ParseString()
			option = fp.ParseLong()
			facility = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if BG__().syslog_device {
		zend.Free(BG__().syslog_device)
	}
	BG__().syslog_device = zend.ZendStrndup(ident, ident_len)
	if BG__().syslog_device == nil {
		return_value.SetFalse()
		return
	}
	PhpOpenlog(BG__().syslog_device, option, facility)
	return_value.SetTrue()
	return
}
func ZifCloselog(executeData zpp.DefEx, return_value zpp.DefReturn) {
	if !executeData.CheckNumArgsNone(false) {
		return
	}
	closelog()
	if BG__().syslog_device {
		zend.Free(BG__().syslog_device)
		BG__().syslog_device = nil
	}
	return_value.SetTrue()
	return
}
func ZifSyslog(executeData zpp.DefEx, return_value zpp.DefReturn, priority *types.Zval, message *types.Zval) {
	var priority zend.ZendLong
	var message *byte
	var message_len int
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			priority = fp.ParseLong()
			message, message_len = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	core.PhpSyslog(priority, "%s", message)
	return_value.SetTrue()
	return
}
