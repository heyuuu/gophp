// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/core"
	"sik/zend"
)

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
	if l1 > LEVENSHTEIN_MAX_LENGTH || l2 > LEVENSHTEIN_MAX_LENGTH {
		return -1
	}
	p1 = zend.SafeEmalloc(l2+1, b.SizeOf("zend_long"), 0)
	p2 = zend.SafeEmalloc(l2+1, b.SizeOf("zend_long"), 0)
	for i2 = 0; i2 <= l2; i2++ {
		p1[i2] = i2 * cost_ins
	}
	for i1 = 0; i1 < l1; i1++ {
		p2[0] = p1[0] + cost_del
		for i2 = 0; i2 < l2; i2++ {
			c0 = p1[i2] + b.Cond(s1[i1] == s2[i2], 0, cost_rep)
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
	zend.Efree(p1)
	zend.Efree(p2)
	return c0
}
func CustomLevdist(str1 *byte, str2 *byte, callback_name *byte) int {
	core.PhpErrorDocref(nil, zend.E_WARNING, "The general Levenshtein support is not there yet")

	/* not there yet */

	return -1

	/* not there yet */
}
func ZifLevenshtein(execute_data *zend.ZendExecuteData, return_value *zend.Zval) {
	var argc int = zend.ZEND_NUM_ARGS()
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
		zend.WRONG_PARAM_COUNT
	}
	if distance < 0 && zend.ZEND_NUM_ARGS() != 3 {
		core.PhpErrorDocref(nil, zend.E_WARNING, "Argument string(s) too long")
	}
	zend.RETVAL_LONG(distance)
	return
}
