// <<generate>>

package zend

// Source: <Zend/zend_modules.h>

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
   +----------------------------------------------------------------------+
*/

const ZEND_MODULE_INFO_FUNC_ARGS = zend_module_entry * zend_module
const ZEND_MODULE_INFO_FUNC_ARGS_PASSTHRU ZEND_MODULE_INFO_FUNC_ARGS = zend_module
const ZEND_MODULE_API_NO = 20190902
const USING_ZTS = 0
const NO_VERSION_YET = nil
const MODULE_PERSISTENT = 1
const MODULE_TEMPORARY = 2

const MODULE_DEP_REQUIRED = 1
const MODULE_DEP_CONFLICTS = 2
const MODULE_DEP_OPTIONAL = 3

var ModuleRegistryRequestStartup func(module *ZendModuleEntry) int
var ModuleRegistryUnloadTemp func(module *ZendModuleEntry) int
