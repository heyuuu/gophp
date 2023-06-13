package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

func PhpClassAttributes(struc *types.Zval) (className string, incompleteClass bool) {
	ce := struc.Object().GetCe()

	if ce == BG__().incomplete_class {
		classNameZstr := PhpLookupClassName(struc)
		if classNameZstr != nil {
			return classNameZstr.GetStr(), true
		} else {
			return INCOMPLETE_CLASS, true
		}
	} else {
		return ce.Name(), false
	}
}
