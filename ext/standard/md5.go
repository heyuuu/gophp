package standard

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/heyuuu/gophp/php/zpp"
	"io/ioutil"
)

func ZifMd5(str string, _ zpp.Opt, rawOutput bool) string {
	sum := md5.Sum([]byte(str))

	if rawOutput {
		return string(sum[:])
	} else {
		return hex.EncodeToString(sum[:])
	}
}
func ZifMd5File(filename string, _ zpp.Opt, rawOutput bool) (string, bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", false
	}

	sum := md5.Sum(data)

	if rawOutput {
		return string(sum[:]), true
	} else {
		return hex.EncodeToString(sum[:]), true
	}
}
