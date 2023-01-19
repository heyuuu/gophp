// <<generate>>

package zend

// Source: <Zend/zend_float.h>

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
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   +----------------------------------------------------------------------+
*/

// #define ZEND_FLOAT_H

/*
  Define functions for FP initialization and de-initialization.
*/

/* Copy of the contents of xpfpa.h (which is under public domain)
   See http://wiki.php.net/rfc/rounding for details.

   Cross Platform Floating Point Arithmetics

   This header file defines several platform-dependent macros that ensure
   equal and deterministic floating point behaviour across several platforms,
   compilers and architectures.

   The current macros are currently only used on x86 and x86_64 architectures,
   on every other architecture, these macros expand to NOPs. This assumes that
   other architectures do not have an internal precision and the operhand types
   define the __special__  computational precision of floating point operations. This
   assumption may be false, in that case, the author is interested in further
   details on the other platform.

   For further details, please visit:
   http://www.christian-seiler.de/projekte/fpmath/

   Version: 20090317 */

/*
  This is either not an x87 FPU or the inline assembly syntax was not
  recognized. In any case, default to NOPs for the macros and hope the
  generated code will behave as planned.
*/

// #define XPFPA_DECLARE

// #define XPFPA_HAVE_CW       0

// #define XPFPA_CW_DATATYPE       unsigned int

// #define XPFPA_STORE_CW(variable)

// #define XPFPA_RESTORE_CW(variable)

// #define XPFPA_SWITCH_DOUBLE()

// #define XPFPA_SWITCH_SINGLE()

// #define XPFPA_SWITCH_DOUBLE_EXTENDED()

// #define XPFPA_RESTORE()

// #define XPFPA_RETURN_DOUBLE(val) return ( val )

// #define XPFPA_RETURN_SINGLE(val) return ( val )

// #define XPFPA_RETURN_DOUBLE_EXTENDED(val) return ( val )

// Source: <Zend/zend_float.c>

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
   | Authors: Christian Seiler <chris_se@gmx.net>                         |
   +----------------------------------------------------------------------+
*/

// # include "zend.h"

// # include "zend_compile.h"

// # include "zend_float.h"

func ZendInitFpu() { EG.SetSavedFpuCwPtr(nil) }

/* }}} */

func ZendShutdownFpu() { EG.SetSavedFpuCwPtr(nil) }

/* }}} */

func ZendEnsureFpuMode() {}

/* }}} */
