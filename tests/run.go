package tests

import (
	"encoding/hex"
	"fmt"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"github.com/heyuuu/gophp/php/perr"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"unicode/utf8"
)

var rawSupportSections = map[string]bool{
	"TEST":        true,
	"FILE":        true,
	"EXPECT":      true,
	"EXPECTF":     true,
	"EXPECTREGEX": true,
	//"EXPECTHEADERS":        true,
	//"POST":                 true,
	//"POST_RAW":             true,
	//"GZIP_POST":            true,
	//"DEFLATE_POST":         true,
	//"PUT":                  true,
	//"GET":                  true,
	//"COOKIE":               true,
	//"ARGS":                 true,
	//"REDIRECTTEST":         true,
	//"CAPTURE_STDIO":        true,
	//"STDIN":                true,
	//"CGI":                  true,
	//"PHPDBG":               true,
	"INI": true,
	//"ENV":                  true,
	//"EXTENSIONS":           true,
	"SKIPIF": true,
	//"XFAIL":                true,
	//"XLEAK":                true,
	"CLEAN":       true,
	"CREDITS":     true,
	"DESCRIPTION": true,
	//"CONFLICTS":            true,
	//"WHITESPACE_SENSITIVE": true,
}

func RunTestFile(testIndex int, testName string, testFile string) (result *TestResult) {
	// 解析测试文件(.phpt文件)
	tc, err := ParseTestFile(testName, testFile)
	if err != nil {
		return NewTestResult(tc, BORK, "parse test case failed", 0)
	}

	sections := tc.Sections

	// 限制目前支持的字段
	var unsupportedSections []string
	for name, _ := range sections {
		if !rawSupportSections[name] {
			unsupportedSections = append(unsupportedSections, name)
		}
	}
	if len(unsupportedSections) > 0 {
		return NewTestResult(tc, SKIP, "unsupported section: "+strings.Join(unsupportedSections, ", "), 0)
	}

	// todo INI 段
	ini := sections["INI"]
	if ini != "" {
		iniReplacer := strings.NewReplacer(
			"{PWD}", filepath.Dir(testFile),
			"{TMP}", os.TempDir(),
		)
		ini = iniReplacer.Replace(ini)
	}

	// 判断是否 SKIP
	if skipIfText, ok := sections["SKIPIF"]; ok {
		output, err := runCodeBuiltin(skipIfText, ini)
		if err != nil {
			return NewTestResult(tc, FAIL, "run SKIPIF code filed: "+err.Error(), 0)
		}
		if output != "" {
			return NewTestResult(tc, SKIP, output, 0)
		}
	}

	// 执行测试case
	var code string
	if fileText, ok := sections["FILE"]; ok {
		code = fileText
	} else {
		return NewTestResult(tc, BORK, "no file section", 0)
	}

	output, err := runCodeBuiltin(code, ini)
	if err != nil {
		return NewTestResult(tc, FAIL, "run code failed: "+err.Error(), 0)
	}
	output = strings.ReplaceAll(output, "\r\n", "\n")

	var pass bool
	var reason string
	if expectText, ok := sections["EXPECT"]; ok {
		pass, reason = compareExpect(output, expectText)
	} else if expectFormatText, ok := sections["EXPECTF"]; ok {
		pass, reason = compareExpectFormat(output, expectFormatText)
	} else if expectRegexText, ok := sections["EXPECTREGEX"]; ok {
		pass, reason = compareExpectRegex(output, expectRegexText)
	}

	if pass {
		result = NewTestResult(tc, PASS, "", 0)
	} else {
		result = NewTestResult(tc, FAIL, reason, 0)
	}
	result.Output = output
	return result
}

func runCodeBuiltin(code string, ini string) (output string, err error) {
	var buf strings.Builder
	defer func() {
		output = buf.String()
		if e := recover(); e != nil && e != perr.ErrExit {
			err = fmt.Errorf("runCodeBuiltin() panic: %v", e)

			// 打印堆栈
			const size = 64 << 10
			stack := make([]byte, size)
			stack = stack[:runtime.Stack(stack, false)]
			log.Printf(">>> runCodeBuiltin() panic: %v\n%s", e, stack)
		}
	}()

	engine := php.NewEngine()
	engine.BaseCtx().INI().AppendIniEntries(ini)
	err = engine.Start()
	if err != nil {
		return "", err
	}

	ctx := engine.NewContext(nil, nil)
	engine.HandleContext(ctx, func(ctx *php.Context) {
		ctx.OG().PushHandler(&buf)
		fileHandle := php.NewFileHandleByCommandLine(code)
		_, err = php.ExecuteScript(ctx, fileHandle, false)
	})
	return
}

func compareExpect(output string, expect string) (equals bool, reason string) {
	if strings.TrimSpace(output) == strings.TrimSpace(expect) {
		return true, ""
	} else {
		reason = fmt.Sprintf("output = \n%s\nexpect =\n%s\n", output, expect)
		return false, reason
	}
}

func compareExpectFormat(output string, expect string) (equals bool, reason string) {
	expect = convertExpectFormat2Regex(expect)
	return compareExpectRegex(output, expect)
}

func convertExpectFormat2Regex(s string) string {
	// do preg_quote, but miss out any %r delimited sections
	var buf strings.Builder
	length := len(s)
	for offset := 0; offset < length; {
		var start, end int
		if start = strpos(s, "%r", offset); start >= 0 {
			// we have found a start tag
			end = strpos(s, "%r", start+2)
			if end < 0 {
				// unbalanced tag, ignore it.
				start = length
				end = length
			}
		} else {
			// no more %r sections
			start = length
			end = length
		}
		// quote a non re portion of the string
		buf.WriteString(pregQuote(s[offset:start]))
		if end > start {
			buf.WriteByte('(')
			buf.WriteString(s[start+2 : end])
			buf.WriteByte(')')
		}
		offset = end + 2
	}
	s = buf.String()

	// Stick to basics
	replacer := strings.NewReplacer(
		"%e", string([]byte{'\\', filepath.Separator}),
		"%s", `[^\r\n]+`,
		"%S", `[^\r\n]*`,
		"%a", `.+`,
		"%A", `.*`,
		"%w", `\s*`,
		"%i", `[+-]?\d+`,
		"%d", `\d+`,
		"%x", `[0-9a-fA-F]+`,
		"%f", `[+-]?\.?\d+\.?\d*(?:[Ee][+-]?\d+)?`,
		"%c", `.`,
		"\x00", "\\0",
	)
	return replacer.Replace(s)
}

func compareExpectRegex(output string, expect string) (equals bool, reason string) {
	equals, err := compareExpectRegexInternal(output, expect)
	if err != nil {
		return false, err.Error()
	}
	if equals {
		return true, ""
	}

	// 目前先规避掉 phpt 换行格式导致的不匹配问题
	equals, err = compareExpectRegexInternal(strings.TrimSpace(output), strings.TrimSpace(expect))
	if err != nil {
		return false, err.Error()
	}
	if equals {
		return true, ""
	}

	// 匹配失败
	reason = fmt.Sprintf("output = \n%s\nexpect =\n%s\n", output, expect)
	return false, reason
}

func compareExpectRegexInternal(output string, expect string) (equals bool, err error) {
	if !utf8.ValidString(expect) {
		expect = utf8SafeString(expect)
		output = utf8SafeString(output)
	}

	rule, err := regexp.Compile(expect)
	if err != nil {
		return false, fmt.Errorf("EXPECTREGEX rule parse fail, err = %w, expect = %s\n", err, expect)
	}
	return rule.MatchString(output), nil
}

func utf8SafeString(s string) string {
	if utf8.ValidString(s) {
		return s
	}
	var buf strings.Builder
	for _, r := range s {
		if utf8.ValidRune(r) {
			buf.WriteRune(r)
		} else {
			buf.WriteString(`\x`)
			buf.WriteString(hex.EncodeToString([]byte(string([]rune{r}))))
		}
	}
	return buf.String()
}
