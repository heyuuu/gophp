package cli

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

const SOCK_EINVAL = EINVAL
const SOCK_EAGAIN = EAGAIN
const SOCK_EINTR = EINTR
const SOCK_EADDRINUSE = EADDRINUSE

var PhpCliServerMaster pid_t
var PhpCliServerWorkers *pid_t
var PhpCliServerWorkersMax zend.ZendLong

var TemplateMap = map[int]string{
	400: "<h1>%s</h1><p>Your browser sent a request that this server could not understand.</p>",
	404: "<h1>%s</h1><p>The requested resource <code class=\"url\">%s</code> was not found on this server.</p>",
	500: "<h1>%s</h1><p>The server is temporarily unavailable.</p>",
	501: "<h1>%s</h1><p>Request method not supported.</p>",
}

const PHP_CLI_SERVER_LOG_PROCESS = 1
const PHP_CLI_SERVER_LOG_ERROR = 2
const PHP_CLI_SERVER_LOG_MESSAGE = 3

var PhpCliServerLogLevel int = 3

var ServerAdditionalFunctions = []types.FunctionEntry{
	DefZifApacheRequestHeaders,
	DefZifApacheResponseHeaders,
	DefZifGetallheaders,
}

// CliModuleData
type CliModuleData struct{}

var _ zend.ModuleData = (*CliModuleData)(nil)

func (d *CliModuleData) Name() string                     { return "cli_server" }
func (d *CliModuleData) Version() string                  { return core.PHP_VERSION }
func (d *CliModuleData) Functions() []types.FunctionEntry { return nil }
func (d *CliModuleData) ModuleStartup(moduleNumber int) bool {
	return ZmStartupCliServer(0, moduleNumber) == types.SUCCESS
}
func (d *CliModuleData) ModuleShutdown(moduleNumber int) bool {
	return ZmShutdownCliServer(0, moduleNumber) == types.SUCCESS
}
func (d *CliModuleData) RequestStartup(moduleNumber int) bool {
	return true
}
func (d *CliModuleData) RequestShutdown(moduleNumber int) bool {
	return true
}

var CliServerModuleEntry = zend.MakeZendModuleEntry(&CliModuleData{}, ZmInfoCliServer)
