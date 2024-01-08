package tests

import (
	"errors"
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
	testFiles, err := findTestFiles(c.SrcDir)
	if err != nil {
		return
	}

	if limit := c.Limit; limit > 0 && len(testFiles) > limit {
		testFiles = testFiles[:limit]
	}

	testCount := len(testFiles)
	c.Events.OnAllStart(time.Now(), testCount)
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

func (c Config) initTestCase(file string) *TestCase {
	shortFileName := file
	if strings.HasPrefix(file, c.SrcDir+"/") {
		shortFileName = file[len(c.SrcDir)+1:]
	}
	return &TestCase{
		File:          file,
		ShortFileName: shortFileName,
	}
}

func (c Config) runTest(testIndex int, testFile string) (*TestResult, error) {
	tc := c.initTestCase(testFile)

	c.Events.OnTestStart(testIndex, tc)

	rawResult, err := runPhpScript(tc.File)
	if err != nil {
		return nil, err
	}

	c.Events.Log(testIndex, strings.TrimSpace(rawResult.Output))

	result := &TestResult{
		Case:     tc,
		Type:     rawResult.Type,
		TestName: rawResult.TestName,
		Reason:   rawResult.Reason,
		UseTime:  time.Duration(rawResult.UseTime),
	}
	c.Events.OnTestEnd(testIndex, tc, result)

	return result, nil
}

func findTestFiles(dir string) ([]string, error) {
	var files []string
	handler := func(file string) error {
		files = append(files, file)
		return nil
	}

	for _, subDir := range []string{"Zend", "tests", "ext", "sapi"} {
		if subDir == "ext" {
			continue
		}
		err := eachTestFiles(filepath.Join(dir, subDir), handler)
		if err != nil {
			return nil, err
		}
	}

	sortTestFiles(files, dir)

	return files, nil
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
