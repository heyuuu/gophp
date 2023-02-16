package replace

import (
	"fmt"
	"regexp"
	"sik/builtin/strutil"
	"sik/script/util"
)

func replaceMakeFunctionEntryDef(code string) string {
	rule := regexp.MustCompile(`(zend\.)?MakeZendFunctionEntry\((.*)\),`)
	sizeRule := regexp.MustCompile(`uint32\(b\.SizeOf\("(arginfo_zend__void)"\)/b.SizeOf\("struct _zend_internal_arg_info"\)-1\)`)

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

		sizeMatches := sizeRule.FindStringSubmatch(num_args)
		if len(sizeMatches) == 0 || strutil.UpperCamelCase(sizeMatches[1]) != arg_info {
			return matches[0]
		}

		return fmt.Sprintf(`%sMakeZendFunctionEntry(%s, %s, %s, %s),`, pkgPrefix, name, flags, handler, arg_info)
	})

	return result
}
