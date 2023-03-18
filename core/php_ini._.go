// <<generate>>

package core

import (
	"sik/zend"
	"sik/zend/types"
)

const PHP_INI_USER = zend.ZEND_INI_USER
const PHP_INI_PERDIR = zend.ZEND_INI_PERDIR
const PHP_INI_SYSTEM = zend.ZEND_INI_SYSTEM
const PHP_INI_ALL = zend.ZEND_INI_ALL
const PHP_INI_STAGE_STARTUP = zend.ZEND_INI_STAGE_STARTUP
const PHP_INI_STAGE_SHUTDOWN = zend.ZEND_INI_STAGE_SHUTDOWN
const PHP_INI_STAGE_ACTIVATE = zend.ZEND_INI_STAGE_ACTIVATE
const PHP_INI_STAGE_DEACTIVATE = zend.ZEND_INI_STAGE_DEACTIVATE
const PHP_INI_STAGE_RUNTIME = zend.ZEND_INI_STAGE_RUNTIME
const PHP_INI_STAGE_HTACCESS = zend.ZEND_INI_STAGE_HTACCESS
const PhpIniColorDisplayerCb = zend.ZendIniColorDisplayerCb

var IsSpecialSection int = 0
var ActiveIniHash *types.HashTable
var HasPerDirConfig int = 0
var HasPerHostConfig int = 0
var PhpIniOpenedPath *byte = nil
var ExtensionLists PhpExtensionLists
var PhpIniScannedPath *byte = nil
var PhpIniScannedFiles *byte = nil

const PHP_EXTENSION_TOKEN = "extension"
const ZEND_EXTENSION_TOKEN = "zend_extension"
