package tests

import (
	"fmt"
	"maps"
	"path/filepath"
	"strings"
	"time"
)

const (
	blockAll   = "ALL"
	blockTest  = "TEST"
	blockSkip  = "SKIP"
	blockClean = "CLEAN"
	blockOut   = "OUT"
	blockExp   = "EXP"
	blockDiff  = "DIFF"
)

type Config struct {
	Logger Logger

	SrcDir    string
	ExtDir    string
	Limit     int
	Workers   int
	PhpBin    string
	PhpCgiBin string

	ConfPassed    string
	IniOverwrites []string

	ShowCfg map[string]bool
	KeepCfg map[string]bool

	PassOptionN bool
	PassOptionE bool

	NoClean     bool
	Quite       bool
	TestTimeout string
	SlowMinTime time.Duration
	Verbose     bool
	X           bool
	Offline     bool
}

const ExtDir = "/__ext__"
const PhpBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php"
const PhpCgiBin = "/opt/homebrew/Cellar/php@7.4/7.4.33_6/bin/php-cgi"

func DefaultConfig() *Config {
	return &Config{
		ExtDir:    ExtDir,
		PhpBin:    PhpBin,
		PhpCgiBin: PhpCgiBin,
		ShowCfg:   map[string]bool{blockAll: true},
		KeepCfg:   map[string]bool{},
	}
}

func (c Config) IsShow(typ string) bool {
	return c.ShowCfg != nil && (c.ShowCfg[blockAll] || c.ShowCfg[typ])
}
func (c Config) IsKeep(typ string) bool {
	return c.KeepCfg != nil && (c.KeepCfg[blockAll] || c.KeepCfg[typ])
}

type TestCase struct {
	index    int
	fileName string `get:""` // 测试case名，一般为文件相对于根目录的相对路径
	filePath string `get:""` // 测试case文件绝对路径

	// 预处理路径，方便使用
	testFile   string
	testSkipif string
	testClean  string
	testPost   string

	// Case 文件解析的信息
	parsed   bool
	parseErr error
	sections map[string]string `get:""`
}

func NewTestCase(fileName string, filePath string) *TestCase {
	tc := &TestCase{fileName: fileName, filePath: filePath}
	tc.initPath()
	return tc
}
func NewTestCaseParsed(fileName string, filePath string, sections map[string]string) *TestCase {
	tc := NewTestCase(fileName, filePath)
	tc.sections = sections
	tc.parseErr = checkFileSections(filePath, sections)
	return tc
}

func (tc *TestCase) ShowName() string {
	return fmt.Sprintf("%s [%s]", tc.TestName(), tc.fileName)
}
func (tc *TestCase) TestName() string {
	if tc.sections == nil {
		return ""
	}
	return strings.TrimSpace(tc.sections["TEST"])
}

func (tc *TestCase) initPath() {
	testDir := filepath.Dir(tc.filePath)
	mainFileName := filepath.Base(tc.filePath)
	if strings.HasSuffix(mainFileName, ".phpt") {
		mainFileName = mainFileName[:len(mainFileName)-5]
	}
	tc.testFile = filepath.Join(testDir, mainFileName+".php")
	tc.testSkipif = filepath.Join(testDir, mainFileName+".skip.php")
	tc.testClean = filepath.Join(testDir, mainFileName+".clean.php")
	tc.testPost = filepath.Join(testDir, mainFileName+".post")
}

func (tc *TestCase) Parse() (map[string]string, error) {
	if !tc.parsed {
		tc.parsed = true
		tc.sections, tc.parseErr = parseTestFileSections(tc.filePath)
	}
	return tc.sections, tc.parseErr
}

type ResultType string

const (
	PASS  ResultType = "PASS"
	BORK  ResultType = "BORK"
	FAIL  ResultType = "FAIL"
	WARN  ResultType = "WARN"
	LEAK  ResultType = "LEAK"
	XFAIL ResultType = "XFAIL"
	XLEAK ResultType = "XLEAK"
	SKIP  ResultType = "SKIP"
	SLOW  ResultType = "SLOW"
)

func ValidResultType(t ResultType) bool {
	switch t {
	case PASS, BORK, FAIL, WARN, LEAK, XFAIL, XLEAK, SLOW, SKIP:
		return true
	default:
		return false
	}
}

type Result struct {
	tc      *TestCase
	types   []ResultType
	info    string
	useTime time.Duration
	output  string
	slow    bool
}

func SimpleResult(tc *TestCase, typ ResultType, info string) *Result {
	return &Result{tc: tc, types: []ResultType{typ}, info: info}
}
func PassResult(tc *TestCase, output string, useTime time.Duration) *Result {
	return &Result{tc: tc, types: []ResultType{PASS}, useTime: useTime, output: output}
}
func ComplexResult(tc *TestCase, types []ResultType, output string, info string, useTime time.Duration) *Result {
	return &Result{tc: tc, types: types, info: info, useTime: useTime, output: output}
}

func (r *Result) Case() *TestCase        { return r.tc }
func (r *Result) Types() []ResultType    { return r.types }
func (r *Result) Info() string           { return r.info }
func (r *Result) UseTime() time.Duration { return r.useTime }
func (r *Result) Output() string         { return r.output }
func (r *Result) Slow() bool             { return r.slow }

func (r *Result) MainType() ResultType {
	if len(r.types) == 0 {
		return ""
	}
	return r.types[0]
}

func (r *Result) ShowTypeNames() string {
	switch len(r.types) {
	case 0:
		return ""
	case 1:
		return string(r.types[0])
	}

	typeNames := make([]string, len(r.types))
	for i, typ := range r.types {
		typeNames[i] = string(typ)
	}
	return strings.Join(typeNames, "&")
}

type Env struct {
	m map[string]string
}

func NewEnv() *Env {
	return &Env{
		m: map[string]string{},
	}
}
func (env *Env) Clone() *Env {
	return &Env{
		m: maps.Clone(env.m),
	}
}
func (env *Env) Set(key string, value string) {
	env.m[key] = value
}
func (env *Env) SetIfEmpty(key string, value string) {
	if env.m[key] == "" {
		env.m[key] = value
	}
}
func (env *Env) Get(key string) string {
	return env.m[key]
}

type IniSettings struct {
	keys           []string
	m              map[string]string
	zendExtensions []string
	extensions     []string
}

func NewIniSettings() *IniSettings {
	return &IniSettings{
		m: map[string]string{},
	}
}

func (setting *IniSettings) Add(key string, value string) {
	switch key {
	case "zend_extension":
		setting.zendExtensions = append(setting.zendExtensions, value)
	case "extension":
		setting.extensions = append(setting.extensions, value)
	default:
		if _, exists := setting.m[key]; !exists {
			setting.keys = append(setting.keys, key)
		}
		setting.m[key] = value
	}
}
func (setting *IniSettings) Get(key string) string {
	return setting.m[key]
}

func (setting *IniSettings) Merge(ini []string) {
	for _, line := range ini {
		if name, value, ok := strings.Cut(line, "="); ok {
			name = strings.TrimSpace(name)
			value = strings.TrimSpace(value)
			setting.Add(name, value)
		}
	}
}

func (setting *IniSettings) Len() int {
	return len(setting.extensions) + len(setting.zendExtensions) + len(setting.keys)
}

func (setting *IniSettings) Each(h func(key, val string)) {
	for _, extension := range setting.extensions {
		h("extension", extension)
	}
	for _, extension := range setting.zendExtensions {
		h("zend_extension", extension)
	}
	for _, key := range setting.keys {
		h(key, setting.m[key])
	}
}

func (setting *IniSettings) ToArgs() []commandArg {
	args := make([]commandArg, 0, setting.Len()*2)
	setting.Each(func(key, val string) {
		args = append(args, commandArg{"-d", false}, commandArg{key + "=" + addslashes(val), true})
	})
	return args
}
