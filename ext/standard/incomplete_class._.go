package standard

import (
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

const INCOMPLETE_CLASS_MSG string = "The script tried to execute a method or " + "access a property of an incomplete object. " + "Please ensure that the class definition \"%s\" of the object " + "you are trying to operate on was loaded _before_ " + "unserialize() gets called or provide an autoloader " + "to load the class definition"

var PhpIncompleteObjectHandlers zend.ZendObjectHandlers

/* {{{ incomplete_class_message
 */
