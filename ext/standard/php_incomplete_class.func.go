package standard

import (
	types2 "github.com/heyuuu/gophp/php/types"
)

func PHP_SET_CLASS_ATTRIBUTES(struc *types2.Zval) {
	if types2.Z_OBJCE_P(struc) == BG__().incomplete_class {
		class_name = PhpLookupClassName(struc)
		if !class_name {
			class_name = types2.NewString(INCOMPLETE_CLASS)
		}
		incomplete_class = 1
	} else {
		class_name = types2.Z_OBJCE_P(struc).GetName().Copy()
	}
}
func PHP_CLEANUP_CLASS_ATTRIBUTES() {
	// types.ZendStringReleaseEx(class_name, 0)
}
