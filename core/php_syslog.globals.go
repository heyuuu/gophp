// <<generate>>

package core

const PHP_SYSLOG_FILTER_ALL = 0
const PHP_SYSLOG_FILTER_NO_CTRL = 1
const PHP_SYSLOG_FILTER_ASCII = 2
const PHP_SYSLOG_FILTER_RAW = 3

var PhpOpenlog func(*byte, int, int)
