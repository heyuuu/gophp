package phpparse

import (
	"fmt"
	"gophp/php/ast"
	"io/ioutil"
	"os/exec"
)

func parseToJsonFile(script string, filename string, output string) error {
	command := exec.Command("php", script, "-s", filename, "-o", output)
	result, err := command.CombinedOutput()
	fmt.Println(string(result))
	if err != nil {
		fmt.Println("命令执行失败: " + err.Error())
		return err
	}

	return nil
}

func loadAstJson(jsonFile string) ([]ast.Stmt, error) {
	binData, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return nil, err
	}

	return decodeAstData(binData)
}
