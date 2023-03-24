package streams

import (
	"sik/core"
)

// Source: <main/streams/plain_wrapper.c>

/*
   +----------------------------------------------------------------------+
   | PHP Version 7                                                        |
   +----------------------------------------------------------------------+
   | Copyright (c) The PHP Group                                          |
   +----------------------------------------------------------------------+
   | This source file is subject to version 3.01 of the PHP license,      |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.php.net/license/3_01.txt                                  |
   | If you did not receive a copy of the PHP license and are unable to   |
   | obtain it through the world-wide-web, please send a note to          |
   | license@php.net so we can mail you a copy immediately.               |
   +----------------------------------------------------------------------+
   | Authors: Wez Furlong <wez@thebrainroom.com>                          |
   +----------------------------------------------------------------------+
*/

var PhpGetUidByName func(name *byte, uid *uid_t) int
var PhpGetGidByName func(name *byte, gid *gid_t) int

/* parse standard "fopen" modes into open() flags */

/* {{{ ------- STDIO stream implementation -------*/

/* This should be "const", but phpdbg overwrite it */

var PhpStreamStdioOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStdiopWrite, PhpStdiopRead, PhpStdiopClose, PhpStdiopFlush, "STDIO", PhpStdiopSeek, PhpStdiopCast, PhpStdiopStat, PhpStdiopSetOption)
var PhpPlainFilesDirstreamOps core.PhpStreamOps = core.MakePhpStreamOps(nil, PhpPlainFilesDirstreamRead, PhpPlainFilesDirstreamClose, nil, "dir", PhpPlainFilesDirstreamRewind, nil, nil, nil)
var PhpPlainFilesWrapperOps core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpPlainFilesStreamOpener, nil, nil, PhpPlainFilesUrlStater, PhpPlainFilesDirOpener, "plainfile", PhpPlainFilesUnlink, PhpPlainFilesRename, PhpPlainFilesMkdir, PhpPlainFilesRmdir, PhpPlainFilesMetadata)

/* TODO: We have to make php_plain_files_wrapper writable to support SWOOLE */

var PhpPlainFilesWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpPlainFilesWrapperOps, nil, 0)

/* {{{ php_stream_fopen_with_path */
