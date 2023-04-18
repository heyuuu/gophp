package streams

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
)

var UrlStreamWrappersHash map[string]*core.PhpStreamWrapper

var LeStream int = types.FAILURE
var LePstream int = types.FAILURE
var LeStreamFilter int = types.FAILURE
