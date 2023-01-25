// <<generate>>

package zend

// Source: <Zend/zend_type_info.h>

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
   | Authors: Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

// #define ZEND_TYPE_INFO_H

// # include "zend_types.h"

const MAY_BE_UNDEF = 1 << IS_UNDEF
const MAY_BE_NULL = 1 << IS_NULL
const MAY_BE_FALSE = 1 << IS_FALSE
const MAY_BE_TRUE = 1 << IS_TRUE
const MAY_BE_LONG = 1 << IS_LONG
const MAY_BE_DOUBLE = 1 << IS_DOUBLE
const MAY_BE_STRING = 1 << IS_STRING
const MAY_BE_ARRAY = 1 << IS_ARRAY
const MAY_BE_OBJECT = 1 << IS_OBJECT
const MAY_BE_RESOURCE = 1 << IS_RESOURCE
const MAY_BE_ANY = MAY_BE_NULL | MAY_BE_FALSE | MAY_BE_TRUE | MAY_BE_LONG | MAY_BE_DOUBLE | MAY_BE_STRING | MAY_BE_ARRAY | MAY_BE_OBJECT | MAY_BE_RESOURCE
const MAY_BE_REF = 1 << IS_REFERENCE
const MAY_BE_ARRAY_SHIFT = IS_REFERENCE
const MAY_BE_ARRAY_OF_NULL = MAY_BE_NULL << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_FALSE = MAY_BE_FALSE << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_TRUE = MAY_BE_TRUE << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_LONG = MAY_BE_LONG << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_DOUBLE = MAY_BE_DOUBLE << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_STRING = MAY_BE_STRING << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_ARRAY = MAY_BE_ARRAY << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_OBJECT = MAY_BE_OBJECT << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_RESOURCE = MAY_BE_RESOURCE << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_ANY = MAY_BE_ANY << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_OF_REF = MAY_BE_REF << MAY_BE_ARRAY_SHIFT
const MAY_BE_ARRAY_KEY_LONG = 1 << 21
const MAY_BE_ARRAY_KEY_STRING = 1 << 22
const MAY_BE_ARRAY_KEY_ANY = MAY_BE_ARRAY_KEY_LONG | MAY_BE_ARRAY_KEY_STRING
const MAY_BE_ERROR = 1 << 23
const MAY_BE_CLASS = 1 << 24
