// <<generate>>

package zend

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
