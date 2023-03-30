package cgi

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/types"
)

var _ core.ISapiModule = (*CgiModuleType)(nil)

var CgiModule = &CgiModuleType{}

type CgiModuleType struct {
	core.BaseSapiModule
	IsFastCgi bool
}

func (c *CgiModuleType) Name() string       { return "cgi-fcgi" }
func (c *CgiModuleType) PrettyName() string { return "CGI/FastCGI" }
func (c *CgiModuleType) Startup() bool {
	return core.PhpModuleStartupEx(c, []zend.ModuleEntry{CgiModuleEntry})
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
	if c.IsFastCgi {
		l := SapiFcgiUbWrite(str, len(str))
		return l, nil
	} else {
		l := SapiCgiUbWrite(str, len(str))
		return l, nil
	}
}

func (c *CgiModuleType) Flush(serverContext any) {
	if c.IsFastCgi {
		SapiFcgiFlush(serverContext)
	} else {
		SapiCgiFlush(serverContext)
	}
}

func (c *CgiModuleType) GetEnv(name string) (string, bool) {
	if c.IsFastCgi {
		return SapiFcgiGetenv(name)
	}

	// SapiCgiGetenv()
	return getenv(name)
}

func (c *CgiModuleType) SendHeaders(headers *core.SapiHeaders) int {
	return SapiCgiSendHeaders(headers)
}

func (c *CgiModuleType) ReadPost(buffer *byte, count_bytes int) int {
	if c.IsFastCgi {
		return SapiFcgiReadPost(buffer, count_bytes)
	}

	return SapiCgiReadPost(buffer, count_bytes)
}

func (c *CgiModuleType) ReadCookies() (string, bool) {
	if c.IsFastCgi {
		return SapiFcgiReadCookies()
	}

	// SapiCgiReadCookies()
	return getenv("HTTP_COOKIE")
}

func (c *CgiModuleType) RegisterServerVariables(trackVarsArray []types.Zval) {
	SapiCgiRegisterVariables(trackVarsArray)
}

func (c *CgiModuleType) LogMessage(message string, syslogType int) {
	SapiCgiLogMessage(message, syslogType)
}
