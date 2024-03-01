package streams

import (
	"github.com/heyuuu/gophp/kits/mapkit"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/assert"
	"github.com/heyuuu/gophp/php/types"
	"maps"
)

var (
	LeStream  = php.InitResourceType("stream")
	LePstream = php.InitResourceType("persistent stream")
	/* Filters are cleaned up by the streams they're attached to */
	LeStreamFilter  = php.InitResourceType("stream filter")
	LeStreamContext = php.InitResourceType("stream-context")
)

var streamGlobalsSharedKey = "streams.globals_shared"
var streamGlobalsKey = "streams.globals"

func StreamG(ctx *php.Context) *StreamGlobals {
	return php.ContextGetOrInitPersistent(ctx, streamGlobalsKey, NewStreamGlobals)
}

func SharedStreamG(ctx *php.Context) *StreamGlobalsShared {
	return php.ContextGetOrInitPersistent(ctx, streamGlobalsSharedKey, initStreamGlobalsShared)
}

func InitSharedStreamGlobals(ctx *php.Context) {
	// assert 必须在具体请求前初始化 StreamGlobals
	assert.Assert(ctx == ctx.Engine().BaseCtx())
	// 注册资源类型
	ctx.Engine().RegisterResourceType(LeStream)
	ctx.Engine().RegisterResourceType(LePstream)
	ctx.Engine().RegisterResourceType(LeStreamFilter)
	ctx.Engine().RegisterResourceType(LeStreamContext)
	// 触发初始化
	SharedStreamG(ctx)
}

// StreamGlobalsShared: 各请求复用的数据，engine范围内唯一
type StreamGlobalsShared struct {
	wrappers map[string]*StreamWrapper      // @see: static HashTable url_stream_wrappers_hash
	filters  map[string]StreamFilterFactory // @see: static HashTable stream_filters_hash
	xports   *types.Table[TransportFactory] // @see: static HashTable xport_hash
}

// @see: php_init_stream_wrappers
func initStreamGlobalsShared() *StreamGlobalsShared {
	g := &StreamGlobalsShared{
		wrappers: map[string]*StreamWrapper{},
		filters:  map[string]StreamFilterFactory{},
		xports:   types.NewTable[TransportFactory](),
	}

	// wrappers
	//g.wrappers["php"] = standard.PhpStreamPhpWrapper
	//g.wrappers["file"] = PhpPlainFilesWrapper
	//g.wrappers["glob"] = PhpGlobStreamWrapper
	//g.wrappers["data"] = PhpStreamRfc2397Wrapper
	//g.wrappers["http"] = standard.PhpStreamHttpWrapper
	//g.wrappers["ftp"] = standard.PhpStreamFtpWrapper

	// filters
	//g.filters = maps.Clone(standard.StandardFiltersMap)

	// xports
	//g.xports.Set("tcp", StreamGenericSocketFactory)
	//g.xports.Set("udp", StreamGenericSocketFactory)

	return g
}

// StreamGlobals: 各请求单独的数据，context范围内唯一，context间不复用
type StreamGlobals struct {
	shared *StreamGlobalsShared

	volatileWrappers map[string]*StreamWrapper // @see: FG(stream_wrappers)
	volatileFilters  map[string]StreamFilterFactory
}

func NewStreamGlobals() *StreamGlobals {
	return &StreamGlobals{}
}

func (g *StreamGlobals) XportNames() []string {
	return g.shared.xports.Keys()
}
func (g *StreamGlobals) XportGet(protocol string) TransportFactory {
	return g.shared.xports.Get(protocol)
}

// wrappers
func (g *StreamGlobals) WrapperNames() []string {
	if g.volatileWrappers != nil {
		return mapkit.Keys(g.volatileWrappers)
	} else {
		return mapkit.Keys(g.shared.wrappers)
	}
}

func (g *StreamGlobals) WrapperGet(protocol string) *StreamWrapper {
	if g.volatileWrappers != nil {
		return g.volatileWrappers[protocol]
	} else {
		return g.shared.wrappers[protocol]
	}
}

func (g *StreamGlobals) HasVolatileWrappers() bool {
	return g.volatileWrappers != nil
}

// @see: php_register_url_stream_wrapper_volatile
func (g *StreamGlobals) WrapperRegisterVolatile(protocol string, wrapper *StreamWrapper) bool {
	if !wrapperSchemeValidate(protocol) {
		return false
	}
	if g.volatileWrappers == nil {
		g.volatileWrappers = maps.Clone(g.volatileWrappers)
	}
	if _, exists := g.volatileWrappers[protocol]; exists {
		return false
	} else {
		g.volatileWrappers[protocol] = wrapper
		return true
	}
}

// @see: php_unregister_url_stream_wrapper_volatile
func (g *StreamGlobals) WrapperUnregisterVolatile(protocol string) bool {
	if g.volatileWrappers == nil {
		g.volatileWrappers = maps.Clone(g.volatileWrappers)
	}
	if _, exists := g.volatileWrappers[protocol]; exists {
		delete(g.volatileWrappers, protocol)
		return true
	} else {
		return false
	}
}

func (g *StreamGlobals) WrapperRestore(protocol string) (existed bool, hasChanged bool) {
	rawWrapper := g.shared.wrappers[protocol]
	if rawWrapper == nil {
		return false, false
	}

	if g.volatileWrappers == nil || g.volatileWrappers[protocol] == rawWrapper {
		return true, false
	}

	if g.volatileWrappers != nil {
		g.volatileWrappers[protocol] = rawWrapper
	}

	return true, true
}

// filters
func (g *StreamGlobals) FilterNames() []string {
	if g.volatileFilters != nil {
		return mapkit.Keys(g.volatileFilters)
	} else {
		return mapkit.Keys(g.shared.filters)
	}
}

func (g *StreamGlobals) FilterHash() map[string]StreamFilterFactory {
	if g.volatileFilters != nil {
		return g.volatileFilters
	} else {
		return g.shared.filters
	}
}

func (g *StreamGlobals) FilterRegisterVolatile(filterPattern string, factory StreamFilterFactory) bool {
	if g.volatileFilters == nil {
		g.volatileFilters = maps.Clone(g.shared.filters)
	}

	if _, exists := g.volatileFilters[filterPattern]; exists {
		return false
	} else {
		g.volatileFilters[filterPattern] = factory
		return true
	}
}
