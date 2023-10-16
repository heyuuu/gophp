package core

/**
 * PhpExtensionLists
 */
type PhpExtensionLists struct {
	zendExtensions []string
	phpExtensions  []string
}

func (l *PhpExtensionLists) Reset() {
	l.zendExtensions = nil
	l.phpExtensions = nil
}
func (l *PhpExtensionLists) AddZendExtension(file string) {
	l.zendExtensions = append(l.zendExtensions, file)
}
func (l *PhpExtensionLists) AddPhpExtension(file string) {
	l.phpExtensions = append(l.phpExtensions, file)
}
