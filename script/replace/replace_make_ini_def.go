package replace

import (
	"fmt"
	"regexp"
	"sik/script/util"
	"strings"
)

func replaceMakeIniEntryDef(code string) string {
	rule := regexp.MustCompile(`MakeZendIniEntryDef\((.*)\)`)

	result := util.RegexReplaceAll(rule, code, func(matches []string) string {
		args := splitArgs(matches[1], 8)
		name := args[0]
		onModify := args[1]
		arg1 := args[2]
		arg2 := args[3]
		arg3 := args[4]
		value := args[5]
		displayer := args[6]
		modifiable := args[7]

		str := fmt.Sprintf(`*NewZendIniEntryDef(%s, %s)`, name, modifiable)
		if value != "nil" {
			str += ".Value(" + value + ")"
		}
		if displayer != "nil" {
			str += ".Displayer(" + displayer + ")"
		}
		if onModify != "nil" {
			str += fmt.Sprintf(".OnModifyArgs(\n%s, %s, %s, %s,\n)", onModify, arg1, arg2, arg3)
		}
		return str
	})

	return result
}

func splitArgs(str string, n int) []string {
	args := make([]string, 0, n)

	var buf strings.Builder
	nest := 0
	for _, c := range []byte(str) {
		if c == ',' && nest == 0 {
			args = append(args, strings.TrimSpace(buf.String()))
			buf.Reset()
		} else {
			buf.WriteByte(c)
			if c == '(' {
				nest++
			} else if c == ')' {
				nest--
			}
		}
	}
	if len(args) < n && buf.Len() > 0 {
		args = append(args, strings.TrimSpace(buf.String()))
		buf.Reset()
	}

	return args
}
