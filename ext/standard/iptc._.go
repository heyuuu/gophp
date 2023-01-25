// <<generate>>

package standard

// Source: <ext/standard/iptc.c>

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
   | Author: Thies C. Arntzen <thies@thieso.net>                          |
   +----------------------------------------------------------------------+
*/

/* some defines for the different JPEG block types */

/* {{{ php_iptc_put1
 */

var Psheader []byte = "xFFxED00Photoshop 3.008BIMx04x040000"

/* {{{ proto array iptcembed(string iptcdata, string jpeg_file_name [, int spool])
   Embed binary IPTC data into a JPEG image. */
