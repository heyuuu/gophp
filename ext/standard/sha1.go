package standard

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/heyuuu/gophp/php/zpp"
	"io/ioutil"
)

func ZifSha1(str string, _ zpp.Opt, rawOutput bool) string {
	sum := sha1.Sum([]byte(str))

	if rawOutput {
		return string(sum[:])
	} else {
		return hex.EncodeToString(sum[:])
	}
}
func ZifSha1File(filename string, _ zpp.Opt, rawOutput bool) (string, bool) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", false
	}

	sum := sha1.Sum(data)

	if rawOutput {
		return string(sum[:]), true
	} else {
		return hex.EncodeToString(sum[:]), true
	}
}
