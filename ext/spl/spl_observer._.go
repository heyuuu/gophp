// <<generate>>

package spl

import (
	b "sik/builtin"
	"sik/zend"
)

// Source: <ext/spl/spl_observer.h>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   +----------------------------------------------------------------------+
*/

var spl_ce_SplObserver *zend.ZendClassEntry
var spl_ce_SplSubject *zend.ZendClassEntry
var spl_ce_SplObjectStorage *zend.ZendClassEntry
var spl_ce_MultipleIterator *zend.ZendClassEntry

// Source: <ext/spl/spl_observer.c>

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
   | Authors: Marcus Boerger <helly@php.net>                              |
   |          Etienne Kneuss <colder@php.net>                             |
   +----------------------------------------------------------------------+
*/

var zim_spl_SplObserver_update func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_attach func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_detach func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var zim_spl_SplSubject_notify func(execute_data *zend.ZendExecuteData, return_value *zend.Zval)
var arginfo_SplObserver_update []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"subject",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("SplSubject", 0),
		0,
		0,
	},
}
var spl_funcs_SplObserver []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"update",
		nil,
		arginfo_SplObserver_update,
		uint32_t(b.SizeOf("arginfo_SplObserver_update")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var arginfo_SplSubject_attach []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{
		"observer",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0),
		0,
		0,
	},
}
var arginfo_SplSubject_void []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}

/*ZEND_BEGIN_ARG_INFO_EX(arginfo_SplSubject_notify, 0, 0, 1)
    ZEND_ARG_OBJ_INFO(0, ignore, SplObserver, 1)
ZEND_END_ARG_INFO();*/

var spl_funcs_SplSubject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"attach",
		nil,
		arginfo_SplSubject_attach,
		uint32_t(b.SizeOf("arginfo_SplSubject_attach")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{
		"detach",
		nil,
		arginfo_SplSubject_attach,
		uint32_t(b.SizeOf("arginfo_SplSubject_attach")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{
		"notify",
		nil,
		arginfo_SplSubject_void,
		uint32_t(b.SizeOf("arginfo_SplSubject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		zend.ZEND_ACC_PUBLIC | zend.ZEND_ACC_ABSTRACT,
	},
	{nil, nil, nil, 0, 0},
}
var spl_handler_SplObjectStorage zend.ZendObjectHandlers

/* {{{ storage is an assoc array of [zend_object*]=>[zval *obj, zval *inf] */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* {{{ proto void SplObjectStorage::attach(object obj, mixed data = NULL)
Attaches an object to the storage if not yet contained */

/* {{{ proto void SplObjectStorage::detach(object obj)
Detaches an object from the storage */

/* {{{ proto string SplObjectStorage::getHash(object obj)
Returns the hash of an object */

/* {{{ proto mixed SplObjectStorage::offsetGet(object obj)
Returns associated information for a stored object */

/* {{{ proto bool SplObjectStorage::addAll(SplObjectStorage $os)
Add all elements contained in $os */

/* {{{ proto bool SplObjectStorage::removeAll(SplObjectStorage $os)
Remove all elements contained in $os */

/* {{{ proto bool SplObjectStorage::removeAllExcept(SplObjectStorage $os)
Remove elements not common to both this SplObjectStorage instance and $os */

/* }}} */

/* {{{ proto int SplObjectStorage::count()
Determine number of objects in storage */

/* {{{ proto void SplObjectStorage::rewind()
Rewind to first position */

/* {{{ proto bool SplObjectStorage::valid()
Returns whether current position is valid */

/* {{{ proto mixed SplObjectStorage::key()
Returns current key */

/* {{{ proto mixed SplObjectStorage::current()
Returns current element */

/* {{{ proto mixed SplObjectStorage::getInfo()
Returns associated information to current element */

/* {{{ proto mixed SplObjectStorage::setInfo(mixed $inf)
Sets associated information of current element to $inf */

/* {{{ proto void SplObjectStorage::next()
Moves position forward */

/* {{{ proto string SplObjectStorage::serialize()
Serializes storage */

/* {{{ proto void SplObjectStorage::unserialize(string serialized)
Unserializes storage */

/* {{{ proto auto SplObjectStorage::__serialize() */

/* {{{ proto void SplObjectStorage::__unserialize(array serialized) */

/* {{{ proto array SplObjectStorage::__debugInfo() */

/* }}} */

var arginfo_Object []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"object", 0, 0, 0},
}
var ArginfoAttach []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"object", 0, 0, 0}, {"data", 0, 0, 0}}
var arginfo_Serialized []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"serialized", 0, 0, 0},
}
var arginfo_setInfo []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"info", 0, 0, 0},
}
var arginfo_getHash []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
	{"object", 0, 0, 0},
}
var arginfo_offsetGet []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"object", 0, 0, 0}}
var ArginfoSplobjectVoid []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(-1)), 0, zend.ZEND_RETURN_VALUE, 0},
}
var spl_funcs_SplObjectStorage []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"attach",
		zim_spl_SplObjectStorage_attach,
		ArginfoAttach,
		uint32_t(b.SizeOf("arginfo_attach")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"detach",
		zim_spl_SplObjectStorage_detach,
		arginfo_Object,
		uint32_t(b.SizeOf("arginfo_Object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"contains",
		zim_spl_SplObjectStorage_contains,
		arginfo_Object,
		uint32_t(b.SizeOf("arginfo_Object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"addAll",
		zim_spl_SplObjectStorage_addAll,
		arginfo_Object,
		uint32_t(b.SizeOf("arginfo_Object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"removeAll",
		zim_spl_SplObjectStorage_removeAll,
		arginfo_Object,
		uint32_t(b.SizeOf("arginfo_Object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"removeAllExcept",
		zim_spl_SplObjectStorage_removeAllExcept,
		arginfo_Object,
		uint32_t(b.SizeOf("arginfo_Object")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getInfo",
		zim_spl_SplObjectStorage_getInfo,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setInfo",
		zim_spl_SplObjectStorage_setInfo,
		arginfo_setInfo,
		uint32_t(b.SizeOf("arginfo_setInfo")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getHash",
		zim_spl_SplObjectStorage_getHash,
		arginfo_getHash,
		uint32_t(b.SizeOf("arginfo_getHash")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__debugInfo",
		zim_spl_SplObjectStorage___debugInfo,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"count",
		zim_spl_SplObjectStorage_count,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		zim_spl_SplObjectStorage_rewind,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"valid",
		zim_spl_SplObjectStorage_valid,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		zim_spl_SplObjectStorage_key,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		zim_spl_SplObjectStorage_current,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		zim_spl_SplObjectStorage_next,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"unserialize",
		zim_spl_SplObjectStorage_unserialize,
		arginfo_Serialized,
		uint32_t(b.SizeOf("arginfo_Serialized")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"serialize",
		zim_spl_SplObjectStorage_serialize,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__unserialize",
		zim_spl_SplObjectStorage___unserialize,
		arginfo_Serialized,
		uint32_t(b.SizeOf("arginfo_Serialized")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__serialize",
		zim_spl_SplObjectStorage___serialize,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetExists",
		zim_spl_SplObjectStorage_contains,
		arginfo_offsetGet,
		uint32_t(b.SizeOf("arginfo_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetSet",
		zim_spl_SplObjectStorage_attach,
		ArginfoAttach,
		uint32_t(b.SizeOf("arginfo_attach")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetUnset",
		zim_spl_SplObjectStorage_detach,
		arginfo_offsetGet,
		uint32_t(b.SizeOf("arginfo_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"offsetGet",
		zim_spl_SplObjectStorage_offsetGet,
		arginfo_offsetGet,
		uint32_t(b.SizeOf("arginfo_offsetGet")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

type MultipleIteratorFlags = int

const (
	MIT_NEED_ANY     = 0
	MIT_NEED_ALL     = 1
	MIT_KEYS_NUMERIC = 0
	MIT_KEYS_ASSOC   = 2
)
const SPL_MULTIPLE_ITERATOR_GET_ALL_CURRENT = 1
const SPL_MULTIPLE_ITERATOR_GET_ALL_KEY = 2

/* {{{ proto MultipleIterator::__construct([int flags = MIT_NEED_ALL|MIT_KEYS_NUMERIC])
   Iterator that iterates over several iterators one after the other */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

/* }}} */

var arginfo_MultipleIterator_attachIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
	{"infos", 0, 0, 0},
}
var arginfo_MultipleIterator_detachIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
}
var arginfo_MultipleIterator_containsIterator []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{
	{(*byte)(zend_uintptr_t(1)), 0, 0, 0},
	{
		"iterator",
		zend.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0),
		0,
		0,
	},
}
var arginfo_MultipleIterator_setflags []zend.ZendInternalArgInfo = []zend.ZendInternalArgInfo{{(*byte)(zend_uintptr_t(1)), 0, 0, 0}, {"flags", 0, 0, 0}}
var spl_funcs_MultipleIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	{
		"__construct",
		zim_spl_MultipleIterator___construct,
		arginfo_MultipleIterator_setflags,
		uint32_t(b.SizeOf("arginfo_MultipleIterator_setflags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"getFlags",
		zim_spl_MultipleIterator_getFlags,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"setFlags",
		zim_spl_MultipleIterator_setFlags,
		arginfo_MultipleIterator_setflags,
		uint32_t(b.SizeOf("arginfo_MultipleIterator_setflags")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"attachIterator",
		zim_spl_MultipleIterator_attachIterator,
		arginfo_MultipleIterator_attachIterator,
		uint32_t(b.SizeOf("arginfo_MultipleIterator_attachIterator")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"detachIterator",
		zim_spl_SplObjectStorage_detach,
		arginfo_MultipleIterator_detachIterator,
		uint32_t(b.SizeOf("arginfo_MultipleIterator_detachIterator")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"containsIterator",
		zim_spl_SplObjectStorage_contains,
		arginfo_MultipleIterator_containsIterator,
		uint32_t(b.SizeOf("arginfo_MultipleIterator_containsIterator")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"countIterators",
		zim_spl_SplObjectStorage_count,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"__debugInfo",
		zim_spl_SplObjectStorage___debugInfo,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"rewind",
		zim_spl_MultipleIterator_rewind,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"valid",
		zim_spl_MultipleIterator_valid,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"key",
		zim_spl_MultipleIterator_key,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"current",
		zim_spl_MultipleIterator_current,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{
		"next",
		zim_spl_MultipleIterator_next,
		ArginfoSplobjectVoid,
		uint32_t(b.SizeOf("arginfo_splobject_void")/b.SizeOf("struct _zend_internal_arg_info") - 1),
		0,
	},
	{nil, nil, nil, 0, 0},
}

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */

/* }}} */
