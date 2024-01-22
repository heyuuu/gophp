package tests

import (
	"fmt"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"strings"
)

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

	expectText := sections["EXPECT"]
	if strings.TrimSpace(output) == strings.TrimSpace(expectText) {
		return NewTestResult(tc, PASS, "", 0)
	} else {
		reason := fmt.Sprintf("output = %s, expect %s", output, expectText)
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
