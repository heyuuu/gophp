package standard

import (
	"sik/core"
	"sik/zend/types"
	"sik/zend/zpp"
)

func ZifUniqid(executeData zpp.Ex, return_value zpp.Ret, _ zpp.Opt, prefix *types.Zval, moreEntropy *types.Zval) {
	var prefix *byte = ""
	var more_entropy types.ZendBool = 0
	var uniqid *types.String
	var sec int
	var usec int
	var prefix_len int = 0
	var tv __struct__timeval
	for {
		for {
			fp := zpp.FastParseStart(executeData, 0, 2, 0)
			fp.StartOptional()
			prefix, prefix_len = fp.ParseString()
			more_entropy = fp.ParseBool()
			if fp.HasError() {
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
