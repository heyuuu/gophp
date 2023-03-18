// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend/types"
)

func PHP_SET_CLASS_ATTRIBUTES(struc *types.Zval) {
	if types.Z_OBJCE_P(struc) == BG__().incomplete_class {
		class_name = PhpLookupClassName(struc)
		if !class_name {
			class_name = types.ZendStringInit(INCOMPLETE_CLASS, b.SizeOf("INCOMPLETE_CLASS")-1, 0)
		}
		incomplete_class = 1
	} else {
		class_name = types.Z_OBJCE_P(struc).GetName().Copy()
	}
}
func PHP_CLEANUP_CLASS_ATTRIBUTES() { types.ZendStringReleaseEx(class_name, 0) }
