package core

import (
	"sik/zend"
)

/**
 * PhpExtensionLists
 */
type PhpExtensionLists struct {
	engine    zend.ZendLlist
	functions zend.ZendLlist
}

func (this *PhpExtensionLists) GetEngine() zend.ZendLlist    { return this.engine }
func (this *PhpExtensionLists) GetFunctions() zend.ZendLlist { return this.functions }
