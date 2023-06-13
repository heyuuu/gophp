package standard

import (
	"github.com/heyuuu/gophp/php/types"
)

func PHP_SET_CLASS_ATTRIBUTES(struc *types.Zval) {
	if types.Z_OBJCE_P(struc) == BG__().incomplete_class {
		class_name = PhpLookupClassName(struc)
		if !class_name {
			class_name = types.NewString(INCOMPLETE_CLASS)
		}
		incomplete_class = 1
	} else {
		class_name = types.Z_OBJCE_P(struc).GetName().Copy()
	}
}
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
