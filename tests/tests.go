package tests

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func Run(conf Config) error {
	return conf.run()
}

type Config struct {
	SrcDir  string
	Events  EventHandler
	Limit   int
	Workers int

	PhpBin     string
	PhpCgiBin  string
	PhpDbgBin  string
	TempSource string
	TempTarget string

	PassOption    string
	IniOverwrites []IniEntry

	ExtensionDir string

	IsWin bool
}

func (c Config) run() (err error) {
	testFiles, extSkipped, ignoreByExt, err := findTestFiles(c.SrcDir)
	if err != nil {
		return
	}

	if limit := c.Limit; limit > 0 && len(testFiles) > limit {
		testFiles = testFiles[:limit]
	}

	testCount := len(testFiles)
	c.Events.OnAllStart(time.Now(), testCount, extSkipped, ignoreByExt)
	if c.Workers > 1 {
		err = c.parallelRunTests(testFiles, c.Workers)
	} else {
		err = c.simpleRunTests(testFiles)
	}
	if err != nil {
		return err
	}

	c.Events.OnAllEnd(time.Now())
	return nil
}

func (c Config) simpleRunTests(testFiles []string) error {
	for i, file := range testFiles {
		if _, err := c.runTest(i, file); err != nil {
			return err
		}
	}
	return nil
}

func (c Config) parallelRunTests(testFiles []string, limit int) (err error) {
	if limit <= 1 {
		return errors.New("parallelRunTests 的并发数 limit 必须大于 1")
	}

	var wg sync.WaitGroup
	wg.Add(len(testFiles))

	// 用于限制并发数
	var limitChan = make(chan struct{}, limit)
	defer close(limitChan)

	// 遍历任务
	for i, testFile := range testFiles {
		limitChan <- struct{}{}
		go func(index int, file string) {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("parallelRunTests painc: %v\n", e)
				}
				wg.Done()
				<-limitChan
			}()

			c.runTest(index, file)
		}(i, testFile)
	}

	wg.Wait()

	return nil
}

func (c Config) runTest(testIndex int, testFile string) (*TestResult, error) {
	testName := testFile
	if strings.HasPrefix(testFile, c.SrcDir+"/") {
		testName = testFile[len(c.SrcDir)+1:]
	}

	tc, err := parseTestCase(testName, testFile)
	if err != nil {
		c.Events.Log(testIndex, "parse test case error: "+err.Error())
		c.Events.OnTestEnd(testIndex, tc, nil)
		return nil, fmt.Errorf("Parse test case file failed: file=%s, err=%w", testFile, err)
	}

	c.Events.OnTestStart(testIndex, tc)
	//result, runErr := c.runTestReal(testIndex, tc)
	result, runErr := c.runTestRealRaw(testIndex, tc)
	c.Events.OnTestEnd(testIndex, tc, result)

	return result, runErr
}

func (c Config) runTestReal(testIndex int, tc *TestCase) (*TestResult, error) {
	rawResult, err := runPhpScript(tc.File)
	if err != nil {
		c.Events.Log(testIndex, "run php script error: "+err.Error())
		return nil, err
	}

	c.Events.Log(testIndex, strings.TrimSpace(rawResult.Output))

	result := &TestResult{
		Case:    tc,
		Type:    rawResult.Type,
		Reason:  rawResult.Reason,
		UseTime: time.Duration(rawResult.UseTime),
	}

	return result, nil
}

func findTestFiles(dir string) (files []string, extSkipped int, ignoreByExt int, err error) {
	// main tests
	for _, subDir := range []string{"Zend", "tests", "sapi"} {
		err = EachTestFile(filepath.Join(dir, subDir), true, func(file string) error {
			files = append(files, file)
			return nil
		})
		if err != nil {
			return
		}
	}
	sortTestFiles(files, dir)

	// ext tests
	extRoot := filepath.Join(dir, "ext")
	exts, err := os.ReadDir(extRoot)
	for _, ext := range exts {
		if ext.IsDir() {
			extSkipped++
			err = EachTestFile(filepath.Join(extRoot, ext.Name()), true, func(file string) error {
				ignoreByExt++
				return nil
			})
			if err != nil {
				return
			}
		}
	}

	return
}
