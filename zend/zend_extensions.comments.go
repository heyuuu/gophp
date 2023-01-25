// <<generate>>

package zend

// Source: <Zend/zend_extensions.h>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/

/*
The constants below are derived from ext/opcache/ZendAccelerator.h

You can use the following macro to check the extension API version for compatibilities:

#define    ZEND_EXTENSION_API_NO_5_0_X __special__     220040412
#define    ZEND_EXTENSION_API_NO_5_1_X __special__     220051025
#define    ZEND_EXTENSION_API_NO_5_2_X __special__     220060519
#define    ZEND_EXTENSION_API_NO_5_3_X __special__     220090626
#define    ZEND_EXTENSION_API_NO_5_4_X __special__     220100525
#define    ZEND_EXTENSION_API_NO_5_5_X __special__     220121212
#define    ZEND_EXTENSION_API_NO_5_6_X __special__     220131226
#define    ZEND_EXTENSION_API_NO_7_0_X __special__     320151012

#if ZEND_EXTENSION_API_NO < ZEND_EXTENSION_API_NO_5_5_X
   // do something for php versions lower than 5.5.x
#endif
*/

/* Typedef's for zend_extension function pointers */

// Source: <Zend/zend_extensions.c>

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
   | Authors: Andi Gutmans <andi@php.net>                                 |
   |          Zeev Suraski <zeev@php.net>                                 |
   +----------------------------------------------------------------------+
*/
