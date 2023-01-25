// <<generate>>

package standard

// Source: <ext/standard/php_math.h>

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
   | Authors: Jim Winstead <jimw@php.net>                                 |
   |          Stig Sæther Bakken <ssb@php.net>                            |
   +----------------------------------------------------------------------+
*/

// #define PHP_MATH_H

/*
   WARNING: these functions are expermental: they could change their names or
   disappear in the next version of PHP!
*/

// # include < math . h >

const M_E = 2.7182817
const M_LOG2E = 1.442695
const M_LOG10E = 0.4342945
const M_LN2 = 0.6931472
const M_LN10 = 2.3025851
const M_PI = 3.1415927
const M_PI_2 = 1.5707964
const M_PI_4 = 0.7853982
const M_1_PI = 0.31830987
const M_2_PI = 0.63661975
const M_SQRTPI = 1.7724539
const M_2_SQRTPI = 1.1283792
const M_LNPI = 1.1447299
const M_EULER = 0.5772157
const M_SQRT2 = 1.4142135
const M_SQRT1_2 = 0.70710677
const M_SQRT3 = 1.7320508

/* Define rounding modes (all are round-to-nearest) */

const PHP_ROUND_HALF_UP = 0x1
const PHP_ROUND_HALF_DOWN = 0x2
const PHP_ROUND_HALF_EVEN = 0x3
const PHP_ROUND_HALF_ODD = 0x4
