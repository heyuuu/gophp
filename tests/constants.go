package tests

import "strconv"

const E_ALL = 32767

var baseIniOverwrites = []string{
	"output_handler=",
	"open_basedir=",
	"disable_functions=",
	"output_buffering=Off",
	"error_reporting=" + strconv.Itoa(E_ALL),
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

const noFileCache = "-d opcache.file_cache= -d opcache.file_cache_only=0"

const helpText = `
Synopsis:
    php run-tests.php [options] [files] [directories]

Options:
    -l <file>   Read the testfiles to be executed from <file>. After the test
                has finished all failed tests are written to the same <file>.
                If the list is empty and no further test is specified then
                all tests are executed (same as: -r <file> -w <file>).

    -r <file>   Read the testfiles to be executed from <file>.

    -w <file>   Write a list of all failed tests to <file>.

    -a <file>   Same as -w but append rather then truncating <file>.

    -W <file>   Write a list of all tests and their result status to <file>.

    -c <file>   Look for php.ini in directory <file> or use <file> as ini.

    -n          Pass -n option to the php binary (Do not use a php.ini).

    -d foo=bar  Pass -d option to the php binary (Define INI entry foo
                with value 'bar').

    -g          Comma separated list of groups to show during test run
                (possible values: PASS, FAIL, XFAIL, XLEAK, SKIP, BORK, WARN, LEAK, REDIRECT).

    -p <php>    Specify PHP executable to run.

    -P          Use PHP_BINARY as PHP executable to run (default).

    -q          Quiet, no user interaction (same as environment NO_INTERACTION).

    -s <file>   Write output to <file>.

    -x          Sets 'SKIP_SLOW_TESTS' environmental variable.

    --offline   Sets 'SKIP_ONLINE_TESTS' environmental variable.

    --verbose
    -v          Verbose mode.

    --help
    -h          This Help.

    --temp-source <sdir>  --temp-target <tdir> [--temp-urlbase <url>]
                Write temporary files to <tdir> by replacing <sdir> from the
                filenames to generate with <tdir>. If --html is being used and
                <url> given then the generated links are relative and prefixed
                with the given url. In general you want to make <sdir> the path
                to your source files and <tdir> some patch in your web page
                hierarchy with <url> pointing to <tdir>.

    --keep-[all|php|skip|clean]
                Do not delete 'all' files, 'php' test file, 'skip' or 'clean'
                file.

    --set-timeout [n]
                Set timeout for individual tests, where [n] is the number of
                seconds. The default value is 60 seconds, or 300 seconds when
                testing for memory leaks.

    --show-[all|php|skip|clean|exp|diff|out|mem]
                Show 'all' files, 'php' test file, 'skip' or 'clean' file. You
                can also use this to show the output 'out', the expected result
                'exp', the difference between them 'diff' or the valgrind log
                'mem'. The result types get written independent of the log format,
                however 'diff' only exists when a test fails.

    --show-slow [n]
                Show all tests that took longer than [n] milliseconds to run.

    --no-clean  Do not execute clean section if any.
`
