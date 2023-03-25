package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

func PhpStdDate(t int64) *byte {
	var tm1 *__struct__tm
	var tmbuf __struct__tm
	var str *byte
	tm1 = core.PhpGmtimeR(&t, &tmbuf)
	str = zend.Emalloc(81)
	str[0] = '0'
	if tm1 == nil {
		return str
	}
	core.Snprintf(str, 80, "%s, %02d %s %04d %02d:%02d:%02d GMT", DayShortNames[tm1.tm_wday], tm1.tm_mday, MonShortNames[tm1.tm_mon], tm1.tm_year+1900, tm1.tm_hour, tm1.tm_min, tm1.tm_sec)
	str[79] = 0
	return str
}
func ZifStrptime(executeData zpp.Ex, return_value zpp.Ret, timestamp *types.Zval, format *types.Zval) {
	var ts *byte
	var ts_length int
	var format *byte
	var format_length int
	var parsed_time __struct__tm
	var unparsed_part *byte
	for {
		for {
			fp := zpp.FastParseStart(executeData, 2, 2, 0)
			ts, ts_length = fp.ParseString()
			format, format_length = fp.ParseString()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	memset(&parsed_time, 0, b.SizeOf("parsed_time"))
	unparsed_part = strptime(ts, format, &parsed_time)
	if unparsed_part == nil {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)
	zend.AddAssocLong(return_value, "tm_sec", parsed_time.tm_sec)
	zend.AddAssocLong(return_value, "tm_min", parsed_time.tm_min)
	zend.AddAssocLong(return_value, "tm_hour", parsed_time.tm_hour)
	zend.AddAssocLong(return_value, "tm_mday", parsed_time.tm_mday)
	zend.AddAssocLong(return_value, "tm_mon", parsed_time.tm_mon)
	zend.AddAssocLong(return_value, "tm_year", parsed_time.tm_year)
	zend.AddAssocLong(return_value, "tm_wday", parsed_time.tm_wday)
	zend.AddAssocLong(return_value, "tm_yday", parsed_time.tm_yday)
	zend.AddAssocString(return_value, "unparsed", unparsed_part)
}
