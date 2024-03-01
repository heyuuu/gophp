package tests

import (
	"github.com/heyuuu/gophp/php/perr"
	"strconv"
)

var baseIniOverwrites = []string{
	"output_handler=",
	"open_basedir=",
	"disable_functions=",
	"output_buffering=Off",
	"error_reporting=" + strconv.Itoa(int(perr.E_ALL)),
	"display_errors=1",
	"display_startup_errors=1",
	"log_errors=0",
	"html_errors=0",
	"track_errors=0",
	"report_memleaks=1",
	"report_zend_debug=0",
	"docref_root=",
	"docref_ext=.html",
	"error_prepend_string=",
	"error_append_string=",
	"auto_prepend_file=",
	"auto_append_file=",
	"ignore_repeated_errors=0",
	"precision=14",
	"memory_limit=128M",
	"log_errors_max_len=0",
	"opcache.fast_shutdown=0",
	"opcache.file_update_protection=0",
	"opcache.revalidate_freq=0",
	"zend.assertions=1",
	"zend.exception_ignore_args=0",
}

var noFileCacheArgs = []commandArg{
	arg("-d"), arg("opcache.file_cache="),
	arg("-d"), arg("opcache.file_cache_only=0"),
}
