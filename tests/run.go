package tests

import "fmt"

var rawSupportSections = map[string]bool{
	"TEST":   true,
	"FILE":   true,
	"EXPECT": true,

	//"EXPECTF":              true,
	//"EXPECTREGEX":          true,
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
	//"INI":                  true,
	//"ENV":                  true,
	//"EXTENSIONS":           true,
	//"SKIPIF":               true,
	//"XFAIL":                true,
	//"XLEAK":                true,
	//"CLEAN":                true,
	//"CREDITS":              true,
	//"DESCRIPTION":          true,
	//"CONFLICTS":            true,
	//"WHITESPACE_SENSITIVE": true,
}

func (c Config) runTestRealRaw(testIndex int, tc *TestCase) (*TestResult, error) {
	if !c.rawSupport(testIndex, tc) {
		return c.runTestReal(testIndex, tc)
	}

	c.Events.Log(testIndex, fmt.Sprintf("unfinished test case: %s", tc.File))

	return nil, nil
}

func (c Config) rawSupport(testIndex int, tc *TestCase) bool {
	// check sections
	for section, _ := range tc.sections {
		if !rawSupportSections[section] {
			//c.Events.Log(testIndex, fmt.Sprintf("unsupport section: %s", section))
			return false
		}
	}
	return true
}
