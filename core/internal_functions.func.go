// <<generate>>

package core

func PhpRegisterInternalExtensions() int {
	return PhpRegisterExtensions(PhpBuiltinExtensions, EXTCOUNT)
}
