package standard

import (
	"github.com/heyuuu/gophp/php"
	"strings"
)

/**
 * Constants
 */

const STR_PAD_LEFT = 0
const STR_PAD_RIGHT = 1
const STR_PAD_BOTH = 2
const PHP_PATHINFO_DIRNAME = 1
const PHP_PATHINFO_BASENAME = 2
const PHP_PATHINFO_EXTENSION = 4
const PHP_PATHINFO_FILENAME = 8
const PHP_PATHINFO_ALL = PHP_PATHINFO_DIRNAME | PHP_PATHINFO_BASENAME | PHP_PATHINFO_EXTENSION | PHP_PATHINFO_FILENAME

const _HEB_BLOCK_TYPE_ENG = 1
const _HEB_BLOCK_TYPE_HEB = 2

const (
	CHAR_MAX    = 127
	LC_CTYPE    = 2
	LC_NUMERIC  = 4
	LC_TIME     = 5
	LC_COLLATE  = 1
	LC_MONETARY = 3
	LC_ALL      = 0
)

func RegisterStringConstants(ctx *php.Context, moduleNumber int) {
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_LEFT", php.Long(STR_PAD_LEFT))
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_RIGHT", php.Long(STR_PAD_RIGHT))
	php.RegisterConstant(ctx, moduleNumber, "STR_PAD_BOTH", php.Long(STR_PAD_BOTH))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_DIRNAME", php.Long(PHP_PATHINFO_DIRNAME))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_BASENAME", php.Long(PHP_PATHINFO_BASENAME))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_EXTENSION", php.Long(PHP_PATHINFO_EXTENSION))
	php.RegisterConstant(ctx, moduleNumber, "PATHINFO_FILENAME", php.Long(PHP_PATHINFO_FILENAME))

	/* If last members of struct lconv equal CHAR_MAX, no grouping is done */

	php.RegisterConstant(ctx, moduleNumber, "CHAR_MAX", php.Long(CHAR_MAX))
	php.RegisterConstant(ctx, moduleNumber, "LC_CTYPE", php.Long(LC_CTYPE))
	php.RegisterConstant(ctx, moduleNumber, "LC_NUMERIC", php.Long(LC_NUMERIC))
	php.RegisterConstant(ctx, moduleNumber, "LC_TIME", php.Long(LC_TIME))
	php.RegisterConstant(ctx, moduleNumber, "LC_COLLATE", php.Long(LC_COLLATE))
	php.RegisterConstant(ctx, moduleNumber, "LC_MONETARY", php.Long(LC_MONETARY))
	php.RegisterConstant(ctx, moduleNumber, "LC_ALL", php.Long(LC_ALL))
}

func ZifUtf8Encode(data string) string {
	var buf strings.Builder
	for _, c := range []byte(data) {
		if c < 0x80 {
			buf.WriteByte(c)
		} else {
			buf.WriteByte(0xc0 | c>>6)
			buf.WriteByte(0x80 | c&0x3f)
		}
	}
	return buf.String()
}
