package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

func PhpClassAttributes(struc *types.Zval) (className string, incompleteClass bool) {
	ce := struc.Object().GetCe()

	if ce == BG__().incomplete_class {
		if className, ok := PhpLookupClassName(struc); ok {
			return className, true
		} else {
			return INCOMPLETE_CLASS, true
		}
	} else {
		return ce.Name(), false
	}
}
