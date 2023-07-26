package standard

import (
	b "github.com/heyuuu/gophp/builtin"
	"github.com/heyuuu/gophp/php/lang"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/zend"
	"github.com/heyuuu/gophp/zend/zpp"
	"strings"
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
	lang.PostInc(&(*q)) = lp
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
				lang.PostInc(&(*q)) = '.'
			}
		} else if isndig(lp) && isdig(*p) || isdig(lp) && isndig(*p) {
			if lq != '.' {
				lang.PostInc(&(*q)) = '.'
			}
			lang.PostInc(&(*q)) = *p
		} else if !(isalnum(*p)) {
			if lq != '.' {
				lang.PostInc(&(*q)) = '.'
			}
		} else {
			lang.PostInc(&(*q)) = *p
		}
		*p++
		lp = (*p) - 1
	}
	lang.PostInc(&(*q)) = '0'
	return buf
}
func CompareSpecialVersionForms(form1 string, form2 string) int {
	var found1 int = -1
	var found2 int = -1
	var specialForms = map[string]int{
		"dev":   0,
		"alpha": 1,
		"a":     1,
		"beta":  2,
		"b":     2,
		"RC":    3,
		"rc":    3,
		"#":     4,
		"pl":    5,
		"p":     5,
	}
	for name, order := range specialForms {
		if strings.HasPrefix(form1, name) {
			found1 = order
			break
		}
	}
	for name, order := range specialForms {
		if strings.HasPrefix(form2, name) {
			found2 = order
			break
		}
	}
	return b.Compare(found1, found2)
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
		if lang.Assign(&n1, strchr(p1, '.')) != nil {
			*n1 = '0'
		}
		if lang.Assign(&n2, strchr(p2, '.')) != nil {
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
func ZifVersionCompare(ver1 *types.Zval, ver2 *types.Zval, _ zpp.Opt, oper *string) *types.Zval {
	compare := PhpVersionCompare(ver1, ver2)
	if oper == nil {
		return types.NewZvalLong(compare)
	}

	var result bool
	switch *oper {
	case "<", "lt":
		result = compare < 0
	case "<=", "le":
		result = compare <= 0
	case ">", "gt":
		result = compare > 0
	case ">=", "ge":
		result = compare >= 0
	case "==", "=", "eq":
		result = compare == 0
	case "!=", "<>", "ne":
		result = compare != 0
	default:
		return types.NewZvalNull()
	}
	return types.NewZvalBool(result)
}
