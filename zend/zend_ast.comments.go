// <<generate>>

package zend

// Source: <Zend/zend_ast.h>

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
   | Authors: Bob Weinand <bwoebi@php.net>                                |
   |          Dmitry Stogov <dmitry@php.net>                              |
   |          Nikita Popov <nikic@php.net>                                |
   +----------------------------------------------------------------------+
*/

/* Same as zend_ast, but with children count, which is updated dynamically */

/* Lineno is stored in val.u2.lineno */

/* Separate structure for function and class declaration, as they need extra information. */

// Source: <Zend/zend_ast.c>

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
   | Authors: Bob Weinand <bwoebi@php.net>                                |
   |          Dmitry Stogov <dmitry@php.net>                              |
   +----------------------------------------------------------------------+
*/

/*
 * Operator Precedence
 * ====================
 * priority  associativity  operators
 * ----------------------------------
 *   10     left            include, include_once, eval, require, require_once
 *   20     left            ,
 *   30     left            or
 *   40     left            xor
 *   50     left            and
 *   60     right           print
 *   70     right           yield
 *   80     right           =>
 *   85     right           yield from
 *   90     right           = += -= *= /= .= %= &= |= ^= <<= >>= **=
 *  100     left            ? :
 *  110     right           ??
 *  120     left            ||
 *  130     left            &&
 *  140     left            |
 *  150     left            ^
 *  160     left            &
 *  170     non-associative == != === !==
 *  180     non-associative < <= > >= <=>
 *  190     left            << >>
 *  200     left            + - .
 *  210     left            * / %
 *  220     right           !
 *  230     non-associative instanceof
 *  240     right           + - ++ -- ~ (type) @
 *  250     right           **
 *  260     left            [
 *  270     non-associative clone new
 */
