package tests

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

type scriptArgs struct {
	TestFile string `json:"test_file"`
}

type scriptResponse struct {
	Ok    bool         `json:"ok"`
	Error string       `json:"error"`
	Data  scriptResult `json:"data"`
}

type scriptResult struct {
	Type     ResultType `json:"type"`
	TestName string     `json:"testName"`
	Reason   string     `json:"reason"`
	UseTime  int        `json:"useTime"`
	Output   string     `json:"output"`
}

const phpBin = "/opt/homebrew/bin/php"
const scriptPath = "/Users/heyu/Code/sik/php-tests/script.php"

func runPhpScript(testFile string) (*scriptResult, error) {
	args := scriptArgs{
		TestFile: testFile,
	}
	argsJson, err := encodeScriptArgs(args)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(phpBin, scriptPath, argsJson)
	//log.Println("Run Cmd: " + cmd.String())
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	//log.Println("Run Cmd Result: " + string(output))

	result, err := decodeScriptOutput(output)
	if err != nil {
		return nil, err
	}

	// 校验 type
	if !ValidResultType(result.Type) {
		return nil, fmt.Errorf(`unexpected type value: "%s"`, result.Type)
	}

	// 校验 output
	if strings.HasPrefix(result.Output, "base64:") {
		result.Output, err = base64Decode(result.Output[len("base64:"):])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func encodeScriptArgs(v scriptArgs) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func decodeScriptOutput(output []byte) (*scriptResult, error) {
	var response scriptResponse
	err := json.Unmarshal(output, &response)
	if err != nil {
		return nil, err
	}
	if response.Ok {
		return &response.Data, nil
	} else {
		return nil, errors.New(response.Error)
	}
}

func base64Decode(output string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(output)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
