// <<generate>>

package zend

import (
	b "sik/builtin"
)

// Source: <Zend/zend_iterators.h>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   |         Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

/* given a zval, returns stuff that can be used to iterate it. */

/* given an iterator, wrap it up as a zval for use by the engine opcodes */

// Source: <Zend/zend_iterators.c>

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
   | Author: Wez Furlong <wez@thebrainroom.com>                           |
   |         Marcus Boerger <helly@php.net>                               |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_API.h"

var ZendIteratorClassEntry ZendClassEntry
var IteratorObjectHandlers ZendObjectHandlers = ZendObjectHandlers{0, IterWrapperFree, IterWrapperDtor, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, IterWrapperGetGc, nil, nil}

func ZendRegisterIteratorWrapper() {
	memset(&ZendIteratorClassEntry, 0, b.SizeOf("zend_class_entry"))
	ZendIteratorClassEntry.SetName(ZendStringInitInterned("__iterator_wrapper", b.SizeOf("\"__iterator_wrapper\"")-1, 1))
	ZendIteratorClassEntry.SetBuiltinFunctions(nil)
}
func IterWrapperFree(object *ZendObject) {
	var iter *ZendObjectIterator = (*ZendObjectIterator)(object)
	iter.GetFuncs().GetDtor()(iter)
}
func IterWrapperDtor(object *ZendObject) {}
func IterWrapperGetGc(object *Zval, table **Zval, n *int) *HashTable {
	/* TODO: We need a get_gc iterator handler */

	*table = nil
	*n = 0
	return nil
}
func ZendIteratorInit(iter *ZendObjectIterator) {
	ZendObjectStdInit(&iter.std, &ZendIteratorClassEntry)
	iter.GetStd().SetHandlers(&IteratorObjectHandlers)
}
func ZendIteratorDtor(iter *ZendObjectIterator) {
	if GC_DELREF(&iter.std) > 0 {
		return
	}
	ZendObjectsStoreDel(&iter.std)
}
func ZendIteratorUnwrap(array_ptr *Zval) *ZendObjectIterator {
	ZEND_ASSERT(Z_TYPE_P(array_ptr) == IS_OBJECT)
	if Z_OBJ_HT_P(array_ptr) == &IteratorObjectHandlers {
		return (*ZendObjectIterator)(Z_OBJ_P(array_ptr))
	}
	return nil
}
