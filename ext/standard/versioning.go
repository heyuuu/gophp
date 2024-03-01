package standard

import (
	"cmp"
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php/types"
	"github.com/heyuuu/gophp/php/zpp"
	"strconv"
	"strings"
)

func splitVersion(version string) []string {
	if version == "" {
		return nil
	} else if version[0] == '#' {
		return strings.Split(version, ".")
	}

	isDigit := func(c byte) bool { return ascii.IsDigit(c) }
	isNotDigit := func(c byte) bool { return !ascii.IsDigit(c) && c != '.' }
	isSpecial := func(c byte) bool { return c == '-' || c == '_' || c == '+' }

	var segments []string
	var start int = 0
	var nextStart int = 0
	for i, c := range []byte(version) {
		if i == 0 {
			continue
		}

		lastC := version[i-1]
		if isSpecial(c) {
			nextStart = i + 1
		} else if isNotDigit(lastC) && isDigit(c) || isDigit(lastC) && isNotDigit(c) {
			nextStart = i
		} else if !ascii.IsAlphaNum(c) {
			nextStart = i + 1
		} else {
			continue
		}

		// split
		if i > start {
			segments = append(segments, version[start:i])
		}
		start = nextStart
	}
	if start < len(version) {
		segments = append(segments, version[start:])
	}
	return segments
}

func canonicalizeVersion(version string) string {
	if version == "" || version[0] == '#' {
		return version
	}

	isDigit := func(c byte) bool { return ascii.IsDigit(c) }
	isNotDigit := func(c byte) bool { return !ascii.IsDigit(c) && c != '.' }
	isSpecial := func(c byte) bool { return c == '-' || c == '_' || c == '+' }

	var buf strings.Builder
	buf.Grow(len(version) * 2)
	for i, c := range []byte(version) {
		if i == 0 {
			buf.WriteByte(c)
			continue
		}

		lastC := version[i-1]
		lastQ := buf.String()[buf.Len()-1]
		if isSpecial(c) {
			if lastQ != '.' {
				buf.WriteByte('.')
			}
		} else if isNotDigit(lastC) && isDigit(c) || isDigit(lastC) && isNotDigit(c) {
			if lastQ != '.' {
				buf.WriteByte('.')
			}
			buf.WriteByte(c)
		} else if !ascii.IsAlphaNum(c) {
			if lastQ != '.' {
				buf.WriteByte('.')
			}
		} else {
			buf.WriteByte(c)
		}
	}
	return buf.String()
}
func CompareSpecialVersionForms(form1 string, form2 string) int {
	var found1 = -1
	var found2 = -1
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
	return cmp.Compare(found1, found2)
}
func PhpVersionCompare(ver1 string, ver2 string) int {
	if ver1 == "" || ver2 == "" {
		return strings.Compare(ver1, ver2)
	}

	var result = 0
	segments1 := splitVersion(ver1)
	segments2 := splitVersion(ver2)
	for i := 0; i < len(segments1) || i < len(segments2); i++ {
		seg1, seg2 := "#N#", "#N#"
		if i < len(segments1) {
			seg1 = segments1[i]
		}
		if i < len(segments2) {
			seg2 = segments2[i]
		}
		if ascii.IsDigit(seg1[0]) && ascii.IsDigit(seg2[0]) {
			/* compare element numerically */
			l1, _ := strconv.Atoi(seg1)
			l2, _ := strconv.Atoi(seg2)
			result = cmp.Compare(l1, l2)
		} else if !ascii.IsDigit(seg1[0]) && !ascii.IsDigit(seg2[0]) {
			/* compare element names */
			result = CompareSpecialVersionForms(seg1, seg2)
		} else {
			/* mix of names and numbers */
			if ascii.IsDigit(seg1[0]) {
				result = CompareSpecialVersionForms("#N#", seg2)
			} else {
				result = CompareSpecialVersionForms(seg1, "#N#")
			}
		}
		if result != 0 {
			break
		}
	}
	return result
}
func ZifVersionCompare(ver1 string, ver2 string, _ zpp.Opt, oper *string) *types.Zval {
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
