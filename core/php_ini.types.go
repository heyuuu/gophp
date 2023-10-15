package core

import (
	"github.com/heyuuu/gophp/zend"
)

/**
 * PhpExtensionLists
 */
type PhpExtensionLists struct {
	engine    zend.ZendLlist[*byte]
	functions zend.ZendLlist[*byte]
}

func (this *PhpExtensionLists) GetEngine() *zend.ZendLlist[*byte]    { return &this.engine }
func (this *PhpExtensionLists) GetFunctions() *zend.ZendLlist[*byte] { return &this.functions }
