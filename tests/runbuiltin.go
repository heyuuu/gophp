package tests

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/php"
	_ "github.com/heyuuu/gophp/php/boot"
	"strings"
)

var ErrUnsupportedSection = errors.New("unsupported section")

func RunCaseBuiltin(testIndex int, testName string, testFile string) (*TestResult, error) {
	// 解析测试文件(.phpt文件)
	tc, err := parseTestCase(testName, testFile)
	if err != nil {
		return nil, err
	}
	sections := tc.sections

	// 限制目前支持的字段
	for name, _ := range sections {
		if !rawSupportSections[name] {
			return nil, ErrUnsupportedSection
		}
	}

	var code string
	if fileText, ok := sections["FILE"]; ok {
		code = fileText
	} else {
		return NewTestResult(tc, BORK, "", 0), nil
	}

	output, err := runCodeBuiltin(code)
	if err != nil {
		return nil, err
	}

	expectText := sections["EXPECT"]
	if strings.TrimSpace(output) == strings.TrimSpace(expectText) {
		return NewTestResult(tc, PASS, "", 0), nil
	} else {
		reason := fmt.Sprintf("output = %s, expect %s", output, expectText)
		return NewTestResult(tc, FAIL, reason, 0), nil
	}
}

func runCodeBuiltin(code string) (output string, err error) {
	var buf strings.Builder
	defer func() {
		output = buf.String()
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
