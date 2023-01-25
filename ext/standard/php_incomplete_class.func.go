// <<generate>>

package standard

import (
	b "sik/builtin"
	"sik/zend"
)

func PHP_SET_CLASS_ATTRIBUTES(struc *zend.Zval) {
	if zend.Z_OBJCE_P(struc) == BG(incomplete_class) {
		class_name = PhpLookupClassName(struc)
		if !class_name {
			class_name = zend.ZendStringInit(INCOMPLETE_CLASS, b.SizeOf("INCOMPLETE_CLASS")-1, 0)
		}
		incomplete_class = 1
	} else {
		class_name = zend.ZendStringCopy(zend.Z_OBJCE_P(struc).name)
	}
}
func PHP_CLEANUP_CLASS_ATTRIBUTES() { zend.ZendStringReleaseEx(class_name, 0) }
