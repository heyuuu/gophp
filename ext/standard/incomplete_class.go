// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

// Source: <ext/standard/incomplete_class.c>

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
   | Author:  Sascha Schumann <sascha@schumann.cx>                        |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include "basic_functions.h"

// # include "php_incomplete_class.h"

const INCOMPLETE_CLASS_MSG string = "The script tried to execute a method or " + "access a property of an incomplete object. " + "Please ensure that the class definition \"%s\" of the object " + "you are trying to operate on was loaded _before_ " + "unserialize() gets called or provide an autoloader " + "to load the class definition"

var PhpIncompleteObjectHandlers zend.ZendObjectHandlers

/* {{{ incomplete_class_message
 */

func IncompleteClassMessage(object *zend.Zval, error_type int) {
	var class_name *zend.ZendString
	class_name = PhpLookupClassName(object)
	if class_name != nil {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, zend.ZSTR_VAL(class_name))
		zend.ZendStringReleaseEx(class_name, 0)
	} else {
		core.PhpErrorDocref(nil, error_type, INCOMPLETE_CLASS_MSG, "unknown")
	}
}

/* }}} */

func IncompleteClassGetProperty(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any, rv *zend.Zval) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	if type_ == zend.BP_VAR_W || type_ == zend.BP_VAR_RW {
		zend.ZVAL_ERROR(rv)
		return rv
	} else {
		return &(zend.ExecutorGlobals.uninitialized_zval)
	}
}

/* }}} */

func IncompleteClassWriteProperty(object *zend.Zval, member *zend.Zval, value *zend.Zval, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return value
}

/* }}} */

func IncompleteClassGetPropertyPtrPtr(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return &(zend.ExecutorGlobals.error_zval)
}

/* }}} */

func IncompleteClassUnsetProperty(object *zend.Zval, member *zend.Zval, cache_slot *any) {
	IncompleteClassMessage(object, zend.E_NOTICE)
}

/* }}} */

func IncompleteClassHasProperty(object *zend.Zval, member *zend.Zval, check_empty int, cache_slot *any) int {
	IncompleteClassMessage(object, zend.E_NOTICE)
	return 0
}

/* }}} */

func IncompleteClassGetMethod(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var zobject zend.Zval
	zend.ZVAL_OBJ(&zobject, *object)
	IncompleteClassMessage(&zobject, zend.E_ERROR)
	return nil
}

/* }}} */

func PhpCreateIncompleteObject(class_type *zend.ZendClassEntry) *zend.ZendObject {
	var object *zend.ZendObject
	object = zend.ZendObjectsNew(class_type)
	object.handlers = &PhpIncompleteObjectHandlers
	zend.ObjectPropertiesInit(object, class_type)
	return object
}
func PhpCreateIncompleteClass() *zend.ZendClassEntry {
	var incomplete_class zend.ZendClassEntry
	memset(&incomplete_class, 0, b.SizeOf("zend_class_entry"))
	incomplete_class.name = zend.ZendStringInitInterned(INCOMPLETE_CLASS, b.SizeOf("INCOMPLETE_CLASS")-1, 1)
	incomplete_class.info.internal.builtin_functions = nil
	incomplete_class.create_object = PhpCreateIncompleteObject
	memcpy(&PhpIncompleteObjectHandlers, &zend.StdObjectHandlers, b.SizeOf("zend_object_handlers"))
	PhpIncompleteObjectHandlers.read_property = IncompleteClassGetProperty
	PhpIncompleteObjectHandlers.has_property = IncompleteClassHasProperty
	PhpIncompleteObjectHandlers.unset_property = IncompleteClassUnsetProperty
	PhpIncompleteObjectHandlers.write_property = IncompleteClassWriteProperty
	PhpIncompleteObjectHandlers.get_property_ptr_ptr = IncompleteClassGetPropertyPtrPtr
	PhpIncompleteObjectHandlers.get_method = IncompleteClassGetMethod
	return zend.ZendRegisterInternalClass(&incomplete_class)
}

/* }}} */

func PhpLookupClassName(object *zend.Zval) *zend.ZendString {
	var val *zend.Zval
	var object_properties *zend.HashTable
	object_properties = zend.Z_OBJPROP_P(object)
	if b.Assign(&val, zend.ZendHashStrFind(object_properties, MAGIC_MEMBER, b.SizeOf("MAGIC_MEMBER")-1)) != nil && zend.Z_TYPE_P(val) == zend.IS_STRING {
		return zend.ZendStringCopy(zend.Z_STR_P(val))
	}
	return nil
}

/* }}} */

func PhpStoreClassName(object *zend.Zval, name *byte, len_ int) {
	var val zend.Zval
	zend.ZVAL_STRINGL(&val, name, len_)
	zend.ZendHashStrUpdate(zend.Z_OBJPROP_P(object), MAGIC_MEMBER, b.SizeOf("MAGIC_MEMBER")-1, &val)
}

/* }}} */
