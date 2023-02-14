package replace

import (
	"fmt"
	"regexp"
	"sik/script/util"
	"strings"
)

func ReplaceMakeArgInfo(code string) string {
	rule := regexp.MustCompile(`(zend\.)?MakeZendInternalArgInfo\((.*)\),`)
	intNameRule := regexp.MustCompile(`\(\*byte\)\(zend_uintptr_t\((-?\d+)\)\)`)

	result := util.RegexReplaceAll(rule, code, func(matches []string) string {
		pkgPrefix := matches[1]
		args := splitArgs(matches[2], 4)
		name := args[0]
		typ := args[1]
		byRef := args[2]
		variadic := args[3]

		var makeArgs []string

		// name
		isNameStr, _ := regexp.MatchString(`"\w+"`, name)
		if isNameStr {
			makeArgs = append(makeArgs, name)
		} else if match := intNameRule.FindStringSubmatch(name); match != nil {
			nameInt := match[1]
			makeArgs = append(makeArgs, nameInt)
		} else {
			panic("未支持的 name 参数: " + name)
		}

		// typ
		if typ != "0" {
			makeArgs = append(makeArgs, `ArgInfoType(`+typ+`)`)
		}

		// byRef
		if byRef != "0" && byRef != "ZEND_RETURN_VALUE" && byRef != "zend.ZEND_RETURN_VALUE" {
			makeArgs = append(makeArgs, `ArgInfoByRef(`+byRef+`)`)
		}

		// variadic
		if variadic != "0" {
			makeArgs = append(makeArgs, `ArgInfoVariadic()`)
		}

		// 组成 Make 方法调用
		if isNameStr {
			return fmt.Sprintf(`%sMakeArgInfo(%s)`, pkgPrefix, strings.Join(makeArgs, ","))
		} else {
			return fmt.Sprintf(`%sMakeArgInfoSpecial(%s)`, pkgPrefix, strings.Join(makeArgs, ","))
		}
	})

	return result
}
