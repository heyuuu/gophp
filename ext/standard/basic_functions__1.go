// <<generate>>

package standard

import (
	g "sik/runtime/grammar"
	"sik/zend"
)

/* }}} */

/* {{{ proto array sys_getloadavg()
 */

func ZifSysGetloadavg(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var load []float64
	if g.CondF2(execute_data.This.u2.num_args == 0, zend.SUCCESS, func() zend.ZEND_RESULT_CODE {
		zend.ZendWrongParametersNoneError()
		return zend.FAILURE
	}) == zend.FAILURE {
		return
	}
	if getloadavg(load, 3) == -1 {
		return_value.u1.type_info = 2
		return
	} else {
		var __arr *zend.ZendArray = zend._zendNewArray(0)
		var __z *zend.Zval = return_value
		__z.value.arr = __arr
		__z.u1.type_info = 7 | 1<<0<<8 | 1<<1<<8
		zend.AddIndexDouble(return_value, 0, load[0])
		zend.AddIndexDouble(return_value, 1, load[1])
		zend.AddIndexDouble(return_value, 2, load[2])
	}
}

/* }}} */
