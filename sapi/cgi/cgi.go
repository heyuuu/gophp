package cgi

import (
	"sik/core"
	"sik/zend"
)

var _ core.ISapiModule = (*CgiModuleType)(nil)

var CgiModule = &CgiModuleType{}

type CgiModuleType struct {
	core.BaseSapiModule
}

func (c *CgiModuleType) Name() string       { return "cgi-fcgi" }
func (c *CgiModuleType) PrettyName() string { return "CGI/FastCGI" }
func (c *CgiModuleType) Startup() bool {
	return core.PhpModuleStartupEx(c, []zend.ZendModuleEntry{CgiModuleEntry})
}

func (c *CgiModuleType) Shutdown() bool {
	core.PhpModuleShutdown()
	return true
}

func (c *CgiModuleType) Activate() {
	SapiCgiActivate()
}

func (c *CgiModuleType) Deactivate() {
	SapiCgiDeactivate()
}

func (c *CgiModuleType) UbWrite(str string) (int, error) {
	// todo
	l := SapiCgiUbWrite(str, len(str))
	return l, nil
}

func (c *CgiModuleType) Flush(serverContext any) {
	// SapiCgiFlush
	SapiCgiFlush(serverContext)
}

func (c *CgiModuleType) GetEnv(name string) (string, bool) {
	// SapiCgiGetenv()
	return getenv(name)
}

func (c *CgiModuleType) SendHeaders(headers *core.SapiHeaders) int {
	return SapiCgiSendHeaders(headers)
}

func (c *CgiModuleType) ReadPost(buffer *byte, count_bytes int) int {
	return SapiCgiReadPost(buffer, count_bytes)
}

func (c *CgiModuleType) ReadCookies() (string, bool) {
	// SapiCgiReadCookies()
	return getenv("HTTP_COOKIE")
}

func (c *CgiModuleType) RegisterServerVariables(trackVarsArray []zend.Zval) {
	SapiCgiRegisterVariables(trackVarsArray)
}

func (c *CgiModuleType) LogMessage(message string, syslogType int) {
	SapiCgiLogMessage(message, syslogType)
}
