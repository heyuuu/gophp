package zend

import "sik/zend/types"

// Source: <Zend/zend_execute_API.c>

/*
   +----------------------------------------------------------------------+
   | Zend Engine                                                          |
   +----------------------------------------------------------------------+
   | Copyright (c) Zend Technologies Ltd. (http://www.zend.com)           |
   +----------------------------------------------------------------------+
   | This source file is subject to version 2.00 of the Zend license,     |
   | that is bundled with this package in the file LICENSE, and is        |
   | available through the world-wide-web at the following url:           |
   | http://www.zend.com/license/2_00.txt.                                |
   | If you did not receive a copy of the Zend license and are unable to  |
   | obtain it through the world-wide-web, please send a note to          |
   | license@zend.com so we can mail you a copy immediately.              |
   +----------------------------------------------------------------------+
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/* true globals */

var EmptyFcallInfo types.ZendFcallInfo = types.MakeZendFcallInfo(0, types.Zval{}, nil, nil, nil, 0, 0)
var EmptyFcallInfoCache types.ZendFcallInfoCache = types.MakeZendFcallInfoCache(nil, nil, nil, nil)

/* This one doesn't exists on QNX */

const SIGPROF = 27
