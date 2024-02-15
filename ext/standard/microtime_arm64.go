package standard

import (
	"github.com/heyuuu/gophp/php/types"
	"syscall"
)

func init() {
	getrusgae = getrusageArm64
}

func getrusageArm64(mode int) (*types.Array, bool) {
	var who int = syscall.RUSAGE_SELF
	if mode == 1 {
		who = syscall.RUSAGE_CHILDREN
	}

	var usg syscall.Rusage
	if err := syscall.Getrusage(who, &usg); err != nil {
		return nil, false
	}

	arr := types.NewArray()

	arr.AddAssocLong("ru_oublock", int(usg.Oublock))
	arr.AddAssocLong("ru_inblock", int(usg.Inblock))
	arr.AddAssocLong("ru_msgsnd", int(usg.Msgsnd))
	arr.AddAssocLong("ru_msgrcv", int(usg.Msgrcv))
	arr.AddAssocLong("ru_maxrss", int(usg.Maxrss))
	arr.AddAssocLong("ru_ixrss", int(usg.Ixrss))
	arr.AddAssocLong("ru_idrss", int(usg.Idrss))
	arr.AddAssocLong("ru_minflt", int(usg.Minflt))
	arr.AddAssocLong("ru_majflt", int(usg.Majflt))
	arr.AddAssocLong("ru_nsignals", int(usg.Nsignals))
	arr.AddAssocLong("ru_nvcsw", int(usg.Nvcsw))
	arr.AddAssocLong("ru_nivcsw", int(usg.Nivcsw))
	arr.AddAssocLong("ru_nswap", int(usg.Nswap))
	arr.AddAssocLong("ru_utime.tv_usec", int(usg.Utime.Usec))
	arr.AddAssocLong("ru_utime.tv_sec", int(usg.Utime.Sec))
	arr.AddAssocLong("ru_stime.tv_usec", int(usg.Stime.Usec))
	arr.AddAssocLong("ru_stime.tv_sec", int(usg.Stime.Sec))

	return arr, true
}
