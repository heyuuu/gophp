package zend

import (
	"github.com/heyuuu/gophp/php/types"
)

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

const ZEND_EXTENSION_API_NO = 320190902

/* Typedef's for zend_extension function pointers */

type StartupFuncT func(extension *ZendExtension) int
type ShutdownFuncT func(extension *ZendExtension)
type ActivateFuncT func()
type DeactivateFuncT func()
type MessageHandlerFuncT func(message int, arg any)
type OpArrayHandlerFuncT func(op_array *types.ZendOpArray)
type StatementHandlerFuncT func(frame *ZendExecuteData)
type FcallBeginHandlerFuncT func(frame *ZendExecuteData)
type FcallEndHandlerFuncT func(frame *ZendExecuteData)
type OpArrayCtorFuncT func(op_array *types.ZendOpArray)
type OpArrayDtorFuncT func(op_array *types.ZendOpArray)
type OpArrayPersistCalcFuncT func(op_array *types.ZendOpArray) int
type OpArrayPersistFuncT func(op_array *types.ZendOpArray, mem any) int

const ZEND_EXTMSG_NEW_EXTENSION = 1

var ZendExtensions ZendLlist

