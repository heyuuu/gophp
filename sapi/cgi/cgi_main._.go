package cgi

import (
	"github.com/heyuuu/gophp/core"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
)

var Act __struct__sigaction
var OldTerm __struct__sigaction
var OldQuit __struct__sigaction
var OldInt __struct__sigaction
var PhpPhpImportEnvironmentVariables func(array_ptr *types.Zval)

/* these globals used for forking children on unix systems */

var Children int = 0

/**
 * Set to non-zero if we are the parent process
 */

var Parent int = 1

/* Did parent received exit signals SIG_TERM/SIG_INT/SIG_QUIT */

var ExitSignal int = 0

/* Is Parent waiting for children to exit */

var ParentWaiting int = 0

/**
 * Process group
 */

var Pgroup pid_t

const PHP_MODE_STANDARD = 1
const PHP_MODE_HIGHLIGHT = 2
const PHP_MODE_LINT = 4
const PHP_MODE_STRIP = 5

var PhpOptarg *byte = nil
var PhpOptind int = 1
var OPTIONS []core.Opt = []core.Opt{
	core.MakeOpt('a', 0, "interactive"),
	core.MakeOpt('b', 1, "bindpath"),
	core.MakeOpt('C', 0, "no-chdir"),
	core.MakeOpt('c', 1, "php-ini"),
	core.MakeOpt('d', 1, "define"),
	core.MakeOpt('e', 0, "profile-info"),
	core.MakeOpt('f', 1, "file"),
	core.MakeOpt('h', 0, "help"),
	core.MakeOpt('i', 0, "info"),
	core.MakeOpt('l', 0, "syntax-check"),
	core.MakeOpt('m', 0, "modules"),
	core.MakeOpt('n', 0, "no-php-ini"),
	core.MakeOpt('q', 0, "no-header"),
	core.MakeOpt('s', 0, "syntax-highlight"),
	core.MakeOpt('s', 0, "syntax-highlighting"),
	core.MakeOpt('w', 0, "strip"),
	core.MakeOpt('?', 0, "usage"),
	core.MakeOpt('v', 0, "version"),
	core.MakeOpt('z', 1, "zend-extension"),
	core.MakeOpt('T', 1, "timing"),
	core.MakeOpt('-', 0, nil),
}

/* {{{ user_config_cache
 *
 * Key for each cache entry is dirname(PATH_TRANSLATED).
 *
 * NOTE: Each cache entry config_hash contains the combination from all user ini files found in
 *       the path starting from doc_root through to dirname(PATH_TRANSLATED).  There is no point
 *       storing per-file entries as it would not be possible to detect added / deleted entries
 *       between separate files.
 */

var php_cgi_globals php_cgi_globals_struct

const STDOUT_FILENO = 1
const SAPI_CGI_MAX_HEADER_LENGTH = 1024
const STDIN_FILENO = 0

var CgiFunctions = []types.FunctionEntry{
	DefZifApacheRequestHeaders,
	DefZifApacheResponseHeaders,
	types.MakeZendFunctionEntryEx("getallheaders", 0, ZifApacheRequestHeaders, []zend.ArgInfo{zend.MakeReturnArgInfo(-1)}),
}
var CgiModuleEntry = zend.MakeZendModuleEntry(
	"cgi-fcgi",
	CgiFunctions,
	ZmStartupCgi,
	ZmShutdownCgi,
	nil,
	nil,
	ZmInfoCgi,
	core.PHP_VERSION,

	0,
	nil,
	nil,
	nil,
)
