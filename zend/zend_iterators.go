// <<generate>>

package zend

import (
	r "sik/runtime"
	g "sik/runtime/grammar"
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
	memset(&ZendIteratorClassEntry, 0, g.SizeOf("zend_class_entry"))
	ZendIteratorClassEntry.SetName(ZendStringInitInterned("__iterator_wrapper", g.SizeOf("\"__iterator_wrapper\"")-1, 1))
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
	if ZendGcDelref(&(&iter.std).GetGc()) > 0 {
		return
	}
	ZendObjectsStoreDel(&iter.std)
}
func ZendIteratorUnwrap(array_ptr *Zval) *ZendObjectIterator {
	r.Assert(array_ptr.GetType() == 8)
	if array_ptr.GetValue().GetObj().GetHandlers() == &IteratorObjectHandlers {
		return (*ZendObjectIterator)(array_ptr.GetValue().GetObj())
	}
	return nil
}
