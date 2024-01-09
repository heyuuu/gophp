package tests

import (
	"errors"
	"fmt"
	"github.com/heyuuu/gophp/shim/cmp"
	"github.com/heyuuu/gophp/shim/slices"
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
	tc, err := parseTestCase(testFile, c.SrcDir)
	if err != nil {
		c.Events.OnTestEnd(testIndex, tc, nil)
		return nil, fmt.Errorf("Parse test case file failed: file=%s, err=%w", testFile, err)
	}

	c.Events.OnTestStart(testIndex, tc)
	result, runErr := c.runTestReal(testIndex, tc)
	c.Events.OnTestEnd(testIndex, tc, result)

	return result, runErr
}

func (c Config) runTestReal(testIndex int, tc *TestCase) (*TestResult, error) {
	rawResult, err := runPhpScript(tc.File)
	if err != nil {
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
		err = eachTestFiles(filepath.Join(dir, subDir), func(file string) error {
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
			err = eachTestFiles(filepath.Join(extRoot, ext.Name()), func(file string) error {
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

func eachTestFiles(dir string, handle func(file string) error) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.Name() == "" || file.Name()[0] == '.' {
			continue
		}

		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			err = eachTestFiles(path, handle)
		} else {
			if strings.HasSuffix(path, ".phpt") {
				err = handle(path)
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func sortTestFiles(files []string, srcDir string) {
	runTestDir := filepath.Join(srcDir, "tests/run-test")
	testDir := filepath.Join(srcDir, "tests")
	scorer := func(file string) int {
		if strings.HasPrefix(file, runTestDir) {
			return 2
		} else if strings.HasPrefix(file, testDir) {
			return 1
		}
		return 0
	}
	slices.SortStableFunc(files, func(file1, file2 string) int {
		score1, score2 := scorer(file1), scorer(file2)
		if score1 == score2 {
			return cmp.Compare(file1, file2)
		} else {
			return -cmp.Compare(score1, score2)
		}
	})
}
