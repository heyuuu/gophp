// <<generate>>

package standard

import (
	"sik/core"
	g "sik/runtime/grammar"
	"sik/zend"
)

// Source: <ext/standard/levenshtein.c>

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
   | Author: Hartmut Holzgraefe <hholzgra@php.net>                        |
   +----------------------------------------------------------------------+
*/

// # include "php.h"

// # include < stdlib . h >

// # include < errno . h >

// # include < ctype . h >

// # include "php_string.h"

// #define LEVENSHTEIN_MAX_LENGTH       255

/* {{{ reference_levdist
 * reference implementation, only optimized for memory usage, not speed */

func ReferenceLevdist(s1 *byte, l1 int, s2 *byte, l2 int, cost_ins zend.ZendLong, cost_rep zend.ZendLong, cost_del zend.ZendLong) zend.ZendLong {
	var p1 *zend.ZendLong
	var p2 *zend.ZendLong
	var tmp *zend.ZendLong
	var c0 zend.ZendLong
	var c1 zend.ZendLong
	var c2 zend.ZendLong
	var i1 int
	var i2 int
	if l1 == 0 {
		return l2 * cost_ins
	}
	if l2 == 0 {
		return l1 * cost_del
	}
	if l1 > 255 || l2 > 255 {
		return -1
	}
	p1 = zend._safeEmalloc(l2+1, g.SizeOf("zend_long"), 0)
	p2 = zend._safeEmalloc(l2+1, g.SizeOf("zend_long"), 0)
	for i2 = 0; i2 <= l2; i2++ {
		p1[i2] = i2 * cost_ins
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + cost_del
		for i2 = 0; i2 < l2; i2++ {
			c0 = p1[i2] + g.Cond(s1[i1] == s2[i2], 0, cost_rep)
			c1 = p1[i2+1] + cost_del
			if c1 < c0 {
				c0 = c1
			}
			c2 = p2[i2] + cost_ins
			if c2 < c0 {
				c0 = c2
			}
			p2[i2+1] = c0
		}
		tmp = p1
		p1 = p2
		p2 = tmp
	}
	c0 = p1[l2]
	zend._efree(p1)
	zend._efree(p2)
	return c0
}

/* }}} */

func CustomLevdist(str1 *byte, str2 *byte, callback_name *byte) int {
	core.PhpErrorDocref(nil, 1<<1, "The general Levenshtein support is not there yet")

	/* not there yet */

	return -1

	/* not there yet */
}

/* }}} */

func ZifLevenshtein(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var argc int = execute_data.This.u2.num_args
	var str1 *byte
	var str2 *byte
	var callback_name *byte
	var str1_len int
	var str2_len int
	var callback_len int
	var cost_ins zend.ZendLong
	var cost_rep zend.ZendLong
	var cost_del zend.ZendLong
	var distance zend.ZendLong = -1
	switch argc {
	case 2:
		if zend.ZendParseParameters(2, "ss", &str1, &str1_len, &str2, &str2_len) == zend.FAILURE {
			return
		}
		distance = ReferenceLevdist(str1, str1_len, str2, str2_len, 1, 1, 1)
		break
	case 5:
		if zend.ZendParseParameters(5, "sslll", &str1, &str1_len, &str2, &str2_len, &cost_ins, &cost_rep, &cost_del) == zend.FAILURE {
			return
		}
		distance = ReferenceLevdist(str1, str1_len, str2, str2_len, cost_ins, cost_rep, cost_del)
		break
	case 3:
		if zend.ZendParseParameters(3, "sss", &str1, &str1_len, &str2, &str2_len, &callback_name, &callback_len) == zend.FAILURE {
			return
		}
		distance = CustomLevdist(str1, str2, callback_name)
		break
	default:
		zend.ZendWrongParamCount()
		return
	}
	if distance < 0 && execute_data.This.u2.num_args != 3 {
		core.PhpErrorDocref(nil, 1<<1, "Argument string(s) too long")
	}
	var __z *zend.Zval = return_value
	__z.value.lval = distance
	__z.u1.type_info = 4
	return
}

/* }}} */
