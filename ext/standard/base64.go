package standard

import (
	"encoding/base64"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="

func PhpBase64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func PhpBase64DecodeEx(str string, strict bool) (string, bool) {
	cleanStr, ok := base64CleanStr(str, strict)
	if !ok {
		return "", false
	}

	// notice: 此处通过过滤 padding + RawStdEncoding 来兼容有/无padding的情况
	cleanStr = strings.TrimRight(cleanStr, string(base64.StdPadding))
	ret, err := base64.RawStdEncoding.DecodeString(cleanStr)
	if err != nil {
		return "", false
	}
	return string(ret), true
}
func base64CleanStr(str string, strict bool) (string, bool) {
	var buf strings.Builder
	for _, c := range []byte(str) {
		// todo 此处可优化为 mask 数组 (参考 strings.asciiSet)
		if strings.IndexByte(base64Table, c) >= 0 {
			buf.WriteByte(c)
		} else if c == ' ' {
			// skip spaces
		} else if strict {
			return "", false
		}
	}
	return buf.String(), true
}

func ZifBase64Encode(str string) string {
	return PhpBase64Encode(str)
}
func ZifBase64Decode(str string, _ zpp.Opt, strict bool) (string, bool) {
	return PhpBase64DecodeEx(str, strict)
}
