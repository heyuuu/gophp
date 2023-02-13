package cli

import (
	"fmt"
	"os"
	"sik/core"
	"sik/zend"
)

var CliServerModule = &CliServerModuleType{}

var _ ICliSapiModule = (*CliServerModuleType)(nil)

type CliServerModuleType struct {
	core.BaseSapiModule
}

func (c *CliServerModuleType) Name() string       { return "cli-server" }
func (c *CliServerModuleType) PrettyName() string { return "Built-in HTTP server" }
func (c *CliServerModuleType) Startup() bool {
	if !core.PhpModuleStartupEx(c, []zend.ZendModuleEntry{CliServerModuleEntry}) {
		return false
	}
	var workers = getenv("PHP_CLI_SERVER_WORKERS")
	if workers == "" {
		_, _ = fmt.Fprintf(os.Stderr, "platform does not support SO_REUSEPORT, cannot create workers\n")
	}
	return true
}

func (c *CliServerModuleType) Shutdown() bool {
	core.PhpModuleShutdown()
	return true
}

func (c *CliServerModuleType) Activate() {}

func (c *CliServerModuleType) Deactivate() {}

func (c *CliServerModuleType) UbWrite(str string) (int, error) {
	var client *PhpCliServerClient = core.SG__().server_context
	if client == nil {
		return 0, nil
	}
	return PhpCliServerClientSendThrough(client, str, len(str))
}

func (c *CliServerModuleType) Flush(serverContext any) {
	var client *PhpCliServerClient = serverContext
	if client == nil {
		return
	}
	if !(zend.ZEND_VALID_SOCKET(client.GetSock())) {
		core.PhpHandleAbortedConnection()
		return
	}
	if !(core.SG__().headers_sent) {
		core.SapiSendHeaders()
		core.SG__().headers_sent = 1
	}
}

func (c *CliServerModuleType) GetEnv(name string) (string, bool) {
	//TODO implement me
	panic("implement me")
}

func (c *CliServerModuleType) HeaderHandler(header *core.SapiHeader, op core.SapiHeaderOpEnum, headers *core.SapiHeaders) int {
	//TODO implement me
	panic("implement me")
}

func (c *CliServerModuleType) SendHeaders(headers *core.SapiHeaders) int {
	return SapiCliServerSendHeaders(headers)
}

func (c *CliServerModuleType) SendHeader(header *core.SapiHeader, serverContext any) {
	//TODO implement me
	panic("implement me")
}

func (c *CliServerModuleType) ReadPost(buffer *byte, count_bytes int) int {
	return SapiCliServerReadPost(buffer, count_bytes)
}

func (c *CliServerModuleType) ReadCookies() (string, bool) {
	val := SapiCliServerReadCookies
	return val, val != nil
}

func (c *CliServerModuleType) RegisterServerVariables(trackVarsArray []zend.Zval) {
	SapiCliServerRegisterVariables(trackVarsArray)
}

func (c *CliServerModuleType) LogMessage(message string, syslogType int) {
	SapiCliServerLogWrite(PHP_CLI_SERVER_LOG_MESSAGE, message)
}
