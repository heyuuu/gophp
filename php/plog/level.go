package plog

import "log/syslog"

type Level int

const (
	Emergency Level = Level(syslog.LOG_EMERG)
	Alert     Level = Level(syslog.LOG_ALERT)
	Critical  Level = Level(syslog.LOG_CRIT)
	Error     Level = Level(syslog.LOG_ERR)
	Warning   Level = Level(syslog.LOG_WARNING)
	Notice    Level = Level(syslog.LOG_NOTICE)
	Info      Level = Level(syslog.LOG_INFO)
	Debug     Level = Level(syslog.LOG_DEBUG)
)
