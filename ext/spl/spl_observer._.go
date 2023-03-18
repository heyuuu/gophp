// <<generate>>

package spl

import (
	"sik/zend"
	"sik/zend/types"
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

var zim_spl_SplObserver_update func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var zim_spl_SplSubject_attach func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var zim_spl_SplSubject_detach func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var zim_spl_SplSubject_notify func(executeData *zend.ZendExecuteData, return_value *types.Zval)
var arginfo_SplObserver_update []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("subject", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("SplSubject", 0))),
}
var spl_funcs_SplObserver []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("update", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, arginfo_SplObserver_update),
}
var arginfo_SplSubject_attach []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("observer", zend.ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("SplObserver", 0))),
}
var arginfo_SplSubject_void []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}

/*ZEND_BEGIN_ARG_INFO_EX(arginfo_SplSubject_notify, 0, 0, 1)
    ZEND_ARG_OBJ_INFO(0, ignore, SplObserver, 1)
ZEND_END_ARG_INFO();*/

var spl_funcs_SplSubject []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("attach", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, arginfo_SplSubject_attach),
	zend.MakeZendFunctionEntryEx("detach", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, arginfo_SplSubject_attach),
	zend.MakeZendFunctionEntryEx("notify", zend.ZEND_ACC_PUBLIC|zend.ZEND_ACC_ABSTRACT, nil, arginfo_SplSubject_void),
}
var spl_handler_SplObjectStorage zend.ZendObjectHandlers

/* {{{ storage is an assoc array of [zend_object*]=>[zval *obj, zval *inf] */

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

var arginfo_Object []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("object"),
}
var ArginfoAttach []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("object"),
	zend.MakeArgInfo("data"),
}
var arginfo_Serialized []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("serialized"),
}
var arginfo_setInfo []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("info"),
}
var arginfo_getHash []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
	zend.MakeArgInfo("object"),
}
var arginfo_offsetGet []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("object"),
}
var ArginfoSplobjectVoid []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(-1),
}
var spl_funcs_SplObjectStorage []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("attach", 0, zim_spl_SplObjectStorage_attach, ArginfoAttach),
	zend.MakeZendFunctionEntryEx("detach", 0, zim_spl_SplObjectStorage_detach, arginfo_Object),
	zend.MakeZendFunctionEntryEx("contains", 0, zim_spl_SplObjectStorage_contains, arginfo_Object),
	zend.MakeZendFunctionEntryEx("addAll", 0, zim_spl_SplObjectStorage_addAll, arginfo_Object),
	zend.MakeZendFunctionEntryEx("removeAll", 0, zim_spl_SplObjectStorage_removeAll, arginfo_Object),
	zend.MakeZendFunctionEntryEx("removeAllExcept", 0, zim_spl_SplObjectStorage_removeAllExcept, arginfo_Object),
	zend.MakeZendFunctionEntryEx("getInfo", 0, zim_spl_SplObjectStorage_getInfo, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("setInfo", 0, zim_spl_SplObjectStorage_setInfo, arginfo_setInfo),
	zend.MakeZendFunctionEntryEx("getHash", 0, zim_spl_SplObjectStorage_getHash, arginfo_getHash),
	zend.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("count", 0, zim_spl_SplObjectStorage_count, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("rewind", 0, zim_spl_SplObjectStorage_rewind, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("valid", 0, zim_spl_SplObjectStorage_valid, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("key", 0, zim_spl_SplObjectStorage_key, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("current", 0, zim_spl_SplObjectStorage_current, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("next", 0, zim_spl_SplObjectStorage_next, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("unserialize", 0, zim_spl_SplObjectStorage_unserialize, arginfo_Serialized),
	zend.MakeZendFunctionEntryEx("serialize", 0, zim_spl_SplObjectStorage_serialize, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("__unserialize", 0, zim_spl_SplObjectStorage___unserialize, arginfo_Serialized),
	zend.MakeZendFunctionEntryEx("__serialize", 0, zim_spl_SplObjectStorage___serialize, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("offsetExists", 0, zim_spl_SplObjectStorage_contains, arginfo_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetSet", 0, zim_spl_SplObjectStorage_attach, ArginfoAttach),
	zend.MakeZendFunctionEntryEx("offsetUnset", 0, zim_spl_SplObjectStorage_detach, arginfo_offsetGet),
	zend.MakeZendFunctionEntryEx("offsetGet", 0, zim_spl_SplObjectStorage_offsetGet, arginfo_offsetGet),
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

var arginfo_MultipleIterator_attachIterator []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
	zend.MakeArgInfo("infos"),
}
var arginfo_MultipleIterator_detachIterator []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
}
var arginfo_MultipleIterator_containsIterator []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("iterator", ArgInfoType(types.ZEND_TYPE_ENCODE_CLASS_CONST("Iterator", 0))),
}
var arginfo_MultipleIterator_setflags []zend.ArgInfo = []zend.ArgInfo{
	zend.MakeReturnArgInfo(1),
	zend.MakeArgInfo("flags"),
}
var spl_funcs_MultipleIterator []zend.ZendFunctionEntry = []zend.ZendFunctionEntry{
	zend.MakeZendFunctionEntryEx("__construct", 0, zim_spl_MultipleIterator___construct, arginfo_MultipleIterator_setflags),
	zend.MakeZendFunctionEntryEx("getFlags", 0, zim_spl_MultipleIterator_getFlags, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("setFlags", 0, zim_spl_MultipleIterator_setFlags, arginfo_MultipleIterator_setflags),
	zend.MakeZendFunctionEntryEx("attachIterator", 0, zim_spl_MultipleIterator_attachIterator, arginfo_MultipleIterator_attachIterator),
	zend.MakeZendFunctionEntryEx("detachIterator", 0, zim_spl_SplObjectStorage_detach, arginfo_MultipleIterator_detachIterator),
	zend.MakeZendFunctionEntryEx("containsIterator", 0, zim_spl_SplObjectStorage_contains, arginfo_MultipleIterator_containsIterator),
	zend.MakeZendFunctionEntryEx("countIterators", 0, zim_spl_SplObjectStorage_count, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("__debugInfo", 0, zim_spl_SplObjectStorage___debugInfo, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("rewind", 0, zim_spl_MultipleIterator_rewind, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("valid", 0, zim_spl_MultipleIterator_valid, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("key", 0, zim_spl_MultipleIterator_key, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("current", 0, zim_spl_MultipleIterator_current, ArginfoSplobjectVoid),
	zend.MakeZendFunctionEntryEx("next", 0, zim_spl_MultipleIterator_next, ArginfoSplobjectVoid),
}

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */
