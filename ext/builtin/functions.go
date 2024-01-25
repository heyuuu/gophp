package builtin

import (
	"github.com/heyuuu/gophp/kits/ascii"
	"github.com/heyuuu/gophp/php"
	"github.com/heyuuu/gophp/php/perr"
	"github.com/heyuuu/gophp/php/types"
	"strings"
)

func ZifGcMemCaches() int     { return 0 }
func ZifGcCollectCycles() int { return 0 }
func ZifGcEnabled() bool      { return true }

func ZifGcStatus() *types.Array {
	arr := types.NewArrayCap(4)
	arr.AddAssocLong("runs", 0)
	arr.AddAssocLong("collected", 0)
	arr.AddAssocLong("threshold", 0)
	arr.AddAssocLong("roots", 0)
	return arr
}

func ZifStrlen(str string) int               { return len(str) }
func ZifStrcmp(str1 string, str2 string) int { return strings.Compare(str1, str2) }
func ZifStrncmp(ctx *php.Context, str1 string, str2 string, len_ int) (int, bool) {
	if len_ < 0 {
		php.Error(ctx, perr.E_WARNING, "Length must be greater than or equal to 0")
		return 0, false
	}
	if len(str1) > len_ {
		str1 = str1[:len_]
	}
	if len(str2) > len_ {
		str2 = str2[:len_]
	}
	return strings.Compare(str1, str2), true
}
func ZifStrcasecmp(str1 string, str2 string) int {
	return ascii.StrCaseCompare(str1, str2)
}