package replace

import (
	"fmt"
	"regexp"
	"sik/builtin/strutil"
	"sik/script/util"
)

func replaceMakeFunctionEntryDef(code string) string {
	rule := regexp.MustCompile(`(zend\.)?MakeZendFunctionEntry\((.*)\),`)
	sizeRule := regexp.MustCompile(`uint32\(b\.SizeOf\("(\w+)"\)/b.SizeOf\("struct _zend_internal_arg_info"\)-1\)`)

	result := util.RegexReplaceAll(rule, code, func(matches []string) string {
		pkgPrefix := matches[1]

		args := splitArgs(matches[2], 5)
		if len(args) < 5 {
			return matches[0]
		}

		name := args[0]
		handler := args[1]
		arg_info := args[2]
		num_args := args[3]
		flags := args[4]

		if !(arg_info == "nil" && num_args == "0") {
			sizeMatches := sizeRule.FindStringSubmatch(num_args)
			if len(sizeMatches) < 2 || !isSameName(arg_info, sizeMatches[1]) {
				fmt.Println([]string{arg_info, num_args})
				fmt.Println(sizeMatches)
				return matches[0]
			}
		}

		return fmt.Sprintf(`%sMakeZendFunctionEntryEx(%s, %s, %s, %s),`, pkgPrefix, name, flags, handler, arg_info)
	})

	return result
}

func isSameName(name string, rawName string) bool {
	if name == "nil" && rawName == "NULL" {
		return true
	}
	if name == rawName {
		return true
	}
	if name == strutil.UpperCamelCase(rawName) {
		return true
	}
	return false
}
