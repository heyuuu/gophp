// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
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

// #define INCOMPLETE_CLASS_MSG       "The script tried to execute a method or " "access a property of an incomplete object. " "Please ensure that the class definition \"%s\" of the object " "you are trying to operate on was loaded _before_ " "unserialize() gets called or provide an autoloader " "to load the class definition"

var PhpIncompleteObjectHandlers zend.ZendObjectHandlers

/* {{{ incomplete_class_message
 */

func IncompleteClassMessage(object *zend.Zval, error_type int) {
	var class_name *zend.ZendString
	class_name = PhpLookupClassName(object)
	if class_name != nil {
		core.PhpErrorDocref(nil, error_type, "The script tried to execute a method or "+"access a property of an incomplete object. "+"Please ensure that the class definition \"%s\" of the object "+"you are trying to operate on was loaded _before_ "+"unserialize() gets called or provide an autoloader "+"to load the class definition", class_name.val)
		zend.ZendStringReleaseEx(class_name, 0)
	} else {
		core.PhpErrorDocref(nil, error_type, "The script tried to execute a method or "+"access a property of an incomplete object. "+"Please ensure that the class definition \"%s\" of the object "+"you are trying to operate on was loaded _before_ "+"unserialize() gets called or provide an autoloader "+"to load the class definition", "unknown")
	}
}

/* }}} */

func IncompleteClassGetProperty(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any, rv *zend.Zval) *zend.Zval {
	IncompleteClassMessage(object, 1<<3)
	if type_ == 1 || type_ == 2 {
		rv.u1.type_info = 15
		return rv
	} else {
		return &zend.EG.uninitialized_zval
	}
}

/* }}} */

func IncompleteClassWriteProperty(object *zend.Zval, member *zend.Zval, value *zend.Zval, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, 1<<3)
	return value
}

/* }}} */

func IncompleteClassGetPropertyPtrPtr(object *zend.Zval, member *zend.Zval, type_ int, cache_slot *any) *zend.Zval {
	IncompleteClassMessage(object, 1<<3)
	return &zend.EG.error_zval
}

/* }}} */

func IncompleteClassUnsetProperty(object *zend.Zval, member *zend.Zval, cache_slot *any) {
	IncompleteClassMessage(object, 1<<3)
}

/* }}} */

func IncompleteClassHasProperty(object *zend.Zval, member *zend.Zval, check_empty int, cache_slot *any) int {
	IncompleteClassMessage(object, 1<<3)
	return 0
}

/* }}} */

func IncompleteClassGetMethod(object **zend.ZendObject, method *zend.ZendString, key *zend.Zval) *zend.ZendFunction {
	var zobject zend.Zval
	var __z *zend.Zval = &zobject
	__z.value.obj = *object
	__z.u1.type_info = 8 | 1<<0<<8 | 1<<1<<8
	IncompleteClassMessage(&zobject, 1<<0)
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
	memset(&incomplete_class, 0, g.SizeOf("zend_class_entry"))
	incomplete_class.name = zend.ZendStringInitInterned("__PHP_Incomplete_Class", g.SizeOf("INCOMPLETE_CLASS")-1, 1)
	incomplete_class.info.internal.builtin_functions = nil
	incomplete_class.create_object = PhpCreateIncompleteObject
	memcpy(&PhpIncompleteObjectHandlers, &zend.StdObjectHandlers, g.SizeOf("zend_object_handlers"))
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
	object_properties = object.value.obj.handlers.get_properties(&(*object))
	if g.Assign(&val, zend.ZendHashStrFind(object_properties, "__PHP_Incomplete_Class_Name", g.SizeOf("MAGIC_MEMBER")-1)) != nil && val.u1.v.type_ == 6 {
		return zend.ZendStringCopy(val.value.str)
	}
	return nil
}

/* }}} */

func PhpStoreClassName(object *zend.Zval, name *byte, len_ int) {
	var val zend.Zval
	var __z *zend.Zval = &val
	var __s *zend.ZendString = zend.ZendStringInit(name, len_, 0)
	__z.value.str = __s
	__z.u1.type_info = 6 | 1<<0<<8
	zend.ZendHashStrUpdate(object.value.obj.handlers.get_properties(&(*object)), "__PHP_Incomplete_Class_Name", g.SizeOf("MAGIC_MEMBER")-1, &val)
}

/* }}} */
