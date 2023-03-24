package standard

import (
	"sik/zend"
	"sik/zend/types"
)

// Source: <ext/standard/dir.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

var DirGlobals PhpDirGlobals
var DirClassEntryPtr *types.ClassEntry

/* {{{ arginfo */

var PhpDirClassFunctions []types.ZendFunctionEntry = []types.ZendFunctionEntry{
	types.MakeZendFunctionEntryEx("close", 0, ZifClosedir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("rewind", 0, ZifRewinddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
	types.MakeZendFunctionEntryEx("read", 0, PhpIfReaddir, []zend.ArgInfo{zend.MakeReturnArgInfo(1),
		zend.MakeArgName("directory"),
		zend.MakeArgName("context"),
	}),
}
