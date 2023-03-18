// <<generate>>

package standard

import (
	"sik/core"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func ZifUniqid(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var prefix *byte = ""
	var more_entropy types.ZendBool = 0
	var uniqid *types.ZendString
	var sec int
	var usec int
	var prefix_len int = 0
	var tv __struct__timeval
	for {
		var _flags int = 0
		var _min_num_args int = 0
		var _max_num_args int = 2
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			fp.StartOptional()
			prefix, prefix_len = fp.ParseString()
			more_entropy = fp.ParseBool()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}

	/* This implementation needs current microsecond to change,
	 * hence we poll time until it does. This is much faster than
	 * calling usleep(1) which may cause the kernel to schedule
	 * another process, causing a pause of around 10ms.
	 */

	for {
		void(gettimeofday((*__struct__timeval)(&tv), (*__struct__timezone)(nil)))
		if !(tv.tv_sec == PrevTv.tv_sec && tv.tv_usec == PrevTv.tv_usec) {
			break
		}
	}
	PrevTv.tv_sec = tv.tv_sec
	PrevTv.tv_usec = tv.tv_usec
	sec = int(tv.tv_sec)
	usec = int(tv.tv_usec % 0x100000)

	/* The max value usec can have is 0xF423F, so we use only five hex
	 * digits for usecs.
	 */

	if more_entropy != 0 {
		uniqid = core.Strpprintf(0, "%s%08x%05x%.8F", prefix, sec, usec, PhpCombinedLcg()*10)
	} else {
		uniqid = core.Strpprintf(0, "%s%08x%05x", prefix, sec, usec)
	}
	return_value.SetString(uniqid)
	return
}
