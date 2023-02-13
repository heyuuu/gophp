package cli

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sik/core"
	r "sik/runtime"
	"sik/zend"
)

var CliModule = &CliModuleType{}

var _ core.ISapiModule = (*CliModuleType)(nil)

type CliModuleType struct {
	core.BaseSapiModule
}

func (c *CliModuleType) Name() string       { return "cli" }
func (c *CliModuleType) PrettyName() string { return "Command Line Interface" }
func (c *CliModuleType) Startup() bool {
	return core.PhpModuleStartupEx(c, nil)
}
func (c *CliModuleType) Shutdown() bool {
	core.PhpModuleShutdown()
	return true
}
func (c *CliModuleType) Activate() {}
func (c *CliModuleType) Deactivate() {
	r.Fflush(stdout)
	if core.SG__().request_info.argv0 {
		zend.Free(core.SG__().request_info.argv0)
		core.SG__().request_info.argv0 = nil
	}
}
func (c *CliModuleType) UbWrite(str string) (int, error) {
	if len(str) == 0 {
		return 0, nil
	}

	file := os.Stdout
	pos := 0
	end := len(str)
	for pos < end {
		count, err := file.WriteString(str[pos:])
		if count > 0 {
			pos += count
		}
		if err != nil && !errors.Is(err, io.ErrShortWrite) {
			// todo
			zend.EG__().SetExitStatus(255)
			core.PhpHandleAbortedConnection()
			return pos, err
		}
	}

	return pos, nil
}
func (c *CliModuleType) Flush(serverContext any) {}

func (c *CliModuleType) GetEnv(name string) (string, bool) {
	return "", false
}

func (c *CliModuleType) HeaderHandler(header *core.SapiHeader, op core.SapiHeaderOpEnum, headers *core.SapiHeaders) int {
	return 0
}

func (c *CliModuleType) SendHeaders(headers *core.SapiHeaders) int {
	/* We do nothing here, this function is needed to prevent that the fallback
	 * header handling is called. */
	return core.SAPI_HEADER_SENT_SUCCESSFULLY
}

func (c *CliModuleType) SendHeader(header *core.SapiHeader, serverContext any) {}

func (c *CliModuleType) ReadPost(buffer *byte, count_bytes int) int {
	//TODO implement me
	panic("implement me")
}

func (c *CliModuleType) ReadCookies() (string, bool) { return "", false }

func (c *CliModuleType) RegisterServerVariables(trackVarsArray []zend.Zval) {
	SapiCliRegisterVariables(trackVarsArray)
}

func (c *CliModuleType) LogMessage(message string, syslogType int) {
	fmt.Fprintf(os.Stderr, "%s\n", message)
}

func (c *CliModuleType) InputFilter(arg int, name string, value string) string {
	//TODO implement me
	panic("implement me")
}

func (c *CliModuleType) GetStat() bool {

	//TODO implement me
	panic("implement me")
}
