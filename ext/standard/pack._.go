// <<generate>>

package standard

// Source: <ext/standard/pack.h>

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
   | Author: Rasmus Lerdorf <rasmus@lerdorf.on.ca>                        |
   +----------------------------------------------------------------------+
*/

// Source: <ext/standard/pack.c>

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
   | Author: Chris Schneider <cschneid@relog.ch>                          |
   +----------------------------------------------------------------------+
*/

/* Whether machine is little endian */

var MachineLittleEndian byte

/* Mapping of byte from char (8bit) to long for machine endian */

var ByteMap []int

/* Mappings of bytes from int (machine dependent) to int for machine endian */

var IntMap []int

/* Mappings of bytes from shorts (16bit) for all endian environments */

var MachineEndianShortMap []int
var BigEndianShortMap []int
var LittleEndianShortMap []int

/* Mappings of bytes from longs (32bit) for all endian environments */

var MachineEndianLongMap []int
var BigEndianLongMap []int
var LittleEndianLongMap []int

/* Mappings of bytes from quads (64bit) for all endian environments */

var MachineEndianLonglongMap []int
var BigEndianLonglongMap []int
var LittleEndianLonglongMap []int

/* {{{ php_pack
 */
