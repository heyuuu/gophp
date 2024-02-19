package php

import "os"

/* var status for backpatching */
// @see: https://wiki.php.net/internals/engine/objects?s[]=bp_var_is
const BP_VAR_R = 0
const BP_VAR_W = 1
const BP_VAR_RW = 2
const BP_VAR_IS = 3
const BP_VAR_FUNC_ARG = 4
const BP_VAR_UNSET = 5

// build-defs

const CONFIGURE_COMMAND = " './configure'  '--with-iconv=/opt/homebrew/Cellar/libiconv/1.16'"
const PHP_PROG_SENDMAIL = "/usr/sbin/sendmail"
const PEAR_INSTALLDIR = ""
const PHP_INCLUDE_PATH = ".:"
const PHP_EXTENSION_DIR = "/usr/local/lib/php/extensions/no-debug-non-zts-20190902"
const PHP_PREFIX = "/usr/local"
const PHP_BINDIR = "/usr/local/bin"
const PHP_MANDIR = "/usr/local/php/man"
const PHP_LIBDIR = "/usr/local/lib/php"
const PHP_DATADIR = "/usr/local/share/php"
const PHP_SYSCONFDIR = "/usr/local/etc"
const PHP_LOCALSTATEDIR = "/usr/local/var"
const PHP_CONFIG_FILE_PATH = "/usr/local/lib"
const PHP_CONFIG_FILE_SCAN_DIR = ""
const PHP_SHLIB_SUFFIX = "so"
const PHP_SHLIB_EXT_PREFIX = ""

// config

const PHP_OS = "Darwin"
const PHP_UNAME = "Darwin bogon 21.6.0 Darwin Kernel Version 21.6.0: Wed Aug 10 14:28:23 PDT 2022; root:xnu-8020.141.5~2/RELEASE_ARM64_T6000 arm64"

// version

const PHP_MAJOR_VERSION = 7
const PHP_MINOR_VERSION = 4
const PHP_RELEASE_VERSION = 33
const PHP_EXTRA_VERSION = ""
const PHP_VERSION = "7.4.33"
const PHP_VERSION_ID = 70433

// php

const PHP_API_VERSION = 20190902
const PHP_OS_FAMILY = "Unknown"
const PHP_DIR_SEPARATOR = os.PathSeparator
const PHP_EOL = "\n"

const EXEC_INPUT_BUF = 4096
