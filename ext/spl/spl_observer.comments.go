// <<generate>>

package spl

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

/*ZEND_BEGIN_ARG_INFO_EX(arginfo_SplSubject_notify, 0, 0, 1)
    ZEND_ARG_OBJ_INFO(0, ignore, SplObserver, 1)
ZEND_END_ARG_INFO();*/

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

/* {{{ PHP_MINIT_FUNCTION(spl_observer) */

/* }}} */
