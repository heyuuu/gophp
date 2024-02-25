package streams

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"time"
)

type StreamFilter struct{}

// StreamFilterFactory
type StreamFilterFactory interface {
	Create(filterName string, filterParams *types.Zval) *StreamFilter
}

type StreamFilterFactoryFunc func(filterName string, filterParams *types.Zval) *StreamFilter

func (f StreamFilterFactoryFunc) Create(filterName string, filterParams *types.Zval) *StreamFilter {
	return f(filterName, filterParams)
}

// TransportFactory
type TransportFactory func(proto string, resource string, persistentId string, options int, flags int, timeout time.Duration, context *StreamContext) *Stream

func wrapperSchemeValidate(protocol string) bool {
	for _, c := range []byte(protocol) {
		if !ascii.IsAlphaNum(c) && c != '+' && c != '-' && c != '.' {
			return false
		}
	}
	return true
}

func streamResourceRegularDtor(res *types.Resource) {
}
