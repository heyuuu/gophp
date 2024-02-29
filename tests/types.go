package tests

import (
	"fmt"
	"github.com/heyuuu/gophp/shim/maps"
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
	DumpRoot string
	Limit    int
	Workers  int
	Logger   Logger

	SrcDir string
	ExtDir string

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
	Asan        bool
}

func DefaultConfig() Config {
	return Config{
		Limit:         -1,
		ShowCfg:       map[string]bool{blockAll: true},
		KeepCfg:       make(map[string]bool),
		IniOverwrites: baseIniOverwrites,
	}
}

type TestCase struct {
	index         int
	file          string
	shortFileName string

	// 预处理路径，方便使用
	testFile   string
	testSkipif string
	testClean  string
	testPost   string

	// Case 文件解析的信息
	sections map[string]string
	testName string
}

func NewTestCase(index int, file string, shortFileName string) *TestCase {
	tc := &TestCase{
		index:         index,
		file:          file,
		shortFileName: shortFileName,
	}
	tc.initPath()

	return tc
}

func (tc *TestCase) ShowName() string {
	return fmt.Sprintf("%s [%s]", tc.testName, tc.shortFileName)
}

func (tc *TestCase) initPath() {
	testDir := filepath.Dir(tc.file)
	mainFileName := filepath.Base(tc.file)
	if strings.HasSuffix(mainFileName, ".phpt") {
		mainFileName = mainFileName[:len(mainFileName)-5]
	}
	tc.testFile = filepath.Join(testDir, mainFileName+".php")
	tc.testSkipif = filepath.Join(testDir, mainFileName+".skip.php")
	tc.testClean = filepath.Join(testDir, mainFileName+".clean.php")
	tc.testPost = filepath.Join(testDir, mainFileName+".post")
}

func (tc *TestCase) parse() error {
	sections, err := parseTestFileSections(tc.file)
	if err != nil {
		return err
	}

	tc.sections = sections
	tc.testName = strings.TrimSpace(sections["TEST"])
	return nil
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

func SimpleResult(tc *TestCase, typ ResultType, info string, useTime time.Duration) *Result {
	return &Result{tc: tc, types: []ResultType{typ}, info: info, useTime: useTime}
}
func ComplexResult(tc *TestCase, types []ResultType, info string, useTime time.Duration) *Result {
	return &Result{tc: tc, types: types, info: info, useTime: useTime}
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

func (setting *IniSettings) ToParams() string {
	var buf strings.Builder
	for _, extension := range setting.extensions {
		_, _ = fmt.Fprintf(&buf, ` -d "%s=%s"`, "extension", addslashes(extension))
	}
	for _, extension := range setting.zendExtensions {
		_, _ = fmt.Fprintf(&buf, ` -d "%s=%s"`, "zend_extension", addslashes(extension))
	}
	for _, key := range setting.keys {
		_, _ = fmt.Fprintf(&buf, ` -d "%s=%s"`, key, addslashes(setting.m[key]))
	}
	return buf.String()
}
