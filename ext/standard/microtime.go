package standard

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"time"
)

func ZifMicrotime(_ zpp.Opt, getAsFloat bool) types.Zval {
	t := time.Now()
	if getAsFloat {
		v := float64(t.UnixNano()) / float64(time.Second)
		return php.Double(v)
	}

	s := fmt.Sprintf("%.8f %d", float64(t.Nanosecond())/float64(time.Second), t.Unix())
	return php.String(s)
}
func ZifGettimeofday(_ zpp.Opt, getAsFloat bool) types.Zval {
	t := time.Now()
	if getAsFloat {
		v := float64(t.UnixNano()) / float64(time.Second)
		return php.Double(v)
	}

	arr := types.NewArrayCap(4)
	arr.KeyUpdate("sec", php.Long(int(t.Unix())))
	arr.KeyUpdate("usec", php.Long(t.Nanosecond()/1000))
	// todo 修正此处的时区问题
	arr.KeyUpdate("minuteswest", php.Long(0))
	arr.KeyUpdate("dsttime", php.Long(0))

	return php.Array(arr)
}

var getrusgae func(mode int) (*types.Array, bool)

func ZifGetrusage(_ zpp.Opt, mode int) (*types.Array, bool) {
	return getrusgae(mode)
}
