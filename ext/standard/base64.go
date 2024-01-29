package standard

import (
	"encoding/base64"
	"github.com/heyuuu/gophp/php/zpp"
	"strings"
)

const base64padding = '='
const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func PhpBase64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
func PhpBase64Decode(str string, strict bool) (string, bool) {
	if !strict {
		str = base64CleanStrWeak(str)
	} else if cleanStr, ok := base64CleanStrStrict(str); ok {
		str = cleanStr
	} else {
		return "", false
	}

	// notice: padding 以被过滤，此处使用 RawStdEncoding 而非 StdEncoding
	ret, err := base64.RawStdEncoding.DecodeString(str)
	if err != nil {
		return "", false
	}
	return string(ret), true
}
func base64CleanStrWeak(str string) string {
	var buf strings.Builder
	buf.Grow(len(str))
	for _, c := range []byte(str) {
		if strings.IndexByte(base64Table, c) >= 0 {
			buf.WriteByte(c)
		}
	}
	str = buf.String()

	if len(str)%4 == 1 {
		str = str[:len(str)-1]
	}
	return str
}

func base64CleanStrStrict(str string) (string, bool) {
	var buf strings.Builder

	var padding int
	for _, c := range []byte(str) {
		if c == base64padding {
			padding++
			continue
		}

		/* skip whitespace */
		if strings.IndexByte("\t\r\n ", c) >= 0 {
			continue
		}
		/* fail if any data follows padding */
		if padding > 0 {
			return "", false
		}
		if strings.IndexByte(base64Table, c) >= 0 {
			buf.WriteByte(c)
		} else {
			/* fail on bad characters */
			return "", false
		}
	}

	str = buf.String()

	/* fail if the input is truncated (only one char in last group) */
	if len(str)%4 == 1 {
		return "", false
	}

	/* fail if the padding length is wrong (not VV==, VVV=), but accept zero padding
	 * RFC 4648: "In some circumstances, the use of padding [--] is not required" */
	if padding > 0 && (padding > 2 || (len(str)+padding)%4 != 0) {
		return "", false
	}

	return str, true
}

func ZifBase64Encode(str string) string {
	return PhpBase64Encode(str)
}
func ZifBase64Decode(str string, _ zpp.Opt, strict bool) (string, bool) {
	return PhpBase64Decode(str, strict)
}
