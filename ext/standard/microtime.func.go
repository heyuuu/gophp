package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/types"
	"sik/zend/zpp"
)

func _phpGettimeofday(executeData *zend.ZendExecuteData, return_value *types.Zval, mode int) {
	var get_as_float types.ZendBool = 0
	var tp __struct__timeval = __struct__timeval{0}
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			get_as_float = fp.ParseBool()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if gettimeofday(&tp, nil) {
		b.Assert(false)
	}
	if get_as_float != 0 {
		return_value.SetDouble(float64(tp.tv_sec + tp.tv_usec/MICRO_IN_SEC))
		return
	}
	if mode != 0 {
		var offset *timelib_time_offset
		offset = timelib_get_time_zone_info(tp.tv_sec, get_timezone_info())
		zend.ArrayInit(return_value)
		zend.AddAssocLong(return_value, "sec", tp.tv_sec)
		zend.AddAssocLong(return_value, "usec", tp.tv_usec)
		zend.AddAssocLong(return_value, "minuteswest", -(offset.offset)/SEC_IN_MIN)
		zend.AddAssocLong(return_value, "dsttime", offset.is_dst)
		timelib_time_offset_dtor(offset)
	} else {
		return_value.SetString(zend.ZendStrpprintf(0, "%.8F %ld", tp.tv_usec/MICRO_IN_SEC, long(tp.tv_sec)))
		return
	}
}
func ZifMicrotime(executeData zpp.DefEx, return_value zpp.DefReturn, _ zpp.DefOpt, getAsFloat *types.Zval) {
	_phpGettimeofday(executeData, return_value, 0)
}
func ZifGettimeofday(executeData zpp.DefEx, return_value zpp.DefReturn, _ zpp.DefOpt, getAsFloat *types.Zval) {
	_phpGettimeofday(executeData, return_value, 1)
}
func ZifGetrusage(executeData zpp.DefEx, return_value zpp.DefReturn, _ zpp.DefOpt, who *types.Zval) {
	var usg __struct__rusage
	var pwho zend.ZendLong = 0
	var who int = RUSAGE_SELF
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 1, 0)
			fp.StartOptional()
			pwho = fp.ParseLong()
			if fp.HasError() {
				return
			}
			break
		}
		break
	}
	if pwho == 1 {
		who = RUSAGE_CHILDREN
	}
	memset(&usg, 0, b.SizeOf("struct rusage"))
	if getrusage(who, &usg) == -1 {
		return_value.SetFalse()
		return
	}
	zend.ArrayInit(return_value)

	// #define PHP_RUSAGE_PARA(a) add_assoc_long ( return_value , # a , usg . a )

	zend.AddAssocLong(return_value, "ru_oublock", usg.ru_oublock)
	zend.AddAssocLong(return_value, "ru_inblock", usg.ru_inblock)
	zend.AddAssocLong(return_value, "ru_msgsnd", usg.ru_msgsnd)
	zend.AddAssocLong(return_value, "ru_msgrcv", usg.ru_msgrcv)
	zend.AddAssocLong(return_value, "ru_maxrss", usg.ru_maxrss)
	zend.AddAssocLong(return_value, "ru_ixrss", usg.ru_ixrss)
	zend.AddAssocLong(return_value, "ru_idrss", usg.ru_idrss)
	zend.AddAssocLong(return_value, "ru_minflt", usg.ru_minflt)
	zend.AddAssocLong(return_value, "ru_majflt", usg.ru_majflt)
	zend.AddAssocLong(return_value, "ru_nsignals", usg.ru_nsignals)
	zend.AddAssocLong(return_value, "ru_nvcsw", usg.ru_nvcsw)
	zend.AddAssocLong(return_value, "ru_nivcsw", usg.ru_nivcsw)
	zend.AddAssocLong(return_value, "ru_nswap", usg.ru_nswap)
	zend.AddAssocLong(return_value, "ru_utime . tv_usec", usg.ru_utime.tv_usec)
	zend.AddAssocLong(return_value, "ru_utime . tv_sec", usg.ru_utime.tv_sec)
	zend.AddAssocLong(return_value, "ru_stime . tv_usec", usg.ru_stime.tv_usec)
	zend.AddAssocLong(return_value, "ru_stime . tv_sec", usg.ru_stime.tv_sec)
}
