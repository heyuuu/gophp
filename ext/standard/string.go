package standard

import (
	"github.com/heyuuu/gophp/builtin/ascii"
	"github.com/heyuuu/gophp/zend/types"
)

func PhpStringToupper(s *types.String) *types.String {
	return types.NewString(ascii.StrToUpper(s.GetStr()))
}

func PhpStringTolower(s *types.String) *types.String {
	return types.NewString(ascii.StrToLower(s.GetStr()))
}

func substr(str string, offset int, length *int) (string, bool) {
	negativeOffset := offset < 0
	rawLen := len(str)
	// handle offset
	if offset > len(str) {
		return "", false
	} else if offset < 0 {
		/* if "offset" position is negative, count start position from the end
		 * of the string
		 */
		offset += len(str)
		if offset < 0 {
			offset = 0
		}
	}
	if offset > 0 {
		str = str[offset:]
	}

	// handle length
	if length != nil {
		l := *length
		if l < 0 {
			/* if "length" position is negative, set it to the length
			 * needed to stop that many chars from the end of the string
			 */
			if -l > len(str) {
				if negativeOffset && -l <= rawLen {
					l = 0
				} else {
					return "", false
				}
			} else {
				l = len(str) + l
			}
		}

		if l < len(str) {
			str = str[:l]
		}
	}

	return str, true
}
