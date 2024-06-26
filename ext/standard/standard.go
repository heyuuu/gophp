package standard

import (
	"github.com/heyuuu/gophp/php"
)

func init() {
	php.AddBuiltinModule(BasicModuleEntry)
}

var BasicModuleEntry = php.ModuleEntry{
	Name:            "standard",
	Functions:       zifFunctions,
	ModuleStartup:   ZmStartupBasic,
	ModuleShutdown:  ZmShutdownBasic,
	RequestStartup:  ZmActivateBasic,
	RequestShutdown: ZmDeactivateBasic,
}

func ZmStartupBasic(ctx *php.Context, moduleNumber int) bool {

	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_SCHEME", PHP_URL_SCHEME)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_HOST", PHP_URL_HOST)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_PORT", PHP_URL_PORT)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_USER", PHP_URL_USER)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_PASS", PHP_URL_PASS)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_PATH", PHP_URL_PATH)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_QUERY", PHP_URL_QUERY)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_URL_FRAGMENT", PHP_URL_FRAGMENT)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_QUERY_RFC1738", PHP_QUERY_RFC1738)
	php.RegisterLongConstant(ctx, moduleNumber, "PHP_QUERY_RFC3986", PHP_QUERY_RFC3986)

	RegisterStringConstants(ctx, moduleNumber)
	RegisterArrayConstants(ctx, moduleNumber)
	RegisterHtmlConstants(ctx, moduleNumber)

	return true
}

func ZmShutdownBasic(ctx *php.Context, moduleNumber int) bool {
	return true
}

func ZmActivateBasic(ctx *php.Context, moduleNumber int) bool {
	return true
}

func ZmDeactivateBasic(ctx *php.Context, moduleNumber int) bool {
	return true
}
