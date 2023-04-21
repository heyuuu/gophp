package types

import "github.com/heyuuu/gophp/zend"

type UserCallable struct {
	Info  ZendFcallInfo
	Cache ZendFcallInfoCache
}

func (c *UserCallable) Call(args ...*Zval) (*Zval, bool) {
	// duplicate args
	realArgs := make([]Zval, len(args))
	for i, arg := range args {
		realArgs[i].CopyFrom(arg)
	}

	// init ZendFcallInfo
	var retval Zval
	c.Info.SetParamCount(uint32(len(args)))
	c.Info.SetParams(realArgs)
	c.Info.SetRetval(&retval)
	c.Info.SetNoSeparation(0)

	// call
	ret := zend.ZendCallFunction(&c.Info, &c.Cache)
	if ret != SUCCESS {
		return nil, false
	}
	return &retval, true
}
