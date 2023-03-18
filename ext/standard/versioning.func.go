// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
	"sik/zend/argparse"
	"sik/zend/types"
)

func PhpCanonicalizeVersion(version *byte) *byte {
	var len_ int = strlen(version)
	var buf *byte = zend.SafeEmalloc(len_, 2, 1)
	var q *byte
	var lp byte
	var lq byte
	var p *byte
	if len_ == 0 {
		*buf = '0'
		return buf
	}
	p = version
	q = buf
	*p++
	lp = (*p) - 1
	b.PostInc(&(*q)) = lp
	for *p {

		/*  s/[-_+]/./g;
		 *  s/([^\d\.])([^\D\.])/$1.$2/g;
		 *  s/([^\D\.])([^\d\.])/$1.$2/g;
		 */

		var isdig func(x byte) bool = func(x byte) bool { return isdigit(x) && x != '.' }
		var isndig func(x byte) bool = func(x byte) bool { return !(isdigit(x)) && x != '.' }
		var isspecialver func(x byte) bool = func(x byte) bool { return x == '-' || x == '_' || x == '+' }
		lq = *(q - 1)
		if isspecialver(*p) {
			if lq != '.' {
				b.PostInc(&(*q)) = '.'
			}
		} else if isndig(lp) && isdig(*p) || isdig(lp) && isndig(*p) {
			if lq != '.' {
				b.PostInc(&(*q)) = '.'
			}
			b.PostInc(&(*q)) = *p
		} else if !(isalnum(*p)) {
			if lq != '.' {
				b.PostInc(&(*q)) = '.'
			}
		} else {
			b.PostInc(&(*q)) = *p
		}
		*p++
		lp = (*p) - 1
	}
	b.PostInc(&(*q)) = '0'
	return buf
}
func CompareSpecialVersionForms(form1 *byte, form2 *byte) int {
	var found1 int = -1
	var found2 int = -1
	var special_forms []SpecialFormsT = []SpecialFormsT{
		MakeSpecialFormsT("dev", 0),
		MakeSpecialFormsT("alpha", 1),
		MakeSpecialFormsT("a", 1),
		MakeSpecialFormsT("beta", 2),
		MakeSpecialFormsT("b", 2),
		MakeSpecialFormsT("RC", 3),
		MakeSpecialFormsT("rc", 3),
		MakeSpecialFormsT("#", 4),
		MakeSpecialFormsT("pl", 5),
		MakeSpecialFormsT("p", 5),
		MakeSpecialFormsT(nil, 0),
	}
	var pp *SpecialFormsT
	for pp = special_forms; pp != nil && pp.GetName() != nil; pp++ {
		if strncmp(form1, pp.GetName(), strlen(pp.GetName())) == 0 {
			found1 = pp.GetOrder()
			break
		}
	}
	for pp = special_forms; pp != nil && pp.GetName() != nil; pp++ {
		if strncmp(form2, pp.GetName(), strlen(pp.GetName())) == 0 {
			found2 = pp.GetOrder()
			break
		}
	}
	return zend.ZEND_NORMALIZE_BOOL(found1 - found2)
}
func PhpVersionCompare(orig_ver1 *byte, orig_ver2 *byte) int {
	var ver1 *byte
	var ver2 *byte
	var p1 *byte
	var p2 *byte
	var n1 *byte
	var n2 *byte
	var l1 long
	var l2 long
	var compare int = 0
	if !(*orig_ver1) || !(*orig_ver2) {
		if !(*orig_ver1) && !(*orig_ver2) {
			return 0
		} else {
			if *orig_ver1 {
				return 1
			} else {
				return -1
			}
		}
	}
	if orig_ver1[0] == '#' {
		ver1 = zend.Estrdup(orig_ver1)
	} else {
		ver1 = PhpCanonicalizeVersion(orig_ver1)
	}
	if orig_ver2[0] == '#' {
		ver2 = zend.Estrdup(orig_ver2)
	} else {
		ver2 = PhpCanonicalizeVersion(orig_ver2)
	}
	n1 = ver1
	p1 = n1
	n2 = ver2
	p2 = n2
	for (*p1) && (*p2) && n1 != nil && n2 != nil {
		if b.Assign(&n1, strchr(p1, '.')) != nil {
			*n1 = '0'
		}
		if b.Assign(&n2, strchr(p2, '.')) != nil {
			*n2 = '0'
		}
		if isdigit(*p1) && isdigit(*p2) {

			/* compare element numerically */

			l1 = strtol(p1, nil, 10)
			l2 = strtol(p2, nil, 10)
			compare = zend.ZEND_NORMALIZE_BOOL(l1 - l2)
		} else if !(isdigit(*p1)) && !(isdigit(*p2)) {

			/* compare element names */

			compare = CompareSpecialVersionForms(p1, p2)

			/* compare element names */

		} else {

			/* mix of names and numbers */

			if isdigit(*p1) {
				compare = CompareSpecialVersionForms("#N#", p2)
			} else {
				compare = CompareSpecialVersionForms(p1, "#N#")
			}

			/* mix of names and numbers */

		}
		if compare != 0 {
			break
		}
		if n1 != nil {
			p1 = n1 + 1
		}
		if n2 != nil {
			p2 = n2 + 1
		}
	}
	if compare == 0 {
		if n1 != nil {
			if isdigit(*p1) {
				compare = 1
			} else {
				compare = PhpVersionCompare(p1, "#N#")
			}
		} else if n2 != nil {
			if isdigit(*p2) {
				compare = -1
			} else {
				compare = PhpVersionCompare("#N#", p2)
			}
		}
	}
	zend.Efree(ver1)
	zend.Efree(ver2)
	return compare
}
func ZifVersionCompare(executeData *zend.ZendExecuteData, return_value *types.Zval) {
	var v1 *byte
	var v2 *byte
	var op *byte = nil
	var v1_len int
	var v2_len int
	var op_len int = 0
	var compare int
	for {
		var _flags int = 0
		var _min_num_args int = 2
		var _max_num_args int = 3
		var _num_args int = executeData.NumArgs()
		var _i int = 0
		var _real_arg *types.Zval
		var _arg *types.Zval = nil
		var _expected_type argparse.ZendExpectedType = argparse.Z_EXPECTED_LONG
		var _error *byte = nil
		var _dummy types.ZendBool
		var _optional types.ZendBool = 0
		var _error_code int = argparse.ZPP_ERROR_OK
		void(_i)
		void(_real_arg)
		void(_arg)
		void(_expected_type)
		void(_error)
		void(_dummy)
		void(_optional)
		for {
			fp := argparse.FastParseStart(executeData, _min_num_args, _max_num_args, _flags)
			v1, v1_len = fp.ParseString()
			v2, v2_len = fp.ParseString()
			fp.StartOptional()
			op, op_len = fp.ParseString()
			if fp.HasError() {
				fp.HandleError()
				return
			}
			break
		}
		break
	}
	compare = PhpVersionCompare(v1, v2)
	if op == nil {
		return_value.SetLong(compare)
		return
	}
	if !(strncmp(op, "<", op_len)) || !(strncmp(op, "lt", op_len)) {
		types.ZVAL_BOOL(return_value, compare == -1)
		return
	}
	if !(strncmp(op, "<=", op_len)) || !(strncmp(op, "le", op_len)) {
		types.ZVAL_BOOL(return_value, compare != 1)
		return
	}
	if !(strncmp(op, ">", op_len)) || !(strncmp(op, "gt", op_len)) {
		types.ZVAL_BOOL(return_value, compare == 1)
		return
	}
	if !(strncmp(op, ">=", op_len)) || !(strncmp(op, "ge", op_len)) {
		types.ZVAL_BOOL(return_value, compare != -1)
		return
	}
	if !(strncmp(op, "==", op_len)) || !(strncmp(op, "=", op_len)) || !(strncmp(op, "eq", op_len)) {
		types.ZVAL_BOOL(return_value, compare == 0)
		return
	}
	if !(strncmp(op, "!=", op_len)) || !(strncmp(op, "<>", op_len)) || !(strncmp(op, "ne", op_len)) {
		types.ZVAL_BOOL(return_value, compare != 0)
		return
	}
	return_value.SetNull()
	return
}
