package tests

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	//"SKIPIF":               true,
	//"XFAIL":                true,
	//"XLEAK":                true,
	//"CLEAN":                true,
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

	// 执行测试case
	var code string
	if fileText, ok := sections["FILE"]; ok {
		code = fileText
	} else {
		return NewTestResult(tc, BORK, "no file section", 0)
	}

	output, err := runCodeBuiltin(code)
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
		return NewTestResult(tc, PASS, "", 0)
	} else {
		return NewTestResult(tc, FAIL, reason, 0)
	}
}

func runCodeBuiltin(code string) (output string, err error) {
	var buf strings.Builder
	defer func() {
		output = buf.String()
		if e := recover(); e != nil {
			err = fmt.Errorf("runCodeBuiltin() panic: %v", e)
		}
	}()

	engine := php.NewEngine()
	err = engine.Start()
	if err != nil {
		return "", err
	}

	ctx := engine.NewContext(nil, nil)
	engine.HandleContext(ctx, func(ctx *php.Context) {
		ctx.OG().PushHandler(&buf)
		fileHandle := php.NewFileHandleByString(code)
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
	)
	return replacer.Replace(s)
}

func compareExpectRegex(output string, expect string) (equals bool, reason string) {
	if regexp.MustCompile(expect).MatchString(output) {
		return true, ""
	} else if regexp.MustCompile(strings.TrimSpace(expect)).MatchString(strings.TrimSpace(output)) {
		// 目前先规避掉 phpt 换行格式导致的不匹配问题
		return true, ""
	} else {
		reason = fmt.Sprintf("output = \n%s\nexpect =\n%s\n", output, expect)
		return false, reason
	}
}
