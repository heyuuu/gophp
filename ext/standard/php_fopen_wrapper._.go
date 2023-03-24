package standard

import (
	"sik/core"
)

// Source: <ext/standard/php_fopen_wrapper.c>

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
   | Authors: Rasmus Lerdorf <rasmus@php.net>                             |
   |          Jim Winstead <jimw@php.net>                                 |
   |          Hartmut Holzgraefe <hholzgra@php.net>                       |
   +----------------------------------------------------------------------+
*/

var PhpStreamOutputOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamOutputWrite, PhpStreamOutputRead, PhpStreamOutputClose, nil, "Output", nil, nil, nil, nil)

type PhpStreamInput = PhpStreamInputT

var PhpStreamInputOps core.PhpStreamOps = core.MakePhpStreamOps(PhpStreamInputWrite, PhpStreamInputRead, PhpStreamInputClose, PhpStreamInputFlush, "Input", PhpStreamInputSeek, nil, nil, nil)
var PhpStdioWops core.PhpStreamWrapperOps = core.MakePhpStreamWrapperOps(PhpStreamUrlWrapPhp, nil, nil, nil, nil, "PHP", nil, nil, nil, nil, nil)
var PhpStreamPhpWrapper core.PhpStreamWrapper = core.MakePhpStreamWrapper(&PhpStdioWops, nil, 0)
